package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_int64map() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_int64map() {
	Hx_Field_loreline_int64map_INITIAL_SIZE = Hx_Init_Hx_Field_loreline_int64map_INITIAL_SIZE()
	Hx_Field_loreline_int64map_LOAD_FACTOR = Hx_Init_Hx_Field_loreline_int64map_LOAD_FACTOR()
}

type Hx_Obj_VTable_loreline_int64map interface {
	Hx_Field_length() int32
	Hx_Field_resize(newCapacity int32)
	Hx_Field_hashCoords(high int32, low int32) int32
	Hx_Field_clear()
	Hx_Field__get(high int32, low int32) any
	Hx_Field__exists(high int32, low int32) bool
	Hx_Field__remove(high int32, low int32) bool
	Hx_Field_iterator() *Hx_Obj_loreline__int64map_int64mapiterator
	Hx_Field__set(high int32, low int32, value any)
	Hx_Field_keyIterator() *Hx_Obj_loreline__int64map_int64mapkeyiterator
	Hx_Field_keyValueIterator() *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator
	Hx_Field_copy() *Hx_Obj_loreline_int64map
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_int64map struct {
	VTable           Hx_Obj_VTable_loreline_int64map `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field__keys1  HxArray[int32]
	Hx_Field__keys2  HxArray[int32]
	Hx_Field__values HxArray[any]
	Hx_Field_size    int32
	Hx_Field_mask    int32
}

func Hx_Obj_loreline_int64map_CreateEmptyInstance() *Hx_Obj_loreline_int64map {
	obj := &Hx_Obj_loreline_int64map{}
	obj.VTable = obj
	obj.Hx_Field__keys1 = HxDefault[HxArray[int32]]()
	obj.Hx_Field__keys2 = HxDefault[HxArray[int32]]()
	obj.Hx_Field__values = HxDefault[HxArray[any]]()
	obj.Hx_Field_size = HxDefault[int32]()
	obj.Hx_Field_mask = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_int64map_CreateInstance() *Hx_Obj_loreline_int64map {
	obj := Hx_Obj_loreline_int64map_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_int64map) Hx_New() {
	{
		this.Hx_Field__keys1 = nil
	}
	{
		this.Hx_Field__keys2 = nil
	}
	{
		this.Hx_Field__values = nil
	}
	{
		this.Hx_Field_size = 0
	}
	{
		this.Hx_Field_mask = 0
	}
	{
		var oldK1 HxArray[int32]
		_ = oldK1
		oldK1 = this.Hx_Field__keys1
		var oldK2 HxArray[int32]
		_ = oldK2
		oldK2 = this.Hx_Field__keys2
		var oldVals HxArray[any]
		_ = oldVals
		oldVals = this.Hx_Field__values
		var length HxNullable[int32]
		_ = length
		length = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (length).Valid != false {
			_hx_tmp_0 = (length).Value
		} else {
			_hx_tmp_0 = 0
		}

		var x []int32
		_ = x
		x = make([]int32, _hx_tmp_0)
		var self []int32
		_ = self
		self = x
		this.Hx_Field__keys1 = HxMakeArray(self...)
		var length1 HxNullable[int32]
		_ = length1
		length1 = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (length1).Valid != false {
			_hx_tmp_1 = (length1).Value
		} else {
			_hx_tmp_1 = 0
		}

		var x1 []int32
		_ = x1
		x1 = make([]int32, _hx_tmp_1)
		var self1 []int32
		_ = self1
		self1 = x1
		this.Hx_Field__keys2 = HxMakeArray(self1...)
		var length2 HxNullable[int32]
		_ = length2
		length2 = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (length2).Valid != false {
			_hx_tmp_2 = (length2).Value
		} else {
			_hx_tmp_2 = 0
		}

		var x2 []any
		_ = x2
		x2 = make([]any, _hx_tmp_2)
		var self2 []any
		_ = self2
		self2 = x2
		this.Hx_Field__values = HxMakeArray(self2...)
		this.Hx_Field_mask = 15
		if oldK1 != nil {
			this.Hx_Field_size = 0
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = HxMakeArrayView[any](oldK1).Len()
				for _g < _g1 {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_3
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(oldVals.Get(i), nil))) {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK1).Get(i))))
						this.VTable.Hx_Field__set(_hx_tmp_4, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK2).Get(i)))), oldVals.Get(i))
					}
				}
			}
		}
	}

}

var Hx_Obj_loreline_int64map_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Int64Map",
	HxMakeArray[string]("INITIAL_SIZE", "LOAD_FACTOR"),
	HxMakeArray[string]("_keys1", "_keys2", "_values", "size", "mask", "resize", "hashCoords", "clear", "copy", "set", "_set", "get", "_get", "exists", "_exists", "remove", "_remove", "iterator", "keyIterator", "keyValueIterator", "length"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_int64map_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_int64map_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_int64map_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "INITIAL_SIZE":
			return &Hx_Field_loreline_int64map_INITIAL_SIZE
		case "LOAD_FACTOR":
			return &Hx_Field_loreline_int64map_LOAD_FACTOR
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_int64map_RTTI
}

func (this *Hx_Obj_loreline_int64map) Hx_As_Hx_Obj_loreline_int64map() *Hx_Obj_loreline_int64map {
	return this
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_resize(newCapacity int32) {
	var oldK1 HxArray[int32]
	_ = oldK1
	oldK1 = this.Hx_Field__keys1
	var oldK2 HxArray[int32]
	_ = oldK2
	oldK2 = this.Hx_Field__keys2
	var oldVals HxArray[any]
	_ = oldVals
	oldVals = this.Hx_Field__values
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{Value: newCapacity, Valid: true}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var x []int32
	_ = x
	x = make([]int32, _hx_tmp_0)
	var self []int32
	_ = self
	self = x
	this.Hx_Field__keys1 = HxMakeArray(self...)
	var length1 HxNullable[int32]
	_ = length1
	length1 = HxNullable[int32]{Value: newCapacity, Valid: true}

	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (length1).Valid != false {
		_hx_tmp_1 = (length1).Value
	} else {
		_hx_tmp_1 = 0
	}

	var x1 []int32
	_ = x1
	x1 = make([]int32, _hx_tmp_1)
	var self1 []int32
	_ = self1
	self1 = x1
	this.Hx_Field__keys2 = HxMakeArray(self1...)
	var length2 HxNullable[int32]
	_ = length2
	length2 = HxNullable[int32]{Value: newCapacity, Valid: true}

	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	if (length2).Valid != false {
		_hx_tmp_2 = (length2).Value
	} else {
		_hx_tmp_2 = 0
	}

	var x2 []any
	_ = x2
	x2 = make([]any, _hx_tmp_2)
	var self2 []any
	_ = self2
	self2 = x2
	this.Hx_Field__values = HxMakeArray(self2...)
	this.Hx_Field_mask = (newCapacity - 1)
	if oldK1 != nil {
		this.Hx_Field_size = 0
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = HxMakeArrayView[any](oldK1).Len()
			for _g < _g1 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_3
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(oldVals.Get(i), nil))) {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK1).Get(i))))
					this.VTable.Hx_Field__set(_hx_tmp_4, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK2).Get(i)))), oldVals.Get(i))
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_hashCoords(high int32, low int32) int32 {
	var h int32
	_ = h
	h = (high + (low << 16))
	h = ((h ^ ((int32)((((uint32)(h)) >> 16)))) * -2048144789)
	h = ((h ^ ((int32)((((uint32)(h)) >> 13)))) * -1028477387)
	return (h ^ ((int32)((((uint32)(h)) >> 16))))
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_clear() {
	this.Hx_Field__keys1 = nil
	this.Hx_Field__keys2 = nil
	this.Hx_Field__values = nil
	{
		var oldK1 HxArray[int32]
		_ = oldK1
		oldK1 = this.Hx_Field__keys1
		var oldK2 HxArray[int32]
		_ = oldK2
		oldK2 = this.Hx_Field__keys2
		var oldVals HxArray[any]
		_ = oldVals
		oldVals = this.Hx_Field__values
		var length HxNullable[int32]
		_ = length
		length = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (length).Valid != false {
			_hx_tmp_0 = (length).Value
		} else {
			_hx_tmp_0 = 0
		}

		var x []int32
		_ = x
		x = make([]int32, _hx_tmp_0)
		var self []int32
		_ = self
		self = x
		this.Hx_Field__keys1 = HxMakeArray(self...)
		var length1 HxNullable[int32]
		_ = length1
		length1 = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (length1).Valid != false {
			_hx_tmp_1 = (length1).Value
		} else {
			_hx_tmp_1 = 0
		}

		var x1 []int32
		_ = x1
		x1 = make([]int32, _hx_tmp_1)
		var self1 []int32
		_ = self1
		self1 = x1
		this.Hx_Field__keys2 = HxMakeArray(self1...)
		var length2 HxNullable[int32]
		_ = length2
		length2 = HxNullable[int32]{Value: 16, Valid: true}

		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (length2).Valid != false {
			_hx_tmp_2 = (length2).Value
		} else {
			_hx_tmp_2 = 0
		}

		var x2 []any
		_ = x2
		x2 = make([]any, _hx_tmp_2)
		var self2 []any
		_ = self2
		self2 = x2
		this.Hx_Field__values = HxMakeArray(self2...)
		this.Hx_Field_mask = 15
		if oldK1 != nil {
			this.Hx_Field_size = 0
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = HxMakeArrayView[any](oldK1).Len()
				for _g < _g1 {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_3
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(oldVals.Get(i), nil))) {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK1).Get(i))))
						this.VTable.Hx_Field__set(_hx_tmp_4, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK2).Get(i)))), oldVals.Get(i))
					}
				}
			}
		}
	}

	this.Hx_Field_size = 0
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_copy() *Hx_Obj_loreline_int64map {
	var result *Hx_Obj_loreline_int64map
	_ = result
	result = Hx_Obj_loreline_int64map_CreateInstance()
	result.Hx_Field__keys1 = (Hx_Array_Copy[int32](HxMakeArrayView[int32](HxMakeArrayView[any](this.Hx_Field__keys1))))
	result.Hx_Field__keys2 = (Hx_Array_Copy[int32](HxMakeArrayView[int32](HxMakeArrayView[any](this.Hx_Field__keys2))))
	result.Hx_Field__values = (Hx_Array_Copy[any](HxMakeArrayView[any](this.Hx_Field__values)))
	result.Hx_Field_size = this.Hx_Field_size
	result.Hx_Field_mask = this.Hx_Field_mask
	return result
}

func (this *Hx_Obj_loreline_int64map) Hx_Field__set(high int32, low int32, value any) {
	if ((float64)(this.Hx_Field_size)) >= (((float64)((this.Hx_Field_mask + 1))) * float64(0.75)) {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = (this.Hx_Field_mask + 1)
		var newCapacity int32
		_ = newCapacity
		newCapacity = (_hx_tmp_0 << 1)
		var oldK1 HxArray[int32]
		_ = oldK1
		oldK1 = this.Hx_Field__keys1
		var oldK2 HxArray[int32]
		_ = oldK2
		oldK2 = this.Hx_Field__keys2
		var oldVals HxArray[any]
		_ = oldVals
		oldVals = this.Hx_Field__values
		var length HxNullable[int32]
		_ = length
		length = HxNullable[int32]{Value: newCapacity, Valid: true}

		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (length).Valid != false {
			_hx_tmp_1 = (length).Value
		} else {
			_hx_tmp_1 = 0
		}

		var x []int32
		_ = x
		x = make([]int32, _hx_tmp_1)
		var self []int32
		_ = self
		self = x
		this.Hx_Field__keys1 = HxMakeArray(self...)
		var length1 HxNullable[int32]
		_ = length1
		length1 = HxNullable[int32]{Value: newCapacity, Valid: true}

		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (length1).Valid != false {
			_hx_tmp_2 = (length1).Value
		} else {
			_hx_tmp_2 = 0
		}

		var x1 []int32
		_ = x1
		x1 = make([]int32, _hx_tmp_2)
		var self1 []int32
		_ = self1
		self1 = x1
		this.Hx_Field__keys2 = HxMakeArray(self1...)
		var length2 HxNullable[int32]
		_ = length2
		length2 = HxNullable[int32]{Value: newCapacity, Valid: true}

		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (length2).Valid != false {
			_hx_tmp_3 = (length2).Value
		} else {
			_hx_tmp_3 = 0
		}

		var x2 []any
		_ = x2
		x2 = make([]any, _hx_tmp_3)
		var self2 []any
		_ = self2
		self2 = x2
		this.Hx_Field__values = HxMakeArray(self2...)
		this.Hx_Field_mask = (newCapacity - 1)
		if oldK1 != nil {
			this.Hx_Field_size = 0
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = HxMakeArrayView[any](oldK1).Len()
				for _g < _g1 {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_4
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(oldVals.Get(i), nil))) {
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK1).Get(i))))
						this.VTable.Hx_Field__set(_hx_tmp_5, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](oldK2).Get(i)))), oldVals.Get(i))
					}
				}
			}
		}
	}

	var h int32
	_ = h
	h = (high + (low << 16))
	h = ((h ^ ((int32)((((uint32)(h)) >> 16)))) * -2048144789)
	h = ((h ^ ((int32)((((uint32)(h)) >> 13)))) * -1028477387)
	var hash int32
	_ = hash
	hash = (h ^ ((int32)((((uint32)(h)) >> 16))))
	var index int32
	_ = index
	index = (hash & this.Hx_Field_mask)
	for {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field__values.Get(index), nil))) {
			HxMakeArrayView[any](this.Hx_Field__keys1).Set(index, ((any)(high)))
			HxMakeArrayView[any](this.Hx_Field__keys2).Set(index, ((any)(low)))
			this.Hx_Field__values.Set(index, value)
			this.Hx_Field_size++
			return
		}

		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys1).Get(index), ((any)(high))))) {
			_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys2).Get(index), ((any)(low))))))
		} else {
			_hx_tmp_6 = false
		}

		if _hx_tmp_6 {
			this.Hx_Field__values.Set(index, value)
			return
		}

		index = ((index + 1) & this.Hx_Field_mask)
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field__get(high int32, low int32) any {
	var h int32
	_ = h
	h = (high + (low << 16))
	h = ((h ^ ((int32)((((uint32)(h)) >> 16)))) * -2048144789)
	h = ((h ^ ((int32)((((uint32)(h)) >> 13)))) * -1028477387)
	var hash int32
	_ = hash
	hash = (h ^ ((int32)((((uint32)(h)) >> 16))))
	var index int32
	_ = index
	index = (hash & this.Hx_Field_mask)
	for {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field__values.Get(index), nil))) {
			return nil
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys1).Get(index), ((any)(high))))) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys2).Get(index), ((any)(low))))))
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			return this.Hx_Field__values.Get(index)
		}

		index = ((index + 1) & this.Hx_Field_mask)
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field__exists(high int32, low int32) bool {
	var h int32
	_ = h
	h = (high + (low << 16))
	h = ((h ^ ((int32)((((uint32)(h)) >> 16)))) * -2048144789)
	h = ((h ^ ((int32)((((uint32)(h)) >> 13)))) * -1028477387)
	var hash int32
	_ = hash
	hash = (h ^ ((int32)((((uint32)(h)) >> 16))))
	var index int32
	_ = index
	index = (hash & this.Hx_Field_mask)
	for {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field__values.Get(index), nil))) {
			return false
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys1).Get(index), ((any)(high))))) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys2).Get(index), ((any)(low))))))
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			return true
		}

		index = ((index + 1) & this.Hx_Field_mask)
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field__remove(high int32, low int32) bool {
	var h int32
	_ = h
	h = (high + (low << 16))
	h = ((h ^ ((int32)((((uint32)(h)) >> 16)))) * -2048144789)
	h = ((h ^ ((int32)((((uint32)(h)) >> 13)))) * -1028477387)
	var hash int32
	_ = hash
	hash = (h ^ ((int32)((((uint32)(h)) >> 16))))
	var index int32
	_ = index
	index = (hash & this.Hx_Field_mask)
	for {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field__values.Get(index), nil))) {
			return false
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys1).Get(index), ((any)(high))))) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](this.Hx_Field__keys2).Get(index), ((any)(low))))))
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			{
				var val any
				_ = val
				this.Hx_Field__values.Set(index, val)
			}

			this.Hx_Field_size--
			index = ((index + 1) & this.Hx_Field_mask)
			for {
				if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field__values.Get(index), nil)))) {
					break
				}

				var k1 int32
				_ = k1
				k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field__keys1).Get(index))))
				var k2 int32
				_ = k2
				k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field__keys2).Get(index))))
				var v any
				_ = v
				v = this.Hx_Field__values.Get(index)
				{
					var val any
					_ = val
					this.Hx_Field__values.Set(index, val)
				}

				this.Hx_Field_size--
				this.VTable.Hx_Field__set(k1, k2, v)
				index = ((index + 1) & this.Hx_Field_mask)
			}

			return true
		}

		index = ((index + 1) & this.Hx_Field_mask)
	}
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_iterator() *Hx_Obj_loreline__int64map_int64mapiterator {
	return Hx_Obj_loreline__int64map_int64mapiterator_CreateInstance(this)
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_keyIterator() *Hx_Obj_loreline__int64map_int64mapkeyiterator {
	return Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateInstance(this)
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_keyValueIterator() *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator {
	return Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateInstance(this)
}

func (this *Hx_Obj_loreline_int64map) Hx_Field_length() int32 {
	return this.Hx_Field_size
}

var Hx_Field_loreline_int64map_INITIAL_SIZE int32

func Hx_Init_Hx_Field_loreline_int64map_INITIAL_SIZE() int32 {
	return 16
}

var Hx_Field_loreline_int64map_LOAD_FACTOR float64

func Hx_Init_Hx_Field_loreline_int64map_LOAD_FACTOR() float64 {
	return float64(0.75)
}

type Hx_Typedef_loreline_int64mapkeyval = any

type Hx_Typedef_loreline_int64mapkey = any

func Hx_ClassInit_Hx_Obj_VTable_loreline__int64map_int64mapiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__int64map_int64mapiterator() {
}

type Hx_Obj_VTable_loreline__int64map_int64mapiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__int64map_int64mapiterator struct {
	VTable         Hx_Obj_VTable_loreline__int64map_int64mapiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__int64map_int64mapiterator_CreateEmptyInstance() *Hx_Obj_loreline__int64map_int64mapiterator {
	obj := &Hx_Obj_loreline__int64map_int64mapiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__int64map_int64mapiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__int64map_int64mapiterator {
	obj := Hx_Obj_loreline__int64map_int64mapiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__int64map_int64mapiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Int64Map.Int64MapIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__int64map_int64mapiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__int64map_int64mapiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__int64map_int64mapiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__int64map_int64mapiterator_RTTI
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_As_Hx_Obj_loreline__int64map_int64mapiterator() *Hx_Obj_loreline__int64map_int64mapiterator {
	return this
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__int64map_int64mapiterator) Hx_Field_next() any {
	var v any
	_ = v
	v = this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index)
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return v
}

func Hx_ClassInit_Hx_Obj_VTable_loreline__int64map_int64mapkeyiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__int64map_int64mapkeyiterator() {
}

type Hx_Obj_VTable_loreline__int64map_int64mapkeyiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__int64map_int64mapkeyiterator struct {
	VTable         Hx_Obj_VTable_loreline__int64map_int64mapkeyiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateEmptyInstance() *Hx_Obj_loreline__int64map_int64mapkeyiterator {
	obj := &Hx_Obj_loreline__int64map_int64mapkeyiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__int64map_int64mapkeyiterator {
	obj := Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__int64map_int64mapkeyiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Int64Map.Int64MapKeyIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__int64map_int64mapkeyiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__int64map_int64mapkeyiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__int64map_int64mapkeyiterator_RTTI
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_As_Hx_Obj_loreline__int64map_int64mapkeyiterator() *Hx_Obj_loreline__int64map_int64mapkeyiterator {
	return this
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyiterator) Hx_Field_next() any {
	var k1 int32
	_ = k1
	k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys1).Get(this.Hx_Field_index))))
	var k2 int32
	_ = k2
	k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys2).Get(this.Hx_Field_index))))
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return any(map[string]any{"high": ((any)(k1)), "low": ((any)(k2))})
}

func Hx_ClassInit_Hx_Obj_VTable_loreline__int64map_int64mapkeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__int64map_int64mapkeyvalueiterator() {
}

type Hx_Obj_VTable_loreline__int64map_int64mapkeyvalueiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__int64map_int64mapkeyvalueiterator struct {
	VTable         Hx_Obj_VTable_loreline__int64map_int64mapkeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator {
	obj := &Hx_Obj_loreline__int64map_int64mapkeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator {
	obj := Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Int64Map.Int64MapKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__int64map_int64mapkeyvalueiterator_RTTI
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_As_Hx_Obj_loreline__int64map_int64mapkeyvalueiterator() *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator {
	return this
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__int64map_int64mapkeyvalueiterator) Hx_Field_next() any {
	var k1 int32
	_ = k1
	k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys1).Get(this.Hx_Field_index))))
	var k2 int32
	_ = k2
	k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys2).Get(this.Hx_Field_index))))
	var v any
	_ = v
	v = this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index)
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return any(map[string]any{"key": any(map[string]any{"high": ((any)(k1)), "low": ((any)(k2))}), "value": v})
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[int32]()
}
