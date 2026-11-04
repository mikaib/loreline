package core

import "strconv"
import "unicode/utf8"
import "strings"
import "reflect"
import "fmt"
import "math"

func Hx_ClassInit_Hx_Obj_VTable_std() {
}

func Hx_StaticInit_Hx_Obj_VTable_std() {
}

type Hx_Obj_VTable_std interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_std struct {
	VTable Hx_Obj_VTable_std `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_std_CreateEmptyInstance() *Hx_Obj_std {
	obj := &Hx_Obj_std{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_std_CreateInstance() *Hx_Obj_std {
	obj := Hx_Obj_std_CreateEmptyInstance()
	return obj
}

var Hx_Obj_std_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Std",
	HxMakeArray[string]("int", "isSpaceChar", "isHexPrefix", "isDecimalDigit", "isHexadecimalDigit", "parseInt", "parseFloat", "string", "stringIter", "isOfType", "downcast", "is", "instance", "random"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_std_CreateInstance()
	},
	func() any {
		return Hx_Obj_std_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_std_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "int":
			return Hx_Field_std_int
		case "isSpaceChar":
			return Hx_Field_std_isSpaceChar
		case "isHexPrefix":
			return Hx_Field_std_isHexPrefix
		case "isDecimalDigit":
			return Hx_Field_std_isDecimalDigit
		case "isHexadecimalDigit":
			return Hx_Field_std_isHexadecimalDigit
		case "parseInt":
			return Hx_Field_std_parseInt
		case "parseFloat":
			return Hx_Field_std_parseFloat
		case "string":
			return Hx_Field_std_string
		case "stringIter":
			return Hx_Field_std_stringIter
		case "isOfType":
			return Hx_Field_std_isOfType
		case "downcast":
			return Hx_Field_std_downcast
		case "is":
			return Hx_Field_std_is
		case "instance":
			return Hx_Field_std_instance
		case "random":
			return Hx_Field_std_random
		}
		return nil
	}
}

func (this *Hx_Obj_std) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_std_RTTI
}

func (this *Hx_Obj_std) Hx_As_Hx_Obj_std() *Hx_Obj_std {
	return this
}

func Hx_Field_std_int(x float64) int32 {
	return ((int32)(int32(x)))
}

func Hx_Field_std_isSpaceChar(code int32) bool {
	return (((code > 8) && (code < 14)) || (code == 32))
}

func Hx_Field_std_isHexPrefix(cur int32, next int32) bool {
	return ((cur == 48) && ((next == 120) || (next == 88)))
}

func Hx_Field_std_isDecimalDigit(code int32) bool {
	return ((48 <= code) && (code <= 57))
}

func Hx_Field_std_isHexadecimalDigit(code int32) bool {
	return ((((48 <= code) && (code <= 57)) || ((97 <= code) && (code <= 102))) || ((65 <= code) && (code <= 70)))
}

func Hx_Field_std_parseInt(x string) HxNullable[int32] {
	if x == HxStringNull {
		return HxNullable[int32]{}
	}

	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(x))
	var index int32
	_ = index
	index = 0
	for index < _hx_reserved_len {
		var code int32
		_ = code
		code = (Hx_Field_go_haxe_hxstring_charCodeAt(x, index)).Value
		if !(((code > 8) && (code < 14)) || (code == 32)) {
			break
		}

		index++
	}

	var isNegative bool
	_ = isNegative
	if index < _hx_reserved_len {
		var sign HxNullable[int32]
		_ = sign
		sign = (Hx_Field_go_haxe_hxstring_charCodeAt(x, index))
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (sign).Valid == true {
			_hx_tmp_1 = ((sign).Value == 45)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if (sign).Valid == true {
				_hx_tmp_2 = ((sign).Value == 43)
			} else {
				_hx_tmp_2 = false
			}

			_hx_tmp_0 = (_hx_tmp_2)
		}

		if _hx_tmp_0 {
			index++
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if (sign).Valid == true {
			_hx_tmp_3 = ((sign).Value == 45)
		} else {
			_hx_tmp_3 = false
		}

		isNegative = (_hx_tmp_3)
	} else {
		isNegative = false
	}

	var isHexadecimal bool
	_ = isHexadecimal
	if (index + 1) < _hx_reserved_len {
		var cur int32
		_ = cur
		cur = (Hx_Field_go_haxe_hxstring_charCodeAt(x, index)).Value
		var next int32
		_ = next
		next = (Hx_Field_go_haxe_hxstring_charCodeAt(x, (index + 1))).Value
		isHexadecimal = ((cur == 48) && ((next == 120) || (next == 88)))
	} else {
		isHexadecimal = false
	}

	if isHexadecimal {
		index = (index + 2)
	}

	var cur int32
	_ = cur
	cur = index
	if isHexadecimal {
		for cur < _hx_reserved_len {
			var code int32
			_ = code
			code = (Hx_Field_go_haxe_hxstring_charCodeAt(x, cur)).Value
			if !((((48 <= code) && (code <= 57)) || ((97 <= code) && (code <= 102))) || ((65 <= code) && (code <= 70))) {
				break
			}

			cur++
		}
	} else {
		for cur < _hx_reserved_len {
			var code int32
			_ = code
			code = (Hx_Field_go_haxe_hxstring_charCodeAt(x, cur)).Value
			if !((48 <= code) && (code <= 57)) {
				break
			}

			cur++
		}
	}

	var firstInvalidIndex int32
	_ = firstInvalidIndex
	firstInvalidIndex = cur
	if index == firstInvalidIndex {
		return HxNullable[int32]{}
	}

	var hx_result_7 struct {
		Error  error
		Result int64
	}
	_ = hx_result_7
	var _hx_tmp_5 string
	_ = _hx_tmp_5
	_hx_tmp_5 = x
	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	_hx_tmp_6 = index
	var _hx_tmp_4 string
	_ = _hx_tmp_4
	_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_5, _hx_tmp_6, HxNullable[int32]{Value: firstInvalidIndex, Valid: true}))
	var _hx_tmp_7 int
	_ = _hx_tmp_7
	if isHexadecimal {
		_hx_tmp_7 = int(((int32)(16)))
	} else {
		_hx_tmp_7 = int(((int32)(10)))
	}

	hx_result_7.Result, hx_result_7.Error = strconv.ParseInt(_hx_tmp_4, _hx_tmp_7, int(((int32)(32))))
	var this1 struct {
		Error  error
		Result int64
	}
	_ = this1
	this1 = hx_result_7
	var this2 int64
	_ = this2
	if this1.Error != nil {
		var e error
		_ = e
		e = this1.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r int64
		_ = r
		r = this1.Result
		this2 = r
	}

	var result int32
	_ = result
	result = ((int32)(this2))
	if isNegative {
		return HxNullable[int32]{Value: -result, Valid: true}
	} else {
		return HxNullable[int32]{Value: result, Valid: true}
	}
}

func Hx_Field_std_parseFloat(x string) float64 {
	var nan float64
	_ = nan
	nan = Hx_Field_math_NaN
	if x == HxStringNull {
		return nan
	}

	x = strings.TrimSpace(x)
	var y string
	_ = y
	y = ""
	var hadE bool
	_ = hadE
	hadE = false
	var hadDigit bool
	_ = hadDigit
	hadDigit = false
	var lastCharWasSign bool
	_ = lastCharWasSign
	lastCharWasSign = false
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(x))
	hx_label_0:
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			{
				var tmp__g1_1 HxNullable[int32]
				_ = tmp__g1_1
				tmp__g1_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(x, i))
				if (tmp__g1_1).Valid == false {
					if !hadDigit {
						return nan
					}

					break
				} else {
					switch (tmp__g1_1).Value {
					case 43, 45:
						{
							if lastCharWasSign {
								return nan
							} else {
								var _hx_tmp_1 string
								_ = _hx_tmp_1
								_hx_tmp_1 = y
								y = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_1, (Hx_Field_go_haxe_hxstring_charAt(x, i)))
								lastCharWasSign = true
							}
						}

					case 46:
						{
							var _hx_tmp_2 string
							_ = _hx_tmp_2
							_hx_tmp_2 = y
							y = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_2, (Hx_Field_go_haxe_hxstring_charAt(x, i)))
							lastCharWasSign = false
						}

					case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
						{
							var _hx_tmp_3 string
							_ = _hx_tmp_3
							_hx_tmp_3 = y
							y = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_3, (Hx_Field_go_haxe_hxstring_charAt(x, i)))
							hadDigit = true
							lastCharWasSign = false
						}

					case 69, 101:
						{
							if hadE {
								return nan
							}

							hadE = true
							lastCharWasSign = false
							var _hx_tmp_4 string
							_ = _hx_tmp_4
							_hx_tmp_4 = y
							y = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_4, (Hx_Field_go_haxe_hxstring_charAt(x, i)))
						}

					default:
						{
							if !hadDigit {
								return nan
							}

							break hx_label_0
						}
					}
				}
			}
		}
	}

	if y == "" {
		return nan
	}

	{
		var hx_try_return float64
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							hx_try_state = 1
							hx_try_return = nan
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_result_8 struct {
					Error  error
					Result float64
				}
				_ = hx_result_8
				var _hx_tmp_5 string
				_ = _hx_tmp_5
				_hx_tmp_5 = y
				hx_result_8.Result, hx_result_8.Error = strconv.ParseFloat(_hx_tmp_5, int(((int32)(64))))
				var this1 struct {
					Error  error
					Result float64
				}
				_ = this1
				this1 = hx_result_8
				if this1.Error != nil {
					var e error
					_ = e
					e = this1.Error
					panic(Hx_Field_haxe_exception_thrown(((any)(e))))
				} else {
					var r float64
					_ = r
					r = this1.Result
					hx_try_state = 1
					hx_try_return = r
					return
				}
			}
		}()

		return hx_try_return
	}
}

func Hx_Field_std_string(s any) string {
	return Hx_Field_std_stringIter(s, 0)
}

func Hx_Field_std_stringIter(s any, depth int32) string {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(s, nil))) {
		return "null"
	}

	if depth >= 5 {
		return "<...>"
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(s)
	if !value.IsValid() {
		return "null"
	}

	var array HxNullable[HxArrayDyn]
	_ = array
	array = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(value.Interface())
	if (array).Valid != false {
		var values HxArray[string]
		_ = values
		values = HxMakeArray[string]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = array.Value.Len()
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				var _hx_tmp_1 HxArray[string]
				_ = _hx_tmp_1
				_hx_tmp_1 = values
				_ = (Hx_Array_Push[string](_hx_tmp_1, Hx_Field_std_stringIter(array.Value.Get_Dyn(i), (depth+1))))
			}
		}

		return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("[", (Hx_Array_Join[string](values, ","))), "]")
	}

	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = kind
	if _hx_tmp_2 == reflect.Ptr {
		return Hx_Field_std_stringIter(value.Elem(), (depth + 1))
	}

	var _hx_tmp_3 reflect.Kind
	_ = _hx_tmp_3
	_hx_tmp_3 = kind
	if _hx_tmp_3 == reflect.Map {
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		var keys []reflect.Value
		_ = keys
		keys = value.MapKeys()
		var _hx_tmp_4 int
		_ = _hx_tmp_4
		_hx_tmp_4 = len(keys)
		if _hx_tmp_4 > int(((int32)(0))) {
			var self []reflect.Value
			_ = self
			self = keys
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[reflect.Value]
			_ = _g_array
			_g_array = HxMakeArray(self...)
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g_current
				if !(_hx_tmp_5 < _g_array.Len()) {
					break
				}

				var _hx_tmp_6 HxArray[reflect.Value]
				_ = _hx_tmp_6
				_hx_tmp_6 = _g_array
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _g_current
				_g_current = (_g_current + 1)
				var key reflect.Value
				_ = key
				key = _hx_tmp_6.Get(_hx_tmp_7)
				var _hx_tmp_8 reflect.Kind
				_ = _hx_tmp_8
				_hx_tmp_8 = key.Kind()
				if _hx_tmp_8 == reflect.String {
					if key.String() == "toString" {
						var _hx_tmp_9 any
						_ = _hx_tmp_9
						_hx_tmp_9 = value.MapIndex(key).Interface()
						return Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any]().Dyn()))
					}
				} else {
					break
				}
			}
		}

		func() func(x string) {
			var _hx_closure_src_50 func(x any)
			_ = _hx_closure_src_50
			_hx_closure_src_50 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_50(((any)(x)))
			}
		}()("[")
		var self []reflect.Value
		_ = self
		self = keys
		var _g HxArray[string]
		_ = _g
		_g = HxMakeArray[string]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[reflect.Value]
			_ = _g2
			_g2 = HxMakeArray(self...)
			for {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = _g1
				if !(_hx_tmp_10 < _g2.Len()) {
					break
				}

				var v reflect.Value
				_ = v
				v = _g2.Get(_g1)
				_g1++
				var _hx_tmp_11 HxArray[string]
				_ = _hx_tmp_11
				_hx_tmp_11 = _g
				var _hx_tmp_12 string
				_ = _hx_tmp_12
				_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("", Hx_Field_std_stringIter(v, (depth+1))), " => ")
				_ = (Hx_Array_Push[string](_hx_tmp_11, Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_12, Hx_Field_std_stringIter(value.MapIndex(v), (depth+1)))))
			}
		}

		func() func(x string) {
			var _hx_closure_src_51 func(x any)
			_ = _hx_closure_src_51
			_hx_closure_src_51 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_51(((any)(x)))
			}
		}()((Hx_Array_Join[string](_g, ", ")))
		func() func(x string) {
			var _hx_closure_src_52 func(x any)
			_ = _hx_closure_src_52
			_hx_closure_src_52 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_52(((any)(x)))
			}
		}()("]")
		return buf.VTable.Hx_Field_toString()
	}

	var _hx_tmp_13 reflect.Kind
	_ = _hx_tmp_13
	_hx_tmp_13 = kind
	if _hx_tmp_13 == reflect.Struct {
		var enumIndexMethod reflect.Value
		_ = enumIndexMethod
		enumIndexMethod = value.MethodByName("Hx_Field_enumIndex")
		var enumTypeMethod reflect.Value
		_ = enumTypeMethod
		enumTypeMethod = value.MethodByName("Hx_Field_enumType")
		var _hx_tmp_14 bool
		_ = _hx_tmp_14
		if enumIndexMethod.IsValid() {
			_hx_tmp_14 = enumTypeMethod.IsValid()
		} else {
			_hx_tmp_14 = false
		}

		if _hx_tmp_14 {
			var self HxArray[reflect.Value]
			_ = self
			self = HxMakeArray[reflect.Value]()
			var enumIndex int32
			_ = enumIndex
			enumIndex = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(enumIndexMethod.Call(self.Underlying())[0].Interface())))
			var self1 HxArray[reflect.Value]
			_ = self1
			self1 = HxMakeArray[reflect.Value]()
			var hx_dyncast_3 any
			_ = hx_dyncast_3
			hx_dyncast_3 = Hx_Field_go_haxe_hxdynamic_castClass(enumTypeMethod.Call(self1.Underlying())[0].Interface(), "Hx_Obj_go_haxe_hxenum")
			var _hx_tmp_15 *Hx_Obj_go_haxe_hxenum
			_ = _hx_tmp_15
			if hx_dyncast_3 != nil {
				_hx_tmp_15 = hx_dyncast_3.(interface{ Hx_As_Hx_Obj_go_haxe_hxenum() *Hx_Obj_go_haxe_hxenum }).Hx_As_Hx_Obj_go_haxe_hxenum()
			} else {
				_hx_tmp_15 = nil
			}

			var enumType *Hx_Obj_go_haxe_hxenum
			_ = enumType
			enumType = _hx_tmp_15
			var enumCtorName string
			_ = enumCtorName
			enumCtorName = enumType.Hx_Field_constructorNames.Get(enumIndex)
			var enumCtorCount int32
			_ = enumCtorCount
			enumCtorCount = enumType.Hx_Field_constructorArgCounts.Get(enumIndex)
			var values HxArray[string]
			_ = values
			values = HxMakeArray[string]()
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = ((int32)(value.NumField()))
				for _g < _g1 {
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					_hx_tmp_16 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_16
					var _hx_tmp_17 HxArray[string]
					_ = _hx_tmp_17
					_hx_tmp_17 = values
					_ = (Hx_Array_Push[string](_hx_tmp_17, Hx_Field_std_stringIter(value.Field(int(i)), (depth+1))))
				}
			}

			if enumCtorCount == 0 {
				return enumCtorName
			} else {
				var _hx_tmp_18 string
				_ = _hx_tmp_18
				_hx_tmp_18 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("", enumCtorName), "(")
				return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_18, (Hx_Array_Join[string](values, ","))), ")")
			}
		}

		var valid reflect.Value
		_ = valid
		valid = value.FieldByName("Valid")
		if valid.IsValid() {
			var val reflect.Value
			_ = val
			val = value.FieldByName("Value")
			var _hx_tmp_19 bool
			_ = _hx_tmp_19
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(valid.Interface(), ((any)(false))))) {
				_hx_tmp_19 = true
			} else {
				_hx_tmp_19 = !val.IsValid()
			}

			if _hx_tmp_19 {
				return "null"
			} else {
				return Hx_Field_std_stringIter(val.Interface(), (depth + 1))
			}
		}

		var vt reflect.Value
		_ = vt
		vt = value.FieldByName("VTable")
		if vt.IsValid() {
			var toStr reflect.Value
			_ = toStr
			toStr = vt.MethodByName("Hx_Field_toString")
			if toStr.IsValid() {
				var self HxArray[reflect.Value]
				_ = self
				self = HxMakeArray[reflect.Value]()
				return Hx_Field_std_stringIter(toStr.Call(self.Underlying())[0], (depth + 1))
			}
		}

		return fmt.Sprintf("%v", value.Interface())
	}

	var _hx_tmp_20 bool
	_ = _hx_tmp_20
	var _hx_tmp_21 reflect.Kind
	_ = _hx_tmp_21
	_hx_tmp_21 = kind
	if _hx_tmp_21 == reflect.Interface {
		_hx_tmp_20 = value.IsNil()
	} else {
		_hx_tmp_20 = false
	}

	if _hx_tmp_20 {
		return "null"
	}

	var _hx_tmp_22 bool
	_ = _hx_tmp_22
	var _hx_tmp_23 reflect.Kind
	_ = _hx_tmp_23
	_hx_tmp_23 = kind
	if _hx_tmp_23 == reflect.Float64 {
		_hx_tmp_22 = true
	} else {
		var _hx_tmp_24 reflect.Kind
		_ = _hx_tmp_24
		_hx_tmp_24 = kind
		var _hx_tmp_25 reflect.Kind
		_ = _hx_tmp_25
		_hx_tmp_25 = _hx_tmp_24
		_hx_tmp_22 = (_hx_tmp_25 == reflect.Float32)
	}

	if _hx_tmp_22 {
		return HxFloatToString(value.Float())
	}

	return fmt.Sprintf("%v", value.Interface())
}

func Hx_Field_std_isOfType(v any, t any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, nil))) {
		return false
	}

	var vt Hx_Enum_valuetype
	_ = vt
	vt = Hx_Field_type_typeof(v)
	switch t {
	case Hx_Obj_array_RTTI:
		{
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if vt.Hx_Field_enumIndex() == 7 {
				_hx_tmp_0 = (vt.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_array_RTTI)
			} else {
				_hx_tmp_0 = false
			}

			return _hx_tmp_0
		}

	case Hx_Abstract_bool_RTTI:
		{
			return (vt == Hx_Enum_valuetype_TBool{})
		}

	case Hx_Abstract_class_RTTI:
		{
			return Hx_Field_go_haxe_hxdynamic_isHxClass(v)
		}

	case Hx_Abstract_dynamic_RTTI:
		{
			return true
		}

	case Hx_Abstract_enum_RTTI:
		{
			return Hx_Field_go_haxe_hxdynamic_isHxEnum(v)
		}

	case Hx_Abstract_float_RTTI:
		{
			return ((vt == Hx_Enum_valuetype_TFloat{}) || (vt == Hx_Enum_valuetype_TInt{}))
		}

	case Hx_Abstract_int_RTTI:
		{
			switch vt.Hx_Field_enumIndex() {
			case 1:
				{
					return true
				}

			case 3:
				{
					var f float64
					_ = f
					f = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v)))
					var _hx_tmp_1 bool
					_ = _hx_tmp_1
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					var _hx_tmp_3 bool
					_ = _hx_tmp_3
					if Hx_Field_math_isFinite(f) {
						_hx_tmp_3 = (Hx_Field_math_ffloor(f) == f)
					} else {
						_hx_tmp_3 = false
					}

					if _hx_tmp_3 {
						var _hx_tmp_4 float64
						_ = _hx_tmp_4
						_hx_tmp_4 = float64(f)
						var _hx_tmp_5 float64
						_ = _hx_tmp_5
						_hx_tmp_5 = _hx_tmp_4
						_hx_tmp_2 = (_hx_tmp_5 >= float64(math.MinInt32))
					} else {
						_hx_tmp_2 = false
					}

					if _hx_tmp_2 {
						var _hx_tmp_6 float64
						_ = _hx_tmp_6
						_hx_tmp_6 = float64(f)
						var _hx_tmp_7 float64
						_ = _hx_tmp_7
						_hx_tmp_7 = _hx_tmp_6
						_hx_tmp_1 = (_hx_tmp_7 <= float64(math.MaxInt32))
					} else {
						_hx_tmp_1 = false
					}

					return _hx_tmp_1
				}

			default:
				{
					return false
				}
			}
		}

	case Hx_Obj_string_RTTI:
		{
			var _hx_tmp_8 bool
			_ = _hx_tmp_8
			if vt.Hx_Field_enumIndex() == 7 {
				_hx_tmp_8 = (vt.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI)
			} else {
				_hx_tmp_8 = false
			}

			return _hx_tmp_8
		}

	default:
		{
			switch vt.Hx_Field_enumIndex() {
			case 7:
				{
					var q *Hx_Obj_go_haxe_hxclass
					_ = q
					q = vt.(Hx_Enum_valuetype_TClass).Hx_Field_c
					if !Hx_Field_go_haxe_hxdynamic_isHxClass(t) {
						return false
					}

					var hx_dyncast_4 any
					_ = hx_dyncast_4
					hx_dyncast_4 = Hx_Field_go_haxe_hxdynamic_castClass(t, "Hx_Obj_go_haxe_hxclass")
					var _hx_tmp_9 *Hx_Obj_go_haxe_hxclass
					_ = _hx_tmp_9
					if hx_dyncast_4 != nil {
						_hx_tmp_9 = hx_dyncast_4.(interface {
							Hx_As_Hx_Obj_go_haxe_hxclass() *Hx_Obj_go_haxe_hxclass
						}).Hx_As_Hx_Obj_go_haxe_hxclass()
					} else {
						_hx_tmp_9 = nil
					}

					var target *Hx_Obj_go_haxe_hxclass
					_ = target
					target = _hx_tmp_9
					if target == nil {
						return false
					}

					var queue HxArray[*Hx_Obj_go_haxe_hxclass]
					_ = queue
					queue = HxMakeArray[*Hx_Obj_go_haxe_hxclass](q)
					for {
						if !(queue.Len() > 0) {
							break
						}

						var current *Hx_Obj_go_haxe_hxclass
						_ = current
						current = (Hx_Array_Shift[*Hx_Obj_go_haxe_hxclass](queue).Value)
						if current == nil {
							continue
						}

						if current.Hx_Field_name == target.Hx_Field_name {
							return true
						}

						if current.Hx_Field_superClass != nil {
							_ = (Hx_Array_Push[*Hx_Obj_go_haxe_hxclass](queue, current.Hx_Field_superClass))
						}

						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 HxArray[*Hx_Obj_go_haxe_hxclass]
							_ = _g1
							_g1 = current.Hx_Field_interfaces
							for {
								var _hx_tmp_10 int32
								_ = _hx_tmp_10
								_hx_tmp_10 = _g
								if !(_hx_tmp_10 < _g1.Len()) {
									break
								}

								var iface *Hx_Obj_go_haxe_hxclass
								_ = iface
								iface = _g1.Get(_g)
								_g++
								_ = (Hx_Array_Push[*Hx_Obj_go_haxe_hxclass](queue, iface))
							}
						}
					}

					return false
				}

			case 8:
				{
					var q *Hx_Obj_go_haxe_hxenum
					_ = q
					q = vt.(Hx_Enum_valuetype_TEnum).Hx_Field_e
					if !Hx_Field_go_haxe_hxdynamic_isHxEnum(t) {
						return false
					}

					var hx_dyncast_5 any
					_ = hx_dyncast_5
					hx_dyncast_5 = Hx_Field_go_haxe_hxdynamic_castClass(t, "Hx_Obj_go_haxe_hxenum")
					var _hx_tmp_11 *Hx_Obj_go_haxe_hxenum
					_ = _hx_tmp_11
					if hx_dyncast_5 != nil {
						_hx_tmp_11 = hx_dyncast_5.(interface{ Hx_As_Hx_Obj_go_haxe_hxenum() *Hx_Obj_go_haxe_hxenum }).Hx_As_Hx_Obj_go_haxe_hxenum()
					} else {
						_hx_tmp_11 = nil
					}

					var v1 *Hx_Obj_go_haxe_hxenum
					_ = v1
					v1 = _hx_tmp_11
					var v2 *Hx_Obj_go_haxe_hxenum
					_ = v2
					v2 = q
					return (((v1 != nil) && (v2 != nil)) && (v1.Hx_Field_name == v2.Hx_Field_name))
				}

			default:
				{
					return false
				}
			}
		}
	}
}

func Hx_Field_std_downcast(value any, c *Hx_Obj_go_haxe_hxclass) any {
	if Hx_Field_std_isOfType(value, c) {
		return value
	} else {
		return nil
	}
}

func Hx_Field_std_is(v any, t any) bool {
	return Hx_Field_std_isOfType(v, t)
}

func Hx_Field_std_instance(value any, c *Hx_Obj_go_haxe_hxclass) any {
	return Hx_Field_std_downcast(value, c)
}

func Hx_Field_std_random(x int32) int32 {
	if x <= 0 {
		return 0
	} else {
		return Hx_Field_std_int((Hx_Field_math_random() * ((float64)(x))))
	}
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_go_haxe_hxclass]()
	hxRegisterArrayConstructor[reflect.Value]()
	hxRegisterArrayConstructor[string]()
}
