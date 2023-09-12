package imagefile

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
)

type DomainEntry struct {
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Price  float64 `xml:"price"`
}

type FileEntry struct {
	ScripteType    string
	scriptDataLink *zip.File
}

type FilePlayground struct {
	Path       string `json:"path"`
	Playground string `json:"playground"`
}

type Manifest struct {
	MainFile        string            `json:"main"`
	FilePlaygrounds []*FilePlayground `json:"file_playground"`
}

type METAINF struct {
	packageSig     []byte
	packageCreator *tls.Certificate
}

type OptInOptions struct {
	Certs   []*x509.Certificate
	Domains []*DomainEntry
}
