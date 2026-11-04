package core

var Hx_Abstract_haxe_ds_map_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.Map",
	HxMakeArray[string]("_hx_new", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "clear", "size", "arrayWrite", "toStringMap", "toIntMap", "toInt64Map", "toEnumValueMapMap", "toObjectMap", "fromStringMap", "fromIntMap", "fromInt64Map", "fromObjectMap"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__map_map_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__map_map_impl_() {
}

type Hx_Obj_VTable_haxe_ds__map_map_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__map_map_impl_ struct {
	VTable Hx_Obj_VTable_haxe_ds__map_map_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_ds__map_map_impl__CreateEmptyInstance() *Hx_Obj_haxe_ds__map_map_impl_ {
	obj := &Hx_Obj_haxe_ds__map_map_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_ds__map_map_impl__CreateInstance() *Hx_Obj_haxe_ds__map_map_impl_ {
	obj := Hx_Obj_haxe_ds__map_map_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_ds__map_map_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._Map.Map_Impl_",
	HxMakeArray[string]("_hx_new", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "clear", "size", "arrayWrite", "toStringMap", "toIntMap", "toInt64Map", "toEnumValueMapMap", "toObjectMap", "fromStringMap", "fromIntMap", "fromInt64Map", "fromObjectMap"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__map_map_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds__map_map_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__map_map_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "set":
			return Hx_Field_haxe_ds__map_map_impl__set
		case "get":
			return Hx_Field_haxe_ds__map_map_impl__get
		case "exists":
			return Hx_Field_haxe_ds__map_map_impl__exists
		case "remove":
			return Hx_Field_haxe_ds__map_map_impl__remove
		case "keys":
			return Hx_Field_haxe_ds__map_map_impl__keys
		case "iterator":
			return Hx_Field_haxe_ds__map_map_impl__iterator
		case "keyValueIterator":
			return Hx_Field_haxe_ds__map_map_impl__keyValueIterator
		case "copy":
			return Hx_Field_haxe_ds__map_map_impl__copy
		case "toString":
			return Hx_Field_haxe_ds__map_map_impl__toString
		case "clear":
			return Hx_Field_haxe_ds__map_map_impl__clear
		case "size":
			return Hx_Field_haxe_ds__map_map_impl__size
		case "arrayWrite":
			return Hx_Field_haxe_ds__map_map_impl__arrayWrite
		case "toStringMap":
			return Hx_Field_haxe_ds__map_map_impl__toStringMap
		case "toIntMap":
			return Hx_Field_haxe_ds__map_map_impl__toIntMap
		case "toInt64Map":
			return Hx_Field_haxe_ds__map_map_impl__toInt64Map
		case "toEnumValueMapMap":
			return Hx_Field_haxe_ds__map_map_impl__toEnumValueMapMap
		case "toObjectMap":
			return Hx_Field_haxe_ds__map_map_impl__toObjectMap
		case "fromStringMap":
			return Hx_Field_haxe_ds__map_map_impl__fromStringMap
		case "fromIntMap":
			return Hx_Field_haxe_ds__map_map_impl__fromIntMap
		case "fromInt64Map":
			return Hx_Field_haxe_ds__map_map_impl__fromInt64Map
		case "fromObjectMap":
			return Hx_Field_haxe_ds__map_map_impl__fromObjectMap
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__map_map_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__map_map_impl__RTTI
}

func (this *Hx_Obj_haxe_ds__map_map_impl_) Hx_As_Hx_Obj_haxe_ds__map_map_impl_() *Hx_Obj_haxe_ds__map_map_impl_ {
	return this
}

func Hx_Field_haxe_ds__map_map_impl__set(this *Hx_Obj_haxe_imap, key any, value any) {
	this.VTable.Hx_Field_set(key, value)
}

func Hx_Field_haxe_ds__map_map_impl__get(this *Hx_Obj_haxe_imap, key any) any {
	return this.VTable.Hx_Field_get(key)
}

func Hx_Field_haxe_ds__map_map_impl__exists(this *Hx_Obj_haxe_imap, key any) bool {
	return this.VTable.Hx_Field_exists(key)
}

func Hx_Field_haxe_ds__map_map_impl__remove(this *Hx_Obj_haxe_imap, key any) bool {
	return this.VTable.Hx_Field_remove(key)
}

func Hx_Field_haxe_ds__map_map_impl__keys(this *Hx_Obj_haxe_imap) any {
	return this.VTable.Hx_Field_keys()
}

func Hx_Field_haxe_ds__map_map_impl__iterator(this *Hx_Obj_haxe_imap) any {
	return this.VTable.Hx_Field_iterator()
}

func Hx_Field_haxe_ds__map_map_impl__keyValueIterator(this *Hx_Obj_haxe_imap) any {
	return this.VTable.Hx_Field_keyValueIterator()
}

func Hx_Field_haxe_ds__map_map_impl__copy(this *Hx_Obj_haxe_imap) *Hx_Obj_haxe_imap {
	return this.VTable.Hx_Field_copy()
}

func Hx_Field_haxe_ds__map_map_impl__toString(this *Hx_Obj_haxe_imap) string {
	return this.VTable.Hx_Field_toString()
}

func Hx_Field_haxe_ds__map_map_impl__clear(this *Hx_Obj_haxe_imap) {
	this.VTable.Hx_Field_clear()
}

func Hx_Field_haxe_ds__map_map_impl__size(this *Hx_Obj_haxe_imap) int32 {
	return this.VTable.Hx_Field_size()
}

func Hx_Field_haxe_ds__map_map_impl__arrayWrite(this *Hx_Obj_haxe_imap, k any, v any) any {
	this.VTable.Hx_Field_set(k, v)
	return v
}

func Hx_Field_haxe_ds__map_map_impl__toStringMap(t *Hx_Obj_haxe_imap) *Hx_Obj_haxe_ds_stringmap {
	return Hx_Obj_haxe_ds_stringmap_CreateInstance()
}

func Hx_Field_haxe_ds__map_map_impl__toIntMap(t *Hx_Obj_haxe_imap) *Hx_Obj_haxe_ds_intmap {
	return Hx_Obj_haxe_ds_intmap_CreateInstance()
}

func Hx_Field_haxe_ds__map_map_impl__toInt64Map(t *Hx_Obj_haxe_imap) *Hx_Obj_haxe_ds_int64map {
	return Hx_Obj_haxe_ds_int64map_CreateInstance()
}

func Hx_Field_haxe_ds__map_map_impl__toEnumValueMapMap(t *Hx_Obj_haxe_imap) *Hx_Obj_haxe_ds_enumvaluemap {
	return Hx_Obj_haxe_ds_enumvaluemap_CreateInstance()
}

func Hx_Field_haxe_ds__map_map_impl__toObjectMap(t *Hx_Obj_haxe_imap) *Hx_Obj_haxe_ds_objectmap {
	return Hx_Obj_haxe_ds_objectmap_CreateInstance()
}

func Hx_Field_haxe_ds__map_map_impl__fromStringMap(_hx_reserved_map *Hx_Obj_haxe_ds_stringmap) *Hx_Obj_haxe_imap {
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_23 := _hx_reserved_map
		if _hx_icast_23 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_23, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_23.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func Hx_Field_haxe_ds__map_map_impl__fromIntMap(_hx_reserved_map *Hx_Obj_haxe_ds_intmap) *Hx_Obj_haxe_imap {
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_24 := _hx_reserved_map
		if _hx_icast_24 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_24, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_24.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func Hx_Field_haxe_ds__map_map_impl__fromInt64Map(_hx_reserved_map *Hx_Obj_haxe_ds_int64map) *Hx_Obj_haxe_imap {
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_25 := _hx_reserved_map
		if _hx_icast_25 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_25, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_25.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func Hx_Field_haxe_ds__map_map_impl__fromObjectMap(_hx_reserved_map *Hx_Obj_haxe_ds_objectmap) *Hx_Obj_haxe_imap {
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_26 := _hx_reserved_map
		if _hx_icast_26 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_26, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_26.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}
