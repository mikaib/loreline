package core

type Hx_Typedef_loreline_lib_godialoguehandler func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 string, _hx_carg2 string, _hx_carg3 HxArray[*Hx_Obj_loreline_texttag], _hx_carg4 func())

type Hx_Typedef_loreline_lib_gochoicehandler func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter, _hx_carg1 HxArray[*Hx_Obj_loreline_choiceoption], _hx_carg2 func(_hx_carg0 int32))

type Hx_Typedef_loreline_lib_gofinishhandler func(_hx_carg0 *Hx_Obj_loreline_lib_gointerpreter)

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_goscript() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_goscript() {
}

type Hx_Obj_VTable_loreline_lib_goscript interface {
	Hx_Field_resume(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, save string, beat string, options *Hx_Obj_loreline_lib_gooptions) struct {
		Err         error
		Interpreter *Hx_Obj_loreline_lib_gointerpreter
	}
	Hx_Field_print(indent string, newline string) string
	Hx_Field_column() int32
	Hx_Field_id() string
	Hx_Field_toJson(pretty bool) string
	Hx_Field_play(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, beat string, options *Hx_Obj_loreline_lib_gooptions) struct {
		Err         error
		Interpreter *Hx_Obj_loreline_lib_gointerpreter
	}
	Hx_Field_type() string
	Hx_Field_line() int32
	Hx_Field_offset() int32
	Hx_Field_length() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_goscript struct {
	Hx_Obj_loreline_lib_gonode
	VTable          Hx_Obj_VTable_loreline_lib_goscript `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_script *Hx_Obj_loreline_script
}

func Hx_Obj_loreline_lib_goscript_CreateEmptyInstance() *Hx_Obj_loreline_lib_goscript {
	obj := &Hx_Obj_loreline_lib_goscript{}
	obj.Hx_Obj_loreline_lib_gonode = *Hx_Obj_loreline_lib_gonode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_lib_gonode.VTable = obj
	obj.Hx_Field_script = HxDefault[*Hx_Obj_loreline_script]()
	return obj
}

func Hx_Obj_loreline_lib_goscript_CreateInstance(script *Hx_Obj_loreline_script) *Hx_Obj_loreline_lib_goscript {
	obj := Hx_Obj_loreline_lib_goscript_CreateEmptyInstance()
	obj.Hx_New(script)
	return obj
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_New(script *Hx_Obj_loreline_script) {
	{
		this.Hx_Obj_loreline_lib_gonode.Hx_New((func() *Hx_Obj_loreline_node {
			_hx_up := script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_script = script
	}

}

var Hx_Obj_loreline_lib_goscript_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoScript",
	HxMakeArray[string](),
	HxMakeArray[string]("script", "play", "resume", "print", "node", "type", "line", "column", "offset", "length", "id", "toJson"),
	Hx_Obj_loreline_lib_gonode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_goscript_CreateInstance(HxConvert[*Hx_Obj_loreline_script](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lib_goscript_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_goscript_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_goscript_RTTI
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_As_Hx_Obj_loreline_lib_goscript() *Hx_Obj_loreline_lib_goscript {
	return this
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_Field_play(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, beat string, options *Hx_Obj_loreline_lib_gooptions) struct {
	Err         error
	Interpreter *Hx_Obj_loreline_lib_gointerpreter
} {
	{
		var hx_try_return struct {
			Err         error
			Interpreter *Hx_Obj_loreline_lib_gointerpreter
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_6 error
							_ = _hx_tmp_6
							_hx_tmp_6 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_374 any
							_ = hx_dyncast_374
							hx_dyncast_374 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gointerpreter")
							var _hx_tmp_7 *Hx_Obj_loreline_lib_gointerpreter
							_ = _hx_tmp_7
							if hx_dyncast_374 != nil {
								_hx_tmp_7 = hx_dyncast_374.(interface {
									Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
								}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
							} else {
								_hx_tmp_7 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err         error
								Interpreter *Hx_Obj_loreline_lib_gointerpreter
							}{_hx_tmp_6, _hx_tmp_7}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 *Hx_Obj_loreline_script
				_ = _hx_tmp_0
				_hx_tmp_0 = this.Hx_Field_script
				var _hx_tmp_1 Hx_Typedef_loreline_dialoguehandler
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_loreline_lib_gointerpreter_dialogueHandler(dialogue)
				var _hx_tmp_2 Hx_Typedef_loreline_choicehandler
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_loreline_lib_gointerpreter_choiceHandler(choice)
				var _hx_tmp_3 Hx_Typedef_loreline_finishhandler
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_lib_gointerpreter_finishHandler(finish)
				var _hx_tmp_4 HxNullable[string]
				_ = _hx_tmp_4
				_hx_tmp_4 = HxNullableString(Hx_Field_loreline_lib_golibrary_emptyToNull(beat))
				var interpreter *Hx_Obj_loreline_interpreter
				_ = interpreter
				interpreter = Hx_Field_loreline_loreline_play(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, Hx_Field_loreline_lib_gooptions_build(options))
				var hx_dyncast_373 any
				_ = hx_dyncast_373
				hx_dyncast_373 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Field_loreline_lib_gointerpreter_of(interpreter))), "Hx_Obj_loreline_lib_gointerpreter")
				var _hx_tmp_5 *Hx_Obj_loreline_lib_gointerpreter
				_ = _hx_tmp_5
				if hx_dyncast_373 != nil {
					_hx_tmp_5 = hx_dyncast_373.(interface {
						Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
					}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
				} else {
					_hx_tmp_5 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err         error
					Interpreter *Hx_Obj_loreline_lib_gointerpreter
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_5}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_goscript) Play(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, beat string, options *Hx_Obj_loreline_lib_gooptions) (*Hx_Obj_loreline_lib_gointerpreter, error) {
	_v := this.VTable.Hx_Field_play(dialogue, choice, finish, beat, options)
	return _v.Interpreter, _v.Err
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_Field_resume(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, save string, beat string, options *Hx_Obj_loreline_lib_gooptions) struct {
	Err         error
	Interpreter *Hx_Obj_loreline_lib_gointerpreter
} {
	{
		var hx_try_return struct {
			Err         error
			Interpreter *Hx_Obj_loreline_lib_gointerpreter
		}
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							var _hx_tmp_7 error
							_ = _hx_tmp_7
							_hx_tmp_7 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_376 any
							_ = hx_dyncast_376
							hx_dyncast_376 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gointerpreter")
							var _hx_tmp_8 *Hx_Obj_loreline_lib_gointerpreter
							_ = _hx_tmp_8
							if hx_dyncast_376 != nil {
								_hx_tmp_8 = hx_dyncast_376.(interface {
									Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
								}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
							} else {
								_hx_tmp_8 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err         error
								Interpreter *Hx_Obj_loreline_lib_gointerpreter
							}{_hx_tmp_7, _hx_tmp_8}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 *Hx_Obj_loreline_script
				_ = _hx_tmp_0
				_hx_tmp_0 = this.Hx_Field_script
				var _hx_tmp_1 Hx_Typedef_loreline_dialoguehandler
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_loreline_lib_gointerpreter_dialogueHandler(dialogue)
				var _hx_tmp_2 Hx_Typedef_loreline_choicehandler
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_loreline_lib_gointerpreter_choiceHandler(choice)
				var _hx_tmp_3 Hx_Typedef_loreline_finishhandler
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_lib_gointerpreter_finishHandler(finish)
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = Hx_Field_loreline_json_parse(save)
				var _hx_tmp_5 HxNullable[string]
				_ = _hx_tmp_5
				_hx_tmp_5 = HxNullableString(Hx_Field_loreline_lib_golibrary_emptyToNull(beat))
				var interpreter *Hx_Obj_loreline_interpreter
				_ = interpreter
				interpreter = Hx_Field_loreline_loreline_resume(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, Hx_Field_loreline_lib_gooptions_build(options))
				var hx_dyncast_375 any
				_ = hx_dyncast_375
				hx_dyncast_375 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Field_loreline_lib_gointerpreter_of(interpreter))), "Hx_Obj_loreline_lib_gointerpreter")
				var _hx_tmp_6 *Hx_Obj_loreline_lib_gointerpreter
				_ = _hx_tmp_6
				if hx_dyncast_375 != nil {
					_hx_tmp_6 = hx_dyncast_375.(interface {
						Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
					}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
				} else {
					_hx_tmp_6 = nil
				}

				hx_try_state = 1
				hx_try_return = struct {
					Err         error
					Interpreter *Hx_Obj_loreline_lib_gointerpreter
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), _hx_tmp_6}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_goscript) Resume(dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler, save string, beat string, options *Hx_Obj_loreline_lib_gooptions) (*Hx_Obj_loreline_lib_gointerpreter, error) {
	_v := this.VTable.Hx_Field_resume(dialogue, choice, finish, save, beat, options)
	return _v.Interpreter, _v.Err
}

func (this *Hx_Obj_loreline_lib_goscript) Hx_Field_print(indent string, newline string) string {
	var _hx_tmp_0 *Hx_Obj_loreline_script
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_script
	var _hx_tmp_1 HxNullable[string]
	_ = _hx_tmp_1
	_hx_tmp_1 = HxNullableString(indent)
	return Hx_Field_loreline_loreline_print(_hx_tmp_0, _hx_tmp_1, HxNullableString(newline))
}

func (this *Hx_Obj_loreline_lib_goscript) Print(indent string, newline string) string {
	return this.VTable.Hx_Field_print(indent, newline)
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_choiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_texttag]()
}
