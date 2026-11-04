package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_exceptions_notimplementedexception() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_exceptions_notimplementedexception() {
}

type Hx_Obj_VTable_haxe_exceptions_notimplementedexception interface {
	Hx_Field_toString() string
	Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_unwrap() any
	Hx_Field_details() string
	Hx_Field_get_message() string
	Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_get_previous() *Hx_Obj_haxe_exception
	Hx_Field_get_native() any
	Hx_Field_error() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_exceptions_notimplementedexception struct {
	Hx_Obj_haxe_exceptions_posexception
	VTable Hx_Obj_VTable_haxe_exceptions_notimplementedexception `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_exceptions_notimplementedexception_CreateEmptyInstance() *Hx_Obj_haxe_exceptions_notimplementedexception {
	obj := &Hx_Obj_haxe_exceptions_notimplementedexception{}
	obj.Hx_Obj_haxe_exceptions_posexception = *Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_exceptions_posexception.VTable = obj
	obj.Hx_Obj_haxe_exception.VTable = obj
	return obj
}

func Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance(message__isOptional any, previous__isOptional any, pos__isOptional any) *Hx_Obj_haxe_exceptions_notimplementedexception {
	obj := Hx_Obj_haxe_exceptions_notimplementedexception_CreateEmptyInstance()
	obj.Hx_New(message__isOptional, previous__isOptional, pos__isOptional)
	return obj
}

func (this *Hx_Obj_haxe_exceptions_notimplementedexception) Hx_New(message__isOptional any, previous__isOptional any, pos__isOptional any) {
	{
		message := HxDefault[string]()
		_ = message

		if message__isOptional != nil {
			message = HxConvert[string](message__isOptional)
		} else {
			message = "Not implemented"
		}
		previous := HxDefault[*Hx_Obj_haxe_exception]()
		_ = previous

		if previous__isOptional != nil {
			previous = HxConvert[*Hx_Obj_haxe_exception](previous__isOptional)
		} else {
			previous = nil
		}
		pos := HxDefault[any]()
		_ = pos

		if pos__isOptional != nil {
			pos = HxConvert[any](pos__isOptional)
		} else {
			pos = nil
		}
		{
			this.Hx_Obj_haxe_exceptions_posexception.Hx_New(message, previous, pos)
		}

	}
}

var Hx_Obj_haxe_exceptions_notimplementedexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.exceptions.NotImplementedException",
	HxMakeArray[string](),
	HxMakeArray[string]("posInfos", "toString", "get_message", "get_stack", "set_stack", "get_previous", "get_native", "_message", "_native", "_stack", "unwrap", "error", "details"),
	Hx_Obj_haxe_exceptions_posexception_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)))
	},
	func() any {
		return Hx_Obj_haxe_exceptions_notimplementedexception_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_exceptions_notimplementedexception_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_exceptions_notimplementedexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_exceptions_notimplementedexception_RTTI
}

func (this *Hx_Obj_haxe_exceptions_notimplementedexception) Hx_As_Hx_Obj_haxe_exceptions_notimplementedexception() *Hx_Obj_haxe_exceptions_notimplementedexception {
	return this
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
