package fsysfile

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// Stellt die Index Datei dar
type Index struct {
	MaxSize string `yaml:"max_size"`
}

// FileSystem stellt das Virteulle Dateisystem dar
type FileSystem struct {
	rootPath   string
	indexYAML  *Index
	rootFolder *Folder
	fsMap      *FsMap
	porcLock   *sync.Mutex
}

// Gibt den Root Ordner zurück
func (o *FileSystem) GetRootDir() *Folder {
	return o.rootFolder
}

// Gibt die Maximale Obergrenze in Bytes an
func (o *FileSystem) GetMaxSize() int64 {
	// Es wird ermittelit ob keine begrenzung vorhanden ist
	if strings.ToLower(o.indexYAML.MaxSize) == "NULL" {
		return -1
	}

	return 0
}

// Gibt einen Ordner anhand seines Vollstäniges Pfaded zurück
func (o *FileSystem) GetDirByFullPath(fqPath string) (*Folder, error) {
	// Die Liste wird gesplittet
	splitedPath, err := splitAndValidatePath(fqPath)
	if err != nil {
		return nil, err
	}

	// Speichert den Aktuellen Ordner zwischen
	var rootPath *Folder

	// Es wird ermittelt ob es sich um den Root Path handelt
	if len(splitedPath) > 1 {
		if splitedPath[0] == "" {
			rootPath = o.rootFolder
		} else {
			// Der Ordner wird geöffnet
			rootPath = o.rootFolder.OpenFolder(splitedPath[0])
			if rootPath == nil {
				return nil, fmt.Errorf("unkown file 2: " + splitedPath[0])
			}
		}
	} else {
		rootPath = o.rootFolder
	}

	// Es wird versucht das nächste Verzeichniss zu ermitteln
	currentDir := rootPath
	for _, i := range splitedPath[1:] {
		currentDir = currentDir.OpenFolder(i)
		if currentDir == nil {
			return nil, fmt.Errorf("GetDirByFullPath: unkown folder " + i)
		}
	}

	// Der Aktuelle Ordner wird zurückgegeben
	return currentDir, nil
}

// Gibt eine Datei anhand ihres Vollständigen Pfaded zurück
func (o *FileSystem) GetFileByFullPath(fqPath string) (*File, error) {
	// Die Liste wird gesplittet
	splitedPath, err := splitAndValidatePath(fqPath)
	if err != nil {
		return nil, err
	}

	// Es wird ermittelt ob es sich bei dem Ersten Ordner um den Root Ordnader handelt
	var passedRootFolderArray []string
	if splitedPath[0] == "" {
		passedRootFolderArray = splitedPath[1:]
	} else {
		passedRootFolderArray = splitedPath
	}

	// Es wird versucht das nächste Verzeichniss zu ermitteln
	var resultFile *File
	currentDir := o.rootFolder
	for h, i := range passedRootFolderArray[:] {
		if h == len(passedRootFolderArray[:])-1 {
			resultFile = currentDir.GetFile(i)
			if currentDir == nil {
				return nil, fmt.Errorf("unkown file: " + i)
			}
		} else {
			currentDir = currentDir.OpenFolder(i)
			if currentDir == nil {
				return nil, fmt.Errorf("unkown folder: " + i)
			}
		}
	}

	// Der Aktuelle Ordner wird zurückgegeben
	return resultFile, nil
}

// Gibt einen Ordner anhand seines Vollstäniges Pfaded zurück
func (o *FileSystem) DeleteDirByFullPath(fqPath string) error {
	// Die Liste wird gesplittet
	splitedPath, err := splitAndValidatePath(fqPath)
	if err != nil {
		return err
	}

	// Speichert den Aktuellen Ordner zwischen
	var rootPath *Folder

	// Es wird ermittelt ob es sich um den Root Path handelt
	if len(splitedPath) > 1 {
		if splitedPath[0] == "" {
			rootPath = o.rootFolder
		} else {
			// Der Ordner wird geöffnet
			rootPath = o.rootFolder.OpenFolder(splitedPath[0])
			if rootPath == nil {
				return fmt.Errorf("unkown file 2: " + splitedPath[0])
			}
		}
	} else {
		rootPath = o.rootFolder
	}

	// Es wird versucht das nächste Verzeichniss zu ermitteln
	currentDir := rootPath
	for _, i := range splitedPath[1:] {
		currentDir = currentDir.OpenFolder(i)
		if currentDir == nil {
			return fmt.Errorf("GetDirByFullPath: unkown folder " + i)
		}
	}

	// Der Aktuelle Ordner wird zurückgegeben
	return currentDir.Delete()
}

// Gibt eine Datei anhand ihres Vollständigen Pfaded zurück
func (o *FileSystem) DeleteFileByFullPath(fqPath string) error {
	// Die Liste wird gesplittet
	splitedPath, err := splitAndValidatePath(fqPath)
	if err != nil {
		return err
	}

	// Speichert den Aktuellen Ordner zwischen
	var rootPath *Folder

	// Es wird ermittelt ob es sich um den Root Path handelt
	if len(splitedPath) > 1 {
		if splitedPath[0] == "" {
			rootPath = o.rootFolder
		} else {
			// Der Ordner wird geöffnet
			rootPath = o.rootFolder.OpenFolder(splitedPath[0])
			if rootPath == nil {
				return fmt.Errorf("unkown file 2: " + splitedPath[0])
			}
		}
	} else {
		rootPath = o.rootFolder
	}

	// Es wird versucht das nächste Verzeichniss zu ermitteln
	var resultFile *File
	currentDir := rootPath
	for h, i := range splitedPath[1:] {
		if h == len(splitedPath[1:])-1 {
			resultFile = currentDir.GetFile(i)
			if currentDir == nil {
				return fmt.Errorf("GetDirByFullPath: unkown file " + i)
			}
		} else {
			currentDir = currentDir.OpenFolder(i)
			if currentDir == nil {
				return fmt.Errorf("GetDirByFullPath: unkown folder " + i)
			}
		}
	}

	// Der Aktuelle Ordner wird zurückgegeben
	return resultFile.Delete()
}

