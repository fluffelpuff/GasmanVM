package coreservice

import (
	"encoding/hex"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	"github.com/fluffelpuff/GasmanVM/core_service/commchannel"
)

// Stellt den ProcessSession Controller dar
type RPCController struct {
	coreService   *CoreService
	memController *MemoryController
	waitGroup     *sync.WaitGroup
}

// Upgraded eine Verbindung, stellt den ProcessSession Server sowie den ProcessSession CLient bereit und fast dass ganze in einer ProcessSession zusammen
func (o *RPCController) upgradeToTPCAndServe(conn net.Conn) {
	// Die Verbindung wird geschlossen wenn die Funktion endet
	defer conn.Close()

	// Zufälligen Wert generieren
	randomValue, err := generateRandomValue()
	if err != nil {
		return
	}

	// Hash des zufälligen Werts erstellen
	hashValue := hashRandomValue(randomValue)

	// Der Hashwert wird in einen Hexstring umgewandelt
	hexedHashValue := hex.EncodeToString(hashValue)

	// Die Verbindug wird Wrappd
	wrappedConnection := commchannel.WrappCommConnection(conn)

	// Es wird ein CommChannel für die Serverseite RPC Funktion erstellt
	coreServiceRPCServerChannel, err := wrappedConnection.OpenSubConnChannel(0)
	if err != nil {
		panic(err)
	}

	// Es wird ein Commchannel für die CLientseitige Server RPC Funktion erstellt
	coreServiceRPCClientChannel, err := wrappedConnection.OpenSubConnChannel(1)
	if err != nil {
		panic(err)
	}

	// Es wird eine neue ROC Sitzung erstellt
	rpcSession := newRPCSession(hexedHashValue, o.coreService, coreServiceRPCClientChannel)

	// Das Closerevent wird festgelegt
	wrappedConnection.AddEventByClose(func(msg string) {
		rpcSession.killConnectionClosed()
	})

	// Der ProcessSession Server wird erzeugt
	server := rpc.NewServer()
	server.Register(rpcSession)

	// ProcessSession-Verbindung einrichten
	server.ServeCodec(jsonrpc.NewServerCodec(coreServiceRPCServerChannel))
}

// Wird im hintergrund ausgeführt, kümmert sich um neue Prozesse (Verbindungen)
func (o *RPCController) backgroundProcess() {
	// Die Schleife wird solange ausgeführt, bis ein CLoser kommt
	for {
		// Es wird auf Eintreffende Verbindungen gewartet
		conn, err := o.memController.AcceptConnection()
		if err != nil {
			panic(err)
		}

		// Es wird Signalisiert dass eine weitere Routine ausgeführt wird
		o.waitGroup.Add(1)

		// Die Verbindung wird zu einer ProcessSession Verbindung geupgraded
		go o.upgradeToTPCAndServe(conn)
	}
}

// Wird in der Main Funktion ausgeführt und hält den ProcessSession Controller am leben
func (o *RPCController) MainRun() error {
	// Es wird ein neuer Main Thread Registriert
	o.waitGroup.Add(1)

	// Der BackgroundProzess wird gestartet
	go o.backgroundProcess()

	// Es wird gewartet bis alle aufgaben fertig sind
	o.waitGroup.Wait()

	// Es ist kein Fehler aufgetreten
	return nil
}

// Wird verwendet um den ProcessSession Controller zu schließen
func (o *RPCController) Close() error {
	return nil
}

// Erstellt einen neuen ProcessSession Controller
func NewRPCController(memController *MemoryController, coreService *CoreService) (*RPCController, error) {
	controller := &RPCController{coreService, memController, &sync.WaitGroup{}}
	if err := coreService.registerRPCController(controller); err != nil {
		return nil, err
	}
	return controller, nil
}
