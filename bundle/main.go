package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"archive/zip"

	"golang.org/x/crypto/sha3"
)

type FileLink struct {
	intrname string
	file     string
	hash     string
}

func calculateSHA3Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha3.New256()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash), nil
}

func listFilesAndFolders(path string) ([]FileLink, error) {
	files := []FileLink{}
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Fehler beim Durchsuchen des Verzeichnisses:", err)
			return err
		}
		if !info.IsDir() {
			hash, err := calculateSHA3Hash(filePath)
			if err != nil {
				panic(err)
			}
			repl := strings.ReplaceAll(filePath, path, "")
			repl = strings.ReplaceAll(repl, "\\", "/")
			repl = strings.TrimPrefix(repl, "/")
			files = append(files, FileLink{repl, filePath, hash})
		}
		return nil
	})

	if err != nil {
		fmt.Println("Fehler beim Durchsuchen des Verzeichnisses:", err)
	}

	return files, nil
}

func checkContainerIsBundleable(files []FileLink) bool {
	// Es wird ermittelt ob es eine Manifest Datei gibt
	var container_file *FileLink
	for _, i := range files {
		if i.intrname == "GasmanVMManifest.xml" {
			container_file = &i
			break
		}
	}

	// Es wird geprüft ob der Container vorhanden ist
	if container_file == nil {
		return false
	}

	// Es wird ermittelt ob es einen META Ordner gibt
	has_meta_folder := 0
	for _, i := range files {
		if strings.HasPrefix(i.intrname, "META/") {
			has_meta_folder++
		}
	}

	return has_meta_folder == 0
}

func writeFileInImage(zipWriter *zip.Writer, name string, host_file_path string, kompressionLevel uint16) error {
	// Konfigurieren Sie die Kompressionsstufe für die Datei
	header, err := zipWriter.CreateHeader(&zip.FileHeader{
		Name:   name,
		Flags:  0,         // Keine weiteren Flags für die Datei
		Method: zip.Store, // Keine Kompression
	})
	if err != nil {
		return err
	}

	// Öffnen Sie die Datei von der Festplatte
	file, err := os.Open(host_file_path)
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return err
	}
	defer file.Close()

	// Inhalt der Datei von der Festplatte in die ZIP-Datei kopieren
	_, err = io.Copy(header, file)
	if err != nil {
		fmt.Println("Fehler beim Kopieren der Datei in die ZIP-Datei:", err)
		return err
	}

	// Schreiben Sie den Inhalt der Datei
	return nil
}

func createAndWriteFileMap(zipWriter *zip.Writer, files []FileLink) error {
	// Der Block wird erstellt
	complete_block := make([]byte, 0)

	// Die Einzelnene Dateieien werden abgerufen
	for h, i := range files {
		var fileType string
		switch {
		case strings.HasSuffix(i.intrname, ".js"):
			fileType = "javascript"
		case strings.HasSuffix(i.intrname, ".lisp"):
			fileType = "lsip"
		case strings.HasSuffix(i.intrname, ".crt"):
			fileType = "certificate"
		case strings.HasSuffix(i.intrname, ".javascript"):
			fileType = "javascript"
		case i.intrname == "GasmanVMManifest.xml":
			fileType = "manifest"
		default:
			return fmt.Errorf("invalid data type: '" + i.intrname + "'")
		}

		// Der Eintrag wird erzeugt
		var completed_line string
		if h == len(files)-1 {
			completed_line = fmt.Sprintf("%s\t%s\t%s", i.intrname, fileType, i.hash)
		} else {
			completed_line = fmt.Sprintf("%s\t%s\t%s\n", i.intrname, fileType, i.hash)
		}

		// Der Eintrag wird hinzugefügr
		complete_block = append(complete_block, []byte(completed_line)...)
	}

	// Der Header wird erzeugt
	file, err := zipWriter.CreateHeader(&zip.FileHeader{
		Name:   "META/files.map",
		Flags:  0,
		Method: zip.Store,
	})
	if err != nil {
		return err
	}

	// Schreiben Sie den Inhalt der Datei
	_, err = file.Write(complete_block)
	return err
}

func main() {
	// Parsen der Befehlszeilenargumente
	flag.Parse()

	// Das erste nicht mit einem Flag versehene Argument ist der Dateipfad
	var folderPath string
	var buildImagePath string
	if flag.NArg() >= 2 {
		folderPath = flag.Arg(0)
		buildImagePath = flag.Arg(1)
	}

	// Überprüfen, ob ein Dateipfad angegeben wurde
	if folderPath == "" || buildImagePath == "" {
		fmt.Println("Verwendung: meinprogramm.exe [<Dateipfad>] [-tag1 <Tag1>] [-tag2 <Tag2>]")
		return
	}

	// Überprüfen, ob der Ordner existiert
	_, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Der Ordner existiert nicht.")
		} else {
			fmt.Println("Fehler beim Überprüfen des Ordners:", err)
		}
		return
	}

	// Es wird eine übersicht aller Dateien und Ordner im Ordner erstellt
	files, err := listFilesAndFolders(folderPath)
	if err != nil {
		panic(err)
	}

	// Es wird geprüft ob der Container alle benötigten Dateien entält
	if !checkContainerIsBundleable(files) {
		panic("invalid container folder")
	}

	// Erstellen oder öffnen Sie die ZIP-Datei
	newZipFile, err := os.Create(buildImagePath)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der ZIP-Datei:", err)
		return
	}
	defer newZipFile.Close()

	// Erstellen Sie einen neuen ZIP-Writer
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Die Dateieen werden in das ZIP Archiv geschrieben
	for _, i := range files {
		// Die Dateien werden geschrieben
		writeFileInImage(zipWriter, i.intrname, i.file, zip.Store)
	}

	// Die File Map wird erstellt und geschrieben
	if err := createAndWriteFileMap(zipWriter, files); err != nil {
		panic(err)
	}
}
