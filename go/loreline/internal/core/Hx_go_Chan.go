package core

var Hx_Abstract_go_chan_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Chan",
	HxMakeArray[string]("_create", "_hx_new", "send", "receive", "tryReceive"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return nil
	},
	func() any {
		return nil
	},
	nil,
	func(name string) any { return nil },
)

func Hx_ClassInit_Hx_Obj_VTable_go__chan_chan_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_go__chan_chan_impl_() {
}

type Hx_Obj_VTable_go__chan_chan_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go__chan_chan_impl_ struct {
	VTable Hx_Obj_VTable_go__chan_chan_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go__chan_chan_impl__CreateEmptyInstance() *Hx_Obj_go__chan_chan_impl_ {
	obj := &Hx_Obj_go__chan_chan_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go__chan_chan_impl__CreateInstance() *Hx_Obj_go__chan_chan_impl_ {
	obj := Hx_Obj_go__chan_chan_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_go__chan_chan_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go._Chan.Chan_Impl_",
	HxMakeArray[string]("_create", "_hx_new", "send", "receive", "tryReceive"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go__chan_chan_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_go__chan_chan_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go__chan_chan_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_hx_new":
			return Hx_Field_go__chan_chan_impl___hx_new
		}
		return nil
	}
}

func (this *Hx_Obj_go__chan_chan_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go__chan_chan_impl__RTTI
}

func (this *Hx_Obj_go__chan_chan_impl_) Hx_As_Hx_Obj_go__chan_chan_impl_() *Hx_Obj_go__chan_chan_impl_ {
	return this
}

func Hx_Field_go__chan_chan_impl___hx_new(length__isOptional any) chan any {
	length := HxDefault[HxNullable[int32]]()
	_ = length

	if length__isOptional != nil {
		length = HxConvert[HxNullable[int32]](length__isOptional)
	} else {

	}
	{
		if (length).Valid == false {
			var hx_nullable_t0 HxNullable[int32]
			_ = hx_nullable_t0
			hx_nullable_t0 = HxNullable[int32]{}

			length = HxNullable[int32]{Value: (hx_nullable_t0).Value, Valid: (hx_nullable_t0).Valid}
		}

		var length1 HxNullable[int32]
		_ = length1
		length1 = length
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (length1).Valid != false {
			_hx_tmp_0 = (length1).Value
		} else {
			_hx_tmp_0 = 0
		}

		return make(chan any, _hx_tmp_0)
	}
}
