package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_restiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_restiterator() {
}

type Hx_Obj_VTable_haxe_iterators_restiterator interface {
	Hx_Field_next() any
	Hx_Field_hasNext() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_restiterator struct {
	VTable           Hx_Obj_VTable_haxe_iterators_restiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_args    HxArray[any]
	Hx_Field_current int32
}

func Hx_Obj_haxe_iterators_restiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_restiterator {
	obj := &Hx_Obj_haxe_iterators_restiterator{}
	obj.VTable = obj
	obj.Hx_Field_args = HxDefault[HxArray[any]]()
	obj.Hx_Field_current = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_iterators_restiterator_CreateInstance(args any) *Hx_Obj_haxe_iterators_restiterator {
	obj := Hx_Obj_haxe_iterators_restiterator_CreateEmptyInstance()
	obj.Hx_New(args)
	return obj
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_New(args any) {
	{
		this.Hx_Field_current = 0
	}
	{
		this.Hx_Field_args = func() HxArray[any] {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(args)
			if _hx_d == nil {
				var _hx_z HxArray[any]
				return _hx_z
			}
			return _hx_d.(HxArray[any])
		}()
	}

}

var Hx_Obj_haxe_iterators_restiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.RestIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("args", "current", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_restiterator_CreateInstance(HxConvert[any](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_restiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_restiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_restiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_As_Hx_Obj_haxe_iterators_restiterator() *Hx_Obj_haxe_iterators_restiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_current
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_args))
}

func (this *Hx_Obj_haxe_iterators_restiterator) Hx_Field_next() any {
	var _hx_tmp_0 HxArray[any]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_args
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_current
	this.Hx_Field_current = (this.Hx_Field_current + 1)
	return _hx_tmp_0.Get(_hx_tmp_1)
}

func init() {
	hxRegisterArrayConstructor[any]()
}
