package core

var Hx_Abstract_haxe_ds_readonlyarray_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.ReadOnlyArray",
	HxMakeArray[string]("length", "get_length", "get", "concat"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_() {
}

type Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ struct {
	VTable Hx_Obj_VTable_haxe_ds__readonlyarray_readonlyarray_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateEmptyInstance() *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ {
	obj := &Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateInstance() *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ {
	obj := Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._ReadOnlyArray.ReadOnlyArray_Impl_",
	HxMakeArray[string]("get_length", "get", "concat"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "get_length":
			return Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get_length
		case "get":
			return Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get
		case "concat":
			return Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__concat
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl__RTTI
}

func (this *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_) Hx_As_Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_() *Hx_Obj_haxe_ds__readonlyarray_readonlyarray_impl_ {
	return this
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get_length(this HxArray[any]) int32 {
	return Hx_Field_go_haxe_hxdynamic_getArrayLength(this)
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__get(this HxArray[any], i int32) any {
	return this.Get(i)
}

func Hx_Field_haxe_ds__readonlyarray_readonlyarray_impl__concat(this HxArray[any], a HxArray[any]) HxArray[any] {
	return (Hx_Array_Concat[any](HxMakeArrayView[any](this), a))
}

func init() {
	hxRegisterArrayConstructor[any]()
}
