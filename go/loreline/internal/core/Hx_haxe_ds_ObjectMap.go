package core

import "reflect"

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_bucket() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_bucket() {
}

type Hx_Obj_VTable_haxe_ds_bucket interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_bucket struct {
	VTable         Hx_Obj_VTable_haxe_ds_bucket `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_key   any
	Hx_Field_value any
}

func Hx_Obj_haxe_ds_bucket_CreateEmptyInstance() *Hx_Obj_haxe_ds_bucket {
	obj := &Hx_Obj_haxe_ds_bucket{}
	obj.VTable = obj
	obj.Hx_Field_key = HxDefault[any]()
	obj.Hx_Field_value = HxDefault[any]()
	return obj
}

func Hx_Obj_haxe_ds_bucket_CreateInstance(key any, value any) *Hx_Obj_haxe_ds_bucket {
	obj := Hx_Obj_haxe_ds_bucket_CreateEmptyInstance()
	obj.Hx_New(key, value)
	return obj
}

func (this *Hx_Obj_haxe_ds_bucket) Hx_New(key any, value any) {
	{
		this.Hx_Field_key = key
		this.Hx_Field_value = value
	}

}

var Hx_Obj_haxe_ds_bucket_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.Bucket",
	HxMakeArray[string](),
	HxMakeArray[string]("key", "value"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_bucket_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_haxe_ds_bucket_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_bucket_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_bucket) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_bucket_RTTI
}

func (this *Hx_Obj_haxe_ds_bucket) Hx_As_Hx_Obj_haxe_ds_bucket() *Hx_Obj_haxe_ds_bucket {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_objectmap() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_objectmap() {
}

type Hx_Obj_VTable_haxe_ds_objectmap interface {
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field_keyValueIterator() any
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_toString() string
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_keys() any
	Hx_Field_iterator() any
	Hx_Field_clear()
	Hx_Field_size() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_objectmap struct {
	Hx_Obj_haxe_imap
	VTable     Hx_Obj_VTable_haxe_ds_objectmap `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_h map[uintptr]*Hx_Obj_haxe_ds_bucket
}

func Hx_Obj_haxe_ds_objectmap_CreateEmptyInstance() *Hx_Obj_haxe_ds_objectmap {
	obj := &Hx_Obj_haxe_ds_objectmap{}
	obj.VTable = obj
	obj.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Field_h = HxDefault[map[uintptr]*Hx_Obj_haxe_ds_bucket]()
	return obj
}

