package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_printer() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_printer() {
}

type Hx_Obj_VTable_loreline_lorscript_printer interface {
	Hx_Field_exprToString(e any) string
	Hx_Field_typeToString(t Hx_Enum_loreline_lorscript_ctype) string
	Hx_Field_add(s any)
	Hx_Field_type(t Hx_Enum_loreline_lorscript_ctype)
	Hx_Field_addType(t Hx_Enum_loreline_lorscript_ctype)
	Hx_Field_expr(e any)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_printer struct {
	VTable        Hx_Obj_VTable_loreline_lorscript_printer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_buf  *Hx_Obj_stringbuf
	Hx_Field_tabs string
}

func Hx_Obj_loreline_lorscript_printer_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_printer {
	obj := &Hx_Obj_loreline_lorscript_printer{}
	obj.VTable = obj
	obj.Hx_Field_buf = HxDefault[*Hx_Obj_stringbuf]()
	obj.Hx_Field_tabs = HxDefault[string]()
	return obj
}

func Hx_Obj_loreline_lorscript_printer_CreateInstance() *Hx_Obj_loreline_lorscript_printer {
	obj := Hx_Obj_loreline_lorscript_printer_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_New() {
	{
	}

}

var Hx_Obj_loreline_lorscript_printer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Printer",
	HxMakeArray[string]("toString", "errorToString"),
	HxMakeArray[string]("buf", "tabs", "exprToString", "typeToString", "add", "type", "addType", "expr"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_printer_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lorscript_printer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_printer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toString":
			return Hx_Field_loreline_lorscript_printer_toString
		case "errorToString":
			return Hx_Field_loreline_lorscript_printer_errorToString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_printer_RTTI
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_As_Hx_Obj_loreline_lorscript_printer() *Hx_Obj_loreline_lorscript_printer {
	return this
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_exprToString(e any) string {
	this.Hx_Field_buf = Hx_Obj_stringbuf_CreateInstance()
	this.Hx_Field_tabs = ""
	this.VTable.Hx_Field_expr(e)
	return this.Hx_Field_buf.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_typeToString(t Hx_Enum_loreline_lorscript_ctype) string {
	this.Hx_Field_buf = Hx_Obj_stringbuf_CreateInstance()
	this.Hx_Field_tabs = ""
	this.VTable.Hx_Field_type(t)
	return this.Hx_Field_buf.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_add(s any) {
	this.Hx_Field_buf.VTable.Hx_Field_add(s)
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_type(t Hx_Enum_loreline_lorscript_ctype) {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			var path HxArray[string]
			_ = path
			path = t.(Hx_Enum_loreline_lorscript_ctype_CTPath).Hx_Field_path
			var params HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = params
			params = t.(Hx_Enum_loreline_lorscript_ctype_CTPath).Hx_Field_params
			{
				{
					var s string
					_ = s
					s = (Hx_Array_Join[string](path, "."))
					func() func(x string) {
						var _hx_closure_src_201 func(x any)
						_ = _hx_closure_src_201
						_hx_closure_src_201 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_201(((any)(x)))
						}
					}()(s)
				}

				if params != nil {
					func() func(x string) {
						var _hx_closure_src_202 func(x any)
						_ = _hx_closure_src_202
						_hx_closure_src_202 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_202(((any)(x)))
						}
					}()("<")
					var first bool
					_ = first
					first = true
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_0 int32
							_ = _hx_tmp_0
							_hx_tmp_0 = _g
							if !(_hx_tmp_0 < params.Len()) {
								break
							}

							var p Hx_Enum_loreline_lorscript_ctype
							_ = p
							p = params.Get(_g)
							_g++
							if first {
								first = false
							} else {
								func() func(x string) {
									var _hx_closure_src_203 func(x any)
									_ = _hx_closure_src_203
									_hx_closure_src_203 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_203(((any)(x)))
									}
								}()(", ")
							}

							this.VTable.Hx_Field_type(p)
						}
					}

					func() func(x string) {
						var _hx_closure_src_204 func(x any)
						_ = _hx_closure_src_204
						_hx_closure_src_204 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_204(((any)(x)))
						}
					}()(">")
				}
			}
		}

	case 1:
		{
			var _gargs HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = _gargs
			_gargs = t.(Hx_Enum_loreline_lorscript_ctype_CTFun).Hx_Field_args
			var _gret Hx_Enum_loreline_lorscript_ctype
			_ = _gret
			_gret = t.(Hx_Enum_loreline_lorscript_ctype_CTFun).Hx_Field_ret
			{
				var args HxArray[Hx_Enum_loreline_lorscript_ctype]
				_ = args
				args = _gargs
				var ret Hx_Enum_loreline_lorscript_ctype
				_ = ret
				ret = _gret
				var _hx_tmp_1 HxArray[Hx_Enum_loreline_lorscript_ctype]
				_ = _hx_tmp_1
				_hx_tmp_1 = args
				if Hx_Field_lambda_exists(_hx_tmp_1, func() func(a any) bool {
					var _hx_closure_src_219 func(a Hx_Enum_loreline_lorscript_ctype) bool
					_ = _hx_closure_src_219
					_hx_closure_src_219 = func(a Hx_Enum_loreline_lorscript_ctype) bool {
						return (a.Hx_Field_enumIndex() == 5)
					}

					return func(a any) bool {
						return _hx_closure_src_219(func() Hx_Enum_loreline_lorscript_ctype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(a)
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ctype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
						}())
					}
				}()) {
					func() func(x string) {
						var _hx_closure_src_205 func(x any)
						_ = _hx_closure_src_205
						_hx_closure_src_205 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_205(((any)(x)))
						}
					}()("(")
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_2 int32
							_ = _hx_tmp_2
							_hx_tmp_2 = _g
							if !(_hx_tmp_2 < args.Len()) {
								break
							}

							var a Hx_Enum_loreline_lorscript_ctype
							_ = a
							a = args.Get(_g)
							_g++
							if a.Hx_Field_enumIndex() == 5 {
								this.VTable.Hx_Field_type(a)
							} else {
								this.VTable.Hx_Field_type(Hx_Enum_loreline_lorscript_ctype_CTNamed{"_", a})
							}
						}
					}

					func() func(x string) {
						var _hx_closure_src_206 func(x any)
						_ = _hx_closure_src_206
						_hx_closure_src_206 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_206(((any)(x)))
						}
					}()(")->")
					this.VTable.Hx_Field_type(ret)
				} else {
					var tmp_args_1 HxArray[Hx_Enum_loreline_lorscript_ctype]
					_ = tmp_args_1
					tmp_args_1 = _gargs
					var tmp_ret_1 Hx_Enum_loreline_lorscript_ctype
					_ = tmp_ret_1
					tmp_ret_1 = _gret
					{
						if tmp_args_1.Len() == 0 {
							func() func(x string) {
								var _hx_closure_src_207 func(x any)
								_ = _hx_closure_src_207
								_hx_closure_src_207 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_207(((any)(x)))
								}
							}()("Void -> ")
						} else {
							var _g int32
							_ = _g
							_g = 0
							for {
								var _hx_tmp_3 int32
								_ = _hx_tmp_3
								_hx_tmp_3 = _g
								if !(_hx_tmp_3 < tmp_args_1.Len()) {
									break
								}

								var a Hx_Enum_loreline_lorscript_ctype
								_ = a
								a = tmp_args_1.Get(_g)
								_g++
								this.VTable.Hx_Field_type(a)
								func() func(x string) {
									var _hx_closure_src_208 func(x any)
									_ = _hx_closure_src_208
									_hx_closure_src_208 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_208(((any)(x)))
									}
								}()(" -> ")
							}
						}

						this.VTable.Hx_Field_type(tmp_ret_1)
					}
				}
			}
		}

	case 2:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_loreline_lorscript_ctype_CTAnon).Hx_Field_fields
			{
				func() func(x string) {
					var _hx_closure_src_209 func(x any)
					_ = _hx_closure_src_209
					_hx_closure_src_209 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_209(((any)(x)))
					}
				}()("{")
				var first bool
				_ = first
				first = true
				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g
						if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fields)) {
							break
						}

						var f any
						_ = f
						f = fields.Get(_g)
						_g++
						if first {
							first = false
							func() func(x string) {
								var _hx_closure_src_210 func(x any)
								_ = _hx_closure_src_210
								_hx_closure_src_210 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_210(((any)(x)))
								}
							}()(" ")
						} else {
							func() func(x string) {
								var _hx_closure_src_211 func(x any)
								_ = _hx_closure_src_211
								_hx_closure_src_211 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_211(((any)(x)))
								}
							}()(", ")
						}

						func() func(x string) {
							var _hx_closure_src_212 func(x any)
							_ = _hx_closure_src_212
							_hx_closure_src_212 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_212(((any)(x)))
							}
						}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), ((any)(" : ")))))
						this.VTable.Hx_Field_type(func() Hx_Enum_loreline_lorscript_ctype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_loreline_lorscript_ctype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
						}())
					}
				}

				var _hx_tmp_5 string
				_ = _hx_tmp_5
				if first {
					_hx_tmp_5 = "}"
				} else {
					_hx_tmp_5 = " }"
				}

				func() func(x string) {
					var _hx_closure_src_213 func(x any)
					_ = _hx_closure_src_213
					_hx_closure_src_213 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_213(((any)(x)))
					}
				}()(_hx_tmp_5)
			}
		}

	case 3:
		{
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTParent).Hx_Field_t
			{
				func() func(x string) {
					var _hx_closure_src_214 func(x any)
					_ = _hx_closure_src_214
					_hx_closure_src_214 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_214(((any)(x)))
					}
				}()("(")
				this.VTable.Hx_Field_type(t1)
				func() func(x string) {
					var _hx_closure_src_215 func(x any)
					_ = _hx_closure_src_215
					_hx_closure_src_215 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_215(((any)(x)))
					}
				}()(")")
			}
		}

	case 4:
		{
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTOpt).Hx_Field_t
			{
				func() func(x string) {
					var _hx_closure_src_216 func(x any)
					_ = _hx_closure_src_216
					_hx_closure_src_216 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_216(((any)(x)))
					}
				}()("?")
				this.VTable.Hx_Field_type(t1)
			}
		}

	case 5:
		{
			var name string
			_ = name
			name = t.(Hx_Enum_loreline_lorscript_ctype_CTNamed).Hx_Field_n
			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = t.(Hx_Enum_loreline_lorscript_ctype_CTNamed).Hx_Field_t
			{
				func() func(x string) {
					var _hx_closure_src_217 func(x any)
					_ = _hx_closure_src_217
					_hx_closure_src_217 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_217(((any)(x)))
					}
				}()(name)
				func() func(x string) {
					var _hx_closure_src_218 func(x any)
					_ = _hx_closure_src_218
					_hx_closure_src_218 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_218(((any)(x)))
					}
				}()(":")
				this.VTable.Hx_Field_type(t1)
			}
		}

	}
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_addType(t Hx_Enum_loreline_lorscript_ctype) {
	if t != nil {
		func() func(x string) {
			var _hx_closure_src_220 func(x any)
			_ = _hx_closure_src_220
			_hx_closure_src_220 = this.Hx_Field_buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_220(((any)(x)))
			}
		}()(" : ")
		this.VTable.Hx_Field_type(t)
	}
}

