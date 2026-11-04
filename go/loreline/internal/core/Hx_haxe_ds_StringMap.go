package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__stringmap_stringmapkeysiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__stringmap_stringmapkeysiterator() {
}

type Hx_Obj_VTable_haxe_ds__stringmap_stringmapkeysiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator struct {
	VTable          Hx_Obj_VTable_haxe_ds__stringmap_stringmapkeysiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_arr    []string
	Hx_Field_pos    int32
	Hx_Field_length int32
}

func Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateEmptyInstance() *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator {
	obj := &Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator{}
	obj.VTable = obj
	obj.Hx_Field_arr = HxDefault[[]string]()
	obj.Hx_Field_pos = HxDefault[int32]()
	obj.Hx_Field_length = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateInstance(h []string) *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator {
	obj := Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateEmptyInstance()
	obj.Hx_New(h)
	return obj
}

func (this *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator) Hx_New(h []string) {
	{
		this.Hx_Field_arr = h
		this.Hx_Field_pos = 0
		this.Hx_Field_length = ((int32)(len(this.Hx_Field_arr)))
	}

}

var Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._StringMap.StringMapKeysIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("arr", "pos", "length", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateInstance(HxConvert[[]string](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_RTTI
}

func (this *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator) Hx_As_Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator() *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator {
	return this
}

func (this *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator) Hx_Field_hasNext() bool {
	return (this.Hx_Field_pos < this.Hx_Field_length)
}

func (this *Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator) Hx_Field_next() string {
	var _hx_tmp_0 []string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_arr
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_pos
	this.Hx_Field_pos = (this.Hx_Field_pos + 1)
	var b string
	_ = b
	b = _hx_tmp_0[_hx_tmp_1]
	return b
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_stringmap() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_stringmap() {
}

type Hx_Obj_VTable_haxe_ds_stringmap interface {
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field_keys() any
	Hx_Field_iterator() any
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_toString() string
	Hx_Field_size() int32
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_keyValueIterator() any
	Hx_Field_clear()
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_stringmap struct {
	Hx_Obj_haxe_imap
	VTable     Hx_Obj_VTable_haxe_ds_stringmap `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_h map[string]any
}

func Hx_Obj_haxe_ds_stringmap_CreateEmptyInstance() *Hx_Obj_haxe_ds_stringmap {
	obj := &Hx_Obj_haxe_ds_stringmap{}
	obj.VTable = obj
	obj.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Field_h = HxDefault[map[string]any]()
	return obj
}

func Hx_Obj_haxe_ds_stringmap_CreateInstance() *Hx_Obj_haxe_ds_stringmap {
	obj := Hx_Obj_haxe_ds_stringmap_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_New() {
	{
		this.Hx_Field_h = nil
	}
	{
		this.Hx_Field_h = map[string]any{}
	}

}

var Hx_Obj_haxe_ds_stringmap_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.StringMap",
	HxMakeArray[string](),
	HxMakeArray[string]("h", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "size", "clear"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_stringmap_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_stringmap_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_stringmap_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_stringmap_RTTI
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap {
	return this
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_set(_hx_param_key any, _hx_param_value any) {
	key := Hx_Field_go_haxe_hxdynamic_toString(_hx_param_key)
	_ = key
	value := _hx_param_value
	_ = value
	this.Hx_Field_h[key] = value
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_get(_hx_param_key any) any {
	key := Hx_Field_go_haxe_hxdynamic_toString(_hx_param_key)
	_ = key
	return this.Hx_Field_h[key]
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_exists(_hx_param_key any) bool {
	key := Hx_Field_go_haxe_hxdynamic_toString(_hx_param_key)
	_ = key
	_, ok := this.Hx_Field_h[key]
	return ok
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_remove(_hx_param_key any) bool {
	key := Hx_Field_go_haxe_hxdynamic_toString(_hx_param_key)
	_ = key
	var prevSize int
	_ = prevSize
	prevSize = len(this.Hx_Field_h)
	delete(this.Hx_Field_h, key)
	var _hx_tmp_0 int
	_ = _hx_tmp_0
	_hx_tmp_0 = prevSize
	return (_hx_tmp_0 != len(this.Hx_Field_h))
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_keys() any {
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

	var keys []string
	_ = keys
	keys = make([]string, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	return ((any)(Hx_Obj_haxe_ds__stringmap_stringmapkeysiterator_CreateInstance(keys)))
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_iterator() any {
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

	var values []any
	_ = values
	values = make([]any, _hx_tmp_0)
	for _, v := range this.Hx_Field_h {
		values = append(values, v)
	}

	var self []any
	_ = self
	self = values
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArray(self...))))
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_keyValueIterator() any {
	return ((any)(Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance((func() *Hx_Obj_haxe_imap {
		_hx_icast_7 := this
		if _hx_icast_7 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_7, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_7.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())))
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_stringmap
	_ = copied
	copied = Hx_Obj_haxe_ds_stringmap_CreateInstance()
	{
		var key any
		_ = key
		key = this.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(key, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(key, "next")
			var key1 string
			_ = key1
			key1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = key1
			copied.VTable.Hx_Field_set(_hx_tmp_2, this.VTable.Hx_Field_get(key1))
		}
	}

	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_8 := copied
		if _hx_icast_8 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_8, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_8.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_toString() string {
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

	var keys []string
	_ = keys
	keys = make([]string, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	var keys1 []string
	_ = keys1
	keys1 = keys
	var length1 HxNullable[int32]
	_ = length1
	length1 = HxNullable[int32]{}

	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (length1).Valid != false {
		_hx_tmp_1 = (length1).Value
	} else {
		_hx_tmp_1 = 0
	}

	var values []any
	_ = values
	values = make([]any, _hx_tmp_1)
	for _, v := range this.Hx_Field_h {
		values = append(values, v)
	}

	var values1 []any
	_ = values1
	values1 = values
	s.VTable.Hx_Field_addChar(91)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ((int32)(len(keys1)))
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_53 func(x any)
					_ = _hx_closure_src_53
					_hx_closure_src_53 = s.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_53(((any)(x)))
					}
				}()(", ")
			}

			var k string
			_ = k
			k = keys1[i]
			func() func(x string) {
				var _hx_closure_src_54 func(x any)
				_ = _hx_closure_src_54
				_hx_closure_src_54 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_54(((any)(x)))
				}
			}()(k)
			func() func(x string) {
				var _hx_closure_src_55 func(x any)
				_ = _hx_closure_src_55
				_hx_closure_src_55 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_55(((any)(x)))
				}
			}()(" => ")
			s.VTable.Hx_Field_add(values1[i])
		}
	}

	s.VTable.Hx_Field_addChar(93)
	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_size() int32 {
	return ((int32)(len(this.Hx_Field_h)))
}

func (this *Hx_Obj_haxe_ds_stringmap) Hx_Field_clear() {
	clear(this.Hx_Field_h)
}
