package core

import "unicode/utf8"

type Hx_Typedef_loreline_importsfilehandler func(path string, callback func(data string))

type Hx_Typedef_loreline_importserrorhandler func(error *Hx_Obj_loreline_error)

type Hx_Typedef_loreline_importscallback func(hasErrors bool, resolvedImports Hx_Typedef_map)

func Hx_ClassInit_Hx_Obj_VTable_loreline__imports_importsloopinfo() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__imports_importsloopinfo() {
}

type Hx_Obj_VTable_loreline__imports_importsloopinfo interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__imports_importsloopinfo struct {
	VTable            Hx_Obj_VTable_loreline__imports_importsloopinfo `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_finished bool
}

func Hx_Obj_loreline__imports_importsloopinfo_CreateEmptyInstance() *Hx_Obj_loreline__imports_importsloopinfo {
	obj := &Hx_Obj_loreline__imports_importsloopinfo{}
	obj.VTable = obj
	obj.Hx_Field_finished = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline__imports_importsloopinfo_CreateInstance(finished bool) *Hx_Obj_loreline__imports_importsloopinfo {
	obj := Hx_Obj_loreline__imports_importsloopinfo_CreateEmptyInstance()
	obj.Hx_New(finished)
	return obj
}

func (this *Hx_Obj_loreline__imports_importsloopinfo) Hx_New(finished bool) {
	{
		this.Hx_Field_finished = finished
	}

}

var Hx_Obj_loreline__imports_importsloopinfo_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Imports.ImportsLoopInfo",
	HxMakeArray[string](),
	HxMakeArray[string]("finished"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__imports_importsloopinfo_CreateInstance(HxConvert[bool](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__imports_importsloopinfo_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__imports_importsloopinfo_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__imports_importsloopinfo) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__imports_importsloopinfo_RTTI
}

func (this *Hx_Obj_loreline__imports_importsloopinfo) Hx_As_Hx_Obj_loreline__imports_importsloopinfo() *Hx_Obj_loreline__imports_importsloopinfo {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_imports() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_imports() {
}

type Hx_Obj_VTable_loreline_imports interface {
	Hx_Field_resolve(rootPath string, tokens Hx_Typedef_loreline_tokens, handleFile Hx_Typedef_loreline_importsfilehandler, handleError Hx_Typedef_loreline_importserrorhandler, done Hx_Typedef_loreline_importscallback)
	Hx_Field_processImports(toImport HxArray[string], visitedImports Hx_Typedef_map, resolvedImports Hx_Typedef_map)
	Hx_Field_handleItemInLoop(item string, loopInfo *Hx_Obj_loreline__imports_importsloopinfo, toImport HxArray[string], visitedImports Hx_Typedef_map, resolvedImports Hx_Typedef_map)
	Hx_Field_extractImports(cwd string, tokens Hx_Typedef_loreline_tokens, toImport HxArray[string], visitedImports Hx_Typedef_map)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_imports struct {
	VTable                    Hx_Obj_VTable_loreline_imports `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_handleFile       Hx_Typedef_loreline_importsfilehandler
	Hx_Field_handleError      Hx_Typedef_loreline_importserrorhandler
	Hx_Field_tokens           Hx_Typedef_loreline_tokens
	Hx_Field_rootPath         string
	Hx_Field_autoAddExtension bool
	Hx_Field_resolvedImports  Hx_Typedef_map
	Hx_Field_done             Hx_Typedef_loreline_importscallback
	Hx_Field_hasErrors        bool
	Hx_Field_pendingImports   int32
	Hx_Field_ext              string
}

