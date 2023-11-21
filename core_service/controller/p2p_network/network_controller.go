package p2pnetwork

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
)

type P2PNetworkController struct {
	libp2pHost *host.Host
}

func NewP2PNetworkController(options []config.Option, privKey crypto.PrivKey) (*P2PNetworkController, error) {
	// Es wird ein neuer LIBP2P Socket erzeugt
	lip2pHost, err := libp2p.New(options...)
	if err != nil {
		return nil, err
	}

	// Das Objekt wird zurückgegeben
	return &P2PNetworkController{libp2pHost: &lip2pHost}, nil
}
