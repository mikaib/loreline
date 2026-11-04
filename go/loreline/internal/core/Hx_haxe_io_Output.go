package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_output() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_output() {
}

type Hx_Obj_VTable_haxe_io_output interface {
	Hx_Field_close()
	Hx_Field_writeDouble(x float64)
	Hx_Field_writeInt24(x int32)
	Hx_Field_writeUInt24(x int32)
	Hx_Field_writeBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) int32
	Hx_Field_write(s *Hx_Obj_haxe_io_bytes)
	Hx_Field_writeInt8(x int32)
	Hx_Field_writeInt32(x int32)
	Hx_Field_writeInput(i *Hx_Obj_haxe_io_input, bufsize__isOptional any)
	Hx_Field_writeFloat(x float64)
	Hx_Field_writeUInt16(x int32)
	Hx_Field_writeByte(c int32)
	Hx_Field_flush()
	Hx_Field_set_bigEndian(b bool) bool
	Hx_Field_writeFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32)
	Hx_Field_writeInt16(x int32)
	Hx_Field_prepare(nbytes int32)
	Hx_Field_writeString(s string, encoding__isOptional any)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_output struct {
	VTable             Hx_Obj_VTable_haxe_io_output `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_bigEndian bool
}

func Hx_Obj_haxe_io_output_CreateEmptyInstance() *Hx_Obj_haxe_io_output {
	obj := &Hx_Obj_haxe_io_output{}
	obj.VTable = obj
	obj.Hx_Field_bigEndian = HxDefault[bool]()
	return obj
}

func Hx_Obj_haxe_io_output_CreateInstance() *Hx_Obj_haxe_io_output {
	obj := Hx_Obj_haxe_io_output_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_io_output_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.Output",
	HxMakeArray[string](),
	HxMakeArray[string]("bigEndian", "writeByte", "writeBytes", "flush", "close", "set_bigEndian", "write", "writeFullBytes", "writeFloat", "writeDouble", "writeInt8", "writeInt16", "writeUInt16", "writeInt24", "writeUInt24", "writeInt32", "prepare", "writeInput", "writeString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_output_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_io_output_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_output_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_output_RTTI
}

func (this *Hx_Obj_haxe_io_output) Hx_As_Hx_Obj_haxe_io_output() *Hx_Obj_haxe_io_output {
	return this
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeByte(c int32) {
	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance(nil, nil, any(map[string]any{"fileName": ((any)("haxe/io/Output.hx")), "lineNumber": ((int32)(49)), "className": ((any)("haxe.io.Output")), "methodName": ((any)("writeByte"))}))))))
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) int32 {
	if ((pos < 0) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > s.Hx_Field_length) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
	}

	var b Hx_Typedef_haxe_io_bytesdata
	_ = b
	b = s.Hx_Field_b
	var k int32
	_ = k
	k = _hx_reserved_len
	for k > 0 {
		this.VTable.Hx_Field_writeByte(((int32)(b.Get(pos))))
		pos++
		k--
	}

	return _hx_reserved_len
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_flush() {
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_close() {
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_set_bigEndian(b bool) bool {
	this.Hx_Field_bigEndian = b
	return b
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_write(s *Hx_Obj_haxe_io_bytes) {
	var l int32
	_ = l
	l = s.Hx_Field_length
	var p int32
	_ = p
	p = 0
	for l > 0 {
		var k int32
		_ = k
		k = this.VTable.Hx_Field_writeBytes(s, p, l)
		if k == 0 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
		}

		p = (p + k)
		l = (l - k)
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) {
	for _hx_reserved_len > 0 {
		var k int32
		_ = k
		k = this.VTable.Hx_Field_writeBytes(s, pos, _hx_reserved_len)
		pos = (pos + k)
		_hx_reserved_len = (_hx_reserved_len - k)
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeFloat(x float64) {
	this.VTable.Hx_Field_writeInt32(Hx_Field_haxe_io_fphelper_floatToI32(Hx_Field__single_single_impl__fromFloat(x)))
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeDouble(x float64) {
	var i64 Hx_Typedef_haxe__int64___int64
	_ = i64
	i64 = Hx_Field_haxe_io_fphelper_doubleToI64(x)
	if this.Hx_Field_bigEndian {
		var _hx_tmp_0 int64
		_ = _hx_tmp_0
		_hx_tmp_0 = ((int64)(i64))
		this.VTable.Hx_Field_writeInt32(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))))
		this.VTable.Hx_Field_writeInt32(((int32)(i64)))
	} else {
		this.VTable.Hx_Field_writeInt32(((int32)(i64)))
		var _hx_tmp_1 int64
		_ = _hx_tmp_1
		_hx_tmp_1 = ((int64)(i64))
		this.VTable.Hx_Field_writeInt32(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))))
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeInt8(x int32) {
	if (x < -128) || (x >= 128) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Overflow{}))))
	}

	this.VTable.Hx_Field_writeByte((x & 255))
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeInt16(x int32) {
	if (x < -32768) || (x >= 32768) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Overflow{}))))
	}

	this.VTable.Hx_Field_writeUInt16((x & 65535))
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeUInt16(x int32) {
	if (x < 0) || (x >= 65536) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Overflow{}))))
	}

	if this.Hx_Field_bigEndian {
		this.VTable.Hx_Field_writeByte((x >> 8))
		this.VTable.Hx_Field_writeByte((x & 255))
	} else {
		this.VTable.Hx_Field_writeByte((x & 255))
		this.VTable.Hx_Field_writeByte((x >> 8))
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeInt24(x int32) {
	if (x < -8388608) || (x >= 8388608) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Overflow{}))))
	}

	this.VTable.Hx_Field_writeUInt24((x & 16777215))
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeUInt24(x int32) {
	if (x < 0) || (x >= 16777216) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Overflow{}))))
	}

	if this.Hx_Field_bigEndian {
		this.VTable.Hx_Field_writeByte((x >> 16))
		this.VTable.Hx_Field_writeByte(((x >> 8) & 255))
		this.VTable.Hx_Field_writeByte((x & 255))
	} else {
		this.VTable.Hx_Field_writeByte((x & 255))
		this.VTable.Hx_Field_writeByte(((x >> 8) & 255))
		this.VTable.Hx_Field_writeByte((x >> 16))
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeInt32(x int32) {
	if this.Hx_Field_bigEndian {
		this.VTable.Hx_Field_writeByte(((int32)((((uint32)(x)) >> 24))))
		this.VTable.Hx_Field_writeByte(((x >> 16) & 255))
		this.VTable.Hx_Field_writeByte(((x >> 8) & 255))
		this.VTable.Hx_Field_writeByte((x & 255))
	} else {
		this.VTable.Hx_Field_writeByte((x & 255))
		this.VTable.Hx_Field_writeByte(((x >> 8) & 255))
		this.VTable.Hx_Field_writeByte(((x >> 16) & 255))
		this.VTable.Hx_Field_writeByte(((int32)((((uint32)(x)) >> 24))))
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_prepare(nbytes int32) {
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeInput(i *Hx_Obj_haxe_io_input, bufsize__isOptional any) {
	bufsize := HxDefault[HxNullable[int32]]()
	_ = bufsize

	if bufsize__isOptional != nil {
		bufsize = HxConvert[HxNullable[int32]](bufsize__isOptional)
	} else {

	}
	{
		if (bufsize).Valid == false {
			var hx_nullable_t51 HxNullable[int32]
			_ = hx_nullable_t51
			hx_nullable_t51 = HxNullable[int32]{}

			bufsize = HxNullable[int32]{Value: (hx_nullable_t51).Value, Valid: (hx_nullable_t51).Valid}
		}

		if (bufsize).Valid == false {
			bufsize = HxNullable[int32]{Value: 4096, Valid: true}
		}

		var buf *Hx_Obj_haxe_io_bytes
		_ = buf
		buf = Hx_Field_haxe_io_bytes_alloc((bufsize).Value)
		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_io_eof").(*Hx_Obj_haxe_io_eof); _hx_ok {
							_ = e
							{
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					for {
						var _hx_tmp_0 *Hx_Obj_haxe_io_bytes
						_ = _hx_tmp_0
						_hx_tmp_0 = buf
						var _hx_reserved_len int32
						_ = _hx_reserved_len
						_hx_reserved_len = i.VTable.Hx_Field_readBytes(_hx_tmp_0, 0, (bufsize).Value)
						if _hx_reserved_len == 0 {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
						}

						var p int32
						_ = p
						p = 0
						for _hx_reserved_len > 0 {
							var k int32
							_ = k
							k = this.VTable.Hx_Field_writeBytes(buf, p, _hx_reserved_len)
							if k == 0 {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
							}

							p = (p + k)
							_hx_reserved_len = (_hx_reserved_len - k)
						}
					}
				}
			}()
			if hx_try_state == 1 {
				return
			}
		}
	}
}

func (this *Hx_Obj_haxe_io_output) Hx_Field_writeString(s string, encoding__isOptional any) {
	encoding := HxDefault[Hx_Enum_haxe_io_encoding]()
	_ = encoding

	if encoding__isOptional != nil {
		encoding = HxConvert[Hx_Enum_haxe_io_encoding](encoding__isOptional)
	} else {
		encoding = nil
	}
	{
		var b *Hx_Obj_haxe_io_bytes
		_ = b
		b = Hx_Field_haxe_io_bytes_ofString(s, encoding)
		this.VTable.Hx_Field_writeFullBytes(b, 0, b.Hx_Field_length)
	}
}
