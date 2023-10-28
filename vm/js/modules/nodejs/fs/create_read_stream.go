package fs

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

func Module_FS_SYNC_createReadStream(vmengine modules.VMInterface, jsrumtime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 1 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "createReadStream", 1)))
	}

	// Die Argumente werden extrahiert
	fullQualPath := parms.Arguments[0].String()
	option := parms.Arguments[1].String()

	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		panic("internal filesystem error, unkown error")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	fileResolve, err := fileSystem.GetFileByFullPath(fullQualPath)
	if err != nil {
		return nil
	}
	if fileResolve == nil {
		panic("file not found")
	}

	// Der Stream wird geöffnet
	_ = option

	return nil
}
