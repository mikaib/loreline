package core

import "reflect"
import "unicode/utf8"
import "math"
import "strings"
import "unsafe"

var Hx_Obj_go_haxe__hxdynamic_unsafe_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe._HxDynamic.Unsafe",
	HxMakeArray[string]("pointer", "add", "sizeof", "slice"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		panic("Cannot dynamically create instance of extern class")
	},
	func() any {
		panic("Cannot dynamically create instance of extern class")
	},
	nil,
	func(name string) any {
		return nil
	},
)

var Hx_Obj_go_haxe_hxdynamicarray_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxDynamicArray",
	HxMakeArray[string](),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		panic("Cannot dynamically create instance of extern class")
	},
	func() any {
		panic("Cannot dynamically create instance of extern class")
	},
	nil,
	func(name string) any {
		return nil
	},
)

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxdynamic() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxdynamic() {
	Hx_Field_go_haxe_hxdynamic_Null = Hx_Init_Hx_Field_go_haxe_hxdynamic_Null()
}

type Hx_Obj_VTable_go_haxe_hxdynamic interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxdynamic struct {
	VTable Hx_Obj_VTable_go_haxe_hxdynamic `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go_haxe_hxdynamic_CreateEmptyInstance() *Hx_Obj_go_haxe_hxdynamic {
	obj := &Hx_Obj_go_haxe_hxdynamic{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go_haxe_hxdynamic_CreateInstance() *Hx_Obj_go_haxe_hxdynamic {
	obj := Hx_Obj_go_haxe_hxdynamic_CreateEmptyInstance()
	return obj
}

var Hx_Obj_go_haxe_hxdynamic_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxDynamic",
	HxMakeArray[string]("Null", "not", "increment", "decrement", "negate", "bitnot", "and", "or", "jointKind", "toAnySlice", "formatField", "isNilValue", "isNull", "equals", "enumEquals", "nequals", "lt", "gtequals", "gt", "ltequals", "add", "stringConcat", "subtract", "multiply", "divide", "modulo", "bitand", "bitor", "bitxor", "lbitshift", "rbitshift", "urbitshift", "toString", "toBool", "valueToInt", "isClass", "classVTable", "convertToType", "createNewClosure", "call", "toInt", "valueToInt64", "toInt64", "valueToFloat", "toFloat", "valueToClass", "toClass", "castClass", "isNullableType", "unwrapNullable", "getField", "tryHxClass", "tryHxEnum", "hxEnumConstructor", "isHxClass", "isHxEnum", "isHxEnumValue", "asClass", "asEnum", "hxClassStaticValue", "getFieldFrom", "setField", "tryDynamicArray", "setArrayIndex", "getArrayIndex", "getArrayLength", "ensureValue", "valueToAssign", "ensureConcreteValue", "ensureInterface", "funcReceiver"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxdynamic_CreateInstance()
	},
	func() any {
		return Hx_Obj_go_haxe_hxdynamic_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxdynamic_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Null":
			return &Hx_Field_go_haxe_hxdynamic_Null
		case "not":
			return Hx_Field_go_haxe_hxdynamic_not
		case "increment":
			return Hx_Field_go_haxe_hxdynamic_increment
		case "decrement":
			return Hx_Field_go_haxe_hxdynamic_decrement
		case "negate":
			return Hx_Field_go_haxe_hxdynamic_negate
		case "bitnot":
			return Hx_Field_go_haxe_hxdynamic_bitnot
		case "and":
			return Hx_Field_go_haxe_hxdynamic_and
		case "or":
			return Hx_Field_go_haxe_hxdynamic_or
		case "jointKind":
			return Hx_Field_go_haxe_hxdynamic_jointKind
		case "toAnySlice":
			return Hx_Field_go_haxe_hxdynamic_toAnySlice
		case "formatField":
			return Hx_Field_go_haxe_hxdynamic_formatField
		case "isNilValue":
			return Hx_Field_go_haxe_hxdynamic_isNilValue
		case "isNull":
			return Hx_Field_go_haxe_hxdynamic_isNull
		case "equals":
			return Hx_Field_go_haxe_hxdynamic_equals
		case "enumEquals":
			return Hx_Field_go_haxe_hxdynamic_enumEquals
		case "nequals":
			return Hx_Field_go_haxe_hxdynamic_nequals
		case "lt":
			return Hx_Field_go_haxe_hxdynamic_lt
		case "gtequals":
			return Hx_Field_go_haxe_hxdynamic_gtequals
		case "gt":
			return Hx_Field_go_haxe_hxdynamic_gt
		case "ltequals":
			return Hx_Field_go_haxe_hxdynamic_ltequals
		case "add":
			return Hx_Field_go_haxe_hxdynamic_add
		case "stringConcat":
			return Hx_Field_go_haxe_hxdynamic_stringConcat
		case "subtract":
			return Hx_Field_go_haxe_hxdynamic_subtract
		case "multiply":
			return Hx_Field_go_haxe_hxdynamic_multiply
		case "divide":
			return Hx_Field_go_haxe_hxdynamic_divide
		case "modulo":
			return Hx_Field_go_haxe_hxdynamic_modulo
		case "bitand":
			return Hx_Field_go_haxe_hxdynamic_bitand
		case "bitor":
			return Hx_Field_go_haxe_hxdynamic_bitor
		case "bitxor":
			return Hx_Field_go_haxe_hxdynamic_bitxor
		case "lbitshift":
			return Hx_Field_go_haxe_hxdynamic_lbitshift
		case "rbitshift":
			return Hx_Field_go_haxe_hxdynamic_rbitshift
		case "urbitshift":
			return Hx_Field_go_haxe_hxdynamic_urbitshift
		case "toString":
			return Hx_Field_go_haxe_hxdynamic_toString
		case "toBool":
			return Hx_Field_go_haxe_hxdynamic_toBool
		case "valueToInt":
			return Hx_Field_go_haxe_hxdynamic_valueToInt
		case "isClass":
			return Hx_Field_go_haxe_hxdynamic_isClass
		case "classVTable":
			return Hx_Field_go_haxe_hxdynamic_classVTable
		case "convertToType":
			return Hx_Field_go_haxe_hxdynamic_convertToType
		case "createNewClosure":
			return Hx_Field_go_haxe_hxdynamic_createNewClosure
		case "call":
			return Hx_Field_go_haxe_hxdynamic_call
		case "toInt":
			return Hx_Field_go_haxe_hxdynamic_toInt
		case "valueToInt64":
			return Hx_Field_go_haxe_hxdynamic_valueToInt64
		case "toInt64":
			return Hx_Field_go_haxe_hxdynamic_toInt64
		case "valueToFloat":
			return Hx_Field_go_haxe_hxdynamic_valueToFloat
		case "toFloat":
			return Hx_Field_go_haxe_hxdynamic_toFloat
		case "valueToClass":
			return Hx_Field_go_haxe_hxdynamic_valueToClass
		case "toClass":
			return Hx_Field_go_haxe_hxdynamic_toClass
		case "castClass":
			return Hx_Field_go_haxe_hxdynamic_castClass
		case "isNullableType":
			return Hx_Field_go_haxe_hxdynamic_isNullableType
		case "unwrapNullable":
			return Hx_Field_go_haxe_hxdynamic_unwrapNullable
		case "getField":
			return Hx_Field_go_haxe_hxdynamic_getField
		case "tryHxClass":
			return Hx_Field_go_haxe_hxdynamic_tryHxClass
		case "tryHxEnum":
			return Hx_Field_go_haxe_hxdynamic_tryHxEnum
		case "hxEnumConstructor":
			return Hx_Field_go_haxe_hxdynamic_hxEnumConstructor
		case "isHxClass":
			return Hx_Field_go_haxe_hxdynamic_isHxClass
		case "isHxEnum":
			return Hx_Field_go_haxe_hxdynamic_isHxEnum
		case "isHxEnumValue":
			return Hx_Field_go_haxe_hxdynamic_isHxEnumValue
		case "asClass":
			return Hx_Field_go_haxe_hxdynamic_asClass
		case "asEnum":
			return Hx_Field_go_haxe_hxdynamic_asEnum
		case "hxClassStaticValue":
			return Hx_Field_go_haxe_hxdynamic_hxClassStaticValue
		case "getFieldFrom":
			return Hx_Field_go_haxe_hxdynamic_getFieldFrom
		case "setField":
			return Hx_Field_go_haxe_hxdynamic_setField
		case "tryDynamicArray":
			return Hx_Field_go_haxe_hxdynamic_tryDynamicArray
		case "setArrayIndex":
			return Hx_Field_go_haxe_hxdynamic_setArrayIndex
		case "getArrayIndex":
			return Hx_Field_go_haxe_hxdynamic_getArrayIndex
		case "getArrayLength":
			return Hx_Field_go_haxe_hxdynamic_getArrayLength
		case "ensureValue":
			return Hx_Field_go_haxe_hxdynamic_ensureValue
		case "valueToAssign":
			return Hx_Field_go_haxe_hxdynamic_valueToAssign
		case "ensureConcreteValue":
			return Hx_Field_go_haxe_hxdynamic_ensureConcreteValue
		case "ensureInterface":
			return Hx_Field_go_haxe_hxdynamic_ensureInterface
		case "funcReceiver":
			return Hx_Field_go_haxe_hxdynamic_funcReceiver
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxdynamic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxdynamic_RTTI
}

func (this *Hx_Obj_go_haxe_hxdynamic) Hx_As_Hx_Obj_go_haxe_hxdynamic() *Hx_Obj_go_haxe_hxdynamic {
	return this
}

var Hx_Field_go_haxe_hxdynamic_Null reflect.Value

func Hx_Init_Hx_Field_go_haxe_hxdynamic_Null() reflect.Value {
	return reflect.ValueOf(nil)
}

func Hx_Field_go_haxe_hxdynamic_not(d any) bool {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 reflect.Kind
	_ = _hx_tmp_0
	_hx_tmp_0 = dV.Kind()
	if _hx_tmp_0 == reflect.Bool {
		var dB bool
		_ = dB
		dB = dV.Bool()
		return !dB
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.not value invalid: ", Hx_Field_std_string(d))))))
}

func Hx_Field_go_haxe_hxdynamic_increment(d any) any {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dV.CanInt() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = dV.CanUint()
	}

	if _hx_tmp_0 {
		return ((any)((Hx_Field_go_haxe_hxdynamic_valueToInt(dV) + 1)))
	} else {
		if dV.CanFloat() {
			return ((any)((Hx_Field_go_haxe_hxdynamic_valueToFloat(dV) + float64(1.0))))
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.increment value invalid: ", Hx_Field_std_string(d))))))
}

func Hx_Field_go_haxe_hxdynamic_decrement(d any) any {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dV.CanInt() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = dV.CanUint()
	}

	if _hx_tmp_0 {
		return ((any)((Hx_Field_go_haxe_hxdynamic_valueToInt(dV) - 1)))
	} else {
		if dV.CanFloat() {
			return ((any)((Hx_Field_go_haxe_hxdynamic_valueToFloat(dV) - float64(1.0))))
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.decrement value invalid: ", Hx_Field_std_string(d))))))
}

func Hx_Field_go_haxe_hxdynamic_negate(d any) any {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dV.CanInt() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = dV.CanUint()
	}

	if _hx_tmp_0 {
		return ((any)((0 - Hx_Field_go_haxe_hxdynamic_valueToInt(dV))))
	} else {
		if dV.CanFloat() {
			return ((any)((float64(0.0) - Hx_Field_go_haxe_hxdynamic_valueToFloat(dV))))
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.negate value invalid: ", Hx_Field_std_string(d))))))
}

func Hx_Field_go_haxe_hxdynamic_bitnot(d any) any {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dV.CanInt() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = dV.CanUint()
	}

	if _hx_tmp_0 {
		return ((any)(^Hx_Field_go_haxe_hxdynamic_valueToInt(dV)))
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.bitnot value invalid: ", Hx_Field_std_string(d))))))
}

func Hx_Field_go_haxe_hxdynamic_and(a any, b any) bool {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = aV.Kind()
	if _hx_tmp_1 == reflect.Bool {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = bV.Kind()
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if aV.Bool() {
			_hx_tmp_4 = bV.Bool()
		} else {
			_hx_tmp_4 = false
		}

		return _hx_tmp_4
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.and invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_or(a any, b any) bool {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = aV.Kind()
	if _hx_tmp_1 == reflect.Bool {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = bV.Kind()
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if aV.Bool() {
			_hx_tmp_4 = true
		} else {
			_hx_tmp_4 = bV.Bool()
		}

		return _hx_tmp_4
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.or invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_jointKind(aV reflect.Value, bV reflect.Value) reflect.Kind {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if aV.CanInt() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = aV.CanUint()
	}

	var avCi bool
	_ = avCi
	avCi = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if bV.CanInt() {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = bV.CanUint()
	}

	var bvCi bool
	_ = bvCi
	bvCi = _hx_tmp_1
	if avCi && bvCi {
		return reflect.Int
	} else {
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if aV.CanFloat() {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = avCi
		}

		if _hx_tmp_3 {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if bV.CanFloat() {
				_hx_tmp_4 = true
			} else {
				_hx_tmp_4 = bvCi
			}

			_hx_tmp_2 = (_hx_tmp_4)
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			return reflect.Float64
		} else {
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			var _hx_tmp_6 reflect.Kind
			_ = _hx_tmp_6
			_hx_tmp_6 = aV.Kind()
			if _hx_tmp_6 == reflect.String {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_7 reflect.Kind
				_ = _hx_tmp_7
				_hx_tmp_7 = bV.Kind()
				var _hx_tmp_8 reflect.Kind
				_ = _hx_tmp_8
				_hx_tmp_8 = _hx_tmp_7
				_hx_tmp_5 = (_hx_tmp_8 == reflect.String)
			}

			if _hx_tmp_5 {
				return reflect.String
			}
		}
	}

	return reflect.Invalid
}

func Hx_Field_go_haxe_hxdynamic_toAnySlice(v any) []any {
	var arr HxNullable[HxArrayDyn]
	_ = arr
	arr = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(v)
	if (arr).Valid != false {
		return arr.Value.Underlying_Dyn()
	}

	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = Hx_Field_go_haxe_hxdynamic_getArrayLength(v)
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{}

	var this1 []any
	_ = this1
	var length1 HxNullable[int32]
	_ = length1
	length1 = length
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length1).Valid != false {
		_hx_tmp_0 = (length1).Value
	} else {
		_hx_tmp_0 = 0
	}

	this1 = make([]any, _hx_tmp_0)
	var slice []any
	_ = slice
	slice = this1
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
			var v1 any
			_ = v1
			v1 = Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getArrayIndex(v, i))
			slice = append(slice, v1)
		}
	}

	return slice
}

func Hx_Field_go_haxe_hxdynamic_formatField(name string) string {
	if int32(utf8.RuneCountInString(name)) == 0 {
		return name
	}

	return Hx_Field_go_haxe_hxdynamic_stringConcat("Hx_Field_", name)
}

func Hx_Field_go_haxe_hxdynamic_isNilValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	var k reflect.Kind
	_ = k
	k = v.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	var _hx_tmp_6 reflect.Kind
	_ = _hx_tmp_6
	_hx_tmp_6 = k
	if _hx_tmp_6 == reflect.Ptr {
		_hx_tmp_5 = true
	} else {
		var _hx_tmp_7 reflect.Kind
		_ = _hx_tmp_7
		_hx_tmp_7 = k
		var _hx_tmp_8 reflect.Kind
		_ = _hx_tmp_8
		_hx_tmp_8 = _hx_tmp_7
		_hx_tmp_5 = (_hx_tmp_8 == reflect.Interface)
	}

	if _hx_tmp_5 {
		_hx_tmp_4 = true
	} else {
		var _hx_tmp_9 reflect.Kind
		_ = _hx_tmp_9
		_hx_tmp_9 = k
		var _hx_tmp_10 reflect.Kind
		_ = _hx_tmp_10
		_hx_tmp_10 = _hx_tmp_9
		_hx_tmp_4 = (_hx_tmp_10 == reflect.Map)
	}

	if _hx_tmp_4 {
		_hx_tmp_3 = true
	} else {
		var _hx_tmp_11 reflect.Kind
		_ = _hx_tmp_11
		_hx_tmp_11 = k
		var _hx_tmp_12 reflect.Kind
		_ = _hx_tmp_12
		_hx_tmp_12 = _hx_tmp_11
		_hx_tmp_3 = (_hx_tmp_12 == reflect.Slice)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_13 reflect.Kind
		_ = _hx_tmp_13
		_hx_tmp_13 = k
		var _hx_tmp_14 reflect.Kind
		_ = _hx_tmp_14
		_hx_tmp_14 = _hx_tmp_13
		_hx_tmp_2 = (_hx_tmp_14 == reflect.Func)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_15 reflect.Kind
		_ = _hx_tmp_15
		_hx_tmp_15 = k
		var _hx_tmp_16 reflect.Kind
		_ = _hx_tmp_16
		_hx_tmp_16 = _hx_tmp_15
		_hx_tmp_1 = (_hx_tmp_16 == reflect.Chan)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_17 reflect.Kind
		_ = _hx_tmp_17
		_hx_tmp_17 = k
		var _hx_tmp_18 reflect.Kind
		_ = _hx_tmp_18
		_hx_tmp_18 = _hx_tmp_17
		_hx_tmp_0 = (_hx_tmp_18 == reflect.UnsafePointer)
	}

	if _hx_tmp_0 {
		return v.IsNil()
	}

	var _hx_tmp_19 reflect.Kind
	_ = _hx_tmp_19
	_hx_tmp_19 = k
	if _hx_tmp_19 == reflect.String {
		return Hx_Field_go_haxe_hxstringnull_isNull(v.String())
	}

	return false
}

func Hx_Field_go_haxe_hxdynamic_isNull(x any) bool {
	if x == nil {
		return true
	}

	return Hx_Field_go_haxe_hxdynamic_isNilValue(Hx_Field_go_haxe_hxdynamic_ensureValue(x))
}

func Hx_Field_go_haxe_hxdynamic_equals(a any, b any) bool {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_unwrapNullable(Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a))
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_unwrapNullable(Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b))
	var aN bool
	_ = aN
	aN = Hx_Field_go_haxe_hxdynamic_isNilValue(aV)
	var bN bool
	_ = bN
	bN = Hx_Field_go_haxe_hxdynamic_isNilValue(bV)
	if aN || bN {
		if aN && bN {
			return true
		} else {
			return false
		}
	}

	var aK reflect.Kind
	_ = aK
	aK = aV.Kind()
	var bK reflect.Kind
	_ = bK
	bK = bV.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = aK
	if _hx_tmp_1 == reflect.Bool {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = bK
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Bool)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = aK
		if _hx_tmp_5 == reflect.Bool {
			var _hx_tmp_6 reflect.Kind
			_ = _hx_tmp_6
			_hx_tmp_6 = bK
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_4 = (_hx_tmp_7 == reflect.Bool)
		} else {
			_hx_tmp_4 = false
		}

		if _hx_tmp_4 {
			var _hx_tmp_8 bool
			_ = _hx_tmp_8
			_hx_tmp_8 = aV.Bool()
			return (_hx_tmp_8 == bV.Bool())
		} else {
			return false
		}
	}

	var _hx_tmp_9 bool
	_ = _hx_tmp_9
	var _hx_tmp_10 reflect.Kind
	_ = _hx_tmp_10
	_hx_tmp_10 = aK
	if _hx_tmp_10 == reflect.Func {
		var _hx_tmp_11 reflect.Kind
		_ = _hx_tmp_11
		_hx_tmp_11 = bK
		var _hx_tmp_12 reflect.Kind
		_ = _hx_tmp_12
		_hx_tmp_12 = _hx_tmp_11
		_hx_tmp_9 = (_hx_tmp_12 == reflect.Func)
	} else {
		_hx_tmp_9 = false
	}

	if _hx_tmp_9 {
		var _hx_tmp_13 uintptr
		_ = _hx_tmp_13
		_hx_tmp_13 = aV.Pointer()
		if _hx_tmp_13 != bV.Pointer() {
			return false
		}

		var _hx_tmp_14 uintptr
		_ = _hx_tmp_14
		_hx_tmp_14 = Hx_Field_go_haxe_hxdynamic_funcReceiver(a)
		return (_hx_tmp_14 == Hx_Field_go_haxe_hxdynamic_funcReceiver(b))
	}

	var _hx_tmp_15 bool
	_ = _hx_tmp_15
	if Hx_Field_go_haxe_hxdynamic_isHxEnumValue(a) {
		_hx_tmp_15 = Hx_Field_go_haxe_hxdynamic_isHxEnumValue(b)
	} else {
		_hx_tmp_15 = false
	}

	if _hx_tmp_15 {
		return Hx_Field_go_haxe_hxdynamic_enumEquals(a, b)
	}

	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_16 bool
	_ = _hx_tmp_16
	var _hx_tmp_17 reflect.Kind
	_ = _hx_tmp_17
	_hx_tmp_17 = k
	if _hx_tmp_17 == reflect.Int32 {
		_hx_tmp_16 = true
	} else {
		var _hx_tmp_18 reflect.Kind
		_ = _hx_tmp_18
		_hx_tmp_18 = k
		var _hx_tmp_19 reflect.Kind
		_ = _hx_tmp_19
		_hx_tmp_19 = _hx_tmp_18
		_hx_tmp_16 = (_hx_tmp_19 == reflect.Int)
	}

	if _hx_tmp_16 {
		var _hx_tmp_20 int32
		_ = _hx_tmp_20
		_hx_tmp_20 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return (_hx_tmp_20 == Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
	} else {
		var _hx_tmp_21 reflect.Kind
		_ = _hx_tmp_21
		_hx_tmp_21 = k
		if _hx_tmp_21 == reflect.Float64 {
			var _hx_tmp_22 float64
			_ = _hx_tmp_22
			_hx_tmp_22 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return (_hx_tmp_22 == Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
		} else {
			var _hx_tmp_23 reflect.Kind
			_ = _hx_tmp_23
			_hx_tmp_23 = k
			if _hx_tmp_23 == reflect.String {
				var _hx_tmp_24 string
				_ = _hx_tmp_24
				_hx_tmp_24 = Hx_Field_go_haxe_hxdynamic_toString(a)
				return (_hx_tmp_24 == Hx_Field_go_haxe_hxdynamic_toString(b))
			} else {
				var _hx_tmp_25 bool
				_ = _hx_tmp_25
				if Hx_Field_go_haxe_hxdynamic_isClass(aV) {
					_hx_tmp_25 = Hx_Field_go_haxe_hxdynamic_isClass(bV)
				} else {
					_hx_tmp_25 = false
				}

				if _hx_tmp_25 {
					var aVt reflect.Value
					_ = aVt
					aVt = Hx_Field_go_haxe_hxdynamic_classVTable(aV)
					var bVt reflect.Value
					_ = bVt
					bVt = Hx_Field_go_haxe_hxdynamic_classVTable(bV)
					var _hx_tmp_26 bool
					_ = _hx_tmp_26
					if aVt.IsValid() {
						_hx_tmp_26 = bVt.IsValid()
					} else {
						_hx_tmp_26 = false
					}

					if _hx_tmp_26 {
						var _hx_tmp_27 uintptr
						_ = _hx_tmp_27
						_hx_tmp_27 = aVt.Pointer()
						return (_hx_tmp_27 == bVt.Pointer())
					}
				}

				var _hx_tmp_28 bool
				_ = _hx_tmp_28
				var _hx_tmp_29 bool
				_ = _hx_tmp_29
				var _hx_tmp_30 bool
				_ = _hx_tmp_30
				var _hx_tmp_31 bool
				_ = _hx_tmp_31
				var _hx_tmp_32 reflect.Kind
				_ = _hx_tmp_32
				_hx_tmp_32 = aK
				if _hx_tmp_32 == reflect.Ptr {
					_hx_tmp_31 = true
				} else {
					var _hx_tmp_33 reflect.Kind
					_ = _hx_tmp_33
					_hx_tmp_33 = aK
					var _hx_tmp_34 reflect.Kind
					_ = _hx_tmp_34
					_hx_tmp_34 = _hx_tmp_33
					_hx_tmp_31 = (_hx_tmp_34 == reflect.Map)
				}

				if _hx_tmp_31 {
					_hx_tmp_30 = true
				} else {
					var _hx_tmp_35 reflect.Kind
					_ = _hx_tmp_35
					_hx_tmp_35 = aK
					var _hx_tmp_36 reflect.Kind
					_ = _hx_tmp_36
					_hx_tmp_36 = _hx_tmp_35
					_hx_tmp_30 = (_hx_tmp_36 == reflect.Chan)
				}

				if _hx_tmp_30 {
					_hx_tmp_29 = true
				} else {
					var _hx_tmp_37 reflect.Kind
					_ = _hx_tmp_37
					_hx_tmp_37 = aK
					var _hx_tmp_38 reflect.Kind
					_ = _hx_tmp_38
					_hx_tmp_38 = _hx_tmp_37
					_hx_tmp_29 = (_hx_tmp_38 == reflect.UnsafePointer)
				}

				if _hx_tmp_29 {
					_hx_tmp_28 = true
				} else {
					var _hx_tmp_39 reflect.Kind
					_ = _hx_tmp_39
					_hx_tmp_39 = aK
					var _hx_tmp_40 reflect.Kind
					_ = _hx_tmp_40
					_hx_tmp_40 = _hx_tmp_39
					_hx_tmp_28 = (_hx_tmp_40 == reflect.Slice)
				}

				if _hx_tmp_28 {
					var _hx_tmp_41 bool
					_ = _hx_tmp_41
					var _hx_tmp_42 reflect.Type
					_ = _hx_tmp_42
					_hx_tmp_42 = aV.Type()
					if _hx_tmp_42 == bV.Type() {
						var _hx_tmp_43 uintptr
						_ = _hx_tmp_43
						_hx_tmp_43 = aV.Pointer()
						var _hx_tmp_44 uintptr
						_ = _hx_tmp_44
						_hx_tmp_44 = _hx_tmp_43
						_hx_tmp_41 = (_hx_tmp_44 == bV.Pointer())
					} else {
						_hx_tmp_41 = false
					}

					return _hx_tmp_41
				}

				var _hx_tmp_45 any
				_ = _hx_tmp_45
				_hx_tmp_45 = aV.Interface()
				return reflect.DeepEqual(_hx_tmp_45, bV.Interface())
			}
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_enumEquals(a any, b any) bool {
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(a, "enumType")
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn())
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(b, "enumType")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())))) {
		return false
	}

	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(a, "enumIndex")
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn())
	var _hx_tmp_5 any
	_ = _hx_tmp_5
	_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(b, "enumIndex")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn())))) {
		return false
	}

	var _hx_tmp_6 any
	_ = _hx_tmp_6
	_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(a, "enumParams")
	var ap HxArray[any]
	_ = ap
	ap = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any]().Dyn()))
	var _hx_tmp_7 any
	_ = _hx_tmp_7
	_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(b, "enumParams")
	var bp HxArray[any]
	_ = bp
	bp = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_7, HxMakeArray[any]().Dyn()))
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(ap, "length"))))
		for _g < _g1 {
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_8
			if !Hx_Field_go_haxe_hxdynamic_equals(ap.Get(i), bp.Get(i)) {
				return false
			}
		}
	}

	return true
}

func Hx_Field_go_haxe_hxdynamic_nequals(a any, b any) bool {
	return !Hx_Field_go_haxe_hxdynamic_equals(a, b)
}

func Hx_Field_go_haxe_hxdynamic_lt(a any, b any) bool {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return (_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return (_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
		} else {
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = k
			if _hx_tmp_7 == reflect.String {
				var _hx_tmp_8 string
				_ = _hx_tmp_8
				_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_toString(a)
				return (_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_toString(b))
			} else {
				var _hx_tmp_9 string
				_ = _hx_tmp_9
				_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.lt invalid operands: ", aV.String()), " and ")
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_9, bV.String())))))
			}
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_gtequals(a any, b any) bool {
	return !Hx_Field_go_haxe_hxdynamic_lt(a, b)
}

func Hx_Field_go_haxe_hxdynamic_gt(a any, b any) bool {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return (_hx_tmp_4 > Hx_Field_go_haxe_hxdynamic_valueToInt(bV))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return (_hx_tmp_6 > Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
		} else {
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = k
			if _hx_tmp_7 == reflect.String {
				var _hx_tmp_8 string
				_ = _hx_tmp_8
				_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_toString(a)
				return (_hx_tmp_8 > Hx_Field_go_haxe_hxdynamic_toString(b))
			} else {
				var _hx_tmp_9 string
				_ = _hx_tmp_9
				_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.gt invalid operands: ", aV.String()), " and ")
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_9, bV.String())))))
			}
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_ltequals(a any, b any) bool {
	return !Hx_Field_go_haxe_hxdynamic_gt(a, b)
}

func Hx_Field_go_haxe_hxdynamic_add(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 + Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return ((any)((_hx_tmp_6 + Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
		} else {
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = k
			if _hx_tmp_7 == reflect.String {
				var aS string
				_ = aS
				aS = Hx_Field_go_haxe_hxdynamic_toString(a)
				var bS string
				_ = bS
				bS = Hx_Field_go_haxe_hxdynamic_toString(b)
				return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(aS, bS)))
			} else {
				var _hx_tmp_8 string
				_ = _hx_tmp_8
				_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.add invalid operands: ", aV.String()), " and ")
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_8, bV.String())))))
			}
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_stringConcat(a string, b string) string {
	return HxStringConcat(a, b)
}

func Hx_Field_go_haxe_hxdynamic_subtract(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 - Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return ((any)((_hx_tmp_6 - Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
		} else {
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.subtract invalid operands: ", aV.String()), " and ")
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_7, bV.String())))))
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_multiply(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 * Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return ((any)((_hx_tmp_6 * Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))
		} else {
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.multiply invalid operands: ", aV.String()), " and ")
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_7, bV.String())))))
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_divide(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Float64)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 float64
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
		return ((any)(((float64)((_hx_tmp_4 / Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.divide invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_modulo(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 % Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = k
		if _hx_tmp_5 == reflect.Float64 {
			var _hx_tmp_6 float64
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_valueToFloat(aV)
			return math.Mod(_hx_tmp_6, Hx_Field_go_haxe_hxdynamic_valueToFloat(bV))
		} else {
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.modulo invalid operands: ", aV.String()), " and ")
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_7, bV.String())))))
		}
	}
}

func Hx_Field_go_haxe_hxdynamic_bitand(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 & Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.bitand invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_bitor(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 | Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.bitor invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_bitxor(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 ^ Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.bitxor invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_lbitshift(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 << Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.lbitshift invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_rbitshift(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)((_hx_tmp_4 >> Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.rbitshift invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_urbitshift(a any, b any) any {
	var aV reflect.Value
	_ = aV
	aV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(a)
	var bV reflect.Value
	_ = bV
	bV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(b)
	var k reflect.Kind
	_ = k
	k = Hx_Field_go_haxe_hxdynamic_jointKind(aV, bV)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_valueToInt(aV)
		return ((any)(((int32)((((uint32)(_hx_tmp_4)) >> Hx_Field_go_haxe_hxdynamic_valueToInt(bV))))))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.urbitshift invalid operands: ", aV.String()), " and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, bV.String())))))
	}
}

func Hx_Field_go_haxe_hxdynamic_toString(d any) string {
	if d == nil {
		return HxStringNull
	}

	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dV.IsValid() {
		var _hx_tmp_1 reflect.Kind
		_ = _hx_tmp_1
		_hx_tmp_1 = dV.Kind()
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 == reflect.String)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		return dV.String()
	}

	var r string
	_ = r
	r = Hx_Field_std_string(d)
	if r == "null" {
		r = HxStringNull
	}

	return r
}

func Hx_Field_go_haxe_hxdynamic_toBool(d any) bool {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 reflect.Kind
	_ = _hx_tmp_0
	_hx_tmp_0 = dV.Kind()
	if _hx_tmp_0 == reflect.Bool {
		return dV.Bool()
	}

	return (Hx_Field_go_haxe_hxdynamic_valueToInt(dV) != 0)
}

func Hx_Field_go_haxe_hxdynamic_valueToInt(dV reflect.Value) int32 {
	dV = Hx_Field_go_haxe_hxdynamic_unwrapNullable(dV)
	if dV.CanUint() {
		var this1 uint64
		_ = this1
		this1 = dV.Uint()
		return ((int32)(this1))
	} else {
		if dV.CanInt() {
			var this1 int64
			_ = this1
			this1 = dV.Int()
			return ((int32)(this1))
		} else {
			if dV.CanFloat() {
				return Hx_Field_std_int(math.Floor(dV.Float()))
			}
		}
	}

	return 0
}

func Hx_Field_go_haxe_hxdynamic_isClass(v reflect.Value) bool {
	var kind reflect.Kind
	_ = kind
	kind = v.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		return Hx_Field_go_haxe_hxdynamic_isClass(v.Elem())
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 != reflect.Struct {
		return false
	}

	return v.FieldByName("VTable").IsValid()
}

func Hx_Field_go_haxe_hxdynamic_classVTable(v reflect.Value) reflect.Value {
	var kind reflect.Kind
	_ = kind
	kind = v.Kind()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 reflect.Kind
		_ = _hx_tmp_1
		_hx_tmp_1 = kind
		if _hx_tmp_1 == reflect.Ptr {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_2 reflect.Kind
			_ = _hx_tmp_2
			_hx_tmp_2 = kind
			var _hx_tmp_3 reflect.Kind
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
		}

		if !(_hx_tmp_0) {
			break
		}

		if v.IsNil() {
			return Hx_Field_go_haxe_hxdynamic_Null
		}

		v = v.Elem()
		kind = v.Kind()
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 != reflect.Struct {
		return Hx_Field_go_haxe_hxdynamic_Null
	}

	var vtable reflect.Value
	_ = vtable
	vtable = v.FieldByName("VTable")
	for {
		if !(vtable.IsValid()) {
			break
		}

		var vk reflect.Kind
		_ = vk
		vk = vtable.Kind()
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = vk
		if _hx_tmp_5 == reflect.Ptr {
			var _hx_tmp_6 reflect.Value
			_ = _hx_tmp_6
			if vtable.IsNil() {
				_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_Null
			} else {
				_hx_tmp_6 = vtable
			}

			return _hx_tmp_6
		} else {
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = vk
			if _hx_tmp_7 == reflect.Interface {
				if vtable.IsNil() {
					return Hx_Field_go_haxe_hxdynamic_Null
				}

				vtable = vtable.Elem()
			} else {
				break
			}
		}
	}

	return Hx_Field_go_haxe_hxdynamic_Null
}

func Hx_Field_go_haxe_hxdynamic_convertToType(v reflect.Value, t reflect.Type) reflect.Value {
	if !v.IsValid() {
		return reflect.Zero(t)
	}

	var cv reflect.Value
	_ = cv
	cv = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(v.Interface())
	var k reflect.Kind
	_ = k
	k = t.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = k
	if _hx_tmp_1 == reflect.Int32 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = k
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Int)
	}

	if _hx_tmp_0 {
		return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_valueToInt(cv))
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = k
	if _hx_tmp_4 == reflect.Float64 {
		return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_valueToFloat(cv))
	}

	var _hx_tmp_5 reflect.Kind
	_ = _hx_tmp_5
	_hx_tmp_5 = k
	if _hx_tmp_5 == reflect.String {
		return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_toString(cv.Interface()))
	}

	var _hx_tmp_6 reflect.Kind
	_ = _hx_tmp_6
	_hx_tmp_6 = k
	if _hx_tmp_6 == reflect.Bool {
		return reflect.ValueOf(Hx_Field_go_haxe_hxdynamic_toBool(cv))
	}

	var _hx_tmp_7 bool
	_ = _hx_tmp_7
	var _hx_tmp_8 reflect.Kind
	_ = _hx_tmp_8
	_hx_tmp_8 = k
	if _hx_tmp_8 == reflect.Ptr {
		var _hx_tmp_9 reflect.Kind
		_ = _hx_tmp_9
		_hx_tmp_9 = t.Elem().Kind()
		var _hx_tmp_10 reflect.Kind
		_ = _hx_tmp_10
		_hx_tmp_10 = _hx_tmp_9
		_hx_tmp_7 = (_hx_tmp_10 == reflect.Struct)
	} else {
		_hx_tmp_7 = false
	}

	if _hx_tmp_7 {
		var _hx_tmp_11 reflect.Value
		_ = _hx_tmp_11
		_hx_tmp_11 = cv
		var _hx_tmp_12 string
		_ = _hx_tmp_12
		_hx_tmp_12 = t.Elem().Name()
		var cls reflect.Value
		_ = cls
		cls = Hx_Field_go_haxe_hxdynamic_valueToClass(_hx_tmp_11, _hx_tmp_12, nil)
		if cls.IsValid() {
			var _hx_tmp_13 reflect.Kind
			_ = _hx_tmp_13
			_hx_tmp_13 = cls.Kind()
			if _hx_tmp_13 == reflect.Ptr {
				return cls
			}

			if cls.CanAddr() {
				return cls.Addr()
			}
		}
	}

	var _hx_tmp_14 bool
	_ = _hx_tmp_14
	var _hx_tmp_15 reflect.Kind
	_ = _hx_tmp_15
	_hx_tmp_15 = k
	if _hx_tmp_15 == reflect.Func {
		var _hx_tmp_16 reflect.Kind
		_ = _hx_tmp_16
		_hx_tmp_16 = cv.Kind()
		var _hx_tmp_17 reflect.Kind
		_ = _hx_tmp_17
		_hx_tmp_17 = _hx_tmp_16
		_hx_tmp_14 = (_hx_tmp_17 == reflect.Func)
	} else {
		_hx_tmp_14 = false
	}

	if _hx_tmp_14 {
		var srcT reflect.Type
		_ = srcT
		srcT = cv.Type()
		if !srcT.AssignableTo(t) {
			return Hx_Field_go_haxe_hxdynamic_createNewClosure(cv, t)
		}

		return cv
	}

	var coerced any
	_ = coerced
	coerced = HxCoerceArray(cv.Interface(), t)
	if !Hx_Field_go_haxe_hxdynamic_isNull(coerced) {
		return reflect.ValueOf(coerced)
	}

	var _hx_tmp_18 bool
	_ = _hx_tmp_18
	if Hx_Field_go_haxe_hxdynamic_isNullableType(t) {
		_hx_tmp_18 = (cv.Type() != t)
	} else {
		_hx_tmp_18 = false
	}

	if _hx_tmp_18 {
		var out reflect.Value
		_ = out
		out = reflect.New(t).Elem()
		var _hx_tmp_19 reflect.Value
		_ = _hx_tmp_19
		_hx_tmp_19 = cv
		out.Field(int(((int32)(0)))).Set(Hx_Field_go_haxe_hxdynamic_convertToType(_hx_tmp_19, t.Field(int(((int32)(0)))).Type))
		out.Field(int(((int32)(1)))).SetBool(true)
		return out
	}

	return cv
}

func Hx_Field_go_haxe_hxdynamic_createNewClosure(fn reflect.Value, t reflect.Type) reflect.Value {
	var srcT reflect.Type
	_ = srcT
	srcT = fn.Type()
	var _hx_tmp_0 reflect.Type
	_ = _hx_tmp_0
	_hx_tmp_0 = t
	return reflect.MakeFunc(_hx_tmp_0, func(args []reflect.Value) []reflect.Value {
		var numIn int
		_ = numIn
		numIn = srcT.NumIn()
		var callArgs HxArray[reflect.Value]
		_ = callArgs
		callArgs = HxMakeArray[reflect.Value]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = ((int32)(numIn))
			for _g < _g1 {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_1
				var pt reflect.Type
				_ = pt
				pt = srcT.In(int(i))
				var _hx_tmp_2 HxArray[reflect.Value]
				_ = _hx_tmp_2
				_hx_tmp_2 = callArgs
				_ = (Hx_Array_Push[reflect.Value](_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_convertToType(args[i], pt)))
			}
		}

		var self HxArray[reflect.Value]
		_ = self
		self = callArgs
		var results []reflect.Value
		_ = results
		results = fn.Call(self.Underlying())
		var length HxNullable[int32]
		_ = length
		length = HxNullable[int32]{}

		var this1 []reflect.Value
		_ = this1
		var length1 HxNullable[int32]
		_ = length1
		length1 = length
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (length1).Valid != false {
			_hx_tmp_3 = (length1).Value
		} else {
			_hx_tmp_3 = 0
		}

		this1 = make([]reflect.Value, _hx_tmp_3)
		var out []reflect.Value
		_ = out
		out = this1
		var numOut int
		_ = numOut
		numOut = t.NumOut()
		{
			var _g2 int32
			_ = _g2
			_g2 = 0
			var _g3 int32
			_ = _g3
			_g3 = ((int32)(numOut))
			for _g2 < _g3 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g2
				_g2 = (_g2 + 1)
				var i int32
				_ = i
				i = _hx_tmp_4
				var _hx_tmp_5 reflect.Value
				_ = _hx_tmp_5
				_hx_tmp_5 = results[i]
				var v reflect.Value
				_ = v
				v = Hx_Field_go_haxe_hxdynamic_convertToType(_hx_tmp_5, t.Out(int(i)))
				out = append(out, v)
			}
		}

		return out
	})
}

func Hx_Field_go_haxe_hxdynamic_call(fn any, args HxArray[any]) any {
	var fV reflect.Value
	_ = fV
	fV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(fn)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if Hx_Field_go_haxe_hxdynamic_isNull(fn) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = !fV.IsValid()
	}

	if _hx_tmp_0 {
		panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.call null function value"))))
	}

	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = fV.Kind()
	if _hx_tmp_1 != reflect.Func {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.call value not callable: ", Hx_Field_std_string(fn))))))
	}

	var fType reflect.Type
	_ = fType
	fType = fV.Type()
	var numIn int
	_ = numIn
	numIn = fType.NumIn()
	var isVariadic bool
	_ = isVariadic
	isVariadic = fType.IsVariadic()
	var argVals HxArray[reflect.Value]
	_ = argVals
	argVals = HxMakeArray[reflect.Value]()
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 int
	_ = _hx_tmp_3
	_hx_tmp_3 = numIn
	if _hx_tmp_3 >= int(((int32)(1))) {
		var a float64
		_ = a
		a = ((float64)(Hx_Field_go_haxe_hxdynamic_getArrayLength(args)))
		var _hx_tmp_4 float64
		_ = _hx_tmp_4
		_hx_tmp_4 = float64(a)
		var _hx_tmp_5 float64
		_ = _hx_tmp_5
		_hx_tmp_5 = _hx_tmp_4
		_hx_tmp_2 = (_hx_tmp_5 > float64(numIn))
	} else {
		_hx_tmp_2 = false
	}

	var actingVariadic bool
	_ = actingVariadic
	actingVariadic = _hx_tmp_2
	if isVariadic {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ((int32)(numIn))
		for _g < _g1 {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_6
			var lhs int32
			_ = lhs
			lhs = i
			var _hx_tmp_7 int
			_ = _hx_tmp_7
			_hx_tmp_7 = numIn
			var rhs int
			_ = rhs
			rhs = (_hx_tmp_7 - int(((int32)(1))))
			var _hx_tmp_8 float64
			_ = _hx_tmp_8
			_hx_tmp_8 = float64(rhs)
			var _hx_tmp_9 float64
			_ = _hx_tmp_9
			_hx_tmp_9 = _hx_tmp_8
			if _hx_tmp_9 == float64(lhs) {
				var elemType reflect.Type
				_ = elemType
				elemType = fType.In(int(i)).Elem()
				{
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = i
					var tmp__g1_1 int32
					_ = tmp__g1_1
					tmp__g1_1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args)
					for tmp__g_1 < tmp__g1_1 {
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						_hx_tmp_10 = tmp__g_1
						tmp__g_1 = (tmp__g_1 + 1)
						var j int32
						_ = j
						j = _hx_tmp_10
						var _hx_tmp_11 HxArray[reflect.Value]
						_ = _hx_tmp_11
						_hx_tmp_11 = argVals
						var _hx_tmp_12 any
						_ = _hx_tmp_12
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = j
						if _hx_tmp_13 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args) {
							_hx_tmp_12 = args.Get(j)
						} else {
							_hx_tmp_12 = nil
						}

						_ = (Hx_Array_Push[reflect.Value](_hx_tmp_11, Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(_hx_tmp_12), elemType)))
					}
				}

				break
			}

			var pt reflect.Type
			_ = pt
			pt = fType.In(int(i))
			var _hx_tmp_14 HxArray[reflect.Value]
			_ = _hx_tmp_14
			_hx_tmp_14 = argVals
			var _hx_tmp_15 any
			_ = _hx_tmp_15
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = i
			if _hx_tmp_16 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args) {
				_hx_tmp_15 = args.Get(i)
			} else {
				_hx_tmp_15 = nil
			}

			_ = (Hx_Array_Push[reflect.Value](_hx_tmp_14, Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(_hx_tmp_15), pt)))
		}
	} else {
		if actingVariadic {
			{
				var _g int32
				_ = _g
				_g = 0
				var _hx_tmp_17 int
				_ = _hx_tmp_17
				_hx_tmp_17 = numIn
				var _g1 int32
				_ = _g1
				_g1 = ((int32)((_hx_tmp_17 - int(((int32)(1))))))
				for _g < _g1 {
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					_hx_tmp_18 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_18
					var pt reflect.Type
					_ = pt
					pt = fType.In(int(i))
					var _hx_tmp_19 HxArray[reflect.Value]
					_ = _hx_tmp_19
					_hx_tmp_19 = argVals
					var _hx_tmp_20 any
					_ = _hx_tmp_20
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					_hx_tmp_21 = i
					if _hx_tmp_21 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args) {
						_hx_tmp_20 = args.Get(i)
					} else {
						_hx_tmp_20 = nil
					}

					_ = (Hx_Array_Push[reflect.Value](_hx_tmp_19, Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(_hx_tmp_20), pt)))
				}
			}

			var _hx_tmp_22 HxArray[any]
			_ = _hx_tmp_22
			_hx_tmp_22 = args
			var _hx_tmp_24 int
			_ = _hx_tmp_24
			_hx_tmp_24 = numIn
			var _hx_tmp_23 int32
			_ = _hx_tmp_23
			_hx_tmp_23 = ((int32)((_hx_tmp_24 - int(((int32)(1))))))
			var rest HxArray[any]
			_ = rest
			rest = (Hx_Array_Slice[any](_hx_tmp_22, _hx_tmp_23, HxNullable[int32]{}))
			var _hx_tmp_25 int
			_ = _hx_tmp_25
			_hx_tmp_25 = numIn
			var pt reflect.Type
			_ = pt
			pt = fType.In((_hx_tmp_25 - int(((int32)(1)))))
			var _hx_tmp_26 HxArray[reflect.Value]
			_ = _hx_tmp_26
			_hx_tmp_26 = argVals
			_ = (Hx_Array_Push[reflect.Value](_hx_tmp_26, Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(((any)(rest))), pt)))
		} else {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = ((int32)(numIn))
			for _g < _g1 {
				var _hx_tmp_27 int32
				_ = _hx_tmp_27
				_hx_tmp_27 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_27
				var pt reflect.Type
				_ = pt
				pt = fType.In(int(i))
				var _hx_tmp_28 any
				_ = _hx_tmp_28
				var _hx_tmp_29 int32
				_ = _hx_tmp_29
				_hx_tmp_29 = i
				if _hx_tmp_29 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args) {
					_hx_tmp_28 = args.Get(i)
				} else {
					_hx_tmp_28 = nil
				}

				var av any
				_ = av
				av = _hx_tmp_28
				var _hx_tmp_30 HxArray[reflect.Value]
				_ = _hx_tmp_30
				_hx_tmp_30 = argVals
				_ = (Hx_Array_Push[reflect.Value](_hx_tmp_30, Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(av), pt)))
			}
		}
	}

	var self HxArray[reflect.Value]
	_ = self
	self = argVals
	var results []reflect.Value
	_ = results
	results = fV.Call(self.Underlying())
	var _hx_tmp_31 int
	_ = _hx_tmp_31
	_hx_tmp_31 = len(results)
	if _hx_tmp_31 == int(((int32)(0))) {
		return nil
	}

	var res0 any
	_ = res0
	res0 = Hx_Field_go_haxe_hxdynamic_ensureInterface(results[0])
	var isNullable bool
	_ = isNullable
	isNullable = false
	_, isNullable = (res0).(HxNullable[interface{}])
	if isNullable {
		if (res0).(HxNullable[interface{}]).Valid {
			res0 = (res0).(HxNullable[interface{}]).Value
		} else {
			res0 = ((any)(Hx_Field_go_haxe_hxdynamic_Null))
		}
	}

	return res0
}

func Hx_Field_go_haxe_hxdynamic_toInt(d any) int32 {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	return Hx_Field_go_haxe_hxdynamic_valueToInt(dV)
}

func Hx_Field_go_haxe_hxdynamic_valueToInt64(dV reflect.Value) int64 {
	dV = Hx_Field_go_haxe_hxdynamic_unwrapNullable(dV)
	if dV.CanInt() {
		return dV.Int()
	} else {
		if dV.CanUint() {
			return int64(dV.Uint())
		} else {
			if dV.CanFloat() {
				return int64(math.Floor(dV.Float()))
			}
		}
	}

	return int64(((int32)(0)))
}

func Hx_Field_go_haxe_hxdynamic_toInt64(d any) int64 {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	return Hx_Field_go_haxe_hxdynamic_valueToInt64(dV)
}

func Hx_Field_go_haxe_hxdynamic_valueToFloat(dV reflect.Value) float64 {
	dV = Hx_Field_go_haxe_hxdynamic_unwrapNullable(dV)
	if dV.CanUint() {
		var this1 uint64
		_ = this1
		this1 = dV.Uint()
		return ((float64)(this1))
	} else {
		if dV.CanInt() {
			var this1 int64
			_ = this1
			this1 = dV.Int()
			return ((float64)(this1))
		} else {
			if dV.CanFloat() {
				return dV.Float()
			}
		}
	}

	return float64(0.0)
}

func Hx_Field_go_haxe_hxdynamic_toFloat(d any) float64 {
	var dV reflect.Value
	_ = dV
	dV = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	return Hx_Field_go_haxe_hxdynamic_valueToFloat(dV)
}

func Hx_Field_go_haxe_hxdynamic_valueToClass(value reflect.Value, className string, hoppedVTable__isOptional any) reflect.Value {
	hoppedVTable := HxDefault[bool]()
	_ = hoppedVTable

	if hoppedVTable__isOptional != nil {
		hoppedVTable = HxConvert[bool](hoppedVTable__isOptional)
	} else {
		hoppedVTable = false
	}
	{
		if !value.IsValid() {
			return Hx_Field_go_haxe_hxdynamic_Null
		}

		var kind reflect.Kind
		_ = kind
		kind = value.Kind()
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 reflect.Kind
		_ = _hx_tmp_1
		_hx_tmp_1 = kind
		if _hx_tmp_1 == reflect.Ptr {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_2 reflect.Kind
			_ = _hx_tmp_2
			_hx_tmp_2 = kind
			var _hx_tmp_3 reflect.Kind
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
		}

		if _hx_tmp_0 {
			if value.IsNil() {
				return Hx_Field_go_haxe_hxdynamic_Null
			}

			return Hx_Field_go_haxe_hxdynamic_valueToClass(value.Elem(), className, hoppedVTable)
		}

		var _hx_tmp_4 reflect.Kind
		_ = _hx_tmp_4
		_hx_tmp_4 = kind
		if _hx_tmp_4 != reflect.Struct {
			return Hx_Field_go_haxe_hxdynamic_Null
		}

		if value.Type().Name() == className {
			return value
		}

		var f reflect.Value
		_ = f
		f = value.FieldByName(className)
		if f.IsValid() {
			return Hx_Field_go_haxe_hxdynamic_valueToClass(f, className, hoppedVTable)
		}

		var boxed reflect.Value
		_ = boxed
		boxed = value.FieldByName("Value")
		if boxed.IsValid() {
			return Hx_Field_go_haxe_hxdynamic_valueToClass(boxed, className, hoppedVTable)
		}

		if !hoppedVTable {
			var vtable reflect.Value
			_ = vtable
			vtable = value.FieldByName("VTable")
			if vtable.IsValid() {
				return Hx_Field_go_haxe_hxdynamic_valueToClass(vtable, className, true)
			}
		}

		return Hx_Field_go_haxe_hxdynamic_Null
	}
}

func Hx_Field_go_haxe_hxdynamic_toClass(d any, className string) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(d) {
		return nil
	}

	var _hx_tmp_0 reflect.Value
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_ensureValue(d)
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = className
	var cls reflect.Value
	_ = cls
	cls = Hx_Field_go_haxe_hxdynamic_valueToClass(_hx_tmp_0, _hx_tmp_1, nil)
	if !cls.IsValid() {
		return nil
	}

	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = cls.Kind()
	if _hx_tmp_2 == reflect.Ptr {
		return cls.Interface()
	}

	if !cls.CanAddr() {
		var boxed reflect.Value
		_ = boxed
		boxed = reflect.New(cls.Type())
		boxed.Elem().Set(cls)
		return boxed.Interface()
	}

	return cls.Addr().Interface()
}

func Hx_Field_go_haxe_hxdynamic_castClass(d any, className string) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(d) {
		return nil
	}

	var v reflect.Value
	_ = v
	v = Hx_Field_go_haxe_hxdynamic_ensureValue(d)
	var _hx_tmp_0 reflect.Value
	_ = _hx_tmp_0
	_hx_tmp_0 = v
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = className
	var cls reflect.Value
	_ = cls
	cls = Hx_Field_go_haxe_hxdynamic_valueToClass(_hx_tmp_0, _hx_tmp_1, nil)
	if !cls.IsValid() {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Can't cast ", v.Type().String()), " to "), className)))))
	}

	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = cls.Kind()
	if _hx_tmp_2 == reflect.Ptr {
		return cls.Interface()
	}

	if !cls.CanAddr() {
		var boxed reflect.Value
		_ = boxed
		boxed = reflect.New(cls.Type())
		boxed.Elem().Set(cls)
		return boxed.Interface()
	}

	return cls.Addr().Interface()
}

func Hx_Field_go_haxe_hxdynamic_isNullableType(t reflect.Type) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = t.Kind()
	if _hx_tmp_2 != reflect.Struct {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_3 int
		_ = _hx_tmp_3
		_hx_tmp_3 = t.NumField()
		var _hx_tmp_4 int
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_3
		_hx_tmp_1 = (_hx_tmp_4 != int(((int32)(2))))
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = !strings.HasPrefix(t.Name(), "HxNullable")
	}

	if _hx_tmp_0 {
		return false
	}

	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	if t.Field(int(((int32)(0)))).Name == "Value" {
		_hx_tmp_6 = (t.Field(int(((int32)(1)))).Name == "Valid")
	} else {
		_hx_tmp_6 = false
	}

	if _hx_tmp_6 {
		var _hx_tmp_7 reflect.Kind
		_ = _hx_tmp_7
		_hx_tmp_7 = t.Field(int(((int32)(1)))).Type.Kind()
		var _hx_tmp_8 reflect.Kind
		_ = _hx_tmp_8
		_hx_tmp_8 = _hx_tmp_7
		_hx_tmp_5 = (_hx_tmp_8 == reflect.Bool)
	} else {
		_hx_tmp_5 = false
	}

	return _hx_tmp_5
}

func Hx_Field_go_haxe_hxdynamic_unwrapNullable(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return value
	}

	var v reflect.Value
	_ = v
	v = value
	var _hx_tmp_0 reflect.Kind
	_ = _hx_tmp_0
	_hx_tmp_0 = v.Kind()
	if _hx_tmp_0 == reflect.Interface {
		if v.IsNil() {
			return value
		}

		v = v.Elem()
	}

	if !Hx_Field_go_haxe_hxdynamic_isNullableType(v.Type()) {
		return value
	}

	var _hx_tmp_1 reflect.Value
	_ = _hx_tmp_1
	if v.Field(int(((int32)(1)))).Bool() {
		_hx_tmp_1 = v.Field(int(((int32)(0))))
	} else {
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_Null
	}

	return _hx_tmp_1
}

func Hx_Field_go_haxe_hxdynamic_getField(dyn any, fieldName string) any {
	return Hx_Field_go_haxe_hxdynamic_getFieldFrom(dyn, fieldName, true)
}

func Hx_Field_go_haxe_hxdynamic_tryHxClass(dyn any) *Hx_Obj_go_haxe_hxclass {
	var cls *Hx_Obj_go_haxe_hxclass
	_ = cls
	var ok bool
	_ = ok
	ok = false
	cls, ok = dyn.(*Hx_Obj_go_haxe_hxclass)
	var _hx_tmp_0 *Hx_Obj_go_haxe_hxclass
	_ = _hx_tmp_0
	if ok {
		_hx_tmp_0 = cls
	} else {
		_hx_tmp_0 = nil
	}

	return _hx_tmp_0
}

func Hx_Field_go_haxe_hxdynamic_tryHxEnum(dyn any) *Hx_Obj_go_haxe_hxenum {
	var en *Hx_Obj_go_haxe_hxenum
	_ = en
	var ok bool
	_ = ok
	ok = false
	en, ok = dyn.(*Hx_Obj_go_haxe_hxenum)
	var _hx_tmp_0 *Hx_Obj_go_haxe_hxenum
	_ = _hx_tmp_0
	if ok {
		_hx_tmp_0 = en
	} else {
		_hx_tmp_0 = nil
	}

	return _hx_tmp_0
}

func Hx_Field_go_haxe_hxdynamic_hxEnumConstructor(en *Hx_Obj_go_haxe_hxenum, fieldName string) any {
	var _hx_tmp_0 HxArray[string]
	_ = _hx_tmp_0
	_hx_tmp_0 = en.Hx_Field_constructorNames
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = fieldName
	var i int32
	_ = i
	i = (Hx_Array_IndexOf[string](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
	if i < 0 {
		return nil
	}

	if en.Hx_Field_constructorArgCounts.Get(i) == 0 {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = i
		return ((any)(en.Hx_Field_createByIndex(_hx_tmp_2, HxMakeArray[any]().Dyn())))
	}

	return func(args ...any) any { return en.Hx_Field_createByIndex(i, HxMakeArray(args...)) }
}

func Hx_Field_go_haxe_hxdynamic_isHxClass(dyn any) bool {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		return false
	}

	var cls any
	_ = cls
	var ok bool
	_ = ok
	ok = false
	cls, ok = dyn.(*Hx_Obj_go_haxe_hxclass)
	return ok
}

func Hx_Field_go_haxe_hxdynamic_isHxEnum(dyn any) bool {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		return false
	}

	var en any
	_ = en
	var ok bool
	_ = ok
	ok = false
	en, ok = dyn.(*Hx_Obj_go_haxe_hxenum)
	return ok
}

func Hx_Field_go_haxe_hxdynamic_isHxEnumValue(dyn any) bool {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		return false
	}

	var en any
	_ = en
	var ok bool
	_ = ok
	ok = false
	en, ok = dyn.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
	return ok
}

func Hx_Field_go_haxe_hxdynamic_asClass(d any) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(d) {
		return nil
	}

	var c any
	_ = c
	var ok bool
	_ = ok
	ok = false
	c, ok = d.(*Hx_Obj_go_haxe_hxclass)
	if ok {
		return c
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(d, "_RTTI")
	return Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
}

func Hx_Field_go_haxe_hxdynamic_asEnum(d any) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(d) {
		return nil
	}

	var e any
	_ = e
	var ok bool
	_ = ok
	ok = false
	e, ok = d.(*Hx_Obj_go_haxe_hxenum)
	if ok {
		return e
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(d, "enumType")
	return Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
}

func Hx_Field_go_haxe_hxdynamic_hxClassStaticValue(cls *Hx_Obj_go_haxe_hxclass, fieldName string) reflect.Value {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(cls.Hx_Field_getStaticFieldPtr)), ((any)(nil))))) {
		return Hx_Field_go_haxe_hxdynamic_Null
	}

	var ptr any
	_ = ptr
	ptr = cls.Hx_Field_getStaticFieldPtr(fieldName)
	if Hx_Field_go_haxe_hxdynamic_isNull(ptr) {
		return Hx_Field_go_haxe_hxdynamic_Null
	}

	var pv reflect.Value
	_ = pv
	pv = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(ptr)
	var _hx_tmp_0 reflect.Kind
	_ = _hx_tmp_0
	_hx_tmp_0 = pv.Kind()
	if _hx_tmp_0 == reflect.Ptr {
		return pv.Elem()
	}

	return pv
}

func Hx_Field_go_haxe_hxdynamic_getFieldFrom(dyn any, fieldName string, hop bool) any {

	switch m := dyn.(type) {
	case map[string]any:
		return m[fieldName]
	default:
		break
	}

	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.field null field access: ", fieldName)))))
	}

	var hxCls *Hx_Obj_go_haxe_hxclass
	_ = hxCls
	hxCls = Hx_Field_go_haxe_hxdynamic_tryHxClass(dyn)
	if hxCls != nil {
		var sv reflect.Value
		_ = sv
		sv = Hx_Field_go_haxe_hxdynamic_hxClassStaticValue(hxCls, fieldName)
		if sv.IsValid() {
			sv = Hx_Field_go_haxe_hxdynamic_unwrapNullable(sv)
			if !sv.IsValid() {
				return nil
			}

			var _hx_tmp_0 any
			_ = _hx_tmp_0
			if sv.CanInterface() {
				_hx_tmp_0 = sv.Interface()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}

	var hxEn *Hx_Obj_go_haxe_hxenum
	_ = hxEn
	hxEn = Hx_Field_go_haxe_hxdynamic_tryHxEnum(dyn)
	if hxEn != nil {
		var ctor any
		_ = ctor
		ctor = Hx_Field_go_haxe_hxdynamic_hxEnumConstructor(hxEn, fieldName)
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(ctor, nil))) {
			return ctor
		}
	}

	var arr HxNullable[HxArrayDyn]
	_ = arr
	arr = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(dyn)
	if (arr).Valid != false {
		switch fieldName {
		case "concat":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(other HxArray[any]) HxArray[any] {
					return Hx_Array_Concat(tmp_arr_1, other)
				}))
			}

		case "contains":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any) bool {
					return Hx_Array_Contains(tmp_arr_1, val)
				}))
			}

		case "copy":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() HxArray[any] {
					return Hx_Array_Copy(tmp_arr_1)
				}))
			}

		case "filter":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(f func(_hx_carg0 any) bool) HxArray[any] {
					return Hx_Array_Filter(tmp_arr_1, f)
				}))
			}

		case "indexOf":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any, fromIndex__isOptional any) int32 {
					fromIndex := HxDefault[HxNullable[int32]]()
					_ = fromIndex

					if fromIndex__isOptional != nil {
						fromIndex = HxConvert[HxNullable[int32]](fromIndex__isOptional)
					} else {

					}
					{
						if (fromIndex).Valid == false {
							var hx_nullable_t14 HxNullable[int32]
							_ = hx_nullable_t14
							hx_nullable_t14 = HxNullable[int32]{}

							fromIndex = HxNullable[int32]{Value: (hx_nullable_t14).Value, Valid: (hx_nullable_t14).Valid}
						}

						return Hx_Array_IndexOf(tmp_arr_1, val, fromIndex)
					}
				}))
			}

		case "insert":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(pos int32, val any) {
					Hx_Array_Insert(tmp_arr_1, pos, val)
				}))
			}

		case "iterator":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() *Hx_Obj_haxe_iterators_arrayiterator {
					return Hx_Array_Iterator(tmp_arr_1)
				}))
			}

		case "join":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(sep string) string {
					return Hx_Array_Join(tmp_arr_1, sep)
				}))
			}

		case "keyValueIterator":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
					return Hx_Array_KeyValueIterator(tmp_arr_1)
				}))
			}

		case "lastIndexOf":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any, fromIndex__isOptional any) int32 {
					fromIndex := HxDefault[HxNullable[int32]]()
					_ = fromIndex

					if fromIndex__isOptional != nil {
						fromIndex = HxConvert[HxNullable[int32]](fromIndex__isOptional)
					} else {

					}
					{
						if (fromIndex).Valid == false {
							var hx_nullable_t15 HxNullable[int32]
							_ = hx_nullable_t15
							hx_nullable_t15 = HxNullable[int32]{}

							fromIndex = HxNullable[int32]{Value: (hx_nullable_t15).Value, Valid: (hx_nullable_t15).Valid}
						}

						return Hx_Array_LastIndexOf(tmp_arr_1, val, fromIndex)
					}
				}))
			}

		case "length":
			{
				return ((any)(arr.Value.Len()))
			}

		case "map":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(f func(_hx_carg0 any) any) HxArray[any] {
					return Hx_Array_Map(tmp_arr_1, f)
				}))
			}

		case "pop":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() any {
					return Hx_Array_Pop(tmp_arr_1)
				}))
			}

		case "push":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any) int32 {
					return Hx_Array_Push(tmp_arr_1, val)
				}))
			}

		case "remove":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any) bool {
					return Hx_Array_Remove(tmp_arr_1, val)
				}))
			}

		case "resize":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(size int32) {
					Hx_Array_Resize(tmp_arr_1, size)
				}))
			}

		case "reverse":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() {
					Hx_Array_Reverse(tmp_arr_1)
				}))
			}

		case "shift":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() any {
					return Hx_Array_Shift(tmp_arr_1)
				}))
			}

		case "slice":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(pos int32, end__isOptional any) HxArray[any] {
					end := HxDefault[HxNullable[int32]]()
					_ = end

					if end__isOptional != nil {
						end = HxConvert[HxNullable[int32]](end__isOptional)
					} else {

					}
					{
						if (end).Valid == false {
							var hx_nullable_t16 HxNullable[int32]
							_ = hx_nullable_t16
							hx_nullable_t16 = HxNullable[int32]{}

							end = HxNullable[int32]{Value: (hx_nullable_t16).Value, Valid: (hx_nullable_t16).Valid}
						}

						return Hx_Array_Slice(tmp_arr_1, pos, end)
					}
				}))
			}

		case "sort":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(f func(_hx_carg0 any, _hx_carg1 any) int32) {
					Hx_Array_Sort(tmp_arr_1, f)
				}))
			}

		case "splice":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(pos int32, _hx_reserved_len int32) HxArray[any] {
					return Hx_Array_Splice(tmp_arr_1, pos, _hx_reserved_len)
				}))
			}

		case "toString":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func() string {
					return Hx_Array_ToString(tmp_arr_1)
				}))
			}

		case "unshift":
			{
				var tmp_arr_1 HxArray[any]
				_ = tmp_arr_1
				tmp_arr_1 = HxAnyToArray(dyn)
				return ((any)(func(val any) any {
					return Hx_Array_Unshift(tmp_arr_1, val)
				}))
			}

		default:
			{
				return nil
			}
		}
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var found bool
	_ = found
	found = false
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = !value.IsValid()
	}

	if _hx_tmp_1 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.field null field access: ", fieldName)))))
	}

	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 reflect.Kind
	_ = _hx_tmp_3
	_hx_tmp_3 = kind
	if _hx_tmp_3 == reflect.Ptr {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_4 reflect.Kind
		_ = _hx_tmp_4
		_hx_tmp_4 = kind
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = _hx_tmp_4
		_hx_tmp_2 = (_hx_tmp_5 == reflect.Interface)
	}

	if _hx_tmp_2 {
		return Hx_Field_go_haxe_hxdynamic_getFieldFrom(value.Elem(), fieldName, hop)
	}

	var _hx_tmp_6 reflect.Kind
	_ = _hx_tmp_6
	_hx_tmp_6 = kind
	if _hx_tmp_6 == reflect.Struct {
		var f reflect.Value
		_ = f
		f = value.FieldByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName))
		if f.IsValid() {
			found = true
		} else {
			var vtable reflect.Value
			_ = vtable
			vtable = value.FieldByName("VTable")
			if vtable.IsValid() {
				f = vtable.MethodByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName))
				if f.IsValid() {
					found = true
				} else {
					var _hx_tmp_7 bool
					_ = _hx_tmp_7
					var _hx_tmp_8 bool
					_ = _hx_tmp_8
					if hop {
						var _hx_tmp_9 reflect.Kind
						_ = _hx_tmp_9
						_hx_tmp_9 = vtable.Kind()
						var _hx_tmp_10 reflect.Kind
						_ = _hx_tmp_10
						_hx_tmp_10 = _hx_tmp_9
						_hx_tmp_8 = (_hx_tmp_10 == reflect.Interface)
					} else {
						_hx_tmp_8 = false
					}

					if _hx_tmp_8 {
						_hx_tmp_7 = !vtable.IsNil()
					} else {
						_hx_tmp_7 = false
					}

					if _hx_tmp_7 {
						var self reflect.Value
						_ = self
						self = vtable.Elem()
						var _hx_tmp_11 bool
						_ = _hx_tmp_11
						if self.IsValid() {
							_hx_tmp_11 = self.CanInterface()
						} else {
							_hx_tmp_11 = false
						}

						if _hx_tmp_11 {
							return Hx_Field_go_haxe_hxdynamic_getFieldFrom(self.Interface(), fieldName, false)
						}
					}
				}
			} else {
				var m reflect.Value
				_ = m
				m = value.MethodByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName))
				if m.IsValid() {
					f = m
					found = true
				}
			}
		}

		value = f
	}

	var _hx_tmp_12 reflect.Kind
	_ = _hx_tmp_12
	_hx_tmp_12 = kind
	if _hx_tmp_12 == reflect.Map {
		var mi reflect.Value
		_ = mi
		mi = value.MapIndex(Hx_Field_go_haxe_hxdynamic_ensureValue(fieldName))
		if mi.IsValid() {
			found = true
			value = mi
		}
	}

	var _hx_tmp_13 bool
	_ = _hx_tmp_13
	var _hx_tmp_14 reflect.Kind
	_ = _hx_tmp_14
	_hx_tmp_14 = kind
	if _hx_tmp_14 == reflect.Array {
		_hx_tmp_13 = true
	} else {
		var _hx_tmp_15 reflect.Kind
		_ = _hx_tmp_15
		_hx_tmp_15 = kind
		var _hx_tmp_16 reflect.Kind
		_ = _hx_tmp_16
		_hx_tmp_16 = _hx_tmp_15
		_hx_tmp_13 = (_hx_tmp_16 == reflect.Slice)
	}

	if _hx_tmp_13 {
		if fieldName == "length" {
			value = Hx_Field_go_haxe_hxdynamic_ensureValue(Hx_Field_go_haxe_hxdynamic_getArrayLength(dyn))
			found = true
		}

		if fieldName == "iterator" {
			value = Hx_Field_go_haxe_hxdynamic_ensureValue(func() *Hx_Obj_haxe_iterators_arrayiterator {
				var this1 []any
				_ = this1
				this1 = Hx_Field_go_haxe_hxdynamic_toAnySlice(dyn)
				var self []any
				_ = self
				self = this1
				var _this HxArray[any]
				_ = _this
				_this = HxMakeArray(self...)
				return Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(_this)
			})
			found = true
		}
	}

	var _hx_tmp_17 reflect.Kind
	_ = _hx_tmp_17
	_hx_tmp_17 = kind
	if _hx_tmp_17 == reflect.String {
		var s string
		_ = s
		s = Hx_Field_go_haxe_hxdynamic_toString(dyn)
		switch fieldName {
		case "charAt":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(index int32) string {
					return Hx_Field_go_haxe_hxstring_charAt(s1, index)
				}))
			}

		case "charCodeAt":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(index int32) HxNullable[int32] {
					return Hx_Field_go_haxe_hxstring_charCodeAt(s1, index)
				}))
			}

		case "indexOf":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(str string, startIndex__isOptional any) int32 {
					startIndex := HxDefault[HxNullable[int32]]()
					_ = startIndex

					if startIndex__isOptional != nil {
						startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
					} else {

					}
					{
						if (startIndex).Valid == false {
							var hx_nullable_t17 HxNullable[int32]
							_ = hx_nullable_t17
							hx_nullable_t17 = HxNullable[int32]{}

							startIndex = HxNullable[int32]{Value: (hx_nullable_t17).Value, Valid: (hx_nullable_t17).Valid}
						}

						return Hx_Field_go_haxe_hxstring_indexOf(s1, str, startIndex)
					}
				}))
			}

		case "lastIndexOf":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(str string, startIndex__isOptional any) int32 {
					startIndex := HxDefault[HxNullable[int32]]()
					_ = startIndex

					if startIndex__isOptional != nil {
						startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
					} else {

					}
					{
						if (startIndex).Valid == false {
							var hx_nullable_t18 HxNullable[int32]
							_ = hx_nullable_t18
							hx_nullable_t18 = HxNullable[int32]{}

							startIndex = HxNullable[int32]{Value: (hx_nullable_t18).Value, Valid: (hx_nullable_t18).Valid}
						}

						return Hx_Field_go_haxe_hxstring_lastIndexOf(s1, str, startIndex)
					}
				}))
			}

		case "length":
			{
				return ((any)(int32(utf8.RuneCountInString(s))))
			}

		case "split":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(delimiter string) HxArray[string] {
					return Hx_Field_go_haxe_hxstring_split(s1, delimiter)
				}))
			}

		case "substr":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(pos int32, _hx_reserved_len__isOptional any) string {
					_hx_reserved_len := HxDefault[HxNullable[int32]]()
					_ = _hx_reserved_len

					if _hx_reserved_len__isOptional != nil {
						_hx_reserved_len = HxConvert[HxNullable[int32]](_hx_reserved_len__isOptional)
					} else {

					}
					{
						if (_hx_reserved_len).Valid == false {
							var hx_nullable_t19 HxNullable[int32]
							_ = hx_nullable_t19
							hx_nullable_t19 = HxNullable[int32]{}

							_hx_reserved_len = HxNullable[int32]{Value: (hx_nullable_t19).Value, Valid: (hx_nullable_t19).Valid}
						}

						return Hx_Field_go_haxe_hxstring_substr(s1, pos, _hx_reserved_len)
					}
				}))
			}

		case "substring":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func(startIndex int32, endIndex__isOptional any) string {
					endIndex := HxDefault[HxNullable[int32]]()
					_ = endIndex

					if endIndex__isOptional != nil {
						endIndex = HxConvert[HxNullable[int32]](endIndex__isOptional)
					} else {

					}
					{
						if (endIndex).Valid == false {
							var hx_nullable_t20 HxNullable[int32]
							_ = hx_nullable_t20
							hx_nullable_t20 = HxNullable[int32]{}

							endIndex = HxNullable[int32]{Value: (hx_nullable_t20).Value, Valid: (hx_nullable_t20).Valid}
						}

						return Hx_Field_go_haxe_hxstring_substring(s1, startIndex, endIndex)
					}
				}))
			}

		case "toLowerCase":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func() string {
					return Hx_Field_go_haxe_hxstring_toLowerCase(s1)
				}))
			}

		case "toString":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func() string {
					return Hx_Field_go_haxe_hxstring_toString(s1)
				}))
			}

		case "toUpperCase":
			{
				var s1 string
				_ = s1
				s1 = s
				return ((any)(func() string {
					return Hx_Field_go_haxe_hxstring_toUpperCase(s1)
				}))
			}

		default:
			{
			}
		}
	}

	if !found {
		return nil
	}

	value = Hx_Field_go_haxe_hxdynamic_unwrapNullable(value)
	if !value.IsValid() {
		return nil
	}

	var _hx_tmp_18 any
	_ = _hx_tmp_18
	if value.CanInterface() {
		_hx_tmp_18 = value.Interface()
	} else {
		_hx_tmp_18 = nil
	}

	return _hx_tmp_18
}

func Hx_Field_go_haxe_hxdynamic_setField(dyn any, fieldName string, v any) any {
	var hxCls *Hx_Obj_go_haxe_hxclass
	_ = hxCls
	hxCls = Hx_Field_go_haxe_hxdynamic_tryHxClass(dyn)
	if hxCls != nil {
		var sv reflect.Value
		_ = sv
		sv = Hx_Field_go_haxe_hxdynamic_hxClassStaticValue(hxCls, fieldName)
		if sv.IsValid() {
			if !sv.CanSet() {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.setField cannot set static \"", fieldName), "\" on \""), hxCls.Hx_Field_name), "\"")))))
			}

			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = v
			sv.Set(Hx_Field_go_haxe_hxdynamic_valueToAssign(_hx_tmp_0, sv.Type()))
			return v
		}
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Interface {
		return Hx_Field_go_haxe_hxdynamic_setField(value.Elem(), fieldName, v)
	}

	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = kind
	if _hx_tmp_2 == reflect.Ptr {
		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_3 reflect.Kind
	_ = _hx_tmp_3
	_hx_tmp_3 = kind
	if _hx_tmp_3 == reflect.Struct {
		var field reflect.Value
		_ = field
		field = value.FieldByName(Hx_Field_go_haxe_hxdynamic_formatField(fieldName))
		if !field.IsValid() {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.setField field \"", fieldName), "\" not present on \"")
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_4, Hx_Field_std_string(value)), "\"")))))
		}

		if !field.CanSet() {
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("runtime.HxDynamic.setField cannot set \"", fieldName), "\" on \"")
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, Hx_Field_std_string(value)), "\"")))))
		}

		var _hx_tmp_6 any
		_ = _hx_tmp_6
		_hx_tmp_6 = v
		field.Set(Hx_Field_go_haxe_hxdynamic_valueToAssign(_hx_tmp_6, field.Type()))
	}

	var _hx_tmp_7 reflect.Kind
	_ = _hx_tmp_7
	_hx_tmp_7 = kind
	if _hx_tmp_7 == reflect.Map {
		var fn reflect.Value
		_ = fn
		fn = Hx_Field_go_haxe_hxdynamic_ensureValue(fieldName)
		var _hx_tmp_8 reflect.Value
		_ = _hx_tmp_8
		_hx_tmp_8 = fn
		var _hx_tmp_9 any
		_ = _hx_tmp_9
		_hx_tmp_9 = v
		value.SetMapIndex(_hx_tmp_8, Hx_Field_go_haxe_hxdynamic_valueToAssign(_hx_tmp_9, value.Type().Elem()))
	}

	return v
}

func Hx_Field_go_haxe_hxdynamic_tryDynamicArray(dyn any) HxNullable[HxArrayDyn] {
	var arr HxArrayDyn
	_ = arr
	var ok bool
	_ = ok
	ok = false
	arr, ok = dyn.(HxArrayDyn)
	var _hx_tmp_0 HxNullable[HxArrayDyn]
	_ = _hx_tmp_0
	if ok {
		_hx_tmp_0 = HxNullable[HxArrayDyn]{Value: arr, Valid: true}
	} else {
		_hx_tmp_0 = HxNullable[HxArrayDyn]{}
	}

	return _hx_tmp_0
}

func Hx_Field_go_haxe_hxdynamic_setArrayIndex(dyn any, index any, v any) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.setArrayIndex null array access"))))
	}

	var arr HxNullable[HxArrayDyn]
	_ = arr
	arr = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(dyn)
	if (arr).Valid != false {
		{
			var hx_try_return any
			_ = hx_try_return
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if _hx_reserved__, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_exception").(*Hx_Obj_haxe_exception); _hx_ok {
							_ = _hx_reserved__
							{
								var _hx_tmp_1 any
								_ = _hx_tmp_1
								_hx_tmp_1 = v
								var assigned reflect.Value
								_ = assigned
								assigned = Hx_Field_go_haxe_hxdynamic_valueToAssign(_hx_tmp_1, arr.Value.ElemType())
								var _hx_tmp_2 int32
								_ = _hx_tmp_2
								_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_toInt(index)
								var _hx_tmp_3 any
								_ = _hx_tmp_3
								if assigned.CanInterface() {
									_hx_tmp_3 = assigned.Interface()
								} else {
									_hx_tmp_3 = nil
								}

								arr.Value.Set_Dyn(_hx_tmp_2, _hx_tmp_3)
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_toInt(index)
					arr.Value.Set_Dyn(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_ensureInterface(v))
				}
			}()
			if hx_try_state == 1 {
				return hx_try_return
			}
		}

		return v
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 == reflect.Interface {
		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_5 reflect.Kind
	_ = _hx_tmp_5
	_hx_tmp_5 = kind
	if _hx_tmp_5 == reflect.Ptr {
		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	var _hx_tmp_7 reflect.Kind
	_ = _hx_tmp_7
	_hx_tmp_7 = kind
	if _hx_tmp_7 == reflect.Slice {
		_hx_tmp_6 = true
	} else {
		var _hx_tmp_8 reflect.Kind
		_ = _hx_tmp_8
		_hx_tmp_8 = kind
		var _hx_tmp_9 reflect.Kind
		_ = _hx_tmp_9
		_hx_tmp_9 = _hx_tmp_8
		_hx_tmp_6 = (_hx_tmp_9 == reflect.Array)
	}

	if _hx_tmp_6 {
		var length int
		_ = length
		length = value.Len()
		var _hx_tmp_10 float64
		_ = _hx_tmp_10
		_hx_tmp_10 = float64(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(index))))
		if _hx_tmp_10 >= float64(length) {
			var _hx_tmp_11 reflect.Kind
			_ = _hx_tmp_11
			_hx_tmp_11 = kind
			if _hx_tmp_11 == reflect.Array {
				panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.setArrayIndex out of bounds exception, cannot grow go array"))))
			}

			var _hx_tmp_12 int
			_ = _hx_tmp_12
			_hx_tmp_12 = (int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(index)))) - length)
			value.Grow((_hx_tmp_12 + int(((int32)(1)))))
			value.SetLen(((int)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_add(index, ((int32)(1)))))))
		}

		var item reflect.Value
		_ = item
		item = value.Index(((int)(Hx_Field_go_haxe_hxdynamic_toInt(index))))
		var _hx_tmp_13 any
		_ = _hx_tmp_13
		_hx_tmp_13 = v
		item.Set(Hx_Field_go_haxe_hxdynamic_valueToAssign(_hx_tmp_13, item.Type()))
	}

	return v
}

func Hx_Field_go_haxe_hxdynamic_getArrayIndex(dyn any, index any) any {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.getArrayIndex null array access"))))
	}

	var arr HxNullable[HxArrayDyn]
	_ = arr
	arr = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(dyn)
	if (arr).Valid != false {
		return arr.Value.Get_Dyn(Hx_Field_go_haxe_hxdynamic_toInt(index))
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		return Hx_Field_go_haxe_hxdynamic_getArrayIndex(value.Elem(), index)
	}

	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	var _hx_tmp_5 reflect.Kind
	_ = _hx_tmp_5
	_hx_tmp_5 = kind
	if _hx_tmp_5 == reflect.Slice {
		_hx_tmp_4 = true
	} else {
		var _hx_tmp_6 reflect.Kind
		_ = _hx_tmp_6
		_hx_tmp_6 = kind
		var _hx_tmp_7 reflect.Kind
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_6
		_hx_tmp_4 = (_hx_tmp_7 == reflect.Array)
	}

	if _hx_tmp_4 {
		var length int
		_ = length
		length = value.Len()
		var _hx_tmp_8 float64
		_ = _hx_tmp_8
		_hx_tmp_8 = float64(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(index))))
		if _hx_tmp_8 >= float64(length) {
			return nil
		}

		return value.Index(((int)(Hx_Field_go_haxe_hxdynamic_toInt(index)))).Interface()
	}

	return value.Interface()
}

func Hx_Field_go_haxe_hxdynamic_getArrayLength(dyn any) int32 {
	if Hx_Field_go_haxe_hxdynamic_isNull(dyn) {
		return 0
	}

	var arr HxNullable[HxArrayDyn]
	_ = arr
	arr = Hx_Field_go_haxe_hxdynamic_tryDynamicArray(dyn)
	if (arr).Valid != false {
		return arr.Value.Len()
	}

	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		return Hx_Field_go_haxe_hxdynamic_getArrayLength(value.Elem())
	}

	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	var _hx_tmp_5 reflect.Kind
	_ = _hx_tmp_5
	_hx_tmp_5 = kind
	if _hx_tmp_5 == reflect.Slice {
		_hx_tmp_4 = true
	} else {
		var _hx_tmp_6 reflect.Kind
		_ = _hx_tmp_6
		_hx_tmp_6 = kind
		var _hx_tmp_7 reflect.Kind
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_6
		_hx_tmp_4 = (_hx_tmp_7 == reflect.Array)
	}

	if _hx_tmp_4 {
		var this1 int
		_ = this1
		this1 = value.Len()
		return ((int32)(this1))
	}

	panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.getArrayLength invalid type"))))
}

func Hx_Field_go_haxe_hxdynamic_ensureValue(dyn any) reflect.Value {
	var ok bool
	_ = ok
	ok = false
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_Null
	value, ok = dyn.(reflect.Value)
	if !ok {
		return reflect.ValueOf(dyn)
	}

	return value
}

func Hx_Field_go_haxe_hxdynamic_valueToAssign(v any, t reflect.Type) reflect.Value {
	if Hx_Field_go_haxe_hxdynamic_isNull(v) {
		return reflect.Zero(t)
	}

	if !Hx_Field_go_haxe_hxdynamic_isNullableType(t) {
		return Hx_Field_go_haxe_hxdynamic_convertToType(Hx_Field_go_haxe_hxdynamic_ensureValue(v), t)
	}

	var unwrappedValue reflect.Value
	_ = unwrappedValue
	unwrappedValue = Hx_Field_go_haxe_hxdynamic_unwrapNullable(Hx_Field_go_haxe_hxdynamic_ensureValue(v))
	if !unwrappedValue.IsValid() {
		return reflect.Zero(t)
	}

	var box reflect.Value
	_ = box
	box = reflect.New(t).Elem()
	var _hx_tmp_0 reflect.Value
	_ = _hx_tmp_0
	_hx_tmp_0 = unwrappedValue
	box.Field(int(((int32)(0)))).Set(Hx_Field_go_haxe_hxdynamic_convertToType(_hx_tmp_0, t.Field(int(((int32)(0)))).Type))
	box.Field(int(((int32)(1)))).SetBool(true)
	return box
}

func Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(dyn any) reflect.Value {
	var v reflect.Value
	_ = v
	v = Hx_Field_go_haxe_hxdynamic_ensureValue(dyn)
	var k reflect.Kind
	_ = k
	k = v.Kind()
	for {
		var _hx_tmp_0 reflect.Kind
		_ = _hx_tmp_0
		_hx_tmp_0 = k
		if !(_hx_tmp_0 == reflect.Interface) {
			break
		}

		v = v.Elem()
		k = v.Kind()
	}

	return v
}

func Hx_Field_go_haxe_hxdynamic_ensureInterface(dyn any) any {
	var ok bool
	_ = ok
	ok = false
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_Null
	value, ok = dyn.(reflect.Value)
	if !ok {
		return dyn
	}

	if !value.CanInterface() {
		panic(Hx_Field_haxe_exception_thrown(((any)("runtime.HxDynamic.ensureInterface cannot convert to iface"))))
	}

	return value.Interface()
}

func Hx_Field_go_haxe_hxdynamic_funcReceiver(fn any) uintptr {
	var header unsafe.Pointer
	_ = header
	header = unsafe.Pointer((&fn))
	var wordSize uintptr
	_ = wordSize
	wordSize = unsafe.Sizeof(header)
	var dataLoc *uintptr
	_ = dataLoc
	dataLoc = ((*uintptr)(unsafe.Add(header, wordSize)))
	var funcval unsafe.Pointer
	_ = funcval
	funcval = unsafe.Pointer((*dataLoc))
	var _hx_tmp_0 unsafe.Pointer
	_ = _hx_tmp_0
	_hx_tmp_0 = funcval
	var _hx_tmp_1 uint
	_ = _hx_tmp_1
	_hx_tmp_1 = ((uint)(wordSize))
	var recvLoc *uintptr
	_ = recvLoc
	recvLoc = ((*uintptr)(unsafe.Add(_hx_tmp_0, ((uintptr)((_hx_tmp_1 * uint(((int32)(2)))))))))
	return (*recvLoc)
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[reflect.Value]()
	hxRegisterArrayConstructor[string]()
}
