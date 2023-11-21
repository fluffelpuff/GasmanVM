package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

// fsCoreReadDir listet den Inhalt eines Verzeichnisses auf und gibt die Ergebnisse zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in JavaScript-Werte.
//   - dirPath: Der Pfad des Verzeichnisses, dessen Inhalt aufgelistet werden soll.
//
// Rückgabewerte:
//   - Ein JavaScript-Array, das die aufgelisteten Verzeichnis- und Dateiobjekte enthält.
//   - Ein Fehler, falls beim Auflisten des Verzeichnisinhalts ein Problem auftritt, z. B. wenn das Verzeichnis nicht gefunden wird.
//
// Die Funktion ruft das Dateisystem über die Schnittstelle vmengine ab und prüft auf Fehler. Dann wird der Inhalt des Verzeichnisses
// aufgelistet und die einzelnen Einträge verarbeitet. Jeder Eintrag wird in ein JavaScript-Objekt umgewandelt und in einem Array
// zwischengespeichert. Das Array wird dann als Ergebnis zurückgegeben. Bei einem Fehler wird ein entsprechender Fehlerwert zurückgegeben.
func fsCoreReadDir(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, dirPath string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal filesystem error, unkown error")
	}

	// Der Inhalt des Verzeichniss wird aufgelistet
	entries, err := fileSystem.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// Die Einzelnen Einträge werden abgeabreitet
	var result []interface{}
	for i := range entries {
		// Es wird ermittelt ob es sich um einen Ordner handelt
		folderItem, isFolder := entries[i].(fsysfile.Folder)
		if isFolder {
			// Es wird ein Objekt erstellt
			newFolderItem := jsruntime.NewObject()
			newFolderItem.Set("type", "folder")
			newFolderItem.Set("name", folderItem.GetName())

			// Der Eintrag wird zwischengespeichert
			result = append(result, newFolderItem)

			// Die Nächste Runde
			continue
		}

		// Es wird ermittelt ob es sich um eine Datei handelt
		fileItem, isFile := entries[i].(fsysfile.File)
		if isFile {
			// Es wird ein Objekt erstellt
			newFolderItem := jsruntime.NewObject()
			newFolderItem.Set("type", "file")
			newFolderItem.Set("name", fileItem.GetName())

			// Der Eintrag wird zwischengespeichert
			result = append(result, newFolderItem)

			// Die Nächste Runde
			continue
		}

		// Es wird ein Fehler ausgelöst
		return nil, fmt.Errorf("unkown entry")
	}

	// Es wird ein Array erstellt
	newResultArray := jsruntime.NewArray(result...)

	// Es wird ein Undefined zurückgegeben
	return newResultArray, nil
}

// Module_FS_SYNC_readdirSync listet den Inhalt eines Verzeichnisses synchron auf und gibt die Ergebnisse zurück.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen JavaScript-Wert.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu durchsuchenden Verzeichnisses enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Array, das die aufgelisteten Verzeichnis- und Dateiobjekte enthält.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad aus den Argumenten.
// Anschließend wird die Kernfunktion fsCoreReadDir aufgerufen, um den Verzeichnisinhalt aufzulisten. Wenn ein Fehler auftritt, wird dieser als Panik ausgelöst.
// Andernfalls gibt die Funktion den aufgelisteten Inhalt als JavaScript-Array zurück.
func Module_FS_SYNC_readdirSync(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readdir", 1)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Der Ordnerinahlt wird eingelesen
	result, err := fsCoreReadDir(vmengine, jsruntime, parms.Arguments[0].String())
	if err != nil {
		panic(err)
	}

	// Es wird ein Undefined zurückgegeben
	return result
}

// Module_FS_SYNC_readdirCallback listet den Inhalt eines Verzeichnisses auf und ruft eine Callback-Funktion auf, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen JavaScript-Wert.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu durchsuchenden Verzeichnisses und eine Callback-Funktion enthält.
//
// Rückgabewert:
//   - Ein Goja-Undefined-Wert.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad und die Callback-Funktion aus den Argumenten.
// Anschließend wird die Kernfunktion fsCoreReadDir aufgerufen, um den Verzeichnisinhalt aufzulisten. Je nach Ergebnis wird die Callback-Funktion mit den
// entsprechenden Parametern aufgerufen, um das Ergebnis zu verarbeiten. Schließlich gibt die Funktion ein Goja-Undefined-Wert zurück.
func Module_FS_SYNC_readdirCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readdir", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[1].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Callback Funktion wird eingelesen
	callbackFunction := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)

	// Der Ordnerinahlt wird eingelesen
	result, err := fsCoreReadDir(vmengine, jsruntime, parms.Arguments[0].String())
	var functionParms goja.FunctionCall
	if err != nil {
		functionParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		functionParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null(), result}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(functionParms)

	// Es wird ein Undefined zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_readdirPromises listet den Inhalt eines Verzeichnisses asynchron auf und gibt eine Promise zurück, um das Ergebnis zu verarbeiten.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in eine JavaScript-Promise.
//   - parms: Ein Goja-Funktionsaufruf, der den Pfad des zu durchsuchenden Verzeichnisses enthält.
//
// Rückgabewert:
//   - Ein JavaScript-Promise-Objekt, das das Ergebnis der Verzeichnisauflistung darstellt.
//
// Die Funktion überprüft, ob die erforderliche Anzahl von Parametern vorhanden ist und extrahiert den Verzeichnispfad aus den Argumenten.
// Anschließend erstellt sie eine neue Promise und die zugehörigen Auflösungs- und Ablehnungsfunktionen. Ein neuer Lesevorgang wird für
// die asynchrone Verzeichnisauflistung registriert. Der Verzeichnisinhalt wird im Hintergrund aufgelistet, und das Ergebnis wird ermittelt.
// Bei einem Fehler wird die Ablehnungsfunktion aufgerufen und gibt den Fehler zurück. Bei Erfolg wird die Auflösungsfunktion aufgerufen,
// um das Ergebnis zurückzugeben. Schließlich gibt die Funktion die Promise zurück, um das Ergebnis der Verzeichnisauflistung asynchron zu verarbeiten.
func Module_FS_ASYNC_readdirPromises(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readdir", 1)))
	}

	// Die Argumente werden extrahiert
	dirPath := parms.Arguments[0].String()

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

		// Der Ordnerinahlt wird eingelesen
		result, err := fsCoreReadDir(vmengine, jsruntime, dirPath)
		if err != nil {
			// Der Fehler wird zurückgegeben
			reject(err)
			return
		}

		// Die Daten werden zurückgegeben
		resolve(result)
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
