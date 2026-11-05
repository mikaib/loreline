package core

type Hx_Enum_haxe_stackitem interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_haxe_stackitem()
}

type Hx_Enum_haxe_stackitem_Module struct {
	Hx_Field_m string
}

func (this Hx_Enum_haxe_stackitem_Module) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_stackitem_Module) M_Hx_Enum_haxe_stackitem()               {}
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_stackitem_RTTI
}
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_m))
}
func (this Hx_Enum_haxe_stackitem_Module) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_m)
	default:
		return nil
	}
}

type Hx_Enum_haxe_stackitem_Method struct {
	Hx_Field_classname HxNullable[string]
	Hx_Field_method    string
}

func (this Hx_Enum_haxe_stackitem_Method) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_stackitem_Method) M_Hx_Enum_haxe_stackitem()               {}
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_stackitem_RTTI
}
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_classname), any(this.Hx_Field_method))
}
func (this Hx_Enum_haxe_stackitem_Method) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_classname)
	case 1:
		return any(this.Hx_Field_method)
	default:
		return nil
	}
}

type Hx_Enum_haxe_stackitem_LocalFunction struct {
	Hx_Field_v HxNullable[int32]
}

func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_stackitem_LocalFunction) M_Hx_Enum_haxe_stackitem()               {}
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumIndex() int32               { return 4 }
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_stackitem_RTTI
}
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_v))
}
func (this Hx_Enum_haxe_stackitem_LocalFunction) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_v)
	default:
		return nil
	}
}

type Hx_Enum_haxe_stackitem_FilePos struct {
	Hx_Field_s      Hx_Enum_haxe_stackitem
	Hx_Field_file   string
	Hx_Field_line   int32
	Hx_Field_column HxNullable[int32]
}

func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_stackitem_FilePos) M_Hx_Enum_haxe_stackitem()               {}
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_stackitem_RTTI
}
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s), any(this.Hx_Field_file), any(this.Hx_Field_line), any(this.Hx_Field_column))
}
func (this Hx_Enum_haxe_stackitem_FilePos) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	case 1:
		return any(this.Hx_Field_file)
	case 2:
		return any(this.Hx_Field_line)
	case 3:
		return any(this.Hx_Field_column)
	default:
		return nil
	}
}

type Hx_Enum_haxe_stackitem_CFunction struct{}

func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_haxe_stackitem_CFunction) M_Hx_Enum_haxe_stackitem()               {}
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_haxe_stackitem_RTTI
}
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_haxe_stackitem_CFunction) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_haxe_stackitem_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"haxe.StackItem",
	HxMakeArray[string]("CFunction", "Module", "FilePos", "Method", "LocalFunction"),
	HxMakeArray[int32](0, 1, 4, 2, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_haxe_stackitem_CFunction{}
		case 1:
			return Hx_Enum_haxe_stackitem_Module{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 2:
			return Hx_Enum_haxe_stackitem_FilePos{params.Get_Dyn(0).(Hx_Enum_haxe_stackitem), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1)))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(2)))), HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(3)))), Valid: params.Get_Dyn(3) != nil}}
		case 3:
			return Hx_Enum_haxe_stackitem_Method{HxNullable[string]{Value: ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), Valid: params.Get_Dyn(0) != nil}, ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1))))}
		case 4:
			return Hx_Enum_haxe_stackitem_LocalFunction{HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(0)))), Valid: params.Get_Dyn(0) != nil}}
		default:
			return nil
		}
	},
	nil,
)

var Hx_Abstract_haxe_callstack_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.CallStack",
	HxMakeArray[string]("length", "get_length", "callStack", "exceptionStack", "toString", "subtract", "copy", "get", "asArray", "equalItems", "exceptionToString", "itemToString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return nil
	},
	func() any {
		return nil
	},
	nil,
	func(name string) any { return nil },
)

func Hx_ClassInit_Hx_Obj_VTable_haxe__callstack_callstack_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__callstack_callstack_impl_() {
}

