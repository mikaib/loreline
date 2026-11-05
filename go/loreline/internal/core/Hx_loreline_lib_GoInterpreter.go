package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_gointerpreter() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_gointerpreter() {
}

type Hx_Obj_VTable_loreline_lib_gointerpreter interface {
	Hx_Field_key() string
	Hx_Field_seedRandomFromClock()
	Hx_Field_character(name string) any
	Hx_Field_topLevelStateField(name string) any
	Hx_Field_currentNode() *Hx_Obj_loreline_lib_gonode
	Hx_Field_start(beat string) error
	Hx_Field_save() struct {
		Err  error
		Save string
	}
	Hx_Field_resumeSpawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) struct {
		Err         error
		Interpreter *Hx_Obj_loreline_lib_gointerpreter
	}
	Hx_Field_characterField(character string, name string) any
	Hx_Field_setStateField(name string, value any)
	Hx_Field_setTopLevelStateField(name string, value any)
	Hx_Field_host() any
	Hx_Field_setHost(value any)
	Hx_Field_restore(save string) error
	Hx_Field_spawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) struct {
		Err         error
		Interpreter *Hx_Obj_loreline_lib_gointerpreter
	}
	Hx_Field_seedRandom(seed float64)
	Hx_Field_dispose() error
	Hx_Field_setCharacterField(character string, name string, value any)
	Hx_Field_isRoot() bool
	Hx_Field_resume() error
	Hx_Field_resumableSpawnKeys() HxArray[string]
	Hx_Field_stateField(name string) any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_gointerpreter struct {
	VTable               Hx_Obj_VTable_loreline_lib_gointerpreter `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_interpreter *Hx_Obj_loreline_interpreter
	Hx_Field_hostValue   any
}

func Hx_Obj_loreline_lib_gointerpreter_CreateEmptyInstance() *Hx_Obj_loreline_lib_gointerpreter {
	obj := &Hx_Obj_loreline_lib_gointerpreter{}
	obj.VTable = obj
	obj.Hx_Field_interpreter = HxDefault[*Hx_Obj_loreline_interpreter]()
	obj.Hx_Field_hostValue = HxDefault[any]()
	return obj
}

func Hx_Obj_loreline_lib_gointerpreter_CreateInstance(interpreter *Hx_Obj_loreline_interpreter) *Hx_Obj_loreline_lib_gointerpreter {
	obj := Hx_Obj_loreline_lib_gointerpreter_CreateEmptyInstance()
	obj.Hx_New(interpreter)
	return obj
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_New(interpreter *Hx_Obj_loreline_interpreter) {
	{
		this.Hx_Field_hostValue = nil
	}
	{
		this.Hx_Field_interpreter = interpreter
	}

}

var Hx_Obj_loreline_lib_gointerpreter_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoInterpreter",
	HxMakeArray[string]("of", "dialogueHandler", "choiceHandler", "finishHandler"),
	HxMakeArray[string]("interpreter", "hostValue", "host", "setHost", "key", "isRoot", "start", "save", "restore", "resume", "spawn", "resumeSpawn", "resumableSpawnKeys", "seedRandom", "seedRandomFromClock", "dispose", "character", "characterField", "setCharacterField", "stateField", "setStateField", "topLevelStateField", "setTopLevelStateField", "currentNode"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_gointerpreter_CreateInstance(HxConvert[*Hx_Obj_loreline_interpreter](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lib_gointerpreter_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_gointerpreter_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "of":
			return Hx_Field_loreline_lib_gointerpreter_of
		case "dialogueHandler":
			return Hx_Field_loreline_lib_gointerpreter_dialogueHandler
		case "choiceHandler":
			return Hx_Field_loreline_lib_gointerpreter_choiceHandler
		case "finishHandler":
			return Hx_Field_loreline_lib_gointerpreter_finishHandler
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_gointerpreter_RTTI
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter {
	return this
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_host() any {
	return this.Hx_Field_hostValue
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Host() any {
	return this.VTable.Hx_Field_host()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_setHost(value any) {
	this.Hx_Field_hostValue = value
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SetHost(value any) {
	this.VTable.Hx_Field_setHost(value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_key() string {
	var tmp HxNullable[string]
	_ = tmp
	tmp = HxNullableString(this.Hx_Field_interpreter.Hx_Field_key)
	if (tmp).Valid != false {
		return (tmp).Value
	} else {
		return ""
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Key() string {
	return this.VTable.Hx_Field_key()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_isRoot() bool {
	return this.Hx_Field_interpreter.VTable.Hx_Field_isRoot()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) IsRoot() bool {
	return this.VTable.Hx_Field_isRoot()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_start(beat string) error {
	{
		var hx_try_return error
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
							hx_try_return = Hx_Field_loreline_lib_golibrary_toError(e)
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				this.Hx_Field_interpreter.VTable.Hx_Field_start(HxNullableString(Hx_Field_loreline_lib_golibrary_emptyToNull(beat)))
				hx_try_state = 1
				hx_try_return = nil
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Start(beat string) error {
	return this.VTable.Hx_Field_start(beat)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_save() struct {
	Err  error
	Save string
} {
	{
		var hx_try_return struct {
			Err  error
			Save string
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
							var _hx_tmp_1 error
							_ = _hx_tmp_1
							_hx_tmp_1 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							hx_try_state = 1
							hx_try_return = struct {
								Err  error
								Save string
							}{_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_toString(((any)("")))}
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
				_hx_tmp_0 = this.Hx_Field_interpreter.VTable.Hx_Field_save()
				hx_try_state = 1
				hx_try_return = struct {
					Err  error
					Save string
				}{func() error {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(nil)))
					if _hx_d == nil {
						var _hx_z error
						return _hx_z
					}
					return _hx_d.(error)
				}(), Hx_Field_go_haxe_hxdynamic_toString(((any)(Hx_Field_loreline_json_stringify(_hx_tmp_0, nil))))}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Save() (string, error) {
	_v := this.VTable.Hx_Field_save()
	return _v.Save, _v.Err
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_restore(save string) error {
	{
		var hx_try_return error
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
							hx_try_return = Hx_Field_loreline_lib_golibrary_toError(e)
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				this.Hx_Field_interpreter.VTable.Hx_Field_restore(Hx_Field_loreline_json_parse(save))
				hx_try_state = 1
				hx_try_return = nil
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Restore(save string) error {
	return this.VTable.Hx_Field_restore(save)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_resume() error {
	{
		var hx_try_return error
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
							hx_try_return = Hx_Field_loreline_lib_golibrary_toError(e)
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				this.Hx_Field_interpreter.VTable.Hx_Field_resume()
				hx_try_state = 1
				hx_try_return = nil
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Resume() error {
	return this.VTable.Hx_Field_resume()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_spawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) struct {
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
							var _hx_tmp_5 error
							_ = _hx_tmp_5
							_hx_tmp_5 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_284 any
							_ = hx_dyncast_284
							hx_dyncast_284 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gointerpreter")
							var _hx_tmp_6 *Hx_Obj_loreline_lib_gointerpreter
							_ = _hx_tmp_6
							if hx_dyncast_284 != nil {
								_hx_tmp_6 = hx_dyncast_284.(interface {
									Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
								}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
							} else {
								_hx_tmp_6 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err         error
								Interpreter *Hx_Obj_loreline_lib_gointerpreter
							}{_hx_tmp_5, _hx_tmp_6}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = key
				var _hx_tmp_1 Hx_Typedef_loreline_dialoguehandler
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_loreline_lib_gointerpreter_dialogueHandler(dialogue)
				var _hx_tmp_2 Hx_Typedef_loreline_choicehandler
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_loreline_lib_gointerpreter_choiceHandler(choice)
				var _hx_tmp_3 Hx_Typedef_loreline_finishhandler
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_lib_gointerpreter_finishHandler(finish)
				var child *Hx_Obj_loreline_interpreter
				_ = child
				child = this.Hx_Field_interpreter.VTable.Hx_Field_spawn(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, nil)
				var hx_dyncast_283 any
				_ = hx_dyncast_283
				hx_dyncast_283 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Field_loreline_lib_gointerpreter_of(child))), "Hx_Obj_loreline_lib_gointerpreter")
				var _hx_tmp_4 *Hx_Obj_loreline_lib_gointerpreter
				_ = _hx_tmp_4
				if hx_dyncast_283 != nil {
					_hx_tmp_4 = hx_dyncast_283.(interface {
						Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
					}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
				} else {
					_hx_tmp_4 = nil
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
				}(), _hx_tmp_4}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Spawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) (*Hx_Obj_loreline_lib_gointerpreter, error) {
	_v := this.VTable.Hx_Field_spawn(key, dialogue, choice, finish)
	return _v.Interpreter, _v.Err
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_resumeSpawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) struct {
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
							var _hx_tmp_5 error
							_ = _hx_tmp_5
							_hx_tmp_5 = func() error {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(((any)(Hx_Field_loreline_lib_golibrary_toError(e))))
								if _hx_d == nil {
									var _hx_z error
									return _hx_z
								}
								return _hx_d.(error)
							}()
							var hx_dyncast_286 any
							_ = hx_dyncast_286
							hx_dyncast_286 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(nil)), "Hx_Obj_loreline_lib_gointerpreter")
							var _hx_tmp_6 *Hx_Obj_loreline_lib_gointerpreter
							_ = _hx_tmp_6
							if hx_dyncast_286 != nil {
								_hx_tmp_6 = hx_dyncast_286.(interface {
									Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
								}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
							} else {
								_hx_tmp_6 = nil
							}

							hx_try_state = 1
							hx_try_return = struct {
								Err         error
								Interpreter *Hx_Obj_loreline_lib_gointerpreter
							}{_hx_tmp_5, _hx_tmp_6}
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = key
				var _hx_tmp_1 Hx_Typedef_loreline_dialoguehandler
				_ = _hx_tmp_1
				_hx_tmp_1 = Hx_Field_loreline_lib_gointerpreter_dialogueHandler(dialogue)
				var _hx_tmp_2 Hx_Typedef_loreline_choicehandler
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_loreline_lib_gointerpreter_choiceHandler(choice)
				var _hx_tmp_3 Hx_Typedef_loreline_finishhandler
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_loreline_lib_gointerpreter_finishHandler(finish)
				var child *Hx_Obj_loreline_interpreter
				_ = child
				child = this.Hx_Field_interpreter.VTable.Hx_Field_resumeSpawn(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, nil)
				var hx_dyncast_285 any
				_ = hx_dyncast_285
				hx_dyncast_285 = Hx_Field_go_haxe_hxdynamic_castClass(((any)(Hx_Field_loreline_lib_gointerpreter_of(child))), "Hx_Obj_loreline_lib_gointerpreter")
				var _hx_tmp_4 *Hx_Obj_loreline_lib_gointerpreter
				_ = _hx_tmp_4
				if hx_dyncast_285 != nil {
					_hx_tmp_4 = hx_dyncast_285.(interface {
						Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
					}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
				} else {
					_hx_tmp_4 = nil
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
				}(), _hx_tmp_4}
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) ResumeSpawn(key string, dialogue Hx_Typedef_loreline_lib_godialoguehandler, choice Hx_Typedef_loreline_lib_gochoicehandler, finish Hx_Typedef_loreline_lib_gofinishhandler) (*Hx_Obj_loreline_lib_gointerpreter, error) {
	_v := this.VTable.Hx_Field_resumeSpawn(key, dialogue, choice, finish)
	return _v.Interpreter, _v.Err
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_resumableSpawnKeys() HxArray[string] {
	return this.Hx_Field_interpreter.VTable.Hx_Field_resumableSpawnKeys()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) ResumableSpawnKeys() HxArray[string] {
	return this.VTable.Hx_Field_resumableSpawnKeys()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_seedRandom(seed float64) {
	this.Hx_Field_interpreter.VTable.Hx_Field_seedRandom(HxNullable[float64]{Value: seed, Valid: true})
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SeedRandom(seed float64) {
	this.VTable.Hx_Field_seedRandom(seed)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_seedRandomFromClock() {
	this.Hx_Field_interpreter.VTable.Hx_Field_seedRandom(nil)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SeedRandomFromClock() {
	this.VTable.Hx_Field_seedRandomFromClock()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_dispose() error {
	{
		var hx_try_return error
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
							hx_try_return = Hx_Field_loreline_lib_golibrary_toError(e)
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				this.Hx_Field_interpreter.VTable.Hx_Field_dispose()
				hx_try_state = 1
				hx_try_return = nil
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Dispose() error {
	return this.VTable.Hx_Field_dispose()
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_character(name string) any {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_interpreter
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_interpreter.VTable.Hx_Field_getCharacter(name)
	return Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_0, _hx_tmp_1, nil)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Character(name string) any {
	return this.VTable.Hx_Field_character(name)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_characterField(character string, name string) any {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_interpreter
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_interpreter.VTable.Hx_Field_getCharacterField(character, name)
	return Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_0, _hx_tmp_1, nil)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) CharacterField(character string, name string) any {
	return this.VTable.Hx_Field_characterField(character, name)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_setCharacterField(character string, name string, value any) {
	this.Hx_Field_interpreter.VTable.Hx_Field_setCharacterField(character, name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SetCharacterField(character string, name string, value any) {
	this.VTable.Hx_Field_setCharacterField(character, name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_stateField(name string) any {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_interpreter
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_interpreter.VTable.Hx_Field_getStateField(name)
	return Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_0, _hx_tmp_1, nil)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) StateField(name string) any {
	return this.VTable.Hx_Field_stateField(name)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_setStateField(name string, value any) {
	this.Hx_Field_interpreter.VTable.Hx_Field_setStateField(name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SetStateField(name string, value any) {
	this.VTable.Hx_Field_setStateField(name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_topLevelStateField(name string) any {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_interpreter
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_interpreter.VTable.Hx_Field_getTopLevelStateField(name)
	return Hx_Field_loreline_lib_golibrary_toHost(_hx_tmp_0, _hx_tmp_1, nil)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) TopLevelStateField(name string) any {
	return this.VTable.Hx_Field_topLevelStateField(name)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_setTopLevelStateField(name string, value any) {
	this.Hx_Field_interpreter.VTable.Hx_Field_setTopLevelStateField(name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) SetTopLevelStateField(name string, value any) {
	this.VTable.Hx_Field_setTopLevelStateField(name, value)
}

func (this *Hx_Obj_loreline_lib_gointerpreter) Hx_Field_currentNode() *Hx_Obj_loreline_lib_gonode {
	var node *Hx_Obj_loreline_astnode
	_ = node
	node = this.Hx_Field_interpreter.VTable.Hx_Field_currentNode()
	if node != nil {
		return Hx_Obj_loreline_lib_gonode_CreateInstance((func() *Hx_Obj_loreline_node {
			_hx_up := node
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
	} else {
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_gointerpreter) CurrentNode() *Hx_Obj_loreline_lib_gonode {
	return this.VTable.Hx_Field_currentNode()
}

func Hx_Field_loreline_lib_gointerpreter_of(interpreter *Hx_Obj_loreline_interpreter) *Hx_Obj_loreline_lib_gointerpreter {
	if interpreter == nil {
		return nil
	}

	var existing any
	_ = existing
	existing = interpreter.Hx_Field_wrapper
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(existing, nil))) {
		var hx_dyncast_287 any
		_ = hx_dyncast_287
		hx_dyncast_287 = Hx_Field_go_haxe_hxdynamic_castClass(existing, "Hx_Obj_loreline_lib_gointerpreter")
		var _hx_tmp_0 *Hx_Obj_loreline_lib_gointerpreter
		_ = _hx_tmp_0
		if hx_dyncast_287 != nil {
			_hx_tmp_0 = hx_dyncast_287.(interface {
				Hx_As_Hx_Obj_loreline_lib_gointerpreter() *Hx_Obj_loreline_lib_gointerpreter
			}).Hx_As_Hx_Obj_loreline_lib_gointerpreter()
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0
	}

	var wrapper *Hx_Obj_loreline_lib_gointerpreter
	_ = wrapper
	wrapper = Hx_Obj_loreline_lib_gointerpreter_CreateInstance(interpreter)
	interpreter.Hx_Field_wrapper = ((any)(wrapper))
	return wrapper
}

func Hx_Field_loreline_lib_gointerpreter_dialogueHandler(handler Hx_Typedef_loreline_lib_godialoguehandler) Hx_Typedef_loreline_dialoguehandler {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(handler)), ((any)(nil))))) {
		return nil
	}

	return ((Hx_Typedef_loreline_dialoguehandler)(func(interpreter *Hx_Obj_loreline_interpreter, character string, text string, tags HxArray[*Hx_Obj_loreline_texttag], callback func()) {
		var tmp HxNullable[string]
		_ = tmp
		tmp = HxNullableString(character)
		var _hx_tmp_0 *Hx_Obj_loreline_lib_gointerpreter
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_loreline_lib_gointerpreter_of(interpreter)
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		if (tmp).Valid != false {
			_hx_tmp_1 = (tmp).Value
		} else {
			_hx_tmp_1 = ""
		}

		handler(_hx_tmp_0, _hx_tmp_1, text, tags, callback)
	}))
}

func Hx_Field_loreline_lib_gointerpreter_choiceHandler(handler Hx_Typedef_loreline_lib_gochoicehandler) Hx_Typedef_loreline_choicehandler {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(handler)), ((any)(nil))))) {
		return nil
	}

	return ((Hx_Typedef_loreline_choicehandler)(func(interpreter *Hx_Obj_loreline_interpreter, options HxArray[*Hx_Obj_loreline_choiceoption], callback func(index int32)) {
		handler(Hx_Field_loreline_lib_gointerpreter_of(interpreter), options, callback)
	}))
}

func Hx_Field_loreline_lib_gointerpreter_finishHandler(handler Hx_Typedef_loreline_lib_gofinishhandler) Hx_Typedef_loreline_finishhandler {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(handler)), ((any)(nil))))) {
		return nil
	}

	return ((Hx_Typedef_loreline_finishhandler)(func(interpreter *Hx_Obj_loreline_interpreter) {
		handler(Hx_Field_loreline_lib_gointerpreter_of(interpreter))
	}))
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_choiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_texttag]()
	hxRegisterArrayConstructor[string]()
}
