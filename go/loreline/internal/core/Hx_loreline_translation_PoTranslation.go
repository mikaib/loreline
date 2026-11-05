package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation_potranslation() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation_potranslation() {
}

type Hx_Obj_VTable_loreline_translation_potranslation interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation_potranslation struct {
	VTable Hx_Obj_VTable_loreline_translation_potranslation `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_translation_potranslation_CreateEmptyInstance() *Hx_Obj_loreline_translation_potranslation {
	obj := &Hx_Obj_loreline_translation_potranslation{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_translation_potranslation_CreateInstance() *Hx_Obj_loreline_translation_potranslation {
	obj := Hx_Obj_loreline_translation_potranslation_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_translation_potranslation_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation.PoTranslation",
	HxMakeArray[string]("toLoreline", "fromScript", "escapeForPoString", "parse", "readKeyword", "readQuotedString", "trimLeft", "escapeLineForComment", "escapeLineForBody"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation_potranslation_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_translation_potranslation_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation_potranslation_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toLoreline":
			return Hx_Field_loreline_translation_potranslation_toLoreline
		case "fromScript":
			return Hx_Field_loreline_translation_potranslation_fromScript
		case "escapeForPoString":
			return Hx_Field_loreline_translation_potranslation_escapeForPoString
		case "parse":
			return Hx_Field_loreline_translation_potranslation_parse
		case "readKeyword":
			return Hx_Field_loreline_translation_potranslation_readKeyword
		case "readQuotedString":
			return Hx_Field_loreline_translation_potranslation_readQuotedString
		case "trimLeft":
			return Hx_Field_loreline_translation_potranslation_trimLeft
		case "escapeLineForComment":
			return Hx_Field_loreline_translation_potranslation_escapeLineForComment
		case "escapeLineForBody":
			return Hx_Field_loreline_translation_potranslation_escapeLineForBody
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_translation_potranslation) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation_potranslation_RTTI
}

func (this *Hx_Obj_loreline_translation_potranslation) Hx_As_Hx_Obj_loreline_translation_potranslation() *Hx_Obj_loreline_translation_potranslation {
	return this
}

func Hx_Field_loreline_translation_potranslation_toLoreline(content string, locale string) string {
	var entries HxArray[*Hx_Obj_loreline_translation__potranslation_poentry]
	_ = entries
	entries = Hx_Field_loreline_translation_potranslation_parse(content)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
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
			if !(_hx_tmp_0 < entries.Len()) {
				break
			}

			var entry *Hx_Obj_loreline_translation__potranslation_poentry
			_ = entry
			entry = entries.Get(_g)
			_g++
			if (entry.Hx_Field_msgstr == HxStringNull) || (entry.Hx_Field_msgstr == "") {
				continue
			}

			var _hx_tmp_1 string
			_ = _hx_tmp_1
			if (entry.Hx_Field_msgctxt != HxStringNull) && (entry.Hx_Field_msgctxt != "") {
				_hx_tmp_1 = entry.Hx_Field_msgctxt
			} else {
				_hx_tmp_1 = entry.Hx_Field_msgid
			}

			var key string
			_ = key
			key = _hx_tmp_1
			if (key == HxStringNull) || (key == "") {
				continue
			}

			if !first {
				func() func(x string) {
					var _hx_closure_src_781 func(x any)
					_ = _hx_closure_src_781
					_hx_closure_src_781 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_781(((any)(x)))
					}
				}()("\n")
			}

			first = false
			func() func(x string) {
				var _hx_closure_src_782 func(x any)
				_ = _hx_closure_src_782
				_hx_closure_src_782 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_782(((any)(x)))
				}
			}()("#")
			func() func(x string) {
				var _hx_closure_src_783 func(x any)
				_ = _hx_closure_src_783
				_hx_closure_src_783 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_783(((any)(x)))
				}
			}()(key)
			if ((entry.Hx_Field_msgctxt != HxStringNull) && (entry.Hx_Field_msgid != HxStringNull)) && (entry.Hx_Field_msgid != "") {
				func() func(x string) {
					var _hx_closure_src_784 func(x any)
					_ = _hx_closure_src_784
					_hx_closure_src_784 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_784(((any)(x)))
					}
				}()(" // ")
				func() func(x string) {
					var _hx_closure_src_785 func(x any)
					_ = _hx_closure_src_785
					_hx_closure_src_785 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_785(((any)(x)))
					}
				}()(Hx_Field_loreline_translation_potranslation_escapeLineForComment(entry.Hx_Field_msgid))
			}

			func() func(x string) {
				var _hx_closure_src_786 func(x any)
				_ = _hx_closure_src_786
				_hx_closure_src_786 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_786(((any)(x)))
				}
			}()("\n")
			func() func(x string) {
				var _hx_closure_src_787 func(x any)
				_ = _hx_closure_src_787
				_hx_closure_src_787 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_787(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_potranslation_escapeLineForBody(entry.Hx_Field_msgstr))
			func() func(x string) {
				var _hx_closure_src_788 func(x any)
				_ = _hx_closure_src_788
				_hx_closure_src_788 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_788(((any)(x)))
				}
			}()("\n")
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_potranslation_fromScript(script *Hx_Obj_loreline_astnode, existing Hx_Typedef_map, locale string) string {
	var entries HxArray[any]
	_ = entries
	entries = Hx_Field_loreline_astutils_extractTranslatableEntries(script)
	var printer *Hx_Obj_loreline_printer
	_ = printer
	printer = Hx_Obj_loreline_printer_CreateInstance(nil, nil)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	func() func(x string) {
		var _hx_closure_src_789 func(x any)
		_ = _hx_closure_src_789
		_hx_closure_src_789 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_789(((any)(x)))
		}
	}()("msgid \"\"\n")
	func() func(x string) {
		var _hx_closure_src_790 func(x any)
		_ = _hx_closure_src_790
		_hx_closure_src_790 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_790(((any)(x)))
		}
	}()("msgstr \"\"\n")
	func() func(x string) {
		var _hx_closure_src_791 func(x any)
		_ = _hx_closure_src_791
		_hx_closure_src_791 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_791(((any)(x)))
		}
	}()("\"Content-Type: text/plain; charset=UTF-8\\n\"\n")
	func() func(x string) {
		var _hx_closure_src_792 func(x any)
		_ = _hx_closure_src_792
		_hx_closure_src_792 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_792(((any)(x)))
		}
	}()("\"Language: ")
	func() func(x string) {
		var _hx_closure_src_793 func(x any)
		_ = _hx_closure_src_793
		_hx_closure_src_793 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_793(((any)(x)))
		}
	}()(Hx_Field_loreline_translation_potranslation_escapeForPoString(locale))
	func() func(x string) {
		var _hx_closure_src_794 func(x any)
		_ = _hx_closure_src_794
		_hx_closure_src_794 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_794(((any)(x)))
		}
	}()("\\n\"\n")
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
			func() func(x string) {
				var _hx_closure_src_795 func(x any)
				_ = _hx_closure_src_795
				_hx_closure_src_795 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_795(((any)(x)))
				}
			}()("\n")
			var hx_dyncast_361 any
			_ = hx_dyncast_361
			hx_dyncast_361 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(entry, "str"), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_1 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_1
			if hx_dyncast_361 != nil {
				_hx_tmp_1 = hx_dyncast_361.(interface {
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
					var _hx_closure_src_796 func(_hx_param_key any) bool
					_ = _hx_closure_src_796
					_hx_closure_src_796 = existing.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_796(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				_hx_tmp_2 = printer.VTable.Hx_Field_printStringLiteralAsText(func() func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
					var _hx_closure_src_797 func(_hx_param_key any) any
					_ = _hx_closure_src_797
					_hx_closure_src_797 = existing.VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
						var hx_dyncast_362 any
						_ = hx_dyncast_362
						hx_dyncast_362 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_797(((any)(_hx_param_key))), "Hx_Obj_loreline_nstringliteral")
						var _hx_tmp_4 *Hx_Obj_loreline_nstringliteral
						_ = _hx_tmp_4
						if hx_dyncast_362 != nil {
							_hx_tmp_4 = hx_dyncast_362.(interface {
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
				var _hx_closure_src_798 func(x any)
				_ = _hx_closure_src_798
				_hx_closure_src_798 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_798(((any)(x)))
				}
			}()("msgctxt \"")
			func() func(x string) {
				var _hx_closure_src_799 func(x any)
				_ = _hx_closure_src_799
				_hx_closure_src_799 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_799(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_potranslation_escapeForPoString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id"))))
			func() func(x string) {
				var _hx_closure_src_800 func(x any)
				_ = _hx_closure_src_800
				_hx_closure_src_800 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_800(((any)(x)))
				}
			}()("\"\n")
			func() func(x string) {
				var _hx_closure_src_801 func(x any)
				_ = _hx_closure_src_801
				_hx_closure_src_801 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_801(((any)(x)))
				}
			}()("msgid \"")
			func() func(x string) {
				var _hx_closure_src_802 func(x any)
				_ = _hx_closure_src_802
				_hx_closure_src_802 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_802(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_potranslation_escapeForPoString(source))
			func() func(x string) {
				var _hx_closure_src_803 func(x any)
				_ = _hx_closure_src_803
				_hx_closure_src_803 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_803(((any)(x)))
				}
			}()("\"\n")
			func() func(x string) {
				var _hx_closure_src_804 func(x any)
				_ = _hx_closure_src_804
				_hx_closure_src_804 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_804(((any)(x)))
				}
			}()("msgstr \"")
			func() func(x string) {
				var _hx_closure_src_805 func(x any)
				_ = _hx_closure_src_805
				_hx_closure_src_805 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_805(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_potranslation_escapeForPoString(translation))
			func() func(x string) {
				var _hx_closure_src_806 func(x any)
				_ = _hx_closure_src_806
				_hx_closure_src_806 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_806(((any)(x)))
				}
			}()("\"\n")
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_potranslation_escapeForPoString(s string) string {
	var out string
	_ = out
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\\")), "\\\\"))
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(out, "\"")), "\\\""))
	return out
}

func Hx_Field_loreline_translation_potranslation_parse(content string) HxArray[*Hx_Obj_loreline_translation__potranslation_poentry] {
	var lines HxArray[string]
	_ = lines
	lines = (Hx_Field_go_haxe_hxstring_split(content, "\n"))
	var entries HxArray[*Hx_Obj_loreline_translation__potranslation_poentry]
	_ = entries
	entries = HxMakeArray[*Hx_Obj_loreline_translation__potranslation_poentry]()
	var current *Hx_Obj_loreline_translation__potranslation_poentry
	_ = current
	var i int32
	_ = i
	i = 0
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < lines.Len()) {
			break
		}

		var line string
		_ = line
		line = lines.Get(i)
		var tmp bool
		_ = tmp
		if int32(utf8.RuneCountInString(line)) > 0 {
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = line
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, (int32(utf8.RuneCountInString(line)) - 1)))
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			if (c).Valid == false {
				_hx_tmp_2 = -1
			} else {
				_hx_tmp_2 = (c).Value
			}

			tmp = ((_hx_tmp_2) == 13)
		} else {
			tmp = false
		}

		if tmp {
			var _hx_reserved_len HxNullable[int32]
			_ = _hx_reserved_len
			_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(line)) - 1), Valid: true}

			var _hx_tmp_3 string
			_ = _hx_tmp_3
			if (_hx_reserved_len).Valid != false {
				_hx_tmp_3 = (Hx_Field_go_haxe_hxstring_substr(line, 0, _hx_reserved_len))
			} else {
				var _hx_tmp_4 string
				_ = _hx_tmp_4
				_hx_tmp_4 = line
				var _hx_tmp_5 string
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_4
				_hx_tmp_3 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, 0, HxNullable[int32]{}))
			}

			line = _hx_tmp_3
		}

		var i1 int32
		_ = i1
		i1 = 0
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = int32(utf8.RuneCountInString(line))
		for i1 < _hx_reserved_len {
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, i1))
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			if (c).Valid == false {
				_hx_tmp_6 = -1
			} else {
				_hx_tmp_6 = (c).Value
			}

			var c1 int32
			_ = c1
			c1 = _hx_tmp_6
			if (c1 != 32) && (c1 != 9) {
				break
			}

			i1++
		}

		var trimmed string
		_ = trimmed
		if i1 == 0 {
			trimmed = line
		} else {
			var tmp__hx_reserved_len_1 HxNullable[int32]
			_ = tmp__hx_reserved_len_1
			tmp__hx_reserved_len_1 = HxNullable[int32]{}

			var _hx_tmp_7 string
			_ = _hx_tmp_7
			if (tmp__hx_reserved_len_1).Valid != false {
				_hx_tmp_7 = (Hx_Field_go_haxe_hxstring_substr(line, i1, tmp__hx_reserved_len_1))
			} else {
				var _hx_tmp_8 string
				_ = _hx_tmp_8
				_hx_tmp_8 = line
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = i1
				var _hx_tmp_10 string
				_ = _hx_tmp_10
				_hx_tmp_10 = _hx_tmp_8
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_9
				_hx_tmp_7 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_10, _hx_tmp_11, HxNullable[int32]{}))
			}

			trimmed = _hx_tmp_7
		}

		if trimmed == "" {
			if current != nil {
				_ = (Hx_Array_Push[*Hx_Obj_loreline_translation__potranslation_poentry](entries, current))
				current = nil
			}

			i++
			continue
		}

		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(trimmed, 0))
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		if (c).Valid == false {
			_hx_tmp_12 = -1
		} else {
			_hx_tmp_12 = (c).Value
		}

		if (_hx_tmp_12) == 35 {
			i++
			continue
		}

		if current == nil {
			current = Hx_Obj_loreline_translation__potranslation_poentry_CreateInstance()
		}

		var keyword string
		_ = keyword
		keyword = Hx_Field_loreline_translation_potranslation_readKeyword(trimmed)
		if keyword == HxStringNull {
			i++
			continue
		}

		var _hx_tmp_13 string
		_ = _hx_tmp_13
		_hx_tmp_13 = trimmed
		var firstString string
		_ = firstString
		firstString = Hx_Field_loreline_translation_potranslation_readQuotedString(_hx_tmp_13, int32(utf8.RuneCountInString(keyword)))
		var _hx_tmp_14 string
		_ = _hx_tmp_14
		if firstString != HxStringNull {
			_hx_tmp_14 = firstString
		} else {
			_hx_tmp_14 = ""
		}

		var combined string
		_ = combined
		combined = _hx_tmp_14
		var j int32
		_ = j
		j = (i + 1)
		for {
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = j
			if !(_hx_tmp_15 < lines.Len()) {
				break
			}

			var next string
			_ = next
			next = lines.Get(j)
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if int32(utf8.RuneCountInString(next)) > 0 {
				var _hx_tmp_16 string
				_ = _hx_tmp_16
				_hx_tmp_16 = next
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_16, (int32(utf8.RuneCountInString(next)) - 1)))
				var _hx_tmp_17 int32
				_ = _hx_tmp_17
				if (tmp_c_1).Valid == false {
					_hx_tmp_17 = -1
				} else {
					_hx_tmp_17 = (tmp_c_1).Value
				}

				tmp_tmp_1 = ((_hx_tmp_17) == 13)
			} else {
				tmp_tmp_1 = false
			}

			if tmp_tmp_1 {
				var tmp__hx_reserved_len_1 HxNullable[int32]
				_ = tmp__hx_reserved_len_1
				tmp__hx_reserved_len_1 = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(next)) - 1), Valid: true}

				var _hx_tmp_18 string
				_ = _hx_tmp_18
				if (tmp__hx_reserved_len_1).Valid != false {
					_hx_tmp_18 = (Hx_Field_go_haxe_hxstring_substr(next, 0, tmp__hx_reserved_len_1))
				} else {
					var _hx_tmp_19 string
					_ = _hx_tmp_19
					_hx_tmp_19 = next
					var _hx_tmp_20 string
					_ = _hx_tmp_20
					_hx_tmp_20 = _hx_tmp_19
					_hx_tmp_18 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_20, 0, HxNullable[int32]{}))
				}

				next = _hx_tmp_18
			}

			var tmp_i_1 int32
			_ = tmp_i_1
			tmp_i_1 = 0
			var tmp__hx_reserved_len_1 int32
			_ = tmp__hx_reserved_len_1
			tmp__hx_reserved_len_1 = int32(utf8.RuneCountInString(next))
			for tmp_i_1 < tmp__hx_reserved_len_1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(next, tmp_i_1))
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
				if (c1 != 32) && (c1 != 9) {
					break
				}

				tmp_i_1++
			}

			var nextTrim string
			_ = nextTrim
			if tmp_i_1 == 0 {
				nextTrim = next
			} else {
				var tmp__hx_reserved_len_2 HxNullable[int32]
				_ = tmp__hx_reserved_len_2
				tmp__hx_reserved_len_2 = HxNullable[int32]{}

				var _hx_tmp_22 string
				_ = _hx_tmp_22
				if (tmp__hx_reserved_len_2).Valid != false {
					_hx_tmp_22 = (Hx_Field_go_haxe_hxstring_substr(next, tmp_i_1, tmp__hx_reserved_len_2))
				} else {
					var _hx_tmp_23 string
					_ = _hx_tmp_23
					_hx_tmp_23 = next
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = tmp_i_1
					var _hx_tmp_25 string
					_ = _hx_tmp_25
					_hx_tmp_25 = _hx_tmp_23
					var _hx_tmp_26 int32
					_ = _hx_tmp_26
					_hx_tmp_26 = _hx_tmp_24
					_hx_tmp_22 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_25, _hx_tmp_26, HxNullable[int32]{}))
				}

				nextTrim = _hx_tmp_22
			}

			var tmp1 bool
			_ = tmp1
			if int32(utf8.RuneCountInString(nextTrim)) != 0 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(nextTrim, 0))
				var _hx_tmp_27 int32
				_ = _hx_tmp_27
				if (tmp_c_1).Valid == false {
					_hx_tmp_27 = -1
				} else {
					_hx_tmp_27 = (tmp_c_1).Value
				}

				tmp1 = ((_hx_tmp_27) != 34)
			} else {
				tmp1 = true
			}

			if tmp1 {
				break
			}

			var more string
			_ = more
			more = Hx_Field_loreline_translation_potranslation_readQuotedString(nextTrim, 0)
			if more != HxStringNull {
				combined = Hx_Field_go_haxe_hxstring_concat(combined, more)
			}

			j++
		}

		switch keyword {
		case "msgctxt":
			{
				current.Hx_Field_msgctxt = combined
			}

		case "msgid":
			{
				current.Hx_Field_msgid = combined
			}

		case "msgstr":
			{
				current.Hx_Field_msgstr = combined
			}

		default:
			{
				if keyword == "msgstr[0]" {
					current.Hx_Field_msgstr = combined
				}
			}
		}

		i = j
	}

	if current != nil {
		_ = (Hx_Array_Push[*Hx_Obj_loreline_translation__potranslation_poentry](entries, current))
	}

	return entries
}

func Hx_Field_loreline_translation_potranslation_readKeyword(line string) string {
	var end int32
	_ = end
	end = 0
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(line))
	for end < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, end))
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (c).Valid == false {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_0
		if (c1 == 32) || (c1 == 9) {
			break
		}

		end++
	}

	if end == 0 {
		return HxStringNull
	}

	if (HxNullable[int32]{Value: end, Valid: true}).Valid != false {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = line
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 0, HxNullable[int32]{Value: end, Valid: true}))
	} else {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = line
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, HxNullable[int32]{}))
	}
}

func Hx_Field_loreline_translation_potranslation_readQuotedString(line string, from int32) string {
	var i int32
	_ = i
	i = from
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(line))
	for i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, i))
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (c).Valid == false {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_0
		if (c1 != 32) && (c1 != 9) {
			break
		}

		i++
	}

	var tmp bool
	_ = tmp
	if i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, i))
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (c).Valid == false {
			_hx_tmp_1 = -1
		} else {
			_hx_tmp_1 = (c).Value
		}

		tmp = ((_hx_tmp_1) != 34)
	} else {
		tmp = true
	}

	if tmp {
		return HxStringNull
	}

	i++
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	for i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, i))
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (c).Valid == false {
			_hx_tmp_2 = -1
		} else {
			_hx_tmp_2 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_2
		if (c1 == 92) && ((i + 1) < _hx_reserved_len) {
			var tmp_c_1 HxNullable[int32]
			_ = tmp_c_1
			tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(line, (i + 1)))
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			if (tmp_c_1).Valid == false {
				_hx_tmp_3 = -1
			} else {
				_hx_tmp_3 = (tmp_c_1).Value
			}

			var esc int32
			_ = esc
			esc = _hx_tmp_3
			switch esc {
			case 34:
				{
					buf.VTable.Hx_Field_addChar(34)
				}

			case 92:
				{
					buf.VTable.Hx_Field_addChar(92)
				}

			case 110:
				{
					buf.VTable.Hx_Field_addChar(10)
				}

			case 114:
				{
					buf.VTable.Hx_Field_addChar(13)
				}

			case 116:
				{
					buf.VTable.Hx_Field_addChar(9)
				}

			default:
				{
					buf.VTable.Hx_Field_addChar(esc)
				}
			}

			i = (i + 2)
		} else {
			if c1 == 34 {
				return buf.VTable.Hx_Field_toString()
			} else {
				buf.VTable.Hx_Field_addChar(c1)
				i++
			}
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Invalid PO: unterminated quoted string", nil)))))
}

func Hx_Field_loreline_translation_potranslation_trimLeft(s string) string {
	var i int32
	_ = i
	i = 0
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(s))
	for i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(s, i))
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (c).Valid == false {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_0
		if (c1 != 32) && (c1 != 9) {
			break
		}

		i++
	}

	if i == 0 {
		return s
	} else {
		var tmp__hx_reserved_len_1 HxNullable[int32]
		_ = tmp__hx_reserved_len_1
		tmp__hx_reserved_len_1 = HxNullable[int32]{}

		if (tmp__hx_reserved_len_1).Valid != false {
			return (Hx_Field_go_haxe_hxstring_substr(s, i, tmp__hx_reserved_len_1))
		} else {
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = s
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = i
			return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{}))
		}
	}
}

func Hx_Field_loreline_translation_potranslation_escapeLineForComment(s string) string {
	return (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\n")), " ")), "\r")), ""))
}

func Hx_Field_loreline_translation_potranslation_escapeLineForBody(s string) string {
	return s
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation__potranslation_poentry() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation__potranslation_poentry() {
}

type Hx_Obj_VTable_loreline_translation__potranslation_poentry interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation__potranslation_poentry struct {
	VTable           Hx_Obj_VTable_loreline_translation__potranslation_poentry `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_msgctxt string
	Hx_Field_msgid   string
	Hx_Field_msgstr  string
}

