package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_systools() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_systools() {
	Hx_Field_haxe_systools_winMetaCharacters = Hx_Init_Hx_Field_haxe_systools_winMetaCharacters()
}

type Hx_Obj_VTable_haxe_systools interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_systools struct {
	VTable Hx_Obj_VTable_haxe_systools `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_systools_CreateEmptyInstance() *Hx_Obj_haxe_systools {
	obj := &Hx_Obj_haxe_systools{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_systools_CreateInstance() *Hx_Obj_haxe_systools {
	obj := Hx_Obj_haxe_systools_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_systools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.SysTools",
	HxMakeArray[string]("winMetaCharacters", "quoteUnixArg", "quoteWinArg"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_systools_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_systools_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_systools_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "winMetaCharacters":
			return &Hx_Field_haxe_systools_winMetaCharacters
		case "quoteUnixArg":
			return Hx_Field_haxe_systools_quoteUnixArg
		case "quoteWinArg":
			return Hx_Field_haxe_systools_quoteWinArg
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_systools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_systools_RTTI
}

func (this *Hx_Obj_haxe_systools) Hx_As_Hx_Obj_haxe_systools() *Hx_Obj_haxe_systools {
	return this
}

var Hx_Field_haxe_systools_winMetaCharacters HxArray[any]

func Hx_Init_Hx_Field_haxe_systools_winMetaCharacters() HxArray[any] {
	return HxMakeArrayView[any](HxMakeArray[int32](32, 40, 41, 37, 33, 94, 34, 60, 62, 38, 124, 10, 13, 44, 59))
}

func Hx_Field_haxe_systools_quoteUnixArg(argument string) string {
	if argument == "" {
		return "''"
	}

	if !Hx_Obj_ereg_CreateInstance("[^a-zA-Z0-9_@%+=:,./-]", "").VTable.Hx_Field_match(argument) {
		return argument
	}

	return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("'", Hx_Field_stringtools_replace(argument, "'", "'\"'\"'")), "'")
}

func Hx_Field_haxe_systools_quoteWinArg(argument string, escapeMetaCharacters bool) string {
	if !Hx_Obj_ereg_CreateInstance("^(/)?[^ \t/\\\\\"]+$", "").VTable.Hx_Field_match(argument) {
		var result *Hx_Obj_stringbuf
		_ = result
		result = Hx_Obj_stringbuf_CreateInstance()
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = argument
		if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, " ", HxNullable[int32]{})) != -1 {
			_hx_tmp_2 = true
		} else {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = argument
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = _hx_tmp_4
			_hx_tmp_2 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "\t", HxNullable[int32]{})) != -1)
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (argument == "")
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = argument
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_0 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_7, "/", HxNullable[int32]{})) > 0)
		}

		var needquote bool
		_ = needquote
		needquote = _hx_tmp_0
		if needquote {
			func() func(x string) {
				var _hx_closure_src_4 func(x any)
				_ = _hx_closure_src_4
				_hx_closure_src_4 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_4(((any)(x)))
				}
			}()("\"")
		}

		var bs_buf *Hx_Obj_stringbuf
		_ = bs_buf
		bs_buf = Hx_Obj_stringbuf_CreateInstance()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = int32(utf8.RuneCountInString(argument))
			for _g < _g1 {
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_8
				{
					var tmp__g1_1 HxNullable[int32]
					_ = tmp__g1_1
					tmp__g1_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(argument, i))
					if (tmp__g1_1).Valid == false {
						var c HxNullable[int32]
						_ = c
						c = tmp__g1_1
						{
							if bs_buf.VTable.Hx_Field_get_length() > 0 {
								func() func(x string) {
									var _hx_closure_src_5 func(x any)
									_ = _hx_closure_src_5
									_hx_closure_src_5 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_5(((any)(x)))
									}
								}()(bs_buf.VTable.Hx_Field_toString())
								bs_buf = Hx_Obj_stringbuf_CreateInstance()
							}

							result.VTable.Hx_Field_addChar((c).Value)
						}
					} else {
						switch (tmp__g1_1).Value {
						case 34:
							{
								var bs string
								_ = bs
								bs = bs_buf.VTable.Hx_Field_toString()
								func() func(x string) {
									var _hx_closure_src_6 func(x any)
									_ = _hx_closure_src_6
									_hx_closure_src_6 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_6(((any)(x)))
									}
								}()(bs)
								func() func(x string) {
									var _hx_closure_src_7 func(x any)
									_ = _hx_closure_src_7
									_hx_closure_src_7 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_7(((any)(x)))
									}
								}()(bs)
								bs_buf = Hx_Obj_stringbuf_CreateInstance()
								func() func(x string) {
									var _hx_closure_src_8 func(x any)
									_ = _hx_closure_src_8
									_hx_closure_src_8 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_8(((any)(x)))
									}
								}()("\\\"")
							}

						case 92:
							{
								func() func(x string) {
									var _hx_closure_src_9 func(x any)
									_ = _hx_closure_src_9
									_hx_closure_src_9 = bs_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_9(((any)(x)))
									}
								}()("\\")
							}

						default:
							{
								var c HxNullable[int32]
								_ = c
								c = tmp__g1_1
								{
									if bs_buf.VTable.Hx_Field_get_length() > 0 {
										func() func(x string) {
											var _hx_closure_src_10 func(x any)
											_ = _hx_closure_src_10
											_hx_closure_src_10 = result.VTable.Hx_Field_add
											return func(x string) {
												_hx_closure_src_10(((any)(x)))
											}
										}()(bs_buf.VTable.Hx_Field_toString())
										bs_buf = Hx_Obj_stringbuf_CreateInstance()
									}

									result.VTable.Hx_Field_addChar((c).Value)
								}
							}
						}
					}
				}
			}
		}

		func() func(x string) {
			var _hx_closure_src_11 func(x any)
			_ = _hx_closure_src_11
			_hx_closure_src_11 = result.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_11(((any)(x)))
			}
		}()(bs_buf.VTable.Hx_Field_toString())
		if needquote {
			func() func(x string) {
				var _hx_closure_src_12 func(x any)
				_ = _hx_closure_src_12
				_hx_closure_src_12 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_12(((any)(x)))
				}
			}()(bs_buf.VTable.Hx_Field_toString())
			func() func(x string) {
				var _hx_closure_src_13 func(x any)
				_ = _hx_closure_src_13
				_hx_closure_src_13 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_13(((any)(x)))
				}
			}()("\"")
		}

		argument = result.VTable.Hx_Field_toString()
	}

	if escapeMetaCharacters {
		var result *Hx_Obj_stringbuf
		_ = result
		result = Hx_Obj_stringbuf_CreateInstance()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = int32(utf8.RuneCountInString(argument))
			for _g < _g1 {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_9
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(argument, i))
				var _hx_tmp_10 HxArray[int32]
				_ = _hx_tmp_10
				_hx_tmp_10 = HxMakeArrayView[int32](Hx_Field_haxe_systools_winMetaCharacters)
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = (c).Value
				if (Hx_Array_IndexOf[int32](_hx_tmp_10, _hx_tmp_11, HxNullable[int32]{})) >= 0 {
					result.VTable.Hx_Field_addChar(94)
				}

				result.VTable.Hx_Field_addChar((c).Value)
			}
		}

		return result.VTable.Hx_Field_toString()
	} else {
		return argument
	}
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[int32]()
}
