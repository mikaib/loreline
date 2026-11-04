package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_gooptions() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_gooptions() {
}

type Hx_Obj_VTable_loreline_lib_gooptions interface {
	Hx_Field_setTranslations(translations *Hx_Obj_loreline_lib_gotranslations)
	Hx_Field_addFunction(name string, fn func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 HxArray[any]) any)
	Hx_Field_setStrictAccess(strictAccess bool)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_gooptions struct {
	VTable                Hx_Obj_VTable_loreline_lib_gooptions `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_functions    Hx_Typedef_loreline_functionsmap
	Hx_Field_strictAccess bool
	Hx_Field_translations *Hx_Obj_loreline_lib_gotranslations
}

func Hx_Obj_loreline_lib_gooptions_CreateEmptyInstance() *Hx_Obj_loreline_lib_gooptions {
	obj := &Hx_Obj_loreline_lib_gooptions{}
	obj.VTable = obj
	obj.Hx_Field_functions = HxDefault[Hx_Typedef_loreline_functionsmap]()
	obj.Hx_Field_strictAccess = HxDefault[bool]()
	obj.Hx_Field_translations = HxDefault[*Hx_Obj_loreline_lib_gotranslations]()
	return obj
}

func Hx_Obj_loreline_lib_gooptions_CreateInstance() *Hx_Obj_loreline_lib_gooptions {
	obj := Hx_Obj_loreline_lib_gooptions_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_New() {
	{
		this.Hx_Field_functions = nil
	}
	{
		this.Hx_Field_strictAccess = false
	}
	{
		this.Hx_Field_translations = nil
	}
	{
	}

}

var Hx_Obj_loreline_lib_gooptions_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoOptions",
	HxMakeArray[string]("build"),
	HxMakeArray[string]("functions", "strictAccess", "translations", "addFunction", "setStrictAccess", "setTranslations"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_gooptions_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lib_gooptions_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_gooptions_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "build":
			return Hx_Field_loreline_lib_gooptions_build
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_gooptions_RTTI
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_As_Hx_Obj_loreline_lib_gooptions() *Hx_Obj_loreline_lib_gooptions {
	return this
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_Field_addFunction(name string, fn func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 HxArray[any]) any) {
	if this.Hx_Field_functions == nil {
		this.Hx_Field_functions = (func() *Hx_Obj_haxe_imap {
			_hx_icast_288 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_288 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_288, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_288.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = name
	this.Hx_Field_functions.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_0, ((any)(func(interpreter *Hx_Obj_loreline_interpreter, args HxArray[any]) any {
		var fn1 func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 HxArray[any]) any
		_ = fn1
		fn1 = fn
		var tmp *Hx_Obj_loreline_lib_gointerpreter
		_ = tmp
		tmp = Hx_Field_loreline_lib_gointerpreter_of(interpreter)
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g1
				if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
					break
				}

				var arg any
				_ = arg
				arg = args.Get(_g1)
				_g1++
				var _hx_tmp_2 HxArray[any]
				_ = _hx_tmp_2
				_hx_tmp_2 = _g
				var _hx_tmp_3 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_3
				_hx_tmp_3 = interpreter
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = arg
				_ = (Hx_Array_Push[any](_hx_tmp_2, Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_3, _hx_tmp_4, nil)))
			}
		}

		return fn1(tmp, _g)
	})))
}

func (this *Hx_Obj_loreline_lib_gooptions) AddFunction(name string, fn func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 HxArray[any]) any) {
	this.VTable.Hx_Field_addFunction(name, fn)
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_Field_setStrictAccess(strictAccess bool) {
	this.Hx_Field_strictAccess = strictAccess
}

func (this *Hx_Obj_loreline_lib_gooptions) SetStrictAccess(strictAccess bool) {
	this.VTable.Hx_Field_setStrictAccess(strictAccess)
}

func (this *Hx_Obj_loreline_lib_gooptions) Hx_Field_setTranslations(translations *Hx_Obj_loreline_lib_gotranslations) {
	this.Hx_Field_translations = translations
}

func (this *Hx_Obj_loreline_lib_gooptions) SetTranslations(translations *Hx_Obj_loreline_lib_gotranslations) {
	this.VTable.Hx_Field_setTranslations(translations)
}

func Hx_Field_loreline_lib_gooptions_build(options *Hx_Obj_loreline_lib_gooptions) *Hx_Obj_loreline_interpreteroptions {
	if options == nil {
		return nil
	}

	var _hx_tmp_0 Hx_Typedef_map
	_ = _hx_tmp_0
	if options.Hx_Field_translations != nil {
		_hx_tmp_0 = options.Hx_Field_translations.Hx_Field_translations
	} else {
		_hx_tmp_0 = nil
	}

	return Hx_Obj_loreline_interpreteroptions_CreateInstance(nil, options.Hx_Field_functions, HxNullable[bool]{Value: options.Hx_Field_strictAccess, Valid: true}, nil, _hx_tmp_0, nil)
}

func init() {
	hxRegisterArrayConstructor[any]()
}
