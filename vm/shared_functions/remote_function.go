package sharedfunctions

// Call ruft eine Goja-Funktion auf und verarbeitet die Rückgabe.
// Die Funktion akzeptiert eine variable Anzahl von Daten, die für den Funktionsaufruf vorbereitet werden.
// Die Funktion gibt ein Interface und einen Fehler zurück.
// - `data ...interface{}`: Die Eingabedaten für den Funktionsaufruf.
// - `interface{}`: Das Ergebnis des Funktionsaufrufs, wenn es erfolgreich ist.
// - `error`: Ein Fehler, der auftritt, wenn die Rückgabe ungültig ist oder ein anderer Fehler auftritt.
func (o *RemoteSharedFunctionCaplse) Call(data ...interface{}) (interface{}, error) {
	// Die Rückgabewerte werden zurückgegeben.
	return nil, nil
}

func (*RemoteSharedFunctionCaplse) ClientFunctionCreator() uint64 {
	return 0
}

func (o *RemoteSharedFunctionCaplse) IsLocal() bool {
	return true
}
