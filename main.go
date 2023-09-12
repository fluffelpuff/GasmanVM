package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/dop251/goja"
)

func printFunc(call goja.FunctionCall) goja.Value {
	// Extrahieren Sie die Argumente, die an die print-Funktion übergeben wurden
	args := make([]interface{}, len(call.Arguments))
	for i, arg := range call.Arguments {
		args[i] = arg.Export()
	}

	// Ausgabe der Argumente auf der Go-Seite
	fmt.Println(args...)

	// Rückgabe eines leeren Werts
	return goja.Undefined()
}

var (
	kernel32DLL          = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitleW = kernel32DLL.NewProc("SetConsoleTitleW")
)

func setConsoleTitle(title string) error {
	ptrTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}

	ret, _, err := procSetConsoleTitleW.Call(uintptr(unsafe.Pointer(ptrTitle)))
	if ret == 0 {
		return err
	}

	return nil
}

type FilePlayground struct {
	Path       string `json:"path"`
	Playground string `json:"playground"`
}

type Manifest struct {
	MainFile        string            `json:"main"`
	FilePlaygrounds []*FilePlayground `json:"file_playground"`
}

func main() {
	// Es wird geprüft ob ein Parameter vorhanden ist
	var imageFilePath string
	var debugMode bool
	var consoleIo bool
	flag.BoolVar(&consoleIo, "enable_console_io", false, "Beschreibung der Flagge")
	flag.StringVar(&imageFilePath, "imgf", "", "Beschreibung der Flagge")
	flag.BoolVar(&debugMode, "debug", false, "Beschreibung der Flagge")
	flag.Parse()

	// Sollte kein Pfad vorhanden sein, wird geprüft ob eine index.jscnt Datei vorhanden ist
	if len(imageFilePath) == 0 {
		// Versuchen Sie, Informationen zur Datei abzurufen
		_, err := os.Stat("index.jscnt")

		// Überprüfen Sie, ob ein Fehler aufgetreten ist
		if os.IsNotExist(err) {
			fmt.Println("Keine Javascript Container Image Datei vorhanden")
			return
		}

		// Der Path wird Aktualisiert
		imageFilePath = "index.jscnt"
	}

	// Öffnen Sie die ZIP-Datei
	zipFile, err := zip.OpenReader(imageFilePath)
	if err != nil {
		log.Fatalf("Fehler beim Öffnen der ZIP-Datei: %v", err)
	}
	defer zipFile.Close()

	// Die Manifestdatei wird ermittelt
	var manifestFile []byte
	for _, file := range zipFile.File {
		if file.Name == ".manifest" {
			// Öffnen und lesen Sie die JavaScript-Datei
			jsFile, err := file.Open()
			if err != nil {
				log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
			}
			defer jsFile.Close()

			manifestFile, err = io.ReadAll(jsFile)
			if err != nil {
				log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
			}
			break
		}
	}

	// Die Manifest Datei wird eingelesen
	var manifest Manifest
	if err := json.Unmarshal(manifestFile, &manifest); err != nil {
		fmt.Println("Fehler beim Dekodieren der JSON-Daten:", err)
		return
	}

	// Es wird ermittelt ob alle benötigten Dateien vorhanden sind
	for _, otem := range manifest.FilePlaygrounds {
		found_file := false
		for _, file := range zipFile.File {
			if string("src/"+otem.Path) == file.Name {
				found_file = true
				break
			}
		}
		if !found_file {
			fmt.Println("Ungülite Container Datei")
			return
		}
	}

	// Es wird ermittelt ob die Datei Hashes korrekt sind
	for _, otem := range zipFile.File {
		if strings.HasPrefix(otem.Name, "src/") {
			found_file := false
			for _, file := range manifest.FilePlaygrounds {
				if otem.Name == string("src/"+file.Path) {
					found_file = true
					break
				}
			}
			if !found_file {
				fmt.Println("Ungülite Container Datei")
				return
			}
		}
	}

	// Es wird geprüft ob die Manifestdatei korrekt ist
	if len(manifestFile) == 0 {
		log.Fatalf("JavaScript-Datei nicht gefunden: %v", "index.js")
	}

	// Die Startdatei wird geladen
	var indexFileBytes []byte
	for _, file := range zipFile.File {
		if file.Name == string("src/"+manifest.MainFile) {
			// Öffnen und lesen Sie die JavaScript-Datei
			jsFile, err := file.Open()
			if err != nil {
				log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
			}
			defer jsFile.Close()

			indexFileBytes, err = io.ReadAll(jsFile)
			if err != nil {
				log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
			}
			break
		}
	}

	// Erstellen Sie eine neue Goja-Laufzeit
	vm := goja.New()

	// Require wird überschrieben
	vm.Set("require", nil)

	// Setzt den Titel der Console
	if consoleIo {
		vm.Set("setTitle", func(value string) {
			setConsoleTitle(value)
		})
	}

	// Überschreiben der `require`-Funktion in Goja
	vm.Set("include", func(call goja.FunctionCall) goja.Value {
		// Holen Sie den Dateinamen, der in require übergeben wurde
		if len(call.Arguments) < 1 {
			log.Print("Fehler: Kein Dateiname in require übergeben")
			return goja.Null()
		}
		filename := call.Argument(0).String()

		// Es wird ermittelt ob es sich um ein Datei oder ein Sysbib Import handelt
		if strings.HasPrefix(filename, "#") {
			return nil
		}

		// Suchen Sie die Datei im ZIP-Archiv
		var content []byte
		for _, file := range zipFile.File {
			if file.Name == string("src/"+filename) {
				fileReader, err := file.Open()
				if err != nil {
					log.Printf("Fehler beim Öffnen der Datei %s im ZIP-Archiv: %v", filename, err)
					return goja.Null()
				}
				defer fileReader.Close()

				content, err = ioutil.ReadAll(fileReader)
				if err != nil {
					log.Printf("Fehler beim Lesen der Datei %s aus dem ZIP-Archiv: %v", filename, err)
					return goja.Null()
				}
				break
			}
		}

		// Erstellen Sie eine neue Goja-Laufzeit für die Module
		moduleVM := goja.New()
		if consoleIo {
			moduleVM.Set("print", printFunc)
		}
		moduleVM.Set("exports", moduleVM.NewObject())

		// Führen Sie den geladenen JavaScript-Code in der neuen Laufzeit aus
		_, err := moduleVM.RunString(string(content))
		if err != nil {
			log.Printf("Fehler beim Ausführen des JavaScript-Codes aus %s: %v", filename, err)
			return goja.Null()
		}

		// Die Exports werden zurückgegeben
		return moduleVM.Get("exports")
	})

	// Überschreiben der `print`-Funktion in Goja
	if consoleIo {
		vm.Set("print", printFunc)
	}

	// Führen Sie den JavaScript-Code in der index.js-Datei aus
	_, err = vm.RunString(string(indexFileBytes))
	if err != nil {
		log.Fatalf("Fehler beim Ausführen des JavaScript-Codes: %v", err)
	}
}