type Hx_Obj_VTable_haxe__callstack_callstack_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__callstack_callstack_impl_ struct {
	VTable Hx_Obj_VTable_haxe__callstack_callstack_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__callstack_callstack_impl__CreateEmptyInstance() *Hx_Obj_haxe__callstack_callstack_impl_ {
	obj := &Hx_Obj_haxe__callstack_callstack_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__callstack_callstack_impl__CreateInstance() *Hx_Obj_haxe__callstack_callstack_impl_ {
	obj := Hx_Obj_haxe__callstack_callstack_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__callstack_callstack_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._CallStack.CallStack_Impl_",
	HxMakeArray[string]("get_length", "callStack", "exceptionStack", "toString", "subtract", "copy", "get", "asArray", "equalItems", "exceptionToString", "itemToString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__callstack_callstack_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__callstack_callstack_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__callstack_callstack_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "get_length":
			return Hx_Field_haxe__callstack_callstack_impl__get_length
		case "callStack":
			return Hx_Field_haxe__callstack_callstack_impl__callStack
		case "exceptionStack":
			return Hx_Field_haxe__callstack_callstack_impl__exceptionStack
		case "toString":
			return Hx_Field_haxe__callstack_callstack_impl__toString
		case "subtract":
			return Hx_Field_haxe__callstack_callstack_impl__subtract
		case "copy":
			return Hx_Field_haxe__callstack_callstack_impl__copy
		case "get":
			return Hx_Field_haxe__callstack_callstack_impl__get
		case "asArray":
			return Hx_Field_haxe__callstack_callstack_impl__asArray
		case "equalItems":
			return Hx_Field_haxe__callstack_callstack_impl__equalItems
		case "exceptionToString":
			return Hx_Field_haxe__callstack_callstack_impl__exceptionToString
		case "itemToString":
			return Hx_Field_haxe__callstack_callstack_impl__itemToString
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__callstack_callstack_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__callstack_callstack_impl__RTTI
}

func (this *Hx_Obj_haxe__callstack_callstack_impl_) Hx_As_Hx_Obj_haxe__callstack_callstack_impl_() *Hx_Obj_haxe__callstack_callstack_impl_ {
	return this
}

func Hx_Field_haxe__callstack_callstack_impl__get_length(this HxArray[Hx_Enum_haxe_stackitem]) int32 {
	return this.Len()
}

func Hx_Field_haxe__callstack_callstack_impl__callStack() HxArray[Hx_Enum_haxe_stackitem] {
	var _hx_tmp_0 Hx_Typedef_haxe__nativestacktrace_nativetrace
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_haxe_nativestacktrace_callStack()
	return Hx_Field_haxe_nativestacktrace_toHaxe(_hx_tmp_0, nil)
}

func Hx_Field_haxe__callstack_callstack_impl__exceptionStack(fullStack__isOptional any) HxArray[Hx_Enum_haxe_stackitem] {
	fullStack := HxDefault[bool]()
	_ = fullStack

	if fullStack__isOptional != nil {
		fullStack = HxConvert[bool](fullStack__isOptional)
	} else {
		fullStack = false
	}
	{
		var _hx_tmp_0 Hx_Typedef_haxe__nativestacktrace_nativetrace
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_haxe_nativestacktrace_exceptionStack()
		var eStack HxArray[Hx_Enum_haxe_stackitem]
		_ = eStack
		eStack = Hx_Field_haxe_nativestacktrace_toHaxe(_hx_tmp_0, nil)
		var _hx_tmp_1 HxArray[Hx_Enum_haxe_stackitem]
		_ = _hx_tmp_1
		if fullStack {
			_hx_tmp_1 = eStack
		} else {
			var _hx_tmp_2 HxArray[Hx_Enum_haxe_stackitem]
			_ = _hx_tmp_2
			_hx_tmp_2 = eStack
			var _hx_tmp_3 HxArray[Hx_Enum_haxe_stackitem]
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = Hx_Field_haxe__callstack_callstack_impl__subtract(_hx_tmp_3, Hx_Field_haxe__callstack_callstack_impl__callStack())
		}

		return _hx_tmp_1
	}
}

func Hx_Field_haxe__callstack_callstack_impl__toString(stack HxArray[Hx_Enum_haxe_stackitem]) string {
	var b *Hx_Obj_stringbuf
	_ = b
	b = Hx_Obj_stringbuf_CreateInstance()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[Hx_Enum_haxe_stackitem]
		_ = _g1
		_g1 = stack
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var s Hx_Enum_haxe_stackitem
			_ = s
			s = _g1.Get(_g)
			_g++
			func() func(x string) {
				var _hx_closure_src_33 func(x any)
				_ = _hx_closure_src_33
				_hx_closure_src_33 = b.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_33(((any)(x)))
				}
			}()("\nCalled from ")
			Hx_Field_haxe__callstack_callstack_impl__itemToString(b, s)
		}
	}

	return b.VTable.Hx_Field_toString()
}

