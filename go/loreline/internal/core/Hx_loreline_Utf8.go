package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_utf8() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_utf8() {
}

type Hx_Obj_VTable_loreline_utf8 interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_utf8 struct {
	VTable Hx_Obj_VTable_loreline_utf8 `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_utf8_CreateEmptyInstance() *Hx_Obj_loreline_utf8 {
	obj := &Hx_Obj_loreline_utf8{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_utf8_CreateInstance() *Hx_Obj_loreline_utf8 {
	obj := Hx_Obj_loreline_utf8_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_utf8_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Utf8",
	HxMakeArray[string]("uLength", "uSubstr", "uSubstring", "uCharCodeAt", "uIndexOf", "uLastIndexOf", "uToChars", "uCharAt", "toCodes"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_utf8_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_utf8_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_utf8_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "uLength":
			return Hx_Field_loreline_utf8_uLength
		case "uSubstr":
			return Hx_Field_loreline_utf8_uSubstr
		case "uSubstring":
			return Hx_Field_loreline_utf8_uSubstring
		case "uCharCodeAt":
			return Hx_Field_loreline_utf8_uCharCodeAt
		case "uIndexOf":
			return Hx_Field_loreline_utf8_uIndexOf
		case "uLastIndexOf":
			return Hx_Field_loreline_utf8_uLastIndexOf
		case "uToChars":
			return Hx_Field_loreline_utf8_uToChars
		case "uCharAt":
			return Hx_Field_loreline_utf8_uCharAt
		case "toCodes":
			return Hx_Field_loreline_utf8_toCodes
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_utf8) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_utf8_RTTI
}

func (this *Hx_Obj_loreline_utf8) Hx_As_Hx_Obj_loreline_utf8() *Hx_Obj_loreline_utf8 {
	return this
}

func Hx_Field_loreline_utf8_uLength(str string) int32 {
	return int32(utf8.RuneCountInString(str))
}

func Hx_Field_loreline_utf8_uSubstr(str string, pos int32, _hx_reserved_len__isOptional any) string {
	_hx_reserved_len := HxDefault[HxNullable[int32]]()
	_ = _hx_reserved_len

	if _hx_reserved_len__isOptional != nil {
		_hx_reserved_len = HxConvert[HxNullable[int32]](_hx_reserved_len__isOptional)
	} else {

	}
	{
		if (_hx_reserved_len).Valid == false {
			var hx_nullable_t21 HxNullable[int32]
			_ = hx_nullable_t21
			hx_nullable_t21 = HxNullable[int32]{}

			_hx_reserved_len = HxNullable[int32]{Value: (hx_nullable_t21).Value, Valid: (hx_nullable_t21).Valid}
		}

		if (_hx_reserved_len).Valid != false {
			return (Hx_Field_go_haxe_hxstring_substr(str, pos, _hx_reserved_len))
		} else {
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = str
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = pos
			return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
		}
	}
}

func Hx_Field_loreline_utf8_uSubstring(str string, startIndex int32, endIndex__isOptional any) string {
	endIndex := HxDefault[HxNullable[int32]]()
	_ = endIndex

	if endIndex__isOptional != nil {
		endIndex = HxConvert[HxNullable[int32]](endIndex__isOptional)
	} else {

	}
	{
		if (endIndex).Valid == false {
			var hx_nullable_t22 HxNullable[int32]
			_ = hx_nullable_t22
			hx_nullable_t22 = HxNullable[int32]{}

			endIndex = HxNullable[int32]{Value: (hx_nullable_t22).Value, Valid: (hx_nullable_t22).Valid}
		}

		if (endIndex).Valid != false {
			return (Hx_Field_go_haxe_hxstring_substring(str, startIndex, endIndex))
		} else {
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = str
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = startIndex
			return (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
		}
	}
}

func Hx_Field_loreline_utf8_uCharCodeAt(str string, pos int32) int32 {
	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(str, pos))
	if (c).Valid == false {
		return -1
	} else {
		return (c).Value
	}
}

func Hx_Field_loreline_utf8_uIndexOf(str string, substr string, startIndex__isOptional any) int32 {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t23 HxNullable[int32]
			_ = hx_nullable_t23
			hx_nullable_t23 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t23).Value, Valid: (hx_nullable_t23).Valid}
		}

		return (Hx_Field_go_haxe_hxstring_indexOf(str, substr, startIndex))
	}
}

func Hx_Field_loreline_utf8_uLastIndexOf(str string, substr string, startIndex__isOptional any) int32 {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t24 HxNullable[int32]
			_ = hx_nullable_t24
			hx_nullable_t24 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t24).Value, Valid: (hx_nullable_t24).Valid}
		}

		return (Hx_Field_go_haxe_hxstring_lastIndexOf(str, substr, startIndex))
	}
}

func Hx_Field_loreline_utf8_uToChars(str string) HxArray[string] {
	return (Hx_Field_go_haxe_hxstring_split(str, ""))
}

func Hx_Field_loreline_utf8_uCharAt(str string, pos int32) string {
	return (Hx_Field_go_haxe_hxstring_charAt(str, pos))
}

func Hx_Field_loreline_utf8_toCodes(str string) HxArray[int32] {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(str))
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{Value: _hx_reserved_len, Valid: true}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var x []int32
	_ = x
	x = make([]int32, _hx_tmp_0)
	var self []int32
	_ = self
	self = x
	var v HxArray[int32]
	_ = v
	v = HxMakeArray(self...)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			{
				var val int32
				_ = val
				val = (Hx_Field_go_haxe_hxstring_charCodeAt(str, i)).Value
				HxMakeArrayView[any](v).Set(i, ((any)(val)))
			}
		}
	}

	return v
}

type Hx_Typedef_loreline_utf8buf = *Hx_Obj_stringbuf

func init() {
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
