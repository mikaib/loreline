package core

var Hx_Abstract_single_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Single",
	HxMakeArray[string]("fromFloat", "fromFloat32", "toFloat", "toFloat32", "toInt"),
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

func Hx_ClassInit_Hx_Obj_VTable__single_single_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable__single_single_impl_() {
}

type Hx_Obj_VTable__single_single_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj__single_single_impl_ struct {
	VTable Hx_Obj_VTable__single_single_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj__single_single_impl__CreateEmptyInstance() *Hx_Obj__single_single_impl_ {
	obj := &Hx_Obj__single_single_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj__single_single_impl__CreateInstance() *Hx_Obj__single_single_impl_ {
	obj := Hx_Obj__single_single_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj__single_single_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"_Single.Single_Impl_",
	HxMakeArray[string]("fromFloat", "fromFloat32", "toFloat", "toFloat32", "toInt"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj__single_single_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj__single_single_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj__single_single_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "fromFloat":
			return Hx_Field__single_single_impl__fromFloat
		case "fromFloat32":
			return Hx_Field__single_single_impl__fromFloat32
		case "toFloat":
			return Hx_Field__single_single_impl__toFloat
		case "toFloat32":
			return Hx_Field__single_single_impl__toFloat32
		case "toInt":
			return Hx_Field__single_single_impl__toInt
		}
		return nil
	}
}

func (this *Hx_Obj__single_single_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj__single_single_impl__RTTI
}

func (this *Hx_Obj__single_single_impl_) Hx_As_Hx_Obj__single_single_impl_() *Hx_Obj__single_single_impl_ {
	return this
}

func Hx_Field__single_single_impl__fromFloat(v float64) float32 {
	return ((float32)(float32(v)))
}

func Hx_Field__single_single_impl__fromFloat32(v float32) float32 {
	return v
}

func Hx_Field__single_single_impl__toFloat(this float32) float64 {
	return ((float64)(float64(this)))
}

func Hx_Field__single_single_impl__toFloat32(this float32) float32 {
	return this
}

func Hx_Field__single_single_impl__toInt(this float32) int32 {
	return ((int32)(int32(this)))
}