func Hx_Obj_loreline_imports_CreateEmptyInstance() *Hx_Obj_loreline_imports {
	obj := &Hx_Obj_loreline_imports{}
	obj.VTable = obj
	obj.Hx_Field_handleFile = HxDefault[Hx_Typedef_loreline_importsfilehandler]()
	obj.Hx_Field_handleError = HxDefault[Hx_Typedef_loreline_importserrorhandler]()
	obj.Hx_Field_tokens = HxDefault[Hx_Typedef_loreline_tokens]()
	obj.Hx_Field_rootPath = HxDefault[string]()
	obj.Hx_Field_autoAddExtension = HxDefault[bool]()
	obj.Hx_Field_resolvedImports = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_done = HxDefault[Hx_Typedef_loreline_importscallback]()
	obj.Hx_Field_hasErrors = HxDefault[bool]()
	obj.Hx_Field_pendingImports = HxDefault[int32]()
	obj.Hx_Field_ext = HxDefault[string]()
	return obj
}

func Hx_Obj_loreline_imports_CreateInstance() *Hx_Obj_loreline_imports {
	obj := Hx_Obj_loreline_imports_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_imports) Hx_New() {
	{
		this.Hx_Field_autoAddExtension = true
	}
	{
		this.Hx_Field_resolvedImports = nil
	}
	{
		this.Hx_Field_done = ((Hx_Typedef_loreline_importscallback)(nil))
	}
	{
		this.Hx_Field_hasErrors = false
	}
	{
		this.Hx_Field_pendingImports = 0
	}
	{
		this.Hx_Field_ext = ".lor"
	}
	{
	}

}

var Hx_Obj_loreline_imports_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Imports",
	HxMakeArray[string]("sameLetterCi", "endsWithLorTxt", "endsWithLor", "lorExtension", "isLorFilePath", "stripLorExtension"),
	HxMakeArray[string]("handleFile", "handleError", "tokens", "rootPath", "autoAddExtension", "resolvedImports", "done", "hasErrors", "pendingImports", "ext", "resolve", "processImports", "handleItemInLoop", "extractImports"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_imports_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_imports_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_imports_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "sameLetterCi":
			return Hx_Field_loreline_imports_sameLetterCi
		case "endsWithLorTxt":
			return Hx_Field_loreline_imports_endsWithLorTxt
		case "endsWithLor":
			return Hx_Field_loreline_imports_endsWithLor
		case "lorExtension":
			return Hx_Field_loreline_imports_lorExtension
		case "isLorFilePath":
			return Hx_Field_loreline_imports_isLorFilePath
		case "stripLorExtension":
			return Hx_Field_loreline_imports_stripLorExtension
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_imports) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_imports_RTTI
}

func (this *Hx_Obj_loreline_imports) Hx_As_Hx_Obj_loreline_imports() *Hx_Obj_loreline_imports {
	return this
}

