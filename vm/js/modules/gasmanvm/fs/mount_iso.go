package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func fsmountisoCore(vmengine modules.VMInterface, jsrumtime *goja.Runtime, sourcePath string) (goja.Value, error) {
	return nil, nil
}

// Module_FS_SYNC_mountisoSync mountisoSync mountet eine ISO-Datei synchron und gibt das Ergebnis sofort zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu mountenden ISO-Datei enthält.
//
// Rückgabewert:
//   - Das Ergebnis der Operation wird direkt zurückgegeben, oder ein Fehler, falls die Operation fehlschlägt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu mountenden ISO-Datei.
// Anschließend wird die ISO-Datei synchron gemountet und das Ergebnis oder ein Fehler wird sofort zurückgegeben.
func Module_FS_SYNC_mountisoSync(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "unlink", 1)))
	}

	// Die Argumente werden extrahiert
	fqSourcePath := parms.Arguments[0].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsmountisoCore(vmengine, jsrumtime, fqSourcePath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_mountisoCallback mountisoCallback mountet eine ISO-Datei synchron und gibt das Ergebnis über eine Callback-Funktion zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu mountenden ISO-Datei und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Ein Wert wird nicht direkt zurückgegeben. Stattdessen wird das Ergebnis oder ein Fehler über die Callback-Funktion übergeben.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu mountenden ISO-Datei und die Callback-Funktion.
// Anschließend mountet sie die ISO-Datei synchron und gibt das Ergebnis oder einen Fehler über die Callback-Funktion zurück.
func Module_FS_SYNC_mountisoCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "unlink", 2)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()
	callbackFunction, isOk := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	if result, err := fsmountisoCore(vmengine, jsruntime, sourcePath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_mountisoPromises mountiso mountet eine ISO-Datei asynchron und gibt das Ergebnis als Promise zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu mountenden ISO-Datei enthält.
//
// Rückgabewert:
//   - Ein Promise, das das Ergebnis oder einen Fehler als Goja-Wert zurückgeben wird.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu mountenden ISO-Datei.
// Anschließend mountet sie die ISO-Datei asynchron und gibt das Ergebnis als Promise zurück.
func Module_FS_ASYNC_mountisoPromises(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "unlink", 1)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()

	// Erstellen Sie eine neue Promise und die Resolving-Funktionen
	promise, resolve, reject := jsrumtime.NewPromise()

	// Es wird ein neuer Lesender Vorgang Registriert
	vmengine.AddNewRoutine()

	// Die Datei wird Asynchrone eingelesen
	go func() {
		// Wird ausgeführt wenn die Funktion fertig ist
		defer func() {
			// Es wird Signalisiert dass die Unit beendet wurde
			vmengine.RemoveRoutine()
		}()

		// Die Datei wird gelöscht
		result, err := fsmountisoCore(vmengine, jsrumtime, sourcePath)
		if err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(result)
	}()

	// Rückgabe der Promise
	return jsrumtime.ToValue(promise)
}
