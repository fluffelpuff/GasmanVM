package imagefile

import (
	"encoding/hex"
	"io"

	"golang.org/x/crypto/sha3"
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

// Gibt den Hash aus
func (o *FileEntry) GetHash() string {
	// Die Datei wird geöffnet
	jsFile, err := o.scriptDataLink.Open()
	if err != nil {
		return ""
	}
	defer jsFile.Close()

	// Erstellen Sie einen SHA-3-Hasher
	hasher := sha3.New256()

	// Kopieren Sie den Inhalt des io.ReadCloser in den Hasher
	_, err = io.Copy(hasher, jsFile)
	if err != nil {
		return ""
	}

	// Berechnen Sie den SHA-3-Hash
	hash := hasher.Sum(nil)

	// Die Daten werden zurückgegeben
	return hex.EncodeToString(hash)
}
