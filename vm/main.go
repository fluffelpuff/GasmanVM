package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fluffelpuff/GasmanVM/imagefile"
)

func main() {
	// Es wird geprüft ob ein Parameter vorhanden ist
	var useBackendService bool
	var disableOpenGL bool
	var disableUI bool
	var imageFilePath string
	var fileSysImage string
	var debugMode bool
	var consoleIo bool
	flag.BoolVar(&disableOpenGL, "disable_opengl", false, "Beschreibung der Flagge")
	flag.BoolVar(&disableUI, "disable_ui", false, "Beschreibung der Flagge")
	flag.BoolVar(&useBackendService, "no_host_service", false, "Beschreibung der Flagge")
	flag.BoolVar(&consoleIo, "enable_console_io", false, "Beschreibung der Flagge")
	flag.StringVar(&fileSysImage, "fsimage", "", "Geben Sie Ihren Namen ein")
	flag.StringVar(&imageFilePath, "imgf", "", "Beschreibung der Flagge")
	flag.BoolVar(&debugMode, "debug", false, "Beschreibung der Flagge")
	flag.Parse()

	// Es wird geprüft ein Backend Service vorhanden ist

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

	// Es wird versucht die Image Datei einzulesen
	image_file, err := imagefile.LoadImageFile(imageFilePath)
	if err != nil {
		panic(err)
	}
	defer image_file.Close()

	// Es wird ermittelt ob eine FilesysImage verwendet werden soll
	if len(fileSysImage) > 0 {
		// Es wird geprüft ob das Image ein Filesystem benötigt
		if image_file.NeedFileSystem() {
			panic("")
		}
	}

	// Die VM wird erzeugt
	vm, err := NewRuntime(image_file)
	if err != nil {
		panic(err)
	}

	// Führt die Runtime aus
	if err := vm.Run(); err != nil {
		panic(err)
	}
}
