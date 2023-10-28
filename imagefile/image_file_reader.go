package imagefile

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

// Schließt das Image
func (o *ImageFileReader) Close() {
	o.zipFile.Close()
}

// Gibt eine Datei anhand ihres Namens zurück
func (o *ImageFileReader) GetFile(path string) (*FileEntry, error) {
	// Es wird versucht zu ermitteln ob die Datei exestiert
	result, found := o.fileMap[path]
	if !found {
		return nil, fmt.Errorf("ImageFileReader>GetFile: file '%s 'not found", path)
	}

	// Zu der Datei werden die Informationen abgerufen
	file_entry, err := o.files.getFile(result)
	if err != nil {
		return nil, fmt.Errorf("ImageFileReader>GetFile: (o.files.getFile) # " + err.Error())
	}

	// Das Objekt wird zurückgegeben
	return file_entry, nil
}

// Gibt eine Codedatei aus
func (o *ImageFileReader) GetSourceFile(path string) (*FileEntry, error) {
	// Es wird versucht die Mainfile abzurufen
	result, err := o.GetFile(fmt.Sprintf("src/%s", path))
	if err != nil {
		return nil, fmt.Errorf("GetMainFile: " + err.Error())
	}

	// Das ergbeniss wird zurückgegeben
	return result, nil
}

// Gibt die Maindatei zurück
func (o *ImageFileReader) GetMainFile() (*FileEntry, error) {
	// Es wird versucht die Mainfile abzurufen
	result, err := o.GetSourceFile(o.manifest.Application.GasmanVmSource)
	if err != nil {
		return nil, fmt.Errorf("GetMainFile: " + err.Error())
	}

	// Das ergbeniss wird zurückgegeben
	return result, nil
}

// Gibt an ob das Image ein Filesysten benötigt
func (o *ImageFileReader) NeedFileSystem() bool {
	return false
}

// Gibt die Rechte an welche das Paket benötigt
func (o *ImageFileReader) GetAllPermissions() []Permission {
	return o.manifest.Permissions
}

// Gibt die META Daten des Paketes zurück
func (o *ImageFileReader) GetMetaData() *METAINF {
	return o.metaInf
}

// Gibt die Manifest Datei zurück
func (o *ImageFileReader) GetManifest() *Manifest {
	return o.manifest
}

