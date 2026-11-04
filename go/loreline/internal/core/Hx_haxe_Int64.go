package core

type Hx_Typedef_haxe__int64___int64 = int64

var Hx_Abstract_haxe_int64_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Int64",
	HxMakeArray[string]("MASK", "make", "_hx_new", "val", "get_val", "set_val", "high", "get_high", "low", "get_low", "copy", "ofInt", "fromInt", "toFloat", "is", "isInt64", "toInt", "getHigh", "getLow", "isNeg", "isZero", "compare", "ucompare", "toStr", "divMod", "toString", "parseString", "fromFloat", "neg", "preIncrement", "postIncrement", "preDecrement", "postDecrement", "add", "addInt", "sub", "subInt", "intSub", "mul", "mulInt", "div", "divInt", "intDiv", "mod", "modInt", "intMod", "eq", "eqInt", "neq", "neqInt", "lt", "ltInt", "intLt", "lte", "lteInt", "intLte", "gt", "gtInt", "intGt", "gte", "gteInt", "intGte", "complement", "and", "or", "xor", "shl", "shr", "ushr"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return nil
	},
	func() any {
		return nil
	},
	nil,
	func(name string) any { return nil },
)

func Hx_ClassInit_Hx_Obj_VTable_haxe__int64_int64_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__int64_int64_impl_() {
	Hx_Field_haxe__int64_int64_impl__MASK = Hx_Init_Hx_Field_haxe__int64_int64_impl__MASK()
}

