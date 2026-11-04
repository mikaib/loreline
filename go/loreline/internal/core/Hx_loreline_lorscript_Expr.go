package core

type Hx_Enum_loreline_lorscript_const interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_const()
}

type Hx_Enum_loreline_lorscript_const_CString struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_const_CString) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_const_CString) M_Hx_Enum_loreline_lorscript_const() {}
func (this Hx_Enum_loreline_lorscript_const_CString) Hx_Field_enumIndex() int32           { return 2 }
func (this Hx_Enum_loreline_lorscript_const_CString) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_const_RTTI
}
func (this Hx_Enum_loreline_lorscript_const_CString) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_const_CString) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_const_CInt struct {
	Hx_Field_v int32
}

func (this Hx_Enum_loreline_lorscript_const_CInt) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_const_CInt) M_Hx_Enum_loreline_lorscript_const() {}
func (this Hx_Enum_loreline_lorscript_const_CInt) Hx_Field_enumIndex() int32           { return 0 }
func (this Hx_Enum_loreline_lorscript_const_CInt) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_const_RTTI
}
func (this Hx_Enum_loreline_lorscript_const_CInt) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_loreline_lorscript_const_CInt) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_const_CFloat struct {
	Hx_Field_f float64
}

func (this Hx_Enum_loreline_lorscript_const_CFloat) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_const_CFloat) M_Hx_Enum_loreline_lorscript_const() {}
func (this Hx_Enum_loreline_lorscript_const_CFloat) Hx_Field_enumIndex() int32           { return 1 }
func (this Hx_Enum_loreline_lorscript_const_CFloat) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_const_RTTI
}
func (this Hx_Enum_loreline_lorscript_const_CFloat) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_f))
}
func (this Hx_Enum_loreline_lorscript_const_CFloat) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_f)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_const_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.Const",
	HxMakeArray[string]("CInt", "CFloat", "CString"),
	HxMakeArray[int32](1, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_const_CInt{((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(0))))}
		case 1:
			return Hx_Enum_loreline_lorscript_const_CFloat{((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(params.Get_Dyn(0))))}
		case 2:
			return Hx_Enum_loreline_lorscript_const_CString{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_loreline_lorscript_expr = any

type Hx_Enum_loreline_lorscript_exprdef interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_exprdef()
}

type Hx_Enum_loreline_lorscript_exprdef_EWhile struct {
	Hx_Field_cond any
	Hx_Field_e    any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) Hx_Field_enumIndex() int32             { return 10 }
func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_cond), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EWhile) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_cond)
	case 1:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EVar struct {
	Hx_Field_n string
	Hx_Field_t Hx_Enum_loreline_lorscript_ctype
	Hx_Field_e any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EVar) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EVar) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EVar) Hx_Field_enumIndex() int32             { return 2 }
func (this Hx_Enum_loreline_lorscript_exprdef_EVar) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EVar) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_n), any(this.Hx_Field_t), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EVar) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_n)
	case 1:
		return any(this.Hx_Field_t)
	case 2:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EUnop struct {
	Hx_Field_op     string
	Hx_Field_prefix bool
	Hx_Field_e      any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) Hx_Field_enumIndex() int32             { return 7 }
func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_op), any(this.Hx_Field_prefix), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EUnop) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_op)
	case 1:
		return any(this.Hx_Field_prefix)
	case 2:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ETry struct {
	Hx_Field_e      any
	Hx_Field_v      string
	Hx_Field_t      Hx_Enum_loreline_lorscript_ctype
	Hx_Field_ecatch any
}

func (this Hx_Enum_loreline_lorscript_exprdef_ETry) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETry) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ETry) Hx_Field_enumIndex() int32             { return 20 }
func (this Hx_Enum_loreline_lorscript_exprdef_ETry) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETry) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_v), any(this.Hx_Field_t), any(this.Hx_Field_ecatch))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETry) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_v)
	case 2:
		return any(this.Hx_Field_t)
	case 3:
		return any(this.Hx_Field_ecatch)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EThrow struct {
	Hx_Field_e any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) Hx_Field_enumIndex() int32             { return 19 }
func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EThrow) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ETernary struct {
	Hx_Field_cond any
	Hx_Field_e1   any
	Hx_Field_e2   any
}

