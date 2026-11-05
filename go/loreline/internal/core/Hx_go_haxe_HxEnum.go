package core

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxenum() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxenum() {
}

type Hx_Obj_VTable_go_haxe_hxenum interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxenum struct {
	VTable                        Hx_Obj_VTable_go_haxe_hxenum `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name                 string
	Hx_Field_constructorNames     HxArray[string]
	Hx_Field_constructorArgCounts HxArray[int32]
	Hx_Field_createByIndex        func(_hx_carg0 int32, _hx_carg1 HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	Hx_Field___meta__             any
}

func Hx_Obj_go_haxe_hxenum_CreateEmptyInstance() *Hx_Obj_go_haxe_hxenum {
	obj := &Hx_Obj_go_haxe_hxenum{}
	obj.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_constructorNames = HxDefault[HxArray[string]]()
	obj.Hx_Field_constructorArgCounts = HxDefault[HxArray[int32]]()
	obj.Hx_Field_createByIndex = HxDefault[func(_hx_carg0 int32, _hx_carg1 HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue]()
	obj.Hx_Field___meta__ = HxDefault[any]()
	return obj
}

func Hx_Obj_go_haxe_hxenum_CreateInstance(name string, constructorNames HxArray[string], constructorArgCounts HxArray[int32], createByIndex func(_hx_carg0 int32, _hx_carg1 HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue, __meta__ any) *Hx_Obj_go_haxe_hxenum {
	obj := Hx_Obj_go_haxe_hxenum_CreateEmptyInstance()
	obj.Hx_New(name, constructorNames, constructorArgCounts, createByIndex, __meta__)
	return obj
}

func (this *Hx_Obj_go_haxe_hxenum) Hx_New(name string, constructorNames HxArray[string], constructorArgCounts HxArray[int32], createByIndex func(_hx_carg0 int32, _hx_carg1 HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue, __meta__ any) {
	{
		this.Hx_Field_name = name
		this.Hx_Field_constructorNames = constructorNames
		this.Hx_Field_constructorArgCounts = constructorArgCounts
		this.Hx_Field_createByIndex = createByIndex
		this.Hx_Field___meta__ = __meta__
		if Hx_Field_go_haxe_hxenum__registry == nil {
			Hx_Field_go_haxe_hxenum__registry = map[string]*Hx_Obj_go_haxe_hxenum{}
		}

		Hx_Field_go_haxe_hxenum__registry[name] = this
	}

}

var Hx_Obj_go_haxe_hxenum_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxEnum",
	HxMakeArray[string]("_registry", "getEnum"),
	HxMakeArray[string]("name", "constructorNames", "constructorArgCounts", "createByIndex", "__meta__"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxenum_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[HxArray[string]](params.Get(1)), HxConvert[HxArray[int32]](params.Get(2)), HxConvert[func(_hx_carg0 int32, _hx_carg1 HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_go_haxe_hxenum_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxenum_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_registry":
			return &Hx_Field_go_haxe_hxenum__registry
		case "getEnum":
			return Hx_Field_go_haxe_hxenum_getEnum
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxenum) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxenum_RTTI
}

func (this *Hx_Obj_go_haxe_hxenum) Hx_As_Hx_Obj_go_haxe_hxenum() *Hx_Obj_go_haxe_hxenum {
	return this
}

var Hx_Field_go_haxe_hxenum__registry map[string]*Hx_Obj_go_haxe_hxenum

func Hx_Field_go_haxe_hxenum_getEnum(name string) *Hx_Obj_go_haxe_hxenum {
	_, ok := Hx_Field_go_haxe_hxenum__registry[name]
	if ok {
		return Hx_Field_go_haxe_hxenum__registry[name]
	} else {
		return nil
	}
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
