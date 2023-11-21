package coreclientbridge

import (
	"fmt"

	"github.com/fluffelpuff/GasmanVM/core_service/argtypes"
)

type ProcessServerSession struct {
	processRegisterCompleted bool
	processSeecret           string
	groups                   map[string]string
}

func (t *ProcessServerSession) ProcessRegisterComplete(args *argtypes.RegisterProcessArgsCompleteArgs, reply *argtypes.EmptyReturn) error {
	*reply = argtypes.EmptyReturn{}
	t.processRegisterCompleted = true
	t.processSeecret = args.ProcessSecret
	// Es ist kein Fehler aufgetreten
	return nil
}

func (t *ProcessServerSession) ProcessGroupedRegisterComplete(args *argtypes.RegisterGroupProcessArgsCompleteArgs, reply *argtypes.EmptyReturn) error {
	*reply = argtypes.EmptyReturn{}
	t.processRegisterCompleted = true
	t.processSeecret = args.ProcessSecret
	t.groups = args.Groups
	fmt.Println(args.Groups)
	// Es ist kein Fehler aufgetreten
	return nil
}

func (t *ProcessServerSession) processRegistrationCompleted() bool {
	return t.processRegisterCompleted
}

func (t *ProcessServerSession) getGroups() map[string]string {
	return t.groups
}

func newVMProcessServerSession(ccb *CoreClientBridge) *ProcessServerSession {
	return &ProcessServerSession{groups: make(map[string]string)}
}