func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) Hx_Field_enumIndex() int32             { return 22 }
func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_cond), any(this.Hx_Field_e1), any(this.Hx_Field_e2))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ETernary) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_cond)
	case 1:
		return any(this.Hx_Field_e1)
	case 2:
		return any(this.Hx_Field_e2)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ESwitch struct {
	Hx_Field_e           any
	Hx_Field_cases       HxArray[any]
	Hx_Field_defaultExpr any
}

func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) Hx_Field_enumIndex() int32             { return 23 }
func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_cases), any(this.Hx_Field_defaultExpr))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ESwitch) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_cases)
	case 2:
		return any(this.Hx_Field_defaultExpr)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EReturn struct {
	Hx_Field_e any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) Hx_Field_enumIndex() int32             { return 15 }
func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EReturn) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EParent struct {
	Hx_Field_e any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EParent) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EParent) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EParent) Hx_Field_enumIndex() int32             { return 3 }
func (this Hx_Enum_loreline_lorscript_exprdef_EParent) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EParent) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EParent) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EObject struct {
	Hx_Field_fl HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_exprdef_EObject) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EObject) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EObject) Hx_Field_enumIndex() int32             { return 21 }
func (this Hx_Enum_loreline_lorscript_exprdef_EObject) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EObject) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_fl))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EObject) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_fl)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ENew struct {
	Hx_Field_cl     string
	Hx_Field_params HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_exprdef_ENew) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ENew) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ENew) Hx_Field_enumIndex() int32             { return 18 }
func (this Hx_Enum_loreline_lorscript_exprdef_ENew) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ENew) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_cl), any(this.Hx_Field_params))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ENew) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_cl)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EMeta struct {
	Hx_Field_name string
	Hx_Field_args HxArray[any]
	Hx_Field_e    any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) Hx_Field_enumIndex() int32             { return 25 }
func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_args), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EMeta) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_args)
	case 2:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EIf struct {
	Hx_Field_cond any
	Hx_Field_e1   any
	Hx_Field_e2   any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EIf) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIf) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EIf) Hx_Field_enumIndex() int32             { return 9 }
func (this Hx_Enum_loreline_lorscript_exprdef_EIf) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIf) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_cond), any(this.Hx_Field_e1), any(this.Hx_Field_e2))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIf) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_cond)
	case 1:
		return any(this.Hx_Field_e1)
	case 2:
		return any(this.Hx_Field_e2)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EIdent struct {
	Hx_Field_v string
}

func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) Hx_Field_enumIndex() int32             { return 1 }
func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EIdent) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EFunction struct {
	Hx_Field_args HxArray[any]
	Hx_Field_e    any
	Hx_Field_name HxNullable[string]
	Hx_Field_ret  Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) Hx_Field_enumIndex() int32             { return 14 }
func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_args), any(this.Hx_Field_e), any(this.Hx_Field_name), any(this.Hx_Field_ret))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFunction) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_args)
	case 1:
		return any(this.Hx_Field_e)
	case 2:
		return any(this.Hx_Field_name)
	case 3:
		return any(this.Hx_Field_ret)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EFor struct {
	Hx_Field_v  string
	Hx_Field_it any
	Hx_Field_e  any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EFor) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFor) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EFor) Hx_Field_enumIndex() int32             { return 11 }
func (this Hx_Enum_loreline_lorscript_exprdef_EFor) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFor) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v), any(this.Hx_Field_it), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EFor) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	case 1:
		return any(this.Hx_Field_it)
	case 2:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EField struct {
	Hx_Field_e any
	Hx_Field_f string
}

func (this Hx_Enum_loreline_lorscript_exprdef_EField) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EField) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EField) Hx_Field_enumIndex() int32             { return 5 }
func (this Hx_Enum_loreline_lorscript_exprdef_EField) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EField) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_f))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EField) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_f)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EDoWhile struct {
	Hx_Field_cond any
	Hx_Field_e    any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) Hx_Field_enumIndex() int32             { return 24 }
func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_cond), any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EDoWhile) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_cond)
	case 1:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EContinue struct{}

func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) Hx_Field_enumIndex() int32             { return 13 }
func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_exprdef_EContinue) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EConst struct {
	Hx_Field_c Hx_Enum_loreline_lorscript_const
}