func (this *Hx_Obj_loreline_imports) Hx_Field_resolve(rootPath string, tokens Hx_Typedef_loreline_tokens, handleFile Hx_Typedef_loreline_importsfilehandler, handleError Hx_Typedef_loreline_importserrorhandler, done Hx_Typedef_loreline_importscallback) {
	this.Hx_Field_rootPath = rootPath
	this.Hx_Field_tokens = tokens
	this.Hx_Field_handleFile = handleFile
	this.Hx_Field_handleError = handleError
	this.Hx_Field_done = done
	this.Hx_Field_hasErrors = false
	this.Hx_Field_pendingImports = 0
	var tmp bool
	_ = tmp
	if rootPath == HxStringNull {
		tmp = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(rootPath))
		if n < 8 {
			tmp = false
		} else {
			var tmp1 bool
			_ = tmp1
			var tmp2 bool
			_ = tmp2
			var tmp3 bool
			_ = tmp3
			var tmp4 bool
			_ = tmp4
			var tmp5 bool
			_ = tmp5
			var tmp6 bool
			_ = tmp6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 8)))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 7)))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_1
				tmp6 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp6 = false
			}

			if tmp6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 6)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				tmp5 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp5 = false
			}

			if tmp5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 5)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				tmp4 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp4 = false
			}

			if tmp4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				tmp3 = ((_hx_tmp_4) == 46)
			} else {
				tmp3 = false
			}

			if tmp3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 3)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_5
				tmp2 = ((c1 == 116) || (c1 == 84))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 2)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				tmp1 = ((c1 == 120) || (c1 == 88))
			} else {
				tmp1 = false
			}

			if tmp1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 1)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				tmp = ((c1 == 116) || (c1 == 84))
			} else {
				tmp = false
			}
		}
	}

	var _hx_tmp_8 string
	_ = _hx_tmp_8
	if tmp {
		_hx_tmp_8 = ".lor.txt"
	} else {
		_hx_tmp_8 = ".lor"
	}

	this.Hx_Field_ext = _hx_tmp_8
	var resolvedImports Hx_Typedef_map
	_ = resolvedImports
	resolvedImports = (func() *Hx_Obj_haxe_imap {
		_hx_icast_289 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_289 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_289, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_289.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var toImport HxArray[string]
	_ = toImport
	toImport = HxMakeArray[string]()
	var visitedImports Hx_Typedef_map
	_ = visitedImports
	visitedImports = (func() *Hx_Obj_haxe_imap {
		_hx_icast_290 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_290 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_290, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_290.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var cwd string
	_ = cwd
	cwd = Hx_Field_haxe_io_path_directory(rootPath)
	this.VTable.Hx_Field_extractImports(cwd, tokens, toImport, visitedImports)
	this.VTable.Hx_Field_processImports(toImport, visitedImports, resolvedImports)
}

func (this *Hx_Obj_loreline_imports) Hx_Field_processImports(toImport HxArray[string], visitedImports Hx_Typedef_map, resolvedImports Hx_Typedef_map) {
	var loopInfo *Hx_Obj_loreline__imports_importsloopinfo
	_ = loopInfo
	loopInfo = Hx_Obj_loreline__imports_importsloopinfo_CreateInstance(false)
	for {
		if !(toImport.Len() > 0) {
			break
		}

		var item HxNullable[string]
		_ = item
		item = (Hx_Array_Shift[string](toImport))
		this.VTable.Hx_Field_handleItemInLoop((item).Value, loopInfo, toImport, visitedImports, resolvedImports)
	}

	loopInfo.Hx_Field_finished = true
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if this.Hx_Field_pendingImports == 0 {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(this.Hx_Field_done)), ((any)(nil))))))
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		this.Hx_Field_resolvedImports = resolvedImports
		this.Hx_Field_done(this.Hx_Field_hasErrors, resolvedImports)
		this.Hx_Field_done = nil
	}
}