func Hx_Field_haxe__callstack_callstack_impl__subtract(this HxArray[Hx_Enum_haxe_stackitem], stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem] {
	var startIndex int32
	_ = startIndex
	startIndex = -1
	var i int32
	_ = i
	i = -1
	for {
		i = (i + 1)
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < this.Len()) {
			break
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = stack.Len()
			for _g < _g1 {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_g = (_g + 1)
				var j int32
				_ = j
				j = _hx_tmp_1
				if Hx_Field_haxe__callstack_callstack_impl__equalItems(this.Get(i), stack.Get(j)) {
					if startIndex < 0 {
						startIndex = i
					}

					i++
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = i
					if _hx_tmp_2 >= this.Len() {
						break
					}
				} else {
					startIndex = -1
				}
			}
		}

		if startIndex >= 0 {
			break
		}
	}

	if startIndex >= 0 {
		var _hx_tmp_3 HxArray[Hx_Enum_haxe_stackitem]
		_ = _hx_tmp_3
		_hx_tmp_3 = this
		return (Hx_Array_Slice[Hx_Enum_haxe_stackitem](_hx_tmp_3, 0, HxNullable[int32]{Value: startIndex, Valid: true}))
	} else {
		return this
	}
}

func Hx_Field_haxe__callstack_callstack_impl__copy(this HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem] {
	return (Hx_Array_Copy[Hx_Enum_haxe_stackitem](this))
}

func Hx_Field_haxe__callstack_callstack_impl__get(this HxArray[Hx_Enum_haxe_stackitem], index int32) Hx_Enum_haxe_stackitem {
	return this.Get(index)
}

func Hx_Field_haxe__callstack_callstack_impl__asArray(this HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem] {
	return this
}

