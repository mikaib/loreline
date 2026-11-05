package core

type Hx_Typedef_haxe_rtti_path = string

type Hx_Typedef_haxe_rtti_platforms = HxArray[string]

type Hx_Typedef_haxe_rtti_functionargument = any

type Hx_Enum_haxe_rtti_ctype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_rtti_ctype()
}

type Hx_Enum_haxe_rtti_ctype_CUnknown struct{}

func (this Hx_Enum_haxe_rtti_ctype_CUnknown) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CUnknown) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CUnknown) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_rtti_ctype_CUnknown) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CUnknown) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_ctype_CUnknown) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CTypedef struct {
	Hx_Field_name   Hx_Typedef_haxe_rtti_path
	Hx_Field_params HxArray[Hx_Enum_haxe_rtti_ctype]
}

func (this Hx_Enum_haxe_rtti_ctype_CTypedef) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CTypedef) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CTypedef) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_haxe_rtti_ctype_CTypedef) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CTypedef) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_params))
}
func (this Hx_Enum_haxe_rtti_ctype_CTypedef) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CFunction struct {
	Hx_Field_args HxArray[any]
	Hx_Field_ret  Hx_Enum_haxe_rtti_ctype
}

func (this Hx_Enum_haxe_rtti_ctype_CFunction) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CFunction) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CFunction) Hx_Field_enumIndex() int32               { return 4 }
func (this Hx_Enum_haxe_rtti_ctype_CFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CFunction) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_args), any(this.Hx_Field_ret))
}
func (this Hx_Enum_haxe_rtti_ctype_CFunction) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_args)
	case 1:
		return any(this.Hx_Field_ret)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CEnum struct {
	Hx_Field_name   Hx_Typedef_haxe_rtti_path
	Hx_Field_params HxArray[Hx_Enum_haxe_rtti_ctype]
}

func (this Hx_Enum_haxe_rtti_ctype_CEnum) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CEnum) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CEnum) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_haxe_rtti_ctype_CEnum) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CEnum) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_params))
}
func (this Hx_Enum_haxe_rtti_ctype_CEnum) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CDynamic struct {
	Hx_Field_t Hx_Enum_haxe_rtti_ctype
}

func (this Hx_Enum_haxe_rtti_ctype_CDynamic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CDynamic) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CDynamic) Hx_Field_enumIndex() int32               { return 6 }
func (this Hx_Enum_haxe_rtti_ctype_CDynamic) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CDynamic) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_haxe_rtti_ctype_CDynamic) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CClass struct {
	Hx_Field_name   Hx_Typedef_haxe_rtti_path
	Hx_Field_params HxArray[Hx_Enum_haxe_rtti_ctype]
}

func (this Hx_Enum_haxe_rtti_ctype_CClass) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CClass) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CClass) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_haxe_rtti_ctype_CClass) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CClass) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_params))
}
func (this Hx_Enum_haxe_rtti_ctype_CClass) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CAnonymous struct {
	Hx_Field_fields HxArray[any]
}

func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) Hx_Field_enumIndex() int32               { return 5 }
func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_fields))
}
func (this Hx_Enum_haxe_rtti_ctype_CAnonymous) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_fields)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_ctype_CAbstract struct {
	Hx_Field_name   Hx_Typedef_haxe_rtti_path
	Hx_Field_params HxArray[Hx_Enum_haxe_rtti_ctype]
}

func (this Hx_Enum_haxe_rtti_ctype_CAbstract) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_ctype_CAbstract) M_Hx_Enum_haxe_rtti_ctype()              {}
func (this Hx_Enum_haxe_rtti_ctype_CAbstract) Hx_Field_enumIndex() int32               { return 7 }
func (this Hx_Enum_haxe_rtti_ctype_CAbstract) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_ctype_RTTI
}
func (this Hx_Enum_haxe_rtti_ctype_CAbstract) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_params))
}
func (this Hx_Enum_haxe_rtti_ctype_CAbstract) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_params)
	default:
		return nil
	}
}

var Hx_Enum_haxe_rtti_ctype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.rtti.CType",
	HxMakeArray[string]("CUnknown", "CEnum", "CClass", "CTypedef", "CFunction", "CAnonymous", "CDynamic", "CAbstract"),
	HxMakeArray[int32](0, 2, 2, 2, 2, 1, 1, 2),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_rtti_ctype_CUnknown{}
		case 1:
			return Hx_Enum_haxe_rtti_ctype_CEnum{params.Get_Dyn(0).(Hx_Typedef_haxe_rtti_path), HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](params.Get_Dyn(1))}
		case 2:
			return Hx_Enum_haxe_rtti_ctype_CClass{params.Get_Dyn(0).(Hx_Typedef_haxe_rtti_path), HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](params.Get_Dyn(1))}
		case 3:
			return Hx_Enum_haxe_rtti_ctype_CTypedef{params.Get_Dyn(0).(Hx_Typedef_haxe_rtti_path), HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](params.Get_Dyn(1))}
		case 4:
			return Hx_Enum_haxe_rtti_ctype_CFunction{HxAnyToTypedArray[any](params.Get_Dyn(0)), params.Get_Dyn(1).(Hx_Enum_haxe_rtti_ctype)}
		case 5:
			return Hx_Enum_haxe_rtti_ctype_CAnonymous{HxAnyToTypedArray[any](params.Get_Dyn(0))}
		case 6:
			return Hx_Enum_haxe_rtti_ctype_CDynamic{params.Get_Dyn(0).(Hx_Enum_haxe_rtti_ctype)}
		case 7:
			return Hx_Enum_haxe_rtti_ctype_CAbstract{params.Get_Dyn(0).(Hx_Typedef_haxe_rtti_path), HxAnyToTypedArray[Hx_Enum_haxe_rtti_ctype](params.Get_Dyn(1))}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_haxe_rtti_pathparams = any

