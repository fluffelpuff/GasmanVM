package imagefile

import (
	"io"
)

// Gibt die Bytes einer Zip Datei aus
func (o *FileEntry) GetBytes() []byte {
	// Die Datei wird geöffnet
	jsFile, err := o.scriptDataLink.Open()
	if err != nil {
		return nil
	}
	defer jsFile.Close()

	// Die Daten werden eingelesen
	file_data, err := io.ReadAll(jsFile)
	if err != nil {
		return nil
	}

	// Die Daten werden zurückgegeben
	return file_data
}