func (this *Hx_Obj_loreline_imports) Hx_Field_handleItemInLoop(item string, loopInfo *Hx_Obj_loreline__imports_importsloopinfo, toImport HxArray[string], visitedImports Hx_Typedef_map, resolvedImports Hx_Typedef_map) {
	var _gthis *Hx_Obj_loreline_imports
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = item
	this.Hx_Field_handleFile(_hx_tmp_0, func(data string) {
		_gthis.Hx_Field_pendingImports--
		if data != HxStringNull {
			{
				var hx_try_state int
				_ = hx_try_state

				func() {
					defer func() {
						if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
							if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
								_ = e
								{
									_gthis.Hx_Field_hasErrors = true
									if Hx_Field_std_isOfType(e, Hx_Obj_loreline_error_RTTI) {
										var hx_dyncast_291 any
										_ = hx_dyncast_291
										hx_dyncast_291 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_error")
										var _hx_tmp_2 *Hx_Obj_loreline_error
										_ = _hx_tmp_2
										if hx_dyncast_291 != nil {
											_hx_tmp_2 = hx_dyncast_291.(interface{ Hx_As_Hx_Obj_loreline_error() *Hx_Obj_loreline_error }).Hx_As_Hx_Obj_loreline_error()
										} else {
											_hx_tmp_2 = nil
										}

										_gthis.Hx_Field_handleError(_hx_tmp_2)
									} else {
										panic(Hx_Field_haxe_exception_thrown(e))
									}
								}
							} else {
								panic(_hx_ex)
							}
						}
					}()
					{
						var lexer *Hx_Obj_loreline_lexer
						_ = lexer
						lexer = Hx_Obj_loreline_lexer_CreateInstance(data)
						var tokens Hx_Typedef_loreline_tokens
						_ = tokens
						tokens = lexer.VTable.Hx_Field_tokenize()
						var lexerErrors HxArray[*Hx_Obj_loreline_lexererror]
						_ = lexerErrors
						lexerErrors = lexer.VTable.Hx_Field_getErrors()
						var _hx_tmp_1 bool
						_ = _hx_tmp_1
						if lexerErrors != nil {
							_hx_tmp_1 = (lexerErrors.Len() > 0)
						} else {
							_hx_tmp_1 = false
						}

						if _hx_tmp_1 {
							_gthis.Hx_Field_handleError((func() *Hx_Obj_loreline_error {
								_hx_up := lexerErrors.Get(0)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_error
								}
								return nil
							})())
						}

						func() func(_hx_param_key string, _hx_param_value Hx_Typedef_loreline_tokens) {
							var _hx_closure_src_688 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_688
							_hx_closure_src_688 = resolvedImports.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value Hx_Typedef_loreline_tokens) {
								_hx_closure_src_688(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(item, tokens)
						_gthis.VTable.Hx_Field_extractImports(Hx_Field_haxe_io_path_directory(item), tokens, toImport, visitedImports)
						if loopInfo.Hx_Field_finished {
							_gthis.VTable.Hx_Field_processImports(toImport, visitedImports, resolvedImports)
						}
					}
				}()
				if hx_try_state == 1 {
					return
				}
			}
		} else {
			_gthis.Hx_Field_hasErrors = true
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if loopInfo.Hx_Field_finished && (_gthis.Hx_Field_pendingImports == 0) {
			_hx_tmp_3 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(_gthis.Hx_Field_done)), ((any)(nil))))))
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			_gthis.Hx_Field_resolvedImports = resolvedImports
			_gthis.Hx_Field_done(_gthis.Hx_Field_hasErrors, resolvedImports)
			_gthis.Hx_Field_done = nil
		}
	})
}

