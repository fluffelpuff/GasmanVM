package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func fsAccessCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, path string, flag interface{}) (goja.Value, error) {
	return nil, nil
}

func Module_FS_SYNC_accessSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "access", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um eienn Integer hadnelt
	if parms.Arguments[1].ExportType().String() != "int64" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a flag")))
	}

	// Es wird versucht die Datei einzulesen
	_, err := fsAccessCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].ToInteger())
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return goja.Undefined()
}

func Module_FS_SYNC_accessCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "access", 3)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um eienn Integer hadnelt
	if parms.Arguments[1].ExportType().String() != "int64" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a flag")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Callback Funktion wird zurückgegeben
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	_, err := fsAccessCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].ToInteger())
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err.Error())}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

func Module_FS_SYNC_accessPromise(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "access", 1)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um eienn Integer hadnelt
	if parms.Arguments[1].ExportType().String() != "int64" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a flag")))
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
		result, err := fsAccessCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].ToInteger())
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