// Diese Funktion erstellt einen neuen Ordner
func (o *FileSystem) CreateNewFolder(fqPath string, mode int64, recursive bool) error {
	fmt.Println("CREATE NEW FOLDER:", fqPath, mode, recursive)
	return nil
}

// Diese Funktion Kopiert eine Datei
func (o *FileSystem) CopyFile(sourceFqdnPath string, destinationFqdnPath string) error {
	return nil
}

// Diese Funktion Gibt an ob es sich bei dem Pfad um einen Link handelt
func (o *FileSystem) GetLinkPath(fqdnPath string) (string, error) {
	return "", nil
}

// Gibt den Inhalt eines Verzeichniss zurück
func (o *FileSystem) ReadDir(fqdnPath string) ([]interface{}, error) {
	return nil, nil
}

// Gibt die Informationen zu einer Datei an
func (o *FileSystem) GetStat(fqdnPath string) (interface{}, error) {
	// Die Liste wird gesplittet
	splitedPath, err := splitAndValidatePath(fqdnPath)
	if err != nil {
		return nil, err
	}

	// Speichert den Aktuellen Ordner zwischen
	var rootPath *Folder

	// Es wird ermittelt ob es sich um den Root Path handelt
	if len(splitedPath) > 1 {
		if splitedPath[0] == "" {
			rootPath = o.rootFolder
		} else {
			// Der Ordner wird geöffnet
			rootPath = o.rootFolder.OpenFolder(splitedPath[0])
			if rootPath == nil {
				return nil, fmt.Errorf("unkown file 2: " + splitedPath[0])
			}
		}
	} else {
		rootPath = o.rootFolder
	}

	// Es wird versucht das nächste Verzeichniss zu ermitteln
	currentDir := rootPath
	for h, i := range splitedPath[1:] {
		// Es wird ermittelt ob es sich um einen Ordner handelt
		currentDir = currentDir.OpenFolder(i)
		if currentDir != nil {
			// Es wird ermittelt ob es sich um den letzten Eintrag handelt
			if len(splitedPath[1:])-1 == h {
				return currentDir.GetStat(), nil
			}

			// Nächste runde
			continue
		}

		// Es wird ermittelt ob es sich um den letzten Eintrag handelt
		if len(splitedPath[1:])-1 == h {
			return nil, fmt.Errorf("invalid path")
		}

		// Es wird ermittelt ob es sich um eine Datei handelt
		currentFile := currentDir.GetFile(i)
		if currentFile != nil {
			return currentFile.GetStat(), nil
		}
	}

	// Der Aktuelle Ordner wird zurückgegeben
	return nil, fmt.Errorf("path not found")
}

// Gibt eine Datei anhand ihres File Descriptors zurück
func (o *FileSystem) GetFileByDescriptor(fileDescriptor string) (*File, error) {
	return nil, nil
}

// Schließt einen File Descriptor
func (o *FileSystem) CloseFileDescriptor(fileDescriptor string) error {
	return nil
}

// Öffnet einen Virteullen Dateisystem Container
func OpenVirtualFileSystemContainer(path string) (*FileSystem, error) {
	// Überprüfe, ob der Pfad ein Ordner ist
	if !IsValidateDirectoryPath(path) {
		return nil, fmt.Errorf("OpenVirtualFileSystemContainer: invalid container root path")
	}

	// Der Ordner Pfad wird formatiert
	formatedPath := FormatRootDirectoryPath(path)
	if formatedPath == "" {
		return nil, fmt.Errorf("internal error, unkown #1")
	}

	// Der Root Ordnerpatht wird formatiert
	rootPath := FormatFolder(formatedPath, "root")

	// Überprüfe, ob der Pfad ein Ordner ist
	if !IsValidateDirectoryPath(rootPath) {
		return nil, fmt.Errorf("OpenVirtualFileSystemContainer: invalid container root path")
	}

	// Es wird ermmittelt ob die Index Datei vorhanden ist
	indexFilePath := FormatPathWithFile(formatedPath, ".index")
	if indexFilePath == "" {
		return nil, fmt.Errorf("internal error, unkown #2")
	}

	// Es wird ermittelt ob der Aktuelle Benutzer berechtigt ist den Ordner zu verwenden
	permissions, err := checkFolderPermissions(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("OpenVirtualFileSystemContainer: " + err.Error())
	}
	if permissions.String() != "-rw-rw-rw-" {
		return nil, fmt.Errorf("OpenVirtualFileSystemContainer: You don't have the necessary rights")
	}

	// YAML-Datei lesen
	yamlData, err := os.ReadFile(indexFilePath)
	if err != nil {
		return nil, nil
	}

	// YAML-Daten in die Config-Struktur einlesen
	var indexYAML Index
	err = yaml.Unmarshal(yamlData, &indexYAML)
	if err != nil {
		return nil, fmt.Errorf("")
	}

	// Gibt an dass die Dateiübersicht erstellt wird
	fileMap := &FsMap{}

	// Das FileSystem Objekt wird erzeugt und zurückgegeben
	fsSystem := &FileSystem{
		rootPath:   rootPath,
		indexYAML:  &indexYAML,
		fsMap:      fileMap,
		rootFolder: &Folder{sysPath: rootPath, name: "", fsMap: fileMap},
		porcLock:   new(sync.Mutex),
	}

	// Das Dateisystem wird zurückgegeben
	return fsSystem, nil
}
