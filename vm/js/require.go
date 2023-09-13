package jsengine

import (
	"fmt"
	"log"
	"strings"

	"github.com/dop251/goja"
)

// Wird ausgeführt um ein JS basiertes Modul zu Impelntieren
func (o *JSEngine) runtimeRequire(call goja.FunctionCall, jsruntime *goja.Runtime) goja.Value {
	// Holen Sie den Dateinamen, der in require übergeben wurde
	if len(call.Arguments) < 1 {
		log.Print("Fehler: Kein Dateiname in require übergeben")
		return goja.Null()
	}

	// Der Dateiname wird extrahiert
	filename := call.Argument(0).String()

	// Es wird ermittelt um was für einen Import es sich handelt
	if strings.HasPrefix(filename, "./") { // Es handelt sich um einen Dateibasierten Import
		// Suchen Sie die Datei im ZIP-Archiv
		img_file := o.motherVM.GetImageFile()

		// Die Datei wird ermittelt
		file_content, err := img_file.GetSourceFile(filename[2:])
		if err != nil {
			return jsruntime.NewGoError(err)
		}

		// Es wird ermittelt ob es sich um eine Javascript Datei handelt
		// if file_content.ScripteType != "js" {}

		// Das Scirpt wird ausgeführt
		moduleVM, err := o.runModuleJSScript(file_content.GetBytes())
		if err != nil {
			return goja.New().NewGoError(err)
		}

		// Die Exports werden zurückgegeben
		return moduleVM.Get("exports")
	} else if strings.HasPrefix(filename, "@") { // Es handelt es um einen Native basierten Import
		return nil
	} else { // Es handelt sich um ein OOB Import
		switch filename {
		// Stellt die Standard NodeJS Funktionen bereit
		case "crypto": // NodeJS Standard + Erweiterungen
			return o.loadCryptoModule(jsruntime)
		case "dns": // NodeJS Standard
			return o.loadDNSModule(jsruntime)
		case "timers": // NodeJS Standard
			return o.loadTimersModule(jsruntime)
		case "url": // NodeJS Standard
			return o.loadURLModule(jsruntime)
		case "https": // NodeJS Standard
			return o.loadHTTPSModule(jsruntime)
		case "zlib": // NodeJS Standard
			return o.loadZLibModule(jsruntime)
		case "buffer": // NodeJS Standard
			return o.loadBufferModule(jsruntime)
		case "vm": // NodeJS Standard + Erweiterungen
			return o.loadVMModule(jsruntime)
		case "path": // NodeJS Standard
			return o.loadPathModule(jsruntime)
		case "fs": // Stellt Dateisystem Funktionen bereit (Achtung, es besteht keine möglichkeit auf das Lokale Dateisystem zugreifen zu können)
			return o.loadFileSystemModule(jsruntime)
		// Spizielle Funktionen
		case "websocket": // Stellt Websocket Funktionen bereit
			return o.loadWebsocketModule(jsruntime)
		case "apisockets": // Stellt API Socket Funktionen bereit
			return o.loadAPISocketsModule(jsruntime)
		case "ssh": // Stellt SSH-Client Funktionen bereit
			return o.loadSSHClientModule(jsruntime)
		case "ipc": // Stellt IPC Funktionen breit
			return o.loadIPCModule(jsruntime)
		case "bip": // Stellt die Bitcoin Funktionen bereit
			return o.loadBIPModule(jsruntime)
		case "sql": // Stellt SQL Funktionen bereit
			return o.loadSQLDBModule(jsruntime)
		case "native": // Stellt die Möglichkeit bereit Native Libs zu verwenden
			return o.loadNativeModule(jsruntime)
		case "encoding_decoding": // Stellt alle Encoding/Decoding Funktionen bereit
			return o.loadEncodingDecodingModule(jsruntime)
		case "webdav": // Stellt WebDav Client Funktionen bereit
			return o.loadingWebDavModule(jsruntime)
		case "git": // Stellt Funktionen für Github bereit
			return o.loadingGitModule(jsruntime)
		case "nostr": // Stellt die Nostr Funktionen bereit
			return o.loadNostrModule(jsruntime)
		case "vmrpc": // Stellt die VM RPC Funktionen bereit
			return o.loadingVMRPCModule(jsruntime)
		default:
			return goja.New().NewGoError(fmt.Errorf("unkown import"))
		}
	}
}
