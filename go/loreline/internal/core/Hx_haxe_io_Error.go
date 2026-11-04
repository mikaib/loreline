package core

type Hx_Enum_haxe_io_error interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_io_error()
}

type Hx_Enum_haxe_io_error_Overflow struct{}

func (this Hx_Enum_haxe_io_error_Overflow) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_error_Overflow) M_Hx_Enum_haxe_io_error()                {}
func (this Hx_Enum_haxe_io_error_Overflow) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_haxe_io_error_Overflow) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_error_RTTI
}
func (this Hx_Enum_haxe_io_error_Overflow) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_io_error_Overflow) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_io_error_OutsideBounds struct{}

func (this Hx_Enum_haxe_io_error_OutsideBounds) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_error_OutsideBounds) M_Hx_Enum_haxe_io_error()                {}
func (this Hx_Enum_haxe_io_error_OutsideBounds) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_haxe_io_error_OutsideBounds) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_error_RTTI
}
func (this Hx_Enum_haxe_io_error_OutsideBounds) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_io_error_OutsideBounds) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_io_error_Custom struct {
	Hx_Field_e any
}

func (this Hx_Enum_haxe_io_error_Custom) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_error_Custom) M_Hx_Enum_haxe_io_error()                {}
func (this Hx_Enum_haxe_io_error_Custom) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_haxe_io_error_Custom) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_error_RTTI
}
func (this Hx_Enum_haxe_io_error_Custom) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_haxe_io_error_Custom) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_haxe_io_error_Blocked struct{}

func (this Hx_Enum_haxe_io_error_Blocked) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_io_error_Blocked) M_Hx_Enum_haxe_io_error()                {}
func (this Hx_Enum_haxe_io_error_Blocked) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_io_error_Blocked) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_io_error_RTTI
}
func (this Hx_Enum_haxe_io_error_Blocked) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_io_error_Blocked) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_haxe_io_error_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.io.Error",
	HxMakeArray[string]("Blocked", "Overflow", "OutsideBounds", "Custom"),
	HxMakeArray[int32](0, 0, 0, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_io_error_Blocked{}
		case 1:
			return Hx_Enum_haxe_io_error_Overflow{}
		case 2:
			return Hx_Enum_haxe_io_error_OutsideBounds{}
		case 3:
			return Hx_Enum_haxe_io_error_Custom{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)
