package core

type Hx_Typedef_haxe__rest_nativerest = HxArray[any]

var Hx_Abstract_haxe_rest_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Rest",
	HxMakeArray[string]("length", "get_length", "of", "_hx_new", "get", "toArray", "iterator", "keyValueIterator", "append", "prepend", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe__rest_rest_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__rest_rest_impl_() {
}

type Hx_Obj_VTable_haxe__rest_rest_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__rest_rest_impl_ struct {
	VTable Hx_Obj_VTable_haxe__rest_rest_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__rest_rest_impl__CreateEmptyInstance() *Hx_Obj_haxe__rest_rest_impl_ {
	obj := &Hx_Obj_haxe__rest_rest_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__rest_rest_impl__CreateInstance() *Hx_Obj_haxe__rest_rest_impl_ {
	obj := Hx_Obj_haxe__rest_rest_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__rest_rest_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._Rest.Rest_Impl_",
	HxMakeArray[string]("get_length", "of", "_hx_new", "get", "toArray", "iterator", "keyValueIterator", "append", "prepend", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__rest_rest_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__rest_rest_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__rest_rest_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "get_length":
			return Hx_Field_haxe__rest_rest_impl__get_length
		case "of":
			return Hx_Field_haxe__rest_rest_impl__of
		case "_hx_new":
			return Hx_Field_haxe__rest_rest_impl___hx_new
		case "get":
			return Hx_Field_haxe__rest_rest_impl__get
		case "toArray":
			return Hx_Field_haxe__rest_rest_impl__toArray
		case "iterator":
			return Hx_Field_haxe__rest_rest_impl__iterator
		case "keyValueIterator":
			return Hx_Field_haxe__rest_rest_impl__keyValueIterator
		case "append":
			return Hx_Field_haxe__rest_rest_impl__append
		case "prepend":
			return Hx_Field_haxe__rest_rest_impl__prepend
		case "toString":
			return Hx_Field_haxe__rest_rest_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__rest_rest_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__rest_rest_impl__RTTI
}

func (this *Hx_Obj_haxe__rest_rest_impl_) Hx_As_Hx_Obj_haxe__rest_rest_impl_() *Hx_Obj_haxe__rest_rest_impl_ {
	return this
}

func Hx_Field_haxe__rest_rest_impl__get_length(this HxArray[any]) int32 {
	return Hx_Field_go_haxe_hxdynamic_getArrayLength(this)
}

func Hx_Field_haxe__rest_rest_impl__of(array HxArray[any]) HxArray[any] {
	return array
}

func Hx_Field_haxe__rest_rest_impl___hx_new(array HxArray[any]) HxArray[any] {
	return array
}

func Hx_Field_haxe__rest_rest_impl__get(this HxArray[any], index int32) any {
	return this.Get(index)
}

func Hx_Field_haxe__rest_rest_impl__toArray(this HxArray[any]) HxArray[any] {
	var slice []any
	_ = slice
	slice = Hx_Field_go_haxe_hxdynamic_toAnySlice(this)
	var self []any
	_ = self
	self = slice
	return HxMakeArray(self...)
}

func Hx_Field_haxe__rest_rest_impl__iterator(this HxArray[any]) *Hx_Obj_haxe_iterators_restiterator {
	return Hx_Obj_haxe_iterators_restiterator_CreateInstance(this)
}

func Hx_Field_haxe__rest_rest_impl__keyValueIterator(this HxArray[any]) *Hx_Obj_haxe_iterators_restkeyvalueiterator {
	return Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateInstance(this)
}

func Hx_Field_haxe__rest_rest_impl__append(this HxArray[any], item any) HxArray[any] {
	var slice []any
	_ = slice
	slice = Hx_Field_go_haxe_hxdynamic_toAnySlice(this)
	var self []any
	_ = self
	self = slice
	var result HxArray[any]
	_ = result
	result = HxMakeArray(self...)
	_ = (Hx_Array_Push[any](HxMakeArrayView[any](result), item))
	return result
}

func Hx_Field_haxe__rest_rest_impl__prepend(this HxArray[any], item any) HxArray[any] {
	var slice []any
	_ = slice
	slice = Hx_Field_go_haxe_hxdynamic_toAnySlice(this)
	var self []any
	_ = self
	self = slice
	var result HxArray[any]
	_ = result
	result = HxMakeArray(self...)
	(Hx_Array_Unshift[any](HxMakeArrayView[any](result), item))
	return result
}

func Hx_Field_haxe__rest_rest_impl__toString(this HxArray[any]) string {
	return Hx_Field_std_string(this)
}

func init() {
	hxRegisterArrayConstructor[any]()
}
