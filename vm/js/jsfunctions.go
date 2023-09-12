package jsengine

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"github.com/dop251/goja"
)

// Setzt den Tittel der Console
func (o *JSEngine) runtimeSetConsoleTitle(title string) error {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitleW := kernel32DLL.NewProc("SetConsoleTitleW")

	ptrTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}

	ret, _, err := procSetConsoleTitleW.Call(uintptr(unsafe.Pointer(ptrTitle)))
	if ret == 0 {
		return err
	}

	return nil
}

// Zeigt ein Print an
func (o *JSEngine) runtimePrintln(call goja.FunctionCall) goja.Value {
	// Extrahieren Sie die Argumente, die an die print-Funktion übergeben wurden
	args := make([]interface{}, len(call.Arguments))
	for i, arg := range call.Arguments {
		args[i] = arg.Export()
	}

	// Ausgabe der Argumente auf der Go-Seite
	fmt.Println(args...)

	// Rückgabe eines leeren Werts
	return goja.Undefined()
}

// Zeigt einen Consolen Log an
func (o *JSEngine) runtimeConsoleLog(call goja.FunctionCall) goja.Value {
	// Extrahieren Sie die Argumente, die an die print-Funktion übergeben wurden
	args := make([]interface{}, len(call.Arguments))
	for i, arg := range call.Arguments {
		args[i] = arg.Export()
	}

	// Ausgabe der Argumente auf der Go-Seite
	log.Println(args...)

	// Rückgabe eines leeren Werts
	return goja.Undefined()
}
