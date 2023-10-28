package sharedfunctions

import "github.com/dop251/goja"

type SharedFunctionCapsle struct {
	JsCall    func(goja.FunctionCall) goja.Value
	JsRuntime *goja.Runtime
}
