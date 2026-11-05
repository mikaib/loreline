package core

import "strings"
import "unicode/utf8"
import "slices"

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxstring() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxstring() {
}

type Hx_Obj_VTable_go_haxe_hxstring interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxstring struct {
	VTable Hx_Obj_VTable_go_haxe_hxstring `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go_haxe_hxstring_CreateEmptyInstance() *Hx_Obj_go_haxe_hxstring {
	obj := &Hx_Obj_go_haxe_hxstring{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go_haxe_hxstring_CreateInstance() *Hx_Obj_go_haxe_hxstring {
	obj := Hx_Obj_go_haxe_hxstring_CreateEmptyInstance()
	return obj
}

var Hx_Obj_go_haxe_hxstring_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxString",
	HxMakeArray[string]("toUpperCase", "toLowerCase", "concat", "charAt", "charCodeAt", "indexOf", "lastIndexOf", "split", "substr", "substring", "toString", "fromCharCode"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxstring_CreateInstance()
	},
	func() any {
		return Hx_Obj_go_haxe_hxstring_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxstring_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toUpperCase":
			return Hx_Field_go_haxe_hxstring_toUpperCase
		case "toLowerCase":
			return Hx_Field_go_haxe_hxstring_toLowerCase
		case "concat":
			return Hx_Field_go_haxe_hxstring_concat
		case "charAt":
			return Hx_Field_go_haxe_hxstring_charAt
		case "charCodeAt":
			return Hx_Field_go_haxe_hxstring_charCodeAt
		case "indexOf":
			return Hx_Field_go_haxe_hxstring_indexOf
		case "lastIndexOf":
			return Hx_Field_go_haxe_hxstring_lastIndexOf
		case "split":
			return Hx_Field_go_haxe_hxstring_split
		case "substr":
			return Hx_Field_go_haxe_hxstring_substr
		case "substring":
			return Hx_Field_go_haxe_hxstring_substring
		case "toString":
			return Hx_Field_go_haxe_hxstring_toString
		case "fromCharCode":
			return Hx_Field_go_haxe_hxstring_fromCharCode
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxstring) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxstring_RTTI
}

func (this *Hx_Obj_go_haxe_hxstring) Hx_As_Hx_Obj_go_haxe_hxstring() *Hx_Obj_go_haxe_hxstring {
	return this
}

func Hx_Field_go_haxe_hxstring_toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func Hx_Field_go_haxe_hxstring_toLowerCase(s string) string {
	return strings.ToLower(s)
}

func Hx_Field_go_haxe_hxstring_concat(a string, b string) string {
	return HxStringConcat(a, b)
}

func Hx_Field_go_haxe_hxstring_charAt(s string, index int32) string {
	var v HxNullable[int32]
	_ = v
	v = Hx_Field_go_haxe_hxstring_charCodeAt(s, index)
	if (v).Valid == false {
		return ""
	}

	return Hx_Field_go_haxe_hxstring_fromCharCode((v).Value)
}

func Hx_Field_go_haxe_hxstring_charCodeAt(s string, index int32) HxNullable[int32] {
	if index < 0 {
		return HxNullable[int32]{}
	}

	var i int32
	_ = i
	i = 0
	var j int32
	_ = j
	j = 0
	var iLen int32
	_ = iLen
	iLen = ((int32)(len(s)))
	for i < iLen {
		if j == index {
			var hx_tuple_1 struct {
				R    rune
				Size int
			}
			_ = hx_tuple_1
			hx_tuple_1.R, hx_tuple_1.Size = utf8.DecodeRuneInString(s[i:])
			var v struct {
				R    rune
				Size int
			}
			_ = v
			v = ((struct {
				R    rune
				Size int
			})(hx_tuple_1))
			var _hx_tmp_0 int
			_ = _hx_tmp_0
			_hx_tmp_0 = v.Size
			if _hx_tmp_0 == int(((int32)(0))) {
				return HxNullable[int32]{}
			} else {
				return HxNullable[int32]{Value: ((int32)(v.R)), Valid: true}
			}
		}

		var b int32
		_ = b
		b = int32(s[i])
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = i
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if b < 128 {
			_hx_tmp_2 = 1
		} else {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			if b < 224 {
				_hx_tmp_3 = 2
			} else {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if b < 240 {
					_hx_tmp_4 = 3
				} else {
					_hx_tmp_4 = 4
				}

				_hx_tmp_3 = _hx_tmp_4
			}

			_hx_tmp_2 = _hx_tmp_3
		}

		i = (_hx_tmp_1 + _hx_tmp_2)
		j++
	}

	return HxNullable[int32]{}
}

func Hx_Field_go_haxe_hxstring_indexOf(s string, str string, startIndex__isOptional any) int32 {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t1 HxNullable[int32]
			_ = hx_nullable_t1
			hx_nullable_t1 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t1).Value, Valid: (hx_nullable_t1).Valid}
		}

		var s1 []rune
		_ = s1
		s1 = (([]rune)(s))
		var str1 []rune
		_ = str1
		str1 = (([]rune)(str))
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = len(s1)
		var _hx_tmp_0 int
		_ = _hx_tmp_0
		_hx_tmp_0 = (_hx_tmp_1 - len(str1))
		var _hx_reserved_len int
		_ = _hx_reserved_len
		_hx_reserved_len = (_hx_tmp_0 + int(((int32)(1))))
		var tmp HxNullable[int32]
		_ = tmp
		tmp = startIndex
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (tmp).Valid != false {
			_hx_tmp_2 = (tmp).Value
		} else {
			_hx_tmp_2 = 0
		}

		startIndex = HxNullable[int32]{Value: _hx_tmp_2, Valid: true}

		if (startIndex).Valid && true && ((startIndex).Value < 0) {
			startIndex = HxNullable[int32]{Value: 0, Valid: true}
		}

		var _hx_tmp_3 int
		_ = _hx_tmp_3
		_hx_tmp_3 = len(str1)
		if _hx_tmp_3 == int(((int32)(0))) {
			var _hx_tmp_4 float64
			_ = _hx_tmp_4
			_hx_tmp_4 = float64(((float64)((startIndex).Value)))
			if _hx_tmp_4 > float64(len(s1)) {
				return ((int32)(len(s1)))
			} else {
				return (startIndex).Value
			}
		}

		var _hx_tmp_5 float64
		_ = _hx_tmp_5
		_hx_tmp_5 = float64(((float64)((startIndex).Value)))
		if _hx_tmp_5 >= float64(len(s1)) {
			return -1
		}

		var _hx_tmp_6 int
		_ = _hx_tmp_6
		_hx_tmp_6 = _hx_reserved_len
		if _hx_tmp_6 < int(((int32)(0))) {
			return -1
		}

		{
			var _g int32
			_ = _g
			_g = (startIndex).Value
			var _g1 int32
			_ = _g1
			_g1 = ((int32)(_hx_reserved_len))
			for _g < _g1 {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_7
				var _hx_tmp_8 rune
				_ = _hx_tmp_8
				_hx_tmp_8 = s1[i]
				if _hx_tmp_8 != str1[0] {
					continue
				}

				var _hx_tmp_9 []rune
				_ = _hx_tmp_9
				_hx_tmp_9 = s1
				var _hx_tmp_10 int
				_ = _hx_tmp_10
				_hx_tmp_10 = int(i)
				var _hx_tmp_11 int
				_ = _hx_tmp_11
				_hx_tmp_11 = len(str1)
				if slices.Equal(_hx_tmp_9[_hx_tmp_10:(_hx_tmp_11+int(i))], str1) {
					return i
				}
			}
		}

		return -1
	}
}

func Hx_Field_go_haxe_hxstring_lastIndexOf(s string, str string, startIndex__isOptional any) int32 {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t2 HxNullable[int32]
			_ = hx_nullable_t2
			hx_nullable_t2 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t2).Value, Valid: (hx_nullable_t2).Valid}
		}

		var r []rune
		_ = r
		r = (([]rune)(s))
		var n []rune
		_ = n
		n = (([]rune)(str))
		var _hx_tmp_0 int
		_ = _hx_tmp_0
		_hx_tmp_0 = len(r)
		var i int
		_ = i
		i = (_hx_tmp_0 - len(n))
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (startIndex).Valid != false {
			var _hx_tmp_2 float64
			_ = _hx_tmp_2
			_hx_tmp_2 = float64(((float64)((startIndex).Value)))
			var _hx_tmp_3 float64
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 < float64(i))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			i = int((startIndex).Value)
		}

		for {
			var _hx_tmp_4 int
			_ = _hx_tmp_4
			_hx_tmp_4 = i
			if !(_hx_tmp_4 >= int(((int32)(0)))) {
				break
			}

			var _hx_tmp_5 []rune
			_ = _hx_tmp_5
			_hx_tmp_5 = r
			var _hx_tmp_6 int
			_ = _hx_tmp_6
			_hx_tmp_6 = i
			var _hx_tmp_7 int
			_ = _hx_tmp_7
			_hx_tmp_7 = i
			if slices.Equal(_hx_tmp_5[_hx_tmp_6:(_hx_tmp_7+len(n))], n) {
				return ((int32)(i))
			}

			i = (i - ((int)(1)))
		}

		return -1
	}
}

func Hx_Field_go_haxe_hxstring_split(s string, delimiter string) HxArray[string] {
	var x []string
	_ = x
	x = strings.Split(s, delimiter)
	var self []string
	_ = self
	self = x
	return HxMakeArray(self...)
}

func Hx_Field_go_haxe_hxstring_substr(s string, pos int32, _hx_reserved_len__isOptional any) string {
	_hx_reserved_len := HxDefault[HxNullable[int32]]()
	_ = _hx_reserved_len

	if _hx_reserved_len__isOptional != nil {
		_hx_reserved_len = HxConvert[HxNullable[int32]](_hx_reserved_len__isOptional)
	} else {

	}
	{
		if (_hx_reserved_len).Valid == false {
			var hx_nullable_t3 HxNullable[int32]
			_ = hx_nullable_t3
			hx_nullable_t3 = HxNullable[int32]{}

			_hx_reserved_len = HxNullable[int32]{Value: (hx_nullable_t3).Value, Valid: (hx_nullable_t3).Valid}
		}

		var sLen int32
		_ = sLen
		sLen = int32(utf8.RuneCountInString(s))
		if pos > sLen {
			return ""
		}

		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if pos < 0 {
			_hx_tmp_0 = (sLen + pos)
		} else {
			_hx_tmp_0 = pos
		}

		var startIndex int32
		_ = startIndex
		startIndex = _hx_tmp_0
		if startIndex < 0 {
			startIndex = 0
		}

		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (_hx_reserved_len).Valid != false {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = startIndex
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 + (_hx_reserved_len).Value)
		} else {
			_hx_tmp_1 = sLen
		}

		var end int32
		_ = end
		end = _hx_tmp_1
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if (_hx_reserved_len).Valid != false {
			_hx_tmp_4 = ((_hx_reserved_len).Valid && true && ((_hx_reserved_len).Value < 0))
		} else {
			_hx_tmp_4 = false
		}

		if _hx_tmp_4 {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = sLen
			end = (_hx_tmp_5 + (_hx_reserved_len).Value)
		}

		if end > sLen {
			end = sLen
		}

		if end <= startIndex {
			return ""
		}

		return string(([]rune)(s)[startIndex:end])
	}
}

func Hx_Field_go_haxe_hxstring_substring(s string, startIndex int32, endIndex__isOptional any) string {
	endIndex := HxDefault[HxNullable[int32]]()
	_ = endIndex

	if endIndex__isOptional != nil {
		endIndex = HxConvert[HxNullable[int32]](endIndex__isOptional)
	} else {

	}
	{
		if (endIndex).Valid == false {
			var hx_nullable_t4 HxNullable[int32]
			_ = hx_nullable_t4
			hx_nullable_t4 = HxNullable[int32]{}

			endIndex = HxNullable[int32]{Value: (hx_nullable_t4).Value, Valid: (hx_nullable_t4).Valid}
		}

		var sLen int32
		_ = sLen
		sLen = int32(utf8.RuneCountInString(s))
		if (endIndex).Valid == false {
			endIndex = HxNullable[int32]{Value: sLen, Valid: true}
		}

		if startIndex < 0 {
			startIndex = 0
		}

		if (endIndex).Valid && true && ((endIndex).Value < 0) {
			endIndex = HxNullable[int32]{Value: 0, Valid: true}
		}

		if true && (endIndex).Valid && (startIndex > (endIndex).Value) {
			var tmp HxNullable[int32]
			_ = tmp
			tmp = endIndex
			endIndex = HxNullable[int32]{Value: startIndex, Valid: true}

			startIndex = (tmp).Value
		}

		if (endIndex).Valid && true && ((endIndex).Value > sLen) {
			endIndex = HxNullable[int32]{Value: sLen, Valid: true}
		}

		if startIndex > sLen {
			return ""
		}

		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = s
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = startIndex
		return string(([]rune)(_hx_tmp_0)[_hx_tmp_1:(endIndex).Value])
	}
}

func Hx_Field_go_haxe_hxstring_toString(s string) string {
	return s
}

func Hx_Field_go_haxe_hxstring_fromCharCode(code int32) string {
	return string(rune(code))
}

func init() {
	hxRegisterArrayConstructor[string]()
}
