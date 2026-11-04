package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds_list() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds_list() {
}

type Hx_Obj_VTable_haxe_ds_list interface {
	Hx_Field_keyValueIterator() *Hx_Obj_haxe_ds__list_listkeyvalueiterator
	Hx_Field_join(sep string) string
	Hx_Field_push(item any)
	Hx_Field_pop() any
	Hx_Field_remove(v any) bool
	Hx_Field_toString() string
	Hx_Field_filter(f func(_hx_carg0 any) bool) *Hx_Obj_haxe_ds_list
	Hx_Field_map(f func(_hx_carg0 any) any) *Hx_Obj_haxe_ds_list
	Hx_Field_add(item any)
	Hx_Field_first() any
	Hx_Field_last() any
	Hx_Field_isEmpty() bool
	Hx_Field_clear()
	Hx_Field_iterator() *Hx_Obj_haxe_ds__list_listiterator
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds_list struct {
	VTable          Hx_Obj_VTable_haxe_ds_list `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_h      *Hx_Obj_haxe_ds__list_listnode
	Hx_Field_q      *Hx_Obj_haxe_ds__list_listnode
	Hx_Field_length int32
}

func Hx_Obj_haxe_ds_list_CreateEmptyInstance() *Hx_Obj_haxe_ds_list {
	obj := &Hx_Obj_haxe_ds_list{}
	obj.VTable = obj
	obj.Hx_Field_h = HxDefault[*Hx_Obj_haxe_ds__list_listnode]()
	obj.Hx_Field_q = HxDefault[*Hx_Obj_haxe_ds__list_listnode]()
	obj.Hx_Field_length = HxDefault[int32]()
	return obj
}

func Hx_Obj_haxe_ds_list_CreateInstance() *Hx_Obj_haxe_ds_list {
	obj := Hx_Obj_haxe_ds_list_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_ds_list) Hx_New() {
	{
		this.Hx_Field_length = 0
	}

}

var Hx_Obj_haxe_ds_list_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds.List",
	HxMakeArray[string](),
	HxMakeArray[string]("h", "q", "length", "add", "push", "first", "last", "pop", "isEmpty", "clear", "remove", "iterator", "keyValueIterator", "toString", "join", "filter", "map"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds_list_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_ds_list_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds_list_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds_list_RTTI
}

func (this *Hx_Obj_haxe_ds_list) Hx_As_Hx_Obj_haxe_ds_list() *Hx_Obj_haxe_ds_list {
	return this
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_add(item any) {
	var x *Hx_Obj_haxe_ds__list_listnode
	_ = x
	x = Hx_Obj_haxe_ds__list_listnode_CreateInstance(item, nil)
	if this.Hx_Field_h == nil {
		this.Hx_Field_h = x
	} else {
		this.Hx_Field_q.Hx_Field_next = x
	}

	this.Hx_Field_q = x
	this.Hx_Field_length++
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_push(item any) {
	var x *Hx_Obj_haxe_ds__list_listnode
	_ = x
	x = Hx_Obj_haxe_ds__list_listnode_CreateInstance(item, this.Hx_Field_h)
	this.Hx_Field_h = x
	if this.Hx_Field_q == nil {
		this.Hx_Field_q = x
	}

	this.Hx_Field_length++
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_first() any {
	if this.Hx_Field_h == nil {
		return nil
	} else {
		return this.Hx_Field_h.Hx_Field_item
	}
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_last() any {
	if this.Hx_Field_q == nil {
		return nil
	} else {
		return this.Hx_Field_q.Hx_Field_item
	}
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_pop() any {
	if this.Hx_Field_h == nil {
		return nil
	}

	var x any
	_ = x
	x = this.Hx_Field_h.Hx_Field_item
	this.Hx_Field_h = this.Hx_Field_h.Hx_Field_next
	if this.Hx_Field_h == nil {
		this.Hx_Field_q = nil
	}

	this.Hx_Field_length--
	return x
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_isEmpty() bool {
	return (this.Hx_Field_h == nil)
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_clear() {
	this.Hx_Field_h = nil
	this.Hx_Field_q = nil
	this.Hx_Field_length = 0
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_remove(v any) bool {
	var prev *Hx_Obj_haxe_ds__list_listnode
	_ = prev
	var l *Hx_Obj_haxe_ds__list_listnode
	_ = l
	l = this.Hx_Field_h
	for l != nil {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l.Hx_Field_item, v))) {
			if prev == nil {
				this.Hx_Field_h = l.Hx_Field_next
			} else {
				prev.Hx_Field_next = l.Hx_Field_next
			}

			if this.Hx_Field_q == l {
				this.Hx_Field_q = prev
			}

			this.Hx_Field_length--
			return true
		}

		prev = l
		l = l.Hx_Field_next
	}

	return false
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_iterator() *Hx_Obj_haxe_ds__list_listiterator {
	return Hx_Obj_haxe_ds__list_listiterator_CreateInstance(this.Hx_Field_h)
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_keyValueIterator() *Hx_Obj_haxe_ds__list_listkeyvalueiterator {
	return Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateInstance(this.Hx_Field_h)
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_toString() string {
	var s *Hx_Obj_stringbuf
	_ = s
	s = Hx_Obj_stringbuf_CreateInstance()
	var first bool
	_ = first
	first = true
	var l *Hx_Obj_haxe_ds__list_listnode
	_ = l
	l = this.Hx_Field_h
	func() func(x string) {
		var _hx_closure_src_196 func(x any)
		_ = _hx_closure_src_196
		_hx_closure_src_196 = s.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_196(((any)(x)))
		}
	}()("{")
	for l != nil {
		if first {
			first = false
		} else {
			func() func(x string) {
				var _hx_closure_src_197 func(x any)
				_ = _hx_closure_src_197
				_hx_closure_src_197 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_197(((any)(x)))
				}
			}()(", ")
		}

		func() func(x string) {
			var _hx_closure_src_198 func(x any)
			_ = _hx_closure_src_198
			_hx_closure_src_198 = s.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_198(((any)(x)))
			}
		}()(Hx_Field_std_string(l.Hx_Field_item))
		l = l.Hx_Field_next
	}

	func() func(x string) {
		var _hx_closure_src_199 func(x any)
		_ = _hx_closure_src_199
		_hx_closure_src_199 = s.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_199(((any)(x)))
		}
	}()("}")
	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_join(sep string) string {
	var s *Hx_Obj_stringbuf
	_ = s
	s = Hx_Obj_stringbuf_CreateInstance()
	var first bool
	_ = first
	first = true
	var l *Hx_Obj_haxe_ds__list_listnode
	_ = l
	l = this.Hx_Field_h
	for l != nil {
		if first {
			first = false
		} else {
			func() func(x string) {
				var _hx_closure_src_200 func(x any)
				_ = _hx_closure_src_200
				_hx_closure_src_200 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_200(((any)(x)))
				}
			}()(sep)
		}

		s.VTable.Hx_Field_add(l.Hx_Field_item)
		l = l.Hx_Field_next
	}

	return s.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_filter(f func(_hx_carg0 any) bool) *Hx_Obj_haxe_ds_list {
	var l2 *Hx_Obj_haxe_ds_list
	_ = l2
	l2 = Hx_Obj_haxe_ds_list_CreateInstance()
	var l *Hx_Obj_haxe_ds__list_listnode
	_ = l
	l = this.Hx_Field_h
	for l != nil {
		var v any
		_ = v
		v = l.Hx_Field_item
		l = l.Hx_Field_next
		if f(v) {
			l2.VTable.Hx_Field_add(v)
		}
	}

	return l2
}

func (this *Hx_Obj_haxe_ds_list) Hx_Field_map(f func(_hx_carg0 any) any) *Hx_Obj_haxe_ds_list {
	var b *Hx_Obj_haxe_ds_list
	_ = b
	b = Hx_Obj_haxe_ds_list_CreateInstance()
	var l *Hx_Obj_haxe_ds__list_listnode
	_ = l
	l = this.Hx_Field_h
	for l != nil {
		var v any
		_ = v
		v = l.Hx_Field_item
		l = l.Hx_Field_next
		b.VTable.Hx_Field_add(f(v))
	}

	return b
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__list_listnode() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__list_listnode() {
}

type Hx_Obj_VTable_haxe_ds__list_listnode interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__list_listnode struct {
	VTable        Hx_Obj_VTable_haxe_ds__list_listnode `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_item any
	Hx_Field_next *Hx_Obj_haxe_ds__list_listnode
}

func Hx_Obj_haxe_ds__list_listnode_CreateEmptyInstance() *Hx_Obj_haxe_ds__list_listnode {
	obj := &Hx_Obj_haxe_ds__list_listnode{}
	obj.VTable = obj
	obj.Hx_Field_item = HxDefault[any]()
	obj.Hx_Field_next = HxDefault[*Hx_Obj_haxe_ds__list_listnode]()
	return obj
}

func Hx_Obj_haxe_ds__list_listnode_CreateInstance(item any, next *Hx_Obj_haxe_ds__list_listnode) *Hx_Obj_haxe_ds__list_listnode {
	obj := Hx_Obj_haxe_ds__list_listnode_CreateEmptyInstance()
	obj.Hx_New(item, next)
	return obj
}

func (this *Hx_Obj_haxe_ds__list_listnode) Hx_New(item any, next *Hx_Obj_haxe_ds__list_listnode) {
	{
		this.Hx_Field_item = item
		this.Hx_Field_next = next
	}

}

var Hx_Obj_haxe_ds__list_listnode_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._List.ListNode",
	HxMakeArray[string]("create"),
	HxMakeArray[string]("item", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__list_listnode_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[*Hx_Obj_haxe_ds__list_listnode](params.Get(1)))
	},
	func() any {
		return Hx_Obj_haxe_ds__list_listnode_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__list_listnode_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__list_listnode) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__list_listnode_RTTI
}

