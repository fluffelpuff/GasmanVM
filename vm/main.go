package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fluffelpuff/GasmanVM/imagefile"
)

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

	// Es wird versucht die Image Datei einzulesen
	image_file, err := imagefile.LoadImageFile(imageFilePath)
	if err != nil {
		panic(err)
	}
	defer image_file.Close()

	// Es wird geprüft ob das Image ein Filesystem Image benötigt

	// Die VM wird erzeugt
	vm, err := NewRuntime(image_file)
	if err != nil {
		panic(err)
	}
	if err := vm.Run(); err != nil {
		panic(err)
	}
}
