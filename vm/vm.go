package vm

import (
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
	sharedfunctions "github.com/fluffelpuff/GasmanVM/vm/shared_functions"
)

func (o *ScriptContainerVM) Run() error {
	// Die Main Datei wird gestartet
	_, err := o.jsEngine.RunMainScript()
	if err != nil {
		return err
	}

	// Der Vorgang wird druchgeführt
	return nil
}

func (o *ScriptContainerVM) setEngines(jsengine *jsengine.JSEngine) error {
	// Es wird geprüft ob die Engines bereits festgelegt wurden
	if o.jsEngine != nil {
		return fmt.Errorf("aborted, engines always seted")
	}

	// Die Javascript Engine wird festgelegt
	o.jsEngine = jsengine

	// Fertig
	return nil
}

func (o *ScriptContainerVM) GetImageFile() *imagefile.ImageFileReader {
	return o.imageFile
}

func (o *ScriptContainerVM) GetFilesystem() *fsysfile.FileSystem {
	return o.fsContainer
}

func (o *ScriptContainerVM) Exit() {
	o.wasExited = true
}

func (o *ScriptContainerVM) GetPermissions() []string {
	return []string{"ABC", "DEF"}
}

func (o *ScriptContainerVM) ConsolePrint(values []interface{}) error {
	// Ausgabe der Argumente auf der Go-Seite
	fmt.Print(exportStringArrayFromArgumentsAndFromateIt(values)...)
	return nil
}

func (o *ScriptContainerVM) ConsolePrintln(values []interface{}) error {
	// Ausgabe der Argumente auf der Go-Seite
	fmt.Println(exportStringArrayFromArgumentsAndFromateIt(values)...)
	return nil
}

func (o *ScriptContainerVM) ConsoleLog(values []interface{}) error {
	// Ausgabe der Argumente auf der Go-Seite
	return runtimeConsoleLog(values)
}

func (o *ScriptContainerVM) ConsoleInfo(values []interface{}) error {
	// Ausgabe der Argumente auf der Go-Seite
	return runtimeInfoLog(values)
}

func (o *ScriptContainerVM) ConsoleError(values []interface{}) error {
	// Ausgabe der Argumente auf der Go-Seite
	return runtimeErrorLog(values)
}

func (o *ScriptContainerVM) ConsoleClear() error {
	clearConsole()
	return nil
}

func (o *ScriptContainerVM) SetWindowTitle(windowTitleValue string) error {
	return runtimeSetConsoleTitle(windowTitleValue)
}

func (o *ScriptContainerVM) RegisterLocalSharedFunction(shareName string, groupName string, sharedFunction func(goja.FunctionCall) goja.Value, runtime *goja.Runtime) error {
	// Die Funktion wird zwischengespeichert
	o.groupFunctionShares[groupName] = make(map[string]SharedFunctionInterface)
	o.groupFunctionShares[groupName][shareName] = &sharedfunctions.SharedFunctionCapsle{JsCall: sharedFunction, JsRuntime: runtime}

	// Es ist kein Fehler aufgetreten
	return nil
}

func (o *ScriptContainerVM) CallSharedFunction(shareName string, groupName string, parms ...interface{}) (interface{}, error) {
	// Es wird ermittelt ob die Gruppe vorhanden ist
	groupResponse, foundGroup := o.groupFunctionShares[groupName]
	if !foundGroup {
		return nil, fmt.Errorf("group '%s' not found", groupName)
	}

	// Es wird ermittelt ob die Funktion in der Gruppe vorhanden ist
	functionResponse, foundFunction := groupResponse[shareName]
	if !foundFunction {
		return nil, fmt.Errorf("function '%s' not found", shareName)
	}

	// Die Funktion wird aufgerufen
	functionResult, functionErr := functionResponse.Call(parms...)
	if functionErr != nil {
		return nil, functionErr
	}

	// Das Ergebniss wird zurückgegeben
	return functionResult, nil
}

func NewRuntime(fsContainer *fsysfile.FileSystem, imageFile *imagefile.ImageFileReader, coreController modules.CoreServiceBridgeInterface) (*ScriptContainerVM, error) {
	// Das Basis Objekt wird erezgut
	base_bundle_runtime := &ScriptContainerVM{new(sync.WaitGroup), fsContainer, imageFile, nil, coreController, make(map[string]map[string]SharedFunctionInterface, 0), false}

	// Die Javascript Runtime wird erstellt
	engine, err := jsengine.NewEngine(base_bundle_runtime)
	if err != nil {
		return nil, fmt.Errorf("NewRuntime: " + err.Error())
	}

	// Die Engines werden geschrieben
	if err := base_bundle_runtime.setEngines(engine); err != nil {
		return nil, fmt.Errorf("NewRuntime: " + err.Error())
	}

	// Das Objekt wird zurückgegeben
	return base_bundle_runtime, nil
}
