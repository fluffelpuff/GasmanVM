package jsengine

import (
	"github.com/dop251/goja"
	apisockets "github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/apisocket"
	bipmodule "github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/bip"
	cryptoNoneNodeJs "github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/crypto"
	encodingdecoding "github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/encoding_decoding"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/fs"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/nostr"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/sql"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/ssh"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/webdav"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/websocket"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/buffer"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/crypto"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/dns"
	nodejsfs "github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/fs"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/https"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/path"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/timers"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/url"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/vm"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/zlib"
)

// Stellt die Crypto Funktionen bereit
func (o *JSEngine) loadCryptoModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	crypto_obj := jsruntime.NewObject()

	// Stellt die Standard NodeJS Crypto Funktionen bereit
	crypto_obj.Set("createCipher", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createCipher(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createCipheriv", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createCipheriv(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createDecipher", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createDecipher(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createDecipheriv", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createDecipheriv(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createDiffieHellman", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createDiffieHellman(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createECDH", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createECDH(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createHash", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createHash(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createHmac", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createHmac(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createSign", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createSign(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("createVerify", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_createVerify(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("getCurves", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_getCurves(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("getDiffieHellman", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_getDiffieHellman(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("getHashes", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_getHashes(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("pbkdf2", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_pbkdf2(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("pbkdf2Sync", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_pbkdf2Sync(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("privateDecrypt", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_privateDecrypt(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("timingSafeEqual", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_timingSafeEqual(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("privateEncrypt", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_privateEncrypt(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("publicDecrypt", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_publicDecrypt(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("publicEncrypt", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_publicEncrypt(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("randomBytes", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_randomBytes(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("setEngine", func(parms goja.FunctionCall) goja.Value {
		return crypto.Module_Crypto_setEngine(o.motherVM, jsruntime, parms)
	})

	// Stellt Hash Funktionen bereit
	hashModule := jsruntime.NewObject()
	crypto_obj.Set("hash", hashModule)
	crypto_obj.Set("computeSha2", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_sha2_compute(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("computeSha3", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_sha3_compute(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("computeKeccak", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_keccak_compute(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("SHA2", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_keccak_compute(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("SHA3", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_keccak_compute(o.motherVM, jsruntime, parms)
	})
	crypto_obj.Set("Keccak", func(parms goja.FunctionCall) goja.Value {
		return cryptoNoneNodeJs.Module_crypto_keccak_compute(o.motherVM, jsruntime, parms)
	})

	// Stellt ECC Funktionen bereit
	eccModule := jsruntime.NewObject()
	crypto_obj.Set("ecc", eccModule)

	// Stellt ED Funktionen bereit
	edModule := jsruntime.NewObject()
	crypto_obj.Set("ed", edModule)

	// Stellt PGP Funktionen bereit
	pgpModule := jsruntime.NewObject()
	crypto_obj.Set("pgp", pgpModule)

	// Stellt DEDIS Funktionen bereit
	dedisModule := jsruntime.NewObject()
	crypto_obj.Set("dedis", dedisModule)

	// Stellt Kyber-K2SO Funktionen bereit
	kyberk2s0Module := jsruntime.NewObject()
	crypto_obj.Set("kyberk2so", kyberk2s0Module)

	// Das Module wird zurückgegeben
	return crypto_obj
}

// Stellt die DNS Funktionen bereit
func (o *JSEngine) loadDNSModule(jsruntime *goja.Runtime) goja.Value {
	// Das DNS Modul Objekt wird erstellt
	dnsModule := jsruntime.NewObject()
	dnsModule.Set("getServers", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_getServers(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("lookup", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_lookup(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("lookupService", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_lookupService(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolve", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolve(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolve4", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolve4(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolve6", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolve6(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveCname", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveCname(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveMx", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveMx(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveNaptr", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveNaptr(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveNs", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveNs(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveSoa", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveSoa(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveSrv", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveSrv(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolvePtr", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolvePtr(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("resolveTxt", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_resolveTxt(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("reverse", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_reverse(o.motherVM, jsruntime, parms)
	})
	dnsModule.Set("setServers", func(parms goja.FunctionCall) goja.Value {
		return dns.Module_DNS_setServers(o.motherVM, jsruntime, parms)
	})

	// Das Objekt wird zurückgegeben
	return dnsModule
}

// Stellt die Timer Funktionen bereit
func (o *JSEngine) loadTimersModule(jsruntime *goja.Runtime) goja.Value {
	// Das Timer Modul wird erezugt
	timersModule := jsruntime.NewObject()
	timersModule.Set("clearImmediate", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_clearImmediate(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("clearInterval", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_clearInterval(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("clearTimeout", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_clearTimeout(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("ref", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_ref(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("setImmediate", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_setImmediate(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("setInterval", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_setInterval(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("setTimeout", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_setTimeout(o.motherVM, jsruntime, parms)
	})
	timersModule.Set("unref", func(parms goja.FunctionCall) goja.Value {
		return timers.Module_Timers_unref(o.motherVM, jsruntime, parms)
	})

	// Das Modul wird zurückgegeben
	return timersModule
}

// Stellt die URL Funktionen bereit
func (o *JSEngine) loadURLModule(jsruntime *goja.Runtime) goja.Value {
	// Das URL Module wird erezugt
	urlModule := jsruntime.NewObject()
	urlModule.Set("format", func(parms goja.FunctionCall) goja.Value {
		return url.Module_URL_format(o.motherVM, jsruntime, parms)
	})
	urlModule.Set("parse", func(parms goja.FunctionCall) goja.Value {
		return url.Module_URL_parse(o.motherVM, jsruntime, parms)
	})
	urlModule.Set("resolve", func(parms goja.FunctionCall) goja.Value {
		return url.Module_URL_resolve(o.motherVM, jsruntime, parms)
	})

	// Das Modul wird zurückgegeben
	return urlModule
}

// Stellt alle Funktionen für einen HTTPS Server bereit
func (o *JSEngine) loadHTTPSModule(jsruntime *goja.Runtime) goja.Value {
	// Das HTTPS Modul wird erezgut
	httpsModule := jsruntime.NewObject()
	httpsModule.Set("createServer", func(parms goja.FunctionCall) goja.Value {
		return https.Module_https_createServer(o.motherVM, jsruntime, parms)
	})
	httpsModule.Set("get", func(parms goja.FunctionCall) goja.Value {
		return https.Module_https_get(o.motherVM, jsruntime, parms)
	})
	httpsModule.Set("globalAgent", func(parms goja.FunctionCall) goja.Value {
		return https.Module_https_globalAgent(o.motherVM, jsruntime, parms)
	})
	httpsModule.Set("request", func(parms goja.FunctionCall) goja.Value {
		return https.Module_https_request(o.motherVM, jsruntime, parms)
	})

	// Das Module wird zurückgegeben
	return httpsModule
}

// Stellt die ZLib Funktionen bereit
func (o *JSEngine) loadZLibModule(jsruntime *goja.Runtime) goja.Value {
	// Das zlib Modul wird erstellt
	zlibMopdule := jsruntime.NewObject()
	zlibMopdule.Set("createDeflate", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createDeflate(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("createDeflateRaw", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createDeflateRaw(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("createGunzip", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createGunzip(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("createGzip", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createGzip(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("createInflate", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createInflate(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("createInflateRaw", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_createInflateRaw(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("deflate", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_deflate(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("deflateSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_deflateSync(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("deflateRaw", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_deflateRaw(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("deflateRawSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_deflateRawSync(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("gunzip", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_gunzip(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("gunzipSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_gunzipSync(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("gzip", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_gzip(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("gzipSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_gzipSync(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("inflate", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_inflate(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("inflateSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_inflateSynce(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("inflateRaw", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_inflateRaw(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("inflateRawSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_inflateRawSync(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("unzip", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_unzip(o.motherVM, jsruntime, parms)
	})
	zlibMopdule.Set("unzipSync", func(parms goja.FunctionCall) goja.Value {
		return zlib.Module_ZLIB_unzipSync(o.motherVM, jsruntime, parms)
	})

	// Das Module wird zurückgegeben
	return zlibMopdule
}

// Stellt VM Funktionen bereit
func (o *JSEngine) loadVMModule(jsruntime *goja.Runtime) goja.Value {
	// Dasa VM Modul wird erzeugt
	vmModule := jsruntime.NewObject()
	vmModule.Set("createContext", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_createContext(o.motherVM, jsruntime, parms)
	})
	vmModule.Set("isContext", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_runInContext(o.motherVM, jsruntime, parms)
	})
	vmModule.Set("runInContext", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_runInContext(o.motherVM, jsruntime, parms)
	})
	vmModule.Set("runInDebug", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_runInDebug(o.motherVM, jsruntime, parms)
	})
	vmModule.Set("runInNewContext", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_runInNewContext(o.motherVM, jsruntime, parms)
	})
	vmModule.Set("runInThisContext", func(parms goja.FunctionCall) goja.Value {
		return vm.Module_VM_runInThisContext(o.motherVM, jsruntime, parms)
	})

	// Gibt das Module zurück
	return vmModule
}

// Stellt die Standard NodeJS Path Funktionen bereit
func (o *JSEngine) loadPathModule(jsruntime *goja.Runtime) goja.Value {
	// Das Path Modul wird erstellt
	pathModule := jsruntime.NewObject()
	pathModule.Set("format", func(parms goja.FunctionCall) goja.Value {
		return path.Module_PATH_format(o.motherVM, jsruntime, parms)
	})
	pathModule.Set("parse", func(parms goja.FunctionCall) goja.Value {
		return path.Module_PATH_parse(o.motherVM, jsruntime, parms)
	})
	pathModule.Set("resolve", func(parms goja.FunctionCall) goja.Value {
		return path.Module_PATH_resolve(o.motherVM, jsruntime, parms)
	})

	// Gibt das Module zurück
	return pathModule
}

// Stellt Websocket Funktionen bereit
func (o *JSEngine) loadWebsocketModule(jsruntime *goja.Runtime) goja.Value {
	// Das Websocket Modul wird erstellt
	websocketModule := jsruntime.NewObject()
	websocketModule.Set("newServer", func(parms goja.FunctionCall) goja.Value {
		return websocket.Module_websocket_newServer(o.motherVM, jsruntime, parms)
	})
	websocketModule.Set("connect", func(parms goja.FunctionCall) goja.Value {
		return websocket.Module_websocket_connect(o.motherVM, jsruntime, parms)
	})

	// Das Module wird zurückgegeben
	return websocketModule
}

// Stellt APISocket Funktionen bereit
func (o *JSEngine) loadAPISocketsModule(jsruntime *goja.Runtime) goja.Value {
	// Das API Socket Modul wird erzeugt
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("newServer", func(parms goja.FunctionCall) goja.Value {
		return apisockets.Module_apisockets_newServer(o.motherVM, jsruntime, parms)
	})
	apiSocketsModule.Set("connect", func(parms goja.FunctionCall) goja.Value {
		return apisockets.Module_websocket_connect(o.motherVM, jsruntime, parms)
	})

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die Nostr Funktionen bereit
func (o *JSEngine) loadNostrModule(jsruntime *goja.Runtime) goja.Value {
	// Das Nostr Modul wird erzeugt
	nostrModule := jsruntime.NewObject()
	nostrModule.Set("newRelayPoolSync", func(parms goja.FunctionCall) goja.Value {
		return nostr.Module_nostr_newRelayPool(o.motherVM, jsruntime, parms)
	})
	nostrModule.Set("newRelayPool", func(parms goja.FunctionCall) goja.Value {
		return nostr.Module_nostr_newRelayPool(o.motherVM, jsruntime, parms)
	})

	// das Module wird zurückgegeben
	return nostrModule
}

// Stellt die Dateisystem Funktionen bereit
func (o *JSEngine) loadFileSystemModule(jsruntime *goja.Runtime) goja.Value {
	// Das Dateisystem für nodejsfs Vorgänge wird erzeugt
	fileSystemSyncModule := jsruntime.NewObject()

	// Wird verwendet um die Zugriffsberechtigung für eine Datei oder einen Ordner zu überprüfen
	fileSystemSyncModule.Set("accessSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_accessSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("access", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_accessCallback(o.motherVM, jsruntime, parms)
	})

	// Fügt einen Datensatz zu einer Datei hinzu
	fileSystemSyncModule.Set("appendFileSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_appendFileSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("appendFile", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_appendFileCallback(o.motherVM, jsruntime, parms)
	})

	// Legt die Berechtigung für eine Datei fest
	fileSystemSyncModule.Set("chmodSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_chmodSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("chmod", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_chmodCallback(o.motherVM, jsruntime, parms)
	})

	// Wird verwendet um den Besitzer einer Datei festzulegen
	fileSystemSyncModule.Set("chownSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_chownSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("chown", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_chownCallback(o.motherVM, jsruntime, parms)
	})

	// Wird verwendet um einen Dateiskriptor zu schließen
	fileSystemSyncModule.Set("closeSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_closeSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("close", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_closeCallback(o.motherVM, jsruntime, parms)
	})

	// Diese Funktion kopiert eine Datei
	fileSystemSyncModule.Set("copyFileSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_copyFileSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("copyFile", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_copyFileCallback(o.motherVM, jsruntime, parms)
	})

	// Gibt an on eine Datei vorhanden ist
	fileSystemSyncModule.Set("existsSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_existsSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("exists", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_existsCallback(o.motherVM, jsruntime, parms)
	})

	// Lese-Stream-Instanz für das Lesen von Daten aus einer Datei
	fileSystemSyncModule.Set("createReadStream", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_createReadStream(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("createWriteStream", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_createWriteStream(o.motherVM, jsruntime, parms)
	})

	// Diese Funktion erstellt einen neuen Ordner
	fileSystemSyncModule.Set("mkdirSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_mkdirSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("mkdir", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_mkdirCallback(o.motherVM, jsruntime, parms)
	})

	// Öffnet eine Datei
	fileSystemSyncModule.Set("openSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_openSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("open", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_openCallback(o.motherVM, jsruntime, parms)
	})

	// Ließt den Inahlt eines Ordner ein
	fileSystemSyncModule.Set("readdirSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_readdirSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("readdir", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_readdirCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktionen zum einlesen einer Datei
	fileSystemSyncModule.Set("readFileSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_readFileSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("readFile", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_readFileCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktionen benenenen eine Funktion um
	fileSystemSyncModule.Set("renameSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_renameSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("rename", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_renameCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktionen zum löschen eines Ordners
	fileSystemSyncModule.Set("rmdirSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_rmdirSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("rmdir", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_rmdirCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktion wird zum Löschen einer Datei verwendet
	fileSystemSyncModule.Set("rmSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_rmdirSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("rm", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_rmdirCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktionen zum Ausgeben der Metadaten
	fileSystemSyncModule.Set("statSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_statSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("stat", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_statCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktion erstellt einen Symlink
	fileSystemSyncModule.Set("symlinkSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_symlinkSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("symlink", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_symlinkCallback(o.motherVM, jsruntime, parms)
	})

	// Wird verwendet um eine Verlinkung zu entfernen
	fileSystemSyncModule.Set("unlinkSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_unlinkSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("unlink", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_unlinkCallback(o.motherVM, jsruntime, parms)
	})

	// Die Funktion wird verwendet, um die Zugriffs- und Änderungszeiten einer Datei festzulegen.
	fileSystemSyncModule.Set("utimesSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("utimes", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesCallback(o.motherVM, jsruntime, parms)
	})

	// Wird verwendet um eine Datei zu schreiben
	fileSystemSyncModule.Set("writeFileSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("writeFile", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesCallback(o.motherVM, jsruntime, parms)
	})

	// Verwendet Daten in eine Datei zu schreiben. Sie blockiert den Ausführungsthread, bis der Schreibvorgang abgeschlossen ist.
	fileSystemSyncModule.Set("writeSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("write", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesCallback(o.motherVM, jsruntime, parms)
	})

	// WritevSync ist eine Methode, die mehrere Puffer (Buffers) in eine geöffnete Datei schreibt, wobei sie synchron arbeitet, was bedeutet, dass sie den Ausführungsthread blockiert, bis der Schreibvorgang abgeschlossen ist.
	fileSystemSyncModule.Set("writevSync", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("writev", func(parms goja.FunctionCall) goja.Value {
		return nodejsfs.Module_FS_SYNC_utimesCallback(o.motherVM, jsruntime, parms)
	})

	/*
	 Die Gasman Modul Funktionen werden hinzugefügt
	*/

	// Stellt eine IMG Datei bereit
	fileSystemSyncModule.Set("mountImgSync", func(parms goja.FunctionCall) goja.Value {
		return fs.Module_FS_SYNC_mountimgSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("mountImg", func(parms goja.FunctionCall) goja.Value {
		return fs.Module_FS_SYNC_mountimgCallback(o.motherVM, jsruntime, parms)
	})

	// Stellt eine ISO Datei bereit
	fileSystemSyncModule.Set("mountIsoSync", func(parms goja.FunctionCall) goja.Value {
		return fs.Module_FS_SYNC_mountisoSync(o.motherVM, jsruntime, parms)
	})
	fileSystemSyncModule.Set("mountIso", func(parms goja.FunctionCall) goja.Value {
		return fs.Module_FS_SYNC_mountisoCallback(o.motherVM, jsruntime, parms)
	})

	// das Module wird zurückgegeben
	return fileSystemSyncModule
}

// Stellt Funktionen für die Verwaltung von Lokalen SQLite oder Lokalen MySQL Datenbanekn bereit
func (o *JSEngine) loadSQLDBModule(jsruntime *goja.Runtime) goja.Value {
	// Die SQL Funktionen werden geschrieben
	sqlModule := jsruntime.NewObject()
	sqlModule.Set("openFile", func(parms goja.FunctionCall) goja.Value {
		return sql.Module_SQL_openFile(o.motherVM, jsruntime, parms)
	})
	sqlModule.Set("connectTo", func(parms goja.FunctionCall) goja.Value {
		return sql.Module_SQL_connectTo(o.motherVM, jsruntime, parms)
	})

	// Gibt das Module zurück
	return sqlModule
}

// Stellt die Webdav Funktionen bereit
func (o *JSEngine) loadingWebDavModule(jsruntime *goja.Runtime) goja.Value {
	// Das WebDav Modul wird erzeugt
	webDavModule := jsruntime.NewObject()
	webDavModule.Set("NewServer", func(parms goja.FunctionCall) goja.Value {
		return webdav.Module_Webdav_newServer(o.motherVM, jsruntime, parms)
	})
	webDavModule.Set("Connect", func(parms goja.FunctionCall) goja.Value {
		return webdav.Module_Webdav_connectTo(o.motherVM, jsruntime, parms)
	})

	// das Module wird zurückgegeben
	return webDavModule
}

// Stellt alle SSH Client Funktionen bereit
func (o *JSEngine) loadSSHClientModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	sshModule := jsruntime.NewObject()
	sshModule.Set("connectTo", func(parms goja.FunctionCall) goja.Value {
		return ssh.Module_SSH_connectTo(o.motherVM, jsruntime, parms)
	})

	// Gibt das Module zurück
	return sshModule
}

// Stellt die Encoding und Decoding Funktionen bereit
func (o *JSEngine) loadEncodingDecodingModule(jsruntime *goja.Runtime) goja.Value {
	// Erzeugt das UTF Mdoul
	utfModule := jsruntime.NewObject()
	utfModule.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_utf_encodeToString(o.motherVM, jsruntime, parms)
	})
	utfModule.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_utf_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das JSON Modul wird erezugt
	jsonModule := jsruntime.NewObject()
	jsonModule.Set("encodeToByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_utf_decodeFromString(o.motherVM, jsruntime, parms)
	})
	jsonModule.Set("decodeFromByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_json_decodeFromByteArray(o.motherVM, jsruntime, parms)
	})
	jsonModule.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_json_encodeToString(o.motherVM, jsruntime, parms)
	})
	jsonModule.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_json_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das XML Modul wird erzeugt
	xmlModule := jsruntime.NewObject()
	xmlModule.Set("encodeToByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_xml_encodeToByteArray(o.motherVM, jsruntime, parms)
	})
	xmlModule.Set("decodeFromByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_xml_decodeFromByteArray(o.motherVM, jsruntime, parms)
	})
	xmlModule.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_xml_encodeToString(o.motherVM, jsruntime, parms)
	})
	xmlModule.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_xml_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das CBOR Modul wird erezeugt
	cborModule := jsruntime.NewObject()
	cborModule.Set("encodeToByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_cbor_encodeToByteArray(o.motherVM, jsruntime, parms)
	})
	cborModule.Set("decodeFromByteArray", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_cbor_decodeFromByteArray(o.motherVM, jsruntime, parms)
	})

	// Das Hex Modul wird erzeugt
	hexModule := jsruntime.NewObject()
	hexModule.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_hex_encodeToString(o.motherVM, jsruntime, parms)
	})
	hexModule.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_hex_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das PEM Modul wird erezugt
	pemModule := jsruntime.NewObject()
	pemModule.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_pem_encodeToString(o.motherVM, jsruntime, parms)
	})
	pemModule.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_pem_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das Base32 Modul wird erezugt
	base32Module := jsruntime.NewObject()
	base32Module.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base32_encodeToString(o.motherVM, jsruntime, parms)
	})
	base32Module.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base32_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das Base58 Modul wird erezeugt
	base58Module := jsruntime.NewObject()
	base58Module.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base32_encodeToString(o.motherVM, jsruntime, parms)
	})
	base58Module.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base32_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das Base64 Modul wird erzeugt
	base64Module := jsruntime.NewObject()
	base64Module.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base64_encodeToString(o.motherVM, jsruntime, parms)
	})
	base64Module.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_base64_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Das Bech32 Modul wird erzeugt
	bech32Module := jsruntime.NewObject()
	bech32Module.Set("encodeToString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_bech32_encodeToString(o.motherVM, jsruntime, parms)
	})
	bech32Module.Set("decodeFromString", func(parms goja.FunctionCall) goja.Value {
		return encodingdecoding.Module_encodingdecoding_bech32_decodeFromString(o.motherVM, jsruntime, parms)
	})

	// Dsa Encoding Decoding Mdoule wird erzeugt
	encoding_decodingModule := jsruntime.NewObject()
	encoding_decodingModule.Set("utf", utfModule)
	encoding_decodingModule.Set("json", jsonModule)
	encoding_decodingModule.Set("xml", xmlModule)
	encoding_decodingModule.Set("cbor", cborModule)
	encoding_decodingModule.Set("hex", hexModule)
	encoding_decodingModule.Set("pem", pemModule)
	encoding_decodingModule.Set("base32", base32Module)
	encoding_decodingModule.Set("base58", base58Module)
	encoding_decodingModule.Set("base64", base64Module)
	encoding_decodingModule.Set("bech32", bech32Module)

	// das Module wird zurückgegeben
	return encoding_decodingModule
}

// Stellt BIP Funktionen sowie Electrum Funktionen bereit
func (o *JSEngine) loadBIPModule(jsruntime *goja.Runtime) goja.Value {
	// Erstellt ein Objekt welches die Module bereitstellt
	bipModule := jsruntime.NewObject()
	bipModule.Set("addnode", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_addnode(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("createrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_createrawtransaction(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("decoderawtransaction", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_decoderawtransaction(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("decodescript", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_decodescript(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("estimatefee", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_estimatefeet(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getaddednodeinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getaddednodeinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getbestblock", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getbestblock(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getbestblockhash", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getbestblockhash(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getblock", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getblock(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getblockchaininfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getblockchaininfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getblockcount", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getblockcount(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getblockhash", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getblockhash(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getblockheader", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getblockheader(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getchaintips", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getchaintips(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getconnectioncount", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getconnectioncount(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getdifficulty", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getdifficulty(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getmempoolinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getmempoolinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getmininginfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getmininginfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getnettotals", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getnettotals(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getnetworkinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getnetworkinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getpeerinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getpeerinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getrawmempool", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getrawmempool(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("getrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_getrawtransaction(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("gettxout", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_gettxout(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("gettxoutproof", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_gettxoutproof(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("gettxoutsetinfo", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_gettxoutsetinfo(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("listbanned", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_listbanned(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("ping", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_ping(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("sendrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_sendrawtransaction(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("stop", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_stop(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("validateaddress", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_validateaddress(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("verifychain", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_verifychain(o.motherVM, jsruntime, parms)
	})
	bipModule.Set("verifymessage", func(parms goja.FunctionCall) goja.Value {
		return bipmodule.Module_BIP_verifymessage(o.motherVM, jsruntime, parms)
	})

	// Das Module wird zurückgegeben
	return bipModule
}

// Stellt die Standard NodeJS Buffer Funktionen zur verfügung
func (o *JSEngine) loadBufferModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	bufferModule := jsruntime.NewObject()
	bufferModule.Set("alloc", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_alloc(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("allocUnsafe", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_allocUnsafe(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("byteLength", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_byteLength(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("compare", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_compare(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("concat", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_concat(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("copy", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_copy(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("entries", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_entries(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("equals", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_equals(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("fill", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_fill(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("from", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_from(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("includes", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_includes(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("indexOf", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_indexOf(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("isBuffer", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_isBuffer(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("isEncoding", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_isEncoding(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("keys", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_keys(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("lastIndexOf", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_lastIndexOf(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readDoubleBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readDoubleBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readDoubleLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readDoubleLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readFloatBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readFloatBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readFloatLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readFloatLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readInt8", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readInt8(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readInt16BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readInt16BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readInt16LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readInt16LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readInt32BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readInt32BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readInt32LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readInt32LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readIntBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readIntBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readIntLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readIntLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUInt8", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUInt8(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUInt16BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUInt16BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUInt16LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUInt16LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUInt32BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUInt32BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUInt32LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUInt32LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUintBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUintBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("readUIntLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_readUIntLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("slice", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_slice(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("swap16", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_swap16(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("swap32", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_swap32(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("swap64", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_swap64(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("toJSON", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_toJSON(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("values", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_values(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("write", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_write(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeDoubleBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeDoubleBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeFloatBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeFloatBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeFloatLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeFloatLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeInt8", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeInt8(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeInt16BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeInt16BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeInt16LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeInt16LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeInt32BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeInt32BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeInt32LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeInt32LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeIntBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeIntBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeIntLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeIntLE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUInt8", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUInt8(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUInt16BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUInt16BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUInt16LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUInt16LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUInt32BE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUInt32BE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUInt32LE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUInt32LE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUIntBE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUIntBE(o.motherVM, jsruntime, parms)
	})
	bufferModule.Set("writeUIntLE", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_writeUIntLE(o.motherVM, jsruntime, parms)
	})

	// Gibt das Module zurück
	return bufferModule
}

// Stellt die Bitcoin sowie Lightning Funktionen bereit
func (o *JSEngine) loadBitcoinLightningModule(jsruntime *goja.Runtime) goja.Value {
	// Stellt die Bitcoin Funktionen bereit
	bitcoinModule := jsruntime.NewObject()
	bitcoinModule.Set("openWalletFile", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_alloc(o.motherVM, jsruntime, parms)
	})
	bitcoinModule.Set("openWallet", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_alloc(o.motherVM, jsruntime, parms)
	})
	bitcoinModule.Set("newWallet", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_alloc(o.motherVM, jsruntime, parms)
	})
	bitcoinModule.Set("getMethode", func(parms goja.FunctionCall) goja.Value {
		return buffer.Module_Buffer_alloc(o.motherVM, jsruntime, parms)
	})

	// Stellt die Lightning Funktionen bereit
	lightningModule := jsruntime.NewObject()

	// Das Bitcoin Lightning Modul wird erstellt
	bitcoinLightningModule := jsruntime.NewObject()
	bitcoinLightningModule.Set("btc", bitcoinModule)
	bitcoinLightningModule.Set("ln", lightningModule)

	// das Module wird zurückgegeben
	return bitcoinLightningModule
}
