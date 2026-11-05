package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation_translationformats() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation_translationformats() {
	Hx_Field_loreline_translation_translationformats_formats = Hx_Init_Hx_Field_loreline_translation_translationformats_formats()
	Hx_Field_loreline_translation_translationformats_enabled = Hx_Init_Hx_Field_loreline_translation_translationformats_enabled()
}

type Hx_Obj_VTable_loreline_translation_translationformats interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation_translationformats struct {
	VTable Hx_Obj_VTable_loreline_translation_translationformats `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_translation_translationformats_CreateEmptyInstance() *Hx_Obj_loreline_translation_translationformats {
	obj := &Hx_Obj_loreline_translation_translationformats{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_translation_translationformats_CreateInstance() *Hx_Obj_loreline_translation_translationformats {
	obj := Hx_Obj_loreline_translation_translationformats_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_translation_translationformats_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation.TranslationFormats",
	HxMakeArray[string]("formats", "convert", "enabled", "translationFormat", "anyEnabled", "wrap", "endsWith"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation_translationformats_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_translation_translationformats_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation_translationformats_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "formats":
			return &Hx_Field_loreline_translation_translationformats_formats
		case "convert":
			return Hx_Field_loreline_translation_translationformats_convert
		case "enabled":
			return &Hx_Field_loreline_translation_translationformats_enabled
		case "translationFormat":
			return Hx_Field_loreline_translation_translationformats_translationFormat
		case "anyEnabled":
			return Hx_Field_loreline_translation_translationformats_anyEnabled
		case "wrap":
			return Hx_Field_loreline_translation_translationformats_wrap
		case "endsWith":
			return Hx_Field_loreline_translation_translationformats_endsWith
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_translation_translationformats) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation_translationformats_RTTI
}

func (this *Hx_Obj_loreline_translation_translationformats) Hx_As_Hx_Obj_loreline_translation_translationformats() *Hx_Obj_loreline_translation_translationformats {
	return this
}

var Hx_Field_loreline_translation_translationformats_formats HxArray[any]

func Hx_Init_Hx_Field_loreline_translation_translationformats_formats() HxArray[any] {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = any(map[string]any{"name": ((any)("po")), "ext": ((any)(".po"))})
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = any(map[string]any{"name": ((any)("xliff")), "ext": ((any)(".xliff"))})
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = any(map[string]any{"name": ((any)("xliff")), "ext": ((any)(".xlf"))})
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = any(map[string]any{"name": ((any)("csv")), "ext": ((any)(".csv"))})
	return HxMakeArray[any](_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, any(map[string]any{"name": ((any)("csv")), "ext": ((any)(".tsv"))})).Dyn()
}

func Hx_Field_loreline_translation_translationformats_convert(ext string, content string, locale string) string {
	switch ext {
	case ".csv":
		{
			return Hx_Field_loreline_translation_csvtranslation_toLoreline(content, locale)
		}

	case ".po":
		{
			return Hx_Field_loreline_translation_potranslation_toLoreline(content, locale)
		}

	case ".tsv":
		{
			return Hx_Field_loreline_translation_csvtranslation_tsvToLoreline(content, locale)
		}

	case ".xlf", ".xliff":
		{
			return Hx_Field_loreline_translation_xlifftranslation_toLoreline(content, locale)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_exception_CreateInstance(Hx_Field_go_haxe_hxstring_concat("Unknown translation format extension: ", ext), nil, nil)))))
		}
	}
}

var Hx_Field_loreline_translation_translationformats_enabled Hx_Typedef_map

func Hx_Init_Hx_Field_loreline_translation_translationformats_enabled() Hx_Typedef_map {
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_369 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_369 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_369, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_369.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func Hx_Field_loreline_translation_translationformats_translationFormat(name string, value bool) {
	func() func(_hx_param_key string, _hx_param_value bool) {
		var _hx_closure_src_840 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_840
		_hx_closure_src_840 = Hx_Field_loreline_translation_translationformats_enabled.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value bool) {
			_hx_closure_src_840(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(name, value)
}

func Hx_Field_loreline_translation_translationformats_anyEnabled() bool {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = Hx_Field_loreline_translation_translationformats_formats
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var f any
			_ = f
			f = _g1.Get(_g)
			_g++
			var hx_cmp_23 HxNullable[bool]
			_ = hx_cmp_23
			hx_cmp_23 = func() func(_hx_param_key string) HxNullable[bool] {
				var _hx_closure_src_842 func(_hx_param_key any) any
				_ = _hx_closure_src_842
				_hx_closure_src_842 = Hx_Field_loreline_translation_translationformats_enabled.VTable.Hx_Field_get
				return func(_hx_param_key string) HxNullable[bool] {
					var hx_nullable_t71 any
					_ = hx_nullable_t71
					hx_nullable_t71 = _hx_closure_src_842(((any)(_hx_param_key)))
					var _hx_tmp_1 HxNullable[bool]
					_ = _hx_tmp_1
					if hx_nullable_t71 == nil {
						_hx_tmp_1 = HxNullable[bool]{}
					} else {
						_hx_tmp_1 = HxNullable[bool]{Value: ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(hx_nullable_t71))), Valid: true}
					}

					return _hx_tmp_1
				}
			}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name")))
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if (hx_cmp_23).Valid == true {
				_hx_tmp_2 = ((hx_cmp_23).Value == true)
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				return true
			}
		}
	}

	return false
}

func Hx_Field_loreline_translation_translationformats_wrap(underlying Hx_Typedef_loreline_importsfilehandler, locale string) *Hx_Obj_loreline_translation_wrappedfilehandler {
	if !Hx_Field_loreline_translation_translationformats_anyEnabled() {
		return Hx_Obj_loreline_translation_wrappedfilehandler_CreateInstance(underlying, ((func() *Hx_Obj_loreline_error)(func() *Hx_Obj_loreline_error {
			return nil
		})))
	}

	var captured *Hx_Obj_loreline_error
	_ = captured
	var tryNext func(stem string, withLocaleSuffix bool, index int32, cb func(content string))
	_ = tryNext
	tryNext = func(stem string, withLocaleSuffix bool, index int32, cb func(content string)) {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = index
		if _hx_tmp_0 >= Hx_Field_go_haxe_hxdynamic_getArrayLength(Hx_Field_loreline_translation_translationformats_formats) {
			if withLocaleSuffix {
				tryNext(stem, false, 0, cb)
			} else {
				cb(HxStringNull)
			}

			return
		}

		var f any
		_ = f
		f = Hx_Field_loreline_translation_translationformats_formats.Get(index)
		var hx_cmp_24 HxNullable[bool]
		_ = hx_cmp_24
		hx_cmp_24 = func() func(_hx_param_key string) HxNullable[bool] {
			var _hx_closure_src_844 func(_hx_param_key any) any
			_ = _hx_closure_src_844
			_hx_closure_src_844 = Hx_Field_loreline_translation_translationformats_enabled.VTable.Hx_Field_get
			return func(_hx_param_key string) HxNullable[bool] {
				var hx_nullable_t73 any
				_ = hx_nullable_t73
				hx_nullable_t73 = _hx_closure_src_844(((any)(_hx_param_key)))
				var _hx_tmp_1 HxNullable[bool]
				_ = _hx_tmp_1
				if hx_nullable_t73 == nil {
					_hx_tmp_1 = HxNullable[bool]{}
				} else {
					_hx_tmp_1 = HxNullable[bool]{Value: ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(hx_nullable_t73))), Valid: true}
				}

				return _hx_tmp_1
			}
		}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "name")))
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if (hx_cmp_24).Valid == false {
			_hx_tmp_2 = true
		} else {
			_hx_tmp_2 = ((hx_cmp_24).Value != true)
		}

		if _hx_tmp_2 {
			tryNext(stem, withLocaleSuffix, (index + 1), cb)
			return
		}

		var _hx_tmp_3 string
		_ = _hx_tmp_3
		if withLocaleSuffix {
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = ((any)(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(stem, "."), locale)))
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = _hx_tmp_4
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_5, Hx_Field_go_haxe_hxdynamic_getField(f, "ext")))
		} else {
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = ((any)(stem))
			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_7, Hx_Field_go_haxe_hxdynamic_getField(f, "ext")))
		}

		var candidate string
		_ = candidate
		candidate = _hx_tmp_3
		var _hx_tmp_8 string
		_ = _hx_tmp_8
		_hx_tmp_8 = candidate
		underlying(_hx_tmp_8, func(content string) {
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
										captured = Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Failed to load translation file \"", candidate), "\": "), e.Hx_Field_message), nil)
										tryNext(stem, withLocaleSuffix, (index + 1), cb)
									}
								} else {
									panic(_hx_ex)
								}
							}
						}()
						{
							cb(Hx_Field_loreline_translation_translationformats_convert(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(f, "ext")), content, locale))
						}
					}()
					if hx_try_state == 1 {
						return
					}
				}
			} else {
				tryNext(stem, withLocaleSuffix, (index + 1), cb)
			}
		})
	}

	var wrapped Hx_Typedef_loreline_importsfilehandler
	_ = wrapped
	wrapped = ((Hx_Typedef_loreline_importsfilehandler)(func(path string, cb func(data string)) {
		var _hx_tmp_9 string
		_ = _hx_tmp_9
		_hx_tmp_9 = path
		underlying(_hx_tmp_9, func(content string) {
			if content != HxStringNull {
				cb(content)
				return
			}

			var isLorPath bool
			_ = isLorPath
			if path == HxStringNull {
				isLorPath = false
			} else {
				var n int32
				_ = n
				n = int32(utf8.RuneCountInString(path))
				if n < 4 {
					isLorPath = false
				} else {
					var isLorPath1 bool
					_ = isLorPath1
					var isLorPath2 bool
					_ = isLorPath2
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					if (c).Valid == false {
						_hx_tmp_10 = -1
					} else {
						_hx_tmp_10 = (c).Value
					}

					if (_hx_tmp_10) == 46 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
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
						isLorPath2 = ((c1 == 108) || (c1 == 76))
					} else {
						isLorPath2 = false
					}

					if isLorPath2 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
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
						isLorPath1 = ((c1 == 111) || (c1 == 79))
					} else {
						isLorPath1 = false
					}

					if isLorPath1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						if (tmp_c_1).Valid == false {
							_hx_tmp_13 = -1
						} else {
							_hx_tmp_13 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_13
						isLorPath = ((c1 == 114) || (c1 == 82))
					} else {
						isLorPath = false
					}
				}
			}

			if !isLorPath {
				cb(HxStringNull)
				return
			}

			var _hx_reserved_len HxNullable[int32]
			_ = _hx_reserved_len
			_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 4), Valid: true}

			var _hx_tmp_14 string
			_ = _hx_tmp_14
			if (_hx_reserved_len).Valid != false {
				_hx_tmp_14 = (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
			} else {
				var _hx_tmp_15 string
				_ = _hx_tmp_15
				_hx_tmp_15 = path
				var _hx_tmp_16 string
				_ = _hx_tmp_16
				_hx_tmp_16 = _hx_tmp_15
				_hx_tmp_14 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_16, 0, HxNullable[int32]{}))
			}

			var withoutLorExt string
			_ = withoutLorExt
			withoutLorExt = _hx_tmp_14
			var localeSuffix string
			_ = localeSuffix
			localeSuffix = Hx_Field_go_haxe_hxstring_concat(".", locale)
			var stem string
			_ = stem
			var sLen int32
			_ = sLen
			sLen = int32(utf8.RuneCountInString(withoutLorExt))
			var suffixLen int32
			_ = suffixLen
			suffixLen = int32(utf8.RuneCountInString(localeSuffix))
			var stem1 bool
			_ = stem1
			if suffixLen > sLen {
				stem1 = false
			} else {
				var pos int32
				_ = pos
				pos = (sLen - suffixLen)
				var _hx_tmp_17 string
				_ = _hx_tmp_17
				if (HxNullable[int32]{Value: suffixLen, Valid: true}).Valid != false {
					var _hx_tmp_18 string
					_ = _hx_tmp_18
					_hx_tmp_18 = withoutLorExt
					var _hx_tmp_19 int32
					_ = _hx_tmp_19
					_hx_tmp_19 = pos
					var _hx_tmp_20 string
					_ = _hx_tmp_20
					_hx_tmp_20 = _hx_tmp_18
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					_hx_tmp_21 = _hx_tmp_19
					_hx_tmp_17 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_20, _hx_tmp_21, HxNullable[int32]{Value: suffixLen, Valid: true}))
				} else {
					var _hx_tmp_22 string
					_ = _hx_tmp_22
					_hx_tmp_22 = withoutLorExt
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = pos
					var _hx_tmp_24 string
					_ = _hx_tmp_24
					_hx_tmp_24 = _hx_tmp_22
					var _hx_tmp_25 int32
					_ = _hx_tmp_25
					_hx_tmp_25 = _hx_tmp_23
					_hx_tmp_17 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_24, _hx_tmp_25, HxNullable[int32]{}))
				}

				stem1 = ((_hx_tmp_17) == localeSuffix)
			}

			if stem1 {
				var _hx_tmp_26 int32
				_ = _hx_tmp_26
				_hx_tmp_26 = int32(utf8.RuneCountInString(withoutLorExt))
				var tmp__hx_reserved_len_1 HxNullable[int32]
				_ = tmp__hx_reserved_len_1
				tmp__hx_reserved_len_1 = HxNullable[int32]{Value: (_hx_tmp_26 - int32(utf8.RuneCountInString(localeSuffix))), Valid: true}

				var _hx_tmp_27 string
				_ = _hx_tmp_27
				if (tmp__hx_reserved_len_1).Valid != false {
					_hx_tmp_27 = (Hx_Field_go_haxe_hxstring_substr(withoutLorExt, 0, tmp__hx_reserved_len_1))
				} else {
					var _hx_tmp_28 string
					_ = _hx_tmp_28
					_hx_tmp_28 = withoutLorExt
					var _hx_tmp_29 string
					_ = _hx_tmp_29
					_hx_tmp_29 = _hx_tmp_28
					_hx_tmp_27 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_29, 0, HxNullable[int32]{}))
				}

				stem = _hx_tmp_27
			} else {
				stem = withoutLorExt
			}

			tryNext(stem, true, 0, cb)
		})
	}))
	return Hx_Obj_loreline_translation_wrappedfilehandler_CreateInstance(wrapped, ((func() *Hx_Obj_loreline_error)(func() *Hx_Obj_loreline_error {
		return captured
	})))
}

func Hx_Field_loreline_translation_translationformats_endsWith(s string, suffix string) bool {
	var sLen int32
	_ = sLen
	sLen = int32(utf8.RuneCountInString(s))
	var suffixLen int32
	_ = suffixLen
	suffixLen = int32(utf8.RuneCountInString(suffix))
	if suffixLen > sLen {
		return false
	}

	var pos int32
	_ = pos
	pos = (sLen - suffixLen)
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if (HxNullable[int32]{Value: suffixLen, Valid: true}).Valid != false {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = pos
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_1
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_2
		_hx_tmp_0 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{Value: suffixLen, Valid: true}))
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = s
		var _hx_tmp_6 int32
		_ = _hx_tmp_6
		_hx_tmp_6 = pos
		var _hx_tmp_7 string
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_5
		var _hx_tmp_8 int32
		_ = _hx_tmp_8
		_hx_tmp_8 = _hx_tmp_6
		_hx_tmp_0 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, _hx_tmp_8, HxNullable[int32]{}))
	}

	return ((_hx_tmp_0) == suffix)
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation_wrappedfilehandler() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation_wrappedfilehandler() {
}

type Hx_Obj_VTable_loreline_translation_wrappedfilehandler interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation_wrappedfilehandler struct {
	VTable             Hx_Obj_VTable_loreline_translation_wrappedfilehandler `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_handler   Hx_Typedef_loreline_importsfilehandler
	Hx_Field_lastError func() *Hx_Obj_loreline_error
}

