package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_path() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_path() {
}

type Hx_Obj_VTable_haxe_io_path interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_path struct {
	VTable             Hx_Obj_VTable_haxe_io_path `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_dir       HxNullable[string]
	Hx_Field_file      string
	Hx_Field_ext       HxNullable[string]
	Hx_Field_backslash bool
}

func Hx_Obj_haxe_io_path_CreateEmptyInstance() *Hx_Obj_haxe_io_path {
	obj := &Hx_Obj_haxe_io_path{}
	obj.VTable = obj
	obj.Hx_Field_dir = HxDefault[HxNullable[string]]()
	obj.Hx_Field_file = HxDefault[string]()
	obj.Hx_Field_ext = HxDefault[HxNullable[string]]()
	obj.Hx_Field_backslash = HxDefault[bool]()
	return obj
}

func Hx_Obj_haxe_io_path_CreateInstance(path string) *Hx_Obj_haxe_io_path {
	obj := Hx_Obj_haxe_io_path_CreateEmptyInstance()
	obj.Hx_New(path)
	return obj
}

func (this *Hx_Obj_haxe_io_path) Hx_New(path string) {
	{
		switch path {
		case ".", "..":
			{
				this.Hx_Field_dir = HxNullableString(path)
				this.Hx_Field_file = ""
				return
			}

		}

		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = path
		var c1 int32
		_ = c1
		c1 = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_0, "/", HxNullable[int32]{}))
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = path
		var c2 int32
		_ = c2
		c2 = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_1, "\\", HxNullable[int32]{}))
		if c1 < c2 {
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = path
			this.Hx_Field_dir = HxNullableString((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, HxNullable[int32]{Value: c2, Valid: true})))
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = path
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = (c2 + 1)
			path = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{}))
			this.Hx_Field_backslash = true
		} else {
			if c2 < c1 {
				var _hx_tmp_5 string
				_ = _hx_tmp_5
				_hx_tmp_5 = path
				this.Hx_Field_dir = HxNullableString((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, 0, HxNullable[int32]{Value: c1, Valid: true})))
				var _hx_tmp_6 string
				_ = _hx_tmp_6
				_hx_tmp_6 = path
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = (c1 + 1)
				path = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_6, _hx_tmp_7, HxNullable[int32]{}))
			} else {
				this.Hx_Field_dir = HxNullable[string]{Value: HxStringNull}
			}
		}

		var _hx_tmp_8 string
		_ = _hx_tmp_8
		_hx_tmp_8 = path
		var cp int32
		_ = cp
		cp = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_8, ".", HxNullable[int32]{}))
		if cp != -1 {
			var _hx_tmp_9 string
			_ = _hx_tmp_9
			_hx_tmp_9 = path
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = (cp + 1)
			this.Hx_Field_ext = HxNullableString((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_9, _hx_tmp_10, HxNullable[int32]{})))
			var _hx_tmp_11 string
			_ = _hx_tmp_11
			_hx_tmp_11 = path
			this.Hx_Field_file = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_11, 0, HxNullable[int32]{Value: cp, Valid: true}))
		} else {
			this.Hx_Field_ext = HxNullable[string]{Value: HxStringNull}

			this.Hx_Field_file = path
		}
	}

}

var Hx_Obj_haxe_io_path_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.Path",
	HxMakeArray[string]("withoutExtension", "withoutDirectory", "directory", "extension", "withExtension", "join", "normalize", "addTrailingSlash", "removeTrailingSlashes", "isAbsolute", "unescape", "escape"),
	HxMakeArray[string]("dir", "file", "ext", "backslash", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_path_CreateInstance(HxConvert[string](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_io_path_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_path_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "withoutExtension":
			return Hx_Field_haxe_io_path_withoutExtension
		case "withoutDirectory":
			return Hx_Field_haxe_io_path_withoutDirectory
		case "directory":
			return Hx_Field_haxe_io_path_directory
		case "extension":
			return Hx_Field_haxe_io_path_extension
		case "withExtension":
			return Hx_Field_haxe_io_path_withExtension
		case "join":
			return Hx_Field_haxe_io_path_join
		case "normalize":
			return Hx_Field_haxe_io_path_normalize
		case "addTrailingSlash":
			return Hx_Field_haxe_io_path_addTrailingSlash
		case "removeTrailingSlashes":
			return Hx_Field_haxe_io_path_removeTrailingSlashes
		case "isAbsolute":
			return Hx_Field_haxe_io_path_isAbsolute
		case "unescape":
			return Hx_Field_haxe_io_path_unescape
		case "escape":
			return Hx_Field_haxe_io_path_escape
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_io_path) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_path_RTTI
}

func (this *Hx_Obj_haxe_io_path) Hx_As_Hx_Obj_haxe_io_path() *Hx_Obj_haxe_io_path {
	return this
}

func (this *Hx_Obj_haxe_io_path) Hx_Field_toString() string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if (this.Hx_Field_dir).Valid == false {
		_hx_tmp_0 = ""
	} else {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = (this.Hx_Field_dir).Value
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		if this.Hx_Field_backslash {
			_hx_tmp_2 = "\\"
		} else {
			_hx_tmp_2 = "/"
		}

		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_1
		_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat(_hx_tmp_3, (_hx_tmp_2))
	}

	var _hx_tmp_4 string
	_ = _hx_tmp_4
	if (this.Hx_Field_ext).Valid == false {
		_hx_tmp_4 = ""
	} else {
		_hx_tmp_4 = Hx_Field_go_haxe_hxstring_concat(".", (this.Hx_Field_ext).Value)
	}

	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat((_hx_tmp_0), this.Hx_Field_file), (_hx_tmp_4))
}

func Hx_Field_haxe_io_path_withoutExtension(path string) string {
	var s *Hx_Obj_haxe_io_path
	_ = s
	s = Hx_Obj_haxe_io_path_CreateInstance(path)
	s.Hx_Field_ext = HxNullable[string]{Value: HxStringNull}

	return s.VTable.Hx_Field_toString()
}

func Hx_Field_haxe_io_path_withoutDirectory(path string) string {
	var s *Hx_Obj_haxe_io_path
	_ = s
	s = Hx_Obj_haxe_io_path_CreateInstance(path)
	s.Hx_Field_dir = HxNullable[string]{Value: HxStringNull}

	return s.VTable.Hx_Field_toString()
}

func Hx_Field_haxe_io_path_directory(path string) string {
	var s *Hx_Obj_haxe_io_path
	_ = s
	s = Hx_Obj_haxe_io_path_CreateInstance(path)
	if (s.Hx_Field_dir).Valid == false {
		return ""
	}

	return (s.Hx_Field_dir).Value
}

func Hx_Field_haxe_io_path_extension(path string) string {
	var s *Hx_Obj_haxe_io_path
	_ = s
	s = Hx_Obj_haxe_io_path_CreateInstance(path)
	if (s.Hx_Field_ext).Valid == false {
		return ""
	}

	return (s.Hx_Field_ext).Value
}

func Hx_Field_haxe_io_path_withExtension(path string, ext HxNullable[string]) string {
	var s *Hx_Obj_haxe_io_path
	_ = s
	s = Hx_Obj_haxe_io_path_CreateInstance(path)
	s.Hx_Field_ext = ext
	return s.VTable.Hx_Field_toString()
}

func Hx_Field_haxe_io_path_join(paths HxArray[string]) string {
	var _g HxArray[string]
	_ = _g
	_g = HxMakeArray[string]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[string]
		_ = _g2
		_g2 = paths
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var v string
			_ = v
			v = _g2.Get(_g1)
			_g1++
			if (v != HxStringNull) && (v != "") {
				_ = (Hx_Array_Push[string](_g, v))
			}
		}
	}

	var paths1 HxArray[string]
	_ = paths1
	paths1 = _g
	if paths1.Len() == 0 {
		return ""
	}

	var path string
	_ = path
	path = paths1.Get(0)
	{
		var _g3 int32
		_ = _g3
		_g3 = 1
		var _g4 int32
		_ = _g4
		_g4 = paths1.Len()
		for _g3 < _g4 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g3
			_g3 = (_g3 + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			path = Hx_Field_haxe_io_path_addTrailingSlash(path)
			path = Hx_Field_go_haxe_hxstring_concat(path, paths1.Get(i))
		}
	}

	return Hx_Field_haxe_io_path_normalize(path)
}

func Hx_Field_haxe_io_path_normalize(path string) string {
	var slash string
	_ = slash
	slash = "/"
	path = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(path, "\\")), slash))
	if path == slash {
		return slash
	}

	var target HxArray[string]
	_ = target
	target = HxMakeArray[string]()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = (Hx_Field_go_haxe_hxstring_split(path, slash))
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var token string
			_ = token
			token = _g1.Get(_g)
			_g++
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if token == ".." {
				_hx_tmp_2 = (target.Len() > 0)
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				var _hx_tmp_3 HxArray[string]
				_ = _hx_tmp_3
				_hx_tmp_3 = target
				var _hx_tmp_4 HxArray[string]
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_1 = (_hx_tmp_4.Get((target.Len() - 1)) != "..")
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				_ = (Hx_Array_Pop[string](target))
			} else {
				if token == "" {
					var _hx_tmp_5 bool
					_ = _hx_tmp_5
					if target.Len() > 0 {
						_hx_tmp_5 = true
					} else {
						var hx_cmp_12 HxNullable[int32]
						_ = hx_cmp_12
						hx_cmp_12 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, 0))
						var _hx_tmp_6 bool
						_ = _hx_tmp_6
						if (hx_cmp_12).Valid == true {
							_hx_tmp_6 = ((hx_cmp_12).Value == 47)
						} else {
							_hx_tmp_6 = false
						}

						_hx_tmp_5 = (_hx_tmp_6)
					}

					if _hx_tmp_5 {
						_ = (Hx_Array_Push[string](target, token))
					}
				} else {
					if token != "." {
						_ = (Hx_Array_Push[string](target, token))
					}
				}
			}
		}
	}

	var tmp string
	_ = tmp
	tmp = (Hx_Array_Join[string](target, slash))
	var acc *Hx_Obj_stringbuf
	_ = acc
	acc = Hx_Obj_stringbuf_CreateInstance()
	var colon bool
	_ = colon
	colon = false
	var slashes bool
	_ = slashes
	slashes = false
	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = int32(utf8.RuneCountInString(tmp))
		for _g2 < _g3 {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_7
			{
				var _g int32
				_ = _g
				_g = (Hx_Field_go_haxe_hxstring_charCodeAt(tmp, i)).Value
				switch _g {
				case 47:
					{
						if !colon {
							slashes = true
						} else {
							var tmp_i_1 int32
							_ = tmp_i_1
							tmp_i_1 = _g
							{
								colon = false
								if slashes {
									func() func(x string) {
										var _hx_closure_src_434 func(x any)
										_ = _hx_closure_src_434
										_hx_closure_src_434 = acc.VTable.Hx_Field_add
										return func(x string) {
											_hx_closure_src_434(((any)(x)))
										}
									}()("/")
									slashes = false
								}

								acc.VTable.Hx_Field_addChar(tmp_i_1)
							}
						}
					}

				case 58:
					{
						func() func(x string) {
							var _hx_closure_src_435 func(x any)
							_ = _hx_closure_src_435
							_hx_closure_src_435 = acc.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_435(((any)(x)))
							}
						}()(":")
						colon = true
					}

				default:
					{
						var tmp_i_1 int32
						_ = tmp_i_1
						tmp_i_1 = _g
						{
							colon = false
							if slashes {
								func() func(x string) {
									var _hx_closure_src_436 func(x any)
									_ = _hx_closure_src_436
									_hx_closure_src_436 = acc.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_436(((any)(x)))
									}
								}()("/")
								slashes = false
							}

							acc.VTable.Hx_Field_addChar(tmp_i_1)
						}
					}
				}
			}
		}
	}

	return acc.VTable.Hx_Field_toString()
}

func Hx_Field_haxe_io_path_addTrailingSlash(path string) string {
	if int32(utf8.RuneCountInString(path)) == 0 {
		return "/"
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = path
	var c1 int32
	_ = c1
	c1 = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_0, "/", HxNullable[int32]{}))
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = path
	var c2 int32
	_ = c2
	c2 = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_1, "\\", HxNullable[int32]{}))
	if c1 < c2 {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = c2
		if _hx_tmp_2 != (int32(utf8.RuneCountInString(path)) - 1) {
			return Hx_Field_go_haxe_hxstring_concat(path, "\\")
		} else {
			return path
		}
	} else {
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = c1
		if _hx_tmp_3 != (int32(utf8.RuneCountInString(path)) - 1) {
			return Hx_Field_go_haxe_hxstring_concat(path, "/")
		} else {
			return path
		}
	}
}

func Hx_Field_haxe_io_path_removeTrailingSlashes(path string) string {
hx_label_0:
	for {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = path
		var _g HxNullable[int32]
		_ = _g
		_g = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, (int32(utf8.RuneCountInString(path)) - 1)))
		if (_g).Valid == false {
			break
		} else {
			switch (_g).Value {
			case 47, 92:
				{
					var _hx_tmp_1 string
					_ = _hx_tmp_1
					_hx_tmp_1 = path
					path = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 0, HxNullable[int32]{Value: -1, Valid: true}))
				}

			default:
				{
					break hx_label_0
				}
			}
		}
	}

	return path
}

func Hx_Field_haxe_io_path_isAbsolute(path string) bool {
	if Hx_Field_stringtools_startsWith(path, "/") {
		return true
	}

	if (Hx_Field_go_haxe_hxstring_charAt(path, 1)) == ":" {
		return true
	}

	if Hx_Field_stringtools_startsWith(path, "\\\\") {
		return true
	}

	return false
}

func Hx_Field_haxe_io_path_unescape(path string) string {
	var regex *Hx_Obj_ereg
	_ = regex
	regex = Hx_Obj_ereg_CreateInstance("-x([0-9][0-9])", "g")
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = path
	return regex.VTable.Hx_Field_map(_hx_tmp_0, func(regex *Hx_Obj_ereg) string {
		return Hx_Field_go_haxe_hxstring_fromCharCode((Hx_Field_std_parseInt(regex.VTable.Hx_Field_matched(1))).Value)
	})
}

func Hx_Field_haxe_io_path_escape(path string, allowSlashes__isOptional any) string {
	allowSlashes := HxDefault[bool]()
	_ = allowSlashes

	if allowSlashes__isOptional != nil {
		allowSlashes = HxConvert[bool](allowSlashes__isOptional)
	} else {
		allowSlashes = false
	}
	{
		var _hx_tmp_0 *Hx_Obj_ereg
		_ = _hx_tmp_0
		if allowSlashes {
			_hx_tmp_0 = Hx_Obj_ereg_CreateInstance("[^A-Za-z0-9_/\\\\\\.]", "g")
		} else {
			_hx_tmp_0 = Hx_Obj_ereg_CreateInstance("[^A-Za-z0-9_\\.]", "g")
		}

		var regex *Hx_Obj_ereg
		_ = regex
		regex = _hx_tmp_0
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = path
		return regex.VTable.Hx_Field_map(_hx_tmp_1, func(v *Hx_Obj_ereg) string {
			var hx_nullable_f200 HxNullable[int32]
			_ = hx_nullable_f200
			hx_nullable_f200 = (Hx_Field_go_haxe_hxstring_charCodeAt(v.VTable.Hx_Field_matched(0), 0))
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			if (hx_nullable_f200).Valid == true {
				_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_toString((hx_nullable_f200).Value)
			} else {
				_hx_tmp_2 = "null"
			}

			return ("-x" + _hx_tmp_2)
		})
	}
}

func init() {
	hxRegisterArrayConstructor[string]()
}
