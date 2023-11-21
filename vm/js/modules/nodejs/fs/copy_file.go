package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

// fsCopyFileCore ist eine interne Funktion, die dazu dient, eine Datei von der Quelladresse zur Zieladresse zu kopieren.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - sourcePath: Der Dateipfad der Quelle, von dem die Datei kopiert werden soll.
// - destinationPath: Der Dateipfad des Ziels, an den die Datei kopiert werden soll.
//
// Rückgabewert:
// Diese Funktion gibt ein Undefined zurück, wenn die Kopie erfolgreich war. Im Fehlerfall wird ein Fehlerwert zurückgegeben.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob das Dateisystem (fileSystem) verfügbar ist. Wenn nicht, wird ein Fehler ausgelöst.
// 2. Die Datei wird von der Quelladresse zur Zieladresse kopiert.
// 3. Wenn die Kopie erfolgreich war, wird ein Undefined zurückgegeben, andernfalls wird ein Fehlerwert zurückgegeben.
//
// Hinweis:
// Diese Funktion wird normalerweise nicht direkt aufgerufen, sondern von anderen Funktionen wie "Module_FS_SYNC_copyFile" verwendet.
func fsCopyFileCore(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, sourcePath string, destinationPath string) error {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return fmt.Errorf("internal error, not fs")
	}

	// Die Datei wird kopiert
	if err := fileSystem.CopyFile(sourcePath, destinationPath); err != nil {
		return err
	}

	// Es wird ein Undefined zurückgegeben
	return nil
}

// Module_FS_SYNC_copyFile ist eine synchrone Funktion, die dazu dient, eine Datei von der Quelladresse zur Zieladresse zu kopieren und das Ergebnis zurückzugeben.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Ein goja.FunctionCall-Objekt, das die Funktion und ihre Argumente darstellt.
//
// Rückgabewert:
// Diese Funktion gibt das kopierte Ergebnis (der Dateipfad des Zielorts) zurück, wenn die Kopie erfolgreich war. Im Fehlerfall wird ein Fehlerwert zurückgegeben.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (zwei) vorhanden ist, andernfalls wird ein Fehler ausgelöst.
// 2. Die Quell- und Ziel-Pfade werden aus den Funktionsparametern extrahiert.
// 3. Es wird versucht, die Datei von der Quell- zur Zieladresse zu kopieren.
// 4. Das Ergebnis des Kopiervorgangs (der Dateipfad des Zielorts) oder ein Fehler wird zurückgegeben, abhängig von Erfolg oder Misserfolg.
//
// Beispiel (JavaScript):
//
//	try {
//	    var result = Module_FS_SYNC_copyFile(vmengine, sourcePath, destPath);
//	    console.log("Datei erfolgreich kopiert: " + result);
//	} catch (error) {
//
//	    console.error("Fehler beim Kopieren der Datei: " + error);
//	}
func Module_FS_SYNC_copyFileSync(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "copyFile", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird versucht die Datei einzulesen
	if err := fsCopyFileCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String()); err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

// Module_FS_SYNC_copyFileCallback ist eine synchrone Funktion, die dazu dient, eine Datei von der Quelladresse zur Zieladresse zu kopieren und das Ergebnis über eine Callback-Funktion zurückzugeben.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Ein goja.FunctionCall-Objekt, das die Funktion und ihre Argumente darstellt.
//
// Rückgabewert:
// Diese Funktion gibt null zurück.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (drei) vorhanden ist, andernfalls wird ein Fehler ausgelöst.
// 2. Die Quell- und Ziel-Pfade werden aus den Funktionsparametern extrahiert.
// 3. Die Callback-Funktion wird aus den Funktionsparametern extrahiert.
// 4. Es wird versucht, die Datei von der Quell- zur Zieladresse zu kopieren.
// 5. Das Ergebnis des Kopiervorgangs oder ein Fehler wird in Abhängigkeit von Erfolg oder Misserfolg an die Callback-Funktion übergeben.
// 6. Die Funktion gibt null zurück, da das Ergebnis über die Callback-Funktion verarbeitet wird.
//
// Beispiel (JavaScript):
//
//	Module_FS_SYNC_copyFileCallback(vmengine, sourcePath, destPath, function(result, error) {
//	    if (error) {
//	        console.error("Fehler beim Kopieren der Datei: " + error);
//	    } else {
//	        console.log("Datei erfolgreich kopiert: " + result);
//	    }
//	});
func Module_FS_SYNC_copyFileCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "copyFile", 3)))
	}

	// Die Argumente werden extrahiert
	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Callback Funktion wird eingelesen
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Die Datei wird Kopiert
	err := fsCopyFileCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String())
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

// Module_FS_ASYNC_copyFilePromises ist eine asynchrone Funktion, die dazu dient, eine Datei von der Quelladresse zur Zieladresse zu kopieren und das Ergebnis über eine Promise zurückzugeben.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Ein goja.FunctionCall-Objekt, das die Funktion und ihre Argumente darstellt.
//
// Rückgabewert:
// Diese Funktion gibt eine Promise zurück, die das Ergebnis des Kopiervorgangs repräsentiert. Die Promise wird entweder aufgelöst, wenn der Vorgang erfolgreich abgeschlossen ist, oder abgelehnt, wenn ein Fehler auftritt.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (zwei) vorhanden ist, andernfalls wird ein Fehler ausgelöst.
// 2. Die Quell- und Ziel-Pfade werden aus den Funktionsparametern extrahiert.
// 3. Es wird eine neue Promise erstellt, die zur Rückgabe des Ergebnisses des Kopiervorgangs verwendet wird.
// 4. Ein neuer asynchroner Vorgang wird registriert, um die Aufgabe auszuführen.
// 5. Die Datei wird asynchron von der Quelle zur Zieladresse kopiert. Wenn die Aufgabe abgeschlossen ist, wird die Promise entweder aufgelöst, wenn erfolgreich, oder abgelehnt, wenn ein Fehler auftritt.
// 6. Das Ergebnis des Kopiervorgangs wird entweder über die "resolve"- oder "reject"-Funktion der Promise zurückgegeben.
//
// Beispiel:
// promise := Module_FS_ASYNC_copyFilePromises(vmengine, jsruntime, parms)
//
//	promise.Then(func(result goja.Value) {
//	    // Erfolgreich kopiert, result enthält das Ergebnis.
//	}).Catch(func(err goja.Value) {
//
//	    // Fehler beim Kopieren, err enthält die Fehlerinformation.
//	})
func Module_FS_ASYNC_copyFilePromises(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "copyFile", 2)))
	}

	// Die Argumente werden extrahiert
	sourcePath := parms.Arguments[0].String()
	destPath := parms.Arguments[1].String()

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
		if err := fsCopyFileCore(vmengine, jsruntime, sourcePath, destPath); err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(goja.Undefined())
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
