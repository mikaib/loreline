package core

var Hx_Abstract_go_map_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Map",
	HxMakeArray[string]("_create", "_hx_new", "set", "get", "exists", "remove", "size", "keys", "values", "copy", "iterator", "clear"),
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

func Hx_ClassInit_Hx_Obj_VTable_go__map_map_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_go__map_map_impl_() {
}

type Hx_Obj_VTable_go__map_map_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go__map_map_impl_ struct {
	VTable Hx_Obj_VTable_go__map_map_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go__map_map_impl__CreateEmptyInstance() *Hx_Obj_go__map_map_impl_ {
	obj := &Hx_Obj_go__map_map_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go__map_map_impl__CreateInstance() *Hx_Obj_go__map_map_impl_ {
	obj := Hx_Obj_go__map_map_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_go__map_map_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go._Map.Map_Impl_",
	HxMakeArray[string]("_create", "_hx_new", "set", "get", "exists", "remove", "size", "keys", "values", "copy", "iterator", "clear"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go__map_map_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_go__map_map_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go__map_map_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_hx_new":
			return Hx_Field_go__map_map_impl___hx_new
		}
		return nil
	}
}

func (this *Hx_Obj_go__map_map_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go__map_map_impl__RTTI
}

func (this *Hx_Obj_go__map_map_impl_) Hx_As_Hx_Obj_go__map_map_impl_() *Hx_Obj_go__map_map_impl_ {
	return this
}

func Hx_Field_go__map_map_impl___hx_new() map[any]any {
	return map[any]any{}
}
