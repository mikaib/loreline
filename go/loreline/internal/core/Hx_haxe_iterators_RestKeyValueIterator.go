package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator() {
}

type Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator interface {
	Hx_Field_next() any
	Hx_Field_hasNext() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_restkeyvalueiterator struct {
	VTable           Hx_Obj_VTable_haxe_iterators_restkeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_args    HxArray[any]
	Hx_Field_current int32
}

func Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_restkeyvalueiterator {
	obj := &Hx_Obj_haxe_iterators_restkeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_args = HxDefault[HxArray[any]]()
	obj.Hx_Field_current = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateInstance(args any) *Hx_Obj_haxe_iterators_restkeyvalueiterator {
	obj := Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(args)
	return obj
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_New(args any) {
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

var Hx_Obj_haxe_iterators_restkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.RestKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("args", "current", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateInstance(HxConvert[any](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_restkeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_restkeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_restkeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_As_Hx_Obj_haxe_iterators_restkeyvalueiterator() *Hx_Obj_haxe_iterators_restkeyvalueiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_current
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_args))
}

func (this *Hx_Obj_haxe_iterators_restkeyvalueiterator) Hx_Field_next() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(this.Hx_Field_current))
	var _hx_tmp_1 HxArray[any]
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_args
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_current
	this.Hx_Field_current = (this.Hx_Field_current + 1)
	return any(map[string]any{"key": _hx_tmp_0, "value": _hx_tmp_1.Get(_hx_tmp_2)})
}

func init() {
	hxRegisterArrayConstructor[any]()
}
