package core

import "unicode/utf8"
import "strings"

func Hx_ClassInit_Hx_Obj_VTable_stringbuf() {
}

func Hx_StaticInit_Hx_Obj_VTable_stringbuf() {
}

type Hx_Obj_VTable_stringbuf interface {
	Hx_Field_clear()
	Hx_Field_toString() string
	Hx_Field_get_length() int32
	Hx_Field_add(x any)
	Hx_Field_addChar(c int32)
	Hx_Field_addSub(s string, pos int32, _hx_reserved_len__isOptional any)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_stringbuf struct {
	VTable     Hx_Obj_VTable_stringbuf `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_b strings.Builder
}

func Hx_Obj_stringbuf_CreateEmptyInstance() *Hx_Obj_stringbuf {
	obj := &Hx_Obj_stringbuf{}
	obj.VTable = obj
	obj.Hx_Field_b = HxDefault[strings.Builder]()
	return obj
}

func Hx_Obj_stringbuf_CreateInstance() *Hx_Obj_stringbuf {
	obj := Hx_Obj_stringbuf_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_stringbuf) Hx_New() {
	{
		this.Hx_Field_b = HxDefault[strings.Builder]()
	}

}

var Hx_Obj_stringbuf_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"StringBuf",
	HxMakeArray[string](),
	HxMakeArray[string]("b", "get_length", "add", "addChar", "addSub", "clear", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_stringbuf_CreateInstance()
	},
	func() any {
		return Hx_Obj_stringbuf_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_stringbuf_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_stringbuf) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_stringbuf_RTTI
}

func (this *Hx_Obj_stringbuf) Hx_As_Hx_Obj_stringbuf() *Hx_Obj_stringbuf {
	return this
}

func (this *Hx_Obj_stringbuf) Hx_Field_get_length() int32 {
	return ((int32)(utf8.RuneCountInString(this.Hx_Field_b.String())))
}

func (this *Hx_Obj_stringbuf) Hx_Field_add(x any) {
	var s string
	_ = s
	s = ""
	var ok bool
	_ = ok
	ok = false
	s, ok = x.(string)
	if ok {
		var hx_result_2 struct {
			Error  error
			Result int
		}
		_ = hx_result_2
		hx_result_2.Result, hx_result_2.Error = this.Hx_Field_b.WriteString(s)
		var this1 struct {
			Error  error
			Result int
		}
		_ = this1
		this1 = hx_result_2
		if this1.Error != nil {
			var e error
			_ = e
			e = this1.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r int
			_ = r
			r = this1.Result
		}
	} else {
		var sl []byte
		_ = sl
		sl, ok = x.([]byte)
		if ok {
			var hx_result_3 struct {
				Error  error
				Result int
			}
			_ = hx_result_3
			hx_result_3.Result, hx_result_3.Error = this.Hx_Field_b.Write(sl)
			var this1 struct {
				Error  error
				Result int
			}
			_ = this1
			this1 = hx_result_3
			if this1.Error != nil {
				var e error
				_ = e
				e = this1.Error
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			} else {
				var r int
				_ = r
				r = this1.Result
			}
		} else {
			var hx_result_4 struct {
				Error  error
				Result int
			}
			_ = hx_result_4
			hx_result_4.Result, hx_result_4.Error = this.Hx_Field_b.WriteString(Hx_Field_std_string(x))
			var this1 struct {
				Error  error
				Result int
			}
			_ = this1
			this1 = hx_result_4
			if this1.Error != nil {
				var e error
				_ = e
				e = this1.Error
				panic(Hx_Field_haxe_exception_thrown(((any)(e))))
			} else {
				var r int
				_ = r
				r = this1.Result
			}
		}
	}
}

func (this *Hx_Obj_stringbuf) Hx_Field_addChar(c int32) {
	var hx_result_5 struct {
		Error  error
		Result int
	}
	_ = hx_result_5
	hx_result_5.Result, hx_result_5.Error = this.Hx_Field_b.WriteRune(rune(c))
	var this1 struct {
		Error  error
		Result int
	}
	_ = this1
	this1 = hx_result_5
	if this1.Error != nil {
		var e error
		_ = e
		e = this1.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r int
		_ = r
		r = this1.Result
	}
}

func (this *Hx_Obj_stringbuf) Hx_Field_addSub(s string, pos int32, _hx_reserved_len__isOptional any) {
	_hx_reserved_len := HxDefault[HxNullable[int32]]()
	_ = _hx_reserved_len

	if _hx_reserved_len__isOptional != nil {
		_hx_reserved_len = HxConvert[HxNullable[int32]](_hx_reserved_len__isOptional)
	} else {

	}
	{
		if (_hx_reserved_len).Valid == false {
			var hx_nullable_t6 HxNullable[int32]
			_ = hx_nullable_t6
			hx_nullable_t6 = HxNullable[int32]{}

			_hx_reserved_len = HxNullable[int32]{Value: (hx_nullable_t6).Value, Valid: (hx_nullable_t6).Valid}
		}

		var hx_result_6 struct {
			Error  error
			Result int
		}
		_ = hx_result_6
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		if (_hx_reserved_len).Valid == false {
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
			_hx_tmp_0 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{}))
		} else {
			_hx_tmp_0 = (Hx_Field_go_haxe_hxstring_substr(s, pos, _hx_reserved_len))
		}

		hx_result_6.Result, hx_result_6.Error = this.Hx_Field_b.WriteString(_hx_tmp_0)
		var this1 struct {
			Error  error
			Result int
		}
		_ = this1
		this1 = hx_result_6
		if this1.Error != nil {
			var e error
			_ = e
			e = this1.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r int
			_ = r
			r = this1.Result
		}
	}
}

func (this *Hx_Obj_stringbuf) Hx_Field_clear() {
	this.Hx_Field_b.Reset()
}

func (this *Hx_Obj_stringbuf) Hx_Field_toString() string {
	return this.Hx_Field_b.String()
}
