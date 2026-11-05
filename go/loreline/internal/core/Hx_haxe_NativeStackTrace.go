package core

import "runtime"
import "unicode/utf8"

type Hx_Typedef_haxe__nativestacktrace_nativetrace = HxArray[Hx_Enum_haxe_stackitem]

func Hx_ClassInit_Hx_Obj_VTable_haxe_nativestacktrace() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_nativestacktrace() {
}

type Hx_Obj_VTable_haxe_nativestacktrace interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_nativestacktrace struct {
	VTable Hx_Obj_VTable_haxe_nativestacktrace `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_nativestacktrace_CreateEmptyInstance() *Hx_Obj_haxe_nativestacktrace {
	obj := &Hx_Obj_haxe_nativestacktrace{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_nativestacktrace_CreateInstance() *Hx_Obj_haxe_nativestacktrace {
	obj := Hx_Obj_haxe_nativestacktrace_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_nativestacktrace_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.NativeStackTrace",
	HxMakeArray[string]("saveStack", "callStack", "exceptionStack", "toHaxe", "isHaxeCallStackFrame", "isConstructorFrame", "capture", "createMethod", "stripClosureSuffix", "isClosureSegment"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_nativestacktrace_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_nativestacktrace_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_nativestacktrace_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "saveStack":
			return Hx_Field_haxe_nativestacktrace_saveStack
		case "callStack":
			return Hx_Field_haxe_nativestacktrace_callStack
		case "exceptionStack":
			return Hx_Field_haxe_nativestacktrace_exceptionStack
		case "toHaxe":
			return Hx_Field_haxe_nativestacktrace_toHaxe
		case "isHaxeCallStackFrame":
			return Hx_Field_haxe_nativestacktrace_isHaxeCallStackFrame
		case "isConstructorFrame":
			return Hx_Field_haxe_nativestacktrace_isConstructorFrame
		case "capture":
			return Hx_Field_haxe_nativestacktrace_capture
		case "createMethod":
			return Hx_Field_haxe_nativestacktrace_createMethod
		case "stripClosureSuffix":
			return Hx_Field_haxe_nativestacktrace_stripClosureSuffix
		case "isClosureSegment":
			return Hx_Field_haxe_nativestacktrace_isClosureSegment
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_nativestacktrace) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_nativestacktrace_RTTI
}

func (this *Hx_Obj_haxe_nativestacktrace) Hx_As_Hx_Obj_haxe_nativestacktrace() *Hx_Obj_haxe_nativestacktrace {
	return this
}

func Hx_Field_haxe_nativestacktrace_saveStack(exception any) {
}

func Hx_Field_haxe_nativestacktrace_callStack() Hx_Typedef_haxe__nativestacktrace_nativetrace {
	return ((Hx_Typedef_haxe__nativestacktrace_nativetrace)(Hx_Field_haxe_nativestacktrace_capture()))
}

func Hx_Field_haxe_nativestacktrace_exceptionStack() Hx_Typedef_haxe__nativestacktrace_nativetrace {
	return ((Hx_Typedef_haxe__nativestacktrace_nativetrace)(Hx_Field_haxe_nativestacktrace_capture()))
}

func Hx_Field_haxe_nativestacktrace_toHaxe(nativeStackTrace Hx_Typedef_haxe__nativestacktrace_nativetrace, skip__isOptional any) HxArray[Hx_Enum_haxe_stackitem] {
	skip := HxDefault[int32]()
	_ = skip

	if skip__isOptional != nil {
		skip = HxConvert[int32](skip__isOptional)
	} else {
		skip = 0
	}
	{
		if nativeStackTrace == nil {
			return HxMakeArray[Hx_Enum_haxe_stackitem]()
		}

		if skip > 0 {
			var _hx_tmp_0 Hx_Typedef_haxe__nativestacktrace_nativetrace
			_ = _hx_tmp_0
			_hx_tmp_0 = nativeStackTrace
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = skip
			return (Hx_Array_Slice[Hx_Enum_haxe_stackitem](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
		}

		return HxAnyToTypedArray[Hx_Enum_haxe_stackitem](nativeStackTrace)
	}
}

func Hx_Field_haxe_nativestacktrace_isHaxeCallStackFrame(f string) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if f == "runtime.Callers" {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = f
		var _hx_tmp_4 string
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_3
		_hx_tmp_2 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_4, "haxe_nativestacktrace", HxNullable[int32]{})) != -1)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = f
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = _hx_tmp_5
		_hx_tmp_1 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_6, "haxe__callstack", HxNullable[int32]{})) != -1)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_haxe_nativestacktrace_isConstructorFrame(f)
	}

	return _hx_tmp_0
}

func Hx_Field_haxe_nativestacktrace_isConstructorFrame(f string) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if Hx_Field_stringtools_endsWith(f, ".Hx_New") {
		_hx_tmp_3 = true
	} else {
		var _hx_tmp_4 string
		_ = _hx_tmp_4
		_hx_tmp_4 = f
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = _hx_tmp_4
		_hx_tmp_3 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_5, "_CreateEmptyInstance", HxNullable[int32]{})) != -1)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = f
		var _hx_tmp_7 string
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_6
		_hx_tmp_2 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_7, "_CreateInstance", HxNullable[int32]{})) != -1)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_8 string
		_ = _hx_tmp_8
		_hx_tmp_8 = f
		var _hx_tmp_9 string
		_ = _hx_tmp_9
		_hx_tmp_9 = _hx_tmp_8
		_hx_tmp_1 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_9, "haxe_exception_thrown", HxNullable[int32]{})) != -1)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_10 string
		_ = _hx_tmp_10
		_hx_tmp_10 = f
		var _hx_tmp_11 string
		_ = _hx_tmp_11
		_hx_tmp_11 = _hx_tmp_10
		_hx_tmp_0 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_11, "haxe_exception_caught", HxNullable[int32]{})) != -1)
	}

	return _hx_tmp_0
}

func Hx_Field_haxe_nativestacktrace_capture() HxArray[Hx_Enum_haxe_stackitem] {
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{Value: 64, Valid: true}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var pcs []uintptr
	_ = pcs
	pcs = make([]uintptr, _hx_tmp_0)
	var n int
	_ = n
	n = runtime.Callers(int(((int32)(1))), pcs)
	var items HxArray[Hx_Enum_haxe_stackitem]
	_ = items
	items = HxMakeArray[Hx_Enum_haxe_stackitem]()
	var _hx_tmp_1 int
	_ = _hx_tmp_1
	_hx_tmp_1 = n
	if _hx_tmp_1 <= int(((int32)(0))) {
		return items
	}

	var _hx_tmp_2 []uintptr
	_ = _hx_tmp_2
	_hx_tmp_2 = pcs
	pcs = _hx_tmp_2[int(((int32)(0))):n]
	var frames *runtime.Frames
	_ = frames
	frames = runtime.CallersFrames(pcs)
	var more bool
	_ = more
	more = true
	var strippingHaxeStackFrames bool
	_ = strippingHaxeStackFrames
	strippingHaxeStackFrames = true
	for more {
		var hx_tuple_7 struct {
			Frame runtime.Frame
			More  bool
		}
		_ = hx_tuple_7
		hx_tuple_7.Frame, hx_tuple_7.More = (*frames).Next()
		var res struct {
			Frame runtime.Frame
			More  bool
		}
		_ = res
		res = ((struct {
			Frame runtime.Frame
			More  bool
		})(hx_tuple_7))
		var frame runtime.Frame
		_ = frame
		frame = res.Frame
		more = res.More
		if frame.Function == "" {
			continue
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if strippingHaxeStackFrames {
			_hx_tmp_3 = Hx_Field_haxe_nativestacktrace_isHaxeCallStackFrame(frame.Function)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			continue
		}

		strippingHaxeStackFrames = false
		var _hx_tmp_4 HxArray[Hx_Enum_haxe_stackitem]
		_ = _hx_tmp_4
		_hx_tmp_4 = items
		var _hx_tmp_5 Hx_Enum_haxe_stackitem
		_ = _hx_tmp_5
		_hx_tmp_5 = Hx_Field_haxe_nativestacktrace_createMethod(frame.Function)
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = frame.File
		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		_hx_tmp_7 = ((int32)(frame.Line))
		_ = (Hx_Array_Push[Hx_Enum_haxe_stackitem](_hx_tmp_4, Hx_Enum_haxe_stackitem_FilePos{_hx_tmp_5, _hx_tmp_6, _hx_tmp_7, HxNullable[int32]{}}))
	}

	return items
}

func Hx_Field_haxe_nativestacktrace_createMethod(fn string) Hx_Enum_haxe_stackitem {
	fn = Hx_Field_haxe_nativestacktrace_stripClosureSuffix(fn)
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = fn
	var dot int32
	_ = dot
	dot = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_0, ".", HxNullable[int32]{}))
	if dot <= 0 {
		return Hx_Enum_haxe_stackitem_Method{HxNullable[string]{Value: HxStringNull}, fn}
	}

	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = fn
	var _hx_tmp_1 HxNullable[string]
	_ = _hx_tmp_1
	_hx_tmp_1 = HxNullableString((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, HxNullable[int32]{Value: dot, Valid: true})))
	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = fn
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = (dot + 1)
	return Hx_Enum_haxe_stackitem_Method{_hx_tmp_1, (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{}))}
}

func Hx_Field_haxe_nativestacktrace_stripClosureSuffix(fn string) string {
	for {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = fn
		var dot int32
		_ = dot
		dot = (Hx_Field_go_haxe_hxstring_lastIndexOf(_hx_tmp_0, ".", HxNullable[int32]{}))
		if dot <= 0 {
			return fn
		}

		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = fn
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = (dot + 1)
		var seg string
		_ = seg
		seg = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{}))
		if Hx_Field_haxe_nativestacktrace_isClosureSegment(seg) {
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = fn
			fn = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: dot, Valid: true}))
		} else {
			return fn
		}
	}
}

func Hx_Field_haxe_nativestacktrace_isClosureSegment(seg string) bool {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	if Hx_Field_stringtools_startsWith(seg, "func") {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = seg
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 4, HxNullable[int32]{}))
	} else {
		_hx_tmp_0 = seg
	}

	var body string
	_ = body
	body = _hx_tmp_0
	if int32(utf8.RuneCountInString(body)) == 0 {
		return false
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(body))
		for _g < _g1 {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_3
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(body, i))
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if (c).Valid && true && ((c).Value < 48) {
				_hx_tmp_4 = true
			} else {
				_hx_tmp_4 = ((c).Valid && true && ((c).Value > 57))
			}

			if _hx_tmp_4 {
				return false
			}
		}
	}

	return true
}

func init() {
	hxRegisterArrayConstructor[Hx_Enum_haxe_stackitem]()
}
