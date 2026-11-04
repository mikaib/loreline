package core

func Hx_ClassInit_Hx_Obj_VTable_lambda() {
}

func Hx_StaticInit_Hx_Obj_VTable_lambda() {
}

type Hx_Obj_VTable_lambda interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_lambda struct {
	VTable Hx_Obj_VTable_lambda `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_lambda_CreateEmptyInstance() *Hx_Obj_lambda {
	obj := &Hx_Obj_lambda{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_lambda_CreateInstance() *Hx_Obj_lambda {
	obj := Hx_Obj_lambda_CreateEmptyInstance()
	return obj
}

var Hx_Obj_lambda_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Lambda",
	HxMakeArray[string]("array", "list", "map", "mapi", "flatten", "flatMap", "has", "exists", "foreach", "iter", "filter", "fold", "foldi", "count", "empty", "indexOf", "find", "findIndex", "concat"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_lambda_CreateInstance()
	},
	func() any {
		return Hx_Obj_lambda_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_lambda_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "array":
			return Hx_Field_lambda_array
		case "list":
			return Hx_Field_lambda_list
		case "map":
			return Hx_Field_lambda_map
		case "mapi":
			return Hx_Field_lambda_mapi
		case "flatten":
			return Hx_Field_lambda_flatten
		case "flatMap":
			return Hx_Field_lambda_flatMap
		case "has":
			return Hx_Field_lambda_has
		case "exists":
			return Hx_Field_lambda_exists
		case "foreach":
			return Hx_Field_lambda_foreach
		case "iter":
			return Hx_Field_lambda_iter
		case "filter":
			return Hx_Field_lambda_filter
		case "fold":
			return Hx_Field_lambda_fold
		case "foldi":
			return Hx_Field_lambda_foldi
		case "count":
			return Hx_Field_lambda_count
		case "empty":
			return Hx_Field_lambda_empty
		case "indexOf":
			return Hx_Field_lambda_indexOf
		case "find":
			return Hx_Field_lambda_find
		case "findIndex":
			return Hx_Field_lambda_findIndex
		case "concat":
			return Hx_Field_lambda_concat
		}
		return nil
	}
}

func (this *Hx_Obj_lambda) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_lambda_RTTI
}

func (this *Hx_Obj_lambda) Hx_As_Hx_Obj_lambda() *Hx_Obj_lambda {
	return this
}

func Hx_Field_lambda_array(it any) HxArray[any] {
	var a HxArray[any]
	_ = a
	a = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gi any
		_ = _gi
		_gi = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gi, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gi, "next")
			var i any
			_ = i
			i = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](a), i))
		}
	}

	return a
}

func Hx_Field_lambda_list(it any) *Hx_Obj_haxe_ds_list {
	var l *Hx_Obj_haxe_ds_list
	_ = l
	l = Hx_Obj_haxe_ds_list_CreateInstance()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gi any
		_ = _gi
		_gi = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gi, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gi, "next")
			var i any
			_ = i
			i = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			l.VTable.Hx_Field_add(i)
		}
	}

	return l
}

func Hx_Field_lambda_map(it any, f func(item any) any) HxArray[any] {
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_3, f(x)))
		}
	}

	return _g
}

func Hx_Field_lambda_mapi(it any, f func(index int32, item any) any) HxArray[any] {
	var i int32
	_ = i
	i = 0
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = HxMakeArrayView[any](_g)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = i
			i = (i + 1)
			_ = (Hx_Array_Push[any](_hx_tmp_3, f(_hx_tmp_4, x)))
		}
	}

	return _g
}

func Hx_Field_lambda_flatten(it any) HxArray[any] {
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _ge any
		_ = _ge
		_ge = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "next")
			var e any
			_ = e
			e = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			{
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(e, "iterator")
				var _gx any
				_ = _gx
				_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn())
				for {
					var _hx_tmp_4 any
					_ = _hx_tmp_4
					_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_5 any
					_ = _hx_tmp_5
					_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
					var x any
					_ = x
					x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn())
					_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), x))
				}
			}
		}
	}

	return _g
}

func Hx_Field_lambda_flatMap(it any, f func(item any) any) HxArray[any] {
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_3, f(x)))
		}
	}

	var _g1 HxArray[any]
	_ = _g1
	_g1 = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(((any)(_g)), "iterator")
		var _ge any
		_ = _ge
		_ge = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "next")
			var e any
			_ = e
			e = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any]().Dyn())
			{
				var _hx_tmp_7 any
				_ = _hx_tmp_7
				_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(e, "iterator")
				var _gx any
				_ = _gx
				_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_7, HxMakeArray[any]().Dyn())
				for {
					var _hx_tmp_8 any
					_ = _hx_tmp_8
					_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_8, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_9 any
					_ = _hx_tmp_9
					_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
					var x any
					_ = x
					x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any]().Dyn())
					_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g1), x))
				}
			}
		}
	}

	return _g1
}

func Hx_Field_lambda_has(it any, elt any) bool {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(x, elt))) {
				return true
			}
		}
	}

	return false
}

func Hx_Field_lambda_exists(it any, f func(item any) bool) bool {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if f(x) {
				return true
			}
		}
	}

	return false
}

func Hx_Field_lambda_foreach(it any, f func(item any) bool) bool {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if !f(x) {
				return false
			}
		}
	}

	return true
}

func Hx_Field_lambda_iter(it any, f func(item any)) {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
	var _gx any
	_ = _gx
	_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
	for {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
		if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
			break
		}

		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
		var x any
		_ = x
		x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
		f(x)
	}
}

func Hx_Field_lambda_filter(it any, f func(item any) bool) HxArray[any] {
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if f(x) {
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), x))
			}
		}
	}

	return _g
}

func Hx_Field_lambda_fold(it any, f func(item any, result any) any, first any) any {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			first = f(x, first)
		}
	}

	return first
}

func Hx_Field_lambda_foldi(it any, f func(item any, result any, index int32) any, first any) any {
	var i int32
	_ = i
	i = 0
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			first = f(x, first, i)
			i++
		}
	}

	return first
}

func Hx_Field_lambda_count(it any, pred__isOptional any) int32 {
	pred := HxDefault[func(item any) bool]()
	_ = pred

	if pred__isOptional != nil {
		pred = HxConvert[func(item any) bool](pred__isOptional)
	} else {
		pred = nil
	}
	{
		var n int32
		_ = n
		n = 0
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(pred)), ((any)(nil))))) {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
			var _g_ any
			_ = _g_
			_g_ = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
			for {
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_, "hasNext")
				if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
					break
				}

				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_g_, "next")
				var _hx_reserved__ any
				_ = _hx_reserved__
				_hx_reserved__ = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
				n++
			}
		} else {
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
			var _gx any
			_ = _gx
			_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn())
			for {
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
				if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()))) {
					break
				}

				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
				var x any
				_ = x
				x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn())
				if pred(x) {
					n++
				}
			}
		}

		return n
	}
}

func Hx_Field_lambda_empty(it any) bool {
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "hasNext")
	return !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))))
}

func Hx_Field_lambda_indexOf(it any, v any) int32 {
	var i int32
	_ = i
	i = 0
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gv any
		_ = _gv
		_gv = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "next")
			var v2 any
			_ = v2
			v2 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, v2))) {
				return i
			}

			i++
		}
	}

	return -1
}

func Hx_Field_lambda_find(it any, f func(item any) bool) any {
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gv any
		_ = _gv
		_gv = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "next")
			var v any
			_ = v
			v = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if f(v) {
				return v
			}
		}
	}

	return nil
}

func Hx_Field_lambda_findIndex(it any, f func(item any) bool) int32 {
	var i int32
	_ = i
	i = 0
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(it, "iterator")
		var _gv any
		_ = _gv
		_gv = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gv, "next")
			var v any
			_ = v
			v = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			if f(v) {
				return i
			}

			i++
		}
	}

	return -1
}

func Hx_Field_lambda_concat(a any, b any) HxArray[any] {
	var l HxArray[any]
	_ = l
	l = HxMakeArray[any]().Dyn()
	{
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(a, "iterator")
		var _gx any
		_ = _gx
		_gx = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gx, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](l), x))
		}
	}

	{
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(b, "iterator")
		var _gx1 any
		_ = _gx1
		_gx1 = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn())
		for {
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gx1, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(_gx1, "next")
			var x any
			_ = x
			x = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn())
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](l), x))
		}
	}

	return l
}

func init() {
	hxRegisterArrayConstructor[any]()
}
