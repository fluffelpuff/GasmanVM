//go:build windows

package jsengine

import "github.com/dop251/goja"

// Stellt die Nativen Funktionen bereit
func (o *JSEngine) loadNativeModule(jsruntime *goja.Runtime) goja.Value {
	// Das Modul für Native Libs wird erstellt
	nativeLibModule := jsruntime.NewObject()
	nativeLibModule.Set("loadLibrary", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// das Module wird zurückgegeben
	return nativeLibModule
}
