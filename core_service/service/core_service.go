package coreservice

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	p2pnetwork "github.com/fluffelpuff/GasmanVM/core_service/controller/p2p_network"
	"github.com/fluffelpuff/GasmanVM/core_service/utils"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multiaddr"
)

type CoreService struct {
	openProcesses             map[uint64]*ProcessSession
	functionalSharingPcoesses map[uint64]*ProcessSession
	newProcessLink            chan *ProcessSession
	sharedFunctionMap         *utils.SharedFunctionMap
	rpcController             *RPCController
	privateKeyEd25519         crypto.PrivKey
	privateSecp256k1Key       *secp256k1.PrivateKey
	p2pNetworkController      *p2pnetwork.P2PNetworkController
}

func (o *CoreService) registerRPCController(rpcController *RPCController) error {
	o.rpcController = rpcController
	return nil
}

func (o *CoreService) MainRun() error {
	return o.rpcController.MainRun()
}

func (o *CoreService) GetPublicKey() []byte {
	// Es wird ein neues Byte Slice erstellt
	nByteSlice := make([]byte, 0)

	// Der Öffentliche ED25519 Schlüssel wird hinzugefügt
	pubKeyEd25519, _ := o.privateKeyEd25519.GetPublic().Raw()
	nByteSlice = append(nByteSlice, pubKeyEd25519...)

	// Der Öffentliche Secp256k1 Schlüssel wird hinzugefügt
	nByteSlice = append(nByteSlice, o.privateSecp256k1Key.PubKey().SerializeCompressed()...)

	// Das Byte Slice wird zurückgegeben
	return nByteSlice
}

func (o *CoreService) AddSharingGroup(groupId string, groupName string) error {
	// Die Gruppe wird in der Shared Function Map hinzugefügt
	if err := o.sharedFunctionMap.AddSharingGroup(groupName, groupId); err != nil {
		return err
	}

	// Log
	fmt.Printf("Sharing group added %s %s\n", groupId, groupName)
	return nil
}

func NewCoreService(hostSeed []byte) (*CoreService, error) {
	// Erzeuge einen SHA-256-Hash des Dateninhalts
	seed256Unit := sha256.New()
	seed256Unit.Write(hostSeed)
	seed256 := seed256Unit.Sum(nil)

	// Erzeuge einen SHA-512-Hash des Dateninhalts
	seed512Unit := sha512.New()
	seed512Unit.Write(hostSeed)
	seed512 := seed512Unit.Sum(nil)

	// Der Secp256k1 Schlüssel wird abgeleitet
	privSecp256k1Key, _ := btcec.PrivKeyFromBytes(seed256)

	// Creates a new RSA key pair for this host.
	privEd25519Key, err := crypto.UnmarshalEd25519PrivateKey(seed512)
	if err != nil {
		return nil, err
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 5554))
	if err != nil {
		return nil, err
	}

	// Die Optionen für den P2P Controller werden erzeugt
	p2pOptions := []libp2p.Option{
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(privEd25519Key),
	}

	// Es wird ein neuer Netzwerk Controller erstellt
	netwController, err := p2pnetwork.NewP2PNetworkController(p2pOptions, privEd25519Key)
	if err != nil {
		return nil, err
	}

	// Das Objekt wird zurückgegeben
	return &CoreService{
		openProcesses:             map[uint64]*ProcessSession{},
		functionalSharingPcoesses: map[uint64]*ProcessSession{},
		newProcessLink:            make(chan *ProcessSession),
		sharedFunctionMap:         utils.NewSharedFunctionMap(),
		p2pNetworkController:      netwController,
		privateSecp256k1Key:       privSecp256k1Key,
		privateKeyEd25519:         privEd25519Key,
		rpcController:             nil,
	}, nil
}
