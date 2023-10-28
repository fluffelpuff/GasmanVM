package fsysfile

import (
	"fmt"
	"strings"
)

// FormatRootDirectoryPath formatiert, einen Path String
func FormatRootDirectoryPath(path string) string {
	if strings.HasSuffix(path, "\\") {
		return path
	}
	return fmt.Sprintf("%s\\", path)
}

// FormatPathWithFile formatiert, einen Path Datei String
func FormatPathWithFile(path string, file string) string {
	return fmt.Sprintf("%s%s", path, file)
}

// FormatFolder formatiert, einen Ordner Path
func FormatFolder(rpath string, name string) string {
	if strings.HasSuffix(rpath, "\\") {
		return fmt.Sprintf("%s%s\\", rpath, name)
	}
	return fmt.Sprintf("%s\\%s\\", rpath, name)
}