type Hx_Typedef_haxe_rtti_typeparams = HxArray[string]

type Hx_Enum_haxe_rtti_rights interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_rtti_rights()
}

type Hx_Enum_haxe_rtti_rights_RNormal struct{}

func (this Hx_Enum_haxe_rtti_rights_RNormal) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RNormal) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RNormal) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_rtti_rights_RNormal) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RNormal) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_rights_RNormal) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_rights_RNo struct{}

func (this Hx_Enum_haxe_rtti_rights_RNo) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RNo) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RNo) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_haxe_rtti_rights_RNo) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RNo) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_rights_RNo) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_rights_RMethod struct{}

func (this Hx_Enum_haxe_rtti_rights_RMethod) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RMethod) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RMethod) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_haxe_rtti_rights_RMethod) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RMethod) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_rights_RMethod) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_rights_RInline struct{}

func (this Hx_Enum_haxe_rtti_rights_RInline) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RInline) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RInline) Hx_Field_enumIndex() int32               { return 5 }
func (this Hx_Enum_haxe_rtti_rights_RInline) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RInline) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_rights_RInline) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_rights_RDynamic struct{}

func (this Hx_Enum_haxe_rtti_rights_RDynamic) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RDynamic) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RDynamic) Hx_Field_enumIndex() int32               { return 4 }
func (this Hx_Enum_haxe_rtti_rights_RDynamic) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RDynamic) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_rtti_rights_RDynamic) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_rights_RCall struct {
	Hx_Field_m string
}

func (this Hx_Enum_haxe_rtti_rights_RCall) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_rights_RCall) M_Hx_Enum_haxe_rtti_rights()             {}
func (this Hx_Enum_haxe_rtti_rights_RCall) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_haxe_rtti_rights_RCall) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_rights_RTTI
}
func (this Hx_Enum_haxe_rtti_rights_RCall) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_m))
}
func (this Hx_Enum_haxe_rtti_rights_RCall) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_m)
	default:
		return nil
	}
}

var Hx_Enum_haxe_rtti_rights_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.rtti.Rights",
	HxMakeArray[string]("RNormal", "RNo", "RCall", "RMethod", "RDynamic", "RInline"),
	HxMakeArray[int32](0, 0, 1, 0, 0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_rtti_rights_RNormal{}
		case 1:
			return Hx_Enum_haxe_rtti_rights_RNo{}
		case 2:
			return Hx_Enum_haxe_rtti_rights_RCall{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 3:
			return Hx_Enum_haxe_rtti_rights_RMethod{}
		case 4:
			return Hx_Enum_haxe_rtti_rights_RDynamic{}
		case 5:
			return Hx_Enum_haxe_rtti_rights_RInline{}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_haxe_rtti_metadata = HxArray[any]

type Hx_Typedef_haxe_rtti_classfield = any

type Hx_Typedef_haxe_rtti_typeinfos = any

type Hx_Typedef_haxe_rtti_classdef = any

type Hx_Typedef_haxe_rtti_enumfield = any

type Hx_Typedef_haxe_rtti_enumdef = any

type Hx_Typedef_haxe_rtti_typedef = any

type Hx_Typedef_haxe_rtti_abstractdef = any

type Hx_Enum_haxe_rtti_typetree interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_rtti_typetree()
}

type Hx_Enum_haxe_rtti_typetree_TTypedecl struct {
	Hx_Field_t any
}

func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) M_Hx_Enum_haxe_rtti_typetree()           {}
func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_typetree_RTTI
}
func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_t))
}
func (this Hx_Enum_haxe_rtti_typetree_TTypedecl) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_t)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_typetree_TPackage struct {
	Hx_Field_name string
	Hx_Field_full string
	Hx_Field_subs HxArray[Hx_Enum_haxe_rtti_typetree]
}

func (this Hx_Enum_haxe_rtti_typetree_TPackage) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_typetree_TPackage) M_Hx_Enum_haxe_rtti_typetree()           {}
func (this Hx_Enum_haxe_rtti_typetree_TPackage) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_rtti_typetree_TPackage) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_typetree_RTTI
}
func (this Hx_Enum_haxe_rtti_typetree_TPackage) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_full), any(this.Hx_Field_subs))
}
func (this Hx_Enum_haxe_rtti_typetree_TPackage) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_full)
	case 2:
		return any(this.Hx_Field_subs)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_typetree_TEnumdecl struct {
	Hx_Field_e any
}