func Hx_Obj_loreline_translation__potranslation_poentry_CreateEmptyInstance() *Hx_Obj_loreline_translation__potranslation_poentry {
	obj := &Hx_Obj_loreline_translation__potranslation_poentry{}
	obj.VTable = obj
	obj.Hx_Field_msgctxt = HxDefault[string]()
	obj.Hx_Field_msgid = HxDefault[string]()
	obj.Hx_Field_msgstr = HxDefault[string]()
	return obj
}

func Hx_Obj_loreline_translation__potranslation_poentry_CreateInstance() *Hx_Obj_loreline_translation__potranslation_poentry {
	obj := Hx_Obj_loreline_translation__potranslation_poentry_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_translation__potranslation_poentry) Hx_New() {
	{
	}

}

var Hx_Obj_loreline_translation__potranslation_poentry_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation._PoTranslation.PoEntry",
	HxMakeArray[string](),
	HxMakeArray[string]("msgctxt", "msgid", "msgstr"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation__potranslation_poentry_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_translation__potranslation_poentry_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation__potranslation_poentry_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_translation__potranslation_poentry) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation__potranslation_poentry_RTTI
}

func (this *Hx_Obj_loreline_translation__potranslation_poentry) Hx_As_Hx_Obj_loreline_translation__potranslation_poentry() *Hx_Obj_loreline_translation__potranslation_poentry {
	return this
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_translation__potranslation_poentry]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
