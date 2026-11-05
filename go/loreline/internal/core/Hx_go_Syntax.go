package core

var Hx_Obj_go_syntax_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Syntax",
	HxMakeArray[string]("code", "defer", "go", "select"),
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

var Hx_Obj_go_select_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Select",
	HxMakeArray[string]("receive", "send", "defaultCase"),
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
