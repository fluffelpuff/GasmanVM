package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsCloseCore(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, descriptor string) error {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return fmt.Errorf("internal fs error")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	if err := fileSystem.CloseFileDescriptor(descriptor); err != nil {
		return err
	}

	// Es ist kein Fehler aufgetreten oder sonstiges
	return nil
}

// Module_FS_SYNC_closeSync ist eine Funktion, die dazu dient, eine Datei oder einen Dateideskriptor synchron zu schließen.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Eine FunktionCall-Struktur, die die übergebenen Parameter enthält.
//
// Rückgabewert:
// Diese Funktion gibt das Ergebnis des Schließvorgangs zurück, normalerweise null, um anzuzeigen, dass der Schließvorgang abgeschlossen ist.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (1 Parameter) vorhanden ist. Andernfalls wird ein Fehler ausgelöst.
// 2. Der übergebene Dateideskriptor (descriptor) wird extrahiert.
// 3. Es wird versucht, die Datei oder den Dateideskriptor zu schließen und dabei wird das Ergebnis und mögliche Fehler ermittelt.
// 4. Abhängig vom Schließergebnis wird das Ergebnis zurückgegeben, normalerweise null, um anzuzeigen, dass der Schließvorgang synchron abgeschlossen wurde.
//
// Hinweis:
// Diese Funktion wird normalerweise nicht direkt aufgerufen, sondern von JavaScript-Code verwendet, um Dateien synchron zu schließen.
func Module_FS_SYNC_closeSync(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "closeSync", 1)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird versucht die Datei einzulesen
	err := fsCloseCore(vmengine, jsruntime, parms.Arguments[0].String())
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

// Module_FS_SYNC_closeCallback ist eine Funktion, die dazu dient, eine Datei oder einen Dateideskriptor synchron zu schließen und eine Callback-Funktion aufzurufen.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Eine FunktionCall-Struktur, die die übergebenen Parameter enthält.
//
// Rückgabewert:
// Diese Funktion gibt null (goja.Null()) zurück, um anzuzeigen, dass der Schließvorgang abgeschlossen ist.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (2 Parameter) vorhanden ist. Andernfalls wird ein Fehler ausgelöst.
// 2. Der übergebene Dateideskriptor (descriptor) und die Callback-Funktion werden extrahiert.
// 3. Es wird versucht, die Datei oder den Dateideskriptor zu schließen und dabei wird das Ergebnis und mögliche Fehler ermittelt.
// 4. Abhängig vom Schließergebnis wird eine Callback-Funktion aufgerufen, der das Ergebnis oder ein Fehler übergeben wird.
// 5. Die Funktion gibt null zurück, um anzuzeigen, dass der Schließvorgang synchron abgeschlossen wurde.
//
// Hinweis:
// Diese Funktion wird normalerweise nicht direkt aufgerufen, sondern von JavaScript-Code verwendet, um Dateien synchron zu schließen und eine Callback-Funktion aufzurufen.
func Module_FS_SYNC_closeCallback(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "close", 2)))
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
	callbackFunction, isOk := parms.Arguments[1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Die Datei wird geschlossen
	err := fsCloseCore(vmengine, jsruntime, parms.Arguments[0].String())
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null()}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

// Module_FS_ASYNC_closePromises ist eine Funktion, die dazu dient, eine Datei oder einen Dateideskriptor asynchron zu schließen und ein Promise-Objekt zurückzugeben.
//
// Parameter:
// - vmengine: Das VMInterface-Objekt, das die Verbindung zur virtuellen Maschine bereitstellt.
// - jsruntime: Das goja.Runtime-Objekt für die JavaScript-Ausführung.
// - parms: Eine FunktionCall-Struktur, die die übergebenen Parameter enthält.
//
// Rückgabewert:
// Diese Funktion gibt ein Promise zurück, das entweder den erfolgreichen Abschluss des Schließens oder einen Fehler darstellt.
//
// Funktionsweise:
// 1. Die Funktion überprüft, ob die erforderliche Anzahl von Parametern (1 Parameter) vorhanden ist. Andernfalls wird ein Fehler ausgelöst.
// 2. Der übergebene Dateideskriptor (descriptor) wird extrahiert.
// 3. Es wird ein neues Promise und die zugehörigen Resolving-Funktionen (resolve und reject) erstellt.
// 4. Ein neuer Leseprozess (Routine) wird im vmengine registriert.
// 5. Die Datei oder der Dateideskriptor wird asynchron geschlossen.
// 6. Nach Abschluss des Schließvorgangs wird entweder das Ergebnis (Erfolg) oder ein Fehler über das Promise-Objekt zurückgegeben.
//
// Hinweis:
// Diese Funktion wird normalerweise nicht direkt aufgerufen, sondern von JavaScript-Code verwendet, um asynchron Dateien zu schließen.
func Module_FS_ASYNC_closePromises(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "close", 1)))
	}

	// Die Argumente werden extrahiert
	descriptor := parms.Arguments[0].String()

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
		if err := fsCloseCore(vmengine, jsruntime, descriptor); err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(goja.Null())
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
