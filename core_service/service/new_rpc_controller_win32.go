package coreservice

import (
	"github.com/Microsoft/go-winio"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func NewCoreServiceController() (*MemoryController, error) {
	// Die Windows Pipe wird geöffnet
	listener, err := winio.ListenPipe(vmpackage.CORE_CONTROLLER_WIN32_PIP_PATH, nil)
	if err != nil {
		return nil, err
	}

	// Das Objekt wird erstellt
	result := &MemoryController{listener}

	// Das Objekt wird zurückgegeben
	return result, nil
}
