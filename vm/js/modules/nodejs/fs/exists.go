package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

// fsExistsCore überprüft die Existenz einer Datei oder eines Verzeichnisses im Dateisystem.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Verarbeitung und Rückgabe der Ergebnisse.
//   - filePath: Der Pfad zur zu überprüfenden Datei oder zum Verzeichnis.
//
// Rückgabewert:
//   - Ein boolescher Wert, der true ist, wenn die Datei oder das Verzeichnis existiert, andernfalls false.
//   - Im Falle eines Fehlers wird ein Fehlerwert zurückgegeben.
//
// Die Funktion ruft das Dateisystem über die vmengine-Schnittstelle auf, sucht die Datei oder das Verzeichnis und gibt true zurück, wenn es existiert, andernfalls false.
func fsExistsCore(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, filePath string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal error, not fs")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	fileResolve, err := fileSystem.GetFileByFullPath(filePath)
	if err != nil {
		return nil, err
	}

	// Die Daten werden zurückgegeben
	return jsruntime.ToValue(fileResolve != nil), nil
}

// Module_FS_SYNC_existsSync überprüft synchron, ob eine Datei oder ein Ordner im Dateisystem existiert.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Verarbeitung und Rückgabe der Ergebnisse.
//   - parms: Die FunktionCall-Struktur, die den übergebenen Dateipfad enthält.
//
// Rückgabewert:
//   - true, wenn die Datei oder das Verzeichnis existiert, andernfalls false.
//   - Im Falle eines Fehlers wird ein Fehlerwert zurückgegeben.
//
// Die Funktion überprüft die Existenz der Datei oder des Ordners im Dateisystem und gibt das Ergebnis zurück.
func Module_FS_SYNC_existsSync(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "exists", 1)))
	}

	// Die Argumente werden extrahiert
	exsPath := parms.Arguments[0].String()

	// Es wird versucht die Datei einzulesen
	result, err := fsExistsCore(vmengine, jsruntime, exsPath)
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return result
}

// Module_FS_SYNC_existsCallback überprüft synchron, ob eine Datei oder ein Ordner im Dateisystem existiert und ruft gegebenenfalls
// eine Callback-Funktion auf.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Verarbeitung und Rückgabe der Ergebnisse.
//   - parms: Die FunktionCall-Struktur, die die übergebenen Parameter enthält.
//
// Rückgabewert:
//   - Die Callback-Funktion wird aufgerufen (falls übergeben), andernfalls wird null zurückgegeben.
//
// Die Funktion überprüft die Existenz der Datei oder des Ordners im Dateisystem und gibt das Ergebnis über die Callback-Funktion zurück.
func Module_FS_SYNC_existsCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "exists", 2)))
	}

	// Die Argumente werden extrahiert
	exsPath := parms.Arguments[0].String()
	callbackFunction, isOk := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	result, err := fsExistsCore(vmengine, jsruntime, exsPath)
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Undefined(), jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{result, goja.Undefined()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

// Module_FS_ASYNC_existsPromises überprüft asynchron, ob eine Datei oder ein Ordner im Dateisystem existiert.
//
// Parameter:
//   - vmengine: Die Schnittstelle zur virtuellen Maschine, die das Dateisystem verwaltet.
//   - jsruntime: Ein Goja-Laufzeitumgebung zur Umwandlung des Ergebnisses in einen Goja-Wert.
//   - parms: Die FunktionCall-Struktur, die die übergebenen Parameter enthält.
//
// Rückgabewert:
//   - Eine Promise, die das Ergebnis der Existenzüberprüfung enthält (true, wenn existiert, sonst false).
//
// Die Funktion extrahiert den Dateipfad aus den übergebenen Parametern, überprüft die Existenz der Datei oder des Ordners im Dateisystem
// und gibt das Ergebnis über eine Promise zurück.
func Module_FS_ASYNC_existsPromises(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "exists", 1)))
	}

	// Die Argumente werden extrahiert
	exsPath := parms.Arguments[0].String()

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
		result, err := fsExistsCore(vmengine, jsruntime, exsPath)
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
