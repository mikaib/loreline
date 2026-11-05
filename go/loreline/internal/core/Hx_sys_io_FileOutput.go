package core

import "os"

func Hx_ClassInit_Hx_Obj_VTable_sys_io_fileoutput() {
}

func Hx_StaticInit_Hx_Obj_VTable_sys_io_fileoutput() {
}

type Hx_Obj_VTable_sys_io_fileoutput interface {
	Hx_Field_writeFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32)
	Hx_Field_writeDouble(x float64)
	Hx_Field_writeInt32(x int32)
	Hx_Field_writeInput(i *Hx_Obj_haxe_io_input, bufsize__isOptional any)
	Hx_Field_writeFloat(x float64)
	Hx_Field_writeUInt16(x int32)
	Hx_Field_writeString(s string, encoding__isOptional any)
	Hx_Field_writeByte(c int32)
	Hx_Field_writeBytes(s *Hx_Obj_haxe_io_bytes, p int32, l int32) int32
	Hx_Field_seek(p int32, pos Hx_Enum_sys_io_fileseek)
	Hx_Field_tell() int32
	Hx_Field_write(s *Hx_Obj_haxe_io_bytes)
	Hx_Field_writeInt8(x int32)
	Hx_Field_writeUInt24(x int32)
	Hx_Field_prepare(nbytes int32)
	Hx_Field_close()
	Hx_Field_flush()
	Hx_Field_writeInt16(x int32)
	Hx_Field_writeInt24(x int32)
	Hx_Field_set_bigEndian(b bool) bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_sys_io_fileoutput struct {
	Hx_Obj_haxe_io_output
	VTable             Hx_Obj_VTable_sys_io_fileoutput `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field___f       *os.File
	Hx_Field___byteBuf []byte
}

func Hx_Obj_sys_io_fileoutput_CreateEmptyInstance() *Hx_Obj_sys_io_fileoutput {
	obj := &Hx_Obj_sys_io_fileoutput{}
	obj.Hx_Obj_haxe_io_output = *Hx_Obj_haxe_io_output_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_io_output.VTable = obj
	obj.Hx_Field___f = HxDefault[*os.File]()
	obj.Hx_Field___byteBuf = HxDefault[[]byte]()
	return obj
}

func Hx_Obj_sys_io_fileoutput_CreateInstance(f *os.File) *Hx_Obj_sys_io_fileoutput {
	obj := Hx_Obj_sys_io_fileoutput_CreateEmptyInstance()
	obj.Hx_New(f)
	return obj
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_New(f *os.File) {
	{
		this.Hx_Field___f = f
		var self HxArray[byte]
		_ = self
		self = HxMakeArray[byte](byte(((int32)(0))))
		this.Hx_Field___byteBuf = self.Underlying()
	}

}

var Hx_Obj_sys_io_fileoutput_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"sys.io.FileOutput",
	HxMakeArray[string](),
	HxMakeArray[string]("__f", "__byteBuf", "writeByte", "writeBytes", "close", "seek", "tell", "bigEndian", "flush", "set_bigEndian", "write", "writeFullBytes", "writeFloat", "writeDouble", "writeInt8", "writeInt16", "writeUInt16", "writeInt24", "writeUInt24", "writeInt32", "prepare", "writeInput", "writeString"),
	Hx_Obj_haxe_io_output_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_sys_io_fileoutput_CreateInstance(HxConvert[*os.File](params.Get(0)))
	},
	func() any {
		return Hx_Obj_sys_io_fileoutput_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_sys_io_fileoutput_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_sys_io_fileoutput_RTTI
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_As_Hx_Obj_sys_io_fileoutput() *Hx_Obj_sys_io_fileoutput {
	return this
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field_writeByte(c int32) {
	{
		var value byte
		_ = value
		value = ((byte)(c))
		this.Hx_Field___byteBuf[0] = value
	}

	{
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{e}))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_result_21 struct {
					Error  error
					Result int
				}
				_ = hx_result_21
				hx_result_21.Result, hx_result_21.Error = (*this.Hx_Field___f).Write(this.Hx_Field___byteBuf)
				var this1 struct {
					Error  error
					Result int
				}
				_ = this1
				this1 = hx_result_21
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
		}()
		if hx_try_state == 1 {
			return
		}
	}
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field_writeBytes(s *Hx_Obj_haxe_io_bytes, p int32, l int32) int32 {
	{
		var hx_try_return int32
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{e}))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var self HxArray[byte]
				_ = self
				self = HxAnyToTypedArray[byte](s.Hx_Field_b)
				var hx_result_22 struct {
					Error  error
					Result int
				}
				_ = hx_result_22
				hx_result_22.Result, hx_result_22.Error = (*this.Hx_Field___f).Write(self.Underlying())
				var this1 struct {
					Error  error
					Result int
				}
				_ = this1
				this1 = hx_result_22
				var this2 int
				_ = this2
				if this1.Error != nil {
					var e error
					_ = e
					e = this1.Error
					panic(Hx_Field_haxe_exception_thrown(((any)(e))))
				} else {
					var r int
					_ = r
					r = this1.Result
					this2 = r
				}

				hx_try_state = 1
				hx_try_return = ((int32)(this2))
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field_close() {
	this.Hx_Obj_haxe_io_output.Hx_Field_close()
	{
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{e}))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var this1 error
				_ = this1
				this1 = (*this.Hx_Field___f).Close()
				if this1 != nil {
					panic(Hx_Field_haxe_exception_thrown(((any)(this1.Error()))))
				}
			}
		}()
		if hx_try_state == 1 {
			return
		}
	}
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field_seek(p int32, pos Hx_Enum_sys_io_fileseek) {
	{
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{e}))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_result_23 struct {
					Error  error
					Result int64
				}
				_ = hx_result_23
				var _hx_tmp_0 int64
				_ = _hx_tmp_0
				_hx_tmp_0 = int64(p)
				var _hx_tmp_1 int
				_ = _hx_tmp_1
				if (pos == Hx_Enum_sys_io_fileseek_SeekBegin{}) {
					_hx_tmp_1 = int(((int32)(0)))
				} else {
					var _hx_tmp_2 int
					_ = _hx_tmp_2
					if (pos == Hx_Enum_sys_io_fileseek_SeekCur{}) {
						_hx_tmp_2 = int(((int32)(1)))
					} else {
						_hx_tmp_2 = int(((int32)(2)))
					}

					_hx_tmp_1 = _hx_tmp_2
				}

				hx_result_23.Result, hx_result_23.Error = (*this.Hx_Field___f).Seek(_hx_tmp_0, _hx_tmp_1)
				var this1 struct {
					Error  error
					Result int64
				}
				_ = this1
				this1 = hx_result_23
				if this1.Error != nil {
					var e error
					_ = e
					e = this1.Error
					panic(Hx_Field_haxe_exception_thrown(((any)(e))))
				} else {
					var r int64
					_ = r
					r = this1.Result
				}
			}
		}()
		if hx_try_state == 1 {
			return
		}
	}
}

func (this *Hx_Obj_sys_io_fileoutput) Hx_Field_tell() int32 {
	{
		var hx_try_return int32
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{e}))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var hx_result_24 struct {
					Error  error
					Result int64
				}
				_ = hx_result_24
				var _hx_tmp_0 int64
				_ = _hx_tmp_0
				_hx_tmp_0 = int64(((int32)(0)))
				hx_result_24.Result, hx_result_24.Error = (*this.Hx_Field___f).Seek(_hx_tmp_0, int(((int32)(1))))
				var this1 struct {
					Error  error
					Result int64
				}
				_ = this1
				this1 = hx_result_24
				var this2 int64
				_ = this2
				if this1.Error != nil {
					var e error
					_ = e
					e = this1.Error
					panic(Hx_Field_haxe_exception_thrown(((any)(e))))
				} else {
					var r int64
					_ = r
					r = this1.Result
					this2 = r
				}

				hx_try_state = 1
				hx_try_return = ((int32)(this2))
				return
			}
		}()

		return hx_try_return
	}
}

func init() {
	hxRegisterArrayConstructor[byte]()
}