func (this *Hx_Obj_loreline_imports) Hx_Field_extractImports(cwd string, tokens Hx_Typedef_loreline_tokens, toImport HxArray[string], visitedImports Hx_Typedef_map) {
	var i int32
	_ = i
	i = 0
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = tokens.Len()
	for i < (_hx_reserved_len - 1) {
		if (tokens.Get(i).Hx_Field_type == Hx_Enum_loreline_tokentype_KwImport{}) {
			var _g Hx_Enum_loreline_tokentype
			_ = _g
			_g = tokens.Get((i + 1)).Hx_Field_type
			if _g.Hx_Field_enumIndex() == 16 {
				var s string
				_ = s
				s = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
				{
					var path string
					_ = path
					path = s
					if !Hx_Field_haxe_io_path_isAbsolute(s) {
						path = Hx_Field_haxe_io_path_join(HxMakeArray[string](cwd, path))
					}

					path = Hx_Field_haxe_io_path_normalize(path)
					var tmp bool
					_ = tmp
					var tmp1 bool
					_ = tmp1
					if path == HxStringNull {
						tmp1 = false
					} else {
						var n int32
						_ = n
						n = int32(utf8.RuneCountInString(path))
						if n < 8 {
							tmp1 = false
						} else {
							var tmp_tmp_1 bool
							_ = tmp_tmp_1
							var tmp2 bool
							_ = tmp2
							var tmp3 bool
							_ = tmp3
							var tmp4 bool
							_ = tmp4
							var tmp5 bool
							_ = tmp5
							var tmp6 bool
							_ = tmp6
							var c HxNullable[int32]
							_ = c
							c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
							var _hx_tmp_0 int32
							_ = _hx_tmp_0
							if (c).Valid == false {
								_hx_tmp_0 = -1
							} else {
								_hx_tmp_0 = (c).Value
							}

							if (_hx_tmp_0) == 46 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
								var _hx_tmp_1 int32
								_ = _hx_tmp_1
								if (tmp_c_1).Valid == false {
									_hx_tmp_1 = -1
								} else {
									_hx_tmp_1 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_1
								tmp6 = ((c1 == 108) || (c1 == 76))
							} else {
								tmp6 = false
							}

							if tmp6 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
								var _hx_tmp_2 int32
								_ = _hx_tmp_2
								if (tmp_c_1).Valid == false {
									_hx_tmp_2 = -1
								} else {
									_hx_tmp_2 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_2
								tmp5 = ((c1 == 111) || (c1 == 79))
							} else {
								tmp5 = false
							}

							if tmp5 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
								var _hx_tmp_3 int32
								_ = _hx_tmp_3
								if (tmp_c_1).Valid == false {
									_hx_tmp_3 = -1
								} else {
									_hx_tmp_3 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_3
								tmp4 = ((c1 == 114) || (c1 == 82))
							} else {
								tmp4 = false
							}

							if tmp4 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
								var _hx_tmp_4 int32
								_ = _hx_tmp_4
								if (tmp_c_1).Valid == false {
									_hx_tmp_4 = -1
								} else {
									_hx_tmp_4 = (tmp_c_1).Value
								}

								tmp3 = ((_hx_tmp_4) == 46)
							} else {
								tmp3 = false
							}

							if tmp3 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
								var _hx_tmp_5 int32
								_ = _hx_tmp_5
								if (tmp_c_1).Valid == false {
									_hx_tmp_5 = -1
								} else {
									_hx_tmp_5 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_5
								tmp2 = ((c1 == 116) || (c1 == 84))
							} else {
								tmp2 = false
							}

							if tmp2 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
								var _hx_tmp_6 int32
								_ = _hx_tmp_6
								if (tmp_c_1).Valid == false {
									_hx_tmp_6 = -1
								} else {
									_hx_tmp_6 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_6
								tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
							} else {
								tmp_tmp_1 = false
							}

							if tmp_tmp_1 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
								var _hx_tmp_7 int32
								_ = _hx_tmp_7
								if (tmp_c_1).Valid == false {
									_hx_tmp_7 = -1
								} else {
									_hx_tmp_7 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_7
								tmp1 = ((c1 == 116) || (c1 == 84))
							} else {
								tmp1 = false
							}
						}
					}

					if !tmp1 {
						if path == HxStringNull {
							tmp = false
						} else {
							var n int32
							_ = n
							n = int32(utf8.RuneCountInString(path))
							if n < 4 {
								tmp = false
							} else {
								var tmp_tmp1_1 bool
								_ = tmp_tmp1_1
								var tmp2 bool
								_ = tmp2
								var c HxNullable[int32]
								_ = c
								c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
								var _hx_tmp_8 int32
								_ = _hx_tmp_8
								if (c).Valid == false {
									_hx_tmp_8 = -1
								} else {
									_hx_tmp_8 = (c).Value
								}

								if (_hx_tmp_8) == 46 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
									var _hx_tmp_9 int32
									_ = _hx_tmp_9
									if (tmp_c_1).Valid == false {
										_hx_tmp_9 = -1
									} else {
										_hx_tmp_9 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_9
									tmp2 = ((c1 == 108) || (c1 == 76))
								} else {
									tmp2 = false
								}

								if tmp2 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
									var _hx_tmp_10 int32
									_ = _hx_tmp_10
									if (tmp_c_1).Valid == false {
										_hx_tmp_10 = -1
									} else {
										_hx_tmp_10 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_10
									tmp_tmp1_1 = ((c1 == 111) || (c1 == 79))
								} else {
									tmp_tmp1_1 = false
								}

								if tmp_tmp1_1 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
									var _hx_tmp_11 int32
									_ = _hx_tmp_11
									if (tmp_c_1).Valid == false {
										_hx_tmp_11 = -1
									} else {
										_hx_tmp_11 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_11
									tmp = ((c1 == 114) || (c1 == 82))
								} else {
									tmp = false
								}
							}
						}
					} else {
						tmp = true
					}

					if !tmp {
						path = Hx_Field_go_haxe_hxdynamic_stringConcat(path, this.Hx_Field_ext)
					}

					if !func() func(_hx_param_key string) bool {
						var _hx_closure_src_689 func(_hx_param_key any) bool
						_ = _hx_closure_src_689
						_hx_closure_src_689 = visitedImports.VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_689(((any)(_hx_param_key)))
						}
					}()(path) {
						this.Hx_Field_pendingImports++
						func() func(_hx_param_key string, _hx_param_value bool) {
							var _hx_closure_src_690 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_690
							_hx_closure_src_690 = visitedImports.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value bool) {
								_hx_closure_src_690(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(path, true)
						_ = (Hx_Array_Push[string](toImport, path))
					}
				}
			}
		}

		i++
	}
}

