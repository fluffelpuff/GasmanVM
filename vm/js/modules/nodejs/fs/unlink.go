package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsUnlinkCore(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, sourcePath string) (goja.Value, error) {
	return nil, nil
}

// Module_FS_SYNC_unlinkSync löscht eine Datei synchron und gibt das Ergebnis oder einen Fehler zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu löschenden Datei enthält.
//
// Rückgabewert:
//   - Das Ergebnis der Löschoperation oder ein Fehler als Goja-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu löschenden Datei.
// Anschließend versucht sie, die Datei synchron zu löschen und gibt das Ergebnis oder einen Fehler zurück.
func Module_FS_SYNC_unlinkSync(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "unlink", 1)))
	}

	// Die Argumente werden extrahiert
	fqSourcePath := parms.Arguments[0].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsUnlinkCore(vmengine, jsrumtime, fqSourcePath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_unlinkCallback löscht eine Datei asynchron und gibt das Ergebnis oder einen Fehler über eine Callback-Funktion zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses für das Callback.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu löschenden Datei und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Ein Undefined-Wert, da das Ergebnis über die Callback-Funktion übertragen wird.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu löschenden Datei und die Callback-Funktion.
// Anschließend versucht sie, die Datei asynchron zu löschen, ruft die Callback-Funktion auf und gibt ein Undefined zurück.
func Module_FS_SYNC_unlinkCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
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
	if result, err := fsUnlinkCore(vmengine, jsruntime, sourcePath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_unlinkPromises löscht eine Datei asynchron und gibt das Ergebnis oder einen Fehler über eine Promise zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses für die Promise.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu löschenden Datei enthält.
//
// Rückgabewert:
//   - Eine Promise, die das Ergebnis der Dateilöschung oder einen Fehler enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu löschenden Datei.
// Anschließend erstellt sie eine neue Promise, registriert einen asynchronen Vorgang, um die Datei zu löschen, und gibt die Promise zurück,
// die das Ergebnis oder einen Fehler nach Abschluss des Löschvorgangs enthält.
func Module_FS_ASYNC_unlinkPromises(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
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
		result, err := fsUnlinkCore(vmengine, jsrumtime, sourcePath)
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
