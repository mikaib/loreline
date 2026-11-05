package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_intmap() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_intmap() {
}

type Hx_Obj_VTable_haxe_ds_intmap interface {
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_toString() string
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_keys() any
	Hx_Field_iterator() any
	Hx_Field_clear()
	Hx_Field_size() int32
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field_keyValueIterator() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_intmap struct {
	Hx_Obj_haxe_imap
	VTable     Hx_Obj_VTable_haxe_ds_intmap `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_h map[int32]any
}

func Hx_Obj_haxe_ds_intmap_CreateEmptyInstance() *Hx_Obj_haxe_ds_intmap {
	obj := &Hx_Obj_haxe_ds_intmap{}
	obj.VTable = obj
	obj.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Field_h = HxDefault[map[int32]any]()
	return obj
}

func Hx_Obj_haxe_ds_intmap_CreateInstance() *Hx_Obj_haxe_ds_intmap {
	obj := Hx_Obj_haxe_ds_intmap_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_New() {
	{
		this.Hx_Field_h = map[int32]any{}
	}

}

var Hx_Obj_haxe_ds_intmap_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.IntMap",
	HxMakeArray[string](),
	HxMakeArray[string]("h", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "clear", "size"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_intmap_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_intmap_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_intmap_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_intmap_RTTI
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_As_Hx_Obj_haxe_ds_intmap() *Hx_Obj_haxe_ds_intmap {
	return this
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_set(_hx_param_key any, _hx_param_value any) {
	key := ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(_hx_param_key)))
	_ = key
	value := _hx_param_value
	_ = value
	this.Hx_Field_h[key] = value
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_get(_hx_param_key any) any {
	key := ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(_hx_param_key)))
	_ = key
	return this.Hx_Field_h[key]
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_exists(_hx_param_key any) bool {
	key := ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(_hx_param_key)))
	_ = key
	_, ok := this.Hx_Field_h[key]
	return ok
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_remove(_hx_param_key any) bool {
	key := ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(_hx_param_key)))
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

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_keys() any {
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

	var keys []int32
	_ = keys
	keys = make([]int32, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	var self []int32
	_ = self
	self = keys
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any](HxMakeArray(self...)))))
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_iterator() any {
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

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_keyValueIterator() any {
	return ((any)(Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance((func() *Hx_Obj_haxe_imap {
		_hx_icast_10 := this
		if _hx_icast_10 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_10, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_10.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())))
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_intmap
	_ = copied
	copied = Hx_Obj_haxe_ds_intmap_CreateInstance()
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
			var key int32
			_ = key
			key = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))))
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = key
			copied.VTable.Hx_Field_set(_hx_tmp_2, this.VTable.Hx_Field_get(key))
		}
	}

	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_11 := copied
		if _hx_icast_11 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_11, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_11.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_toString() string {
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

	var keys []int32
	_ = keys
	keys = make([]int32, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	var keys1 []int32
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
					var _hx_closure_src_56 func(x any)
					_ = _hx_closure_src_56
					_hx_closure_src_56 = s.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_56(((any)(x)))
					}
				}()(", ")
			}

			func() func(x int32) {
				var _hx_closure_src_57 func(x any)
				_ = _hx_closure_src_57
				_hx_closure_src_57 = s.VTable.Hx_Field_add
				return func(x int32) {
					_hx_closure_src_57(((any)(x)))
				}
			}()(keys1[i])
			func() func(x string) {
				var _hx_closure_src_58 func(x any)
				_ = _hx_closure_src_58
				_hx_closure_src_58 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_58(((any)(x)))
				}
			}()(" => ")
			s.VTable.Hx_Field_add(values1[i])
		}
	}

	s.VTable.Hx_Field_addChar(93)
	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_clear() {
	clear(this.Hx_Field_h)
}

func (this *Hx_Obj_haxe_ds_intmap) Hx_Field_size() int32 {
	return ((int32)(len(this.Hx_Field_h)))
}
