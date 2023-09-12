package jsengine

import "github.com/dop251/goja"

// Initalisiert alle Standard Funktionen
func (o *JSEngine) initRuntimeBaseFunctions(jsruntime *goja.Runtime) {
	// Die Requiere und Include Funktionen werden beschrieben
	req_inc_wrapper_func := func(call goja.FunctionCall) goja.Value {
		return o.runtimeRequire(call, jsruntime)
	}
	jsruntime.Set("require", req_inc_wrapper_func)
	jsruntime.Set("include", req_inc_wrapper_func)

	// Die Print Funktion wird geschrieben
	jsruntime.Set("println", o.runtimePrintln)

	// Die Consolen Funktionen werden geschrieben
	console_obj := jsruntime.NewObject()
	jsruntime.Set("console", console_obj)
	console_obj.Set("log", o.runtimeConsoleLog)

	// Die Fenster Funktionen werden geschrieben
	terminal_obj := jsruntime.NewObject()
	jsruntime.Set("terminal", terminal_obj)
	terminal_obj.Set("setWindowTitle", o.runtimeSetConsoleTitle)
}
