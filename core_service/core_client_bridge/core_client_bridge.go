package coreclientbridge

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	"github.com/fluffelpuff/GasmanVM/core_service/argtypes"
	"github.com/fluffelpuff/GasmanVM/core_service/commchannel"
	"github.com/fluffelpuff/GasmanVM/imagefile"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

type CoreClientBridge struct {
	conn             net.Conn
	rpcClinet        *rpc.Client
	rpcServerChannel *commchannel.ConnChannel
	rpcServer        *rpc.Server
	vm               vmpackage.VMInterface
	waitGroup        *sync.WaitGroup
	processSecret    string
}

// Setup ist eine Methode des CoreClientBridge-Typs, die dazu verwendet wird, den Prozess zu konfigurieren und zu starten.
// Sie nimmt ein Manifest-Objekt als Eingabe, registriert den Prozess beim Remote-Service "ProcessSession.RegisterVMProcess"
// und startet einen RPC-Server, um Prozessanfragen zu behandeln.
// Diese Methode gibt einen Fehler zurück, wenn bei der Konfiguration oder Kommunikation Probleme auftreten.
// Ansonsten gibt sie `nil` zurück, um anzuzeigen, dass die Konfiguration und der Start erfolgreich abgeschlossen wurden.
func (o *CoreClientBridge) Setup(manifest *imagefile.Manifest) error {
	// DEBUG
	fmt.Print("Setup the process: ")

	// Die Kernparameter werden zusammengefasst
	request := &argtypes.RegisterProcessArgs{
		ManifestData: manifest,
		Version:      vmpackage.VERSION,
	}

	// Es wird ein neuer RPC Server erstellt
	o.rpcServer = rpc.NewServer()

	// Das SubServerModul wird erstellt und registriert
	sessionRPC := newVMProcessServerSession(o)
	o.rpcServer.Register(sessionRPC)

	// Der Server wird gestartet
	o.waitGroup.Add(1)
	go func() {
		// Führt den Server aus
		o.rpcServer.ServeCodec(jsonrpc.NewServerCodec(o.rpcServerChannel))

		// Signalisiert dass die Routine beendet wurde
		o.waitGroup.Done()
	}()

	// ProcessSession-Aufruf
	var reply argtypes.RegisterVMProcessReturn
	if err := o.rpcClinet.Call("ProcessSession.RegisterVMProcess", request, &reply); err != nil {
		return err
	}

	// Es wird geprüft ob der Process
	if !sessionRPC.processRegistrationCompleted() {
		return fmt.Errorf("invalid process registration process")
	}

	// Es wird ermittelt ob die Anwendung Gruppen verwendet
	if manifest.HasSharingGroups() {
		// Es wird geprüft ob die benötigten Gruppen ermittelt werden konnten
		notDeterminedGroups := []string{}
		for _, item := range manifest.GetSharingGroups() {
			_, foundIt := sessionRPC.getGroups()[item]
			if !foundIt {
				notDeterminedGroups = append(notDeterminedGroups, item)
				continue
			}
		}

		// Sollte es Gruppen geben welche nicht Ermittelt werden konnten wird das Programm abgeschlossen
		if len(notDeterminedGroups) > 0 {
			fmt.Println(sessionRPC.getGroups())
			return fmt.Errorf("unkown groups: %s", formatStringList(notDeterminedGroups))
		}
	}

	// Die ID wird zwischengespeichert
	o.processSecret = reply.ProcessSecret

	// DEBUG
	fmt.Println("DONE!")

	// Der Vorgang wurde ohne Fehler durchgeführt
	return nil
}

// Provide ist eine Methode des CoreClientBridge-Typs, die dazu dient, ein Bild bereitzustellen.
// Sie ruft den Remote-Service "ProcessSession.ProvideProcess" auf, um die Bildbereitstellung durchzuführen.
// Diese Methode gibt einen Fehler zurück, falls bei der Kommunikation mit dem Remote-Service Probleme auftreten.
// Ansonsten gibt sie `nil` zurück, um anzuzeigen, dass die Bereitstellung erfolgreich abgeschlossen wurde.
func (o *CoreClientBridge) Provide() error {
	// DEBUG
	fmt.Print("Provide the image: ")

	// Die Kernparameter werden zusammengefasst
	request := &argtypes.ProvideProcessArgs{ProcessSecret: o.processSecret}

	// ProcessSession-Aufruf
	var reply argtypes.ProvideProcessReturn
	if err := o.rpcClinet.Call("ProcessSession.ProvideProcess", request, &reply); err != nil {
		return err
	}

	// DEBUG
	fmt.Println("DONE!")

	// Es ist kein Fehler aufgetreten
	return nil
}
