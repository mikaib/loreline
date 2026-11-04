package core

import "encoding/json"
import "strings"
import "io"

func Hx_ClassInit_Hx_Obj_VTable_haxe_json() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_json() {
}

type Hx_Obj_VTable_haxe_json interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_json struct {
	VTable Hx_Obj_VTable_haxe_json `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_json_CreateEmptyInstance() *Hx_Obj_haxe_json {
	obj := &Hx_Obj_haxe_json{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_json_CreateInstance() *Hx_Obj_haxe_json {
	obj := Hx_Obj_haxe_json_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_json_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Json",
	HxMakeArray[string]("parse", "stringify", "convert"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_json_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_json_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_json_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "parse":
			return Hx_Field_haxe_json_parse
		case "stringify":
			return Hx_Field_haxe_json_stringify
		case "convert":
			return Hx_Field_haxe_json_convert
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_json) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_json_RTTI
}

func (this *Hx_Obj_haxe_json) Hx_As_Hx_Obj_haxe_json() *Hx_Obj_haxe_json {
	return this
}

func Hx_Field_haxe_json_parse(text string) any {
	var decoder *json.Decoder
	_ = decoder
	decoder = json.NewDecoder(((io.Reader)(strings.NewReader(text))))
	(*decoder).UseNumber()
	var raw any
	_ = raw
	var err error
	_ = err
	err = (*decoder).Decode(&raw)
	if err != nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid JSON: ", err.Error())))))
	}

	{
		var hx_result_12 struct {
			Error  error
			Result json.Token
		}
		_ = hx_result_12
		hx_result_12.Result, hx_result_12.Error = (*decoder).Token()
		var _g struct {
			Error  error
			Result json.Token
		}
		_ = _g
		_g = hx_result_12
		if _g.Error != nil {
			var e error
			_ = e
			e = ((error)(_g.Error))
			if e.Error() != "EOF" {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid JSON: ", e.Error())))))
			}
		} else {
			panic(Hx_Field_haxe_exception_thrown(((any)("Invalid JSON: unexpected data after top-level value"))))
		}
	}

	return Hx_Field_haxe_json_convert(raw)
}

func Hx_Field_haxe_json_stringify(value any, replacer__isOptional any, space__isOptional any) string {
	replacer := HxDefault[func(key any, value any) any]()
	_ = replacer

	if replacer__isOptional != nil {
		replacer = HxConvert[func(key any, value any) any](replacer__isOptional)
	} else {
		replacer = nil
	}
	space := HxDefault[HxNullable[string]]()
	_ = space

	if space__isOptional != nil {
		space = HxConvert[HxNullable[string]](space__isOptional)
	} else {

	}
	{
		if (space).Valid == false {
			var hx_nullable_t32 HxNullable[string]
			_ = hx_nullable_t32
			hx_nullable_t32 = HxNullable[string]{Value: HxStringNull}

			space = HxNullable[string]{Value: (hx_nullable_t32).Value, Valid: (hx_nullable_t32).Valid}
		}

		return Hx_Field_haxe_format_jsonprinter_print(value, replacer, space)
	}
}

func Hx_Field_haxe_json_convert(raw any) any {
	if func() bool { _, ok := raw.(map[string]any); return ok }() {
		var obj map[string]any
		_ = obj
		obj = raw.(map[string]any)
		{
			var length HxNullable[int32]
			_ = length
			length = HxNullable[int32]{}

			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (length).Valid != false {
				_hx_tmp_0 = (length).Value
			} else {
				_hx_tmp_0 = 0
			}

			var keys []string
			_ = keys
			keys = make([]string, _hx_tmp_0)
			for k := range obj {
				keys = append(keys, k)
			}

			var self []string
			_ = self
			self = keys
			var _gkey_current int32
			_ = _gkey_current
			_gkey_current = 0
			var _gkey_array HxArray[string]
			_ = _gkey_array
			_gkey_array = HxMakeArray(self...)
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _gkey_current
				if !(_hx_tmp_1 < _gkey_array.Len()) {
					break
				}

				var _hx_tmp_2 HxArray[string]
				_ = _hx_tmp_2
				_hx_tmp_2 = _gkey_array
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _gkey_current
				_gkey_current = (_gkey_current + 1)
				var key string
				_ = key
				key = _hx_tmp_2.Get(_hx_tmp_3)
				var _hx_tmp_4 map[string]any
				_ = _hx_tmp_4
				_hx_tmp_4 = obj
				var _hx_tmp_5 string
				_ = _hx_tmp_5
				_hx_tmp_5 = key
				_hx_tmp_4[_hx_tmp_5] = Hx_Field_haxe_json_convert(obj[key])
			}
		}

		return raw
	}

	if func() bool { _, ok := raw.([]any); return ok }() {
		var items []any
		_ = items
		items = raw.([]any)
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var self []any
			_ = self
			self = items
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[any]
			_ = _g_array
			_g_array = HxMakeArray(self...)
			for {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = _g_current
				if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g_array)) {
					break
				}

				var _hx_tmp_7 HxArray[any]
				_ = _hx_tmp_7
				_hx_tmp_7 = _g_array
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _g_current
				_g_current = (_g_current + 1)
				var item any
				_ = item
				item = _hx_tmp_7.Get(_hx_tmp_8)
				var _hx_tmp_9 HxArray[any]
				_ = _hx_tmp_9
				_hx_tmp_9 = _g
				_ = (Hx_Array_Push[any](_hx_tmp_9, Hx_Field_haxe_json_convert(item)))
			}
		}

		return ((any)(_g))
	}

	if func() bool { _, ok := raw.(json.Number); return ok }() {
		var text string
		_ = text
		text = string(raw.(json.Number))
		var f float64
		_ = f
		f = Hx_Field_std_parseFloat(text)
		var _hx_tmp_10 string
		_ = _hx_tmp_10
		_hx_tmp_10 = text
		if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_10, ".", HxNullable[int32]{})) >= 0 {
			return ((any)(f))
		}

		var i int32
		_ = i
		i = Hx_Field_std_int(f)
		if ((float64)(i)) == f {
			return ((any)(i))
		} else {
			return ((any)(f))
		}
	}

	return raw
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
