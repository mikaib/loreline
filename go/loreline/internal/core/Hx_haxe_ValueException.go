package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_valueexception() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_valueexception() {
}

type Hx_Obj_VTable_haxe_valueexception interface {
	Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_get_previous() *Hx_Obj_haxe_exception
	Hx_Field_error() string
	Hx_Field_unwrap() any
	Hx_Field_get_message() string
	Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_get_native() any
	Hx_Field_toString() string
	Hx_Field_details() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_valueexception struct {
	Hx_Obj_haxe_exception
	VTable         Hx_Obj_VTable_haxe_valueexception `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_value any
}

func Hx_Obj_haxe_valueexception_CreateEmptyInstance() *Hx_Obj_haxe_valueexception {
	obj := &Hx_Obj_haxe_valueexception{}
	obj.Hx_Obj_haxe_exception = *Hx_Obj_haxe_exception_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_exception.VTable = obj
	obj.Hx_Field_value = HxDefault[any]()
	return obj
}

func Hx_Obj_haxe_valueexception_CreateInstance(value any, previous__isOptional any, native__isOptional any) *Hx_Obj_haxe_valueexception {
	obj := Hx_Obj_haxe_valueexception_CreateEmptyInstance()
	obj.Hx_New(value, previous__isOptional, native__isOptional)
	return obj
}

func (this *Hx_Obj_haxe_valueexception) Hx_New(value any, previous__isOptional any, native__isOptional any) {
	{
		previous := HxDefault[*Hx_Obj_haxe_exception]()
		_ = previous

		if previous__isOptional != nil {
			previous = HxConvert[*Hx_Obj_haxe_exception](previous__isOptional)
		} else {
			previous = nil
		}
		native := HxDefault[any]()
		_ = native

		if native__isOptional != nil {
			native = HxConvert[any](native__isOptional)
		} else {
			native = nil
		}
		{
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
				_hx_tmp_0 = "null"
			} else {
				_hx_tmp_0 = Hx_Field_std_string(value)
			}

			this.Hx_Obj_haxe_exception.Hx_New(_hx_tmp_0, previous, native)
			this.Hx_Field_value = value
		}

	}
}

var Hx_Obj_haxe_valueexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ValueException",
	HxMakeArray[string](),
	HxMakeArray[string]("value", "unwrap", "get_message", "get_stack", "set_stack", "get_previous", "get_native", "_message", "_native", "_stack", "toString", "error", "details"),
	Hx_Obj_haxe_exception_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_valueexception_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)))
	},
	func() any {
		return Hx_Obj_haxe_valueexception_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_valueexception_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_valueexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_valueexception_RTTI
}

func (this *Hx_Obj_haxe_valueexception) Hx_As_Hx_Obj_haxe_valueexception() *Hx_Obj_haxe_valueexception {
	return this
}

func (this *Hx_Obj_haxe_valueexception) Hx_Field_unwrap() any {
	return this.Hx_Field_value
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
