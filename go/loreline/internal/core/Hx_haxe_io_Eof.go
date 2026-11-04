package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_eof() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_eof() {
}

type Hx_Obj_VTable_haxe_io_eof interface {
	Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_get_previous() *Hx_Obj_haxe_exception
	Hx_Field_get_native() any
	Hx_Field_toString() string
	Hx_Field_error() string
	Hx_Field_details() string
	Hx_Field_get_message() string
	Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_unwrap() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_eof struct {
	Hx_Obj_haxe_exception
	VTable Hx_Obj_VTable_haxe_io_eof `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_io_eof_CreateEmptyInstance() *Hx_Obj_haxe_io_eof {
	obj := &Hx_Obj_haxe_io_eof{}
	obj.Hx_Obj_haxe_exception = *Hx_Obj_haxe_exception_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_exception.VTable = obj
	return obj
}

func Hx_Obj_haxe_io_eof_CreateInstance() *Hx_Obj_haxe_io_eof {
	obj := Hx_Obj_haxe_io_eof_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_io_eof) Hx_New() {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = nil
		this.Hx_Obj_haxe_exception.Hx_New("Eof", _hx_tmp_0, nil)
	}

}

var Hx_Obj_haxe_io_eof_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.Eof",
	HxMakeArray[string](),
	HxMakeArray[string]("get_message", "get_stack", "set_stack", "get_previous", "get_native", "_message", "_native", "_stack", "unwrap", "toString", "error", "details"),
	Hx_Obj_haxe_exception_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_eof_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_io_eof_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_eof_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_io_eof) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_eof_RTTI
}

func (this *Hx_Obj_haxe_io_eof) Hx_As_Hx_Obj_haxe_io_eof() *Hx_Obj_haxe_io_eof {
	return this
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
