package jsengine

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
	"github.com/fluffelpuff/GasmanVM/vm/utils"
)

// runtimeGroupShareCallback ist eine Callback-Funktion, die von der Runtime-Gruppe aufgerufen wird, um eine geteilte Funktion zu registrieren.
// Die Funktion überprüft die Parameter und registriert die Funktion in der VM-Engine.
// Die Funktion gibt keinen Wert zurück, sondern verwendet eine Callback-Funktion zur Weitergabe von Ergebnissen oder Fehlern.
// - `vmengine modules.VMInterface`: Die VM-Engine, in der die geteilte Funktion registriert wird.
// - `jsruntime *goja.Runtime`: Die Goja-Laufzeitumgebung, in der die Callback-Funktion ausgeführt wird.
// - `call goja.FunctionCall`: Die Argumente für die Callback-Funktion.
// - `goja.Value`: Der Rückgabewert der Callback-Funktion (Undefined).
func runtimeGroupShareCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, call goja.FunctionCall) goja.Value {
	// Es wird geprüft, ob keine Argumente an die Funktion übergeben wurden.
	if len(call.Arguments) != 4 {
		panic(goja.New().NewGoError(fmt.Errorf("the function 'shareFunction' requires 4 parameters")))
	}

	// Es wird ermittelt, ob es sich bei der vierten Variable um eine Funktion handelt.
	if call.Arguments[3].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the 'shareFunction' requires a callback function")))
	}

	// Die Callback-Funktion wird eingelesen.
	callbackFunction, isOkCallback := call.Arguments[3].Export().(func(goja.FunctionCall) goja.Value)
	if !isOkCallback {
		panic(goja.New().NewGoError(fmt.Errorf("the 'shareFunction' requires a callback function")))
	}

	// Es wird ermittelt, ob es sich um einen String handelt.
	if call.Arguments[0].ExportType().String() != "string" {
		// Der Fehler wird zurückgegeben
		callbackFunction(goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(fmt.Errorf("the name of the shared function must be specified using a string"))}})

		// Die Funktion gibt keinen Wert zurück
		return goja.Undefined()
	}

	// Es wird ermittelt, ob es sich bei der dritten Variable um eine Funktion handelt.
	if call.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		// Der Fehler wird zurückgegeben
		callbackFunction(goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(fmt.Errorf("you need to create as a function"))}})

		// Die Funktion gibt keinen Wert zurück
		return goja.Undefined()
	}

	// Die geteilte Funktion wird eingelesen.
	sharedFunction, isOkShareFunction := call.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOkShareFunction {
		// Der Fehler wird zurückgegeben
		callbackFunction(goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(fmt.Errorf("you need to create as a function"))}})

		// Die Funktion gibt keinen Wert zurück
		return goja.Undefined()
	}

	// Die Share-Funktion wird aufgerufen.
	err := vmengine.RegisterLocalSharedFunction(call.Arguments[0].String(), "<GROUP_NAME>", sharedFunction, jsruntime)
	var functionParams goja.FunctionCall
	if err != nil {
		functionParams = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		functionParams = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null()}}
	}

	// Die Callback-Funktion wird aufgerufen.
	callbackFunction(functionParams)

	// Rückgabe eines leeren Werts.
	return goja.Undefined()
}

// runtimeGroupFunctionCallCallback ist eine Callback-Funktion, die von der Runtime-Gruppe aufgerufen wird.
// Die Funktion überprüft die Parameter und ruft eine geteilte Funktion in Go mit den entsprechenden Parametern auf.
// Die Funktion gibt keinen Wert zurück, sondern verwendet eine Callback-Funktion zur Weitergabe von Ergebnissen oder Fehlern.
// - `vmengine modules.VMInterface`: Die VM-Engine, die die geteilte Funktion aufruft.
// - `jsruntime *goja.Runtime`: Die Goja-Laufzeitumgebung, in der die Callback-Funktion ausgeführt wird.
// - `call goja.FunctionCall`: Die Argumente für die Callback-Funktion.
// - `goja.Value`: Der Rückgabewert der Callback-Funktion (Undefined).
func runtimeGroupFunctionCallCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, call goja.FunctionCall) goja.Value {
	// Es wird geprüft, ob keine Argumente an die Funktion übergeben wurden.
	if len(call.Arguments) < 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function 'callFunction' requires 4 parameters")))
	}

	// Es wird ermittelt, ob es sich bei der dritten Variable um eine Funktion handelt.
	if call.Arguments[len(call.Arguments)-1].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the 'callFunction' requires a callback function")))
	}

	// Die Callback-Funktion wird zurückgegeben.
	callbackFunction, isOk := call.Arguments[len(call.Arguments)-1].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(goja.New().NewGoError(fmt.Errorf("the 'callFunction' requires a callback function")))
	}

	// Es wird ermittelt, ob es sich um einen String handelt.
	if call.Arguments[0].ExportType().String() != "string" {
		// Der Fehler wird zurückgegeben
		callbackFunction(goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(fmt.Errorf("the name of the shared function must be specified using a string"))}})

		// Die Funktion gibt keinen Wert zurück
		return goja.Undefined()
	}

	// Die Parameter werden ausgelesen.
	retrievedItems := make([]interface{}, 0)
	for _, item := range call.Arguments[2 : len(call.Arguments)-1] {
		// Das Item wird Exportiert
		exportedItem := item.Export()

		// Es wird ermittelt ob es sich um einen zulässigen Wert handelt
		if !utils.CheckDataValues(exportedItem) {
			// Der Fehler wird zurückgegeben
			callbackFunction(goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(fmt.Errorf("an inadmissible value was recognized"))}})

			// Die Funktion gibt keinen Wert zurück
			return goja.Undefined()
		}

		// Der Wert wird zwischengespeichert
		retrievedItems = append(retrievedItems, exportedItem)
	}

	// Es wird versucht, die geteilte Funktion aufzurufen.
	result, err := vmengine.CallSharedFunction(call.Arguments[0].String(), "<GROUP_NAME>", retrievedItems...)
	var callbackParams goja.FunctionCall
	if err != nil {
		callbackParams = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParams = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null(), jsruntime.ToValue(result)}}
	}

	// Die Callback-Funktion wird aufgerufen.
	callbackFunction(callbackParams)

	// Rückgabe eines leeren Werts.
	return goja.Undefined()
}
