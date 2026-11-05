package core

import "net/url"
import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_stringtools() {
}

func Hx_StaticInit_Hx_Obj_VTable_stringtools() {
	Hx_Field_stringtools_winMetaCharacters = Hx_Init_Hx_Field_stringtools_winMetaCharacters()
}

type Hx_Obj_VTable_stringtools interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_stringtools struct {
	VTable Hx_Obj_VTable_stringtools `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_stringtools_CreateEmptyInstance() *Hx_Obj_stringtools {
	obj := &Hx_Obj_stringtools{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_stringtools_CreateInstance() *Hx_Obj_stringtools {
	obj := Hx_Obj_stringtools_CreateEmptyInstance()
	return obj
}

var Hx_Obj_stringtools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"StringTools",
	HxMakeArray[string]("urlEncode", "urlDecode", "htmlEscape", "htmlUnescape", "contains", "startsWith", "endsWith", "isSpace", "ltrim", "rtrim", "trim", "lpad", "rpad", "replace", "hex", "fastCodeAt", "unsafeCodeAt", "iterator", "keyValueIterator", "isEof", "quoteUnixArg", "winMetaCharacters", "quoteWinArg"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_stringtools_CreateInstance()
	},
	func() any {
		return Hx_Obj_stringtools_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_stringtools_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "urlEncode":
			return Hx_Field_stringtools_urlEncode
		case "urlDecode":
			return Hx_Field_stringtools_urlDecode
		case "htmlEscape":
			return Hx_Field_stringtools_htmlEscape
		case "htmlUnescape":
			return Hx_Field_stringtools_htmlUnescape
		case "contains":
			return Hx_Field_stringtools_contains
		case "startsWith":
			return Hx_Field_stringtools_startsWith
		case "endsWith":
			return Hx_Field_stringtools_endsWith
		case "isSpace":
			return Hx_Field_stringtools_isSpace
		case "ltrim":
			return Hx_Field_stringtools_ltrim
		case "rtrim":
			return Hx_Field_stringtools_rtrim
		case "trim":
			return Hx_Field_stringtools_trim
		case "lpad":
			return Hx_Field_stringtools_lpad
		case "rpad":
			return Hx_Field_stringtools_rpad
		case "replace":
			return Hx_Field_stringtools_replace
		case "hex":
			return Hx_Field_stringtools_hex
		case "fastCodeAt":
			return Hx_Field_stringtools_fastCodeAt
		case "unsafeCodeAt":
			return Hx_Field_stringtools_unsafeCodeAt
		case "iterator":
			return Hx_Field_stringtools_iterator
		case "keyValueIterator":
			return Hx_Field_stringtools_keyValueIterator
		case "isEof":
			return Hx_Field_stringtools_isEof
		case "quoteUnixArg":
			return Hx_Field_stringtools_quoteUnixArg
		case "winMetaCharacters":
			return &Hx_Field_stringtools_winMetaCharacters
		case "quoteWinArg":
			return Hx_Field_stringtools_quoteWinArg
		}
		return nil
	}
}

func (this *Hx_Obj_stringtools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_stringtools_RTTI
}

func (this *Hx_Obj_stringtools) Hx_As_Hx_Obj_stringtools() *Hx_Obj_stringtools {
	return this
}

func Hx_Field_stringtools_urlEncode(s string) string {
	return url.QueryEscape(s)
}

func Hx_Field_stringtools_urlDecode(s string) string {
	var hx_result_1 struct {
		Error  error
		Result string
	}
	_ = hx_result_1
	hx_result_1.Result, hx_result_1.Error = url.QueryUnescape(s)
	var this1 struct {
		Error  error
		Result string
	}
	_ = this1
	this1 = hx_result_1
	if this1.Error != nil {
		var e error
		_ = e
		e = this1.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r string
		_ = r
		r = this1.Result
		return r
	}
}

func Hx_Field_stringtools_htmlEscape(s string, quotes__isOptional any) string {
	quotes := HxDefault[bool]()
	_ = quotes

	if quotes__isOptional != nil {
		quotes = HxConvert[bool](quotes__isOptional)
	} else {
		quotes = false
	}
	{
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		{
			var _g_offset int32
			_ = _g_offset
			_g_offset = 0
			var _g_s string
			_ = _g_s
			_g_s = s
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g_offset
				if !(_hx_tmp_0 < int32(utf8.RuneCountInString(_g_s))) {
					break
				}

				var _hx_tmp_1 string
				_ = _hx_tmp_1
				_hx_tmp_1 = _g_s
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g_offset
				_g_offset = (_g_offset + 1)
				var code int32
				_ = code
				code = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, _hx_tmp_2)).Value
				switch code {
				case 34:
					{
						if quotes {
							func() func(x string) {
								var _hx_closure_src_14 func(x any)
								_ = _hx_closure_src_14
								_hx_closure_src_14 = buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_14(((any)(x)))
								}
							}()("&quot;")
						} else {
							buf.VTable.Hx_Field_addChar(code)
						}
					}

				case 38:
					{
						func() func(x string) {
							var _hx_closure_src_15 func(x any)
							_ = _hx_closure_src_15
							_hx_closure_src_15 = buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_15(((any)(x)))
							}
						}()("&amp;")
					}

				case 39:
					{
						if quotes {
							func() func(x string) {
								var _hx_closure_src_16 func(x any)
								_ = _hx_closure_src_16
								_hx_closure_src_16 = buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_16(((any)(x)))
								}
							}()("&#039;")
						} else {
							buf.VTable.Hx_Field_addChar(code)
						}
					}

				case 60:
					{
						func() func(x string) {
							var _hx_closure_src_17 func(x any)
							_ = _hx_closure_src_17
							_hx_closure_src_17 = buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_17(((any)(x)))
							}
						}()("&lt;")
					}

				case 62:
					{
						func() func(x string) {
							var _hx_closure_src_18 func(x any)
							_ = _hx_closure_src_18
							_hx_closure_src_18 = buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_18(((any)(x)))
							}
						}()("&gt;")
					}

				default:
					{
						buf.VTable.Hx_Field_addChar(code)
					}
				}
			}
		}

		return buf.VTable.Hx_Field_toString()
	}
}

func Hx_Field_stringtools_htmlUnescape(s string) string {
	return (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "&gt;")), ">")), "&lt;")), "<")), "&quot;")), "\"")), "&#039;")), "'")), "&amp;")), "&"))
}

func Hx_Field_stringtools_contains(s string, value string) bool {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = s
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = value
	return ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) != -1)
}

func Hx_Field_stringtools_startsWith(s string, start string) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = int32(utf8.RuneCountInString(s))
	if _hx_tmp_1 >= int32(utf8.RuneCountInString(start)) {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = s
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = start
		var _hx_tmp_4 string
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_2
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = _hx_tmp_3
		_hx_tmp_0 = ((Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{Value: 0, Valid: true})) == 0)
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}

func Hx_Field_stringtools_endsWith(s string, end string) bool {
	var elen int32
	_ = elen
	elen = int32(utf8.RuneCountInString(end))
	var slen int32
	_ = slen
	slen = int32(utf8.RuneCountInString(s))
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if slen >= elen {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = end
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_1
		var _hx_tmp_4 string
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_2
		_hx_tmp_0 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{Value: (slen - elen), Valid: true})) == (slen - elen))
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}

func Hx_Field_stringtools_isSpace(s string, pos int32) bool {
	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(s, pos))
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (c).Valid && true && ((c).Value > 8) {
		_hx_tmp_1 = ((c).Valid && true && ((c).Value < 14))
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if (c).Valid == true {
			_hx_tmp_2 = ((c).Value == 32)
		} else {
			_hx_tmp_2 = false
		}

		_hx_tmp_0 = (_hx_tmp_2)
	}

	return _hx_tmp_0
}

func Hx_Field_stringtools_ltrim(s string) string {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(s))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(s, r)
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	if r > 0 {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = r
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		return s
	}
}

func Hx_Field_stringtools_rtrim(s string) string {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(s))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	if r > 0 {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		return s
	}
}

func Hx_Field_stringtools_trim(s string) string {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(s))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = s
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = s
	}

	var s1 string
	_ = s1
	s1 = _hx_tmp_1
	var l1 int32
	_ = l1
	l1 = int32(utf8.RuneCountInString(s1))
	var r1 int32
	_ = r1
	r1 = 0
	for {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if r1 < l1 {
			_hx_tmp_4 = Hx_Field_stringtools_isSpace(s1, r1)
		} else {
			_hx_tmp_4 = false
		}

		if !(_hx_tmp_4) {
			break
		}

		r1++
	}

	if r1 > 0 {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = s1
		var _hx_tmp_6 int32
		_ = _hx_tmp_6
		_hx_tmp_6 = r1
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, _hx_tmp_6, HxNullable[int32]{Value: (l1 - r1), Valid: true}))
	} else {
		return s1
	}
}

func Hx_Field_stringtools_lpad(s string, c string, l int32) string {
	if int32(utf8.RuneCountInString(c)) <= 0 {
		return s
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = l
	l = (_hx_tmp_0 - int32(utf8.RuneCountInString(s)))
	for {
		if !(buf.VTable.Hx_Field_get_length() < l) {
			break
		}

		func() func(x string) {
			var _hx_closure_src_19 func(x any)
			_ = _hx_closure_src_19
			_hx_closure_src_19 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_19(((any)(x)))
			}
		}()(c)
	}

	func() func(x string) {
		var _hx_closure_src_20 func(x any)
		_ = _hx_closure_src_20
		_hx_closure_src_20 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_20(((any)(x)))
		}
	}()(s)
	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_stringtools_rpad(s string, c string, l int32) string {
	if int32(utf8.RuneCountInString(c)) <= 0 {
		return s
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	func() func(x string) {
		var _hx_closure_src_21 func(x any)
		_ = _hx_closure_src_21
		_hx_closure_src_21 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_21(((any)(x)))
		}
	}()(s)
	for {
		if !(buf.VTable.Hx_Field_get_length() < l) {
			break
		}

		func() func(x string) {
			var _hx_closure_src_22 func(x any)
			_ = _hx_closure_src_22
			_hx_closure_src_22 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_22(((any)(x)))
			}
		}()(c)
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_stringtools_replace(s string, sub string, by string) string {
	return (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, sub)), by))
}

func Hx_Field_stringtools_hex(n int32, digits__isOptional any) string {
	digits := HxDefault[HxNullable[int32]]()
	_ = digits

	if digits__isOptional != nil {
		digits = HxConvert[HxNullable[int32]](digits__isOptional)
	} else {

	}
	{
		if (digits).Valid == false {
			var hx_nullable_t5 HxNullable[int32]
			_ = hx_nullable_t5
			hx_nullable_t5 = HxNullable[int32]{}

			digits = HxNullable[int32]{Value: (hx_nullable_t5).Value, Valid: (hx_nullable_t5).Valid}
		}

		var s string
		_ = s
		s = ""
		var hexChars string
		_ = hexChars
		hexChars = "0123456789ABCDEF"
		for {
			s = Hx_Field_go_haxe_hxstring_concat((Hx_Field_go_haxe_hxstring_charAt(hexChars, (n & 15))), s)
			n = ((int32)((((uint32)(n)) >> 4)))
			if !(n > 0) {
				break
			}
		}

		if (digits).Valid != false {
			for {
				if !(true && (digits).Valid && (int32(utf8.RuneCountInString(s)) < (digits).Value)) {
					break
				}

				s = Hx_Field_go_haxe_hxstring_concat("0", s)
			}
		}

		return s
	}
}

func Hx_Field_stringtools_fastCodeAt(s string, index int32) int32 {
	return (Hx_Field_go_haxe_hxstring_charCodeAt(s, index)).Value
}

func Hx_Field_stringtools_unsafeCodeAt(s string, index int32) int32 {
	return (Hx_Field_go_haxe_hxstring_charCodeAt(s, index)).Value
}

func Hx_Field_stringtools_iterator(s string) *Hx_Obj_haxe_iterators_stringiterator {
	return Hx_Obj_haxe_iterators_stringiterator_CreateInstance(s)
}

func Hx_Field_stringtools_keyValueIterator(s string) *Hx_Obj_haxe_iterators_stringkeyvalueiterator {
	return Hx_Obj_haxe_iterators_stringkeyvalueiterator_CreateInstance(s)
}

func Hx_Field_stringtools_isEof(c int32) bool {
	return (c == 0)
}

func Hx_Field_stringtools_quoteUnixArg(argument string) string {
	if argument == "" {
		return "''"
	} else {
		if !Hx_Obj_ereg_CreateInstance("[^a-zA-Z0-9_@%+=:,./-]", "").VTable.Hx_Field_match(argument) {
			return argument
		} else {
			return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("'", Hx_Field_stringtools_replace(argument, "'", "'\"'\"'")), "'")
		}
	}
}

var Hx_Field_stringtools_winMetaCharacters HxArray[int32]

func Hx_Init_Hx_Field_stringtools_winMetaCharacters() HxArray[int32] {
	return HxMakeArrayView[int32](Hx_Field_haxe_systools_winMetaCharacters)
}

func Hx_Field_stringtools_quoteWinArg(argument string, escapeMetaCharacters bool) string {
	var argument1 string
	_ = argument1
	argument1 = argument
	if !Hx_Obj_ereg_CreateInstance("^(/)?[^ \t/\\\\\"]+$", "").VTable.Hx_Field_match(argument1) {
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
		_hx_tmp_3 = argument1
		if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, " ", HxNullable[int32]{})) != -1 {
			_hx_tmp_2 = true
		} else {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = argument1
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = _hx_tmp_4
			_hx_tmp_2 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "\t", HxNullable[int32]{})) != -1)
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (argument1 == "")
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = argument1
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
				var _hx_closure_src_23 func(x any)
				_ = _hx_closure_src_23
				_hx_closure_src_23 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_23(((any)(x)))
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
			_g1 = int32(utf8.RuneCountInString(argument1))
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
					tmp__g1_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(argument1, i))
					if (tmp__g1_1).Valid == false {
						var c HxNullable[int32]
						_ = c
						c = tmp__g1_1
						{
							if bs_buf.VTable.Hx_Field_get_length() > 0 {
								func() func(x string) {
									var _hx_closure_src_24 func(x any)
									_ = _hx_closure_src_24
									_hx_closure_src_24 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_24(((any)(x)))
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
									var _hx_closure_src_25 func(x any)
									_ = _hx_closure_src_25
									_hx_closure_src_25 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_25(((any)(x)))
									}
								}()(bs)
								func() func(x string) {
									var _hx_closure_src_26 func(x any)
									_ = _hx_closure_src_26
									_hx_closure_src_26 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_26(((any)(x)))
									}
								}()(bs)
								bs_buf = Hx_Obj_stringbuf_CreateInstance()
								func() func(x string) {
									var _hx_closure_src_27 func(x any)
									_ = _hx_closure_src_27
									_hx_closure_src_27 = result.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_27(((any)(x)))
									}
								}()("\\\"")
							}

						case 92:
							{
								func() func(x string) {
									var _hx_closure_src_28 func(x any)
									_ = _hx_closure_src_28
									_hx_closure_src_28 = bs_buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_28(((any)(x)))
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
											var _hx_closure_src_29 func(x any)
											_ = _hx_closure_src_29
											_hx_closure_src_29 = result.VTable.Hx_Field_add
											return func(x string) {
												_hx_closure_src_29(((any)(x)))
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
			var _hx_closure_src_30 func(x any)
			_ = _hx_closure_src_30
			_hx_closure_src_30 = result.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_30(((any)(x)))
			}
		}()(bs_buf.VTable.Hx_Field_toString())
		if needquote {
			func() func(x string) {
				var _hx_closure_src_31 func(x any)
				_ = _hx_closure_src_31
				_hx_closure_src_31 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_31(((any)(x)))
				}
			}()(bs_buf.VTable.Hx_Field_toString())
			func() func(x string) {
				var _hx_closure_src_32 func(x any)
				_ = _hx_closure_src_32
				_hx_closure_src_32 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_32(((any)(x)))
				}
			}()("\"")
		}

		argument1 = result.VTable.Hx_Field_toString()
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
			_g1 = int32(utf8.RuneCountInString(argument1))
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
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(argument1, i))
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
		return argument1
	}
}

func init() {
	hxRegisterArrayConstructor[int32]()
}
