package core

import "unicode/utf8"
import "math"

func Hx_ClassInit_Hx_Obj_VTable_haxe_int64helper() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_int64helper() {
}

type Hx_Obj_VTable_haxe_int64helper interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_int64helper struct {
	VTable Hx_Obj_VTable_haxe_int64helper `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_int64helper_CreateEmptyInstance() *Hx_Obj_haxe_int64helper {
	obj := &Hx_Obj_haxe_int64helper{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_int64helper_CreateInstance() *Hx_Obj_haxe_int64helper {
	obj := Hx_Obj_haxe_int64helper_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_int64helper_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Int64Helper",
	HxMakeArray[string]("parseString", "fromFloat", "toFloat"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_int64helper_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_int64helper_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_int64helper_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "parseString":
			return Hx_Field_haxe_int64helper_parseString
		case "fromFloat":
			return Hx_Field_haxe_int64helper_fromFloat
		case "toFloat":
			return Hx_Field_haxe_int64helper_toFloat
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_int64helper) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_int64helper_RTTI
}

func (this *Hx_Obj_haxe_int64helper) Hx_As_Hx_Obj_haxe_int64helper() *Hx_Obj_haxe_int64helper {
	return this
}

func Hx_Field_haxe_int64helper_parseString(sParam string) Hx_Typedef_haxe__int64___int64 {
	var _v int32
	_ = _v
	_v = 10
	var base Hx_Typedef_haxe__int64___int64
	_ = base
	base = ((Hx_Typedef_haxe__int64___int64)(_v))
	var _v1 int32
	_ = _v1
	_v1 = 0
	var current Hx_Typedef_haxe__int64___int64
	_ = current
	current = ((Hx_Typedef_haxe__int64___int64)(_v1))
	var _v2 int32
	_ = _v2
	_v2 = 1
	var multiplier Hx_Typedef_haxe__int64___int64
	_ = multiplier
	multiplier = ((Hx_Typedef_haxe__int64___int64)(_v2))
	var sIsNegative bool
	_ = sIsNegative
	sIsNegative = false
	var s string
	_ = s
	s = Hx_Field_stringtools_trim(sParam)
	if (Hx_Field_go_haxe_hxstring_charAt(s, 0)) == "-" {
		sIsNegative = true
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = s
		s = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_0, 1, ((HxNullable[int32])(HxNullable[int32]{Value: int32(utf8.RuneCountInString(s)), Valid: true}))))
	}

	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(s))
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
			var digitInt int32
			_ = digitInt
			digitInt = ((Hx_Field_go_haxe_hxstring_charCodeAt(s, ((_hx_reserved_len - 1) - i))).Value - 48)
			if (digitInt < 0) || (digitInt > 9) {
				panic(Hx_Field_haxe_exception_thrown(((any)("NumberFormatError"))))
			}

			if digitInt != 0 {
				var tmp__v_1 int32
				_ = tmp__v_1
				tmp__v_1 = digitInt
				var digit Hx_Typedef_haxe__int64___int64
				_ = digit
				digit = ((Hx_Typedef_haxe__int64___int64)(tmp__v_1))
				if sIsNegative {
					current = ((Hx_Typedef_haxe__int64___int64)((((int64)(current)) - ((int64)(((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)((((int64)(multiplier)) * ((int64)(digit))))))))))))
					var _hx_tmp_2 Hx_Typedef_haxe__int64___int64
					_ = _hx_tmp_2
					_hx_tmp_2 = current
					if !(_hx_tmp_2 < int64(((int32)(0)))) {
						panic(Hx_Field_haxe_exception_thrown(((any)("NumberFormatError: Underflow"))))
					}
				} else {
					current = ((Hx_Typedef_haxe__int64___int64)((((int64)(current)) + ((int64)(((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)((((int64)(multiplier)) * ((int64)(digit))))))))))))
					var _hx_tmp_3 Hx_Typedef_haxe__int64___int64
					_ = _hx_tmp_3
					_hx_tmp_3 = current
					if _hx_tmp_3 < int64(((int32)(0))) {
						panic(Hx_Field_haxe_exception_thrown(((any)("NumberFormatError: Overflow"))))
					}
				}
			}

			multiplier = ((Hx_Typedef_haxe__int64___int64)((((int64)(multiplier)) * ((int64)(base)))))
		}
	}

	return current
}

func Hx_Field_haxe_int64helper_fromFloat(f float64) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if Hx_Field_math_isNaN(f) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = !Hx_Field_math_isFinite(f)
	}

	if _hx_tmp_0 {
		panic(Hx_Field_haxe_exception_thrown(((any)("Number is NaN or Infinite"))))
	}

	var _hx_tmp_1 float64
	_ = _hx_tmp_1
	_hx_tmp_1 = f
	var noFractions float64
	_ = noFractions
	noFractions = (_hx_tmp_1 - math.Mod(f, ((float64)(1))))
	if noFractions > float64(9007199254740991) {
		panic(Hx_Field_haxe_exception_thrown(((any)("Conversion overflow"))))
	}

	if noFractions < float64(-9007199254740991) {
		panic(Hx_Field_haxe_exception_thrown(((any)("Conversion underflow"))))
	}

	var _v int32
	_ = _v
	_v = 0
	var result Hx_Typedef_haxe__int64___int64
	_ = result
	result = ((Hx_Typedef_haxe__int64___int64)(_v))
	var neg bool
	_ = neg
	neg = (noFractions < ((float64)(0)))
	var _hx_tmp_2 float64
	_ = _hx_tmp_2
	if neg {
		_hx_tmp_2 = -noFractions
	} else {
		_hx_tmp_2 = noFractions
	}

	var rest float64
	_ = rest
	rest = _hx_tmp_2
	var i int32
	_ = i
	i = 0
	for rest >= ((float64)(1)) {
		var curr float64
		_ = curr
		curr = math.Mod(rest, ((float64)(2)))
		rest = (rest / ((float64)(2)))
		if curr >= ((float64)(1)) {
			var tmp__v_1 int32
			_ = tmp__v_1
			tmp__v_1 = 1
			var a Hx_Typedef_haxe__int64___int64
			_ = a
			a = ((Hx_Typedef_haxe__int64___int64)(tmp__v_1))
			var _hx_tmp_3 int64
			_ = _hx_tmp_3
			_hx_tmp_3 = ((int64)(result))
			var _hx_tmp_4 Hx_Typedef_haxe__int64___int64
			_ = _hx_tmp_4
			if (i <= 0) || (i >= 64) {
				_hx_tmp_4 = a
			} else {
				var _hx_tmp_5 int64
				_ = _hx_tmp_5
				_hx_tmp_5 = ((int64)(a))
				var _hx_tmp_6 int64
				_ = _hx_tmp_6
				_hx_tmp_6 = _hx_tmp_5
				_hx_tmp_4 = ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_6 << int64(i))))
			}

			result = ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_3 + ((int64)((_hx_tmp_4))))))
		}

		i++
	}

	if neg {
		result = ((Hx_Typedef_haxe__int64___int64)(Hx_Field_haxe__int64_int64_impl__neg(result)))
	}

	return result
}

func Hx_Field_haxe_int64helper_toFloat(x Hx_Typedef_haxe__int64___int64) float64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(x))
	var high int32
	_ = high
	high = ((int32)((_hx_tmp_0 >> int64(((int32)(32))))))
	var low int32
	_ = low
	low = ((int32)(x))
	var _hx_tmp_1 float64
	_ = _hx_tmp_1
	_hx_tmp_1 = (((float64)(high)) * float64(4294967296.))
	var _hx_tmp_2 float64
	_ = _hx_tmp_2
	if low < 0 {
		_hx_tmp_2 = (((float64)(low)) + float64(4294967296.))
	} else {
		_hx_tmp_2 = ((float64)(low))
	}

	return (_hx_tmp_1 + (_hx_tmp_2))
}
