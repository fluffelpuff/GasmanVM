package main

import (
	"fmt"

	"github.com/fluffelpuff/GasmanVM/imagefile"
	jsengine "github.com/fluffelpuff/GasmanVM/vm/js"
)

func (o *ScriptContainerVM) Run() error {
	// Die Main Datei wird gestartet
	_, err := o.jsEngine.RunMainScript()
	if err != nil {
		return fmt.Errorf("ScriptContainerVM>Run" + err.Error())
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

func (o *ScriptContainerVM) GetImageFile() *imagefile.ImageFile {
	return o.imageFile
}

func NewRuntime(imageFile *imagefile.ImageFile) (*ScriptContainerVM, error) {
	// Das Basis Objekt wird erezgut
	base_bundle_runtime := &ScriptContainerVM{imageFile, nil}

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
