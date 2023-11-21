package coreclientbridge

import (
	"fmt"

	"github.com/fluffelpuff/GasmanVM/core_service/argtypes"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
)

func (o *CoreClientBridge) SetVM(vm vmpackage.VMInterface) error {
	if o.vm != nil {
		return fmt.Errorf("vm always seted")
	}
	o.vm = vm
	fmt.Println("VM SETED")
	return nil
}

func (o *CoreClientBridge) RegisterSharedFunction(groupName string, functionName string, vm vmpackage.VMInterface) (*argtypes.RegisterSharedFunctionReturn, error) {
	// Die Kernparameter werden zusammengefasst
	request := &argtypes.RegisterSharedFunctionArgs{
		GroupName:    groupName,
		FunctionName: functionName,
	}

	// ProcessSession-Aufruf
	var reply *argtypes.RegisterSharedFunctionReturn
	err := o.rpcClinet.Call("ProcessSession.RegisterSharedFunction", request, &reply)
	if err != nil {
		return nil, err
	}

	// Die Daten werden zurückgegeben
	return reply, nil
}

func (o *CoreClientBridge) CallSharedFunction(packageIdentifyer vmpackage.PackageIdentifyerInterface, groupName string, functionName string, timeout uint64, args []interface{}) (interface{}, error) {
	// Die Kernparameter werden zusammengefasst
	request := &argtypes.CallSharedFunctionArgs{
		FunctionId:     packageIdentifyer.GetID(),
		FunctionIdType: packageIdentifyer.GetType(),
		Timeout:        timeout,
		GroupName:      groupName,
		FunctionName:   functionName,
		Args:           args,
	}

	// ProcessSession-Aufruf
	var reply *argtypes.CallSharedFunctionReturn
	err := o.rpcClinet.Call("ProcessSession.CallSharedFunction", request, &reply)
	if err != nil {
		return nil, err
	}

	// Die Daten werden zurückgegeben
	return reply, nil
}
