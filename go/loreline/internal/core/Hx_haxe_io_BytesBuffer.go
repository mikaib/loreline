package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_bytesbuffer() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_bytesbuffer() {
}

type Hx_Obj_VTable_haxe_io_bytesbuffer interface {
	Hx_Field_get_length() int32
	Hx_Field_addByte(_hx_reserved_byte int32)
	Hx_Field_add(src *Hx_Obj_haxe_io_bytes)
	Hx_Field_addInt64(v Hx_Typedef_haxe__int64___int64)
	Hx_Field_addBytes(src *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32)
	Hx_Field_addString(v string, encoding__isOptional any)
	Hx_Field_addInt32(v int32)
	Hx_Field_addFloat(v float64)
	Hx_Field_addDouble(v float64)
	Hx_Field_getBytes() *Hx_Obj_haxe_io_bytes
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_bytesbuffer struct {
	VTable     Hx_Obj_VTable_haxe_io_bytesbuffer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_b Hx_Typedef_haxe_io_bytesdata
}

func Hx_Obj_haxe_io_bytesbuffer_CreateEmptyInstance() *Hx_Obj_haxe_io_bytesbuffer {
	obj := &Hx_Obj_haxe_io_bytesbuffer{}
	obj.VTable = obj
	obj.Hx_Field_b = HxDefault[Hx_Typedef_haxe_io_bytesdata]()
	return obj
}

func Hx_Obj_haxe_io_bytesbuffer_CreateInstance() *Hx_Obj_haxe_io_bytesbuffer {
	obj := Hx_Obj_haxe_io_bytesbuffer_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_New() {
	{
		this.Hx_Field_b = ((Hx_Typedef_haxe_io_bytesdata)(HxMakeArray[byte]()))
	}

}

var Hx_Obj_haxe_io_bytesbuffer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.BytesBuffer",
	HxMakeArray[string](),
	HxMakeArray[string]("b", "get_length", "addByte", "add", "addString", "addInt32", "addInt64", "addFloat", "addDouble", "addBytes", "getBytes"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_bytesbuffer_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_io_bytesbuffer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_bytesbuffer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_bytesbuffer_RTTI
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_As_Hx_Obj_haxe_io_bytesbuffer() *Hx_Obj_haxe_io_bytesbuffer {
	return this
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_get_length() int32 {
	return this.Hx_Field_b.Len()
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addByte(_hx_reserved_byte int32) {
	var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_0, byte(_hx_reserved_byte)))
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_add(src *Hx_Obj_haxe_io_bytes) {
	var b2 Hx_Typedef_haxe_io_bytesdata
	_ = b2
	b2 = src.Hx_Field_b
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = src.Hx_Field_length
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			_ = (Hx_Array_Push[byte](this.Hx_Field_b, b2.Get(i)))
		}
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addString(v string, encoding__isOptional any) {
	encoding := HxDefault[Hx_Enum_haxe_io_encoding]()
	_ = encoding

	if encoding__isOptional != nil {
		encoding = HxConvert[Hx_Enum_haxe_io_encoding](encoding__isOptional)
	} else {
		encoding = nil
	}
	{
		var src *Hx_Obj_haxe_io_bytes
		_ = src
		src = Hx_Field_haxe_io_bytes_ofString(v, encoding)
		var b2 Hx_Typedef_haxe_io_bytesdata
		_ = b2
		b2 = src.Hx_Field_b
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = src.Hx_Field_length
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				_ = (Hx_Array_Push[byte](this.Hx_Field_b, b2.Get(i)))
			}
		}
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addInt32(v int32) {
	var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_0, byte((v & 255))))
	var _hx_tmp_1 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_1, byte(((v >> 8) & 255))))
	var _hx_tmp_2 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_2, byte(((v >> 16) & 255))))
	var _hx_tmp_3 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_3
	_hx_tmp_3 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_3, byte(((int32)((((uint32)(v)) >> 24))))))
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addInt64(v Hx_Typedef_haxe__int64___int64) {
	{
		var v1 int32
		_ = v1
		v1 = ((int32)(v))
		var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_0, byte((v1 & 255))))
		var _hx_tmp_1 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_1, byte(((v1 >> 8) & 255))))
		var _hx_tmp_2 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_2
		_hx_tmp_2 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_2, byte(((v1 >> 16) & 255))))
		var _hx_tmp_3 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_3
		_hx_tmp_3 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_3, byte(((int32)((((uint32)(v1)) >> 24))))))
	}

	{
		var _hx_tmp_4 int64
		_ = _hx_tmp_4
		_hx_tmp_4 = ((int64)(v))
		var v2 int32
		_ = v2
		v2 = ((int32)((_hx_tmp_4 >> int64(((int32)(32))))))
		var _hx_tmp_5 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_5
		_hx_tmp_5 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_5, byte((v2 & 255))))
		var _hx_tmp_6 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_6
		_hx_tmp_6 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_6, byte(((v2 >> 8) & 255))))
		var _hx_tmp_7 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_7
		_hx_tmp_7 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_7, byte(((v2 >> 16) & 255))))
		var _hx_tmp_8 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_8
		_hx_tmp_8 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_8, byte(((int32)((((uint32)(v2)) >> 24))))))
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addFloat(v float64) {
	var v1 int32
	_ = v1
	v1 = Hx_Field_haxe_io_fphelper_floatToI32(Hx_Field__single_single_impl__fromFloat(v))
	var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_0, byte((v1 & 255))))
	var _hx_tmp_1 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_1, byte(((v1 >> 8) & 255))))
	var _hx_tmp_2 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_2, byte(((v1 >> 16) & 255))))
	var _hx_tmp_3 Hx_Typedef_haxe_io_bytesdata
	_ = _hx_tmp_3
	_hx_tmp_3 = this.Hx_Field_b
	_ = (Hx_Array_Push[byte](_hx_tmp_3, byte(((int32)((((uint32)(v1)) >> 24))))))
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addDouble(v float64) {
	var v1 Hx_Typedef_haxe__int64___int64
	_ = v1
	v1 = Hx_Field_haxe_io_fphelper_doubleToI64(v)
	{
		var v2 int32
		_ = v2
		v2 = ((int32)(v1))
		var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_0, byte((v2 & 255))))
		var _hx_tmp_1 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_1, byte(((v2 >> 8) & 255))))
		var _hx_tmp_2 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_2
		_hx_tmp_2 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_2, byte(((v2 >> 16) & 255))))
		var _hx_tmp_3 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_3
		_hx_tmp_3 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_3, byte(((int32)((((uint32)(v2)) >> 24))))))
	}

	{
		var _hx_tmp_4 int64
		_ = _hx_tmp_4
		_hx_tmp_4 = ((int64)(v1))
		var v3 int32
		_ = v3
		v3 = ((int32)((_hx_tmp_4 >> int64(((int32)(32))))))
		var _hx_tmp_5 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_5
		_hx_tmp_5 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_5, byte((v3 & 255))))
		var _hx_tmp_6 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_6
		_hx_tmp_6 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_6, byte(((v3 >> 8) & 255))))
		var _hx_tmp_7 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_7
		_hx_tmp_7 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_7, byte(((v3 >> 16) & 255))))
		var _hx_tmp_8 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_8
		_hx_tmp_8 = this.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_8, byte(((int32)((((uint32)(v3)) >> 24))))))
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_addBytes(src *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) {
	if ((pos < 0) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > src.Hx_Field_length) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
	}

	var b2 Hx_Typedef_haxe_io_bytesdata
	_ = b2
	b2 = src.Hx_Field_b
	{
		var _g int32
		_ = _g
		_g = pos
		var _g1 int32
		_ = _g1
		_g1 = (pos + _hx_reserved_len)
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			_ = (Hx_Array_Push[byte](this.Hx_Field_b, b2.Get(i)))
		}
	}
}

func (this *Hx_Obj_haxe_io_bytesbuffer) Hx_Field_getBytes() *Hx_Obj_haxe_io_bytes {
	var bytes *Hx_Obj_haxe_io_bytes
	_ = bytes
	bytes = Hx_Field_haxe_io_bytes_ofData(this.Hx_Field_b)
	this.Hx_Field_b = nil
	return bytes
}

func init() {
	hxRegisterArrayConstructor[byte]()
}
