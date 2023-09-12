package jsengine

import (
	"github.com/dop251/goja"
)

// Stellt die Crypto Funktionen bereit
func (o *JSEngine) loadCryptoModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	crypto_obj := jsruntime.NewObject()

	// Stellt die Standard NodeJS Crypto Funktionen bereit
	crypto_obj.Set("createCipher", nil)
	crypto_obj.Set("createCipheriv", nil)
	crypto_obj.Set("createDecipher", nil)
	crypto_obj.Set("createDecipheriv", nil)
	crypto_obj.Set("createDiffieHellman", nil)
	crypto_obj.Set("createECDH", nil)
	crypto_obj.Set("createHash", nil)
	crypto_obj.Set("createHmac", nil)
	crypto_obj.Set("createSign", nil)
	crypto_obj.Set("createVerify", nil)
	crypto_obj.Set("getCurves", nil)
	crypto_obj.Set("getDiffieHellman", nil)
	crypto_obj.Set("getHashes", nil)
	crypto_obj.Set("pbkdf2", nil)
	crypto_obj.Set("pbkdf2Sync", nil)
	crypto_obj.Set("privateDecrypt", nil)
	crypto_obj.Set("timingSafeEqual", nil)
	crypto_obj.Set("privateEncrypt", nil)
	crypto_obj.Set("publicDecrypt", nil)
	crypto_obj.Set("publicEncrypt", nil)
	crypto_obj.Set("randomBytes", nil)
	crypto_obj.Set("setEngine", nil)

	// Stellt zuzäzliche Post-quantum cryptography funktionen bereit
	poscryp := jsruntime.NewObject()
	crypto_obj.Set("pqc", poscryp)

	return crypto_obj
}

// Stellt die DNS Funktionen bereit
func (o *JSEngine) loadDNSModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	crypto_obj := jsruntime.NewObject()
	crypto_obj.Set("getServers", nil)
	crypto_obj.Set("lookup", nil)
	crypto_obj.Set("lookupService", nil)
	crypto_obj.Set("resolve", nil)
	crypto_obj.Set("resolve4", nil)
	crypto_obj.Set("resolve6", nil)
	crypto_obj.Set("resolveCname", nil)
	crypto_obj.Set("resolveMx", nil)
	crypto_obj.Set("resolveNaptr", nil)
	crypto_obj.Set("resolveNs", nil)
	crypto_obj.Set("resolveSoa", nil)
	crypto_obj.Set("resolveSrv", nil)
	crypto_obj.Set("resolvePtr", nil)
	crypto_obj.Set("resolveTxt", nil)
	crypto_obj.Set("reverse", nil)
	crypto_obj.Set("setServers", nil)
	return crypto_obj
}

// Stellt die Timer Funktionen bereit
func (o *JSEngine) loadTimersModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	crypto_obj := jsruntime.NewObject()
	crypto_obj.Set("clearImmediate", nil)
	crypto_obj.Set("clearInterval", nil)
	crypto_obj.Set("clearTimeout", nil)
	crypto_obj.Set("ref", nil)
	crypto_obj.Set("setImmediate", nil)
	crypto_obj.Set("setInterval", nil)
	crypto_obj.Set("setTimeout", nil)
	crypto_obj.Set("unref", nil)
	return crypto_obj
}

// Stellt die URL Funktionen bereit
func (o *JSEngine) loadURLModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	crypto_obj := jsruntime.NewObject()
	crypto_obj.Set("format", nil)
	crypto_obj.Set("parse", nil)
	crypto_obj.Set("resolve", nil)
	return crypto_obj
}

// Stellt Funktionen für Bitcoin bereit
func (o *JSEngine) loadBIPModule(jsruntime *goja.Runtime) goja.Value {
	// Stellt die Blockchain Funktionen bereit
	blockchain_object := jsruntime.NewObject()

	// Stellt die Netzwerk Funktionen bereit
	network_object := jsruntime.NewObject()

	// Das Objekt für die Transkationsfunktionen wird erzeugt
	transactions_object := jsruntime.NewObject()

	// Das Objekt für die Uitls wird erzeugt
	utils_oject := jsruntime.NewObject()

	// Das Objekt für die Wallet Funktionen wird erzeugt
	wallet_object := jsruntime.NewObject()

	// Die Funktionen werden hinzugefügt
	bipModule := jsruntime.NewObject()
	bipModule.Set("blockchain", blockchain_object)
	bipModule.Set("network", network_object)
	bipModule.Set("transactions", transactions_object)
	bipModule.Set("utils", utils_oject)
	bipModule.Set("wallet", wallet_object)

	// Das Module wird zurückgegeben
	return bipModule
}

