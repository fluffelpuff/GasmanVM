package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
	"github.com/danieljoos/wincred"
	"github.com/fatih/color"
	coreservice "github.com/fluffelpuff/GasmanVM/core_service/service"
	"github.com/fluffelpuff/GasmanVM/vmpackage"
	"github.com/kardianos/service"
	"github.com/tyler-smith/go-bip39"
)

var (
	green  = color.New(color.FgHiGreen)
	fgCyan = color.New(color.FgCyan)
)

func printWelcomeText() {
	// Welcomes Textt
	text, _ := base64.StdEncoding.DecodeString("ICAgIF9fX19fICAgICAgICAgICAgICAgICAgICAgICAgICAgICBfXyAgICAgIF9fX18gIF9fICAgICANCiAgIC8gX19fX3wgICAgICAgICAgICAgICAgICAgICAgICAgICAgXCBcICAgIC8gLyAgXC8gIHwgICAgDQogIHwgfCAgX18gIF9fIF8gX19fIF8gX18gX19fICAgX18gXyBfIF9cIFwgIC8gL3wgXCAgLyB8ICAgIA0KICB8IHwgfF8gfC8gX2AgLyBfX3wgJ18gYCBfIFwgLyBfYCB8ICdfIFwgXC8gLyB8IHxcL3wgfCAgICANCiAgfCB8X198IHwgKF98IFxfXyBcIHwgfCB8IHwgfCAoX3wgfCB8IHwgXCAgLyAgfCB8ICB8IHwgICAgDQogICBcX19fX198XF9fLF98X19fL198IHxffF98X3xcX18sX3xffCB8X3xcLyAgIHxffCAgfF98ICAgIA0KICAvIF9fX198ICAgICAgICAgICAgICAgIC8gX19fX3wgICAgICAgICAgICAgICAoXykgICAgICAgICANCiB8IHwgICAgIF9fXyAgXyBfXyBfX18gIHwgKF9fXyAgIF9fXyBfIF9fX18gICBfX18gIF9fXyBfX18gDQogfCB8ICAgIC8gXyBcfCAnX18vIF8gXCAgXF9fXyBcIC8gXyBcICdfX1wgXCAvIC8gfC8gX18vIF8gXA0KIHwgfF9fX3wgKF8pIHwgfCB8ICBfXy8gIF9fX18pIHwgIF9fLyB8ICAgXCBWIC98IHwgKF98ICBfXy8NCiAgXF9fX19fXF9fXy98X3wgIFxfX198IHxfX19fXy8gXF9fX3xffCAgICBcXy8gfF98XF9fX1xfX198DQogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIA0KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA=")
	fgCyan.Println(string(text))

	// Die Version wird angezeigt
	fmt.Printf("Version: %s\n", green.Sprintf("23.11.BETA-1"))
}

func retrivePriveKey() ([]byte, error) {
	// Es wird versucht den Privaten Schlüssel abzurufen
	cred, err := wincred.GetGenericCredential(vmpackage.WIN32_CRED_PROGRAMM_NAME)
	if err != nil {
		// Erstelle ein zufälliges Mnemonic mit einer bestimmten Entropie (128, 160, 192, 224 oder 256 Bits)
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatal(err)
		}

		// Erstelle ein Mnemonic aus der Entropie
		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			log.Fatal(err)
		}

		// Erstelle einen Seed aus dem Mnemonic
		seed := bip39.NewSeed(mnemonic, "passphrase")

		// Es wird ein neuer Privatekey abgespeichert
		cred := wincred.NewGenericCredential(vmpackage.WIN32_CRED_PROGRAMM_NAME)
		cred.UserName = "GasmanVM"
		cred.CredentialBlob = seed

		// Credential im Windows Credential Manager speichern
		if err := cred.Write(); err != nil {
			return nil, err
		}

		// Der Cred wird zuröckgegeben
		return seed, nil
	}

	// Der Private Schlüssel wird zurckgegeben
	return cred.CredentialBlob, nil
}

func initServices(coreService *coreservice.CoreService, isInteractiveProcess bool) error {
	// Der P2P Controller wird gestartet
	fmt.Println("Starting Network service...")

	// Der WebRPC Controller wird gestartet
	fmt.Println("Starting WebRPC service...")

	// Der DNS Controller wird gesartet
	fmt.Println("Starting DNS service...")

	// Der Bitcoin Controller wird gestartet
	fmt.Println("Starting Bitcoin Controller service...")

	// Der Tor Addapotor wird gestartet
	fmt.Println("Starting Tor client service...")

	// Es ist kein Fehler aufgetreten
	return nil
}

func runAsInteractiveProcess(coreService *coreservice.CoreService) error {
	// Printing CoreService Informations
	printWelcomeText()

	// Der Öffentliche Schlüssel wird angezeigt
	pubKeyBase58 := base58.Encode(coreService.GetPublicKey())
	fmt.Printf("Public key: %s\n", color.GreenString(fmt.Sprintf("%s%s", vmpackage.EXPORTED_PUBLIC_KEY_PREFIC, pubKeyBase58)))

	// Der Controller wird gestartet
	if err := coreService.AddSharingGroup("cdfe0be79e9870df311eb89ebe20d9b005318e7953a98902e394ed7bb2920b2f", "internal.test.group"); err != nil {
		panic(err)
	}

	// Die Dienste werden Initialkisiert
	if err := initServices(coreService, true); err != nil {
		return err
	}

	// Der RPC Memory Controller wird ausgeführt
	return coreService.MainRun()
}

func runAsWin32Service(coreService *coreservice.CoreService, progr program) error {
	return nil
}

func main() {
	// Die Windows Dienste Einstellungen werden erzeugt
	svcConfig := &service.Config{
		Name:        "GasmanVMCoreService",
		DisplayName: "GasmanVM - CoreService",
		Description: "",
	}

	// Der Private Schlüssel des Hosts wird abgerufen
	hostSeed, err := retrivePriveKey()
	if err != nil {
		panic(err)
	}

	// Es wird ein neuer CoreService gestartet
	coreService, err := coreservice.NewCoreService(hostSeed)
	if err != nil {
		panic(err)
	}

	// Der Service Controller wird gestartet
	socketController, err := coreservice.NewCoreServiceController()
	if err != nil {
		panic(err)
	}

	// Es wird ein neuer ProcessSession Controller erstellt
	rpcController, err := coreservice.NewRPCController(socketController, coreService)
	if err != nil {
		socketController.Close()
		panic(err)
	}

	// Die Service Instanz wird erzeugt
	prg := &program{}
	_, err = service.New(prg, svcConfig)
	if err != nil {
		panic(err)
	}

	// Es wird geprüft ob der Dienst Interaktiv ausgeführt wird
	if service.Interactive() {
		if err := runAsInteractiveProcess(coreService); err != nil {
			panic(err)
		}
	} else {
		if err := runAsWin32Service(coreService, *prg); err != nil {
			panic(err)
		}
	}

	// Wird ausgeführt wenn das Programm geschlossen werden soll
	defer func() {
		// Der SocketController wird geschlossen
		socketController.Close()

		// Der ProcessSession Controller wird geschlossen
		rpcController.Close()
	}()
}
