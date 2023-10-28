package fs

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func isValidOpenFlag(flag string) bool {
	// Gültige Flags für fs.open in Node.js
	validFlags := regexp.MustCompile(`^[rwa]+\+?$`)
	return validFlags.MatchString(flag)
}

// fsOpenCore öffnet eine Datei im Dateisystem und gibt ihren Inhalt basierend auf den angegebenen Optionen zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Dateiinhalts in einen Goja-Wert.
//   - filePath: Der Pfad zur zu öffnenden Datei.
//   - fileOption: Die Option, wie der Dateiinhalt zurückgegeben werden soll (z. B. "utf8", "binary", "base64").
//
// Rückgabewert:
//   - Ein Goja-Wert, der den Dateiinhalt basierend auf der Option enthält.
//   - Ein Fehler, falls ein Problem beim Öffnen der Datei oder bei der Verarbeitung der Option auftritt.
//
// Die Funktion ruft das Dateisystem von der virtuellen Maschine ab, öffnet die Datei und gibt ihren Inhalt basierend auf
// der angegebenen Option zurück. Sie gibt den Dateiinhalt als Goja-Wert zurück und meldet einen Fehler, wenn ein Problem auftritt.
func fsOpenCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, filePath string, flags string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal fs error")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	resolvedFile, err := fileSystem.GetFileByFullPath(filePath)
	if err != nil {
		return nil, err
	}

	// Es wird ermittelt ob es sich um einen zulässigen Flag handelt
	if !isValidOpenFlag(strings.ToLower(flags)) {
		return nil, fmt.Errorf("invalid operation flag")
	}

	// Es wird ein File Descripto erstellt
	fileDescriptor, err := resolvedFile.GetFileDescriptor(strings.ToLower(flags))
	if err != nil {
		return nil, err
	}

	// Der Datei Discriptor wird zurückgegeben
	return jsruntime.ToValue(fileDescriptor), nil
}

// Module_FS_SYNC_openSync öffnet eine Datei synchron und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu öffnenden Datei und die Dateioptionen enthält.
//
// Rückgabewert:
//   - Das Ergebnis der synchronen Dateiöffnung.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Dateipfad und die Dateioptionen.
// Anschließend wird die Datei synchron geöffnet, und das Ergebnis oder ein Fehler wird zurückgegeben.
func Module_FS_SYNC_openSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "openSync", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird versucht die Datei einzulesen
	result, err := fsOpenCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String())
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_openCallback öffnet eine Datei synchron und ruft eine Callback-Funktion auf.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu öffnenden Datei, die Dateioptionen und die Callback-Funktion enthält.
//
// Rückgabewert:
//   - Null, da die Funktion synchron ist und das Ergebnis über die Callback-Funktion übergeben wird.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Dateipfad, die Dateioptionen und die Callback-Funktion.
// Anschließend wird die Datei synchron geöffnet, und das Ergebnis oder ein Fehler wird an die Callback-Funktion übergeben.
func Module_FS_SYNC_openCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "open", 3)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Callback Funktion wird eingelesen
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	result, err := fsOpenCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String())

	// Die Callback Parameter werden erzeugt
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err.Error()), goja.Undefined()}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null(), result}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

// Module_FS_ASYNC_openPromises öffnet eine Datei asynchron und gibt eine Promise zurück, die das Ergebnis enthält.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu öffnenden Datei und die Dateioptionen enthält.
//
// Rückgabewert:
//   - Eine Promise, die das Ergebnis der Operation oder einen Fehler enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Dateipfad und die Dateioptionen.
// Anschließend wird die Datei asynchron geöffnet, und das Ergebnis oder ein Fehler wird in einer Promise zurückgegeben.
func Module_FS_ASYNC_openPromises(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "open", 2)))
	}

	// Die Argumente werden extrahiert
	filePath := parms.Arguments[0].String()
	fileOption := strings.ToLower(strings.TrimSpace(parms.Arguments[1].String()))

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
		result, err := fsOpenCore(vmengine, jsruntime, filePath, fileOption)
		if err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(result)
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
