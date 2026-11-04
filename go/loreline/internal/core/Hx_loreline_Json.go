package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_json() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_json() {
}

type Hx_Obj_VTable_loreline_json interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_json struct {
	VTable Hx_Obj_VTable_loreline_json `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_json_CreateEmptyInstance() *Hx_Obj_loreline_json {
	obj := &Hx_Obj_loreline_json{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_json_CreateInstance() *Hx_Obj_loreline_json {
	obj := Hx_Obj_loreline_json_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_json_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Json",
	HxMakeArray[string]("stringify", "parse"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_json_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_json_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_json_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "stringify":
			return Hx_Field_loreline_json_stringify
		case "parse":
			return Hx_Field_loreline_json_parse
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_json) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_json_RTTI
}

func (this *Hx_Obj_loreline_json) Hx_As_Hx_Obj_loreline_json() *Hx_Obj_loreline_json {
	return this
}

func Hx_Field_loreline_json_stringify(value any, pretty__isOptional any) string {
	pretty := HxDefault[bool]()
	_ = pretty

	if pretty__isOptional != nil {
		pretty = HxConvert[bool](pretty__isOptional)
	} else {
		pretty = false
	}
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = value
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = nil
		var _hx_tmp_2 HxNullable[string]
		_ = _hx_tmp_2
		if pretty {
			_hx_tmp_2 = HxNullableString("  ")
		} else {
			_hx_tmp_2 = HxNullable[string]{Value: HxStringNull}
		}

		return Hx_Field_haxe_format_jsonprinter_print(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2)
	}
}

func Hx_Field_loreline_json_parse(json string) any {
	return Hx_Field_haxe_json_parse(json)
}
