package core

var Hx_Obj_go_runtime_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Runtime",
	HxMakeArray[string]("compiler", "GOARCH", "GOOS", "memProfileRate", "addCleanup", "blockProfile", "breakpoint", "cPUProfile", "caller", "callers", "callersFrames", "funcForPC", "GC", "GOMAXPROCS", "GOROOT", "goexit", "goroutineProfile", "gosched", "keepAlive", "lockOSThread", "memProfile", "mutexProfile", "numCPU", "numCgoCall", "numGoroutine", "readMemStats", "readTrace", "setBlockProfileRate", "setCPUProfileRate", "setCgoTraceback", "setDefaultGOMAXPROCS", "setFinalizer", "setMutexProfileFraction", "stack", "startTrace", "stopTrace", "threadCreateProfile", "unlockOSThread", "version"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		panic("Cannot dynamically create instance of extern class")
	},
	func() any {
		panic("Cannot dynamically create instance of extern class")
	},
	nil,
	func(name string) any {
		return nil
	},
)
