package core

type Hx_Typedef_haxe_ds__vector_vectordata = HxArray[any]

var Hx_Abstract_haxe_ds_vector_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.Vector",
	HxMakeArray[string]("_hx_new", "get", "set", "length", "get_length", "fill", "blit", "toArray", "toData", "fromData", "fromArrayCopy", "copy", "join", "sort", "map"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__vector_vector_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__vector_vector_impl_() {
}

type Hx_Obj_VTable_haxe_ds__vector_vector_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__vector_vector_impl_ struct {
	VTable Hx_Obj_VTable_haxe_ds__vector_vector_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_ds__vector_vector_impl__CreateEmptyInstance() *Hx_Obj_haxe_ds__vector_vector_impl_ {
	obj := &Hx_Obj_haxe_ds__vector_vector_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_ds__vector_vector_impl__CreateInstance() *Hx_Obj_haxe_ds__vector_vector_impl_ {
	obj := Hx_Obj_haxe_ds__vector_vector_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_ds__vector_vector_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._Vector.Vector_Impl_",
	HxMakeArray[string]("_hx_new", "get", "set", "get_length", "fill", "blit", "toArray", "toData", "fromData", "fromArrayCopy", "copy", "join", "sort", "map"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__vector_vector_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds__vector_vector_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__vector_vector_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "get_length":
			return Hx_Field_haxe_ds__vector_vector_impl__get_length
		case "fill":
			return Hx_Field_haxe_ds__vector_vector_impl__fill
		case "join":
			return Hx_Field_haxe_ds__vector_vector_impl__join
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__vector_vector_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__vector_vector_impl__RTTI
}

func (this *Hx_Obj_haxe_ds__vector_vector_impl_) Hx_As_Hx_Obj_haxe_ds__vector_vector_impl_() *Hx_Obj_haxe_ds__vector_vector_impl_ {
	return this
}

func Hx_Field_haxe_ds__vector_vector_impl__get_length(this HxArray[any]) int32 {
	return Hx_Field_go_haxe_hxdynamic_getArrayLength(this)
}

func Hx_Field_haxe_ds__vector_vector_impl__fill(this HxArray[any], value any) {
	var _g int32
	_ = _g
	_g = 0
	var _g1 int32
	_ = _g1
	_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(this)
	for _g < _g1 {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = _g
		_g = (_g + 1)
		var i int32
		_ = i
		i = _hx_tmp_0
		this.Set(i, value)
	}
}

func Hx_Field_haxe_ds__vector_vector_impl__join(this HxArray[any], sep string) string {
	return (Hx_Array_Join[any](HxMakeArrayView[any](this), sep))
}

func init() {
	hxRegisterArrayConstructor[any]()
}
