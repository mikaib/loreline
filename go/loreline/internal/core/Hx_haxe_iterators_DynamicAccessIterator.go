package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_dynamicaccessiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_dynamicaccessiterator() {
}

type Hx_Obj_VTable_haxe_iterators_dynamicaccessiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_dynamicaccessiterator struct {
	VTable          Hx_Obj_VTable_haxe_iterators_dynamicaccessiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_access any
	Hx_Field_keys   HxArray[string]
	Hx_Field_index  int32
}

func Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_dynamicaccessiterator {
	obj := &Hx_Obj_haxe_iterators_dynamicaccessiterator{}
	obj.VTable = obj
	obj.Hx_Field_access = HxDefault[any]()
	obj.Hx_Field_keys = HxDefault[HxArray[string]]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateInstance(access any) *Hx_Obj_haxe_iterators_dynamicaccessiterator {
	obj := Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateEmptyInstance()
	obj.Hx_New(access)
	return obj
}

func (this *Hx_Obj_haxe_iterators_dynamicaccessiterator) Hx_New(access any) {
	{
		this.Hx_Field_access = access
		this.Hx_Field_keys = Hx_Field_reflect_fields(access)
		this.Hx_Field_index = 0
	}

}

var Hx_Obj_haxe_iterators_dynamicaccessiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.DynamicAccessIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("access", "keys", "index", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateInstance(HxConvert[any](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_dynamicaccessiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_dynamicaccessiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_dynamicaccessiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_dynamicaccessiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_dynamicaccessiterator) Hx_As_Hx_Obj_haxe_iterators_dynamicaccessiterator() *Hx_Obj_haxe_iterators_dynamicaccessiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_dynamicaccessiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < this.Hx_Field_keys.Len())
}

func (this *Hx_Obj_haxe_iterators_dynamicaccessiterator) Hx_Field_next() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_access
	var _hx_tmp_1 HxArray[string]
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_keys
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_index
	this.Hx_Field_index = (this.Hx_Field_index + 1)
	return Hx_Field_reflect_field(_hx_tmp_0, _hx_tmp_1.Get(_hx_tmp_2))
}

func init() {
	hxRegisterArrayConstructor[string]()
}
