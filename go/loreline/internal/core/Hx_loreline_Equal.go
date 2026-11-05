package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_equal() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_equal() {
}

type Hx_Obj_VTable_loreline_equal interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_equal struct {
	VTable Hx_Obj_VTable_loreline_equal `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_equal_CreateEmptyInstance() *Hx_Obj_loreline_equal {
	obj := &Hx_Obj_loreline_equal{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_equal_CreateInstance() *Hx_Obj_loreline_equal {
	obj := Hx_Obj_loreline_equal_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_equal_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Equal",
	HxMakeArray[string]("equal", "objectFieldsEqual", "arrayEqual", "stringMapEqual", "intMapEqual"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_equal_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_equal_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_equal_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "equal":
			return Hx_Field_loreline_equal_equal
		case "objectFieldsEqual":
			return Hx_Field_loreline_equal_objectFieldsEqual
		case "arrayEqual":
			return Hx_Field_loreline_equal_arrayEqual
		case "stringMapEqual":
			return Hx_Field_loreline_equal_stringMapEqual
		case "intMapEqual":
			return Hx_Field_loreline_equal_intMapEqual
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_equal) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_equal_RTTI
}

func (this *Hx_Obj_loreline_equal) Hx_As_Hx_Obj_loreline_equal() *Hx_Obj_loreline_equal {
	return this
}

func Hx_Field_loreline_equal_equal(interpreter *Hx_Obj_loreline_interpreter, a any, b any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, b))) {
		return true
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(b, nil))))
	}

	if _hx_tmp_0 {
		return false
	}

	if Hx_Field_loreline_arrays_isArray(a) {
		if Hx_Field_loreline_arrays_isArray(b) {
			return Hx_Field_loreline_equal_arrayEqual(interpreter, a, b)
		}

		return false
	} else {
		if Hx_Field_std_isOfType(a, Hx_Obj_haxe_ds_stringmap_RTTI) {
			if Hx_Field_std_isOfType(b, Hx_Obj_haxe_ds_stringmap_RTTI) {
				var _hx_tmp_1 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_1
				_hx_tmp_1 = interpreter
				var hx_dyncast_171 any
				_ = hx_dyncast_171
				hx_dyncast_171 = Hx_Field_go_haxe_hxdynamic_castClass(a, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_2 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_2
				if hx_dyncast_171 != nil {
					_hx_tmp_2 = hx_dyncast_171.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_2 = nil
				}

				var hx_dyncast_172 any
				_ = hx_dyncast_172
				hx_dyncast_172 = Hx_Field_go_haxe_hxdynamic_castClass(b, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_3 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_3
				if hx_dyncast_172 != nil {
					_hx_tmp_3 = hx_dyncast_172.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_3 = nil
				}

				return Hx_Field_loreline_equal_stringMapEqual(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3)
			}

			return false
		} else {
			if Hx_Field_std_isOfType(a, Hx_Obj_haxe_ds_intmap_RTTI) {
				if Hx_Field_std_isOfType(b, Hx_Obj_haxe_ds_intmap_RTTI) {
					var _hx_tmp_4 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_4
					_hx_tmp_4 = interpreter
					var hx_dyncast_173 any
					_ = hx_dyncast_173
					hx_dyncast_173 = Hx_Field_go_haxe_hxdynamic_castClass(a, "Hx_Obj_haxe_ds_intmap")
					var _hx_tmp_5 *Hx_Obj_haxe_ds_intmap
					_ = _hx_tmp_5
					if hx_dyncast_173 != nil {
						_hx_tmp_5 = hx_dyncast_173.(interface{ Hx_As_Hx_Obj_haxe_ds_intmap() *Hx_Obj_haxe_ds_intmap }).Hx_As_Hx_Obj_haxe_ds_intmap()
					} else {
						_hx_tmp_5 = nil
					}

					var hx_dyncast_174 any
					_ = hx_dyncast_174
					hx_dyncast_174 = Hx_Field_go_haxe_hxdynamic_castClass(b, "Hx_Obj_haxe_ds_intmap")
					var _hx_tmp_6 *Hx_Obj_haxe_ds_intmap
					_ = _hx_tmp_6
					if hx_dyncast_174 != nil {
						_hx_tmp_6 = hx_dyncast_174.(interface{ Hx_As_Hx_Obj_haxe_ds_intmap() *Hx_Obj_haxe_ds_intmap }).Hx_As_Hx_Obj_haxe_ds_intmap()
					} else {
						_hx_tmp_6 = nil
					}

					return Hx_Field_loreline_equal_intMapEqual(_hx_tmp_4, _hx_tmp_5, _hx_tmp_6)
				}

				return false
			} else {
				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				if Hx_Field_loreline_runtimebeatref_beatOf(a) != nil {
					_hx_tmp_7 = true
				} else {
					_hx_tmp_7 = (Hx_Field_loreline_runtimebeatref_beatOf(b) != nil)
				}

				if _hx_tmp_7 {
					var aBeat *Hx_Obj_loreline_nbeatdecl
					_ = aBeat
					aBeat = Hx_Field_loreline_runtimebeatref_beatOf(a)
					var bBeat *Hx_Obj_loreline_nbeatdecl
					_ = bBeat
					bBeat = Hx_Field_loreline_runtimebeatref_beatOf(b)
					return (((aBeat != nil) && (bBeat != nil)) && (aBeat == bBeat))
				} else {
					var _hx_tmp_8 bool
					_ = _hx_tmp_8
					if Hx_Field_loreline_runtimecharacterref_characterOf(a) != nil {
						_hx_tmp_8 = true
					} else {
						_hx_tmp_8 = (Hx_Field_loreline_runtimecharacterref_characterOf(b) != nil)
					}

					if _hx_tmp_8 {
						var aValue any
						_ = aValue
						aValue = Hx_Field_loreline_runtimecharacterref_fieldsOf(a)
						var bValue any
						_ = bValue
						bValue = Hx_Field_loreline_runtimecharacterref_fieldsOf(b)
						return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(aValue, bValue))))
					} else {
						if Hx_Field_loreline_objects_isFields(a) {
							if Hx_Field_loreline_objects_isFields(b) {
								return Hx_Field_loreline_equal_objectFieldsEqual(interpreter, a, b)
							}

							return false
						}
					}
				}
			}
		}
	}

	return false
}

func Hx_Field_loreline_equal_objectFieldsEqual(interpreter *Hx_Obj_loreline_interpreter, a any, b any) bool {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = Hx_Field_loreline_objects_getFields(interpreter, a)
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var field string
			_ = field
			field = _g1.Get(_g)
			_g++
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if !Hx_Field_loreline_objects_fieldExists(interpreter, b, field) {
				_hx_tmp_1 = true
			} else {
				var _hx_tmp_2 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_2
				_hx_tmp_2 = interpreter
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_objects_getField(interpreter, a, field)
				var _hx_tmp_4 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_2
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_3
				_hx_tmp_1 = !Hx_Field_loreline_equal_equal(_hx_tmp_4, _hx_tmp_5, Hx_Field_loreline_objects_getField(interpreter, b, field))
			}

			if _hx_tmp_1 {
				return false
			}
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 HxArray[string]
		_ = _g3
		_g3 = Hx_Field_loreline_objects_getFields(interpreter, b)
		for {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = _g2
			if !(_hx_tmp_6 < _g3.Len()) {
				break
			}

			var field string
			_ = field
			field = _g3.Get(_g2)
			_g2++
			if !Hx_Field_loreline_objects_fieldExists(interpreter, a, field) {
				return false
			}
		}
	}

	return true
}

func Hx_Field_loreline_equal_arrayEqual(interpreter *Hx_Obj_loreline_interpreter, a any, b any) bool {
	var lenA int32
	_ = lenA
	lenA = Hx_Field_loreline_arrays_arrayLength(a)
	var lenB int32
	_ = lenB
	lenB = Hx_Field_loreline_arrays_arrayLength(b)
	if lenA != lenB {
		return false
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = lenA
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var _hx_tmp_1 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_1
			_hx_tmp_1 = interpreter
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_loreline_arrays_arrayGet(a, i)
			if !Hx_Field_loreline_equal_equal(_hx_tmp_1, _hx_tmp_2, Hx_Field_loreline_arrays_arrayGet(b, i)) {
				return false
			}
		}
	}

	return true
}

func Hx_Field_loreline_equal_stringMapEqual(interpreter *Hx_Obj_loreline_interpreter, a *Hx_Obj_haxe_ds_stringmap, b *Hx_Obj_haxe_ds_stringmap) bool {
	{
		var _hx_reserved_map *Hx_Obj_haxe_imap
		_ = _hx_reserved_map
		_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
			_hx_icast_175 := a
			if _hx_icast_175 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_175, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_175.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		var _g_map *Hx_Obj_haxe_imap
		_ = _g_map
		_g_map = _hx_reserved_map
		var _g_keys any
		_ = _g_keys
		_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
			var key string
			_ = key
			key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _g_value any
			_ = _g_value
			_g_value = _g_map.VTable.Hx_Field_get(key)
			var _g_key string
			_ = _g_key
			_g_key = key
			var key1 string
			_ = key1
			key1 = _g_key
			var val any
			_ = val
			val = _g_value
			{
				if !func() func(_hx_param_key string) bool {
					var _hx_closure_src_536 func(_hx_param_key any) bool
					_ = _hx_closure_src_536
					_hx_closure_src_536 = b.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_536(((any)(_hx_param_key)))
					}
				}()(key1) {
					return false
				}

				var _hx_tmp_2 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_2
				_hx_tmp_2 = interpreter
				if !Hx_Field_loreline_equal_equal(_hx_tmp_2, b.VTable.Hx_Field_get(key1), val) {
					return false
				}
			}
		}
	}

	{
		var _gkey any
		_ = _gkey
		_gkey = b.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "next")
			var key string
			_ = key
			key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()))
			if !func() func(_hx_param_key string) bool {
				var _hx_closure_src_537 func(_hx_param_key any) bool
				_ = _hx_closure_src_537
				_hx_closure_src_537 = a.VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_537(((any)(_hx_param_key)))
				}
			}()(key) {
				return false
			}
		}
	}

	return true
}

func Hx_Field_loreline_equal_intMapEqual(interpreter *Hx_Obj_loreline_interpreter, a *Hx_Obj_haxe_ds_intmap, b *Hx_Obj_haxe_ds_intmap) bool {
	{
		var _hx_reserved_map *Hx_Obj_haxe_imap
		_ = _hx_reserved_map
		_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
			_hx_icast_176 := a
			if _hx_icast_176 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_176, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_176.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		var _g_map *Hx_Obj_haxe_imap
		_ = _g_map
		_g_map = _hx_reserved_map
		var _g_keys any
		_ = _g_keys
		_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
			var key int32
			_ = key
			key = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))))
			var _g_value any
			_ = _g_value
			_g_value = _g_map.VTable.Hx_Field_get(key)
			var _g_key int32
			_ = _g_key
			_g_key = key
			var key1 int32
			_ = key1
			key1 = _g_key
			var val any
			_ = val
			val = _g_value
			{
				if !func() func(_hx_param_key int32) bool {
					var _hx_closure_src_538 func(_hx_param_key any) bool
					_ = _hx_closure_src_538
					_hx_closure_src_538 = b.VTable.Hx_Field_exists
					return func(_hx_param_key int32) bool {
						return _hx_closure_src_538(((any)(_hx_param_key)))
					}
				}()(key1) {
					return false
				}

				var _hx_tmp_2 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_2
				_hx_tmp_2 = interpreter
				if !Hx_Field_loreline_equal_equal(_hx_tmp_2, b.VTable.Hx_Field_get(key1), val) {
					return false
				}
			}
		}
	}

	{
		var _gkey any
		_ = _gkey
		_gkey = b.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "next")
			var key int32
			_ = key
			key = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()))))
			if !func() func(_hx_param_key int32) bool {
				var _hx_closure_src_539 func(_hx_param_key any) bool
				_ = _hx_closure_src_539
				_hx_closure_src_539 = a.VTable.Hx_Field_exists
				return func(_hx_param_key int32) bool {
					return _hx_closure_src_539(((any)(_hx_param_key)))
				}
			}()(key) {
				return false
			}
		}
	}

	return true
}

func init() {
	hxRegisterArrayConstructor[string]()
}