func (this Hx_Enum_loreline_lorscript_exprdef_EConst) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EConst) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EConst) Hx_Field_enumIndex() int32             { return 0 }
func (this Hx_Enum_loreline_lorscript_exprdef_EConst) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EConst) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EConst) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ECheckType struct {
	Hx_Field_e any
	Hx_Field_t Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) Hx_Field_enumIndex() int32             { return 26 }
func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECheckType) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_ECall struct {
	Hx_Field_e      any
	Hx_Field_params HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_exprdef_ECall) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECall) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_ECall) Hx_Field_enumIndex() int32             { return 8 }
func (this Hx_Enum_loreline_lorscript_exprdef_ECall) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECall) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_params))
}
func (this Hx_Enum_loreline_lorscript_exprdef_ECall) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EBreak struct{}

func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) Hx_Field_enumIndex() int32             { return 12 }
func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBreak) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EBlock struct {
	Hx_Field_e HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) Hx_Field_enumIndex() int32             { return 4 }
func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBlock) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EBinop struct {
	Hx_Field_op string
	Hx_Field_e1 any
	Hx_Field_e2 any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) Hx_Field_enumIndex() int32             { return 6 }
func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_op), any(this.Hx_Field_e1), any(this.Hx_Field_e2))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EBinop) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_op)
	case 1:
		return any(this.Hx_Field_e1)
	case 2:
		return any(this.Hx_Field_e2)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EArrayDecl struct {
	Hx_Field_e HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) Hx_Field_enumIndex() int32             { return 17 }
func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArrayDecl) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_exprdef_EArray struct {
	Hx_Field_e     any
	Hx_Field_index any
}