func Hx_Field_loreline_imports_sameLetterCi(c int32, lowerExpected int32) bool {
	return ((c == lowerExpected) || (c == (lowerExpected - 32)))
}

func Hx_Field_loreline_imports_endsWithLorTxt(path string) bool {
	if path == HxStringNull {
		return false
	}

	var n int32
	_ = n
	n = int32(utf8.RuneCountInString(path))
	if n < 8 {
		return false
	}

	var tmp bool
	_ = tmp
	var tmp1 bool
	_ = tmp1
	var tmp2 bool
	_ = tmp2
	var tmp3 bool
	_ = tmp3
	var tmp4 bool
	_ = tmp4
	var tmp5 bool
	_ = tmp5
	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (c).Valid == false {
		_hx_tmp_0 = -1
	} else {
		_hx_tmp_0 = (c).Value
	}

	if (_hx_tmp_0) == 46 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (tmp_c_1).Valid == false {
			_hx_tmp_1 = -1
		} else {
			_hx_tmp_1 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_1
		tmp5 = ((c1 == 108) || (c1 == 76))
	} else {
		tmp5 = false
	}

	if tmp5 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (tmp_c_1).Valid == false {
			_hx_tmp_2 = -1
		} else {
			_hx_tmp_2 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_2
		tmp4 = ((c1 == 111) || (c1 == 79))
	} else {
		tmp4 = false
	}

	if tmp4 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (tmp_c_1).Valid == false {
			_hx_tmp_3 = -1
		} else {
			_hx_tmp_3 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_3
		tmp3 = ((c1 == 114) || (c1 == 82))
	} else {
		tmp3 = false
	}

	if tmp3 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		if (tmp_c_1).Valid == false {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = (tmp_c_1).Value
		}

		tmp2 = ((_hx_tmp_4) == 46)
	} else {
		tmp2 = false
	}

	if tmp2 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		if (tmp_c_1).Valid == false {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_5
		tmp1 = ((c1 == 116) || (c1 == 84))
	} else {
		tmp1 = false
	}

	if tmp1 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
		var _hx_tmp_6 int32
		_ = _hx_tmp_6
		if (tmp_c_1).Valid == false {
			_hx_tmp_6 = -1
		} else {
			_hx_tmp_6 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_6
		tmp = ((c1 == 120) || (c1 == 88))
	} else {
		tmp = false
	}

	if tmp {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		if (tmp_c_1).Valid == false {
			_hx_tmp_7 = -1
		} else {
			_hx_tmp_7 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_7
		return ((c1 == 116) || (c1 == 84))
	} else {
		return false
	}
}

