package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func extractAppendData(parm goja.Value) ([]byte, error) {
	var appendData []byte
	switch parm.ExportType().String() {
	case "string":
		appendData = []byte(parm.String())
	default:
		return nil, fmt.Errorf("invalid unsupported type")
	}
	return appendData, nil
}

func fsAppendFileCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, fileName string, appendData []byte) (goja.Value, error) {
	return nil, nil
}

func Module_FS_SYNC_appendFileSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "appendFile", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Der Appened Daten werden extrahiert
	appData, err := extractAppendData(parms.Arguments[1])
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Es wird versucht die Datei einzulesen
	result, err := fsAppendFileCore(vmengine, jsruntime, parms.Arguments[0].String(), appData)
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return result
}

func Module_FS_SYNC_appendFileCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "appendFile", 3)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Der Appened Daten werden extrahiert
	appData, err := extractAppendData(parms.Arguments[1])
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Callback Funktion wird zurückgegeben
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	var callbackParms goja.FunctionCall
	result, err := fsAppendFileCore(vmengine, jsruntime, parms.Arguments[0].String(), appData)
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

func Module_FS_SYNC_appendFilePromises(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "appendFile", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Der Appened Daten werden extrahiert
	appData, err := extractAppendData(parms.Arguments[1])
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

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
		result, err := fsAppendFileCore(vmengine, jsruntime, parms.Arguments[0].String(), appData)
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