type Hx_Obj_VTable_haxe__int64_int64_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__int64_int64_impl_ struct {
	VTable Hx_Obj_VTable_haxe__int64_int64_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__int64_int64_impl__CreateEmptyInstance() *Hx_Obj_haxe__int64_int64_impl_ {
	obj := &Hx_Obj_haxe__int64_int64_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__int64_int64_impl__CreateInstance() *Hx_Obj_haxe__int64_int64_impl_ {
	obj := Hx_Obj_haxe__int64_int64_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__int64_int64_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._Int64.Int64_Impl_",
	HxMakeArray[string]("MASK", "make", "_hx_new", "get_val", "set_val", "get_high", "get_low", "copy", "ofInt", "fromInt", "toFloat", "is", "isInt64", "toInt", "getHigh", "getLow", "isNeg", "isZero", "compare", "ucompare", "toStr", "divMod", "toString", "parseString", "fromFloat", "neg", "preIncrement", "postIncrement", "preDecrement", "postDecrement", "add", "addInt", "sub", "subInt", "intSub", "mul", "mulInt", "div", "divInt", "intDiv", "mod", "modInt", "intMod", "eq", "eqInt", "neq", "neqInt", "lt", "ltInt", "intLt", "lte", "lteInt", "intLte", "gt", "gtInt", "intGt", "gte", "gteInt", "intGte", "complement", "and", "or", "xor", "shl", "shr", "ushr"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__int64_int64_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__int64_int64_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__int64_int64_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "MASK":
			return &Hx_Field_haxe__int64_int64_impl__MASK
		case "make":
			return Hx_Field_haxe__int64_int64_impl__make
		case "_hx_new":
			return Hx_Field_haxe__int64_int64_impl___hx_new
		case "get_val":
			return Hx_Field_haxe__int64_int64_impl__get_val
		case "set_val":
			return Hx_Field_haxe__int64_int64_impl__set_val
		case "get_high":
			return Hx_Field_haxe__int64_int64_impl__get_high
		case "get_low":
			return Hx_Field_haxe__int64_int64_impl__get_low
		case "copy":
			return Hx_Field_haxe__int64_int64_impl__copy
		case "ofInt":
			return Hx_Field_haxe__int64_int64_impl__ofInt
		case "fromInt":
			return Hx_Field_haxe__int64_int64_impl__fromInt
		case "toFloat":
			return Hx_Field_haxe__int64_int64_impl__toFloat
		case "is":
			return Hx_Field_haxe__int64_int64_impl__is
		case "isInt64":
			return Hx_Field_haxe__int64_int64_impl__isInt64
		case "toInt":
			return Hx_Field_haxe__int64_int64_impl__toInt
		case "getHigh":
			return Hx_Field_haxe__int64_int64_impl__getHigh
		case "getLow":
			return Hx_Field_haxe__int64_int64_impl__getLow
		case "isNeg":
			return Hx_Field_haxe__int64_int64_impl__isNeg
		case "isZero":
			return Hx_Field_haxe__int64_int64_impl__isZero
		case "compare":
			return Hx_Field_haxe__int64_int64_impl__compare
		case "ucompare":
			return Hx_Field_haxe__int64_int64_impl__ucompare
		case "toStr":
			return Hx_Field_haxe__int64_int64_impl__toStr
		case "divMod":
			return Hx_Field_haxe__int64_int64_impl__divMod
		case "toString":
			return Hx_Field_haxe__int64_int64_impl__toString
		case "parseString":
			return Hx_Field_haxe__int64_int64_impl__parseString
		case "fromFloat":
			return Hx_Field_haxe__int64_int64_impl__fromFloat
		case "neg":
			return Hx_Field_haxe__int64_int64_impl__neg
		case "preIncrement":
			return Hx_Field_haxe__int64_int64_impl__preIncrement
		case "postIncrement":
			return Hx_Field_haxe__int64_int64_impl__postIncrement
		case "preDecrement":
			return Hx_Field_haxe__int64_int64_impl__preDecrement
		case "postDecrement":
			return Hx_Field_haxe__int64_int64_impl__postDecrement
		case "add":
			return Hx_Field_haxe__int64_int64_impl__add
		case "addInt":
			return Hx_Field_haxe__int64_int64_impl__addInt
		case "sub":
			return Hx_Field_haxe__int64_int64_impl__sub
		case "subInt":
			return Hx_Field_haxe__int64_int64_impl__subInt
		case "intSub":
			return Hx_Field_haxe__int64_int64_impl__intSub
		case "mul":
			return Hx_Field_haxe__int64_int64_impl__mul
		case "mulInt":
			return Hx_Field_haxe__int64_int64_impl__mulInt
		case "div":
			return Hx_Field_haxe__int64_int64_impl__div
		case "divInt":
			return Hx_Field_haxe__int64_int64_impl__divInt
		case "intDiv":
			return Hx_Field_haxe__int64_int64_impl__intDiv
		case "mod":
			return Hx_Field_haxe__int64_int64_impl__mod
		case "modInt":
			return Hx_Field_haxe__int64_int64_impl__modInt
		case "intMod":
			return Hx_Field_haxe__int64_int64_impl__intMod
		case "eq":
			return Hx_Field_haxe__int64_int64_impl__eq
		case "eqInt":
			return Hx_Field_haxe__int64_int64_impl__eqInt
		case "neq":
			return Hx_Field_haxe__int64_int64_impl__neq
		case "neqInt":
			return Hx_Field_haxe__int64_int64_impl__neqInt
		case "lt":
			return Hx_Field_haxe__int64_int64_impl__lt
		case "ltInt":
			return Hx_Field_haxe__int64_int64_impl__ltInt
		case "intLt":
			return Hx_Field_haxe__int64_int64_impl__intLt
		case "lte":
			return Hx_Field_haxe__int64_int64_impl__lte
		case "lteInt":
			return Hx_Field_haxe__int64_int64_impl__lteInt
		case "intLte":
			return Hx_Field_haxe__int64_int64_impl__intLte
		case "gt":
			return Hx_Field_haxe__int64_int64_impl__gt
		case "gtInt":
			return Hx_Field_haxe__int64_int64_impl__gtInt
		case "intGt":
			return Hx_Field_haxe__int64_int64_impl__intGt
		case "gte":
			return Hx_Field_haxe__int64_int64_impl__gte
		case "gteInt":
			return Hx_Field_haxe__int64_int64_impl__gteInt
		case "intGte":
			return Hx_Field_haxe__int64_int64_impl__intGte
		case "complement":
			return Hx_Field_haxe__int64_int64_impl__complement
		case "and":
			return Hx_Field_haxe__int64_int64_impl__and
		case "or":
			return Hx_Field_haxe__int64_int64_impl__or
		case "xor":
			return Hx_Field_haxe__int64_int64_impl__xor
		case "shl":
			return Hx_Field_haxe__int64_int64_impl__shl
		case "shr":
			return Hx_Field_haxe__int64_int64_impl__shr
		case "ushr":
			return Hx_Field_haxe__int64_int64_impl__ushr
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__int64_int64_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__int64_int64_impl__RTTI
}

func (this *Hx_Obj_haxe__int64_int64_impl_) Hx_As_Hx_Obj_haxe__int64_int64_impl_() *Hx_Obj_haxe__int64_int64_impl_ {
	return this
}

var Hx_Field_haxe__int64_int64_impl__MASK int64

func Hx_Init_Hx_Field_haxe__int64_int64_impl__MASK() int64 {
	var v int64
	_ = v
	v = int64(((int32)(65535)))
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = v
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = v
	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = _hx_tmp_0
	var _hx_tmp_3 int64
	_ = _hx_tmp_3
	_hx_tmp_3 = _hx_tmp_1
	return (_hx_tmp_2 | (_hx_tmp_3 << int64(((int32)(16)))))
}

func Hx_Field_haxe__int64_int64_impl__make(high int32, low int32) Hx_Typedef_haxe__int64___int64 {
	var h int64
	_ = h
	h = int64(high)
	var l int64
	_ = l
	l = int64(low)
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = h
	return ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_0 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
}

func Hx_Field_haxe__int64_int64_impl___hx_new(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return x
}

func Hx_Field_haxe__int64_int64_impl__get_val(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return this
}

func Hx_Field_haxe__int64_int64_impl__set_val(this Hx_Typedef_haxe__int64___int64, x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	this = x
	return this
}

func Hx_Field_haxe__int64_int64_impl__get_high(this Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	return ((int32)((_hx_tmp_0 >> int64(((int32)(32))))))
}

func Hx_Field_haxe__int64_int64_impl__get_low(this Hx_Typedef_haxe__int64___int64) int32 {
	return ((int32)(this))
}

func Hx_Field_haxe__int64_int64_impl__copy(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return this
}

func Hx_Field_haxe__int64_int64_impl__ofInt(x int32) Hx_Typedef_haxe__int64___int64 {
	var _v int32
	_ = _v
	_v = x
	return ((Hx_Typedef_haxe__int64___int64)(_v))
}

func Hx_Field_haxe__int64_int64_impl__fromInt(x int32) Hx_Typedef_haxe__int64___int64 {
	var _v int32
	_ = _v
	_v = x
	return ((Hx_Typedef_haxe__int64___int64)(_v))
}

func Hx_Field_haxe__int64_int64_impl__toFloat(this Hx_Typedef_haxe__int64___int64) float64 {
	return Hx_Field_haxe_int64helper_toFloat(this)
}

func Hx_Field_haxe__int64_int64_impl__is(val any) bool {
	return false
}

func Hx_Field_haxe__int64_int64_impl__isInt64(val any) bool {
	return false
}

func Hx_Field_haxe__int64_int64_impl__toInt(x Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_1
	_hx_tmp_1 = x
	if _hx_tmp_1 < int64(((int32)(-2147483648))) {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 Hx_Typedef_haxe__int64___int64
		_ = _hx_tmp_2
		_hx_tmp_2 = x
		var _hx_tmp_3 Hx_Typedef_haxe__int64___int64
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 > int64(((int32)(2147483647))))
	}

	if _hx_tmp_0 {
		panic(Hx_Field_haxe_exception_thrown(((any)("Overflow"))))
	}

	return ((int32)(x))
}

func Hx_Field_haxe__int64_int64_impl__getHigh(x Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(x))
	return ((int32)((_hx_tmp_0 >> int64(((int32)(32))))))
}

func Hx_Field_haxe__int64_int64_impl__getLow(x Hx_Typedef_haxe__int64___int64) int32 {
	return ((int32)(x))
}

func Hx_Field_haxe__int64_int64_impl__isNeg(x Hx_Typedef_haxe__int64___int64) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = x
	return (_hx_tmp_0 < int64(((int32)(0))))
}

