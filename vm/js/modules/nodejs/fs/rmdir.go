package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

// fsCoreRmDir löscht ein Verzeichnis und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - orig_path: Der Pfad des zu löschenden Verzeichnisses.
//   - options: Optionale Parameter, die beim Löschen des Verzeichnisses verwendet werden können (als Schnittstelle definiert).
//
// Rückgabewerte:
//   - Ein Goja-Null-Wert, um den erfolgreichen Abschluss des Löschvorgangs darzustellen.
//   - Ein Fehler, falls beim Löschen des Verzeichnisses ein Problem auftritt, z. B. wenn das Verzeichnis nicht gefunden wird.
//
// Die Funktion ruft das Dateisystem über die Schnittstelle vmengine ab und prüft auf Fehler. Dann wird das zu löschende Verzeichnis
// anhand des Pfads ermittelt. Anschließend wird das Verzeichnis gelöscht, und bei einem erfolgreichen Vorgang wird ein Goja-Null-Wert
// als Ergebnis zurückgegeben. Im Falle eines Fehlers wird ein entsprechender Fehlerwert zurückgegeben.
func fsCoreRmDir(vmengine modules.VMInterface, jsrumtime *goja.Runtime, orig_path string, options interface{}) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal filesystem error, unkown error")
	}

	// Der Ordner wird ermittelt
	folder, err := fileSystem.GetDirByFullPath(orig_path)
	if err != nil {
		return nil, err
	}

	// Der Ordner wird gelöscht
	if err := folder.Delete(); err != nil {
		return nil, err
	}

	// Es wird ein Null Zurückgegeben
	return goja.Null(), nil
}

// Module_FS_SYNC_rmdirSync löscht ein Verzeichnis synchron und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen JavaScript-Wert.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu löschenden Verzeichnisses und optionale Parameter enthält.
//
// Rückgabewert:
//   - Ein Goja-Null-Wert, um den erfolgreichen Abschluss des Löschvorgangs darzustellen.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad und optionale Parameter aus den Argumenten.
// Anschließend wird die Kernfunktion fsCoreRmDir aufgerufen, um das Verzeichnis synchron zu löschen. Bei einem Fehler wird dieser als Panik ausgelöst,
// andernfalls gibt die Funktion ein Goja-Null-Wert zurück, um den erfolgreichen Abschluss des Löschvorgangs darzustellen.
func Module_FS_SYNC_rmdirSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "rmdir", 2)))
	}

	// Die Argumente werden extrahiert
	folderPath := parms.Arguments[0].String()
	options := parms.Arguments[1].Export()

	// Der Ordner wird entfernt
	_, err := fsCoreRmDir(vmengine, jsruntime, folderPath, options)
	if err != nil {
		panic(err)
	}

	// Das Ergebniss wird zurückgegeben
	return goja.Undefined()
}

// Module_FS_SYNC_rmdirCallback löscht ein Verzeichnis und ruft eine Callback-Funktion auf, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu löschenden Verzeichnisses, optionale Parameter und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Ein Goja-Null-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad, optionale Parameter und die Callback-Funktion aus den Argumenten.
// Anschließend wird die Kernfunktion fsCoreRmDir aufgerufen, um das Verzeichnis zu löschen. Je nach Ergebnis wird die Callback-Funktion mit den entsprechenden Parametern aufgerufen, um das Ergebnis zu verarbeiten.
// Schließlich gibt die Funktion ein Goja-Null-Wert zurück.
func Module_FS_SYNC_rmdirCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "rmdir", 3)))
	}

	// Die Argumente werden extrahiert
	folderPath := parms.Arguments[0].String()
	options := parms.Arguments[1].Export()
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	_, err := fsCoreRmDir(vmengine, jsruntime, folderPath, options)
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Undefined()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_rmdirPromises löscht ein Verzeichnis asynchron und gibt eine Promise zurück, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in eine JavaScript-Promise.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu löschenden Verzeichnisses und optionale Parameter enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Promise-Objekt, das das Ergebnis des Löschvorgangs darstellt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad und optionale Parameter aus den Argumenten.
// Anschließend erstellt sie eine neue Promise und die zugehörigen Auflösungs- und Ablehnungsfunktionen. Ein neuer Lesevorgang wird für die asynchrone
// Löschung des Verzeichnisses registriert. Der Löschvorgang wird im Hintergrund ausgeführt, und das Ergebnis wird ermittelt. Bei einem Fehler wird
// die Ablehnungsfunktion aufgerufen, um den Fehler zurückzugeben. Bei Erfolg wird die Auflösungsfunktion aufgerufen, um das Ergebnis zurückzugeben.
// Schließlich gibt die Funktion die Promise zurück, um das Ergebnis des Löschvorgangs asynchron zu verarbeiten.
func Module_FS_ASYNC_rmdirPromises(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "rmdir", 2)))
	}

	// Die Argumente werden extrahiert
	folderPath := parms.Arguments[0].String()
	options := parms.Arguments[1].Export()

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
		result, err := fsCoreRmDir(vmengine, jsrumtime, folderPath, options)
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
