package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_exceptions_posexception() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_exceptions_posexception() {
}

type Hx_Obj_VTable_haxe_exceptions_posexception interface {
	Hx_Field_get_previous() *Hx_Obj_haxe_exception
	Hx_Field_get_native() any
	Hx_Field_unwrap() any
	Hx_Field_error() string
	Hx_Field_details() string
	Hx_Field_toString() string
	Hx_Field_get_message() string
	Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_exceptions_posexception struct {
	Hx_Obj_haxe_exception
	VTable            Hx_Obj_VTable_haxe_exceptions_posexception `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_posInfos any
}

func Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance() *Hx_Obj_haxe_exceptions_posexception {
	obj := &Hx_Obj_haxe_exceptions_posexception{}
	obj.Hx_Obj_haxe_exception = *Hx_Obj_haxe_exception_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_exception.VTable = obj
	obj.Hx_Field_posInfos = HxDefault[any]()
	return obj
}

func Hx_Obj_haxe_exceptions_posexception_CreateInstance(message string, previous__isOptional any, pos__isOptional any) *Hx_Obj_haxe_exceptions_posexception {
	obj := Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance()
	obj.Hx_New(message, previous__isOptional, pos__isOptional)
	return obj
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_New(message string, previous__isOptional any, pos__isOptional any) {
	{
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
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = message
			var _hx_tmp_1 *Hx_Obj_haxe_exception
			_ = _hx_tmp_1
			_hx_tmp_1 = previous
			this.Hx_Obj_haxe_exception.Hx_New(_hx_tmp_0, _hx_tmp_1, nil)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(pos, nil))) {
				this.Hx_Field_posInfos = any(map[string]any{"fileName": ((any)("(unknown)")), "lineNumber": ((int32)(0)), "className": ((any)("(unknown)")), "methodName": ((any)("(unknown)"))})
			} else {
				this.Hx_Field_posInfos = pos
			}
		}

	}
}

var Hx_Obj_haxe_exceptions_posexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.exceptions.PosException",
	HxMakeArray[string](),
	HxMakeArray[string]("posInfos", "toString", "get_message", "get_stack", "set_stack", "get_previous", "get_native", "_message", "_native", "_stack", "unwrap", "error", "details"),
	Hx_Obj_haxe_exception_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_exceptions_posexception_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)))
	},
	func() any {
		return Hx_Obj_haxe_exceptions_posexception_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_exceptions_posexception_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_exceptions_posexception_RTTI
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_As_Hx_Obj_haxe_exceptions_posexception() *Hx_Obj_haxe_exceptions_posexception {
	return this
}

func (this *Hx_Obj_haxe_exceptions_posexception) Hx_Field_toString() string {
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = ((any)(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", this.Hx_Obj_haxe_exception.Hx_Field_toString()), " in ")))
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "className")), ((any)(".")))
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "methodName")), ((any)(" at ")))
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "fileName")), ((any)(":")))
	return Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_posInfos, "lineNumber")))
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
