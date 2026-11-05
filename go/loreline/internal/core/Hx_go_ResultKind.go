package core

type Hx_Enum_go_resultkind interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_go_resultkind()
}

type Hx_Enum_go_resultkind_Ok struct {
	Hx_Field_r any
}

func (this Hx_Enum_go_resultkind_Ok) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_go_resultkind_Ok) M_Hx_Enum_go_resultkind()                {}
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_go_resultkind_RTTI
}
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_r))
}
func (this Hx_Enum_go_resultkind_Ok) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_r)
	default:
		return nil
	}
}

type Hx_Enum_go_resultkind_Err struct {
	Hx_Field_e any
}

func (this Hx_Enum_go_resultkind_Err) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_go_resultkind_Err) M_Hx_Enum_go_resultkind()                {}
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_go_resultkind_RTTI
}
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_go_resultkind_Err) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

var Hx_Enum_go_resultkind_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"go.ResultKind",
	HxMakeArray[string]("Ok", "Err"),
	HxMakeArray[int32](1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_go_resultkind_Ok{params.Get_Dyn(0).(any)}
		case 1:
			return Hx_Enum_go_resultkind_Err{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)
