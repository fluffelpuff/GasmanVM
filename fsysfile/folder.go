package fsysfile

import (
	"fmt"
	"os"
)

// Stellt einen Ordner dar
type Folder struct {
	fsMap   *FsMap
	sysPath string
	name    string
}

// Öffnet einen Ordner
func (o *Folder) OpenFolder(name string) *Folder {
	// Der neue Path wird Formatiert
	newPath := FormatFolder(o.sysPath, name)

	// Es wird versucht den Ordner zu öffnen
	if !IsDirectory(newPath) {
		return nil
	}

	// Das neue Objekt wird zurückgegeben
	return &Folder{sysPath: newPath, name: name, fsMap: o.fsMap}
}

// Öffnet eine Datei
func (o *Folder) GetFile(name string) *File {
	// Es wird ermittelt ob es sich um eine FileSystem Datei handelt
	if IsFsFileExt(name) {
		return nil
	}

	// Der neue Path wird Formatiert
	newPath := FormatPathWithFile(o.sysPath, name)

	// Es wird geprüft ob es sich um eine Datei handelt
	if !IsFile(newPath) {
		return nil
	}

	// Das Datei Objekt wird erstellt und zurückgegeben
	return &File{sysPath: newPath, fsMap: o.fsMap}
}

// Erstellt einen neuen Ordner
func (o *Folder) CreateNewFolder(name string) error {
	// Der neue Path wird Formatiert
	newPath := FormatFolder(o.sysPath, name)

	// Der Ordner wrid erstellt
	err := os.Mkdir(newPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Der Vorgang wurde ohne Fehler durchgeführt
	return nil
}

// Löscht einen Ordner
func (o *Folder) DeleteFolder(name string) error {
	// Der neue Path wird Formatiert
	newPath := FormatFolder(o.sysPath, name)

	// Der Ordner wird gelöscht
	err := os.RemoveAll(newPath)
	if err != nil {
		return err
	}

	// Der Vorgang wurde ohne Fehler durchgeführt
	return nil
}

// Löscht eine Datei im Ordner
func (o *Folder) DeleteFile(name string) error {
	// Die Datei wird abgerufen
	file := o.GetFile(name)
	if file == nil {
		return fmt.Errorf("unkown file")
	}

	// Die Datei wird gelöscht
	return file.Delete()
}

// Löscht den Aktuellen Ordner
func (o *Folder) Delete() error {
	return nil
}

// Gibt den Namen des Ordners zurück
func (o *Folder) GetName() string {
	return o.name
}

// Gibt den Folder Stat zurück
func (o *Folder) GetStat() *FolderStat {
	return nil
}
