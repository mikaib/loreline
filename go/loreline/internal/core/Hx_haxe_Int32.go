package core

var Hx_Abstract_haxe_int32_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Int32",
	HxMakeArray[string]("negate", "preIncrement", "postIncrement", "preDecrement", "postDecrement", "add", "addInt", "addFloat", "sub", "subInt", "intSub", "subFloat", "floatSub", "mul", "mulInt", "mulFloat", "div", "divInt", "intDiv", "divFloat", "floatDiv", "mod", "modInt", "intMod", "modFloat", "floatMod", "eq", "eqInt", "eqFloat", "neq", "neqInt", "neqFloat", "lt", "ltInt", "intLt", "ltFloat", "floatLt", "lte", "lteInt", "intLte", "lteFloat", "floatLte", "gt", "gtInt", "intGt", "gtFloat", "floatGt", "gte", "gteInt", "intGte", "gteFloat", "floatGte", "complement", "and", "andInt", "or", "orInt", "xor", "xorInt", "shr", "shrInt", "intShr", "ushr", "ushrInt", "intUshr", "shl", "shlInt", "intShl", "toFloat", "ucompare", "clamp"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe__int32_int32_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__int32_int32_impl_() {
}

type Hx_Obj_VTable_haxe__int32_int32_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__int32_int32_impl_ struct {
	VTable Hx_Obj_VTable_haxe__int32_int32_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__int32_int32_impl__CreateEmptyInstance() *Hx_Obj_haxe__int32_int32_impl_ {
	obj := &Hx_Obj_haxe__int32_int32_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__int32_int32_impl__CreateInstance() *Hx_Obj_haxe__int32_int32_impl_ {
	obj := Hx_Obj_haxe__int32_int32_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__int32_int32_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._Int32.Int32_Impl_",
	HxMakeArray[string]("negate", "preIncrement", "postIncrement", "preDecrement", "postDecrement", "add", "addInt", "addFloat", "sub", "subInt", "intSub", "subFloat", "floatSub", "mul", "mulInt", "mulFloat", "div", "divInt", "intDiv", "divFloat", "floatDiv", "mod", "modInt", "intMod", "modFloat", "floatMod", "eq", "eqInt", "eqFloat", "neq", "neqInt", "neqFloat", "lt", "ltInt", "intLt", "ltFloat", "floatLt", "lte", "lteInt", "intLte", "lteFloat", "floatLte", "gt", "gtInt", "intGt", "gtFloat", "floatGt", "gte", "gteInt", "intGte", "gteFloat", "floatGte", "complement", "and", "andInt", "or", "orInt", "xor", "xorInt", "shr", "shrInt", "intShr", "ushr", "ushrInt", "intUshr", "shl", "shlInt", "intShl", "toFloat", "ucompare", "clamp"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__int32_int32_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__int32_int32_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__int32_int32_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "negate":
			return Hx_Field_haxe__int32_int32_impl__negate
		case "preIncrement":
			return Hx_Field_haxe__int32_int32_impl__preIncrement
		case "postIncrement":
			return Hx_Field_haxe__int32_int32_impl__postIncrement
		case "preDecrement":
			return Hx_Field_haxe__int32_int32_impl__preDecrement
		case "postDecrement":
			return Hx_Field_haxe__int32_int32_impl__postDecrement
		case "add":
			return Hx_Field_haxe__int32_int32_impl__add
		case "addInt":
			return Hx_Field_haxe__int32_int32_impl__addInt
		case "sub":
			return Hx_Field_haxe__int32_int32_impl__sub
		case "subInt":
			return Hx_Field_haxe__int32_int32_impl__subInt
		case "intSub":
			return Hx_Field_haxe__int32_int32_impl__intSub
		case "toFloat":
			return Hx_Field_haxe__int32_int32_impl__toFloat
		case "ucompare":
			return Hx_Field_haxe__int32_int32_impl__ucompare
		case "clamp":
			return Hx_Field_haxe__int32_int32_impl__clamp
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__int32_int32_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__int32_int32_impl__RTTI
}

func (this *Hx_Obj_haxe__int32_int32_impl_) Hx_As_Hx_Obj_haxe__int32_int32_impl_() *Hx_Obj_haxe__int32_int32_impl_ {
	return this
}

func Hx_Field_haxe__int32_int32_impl__negate(this int32) int32 {
	return (^this + 1)
}

func Hx_Field_haxe__int32_int32_impl__preIncrement(this int32) int32 {
	this = (this + 1)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	this = _hx_tmp_0
	return this
}

func Hx_Field_haxe__int32_int32_impl__postIncrement(this int32) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	this = (this + 1)
	var ret int32
	_ = ret
	ret = _hx_tmp_0
	return ret
}

func Hx_Field_haxe__int32_int32_impl__preDecrement(this int32) int32 {
	this = (this - 1)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	this = _hx_tmp_0
	return this
}

func Hx_Field_haxe__int32_int32_impl__postDecrement(this int32) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	this = (this - 1)
	var ret int32
	_ = ret
	ret = _hx_tmp_0
	return ret
}

func Hx_Field_haxe__int32_int32_impl__add(a int32, b int32) int32 {
	return (a + b)
}

func Hx_Field_haxe__int32_int32_impl__addInt(a int32, b int32) int32 {
	return (a + b)
}

func Hx_Field_haxe__int32_int32_impl__sub(a int32, b int32) int32 {
	return (a - b)
}

func Hx_Field_haxe__int32_int32_impl__subInt(a int32, b int32) int32 {
	return (a - b)
}

func Hx_Field_haxe__int32_int32_impl__intSub(a int32, b int32) int32 {
	return (a - b)
}

func Hx_Field_haxe__int32_int32_impl__toFloat(this int32) float64 {
	return ((float64)(this))
}

func Hx_Field_haxe__int32_int32_impl__ucompare(a int32, b int32) int32 {
	if a < 0 {
		if b < 0 {
			return ((int32)((((int32)(^b)) - ((int32)(^a)))))
		} else {
			return 1
		}
	}

	if b < 0 {
		return -1
	} else {
		return ((int32)((a - b)))
	}
}

func Hx_Field_haxe__int32_int32_impl__clamp(x int32) int32 {
	return x
}
