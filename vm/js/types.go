package jsengine

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/imagefile"
)

type VMInterface interface {
	GetImageFile() *imagefile.ImageFile
}

type JSEngine struct {
	motherVM       VMInterface
	jsInterpreters []*goja.Runtime
}
