package fs

import (
	"fmt"
	"strings"

	"github.com/dop251/goja"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules"
)

// Diese Funktion wird verwendet um eine Datei im allgemeinen zu Laden und einen Goja Wert zurückzugeben
func fsReadFileCore(vmengine modules.VMInterface, jsruntime *goja.Runtime, filePath string, encoding string) (goja.Value, error) {
	// Das Dateisystem wird abgerufen
	fileSystem := vmengine.GetFilesystem()
	if fileSystem == nil {
		return nil, fmt.Errorf("internal filesystem error, unkown error")
	}

	// Es wird ermittelt ob die Datei verfügbar ist
	fileResolve, err := fileSystem.GetFileByFullPath(filePath)
	if err != nil {
		return nil, err
	}
	if fileResolve == nil {
		return nil, fmt.Errorf("file not found")
	}

	// Die Option wird geprüft
	switch encoding {
	case "utf8":
		// Der Inhalt der Datei wird als UTF8 String abgerufen
		utf8Value, err := fileResolve.OpenUtf8()
		if err != nil {
			return nil, err
		}

		// Gibt die Daten zurück
		return jsruntime.ToValue(utf8Value), nil
	case "binary", "buffer":
		// Der Inhalt der Datei wird als Byte Array eingelesen
		byteArrayValue, err := fileResolve.OpenBinary()
		if err != nil {
			return nil, err
		}

		// Es wird ein GoJa Byte Array aus den Bytes erstellt
		gojaByteBuffer := jsruntime.NewArrayBuffer(byteArrayValue)

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(gojaByteBuffer), nil
	case "base64":
		// Der Inhalt der Datei wird als Base64 String abgerufen
		base64Value, err := fileResolve.OpenBase64()
		if err != nil {
			return nil, err
		}

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(base64Value), nil
	case "base32":
		// Der Inhalt der Datei wird als Base32 String abgerufen
		base32Value, err := fileResolve.OpenBase32()
		if err != nil {
			return nil, err
		}

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(base32Value), nil
	case "base58":
		// Der Inhalt der Datei wird als Base58 String abgerufen
		base58Value, err := fileResolve.OpenBase58()
		if err != nil {
			return nil, err
		}

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(base58Value), nil
	case "hex":
		// Der Inhalt der Datei wird als Hex String abgerufen
		hexValue, err := fileResolve.OpenHex()
		if err != nil {
			return nil, err
		}

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(hexValue), nil
	case "latin1":
		// Der Inhalt der Datei wird als latin1 String abgerufen
		latin1Value, err := fileResolve.OpenLatin1()
		if err != nil {
			// Der Fehler wird zurückgegeben
			return nil, err
		}

		// Die Daten werden zurückgegeben
		return jsruntime.ToValue(latin1Value), nil
	default:
		// Es wird ein Fehler zurückgegeben
		return nil, fmt.Errorf("unkown encoding")
	}
}

// Diese Funktion wird für das Synchrone Laden einer Datei verwendet
func Module_FS_SYNC_readFileSync(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readFileSync", 2)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird versucht die Datei einzulesen
	result, err := fsReadFileCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String())
	if err != nil {
		panic(goja.New().NewGoError(err))
	}

	// Die Daten werden zurückgegeben
	return result
}

// Diese Funktion wird für das Synchrone Laden einer Datei verwendet
func Module_FS_SYNC_readFileCallback(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 3 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readFile", 3)))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[0].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich um einen String handelt
	if parms.Arguments[1].ExportType().String() != "string" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a string")))
	}

	// Es wird ermittelt ob es sich bei der Dritten Variable um eine Funktion handelt
	if parms.Arguments[2].ExportType().String() != "func(goja.FunctionCall) goja.Value" {
		panic(goja.New().NewGoError(fmt.Errorf("the function need a callback function")))
	}

	// Die Callback Funktion wird eingelesen
	callbackFunction, isOk := parms.Arguments[2].Export().(func(goja.FunctionCall) goja.Value)
	if !isOk {
		panic(fmt.Errorf("invalid callback function"))
	}

	// Es wird versucht die Datei einzulesen
	result, err := fsReadFileCore(vmengine, jsruntime, parms.Arguments[0].String(), parms.Arguments[1].String())
	var callbackParms goja.FunctionCall
	if err != nil {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{jsruntime.ToValue(err)}}
	} else {
		callbackParms = goja.FunctionCall{This: jsruntime.GlobalObject(), Arguments: []goja.Value{goja.Null(), result}}
	}

	// Die Callback Funktion wird aufgerufen
	callbackFunction(callbackParms)

	// Die Daten werden zurückgegeben
	return goja.Null()
}

// Diese Funktion wird für das Asynchrone Laden einer Datei verwendet
func Module_FS_ASYNC_readFilePromises(vmengine modules.VMInterface, jsruntime *goja.Runtime, parms goja.FunctionCall) goja.Value {
	// Es wird geprüft ob die Benötigte Anzahl von Parametern vorhanden ist
	if len(parms.Arguments) != 2 {
		panic(goja.New().NewGoError(fmt.Errorf("the function '%s' requires %d parameters", "readFile", 2)))
	}

	// Die Argumente werden extrahiert
	filePath := parms.Arguments[0].String()
	fileOption := strings.ToLower(strings.TrimSpace(parms.Arguments[1].String()))

	// Erstellen Sie eine neue Promise und die Resolving-Funktionen
	promise, resolve, reject := jsruntime.NewPromise()

	// Es wird ein neuer Lesender Vorgang Registriert
	vmengine.AddNewRoutine()

	// Die Datei wird Asynchrone eingelesen
	go func() {
		// Wird ausgeführt wenn die Funktion fertig ist
		defer func() {
			// Es wird Signalisiert dass die Unit beendet wurde
			vmengine.RemoveRoutine()
		}()

		// Die Datei wird eingelesen
		result, err := fsReadFileCore(vmengine, jsruntime, filePath, fileOption)
		if err != nil {
			// Der Fehlelr wird zurückgegeben
			reject(goja.New().NewGoError(err))
			return
		}

		// Das Ergebniss wird zurückgegeben
		resolve(result)
	}()

	// Rückgabe der Promise
	return jsruntime.ToValue(promise)
}
