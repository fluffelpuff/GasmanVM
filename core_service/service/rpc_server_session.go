package coreservice

import (
	"fmt"

	"github.com/fluffelpuff/GasmanVM/core_service/logging"
	"github.com/fluffelpuff/GasmanVM/core_service/utils"
)

func (o *CoreService) registerVMProcess(vmProcess *ProcessSession) error {
	// Der Prozess wird Registriert
	o.openProcesses[vmProcessPointerId(vmProcess)] = vmProcess

	// Debug
	logging.LogInfo(fmt.Sprintf("New package '%s' on process '%s' registrated", vmProcess.ManifestData.Package.GasmanVmPackage, vmProcess.processId))

	// Es ist kein Fehler aufgetreten
	return nil
}

func (o *CoreService) unregisterVMProcess(vmProcess *ProcessSession) {
	// Die Geteilten Funktionen werden gespeichert
	delete(o.functionalSharingPcoesses, vmProcessPointerId(vmProcess))

	// Der Hauptprozess wird entfertn
	delete(o.openProcesses, vmProcessPointerId(vmProcess))

	// Debug
	logging.LogInfo(fmt.Sprintf("Process '%s' unregistered", vmProcess.processId))
}

func (o *CoreService) enableFunctionShare(vmProcess *ProcessSession) error {
	// Es wird ermittelt ob der Prozess Registriert
	if _, found := o.openProcesses[vmProcessPointerId(vmProcess)]; !found {
		return fmt.Errorf("unkown process")
	}

	// Der Prozess wird im Sharing Funktion Storage gespeichert
	o.functionalSharingPcoesses[vmProcessPointerId(vmProcess)] = vmProcess

	// Debug
	logging.LogInfo(fmt.Sprintf("Functional sharing is enabeld on process '%s'", vmProcess.processId))

	// Es ist kein Fehler aufgetreten
	return nil
}

func (o *CoreService) shareFunction(groupName string, functionName string, endPoint *ProcessSession) (string, error) {
	// Es wird ermittelt ob es sich um eine bekannte Gruppe handelt
	groupItem := o.sharedFunctionMap.GetGroup(groupName)
	if groupItem == nil {
		return "", fmt.Errorf("unkown sharing group")
	}

	// Die Daten werden zurückgegeben
	id, err := o.sharedFunctionMap.AddProcessFunctionShare(functionName, groupItem, endPoint)
	if err != nil {
		return "", err
	}

	// LOG
	logging.LogInfo(fmt.Sprintf("Process '%s' shared function '%s' in group '%s'", endPoint.processId, functionName, groupName))

	// Die ID wird zurückgegeben
	return id, nil
}

func (o *CoreService) registerProcessInSharingGroupAndReturnGroup(groupID string, sourceCallerSession *ProcessSession) (*utils.SharingGroup, error) {
	// Es wird versucht die Gruppe abzurufen
	item := o.sharedFunctionMap.GetGroup(groupID)
	if item == nil {
		return nil, fmt.Errorf("service: unkown group")
	}

	// Es wird versucht den Prozess sowie die Gruppe in der Shared Function Map zuzuordnen
	err := o.sharedFunctionMap.RegisterProcessWithGroup(item, sourceCallerSession)
	if err != nil {
		return nil, err
	}

	// Debug
	logging.LogInfo(fmt.Sprintf("Process '%s' was registered in sharing group '%s'", sourceCallerSession.processId, item.Name))

	// Der Prozess sowie die Gruppe werden zusammen
	return item, nil
}

func (o *CoreService) callFunction(groupName string, functionName string, sourceCallerSession *ProcessSession, args []interface{}) (bool, interface{}, error) {
	// Es wird versucht die Gruppe abzurufen
	item := o.sharedFunctionMap.GetGroup(groupName)
	if item == nil {
		return false, nil, fmt.Errorf("unkown group")
	}

	// LOG
	logging.LogInfo(fmt.Sprintf("Process '%s' tries to call function '%s' in group '%s'", sourceCallerSession.processId, functionName, groupName))

	// Die Funktion wird aufgerufen
	found, result, funcError, internalError := o.sharedFunctionMap.CallFunction(functionName, item, args)
	if internalError != nil {
		return found, result, fmt.Errorf("internal error")
	}

	// Es wird geprüft ob die Funktion gefunden wurde
	if found {
		logging.LogInfo(fmt.Sprintf("Process '%s' tried to call function '%s' in group '%s', successfully", sourceCallerSession.processId, functionName, groupName))
	} else {
		logging.LogInfo(fmt.Sprintf("Process '%s' tried to call function '%s' in group '%s', but the function could not be called.", sourceCallerSession.processId, functionName, groupName))
	}

	// Die Daten werden zurückgegeben
	return found, result, funcError
}
