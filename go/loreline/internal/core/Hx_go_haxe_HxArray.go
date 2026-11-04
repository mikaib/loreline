package core

var Hx_Obj_go_haxe__hxarray_dynamicarray_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe._HxArray.DynamicArray_Impl_",
	HxMakeArray[string]("fromDynamic"),
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

var Hx_Obj_go_haxe_hxarray_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxArray",
	HxMakeArray[string]("fastGet", "fastSet", "elemType", "iterator", "keyValueIterator", "push", "pop", "shift", "unshift", "insert", "remove", "contains", "indexOf", "lastIndexOf", "concat", "copy", "slice", "splice", "reverse", "sort", "map", "filter", "resize", "join", "toString"),
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
