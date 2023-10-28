package fsysfile

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// Gibt die Verfügbare Größe des RAMS zurück
func getAvailableRAM() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Sys
}

// Gibt die Größe einer Datei zurück
func fileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// IsValidName prüft, ob der übergebene Name ein gültiger Ordner- oder Dateiname nach Unix-Regeln ist.
func IsValidName(name string) bool {
	// Unix-Regeln: Ein Name darf nur aus Buchstaben, Zahlen, Bindestrichen und Punkten bestehen.
	// Ein Name darf nicht mit einem Punkt oder Bindestrich beginnen und nicht aus mehr als einem Punkt bestehen.
	// Ein Name darf nicht mehr als 255 Zeichen haben.

	// Prüfe, ob der Name länger als 255 Zeichen ist.
	if len(name) > 255 {
		return false
	}

	// Überprüfe, ob der Name mit einem Punkt oder Bindestrich beginnt.
	if name[0] == '.' || name[0] == '-' {
		return false
	}

	// Überprüfe, ob der Name mehr als einen Punkt enthält.
	if regexp.MustCompile(`\.{2,}`).MatchString(name) {
		return false
	}

	// Überprüfe, ob der Name nur aus erlaubten Zeichen besteht.
	if !regexp.MustCompile(`^[a-zA-Z0-9\.-]+$`).MatchString(name) {
		return false
	}

	return true
}

// IsValidUnixPath überprüft, ob der übergebene Pfad ein gültiger Unix-Pfad ist.
func isValidUnixPath(path string) bool {
	// Unix-Pfade dürfen nur aus Buchstaben, Zahlen, Bindestrichen, Unterstrichen, Punkten und Schrägstrichen bestehen.
	// Ein Pfad darf nicht mit einem Schrägstrich enden.

	// Überprüfe, ob der Pfad mit einem Schrägstrich endet.
	if path[len(path)-1] == '/' {
		return false
	}

	// Überprüfe, ob der Pfad nur aus erlaubten Zeichen besteht.
	if !regexp.MustCompile(`^[a-zA-Z0-9_./-]+$`).MatchString(path) {
		return false
	}

	return true
}

// Splitet einen Path auf
func splitPath(fqdnPath string) []string {
	return strings.Split(fqdnPath, "/")
}

// Splittet einen Path und Validiert diesen dann
func splitAndValidatePath(fqdnPath string) ([]string, error) {
	// Es wird ermittelt ob es sich um einen zulässigen Path handelt
	if result := isValidUnixPath(fqdnPath); !result {
		return nil, fmt.Errorf("invalid path: " + fqdnPath)
	}

	// Der Path wird gesplittet
	splittetPath := splitPath(fqdnPath)

	// Der Splittet Path wird zurückgegeben
	return splittetPath, nil
}
