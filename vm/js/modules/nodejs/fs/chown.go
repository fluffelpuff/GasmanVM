package fs

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func fsChownCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, descriptor interface{}) (goja.Value, error) {
	return nil, nil
}

func Module_FS_SYNC_chownSync(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chownCallback(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}

func Module_FS_SYNC_chownPromises(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	return nil
}
