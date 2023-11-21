package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsmountimgCore(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, sourcePath string) (goja.Value, error) {
	return nil, nil
}

// Module_FS_SYNC_mountimgSync mountimg mountet ein Image synchron und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu mountenden Image enthält.
//
// Rückgabewert:
//   - Das Ergebnis des Mountvorgangs oder ein Fehler als Goja-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zum zu mountenden Image.
// Anschließend versucht sie, das Image zu mounten, und gibt das Ergebnis oder einen Fehler als Goja-Wert zurück.
func Module_FS_SYNC_mountimgSync(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "unlink", 1)))
	}

	// Die Argumente werden extrahiert
	fqSourcePath := parms.Arguments[0].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsmountimgCore(vmengine, jsrumtime, fqSourcePath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_mountimgCallback mountimg mountet ein Image synchron mit einer Callback-Funktion.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu mountenden Image und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - undefined
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zum zu mountenden Image und die Callback-Funktion.
// Anschließend versucht sie, das Image zu mounten, und ruft die Callback-Funktion mit dem Ergebnis oder einem Fehler auf.
// Schließlich gibt sie `undefined` zurück.
func Module_FS_SYNC_mountimgCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
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
	if result, err := fsmountimgCore(vmengine, jsruntime, sourcePath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_mountimgPromises mountimg mountet ein Image asynchron mit Promises und gibt ein Promise-Objekt zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zum zu mountenden Image enthält.
//
// Rückgabewert:
//   - Ein Promise-Objekt, das das Ergebnis oder einen Fehler enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Pfad zum zu mountenden Image.
// Anschließend erstellt sie ein neues Promise-Objekt und fügt einen asynchronen Vorgang hinzu, um das Image zu mounten.
// Das Promise-Objekt enthält das Ergebnis oder einen Fehler, der durch die Promise-Resolving-Funktionen zurückgegeben wird.
func Module_FS_ASYNC_mountimgPromises(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
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
		result, err := fsmountimgCore(vmengine, jsrumtime, sourcePath)
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
