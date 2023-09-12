package main

import (
	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
)

type ScriptContainerVM struct {
	imageFile *imagefile.ImageFile
	jsEngine  *jsengine.JSEngine
}
