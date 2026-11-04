package core

import "reflect"

func Hx_ClassInit_Hx_Obj_VTable_reflect() {
}

func Hx_StaticInit_Hx_Obj_VTable_reflect() {
}

type Hx_Obj_VTable_reflect interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_reflect struct {
	VTable Hx_Obj_VTable_reflect `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_reflect_CreateEmptyInstance() *Hx_Obj_reflect {
	obj := &Hx_Obj_reflect{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_reflect_CreateInstance() *Hx_Obj_reflect {
	obj := Hx_Obj_reflect_CreateEmptyInstance()
	return obj
}

var Hx_Obj_reflect_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Reflect",
	HxMakeArray[string]("hasField", "field", "setField", "getProperty", "setProperty", "callMethod", "fields", "isFunction", "compare", "compareMethods", "isObject", "isEnumValue", "deleteField", "copy", "makeVarArgs"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_reflect_CreateInstance()
	},
	func() any {
		return Hx_Obj_reflect_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_reflect_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "hasField":
			return Hx_Field_reflect_hasField
		case "field":
			return Hx_Field_reflect_field
		case "setField":
			return Hx_Field_reflect_setField
		case "getProperty":
			return Hx_Field_reflect_getProperty
		case "setProperty":
			return Hx_Field_reflect_setProperty
		case "callMethod":
			return Hx_Field_reflect_callMethod
		case "fields":
			return Hx_Field_reflect_fields
		case "isFunction":
			return Hx_Field_reflect_isFunction
		case "compare":
			return Hx_Field_reflect_compare
		case "compareMethods":
			return Hx_Field_reflect_compareMethods
		case "isObject":
			return Hx_Field_reflect_isObject
		case "isEnumValue":
			return Hx_Field_reflect_isEnumValue
		case "deleteField":
			return Hx_Field_reflect_deleteField
		case "copy":
			return Hx_Field_reflect_copy
		case "makeVarArgs":
			return Hx_Field_reflect_makeVarArgs
		}
		return nil
	}
}

func (this *Hx_Obj_reflect) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_reflect_RTTI
}

func (this *Hx_Obj_reflect) Hx_As_Hx_Obj_reflect() *Hx_Obj_reflect {
	return this
}

func Hx_Field_reflect_hasField(o any, field string) bool {
	var _hx_tmp_0 HxArray[string]
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_reflect_fields(o)
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = field
	return ((Hx_Array_IndexOf[string](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) != -1)
}

func Hx_Field_reflect_field(o any, field string) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(o, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (field == HxStringNull)
	}

	if _hx_tmp_0 {
		return nil
	}

	return Hx_Field_go_haxe_hxdynamic_getField(o, field)
}

func Hx_Field_reflect_setField(o any, field string, value any) {
	Hx_Field_go_haxe_hxdynamic_setField(o, field, value)
}

func Hx_Field_reflect_getProperty(o any, field string) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(o, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (field == HxStringNull)
	}

	if _hx_tmp_0 {
		return nil
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = o
	var fn any
	_ = fn
	fn = Hx_Field_go_haxe_hxdynamic_getField(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_stringConcat("get_", field))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(fn, nil))) {
		return Hx_Field_go_haxe_hxdynamic_getField(o, field)
	}

	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = fn
	return Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn())
}

func Hx_Field_reflect_setProperty(o any, field string, value any) {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = o
	var fn any
	_ = fn
	fn = Hx_Field_go_haxe_hxdynamic_getField(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_stringConcat("set_", field))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(fn, nil))) {
		Hx_Field_go_haxe_hxdynamic_setField(o, field, value)
		return
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = fn
	Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any](value).Dyn())
}

func Hx_Field_reflect_callMethod(o any, _hx_reserved_func any, args HxArray[any]) any {
	return Hx_Field_go_haxe_hxdynamic_call(_hx_reserved_func, args)
}

func Hx_Field_reflect_fields(o any) HxArray[string] {
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(o)
	if !value.IsValid() {
		return HxMakeArray[string]()
	}

	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		if value.IsNil() {
			return HxMakeArray[string]()
		}

		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 == reflect.Map {
		var keys []reflect.Value
		_ = keys
		keys = value.MapKeys()
		var self []reflect.Value
		_ = self
		self = keys
		var _g HxArray[string]
		_ = _g
		_g = HxMakeArray[string]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[reflect.Value]
			_ = _g2
			_g2 = HxMakeArray(self...)
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g1
				if !(_hx_tmp_5 < _g2.Len()) {
					break
				}

				var v reflect.Value
				_ = v
				v = _g2.Get(_g1)
				_g1++
				var _hx_tmp_6 HxArray[string]
				_ = _hx_tmp_6
				_hx_tmp_6 = _g
				_ = (Hx_Array_Push[string](_hx_tmp_6, Hx_Field_std_string(v)))
			}
		}

		var names HxArray[string]
		_ = names
		names = _g
		var _hx_tmp_7 HxArray[string]
		_ = _hx_tmp_7
		_hx_tmp_7 = names
		(Hx_Array_Sort[string](_hx_tmp_7, func() func(a string, b string) int32 {
			var _hx_closure_src_49 func(a any, b any) int32
			_ = _hx_closure_src_49
			_hx_closure_src_49 = Hx_Field_reflect_compare
			return func(a string, b string) int32 {
				return _hx_closure_src_49(((any)(a)), ((any)(b)))
			}
		}()))
		return names
	}

	var _hx_tmp_8 reflect.Kind
	_ = _hx_tmp_8
	_hx_tmp_8 = kind
	if _hx_tmp_8 == reflect.Struct {
		var t Hx_Enum_valuetype
		_ = t
		t = Hx_Field_type_typeof(o)
		if t.Hx_Field_enumIndex() == 7 {
			var c *Hx_Obj_go_haxe_hxclass
			_ = c
			c = t.(Hx_Enum_valuetype_TClass).Hx_Field_c
			var _g HxArray[string]
			_ = _g
			_g = HxMakeArray[string]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[string]
				_ = _g2
				_g2 = Hx_Field_type_getInstanceFields(c)
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g1
					if !(_hx_tmp_9 < _g2.Len()) {
						break
					}

					var v string
					_ = v
					v = _g2.Get(_g1)
					_g1++
					if value.FieldByName(Hx_Field_go_haxe_hxdynamic_formatField(v)).IsValid() {
						_ = (Hx_Array_Push[string](_g, v))
					}
				}
			}

			return _g
		} else {
			return HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_getField(o, "staticFields"))
		}
	}

	return HxMakeArray[string]()
}

func Hx_Field_reflect_isFunction(f any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f, nil))) {
		return false
	}

	var _hx_tmp_0 reflect.Kind
	_ = _hx_tmp_0
	_hx_tmp_0 = reflect.TypeOf(f).Kind()
	return (_hx_tmp_0 == reflect.Func)
}

func Hx_Field_reflect_compare(a any, b any) int32 {
	var aIsNull bool
	_ = aIsNull
	aIsNull = Hx_Field_go_haxe_hxstringnull_dynIsStringNull(a)
	var bIsNull bool
	_ = bIsNull
	bIsNull = Hx_Field_go_haxe_hxstringnull_dynIsStringNull(b)
	if aIsNull || bIsNull {
		if aIsNull && bIsNull {
			return 0
		} else {
			return -1
		}
	}

	{
		var hx_try_return int32
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if _hx_reserved__, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_exception").(*Hx_Obj_haxe_exception); _hx_ok {
						_ = _hx_reserved__
						{
							hx_try_state = 1
							hx_try_return = -1
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, b))) {
					hx_try_state = 1
					hx_try_return = 0
					return
				} else {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_gt(a, b))) {
						hx_try_state = 1
						hx_try_return = 1
						return
					} else {
						hx_try_state = 1
						hx_try_return = -1
						return
					}
				}
			}
		}()

		return hx_try_return
	}
}

func Hx_Field_reflect_compareMethods(f1 any, f2 any) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f1, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f2, nil))))
	}

	if _hx_tmp_0 {
		return false
	}

	return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(f1, f2))))
}

func Hx_Field_reflect_isObject(v any) bool {
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(v)
	if !value.IsValid() {
		return false
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(v, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_reflect_isEnumValue(value)
	}

	if _hx_tmp_0 {
		return false
	}

	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 reflect.Kind
	_ = _hx_tmp_2
	_hx_tmp_2 = kind
	if _hx_tmp_2 == reflect.Ptr {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = kind
		var _hx_tmp_4 reflect.Kind
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_3
		_hx_tmp_1 = (_hx_tmp_4 == reflect.Interface)
	}

	if _hx_tmp_1 {
		if value.IsNil() {
			return false
		}

		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	var _hx_tmp_6 reflect.Kind
	_ = _hx_tmp_6
	_hx_tmp_6 = kind
	if _hx_tmp_6 == reflect.String {
		_hx_tmp_5 = Hx_Field_go_haxe_hxstringnull_isNull(value.String())
	} else {
		_hx_tmp_5 = false
	}

	if _hx_tmp_5 {
		return false
	}

	var _hx_reserved_type reflect.Type
	_ = _hx_reserved_type
	_hx_reserved_type = value.Type()
	var _hx_tmp_7 bool
	_ = _hx_tmp_7
	var _hx_tmp_8 bool
	_ = _hx_tmp_8
	var _hx_tmp_9 reflect.Kind
	_ = _hx_tmp_9
	_hx_tmp_9 = kind
	if _hx_tmp_9 == reflect.String {
		_hx_tmp_8 = true
	} else {
		var _hx_tmp_10 reflect.Kind
		_ = _hx_tmp_10
		_hx_tmp_10 = kind
		var _hx_tmp_11 reflect.Kind
		_ = _hx_tmp_11
		_hx_tmp_11 = _hx_tmp_10
		_hx_tmp_8 = (_hx_tmp_11 == reflect.Struct)
	}

	if _hx_tmp_8 {
		_hx_tmp_7 = true
	} else {
		var _hx_tmp_12 bool
		_ = _hx_tmp_12
		var _hx_tmp_13 bool
		_ = _hx_tmp_13
		var _hx_tmp_14 reflect.Kind
		_ = _hx_tmp_14
		_hx_tmp_14 = kind
		if _hx_tmp_14 == reflect.Map {
			var _hx_tmp_15 reflect.Kind
			_ = _hx_tmp_15
			_hx_tmp_15 = _hx_reserved_type.Key().Kind()
			var _hx_tmp_16 reflect.Kind
			_ = _hx_tmp_16
			_hx_tmp_16 = _hx_tmp_15
			_hx_tmp_13 = (_hx_tmp_16 == reflect.String)
		} else {
			_hx_tmp_13 = false
		}

		if _hx_tmp_13 {
			var _hx_tmp_17 reflect.Kind
			_ = _hx_tmp_17
			_hx_tmp_17 = _hx_reserved_type.Elem().Kind()
			var _hx_tmp_18 reflect.Kind
			_ = _hx_tmp_18
			_hx_tmp_18 = _hx_tmp_17
			_hx_tmp_12 = (_hx_tmp_18 == reflect.Interface)
		} else {
			_hx_tmp_12 = false
		}

		_hx_tmp_7 = _hx_tmp_12
	}

	return _hx_tmp_7
}

func Hx_Field_reflect_isEnumValue(v any) bool {
	return Hx_Field_go_haxe_hxdynamic_isHxEnumValue(v)
}

func Hx_Field_reflect_deleteField(o any, field string) bool {
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(o)
	if !value.IsValid() {
		return false
	}

	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		if value.IsNil() {
			return false
		}

		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 == reflect.Map {
		var kt reflect.Type
		_ = kt
		kt = value.Type().Key()
		var _hx_tmp_5 reflect.Kind
		_ = _hx_tmp_5
		_hx_tmp_5 = kt.Kind()
		if _hx_tmp_5 != reflect.String {
			return false
		}

		var kv reflect.Value
		_ = kv
		kv = reflect.ValueOf(field)
		if !value.MapIndex(kv).IsValid() {
			return false
		}

		var _hx_tmp_6 reflect.Value
		_ = _hx_tmp_6
		_hx_tmp_6 = kv
		value.SetMapIndex(_hx_tmp_6, reflect.ValueOf(nil))
		return true
	}

	return false
}

func Hx_Field_reflect_copy(o any) any {
	var value reflect.Value
	_ = value
	value = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(o)
	if !value.IsValid() {
		return nil
	}

	var kind reflect.Kind
	_ = kind
	kind = value.Kind()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 reflect.Kind
	_ = _hx_tmp_1
	_hx_tmp_1 = kind
	if _hx_tmp_1 == reflect.Ptr {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = kind
		var _hx_tmp_3 reflect.Kind
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = (_hx_tmp_3 == reflect.Interface)
	}

	if _hx_tmp_0 {
		if value.IsNil() {
			return nil
		}

		value = value.Elem()
		kind = value.Kind()
	}

	var _hx_tmp_4 reflect.Kind
	_ = _hx_tmp_4
	_hx_tmp_4 = kind
	if _hx_tmp_4 == reflect.Map {
		var mt reflect.Type
		_ = mt
		mt = value.Type()
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		var _hx_tmp_6 reflect.Kind
		_ = _hx_tmp_6
		_hx_tmp_6 = mt.Key().Kind()
		if _hx_tmp_6 != reflect.String {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_7 reflect.Kind
			_ = _hx_tmp_7
			_hx_tmp_7 = mt.Elem().Kind()
			var _hx_tmp_8 reflect.Kind
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_5 = (_hx_tmp_8 != reflect.Interface)
		}

		if _hx_tmp_5 {
			return nil
		}

		var newMap reflect.Value
		_ = newMap
		newMap = reflect.MakeMap(mt)
		var keys []reflect.Value
		_ = keys
		keys = value.MapKeys()
		{
			var self []reflect.Value
			_ = self
			self = keys
			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[reflect.Value]
			_ = _g_array
			_g_array = HxMakeArray(self...)
			for {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g_current
				if !(_hx_tmp_9 < _g_array.Len()) {
					break
				}

				var _hx_tmp_10 HxArray[reflect.Value]
				_ = _hx_tmp_10
				_hx_tmp_10 = _g_array
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _g_current
				_g_current = (_g_current + 1)
				var key reflect.Value
				_ = key
				key = _hx_tmp_10.Get(_hx_tmp_11)
				var v reflect.Value
				_ = v
				v = value.MapIndex(key)
				newMap.SetMapIndex(key, v)
			}
		}

		return newMap.Interface()
	}

	panic(Hx_Field_haxe_exception_thrown(((any)("o is not an anon struct"))))
}

func Hx_Field_reflect_makeVarArgs(f func(_hx_carg0 HxArray[any]) any) any {
	return func(v ...any) any { return f(HxMakeArray[any](v...)) }
}

func init() {
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[reflect.Value]()
	hxRegisterArrayConstructor[string]()
}