func (this Hx_Enum_loreline_lorscript_exprdef_EArray) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArray) M_Hx_Enum_loreline_lorscript_exprdef() {}
func (this Hx_Enum_loreline_lorscript_exprdef_EArray) Hx_Field_enumIndex() int32             { return 16 }
func (this Hx_Enum_loreline_lorscript_exprdef_EArray) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_exprdef_RTTI
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArray) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e), any(this.Hx_Field_index))
}
func (this Hx_Enum_loreline_lorscript_exprdef_EArray) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	case 1:
		return any(this.Hx_Field_index)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_exprdef_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.ExprDef",
	HxMakeArray[string]("EConst", "EIdent", "EVar", "EParent", "EBlock", "EField", "EBinop", "EUnop", "ECall", "EIf", "EWhile", "EFor", "EBreak", "EContinue", "EFunction", "EReturn", "EArray", "EArrayDecl", "ENew", "EThrow", "ETry", "EObject", "ETernary", "ESwitch", "EDoWhile", "EMeta", "ECheckType"),
	HxMakeArray[int32](1, 1, 3, 1, 1, 2, 3, 3, 2, 3, 2, 3, 0, 0, 4, 1, 2, 1, 2, 1, 4, 1, 3, 3, 2, 3, 2),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_exprdef_EConst{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_const)}
		case 1:
			return Hx_Enum_loreline_lorscript_exprdef_EIdent{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 2:
			return Hx_Enum_loreline_lorscript_exprdef_EVar{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), params.Get_Dyn(1).(Hx_Enum_loreline_lorscript_ctype), params.Get_Dyn(2).(any)}
		case 3:
			return Hx_Enum_loreline_lorscript_exprdef_EParent{params.Get_Dyn(0).(any)}
		case 4:
			return Hx_Enum_loreline_lorscript_exprdef_EBlock{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 5:
			return Hx_Enum_loreline_lorscript_exprdef_EField{params.Get_Dyn(0).(any), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1))))}
		case 6:
			return Hx_Enum_loreline_lorscript_exprdef_EBinop{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), params.Get_Dyn(1).(any), params.Get_Dyn(2).(any)}
		case 7:
			return Hx_Enum_loreline_lorscript_exprdef_EUnop{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(1)))), params.Get_Dyn(2).(any)}
		case 8:
			return Hx_Enum_loreline_lorscript_exprdef_ECall{params.Get_Dyn(0).(any), HxAnyToTypedArray[any](params.Get_Dyn(1))}
		case 9:
			return Hx_Enum_loreline_lorscript_exprdef_EIf{params.Get_Dyn(0).(any), params.Get_Dyn(1).(any), params.Get_Dyn(2).(any)}
		case 10:
			return Hx_Enum_loreline_lorscript_exprdef_EWhile{params.Get_Dyn(0).(any), params.Get_Dyn(1).(any)}
		case 11:
			return Hx_Enum_loreline_lorscript_exprdef_EFor{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), params.Get_Dyn(1).(any), params.Get_Dyn(2).(any)}
		case 12:
			return Hx_Enum_loreline_lorscript_exprdef_EBreak{}
		case 13:
			return Hx_Enum_loreline_lorscript_exprdef_EContinue{}
		case 14:
			return Hx_Enum_loreline_lorscript_exprdef_EFunction{HxAnyToTypedArray[any](params.Get_Dyn(0)), params.Get_Dyn(1).(any), HxNullable[string]{Value: ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(2)))), Valid: params.Get_Dyn(2) != nil}, params.Get_Dyn(3).(Hx_Enum_loreline_lorscript_ctype)}
		case 15:
			return Hx_Enum_loreline_lorscript_exprdef_EReturn{params.Get_Dyn(0).(any)}
		case 16:
			return Hx_Enum_loreline_lorscript_exprdef_EArray{params.Get_Dyn(0).(any), params.Get_Dyn(1).(any)}
		case 17:
			return Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 18:
			return Hx_Enum_loreline_lorscript_exprdef_ENew{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), HxAnyToTypedArray[any](params.Get_Dyn(1))}
		case 19:
			return Hx_Enum_loreline_lorscript_exprdef_EThrow{params.Get_Dyn(0).(any)}
		case 20:
			return Hx_Enum_loreline_lorscript_exprdef_ETry{params.Get_Dyn(0).(any), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1)))), params.Get_Dyn(2).(Hx_Enum_loreline_lorscript_ctype), params.Get_Dyn(3).(any)}
		case 21:
			return Hx_Enum_loreline_lorscript_exprdef_EObject{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 22:
			return Hx_Enum_loreline_lorscript_exprdef_ETernary{params.Get_Dyn(0).(any), params.Get_Dyn(1).(any), params.Get_Dyn(2).(any)}
		case 23:
			return Hx_Enum_loreline_lorscript_exprdef_ESwitch{params.Get_Dyn(0).(any), HxAnyToTypedArray[any](params.Get_Dyn(1)), params.Get_Dyn(2).(any)}
		case 24:
			return Hx_Enum_loreline_lorscript_exprdef_EDoWhile{params.Get_Dyn(0).(any), params.Get_Dyn(1).(any)}
		case 25:
			return Hx_Enum_loreline_lorscript_exprdef_EMeta{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), HxAnyToTypedArray[any](params.Get_Dyn(1)), params.Get_Dyn(2).(any)}
		case 26:
			return Hx_Enum_loreline_lorscript_exprdef_ECheckType{params.Get_Dyn(0).(any), params.Get_Dyn(1).(Hx_Enum_loreline_lorscript_ctype)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_loreline_lorscript_argument = any

type Hx_Typedef_loreline_lorscript_metadata = HxArray[any]

type Hx_Enum_loreline_lorscript_ctype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_ctype()
}

type Hx_Enum_loreline_lorscript_ctype_CTPath struct {
	Hx_Field_path   HxArray[string]
	Hx_Field_params HxArray[Hx_Enum_loreline_lorscript_ctype]
}

func (this Hx_Enum_loreline_lorscript_ctype_CTPath) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTPath) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTPath) Hx_Field_enumIndex() int32           { return 0 }
func (this Hx_Enum_loreline_lorscript_ctype_CTPath) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTPath) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_path), any(this.Hx_Field_params))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTPath) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_path)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctype_CTParent struct {
	Hx_Field_t Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_ctype_CTParent) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTParent) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTParent) Hx_Field_enumIndex() int32           { return 3 }
func (this Hx_Enum_loreline_lorscript_ctype_CTParent) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTParent) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTParent) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctype_CTOpt struct {
	Hx_Field_t Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) Hx_Field_enumIndex() int32           { return 4 }
func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTOpt) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctype_CTNamed struct {
	Hx_Field_n string
	Hx_Field_t Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) Hx_Field_enumIndex() int32           { return 5 }
func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_n), any(this.Hx_Field_t))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTNamed) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_n)
	case 1:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctype_CTFun struct {
	Hx_Field_args HxArray[Hx_Enum_loreline_lorscript_ctype]
	Hx_Field_ret  Hx_Enum_loreline_lorscript_ctype
}

