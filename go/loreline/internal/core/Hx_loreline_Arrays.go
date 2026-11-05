package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_arrays() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_arrays() {
}

type Hx_Obj_VTable_loreline_arrays interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_arrays struct {
	VTable Hx_Obj_VTable_loreline_arrays `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_arrays_CreateEmptyInstance() *Hx_Obj_loreline_arrays {
	obj := &Hx_Obj_loreline_arrays{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_arrays_CreateInstance() *Hx_Obj_loreline_arrays {
	obj := Hx_Obj_loreline_arrays_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_arrays_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Arrays",
	HxMakeArray[string]("isArray", "arrayLength", "arrayGet", "arraySet", "createArray", "arrayPush", "arrayPop", "arrayShift", "arrayInsert", "arrayRemoveAt", "getIterator", "arrayCopy", "arraySort", "arrayReverse", "arrayJoin"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_arrays_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_arrays_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_arrays_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "isArray":
			return Hx_Field_loreline_arrays_isArray
		case "arrayLength":
			return Hx_Field_loreline_arrays_arrayLength
		case "arrayGet":
			return Hx_Field_loreline_arrays_arrayGet
		case "arraySet":
			return Hx_Field_loreline_arrays_arraySet
		case "createArray":
			return Hx_Field_loreline_arrays_createArray
		case "arrayPush":
			return Hx_Field_loreline_arrays_arrayPush
		case "arrayPop":
			return Hx_Field_loreline_arrays_arrayPop
		case "arrayShift":
			return Hx_Field_loreline_arrays_arrayShift
		case "arrayInsert":
			return Hx_Field_loreline_arrays_arrayInsert
		case "arrayRemoveAt":
			return Hx_Field_loreline_arrays_arrayRemoveAt
		case "getIterator":
			return Hx_Field_loreline_arrays_getIterator
		case "arrayCopy":
			return Hx_Field_loreline_arrays_arrayCopy
		case "arraySort":
			return Hx_Field_loreline_arrays_arraySort
		case "arrayReverse":
			return Hx_Field_loreline_arrays_arrayReverse
		case "arrayJoin":
			return Hx_Field_loreline_arrays_arrayJoin
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_arrays) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_arrays_RTTI
}

func (this *Hx_Obj_loreline_arrays) Hx_As_Hx_Obj_loreline_arrays() *Hx_Obj_loreline_arrays {
	return this
}

func Hx_Field_loreline_arrays_isArray(array any) bool {
	if Hx_Field_std_isOfType(array, Hx_Obj_array_RTTI) {
		return true
	}

	return false
}

func Hx_Field_loreline_arrays_arrayLength(array any) int32 {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	return Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)
}

func Hx_Field_loreline_arrays_arrayGet(array any, index int32) any {
	var i int32
	_ = i
	i = Hx_Field_std_int(((float64)(index)))
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if i >= 0 {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = i
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr))
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		return arr.Get(i)
	} else {
		return nil
	}
}

func Hx_Field_loreline_arrays_arraySet(array any, index int32, value any) {
	var i int32
	_ = i
	i = Hx_Field_std_int(((float64)(index)))
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	arr.Set(i, value)
}

func Hx_Field_loreline_arrays_createArray() any {
	var arr HxArray[any]
	_ = arr
	arr = HxMakeArray[any]().Dyn()
	return ((any)(arr))
}

func Hx_Field_loreline_arrays_arrayPush(array any, value any) {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	_ = (Hx_Array_Push[any](arr, value))
}

func Hx_Field_loreline_arrays_arrayPop(array any) any {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	return (Hx_Array_Pop[any](arr).Value)
}

func Hx_Field_loreline_arrays_arrayShift(array any) any {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	return (Hx_Array_Shift[any](arr).Value)
}

func Hx_Field_loreline_arrays_arrayInsert(array any, index int32, value any) {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	(Hx_Array_Insert[any](arr, index, value))
}

func Hx_Field_loreline_arrays_arrayRemoveAt(array any, index int32) {
	var arr HxArray[any]
	_ = arr
	arr = HxAnyToArray(array)
	_ = (Hx_Array_Splice[any](arr, index, 1))
}

func Hx_Field_loreline_arrays_getIterator(array any) any {
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxAnyToArray(array))))
}

func Hx_Field_loreline_arrays_arrayCopy(array any) any {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
	var _hx_reserved_copy any
	_ = _hx_reserved_copy
	_hx_reserved_copy = Hx_Field_loreline_arrays_createArray()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = _hx_reserved_copy
			Hx_Field_loreline_arrays_arrayPush(_hx_tmp_1, Hx_Field_loreline_arrays_arrayGet(array, i))
		}
	}

	return _hx_reserved_copy
}

func Hx_Field_loreline_arrays_arraySort(array any, cmp func(_hx_carg0 any, _hx_carg1 any) int32) {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
	{
		var _g int32
		_ = _g
		_g = 1
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var key any
			_ = key
			key = Hx_Field_loreline_arrays_arrayGet(array, i)
			var j int32
			_ = j
			j = (i - 1)
			for {
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if j >= 0 {
					_hx_tmp_1 = (cmp(Hx_Field_loreline_arrays_arrayGet(array, j), key) > 0)
				} else {
					_hx_tmp_1 = false
				}

				if !(_hx_tmp_1) {
					break
				}

				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = array
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = (j + 1)
				Hx_Field_loreline_arrays_arraySet(_hx_tmp_2, _hx_tmp_3, Hx_Field_loreline_arrays_arrayGet(array, j))
				j--
			}

			Hx_Field_loreline_arrays_arraySet(array, (j + 1), key)
		}
	}
}

func Hx_Field_loreline_arrays_arrayReverse(array any) {
	var i int32
	_ = i
	i = 0
	var j int32
	_ = j
	j = (Hx_Field_loreline_arrays_arrayLength(array) - 1)
	for i < j {
		var tmp any
		_ = tmp
		tmp = Hx_Field_loreline_arrays_arrayGet(array, i)
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = array
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = i
		Hx_Field_loreline_arrays_arraySet(_hx_tmp_0, _hx_tmp_1, Hx_Field_loreline_arrays_arrayGet(array, j))
		Hx_Field_loreline_arrays_arraySet(array, j, tmp)
		i++
		j--
	}
}

func Hx_Field_loreline_arrays_arrayJoin(array any, sep string) string {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = _hx_reserved_len
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_192 func(x any)
					_ = _hx_closure_src_192
					_hx_closure_src_192 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_192(((any)(x)))
					}
				}()(sep)
			}

			var _hx_tmp_1 string
			_ = _hx_tmp_1
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_loreline_arrays_arrayGet(array, i), nil))) {
				_hx_tmp_1 = "null"
			} else {
				_hx_tmp_1 = Hx_Field_std_string(Hx_Field_loreline_arrays_arrayGet(array, i))
			}

			func() func(x string) {
				var _hx_closure_src_193 func(x any)
				_ = _hx_closure_src_193
				_hx_closure_src_193 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_193(((any)(x)))
				}
			}()(_hx_tmp_1)
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func init() {
	hxRegisterArrayConstructor[any]()
}