func Hx_Field_loreline_imports_endsWithLor(path string) bool {
	if path == HxStringNull {
		return false
	}

	var n int32
	_ = n
	n = int32(utf8.RuneCountInString(path))
	if n < 4 {
		return false
	}

	var tmp bool
	_ = tmp
	var tmp1 bool
	_ = tmp1
	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (c).Valid == false {
		_hx_tmp_0 = -1
	} else {
		_hx_tmp_0 = (c).Value
	}

	if (_hx_tmp_0) == 46 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (tmp_c_1).Valid == false {
			_hx_tmp_1 = -1
		} else {
			_hx_tmp_1 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_1
		tmp1 = ((c1 == 108) || (c1 == 76))
	} else {
		tmp1 = false
	}

	if tmp1 {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (tmp_c_1).Valid == false {
			_hx_tmp_2 = -1
		} else {
			_hx_tmp_2 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_2
		tmp = ((c1 == 111) || (c1 == 79))
	} else {
		tmp = false
	}

	if tmp {
		var tmp_c_1 HxNullable[int32]
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (tmp_c_1).Valid == false {
			_hx_tmp_3 = -1
		} else {
			_hx_tmp_3 = (tmp_c_1).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_3
		return ((c1 == 114) || (c1 == 82))
	} else {
		return false
	}
}

func Hx_Field_loreline_imports_lorExtension(path string) string {
	var tmp bool
	_ = tmp
	if path == HxStringNull {
		tmp = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(path))
		if n < 8 {
			tmp = false
		} else {
			var tmp1 bool
			_ = tmp1
			var tmp2 bool
			_ = tmp2
			var tmp3 bool
			_ = tmp3
			var tmp4 bool
			_ = tmp4
			var tmp5 bool
			_ = tmp5
			var tmp6 bool
			_ = tmp6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_1
				tmp6 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp6 = false
			}

			if tmp6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				tmp5 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp5 = false
			}

			if tmp5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				tmp4 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp4 = false
			}

			if tmp4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				tmp3 = ((_hx_tmp_4) == 46)
			} else {
				tmp3 = false
			}

			if tmp3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_5
				tmp2 = ((c1 == 116) || (c1 == 84))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				tmp1 = ((c1 == 120) || (c1 == 88))
			} else {
				tmp1 = false
			}

			if tmp1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				tmp = ((c1 == 116) || (c1 == 84))
			} else {
				tmp = false
			}
		}
	}

	if tmp {
		return ".lor.txt"
	} else {
		return ".lor"
	}
}

func Hx_Field_loreline_imports_isLorFilePath(path string) bool {
	var tmp bool
	_ = tmp
	if path == HxStringNull {
		tmp = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(path))
		if n < 8 {
			tmp = false
		} else {
			var tmp1 bool
			_ = tmp1
			var tmp2 bool
			_ = tmp2
			var tmp3 bool
			_ = tmp3
			var tmp4 bool
			_ = tmp4
			var tmp5 bool
			_ = tmp5
			var tmp6 bool
			_ = tmp6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_1
				tmp6 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp6 = false
			}

			if tmp6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				tmp5 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp5 = false
			}

			if tmp5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				tmp4 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp4 = false
			}

			if tmp4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				tmp3 = ((_hx_tmp_4) == 46)
			} else {
				tmp3 = false
			}

			if tmp3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_5
				tmp2 = ((c1 == 116) || (c1 == 84))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				tmp1 = ((c1 == 120) || (c1 == 88))
			} else {
				tmp1 = false
			}

			if tmp1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				tmp = ((c1 == 116) || (c1 == 84))
			} else {
				tmp = false
			}
		}
	}

	if !tmp {
		if path == HxStringNull {
			return false
		} else {
			var n int32
			_ = n
			n = int32(utf8.RuneCountInString(path))
			if n < 4 {
				return false
			} else {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				var tmp1 bool
				_ = tmp1
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				if (c).Valid == false {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = (c).Value
				}

				if (_hx_tmp_8) == 46 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					if (tmp_c_1).Valid == false {
						_hx_tmp_9 = -1
					} else {
						_hx_tmp_9 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_9
					tmp1 = ((c1 == 108) || (c1 == 76))
				} else {
					tmp1 = false
				}

				if tmp1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					if (tmp_c_1).Valid == false {
						_hx_tmp_10 = -1
					} else {
						_hx_tmp_10 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_10
					tmp_tmp_1 = ((c1 == 111) || (c1 == 79))
				} else {
					tmp_tmp_1 = false
				}

				if tmp_tmp_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					if (tmp_c_1).Valid == false {
						_hx_tmp_11 = -1
					} else {
						_hx_tmp_11 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_11
					return ((c1 == 114) || (c1 == 82))
				} else {
					return false
				}
			}
		}
	} else {
		return true
	}
}

