package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_tools() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_tools() {
}

type Hx_Obj_VTable_loreline_lorscript_tools interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_tools struct {
	VTable Hx_Obj_VTable_loreline_lorscript_tools `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_lorscript_tools_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_tools {
	obj := &Hx_Obj_loreline_lorscript_tools{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_lorscript_tools_CreateInstance() *Hx_Obj_loreline_lorscript_tools {
	obj := Hx_Obj_loreline_lorscript_tools_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_lorscript_tools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Tools",
	HxMakeArray[string]("iter", "map", "expr", "mk"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_tools_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lorscript_tools_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_tools_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "iter":
			return Hx_Field_loreline_lorscript_tools_iter
		case "map":
			return Hx_Field_loreline_lorscript_tools_map
		case "expr":
			return Hx_Field_loreline_lorscript_tools_expr
		case "mk":
			return Hx_Field_loreline_lorscript_tools_mk
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_tools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_tools_RTTI
}

func (this *Hx_Obj_loreline_lorscript_tools) Hx_As_Hx_Obj_loreline_lorscript_tools() *Hx_Obj_loreline_lorscript_tools {
	return this
}

func Hx_Field_loreline_lorscript_tools_iter(e any, f func(_hx_carg0 any)) {
	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 0:
		{
		}

	case 1:
		{
		}

	case 2:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_e
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				f(e)
			}
		}

	case 3:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
			f(e)
		}

	case 4:
		{
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
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
					f(e)
				}
			}
		}

	case 5:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
			f(e)
		}

	case 6:
		{
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			{
				f(e1)
				f(e2)
			}
		}

	case 7:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
			f(e)
		}

	case 8:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_e
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_params
			{
				f(e)
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_1 int32
						_ = _hx_tmp_1
						_hx_tmp_1 = _g1
						if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
							break
						}

						var a any
						_ = a
						a = args.Get(_g1)
						_g1++
						f(a)
					}
				}
			}
		}

	case 9:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_cond
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
			{
				f(c)
				f(e1)
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e2, nil))) {
					f(e2)
				}
			}
		}

	case 10:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_cond
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
			{
				f(c)
				f(e)
			}
		}

	case 11:
		{
			var it any
			_ = it
			it = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_it
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_e
			{
				f(it)
				f(e)
			}
		}

	case 12, 13:
		{
		}

	case 14:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
			f(e)
		}

	case 15:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				f(e)
			}
		}

	case 16:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
			var i any
			_ = i
			i = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
			{
				f(e)
				f(i)
			}
		}

	case 17:
		{
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_EArrayDecl).Hx_Field_e
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g1
					if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
						break
					}

					var e any
					_ = e
					e = el.Get(_g1)
					_g1++
					f(e)
				}
			}
		}

	case 18:
		{
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_params
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g1
					if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
						break
					}

					var e any
					_ = e
					e = el.Get(_g1)
					_g1++
					f(e)
				}
			}
		}

	case 19:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EThrow).Hx_Field_e
			f(e)
		}

	case 20:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_e
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
			{
				f(e)
				f(c)
			}
		}

	case 21:
		{
			var fl HxArray[any]
			_ = fl
			fl = _g.(Hx_Enum_loreline_lorscript_exprdef_EObject).Hx_Field_fl
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g1
					if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
						break
					}

					var fi any
					_ = fi
					fi = fl.Get(_g1)
					_g1++
					f(Hx_Field_go_haxe_hxdynamic_getField(fi, "e"))
				}
			}
		}

	case 22:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
			{
				f(c)
				f(e1)
				f(e2)
			}
		}

	case 23:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_e
			var cases HxArray[any]
			_ = cases
			cases = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_cases
			var def any
			_ = def
			def = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_defaultExpr
			{
				f(e)
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = _g1
						if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(cases)) {
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
								var _hx_tmp_6 int32
								_ = _hx_tmp_6
								_hx_tmp_6 = tmp__g_1
								if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
									break
								}

								var v any
								_ = v
								v = _g2.Get(tmp__g_1)
								tmp__g_1++
								f(v)
							}
						}

						f(Hx_Field_go_haxe_hxdynamic_getField(c, "expr"))
					}
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(def, nil))) {
					f(def)
				}
			}
		}

	case 24:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_cond
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
			{
				f(c)
				f(e)
			}
		}

	case 25:
		{
			var name string
			_ = name
			name = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_name
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_args
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
			{
				if args != nil {
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					for {
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = tmp__g_1
						if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
							break
						}

						var a any
						_ = a
						a = args.Get(tmp__g_1)
						tmp__g_1++
						f(a)
					}
				}

				f(e)
			}
		}

	case 26:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e
			f(e)
		}

	}
}

func Hx_Field_loreline_lorscript_tools_map(e any, f func(_hx_carg0 any) any) any {
	var edef Hx_Enum_loreline_lorscript_exprdef
	_ = edef
	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 0:
		{
			edef = func() Hx_Enum_loreline_lorscript_exprdef {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_exprdef
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
			}()
		}

	case 1:
		{
			edef = func() Hx_Enum_loreline_lorscript_exprdef {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_exprdef
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
			}()
		}

	case 2:
		{
			var n string
			_ = n
			n = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_n
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			t = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_t
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_e
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = n
			var _hx_tmp_1 Hx_Enum_loreline_lorscript_ctype
			_ = _hx_tmp_1
			_hx_tmp_1 = t
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				_hx_tmp_2 = f(e)
			} else {
				_hx_tmp_2 = nil
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EVar{_hx_tmp_0, _hx_tmp_1, _hx_tmp_2}
		}

	case 3:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
			edef = Hx_Enum_loreline_lorscript_exprdef_EParent{f(e)}
		}

	case 4:
		{
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g2
					if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
						break
					}

					var e any
					_ = e
					e = el.Get(_g2)
					_g2++
					var _hx_tmp_4 HxArray[any]
					_ = _hx_tmp_4
					_hx_tmp_4 = HxMakeArrayView[any](_g1)
					_ = (Hx_Array_Push[any](_hx_tmp_4, f(e)))
				}
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EBlock{_g1}
		}

	case 5:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
			var fi string
			_ = fi
			fi = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
			edef = Hx_Enum_loreline_lorscript_exprdef_EField{f(e), fi}
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
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = op
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = f(e1)
			edef = Hx_Enum_loreline_lorscript_exprdef_EBinop{_hx_tmp_5, _hx_tmp_6, f(e2)}
		}

	case 7:
		{
			var op string
			_ = op
			op = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_op
			var pre bool
			_ = pre
			pre = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_prefix
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = op
			var _hx_tmp_8 bool
			_ = _hx_tmp_8
			_hx_tmp_8 = pre
			edef = Hx_Enum_loreline_lorscript_exprdef_EUnop{_hx_tmp_7, _hx_tmp_8, f(e)}
		}

	case 8:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_e
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_exprdef_ECall).Hx_Field_params
			var edef1 any
			_ = edef1
			edef1 = f(e)
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g2
					if !(_hx_tmp_9 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
						break
					}

					var a any
					_ = a
					a = args.Get(_g2)
					_g2++
					var _hx_tmp_10 HxArray[any]
					_ = _hx_tmp_10
					_hx_tmp_10 = HxMakeArrayView[any](_g1)
					_ = (Hx_Array_Push[any](_hx_tmp_10, f(a)))
				}
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_ECall{edef1, _g1}
		}

	case 9:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_cond
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
			var _hx_tmp_11 any
			_ = _hx_tmp_11
			_hx_tmp_11 = f(c)
			var _hx_tmp_12 any
			_ = _hx_tmp_12
			_hx_tmp_12 = f(e1)
			var _hx_tmp_13 any
			_ = _hx_tmp_13
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e2, nil))) {
				_hx_tmp_13 = f(e2)
			} else {
				_hx_tmp_13 = nil
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EIf{_hx_tmp_11, _hx_tmp_12, _hx_tmp_13}
		}

	case 10:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_cond
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
			var _hx_tmp_14 any
			_ = _hx_tmp_14
			_hx_tmp_14 = f(c)
			edef = Hx_Enum_loreline_lorscript_exprdef_EWhile{_hx_tmp_14, f(e)}
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
			var _hx_tmp_15 string
			_ = _hx_tmp_15
			_hx_tmp_15 = v
			var _hx_tmp_16 any
			_ = _hx_tmp_16
			_hx_tmp_16 = f(it)
			edef = Hx_Enum_loreline_lorscript_exprdef_EFor{_hx_tmp_15, _hx_tmp_16, f(e)}
		}

	case 12, 13:
		{
			edef = func() Hx_Enum_loreline_lorscript_exprdef {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_exprdef
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
			}()
		}

	case 14:
		{
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_args
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
			var name HxNullable[string]
			_ = name
			name = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_name
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			t = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_ret
			var _hx_tmp_17 HxArray[any]
			_ = _hx_tmp_17
			_hx_tmp_17 = args
			edef = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_17, f(e), name, t}
		}

	case 15:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
			var _hx_tmp_18 any
			_ = _hx_tmp_18
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				_hx_tmp_18 = f(e)
			} else {
				_hx_tmp_18 = nil
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EReturn{_hx_tmp_18}
		}

	case 16:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
			var i any
			_ = i
			i = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
			var _hx_tmp_19 any
			_ = _hx_tmp_19
			_hx_tmp_19 = f(e)
			edef = Hx_Enum_loreline_lorscript_exprdef_EArray{_hx_tmp_19, f(i)}
		}

	case 17:
		{
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_EArrayDecl).Hx_Field_e
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					_hx_tmp_20 = _g2
					if !(_hx_tmp_20 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
						break
					}

					var e any
					_ = e
					e = el.Get(_g2)
					_g2++
					var _hx_tmp_21 HxArray[any]
					_ = _hx_tmp_21
					_hx_tmp_21 = HxMakeArrayView[any](_g1)
					_ = (Hx_Array_Push[any](_hx_tmp_21, f(e)))
				}
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{_g1}
		}

	case 18:
		{
			var cl string
			_ = cl
			cl = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_cl
			var el HxArray[any]
			_ = el
			el = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_params
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_22 int32
					_ = _hx_tmp_22
					_hx_tmp_22 = _g2
					if !(_hx_tmp_22 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
						break
					}

					var e any
					_ = e
					e = el.Get(_g2)
					_g2++
					var _hx_tmp_23 HxArray[any]
					_ = _hx_tmp_23
					_hx_tmp_23 = HxMakeArrayView[any](_g1)
					_ = (Hx_Array_Push[any](_hx_tmp_23, f(e)))
				}
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_ENew{cl, _g1}
		}

	case 19:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EThrow).Hx_Field_e
			edef = Hx_Enum_loreline_lorscript_exprdef_EThrow{f(e)}
		}

	case 20:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_e
			var v string
			_ = v
			v = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_v
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			t = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_t
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
			var _hx_tmp_24 any
			_ = _hx_tmp_24
			_hx_tmp_24 = f(e)
			var _hx_tmp_25 string
			_ = _hx_tmp_25
			_hx_tmp_25 = v
			var _hx_tmp_26 Hx_Enum_loreline_lorscript_ctype
			_ = _hx_tmp_26
			_hx_tmp_26 = t
			edef = Hx_Enum_loreline_lorscript_exprdef_ETry{_hx_tmp_24, _hx_tmp_25, _hx_tmp_26, f(c)}
		}

	case 21:
		{
			var fl HxArray[any]
			_ = fl
			fl = _g.(Hx_Enum_loreline_lorscript_exprdef_EObject).Hx_Field_fl
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					_hx_tmp_27 = _g2
					if !(_hx_tmp_27 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
						break
					}

					var fi any
					_ = fi
					fi = fl.Get(_g2)
					_g2++
					var _hx_tmp_28 HxArray[any]
					_ = _hx_tmp_28
					_hx_tmp_28 = HxMakeArrayView[any](_g1)
					var _hx_tmp_29 any
					_ = _hx_tmp_29
					_hx_tmp_29 = Hx_Field_go_haxe_hxdynamic_getField(fi, "name")
					_ = (Hx_Array_Push[any](_hx_tmp_28, any(map[string]any{"name": _hx_tmp_29, "e": f(Hx_Field_go_haxe_hxdynamic_getField(fi, "e"))})))
				}
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_EObject{_g1}
		}

	case 22:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
			var _hx_tmp_30 any
			_ = _hx_tmp_30
			_hx_tmp_30 = f(c)
			var _hx_tmp_31 any
			_ = _hx_tmp_31
			_hx_tmp_31 = f(e1)
			edef = Hx_Enum_loreline_lorscript_exprdef_ETernary{_hx_tmp_30, _hx_tmp_31, f(e2)}
		}

	case 23:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_e
			var cases HxArray[any]
			_ = cases
			cases = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_cases
			var def any
			_ = def
			def = _g.(Hx_Enum_loreline_lorscript_exprdef_ESwitch).Hx_Field_defaultExpr
			var edef1 any
			_ = edef1
			edef1 = f(e)
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxMakeArray[any]().Dyn()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				for {
					var _hx_tmp_32 int32
					_ = _hx_tmp_32
					_hx_tmp_32 = _g2
					if !(_hx_tmp_32 < Hx_Field_go_haxe_hxdynamic_getArrayLength(cases)) {
						break
					}

					var c any
					_ = c
					c = cases.Get(_g2)
					_g2++
					var tmp__g_1 HxArray[any]
					_ = tmp__g_1
					tmp__g_1 = HxMakeArray[any]().Dyn()
					{
						var _g3 int32
						_ = _g3
						_g3 = 0
						var _g4 HxArray[any]
						_ = _g4
						_g4 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "values"))
						for {
							var _hx_tmp_33 int32
							_ = _hx_tmp_33
							_hx_tmp_33 = _g3
							if !(_hx_tmp_33 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g4)) {
								break
							}

							var v any
							_ = v
							v = _g4.Get(_g3)
							_g3++
							var _hx_tmp_34 HxArray[any]
							_ = _hx_tmp_34
							_hx_tmp_34 = HxMakeArrayView[any](tmp__g_1)
							_ = (Hx_Array_Push[any](_hx_tmp_34, f(v)))
						}
					}

					var _hx_tmp_35 HxArray[any]
					_ = _hx_tmp_35
					_hx_tmp_35 = HxMakeArrayView[any](_g1)
					var _hx_tmp_36 any
					_ = _hx_tmp_36
					_hx_tmp_36 = ((any)(tmp__g_1))
					_ = (Hx_Array_Push[any](_hx_tmp_35, any(map[string]any{"values": _hx_tmp_36, "expr": f(Hx_Field_go_haxe_hxdynamic_getField(c, "expr"))})))
				}
			}

			var _hx_tmp_37 any
			_ = _hx_tmp_37
			_hx_tmp_37 = edef1
			var _hx_tmp_38 HxArray[any]
			_ = _hx_tmp_38
			_hx_tmp_38 = _g1
			var _hx_tmp_39 any
			_ = _hx_tmp_39
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(def, nil))) {
				_hx_tmp_39 = nil
			} else {
				_hx_tmp_39 = f(def)
			}

			edef = Hx_Enum_loreline_lorscript_exprdef_ESwitch{_hx_tmp_37, _hx_tmp_38, _hx_tmp_39}
		}

	case 24:
		{
			var c any
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_cond
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
			var _hx_tmp_40 any
			_ = _hx_tmp_40
			_hx_tmp_40 = f(c)
			edef = Hx_Enum_loreline_lorscript_exprdef_EDoWhile{_hx_tmp_40, f(e)}
		}

	case 25:
		{
			var name string
			_ = name
			name = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_name
			var args HxArray[any]
			_ = args
			args = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_args
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
			var edef1 HxArray[any]
			_ = edef1
			if args == nil {
				edef1 = nil
			} else {
				var tmp__g_1 HxArray[any]
				_ = tmp__g_1
				tmp__g_1 = HxMakeArray[any]().Dyn()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_41 int32
						_ = _hx_tmp_41
						_hx_tmp_41 = _g1
						if !(_hx_tmp_41 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
							break
						}

						var a any
						_ = a
						a = args.Get(_g1)
						_g1++
						var _hx_tmp_42 HxArray[any]
						_ = _hx_tmp_42
						_hx_tmp_42 = HxMakeArrayView[any](tmp__g_1)
						_ = (Hx_Array_Push[any](_hx_tmp_42, f(a)))
					}
				}

				edef1 = tmp__g_1
			}

			var _hx_tmp_43 string
			_ = _hx_tmp_43
			_hx_tmp_43 = name
			var _hx_tmp_44 HxArray[any]
			_ = _hx_tmp_44
			_hx_tmp_44 = edef1
			edef = Hx_Enum_loreline_lorscript_exprdef_EMeta{_hx_tmp_43, _hx_tmp_44, f(e)}
		}

	case 26:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			t = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_t
			edef = Hx_Enum_loreline_lorscript_exprdef_ECheckType{f(e), t}
		}

	}

	var _hx_tmp_45 any
	_ = _hx_tmp_45
	_hx_tmp_45 = ((any)(edef))
	var _hx_tmp_46 any
	_ = _hx_tmp_46
	_hx_tmp_46 = Hx_Field_go_haxe_hxdynamic_getField(e, "pmin")
	var _hx_tmp_47 any
	_ = _hx_tmp_47
	_hx_tmp_47 = Hx_Field_go_haxe_hxdynamic_getField(e, "pmax")
	var _hx_tmp_48 any
	_ = _hx_tmp_48
	_hx_tmp_48 = Hx_Field_go_haxe_hxdynamic_getField(e, "origin")
	return any(map[string]any{"e": _hx_tmp_45, "pmin": _hx_tmp_46, "pmax": _hx_tmp_47, "origin": _hx_tmp_48, "line": Hx_Field_go_haxe_hxdynamic_getField(e, "line")})
}

func Hx_Field_loreline_lorscript_tools_expr(e any) Hx_Enum_loreline_lorscript_exprdef {
	return func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
}

func Hx_Field_loreline_lorscript_tools_mk(e Hx_Enum_loreline_lorscript_exprdef, p any) any {
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

func init() {
	hxRegisterArrayConstructor[any]()
}
