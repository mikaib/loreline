package core

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxclass() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxclass() {
}

type Hx_Obj_VTable_go_haxe_hxclass interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxclass struct {
	VTable                       Hx_Obj_VTable_go_haxe_hxclass `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name                string
	Hx_Field_staticFields        HxArray[string]
	Hx_Field_instanceFields      HxArray[string]
	Hx_Field_superClass          *Hx_Obj_go_haxe_hxclass
	Hx_Field_interfaces          HxArray[*Hx_Obj_go_haxe_hxclass]
	Hx_Field_createInstance      func(_hx_carg0 HxArray[any]) any
	Hx_Field_createEmptyInstance func() any
	Hx_Field___meta__            any
	Hx_Field_getStaticFieldPtr   func(_hx_carg0 string) any
}

func Hx_Obj_go_haxe_hxclass_CreateEmptyInstance() *Hx_Obj_go_haxe_hxclass {
	obj := &Hx_Obj_go_haxe_hxclass{}
	obj.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_staticFields = HxDefault[HxArray[string]]()
	obj.Hx_Field_instanceFields = HxDefault[HxArray[string]]()
	obj.Hx_Field_superClass = HxDefault[*Hx_Obj_go_haxe_hxclass]()
	obj.Hx_Field_interfaces = HxDefault[HxArray[*Hx_Obj_go_haxe_hxclass]]()
	obj.Hx_Field_createInstance = HxDefault[func(_hx_carg0 HxArray[any]) any]()
	obj.Hx_Field_createEmptyInstance = HxDefault[func() any]()
	obj.Hx_Field___meta__ = HxDefault[any]()
	obj.Hx_Field_getStaticFieldPtr = HxDefault[func(_hx_carg0 string) any]()
	return obj
}

func Hx_Obj_go_haxe_hxclass_CreateInstance(name string, staticFields HxArray[string], instanceFields HxArray[string], superClass *Hx_Obj_go_haxe_hxclass, interfaces HxArray[*Hx_Obj_go_haxe_hxclass], createInstance func(_hx_carg0 HxArray[any]) any, createEmptyInstance func() any, __meta__ any, getStaticFieldPtr func(_hx_carg0 string) any) *Hx_Obj_go_haxe_hxclass {
	obj := Hx_Obj_go_haxe_hxclass_CreateEmptyInstance()
	obj.Hx_New(name, staticFields, instanceFields, superClass, interfaces, createInstance, createEmptyInstance, __meta__, getStaticFieldPtr)
	return obj
}

func (this *Hx_Obj_go_haxe_hxclass) Hx_New(name string, staticFields HxArray[string], instanceFields HxArray[string], superClass *Hx_Obj_go_haxe_hxclass, interfaces HxArray[*Hx_Obj_go_haxe_hxclass], createInstance func(_hx_carg0 HxArray[any]) any, createEmptyInstance func() any, __meta__ any, getStaticFieldPtr func(_hx_carg0 string) any) {
	{
		this.Hx_Field_name = name
		this.Hx_Field_staticFields = staticFields
		this.Hx_Field_instanceFields = instanceFields
		this.Hx_Field_superClass = superClass
		this.Hx_Field_interfaces = interfaces
		this.Hx_Field_createInstance = createInstance
		this.Hx_Field_createEmptyInstance = createEmptyInstance
		this.Hx_Field___meta__ = __meta__
		this.Hx_Field_getStaticFieldPtr = getStaticFieldPtr
		if Hx_Field_go_haxe_hxclass__registry == nil {
			Hx_Field_go_haxe_hxclass__registry = map[string]*Hx_Obj_go_haxe_hxclass{}
		}

		Hx_Field_go_haxe_hxclass__registry[name] = this
	}

}

var Hx_Obj_go_haxe_hxclass_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxClass",
	HxMakeArray[string]("_registry", "getClass"),
	HxMakeArray[string]("name", "staticFields", "instanceFields", "superClass", "interfaces", "createInstance", "createEmptyInstance", "__meta__", "getStaticFieldPtr"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxclass_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[HxArray[string]](params.Get(1)), HxConvert[HxArray[string]](params.Get(2)), HxConvert[*Hx_Obj_go_haxe_hxclass](params.Get(3)), HxConvert[HxArray[*Hx_Obj_go_haxe_hxclass]](params.Get(4)), HxConvert[func(_hx_carg0 HxArray[any]) any](params.Get(5)), HxConvert[func() any](params.Get(6)), HxConvert[any](params.Get(7)), HxConvert[func(_hx_carg0 string) any](params.Get(8)))
	},
	func() any {
		return Hx_Obj_go_haxe_hxclass_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxclass_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "_registry":
			return &Hx_Field_go_haxe_hxclass__registry
		case "getClass":
			return Hx_Field_go_haxe_hxclass_getClass
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxclass) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxclass_RTTI
}

func (this *Hx_Obj_go_haxe_hxclass) Hx_As_Hx_Obj_go_haxe_hxclass() *Hx_Obj_go_haxe_hxclass {
	return this
}

var Hx_Field_go_haxe_hxclass__registry map[string]*Hx_Obj_go_haxe_hxclass

func Hx_Field_go_haxe_hxclass_getClass(name string) *Hx_Obj_go_haxe_hxclass {
	_, ok := Hx_Field_go_haxe_hxclass__registry[name]
	if ok {
		return Hx_Field_go_haxe_hxclass__registry[name]
	} else {
		return nil
	}
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_go_haxe_hxclass]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
