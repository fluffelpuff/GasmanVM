package coreclientbridge

import (
	"fmt"
	"net/rpc/jsonrpc"
	"sync"

	"github.com/Microsoft/go-winio"
	"github.com/fluffelpuff/GasmanVM/core_service/commchannel"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func OpenBridgeConnection() (*CoreClientBridge, error) {
	// DEBUG
	fmt.Print("Try to connect CoreService: ")

	// Öffnen der Named Pipe-Verbindung
	conn, err := winio.DialPipe(vmpackage.CORE_CONTROLLER_WIN32_PIP_PATH, nil)
	if err != nil {
		return nil, err
	}

	// Die Verbindung wird geupgraded
	upgradedConnection := commchannel.WrappCommConnection(conn)

	// Der ClientRPC Channel wird erstellt
	clientRPC, err := upgradedConnection.OpenSubConnChannel(0)
	if err != nil {
		return nil, err
	}

	// Der ServerRPC Channel wird erstellt
	serverRPCChannel, err := upgradedConnection.OpenSubConnChannel(1)
	if err != nil {
		return nil, err
	}

	// ProcessSession-Client erstellen
	client := jsonrpc.NewClient(clientRPC)

	// Es wird eine neue CoreClient Bridge erstellt
	coreClientBridge := &CoreClientBridge{conn, client, serverRPCChannel, nil, nil, new(sync.WaitGroup), ""}

	// DEBUG
	fmt.Println("DONE!")

	// Die Aktuelle Verbindung wird zurpückgegeben
	return coreClientBridge, nil
}
