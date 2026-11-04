package core

import "unicode/utf8"

type Hx_Typedef_loreline_lorscript_typecheck = any

type Hx_Enum_loreline_lorscript_ttype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_ttype()
}

type Hx_Enum_loreline_lorscript_ttype_TVoid struct{}

func (this Hx_Enum_loreline_lorscript_ttype_TVoid) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TVoid) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TVoid) Hx_Field_enumIndex() int32           { return 1 }
func (this Hx_Enum_loreline_lorscript_ttype_TVoid) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TVoid) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_ttype_TVoid) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TUnresolved struct {
	Hx_Field_name string
}

func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) Hx_Field_enumIndex() int32           { return 7 }
func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_lorscript_ttype_TUnresolved) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TType struct {
	Hx_Field_t    any
	Hx_Field_args HxArray[Hx_Enum_loreline_lorscript_ttype]
}

func (this Hx_Enum_loreline_lorscript_ttype_TType) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TType) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TType) Hx_Field_enumIndex() int32           { return 11 }
func (this Hx_Enum_loreline_lorscript_ttype_TType) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TType) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t), any(this.Hx_Field_args))
}
func (this Hx_Enum_loreline_lorscript_ttype_TType) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	case 1:
		return any(this.Hx_Field_args)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TParam struct {
	Hx_Field_name string
}

func (this Hx_Enum_loreline_lorscript_ttype_TParam) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TParam) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TParam) Hx_Field_enumIndex() int32           { return 6 }
func (this Hx_Enum_loreline_lorscript_ttype_TParam) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TParam) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_lorscript_ttype_TParam) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TNull struct {
	Hx_Field_t Hx_Enum_loreline_lorscript_ttype
}

func (this Hx_Enum_loreline_lorscript_ttype_TNull) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TNull) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TNull) Hx_Field_enumIndex() int32           { return 8 }
func (this Hx_Enum_loreline_lorscript_ttype_TNull) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TNull) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ttype_TNull) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TMono struct {
	Hx_Field_r any
}

func (this Hx_Enum_loreline_lorscript_ttype_TMono) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TMono) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TMono) Hx_Field_enumIndex() int32           { return 0 }
func (this Hx_Enum_loreline_lorscript_ttype_TMono) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TMono) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_r))
}
func (this Hx_Enum_loreline_lorscript_ttype_TMono) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_r)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TLazy struct {
	Hx_Field_f func() Hx_Enum_loreline_lorscript_ttype
}

func (this Hx_Enum_loreline_lorscript_ttype_TLazy) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TLazy) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TLazy) Hx_Field_enumIndex() int32           { return 15 }
func (this Hx_Enum_loreline_lorscript_ttype_TLazy) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TLazy) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_f))
}
func (this Hx_Enum_loreline_lorscript_ttype_TLazy) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_f)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TInt struct{}

func (this Hx_Enum_loreline_lorscript_ttype_TInt) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TInt) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TInt) Hx_Field_enumIndex() int32           { return 2 }
func (this Hx_Enum_loreline_lorscript_ttype_TInt) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TInt) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_ttype_TInt) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TInst struct {
	Hx_Field_c    any
	Hx_Field_args HxArray[Hx_Enum_loreline_lorscript_ttype]
}

func (this Hx_Enum_loreline_lorscript_ttype_TInst) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TInst) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TInst) Hx_Field_enumIndex() int32           { return 9 }
func (this Hx_Enum_loreline_lorscript_ttype_TInst) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TInst) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c), any(this.Hx_Field_args))
}
func (this Hx_Enum_loreline_lorscript_ttype_TInst) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	case 1:
		return any(this.Hx_Field_args)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TFun struct {
	Hx_Field_args HxArray[any]
	Hx_Field_ret  Hx_Enum_loreline_lorscript_ttype
}

func (this Hx_Enum_loreline_lorscript_ttype_TFun) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TFun) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TFun) Hx_Field_enumIndex() int32           { return 13 }
func (this Hx_Enum_loreline_lorscript_ttype_TFun) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TFun) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_args), any(this.Hx_Field_ret))
}
func (this Hx_Enum_loreline_lorscript_ttype_TFun) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_args)
	case 1:
		return any(this.Hx_Field_ret)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TFloat struct{}

func (this Hx_Enum_loreline_lorscript_ttype_TFloat) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TFloat) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TFloat) Hx_Field_enumIndex() int32           { return 3 }
func (this Hx_Enum_loreline_lorscript_ttype_TFloat) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TFloat) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_ttype_TFloat) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TEnum struct {
	Hx_Field_e    any
	Hx_Field_args HxArray[Hx_Enum_loreline_lorscript_ttype]
}

func (this Hx_Enum_loreline_lorscript_ttype_TEnum) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TEnum) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TEnum) Hx_Field_enumIndex() int32           { return 10 }
func (this Hx_Enum_loreline_lorscript_ttype_TEnum) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TEnum) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_args))
}
func (this Hx_Enum_loreline_lorscript_ttype_TEnum) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_args)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TDynamic struct{}

func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) Hx_Field_enumIndex() int32           { return 5 }
func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_ttype_TDynamic) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TBool struct{}

func (this Hx_Enum_loreline_lorscript_ttype_TBool) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TBool) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TBool) Hx_Field_enumIndex() int32           { return 4 }
func (this Hx_Enum_loreline_lorscript_ttype_TBool) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TBool) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_ttype_TBool) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TAnon struct {
	Hx_Field_fields HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_ttype_TAnon) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TAnon) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TAnon) Hx_Field_enumIndex() int32           { return 14 }
func (this Hx_Enum_loreline_lorscript_ttype_TAnon) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TAnon) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_fields))
}
func (this Hx_Enum_loreline_lorscript_ttype_TAnon) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_fields)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ttype_TAbstract struct {
	Hx_Field_a    any
	Hx_Field_args HxArray[Hx_Enum_loreline_lorscript_ttype]
}

func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) M_Hx_Enum_loreline_lorscript_ttype() {}
func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) Hx_Field_enumIndex() int32           { return 12 }
func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ttype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_a), any(this.Hx_Field_args))
}
func (this Hx_Enum_loreline_lorscript_ttype_TAbstract) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_a)
	case 1:
		return any(this.Hx_Field_args)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_ttype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.TType",
	HxMakeArray[string]("TMono", "TVoid", "TInt", "TFloat", "TBool", "TDynamic", "TParam", "TUnresolved", "TNull", "TInst", "TEnum", "TType", "TAbstract", "TFun", "TAnon", "TLazy"),
	HxMakeArray[int32](1, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_ttype_TMono{params.Get_Dyn(0).(any)}
		case 1:
			return Hx_Enum_loreline_lorscript_ttype_TVoid{}
		case 2:
			return Hx_Enum_loreline_lorscript_ttype_TInt{}
		case 3:
			return Hx_Enum_loreline_lorscript_ttype_TFloat{}
		case 4:
			return Hx_Enum_loreline_lorscript_ttype_TBool{}
		case 5:
			return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
		case 6:
			return Hx_Enum_loreline_lorscript_ttype_TParam{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 7:
			return Hx_Enum_loreline_lorscript_ttype_TUnresolved{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 8:
			return Hx_Enum_loreline_lorscript_ttype_TNull{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_ttype)}
		case 9:
			return Hx_Enum_loreline_lorscript_ttype_TInst{params.Get_Dyn(0).(any), HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](params.Get_Dyn(1))}
		case 10:
			return Hx_Enum_loreline_lorscript_ttype_TEnum{params.Get_Dyn(0).(any), HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](params.Get_Dyn(1))}
		case 11:
			return Hx_Enum_loreline_lorscript_ttype_TType{params.Get_Dyn(0).(any), HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](params.Get_Dyn(1))}
		case 12:
			return Hx_Enum_loreline_lorscript_ttype_TAbstract{params.Get_Dyn(0).(any), HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](params.Get_Dyn(1))}
		case 13:
			return Hx_Enum_loreline_lorscript_ttype_TFun{HxAnyToTypedArray[any](params.Get_Dyn(0)), params.Get_Dyn(1).(Hx_Enum_loreline_lorscript_ttype)}
		case 14:
			return Hx_Enum_loreline_lorscript_ttype_TAnon{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 15:
			return Hx_Enum_loreline_lorscript_ttype_TLazy{params.Get_Dyn(0).(func() Hx_Enum_loreline_lorscript_ttype)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Enum_loreline_lorscript__checker_withtype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript__checker_withtype()
}

type Hx_Enum_loreline_lorscript__checker_withtype_WithType struct {
	Hx_Field_t Hx_Enum_loreline_lorscript_ttype
}

func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) M_Hx_Enum_loreline_lorscript__checker_withtype() {
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) Hx_Field_enumIndex() int32 {
	return 2
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__checker_withtype_RTTI
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_WithType) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript__checker_withtype_Value struct{}

func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) M_Hx_Enum_loreline_lorscript__checker_withtype() {
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) Hx_Field_enumIndex() int32 { return 1 }
func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__checker_withtype_RTTI
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_Value) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript__checker_withtype_NoValue struct{}

func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) M_Hx_Enum_loreline_lorscript__checker_withtype() {
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) Hx_Field_enumIndex() int32 { return 0 }
func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__checker_withtype_RTTI
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript__checker_withtype_NoValue) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript__checker_withtype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript._Checker.WithType",
	HxMakeArray[string]("NoValue", "Value", "WithType"),
	HxMakeArray[int32](0, 0, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}
		case 1:
			return Hx_Enum_loreline_lorscript__checker_withtype_Value{}
		case 2:
			return Hx_Enum_loreline_lorscript__checker_withtype_WithType{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_ttype)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Enum_loreline_lorscript_ctypedecl interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_ctypedecl()
}

type Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef struct {
	Hx_Field_t any
}

func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) M_Hx_Enum_loreline_lorscript_ctypedecl() {}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctypedecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctypedecl_CTEnum struct {
	Hx_Field_e any
}

func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) M_Hx_Enum_loreline_lorscript_ctypedecl() {}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctypedecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTEnum) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctypedecl_CTClass struct {
	Hx_Field_c any
}

func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) M_Hx_Enum_loreline_lorscript_ctypedecl() {}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctypedecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTClass) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctypedecl_CTAlias struct {
	Hx_Field_t Hx_Enum_loreline_lorscript_ttype
}

func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) M_Hx_Enum_loreline_lorscript_ctypedecl() {}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctypedecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAlias) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract struct {
	Hx_Field_a any
}

func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) M_Hx_Enum_loreline_lorscript_ctypedecl() {
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) Hx_Field_enumIndex() int32 { return 4 }
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctypedecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_a))
}
func (this Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_a)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_ctypedecl_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.CTypedecl",
	HxMakeArray[string]("CTClass", "CTEnum", "CTTypedef", "CTAlias", "CTAbstract"),
	HxMakeArray[int32](1, 1, 1, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_ctypedecl_CTClass{params.Get_Dyn(0).(any)}
		case 1:
			return Hx_Enum_loreline_lorscript_ctypedecl_CTEnum{params.Get_Dyn(0).(any)}
		case 2:
			return Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef{params.Get_Dyn(0).(any)}
		case 3:
			return Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_ttype)}
		case 4:
			return Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_loreline_lorscript_cmetadata = HxArray[any]

type Hx_Typedef_loreline_lorscript_cnamedtype = any

type Hx_Typedef_loreline_lorscript_cclass = any

type Hx_Typedef_loreline_lorscript_cfield = any

type Hx_Typedef_loreline_lorscript_cenum = any

type Hx_Typedef_loreline_lorscript_ctypedef = any

type Hx_Typedef_loreline_lorscript_cabstract = any

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_completion() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_completion() {
}

type Hx_Obj_VTable_loreline_lorscript_completion interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_completion struct {
	VTable        Hx_Obj_VTable_loreline_lorscript_completion `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_expr any
	Hx_Field_t    Hx_Enum_loreline_lorscript_ttype
}

func Hx_Obj_loreline_lorscript_completion_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_completion {
	obj := &Hx_Obj_loreline_lorscript_completion{}
	obj.VTable = obj
	obj.Hx_Field_expr = HxDefault[any]()
	obj.Hx_Field_t = HxDefault[Hx_Enum_loreline_lorscript_ttype]()
	return obj
}

func Hx_Obj_loreline_lorscript_completion_CreateInstance(expr any, t Hx_Enum_loreline_lorscript_ttype) *Hx_Obj_loreline_lorscript_completion {
	obj := Hx_Obj_loreline_lorscript_completion_CreateEmptyInstance()
	obj.Hx_New(expr, t)
	return obj
}

func (this *Hx_Obj_loreline_lorscript_completion) Hx_New(expr any, t Hx_Enum_loreline_lorscript_ttype) {
	{
		this.Hx_Field_expr = expr
		this.Hx_Field_t = t
	}

}

var Hx_Obj_loreline_lorscript_completion_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Completion",
	HxMakeArray[string](),
	HxMakeArray[string]("expr", "t"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_completion_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[Hx_Enum_loreline_lorscript_ttype](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_lorscript_completion_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_completion_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_completion) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_completion_RTTI
}

func (this *Hx_Obj_loreline_lorscript_completion) Hx_As_Hx_Obj_loreline_lorscript_completion() *Hx_Obj_loreline_lorscript_completion {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_checkertypes() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_checkertypes() {
}

type Hx_Obj_VTable_loreline_lorscript_checkertypes interface {
	Hx_Field_addMeta(src any, to any)
	Hx_Field_makeXmlType(t Hx_Enum_haxe_rtti_ctype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_getType(name string, args__isOptional any) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_resolve(name string, args__isOptional any) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_addXmlApi(api *Hx_Obj_xml)
	Hx_Field_defineClass(name string, ct__isOptional any) any
	Hx_Field_addXmlType(x Hx_Enum_haxe_rtti_typetree, todo HxArray[func()])
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_checkertypes struct {
	VTable               Hx_Obj_VTable_loreline_lorscript_checkertypes `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_types       Hx_Typedef_map
	Hx_Field_t_string    Hx_Enum_loreline_lorscript_ttype
	Hx_Field_localParams Hx_Typedef_map
	Hx_Field_parser      *Hx_Obj_loreline_lorscript_parser
}

func Hx_Obj_loreline_lorscript_checkertypes_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_checkertypes {
	obj := &Hx_Obj_loreline_lorscript_checkertypes{}
	obj.VTable = obj
	obj.Hx_Field_types = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_t_string = HxDefault[Hx_Enum_loreline_lorscript_ttype]()
	obj.Hx_Field_localParams = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_parser = HxDefault[*Hx_Obj_loreline_lorscript_parser]()
	return obj
}