func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) M_Hx_Enum_haxe_rtti_typetree()           {}
func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_typetree_RTTI
}
func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_e))
}
func (this Hx_Enum_haxe_rtti_typetree_TEnumdecl) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_e)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_typetree_TClassdecl struct {
	Hx_Field_c any
}

func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) M_Hx_Enum_haxe_rtti_typetree() {}
func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) Hx_Field_enumIndex() int32     { return 1 }
func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_typetree_RTTI
}
func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_haxe_rtti_typetree_TClassdecl) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_haxe_rtti_typetree_TAbstractdecl struct {
	Hx_Field_a any
}

func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) M_Hx_Enum_haxe_rtti_typetree() {}
func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) Hx_Field_enumIndex() int32     { return 4 }
func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_rtti_typetree_RTTI
}
func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_a))
}
func (this Hx_Enum_haxe_rtti_typetree_TAbstractdecl) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_a)
	default:
		return nil
	}
}

var Hx_Enum_haxe_rtti_typetree_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.rtti.TypeTree",
	HxMakeArray[string]("TPackage", "TClassdecl", "TEnumdecl", "TTypedecl", "TAbstractdecl"),
	HxMakeArray[int32](3, 1, 1, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_rtti_typetree_TPackage{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1)))), HxAnyToTypedArray[Hx_Enum_haxe_rtti_typetree](params.Get_Dyn(2))}
		case 1:
			return Hx_Enum_haxe_rtti_typetree_TClassdecl{params.Get_Dyn(0).(any)}
		case 2:
			return Hx_Enum_haxe_rtti_typetree_TEnumdecl{params.Get_Dyn(0).(any)}
		case 3:
			return Hx_Enum_haxe_rtti_typetree_TTypedecl{params.Get_Dyn(0).(any)}
		case 4:
			return Hx_Enum_haxe_rtti_typetree_TAbstractdecl{params.Get_Dyn(0).(any)}
		default:
			return nil
		}
	},
	nil,
)

type Hx_Typedef_haxe_rtti_typeroot = HxArray[Hx_Enum_haxe_rtti_typetree]

func Hx_ClassInit_Hx_Obj_VTable_haxe_rtti_typeapi() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_rtti_typeapi() {
}