func Hx_Field_haxe__callstack_callstack_impl__equalItems(item1 Hx_Enum_haxe_stackitem, item2 Hx_Enum_haxe_stackitem) bool {
	if item1 == nil {
		return (item2 == nil)
	} else {
		switch item1.Hx_Field_enumIndex() {
		case 0:
			{
				if item2 == nil {
					return false
				} else {
					return (item2.Hx_Field_enumIndex() == 0)
				}
			}

		case 1:
			{
				if item2 == nil {
					return false
				} else {
					if item2.Hx_Field_enumIndex() == 1 {
						var m2 string
						_ = m2
						m2 = item2.(Hx_Enum_haxe_stackitem_Module).Hx_Field_m
						var m1 string
						_ = m1
						m1 = item1.(Hx_Enum_haxe_stackitem_Module).Hx_Field_m
						return (m1 == m2)
					} else {
						return false
					}
				}
			}

		case 2:
			{
				if item2 == nil {
					return false
				} else {
					if item2.Hx_Field_enumIndex() == 2 {
						var item21 Hx_Enum_haxe_stackitem
						_ = item21
						item21 = item2.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_s
						var file2 string
						_ = file2
						file2 = item2.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_file
						var line2 int32
						_ = line2
						line2 = item2.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_line
						var col2 HxNullable[int32]
						_ = col2
						col2 = item2.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_column
						var col1 HxNullable[int32]
						_ = col1
						col1 = item1.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_column
						var line1 int32
						_ = line1
						line1 = item1.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_line
						var file1 string
						_ = file1
						file1 = item1.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_file
						var item11 Hx_Enum_haxe_stackitem
						_ = item11
						item11 = item1.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_s
						var _hx_tmp_0 bool
						_ = _hx_tmp_0
						var _hx_tmp_1 bool
						_ = _hx_tmp_1
						if (file1 == file2) && (line1 == line2) {
							var _hx_tmp_2 bool
							_ = _hx_tmp_2
							var _hx_tmp_3 bool
							_ = _hx_tmp_3
							_hx_tmp_3 = (col1).Valid
							if _hx_tmp_3 == (col2).Valid {
								var _hx_tmp_4 bool
								_ = _hx_tmp_4
								if (col1).Valid == false {
									_hx_tmp_4 = true
								} else {
									var _hx_tmp_5 int32
									_ = _hx_tmp_5
									_hx_tmp_5 = (col1).Value
									var _hx_tmp_6 int32
									_ = _hx_tmp_6
									_hx_tmp_6 = _hx_tmp_5
									_hx_tmp_4 = (_hx_tmp_6 == (col2).Value)
								}

								_hx_tmp_2 = _hx_tmp_4
							} else {
								_hx_tmp_2 = false
							}

							_hx_tmp_1 = (_hx_tmp_2)
						} else {
							_hx_tmp_1 = false
						}

						if _hx_tmp_1 {
							_hx_tmp_0 = Hx_Field_haxe__callstack_callstack_impl__equalItems(item11, item21)
						} else {
							_hx_tmp_0 = false
						}

						return _hx_tmp_0
					} else {
						return false
					}
				}
			}

		case 3:
			{
				if item2 == nil {
					return false
				} else {
					if item2.Hx_Field_enumIndex() == 3 {
						var class2 HxNullable[string]
						_ = class2
						class2 = item2.(Hx_Enum_haxe_stackitem_Method).Hx_Field_classname
						var method2 string
						_ = method2
						method2 = item2.(Hx_Enum_haxe_stackitem_Method).Hx_Field_method
						var method1 string
						_ = method1
						method1 = item1.(Hx_Enum_haxe_stackitem_Method).Hx_Field_method
						var class1 HxNullable[string]
						_ = class1
						class1 = item1.(Hx_Enum_haxe_stackitem_Method).Hx_Field_classname
						var _hx_tmp_7 bool
						_ = _hx_tmp_7
						var _hx_tmp_8 bool
						_ = _hx_tmp_8
						var _hx_tmp_9 bool
						_ = _hx_tmp_9
						_hx_tmp_9 = (class1).Valid
						if _hx_tmp_9 == (class2).Valid {
							var _hx_tmp_10 bool
							_ = _hx_tmp_10
							if (class1).Valid == false {
								_hx_tmp_10 = true
							} else {
								var _hx_tmp_11 string
								_ = _hx_tmp_11
								_hx_tmp_11 = (class1).Value
								var _hx_tmp_12 string
								_ = _hx_tmp_12
								_hx_tmp_12 = _hx_tmp_11
								_hx_tmp_10 = (_hx_tmp_12 == (class2).Value)
							}

							_hx_tmp_8 = _hx_tmp_10
						} else {
							_hx_tmp_8 = false
						}

						if _hx_tmp_8 {
							_hx_tmp_7 = (method1 == method2)
						} else {
							_hx_tmp_7 = false
						}

						return _hx_tmp_7
					} else {
						return false
					}
				}
			}

		case 4:
			{
				if item2 == nil {
					return false
				} else {
					if item2.Hx_Field_enumIndex() == 4 {
						var v2 HxNullable[int32]
						_ = v2
						v2 = item2.(Hx_Enum_haxe_stackitem_LocalFunction).Hx_Field_v
						var v1 HxNullable[int32]
						_ = v1
						v1 = item1.(Hx_Enum_haxe_stackitem_LocalFunction).Hx_Field_v
						var _hx_tmp_13 bool
						_ = _hx_tmp_13
						var _hx_tmp_14 bool
						_ = _hx_tmp_14
						_hx_tmp_14 = (v1).Valid
						if _hx_tmp_14 == (v2).Valid {
							var _hx_tmp_15 bool
							_ = _hx_tmp_15
							if (v1).Valid == false {
								_hx_tmp_15 = true
							} else {
								var _hx_tmp_16 int32
								_ = _hx_tmp_16
								_hx_tmp_16 = (v1).Value
								var _hx_tmp_17 int32
								_ = _hx_tmp_17
								_hx_tmp_17 = _hx_tmp_16
								_hx_tmp_15 = (_hx_tmp_17 == (v2).Value)
							}

							_hx_tmp_13 = _hx_tmp_15
						} else {
							_hx_tmp_13 = false
						}

						return (_hx_tmp_13)
					} else {
						return false
					}
				}
			}

		default:
			{
				panic(Hx_Field_haxe_exception_thrown("unreachable"))
			}
		}
	}
}

