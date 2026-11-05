package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_format_jsonprinter() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_format_jsonprinter() {
}

type Hx_Obj_VTable_haxe_format_jsonprinter interface {
	Hx_Field_objString(v any)
	Hx_Field_beginStack(v any)
	Hx_Field_endStack(v any)
	Hx_Field_ipad()
	Hx_Field_newl()
	Hx_Field_write(k any, v any)
	Hx_Field_classString(v any)
	Hx_Field_fieldsString(v any, fields HxArray[string])
	Hx_Field_quote(s string)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_format_jsonprinter struct {
	VTable            Hx_Obj_VTable_haxe_format_jsonprinter `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_buf      *Hx_Obj_stringbuf
	Hx_Field_replacer func(key any, value any) any
	Hx_Field_indent   string
	Hx_Field_pretty   bool
	Hx_Field_nind     int32
	Hx_Field_stack    Hx_Typedef_map
	Hx_Field_depth    int32
}

func Hx_Obj_haxe_format_jsonprinter_CreateEmptyInstance() *Hx_Obj_haxe_format_jsonprinter {
	obj := &Hx_Obj_haxe_format_jsonprinter{}
	obj.VTable = obj
	obj.Hx_Field_buf = HxDefault[*Hx_Obj_stringbuf]()
	obj.Hx_Field_replacer = HxDefault[func(key any, value any) any]()
	obj.Hx_Field_indent = HxDefault[string]()
	obj.Hx_Field_pretty = HxDefault[bool]()
	obj.Hx_Field_nind = HxDefault[int32]()
	obj.Hx_Field_stack = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_depth = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_format_jsonprinter_CreateInstance(replacer func(key any, value any) any, space HxNullable[string]) *Hx_Obj_haxe_format_jsonprinter {
	obj := Hx_Obj_haxe_format_jsonprinter_CreateEmptyInstance()
	obj.Hx_New(replacer, space)
	return obj
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_New(replacer func(key any, value any) any, space HxNullable[string]) {
	{
		this.Hx_Field_stack = (func() *Hx_Obj_haxe_imap {
			_hx_icast_31 := Hx_Obj_haxe_ds_objectmap_CreateInstance()
			if _hx_icast_31 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_31, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_31.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_depth = 0
	}
	{
		this.Hx_Field_replacer = replacer
		this.Hx_Field_indent = (space).Value
		this.Hx_Field_pretty = ((space).Valid != false)
		this.Hx_Field_nind = 0
		this.Hx_Field_buf = Hx_Obj_stringbuf_CreateInstance()
	}

}

var Hx_Obj_haxe_format_jsonprinter_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.format.JsonPrinter",
	HxMakeArray[string]("print"),
	HxMakeArray[string]("buf", "replacer", "indent", "pretty", "nind", "stack", "depth", "ipad", "newl", "write", "addChar", "add", "classString", "objString", "beginStack", "endStack", "fieldsString", "quote"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_format_jsonprinter_CreateInstance(HxConvert[func(key any, value any) any](params.Get(0)), HxConvert[HxNullable[string]](params.Get(1)))
	},
	func() any {
		return Hx_Obj_haxe_format_jsonprinter_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_format_jsonprinter_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "print":
			return Hx_Field_haxe_format_jsonprinter_print
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_format_jsonprinter_RTTI
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_As_Hx_Obj_haxe_format_jsonprinter() *Hx_Obj_haxe_format_jsonprinter {
	return this
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_ipad() {
	if this.Hx_Field_pretty {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_indent
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_nind
		var v string
		_ = v
		v = Hx_Field_stringtools_lpad("", _hx_tmp_0, (_hx_tmp_1 * int32(utf8.RuneCountInString(this.Hx_Field_indent))))
		func() func(x string) {
			var _hx_closure_src_173 func(x any)
			_ = _hx_closure_src_173
			_hx_closure_src_173 = this.Hx_Field_buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_173(((any)(x)))
			}
		}()(v)
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_newl() {
	if this.Hx_Field_pretty {
		this.Hx_Field_buf.VTable.Hx_Field_addChar(10)
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_write(k any, v any) {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(this.Hx_Field_replacer)), ((any)(nil))))) {
		v = this.Hx_Field_replacer(k, v)
	}

	{
		var _g Hx_Enum_valuetype
		_ = _g
		_g = Hx_Field_type_typeof(v)
		switch _g.Hx_Field_enumIndex() {
		case 0:
			{
				func() func(x string) {
					var _hx_closure_src_174 func(x any)
					_ = _hx_closure_src_174
					_hx_closure_src_174 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_174(((any)(x)))
					}
				}()("null")
			}

		case 1, 2:
			{
				func() func(x string) {
					var _hx_closure_src_175 func(x any)
					_ = _hx_closure_src_175
					_hx_closure_src_175 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_175(((any)(x)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(v))
			}

		case 3:
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				if Hx_Field_math_isFinite(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(v)))) {
					_hx_tmp_0 = Hx_Field_std_string(v)
				} else {
					_hx_tmp_0 = "null"
				}

				var v1 string
				_ = v1
				v1 = _hx_tmp_0
				func() func(x string) {
					var _hx_closure_src_176 func(x any)
					_ = _hx_closure_src_176
					_hx_closure_src_176 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_176(((any)(x)))
					}
				}()(v1)
			}

		case 4:
			{
				func() func(x string) {
					var _hx_closure_src_177 func(x any)
					_ = _hx_closure_src_177
					_hx_closure_src_177 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_177(((any)(x)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(v))
			}

		case 5:
			{
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = v
				this.VTable.Hx_Field_fieldsString(_hx_tmp_1, Hx_Field_reflect_fields(v))
			}

		case 6:
			{
				func() func(x string) {
					var _hx_closure_src_178 func(x any)
					_ = _hx_closure_src_178
					_hx_closure_src_178 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_178(((any)(x)))
					}
				}()("\"<fun>\"")
			}

		case 7:
			{
				var c *Hx_Obj_go_haxe_hxclass
				_ = c
				c = _g.(Hx_Enum_valuetype_TClass).Hx_Field_c
				if c == Hx_Obj_string_RTTI {
					this.VTable.Hx_Field_quote(Hx_Field_go_haxe_hxdynamic_toString(v))
				} else {
					if c == Hx_Obj_array_RTTI {
						var v1 HxArray[any]
						_ = v1
						v1 = HxAnyToArray(v)
						this.Hx_Field_buf.VTable.Hx_Field_addChar(91)
						var _hx_reserved_len int32
						_ = _hx_reserved_len
						_hx_reserved_len = Hx_Field_go_haxe_hxdynamic_getArrayLength(v1)
						var last int32
						_ = last
						last = (_hx_reserved_len - 1)
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							var _g1 int32
							_ = _g1
							_g1 = _hx_reserved_len
							for tmp__g_1 < _g1 {
								var _hx_tmp_2 int32
								_ = _hx_tmp_2
								_hx_tmp_2 = tmp__g_1
								tmp__g_1 = (tmp__g_1 + 1)
								var i int32
								_ = i
								i = _hx_tmp_2
								if i > 0 {
									this.Hx_Field_buf.VTable.Hx_Field_addChar(44)
								} else {
									this.Hx_Field_nind++
								}

								if this.Hx_Field_pretty {
									this.Hx_Field_buf.VTable.Hx_Field_addChar(10)
								}

								if this.Hx_Field_pretty {
									var _hx_tmp_3 string
									_ = _hx_tmp_3
									_hx_tmp_3 = this.Hx_Field_indent
									var _hx_tmp_4 int32
									_ = _hx_tmp_4
									_hx_tmp_4 = this.Hx_Field_nind
									var v string
									_ = v
									v = Hx_Field_stringtools_lpad("", _hx_tmp_3, (_hx_tmp_4 * int32(utf8.RuneCountInString(this.Hx_Field_indent))))
									func() func(x string) {
										var _hx_closure_src_179 func(x any)
										_ = _hx_closure_src_179
										_hx_closure_src_179 = this.Hx_Field_buf.VTable.Hx_Field_add
										return func(x string) {
											_hx_closure_src_179(((any)(x)))
										}
									}()(v)
								}

								this.VTable.Hx_Field_write(i, v1.Get(i))
								if i == last {
									this.Hx_Field_nind--
									if this.Hx_Field_pretty {
										this.Hx_Field_buf.VTable.Hx_Field_addChar(10)
									}

									if this.Hx_Field_pretty {
										var _hx_tmp_5 string
										_ = _hx_tmp_5
										_hx_tmp_5 = this.Hx_Field_indent
										var _hx_tmp_6 int32
										_ = _hx_tmp_6
										_hx_tmp_6 = this.Hx_Field_nind
										var v string
										_ = v
										v = Hx_Field_stringtools_lpad("", _hx_tmp_5, (_hx_tmp_6 * int32(utf8.RuneCountInString(this.Hx_Field_indent))))
										func() func(x string) {
											var _hx_closure_src_180 func(x any)
											_ = _hx_closure_src_180
											_hx_closure_src_180 = this.Hx_Field_buf.VTable.Hx_Field_add
											return func(x string) {
												_hx_closure_src_180(((any)(x)))
											}
										}()(v)
									}
								}
							}
						}

						this.Hx_Field_buf.VTable.Hx_Field_addChar(93)
					} else {
						if c == Hx_Obj_haxe_ds_stringmap_RTTI {
							var hx_dyncast_32 any
							_ = hx_dyncast_32
							hx_dyncast_32 = Hx_Field_go_haxe_hxdynamic_castClass(v, "Hx_Obj_haxe_ds_stringmap")
							var _hx_tmp_7 *Hx_Obj_haxe_ds_stringmap
							_ = _hx_tmp_7
							if hx_dyncast_32 != nil {
								_hx_tmp_7 = hx_dyncast_32.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap()
							} else {
								_hx_tmp_7 = nil
							}

							var v1 *Hx_Obj_haxe_ds_stringmap
							_ = v1
							v1 = _hx_tmp_7
							{
								var v2 any
								_ = v2
								v2 = ((any)(v1))
								this.Hx_Field_depth++
								if this.Hx_Field_depth > 10 {
									if this.Hx_Field_stack.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
									}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_exists(v2) {
										panic(Hx_Field_haxe_exception_thrown(((any)("Cannot stringify circular structure"))))
									}

									this.Hx_Field_stack.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
									}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_set(v2, true)
								}
							}

							var o any
							_ = o
							o = any(map[string]any{})
							{
								var k any
								_ = k
								k = v1.VTable.Hx_Field_keys()
								for {
									var _hx_tmp_8 any
									_ = _hx_tmp_8
									_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(k, "hasNext")
									if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_8, HxMakeArray[any]().Dyn()))) {
										break
									}

									var _hx_tmp_9 any
									_ = _hx_tmp_9
									_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(k, "next")
									var k1 string
									_ = k1
									k1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any]().Dyn()))
									var _hx_tmp_10 any
									_ = _hx_tmp_10
									_hx_tmp_10 = o
									var _hx_tmp_11 string
									_ = _hx_tmp_11
									_hx_tmp_11 = k1
									Hx_Field_reflect_setField(_hx_tmp_10, _hx_tmp_11, v1.VTable.Hx_Field_get(k1))
								}
							}

							{
								var v3 any
								_ = v3
								v3 = o
								var _hx_tmp_12 any
								_ = _hx_tmp_12
								_hx_tmp_12 = v3
								this.VTable.Hx_Field_fieldsString(_hx_tmp_12, Hx_Field_reflect_fields(v3))
							}

							{
								if this.Hx_Field_depth > 10 {
									this.Hx_Field_stack.VTable.(interface {
										Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
									}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_remove(((any)(v1)))
								}

								this.Hx_Field_depth--
							}
						} else {
							if c == Hx_Obj_date_RTTI {
								var hx_dyncast_33 any
								_ = hx_dyncast_33
								hx_dyncast_33 = Hx_Field_go_haxe_hxdynamic_castClass(v, "Hx_Obj_date")
								var _hx_tmp_13 *Hx_Obj_date
								_ = _hx_tmp_13
								if hx_dyncast_33 != nil {
									_hx_tmp_13 = hx_dyncast_33.(interface{ Hx_As_Hx_Obj_date() *Hx_Obj_date }).Hx_As_Hx_Obj_date()
								} else {
									_hx_tmp_13 = nil
								}

								var v1 *Hx_Obj_date
								_ = v1
								v1 = _hx_tmp_13
								this.VTable.Hx_Field_quote(v1.VTable.Hx_Field_toString())
							} else {
								this.VTable.Hx_Field_classString(v)
							}
						}
					}
				}
			}

		case 8:
			{
				var i int32
				_ = i
				i = Hx_Field_type_enumIndex(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(v)
					if _hx_d == nil {
						var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
						return _hx_z
					}
					return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
				}())
				{
					var v1 string
					_ = v1
					v1 = Hx_Field_std_string(i)
					func() func(x string) {
						var _hx_closure_src_181 func(x any)
						_ = _hx_closure_src_181
						_hx_closure_src_181 = this.Hx_Field_buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_181(((any)(x)))
						}
					}()(v1)
				}
			}

		case 9:
			{
				func() func(x string) {
					var _hx_closure_src_182 func(x any)
					_ = _hx_closure_src_182
					_hx_closure_src_182 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_182(((any)(x)))
					}
				}()("\"???\"")
			}

		}
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_classString(v any) {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = v
	this.VTable.Hx_Field_fieldsString(_hx_tmp_0, Hx_Field_type_getInstanceFields(Hx_Field_type_getClass(v)))
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_objString(v any) {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = v
	this.VTable.Hx_Field_fieldsString(_hx_tmp_0, Hx_Field_reflect_fields(v))
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_beginStack(v any) {
	this.Hx_Field_depth++
	if this.Hx_Field_depth > 10 {
		if this.Hx_Field_stack.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
		}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_exists(v) {
			panic(Hx_Field_haxe_exception_thrown(((any)("Cannot stringify circular structure"))))
		}

		this.Hx_Field_stack.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
		}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_set(v, true)
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_endStack(v any) {
	if this.Hx_Field_depth > 10 {
		this.Hx_Field_stack.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
		}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_remove(v)
	}

	this.Hx_Field_depth--
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_fieldsString(v any, fields HxArray[string]) {
	{
		this.Hx_Field_depth++
		if this.Hx_Field_depth > 10 {
			if this.Hx_Field_stack.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
			}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_exists(v) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Cannot stringify circular structure"))))
			}

			this.Hx_Field_stack.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
			}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_set(v, true)
		}
	}

	this.Hx_Field_buf.VTable.Hx_Field_addChar(123)
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = fields.Len()
	var empty bool
	_ = empty
	empty = true
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var f string
			_ = f
			f = fields.Get(i)
			var value any
			_ = value
			value = Hx_Field_reflect_field(v, f)
			if Hx_Field_reflect_isFunction(value) {
				continue
			}

			if empty {
				this.Hx_Field_nind++
				empty = false
			} else {
				this.Hx_Field_buf.VTable.Hx_Field_addChar(44)
			}

			if this.Hx_Field_pretty {
				this.Hx_Field_buf.VTable.Hx_Field_addChar(10)
			}

			if this.Hx_Field_pretty {
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				_hx_tmp_1 = this.Hx_Field_indent
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = this.Hx_Field_nind
				var v string
				_ = v
				v = Hx_Field_stringtools_lpad("", _hx_tmp_1, (_hx_tmp_2 * int32(utf8.RuneCountInString(this.Hx_Field_indent))))
				func() func(x string) {
					var _hx_closure_src_183 func(x any)
					_ = _hx_closure_src_183
					_hx_closure_src_183 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_183(((any)(x)))
					}
				}()(v)
			}

			this.VTable.Hx_Field_quote(f)
			this.Hx_Field_buf.VTable.Hx_Field_addChar(58)
			if this.Hx_Field_pretty {
				this.Hx_Field_buf.VTable.Hx_Field_addChar(32)
			}

			this.VTable.Hx_Field_write(f, value)
		}
	}

	if !empty {
		this.Hx_Field_nind--
		if this.Hx_Field_pretty {
			this.Hx_Field_buf.VTable.Hx_Field_addChar(10)
		}

		if this.Hx_Field_pretty {
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = this.Hx_Field_indent
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = this.Hx_Field_nind
			var v string
			_ = v
			v = Hx_Field_stringtools_lpad("", _hx_tmp_3, (_hx_tmp_4 * int32(utf8.RuneCountInString(this.Hx_Field_indent))))
			func() func(x string) {
				var _hx_closure_src_184 func(x any)
				_ = _hx_closure_src_184
				_hx_closure_src_184 = this.Hx_Field_buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_184(((any)(x)))
				}
			}()(v)
		}
	}

	this.Hx_Field_buf.VTable.Hx_Field_addChar(125)
	{
		if this.Hx_Field_depth > 10 {
			this.Hx_Field_stack.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
			}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_remove(v)
		}

		this.Hx_Field_depth--
	}
}

func (this *Hx_Obj_haxe_format_jsonprinter) Hx_Field_quote(s string) {
	this.Hx_Field_buf.VTable.Hx_Field_addChar(34)
	var i int32
	_ = i
	i = 0
	var length int32
	_ = length
	length = int32(utf8.RuneCountInString(s))
	for i < length {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = s
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = i
		i = (i + 1)
		var c int32
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, _hx_tmp_1)).Value
		switch c {
		case 8:
			{
				func() func(x string) {
					var _hx_closure_src_185 func(x any)
					_ = _hx_closure_src_185
					_hx_closure_src_185 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_185(((any)(x)))
					}
				}()("\\b")
			}

		case 9:
			{
				func() func(x string) {
					var _hx_closure_src_186 func(x any)
					_ = _hx_closure_src_186
					_hx_closure_src_186 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_186(((any)(x)))
					}
				}()("\\t")
			}

		case 10:
			{
				func() func(x string) {
					var _hx_closure_src_187 func(x any)
					_ = _hx_closure_src_187
					_hx_closure_src_187 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_187(((any)(x)))
					}
				}()("\\n")
			}

		case 12:
			{
				func() func(x string) {
					var _hx_closure_src_188 func(x any)
					_ = _hx_closure_src_188
					_hx_closure_src_188 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_188(((any)(x)))
					}
				}()("\\f")
			}

		case 13:
			{
				func() func(x string) {
					var _hx_closure_src_189 func(x any)
					_ = _hx_closure_src_189
					_hx_closure_src_189 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_189(((any)(x)))
					}
				}()("\\r")
			}

		case 34:
			{
				func() func(x string) {
					var _hx_closure_src_190 func(x any)
					_ = _hx_closure_src_190
					_hx_closure_src_190 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_190(((any)(x)))
					}
				}()("\\\"")
			}

		case 92:
			{
				func() func(x string) {
					var _hx_closure_src_191 func(x any)
					_ = _hx_closure_src_191
					_hx_closure_src_191 = this.Hx_Field_buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_191(((any)(x)))
					}
				}()("\\\\")
			}

		default:
			{
				this.Hx_Field_buf.VTable.Hx_Field_addChar(c)
			}
		}
	}

	this.Hx_Field_buf.VTable.Hx_Field_addChar(34)
}

func Hx_Field_haxe_format_jsonprinter_print(o any, replacer__isOptional any, space__isOptional any) string {
	replacer := HxDefault[func(key any, value any) any]()
	_ = replacer

	if replacer__isOptional != nil {
		replacer = HxConvert[func(key any, value any) any](replacer__isOptional)
	} else {
		replacer = nil
	}
	space := HxDefault[HxNullable[string]]()
	_ = space

	if space__isOptional != nil {
		space = HxConvert[HxNullable[string]](space__isOptional)
	} else {

	}
	{
		if (space).Valid == false {
			var hx_nullable_t31 HxNullable[string]
			_ = hx_nullable_t31
			hx_nullable_t31 = HxNullable[string]{Value: HxStringNull}

			space = HxNullable[string]{Value: (hx_nullable_t31).Value, Valid: (hx_nullable_t31).Valid}
		}

		var printer *Hx_Obj_haxe_format_jsonprinter
		_ = printer
		printer = Hx_Obj_haxe_format_jsonprinter_CreateInstance(replacer, space)
		printer.VTable.Hx_Field_write("", o)
		return printer.Hx_Field_buf.VTable.Hx_Field_toString()
	}
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
