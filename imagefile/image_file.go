package imagefile

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

// Stellt das Image dar
type ImageFile struct {
	optIn    *OptInOptions
	manifest *Manifest
	metaInf  *METAINF
	files    map[string]*FileEntry
	zipObj   *zip.ReadCloser
}

// Schließt das Image
func (o *ImageFile) Close() {
	o.zipObj.Close()
}

// Gibt eine Datei anhand ihres Namens zurück
func (o *ImageFile) GetFile(path string) (*FileEntry, error) {
	// Es wird versucht zu ermitteln ob die Datei exestiert
	result, found := o.files[path]
	if !found {
		return nil, fmt.Errorf("ImageFile>GetFile: file not found")
	}

	// Das Objekt wird zurückgegeben
	return result, nil
}

// Gibt eine Codedatei aus
func (o *ImageFile) GetSourceFile(path string) (*FileEntry, error) {
	// Es wird versucht die Mainfile abzurufen
	result, err := o.GetFile(fmt.Sprintf("src/%s", path))
	if err != nil {
		return nil, fmt.Errorf("GetMainFile: " + err.Error())
	}

	// Das ergbeniss wird zurückgegeben
	return result, nil
}

// Gibt die Maindatei zurück
func (o *ImageFile) GetMainFile() (*FileEntry, error) {
	// Es wird versucht die Mainfile abzurufen
	result, err := o.GetSourceFile(o.manifest.MainFile)
	if err != nil {
		return nil, fmt.Errorf("GetMainFile: " + err.Error())
	}

	// Das ergbeniss wird zurückgegeben
	return result, nil
}

// Lädt eine Image Datei
func LoadImageFile(path string) (*ImageFile, error) {
	// Öffnen Sie die ZIP-Datei
	zipFile, err := zip.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("LoadImageFile: " + err.Error())
	}

	// Die Manifestdatei wird ermittelt
	var manifestFile []byte
	for _, file := range zipFile.File {
		if file.Name == "GasmanVMManifest.xml" {
			// Öffnen und lesen Sie die JavaScript-Datei
			jsFile, err := file.Open()
			if err != nil {
				log.Fatalf("Fehler beim Öffnen der JavaScript-Datei im ZIP-Archiv: %v", err)
			}
			defer jsFile.Close()

			manifestFile, err = io.ReadAll(jsFile)
			if err != nil {
				log.Fatalf("Fehler beim Lesen der JavaScript-Datei aus dem ZIP-Archiv: %v", err)
			}
			break
		}
	}

	// Die Manifest Datei wird eingelesen
	var manifest Manifest
	if err := json.Unmarshal(manifestFile, &manifest); err != nil {
		return nil, fmt.Errorf("LoadImageFile: " + err.Error())
	}

	// Es wird ermittelt ob alle benötigten Dateien vorhanden sind
	for _, otem := range manifest.FilePlaygrounds {
		found_file := false
		for _, file := range zipFile.File {
			if string("src/"+otem.Path) == file.Name {
				found_file = true
				break
			}
		}
		if !found_file {
			return nil, fmt.Errorf("LoadImageFile: file not found")
		}
	}

	// Es wird ermittelt ob die Datei Hashes korrekt sind
	file_entrys := make(map[string]*FileEntry)
	for _, otem := range zipFile.File {
		if strings.HasPrefix(otem.Name, "src/") {
			found_file := false
			for _, file := range manifest.FilePlaygrounds {
				if otem.Name == string("src/"+file.Path) {
					file_entrys[otem.Name] = &FileEntry{file.Playground, otem}
					found_file = true
					break
				}
			}
			if !found_file {
				return nil, fmt.Errorf("LoadImageFile: invalid script ")
			}
		}
	}

	// Die META Informationen werden eingelesen
	package_sig := make([]byte, 0)
	var owner_cert tls.Certificate
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
				owner_cert, err = tls.X509KeyPair(readed_cert, nil)
				if err != nil {
					return nil, fmt.Errorf("LoadImageFile: " + err.Error())
				}

				// Nächste Runde
				continue
			} else { // Es handelt sich um ein unebaknntes Paket
				return nil, fmt.Errorf("LoadImageFile: unkown META-DATA")
			}
		}
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
	minf_obj := &METAINF{package_sig, &owner_cert}

	// Die Imagedatei wird erstellt
	image_object := &ImageFile{optin_obj, &manifest, minf_obj, file_entrys, zipFile}

	// Das Objekt wurd ohne einen Fehler erstellt und wird zurückgegeben
	return image_object, nil
}