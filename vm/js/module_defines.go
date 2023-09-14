package jsengine

import (
	"github.com/dop251/goja"
	apisockets "github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/apisocket"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/nostr"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/gasmanvm/websocket"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/crypto"
	"github.com/fluffelpuff/GasmanVM/vm/js/modules/nodejs/dns"
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

	// Stellt zuzäzliche Post-quantum cryptography funktionen bereit
	poscryp := jsruntime.NewObject()
	crypto_obj.Set("pqc", poscryp)

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

// Stellt die Nativen Funktionen bereit
func (o *JSEngine) loadNativeModule(jsruntime *goja.Runtime) goja.Value {
	// Das Modul für Native Libs wird erstellt
	nativeLibModule := jsruntime.NewObject()
	nativeLibModule.Set("loadLibrary", func(parms goja.FunctionCall) goja.Value {
		return o.loadNativeLibrary(parms, jsruntime)
	})

	// das Module wird zurückgegeben
	return nativeLibModule
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
	nostrModule.Set("newRelayPool", func(parms goja.FunctionCall) goja.Value {
		return nostr.Module_nostr_newRelayPool(o.motherVM, jsruntime, parms)
	})

	// das Module wird zurückgegeben
	return nostrModule
}

// Stellt die VMRPC Funktionen bereit
func (o *JSEngine) loadingVMRPCModule(jsruntime *goja.Runtime) goja.Value {
	// Das VMRPC Modul wird erzeugt
	vmRpcModule := jsruntime.NewObject()
	vmRpcModule.Set("openChannel", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	vmRpcModule.Set("closeChannel", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	vmRpcModule.Set("writeChannel", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	vmRpcModule.Set("readChannel", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// das Module wird zurückgegeben
	return vmRpcModule
}

// Stellt die Bitcoin sowie Lightning Funktionen bereit
func (o *JSEngine) loadBitcoinLightningModule(jsruntime *goja.Runtime) goja.Value {
	// Stellt die Bitcoin Funktionen bereit
	bitcoinModule := jsruntime.NewObject()

	// Stellt die Lightning Funktionen bereit
	lightningModule := jsruntime.NewObject()

	// Das Bitcoin Lightning Modul wird erstellt
	bitcoinLightningModule := jsruntime.NewObject()
	bitcoinLightningModule.Set("btc", bitcoinModule)
	bitcoinLightningModule.Set("ln", lightningModule)

	// das Module wird zurückgegeben
	return bitcoinLightningModule
}

// Stellt BIP Funktionen sowie Electrum Funktionen bereit
func (o *JSEngine) loadBIPModule(jsruntime *goja.Runtime) goja.Value {
	// Stellt die Blockchain Funktionen bereit
	blockchain_object := jsruntime.NewObject()
	blockchain_object.Set("getbestblockhash", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblock", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockchaininfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockcount", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockfilter", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockhash", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockheader", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getblockstats", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getchaintips", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getchaintxstats", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getdifficulty", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getmempoolancestors", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getmempooldescendants", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getmempoolentry", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getmempoolinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("getrawmempool", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("gettxout", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("gettxoutproof", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("gettxoutsetinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("gettxoutproof", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("preciousblock", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("pruneblockchain", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("savemempool", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("scantxoutset", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("verifychain", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("scantxoutset", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	blockchain_object.Set("verifytxoutproof", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Stellt die Netzwerk Funktionen bereit
	network_object := jsruntime.NewObject()
	network_object.Set("addnode", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("clearbanned", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("disconnectnode", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getaddednodeinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getconnectioncount", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getnettotals", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getnetworkinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getnodeaddresses", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("getpeerinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("listbanned", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("ping", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("setban", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	network_object.Set("setnetworkactive", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Das Objekt für die Transkationsfunktionen wird erzeugt
	transactions_object := jsruntime.NewObject()
	transactions_object.Set("analyzepsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("combinepsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("combinerawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("converttopsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("createpsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("createrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("decodepsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("decoderawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("decodescript", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("finalizepsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("fundrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("getrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("joinpsbts", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("sendrawtransaction", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("signrawtransactionwithkey", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("testmempoolaccept", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	transactions_object.Set("utxoupdatepsbt", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Das Objekt für die Uitls wird erzeugt
	utils_oject := jsruntime.NewObject()
	utils_oject.Set("createmultisig", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("deriveaddresses", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("estimatesmartfee", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("getdescriptorinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("getindexinfo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("signmessagewithprivkey", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("validateaddress", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	utils_oject.Set("verifymessage", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Das Objekt für die Wallet Funktionen wird erzeugt
	wallet_object := jsruntime.NewObject()
	wallet_object.Set("new", nil)

	// Die Funktionen werden hinzugefügt
	bipCoreModule := jsruntime.NewObject()
	bipCoreModule.Set("blockchain", blockchain_object)
	bipCoreModule.Set("network", network_object)
	bipCoreModule.Set("transactions", transactions_object)
	bipCoreModule.Set("utils", utils_oject)
	bipCoreModule.Set("wallet", wallet_object)

	// Das Modul für die Electrum Funktionen wird angelegt
	electrumModule := jsruntime.NewObject()

	// Das Finale Modul wird erstellt
	bipModule := jsruntime.NewObject()
	bipModule.Set("bipcore", bipCoreModule)
	bipModule.Set("electrum", electrumModule)

	// Das Module wird zurückgegeben
	return bipModule
}

// Stellt die Encoding und Decoding Funktionen bereit
func (o *JSEngine) loadEncodingDecodingModule(jsruntime *goja.Runtime) goja.Value {
	// Erzeugt das UTF Mdoul
	utfModule := jsruntime.NewObject()

	// Das JSON Modul wird erezugt
	jsonModule := jsruntime.NewObject()

	// Das XML Modul wird erzeugt
	xmlModule := jsruntime.NewObject()

	// Das CBOR Modul wird erezeugt
	cborModule := jsruntime.NewObject()

	// Das Hex Modul wird erzeugt
	hexModule := jsruntime.NewObject()

	// Das PEM Modul wird erezugt
	pemModule := jsruntime.NewObject()

	// Das Base32 Modul wird erezugt
	base32Module := jsruntime.NewObject()

	// Das Base58 Modul wird erezeugt
	base58Module := jsruntime.NewObject()

	// Das Base64 Modul wird erzeugt
	base64Module := jsruntime.NewObject()

	// Das Bech32 Modul wird erzeugt
	bech32Module := jsruntime.NewObject()

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

// Stellt die Dateisystem Funktionen bereit
func (o *JSEngine) loadFileSystemModule(jsruntime *goja.Runtime) goja.Value {
	// Das Filesystem Module wird erzeugt
	fileSystemModule := jsruntime.NewObject()
	fileSystemModule.Set("access", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("accessSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("appendFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("appendFileSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("chmod", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("chmodSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("chown", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("chownSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("close", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("closeSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("copyFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("copyFileSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("createReadStream", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("createWriteStream", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("dirent", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("exists", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("existsSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fchmod", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fchmodSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fchown", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fchownSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fdatasync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fdatasyncSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fstat", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fstatSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fsync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("fsyncSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("ftruncate", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("ftruncateSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("futimes", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("futimesSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lchmod", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lchmodSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lchown", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lchownSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("link", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("linkSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lstat", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lstatSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lutimes", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("lutimesSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("mkdir", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("mkdirSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("mkdtemp", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("mkdtempSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("open", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("openSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("opendir", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("read", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readdir", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readdirSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readFileSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readlink", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("readlinkSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("realpath", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("realpathSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("rename", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("renameSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("rmdir", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("rmdirSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("stat", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("statSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("symlink", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("symlinkSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("truncate", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("truncateSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("unlink", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("unlinkSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("unwatchFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("utimes", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("utimesSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("watch", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("watchFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("write", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("writeSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("writeFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	fileSystemModule.Set("writeFileSync", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// das Module wird zurückgegeben
	return fileSystemModule
}

// Stellt die GIT Funktionen bereit
func (o *JSEngine) loadingGitModule(jsruntime *goja.Runtime) goja.Value {
	// Das Repository Modul wird erzeugt
	repoModule := jsruntime.NewObject()

	// Das Commit Modul wird erzeugt
	commitModule := jsruntime.NewObject()

	// Das Branch Modul wird erzeugt
	branchModule := jsruntime.NewObject()

	// Das Index Modul wird erezuegt
	indexModule := jsruntime.NewObject()

	// Das Tag Modul wird erezeugt
	tagModule := jsruntime.NewObject()

	// Das Change Module wird erzeugt
	changeModule := jsruntime.NewObject()

	// Das Remote Modul wird erzeugt
	remoteModule := jsruntime.NewObject()

	// Das Konfig Modul wird erzeugt
	configModule := jsruntime.NewObject()

	// Das Git Module wird erzeugt
	gitModule := jsruntime.NewObject()
	gitModule.Set("repository", repoModule)
	gitModule.Set("commit", commitModule)
	gitModule.Set("branch", branchModule)
	gitModule.Set("index", indexModule)
	gitModule.Set("tag", tagModule)
	gitModule.Set("change", changeModule)
	gitModule.Set("remote", remoteModule)
	gitModule.Set("config", configModule)

	// das Module wird zurückgegeben
	return gitModule
}

// Stellt die Webdav Funktionen bereit
func (o *JSEngine) loadingWebDavModule(jsruntime *goja.Runtime) goja.Value {
	// Das WebDav Modul wird erzeugt
	webDavModule := jsruntime.NewObject()
	webDavModule.Set("NewServer", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	webDavModule.Set("Connect", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// das Module wird zurückgegeben
	return webDavModule
}

// Stellt Funktionen für die Verwaltung von Lokalen SQLite oder Lokalen MySQL Datenbanekn bereit
func (o *JSEngine) loadSQLDBModule(jsruntime *goja.Runtime) goja.Value {
	// Die SQL Funktionen werden geschrieben
	sqlModule := jsruntime.NewObject()
	sqlModule.Set("openFile", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	sqlModule.Set("connectTo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Gibt das Module zurück
	return sqlModule
}

// Stellt alle SSH Client Funktionen bereit
func (o *JSEngine) loadSSHClientModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	sshModule := jsruntime.NewObject()
	sshModule.Set("connectTo", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Gibt das Module zurück
	return sshModule
}

// Stellt die Standard NodeJS Buffer Funktionen zur verfügung
func (o *JSEngine) loadBufferModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	bufferModule := jsruntime.NewObject()
	bufferModule.Set("alloc", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("allocUnsafe", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("byteLength", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("compare", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("concat", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("copy", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("entries", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("equals", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("fill", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("from", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("includes", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("indexOf", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("isBuffer", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("isEncoding", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("keys", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("lastIndexOf", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readDoubleBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readDoubleLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readFloatBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readFloatLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readInt8", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readInt16BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readInt16LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readInt32BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readInt32LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readIntBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readIntLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUInt8", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUInt16BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUInt16LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUInt32BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUInt32LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUintBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("readUIntLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("slice", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("swap16", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("swap32", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("swap64", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("toString", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("toJSON", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("values", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("write", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeDoubleBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeFloatBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeFloatLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeInt8", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeInt16BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeInt16LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeInt32BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeInt32LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeIntBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeIntLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUInt8", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUInt16BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUInt16LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUInt32BE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUInt32LE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUIntBE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})
	bufferModule.Set("writeUIntLE", func(parms goja.FunctionCall) goja.Value {
		return nil
	})

	// Gibt das Module zurück
	return bufferModule
}
