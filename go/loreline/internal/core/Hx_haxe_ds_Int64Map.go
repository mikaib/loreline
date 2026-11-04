package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_int64map() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_int64map() {
}

type Hx_Obj_VTable_haxe_ds_int64map interface {
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field_iterator() any
	Hx_Field_clear()
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_keys() any
	Hx_Field_keyValueIterator() any
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_toString() string
	Hx_Field_size() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_int64map struct {
	Hx_Obj_haxe_imap
	VTable     Hx_Obj_VTable_haxe_ds_int64map `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_h map[Hx_Typedef_haxe__int64___int64]any
}

func Hx_Obj_haxe_ds_int64map_CreateEmptyInstance() *Hx_Obj_haxe_ds_int64map {
	obj := &Hx_Obj_haxe_ds_int64map{}
	obj.VTable = obj
	obj.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Field_h = HxDefault[map[Hx_Typedef_haxe__int64___int64]any]()
	return obj
}

func Hx_Obj_haxe_ds_int64map_CreateInstance() *Hx_Obj_haxe_ds_int64map {
	obj := Hx_Obj_haxe_ds_int64map_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_New() {
	{
		this.Hx_Field_h = nil
	}
	{
		this.Hx_Field_h = map[Hx_Typedef_haxe__int64___int64]any{}
	}

}

var Hx_Obj_haxe_ds_int64map_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.Int64Map",
	HxMakeArray[string](),
	HxMakeArray[string]("h", "set", "get", "exists", "remove", "keys", "iterator", "keyValueIterator", "copy", "toString", "clear", "size"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_int64map_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_int64map_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_int64map_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_int64map_RTTI
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_As_Hx_Obj_haxe_ds_int64map() *Hx_Obj_haxe_ds_int64map {
	return this
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_set(_hx_param_key any, _hx_param_value any) {
	key := ((Hx_Typedef_haxe__int64___int64)(Hx_Field_go_haxe_hxdynamic_toInt64(_hx_param_key)))
	_ = key
	value := _hx_param_value
	_ = value
	this.Hx_Field_h[key] = value
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_get(_hx_param_key any) any {
	key := ((Hx_Typedef_haxe__int64___int64)(Hx_Field_go_haxe_hxdynamic_toInt64(_hx_param_key)))
	_ = key
	return this.Hx_Field_h[key]
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_exists(_hx_param_key any) bool {
	key := ((Hx_Typedef_haxe__int64___int64)(Hx_Field_go_haxe_hxdynamic_toInt64(_hx_param_key)))
	_ = key
	_, ok := this.Hx_Field_h[key]
	return ok
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_remove(_hx_param_key any) bool {
	key := ((Hx_Typedef_haxe__int64___int64)(Hx_Field_go_haxe_hxdynamic_toInt64(_hx_param_key)))
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

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_keys() any {
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

	var keys []Hx_Typedef_haxe__int64___int64
	_ = keys
	keys = make([]Hx_Typedef_haxe__int64___int64, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	var self []Hx_Typedef_haxe__int64___int64
	_ = self
	self = keys
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any](HxMakeArray(self...)))))
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_iterator() any {
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

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_keyValueIterator() any {
	return ((any)(Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance((func() *Hx_Obj_haxe_imap {
		_hx_icast_13 := this
		if _hx_icast_13 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_13, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_13.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())))
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_int64map
	_ = copied
	copied = Hx_Obj_haxe_ds_int64map_CreateInstance()
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
			var key Hx_Typedef_haxe__int64___int64
			_ = key
			key = ((Hx_Typedef_haxe__int64___int64)(Hx_Field_go_haxe_hxdynamic_toInt64(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))))
			var _hx_tmp_2 Hx_Typedef_haxe__int64___int64
			_ = _hx_tmp_2
			_hx_tmp_2 = key
			copied.VTable.Hx_Field_set(_hx_tmp_2, this.VTable.Hx_Field_get(key))
		}
	}

	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_14 := copied
		if _hx_icast_14 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_14, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_14.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_toString() string {
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

	var keys []Hx_Typedef_haxe__int64___int64
	_ = keys
	keys = make([]Hx_Typedef_haxe__int64___int64, _hx_tmp_0)
	for k := range this.Hx_Field_h {
		keys = append(keys, k)
	}

	var keys1 []Hx_Typedef_haxe__int64___int64
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
					var _hx_closure_src_59 func(x any)
					_ = _hx_closure_src_59
					_hx_closure_src_59 = s.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_59(((any)(x)))
					}
				}()(", ")
			}

			func() func(x Hx_Typedef_haxe__int64___int64) {
				var _hx_closure_src_60 func(x any)
				_ = _hx_closure_src_60
				_hx_closure_src_60 = s.VTable.Hx_Field_add
				return func(x Hx_Typedef_haxe__int64___int64) {
					_hx_closure_src_60(((any)(x)))
				}
			}()(keys1[i])
			func() func(x string) {
				var _hx_closure_src_61 func(x any)
				_ = _hx_closure_src_61
				_hx_closure_src_61 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_61(((any)(x)))
				}
			}()(" => ")
			s.VTable.Hx_Field_add(values1[i])
		}
	}

	s.VTable.Hx_Field_addChar(93)
	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_clear() {
	clear(this.Hx_Field_h)
}

func (this *Hx_Obj_haxe_ds_int64map) Hx_Field_size() int32 {
	return ((int32)(len(this.Hx_Field_h)))
}
