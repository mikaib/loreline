package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_async() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_async() {
}

type Hx_Obj_VTable_loreline_async interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_async struct {
	VTable        Hx_Obj_VTable_loreline_async `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_func func(done func())
}

func Hx_Obj_loreline_async_CreateEmptyInstance() *Hx_Obj_loreline_async {
	obj := &Hx_Obj_loreline_async{}
	obj.VTable = obj
	obj.Hx_Field_func = HxDefault[func(done func())]()
	return obj
}

func Hx_Obj_loreline_async_CreateInstance(_hx_reserved_func func(done func())) *Hx_Obj_loreline_async {
	obj := Hx_Obj_loreline_async_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_func)
	return obj
}

func (this *Hx_Obj_loreline_async) Hx_New(_hx_reserved_func func(done func())) {
	{
		this.Hx_Field_func = _hx_reserved_func
	}

}

var Hx_Obj_loreline_async_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Async",
	HxMakeArray[string](),
	HxMakeArray[string]("func"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_async_CreateInstance(HxConvert[func(done func())](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_async_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_async_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_async) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_async_RTTI
}

func (this *Hx_Obj_loreline_async) Hx_As_Hx_Obj_loreline_async() *Hx_Obj_loreline_async {
	return this
}