func Hx_Field_haxe__callstack_callstack_impl__exceptionToString(e *Hx_Obj_haxe_exception) string {
	if e.VTable.Hx_Field_get_previous() == nil {
		var tmp string
		_ = tmp
		tmp = Hx_Field_go_haxe_hxstring_concat("Exception: ", e.VTable.Hx_Field_toString())
		var tmp1 HxArray[Hx_Enum_haxe_stackitem]
		_ = tmp1
		tmp1 = e.VTable.Hx_Field_get_stack()
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = tmp
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		if tmp1 == nil {
			_hx_tmp_1 = "null"
		} else {
			_hx_tmp_1 = Hx_Field_haxe__callstack_callstack_impl__toString(tmp1)
		}

		return Hx_Field_go_haxe_hxstring_concat(_hx_tmp_0, (_hx_tmp_1))
	}

	var result string
	_ = result
	result = ""
	var e1 *Hx_Obj_haxe_exception
	_ = e1
	e1 = e
	var prev *Hx_Obj_haxe_exception
	_ = prev
	for e1 != nil {
		if prev == nil {
			var result1 string
			_ = result1
			result1 = Hx_Field_go_haxe_hxstring_concat("Exception: ", e1.VTable.Hx_Field_get_message())
			var tmp HxArray[Hx_Enum_haxe_stackitem]
			_ = tmp
			tmp = e1.VTable.Hx_Field_get_stack()
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = result1
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			if tmp == nil {
				_hx_tmp_3 = "null"
			} else {
				_hx_tmp_3 = Hx_Field_haxe__callstack_callstack_impl__toString(tmp)
			}

			result = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_2, (_hx_tmp_3)), result)
		} else {
			var _hx_tmp_4 HxArray[Hx_Enum_haxe_stackitem]
			_ = _hx_tmp_4
			_hx_tmp_4 = e1.VTable.Hx_Field_get_stack()
			var prevStack HxArray[Hx_Enum_haxe_stackitem]
			_ = prevStack
			prevStack = Hx_Field_haxe__callstack_callstack_impl__subtract(_hx_tmp_4, prev.VTable.Hx_Field_get_stack())
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Field_go_haxe_hxstring_concat("Exception: ", e1.VTable.Hx_Field_get_message())
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			if prevStack == nil {
				_hx_tmp_6 = "null"
			} else {
				_hx_tmp_6 = Hx_Field_haxe__callstack_callstack_impl__toString(prevStack)
			}

			result = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_5, (_hx_tmp_6)), "\n\nNext "), result)
		}

		prev = e1
		e1 = e1.VTable.Hx_Field_get_previous()
	}

	return result
}