func Hx_Obj_loreline_translation_wrappedfilehandler_CreateEmptyInstance() *Hx_Obj_loreline_translation_wrappedfilehandler {
	obj := &Hx_Obj_loreline_translation_wrappedfilehandler{}
	obj.VTable = obj
	obj.Hx_Field_handler = HxDefault[Hx_Typedef_loreline_importsfilehandler]()
	obj.Hx_Field_lastError = HxDefault[func() *Hx_Obj_loreline_error]()
	return obj
}

func Hx_Obj_loreline_translation_wrappedfilehandler_CreateInstance(handler Hx_Typedef_loreline_importsfilehandler, lastError func() *Hx_Obj_loreline_error) *Hx_Obj_loreline_translation_wrappedfilehandler {
	obj := Hx_Obj_loreline_translation_wrappedfilehandler_CreateEmptyInstance()
	obj.Hx_New(handler, lastError)
	return obj
}

func (this *Hx_Obj_loreline_translation_wrappedfilehandler) Hx_New(handler Hx_Typedef_loreline_importsfilehandler, lastError func() *Hx_Obj_loreline_error) {
	{
		this.Hx_Field_handler = handler
		this.Hx_Field_lastError = lastError
	}

}

var Hx_Obj_loreline_translation_wrappedfilehandler_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation.WrappedFileHandler",
	HxMakeArray[string](),
	HxMakeArray[string]("handler", "lastError"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation_wrappedfilehandler_CreateInstance(HxConvert[Hx_Typedef_loreline_importsfilehandler](params.Get(0)), HxConvert[func() *Hx_Obj_loreline_error](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_translation_wrappedfilehandler_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation_wrappedfilehandler_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_translation_wrappedfilehandler) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation_wrappedfilehandler_RTTI
}

func (this *Hx_Obj_loreline_translation_wrappedfilehandler) Hx_As_Hx_Obj_loreline_translation_wrappedfilehandler() *Hx_Obj_loreline_translation_wrappedfilehandler {
	return this
}

func init() {
	hxRegisterArrayConstructor[any]()
}
