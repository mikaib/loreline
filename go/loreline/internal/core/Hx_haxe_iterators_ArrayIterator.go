package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_arrayiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_arrayiterator() {
}

type Hx_Obj_VTable_haxe_iterators_arrayiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_arrayiterator struct {
	VTable           Hx_Obj_VTable_haxe_iterators_arrayiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_array   HxArray[any]
	Hx_Field_current int32
}

func Hx_Obj_haxe_iterators_arrayiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_arrayiterator {
	obj := &Hx_Obj_haxe_iterators_arrayiterator{}
	obj.VTable = obj
	obj.Hx_Field_array = HxDefault[HxArray[any]]()
	obj.Hx_Field_current = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(array HxArray[any]) *Hx_Obj_haxe_iterators_arrayiterator {
	obj := Hx_Obj_haxe_iterators_arrayiterator_CreateEmptyInstance()
	obj.Hx_New(array)
	return obj
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_New(array HxArray[any]) {
	{
		this.Hx_Field_current = 0
	}
	{
		this.Hx_Field_array = array
	}

}

var Hx_Obj_haxe_iterators_arrayiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.ArrayIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("array", "current", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxConvert[HxArray[any]](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_arrayiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_arrayiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_arrayiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_As_Hx_Obj_haxe_iterators_arrayiterator() *Hx_Obj_haxe_iterators_arrayiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_current
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_array))
}

func (this *Hx_Obj_haxe_iterators_arrayiterator) Hx_Field_next() any {
	var _hx_tmp_0 HxArray[any]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_array
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_current
	this.Hx_Field_current = (this.Hx_Field_current + 1)
	return _hx_tmp_0.Get(_hx_tmp_1)
}

func init() {
	hxRegisterArrayConstructor[any]()
}
