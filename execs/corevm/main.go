package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	coreclientbridge "github.com/fluffelpuff/GasmanVM/core_service/core_client_bridge"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	"github.com/fluffelpuff/GasmanVM/vm"
)

var disableSigCheck bool
var imageFilePath string
var fileSysImage string
var debugMode bool
var showInfo bool

func main() {
	// Es wird versucht die Verbindung mit dem Core Service herzustellen
	coreControllerBridge, err := coreclientbridge.OpenBridgeConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Die Flags werden ausgewertet
	flag.BoolVar(&disableSigCheck, "disableSigCheck", false, "Beschreibung der Flagge")
	flag.BoolVar(&showInfo, "showInfoPrint", false, "Beschreibung der Flagge")
	flag.StringVar(&fileSysImage, "fsroot", "", "Geben Sie Ihren Namen ein")
	flag.StringVar(&imageFilePath, "imageFile", "", "Beschreibung der Flagge")
	flag.BoolVar(&debugMode, "debug", false, "Beschreibung der Flagge")
	flag.Parse()

	// Die Erlaubten Modi werden abgerufen
	allowed_modes := strings.Split(os.Getenv("GASMANVM_ALLOWED_MODES"), ",")
	if len(allowed_modes) == 0 {
		allowed_modes = append(allowed_modes, "service")
		allowed_modes = append(allowed_modes, "app")
	}

	// VM name and version
	if showInfo {
		fmt.Println("GasmanVM v 0.1.1")
	}

	// Sollte kein Pfad vorhanden sein, wird geprüft ob eine index.gimg Datei vorhanden ist
	if len(imageFilePath) == 0 {
		// Versuchen Sie, Informationen zur Datei abzurufen
		_, err := os.Stat("index.gimg")

		// Überprüfen Sie, ob ein Fehler aufgetreten ist
		if os.IsNotExist(err) {
			fmt.Println("Keine Javascript Container Image Datei vorhanden")
			return
		}

		// Der Path wird Aktualisiert
		imageFilePath = "index.gimg"
	}

	// Info
	if showInfo {
		fmt.Printf("Container image file '%s' is loading...\n", imageFilePath)
	}

	// Es wird versucht die Image Datei einzulesen
	image_file, err := imagefile.LoadImageFile(imageFilePath, showInfo)
	if err != nil {
		panic(err)
	}
	defer image_file.Close()

	// Es wird ermittelt ob der Modus des Images zulässig ist
	allowed_mode := false
	for _, i := range allowed_modes {
		if i == image_file.GetManifest().Application.GasmanVmMode {
			allowed_mode = true
			break
		}
	}
	if !allowed_mode {
		fmt.Println("Not allowed image mode")
		return
	}

	// Es wird geprüft ob die Signatu prüfung deaktiviert wurde
	if !disableSigCheck {
		metad := image_file.GetMetaData()
		if metad.PackageCreator == nil {
			fmt.Printf("The image '%s' cannot be opened, it has no issuer information.\n", imageFilePath)
			return
		}
	}

	// Das Manifest wird abgerufen
	manifest := image_file.GetManifest()

	// Der CoreController wird Vorbereitet
	if err := coreControllerBridge.Setup(manifest); err != nil {
		fmt.Println(err)
		return
	}

	// Sollte das Image ein Virtuelles Dateisystem benötigen und sollte keines vorhanden sein, wird der Vorgang abgebrochen
	var fsSysContainer *fsysfile.FileSystem
	if manifest.Application.GasmanVmFilesysimage != "" || len(manifest.Application.GasmanVmFilesysimage) != 0 {
		if manifest.Application.GasmanVmFilesysimage != "no" {
			if len(fileSysImage) == 0 {
				if os.Getenv("GASMANVM_SIMULATE_FILE_SYSTEM") != "yes" {
					fmt.Println("No virtual file system was passed for the image, the process is aborted because a virtual file system is required by your image.")
					return
				}
			}
		}
	}

	// Sofern vorhanden wird das Dateisystem geladen
	if len(fileSysImage) > 0 {
		// Es wird versucht das Dateisystem einzulesen
		fsSysContainer, err = fsysfile.OpenVirtualFileSystemContainer(fileSysImage)
		if err != nil {
			panic(err)
		}
	}

	// Sollte der DisableSigCheck verwendet werden, muss der Benutzer die Ausführung des Paketes Manuell bestätigen
	if disableSigCheck && os.Getenv("GASMANVM_ALLOW_UNSECURE") != "yes" {
		// Die Informationen zu dem Image werden angzeigt
		fmt.Println("Details about the image you opened:")
		fmt.Println(" -> Name\t\t\t : " + manifest.Package.GasmanVmPackage)
		fmt.Println(" -> Version\t\t\t : " + manifest.Package.GasmanVmVersion)
		fmt.Println(" -> Label\t\t\t : " + manifest.Application.GasmanVmLabel)
		fmt.Println(" -> Description\t\t\t : " + manifest.Application.GasmanVmDescription)
		fmt.Println(" -> Running mode\t\t : " + manifest.Application.GasmanVmMode)
		fmt.Println(" -> Main source file\t\t : " + manifest.Application.GasmanVmSource)
		fmt.Println(" -> File system needed\t\t : " + manifest.Application.GasmanVmFilesysimage)
		fmt.Println(" -> Required permissions\t # ")

		// Die Einzelnen Berechtigungen werden angezeigt
		total_prints := 0
		for _, i := range manifest.Permissions {
			// Es wird ermittelt ob diese Berechtigung zulässig ist
			total_prints++
			foundit, gasmnaservice := validatePermission(i.GasmanVmName)
			if foundit && gasmnaservice {
				fmt.Println("\t-> " + i.GasmanVmName)
			} else if foundit && !gasmnaservice {
				fmt.Println("\t-> " + i.GasmanVmName)
				fmt.Println("Aborted, needed GasmanVM Core Service")
				return
			} else {
				fmt.Println("\t-> " + i.GasmanVmName)
				fmt.Println("Aborted, unkown permission")
				return
			}
		}
		if total_prints < 1 {
			fmt.Println("None")
		}

		// F+ge eine Leere Zeile hinzu
		fmt.Println()

		// Die Gruppen werden angezeigt
		fmt.Println(manifest.Groups)

		// Es wird auf die Bestätigung durch den Benutzer gewartet
		if !vm.YesOrNoTextEnter("Are you sure you want to open this package? (yes/No) # ") {
			fmt.Println("Aborted.")
			return
		}
	}

	// Die VM wird erzeugt
	vm, err := vm.NewRuntime(fsSysContainer, image_file, coreControllerBridge)
	if err != nil {
		panic(err)
	}

	// Die Anendung wird bereitgestellt
	if err := coreControllerBridge.Provide(); err != nil {
		panic(err)
	}

	// Führt die Runtime aus
	if err := vm.Run(); err != nil {
		panic(err)
	}

	// Es wird gewartet bis die Anwendung fertig ist
	vm.Wait()
}

func validatePermission(permission_name string) (bool, bool) {
	switch permission_name {
	case "INTERNET":
		return true, true
	case "LOCAL_USER_FOLDERS":
		return true, true
	case "CAMERA":
		return true, true
	case "MICROPHONE":
		return true, true
	case "EXTERNAL_IPC":
		return true, true
	case "HOST_INFORMATIONS":
		return true, true
	case "XSERVERSTREAM":
		return true, true
	case "WIN32UISTREAM":
		return true, true
	default:
		return false, false
	}
}
