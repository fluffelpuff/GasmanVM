package sharedfunctions

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/utils"
)

// Call ruft eine Goja-Funktion auf und verarbeitet die Rückgabe.
// Die Funktion akzeptiert eine variable Anzahl von Daten, die für den Funktionsaufruf vorbereitet werden.
// Die Funktion gibt ein Interface und einen Fehler zurück.
// - `data ...interface{}`: Die Eingabedaten für den Funktionsaufruf.
// - `interface{}`: Das Ergebnis des Funktionsaufrufs, wenn es erfolgreich ist.
// - `error`: Ein Fehler, der auftritt, wenn die Rückgabe ungültig ist oder ein anderer Fehler auftritt.
func (o *LocalSharedFunctionCapsle) Call(data ...interface{}) (interface{}, error) {
	// Die Daten werden für Goja vorbereitet.
	gojaParms := make([]goja.Value, 0)
	for _, item := range data {
		gojaParms = append(gojaParms, o.JsRuntime.ToValue(item))
	}

	// Die Parameter für den Funktionsaufruf werden erstellt.
	parms := goja.FunctionCall{This: o.JsRuntime.GlobalObject(), Arguments: gojaParms}

	// Die Funktion wird aufgerufen.
	functionCallResult := o.JsCall(parms)

	// Die Rückgabewerte werden exportiert
	exportedResult := functionCallResult.Export()

	// Die Rückgabe wird validiert.
	if !utils.CheckDataValues(exportedResult) {
		return nil, fmt.Errorf("Ungültiger Rückgabewert")
	}

	// Die Rückgabewerte werden zurückgegeben.
	return exportedResult, nil
}

func (*LocalSharedFunctionCapsle) ClientFunctionCreator() uint64 {
	return 0
}

func (o *LocalSharedFunctionCapsle) IsLocal() bool {
	return true
}
