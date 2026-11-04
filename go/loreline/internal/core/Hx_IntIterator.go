package core

func Hx_ClassInit_Hx_Obj_VTable_intiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_intiterator() {
}

type Hx_Obj_VTable_intiterator interface {
	Hx_Field_next() int32
	Hx_Field_hasNext() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_intiterator struct {
	VTable       Hx_Obj_VTable_intiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_min int32
	Hx_Field_max int32
}

func Hx_Obj_intiterator_CreateEmptyInstance() *Hx_Obj_intiterator {
	obj := &Hx_Obj_intiterator{}
	obj.VTable = obj
	obj.Hx_Field_min = HxDefault[int32]()
	obj.Hx_Field_max = HxDefault[int32]()
	return obj
}

func Hx_Obj_intiterator_CreateInstance(min int32, max int32) *Hx_Obj_intiterator {
	obj := Hx_Obj_intiterator_CreateEmptyInstance()
	obj.Hx_New(min, max)
	return obj
}

func (this *Hx_Obj_intiterator) Hx_New(min int32, max int32) {
	{
		this.Hx_Field_min = min
		this.Hx_Field_max = max
	}

}

var Hx_Obj_intiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"IntIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("min", "max", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_intiterator_CreateInstance(HxConvert[int32](params.Get(0)), HxConvert[int32](params.Get(1)))
	},
	func() any {
		return Hx_Obj_intiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_intiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_intiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_intiterator_RTTI
}

func (this *Hx_Obj_intiterator) Hx_As_Hx_Obj_intiterator() *Hx_Obj_intiterator {
	return this
}

func (this *Hx_Obj_intiterator) Hx_Field_hasNext() bool {
	return (this.Hx_Field_min < this.Hx_Field_max)
}

func (this *Hx_Obj_intiterator) Hx_Field_next() int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_min
	this.Hx_Field_min = (this.Hx_Field_min + 1)
	return _hx_tmp_0
}