func Hx_Field_haxe__int64_int64_impl__isZero(x Hx_Typedef_haxe__int64___int64) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = x
	return (_hx_tmp_0 == int64(((int32)(0))))
}

func Hx_Field_haxe__int64_int64_impl__compare(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) int32 {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

func Hx_Field_haxe__int64_int64_impl__ucompare(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	if _hx_tmp_0 < int64(((int32)(0))) {
		var _hx_tmp_1 Hx_Typedef_haxe__int64___int64
		_ = _hx_tmp_1
		_hx_tmp_1 = b
		if _hx_tmp_1 < int64(((int32)(0))) {
			if a < b {
				return -1
			} else {
				if a > b {
					return 1
				} else {
					return 0
				}
			}
		} else {
			return 1
		}
	}

	var _hx_tmp_2 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_2
	_hx_tmp_2 = b
	if _hx_tmp_2 < int64(((int32)(0))) {
		return -1
	} else {
		if a < b {
			return -1
		} else {
			if a > b {
				return 1
			} else {
				return 0
			}
		}
	}
}

func Hx_Field_haxe__int64_int64_impl__toStr(x Hx_Typedef_haxe__int64___int64) string {
	return Hx_Field_go_haxe_hxdynamic_stringConcat("", Hx_Field_std_string(x))
}

func Hx_Field_haxe__int64_int64_impl__divMod(dividend Hx_Typedef_haxe__int64___int64, divisor Hx_Typedef_haxe__int64___int64) any {
	return any(map[string]any{"quotient": ((any)(((Hx_Typedef_haxe__int64___int64)((((int64)(dividend)) / ((int64)(divisor))))))), "modulus": ((any)(((Hx_Typedef_haxe__int64___int64)((((int64)(dividend)) % ((int64)(divisor)))))))})
}

func Hx_Field_haxe__int64_int64_impl__toString(this Hx_Typedef_haxe__int64___int64) string {
	return Hx_Field_go_haxe_hxdynamic_stringConcat("", Hx_Field_std_string(this))
}

func Hx_Field_haxe__int64_int64_impl__parseString(sParam string) Hx_Typedef_haxe__int64___int64 {
	return Hx_Field_haxe_int64helper_parseString(sParam)
}

func Hx_Field_haxe__int64_int64_impl__fromFloat(f float64) Hx_Typedef_haxe__int64___int64 {
	return Hx_Field_haxe_int64helper_fromFloat(f)
}

func Hx_Field_haxe__int64_int64_impl__neg(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)(-((int64)(x))))
}

