package core

func Hx_ClassInit_Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue() {
}

type Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue interface {
	Hx_Field_enumParameter(index int32) any
	Hx_Field_enumIndex() int32
	Hx_Field_enumParams() HxArray[any]
	Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe__hxenumvalue__hxenumvalue struct {
	VTable Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

var Hx_Obj_go_haxe__hxenumvalue__hxenumvalue_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe._HxEnumValue._HxEnumValue",
	HxMakeArray[string](),
	HxMakeArray[string]("enumIndex", "enumParams", "enumType", "enumParameter"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return nil
	},
	func() any {
		return nil
	},
	nil,
	func(name string) any {
		return nil
	},
)

var Hx_Obj_go_haxe__hxenumvalue_hxenumvalue_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe._HxEnumValue.HxEnumValue_Impl_",
	HxMakeArray[string](),
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

func init() {
	hxRegisterArrayConstructor[any]()
}