func (this Hx_Enum_loreline_lorscript_ctype_CTFun) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTFun) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTFun) Hx_Field_enumIndex() int32           { return 1 }
func (this Hx_Enum_loreline_lorscript_ctype_CTFun) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTFun) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_args), any(this.Hx_Field_ret))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTFun) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_args)
	case 1:
		return any(this.Hx_Field_ret)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_ctype_CTAnon struct {
	Hx_Field_fields HxArray[any]
}

func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) M_Hx_Enum_loreline_lorscript_ctype() {}
func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) Hx_Field_enumIndex() int32           { return 2 }
func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_ctype_RTTI
}
func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_fields))
}
func (this Hx_Enum_loreline_lorscript_ctype_CTAnon) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_fields)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_ctype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.CType",
	HxMakeArray[string]("CTPath", "CTFun", "CTAnon", "CTParent", "CTOpt", "CTNamed"),
	HxMakeArray[int32](2, 2, 1, 1, 1, 2),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_ctype_CTPath{HxAnyToTypedArray[string](params.Get_Dyn(0)), params.Get_Dyn(1).(HxArray[Hx_Enum_loreline_lorscript_ctype])}
		case 1:
			return Hx_Enum_loreline_lorscript_ctype_CTFun{HxAnyToTypedArray[Hx_Enum_loreline_lorscript_ctype](params.Get_Dyn(0)), params.Get_Dyn(1).(Hx_Enum_loreline_lorscript_ctype)}
		case 2:
			return Hx_Enum_loreline_lorscript_ctype_CTAnon{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 3:
			return Hx_Enum_loreline_lorscript_ctype_CTParent{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_ctype)}
		case 4:
			return Hx_Enum_loreline_lorscript_ctype_CTOpt{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_ctype)}
		case 5:
			return Hx_Enum_loreline_lorscript_ctype_CTNamed{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), params.Get_Dyn(1).(Hx_Enum_loreline_lorscript_ctype)}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_error() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_error() {
}

