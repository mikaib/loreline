package core

import "os"

func Hx_ClassInit_Hx_Obj_VTable_sys_io_fileinput() {
}

func Hx_StaticInit_Hx_Obj_VTable_sys_io_fileinput() {
}

type Hx_Obj_VTable_sys_io_fileinput interface {
	Hx_Field_seek(p int32, pos Hx_Enum_sys_io_fileseek)
	Hx_Field_tell() int32
	Hx_Field_readAll(bufsize__isOptional any) *Hx_Obj_haxe_io_bytes
	Hx_Field_readFullBytes(s *Hx_Obj_haxe_io_bytes, pos int32, _hx_reserved_len int32)
	Hx_Field_read(nbytes int32) *Hx_Obj_haxe_io_bytes
	Hx_Field_readUntil(end int32) string
	Hx_Field_readUInt16() int32
	Hx_Field_readByte() int32
	Hx_Field_eof() bool
	Hx_Field_set_bigEndian(b bool) bool
	Hx_Field_readLine() string
	Hx_Field_readDouble() float64
	Hx_Field_readInt16() int32
	Hx_Field_readUInt24() int32
	Hx_Field_readInt32() int32
	Hx_Field_readInt8() int32
	Hx_Field_readInt24() int32
	Hx_Field_readString(_hx_reserved_len int32, encoding__isOptional any) string
	Hx_Field_close()
	Hx_Field_readFloat() float64
	Hx_Field_readBytes(s *Hx_Obj_haxe_io_bytes, p int32, l int32) int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_sys_io_fileinput struct {
	Hx_Obj_haxe_io_input
	VTable             Hx_Obj_VTable_sys_io_fileinput `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field___f       *os.File
	Hx_Field___eof     bool
	Hx_Field___byteBuf []byte
}

func Hx_Obj_sys_io_fileinput_CreateEmptyInstance() *Hx_Obj_sys_io_fileinput {
	obj := &Hx_Obj_sys_io_fileinput{}
	obj.Hx_Obj_haxe_io_input = *Hx_Obj_haxe_io_input_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_io_input.VTable = obj
	obj.Hx_Field___f = HxDefault[*os.File]()
	obj.Hx_Field___eof = HxDefault[bool]()
	obj.Hx_Field___byteBuf = HxDefault[[]byte]()
	return obj
}

func Hx_Obj_sys_io_fileinput_CreateInstance(f *os.File) *Hx_Obj_sys_io_fileinput {
	obj := Hx_Obj_sys_io_fileinput_CreateEmptyInstance()
	obj.Hx_New(f)
	return obj
}

func (this *Hx_Obj_sys_io_fileinput) Hx_New(f *os.File) {
	{
		this.Hx_Field___f = f
		this.Hx_Field___eof = false
		var self HxArray[byte]
		_ = self
		self = HxMakeArray[byte](byte(((int32)(0))))
		this.Hx_Field___byteBuf = self.Underlying()
	}

}

var Hx_Obj_sys_io_fileinput_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"sys.io.FileInput",
	HxMakeArray[string](),
	HxMakeArray[string]("__f", "__eof", "__byteBuf", "readByte", "readBytes", "close", "seek", "tell", "eof", "bigEndian", "set_bigEndian", "readAll", "readFullBytes", "read", "readUntil", "readLine", "readFloat", "readDouble", "readInt8", "readInt16", "readUInt16", "readInt24", "readUInt24", "readInt32", "readString"),
	Hx_Obj_haxe_io_input_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_sys_io_fileinput_CreateInstance(HxConvert[*os.File](params.Get(0)))
	},
	func() any {
		return Hx_Obj_sys_io_fileinput_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_sys_io_fileinput_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_sys_io_fileinput_RTTI
}

func (this *Hx_Obj_sys_io_fileinput) Hx_As_Hx_Obj_sys_io_fileinput() *Hx_Obj_sys_io_fileinput {
	return this
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_readByte() int32 {
	var hx_result_16 struct {
		Error  error
		Result int
	}
	_ = hx_result_16
	hx_result_16.Result, hx_result_16.Error = (*this.Hx_Field___f).Read(this.Hx_Field___byteBuf)
	var t struct {
		Error  error
		Result int
	}
	_ = t
	t = ((struct {
		Error  error
		Result int
	})(((struct {
		Error  error
		Result int
	})(hx_result_16))))
	if t.Error != nil {
		if t.Error.Error() == "EOF" {
			this.Hx_Field___eof = true
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_io_eof_CreateInstance()))))
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{t.Error}))))
	}

	return ((int32)(this.Hx_Field___byteBuf[0]))
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_readBytes(s *Hx_Obj_haxe_io_bytes, p int32, l int32) int32 {
	var self HxArray[byte]
	_ = self
	self = HxAnyToTypedArray[byte](s.Hx_Field_b)
	var hx_result_18 struct {
		Error  error
		Result int
	}
	_ = hx_result_18
	hx_result_18.Result, hx_result_18.Error = (*this.Hx_Field___f).Read(self.Underlying())
	var t struct {
		Error  error
		Result int
	}
	_ = t
	t = ((struct {
		Error  error
		Result int
	})(((struct {
		Error  error
		Result int
	})(hx_result_18))))
	if t.Error != nil {
		if t.Error.Error() == "EOF" {
			this.Hx_Field___eof = true
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_io_eof_CreateInstance()))))
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Enum_haxe_io_error_Custom{t.Error}))))
	}

	return ((int32)(t.Result))
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_close() {
	var this1 error
	_ = this1
	this1 = (*this.Hx_Field___f).Close()
	if this1 != nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(this1.Error()))))
	}
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_seek(p int32, pos Hx_Enum_sys_io_fileseek) {
	{
		var hx_result_19 struct {
			Error  error
			Result int64
		}
		_ = hx_result_19
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

		hx_result_19.Result, hx_result_19.Error = (*this.Hx_Field___f).Seek(_hx_tmp_0, _hx_tmp_1)
		var this1 struct {
			Error  error
			Result int64
		}
		_ = this1
		this1 = hx_result_19
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

	this.Hx_Field___eof = false
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_tell() int32 {
	var hx_result_20 struct {
		Error  error
		Result int64
	}
	_ = hx_result_20
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = int64(((int32)(0)))
	hx_result_20.Result, hx_result_20.Error = (*this.Hx_Field___f).Seek(_hx_tmp_0, int(((int32)(1))))
	var this1 struct {
		Error  error
		Result int64
	}
	_ = this1
	this1 = hx_result_20
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

	return ((int32)(this2))
}

func (this *Hx_Obj_sys_io_fileinput) Hx_Field_eof() bool {
	return this.Hx_Field___eof
}

func init() {
	hxRegisterArrayConstructor[byte]()
}
