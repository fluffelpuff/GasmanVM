package coreservice

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	"github.com/fluffelpuff/GasmanVM/core_service/argtypes"
	"github.com/fluffelpuff/GasmanVM/core_service/commchannel"
	"github.com/fluffelpuff/GasmanVM/imagefile"
)

type ProcessSession struct {
	coreService         *CoreService
	ManifestData        *imagefile.Manifest
	commRPCClientServer *commchannel.ConnChannel
	rpcClient           *rpc.Client
	mu                  sync.Mutex
	processId           string
	processRegistered   bool
	processEnabled      bool
}

func (t *ProcessSession) RegisterVMProcess(args *argtypes.RegisterProcessArgs, reply *argtypes.RegisterVMProcessReturn) error {
	// Der Mutex wird verwendet
	t.mu.Lock()
	defer t.mu.Unlock()

	// Es wird geprüft ob der Prozess bereits Registriert wurde
	if t.processRegistered {
		return fmt.Errorf("process always registrated")
	}

	// Es wird ein neuer RPC Client erstellt
	t.rpcClient = jsonrpc.NewClient(t.commRPCClientServer)

	// Die Manifestdaten werden zwischengespeichert
	t.ManifestData = args.ManifestData

	// Es wird Signalisiert dass die VM Einsatzbereit ist
	if err := t.coreService.registerVMProcess(t); err != nil {
		return fmt.Errorf("invalid internal process")
	}

	// Es wird ermittel ob die Gruppen Welche das Paket benötigt bekannt sind
	if args.ManifestData.HasSharingGroups() {
		// Speichert die Ermittelten Gruppen ab
		retrivedGroups := map[string]string{}

		// Die Gruppen werden ermittelt
		for _, item := range args.ManifestData.GetSharingGroups() {
			// Es wird geprüft ob der Host Mitglied in dieser Gruppe ist
			if !t.coreService.sharedFunctionMap.IsKnownSharingGroup(item) {
				continue
			}

			// Der Prozess wird in der Gruppe Registriert
			sharingGroup, err := t.coreService.registerProcessInSharingGroupAndReturnGroup(item, t)
			if err != nil {
				return err
			}

			// Der Gruppenname wird Zwischengespeichert
			// der Gruppenname wird der GruppenID zugeordnet
			retrivedGroups[item] = sharingGroup.GetGroupName()
		}

		// Die Daten werden zusammengefasst
		args := argtypes.RegisterGroupProcessArgsCompleteArgs{
			Groups:        retrivedGroups,
			ProcessSecret: t.processId,
		}

		// Es wird dem Client Signalisiert dass der Vorgang erfolgreich beendet wurde
		if err := t.rpcClient.Call("ProcessServerSession.ProcessGroupedRegisterComplete", &args, &argtypes.EmptyReturn{}); err != nil {
			return fmt.Errorf("process registration failed")
		}
	} else {
		// Die Daten werden zusammengefasst
		args := argtypes.RegisterProcessArgsCompleteArgs{ProcessSecret: t.processId}

		// Es wird dem Client Signalisiert dass der Vorgang erfolgreich beendet wurde
		if err := t.rpcClient.Call("ProcessServerSession.ProcessRegisterComplete", &args, &argtypes.EmptyReturn{}); err != nil {
			return fmt.Errorf("process registration failed")
		}
	}

	// Die Daten werden übermittelt
	*reply = argtypes.RegisterVMProcessReturn{ProcessSecret: t.processId}

	// Es wird Signalisiert dass der Prozess Registriert wurde
	t.processRegistered = true

	// Es ist kein Fehler aufgetreten
	return nil
}

func (t *ProcessSession) ProvideProcess(args *argtypes.ProvideProcessArgs, reply *argtypes.ProvideProcessReturn) error {
	// Der Mutex wird verwendet
	t.mu.Lock()
	defer t.mu.Unlock()

	// Es wird geprüft ob der Prozess bereits Registriert wurde
	if !t.processRegistered {
		return fmt.Errorf("process isnt registrated")
	}

	// Es wird geprüft ob das ProcessSecret korrekt ist
	if t.processId != args.ProcessSecret {
		return fmt.Errorf("invalid process secret")
	}

	// Die Shared Group Functions werden Aktiviert
	if err := t.coreService.enableFunctionShare(t); err != nil {
		return fmt.Errorf("process providing error")
	}

	// Die Antwort wird zurückgesendet
	*reply = argtypes.ProvideProcessReturn{}

	// Der Prozess wird als Aktiv Markiert
	t.processEnabled = true

	// Die Daten werden zurückgegeben
	return nil
}

func (t *ProcessSession) CallSharedFunction(args *argtypes.CallSharedFunctionArgs, reply *argtypes.CallSharedFunctionReturn) error {
	// Die Funktion wird ausgeführt
	foundFunction, result, err := t.coreService.callFunction(args.GroupName, args.FunctionName, t, args.Args)
	if err != nil {
		return err
	}

	// Es wird geprüft ob die Funktion gefunden wurde
	if !foundFunction {
		return fmt.Errorf("function not found")
	}

	*reply = argtypes.CallSharedFunctionReturn{
		Result: result,
	}

	// Die Daten werden zurückgegeben
	return nil
}

func (t *ProcessSession) RegisterSharedFunction(args *argtypes.RegisterSharedFunctionArgs, reply *argtypes.RegisterSharedFunctionReturn) error {
	// Der Mutex wird verwendet
	t.mu.Lock()
	defer t.mu.Unlock()

	// Die Funktion wird global Registriert
	procFuncId, err := t.coreService.shareFunction(args.GroupName, args.FunctionName, t)
	if err != nil {
		return err
	}

	// Die Antwort wird zurückgesendet
	*reply = argtypes.RegisterSharedFunctionReturn{FunctionId: procFuncId}

	// Es ist kein Fehler aufgetreten
	return nil
}

func (t *ProcessSession) killConnectionClosed() {
	// Der Mutex wird verwendet
	t.mu.Lock()
	defer t.mu.Unlock()

	// Der RPC Client wird geschlossen
	t.rpcClient.Close()

	// Der Prozess wird Entfernt
	t.coreService.unregisterVMProcess(t)
}

func newRPCSession(procID string, coreSerice *CoreService, clientServerRPCChannel *commchannel.ConnChannel) *ProcessSession {
	// Das ProcessSession Sitzungsobjekt wird erstellt
	rpcSession := ProcessSession{}
	rpcSession.coreService = coreSerice
	rpcSession.processId = procID
	rpcSession.processEnabled = false
	rpcSession.processRegistered = false
	rpcSession.ManifestData = nil
	rpcSession.commRPCClientServer = clientServerRPCChannel
	rpcSession.rpcClient = nil
	rpcSession.mu = sync.Mutex{}
	return &rpcSession
}
