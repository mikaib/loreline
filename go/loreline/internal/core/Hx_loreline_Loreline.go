package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_loreline() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_loreline() {
	Hx_Field_loreline_loreline__lastError = Hx_Init_Hx_Field_loreline_loreline__lastError()
}

type Hx_Obj_VTable_loreline_loreline interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_loreline struct {
	VTable Hx_Obj_VTable_loreline_loreline `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_loreline_CreateEmptyInstance() *Hx_Obj_loreline_loreline {
	obj := &Hx_Obj_loreline_loreline{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_loreline_CreateInstance() *Hx_Obj_loreline_loreline {
	obj := Hx_Obj_loreline_loreline_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_loreline_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Loreline",
	HxMakeArray[string]("_lastError", "lastError", "parse", "_reportParseError", "loadLocale", "relativePath", "play", "resume", "extractTranslations", "translationFormat", "generateTranslationFile", "insertLocalizationKeys", "collectHashIds", "removeLocalizationKeys", "extractTranslatableEntries", "hasUntaggedTranslatableStrings", "print", "update"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_loreline_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_loreline_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_loreline_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_lastError":
			return &Hx_Field_loreline_loreline__lastError
		case "lastError":
			return Hx_Field_loreline_loreline_lastError
		case "parse":
			return Hx_Field_loreline_loreline_parse
		case "_reportParseError":
			return Hx_Field_loreline_loreline__reportParseError
		case "loadLocale":
			return Hx_Field_loreline_loreline_loadLocale
		case "relativePath":
			return Hx_Field_loreline_loreline_relativePath
		case "play":
			return Hx_Field_loreline_loreline_play
		case "resume":
			return Hx_Field_loreline_loreline_resume
		case "extractTranslations":
			return Hx_Field_loreline_loreline_extractTranslations
		case "translationFormat":
			return Hx_Field_loreline_loreline_translationFormat
		case "generateTranslationFile":
			return Hx_Field_loreline_loreline_generateTranslationFile
		case "insertLocalizationKeys":
			return Hx_Field_loreline_loreline_insertLocalizationKeys
		case "collectHashIds":
			return Hx_Field_loreline_loreline_collectHashIds
		case "removeLocalizationKeys":
			return Hx_Field_loreline_loreline_removeLocalizationKeys
		case "extractTranslatableEntries":
			return Hx_Field_loreline_loreline_extractTranslatableEntries
		case "hasUntaggedTranslatableStrings":
			return Hx_Field_loreline_loreline_hasUntaggedTranslatableStrings
		case "print":
			return Hx_Field_loreline_loreline_print
		case "update":
			return Hx_Field_loreline_loreline_update
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_loreline) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_loreline_RTTI
}

func (this *Hx_Obj_loreline_loreline) Hx_As_Hx_Obj_loreline_loreline() *Hx_Obj_loreline_loreline {
	return this
}

var Hx_Field_loreline_loreline__lastError *Hx_Obj_loreline_error

func Hx_Init_Hx_Field_loreline_loreline__lastError() *Hx_Obj_loreline_error {
	return nil
}

func Hx_Field_loreline_loreline_lastError() *Hx_Obj_loreline_error {
	return Hx_Field_loreline_loreline__lastError
}

func Hx_Field_loreline_loreline_parse(input string, filePath__isOptional any, handleFile__isOptional any, callback__isOptional any) *Hx_Obj_loreline_script {
	filePath := HxDefault[HxNullable[string]]()
	_ = filePath

	if filePath__isOptional != nil {
		filePath = HxConvert[HxNullable[string]](filePath__isOptional)
	} else {

	}
	handleFile := HxDefault[Hx_Typedef_loreline_importsfilehandler]()
	_ = handleFile

	if handleFile__isOptional != nil {
		handleFile = HxConvert[Hx_Typedef_loreline_importsfilehandler](handleFile__isOptional)
	} else {
		handleFile = nil
	}
	callback := HxDefault[func(script *Hx_Obj_loreline_script)]()
	_ = callback

	if callback__isOptional != nil {
		callback = HxConvert[func(script *Hx_Obj_loreline_script)](callback__isOptional)
	} else {
		callback = nil
	}
	{
		if (filePath).Valid == false {
			var hx_nullable_t74 HxNullable[string]
			_ = hx_nullable_t74
			hx_nullable_t74 = HxNullable[string]{Value: HxStringNull}

			filePath = HxNullable[string]{Value: (hx_nullable_t74).Value, Valid: (hx_nullable_t74).Valid}
		}

		Hx_Field_loreline_loreline__lastError = nil
		var lexer *Hx_Obj_loreline_lexer
		_ = lexer
		lexer = Hx_Obj_loreline_lexer_CreateInstance(input)
		var tokens Hx_Typedef_loreline_tokens
		_ = tokens
		tokens = lexer.VTable.Hx_Field_tokenize()
		var lexerErrors HxArray[*Hx_Obj_loreline_lexererror]
		_ = lexerErrors
		lexerErrors = lexer.VTable.Hx_Field_getErrors()
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if lexerErrors != nil {
			_hx_tmp_0 = (lexerErrors.Len() > 0)
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			return Hx_Field_loreline_loreline__reportParseError((func() *Hx_Obj_loreline_error {
				_hx_up := lexerErrors.Get(0)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_error
				}
				return nil
			})(), callback)
		}

		var result *Hx_Obj_loreline_script
		_ = result
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (filePath).Valid != false {
			_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(handleFile)), ((any)(nil))))))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var imports *Hx_Obj_loreline_imports
			_ = imports
			imports = Hx_Obj_loreline_imports_CreateInstance()
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = (filePath).Value
			var _hx_tmp_3 Hx_Typedef_loreline_tokens
			_ = _hx_tmp_3
			_hx_tmp_3 = tokens
			var _hx_tmp_4 Hx_Typedef_loreline_importsfilehandler
			_ = _hx_tmp_4
			_hx_tmp_4 = handleFile
			var _hx_tmp_5 Hx_Typedef_loreline_importserrorhandler
			_ = _hx_tmp_5
			_hx_tmp_5 = ((Hx_Typedef_loreline_importserrorhandler)(func(_hx_reserved_error *Hx_Obj_loreline_error) {
				Hx_Field_loreline_loreline__reportParseError(_hx_reserved_error, callback)
			}))
			imports.VTable.Hx_Field_resolve(_hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, ((Hx_Typedef_loreline_importscallback)(func(hasErrors bool, resolvedImports Hx_Typedef_map) {
				if Hx_Field_loreline_loreline__lastError != nil {
					return
				}

				var parser *Hx_Obj_loreline_parser
				_ = parser
				parser = Hx_Obj_loreline_parser_CreateInstance(tokens, Hx_Obj_loreline_parsercontext_CreateInstance((filePath).Value, (filePath).Value, resolvedImports, nil))
				result = parser.VTable.Hx_Field_parse()
				result.Hx_Field_indentSize = lexer.Hx_Field_detectedIndentSize
				var parseErrors HxArray[*Hx_Obj_loreline_parseerror]
				_ = parseErrors
				parseErrors = parser.VTable.Hx_Field_getErrors()
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if parseErrors != nil {
					_hx_tmp_6 = (parseErrors.Len() > 0)
				} else {
					_hx_tmp_6 = false
				}

				if _hx_tmp_6 {
					result = nil
					Hx_Field_loreline_loreline__reportParseError((func() *Hx_Obj_loreline_error {
						_hx_up := parseErrors.Get(0)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_error
						}
						return nil
					})(), callback)
					return
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(callback)), ((any)(nil))))) {
					callback(result)
				}
			})))
			return result
		}

		var parser *Hx_Obj_loreline_parser
		_ = parser
		parser = Hx_Obj_loreline_parser_CreateInstance(tokens, nil)
		result = parser.VTable.Hx_Field_parse()
		result.Hx_Field_indentSize = lexer.Hx_Field_detectedIndentSize
		var parseErrors HxArray[*Hx_Obj_loreline_parseerror]
		_ = parseErrors
		parseErrors = parser.VTable.Hx_Field_getErrors()
		var _hx_tmp_7 bool
		_ = _hx_tmp_7
		if parseErrors != nil {
			_hx_tmp_7 = (parseErrors.Len() > 0)
		} else {
			_hx_tmp_7 = false
		}

		if _hx_tmp_7 {
			return Hx_Field_loreline_loreline__reportParseError((func() *Hx_Obj_loreline_error {
				_hx_up := parseErrors.Get(0)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_error
				}
				return nil
			})(), callback)
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(callback)), ((any)(nil))))) {
			callback(result)
		}

		return result
	}
}

func Hx_Field_loreline_loreline__reportParseError(e *Hx_Obj_loreline_error, callback__isOptional any) *Hx_Obj_loreline_script {
	callback := HxDefault[func(script *Hx_Obj_loreline_script)]()
	_ = callback

	if callback__isOptional != nil {
		callback = HxConvert[func(script *Hx_Obj_loreline_script)](callback__isOptional)
	} else {
		callback = nil
	}
	{
		Hx_Field_loreline_loreline__lastError = e
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(callback)), ((any)(nil))))) {
			callback(nil)
			return nil
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	}
}

func Hx_Field_loreline_loreline_loadLocale(locale string, script *Hx_Obj_loreline_script, filePath__isOptional any, handleFile__isOptional any, callback__isOptional any) Hx_Typedef_map {
	filePath := HxDefault[HxNullable[string]]()
	_ = filePath

	if filePath__isOptional != nil {
		filePath = HxConvert[HxNullable[string]](filePath__isOptional)
	} else {

	}
	handleFile := HxDefault[Hx_Typedef_loreline_importsfilehandler]()
	_ = handleFile

	if handleFile__isOptional != nil {
		handleFile = HxConvert[Hx_Typedef_loreline_importsfilehandler](handleFile__isOptional)
	} else {
		handleFile = nil
	}
	callback := HxDefault[func(translations Hx_Typedef_map)]()
	_ = callback

	if callback__isOptional != nil {
		callback = HxConvert[func(translations Hx_Typedef_map)](callback__isOptional)
	} else {
		callback = nil
	}
	{
		if (filePath).Valid == false {
			var hx_nullable_t75 HxNullable[string]
			_ = hx_nullable_t75
			hx_nullable_t75 = HxNullable[string]{Value: HxStringNull}

			filePath = HxNullable[string]{Value: (hx_nullable_t75).Value, Valid: (hx_nullable_t75).Valid}
		}

		Hx_Field_loreline_loreline__lastError = nil
		if (filePath).Valid == false {
			filePath = HxNullableString(script.Hx_Field_filePath)
		}

		if (filePath).Valid == false {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Cannot load locale: no filePath provided and script.filePath is null", nil)))))
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(handleFile)), ((any)(nil))))) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Cannot load locale: handleFile is required", nil)))))
		}

		var wrapped *Hx_Obj_loreline_translation_wrappedfilehandler
		_ = wrapped
		wrapped = Hx_Field_loreline_translation_translationformats_wrap(handleFile, locale)
		handleFile = wrapped.Hx_Field_handler
		var wrappedLastError func() *Hx_Obj_loreline_error
		_ = wrappedLastError
		wrappedLastError = wrapped.Hx_Field_lastError
		var ext string
		_ = ext
		ext = ".lor"
		var translationsBaseDir string
		_ = translationsBaseDir
		var translationsBaseDir1 bool
		_ = translationsBaseDir1
		var translationsBaseDir2 bool
		_ = translationsBaseDir2
		if (filePath).Valid == false {
			translationsBaseDir2 = false
		} else {
			var n int32
			_ = n
			n = int32(utf8.RuneCountInString(filePath.Value))
			if n < 8 {
				translationsBaseDir2 = false
			} else {
				var tmp_translationsBaseDir_1 bool
				_ = tmp_translationsBaseDir_1
				var tmp_translationsBaseDir1_1 bool
				_ = tmp_translationsBaseDir1_1
				var translationsBaseDir3 bool
				_ = translationsBaseDir3
				var translationsBaseDir4 bool
				_ = translationsBaseDir4
				var translationsBaseDir5 bool
				_ = translationsBaseDir5
				var translationsBaseDir6 bool
				_ = translationsBaseDir6
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 8)))
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
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 7)))
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
					translationsBaseDir6 = ((c1 == 108) || (c1 == 76))
				} else {
					translationsBaseDir6 = false
				}

				if translationsBaseDir6 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 6)))
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
					translationsBaseDir5 = ((c1 == 111) || (c1 == 79))
				} else {
					translationsBaseDir5 = false
				}

				if translationsBaseDir5 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 5)))
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
					translationsBaseDir4 = ((c1 == 114) || (c1 == 82))
				} else {
					translationsBaseDir4 = false
				}

				if translationsBaseDir4 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 4)))
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					if (tmp_c_1).Valid == false {
						_hx_tmp_4 = -1
					} else {
						_hx_tmp_4 = (tmp_c_1).Value
					}

					translationsBaseDir3 = ((_hx_tmp_4) == 46)
				} else {
					translationsBaseDir3 = false
				}

				if translationsBaseDir3 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 3)))
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
					tmp_translationsBaseDir1_1 = ((c1 == 116) || (c1 == 84))
				} else {
					tmp_translationsBaseDir1_1 = false
				}

				if tmp_translationsBaseDir1_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 2)))
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
					tmp_translationsBaseDir_1 = ((c1 == 120) || (c1 == 88))
				} else {
					tmp_translationsBaseDir_1 = false
				}

				if tmp_translationsBaseDir_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 1)))
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
					translationsBaseDir2 = ((c1 == 116) || (c1 == 84))
				} else {
					translationsBaseDir2 = false
				}
			}
		}

		if !translationsBaseDir2 {
			if (filePath).Valid == false {
				translationsBaseDir1 = false
			} else {
				var n int32
				_ = n
				n = int32(utf8.RuneCountInString(filePath.Value))
				if n < 4 {
					translationsBaseDir1 = false
				} else {
					var tmp_translationsBaseDir_1 bool
					_ = tmp_translationsBaseDir_1
					var tmp_translationsBaseDir2_1 bool
					_ = tmp_translationsBaseDir2_1
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 4)))
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
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 3)))
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
						tmp_translationsBaseDir2_1 = ((c1 == 108) || (c1 == 76))
					} else {
						tmp_translationsBaseDir2_1 = false
					}

					if tmp_translationsBaseDir2_1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 2)))
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
						tmp_translationsBaseDir_1 = ((c1 == 111) || (c1 == 79))
					} else {
						tmp_translationsBaseDir_1 = false
					}

					if tmp_translationsBaseDir_1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (n - 1)))
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
						translationsBaseDir1 = ((c1 == 114) || (c1 == 82))
					} else {
						translationsBaseDir1 = false
					}
				}
			}
		} else {
			translationsBaseDir1 = true
		}

		if translationsBaseDir1 {
			translationsBaseDir = Hx_Field_haxe_io_path_directory((filePath).Value)
		} else {
			var fpLen int32
			_ = fpLen
			fpLen = int32(utf8.RuneCountInString(filePath.Value))
			var lastChar int32
			_ = lastChar
			if fpLen > 0 {
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(filePath.Value, (fpLen - 1)))
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				if (c).Valid == false {
					_hx_tmp_12 = -1
				} else {
					_hx_tmp_12 = (c).Value
				}

				lastChar = _hx_tmp_12
			} else {
				lastChar = 0
			}

			if (lastChar == 47) || (lastChar == 92) {
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (fpLen - 1), Valid: true}

				var _hx_tmp_13 string
				_ = _hx_tmp_13
				if (_hx_reserved_len).Valid != false {
					_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(filePath.Value, 0, _hx_reserved_len))
				} else {
					var _hx_tmp_14 string
					_ = _hx_tmp_14
					_hx_tmp_14 = filePath.Value
					var _hx_tmp_15 string
					_ = _hx_tmp_15
					_hx_tmp_15 = _hx_tmp_14
					_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, 0, HxNullable[int32]{}))
				}

				translationsBaseDir = _hx_tmp_13
			} else {
				translationsBaseDir = (filePath).Value
			}
		}

		var sourceRootPath string
		_ = sourceRootPath
		sourceRootPath = script.Hx_Field_filePath
		var relativePaths HxArray[string]
		_ = relativePaths
		relativePaths = HxMakeArray[string]()
		var translationPaths HxArray[string]
		_ = translationPaths
		translationPaths = HxMakeArray[string]()
		if sourceRootPath != HxStringNull {
			var rootBaseName string
			_ = rootBaseName
			rootBaseName = Hx_Field_haxe_io_path_withoutDirectory(sourceRootPath)
			var rootStem HxNullable[string]
			_ = rootStem
			if rootBaseName == HxStringNull {
				rootStem = HxNullableString(HxStringNull)
			} else {
				var rootStem1 bool
				_ = rootStem1
				if rootBaseName == HxStringNull {
					rootStem1 = false
				} else {
					var n int32
					_ = n
					n = int32(utf8.RuneCountInString(rootBaseName))
					if n < 8 {
						rootStem1 = false
					} else {
						var tmp_rootStem_1 bool
						_ = tmp_rootStem_1
						var rootStem2 bool
						_ = rootStem2
						var rootStem3 bool
						_ = rootStem3
						var rootStem4 bool
						_ = rootStem4
						var rootStem5 bool
						_ = rootStem5
						var rootStem6 bool
						_ = rootStem6
						var c HxNullable[int32]
						_ = c
						c = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 8)))
						var _hx_tmp_16 int32
						_ = _hx_tmp_16
						if (c).Valid == false {
							_hx_tmp_16 = -1
						} else {
							_hx_tmp_16 = (c).Value
						}

						if (_hx_tmp_16) == 46 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 7)))
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							if (tmp_c_1).Valid == false {
								_hx_tmp_17 = -1
							} else {
								_hx_tmp_17 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_17
							rootStem6 = ((c1 == 108) || (c1 == 76))
						} else {
							rootStem6 = false
						}

						if rootStem6 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 6)))
							var _hx_tmp_18 int32
							_ = _hx_tmp_18
							if (tmp_c_1).Valid == false {
								_hx_tmp_18 = -1
							} else {
								_hx_tmp_18 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_18
							rootStem5 = ((c1 == 111) || (c1 == 79))
						} else {
							rootStem5 = false
						}

						if rootStem5 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 5)))
							var _hx_tmp_19 int32
							_ = _hx_tmp_19
							if (tmp_c_1).Valid == false {
								_hx_tmp_19 = -1
							} else {
								_hx_tmp_19 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_19
							rootStem4 = ((c1 == 114) || (c1 == 82))
						} else {
							rootStem4 = false
						}

						if rootStem4 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 4)))
							var _hx_tmp_20 int32
							_ = _hx_tmp_20
							if (tmp_c_1).Valid == false {
								_hx_tmp_20 = -1
							} else {
								_hx_tmp_20 = (tmp_c_1).Value
							}

							rootStem3 = ((_hx_tmp_20) == 46)
						} else {
							rootStem3 = false
						}

						if rootStem3 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 3)))
							var _hx_tmp_21 int32
							_ = _hx_tmp_21
							if (tmp_c_1).Valid == false {
								_hx_tmp_21 = -1
							} else {
								_hx_tmp_21 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_21
							rootStem2 = ((c1 == 116) || (c1 == 84))
						} else {
							rootStem2 = false
						}

						if rootStem2 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 2)))
							var _hx_tmp_22 int32
							_ = _hx_tmp_22
							if (tmp_c_1).Valid == false {
								_hx_tmp_22 = -1
							} else {
								_hx_tmp_22 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_22
							tmp_rootStem_1 = ((c1 == 120) || (c1 == 88))
						} else {
							tmp_rootStem_1 = false
						}

						if tmp_rootStem_1 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 1)))
							var _hx_tmp_23 int32
							_ = _hx_tmp_23
							if (tmp_c_1).Valid == false {
								_hx_tmp_23 = -1
							} else {
								_hx_tmp_23 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_23
							rootStem1 = ((c1 == 116) || (c1 == 84))
						} else {
							rootStem1 = false
						}
					}
				}

				if rootStem1 {
					var _hx_reserved_len HxNullable[int32]
					_ = _hx_reserved_len
					_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(rootBaseName)) - 8), Valid: true}

					var _hx_tmp_24 string
					_ = _hx_tmp_24
					if (_hx_reserved_len).Valid != false {
						_hx_tmp_24 = (Hx_Field_go_haxe_hxstring_substr(rootBaseName, 0, _hx_reserved_len))
					} else {
						var _hx_tmp_25 string
						_ = _hx_tmp_25
						_hx_tmp_25 = rootBaseName
						var _hx_tmp_26 string
						_ = _hx_tmp_26
						_hx_tmp_26 = _hx_tmp_25
						_hx_tmp_24 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_26, 0, HxNullable[int32]{}))
					}

					rootStem = HxNullableString(_hx_tmp_24)
				} else {
					var tmp_rootStem1_1 bool
					_ = tmp_rootStem1_1
					if rootBaseName == HxStringNull {
						tmp_rootStem1_1 = false
					} else {
						var n int32
						_ = n
						n = int32(utf8.RuneCountInString(rootBaseName))
						if n < 4 {
							tmp_rootStem1_1 = false
						} else {
							var tmp_rootStem_1 bool
							_ = tmp_rootStem_1
							var rootStem2 bool
							_ = rootStem2
							var c HxNullable[int32]
							_ = c
							c = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 4)))
							var _hx_tmp_27 int32
							_ = _hx_tmp_27
							if (c).Valid == false {
								_hx_tmp_27 = -1
							} else {
								_hx_tmp_27 = (c).Value
							}

							if (_hx_tmp_27) == 46 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 3)))
								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								if (tmp_c_1).Valid == false {
									_hx_tmp_28 = -1
								} else {
									_hx_tmp_28 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_28
								rootStem2 = ((c1 == 108) || (c1 == 76))
							} else {
								rootStem2 = false
							}

							if rootStem2 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 2)))
								var _hx_tmp_29 int32
								_ = _hx_tmp_29
								if (tmp_c_1).Valid == false {
									_hx_tmp_29 = -1
								} else {
									_hx_tmp_29 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_29
								tmp_rootStem_1 = ((c1 == 111) || (c1 == 79))
							} else {
								tmp_rootStem_1 = false
							}

							if tmp_rootStem_1 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootBaseName, (n - 1)))
								var _hx_tmp_30 int32
								_ = _hx_tmp_30
								if (tmp_c_1).Valid == false {
									_hx_tmp_30 = -1
								} else {
									_hx_tmp_30 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_30
								tmp_rootStem1_1 = ((c1 == 114) || (c1 == 82))
							} else {
								tmp_rootStem1_1 = false
							}
						}
					}

					if tmp_rootStem1_1 {
						var _hx_reserved_len HxNullable[int32]
						_ = _hx_reserved_len
						_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(rootBaseName)) - 4), Valid: true}

						var _hx_tmp_31 string
						_ = _hx_tmp_31
						if (_hx_reserved_len).Valid != false {
							_hx_tmp_31 = (Hx_Field_go_haxe_hxstring_substr(rootBaseName, 0, _hx_reserved_len))
						} else {
							var _hx_tmp_32 string
							_ = _hx_tmp_32
							_hx_tmp_32 = rootBaseName
							var _hx_tmp_33 string
							_ = _hx_tmp_33
							_hx_tmp_33 = _hx_tmp_32
							_hx_tmp_31 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_33, 0, HxNullable[int32]{}))
						}

						rootStem = HxNullableString(_hx_tmp_31)
					} else {
						rootStem = HxNullableString(rootBaseName)
					}
				}
			}

			var rootStem1 string
			_ = rootStem1
			rootStem1 = (rootStem).Value
			var rootTransName string
			_ = rootTransName
			rootTransName = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(rootStem1, "."), locale), ext)
			_ = (Hx_Array_Push[string](relativePaths, "."))
			var _hx_tmp_34 HxArray[string]
			_ = _hx_tmp_34
			_hx_tmp_34 = translationPaths
			_ = (Hx_Array_Push[string](_hx_tmp_34, Hx_Field_haxe_io_path_normalize(Hx_Field_haxe_io_path_join(HxMakeArray[string](translationsBaseDir, rootTransName)))))
		}

		if sourceRootPath != HxStringNull {
			var lens *Hx_Obj_loreline_lens
			_ = lens
			lens = Hx_Obj_loreline_lens_CreateInstance(script)
			var sourceRootDir string
			_ = sourceRootDir
			sourceRootDir = Hx_Field_haxe_io_path_directory(sourceRootPath)
			var importedAbsPaths HxArray[string]
			_ = importedAbsPaths
			importedAbsPaths = lens.VTable.Hx_Field_getImportedPaths(sourceRootPath)
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_35 int32
					_ = _hx_tmp_35
					_hx_tmp_35 = _g
					if !(_hx_tmp_35 < importedAbsPaths.Len()) {
						break
					}

					var importAbsPath string
					_ = importAbsPath
					importAbsPath = importedAbsPaths.Get(_g)
					_g++
					var importRelPath string
					_ = importRelPath
					importRelPath = Hx_Field_loreline_loreline_relativePath(sourceRootDir, importAbsPath)
					var importStem HxNullable[string]
					_ = importStem
					if importRelPath == HxStringNull {
						importStem = HxNullableString(HxStringNull)
					} else {
						var importStem1 bool
						_ = importStem1
						if importRelPath == HxStringNull {
							importStem1 = false
						} else {
							var n int32
							_ = n
							n = int32(utf8.RuneCountInString(importRelPath))
							if n < 8 {
								importStem1 = false
							} else {
								var tmp_importStem_1 bool
								_ = tmp_importStem_1
								var importStem2 bool
								_ = importStem2
								var importStem3 bool
								_ = importStem3
								var importStem4 bool
								_ = importStem4
								var importStem5 bool
								_ = importStem5
								var importStem6 bool
								_ = importStem6
								var c HxNullable[int32]
								_ = c
								c = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 8)))
								var _hx_tmp_36 int32
								_ = _hx_tmp_36
								if (c).Valid == false {
									_hx_tmp_36 = -1
								} else {
									_hx_tmp_36 = (c).Value
								}

								if (_hx_tmp_36) == 46 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 7)))
									var _hx_tmp_37 int32
									_ = _hx_tmp_37
									if (tmp_c_1).Valid == false {
										_hx_tmp_37 = -1
									} else {
										_hx_tmp_37 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_37
									importStem6 = ((c1 == 108) || (c1 == 76))
								} else {
									importStem6 = false
								}

								if importStem6 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 6)))
									var _hx_tmp_38 int32
									_ = _hx_tmp_38
									if (tmp_c_1).Valid == false {
										_hx_tmp_38 = -1
									} else {
										_hx_tmp_38 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_38
									importStem5 = ((c1 == 111) || (c1 == 79))
								} else {
									importStem5 = false
								}

								if importStem5 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 5)))
									var _hx_tmp_39 int32
									_ = _hx_tmp_39
									if (tmp_c_1).Valid == false {
										_hx_tmp_39 = -1
									} else {
										_hx_tmp_39 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_39
									importStem4 = ((c1 == 114) || (c1 == 82))
								} else {
									importStem4 = false
								}

								if importStem4 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 4)))
									var _hx_tmp_40 int32
									_ = _hx_tmp_40
									if (tmp_c_1).Valid == false {
										_hx_tmp_40 = -1
									} else {
										_hx_tmp_40 = (tmp_c_1).Value
									}

									importStem3 = ((_hx_tmp_40) == 46)
								} else {
									importStem3 = false
								}

								if importStem3 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 3)))
									var _hx_tmp_41 int32
									_ = _hx_tmp_41
									if (tmp_c_1).Valid == false {
										_hx_tmp_41 = -1
									} else {
										_hx_tmp_41 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_41
									importStem2 = ((c1 == 116) || (c1 == 84))
								} else {
									importStem2 = false
								}

								if importStem2 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 2)))
									var _hx_tmp_42 int32
									_ = _hx_tmp_42
									if (tmp_c_1).Valid == false {
										_hx_tmp_42 = -1
									} else {
										_hx_tmp_42 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_42
									tmp_importStem_1 = ((c1 == 120) || (c1 == 88))
								} else {
									tmp_importStem_1 = false
								}

								if tmp_importStem_1 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 1)))
									var _hx_tmp_43 int32
									_ = _hx_tmp_43
									if (tmp_c_1).Valid == false {
										_hx_tmp_43 = -1
									} else {
										_hx_tmp_43 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_43
									importStem1 = ((c1 == 116) || (c1 == 84))
								} else {
									importStem1 = false
								}
							}
						}

						if importStem1 {
							var _hx_reserved_len HxNullable[int32]
							_ = _hx_reserved_len
							_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(importRelPath)) - 8), Valid: true}

							var _hx_tmp_44 string
							_ = _hx_tmp_44
							if (_hx_reserved_len).Valid != false {
								_hx_tmp_44 = (Hx_Field_go_haxe_hxstring_substr(importRelPath, 0, _hx_reserved_len))
							} else {
								var _hx_tmp_45 string
								_ = _hx_tmp_45
								_hx_tmp_45 = importRelPath
								var _hx_tmp_46 string
								_ = _hx_tmp_46
								_hx_tmp_46 = _hx_tmp_45
								_hx_tmp_44 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_46, 0, HxNullable[int32]{}))
							}

							importStem = HxNullableString(_hx_tmp_44)
						} else {
							var tmp_importStem1_1 bool
							_ = tmp_importStem1_1
							if importRelPath == HxStringNull {
								tmp_importStem1_1 = false
							} else {
								var n int32
								_ = n
								n = int32(utf8.RuneCountInString(importRelPath))
								if n < 4 {
									tmp_importStem1_1 = false
								} else {
									var tmp_importStem_1 bool
									_ = tmp_importStem_1
									var importStem2 bool
									_ = importStem2
									var c HxNullable[int32]
									_ = c
									c = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 4)))
									var _hx_tmp_47 int32
									_ = _hx_tmp_47
									if (c).Valid == false {
										_hx_tmp_47 = -1
									} else {
										_hx_tmp_47 = (c).Value
									}

									if (_hx_tmp_47) == 46 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 3)))
										var _hx_tmp_48 int32
										_ = _hx_tmp_48
										if (tmp_c_1).Valid == false {
											_hx_tmp_48 = -1
										} else {
											_hx_tmp_48 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_48
										importStem2 = ((c1 == 108) || (c1 == 76))
									} else {
										importStem2 = false
									}

									if importStem2 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 2)))
										var _hx_tmp_49 int32
										_ = _hx_tmp_49
										if (tmp_c_1).Valid == false {
											_hx_tmp_49 = -1
										} else {
											_hx_tmp_49 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_49
										tmp_importStem_1 = ((c1 == 111) || (c1 == 79))
									} else {
										tmp_importStem_1 = false
									}

									if tmp_importStem_1 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importRelPath, (n - 1)))
										var _hx_tmp_50 int32
										_ = _hx_tmp_50
										if (tmp_c_1).Valid == false {
											_hx_tmp_50 = -1
										} else {
											_hx_tmp_50 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_50
										tmp_importStem1_1 = ((c1 == 114) || (c1 == 82))
									} else {
										tmp_importStem1_1 = false
									}
								}
							}

							if tmp_importStem1_1 {
								var _hx_reserved_len HxNullable[int32]
								_ = _hx_reserved_len
								_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(importRelPath)) - 4), Valid: true}

								var _hx_tmp_51 string
								_ = _hx_tmp_51
								if (_hx_reserved_len).Valid != false {
									_hx_tmp_51 = (Hx_Field_go_haxe_hxstring_substr(importRelPath, 0, _hx_reserved_len))
								} else {
									var _hx_tmp_52 string
									_ = _hx_tmp_52
									_hx_tmp_52 = importRelPath
									var _hx_tmp_53 string
									_ = _hx_tmp_53
									_hx_tmp_53 = _hx_tmp_52
									_hx_tmp_51 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_53, 0, HxNullable[int32]{}))
								}

								importStem = HxNullableString(_hx_tmp_51)
							} else {
								importStem = HxNullableString(importRelPath)
							}
						}
					}

					var importStem1 string
					_ = importStem1
					importStem1 = (importStem).Value
					var importTransRelPath string
					_ = importTransRelPath
					importTransRelPath = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(importStem1, "."), locale), ext)
					var path string
					_ = path
					path = Hx_Field_haxe_io_path_normalize(importRelPath)
					var tmp string
					_ = tmp
					if path == HxStringNull {
						tmp = HxStringNull
					} else {
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
								var _hx_tmp_54 int32
								_ = _hx_tmp_54
								if (c).Valid == false {
									_hx_tmp_54 = -1
								} else {
									_hx_tmp_54 = (c).Value
								}

								if (_hx_tmp_54) == 46 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
									var _hx_tmp_55 int32
									_ = _hx_tmp_55
									if (tmp_c_1).Valid == false {
										_hx_tmp_55 = -1
									} else {
										_hx_tmp_55 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_55
									tmp6 = ((c1 == 108) || (c1 == 76))
								} else {
									tmp6 = false
								}

								if tmp6 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
									var _hx_tmp_56 int32
									_ = _hx_tmp_56
									if (tmp_c_1).Valid == false {
										_hx_tmp_56 = -1
									} else {
										_hx_tmp_56 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_56
									tmp5 = ((c1 == 111) || (c1 == 79))
								} else {
									tmp5 = false
								}

								if tmp5 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
									var _hx_tmp_57 int32
									_ = _hx_tmp_57
									if (tmp_c_1).Valid == false {
										_hx_tmp_57 = -1
									} else {
										_hx_tmp_57 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_57
									tmp4 = ((c1 == 114) || (c1 == 82))
								} else {
									tmp4 = false
								}

								if tmp4 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
									var _hx_tmp_58 int32
									_ = _hx_tmp_58
									if (tmp_c_1).Valid == false {
										_hx_tmp_58 = -1
									} else {
										_hx_tmp_58 = (tmp_c_1).Value
									}

									tmp3 = ((_hx_tmp_58) == 46)
								} else {
									tmp3 = false
								}

								if tmp3 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
									var _hx_tmp_59 int32
									_ = _hx_tmp_59
									if (tmp_c_1).Valid == false {
										_hx_tmp_59 = -1
									} else {
										_hx_tmp_59 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_59
									tmp2 = ((c1 == 116) || (c1 == 84))
								} else {
									tmp2 = false
								}

								if tmp2 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
									var _hx_tmp_60 int32
									_ = _hx_tmp_60
									if (tmp_c_1).Valid == false {
										_hx_tmp_60 = -1
									} else {
										_hx_tmp_60 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_60
									tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
								} else {
									tmp_tmp_1 = false
								}

								if tmp_tmp_1 {
									var tmp_c_1 HxNullable[int32]
									_ = tmp_c_1
									tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
									var _hx_tmp_61 int32
									_ = _hx_tmp_61
									if (tmp_c_1).Valid == false {
										_hx_tmp_61 = -1
									} else {
										_hx_tmp_61 = (tmp_c_1).Value
									}

									var c1 int32
									_ = c1
									c1 = _hx_tmp_61
									tmp1 = ((c1 == 116) || (c1 == 84))
								} else {
									tmp1 = false
								}
							}
						}

						if tmp1 {
							var _hx_reserved_len HxNullable[int32]
							_ = _hx_reserved_len
							_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 8), Valid: true}

							var _hx_tmp_62 string
							_ = _hx_tmp_62
							if (_hx_reserved_len).Valid != false {
								_hx_tmp_62 = (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
							} else {
								var _hx_tmp_63 string
								_ = _hx_tmp_63
								_hx_tmp_63 = path
								var _hx_tmp_64 string
								_ = _hx_tmp_64
								_hx_tmp_64 = _hx_tmp_63
								_hx_tmp_62 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_64, 0, HxNullable[int32]{}))
							}

							tmp = _hx_tmp_62
						} else {
							var tmp_tmp1_1 bool
							_ = tmp_tmp1_1
							if path == HxStringNull {
								tmp_tmp1_1 = false
							} else {
								var n int32
								_ = n
								n = int32(utf8.RuneCountInString(path))
								if n < 4 {
									tmp_tmp1_1 = false
								} else {
									var tmp_tmp_1 bool
									_ = tmp_tmp_1
									var tmp2 bool
									_ = tmp2
									var c HxNullable[int32]
									_ = c
									c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
									var _hx_tmp_65 int32
									_ = _hx_tmp_65
									if (c).Valid == false {
										_hx_tmp_65 = -1
									} else {
										_hx_tmp_65 = (c).Value
									}

									if (_hx_tmp_65) == 46 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
										var _hx_tmp_66 int32
										_ = _hx_tmp_66
										if (tmp_c_1).Valid == false {
											_hx_tmp_66 = -1
										} else {
											_hx_tmp_66 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_66
										tmp2 = ((c1 == 108) || (c1 == 76))
									} else {
										tmp2 = false
									}

									if tmp2 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
										var _hx_tmp_67 int32
										_ = _hx_tmp_67
										if (tmp_c_1).Valid == false {
											_hx_tmp_67 = -1
										} else {
											_hx_tmp_67 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_67
										tmp_tmp_1 = ((c1 == 111) || (c1 == 79))
									} else {
										tmp_tmp_1 = false
									}

									if tmp_tmp_1 {
										var tmp_c_1 HxNullable[int32]
										_ = tmp_c_1
										tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
										var _hx_tmp_68 int32
										_ = _hx_tmp_68
										if (tmp_c_1).Valid == false {
											_hx_tmp_68 = -1
										} else {
											_hx_tmp_68 = (tmp_c_1).Value
										}

										var c1 int32
										_ = c1
										c1 = _hx_tmp_68
										tmp_tmp1_1 = ((c1 == 114) || (c1 == 82))
									} else {
										tmp_tmp1_1 = false
									}
								}
							}

							if tmp_tmp1_1 {
								var _hx_reserved_len HxNullable[int32]
								_ = _hx_reserved_len
								_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 4), Valid: true}

								var _hx_tmp_69 string
								_ = _hx_tmp_69
								if (_hx_reserved_len).Valid != false {
									_hx_tmp_69 = (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
								} else {
									var _hx_tmp_70 string
									_ = _hx_tmp_70
									_hx_tmp_70 = path
									var _hx_tmp_71 string
									_ = _hx_tmp_71
									_hx_tmp_71 = _hx_tmp_70
									_hx_tmp_69 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_71, 0, HxNullable[int32]{}))
								}

								tmp = _hx_tmp_69
							} else {
								tmp = path
							}
						}
					}

					_ = (Hx_Array_Push[string](relativePaths, tmp))
					var _hx_tmp_72 HxArray[string]
					_ = _hx_tmp_72
					_hx_tmp_72 = translationPaths
					_ = (Hx_Array_Push[string](_hx_tmp_72, Hx_Field_haxe_io_path_normalize(Hx_Field_haxe_io_path_join(HxMakeArray[string](translationsBaseDir, importTransRelPath)))))
				}
			}
		}

		var result Hx_Typedef_map
		_ = result
		result = (func() *Hx_Obj_haxe_imap {
			_hx_icast_370 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_370 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_370, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_370.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		var pending int32
		_ = pending
		pending = relativePaths.Len()
		var allProcessed bool
		_ = allProcessed
		allProcessed = false
		var firstParseError *Hx_Obj_loreline_error
		_ = firstParseError
		var noopHandle Hx_Typedef_loreline_importsfilehandler
		_ = noopHandle
		noopHandle = ((Hx_Typedef_loreline_importsfilehandler)(func(path string, cb func(data string)) {
			cb(HxStringNull)
		}))
		var dispatchResult func()
		_ = dispatchResult
		dispatchResult = func() {
			var _hx_tmp_73 *Hx_Obj_loreline_error
			_ = _hx_tmp_73
			if firstParseError != nil {
				_hx_tmp_73 = firstParseError
			} else {
				_hx_tmp_73 = wrappedLastError()
			}

			var finalError *Hx_Obj_loreline_error
			_ = finalError
			finalError = _hx_tmp_73
			if finalError != nil {
				Hx_Field_loreline_loreline__lastError = finalError
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(callback)), ((any)(nil))))) {
					callback(nil)
					return
				}

				panic(Hx_Field_haxe_exception_thrown(((any)(finalError))))
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(callback)), ((any)(nil))))) {
				callback(result)
			}
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = relativePaths.Len()
			for _g < _g1 {
				var _hx_tmp_74 int32
				_ = _hx_tmp_74
				_hx_tmp_74 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_74
				var relPath string
				_ = relPath
				relPath = relativePaths.Get(i)
				var transPath string
				_ = transPath
				transPath = translationPaths.Get(i)
				var _hx_tmp_75 string
				_ = _hx_tmp_75
				_hx_tmp_75 = transPath
				handleFile(_hx_tmp_75, func(content string) {
					pending = (pending - 1)
					if content != HxStringNull {
						{
							var hx_try_state int
							_ = hx_try_state

							func() {
								defer func() {
									if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
										if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_error").(*Hx_Obj_loreline_error); _hx_ok {
											_ = e
											{
												if firstParseError == nil {
													firstParseError = Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Failed to load translation file \"", transPath), "\": "), e.Hx_Field_message), nil)
												}
											}
										} else if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
											_ = e
											{
												if firstParseError == nil {
													var _hx_tmp_82 string
													_ = _hx_tmp_82
													_hx_tmp_82 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Failed to load translation file \"", transPath), "\": ")
													var _hx_tmp_83 string
													_ = _hx_tmp_83
													if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
														_hx_tmp_83 = "null"
													} else {
														_hx_tmp_83 = Hx_Field_std_string(e)
													}

													firstParseError = Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_82, (_hx_tmp_83)), nil)
												}
											}
										} else {
											panic(_hx_ex)
										}
									}
								}()
								{
									var _hx_tmp_76 string
									_ = _hx_tmp_76
									_hx_tmp_76 = content
									var _hx_tmp_77 HxNullable[string]
									_ = _hx_tmp_77
									_hx_tmp_77 = HxNullableString(transPath)
									var _hx_tmp_78 Hx_Typedef_loreline_importsfilehandler
									_ = _hx_tmp_78
									_hx_tmp_78 = noopHandle
									var transScript *Hx_Obj_loreline_script
									_ = transScript
									transScript = Hx_Field_loreline_loreline_parse(_hx_tmp_76, _hx_tmp_77, _hx_tmp_78, nil)
									if transScript != nil {
										var translations Hx_Typedef_map
										_ = translations
										translations = Hx_Field_loreline_astutils_extractTranslations((func() *Hx_Obj_loreline_astnode {
											_hx_up := transScript
											if _hx_up != nil {
												return &_hx_up.Hx_Obj_loreline_astnode
											}
											return nil
										})())
										{
											var _hx_reserved_map *Hx_Obj_haxe_imap
											_ = _hx_reserved_map
											_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
												_hx_icast_371 := translations
												if _hx_icast_371 == nil {
													return nil
												}
												Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_371, "Hx_Obj_haxe_imap")
												return &Hx_Obj_haxe_imap{VTable: _hx_icast_371.VTable.(Hx_Obj_VTable_haxe_imap)}
											})()
											var _g_map *Hx_Obj_haxe_imap
											_ = _g_map
											_g_map = _hx_reserved_map
											var _g_keys any
											_ = _g_keys
											_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
											for {
												var _hx_tmp_79 any
												_ = _hx_tmp_79
												_hx_tmp_79 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
												if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_79, HxMakeArray[any]().Dyn()))) {
													break
												}

												var _hx_tmp_80 any
												_ = _hx_tmp_80
												_hx_tmp_80 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
												var key string
												_ = key
												key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_80, HxMakeArray[any]().Dyn()))
												var _g_value *Hx_Obj_loreline_nstringliteral
												_ = _g_value
												_g_value = func() func(k string) *Hx_Obj_loreline_nstringliteral {
													var _hx_closure_src_845 func(k any) any
													_ = _hx_closure_src_845
													_hx_closure_src_845 = _g_map.VTable.Hx_Field_get
													return func(k string) *Hx_Obj_loreline_nstringliteral {
														var hx_dyncast_372 any
														_ = hx_dyncast_372
														hx_dyncast_372 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_845(((any)(k))), "Hx_Obj_loreline_nstringliteral")
														var _hx_tmp_81 *Hx_Obj_loreline_nstringliteral
														_ = _hx_tmp_81
														if hx_dyncast_372 != nil {
															_hx_tmp_81 = hx_dyncast_372.(interface {
																Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
															}).Hx_As_Hx_Obj_loreline_nstringliteral()
														} else {
															_hx_tmp_81 = nil
														}

														return _hx_tmp_81
													}
												}()(key)
												var _g_key string
												_ = _g_key
												_g_key = key
												var id string
												_ = id
												id = _g_key
												var str *Hx_Obj_loreline_nstringliteral
												_ = str
												str = _g_value
												func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nstringliteral) {
													var _hx_closure_src_846 func(_hx_param_key any, _hx_param_value any)
													_ = _hx_closure_src_846
													_hx_closure_src_846 = result.VTable.Hx_Field_set
													return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nstringliteral) {
														_hx_closure_src_846(((any)(_hx_param_key)), ((any)(_hx_param_value)))
													}
												}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(relPath, "#"), id), str)
											}
										}
									}
								}
							}()
							if hx_try_state == 1 {
								return
							}
						}
					}

					if (pending == 0) && allProcessed {
						dispatchResult()
					}
				})
			}
		}

		allProcessed = true
		if pending == 0 {
			dispatchResult()
		}

		if Hx_Field_loreline_loreline__lastError != nil {
			return nil
		} else {
			return result
		}
	}
}

func Hx_Field_loreline_loreline_relativePath(fromDir string, toPath string) string {
	var from string
	_ = from
	from = Hx_Field_haxe_io_path_normalize(fromDir)
	var to string
	_ = to
	to = Hx_Field_haxe_io_path_normalize(toPath)
	var fromParts HxArray[string]
	_ = fromParts
	fromParts = (Hx_Field_go_haxe_hxstring_split(from, "/"))
	var toParts HxArray[string]
	_ = toParts
	toParts = (Hx_Field_go_haxe_hxstring_split(to, "/"))
	var commonLen int32
	_ = commonLen
	commonLen = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = commonLen
		if _hx_tmp_2 < fromParts.Len() {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = commonLen
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_1 = (_hx_tmp_4 < toParts.Len())
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = (fromParts.Get(commonLen) == toParts.Get(commonLen))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		commonLen++
	}

	var upCount int32
	_ = upCount
	upCount = (fromParts.Len() - commonLen)
	var result HxArray[string]
	_ = result
	result = HxMakeArray[string]()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = upCount
		for _g < _g1 {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_5
			_ = (Hx_Array_Push[string](result, ".."))
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = commonLen
		var _g3 int32
		_ = _g3
		_g3 = toParts.Len()
		for _g2 < _g3 {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_6
			_ = (Hx_Array_Push[string](result, toParts.Get(i)))
		}
	}

	if result.Len() == 0 {
		return "."
	}

	return (Hx_Array_Join[string](result, "/"))
}

func Hx_Field_loreline_loreline_play(script *Hx_Obj_loreline_script, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, beatName__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter {
	beatName := HxDefault[HxNullable[string]]()
	_ = beatName

	if beatName__isOptional != nil {
		beatName = HxConvert[HxNullable[string]](beatName__isOptional)
	} else {

	}
	options := HxDefault[*Hx_Obj_loreline_interpreteroptions]()
	_ = options

	if options__isOptional != nil {
		options = HxConvert[*Hx_Obj_loreline_interpreteroptions](options__isOptional)
	} else {
		options = nil
	}
	{
		if (beatName).Valid == false {
			var hx_nullable_t76 HxNullable[string]
			_ = hx_nullable_t76
			hx_nullable_t76 = HxNullable[string]{Value: HxStringNull}

			beatName = HxNullable[string]{Value: (hx_nullable_t76).Value, Valid: (hx_nullable_t76).Valid}
		}

		var interpreter *Hx_Obj_loreline_interpreter
		_ = interpreter
		interpreter = Hx_Obj_loreline_interpreter_CreateInstance(script, handleDialogue, handleChoice, handleFinish, options, nil)
		interpreter.VTable.Hx_Field_start(beatName)
		return interpreter
	}
}

func Hx_Field_loreline_loreline_resume(script *Hx_Obj_loreline_script, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, saveData any, beatName__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter {
	beatName := HxDefault[HxNullable[string]]()
	_ = beatName

	if beatName__isOptional != nil {
		beatName = HxConvert[HxNullable[string]](beatName__isOptional)
	} else {

	}
	options := HxDefault[*Hx_Obj_loreline_interpreteroptions]()
	_ = options

	if options__isOptional != nil {
		options = HxConvert[*Hx_Obj_loreline_interpreteroptions](options__isOptional)
	} else {
		options = nil
	}
	{
		if (beatName).Valid == false {
			var hx_nullable_t77 HxNullable[string]
			_ = hx_nullable_t77
			hx_nullable_t77 = HxNullable[string]{Value: HxStringNull}

			beatName = HxNullable[string]{Value: (hx_nullable_t77).Value, Valid: (hx_nullable_t77).Valid}
		}

		var interpreter *Hx_Obj_loreline_interpreter
		_ = interpreter
		interpreter = Hx_Obj_loreline_interpreter_CreateInstance(script, handleDialogue, handleChoice, handleFinish, options, nil)
		interpreter.VTable.Hx_Field_restore(saveData)
		if (beatName).Valid != false {
			interpreter.VTable.Hx_Field_start(beatName)
		} else {
			interpreter.VTable.Hx_Field_resume()
		}

		return interpreter
	}
}

func Hx_Field_loreline_loreline_extractTranslations(script *Hx_Obj_loreline_script) Hx_Typedef_map {
	return Hx_Field_loreline_astutils_extractTranslations((func() *Hx_Obj_loreline_astnode {
		_hx_up := script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func Hx_Field_loreline_loreline_translationFormat(name string, enabled bool) {
	Hx_Field_loreline_translation_translationformats_translationFormat(name, enabled)
}

func Hx_Field_loreline_loreline_generateTranslationFile(script *Hx_Obj_loreline_script, existing__isOptional any, format__isOptional any, locale__isOptional any) string {
	existing := HxDefault[Hx_Typedef_map]()
	_ = existing

	if existing__isOptional != nil {
		existing = HxConvert[Hx_Typedef_map](existing__isOptional)
	} else {
		existing = nil
	}
	format := HxDefault[HxNullable[string]]()
	_ = format

	if format__isOptional != nil {
		format = HxConvert[HxNullable[string]](format__isOptional)
	} else {

	}
	locale := HxDefault[HxNullable[string]]()
	_ = locale

	if locale__isOptional != nil {
		locale = HxConvert[HxNullable[string]](locale__isOptional)
	} else {

	}
	{
		if (format).Valid == false {
			var hx_nullable_t79 HxNullable[string]
			_ = hx_nullable_t79
			hx_nullable_t79 = HxNullable[string]{Value: HxStringNull}

			format = HxNullable[string]{Value: (hx_nullable_t79).Value, Valid: (hx_nullable_t79).Valid}
		}

		if (locale).Valid == false {
			var hx_nullable_t78 HxNullable[string]
			_ = hx_nullable_t78
			hx_nullable_t78 = HxNullable[string]{Value: HxStringNull}

			locale = HxNullable[string]{Value: (hx_nullable_t78).Value, Valid: (hx_nullable_t78).Valid}
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (format).Valid == false {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if (format).Valid == true {
				_hx_tmp_1 = ((format).Value == "lor")
			} else {
				_hx_tmp_1 = false
			}

			_hx_tmp_0 = (_hx_tmp_1)
		}

		if _hx_tmp_0 {
			var _hx_tmp_2 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := script
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_3 Hx_Typedef_map
			_ = _hx_tmp_3
			_hx_tmp_3 = existing
			return Hx_Field_loreline_astutils_generateTranslationFile(_hx_tmp_2, _hx_tmp_3, Hx_Obj_loreline_printer_CreateInstance(nil, nil))
		}

		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if (locale).Valid == false {
			_hx_tmp_4 = true
		} else {
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if (locale).Valid == true {
				_hx_tmp_5 = ((locale).Value == "")
			} else {
				_hx_tmp_5 = false
			}

			_hx_tmp_4 = (_hx_tmp_5)
		}

		if _hx_tmp_4 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("A locale is required when generating a translation file in format '", (format).Value), "'"), nil)))))
		}

		if (format).Valid == false {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unknown translation file format: '", (format).Value), "'"), nil)))))
		} else {
			switch (format).Value {
			case "csv":
				{
					var _hx_tmp_6 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_6
					_hx_tmp_6 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := script
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_7 Hx_Typedef_map
					_ = _hx_tmp_7
					_hx_tmp_7 = existing
					return Hx_Field_loreline_translation_csvtranslation_fromScript(_hx_tmp_6, _hx_tmp_7, (locale).Value)
				}

			case "po":
				{
					var _hx_tmp_8 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_8
					_hx_tmp_8 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := script
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_9 Hx_Typedef_map
					_ = _hx_tmp_9
					_hx_tmp_9 = existing
					return Hx_Field_loreline_translation_potranslation_fromScript(_hx_tmp_8, _hx_tmp_9, (locale).Value)
				}

			case "tsv":
				{
					var _hx_tmp_10 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_10
					_hx_tmp_10 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := script
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_11 Hx_Typedef_map
					_ = _hx_tmp_11
					_hx_tmp_11 = existing
					return Hx_Field_loreline_translation_csvtranslation_tsvFromScript(_hx_tmp_10, _hx_tmp_11, (locale).Value)
				}

			case "xliff":
				{
					var _hx_tmp_12 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_12
					_hx_tmp_12 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := script
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_13 Hx_Typedef_map
					_ = _hx_tmp_13
					_hx_tmp_13 = existing
					return Hx_Field_loreline_translation_xlifftranslation_fromScript(_hx_tmp_12, _hx_tmp_13, (locale).Value)
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unknown translation file format: '", (format).Value), "'"), nil)))))
				}
			}
		}
	}
}

func Hx_Field_loreline_loreline_insertLocalizationKeys(content string, script *Hx_Obj_loreline_script, includeImports__isOptional any, reservedIds__isOptional any) string {
	includeImports := HxDefault[bool]()
	_ = includeImports

	if includeImports__isOptional != nil {
		includeImports = HxConvert[bool](includeImports__isOptional)
	} else {
		includeImports = true
	}
	reservedIds := HxDefault[Hx_Typedef_map]()
	_ = reservedIds

	if reservedIds__isOptional != nil {
		reservedIds = HxConvert[Hx_Typedef_map](reservedIds__isOptional)
	} else {
		reservedIds = nil
	}
	{
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = content
		var _hx_tmp_1 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
			_hx_up := script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		_hx_tmp_2 = includeImports
		var _hx_tmp_3 Hx_Typedef_map
		_ = _hx_tmp_3
		_hx_tmp_3 = reservedIds
		return Hx_Field_loreline_astutils_insertLocalizationKeys(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, nil)
	}
}

func Hx_Field_loreline_loreline_collectHashIds(content string, out__isOptional any) Hx_Typedef_map {
	out := HxDefault[Hx_Typedef_map]()
	_ = out

	if out__isOptional != nil {
		out = HxConvert[Hx_Typedef_map](out__isOptional)
	} else {
		out = nil
	}
	{
		return Hx_Field_loreline_astutils_collectHashIds(content, out)
	}
}

func Hx_Field_loreline_loreline_removeLocalizationKeys(content string, script *Hx_Obj_loreline_script) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = content
	return Hx_Field_loreline_astutils_removeLocalizationKeys(_hx_tmp_0, (func() *Hx_Obj_loreline_astnode {
		_hx_up := script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func Hx_Field_loreline_loreline_extractTranslatableEntries(script *Hx_Obj_loreline_script) HxArray[any] {
	return Hx_Field_loreline_astutils_extractTranslatableEntries((func() *Hx_Obj_loreline_astnode {
		_hx_up := script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func Hx_Field_loreline_loreline_hasUntaggedTranslatableStrings(script *Hx_Obj_loreline_script) bool {
	return Hx_Field_loreline_astutils_hasUntaggedTranslatableStrings((func() *Hx_Obj_loreline_astnode {
		_hx_up := script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func Hx_Field_loreline_loreline_print(script *Hx_Obj_loreline_script, indent__isOptional any, newline__isOptional any) string {
	indent := HxDefault[HxNullable[string]]()
	_ = indent

	if indent__isOptional != nil {
		indent = HxConvert[HxNullable[string]](indent__isOptional)
	} else {

	}
	newline := HxDefault[HxNullable[string]]()
	_ = newline

	if newline__isOptional != nil {
		newline = HxConvert[HxNullable[string]](newline__isOptional)
	} else {

	}
	{
		if (indent).Valid == false {
			var hx_nullable_t81 HxNullable[string]
			_ = hx_nullable_t81
			hx_nullable_t81 = HxNullable[string]{Value: HxStringNull}

			indent = HxNullable[string]{Value: (hx_nullable_t81).Value, Valid: (hx_nullable_t81).Valid}
		}

		if (newline).Valid == false {
			var hx_nullable_t80 HxNullable[string]
			_ = hx_nullable_t80
			hx_nullable_t80 = HxNullable[string]{Value: HxStringNull}

			newline = HxNullable[string]{Value: (hx_nullable_t80).Value, Valid: (hx_nullable_t80).Valid}
		}

		var indent1 HxNullable[string]
		_ = indent1
		indent1 = indent
		var newline1 HxNullable[string]
		_ = newline1
		newline1 = newline
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		if (indent1).Valid != false {
			_hx_tmp_0 = (indent1).Value
		} else {
			_hx_tmp_0 = "  "
		}

		var _hx_tmp_1 string
		_ = _hx_tmp_1
		if (newline1).Valid != false {
			_hx_tmp_1 = (newline1).Value
		} else {
			_hx_tmp_1 = "\n"
		}

		var printer *Hx_Obj_loreline_printer
		_ = printer
		printer = Hx_Obj_loreline_printer_CreateInstance(_hx_tmp_0, _hx_tmp_1)
		return printer.VTable.Hx_Field_print((func() *Hx_Obj_loreline_node {
			_hx_up := script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
	}
}

func Hx_Field_loreline_loreline_update(delta float64) {
	Hx_Field_loreline_timer_update(delta)
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_lexererror]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_parseerror]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
