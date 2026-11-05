package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_fields() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_fields() {
}

type Hx_Obj_VTable_loreline_fields interface {
	Hx_Field_lorelineSet(interpreter *Hx_Obj_loreline_interpreter, key string, value any)
	Hx_Field_lorelineRemove(interpreter *Hx_Obj_loreline_interpreter, key string) bool
	Hx_Field_lorelineExists(interpreter *Hx_Obj_loreline_interpreter, key string) bool
	Hx_Field_lorelineFields(interpreter *Hx_Obj_loreline_interpreter) HxArray[string]
	Hx_Field_lorelineCreate(interpreter *Hx_Obj_loreline_interpreter)
	Hx_Field_lorelineGet(interpreter *Hx_Obj_loreline_interpreter, key string) any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_fields struct {
	VTable Hx_Obj_VTable_loreline_fields `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

var Hx_Obj_loreline_fields_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Fields",
	HxMakeArray[string](),
	HxMakeArray[string]("lorelineCreate", "lorelineGet", "lorelineSet", "lorelineRemove", "lorelineExists", "lorelineFields"),
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

func init() {
	hxRegisterArrayConstructor[string]()
}
