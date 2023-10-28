package jsengine

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

type JSEngine struct {
	wg             sync.WaitGroup
	motherVM       modules.VMInterface
	jsInterpreters []*goja.Runtime
}
