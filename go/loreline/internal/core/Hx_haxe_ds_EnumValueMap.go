package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_enumvaluemap() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_enumvaluemap() {
}

type Hx_Obj_VTable_haxe_ds_enumvaluemap interface {
	Hx_Field_compareArg(v1 any, v2 any) int32
	Hx_Field_keyValueIterator() any
	Hx_Field_setLoop(k any, v any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_merge(t1 *Hx_Obj_haxe_ds_treenode, t2 *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_minBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_removeMinBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_getSortIndex(v any) HxNullable[int32]
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_iterator() any
	Hx_Field_compare(_hx_param_k1 any, _hx_param_k2 any) int32
	Hx_Field_compareArgs(a1 HxArray[any], a2 HxArray[any]) int32
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_keys() any
	Hx_Field_removeLoop(k any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_keysLoop(node *Hx_Obj_haxe_ds_treenode, acc HxArray[any])
	Hx_Field_balance(l *Hx_Obj_haxe_ds_treenode, k any, v any, r *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_toString() string
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_clear()
	Hx_Field_size() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_enumvaluemap struct {
	Hx_Obj_haxe_ds_balancedtree
	VTable                      Hx_Obj_VTable_haxe_ds_enumvaluemap `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_sortIndices        *Hx_Obj_haxe_imap
	Hx_Field_sortIndicesCounter int32
}

func Hx_Obj_haxe_ds_enumvaluemap_CreateEmptyInstance() *Hx_Obj_haxe_ds_enumvaluemap {
	obj := &Hx_Obj_haxe_ds_enumvaluemap{}
	obj.Hx_Obj_haxe_ds_balancedtree = *Hx_Obj_haxe_ds_balancedtree_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_haxe_ds_balancedtree.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Obj_haxe_ds_balancedtree.VTable = obj
	obj.Hx_Field_sortIndices = HxDefault[*Hx_Obj_haxe_imap]()
	obj.Hx_Field_sortIndicesCounter = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_ds_enumvaluemap_CreateInstance() *Hx_Obj_haxe_ds_enumvaluemap {
	obj := Hx_Obj_haxe_ds_enumvaluemap_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_New() {
	{
		this.Hx_Field_sortIndices = (func() *Hx_Obj_haxe_imap {
			_hx_icast_21 := Hx_Obj_haxe_ds_objectmap_CreateInstance()
			if _hx_icast_21 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_21, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_21.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_sortIndicesCounter = 0
	}
	{
		this.Hx_Obj_haxe_ds_balancedtree.Hx_New()
	}

}

var Hx_Obj_haxe_ds_enumvaluemap_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.EnumValueMap",
	HxMakeArray[string](),
	HxMakeArray[string]("sortIndices", "sortIndicesCounter", "compare", "compareArgs", "getSortIndex", "compareArg", "copy", "root", "set", "get", "remove", "exists", "iterator", "keyValueIterator", "keys", "setLoop", "removeLoop", "keysLoop", "merge", "minBinding", "removeMinBinding", "balance", "toString", "clear", "size"),
	Hx_Obj_haxe_ds_balancedtree_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_enumvaluemap_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_enumvaluemap_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_enumvaluemap_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_enumvaluemap_RTTI
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_As_Hx_Obj_haxe_ds_enumvaluemap() *Hx_Obj_haxe_ds_enumvaluemap {
	return this
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field_compare(_hx_param_k1 any, _hx_param_k2 any) int32 {
	k1 := func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_param_k1)
		if _hx_d == nil {
			var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
			return _hx_z
		}
		return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
	}()
	_ = k1
	k2 := func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_param_k2)
		if _hx_d == nil {
			var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
			return _hx_z
		}
		return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
	}()
	_ = k2
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_type_enumIndex(k1)
	var d int32
	_ = d
	d = (_hx_tmp_0 - Hx_Field_type_enumIndex(k2))
	if d != 0 {
		return d
	}

	var p1 HxArray[any]
	_ = p1
	p1 = Hx_Field_type_enumParameters(k1)
	var p2 HxArray[any]
	_ = p2
	p2 = Hx_Field_type_enumParameters(k2)
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if Hx_Field_go_haxe_hxdynamic_getArrayLength(p1) == 0 {
		_hx_tmp_1 = (Hx_Field_go_haxe_hxdynamic_getArrayLength(p2) == 0)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		return 0
	}

	return this.VTable.Hx_Field_compareArgs(p1, p2)
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field_compareArgs(a1 HxArray[any], a2 HxArray[any]) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getArrayLength(a1)
	var ld int32
	_ = ld
	ld = (_hx_tmp_0 - Hx_Field_go_haxe_hxdynamic_getArrayLength(a2))
	if ld != 0 {
		return ld
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(a1)
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			var d int32
			_ = d
			d = this.VTable.Hx_Field_compareArg(a1.Get(i), a2.Get(i))
			if d != 0 {
				return d
			}
		}
	}

	return 0
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field_getSortIndex(v any) HxNullable[int32] {
	if this.Hx_Field_sortIndices.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
	}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_exists(v) {
		var hx_nullable_t28 any
		_ = hx_nullable_t28
		hx_nullable_t28 = this.Hx_Field_sortIndices.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
		}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_get(v)
		var _hx_tmp_0 HxNullable[int32]
		_ = _hx_tmp_0
		if hx_nullable_t28 == nil {
			_hx_tmp_0 = HxNullable[int32]{}
		} else {
			_hx_tmp_0 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t28))), Valid: true}
		}

		return _hx_tmp_0
	}

	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_sortIndicesCounter
	this.Hx_Field_sortIndicesCounter = (this.Hx_Field_sortIndicesCounter + 1)
	var sort int32
	_ = sort
	sort = _hx_tmp_1
	this.Hx_Field_sortIndices.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_objectmap() *Hx_Obj_haxe_ds_objectmap
	}).Hx_As_Hx_Obj_haxe_ds_objectmap().VTable.Hx_Field_set(v, sort)
	return HxNullable[int32]{Value: sort, Valid: true}
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field_compareArg(v1 any, v2 any) int32 {
	var vt1 Hx_Enum_valuetype
	_ = vt1
	vt1 = Hx_Field_type_typeof(v1)
	var vt2 Hx_Enum_valuetype
	_ = vt2
	vt2 = Hx_Field_type_typeof(v2)
	switch vt1.Hx_Field_enumIndex() {
	case 0:
		{
			if vt2.Hx_Field_enumIndex() == 0 {
				return 0
			} else {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_64 func(a any, b any) int32
					_ = _hx_closure_src_64
					_hx_closure_src_64 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_64(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_0, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 1:
		{
			if vt2.Hx_Field_enumIndex() == 1 {
				return Hx_Field_reflect_compare(v1, v2)
			} else {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_65 func(a any, b any) int32
					_ = _hx_closure_src_65
					_hx_closure_src_65 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_65(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_1, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 3:
		{
			if vt2.Hx_Field_enumIndex() == 3 {
				return Hx_Field_reflect_compare(v1, v2)
			} else {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_66 func(a any, b any) int32
					_ = _hx_closure_src_66
					_hx_closure_src_66 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_66(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_2, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 4:
		{
			if vt2.Hx_Field_enumIndex() == 4 {
				return Hx_Field_reflect_compare(v1, v2)
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_67 func(a any, b any) int32
					_ = _hx_closure_src_67
					_hx_closure_src_67 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_67(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_3, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 5:
		{
			if vt2.Hx_Field_enumIndex() == 5 {
				var _hx_tmp_6 HxNullable[int32]
				_ = _hx_tmp_6
				_hx_tmp_6 = this.VTable.Hx_Field_getSortIndex(v1)
				return func() func(a HxNullable[int32], b HxNullable[int32]) int32 {
					var _hx_closure_src_68 func(a any, b any) int32
					_ = _hx_closure_src_68
					_hx_closure_src_68 = Hx_Field_reflect_compare
					return func(a HxNullable[int32], b HxNullable[int32]) int32 {
						var hx_nullable_f4 HxNullable[int32]
						_ = hx_nullable_f4
						hx_nullable_f4 = a
						var _hx_tmp_4 any
						_ = _hx_tmp_4
						if (hx_nullable_f4).Valid == true {
							_hx_tmp_4 = (hx_nullable_f4).Value
						} else {
							_hx_tmp_4 = nil
						}

						var hx_nullable_f5 HxNullable[int32]
						_ = hx_nullable_f5
						hx_nullable_f5 = b
						var _hx_tmp_5 any
						_ = _hx_tmp_5
						if (hx_nullable_f5).Valid == true {
							_hx_tmp_5 = (hx_nullable_f5).Value
						} else {
							_hx_tmp_5 = nil
						}

						return _hx_closure_src_68(_hx_tmp_4, _hx_tmp_5)
					}
				}()(_hx_tmp_6, this.VTable.Hx_Field_getSortIndex(v2))
			} else {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_69 func(a any, b any) int32
					_ = _hx_closure_src_69
					_hx_closure_src_69 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_69(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_7, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 6:
		{
			if vt2.Hx_Field_enumIndex() == 6 {
				var _hx_tmp_8 string
				_ = _hx_tmp_8
				_hx_tmp_8 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unsupported comparison types: ", Hx_Field_std_string(vt1)), " ")
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_8, Hx_Field_std_string(vt2))))))
			} else {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_70 func(a any, b any) int32
					_ = _hx_closure_src_70
					_hx_closure_src_70 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_70(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_9, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 7:
		{
			if vt1.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
				if vt2.Hx_Field_enumIndex() == 7 {
					if vt2.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
						return Hx_Field_reflect_compare(v1, v2)
					} else {
						var _hx_tmp_12 HxNullable[int32]
						_ = _hx_tmp_12
						_hx_tmp_12 = this.VTable.Hx_Field_getSortIndex(v1)
						return func() func(a HxNullable[int32], b HxNullable[int32]) int32 {
							var _hx_closure_src_71 func(a any, b any) int32
							_ = _hx_closure_src_71
							_hx_closure_src_71 = Hx_Field_reflect_compare
							return func(a HxNullable[int32], b HxNullable[int32]) int32 {
								var hx_nullable_f6 HxNullable[int32]
								_ = hx_nullable_f6
								hx_nullable_f6 = a
								var _hx_tmp_10 any
								_ = _hx_tmp_10
								if (hx_nullable_f6).Valid == true {
									_hx_tmp_10 = (hx_nullable_f6).Value
								} else {
									_hx_tmp_10 = nil
								}

								var hx_nullable_f7 HxNullable[int32]
								_ = hx_nullable_f7
								hx_nullable_f7 = b
								var _hx_tmp_11 any
								_ = _hx_tmp_11
								if (hx_nullable_f7).Valid == true {
									_hx_tmp_11 = (hx_nullable_f7).Value
								} else {
									_hx_tmp_11 = nil
								}

								return _hx_closure_src_71(_hx_tmp_10, _hx_tmp_11)
							}
						}()(_hx_tmp_12, this.VTable.Hx_Field_getSortIndex(v2))
					}
				} else {
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
					return func() func(a int32, b int32) int32 {
						var _hx_closure_src_72 func(a any, b any) int32
						_ = _hx_closure_src_72
						_hx_closure_src_72 = Hx_Field_reflect_compare
						return func(a int32, b int32) int32 {
							return _hx_closure_src_72(((any)(a)), ((any)(b)))
						}
					}()(_hx_tmp_13, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
				}
			} else {
				if vt2.Hx_Field_enumIndex() == 7 {
					var _hx_tmp_16 HxNullable[int32]
					_ = _hx_tmp_16
					_hx_tmp_16 = this.VTable.Hx_Field_getSortIndex(v1)
					return func() func(a HxNullable[int32], b HxNullable[int32]) int32 {
						var _hx_closure_src_73 func(a any, b any) int32
						_ = _hx_closure_src_73
						_hx_closure_src_73 = Hx_Field_reflect_compare
						return func(a HxNullable[int32], b HxNullable[int32]) int32 {
							var hx_nullable_f8 HxNullable[int32]
							_ = hx_nullable_f8
							hx_nullable_f8 = a
							var _hx_tmp_14 any
							_ = _hx_tmp_14
							if (hx_nullable_f8).Valid == true {
								_hx_tmp_14 = (hx_nullable_f8).Value
							} else {
								_hx_tmp_14 = nil
							}

							var hx_nullable_f9 HxNullable[int32]
							_ = hx_nullable_f9
							hx_nullable_f9 = b
							var _hx_tmp_15 any
							_ = _hx_tmp_15
							if (hx_nullable_f9).Valid == true {
								_hx_tmp_15 = (hx_nullable_f9).Value
							} else {
								_hx_tmp_15 = nil
							}

							return _hx_closure_src_73(_hx_tmp_14, _hx_tmp_15)
						}
					}()(_hx_tmp_16, this.VTable.Hx_Field_getSortIndex(v2))
				} else {
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					_hx_tmp_17 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
					return func() func(a int32, b int32) int32 {
						var _hx_closure_src_74 func(a any, b any) int32
						_ = _hx_closure_src_74
						_hx_closure_src_74 = Hx_Field_reflect_compare
						return func(a int32, b int32) int32 {
							return _hx_closure_src_74(((any)(a)), ((any)(b)))
						}
					}()(_hx_tmp_17, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
				}
			}
		}

	case 8:
		{
			if vt2.Hx_Field_enumIndex() == 8 {
				return func() func(_hx_param_k1 Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue, _hx_param_k2 Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue) int32 {
					var _hx_closure_src_75 func(_hx_param_k1 any, _hx_param_k2 any) int32
					_ = _hx_closure_src_75
					_hx_closure_src_75 = this.VTable.Hx_Field_compare
					return func(_hx_param_k1 Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue, _hx_param_k2 Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue) int32 {
						return _hx_closure_src_75(((any)(_hx_param_k1)), ((any)(_hx_param_k2)))
					}
				}()(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(v1)
					if _hx_d == nil {
						var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
						return _hx_z
					}
					return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
				}(), func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(v2)
					if _hx_d == nil {
						var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
						return _hx_z
					}
					return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
				}())
			} else {
				var _hx_tmp_18 int32
				_ = _hx_tmp_18
				_hx_tmp_18 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_76 func(a any, b any) int32
					_ = _hx_closure_src_76
					_hx_closure_src_76 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_76(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_18, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	case 9:
		{
			if vt2.Hx_Field_enumIndex() == 9 {
				var _hx_tmp_19 string
				_ = _hx_tmp_19
				_hx_tmp_19 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unsupported comparison types: ", Hx_Field_std_string(vt1)), " ")
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_19, Hx_Field_std_string(vt2))))))
			} else {
				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				_hx_tmp_20 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
				return func() func(a int32, b int32) int32 {
					var _hx_closure_src_77 func(a any, b any) int32
					_ = _hx_closure_src_77
					_hx_closure_src_77 = Hx_Field_reflect_compare
					return func(a int32, b int32) int32 {
						return _hx_closure_src_77(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_20, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
			}
		}

	default:
		{
			var _hx_tmp_21 int32
			_ = _hx_tmp_21
			_hx_tmp_21 = Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt1)))
			return func() func(a int32, b int32) int32 {
				var _hx_closure_src_78 func(a any, b any) int32
				_ = _hx_closure_src_78
				_hx_closure_src_78 = Hx_Field_reflect_compare
				return func(a int32, b int32) int32 {
					return _hx_closure_src_78(((any)(a)), ((any)(b)))
				}
			}()(_hx_tmp_21, Hx_Field_type_enumIndex(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(vt2))))
		}
	}
}

func (this *Hx_Obj_haxe_ds_enumvaluemap) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_enumvaluemap
	_ = copied
	copied = Hx_Obj_haxe_ds_enumvaluemap_CreateInstance()
	copied.Hx_Field_root = this.Hx_Field_root
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_22 := copied
		if _hx_icast_22 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_22, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_22.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func init() {
	hxRegisterArrayConstructor[any]()
}