// Stellt Funktionen für die Verwaltung von Lokalen SQLite oder Lokalen MySQL Datenbanekn bereit
func (o *JSEngine) loadSQLDBModule(jsruntime *goja.Runtime) goja.Value {
	// Die SQL Funktionen werden geschrieben
	sqlModule := jsruntime.NewObject()
	sqlModule.Set("openFile", nil)
	sqlModule.Set("connectTo", nil)

	// Gibt das Module zurück
	return sqlModule
}

// Stellt alle Funktionen für einen HTTPS Server bereit
func (o *JSEngine) loadHTTPSModule(jsruntime *goja.Runtime) goja.Value {
	// Die HTTPS Funktionen werden erzeugt
	httpsModule := jsruntime.NewObject()
	httpsModule.Set("createServer", nil)
	httpsModule.Set("get", nil)
	httpsModule.Set("globalAgent", nil)
	httpsModule.Set("request", nil)

	// Das Module wird zurückgegeben
	return httpsModule
}

// Stellt die ZLib Funktionen bereit
func (o *JSEngine) loadZLibModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	zlibMopdule := jsruntime.NewObject()
	zlibMopdule.Set("createDeflate", nil)
	zlibMopdule.Set("createDeflateRaw", nil)
	zlibMopdule.Set("createGunzip", nil)
	zlibMopdule.Set("createGzip", nil)
	zlibMopdule.Set("createInflate", nil)
	zlibMopdule.Set("createInflateRaw", nil)
	zlibMopdule.Set("deflate", nil)
	zlibMopdule.Set("deflateSync", nil)
	zlibMopdule.Set("deflateRaw", nil)
	zlibMopdule.Set("deflateRawSync", nil)
	zlibMopdule.Set("gunzip", nil)
	zlibMopdule.Set("gunzipSync", nil)
	zlibMopdule.Set("gzip", nil)
	zlibMopdule.Set("gzipSync", nil)
	zlibMopdule.Set("inflate", nil)
	zlibMopdule.Set("inflateSync", nil)
	zlibMopdule.Set("inflateRaw", nil)
	zlibMopdule.Set("inflateRawSync", nil)
	zlibMopdule.Set("unzip", nil)
	zlibMopdule.Set("unzipSync", nil)

	// Das Module wird zurückgegeben
	return zlibMopdule
}

// Stellt alle SSH Client Funktionen bereit
func (o *JSEngine) loadSSHClientModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	sshModule := jsruntime.NewObject()
	sshModule.Set("connectTo", nil)

	// Gibt das Module zurück
	return sshModule
}

// Stellt IPC Funktionen bereit
func (o *JSEngine) loadIPCModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	ipcModule := jsruntime.NewObject()
	ipcModule.Set("connectTo", nil)
	ipcModule.Set("createServer", nil)

	// Das Module wird zurückgegeben
	return ipcModule
}

