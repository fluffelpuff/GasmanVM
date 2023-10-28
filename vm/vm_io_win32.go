package vm

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

// Räumt die Console unter Windows auf
func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Setzt den Tittel der Console
func runtimeSetConsoleTitle(title string) error {
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitleW := kernel32DLL.NewProc("SetConsoleTitleW")

	ptrTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		panic(err)
	}

	ret, _, err := procSetConsoleTitleW.Call(uintptr(unsafe.Pointer(ptrTitle)))
	if ret == 0 {
		panic(err)
	}

	return nil
}
