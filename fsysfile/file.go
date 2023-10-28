package fsysfile

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

// Stellt eine Datei dar
type File struct {
	fsMap   *FsMap
	sysPath string
}

// Prüft ob eine Datei exestiert und Öffnet diese
func openFileAndCheckBase(filePath string) ([]byte, error) {
	// Es wird ermittelt ob der Path exestiert und
	if !IsFile(filePath) {
		return nil, fmt.Errorf("file not found")
	}

	// Es wird ermittelt ob ausreichend RAM verfügbar ist um die Datei vollständig zu laden
	if !ValidateFileSizeRAMAvail(filePath) {
		return nil, fmt.Errorf("file to big")
	}

	// Die Datei wird eingelesen
	rFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("internal error, file not found")
	}
	if rFile == nil {
		return nil, fmt.Errorf("internal error, file not found")
	}

	// Die Datei wird zurückgegeben
	return rFile, nil
}

// Löscht die Aktuelle Datei
func (o *File) Delete() error {
	return nil
}

// Gibt den Inhalt der Datei als UTF-8 String zurück
func (o *File) OpenUtf8() (string, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return "", err
	}

	// Die Datei wird in einen String umgewandelt
	utf8ConvertedString := string(rFile)

	// Der Inhalt der Datei wird zurückgegeben
	return utf8ConvertedString, nil
}

// Gibt den Inhalt der Datei als Bytes zurück
func (o *File) OpenBinary() ([]byte, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return nil, err
	}

	// Die Bytes werden zurückgegeben
	return rFile, nil
}

// Gibt den Inahlt der Datei als latin1 zurück
func (o *File) OpenLatin1() (string, error) {
	// Es wird ermittelt ob der Path exestiert und
	if !IsFile(o.sysPath) {
		return "", fmt.Errorf("file not found")
	}

	// Es wird ermittelt ob ausreichend RAM verfügbar ist um die Datei vollständig zu laden
	if !ValidateFileSizeRAMAvail(o.sysPath) {
		return "", fmt.Errorf("file to big")
	}

	// Öffne die Datei im Latin-1 (ISO-8859-1) Codierungsmodus
	file, err := os.OpenFile(o.sysPath, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Erstelle einen Reader mit Latin-1 Codierung
	reader := bufio.NewReader(file)

	// Verwende einen strings.Builder, um den Inhalt in einen String zu schreiben
	var contentBuilder strings.Builder

	// Lese den Inhalt der Datei Zeile für Zeile
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // Dateiende erreicht
		}
		if err != nil {
			return "", err
		}

		// Hier haben Sie 'line' als String mit Latin-1 Codierung
		contentBuilder.WriteString(line)
	}

	// Den gesamten Inhalt als String erhalten
	content := contentBuilder.String()

	// Der String wird zurückgegeben
	return content, nil
}

// Gibt den Inhalt der Datei als Base64 zurück
func (o *File) OpenBase64() (string, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return "", err
	}

	// Die Daten werden mittels Base64 Kodiert
	base64String := base64.StdEncoding.EncodeToString(rFile)

	// Der Inhalt wird in einen Base64 String umgewandelt
	return base64String, nil
}

// Gibt den Inhalt der Datei als Base58 zurück
func (o *File) OpenBase58() (string, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return "", err
	}

	// Die Daten werden mittels Hex kodiert
	base58String := base58.Encode(rFile)

	// Der Wert wird zurückgegeben
	return base58String, nil
}

// Gibt den Inhalt der Datei als Base32 zurück
func (o *File) OpenBase32() (string, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return "", err
	}

	// Die Daten werden mittels Hex kodiert
	base32String := base64.StdEncoding.EncodeToString(rFile)

	// Der Wert wird zurückgegeben
	return base32String, nil
}

// Gibt den Inhalt der Datei als Hexstring zurück
func (o *File) OpenHex() (string, error) {
	// Es wird geprüft ob die Datei vorhanden ist, wenn ja wird diese geöffnet
	rFile, err := openFileAndCheckBase(o.sysPath)
	if err != nil {
		return "", err
	}

	// Die Daten werden mittels Hex kodiert
	hexString := hex.EncodeToString(rFile)

	// Der Wert wird zurückgegeben
	return hexString, nil
}

// Wird verwendet um die Datei umzubenenen
func (o *File) Rename(newName string) error {
	return nil
}

// Gibt den Aktuellen Namen zurück
func (o *File) GetName() string {
	return o.sysPath
}

// Gibt die FSStat Informationen zurück
func (o *File) GetStat() *FileStat {
	return nil
}

// Erzeugt einen neuen File Descriptor
func (o *File) GetFileDescriptor(flags string) (string, error) {
	return "0x111", nil
}