func Hx_Obj_haxe_ds_objectmap_CreateInstance() *Hx_Obj_haxe_ds_objectmap {
	obj := Hx_Obj_haxe_ds_objectmap_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_New() {
	{
		this.Hx_Field_h = map[uintptr]*Hx_Obj_haxe_ds_bucket{}
	}

}

var Hx_Obj_haxe_ds_objectmap_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.ObjectMap",
	HxMakeArray[string]("id"),
	HxMakeArray[string]("h", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "clear", "size"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_objectmap_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_objectmap_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_objectmap_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "id":
			return Hx_Field_haxe_ds_objectmap_id
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_objectmap_RTTI
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap {
	return this
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_set(_hx_param_key any, _hx_param_value any) {
	key := _hx_param_key
	_ = key
	value := _hx_param_value
	_ = value
	var _hx_tmp_0 map[uintptr]*Hx_Obj_haxe_ds_bucket
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_h
	var _hx_tmp_1 uintptr
	_ = _hx_tmp_1
	_hx_tmp_1 = reflect.ValueOf(key).Pointer()
	_hx_tmp_0[_hx_tmp_1] = Hx_Obj_haxe_ds_bucket_CreateInstance(key, value)
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_get(_hx_param_key any) any {
	key := _hx_param_key
	_ = key
	var _hx_tmp_0 map[uintptr]*Hx_Obj_haxe_ds_bucket
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_h
	_, ok := _hx_tmp_0[reflect.ValueOf(key).Pointer()]
	if !ok {
		return nil
	}

	var _hx_tmp_1 map[uintptr]*Hx_Obj_haxe_ds_bucket
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_h
	return _hx_tmp_1[reflect.ValueOf(key).Pointer()].Hx_Field_value
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_exists(_hx_param_key any) bool {
	key := _hx_param_key
	_ = key
	var _hx_tmp_0 map[uintptr]*Hx_Obj_haxe_ds_bucket
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_h
	_, ok := _hx_tmp_0[reflect.ValueOf(key).Pointer()]
	return ok
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_remove(_hx_param_key any) bool {
	key := _hx_param_key
	_ = key
	var k uintptr
	_ = k
	k = reflect.ValueOf(key).Pointer()
	_, ok := this.Hx_Field_h[k]
	if !ok {
		return false
	}

	delete(this.Hx_Field_h, k)
	return true
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_keys() any {
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var values []*Hx_Obj_haxe_ds_bucket
	_ = values
	values = make([]*Hx_Obj_haxe_ds_bucket, _hx_tmp_0)
	for _, v := range this.Hx_Field_h {
		values = append(values, v)
	}

	var self []*Hx_Obj_haxe_ds_bucket
	_ = self
	self = values
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_haxe_ds_bucket]
		_ = _g2
		_g2 = HxMakeArray(self...)
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < _g2.Len()) {
				break
			}

			var v *Hx_Obj_haxe_ds_bucket
			_ = v
			v = _g2.Get(_g1)
			_g1++
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), v.Hx_Field_key))
		}
	}

	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(_g)))
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_iterator() any {
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var values []*Hx_Obj_haxe_ds_bucket
	_ = values
	values = make([]*Hx_Obj_haxe_ds_bucket, _hx_tmp_0)
	for _, v := range this.Hx_Field_h {
		values = append(values, v)
	}

	var self []*Hx_Obj_haxe_ds_bucket
	_ = self
	self = values
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_haxe_ds_bucket]
		_ = _g2
		_g2 = HxMakeArray(self...)
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < _g2.Len()) {
				break
			}

			var v *Hx_Obj_haxe_ds_bucket
			_ = v
			v = _g2.Get(_g1)
			_g1++
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), v.Hx_Field_value))
		}
	}

	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(_g)))
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_keyValueIterator() any {
	return ((any)(Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance((func() *Hx_Obj_haxe_imap {
		_hx_icast_19 := this
		if _hx_icast_19 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_19, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_19.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())))
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_objectmap
	_ = copied
	copied = Hx_Obj_haxe_ds_objectmap_CreateInstance()
	{
		var _gkey any
		_ = _gkey
		_gkey = this.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "next")
			var key any
			_ = key
			key = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn())
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = key
			copied.VTable.Hx_Field_set(_hx_tmp_2, this.VTable.Hx_Field_get(key))
		}
	}

	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_20 := copied
		if _hx_icast_20 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_20, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_20.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_toString() string {
	var s *Hx_Obj_stringbuf
	_ = s
	s = Hx_Obj_stringbuf_CreateInstance()
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var values []*Hx_Obj_haxe_ds_bucket
	_ = values
	values = make([]*Hx_Obj_haxe_ds_bucket, _hx_tmp_0)
	for _, v := range this.Hx_Field_h {
		values = append(values, v)
	}

	var buckets []*Hx_Obj_haxe_ds_bucket
	_ = buckets
	buckets = values
	s.VTable.Hx_Field_addChar(91)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ((int32)(len(buckets)))
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_62 func(x any)
					_ = _hx_closure_src_62
					_hx_closure_src_62 = s.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_62(((any)(x)))
					}
				}()(", ")
			}

			s.VTable.Hx_Field_add(buckets[i].Hx_Field_key)
			func() func(x string) {
				var _hx_closure_src_63 func(x any)
				_ = _hx_closure_src_63
				_hx_closure_src_63 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_63(((any)(x)))
				}
			}()(" => ")
			s.VTable.Hx_Field_add(buckets[i].Hx_Field_value)
		}
	}

	s.VTable.Hx_Field_addChar(93)
	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_clear() {
	clear(this.Hx_Field_h)
}

func (this *Hx_Obj_haxe_ds_objectmap) Hx_Field_size() int32 {
	return ((int32)(len(this.Hx_Field_h)))
}

func Hx_Field_haxe_ds_objectmap_id(key any) uintptr {
	return reflect.ValueOf(key).Pointer()
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_haxe_ds_bucket]()
	hxRegisterArrayConstructor[any]()
}
