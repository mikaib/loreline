package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_wrappederror() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_wrappederror() {
}

type Hx_Obj_VTable_loreline_wrappederror interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_wrappederror struct {
	Hx_Obj_loreline_error
	VTable           Hx_Obj_VTable_loreline_wrappederror `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_wrapped any
}

func Hx_Obj_loreline_wrappederror_CreateEmptyInstance() *Hx_Obj_loreline_wrappederror {
	obj := &Hx_Obj_loreline_wrappederror{}
	obj.Hx_Obj_loreline_error = *Hx_Obj_loreline_error_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_error.VTable = obj
	obj.Hx_Field_wrapped = HxDefault[any]()
	return obj
}

func Hx_Obj_loreline_wrappederror_CreateInstance(wrapped any, message string, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_wrappederror {
	obj := Hx_Obj_loreline_wrappederror_CreateEmptyInstance()
	obj.Hx_New(wrapped, message, pos)
	return obj
}

func (this *Hx_Obj_loreline_wrappederror) Hx_New(wrapped any, message string, pos *Hx_Obj_loreline_position) {
	{
		this.Hx_Obj_loreline_error.Hx_New(message, pos)
		this.Hx_Field_wrapped = wrapped
	}

}

var Hx_Obj_loreline_wrappederror_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.WrappedError",
	HxMakeArray[string](),
	HxMakeArray[string]("wrapped", "message", "pos", "stack", "toString"),
	Hx_Obj_loreline_error_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_wrappederror_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[string](params.Get(1)), HxConvert[*Hx_Obj_loreline_position](params.Get(2)))
	},
	func() any {
		return Hx_Obj_loreline_wrappederror_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_wrappederror_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_wrappederror) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_wrappederror_RTTI
}

func (this *Hx_Obj_loreline_wrappederror) Hx_As_Hx_Obj_loreline_wrappederror() *Hx_Obj_loreline_wrappederror {
	return this
}