func Hx_Field_haxe__int64_int64_impl__preIncrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	this = (this + ((Hx_Typedef_haxe__int64___int64)(1)))
	return this
}

func Hx_Field_haxe__int64_int64_impl__postIncrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	this = (this + ((Hx_Typedef_haxe__int64___int64)(1)))
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(this)) - ((int64)(1)))))
}

func Hx_Field_haxe__int64_int64_impl__preDecrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	this = (this - ((Hx_Typedef_haxe__int64___int64)(1)))
	return this
}

func Hx_Field_haxe__int64_int64_impl__postDecrement(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	this = (this - ((Hx_Typedef_haxe__int64___int64)(1)))
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(this)) + ((int64)(1)))))
}

func Hx_Field_haxe__int64_int64_impl__add(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) + ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__addInt(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 + int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__sub(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) - ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__subInt(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 - int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__intSub(a int32, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((int64(a) - ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__mul(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) * ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__mulInt(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 * int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__div(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) / ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__divInt(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 / int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__intDiv(a int32, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((int64(a) / ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__mod(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) % ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__modInt(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 % int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__intMod(a int32, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((int64(a) % ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__eq(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a == b)
}

func Hx_Field_haxe__int64_int64_impl__eqInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 == int64(b))
}

func Hx_Field_haxe__int64_int64_impl__neq(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a != b)
}

func Hx_Field_haxe__int64_int64_impl__neqInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 != int64(b))
}

func Hx_Field_haxe__int64_int64_impl__lt(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a < b)
}

func Hx_Field_haxe__int64_int64_impl__ltInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 < int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intLt(a int32, b Hx_Typedef_haxe__int64___int64) bool {
	return (int64(a) < b)
}

func Hx_Field_haxe__int64_int64_impl__lte(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a <= b)
}

func Hx_Field_haxe__int64_int64_impl__lteInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 <= int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intLte(a int32, b Hx_Typedef_haxe__int64___int64) bool {
	return (int64(a) <= b)
}

func Hx_Field_haxe__int64_int64_impl__gt(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a > b)
}

func Hx_Field_haxe__int64_int64_impl__gtInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 > int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intGt(a int32, b Hx_Typedef_haxe__int64___int64) bool {
	return (int64(a) > b)
}

func Hx_Field_haxe__int64_int64_impl__gte(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (a >= b)
}

func Hx_Field_haxe__int64_int64_impl__gteInt(a Hx_Typedef_haxe__int64___int64, b int32) bool {
	var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
	_ = _hx_tmp_0
	_hx_tmp_0 = a
	return (_hx_tmp_0 >= int64(b))
}

func Hx_Field_haxe__int64_int64_impl__intGte(a int32, b Hx_Typedef_haxe__int64___int64) bool {
	return (int64(a) >= b)
}

func Hx_Field_haxe__int64_int64_impl__complement(x Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)(^((int64)(x))))
}

func Hx_Field_haxe__int64_int64_impl__and(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) & ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__or(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) | ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__xor(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return ((Hx_Typedef_haxe__int64___int64)((((int64)(a)) ^ ((int64)(b)))))
}

func Hx_Field_haxe__int64_int64_impl__shl(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	if (b <= 0) || (b >= 64) {
		return a
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 << int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__shr(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	if (b <= 0) || (b >= 64) {
		return a
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)((_hx_tmp_0 >> int64(b))))
}

func Hx_Field_haxe__int64_int64_impl__ushr(a Hx_Typedef_haxe__int64___int64, b int32) Hx_Typedef_haxe__int64___int64 {
	if (b <= 0) || (b >= 64) {
		return a
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(a))
	return ((Hx_Typedef_haxe__int64___int64)(((int64)((((uint64)(_hx_tmp_0)) >> int64(b))))))
}
