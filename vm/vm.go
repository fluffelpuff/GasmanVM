package vm

import (
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/fsysfile"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
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
	// Die Funktion wird im CoreController Registriert
	registrationResult, err := o.coreServiceBridgeInterface.RegisterSharedFunction(groupName, shareName, o)
	if err != nil {
		return err
	}

	// Die ID unter welcher die Funktion aufzurufen ist, wird abgespeichert
	o.functionSharingIdMap[registrationResult.FunctionId] = sharedFunction

	// DEBUG
	fmt.Printf("New local function share registrated '%s/%s'\n", groupName, shareName)

	// Es ist kein Fehler aufgetreten
	return nil
}

func (o *ScriptContainerVM) RegisterSharedFunction(groupName string, functionName string, functionCallId string, coreBridge vmpackage.CoreServiceBridgeInterface) error {
	// DEBUG
	fmt.Println(fmt.Printf("New remote function share registrated '%s/%s'", groupName, functionName))

	// Es ist kein Fehler aufgetreten
	return nil
}

func (o *ScriptContainerVM) CallSharedFunction(packageIdentifyer vmpackage.PackageIdentifyerInterface, groupName string, functionName string, timeout uint64, parms ...interface{}) (interface{}, error) {
	// Die Funktion wird aufgerufen
	functionCallresult, err := o.coreServiceBridgeInterface.CallSharedFunction(packageIdentifyer, groupName, functionName, timeout, parms)
	if err != nil {
		return nil, err
	}

	// DEBUG
	fmt.Printf("Call shared function '%s/%s' with timeout '%d ms'\n", groupName, functionName, timeout)

	// Das Ergebniss wird zurückgegeben
	return functionCallresult, nil
}

func NewRuntime(fsContainer *fsysfile.FileSystem, imageFile *imagefile.ImageFileReader, coreController vmpackage.CoreServiceBridgeInterface) (*ScriptContainerVM, error) {
	// Das Basis Objekt wird erezgut
	base_bundle_runtime := &ScriptContainerVM{new(sync.WaitGroup), fsContainer, imageFile, nil, coreController, false, make(map[string]func(goja.FunctionCall) goja.Value)}

	// Die VM wird in dem CoreController Registriert
	if err := coreController.SetVM(base_bundle_runtime); err != nil {
		return nil, err
	}

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
