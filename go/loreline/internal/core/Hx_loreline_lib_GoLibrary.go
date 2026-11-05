package core

import "errors"
import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_golibrary() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_golibrary() {
}

type Hx_Obj_VTable_loreline_lib_golibrary interface {
	Hx_Field_parse(source string, path string, readFile func(_hx_carg0 string) string) struct {
		Err    error
		Script *Hx_Obj_loreline_lib_goscript
	}
	Hx_Field_scriptFromJson(json string) struct {
		Err    error
		Script *Hx_Obj_loreline_lib_goscript
	}
	Hx_Field_nodeFromJson(json string) struct {
		Err  error
		Node *Hx_Obj_loreline_lib_gonode
	}
	Hx_Field_extractTranslations(script *Hx_Obj_loreline_lib_goscript) *Hx_Obj_loreline_lib_gotranslations
	Hx_Field_loadLocale(locale string, script *Hx_Obj_loreline_lib_goscript, path string, readFile func(_hx_carg0 string) string) struct {
		Err          error
		Translations *Hx_Obj_loreline_lib_gotranslations
	}
	Hx_Field_translationFormat(name string, enabled bool)
	Hx_Field_update(delta float64)
	Hx_Field_errorMessage(value any) string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_golibrary struct {
	VTable Hx_Obj_VTable_loreline_lib_golibrary `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_lib_golibrary_CreateEmptyInstance() *Hx_Obj_loreline_lib_golibrary {
	obj := &Hx_Obj_loreline_lib_golibrary{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_lib_golibrary_CreateInstance() *Hx_Obj_loreline_lib_golibrary {
	obj := Hx_Obj_loreline_lib_golibrary_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_New() {
	{
	}

}

var Hx_Obj_loreline_lib_golibrary_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoLibrary",
	HxMakeArray[string]("main", "message", "toError", "emptyToNull", "fileHandler", "toHost"),
	HxMakeArray[string]("parse", "scriptFromJson", "nodeFromJson", "extractTranslations", "loadLocale", "translationFormat", "update", "errorMessage"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_golibrary_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lib_golibrary_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_golibrary_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "main":
			return Hx_Field_loreline_lib_golibrary_main
		case "message":
			return Hx_Field_loreline_lib_golibrary_message
		case "toError":
			return Hx_Field_loreline_lib_golibrary_toError
		case "emptyToNull":
			return Hx_Field_loreline_lib_golibrary_emptyToNull
		case "fileHandler":
			return Hx_Field_loreline_lib_golibrary_fileHandler
		case "toHost":
			return Hx_Field_loreline_lib_golibrary_toHost
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_golibrary_RTTI
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_As_Hx_Obj_loreline_lib_golibrary() *Hx_Obj_loreline_lib_golibrary {
	return this
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_parse(source string, path string, readFile func(_hx_carg0 string) string) struct {
	Err    error
	Script *Hx_Obj_loreline_lib_goscript
} {
	{
		var hx_try_return struct {
			Err    error
			Script *Hx_Obj_loreline_lib_goscript
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_6 error
							_ = _hx_tmp_6
							_hx_tmp_6 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_379 any
							_ = hx_dyncast_379
							hx_dyncast_379 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_goscript")
							var _hx_tmp_7 *Hx_Obj_loreline_lib_goscript
							_ = _hx_tmp_7
							if hx_dyncast_379 != nil {
								_hx_tmp_7 = hx_dyncast_379.(interface {
									Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript
								}).Hx_As_Hx_Obj_loreline_lib_goscript()
							} else {
								_hx_tmp_7 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err    error
								Script *Hx_Obj_loreline_lib_goscript
							}{_hx_tmp_6, _hx_tmp_7}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = source
				var _hx_tmp_1 HxNullable[string]
				_ = _hx_tmp_1
				_hx_tmp_1 = HxNullableString(Hx_Field_loreline_lib_golibrary_emptyToNull(path))
				var _hx_tmp_2 Hx_Typedef_loreline_importsfilehandler
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_loreline_lib_golibrary_fileHandler(readFile)
				var script *Hx_Obj_loreline_script
				_ = script
				script = Hx_Field_loreline_loreline_parse(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, nil)
				if script == nil {
					var e *Hx_Obj_loreline_error
					_ = e
					e = Hx_Field_loreline_loreline_lastError()
					var _hx_tmp_3 string
					_ = _hx_tmp_3
					if e != nil {
						_hx_tmp_3 = Hx_Field_std_string(e)
					} else {
						_hx_tmp_3 = "Failed to parse script"
					}

					var hx_dyncast_377 any
					_ = hx_dyncast_377
					hx_dyncast_377 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_goscript")
					var _hx_tmp_4 *Hx_Obj_loreline_lib_goscript
					_ = _hx_tmp_4
					if hx_dyncast_377 != nil {
						_hx_tmp_4 = hx_dyncast_377.(interface {
							Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript
						}).Hx_As_Hx_Obj_loreline_lib_goscript()
					} else {
						_hx_tmp_4 = nil
					}

					hx_try_state = 1
					hx_try_return = struct {
						Err    error
						Script *Hx_Obj_loreline_lib_goscript
					}{func() error {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(errors.New(_hx_tmp_3))))
						if _hx_d == nil {
							var _hx_z error
							return _hx_z
						}
						return _hx_d.(error)
					}(), _hx_tmp_4}
					return
				}

				var hx_dyncast_378 any
				_ = hx_dyncast_378
				hx_dyncast_378 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Obj_loreline_lib_goscript_CreateInstance(script))), "Hx_Obj_loreline_lib_goscript")
				var _hx_tmp_5 *Hx_Obj_loreline_lib_goscript
				_ = _hx_tmp_5
				if hx_dyncast_378 != nil {
					_hx_tmp_5 = hx_dyncast_378.(interface {
						Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript
					}).Hx_As_Hx_Obj_loreline_lib_goscript()
				} else {
					_hx_tmp_5 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err    error
					Script *Hx_Obj_loreline_lib_goscript
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_5}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_golibrary) Parse(source string, path string, readFile func(_hx_carg0 string) string) (*Hx_Obj_loreline_lib_goscript, error) {
	_v := this.VTable.Hx_Field_parse(source, path, readFile)
	return _v.Script, _v.Err
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_scriptFromJson(json string) struct {
	Err    error
	Script *Hx_Obj_loreline_lib_goscript
} {
	{
		var hx_try_return struct {
			Err    error
			Script *Hx_Obj_loreline_lib_goscript
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_1 error
							_ = _hx_tmp_1
							_hx_tmp_1 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_381 any
							_ = hx_dyncast_381
							hx_dyncast_381 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_goscript")
							var _hx_tmp_2 *Hx_Obj_loreline_lib_goscript
							_ = _hx_tmp_2
							if hx_dyncast_381 != nil {
								_hx_tmp_2 = hx_dyncast_381.(interface {
									Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript
								}).Hx_As_Hx_Obj_loreline_lib_goscript()
							} else {
								_hx_tmp_2 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err    error
								Script *Hx_Obj_loreline_lib_goscript
							}{_hx_tmp_1, _hx_tmp_2}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_dyncast_380 any
				_ = hx_dyncast_380
				hx_dyncast_380 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Obj_loreline_lib_goscript_CreateInstance(Hx_Field_loreline_script_fromJson(Hx_Field_loreline_json_parse(json))))), "Hx_Obj_loreline_lib_goscript")
				var _hx_tmp_0 *Hx_Obj_loreline_lib_goscript
				_ = _hx_tmp_0
				if hx_dyncast_380 != nil {
					_hx_tmp_0 = hx_dyncast_380.(interface {
						Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript
					}).Hx_As_Hx_Obj_loreline_lib_goscript()
				} else {
					_hx_tmp_0 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err    error
					Script *Hx_Obj_loreline_lib_goscript
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_0}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_golibrary) ScriptFromJson(json string) (*Hx_Obj_loreline_lib_goscript, error) {
	_v := this.VTable.Hx_Field_scriptFromJson(json)
	return _v.Script, _v.Err
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_nodeFromJson(json string) struct {
	Err  error
	Node *Hx_Obj_loreline_lib_gonode
} {
	{
		var hx_try_return struct {
			Err  error
			Node *Hx_Obj_loreline_lib_gonode
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_1 error
							_ = _hx_tmp_1
							_hx_tmp_1 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_383 any
							_ = hx_dyncast_383
							hx_dyncast_383 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gonode")
							var _hx_tmp_2 *Hx_Obj_loreline_lib_gonode
							_ = _hx_tmp_2
							if hx_dyncast_383 != nil {
								_hx_tmp_2 = hx_dyncast_383.(interface {
									Hx_As_Hx_Obj_loreline_lib_gonode() *Hx_Obj_loreline_lib_gonode
								}).Hx_As_Hx_Obj_loreline_lib_gonode()
							} else {
								_hx_tmp_2 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err  error
								Node *Hx_Obj_loreline_lib_gonode
							}{_hx_tmp_1, _hx_tmp_2}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_dyncast_382 any
				_ = hx_dyncast_382
				hx_dyncast_382 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Obj_loreline_lib_gonode_CreateInstance(Hx_Field_loreline_node_fromJson(Hx_Field_loreline_json_parse(json))))), "Hx_Obj_loreline_lib_gonode")
				var _hx_tmp_0 *Hx_Obj_loreline_lib_gonode
				_ = _hx_tmp_0
				if hx_dyncast_382 != nil {
					_hx_tmp_0 = hx_dyncast_382.(interface {
						Hx_As_Hx_Obj_loreline_lib_gonode() *Hx_Obj_loreline_lib_gonode
					}).Hx_As_Hx_Obj_loreline_lib_gonode()
				} else {
					_hx_tmp_0 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err  error
					Node *Hx_Obj_loreline_lib_gonode
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_0}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_golibrary) NodeFromJson(json string) (*Hx_Obj_loreline_lib_gonode, error) {
	_v := this.VTable.Hx_Field_nodeFromJson(json)
	return _v.Node, _v.Err
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_extractTranslations(script *Hx_Obj_loreline_lib_goscript) *Hx_Obj_loreline_lib_gotranslations {
	return Hx_Obj_loreline_lib_gotranslations_CreateInstance(Hx_Field_loreline_loreline_extractTranslations(script.Hx_Field_script))
}

func (this *Hx_Obj_loreline_lib_golibrary) ExtractTranslations(script *Hx_Obj_loreline_lib_goscript) *Hx_Obj_loreline_lib_gotranslations {
	return this.VTable.Hx_Field_extractTranslations(script)
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_loadLocale(locale string, script *Hx_Obj_loreline_lib_goscript, path string, readFile func(_hx_carg0 string) string) struct {
	Err          error
	Translations *Hx_Obj_loreline_lib_gotranslations
} {
	{
		var hx_try_return struct {
			Err          error
			Translations *Hx_Obj_loreline_lib_gotranslations
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_7 error
							_ = _hx_tmp_7
							_hx_tmp_7 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_386 any
							_ = hx_dyncast_386
							hx_dyncast_386 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gotranslations")
							var _hx_tmp_8 *Hx_Obj_loreline_lib_gotranslations
							_ = _hx_tmp_8
							if hx_dyncast_386 != nil {
								_hx_tmp_8 = hx_dyncast_386.(interface {
									Hx_As_Hx_Obj_loreline_lib_gotranslations() *Hx_Obj_loreline_lib_gotranslations
								}).Hx_As_Hx_Obj_loreline_lib_gotranslations()
							} else {
								_hx_tmp_8 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err          error
								Translations *Hx_Obj_loreline_lib_gotranslations
							}{_hx_tmp_7, _hx_tmp_8}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = locale
				var _hx_tmp_1 *Hx_Obj_loreline_script
				_ = _hx_tmp_1
				_hx_tmp_1 = script.Hx_Field_script
				var _hx_tmp_2 HxNullable[string]
				_ = _hx_tmp_2
				_hx_tmp_2 = HxNullableString(Hx_Field_loreline_lib_golibrary_emptyToNull(path))
				var _hx_tmp_3 Hx_Typedef_loreline_importsfilehandler
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_lib_golibrary_fileHandler(readFile)
				var translations Hx_Typedef_map
				_ = translations
				translations = Hx_Field_loreline_loreline_loadLocale(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, nil)
				if translations == nil {
					var e *Hx_Obj_loreline_error
					_ = e
					e = Hx_Field_loreline_loreline_lastError()
					var _hx_tmp_4 string
					_ = _hx_tmp_4
					if e != nil {
						_hx_tmp_4 = Hx_Field_std_string(e)
					} else {
						_hx_tmp_4 = Hx_Field_go_haxe_hxstring_concat("Failed to load locale ", locale)
					}

					var hx_dyncast_384 any
					_ = hx_dyncast_384
					hx_dyncast_384 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gotranslations")
					var _hx_tmp_5 *Hx_Obj_loreline_lib_gotranslations
					_ = _hx_tmp_5
					if hx_dyncast_384 != nil {
						_hx_tmp_5 = hx_dyncast_384.(interface {
							Hx_As_Hx_Obj_loreline_lib_gotranslations() *Hx_Obj_loreline_lib_gotranslations
						}).Hx_As_Hx_Obj_loreline_lib_gotranslations()
					} else {
						_hx_tmp_5 = nil
					}

					hx_try_state = 1
					hx_try_return = struct {
						Err          error
						Translations *Hx_Obj_loreline_lib_gotranslations
					}{func() error {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(errors.New(_hx_tmp_4))))
						if _hx_d == nil {
							var _hx_z error
							return _hx_z
						}
						return _hx_d.(error)
					}(), _hx_tmp_5}
					return
				}

				var hx_dyncast_385 any
				_ = hx_dyncast_385
				hx_dyncast_385 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Obj_loreline_lib_gotranslations_CreateInstance(translations))), "Hx_Obj_loreline_lib_gotranslations")
				var _hx_tmp_6 *Hx_Obj_loreline_lib_gotranslations
				_ = _hx_tmp_6
				if hx_dyncast_385 != nil {
					_hx_tmp_6 = hx_dyncast_385.(interface {
						Hx_As_Hx_Obj_loreline_lib_gotranslations() *Hx_Obj_loreline_lib_gotranslations
					}).Hx_As_Hx_Obj_loreline_lib_gotranslations()
				} else {
					_hx_tmp_6 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err          error
					Translations *Hx_Obj_loreline_lib_gotranslations
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_6}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_golibrary) LoadLocale(locale string, script *Hx_Obj_loreline_lib_goscript, path string, readFile func(_hx_carg0 string) string) (*Hx_Obj_loreline_lib_gotranslations, error) {
	_v := this.VTable.Hx_Field_loadLocale(locale, script, path, readFile)
	return _v.Translations, _v.Err
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_translationFormat(name string, enabled bool) {
	Hx_Field_loreline_loreline_translationFormat(name, enabled)
}

func (this *Hx_Obj_loreline_lib_golibrary) TranslationFormat(name string, enabled bool) {
	this.VTable.Hx_Field_translationFormat(name, enabled)
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_update(delta float64) {
	Hx_Field_loreline_timer_update(delta)
}

func (this *Hx_Obj_loreline_lib_golibrary) Update(delta float64) {
	this.VTable.Hx_Field_update(delta)
}

func (this *Hx_Obj_loreline_lib_golibrary) Hx_Field_errorMessage(value any) string {
	return Hx_Field_loreline_lib_golibrary_message(value)
}

func (this *Hx_Obj_loreline_lib_golibrary) ErrorMessage(value any) string {
	return this.VTable.Hx_Field_errorMessage(value)
}

func Hx_Field_loreline_lib_golibrary_main() {
}

func Hx_Field_loreline_lib_golibrary_message(e any) string {
	if Hx_Field_std_isOfType(e, Hx_Obj_haxe_valueexception_RTTI) {
		var hx_dyncast_387 any
		_ = hx_dyncast_387
		hx_dyncast_387 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_haxe_valueexception")
		var _hx_tmp_0 *Hx_Obj_haxe_valueexception
		_ = _hx_tmp_0
		if hx_dyncast_387 != nil {
			_hx_tmp_0 = hx_dyncast_387.(interface {
				Hx_As_Hx_Obj_haxe_valueexception() *Hx_Obj_haxe_valueexception
			}).Hx_As_Hx_Obj_haxe_valueexception()
		} else {
			_hx_tmp_0 = nil
		}

		var tmp any
		_ = tmp
		tmp = _hx_tmp_0.Hx_Field_value
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(tmp, nil))) {
			return "null"
		} else {
			return Hx_Field_std_string(tmp)
		}
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		return "null"
	} else {
		return Hx_Field_std_string(e)
	}
}

func Hx_Field_loreline_lib_golibrary_toError(e any) error {
	return errors.New(Hx_Field_loreline_lib_golibrary_message(e))
}

func Hx_Field_loreline_lib_golibrary_emptyToNull(value string) string {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if value == HxStringNull {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (int32(utf8.RuneCountInString(value)) == 0)
	}

	if _hx_tmp_0 {
		return HxStringNull
	} else {
		return value
	}
}

func Hx_Field_loreline_lib_golibrary_fileHandler(readFile func(_hx_carg0 string) string) Hx_Typedef_loreline_importsfilehandler {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(readFile)), ((any)(nil))))) {
		return nil
	} else {
		return ((Hx_Typedef_loreline_importsfilehandler)(func(path string, callback func(data string)) {
			callback(readFile(path))
		}))
	}
}

func Hx_Field_loreline_lib_golibrary_toHost(interpreter *Hx_Obj_loreline_interpreter, value any, depth__isOptional any) any {
	depth := HxDefault[int32]()
	_ = depth

	if depth__isOptional != nil {
		depth = HxConvert[int32](depth__isOptional)
	} else {
		depth = 0
	}
	{
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
			_hx_tmp_4 = true
		} else {
			_hx_tmp_4 = Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI)
		}

		if _hx_tmp_4 {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
		}

		if _hx_tmp_3 {
			_hx_tmp_2 = true
		} else {
			_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (depth > 32)
		}

		if _hx_tmp_0 {
			return value
		}

		if Hx_Field_loreline_arrays_isArray(value) {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 int32
				_ = _g2
				_g2 = Hx_Field_loreline_arrays_arrayLength(value)
				for _g1 < _g2 {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = _g1
					_g1 = (_g1 + 1)
					var i int32
					_ = i
					i = _hx_tmp_5
					var _hx_tmp_6 HxArray[any]
					_ = _hx_tmp_6
					_hx_tmp_6 = _g
					var _hx_tmp_7 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_7
					_hx_tmp_7 = interpreter
					_ = (Hx_Array_Push[any](_hx_tmp_6, Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_7, Hx_Field_loreline_arrays_arrayGet(value, i), (depth+1))))
				}
			}

			return ((any)(_g))
		}

		var _hx_tmp_8 bool
		_ = _hx_tmp_8
		if Hx_Field_std_isOfType(value, Hx_Obj_loreline_fields_RTTI) {
			_hx_tmp_8 = true
		} else {
			_hx_tmp_8 = Hx_Field_std_isOfType(value, Hx_Obj_haxe_ds_stringmap_RTTI)
		}

		if _hx_tmp_8 {
			var result any
			_ = result
			result = any(map[string]any{})
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[string]
				_ = _g1
				_g1 = Hx_Field_loreline_objects_getFields(interpreter, value)
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					if !(_hx_tmp_9 < _g1.Len()) {
						break
					}

					var name string
					_ = name
					name = _g1.Get(_g)
					_g++
					{
						var _hx_tmp_10 *Hx_Obj_loreline_interpreter
						_ = _hx_tmp_10
						_hx_tmp_10 = interpreter
						var value1 any
						_ = value1
						value1 = Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_10, Hx_Field_loreline_objects_getField(interpreter, value, name), (depth + 1))
						Hx_Field_reflect_setField(result, name, value1)
					}
				}
			}

			return result
		}

		return value
	}
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