func Hx_Field_loreline_imports_stripLorExtension(path string) string {
	if path == HxStringNull {
		return HxStringNull
	}

	var tmp bool
	_ = tmp
	if path == HxStringNull {
		tmp = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(path))
		if n < 8 {
			tmp = false
		} else {
			var tmp1 bool
			_ = tmp1
			var tmp2 bool
			_ = tmp2
			var tmp3 bool
			_ = tmp3
			var tmp4 bool
			_ = tmp4
			var tmp5 bool
			_ = tmp5
			var tmp6 bool
			_ = tmp6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_1
				tmp6 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp6 = false
			}

			if tmp6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				tmp5 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp5 = false
			}

			if tmp5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				tmp4 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp4 = false
			}

			if tmp4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				tmp3 = ((_hx_tmp_4) == 46)
			} else {
				tmp3 = false
			}

			if tmp3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_5
				tmp2 = ((c1 == 116) || (c1 == 84))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				tmp1 = ((c1 == 120) || (c1 == 88))
			} else {
				tmp1 = false
			}

			if tmp1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				tmp = ((c1 == 116) || (c1 == 84))
			} else {
				tmp = false
			}
		}
	}

	if tmp {
		var _hx_reserved_len HxNullable[int32]
		_ = _hx_reserved_len
		_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 8), Valid: true}

		if (_hx_reserved_len).Valid != false {
			return (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
		} else {
			var _hx_tmp_8 string
			_ = _hx_tmp_8
			_hx_tmp_8 = path
			return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_8, 0, HxNullable[int32]{}))
		}
	}

	var tmp1 bool
	_ = tmp1
	if path == HxStringNull {
		tmp1 = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(path))
		if n < 4 {
			tmp1 = false
		} else {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			var tmp2 bool
			_ = tmp2
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
			var _hx_tmp_9 int32
			_ = _hx_tmp_9
			if (c).Valid == false {
				_hx_tmp_9 = -1
			} else {
				_hx_tmp_9 = (c).Value
			}

			if (_hx_tmp_9) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				if (tmp_c_1).Valid == false {
					_hx_tmp_10 = -1
				} else {
					_hx_tmp_10 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_10
				tmp2 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				if (tmp_c_1).Valid == false {
					_hx_tmp_11 = -1
				} else {
					_hx_tmp_11 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_11
				tmp_tmp_1 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp_tmp_1 = false
			}

			if tmp_tmp_1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				if (tmp_c_1).Valid == false {
					_hx_tmp_12 = -1
				} else {
					_hx_tmp_12 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_12
				tmp1 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp1 = false
			}
		}
	}

	if tmp1 {
		var _hx_reserved_len HxNullable[int32]
		_ = _hx_reserved_len
		_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 4), Valid: true}

		if (_hx_reserved_len).Valid != false {
			return (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
		} else {
			var _hx_tmp_13 string
			_ = _hx_tmp_13
			_hx_tmp_13 = path
			return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_13, 0, HxNullable[int32]{}))
		}
	}

	return path
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_lexererror]()
	hxRegisterArrayConstructor[string]()
}
