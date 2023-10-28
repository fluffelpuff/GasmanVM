package fs

import (
	"fmt"
	"reflect"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func validateOptions(options map[string]interface{}) bool {
	recursive, mode := 0, 0
	for key := range options {
		switch key {
		case "mode":
			mode++
		case "recursive":
			recursive++
		default:
			return false
		}
	}
	if recursive > 1 {
		return false
	}
	if mode > 1 {
		return false
	}
	return true
}

// fsMkdirCore erstellt ein Verzeichnis im Dateisystem.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - folderPath: Der Pfad zum zu erstellenden Verzeichnis.
//
// Rückgabewert:
//   - Ein Goja-Wert, der das Ergebnis des Verzeichniserstellungsprozesses enthält.
//   - Ein Fehler, falls ein Problem bei der Erstellung des Verzeichnisses auftritt.
//
// Die Funktion ruft das Dateisystem von der virtuellen Maschine ab und verwendet es, um das Verzeichnis zu erstellen.
// Sie gibt ein Goja-Undefined zurück, um den Erfolg anzuzeigen, oder einen Fehler, falls die Erstellung fehlschlägt.
func fsMkdirCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, folderPath string, options map[string]interface{}) error {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return fmt.Errorf("internal error, not fs")
	}

	// Es wird geprüft ob die mkdir Option zulässig ist
	if !validateOptions(options) {
		return fmt.Errorf("invalid operation options passed")
	}

	// Die Werte werden abgerufen
	recursiveRetrived, foundRecursive := options["recursive"]
	modeRetrived, foundMode := options["mode"]

	// Speichert die Finalen Werte ab
	var recursive bool
	var mode int64

	// Es wird ermittelt ob der Vorgang Rekrusiv durchgeführt werden soll
	if foundRecursive {
		conve, ok := recursiveRetrived.(bool)
		if !ok {
			return fmt.Errorf("invalid recursive data type")
		}
		recursive = conve
	} else {
		recursive = false
	}

	// Es wird ermittelt welcher Modus verwendet werden soll
	if foundMode {
		conve, ok := modeRetrived.(int64)
		if !ok {
			return fmt.Errorf("invalid recursive data type")
		}
		mode = conve
	} else {
		mode = 0
	}

	// Der Ordner wird erstellt
	err := fileSystem.CreateNewFolder(folderPath, mode, recursive)
	if err != nil {
		return err
	}

	// Die Daten werden zurückgegeben
	return nil
}

// Module_FS_SYNC_mkdirSync erstellt ein Verzeichnis synchron und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu erstellenden Verzeichnis enthält.
//
// Rückgabewert:
//   - Ein Wert, der das Ergebnis des Verzeichniserstellungsprozesses enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist, extrahiert den Verzeichnispfad und erstellt das Verzeichnis synchron. Das Ergebnis wird zurückgegeben.
func Module_FS_SYNC_mkdirSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "mkdirSync", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um ein Objekt handelt
	if parms.Arguments[1].ExportType() != reflect.TypeOf(make(map[string]interface{})) {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a config object")))
	}

	// Die Optionen werden Extrahiert
	options, isOk := parms.Arguments[1].Export().(map[string]interface{})
	if !isOk {
		panic(goja.New().NewGoError(fmt.Errorf("invalid config")))
	}

	// Es wird versucht die Datei einzulesen
	err := fsMkdirCore(vmengine, jsruntime, parms.Arguments[0].String(), options)
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

// Module_FS_SYNC_mkdirCallback erstellt ein Verzeichnis synchron und ruft eine Callback-Funktion auf.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu erstellenden Verzeichnis und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Null, da die Funktion die Daten über eine Callback-Funktion zurückgibt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad und die Callback-Funktion.
// Anschließend wird das Verzeichnis synchron erstellt, und die Callback-Funktion wird mit den Ergebnissen aufgerufen.
func Module_FS_SYNC_mkdirCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "mkdir", 3)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um ein Objekt handelt
	if parms.Arguments[1].ExportType() != reflect.TypeOf(make(map[string]interface{})) {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a config object")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Optionen werden Extrahiert
	options, isOk := parms.Arguments[1].Export().(map[string]interface{})
	if !isOk {
		panic(goja.New().NewGoError(fmt.Errorf("invalid config")))
	}

	// Die Callback Funktion wird zurückgegeben
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	err := fsMkdirCore(vmengine, jsruntime, parms.Arguments[0].String(), options)
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

// Module_FS_ASYNC_mkdirPromises erstellt ein Verzeichnis asynchron und gibt eine Promise zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu erstellenden Verzeichnis enthält.
//
// Rückgabewert:
//   - Eine Promise, die das Ergebnis des asynchronen Erstellens des Verzeichnisses enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad.
// Anschließend wird das Verzeichnis asynchron erstellt, und das Ergebnis oder ein Fehler wird in einer Promise zurückgegeben.
func Module_FS_ASYNC_mkdirPromises(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "mkdir", 1)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um ein Objekt handelt
	if parms.Arguments[1].ExportType() != reflect.TypeOf(make(map[string]interface{})) {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a config object")))
	}

	// Die Optionen werden Extrahiert
	options, isOk := parms.Arguments[1].Export().(map[string]interface{})
	if !isOk {
		panic(goja.New().NewGoError(fmt.Errorf("invalid config")))
	}

	// Erstellen Sie eine neue Promise und die Resolving-Funktionen
	promise, resolve, reject := jsruntime.NewPromise()

	// Es wird ein neuer Lesender Vorgang Registriert
	vmengine.AddNewRoutine()

	// Die Datei wird Asynchrone eingelesen
	go func() {
		// Wird ausgeführt wenn die Funktion fertig ist
		defer func() {
			// Es wird Signalisiert dass die Unit beendet wurde
			vmengine.RemoveRoutine()
		}()

		// Die Datei wird eingelesen
		err := fsMkdirCore(vmengine, jsruntime, parms.Arguments[0].String(), options)
		if err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(goja.Undefined())
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
