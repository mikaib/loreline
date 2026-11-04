package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_input() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_input() {
}

type Hx_Obj_VTable_haxe_io_input interface {
	Hx_Field_set_bigEndian(b bool) bool
	Hx_Field_readInt16() int32
	Hx_Field_readByte() int32
	Hx_Field_readBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) int32
	Hx_Field_readAll(bufsize__isOptional any) *Hx_Obj_haxe_io_bytes
	Hx_Field_readFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32)
	Hx_Field_readDouble() float64
	Hx_Field_readInt8() int32
	Hx_Field_readUInt16() int32
	Hx_Field_readUInt24() int32
	Hx_Field_close()
	Hx_Field_read(nbytes int32) *Hx_Obj_haxe_io_bytes
	Hx_Field_readLine() string
	Hx_Field_readFloat() float64
	Hx_Field_readInt24() int32
	Hx_Field_readString(_hx_reserved_len int32, encoding__isOptional any) string
	Hx_Field_readUntil(end int32) string
	Hx_Field_readInt32() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_input struct {
	VTable             Hx_Obj_VTable_haxe_io_input `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_bigEndian bool
}

func Hx_Obj_haxe_io_input_CreateEmptyInstance() *Hx_Obj_haxe_io_input {
	obj := &Hx_Obj_haxe_io_input{}
	obj.VTable = obj
	obj.Hx_Field_bigEndian = HxDefault[bool]()
	return obj
}

func Hx_Obj_haxe_io_input_CreateInstance() *Hx_Obj_haxe_io_input {
	obj := Hx_Obj_haxe_io_input_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_io_input_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.Input",
	HxMakeArray[string](),
	HxMakeArray[string]("bigEndian", "readByte", "readBytes", "close", "set_bigEndian", "readAll", "readFullBytes", "read", "readUntil", "readLine", "readFloat", "readDouble", "readInt8", "readInt16", "readUInt16", "readInt24", "readUInt24", "readInt32", "readString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_input_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_io_input_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_input_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_input_RTTI
}

func (this *Hx_Obj_haxe_io_input) Hx_As_Hx_Obj_haxe_io_input() *Hx_Obj_haxe_io_input {
	return this
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readByte() int32 {
	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_exceptions_notimplementedexception_CreateInstance(nil, nil, any(map[string]any{"fileName": ((any)("haxe/io/Input.hx")), "lineNumber": ((int32)(53)), "className": ((any)("haxe.io.Input")), "methodName": ((any)("readByte"))}))))))
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) int32 {
	var k int32
	_ = k
	k = _hx_reserved_len
	var b Hx_Typedef_haxe_io_bytesdata
	_ = b
	b = s.Hx_Field_b
	if ((pos < 0) || (_hx_reserved_len < 0)) || ((pos + _hx_reserved_len) > s.Hx_Field_length) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
	}

	{
		var hx_try_return int32
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if eof, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_io_eof").(*Hx_Obj_haxe_io_eof); _hx_ok {
						_ = eof
						{
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				for k > 0 {
					b.Set(pos, ((byte)(this.VTable.Hx_Field_readByte())))
					pos++
					k--
				}
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	return (_hx_reserved_len - k)
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_close() {
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_set_bigEndian(b bool) bool {
	this.Hx_Field_bigEndian = b
	return b
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readAll(bufsize__isOptional any) *Hx_Obj_haxe_io_bytes {
	bufsize := HxDefault[HxNullable[int32]]()
	_ = bufsize

	if bufsize__isOptional != nil {
		bufsize = HxConvert[HxNullable[int32]](bufsize__isOptional)
	} else {

	}
	{
		if (bufsize).Valid == false {
			var hx_nullable_t50 HxNullable[int32]
			_ = hx_nullable_t50
			hx_nullable_t50 = HxNullable[int32]{}

			bufsize = HxNullable[int32]{Value: (hx_nullable_t50).Value, Valid: (hx_nullable_t50).Valid}
		}

		if (bufsize).Valid == false {
			bufsize = HxNullable[int32]{Value: 16384, Valid: true}
		}

		var buf *Hx_Obj_haxe_io_bytes
		_ = buf
		buf = Hx_Field_haxe_io_bytes_alloc((bufsize).Value)
		var total *Hx_Obj_haxe_io_bytesbuffer
		_ = total
		total = Hx_Obj_haxe_io_bytesbuffer_CreateInstance()
		{
			var hx_try_return *Hx_Obj_haxe_io_bytes
			_ = hx_try_return
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
						_hx_reserved_len = this.VTable.Hx_Field_readBytes(_hx_tmp_0, 0, (bufsize).Value)
						if _hx_reserved_len == 0 {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
						}

						{
							if (_hx_reserved_len < 0) || (_hx_reserved_len > buf.Hx_Field_length) {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_OutsideBounds{}))))
							}

							var b2 Hx_Typedef_haxe_io_bytesdata
							_ = b2
							b2 = buf.Hx_Field_b
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
									_ = (Hx_Array_Push[byte](total.Hx_Field_b, b2.Get(i)))
								}
							}
						}
					}
				}
			}()
			if hx_try_state == 1 {
				return hx_try_return
			}
		}

		return total.VTable.Hx_Field_getBytes()
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32) {
	for _hx_reserved_len > 0 {
		var k int32
		_ = k
		k = this.VTable.Hx_Field_readBytes(s, pos, _hx_reserved_len)
		if k == 0 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
		}

		pos = (pos + k)
		_hx_reserved_len = (_hx_reserved_len - k)
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_read(nbytes int32) *Hx_Obj_haxe_io_bytes {
	var s *Hx_Obj_haxe_io_bytes
	_ = s
	s = Hx_Field_haxe_io_bytes_alloc(nbytes)
	var p int32
	_ = p
	p = 0
	for nbytes > 0 {
		var k int32
		_ = k
		k = this.VTable.Hx_Field_readBytes(s, p, nbytes)
		if k == 0 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Blocked{}))))
		}

		p = (p + k)
		nbytes = (nbytes - k)
	}

	return s
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readUntil(end int32) string {
	var buf *Hx_Obj_haxe_io_bytesbuffer
	_ = buf
	buf = Hx_Obj_haxe_io_bytesbuffer_CreateInstance()
	var last int32
	_ = last
	for {
		last = this.VTable.Hx_Field_readByte()
		if last == end {
			break
		}

		var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
		_ = _hx_tmp_0
		_hx_tmp_0 = buf.Hx_Field_b
		_ = (Hx_Array_Push[byte](_hx_tmp_0, byte(last)))
	}

	return buf.VTable.Hx_Field_getBytes().VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readLine() string {
	var buf *Hx_Obj_haxe_io_bytesbuffer
	_ = buf
	buf = Hx_Obj_haxe_io_bytesbuffer_CreateInstance()
	var last int32
	_ = last
	var s string
	_ = s
	{
		var hx_try_return string
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_io_eof").(*Hx_Obj_haxe_io_eof); _hx_ok {
						_ = e
						{
							s = buf.VTable.Hx_Field_getBytes().VTable.Hx_Field_toString()
							if int32(utf8.RuneCountInString(s)) == 0 {
								panic(Hx_Field_haxe_exception_thrown(((any)(e))))
							}
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				for {
					last = this.VTable.Hx_Field_readByte()
					if last == 10 {
						break
					}

					var _hx_tmp_0 Hx_Typedef_haxe_io_bytesdata
					_ = _hx_tmp_0
					_hx_tmp_0 = buf.Hx_Field_b
					_ = (Hx_Array_Push[byte](_hx_tmp_0, byte(last)))
				}

				s = buf.VTable.Hx_Field_getBytes().VTable.Hx_Field_toString()
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				_hx_tmp_1 = s
				var hx_cmp_13 HxNullable[int32]
				_ = hx_cmp_13
				hx_cmp_13 = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, (int32(utf8.RuneCountInString(s)) - 1)))
				var _hx_tmp_2 bool
				_ = _hx_tmp_2
				if (hx_cmp_13).Valid == true {
					_hx_tmp_2 = ((hx_cmp_13).Value == 13)
				} else {
					_hx_tmp_2 = false
				}

				if _hx_tmp_2 {
					var _hx_tmp_3 string
					_ = _hx_tmp_3
					_hx_tmp_3 = s
					s = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: -1, Valid: true}))
				}
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	return s
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readFloat() float64 {
	return Hx_Field__single_single_impl__toFloat(Hx_Field_haxe_io_fphelper_i32ToFloat(this.VTable.Hx_Field_readInt32()))
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readDouble() float64 {
	var i1 int32
	_ = i1
	i1 = this.VTable.Hx_Field_readInt32()
	var i2 int32
	_ = i2
	i2 = this.VTable.Hx_Field_readInt32()
	if this.Hx_Field_bigEndian {
		return Hx_Field_haxe_io_fphelper_i64ToDouble(i2, i1)
	} else {
		return Hx_Field_haxe_io_fphelper_i64ToDouble(i1, i2)
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readInt8() int32 {
	var n int32
	_ = n
	n = this.VTable.Hx_Field_readByte()
	if n >= 128 {
		return (n - 256)
	}

	return n
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readInt16() int32 {
	var ch1 int32
	_ = ch1
	ch1 = this.VTable.Hx_Field_readByte()
	var ch2 int32
	_ = ch2
	ch2 = this.VTable.Hx_Field_readByte()
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if this.Hx_Field_bigEndian {
		_hx_tmp_0 = (ch2 | (ch1 << 8))
	} else {
		_hx_tmp_0 = (ch1 | (ch2 << 8))
	}

	var n int32
	_ = n
	n = _hx_tmp_0
	if (n & 32768) != 0 {
		return (n - 65536)
	}

	return n
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readUInt16() int32 {
	var ch1 int32
	_ = ch1
	ch1 = this.VTable.Hx_Field_readByte()
	var ch2 int32
	_ = ch2
	ch2 = this.VTable.Hx_Field_readByte()
	if this.Hx_Field_bigEndian {
		return (ch2 | (ch1 << 8))
	} else {
		return (ch1 | (ch2 << 8))
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readInt24() int32 {
	var ch1 int32
	_ = ch1
	ch1 = this.VTable.Hx_Field_readByte()
	var ch2 int32
	_ = ch2
	ch2 = this.VTable.Hx_Field_readByte()
	var ch3 int32
	_ = ch3
	ch3 = this.VTable.Hx_Field_readByte()
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if this.Hx_Field_bigEndian {
		_hx_tmp_0 = ((ch3 | (ch2 << 8)) | (ch1 << 16))
	} else {
		_hx_tmp_0 = ((ch1 | (ch2 << 8)) | (ch3 << 16))
	}

	var n int32
	_ = n
	n = _hx_tmp_0
	if (n & 8388608) != 0 {
		return (n - 16777216)
	}

	return n
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readUInt24() int32 {
	var ch1 int32
	_ = ch1
	ch1 = this.VTable.Hx_Field_readByte()
	var ch2 int32
	_ = ch2
	ch2 = this.VTable.Hx_Field_readByte()
	var ch3 int32
	_ = ch3
	ch3 = this.VTable.Hx_Field_readByte()
	if this.Hx_Field_bigEndian {
		return ((ch3 | (ch2 << 8)) | (ch1 << 16))
	} else {
		return ((ch1 | (ch2 << 8)) | (ch3 << 16))
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readInt32() int32 {
	var ch1 int32
	_ = ch1
	ch1 = this.VTable.Hx_Field_readByte()
	var ch2 int32
	_ = ch2
	ch2 = this.VTable.Hx_Field_readByte()
	var ch3 int32
	_ = ch3
	ch3 = this.VTable.Hx_Field_readByte()
	var ch4 int32
	_ = ch4
	ch4 = this.VTable.Hx_Field_readByte()
	if this.Hx_Field_bigEndian {
		return (((ch4 | (ch3 << 8)) | (ch2 << 16)) | (ch1 << 24))
	} else {
		return (((ch1 | (ch2 << 8)) | (ch3 << 16)) | (ch4 << 24))
	}
}

func (this *Hx_Obj_haxe_io_input) Hx_Field_readString(_hx_reserved_len int32, encoding__isOptional any) string {
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
		b = Hx_Field_haxe_io_bytes_alloc(_hx_reserved_len)
		this.VTable.Hx_Field_readFullBytes(b, 0, _hx_reserved_len)
		return b.VTable.Hx_Field_getString(0, _hx_reserved_len, encoding)
	}
}
