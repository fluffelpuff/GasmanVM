package fs

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

// fsCoreRename benennt eine Datei um und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - orig_path: Der ursprüngliche Pfad der Datei, die umbenannt werden soll.
//   - new_name: Der neue Name, den die Datei erhalten soll.
//
// Rückgabewerte:
//   - Ein Goja-Null-Wert, um den erfolgreichen Abschluss der Umbenennung darzustellen.
//   - Ein Fehler, falls beim Umbenennen der Datei ein Problem auftritt, z. B. wenn die Datei nicht gefunden wird.
//
// Die Funktion ruft das Dateisystem über die Schnittstelle vmengine ab und prüft auf Fehler. Dann wird überprüft, ob die Datei anhand des ursprünglichen Pfads verfügbar ist.
// Wenn die Datei gefunden wird, wird sie umbenannt. Ein erfolgreicher Vorgang wird durch einen Goja-Null-Wert signalisiert. Im Falle eines Fehlers wird ein entsprechender Fehlerwert zurückgegeben.
func fsCoreRename(vmengine modules.VMInterface, jsrumtime *goja.Runtime, orig_path string, new_name string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal filesystem error, unkown error")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	fileResolve, err := fileSystem.GetFileByFullPath(orig_path)
	if err != nil {
		return nil, err
	}
	if fileResolve == nil {
		return nil, fmt.Errorf("file not found")
	}

	// Die Datei wird umbenannt
	if err := fileResolve.Rename(new_name); err != nil {
		return nil, err
	}

	// Der Vorgang wurde erfolgreich durchgeführt
	return goja.Null(), nil
}

// Module_FS_SYNC_renameSync benennt eine Datei synchron um.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen JavaScript-Wert.
//   - parms: Ein Goja-Funktionsaufruf, der die Argumente für diese Funktion enthält.
//
// Rückgabewert:
//   - Ein Goja-Undefined-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist. Sie extrahiert den Quellpfad der Datei und den neuen Namen
// aus den übergebenen Argumenten. Die Datei wird synchron umbenannt, und bei einem Fehler wird dieser als Panik ausgelöst. Andernfalls gibt
// die Funktion ein Goja-Undefined-Wert zurück
func Module_FS_SYNC_renameSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "renameSync", 2)))
	}

	// Die Argumente werden extrahiert
	sourceFullPathName := parms.Arguments[0].String()
	newName := strings.ToLower(strings.TrimSpace(parms.Arguments[1].String()))

	// Die Datei wird umbeannt
	if _, err := fsCoreRename(vmengine, jsruntime, sourceFullPathName, newName); err != nil {
		panic(err)
	}

	// Es wird ein Undefined zurückgegeben
	return goja.Null()
}

// Module_FS_SYNC_renameCallback benennt eine Datei um und ruft eine Callback-Funktion auf, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen JavaScript-Wert.
//   - parms: Ein Goja-Funktionsaufruf, der die Argumente für diese Funktion enthält.
//
// Rückgabewert:
//   - Ein Goja-Undefined-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist. Sie extrahiert den Quellpfad der Datei und den neuen Namen
// aus den übergebenen Argumenten. Außerdem wird eine Callback-Funktion extrahiert. Die Datei wird umbenannt, und das Ergebnis wird überprüft.
// Wenn ein Fehler auftritt, wird die Callback-Funktion mit dem Fehlerwert aufgerufen. Andernfalls wird die Callback-Funktion ohne Fehler
// aufgerufen. Schließlich gibt die Funktion ein Goja-Undefined-Wert zurück.
func Module_FS_SYNC_renameCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readFile", 3)))
	}

	// Die Argumente werden extrahiert
	sourceFullPathName := parms.Arguments[0].String()
	newName := strings.ToLower(strings.TrimSpace(parms.Arguments[1].String()))
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Die Datei wird umbenannt
	var callbackParms goja.FunctionCall
	_, err := fsCoreRename(vmengine, jsruntime, sourceFullPathName, newName)
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Undefined()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_renamePromises benennt eine Datei asynchron um und gibt eine Promise zurück, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in ein JavaScript-Promise.
//   - parms: Ein Goja-Funktionsaufruf, der die Argumente für diese Funktion enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Promise-Objekt, das das Ergebnis der Umbenennung der Datei darstellt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist. Sie extrahiert den Quellpfad der Datei und den neuen Namen
// aus den übergebenen Argumenten. Dann erstellt sie eine neue Promise und die zugehörigen Auflösungs- und Ablehnungsfunktionen. Ein neuer
// Lesevorgang wird für die asynchrone Umbenennung registriert. Die Datei wird im Hintergrund umbenannt, und das Ergebnis wird ermittelt.
// Bei einem Fehler wird die Ablehnungsfunktion aufgerufen und ein Fehlerwert zurückgegeben. Bei Erfolg wird die Auflösungsfunktion aufgerufen,
// um das Ergebnis zurückzugeben. Schließlich gibt die Funktion die Promise zurück, um das Ergebnis der Umbenennung asynchron zu verarbeiten.
func Module_FS_ASYNC_renamePromises(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "rename", 2)))
	}

	// Die Argumente werden extrahiert
	sourceFullPathName := parms.Arguments[0].String()
	newName := strings.ToLower(strings.TrimSpace(parms.Arguments[1].String()))

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
		result, err := fsCoreRename(vmengine, jsruntime, sourceFullPathName, newName)
		if err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Die Daten werden zurückgegeben
		resolve(result)
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