// Stellt die Standard NodeJS Buffer Funktionen zur verfügung
func (o *JSEngine) loadBufferModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	bufferModule := jsruntime.NewObject()
	bufferModule.Set("alloc", nil)
	bufferModule.Set("allocUnsafe", nil)
	bufferModule.Set("byteLength", nil)
	bufferModule.Set("compare", nil)
	bufferModule.Set("concat", nil)
	bufferModule.Set("copy", nil)
	bufferModule.Set("entries", nil)
	bufferModule.Set("equals", nil)
	bufferModule.Set("fill", nil)
	bufferModule.Set("from", nil)
	bufferModule.Set("includes", nil)
	bufferModule.Set("indexOf", nil)
	bufferModule.Set("isBuffer", nil)
	bufferModule.Set("isEncoding", nil)
	bufferModule.Set("keys", nil)
	bufferModule.Set("lastIndexOf", nil)
	bufferModule.Set("readDoubleBE", nil)
	bufferModule.Set("readDoubleLE", nil)
	bufferModule.Set("readFloatBE", nil)
	bufferModule.Set("readFloatLE", nil)
	bufferModule.Set("readInt8", nil)
	bufferModule.Set("readInt16BE", nil)
	bufferModule.Set("readInt16LE", nil)
	bufferModule.Set("readInt32BE", nil)
	bufferModule.Set("readInt32LE", nil)
	bufferModule.Set("readIntBE", nil)
	bufferModule.Set("readIntLE", nil)
	bufferModule.Set("readUInt8", nil)
	bufferModule.Set("readUInt16BE", nil)
	bufferModule.Set("readUInt16LE", nil)
	bufferModule.Set("readUInt32BE", nil)
	bufferModule.Set("readUInt32LE", nil)
	bufferModule.Set("readUintBE", nil)
	bufferModule.Set("readUIntLE", nil)
	bufferModule.Set("slice", nil)
	bufferModule.Set("swap16", nil)
	bufferModule.Set("swap32", nil)
	bufferModule.Set("swap64", nil)
	bufferModule.Set("toString", nil)
	bufferModule.Set("toJSON", nil)
	bufferModule.Set("values", nil)
	bufferModule.Set("write", nil)
	bufferModule.Set("writeDoubleBE", nil)
	bufferModule.Set("writeFloatBE", nil)
	bufferModule.Set("writeFloatLE", nil)
	bufferModule.Set("writeInt8", nil)
	bufferModule.Set("writeInt16BE", nil)
	bufferModule.Set("writeInt16LE", nil)
	bufferModule.Set("writeInt32BE", nil)
	bufferModule.Set("writeInt32LE", nil)
	bufferModule.Set("writeIntBE", nil)
	bufferModule.Set("writeIntLE", nil)
	bufferModule.Set("writeUInt8", nil)
	bufferModule.Set("writeUInt16BE", nil)
	bufferModule.Set("writeUInt16LE", nil)
	bufferModule.Set("writeUInt32BE", nil)
	bufferModule.Set("writeUInt32LE", nil)
	bufferModule.Set("writeUIntBE", nil)
	bufferModule.Set("writeUIntLE", nil)

	// Gibt das Module zurück
	return bufferModule
}

// Stellt VM Funktionen bereit
func (o *JSEngine) loadVMModule(jsruntime *goja.Runtime) goja.Value {
	// Die Basis VM Modul Funktionen werden hinzugefügt
	vmModule := jsruntime.NewObject()
	vmModule.Set("createContext", nil)
	vmModule.Set("isContext", nil)
	vmModule.Set("runInContext", nil)
	vmModule.Set("runInDebug", nil)
	vmModule.Set("runInNewContext", nil)
	vmModule.Set("runInThisContext", nil)

	// Gibt das Module zurück
	return vmModule
}

// Stellt die Standard NodeJS Path Funktionen bereit
func (o *JSEngine) loadPathModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	pathModule := jsruntime.NewObject()
	pathModule.Set("format", nil)
	pathModule.Set("parse", nil)
	pathModule.Set("resolve", nil)

	// Gibt das Module zurück
	return pathModule
}

// Stellt Websocket Funktionen bereit
func (o *JSEngine) loadWebsocketModule(jsruntime *goja.Runtime) goja.Value {
	// Die Basis wird erstellt
	websocketModule := jsruntime.NewObject()
	websocketModule.Set("NewServer", nil)
	websocketModule.Set("Connect", nil)

	// Das Module wird zurückgegeben
	return websocketModule
}

