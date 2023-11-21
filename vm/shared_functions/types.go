package sharedfunctions

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

type LocalSharedFunctionCapsle struct {
	JsCall    func(goja.FunctionCall) goja.Value
	JsRuntime *goja.Runtime
}

type RemoteSharedFunctionCaplse struct {
	FunctionCallId     string
	CoreBridgeFunction vmpackage.CoreServiceBridgeInterface
}
