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
	CRC32          uint32
}

type FilePlayground struct {
	Path       string
	Playground string
	Hash       string
	CRC32      uint32
}

type Manifest struct {
	MainFile    string `xml:"main,attr"`
	Rights      Rights `xml:"rights"`
	SourceFiles Files  `xml:"files"`
}

type METAINF struct {
	packageSig     []byte
	packageCreator *tls.Certificate
}

type OptInOptions struct {
	Certs   []*x509.Certificate
	Domains []*DomainEntry
}

type Right struct {
	Name string `xml:"name,attr"`
}

type Rights struct {
	Rights []Right `xml:"right"`
}

type File struct {
	Path       string `xml:"path,attr"`
	Playground string `xml:"playground,attr"`
	Hash       string `xml:"hash,attr"`
}

type Files struct {
	Files []File `xml:"file"`
}

type Data struct {
	Main           string `xml:"main,attr"`
	Rights         Rights `xml:"rights"`
	FilePlayground Files  `xml:"files"`
}
