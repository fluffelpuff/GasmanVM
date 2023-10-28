package fsysfile

import (
	"os"
	"path/filepath"
)

// IsValidPath überprüft, ob der Pfad gültig ist
func IsValidPath(path string) bool {
	return filepath.IsAbs(path)
}

// IsDirectory überprüft, ob es sich bei einem Pfad um einen Ordner handelt
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// IsValidateDirectoryPath überprüft, ob es sich um einen Ordner handelt und überpüft ob der Path korrekt ist
func IsValidateDirectoryPath(path string) bool {
	if !IsValidPath(path) {
		return false
	}
	return IsDirectory(path)
}

// IsFile prüft ob eine Datei vorhanden ist
func IsFile(path string) bool {
	// Überprüfen, ob die Datei existiert
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// IsFsFile gibt an, ob es sich um einen
func IsFsFileExt(fName string) bool {
	return false
}

// Ermittelt die Berechtigungen für den Ordner
func checkFolderPermissions(path string) (os.FileMode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return info.Mode().Perm(), nil
}

// Ermittelt ob eine Datei Größer als der RAM ist
func ValidateFileSizeRAMAvail(path string) bool {
	fileSize, err := fileSize(path)
	if err != nil {
		return false
	}

	// Vergleiche die Dateigröße mit der verfügbaren RAM-Größe
	if int64(fileSize) > int64(getAvailableRAM()) {
		return false
	} else {
		return true
	}
}
