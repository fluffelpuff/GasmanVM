package jsengine

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
)

// Initalisiert ein Javascript basiertes Modules
func (o *JSEngine) runModuleJSScript(script_bytes []byte) (*goja.Runtime, error) {
	// Erstellen Sie eine neue Goja-Laufzeit für die Module
	moduleVM := goja.New()

	// Die Standard Funktionen werden Aktiviert
	o.initRuntimeBaseFunctions(moduleVM)

	// Die Export Vairable wird erzeugt
	moduleVM.Set("exports", moduleVM.NewObject())

	// Führen Sie den geladenen JavaScript-Code in der neuen Laufzeit aus
	_, err := moduleVM.RunString(string(script_bytes))
	if err != nil {
		return nil, fmt.Errorf("JSEngine:runModuleJSScript" + err.Error())
	}

	// Das Ergebniss wird zurückgegeben
	return moduleVM, nil
}

// Führt das Mainscript aus
func (o *JSEngine) RunMainScript() (goja.Value, error) {
	// Es wird georpüft ob Bereits ein Main Script ausgefüht wird
	if len(o.jsInterpreters) != 0 {
		return nil, fmt.Errorf("JSEngine>RunMainScript: always running main")
	}

	// Die Imagedatei wird abgerufen
	image_file := o.motherVM.GetImageFile()
	if image_file == nil {
		return nil, fmt.Errorf("internal error")
	}

	// Die Main Datei wird abgerufen
	main_file, err := image_file.GetMainFile()
	if err != nil {
		return nil, fmt.Errorf("JSEngine>RunMainScript: " + err.Error())
	}

	// Es wird ermittelt ob es sich um eine Javascript Datei handelt
	if main_file.File.Playground != "js" && main_file.File.Playground != "javascript" {
		return nil, fmt.Errorf("JSEngine>RunMainScript: the image main file isn a javascript file")
	}

	// Es wird eine neie Engine erzeugt
	main_engine := goja.New()

	// Die Basis Funktionen werden registiert
	o.initRuntimeBaseFunctions(main_engine)

	// Die Main Engie wird zwischengespeichert
	o.jsInterpreters = append(o.jsInterpreters, main_engine)

	// Das Script wird eingelesen
	parsed_script := string(main_file.GetBytes())

	// Das Script wird ausgeführt
	result, err := main_engine.RunString(parsed_script)
	if err != nil {
		// Überprüfe, ob es sich um einen InterruptedError handelt
		if interruptedErr, ok := err.(*goja.InterruptedError); ok {
			// Es wird ermittelt ob sich um den Closer Command Signature String handelt, wenn ja wird die VM ohne Fehler geschlossen, andererseits wird die VM mit einem Feler abgebrochen
			if strings.HasPrefix(interruptedErr.Error(), o.getCloserValue()) {
				return nil, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Es wird gewartet bis die WaitGroup fertig ist
	o.wg.Wait()

	// Das Ergebniss wird zurückgegeben
	return result, nil
}
