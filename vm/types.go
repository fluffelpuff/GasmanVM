package vm

import (
	"sync"

	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

type SharedFunctionInterface interface {
	Call(...interface{}) (interface{}, error)
	ClientFunctionCreator() uint64
	IsLocal() bool
}

type ScriptContainerVM struct {
	wait                       *sync.WaitGroup
	fsContainer                *fsysfile.FileSystem
	imageFile                  *imagefile.ImageFileReader
	jsEngine                   *jsengine.JSEngine
	coreServiceBridgeInterface modules.CoreServiceBridgeInterface
	groupFunctionShares        map[string]map[string]SharedFunctionInterface
	wasExited                  bool
}
