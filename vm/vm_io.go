package vm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// Gibt an ob es sich um ein Objekt handelt
func isMap(input interface{}) bool {
	// Verwende reflect.TypeOf, um den Typ des übergebenen Interface zu erhalten
	t := reflect.TypeOf(input)

	// Überprüfe, ob es sich um einen Map-Typ handelt
	return t.Kind() == reflect.Map
}

// Extrahiert die Print Werte aus den Goja Argumenten
func exportStringArrayFromArgumentsAndFromateIt(arguments []interface{}) []interface{} {
	// Extrahieren Sie die Argumente, die an die print-Funktion übergeben wurden
	args := make([]interface{}, len(arguments))
	for i, arg := range arguments {
		// Es wird geprüft ob es sich um eien Map handelt, wenn ja wird diese in JSON Umgewandelt und ausgegeben
		if isMap(arg) {
			converted, err := customJSONMarshal(arg)
			if err != nil {
				panic(err)
			}
			args[i] = converted
			continue
		} else {
			args[i] = arg
			continue
		}
	}
	return args
}

// Wandelt ein Objekt in eine Formatierte ausgabe um
func customJSONMarshal(v interface{}) (string, error) {
	// Benutzerdefinierte JSON-Umwandlung mit Einrückungen
	var sb strings.Builder
	encoder := json.NewEncoder(&sb)
	encoder.SetEscapeHTML(false) // Escaping deaktivieren
	encoder.SetIndent("", "  ")  // Einrückungen mit zwei Leerzeichen
	err := encoder.Encode(v)
	if err != nil {
		return "", err
	}
	return string([]byte(sb.String())), nil
}

// Gibt einen JA/Nein Fragedialog aus
func YesOrNoTextEnter(banner string) bool {
	for {
		fmt.Print(banner)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.ToLower(strings.TrimSpace(input)) // Leerzeichen am Anfang und Ende entfernen
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" || input == "" {
			return false
		}
	}
}

// Zeigt einen Consolen Log an
func runtimeConsoleLog(values []interface{}) error {
	// Die Printline Werte werden ausgelesen
	printValues := exportStringArrayFromArgumentsAndFromateIt(values)

	// Ausgabe der Argumente auf der Go-Seite
	log.Println(printValues...)

	// Es ist kein Fehler aufgetreten
	return nil
}

// Zeigt einen Consolen Info-Log an
func runtimeInfoLog(values []interface{}) error {
	// Die Printline Werte werden ausgelesen
	printValues := exportStringArrayFromArgumentsAndFromateIt(values)

	// Ausgabe der Argumente auf der Go-Seite
	log.Println(printValues...)

	// Es ist kein Fehler aufgetreten
	return nil
}

// Zeigt eine Fehlermeldung auf der Console an
func runtimeErrorLog(values []interface{}) error {
	// Die Printline Werte werden ausgelesen
	printValues := exportStringArrayFromArgumentsAndFromateIt(values)

	// Ausgabe der Argumente auf der Go-Seite
	log.Println(printValues...)

	// Es ist kein Fehler aufgetreten
	return nil
}
