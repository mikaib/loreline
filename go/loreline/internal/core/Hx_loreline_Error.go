package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_error() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_error() {
}

type Hx_Obj_VTable_loreline_error interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_error struct {
	VTable           Hx_Obj_VTable_loreline_error `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_message string
	Hx_Field_pos     *Hx_Obj_loreline_position
	Hx_Field_stack   HxArray[Hx_Enum_haxe_stackitem]
}

func Hx_Obj_loreline_error_CreateEmptyInstance() *Hx_Obj_loreline_error {
	obj := &Hx_Obj_loreline_error{}
	obj.VTable = obj
	obj.Hx_Field_message = HxDefault[string]()
	obj.Hx_Field_pos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_stack = HxDefault[HxArray[Hx_Enum_haxe_stackitem]]()
	return obj
}

func Hx_Obj_loreline_error_CreateInstance(message string, pos__isOptional any) *Hx_Obj_loreline_error {
	obj := Hx_Obj_loreline_error_CreateEmptyInstance()
	obj.Hx_New(message, pos__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_error) Hx_New(message string, pos__isOptional any) {
	{
		pos := HxDefault[*Hx_Obj_loreline_position]()
		_ = pos

		if pos__isOptional != nil {
			pos = HxConvert[*Hx_Obj_loreline_position](pos__isOptional)
		} else {
			pos = nil
		}
		{
			this.Hx_Field_message = message
			this.Hx_Field_pos = pos
			this.Hx_Field_stack = Hx_Field_haxe__callstack_callstack_impl__callStack()
		}

	}
}

var Hx_Obj_loreline_error_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Error",
	HxMakeArray[string](),
	HxMakeArray[string]("message", "pos", "stack", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_error_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_error_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_error_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_error) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_error_RTTI
}

func (this *Hx_Obj_loreline_error) Hx_As_Hx_Obj_loreline_error() *Hx_Obj_loreline_error {
	return this
}

func (this *Hx_Obj_loreline_error) Hx_Field_toString() string {
	if this.Hx_Field_pos != nil {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("", this.Hx_Field_message), " at ")
		return Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_0, this.Hx_Field_pos.VTable.Hx_Field_toString())
	} else {
		return this.Hx_Field_message
	}
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
