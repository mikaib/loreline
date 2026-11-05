package core

var Hx_Abstract_go_result_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.Result",
	HxMakeArray[string]("sure", "or", "isOk", "tuple"),
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

func Hx_ClassInit_Hx_Obj_VTable_go__result_result_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_go__result_result_impl_() {
}

type Hx_Obj_VTable_go__result_result_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go__result_result_impl_ struct {
	VTable Hx_Obj_VTable_go__result_result_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go__result_result_impl__CreateEmptyInstance() *Hx_Obj_go__result_result_impl_ {
	obj := &Hx_Obj_go__result_result_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go__result_result_impl__CreateInstance() *Hx_Obj_go__result_result_impl_ {
	obj := Hx_Obj_go__result_result_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_go__result_result_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go._Result.Result_Impl_",
	HxMakeArray[string]("sure", "or", "isOk", "tuple"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go__result_result_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_go__result_result_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go__result_result_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_go__result_result_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go__result_result_impl__RTTI
}

func (this *Hx_Obj_go__result_result_impl_) Hx_As_Hx_Obj_go__result_result_impl_() *Hx_Obj_go__result_result_impl_ {
	return this
}
