package modules

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
)

type VMInterface interface {
	GetFilesystem() *fsysfile.FileSystem
	GetImageFile() *imagefile.ImageFileReader
	GetPermissions() []string
	ConsolePrint([]interface{}) error
	ConsolePrintln([]interface{}) error
	ConsoleLog([]interface{}) error
	ConsoleInfo([]interface{}) error
	ConsoleError([]interface{}) error
	SetWindowTitle(string) error
	RegisterLocalSharedFunction(string, string, func(goja.FunctionCall) goja.Value, *goja.Runtime) error
	CallSharedFunction(string, string, ...interface{}) (interface{}, error)
	ConsoleClear() error
	AddNewRoutine()
	RemoveRoutine()
}

type CoreServiceBridgeInterface interface {
}
