package core

import "math"

type Hx_Enum_loreline_lorscript__interp_stop interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript__interp_stop()
}

type Hx_Enum_loreline_lorscript__interp_stop_SReturn struct{}

func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) M_Hx_Enum_loreline_lorscript__interp_stop() {
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) Hx_Field_enumIndex() int32 { return 2 }
func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__interp_stop_RTTI
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SReturn) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript__interp_stop_SContinue struct{}

func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) M_Hx_Enum_loreline_lorscript__interp_stop() {
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) Hx_Field_enumIndex() int32 { return 1 }
func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__interp_stop_RTTI
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SContinue) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript__interp_stop_SBreak struct{}

func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) M_Hx_Enum_loreline_lorscript__interp_stop() {
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) Hx_Field_enumIndex() int32 { return 0 }
func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript__interp_stop_RTTI
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript__interp_stop_SBreak) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript__interp_stop_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript._Interp.Stop",
	HxMakeArray[string]("SBreak", "SContinue", "SReturn"),
	HxMakeArray[int32](0, 0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript__interp_stop_SBreak{}
		case 1:
			return Hx_Enum_loreline_lorscript__interp_stop_SContinue{}
		case 2:
			return Hx_Enum_loreline_lorscript__interp_stop_SReturn{}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_interp() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_interp() {
}

type Hx_Obj_VTable_loreline_lorscript_interp interface {
	Hx_Field_assignOp(op string, fop func(_hx_carg0 any, _hx_carg1 any) any)
	Hx_Field_restore(old int32)
	Hx_Field_rethrow(e any)
	Hx_Field_fcall(o any, f string, args HxArray[any]) any
	Hx_Field_posInfos() any
	Hx_Field_execute(expr any) any
	Hx_Field_doWhileLoop(econd any, e any)
	Hx_Field_whileLoop(econd any, e any)
	Hx_Field_isMap(o any) bool
	Hx_Field_resetVariables()
	Hx_Field_evalAssignOp(op string, fop func(_hx_carg0 any, _hx_carg1 any) any, e1 any, e2 any) any
	Hx_Field_makeIterator(v any) any
	Hx_Field_loopRun(f func()) bool
	Hx_Field_setVar(name string, v any) any
	Hx_Field_forLoop(n string, it any, e any)
	Hx_Field_dynamicCall(args HxArray[any]) any
	Hx_Field_duplicate(h Hx_Typedef_map) *Hx_Obj_haxe_imap
	Hx_Field_error(e Hx_Enum_loreline_lorscript_errordef, rethrow__isOptional any) any
	Hx_Field_setMapValue(_hx_reserved_map any, key any, value any)
	Hx_Field_call(o any, f any, args HxArray[any]) any
	Hx_Field_cnew(cl string, args HxArray[any]) any
	Hx_Field_initOps()
	Hx_Field_resolve(id string) any
	Hx_Field_makeMap(keys HxArray[any], values HxArray[any]) any
	Hx_Field_set(o any, f string, v any) any
	Hx_Field_exprReturn(e any) any
	Hx_Field_getMapValue(_hx_reserved_map any, key any) any
	Hx_Field_callIsShadowed() bool
	Hx_Field_assign(e1 any, e2 any) any
	Hx_Field_increment(e any, prefix bool, delta int32) any
	Hx_Field_expr(e any) any
	Hx_Field_get(o any, f string) any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_interp struct {
	VTable               Hx_Obj_VTable_loreline_lorscript_interp `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_variables   Hx_Typedef_map
	Hx_Field_locals      Hx_Typedef_map
	Hx_Field_binops      Hx_Typedef_map
	Hx_Field_depth       int32
	Hx_Field_inTry       bool
	Hx_Field_declared    HxArray[any]
	Hx_Field_returnValue any
	Hx_Field_curExpr     any
	Hx_Field_interpreter *Hx_Obj_loreline_interpreter
}

func Hx_Obj_loreline_lorscript_interp_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_interp {
	obj := &Hx_Obj_loreline_lorscript_interp{}
	obj.VTable = obj
	obj.Hx_Field_variables = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_locals = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_binops = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_depth = HxDefault[int32]()
	obj.Hx_Field_inTry = HxDefault[bool]()
	obj.Hx_Field_declared = HxDefault[HxArray[any]]()
	obj.Hx_Field_returnValue = HxDefault[any]()
	obj.Hx_Field_curExpr = HxDefault[any]()
	obj.Hx_Field_interpreter = HxDefault[*Hx_Obj_loreline_interpreter]()
	return obj
}

func Hx_Obj_loreline_lorscript_interp_CreateInstance(interpreter *Hx_Obj_loreline_interpreter) *Hx_Obj_loreline_lorscript_interp {
	obj := Hx_Obj_loreline_lorscript_interp_CreateEmptyInstance()
	obj.Hx_New(interpreter)
	return obj
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_New(interpreter *Hx_Obj_loreline_interpreter) {
	{
		this.Hx_Field_interpreter = interpreter
		this.Hx_Field_locals = (func() *Hx_Obj_haxe_imap {
			_hx_icast_206 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_206 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_206, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_206.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		this.Hx_Field_declared = HxMakeArray[any]().Dyn()
		this.VTable.Hx_Field_resetVariables()
		this.VTable.Hx_Field_initOps()
	}

}

var Hx_Obj_loreline_lorscript_interp_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Interp",
	HxMakeArray[string]("__keepIntIterator"),
	HxMakeArray[string]("variables", "locals", "binops", "depth", "inTry", "declared", "returnValue", "curExpr", "interpreter", "resetVariables", "posInfos", "initOps", "setVar", "assign", "assignOp", "evalAssignOp", "increment", "execute", "exprReturn", "duplicate", "restore", "error", "rethrow", "resolve", "expr", "doWhileLoop", "whileLoop", "makeIterator", "forLoop", "loopRun", "isMap", "getMapValue", "setMapValue", "makeMap", "get", "set", "fcall", "call", "callIsShadowed", "dynamicCall", "cnew"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_interp_CreateInstance(HxConvert[*Hx_Obj_loreline_interpreter](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lorscript_interp_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_interp_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "__keepIntIterator":
			return Hx_Field_loreline_lorscript_interp___keepIntIterator
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_interp_RTTI
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_As_Hx_Obj_loreline_lorscript_interp() *Hx_Obj_loreline_lorscript_interp {
	return this
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_resetVariables() {
	this.Hx_Field_variables = (func() *Hx_Obj_haxe_imap {
		_hx_icast_177 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_177 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_177, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_177.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field_variables.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("null", nil)
	this.Hx_Field_variables.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("true", ((any)(true)))
	this.Hx_Field_variables.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("false", ((any)(false)))
	var _gthis *Hx_Obj_loreline_lorscript_interp
	_ = _gthis
	_gthis = this
	this.Hx_Field_variables.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("trace", Hx_Field_reflect_makeVarArgs(func() func(el HxArray[any]) any {
		var _hx_closure_src_540 func(el HxArray[any])
		_ = _hx_closure_src_540
		_hx_closure_src_540 = func(el HxArray[any]) {
			var inf any
			_ = inf
			inf = _gthis.VTable.Hx_Field_posInfos()
			var v any
			_ = v
			v = (Hx_Array_Shift[any](el).Value)
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(el) > 0 {
				Hx_Field_go_haxe_hxdynamic_setField(inf, "customParams", el)
			}

			Hx_Field_haxe_log_trace(Hx_Field_std_string(v), inf)
		}

		return func(el HxArray[any]) any {
			_hx_closure_src_540(el)
			return nil
		}
	}()))
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_posInfos() any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_curExpr, nil))) {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")
		return any(map[string]any{"fileName": _hx_tmp_0, "lineNumber": Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")})
	}

	return any(map[string]any{"fileName": ((any)("lorscript")), "lineNumber": ((int32)(0))})
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_initOps() {
	var me *Hx_Obj_loreline_lorscript_interp
	_ = me
	me = this
	this.Hx_Field_binops = (func() *Hx_Obj_haxe_imap {
		_hx_icast_178 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_178 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_178, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_178.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("+", func(e1 any, e2 any) any {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = me.VTable.Hx_Field_expr(e1)
		return Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, me.VTable.Hx_Field_expr(e2))
	})
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("-", func() func(e1 any, e2 any) any {
		var _hx_closure_src_541 func(e1 any, e2 any) float64
		_ = _hx_closure_src_541
		_hx_closure_src_541 = func(e1 any, e2 any) float64 {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = me.VTable.Hx_Field_expr(e1)
			return ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_subtract(_hx_tmp_1, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_541(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("*", func() func(e1 any, e2 any) any {
		var _hx_closure_src_542 func(e1 any, e2 any) float64
		_ = _hx_closure_src_542
		_hx_closure_src_542 = func(e1 any, e2 any) float64 {
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = me.VTable.Hx_Field_expr(e1)
			return ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_multiply(_hx_tmp_2, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_542(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("/", func() func(e1 any, e2 any) any {
		var _hx_closure_src_543 func(e1 any, e2 any) float64
		_ = _hx_closure_src_543
		_hx_closure_src_543 = func(e1 any, e2 any) float64 {
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = me.VTable.Hx_Field_expr(e1)
			return ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_divide(_hx_tmp_3, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_543(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("%", func() func(e1 any, e2 any) any {
		var _hx_closure_src_544 func(e1 any, e2 any) float64
		_ = _hx_closure_src_544
		_hx_closure_src_544 = func(e1 any, e2 any) float64 {
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = me.VTable.Hx_Field_expr(e1)
			return ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_modulo(_hx_tmp_4, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_544(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("&", func() func(e1 any, e2 any) any {
		var _hx_closure_src_545 func(e1 any, e2 any) int32
		_ = _hx_closure_src_545
		_hx_closure_src_545 = func(e1 any, e2 any) int32 {
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_bitand(_hx_tmp_5, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_545(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("|", func() func(e1 any, e2 any) any {
		var _hx_closure_src_546 func(e1 any, e2 any) int32
		_ = _hx_closure_src_546
		_hx_closure_src_546 = func(e1 any, e2 any) int32 {
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_bitor(_hx_tmp_6, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_546(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("^", func() func(e1 any, e2 any) any {
		var _hx_closure_src_547 func(e1 any, e2 any) int32
		_ = _hx_closure_src_547
		_hx_closure_src_547 = func(e1 any, e2 any) int32 {
			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_bitxor(_hx_tmp_7, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_547(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("<<", func() func(e1 any, e2 any) any {
		var _hx_closure_src_548 func(e1 any, e2 any) int32
		_ = _hx_closure_src_548
		_hx_closure_src_548 = func(e1 any, e2 any) int32 {
			var _hx_tmp_8 any
			_ = _hx_tmp_8
			_hx_tmp_8 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_lbitshift(_hx_tmp_8, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_548(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(">>", func() func(e1 any, e2 any) any {
		var _hx_closure_src_549 func(e1 any, e2 any) int32
		_ = _hx_closure_src_549
		_hx_closure_src_549 = func(e1 any, e2 any) int32 {
			var _hx_tmp_9 any
			_ = _hx_tmp_9
			_hx_tmp_9 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_rbitshift(_hx_tmp_9, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_549(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(">>>", func() func(e1 any, e2 any) any {
		var _hx_closure_src_550 func(e1 any, e2 any) int32
		_ = _hx_closure_src_550
		_hx_closure_src_550 = func(e1 any, e2 any) int32 {
			var _hx_tmp_10 any
			_ = _hx_tmp_10
			_hx_tmp_10 = me.VTable.Hx_Field_expr(e1)
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_urbitshift(_hx_tmp_10, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_550(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("==", func() func(e1 any, e2 any) any {
		var _hx_closure_src_551 func(e1 any, e2 any) bool
		_ = _hx_closure_src_551
		_hx_closure_src_551 = func(e1 any, e2 any) bool {
			var _hx_tmp_11 any
			_ = _hx_tmp_11
			_hx_tmp_11 = me.VTable.Hx_Field_expr(e1)
			return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_11, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_551(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("!=", func() func(e1 any, e2 any) any {
		var _hx_closure_src_552 func(e1 any, e2 any) bool
		_ = _hx_closure_src_552
		_hx_closure_src_552 = func(e1 any, e2 any) bool {
			var _hx_tmp_12 any
			_ = _hx_tmp_12
			_hx_tmp_12 = me.VTable.Hx_Field_expr(e1)
			return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_12, me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_552(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(">=", func() func(e1 any, e2 any) any {
		var _hx_closure_src_553 func(e1 any, e2 any) bool
		_ = _hx_closure_src_553
		_hx_closure_src_553 = func(e1 any, e2 any) bool {
			var v1 int32
			_ = v1
			v1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e1)))))
			var v2 int32
			_ = v2
			v2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e2)))))
			return (v1 >= v2)
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_553(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("<=", func() func(e1 any, e2 any) any {
		var _hx_closure_src_554 func(e1 any, e2 any) bool
		_ = _hx_closure_src_554
		_hx_closure_src_554 = func(e1 any, e2 any) bool {
			var v1 int32
			_ = v1
			v1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e1)))))
			var v2 int32
			_ = v2
			v2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e2)))))
			return (v1 <= v2)
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_554(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(">", func() func(e1 any, e2 any) any {
		var _hx_closure_src_555 func(e1 any, e2 any) bool
		_ = _hx_closure_src_555
		_hx_closure_src_555 = func(e1 any, e2 any) bool {
			var v1 int32
			_ = v1
			v1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e1)))))
			var v2 int32
			_ = v2
			v2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e2)))))
			return (v1 > v2)
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_555(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("<", func() func(e1 any, e2 any) any {
		var _hx_closure_src_556 func(e1 any, e2 any) bool
		_ = _hx_closure_src_556
		_hx_closure_src_556 = func(e1 any, e2 any) bool {
			var v1 int32
			_ = v1
			v1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e1)))))
			var v2 int32
			_ = v2
			v2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e2)))))
			return (v1 < v2)
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_556(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("||", func() func(e1 any, e2 any) any {
		var _hx_closure_src_557 func(e1 any, e2 any) bool
		_ = _hx_closure_src_557
		_hx_closure_src_557 = func(e1 any, e2 any) bool {
			var _hx_tmp_13 bool
			_ = _hx_tmp_13
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(me.VTable.Hx_Field_expr(e1), ((any)(true)))))) {
				_hx_tmp_13 = true
			} else {
				_hx_tmp_13 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(me.VTable.Hx_Field_expr(e2), ((any)(true)))))))
			}

			return _hx_tmp_13
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_557(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("&&", func() func(e1 any, e2 any) any {
		var _hx_closure_src_558 func(e1 any, e2 any) bool
		_ = _hx_closure_src_558
		_hx_closure_src_558 = func(e1 any, e2 any) bool {
			var _hx_tmp_14 bool
			_ = _hx_tmp_14
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(me.VTable.Hx_Field_expr(e1), ((any)(true)))))) {
				_hx_tmp_14 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(me.VTable.Hx_Field_expr(e2), ((any)(true)))))))
			} else {
				_hx_tmp_14 = false
			}

			return _hx_tmp_14
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_558(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("=", this.VTable.Hx_Field_assign)
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("...", func() func(e1 any, e2 any) any {
		var _hx_closure_src_559 func(e1 any, e2 any) *Hx_Obj_intiterator
		_ = _hx_closure_src_559
		_hx_closure_src_559 = func(e1 any, e2 any) *Hx_Obj_intiterator {
			return Hx_Obj_intiterator_CreateInstance(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e1))))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(me.VTable.Hx_Field_expr(e2))))))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_559(e1, e2)))
		}
	}())
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set("is", func() func(e1 any, e2 any) any {
		var _hx_closure_src_560 func(e1 any, e2 any) bool
		_ = _hx_closure_src_560
		_hx_closure_src_560 = func(e1 any, e2 any) bool {
			var _hx_tmp_15 any
			_ = _hx_tmp_15
			_hx_tmp_15 = me.VTable.Hx_Field_expr(e1)
			return Hx_Field_std_isOfType(_hx_tmp_15, me.VTable.Hx_Field_expr(e2))
		}

		return func(e1 any, e2 any) any {
			return ((any)(_hx_closure_src_560(e1, e2)))
		}
	}())
	this.VTable.Hx_Field_assignOp("+=", func(v1 any, v2 any) any {
		return Hx_Field_go_haxe_hxdynamic_add(v1, v2)
	})
	this.VTable.Hx_Field_assignOp("-=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_561 func(v1 float64, v2 float64) float64
		_ = _hx_closure_src_561
		_hx_closure_src_561 = func(v1 float64, v2 float64) float64 {
			return (v1 - v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_16 float64
			_ = _hx_tmp_16
			_hx_tmp_16 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v1)))
			return ((any)(_hx_closure_src_561(_hx_tmp_16, ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("*=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_562 func(v1 float64, v2 float64) float64
		_ = _hx_closure_src_562
		_hx_closure_src_562 = func(v1 float64, v2 float64) float64 {
			return (v1 * v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_17 float64
			_ = _hx_tmp_17
			_hx_tmp_17 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v1)))
			return ((any)(_hx_closure_src_562(_hx_tmp_17, ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("/=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_563 func(v1 float64, v2 float64) float64
		_ = _hx_closure_src_563
		_hx_closure_src_563 = func(v1 float64, v2 float64) float64 {
			return ((float64)((v1 / v2)))
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_18 float64
			_ = _hx_tmp_18
			_hx_tmp_18 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v1)))
			return ((any)(_hx_closure_src_563(_hx_tmp_18, ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("%=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_564 func(v1 float64, v2 float64) float64
		_ = _hx_closure_src_564
		_hx_closure_src_564 = func(v1 float64, v2 float64) float64 {
			return math.Mod(v1, v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_19 float64
			_ = _hx_tmp_19
			_hx_tmp_19 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v1)))
			return ((any)(_hx_closure_src_564(_hx_tmp_19, ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("&=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_565 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_565
		_hx_closure_src_565 = func(v1 int32, v2 int32) int32 {
			return (v1 & v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_20 int32
			_ = _hx_tmp_20
			_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_565(_hx_tmp_20, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("|=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_566 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_566
		_hx_closure_src_566 = func(v1 int32, v2 int32) int32 {
			return (v1 | v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_21 int32
			_ = _hx_tmp_21
			_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_566(_hx_tmp_21, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("^=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_567 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_567
		_hx_closure_src_567 = func(v1 int32, v2 int32) int32 {
			return (v1 ^ v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_22 int32
			_ = _hx_tmp_22
			_hx_tmp_22 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_567(_hx_tmp_22, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp("<<=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_568 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_568
		_hx_closure_src_568 = func(v1 int32, v2 int32) int32 {
			return (v1 << v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_23 int32
			_ = _hx_tmp_23
			_hx_tmp_23 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_568(_hx_tmp_23, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp(">>=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_569 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_569
		_hx_closure_src_569 = func(v1 int32, v2 int32) int32 {
			return (v1 >> v2)
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_24 int32
			_ = _hx_tmp_24
			_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_569(_hx_tmp_24, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
	this.VTable.Hx_Field_assignOp(">>>=", func() func(v1 any, v2 any) any {
		var _hx_closure_src_570 func(v1 int32, v2 int32) int32
		_ = _hx_closure_src_570
		_hx_closure_src_570 = func(v1 int32, v2 int32) int32 {
			return ((int32)((((uint32)(v1)) >> v2)))
		}

		return func(v1 any, v2 any) any {
			var _hx_tmp_25 int32
			_ = _hx_tmp_25
			_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v1)))
			return ((any)(_hx_closure_src_570(_hx_tmp_25, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(v2))))))
		}
	}())
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_setVar(name string, v any) any {
	if !func() func(_hx_param_key string) bool {
		var _hx_closure_src_571 func(_hx_param_key any) bool
		_ = _hx_closure_src_571
		_hx_closure_src_571 = this.Hx_Field_variables.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_571(((any)(_hx_param_key)))
		}
	}()(name) {
		Hx_Field_loreline_objects_setField(this.Hx_Field_interpreter, this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name, v)
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)("Invalid assign"))))
	}

	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_assign(e1 any, e2 any) any {
	var v any
	_ = v
	v = this.VTable.Hx_Field_expr(e2)
	{
		var _g Hx_Enum_loreline_lorscript_exprdef
		_ = _g
		_g = func() Hx_Enum_loreline_lorscript_exprdef {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_loreline_lorscript_exprdef
				return _hx_z
			}
			return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
		}()
		switch _g.Hx_Field_enumIndex() {
		case 1:
			{
				var id string
				_ = id
				id = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
				{
					var l any
					_ = l
					l = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, nil))) {
						this.VTable.Hx_Field_setVar(id, v)
					} else {
						Hx_Field_go_haxe_hxdynamic_setField(l, "r", v)
					}
				}
			}

		case 5:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
				var f string
				_ = f
				f = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
				v = this.VTable.Hx_Field_set(this.VTable.Hx_Field_expr(e), f, v)
			}

		case 16:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
				var index any
				_ = index
				index = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
				{
					var arr any
					_ = arr
					arr = this.VTable.Hx_Field_expr(e)
					var index1 any
					_ = index1
					index1 = this.VTable.Hx_Field_expr(index)
					if Hx_Field_loreline_arrays_isArray(arr) {
						var _hx_tmp_0 any
						_ = _hx_tmp_0
						_hx_tmp_0 = arr
						Hx_Field_loreline_arrays_arraySet(_hx_tmp_0, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
					} else {
						if Hx_Field_loreline_objects_isFields(arr) {
							var _hx_tmp_1 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_1
							_hx_tmp_1 = this.Hx_Field_interpreter
							var _hx_tmp_2 any
							_ = _hx_tmp_2
							_hx_tmp_2 = arr
							Hx_Field_loreline_objects_setField(_hx_tmp_1, _hx_tmp_2, Hx_Field_go_haxe_hxdynamic_toString(index1), v)
						} else {
							if Hx_Field_std_isOfType(arr, Hx_Obj_haxe_imap_RTTI) {
								var hx_cast_ex_179 any
								_ = hx_cast_ex_179
								hx_cast_ex_179 = arr
								if !Hx_Field_std_isOfType(hx_cast_ex_179, Hx_Obj_haxe_imap_RTTI) {
									panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
								}

								var hx_dyncast_180 any
								_ = hx_dyncast_180
								hx_dyncast_180 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_179, "Hx_Obj_haxe_imap")
								var _hx_tmp_3 *Hx_Obj_haxe_imap
								_ = _hx_tmp_3
								if hx_dyncast_180 != nil {
									_hx_tmp_3 = hx_dyncast_180.(*Hx_Obj_haxe_imap)
								} else {
									_hx_tmp_3 = nil
								}

								(_hx_tmp_3).VTable.Hx_Field_set(index1, v)
							} else {
								var _hx_tmp_4 any
								_ = _hx_tmp_4
								_hx_tmp_4 = arr
								Hx_Field_go_haxe_hxdynamic_setArrayIndex(_hx_tmp_4, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
							}
						}
					}
				}
			}

		default:
			{
				var e *Hx_Obj_loreline_lorscript_error
				_ = e
				e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{"="}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			}
		}
	}

	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_assignOp(op string, fop func(_hx_carg0 any, _hx_carg1 any) any) {
	var me *Hx_Obj_loreline_lorscript_interp
	_ = me
	me = this
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = op
	this.Hx_Field_binops.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_0, func(e1 any, e2 any) any {
		return me.VTable.Hx_Field_evalAssignOp(op, fop, e1, e2)
	})
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_evalAssignOp(op string, fop func(_hx_carg0 any, _hx_carg1 any) any, e1 any, e2 any) any {
	var v any
	_ = v
	{
		var _g Hx_Enum_loreline_lorscript_exprdef
		_ = _g
		_g = func() Hx_Enum_loreline_lorscript_exprdef {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_loreline_lorscript_exprdef
				return _hx_z
			}
			return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
		}()
		switch _g.Hx_Field_enumIndex() {
		case 1:
			{
				var id string
				_ = id
				id = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
				{
					var l any
					_ = l
					l = this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
					var _hx_tmp_0 any
					_ = _hx_tmp_0
					_hx_tmp_0 = this.VTable.Hx_Field_expr(e1)
					v = fop(_hx_tmp_0, this.VTable.Hx_Field_expr(e2))
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, nil))) {
						this.VTable.Hx_Field_setVar(id, v)
					} else {
						Hx_Field_go_haxe_hxdynamic_setField(l, "r", v)
					}
				}
			}

		case 5:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
				var f string
				_ = f
				f = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
				{
					var obj any
					_ = obj
					obj = this.VTable.Hx_Field_expr(e)
					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = this.VTable.Hx_Field_get(obj, f)
					v = fop(_hx_tmp_1, this.VTable.Hx_Field_expr(e2))
					v = this.VTable.Hx_Field_set(obj, f, v)
				}
			}

		case 16:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
				var index any
				_ = index
				index = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
				{
					var arr any
					_ = arr
					arr = this.VTable.Hx_Field_expr(e)
					var index1 any
					_ = index1
					index1 = this.VTable.Hx_Field_expr(index)
					if Hx_Field_loreline_arrays_isArray(arr) {
						var _hx_tmp_3 any
						_ = _hx_tmp_3
						_hx_tmp_3 = arr
						var _hx_tmp_2 any
						_ = _hx_tmp_2
						_hx_tmp_2 = Hx_Field_loreline_arrays_arrayGet(_hx_tmp_3, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))))
						v = fop(_hx_tmp_2, this.VTable.Hx_Field_expr(e2))
						var _hx_tmp_4 any
						_ = _hx_tmp_4
						_hx_tmp_4 = arr
						Hx_Field_loreline_arrays_arraySet(_hx_tmp_4, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
					} else {
						if Hx_Field_loreline_objects_isFields(arr) {
							var _hx_tmp_6 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_6
							_hx_tmp_6 = this.Hx_Field_interpreter
							var _hx_tmp_7 any
							_ = _hx_tmp_7
							_hx_tmp_7 = arr
							var _hx_tmp_5 any
							_ = _hx_tmp_5
							_hx_tmp_5 = Hx_Field_loreline_objects_getField(_hx_tmp_6, _hx_tmp_7, Hx_Field_go_haxe_hxdynamic_toString(index1))
							v = fop(_hx_tmp_5, this.VTable.Hx_Field_expr(e2))
							var _hx_tmp_8 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_8
							_hx_tmp_8 = this.Hx_Field_interpreter
							var _hx_tmp_9 any
							_ = _hx_tmp_9
							_hx_tmp_9 = arr
							Hx_Field_loreline_objects_setField(_hx_tmp_8, _hx_tmp_9, Hx_Field_go_haxe_hxdynamic_toString(index1), v)
						} else {
							if Hx_Field_std_isOfType(arr, Hx_Obj_haxe_imap_RTTI) {
								var hx_cast_ex_181 any
								_ = hx_cast_ex_181
								hx_cast_ex_181 = arr
								if !Hx_Field_std_isOfType(hx_cast_ex_181, Hx_Obj_haxe_imap_RTTI) {
									panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
								}

								var hx_dyncast_182 any
								_ = hx_dyncast_182
								hx_dyncast_182 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_181, "Hx_Obj_haxe_imap")
								var _hx_tmp_10 *Hx_Obj_haxe_imap
								_ = _hx_tmp_10
								if hx_dyncast_182 != nil {
									_hx_tmp_10 = hx_dyncast_182.(*Hx_Obj_haxe_imap)
								} else {
									_hx_tmp_10 = nil
								}

								v = fop((_hx_tmp_10).VTable.Hx_Field_get(index1), this.VTable.Hx_Field_expr(e2))
								var hx_cast_ex_183 any
								_ = hx_cast_ex_183
								hx_cast_ex_183 = arr
								if !Hx_Field_std_isOfType(hx_cast_ex_183, Hx_Obj_haxe_imap_RTTI) {
									panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
								}

								var hx_dyncast_184 any
								_ = hx_dyncast_184
								hx_dyncast_184 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_183, "Hx_Obj_haxe_imap")
								var _hx_tmp_11 *Hx_Obj_haxe_imap
								_ = _hx_tmp_11
								if hx_dyncast_184 != nil {
									_hx_tmp_11 = hx_dyncast_184.(*Hx_Obj_haxe_imap)
								} else {
									_hx_tmp_11 = nil
								}

								(_hx_tmp_11).VTable.Hx_Field_set(index1, v)
							} else {
								var _hx_tmp_13 any
								_ = _hx_tmp_13
								_hx_tmp_13 = arr
								var _hx_tmp_12 any
								_ = _hx_tmp_12
								_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_13, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))))
								v = fop(_hx_tmp_12, this.VTable.Hx_Field_expr(e2))
								var _hx_tmp_14 any
								_ = _hx_tmp_14
								_hx_tmp_14 = arr
								Hx_Field_go_haxe_hxdynamic_setArrayIndex(_hx_tmp_14, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
							}
						}
					}
				}
			}

		default:
			{
				var e *Hx_Obj_loreline_lorscript_error
				_ = e
				e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{op}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			}
		}
	}

	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_increment(e any, prefix bool, delta int32) any {
	this.Hx_Field_curExpr = e
	var e1 Hx_Enum_loreline_lorscript_exprdef
	_ = e1
	e1 = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch e1.Hx_Field_enumIndex() {
	case 1:
		{
			var id string
			_ = id
			id = e1.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
			var l any
			_ = l
			l = this.Hx_Field_locals.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, nil))) {
				_hx_tmp_0 = this.VTable.Hx_Field_resolve(id)
			} else {
				_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(l, "r")
			}

			var v any
			_ = v
			v = _hx_tmp_0
			if prefix {
				v = Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta)))
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, nil))) {
					this.VTable.Hx_Field_setVar(id, v)
				} else {
					Hx_Field_go_haxe_hxdynamic_setField(l, "r", v)
				}
			} else {
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, nil))) {
					var _hx_tmp_1 string
					_ = _hx_tmp_1
					_hx_tmp_1 = id
					this.VTable.Hx_Field_setVar(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta))))
				} else {
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = l
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "r", Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta))))
				}
			}

			return v
		}

	case 5:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
			var f string
			_ = f
			f = e1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
			var obj any
			_ = obj
			obj = this.VTable.Hx_Field_expr(e)
			var v any
			_ = v
			v = this.VTable.Hx_Field_get(obj, f)
			if prefix {
				v = Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta)))
				this.VTable.Hx_Field_set(obj, f, v)
			} else {
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = obj
				var _hx_tmp_4 string
				_ = _hx_tmp_4
				_hx_tmp_4 = f
				this.VTable.Hx_Field_set(_hx_tmp_3, _hx_tmp_4, Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta))))
			}

			return v
		}

	case 16:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
			var index any
			_ = index
			index = e1.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
			var arr any
			_ = arr
			arr = this.VTable.Hx_Field_expr(e)
			var index1 any
			_ = index1
			index1 = this.VTable.Hx_Field_expr(index)
			if Hx_Field_loreline_arrays_isArray(arr) {
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = arr
				var v any
				_ = v
				v = Hx_Field_loreline_arrays_arrayGet(_hx_tmp_5, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))))
				if prefix {
					v = Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta)))
					var _hx_tmp_6 any
					_ = _hx_tmp_6
					_hx_tmp_6 = arr
					Hx_Field_loreline_arrays_arraySet(_hx_tmp_6, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
				} else {
					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = arr
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1)))
					Hx_Field_loreline_arrays_arraySet(_hx_tmp_7, _hx_tmp_8, Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta))))
				}

				return v
			} else {
				if Hx_Field_loreline_objects_isFields(arr) {
					var _hx_tmp_9 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_9
					_hx_tmp_9 = this.Hx_Field_interpreter
					var _hx_tmp_10 any
					_ = _hx_tmp_10
					_hx_tmp_10 = arr
					var v any
					_ = v
					v = Hx_Field_loreline_objects_getField(_hx_tmp_9, _hx_tmp_10, Hx_Field_go_haxe_hxdynamic_toString(index1))
					if prefix {
						v = Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta)))
						var _hx_tmp_11 *Hx_Obj_loreline_interpreter
						_ = _hx_tmp_11
						_hx_tmp_11 = this.Hx_Field_interpreter
						var _hx_tmp_12 any
						_ = _hx_tmp_12
						_hx_tmp_12 = arr
						Hx_Field_loreline_objects_setField(_hx_tmp_11, _hx_tmp_12, Hx_Field_go_haxe_hxdynamic_toString(index1), v)
					} else {
						var _hx_tmp_13 *Hx_Obj_loreline_interpreter
						_ = _hx_tmp_13
						_hx_tmp_13 = this.Hx_Field_interpreter
						var _hx_tmp_14 any
						_ = _hx_tmp_14
						_hx_tmp_14 = arr
						var _hx_tmp_15 string
						_ = _hx_tmp_15
						_hx_tmp_15 = Hx_Field_go_haxe_hxdynamic_toString(index1)
						Hx_Field_loreline_objects_setField(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15, Hx_Field_go_haxe_hxdynamic_add(v, ((any)(delta))))
					}

					return v
				} else {
					if Hx_Field_std_isOfType(arr, Hx_Obj_haxe_imap_RTTI) {
						var hx_cast_ex_191 any
						_ = hx_cast_ex_191
						hx_cast_ex_191 = arr
						if !Hx_Field_std_isOfType(hx_cast_ex_191, Hx_Obj_haxe_imap_RTTI) {
							panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
						}

						var hx_dyncast_192 any
						_ = hx_dyncast_192
						hx_dyncast_192 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_191, "Hx_Obj_haxe_imap")
						var _hx_tmp_16 *Hx_Obj_haxe_imap
						_ = _hx_tmp_16
						if hx_dyncast_192 != nil {
							_hx_tmp_16 = hx_dyncast_192.(*Hx_Obj_haxe_imap)
						} else {
							_hx_tmp_16 = nil
						}

						var v int32
						_ = v
						v = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt((_hx_tmp_16).VTable.Hx_Field_get(index1)))))
						if prefix {
							v = (v + delta)
							var hx_cast_ex_187 any
							_ = hx_cast_ex_187
							hx_cast_ex_187 = arr
							if !Hx_Field_std_isOfType(hx_cast_ex_187, Hx_Obj_haxe_imap_RTTI) {
								panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
							}

							var hx_dyncast_188 any
							_ = hx_dyncast_188
							hx_dyncast_188 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_187, "Hx_Obj_haxe_imap")
							var _hx_tmp_17 *Hx_Obj_haxe_imap
							_ = _hx_tmp_17
							if hx_dyncast_188 != nil {
								_hx_tmp_17 = hx_dyncast_188.(*Hx_Obj_haxe_imap)
							} else {
								_hx_tmp_17 = nil
							}

							(_hx_tmp_17).VTable.Hx_Field_set(index1, ((any)(v)))
						} else {
							var hx_cast_ex_189 any
							_ = hx_cast_ex_189
							hx_cast_ex_189 = arr
							if !Hx_Field_std_isOfType(hx_cast_ex_189, Hx_Obj_haxe_imap_RTTI) {
								panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
							}

							var hx_dyncast_190 any
							_ = hx_dyncast_190
							hx_dyncast_190 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_189, "Hx_Obj_haxe_imap")
							var _hx_tmp_18 *Hx_Obj_haxe_imap
							_ = _hx_tmp_18
							if hx_dyncast_190 != nil {
								_hx_tmp_18 = hx_dyncast_190.(*Hx_Obj_haxe_imap)
							} else {
								_hx_tmp_18 = nil
							}

							(_hx_tmp_18).VTable.Hx_Field_set(index1, ((any)((v + delta))))
						}

						return ((any)(v))
					} else {
						var _hx_tmp_19 any
						_ = _hx_tmp_19
						_hx_tmp_19 = arr
						var v int32
						_ = v
						v = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_19, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1)))))))
						if prefix {
							v = (v + delta)
							var _hx_tmp_20 any
							_ = _hx_tmp_20
							_hx_tmp_20 = arr
							Hx_Field_go_haxe_hxdynamic_setArrayIndex(_hx_tmp_20, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), v)
						} else {
							var _hx_tmp_21 any
							_ = _hx_tmp_21
							_hx_tmp_21 = arr
							Hx_Field_go_haxe_hxdynamic_setArrayIndex(_hx_tmp_21, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))), (v + delta))
						}

						return ((any)(v))
					}
				}
			}
		}

	default:
		{
			var _hx_tmp_22 string
			_ = _hx_tmp_22
			if delta > 0 {
				_hx_tmp_22 = "++"
			} else {
				_hx_tmp_22 = "--"
			}

			var e *Hx_Obj_loreline_lorscript_error
			_ = e
			e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{_hx_tmp_22}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_execute(expr any) any {
	this.Hx_Field_depth = 0
	this.Hx_Field_locals = (func() *Hx_Obj_haxe_imap {
		_hx_icast_193 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_193 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_193, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_193.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field_declared = HxMakeArray[any]().Dyn()
	return this.VTable.Hx_Field_exprReturn(expr)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_exprReturn(e any) any {
	{
		var hx_try_return any
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
						_ = e
						{
							switch e.Hx_Field_enumIndex() {
							case 0:
								{
									panic(Hx_Field_haxe_exception_thrown(((any)("Invalid break"))))
								}

							case 1:
								{
									panic(Hx_Field_haxe_exception_thrown(((any)("Invalid continue"))))
								}

							case 2:
								{
									var v any
									_ = v
									v = this.Hx_Field_returnValue
									this.Hx_Field_returnValue = nil
									hx_try_state = 1
									hx_try_return = v
									return
								}

							default:
								{
									panic(Hx_Field_haxe_exception_thrown("unreachable"))
								}
							}
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				hx_try_state = 1
				hx_try_return = this.VTable.Hx_Field_expr(e)
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_duplicate(h Hx_Typedef_map) *Hx_Obj_haxe_imap {
	var h2 *Hx_Obj_haxe_imap
	_ = h2
	h2 = (func() *Hx_Obj_haxe_imap {
		_hx_icast_194 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_194 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_194, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_194.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	{
		var _gk any
		_ = _gk
		_gk = h.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gk, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gk, "next")
			var k string
			_ = k
			k = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = k
			h2.VTable.Hx_Field_set(_hx_tmp_2, h.VTable.Hx_Field_get(k))
		}
	}

	return h2
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_restore(old int32) {
	for {
		if !(Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared) > old) {
			break
		}

		var d any
		_ = d
		d = (Hx_Array_Pop[any](HxMakeArrayView[any](this.Hx_Field_declared)).Value)
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(d, "n")
		this.Hx_Field_locals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(d, "old"))
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_error(e Hx_Enum_loreline_lorscript_errordef, rethrow__isOptional any) any {
	rethrow := HxDefault[bool]()
	_ = rethrow

	if rethrow__isOptional != nil {
		rethrow = HxConvert[bool](rethrow__isOptional)
	} else {
		rethrow = false
	}
	{
		var e1 *Hx_Obj_loreline_lorscript_error
		_ = e1
		e1 = Hx_Obj_loreline_lorscript_error_CreateInstance(e, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		if rethrow {
			panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
		} else {
			panic(Hx_Field_haxe_exception_thrown(((any)(e1))))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_rethrow(e any) {
	panic(Hx_Field_haxe_exception_thrown(e))
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_resolve(id string) any {
	var v any
	_ = v
	v = this.Hx_Field_variables.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, nil))) {
		_hx_tmp_0 = !func() func(_hx_param_key string) bool {
			var _hx_closure_src_572 func(_hx_param_key any) bool
			_ = _hx_closure_src_572
			_hx_closure_src_572 = this.Hx_Field_variables.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_572(((any)(_hx_param_key)))
			}
		}()(id)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		if Hx_Field_loreline_objects_fieldExists(this.Hx_Field_interpreter, this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, id) {
			return Hx_Field_loreline_objects_getField(this.Hx_Field_interpreter, this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, id)
		}

		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_573 func(_hx_param_key any) bool
			_ = _hx_closure_src_573
			_hx_closure_src_573 = this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_573(((any)(_hx_param_key)))
			}
		}()(id) {
			return func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
				var _hx_closure_src_576 func(_hx_param_key any) any
				_ = _hx_closure_src_576
				_hx_closure_src_576 = this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
					var hx_dyncast_196 any
					_ = hx_dyncast_196
					hx_dyncast_196 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_576(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
					var _hx_tmp_1 *Hx_Obj_loreline_runtimecharacter
					_ = _hx_tmp_1
					if hx_dyncast_196 != nil {
						_hx_tmp_1 = hx_dyncast_196.(interface {
							Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
						}).Hx_As_Hx_Obj_loreline_runtimecharacter()
					} else {
						_hx_tmp_1 = nil
					}

					return _hx_tmp_1
				}
			}()(id).Hx_Field_fields
		}

		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_575 func(_hx_param_key any) bool
			_ = _hx_closure_src_575
			_hx_closure_src_575 = this.Hx_Field_interpreter.Hx_Field_topLevelFunctions.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_575(((any)(_hx_param_key)))
			}
		}()(id) {
			return this.Hx_Field_interpreter.Hx_Field_topLevelFunctions.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
		}

		var beatRef *Hx_Obj_loreline_runtimebeatref
		_ = beatRef
		beatRef = this.Hx_Field_interpreter.VTable.Hx_Field_resolveBeatRefByName(id)
		if beatRef != nil {
			return ((any)(beatRef))
		}

		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EUnknownVariable{id}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_expr(e any) any {
	this.Hx_Field_curExpr = e
	var e1 Hx_Enum_loreline_lorscript_exprdef
	_ = e1
	e1 = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	var _gthis *Hx_Obj_loreline_lorscript_interp
	_ = _gthis
	_gthis = this
	switch e1.Hx_Field_enumIndex() {
	case 0:
		{
			var c Hx_Enum_loreline_lorscript_const
			_ = c
			c = e1.(Hx_Enum_loreline_lorscript_exprdef_EConst).Hx_Field_c
			switch c.Hx_Field_enumIndex() {
			case 0:
				{
					var v int32
					_ = v
					v = c.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
					return ((any)(v))
				}

			case 1:
				{
					var f float64
					_ = f
					f = c.(Hx_Enum_loreline_lorscript_const_CFloat).Hx_Field_f
					return ((any)(f))
				}

			case 2:
				{
					var s string
					_ = s
					s = c.(Hx_Enum_loreline_lorscript_const_CString).Hx_Field_s
					return ((any)(s))
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown("unreachable"))
				}
			}
		}

	case 1:
		{
			var id string
			_ = id
			id = e1.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
			var l any
			_ = l
			l = this.Hx_Field_locals.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(l, nil))) {
				return Hx_Field_go_haxe_hxdynamic_getField(l, "r")
			}

			return this.VTable.Hx_Field_resolve(id)
		}

	case 2:
		{
			var n string
			_ = n
			n = e1.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_n
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_e
			var _hx_tmp_0 HxArray[any]
			_ = _hx_tmp_0
			_hx_tmp_0 = HxMakeArrayView[any](this.Hx_Field_declared)
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = ((any)(n))
			_ = (Hx_Array_Push[any](_hx_tmp_0, any(map[string]any{"n": _hx_tmp_1, "old": this.Hx_Field_locals.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(n)})))
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = n
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_3 = nil
			} else {
				_hx_tmp_3 = this.VTable.Hx_Field_expr(e)
			}

			this.Hx_Field_locals.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_2, any(map[string]any{"r": _hx_tmp_3}))
			return nil
		}

	case 3:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
			return this.VTable.Hx_Field_expr(e)
		}

	case 4:
		{
			var exprs HxArray[any]
			_ = exprs
			exprs = e1.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
			var old int32
			_ = old
			old = Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared)
			var v any
			_ = v
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs)) {
						break
					}

					var e any
					_ = e
					e = exprs.Get(_g)
					_g++
					v = this.VTable.Hx_Field_expr(e)
				}
			}

			this.VTable.Hx_Field_restore(old)
			return v
		}

	case 5:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
			var f string
			_ = f
			f = e1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
			return this.VTable.Hx_Field_get(this.VTable.Hx_Field_expr(e), f)
		}

	case 6:
		{
			var op string
			_ = op
			op = e1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op
			var e11 any
			_ = e11
			e11 = e1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = e1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			var fop func(_hx_carg0 any, _hx_carg1 any) any
			_ = fop
			fop = func() func(_hx_carg0 any, _hx_carg1 any) any {
				var _hx_closure_src_577 any
				_ = _hx_closure_src_577
				_hx_closure_src_577 = this.Hx_Field_binops.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(op)
				return func(_hx_carg0 any, _hx_carg1 any) any {
					var _hx_tmp_5 any
					_ = _hx_tmp_5
					_hx_tmp_5 = _hx_closure_src_577
					return Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any](_hx_carg0, _hx_carg1).Dyn())
				}
			}()
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(fop)), ((any)(nil))))) {
				var e *Hx_Obj_loreline_lorscript_error
				_ = e
				e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{op}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			}

			return fop(e11, e2)
		}

	case 7:
		{
			var op string
			_ = op
			op = e1.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_op
			var prefix bool
			_ = prefix
			prefix = e1.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_prefix
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
			switch op {
			case "!":
				{
					return ((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.VTable.Hx_Field_expr(e), ((any)(true))))))))
				}

			case "++":
				{
					return this.VTable.Hx_Field_increment(e, prefix, 1)
				}

			case "-":
				{
					return ((any)(-((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(this.VTable.Hx_Field_expr(e))))))
				}

			case "--":
				{
					return this.VTable.Hx_Field_increment(e, prefix, -1)
				}

			case "~":
				{
					return ((any)(^((int32)(Hx_Field_go_haxe_hxdynamic_toInt(this.VTable.Hx_Field_expr(e))))))
				}

			default:
				{
					var tmp_e_1 *Hx_Obj_loreline_lorscript_error
					_ = tmp_e_1
					tmp_e_1 = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{op}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
					panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e_1))))
				}
			}
		}

	case 8:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_e
			var params HxArray[any]
			_ = params
			params = e1.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_params
			var args HxArray[any]
			_ = args
			args = HxMakeArray[any]().Dyn()
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _g
					if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
						break
					}

					var p any
					_ = p
					p = params.Get(_g)
					_g++
					var _hx_tmp_7 HxArray[any]
					_ = _hx_tmp_7
					_hx_tmp_7 = args
					_ = (Hx_Array_Push[any](_hx_tmp_7, this.VTable.Hx_Field_expr(p)))
				}
			}

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
			switch _g1.Hx_Field_enumIndex() {
			case 1:
				{
					var id string
					_ = id
					id = _g1.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
					var _hx_tmp_8 bool
					_ = _hx_tmp_8
					if id == "call" {
						_hx_tmp_8 = !this.VTable.Hx_Field_callIsShadowed()
					} else {
						_hx_tmp_8 = false
					}

					if _hx_tmp_8 {
						return this.VTable.Hx_Field_dynamicCall(args)
					} else {
						return this.VTable.Hx_Field_call(nil, this.VTable.Hx_Field_expr(e), args)
					}
				}

			case 5:
				{
					var tmp_e_1 any
					_ = tmp_e_1
					tmp_e_1 = _g1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
					var f string
					_ = f
					f = _g1.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
					var obj any
					_ = obj
					obj = this.VTable.Hx_Field_expr(tmp_e_1)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(obj, nil))) {
						var tmp_e_2 *Hx_Obj_loreline_lorscript_error
						_ = tmp_e_2
						tmp_e_2 = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidAccess{f}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
						panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e_2))))
					}

					return this.VTable.Hx_Field_fcall(obj, f, args)
				}

			default:
				{
					return this.VTable.Hx_Field_call(nil, this.VTable.Hx_Field_expr(e), args)
				}
			}
		}

	case 9:
		{
			var econd any
			_ = econd
			econd = e1.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_cond
			var e11 any
			_ = e11
			e11 = e1.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = e1.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.VTable.Hx_Field_expr(econd), ((any)(true))))) {
				return this.VTable.Hx_Field_expr(e11)
			} else {
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
					return nil
				} else {
					return this.VTable.Hx_Field_expr(e2)
				}
			}
		}

	case 10:
		{
			var econd any
			_ = econd
			econd = e1.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_cond
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
			this.VTable.Hx_Field_whileLoop(econd, e)
			return nil
		}

	case 11:
		{
			var v string
			_ = v
			v = e1.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_v
			var it any
			_ = it
			it = e1.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_it
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_e
			this.VTable.Hx_Field_forLoop(v, it, e)
			return nil
		}

	case 12:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_loreline_lorscript__interp_stop_SBreak{}))))
		}

	case 13:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_loreline_lorscript__interp_stop_SContinue{}))))
		}

	case 14:
		{
			var params HxArray[any]
			_ = params
			params = e1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_args
			var fexpr any
			_ = fexpr
			fexpr = e1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
			var name HxNullable[string]
			_ = name
			name = e1.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_name
			var capturedLocals *Hx_Obj_haxe_imap
			_ = capturedLocals
			capturedLocals = this.VTable.Hx_Field_duplicate(this.Hx_Field_locals)
			var me *Hx_Obj_loreline_lorscript_interp
			_ = me
			me = this
			var hasOpt bool
			_ = hasOpt
			hasOpt = false
			var minParams int32
			_ = minParams
			minParams = 0
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					if !(_hx_tmp_9 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
						break
					}

					var p any
					_ = p
					p = params.Get(_g)
					_g++
					var _hx_tmp_10 bool
					_ = _hx_tmp_10
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(p, "opt"))) {
						_hx_tmp_10 = true
					} else {
						_hx_tmp_10 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(p, "value"), nil))))
					}

					if _hx_tmp_10 {
						hasOpt = true
					} else {
						minParams = (minParams + 1)
					}
				}
			}

			var f func(args HxArray[any]) any
			_ = f
			f = func(args HxArray[any]) any {
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				if args == nil {
					_hx_tmp_11 = 0
				} else {
					_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_getArrayLength(args)
				}

				if (_hx_tmp_11) != Hx_Field_go_haxe_hxdynamic_getArrayLength(params) {
					if Hx_Field_go_haxe_hxdynamic_getArrayLength(args) < minParams {
						var _hx_tmp_12 string
						_ = _hx_tmp_12
						_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_stringConcat(("Invalid number of parameters. Got " + Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getArrayLength(args))), ", required ")
						var str string
						_ = str
						str = (_hx_tmp_12 + Hx_Field_go_haxe_hxdynamic_toString(minParams))
						if (name).Valid != false {
							var _hx_tmp_13 string
							_ = _hx_tmp_13
							_hx_tmp_13 = str
							str = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_13, Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(" for function '", (name).Value), "'"))
						}

						var e *Hx_Obj_loreline_lorscript_error
						_ = e
						e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{str}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(_gthis.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(_gthis.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(_gthis.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(_gthis.Hx_Field_curExpr, "line")))))
						panic(Hx_Field_haxe_exception_thrown(((any)(e))))
					}

					var args2 HxArray[any]
					_ = args2
					args2 = HxMakeArray[any]().Dyn()
					var extraParams int32
					_ = extraParams
					extraParams = (Hx_Field_go_haxe_hxdynamic_getArrayLength(args) - minParams)
					var pos int32
					_ = pos
					pos = 0
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							_hx_tmp_14 = _g
							if !(_hx_tmp_14 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
								break
							}

							var p any
							_ = p
							p = params.Get(_g)
							_g++
							var _hx_tmp_15 bool
							_ = _hx_tmp_15
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(p, "opt"))) {
								_hx_tmp_15 = true
							} else {
								_hx_tmp_15 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(p, "value"), nil))))
							}

							if _hx_tmp_15 {
								if extraParams > 0 {
									var _hx_tmp_16 HxArray[any]
									_ = _hx_tmp_16
									_hx_tmp_16 = HxMakeArrayView[any](args2)
									var _hx_tmp_17 HxArray[any]
									_ = _hx_tmp_17
									_hx_tmp_17 = args
									var _hx_tmp_18 int32
									_ = _hx_tmp_18
									_hx_tmp_18 = pos
									pos = (pos + 1)
									_ = (Hx_Array_Push[any](_hx_tmp_16, _hx_tmp_17.Get(_hx_tmp_18)))
									extraParams--
								} else {
									var _hx_tmp_19 HxArray[any]
									_ = _hx_tmp_19
									_hx_tmp_19 = HxMakeArrayView[any](args2)
									var _hx_tmp_20 any
									_ = _hx_tmp_20
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(p, "value"), nil))) {
										_hx_tmp_20 = me.VTable.Hx_Field_expr(Hx_Field_go_haxe_hxdynamic_getField(p, "value"))
									} else {
										_hx_tmp_20 = nil
									}

									_ = (Hx_Array_Push[any](_hx_tmp_19, _hx_tmp_20))
								}
							} else {
								var _hx_tmp_21 HxArray[any]
								_ = _hx_tmp_21
								_hx_tmp_21 = HxMakeArrayView[any](args2)
								var _hx_tmp_22 HxArray[any]
								_ = _hx_tmp_22
								_hx_tmp_22 = args
								var _hx_tmp_23 int32
								_ = _hx_tmp_23
								_hx_tmp_23 = pos
								pos = (pos + 1)
								_ = (Hx_Array_Push[any](_hx_tmp_21, _hx_tmp_22.Get(_hx_tmp_23)))
							}
						}
					}

					args = args2
				}

				var old Hx_Typedef_map
				_ = old
				old = me.Hx_Field_locals
				var depth int32
				_ = depth
				depth = me.Hx_Field_depth
				me.Hx_Field_depth++
				me.Hx_Field_locals = me.VTable.Hx_Field_duplicate(capturedLocals)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(params)
					for _g < _g1 {
						var _hx_tmp_24 int32
						_ = _hx_tmp_24
						_hx_tmp_24 = _g
						_g = (_g + 1)
						var i int32
						_ = i
						i = _hx_tmp_24
						var _hx_tmp_25 any
						_ = _hx_tmp_25
						_hx_tmp_25 = Hx_Field_go_haxe_hxdynamic_getField(params.Get(i), "name")
						me.Hx_Field_locals.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_25, any(map[string]any{"r": args.Get(i)}))
					}
				}

				var r any
				_ = r
				var oldDecl int32
				_ = oldDecl
				oldDecl = Hx_Field_go_haxe_hxdynamic_getArrayLength(_gthis.Hx_Field_declared)
				if _gthis.Hx_Field_inTry {
					{
						var hx_try_return any
						_ = hx_try_return
						var hx_try_state int
						_ = hx_try_state

						func() {
							defer func() {
								if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
									if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
										_ = e
										{
											_gthis.VTable.Hx_Field_restore(oldDecl)
											me.Hx_Field_locals = old
											me.Hx_Field_depth = depth
											panic(Hx_Field_haxe_exception_thrown(e))
										}
									} else {
										panic(_hx_ex)
									}
								}
							}()
							{
								r = me.VTable.Hx_Field_exprReturn(fexpr)
							}
						}()
						if hx_try_state == 1 {
							return hx_try_return
						}
					}
				} else {
					r = me.VTable.Hx_Field_exprReturn(fexpr)
				}

				_gthis.VTable.Hx_Field_restore(oldDecl)
				me.Hx_Field_locals = old
				me.Hx_Field_depth = depth
				return r
			}

			var f1 any
			_ = f1
			f1 = Hx_Field_reflect_makeVarArgs(f)
			if (name).Valid != false {
				if this.Hx_Field_depth == 0 {
					var hx_nullable_f203 HxNullable[string]
					_ = hx_nullable_f203
					hx_nullable_f203 = name
					var _hx_tmp_26 any
					_ = _hx_tmp_26
					if (hx_nullable_f203).Valid == true {
						_hx_tmp_26 = (hx_nullable_f203).Value
					} else {
						_hx_tmp_26 = nil
					}

					this.Hx_Field_variables.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_26, f1)
				} else {
					var _hx_tmp_27 HxArray[any]
					_ = _hx_tmp_27
					_hx_tmp_27 = HxMakeArrayView[any](this.Hx_Field_declared)
					var hx_nullable_f207 HxNullable[string]
					_ = hx_nullable_f207
					hx_nullable_f207 = name
					var _hx_tmp_28 any
					_ = _hx_tmp_28
					if (hx_nullable_f207).Valid == true {
						_hx_tmp_28 = (hx_nullable_f207).Value
					} else {
						_hx_tmp_28 = nil
					}

					var hx_nullable_f204 HxNullable[string]
					_ = hx_nullable_f204
					hx_nullable_f204 = name
					var _hx_tmp_29 any
					_ = _hx_tmp_29
					if (hx_nullable_f204).Valid == true {
						_hx_tmp_29 = (hx_nullable_f204).Value
					} else {
						_hx_tmp_29 = nil
					}

					_ = (Hx_Array_Push[any](_hx_tmp_27, any(map[string]any{"n": _hx_tmp_28, "old": this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(_hx_tmp_29)})))
					var ref any
					_ = ref
					ref = any(map[string]any{"r": f1})
					var hx_nullable_f205 HxNullable[string]
					_ = hx_nullable_f205
					hx_nullable_f205 = name
					var _hx_tmp_30 any
					_ = _hx_tmp_30
					if (hx_nullable_f205).Valid == true {
						_hx_tmp_30 = (hx_nullable_f205).Value
					} else {
						_hx_tmp_30 = nil
					}

					this.Hx_Field_locals.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_30, ref)
					var hx_nullable_f206 HxNullable[string]
					_ = hx_nullable_f206
					hx_nullable_f206 = name
					var _hx_tmp_31 any
					_ = _hx_tmp_31
					if (hx_nullable_f206).Valid == true {
						_hx_tmp_31 = (hx_nullable_f206).Value
					} else {
						_hx_tmp_31 = nil
					}

					capturedLocals.VTable.Hx_Field_set(_hx_tmp_31, ref)
				}
			}

			return f1
		}

	case 15:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
			var _hx_tmp_32 any
			_ = _hx_tmp_32
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_32 = nil
			} else {
				_hx_tmp_32 = this.VTable.Hx_Field_expr(e)
			}

			this.Hx_Field_returnValue = _hx_tmp_32
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_loreline_lorscript__interp_stop_SReturn{}))))
		}

	case 16:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
			var index any
			_ = index
			index = e1.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
			var arr any
			_ = arr
			arr = this.VTable.Hx_Field_expr(e)
			var index1 any
			_ = index1
			index1 = this.VTable.Hx_Field_expr(index)
			if Hx_Field_loreline_arrays_isArray(arr) {
				var _hx_tmp_33 any
				_ = _hx_tmp_33
				_hx_tmp_33 = arr
				return Hx_Field_loreline_arrays_arrayGet(_hx_tmp_33, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))))
			} else {
				if Hx_Field_loreline_objects_isFields(arr) {
					var _hx_tmp_34 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_34
					_hx_tmp_34 = this.Hx_Field_interpreter
					var _hx_tmp_35 any
					_ = _hx_tmp_35
					_hx_tmp_35 = arr
					return Hx_Field_loreline_objects_getField(_hx_tmp_34, _hx_tmp_35, Hx_Field_go_haxe_hxdynamic_toString(index1))
				} else {
					if Hx_Field_std_isOfType(arr, Hx_Obj_haxe_imap_RTTI) {
						var hx_cast_ex_197 any
						_ = hx_cast_ex_197
						hx_cast_ex_197 = arr
						if !Hx_Field_std_isOfType(hx_cast_ex_197, Hx_Obj_haxe_imap_RTTI) {
							panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
						}

						var hx_dyncast_198 any
						_ = hx_dyncast_198
						hx_dyncast_198 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_197, "Hx_Obj_haxe_imap")
						var _hx_tmp_36 *Hx_Obj_haxe_imap
						_ = _hx_tmp_36
						if hx_dyncast_198 != nil {
							_hx_tmp_36 = hx_dyncast_198.(*Hx_Obj_haxe_imap)
						} else {
							_hx_tmp_36 = nil
						}

						return (_hx_tmp_36).VTable.Hx_Field_get(index1)
					}
				}
			}

			var _hx_tmp_37 any
			_ = _hx_tmp_37
			_hx_tmp_37 = arr
			return Hx_Field_go_haxe_hxdynamic_getArrayIndex(_hx_tmp_37, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(index1))))
		}

	case 17:
		{
			var arr HxArray[any]
			_ = arr
			arr = e1.(Hx_Enum_loreline_lorscript_exprdef_EArrayDecl).Hx_Field_e
			var tmp bool
			_ = tmp
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(arr) > 0 {
				var _g Hx_Enum_loreline_lorscript_exprdef
				_ = _g
				_g = func() Hx_Enum_loreline_lorscript_exprdef {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arr.Get(0), "e"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_exprdef
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
				}()
				var _hx_tmp_38 bool
				_ = _hx_tmp_38
				if _g.Hx_Field_enumIndex() == 6 {
					_hx_tmp_38 = (_g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op == "=>")
				} else {
					_hx_tmp_38 = false
				}

				tmp = _hx_tmp_38
			} else {
				tmp = false
			}

			if tmp {
				var keys HxArray[any]
				_ = keys
				keys = HxMakeArray[any]().Dyn()
				var values HxArray[any]
				_ = values
				values = HxMakeArray[any]().Dyn()
				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_39 int32
						_ = _hx_tmp_39
						_hx_tmp_39 = _g
						if !(_hx_tmp_39 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
							break
						}

						var e any
						_ = e
						e = arr.Get(_g)
						_g++
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
							if _g1.Hx_Field_enumIndex() == 6 {
								if _g1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op == "=>" {
									var eKey any
									_ = eKey
									eKey = _g1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
									var eValue any
									_ = eValue
									eValue = _g1.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
									{
										var _hx_tmp_40 HxArray[any]
										_ = _hx_tmp_40
										_hx_tmp_40 = keys
										_ = (Hx_Array_Push[any](_hx_tmp_40, this.VTable.Hx_Field_expr(eKey)))
										var _hx_tmp_41 HxArray[any]
										_ = _hx_tmp_41
										_hx_tmp_41 = values
										_ = (Hx_Array_Push[any](_hx_tmp_41, this.VTable.Hx_Field_expr(eValue)))
									}
								} else {
									this.Hx_Field_curExpr = e
									var tmp_e1_1 *Hx_Obj_loreline_lorscript_error
									_ = tmp_e1_1
									tmp_e1_1 = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"Invalid map key=>value expression"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
									panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e1_1))))
								}
							} else {
								this.Hx_Field_curExpr = e
								var tmp_e1_1 *Hx_Obj_loreline_lorscript_error
								_ = tmp_e1_1
								tmp_e1_1 = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"Invalid map key=>value expression"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
								panic(Hx_Field_haxe_exception_thrown(((any)(tmp_e1_1))))
							}
						}
					}
				}

				return this.VTable.Hx_Field_makeMap(keys, values)
			} else {
				var a HxArray[any]
				_ = a
				a = HxMakeArray[any]().Dyn()
				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_42 int32
						_ = _hx_tmp_42
						_hx_tmp_42 = _g
						if !(_hx_tmp_42 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
							break
						}

						var e any
						_ = e
						e = arr.Get(_g)
						_g++
						var _hx_tmp_43 HxArray[any]
						_ = _hx_tmp_43
						_hx_tmp_43 = a
						_ = (Hx_Array_Push[any](_hx_tmp_43, this.VTable.Hx_Field_expr(e)))
					}
				}

				return ((any)(a))
			}
		}

	case 18:
		{
			var cl string
			_ = cl
			cl = e1.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_cl
			var params HxArray[any]
			_ = params
			params = e1.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_params
			var a HxArray[any]
			_ = a
			a = HxMakeArray[any]().Dyn()
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_44 int32
					_ = _hx_tmp_44
					_hx_tmp_44 = _g
					if !(_hx_tmp_44 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
						break
					}

					var e any
					_ = e
					e = params.Get(_g)
					_g++
					var _hx_tmp_45 HxArray[any]
					_ = _hx_tmp_45
					_hx_tmp_45 = a
					_ = (Hx_Array_Push[any](_hx_tmp_45, this.VTable.Hx_Field_expr(e)))
				}
			}

			return this.VTable.Hx_Field_cnew(cl, a)
		}

	case 19:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EThrow).Hx_Field_e
			panic(Hx_Field_haxe_exception_thrown(this.VTable.Hx_Field_expr(e)))
		}

	case 20:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_e
			var n string
			_ = n
			n = e1.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_v
			var ecatch any
			_ = ecatch
			ecatch = e1.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
			var old int32
			_ = old
			old = Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared)
			var oldTry bool
			_ = oldTry
			oldTry = this.Hx_Field_inTry
			{
				var hx_try_return any
				_ = hx_try_return
				var hx_try_state int
				_ = hx_try_state

				func() {
					defer func() {
						if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
							if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
								_ = err
								{
									this.Hx_Field_inTry = oldTry
									panic(Hx_Field_haxe_exception_thrown(((any)(err))))
								}
							} else if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
								_ = err
								{
									this.VTable.Hx_Field_restore(old)
									this.Hx_Field_inTry = oldTry
									var _hx_tmp_46 HxArray[any]
									_ = _hx_tmp_46
									_hx_tmp_46 = HxMakeArrayView[any](this.Hx_Field_declared)
									var _hx_tmp_47 any
									_ = _hx_tmp_47
									_hx_tmp_47 = ((any)(n))
									_ = (Hx_Array_Push[any](_hx_tmp_46, any(map[string]any{"n": _hx_tmp_47, "old": this.Hx_Field_locals.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
									}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(n)})))
									var _hx_tmp_48 string
									_ = _hx_tmp_48
									_hx_tmp_48 = n
									this.Hx_Field_locals.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
									}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_48, any(map[string]any{"r": err}))
									var v any
									_ = v
									v = this.VTable.Hx_Field_expr(ecatch)
									this.VTable.Hx_Field_restore(old)
									hx_try_state = 1
									hx_try_return = v
									return
								}
							} else {
								panic(_hx_ex)
							}
						}
					}()
					{
						this.Hx_Field_inTry = true
						var v any
						_ = v
						v = this.VTable.Hx_Field_expr(e)
						this.VTable.Hx_Field_restore(old)
						this.Hx_Field_inTry = oldTry
						hx_try_state = 1
						hx_try_return = v
						return
					}
				}()

				return hx_try_return
			}
		}

	case 21:
		{
			var fl HxArray[any]
			_ = fl
			fl = e1.(Hx_Enum_loreline_lorscript_exprdef_EObject).Hx_Field_fl
			var _hx_tmp_49 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_49
			_hx_tmp_49 = this.Hx_Field_interpreter
			var _hx_tmp_50 any
			_ = _hx_tmp_50
			_hx_tmp_50 = nil
			var o any
			_ = o
			o = Hx_Field_loreline_objects_createFields(_hx_tmp_49, _hx_tmp_50, nil)
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_51 int32
					_ = _hx_tmp_51
					_hx_tmp_51 = _g
					if !(_hx_tmp_51 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
						break
					}

					var f any
					_ = f
					f = fl.Get(_g)
					_g++
					var _hx_tmp_52 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_52
					_hx_tmp_52 = this.Hx_Field_interpreter
					var _hx_tmp_53 any
					_ = _hx_tmp_53
					_hx_tmp_53 = o
					var _hx_tmp_54 string
					_ = _hx_tmp_54
					_hx_tmp_54 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name"))
					Hx_Field_loreline_objects_setField(_hx_tmp_52, _hx_tmp_53, _hx_tmp_54, this.VTable.Hx_Field_expr(Hx_Field_go_haxe_hxdynamic_getField(f, "e")))
				}
			}

			return o
		}

	case 22:
		{
			var econd any
			_ = econd
			econd = e1.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
			var e11 any
			_ = e11
			e11 = e1.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = e1.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.VTable.Hx_Field_expr(econd), ((any)(true))))) {
				return this.VTable.Hx_Field_expr(e11)
			} else {
				return this.VTable.Hx_Field_expr(e2)
			}
		}

	case 23:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_e
			var cases HxArray[any]
			_ = cases
			cases = e1.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_cases
			var def any
			_ = def
			def = e1.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_defaultExpr
			var val any
			_ = val
			val = this.VTable.Hx_Field_expr(e)
			var match bool
			_ = match
			match = false
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_55 int32
					_ = _hx_tmp_55
					_hx_tmp_55 = _g
					if !(_hx_tmp_55 < Hx_Field_go_haxe_hxdynamic_getArrayLength(cases)) {
						break
					}

					var c any
					_ = c
					c = cases.Get(_g)
					_g++
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[any]
						_ = _g2
						_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "values"))
						for {
							var _hx_tmp_56 int32
							_ = _hx_tmp_56
							_hx_tmp_56 = _g1
							if !(_hx_tmp_56 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
								break
							}

							var v any
							_ = v
							v = _g2.Get(_g1)
							_g1++
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.VTable.Hx_Field_expr(v), val))) {
								match = true
								break
							}
						}
					}

					if match {
						val = this.VTable.Hx_Field_expr(Hx_Field_go_haxe_hxdynamic_getField(c, "expr"))
						break
					}
				}
			}

			if !match {
				var _hx_tmp_57 any
				_ = _hx_tmp_57
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(def, nil))) {
					_hx_tmp_57 = nil
				} else {
					_hx_tmp_57 = this.VTable.Hx_Field_expr(def)
				}

				val = _hx_tmp_57
			}

			return val
		}

	case 24:
		{
			var econd any
			_ = econd
			econd = e1.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_cond
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
			this.VTable.Hx_Field_doWhileLoop(econd, e)
			return nil
		}

	case 25:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
			return this.VTable.Hx_Field_expr(e)
		}

	case 26:
		{
			var e any
			_ = e
			e = e1.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e
			return this.VTable.Hx_Field_expr(e)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_doWhileLoop(econd any, e any) {
	var old int32
	_ = old
	old = Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared)
	var _gthis *Hx_Obj_loreline_lorscript_interp
	_ = _gthis
	_gthis = this
	for {
		var cont bool
		_ = cont
		cont = true
		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
							_ = err
							{
								switch err.Hx_Field_enumIndex() {
								case 0:
									{
										cont = false
									}

								case 1:
									{
									}

								case 2:
									{
										panic(Hx_Field_haxe_exception_thrown(((any)(err))))
									}

								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					_gthis.VTable.Hx_Field_expr(e)
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

		if !cont {
			break
		}

		if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.VTable.Hx_Field_expr(econd), ((any)(true)))))) {
			break
		}
	}

	this.VTable.Hx_Field_restore(old)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_whileLoop(econd any, e any) {
	var old int32
	_ = old
	old = Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared)
	var _gthis *Hx_Obj_loreline_lorscript_interp
	_ = _gthis
	_gthis = this
	for {
		if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.VTable.Hx_Field_expr(econd), ((any)(true)))))) {
			break
		}

		var cont bool
		_ = cont
		cont = true
		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
							_ = err
							{
								switch err.Hx_Field_enumIndex() {
								case 0:
									{
										cont = false
									}

								case 1:
									{
									}

								case 2:
									{
										panic(Hx_Field_haxe_exception_thrown(((any)(err))))
									}

								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					_gthis.VTable.Hx_Field_expr(e)
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

		if !cont {
			break
		}
	}

	this.VTable.Hx_Field_restore(old)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_makeIterator(v any) any {
	if Hx_Field_loreline_arrays_isArray(v) {
		return Hx_Field_loreline_arrays_getIterator(v)
	}

	{
		var hx_try_return any
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 any
				_ = _hx_tmp_0
				_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(v, "iterator")
				v = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(v, "hasNext"), nil))) {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(v, "next"), nil))))
	}

	if _hx_tmp_1 {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidIterator{Hx_Field_go_haxe_hxdynamic_toString(v)}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_forLoop(n string, it any, e any) {
	var old int32
	_ = old
	old = Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_declared)
	var _hx_tmp_0 HxArray[any]
	_ = _hx_tmp_0
	_hx_tmp_0 = HxMakeArrayView[any](this.Hx_Field_declared)
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = ((any)(n))
	_ = (Hx_Array_Push[any](_hx_tmp_0, any(map[string]any{"n": _hx_tmp_1, "old": this.Hx_Field_locals.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(n)})))
	var it1 any
	_ = it1
	it1 = this.VTable.Hx_Field_makeIterator(this.VTable.Hx_Field_expr(it))
	var _gthis *Hx_Obj_loreline_lorscript_interp
	_ = _gthis
	_gthis = this
	for {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(it1, "hasNext")
		if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
			break
		}

		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = n
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(it1, "next")
		this.Hx_Field_locals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_3, any(map[string]any{"r": Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn())}))
		var cont bool
		_ = cont
		cont = true
		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
							_ = err
							{
								switch err.Hx_Field_enumIndex() {
								case 0:
									{
										cont = false
									}

								case 1:
									{
									}

								case 2:
									{
										panic(Hx_Field_haxe_exception_thrown(((any)(err))))
									}

								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					_gthis.VTable.Hx_Field_expr(e)
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

		if !cont {
			break
		}
	}

	this.VTable.Hx_Field_restore(old)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_loopRun(f func()) bool {
	var cont bool
	_ = cont
	cont = true
	{
		var hx_try_return bool
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if err, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "Hx_Enum_loreline_lorscript__interp_stop").(Hx_Enum_loreline_lorscript__interp_stop); _hx_ok {
						_ = err
						{
							switch err.Hx_Field_enumIndex() {
							case 0:
								{
									cont = false
								}

							case 1:
								{
								}

							case 2:
								{
									panic(Hx_Field_haxe_exception_thrown(((any)(err))))
								}

							}
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				f()
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	return cont
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_isMap(o any) bool {
	return Hx_Field_std_isOfType(o, Hx_Obj_haxe_imap_RTTI)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_getMapValue(_hx_reserved_map any, key any) any {
	var hx_cast_ex_199 any
	_ = hx_cast_ex_199
	hx_cast_ex_199 = _hx_reserved_map
	if !Hx_Field_std_isOfType(hx_cast_ex_199, Hx_Obj_haxe_imap_RTTI) {
		panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
	}

	var hx_dyncast_200 any
	_ = hx_dyncast_200
	hx_dyncast_200 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_199, "Hx_Obj_haxe_imap")
	var _hx_tmp_0 *Hx_Obj_haxe_imap
	_ = _hx_tmp_0
	if hx_dyncast_200 != nil {
		_hx_tmp_0 = hx_dyncast_200.(*Hx_Obj_haxe_imap)
	} else {
		_hx_tmp_0 = nil
	}

	return (_hx_tmp_0).VTable.Hx_Field_get(key)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_setMapValue(_hx_reserved_map any, key any, value any) {
	var hx_cast_ex_201 any
	_ = hx_cast_ex_201
	hx_cast_ex_201 = _hx_reserved_map
	if !Hx_Field_std_isOfType(hx_cast_ex_201, Hx_Obj_haxe_imap_RTTI) {
		panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
	}

	var hx_dyncast_202 any
	_ = hx_dyncast_202
	hx_dyncast_202 = Hx_Field_go_haxe_hxdynamic_castClass(hx_cast_ex_201, "Hx_Obj_haxe_imap")
	var _hx_tmp_0 *Hx_Obj_haxe_imap
	_ = _hx_tmp_0
	if hx_dyncast_202 != nil {
		_hx_tmp_0 = hx_dyncast_202.(*Hx_Obj_haxe_imap)
	} else {
		_hx_tmp_0 = nil
	}

	(_hx_tmp_0).VTable.Hx_Field_set(key, value)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_makeMap(keys HxArray[any], values HxArray[any]) any {
	var isAllString bool
	_ = isAllString
	isAllString = true
	var isAllInt bool
	_ = isAllInt
	isAllInt = true
	var isAllObject bool
	_ = isAllObject
	isAllObject = true
	var isAllEnum bool
	_ = isAllEnum
	isAllEnum = true
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(keys)) {
				break
			}

			var key any
			_ = key
			key = keys.Get(_g)
			_g++
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if isAllString {
				_hx_tmp_1 = Hx_Field_std_isOfType(key, Hx_Obj_string_RTTI)
			} else {
				_hx_tmp_1 = false
			}

			isAllString = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if isAllInt {
				_hx_tmp_2 = Hx_Field_std_isOfType(key, Hx_Abstract_int_RTTI)
			} else {
				_hx_tmp_2 = false
			}

			isAllInt = _hx_tmp_2
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if isAllObject {
				_hx_tmp_3 = Hx_Field_reflect_isObject(key)
			} else {
				_hx_tmp_3 = false
			}

			isAllObject = _hx_tmp_3
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if isAllEnum {
				_hx_tmp_4 = Hx_Field_reflect_isEnumValue(key)
			} else {
				_hx_tmp_4 = false
			}

			isAllEnum = _hx_tmp_4
		}
	}

	if isAllInt {
		var m *Hx_Obj_haxe_imap
		_ = m
		m = (func() *Hx_Obj_haxe_imap {
			_hx_icast_203 := Hx_Obj_haxe_ds_intmap_CreateInstance()
			if _hx_icast_203 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_203, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_203.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		{
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[any]
			_ = _g_array
			_g_array = keys
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g_current
				if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g_array)) {
					break
				}

				var _g_value any
				_ = _g_value
				_g_value = _g_array.Get(_g_current)
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = _g_current
				_g_current = (_g_current + 1)
				var _g_key int32
				_ = _g_key
				_g_key = _hx_tmp_6
				var i int32
				_ = i
				i = _g_key
				var key any
				_ = key
				key = _g_value
				m.VTable.Hx_Field_set(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(key))), values.Get(i))
			}
		}

		return any(m.VTable)
	}

	if isAllString {
		var m *Hx_Obj_haxe_imap
		_ = m
		m = (func() *Hx_Obj_haxe_imap {
			_hx_icast_204 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_204 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_204, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_204.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		{
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[any]
			_ = _g_array
			_g_array = keys
			for {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _g_current
				if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g_array)) {
					break
				}

				var _g_value any
				_ = _g_value
				_g_value = _g_array.Get(_g_current)
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _g_current
				_g_current = (_g_current + 1)
				var _g_key int32
				_ = _g_key
				_g_key = _hx_tmp_8
				var i int32
				_ = i
				i = _g_key
				var key any
				_ = key
				key = _g_value
				m.VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_toString(key), values.Get(i))
			}
		}

		return any(m.VTable)
	}

	if isAllEnum {
		var m *Hx_Obj_haxe_ds_enumvaluemap
		_ = m
		m = Hx_Obj_haxe_ds_enumvaluemap_CreateInstance()
		{
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[any]
			_ = _g_array
			_g_array = keys
			for {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g_current
				if !(_hx_tmp_9 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g_array)) {
					break
				}

				var _g_value any
				_ = _g_value
				_g_value = _g_array.Get(_g_current)
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = _g_current
				_g_current = (_g_current + 1)
				var _g_key int32
				_ = _g_key
				_g_key = _hx_tmp_10
				var i int32
				_ = i
				i = _g_key
				var key any
				_ = key
				key = _g_value
				m.VTable.Hx_Field_set(key, values.Get(i))
			}
		}

		return ((any)(m))
	}

	if isAllObject {
		var m *Hx_Obj_haxe_imap
		_ = m
		m = (func() *Hx_Obj_haxe_imap {
			_hx_icast_205 := Hx_Obj_haxe_ds_objectmap_CreateInstance()
			if _hx_icast_205 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_205, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_205.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		{
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[any]
			_ = _g_array
			_g_array = keys
			for {
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _g_current
				if !(_hx_tmp_11 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g_array)) {
					break
				}

				var _g_value any
				_ = _g_value
				_g_value = _g_array.Get(_g_current)
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				_hx_tmp_12 = _g_current
				_g_current = (_g_current + 1)
				var _g_key int32
				_ = _g_key
				_g_key = _hx_tmp_12
				var i int32
				_ = i
				i = _g_key
				var key any
				_ = key
				key = _g_value
				m.VTable.Hx_Field_set(key, values.Get(i))
			}
		}

		return any(m.VTable)
	}

	var e *Hx_Obj_loreline_lorscript_error
	_ = e
	e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid map keys ", Hx_Field_std_string(keys))}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
	panic(Hx_Field_haxe_exception_thrown(((any)(e))))
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_get(o any, f string) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(o, nil))) {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidAccess{f}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	o = Hx_Field_loreline_runtimecharacterref_fieldsOf(o)
	if Hx_Field_loreline_objects_isFields(o) {
		return Hx_Field_loreline_objects_getField(this.Hx_Field_interpreter, o, f)
	}

	return Hx_Field_reflect_getProperty(o, f)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_set(o any, f string, v any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(o, nil))) {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidAccess{f}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	o = Hx_Field_loreline_runtimecharacterref_fieldsOf(o)
	if Hx_Field_loreline_objects_isFields(o) {
		Hx_Field_loreline_objects_setField(this.Hx_Field_interpreter, o, f, v)
		return v
	}

	Hx_Field_reflect_setProperty(o, f, v)
	return v
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_fcall(o any, f string, args HxArray[any]) any {
	o = Hx_Field_loreline_runtimecharacterref_fieldsOf(o)
	var helper any
	_ = helper
	var oBeat *Hx_Obj_loreline_nbeatdecl
	_ = oBeat
	oBeat = Hx_Field_loreline_runtimebeatref_beatOf(o)
	if Hx_Field_std_isOfType(o, Hx_Obj_string_RTTI) {
		helper = Hx_Field_loreline_objects_getStringHelper(this.Hx_Field_interpreter, f)
	} else {
		if Hx_Field_loreline_arrays_isArray(o) {
			helper = Hx_Field_loreline_objects_getArrayHelper(this.Hx_Field_interpreter, f)
		} else {
			if oBeat != nil {
				helper = Hx_Field_loreline_objects_getBeatHelper(this.Hx_Field_interpreter, f)
			} else {
				if Hx_Field_loreline_objects_isFields(o) {
					helper = Hx_Field_loreline_objects_getMapHelper(this.Hx_Field_interpreter, f)
				}
			}
		}
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(helper, nil))) {
		_hx_tmp_0 = Hx_Field_reflect_isFunction(helper)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var _hx_tmp_1 HxArray[any]
		_ = _hx_tmp_1
		_hx_tmp_1 = args
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		if oBeat != nil {
			_hx_tmp_2 = ((any)(oBeat))
		} else {
			_hx_tmp_2 = o
		}

		(Hx_Array_Insert[any](_hx_tmp_1, 0, _hx_tmp_2))
		return this.VTable.Hx_Field_call(nil, helper, args)
	}

	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = o
	return this.VTable.Hx_Field_call(_hx_tmp_3, this.VTable.Hx_Field_get(o, f), args)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_call(o any, f any, args HxArray[any]) any {
	return Hx_Field_reflect_callMethod(o, f, args)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_callIsShadowed() bool {
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_578 func(_hx_param_key any) bool
		_ = _hx_closure_src_578
		_hx_closure_src_578 = this.Hx_Field_locals.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_578(((any)(_hx_param_key)))
		}
	}()("call") {
		return true
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_579 func(_hx_param_key any) bool
		_ = _hx_closure_src_579
		_hx_closure_src_579 = this.Hx_Field_variables.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_579(((any)(_hx_param_key)))
		}
	}()("call") {
		return true
	}

	if Hx_Field_loreline_objects_fieldExists(this.Hx_Field_interpreter, this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, "call") {
		return true
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_580 func(_hx_param_key any) bool
		_ = _hx_closure_src_580
		_hx_closure_src_580 = this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_580(((any)(_hx_param_key)))
		}
	}()("call") {
		return true
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_581 func(_hx_param_key any) bool
		_ = _hx_closure_src_581
		_hx_closure_src_581 = this.Hx_Field_interpreter.Hx_Field_topLevelFunctions.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_581(((any)(_hx_param_key)))
		}
	}()("call") {
		return true
	}

	if this.Hx_Field_interpreter.VTable.Hx_Field_resolveBeatByName("call", nil) != nil {
		return true
	}

	return false
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_dynamicCall(args HxArray[any]) any {
	if Hx_Field_go_haxe_hxdynamic_getArrayLength(args) == 0 {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"call() requires a target argument"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	var target any
	_ = target
	target = args.Get(0)
	var _hx_tmp_0 HxArray[any]
	_ = _hx_tmp_0
	_hx_tmp_0 = args
	var rest HxArray[any]
	_ = rest
	rest = (Hx_Array_Slice[any](_hx_tmp_0, 1, HxNullable[int32]{}))
	if Hx_Field_loreline_runtimebeatref_beatOf(target) != nil {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"call() can only call functions. Beats cannot run inside function code."}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	if Hx_Field_std_isOfType(target, Hx_Obj_string_RTTI) {
		var name string
		_ = name
		name = Hx_Field_go_haxe_hxdynamic_toString(target)
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = name
		if this.Hx_Field_interpreter.VTable.Hx_Field_resolveBeatByName(_hx_tmp_1, nil) != nil {
			var e *Hx_Obj_loreline_lorscript_error
			_ = e
			e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"call() can only call functions. Beats cannot run inside function code."}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		}

		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_582 func(_hx_param_key any) bool
			_ = _hx_closure_src_582
			_hx_closure_src_582 = this.Hx_Field_interpreter.Hx_Field_topLevelFunctions.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_582(((any)(_hx_param_key)))
			}
		}()(name) {
			target = this.Hx_Field_interpreter.Hx_Field_topLevelFunctions.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
		} else {
			var e *Hx_Obj_loreline_lorscript_error
			_ = e
			e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{Hx_Field_go_haxe_hxdynamic_stringConcat("call() target not found: ", name)}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		}
	}

	if !Hx_Field_reflect_isFunction(target) {
		var e *Hx_Obj_loreline_lorscript_error
		_ = e
		e = Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"call() target is not a function"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "pmax")))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "origin")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_curExpr, "line")))))
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}

	return this.VTable.Hx_Field_call(nil, target, rest)
}

func (this *Hx_Obj_loreline_lorscript_interp) Hx_Field_cnew(cl string, args HxArray[any]) any {
	var c *Hx_Obj_go_haxe_hxclass
	_ = c
	c = Hx_Field_type_resolveClass(cl)
	if c == nil {
		c = func() *Hx_Obj_go_haxe_hxclass {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_resolve(cl))
			if _hx_d == nil {
				var _hx_z *Hx_Obj_go_haxe_hxclass
				return _hx_z
			}
			return _hx_d.(*Hx_Obj_go_haxe_hxclass)
		}()
	}

	return Hx_Field_type_createInstance(c, args)
}

func Hx_Field_loreline_lorscript_interp___keepIntIterator() HxArray[any] {
	var it *Hx_Obj_intiterator
	_ = it
	it = Hx_Obj_intiterator_CreateInstance(0, 0)
	return HxMakeArray[any](((any)(it.VTable.Hx_Field_hasNext)), ((any)(it.VTable.Hx_Field_next))).Dyn()
}

func init() {
	hxRegisterArrayConstructor[any]()
}
