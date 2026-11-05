package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_objects() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_objects() {
}

type Hx_Obj_VTable_loreline_objects interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_objects struct {
	VTable Hx_Obj_VTable_loreline_objects `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_objects_CreateEmptyInstance() *Hx_Obj_loreline_objects {
	obj := &Hx_Obj_loreline_objects{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_objects_CreateInstance() *Hx_Obj_loreline_objects {
	obj := Hx_Obj_loreline_objects_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_objects_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Objects",
	HxMakeArray[string]("isFields", "getField", "getFields", "setField", "removeField", "fieldExists", "getStringHelper", "getArrayHelper", "getMapHelper", "getBeatHelper", "createFields"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_objects_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_objects_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_objects_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "isFields":
			return Hx_Field_loreline_objects_isFields
		case "getField":
			return Hx_Field_loreline_objects_getField
		case "getFields":
			return Hx_Field_loreline_objects_getFields
		case "setField":
			return Hx_Field_loreline_objects_setField
		case "removeField":
			return Hx_Field_loreline_objects_removeField
		case "fieldExists":
			return Hx_Field_loreline_objects_fieldExists
		case "getStringHelper":
			return Hx_Field_loreline_objects_getStringHelper
		case "getArrayHelper":
			return Hx_Field_loreline_objects_getArrayHelper
		case "getMapHelper":
			return Hx_Field_loreline_objects_getMapHelper
		case "getBeatHelper":
			return Hx_Field_loreline_objects_getBeatHelper
		case "createFields":
			return Hx_Field_loreline_objects_createFields
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_objects) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_objects_RTTI
}

func (this *Hx_Obj_loreline_objects) Hx_As_Hx_Obj_loreline_objects() *Hx_Obj_loreline_objects {
	return this
}

func Hx_Field_loreline_objects_isFields(value any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return false
	} else {
		if Hx_Field_std_isOfType(value, Hx_Obj_loreline_fields_RTTI) {
			return true
		} else {
			if Hx_Field_std_isOfType(value, Hx_Obj_haxe_ds_stringmap_RTTI) {
				return true
			} else {
				if Hx_Field_std_isOfType(value, Hx_Obj_haxe_ds_intmap_RTTI) {
					return false
				} else {
					var _hx_tmp_0 bool
					_ = _hx_tmp_0
					var _hx_tmp_1 bool
					_ = _hx_tmp_1
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					if Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI) {
						_hx_tmp_2 = true
					} else {
						_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
					}

					if _hx_tmp_2 {
						_hx_tmp_1 = true
					} else {
						_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
					}

					if _hx_tmp_1 {
						_hx_tmp_0 = true
					} else {
						_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI)
					}

					if _hx_tmp_0 {
						return false
					} else {
						if Hx_Field_loreline_arrays_isArray(value) {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

func Hx_Field_loreline_objects_getField(interpreter *Hx_Obj_loreline_interpreter, fields any, name string) any {
	if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
		var hx_dyncast_34 any
		_ = hx_dyncast_34
		hx_dyncast_34 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
		var _hx_tmp_0 *Hx_Obj_loreline_fields
		_ = _hx_tmp_0
		if hx_dyncast_34 != nil {
			_hx_tmp_0 = hx_dyncast_34.(*Hx_Obj_loreline_fields)
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0.VTable.Hx_Field_lorelineGet(interpreter, name)
	} else {
		if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
			var hx_dyncast_35 any
			_ = hx_dyncast_35
			hx_dyncast_35 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
			var _hx_tmp_1 *Hx_Obj_haxe_ds_stringmap
			_ = _hx_tmp_1
			if hx_dyncast_35 != nil {
				_hx_tmp_1 = hx_dyncast_35.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap()
			} else {
				_hx_tmp_1 = nil
			}

			return _hx_tmp_1.VTable.Hx_Field_get(name)
		} else {
			return Hx_Field_reflect_getProperty(fields, name)
		}
	}
}

func Hx_Field_loreline_objects_getFields(interpreter *Hx_Obj_loreline_interpreter, fields any) HxArray[string] {
	if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
		var hx_dyncast_36 any
		_ = hx_dyncast_36
		hx_dyncast_36 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
		var _hx_tmp_0 *Hx_Obj_loreline_fields
		_ = _hx_tmp_0
		if hx_dyncast_36 != nil {
			_hx_tmp_0 = hx_dyncast_36.(*Hx_Obj_loreline_fields)
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0.VTable.Hx_Field_lorelineFields(interpreter)
	} else {
		if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
			var _g HxArray[string]
			_ = _g
			_g = HxMakeArray[string]()
			{
				var hx_dyncast_37 any
				_ = hx_dyncast_37
				hx_dyncast_37 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_1 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_1
				if hx_dyncast_37 != nil {
					_hx_tmp_1 = hx_dyncast_37.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_1 = nil
				}

				var key any
				_ = key
				key = _hx_tmp_1.VTable.Hx_Field_keys()
				for {
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(key, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_3 any
					_ = _hx_tmp_3
					_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(key, "next")
					var key1 string
					_ = key1
					key1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))
					_ = (Hx_Array_Push[string](_g, key1))
				}
			}

			return _g
		} else {
			return Hx_Field_reflect_fields(fields)
		}
	}
}

func Hx_Field_loreline_objects_setField(interpreter *Hx_Obj_loreline_interpreter, fields any, name string, value any) {
	if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
		var hx_dyncast_38 any
		_ = hx_dyncast_38
		hx_dyncast_38 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
		var _hx_tmp_0 *Hx_Obj_loreline_fields
		_ = _hx_tmp_0
		if hx_dyncast_38 != nil {
			_hx_tmp_0 = hx_dyncast_38.(*Hx_Obj_loreline_fields)
		} else {
			_hx_tmp_0 = nil
		}

		_hx_tmp_0.VTable.Hx_Field_lorelineSet(interpreter, name, value)
	} else {
		if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
			var hx_dyncast_39 any
			_ = hx_dyncast_39
			hx_dyncast_39 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
			var _hx_tmp_1 *Hx_Obj_haxe_ds_stringmap
			_ = _hx_tmp_1
			if hx_dyncast_39 != nil {
				_hx_tmp_1 = hx_dyncast_39.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap()
			} else {
				_hx_tmp_1 = nil
			}

			_hx_tmp_1.VTable.Hx_Field_set(name, value)
		} else {
			Hx_Field_reflect_setProperty(fields, name, value)
		}
	}
}

func Hx_Field_loreline_objects_removeField(interpreter *Hx_Obj_loreline_interpreter, fields any, name string) bool {
	if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
		var hx_dyncast_40 any
		_ = hx_dyncast_40
		hx_dyncast_40 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
		var _hx_tmp_0 *Hx_Obj_loreline_fields
		_ = _hx_tmp_0
		if hx_dyncast_40 != nil {
			_hx_tmp_0 = hx_dyncast_40.(*Hx_Obj_loreline_fields)
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0.VTable.Hx_Field_lorelineRemove(interpreter, name)
	} else {
		if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
			return func() func(_hx_param_key string) bool {
				var hx_dyncast_41 any
				_ = hx_dyncast_41
				hx_dyncast_41 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_1 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_1
				if hx_dyncast_41 != nil {
					_hx_tmp_1 = hx_dyncast_41.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_1 = nil
				}

				var _hx_closure_src_194 func(_hx_param_key any) bool
				_ = _hx_closure_src_194
				_hx_closure_src_194 = _hx_tmp_1.VTable.Hx_Field_remove
				return func(_hx_param_key string) bool {
					return _hx_closure_src_194(((any)(_hx_param_key)))
				}
			}()(name)
		} else {
			return Hx_Field_reflect_deleteField(fields, name)
		}
	}
}

func Hx_Field_loreline_objects_fieldExists(interpreter *Hx_Obj_loreline_interpreter, fields any, name string) bool {
	if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
		var hx_dyncast_42 any
		_ = hx_dyncast_42
		hx_dyncast_42 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
		var _hx_tmp_0 *Hx_Obj_loreline_fields
		_ = _hx_tmp_0
		if hx_dyncast_42 != nil {
			_hx_tmp_0 = hx_dyncast_42.(*Hx_Obj_loreline_fields)
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0.VTable.Hx_Field_lorelineExists(interpreter, name)
	} else {
		if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
			return func() func(_hx_param_key string) bool {
				var hx_dyncast_43 any
				_ = hx_dyncast_43
				hx_dyncast_43 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_1 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_1
				if hx_dyncast_43 != nil {
					_hx_tmp_1 = hx_dyncast_43.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_1 = nil
				}

				var _hx_closure_src_195 func(_hx_param_key any) bool
				_ = _hx_closure_src_195
				_hx_closure_src_195 = _hx_tmp_1.VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_195(((any)(_hx_param_key)))
				}
			}()(name)
		} else {
			return Hx_Field_reflect_hasField(fields, name)
		}
	}
}

func Hx_Field_loreline_objects_getStringHelper(interpreter *Hx_Obj_loreline_interpreter, name string) any {
	return interpreter.Hx_Field_stringHelpers.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
}

func Hx_Field_loreline_objects_getArrayHelper(interpreter *Hx_Obj_loreline_interpreter, name string) any {
	return interpreter.Hx_Field_arrayHelpers.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
}

func Hx_Field_loreline_objects_getMapHelper(interpreter *Hx_Obj_loreline_interpreter, name string) any {
	return interpreter.Hx_Field_mapHelpers.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
}

func Hx_Field_loreline_objects_getBeatHelper(interpreter *Hx_Obj_loreline_interpreter, name string) any {
	return interpreter.Hx_Field_beatHelpers.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
}

func Hx_Field_loreline_objects_createFields(interpreter__isOptional any, _hx_reserved_type__isOptional any, node__isOptional any) any {
	interpreter := HxDefault[*Hx_Obj_loreline_interpreter]()
	_ = interpreter

	if interpreter__isOptional != nil {
		interpreter = HxConvert[*Hx_Obj_loreline_interpreter](interpreter__isOptional)
	} else {
		interpreter = nil
	}
	_hx_reserved_type := HxDefault[HxNullable[string]]()
	_ = _hx_reserved_type

	if _hx_reserved_type__isOptional != nil {
		_hx_reserved_type = HxConvert[HxNullable[string]](_hx_reserved_type__isOptional)
	} else {

	}
	node := HxDefault[*Hx_Obj_loreline_node]()
	_ = node

	if node__isOptional != nil {
		node = HxConvert[*Hx_Obj_loreline_node](node__isOptional)
	} else {
		node = nil
	}
	{
		if (_hx_reserved_type).Valid == false {
			var hx_nullable_t33 HxNullable[string]
			_ = hx_nullable_t33
			hx_nullable_t33 = HxNullable[string]{Value: HxStringNull}

			_hx_reserved_type = HxNullable[string]{Value: (hx_nullable_t33).Value, Valid: (hx_nullable_t33).Valid}
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if interpreter != nil {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(interpreter.Hx_Field_customCreateFields)), ((any)(nil))))))
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			var _hx_tmp_1 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_1
			_hx_tmp_1 = interpreter
			var customInstance any
			_ = customInstance
			customInstance = interpreter.Hx_Field_customCreateFields(_hx_tmp_1, (_hx_reserved_type).Value, node)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(customInstance, nil))) {
				return customInstance
			}
		}

		if (_hx_reserved_type).Valid != false {
			var instance any
			_ = instance
			instance = Hx_Field_type_createEmptyInstance(Hx_Field_type_resolveClass((_hx_reserved_type).Value))
			if Hx_Field_std_isOfType(instance, Hx_Obj_loreline_fields_RTTI) {
				var hx_dyncast_44 any
				_ = hx_dyncast_44
				hx_dyncast_44 = Hx_Field_go_haxe_hxdynamic_castClass(instance, "Hx_Obj_loreline_fields")
				var _hx_tmp_2 *Hx_Obj_loreline_fields
				_ = _hx_tmp_2
				if hx_dyncast_44 != nil {
					_hx_tmp_2 = hx_dyncast_44.(*Hx_Obj_loreline_fields)
				} else {
					_hx_tmp_2 = nil
				}

				var fields *Hx_Obj_loreline_fields
				_ = fields
				fields = _hx_tmp_2
				fields.VTable.Hx_Field_lorelineCreate(interpreter)
			}

			return instance
		}

		return ((any)(Hx_Obj_haxe_ds_stringmap_CreateInstance()))
	}
}

func init() {
	hxRegisterArrayConstructor[string]()
}
