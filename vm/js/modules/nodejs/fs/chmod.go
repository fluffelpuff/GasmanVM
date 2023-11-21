package fs

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsChmodCore(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, descriptor interface{}) (goja.Value, error) {
	return nil, nil
}

func Module_FS_SYNC_chmodSync(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chmodCallback(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chmodPromises(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}
