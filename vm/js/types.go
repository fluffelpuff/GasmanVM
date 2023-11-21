package jsengine

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

type JSEngine struct {
	wg             sync.WaitGroup
	motherVM       vmpackage.VMInterface
	mutex          *sync.Mutex
	jsInterpreters []*goja.Runtime
	hyperWait      bool
}

type PackageIdentifyer struct {
	Type string
	Id   string
}

func (o *PackageIdentifyer) GetID() string {
	return o.Id
}

func (o *PackageIdentifyer) GetType() string {
	return o.Type
}
