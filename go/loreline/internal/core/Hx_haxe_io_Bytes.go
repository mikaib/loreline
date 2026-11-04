package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_bytes() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_bytes() {
}

type Hx_Obj_VTable_haxe_io_bytes interface {
	Hx_Field_getData() Hx_Typedef_haxe_io_bytesdata
	Hx_Field_get(pos int32) int32
	Hx_Field_fill(pos int32, _hx_reserved_len int32, value int32)
	Hx_Field_getDouble(pos int32) float64
	Hx_Field_getInt32(pos int32) int32
	Hx_Field_setInt32(pos int32, v int32)
	Hx_Field_toString() string
	Hx_Field_set(pos int32, v int32)
	Hx_Field_getFloat(pos int32) float64
	Hx_Field_setFloat(pos int32, v float64)
	Hx_Field_setDouble(pos int32, v float64)
	Hx_Field_getUInt16(pos int32) int32
	Hx_Field_setUInt16(pos int32, v int32)
	Hx_Field_setInt64(pos int32, v Hx_Typedef_haxe__int64___int64)
	Hx_Field_blit(pos int32, src *Hx_Obj_haxe_io_bytes, srcpos int32, _hx_reserved_len int32)
	Hx_Field_sub(pos int32, _hx_reserved_len int32) *Hx_Obj_haxe_io_bytes
	Hx_Field_compare(other *Hx_Obj_haxe_io_bytes) int32
	Hx_Field_getInt64(pos int32) Hx_Typedef_haxe__int64___int64
	Hx_Field_getString(pos int32, _hx_reserved_len int32, encoding__isOptional any) string
	Hx_Field_readString(pos int32, _hx_reserved_len int32) string
	Hx_Field_toHex() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_bytes struct {
	VTable          Hx_Obj_VTable_haxe_io_bytes `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_length int32
	Hx_Field_b      Hx_Typedef_haxe_io_bytesdata
}

func Hx_Obj_haxe_io_bytes_CreateEmptyInstance() *Hx_Obj_haxe_io_bytes {
	obj := &Hx_Obj_haxe_io_bytes{}
	obj.VTable = obj
	obj.Hx_Field_length = HxDefault[int32]()
	obj.Hx_Field_b = HxDefault[Hx_Typedef_haxe_io_bytesdata]()
	return obj
}

func Hx_Obj_haxe_io_bytes_CreateInstance(length int32, b Hx_Typedef_haxe_io_bytesdata) *Hx_Obj_haxe_io_bytes {
	obj := Hx_Obj_haxe_io_bytes_CreateEmptyInstance()
	obj.Hx_New(length, b)
	return obj
}

func (this *Hx_Obj_haxe_io_bytes) Hx_New(length int32, b Hx_Typedef_haxe_io_bytesdata) {
	{
		this.Hx_Field_length = length
		this.Hx_Field_b = b
	}

}

var Hx_Obj_haxe_io_bytes_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.Bytes",
	HxMakeArray[string]("alloc", "ofString", "ofData", "ofHex", "fastGet"),
	HxMakeArray[string]("length", "b", "get", "set", "blit", "fill", "sub", "compare", "getDouble", "getFloat", "setDouble", "setFloat", "getUInt16", "setUInt16", "getInt32", "getInt64", "setInt32", "setInt64", "getString", "readString", "toString", "toHex", "getData"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_bytes_CreateInstance(HxConvert[int32](params.Get(0)), HxConvert[Hx_Typedef_haxe_io_bytesdata](params.Get(1)))
	},
	func() any {
		return Hx_Obj_haxe_io_bytes_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_bytes_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "alloc":
			return Hx_Field_haxe_io_bytes_alloc
		case "ofString":
			return Hx_Field_haxe_io_bytes_ofString
		case "ofData":
			return Hx_Field_haxe_io_bytes_ofData
		case "ofHex":
			return Hx_Field_haxe_io_bytes_ofHex
		case "fastGet":
			return Hx_Field_haxe_io_bytes_fastGet
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_bytes_RTTI
}

func (this *Hx_Obj_haxe_io_bytes) Hx_As_Hx_Obj_haxe_io_bytes() *Hx_Obj_haxe_io_bytes {
	return this
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_get(pos int32) int32 {
	return ((int32)(this.Hx_Field_b.Get(pos)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_set(pos int32, v int32) {
	this.Hx_Field_b.Set(pos, byte((v & 255)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_blit(pos int32, src *Hx_Obj_haxe_io_bytes, srcpos int32, _hx_reserved_len int32) {
	if ((((pos < 0) || (srcpos < 0)) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > this.Hx_Field_length)) || ((srcpos + _hx_reserved_len) > src.Hx_Field_length) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
	}

	var b1 Hx_Typedef_haxe_io_bytesdata
	_ = b1
	b1 = this.Hx_Field_b
	var b2 Hx_Typedef_haxe_io_bytesdata
	_ = b2
	b2 = src.Hx_Field_b
	if (b1 == b2) && (pos > srcpos) {
		var i int32
		_ = i
		i = _hx_reserved_len
		for i > 0 {
			i--
			b1.Set((i + pos), b2.Get((i + srcpos)))
		}

		return
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			b1.Set((i + pos), b2.Get((i + srcpos)))
		}
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_fill(pos int32, _hx_reserved_len int32, value int32) {
	var _g int32
	_ = _g
	_g = 0
	var _g1 int32
	_ = _g1
	_g1 = _hx_reserved_len
	for _g < _g1 {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = _g
		_g = (_g + 1)
		var i int32
		_ = i
		i = _hx_tmp_0
		var _hx_tmp_1 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_b
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = pos
		pos = (pos + 1)
		_hx_tmp_1.Set(_hx_tmp_2, byte((value & 255)))
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_sub(pos int32, _hx_reserved_len int32) *Hx_Obj_haxe_io_bytes {
	if ((pos < 0) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > this.Hx_Field_length) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
	}

	var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_b
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = pos
	return Hx_Obj_haxe_io_bytes_CreateInstance(_hx_reserved_len, ((Hx_Typedef_haxe_io_bytesdata)((Hx_Array_Slice[byte](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{Value: (pos + _hx_reserved_len), Valid: true})))))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_compare(other *Hx_Obj_haxe_io_bytes) int32 {
	var b1 Hx_Typedef_haxe_io_bytesdata
	_ = b1
	b1 = this.Hx_Field_b
	var b2 Hx_Typedef_haxe_io_bytesdata
	_ = b2
	b2 = other.Hx_Field_b
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if this.Hx_Field_length < other.Hx_Field_length {
		_hx_tmp_0 = this.Hx_Field_length
	} else {
		_hx_tmp_0 = other.Hx_Field_length
	}

	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = _hx_tmp_0
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			if b1.Get(i) != b2.Get(i) {
				return (((int32)(b1.Get(i))) - ((int32)(b2.Get(i))))
			}
		}
	}

	return (this.Hx_Field_length - other.Hx_Field_length)
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getDouble(pos int32) float64 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this.Hx_Field_b.Get((pos + 1))))
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)(this.Hx_Field_b.Get((pos + 2))))
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int32)(this.Hx_Field_b.Get((pos + 3))))
	var v int32
	_ = v
	v = (((((int32)(this.Hx_Field_b.Get(pos))) | (_hx_tmp_0 << 8)) | (_hx_tmp_1 << 16)) | (_hx_tmp_2 << 24))
	var tmp int32
	_ = tmp
	tmp = int32(v)
	var pos1 int32
	_ = pos1
	pos1 = (pos + 4)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = ((int32)(this.Hx_Field_b.Get((pos1 + 1))))
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = ((int32)(this.Hx_Field_b.Get((pos1 + 2))))
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	_hx_tmp_5 = ((int32)(this.Hx_Field_b.Get((pos1 + 3))))
	var v1 int32
	_ = v1
	v1 = (((((int32)(this.Hx_Field_b.Get(pos1))) | (_hx_tmp_3 << 8)) | (_hx_tmp_4 << 16)) | (_hx_tmp_5 << 24))
	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	_hx_tmp_6 = tmp
	return Hx_Field_haxe_io_fphelper_i64ToDouble(_hx_tmp_6, int32(v1))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getFloat(pos int32) float64 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this.Hx_Field_b.Get((pos + 1))))
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)(this.Hx_Field_b.Get((pos + 2))))
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int32)(this.Hx_Field_b.Get((pos + 3))))
	var v int32
	_ = v
	v = (((((int32)(this.Hx_Field_b.Get(pos))) | (_hx_tmp_0 << 8)) | (_hx_tmp_1 << 16)) | (_hx_tmp_2 << 24))
	return Hx_Field__single_single_impl__toFloat(Hx_Field_haxe_io_fphelper_i32ToFloat(int32(v)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_setDouble(pos int32, v float64) {
	var i Hx_Typedef_haxe__int64___int64
	_ = i
	i = Hx_Field_haxe_io_fphelper_doubleToI64(v)
	{
		var v1 int32
		_ = v1
		v1 = ((int32)(i))
		this.Hx_Field_b.Set(pos, byte((v1 & 255)))
		this.Hx_Field_b.Set((pos + 1), byte(((v1 >> 8) & 255)))
		this.Hx_Field_b.Set((pos + 2), byte(((v1 >> 16) & 255)))
		this.Hx_Field_b.Set((pos + 3), byte((((int32)((((uint32)(v1)) >> 24))) & 255)))
	}

	{
		var pos1 int32
		_ = pos1
		pos1 = (pos + 4)
		var _hx_tmp_0 int64
		_ = _hx_tmp_0
		_hx_tmp_0 = ((int64)(i))
		var v2 int32
		_ = v2
		v2 = ((int32)((_hx_tmp_0 >> int64(((int32)(32))))))
		this.Hx_Field_b.Set(pos1, byte((v2 & 255)))
		this.Hx_Field_b.Set((pos1 + 1), byte(((v2 >> 8) & 255)))
		this.Hx_Field_b.Set((pos1 + 2), byte(((v2 >> 16) & 255)))
		this.Hx_Field_b.Set((pos1 + 3), byte((((int32)((((uint32)(v2)) >> 24))) & 255)))
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_setFloat(pos int32, v float64) {
	var v1 int32
	_ = v1
	v1 = Hx_Field_haxe_io_fphelper_floatToI32(Hx_Field__single_single_impl__fromFloat(v))
	this.Hx_Field_b.Set(pos, byte((v1 & 255)))
	this.Hx_Field_b.Set((pos + 1), byte(((v1 >> 8) & 255)))
	this.Hx_Field_b.Set((pos + 2), byte(((v1 >> 16) & 255)))
	this.Hx_Field_b.Set((pos + 3), byte((((int32)((((uint32)(v1)) >> 24))) & 255)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getUInt16(pos int32) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this.Hx_Field_b.Get((pos + 1))))
	return (((int32)(this.Hx_Field_b.Get(pos))) | (_hx_tmp_0 << 8))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_setUInt16(pos int32, v int32) {
	this.Hx_Field_b.Set(pos, byte((v & 255)))
	this.Hx_Field_b.Set((pos + 1), byte(((v >> 8) & 255)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getInt32(pos int32) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this.Hx_Field_b.Get((pos + 1))))
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)(this.Hx_Field_b.Get((pos + 2))))
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int32)(this.Hx_Field_b.Get((pos + 3))))
	var v int32
	_ = v
	v = (((((int32)(this.Hx_Field_b.Get(pos))) | (_hx_tmp_0 << 8)) | (_hx_tmp_1 << 16)) | (_hx_tmp_2 << 24))
	return int32(v)
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getInt64(pos int32) Hx_Typedef_haxe__int64___int64 {
	var pos1 int32
	_ = pos1
	pos1 = (pos + 4)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this.Hx_Field_b.Get((pos1 + 1))))
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)(this.Hx_Field_b.Get((pos1 + 2))))
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int32)(this.Hx_Field_b.Get((pos1 + 3))))
	var v int32
	_ = v
	v = (((((int32)(this.Hx_Field_b.Get(pos1))) | (_hx_tmp_0 << 8)) | (_hx_tmp_1 << 16)) | (_hx_tmp_2 << 24))
	var high int32
	_ = high
	high = int32(v)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = ((int32)(this.Hx_Field_b.Get((pos + 1))))
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = ((int32)(this.Hx_Field_b.Get((pos + 2))))
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	_hx_tmp_5 = ((int32)(this.Hx_Field_b.Get((pos + 3))))
	var v1 int32
	_ = v1
	v1 = (((((int32)(this.Hx_Field_b.Get(pos))) | (_hx_tmp_3 << 8)) | (_hx_tmp_4 << 16)) | (_hx_tmp_5 << 24))
	var low int32
	_ = low
	low = int32(v1)
	var h int64
	_ = h
	h = int64(high)
	var l int64
	_ = l
	l = int64(low)
	var _hx_tmp_6 int64
	_ = _hx_tmp_6
	_hx_tmp_6 = h
	return ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_6 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_setInt32(pos int32, v int32) {
	this.Hx_Field_b.Set(pos, byte((v & 255)))
	this.Hx_Field_b.Set((pos + 1), byte(((v >> 8) & 255)))
	this.Hx_Field_b.Set((pos + 2), byte(((v >> 16) & 255)))
	this.Hx_Field_b.Set((pos + 3), byte((((int32)((((uint32)(v)) >> 24))) & 255)))
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_setInt64(pos int32, v Hx_Typedef_haxe__int64___int64) {
	{
		var v1 int32
		_ = v1
		v1 = ((int32)(v))
		this.Hx_Field_b.Set(pos, byte((v1 & 255)))
		this.Hx_Field_b.Set((pos + 1), byte(((v1 >> 8) & 255)))
		this.Hx_Field_b.Set((pos + 2), byte(((v1 >> 16) & 255)))
		this.Hx_Field_b.Set((pos + 3), byte((((int32)((((uint32)(v1)) >> 24))) & 255)))
	}

	{
		var pos1 int32
		_ = pos1
		pos1 = (pos + 4)
		var _hx_tmp_0 int64
		_ = _hx_tmp_0
		_hx_tmp_0 = ((int64)(v))
		var v2 int32
		_ = v2
		v2 = ((int32)((_hx_tmp_0 >> int64(((int32)(32))))))
		this.Hx_Field_b.Set(pos1, byte((v2 & 255)))
		this.Hx_Field_b.Set((pos1 + 1), byte(((v2 >> 8) & 255)))
		this.Hx_Field_b.Set((pos1 + 2), byte(((v2 >> 16) & 255)))
		this.Hx_Field_b.Set((pos1 + 3), byte((((int32)((((uint32)(v2)) >> 24))) & 255)))
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getString(pos int32, _hx_reserved_len int32, encoding__isOptional any) string {
	encoding := HxDefault[Hx_Enum_haxe_io_encoding]()
	_ = encoding

	if encoding__isOptional != nil {
		encoding = HxConvert[Hx_Enum_haxe_io_encoding](encoding__isOptional)
	} else {
		encoding = nil
	}
	{
		if ((pos < 0) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > this.Hx_Field_length) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
		}

		var self HxArray[byte]
		_ = self
		self = HxAnyToTypedArray[byte](this.Hx_Field_b)
		var _hx_tmp_0 []byte
		_ = _hx_tmp_0
		_hx_tmp_0 = self.Underlying()
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = int(pos)
		return string(_hx_tmp_0[_hx_tmp_1:int((pos + _hx_reserved_len))])
	}
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_readString(pos int32, _hx_reserved_len int32) string {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = pos
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = _hx_reserved_len
	return this.VTable.Hx_Field_getString(_hx_tmp_0, _hx_tmp_1, nil)
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_toString() string {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_length
	return this.VTable.Hx_Field_getString(0, _hx_tmp_0, nil)
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_toHex() string {
	var s *Hx_Obj_stringbuf
	_ = s
	s = Hx_Obj_stringbuf_CreateInstance()
	var chars HxArray[HxNullable[int32]]
	_ = chars
	chars = HxMakeArray[HxNullable[int32]]()
	var str string
	_ = str
	str = "0123456789abcdef"
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(str))
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var _hx_tmp_1 HxArray[HxNullable[int32]]
			_ = _hx_tmp_1
			_hx_tmp_1 = chars
			_ = (Hx_Array_Push[HxNullable[int32]](_hx_tmp_1, (Hx_Field_go_haxe_hxstring_charCodeAt(str, i))))
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = this.Hx_Field_length
		for _g2 < _g3 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			var c int32
			_ = c
			c = ((int32)(this.Hx_Field_b.Get(i)))
			s.VTable.Hx_Field_addChar((chars.Get((c >> 4))).Value)
			s.VTable.Hx_Field_addChar((chars.Get((c & 15))).Value)
		}
	}

	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_io_bytes) Hx_Field_getData() Hx_Typedef_haxe_io_bytesdata {
	return this.Hx_Field_b
}

func Hx_Field_haxe_io_bytes_alloc(length int32) *Hx_Obj_haxe_io_bytes {
	var length1 HxNullable[int32]
	_ = length1
	length1 = HxNullable[int32]{Value: length, Valid: true}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length1).Valid != false {
		_hx_tmp_0 = (length1).Value
	} else {
		_hx_tmp_0 = 0
	}

	var self []byte
	_ = self
	self = make([]byte, _hx_tmp_0)
	return Hx_Obj_haxe_io_bytes_CreateInstance(length, ((Hx_Typedef_haxe_io_bytesdata)(HxMakeArray(self...))))
}

func Hx_Field_haxe_io_bytes_ofString(s string, encoding__isOptional any) *Hx_Obj_haxe_io_bytes {
	encoding := HxDefault[Hx_Enum_haxe_io_encoding]()
	_ = encoding

	if encoding__isOptional != nil {
		encoding = HxConvert[Hx_Enum_haxe_io_encoding](encoding__isOptional)
	} else {
		encoding = nil
	}
	{
		var self []byte
		_ = self
		self = (([]byte)(s))
		var self1 []byte
		_ = self1
		self1 = self
		return Hx_Field_haxe_io_bytes_ofData(((Hx_Typedef_haxe_io_bytesdata)(HxMakeArray(self1...))))
	}
}

func Hx_Field_haxe_io_bytes_ofData(b Hx_Typedef_haxe_io_bytesdata) *Hx_Obj_haxe_io_bytes {
	return Hx_Obj_haxe_io_bytes_CreateInstance(b.Len(), b)
}

func Hx_Field_haxe_io_bytes_ofHex(s string) *Hx_Obj_haxe_io_bytes {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(s))
	if (_hx_reserved_len & 1) != 0 {
		panic(Hx_Field_haxe_exception_thrown(((any)("Not a hex string (odd number of digits)"))))
	}

	var ret *Hx_Obj_haxe_io_bytes
	_ = ret
	ret = Hx_Field_haxe_io_bytes_alloc((_hx_reserved_len >> 1))
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ret.Hx_Field_length
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var high int32
			_ = high
			high = (Hx_Field_go_haxe_hxstring_charCodeAt(s, (i * 2))).Value
			var low int32
			_ = low
			low = (Hx_Field_go_haxe_hxstring_charCodeAt(s, ((i * 2) + 1))).Value
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = (high & 64)
			high = ((high & 15) + ((_hx_tmp_1 >> 6) * 9))
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = (low & 64)
			low = ((low & 15) + ((_hx_tmp_2 >> 6) * 9))
			ret.Hx_Field_b.Set(i, byte(((((high << 4) | low) & 255) & 255)))
		}
	}

	return ret
}

func Hx_Field_haxe_io_bytes_fastGet(b Hx_Typedef_haxe_io_bytesdata, pos int32) int32 {
	return ((int32)(b.Get(pos)))
}

func init() {
	hxRegisterArrayConstructor[HxNullable[int32]]()
	hxRegisterArrayConstructor[byte]()
}
