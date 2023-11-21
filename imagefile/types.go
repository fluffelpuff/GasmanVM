package imagefile

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"os"
)

type DomainEntry struct {
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Price  float64 `xml:"price"`
}

type ContainerFile struct {
	Path       string
	Playground string
	Hash       string
	Validated  bool
}

type FileEntry struct {
	File           *ContainerFile
	scriptDataLink *zip.File
}

type ContainerFilesMap struct {
	files map[string]*ContainerFile
}

type Manifest struct {
	XMLName     xml.Name     `xml:"manifest"`
	Package     Package      `xml:"package"`
	Application Application  `xml:"application"`
	Domains     Domains      `xml:"domains"`
	Permissions []Permission `xml:"permission"`
	Groups      Groups       `xml:"groups"`
}

type Package struct {
	GasmanVmPackage string `xml:"gasmanvm-package,attr"`
	GasmanVmVersion string `xml:"gasmanvm-version,attr"`
	GasmanVmGit     string `xml:"gasmanvm-git,attr"`
}

type Group struct {
	Certificate string `xml:",innerxml"`
}

type Groups struct {
	Groups []Group `xml:"group"`
}

type Application struct {
	GasmanVmFilesysimage string `xml:"gasmanvm-filesysimage,attr"`
	GasmanVmMode         string `xml:"gasmanvm-mode,attr"`
	GasmanVmAssets       string `xml:"gasmanvm-assets,attr"`
	GasmanVmSyntic       string `xml:"gasmanvm-syntic,attr"`
	GasmanVmSource       string `xml:"gasmanvm-source,attr"`
	GasmanVmSourceMain   string `xml:"gasmanvm-source_main,attr"`
	GasmanVmDescription  string `xml:"gasmanvm-description,attr"`
	GasmanVmLabel        string `xml:"gasmanvm-label,attr"`
	GasmanVmSslstore     string `xml:"gasmanvm-sslstore,attr"`
}

type Domains struct {
	Domain []Domain `xml:"domain"`
}

type Domain struct {
	GasmanVmHost string `xml:"gasmanvm-host,attr"`
	Specific     string `xml:"spefic,attr"`
}

type Permission struct {
	GasmanVmName string `xml:"gasmanvm-name,attr"`
}

type OptInOptions struct {
	Certs   []*x509.Certificate
	Domains []*DomainEntry
}

type METAINF struct {
	PackageSig     []byte
	PackageCreator *tls.Certificate
}

// Stellt das Image Reader dar
type ImageFileReader struct {
	optIn    *OptInOptions
	manifest *Manifest
	metaInf  *METAINF
	files    *ContainerFilesMap
	fileMap  map[string]*zip.File
	zipObj   *zip.Reader
	zipFile  *os.File
}
