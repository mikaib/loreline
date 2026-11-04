package core

var Hx_Obj_go_errors_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Errors",
	HxMakeArray[string]("as", "asType", "_is", "join", "_new", "unwrap"),
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