func Hx_Obj_loreline_lorscript_checkertypes_CreateInstance() *Hx_Obj_loreline_lorscript_checkertypes {
	obj := Hx_Obj_loreline_lorscript_checkertypes_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_New() {
	{
		this.Hx_Field_types = (func() *Hx_Obj_haxe_imap {
			_hx_icast_70 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_70 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_70, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_70.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_types = (func() *Hx_Obj_haxe_imap {
			_hx_icast_75 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_75 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_75, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_75.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_406 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_406
			_hx_closure_src_406 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_406(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()("Void", Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{Hx_Enum_loreline_lorscript_ttype_TVoid{}})
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_407 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_407
			_hx_closure_src_407 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_407(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()("Int", Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{Hx_Enum_loreline_lorscript_ttype_TInt{}})
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_408 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_408
			_hx_closure_src_408 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_408(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()("Float", Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{Hx_Enum_loreline_lorscript_ttype_TFloat{}})
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_409 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_409
			_hx_closure_src_409 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_409(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()("Bool", Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{Hx_Enum_loreline_lorscript_ttype_TBool{}})
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_410 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_410
			_hx_closure_src_410 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_410(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()("Dynamic", Hx_Enum_loreline_lorscript_ctypedecl_CTAlias{Hx_Enum_loreline_lorscript_ttype_TDynamic{}})
		this.Hx_Field_parser = Hx_Obj_loreline_lorscript_parser_CreateInstance()
	}

}

var Hx_Obj_loreline_lorscript_checkertypes_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.CheckerTypes",
	HxMakeArray[string](),
	HxMakeArray[string]("types", "t_string", "localParams", "parser", "addXmlApi", "defineClass", "addXmlType", "addMeta", "makeXmlType", "getType", "resolve"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_checkertypes_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lorscript_checkertypes_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_checkertypes_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_checkertypes_RTTI
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_As_Hx_Obj_loreline_lorscript_checkertypes() *Hx_Obj_loreline_lorscript_checkertypes {
	return this
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_addXmlApi(api *Hx_Obj_xml) {
	var types *Hx_Obj_haxe_rtti_xmlparser
	_ = types
	types = Hx_Obj_haxe_rtti_xmlparser_CreateInstance()
	types.VTable.Hx_Field_process(api, "")
	var todo HxArray[func()]
	_ = todo
	todo = HxMakeArray[func()]()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 Hx_Typedef_haxe_rtti_typeroot
		_ = _g1
		_g1 = types.Hx_Field_root
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var v Hx_Enum_haxe_rtti_typetree
			_ = v
			v = _g1.Get(_g)
			_g++
			this.VTable.Hx_Field_addXmlType(v, todo)
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g2
			if !(_hx_tmp_1 < todo.Len()) {
				break
			}

			var f func()
			_ = f
			f = todo.Get(_g2)
			_g2++
			f()
		}
	}

	this.Hx_Field_t_string = this.VTable.Hx_Field_getType("String", nil)
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_defineClass(name string, ct__isOptional any) any {
	ct := HxDefault[any]()
	_ = ct

	if ct__isOptional != nil {
		ct = HxConvert[any](ct__isOptional)
	} else {
		ct = nil
	}
	{
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(ct, nil))) {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = ((any)(name))
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = ((any)(Hx_Obj_haxe_ds_stringmap_CreateInstance()))
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = ((any)(Hx_Obj_haxe_ds_stringmap_CreateInstance()))
			ct = any(map[string]any{"name": _hx_tmp_0, "fields": _hx_tmp_1, "statics": _hx_tmp_2, "params": ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()))})
		}

		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = name
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
			var _hx_closure_src_389 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_389
			_hx_closure_src_389 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
				_hx_closure_src_389(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()(_hx_tmp_3, Hx_Enum_loreline_lorscript_ctypedecl_CTClass{ct})
		return ct
	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_addXmlType(x Hx_Enum_haxe_rtti_typetree, todo HxArray[func()]) {
	var _gthis *Hx_Obj_loreline_lorscript_checkertypes
	_ = _gthis
	_gthis = this
	switch x.Hx_Field_enumIndex() {
	case 0:
		{
			var name string
			_ = name
			name = x.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_name
			var full string
			_ = full
			full = x.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_full
			var subs HxArray[Hx_Enum_haxe_rtti_typetree]
			_ = subs
			subs = x.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_subs
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = _g
					if !(_hx_tmp_0 < subs.Len()) {
						break
					}

					var s Hx_Enum_haxe_rtti_typetree
					_ = s
					s = subs.Get(_g)
					_g++
					this.VTable.Hx_Field_addXmlType(s, todo)
				}
			}
		}

	case 1:
		{
			var c any
			_ = c
			c = x.(Hx_Enum_haxe_rtti_typetree_TClassdecl).Hx_Field_c
			{
				if func() func(_hx_param_key string) bool {
					var _hx_closure_src_390 func(_hx_param_key any) bool
					_ = _hx_closure_src_390
					_hx_closure_src_390 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_390(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(c, "path"))) {
					return
				}

				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(c, "path")
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()))
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = ((any)(Hx_Obj_haxe_ds_stringmap_CreateInstance()))
				var cl any
				_ = cl
				cl = any(map[string]any{"name": _hx_tmp_1, "params": _hx_tmp_2, "fields": _hx_tmp_3, "statics": ((any)(Hx_Obj_haxe_ds_stringmap_CreateInstance()))})
				this.VTable.Hx_Field_addMeta(c, cl)
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(c, "isInterface")))) {
					Hx_Field_go_haxe_hxdynamic_setField(cl, "isInterface", true)
				}

				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 Hx_Typedef_haxe_rtti_typeparams
					_ = _g1
					_g1 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(c, "params"))
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g
						if !(_hx_tmp_4 < _g1.Len()) {
							break
						}

						var p string
						_ = p
						p = _g1.Get(_g)
						_g++
						var _hx_tmp_5 any
						_ = _hx_tmp_5
						_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(cl, "params"), "push")
						Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any](Hx_Enum_loreline_lorscript_ttype_TParam{p}).Dyn())
					}
				}

				var _hx_tmp_6 HxArray[func()]
				_ = _hx_tmp_6
				_hx_tmp_6 = todo
				_ = (Hx_Array_Push[func()](_hx_tmp_6, func() {
					var _g *Hx_Obj_haxe_imap
					_ = _g
					_g = (func() *Hx_Obj_haxe_imap {
						_hx_icast_71 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
						if _hx_icast_71 == nil {
							return nil
						}
						Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_71, "Hx_Obj_haxe_imap")
						return &Hx_Obj_haxe_imap{VTable: _hx_icast_71.VTable.(Hx_Obj_VTable_haxe_imap)}
					})()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _g2
						_g2 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl, "params"))
						for {
							var _hx_tmp_7 int32
							_ = _hx_tmp_7
							_hx_tmp_7 = _g1
							if !(_hx_tmp_7 < _g2.Len()) {
								break
							}

							var t Hx_Enum_loreline_lorscript_ttype
							_ = t
							t = _g2.Get(_g1)
							_g1++
							var _hx_tmp_8 any
							_ = _hx_tmp_8
							_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(c, "path"), ((any)(".")))
							func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								var _hx_closure_src_391 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_391
								_hx_closure_src_391 = _g.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
									_hx_closure_src_391(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_8, ((any)(Hx_Field_loreline_lorscript_checker_typeStr(t))))), t)
						}
					}

					_gthis.Hx_Field_localParams = _g
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"), nil))) {
						var _gthis1 *Hx_Obj_loreline_lorscript_checkertypes
						_ = _gthis1
						_gthis1 = _gthis
						var c1 Hx_Typedef_haxe_rtti_path
						_ = c1
						c1 = ((Hx_Typedef_haxe_rtti_path)(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"), "path"))))
						var tmp__g_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = tmp__g_1
						tmp__g_1 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							var _g2 HxArray[Hx_Enum_haxe_rtti_ctype]
							_ = _g2
							_g2 = HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"), "params"))
							for {
								var _hx_tmp_9 int32
								_ = _hx_tmp_9
								_hx_tmp_9 = _g1
								if !(_hx_tmp_9 < _g2.Len()) {
									break
								}

								var t Hx_Enum_haxe_rtti_ctype
								_ = t
								t = _g2.Get(_g1)
								_g1++
								var _hx_tmp_10 HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = _hx_tmp_10
								_hx_tmp_10 = tmp__g_1
								_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_10, _gthis.VTable.Hx_Field_makeXmlType(t)))
							}
						}

						var _hx_tmp_11 any
						_ = _hx_tmp_11
						_hx_tmp_11 = cl
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_11, "superClass", _gthis1.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(c1), tmp__g_1))
					}

					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"), ((any)(nil))))) {
						var _hx_tmp_12 any
						_ = _hx_tmp_12
						_hx_tmp_12 = cl
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_12, "interfaces", HxMakeArray[Hx_Enum_loreline_lorscript_ttype]())
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g1 HxArray[any]
							_ = _g1
							_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"))
							for {
								var _hx_tmp_13 int32
								_ = _hx_tmp_13
								_hx_tmp_13 = tmp__g_1
								if !(_hx_tmp_13 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
									break
								}

								var i any
								_ = i
								i = _g1.Get(tmp__g_1)
								tmp__g_1++
								var cl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = cl1
								cl1 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl, "interfaces"))
								var _gthis1 *Hx_Obj_loreline_lorscript_checkertypes
								_ = _gthis1
								_gthis1 = _gthis
								var i1 Hx_Typedef_haxe_rtti_path
								_ = i1
								i1 = ((Hx_Typedef_haxe_rtti_path)(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(i, "path"))))
								var _g2 HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = _g2
								_g2 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
								{
									var _g3 int32
									_ = _g3
									_g3 = 0
									var _g4 HxArray[Hx_Enum_haxe_rtti_ctype]
									_ = _g4
									_g4 = HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](Hx_Field_go_haxe_hxdynamic_getField(i, "params"))
									for {
										var _hx_tmp_14 int32
										_ = _hx_tmp_14
										_hx_tmp_14 = _g3
										if !(_hx_tmp_14 < _g4.Len()) {
											break
										}

										var t Hx_Enum_haxe_rtti_ctype
										_ = t
										t = _g4.Get(_g3)
										_g3++
										var _hx_tmp_15 HxArray[Hx_Enum_loreline_lorscript_ttype]
										_ = _hx_tmp_15
										_hx_tmp_15 = _g2
										_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_15, _gthis.VTable.Hx_Field_makeXmlType(t)))
									}
								}

								var _hx_tmp_16 HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = _hx_tmp_16
								_hx_tmp_16 = cl1
								_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_16, _gthis1.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(i1), _g2)))
							}
						}
					}

					var pkeys HxArray[string]
					_ = pkeys
					pkeys = HxMakeArray[string]()
					var initField func(f any, fields any)
					_ = initField
					initField = func(f any, fields any) {
						var _hx_tmp_17 bool
						_ = _hx_tmp_17
						var _hx_tmp_18 bool
						_ = _hx_tmp_18
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(f, "isOverride"))) {
							_hx_tmp_18 = true
						} else {
							var _hx_tmp_19 any
							_ = _hx_tmp_19
							_hx_tmp_19 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), "substr")
							var _hx_tmp_20 any
							_ = _hx_tmp_20
							_hx_tmp_20 = _hx_tmp_19
							_hx_tmp_18 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_20, HxMakeArray[any](((int32)(0)), ((int32)(4))).Dyn()), ((any)("get_"))))))
						}

						if _hx_tmp_18 {
							_hx_tmp_17 = true
						} else {
							var _hx_tmp_21 any
							_ = _hx_tmp_21
							_hx_tmp_21 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), "substr")
							var _hx_tmp_22 any
							_ = _hx_tmp_22
							_hx_tmp_22 = _hx_tmp_21
							_hx_tmp_17 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_22, HxMakeArray[any](((int32)(0)), ((int32)(4))).Dyn()), ((any)("set_"))))))
						}

						if _hx_tmp_17 {
							return
						}

						var complete bool
						_ = complete
						complete = !Hx_Field_stringtools_startsWith(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name")), "__")
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g1 HxArray[any]
							_ = _g1
							_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(f, "meta"))
							for {
								var _hx_tmp_23 int32
								_ = _hx_tmp_23
								_hx_tmp_23 = tmp__g_1
								if !(_hx_tmp_23 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
									break
								}

								var m any
								_ = m
								m = _g1.Get(tmp__g_1)
								tmp__g_1++
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(m, "name"), ((any)(":noScript"))))) {
									return
								}

								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(m, "name"), ((any)(":noCompletion"))))) {
									complete = false
								}
							}
						}

						var fl bool
						_ = fl
						switch Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "set")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() {
						case 2:
							{
								fl = true
							}

						case 0, 4:
							{
								fl = true
							}

						default:
							{
								fl = false
							}
						}

						var _hx_tmp_24 any
						_ = _hx_tmp_24
						_hx_tmp_24 = Hx_Field_go_haxe_hxdynamic_getField(f, "isPublic")
						var _hx_tmp_25 any
						_ = _hx_tmp_25
						_hx_tmp_25 = ((any)(fl))
						var _hx_tmp_26 any
						_ = _hx_tmp_26
						_hx_tmp_26 = ((any)(complete))
						var _hx_tmp_27 any
						_ = _hx_tmp_27
						_hx_tmp_27 = ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()))
						var fl1 any
						_ = fl1
						fl1 = any(map[string]any{"isPublic": _hx_tmp_24, "canWrite": _hx_tmp_25, "complete": _hx_tmp_26, "params": _hx_tmp_27, "name": Hx_Field_go_haxe_hxdynamic_getField(f, "name"), "t": ((any)(nil))})
						{
							var _g2 int32
							_ = _g2
							_g2 = 0
							var _g3 Hx_Typedef_haxe_rtti_typeparams
							_ = _g3
							_g3 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(f, "params"))
							for {
								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								_hx_tmp_28 = _g2
								if !(_hx_tmp_28 < _g3.Len()) {
									break
								}

								var p string
								_ = p
								p = _g3.Get(_g2)
								_g2++
								var pt Hx_Enum_loreline_lorscript_ttype
								_ = pt
								pt = Hx_Enum_loreline_lorscript_ttype_TParam{p}

								var key string
								_ = key
								key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), ((any)("."))), ((any)(p))))
								_ = (Hx_Array_Push[string](pkeys, key))
								var _hx_tmp_29 any
								_ = _hx_tmp_29
								_hx_tmp_29 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(fl1, "params"), "push")
								Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_29, HxMakeArray[any](pt).Dyn())
								func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
									var _hx_closure_src_392 func(_hx_param_key any, _hx_param_value any)
									_ = _hx_closure_src_392
									_hx_closure_src_392 = _gthis.Hx_Field_localParams.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
									}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
									return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
										_hx_closure_src_392(((any)(_hx_param_key)), ((any)(_hx_param_value)))
									}
								}()(key, pt)
							}
						}

						var _hx_tmp_30 any
						_ = _hx_tmp_30
						_hx_tmp_30 = fl1
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_30, "t", _gthis.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "type"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_haxe_rtti_ctype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_haxe_rtti_ctype)
						}()))
						var _hx_tmp_31 bool
						_ = _hx_tmp_31
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(f, "meta"), ((any)(nil))))) {
							_hx_tmp_31 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f, "meta"), "length"), ((int32)(0))))))
						} else {
							_hx_tmp_31 = false
						}

						if _hx_tmp_31 {
							var _hx_tmp_32 any
							_ = _hx_tmp_32
							_hx_tmp_32 = fl1
							Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_32, "meta", HxMakeArray[any]().Dyn())
							{
								var tmp__g_1 int32
								_ = tmp__g_1
								tmp__g_1 = 0
								var _g1 HxArray[any]
								_ = _g1
								_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(f, "meta"))
								for {
									var _hx_tmp_33 int32
									_ = _hx_tmp_33
									_hx_tmp_33 = tmp__g_1
									if !(_hx_tmp_33 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
										break
									}

									var m any
									_ = m
									m = _g1.Get(tmp__g_1)
									tmp__g_1++
									var tmp_fl_1 HxArray[any]
									_ = tmp_fl_1
									tmp_fl_1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(fl1, "meta"))
									var m1 string
									_ = m1
									m1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(m, "name"))
									var _g2 HxArray[any]
									_ = _g2
									_g2 = HxMakeArray[any]().Dyn()
									{
										var _g3 int32
										_ = _g3
										_g3 = 0
										var _g4 HxArray[string]
										_ = _g4
										_g4 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(m, "params"))
										for {
											var _hx_tmp_34 int32
											_ = _hx_tmp_34
											_hx_tmp_34 = _g3
											if !(_hx_tmp_34 < _g4.Len()) {
												break
											}

											var p string
											_ = p
											p = _g4.Get(_g3)
											_g3++
											var tmp_initField_1 any
											_ = tmp_initField_1
											{
												var hx_try_state int
												_ = hx_try_state

												func() {
													defer func() {
														if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
															if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_lorscript_error").(*Hx_Obj_loreline_lorscript_error); _hx_ok {
																_ = e
																{
																	tmp_initField_1 = nil
																}
															} else {
																panic(_hx_ex)
															}
														}
													}()
													{
														var _hx_tmp_35 string
														_ = _hx_tmp_35
														_hx_tmp_35 = p
														var _hx_tmp_36 any
														_ = _hx_tmp_36
														_hx_tmp_36 = nil
														tmp_initField_1 = _gthis.Hx_Field_parser.VTable.Hx_Field_parseString(_hx_tmp_35, _hx_tmp_36, nil)
													}
												}()
												if hx_try_state == 1 {
													return
												}
												if hx_try_state == 2 {
													break
												}
												if hx_try_state == 3 {
													continue
												}
											}

											_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g2), tmp_initField_1))
										}
									}

									var _hx_tmp_37 HxArray[any]
									_ = _hx_tmp_37
									_hx_tmp_37 = HxMakeArrayView[any](tmp_fl_1)
									_ = (Hx_Array_Push[any](_hx_tmp_37, any(map[string]any{"name": ((any)(m1)), "params": ((any)(_g2))})))
								}
							}
						}

						for {
							if !(pkeys.Len() > 0) {
								break
							}

							func() func(_hx_param_key string) bool {
								var _hx_closure_src_393 func(_hx_param_key any) bool
								_ = _hx_closure_src_393
								_hx_closure_src_393 = _gthis.Hx_Field_localParams.VTable.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_remove
								return func(_hx_param_key string) bool {
									return _hx_closure_src_393(((any)(_hx_param_key)))
								}
							}()((Hx_Array_Pop[string](pkeys)).Value)
						}

						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(fl1, "name"), ((any)("new"))))) {
							Hx_Field_go_haxe_hxdynamic_setField(cl, "constructor", fl1)
						} else {
							var _hx_tmp_38 any
							_ = _hx_tmp_38
							_hx_tmp_38 = Hx_Field_go_haxe_hxdynamic_getField(fields, "set")
							Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_38, HxMakeArray[any](Hx_Field_go_haxe_hxdynamic_getField(f, "name"), fl1).Dyn())
						}
					}

					{
						var _g3 int32
						_ = _g3
						_g3 = 0
						var _g4 HxArray[any]
						_ = _g4
						_g4 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"))
						for {
							var _hx_tmp_39 int32
							_ = _hx_tmp_39
							_hx_tmp_39 = _g3
							if !(_hx_tmp_39 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g4)) {
								break
							}

							var f any
							_ = f
							f = _g4.Get(_g3)
							_g3++
							var _hx_tmp_40 any
							_ = _hx_tmp_40
							_hx_tmp_40 = f
							initField(_hx_tmp_40, Hx_Field_go_haxe_hxdynamic_getField(cl, "fields"))
						}
					}

					{
						var _g5 int32
						_ = _g5
						_g5 = 0
						var _g6 HxArray[any]
						_ = _g6
						_g6 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "statics"))
						for {
							var _hx_tmp_41 int32
							_ = _hx_tmp_41
							_hx_tmp_41 = _g5
							if !(_hx_tmp_41 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g6)) {
								break
							}

							var f any
							_ = f
							f = _g6.Get(_g5)
							_g5++
							var _hx_tmp_42 any
							_ = _hx_tmp_42
							_hx_tmp_42 = f
							initField(_hx_tmp_42, Hx_Field_go_haxe_hxdynamic_getField(cl, "statics"))
						}
					}

					_gthis.Hx_Field_localParams = nil
				}))
				var _hx_tmp_43 string
				_ = _hx_tmp_43
				_hx_tmp_43 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(cl, "name"))
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
					var _hx_closure_src_394 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_394
					_hx_closure_src_394 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
						_hx_closure_src_394(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_43, Hx_Enum_loreline_lorscript_ctypedecl_CTClass{cl})
			}
		}

	case 2:
		{
			var e any
			_ = e
			e = x.(Hx_Enum_haxe_rtti_typetree_TEnumdecl).Hx_Field_e
			{
				if func() func(_hx_param_key string) bool {
					var _hx_closure_src_395 func(_hx_param_key any) bool
					_ = _hx_closure_src_395
					_hx_closure_src_395 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_395(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "path"))) {
					return
				}

				var _hx_tmp_44 any
				_ = _hx_tmp_44
				_hx_tmp_44 = Hx_Field_go_haxe_hxdynamic_getField(e, "path")
				var _hx_tmp_45 any
				_ = _hx_tmp_45
				_hx_tmp_45 = ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()))
				var en any
				_ = en
				en = any(map[string]any{"name": _hx_tmp_44, "params": _hx_tmp_45, "constructors": ((any)(HxMakeArray[any]().Dyn()))})
				this.VTable.Hx_Field_addMeta(e, en)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 Hx_Typedef_haxe_rtti_typeparams
					_ = _g1
					_g1 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(e, "params"))
					for {
						var _hx_tmp_46 int32
						_ = _hx_tmp_46
						_hx_tmp_46 = _g
						if !(_hx_tmp_46 < _g1.Len()) {
							break
						}

						var p string
						_ = p
						p = _g1.Get(_g)
						_g++
						var _hx_tmp_47 any
						_ = _hx_tmp_47
						_hx_tmp_47 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(en, "params"), "push")
						Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_47, HxMakeArray[any](Hx_Enum_loreline_lorscript_ttype_TParam{p}).Dyn())
					}
				}

				var _hx_tmp_48 HxArray[func()]
				_ = _hx_tmp_48
				_hx_tmp_48 = todo
				_ = (Hx_Array_Push[func()](_hx_tmp_48, func() {
					var _g *Hx_Obj_haxe_imap
					_ = _g
					_g = (func() *Hx_Obj_haxe_imap {
						_hx_icast_72 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
						if _hx_icast_72 == nil {
							return nil
						}
						Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_72, "Hx_Obj_haxe_imap")
						return &Hx_Obj_haxe_imap{VTable: _hx_icast_72.VTable.(Hx_Obj_VTable_haxe_imap)}
					})()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _g2
						_g2 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(en, "params"))
						for {
							var _hx_tmp_49 int32
							_ = _hx_tmp_49
							_hx_tmp_49 = _g1
							if !(_hx_tmp_49 < _g2.Len()) {
								break
							}

							var t Hx_Enum_loreline_lorscript_ttype
							_ = t
							t = _g2.Get(_g1)
							_g1++
							var _hx_tmp_50 any
							_ = _hx_tmp_50
							_hx_tmp_50 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(e, "path"), ((any)(".")))
							func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								var _hx_closure_src_396 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_396
								_hx_closure_src_396 = _g.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
									_hx_closure_src_396(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_50, ((any)(Hx_Field_loreline_lorscript_checker_typeStr(t))))), t)
						}
					}

					_gthis.Hx_Field_localParams = _g
					{
						var _g3 int32
						_ = _g3
						_g3 = 0
						var _g4 HxArray[any]
						_ = _g4
						_g4 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(e, "constructors"))
						for {
							var _hx_tmp_51 int32
							_ = _hx_tmp_51
							_hx_tmp_51 = _g3
							if !(_hx_tmp_51 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g4)) {
								break
							}

							var c any
							_ = c
							c = _g4.Get(_g3)
							_g3++
							var en1 HxArray[any]
							_ = en1
							en1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(en, "constructors"))
							var c1 string
							_ = c1
							c1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(c, "name"))
							var tmp HxArray[any]
							_ = tmp
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c, "args"), ((any)(nil))))) {
								tmp = nil
							} else {
								var tmp__g_1 HxArray[any]
								_ = tmp__g_1
								tmp__g_1 = HxMakeArray[any]().Dyn()
								{
									var _g1 int32
									_ = _g1
									_g1 = 0
									var _g2 HxArray[any]
									_ = _g2
									_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "args"))
									for {
										var _hx_tmp_52 int32
										_ = _hx_tmp_52
										_hx_tmp_52 = _g1
										if !(_hx_tmp_52 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
											break
										}

										var a any
										_ = a
										a = _g2.Get(_g1)
										_g1++
										var _hx_tmp_53 HxArray[any]
										_ = _hx_tmp_53
										_hx_tmp_53 = HxMakeArrayView[any](tmp__g_1)
										var _hx_tmp_54 any
										_ = _hx_tmp_54
										_hx_tmp_54 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
										var _hx_tmp_55 any
										_ = _hx_tmp_55
										_hx_tmp_55 = Hx_Field_go_haxe_hxdynamic_getField(a, "opt")
										_ = (Hx_Array_Push[any](_hx_tmp_53, any(map[string]any{"name": _hx_tmp_54, "opt": _hx_tmp_55, "t": ((any)(_gthis.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
											_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
											if _hx_d == nil {
												var _hx_z Hx_Enum_haxe_rtti_ctype
												return _hx_z
											}
											return _hx_d.(Hx_Enum_haxe_rtti_ctype)
										}())))})))
									}
								}

								tmp = tmp__g_1
							}

							var _hx_tmp_56 HxArray[any]
							_ = _hx_tmp_56
							_hx_tmp_56 = HxMakeArrayView[any](en1)
							_ = (Hx_Array_Push[any](_hx_tmp_56, any(map[string]any{"name": ((any)(c1)), "args": ((any)(tmp))})))
						}
					}

					_gthis.Hx_Field_localParams = nil
				}))
				var _hx_tmp_57 string
				_ = _hx_tmp_57
				_hx_tmp_57 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(en, "name"))
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
					var _hx_closure_src_397 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_397
					_hx_closure_src_397 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
						_hx_closure_src_397(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_57, Hx_Enum_loreline_lorscript_ctypedecl_CTEnum{en})
			}
		}

	case 3:
		{
			var t any
			_ = t
			t = x.(Hx_Enum_haxe_rtti_typetree_TTypedecl).Hx_Field_t
			{
				if func() func(_hx_param_key string) bool {
					var _hx_closure_src_398 func(_hx_param_key any) bool
					_ = _hx_closure_src_398
					_hx_closure_src_398 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_398(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(t, "path"))) {
					return
				}

				var _hx_tmp_58 any
				_ = _hx_tmp_58
				_hx_tmp_58 = Hx_Field_go_haxe_hxdynamic_getField(t, "path")
				var td any
				_ = td
				td = any(map[string]any{"name": _hx_tmp_58, "params": ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]())), "t": ((any)(nil))})
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 Hx_Typedef_haxe_rtti_typeparams
					_ = _g1
					_g1 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(t, "params"))
					for {
						var _hx_tmp_59 int32
						_ = _hx_tmp_59
						_hx_tmp_59 = _g
						if !(_hx_tmp_59 < _g1.Len()) {
							break
						}

						var p string
						_ = p
						p = _g1.Get(_g)
						_g++
						var _hx_tmp_60 any
						_ = _hx_tmp_60
						_hx_tmp_60 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(td, "params"), "push")
						Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_60, HxMakeArray[any](Hx_Enum_loreline_lorscript_ttype_TParam{p}).Dyn())
					}
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t, "path"), ((any)("loreline.lorscript.TypeCheck"))))) {
					var _hx_tmp_61 any
					_ = _hx_tmp_61
					_hx_tmp_61 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(td, "params"), "reverse")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_61, HxMakeArray[any]().Dyn())
				}

				var _hx_tmp_62 HxArray[func()]
				_ = _hx_tmp_62
				_hx_tmp_62 = todo
				_ = (Hx_Array_Push[func()](_hx_tmp_62, func() {
					var _g *Hx_Obj_haxe_imap
					_ = _g
					_g = (func() *Hx_Obj_haxe_imap {
						_hx_icast_73 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
						if _hx_icast_73 == nil {
							return nil
						}
						Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_73, "Hx_Obj_haxe_imap")
						return &Hx_Obj_haxe_imap{VTable: _hx_icast_73.VTable.(Hx_Obj_VTable_haxe_imap)}
					})()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _g2
						_g2 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(td, "params"))
						for {
							var _hx_tmp_63 int32
							_ = _hx_tmp_63
							_hx_tmp_63 = _g1
							if !(_hx_tmp_63 < _g2.Len()) {
								break
							}

							var pt Hx_Enum_loreline_lorscript_ttype
							_ = pt
							pt = _g2.Get(_g1)
							_g1++
							var _hx_tmp_64 any
							_ = _hx_tmp_64
							_hx_tmp_64 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(t, "path"), ((any)(".")))
							func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								var _hx_closure_src_399 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_399
								_hx_closure_src_399 = _g.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
									_hx_closure_src_399(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_64, ((any)(Hx_Field_loreline_lorscript_checker_typeStr(pt))))), pt)
						}
					}

					_gthis.Hx_Field_localParams = _g
					var _hx_tmp_65 any
					_ = _hx_tmp_65
					_hx_tmp_65 = td
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_65, "t", _gthis.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t, "type"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_haxe_rtti_ctype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_haxe_rtti_ctype)
					}()))
					_gthis.Hx_Field_localParams = nil
				}))
				var _hx_tmp_66 string
				_ = _hx_tmp_66
				_hx_tmp_66 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(t, "path"))
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
					var _hx_closure_src_400 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_400
					_hx_closure_src_400 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
						_hx_closure_src_400(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_66, Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef{td})
			}
		}

	case 4:
		{
			var a any
			_ = a
			a = x.(Hx_Enum_haxe_rtti_typetree_TAbstractdecl).Hx_Field_a
			{
				if func() func(_hx_param_key string) bool {
					var _hx_closure_src_401 func(_hx_param_key any) bool
					_ = _hx_closure_src_401
					_hx_closure_src_401 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_401(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "path"))) {
					return
				}

				var _hx_tmp_67 any
				_ = _hx_tmp_67
				_hx_tmp_67 = Hx_Field_go_haxe_hxdynamic_getField(a, "path")
				var ta any
				_ = ta
				ta = any(map[string]any{"name": _hx_tmp_67, "params": ((any)(HxMakeArray[Hx_Enum_loreline_lorscript_ttype]())), "t": ((any)(nil))})
				this.VTable.Hx_Field_addMeta(a, ta)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 Hx_Typedef_haxe_rtti_typeparams
					_ = _g1
					_g1 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(a, "params"))
					for {
						var _hx_tmp_68 int32
						_ = _hx_tmp_68
						_hx_tmp_68 = _g
						if !(_hx_tmp_68 < _g1.Len()) {
							break
						}

						var p string
						_ = p
						p = _g1.Get(_g)
						_g++
						var _hx_tmp_69 any
						_ = _hx_tmp_69
						_hx_tmp_69 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(ta, "params"), "push")
						Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_69, HxMakeArray[any](Hx_Enum_loreline_lorscript_ttype_TParam{p}).Dyn())
					}
				}

				var _hx_tmp_70 HxArray[func()]
				_ = _hx_tmp_70
				_hx_tmp_70 = todo
				_ = (Hx_Array_Push[func()](_hx_tmp_70, func() {
					var _g *Hx_Obj_haxe_imap
					_ = _g
					_g = (func() *Hx_Obj_haxe_imap {
						_hx_icast_74 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
						if _hx_icast_74 == nil {
							return nil
						}
						Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_74, "Hx_Obj_haxe_imap")
						return &Hx_Obj_haxe_imap{VTable: _hx_icast_74.VTable.(Hx_Obj_VTable_haxe_imap)}
					})()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _g2
						_g2 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(ta, "params"))
						for {
							var _hx_tmp_71 int32
							_ = _hx_tmp_71
							_hx_tmp_71 = _g1
							if !(_hx_tmp_71 < _g2.Len()) {
								break
							}

							var t Hx_Enum_loreline_lorscript_ttype
							_ = t
							t = _g2.Get(_g1)
							_g1++
							var _hx_tmp_72 any
							_ = _hx_tmp_72
							_hx_tmp_72 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(a, "path"), ((any)(".")))
							func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								var _hx_closure_src_402 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_402
								_hx_closure_src_402 = _g.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
									_hx_closure_src_402(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_72, ((any)(Hx_Field_loreline_lorscript_checker_typeStr(t))))), t)
						}
					}

					_gthis.Hx_Field_localParams = _g
					var _hx_tmp_73 any
					_ = _hx_tmp_73
					_hx_tmp_73 = ta
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_73, "t", _gthis.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "athis"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_haxe_rtti_ctype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_haxe_rtti_ctype)
					}()))
					_gthis.Hx_Field_localParams = nil
				}))
				var _hx_tmp_74 string
				_ = _hx_tmp_74
				_hx_tmp_74 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "path"))
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
					var _hx_closure_src_403 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_403
					_hx_closure_src_403 = this.Hx_Field_types.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ctypedecl) {
						_hx_closure_src_403(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_74, Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract{ta})
			}
		}

	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_addMeta(src any, to any) {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(src, "meta"), ((any)(nil))))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(src, "meta"), "length"), ((int32)(0))))))
	}

	if _hx_tmp_0 {
		return
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = to
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "meta", HxMakeArray[any]().Dyn())
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(src, "meta"))
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var m any
			_ = m
			m = _g1.Get(_g)
			_g++
			var to1 HxArray[any]
			_ = to1
			to1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(to, "meta"))
			var m1 string
			_ = m1
			m1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(m, "name"))
			var _g2 HxArray[any]
			_ = _g2
			_g2 = HxMakeArray[any]().Dyn()
			{
				var _g3 int32
				_ = _g3
				_g3 = 0
				var _g4 HxArray[string]
				_ = _g4
				_g4 = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(m, "params"))
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g3
					if !(_hx_tmp_3 < _g4.Len()) {
						break
					}

					var p string
					_ = p
					p = _g4.Get(_g3)
					_g3++
					var tmp any
					_ = tmp
					{
						var hx_try_state int
						_ = hx_try_state

						func() {
							defer func() {
								if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
									if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_lorscript_error").(*Hx_Obj_loreline_lorscript_error); _hx_ok {
										_ = e
										{
											tmp = nil
										}
									} else {
										panic(_hx_ex)
									}
								}
							}()
							{
								var _hx_tmp_4 string
								_ = _hx_tmp_4
								_hx_tmp_4 = p
								var _hx_tmp_5 any
								_ = _hx_tmp_5
								_hx_tmp_5 = nil
								tmp = this.Hx_Field_parser.VTable.Hx_Field_parseString(_hx_tmp_4, _hx_tmp_5, nil)
							}
						}()
						if hx_try_state == 1 {
							return
						}
						if hx_try_state == 2 {
							break
						}
						if hx_try_state == 3 {
							continue
						}
					}

					_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g2), tmp))
				}
			}

			var _hx_tmp_6 HxArray[any]
			_ = _hx_tmp_6
			_hx_tmp_6 = HxMakeArrayView[any](to1)
			_ = (Hx_Array_Push[any](_hx_tmp_6, any(map[string]any{"name": ((any)(m1)), "params": ((any)(_g2))})))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_makeXmlType(t Hx_Enum_haxe_rtti_ctype) Hx_Enum_loreline_lorscript_ttype {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			return Hx_Enum_loreline_lorscript_ttype_TUnresolved{"Unknown"}
		}

	case 1:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_params
			var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _g
			_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = _g1
					if !(_hx_tmp_0 < params.Len()) {
						break
					}

					var t Hx_Enum_haxe_rtti_ctype
					_ = t
					t = params.Get(_g1)
					_g1++
					var _hx_tmp_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = _hx_tmp_1
					_hx_tmp_1 = _g
					_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_1, this.VTable.Hx_Field_makeXmlType(t)))
				}
			}

			return this.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(name), _g)
		}

	case 2:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_params
			var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _g
			_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g1
					if !(_hx_tmp_2 < params.Len()) {
						break
					}

					var t Hx_Enum_haxe_rtti_ctype
					_ = t
					t = params.Get(_g1)
					_g1++
					var _hx_tmp_3 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_3, this.VTable.Hx_Field_makeXmlType(t)))
				}
			}

			return this.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(name), _g)
		}

	case 3:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_params
			var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _g
			_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g1
					if !(_hx_tmp_4 < params.Len()) {
						break
					}

					var t Hx_Enum_haxe_rtti_ctype
					_ = t
					t = params.Get(_g1)
					_g1++
					var _hx_tmp_5 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = _hx_tmp_5
					_hx_tmp_5 = _g
					_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_5, this.VTable.Hx_Field_makeXmlType(t)))
				}
			}

			return this.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(name), _g)
		}

	case 4:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_args
			var ret Hx_Enum_haxe_rtti_ctype
			_ = ret
			ret = t.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_ret
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _g1
					if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
						break
					}

					var a any
					_ = a
					a = args.Get(_g1)
					_g1++
					var _hx_tmp_7 HxArray[any]
					_ = _hx_tmp_7
					_hx_tmp_7 = HxMakeArrayView[any](_g)
					var _hx_tmp_8 any
					_ = _hx_tmp_8
					_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
					var _hx_tmp_9 any
					_ = _hx_tmp_9
					_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(a, "opt")
					_ = (Hx_Array_Push[any](_hx_tmp_7, any(map[string]any{"name": _hx_tmp_8, "opt": _hx_tmp_9, "t": ((any)(this.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_haxe_rtti_ctype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_haxe_rtti_ctype)
					}())))})))
				}
			}

			var _hx_tmp_10 HxArray[any]
			_ = _hx_tmp_10
			_hx_tmp_10 = _g
			return Hx_Enum_loreline_lorscript_ttype_TFun{_hx_tmp_10, this.VTable.Hx_Field_makeXmlType(ret)}
		}

	case 5:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_haxe_rtti_ctype_CAnonymous).Hx_Field_fields
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					_hx_tmp_11 = _g1
					if !(_hx_tmp_11 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
						break
					}

					var f any
					_ = f
					f = fields.Get(_g1)
					_g1++
					var f1 string
					_ = f1
					f1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name"))
					var tmp Hx_Enum_loreline_lorscript_ttype
					_ = tmp
					tmp = this.VTable.Hx_Field_makeXmlType(func() Hx_Enum_haxe_rtti_ctype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "type"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_haxe_rtti_ctype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_haxe_rtti_ctype)
					}())
					var m HxArray[any]
					_ = m
					m = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(f, "meta"))
					var tmp1 bool
					_ = tmp1
					if m == nil {
						tmp1 = false
					} else {
						var b bool
						_ = b
						b = false
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							for {
								var _hx_tmp_12 int32
								_ = _hx_tmp_12
								_hx_tmp_12 = tmp__g_1
								if !(_hx_tmp_12 < Hx_Field_go_haxe_hxdynamic_getArrayLength(m)) {
									break
								}

								var m1 any
								_ = m1
								m1 = m.Get(tmp__g_1)
								tmp__g_1++
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(m1, "name"), ((any)(":optional"))))) {
									b = true
									break
								}
							}
						}

						tmp1 = b
					}

					var _hx_tmp_13 HxArray[any]
					_ = _hx_tmp_13
					_hx_tmp_13 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_13, any(map[string]any{"name": ((any)(f1)), "t": ((any)(tmp)), "opt": ((any)(tmp1))})))
				}
			}

			return Hx_Enum_loreline_lorscript_ttype_TAnon{_g}
		}

	case 6:
		{
			var t1 Hx_Enum_haxe_rtti_ctype
			_ = t1
			t1 = t.(Hx_Enum_haxe_rtti_ctype_CDynamic).Hx_Field_t
			return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
		}

	case 7:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_params
			var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _g
			_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_14 int32
					_ = _hx_tmp_14
					_hx_tmp_14 = _g1
					if !(_hx_tmp_14 < params.Len()) {
						break
					}

					var t Hx_Enum_haxe_rtti_ctype
					_ = t
					t = params.Get(_g1)
					_g1++
					var _hx_tmp_15 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = _hx_tmp_15
					_hx_tmp_15 = _g
					_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_15, this.VTable.Hx_Field_makeXmlType(t)))
				}
			}

			return this.VTable.Hx_Field_getType(Hx_Field_go_haxe_hxdynamic_toString(name), _g)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_getType(name string, args__isOptional any) Hx_Enum_loreline_lorscript_ttype {
	args := HxDefault[HxArray[Hx_Enum_loreline_lorscript_ttype]]()
	_ = args

	if args__isOptional != nil {
		args = HxConvert[HxArray[Hx_Enum_loreline_lorscript_ttype]](args__isOptional)
	} else {
		args = nil
	}
	{
		if this.Hx_Field_localParams != nil {
			var t Hx_Enum_loreline_lorscript_ttype
			_ = t
			t = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
				var _hx_closure_src_404 func(_hx_param_key any) any
				_ = _hx_closure_src_404
				_hx_closure_src_404 = this.Hx_Field_localParams.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
					return func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_404(((any)(_hx_param_key))))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
				}
			}()(name)
			if t != nil {
				return t
			}
		}

		var t Hx_Enum_loreline_lorscript_ttype
		_ = t
		t = this.VTable.Hx_Field_resolve(name, args)
		if t == nil {
			var pack HxArray[string]
			_ = pack
			pack = (Hx_Field_go_haxe_hxstring_split(name, "."))
			if pack.Len() > 1 {
				var _hx_tmp_0 HxArray[string]
				_ = _hx_tmp_0
				_hx_tmp_0 = pack
				var priv string
				_ = priv
				priv = _hx_tmp_0.Get((pack.Len() - 2))
				var hx_cmp_10 HxNullable[int32]
				_ = hx_cmp_10
				hx_cmp_10 = (Hx_Field_go_haxe_hxstring_charCodeAt(priv, 0))
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if (hx_cmp_10).Valid == true {
					_hx_tmp_1 = ((hx_cmp_10).Value == 95)
				} else {
					_hx_tmp_1 = false
				}

				if _hx_tmp_1 {
					_ = (Hx_Array_Remove[string](pack, priv))
					return this.VTable.Hx_Field_getType((Hx_Array_Join[string](pack, ".")), args)
				}
			}

			return Hx_Enum_loreline_lorscript_ttype_TUnresolved{name}
		}

		return t
	}
}

