package main

import coreservice "github.com/fluffelpuff/GasmanVM/core_service/service"

func main() {
	// Der Service Controller wird gestartet
	socketController, err := coreservice.NewCoreServiceController()
	if err != nil {
		panic(err)
	}

	// Es wird ein neuer RPC Controller erstellt
	rpcController, err := coreservice.NewRPCController(socketController)
	if err != nil {
		panic(err)
	}

}
