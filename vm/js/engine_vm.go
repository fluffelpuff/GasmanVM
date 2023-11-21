package jsengine

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

// Gibt die VM Informationen zurück
func runtimeVMGetInfo(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, call goja.FunctionCall) goja.Value {
	// Es wird geprüft ob keine Argumente an die Funktion übergeben wurden
	if len(call.Arguments) != 0 {
		panic("the function does not require any parameters")
	}

	// Es wird ein neues Objekt erstellt welches die Informationen der VM enthält
	vmInfoObject := jsruntime.NewObject()

	// Die Aktuelle Version wird angegeben
	vmInfoObject.Set("version", vmpackage.VERSION)

	// Der Name der VM und die VM Architektur werden angegeben
	vmInfoObject.Set("arc", "XSCRIPT_EMC_100")

	// Der Öffentliche Schlüssel der VM wird geschrieben
	vmInfoObject.Set("vmpkey", "<PUBLIC-KEY>")

	// Der Öffentliche Node Schlüssel wird geschrieben
	vmInfoObject.Set("nodepkey", "<PUBLIC-NODE-KEY>")

	// Die Berechtigungen werden geschrieben
	vmInfoObject.Set("permissions", jsruntime.ToValue(vmengine.GetPermissions()))

	// Rückgabe eines leeren Werts
	return vmInfoObject
}

// Die VM Funktionen werden bereitsgestellt
func (o *JSEngine) getVMModuleFunctions(jsruntime *goja.Runtime) goja.Value {
	// Das VM Objekt wird erstellt
	vmObject := jsruntime.NewObject()

	// Diese Funktion gibt die Aktuellen Informationen über die VM aus
	err := vmObject.Set("getInfo", func(parms goja.FunctionCall) goja.Value {
		return runtimeVMGetInfo(o.motherVM, jsruntime, parms)
	})
	if err != nil {
		panic(err)
	}

	// Diese Funktion wird verwendet um die VM zu beenden
	err = vmObject.Set("kill", func(call goja.FunctionCall) goja.Value {
		jsruntime.Interrupt(o.getCloserValue())
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Diese Funktion wird verwendet um sicherzustellen das die VM weitläuft, auch wenn das Mainscript durchgelaufen ist
	err = vmObject.Set("enableHyperWaitMode", func(call goja.FunctionCall) goja.Value {
		// Es wird ermittelt wieviele Parameter angegeben wurden
		if len(call.Arguments) != 0 {
			panic("the 'hyperWait' functio dosen't need parameters")
		}

		// Der HyperWait Modus wird Aktiviert
		o.enableHyperWaitMode()

		// Es wird Signalisiert dass der HyperWait Modus Aktiviert werden soll
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Das Objekt wird zurückgegeben
	return vmObject
}