func (this *Hx_Obj_loreline_lorscript_checkertypes) Hx_Field_resolve(name string, args__isOptional any) Hx_Enum_loreline_lorscript_ttype {
	args := HxDefault[HxArray[Hx_Enum_loreline_lorscript_ttype]]()
	_ = args

	if args__isOptional != nil {
		args = HxConvert[HxArray[Hx_Enum_loreline_lorscript_ttype]](args__isOptional)
	} else {
		args = nil
	}
	{
		if name == "Null" {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if args == nil {
				_hx_tmp_0 = true
			} else {
				_hx_tmp_0 = (args.Len() != 1)
			}

			if _hx_tmp_0 {
				panic(Hx_Field_haxe_exception_thrown(((any)("Missing Null<T> parameter"))))
			}

			return Hx_Enum_loreline_lorscript_ttype_TNull{args.Get(0)}
		}

		var t Hx_Enum_loreline_lorscript_ctypedecl
		_ = t
		t = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ctypedecl {
			var _hx_closure_src_405 func(_hx_param_key any) any
			_ = _hx_closure_src_405
			_hx_closure_src_405 = this.Hx_Field_types.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ctypedecl {
				return func() Hx_Enum_loreline_lorscript_ctypedecl {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_405(((any)(_hx_param_key))))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ctypedecl
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ctypedecl)
				}()
			}
		}()(name)
		if t == nil {
			return nil
		}

		if args == nil {
			args = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
		}

		switch t.Hx_Field_enumIndex() {
		case 0:
			{
				var c any
				_ = c
				c = t.(Hx_Enum_loreline_lorscript_ctypedecl_CTClass).Hx_Field_c
				return Hx_Enum_loreline_lorscript_ttype_TInst{c, args}
			}

		case 1:
			{
				var e any
				_ = e
				e = t.(Hx_Enum_loreline_lorscript_ctypedecl_CTEnum).Hx_Field_e
				return Hx_Enum_loreline_lorscript_ttype_TEnum{e, args}
			}

		case 2:
			{
				var t1 any
				_ = t1
				t1 = t.(Hx_Enum_loreline_lorscript_ctypedecl_CTTypedef).Hx_Field_t
				return Hx_Enum_loreline_lorscript_ttype_TType{t1, args}
			}

		case 3:
			{
				var t1 Hx_Enum_loreline_lorscript_ttype
				_ = t1
				t1 = t.(Hx_Enum_loreline_lorscript_ctypedecl_CTAlias).Hx_Field_t
				return t1
			}

		case 4:
			{
				var a any
				_ = a
				a = t.(Hx_Enum_loreline_lorscript_ctypedecl_CTAbstract).Hx_Field_a
				return Hx_Enum_loreline_lorscript_ttype_TAbstract{a, args}
			}

		default:
			{
				panic(Hx_Field_haxe_exception_thrown("unreachable"))
			}
		}
	}
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_checker() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_checker() {
}

