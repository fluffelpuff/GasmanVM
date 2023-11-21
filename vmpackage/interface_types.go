package vmpackage

import (
	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/core_service/argtypes"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
)

type SharedFunctionInterface interface {
	Call(...interface{}) (interface{}, error)
	ClientFunctionCreator() uint64
	IsLocal() bool
}

type CoreServiceBridgeInterface interface {
	CallSharedFunction(PackageIdentifyerInterface, string, string, uint64, []interface{}) (interface{}, error)
	RegisterSharedFunction(string, string, VMInterface) (*argtypes.RegisterSharedFunctionReturn, error)
	SetVM(VMInterface) error
	Setup(manifest *imagefile.Manifest) error
	Provide() error
}

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
	RegisterSharedFunction(string, string, string, CoreServiceBridgeInterface) error
	CallSharedFunction(PackageIdentifyerInterface, string, string, uint64, ...interface{}) (interface{}, error)
	ConsoleClear() error
	AddNewRoutine()
	RemoveRoutine()
}

type PackageIdentifyerInterface interface {
	GetID() string
	GetType() string
}
