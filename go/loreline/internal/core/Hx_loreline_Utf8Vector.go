package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_utf8vector() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_utf8vector() {
}

type Hx_Obj_VTable_loreline_utf8vector interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_utf8vector struct {
	VTable Hx_Obj_VTable_loreline_utf8vector `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_utf8vector_CreateEmptyInstance() *Hx_Obj_loreline_utf8vector {
	obj := &Hx_Obj_loreline_utf8vector{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_utf8vector_CreateInstance() *Hx_Obj_loreline_utf8vector {
	obj := Hx_Obj_loreline_utf8vector_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_utf8vector_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Utf8Vector",
	HxMakeArray[string]("uLength", "uCharCodeAt", "uSubstr", "uSubstring", "uCharAt", "uIndexOf", "slice"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_utf8vector_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_utf8vector_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_utf8vector_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "uLength":
			return Hx_Field_loreline_utf8vector_uLength
		case "uCharCodeAt":
			return Hx_Field_loreline_utf8vector_uCharCodeAt
		case "uSubstr":
			return Hx_Field_loreline_utf8vector_uSubstr
		case "uSubstring":
			return Hx_Field_loreline_utf8vector_uSubstring
		case "uCharAt":
			return Hx_Field_loreline_utf8vector_uCharAt
		case "uIndexOf":
			return Hx_Field_loreline_utf8vector_uIndexOf
		case "slice":
			return Hx_Field_loreline_utf8vector_slice
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_utf8vector) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_utf8vector_RTTI
}

func (this *Hx_Obj_loreline_utf8vector) Hx_As_Hx_Obj_loreline_utf8vector() *Hx_Obj_loreline_utf8vector {
	return this
}

func Hx_Field_loreline_utf8vector_uLength(v HxArray[int32]) int32 {
	return HxMakeArrayView[any](v).Len()
}

func Hx_Field_loreline_utf8vector_uCharCodeAt(v HxArray[int32], pos int32) int32 {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if pos < 0 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = pos
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_0 {
		return -1
	} else {
		return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}
}

func Hx_Field_loreline_utf8vector_uSubstr(v HxArray[int32], pos int32, _hx_reserved_len__isOptional any) string {
	_hx_reserved_len := HxDefault[HxNullable[int32]]()
	_ = _hx_reserved_len

	if _hx_reserved_len__isOptional != nil {
		_hx_reserved_len = HxConvert[HxNullable[int32]](_hx_reserved_len__isOptional)
	} else {

	}
	{
		if (_hx_reserved_len).Valid == false {
			var hx_nullable_t25 HxNullable[int32]
			_ = hx_nullable_t25
			hx_nullable_t25 = HxNullable[int32]{}

			_hx_reserved_len = HxNullable[int32]{Value: (hx_nullable_t25).Value, Valid: (hx_nullable_t25).Valid}
		}

		var _hx_tmp_0 HxArray[int32]
		_ = _hx_tmp_0
		_hx_tmp_0 = v
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = pos
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (_hx_reserved_len).Valid != false {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = pos
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_2 = (_hx_tmp_4 + (_hx_reserved_len).Value)
		} else {
			_hx_tmp_2 = HxMakeArrayView[any](v).Len()
		}

		return Hx_Field_loreline_utf8vector_slice(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2)
	}
}

func Hx_Field_loreline_utf8vector_uSubstring(v HxArray[int32], startIndex int32, endIndex__isOptional any) string {
	endIndex := HxDefault[HxNullable[int32]]()
	_ = endIndex

	if endIndex__isOptional != nil {
		endIndex = HxConvert[HxNullable[int32]](endIndex__isOptional)
	} else {

	}
	{
		if (endIndex).Valid == false {
			var hx_nullable_t26 HxNullable[int32]
			_ = hx_nullable_t26
			hx_nullable_t26 = HxNullable[int32]{}

			endIndex = HxNullable[int32]{Value: (hx_nullable_t26).Value, Valid: (hx_nullable_t26).Valid}
		}

		var _hx_tmp_0 HxArray[int32]
		_ = _hx_tmp_0
		_hx_tmp_0 = v
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = startIndex
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (endIndex).Valid != false {
			_hx_tmp_2 = (endIndex).Value
		} else {
			_hx_tmp_2 = HxMakeArrayView[any](v).Len()
		}

		return Hx_Field_loreline_utf8vector_slice(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2)
	}
}

func Hx_Field_loreline_utf8vector_uCharAt(v HxArray[int32], pos int32) string {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if pos < 0 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = pos
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_0 {
		return ""
	} else {
		return Hx_Field_loreline_utf8vector_slice(v, pos, (pos + 1))
	}
}

func Hx_Field_loreline_utf8vector_uIndexOf(v HxArray[int32], substr string, startIndex__isOptional any) int32 {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t27 HxNullable[int32]
			_ = hx_nullable_t27
			hx_nullable_t27 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t27).Value, Valid: (hx_nullable_t27).Valid}
		}

		var needle HxArray[int32]
		_ = needle
		needle = Hx_Field_loreline_utf8_toCodes(substr)
		var nlen int32
		_ = nlen
		nlen = HxMakeArrayView[any](needle).Len()
		if nlen == 0 {
			if (startIndex).Valid != false {
				return (startIndex).Value
			} else {
				return 0
			}
		}

		var _hx_tmp_0 HxNullable[int32]
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (startIndex).Valid != false {
			_hx_tmp_1 = ((startIndex).Valid && true && ((startIndex).Value > 0))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = startIndex
		} else {
			_hx_tmp_0 = HxNullable[int32]{Value: 0, Valid: true}
		}

		var i HxNullable[int32]
		_ = i
		i = _hx_tmp_0
		var last int32
		_ = last
		last = (HxMakeArrayView[any](v).Len() - nlen)
		for {
			if !((i).Valid && true && ((i).Value <= last) || (!(i).Valid && !true)) {
				break
			}

			var k int32
			_ = k
			k = 0
			for {
				var _hx_tmp_2 bool
				_ = _hx_tmp_2
				if k < nlen {
					var _hx_tmp_3 HxArray[any]
					_ = _hx_tmp_3
					_hx_tmp_3 = HxMakeArrayView[any](v)
					var _hx_tmp_4 HxArray[any]
					_ = _hx_tmp_4
					_hx_tmp_4 = _hx_tmp_3
					_hx_tmp_2 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_4.Get(((i).Value + k)), HxMakeArrayView[any](needle).Get(k)))))
				} else {
					_hx_tmp_2 = false
				}

				if !(_hx_tmp_2) {
					break
				}

				k++
			}

			if k == nlen {
				return (i).Value
			}

			(i).Value++
		}

		return -1
	}
}

func Hx_Field_loreline_utf8vector_slice(v HxArray[int32], start int32, end int32) string {
	if start < 0 {
		start = 0
	}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = end
	if _hx_tmp_0 > HxMakeArrayView[any](v).Len() {
		end = HxMakeArrayView[any](v).Len()
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var i int32
	_ = i
	i = start
	for i < end {
		buf.VTable.Hx_Field_addChar(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(i)))))
		i++
	}

	return buf.VTable.Hx_Field_toString()
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[int32]()
}
