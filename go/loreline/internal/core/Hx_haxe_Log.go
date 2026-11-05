package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_log() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_log() {
}

type Hx_Obj_VTable_haxe_log interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_log struct {
	VTable Hx_Obj_VTable_haxe_log `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_log_CreateEmptyInstance() *Hx_Obj_haxe_log {
	obj := &Hx_Obj_haxe_log{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_log_CreateInstance() *Hx_Obj_haxe_log {
	obj := Hx_Obj_haxe_log_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_log_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Log",
	HxMakeArray[string]("formatOutput", "trace"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_log_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_log_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_log_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "formatOutput":
			return Hx_Field_haxe_log_formatOutput
		case "trace":
			return Hx_Field_haxe_log_trace
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_log) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_log_RTTI
}

func (this *Hx_Obj_haxe_log) Hx_As_Hx_Obj_haxe_log() *Hx_Obj_haxe_log {
	return this
}

func Hx_Field_haxe_log_formatOutput(v any, infos any) string {
	var str string
	_ = str
	str = Hx_Field_std_string(v)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(infos, nil))) {
		return str
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(infos, "fileName"), ((any)(":")))
	var pstr string
	_ = pstr
	pstr = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(infos, "lineNumber")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(infos, "customParams"), ((any)(nil))))) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(infos, "customParams"))
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var v any
			_ = v
			v = _g1.Get(_g)
			_g++
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = str
			str = Hx_Field_go_haxe_hxstring_concat(_hx_tmp_2, Hx_Field_go_haxe_hxstring_concat(", ", Hx_Field_std_string(v)))
		}
	}

	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(pstr, ": "), str)
}

var Hx_Field_haxe_log_trace = func(v any, infos__isOptional any) {
	infos := HxDefault[any]()
	_ = infos

	if infos__isOptional != nil {
		infos = HxConvert[any](infos__isOptional)
	} else {
		infos = nil
	}
	{
		var str string
		_ = str
		str = Hx_Field_haxe_log_formatOutput(v, infos)
		Hx_Field_sys_println(str)
	}
}

func init() {
	hxRegisterArrayConstructor[any]()
}
