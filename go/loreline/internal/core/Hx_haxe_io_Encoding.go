package core

type Hx_Enum_haxe_io_encoding interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_io_encoding()
}

type Hx_Enum_haxe_io_encoding_UTF8 struct{}

func (this Hx_Enum_haxe_io_encoding_UTF8) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_encoding_UTF8) M_Hx_Enum_haxe_io_encoding()             {}
func (this Hx_Enum_haxe_io_encoding_UTF8) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_io_encoding_UTF8) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_encoding_RTTI
}
func (this Hx_Enum_haxe_io_encoding_UTF8) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_io_encoding_UTF8) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_io_encoding_RawNative struct{}

func (this Hx_Enum_haxe_io_encoding_RawNative) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_encoding_RawNative) M_Hx_Enum_haxe_io_encoding()             {}
func (this Hx_Enum_haxe_io_encoding_RawNative) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_haxe_io_encoding_RawNative) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_encoding_RTTI
}
func (this Hx_Enum_haxe_io_encoding_RawNative) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_io_encoding_RawNative) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_haxe_io_encoding_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.io.Encoding",
	HxMakeArray[string]("UTF8", "RawNative"),
	HxMakeArray[int32](0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_io_encoding_UTF8{}
		case 1:
			return Hx_Enum_haxe_io_encoding_RawNative{}
		default:
			return nil
		}
	},
	nil,
)
