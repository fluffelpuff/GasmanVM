package jsengine

import (
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

// Diese Funktion stellt die Top Funktionen bereit
func (o *JSEngine) loadMainFunctions(jsruntime *goja.Runtime) error {
	// Die Require Funktion wird definiert
	err := jsruntime.Set("require", func(call goja.FunctionCall) goja.Value {
		return o.runtimeRequire(call, jsruntime)
	})
	if err != nil {
		return err
	}

	// Die Includ Funktion wird definiert
	err = jsruntime.Set("include", func(call goja.FunctionCall) goja.Value {
		return o.runtimeRequire(call, jsruntime)
	})
	if err != nil {
		return err
	}

	// Die Println Funktion wird definiert
	err = jsruntime.Set("println", func(parms goja.FunctionCall) goja.Value {
		// Die Printline Werte werden ausgelesen
		printValues := convertoToInterfacesArray(parms.Arguments)

		// Der Inhalt wird angezeigt
		o.motherVM.ConsolePrintln(printValues)

		// Die Werte werden an die VM übergeben
		return goja.Undefined()
	})
	if err != nil {
		return err
	}

	// Die Print Funktion wird definiert
	err = jsruntime.Set("print", func(parms goja.FunctionCall) goja.Value {
		// Die Printline Werte werden ausgelesen
		printValues := convertoToInterfacesArray(parms.Arguments)

		// Der Inhalt wird angezeigt
		o.motherVM.ConsolePrint(printValues)

		// Die Werte werden an die VM übergeben
		return goja.Undefined()
	})
	if err != nil {
		return err
	}

	// Wird verwendet um die Ausführung zu beenden
	err = jsruntime.Set("exit", func(parms goja.FunctionCall) goja.Value {
		jsruntime.Interrupt(o.getCloserValue())
		return goja.Undefined()
	})
	if err != nil {
		return err
	}

	// Es ist kein Fehler aufgetreten
	return nil
}

// Stellt die Group Funktionen bereit
func (o *JSEngine) runtimeGroupFunctions(jsruntime *goja.Runtime) goja.Value {
	// Das Group Objekt wird erstellt
	groupObject := jsruntime.NewObject()

	// Die 'share' Funktion wird festgelegt
	err := groupObject.Set("shareFunction", func(parms goja.FunctionCall) goja.Value {
		return runtimeGroupShareCallback(o.motherVM, jsruntime, parms)
	})
	if err != nil {
		panic(err)
	}

	// Die 'call' Funktion wird gestgelegt
	err = groupObject.Set("callFunction", func(parms goja.FunctionCall) goja.Value {
		return runtimeGroupFunctionCallCallback(o.motherVM, jsruntime, parms)
	})
	if err != nil {
		panic(err)
	}

	// Das Group Objekt wird zurückgegeben
	return groupObject
}

// Gibt die Consolen Funktion zurück
func (o *JSEngine) getConsoleModuleFunctions(jsruntime *goja.Runtime) goja.Value {
	// Die Consolen Funktionen werden bereitgestellt
	consoleObject := jsruntime.NewObject()

	// Gibt einen Log an
	err := consoleObject.Set("log", func(parms goja.FunctionCall) goja.Value {
		// Die Argumente werden Exportiert
		args := convertoToInterfacesArray(parms.Arguments)

		// Der Log Eintrag wird angezeigt
		o.motherVM.ConsoleLog(args)

		// Es wird ein Undefined zurückgegeben
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Gibt einen Info Text aus
	err = consoleObject.Set("info", func(parms goja.FunctionCall) goja.Value {
		// Die Argumente werden Exportiert
		args := convertoToInterfacesArray(parms.Arguments)

		// Der Log Eintrag wird angezeigt
		o.motherVM.ConsoleInfo(args)

		// Es wird ein Undefined zurückgegeben
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Zeigt einen Fehler an
	err = consoleObject.Set("error", func(parms goja.FunctionCall) goja.Value {
		// Die Argumente werden Exportiert
		args := convertoToInterfacesArray(parms.Arguments)

		// Der Log Eintrag wird angezeigt
		o.motherVM.ConsoleError(args)

		// Es wird ein Undefined zurückgegeben
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Löscht den Inahlt der Console
	err = consoleObject.Set("clear", func(parms goja.FunctionCall) goja.Value {
		// Die Console wird geleert
		o.motherVM.ConsoleClear()

		// Es wird ein Undefined zurückgegeben
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	// Legt den Titel einer Konsole fest
	err = consoleObject.Set("setWindowTitle", func(parms goja.FunctionCall) goja.Value {
		// Sollte mehr als 1 Argument vorhanden sein, wird der Vorgang abgebrochen
		if len(parms.Arguments) != 1 {
			panic("setWindowTitle need one argument, the new window title")
		}

		// Der Titel wird ausgelesen
		readTitle := parms.Arguments[0].String()

		// Die Console wird geleert
		o.motherVM.SetWindowTitle(readTitle)

		// Es wird ein Undefined zurückgegeben
		return goja.Undefined()
	})
	if err != nil {
		panic(err)
	}

	//consoleObject.Set("groupEnd", o.runtimeConsoleLog)
	//consoleObject.Set("groupCollapsed", o.runtimeConsoleLog)
	//consoleObject.Set("group", o.runtimeConsoleLog)
	//consoleObject.Set("debug", o.runtimeConsoleLog)
	//consoleObject.Set("count", o.runtimeConsoleLog)
	//consoleObject.Set("assert", o.runtimeConsoleLog)
	//consoleObject.Set("table", o.runtimeConsoleLog)
	//consoleObject.Set("time", o.runtimeConsoleLog)
	//consoleObject.Set("timeEnd", o.runtimeConsoleLog)
	//consoleObject.Set("timeLog", o.runtimeConsoleLog)

	// Dasa Objekt wird zurückgegeebn
	return consoleObject
}

// Stellt die Driver Funktionen bereit
func (o *JSEngine) runtimeDriverFunctions(jsruntime *goja.Runtime) goja.Value {
	return nil
}

// Initalisiert alle Standard Funktionen
func (o *JSEngine) initRuntimeBaseFunctions(jsruntime *goja.Runtime) error {
	// Die Basis Funktionen werden geladen
	if err := o.loadMainFunctions(jsruntime); err != nil {
		return fmt.Errorf("initRuntimeBaseFunctions: " + err.Error())
	}

	// Die Console Funktionen werden geladen
	if err := jsruntime.Set("console", o.getConsoleModuleFunctions(jsruntime)); err != nil {
		return fmt.Errorf("initRuntimeBaseFunctions: " + err.Error())
	}

	// Die VM Funktionen wird geladen
	if err := jsruntime.Set("vm", o.getVMModuleFunctions(jsruntime)); err != nil {
		return fmt.Errorf("initRuntimeBaseFunctions: " + err.Error())
	}

	// Die Treiber Modul Funktionen werden geladen
	if err := jsruntime.Set("driver", o.runtimeDriverFunctions(jsruntime)); err != nil {
		return fmt.Errorf("initRuntimeBaseFunctions: " + err.Error())
	}

	// Die Group Modul Funktionen werden geladen
	if err := jsruntime.Set("group", o.runtimeGroupFunctions(jsruntime)); err != nil {
		return fmt.Errorf("initRuntimeBaseFunctions: " + err.Error())
	}

	// Es ist kein Fehler aufgetreten
	return nil
}

// Gibt den Aktuellen Closer Wert zurück
func (o *JSEngine) getCloserValue() string {
	return "__vmintravalue:0a4ec9d8dce6bbf33a230c16443a20447bc3a7616d813ab1b1525fb95c1392f7"
}

// Aktiviert den HyperWait Modus
func (o *JSEngine) enableHyperWaitMode() {
	// Der Mutex wird verwendet
	o.mutex.Lock()
	defer o.mutex.Unlock()

	// Der Wert wird Aktualisiert, sofern er nicht bereits Aktiviert wurde
	if !o.hyperWait {
		// Der Wert wird gesetzt
		o.hyperWait = true

		// Es wird ein Waiter auf VM ebene hinzugefügt
		o.motherVM.AddNewRoutine()

		// DEBUG
		fmt.Println("Hyper waiting is active!")
	}
}

// Erstellt eine neue Engine
func NewEngine(mother_vm vmpackage.VMInterface) (*JSEngine, error) {
	// Das Basis Objekt wird erezgut
	base_bundle_runtime := &JSEngine{
		motherVM:       mother_vm,
		jsInterpreters: make([]*goja.Runtime, 0),
		hyperWait:      false,
		mutex:          new(sync.Mutex),
	}

	// Die Daten werden zurückgegeben
	return base_bundle_runtime, nil
}
