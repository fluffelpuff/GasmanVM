package imagefile

import (
	"archive/zip"
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/sha3"
)

func loadFilesMapFromReadCloser(file_stream io.ReadCloser) (*ContainerFilesMap, error) {
	// Erstellen Sie einen bufio.Scanner, der Zeile für Zeile liest
	reader := bufio.NewReader(file_stream)

	// Speichert die Final ausgelsenen Daten ab
	final_readed_data := make([]string, 0)

	// Speichert den Aktuellen Datensatz ab
	current_data_set := make([]byte, 0)

	// Schleife zum Lesen zeichenweise
	for {
		byter, err := reader.ReadByte()
		if err == io.EOF {
			final_readed_data = append(final_readed_data, string(current_data_set))
			break
		} else if err != nil {
			return nil, err
		}

		// Es wird geprüft ob es sich um einen Zeilenumbrauch handelt
		if byter == '\n' {
			final_readed_data = append(final_readed_data, string(current_data_set))
			current_data_set = make([]byte, 0)
			continue
		}

		// Der Datensatz wird hinzugefügt
		current_data_set = append(current_data_set, byter)
	}

	// Die Einzelenen Einträge werden ausgewertet
	file_infos := make(map[string]*ContainerFile)
	for _, i := range final_readed_data {
		spd := strings.Split(i, "\t")
		file_infos[spd[0]] = &ContainerFile{spd[0], spd[1], spd[2], false}
	}

	// Das Container File Map Objekt wird zurückgegeben
	cont_file_map := &ContainerFilesMap{file_infos}

	// Die Daten werden zurückgegeben
	return cont_file_map, nil
}

// Überprüft ob eine Datei so in der META/files.map vorhanden und gültig ist
func (o *ContainerFilesMap) validateFile(file *zip.File) bool {
	// Es wird geprüft ob die Datei vorhanden ist
	file_entry, foundit := o.files[file.Name]
	if !foundit {
		return false
	}

	// Die Datei wird geöffnet
	opened_file, err := file.Open()
	if err != nil {
		panic(err)
	}

	// Es wird ein Hash aus dieser Datei erstellt
	hasher := sha3.New256()
	_, err = io.Copy(hasher, opened_file)
	if err != nil {
		panic(err)
	}

	// Erhalten Sie den berechneten Hash
	result := hex.EncodeToString(hasher.Sum(nil))

	// Gültige Datei
	return result == file_entry.Hash
}

// Überprüft ob eine Datei so in der META/files.map vorhanden und gültig ist
func (o *ContainerFilesMap) validateFileAndMarkAsValidated(file *zip.File) bool {
	// Es wird geprüft ob die Datei bereits geprüft wurde
	result, foundit := o.files[file.Name]
	if !foundit {
		return false
	}
	if result.Validated {
		return true
	}

	// Es wird geprüft ob die Datei gültig ist
	if !o.validateFile(file) {
		return false
	}

	// Die Datei wird Validiert abgespeid
	o.files[file.Name].Validated = true

	return true
}

// Gibt an ob alles Daten der META/files.map validiert wurden
func (o *ContainerFilesMap) allFilesAreValidated() bool {
	for _, i := range o.files {
		if !i.Validated {
			return false
		}
	}
	return true
}

// Gibt alle Verfügabren Dateien an
func (o *ContainerFilesMap) getFile(file *zip.File) (*FileEntry, error) {
	filetr, foundit := o.files[file.Name]
	if !foundit {
		return nil, fmt.Errorf("file not found")
	}
	return &FileEntry{filetr, file}, nil
}
