package jsengine

import (
	"github.com/dop251/goja"
)

func NewEngine(mother_vm VMInterface) (*JSEngine, error) {
	// Das Basis Objekt wird erezgut
	base_bundle_runtime := &JSEngine{mother_vm, make([]*goja.Runtime, 0)}

	// Die Daten werden zurückgegeben
	return base_bundle_runtime, nil
}