// Lädt eine Image Datei
func LoadImageFile(path string, show_info bool) (*ImageFileReader, error) {
	// Die Datei wird geöffnet
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("LoadImageFile: " + err.Error())
	}

	// Lesen Sie die ersten 2 Bytes der Datei
	header := make([]byte, 2)
	_, err = io.ReadFull(file, header)
	if err != nil {
		return nil, err
	}

	// Es wird geprüft ob es sich um eine ZIP handelt
	if string(header) != "PK" {
		return nil, fmt.Errorf("unkown image datatype")
	}

	// Ermitteln Sie die Größe der ZIP-Datei
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("LoadImageFile: " + err.Error())
	}

	// Öffnen Sie die ZIP-Datei
	zipFile, err := zip.NewReader(file, fileInfo.Size())
	if err != nil {
		return nil, fmt.Errorf("LoadImageFile: " + err.Error())
	}

	// Es wird ermittelt wie die Dekomprimierungsstufe ist
	for _, file := range zipFile.File {
		if file.Method != zip.Store {
			return nil, fmt.Errorf("LoadImageFile: It is a compressed ZIP package")
		}
	}

	// Die Dateiübersicht wird geladen "/META/files.map"
	var filesMap *ContainerFilesMap
	for _, file := range zipFile.File {
		if file.Name == "META/files.map" {
			// Öffnen und lesen Sie die JavaScript-Datei
			jsFile, err := file.Open()
			if err != nil {
				log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
			}
			defer jsFile.Close()

			// Die Datei wird eingelesen
			filesMap, err = loadFilesMapFromReadCloser(jsFile)
			if err != nil {
				log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
			}

			break
		}
	}
	if filesMap == nil {
		return nil, fmt.Errorf("unkown internal error")
	}

	// Es wird geprüft ob alle benötigten Dateien vorhanden sind
	fileMapZipLink := make(map[string]*zip.File, 0)
	for _, zip_file := range zipFile.File {
		// Sollte es sich um die META/files.map handelm wird der Vorgang übersprungen
		if zip_file.Name == "META/files.map" {
			continue
		}

		// Die Datei wird vorab geprüft
		if !filesMap.validateFileAndMarkAsValidated(zip_file) {
			fmt.Println(zip_file.Name)
			return nil, fmt.Errorf("broken container file, invalid file")
		}

		// Die Datei wird zwischengespeichert
		fileMapZipLink[zip_file.Name] = zip_file
	}

	// Es wird sichergestellt dass alle Dateien geprüft wurden
	if !filesMap.allFilesAreValidated() {
		return nil, fmt.Errorf("broken container file")
	}

	// Info
	if show_info {
		fmt.Println("The manifest file is read")
	}

	// Die Manifestdatei wird ermittelt
	var manifestFile io.ReadCloser
	for _, file := range zipFile.File {
		if file.Name == "GasmanVMManifest.xml" {
			// Öffnen und lesen Sie die JavaScript-Datei
			jsFile, err := file.Open()
			if err != nil {
				log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
			}

			manifestFile = jsFile
			break
		}
	}
	if manifestFile == nil {
		return nil, fmt.Errorf("invalid manifest")
	}
	defer manifestFile.Close()

	var manifest Manifest
	if err := xml.NewDecoder(manifestFile).Decode(&manifest); err != nil {
		log.Fatal(err)
	}

	// Info
	if show_info {
		fmt.Println("  > Data integrity is checked")
	}

	// Die META Informationen werden eingelesen
	package_sig := make([]byte, 0)
	var owner_cert *tls.Certificate
	for _, file := range zipFile.File {
		if strings.HasPrefix(file.Name, "META/") {
			if strings.HasSuffix(file.Name, "package.sig") { // Es handelt sich um die Paket Signatur
				// Öffnen und lesen Sie die JavaScript-Datei
				fx, err := file.Open()
				if err != nil {
					log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
				}
				defer fx.Close()

				// Die Datei wird eingelesen
				package_sig, err = io.ReadAll(fx)
				if err != nil {
					log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
				}

				// Nächste Runde
				continue
			} else if strings.HasSuffix(file.Name, "builder.crt") { // Es handelt sich um ein Zertifikat
				// Öffnen und lesen Sie die JavaScript-Datei
				fx, err := file.Open()
				if err != nil {
					return nil, fmt.Errorf("LoadImageFile: " + err.Error())
				}
				defer fx.Close()

				// Die Datei wird eingelesen
				readed_cert, err := io.ReadAll(fx)
				if err != nil {
					return nil, fmt.Errorf("LoadImageFile: " + err.Error())
				}

				// Das Zert wird eingelesen
				owner_cert_r, err := tls.X509KeyPair(readed_cert, nil)
				if err != nil {
					return nil, fmt.Errorf("LoadImageFile: " + err.Error())
				}

				// Das Cert wird zwischengespesichert
				owner_cert = &owner_cert_r

				// Nächste Runde
				continue
			} else {
				// Sollte es sich um die META/files.map handeln wird der Vorgang abgebrochen
				if file.Name == "META/files.map" {
					continue
				}

				// Der Vorgang wird übersprungen
				return nil, fmt.Errorf("LoadImageFile: unkown META-DATA >> " + file.Name)
			}
		}
	}

	// Info
	if show_info {
		fmt.Println("The optional information is read in:")
	}

	// Die OPT-IN Informationen werden eingelesen
	certs, domains := make([]*x509.Certificate, 0), make([]*DomainEntry, 0)
	for _, file := range zipFile.File {
		if strings.HasPrefix(file.Name, "optin/") {
			if strings.HasPrefix(file.Name, "optin/ssl") {
				if strings.HasSuffix(file.Name, ".crt") { // Es handelt sich um die Paket Signatur
					// Öffnen und lesen Sie die JavaScript-Datei
					fx, err := file.Open()
					if err != nil {
						log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
					}
					defer fx.Close()

					// Die Datei wird eingelesen
					cert_file, err := io.ReadAll(fx)
					if err != nil {
						log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
					}

					// Extrahieren Sie das Zertifikat aus der Datei
					block, _ := pem.Decode(cert_file)
					if block == nil {
						fmt.Println("Kein gültiges PEM-Format in der Datei gefunden.")
						return nil, err
					}

					// Erstellen Sie ein X.509-Zertifikat aus den Zertifikatsdaten
					cert, err := x509.ParseCertificate(block.Bytes)
					if err != nil {
						fmt.Println("Fehler beim Laden des Zertifikats:", err)
					}

					// Das Zert wird zwischengespeichert
					certs = append(certs, cert)

					// Info
					if show_info {
						// Erstellen Sie einen neuen SHA-3-256-Hasher
						hasher := sha3.New224()

						// Fügen Sie die Daten dem Hasher hinzu
						hasher.Write(cert.Raw)

						// Holen Sie den 256-Bit-Hashwert als Slice von Bytes
						hashValue := hasher.Sum(nil)

						fmt.Printf("  > Certificate '%x' added\n", hashValue)
					}

					// Nächste Runde
					continue
				} else {
					return nil, fmt.Errorf("LoadImageFile: not allowed file type")
				}
			} else if strings.HasPrefix(file.Name, "optin/domains") {
				if strings.HasSuffix(file.Name, ".domain") { // Es handelt sich um die Paket Signatur
					// Öffnen und lesen Sie die JavaScript-Datei
					fx, err := file.Open()
					if err != nil {
						log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
					}
					defer fx.Close()

					// Dekodieren der XML-Daten
					var domainstore *DomainEntry
					decoder := xml.NewDecoder(fx)
					if err := decoder.Decode(&domainstore); err != nil {
						fmt.Println("Fehler beim Dekodieren der XML-Daten:", err)
					}

					// Die Domain wird zwischen gespeichert
					domains = append(domains, domainstore)

					// Info
					if show_info {
						fmt.Println("  > Domain added")
					}

					// Nächste Runde
					continue
				} else {
					return nil, fmt.Errorf("LoadImageFile: not allowed file type")
				}
			} else {
				return nil, fmt.Errorf("LoadImageFile: unkown optin data")
			}
		}
	}

	// Die Domains werden zusammengefasst
	optin_obj := &OptInOptions{certs, domains}

	// Das Meta INF Objekt wird erstellt
	minf_obj := &METAINF{package_sig, owner_cert}

	// Die Imagedatei wird erstellt
	image_object := &ImageFileReader{optin_obj, &manifest, minf_obj, filesMap, fileMapZipLink, zipFile, file}

	// Das Objekt wurd ohne einen Fehler erstellt und wird zurückgegeben
	return image_object, nil
}
