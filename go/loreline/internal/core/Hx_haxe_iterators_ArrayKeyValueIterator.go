package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator() {
}

type Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_arraykeyvalueiterator struct {
	VTable           Hx_Obj_VTable_haxe_iterators_arraykeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_current int32
	Hx_Field_array   HxArray[any]
}

func Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
	obj := &Hx_Obj_haxe_iterators_arraykeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_current = HxDefault[int32]()
	obj.Hx_Field_array = HxDefault[HxArray[any]]()
	return obj
}

func Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateInstance(array HxArray[any]) *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
	obj := Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(array)
	return obj
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_New(array HxArray[any]) {
	{
		this.Hx_Field_current = 0
	}
	{
		this.Hx_Field_array = array
	}

}

var Hx_Obj_haxe_iterators_arraykeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.ArrayKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("current", "array", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateInstance(HxConvert[HxArray[any]](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_arraykeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_arraykeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_arraykeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_As_Hx_Obj_haxe_iterators_arraykeyvalueiterator() *Hx_Obj_haxe_iterators_arraykeyvalueiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_current
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_array))
}

func (this *Hx_Obj_haxe_iterators_arraykeyvalueiterator) Hx_Field_next() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_array.Get(this.Hx_Field_current)
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_current
	this.Hx_Field_current = (this.Hx_Field_current + 1)
	return any(map[string]any{"value": _hx_tmp_0, "key": ((any)(_hx_tmp_1))})
}

func init() {
	hxRegisterArrayConstructor[any]()
}
