package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

// fsCoreState prüft den Status eines Dateisystems und gibt ein Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - orig_path: Der Pfad oder die Identifikation des Dateisystems, dessen Status überprüft werden soll.
//
// Rückgabewerte:
//   - Ein Goja-Undefined-Wert, um den Status des Dateisystems anzuzeigen.
//   - Ein Fehler, falls beim Überprüfen des Dateisystemstatus ein Problem auftritt.
//
// Die Funktion nimmt die Schnittstelle vmengine und den Pfad oder die Identifikation des Dateisystems als Parameter entgegen. Sie überprüft
// den Status des Dateisystems und gibt ein Goja-Undefined-Wert zurück, um den Status anzuzeigen. Falls beim Überprüfen des Status ein Problem
// auftritt, wird ein entsprechender Fehlerwert zurückgegeben.
func fsCoreState(vmengine modules.VMInterface, jsrumtime *goja.Runtime, orig_path string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal filesystem error, unkown error")
	}

	// Es wird ermittelt ob es sich um eine Datei handelt
	result, err := fileSystem.GetStat(orig_path)
	if err != nil {
		return nil, err
	}

	// Das JS Objekt wird erstellt
	jsObject := jsrumtime.NewObject()

	// Es wird ermittelt ob es sich um einen FileStat handelt
	fileStateResult, isFileStat := result.(*fsysfile.FileStat)
	if isFileStat {
		// Der Typ wird definiert
		jsObject.Set("type", "file")

		// Die Restlichen Atribute werden abgerufen
		fstatAtrib, err := fileStateResult.GetAtributes()
		if err != nil {
			return nil, err
		}

		// Die Restlichen Atribute werden gesetzt
		for ai := range fstatAtrib {
			jsObject.Set(fstatAtrib[ai].Name, fstatAtrib[ai].Value)
		}

		// Das Objekt wird zurückgegeben
		return jsObject, nil
	}

	// Es wird ermittelt ob es sich um einen FolderStat handelt
	folderStatResult, isFolderStat := result.(*fsysfile.FolderStat)
	if isFolderStat {
		// Der Typ wird definiert
		jsObject.Set("type", "folder")

		// Die Restlichen Atribute werden abgerufen
		fstatAtrib, err := folderStatResult.GetAtributes()
		if err != nil {
			return nil, err
		}

		// Die Restlichen Atribute werden gesetzt
		for ai := range fstatAtrib {
			jsObject.Set(fstatAtrib[ai].Name, fstatAtrib[ai].Value)
		}

		// Das Objekt wird zurückgegeben
		return jsObject, nil
	}

	// Es wird ermittelt ob es sich um eine Datei handelt
	return nil, fmt.Errorf("invalid path: %s", orig_path)
}

// Module_FS_SYNC_statSync ruft synchron Metadaten über eine Datei oder ein Verzeichnis ab und gibt das Ergebnis zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - parms: Ein Goja-Funktionsaufruf, der den vollständigen Pfad zur Datei oder zum Verzeichnis enthält, über das Metadaten abgerufen werden sollen.
//
// Rückgabewert:
//   - Ein JavaScript-Objekt, das Metadaten über die Datei oder das Verzeichnis enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den vollständigen Pfad zur Datei oder zum Verzeichnis aus den Argumenten.
// Anschließend ruft sie die Kernfunktion fsCoreState auf, um Metadaten über die Datei oder das Verzeichnis synchron abzurufen. Bei einem Fehler wird
// der Fehler als JavaScript-Wert zurückgegeben, ansonsten wird ein JavaScript-Objekt mit den Metadaten zurückgegeben.
func Module_FS_SYNC_statSync(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "stat", 1)))
	}

	// Die Argumente werden extrahiert
	fqPath := parms.Arguments[0].String()

	// Die Metadaten über die Datei bzw den Ordner werden ermittelt
	result, err := fsCoreState(vmengine, jsrumtime, fqPath)
	if err != nil {
		return jsrumtime.ToValue(err)
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_statCallback ruft Metadaten über eine Datei oder ein Verzeichnis ab und ruft eine Callback-Funktion auf, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - parms: Ein Goja-Funktionsaufruf, der den vollständigen Pfad zur Datei oder zum Verzeichnis enthält, über das Metadaten abgerufen werden sollen, sowie eine Callback-Funktion.
//
// Rückgabewert:
//   - Ein Goja-Undefined-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den vollständigen Pfad zur Datei oder zum Verzeichnis aus den Argumenten.
// Sie überprüft auch, ob eine gültige Callback-Funktion übergeben wurde. Anschließend ruft sie die Kernfunktion fsCoreState auf, um Metadaten über die Datei oder das Verzeichnis abzurufen.
// Je nach Ergebnis wird die Callback-Funktion mit den entsprechenden Parametern aufgerufen, um das Ergebnis zu verarbeiten.
// Schließlich gibt die Funktion ein Goja-Undefined-Wert zurück.
func Module_FS_SYNC_statCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "stat", 2)))
	}

	// Die Argumente werden extrahiert
	fqPath := parms.Arguments[0].String()
	callbackFunction, isOk := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	if result, err := fsCoreState(vmengine, jsruntime, fqPath); err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(result)}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_statPromises ruft asynchron Metadaten über eine Datei oder ein Verzeichnis ab und gibt eine Promise zurück, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsrumtime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in eine JavaScript-Promise.
//   - parms: Ein Goja-Funktionsaufruf, der den vollständigen Pfad zur Datei oder zum Verzeichnis enthält, über das Metadaten abgerufen werden sollen.
//
// Rückgabewert:
//   - Ein JavaScript-Promise-Objekt, das das Ergebnis des Metadatenabrufs darstellt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den vollständigen Pfad zur Datei oder zum Verzeichnis aus den Argumenten.
// Anschließend erstellt sie eine neue Promise und die zugehörigen Auflösungs- und Ablehnungsfunktionen. Ein neuer Lesevorgang wird für die asynchrone
// Abfrage der Metadaten registriert. Die Abfrage wird im Hintergrund ausgeführt, und das Ergebnis wird ermittelt. Bei einem Fehler wird die
// Ablehnungsfunktion aufgerufen, um den Fehler zurückzugeben. Bei Erfolg wird die Auflösungsfunktion aufgerufen, um das Ergebnis zurückzugeben.
// Schließlich gibt die Funktion die Promise zurück, um das Ergebnis des Metadatenabrufs asynchron zu verarbeiten.
func Module_FS_ASYNC_statPromises(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "stat", 1)))
	}

	// Die Argumente werden extrahiert
	fqPath := parms.Arguments[0].String()

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
		result, err := fsCoreState(vmengine, jsrumtime, fqPath)
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
