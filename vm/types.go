package vm

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

type ScriptContainerVM struct {
	wait                       *sync.WaitGroup
	fsContainer                *fsysfile.FileSystem
	imageFile                  *imagefile.ImageFileReader
	jsEngine                   *jsengine.JSEngine
	coreServiceBridgeInterface vmpackage.CoreServiceBridgeInterface
	wasExited                  bool
	functionSharingIdMap       map[string]func(goja.FunctionCall) goja.Value
}
