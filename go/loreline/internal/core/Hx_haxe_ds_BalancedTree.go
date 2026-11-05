package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_balancedtree() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_balancedtree() {
}

type Hx_Obj_VTable_haxe_ds_balancedtree interface {
	Hx_Field_keys() any
	Hx_Field_clear()
	Hx_Field_size() int32
	Hx_Field_exists(_hx_param_key any) bool
	Hx_Field_minBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_balance(l *Hx_Obj_haxe_ds_treenode, k any, v any, r *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_compare(k1 any, k2 any) int32
	Hx_Field_setLoop(k any, v any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_removeLoop(k any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_set(_hx_param_key any, _hx_param_value any)
	Hx_Field_iterator() any
	Hx_Field_keyValueIterator() any
	Hx_Field_copy() *Hx_Obj_haxe_imap
	Hx_Field_keysLoop(node *Hx_Obj_haxe_ds_treenode, acc HxArray[any])
	Hx_Field_merge(t1 *Hx_Obj_haxe_ds_treenode, t2 *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_removeMinBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode
	Hx_Field_toString() string
	Hx_Field_get(_hx_param_key any) any
	Hx_Field_remove(_hx_param_key any) bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_balancedtree struct {
	Hx_Obj_haxe_imap
	VTable        Hx_Obj_VTable_haxe_ds_balancedtree `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_root *Hx_Obj_haxe_ds_treenode
}

func Hx_Obj_haxe_ds_balancedtree_CreateEmptyInstance() *Hx_Obj_haxe_ds_balancedtree {
	obj := &Hx_Obj_haxe_ds_balancedtree{}
	obj.VTable = obj
	obj.Hx_Obj_haxe_imap.VTable = obj
	obj.Hx_Field_root = HxDefault[*Hx_Obj_haxe_ds_treenode]()
	return obj
}

func Hx_Obj_haxe_ds_balancedtree_CreateInstance() *Hx_Obj_haxe_ds_balancedtree {
	obj := Hx_Obj_haxe_ds_balancedtree_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_New() {
	{
	}

}

var Hx_Obj_haxe_ds_balancedtree_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.BalancedTree",
	HxMakeArray[string]("iteratorLoop", "sizeLoop"),
	HxMakeArray[string]("root", "set", "get", "remove", "exists", "iterator", "keyValueIterator", "keys", "copy", "setLoop", "removeLoop", "keysLoop", "merge", "minBinding", "removeMinBinding", "balance", "compare", "toString", "clear", "size"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](Hx_Obj_haxe_imap_RTTI),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_balancedtree_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_balancedtree_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_balancedtree_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "iteratorLoop":
			return Hx_Field_haxe_ds_balancedtree_iteratorLoop
		case "sizeLoop":
			return Hx_Field_haxe_ds_balancedtree_sizeLoop
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_balancedtree_RTTI
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_As_Hx_Obj_haxe_ds_balancedtree() *Hx_Obj_haxe_ds_balancedtree {
	return this
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_set(_hx_param_key any, _hx_param_value any) {
	key := _hx_param_key
	_ = key
	value := _hx_param_value
	_ = value
	this.Hx_Field_root = this.VTable.Hx_Field_setLoop(key, value, this.Hx_Field_root)
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_get(_hx_param_key any) any {
	key := _hx_param_key
	_ = key
	var node *Hx_Obj_haxe_ds_treenode
	_ = node
	node = this.Hx_Field_root
	for node != nil {
		var c int32
		_ = c
		c = this.VTable.Hx_Field_compare(key, node.Hx_Field_key)
		if c == 0 {
			return node.Hx_Field_value
		}

		if c < 0 {
			node = node.Hx_Field_left
		} else {
			node = node.Hx_Field_right
		}
	}

	return nil
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_remove(_hx_param_key any) bool {
	key := _hx_param_key
	_ = key
	{
		var hx_try_return bool
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "string").(string); _hx_ok {
						_ = e
						{
							hx_try_state = 1
							hx_try_return = false
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				this.Hx_Field_root = this.VTable.Hx_Field_removeLoop(key, this.Hx_Field_root)
				hx_try_state = 1
				hx_try_return = true
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_exists(_hx_param_key any) bool {
	key := _hx_param_key
	_ = key
	var node *Hx_Obj_haxe_ds_treenode
	_ = node
	node = this.Hx_Field_root
	for node != nil {
		var c int32
		_ = c
		c = this.VTable.Hx_Field_compare(key, node.Hx_Field_key)
		if c == 0 {
			return true
		} else {
			if c < 0 {
				node = node.Hx_Field_left
			} else {
				node = node.Hx_Field_right
			}
		}
	}

	return false
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_iterator() any {
	var ret HxArray[any]
	_ = ret
	ret = HxMakeArray[any]().Dyn()
	Hx_Field_haxe_ds_balancedtree_iteratorLoop(this.Hx_Field_root, ret)
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(ret)))
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_keyValueIterator() any {
	return ((any)(Hx_Obj_haxe_iterators_mapkeyvalueiterator_CreateInstance((func() *Hx_Obj_haxe_imap {
		_hx_icast_16 := this
		if _hx_icast_16 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_16, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_16.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())))
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_keys() any {
	var ret HxArray[any]
	_ = ret
	ret = HxMakeArray[any]().Dyn()
	this.VTable.Hx_Field_keysLoop(this.Hx_Field_root, ret)
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(ret)))
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_copy() *Hx_Obj_haxe_imap {
	var copied *Hx_Obj_haxe_ds_balancedtree
	_ = copied
	copied = Hx_Obj_haxe_ds_balancedtree_CreateInstance()
	copied.Hx_Field_root = this.Hx_Field_root
	return (func() *Hx_Obj_haxe_imap {
		_hx_icast_17 := copied
		if _hx_icast_17 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_17, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_17.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_setLoop(k any, v any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	if node == nil {
		return Hx_Obj_haxe_ds_treenode_CreateInstance(nil, k, v, nil, nil)
	}

	var c int32
	_ = c
	c = this.VTable.Hx_Field_compare(k, node.Hx_Field_key)
	if c == 0 {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if node == nil {
			_hx_tmp_0 = 0
		} else {
			_hx_tmp_0 = node.Hx_Field__height
		}

		return Hx_Obj_haxe_ds_treenode_CreateInstance(node.Hx_Field_left, k, v, node.Hx_Field_right, _hx_tmp_0)
	} else {
		if c < 0 {
			var nl *Hx_Obj_haxe_ds_treenode
			_ = nl
			nl = this.VTable.Hx_Field_setLoop(k, v, node.Hx_Field_left)
			return this.VTable.Hx_Field_balance(nl, node.Hx_Field_key, node.Hx_Field_value, node.Hx_Field_right)
		} else {
			var nr *Hx_Obj_haxe_ds_treenode
			_ = nr
			nr = this.VTable.Hx_Field_setLoop(k, v, node.Hx_Field_right)
			return this.VTable.Hx_Field_balance(node.Hx_Field_left, node.Hx_Field_key, node.Hx_Field_value, nr)
		}
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_removeLoop(k any, node *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	if node == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)("Not_found"))))
	}

	var c int32
	_ = c
	c = this.VTable.Hx_Field_compare(k, node.Hx_Field_key)
	if c == 0 {
		return this.VTable.Hx_Field_merge(node.Hx_Field_left, node.Hx_Field_right)
	} else {
		if c < 0 {
			return this.VTable.Hx_Field_balance(this.VTable.Hx_Field_removeLoop(k, node.Hx_Field_left), node.Hx_Field_key, node.Hx_Field_value, node.Hx_Field_right)
		} else {
			var _hx_tmp_0 *Hx_Obj_haxe_ds_treenode
			_ = _hx_tmp_0
			_hx_tmp_0 = node.Hx_Field_left
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = node.Hx_Field_key
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = node.Hx_Field_value
			return this.VTable.Hx_Field_balance(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, this.VTable.Hx_Field_removeLoop(k, node.Hx_Field_right))
		}
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_keysLoop(node *Hx_Obj_haxe_ds_treenode, acc HxArray[any]) {
	if node != nil {
		this.VTable.Hx_Field_keysLoop(node.Hx_Field_left, acc)
		_ = (Hx_Array_Push[any](HxMakeArrayView[any](acc), node.Hx_Field_key))
		this.VTable.Hx_Field_keysLoop(node.Hx_Field_right, acc)
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_merge(t1 *Hx_Obj_haxe_ds_treenode, t2 *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	var t *Hx_Obj_haxe_ds_treenode
	_ = t
	t = this.VTable.Hx_Field_minBinding(t2)
	var _hx_tmp_0 *Hx_Obj_haxe_ds_treenode
	_ = _hx_tmp_0
	_hx_tmp_0 = t1
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = t.Hx_Field_key
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = t.Hx_Field_value
	return this.VTable.Hx_Field_balance(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, this.VTable.Hx_Field_removeMinBinding(t2))
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_minBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	if t == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)("Not_found"))))
	} else {
		if t.Hx_Field_left == nil {
			return t
		} else {
			return this.VTable.Hx_Field_minBinding(t.Hx_Field_left)
		}
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_removeMinBinding(t *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	if t.Hx_Field_left == nil {
		return t.Hx_Field_right
	} else {
		return this.VTable.Hx_Field_balance(this.VTable.Hx_Field_removeMinBinding(t.Hx_Field_left), t.Hx_Field_key, t.Hx_Field_value, t.Hx_Field_right)
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_balance(l *Hx_Obj_haxe_ds_treenode, k any, v any, r *Hx_Obj_haxe_ds_treenode) *Hx_Obj_haxe_ds_treenode {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if l == nil {
		_hx_tmp_0 = 0
	} else {
		_hx_tmp_0 = l.Hx_Field__height
	}

	var hl int32
	_ = hl
	hl = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if r == nil {
		_hx_tmp_1 = 0
	} else {
		_hx_tmp_1 = r.Hx_Field__height
	}

	var hr int32
	_ = hr
	hr = _hx_tmp_1
	if hl > (hr + 2) {
		var _this *Hx_Obj_haxe_ds_treenode
		_ = _this
		_this = l.Hx_Field_left
		var _this1 *Hx_Obj_haxe_ds_treenode
		_ = _this1
		_this1 = l.Hx_Field_right
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if _this == nil {
			_hx_tmp_2 = 0
		} else {
			_hx_tmp_2 = _this.Hx_Field__height
		}

		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if _this1 == nil {
			_hx_tmp_3 = 0
		} else {
			_hx_tmp_3 = _this1.Hx_Field__height
		}

		if (_hx_tmp_2) >= (_hx_tmp_3) {
			return Hx_Obj_haxe_ds_treenode_CreateInstance(l.Hx_Field_left, l.Hx_Field_key, l.Hx_Field_value, Hx_Obj_haxe_ds_treenode_CreateInstance(l.Hx_Field_right, k, v, r, nil), nil)
		} else {
			return Hx_Obj_haxe_ds_treenode_CreateInstance(Hx_Obj_haxe_ds_treenode_CreateInstance(l.Hx_Field_left, l.Hx_Field_key, l.Hx_Field_value, l.Hx_Field_right.Hx_Field_left, nil), l.Hx_Field_right.Hx_Field_key, l.Hx_Field_right.Hx_Field_value, Hx_Obj_haxe_ds_treenode_CreateInstance(l.Hx_Field_right.Hx_Field_right, k, v, r, nil), nil)
		}
	} else {
		if hr > (hl + 2) {
			var _this *Hx_Obj_haxe_ds_treenode
			_ = _this
			_this = r.Hx_Field_right
			var _this1 *Hx_Obj_haxe_ds_treenode
			_ = _this1
			_this1 = r.Hx_Field_left
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			if _this == nil {
				_hx_tmp_4 = 0
			} else {
				_hx_tmp_4 = _this.Hx_Field__height
			}

			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			if _this1 == nil {
				_hx_tmp_5 = 0
			} else {
				_hx_tmp_5 = _this1.Hx_Field__height
			}

			if (_hx_tmp_4) > (_hx_tmp_5) {
				return Hx_Obj_haxe_ds_treenode_CreateInstance(Hx_Obj_haxe_ds_treenode_CreateInstance(l, k, v, r.Hx_Field_left, nil), r.Hx_Field_key, r.Hx_Field_value, r.Hx_Field_right, nil)
			} else {
				return Hx_Obj_haxe_ds_treenode_CreateInstance(Hx_Obj_haxe_ds_treenode_CreateInstance(l, k, v, r.Hx_Field_left.Hx_Field_left, nil), r.Hx_Field_left.Hx_Field_key, r.Hx_Field_left.Hx_Field_value, Hx_Obj_haxe_ds_treenode_CreateInstance(r.Hx_Field_left.Hx_Field_right, r.Hx_Field_key, r.Hx_Field_value, r.Hx_Field_right, nil), nil)
			}
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			if hl > hr {
				_hx_tmp_6 = hl
			} else {
				_hx_tmp_6 = hr
			}

			return Hx_Obj_haxe_ds_treenode_CreateInstance(l, k, v, r, ((_hx_tmp_6) + 1))
		}
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_compare(k1 any, k2 any) int32 {
	return Hx_Field_reflect_compare(k1, k2)
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_toString() string {
	if this.Hx_Field_root == nil {
		return "[]"
	} else {
		return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("[", this.Hx_Field_root.VTable.Hx_Field_toString()), "]")
	}
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_clear() {
	this.Hx_Field_root = nil
}

func (this *Hx_Obj_haxe_ds_balancedtree) Hx_Field_size() int32 {
	return Hx_Field_haxe_ds_balancedtree_sizeLoop(this.Hx_Field_root)
}

func Hx_Field_haxe_ds_balancedtree_iteratorLoop(node *Hx_Obj_haxe_ds_treenode, acc HxArray[any]) {
	if node != nil {
		Hx_Field_haxe_ds_balancedtree_iteratorLoop(node.Hx_Field_left, acc)
		_ = (Hx_Array_Push[any](HxMakeArrayView[any](acc), node.Hx_Field_value))
		Hx_Field_haxe_ds_balancedtree_iteratorLoop(node.Hx_Field_right, acc)
	}
}

func Hx_Field_haxe_ds_balancedtree_sizeLoop(node *Hx_Obj_haxe_ds_treenode) int32 {
	if node != nil {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = (Hx_Field_haxe_ds_balancedtree_sizeLoop(node.Hx_Field_left) + 1)
		return (_hx_tmp_0 + Hx_Field_haxe_ds_balancedtree_sizeLoop(node.Hx_Field_right))
	} else {
		return 0
	}
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_treenode() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_treenode() {
}

type Hx_Obj_VTable_haxe_ds_treenode interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_treenode struct {
	VTable           Hx_Obj_VTable_haxe_ds_treenode `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_left    *Hx_Obj_haxe_ds_treenode
	Hx_Field_right   *Hx_Obj_haxe_ds_treenode
	Hx_Field_key     any
	Hx_Field_value   any
	Hx_Field__height int32
}

func Hx_Obj_haxe_ds_treenode_CreateEmptyInstance() *Hx_Obj_haxe_ds_treenode {
	obj := &Hx_Obj_haxe_ds_treenode{}
	obj.VTable = obj
	obj.Hx_Field_left = HxDefault[*Hx_Obj_haxe_ds_treenode]()
	obj.Hx_Field_right = HxDefault[*Hx_Obj_haxe_ds_treenode]()
	obj.Hx_Field_key = HxDefault[any]()
	obj.Hx_Field_value = HxDefault[any]()
	obj.Hx_Field__height = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_ds_treenode_CreateInstance(l *Hx_Obj_haxe_ds_treenode, k any, v any, r *Hx_Obj_haxe_ds_treenode, h__isOptional any) *Hx_Obj_haxe_ds_treenode {
	obj := Hx_Obj_haxe_ds_treenode_CreateEmptyInstance()
	obj.Hx_New(l, k, v, r, h__isOptional)
	return obj
}

func (this *Hx_Obj_haxe_ds_treenode) Hx_New(l *Hx_Obj_haxe_ds_treenode, k any, v any, r *Hx_Obj_haxe_ds_treenode, h__isOptional any) {
	{
		h := HxDefault[int32]()
		_ = h

		if h__isOptional != nil {
			h = HxConvert[int32](h__isOptional)
		} else {
			h = -1
		}
		{
			this.Hx_Field_left = l
			this.Hx_Field_key = k
			this.Hx_Field_value = v
			this.Hx_Field_right = r
			if h == -1 {
				var tmp int32
				_ = tmp
				var _this *Hx_Obj_haxe_ds_treenode
				_ = _this
				_this = this.Hx_Field_left
				var _this1 *Hx_Obj_haxe_ds_treenode
				_ = _this1
				_this1 = this.Hx_Field_right
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				if _this == nil {
					_hx_tmp_0 = 0
				} else {
					_hx_tmp_0 = _this.Hx_Field__height
				}

				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if _this1 == nil {
					_hx_tmp_1 = 0
				} else {
					_hx_tmp_1 = _this1.Hx_Field__height
				}

				if (_hx_tmp_0) > (_hx_tmp_1) {
					var tmp__this_1 *Hx_Obj_haxe_ds_treenode
					_ = tmp__this_1
					tmp__this_1 = this.Hx_Field_left
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					if tmp__this_1 == nil {
						_hx_tmp_2 = 0
					} else {
						_hx_tmp_2 = tmp__this_1.Hx_Field__height
					}

					tmp = _hx_tmp_2
				} else {
					var tmp__this_1 *Hx_Obj_haxe_ds_treenode
					_ = tmp__this_1
					tmp__this_1 = this.Hx_Field_right
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					if tmp__this_1 == nil {
						_hx_tmp_3 = 0
					} else {
						_hx_tmp_3 = tmp__this_1.Hx_Field__height
					}

					tmp = _hx_tmp_3
				}

				this.Hx_Field__height = (tmp + 1)
			} else {
				this.Hx_Field__height = h
			}
		}

	}
}

var Hx_Obj_haxe_ds_treenode_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.TreeNode",
	HxMakeArray[string](),
	HxMakeArray[string]("left", "right", "key", "value", "_height", "get_height", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_treenode_CreateInstance(HxConvert[*Hx_Obj_haxe_ds_treenode](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[*Hx_Obj_haxe_ds_treenode](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_haxe_ds_treenode_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_treenode_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_treenode) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_treenode_RTTI
}

func (this *Hx_Obj_haxe_ds_treenode) Hx_As_Hx_Obj_haxe_ds_treenode() *Hx_Obj_haxe_ds_treenode {
	return this
}

func (this *Hx_Obj_haxe_ds_treenode) Hx_Field_toString() string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if this.Hx_Field_left == nil {
		_hx_tmp_0 = ""
	} else {
		_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_std_string(this.Hx_Field_left.VTable.Hx_Field_toString()), ", ")
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", Hx_Field_std_string(this.Hx_Field_key)), " => ")
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	if this.Hx_Field_right == nil {
		_hx_tmp_2 = ""
	} else {
		_hx_tmp_2 = Hx_Field_go_haxe_hxstring_concat(", ", Hx_Field_std_string(this.Hx_Field_right.VTable.Hx_Field_toString()))
	}

	return Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat((_hx_tmp_0), (Hx_Field_go_haxe_hxstring_concat(_hx_tmp_1, Hx_Field_std_string(this.Hx_Field_value)))), (_hx_tmp_2))
}

func init() {
	hxRegisterArrayConstructor[any]()
}
