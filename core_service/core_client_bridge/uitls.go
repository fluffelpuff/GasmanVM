package coreclientbridge

import "strings"

// Funktion zum Formatieren einer Liste von Strings
func formatStringList(list []string) string {
	// Verwendung von strings.Join, um die Liste zu verbinden
	// Der Separator ist ", "
	result := strings.Join(list, ", ")

	return result
}