type Hx_Obj_VTable_loreline_lorscript_checker interface {
	Hx_Field_setGlobals(cl any, allowPrivate__isOptional any)
	Hx_Field_removeGlobal(name string)
	Hx_Field_apply(t Hx_Enum_loreline_lorscript_ttype, params HxArray[Hx_Enum_loreline_lorscript_ttype], args HxArray[Hx_Enum_loreline_lorscript_ttype]) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_mapType(t Hx_Enum_loreline_lorscript_ttype, f func(_hx_carg0 Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_makeMono() Hx_Enum_loreline_lorscript_ttype
	Hx_Field_typeExprWith(expr any, t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_mk(e Hx_Enum_loreline_lorscript_exprdef, p any) any
	Hx_Field_onCompletion(expr any, t Hx_Enum_loreline_lorscript_ttype)
	Hx_Field_typeField(o any, f string, expr any, forWrite bool) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_setEvent(name string, _hx_reserved_type Hx_Enum_loreline_lorscript_ttype)
	Hx_Field_typeArgs(args HxArray[any], pos any) HxArray[any]
	Hx_Field_makeType(t Hx_Enum_loreline_lorscript_ctype, e any) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_getField(t Hx_Enum_loreline_lorscript_ttype, f string, e any, forWrite__isOptional any) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_edef(e any) Hx_Enum_loreline_lorscript_exprdef
	Hx_Field_tryUnify(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype) bool
	Hx_Field_unify(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype, e any)
	Hx_Field_unasync(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_getIteratorType(it any, itt Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_makeIterator(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_isString(t Hx_Enum_loreline_lorscript_ttype) bool
	Hx_Field_link(a Hx_Enum_loreline_lorscript_ttype, b Hx_Enum_loreline_lorscript_ttype, r any) bool
	Hx_Field_onTopDownEnum(en any, field string) bool
	Hx_Field_linkLoop(a Hx_Enum_loreline_lorscript_ttype, t Hx_Enum_loreline_lorscript_ttype) bool
	Hx_Field_error(msg string, curExpr any)
	Hx_Field_typeEq(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype) bool
	Hx_Field_getFields(t Hx_Enum_loreline_lorscript_ttype) HxArray[any]
	Hx_Field_typeExpr(expr any, withType Hx_Enum_loreline_lorscript__checker_withtype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_getGlobals() Hx_Typedef_map
	Hx_Field_setGlobal(name string, _hx_reserved_type Hx_Enum_loreline_lorscript_ttype)
	Hx_Field_check(expr any, withType__isOptional any, isCompletion__isOptional any) Hx_Enum_loreline_lorscript_ttype
	Hx_Field_saveLocals() *Hx_Obj_haxe_imap
	Hx_Field_follow(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_checker struct {
	VTable                      Hx_Obj_VTable_loreline_lorscript_checker `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_types              *Hx_Obj_loreline_lorscript_checkertypes
	Hx_Field_locals             Hx_Typedef_map
	Hx_Field_globals            Hx_Typedef_map
	Hx_Field_events             Hx_Typedef_map
	Hx_Field_currentFunType     Hx_Enum_loreline_lorscript_ttype
	Hx_Field_isCompletion       bool
	Hx_Field_allowDefine        bool
	Hx_Field_hasReturn          bool
	Hx_Field_allowAsync         bool
	Hx_Field_allowReturn        Hx_Enum_loreline_lorscript_ttype
	Hx_Field_allowGlobalsDefine bool
	Hx_Field_allowUntypedMeta   bool
	Hx_Field_onTopDownEnum_Dyn  func(this *Hx_Obj_loreline_lorscript_checker, en any, field string) bool
}

func Hx_Obj_loreline_lorscript_checker_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_checker {
	obj := &Hx_Obj_loreline_lorscript_checker{}
	obj.VTable = obj
	obj.Hx_Field_onTopDownEnum_Dyn = func(this *Hx_Obj_loreline_lorscript_checker, en any, field string) bool {
		return false
	}
	obj.Hx_Field_types = HxDefault[*Hx_Obj_loreline_lorscript_checkertypes]()
	obj.Hx_Field_locals = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_globals = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_events = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_currentFunType = HxDefault[Hx_Enum_loreline_lorscript_ttype]()
	obj.Hx_Field_isCompletion = HxDefault[bool]()
	obj.Hx_Field_allowDefine = HxDefault[bool]()
	obj.Hx_Field_hasReturn = HxDefault[bool]()
	obj.Hx_Field_allowAsync = HxDefault[bool]()
	obj.Hx_Field_allowReturn = HxDefault[Hx_Enum_loreline_lorscript_ttype]()
	obj.Hx_Field_allowGlobalsDefine = HxDefault[bool]()
	obj.Hx_Field_allowUntypedMeta = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_lorscript_checker_CreateInstance(types__isOptional any) *Hx_Obj_loreline_lorscript_checker {
	obj := Hx_Obj_loreline_lorscript_checker_CreateEmptyInstance()
	obj.Hx_New(types__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_New(types__isOptional any) {
	{
		types := HxDefault[*Hx_Obj_loreline_lorscript_checkertypes]()
		_ = types

		if types__isOptional != nil {
			types = HxConvert[*Hx_Obj_loreline_lorscript_checkertypes](types__isOptional)
		} else {
			types = nil
		}
		{
			this.Hx_Field_globals = (func() *Hx_Obj_haxe_imap {
				_hx_icast_76 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_76 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_76, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_76.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_events = (func() *Hx_Obj_haxe_imap {
				_hx_icast_77 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_77 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_77, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_77.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			if types == nil {
				types = Hx_Obj_loreline_lorscript_checkertypes_CreateInstance()
			}

			this.Hx_Field_types = types
		}

	}
}

var Hx_Obj_loreline_lorscript_checker_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Checker",
	HxMakeArray[string]("typeStr", "typeIter"),
	HxMakeArray[string]("types", "locals", "globals", "events", "currentFunType", "isCompletion", "allowDefine", "hasReturn", "allowAsync", "allowReturn", "allowGlobalsDefine", "allowUntypedMeta", "setGlobals", "removeGlobal", "setGlobal", "setEvent", "getGlobals", "onTopDownEnum", "typeArgs", "check", "edef", "error", "saveLocals", "makeType", "linkLoop", "link", "typeEq", "tryUnify", "unify", "apply", "mapType", "follow", "getFields", "getField", "unasync", "typeExprWith", "makeMono", "makeIterator", "mk", "isString", "onCompletion", "typeField", "typeExpr", "getIteratorType"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_checker_CreateInstance(HxConvert[any](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lorscript_checker_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_checker_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "typeStr":
			return Hx_Field_loreline_lorscript_checker_typeStr
		case "typeIter":
			return Hx_Field_loreline_lorscript_checker_typeIter
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_checker_RTTI
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_As_Hx_Obj_loreline_lorscript_checker() *Hx_Obj_loreline_lorscript_checker {
	return this
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_setGlobals(cl any, allowPrivate__isOptional any) {
	allowPrivate := HxDefault[bool]()
	_ = allowPrivate

	if allowPrivate__isOptional != nil {
		allowPrivate = HxConvert[bool](allowPrivate__isOptional)
	} else {
		allowPrivate = false
	}
	{
		var _gthis *Hx_Obj_loreline_lorscript_checker
		_ = _gthis
		_gthis = this
		for {
			{
				var hx_dyncast_78 any
				_ = hx_dyncast_78
				hx_dyncast_78 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(cl, "fields"), "Hx_Obj_haxe_imap")
				var _hx_tmp_0 *Hx_Obj_haxe_imap
				_ = _hx_tmp_0
				if hx_dyncast_78 != nil {
					_hx_tmp_0 = hx_dyncast_78.(*Hx_Obj_haxe_imap)
				} else {
					_hx_tmp_0 = nil
				}

				var f any
				_ = f
				f = _hx_tmp_0.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_iterator()
				for {
					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(f, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(f, "next")
					var f1 any
					_ = f1
					f1 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
					var _hx_tmp_3 bool
					_ = _hx_tmp_3
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(f1, "isPublic"))) {
						_hx_tmp_3 = true
					} else {
						_hx_tmp_3 = allowPrivate
					}

					if _hx_tmp_3 {
						var _hx_tmp_4 Hx_Enum_loreline_lorscript_ttype
						_ = _hx_tmp_4
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f1, "params"), "length"), ((int32)(0))))) {
							_hx_tmp_4 = func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
						} else {
							_hx_tmp_4 = Hx_Enum_loreline_lorscript_ttype_TLazy{func() Hx_Enum_loreline_lorscript_ttype {
								var _gthis1 *Hx_Obj_loreline_lorscript_checker
								_ = _gthis1
								_gthis1 = _gthis
								var tmp_f_1 Hx_Enum_loreline_lorscript_ttype
								_ = tmp_f_1
								tmp_f_1 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								var f2 HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = f2
								f2 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(f1, "params"))
								var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
								_ = _g
								_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
								{
									var _g1 int32
									_ = _g1
									_g1 = 0
									var _g2 int32
									_ = _g2
									_g2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f1, "params"), "length"))))
									for _g1 < _g2 {
										var _hx_tmp_5 int32
										_ = _hx_tmp_5
										_hx_tmp_5 = _g1
										_g1 = (_g1 + 1)
										var i int32
										_ = i
										i = _hx_tmp_5
										var _hx_tmp_6 HxArray[Hx_Enum_loreline_lorscript_ttype]
										_ = _hx_tmp_6
										_hx_tmp_6 = _g
										var _hx_tmp_7 HxArray[Hx_Enum_loreline_lorscript_ttype]
										_ = _hx_tmp_7
										_hx_tmp_7 = _hx_tmp_6
										_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_7, _gthis.VTable.Hx_Field_makeMono()))
									}
								}

								return _gthis1.VTable.Hx_Field_apply(tmp_f_1, f2, _g)
							}}
						}

						var tmp Hx_Enum_loreline_lorscript_ttype
						_ = tmp
						tmp = _hx_tmp_4
						this.VTable.Hx_Field_setGlobal(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f1, "name")), tmp)
					}
				}
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(cl, "superClass"), ((any)(nil))))) {
				break
			}

			var _g Hx_Enum_loreline_lorscript_ttype
			_ = _g
			_g = func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cl, "superClass"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}()
			if _g == nil {
				panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
			} else {
				if _g.Hx_Field_enumIndex() == 9 {
					var c any
					_ = c
					c = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
					cl = c
				} else {
					panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_removeGlobal(name string) {
	func() func(_hx_param_key string) bool {
		var _hx_closure_src_411 func(_hx_param_key any) bool
		_ = _hx_closure_src_411
		_hx_closure_src_411 = this.Hx_Field_globals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_remove
		return func(_hx_param_key string) bool {
			return _hx_closure_src_411(((any)(_hx_param_key)))
		}
	}()(name)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_setGlobal(name string, _hx_reserved_type Hx_Enum_loreline_lorscript_ttype) {
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
		var _hx_closure_src_412 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_412
		_hx_closure_src_412 = this.Hx_Field_globals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
			_hx_closure_src_412(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(name, _hx_reserved_type)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_setEvent(name string, _hx_reserved_type Hx_Enum_loreline_lorscript_ttype) {
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
		var _hx_closure_src_413 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_413
		_hx_closure_src_413 = this.Hx_Field_events.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
			_hx_closure_src_413(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(name, _hx_reserved_type)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_getGlobals() Hx_Typedef_map {
	return this.Hx_Field_globals
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_onTopDownEnum(en any, field string) bool {
	return this.Hx_Field_onTopDownEnum_Dyn(this, en, field)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_typeArgs(args HxArray[any], pos any) HxArray[any] {
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 int32
		_ = _g2
		_g2 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args)
		for _g1 < _g2 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			_g1 = (_g1 + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var a any
			_ = a
			a = args.Get(i)
			var _hx_tmp_1 Hx_Enum_loreline_lorscript_ttype
			_ = _hx_tmp_1
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(a, "t"), ((any)(nil))))) {
				_hx_tmp_1 = this.VTable.Hx_Field_makeMono()
			} else {
				_hx_tmp_1 = this.VTable.Hx_Field_makeType(func() Hx_Enum_loreline_lorscript_ctype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ctype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
				}(), pos)
			}

			var at Hx_Enum_loreline_lorscript_ttype
			_ = at
			at = _hx_tmp_1
			var _hx_tmp_2 HxArray[any]
			_ = _hx_tmp_2
			_hx_tmp_2 = HxMakeArrayView[any](_g)
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
			_ = (Hx_Array_Push[any](_hx_tmp_2, any(map[string]any{"name": _hx_tmp_3, "opt": Hx_Field_go_haxe_hxdynamic_getField(a, "opt"), "t": ((any)(at))})))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_check(expr any, withType__isOptional any, isCompletion__isOptional any) Hx_Enum_loreline_lorscript_ttype {
	withType := HxDefault[Hx_Enum_loreline_lorscript__checker_withtype]()
	_ = withType

	if withType__isOptional != nil {
		withType = HxConvert[Hx_Enum_loreline_lorscript__checker_withtype](withType__isOptional)
	} else {
		withType = nil
	}
	isCompletion := HxDefault[HxNullable[bool]]()
	_ = isCompletion

	if isCompletion__isOptional != nil {
		isCompletion = HxConvert[HxNullable[bool]](isCompletion__isOptional)
	} else {

	}
	{
		if (isCompletion).Valid == false {
			isCompletion = HxNullable[bool]{Value: false, Valid: true}
		}

		if withType == nil {
			withType = Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}
		}

		this.Hx_Field_locals = (func() *Hx_Obj_haxe_imap {
			_hx_icast_79 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_79 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_79, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_79.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		if this.Hx_Field_types.Hx_Field_t_string == nil {
			this.Hx_Field_types.Hx_Field_t_string = this.Hx_Field_types.VTable.Hx_Field_getType("String", nil)
		}

		this.Hx_Field_allowDefine = this.Hx_Field_allowGlobalsDefine
		this.Hx_Field_isCompletion = (isCompletion).Value
		if Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(expr, "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() == 14 {
			expr = this.VTable.Hx_Field_mk(Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](expr).Dyn()}, expr)
		}

		var _gthis *Hx_Obj_loreline_lorscript_checker
		_ = _gthis
		_gthis = this
		{
			var _g Hx_Enum_loreline_lorscript_exprdef
			_ = _g
			_g = func() Hx_Enum_loreline_lorscript_exprdef {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(expr, "e"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_exprdef
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
			}()
			if _g.Hx_Field_enumIndex() == 4 {
				var el HxArray[any]
				_ = el
				el = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
				var delayed HxArray[func() Hx_Enum_loreline_lorscript_ttype]
				_ = delayed
				delayed = HxMakeArray[func() Hx_Enum_loreline_lorscript_ttype]()
				var last Hx_Enum_loreline_lorscript_ttype
				_ = last
				last = Hx_Enum_loreline_lorscript_ttype_TVoid{}

				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
							break
						}

						var e any
						_ = e
						e = el.Get(_g1)
						_g1++
						for {
							var tmp__g_1 Hx_Enum_loreline_lorscript_exprdef
							_ = tmp__g_1
							tmp__g_1 = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if tmp__g_1.Hx_Field_enumIndex() == 25 {
								var e2 any
								_ = e2
								e2 = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
								e = e2
							} else {
								break
							}
						}

						{
							var tmp__g_1 Hx_Enum_loreline_lorscript_exprdef
							_ = tmp__g_1
							tmp__g_1 = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if tmp__g_1.Hx_Field_enumIndex() == 14 {
								var args HxArray[any]
								_ = args
								args = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_args
								var name HxNullable[string]
								_ = name
								name = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_name
								var ret Hx_Enum_loreline_lorscript_ctype
								_ = ret
								ret = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_ret
								if (name).Valid != false {
									var _hx_tmp_1 Hx_Enum_loreline_lorscript_ttype
									_ = _hx_tmp_1
									if ret == nil {
										_hx_tmp_1 = this.VTable.Hx_Field_makeMono()
									} else {
										_hx_tmp_1 = this.VTable.Hx_Field_makeType(ret, e)
									}

									var tret Hx_Enum_loreline_lorscript_ttype
									_ = tret
									tret = _hx_tmp_1
									var ft Hx_Enum_loreline_lorscript_ttype
									_ = ft
									ft = Hx_Enum_loreline_lorscript_ttype_TFun{this.VTable.Hx_Field_typeArgs(args, e), tret}

									func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
										var _hx_closure_src_414 func(_hx_param_key any, _hx_param_value any)
										_ = _hx_closure_src_414
										_hx_closure_src_414 = this.Hx_Field_locals.VTable.(interface {
											Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
										}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
										return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
											_hx_closure_src_414(((any)(_hx_param_key)), ((any)(_hx_param_value)))
										}
									}()((name).Value, ft)
									var _hx_tmp_2 HxArray[func() Hx_Enum_loreline_lorscript_ttype]
									_ = _hx_tmp_2
									_hx_tmp_2 = delayed
									_ = (Hx_Array_Push[func() Hx_Enum_loreline_lorscript_ttype](_hx_tmp_2, func() Hx_Enum_loreline_lorscript_ttype {
										_gthis.Hx_Field_currentFunType = ft
										_gthis.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
										return ft
									}))
								} else {
									{
										var tmp__g_2 int32
										_ = tmp__g_2
										tmp__g_2 = 0
										for {
											var _hx_tmp_3 int32
											_ = _hx_tmp_3
											_hx_tmp_3 = tmp__g_2
											if !(_hx_tmp_3 < delayed.Len()) {
												break
											}

											var f func() Hx_Enum_loreline_lorscript_ttype
											_ = f
											f = delayed.Get(tmp__g_2)
											tmp__g_2++
											f()
										}
									}

									delayed = HxMakeArray[func() Hx_Enum_loreline_lorscript_ttype]()
									var _hx_tmp_4 HxArray[any]
									_ = _hx_tmp_4
									_hx_tmp_4 = el
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_4.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(el) - 1)), e))) {
										last = this.VTable.Hx_Field_typeExpr(e, withType)
									} else {
										this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
									}
								}
							} else {
								{
									var tmp__g_2 int32
									_ = tmp__g_2
									tmp__g_2 = 0
									for {
										var _hx_tmp_5 int32
										_ = _hx_tmp_5
										_hx_tmp_5 = tmp__g_2
										if !(_hx_tmp_5 < delayed.Len()) {
											break
										}

										var f func() Hx_Enum_loreline_lorscript_ttype
										_ = f
										f = delayed.Get(tmp__g_2)
										tmp__g_2++
										f()
									}
								}

								delayed = HxMakeArray[func() Hx_Enum_loreline_lorscript_ttype]()
								var _hx_tmp_6 HxArray[any]
								_ = _hx_tmp_6
								_hx_tmp_6 = el
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_6.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(el) - 1)), e))) {
									last = this.VTable.Hx_Field_typeExpr(e, withType)
								} else {
									this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
								}
							}
						}
					}
				}

				{
					var _g2 int32
					_ = _g2
					_g2 = 0
					for {
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = _g2
						if !(_hx_tmp_7 < delayed.Len()) {
							break
						}

						var f func() Hx_Enum_loreline_lorscript_ttype
						_ = f
						f = delayed.Get(_g2)
						_g2++
						last = f()
					}
				}

				return last
			}
		}

		return this.VTable.Hx_Field_typeExpr(expr, withType)
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_edef(e any) Hx_Enum_loreline_lorscript_exprdef {
	return func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_error(msg string, curExpr any) {
	var e Hx_Enum_loreline_lorscript_errordef
	_ = e
	e = Hx_Enum_loreline_lorscript_errordef_ECustom{msg}

	var e1 *Hx_Obj_loreline_lorscript_error
	_ = e1
	e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "line")))))
	if !this.Hx_Field_isCompletion {
		panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_saveLocals() *Hx_Obj_haxe_imap {
	var _g *Hx_Obj_haxe_imap
	_ = _g
	_g = (func() *Hx_Obj_haxe_imap {
		_hx_icast_80 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_80 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_80, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_80.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	{
		var k any
		_ = k
		k = this.Hx_Field_locals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(k, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(k, "next")
			var k1 string
			_ = k1
			k1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = k1
			func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
				var _hx_closure_src_415 func(_hx_param_key any, _hx_param_value any)
				_ = _hx_closure_src_415
				_hx_closure_src_415 = _g.VTable.Hx_Field_set
				return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
					_hx_closure_src_415(((any)(_hx_param_key)), ((any)(_hx_param_value)))
				}
			}()(_hx_tmp_2, func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
				var _hx_closure_src_416 func(_hx_param_key any) any
				_ = _hx_closure_src_416
				_hx_closure_src_416 = this.Hx_Field_locals.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
					return func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_416(((any)(_hx_param_key))))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
				}
			}()(k1))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_makeType(t Hx_Enum_loreline_lorscript_ctype, e any) Hx_Enum_loreline_lorscript_ttype {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var path HxArray[string]
			_ = path
			path = t.(Hx_Enum_loreline_lorscript_ctype_CTPath).Hx_Field_path
			var params HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = params
			params = t.(Hx_Enum_loreline_lorscript_ctype_CTPath).Hx_Field_params
			var ct *Hx_Obj_loreline_lorscript_checkertypes
			_ = ct
			ct = this.Hx_Field_types
			var ct1 string
			_ = ct1
			ct1 = (Hx_Array_Join[string](path, "."))
			var ct2 HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = ct2
			if params == nil {
				ct2 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
			} else {
				var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = _g
				_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < params.Len()) {
							break
						}

						var p Hx_Enum_loreline_lorscript_ctype
						_ = p
						p = params.Get(_g1)
						_g1++
						var _hx_tmp_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _hx_tmp_1
						_hx_tmp_1 = _g
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_1, this.VTable.Hx_Field_makeType(p, e)))
					}
				}

				ct2 = _g
			}

			var ct3 Hx_Enum_loreline_lorscript_ttype
			_ = ct3
			ct3 = ct.VTable.Hx_Field_resolve(ct1, ct2)
			if ct3 == nil {
				{
					var e1 Hx_Enum_loreline_lorscript_errordef
					_ = e1
					e1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Unknown type ", Hx_Field_std_string(path))}

					var e2 *Hx_Obj_loreline_lorscript_error
					_ = e2
					e2 = Hx_Obj_loreline_lorscript_error_CreateInstance(e1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e2))))
					}
				}

				ct3 = Hx_Enum_loreline_lorscript_ttype_TDynamic{}
			}

			return ct3
		}

	case 1:
		{
			var args HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ctype_CTFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ctype
			_ = ret
			ret = t.(Hx_Enum_loreline_lorscript_ctype_CTFun).Hx_Field_ret
			var i int32
			_ = i
			i = 0
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g1
					if !(_hx_tmp_2 < args.Len()) {
						break
					}

					var a Hx_Enum_loreline_lorscript_ctype
					_ = a
					a = args.Get(_g1)
					_g1++
					var _hx_tmp_3 HxArray[any]
					_ = _hx_tmp_3
					_hx_tmp_3 = HxMakeArrayView[any](_g)
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = i
					i = (i + 1)
					_ = (Hx_Array_Push[any](_hx_tmp_3, any(map[string]any{"name": ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("p", Hx_Field_go_haxe_hxdynamic_toString(_hx_tmp_4)))), "opt": ((any)(false)), "t": ((any)(this.VTable.Hx_Field_makeType(a, e)))})))
				}
			}

			var _hx_tmp_5 HxArray[any]
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			return Hx_Enum_loreline_lorscript_ttype_TFun{_hx_tmp_5, this.VTable.Hx_Field_makeType(ret, e)}
		}

	case 2:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_loreline_lorscript_ctype_CTAnon).Hx_Field_fields
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _g1
					if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
						break
					}

					var f any
					_ = f
					f = fields.Get(_g1)
					_g1++
					var _hx_tmp_7 HxArray[any]
					_ = _hx_tmp_7
					_hx_tmp_7 = HxMakeArrayView[any](_g)
					var _hx_tmp_8 any
					_ = _hx_tmp_8
					_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(f, "name")
					_ = (Hx_Array_Push[any](_hx_tmp_7, any(map[string]any{"name": _hx_tmp_8, "opt": ((any)(false)), "t": ((any)(this.VTable.Hx_Field_makeType(func() Hx_Enum_loreline_lorscript_ctype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ctype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
					}(), e)))})))
				}
			}

			return Hx_Enum_loreline_lorscript_ttype_TAnon{_g}
		}

	case 3:
		{
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTParent).Hx_Field_t
			return this.VTable.Hx_Field_makeType(t1, e)
		}

	case 4:
		{
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTOpt).Hx_Field_t
			return this.VTable.Hx_Field_makeType(t1, e)
		}

	case 5:
		{
			var n string
			_ = n
			n = t.(Hx_Enum_loreline_lorscript_ctype_CTNamed).Hx_Field_n
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTNamed).Hx_Field_t
			return this.VTable.Hx_Field_makeType(t1, e)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_linkLoop(a Hx_Enum_loreline_lorscript_ttype, t Hx_Enum_loreline_lorscript_ttype) bool {
	if t == a {
		return true
	}

	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				return false
			}

			var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
			_ = _hx_tmp_0
			_hx_tmp_0 = a
			return this.VTable.Hx_Field_linkLoop(_hx_tmp_0, func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}())
		}

	case 5:
		{
			if (t == Hx_Enum_loreline_lorscript_ttype_TDynamic{}) {
				return false
			}

			return this.VTable.Hx_Field_linkLoop(a, Hx_Enum_loreline_lorscript_ttype_TDynamic{})
		}

	case 9:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = _g
					if !(_hx_tmp_1 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					if this.VTable.Hx_Field_linkLoop(a, t) {
						return true
					}
				}
			}

			return false
		}

	case 10:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g
					if !(_hx_tmp_2 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					if this.VTable.Hx_Field_linkLoop(a, t) {
						return true
					}
				}
			}

			return false
		}

	case 11:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					if !(_hx_tmp_3 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					if this.VTable.Hx_Field_linkLoop(a, t) {
						return true
					}
				}
			}

			return false
		}

	case 12:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					if !(_hx_tmp_4 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					if this.VTable.Hx_Field_linkLoop(a, t) {
						return true
					}
				}
			}

			return false
		}

	case 13:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ttype
			_ = ret
			ret = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = _g
					if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
						break
					}

					var arg any
					_ = arg
					arg = args.Get(_g)
					_g++
					var _hx_tmp_6 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_6
					_hx_tmp_6 = a
					if this.VTable.Hx_Field_linkLoop(_hx_tmp_6, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()) {
						return true
					}
				}
			}

			return this.VTable.Hx_Field_linkLoop(a, ret)
		}

	case 14:
		{
			var fl HxArray[any]
			_ = fl
			fl = t.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _g
					if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
						break
					}

					var f any
					_ = f
					f = fl.Get(_g)
					_g++
					var _hx_tmp_8 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_8
					_hx_tmp_8 = a
					if this.VTable.Hx_Field_linkLoop(_hx_tmp_8, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()) {
						return true
					}
				}
			}

			return false
		}

	default:
		{
			return false
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_link(a Hx_Enum_loreline_lorscript_ttype, b Hx_Enum_loreline_lorscript_ttype, r any) bool {
	if this.VTable.Hx_Field_linkLoop(a, b) {
		return (this.VTable.Hx_Field_follow(b) == a)
	}

	if (b == Hx_Enum_loreline_lorscript_ttype_TDynamic{}) {
		return true
	}

	Hx_Field_go_haxe_hxdynamic_setField(r, "r", b)
	return true
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_typeEq(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype) bool {
	if t1 == t2 {
		return true
	}

	switch t1.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t1.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				if !this.VTable.Hx_Field_link(t1, t2, r) {
					return false
				}

				Hx_Field_go_haxe_hxdynamic_setField(r, "r", t2)
				return true
			}

			return this.VTable.Hx_Field_typeEq(func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}(), t2)
		}

	case 8:
		{
			var _gt Hx_Enum_loreline_lorscript_ttype
			_ = _gt
			_gt = t1.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_0
					_hx_tmp_0 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_0, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = _gt
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_1 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_1
					_hx_tmp_1 = t1
					var _hx_tmp_2 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_2
					_hx_tmp_2 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_1, this.VTable.Hx_Field_apply(_hx_tmp_2, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = _gt
					return this.VTable.Hx_Field_typeEq(t1, t2)
				}
			}
		}

	case 9:
		{
			var _gc any
			_ = _gc
			_gc = t1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_3 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_3
					_hx_tmp_3 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_3, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 9:
				{
					var cl2 any
					_ = cl2
					cl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var cl1 any
					_ = cl1
					cl1 = _gc
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(cl1, cl2))) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = pl1.Len()
							for _g < _g1 {
								var _hx_tmp_4 int32
								_ = _hx_tmp_4
								_hx_tmp_4 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_4
								if !this.VTable.Hx_Field_typeEq(pl1.Get(i), pl2.Get(i)) {
									return false
								}
							}
						}

						return true
					}
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_5 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_5
					_hx_tmp_5 = t1
					var _hx_tmp_6 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_6
					_hx_tmp_6 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_5, this.VTable.Hx_Field_apply(_hx_tmp_6, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
				}
			}
		}

	case 10:
		{
			var _ge any
			_ = _ge
			_ge = t1.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_e
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_7 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_7
					_hx_tmp_7 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_7, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 10:
				{
					var e2 any
					_ = e2
					e2 = t2.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_e
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var e1 any
					_ = e1
					e1 = _ge
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, e2))) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = pl1.Len()
							for _g < _g1 {
								var _hx_tmp_8 int32
								_ = _hx_tmp_8
								_hx_tmp_8 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_8
								if !this.VTable.Hx_Field_typeEq(pl1.Get(i), pl2.Get(i)) {
									return false
								}
							}
						}

						return true
					}
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_9 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_9
					_hx_tmp_9 = t1
					var _hx_tmp_10 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_10
					_hx_tmp_10 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_9, this.VTable.Hx_Field_apply(_hx_tmp_10, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
				}
			}
		}

	case 11:
		{
			var _gt any
			_ = _gt
			_gt = t1.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_11 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_11
					_hx_tmp_11 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_11, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_12 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_12
					_hx_tmp_12 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(this.VTable.Hx_Field_apply(_hx_tmp_12, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(t1, t21))) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = pl1.Len()
							for _g < _g1 {
								var _hx_tmp_13 int32
								_ = _hx_tmp_13
								_hx_tmp_13 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_13
								if !this.VTable.Hx_Field_typeEq(pl1.Get(i), pl2.Get(i)) {
									return false
								}
							}
						}

						return true
					} else {
						var tmp_pl1_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = tmp_pl1_1
						tmp_pl1_1 = _gargs
						var tmp_t1_1 any
						_ = tmp_t1_1
						tmp_t1_1 = _gt
						var _hx_tmp_14 Hx_Enum_loreline_lorscript_ttype
						_ = _hx_tmp_14
						_hx_tmp_14 = func() Hx_Enum_loreline_lorscript_ttype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(tmp_t1_1, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ttype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
						}()
						return this.VTable.Hx_Field_typeEq(this.VTable.Hx_Field_apply(_hx_tmp_14, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(tmp_t1_1, "params")), tmp_pl1_1), t2)
					}
				}

			default:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_15 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_15
					_hx_tmp_15 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(this.VTable.Hx_Field_apply(_hx_tmp_15, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}
			}
		}

	case 12:
		{
			var _ga any
			_ = _ga
			_ga = t1.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_16 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_16
					_hx_tmp_16 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_16, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_17 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_17
					_hx_tmp_17 = t1
					var _hx_tmp_18 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_18
					_hx_tmp_18 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_17, this.VTable.Hx_Field_apply(_hx_tmp_18, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 12:
				{
					var a2 any
					_ = a2
					a2 = t2.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var a1 any
					_ = a1
					a1 = _ga
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a1, a2))) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = pl1.Len()
							for _g < _g1 {
								var _hx_tmp_19 int32
								_ = _hx_tmp_19
								_hx_tmp_19 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_19
								if !this.VTable.Hx_Field_typeEq(pl1.Get(i), pl2.Get(i)) {
									return false
								}
							}
						}

						return true
					}
				}

			default:
				{
				}
			}
		}

	case 13:
		{
			var _gargs HxArray[any]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var _gret Hx_Enum_loreline_lorscript_ttype
			_ = _gret
			_gret = t1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_20 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_20
					_hx_tmp_20 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_20, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_21 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_21
					_hx_tmp_21 = t1
					var _hx_tmp_22 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_22
					_hx_tmp_22 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_21, this.VTable.Hx_Field_apply(_hx_tmp_22, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 13:
				{
					var args2 HxArray[any]
					_ = args2
					args2 = t2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
					var r2 Hx_Enum_loreline_lorscript_ttype
					_ = r2
					r2 = t2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
					var r1 Hx_Enum_loreline_lorscript_ttype
					_ = r1
					r1 = _gret
					var args1 HxArray[any]
					_ = args1
					args1 = _gargs
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args1)
					if _hx_tmp_23 == Hx_Field_go_haxe_hxdynamic_getArrayLength(args2) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args1)
							for _g < _g1 {
								var _hx_tmp_24 int32
								_ = _hx_tmp_24
								_hx_tmp_24 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_24
								var _hx_tmp_25 Hx_Enum_loreline_lorscript_ttype
								_ = _hx_tmp_25
								_hx_tmp_25 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(args1.Get(i), "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								if !this.VTable.Hx_Field_typeEq(_hx_tmp_25, func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(args2.Get(i), "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()) {
									return false
								}
							}
						}

						return this.VTable.Hx_Field_typeEq(r1, r2)
					}
				}

			default:
				{
				}
			}
		}

	case 14:
		{
			var _gfields HxArray[any]
			_ = _gfields
			_gfields = t1.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_26 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_26
					_hx_tmp_26 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_26, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_27 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_27
					_hx_tmp_27 = t1
					var _hx_tmp_28 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_28
					_hx_tmp_28 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_27, this.VTable.Hx_Field_apply(_hx_tmp_28, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 14:
				{
					var a2 HxArray[any]
					_ = a2
					a2 = t2.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
					var a1 HxArray[any]
					_ = a1
					a1 = _gfields
					var _hx_tmp_29 int32
					_ = _hx_tmp_29
					_hx_tmp_29 = Hx_Field_go_haxe_hxdynamic_getArrayLength(a1)
					if _hx_tmp_29 == Hx_Field_go_haxe_hxdynamic_getArrayLength(a2) {
						var m *Hx_Obj_haxe_imap
						_ = m
						m = (func() *Hx_Obj_haxe_imap {
							_hx_icast_81 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
							if _hx_icast_81 == nil {
								return nil
							}
							Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_81, "Hx_Obj_haxe_imap")
							return &Hx_Obj_haxe_imap{VTable: _hx_icast_81.VTable.(Hx_Obj_VTable_haxe_imap)}
						})()
						{
							var _g int32
							_ = _g
							_g = 0
							for {
								var _hx_tmp_30 int32
								_ = _hx_tmp_30
								_hx_tmp_30 = _g
								if !(_hx_tmp_30 < Hx_Field_go_haxe_hxdynamic_getArrayLength(a2)) {
									break
								}

								var f any
								_ = f
								f = a2.Get(_g)
								_g++
								m.VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), f)
							}
						}

						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							for {
								var _hx_tmp_31 int32
								_ = _hx_tmp_31
								_hx_tmp_31 = _g1
								if !(_hx_tmp_31 < Hx_Field_go_haxe_hxdynamic_getArrayLength(a1)) {
									break
								}

								var f1 any
								_ = f1
								f1 = a1.Get(_g1)
								_g1++
								var f2 any
								_ = f2
								f2 = m.VTable.Hx_Field_get(Hx_Field_go_haxe_hxdynamic_getField(f1, "name"))
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f2, nil))) {
									return false
								}

								var _hx_tmp_32 Hx_Enum_loreline_lorscript_ttype
								_ = _hx_tmp_32
								_hx_tmp_32 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								if !this.VTable.Hx_Field_typeEq(_hx_tmp_32, func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()) {
									return false
								}
							}
						}

						return true
					}
				}

			default:
				{
				}
			}
		}

	default:
		{
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_33 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_33
					_hx_tmp_33 = t1
					return this.VTable.Hx_Field_typeEq(_hx_tmp_33, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_typeEq(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_34 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_34
					_hx_tmp_34 = t1
					var _hx_tmp_35 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_35
					_hx_tmp_35 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_typeEq(_hx_tmp_34, this.VTable.Hx_Field_apply(_hx_tmp_35, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_tryUnify(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype) bool {
	if t1 == t2 {
		return true
	}

	switch t1.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t1.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				if !this.VTable.Hx_Field_link(t1, t2, r) {
					return false
				}

				Hx_Field_go_haxe_hxdynamic_setField(r, "r", t2)
				return true
			}

			return this.VTable.Hx_Field_tryUnify(func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}(), t2)
		}

	case 2:
		{
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_0
					_hx_tmp_0 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_0, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 3:
				{
					return true
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_1 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_1
					_hx_tmp_1 = t1
					var _hx_tmp_2 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_2
					_hx_tmp_2 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_1, this.VTable.Hx_Field_apply(_hx_tmp_2, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
				}
			}
		}

	case 5:
		{
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_3 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_3
					_hx_tmp_3 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_3, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_4 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_4
					_hx_tmp_4 = t1
					var _hx_tmp_5 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_5
					_hx_tmp_5 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_4, this.VTable.Hx_Field_apply(_hx_tmp_5, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
					return true
				}
			}
		}

	case 8:
		{
			var _gt Hx_Enum_loreline_lorscript_ttype
			_ = _gt
			_gt = t1.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_6 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_6
					_hx_tmp_6 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_6, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = _gt
					return this.VTable.Hx_Field_tryUnify(t1, t2)
				}

			case 8:
				{
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = _gt
					return this.VTable.Hx_Field_tryUnify(t1, t2)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_7 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_7
					_hx_tmp_7 = t1
					var _hx_tmp_8 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_8
					_hx_tmp_8 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_7, this.VTable.Hx_Field_apply(_hx_tmp_8, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = _gt
					return this.VTable.Hx_Field_tryUnify(t1, t2)
				}
			}
		}

	case 9:
		{
			var _gc any
			_ = _gc
			_gc = t1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_9 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_9
					_hx_tmp_9 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_9, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 9:
				{
					var cl2 any
					_ = cl2
					cl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var cl1 any
					_ = cl1
					cl1 = _gc
					for {
						if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(cl1, cl2)))) {
							break
						}

						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(cl1, "interfaces"), ((any)(nil))))) {
							var _g int32
							_ = _g
							_g = 0
							var _g1 HxArray[Hx_Enum_loreline_lorscript_ttype]
							_ = _g1
							_g1 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl1, "interfaces"))
							for {
								var _hx_tmp_10 int32
								_ = _hx_tmp_10
								_hx_tmp_10 = _g
								if !(_hx_tmp_10 < _g1.Len()) {
									break
								}

								var i Hx_Enum_loreline_lorscript_ttype
								_ = i
								i = _g1.Get(_g)
								_g++
								if i.Hx_Field_enumIndex() == 9 {
									var cli any
									_ = cli
									cli = i.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
									var args HxArray[Hx_Enum_loreline_lorscript_ttype]
									_ = args
									args = i.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
									{
										var tmp__g_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
										_ = tmp__g_1
										tmp__g_1 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
										{
											var tmp__g1_1 int32
											_ = tmp__g1_1
											tmp__g1_1 = 0
											for {
												var _hx_tmp_11 int32
												_ = _hx_tmp_11
												_hx_tmp_11 = tmp__g1_1
												if !(_hx_tmp_11 < args.Len()) {
													break
												}

												var a Hx_Enum_loreline_lorscript_ttype
												_ = a
												a = args.Get(tmp__g1_1)
												tmp__g1_1++
												var _hx_tmp_12 HxArray[Hx_Enum_loreline_lorscript_ttype]
												_ = _hx_tmp_12
												_hx_tmp_12 = tmp__g_1
												var _hx_tmp_13 Hx_Enum_loreline_lorscript_ttype
												_ = _hx_tmp_13
												_hx_tmp_13 = a
												_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_12, this.VTable.Hx_Field_apply(_hx_tmp_13, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl1, "params")), pl1)))
											}
										}

										var i1 Hx_Enum_loreline_lorscript_ttype
										_ = i1
										i1 = Hx_Enum_loreline_lorscript_ttype_TInst{cli, tmp__g_1}

										if this.VTable.Hx_Field_tryUnify(i1, t2) {
											return true
										}
									}
								} else {
									panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
								}
							}
						}

						{
							var _g Hx_Enum_loreline_lorscript_ttype
							_ = _g
							_g = func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cl1, "superClass"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
							if _g == nil {
								return false
							} else {
								if _g.Hx_Field_enumIndex() == 9 {
									var c any
									_ = c
									c = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
									var args HxArray[Hx_Enum_loreline_lorscript_ttype]
									_ = args
									args = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
									{
										var _g1 HxArray[Hx_Enum_loreline_lorscript_ttype]
										_ = _g1
										_g1 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
										{
											var _g2 int32
											_ = _g2
											_g2 = 0
											for {
												var _hx_tmp_14 int32
												_ = _hx_tmp_14
												_hx_tmp_14 = _g2
												if !(_hx_tmp_14 < args.Len()) {
													break
												}

												var a Hx_Enum_loreline_lorscript_ttype
												_ = a
												a = args.Get(_g2)
												_g2++
												var _hx_tmp_15 HxArray[Hx_Enum_loreline_lorscript_ttype]
												_ = _hx_tmp_15
												_hx_tmp_15 = _g1
												var _hx_tmp_16 Hx_Enum_loreline_lorscript_ttype
												_ = _hx_tmp_16
												_hx_tmp_16 = a
												_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_15, this.VTable.Hx_Field_apply(_hx_tmp_16, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl1, "params")), pl1)))
											}
										}

										pl1 = _g1
										cl1 = c
									}
								} else {
									panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
								}
							}
						}
					}

					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = pl1.Len()
						for _g < _g1 {
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							_hx_tmp_17 = _g
							_g = (_g + 1)
							var i int32
							_ = i
							i = _hx_tmp_17
							if !this.VTable.Hx_Field_typeEq(pl1.Get(i), pl2.Get(i)) {
								return false
							}
						}
					}

					return true
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_18 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_18
					_hx_tmp_18 = t1
					var _hx_tmp_19 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_19
					_hx_tmp_19 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_18, this.VTable.Hx_Field_apply(_hx_tmp_19, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 14:
				{
					var fl HxArray[any]
					_ = fl
					fl = t2.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var cl1 any
					_ = cl1
					cl1 = _gc
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)
						for _g < _g1 {
							var _hx_tmp_20 int32
							_ = _hx_tmp_20
							_hx_tmp_20 = _g
							_g = (_g + 1)
							var i int32
							_ = i
							i = _hx_tmp_20
							var f2 any
							_ = f2
							f2 = fl.Get(i)
							var f1 any
							_ = f1
							var cl any
							_ = cl
							cl = cl1
							for {
								var hx_dyncast_82 any
								_ = hx_dyncast_82
								hx_dyncast_82 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(cl, "fields"), "Hx_Obj_haxe_imap")
								var _hx_tmp_21 *Hx_Obj_haxe_imap
								_ = _hx_tmp_21
								if hx_dyncast_82 != nil {
									_hx_tmp_21 = hx_dyncast_82.(*Hx_Obj_haxe_imap)
								} else {
									_hx_tmp_21 = nil
								}

								f1 = _hx_tmp_21.VTable.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(Hx_Field_go_haxe_hxdynamic_getField(f2, "name"))
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(f1, nil))) {
									break
								}

								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(cl, "superClass"), ((any)(nil))))) {
									return false
								}

								var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
								_ = tmp__g_1
								tmp__g_1 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cl, "superClass"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								if tmp__g_1 == nil {
									panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
								} else {
									if tmp__g_1.Hx_Field_enumIndex() == 9 {
										var c any
										_ = c
										c = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
										cl = c
									} else {
										panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
									}
								}
							}

							var _hx_tmp_23 Hx_Enum_loreline_lorscript_ttype
							_ = _hx_tmp_23
							_hx_tmp_23 = func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
							var _hx_tmp_22 Hx_Enum_loreline_lorscript_ttype
							_ = _hx_tmp_22
							_hx_tmp_22 = this.VTable.Hx_Field_apply(_hx_tmp_23, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cl1, "params")), pl1)
							if !this.VTable.Hx_Field_typeEq(_hx_tmp_22, func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()) {
								return false
							}
						}
					}

					return true
				}

			default:
				{
				}
			}
		}

	case 11:
		{
			var _gt any
			_ = _gt
			_gt = t1.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_24 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_24
					_hx_tmp_24 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_24, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_25 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_25
					_hx_tmp_25 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(this.VTable.Hx_Field_apply(_hx_tmp_25, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}

			case 8:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_26 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_26
					_hx_tmp_26 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(this.VTable.Hx_Field_apply(_hx_tmp_26, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}

			case 11:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_27 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_27
					_hx_tmp_27 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(this.VTable.Hx_Field_apply(_hx_tmp_27, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}

			default:
				{
					var pl1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl1
					pl1 = _gargs
					var t1 any
					_ = t1
					t1 = _gt
					var _hx_tmp_28 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_28
					_hx_tmp_28 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(this.VTable.Hx_Field_apply(_hx_tmp_28, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), pl1), t2)
				}
			}
		}

	case 13:
		{
			var _gargs HxArray[any]
			_ = _gargs
			_gargs = t1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var _gret Hx_Enum_loreline_lorscript_ttype
			_ = _gret
			_gret = t1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_29 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_29
					_hx_tmp_29 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_29, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_30 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_30
					_hx_tmp_30 = t1
					var _hx_tmp_31 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_31
					_hx_tmp_31 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_30, this.VTable.Hx_Field_apply(_hx_tmp_31, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 12:
				{
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t2.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a, "name"), ((any)("haxe.Function"))))) {
						return true
					}
				}

			case 13:
				{
					var args2 HxArray[any]
					_ = args2
					args2 = t2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
					var r2 Hx_Enum_loreline_lorscript_ttype
					_ = r2
					r2 = t2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
					var r1 Hx_Enum_loreline_lorscript_ttype
					_ = r1
					r1 = _gret
					var args1 HxArray[any]
					_ = args1
					args1 = _gargs
					var _hx_tmp_32 int32
					_ = _hx_tmp_32
					_hx_tmp_32 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args1)
					if _hx_tmp_32 == Hx_Field_go_haxe_hxdynamic_getArrayLength(args2) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args1)
							for _g < _g1 {
								var _hx_tmp_33 int32
								_ = _hx_tmp_33
								_hx_tmp_33 = _g
								_g = (_g + 1)
								var i int32
								_ = i
								i = _hx_tmp_33
								var a1 any
								_ = a1
								a1 = args1.Get(i)
								var a2 any
								_ = a2
								a2 = args2.Get(i)
								var _hx_tmp_34 bool
								_ = _hx_tmp_34
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(a2, "opt"))) {
									_hx_tmp_34 = !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(a1, "opt"))))
								} else {
									_hx_tmp_34 = false
								}

								if _hx_tmp_34 {
									return false
								}

								var _hx_tmp_35 Hx_Enum_loreline_lorscript_ttype
								_ = _hx_tmp_35
								_hx_tmp_35 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a2, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								if !this.VTable.Hx_Field_tryUnify(_hx_tmp_35, func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a1, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()) {
									return false
								}
							}
						}

						return this.VTable.Hx_Field_tryUnify(r1, r2)
					}
				}

			default:
				{
				}
			}
		}

	case 14:
		{
			var _gfields HxArray[any]
			_ = _gfields
			_gfields = t1.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_36 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_36
					_hx_tmp_36 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_36, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_37 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_37
					_hx_tmp_37 = t1
					var _hx_tmp_38 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_38
					_hx_tmp_38 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_37, this.VTable.Hx_Field_apply(_hx_tmp_38, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			case 14:
				{
					var a2 HxArray[any]
					_ = a2
					a2 = t2.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
					var a1 HxArray[any]
					_ = a1
					a1 = _gfields
					if Hx_Field_go_haxe_hxdynamic_getArrayLength(a2) == 0 {
						return true
					}

					var m *Hx_Obj_haxe_imap
					_ = m
					m = (func() *Hx_Obj_haxe_imap {
						_hx_icast_83 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
						if _hx_icast_83 == nil {
							return nil
						}
						Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_83, "Hx_Obj_haxe_imap")
						return &Hx_Obj_haxe_imap{VTable: _hx_icast_83.VTable.(Hx_Obj_VTable_haxe_imap)}
					})()
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_39 int32
							_ = _hx_tmp_39
							_hx_tmp_39 = _g
							if !(_hx_tmp_39 < Hx_Field_go_haxe_hxdynamic_getArrayLength(a1)) {
								break
							}

							var f any
							_ = f
							f = a1.Get(_g)
							_g++
							m.VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), f)
						}
					}

					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_40 int32
							_ = _hx_tmp_40
							_hx_tmp_40 = _g1
							if !(_hx_tmp_40 < Hx_Field_go_haxe_hxdynamic_getArrayLength(a2)) {
								break
							}

							var f2 any
							_ = f2
							f2 = a2.Get(_g1)
							_g1++
							var f1 any
							_ = f1
							f1 = m.VTable.Hx_Field_get(Hx_Field_go_haxe_hxdynamic_getField(f2, "name"))
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f1, nil))) {
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(f2, "opt")))) {
									continue
								}

								return false
							}

							var _hx_tmp_41 Hx_Enum_loreline_lorscript_ttype
							_ = _hx_tmp_41
							_hx_tmp_41 = func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
							if !this.VTable.Hx_Field_typeEq(_hx_tmp_41, func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()) {
								return false
							}
						}
					}

					return true
				}

			default:
				{
				}
			}
		}

	default:
		{
			switch t2.Hx_Field_enumIndex() {
			case 0:
				{
					var r any
					_ = r
					r = t2.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
						if !this.VTable.Hx_Field_link(t2, t1, r) {
							return false
						}

						Hx_Field_go_haxe_hxdynamic_setField(r, "r", t1)
						return true
					}

					var _hx_tmp_42 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_42
					_hx_tmp_42 = t1
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_42, func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}

			case 5:
				{
					return true
				}

			case 8:
				{
					var t21 Hx_Enum_loreline_lorscript_ttype
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
					return this.VTable.Hx_Field_tryUnify(t1, t21)
				}

			case 11:
				{
					var t21 any
					_ = t21
					t21 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
					var pl2 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = pl2
					pl2 = t2.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
					var _hx_tmp_43 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_43
					_hx_tmp_43 = t1
					var _hx_tmp_44 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_44
					_hx_tmp_44 = func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t21, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()
					return this.VTable.Hx_Field_tryUnify(_hx_tmp_43, this.VTable.Hx_Field_apply(_hx_tmp_44, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t21, "params")), pl2))
				}

			default:
				{
				}
			}
		}
	}

	return this.VTable.Hx_Field_typeEq(t1, t2)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_unify(t1 Hx_Enum_loreline_lorscript_ttype, t2 Hx_Enum_loreline_lorscript_ttype, e any) {
	if !this.VTable.Hx_Field_tryUnify(t1, t2) {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(t1), " should be ")
		var e1 Hx_Enum_loreline_lorscript_errordef
		_ = e1
		e1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_0, Hx_Field_loreline_lorscript_checker_typeStr(t2))}

		var e2 *Hx_Obj_loreline_lorscript_error
		_ = e2
		e2 = Hx_Obj_loreline_lorscript_error_CreateInstance(e1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "line")))))
		if !this.Hx_Field_isCompletion {
			panic(Hx_Field_haxe_exception_thrown(((any)(e2))))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_apply(t Hx_Enum_loreline_lorscript_ttype, params HxArray[Hx_Enum_loreline_lorscript_ttype], args HxArray[Hx_Enum_loreline_lorscript_ttype]) Hx_Enum_loreline_lorscript_ttype {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = args.Len()
	if _hx_tmp_0 != params.Len() {
		panic(Hx_Field_haxe_exception_thrown(((any)("Invalid number of type parameters"))))
	}

	if args.Len() == 0 {
		return t
	}

	var subst *Hx_Obj_haxe_imap
	_ = subst
	subst = (func() *Hx_Obj_haxe_imap {
		_hx_icast_84 := Hx_Obj_haxe_ds_enumvaluemap_CreateInstance()
		if _hx_icast_84 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_84, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_84.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = params.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			func() func(_hx_param_key Hx_Enum_loreline_lorscript_ttype, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
				var _hx_closure_src_417 func(_hx_param_key any, _hx_param_value any)
				_ = _hx_closure_src_417
				_hx_closure_src_417 = subst.VTable.Hx_Field_set
				return func(_hx_param_key Hx_Enum_loreline_lorscript_ttype, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
					_hx_closure_src_417(((any)(_hx_param_key)), ((any)(_hx_param_value)))
				}
			}()(params.Get(i), args.Get(i))
		}
	}

	var _hx_reserved_map func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
	_ = _hx_reserved_map
	var _gthis *Hx_Obj_loreline_lorscript_checker
	_ = _gthis
	_gthis = this
	_hx_reserved_map = func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
		var st Hx_Enum_loreline_lorscript_ttype
		_ = st
		st = func() func(_hx_param_key Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
			var _hx_closure_src_418 func(_hx_param_key any) any
			_ = _hx_closure_src_418
			_hx_closure_src_418 = subst.VTable.Hx_Field_get
			return func(_hx_param_key Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
				return func() Hx_Enum_loreline_lorscript_ttype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_418(((any)(_hx_param_key))))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ttype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
				}()
			}
		}()(t)
		if st != nil {
			return st
		}

		return _gthis.VTable.Hx_Field_mapType(t, _hx_reserved_map)
	}

	return _hx_reserved_map(t)
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_mapType(t Hx_Enum_loreline_lorscript_ttype, f func(_hx_carg0 Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				return t
			}

			return f(t)
		}

	case 1, 2, 3, 4, 5:
		{
			return t
		}

	case 6:
		{
			return t
		}

	case 7:
		{
			return t
		}

	case 8:
		{
			var t1 Hx_Enum_loreline_lorscript_ttype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			return Hx_Enum_loreline_lorscript_ttype_TNull{f(t1)}
		}

	case 9:
		{
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			if _gargs.Len() == 0 {
				return t
			} else {
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _gargs
				var c any
				_ = c
				c = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
				var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = _g
				_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _hx_tmp_1
						_hx_tmp_1 = _g
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_1, f(t)))
					}
				}

				return Hx_Enum_loreline_lorscript_ttype_TInst{c, _g}
			}
		}

	case 10:
		{
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
			if _gargs.Len() == 0 {
				return t
			} else {
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _gargs
				var e any
				_ = e
				e = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_e
				var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = _g
				_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = _g1
						if !(_hx_tmp_2 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_3 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_3, f(t)))
					}
				}

				return Hx_Enum_loreline_lorscript_ttype_TEnum{e, _g}
			}
		}

	case 11:
		{
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			if _gargs.Len() == 0 {
				return t
			} else {
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _gargs
				var t1 any
				_ = t1
				t1 = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
				var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = _g
				_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g1
						if !(_hx_tmp_4 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_5 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _hx_tmp_5
						_hx_tmp_5 = _g
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_5, f(t)))
					}
				}

				return Hx_Enum_loreline_lorscript_ttype_TType{t1, _g}
			}
		}

	case 12:
		{
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			if _gargs.Len() == 0 {
				return t
			} else {
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _gargs
				var a any
				_ = a
				a = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a
				var _g HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = _g
				_g = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = _g1
						if !(_hx_tmp_6 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_7 HxArray[Hx_Enum_loreline_lorscript_ttype]
						_ = _hx_tmp_7
						_hx_tmp_7 = _g
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_7, f(t)))
					}
				}

				return Hx_Enum_loreline_lorscript_ttype_TAbstract{a, _g}
			}
		}

	case 13:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ttype
			_ = ret
			ret = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _g1
					if !(_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
						break
					}

					var a any
					_ = a
					a = args.Get(_g1)
					_g1++
					var _hx_tmp_9 HxArray[any]
					_ = _hx_tmp_9
					_hx_tmp_9 = HxMakeArrayView[any](_g)
					var _hx_tmp_10 any
					_ = _hx_tmp_10
					_hx_tmp_10 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
					var _hx_tmp_11 any
					_ = _hx_tmp_11
					_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_getField(a, "opt")
					_ = (Hx_Array_Push[any](_hx_tmp_9, any(map[string]any{"name": _hx_tmp_10, "opt": _hx_tmp_11, "t": ((any)(f(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())))})))
				}
			}

			var _hx_tmp_12 HxArray[any]
			_ = _hx_tmp_12
			_hx_tmp_12 = _g
			return Hx_Enum_loreline_lorscript_ttype_TFun{_hx_tmp_12, f(ret)}
		}

	case 14:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = _g1
					if !(_hx_tmp_13 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
						break
					}

					var af any
					_ = af
					af = fields.Get(_g1)
					_g1++
					var _hx_tmp_14 HxArray[any]
					_ = _hx_tmp_14
					_hx_tmp_14 = HxMakeArrayView[any](_g)
					var _hx_tmp_15 any
					_ = _hx_tmp_15
					_hx_tmp_15 = Hx_Field_go_haxe_hxdynamic_getField(af, "name")
					var _hx_tmp_16 any
					_ = _hx_tmp_16
					_hx_tmp_16 = Hx_Field_go_haxe_hxdynamic_getField(af, "opt")
					_ = (Hx_Array_Push[any](_hx_tmp_14, any(map[string]any{"name": _hx_tmp_15, "opt": _hx_tmp_16, "t": ((any)(f(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(af, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())))})))
				}
			}

			return Hx_Enum_loreline_lorscript_ttype_TAnon{_g}
		}

	case 15:
		{
			var l func() Hx_Enum_loreline_lorscript_ttype
			_ = l
			l = t.(Hx_Enum_loreline_lorscript_ttype_TLazy).Hx_Field_f
			return f(l())
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_follow(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				return this.VTable.Hx_Field_follow(func() Hx_Enum_loreline_lorscript_ttype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ttype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
				}())
			} else {
				return t
			}
		}

	case 8:
		{
			var t1 Hx_Enum_loreline_lorscript_ttype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			return this.VTable.Hx_Field_follow(t1)
		}

	case 11:
		{
			var t1 any
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
			_ = _hx_tmp_0
			_hx_tmp_0 = func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t1, "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}()
			return this.VTable.Hx_Field_follow(this.VTable.Hx_Field_apply(_hx_tmp_0, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(t1, "params")), args))
		}

	case 15:
		{
			var f func() Hx_Enum_loreline_lorscript_ttype
			_ = f
			f = t.(Hx_Enum_loreline_lorscript_ttype_TLazy).Hx_Field_f
			return this.VTable.Hx_Field_follow(f())
		}

	default:
		{
			return t
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_getFields(t Hx_Enum_loreline_lorscript_ttype) HxArray[any] {
	var fields HxArray[any]
	_ = fields
	fields = HxMakeArray[any]().Dyn()
	var _gthis *Hx_Obj_loreline_lorscript_checker
	_ = _gthis
	_gthis = this
	{
		var _g Hx_Enum_loreline_lorscript_ttype
		_ = _g
		_g = this.VTable.Hx_Field_follow(t)
		switch _g.Hx_Field_enumIndex() {
		case 9:
			{
				var c any
				_ = c
				c = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
				{
					var _hx_reserved_map func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
					_ = _hx_reserved_map
					_hx_reserved_map = func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
						var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
						_ = _hx_tmp_0
						_hx_tmp_0 = t
						return _gthis.VTable.Hx_Field_apply(_hx_tmp_0, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "params")), args)
					}

					for {
						if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(c, nil)))) {
							break
						}

						{
							var hx_dyncast_85 any
							_ = hx_dyncast_85
							hx_dyncast_85 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"), "Hx_Obj_haxe_imap")
							var _hx_tmp_1 *Hx_Obj_haxe_imap
							_ = _hx_tmp_1
							if hx_dyncast_85 != nil {
								_hx_tmp_1 = hx_dyncast_85.(*Hx_Obj_haxe_imap)
							} else {
								_hx_tmp_1 = nil
							}

							var fname any
							_ = fname
							fname = _hx_tmp_1.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_keys()
							for {
								var _hx_tmp_2 any
								_ = _hx_tmp_2
								_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(fname, "hasNext")
								if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
									break
								}

								var _hx_tmp_3 any
								_ = _hx_tmp_3
								_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(fname, "next")
								var fname1 string
								_ = fname1
								fname1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))
								var hx_dyncast_86 any
								_ = hx_dyncast_86
								hx_dyncast_86 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"), "Hx_Obj_haxe_imap")
								var _hx_tmp_4 *Hx_Obj_haxe_imap
								_ = _hx_tmp_4
								if hx_dyncast_86 != nil {
									_hx_tmp_4 = hx_dyncast_86.(*Hx_Obj_haxe_imap)
								} else {
									_hx_tmp_4 = nil
								}

								var f any
								_ = f
								f = _hx_tmp_4.VTable.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(fname1)
								var _hx_tmp_5 bool
								_ = _hx_tmp_5
								if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(f, "isPublic")))) {
									_hx_tmp_5 = true
								} else {
									_hx_tmp_5 = !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(f, "complete"))))
								}

								if _hx_tmp_5 {
									continue
								}

								var name string
								_ = name
								name = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name"))
								var t Hx_Enum_loreline_lorscript_ttype
								_ = t
								t = _hx_reserved_map(func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}())
								var _hx_tmp_6 bool
								_ = _hx_tmp_6
								if this.Hx_Field_allowAsync {
									_hx_tmp_6 = Hx_Field_stringtools_startsWith(name, "a_")
								} else {
									_hx_tmp_6 = false
								}

								if _hx_tmp_6 {
									t = this.VTable.Hx_Field_unasync(t)
									var _hx_tmp_7 string
									_ = _hx_tmp_7
									_hx_tmp_7 = name
									name = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, 2, HxNullable[int32]{}))
								}

								var _hx_tmp_8 HxArray[any]
								_ = _hx_tmp_8
								_hx_tmp_8 = HxMakeArrayView[any](fields)
								_ = (Hx_Array_Push[any](_hx_tmp_8, any(map[string]any{"name": ((any)(name)), "t": ((any)(t))})))
							}
						}

						var _hx_tmp_9 bool
						_ = _hx_tmp_9
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(c, "isInterface"))) {
							_hx_tmp_9 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"), ((any)(nil))))))
						} else {
							_hx_tmp_9 = false
						}

						if _hx_tmp_9 {
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g1 HxArray[Hx_Enum_loreline_lorscript_ttype]
							_ = _g1
							_g1 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"))
							for {
								var _hx_tmp_10 int32
								_ = _hx_tmp_10
								_hx_tmp_10 = tmp__g_1
								if !(_hx_tmp_10 < _g1.Len()) {
									break
								}

								var i Hx_Enum_loreline_lorscript_ttype
								_ = i
								i = _g1.Get(tmp__g_1)
								tmp__g_1++
								{
									var _g2 int32
									_ = _g2
									_g2 = 0
									var _g3 HxArray[any]
									_ = _g3
									_g3 = this.VTable.Hx_Field_getFields(i)
									for {
										var _hx_tmp_11 int32
										_ = _hx_tmp_11
										_hx_tmp_11 = _g2
										if !(_hx_tmp_11 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g3)) {
											break
										}

										var f any
										_ = f
										f = _g3.Get(_g2)
										_g2++
										var _hx_tmp_12 HxArray[any]
										_ = _hx_tmp_12
										_hx_tmp_12 = HxMakeArrayView[any](fields)
										var _hx_tmp_13 any
										_ = _hx_tmp_13
										_hx_tmp_13 = Hx_Field_go_haxe_hxdynamic_getField(f, "name")
										_ = (Hx_Array_Push[any](_hx_tmp_12, any(map[string]any{"name": _hx_tmp_13, "t": ((any)(_hx_reserved_map(func() Hx_Enum_loreline_lorscript_ttype {
											_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
											if _hx_d == nil {
												var _hx_z Hx_Enum_loreline_lorscript_ttype
												return _hx_z
											}
											return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
										}())))})))
									}
								}
							}
						}

						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"), ((any)(nil))))) {
							break
						}

						{
							var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp__g_1
							tmp__g_1 = func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
							if tmp__g_1 == nil {
								break
							} else {
								if tmp__g_1.Hx_Field_enumIndex() == 9 {
									var csup any
									_ = csup
									csup = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
									var tmp_args_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
									_ = tmp_args_1
									tmp_args_1 = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
									{
										var curMap func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype
										_ = curMap
										curMap = _hx_reserved_map
										_hx_reserved_map = func(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
											var _hx_tmp_14 Hx_Enum_loreline_lorscript_ttype
											_ = _hx_tmp_14
											_hx_tmp_14 = t
											return curMap(_gthis.VTable.Hx_Field_apply(_hx_tmp_14, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(csup, "params")), tmp_args_1))
										}

										c = csup
									}
								} else {
									break
								}
							}
						}
					}
				}
			}

		case 13:
			{
				var args HxArray[any]
				_ = args
				args = _g.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
				var ret Hx_Enum_loreline_lorscript_ttype
				_ = ret
				ret = _g.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
				if this.Hx_Field_isCompletion {
					var _hx_tmp_15 HxArray[any]
					_ = _hx_tmp_15
					_hx_tmp_15 = HxMakeArrayView[any](fields)
					_ = (Hx_Array_Push[any](_hx_tmp_15, any(map[string]any{"name": ((any)("bind")), "t": ((any)(Hx_Enum_loreline_lorscript_ttype_TFun{args, Hx_Enum_loreline_lorscript_ttype_TVoid{}}))})))
				}
			}

		case 14:
			{
				var fl HxArray[any]
				_ = fl
				fl = _g.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_16 int32
						_ = _hx_tmp_16
						_hx_tmp_16 = _g1
						if !(_hx_tmp_16 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
							break
						}

						var f any
						_ = f
						f = fl.Get(_g1)
						_g1++
						var _hx_tmp_17 HxArray[any]
						_ = _hx_tmp_17
						_hx_tmp_17 = HxMakeArrayView[any](fields)
						var _hx_tmp_18 any
						_ = _hx_tmp_18
						_hx_tmp_18 = Hx_Field_go_haxe_hxdynamic_getField(f, "name")
						_ = (Hx_Array_Push[any](_hx_tmp_17, any(map[string]any{"name": _hx_tmp_18, "t": Hx_Field_go_haxe_hxdynamic_getField(f, "t")})))
					}
				}
			}

		default:
			{
			}
		}
	}

	return fields
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_getField(t Hx_Enum_loreline_lorscript_ttype, f string, e any, forWrite__isOptional any) Hx_Enum_loreline_lorscript_ttype {
	forWrite := HxDefault[bool]()
	_ = forWrite

	if forWrite__isOptional != nil {
		forWrite = HxConvert[bool](forWrite__isOptional)
	} else {
		forWrite = false
	}
	{
		var _g Hx_Enum_loreline_lorscript_ttype
		_ = _g
		_g = this.VTable.Hx_Field_follow(t)
		switch _g.Hx_Field_enumIndex() {
		case 5:
			{
				return this.VTable.Hx_Field_makeMono()
			}

		case 9:
			{
				var c any
				_ = c
				c = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
				var args HxArray[Hx_Enum_loreline_lorscript_ttype]
				_ = args
				args = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
				var hx_dyncast_87 any
				_ = hx_dyncast_87
				hx_dyncast_87 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"), "Hx_Obj_haxe_imap")
				var _hx_tmp_0 *Hx_Obj_haxe_imap
				_ = _hx_tmp_0
				if hx_dyncast_87 != nil {
					_hx_tmp_0 = hx_dyncast_87.(*Hx_Obj_haxe_imap)
				} else {
					_hx_tmp_0 = nil
				}

				var cf any
				_ = cf
				cf = _hx_tmp_0.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(f)
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(cf, nil))) {
					_hx_tmp_1 = this.Hx_Field_allowAsync
				} else {
					_hx_tmp_1 = false
				}

				if _hx_tmp_1 {
					var hx_dyncast_88 any
					_ = hx_dyncast_88
					hx_dyncast_88 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"), "Hx_Obj_haxe_imap")
					var _hx_tmp_2 *Hx_Obj_haxe_imap
					_ = _hx_tmp_2
					if hx_dyncast_88 != nil {
						_hx_tmp_2 = hx_dyncast_88.(*Hx_Obj_haxe_imap)
					} else {
						_hx_tmp_2 = nil
					}

					cf = _hx_tmp_2.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(Hx_Field_go_haxe_hxdynamic_stringConcat("a_", f))
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(cf, nil))) {
						var isPublic bool
						_ = isPublic
						isPublic = true
						var _hx_tmp_3 any
						_ = _hx_tmp_3
						_hx_tmp_3 = ((any)(isPublic))
						var _hx_tmp_4 any
						_ = _hx_tmp_4
						_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(cf, "params")
						var _hx_tmp_5 any
						_ = _hx_tmp_5
						_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(cf, "name")
						var _hx_tmp_6 any
						_ = _hx_tmp_6
						_hx_tmp_6 = ((any)(this.VTable.Hx_Field_unasync(func() Hx_Enum_loreline_lorscript_ttype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cf, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ttype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
						}())))
						cf = any(map[string]any{"isPublic": _hx_tmp_3, "canWrite": ((any)(false)), "params": _hx_tmp_4, "name": _hx_tmp_5, "t": _hx_tmp_6, "complete": Hx_Field_go_haxe_hxdynamic_getField(cf, "complete")})
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(cf, "t"), ((any)(nil))))) {
							cf = nil
						}
					}
				}

				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(cf, nil))) {
					_hx_tmp_8 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(c, "isInterface"))))
				} else {
					_hx_tmp_8 = false
				}

				if _hx_tmp_8 {
					_hx_tmp_7 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"), ((any)(nil))))))
				} else {
					_hx_tmp_7 = false
				}

				if _hx_tmp_7 {
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var _g1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = _g1
					_g1 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "interfaces"))
					for {
						var _hx_tmp_9 int32
						_ = _hx_tmp_9
						_hx_tmp_9 = tmp__g_1
						if !(_hx_tmp_9 < _g1.Len()) {
							break
						}

						var i Hx_Enum_loreline_lorscript_ttype
						_ = i
						i = _g1.Get(tmp__g_1)
						tmp__g_1++
						var ft Hx_Enum_loreline_lorscript_ttype
						_ = ft
						ft = this.VTable.Hx_Field_getField(i, f, e, forWrite)
						if ft != nil {
							var _hx_tmp_10 Hx_Enum_loreline_lorscript_ttype
							_ = _hx_tmp_10
							_hx_tmp_10 = ft
							return this.VTable.Hx_Field_apply(_hx_tmp_10, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "params")), args)
						}
					}
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(cf, nil))) {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"), ((any)(nil))))) {
						return nil
					}

					var ft Hx_Enum_loreline_lorscript_ttype
					_ = ft
					ft = this.VTable.Hx_Field_getField(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(c, "superClass"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}(), f, e, forWrite)
					if ft != nil {
						var _hx_tmp_11 Hx_Enum_loreline_lorscript_ttype
						_ = _hx_tmp_11
						_hx_tmp_11 = ft
						ft = this.VTable.Hx_Field_apply(_hx_tmp_11, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "params")), args)
					}

					return ft
				}

				if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(cf, "isPublic")))) {
					var _hx_tmp_12 any
					_ = _hx_tmp_12
					_hx_tmp_12 = ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Can't access private field ", f), " on ")))
					var e1 Hx_Enum_loreline_lorscript_errordef
					_ = e1
					e1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_12, Hx_Field_go_haxe_hxdynamic_getField(c, "name")))}

					var e2 *Hx_Obj_loreline_lorscript_error
					_ = e2
					e2 = Hx_Obj_loreline_lorscript_error_CreateInstance(e1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e2))))
					}
				}

				var _hx_tmp_13 bool
				_ = _hx_tmp_13
				if forWrite {
					_hx_tmp_13 = !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(cf, "canWrite"))))
				} else {
					_hx_tmp_13 = false
				}

				if _hx_tmp_13 {
					var _hx_tmp_14 any
					_ = _hx_tmp_14
					_hx_tmp_14 = ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Can't write readonly field ", f), " on ")))
					var e1 Hx_Enum_loreline_lorscript_errordef
					_ = e1
					e1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_14, Hx_Field_go_haxe_hxdynamic_getField(c, "name")))}

					var e2 *Hx_Obj_loreline_lorscript_error
					_ = e2
					e2 = Hx_Obj_loreline_lorscript_error_CreateInstance(e1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e2))))
					}
				}

				var t Hx_Enum_loreline_lorscript_ttype
				_ = t
				t = func() Hx_Enum_loreline_lorscript_ttype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cf, "t"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ttype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
				}()
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(cf, "params"), ((any)(nil))))) {
					var cf1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = cf1
					cf1 = HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(cf, "params"))
					var tmp__g_1 HxArray[Hx_Enum_loreline_lorscript_ttype]
					_ = tmp__g_1
					tmp__g_1 = HxMakeArray[Hx_Enum_loreline_lorscript_ttype]()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 int32
						_ = _g2
						_g2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(cf, "params"), "length"))))
						for _g1 < _g2 {
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							_hx_tmp_15 = _g1
							_g1 = (_g1 + 1)
							var i int32
							_ = i
							i = _hx_tmp_15
							var _hx_tmp_16 HxArray[Hx_Enum_loreline_lorscript_ttype]
							_ = _hx_tmp_16
							_hx_tmp_16 = tmp__g_1
							_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ttype](_hx_tmp_16, this.VTable.Hx_Field_makeMono()))
						}
					}

					t = this.VTable.Hx_Field_apply(t, cf1, tmp__g_1)
				}

				var _hx_tmp_17 Hx_Enum_loreline_lorscript_ttype
				_ = _hx_tmp_17
				_hx_tmp_17 = t
				return this.VTable.Hx_Field_apply(_hx_tmp_17, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(c, "params")), args)
			}

		case 14:
			{
				var fields HxArray[any]
				_ = fields
				fields = _g.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = _g1
						if !(_hx_tmp_18 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
							break
						}

						var af any
						_ = af
						af = fields.Get(_g1)
						_g1++
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(af, "name"), ((any)(f))))) {
							return func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(af, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
						}
					}
				}

				return nil
			}

		default:
			{
				return nil
			}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_unasync(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	{
		var _g Hx_Enum_loreline_lorscript_ttype
		_ = _g
		_g = this.VTable.Hx_Field_follow(t)
		if _g.Hx_Field_enumIndex() == 13 {
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ttype
			_ = ret
			ret = _g.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(args) > 0 {
				var rargs HxArray[any]
				_ = rargs
				rargs = (Hx_Array_Copy[any](HxMakeArrayView[any](args)))
				{
					var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
					_ = tmp__g_1
					tmp__g_1 = this.VTable.Hx_Field_follow(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField((Hx_Array_Shift[any](HxMakeArrayView[any](rargs)).Value), "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
					if tmp__g_1.Hx_Field_enumIndex() == 13 {
						var _gargs HxArray[any]
						_ = _gargs
						_gargs = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
						if Hx_Field_go_haxe_hxdynamic_getArrayLength(_gargs) == 1 {
							var r any
							_ = r
							r = _gargs.Get(0)
							var _hx_tmp_0 HxArray[any]
							_ = _hx_tmp_0
							_hx_tmp_0 = rargs
							return Hx_Enum_loreline_lorscript_ttype_TFun{_hx_tmp_0, func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()}
						}
					}
				}
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_typeExprWith(expr any, t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = expr
	var et Hx_Enum_loreline_lorscript_ttype
	_ = et
	et = this.VTable.Hx_Field_typeExpr(_hx_tmp_0, Hx_Enum_loreline_lorscript__checker_withtype_WithType{t})
	this.VTable.Hx_Field_unify(et, t, expr)
	return t
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_makeMono() Hx_Enum_loreline_lorscript_ttype {
	return Hx_Enum_loreline_lorscript_ttype_TMono{any(map[string]any{"r": ((any)(nil))})}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_makeIterator(t Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = any(map[string]any{"name": ((any)("next")), "opt": ((any)(false)), "t": ((any)(Hx_Enum_loreline_lorscript_ttype_TFun{HxMakeArray[any]().Dyn(), t}))})
	return Hx_Enum_loreline_lorscript_ttype_TAnon{HxMakeArray[any](_hx_tmp_0, any(map[string]any{"name": ((any)("hasNext")), "opt": ((any)(false)), "t": ((any)(Hx_Enum_loreline_lorscript_ttype_TFun{HxMakeArray[any]().Dyn(), Hx_Enum_loreline_lorscript_ttype_TBool{}}))})).Dyn()}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_mk(e Hx_Enum_loreline_lorscript_exprdef, p any) any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(e))
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(p, "pmin")
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(p, "pmax")
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(p, "origin")
	return any(map[string]any{"e": _hx_tmp_0, "pmin": _hx_tmp_1, "pmax": _hx_tmp_2, "origin": _hx_tmp_3, "line": Hx_Field_go_haxe_hxdynamic_getField(p, "line")})
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_isString(t Hx_Enum_loreline_lorscript_ttype) bool {
	t = this.VTable.Hx_Field_follow(t)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if t.Hx_Field_enumIndex() == 9 {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c, "name"), ((any)("String"))))))
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_onCompletion(expr any, t Hx_Enum_loreline_lorscript_ttype) {
	if this.Hx_Field_isCompletion {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_completion_CreateInstance(expr, t)))))
	}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_typeField(o any, f string, expr any, forWrite bool) Hx_Enum_loreline_lorscript_ttype {
	var ot Hx_Enum_loreline_lorscript_ttype
	_ = ot
	ot = this.VTable.Hx_Field_typeExpr(o, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
	if f == HxStringNull {
		this.VTable.Hx_Field_onCompletion(expr, ot)
	}

	var ft Hx_Enum_loreline_lorscript_ttype
	_ = ft
	ft = this.VTable.Hx_Field_getField(ot, f, expr, forWrite)
	if ft == nil {
		{
			var e Hx_Enum_loreline_lorscript_errordef
			_ = e
			e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(ot), " has no field "), f)}

			var e1 *Hx_Obj_loreline_lorscript_error
			_ = e1
			e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
			if !this.Hx_Field_isCompletion {
				panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
			}
		}

		ft = Hx_Enum_loreline_lorscript_ttype_TDynamic{}
	}

	return ft
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_typeExpr(expr any, withType Hx_Enum_loreline_lorscript__checker_withtype) Hx_Enum_loreline_lorscript_ttype {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(expr, nil))) {
		_hx_tmp_0 = this.Hx_Field_isCompletion
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		if withType.Hx_Field_enumIndex() == 2 {
			var t Hx_Enum_loreline_lorscript_ttype
			_ = t
			t = withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t
			return t
		} else {
			return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
		}
	}

	var _gthis *Hx_Obj_loreline_lorscript_checker
	_ = _gthis
	_gthis = this
	{
		var _g Hx_Enum_loreline_lorscript_exprdef
		_ = _g
		_g = func() Hx_Enum_loreline_lorscript_exprdef {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(expr, "e"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_loreline_lorscript_exprdef
				return _hx_z
			}
			return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
		}()
		switch _g.Hx_Field_enumIndex() {
		case 0:
			{
				var c Hx_Enum_loreline_lorscript_const
				_ = c
				c = _g.(Hx_Enum_loreline_lorscript_exprdef_EConst).Hx_Field_c
				switch c.Hx_Field_enumIndex() {
				case 0:
					{
						return Hx_Enum_loreline_lorscript_ttype_TInt{}
					}

				case 1:
					{
						return Hx_Enum_loreline_lorscript_ttype_TFloat{}
					}

				case 2:
					{
						return this.Hx_Field_types.Hx_Field_t_string
					}

				default:
					{
						panic(Hx_Field_haxe_exception_thrown("unreachable"))
					}
				}
			}

		case 1:
			{
				var v string
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
				{
					var l Hx_Enum_loreline_lorscript_ttype
					_ = l
					l = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
						var _hx_closure_src_419 func(_hx_param_key any) any
						_ = _hx_closure_src_419
						_hx_closure_src_419 = this.Hx_Field_locals.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
						return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
							return func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_419(((any)(_hx_param_key))))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
						}
					}()(v)
					if l != nil {
						return l
					}

					var g Hx_Enum_loreline_lorscript_ttype
					_ = g
					g = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
						var _hx_closure_src_420 func(_hx_param_key any) any
						_ = _hx_closure_src_420
						_hx_closure_src_420 = this.Hx_Field_globals.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
						return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
							return func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_420(((any)(_hx_param_key))))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
						}
					}()(v)
					if g != nil {
						if g == nil {
							return g
						} else {
							if g.Hx_Field_enumIndex() == 15 {
								var f func() Hx_Enum_loreline_lorscript_ttype
								_ = f
								f = g.(Hx_Enum_loreline_lorscript_ttype_TLazy).Hx_Field_f
								return f()
							} else {
								return g
							}
						}
					}

					if this.Hx_Field_allowAsync {
						g = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
							var _hx_closure_src_421 func(_hx_param_key any) any
							_ = _hx_closure_src_421
							_hx_closure_src_421 = this.Hx_Field_globals.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
							return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
								return func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_421(((any)(_hx_param_key))))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
							}
						}()(Hx_Field_go_haxe_hxdynamic_stringConcat("a_", v))
						if g != nil {
							g = this.VTable.Hx_Field_unasync(g)
						}

						if g != nil {
							return g
						}
					}

					switch v {
					case "false", "true":
						{
							return Hx_Enum_loreline_lorscript_ttype_TBool{}
						}

					case "null":
						{
							return this.VTable.Hx_Field_makeMono()
						}

					case "trace":
						{
							return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
						}

					default:
						{
							if withType.Hx_Field_enumIndex() == 2 {
								var _gt Hx_Enum_loreline_lorscript_ttype
								_ = _gt
								_gt = withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t
								if _gt.Hx_Field_enumIndex() == 10 {
									var et Hx_Enum_loreline_lorscript_ttype
									_ = et
									et = _gt
									var e any
									_ = e
									e = _gt.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_e
									var args HxArray[Hx_Enum_loreline_lorscript_ttype]
									_ = args
									args = _gt.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
									{
										var tmp__g_1 int32
										_ = tmp__g_1
										tmp__g_1 = 0
										var _g1 HxArray[any]
										_ = _g1
										_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(e, "constructors"))
										for {
											var _hx_tmp_1 int32
											_ = _hx_tmp_1
											_hx_tmp_1 = tmp__g_1
											if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
												break
											}

											var c any
											_ = c
											c = _g1.Get(tmp__g_1)
											tmp__g_1++
											if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c, "name"), ((any)(v))))) {
												if this.VTable.Hx_Field_onTopDownEnum(e, v) {
													var _hx_tmp_2 Hx_Enum_loreline_lorscript_ttype
													_ = _hx_tmp_2
													if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c, "args"), ((any)(nil))))) {
														_hx_tmp_2 = et
													} else {
														_hx_tmp_2 = Hx_Enum_loreline_lorscript_ttype_TFun{HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "args")), et}
													}

													var ct Hx_Enum_loreline_lorscript_ttype
													_ = ct
													ct = _hx_tmp_2
													var _hx_tmp_3 Hx_Enum_loreline_lorscript_ttype
													_ = _hx_tmp_3
													_hx_tmp_3 = ct
													return this.VTable.Hx_Field_apply(_hx_tmp_3, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(e, "params")), args)
												}

												break
											}
										}
									}
								}
							}

							if this.Hx_Field_isCompletion {
								return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
							}

							{
								var e Hx_Enum_loreline_lorscript_errordef
								_ = e
								e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Unknown identifier ", v)}

								var e1 *Hx_Obj_loreline_lorscript_error
								_ = e1
								e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
								if !this.Hx_Field_isCompletion {
									panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
								}
							}
						}
					}
				}
			}

		case 2:
			{
				var n string
				_ = n
				n = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_n
				var t Hx_Enum_loreline_lorscript_ctype
				_ = t
				t = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_t
				var init any
				_ = init
				init = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_e
				var _hx_tmp_4 Hx_Enum_loreline_lorscript_ttype
				_ = _hx_tmp_4
				if t == nil {
					_hx_tmp_4 = this.VTable.Hx_Field_makeMono()
				} else {
					_hx_tmp_4 = this.VTable.Hx_Field_makeType(t, expr)
				}

				var vt Hx_Enum_loreline_lorscript_ttype
				_ = vt
				vt = _hx_tmp_4
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(init, nil))) {
					var _hx_tmp_5 any
					_ = _hx_tmp_5
					_hx_tmp_5 = init
					var _hx_tmp_6 Hx_Enum_loreline_lorscript__checker_withtype
					_ = _hx_tmp_6
					if t == nil {
						_hx_tmp_6 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
					} else {
						_hx_tmp_6 = Hx_Enum_loreline_lorscript__checker_withtype_WithType{vt}
					}

					var et Hx_Enum_loreline_lorscript_ttype
					_ = et
					et = this.VTable.Hx_Field_typeExpr(_hx_tmp_5, _hx_tmp_6)
					if t == nil {
						vt = et
					} else {
						this.VTable.Hx_Field_unify(et, vt, init)
					}
				}

				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
					var _hx_closure_src_422 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_422
					_hx_closure_src_422 = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
						_hx_closure_src_422(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(n, vt)
				return Hx_Enum_loreline_lorscript_ttype_TVoid{}
			}

		case 3:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
				return this.VTable.Hx_Field_typeExpr(e, withType)
			}

		case 4:
			{
				var el HxArray[any]
				_ = el
				el = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
				var t Hx_Enum_loreline_lorscript_ttype
				_ = t
				t = Hx_Enum_loreline_lorscript_ttype_TVoid{}

				var locals *Hx_Obj_haxe_imap
				_ = locals
				locals = this.VTable.Hx_Field_saveLocals()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = _g1
						if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
							break
						}

						var e any
						_ = e
						e = el.Get(_g1)
						_g1++
						var _hx_tmp_8 any
						_ = _hx_tmp_8
						_hx_tmp_8 = e
						var _hx_tmp_9 Hx_Enum_loreline_lorscript__checker_withtype
						_ = _hx_tmp_9
						var _hx_tmp_10 any
						_ = _hx_tmp_10
						_hx_tmp_10 = e
						var _hx_tmp_11 HxArray[any]
						_ = _hx_tmp_11
						_hx_tmp_11 = el
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_10, _hx_tmp_11.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(el) - 1))))) {
							_hx_tmp_9 = withType
						} else {
							_hx_tmp_9 = Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}
						}

						t = this.VTable.Hx_Field_typeExpr(_hx_tmp_8, _hx_tmp_9)
					}
				}

				this.Hx_Field_locals = locals
				return t
			}

		case 5:
			{
				var o any
				_ = o
				o = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
				var f string
				_ = f
				f = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
				return this.VTable.Hx_Field_typeField(o, f, expr, false)
			}

		case 6:
			{
				var op string
				_ = op
				op = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op
				var e1 any
				_ = e1
				e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
				var e2 any
				_ = e2
				e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
				switch op {
				case "!=", "==":
					{
						var t1 Hx_Enum_loreline_lorscript_ttype
						_ = t1
						t1 = this.VTable.Hx_Field_typeExpr(e1, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
						var _hx_tmp_12 any
						_ = _hx_tmp_12
						_hx_tmp_12 = e2
						var t2 Hx_Enum_loreline_lorscript_ttype
						_ = t2
						t2 = this.VTable.Hx_Field_typeExpr(_hx_tmp_12, Hx_Enum_loreline_lorscript__checker_withtype_WithType{t1})
						if !this.VTable.Hx_Field_tryUnify(t1, t2) {
							this.VTable.Hx_Field_unify(t2, t1, e2)
						}

						return Hx_Enum_loreline_lorscript_ttype_TBool{}
					}

				case "%", "*", "-", "/":
					{
						var _hx_tmp_13 any
						_ = _hx_tmp_13
						_hx_tmp_13 = e1
						var t1 Hx_Enum_loreline_lorscript_ttype
						_ = t1
						t1 = this.VTable.Hx_Field_typeExpr(_hx_tmp_13, Hx_Enum_loreline_lorscript__checker_withtype_WithType{Hx_Enum_loreline_lorscript_ttype_TInt{}})
						var _hx_tmp_14 any
						_ = _hx_tmp_14
						_hx_tmp_14 = e2
						var t2 Hx_Enum_loreline_lorscript_ttype
						_ = t2
						t2 = this.VTable.Hx_Field_typeExpr(_hx_tmp_14, Hx_Enum_loreline_lorscript__checker_withtype_WithType{t1})
						if !this.VTable.Hx_Field_tryUnify(t1, t2) {
							this.VTable.Hx_Field_unify(t2, t1, e2)
						}

						{
							var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp__g_1
							tmp__g_1 = this.VTable.Hx_Field_follow(t1)
							var _g1 Hx_Enum_loreline_lorscript_ttype
							_ = _g1
							_g1 = this.VTable.Hx_Field_follow(t2)
							switch tmp__g_1.Hx_Field_enumIndex() {
							case 2:
								{
									switch _g1.Hx_Field_enumIndex() {
									case 2:
										{
											if op == "/" {
												return Hx_Enum_loreline_lorscript_ttype_TFloat{}
											}

											return Hx_Enum_loreline_lorscript_ttype_TInt{}
										}

									case 3, 5:
										{
											return Hx_Enum_loreline_lorscript_ttype_TFloat{}
										}

									default:
										{
											this.VTable.Hx_Field_unify(t1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
											this.VTable.Hx_Field_unify(t2, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
										}
									}
								}

							case 3:
								{
									switch _g1.Hx_Field_enumIndex() {
									case 2, 3, 5:
										{
											return Hx_Enum_loreline_lorscript_ttype_TFloat{}
										}

									default:
										{
											this.VTable.Hx_Field_unify(t1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
											this.VTable.Hx_Field_unify(t2, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
										}
									}
								}

							case 5:
								{
									switch _g1.Hx_Field_enumIndex() {
									case 2, 3, 5:
										{
											return Hx_Enum_loreline_lorscript_ttype_TFloat{}
										}

									default:
										{
											this.VTable.Hx_Field_unify(t1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
											this.VTable.Hx_Field_unify(t2, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
										}
									}
								}

							default:
								{
									this.VTable.Hx_Field_unify(t1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
									this.VTable.Hx_Field_unify(t2, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
								}
							}
						}
					}

				case "+":
					{
						var _hx_tmp_15 any
						_ = _hx_tmp_15
						_hx_tmp_15 = e1
						var t1 Hx_Enum_loreline_lorscript_ttype
						_ = t1
						t1 = this.VTable.Hx_Field_typeExpr(_hx_tmp_15, Hx_Enum_loreline_lorscript__checker_withtype_WithType{Hx_Enum_loreline_lorscript_ttype_TInt{}})
						var _hx_tmp_16 any
						_ = _hx_tmp_16
						_hx_tmp_16 = e2
						var t2 Hx_Enum_loreline_lorscript_ttype
						_ = t2
						t2 = this.VTable.Hx_Field_typeExpr(_hx_tmp_16, Hx_Enum_loreline_lorscript__checker_withtype_WithType{t1})
						this.VTable.Hx_Field_tryUnify(t1, t2)
						{
							var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp__g_1
							tmp__g_1 = this.VTable.Hx_Field_follow(t1)
							var _g1 Hx_Enum_loreline_lorscript_ttype
							_ = _g1
							_g1 = this.VTable.Hx_Field_follow(t2)
							switch tmp__g_1.Hx_Field_enumIndex() {
							case 2:
								{
									switch _g1.Hx_Field_enumIndex() {
									case 2:
										{
											return Hx_Enum_loreline_lorscript_ttype_TInt{}
										}

									case 3:
										{
											return Hx_Enum_loreline_lorscript_ttype_TFloat{}
										}

									case 5:
										{
											return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
										}

									default:
										{
											var tmp_t1_1 Hx_Enum_loreline_lorscript_ttype
											_ = tmp_t1_1
											tmp_t1_1 = tmp__g_1
											var tmp_t2_1 Hx_Enum_loreline_lorscript_ttype
											_ = tmp_t2_1
											tmp_t2_1 = _g1
											{
												var _hx_tmp_17 bool
												_ = _hx_tmp_17
												if this.VTable.Hx_Field_isString(tmp_t1_1) {
													_hx_tmp_17 = true
												} else {
													_hx_tmp_17 = this.VTable.Hx_Field_isString(tmp_t2_1)
												}

												if _hx_tmp_17 {
													return this.Hx_Field_types.Hx_Field_t_string
												}

												this.VTable.Hx_Field_unify(tmp_t1_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
												this.VTable.Hx_Field_unify(tmp_t2_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
											}
										}
									}
								}

							case 3:
								{
									switch _g1.Hx_Field_enumIndex() {
									case 2, 3:
										{
											return Hx_Enum_loreline_lorscript_ttype_TFloat{}
										}

									case 5:
										{
											return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
										}

									default:
										{
											var tmp_t1_1 Hx_Enum_loreline_lorscript_ttype
											_ = tmp_t1_1
											tmp_t1_1 = tmp__g_1
											var tmp_t2_1 Hx_Enum_loreline_lorscript_ttype
											_ = tmp_t2_1
											tmp_t2_1 = _g1
											{
												var _hx_tmp_18 bool
												_ = _hx_tmp_18
												if this.VTable.Hx_Field_isString(tmp_t1_1) {
													_hx_tmp_18 = true
												} else {
													_hx_tmp_18 = this.VTable.Hx_Field_isString(tmp_t2_1)
												}

												if _hx_tmp_18 {
													return this.Hx_Field_types.Hx_Field_t_string
												}

												this.VTable.Hx_Field_unify(tmp_t1_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
												this.VTable.Hx_Field_unify(tmp_t2_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
											}
										}
									}
								}

							case 5:
								{
									return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
								}

							default:
								{
									if _g1.Hx_Field_enumIndex() == 5 {
										return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
									} else {
										var tmp_t1_1 Hx_Enum_loreline_lorscript_ttype
										_ = tmp_t1_1
										tmp_t1_1 = tmp__g_1
										var tmp_t2_1 Hx_Enum_loreline_lorscript_ttype
										_ = tmp_t2_1
										tmp_t2_1 = _g1
										{
											var _hx_tmp_19 bool
											_ = _hx_tmp_19
											if this.VTable.Hx_Field_isString(tmp_t1_1) {
												_hx_tmp_19 = true
											} else {
												_hx_tmp_19 = this.VTable.Hx_Field_isString(tmp_t2_1)
											}

											if _hx_tmp_19 {
												return this.Hx_Field_types.Hx_Field_t_string
											}

											this.VTable.Hx_Field_unify(tmp_t1_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e1)
											this.VTable.Hx_Field_unify(tmp_t2_1, Hx_Enum_loreline_lorscript_ttype_TFloat{}, e2)
										}
									}
								}
							}
						}
					}

				case "...":
					{
						this.VTable.Hx_Field_typeExprWith(e1, Hx_Enum_loreline_lorscript_ttype_TInt{})
						this.VTable.Hx_Field_typeExprWith(e2, Hx_Enum_loreline_lorscript_ttype_TInt{})
						return this.VTable.Hx_Field_makeIterator(Hx_Enum_loreline_lorscript_ttype_TInt{})
					}

				case "&", "<<", ">>", ">>>", "^", "|":
					{
						this.VTable.Hx_Field_typeExprWith(e1, Hx_Enum_loreline_lorscript_ttype_TInt{})
						this.VTable.Hx_Field_typeExprWith(e2, Hx_Enum_loreline_lorscript_ttype_TInt{})
						return Hx_Enum_loreline_lorscript_ttype_TInt{}
					}

				case "<", "<=", ">", ">=":
					{
						var t1 Hx_Enum_loreline_lorscript_ttype
						_ = t1
						t1 = this.VTable.Hx_Field_typeExpr(e1, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
						var _hx_tmp_20 any
						_ = _hx_tmp_20
						_hx_tmp_20 = e2
						var t2 Hx_Enum_loreline_lorscript_ttype
						_ = t2
						t2 = this.VTable.Hx_Field_typeExpr(_hx_tmp_20, Hx_Enum_loreline_lorscript__checker_withtype_WithType{t1})
						if !this.VTable.Hx_Field_tryUnify(t1, t2) {
							this.VTable.Hx_Field_unify(t2, t1, e2)
						}

						{
							var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp__g_1
							tmp__g_1 = this.VTable.Hx_Field_follow(t1)
							switch tmp__g_1.Hx_Field_enumIndex() {
							case 2, 3, 4:
								{
								}

							case 9:
								{
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c, "name"), ((any)("String"))))) {
										var e Hx_Enum_loreline_lorscript_errordef
										_ = e
										e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot compare ", Hx_Field_loreline_lorscript_checker_typeStr(t1))}

										var tmp_e1_1 *Hx_Obj_loreline_lorscript_error
										_ = tmp_e1_1
										tmp_e1_1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
										if !this.Hx_Field_isCompletion {
											panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e1_1))))
										}
									}
								}

							default:
								{
									var e Hx_Enum_loreline_lorscript_errordef
									_ = e
									e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot compare ", Hx_Field_loreline_lorscript_checker_typeStr(t1))}

									var tmp_e1_1 *Hx_Obj_loreline_lorscript_error
									_ = tmp_e1_1
									tmp_e1_1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
									if !this.Hx_Field_isCompletion {
										panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e1_1))))
									}
								}
							}
						}

						return Hx_Enum_loreline_lorscript_ttype_TBool{}
					}

				case "=":
					{
						if this.Hx_Field_allowDefine {
							var tmp__g_1 Hx_Enum_loreline_lorscript_exprdef
							_ = tmp__g_1
							tmp__g_1 = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if tmp__g_1.Hx_Field_enumIndex() == 1 {
								var i string
								_ = i
								i = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
								var _hx_tmp_21 bool
								_ = _hx_tmp_21
								if !func() func(_hx_param_key string) bool {
									var _hx_closure_src_423 func(_hx_param_key any) bool
									_ = _hx_closure_src_423
									_hx_closure_src_423 = this.Hx_Field_locals.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
									}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
									return func(_hx_param_key string) bool {
										return _hx_closure_src_423(((any)(_hx_param_key)))
									}
								}()(i) {
									_hx_tmp_21 = !func() func(_hx_param_key string) bool {
										var _hx_closure_src_424 func(_hx_param_key any) bool
										_ = _hx_closure_src_424
										_hx_closure_src_424 = this.Hx_Field_globals.VTable.(interface {
											Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
										}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
										return func(_hx_param_key string) bool {
											return _hx_closure_src_424(((any)(_hx_param_key)))
										}
									}()(i)
								} else {
									_hx_tmp_21 = false
								}

								if _hx_tmp_21 {
									var vt Hx_Enum_loreline_lorscript_ttype
									_ = vt
									vt = this.VTable.Hx_Field_typeExpr(e2, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
									func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
										var _hx_closure_src_425 func(_hx_param_key any, _hx_param_value any)
										_ = _hx_closure_src_425
										_hx_closure_src_425 = this.Hx_Field_locals.VTable.(interface {
											Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
										}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
										return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
											_hx_closure_src_425(((any)(_hx_param_key)), ((any)(_hx_param_value)))
										}
									}()(i, vt)
									return vt
								}
							}
						}

						var vt Hx_Enum_loreline_lorscript_ttype
						_ = vt
						var tmp__g_1 Hx_Enum_loreline_lorscript_exprdef
						_ = tmp__g_1
						tmp__g_1 = func() Hx_Enum_loreline_lorscript_exprdef {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_exprdef
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
						}()
						if tmp__g_1.Hx_Field_enumIndex() == 5 {
							var o any
							_ = o
							o = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
							var f string
							_ = f
							f = tmp__g_1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
							vt = this.VTable.Hx_Field_typeField(o, f, e1, true)
						} else {
							vt = this.VTable.Hx_Field_typeExpr(e1, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
						}

						this.VTable.Hx_Field_typeExprWith(e2, vt)
						return vt
					}

				case "&&", "||":
					{
						this.VTable.Hx_Field_typeExprWith(e1, Hx_Enum_loreline_lorscript_ttype_TBool{})
						this.VTable.Hx_Field_typeExprWith(e2, Hx_Enum_loreline_lorscript_ttype_TBool{})
						return Hx_Enum_loreline_lorscript_ttype_TBool{}
					}

				default:
					{
						var _hx_tmp_22 string
						_ = _hx_tmp_22
						_hx_tmp_22 = op
						var hx_cmp_11 HxNullable[int32]
						_ = hx_cmp_11
						hx_cmp_11 = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_22, (int32(utf8.RuneCountInString(op)) - 1)))
						var _hx_tmp_23 bool
						_ = _hx_tmp_23
						if (hx_cmp_11).Valid == true {
							_hx_tmp_23 = ((hx_cmp_11).Value == 61)
						} else {
							_hx_tmp_23 = false
						}

						if _hx_tmp_23 {
							var _hx_tmp_24 string
							_ = _hx_tmp_24
							_hx_tmp_24 = op
							var t Hx_Enum_loreline_lorscript_ttype
							_ = t
							t = this.VTable.Hx_Field_typeExpr(this.VTable.Hx_Field_mk(Hx_Enum_loreline_lorscript_exprdef_EBinop{(Hx_Field_go_haxe_hxstring_substr(_hx_tmp_24, 0, HxNullable[int32]{Value: (int32(utf8.RuneCountInString(op)) - 1), Valid: true})), e1, e2}, expr), withType)
							return this.VTable.Hx_Field_typeExpr(this.VTable.Hx_Field_mk(Hx_Enum_loreline_lorscript_exprdef_EBinop{"=", e1, e2}, expr), withType)
						}

						{
							var e Hx_Enum_loreline_lorscript_errordef
							_ = e
							e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Unsupported operation ", op)}

							var e3 *Hx_Obj_loreline_lorscript_error
							_ = e3
							e3 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
							if !this.Hx_Field_isCompletion {
								panic(Hx_Field_haxe_exception_thrown(((any)(e3))))
							}
						}
					}
				}
			}

		case 7:
			{
				var op string
				_ = op
				op = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_op
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
				{
					var et Hx_Enum_loreline_lorscript_ttype
					_ = et
					et = this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
					switch op {
					case "!":
						{
							this.VTable.Hx_Field_unify(et, Hx_Enum_loreline_lorscript_ttype_TBool{}, e)
							return et
						}

					case "++", "-", "--":
						{
							this.VTable.Hx_Field_unify(et, Hx_Enum_loreline_lorscript_ttype_TInt{}, e)
							return et
						}

					default:
						{
						}
					}
				}
			}

		case 8:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_e
				var params HxArray[any]
				_ = params
				params = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_params
				{
					var _g1 Hx_Enum_loreline_lorscript_exprdef
					_ = _g1
					_g1 = func() Hx_Enum_loreline_lorscript_exprdef {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_exprdef
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
					}()
					if _g1.Hx_Field_enumIndex() == 5 {
						if _g1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f == "bind" {
							var val any
							_ = val
							val = _g1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
							{
								var ft Hx_Enum_loreline_lorscript_ttype
								_ = ft
								ft = this.VTable.Hx_Field_typeExpr(val, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
								if ft.Hx_Field_enumIndex() == 13 {
									var args HxArray[any]
									_ = args
									args = ft.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
									var ret Hx_Enum_loreline_lorscript_ttype
									_ = ret
									ret = ft.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
									var remainArgs HxArray[any]
									_ = remainArgs
									remainArgs = (Hx_Array_Copy[any](HxMakeArrayView[any](args)))
									{
										var tmp__g_1 int32
										_ = tmp__g_1
										tmp__g_1 = 0
										for {
											var _hx_tmp_25 int32
											_ = _hx_tmp_25
											_hx_tmp_25 = tmp__g_1
											if !(_hx_tmp_25 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
												break
											}

											var p any
											_ = p
											p = params.Get(tmp__g_1)
											tmp__g_1++
											var a any
											_ = a
											a = (Hx_Array_Shift[any](HxMakeArrayView[any](remainArgs)).Value)
											if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
												{
													var tmp_e_1 Hx_Enum_loreline_lorscript_errordef
													_ = tmp_e_1
													tmp_e_1 = Hx_Enum_loreline_lorscript_errordef_ECustom{"Too many arguments"}

													var e1 *Hx_Obj_loreline_lorscript_error
													_ = e1
													e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(tmp_e_1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(p, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(p, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(p, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(p, "line")))))
													if !this.Hx_Field_isCompletion {
														panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
													}
												}

												return Hx_Enum_loreline_lorscript_ttype_TFun{HxMakeArray[any]().Dyn(), ret}
											}

											var _hx_tmp_26 any
											_ = _hx_tmp_26
											_hx_tmp_26 = p
											this.VTable.Hx_Field_typeExprWith(_hx_tmp_26, func() Hx_Enum_loreline_lorscript_ttype {
												_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
												if _hx_d == nil {
													var _hx_z Hx_Enum_loreline_lorscript_ttype
													return _hx_z
												}
												return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
											}())
										}
									}

									return Hx_Enum_loreline_lorscript_ttype_TFun{remainArgs, ret}
								}
							}
						}
					}
				}

				var _hx_tmp_27 any
				_ = _hx_tmp_27
				_hx_tmp_27 = e
				var _hx_tmp_28 Hx_Enum_loreline_lorscript__checker_withtype
				_ = _hx_tmp_28
				if Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() == 1 {
					var _hx_tmp_29 Hx_Enum_loreline_lorscript__checker_withtype
					_ = _hx_tmp_29
					if withType.Hx_Field_enumIndex() == 2 {
						var _hx_tmp_30 Hx_Enum_loreline_lorscript__checker_withtype
						_ = _hx_tmp_30
						if withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t.Hx_Field_enumIndex() == 10 {
							_hx_tmp_30 = withType
						} else {
							_hx_tmp_30 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
						}

						_hx_tmp_29 = _hx_tmp_30
					} else {
						_hx_tmp_29 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
					}

					_hx_tmp_28 = _hx_tmp_29
				} else {
					_hx_tmp_28 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
				}

				var ft Hx_Enum_loreline_lorscript_ttype
				_ = ft
				ft = this.VTable.Hx_Field_typeExpr(_hx_tmp_27, _hx_tmp_28)
				var _g2 Hx_Enum_loreline_lorscript_ttype
				_ = _g2
				_g2 = this.VTable.Hx_Field_follow(ft)
				switch _g2.Hx_Field_enumIndex() {
				case 5:
					{
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							for {
								var _hx_tmp_31 int32
								_ = _hx_tmp_31
								_hx_tmp_31 = tmp__g_1
								if !(_hx_tmp_31 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
									break
								}

								var p any
								_ = p
								p = params.Get(tmp__g_1)
								tmp__g_1++
								this.VTable.Hx_Field_typeExpr(p, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
							}
						}

						return this.VTable.Hx_Field_makeMono()
					}

				case 13:
					{
						var args HxArray[any]
						_ = args
						args = _g2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
						var ret Hx_Enum_loreline_lorscript_ttype
						_ = ret
						ret = _g2.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g1 int32
							_ = _g1
							_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(params)
							for tmp__g_1 < _g1 {
								var _hx_tmp_32 int32
								_ = _hx_tmp_32
								_hx_tmp_32 = tmp__g_1
								tmp__g_1 = (tmp__g_1 + 1)
								var i int32
								_ = i
								i = _hx_tmp_32
								var a any
								_ = a
								a = args.Get(i)
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
									{
										var curExpr any
										_ = curExpr
										curExpr = params.Get(i)
										var tmp_e_1 Hx_Enum_loreline_lorscript_errordef
										_ = tmp_e_1
										tmp_e_1 = Hx_Enum_loreline_lorscript_errordef_ECustom{"Too many arguments"}

										var e1 *Hx_Obj_loreline_lorscript_error
										_ = e1
										e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(tmp_e_1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "line")))))
										if !this.Hx_Field_isCompletion {
											panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
										}
									}

									break
								}

								var _hx_tmp_33 any
								_ = _hx_tmp_33
								_hx_tmp_33 = params.Get(i)
								var _hx_tmp_34 Hx_Enum_loreline_lorscript__checker_withtype
								_ = _hx_tmp_34
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
									_hx_tmp_34 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
								} else {
									_hx_tmp_34 = Hx_Enum_loreline_lorscript__checker_withtype_WithType{func() Hx_Enum_loreline_lorscript_ttype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
										if _hx_d == nil {
											var _hx_z Hx_Enum_loreline_lorscript_ttype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
									}()}
								}

								var t Hx_Enum_loreline_lorscript_ttype
								_ = t
								t = this.VTable.Hx_Field_typeExpr(_hx_tmp_33, _hx_tmp_34)
								var _hx_tmp_35 Hx_Enum_loreline_lorscript_ttype
								_ = _hx_tmp_35
								_hx_tmp_35 = t
								this.VTable.Hx_Field_unify(_hx_tmp_35, func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}(), params.Get(i))
							}
						}

						{
							var _g3 int32
							_ = _g3
							_g3 = Hx_Field_go_haxe_hxdynamic_getArrayLength(params)
							var _g4 int32
							_ = _g4
							_g4 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args)
							for _g3 < _g4 {
								var _hx_tmp_36 int32
								_ = _hx_tmp_36
								_hx_tmp_36 = _g3
								_g3 = (_g3 + 1)
								var i int32
								_ = i
								i = _hx_tmp_36
								if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(args.Get(i), "opt")))) {
									var _hx_tmp_37 any
									_ = _hx_tmp_37
									_hx_tmp_37 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(((any)("Missing argument ")), Hx_Field_go_haxe_hxdynamic_getField(args.Get(i), "name")), ((any)(":")))
									var tmp_e_1 Hx_Enum_loreline_lorscript_errordef
									_ = tmp_e_1
									tmp_e_1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_37, ((any)(Hx_Field_loreline_lorscript_checker_typeStr(func() Hx_Enum_loreline_lorscript_ttype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(args.Get(i), "t"))
										if _hx_d == nil {
											var _hx_z Hx_Enum_loreline_lorscript_ttype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
									}())))))}

									var e1 *Hx_Obj_loreline_lorscript_error
									_ = e1
									e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(tmp_e_1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
									if !this.Hx_Field_isCompletion {
										panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
									}
								}
							}
						}

						return ret
					}

				default:
					{
						{
							var e1 Hx_Enum_loreline_lorscript_errordef
							_ = e1
							e1 = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(ft), " cannot be called")}

							var e2 *Hx_Obj_loreline_lorscript_error
							_ = e2
							e2 = Hx_Obj_loreline_lorscript_error_CreateInstance(e1, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "line")))))
							if !this.Hx_Field_isCompletion {
								panic(Hx_Field_haxe_exception_thrown(((any)(e2))))
							}
						}

						return this.VTable.Hx_Field_makeMono()
					}
				}
			}

		case 9:
			{
				var cond any
				_ = cond
				cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_cond
				var e1 any
				_ = e1
				e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
				var e2 any
				_ = e2
				e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
				{
					this.VTable.Hx_Field_typeExprWith(cond, Hx_Enum_loreline_lorscript_ttype_TBool{})
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = this.VTable.Hx_Field_typeExpr(e1, withType)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
						return t1
					}

					var t2 Hx_Enum_loreline_lorscript_ttype
					_ = t2
					t2 = this.VTable.Hx_Field_typeExpr(e2, withType)
					if (withType == Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
						return Hx_Enum_loreline_lorscript_ttype_TVoid{}
					}

					if this.VTable.Hx_Field_tryUnify(t2, t1) {
						return t1
					}

					if this.VTable.Hx_Field_tryUnify(t1, t2) {
						return t2
					}

					this.VTable.Hx_Field_unify(t2, t1, e2)
				}
			}

		case 10:
			{
				var cond any
				_ = cond
				cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_cond
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
				this.VTable.Hx_Field_typeExprWith(cond, Hx_Enum_loreline_lorscript_ttype_TBool{})
				this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
				return Hx_Enum_loreline_lorscript_ttype_TVoid{}
			}

		case 11:
			{
				var v string
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_v
				var it any
				_ = it
				it = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_it
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_e
				var locals *Hx_Obj_haxe_imap
				_ = locals
				locals = this.VTable.Hx_Field_saveLocals()
				var itt Hx_Enum_loreline_lorscript_ttype
				_ = itt
				itt = this.VTable.Hx_Field_typeExpr(it, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
				var vt Hx_Enum_loreline_lorscript_ttype
				_ = vt
				vt = this.VTable.Hx_Field_getIteratorType(it, itt)
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
					var _hx_closure_src_426 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_426
					_hx_closure_src_426 = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
						_hx_closure_src_426(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(v, vt)
				this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
				this.Hx_Field_locals = locals
				return Hx_Enum_loreline_lorscript_ttype_TVoid{}
			}

		case 12, 13:
			{
				return Hx_Enum_loreline_lorscript_ttype_TVoid{}
			}

		case 14:
			{
				var args HxArray[any]
				_ = args
				args = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_args
				var body any
				_ = body
				body = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
				var name HxNullable[string]
				_ = name
				name = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_name
				var ret Hx_Enum_loreline_lorscript_ctype
				_ = ret
				ret = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_ret
				var ft Hx_Enum_loreline_lorscript_ttype
				_ = ft
				var tret Hx_Enum_loreline_lorscript_ttype
				_ = tret
				var targs HxArray[any]
				_ = targs
				if this.Hx_Field_currentFunType != nil {
					{
						var tmp__g_1 Hx_Enum_loreline_lorscript_ttype
						_ = tmp__g_1
						tmp__g_1 = this.Hx_Field_currentFunType
						if tmp__g_1.Hx_Field_enumIndex() == 13 {
							var tmp_args_1 HxArray[any]
							_ = tmp_args_1
							tmp_args_1 = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
							var tmp_ret_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp_ret_1
							tmp_ret_1 = tmp__g_1.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
							{
								ft = this.Hx_Field_currentFunType
								tret = tmp_ret_1
								targs = tmp_args_1
							}
						} else {
							panic(Hx_Field_haxe_exception_thrown(((any)("assert"))))
						}
					}

					this.Hx_Field_currentFunType = nil
				} else {
					var _hx_tmp_38 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_38
					if ret == nil {
						_hx_tmp_38 = this.VTable.Hx_Field_makeMono()
					} else {
						_hx_tmp_38 = this.VTable.Hx_Field_makeType(ret, expr)
					}

					tret = _hx_tmp_38
				}

				var locals *Hx_Obj_haxe_imap
				_ = locals
				locals = this.VTable.Hx_Field_saveLocals()
				var oldRet Hx_Enum_loreline_lorscript_ttype
				_ = oldRet
				oldRet = this.Hx_Field_allowReturn
				var oldGDef bool
				_ = oldGDef
				oldGDef = this.Hx_Field_allowDefine
				var oldHasRet bool
				_ = oldHasRet
				oldHasRet = this.Hx_Field_hasReturn
				this.Hx_Field_allowReturn = tret
				this.Hx_Field_allowDefine = false
				this.Hx_Field_hasReturn = false
				var withArgs HxArray[any]
				_ = withArgs
				var _hx_tmp_39 bool
				_ = _hx_tmp_39
				if (name).Valid != false {
					var _hx_tmp_40 bool
					_ = _hx_tmp_40
					if withType.Hx_Field_enumIndex() == 2 {
						_hx_tmp_40 = (this.VTable.Hx_Field_follow(withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t).Hx_Field_enumIndex() == 13)
					} else {
						_hx_tmp_40 = false
					}

					_hx_tmp_39 = !(_hx_tmp_40)
				} else {
					_hx_tmp_39 = false
				}

				if _hx_tmp_39 {
					var ev Hx_Enum_loreline_lorscript_ttype
					_ = ev
					ev = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
						var _hx_closure_src_427 func(_hx_param_key any) any
						_ = _hx_closure_src_427
						_hx_closure_src_427 = this.Hx_Field_events.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
						return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
							return func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_427(((any)(_hx_param_key))))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}()
						}
					}()((name).Value)
					if ev != nil {
						withType = Hx_Enum_loreline_lorscript__checker_withtype_WithType{ev}
					}
				}

				if withType.Hx_Field_enumIndex() == 2 {
					var _hx_tmp Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp
					_hx_tmp = this.VTable.Hx_Field_follow(withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t)
					if _hx_tmp.Hx_Field_enumIndex() == 13 {
						var tmp_args_1 HxArray[any]
						_ = tmp_args_1
						tmp_args_1 = _hx_tmp.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
						var tmp_ret_1 Hx_Enum_loreline_lorscript_ttype
						_ = tmp_ret_1
						tmp_ret_1 = _hx_tmp.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
						{
							withArgs = tmp_args_1
							this.VTable.Hx_Field_unify(tret, tmp_ret_1, expr)
						}
					}
				}

				if targs == nil {
					targs = this.VTable.Hx_Field_typeArgs(args, expr)
				}

				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 int32
					_ = _g2
					_g2 = Hx_Field_go_haxe_hxdynamic_getArrayLength(targs)
					for _g1 < _g2 {
						var _hx_tmp_41 int32
						_ = _hx_tmp_41
						_hx_tmp_41 = _g1
						_g1 = (_g1 + 1)
						var i int32
						_ = i
						i = _hx_tmp_41
						var a any
						_ = a
						a = targs.Get(i)
						if withArgs != nil {
							var _hx_tmp_42 int32
							_ = _hx_tmp_42
							_hx_tmp_42 = i
							if _hx_tmp_42 < Hx_Field_go_haxe_hxdynamic_getArrayLength(withArgs) {
								var _hx_tmp_43 Hx_Enum_loreline_lorscript_ttype
								_ = _hx_tmp_43
								_hx_tmp_43 = func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(withArgs.Get(i), "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}()
								this.VTable.Hx_Field_unify(_hx_tmp_43, func() Hx_Enum_loreline_lorscript_ttype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ttype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
								}(), expr)
							} else {
								var e Hx_Enum_loreline_lorscript_errordef
								_ = e
								e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(((any)("Extra argument ")), Hx_Field_go_haxe_hxdynamic_getField(a, "name")))}

								var e1 *Hx_Obj_loreline_lorscript_error
								_ = e1
								e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
								if !this.Hx_Field_isCompletion {
									panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
								}
							}
						}

						var _hx_tmp_44 string
						_ = _hx_tmp_44
						_hx_tmp_44 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "name"))
						func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
							var _hx_closure_src_428 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_428
							_hx_closure_src_428 = this.Hx_Field_locals.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								_hx_closure_src_428(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(_hx_tmp_44, func() Hx_Enum_loreline_lorscript_ttype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ttype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
						}())
					}
				}

				var _hx_tmp_45 bool
				_ = _hx_tmp_45
				if withArgs != nil {
					var _hx_tmp_46 int32
					_ = _hx_tmp_46
					_hx_tmp_46 = Hx_Field_go_haxe_hxdynamic_getArrayLength(targs)
					var _hx_tmp_47 int32
					_ = _hx_tmp_47
					_hx_tmp_47 = _hx_tmp_46
					_hx_tmp_45 = (_hx_tmp_47 < Hx_Field_go_haxe_hxdynamic_getArrayLength(withArgs))
				} else {
					_hx_tmp_45 = false
				}

				if _hx_tmp_45 {
					var _hx_tmp_48 int32
					_ = _hx_tmp_48
					_hx_tmp_48 = Hx_Field_go_haxe_hxdynamic_getArrayLength(withArgs)
					var msg string
					_ = msg
					msg = Hx_Field_go_haxe_hxdynamic_stringConcat(("Missing " + Hx_Field_go_haxe_hxdynamic_toString((_hx_tmp_48 - Hx_Field_go_haxe_hxdynamic_getArrayLength(targs)))), " arguments (")
					var tmp__g_1 HxArray[string]
					_ = tmp__g_1
					tmp__g_1 = HxMakeArray[string]()
					{
						var _g1 int32
						_ = _g1
						_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(targs)
						var _g2 int32
						_ = _g2
						_g2 = Hx_Field_go_haxe_hxdynamic_getArrayLength(withArgs)
						for _g1 < _g2 {
							var _hx_tmp_49 int32
							_ = _hx_tmp_49
							_hx_tmp_49 = _g1
							_g1 = (_g1 + 1)
							var i int32
							_ = i
							i = _hx_tmp_49
							var _hx_tmp_50 HxArray[string]
							_ = _hx_tmp_50
							_hx_tmp_50 = tmp__g_1
							_ = (Hx_Array_Push[string](_hx_tmp_50, Hx_Field_loreline_lorscript_checker_typeStr(func() Hx_Enum_loreline_lorscript_ttype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(withArgs.Get(i), "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ttype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
							}())))
						}
					}

					var _hx_tmp_51 string
					_ = _hx_tmp_51
					_hx_tmp_51 = msg
					var e Hx_Enum_loreline_lorscript_errordef
					_ = e
					e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_51, (Hx_Array_Join[string](tmp__g_1, ","))), ")")}

					var e1 *Hx_Obj_loreline_lorscript_error
					_ = e1
					e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
					}
				}

				this.VTable.Hx_Field_typeExpr(body, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
				var _hx_tmp_52 bool
				_ = _hx_tmp_52
				if !this.Hx_Field_hasReturn {
					_hx_tmp_52 = !this.VTable.Hx_Field_tryUnify(tret, Hx_Enum_loreline_lorscript_ttype_TVoid{})
				} else {
					_hx_tmp_52 = false
				}

				if _hx_tmp_52 {
					var e Hx_Enum_loreline_lorscript_errordef
					_ = e
					e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Missing return ", Hx_Field_loreline_lorscript_checker_typeStr(tret))}

					var e1 *Hx_Obj_loreline_lorscript_error
					_ = e1
					e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
					}
				}

				this.Hx_Field_allowDefine = oldGDef
				this.Hx_Field_allowReturn = oldRet
				this.Hx_Field_hasReturn = oldHasRet
				this.Hx_Field_locals = locals
				if ft == nil {
					ft = Hx_Enum_loreline_lorscript_ttype_TFun{targs, tret}

					if (name).Valid != false {
						func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
							var _hx_closure_src_429 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_429
							_hx_closure_src_429 = locals.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
								_hx_closure_src_429(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()((name).Value, ft)
					}
				}

				return ft
			}

		case 15:
			{
				var v any
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
				var _hx_tmp_53 Hx_Enum_loreline_lorscript_ttype
				_ = _hx_tmp_53
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, nil))) {
					_hx_tmp_53 = Hx_Enum_loreline_lorscript_ttype_TVoid{}
				} else {
					var _hx_tmp_54 any
					_ = _hx_tmp_54
					_hx_tmp_54 = v
					var _hx_tmp_55 Hx_Enum_loreline_lorscript__checker_withtype
					_ = _hx_tmp_55
					if this.Hx_Field_allowReturn == nil {
						_hx_tmp_55 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
					} else {
						_hx_tmp_55 = Hx_Enum_loreline_lorscript__checker_withtype_WithType{this.Hx_Field_allowReturn}
					}

					_hx_tmp_53 = this.VTable.Hx_Field_typeExpr(_hx_tmp_54, _hx_tmp_55)
				}

				var et Hx_Enum_loreline_lorscript_ttype
				_ = et
				et = _hx_tmp_53
				this.Hx_Field_hasReturn = true
				if this.Hx_Field_allowReturn == nil {
					var e Hx_Enum_loreline_lorscript_errordef
					_ = e
					e = Hx_Enum_loreline_lorscript_errordef_ECustom{"Return not allowed here"}

					var e1 *Hx_Obj_loreline_lorscript_error
					_ = e1
					e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
					if !this.Hx_Field_isCompletion {
						panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
					}
				} else {
					var _hx_tmp_56 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_56
					_hx_tmp_56 = et
					var _hx_tmp_57 Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp_57
					_hx_tmp_57 = this.Hx_Field_allowReturn
					var _hx_tmp_58 any
					_ = _hx_tmp_58
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, nil))) {
						_hx_tmp_58 = expr
					} else {
						_hx_tmp_58 = v
					}

					this.VTable.Hx_Field_unify(_hx_tmp_56, _hx_tmp_57, _hx_tmp_58)
				}

				return this.VTable.Hx_Field_makeMono()
			}

		case 16:
			{
				var a any
				_ = a
				a = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
				var index any
				_ = index
				index = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
				{
					this.VTable.Hx_Field_typeExprWith(index, Hx_Enum_loreline_lorscript_ttype_TInt{})
					var at Hx_Enum_loreline_lorscript_ttype
					_ = at
					at = this.VTable.Hx_Field_typeExpr(a, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
					{
						var _g1 Hx_Enum_loreline_lorscript_ttype
						_ = _g1
						_g1 = this.VTable.Hx_Field_follow(at)
						if _g1.Hx_Field_enumIndex() == 9 {
							var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
							_ = _gargs
							_gargs = _g1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(_g1.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c, "name"), ((any)("Array"))))) {
								if _gargs.Len() == 1 {
									var et Hx_Enum_loreline_lorscript_ttype
									_ = et
									et = _gargs.Get(0)
									return et
								} else {
									var e Hx_Enum_loreline_lorscript_errordef
									_ = e
									e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(at), " is not an Array")}

									var e1 *Hx_Obj_loreline_lorscript_error
									_ = e1
									e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "line")))))
									if !this.Hx_Field_isCompletion {
										panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
									}
								}
							} else {
								var e Hx_Enum_loreline_lorscript_errordef
								_ = e
								e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(at), " is not an Array")}

								var e1 *Hx_Obj_loreline_lorscript_error
								_ = e1
								e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "line")))))
								if !this.Hx_Field_isCompletion {
									panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
								}
							}
						} else {
							var e Hx_Enum_loreline_lorscript_errordef
							_ = e
							e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_loreline_lorscript_checker_typeStr(at), " is not an Array")}

							var e1 *Hx_Obj_loreline_lorscript_error
							_ = e1
							e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(a, "line")))))
							if !this.Hx_Field_isCompletion {
								panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
							}
						}
					}
				}
			}

		case 17:
			{
				var el HxArray[any]
				_ = el
				el = _g.(Hx_Enum_loreline_lorscript_exprdef_EArrayDecl).Hx_Field_e
				var et Hx_Enum_loreline_lorscript_ttype
				_ = et
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_59 int32
						_ = _hx_tmp_59
						_hx_tmp_59 = _g1
						if !(_hx_tmp_59 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
							break
						}

						var v any
						_ = v
						v = el.Get(_g1)
						_g1++
						var _hx_tmp_60 any
						_ = _hx_tmp_60
						_hx_tmp_60 = v
						var _hx_tmp_61 Hx_Enum_loreline_lorscript__checker_withtype
						_ = _hx_tmp_61
						if et == nil {
							_hx_tmp_61 = Hx_Enum_loreline_lorscript__checker_withtype_Value{}
						} else {
							_hx_tmp_61 = Hx_Enum_loreline_lorscript__checker_withtype_WithType{et}
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = this.VTable.Hx_Field_typeExpr(_hx_tmp_60, _hx_tmp_61)
						if et == nil {
							et = t
						} else {
							if !this.VTable.Hx_Field_tryUnify(t, et) {
								if this.VTable.Hx_Field_tryUnify(et, t) {
									et = t
								} else {
									this.VTable.Hx_Field_unify(t, et, v)
								}
							}
						}
					}
				}

				if et == nil {
					et = this.VTable.Hx_Field_makeMono()
				}

				return this.Hx_Field_types.VTable.Hx_Field_getType("Array", HxMakeArray[Hx_Enum_loreline_lorscript_ttype](et))
			}

		case 18:
			{
				var cl string
				_ = cl
				cl = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_cl
				var params HxArray[any]
				_ = params
				params = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_params
			}

		case 19:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EThrow).Hx_Field_e
				this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
				return this.VTable.Hx_Field_makeMono()
			}

		case 20:
			{
				var etry any
				_ = etry
				etry = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_e
				var v string
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_v
				var et Hx_Enum_loreline_lorscript_ctype
				_ = et
				et = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_t
				var ecatch any
				_ = ecatch
				ecatch = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
				var vt Hx_Enum_loreline_lorscript_ttype
				_ = vt
				vt = this.VTable.Hx_Field_typeExpr(etry, withType)
				var old Hx_Enum_loreline_lorscript_ttype
				_ = old
				old = func() func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
					var _hx_closure_src_430 func(_hx_param_key any) any
					_ = _hx_closure_src_430
					_hx_closure_src_430 = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
					return func(_hx_param_key string) Hx_Enum_loreline_lorscript_ttype {
						return func() Hx_Enum_loreline_lorscript_ttype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_430(((any)(_hx_param_key))))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ttype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
						}()
					}
				}()(v)
				var _hx_tmp_62 string
				_ = _hx_tmp_62
				_hx_tmp_62 = v
				func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
					var _hx_closure_src_431 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_431
					_hx_closure_src_431 = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
						_hx_closure_src_431(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_62, this.VTable.Hx_Field_makeType(et, ecatch))
				var ct Hx_Enum_loreline_lorscript_ttype
				_ = ct
				ct = this.VTable.Hx_Field_typeExpr(ecatch, withType)
				if old != nil {
					func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
						var _hx_closure_src_432 func(_hx_param_key any, _hx_param_value any)
						_ = _hx_closure_src_432
						_hx_closure_src_432 = this.Hx_Field_locals.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
						return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_lorscript_ttype) {
							_hx_closure_src_432(((any)(_hx_param_key)), ((any)(_hx_param_value)))
						}
					}()(v, old)
				} else {
					func() func(_hx_param_key string) bool {
						var _hx_closure_src_433 func(_hx_param_key any) bool
						_ = _hx_closure_src_433
						_hx_closure_src_433 = this.Hx_Field_locals.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_remove
						return func(_hx_param_key string) bool {
							return _hx_closure_src_433(((any)(_hx_param_key)))
						}
					}()(v)
				}

				if (withType == Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
					return Hx_Enum_loreline_lorscript_ttype_TVoid{}
				}

				if this.VTable.Hx_Field_tryUnify(vt, ct) {
					return ct
				}

				this.VTable.Hx_Field_unify(ct, vt, ecatch)
				return vt
			}

		case 21:
			{
				var fl HxArray[any]
				_ = fl
				fl = _g.(Hx_Enum_loreline_lorscript_exprdef_EObject).Hx_Field_fl
				if withType.Hx_Field_enumIndex() == 2 {
					var _hx_tmp Hx_Enum_loreline_lorscript_ttype
					_ = _hx_tmp
					_hx_tmp = this.VTable.Hx_Field_follow(withType.(Hx_Enum_loreline_lorscript__checker_withtype_WithType).Hx_Field_t)
					if _hx_tmp.Hx_Field_enumIndex() == 14 {
						var tfields HxArray[any]
						_ = tfields
						tfields = _hx_tmp.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
						if Hx_Field_go_haxe_hxdynamic_getArrayLength(tfields) > 0 {
							var tmp__g_1 *Hx_Obj_haxe_imap
							_ = tmp__g_1
							tmp__g_1 = (func() *Hx_Obj_haxe_imap {
								_hx_icast_89 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
								if _hx_icast_89 == nil {
									return nil
								}
								Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_89, "Hx_Obj_haxe_imap")
								return &Hx_Obj_haxe_imap{VTable: _hx_icast_89.VTable.(Hx_Obj_VTable_haxe_imap)}
							})()
							{
								var _g1 int32
								_ = _g1
								_g1 = 0
								for {
									var _hx_tmp_63 int32
									_ = _hx_tmp_63
									_hx_tmp_63 = _g1
									if !(_hx_tmp_63 < Hx_Field_go_haxe_hxdynamic_getArrayLength(tfields)) {
										break
									}

									var f any
									_ = f
									f = tfields.Get(_g1)
									_g1++
									tmp__g_1.VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), f)
								}
							}

							var _hx_reserved_map *Hx_Obj_haxe_imap
							_ = _hx_reserved_map
							_hx_reserved_map = tmp__g_1
							var _g2 HxArray[any]
							_ = _g2
							_g2 = HxMakeArray[any]().Dyn()
							{
								var _g3 int32
								_ = _g3
								_g3 = 0
								for {
									var _hx_tmp_64 int32
									_ = _hx_tmp_64
									_hx_tmp_64 = _g3
									if !(_hx_tmp_64 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
										break
									}

									var f any
									_ = f
									f = fl.Get(_g3)
									_g3++
									var ft any
									_ = ft
									ft = _hx_reserved_map.VTable.Hx_Field_get(Hx_Field_go_haxe_hxdynamic_getField(f, "name"))
									var ft1 Hx_Enum_loreline_lorscript_ttype
									_ = ft1
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(ft, nil))) {
										{
											var curExpr any
											_ = curExpr
											curExpr = Hx_Field_go_haxe_hxdynamic_getField(f, "e")
											var e Hx_Enum_loreline_lorscript_errordef
											_ = e
											e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(((any)("Extra field ")), Hx_Field_go_haxe_hxdynamic_getField(f, "name")))}

											var e1 *Hx_Obj_loreline_lorscript_error
											_ = e1
											e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(curExpr, "line")))))
											if !this.Hx_Field_isCompletion {
												panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
											}
										}

										ft1 = Hx_Enum_loreline_lorscript_ttype_TDynamic{}
									} else {
										ft1 = func() Hx_Enum_loreline_lorscript_ttype {
											_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(ft, "t"))
											if _hx_d == nil {
												var _hx_z Hx_Enum_loreline_lorscript_ttype
												return _hx_z
											}
											return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
										}()
									}

									var _hx_tmp_65 HxArray[any]
									_ = _hx_tmp_65
									_hx_tmp_65 = HxMakeArrayView[any](_g2)
									var _hx_tmp_66 any
									_ = _hx_tmp_66
									_hx_tmp_66 = ((any)(this.VTable.Hx_Field_typeExprWith(Hx_Field_go_haxe_hxdynamic_getField(f, "e"), ft1)))
									_ = (Hx_Array_Push[any](_hx_tmp_65, any(map[string]any{"t": _hx_tmp_66, "opt": ((any)(false)), "name": Hx_Field_go_haxe_hxdynamic_getField(f, "name")})))
								}
							}

							return Hx_Enum_loreline_lorscript_ttype_TAnon{_g2}
						} else {
							var tmp__g_1 HxArray[any]
							_ = tmp__g_1
							tmp__g_1 = HxMakeArray[any]().Dyn()
							{
								var _g1 int32
								_ = _g1
								_g1 = 0
								for {
									var _hx_tmp_67 int32
									_ = _hx_tmp_67
									_hx_tmp_67 = _g1
									if !(_hx_tmp_67 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
										break
									}

									var f any
									_ = f
									f = fl.Get(_g1)
									_g1++
									var _hx_tmp_68 HxArray[any]
									_ = _hx_tmp_68
									_hx_tmp_68 = HxMakeArrayView[any](tmp__g_1)
									var _hx_tmp_69 any
									_ = _hx_tmp_69
									_hx_tmp_69 = ((any)(this.VTable.Hx_Field_typeExpr(Hx_Field_go_haxe_hxdynamic_getField(f, "e"), Hx_Enum_loreline_lorscript__checker_withtype_Value{})))
									_ = (Hx_Array_Push[any](_hx_tmp_68, any(map[string]any{"t": _hx_tmp_69, "opt": ((any)(false)), "name": Hx_Field_go_haxe_hxdynamic_getField(f, "name")})))
								}
							}

							return Hx_Enum_loreline_lorscript_ttype_TAnon{tmp__g_1}
						}
					} else {
						var tmp__g_1 HxArray[any]
						_ = tmp__g_1
						tmp__g_1 = HxMakeArray[any]().Dyn()
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							for {
								var _hx_tmp_70 int32
								_ = _hx_tmp_70
								_hx_tmp_70 = _g1
								if !(_hx_tmp_70 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
									break
								}

								var f any
								_ = f
								f = fl.Get(_g1)
								_g1++
								var _hx_tmp_71 HxArray[any]
								_ = _hx_tmp_71
								_hx_tmp_71 = HxMakeArrayView[any](tmp__g_1)
								var _hx_tmp_72 any
								_ = _hx_tmp_72
								_hx_tmp_72 = ((any)(this.VTable.Hx_Field_typeExpr(Hx_Field_go_haxe_hxdynamic_getField(f, "e"), Hx_Enum_loreline_lorscript__checker_withtype_Value{})))
								_ = (Hx_Array_Push[any](_hx_tmp_71, any(map[string]any{"t": _hx_tmp_72, "opt": ((any)(false)), "name": Hx_Field_go_haxe_hxdynamic_getField(f, "name")})))
							}
						}

						return Hx_Enum_loreline_lorscript_ttype_TAnon{tmp__g_1}
					}
				} else {
					var tmp__g_1 HxArray[any]
					_ = tmp__g_1
					tmp__g_1 = HxMakeArray[any]().Dyn()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_73 int32
							_ = _hx_tmp_73
							_hx_tmp_73 = _g1
							if !(_hx_tmp_73 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
								break
							}

							var f any
							_ = f
							f = fl.Get(_g1)
							_g1++
							var _hx_tmp_74 HxArray[any]
							_ = _hx_tmp_74
							_hx_tmp_74 = HxMakeArrayView[any](tmp__g_1)
							var _hx_tmp_75 any
							_ = _hx_tmp_75
							_hx_tmp_75 = ((any)(this.VTable.Hx_Field_typeExpr(Hx_Field_go_haxe_hxdynamic_getField(f, "e"), Hx_Enum_loreline_lorscript__checker_withtype_Value{})))
							_ = (Hx_Array_Push[any](_hx_tmp_74, any(map[string]any{"t": _hx_tmp_75, "opt": ((any)(false)), "name": Hx_Field_go_haxe_hxdynamic_getField(f, "name")})))
						}
					}

					return Hx_Enum_loreline_lorscript_ttype_TAnon{tmp__g_1}
				}
			}

		case 22:
			{
				var cond any
				_ = cond
				cond = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
				var e1 any
				_ = e1
				e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
				var e2 any
				_ = e2
				e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
				{
					this.VTable.Hx_Field_typeExprWith(cond, Hx_Enum_loreline_lorscript_ttype_TBool{})
					var t1 Hx_Enum_loreline_lorscript_ttype
					_ = t1
					t1 = this.VTable.Hx_Field_typeExpr(e1, withType)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
						return t1
					}

					var t2 Hx_Enum_loreline_lorscript_ttype
					_ = t2
					t2 = this.VTable.Hx_Field_typeExpr(e2, withType)
					if (withType == Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
						return Hx_Enum_loreline_lorscript_ttype_TVoid{}
					}

					if this.VTable.Hx_Field_tryUnify(t2, t1) {
						return t1
					}

					if this.VTable.Hx_Field_tryUnify(t1, t2) {
						return t2
					}

					this.VTable.Hx_Field_unify(t2, t1, e2)
				}
			}

		case 23:
			{
				var value any
				_ = value
				value = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_e
				var cases HxArray[any]
				_ = cases
				cases = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_cases
				var defaultExpr any
				_ = defaultExpr
				defaultExpr = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_defaultExpr
				var tmin Hx_Enum_loreline_lorscript_ttype
				_ = tmin
				var vt Hx_Enum_loreline_lorscript_ttype
				_ = vt
				vt = this.VTable.Hx_Field_typeExpr(value, Hx_Enum_loreline_lorscript__checker_withtype_Value{})
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_76 int32
						_ = _hx_tmp_76
						_hx_tmp_76 = _g1
						if !(_hx_tmp_76 < Hx_Field_go_haxe_hxdynamic_getArrayLength(cases)) {
							break
						}

						var c any
						_ = c
						c = cases.Get(_g1)
						_g1++
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g2 HxArray[any]
							_ = _g2
							_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "values"))
							for {
								var _hx_tmp_77 int32
								_ = _hx_tmp_77
								_hx_tmp_77 = tmp__g_1
								if !(_hx_tmp_77 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
									break
								}

								var v any
								_ = v
								v = _g2.Get(tmp__g_1)
								tmp__g_1++
								var _hx_tmp_78 any
								_ = _hx_tmp_78
								_hx_tmp_78 = v
								var ct Hx_Enum_loreline_lorscript_ttype
								_ = ct
								ct = this.VTable.Hx_Field_typeExpr(_hx_tmp_78, Hx_Enum_loreline_lorscript__checker_withtype_WithType{vt})
								this.VTable.Hx_Field_unify(ct, vt, v)
							}
						}

						var et Hx_Enum_loreline_lorscript_ttype
						_ = et
						et = this.VTable.Hx_Field_typeExpr(Hx_Field_go_haxe_hxdynamic_getField(c, "expr"), withType)
						{
							var p any
							_ = p
							p = Hx_Field_go_haxe_hxdynamic_getField(c, "expr")
							if (withType != Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
								if tmin == nil {
									tmin = et
								} else {
									if !_gthis.VTable.Hx_Field_tryUnify(et, tmin) {
										_gthis.VTable.Hx_Field_unify(tmin, et, p)
										tmin = et
									}
								}
							}
						}
					}
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(defaultExpr, nil))) {
					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = this.VTable.Hx_Field_typeExpr(defaultExpr, withType)
					if (withType != Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
						if tmin == nil {
							tmin = t
						} else {
							if !_gthis.VTable.Hx_Field_tryUnify(t, tmin) {
								_gthis.VTable.Hx_Field_unify(tmin, t, defaultExpr)
								tmin = t
							}
						}
					}
				}

				if (withType == Hx_Enum_loreline_lorscript__checker_withtype_NoValue{}) {
					return Hx_Enum_loreline_lorscript_ttype_TVoid{}
				} else {
					if tmin == nil {
						return this.VTable.Hx_Field_makeMono()
					} else {
						return tmin
					}
				}
			}

		case 24:
			{
				var cond any
				_ = cond
				cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_cond
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
				this.VTable.Hx_Field_typeExprWith(cond, Hx_Enum_loreline_lorscript_ttype_TBool{})
				this.VTable.Hx_Field_typeExpr(e, Hx_Enum_loreline_lorscript__checker_withtype_NoValue{})
				return Hx_Enum_loreline_lorscript_ttype_TVoid{}
			}

		case 25:
			{
				var m string
				_ = m
				m = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_name
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
				if (m == ":untyped") && this.Hx_Field_allowUntypedMeta {
					return this.VTable.Hx_Field_makeMono()
				}

				return this.VTable.Hx_Field_typeExpr(e, withType)
			}

		case 26:
			{
				var v any
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e
				var t Hx_Enum_loreline_lorscript_ctype
				_ = t
				t = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_t
				var ct Hx_Enum_loreline_lorscript_ttype
				_ = ct
				ct = this.VTable.Hx_Field_makeType(t, expr)
				var _hx_tmp_79 any
				_ = _hx_tmp_79
				_hx_tmp_79 = v
				var vt Hx_Enum_loreline_lorscript_ttype
				_ = vt
				vt = this.VTable.Hx_Field_typeExpr(_hx_tmp_79, Hx_Enum_loreline_lorscript__checker_withtype_WithType{ct})
				this.VTable.Hx_Field_unify(vt, ct, v)
				return ct
			}

		}
	}

	{
		var e Hx_Enum_loreline_lorscript_errordef
		_ = e
		e = Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Don't know how to type ", Hx_Field_type_enumConstructor(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(expr, "e"))
			if _hx_d == nil {
				var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
				return _hx_z
			}
			return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
		}()))}

		var e1 *Hx_Obj_loreline_lorscript_error
		_ = e1
		e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(expr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "line")))))
		if !this.Hx_Field_isCompletion {
			panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
		}
	}

	return Hx_Enum_loreline_lorscript_ttype_TDynamic{}
}

func (this *Hx_Obj_loreline_lorscript_checker) Hx_Field_getIteratorType(it any, itt Hx_Enum_loreline_lorscript_ttype) Hx_Enum_loreline_lorscript_ttype {
	{
		var _g Hx_Enum_loreline_lorscript_ttype
		_ = _g
		_g = this.VTable.Hx_Field_follow(itt)
		if _g.Hx_Field_enumIndex() == 9 {
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = _gargs
			_gargs = _g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(_g.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c, "name"), ((any)("Array"))))) {
				if _gargs.Len() == 1 {
					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = _gargs.Get(0)
					return t
				}
			}
		}
	}

	var _hx_tmp_0 Hx_Enum_loreline_lorscript_ttype
	_ = _hx_tmp_0
	_hx_tmp_0 = itt
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = it
	var ft Hx_Enum_loreline_lorscript_ttype
	_ = ft
	ft = this.VTable.Hx_Field_getField(_hx_tmp_0, "iterator", _hx_tmp_1, nil)
	if ft == nil {
		if itt.Hx_Field_enumIndex() == 12 {
			var a any
			_ = a
			a = itt.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = itt.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			var _hx_tmp_2 Hx_Enum_loreline_lorscript_ttype
			_ = _hx_tmp_2
			_hx_tmp_2 = func() Hx_Enum_loreline_lorscript_ttype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ttype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
			}()
			var at Hx_Enum_loreline_lorscript_ttype
			_ = at
			at = this.VTable.Hx_Field_apply(_hx_tmp_2, HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ttype](Hx_Field_go_haxe_hxdynamic_getField(a, "params")), args)
			return this.VTable.Hx_Field_getIteratorType(it, at)
		}
	}

	if ft != nil {
		if ft == nil {
			ft = nil
		} else {
			if ft.Hx_Field_enumIndex() == 13 {
				if Hx_Field_go_haxe_hxdynamic_getArrayLength(ft.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args) == 0 {
					var ret Hx_Enum_loreline_lorscript_ttype
					_ = ret
					ret = ft.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
					ft = ret
				} else {
					ft = nil
				}
			} else {
				ft = nil
			}
		}
	}

	var t Hx_Enum_loreline_lorscript_ttype
	_ = t
	t = this.VTable.Hx_Field_makeMono()
	var iter Hx_Enum_loreline_lorscript_ttype
	_ = iter
	iter = this.VTable.Hx_Field_makeIterator(t)
	var _hx_tmp_3 Hx_Enum_loreline_lorscript_ttype
	_ = _hx_tmp_3
	if ft != nil {
		_hx_tmp_3 = ft
	} else {
		_hx_tmp_3 = itt
	}

	this.VTable.Hx_Field_unify(_hx_tmp_3, iter, it)
	return t
}

func Hx_Field_loreline_lorscript_checker_typeStr(t Hx_Enum_loreline_lorscript_ttype) string {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				return "Unknown"
			} else {
				return Hx_Field_loreline_lorscript_checker_typeStr(func() Hx_Enum_loreline_lorscript_ttype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ttype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
				}())
			}
		}

	case 6:
		{
			var name string
			_ = name
			name = t.(Hx_Enum_loreline_lorscript_ttype_TParam).Hx_Field_name
			return name
		}

	case 7:
		{
			var name string
			_ = name
			name = t.(Hx_Enum_loreline_lorscript_ttype_TUnresolved).Hx_Field_name
			return Hx_Field_go_haxe_hxdynamic_stringConcat("?", name)
		}

	case 8:
		{
			var t1 Hx_Enum_loreline_lorscript_ttype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Null<", Hx_Field_loreline_lorscript_checker_typeStr(t1)), ">")
		}

	case 9:
		{
			var c any
			_ = c
			c = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_c
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			var c1 string
			_ = c1
			c1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(c, "name"))
			var tmp string
			_ = tmp
			if args.Len() == 0 {
				tmp = ""
			} else {
				var _g HxArray[string]
				_ = _g
				_g = HxMakeArray[string]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_1 HxArray[string]
						_ = _hx_tmp_1
						_hx_tmp_1 = _g
						_ = (Hx_Array_Push[string](_hx_tmp_1, Hx_Field_loreline_lorscript_checker_typeStr(t)))
					}
				}

				tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("<", (Hx_Array_Join[string](_g, ","))), ">")
			}

			return Hx_Field_go_haxe_hxdynamic_stringConcat(c1, tmp)
		}

	case 10:
		{
			var e any
			_ = e
			e = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_e
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
			var e1 string
			_ = e1
			e1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(e, "name"))
			var tmp string
			_ = tmp
			if args.Len() == 0 {
				tmp = ""
			} else {
				var _g HxArray[string]
				_ = _g
				_g = HxMakeArray[string]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = _g1
						if !(_hx_tmp_2 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_3 HxArray[string]
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						_ = (Hx_Array_Push[string](_hx_tmp_3, Hx_Field_loreline_lorscript_checker_typeStr(t)))
					}
				}

				tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("<", (Hx_Array_Join[string](_g, ","))), ">")
			}

			return Hx_Field_go_haxe_hxdynamic_stringConcat(e1, tmp)
		}

	case 11:
		{
			var t1 any
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_t
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t1, "name"), ((any)("loreline.lorscript.TypeCheck"))))) {
				return Hx_Field_loreline_lorscript_checker_typeStr(args.Get(1))
			} else {
				var t string
				_ = t
				t = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(t1, "name"))
				var tmp string
				_ = tmp
				if args.Len() == 0 {
					tmp = ""
				} else {
					var _g HxArray[string]
					_ = _g
					_g = HxMakeArray[string]()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = _g1
							if !(_hx_tmp_4 < args.Len()) {
								break
							}

							var tmp_t_1 Hx_Enum_loreline_lorscript_ttype
							_ = tmp_t_1
							tmp_t_1 = args.Get(_g1)
							_g1++
							var _hx_tmp_5 HxArray[string]
							_ = _hx_tmp_5
							_hx_tmp_5 = _g
							_ = (Hx_Array_Push[string](_hx_tmp_5, Hx_Field_loreline_lorscript_checker_typeStr(tmp_t_1)))
						}
					}

					tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("<", (Hx_Array_Join[string](_g, ","))), ">")
				}

				return Hx_Field_go_haxe_hxdynamic_stringConcat(t, tmp)
			}
		}

	case 12:
		{
			var a any
			_ = a
			a = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_a
			var args HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			var a1 string
			_ = a1
			a1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "name"))
			var tmp string
			_ = tmp
			if args.Len() == 0 {
				tmp = ""
			} else {
				var _g HxArray[string]
				_ = _g
				_g = HxMakeArray[string]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = _g1
						if !(_hx_tmp_6 < args.Len()) {
							break
						}

						var t Hx_Enum_loreline_lorscript_ttype
						_ = t
						t = args.Get(_g1)
						_g1++
						var _hx_tmp_7 HxArray[string]
						_ = _hx_tmp_7
						_hx_tmp_7 = _g
						_ = (Hx_Array_Push[string](_hx_tmp_7, Hx_Field_loreline_lorscript_checker_typeStr(t)))
					}
				}

				tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("<", (Hx_Array_Join[string](_g, ","))), ">")
			}

			return Hx_Field_go_haxe_hxdynamic_stringConcat(a1, tmp)
		}

	case 13:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ttype
			_ = ret
			ret = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			var _g HxArray[string]
			_ = _g
			_g = HxMakeArray[string]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _g1
					if !(_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
						break
					}

					var a any
					_ = a
					a = args.Get(_g1)
					_g1++
					var _hx_tmp_9 HxArray[string]
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					var _hx_tmp_10 string
					_ = _hx_tmp_10
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(a, "opt")))) {
						_hx_tmp_10 = "?"
					} else {
						_hx_tmp_10 = ""
					}

					var _hx_tmp_11 string
					_ = _hx_tmp_11
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(a, "name"), ((any)(""))))) {
						_hx_tmp_11 = ""
					} else {
						_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(a, "name"), ((any)(":"))))
					}

					_ = (Hx_Array_Push[string](_hx_tmp_9, Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat((_hx_tmp_10), (_hx_tmp_11)), Hx_Field_loreline_lorscript_checker_typeStr(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}()))))
				}
			}

			var _hx_tmp_12 string
			_ = _hx_tmp_12
			_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("(", (Hx_Array_Join[string](_g, ", "))), ") -> ")
			return Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_12, Hx_Field_loreline_lorscript_checker_typeStr(ret))
		}

	case 14:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			var _g HxArray[string]
			_ = _g
			_g = HxMakeArray[string]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = _g1
					if !(_hx_tmp_13 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
						break
					}

					var f any
					_ = f
					f = fields.Get(_g1)
					_g1++
					var _hx_tmp_14 HxArray[string]
					_ = _hx_tmp_14
					_hx_tmp_14 = _g
					var _hx_tmp_15 string
					_ = _hx_tmp_15
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(f, "opt")))) {
						_hx_tmp_15 = "?"
					} else {
						_hx_tmp_15 = ""
					}

					_ = (Hx_Array_Push[string](_hx_tmp_14, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(((any)((_hx_tmp_15))), Hx_Field_go_haxe_hxdynamic_getField(f, "name")), ((any)(":"))), ((any)(Hx_Field_loreline_lorscript_checker_typeStr(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())))))))
				}
			}

			return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("{", (Hx_Array_Join[string](_g, ", "))), "}")
		}

	default:
		{
			var _hx_tmp_16 string
			_ = _hx_tmp_16
			_hx_tmp_16 = Hx_Field_type_enumConstructor(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(t)))
			return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_16, 1, HxNullable[int32]{}))
		}
	}
}

func Hx_Field_loreline_lorscript_checker_typeIter(t Hx_Enum_loreline_lorscript_ttype, callb func(_hx_carg0 Hx_Enum_loreline_lorscript_ttype)) {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var r any
			_ = r
			r = t.(Hx_Enum_loreline_lorscript_ttype_TMono).Hx_Field_r
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(r, "r"), ((any)(nil))))) {
				callb(func() Hx_Enum_loreline_lorscript_ttype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(r, "r"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_ttype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
				}())
			}
		}

	case 8:
		{
			var t1 Hx_Enum_loreline_lorscript_ttype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ttype_TNull).Hx_Field_t
			callb(t1)
		}

	case 9:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TInst).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = _g
					if !(_hx_tmp_0 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					callb(t)
				}
			}
		}

	case 10:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TEnum).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = _g
					if !(_hx_tmp_1 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					callb(t)
				}
			}
		}

	case 11:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TType).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g
					if !(_hx_tmp_2 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					callb(t)
				}
			}
		}

	case 12:
		{
			var tl HxArray[Hx_Enum_loreline_lorscript_ttype]
			_ = tl
			tl = t.(Hx_Enum_loreline_lorscript_ttype_TAbstract).Hx_Field_args
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					if !(_hx_tmp_3 < tl.Len()) {
						break
					}

					var t Hx_Enum_loreline_lorscript_ttype
					_ = t
					t = tl.Get(_g)
					_g++
					callb(t)
				}
			}
		}

	case 13:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_args
			var ret Hx_Enum_loreline_lorscript_ttype
			_ = ret
			ret = t.(Hx_Enum_loreline_lorscript_ttype_TFun).Hx_Field_ret
			{
				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g
						if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
							break
						}

						var t any
						_ = t
						t = args.Get(_g)
						_g++
						callb(func() Hx_Enum_loreline_lorscript_ttype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ttype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
						}())
					}
				}

				callb(ret)
			}
		}

	case 14:
		{
			var fl HxArray[any]
			_ = fl
			fl = t.(Hx_Enum_loreline_lorscript_ttype_TAnon).Hx_Field_fields
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = _g
					if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
						break
					}

					var f any
					_ = f
					f = fl.Get(_g)
					_g++
					callb(func() Hx_Enum_loreline_lorscript_ttype {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_ttype
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_ttype)
					}())
				}
			}
		}

	case 15:
		{
			var f func() Hx_Enum_loreline_lorscript_ttype
			_ = f
			f = t.(Hx_Enum_loreline_lorscript_ttype_TLazy).Hx_Field_f
			callb(f())
		}

	default:
		{
		}
	}
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_ctype]()
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_typetree]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_ctype]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_ttype]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[func()]()
	hxRegisterArrayConstructor[func() Hx_Enum_loreline_lorscript_ttype]()
	hxRegisterArrayConstructor[string]()
}
