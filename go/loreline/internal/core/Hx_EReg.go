package core

import "unicode/utf8"
import "regexp"

func Hx_ClassInit_Hx_Obj_VTable_ereg() {
}

func Hx_StaticInit_Hx_Obj_VTable_ereg() {
}

type Hx_Obj_VTable_ereg interface {
	Hx_Field_matchedLeft() string
	Hx_Field_matchedRight() string
	Hx_Field_matchedPos() any
	Hx_Field_matchSub(s string, pos int32, _hx_reserved_len__isOptional any) bool
	Hx_Field_matchedNum() int32
	Hx_Field_split(s string) HxArray[string]
	Hx_Field_replace(s string, by string) string
	Hx_Field_map(s string, f func(_hx_carg0 *Hx_Obj_ereg) string) string
	Hx_Field_byteSub(s string, a int32, b int32) string
	Hx_Field_byteFrom(s string, a int32) string
	Hx_Field_byteLen(s string) int32
	Hx_Field_charToByte(s string, pos int32) int32
	Hx_Field_match(s string) bool
	Hx_Field_matched(n int32) string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_ereg struct {
	VTable            Hx_Obj_VTable_ereg `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_re       *regexp.Regexp
	Hx_Field_str      string
	Hx_Field_matchPos HxArray[int]
	Hx_Field_global   bool
}

func Hx_Obj_ereg_CreateEmptyInstance() *Hx_Obj_ereg {
	obj := &Hx_Obj_ereg{}
	obj.VTable = obj
	obj.Hx_Field_re = HxDefault[*regexp.Regexp]()
	obj.Hx_Field_str = HxDefault[string]()
	obj.Hx_Field_matchPos = HxDefault[HxArray[int]]()
	obj.Hx_Field_global = HxDefault[bool]()
	return obj
}

func Hx_Obj_ereg_CreateInstance(r string, opt string) *Hx_Obj_ereg {
	obj := Hx_Obj_ereg_CreateEmptyInstance()
	obj.Hx_New(r, opt)
	return obj
}

func (this *Hx_Obj_ereg) Hx_New(r string, opt string) {
	{
		var flags string
		_ = flags
		flags = ""
		if opt != HxStringNull {
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = opt
			if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "i", HxNullable[int32]{})) != -1 {
				flags = Hx_Field_go_haxe_hxdynamic_stringConcat(flags, "i")
			}

			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = opt
			if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, "m", HxNullable[int32]{})) != -1 {
				flags = Hx_Field_go_haxe_hxdynamic_stringConcat(flags, "m")
			}

			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = opt
			if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_2, "s", HxNullable[int32]{})) != -1 {
				flags = Hx_Field_go_haxe_hxdynamic_stringConcat(flags, "s")
			}

			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = opt
			this.Hx_Field_global = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, "g", HxNullable[int32]{})) != -1)
		}

		var _hx_tmp_4 string
		_ = _hx_tmp_4
		if int32(utf8.RuneCountInString(flags)) > 0 {
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("(?", flags), ")"), r)
		} else {
			_hx_tmp_4 = r
		}

		var pattern string
		_ = pattern
		pattern = _hx_tmp_4
		var hx_result_0 struct {
			Error  error
			Result *regexp.Regexp
		}
		_ = hx_result_0
		hx_result_0.Result, hx_result_0.Error = regexp.Compile(pattern)
		var compiled struct {
			Error  error
			Result *regexp.Regexp
		}
		_ = compiled
		compiled = hx_result_0
		var tmp bool
		_ = tmp
		if compiled.Error != nil {
			tmp = false
		} else {
			tmp = true
		}

		if tmp {
			var tmp_tmp_1 *regexp.Regexp
			_ = tmp_tmp_1
			if compiled.Error != nil {
				var e error
				_ = e
				e = compiled.Error
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			} else {
				var r *regexp.Regexp
				_ = r
				r = compiled.Result
				tmp_tmp_1 = r
			}

			this.Hx_Field_re = tmp_tmp_1
		}
	}

}

var Hx_Obj_ereg_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"EReg",
	HxMakeArray[string]("escape"),
	HxMakeArray[string]("re", "str", "matchPos", "global", "byteSub", "byteFrom", "byteLen", "charToByte", "match", "matched", "matchedLeft", "matchedRight", "matchedPos", "matchSub", "matchedNum", "split", "replace", "map"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_ereg_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[string](params.Get(1)))
	},
	func() any {
		return Hx_Obj_ereg_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_ereg_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "escape":
			return Hx_Field_ereg_escape
		}
		return nil
	}
}

func (this *Hx_Obj_ereg) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_ereg_RTTI
}

func (this *Hx_Obj_ereg) Hx_As_Hx_Obj_ereg() *Hx_Obj_ereg {
	return this
}

func (this *Hx_Obj_ereg) Hx_Field_byteSub(s string, a int32, b int32) string {
	return s[a:b]
}

func (this *Hx_Obj_ereg) Hx_Field_byteFrom(s string, a int32) string {
	return s[a:]
}

func (this *Hx_Obj_ereg) Hx_Field_byteLen(s string) int32 {
	return ((int32)(len(s)))
}

func (this *Hx_Obj_ereg) Hx_Field_charToByte(s string, pos int32) int32 {
	var i int32
	_ = i
	i = 0
	var n int32
	_ = n
	n = ((int32)(len(s)))
	for (pos > 0) && (i < n) {
		var hx_tuple_2 struct {
			R    rune
			Size int
		}
		_ = hx_tuple_2
		hx_tuple_2.R, hx_tuple_2.Size = utf8.DecodeRuneInString(s[i:])
		i = ((int32)((hx_tuple_2.Size + int(i))))
		pos--
	}

	return i
}

func (this *Hx_Obj_ereg) Hx_Field_match(s string) bool {
	var idx []int
	_ = idx
	idx = (*this.Hx_Field_re).FindStringSubmatchIndex(s)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if idx == nil {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = len(idx)
		var _hx_tmp_2 int
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 == int(((int32)(0))))
	}

	if _hx_tmp_0 {
		this.Hx_Field_matchPos = nil
		return false
	}

	var self []int
	_ = self
	self = idx
	this.Hx_Field_matchPos = HxMakeArray(self...)
	this.Hx_Field_str = s
	return true
}

func (this *Hx_Obj_ereg) Hx_Field_matched(n int32) string {
	if this.Hx_Field_matchPos == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)("No string matched"))))
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if n < 0 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = ((n * 2) + 1)
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 >= this.Hx_Field_matchPos.Len())
	}

	if _hx_tmp_0 {
		return HxStringNull
	}

	var a int
	_ = a
	a = this.Hx_Field_matchPos.Get((n * 2))
	var b int
	_ = b
	b = this.Hx_Field_matchPos.Get(((n * 2) + 1))
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	var _hx_tmp_4 int
	_ = _hx_tmp_4
	_hx_tmp_4 = a
	if _hx_tmp_4 < int(((int32)(0))) {
		_hx_tmp_3 = true
	} else {
		var _hx_tmp_5 int
		_ = _hx_tmp_5
		_hx_tmp_5 = b
		var _hx_tmp_6 int
		_ = _hx_tmp_6
		_hx_tmp_6 = _hx_tmp_5
		_hx_tmp_3 = (_hx_tmp_6 < int(((int32)(0))))
	}

	if _hx_tmp_3 {
		return HxStringNull
	}

	return this.Hx_Field_str[((int32)(a)):((int32)(b))]
}

func (this *Hx_Obj_ereg) Hx_Field_matchedLeft() string {
	if this.Hx_Field_matchPos == nil {
		return HxStringNull
	}

	return this.Hx_Field_str[0:((int32)(this.Hx_Field_matchPos.Get(0)))]
}

func (this *Hx_Obj_ereg) Hx_Field_matchedRight() string {
	if this.Hx_Field_matchPos == nil {
		return HxStringNull
	}

	return this.Hx_Field_str[((int32)(this.Hx_Field_matchPos.Get(1))):]
}

func (this *Hx_Obj_ereg) Hx_Field_matchedPos() any {
	if this.Hx_Field_matchPos == nil {
		return nil
	}

	var begin int
	_ = begin
	begin = this.Hx_Field_matchPos.Get(0)
	var end int
	_ = end
	end = this.Hx_Field_matchPos.Get(1)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(((int32)(utf8.RuneCountInString(this.Hx_Field_str[0:((int32)(begin))])))))
	return any(map[string]any{"pos": _hx_tmp_0, "len": ((any)(((int32)(utf8.RuneCountInString(this.Hx_Field_str[((int32)(begin)):((int32)(end))])))))})
}

func (this *Hx_Obj_ereg) Hx_Field_matchSub(s string, pos int32, _hx_reserved_len__isOptional any) bool {
	_hx_reserved_len := HxDefault[int32]()
	_ = _hx_reserved_len

	if _hx_reserved_len__isOptional != nil {
		_hx_reserved_len = HxConvert[int32](_hx_reserved_len__isOptional)
	} else {
		_hx_reserved_len = -1
	}
	{
		var pos1 int32
		_ = pos1
		pos1 = pos
		var i int32
		_ = i
		i = 0
		var n int32
		_ = n
		n = ((int32)(len(s)))
		for (pos1 > 0) && (i < n) {
			var hx_tuple_3 struct {
				R    rune
				Size int
			}
			_ = hx_tuple_3
			hx_tuple_3.R, hx_tuple_3.Size = utf8.DecodeRuneInString(s[i:])
			i = ((int32)((hx_tuple_3.Size + int(i))))
			pos1--
		}

		var _hx_reserved_byte int32
		_ = _hx_reserved_byte
		_hx_reserved_byte = i
		var sub string
		_ = sub
		if _hx_reserved_len < 0 {
			sub = s[_hx_reserved_byte:]
		} else {
			var tmp_pos1_1 int32
			_ = tmp_pos1_1
			tmp_pos1_1 = (pos + _hx_reserved_len)
			var tmp_i_1 int32
			_ = tmp_i_1
			tmp_i_1 = 0
			var tmp_n_1 int32
			_ = tmp_n_1
			tmp_n_1 = ((int32)(len(s)))
			for (tmp_pos1_1 > 0) && (tmp_i_1 < tmp_n_1) {
				var hx_tuple_4 struct {
					R    rune
					Size int
				}
				_ = hx_tuple_4
				hx_tuple_4.R, hx_tuple_4.Size = utf8.DecodeRuneInString(s[tmp_i_1:])
				tmp_i_1 = ((int32)((hx_tuple_4.Size + int(tmp_i_1))))
				tmp_pos1_1--
			}

			sub = s[_hx_reserved_byte:tmp_i_1]
		}

		var idx []int
		_ = idx
		idx = (*this.Hx_Field_re).FindStringSubmatchIndex(sub)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if idx == nil {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_1 int
			_ = _hx_tmp_1
			_hx_tmp_1 = len(idx)
			var _hx_tmp_2 int
			_ = _hx_tmp_2
			_hx_tmp_2 = _hx_tmp_1
			_hx_tmp_0 = (_hx_tmp_2 == int(((int32)(0))))
		}

		if _hx_tmp_0 {
			this.Hx_Field_matchPos = nil
			return false
		}

		var self []int
		_ = self
		self = idx
		this.Hx_Field_matchPos = HxMakeArray(self...)
		this.Hx_Field_str = s
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = this.Hx_Field_matchPos.Len()
			for _g < _g1 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_g = (_g + 1)
				var tmp_i_1 int32
				_ = tmp_i_1
				tmp_i_1 = _hx_tmp_3
				var _hx_tmp_4 int
				_ = _hx_tmp_4
				_hx_tmp_4 = this.Hx_Field_matchPos.Get(tmp_i_1)
				if _hx_tmp_4 >= int(((int32)(0))) {
					var _hx_tmp_5 HxArray[int]
					_ = _hx_tmp_5
					_hx_tmp_5 = this.Hx_Field_matchPos
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = tmp_i_1
					var _hx_tmp_7 int
					_ = _hx_tmp_7
					_hx_tmp_7 = _hx_tmp_5.Get(_hx_tmp_6)
					_hx_tmp_5.Set(_hx_tmp_6, (_hx_tmp_7 + int(_hx_reserved_byte)))
				}
			}
		}

		return true
	}
}

func (this *Hx_Obj_ereg) Hx_Field_matchedNum() int32 {
	if this.Hx_Field_matchPos == nil {
		return 0
	} else {
		return Hx_Field_std_int(((float64)((((float64)(this.Hx_Field_matchPos.Len())) / ((float64)(2))))))
	}
}

func (this *Hx_Obj_ereg) Hx_Field_split(s string) HxArray[string] {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = s
	var _hx_tmp_1 int
	_ = _hx_tmp_1
	if this.Hx_Field_global {
		_hx_tmp_1 = int(((int32)(-1)))
	} else {
		_hx_tmp_1 = int(((int32)(2)))
	}

	var self []string
	_ = self
	self = (*this.Hx_Field_re).Split(_hx_tmp_0, _hx_tmp_1)
	return HxMakeArray(self...)
}

func (this *Hx_Obj_ereg) Hx_Field_replace(s string, by string) string {
	if this.Hx_Field_global {
		return (*this.Hx_Field_re).ReplaceAllString(s, by)
	}

	var idx []int
	_ = idx
	idx = (*this.Hx_Field_re).FindStringSubmatchIndex(s)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if idx == nil {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = len(idx)
		var _hx_tmp_2 int
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 == int(((int32)(0))))
	}

	if _hx_tmp_0 {
		return s
	}

	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = s
	var before string
	_ = before
	before = _hx_tmp_3[0:((int32)(idx[0]))]
	var _hx_tmp_4 string
	_ = _hx_tmp_4
	_hx_tmp_4 = s
	var after string
	_ = after
	after = _hx_tmp_4[((int32)(idx[1])):]
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{Value: 0, Valid: true}

	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	if (length).Valid != false {
		_hx_tmp_5 = (length).Value
	} else {
		_hx_tmp_5 = 0
	}

	var dst []byte
	_ = dst
	dst = make([]byte, _hx_tmp_5)
	dst = (*this.Hx_Field_re).ExpandString(dst, by, s, idx)
	var _hx_tmp_6 string
	_ = _hx_tmp_6
	_hx_tmp_6 = before
	return Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_6, string(dst)), after)
}

func (this *Hx_Obj_ereg) Hx_Field_map(s string, f func(_hx_carg0 *Hx_Obj_ereg) string) string {
	var matches [][]int
	_ = matches
	if this.Hx_Field_global {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = s
		matches = (*this.Hx_Field_re).FindAllStringSubmatchIndex(_hx_tmp_0, int(((int32)(-1))))
	} else {
		var m []int
		_ = m
		m = (*this.Hx_Field_re).FindStringSubmatchIndex(s)
		if m == nil {
			var self HxArray[[]int]
			_ = self
			self = HxMakeArray[[]int]()
			matches = self.Underlying()
		} else {
			var self HxArray[[]int]
			_ = self
			self = HxMakeArray[[]int](m)
			matches = self.Underlying()
		}
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var pos int32
	_ = pos
	pos = 0
	var mi int32
	_ = mi
	mi = 0
	var bytes int32
	_ = bytes
	bytes = ((int32)(len(s)))
	for pos < bytes {
		for {
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			var _hx_tmp_2 float64
			_ = _hx_tmp_2
			_hx_tmp_2 = float64(mi)
			if _hx_tmp_2 < float64(len(matches)) {
				var _hx_tmp_3 int
				_ = _hx_tmp_3
				_hx_tmp_3 = matches[mi][0]
				var _hx_tmp_4 int
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_1 = (_hx_tmp_4 < int(pos))
			} else {
				_hx_tmp_1 = false
			}

			if !(_hx_tmp_1) {
				break
			}

			mi++
		}

		var _hx_tmp_5 float64
		_ = _hx_tmp_5
		_hx_tmp_5 = float64(mi)
		if _hx_tmp_5 >= float64(len(matches)) {
			break
		}

		var idx []int
		_ = idx
		idx = matches[mi]
		var start int
		_ = start
		start = idx[0]
		var end int
		_ = end
		end = idx[1]
		func() func(x string) {
			var _hx_closure_src_0 func(x any)
			_ = _hx_closure_src_0
			_hx_closure_src_0 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_0(((any)(x)))
			}
		}()(s[pos:((int32)(start))])
		var self []int
		_ = self
		self = idx
		this.Hx_Field_matchPos = HxMakeArray(self...)
		this.Hx_Field_str = s
		func() func(x string) {
			var _hx_closure_src_1 func(x any)
			_ = _hx_closure_src_1
			_hx_closure_src_1 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_1(((any)(x)))
			}
		}()(f(this))
		if end == start {
			var _hx_tmp_6 int
			_ = _hx_tmp_6
			_hx_tmp_6 = end
			if _hx_tmp_6 < int(int32(utf8.RuneCountInString(s))) {
				var hx_tuple_5 struct {
					R    rune
					Size int
				}
				_ = hx_tuple_5
				hx_tuple_5.R, hx_tuple_5.Size = utf8.DecodeRuneInString(s[((int32)(end)):])
				var size int
				_ = size
				size = hx_tuple_5.Size
				func() func(x string) {
					var _hx_closure_src_2 func(x any)
					_ = _hx_closure_src_2
					_hx_closure_src_2 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_2(((any)(x)))
					}
				}()(s[((int32)(end)):((int32)((end + size)))])
				pos = ((int32)((end + size)))
			} else {
				var _hx_tmp_7 int
				_ = _hx_tmp_7
				_hx_tmp_7 = end
				pos = ((int32)((_hx_tmp_7 + int(((int32)(1))))))
			}
		} else {
			pos = ((int32)(end))
		}

		mi++
		if !this.Hx_Field_global {
			break
		}
	}

	var _hx_tmp_8 string
	_ = _hx_tmp_8
	_hx_tmp_8 = s
	var _hx_tmp_9 int32
	_ = _hx_tmp_9
	if pos > bytes {
		_hx_tmp_9 = bytes
	} else {
		_hx_tmp_9 = pos
	}

	func() func(x string) {
		var _hx_closure_src_3 func(x any)
		_ = _hx_closure_src_3
		_hx_closure_src_3 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_3(((any)(x)))
		}
	}()(_hx_tmp_8[_hx_tmp_9:])
	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_ereg_escape(s string) string {
	return regexp.QuoteMeta(s)
}

func init() {
	hxRegisterArrayConstructor[[]int]()
	hxRegisterArrayConstructor[int]()
	hxRegisterArrayConstructor[string]()
}