type Hx_Obj_VTable_haxe_rtti_typeapi interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_rtti_typeapi struct {
	VTable Hx_Obj_VTable_haxe_rtti_typeapi `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_rtti_typeapi_CreateEmptyInstance() *Hx_Obj_haxe_rtti_typeapi {
	obj := &Hx_Obj_haxe_rtti_typeapi{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_rtti_typeapi_CreateInstance() *Hx_Obj_haxe_rtti_typeapi {
	obj := Hx_Obj_haxe_rtti_typeapi_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_rtti_typeapi_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.rtti.TypeApi",
	HxMakeArray[string]("typeInfos", "isVar", "leq", "rightsEq", "typeEq", "fieldEq", "constructorEq"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_rtti_typeapi_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_rtti_typeapi_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_rtti_typeapi_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "typeInfos":
			return Hx_Field_haxe_rtti_typeapi_typeInfos
		case "isVar":
			return Hx_Field_haxe_rtti_typeapi_isVar
		case "leq":
			return Hx_Field_haxe_rtti_typeapi_leq
		case "rightsEq":
			return Hx_Field_haxe_rtti_typeapi_rightsEq
		case "typeEq":
			return Hx_Field_haxe_rtti_typeapi_typeEq
		case "fieldEq":
			return Hx_Field_haxe_rtti_typeapi_fieldEq
		case "constructorEq":
			return Hx_Field_haxe_rtti_typeapi_constructorEq
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_rtti_typeapi) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_rtti_typeapi_RTTI
}

func (this *Hx_Obj_haxe_rtti_typeapi) Hx_As_Hx_Obj_haxe_rtti_typeapi() *Hx_Obj_haxe_rtti_typeapi {
	return this
}

func Hx_Field_haxe_rtti_typeapi_typeInfos(t Hx_Enum_haxe_rtti_typetree) any {
	var inf any
	_ = inf
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)("Unexpected Package"))))
		}

	case 1:
		{
			var c any
			_ = c
			c = t.(Hx_Enum_haxe_rtti_typetree_TClassdecl).Hx_Field_c
			inf = c
		}

	case 2:
		{
			var e any
			_ = e
			e = t.(Hx_Enum_haxe_rtti_typetree_TEnumdecl).Hx_Field_e
			inf = e
		}

	case 3:
		{
			var t1 any
			_ = t1
			t1 = t.(Hx_Enum_haxe_rtti_typetree_TTypedecl).Hx_Field_t
			inf = t1
		}

	case 4:
		{
			var a any
			_ = a
			a = t.(Hx_Enum_haxe_rtti_typetree_TAbstractdecl).Hx_Field_a
			inf = a
		}

	}

	return inf
}

func Hx_Field_haxe_rtti_typeapi_isVar(t Hx_Enum_haxe_rtti_ctype) bool {
	if t.Hx_Field_enumIndex() == 4 {
		return false
	} else {
		return true
	}
}

func Hx_Field_haxe_rtti_typeapi_leq(f func(_hx_carg0 any, _hx_carg1 any) bool, l1 HxArray[any], l2 HxArray[any]) bool {
	var it_current int32
	_ = it_current
	it_current = 0
	var it_array HxArray[any]
	_ = it_array
	it_array = l2
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(l1)) {
				break
			}

			var e1 any
			_ = e1
			e1 = l1.Get(_g)
			_g++
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = it_current
			if _hx_tmp_1 >= Hx_Field_go_haxe_hxdynamic_getArrayLength(it_array) {
				return false
			}

			var _hx_tmp_2 HxArray[any]
			_ = _hx_tmp_2
			_hx_tmp_2 = it_array
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = it_current
			it_current = (it_current + 1)
			var e2 any
			_ = e2
			e2 = _hx_tmp_2.Get(_hx_tmp_3)
			if !f(e1, e2) {
				return false
			}
		}
	}

	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = it_current
	if _hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(it_array) {
		return false
	}

	return true
}

func Hx_Field_haxe_rtti_typeapi_rightsEq(r1 Hx_Enum_haxe_rtti_rights, r2 Hx_Enum_haxe_rtti_rights) bool {
	if r1 == r2 {
		return true
	}

	if r1.Hx_Field_enumIndex() == 2 {
		var m1 string
		_ = m1
		m1 = r1.(Hx_Enum_haxe_rtti_rights_RCall).Hx_Field_m
		if r2.Hx_Field_enumIndex() == 2 {
			var m2 string
			_ = m2
			m2 = r2.(Hx_Enum_haxe_rtti_rights_RCall).Hx_Field_m
			return (m1 == m2)
		}
	}

	return false
}

func Hx_Field_haxe_rtti_typeapi_typeEq(t1 Hx_Enum_haxe_rtti_ctype, t2 Hx_Enum_haxe_rtti_ctype) bool {
	switch t1.Hx_Field_enumIndex() {
	case 0:
		{
			return (t2 == Hx_Enum_haxe_rtti_ctype_CUnknown{})
		}

	case 1:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t1.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t1.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_params
			if t2.Hx_Field_enumIndex() == 1 {
				var name2 Hx_Typedef_haxe_rtti_path
				_ = name2
				name2 = t2.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_name
				var params2 HxArray[Hx_Enum_haxe_rtti_ctype]
				_ = params2
				params2 = t2.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_params
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				if name == name2 {
					_hx_tmp_0 = func() func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
						var _hx_closure_src_376 func(f func(_hx_carg0 any, _hx_carg1 any) bool, l1 HxArray[any], l2 HxArray[any]) bool
						_ = _hx_closure_src_376
						_hx_closure_src_376 = Hx_Field_haxe_rtti_typeapi_leq
						return func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
							return _hx_closure_src_376(func() func(_hx_carg0 any, _hx_carg1 any) bool {
								var _hx_closure_src_380 func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool
								_ = _hx_closure_src_380
								_hx_closure_src_380 = f
								return func(_hx_carg0 any, _hx_carg1 any) bool {
									return _hx_closure_src_380(func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg0)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}(), func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg1)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}())
								}
							}(), HxMakeArrayView[any](l1), HxMakeArrayView[any](l2))
						}
					}()(Hx_Field_haxe_rtti_typeapi_typeEq, params, params2)
				} else {
					_hx_tmp_0 = false
				}

				return _hx_tmp_0
			}
		}

	case 2:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t1.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t1.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_params
			if t2.Hx_Field_enumIndex() == 2 {
				var name2 Hx_Typedef_haxe_rtti_path
				_ = name2
				name2 = t2.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_name
				var params2 HxArray[Hx_Enum_haxe_rtti_ctype]
				_ = params2
				params2 = t2.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_params
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if name == name2 {
					_hx_tmp_1 = func() func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
						var _hx_closure_src_377 func(f func(_hx_carg0 any, _hx_carg1 any) bool, l1 HxArray[any], l2 HxArray[any]) bool
						_ = _hx_closure_src_377
						_hx_closure_src_377 = Hx_Field_haxe_rtti_typeapi_leq
						return func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
							return _hx_closure_src_377(func() func(_hx_carg0 any, _hx_carg1 any) bool {
								var _hx_closure_src_381 func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool
								_ = _hx_closure_src_381
								_hx_closure_src_381 = f
								return func(_hx_carg0 any, _hx_carg1 any) bool {
									return _hx_closure_src_381(func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg0)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}(), func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg1)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}())
								}
							}(), HxMakeArrayView[any](l1), HxMakeArrayView[any](l2))
						}
					}()(Hx_Field_haxe_rtti_typeapi_typeEq, params, params2)
				} else {
					_hx_tmp_1 = false
				}

				return _hx_tmp_1
			}
		}

	case 3:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t1.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t1.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_params
			if t2.Hx_Field_enumIndex() == 3 {
				var name2 Hx_Typedef_haxe_rtti_path
				_ = name2
				name2 = t2.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_name
				var params2 HxArray[Hx_Enum_haxe_rtti_ctype]
				_ = params2
				params2 = t2.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_params
				var _hx_tmp_2 bool
				_ = _hx_tmp_2
				if name == name2 {
					_hx_tmp_2 = func() func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
						var _hx_closure_src_378 func(f func(_hx_carg0 any, _hx_carg1 any) bool, l1 HxArray[any], l2 HxArray[any]) bool
						_ = _hx_closure_src_378
						_hx_closure_src_378 = Hx_Field_haxe_rtti_typeapi_leq
						return func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
							return _hx_closure_src_378(func() func(_hx_carg0 any, _hx_carg1 any) bool {
								var _hx_closure_src_382 func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool
								_ = _hx_closure_src_382
								_hx_closure_src_382 = f
								return func(_hx_carg0 any, _hx_carg1 any) bool {
									return _hx_closure_src_382(func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg0)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}(), func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg1)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}())
								}
							}(), HxMakeArrayView[any](l1), HxMakeArrayView[any](l2))
						}
					}()(Hx_Field_haxe_rtti_typeapi_typeEq, params, params2)
				} else {
					_hx_tmp_2 = false
				}

				return _hx_tmp_2
			}
		}

	case 4:
		{
			var args HxArray[any]
			_ = args
			args = t1.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_args
			var ret Hx_Enum_haxe_rtti_ctype
			_ = ret
			ret = t1.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_ret
			if t2.Hx_Field_enumIndex() == 4 {
				var args2 HxArray[any]
				_ = args2
				args2 = t2.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_args
				var ret2 Hx_Enum_haxe_rtti_ctype
				_ = ret2
				ret2 = t2.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_ret
				var _hx_tmp_3 bool
				_ = _hx_tmp_3
				if Hx_Field_haxe_rtti_typeapi_leq(func(a any, b any) bool {
					var _hx_tmp_4 bool
					_ = _hx_tmp_4
					var _hx_tmp_5 bool
					_ = _hx_tmp_5
					var _hx_tmp_6 any
					_ = _hx_tmp_6
					_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_6, Hx_Field_go_haxe_hxdynamic_getField(b, "name")))) {
						var _hx_tmp_7 any
						_ = _hx_tmp_7
						_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(a, "opt")
						var _hx_tmp_8 any
						_ = _hx_tmp_8
						_hx_tmp_8 = _hx_tmp_7
						_hx_tmp_5 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_8, Hx_Field_go_haxe_hxdynamic_getField(b, "opt")))))
					} else {
						_hx_tmp_5 = false
					}

					if _hx_tmp_5 {
						var _hx_tmp_9 Hx_Enum_haxe_rtti_ctype
						_ = _hx_tmp_9
						_hx_tmp_9 = func() Hx_Enum_haxe_rtti_ctype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_haxe_rtti_ctype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_haxe_rtti_ctype)
						}()
						var _hx_tmp_10 Hx_Enum_haxe_rtti_ctype
						_ = _hx_tmp_10
						_hx_tmp_10 = _hx_tmp_9
						_hx_tmp_4 = Hx_Field_haxe_rtti_typeapi_typeEq(_hx_tmp_10, func() Hx_Enum_haxe_rtti_ctype {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(b, "t"))
							if _hx_d == nil {
								var _hx_z Hx_Enum_haxe_rtti_ctype
								return _hx_z
							}
							return _hx_d.(Hx_Enum_haxe_rtti_ctype)
						}())
					} else {
						_hx_tmp_4 = false
					}

					return _hx_tmp_4
				}, args, args2) {
					_hx_tmp_3 = Hx_Field_haxe_rtti_typeapi_typeEq(ret, ret2)
				} else {
					_hx_tmp_3 = false
				}

				return _hx_tmp_3
			}
		}

	case 5:
		{
			var fields HxArray[any]
			_ = fields
			fields = t1.(Hx_Enum_haxe_rtti_ctype_CAnonymous).Hx_Field_fields
			if t2.Hx_Field_enumIndex() == 5 {
				var fields2 HxArray[any]
				_ = fields2
				fields2 = t2.(Hx_Enum_haxe_rtti_ctype_CAnonymous).Hx_Field_fields
				return Hx_Field_haxe_rtti_typeapi_leq(func(a any, b any) bool {
					return Hx_Field_haxe_rtti_typeapi_fieldEq(a, b)
				}, fields, fields2)
			}
		}

	case 6:
		{
			var t Hx_Enum_haxe_rtti_ctype
			_ = t
			t = t1.(Hx_Enum_haxe_rtti_ctype_CDynamic).Hx_Field_t
			if t2.Hx_Field_enumIndex() == 6 {
				var t21 Hx_Enum_haxe_rtti_ctype
				_ = t21
				t21 = t2.(Hx_Enum_haxe_rtti_ctype_CDynamic).Hx_Field_t
				if (t == nil) != (t21 == nil) {
					return false
				}

				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				if t == nil {
					_hx_tmp_11 = true
				} else {
					_hx_tmp_11 = Hx_Field_haxe_rtti_typeapi_typeEq(t, t21)
				}

				return _hx_tmp_11
			}
		}

	case 7:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t1.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t1.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_params
			if t2.Hx_Field_enumIndex() == 7 {
				var name2 Hx_Typedef_haxe_rtti_path
				_ = name2
				name2 = t2.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_name
				var params2 HxArray[Hx_Enum_haxe_rtti_ctype]
				_ = params2
				params2 = t2.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_params
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if name == name2 {
					_hx_tmp_12 = func() func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
						var _hx_closure_src_379 func(f func(_hx_carg0 any, _hx_carg1 any) bool, l1 HxArray[any], l2 HxArray[any]) bool
						_ = _hx_closure_src_379
						_hx_closure_src_379 = Hx_Field_haxe_rtti_typeapi_leq
						return func(f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool, l1 HxArray[Hx_Enum_haxe_rtti_ctype], l2 HxArray[Hx_Enum_haxe_rtti_ctype]) bool {
							return _hx_closure_src_379(func() func(_hx_carg0 any, _hx_carg1 any) bool {
								var _hx_closure_src_383 func(_hx_carg0 Hx_Enum_haxe_rtti_ctype, _hx_carg1 Hx_Enum_haxe_rtti_ctype) bool
								_ = _hx_closure_src_383
								_hx_closure_src_383 = f
								return func(_hx_carg0 any, _hx_carg1 any) bool {
									return _hx_closure_src_383(func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg0)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}(), func() Hx_Enum_haxe_rtti_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_carg1)
										if _hx_d == nil {
											var _hx_z Hx_Enum_haxe_rtti_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_haxe_rtti_ctype)
									}())
								}
							}(), HxMakeArrayView[any](l1), HxMakeArrayView[any](l2))
						}
					}()(Hx_Field_haxe_rtti_typeapi_typeEq, params, params2)
				} else {
					_hx_tmp_12 = false
				}

				return _hx_tmp_12
			}
		}

	}

	return false
}

func Hx_Field_haxe_rtti_typeapi_fieldEq(f1 any, f2 any) bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(f1, "name")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(f2, "name")))) {
		return false
	}

	var _hx_tmp_1 Hx_Enum_haxe_rtti_ctype
	_ = _hx_tmp_1
	_hx_tmp_1 = func() Hx_Enum_haxe_rtti_ctype {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "type"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_ctype
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_ctype)
	}()
	if !Hx_Field_haxe_rtti_typeapi_typeEq(_hx_tmp_1, func() Hx_Enum_haxe_rtti_ctype {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "type"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_ctype
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_ctype)
	}()) {
		return false
	}

	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(f1, "isPublic")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_getField(f2, "isPublic")))) {
		return false
	}

	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(f1, "doc")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(f2, "doc")))) {
		return false
	}

	var _hx_tmp_4 Hx_Enum_haxe_rtti_rights
	_ = _hx_tmp_4
	_hx_tmp_4 = func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "get"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()
	if !Hx_Field_haxe_rtti_typeapi_rightsEq(_hx_tmp_4, func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "get"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()) {
		return false
	}

	var _hx_tmp_5 Hx_Enum_haxe_rtti_rights
	_ = _hx_tmp_5
	_hx_tmp_5 = func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "set"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()
	if !Hx_Field_haxe_rtti_typeapi_rightsEq(_hx_tmp_5, func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "set"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()) {
		return false
	}

	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f1, "params"), ((any)(nil))))))
	if _hx_tmp_6 != ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f2, "params"), ((any)(nil)))))) {
		return false
	}

	var _hx_tmp_7 bool
	_ = _hx_tmp_7
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(f1, "params"), ((any)(nil))))) {
		var _hx_tmp_9 any
		_ = _hx_tmp_9
		_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f1, "params"), "join")
		var _hx_tmp_8 any
		_ = _hx_tmp_8
		_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any](((any)(":"))).Dyn())
		var _hx_tmp_10 any
		_ = _hx_tmp_10
		_hx_tmp_10 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(f2, "params"), "join")
		var _hx_tmp_11 any
		_ = _hx_tmp_11
		_hx_tmp_11 = _hx_tmp_8
		var _hx_tmp_12 any
		_ = _hx_tmp_12
		_hx_tmp_12 = _hx_tmp_10
		_hx_tmp_7 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_11, Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_12, HxMakeArray[any](((any)(":"))).Dyn())))))
	} else {
		_hx_tmp_7 = false
	}

	if _hx_tmp_7 {
		return false
	}

	return true
}

func Hx_Field_haxe_rtti_typeapi_constructorEq(c1 any, c2 any) bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(c1, "name")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(c2, "name")))) {
		return false
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(c1, "doc")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(c2, "doc")))) {
		return false
	}

	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	_hx_tmp_2 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c1, "args"), ((any)(nil))))))
	if _hx_tmp_2 != ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(c2, "args"), ((any)(nil)))))) {
		return false
	}

	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(c1, "args"), ((any)(nil))))) {
		var _hx_tmp_4 func(a any, b any) bool
		_ = _hx_tmp_4
		_hx_tmp_4 = func(a any, b any) bool {
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(a, "name")
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_7, Hx_Field_go_haxe_hxdynamic_getField(b, "name")))) {
				var _hx_tmp_8 any
				_ = _hx_tmp_8
				_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(a, "opt")
				var _hx_tmp_9 any
				_ = _hx_tmp_9
				_hx_tmp_9 = _hx_tmp_8
				_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_9, Hx_Field_go_haxe_hxdynamic_getField(b, "opt")))))
			} else {
				_hx_tmp_6 = false
			}

			if _hx_tmp_6 {
				var _hx_tmp_10 Hx_Enum_haxe_rtti_ctype
				_ = _hx_tmp_10
				_hx_tmp_10 = func() Hx_Enum_haxe_rtti_ctype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a, "t"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_haxe_rtti_ctype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_haxe_rtti_ctype)
				}()
				var _hx_tmp_11 Hx_Enum_haxe_rtti_ctype
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_10
				_hx_tmp_5 = Hx_Field_haxe_rtti_typeapi_typeEq(_hx_tmp_11, func() Hx_Enum_haxe_rtti_ctype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(b, "t"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_haxe_rtti_ctype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_haxe_rtti_ctype)
				}())
			} else {
				_hx_tmp_5 = false
			}

			return _hx_tmp_5
		}

		var _hx_tmp_12 HxArray[any]
		_ = _hx_tmp_12
		_hx_tmp_12 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c1, "args"))
		var _hx_tmp_13 func(a any, b any) bool
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_4
		var _hx_tmp_14 HxArray[any]
		_ = _hx_tmp_14
		_hx_tmp_14 = _hx_tmp_12
		_hx_tmp_3 = !Hx_Field_haxe_rtti_typeapi_leq(_hx_tmp_13, _hx_tmp_14, HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c2, "args")))
	} else {
		_hx_tmp_3 = false
	}

	if _hx_tmp_3 {
		return false
	}

	return true
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_rtti_ctypetools() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_rtti_ctypetools() {
}

type Hx_Obj_VTable_haxe_rtti_ctypetools interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_rtti_ctypetools struct {
	VTable Hx_Obj_VTable_haxe_rtti_ctypetools `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_rtti_ctypetools_CreateEmptyInstance() *Hx_Obj_haxe_rtti_ctypetools {
	obj := &Hx_Obj_haxe_rtti_ctypetools{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_rtti_ctypetools_CreateInstance() *Hx_Obj_haxe_rtti_ctypetools {
	obj := Hx_Obj_haxe_rtti_ctypetools_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_rtti_ctypetools_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.rtti.CTypeTools",
	HxMakeArray[string]("toString", "nameWithParams", "functionArgumentName", "classField"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_rtti_ctypetools_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_rtti_ctypetools_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_rtti_ctypetools_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toString":
			return Hx_Field_haxe_rtti_ctypetools_toString
		case "nameWithParams":
			return Hx_Field_haxe_rtti_ctypetools_nameWithParams
		case "functionArgumentName":
			return Hx_Field_haxe_rtti_ctypetools_functionArgumentName
		case "classField":
			return Hx_Field_haxe_rtti_ctypetools_classField
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_rtti_ctypetools) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_rtti_ctypetools_RTTI
}

func (this *Hx_Obj_haxe_rtti_ctypetools) Hx_As_Hx_Obj_haxe_rtti_ctypetools() *Hx_Obj_haxe_rtti_ctypetools {
	return this
}

func Hx_Field_haxe_rtti_ctypetools_toString(t Hx_Enum_haxe_rtti_ctype) string {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			return "unknown"
		}

	case 1:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CEnum).Hx_Field_params
			return Hx_Field_haxe_rtti_ctypetools_nameWithParams(Hx_Field_go_haxe_hxdynamic_toString(name), params)
		}

	case 2:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CClass).Hx_Field_params
			return Hx_Field_haxe_rtti_ctypetools_nameWithParams(Hx_Field_go_haxe_hxdynamic_toString(name), params)
		}

	case 3:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CTypedef).Hx_Field_params
			return Hx_Field_haxe_rtti_ctypetools_nameWithParams(Hx_Field_go_haxe_hxdynamic_toString(name), params)
		}

	case 4:
		{
			var args HxArray[any]
			_ = args
			args = t.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_args
			var ret Hx_Enum_haxe_rtti_ctype
			_ = ret
			ret = t.(Hx_Enum_haxe_rtti_ctype_CFunction).Hx_Field_ret
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(args) == 0 {
				return Hx_Field_go_haxe_hxstring_concat("Void -> ", Hx_Field_haxe_rtti_ctypetools_toString(ret))
			} else {
				var f func(_hx_carg0 any) string
				_ = f
				f = Hx_Field_haxe_rtti_ctypetools_functionArgumentName
				var _g HxArray[string]
				_ = _g
				_g = HxMakeArray[string]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 HxArray[any]
					_ = _g2
					_g2 = args
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
							break
						}

						var v any
						_ = v
						v = _g2.Get(_g1)
						_g1++
						var _hx_tmp_1 HxArray[string]
						_ = _hx_tmp_1
						_hx_tmp_1 = _g
						_ = (Hx_Array_Push[string](_hx_tmp_1, f(v)))
					}
				}

				var _hx_tmp_2 string
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_go_haxe_hxstring_concat((Hx_Array_Join[string](_g, " -> ")), " -> ")
				return Hx_Field_go_haxe_hxstring_concat(_hx_tmp_2, Hx_Field_haxe_rtti_ctypetools_toString(ret))
			}
		}

	case 5:
		{
			var fields HxArray[any]
			_ = fields
			fields = t.(Hx_Enum_haxe_rtti_ctype_CAnonymous).Hx_Field_fields
			var f func(_hx_carg0 any) string
			_ = f
			f = Hx_Field_haxe_rtti_ctypetools_classField
			var _g HxArray[string]
			_ = _g
			_g = HxMakeArray[string]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[any]
				_ = _g2
				_g2 = fields
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g1
					if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
						break
					}

					var v any
					_ = v
					v = _g2.Get(_g1)
					_g1++
					var _hx_tmp_4 HxArray[string]
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					_ = (Hx_Array_Push[string](_hx_tmp_4, f(v)))
				}
			}

			return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("{ ", (Hx_Array_Join[string](_g, ", "))), "}")
		}

	case 6:
		{
			var d Hx_Enum_haxe_rtti_ctype
			_ = d
			d = t.(Hx_Enum_haxe_rtti_ctype_CDynamic).Hx_Field_t
			if d == nil {
				return "Dynamic"
			} else {
				return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Dynamic<", Hx_Field_haxe_rtti_ctypetools_toString(d)), ">")
			}
		}

	case 7:
		{
			var name Hx_Typedef_haxe_rtti_path
			_ = name
			name = t.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_name
			var params HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = params
			params = t.(Hx_Enum_haxe_rtti_ctype_CAbstract).Hx_Field_params
			return Hx_Field_haxe_rtti_ctypetools_nameWithParams(Hx_Field_go_haxe_hxdynamic_toString(name), params)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func Hx_Field_haxe_rtti_ctypetools_nameWithParams(name string, params HxArray[Hx_Enum_haxe_rtti_ctype]) string {
	if params.Len() == 0 {
		return name
	}

	var tmp string
	_ = tmp
	tmp = Hx_Field_go_haxe_hxstring_concat(name, "<")
	var f func(_hx_carg0 Hx_Enum_haxe_rtti_ctype) string
	_ = f
	f = Hx_Field_haxe_rtti_ctypetools_toString
	var _g HxArray[string]
	_ = _g
	_g = HxMakeArray[string]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[Hx_Enum_haxe_rtti_ctype]
		_ = _g2
		_g2 = params
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var v Hx_Enum_haxe_rtti_ctype
			_ = v
			v = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[string]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[string](_hx_tmp_1, f(v)))
		}
	}

	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = tmp
	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_2, (Hx_Array_Join[string](_g, ", "))), ">")
}

func Hx_Field_haxe_rtti_ctypetools_functionArgumentName(arg any) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(arg, "opt")))) {
		_hx_tmp_0 = "?"
	} else {
		_hx_tmp_0 = ""
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(arg, "name"), ((any)(""))))) {
		_hx_tmp_1 = ""
	} else {
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(arg, "name"), ((any)(":"))))
	}

	var _hx_tmp_2 string
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(arg, "value"), nil)))))), ((any)(false))))) {
		_hx_tmp_2 = ""
	} else {
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(((any)(" = ")), Hx_Field_go_haxe_hxdynamic_getField(arg, "value")))
	}

	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat((_hx_tmp_0), (_hx_tmp_1)), Hx_Field_haxe_rtti_ctypetools_toString(func() Hx_Enum_haxe_rtti_ctype {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_ctype
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_ctype)
	}())), (_hx_tmp_2))
}

func Hx_Field_haxe_rtti_ctypetools_classField(cf any) string {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_getField(cf, "name"), ((any)(":")))
	return Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_0, ((any)(Hx_Field_haxe_rtti_ctypetools_toString(func() Hx_Enum_haxe_rtti_ctype {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(cf, "type"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_ctype
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_ctype)
	}())))))
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_ctype]()
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_typetree]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