func Hx_Field_haxe__callstack_callstack_impl__itemToString(b *Hx_Obj_stringbuf, s Hx_Enum_haxe_stackitem) {
	switch s.Hx_Field_enumIndex() {
	case 0:
		{
			func() func(x string) {
				var _hx_closure_src_34 func(x any)
				_ = _hx_closure_src_34
				_hx_closure_src_34 = b.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_34(((any)(x)))
				}
			}()("a C function")
		}

	case 1:
		{
			var m string
			_ = m
			m = s.(Hx_Enum_haxe_stackitem_Module).Hx_Field_m
			{
				func() func(x string) {
					var _hx_closure_src_35 func(x any)
					_ = _hx_closure_src_35
					_hx_closure_src_35 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_35(((any)(x)))
					}
				}()("module ")
				func() func(x string) {
					var _hx_closure_src_36 func(x any)
					_ = _hx_closure_src_36
					_hx_closure_src_36 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_36(((any)(x)))
					}
				}()(m)
			}
		}

	case 2:
		{
			var s1 Hx_Enum_haxe_stackitem
			_ = s1
			s1 = s.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_s
			var file string
			_ = file
			file = s.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_file
			var line int32
			_ = line
			line = s.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_line
			var col HxNullable[int32]
			_ = col
			col = s.(Hx_Enum_haxe_stackitem_FilePos).Hx_Field_column
			{
				if s1 != nil {
					Hx_Field_haxe__callstack_callstack_impl__itemToString(b, s1)
					func() func(x string) {
						var _hx_closure_src_37 func(x any)
						_ = _hx_closure_src_37
						_hx_closure_src_37 = b.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_37(((any)(x)))
						}
					}()(" (")
				}

				func() func(x string) {
					var _hx_closure_src_38 func(x any)
					_ = _hx_closure_src_38
					_hx_closure_src_38 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_38(((any)(x)))
					}
				}()(file)
				func() func(x string) {
					var _hx_closure_src_39 func(x any)
					_ = _hx_closure_src_39
					_hx_closure_src_39 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_39(((any)(x)))
					}
				}()(" line ")
				func() func(x int32) {
					var _hx_closure_src_40 func(x any)
					_ = _hx_closure_src_40
					_hx_closure_src_40 = b.VTable.Hx_Field_add
					return func(x int32) {
						_hx_closure_src_40(((any)(x)))
					}
				}()(line)
				if (col).Valid != false {
					func() func(x string) {
						var _hx_closure_src_41 func(x any)
						_ = _hx_closure_src_41
						_hx_closure_src_41 = b.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_41(((any)(x)))
						}
					}()(" column ")
					func() func(x HxNullable[int32]) {
						var _hx_closure_src_42 func(x any)
						_ = _hx_closure_src_42
						_hx_closure_src_42 = b.VTable.Hx_Field_add
						return func(x HxNullable[int32]) {
							var hx_nullable_f0 HxNullable[int32]
							_ = hx_nullable_f0
							hx_nullable_f0 = x
							var _hx_tmp_0 any
							_ = _hx_tmp_0
							if (hx_nullable_f0).Valid == true {
								_hx_tmp_0 = (hx_nullable_f0).Value
							} else {
								_hx_tmp_0 = nil
							}

							_hx_closure_src_42(_hx_tmp_0)
						}
					}()(col)
				}

				if s1 != nil {
					func() func(x string) {
						var _hx_closure_src_43 func(x any)
						_ = _hx_closure_src_43
						_hx_closure_src_43 = b.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_43(((any)(x)))
						}
					}()(")")
				}
			}
		}

	case 3:
		{
			var cname HxNullable[string]
			_ = cname
			cname = s.(Hx_Enum_haxe_stackitem_Method).Hx_Field_classname
			var meth string
			_ = meth
			meth = s.(Hx_Enum_haxe_stackitem_Method).Hx_Field_method
			{
				var _hx_tmp_2 HxNullable[string]
				_ = _hx_tmp_2
				if (cname).Valid == false {
					_hx_tmp_2 = HxNullableString("<unknown>")
				} else {
					_hx_tmp_2 = cname
				}

				func() func(x HxNullable[string]) {
					var _hx_closure_src_44 func(x any)
					_ = _hx_closure_src_44
					_hx_closure_src_44 = b.VTable.Hx_Field_add
					return func(x HxNullable[string]) {
						var hx_nullable_f1 HxNullable[string]
						_ = hx_nullable_f1
						hx_nullable_f1 = x
						var _hx_tmp_1 any
						_ = _hx_tmp_1
						if (hx_nullable_f1).Valid == true {
							_hx_tmp_1 = (hx_nullable_f1).Value
						} else {
							_hx_tmp_1 = nil
						}

						_hx_closure_src_44(_hx_tmp_1)
					}
				}()(_hx_tmp_2)
				func() func(x string) {
					var _hx_closure_src_45 func(x any)
					_ = _hx_closure_src_45
					_hx_closure_src_45 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_45(((any)(x)))
					}
				}()(".")
				func() func(x string) {
					var _hx_closure_src_46 func(x any)
					_ = _hx_closure_src_46
					_hx_closure_src_46 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_46(((any)(x)))
					}
				}()(meth)
			}
		}

	case 4:
		{
			var n HxNullable[int32]
			_ = n
			n = s.(Hx_Enum_haxe_stackitem_LocalFunction).Hx_Field_v
			{
				func() func(x string) {
					var _hx_closure_src_47 func(x any)
					_ = _hx_closure_src_47
					_hx_closure_src_47 = b.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_47(((any)(x)))
					}
				}()("local function #")
				func() func(x HxNullable[int32]) {
					var _hx_closure_src_48 func(x any)
					_ = _hx_closure_src_48
					_hx_closure_src_48 = b.VTable.Hx_Field_add
					return func(x HxNullable[int32]) {
						var hx_nullable_f2 HxNullable[int32]
						_ = hx_nullable_f2
						hx_nullable_f2 = x
						var _hx_tmp_3 any
						_ = _hx_tmp_3
						if (hx_nullable_f2).Valid == true {
							_hx_tmp_3 = (hx_nullable_f2).Value
						} else {
							_hx_tmp_3 = nil
						}

						_hx_closure_src_48(_hx_tmp_3)
					}
				}()(n)
			}
		}

	}
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
