package fs

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func fsChownCore(vmengine vmpackage.VMInterface, jsruntime *goja.Runtime, descriptor interface{}) (goja.Value, error) {
	return nil, nil
}

func Module_FS_SYNC_chownSync(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chownCallback(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chownPromises(vmengine vmpackage.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}
