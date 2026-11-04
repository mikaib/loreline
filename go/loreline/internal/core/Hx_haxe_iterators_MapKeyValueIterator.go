package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_mapkeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_mapkeyvalueiterator() {
}

type Hx_Obj_VTable_haxe_iterators_mapkeyvalueiterator interface {
	Hx_Field_next() any
	Hx_Field_hasNext() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_mapkeyvalueiterator struct {
	VTable        Hx_Obj_VTable_haxe_iterators_mapkeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map  *Hx_Obj_haxe_imap
	Hx_Field_keys any
}

func Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_mapkeyvalueiterator {
	obj := &Hx_Obj_haxe_iterators_mapkeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_haxe_imap]()
	obj.Hx_Field_keys = HxDefault[any]()
	return obj
}

func Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance(_hx_reserved_map *Hx_Obj_haxe_imap) *Hx_Obj_haxe_iterators_mapkeyvalueiterator {
	obj := Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_haxe_iterators_mapkeyvalueiterator) Hx_New(_hx_reserved_map *Hx_Obj_haxe_imap) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_keys = _hx_reserved_map.VTable.Hx_Field_keys()
	}

}

var Hx_Obj_haxe_iterators_mapkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.MapKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "keys", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance(HxConvert[*Hx_Obj_haxe_imap](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_mapkeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_mapkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_mapkeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_mapkeyvalueiterator) Hx_As_Hx_Obj_haxe_iterators_mapkeyvalueiterator() *Hx_Obj_haxe_iterators_mapkeyvalueiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_mapkeyvalueiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_keys, "hasNext")
	return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))))
}

func (this *Hx_Obj_haxe_iterators_mapkeyvalueiterator) Hx_Field_next() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(this.Hx_Field_keys, "next")
	var key any
	_ = key
	key = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
	return any(map[string]any{"value": this.Hx_Field_map.VTable.Hx_Field_get(key), "key": key})
}
