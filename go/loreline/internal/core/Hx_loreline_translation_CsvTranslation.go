package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation_csvtranslation() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation_csvtranslation() {
}

type Hx_Obj_VTable_loreline_translation_csvtranslation interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation_csvtranslation struct {
	VTable Hx_Obj_VTable_loreline_translation_csvtranslation `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_translation_csvtranslation_CreateEmptyInstance() *Hx_Obj_loreline_translation_csvtranslation {
	obj := &Hx_Obj_loreline_translation_csvtranslation{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_translation_csvtranslation_CreateInstance() *Hx_Obj_loreline_translation_csvtranslation {
	obj := Hx_Obj_loreline_translation_csvtranslation_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_translation_csvtranslation_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation.CsvTranslation",
	HxMakeArray[string]("toLoreline", "tsvToLoreline", "convert", "fromScript", "tsvFromScript", "generate", "escapeForCsvField", "findColumn", "findLocaleColumn", "primarySubtag", "parseRows", "escapeLineForComment"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation_csvtranslation_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_translation_csvtranslation_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation_csvtranslation_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toLoreline":
			return Hx_Field_loreline_translation_csvtranslation_toLoreline
		case "tsvToLoreline":
			return Hx_Field_loreline_translation_csvtranslation_tsvToLoreline
		case "convert":
			return Hx_Field_loreline_translation_csvtranslation_convert
		case "fromScript":
			return Hx_Field_loreline_translation_csvtranslation_fromScript
		case "tsvFromScript":
			return Hx_Field_loreline_translation_csvtranslation_tsvFromScript
		case "generate":
			return Hx_Field_loreline_translation_csvtranslation_generate
		case "escapeForCsvField":
			return Hx_Field_loreline_translation_csvtranslation_escapeForCsvField
		case "findColumn":
			return Hx_Field_loreline_translation_csvtranslation_findColumn
		case "findLocaleColumn":
			return Hx_Field_loreline_translation_csvtranslation_findLocaleColumn
		case "primarySubtag":
			return Hx_Field_loreline_translation_csvtranslation_primarySubtag
		case "parseRows":
			return Hx_Field_loreline_translation_csvtranslation_parseRows
		case "escapeLineForComment":
			return Hx_Field_loreline_translation_csvtranslation_escapeLineForComment
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_translation_csvtranslation) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation_csvtranslation_RTTI
}

func (this *Hx_Obj_loreline_translation_csvtranslation) Hx_As_Hx_Obj_loreline_translation_csvtranslation() *Hx_Obj_loreline_translation_csvtranslation {
	return this
}

func Hx_Field_loreline_translation_csvtranslation_toLoreline(content string, locale string) string {
	return Hx_Field_loreline_translation_csvtranslation_convert(content, 44, locale)
}

func Hx_Field_loreline_translation_csvtranslation_tsvToLoreline(content string, locale string) string {
	return Hx_Field_loreline_translation_csvtranslation_convert(content, 9, locale)
}

func Hx_Field_loreline_translation_csvtranslation_convert(content string, sep int32, locale string) string {
	var rows HxArray[HxArray[string]]
	_ = rows
	rows = Hx_Field_loreline_translation_csvtranslation_parseRows(content, sep)
	if rows.Len() < 1 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Invalid CSV: header row required", nil)))))
	}

	if rows.Len() < 2 {
		return ""
	}

	var header HxArray[string]
	_ = header
	header = rows.Get(0)
	var _hx_tmp_0 HxArray[string]
	_ = _hx_tmp_0
	_hx_tmp_0 = header
	var keyCol int32
	_ = keyCol
	keyCol = Hx_Field_loreline_translation_csvtranslation_findColumn(_hx_tmp_0, HxMakeArray[string]("key"))
	if keyCol < 0 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Invalid CSV: 'key' column not found in header", nil)))))
	}

	var _hx_tmp_1 HxArray[string]
	_ = _hx_tmp_1
	_hx_tmp_1 = header
	var sourceCol int32
	_ = sourceCol
	sourceCol = Hx_Field_loreline_translation_csvtranslation_findColumn(_hx_tmp_1, HxMakeArray[string]("source", "en", "original"))
	var transCol int32
	_ = transCol
	transCol = Hx_Field_loreline_translation_csvtranslation_findLocaleColumn(header, locale)
	if transCol < 0 {
		if header.Len() >= 2 {
			transCol = (header.Len() - 1)
		} else {
			return ""
		}
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var first bool
	_ = first
	first = true
	{
		var _g int32
		_ = _g
		_g = 1
		var _g1 int32
		_ = _g1
		_g1 = rows.Len()
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			var row HxArray[string]
			_ = row
			row = rows.Get(i)
			if row.Len() <= keyCol {
				continue
			}

			var key string
			_ = key
			key = row.Get(keyCol)
			if (key == HxStringNull) || (key == "") {
				continue
			}

			if row.Len() <= transCol {
				continue
			}

			var translation string
			_ = translation
			translation = row.Get(transCol)
			if (translation == HxStringNull) || (translation == "") {
				continue
			}

			var _hx_tmp_3 HxNullable[string]
			_ = _hx_tmp_3
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if sourceCol >= 0 {
				_hx_tmp_4 = (row.Len() > sourceCol)
			} else {
				_hx_tmp_4 = false
			}

			if _hx_tmp_4 {
				var hx_nullable_a68 HxArray[string]
				_ = hx_nullable_a68
				hx_nullable_a68 = row
				var hx_nullable_i69 int32
				_ = hx_nullable_i69
				hx_nullable_i69 = sourceCol
				var _hx_tmp_5 HxNullable[string]
				_ = _hx_tmp_5
				if hx_nullable_i69 >= 0 && hx_nullable_i69 < hx_nullable_a68.Len() {
					_hx_tmp_5 = HxNullable[string]{Value: hx_nullable_a68.Get(hx_nullable_i69), Valid: true}
				} else {
					_hx_tmp_5 = HxNullable[string]{Value: HxStringNull}
				}

				_hx_tmp_3 = _hx_tmp_5
			} else {
				_hx_tmp_3 = HxNullable[string]{Value: HxStringNull}
			}

			var src HxNullable[string]
			_ = src
			src = _hx_tmp_3
			if !first {
				func() func(x string) {
					var _hx_closure_src_759 func(x any)
					_ = _hx_closure_src_759
					_hx_closure_src_759 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_759(((any)(x)))
					}
				}()("\n")
			}

			first = false
			func() func(x string) {
				var _hx_closure_src_760 func(x any)
				_ = _hx_closure_src_760
				_hx_closure_src_760 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_760(((any)(x)))
				}
			}()("#")
			func() func(x string) {
				var _hx_closure_src_761 func(x any)
				_ = _hx_closure_src_761
				_hx_closure_src_761 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_761(((any)(x)))
				}
			}()(key)
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			if (src).Valid != false {
				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				if (src).Valid == false {
					_hx_tmp_7 = true
				} else {
					_hx_tmp_7 = ((src).Value != "")
				}

				_hx_tmp_6 = (_hx_tmp_7)
			} else {
				_hx_tmp_6 = false
			}

			if _hx_tmp_6 {
				func() func(x string) {
					var _hx_closure_src_762 func(x any)
					_ = _hx_closure_src_762
					_hx_closure_src_762 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_762(((any)(x)))
					}
				}()(" // ")
				func() func(x string) {
					var _hx_closure_src_763 func(x any)
					_ = _hx_closure_src_763
					_hx_closure_src_763 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_763(((any)(x)))
					}
				}()((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(src.Value, "\n")), " ")), "\r")), "")))
			}

			func() func(x string) {
				var _hx_closure_src_764 func(x any)
				_ = _hx_closure_src_764
				_hx_closure_src_764 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_764(((any)(x)))
				}
			}()("\n")
			func() func(x string) {
				var _hx_closure_src_765 func(x any)
				_ = _hx_closure_src_765
				_hx_closure_src_765 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_765(((any)(x)))
				}
			}()(translation)
			func() func(x string) {
				var _hx_closure_src_766 func(x any)
				_ = _hx_closure_src_766
				_hx_closure_src_766 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_766(((any)(x)))
				}
			}()("\n")
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_csvtranslation_fromScript(script *Hx_Obj_loreline_astnode, existing Hx_Typedef_map, locale string) string {
	return Hx_Field_loreline_translation_csvtranslation_generate(script, existing, locale, 44)
}

func Hx_Field_loreline_translation_csvtranslation_tsvFromScript(script *Hx_Obj_loreline_astnode, existing Hx_Typedef_map, locale string) string {
	return Hx_Field_loreline_translation_csvtranslation_generate(script, existing, locale, 9)
}

func Hx_Field_loreline_translation_csvtranslation_generate(script *Hx_Obj_loreline_astnode, existing Hx_Typedef_map, locale string, sep int32) string {
	var entries HxArray[any]
	_ = entries
	entries = Hx_Field_loreline_astutils_extractTranslatableEntries(script)
	var printer *Hx_Obj_loreline_printer
	_ = printer
	printer = Hx_Obj_loreline_printer_CreateInstance(nil, nil)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var sepStr string
	_ = sepStr
	sepStr = Hx_Field_go_haxe_hxstring_fromCharCode(sep)
	func() func(x string) {
		var _hx_closure_src_767 func(x any)
		_ = _hx_closure_src_767
		_hx_closure_src_767 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_767(((any)(x)))
		}
	}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField("key", sep))
	func() func(x string) {
		var _hx_closure_src_768 func(x any)
		_ = _hx_closure_src_768
		_hx_closure_src_768 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_768(((any)(x)))
		}
	}()(sepStr)
	func() func(x string) {
		var _hx_closure_src_769 func(x any)
		_ = _hx_closure_src_769
		_hx_closure_src_769 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_769(((any)(x)))
		}
	}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField("source", sep))
	func() func(x string) {
		var _hx_closure_src_770 func(x any)
		_ = _hx_closure_src_770
		_hx_closure_src_770 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_770(((any)(x)))
		}
	}()(sepStr)
	func() func(x string) {
		var _hx_closure_src_771 func(x any)
		_ = _hx_closure_src_771
		_hx_closure_src_771 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_771(((any)(x)))
		}
	}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField(locale, sep))
	func() func(x string) {
		var _hx_closure_src_772 func(x any)
		_ = _hx_closure_src_772
		_hx_closure_src_772 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_772(((any)(x)))
		}
	}()("\n")
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(entries)) {
				break
			}

			var entry any
			_ = entry
			entry = entries.Get(_g)
			_g++
			var hx_dyncast_359 any
			_ = hx_dyncast_359
			hx_dyncast_359 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(entry, "str"), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_1 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_1
			if hx_dyncast_359 != nil {
				_hx_tmp_1 = hx_dyncast_359.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
			} else {
				_hx_tmp_1 = nil
			}

			var source string
			_ = source
			source = printer.VTable.Hx_Field_printStringLiteralAsText(_hx_tmp_1)
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if existing != nil {
				_hx_tmp_3 = func() func(_hx_param_key string) bool {
					var _hx_closure_src_773 func(_hx_param_key any) bool
					_ = _hx_closure_src_773
					_hx_closure_src_773 = existing.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_773(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				_hx_tmp_2 = printer.VTable.Hx_Field_printStringLiteralAsText(func() func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
					var _hx_closure_src_774 func(_hx_param_key any) any
					_ = _hx_closure_src_774
					_hx_closure_src_774 = existing.VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
						var hx_dyncast_360 any
						_ = hx_dyncast_360
						hx_dyncast_360 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_774(((any)(_hx_param_key))), "Hx_Obj_loreline_nstringliteral")
						var _hx_tmp_4 *Hx_Obj_loreline_nstringliteral
						_ = _hx_tmp_4
						if hx_dyncast_360 != nil {
							_hx_tmp_4 = hx_dyncast_360.(interface {
								Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
							}).Hx_As_Hx_Obj_loreline_nstringliteral()
						} else {
							_hx_tmp_4 = nil
						}

						return _hx_tmp_4
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id"))))
			} else {
				_hx_tmp_2 = source
			}

			var translation string
			_ = translation
			translation = _hx_tmp_2
			func() func(x string) {
				var _hx_closure_src_775 func(x any)
				_ = _hx_closure_src_775
				_hx_closure_src_775 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_775(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")), sep))
			func() func(x string) {
				var _hx_closure_src_776 func(x any)
				_ = _hx_closure_src_776
				_hx_closure_src_776 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_776(((any)(x)))
				}
			}()(sepStr)
			func() func(x string) {
				var _hx_closure_src_777 func(x any)
				_ = _hx_closure_src_777
				_hx_closure_src_777 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_777(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField(source, sep))
			func() func(x string) {
				var _hx_closure_src_778 func(x any)
				_ = _hx_closure_src_778
				_hx_closure_src_778 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_778(((any)(x)))
				}
			}()(sepStr)
			func() func(x string) {
				var _hx_closure_src_779 func(x any)
				_ = _hx_closure_src_779
				_hx_closure_src_779 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_779(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_csvtranslation_escapeForCsvField(translation, sep))
			func() func(x string) {
				var _hx_closure_src_780 func(x any)
				_ = _hx_closure_src_780
				_hx_closure_src_780 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_780(((any)(x)))
				}
			}()("\n")
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_csvtranslation_escapeForCsvField(s string, sep int32) string {
	var needsQuoting bool
	_ = needsQuoting
	needsQuoting = false
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(s))
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(s, i))
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			if (c).Valid == false {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = (c).Value
			}

			var c1 int32
			_ = c1
			c1 = _hx_tmp_1
			if (((c1 == sep) || (c1 == 34)) || (c1 == 10)) || (c1 == 13) {
				needsQuoting = true
				break
			}
		}
	}

	if !needsQuoting {
		return s
	}

	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("\"", (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\"")), "\"\""))), "\"")
}

func Hx_Field_loreline_translation_csvtranslation_findColumn(header HxArray[string], names HxArray[string]) int32 {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = header.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var h string
			_ = h
			h = Hx_Field_go_haxe_hxstring_toLowerCase(header.Get(i))
			{
				var tmp__g1_1 int32
				_ = tmp__g1_1
				tmp__g1_1 = 0
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = tmp__g1_1
					if !(_hx_tmp_1 < names.Len()) {
						break
					}

					var n string
					_ = n
					n = names.Get(tmp__g1_1)
					tmp__g1_1++
					if h == n {
						return i
					}
				}
			}
		}
	}

	return -1
}

func Hx_Field_loreline_translation_csvtranslation_findLocaleColumn(header HxArray[string], locale string) int32 {
	var locLow string
	_ = locLow
	locLow = Hx_Field_go_haxe_hxstring_toLowerCase(locale)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = header.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			if Hx_Field_go_haxe_hxstring_toLowerCase(header.Get(i)) == locLow {
				return i
			}
		}
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = locLow
	var dash int32
	_ = dash
	dash = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, "-", HxNullable[int32]{}))
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	if dash < 0 {
		_hx_tmp_2 = locLow
	} else {
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		if (HxNullable[int32]{Value: dash, Valid: true}).Valid != false {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = locLow
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = _hx_tmp_4
			_hx_tmp_3 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, 0, HxNullable[int32]{Value: dash, Valid: true}))
		} else {
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = locLow
			var _hx_tmp_7 string
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_3 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, 0, HxNullable[int32]{}))
		}

		_hx_tmp_2 = _hx_tmp_3
	}

	var primary string
	_ = primary
	primary = _hx_tmp_2
	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = header.Len()
		for _g2 < _g3 {
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_8
			var s string
			_ = s
			s = Hx_Field_go_haxe_hxstring_toLowerCase(header.Get(i))
			var _hx_tmp_9 string
			_ = _hx_tmp_9
			_hx_tmp_9 = s
			var tmp_dash_1 int32
			_ = tmp_dash_1
			tmp_dash_1 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_9, "-", HxNullable[int32]{}))
			var _hx_tmp_10 string
			_ = _hx_tmp_10
			if tmp_dash_1 < 0 {
				_hx_tmp_10 = s
			} else {
				var _hx_tmp_11 string
				_ = _hx_tmp_11
				if (HxNullable[int32]{Value: tmp_dash_1, Valid: true}).Valid != false {
					var _hx_tmp_12 string
					_ = _hx_tmp_12
					_hx_tmp_12 = s
					var _hx_tmp_13 string
					_ = _hx_tmp_13
					_hx_tmp_13 = _hx_tmp_12
					_hx_tmp_11 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_13, 0, HxNullable[int32]{Value: tmp_dash_1, Valid: true}))
				} else {
					var _hx_tmp_14 string
					_ = _hx_tmp_14
					_hx_tmp_14 = s
					var _hx_tmp_15 string
					_ = _hx_tmp_15
					_hx_tmp_15 = _hx_tmp_14
					_hx_tmp_11 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, 0, HxNullable[int32]{}))
				}

				_hx_tmp_10 = _hx_tmp_11
			}

			if (_hx_tmp_10) == primary {
				return i
			}
		}
	}

	return -1
}

func Hx_Field_loreline_translation_csvtranslation_primarySubtag(s string) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = s
	var dash int32
	_ = dash
	dash = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "-", HxNullable[int32]{}))
	if dash < 0 {
		return s
	}

	if (HxNullable[int32]{Value: dash, Valid: true}).Valid != false {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 0, HxNullable[int32]{Value: dash, Valid: true}))
	} else {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = s
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, HxNullable[int32]{}))
	}
}

func Hx_Field_loreline_translation_csvtranslation_parseRows(content string, sep int32) HxArray[HxArray[string]] {
	var rows HxArray[HxArray[string]]
	_ = rows
	rows = HxMakeArray[HxArray[string]]()
	var row HxArray[string]
	_ = row
	row = HxMakeArray[string]()
	var field *Hx_Obj_stringbuf
	_ = field
	field = Hx_Obj_stringbuf_CreateInstance()
	var inQuotes bool
	_ = inQuotes
	inQuotes = false
	var i int32
	_ = i
	i = 0
	var chars HxArray[int32]
	_ = chars
	chars = Hx_Field_loreline_utf8_toCodes(content)
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = HxMakeArrayView[any](chars).Len()
	for i < _hx_reserved_len {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if i < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = i
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](chars).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(i))))
		}

		var c int32
		_ = c
		c = _hx_tmp_0
		if inQuotes {
			if c == 34 {
				var tmp bool
				_ = tmp
				if (i + 1) < _hx_reserved_len {
					var pos int32
					_ = pos
					pos = (i + 1)
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					var _hx_tmp_5 bool
					_ = _hx_tmp_5
					if pos < 0 {
						_hx_tmp_5 = true
					} else {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = pos
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = _hx_tmp_6
						_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](chars).Len())
					}

					if _hx_tmp_5 {
						_hx_tmp_4 = -1
					} else {
						_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
					}

					tmp = ((_hx_tmp_4) == 34)
				} else {
					tmp = false
				}

				if tmp {
					field.VTable.Hx_Field_addChar(34)
					i = (i + 2)
				} else {
					inQuotes = false
					i++
				}
			} else {
				field.VTable.Hx_Field_addChar(c)
				i++
			}
		} else {
			if c == 34 {
				inQuotes = true
				i++
			} else {
				if c == sep {
					var _hx_tmp_8 HxArray[string]
					_ = _hx_tmp_8
					_hx_tmp_8 = row
					_ = (Hx_Array_Push[string](_hx_tmp_8, field.VTable.Hx_Field_toString()))
					field = Hx_Obj_stringbuf_CreateInstance()
					i++
				} else {
					if (c == 10) || (c == 13) {
						var _hx_tmp_9 HxArray[string]
						_ = _hx_tmp_9
						_hx_tmp_9 = row
						_ = (Hx_Array_Push[string](_hx_tmp_9, field.VTable.Hx_Field_toString()))
						field = Hx_Obj_stringbuf_CreateInstance()
						var _hx_tmp_10 bool
						_ = _hx_tmp_10
						if row.Len() > 1 {
							_hx_tmp_10 = true
						} else {
							var _hx_tmp_11 bool
							_ = _hx_tmp_11
							if row.Len() == 1 {
								_hx_tmp_11 = (row.Get(0) != "")
							} else {
								_hx_tmp_11 = false
							}

							_hx_tmp_10 = _hx_tmp_11
						}

						if _hx_tmp_10 {
							_ = (Hx_Array_Push[HxArray[string]](rows, row))
						}

						row = HxMakeArray[string]()
						var tmp bool
						_ = tmp
						if (c == 13) && ((i + 1) < _hx_reserved_len) {
							var pos int32
							_ = pos
							pos = (i + 1)
							var _hx_tmp_12 int32
							_ = _hx_tmp_12
							var _hx_tmp_13 bool
							_ = _hx_tmp_13
							if pos < 0 {
								_hx_tmp_13 = true
							} else {
								var _hx_tmp_14 int32
								_ = _hx_tmp_14
								_hx_tmp_14 = pos
								var _hx_tmp_15 int32
								_ = _hx_tmp_15
								_hx_tmp_15 = _hx_tmp_14
								_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](chars).Len())
							}

							if _hx_tmp_13 {
								_hx_tmp_12 = -1
							} else {
								_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
							}

							tmp = ((_hx_tmp_12) == 10)
						} else {
							tmp = false
						}

						if tmp {
							i = (i + 2)
						} else {
							i++
						}
					} else {
						field.VTable.Hx_Field_addChar(c)
						i++
					}
				}
			}
		}
	}

	var _hx_tmp_16 bool
	_ = _hx_tmp_16
	if field.VTable.Hx_Field_toString() != "" {
		_hx_tmp_16 = true
	} else {
		_hx_tmp_16 = (row.Len() > 0)
	}

	if _hx_tmp_16 {
		var _hx_tmp_17 HxArray[string]
		_ = _hx_tmp_17
		_hx_tmp_17 = row
		_ = (Hx_Array_Push[string](_hx_tmp_17, field.VTable.Hx_Field_toString()))
		var _hx_tmp_18 bool
		_ = _hx_tmp_18
		if row.Len() > 1 {
			_hx_tmp_18 = true
		} else {
			var _hx_tmp_19 bool
			_ = _hx_tmp_19
			if row.Len() == 1 {
				_hx_tmp_19 = (row.Get(0) != "")
			} else {
				_hx_tmp_19 = false
			}

			_hx_tmp_18 = _hx_tmp_19
		}

		if _hx_tmp_18 {
			_ = (Hx_Array_Push[HxArray[string]](rows, row))
		}
	}

	return rows
}

func Hx_Field_loreline_translation_csvtranslation_escapeLineForComment(s string) string {
	return (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\n")), " ")), "\r")), ""))
}

func init() {
	hxRegisterArrayConstructor[HxArray[string]]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
