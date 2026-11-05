package core

type Hx_Enum_sys_io_fileseek interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_sys_io_fileseek()
}

type Hx_Enum_sys_io_fileseek_SeekEnd struct{}

func (this Hx_Enum_sys_io_fileseek_SeekEnd) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_sys_io_fileseek_SeekEnd) M_Hx_Enum_sys_io_fileseek()              {}
func (this Hx_Enum_sys_io_fileseek_SeekEnd) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_sys_io_fileseek_SeekEnd) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_sys_io_fileseek_RTTI
}
func (this Hx_Enum_sys_io_fileseek_SeekEnd) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_sys_io_fileseek_SeekEnd) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_sys_io_fileseek_SeekCur struct{}

func (this Hx_Enum_sys_io_fileseek_SeekCur) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_sys_io_fileseek_SeekCur) M_Hx_Enum_sys_io_fileseek()              {}
func (this Hx_Enum_sys_io_fileseek_SeekCur) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_sys_io_fileseek_SeekCur) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_sys_io_fileseek_RTTI
}
func (this Hx_Enum_sys_io_fileseek_SeekCur) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_sys_io_fileseek_SeekCur) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_sys_io_fileseek_SeekBegin struct{}

func (this Hx_Enum_sys_io_fileseek_SeekBegin) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_sys_io_fileseek_SeekBegin) M_Hx_Enum_sys_io_fileseek()              {}
func (this Hx_Enum_sys_io_fileseek_SeekBegin) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_sys_io_fileseek_SeekBegin) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_sys_io_fileseek_RTTI
}
func (this Hx_Enum_sys_io_fileseek_SeekBegin) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_sys_io_fileseek_SeekBegin) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_sys_io_fileseek_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"sys.io.FileSeek",
	HxMakeArray[string]("SeekBegin", "SeekCur", "SeekEnd"),
	HxMakeArray[int32](0, 0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_sys_io_fileseek_SeekBegin{}
		case 1:
			return Hx_Enum_sys_io_fileseek_SeekCur{}
		case 2:
			return Hx_Enum_sys_io_fileseek_SeekEnd{}
		default:
			return nil
		}
	},
	nil,
)
