package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_exception() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_exception() {
}

type Hx_Obj_VTable_haxe_exception interface {
	Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_get_previous() *Hx_Obj_haxe_exception
	Hx_Field_get_native() any
	Hx_Field_toString() string
	Hx_Field_details() string
	Hx_Field_get_message() string
	Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem]
	Hx_Field_unwrap() any
	Hx_Field_error() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_exception struct {
	VTable            Hx_Obj_VTable_haxe_exception `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field__message string
	Hx_Field__native  any
	Hx_Field__stack   HxArray[Hx_Enum_haxe_stackitem]
}

func Hx_Obj_haxe_exception_CreateEmptyInstance() *Hx_Obj_haxe_exception {
	obj := &Hx_Obj_haxe_exception{}
	obj.VTable = obj
	obj.Hx_Field__message = HxDefault[string]()
	obj.Hx_Field__native = HxDefault[any]()
	obj.Hx_Field__stack = HxDefault[HxArray[Hx_Enum_haxe_stackitem]]()
	return obj
}

func Hx_Obj_haxe_exception_CreateInstance(message string, previous__isOptional any, native__isOptional any) *Hx_Obj_haxe_exception {
	obj := Hx_Obj_haxe_exception_CreateEmptyInstance()
	obj.Hx_New(message, previous__isOptional, native__isOptional)
	return obj
}

func (this *Hx_Obj_haxe_exception) Hx_New(message string, previous__isOptional any, native__isOptional any) {
	{
		previous := HxDefault[*Hx_Obj_haxe_exception]()
		_ = previous

		if previous__isOptional != nil {
			previous = HxConvert[*Hx_Obj_haxe_exception](previous__isOptional)
		} else {
			previous = nil
		}
		native := HxDefault[any]()
		_ = native

		if native__isOptional != nil {
			native = HxConvert[any](native__isOptional)
		} else {
			native = nil
		}
		{
			this.Hx_Field__message = message
			this.Hx_Field__native = native
			var _hx_tmp_0 Hx_Typedef_haxe__nativestacktrace_nativetrace
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_haxe_nativestacktrace_callStack()
			this.Hx_Field__stack = Hx_Field_haxe_nativestacktrace_toHaxe(_hx_tmp_0, nil)
		}

	}
}

var Hx_Obj_haxe_exception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.Exception",
	HxMakeArray[string]("caught", "thrown", "catchValue"),
	HxMakeArray[string]("get_message", "get_stack", "set_stack", "get_previous", "get_native", "_message", "_native", "_stack", "unwrap", "toString", "error", "details"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_exception_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)))
	},
	func() any {
		return Hx_Obj_haxe_exception_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_exception_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "caught":
			return Hx_Field_haxe_exception_caught
		case "thrown":
			return Hx_Field_haxe_exception_thrown
		case "catchValue":
			return Hx_Field_haxe_exception_catchValue
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_exception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_exception_RTTI
}

func (this *Hx_Obj_haxe_exception) Hx_As_Hx_Obj_haxe_exception() *Hx_Obj_haxe_exception {
	return this
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_message() string {
	return this.Hx_Field__message
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_stack() HxArray[Hx_Enum_haxe_stackitem] {
	return this.Hx_Field__stack
}

func (this *Hx_Obj_haxe_exception) Hx_Field_set_stack(stack HxArray[Hx_Enum_haxe_stackitem]) HxArray[Hx_Enum_haxe_stackitem] {
	this.Hx_Field__stack = stack
	var _hx_tmp_0 HxArray[Hx_Enum_haxe_stackitem]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field__stack
	return _hx_tmp_0
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_previous() *Hx_Obj_haxe_exception {
	return nil
}

func (this *Hx_Obj_haxe_exception) Hx_Field_get_native() any {
	return this.Hx_Field__native
}

func (this *Hx_Obj_haxe_exception) Hx_Field_unwrap() any {
	return ((any)(this))
}

func (this *Hx_Obj_haxe_exception) Hx_Field_toString() string {
	return Hx_Field_go_haxe_hxdynamic_stringConcat("", this.VTable.Hx_Field_get_message())
}

func (this *Hx_Obj_haxe_exception) Hx_Field_error() string {
	return this.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_haxe_exception) Error() string {
	return this.VTable.Hx_Field_error()
}

func (this *Hx_Obj_haxe_exception) Hx_Field_details() string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_stringConcat(this.VTable.Hx_Field_toString(), "\n")
	return Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_0, Hx_Field_haxe__callstack_callstack_impl__toString(this.VTable.Hx_Field_get_stack()))
}

func Hx_Field_haxe_exception_caught(value any) *Hx_Obj_haxe_exception {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		_hx_tmp_0 = "null"
	} else {
		_hx_tmp_0 = Hx_Field_std_string(value)
	}

	return Hx_Obj_haxe_exception_CreateInstance(_hx_tmp_0, nil, nil)
}

func Hx_Field_haxe_exception_thrown(value any) *Hx_Obj_haxe_exception {
	var exc any
	_ = exc
	exc = Hx_Field_go_haxe_hxdynamic_toClass(value, "Hx_Obj_haxe_exception")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(exc, nil))) {
		var hx_dyncast_0 any
		_ = hx_dyncast_0
		hx_dyncast_0 = Hx_Field_go_haxe_hxdynamic_castClass(exc, "Hx_Obj_haxe_exception")
		var _hx_tmp_0 *Hx_Obj_haxe_exception
		_ = _hx_tmp_0
		if hx_dyncast_0 != nil {
			_hx_tmp_0 = hx_dyncast_0.(interface{ Hx_As_Hx_Obj_haxe_exception() *Hx_Obj_haxe_exception }).Hx_As_Hx_Obj_haxe_exception()
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0
	}

	return (func() *Hx_Obj_haxe_exception {
		_hx_up := Hx_Obj_haxe_valueexception_CreateInstance(value, nil, nil)
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_haxe_exception
		}
		return nil
	})()
}

func Hx_Field_haxe_exception_catchValue(ex any, goType string) any {
	if goType == "any" {
		{
			var hx_try_return any
			_ = hx_try_return
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
							_ = e
							{
								hx_try_state = 1
								hx_try_return = ex
								return
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					var _hx_tmp_0 any
					_ = _hx_tmp_0
					_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(ex, "unwrap")
					var r any
					_ = r
					r = Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn())
					hx_try_state = 1
					hx_try_return = r
					return
				}
			}()

			return hx_try_return
		}
	}

	var hx_dyncast_1 any
	_ = hx_dyncast_1
	hx_dyncast_1 = Hx_Field_go_haxe_hxdynamic_castClass(ex, "Hx_Obj_haxe_exception")
	var _hx_tmp_1 *Hx_Obj_haxe_exception
	_ = _hx_tmp_1
	if hx_dyncast_1 != nil {
		_hx_tmp_1 = hx_dyncast_1.(interface{ Hx_As_Hx_Obj_haxe_exception() *Hx_Obj_haxe_exception }).Hx_As_Hx_Obj_haxe_exception()
	} else {
		_hx_tmp_1 = nil
	}

	var e *Hx_Obj_haxe_exception
	_ = e
	e = _hx_tmp_1
	var value any
	_ = value
	value = e.VTable.Hx_Field_unwrap()
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = goType
	var name string
	_ = name
	name = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 1, HxNullable[int32]{}))
	var hit any
	_ = hit
	hit = Hx_Field_go_haxe_hxdynamic_toClass(value, name)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(hit, nil))) {
		hit = Hx_Field_go_haxe_hxdynamic_toClass(e, name)
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(hit, nil))) {
		return hit
	} else {
		return value
	}
}

func Hx_ClassInit_Hx_Obj_VTable_haxe__exception_exception_fields_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe__exception_exception_fields_() {
}

type Hx_Obj_VTable_haxe__exception_exception_fields_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe__exception_exception_fields_ struct {
	VTable Hx_Obj_VTable_haxe__exception_exception_fields_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe__exception_exception_fields__CreateEmptyInstance() *Hx_Obj_haxe__exception_exception_fields_ {
	obj := &Hx_Obj_haxe__exception_exception_fields_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe__exception_exception_fields__CreateInstance() *Hx_Obj_haxe__exception_exception_fields_ {
	obj := Hx_Obj_haxe__exception_exception_fields__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe__exception_exception_fields__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe._Exception.Exception_Fields_",
	HxMakeArray[string]("checkException"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe__exception_exception_fields__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe__exception_exception_fields__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe__exception_exception_fields__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "checkException":
			return Hx_Field_haxe__exception_exception_fields__checkException
		}
		return nil
	}
}

func (this *Hx_Obj_haxe__exception_exception_fields_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe__exception_exception_fields__RTTI
}

func (this *Hx_Obj_haxe__exception_exception_fields_) Hx_As_Hx_Obj_haxe__exception_exception_fields_() *Hx_Obj_haxe__exception_exception_fields_ {
	return this
}

func Hx_Field_haxe__exception_exception_fields__checkException(e any) any {
	var thrown any
	_ = thrown
	thrown = Hx_Field_go_haxe_hxdynamic_toClass(e, "Hx_Obj_haxe_exception")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(thrown, nil))) {
		return thrown
	}

	var isGoError bool
	_ = isGoError
	isGoError = false
	var goError error
	_ = goError
	goError, isGoError = e.(error)
	if isGoError {
		return ((any)((func() *Hx_Obj_haxe_exception {
			_hx_up := Hx_Obj_haxe_valueexception_CreateInstance(goError.Error(), nil, goError)
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_haxe_exception
			}
			return nil
		})()))
	}

	if Hx_Field_std_isOfType(e, Hx_Obj_string_RTTI) {
		return ((any)((func() *Hx_Obj_haxe_exception {
			_hx_up := Hx_Obj_haxe_valueexception_CreateInstance(e, nil, nil)
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_haxe_exception
			}
			return nil
		})()))
	}

	return e
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