func (this *Hx_Obj_loreline_lorscript_printer) Hx_Field_expr(e any) {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		func() func(x string) {
			var _hx_closure_src_221 func(x any)
			_ = _hx_closure_src_221
			_hx_closure_src_221 = this.Hx_Field_buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_221(((any)(x)))
			}
		}()("??NULL??")
		return
	}

	{
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
				var c Hx_Enum_loreline_lorscript_const
				_ = c
				c = _g.(Hx_Enum_loreline_lorscript_exprdef_EConst).Hx_Field_c
				switch c.Hx_Field_enumIndex() {
				case 0:
					{
						var i int32
						_ = i
						i = c.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
						func() func(x int32) {
							var _hx_closure_src_222 func(x any)
							_ = _hx_closure_src_222
							_hx_closure_src_222 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x int32) {
								_hx_closure_src_222(((any)(x)))
							}
						}()(i)
					}

				case 1:
					{
						var f float64
						_ = f
						f = c.(Hx_Enum_loreline_lorscript_const_CFloat).Hx_Field_f
						func() func(x float64) {
							var _hx_closure_src_223 func(x any)
							_ = _hx_closure_src_223
							_hx_closure_src_223 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x float64) {
								_hx_closure_src_223(((any)(x)))
							}
						}()(f)
					}

				case 2:
					{
						var s string
						_ = s
						s = c.(Hx_Enum_loreline_lorscript_const_CString).Hx_Field_s
						{
							func() func(x string) {
								var _hx_closure_src_224 func(x any)
								_ = _hx_closure_src_224
								_hx_closure_src_224 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_224(((any)(x)))
								}
							}()("\"")
							{
								var s1 string
								_ = s1
								s1 = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\"")), "\\\"")), "\n")), "\\n")), "\r")), "\\r")), "\t")), "\\t"))
								func() func(x string) {
									var _hx_closure_src_225 func(x any)
									_ = _hx_closure_src_225
									_hx_closure_src_225 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_225(((any)(x)))
									}
								}()(s1)
							}

							func() func(x string) {
								var _hx_closure_src_226 func(x any)
								_ = _hx_closure_src_226
								_hx_closure_src_226 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_226(((any)(x)))
								}
							}()("\"")
						}
					}

				}
			}

		case 1:
			{
				var v string
				_ = v
				v = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
				func() func(x string) {
					var _hx_closure_src_227 func(x any)
					_ = _hx_closure_src_227
					_hx_closure_src_227 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_227(((any)(x)))
					}
				}()(v)
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
				{
					func() func(x string) {
						var _hx_closure_src_228 func(x any)
						_ = _hx_closure_src_228
						_hx_closure_src_228 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_228(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat("var ", n))
					this.VTable.Hx_Field_addType(t)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
						func() func(x string) {
							var _hx_closure_src_229 func(x any)
							_ = _hx_closure_src_229
							_hx_closure_src_229 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_229(((any)(x)))
							}
						}()(" = ")
						this.VTable.Hx_Field_expr(e)
					}
				}
			}

		case 3:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
				{
					func() func(x string) {
						var _hx_closure_src_230 func(x any)
						_ = _hx_closure_src_230
						_hx_closure_src_230 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_230(((any)(x)))
						}
					}()("(")
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_231 func(x any)
						_ = _hx_closure_src_231
						_hx_closure_src_231 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_231(((any)(x)))
						}
					}()(")")
				}
			}

		case 4:
			{
				var el HxArray[any]
				_ = el
				el = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
				if Hx_Field_go_haxe_hxdynamic_getArrayLength(el) == 0 {
					func() func(x string) {
						var _hx_closure_src_232 func(x any)
						_ = _hx_closure_src_232
						_hx_closure_src_232 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_232(((any)(x)))
						}
					}()("{}")
				} else {
					this.Hx_Field_tabs = Hx_Field_go_haxe_hxstring_concat(this.Hx_Field_tabs, "\t")
					func() func(x string) {
						var _hx_closure_src_233 func(x any)
						_ = _hx_closure_src_233
						_hx_closure_src_233 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_233(((any)(x)))
						}
					}()("{\n")
					{
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						for {
							var _hx_tmp_0 int32
							_ = _hx_tmp_0
							_hx_tmp_0 = tmp__g_1
							if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
								break
							}

							var e any
							_ = e
							e = el.Get(tmp__g_1)
							tmp__g_1++
							func() func(x string) {
								var _hx_closure_src_234 func(x any)
								_ = _hx_closure_src_234
								_hx_closure_src_234 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_234(((any)(x)))
								}
							}()(this.Hx_Field_tabs)
							this.VTable.Hx_Field_expr(e)
							func() func(x string) {
								var _hx_closure_src_235 func(x any)
								_ = _hx_closure_src_235
								_hx_closure_src_235 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_235(((any)(x)))
								}
							}()(";\n")
						}
					}

					var _hx_tmp_1 string
					_ = _hx_tmp_1
					_hx_tmp_1 = this.Hx_Field_tabs
					this.Hx_Field_tabs = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 1, HxNullable[int32]{}))
					func() func(x string) {
						var _hx_closure_src_236 func(x any)
						_ = _hx_closure_src_236
						_hx_closure_src_236 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_236(((any)(x)))
						}
					}()("}")
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
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_237 func(x any)
						_ = _hx_closure_src_237
						_hx_closure_src_237 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_237(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(".", f))
				}
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
				{
					this.VTable.Hx_Field_expr(e1)
					func() func(x string) {
						var _hx_closure_src_238 func(x any)
						_ = _hx_closure_src_238
						_hx_closure_src_238 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_238(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(" ", op), " "))
					this.VTable.Hx_Field_expr(e2)
				}
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
				if pre {
					func() func(x string) {
						var _hx_closure_src_239 func(x any)
						_ = _hx_closure_src_239
						_hx_closure_src_239 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_239(((any)(x)))
						}
					}()(op)
					this.VTable.Hx_Field_expr(e)
				} else {
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_240 func(x any)
						_ = _hx_closure_src_240
						_hx_closure_src_240 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_240(((any)(x)))
						}
					}()(op)
				}
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
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
						this.VTable.Hx_Field_expr(e)
					} else {
						switch Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() {
						case 0:
							{
								this.VTable.Hx_Field_expr(e)
							}

						case 1:
							{
								this.VTable.Hx_Field_expr(e)
							}

						case 5:
							{
								this.VTable.Hx_Field_expr(e)
							}

						default:
							{
								func() func(x string) {
									var _hx_closure_src_241 func(x any)
									_ = _hx_closure_src_241
									_hx_closure_src_241 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_241(((any)(x)))
									}
								}()("(")
								this.VTable.Hx_Field_expr(e)
								func() func(x string) {
									var _hx_closure_src_242 func(x any)
									_ = _hx_closure_src_242
									_hx_closure_src_242 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_242(((any)(x)))
									}
								}()(")")
							}
						}
					}

					func() func(x string) {
						var _hx_closure_src_243 func(x any)
						_ = _hx_closure_src_243
						_hx_closure_src_243 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_243(((any)(x)))
						}
					}()("(")
					var first bool
					_ = first
					first = true
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_2 int32
							_ = _hx_tmp_2
							_hx_tmp_2 = _g1
							if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
								break
							}

							var a any
							_ = a
							a = args.Get(_g1)
							_g1++
							if first {
								first = false
							} else {
								func() func(x string) {
									var _hx_closure_src_244 func(x any)
									_ = _hx_closure_src_244
									_hx_closure_src_244 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_244(((any)(x)))
									}
								}()(", ")
							}

							this.VTable.Hx_Field_expr(a)
						}
					}

					func() func(x string) {
						var _hx_closure_src_245 func(x any)
						_ = _hx_closure_src_245
						_hx_closure_src_245 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_245(((any)(x)))
						}
					}()(")")
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
					func() func(x string) {
						var _hx_closure_src_246 func(x any)
						_ = _hx_closure_src_246
						_hx_closure_src_246 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_246(((any)(x)))
						}
					}()("if( ")
					this.VTable.Hx_Field_expr(cond)
					func() func(x string) {
						var _hx_closure_src_247 func(x any)
						_ = _hx_closure_src_247
						_hx_closure_src_247 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_247(((any)(x)))
						}
					}()(" ) ")
					this.VTable.Hx_Field_expr(e1)
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e2, nil))) {
						func() func(x string) {
							var _hx_closure_src_248 func(x any)
							_ = _hx_closure_src_248
							_hx_closure_src_248 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_248(((any)(x)))
							}
						}()(" else ")
						this.VTable.Hx_Field_expr(e2)
					}
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
				{
					func() func(x string) {
						var _hx_closure_src_249 func(x any)
						_ = _hx_closure_src_249
						_hx_closure_src_249 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_249(((any)(x)))
						}
					}()("while( ")
					this.VTable.Hx_Field_expr(cond)
					func() func(x string) {
						var _hx_closure_src_250 func(x any)
						_ = _hx_closure_src_250
						_hx_closure_src_250 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_250(((any)(x)))
						}
					}()(" ) ")
					this.VTable.Hx_Field_expr(e)
				}
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
				{
					func() func(x string) {
						var _hx_closure_src_251 func(x any)
						_ = _hx_closure_src_251
						_hx_closure_src_251 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_251(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("for( ", v), " in "))
					this.VTable.Hx_Field_expr(it)
					func() func(x string) {
						var _hx_closure_src_252 func(x any)
						_ = _hx_closure_src_252
						_hx_closure_src_252 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_252(((any)(x)))
						}
					}()(" ) ")
					this.VTable.Hx_Field_expr(e)
				}
			}

		case 12:
			{
				func() func(x string) {
					var _hx_closure_src_253 func(x any)
					_ = _hx_closure_src_253
					_hx_closure_src_253 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_253(((any)(x)))
					}
				}()("break")
			}

		case 13:
			{
				func() func(x string) {
					var _hx_closure_src_254 func(x any)
					_ = _hx_closure_src_254
					_hx_closure_src_254 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_254(((any)(x)))
					}
				}()("continue")
			}

		case 14:
			{
				var params HxArray[any]
				_ = params
				params = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_args
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
				var name HxNullable[string]
				_ = name
				name = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_name
				var ret Hx_Enum_loreline_lorscript_ctype
				_ = ret
				ret = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_ret
				{
					func() func(x string) {
						var _hx_closure_src_255 func(x any)
						_ = _hx_closure_src_255
						_hx_closure_src_255 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_255(((any)(x)))
						}
					}()("function")
					if (name).Valid != false {
						func() func(x string) {
							var _hx_closure_src_256 func(x any)
							_ = _hx_closure_src_256
							_hx_closure_src_256 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_256(((any)(x)))
							}
						}()(Hx_Field_go_haxe_hxstring_concat(" ", (name).Value))
					}

					func() func(x string) {
						var _hx_closure_src_257 func(x any)
						_ = _hx_closure_src_257
						_hx_closure_src_257 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_257(((any)(x)))
						}
					}()("(")
					var first bool
					_ = first
					first = true
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_3 int32
							_ = _hx_tmp_3
							_hx_tmp_3 = _g1
							if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(params)) {
								break
							}

							var a any
							_ = a
							a = params.Get(_g1)
							_g1++
							if first {
								first = false
							} else {
								func() func(x string) {
									var _hx_closure_src_258 func(x any)
									_ = _hx_closure_src_258
									_hx_closure_src_258 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_258(((any)(x)))
									}
								}()(", ")
							}

							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(a, "opt")))) {
								func() func(x string) {
									var _hx_closure_src_259 func(x any)
									_ = _hx_closure_src_259
									_hx_closure_src_259 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_259(((any)(x)))
									}
								}()("?")
							}

							func() func(x string) {
								var _hx_closure_src_260 func(x any)
								_ = _hx_closure_src_260
								_hx_closure_src_260 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_260(((any)(x)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(a, "name")))
							this.VTable.Hx_Field_addType(func() Hx_Enum_loreline_lorscript_ctype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ctype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
							}())
						}
					}

					func() func(x string) {
						var _hx_closure_src_261 func(x any)
						_ = _hx_closure_src_261
						_hx_closure_src_261 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_261(((any)(x)))
						}
					}()(")")
					this.VTable.Hx_Field_addType(ret)
					func() func(x string) {
						var _hx_closure_src_262 func(x any)
						_ = _hx_closure_src_262
						_hx_closure_src_262 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_262(((any)(x)))
						}
					}()(" ")
					this.VTable.Hx_Field_expr(e)
				}
			}

		case 15:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
				{
					func() func(x string) {
						var _hx_closure_src_263 func(x any)
						_ = _hx_closure_src_263
						_hx_closure_src_263 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_263(((any)(x)))
						}
					}()("return")
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
						func() func(x string) {
							var _hx_closure_src_264 func(x any)
							_ = _hx_closure_src_264
							_hx_closure_src_264 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_264(((any)(x)))
							}
						}()(" ")
						this.VTable.Hx_Field_expr(e)
					}
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
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_265 func(x any)
						_ = _hx_closure_src_265
						_hx_closure_src_265 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_265(((any)(x)))
						}
					}()("[")
					this.VTable.Hx_Field_expr(index)
					func() func(x string) {
						var _hx_closure_src_266 func(x any)
						_ = _hx_closure_src_266
						_hx_closure_src_266 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_266(((any)(x)))
						}
					}()("]")
				}
			}

		case 17:
			{
				var el HxArray[any]
				_ = el
				el = _g.(Hx_Enum_loreline_lorscript_exprdef_EArrayDecl).Hx_Field_e
				{
					func() func(x string) {
						var _hx_closure_src_267 func(x any)
						_ = _hx_closure_src_267
						_hx_closure_src_267 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_267(((any)(x)))
						}
					}()("[")
					var first bool
					_ = first
					first = true
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = _g1
							if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(el)) {
								break
							}

							var e any
							_ = e
							e = el.Get(_g1)
							_g1++
							if first {
								first = false
							} else {
								func() func(x string) {
									var _hx_closure_src_268 func(x any)
									_ = _hx_closure_src_268
									_hx_closure_src_268 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_268(((any)(x)))
									}
								}()(", ")
							}

							this.VTable.Hx_Field_expr(e)
						}
					}

					func() func(x string) {
						var _hx_closure_src_269 func(x any)
						_ = _hx_closure_src_269
						_hx_closure_src_269 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_269(((any)(x)))
						}
					}()("]")
				}
			}

		case 18:
			{
				var cl string
				_ = cl
				cl = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_cl
				var args HxArray[any]
				_ = args
				args = _g.(Hx_Enum_loreline_lorscript_exprdef_ENew).Hx_Field_params
				{
					func() func(x string) {
						var _hx_closure_src_270 func(x any)
						_ = _hx_closure_src_270
						_hx_closure_src_270 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_270(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("new ", cl), "("))
					var first bool
					_ = first
					first = true
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_5 int32
							_ = _hx_tmp_5
							_hx_tmp_5 = _g1
							if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
								break
							}

							var e any
							_ = e
							e = args.Get(_g1)
							_g1++
							if first {
								first = false
							} else {
								func() func(x string) {
									var _hx_closure_src_271 func(x any)
									_ = _hx_closure_src_271
									_hx_closure_src_271 = this.Hx_Field_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_271(((any)(x)))
									}
								}()(", ")
							}

							this.VTable.Hx_Field_expr(e)
						}
					}

					func() func(x string) {
						var _hx_closure_src_272 func(x any)
						_ = _hx_closure_src_272
						_hx_closure_src_272 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_272(((any)(x)))
						}
					}()(")")
				}
			}

		case 19:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EThrow).Hx_Field_e
				{
					func() func(x string) {
						var _hx_closure_src_273 func(x any)
						_ = _hx_closure_src_273
						_hx_closure_src_273 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_273(((any)(x)))
						}
					}()("throw ")
					this.VTable.Hx_Field_expr(e)
				}
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
				var ecatch any
				_ = ecatch
				ecatch = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
				{
					func() func(x string) {
						var _hx_closure_src_274 func(x any)
						_ = _hx_closure_src_274
						_hx_closure_src_274 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_274(((any)(x)))
						}
					}()("try ")
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_275 func(x any)
						_ = _hx_closure_src_275
						_hx_closure_src_275 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_275(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(" catch( ", v))
					this.VTable.Hx_Field_addType(t)
					func() func(x string) {
						var _hx_closure_src_276 func(x any)
						_ = _hx_closure_src_276
						_hx_closure_src_276 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_276(((any)(x)))
						}
					}()(") ")
					this.VTable.Hx_Field_expr(ecatch)
				}
			}

		case 21:
			{
				var fl HxArray[any]
				_ = fl
				fl = _g.(Hx_Enum_loreline_lorscript_exprdef_EObject).Hx_Field_fl
				if Hx_Field_go_haxe_hxdynamic_getArrayLength(fl) == 0 {
					func() func(x string) {
						var _hx_closure_src_277 func(x any)
						_ = _hx_closure_src_277
						_hx_closure_src_277 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_277(((any)(x)))
						}
					}()("{}")
				} else {
					this.Hx_Field_tabs = Hx_Field_go_haxe_hxstring_concat(this.Hx_Field_tabs, "\t")
					func() func(x string) {
						var _hx_closure_src_278 func(x any)
						_ = _hx_closure_src_278
						_hx_closure_src_278 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_278(((any)(x)))
						}
					}()("{\n")
					{
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						for {
							var _hx_tmp_6 int32
							_ = _hx_tmp_6
							_hx_tmp_6 = tmp__g_1
							if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fl)) {
								break
							}

							var f any
							_ = f
							f = fl.Get(tmp__g_1)
							tmp__g_1++
							func() func(x string) {
								var _hx_closure_src_279 func(x any)
								_ = _hx_closure_src_279
								_hx_closure_src_279 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_279(((any)(x)))
								}
							}()(this.Hx_Field_tabs)
							func() func(x string) {
								var _hx_closure_src_280 func(x any)
								_ = _hx_closure_src_280
								_hx_closure_src_280 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_280(((any)(x)))
								}
							}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(f, "name"), ((any)(" : ")))))
							this.VTable.Hx_Field_expr(Hx_Field_go_haxe_hxdynamic_getField(f, "e"))
							func() func(x string) {
								var _hx_closure_src_281 func(x any)
								_ = _hx_closure_src_281
								_hx_closure_src_281 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_281(((any)(x)))
								}
							}()(",\n")
						}
					}

					var _hx_tmp_7 string
					_ = _hx_tmp_7
					_hx_tmp_7 = this.Hx_Field_tabs
					this.Hx_Field_tabs = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, 1, HxNullable[int32]{}))
					func() func(x string) {
						var _hx_closure_src_282 func(x any)
						_ = _hx_closure_src_282
						_hx_closure_src_282 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_282(((any)(x)))
						}
					}()("}")
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
					this.VTable.Hx_Field_expr(c)
					func() func(x string) {
						var _hx_closure_src_283 func(x any)
						_ = _hx_closure_src_283
						_hx_closure_src_283 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_283(((any)(x)))
						}
					}()(" ? ")
					this.VTable.Hx_Field_expr(e1)
					func() func(x string) {
						var _hx_closure_src_284 func(x any)
						_ = _hx_closure_src_284
						_hx_closure_src_284 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_284(((any)(x)))
						}
					}()(" : ")
					this.VTable.Hx_Field_expr(e2)
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
					func() func(x string) {
						var _hx_closure_src_285 func(x any)
						_ = _hx_closure_src_285
						_hx_closure_src_285 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_285(((any)(x)))
						}
					}()("switch( ")
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_286 func(x any)
						_ = _hx_closure_src_286
						_hx_closure_src_286 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_286(((any)(x)))
						}
					}()(") {")
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_8 int32
							_ = _hx_tmp_8
							_hx_tmp_8 = _g1
							if !(_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_getArrayLength(cases)) {
								break
							}

							var c any
							_ = c
							c = cases.Get(_g1)
							_g1++
							func() func(x string) {
								var _hx_closure_src_287 func(x any)
								_ = _hx_closure_src_287
								_hx_closure_src_287 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_287(((any)(x)))
								}
							}()("case ")
							var first bool
							_ = first
							first = true
							{
								var tmp__g_1 int32
								_ = tmp__g_1
								tmp__g_1 = 0
								var _g2 HxArray[any]
								_ = _g2
								_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "values"))
								for {
									var _hx_tmp_9 int32
									_ = _hx_tmp_9
									_hx_tmp_9 = tmp__g_1
									if !(_hx_tmp_9 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
										break
									}

									var v any
									_ = v
									v = _g2.Get(tmp__g_1)
									tmp__g_1++
									if first {
										first = false
									} else {
										func() func(x string) {
											var _hx_closure_src_288 func(x any)
											_ = _hx_closure_src_288
											_hx_closure_src_288 = this.Hx_Field_buf.VTable.Hx_Field_add
											return func(x string) {
												_hx_closure_src_288(((any)(x)))
											}
										}()(", ")
									}

									this.VTable.Hx_Field_expr(v)
								}
							}

							func() func(x string) {
								var _hx_closure_src_289 func(x any)
								_ = _hx_closure_src_289
								_hx_closure_src_289 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_289(((any)(x)))
								}
							}()(": ")
							this.VTable.Hx_Field_expr(Hx_Field_go_haxe_hxdynamic_getField(c, "expr"))
							func() func(x string) {
								var _hx_closure_src_290 func(x any)
								_ = _hx_closure_src_290
								_hx_closure_src_290 = this.Hx_Field_buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_290(((any)(x)))
								}
							}()(";\n")
						}
					}

					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(def, nil))) {
						func() func(x string) {
							var _hx_closure_src_291 func(x any)
							_ = _hx_closure_src_291
							_hx_closure_src_291 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_291(((any)(x)))
							}
						}()("default: ")
						this.VTable.Hx_Field_expr(def)
						func() func(x string) {
							var _hx_closure_src_292 func(x any)
							_ = _hx_closure_src_292
							_hx_closure_src_292 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_292(((any)(x)))
							}
						}()(";\n")
					}

					func() func(x string) {
						var _hx_closure_src_293 func(x any)
						_ = _hx_closure_src_293
						_hx_closure_src_293 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_293(((any)(x)))
						}
					}()("}")
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
				{
					func() func(x string) {
						var _hx_closure_src_294 func(x any)
						_ = _hx_closure_src_294
						_hx_closure_src_294 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_294(((any)(x)))
						}
					}()("do ")
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_295 func(x any)
						_ = _hx_closure_src_295
						_hx_closure_src_295 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_295(((any)(x)))
						}
					}()(" while ( ")
					this.VTable.Hx_Field_expr(cond)
					func() func(x string) {
						var _hx_closure_src_296 func(x any)
						_ = _hx_closure_src_296
						_hx_closure_src_296 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_296(((any)(x)))
						}
					}()(" )")
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
					func() func(x string) {
						var _hx_closure_src_297 func(x any)
						_ = _hx_closure_src_297
						_hx_closure_src_297 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_297(((any)(x)))
						}
					}()("@")
					func() func(x string) {
						var _hx_closure_src_298 func(x any)
						_ = _hx_closure_src_298
						_hx_closure_src_298 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_298(((any)(x)))
						}
					}()(name)
					var _hx_tmp_10 bool
					_ = _hx_tmp_10
					if args != nil {
						_hx_tmp_10 = (Hx_Field_go_haxe_hxdynamic_getArrayLength(args) > 0)
					} else {
						_hx_tmp_10 = false
					}

					if _hx_tmp_10 {
						func() func(x string) {
							var _hx_closure_src_299 func(x any)
							_ = _hx_closure_src_299
							_hx_closure_src_299 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_299(((any)(x)))
							}
						}()("(")
						var first bool
						_ = first
						first = true
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							for {
								var _hx_tmp_11 int32
								_ = _hx_tmp_11
								_hx_tmp_11 = tmp__g_1
								if !(_hx_tmp_11 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
									break
								}

								var a any
								_ = a
								a = args.Get(tmp__g_1)
								tmp__g_1++
								if first {
									first = false
								} else {
									func() func(x string) {
										var _hx_closure_src_300 func(x any)
										_ = _hx_closure_src_300
										_hx_closure_src_300 = this.Hx_Field_buf.VTable.Hx_Field_add
										return func(x string) {
											_hx_closure_src_300(((any)(x)))
										}
									}()(", ")
								}

								this.VTable.Hx_Field_expr(e)
							}
						}

						func() func(x string) {
							var _hx_closure_src_301 func(x any)
							_ = _hx_closure_src_301
							_hx_closure_src_301 = this.Hx_Field_buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_301(((any)(x)))
							}
						}()(")")
					}

					func() func(x string) {
						var _hx_closure_src_302 func(x any)
						_ = _hx_closure_src_302
						_hx_closure_src_302 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_302(((any)(x)))
						}
					}()(" ")
					this.VTable.Hx_Field_expr(e)
				}
			}

		case 26:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e
				var t Hx_Enum_loreline_lorscript_ctype
				_ = t
				t = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_t
				{
					func() func(x string) {
						var _hx_closure_src_303 func(x any)
						_ = _hx_closure_src_303
						_hx_closure_src_303 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_303(((any)(x)))
						}
					}()("(")
					this.VTable.Hx_Field_expr(e)
					func() func(x string) {
						var _hx_closure_src_304 func(x any)
						_ = _hx_closure_src_304
						_hx_closure_src_304 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_304(((any)(x)))
						}
					}()(" : ")
					this.VTable.Hx_Field_addType(t)
					func() func(x string) {
						var _hx_closure_src_305 func(x any)
						_ = _hx_closure_src_305
						_hx_closure_src_305 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_305(((any)(x)))
						}
					}()(")")
				}
			}

		}
	}
}