func (this *Hx_Obj_haxe_ds__list_listnode) Hx_As_Hx_Obj_haxe_ds__list_listnode() *Hx_Obj_haxe_ds__list_listnode {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__list_listiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__list_listiterator() {
}

type Hx_Obj_VTable_haxe_ds__list_listiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__list_listiterator struct {
	VTable        Hx_Obj_VTable_haxe_ds__list_listiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_head *Hx_Obj_haxe_ds__list_listnode
}

func Hx_Obj_haxe_ds__list_listiterator_CreateEmptyInstance() *Hx_Obj_haxe_ds__list_listiterator {
	obj := &Hx_Obj_haxe_ds__list_listiterator{}
	obj.VTable = obj
	obj.Hx_Field_head = HxDefault[*Hx_Obj_haxe_ds__list_listnode]()
	return obj
}

func Hx_Obj_haxe_ds__list_listiterator_CreateInstance(head *Hx_Obj_haxe_ds__list_listnode) *Hx_Obj_haxe_ds__list_listiterator {
	obj := Hx_Obj_haxe_ds__list_listiterator_CreateEmptyInstance()
	obj.Hx_New(head)
	return obj
}

func (this *Hx_Obj_haxe_ds__list_listiterator) Hx_New(head *Hx_Obj_haxe_ds__list_listnode) {
	{
		this.Hx_Field_head = head
	}

}

var Hx_Obj_haxe_ds__list_listiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._List.ListIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("head", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__list_listiterator_CreateInstance(HxConvert[*Hx_Obj_haxe_ds__list_listnode](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_ds__list_listiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__list_listiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__list_listiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__list_listiterator_RTTI
}

func (this *Hx_Obj_haxe_ds__list_listiterator) Hx_As_Hx_Obj_haxe_ds__list_listiterator() *Hx_Obj_haxe_ds__list_listiterator {
	return this
}

func (this *Hx_Obj_haxe_ds__list_listiterator) Hx_Field_hasNext() bool {
	return (this.Hx_Field_head != nil)
}

func (this *Hx_Obj_haxe_ds__list_listiterator) Hx_Field_next() any {
	var val any
	_ = val
	val = this.Hx_Field_head.Hx_Field_item
	this.Hx_Field_head = this.Hx_Field_head.Hx_Field_next
	return val
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_ds__list_listkeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_ds__list_listkeyvalueiterator() {
}

type Hx_Obj_VTable_haxe_ds__list_listkeyvalueiterator interface {
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_ds__list_listkeyvalueiterator struct {
	VTable        Hx_Obj_VTable_haxe_ds__list_listkeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_idx  int32
	Hx_Field_head *Hx_Obj_haxe_ds__list_listnode
}

func Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_haxe_ds__list_listkeyvalueiterator {
	obj := &Hx_Obj_haxe_ds__list_listkeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_idx = HxDefault[int32]()
	obj.Hx_Field_head = HxDefault[*Hx_Obj_haxe_ds__list_listnode]()
	return obj
}

func Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateInstance(head *Hx_Obj_haxe_ds__list_listnode) *Hx_Obj_haxe_ds__list_listkeyvalueiterator {
	obj := Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(head)
	return obj
}

func (this *Hx_Obj_haxe_ds__list_listkeyvalueiterator) Hx_New(head *Hx_Obj_haxe_ds__list_listnode) {
	{
		this.Hx_Field_head = head
		this.Hx_Field_idx = 0
	}

}

var Hx_Obj_haxe_ds__list_listkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.ds._List.ListKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("idx", "head", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateInstance(HxConvert[*Hx_Obj_haxe_ds__list_listnode](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_ds__list_listkeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_ds__list_listkeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_ds__list_listkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_ds__list_listkeyvalueiterator_RTTI
}

func (this *Hx_Obj_haxe_ds__list_listkeyvalueiterator) Hx_As_Hx_Obj_haxe_ds__list_listkeyvalueiterator() *Hx_Obj_haxe_ds__list_listkeyvalueiterator {
	return this
}

func (this *Hx_Obj_haxe_ds__list_listkeyvalueiterator) Hx_Field_hasNext() bool {
	return (this.Hx_Field_head != nil)
}

func (this *Hx_Obj_haxe_ds__list_listkeyvalueiterator) Hx_Field_next() any {
	var val any
	_ = val
	val = this.Hx_Field_head.Hx_Field_item
	this.Hx_Field_head = this.Hx_Field_head.Hx_Field_next
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = val
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_idx
	this.Hx_Field_idx = (this.Hx_Field_idx + 1)
	return any(map[string]any{"value": _hx_tmp_0, "key": ((any)(_hx_tmp_1))})
}