type Hx_Obj_VTable_loreline_lorscript_error interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_error struct {
	VTable          Hx_Obj_VTable_loreline_lorscript_error `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_e      Hx_Enum_loreline_lorscript_errordef
	Hx_Field_pmin   int32
	Hx_Field_pmax   int32
	Hx_Field_origin string
	Hx_Field_line   int32
}

func Hx_Obj_loreline_lorscript_error_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_error {
	obj := &Hx_Obj_loreline_lorscript_error{}
	obj.VTable = obj
	obj.Hx_Field_e = HxDefault[Hx_Enum_loreline_lorscript_errordef]()
	obj.Hx_Field_pmin = HxDefault[int32]()
	obj.Hx_Field_pmax = HxDefault[int32]()
	obj.Hx_Field_origin = HxDefault[string]()
	obj.Hx_Field_line = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_lorscript_error_CreateInstance(e Hx_Enum_loreline_lorscript_errordef, pmin int32, pmax int32, origin string, line int32) *Hx_Obj_loreline_lorscript_error {
	obj := Hx_Obj_loreline_lorscript_error_CreateEmptyInstance()
	obj.Hx_New(e, pmin, pmax, origin, line)
	return obj
}

func (this *Hx_Obj_loreline_lorscript_error) Hx_New(e Hx_Enum_loreline_lorscript_errordef, pmin int32, pmax int32, origin string, line int32) {
	{
		this.Hx_Field_e = e
		this.Hx_Field_pmin = pmin
		this.Hx_Field_pmax = pmax
		this.Hx_Field_origin = origin
		this.Hx_Field_line = line
	}

}

var Hx_Obj_loreline_lorscript_error_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Error",
	HxMakeArray[string](),
	HxMakeArray[string]("e", "pmin", "pmax", "origin", "line", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_error_CreateInstance(HxConvert[Hx_Enum_loreline_lorscript_errordef](params.Get(0)), HxConvert[int32](params.Get(1)), HxConvert[int32](params.Get(2)), HxConvert[string](params.Get(3)), HxConvert[int32](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_lorscript_error_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_error_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_error) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_error_RTTI
}

func (this *Hx_Obj_loreline_lorscript_error) Hx_As_Hx_Obj_loreline_lorscript_error() *Hx_Obj_loreline_lorscript_error {
	return this
}

func (this *Hx_Obj_loreline_lorscript_error) Hx_Field_toString() string {
	return Hx_Field_loreline_lorscript_printer_errorToString(this)
}

type Hx_Enum_loreline_lorscript_errordef interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_errordef()
}

type Hx_Enum_loreline_lorscript_errordef_EUnterminatedString struct{}

func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) Hx_Field_enumIndex() int32 {
	return 2
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedString) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment struct{}

func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) Hx_Field_enumIndex() int32 {
	return 3
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EUnknownVariable struct {
	Hx_Field_v string
}

func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) Hx_Field_enumIndex() int32 { return 5 }
func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnknownVariable) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EUnexpected struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) M_Hx_Enum_loreline_lorscript_errordef() {}
func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) Hx_Field_enumIndex() int32              { return 1 }
func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_errordef_EUnexpected) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor struct {
	Hx_Field_msg string
}

func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) Hx_Field_enumIndex() int32 {
	return 4
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_msg))
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_msg)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EInvalidOp struct {
	Hx_Field_op string
}

func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) M_Hx_Enum_loreline_lorscript_errordef() {}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) Hx_Field_enumIndex() int32              { return 7 }
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_op))
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidOp) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_op)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EInvalidIterator struct {
	Hx_Field_v string
}

func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) Hx_Field_enumIndex() int32 { return 6 }
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidIterator) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EInvalidChar struct {
	Hx_Field_c int32
}

func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) Hx_Field_enumIndex() int32 { return 0 }
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidChar) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_EInvalidAccess struct {
	Hx_Field_f string
}

func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) M_Hx_Enum_loreline_lorscript_errordef() {
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) Hx_Field_enumIndex() int32 { return 8 }
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_f))
}
func (this Hx_Enum_loreline_lorscript_errordef_EInvalidAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_f)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_errordef_ECustom struct {
	Hx_Field_msg string
}

func (this Hx_Enum_loreline_lorscript_errordef_ECustom) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_errordef_ECustom) M_Hx_Enum_loreline_lorscript_errordef() {}
func (this Hx_Enum_loreline_lorscript_errordef_ECustom) Hx_Field_enumIndex() int32              { return 9 }
func (this Hx_Enum_loreline_lorscript_errordef_ECustom) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_errordef_RTTI
}
func (this Hx_Enum_loreline_lorscript_errordef_ECustom) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_msg))
}
func (this Hx_Enum_loreline_lorscript_errordef_ECustom) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_msg)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_errordef_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.ErrorDef",
	HxMakeArray[string]("EInvalidChar", "EUnexpected", "EUnterminatedString", "EUnterminatedComment", "EInvalidPreprocessor", "EUnknownVariable", "EInvalidIterator", "EInvalidOp", "EInvalidAccess", "ECustom"),
	HxMakeArray[int32](1, 1, 0, 0, 1, 1, 1, 1, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_errordef_EInvalidChar{((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(0))))}
		case 1:
			return Hx_Enum_loreline_lorscript_errordef_EUnexpected{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 2:
			return Hx_Enum_loreline_lorscript_errordef_EUnterminatedString{}
		case 3:
			return Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment{}
		case 4:
			return Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 5:
			return Hx_Enum_loreline_lorscript_errordef_EUnknownVariable{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 6:
			return Hx_Enum_loreline_lorscript_errordef_EInvalidIterator{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 7:
			return Hx_Enum_loreline_lorscript_errordef_EInvalidOp{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 8:
			return Hx_Enum_loreline_lorscript_errordef_EInvalidAccess{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 9:
			return Hx_Enum_loreline_lorscript_errordef_ECustom{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Enum_loreline_lorscript_moduledecl interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_moduledecl()
}

type Hx_Enum_loreline_lorscript_moduledecl_DTypedef struct {
	Hx_Field_c any
}

func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) M_Hx_Enum_loreline_lorscript_moduledecl() {
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) Hx_Field_enumIndex() int32 { return 3 }
func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_moduledecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DTypedef) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_moduledecl_DPackage struct {
	Hx_Field_path HxArray[string]
}

func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) M_Hx_Enum_loreline_lorscript_moduledecl() {
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) Hx_Field_enumIndex() int32 { return 0 }
func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_moduledecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_path))
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DPackage) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_path)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_moduledecl_DImport struct {
	Hx_Field_path       HxArray[string]
	Hx_Field_everything HxNullable[bool]
}

func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) M_Hx_Enum_loreline_lorscript_moduledecl() {}
func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) Hx_Field_enumIndex() int32                { return 1 }
func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_moduledecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_path), any(this.Hx_Field_everything))
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DImport) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_path)
	case 1:
		return any(this.Hx_Field_everything)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_moduledecl_DClass struct {
	Hx_Field_c any
}

func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) M_Hx_Enum_loreline_lorscript_moduledecl() {}
func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) Hx_Field_enumIndex() int32                { return 2 }
func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_moduledecl_RTTI
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_moduledecl_DClass) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_moduledecl_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.ModuleDecl",
	HxMakeArray[string]("DPackage", "DImport", "DClass", "DTypedef"),
	HxMakeArray[int32](1, 2, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_moduledecl_DPackage{HxAnyToTypedArray[string](params.Get_Dyn(0))}
		case 1:
			return Hx_Enum_loreline_lorscript_moduledecl_DImport{HxAnyToTypedArray[string](params.Get_Dyn(0)), HxNullable[bool]{Value: ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(1)))), Valid: params.Get_Dyn(1) != nil}}
		case 2:
			return Hx_Enum_loreline_lorscript_moduledecl_DClass{params.Get_Dyn(0).(any)}
		case 3:
			return Hx_Enum_loreline_lorscript_moduledecl_DTypedef{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_loreline_lorscript_moduletype = any

type Hx_Typedef_loreline_lorscript_classdecl = any

type Hx_Typedef_loreline_lorscript_typedecl = any

type Hx_Typedef_loreline_lorscript_fielddecl = any

type Hx_Enum_loreline_lorscript_fieldaccess interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_fieldaccess()
}

type Hx_Enum_loreline_lorscript_fieldaccess_AStatic struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) Hx_Field_enumIndex() int32 { return 4 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AStatic) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldaccess_APublic struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) Hx_Field_enumIndex() int32 { return 0 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APublic) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldaccess_APrivate struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) Hx_Field_enumIndex() int32 { return 1 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_APrivate) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldaccess_AOverride struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) Hx_Field_enumIndex() int32 { return 3 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AOverride) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldaccess_AMacro struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) Hx_Field_enumIndex() int32 { return 5 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AMacro) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldaccess_AInline struct{}

func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) M_Hx_Enum_loreline_lorscript_fieldaccess() {
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) Hx_Field_enumIndex() int32 { return 2 }
func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldaccess_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_fieldaccess_AInline) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_fieldaccess_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.FieldAccess",
	HxMakeArray[string]("APublic", "APrivate", "AInline", "AOverride", "AStatic", "AMacro"),
	HxMakeArray[int32](0, 0, 0, 0, 0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_fieldaccess_APublic{}
		case 1:
			return Hx_Enum_loreline_lorscript_fieldaccess_APrivate{}
		case 2:
			return Hx_Enum_loreline_lorscript_fieldaccess_AInline{}
		case 3:
			return Hx_Enum_loreline_lorscript_fieldaccess_AOverride{}
		case 4:
			return Hx_Enum_loreline_lorscript_fieldaccess_AStatic{}
		case 5:
			return Hx_Enum_loreline_lorscript_fieldaccess_AMacro{}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Enum_loreline_lorscript_fieldkind interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_fieldkind()
}

type Hx_Enum_loreline_lorscript_fieldkind_KVar struct {
	Hx_Field_v any
}

func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) M_Hx_Enum_loreline_lorscript_fieldkind() {}
func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldkind_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KVar) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_fieldkind_KFunction struct {
	Hx_Field_f any
}

func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) M_Hx_Enum_loreline_lorscript_fieldkind() {}
func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_fieldkind_RTTI
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_f))
}
func (this Hx_Enum_loreline_lorscript_fieldkind_KFunction) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_f)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_fieldkind_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.FieldKind",
	HxMakeArray[string]("KFunction", "KVar"),
	HxMakeArray[int32](1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_fieldkind_KFunction{params.Get_Dyn(0).(any)}
		case 1:
			return Hx_Enum_loreline_lorscript_fieldkind_KVar{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_loreline_lorscript_functiondecl = any

type Hx_Typedef_loreline_lorscript_vardecl = any

func init() {
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_ctype]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