func Hx_Field_loreline_lorscript_printer_toString(e any) string {
	return Hx_Obj_loreline_lorscript_printer_CreateInstance().VTable.Hx_Field_exprToString(e)
}

func Hx_Field_loreline_lorscript_printer_errorToString(e *Hx_Obj_loreline_lorscript_error) string {
	var message string
	_ = message
	var _g Hx_Enum_loreline_lorscript_errordef
	_ = _g
	_g = e.Hx_Field_e
	switch _g.Hx_Field_enumIndex() {
	case 0:
		{
			var c int32
			_ = c
			c = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidChar).Hx_Field_c
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			if c == 0 {
				_hx_tmp_0 = "EOF"
			} else {
				_hx_tmp_0 = Hx_Field_go_haxe_hxstring_fromCharCode(c)
			}

			message = Hx_Field_go_haxe_hxstring_concat((Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Invalid character: '", (_hx_tmp_0)), "' (") + Hx_Field_go_haxe_hxdynamic_toString(c)), ")")
		}

	case 1:
		{
			var s string
			_ = s
			s = _g.(Hx_Enum_loreline_lorscript_errordef_EUnexpected).Hx_Field_s
			message = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unexpected token: \"", s), "\"")
		}

	case 2:
		{
			message = "Unterminated string"
		}

	case 3:
		{
			message = "Unterminated comment"
		}

	case 4:
		{
			var str string
			_ = str
			str = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor).Hx_Field_msg
			message = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Invalid preprocessor (", str), ")")
		}

	case 5:
		{
			var v string
			_ = v
			v = _g.(Hx_Enum_loreline_lorscript_errordef_EUnknownVariable).Hx_Field_v
			message = Hx_Field_go_haxe_hxstring_concat("Unknown variable: ", v)
		}

	case 6:
		{
			var v string
			_ = v
			v = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidIterator).Hx_Field_v
			message = Hx_Field_go_haxe_hxstring_concat("Invalid iterator: ", v)
		}

	case 7:
		{
			var op string
			_ = op
			op = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidOp).Hx_Field_op
			message = Hx_Field_go_haxe_hxstring_concat("Invalid operator: ", op)
		}

	case 8:
		{
			var f string
			_ = f
			f = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidAccess).Hx_Field_f
			message = Hx_Field_go_haxe_hxstring_concat("Invalid access to field ", f)
		}

	case 9:
		{
			var msg string
			_ = msg
			msg = _g.(Hx_Enum_loreline_lorscript_errordef_ECustom).Hx_Field_msg
			message = msg
		}

	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxstring_concat(e.Hx_Field_origin, ":")
	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat((_hx_tmp_1+Hx_Field_go_haxe_hxdynamic_toString(e.Hx_Field_line)), ": "), message)
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_ctype]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
