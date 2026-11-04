package core

var Hx_Obj_go_go_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Go",
	HxMakeArray[string]("int", "uint", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "byte", "rune", "complex64", "complex128", "print", "println", "panic", "len", "append", "copy", "cap", "string", "hxFloat", "hxInt"),
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
