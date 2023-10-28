package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func fsSymlinkCore(vmengine modules.VMInterface, jsrumtime *goja.Runtime, sourcePath string, destinationPath string) (goja.Value, error) {
	return nil, nil
}

// Module_FS_SYNC_symlinkSync erstellt einen symbolischen Link synchron zwischen den angegebenen Quell- und Ziel-Pfaden und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur Quelle und den Pfad zum Ziel für den symbolischen Link enthält.
//
// Rückgabewert:
//   - Das Ergebnis der symbolischen Link-Erstellung oder ein Fehler, falls aufgetreten.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert die Quell- und Ziel-Pfade aus den Argumenten.
// Anschließend erstellt sie synchron einen symbolischen Link zwischen den Pfaden und gibt das Ergebnis oder einen Fehler zurück, falls aufgetreten.
func Module_FS_SYNC_symlinkSync(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "symlinkSync", 2)))
	}

	// Die Argumente werden extrahiert
	fqSourcePath := parms.Arguments[0].String()
	fqDestinationPath := parms.Arguments[1].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsSymlinkCore(vmengine, jsrumtime, fqSourcePath, fqDestinationPath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_symlinkCallback erstellt einen symbolischen Link zwischen den angegebenen Quell- und Ziel-Pfaden und ruft dann eine Callback-Funktion auf.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses für die Callback-Funktion.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur Quelle, den Pfad zum Ziel und die Callback-Funktion für die symbolische Link-Erstellung enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Undefined-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert die Quell- und Ziel-Pfade sowie die Callback-Funktion aus den Argumenten.
// Anschließend versucht sie, den symbolischen Link zwischen den Pfaden zu erstellen, und ruft dann die Callback-Funktion auf. Wenn ein Fehler auftritt,
// wird der Fehler als Argument an die Callback-Funktion übergeben. Bei erfolgreichem Abschluss wird `undefined` als Argument übergeben.
// Schließlich gibt die Funktion `undefined` zurück, um den Abschluss der Operation zu signalisieren.
func Module_FS_SYNC_symlinkCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "symlink", 3)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()
	destinationPath := parms.Arguments[1].String()
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	if result, err := fsSymlinkCore(vmengine, jsruntime, sourcePath, destinationPath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_symlinkPromises erstellt einen symbolischen Link asynchron zwischen den angegebenen Quell- und Ziel-Pfaden und gibt eine Promise zurück, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in eine JavaScript-Promise.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad zur Quelle und den Pfad zum Ziel für den symbolischen Link enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Promise-Objekt, das das Ergebnis der Erstellung des symbolischen Links darstellt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert die Quell- und Ziel-Pfade aus den Argumenten.
// Anschließend erstellt sie eine neue Promise und die zugehörigen Auflösungs- und Ablehnungsfunktionen. Ein neuer Vorgang wird registriert,
// um den symbolischen Link asynchron zu erstellen. Die Erstellung wird im Hintergrund ausgeführt, und das Ergebnis wird ermittelt.
// Bei einem Fehler wird die Ablehnungsfunktion aufgerufen, um den Fehler zurückzugeben. Bei Erfolg wird die Auflösungsfunktion aufgerufen,
// um das Ergebnis zurückzugeben. Schließlich gibt die Funktion die Promise zurück, um das Ergebnis der symbolischen Link-Erstellung asynchron zu verarbeiten.
func Module_FS_ASYNC_symlinkPromises(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "symlink", 2)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()
	destinationPath := parms.Arguments[1].String()

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
		result, err := fsSymlinkCore(vmengine, jsrumtime, sourcePath, destinationPath)
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
