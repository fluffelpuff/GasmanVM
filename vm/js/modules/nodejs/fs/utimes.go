package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsUtimesCore(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, sourcePath string) (goja.Value, error) {
	return nil, nil
}

// Module_FS_SYNC_utimesSync aktualisiert die Zugriffs- und Modifikationszeiten einer Datei oder eines Ordners synchron und gibt das Ergebnis oder einen Fehler zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu aktualisierenden Datei oder zum Ordner enthält.
//
// Rückgabewert:
//   - Das Ergebnis der Aktualisierung oder ein Fehler als Goja-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu aktualisierenden Datei oder zum Ordner.
// Anschließend versucht sie, die Zugriffs- und Modifikationszeiten synchron zu aktualisieren und gibt das Ergebnis oder einen Fehler zurück.
func Module_FS_SYNC_utimesSync(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "utimesSync", 1)))
	}

	// Die Argumente werden extrahiert
	fqSourcePath := parms.Arguments[0].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsUtimesCore(vmengine, jsrumtime, fqSourcePath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_utimesCallback aktualisiert die Zugriffs- und Modifikationszeiten einer Datei oder eines Ordners asynchron mit einer Callback-Funktion und gibt Undefined zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu aktualisierenden Datei oder zum Ordner sowie eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Undefined.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu aktualisierenden Datei oder zum Ordner sowie die Callback-Funktion.
// Anschließend versucht sie, die Zugriffs- und Modifikationszeiten asynchron zu aktualisieren und ruft die Callback-Funktion mit dem Ergebnis oder einem Fehler auf.
func Module_FS_SYNC_utimesCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "utimes", 2)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()
	callbackFunction, isOk := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	if result, err := fsUtimesCore(vmengine, jsruntime, sourcePath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_utimesPromises aktualisiert die Zugriffs- und Modifikationszeiten einer Datei oder eines Ordners asynchron mit Promises und gibt ein Promise-Objekt zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur zu aktualisierenden Datei oder zum Ordner enthält.
//
// Rückgabewert:
//   - Ein Promise-Objekt, das das Ergebnis oder einen Fehler enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zur zu aktualisierenden Datei oder zum Ordner.
// Anschließend erstellt sie ein neues Promise-Objekt und fügt einen asynchronen Vorgang hinzu, um die Zugriffs- und Modifikationszeiten zu aktualisieren.
// Das Promise-Objekt enthält das Ergebnis oder einen Fehler, der durch die Promise-Resolving-Funktionen zurückgegeben wird.
func Module_FS_ASYNC_utimesPromises(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "utimes", 1)))
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
		result, err := fsUtimesCore(vmengine, jsrumtime, sourcePath)
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