// Stellt APISocket Funktionen bereit
func (o *JSEngine) loadAPISocketsModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die Dateisystem Funktionen bereit
func (o *JSEngine) loadFileSystemModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	fileSystemModule := jsruntime.NewObject()
	fileSystemModule.Set("access", nil)
	fileSystemModule.Set("accessSync", nil)
	fileSystemModule.Set("appendFile", nil)
	fileSystemModule.Set("appendFileSync", nil)
	fileSystemModule.Set("chmod", nil)
	fileSystemModule.Set("chmodSync", nil)
	fileSystemModule.Set("chown", nil)
	fileSystemModule.Set("chownSync", nil)
	fileSystemModule.Set("close", nil)
	fileSystemModule.Set("closeSync", nil)
	fileSystemModule.Set("copyFile", nil)
	fileSystemModule.Set("copyFileSync", nil)
	fileSystemModule.Set("createReadStream", nil)
	fileSystemModule.Set("createWriteStream", nil)
	fileSystemModule.Set("dirent", nil)
	fileSystemModule.Set("exists", nil)
	fileSystemModule.Set("existsSync", nil)
	fileSystemModule.Set("fchmod", nil)
	fileSystemModule.Set("fchmodSync", nil)
	fileSystemModule.Set("fchown", nil)
	fileSystemModule.Set("fchownSync", nil)
	fileSystemModule.Set("fdatasync", nil)
	fileSystemModule.Set("fdatasyncSync", nil)
	fileSystemModule.Set("fstat", nil)
	fileSystemModule.Set("fstatSync", nil)
	fileSystemModule.Set("fsync", nil)
	fileSystemModule.Set("fsyncSync", nil)
	fileSystemModule.Set("ftruncate", nil)
	fileSystemModule.Set("ftruncateSync", nil)
	fileSystemModule.Set("futimes", nil)
	fileSystemModule.Set("futimesSync", nil)
	fileSystemModule.Set("lchmod", nil)
	fileSystemModule.Set("lchmodSync", nil)
	fileSystemModule.Set("lchown", nil)
	fileSystemModule.Set("lchownSync", nil)
	fileSystemModule.Set("link", nil)
	fileSystemModule.Set("linkSync", nil)
	fileSystemModule.Set("lstat", nil)
	fileSystemModule.Set("lstatSync", nil)
	fileSystemModule.Set("lutimes", nil)
	fileSystemModule.Set("lutimesSync", nil)
	fileSystemModule.Set("mkdir", nil)
	fileSystemModule.Set("mkdirSync", nil)
	fileSystemModule.Set("mkdtemp", nil)
	fileSystemModule.Set("mkdtempSync", nil)
	fileSystemModule.Set("open", nil)
	fileSystemModule.Set("openSync", nil)
	fileSystemModule.Set("opendir", nil)
	fileSystemModule.Set("read", nil)
	fileSystemModule.Set("readSync", nil)
	fileSystemModule.Set("readdir", nil)
	fileSystemModule.Set("readdirSync", nil)
	fileSystemModule.Set("readFile", nil)
	fileSystemModule.Set("readFileSync", nil)
	fileSystemModule.Set("readlink", nil)
	fileSystemModule.Set("readlinkSync", nil)
	fileSystemModule.Set("realpath", nil)
	fileSystemModule.Set("realpathSync", nil)
	fileSystemModule.Set("rename", nil)
	fileSystemModule.Set("renameSync", nil)
	fileSystemModule.Set("rmdir", nil)
	fileSystemModule.Set("rmdirSync", nil)
	fileSystemModule.Set("stat", nil)
	fileSystemModule.Set("statSync", nil)
	fileSystemModule.Set("symlink", nil)
	fileSystemModule.Set("symlinkSync", nil)
	fileSystemModule.Set("truncate", nil)
	fileSystemModule.Set("truncateSync", nil)
	fileSystemModule.Set("unlink", nil)
	fileSystemModule.Set("unlinkSync", nil)
	fileSystemModule.Set("unwatchFile", nil)
	fileSystemModule.Set("utimes", nil)
	fileSystemModule.Set("utimesSync", nil)
	fileSystemModule.Set("watch", nil)
	fileSystemModule.Set("watchFile", nil)
	fileSystemModule.Set("write", nil)
	fileSystemModule.Set("writeSync", nil)
	fileSystemModule.Set("writeFile", nil)
	fileSystemModule.Set("writeFileSync", nil)

	// das Module wird zurückgegeben
	return fileSystemModule
}

// Stellt die Nativen Funktionen bereit
func (o *JSEngine) loadNativeModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die Encoding und Decoding Funktionen bereit
func (o *JSEngine) loadEncodingDecodingModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die GIT Funktionen bereit
func (o *JSEngine) loadingGitModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die Nostr Funktionen bereit
func (o *JSEngine) loadNostrModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die VMRPC Funktionen bereit
func (o *JSEngine) loadingVMRPCModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}

// Stellt die Webdav Funktionen bereit
func (o *JSEngine) loadingWebDavModule(jsruntime *goja.Runtime) goja.Value {
	// Die Crypto Funktionen werden geschrieben
	apiSocketsModule := jsruntime.NewObject()
	apiSocketsModule.Set("NewServer", nil)
	apiSocketsModule.Set("Connect", nil)

	// das Module wird zurückgegeben
	return apiSocketsModule
}
