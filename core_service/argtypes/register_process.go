package argtypes

import "github.com/fluffelpuff/GasmanVM/imagefile"

type RegisterProcessArgs struct {
	ManifestData *imagefile.Manifest
	Version      uint64
}

type RegisterVMProcessReturn struct {
	ProcessSecret string
}

type RegisterGroupProcessArgsCompleteArgs struct {
	ProcessSecret string
	Groups        map[string]string
}

type RegisterProcessArgsCompleteArgs struct {
	ProcessSecret string
}
