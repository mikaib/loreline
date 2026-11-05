package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_iterators_stringiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_iterators_stringiterator() {
}

type Hx_Obj_VTable_haxe_iterators_stringiterator interface {
	Hx_Field_next() int32
	Hx_Field_hasNext() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_iterators_stringiterator struct {
	VTable          Hx_Obj_VTable_haxe_iterators_stringiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_offset int32
	Hx_Field_s      string
}

func Hx_Obj_haxe_iterators_stringiterator_CreateEmptyInstance() *Hx_Obj_haxe_iterators_stringiterator {
	obj := &Hx_Obj_haxe_iterators_stringiterator{}
	obj.VTable = obj
	obj.Hx_Field_offset = HxDefault[int32]()
	obj.Hx_Field_s = HxDefault[string]()
	return obj
}

func Hx_Obj_haxe_iterators_stringiterator_CreateInstance(s string) *Hx_Obj_haxe_iterators_stringiterator {
	obj := Hx_Obj_haxe_iterators_stringiterator_CreateEmptyInstance()
	obj.Hx_New(s)
	return obj
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_New(s string) {
	{
		this.Hx_Field_offset = 0
	}
	{
		this.Hx_Field_s = s
	}

}

var Hx_Obj_haxe_iterators_stringiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.iterators.StringIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("offset", "s", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_iterators_stringiterator_CreateInstance(HxConvert[string](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_iterators_stringiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_iterators_stringiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_iterators_stringiterator_RTTI
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_As_Hx_Obj_haxe_iterators_stringiterator() *Hx_Obj_haxe_iterators_stringiterator {
	return this
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_offset
	return (_hx_tmp_0 < int32(utf8.RuneCountInString(this.Hx_Field_s)))
}

func (this *Hx_Obj_haxe_iterators_stringiterator) Hx_Field_next() int32 {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_s
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_offset
	this.Hx_Field_offset = (this.Hx_Field_offset + 1)
	return (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, _hx_tmp_1)).Value
}
