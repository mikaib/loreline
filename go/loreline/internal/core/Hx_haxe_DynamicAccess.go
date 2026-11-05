package core

var Hx_Abstract_haxe_dynamicaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.DynamicAccess",
	HxMakeArray[string]("_hx_new", "get", "set", "exists", "remove", "keys", "copy", "iterator", "keyValueIterator"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe__dynamicaccess_dynamicaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__dynamicaccess_dynamicaccess_impl_() {
}

type Hx_Obj_VTable_haxe__dynamicaccess_dynamicaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe__dynamicaccess_dynamicaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_ {
	obj := &Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__CreateInstance() *Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_ {
	obj := Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._DynamicAccess.DynamicAccess_Impl_",
	HxMakeArray[string]("_hx_new", "get", "set", "exists", "remove", "keys", "copy", "iterator", "keyValueIterator"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_hx_new":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl___hx_new
		case "get":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__get
		case "set":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__set
		case "exists":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__exists
		case "remove":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__remove
		case "keys":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__keys
		case "copy":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__copy
		case "iterator":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__iterator
		case "keyValueIterator":
			return Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__keyValueIterator
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl__RTTI
}

func (this *Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_) Hx_As_Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_() *Hx_Obj_haxe__dynamicaccess_dynamicaccess_impl_ {
	return this
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl___hx_new() any {
	return any(map[string]any{})
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__get(this any, key string) any {
	return Hx_Field_reflect_field(this, key)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__set(this any, key string, value any) any {
	Hx_Field_reflect_setField(this, key, value)
	return value
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__exists(this any, key string) bool {
	return Hx_Field_reflect_hasField(this, key)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__remove(this any, key string) bool {
	return Hx_Field_reflect_deleteField(this, key)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__keys(this any) HxArray[string] {
	return Hx_Field_reflect_fields(this)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__copy(this any) any {
	return Hx_Field_reflect_copy(this)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__iterator(this any) *Hx_Obj_haxe_iterators_dynamicaccessiterator {
	return Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateInstance(this)
}

func Hx_Field_haxe__dynamicaccess_dynamicaccess_impl__keyValueIterator(this any) *Hx_Obj_haxe_iterators_dynamicaccesskeyvalueiterator {
	return Hx_Obj_haxe_iterators_dynamicaccesskeyvalueiterator_CreateInstance(this)
}

func init() {
	hxRegisterArrayConstructor[string]()
}
