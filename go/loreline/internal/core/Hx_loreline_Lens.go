package core

import "unicode/utf8"

type Hx_Typedef_loreline_whenwarning = any

func Hx_ClassInit_Hx_Obj_VTable_loreline_reference() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_reference() {
}

type Hx_Obj_VTable_loreline_reference interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_reference struct {
	VTable          Hx_Obj_VTable_loreline_reference `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target any
	Hx_Field_origin *Hx_Obj_loreline_node
}

func Hx_Obj_loreline_reference_CreateEmptyInstance() *Hx_Obj_loreline_reference {
	obj := &Hx_Obj_loreline_reference{}
	obj.VTable = obj
	obj.Hx_Field_target = HxDefault[any]()
	obj.Hx_Field_origin = HxDefault[*Hx_Obj_loreline_node]()
	return obj
}

func Hx_Obj_loreline_reference_CreateInstance(target any, origin *Hx_Obj_loreline_node) *Hx_Obj_loreline_reference {
	obj := Hx_Obj_loreline_reference_CreateEmptyInstance()
	obj.Hx_New(target, origin)
	return obj
}

func (this *Hx_Obj_loreline_reference) Hx_New(target any, origin *Hx_Obj_loreline_node) {
	{
		this.Hx_Field_target = target
		this.Hx_Field_origin = origin
	}

}

var Hx_Obj_loreline_reference_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Reference",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "origin"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_reference_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[*Hx_Obj_loreline_node](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_reference_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_reference_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_reference) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_reference_RTTI
}

func (this *Hx_Obj_loreline_reference) Hx_As_Hx_Obj_loreline_reference() *Hx_Obj_loreline_reference {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_funclorscript() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_funclorscript() {
}

type Hx_Obj_VTable_loreline_funclorscript interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_funclorscript struct {
	VTable                   Hx_Obj_VTable_loreline_funclorscript `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_func            *Hx_Obj_loreline_nfunctiondecl
	Hx_Field_codeToLorscript *Hx_Obj_loreline_codetolorscript
	Hx_Field_lorscript       string
	Hx_Field_expr            any
	Hx_Field_error           *Hx_Obj_loreline_error
}

func Hx_Obj_loreline_funclorscript_CreateEmptyInstance() *Hx_Obj_loreline_funclorscript {
	obj := &Hx_Obj_loreline_funclorscript{}
	obj.VTable = obj
	obj.Hx_Field_func = HxDefault[*Hx_Obj_loreline_nfunctiondecl]()
	obj.Hx_Field_codeToLorscript = HxDefault[*Hx_Obj_loreline_codetolorscript]()
	obj.Hx_Field_lorscript = HxDefault[string]()
	obj.Hx_Field_expr = HxDefault[any]()
	obj.Hx_Field_error = HxDefault[*Hx_Obj_loreline_error]()
	return obj
}

func Hx_Obj_loreline_funclorscript_CreateInstance(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) *Hx_Obj_loreline_funclorscript {
	obj := Hx_Obj_loreline_funclorscript_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_func)
	return obj
}

func (this *Hx_Obj_loreline_funclorscript) Hx_New(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) {
	{
		this.Hx_Field_lorscript = HxStringNull
	}
	{
		this.Hx_Field_expr = nil
	}
	{
		this.Hx_Field_error = nil
	}
	{
		this.Hx_Field_func = _hx_reserved_func
		this.Hx_Field_codeToLorscript = Hx_Obj_loreline_codetolorscript_CreateInstance()
		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
							_ = e
							{
								if Hx_Field_std_isOfType(e, Hx_Obj_loreline_error_RTTI) {
									var hx_dyncast_130 any
									_ = hx_dyncast_130
									hx_dyncast_130 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_error")
									var _hx_tmp_0 *Hx_Obj_loreline_error
									_ = _hx_tmp_0
									if hx_dyncast_130 != nil {
										_hx_tmp_0 = hx_dyncast_130.(interface{ Hx_As_Hx_Obj_loreline_error() *Hx_Obj_loreline_error }).Hx_As_Hx_Obj_loreline_error()
									} else {
										_hx_tmp_0 = nil
									}

									this.Hx_Field_error = _hx_tmp_0
									this.Hx_Field_error.Hx_Field_pos = _hx_reserved_func.Hx_Field_pos.VTable.Hx_Field_withOffset(this.Hx_Field_codeToLorscript.Hx_Field_inputText, this.Hx_Field_error.Hx_Field_pos.Hx_Field_offset, this.Hx_Field_error.Hx_Field_pos.Hx_Field_length, _hx_reserved_func.Hx_Field_pos.Hx_Field_offset)
								}

								this.Hx_Field_lorscript = this.Hx_Field_codeToLorscript.Hx_Field_output.VTable.Hx_Field_toString()
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					this.Hx_Field_lorscript = this.Hx_Field_codeToLorscript.VTable.Hx_Field_process(_hx_reserved_func.Hx_Field_code)
				}
			}()
			if hx_try_state == 1 {
				return
			}
		}

		{
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
							_ = e
							{
								if this.Hx_Field_error == nil {
									if Hx_Field_std_isOfType(e, Hx_Obj_loreline_lorscript_error_RTTI) {
										var hx_dyncast_128 any
										_ = hx_dyncast_128
										hx_dyncast_128 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_lorscript_error")
										var _hx_tmp_3 *Hx_Obj_loreline_lorscript_error
										_ = _hx_tmp_3
										if hx_dyncast_128 != nil {
											_hx_tmp_3 = hx_dyncast_128.(interface {
												Hx_As_Hx_Obj_loreline_lorscript_error() *Hx_Obj_loreline_lorscript_error
											}).Hx_As_Hx_Obj_loreline_lorscript_error()
										} else {
											_hx_tmp_3 = nil
										}

										var lorscriptError *Hx_Obj_loreline_lorscript_error
										_ = lorscriptError
										lorscriptError = _hx_tmp_3
										var tmp string
										_ = tmp
										var _g Hx_Enum_loreline_lorscript_errordef
										_ = _g
										_g = lorscriptError.Hx_Field_e
										switch _g.Hx_Field_enumIndex() {
										case 0:
											{
												var c int32
												_ = c
												c = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidChar).Hx_Field_c
												tmp = ("Invalid character: " + Hx_Field_go_haxe_hxdynamic_toString(c))
											}

										case 1:
											{
												var s string
												_ = s
												s = _g.(Hx_Enum_loreline_lorscript_errordef_EUnexpected).Hx_Field_s
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Unexpected: ", s)
											}

										case 2:
											{
												tmp = "Unterminated string"
											}

										case 3:
											{
												tmp = "Unterminated comment"
											}

										case 4:
											{
												var msg string
												_ = msg
												msg = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor).Hx_Field_msg
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid preprocessor: ", msg)
											}

										case 5:
											{
												var v string
												_ = v
												v = _g.(Hx_Enum_loreline_lorscript_errordef_EUnknownVariable).Hx_Field_v
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Unknown variable: ", v)
											}

										case 6:
											{
												var v string
												_ = v
												v = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidIterator).Hx_Field_v
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid iterator: ", v)
											}

										case 7:
											{
												var op string
												_ = op
												op = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidOp).Hx_Field_op
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid operator: ", op)
											}

										case 8:
											{
												var f string
												_ = f
												f = _g.(Hx_Enum_loreline_lorscript_errordef_EInvalidAccess).Hx_Field_f
												tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid access: ", f)
											}

										case 9:
											{
												var msg string
												_ = msg
												msg = _g.(Hx_Enum_loreline_lorscript_errordef_ECustom).Hx_Field_msg
												tmp = msg
											}

										}

										this.Hx_Field_error = (func() *Hx_Obj_loreline_error {
											_hx_up := Hx_Obj_loreline_wrappederror_CreateInstance(lorscriptError, tmp, this.Hx_Field_codeToLorscript.VTable.Hx_Field_toLorelinePos(_hx_reserved_func.Hx_Field_pos, lorscriptError.Hx_Field_pmin, lorscriptError.Hx_Field_pmax))
											if _hx_up != nil {
												return &_hx_up.Hx_Obj_loreline_error
											}
											return nil
										})()
									} else {
										if Hx_Field_std_isOfType(e, Hx_Obj_loreline_error_RTTI) {
											var hx_dyncast_129 any
											_ = hx_dyncast_129
											hx_dyncast_129 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_error")
											var _hx_tmp_4 *Hx_Obj_loreline_error
											_ = _hx_tmp_4
											if hx_dyncast_129 != nil {
												_hx_tmp_4 = hx_dyncast_129.(interface{ Hx_As_Hx_Obj_loreline_error() *Hx_Obj_loreline_error }).Hx_As_Hx_Obj_loreline_error()
											} else {
												_hx_tmp_4 = nil
											}

											this.Hx_Field_error = _hx_tmp_4
											this.Hx_Field_error.Hx_Field_pos = _hx_reserved_func.Hx_Field_pos.VTable.Hx_Field_withOffset(this.Hx_Field_codeToLorscript.Hx_Field_inputText, this.Hx_Field_error.Hx_Field_pos.Hx_Field_offset, this.Hx_Field_error.Hx_Field_pos.Hx_Field_length, _hx_reserved_func.Hx_Field_pos.Hx_Field_offset)
										}
									}
								}

								{
									func() {
										defer func() {
											if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
												if e2, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_haxe_exception").(*Hx_Obj_haxe_exception); _hx_ok {
													_ = e2
													{
													}
												} else {
													panic(_hx_ex)
												}
											}
										}()
										{
											var parser *Hx_Obj_loreline_lorscript_parser
											_ = parser
											parser = Hx_Obj_loreline_lorscript_parser_CreateInstance()
											parser.Hx_Field_resumeErrors = true
											parser.Hx_Field_allowJSON = true
											parser.Hx_Field_allowTypes = true
											var origin HxNullable[string]
											_ = origin
											origin = _hx_reserved_func.Hx_Field_name
											var _hx_tmp_5 string
											_ = _hx_tmp_5
											_hx_tmp_5 = this.Hx_Field_lorscript
											var _hx_tmp_6 string
											_ = _hx_tmp_6
											if (origin).Valid != false {
												_hx_tmp_6 = (origin).Value
											} else {
												_hx_tmp_6 = "?"
											}

											this.Hx_Field_expr = parser.VTable.Hx_Field_parseString(_hx_tmp_5, HxNullableString(_hx_tmp_6), nil)
										}
									}()

								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					var parser *Hx_Obj_loreline_lorscript_parser
					_ = parser
					parser = Hx_Obj_loreline_lorscript_parser_CreateInstance()
					parser.Hx_Field_resumeErrors = false
					parser.Hx_Field_allowJSON = true
					parser.Hx_Field_allowTypes = true
					var origin HxNullable[string]
					_ = origin
					origin = _hx_reserved_func.Hx_Field_name
					var _hx_tmp_1 string
					_ = _hx_tmp_1
					_hx_tmp_1 = this.Hx_Field_lorscript
					var _hx_tmp_2 string
					_ = _hx_tmp_2
					if (origin).Valid != false {
						_hx_tmp_2 = (origin).Value
					} else {
						_hx_tmp_2 = "?"
					}

					this.Hx_Field_expr = parser.VTable.Hx_Field_parseString(_hx_tmp_1, HxNullableString(_hx_tmp_2), nil)
				}
			}()
			if hx_try_state == 1 {
				return
			}
		}
	}

}

var Hx_Obj_loreline_funclorscript_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.FuncLorscript",
	HxMakeArray[string](),
	HxMakeArray[string]("func", "codeToLorscript", "lorscript", "expr", "error"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_funclorscript_CreateInstance(HxConvert[*Hx_Obj_loreline_nfunctiondecl](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_funclorscript_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_funclorscript_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_funclorscript) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_funclorscript_RTTI
}

func (this *Hx_Obj_loreline_funclorscript) Hx_As_Hx_Obj_loreline_funclorscript() *Hx_Obj_loreline_funclorscript {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscriptcompletion() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscriptcompletion() {
}

type Hx_Obj_VTable_loreline_lorscriptcompletion interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscriptcompletion struct {
	VTable              Hx_Obj_VTable_loreline_lorscriptcompletion `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_locals     Hx_Typedef_map
	Hx_Field_completion *Hx_Obj_loreline_lorscript_completion
}

func Hx_Obj_loreline_lorscriptcompletion_CreateEmptyInstance() *Hx_Obj_loreline_lorscriptcompletion {
	obj := &Hx_Obj_loreline_lorscriptcompletion{}
	obj.VTable = obj
	obj.Hx_Field_locals = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_completion = HxDefault[*Hx_Obj_loreline_lorscript_completion]()
	return obj
}

func Hx_Obj_loreline_lorscriptcompletion_CreateInstance(locals__isOptional any, completion__isOptional any) *Hx_Obj_loreline_lorscriptcompletion {
	obj := Hx_Obj_loreline_lorscriptcompletion_CreateEmptyInstance()
	obj.Hx_New(locals__isOptional, completion__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_lorscriptcompletion) Hx_New(locals__isOptional any, completion__isOptional any) {
	{
		locals := HxDefault[Hx_Typedef_map]()
		_ = locals

		if locals__isOptional != nil {
			locals = HxConvert[Hx_Typedef_map](locals__isOptional)
		} else {
		}
		completion := HxDefault[*Hx_Obj_loreline_lorscript_completion]()
		_ = completion

		if completion__isOptional != nil {
			completion = HxConvert[*Hx_Obj_loreline_lorscript_completion](completion__isOptional)
		} else {
		}
		{
			this.Hx_Field_locals = nil
		}
		{
			this.Hx_Field_completion = nil
		}
		{
			if locals != nil {
				this.Hx_Field_locals = locals
			}

			if completion != nil {
				this.Hx_Field_completion = completion
			}
		}

	}
}

var Hx_Obj_loreline_lorscriptcompletion_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.LorscriptCompletion",
	HxMakeArray[string](),
	HxMakeArray[string]("locals", "completion"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscriptcompletion_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_lorscriptcompletion_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscriptcompletion_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscriptcompletion) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscriptcompletion_RTTI
}

func (this *Hx_Obj_loreline_lorscriptcompletion) Hx_As_Hx_Obj_loreline_lorscriptcompletion() *Hx_Obj_loreline_lorscriptcompletion {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_lens() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lens() {
}

type Hx_Obj_VTable_loreline_lens interface {
	Hx_Field_getVisibleBeats(fromNode *Hx_Obj_loreline_node) HxArray[*Hx_Obj_loreline_nbeatdecl]
	Hx_Field_getAllTags() HxArray[string]
	Hx_Field_isTopLevelNode(node *Hx_Obj_loreline_node) bool
	Hx_Field_getClosestNodeAtOrBeforePosition(pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node
	Hx_Field_resolveArrayAccess(access *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_node
	Hx_Field_findBeatFromAccess(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_findCallArgument(node *Hx_Obj_loreline_node) any
	Hx_Field_resolveLorscriptAccess(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, expr any) *Hx_Obj_loreline_node
	Hx_Field_getNodeById(id Hx_Typedef_haxe__int64___int64) *Hx_Obj_loreline_node
	Hx_Field_getDialogueSpeakers(node *Hx_Obj_loreline_node) HxArray[string]
	Hx_Field_countTags() Hx_Typedef_map
	Hx_Field_findStateField(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_nobjectfield
	Hx_Field_findOutboundBeats(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_findBeatCharacters(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_findReadCharacterFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_getNodeAncestorFilePaths(node *Hx_Obj_loreline_node) HxArray[string]
	Hx_Field_getNodeRelativeFilePath(node *Hx_Obj_loreline_node) string
	Hx_Field_findBeatFromTransition(transition *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_findTopLevelBeatFromNode(node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_findModifiedCharacterFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_isLorscriptExpr(input any) bool
	Hx_Field_resolveAccessInFunction(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node
	Hx_Field_getParentNode(node *Hx_Obj_loreline_node) *Hx_Obj_loreline_node
	Hx_Field_findBeatParamFromAccess(access *Hx_Obj_loreline_naccess) any
	Hx_Field_findBeatFromInsertion(insertion *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_getVisibleFunctions() HxArray[*Hx_Obj_loreline_nfunctiondecl]
	Hx_Field_findReferencesToBeat(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_getFuncLorscript(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) *Hx_Obj_loreline_funclorscript
	Hx_Field_getNodeAtPosition(pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node
	Hx_Field_getImportedPaths(rootPath string) HxArray[string]
	Hx_Field__getImportedPaths(rootPath string, script *Hx_Obj_loreline_script, result HxArray[string], used Hx_Typedef_map) HxArray[string]
	Hx_Field_findBeatByPathFromNode(path string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_getVisibleStateFields(fromNode *Hx_Obj_loreline_node) HxArray[*Hx_Obj_loreline_nobjectfield]
	Hx_Field_findReadStateFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_traverse(node *Hx_Obj_loreline_node, callback func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool)
	Hx_Field_initialize()
	Hx_Field_getNodesOfType(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any]
	Hx_Field_getNodeFilePath(node *Hx_Obj_loreline_node, rootPath string) string
	Hx_Field_findCharacterFromDialogue(dialogue *Hx_Obj_loreline_ndialoguestatement) *Hx_Obj_loreline_ncharacterdecl
	Hx_Field_findFunctionByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nfunctiondecl
	Hx_Field_getVisibleCharacters() HxArray[*Hx_Obj_loreline_ncharacterdecl]
	Hx_Field_findModifiedStateFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference]
	Hx_Field_getLorscriptExpr(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) any
	Hx_Field_findBeatByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_getLorscriptCompletion(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_lorscriptcompletion
	Hx_Field_getFirstParentOfType(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
	Hx_Field_resolveAccess(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_node
	Hx_Field_findCharacterByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_ncharacterdecl
	Hx_Field_getWhenWarnings(when *Hx_Obj_loreline_nwhenstatement) HxArray[any]
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lens struct {
	VTable                      Hx_Obj_VTable_loreline_lens `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_script             *Hx_Obj_loreline_script
	Hx_Field_nodesById          *Hx_Obj_loreline_nodeidmap
	Hx_Field_parentNodes        *Hx_Obj_loreline_nodeidmap
	Hx_Field_childNodes         *Hx_Obj_loreline_nodeidmap
	Hx_Field_lorscriptFunctions *Hx_Obj_loreline_nodeidmap
}

func Hx_Obj_loreline_lens_CreateEmptyInstance() *Hx_Obj_loreline_lens {
	obj := &Hx_Obj_loreline_lens{}
	obj.VTable = obj
	obj.Hx_Field_script = HxDefault[*Hx_Obj_loreline_script]()
	obj.Hx_Field_nodesById = HxDefault[*Hx_Obj_loreline_nodeidmap]()
	obj.Hx_Field_parentNodes = HxDefault[*Hx_Obj_loreline_nodeidmap]()
	obj.Hx_Field_childNodes = HxDefault[*Hx_Obj_loreline_nodeidmap]()
	obj.Hx_Field_lorscriptFunctions = HxDefault[*Hx_Obj_loreline_nodeidmap]()
	return obj
}

func Hx_Obj_loreline_lens_CreateInstance(script *Hx_Obj_loreline_script) *Hx_Obj_loreline_lens {
	obj := Hx_Obj_loreline_lens_CreateEmptyInstance()
	obj.Hx_New(script)
	return obj
}

func (this *Hx_Obj_loreline_lens) Hx_New(script *Hx_Obj_loreline_script) {
	{
		this.Hx_Field_nodesById = Hx_Obj_loreline_nodeidmap_CreateInstance()
	}
	{
		this.Hx_Field_parentNodes = Hx_Obj_loreline_nodeidmap_CreateInstance()
	}
	{
		this.Hx_Field_childNodes = Hx_Obj_loreline_nodeidmap_CreateInstance()
	}
	{
		this.Hx_Field_lorscriptFunctions = Hx_Obj_loreline_nodeidmap_CreateInstance()
	}
	{
		this.Hx_Field_script = script
		this.VTable.Hx_Field_initialize()
	}

}

var Hx_Obj_loreline_lens_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Lens",
	HxMakeArray[string]("leftmostOperand", "topLevelClauses", "comparedValue", "isNumberLiteral"),
	HxMakeArray[string]("script", "nodesById", "parentNodes", "childNodes", "lorscriptFunctions", "initialize", "getNodeById", "getNodeAtPosition", "getClosestNodeAtOrBeforePosition", "getNodesOfType", "getDialogueSpeakers", "getParentNode", "getFirstParentOfType", "getNodeFilePath", "getNodeAncestorFilePaths", "getNodeRelativeFilePath", "getImportedPaths", "_getImportedPaths", "resolveArrayAccess", "resolveAccess", "findBeatFromAccess", "findBeatParamFromAccess", "findCallArgument", "findBeatFromTransition", "findBeatFromInsertion", "findBeatByPathFromNode", "isTopLevelNode", "findTopLevelBeatFromNode", "findBeatByNameFromNode", "findCharacterFromDialogue", "findCharacterByNameFromNode", "getWhenWarnings", "findFunctionByNameFromNode", "getVisibleCharacters", "getVisibleFunctions", "getVisibleStateFields", "getVisibleBeats", "getAllTags", "countTags", "findStateField", "findOutboundBeats", "findReferencesToBeat", "findModifiedStateFields", "findReadStateFields", "findBeatCharacters", "findModifiedCharacterFields", "findReadCharacterFields", "traverse", "getFuncLorscript", "getLorscriptExpr", "resolveLorscriptAccess", "isLorscriptExpr", "getLorscriptCompletion", "resolveAccessInFunction"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lens_CreateInstance(HxConvert[*Hx_Obj_loreline_script](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lens_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lens_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "leftmostOperand":
			return Hx_Field_loreline_lens_leftmostOperand
		case "topLevelClauses":
			return Hx_Field_loreline_lens_topLevelClauses
		case "comparedValue":
			return Hx_Field_loreline_lens_comparedValue
		case "isNumberLiteral":
			return Hx_Field_loreline_lens_isNumberLiteral
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lens) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lens_RTTI
}

func (this *Hx_Obj_loreline_lens) Hx_As_Hx_Obj_loreline_lens() *Hx_Obj_loreline_lens {
	return this
}

func (this *Hx_Obj_loreline_lens) Hx_Field_initialize() {
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		{
			var key Hx_Typedef_haxe__int64___int64
			_ = key
			key = node.Hx_Field_id
			var _hx_tmp_0 int64
			_ = _hx_tmp_0
			_hx_tmp_0 = ((int64)(key))
			func() func(high int32, low int32, value *Hx_Obj_loreline_node) {
				var _hx_closure_src_473 func(high int32, low int32, value any)
				_ = _hx_closure_src_473
				_hx_closure_src_473 = _gthis.Hx_Field_nodesById.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_node) {
					_hx_closure_src_473(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key)), node)
		}

		if parent != nil {
			{
				var key Hx_Typedef_haxe__int64___int64
				_ = key
				key = node.Hx_Field_id
				var _hx_tmp_1 int64
				_ = _hx_tmp_1
				_hx_tmp_1 = ((int64)(key))
				func() func(high int32, low int32, value *Hx_Obj_loreline_node) {
					var _hx_closure_src_474 func(high int32, low int32, value any)
					_ = _hx_closure_src_474
					_hx_closure_src_474 = _gthis.Hx_Field_parentNodes.Hx_Field_map.VTable.Hx_Field__set
					return func(high int32, low int32, value *Hx_Obj_loreline_node) {
						_hx_closure_src_474(high, low, ((any)(value)))
					}
				}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)), parent)
			}

			var key1 Hx_Typedef_haxe__int64___int64
			_ = key1
			key1 = parent.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(key1))
			var children HxArray[*Hx_Obj_loreline_node]
			_ = children
			children = func() func(high int32, low int32) HxArray[*Hx_Obj_loreline_node] {
				var _hx_closure_src_475 func(high int32, low int32) any
				_ = _hx_closure_src_475
				_hx_closure_src_475 = _gthis.Hx_Field_childNodes.Hx_Field_map.VTable.Hx_Field__get
				return func(high int32, low int32) HxArray[*Hx_Obj_loreline_node] {
					return HxAnyToTypedArray[*Hx_Obj_loreline_node](_hx_closure_src_475(high, low))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(key1)))
			if children == nil {
				children = HxMakeArray[*Hx_Obj_loreline_node]()
				{
					var key Hx_Typedef_haxe__int64___int64
					_ = key
					key = parent.Hx_Field_id
					var _hx_tmp_3 int64
					_ = _hx_tmp_3
					_hx_tmp_3 = ((int64)(key))
					func() func(high int32, low int32, value HxArray[*Hx_Obj_loreline_node]) {
						var _hx_closure_src_476 func(high int32, low int32, value any)
						_ = _hx_closure_src_476
						_hx_closure_src_476 = _gthis.Hx_Field_childNodes.Hx_Field_map.VTable.Hx_Field__set
						return func(high int32, low int32, value HxArray[*Hx_Obj_loreline_node]) {
							_hx_closure_src_476(high, low, ((any)(value)))
						}
					}()(((int32)((_hx_tmp_3 >> int64(((int32)(32)))))), ((int32)(key)), children)
				}
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_node](children, node))
		}
	})
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodeById(id Hx_Typedef_haxe__int64___int64) *Hx_Obj_loreline_node {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	return func() func(high int32, low int32) *Hx_Obj_loreline_node {
		var _hx_closure_src_477 func(high int32, low int32) any
		_ = _hx_closure_src_477
		_hx_closure_src_477 = this.Hx_Field_nodesById.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_node {
			var hx_dyncast_131 any
			_ = hx_dyncast_131
			hx_dyncast_131 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_477(high, low), "Hx_Obj_loreline_node")
			var _hx_tmp_0 *Hx_Obj_loreline_node
			_ = _hx_tmp_0
			if hx_dyncast_131 != nil {
				_hx_tmp_0 = hx_dyncast_131.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodeAtPosition(pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node {
	var bestMatch *Hx_Obj_loreline_node
	_ = bestMatch
	this.Hx_Field_script.VTable.Hx_Field_eachExcludingImported(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var nodePos *Hx_Obj_loreline_position
		_ = nodePos
		nodePos = node.Hx_Field_pos
		if ((nodePos.Hx_Field_length > 0) && (nodePos.Hx_Field_offset <= pos.Hx_Field_offset)) && ((nodePos.Hx_Field_offset + nodePos.Hx_Field_length) >= pos.Hx_Field_offset) {
			bestMatch = node
		}
	})
	return bestMatch
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getClosestNodeAtOrBeforePosition(pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node {
	var bestMatch *Hx_Obj_loreline_node
	_ = bestMatch
	var bestDistance int32
	_ = bestDistance
	bestDistance = 999999999
	this.Hx_Field_script.VTable.Hx_Field_eachExcludingImported(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var nodePos *Hx_Obj_loreline_position
		_ = nodePos
		nodePos = node.Hx_Field_pos
		var distance int32
		_ = distance
		distance = (pos.Hx_Field_offset - nodePos.Hx_Field_offset)
		if (distance >= 0) && (distance < bestDistance) {
			bestDistance = distance
			bestMatch = node
		}
	})
	return bestMatch
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodesOfType(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any] {
	includeImported := HxDefault[bool]()
	_ = includeImported

	if includeImported__isOptional != nil {
		includeImported = HxConvert[bool](includeImported__isOptional)
	} else {
		includeImported = false
	}
	{
		var matches HxArray[any]
		_ = matches
		matches = HxMakeArray[any]().Dyn()
		var _gthis *Hx_Obj_loreline_lens
		_ = _gthis
		_gthis = this
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
			if Hx_Field_std_isOfType(node, nodeType) {
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](matches), ((any)(node))))
			}

			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if includeImported {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 *Hx_Obj_loreline_node
				_ = _hx_tmp_3
				_hx_tmp_3 = node
				var _hx_tmp_4 *Hx_Obj_loreline_node
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 == (func() *Hx_Obj_loreline_node {
					_hx_up := _gthis.Hx_Field_script
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = true
			} else {
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nimportstatement_RTTI) {
					_hx_tmp_5 = true
				} else {
					_hx_tmp_5 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_script_RTTI)
				}

				_hx_tmp_1 = !(_hx_tmp_5)
			}

			return _hx_tmp_1
		})
		return matches
	}
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getDialogueSpeakers(node *Hx_Obj_loreline_node) HxArray[string] {
	var speakers HxArray[string]
	_ = speakers
	speakers = HxMakeArray[string]()
	var seen *Hx_Obj_haxe_imap
	_ = seen
	seen = (func() *Hx_Obj_haxe_imap {
		_hx_icast_132 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_132 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_132, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_132.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = node
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(child *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ndialoguestatement_RTTI) {
			var dialogue *Hx_Obj_loreline_ndialoguestatement
			_ = dialogue
			dialogue = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
			}).Hx_As_Hx_Obj_loreline_ndialoguestatement()
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if dialogue.Hx_Field_character != HxStringNull {
				_hx_tmp_1 = !func() func(_hx_param_key string) bool {
					var _hx_closure_src_478 func(_hx_param_key any) bool
					_ = _hx_closure_src_478
					_hx_closure_src_478 = seen.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_478(((any)(_hx_param_key)))
					}
				}()(dialogue.Hx_Field_character)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				func() func(_hx_param_key string, _hx_param_value bool) {
					var _hx_closure_src_479 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_479
					_hx_closure_src_479 = seen.VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value bool) {
						_hx_closure_src_479(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(dialogue.Hx_Field_character, true)
				_ = (Hx_Array_Push[string](speakers, dialogue.Hx_Field_character))
			}
		}

		return true
	})
	return speakers
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getParentNode(node *Hx_Obj_loreline_node) *Hx_Obj_loreline_node {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = node.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	return func() func(high int32, low int32) *Hx_Obj_loreline_node {
		var _hx_closure_src_480 func(high int32, low int32) any
		_ = _hx_closure_src_480
		_hx_closure_src_480 = this.Hx_Field_parentNodes.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_node {
			var hx_dyncast_133 any
			_ = hx_dyncast_133
			hx_dyncast_133 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_480(high, low), "Hx_Obj_loreline_node")
			var _hx_tmp_0 *Hx_Obj_loreline_node
			_ = _hx_tmp_0
			if hx_dyncast_133 != nil {
				_hx_tmp_0 = hx_dyncast_133.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getFirstParentOfType(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any {
	var current any
	_ = current
	current = ((any)(node))
	for {
		if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(current, nil)))) {
			break
		}

		var hx_dyncast_135 any
		_ = hx_dyncast_135
		hx_dyncast_135 = Hx_Field_go_haxe_hxdynamic_castClass(current, "Hx_Obj_loreline_node")
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		if hx_dyncast_135 != nil {
			_hx_tmp_0 = hx_dyncast_135.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
		} else {
			_hx_tmp_0 = nil
		}

		current = ((any)(this.VTable.Hx_Field_getParentNode(_hx_tmp_0)))
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(current, nil))) {
			_hx_tmp_1 = (((*Hx_Obj_go_haxe_hxclass)(Hx_Field_type_getClass(current))) == _hx_reserved_type)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			return current
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodeFilePath(node *Hx_Obj_loreline_node, rootPath string) string {
	var importChain HxArray[*Hx_Obj_loreline_nimportstatement]
	_ = importChain
	importChain = HxMakeArray[*Hx_Obj_loreline_nimportstatement]()
	var current *Hx_Obj_loreline_node
	_ = current
	current = node
	for current != nil {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
			var _hx_closure_src_481 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_481
			_hx_closure_src_481 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
				var hx_dyncast_136 any
				_ = hx_dyncast_136
				hx_dyncast_136 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_481(node, _hx_reserved_type), "Hx_Obj_loreline_nimportstatement")
				var _hx_tmp_0 *Hx_Obj_loreline_nimportstatement
				_ = _hx_tmp_0
				if hx_dyncast_136 != nil {
					_hx_tmp_0 = hx_dyncast_136.(interface {
						Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
					}).Hx_As_Hx_Obj_loreline_nimportstatement()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(current, ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nimportstatement_RTTI)))
		if importStmt != nil {
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nimportstatement](importChain, importStmt))
			current = (func() *Hx_Obj_loreline_node {
				_hx_up := importStmt
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		} else {
			break
		}
	}

	if importChain.Len() == 0 {
		return rootPath
	}

	var currentDir string
	_ = currentDir
	currentDir = Hx_Field_haxe_io_path_directory(rootPath)
	var ext bool
	_ = ext
	if rootPath == HxStringNull {
		ext = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(rootPath))
		if n < 8 {
			ext = false
		} else {
			var ext1 bool
			_ = ext1
			var ext2 bool
			_ = ext2
			var ext3 bool
			_ = ext3
			var ext4 bool
			_ = ext4
			var ext5 bool
			_ = ext5
			var ext6 bool
			_ = ext6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 8)))
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			if (c).Valid == false {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = (c).Value
			}

			if (_hx_tmp_1) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 7)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				ext6 = ((c1 == 108) || (c1 == 76))
			} else {
				ext6 = false
			}

			if ext6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 6)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				ext5 = ((c1 == 111) || (c1 == 79))
			} else {
				ext5 = false
			}

			if ext5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 5)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_4
				ext4 = ((c1 == 114) || (c1 == 82))
			} else {
				ext4 = false
			}

			if ext4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 4)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				ext3 = ((_hx_tmp_5) == 46)
			} else {
				ext3 = false
			}

			if ext3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 3)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				ext2 = ((c1 == 116) || (c1 == 84))
			} else {
				ext2 = false
			}

			if ext2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 2)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				ext1 = ((c1 == 120) || (c1 == 88))
			} else {
				ext1 = false
			}

			if ext1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 1)))
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				if (tmp_c_1).Valid == false {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_8
				ext = ((c1 == 116) || (c1 == 84))
			} else {
				ext = false
			}
		}
	}

	var _hx_tmp_9 string
	_ = _hx_tmp_9
	if ext {
		_hx_tmp_9 = ".lor.txt"
	} else {
		_hx_tmp_9 = ".lor"
	}

	var ext1 string
	_ = ext1
	ext1 = _hx_tmp_9
	var resolvedPath string
	_ = resolvedPath
	resolvedPath = rootPath
	var i int32
	_ = i
	i = (importChain.Len() - 1)
	for i >= 0 {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = importChain.Get(i)
		var importPath string
		_ = importPath
		var _g Hx_Enum_loreline_stringparttype
		_ = _g
		_g = importStmt.Hx_Field_path.Hx_Field_parts.Get(0).Hx_Field_partType
		if _g.Hx_Field_enumIndex() == 0 {
			var text string
			_ = text
			text = _g.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
			importPath = text
		} else {
			importPath = ""
		}

		if !Hx_Field_haxe_io_path_isAbsolute(importPath) {
			importPath = Hx_Field_haxe_io_path_join(HxMakeArray[string](currentDir, importPath))
		}

		importPath = Hx_Field_haxe_io_path_normalize(importPath)
		var tmp bool
		_ = tmp
		var tmp1 bool
		_ = tmp1
		if importPath == HxStringNull {
			tmp1 = false
		} else {
			var n int32
			_ = n
			n = int32(utf8.RuneCountInString(importPath))
			if n < 8 {
				tmp1 = false
			} else {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				var tmp2 bool
				_ = tmp2
				var tmp3 bool
				_ = tmp3
				var tmp4 bool
				_ = tmp4
				var tmp5 bool
				_ = tmp5
				var tmp6 bool
				_ = tmp6
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 8)))
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				if (c).Valid == false {
					_hx_tmp_10 = -1
				} else {
					_hx_tmp_10 = (c).Value
				}

				if (_hx_tmp_10) == 46 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 7)))
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					if (tmp_c_1).Valid == false {
						_hx_tmp_11 = -1
					} else {
						_hx_tmp_11 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_11
					tmp6 = ((c1 == 108) || (c1 == 76))
				} else {
					tmp6 = false
				}

				if tmp6 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 6)))
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					if (tmp_c_1).Valid == false {
						_hx_tmp_12 = -1
					} else {
						_hx_tmp_12 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_12
					tmp5 = ((c1 == 111) || (c1 == 79))
				} else {
					tmp5 = false
				}

				if tmp5 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 5)))
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					if (tmp_c_1).Valid == false {
						_hx_tmp_13 = -1
					} else {
						_hx_tmp_13 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_13
					tmp4 = ((c1 == 114) || (c1 == 82))
				} else {
					tmp4 = false
				}

				if tmp4 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 4)))
					var _hx_tmp_14 int32
					_ = _hx_tmp_14
					if (tmp_c_1).Valid == false {
						_hx_tmp_14 = -1
					} else {
						_hx_tmp_14 = (tmp_c_1).Value
					}

					tmp3 = ((_hx_tmp_14) == 46)
				} else {
					tmp3 = false
				}

				if tmp3 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
					var _hx_tmp_15 int32
					_ = _hx_tmp_15
					if (tmp_c_1).Valid == false {
						_hx_tmp_15 = -1
					} else {
						_hx_tmp_15 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_15
					tmp2 = ((c1 == 116) || (c1 == 84))
				} else {
					tmp2 = false
				}

				if tmp2 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					if (tmp_c_1).Valid == false {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_16
					tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
				} else {
					tmp_tmp_1 = false
				}

				if tmp_tmp_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					if (tmp_c_1).Valid == false {
						_hx_tmp_17 = -1
					} else {
						_hx_tmp_17 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_17
					tmp1 = ((c1 == 116) || (c1 == 84))
				} else {
					tmp1 = false
				}
			}
		}

		if !tmp1 {
			if importPath == HxStringNull {
				tmp = false
			} else {
				var n int32
				_ = n
				n = int32(utf8.RuneCountInString(importPath))
				if n < 4 {
					tmp = false
				} else {
					var tmp_tmp1_1 bool
					_ = tmp_tmp1_1
					var tmp2 bool
					_ = tmp2
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 4)))
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					if (c).Valid == false {
						_hx_tmp_18 = -1
					} else {
						_hx_tmp_18 = (c).Value
					}

					if (_hx_tmp_18) == 46 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						if (tmp_c_1).Valid == false {
							_hx_tmp_19 = -1
						} else {
							_hx_tmp_19 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_19
						tmp2 = ((c1 == 108) || (c1 == 76))
					} else {
						tmp2 = false
					}

					if tmp2 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						if (tmp_c_1).Valid == false {
							_hx_tmp_20 = -1
						} else {
							_hx_tmp_20 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_20
						tmp_tmp1_1 = ((c1 == 111) || (c1 == 79))
					} else {
						tmp_tmp1_1 = false
					}

					if tmp_tmp1_1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
						var _hx_tmp_21 int32
						_ = _hx_tmp_21
						if (tmp_c_1).Valid == false {
							_hx_tmp_21 = -1
						} else {
							_hx_tmp_21 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_21
						tmp = ((c1 == 114) || (c1 == 82))
					} else {
						tmp = false
					}
				}
			}
		} else {
			tmp = true
		}

		if !tmp {
			importPath = Hx_Field_go_haxe_hxdynamic_stringConcat(importPath, ext1)
		}

		resolvedPath = importPath
		currentDir = Hx_Field_haxe_io_path_directory(resolvedPath)
		i--
	}

	return resolvedPath
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodeAncestorFilePaths(node *Hx_Obj_loreline_node) HxArray[string] {
	var importChain HxArray[*Hx_Obj_loreline_nimportstatement]
	_ = importChain
	importChain = HxMakeArray[*Hx_Obj_loreline_nimportstatement]()
	var current *Hx_Obj_loreline_node
	_ = current
	current = node
	for current != nil {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
			var _hx_closure_src_482 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_482
			_hx_closure_src_482 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
				var hx_dyncast_137 any
				_ = hx_dyncast_137
				hx_dyncast_137 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_482(node, _hx_reserved_type), "Hx_Obj_loreline_nimportstatement")
				var _hx_tmp_0 *Hx_Obj_loreline_nimportstatement
				_ = _hx_tmp_0
				if hx_dyncast_137 != nil {
					_hx_tmp_0 = hx_dyncast_137.(interface {
						Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
					}).Hx_As_Hx_Obj_loreline_nimportstatement()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(current, ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nimportstatement_RTTI)))
		if importStmt != nil {
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nimportstatement](importChain, importStmt))
			current = (func() *Hx_Obj_loreline_node {
				_hx_up := importStmt
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		} else {
			break
		}
	}

	var result HxArray[string]
	_ = result
	result = HxMakeArray[string]()
	var resolved string
	_ = resolved
	resolved = HxStringNull
	var i int32
	_ = i
	i = (importChain.Len() - 1)
	for i >= 0 {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = importChain.Get(i)
		var importPath string
		_ = importPath
		var _g Hx_Enum_loreline_stringparttype
		_ = _g
		_g = importStmt.Hx_Field_path.Hx_Field_parts.Get(0).Hx_Field_partType
		if _g.Hx_Field_enumIndex() == 0 {
			var text string
			_ = text
			text = _g.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
			importPath = text
		} else {
			importPath = ""
		}

		if resolved == HxStringNull {
			resolved = importPath
		} else {
			resolved = Hx_Field_haxe_io_path_join(HxMakeArray[string](Hx_Field_haxe_io_path_directory(resolved), importPath))
		}

		var path string
		_ = path
		path = Hx_Field_haxe_io_path_normalize(resolved)
		var tmp string
		_ = tmp
		if path == HxStringNull {
			tmp = HxStringNull
		} else {
			var tmp1 bool
			_ = tmp1
			if path == HxStringNull {
				tmp1 = false
			} else {
				var n int32
				_ = n
				n = int32(utf8.RuneCountInString(path))
				if n < 8 {
					tmp1 = false
				} else {
					var tmp_tmp_1 bool
					_ = tmp_tmp_1
					var tmp2 bool
					_ = tmp2
					var tmp3 bool
					_ = tmp3
					var tmp4 bool
					_ = tmp4
					var tmp5 bool
					_ = tmp5
					var tmp6 bool
					_ = tmp6
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					if (c).Valid == false {
						_hx_tmp_1 = -1
					} else {
						_hx_tmp_1 = (c).Value
					}

					if (_hx_tmp_1) == 46 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						if (tmp_c_1).Valid == false {
							_hx_tmp_2 = -1
						} else {
							_hx_tmp_2 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_2
						tmp6 = ((c1 == 108) || (c1 == 76))
					} else {
						tmp6 = false
					}

					if tmp6 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						if (tmp_c_1).Valid == false {
							_hx_tmp_3 = -1
						} else {
							_hx_tmp_3 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_3
						tmp5 = ((c1 == 111) || (c1 == 79))
					} else {
						tmp5 = false
					}

					if tmp5 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						if (tmp_c_1).Valid == false {
							_hx_tmp_4 = -1
						} else {
							_hx_tmp_4 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_4
						tmp4 = ((c1 == 114) || (c1 == 82))
					} else {
						tmp4 = false
					}

					if tmp4 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						if (tmp_c_1).Valid == false {
							_hx_tmp_5 = -1
						} else {
							_hx_tmp_5 = (tmp_c_1).Value
						}

						tmp3 = ((_hx_tmp_5) == 46)
					} else {
						tmp3 = false
					}

					if tmp3 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						if (tmp_c_1).Valid == false {
							_hx_tmp_6 = -1
						} else {
							_hx_tmp_6 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_6
						tmp2 = ((c1 == 116) || (c1 == 84))
					} else {
						tmp2 = false
					}

					if tmp2 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						if (tmp_c_1).Valid == false {
							_hx_tmp_7 = -1
						} else {
							_hx_tmp_7 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_7
						tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
					} else {
						tmp_tmp_1 = false
					}

					if tmp_tmp_1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						if (tmp_c_1).Valid == false {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_8
						tmp1 = ((c1 == 116) || (c1 == 84))
					} else {
						tmp1 = false
					}
				}
			}

			if tmp1 {
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 8), Valid: true}

				var _hx_tmp_9 string
				_ = _hx_tmp_9
				if (_hx_reserved_len).Valid != false {
					_hx_tmp_9 = (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
				} else {
					var _hx_tmp_10 string
					_ = _hx_tmp_10
					_hx_tmp_10 = path
					var _hx_tmp_11 string
					_ = _hx_tmp_11
					_hx_tmp_11 = _hx_tmp_10
					_hx_tmp_9 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_11, 0, HxNullable[int32]{}))
				}

				tmp = _hx_tmp_9
			} else {
				var tmp_tmp1_1 bool
				_ = tmp_tmp1_1
				if path == HxStringNull {
					tmp_tmp1_1 = false
				} else {
					var n int32
					_ = n
					n = int32(utf8.RuneCountInString(path))
					if n < 4 {
						tmp_tmp1_1 = false
					} else {
						var tmp_tmp_1 bool
						_ = tmp_tmp_1
						var tmp2 bool
						_ = tmp2
						var c HxNullable[int32]
						_ = c
						c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						if (c).Valid == false {
							_hx_tmp_12 = -1
						} else {
							_hx_tmp_12 = (c).Value
						}

						if (_hx_tmp_12) == 46 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
							var _hx_tmp_13 int32
							_ = _hx_tmp_13
							if (tmp_c_1).Valid == false {
								_hx_tmp_13 = -1
							} else {
								_hx_tmp_13 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_13
							tmp2 = ((c1 == 108) || (c1 == 76))
						} else {
							tmp2 = false
						}

						if tmp2 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							if (tmp_c_1).Valid == false {
								_hx_tmp_14 = -1
							} else {
								_hx_tmp_14 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_14
							tmp_tmp_1 = ((c1 == 111) || (c1 == 79))
						} else {
							tmp_tmp_1 = false
						}

						if tmp_tmp_1 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							if (tmp_c_1).Valid == false {
								_hx_tmp_15 = -1
							} else {
								_hx_tmp_15 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_15
							tmp_tmp1_1 = ((c1 == 114) || (c1 == 82))
						} else {
							tmp_tmp1_1 = false
						}
					}
				}

				if tmp_tmp1_1 {
					var _hx_reserved_len HxNullable[int32]
					_ = _hx_reserved_len
					_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 4), Valid: true}

					var _hx_tmp_16 string
					_ = _hx_tmp_16
					if (_hx_reserved_len).Valid != false {
						_hx_tmp_16 = (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
					} else {
						var _hx_tmp_17 string
						_ = _hx_tmp_17
						_hx_tmp_17 = path
						var _hx_tmp_18 string
						_ = _hx_tmp_18
						_hx_tmp_18 = _hx_tmp_17
						_hx_tmp_16 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_18, 0, HxNullable[int32]{}))
					}

					tmp = _hx_tmp_16
				} else {
					tmp = path
				}
			}
		}

		_ = (Hx_Array_Push[string](result, tmp))
		i--
	}

	(Hx_Array_Reverse[string](result))
	_ = (Hx_Array_Push[string](result, "."))
	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getNodeRelativeFilePath(node *Hx_Obj_loreline_node) string {
	var importChain HxArray[*Hx_Obj_loreline_nimportstatement]
	_ = importChain
	var current *Hx_Obj_loreline_node
	_ = current
	current = node
	for current != nil {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
			var _hx_closure_src_483 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_483
			_hx_closure_src_483 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nimportstatement {
				var hx_dyncast_138 any
				_ = hx_dyncast_138
				hx_dyncast_138 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_483(node, _hx_reserved_type), "Hx_Obj_loreline_nimportstatement")
				var _hx_tmp_0 *Hx_Obj_loreline_nimportstatement
				_ = _hx_tmp_0
				if hx_dyncast_138 != nil {
					_hx_tmp_0 = hx_dyncast_138.(interface {
						Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
					}).Hx_As_Hx_Obj_loreline_nimportstatement()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(current, ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nimportstatement_RTTI)))
		if importStmt != nil {
			if importChain == nil {
				importChain = HxMakeArray[*Hx_Obj_loreline_nimportstatement]()
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_nimportstatement](importChain, importStmt))
			current = (func() *Hx_Obj_loreline_node {
				_hx_up := importStmt
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		} else {
			break
		}
	}

	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if importChain == nil {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = (importChain.Len() == 0)
	}

	if _hx_tmp_1 {
		return "."
	}

	var resolved string
	_ = resolved
	resolved = HxStringNull
	var i int32
	_ = i
	i = (importChain.Len() - 1)
	for i >= 0 {
		var importStmt *Hx_Obj_loreline_nimportstatement
		_ = importStmt
		importStmt = importChain.Get(i)
		var importPath string
		_ = importPath
		var _g Hx_Enum_loreline_stringparttype
		_ = _g
		_g = importStmt.Hx_Field_path.Hx_Field_parts.Get(0).Hx_Field_partType
		if _g.Hx_Field_enumIndex() == 0 {
			var text string
			_ = text
			text = _g.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
			importPath = text
		} else {
			importPath = ""
		}

		if resolved == HxStringNull {
			resolved = importPath
		} else {
			resolved = Hx_Field_haxe_io_path_join(HxMakeArray[string](Hx_Field_haxe_io_path_directory(resolved), importPath))
		}

		i--
	}

	var path string
	_ = path
	path = Hx_Field_haxe_io_path_normalize(resolved)
	if path == HxStringNull {
		return HxStringNull
	} else {
		var tmp bool
		_ = tmp
		if path == HxStringNull {
			tmp = false
		} else {
			var n int32
			_ = n
			n = int32(utf8.RuneCountInString(path))
			if n < 8 {
				tmp = false
			} else {
				var tmp1 bool
				_ = tmp1
				var tmp2 bool
				_ = tmp2
				var tmp3 bool
				_ = tmp3
				var tmp4 bool
				_ = tmp4
				var tmp5 bool
				_ = tmp5
				var tmp6 bool
				_ = tmp6
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (c).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (c).Value
				}

				if (_hx_tmp_2) == 46 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					if (tmp_c_1).Valid == false {
						_hx_tmp_3 = -1
					} else {
						_hx_tmp_3 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_3
					tmp6 = ((c1 == 108) || (c1 == 76))
				} else {
					tmp6 = false
				}

				if tmp6 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					if (tmp_c_1).Valid == false {
						_hx_tmp_4 = -1
					} else {
						_hx_tmp_4 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_4
					tmp5 = ((c1 == 111) || (c1 == 79))
				} else {
					tmp5 = false
				}

				if tmp5 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					if (tmp_c_1).Valid == false {
						_hx_tmp_5 = -1
					} else {
						_hx_tmp_5 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_5
					tmp4 = ((c1 == 114) || (c1 == 82))
				} else {
					tmp4 = false
				}

				if tmp4 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					if (tmp_c_1).Valid == false {
						_hx_tmp_6 = -1
					} else {
						_hx_tmp_6 = (tmp_c_1).Value
					}

					tmp3 = ((_hx_tmp_6) == 46)
				} else {
					tmp3 = false
				}

				if tmp3 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					if (tmp_c_1).Valid == false {
						_hx_tmp_7 = -1
					} else {
						_hx_tmp_7 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_7
					tmp2 = ((c1 == 116) || (c1 == 84))
				} else {
					tmp2 = false
				}

				if tmp2 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					if (tmp_c_1).Valid == false {
						_hx_tmp_8 = -1
					} else {
						_hx_tmp_8 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_8
					tmp1 = ((c1 == 120) || (c1 == 88))
				} else {
					tmp1 = false
				}

				if tmp1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					if (tmp_c_1).Valid == false {
						_hx_tmp_9 = -1
					} else {
						_hx_tmp_9 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_9
					tmp = ((c1 == 116) || (c1 == 84))
				} else {
					tmp = false
				}
			}
		}

		if tmp {
			var _hx_reserved_len HxNullable[int32]
			_ = _hx_reserved_len
			_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 8), Valid: true}

			if (_hx_reserved_len).Valid != false {
				return (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
			} else {
				var _hx_tmp_10 string
				_ = _hx_tmp_10
				_hx_tmp_10 = path
				return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_10, 0, HxNullable[int32]{}))
			}
		} else {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if path == HxStringNull {
				tmp_tmp_1 = false
			} else {
				var n int32
				_ = n
				n = int32(utf8.RuneCountInString(path))
				if n < 4 {
					tmp_tmp_1 = false
				} else {
					var tmp1 bool
					_ = tmp1
					var tmp2 bool
					_ = tmp2
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					if (c).Valid == false {
						_hx_tmp_11 = -1
					} else {
						_hx_tmp_11 = (c).Value
					}

					if (_hx_tmp_11) == 46 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						if (tmp_c_1).Valid == false {
							_hx_tmp_12 = -1
						} else {
							_hx_tmp_12 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_12
						tmp2 = ((c1 == 108) || (c1 == 76))
					} else {
						tmp2 = false
					}

					if tmp2 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						if (tmp_c_1).Valid == false {
							_hx_tmp_13 = -1
						} else {
							_hx_tmp_13 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_13
						tmp1 = ((c1 == 111) || (c1 == 79))
					} else {
						tmp1 = false
					}

					if tmp1 {
						var tmp_c_1 HxNullable[int32]
						_ = tmp_c_1
						tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						if (tmp_c_1).Valid == false {
							_hx_tmp_14 = -1
						} else {
							_hx_tmp_14 = (tmp_c_1).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_14
						tmp_tmp_1 = ((c1 == 114) || (c1 == 82))
					} else {
						tmp_tmp_1 = false
					}
				}
			}

			if tmp_tmp_1 {
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(path)) - 4), Valid: true}

				if (_hx_reserved_len).Valid != false {
					return (Hx_Field_go_haxe_hxstring_substr(path, 0, _hx_reserved_len))
				} else {
					var _hx_tmp_15 string
					_ = _hx_tmp_15
					_hx_tmp_15 = path
					return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, 0, HxNullable[int32]{}))
				}
			} else {
				return path
			}
		}
	}
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getImportedPaths(rootPath string) HxArray[string] {
	var result HxArray[string]
	_ = result
	result = HxMakeArray[string]()
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = rootPath
	var _hx_tmp_1 *Hx_Obj_loreline_script
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_script
	var _hx_tmp_2 HxArray[string]
	_ = _hx_tmp_2
	_hx_tmp_2 = result
	this.VTable.Hx_Field__getImportedPaths(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, (func() *Hx_Obj_haxe_imap {
		_hx_icast_139 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_139 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_139, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_139.VTable.(Hx_Obj_VTable_haxe_imap)}
	})())
	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field__getImportedPaths(rootPath string, script *Hx_Obj_loreline_script, result HxArray[string], used Hx_Typedef_map) HxArray[string] {
	var rootDir string
	_ = rootDir
	rootDir = Hx_Field_haxe_io_path_directory(rootPath)
	if script == nil {
		script = this.Hx_Field_script
	}

	var ext bool
	_ = ext
	if rootPath == HxStringNull {
		ext = false
	} else {
		var n int32
		_ = n
		n = int32(utf8.RuneCountInString(rootPath))
		if n < 8 {
			ext = false
		} else {
			var ext1 bool
			_ = ext1
			var ext2 bool
			_ = ext2
			var ext3 bool
			_ = ext3
			var ext4 bool
			_ = ext4
			var ext5 bool
			_ = ext5
			var ext6 bool
			_ = ext6
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 8)))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 46 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 7)))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_1
				ext6 = ((c1 == 108) || (c1 == 76))
			} else {
				ext6 = false
			}

			if ext6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 6)))
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				if (tmp_c_1).Valid == false {
					_hx_tmp_2 = -1
				} else {
					_hx_tmp_2 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_2
				ext5 = ((c1 == 111) || (c1 == 79))
			} else {
				ext5 = false
			}

			if ext5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 5)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (tmp_c_1).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_3
				ext4 = ((c1 == 114) || (c1 == 82))
			} else {
				ext4 = false
			}

			if ext4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				ext3 = ((_hx_tmp_4) == 46)
			} else {
				ext3 = false
			}

			if ext3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 3)))
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (tmp_c_1).Valid == false {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_5
				ext2 = ((c1 == 116) || (c1 == 84))
			} else {
				ext2 = false
			}

			if ext2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 2)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (tmp_c_1).Valid == false {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_6
				ext1 = ((c1 == 120) || (c1 == 88))
			} else {
				ext1 = false
			}

			if ext1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(rootPath, (n - 1)))
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				if (tmp_c_1).Valid == false {
					_hx_tmp_7 = -1
				} else {
					_hx_tmp_7 = (tmp_c_1).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_7
				ext = ((c1 == 116) || (c1 == 84))
			} else {
				ext = false
			}
		}
	}

	var _hx_tmp_8 string
	_ = _hx_tmp_8
	if ext {
		_hx_tmp_8 = ".lor.txt"
	} else {
		_hx_tmp_8 = ".lor"
	}

	var ext1 string
	_ = ext1
	ext1 = _hx_tmp_8
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g1
		_g1 = script.Hx_Field_body
		for {
			var _hx_tmp_9 int32
			_ = _hx_tmp_9
			_hx_tmp_9 = _g
			if !(_hx_tmp_9 < _g1.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g1.Get(_g)
			_g++
			if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nimportstatement_RTTI) {
				var importNode *Hx_Obj_loreline_nimportstatement
				_ = importNode
				importNode = node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
				}).Hx_As_Hx_Obj_loreline_nimportstatement()
				var importPath string
				_ = importPath
				var tmp__g_1 Hx_Enum_loreline_stringparttype
				_ = tmp__g_1
				tmp__g_1 = importNode.Hx_Field_path.Hx_Field_parts.Get(0).Hx_Field_partType
				if tmp__g_1.Hx_Field_enumIndex() == 0 {
					var text string
					_ = text
					text = tmp__g_1.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
					importPath = text
				} else {
					importPath = ""
				}

				if !Hx_Field_haxe_io_path_isAbsolute(importPath) {
					importPath = Hx_Field_haxe_io_path_join(HxMakeArray[string](rootDir, importPath))
				}

				importPath = Hx_Field_haxe_io_path_normalize(importPath)
				var tmp bool
				_ = tmp
				var tmp1 bool
				_ = tmp1
				if importPath == HxStringNull {
					tmp1 = false
				} else {
					var n int32
					_ = n
					n = int32(utf8.RuneCountInString(importPath))
					if n < 8 {
						tmp1 = false
					} else {
						var tmp_tmp_1 bool
						_ = tmp_tmp_1
						var tmp2 bool
						_ = tmp2
						var tmp3 bool
						_ = tmp3
						var tmp4 bool
						_ = tmp4
						var tmp5 bool
						_ = tmp5
						var tmp6 bool
						_ = tmp6
						var c HxNullable[int32]
						_ = c
						c = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 8)))
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						if (c).Valid == false {
							_hx_tmp_10 = -1
						} else {
							_hx_tmp_10 = (c).Value
						}

						if (_hx_tmp_10) == 46 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 7)))
							var _hx_tmp_11 int32
							_ = _hx_tmp_11
							if (tmp_c_1).Valid == false {
								_hx_tmp_11 = -1
							} else {
								_hx_tmp_11 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_11
							tmp6 = ((c1 == 108) || (c1 == 76))
						} else {
							tmp6 = false
						}

						if tmp6 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 6)))
							var _hx_tmp_12 int32
							_ = _hx_tmp_12
							if (tmp_c_1).Valid == false {
								_hx_tmp_12 = -1
							} else {
								_hx_tmp_12 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_12
							tmp5 = ((c1 == 111) || (c1 == 79))
						} else {
							tmp5 = false
						}

						if tmp5 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 5)))
							var _hx_tmp_13 int32
							_ = _hx_tmp_13
							if (tmp_c_1).Valid == false {
								_hx_tmp_13 = -1
							} else {
								_hx_tmp_13 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_13
							tmp4 = ((c1 == 114) || (c1 == 82))
						} else {
							tmp4 = false
						}

						if tmp4 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 4)))
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							if (tmp_c_1).Valid == false {
								_hx_tmp_14 = -1
							} else {
								_hx_tmp_14 = (tmp_c_1).Value
							}

							tmp3 = ((_hx_tmp_14) == 46)
						} else {
							tmp3 = false
						}

						if tmp3 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							if (tmp_c_1).Valid == false {
								_hx_tmp_15 = -1
							} else {
								_hx_tmp_15 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_15
							tmp2 = ((c1 == 116) || (c1 == 84))
						} else {
							tmp2 = false
						}

						if tmp2 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
							var _hx_tmp_16 int32
							_ = _hx_tmp_16
							if (tmp_c_1).Valid == false {
								_hx_tmp_16 = -1
							} else {
								_hx_tmp_16 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_16
							tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
						} else {
							tmp_tmp_1 = false
						}

						if tmp_tmp_1 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							if (tmp_c_1).Valid == false {
								_hx_tmp_17 = -1
							} else {
								_hx_tmp_17 = (tmp_c_1).Value
							}

							var c1 int32
							_ = c1
							c1 = _hx_tmp_17
							tmp1 = ((c1 == 116) || (c1 == 84))
						} else {
							tmp1 = false
						}
					}
				}

				if !tmp1 {
					if importPath == HxStringNull {
						tmp = false
					} else {
						var n int32
						_ = n
						n = int32(utf8.RuneCountInString(importPath))
						if n < 4 {
							tmp = false
						} else {
							var tmp_tmp1_1 bool
							_ = tmp_tmp1_1
							var tmp2 bool
							_ = tmp2
							var c HxNullable[int32]
							_ = c
							c = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 4)))
							var _hx_tmp_18 int32
							_ = _hx_tmp_18
							if (c).Valid == false {
								_hx_tmp_18 = -1
							} else {
								_hx_tmp_18 = (c).Value
							}

							if (_hx_tmp_18) == 46 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
								var _hx_tmp_19 int32
								_ = _hx_tmp_19
								if (tmp_c_1).Valid == false {
									_hx_tmp_19 = -1
								} else {
									_hx_tmp_19 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_19
								tmp2 = ((c1 == 108) || (c1 == 76))
							} else {
								tmp2 = false
							}

							if tmp2 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
								var _hx_tmp_20 int32
								_ = _hx_tmp_20
								if (tmp_c_1).Valid == false {
									_hx_tmp_20 = -1
								} else {
									_hx_tmp_20 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_20
								tmp_tmp1_1 = ((c1 == 111) || (c1 == 79))
							} else {
								tmp_tmp1_1 = false
							}

							if tmp_tmp1_1 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
								var _hx_tmp_21 int32
								_ = _hx_tmp_21
								if (tmp_c_1).Valid == false {
									_hx_tmp_21 = -1
								} else {
									_hx_tmp_21 = (tmp_c_1).Value
								}

								var c1 int32
								_ = c1
								c1 = _hx_tmp_21
								tmp = ((c1 == 114) || (c1 == 82))
							} else {
								tmp = false
							}
						}
					}
				} else {
					tmp = true
				}

				if !tmp {
					importPath = Hx_Field_go_haxe_hxdynamic_stringConcat(importPath, ext1)
				}

				if !func() func(_hx_param_key string) bool {
					var _hx_closure_src_484 func(_hx_param_key any) bool
					_ = _hx_closure_src_484
					_hx_closure_src_484 = used.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_484(((any)(_hx_param_key)))
					}
				}()(importPath) {
					func() func(_hx_param_key string, _hx_param_value bool) {
						var _hx_closure_src_485 func(_hx_param_key any, _hx_param_value any)
						_ = _hx_closure_src_485
						_hx_closure_src_485 = used.VTable.Hx_Field_set
						return func(_hx_param_key string, _hx_param_value bool) {
							_hx_closure_src_485(((any)(_hx_param_key)), ((any)(_hx_param_value)))
						}
					}()(importPath, true)
					_ = (Hx_Array_Push[string](result, importPath))
					if importNode.Hx_Field_script != nil {
						this.VTable.Hx_Field__getImportedPaths(importPath, importNode.Hx_Field_script, result, used)
					}
				}
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_resolveArrayAccess(access *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_node {
	var targetNode *Hx_Obj_loreline_node
	_ = targetNode
	if Hx_Field_std_isOfType(access.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
		targetNode = this.VTable.Hx_Field_resolveAccess(access.Hx_Field_target.VTable.(interface {
			Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
		}).Hx_As_Hx_Obj_loreline_naccess())
	} else {
		if Hx_Field_std_isOfType(access.Hx_Field_target, Hx_Obj_loreline_narrayaccess_RTTI) {
			targetNode = this.VTable.Hx_Field_resolveArrayAccess(access.Hx_Field_target.VTable.(interface {
				Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
			}).Hx_As_Hx_Obj_loreline_narrayaccess())
		}
	}

	if targetNode == nil {
		return nil
	}

	{
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_486 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_486
			_hx_closure_src_486 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_486(((any)(o)))))
			}
		}()(targetNode)
		if _g != nil {
			switch _g {
			case Hx_Obj_loreline_nliteral_RTTI:
				{
					var literal *Hx_Obj_loreline_nliteral
					_ = literal
					literal = targetNode.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
					}).Hx_As_Hx_Obj_loreline_nliteral()
					if (literal.Hx_Field_literalType == Hx_Enum_loreline_literaltype_Array{}) {
						var elements HxArray[any]
						_ = elements
						elements = HxAnyToArray(literal.Hx_Field_value)
						if Hx_Field_std_isOfType(access.Hx_Field_index, Hx_Obj_loreline_nliteral_RTTI) {
							var indexLit *Hx_Obj_loreline_nliteral
							_ = indexLit
							indexLit = access.Hx_Field_index.VTable.(interface {
								Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
							}).Hx_As_Hx_Obj_loreline_nliteral()
							if (indexLit.Hx_Field_literalType == Hx_Enum_loreline_literaltype_Number{}) {
								var index int32
								_ = index
								index = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(indexLit.Hx_Field_value)))))
								var _hx_tmp_0 bool
								_ = _hx_tmp_0
								if index >= 0 {
									var _hx_tmp_1 int32
									_ = _hx_tmp_1
									_hx_tmp_1 = index
									var _hx_tmp_2 int32
									_ = _hx_tmp_2
									_hx_tmp_2 = _hx_tmp_1
									_hx_tmp_0 = (_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(elements))
								} else {
									_hx_tmp_0 = false
								}

								if _hx_tmp_0 {
									var element any
									_ = element
									element = elements.Get(index)
									if Hx_Field_std_isOfType(element, Hx_Obj_loreline_node_RTTI) {
										var hx_dyncast_140 any
										_ = hx_dyncast_140
										hx_dyncast_140 = Hx_Field_go_haxe_hxdynamic_castClass(element, "Hx_Obj_loreline_node")
										var _hx_tmp_3 *Hx_Obj_loreline_node
										_ = _hx_tmp_3
										if hx_dyncast_140 != nil {
											_hx_tmp_3 = hx_dyncast_140.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
										} else {
											_hx_tmp_3 = nil
										}

										return _hx_tmp_3
									}
								}
							}
						}
					}
				}

			case Hx_Obj_loreline_nobjectfield_RTTI:
				{
					var field *Hx_Obj_loreline_nobjectfield
					_ = field
					field = targetNode.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
					}).Hx_As_Hx_Obj_loreline_nobjectfield()
					if Hx_Field_std_isOfType(field.Hx_Field_value, Hx_Obj_loreline_nliteral_RTTI) {
						var literal *Hx_Obj_loreline_nliteral
						_ = literal
						literal = field.Hx_Field_value.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
						}).Hx_As_Hx_Obj_loreline_nliteral()
						if (literal.Hx_Field_literalType == Hx_Enum_loreline_literaltype_Array{}) {
							var elements HxArray[any]
							_ = elements
							elements = HxAnyToArray(literal.Hx_Field_value)
							if Hx_Field_std_isOfType(access.Hx_Field_index, Hx_Obj_loreline_nliteral_RTTI) {
								var indexLit *Hx_Obj_loreline_nliteral
								_ = indexLit
								indexLit = access.Hx_Field_index.VTable.(interface {
									Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
								}).Hx_As_Hx_Obj_loreline_nliteral()
								if (indexLit.Hx_Field_literalType == Hx_Enum_loreline_literaltype_Number{}) {
									var index int32
									_ = index
									index = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(indexLit.Hx_Field_value)))))
									var _hx_tmp_4 bool
									_ = _hx_tmp_4
									if index >= 0 {
										var _hx_tmp_5 int32
										_ = _hx_tmp_5
										_hx_tmp_5 = index
										var _hx_tmp_6 int32
										_ = _hx_tmp_6
										_hx_tmp_6 = _hx_tmp_5
										_hx_tmp_4 = (_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(elements))
									} else {
										_hx_tmp_4 = false
									}

									if _hx_tmp_4 {
										var element any
										_ = element
										element = elements.Get(index)
										if Hx_Field_std_isOfType(element, Hx_Obj_loreline_node_RTTI) {
											var hx_dyncast_141 any
											_ = hx_dyncast_141
											hx_dyncast_141 = Hx_Field_go_haxe_hxdynamic_castClass(element, "Hx_Obj_loreline_node")
											var _hx_tmp_7 *Hx_Obj_loreline_node
											_ = _hx_tmp_7
											if hx_dyncast_141 != nil {
												_hx_tmp_7 = hx_dyncast_141.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
											} else {
												_hx_tmp_7 = nil
											}

											return _hx_tmp_7
										}
									}
								}
							}
						}
					}
				}

			default:
				{
				}
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_resolveAccess(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_node {
	if access.Hx_Field_target != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		if Hx_Field_std_isOfType(access.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
			_hx_tmp_0 = this.VTable.Hx_Field_resolveAccess(access.Hx_Field_target.VTable.(interface {
				Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
			}).Hx_As_Hx_Obj_loreline_naccess())
		} else {
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			if Hx_Field_std_isOfType(access.Hx_Field_target, Hx_Obj_loreline_narrayaccess_RTTI) {
				_hx_tmp_1 = this.VTable.Hx_Field_resolveArrayAccess(access.Hx_Field_target.VTable.(interface {
					Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
				}).Hx_As_Hx_Obj_loreline_narrayaccess())
			} else {
				_hx_tmp_1 = nil
			}

			_hx_tmp_0 = _hx_tmp_1
		}

		var targetNode *Hx_Obj_loreline_node
		_ = targetNode
		targetNode = _hx_tmp_0
		if targetNode != nil {
			if Hx_Field_std_isOfType(targetNode, Hx_Obj_loreline_nobjectfield_RTTI) {
				targetNode = (func() *Hx_Obj_loreline_node {
					_hx_up := targetNode.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
					}).Hx_As_Hx_Obj_loreline_nobjectfield().Hx_Field_value
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})()
			}

			{
				var _g *Hx_Obj_go_haxe_hxclass
				_ = _g
				_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					var _hx_closure_src_487 func(o any) *Hx_Obj_go_haxe_hxclass
					_ = _hx_closure_src_487
					_hx_closure_src_487 = Hx_Field_type_getClass
					return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
						return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_487(((any)(o)))))
					}
				}()(targetNode)
				if _g != nil {
					switch _g {
					case Hx_Obj_loreline_ncharacterdecl_RTTI:
						{
							var characterDecl *Hx_Obj_loreline_ncharacterdecl
							_ = characterDecl
							characterDecl = targetNode.VTable.(interface {
								Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
							}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
							{
								var tmp__g_1 int32
								_ = tmp__g_1
								tmp__g_1 = 0
								var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
								_ = _g1
								_g1 = characterDecl.Hx_Field_fields
								for {
									var _hx_tmp_2 int32
									_ = _hx_tmp_2
									_hx_tmp_2 = tmp__g_1
									if !(_hx_tmp_2 < _g1.Len()) {
										break
									}

									var prop *Hx_Obj_loreline_nobjectfield
									_ = prop
									prop = _g1.Get(tmp__g_1)
									tmp__g_1++
									if prop.Hx_Field_name == access.Hx_Field_name {
										return (func() *Hx_Obj_loreline_node {
											_hx_up := prop
											if _hx_up != nil {
												return &_hx_up.Hx_Obj_loreline_node
											}
											return nil
										})()
									}
								}
							}
						}

					case Hx_Obj_loreline_nliteral_RTTI:
						{
							var literal *Hx_Obj_loreline_nliteral
							_ = literal
							literal = targetNode.VTable.(interface {
								Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
							}).Hx_As_Hx_Obj_loreline_nliteral()
							{
								var tmp__g_1 Hx_Enum_loreline_literaltype
								_ = tmp__g_1
								tmp__g_1 = literal.Hx_Field_literalType
								if tmp__g_1.Hx_Field_enumIndex() == 4 {
									var style int32
									_ = style
									style = tmp__g_1.(Hx_Enum_loreline_literaltype_Object).Hx_Field_style
									{
										var fields HxArray[*Hx_Obj_loreline_nobjectfield]
										_ = fields
										fields = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](literal.Hx_Field_value)
										{
											var _g1 int32
											_ = _g1
											_g1 = 0
											for {
												var _hx_tmp_3 int32
												_ = _hx_tmp_3
												_hx_tmp_3 = _g1
												if !(_hx_tmp_3 < fields.Len()) {
													break
												}

												var field *Hx_Obj_loreline_nobjectfield
												_ = field
												field = fields.Get(_g1)
												_g1++
												if field.Hx_Field_name == access.Hx_Field_name {
													return (func() *Hx_Obj_loreline_node {
														_hx_up := field
														if _hx_up != nil {
															return &_hx_up.Hx_Obj_loreline_node
														}
														return nil
													})()
												}
											}
										}
									}
								}
							}
						}

					default:
						{
						}
					}
				}
			}
		}

		return nil
	}

	var name string
	_ = name
	name = access.Hx_Field_name
	if name == HxStringNull {
		return nil
	}

	var currentBeat *Hx_Obj_loreline_nbeatdecl
	_ = currentBeat
	currentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_488 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
		_ = _hx_closure_src_488
		_hx_closure_src_488 = this.VTable.Hx_Field_getFirstParentOfType
		return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_142 any
			_ = hx_dyncast_142
			hx_dyncast_142 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_488(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_4 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_4
			if hx_dyncast_142 != nil {
				_hx_tmp_4 = hx_dyncast_142.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_4 = nil
			}

			return _hx_tmp_4
		}
	}()((func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	for currentBeat != nil {
		var result *Hx_Obj_loreline_node
		_ = result
		var _hx_tmp_5 *Hx_Obj_loreline_node
		_ = _hx_tmp_5
		_hx_tmp_5 = (func() *Hx_Obj_loreline_node {
			_hx_up := currentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_traverse(_hx_tmp_5, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
			if result != nil {
				return false
			}

			if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nstatedecl_RTTI) {
				var stateDecl *Hx_Obj_loreline_nstatedecl
				_ = stateDecl
				stateDecl = node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
				}).Hx_As_Hx_Obj_loreline_nstatedecl()
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
					_ = _g1
					_g1 = stateDecl.Hx_Field_fields
					for {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = _g
						if !(_hx_tmp_6 < _g1.Len()) {
							break
						}

						var field *Hx_Obj_loreline_nobjectfield
						_ = field
						field = _g1.Get(_g)
						_g++
						if field.Hx_Field_name == name {
							result = (func() *Hx_Obj_loreline_node {
								_hx_up := field
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
							return false
						}
					}
				}
			}

			return true
		})
		if result != nil {
			return result
		}

		currentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_489 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_489
			_hx_closure_src_489 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_143 any
				_ = hx_dyncast_143
				hx_dyncast_143 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_489(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_7 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_7
				if hx_dyncast_143 != nil {
					_hx_tmp_7 = hx_dyncast_143.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_7 = nil
				}

				return _hx_tmp_7
			}
		}()((func() *Hx_Obj_loreline_node {
			_hx_up := currentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	}

	var stateField *Hx_Obj_loreline_node
	_ = stateField
	var _hx_tmp_8 *Hx_Obj_loreline_node
	_ = _hx_tmp_8
	_hx_tmp_8 = (func() *Hx_Obj_loreline_node {
		_hx_up := this.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_8, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		if stateField != nil {
			return false
		}

		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nstatedecl_RTTI) {
			var stateDecl *Hx_Obj_loreline_nstatedecl
			_ = stateDecl
			stateDecl = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
			}).Hx_As_Hx_Obj_loreline_nstatedecl()
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
				_ = _g1
				_g1 = stateDecl.Hx_Field_fields
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					if !(_hx_tmp_9 < _g1.Len()) {
						break
					}

					var field *Hx_Obj_loreline_nobjectfield
					_ = field
					field = _g1.Get(_g)
					_g++
					if field.Hx_Field_name == name {
						stateField = (func() *Hx_Obj_loreline_node {
							_hx_up := field
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})()
						return false
					}
				}
			}
		}

		var _hx_tmp_10 bool
		_ = _hx_tmp_10
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nimportstatement_RTTI) {
			_hx_tmp_10 = true
		} else {
			_hx_tmp_10 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_script_RTTI)
		}

		return _hx_tmp_10
	})
	if stateField != nil {
		return stateField
	}

	var _hx_tmp_11 string
	_ = _hx_tmp_11
	_hx_tmp_11 = name
	var characterDecl *Hx_Obj_loreline_ncharacterdecl
	_ = characterDecl
	characterDecl = this.VTable.Hx_Field_findCharacterByNameFromNode(_hx_tmp_11, (func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
	if characterDecl != nil {
		return (func() *Hx_Obj_loreline_node {
			_hx_up := characterDecl
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
	}

	var _hx_tmp_12 string
	_ = _hx_tmp_12
	_hx_tmp_12 = name
	var beatDecl *Hx_Obj_loreline_nbeatdecl
	_ = beatDecl
	beatDecl = this.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_12, (func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
	if beatDecl != nil {
		return (func() *Hx_Obj_loreline_node {
			_hx_up := beatDecl
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
	}

	var _hx_tmp_13 string
	_ = _hx_tmp_13
	_hx_tmp_13 = name
	var functionDecl *Hx_Obj_loreline_nfunctiondecl
	_ = functionDecl
	functionDecl = this.VTable.Hx_Field_findFunctionByNameFromNode(_hx_tmp_13, (func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
	if functionDecl != nil {
		return (func() *Hx_Obj_loreline_node {
			_hx_up := functionDecl
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatFromAccess(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_nbeatdecl {
	if (access.Hx_Field_target == nil) && (access.Hx_Field_name != HxStringNull) {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = access.Hx_Field_name
		return this.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := access
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatParamFromAccess(access *Hx_Obj_loreline_naccess) any {
	if (access.Hx_Field_target != nil) || (access.Hx_Field_name == HxStringNull) {
		return nil
	}

	var parentBeat *Hx_Obj_loreline_nbeatdecl
	_ = parentBeat
	parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_490 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
		_ = _hx_closure_src_490
		_hx_closure_src_490 = this.VTable.Hx_Field_getFirstParentOfType
		return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_144 any
			_ = hx_dyncast_144
			hx_dyncast_144 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_490(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_0
			if hx_dyncast_144 != nil {
				_hx_tmp_0 = hx_dyncast_144.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()((func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	for parentBeat != nil {
		if parentBeat.Hx_Field_params != nil {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_nbeatparam]
			_ = _g1
			_g1 = parentBeat.Hx_Field_params
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < _g1.Len()) {
					break
				}

				var param *Hx_Obj_loreline_nbeatparam
				_ = param
				param = _g1.Get(_g)
				_g++
				if param.Hx_Field_name == access.Hx_Field_name {
					return any(map[string]any{"param": ((any)(param)), "beat": ((any)(parentBeat))})
				}
			}
		}

		parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_491 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_491
			_hx_closure_src_491 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_145 any
				_ = hx_dyncast_145
				hx_dyncast_145 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_491(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_2
				if hx_dyncast_145 != nil {
					_hx_tmp_2 = hx_dyncast_145.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_2 = nil
				}

				return _hx_tmp_2
			}
		}()((func() *Hx_Obj_loreline_node {
			_hx_up := parentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findCallArgument(node *Hx_Obj_loreline_node) any {
	var child *Hx_Obj_loreline_node
	_ = child
	child = node
	var parent *Hx_Obj_loreline_node
	_ = parent
	parent = this.VTable.Hx_Field_getParentNode(child)
	for parent != nil {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_492 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_492
				_hx_closure_src_492 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_492(((any)(o)))))
				}
			}()(parent)
			if _g == nil {
				if !Hx_Field_std_isOfType(parent, Hx_Obj_loreline_nexpr_RTTI) {
					return nil
				}
			} else {
				switch _g {
				case Hx_Obj_loreline_nbeatcall_RTTI:
					{
						var beatCall *Hx_Obj_loreline_nbeatcall
						_ = beatCall
						beatCall = parent.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
						}).Hx_As_Hx_Obj_loreline_nbeatcall()
						if beatCall.Hx_Field_args != nil {
							var _hx_tmp_0 HxArray[*Hx_Obj_loreline_nexpr]
							_ = _hx_tmp_0
							_hx_tmp_0 = beatCall.Hx_Field_args
							var _hx_tmp_1 *Hx_Obj_loreline_nexpr
							_ = _hx_tmp_1
							_hx_tmp_1 = child.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
							var index int32
							_ = index
							index = (Hx_Array_IndexOf[*Hx_Obj_loreline_nexpr](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
							if index != -1 {
								return any(map[string]any{"owner": ((any)(parent)), "index": ((any)(index))})
							}
						}

						return nil
					}

				case Hx_Obj_loreline_ncall_RTTI:
					{
						var call *Hx_Obj_loreline_ncall
						_ = call
						call = parent.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()
						if call.Hx_Field_args != nil {
							var _hx_tmp_2 HxArray[*Hx_Obj_loreline_nexpr]
							_ = _hx_tmp_2
							_hx_tmp_2 = call.Hx_Field_args
							var _hx_tmp_3 *Hx_Obj_loreline_nexpr
							_ = _hx_tmp_3
							_hx_tmp_3 = child.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
							var index int32
							_ = index
							index = (Hx_Array_IndexOf[*Hx_Obj_loreline_nexpr](_hx_tmp_2, _hx_tmp_3, HxNullable[int32]{}))
							if index != -1 {
								return any(map[string]any{"owner": ((any)(parent)), "index": ((any)(index))})
							}
						}

						return nil
					}

				case Hx_Obj_loreline_ninsertion_RTTI:
					{
						var insertion *Hx_Obj_loreline_ninsertion
						_ = insertion
						insertion = parent.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
						}).Hx_As_Hx_Obj_loreline_ninsertion()
						if insertion.Hx_Field_args != nil {
							var _hx_tmp_4 HxArray[*Hx_Obj_loreline_nexpr]
							_ = _hx_tmp_4
							_hx_tmp_4 = insertion.Hx_Field_args
							var _hx_tmp_5 *Hx_Obj_loreline_nexpr
							_ = _hx_tmp_5
							_hx_tmp_5 = child.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
							var index int32
							_ = index
							index = (Hx_Array_IndexOf[*Hx_Obj_loreline_nexpr](_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{}))
							if index != -1 {
								return any(map[string]any{"owner": ((any)(parent)), "index": ((any)(index))})
							}
						}

						return nil
					}

				case Hx_Obj_loreline_ntransition_RTTI:
					{
						var transition *Hx_Obj_loreline_ntransition
						_ = transition
						transition = parent.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
						}).Hx_As_Hx_Obj_loreline_ntransition()
						if transition.Hx_Field_args != nil {
							var _hx_tmp_6 HxArray[*Hx_Obj_loreline_nexpr]
							_ = _hx_tmp_6
							_hx_tmp_6 = transition.Hx_Field_args
							var _hx_tmp_7 *Hx_Obj_loreline_nexpr
							_ = _hx_tmp_7
							_hx_tmp_7 = child.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
							var index int32
							_ = index
							index = (Hx_Array_IndexOf[*Hx_Obj_loreline_nexpr](_hx_tmp_6, _hx_tmp_7, HxNullable[int32]{}))
							if index != -1 {
								return any(map[string]any{"owner": ((any)(parent)), "index": ((any)(index))})
							}
						}

						return nil
					}

				default:
					{
						if !Hx_Field_std_isOfType(parent, Hx_Obj_loreline_nexpr_RTTI) {
							return nil
						}
					}
				}
			}
		}

		child = parent
		parent = this.VTable.Hx_Field_getParentNode(child)
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatFromTransition(transition *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_nbeatdecl {
	if transition.Hx_Field_targetExpr != nil {
		return nil
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = transition.Hx_Field_target
	return this.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
		_hx_up := transition
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatFromInsertion(insertion *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_nbeatdecl {
	if insertion.Hx_Field_targetExpr != nil {
		return nil
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = insertion.Hx_Field_target
	return this.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
		_hx_up := insertion
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatByPathFromNode(path string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl {
	var parts HxArray[string]
	_ = parts
	parts = (Hx_Field_go_haxe_hxstring_split(path, "."))
	var beat *Hx_Obj_loreline_nbeatdecl
	_ = beat
	beat = this.VTable.Hx_Field_findBeatByNameFromNode(parts.Get(0), node)
	var i int32
	_ = i
	i = 1
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if beat != nil {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = i
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _hx_tmp_1
			_hx_tmp_0 = (_hx_tmp_2 < parts.Len())
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		var name string
		_ = name
		name = parts.Get(i)
		var foundBeat *Hx_Obj_loreline_nbeatdecl
		_ = foundBeat
		var _hx_tmp_3 *Hx_Obj_loreline_node
		_ = _hx_tmp_3
		_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_traverse(_hx_tmp_3, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
			if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nbeatdecl_RTTI) {
				var tmp_beat_1 *Hx_Obj_loreline_nbeatdecl
				_ = tmp_beat_1
				tmp_beat_1 = node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				if tmp_beat_1.Hx_Field_name == name {
					foundBeat = tmp_beat_1
				}

				return false
			} else {
				return (foundBeat == nil)
			}
		})
		beat = foundBeat
		i++
	}

	return beat
}

func (this *Hx_Obj_loreline_lens) Hx_Field_isTopLevelNode(node *Hx_Obj_loreline_node) bool {
	if Hx_Field_std_isOfType(node, Hx_Obj_loreline_astnode_RTTI) {
		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		astNode = node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
		}).Hx_As_Hx_Obj_loreline_astnode()
		var _hx_tmp_0 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_script.Hx_Field_body
		var _hx_tmp_1 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_1
		_hx_tmp_1 = astNode
		return ((Hx_Array_IndexOf[*Hx_Obj_loreline_astnode](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) != -1)
	}

	return false
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findTopLevelBeatFromNode(node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl {
	if this.VTable.Hx_Field_isTopLevelNode(node) {
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nbeatdecl_RTTI) {
			return node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
			}).Hx_As_Hx_Obj_loreline_nbeatdecl()
		} else {
			return nil
		}
	} else {
		var resolved *Hx_Obj_loreline_nbeatdecl
		_ = resolved
		resolved = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_493 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_493
			_hx_closure_src_493 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_146 any
				_ = hx_dyncast_146
				hx_dyncast_146 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_493(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_0
				if hx_dyncast_146 != nil {
					_hx_tmp_0 = hx_dyncast_146.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(node, ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
		for {
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if resolved != nil {
				_hx_tmp_1 = !this.VTable.Hx_Field_isTopLevelNode((func() *Hx_Obj_loreline_node {
					_hx_up := resolved
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})())
			} else {
				_hx_tmp_1 = false
			}

			if !(_hx_tmp_1) {
				break
			}

			resolved = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var _hx_closure_src_494 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
				_ = _hx_closure_src_494
				_hx_closure_src_494 = this.VTable.Hx_Field_getFirstParentOfType
				return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
					var hx_dyncast_147 any
					_ = hx_dyncast_147
					hx_dyncast_147 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_494(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
					var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
					_ = _hx_tmp_2
					if hx_dyncast_147 != nil {
						_hx_tmp_2 = hx_dyncast_147.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					} else {
						_hx_tmp_2 = nil
					}

					return _hx_tmp_2
				}
			}()((func() *Hx_Obj_loreline_node {
				_hx_up := resolved
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if resolved != nil {
			_hx_tmp_3 = Hx_Field_std_isOfType(resolved, Hx_Obj_loreline_nbeatdecl_RTTI)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			return resolved
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nbeatdecl {
	var result *Hx_Obj_loreline_nbeatdecl
	_ = result
	var parentBeat *Hx_Obj_loreline_nbeatdecl
	_ = parentBeat
	parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_495 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
		_ = _hx_closure_src_495
		_hx_closure_src_495 = this.VTable.Hx_Field_getFirstParentOfType
		return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_148 any
			_ = hx_dyncast_148
			hx_dyncast_148 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_495(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_0
			if hx_dyncast_148 != nil {
				_hx_tmp_0 = hx_dyncast_148.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(node, ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	for parentBeat != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := parentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_traverse(_hx_tmp_1, func(child *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
			if (result != nil) || (child == node) {
				return false
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nbeatdecl_RTTI) {
					var beatDecl *Hx_Obj_loreline_nbeatdecl
					_ = beatDecl
					beatDecl = child.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					if beatDecl.Hx_Field_name == name {
						result = beatDecl
					}

					return false
				}
			}

			return true
		})
		parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_496 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_496
			_hx_closure_src_496 = this.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_149 any
				_ = hx_dyncast_149
				hx_dyncast_149 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_496(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_2
				if hx_dyncast_149 != nil {
					_hx_tmp_2 = hx_dyncast_149.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_2 = nil
				}

				return _hx_tmp_2
			}
		}()((func() *Hx_Obj_loreline_node {
			_hx_up := parentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	}

	if result == nil {
		var _hx_tmp_3 *Hx_Obj_loreline_node
		_ = _hx_tmp_3
		_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_traverse(_hx_tmp_3, func(child *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if result == nil {
				_hx_tmp_4 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_nbeatdecl_RTTI)
			} else {
				_hx_tmp_4 = false
			}

			if _hx_tmp_4 {
				var beatDecl *Hx_Obj_loreline_nbeatdecl
				_ = beatDecl
				beatDecl = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				if beatDecl.Hx_Field_name == name {
					result = beatDecl
				}
			}

			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nimportstatement_RTTI) {
				_hx_tmp_5 = true
			} else {
				_hx_tmp_5 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_script_RTTI)
			}

			return _hx_tmp_5
		})
	}

	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findCharacterFromDialogue(dialogue *Hx_Obj_loreline_ndialoguestatement) *Hx_Obj_loreline_ncharacterdecl {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = dialogue.Hx_Field_character
	return this.VTable.Hx_Field_findCharacterByNameFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
		_hx_up := dialogue
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findCharacterByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_ncharacterdecl {
	var result *Hx_Obj_loreline_ncharacterdecl
	_ = result
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := this.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(child *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if result == nil {
			_hx_tmp_1 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_ncharacterdecl_RTTI)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var characterDecl *Hx_Obj_loreline_ncharacterdecl
			_ = characterDecl
			characterDecl = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
			}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
			if characterDecl.Hx_Field_name == name {
				result = characterDecl
			}
		}

		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nimportstatement_RTTI) {
			_hx_tmp_2 = true
		} else {
			_hx_tmp_2 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_script_RTTI)
		}

		return _hx_tmp_2
	})
	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getWhenWarnings(when *Hx_Obj_loreline_nwhenstatement) HxArray[any] {
	var warnings HxArray[any]
	_ = warnings
	warnings = HxMakeArray[any]().Dyn()
	var strategy HxNullable[string]
	_ = strategy
	strategy = when.Hx_Field_strategy
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if (strategy).Valid != false {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if (strategy).Valid == false {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = ((strategy).Value != "first")
		}

		_hx_tmp_2 = (_hx_tmp_3)
	} else {
		_hx_tmp_2 = false
	}

	if _hx_tmp_2 {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if (strategy).Valid == false {
			_hx_tmp_4 = true
		} else {
			_hx_tmp_4 = ((strategy).Value != "pick")
		}

		_hx_tmp_1 = (_hx_tmp_4)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var _hx_tmp_5 string
		_ = _hx_tmp_5
		_hx_tmp_5 = (strategy).Value
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = _hx_tmp_5
		_hx_tmp_0 = (this.VTable.Hx_Field_findFunctionByNameFromNode(_hx_tmp_6, (func() *Hx_Obj_loreline_node {
			_hx_up := when
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()) == nil)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var pos *Hx_Obj_loreline_position
		_ = pos
		pos = when.Hx_Field_strategyPos
		var _hx_tmp_7 HxArray[any]
		_ = _hx_tmp_7
		_hx_tmp_7 = HxMakeArrayView[any](warnings)
		var _hx_tmp_8 *Hx_Obj_loreline_position
		_ = _hx_tmp_8
		if pos != nil {
			_hx_tmp_8 = pos
		} else {
			_hx_tmp_8 = when.Hx_Field_pos
		}

		var _hx_tmp_9 string
		_ = _hx_tmp_9
		_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Unknown strategy: ", (strategy).Value), ". Use first, pick, or the name of a function. A function provided by the host can be declared with `function ")
		_ = (Hx_Array_Push[any](_hx_tmp_7, any(map[string]any{"pos": ((any)(_hx_tmp_8)), "message": ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_9, (strategy).Value), "(rules)`."))), "isWarning": ((any)(true))})))
	}

	var comparedValues Hx_Typedef_map
	_ = comparedValues
	comparedValues = (func() *Hx_Obj_haxe_imap {
		_hx_icast_150 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_150 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_150, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_150.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var thresholdWarned bool
	_ = thresholdWarned
	thresholdWarned = false
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nwhenrule]
		_ = _g1
		_g1 = when.Hx_Field_rules
		for {
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = _g
			if !(_hx_tmp_10 < _g1.Len()) {
				break
			}

			var rule *Hx_Obj_loreline_nwhenrule
			_ = rule
			rule = _g1.Get(_g)
			_g++
			if (rule.Hx_Field_insertion != nil) || (rule.Hx_Field_condition == nil) {
				continue
			}

			var first *Hx_Obj_loreline_nexpr
			_ = first
			first = Hx_Field_loreline_lens_leftmostOperand(rule.Hx_Field_condition)
			var _hx_tmp_11 bool
			_ = _hx_tmp_11
			var _hx_tmp_12 bool
			_ = _hx_tmp_12
			if !rule.Hx_Field_once {
				_hx_tmp_12 = Hx_Field_std_isOfType(first, Hx_Obj_loreline_nunary_RTTI)
			} else {
				_hx_tmp_12 = false
			}

			if _hx_tmp_12 {
				_hx_tmp_11 = (first.Hx_Field_pos.Hx_Field_offset == rule.Hx_Field_pos.Hx_Field_offset)
			} else {
				_hx_tmp_11 = false
			}

			if _hx_tmp_11 {
				if first.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
				}).Hx_As_Hx_Obj_loreline_nunary().Hx_Field_op.Hx_Field_enumIndex() == 28 {
					var _hx_tmp_13 HxArray[any]
					_ = _hx_tmp_13
					_hx_tmp_13 = HxMakeArrayView[any](warnings)
					_ = (Hx_Array_Push[any](_hx_tmp_13, any(map[string]any{"pos": ((any)(Hx_Obj_loreline_position_CreateInstance(first.Hx_Field_pos.Hx_Field_line, first.Hx_Field_pos.Hx_Field_column, first.Hx_Field_pos.Hx_Field_offset, 1))), "message": ((any)("This `-` is a minus sign on the condition. To play this rule only once, write `- ` with a space after it.")), "isWarning": ((any)(true))})))
				}
			}

			var _hx_tmp_14 bool
			_ = _hx_tmp_14
			if (strategy).Valid == false {
				_hx_tmp_14 = !thresholdWarned
			} else {
				_hx_tmp_14 = false
			}

			if _hx_tmp_14 {
				var score int32
				_ = score
				score = Hx_Field_loreline_astutils_whenRuleScore(rule)
				{
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var tmp__g1_1 HxArray[*Hx_Obj_loreline_nexpr]
					_ = tmp__g1_1
					tmp__g1_1 = Hx_Field_loreline_lens_topLevelClauses(rule.Hx_Field_condition)
					for {
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = tmp__g_1
						if !(_hx_tmp_15 < tmp__g1_1.Len()) {
							break
						}

						var clause *Hx_Obj_loreline_nexpr
						_ = clause
						clause = tmp__g1_1.Get(tmp__g_1)
						tmp__g_1++
						var value HxNullable[string]
						_ = value
						value = Hx_Field_loreline_lens_comparedValue(clause)
						if (value).Valid == false {
							continue
						}

						var _hx_tmp_16 string
						_ = _hx_tmp_16
						_hx_tmp_16 = (Hx_Field_go_haxe_hxdynamic_toString(score) + ":")
						var key string
						_ = key
						key = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_16, (value).Value)
						if func() func(_hx_param_key string) bool {
							var _hx_closure_src_497 func(_hx_param_key any) bool
							_ = _hx_closure_src_497
							_hx_closure_src_497 = comparedValues.VTable.Hx_Field_exists
							return func(_hx_param_key string) bool {
								return _hx_closure_src_497(((any)(_hx_param_key)))
							}
						}()(key) {
							var _hx_tmp_17 HxArray[any]
							_ = _hx_tmp_17
							_hx_tmp_17 = HxMakeArrayView[any](warnings)
							var _hx_tmp_18 any
							_ = _hx_tmp_18
							_hx_tmp_18 = ((any)(clause.Hx_Field_pos))
							_ = (Hx_Array_Push[any](_hx_tmp_17, any(map[string]any{"pos": _hx_tmp_18, "message": ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Rules comparing `", (value).Value), "` with a number, with as many criteria, tie with the default strategy: the one played least recently wins, not the first that matches. Use `when first` for thresholds."))), "isWarning": ((any)(true))})))
							thresholdWarned = true
							break
						}

						func() func(_hx_param_key string, _hx_param_value bool) {
							var _hx_closure_src_498 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_498
							_hx_closure_src_498 = comparedValues.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value bool) {
								_hx_closure_src_498(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(key, true)
					}
				}
			}
		}
	}

	return warnings
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findFunctionByNameFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_nfunctiondecl {
	var result *Hx_Obj_loreline_nfunctiondecl
	_ = result
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := this.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(child *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if result == nil {
			_hx_tmp_1 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_nfunctiondecl_RTTI)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var functionDecl *Hx_Obj_loreline_nfunctiondecl
			_ = functionDecl
			functionDecl = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
			}).Hx_As_Hx_Obj_loreline_nfunctiondecl()
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if (functionDecl.Hx_Field_name).Valid == true {
				_hx_tmp_2 = ((functionDecl.Hx_Field_name).Value == name)
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				result = functionDecl
			}
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nimportstatement_RTTI) {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = Hx_Field_std_isOfType(child, Hx_Obj_loreline_script_RTTI)
		}

		return _hx_tmp_3
	})
	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getVisibleCharacters() HxArray[*Hx_Obj_loreline_ncharacterdecl] {
	var result HxArray[*Hx_Obj_loreline_ncharacterdecl]
	_ = result
	result = HxMakeArray[*Hx_Obj_loreline_ncharacterdecl]()
	{
		var node any
		_ = node
		node = ((any)(this.Hx_Field_script.VTable.Hx_Field_iterator()))
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(node, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(node, "next")
			var hx_dyncast_151 any
			_ = hx_dyncast_151
			hx_dyncast_151 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_astnode")
			var _hx_tmp_2 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_2
			if hx_dyncast_151 != nil {
				_hx_tmp_2 = hx_dyncast_151.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
			} else {
				_hx_tmp_2 = nil
			}

			var node1 *Hx_Obj_loreline_astnode
			_ = node1
			node1 = _hx_tmp_2
			if Hx_Field_std_isOfType(node1, Hx_Obj_loreline_ncharacterdecl_RTTI) {
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_ncharacterdecl]
				_ = _hx_tmp_3
				_hx_tmp_3 = result
				_ = (Hx_Array_Push[*Hx_Obj_loreline_ncharacterdecl](_hx_tmp_3, node1.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
				}).Hx_As_Hx_Obj_loreline_ncharacterdecl()))
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getVisibleFunctions() HxArray[*Hx_Obj_loreline_nfunctiondecl] {
	var result HxArray[*Hx_Obj_loreline_nfunctiondecl]
	_ = result
	result = HxMakeArray[*Hx_Obj_loreline_nfunctiondecl]()
	{
		var node any
		_ = node
		node = ((any)(this.Hx_Field_script.VTable.Hx_Field_iterator()))
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(node, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(node, "next")
			var hx_dyncast_152 any
			_ = hx_dyncast_152
			hx_dyncast_152 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_astnode")
			var _hx_tmp_2 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_2
			if hx_dyncast_152 != nil {
				_hx_tmp_2 = hx_dyncast_152.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
			} else {
				_hx_tmp_2 = nil
			}

			var node1 *Hx_Obj_loreline_astnode
			_ = node1
			node1 = _hx_tmp_2
			if Hx_Field_std_isOfType(node1, Hx_Obj_loreline_nfunctiondecl_RTTI) {
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_nfunctiondecl]
				_ = _hx_tmp_3
				_hx_tmp_3 = result
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nfunctiondecl](_hx_tmp_3, node1.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
				}).Hx_As_Hx_Obj_loreline_nfunctiondecl()))
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getVisibleStateFields(fromNode *Hx_Obj_loreline_node) HxArray[*Hx_Obj_loreline_nobjectfield] {
	var fields HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = fields
	fields = HxMakeArray[*Hx_Obj_loreline_nobjectfield]()
	var seenFields_map *Hx_Obj_loreline_int64map
	_ = seenFields_map
	seenFields_map = Hx_Obj_loreline_int64map_CreateInstance()
	var current *Hx_Obj_loreline_node
	_ = current
	current = fromNode
	for current != nil {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_499 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_499
				_hx_closure_src_499 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_499(((any)(o)))))
				}
			}()(current)
			if _g != nil {
				if _g == Hx_Obj_loreline_nstatedecl_RTTI {
					var state *Hx_Obj_loreline_nstatedecl
					_ = state
					state = current.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
					}).Hx_As_Hx_Obj_loreline_nstatedecl()
					{
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
						_ = _g1
						_g1 = state.Hx_Field_fields
						for {
							var _hx_tmp_0 int32
							_ = _hx_tmp_0
							_hx_tmp_0 = tmp__g_1
							if !(_hx_tmp_0 < _g1.Len()) {
								break
							}

							var field *Hx_Obj_loreline_nobjectfield
							_ = field
							field = _g1.Get(tmp__g_1)
							tmp__g_1++
							var key Hx_Typedef_haxe__int64___int64
							_ = key
							key = field.Hx_Field_id
							var _hx_tmp_1 int64
							_ = _hx_tmp_1
							_hx_tmp_1 = ((int64)(key))
							if !seenFields_map.VTable.Hx_Field__exists(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key))) {
								{
									var tmp_key_1 Hx_Typedef_haxe__int64___int64
									_ = tmp_key_1
									tmp_key_1 = field.Hx_Field_id
									var _hx_tmp_2 int64
									_ = _hx_tmp_2
									_hx_tmp_2 = ((int64)(tmp_key_1))
									func() func(high int32, low int32, value bool) {
										var _hx_closure_src_500 func(high int32, low int32, value any)
										_ = _hx_closure_src_500
										_hx_closure_src_500 = seenFields_map.VTable.Hx_Field__set
										return func(high int32, low int32, value bool) {
											_hx_closure_src_500(high, low, ((any)(value)))
										}
									}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), true)
								}

								_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](fields, field))
							}
						}
					}
				}
			}
		}

		var key Hx_Typedef_haxe__int64___int64
		_ = key
		key = current.Hx_Field_id
		var _hx_tmp_4 int64
		_ = _hx_tmp_4
		_hx_tmp_4 = ((int64)(key))
		current = func() func(high int32, low int32) *Hx_Obj_loreline_node {
			var _hx_closure_src_501 func(high int32, low int32) any
			_ = _hx_closure_src_501
			_hx_closure_src_501 = this.Hx_Field_parentNodes.Hx_Field_map.VTable.Hx_Field__get
			return func(high int32, low int32) *Hx_Obj_loreline_node {
				var hx_dyncast_153 any
				_ = hx_dyncast_153
				hx_dyncast_153 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_501(high, low), "Hx_Obj_loreline_node")
				var _hx_tmp_3 *Hx_Obj_loreline_node
				_ = _hx_tmp_3
				if hx_dyncast_153 != nil {
					_hx_tmp_3 = hx_dyncast_153.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
				} else {
					_hx_tmp_3 = nil
				}

				return _hx_tmp_3
			}
		}()(((int32)((_hx_tmp_4 >> int64(((int32)(32)))))), ((int32)(key)))
	}

	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_502 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_502
			_hx_closure_src_502 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_502(((any)(o)))))
			}
		}()(node)
		if _g != nil {
			if _g == Hx_Obj_loreline_nstatedecl_RTTI {
				var state *Hx_Obj_loreline_nstatedecl
				_ = state
				state = node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
				}).Hx_As_Hx_Obj_loreline_nstatedecl()
				if Hx_Field_std_isOfType(parent, Hx_Obj_loreline_script_RTTI) {
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
					_ = _g1
					_g1 = state.Hx_Field_fields
					for {
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = tmp__g_1
						if !(_hx_tmp_5 < _g1.Len()) {
							break
						}

						var field *Hx_Obj_loreline_nobjectfield
						_ = field
						field = _g1.Get(tmp__g_1)
						tmp__g_1++
						var key Hx_Typedef_haxe__int64___int64
						_ = key
						key = field.Hx_Field_id
						var _hx_tmp_6 int64
						_ = _hx_tmp_6
						_hx_tmp_6 = ((int64)(key))
						if !seenFields_map.VTable.Hx_Field__exists(((int32)((_hx_tmp_6 >> int64(((int32)(32)))))), ((int32)(key))) {
							{
								var tmp_key_1 Hx_Typedef_haxe__int64___int64
								_ = tmp_key_1
								tmp_key_1 = field.Hx_Field_id
								var _hx_tmp_7 int64
								_ = _hx_tmp_7
								_hx_tmp_7 = ((int64)(tmp_key_1))
								func() func(high int32, low int32, value bool) {
									var _hx_closure_src_503 func(high int32, low int32, value any)
									_ = _hx_closure_src_503
									_hx_closure_src_503 = seenFields_map.VTable.Hx_Field__set
									return func(high int32, low int32, value bool) {
										_hx_closure_src_503(high, low, ((any)(value)))
									}
								}()(((int32)((_hx_tmp_7 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), true)
							}

							_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](fields, field))
						}
					}
				}
			}
		}
	})
	return fields
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getVisibleBeats(fromNode *Hx_Obj_loreline_node) HxArray[*Hx_Obj_loreline_nbeatdecl] {
	var beats HxArray[*Hx_Obj_loreline_nbeatdecl]
	_ = beats
	beats = HxMakeArray[*Hx_Obj_loreline_nbeatdecl]()
	var seenBeats *Hx_Obj_haxe_imap
	_ = seenBeats
	seenBeats = (func() *Hx_Obj_haxe_imap {
		_hx_icast_154 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_154 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_154, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_154.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var current *Hx_Obj_loreline_node
	_ = current
	current = this.VTable.Hx_Field_getParentNode(fromNode)
	for current != nil {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_504 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_504
				_hx_closure_src_504 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_504(((any)(o)))))
				}
			}()(current)
			if _g != nil {
				if _g == Hx_Obj_loreline_nbeatdecl_RTTI {
					var parent *Hx_Obj_loreline_node
					_ = parent
					parent = this.VTable.Hx_Field_getParentNode(current)
					if parent != nil {
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						var _hx_tmp_2 *Hx_Obj_go_haxe_hxclass
						_ = _hx_tmp_2
						_hx_tmp_2 = ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI))
						var hx_nullable_f202 HxNullable[bool]
						_ = hx_nullable_f202
						hx_nullable_f202 = HxNullable[bool]{}

						var _hx_tmp_3 any
						_ = _hx_tmp_3
						if (hx_nullable_f202).Valid == true {
							_hx_tmp_3 = (hx_nullable_f202).Value
						} else {
							_hx_tmp_3 = nil
						}

						var _g1 HxArray[*Hx_Obj_loreline_nbeatdecl]
						_ = _g1
						_g1 = func() func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nbeatdecl] {
							var _hx_closure_src_505 func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any]
							_ = _hx_closure_src_505
							_hx_closure_src_505 = this.VTable.Hx_Field_getNodesOfType
							return func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nbeatdecl] {
								includeImported := HxDefault[bool]()
								_ = includeImported

								if includeImported__isOptional != nil {
									includeImported = HxConvert[bool](includeImported__isOptional)
								} else {
								}
								{
									var _hx_tmp_0 *Hx_Obj_go_haxe_hxclass
									_ = _hx_tmp_0
									_hx_tmp_0 = nodeType
									var _hx_tmp_1 any
									_ = _hx_tmp_1
									if includeImported__isOptional == nil {
										_hx_tmp_1 = nil
									} else {
										_hx_tmp_1 = includeImported
									}

									return HxMakeArrayView[*Hx_Obj_loreline_nbeatdecl](_hx_closure_src_505(_hx_tmp_0, _hx_tmp_1))
								}
							}
						}()(_hx_tmp_2, _hx_tmp_3)
						for {
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = tmp__g_1
							if !(_hx_tmp_4 < _g1.Len()) {
								break
							}

							var child *Hx_Obj_loreline_nbeatdecl
							_ = child
							child = _g1.Get(tmp__g_1)
							tmp__g_1++
							if !func() func(_hx_param_key string) bool {
								var _hx_closure_src_506 func(_hx_param_key any) bool
								_ = _hx_closure_src_506
								_hx_closure_src_506 = seenBeats.VTable.Hx_Field_exists
								return func(_hx_param_key string) bool {
									return _hx_closure_src_506(((any)(_hx_param_key)))
								}
							}()(child.Hx_Field_name) {
								func() func(_hx_param_key string, _hx_param_value bool) {
									var _hx_closure_src_507 func(_hx_param_key any, _hx_param_value any)
									_ = _hx_closure_src_507
									_hx_closure_src_507 = seenBeats.VTable.Hx_Field_set
									return func(_hx_param_key string, _hx_param_value bool) {
										_hx_closure_src_507(((any)(_hx_param_key)), ((any)(_hx_param_value)))
									}
								}()(child.Hx_Field_name, true)
								_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatdecl](beats, child))
							}
						}
					}
				}
			}
		}

		current = this.VTable.Hx_Field_getParentNode(current)
	}

	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_508 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_508
			_hx_closure_src_508 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_508(((any)(o)))))
			}
		}()(node)
		if _g != nil {
			if _g == Hx_Obj_loreline_nbeatdecl_RTTI {
				if Hx_Field_std_isOfType(parent, Hx_Obj_loreline_script_RTTI) {
					var beat *Hx_Obj_loreline_nbeatdecl
					_ = beat
					beat = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					if !func() func(_hx_param_key string) bool {
						var _hx_closure_src_509 func(_hx_param_key any) bool
						_ = _hx_closure_src_509
						_hx_closure_src_509 = seenBeats.VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_509(((any)(_hx_param_key)))
						}
					}()(beat.Hx_Field_name) {
						func() func(_hx_param_key string, _hx_param_value bool) {
							var _hx_closure_src_510 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_510
							_hx_closure_src_510 = seenBeats.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value bool) {
								_hx_closure_src_510(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(beat.Hx_Field_name, true)
						_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatdecl](beats, beat))
					}
				}
			}
		}
	})
	return beats
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getAllTags() HxArray[string] {
	var tags *Hx_Obj_haxe_imap
	_ = tags
	tags = (func() *Hx_Obj_haxe_imap {
		_hx_icast_155 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_155 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_155, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_155.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var _hx_tmp_0 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_0
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_astnode_RTTI) {
			_hx_tmp_0 = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
		} else {
			_hx_tmp_0 = nil
		}

		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		astNode = _hx_tmp_0
		if (astNode != nil) && (astNode.Hx_Field_trailingComments != nil) {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_comment]
			_ = _g1
			_g1 = astNode.Hx_Field_trailingComments
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < _g1.Len()) {
					break
				}

				var comment *Hx_Obj_loreline_comment
				_ = comment
				comment = _g1.Get(_g)
				_g++
				if comment.Hx_Field_isHash {
					var key string
					_ = key
					key = Hx_Field_stringtools_trim(comment.Hx_Field_content)
					func() func(_hx_param_key string, _hx_param_value bool) {
						var _hx_closure_src_511 func(_hx_param_key any, _hx_param_value any)
						_ = _hx_closure_src_511
						_hx_closure_src_511 = tags.VTable.Hx_Field_set
						return func(_hx_param_key string, _hx_param_value bool) {
							_hx_closure_src_511(((any)(_hx_param_key)), ((any)(_hx_param_value)))
						}
					}()(key, true)
				}
			}
		}
	})
	var _g HxArray[string]
	_ = _g
	_g = HxMakeArray[string]()
	{
		var tag any
		_ = tag
		tag = tags.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(tag, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(tag, "next")
			var tag1 string
			_ = tag1
			tag1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))
			_ = (Hx_Array_Push[string](_g, tag1))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_lens) Hx_Field_countTags() Hx_Typedef_map {
	var tags *Hx_Obj_haxe_imap
	_ = tags
	tags = (func() *Hx_Obj_haxe_imap {
		_hx_icast_156 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_156 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_156, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_156.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var _hx_tmp_0 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_0
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_astnode_RTTI) {
			_hx_tmp_0 = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
		} else {
			_hx_tmp_0 = nil
		}

		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		astNode = _hx_tmp_0
		if (astNode != nil) && (astNode.Hx_Field_trailingComments != nil) {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_comment]
			_ = _g1
			_g1 = astNode.Hx_Field_trailingComments
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < _g1.Len()) {
					break
				}

				var comment *Hx_Obj_loreline_comment
				_ = comment
				comment = _g1.Get(_g)
				_g++
				if comment.Hx_Field_isHash {
					var text string
					_ = text
					text = Hx_Field_stringtools_trim(comment.Hx_Field_content)
					var prevCount HxNullable[int32]
					_ = prevCount
					prevCount = func() func(_hx_param_key string) HxNullable[int32] {
						var _hx_closure_src_512 func(_hx_param_key any) any
						_ = _hx_closure_src_512
						_hx_closure_src_512 = tags.VTable.Hx_Field_get
						return func(_hx_param_key string) HxNullable[int32] {
							var hx_nullable_t56 any
							_ = hx_nullable_t56
							hx_nullable_t56 = _hx_closure_src_512(((any)(_hx_param_key)))
							var _hx_tmp_2 HxNullable[int32]
							_ = _hx_tmp_2
							if hx_nullable_t56 == nil {
								_hx_tmp_2 = HxNullable[int32]{}
							} else {
								_hx_tmp_2 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t56))), Valid: true}
							}

							return _hx_tmp_2
						}
					}()(text)
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					if (prevCount).Valid != false {
						_hx_tmp_3 = (prevCount).Value
					} else {
						_hx_tmp_3 = 0
					}

					var prevCount1 int32
					_ = prevCount1
					prevCount1 = _hx_tmp_3
					func() func(_hx_param_key string, _hx_param_value int32) {
						var _hx_closure_src_513 func(_hx_param_key any, _hx_param_value any)
						_ = _hx_closure_src_513
						_hx_closure_src_513 = tags.VTable.Hx_Field_set
						return func(_hx_param_key string, _hx_param_value int32) {
							_hx_closure_src_513(((any)(_hx_param_key)), ((any)(_hx_param_value)))
						}
					}()(text, (prevCount1 + 1))
				}
			}
		}
	})
	return tags
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findStateField(access *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_nobjectfield {
	if access.Hx_Field_target != nil {
		return nil
	}

	var stateFields HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = stateFields
	stateFields = this.VTable.Hx_Field_getVisibleStateFields((func() *Hx_Obj_loreline_node {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < stateFields.Len()) {
				break
			}

			var field *Hx_Obj_loreline_nobjectfield
			_ = field
			field = stateFields.Get(_g)
			_g++
			if field.Hx_Field_name == access.Hx_Field_name {
				return field
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findOutboundBeats(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var targetBeats_map *Hx_Obj_loreline_int64map
	_ = targetBeats_map
	targetBeats_map = Hx_Obj_loreline_int64map_CreateInstance()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_514 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_514
				_hx_closure_src_514 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_514(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				switch _g {
				case Hx_Obj_loreline_ncall_RTTI:
					{
						var call *Hx_Obj_loreline_ncall
						_ = call
						call = node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()
						if Hx_Field_std_isOfType(call.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
							var access *Hx_Obj_loreline_naccess
							_ = access
							access = call.Hx_Field_target.VTable.(interface {
								Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
							}).Hx_As_Hx_Obj_loreline_naccess()
							if access.Hx_Field_target == nil {
								var targetBeat *Hx_Obj_loreline_nbeatdecl
								_ = targetBeat
								targetBeat = _gthis.VTable.Hx_Field_findBeatFromAccess(access)
								if targetBeat != nil {
									var key Hx_Typedef_haxe__int64___int64
									_ = key
									key = targetBeat.Hx_Field_id
									var value *Hx_Obj_loreline_reference
									_ = value
									value = Hx_Obj_loreline_reference_CreateInstance(targetBeat, (func() *Hx_Obj_loreline_node {
										_hx_up := call
										if _hx_up != nil {
											return &_hx_up.Hx_Obj_loreline_node
										}
										return nil
									})())
									{
										var key1 Hx_Typedef_haxe__int64___int64
										_ = key1
										key1 = key
										var _hx_tmp_1 int64
										_ = _hx_tmp_1
										_hx_tmp_1 = ((int64)(key1))
										func() func(high int32, low int32, value *Hx_Obj_loreline_reference) {
											var _hx_closure_src_515 func(high int32, low int32, value any)
											_ = _hx_closure_src_515
											_hx_closure_src_515 = targetBeats_map.VTable.Hx_Field__set
											return func(high int32, low int32, value *Hx_Obj_loreline_reference) {
												_hx_closure_src_515(high, low, ((any)(value)))
											}
										}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key1)), value)
									}
								}
							}
						}
					}

				case Hx_Obj_loreline_ninsertion_RTTI:
					{
						var insertion *Hx_Obj_loreline_ninsertion
						_ = insertion
						insertion = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
						}).Hx_As_Hx_Obj_loreline_ninsertion()
						if insertion.Hx_Field_targetExpr == nil {
							var _hx_tmp_2 string
							_ = _hx_tmp_2
							_hx_tmp_2 = insertion.Hx_Field_target
							var targetBeat *Hx_Obj_loreline_nbeatdecl
							_ = targetBeat
							targetBeat = _gthis.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
								_hx_up := insertion
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							if targetBeat != nil {
								var key Hx_Typedef_haxe__int64___int64
								_ = key
								key = targetBeat.Hx_Field_id
								var value *Hx_Obj_loreline_reference
								_ = value
								value = Hx_Obj_loreline_reference_CreateInstance(targetBeat, (func() *Hx_Obj_loreline_node {
									_hx_up := insertion
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_node
									}
									return nil
								})())
								{
									var key1 Hx_Typedef_haxe__int64___int64
									_ = key1
									key1 = key
									var _hx_tmp_3 int64
									_ = _hx_tmp_3
									_hx_tmp_3 = ((int64)(key1))
									func() func(high int32, low int32, value *Hx_Obj_loreline_reference) {
										var _hx_closure_src_516 func(high int32, low int32, value any)
										_ = _hx_closure_src_516
										_hx_closure_src_516 = targetBeats_map.VTable.Hx_Field__set
										return func(high int32, low int32, value *Hx_Obj_loreline_reference) {
											_hx_closure_src_516(high, low, ((any)(value)))
										}
									}()(((int32)((_hx_tmp_3 >> int64(((int32)(32)))))), ((int32)(key1)), value)
								}
							}
						}
					}

				case Hx_Obj_loreline_ntransition_RTTI:
					{
						var transition *Hx_Obj_loreline_ntransition
						_ = transition
						transition = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
						}).Hx_As_Hx_Obj_loreline_ntransition()
						if transition.Hx_Field_targetExpr == nil {
							var _hx_tmp_4 string
							_ = _hx_tmp_4
							_hx_tmp_4 = transition.Hx_Field_target
							var targetBeat *Hx_Obj_loreline_nbeatdecl
							_ = targetBeat
							targetBeat = _gthis.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
								_hx_up := transition
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							if targetBeat != nil {
								var key Hx_Typedef_haxe__int64___int64
								_ = key
								key = targetBeat.Hx_Field_id
								var value *Hx_Obj_loreline_reference
								_ = value
								value = Hx_Obj_loreline_reference_CreateInstance(targetBeat, (func() *Hx_Obj_loreline_node {
									_hx_up := transition
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_node
									}
									return nil
								})())
								{
									var key1 Hx_Typedef_haxe__int64___int64
									_ = key1
									key1 = key
									var _hx_tmp_5 int64
									_ = _hx_tmp_5
									_hx_tmp_5 = ((int64)(key1))
									func() func(high int32, low int32, value *Hx_Obj_loreline_reference) {
										var _hx_closure_src_517 func(high int32, low int32, value any)
										_ = _hx_closure_src_517
										_hx_closure_src_517 = targetBeats_map.VTable.Hx_Field__set
										return func(high int32, low int32, value *Hx_Obj_loreline_reference) {
											_hx_closure_src_517(high, low, ((any)(value)))
										}
									}()(((int32)((_hx_tmp_5 >> int64(((int32)(32)))))), ((int32)(key1)), value)
								}
							}
						}
					}

				default:
					{
					}
				}
			}
		}

		return true
	})
	var _g HxArray[*Hx_Obj_loreline_reference]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_reference]()
	{
		var _g_map *Hx_Obj_loreline_int64map
		_ = _g_map
		_g_map = targetBeats_map
		var _g_index int32
		_ = _g_index
		_g_index = 0
		for {
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _g_index
			if _hx_tmp_7 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len() {
				_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
			} else {
				_hx_tmp_6 = false
			}

			if !(_hx_tmp_6) {
				break
			}

			_g_index++
		}

		for {
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _g_index
			if !(_hx_tmp_8 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len()) {
				break
			}

			var hx_dyncast_157 any
			_ = hx_dyncast_157
			hx_dyncast_157 = Hx_Field_go_haxe_hxdynamic_castClass(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), "Hx_Obj_loreline_reference")
			var _hx_tmp_9 *Hx_Obj_loreline_reference
			_ = _hx_tmp_9
			if hx_dyncast_157 != nil {
				_hx_tmp_9 = hx_dyncast_157.(interface {
					Hx_As_Hx_Obj_loreline_reference() *Hx_Obj_loreline_reference
				}).Hx_As_Hx_Obj_loreline_reference()
			} else {
				_hx_tmp_9 = nil
			}

			var v *Hx_Obj_loreline_reference
			_ = v
			v = _hx_tmp_9
			_g_index++
			for {
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _g_index
				if _hx_tmp_11 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len() {
					_hx_tmp_10 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
				} else {
					_hx_tmp_10 = false
				}

				if !(_hx_tmp_10) {
					break
				}

				_g_index++
			}

			var ref *Hx_Obj_loreline_reference
			_ = ref
			ref = v
			_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_g, ref))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findReferencesToBeat(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var references HxArray[*Hx_Obj_loreline_reference]
	_ = references
	references = HxMakeArray[*Hx_Obj_loreline_reference]()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	this.Hx_Field_script.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) {
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_518 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_518
			_hx_closure_src_518 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_518(((any)(o)))))
			}
		}()(node)
		if _g != nil {
			switch _g {
			case Hx_Obj_loreline_ncall_RTTI:
				{
					var call *Hx_Obj_loreline_ncall
					_ = call
					call = node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()
					if Hx_Field_std_isOfType(call.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
						var access *Hx_Obj_loreline_naccess
						_ = access
						access = call.Hx_Field_target.VTable.(interface {
							Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
						}).Hx_As_Hx_Obj_loreline_naccess()
						if (access.Hx_Field_target == nil) && (access.Hx_Field_name == beatDecl.Hx_Field_name) {
							var foundBeat *Hx_Obj_loreline_nbeatdecl
							_ = foundBeat
							foundBeat = _gthis.VTable.Hx_Field_findBeatFromAccess(access)
							var _hx_tmp_0 bool
							_ = _hx_tmp_0
							if foundBeat != nil {
								_hx_tmp_0 = Hx_Field_loreline__node_nodeid_impl__equals(foundBeat.Hx_Field_id, beatDecl.Hx_Field_id)
							} else {
								_hx_tmp_0 = false
							}

							if _hx_tmp_0 {
								var _hx_tmp_1 HxArray[*Hx_Obj_loreline_reference]
								_ = _hx_tmp_1
								_hx_tmp_1 = references
								_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_hx_tmp_1, Hx_Obj_loreline_reference_CreateInstance(beatDecl, (func() *Hx_Obj_loreline_node {
									_hx_up := call
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_node
									}
									return nil
								})())))
							}
						}
					}
				}

			case Hx_Obj_loreline_ninsertion_RTTI:
				{
					var insertion *Hx_Obj_loreline_ninsertion
					_ = insertion
					insertion = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
					}).Hx_As_Hx_Obj_loreline_ninsertion()
					if insertion.Hx_Field_target == beatDecl.Hx_Field_name {
						var _hx_tmp_2 HxArray[*Hx_Obj_loreline_reference]
						_ = _hx_tmp_2
						_hx_tmp_2 = references
						_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_hx_tmp_2, Hx_Obj_loreline_reference_CreateInstance(beatDecl, (func() *Hx_Obj_loreline_node {
							_hx_up := insertion
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})())))
					}
				}

			case Hx_Obj_loreline_ntransition_RTTI:
				{
					var transition *Hx_Obj_loreline_ntransition
					_ = transition
					transition = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
					}).Hx_As_Hx_Obj_loreline_ntransition()
					if transition.Hx_Field_target == beatDecl.Hx_Field_name {
						var _hx_tmp_3 HxArray[*Hx_Obj_loreline_reference]
						_ = _hx_tmp_3
						_hx_tmp_3 = references
						_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_hx_tmp_3, Hx_Obj_loreline_reference_CreateInstance(beatDecl, (func() *Hx_Obj_loreline_node {
							_hx_up := transition
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})())))
					}
				}

			default:
				{
				}
			}
		}
	})
	return references
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findModifiedStateFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var modifiedFields Hx_Typedef_map
	_ = modifiedFields
	modifiedFields = (func() *Hx_Obj_haxe_imap {
		_hx_icast_160 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_160 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_160, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_160.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_519 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_519
				_hx_closure_src_519 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_519(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				if _g == Hx_Obj_loreline_nassign_RTTI {
					var assign *Hx_Obj_loreline_nassign
					_ = assign
					assign = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
					}).Hx_As_Hx_Obj_loreline_nassign()
					if Hx_Field_std_isOfType(assign.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
						var access *Hx_Obj_loreline_naccess
						_ = access
						access = assign.Hx_Field_target.VTable.(interface {
							Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
						}).Hx_As_Hx_Obj_loreline_naccess()
						var field *Hx_Obj_loreline_nobjectfield
						_ = field
						field = _gthis.VTable.Hx_Field_findStateField(access)
						if field != nil {
							var key string
							_ = key
							key = field.Hx_Field_name
							var value *Hx_Obj_loreline_reference
							_ = value
							value = Hx_Obj_loreline_reference_CreateInstance(field, (func() *Hx_Obj_loreline_node {
								_hx_up := assign
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_reference) {
								var _hx_closure_src_520 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_520
								_hx_closure_src_520 = modifiedFields.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_reference) {
									_hx_closure_src_520(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(key, value)
						}
					}
				}
			}
		}

		return true
	})
	var _g HxArray[*Hx_Obj_loreline_reference]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_reference]()
	{
		var ref any
		_ = ref
		ref = modifiedFields.VTable.Hx_Field_iterator()
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(ref, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(ref, "next")
			var hx_dyncast_161 any
			_ = hx_dyncast_161
			hx_dyncast_161 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_reference")
			var _hx_tmp_3 *Hx_Obj_loreline_reference
			_ = _hx_tmp_3
			if hx_dyncast_161 != nil {
				_hx_tmp_3 = hx_dyncast_161.(interface {
					Hx_As_Hx_Obj_loreline_reference() *Hx_Obj_loreline_reference
				}).Hx_As_Hx_Obj_loreline_reference()
			} else {
				_hx_tmp_3 = nil
			}

			var ref1 *Hx_Obj_loreline_reference
			_ = ref1
			ref1 = _hx_tmp_3
			_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_g, ref1))
		}
	}

	var refs HxArray[*Hx_Obj_loreline_reference]
	_ = refs
	refs = _g
	var _hx_tmp_4 HxArray[*Hx_Obj_loreline_reference]
	_ = _hx_tmp_4
	_hx_tmp_4 = refs
	(Hx_Array_Sort[*Hx_Obj_loreline_reference](_hx_tmp_4, func(a *Hx_Obj_loreline_reference, b *Hx_Obj_loreline_reference) int32 {
		var hx_dyncast_158 any
		_ = hx_dyncast_158
		hx_dyncast_158 = Hx_Field_go_haxe_hxdynamic_castClass(a.Hx_Field_target, "Hx_Obj_loreline_nobjectfield")
		var _hx_tmp_5 *Hx_Obj_loreline_nobjectfield
		_ = _hx_tmp_5
		if hx_dyncast_158 != nil {
			_hx_tmp_5 = hx_dyncast_158.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()
		} else {
			_hx_tmp_5 = nil
		}

		var aName string
		_ = aName
		aName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_5.Hx_Field_name)
		var hx_dyncast_159 any
		_ = hx_dyncast_159
		hx_dyncast_159 = Hx_Field_go_haxe_hxdynamic_castClass(b.Hx_Field_target, "Hx_Obj_loreline_nobjectfield")
		var _hx_tmp_6 *Hx_Obj_loreline_nobjectfield
		_ = _hx_tmp_6
		if hx_dyncast_159 != nil {
			_hx_tmp_6 = hx_dyncast_159.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()
		} else {
			_hx_tmp_6 = nil
		}

		var bName string
		_ = bName
		bName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_6.Hx_Field_name)
		if aName < bName {
			return -1
		} else {
			if aName > bName {
				return 1
			} else {
				return 0
			}
		}
	}))
	return refs
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findReadStateFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var readFields Hx_Typedef_map
	_ = readFields
	readFields = (func() *Hx_Obj_haxe_imap {
		_hx_icast_164 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_164 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_164, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_164.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_521 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_521
				_hx_closure_src_521 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_521(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				if _g == Hx_Obj_loreline_naccess_RTTI {
					var access *Hx_Obj_loreline_naccess
					_ = access
					access = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
					}).Hx_As_Hx_Obj_loreline_naccess()
					if Hx_Field_std_isOfType(parent, Hx_Obj_loreline_nassign_RTTI) {
						var assign *Hx_Obj_loreline_nassign
						_ = assign
						assign = parent.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
						}).Hx_As_Hx_Obj_loreline_nassign()
						if (func() *Hx_Obj_loreline_node {
							_hx_up := assign.Hx_Field_target
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})() == node {
							return true
						}
					}

					var field *Hx_Obj_loreline_nobjectfield
					_ = field
					field = _gthis.VTable.Hx_Field_findStateField(access)
					if field != nil {
						var key string
						_ = key
						key = field.Hx_Field_name
						var value *Hx_Obj_loreline_reference
						_ = value
						value = Hx_Obj_loreline_reference_CreateInstance(field, (func() *Hx_Obj_loreline_node {
							_hx_up := access
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})())
						func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_reference) {
							var _hx_closure_src_522 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_522
							_hx_closure_src_522 = readFields.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_reference) {
								_hx_closure_src_522(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(key, value)
					}
				}
			}
		}

		return true
	})
	var _g HxArray[*Hx_Obj_loreline_reference]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_reference]()
	{
		var ref any
		_ = ref
		ref = readFields.VTable.Hx_Field_iterator()
		for {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(ref, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(ref, "next")
			var hx_dyncast_165 any
			_ = hx_dyncast_165
			hx_dyncast_165 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_reference")
			var _hx_tmp_3 *Hx_Obj_loreline_reference
			_ = _hx_tmp_3
			if hx_dyncast_165 != nil {
				_hx_tmp_3 = hx_dyncast_165.(interface {
					Hx_As_Hx_Obj_loreline_reference() *Hx_Obj_loreline_reference
				}).Hx_As_Hx_Obj_loreline_reference()
			} else {
				_hx_tmp_3 = nil
			}

			var ref1 *Hx_Obj_loreline_reference
			_ = ref1
			ref1 = _hx_tmp_3
			_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_g, ref1))
		}
	}

	var refs HxArray[*Hx_Obj_loreline_reference]
	_ = refs
	refs = _g
	var _hx_tmp_4 HxArray[*Hx_Obj_loreline_reference]
	_ = _hx_tmp_4
	_hx_tmp_4 = refs
	(Hx_Array_Sort[*Hx_Obj_loreline_reference](_hx_tmp_4, func(a *Hx_Obj_loreline_reference, b *Hx_Obj_loreline_reference) int32 {
		var hx_dyncast_162 any
		_ = hx_dyncast_162
		hx_dyncast_162 = Hx_Field_go_haxe_hxdynamic_castClass(a.Hx_Field_target, "Hx_Obj_loreline_nobjectfield")
		var _hx_tmp_5 *Hx_Obj_loreline_nobjectfield
		_ = _hx_tmp_5
		if hx_dyncast_162 != nil {
			_hx_tmp_5 = hx_dyncast_162.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()
		} else {
			_hx_tmp_5 = nil
		}

		var aName string
		_ = aName
		aName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_5.Hx_Field_name)
		var hx_dyncast_163 any
		_ = hx_dyncast_163
		hx_dyncast_163 = Hx_Field_go_haxe_hxdynamic_castClass(b.Hx_Field_target, "Hx_Obj_loreline_nobjectfield")
		var _hx_tmp_6 *Hx_Obj_loreline_nobjectfield
		_ = _hx_tmp_6
		if hx_dyncast_163 != nil {
			_hx_tmp_6 = hx_dyncast_163.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()
		} else {
			_hx_tmp_6 = nil
		}

		var bName string
		_ = bName
		bName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_6.Hx_Field_name)
		if aName < bName {
			return -1
		} else {
			if aName > bName {
				return 1
			} else {
				return 0
			}
		}
	}))
	return refs
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findBeatCharacters(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var characters_map *Hx_Obj_loreline_int64map
	_ = characters_map
	characters_map = Hx_Obj_loreline_int64map_CreateInstance()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_523 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_523
				_hx_closure_src_523 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_523(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				switch _g {
				case Hx_Obj_loreline_naccess_RTTI:
					{
						var access *Hx_Obj_loreline_naccess
						_ = access
						access = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
						}).Hx_As_Hx_Obj_loreline_naccess()
						if access.Hx_Field_target == nil {
							var _hx_tmp_1 string
							_ = _hx_tmp_1
							_hx_tmp_1 = access.Hx_Field_name
							var character *Hx_Obj_loreline_ncharacterdecl
							_ = character
							character = _gthis.VTable.Hx_Field_findCharacterByNameFromNode(_hx_tmp_1, (func() *Hx_Obj_loreline_node {
								_hx_up := access
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							if character != nil {
								var key Hx_Typedef_haxe__int64___int64
								_ = key
								key = character.Hx_Field_id
								var value *Hx_Obj_loreline_reference
								_ = value
								value = Hx_Obj_loreline_reference_CreateInstance(character, (func() *Hx_Obj_loreline_node {
									_hx_up := access
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_node
									}
									return nil
								})())
								{
									var key1 Hx_Typedef_haxe__int64___int64
									_ = key1
									key1 = key
									var _hx_tmp_2 int64
									_ = _hx_tmp_2
									_hx_tmp_2 = ((int64)(key1))
									func() func(high int32, low int32, value *Hx_Obj_loreline_reference) {
										var _hx_closure_src_524 func(high int32, low int32, value any)
										_ = _hx_closure_src_524
										_hx_closure_src_524 = characters_map.VTable.Hx_Field__set
										return func(high int32, low int32, value *Hx_Obj_loreline_reference) {
											_hx_closure_src_524(high, low, ((any)(value)))
										}
									}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(key1)), value)
								}
							}
						}
					}

				case Hx_Obj_loreline_ndialoguestatement_RTTI:
					{
						var dialogue *Hx_Obj_loreline_ndialoguestatement
						_ = dialogue
						dialogue = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
						}).Hx_As_Hx_Obj_loreline_ndialoguestatement()
						var character *Hx_Obj_loreline_ncharacterdecl
						_ = character
						character = _gthis.VTable.Hx_Field_findCharacterFromDialogue(dialogue)
						if character != nil {
							var key Hx_Typedef_haxe__int64___int64
							_ = key
							key = character.Hx_Field_id
							var value *Hx_Obj_loreline_reference
							_ = value
							value = Hx_Obj_loreline_reference_CreateInstance(character, (func() *Hx_Obj_loreline_node {
								_hx_up := dialogue
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							{
								var key1 Hx_Typedef_haxe__int64___int64
								_ = key1
								key1 = key
								var _hx_tmp_3 int64
								_ = _hx_tmp_3
								_hx_tmp_3 = ((int64)(key1))
								func() func(high int32, low int32, value *Hx_Obj_loreline_reference) {
									var _hx_closure_src_525 func(high int32, low int32, value any)
									_ = _hx_closure_src_525
									_hx_closure_src_525 = characters_map.VTable.Hx_Field__set
									return func(high int32, low int32, value *Hx_Obj_loreline_reference) {
										_hx_closure_src_525(high, low, ((any)(value)))
									}
								}()(((int32)((_hx_tmp_3 >> int64(((int32)(32)))))), ((int32)(key1)), value)
							}
						}
					}

				default:
					{
					}
				}
			}
		}

		return true
	})
	var _g HxArray[*Hx_Obj_loreline_reference]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_reference]()
	{
		var _g_map *Hx_Obj_loreline_int64map
		_ = _g_map
		_g_map = characters_map
		var _g_index int32
		_ = _g_index
		_g_index = 0
		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g_index
			if _hx_tmp_5 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len() {
				_hx_tmp_4 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
			} else {
				_hx_tmp_4 = false
			}

			if !(_hx_tmp_4) {
				break
			}

			_g_index++
		}

		for {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = _g_index
			if !(_hx_tmp_6 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len()) {
				break
			}

			var hx_dyncast_168 any
			_ = hx_dyncast_168
			hx_dyncast_168 = Hx_Field_go_haxe_hxdynamic_castClass(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), "Hx_Obj_loreline_reference")
			var _hx_tmp_7 *Hx_Obj_loreline_reference
			_ = _hx_tmp_7
			if hx_dyncast_168 != nil {
				_hx_tmp_7 = hx_dyncast_168.(interface {
					Hx_As_Hx_Obj_loreline_reference() *Hx_Obj_loreline_reference
				}).Hx_As_Hx_Obj_loreline_reference()
			} else {
				_hx_tmp_7 = nil
			}

			var v *Hx_Obj_loreline_reference
			_ = v
			v = _hx_tmp_7
			_g_index++
			for {
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g_index
				if _hx_tmp_9 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Len() {
					_hx_tmp_8 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_reference](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
				} else {
					_hx_tmp_8 = false
				}

				if !(_hx_tmp_8) {
					break
				}

				_g_index++
			}

			var ref *Hx_Obj_loreline_reference
			_ = ref
			ref = v
			_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_g, ref))
		}
	}

	var refs HxArray[*Hx_Obj_loreline_reference]
	_ = refs
	refs = _g
	var _hx_tmp_10 HxArray[*Hx_Obj_loreline_reference]
	_ = _hx_tmp_10
	_hx_tmp_10 = refs
	(Hx_Array_Sort[*Hx_Obj_loreline_reference](_hx_tmp_10, func(a *Hx_Obj_loreline_reference, b *Hx_Obj_loreline_reference) int32 {
		var hx_dyncast_166 any
		_ = hx_dyncast_166
		hx_dyncast_166 = Hx_Field_go_haxe_hxdynamic_castClass(a.Hx_Field_target, "Hx_Obj_loreline_ncharacterdecl")
		var _hx_tmp_11 *Hx_Obj_loreline_ncharacterdecl
		_ = _hx_tmp_11
		if hx_dyncast_166 != nil {
			_hx_tmp_11 = hx_dyncast_166.(interface {
				Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
			}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
		} else {
			_hx_tmp_11 = nil
		}

		var aName string
		_ = aName
		aName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_11.Hx_Field_name)
		var hx_dyncast_167 any
		_ = hx_dyncast_167
		hx_dyncast_167 = Hx_Field_go_haxe_hxdynamic_castClass(b.Hx_Field_target, "Hx_Obj_loreline_ncharacterdecl")
		var _hx_tmp_12 *Hx_Obj_loreline_ncharacterdecl
		_ = _hx_tmp_12
		if hx_dyncast_167 != nil {
			_hx_tmp_12 = hx_dyncast_167.(interface {
				Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
			}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
		} else {
			_hx_tmp_12 = nil
		}

		var bName string
		_ = bName
		bName = Hx_Field_go_haxe_hxstring_toLowerCase(_hx_tmp_12.Hx_Field_name)
		if aName < bName {
			return -1
		} else {
			if aName > bName {
				return 1
			} else {
				return 0
			}
		}
	}))
	return refs
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findModifiedCharacterFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var used_map *Hx_Obj_loreline_int64map
	_ = used_map
	used_map = Hx_Obj_loreline_int64map_CreateInstance()
	var refs HxArray[*Hx_Obj_loreline_reference]
	_ = refs
	refs = HxMakeArray[*Hx_Obj_loreline_reference]()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_526 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_526
				_hx_closure_src_526 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_526(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				if _g == Hx_Obj_loreline_nassign_RTTI {
					var assign *Hx_Obj_loreline_nassign
					_ = assign
					assign = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
					}).Hx_As_Hx_Obj_loreline_nassign()
					if Hx_Field_std_isOfType(assign.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
						var access *Hx_Obj_loreline_naccess
						_ = access
						access = assign.Hx_Field_target.VTable.(interface {
							Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
						}).Hx_As_Hx_Obj_loreline_naccess()
						var resolved *Hx_Obj_loreline_node
						_ = resolved
						resolved = _gthis.VTable.Hx_Field_resolveAccess(access)
						if Hx_Field_std_isOfType(resolved, Hx_Obj_loreline_nobjectfield_RTTI) {
							var key Hx_Typedef_haxe__int64___int64
							_ = key
							key = resolved.Hx_Field_id
							var _hx_tmp_1 int64
							_ = _hx_tmp_1
							_hx_tmp_1 = ((int64)(key))
							if !used_map.VTable.Hx_Field__exists(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key))) {
								var parent *Hx_Obj_loreline_node
								_ = parent
								parent = _gthis.VTable.Hx_Field_getParentNode(resolved)
								if Hx_Field_std_isOfType(parent, Hx_Obj_loreline_ncharacterdecl_RTTI) {
									{
										var tmp_key_1 Hx_Typedef_haxe__int64___int64
										_ = tmp_key_1
										tmp_key_1 = resolved.Hx_Field_id
										var _hx_tmp_2 int64
										_ = _hx_tmp_2
										_hx_tmp_2 = ((int64)(tmp_key_1))
										func() func(high int32, low int32, value bool) {
											var _hx_closure_src_527 func(high int32, low int32, value any)
											_ = _hx_closure_src_527
											_hx_closure_src_527 = used_map.VTable.Hx_Field__set
											return func(high int32, low int32, value bool) {
												_hx_closure_src_527(high, low, ((any)(value)))
											}
										}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), true)
									}

									var _hx_tmp_3 HxArray[*Hx_Obj_loreline_reference]
									_ = _hx_tmp_3
									_hx_tmp_3 = refs
									_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_hx_tmp_3, Hx_Obj_loreline_reference_CreateInstance(resolved.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
									}).Hx_As_Hx_Obj_loreline_nobjectfield(), node)))
								}
							}
						}
					}
				}
			}
		}

		return true
	})
	return refs
}

func (this *Hx_Obj_loreline_lens) Hx_Field_findReadCharacterFields(beatDecl *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_reference] {
	var used_map *Hx_Obj_loreline_int64map
	_ = used_map
	used_map = Hx_Obj_loreline_int64map_CreateInstance()
	var refs HxArray[*Hx_Obj_loreline_reference]
	_ = refs
	refs = HxMakeArray[*Hx_Obj_loreline_reference]()
	var _gthis *Hx_Obj_loreline_lens
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := beatDecl
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
		{
			var _g *Hx_Obj_go_haxe_hxclass
			_ = _g
			_g = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_528 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_528
				_hx_closure_src_528 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_528(((any)(o)))))
				}
			}()(node)
			if _g != nil {
				if _g == Hx_Obj_loreline_naccess_RTTI {
					var access *Hx_Obj_loreline_naccess
					_ = access
					access = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
					}).Hx_As_Hx_Obj_loreline_naccess()
					var parent *Hx_Obj_loreline_node
					_ = parent
					parent = _gthis.VTable.Hx_Field_getParentNode((func() *Hx_Obj_loreline_node {
						_hx_up := access
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})())
					if Hx_Field_std_isOfType(parent, Hx_Obj_loreline_nassign_RTTI) {
						var assign *Hx_Obj_loreline_nassign
						_ = assign
						assign = parent.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
						}).Hx_As_Hx_Obj_loreline_nassign()
						var _hx_tmp_1 *Hx_Obj_loreline_nexpr
						_ = _hx_tmp_1
						_hx_tmp_1 = assign.Hx_Field_target
						if _hx_tmp_1 == (func() *Hx_Obj_loreline_nexpr {
							_hx_up := access
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_nexpr
							}
							return nil
						})() {
							return true
						}
					}

					var resolved *Hx_Obj_loreline_node
					_ = resolved
					resolved = _gthis.VTable.Hx_Field_resolveAccess(access)
					if Hx_Field_std_isOfType(resolved, Hx_Obj_loreline_nobjectfield_RTTI) {
						var key Hx_Typedef_haxe__int64___int64
						_ = key
						key = resolved.Hx_Field_id
						var _hx_tmp_2 int64
						_ = _hx_tmp_2
						_hx_tmp_2 = ((int64)(key))
						if !used_map.VTable.Hx_Field__exists(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(key))) {
							var tmp_parent_1 *Hx_Obj_loreline_node
							_ = tmp_parent_1
							tmp_parent_1 = _gthis.VTable.Hx_Field_getParentNode(resolved)
							if Hx_Field_std_isOfType(tmp_parent_1, Hx_Obj_loreline_ncharacterdecl_RTTI) {
								{
									var tmp_key_1 Hx_Typedef_haxe__int64___int64
									_ = tmp_key_1
									tmp_key_1 = resolved.Hx_Field_id
									var _hx_tmp_3 int64
									_ = _hx_tmp_3
									_hx_tmp_3 = ((int64)(tmp_key_1))
									func() func(high int32, low int32, value bool) {
										var _hx_closure_src_529 func(high int32, low int32, value any)
										_ = _hx_closure_src_529
										_hx_closure_src_529 = used_map.VTable.Hx_Field__set
										return func(high int32, low int32, value bool) {
											_hx_closure_src_529(high, low, ((any)(value)))
										}
									}()(((int32)((_hx_tmp_3 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), true)
								}

								var _hx_tmp_4 HxArray[*Hx_Obj_loreline_reference]
								_ = _hx_tmp_4
								_hx_tmp_4 = refs
								_ = (Hx_Array_Push[*Hx_Obj_loreline_reference](_hx_tmp_4, Hx_Obj_loreline_reference_CreateInstance(resolved.VTable.(interface {
									Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
								}).Hx_As_Hx_Obj_loreline_nobjectfield(), node)))
							}
						}
					}
				}
			}
		}

		return true
	})
	return refs
}

func (this *Hx_Obj_loreline_lens) Hx_Field_traverse(node *Hx_Obj_loreline_node, callback func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool) {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = node.Hx_Field_id
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(key))
	var children HxArray[*Hx_Obj_loreline_node]
	_ = children
	children = func() func(high int32, low int32) HxArray[*Hx_Obj_loreline_node] {
		var _hx_closure_src_530 func(high int32, low int32) any
		_ = _hx_closure_src_530
		_hx_closure_src_530 = this.Hx_Field_childNodes.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) HxArray[*Hx_Obj_loreline_node] {
			return HxAnyToTypedArray[*Hx_Obj_loreline_node](_hx_closure_src_530(high, low))
		}
	}()(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key)))
	if children != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = children.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			var child *Hx_Obj_loreline_node
			_ = child
			child = children.Get(i)
			if callback(child, node) {
				this.VTable.Hx_Field_traverse(child, callback)
			}
		}
	}
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getFuncLorscript(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) *Hx_Obj_loreline_funclorscript {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = _hx_reserved_func.Hx_Field_id
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var info *Hx_Obj_loreline_funclorscript
	_ = info
	info = func() func(high int32, low int32) *Hx_Obj_loreline_funclorscript {
		var _hx_closure_src_531 func(high int32, low int32) any
		_ = _hx_closure_src_531
		_hx_closure_src_531 = this.Hx_Field_lorscriptFunctions.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_funclorscript {
			var hx_dyncast_169 any
			_ = hx_dyncast_169
			hx_dyncast_169 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_531(high, low), "Hx_Obj_loreline_funclorscript")
			var _hx_tmp_0 *Hx_Obj_loreline_funclorscript
			_ = _hx_tmp_0
			if hx_dyncast_169 != nil {
				_hx_tmp_0 = hx_dyncast_169.(interface {
					Hx_As_Hx_Obj_loreline_funclorscript() *Hx_Obj_loreline_funclorscript
				}).Hx_As_Hx_Obj_loreline_funclorscript()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if info == nil {
		info = Hx_Obj_loreline_funclorscript_CreateInstance(_hx_reserved_func)
		{
			var tmp_key_1 Hx_Typedef_haxe__int64___int64
			_ = tmp_key_1
			tmp_key_1 = id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(tmp_key_1))
			func() func(high int32, low int32, value *Hx_Obj_loreline_funclorscript) {
				var _hx_closure_src_532 func(high int32, low int32, value any)
				_ = _hx_closure_src_532
				_hx_closure_src_532 = this.Hx_Field_lorscriptFunctions.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_funclorscript) {
					_hx_closure_src_532(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), info)
		}
	}

	return info
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getLorscriptExpr(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) any {
	var info *Hx_Obj_loreline_funclorscript
	_ = info
	info = this.VTable.Hx_Field_getFuncLorscript(_hx_reserved_func)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(info.Hx_Field_expr, nil))) {
		return nil
	}

	var offset int32
	_ = offset
	offset = (pos.Hx_Field_offset - _hx_reserved_func.Hx_Field_pos.Hx_Field_offset)
	var bestExpr any
	_ = bestExpr
	var handler func(expr any)
	_ = handler
	handler = func(expr any) {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(expr, nil))) {
			var min int32
			_ = min
			min = info.Hx_Field_codeToLorscript.VTable.Hx_Field_inputPosFromProcessedPos(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmin")))))
			var max int32
			_ = max
			max = info.Hx_Field_codeToLorscript.VTable.Hx_Field_inputPosFromProcessedPos(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(expr, "pmax")))))
			if (offset >= min) && (offset <= max) {
				bestExpr = expr
			}

			Hx_Field_loreline_lorscript_tools_iter(expr, handler)
		}
	}

	Hx_Field_loreline_lorscript_tools_iter(info.Hx_Field_expr, handler)
	return bestExpr
}

func (this *Hx_Obj_loreline_lens) Hx_Field_resolveLorscriptAccess(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, expr any) *Hx_Obj_loreline_node {
	{
		var _g Hx_Enum_loreline_lorscript_exprdef
		_ = _g
		_g = func() Hx_Enum_loreline_lorscript_exprdef {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(expr, "e"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_loreline_lorscript_exprdef
				return _hx_z
			}
			return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
		}()
		switch _g.Hx_Field_enumIndex() {
		case 1:
			{
				var name string
				_ = name
				name = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
				{
					var stateField *Hx_Obj_loreline_node
					_ = stateField
					var _hx_tmp_0 *Hx_Obj_loreline_node
					_ = _hx_tmp_0
					_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
						_hx_up := this.Hx_Field_script
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					this.VTable.Hx_Field_traverse(_hx_tmp_0, func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node) bool {
						if stateField != nil {
							return false
						}

						if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nstatedecl_RTTI) {
							var stateDecl *Hx_Obj_loreline_nstatedecl
							_ = stateDecl
							stateDecl = node.VTable.(interface {
								Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
							}).Hx_As_Hx_Obj_loreline_nstatedecl()
							{
								var tmp__g_1 int32
								_ = tmp__g_1
								tmp__g_1 = 0
								var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
								_ = _g1
								_g1 = stateDecl.Hx_Field_fields
								for {
									var _hx_tmp_1 int32
									_ = _hx_tmp_1
									_hx_tmp_1 = tmp__g_1
									if !(_hx_tmp_1 < _g1.Len()) {
										break
									}

									var field *Hx_Obj_loreline_nobjectfield
									_ = field
									field = _g1.Get(tmp__g_1)
									tmp__g_1++
									if field.Hx_Field_name == name {
										stateField = (func() *Hx_Obj_loreline_node {
											_hx_up := field
											if _hx_up != nil {
												return &_hx_up.Hx_Obj_loreline_node
											}
											return nil
										})()
										return false
									}
								}
							}
						}

						var _hx_tmp_2 bool
						_ = _hx_tmp_2
						if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nimportstatement_RTTI) {
							_hx_tmp_2 = true
						} else {
							_hx_tmp_2 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_script_RTTI)
						}

						return _hx_tmp_2
					})
					if stateField != nil {
						return stateField
					}

					var _hx_tmp_3 string
					_ = _hx_tmp_3
					_hx_tmp_3 = name
					var characterDecl *Hx_Obj_loreline_ncharacterdecl
					_ = characterDecl
					characterDecl = this.VTable.Hx_Field_findCharacterByNameFromNode(_hx_tmp_3, (func() *Hx_Obj_loreline_node {
						_hx_up := _hx_reserved_func
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})())
					if characterDecl != nil {
						return (func() *Hx_Obj_loreline_node {
							_hx_up := characterDecl
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})()
					}

					var _hx_tmp_4 string
					_ = _hx_tmp_4
					_hx_tmp_4 = name
					var beatDecl *Hx_Obj_loreline_nbeatdecl
					_ = beatDecl
					beatDecl = this.VTable.Hx_Field_findBeatByNameFromNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
						_hx_up := _hx_reserved_func
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})())
					if beatDecl != nil {
						return (func() *Hx_Obj_loreline_node {
							_hx_up := beatDecl
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})()
					}
				}
			}

		case 5:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_e
				var f string
				_ = f
				f = _g.(Hx_Enum_loreline_lorscript_exprdef_EField).Hx_Field_f
				{
					var _hx_tmp_5 *Hx_Obj_loreline_node
					_ = _hx_tmp_5
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
						_hx_tmp_5 = this.VTable.Hx_Field_resolveLorscriptAccess(_hx_reserved_func, e)
					} else {
						_hx_tmp_5 = nil
					}

					var targetNode *Hx_Obj_loreline_node
					_ = targetNode
					targetNode = _hx_tmp_5
					if targetNode != nil {
						if Hx_Field_std_isOfType(targetNode, Hx_Obj_loreline_nobjectfield_RTTI) {
							targetNode = (func() *Hx_Obj_loreline_node {
								_hx_up := targetNode.VTable.(interface {
									Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
								}).Hx_As_Hx_Obj_loreline_nobjectfield().Hx_Field_value
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
						}

						{
							var tmp__g_1 *Hx_Obj_go_haxe_hxclass
							_ = tmp__g_1
							tmp__g_1 = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
								var _hx_closure_src_533 func(o any) *Hx_Obj_go_haxe_hxclass
								_ = _hx_closure_src_533
								_hx_closure_src_533 = Hx_Field_type_getClass
								return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
									return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_533(((any)(o)))))
								}
							}()(targetNode)
							if tmp__g_1 != nil {
								switch tmp__g_1 {
								case Hx_Obj_loreline_ncharacterdecl_RTTI:
									{
										var characterDecl *Hx_Obj_loreline_ncharacterdecl
										_ = characterDecl
										characterDecl = targetNode.VTable.(interface {
											Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
										}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
										{
											var tmp__g_2 int32
											_ = tmp__g_2
											tmp__g_2 = 0
											var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
											_ = _g1
											_g1 = characterDecl.Hx_Field_fields
											for {
												var _hx_tmp_6 int32
												_ = _hx_tmp_6
												_hx_tmp_6 = tmp__g_2
												if !(_hx_tmp_6 < _g1.Len()) {
													break
												}

												var prop *Hx_Obj_loreline_nobjectfield
												_ = prop
												prop = _g1.Get(tmp__g_2)
												tmp__g_2++
												if prop.Hx_Field_name == f {
													return (func() *Hx_Obj_loreline_node {
														_hx_up := prop
														if _hx_up != nil {
															return &_hx_up.Hx_Obj_loreline_node
														}
														return nil
													})()
												}
											}
										}
									}

								case Hx_Obj_loreline_nliteral_RTTI:
									{
										var literal *Hx_Obj_loreline_nliteral
										_ = literal
										literal = targetNode.VTable.(interface {
											Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
										}).Hx_As_Hx_Obj_loreline_nliteral()
										{
											var tmp__g_2 Hx_Enum_loreline_literaltype
											_ = tmp__g_2
											tmp__g_2 = literal.Hx_Field_literalType
											if tmp__g_2.Hx_Field_enumIndex() == 4 {
												var style int32
												_ = style
												style = tmp__g_2.(Hx_Enum_loreline_literaltype_Object).Hx_Field_style
												{
													var fields HxArray[*Hx_Obj_loreline_nobjectfield]
													_ = fields
													fields = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](literal.Hx_Field_value)
													{
														var _g1 int32
														_ = _g1
														_g1 = 0
														for {
															var _hx_tmp_7 int32
															_ = _hx_tmp_7
															_hx_tmp_7 = _g1
															if !(_hx_tmp_7 < fields.Len()) {
																break
															}

															var field *Hx_Obj_loreline_nobjectfield
															_ = field
															field = fields.Get(_g1)
															_g1++
															if field.Hx_Field_name == f {
																return (func() *Hx_Obj_loreline_node {
																	_hx_up := field
																	if _hx_up != nil {
																		return &_hx_up.Hx_Obj_loreline_node
																	}
																	return nil
																})()
															}
														}
													}
												}
											}
										}
									}

								default:
									{
									}
								}
							}
						}
					}
				}
			}

		case 16:
			{
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_e
				var index any
				_ = index
				index = _g.(Hx_Enum_loreline_lorscript_exprdef_EArray).Hx_Field_index
			}

		default:
			{
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lens) Hx_Field_isLorscriptExpr(input any) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(input, nil))) {
		_hx_tmp_1 = Hx_Field_reflect_hasField(input, "e")
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = Hx_Field_std_isOfType(Hx_Field_reflect_field(input, "e"), Hx_Enum_loreline_lorscript_exprdef_RTTI)
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_lens) Hx_Field_getLorscriptCompletion(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_lorscriptcompletion {
	var info *Hx_Obj_loreline_funclorscript
	_ = info
	info = this.VTable.Hx_Field_getFuncLorscript(_hx_reserved_func)
	var inputPos int32
	_ = inputPos
	inputPos = ((pos.Hx_Field_offset - _hx_reserved_func.Hx_Field_pos.Hx_Field_offset) + pos.Hx_Field_length)
	var processedPos int32
	_ = processedPos
	processedPos = info.Hx_Field_codeToLorscript.VTable.Hx_Field_processedPosFromInputPos(inputPos)
	var str string
	_ = str
	str = info.Hx_Field_lorscript
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((float64)(processedPos))
	var endIndex HxNullable[int32]
	_ = endIndex
	endIndex = HxNullable[int32]{Value: Hx_Field_std_int(Hx_Field_math_min(_hx_tmp_0, ((float64)(int32(utf8.RuneCountInString(info.Hx_Field_lorscript)))))), Valid: true}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if (endIndex).Valid != false {
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substring(str, 0, endIndex))
	} else {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = str
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_3, 0, HxNullable[int32]{}))
	}

	var truncated string
	_ = truncated
	truncated = _hx_tmp_1
	var fullLen int32
	_ = fullLen
	fullLen = int32(utf8.RuneCountInString(truncated))
	var _hx_tmp_4 string
	_ = _hx_tmp_4
	_hx_tmp_4 = truncated
	var startIndex int32
	_ = startIndex
	startIndex = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_4, "{", HxNullable[int32]{})) + 1)
	var _hx_tmp_5 string
	_ = _hx_tmp_5
	if (HxNullable[int32]{Value: fullLen, Valid: true}).Valid != false {
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = truncated
		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		_hx_tmp_7 = startIndex
		var _hx_tmp_8 string
		_ = _hx_tmp_8
		_hx_tmp_8 = _hx_tmp_6
		var _hx_tmp_9 int32
		_ = _hx_tmp_9
		_hx_tmp_9 = _hx_tmp_7
		_hx_tmp_5 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_8, _hx_tmp_9, HxNullable[int32]{Value: fullLen, Valid: true}))
	} else {
		var _hx_tmp_10 string
		_ = _hx_tmp_10
		_hx_tmp_10 = truncated
		var _hx_tmp_11 int32
		_ = _hx_tmp_11
		_hx_tmp_11 = startIndex
		var _hx_tmp_12 string
		_ = _hx_tmp_12
		_hx_tmp_12 = _hx_tmp_10
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_11
		_hx_tmp_5 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_12, _hx_tmp_13, HxNullable[int32]{}))
	}

	truncated = _hx_tmp_5
	var truncatedLen int32
	_ = truncatedLen
	truncatedLen = int32(utf8.RuneCountInString(truncated))
	var spaces *Hx_Obj_stringbuf
	_ = spaces
	spaces = Hx_Obj_stringbuf_CreateInstance()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = (fullLen - truncatedLen)
		for _g < _g1 {
			var _hx_tmp_14 int32
			_ = _hx_tmp_14
			_hx_tmp_14 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_14
			spaces.VTable.Hx_Field_addChar(32)
		}
	}

	truncated = Hx_Field_go_haxe_hxdynamic_stringConcat(spaces.VTable.Hx_Field_toString(), truncated)
	var completion *Hx_Obj_loreline_lorscript_completion
	_ = completion
	var checker *Hx_Obj_loreline_lorscript_checker
	_ = checker
	checker = Hx_Obj_loreline_lorscript_checker_CreateInstance(nil)
	{
		var hx_try_return *Hx_Obj_loreline_lorscriptcompletion
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							if Hx_Field_std_isOfType(e, Hx_Obj_loreline_lorscript_completion_RTTI) {
								var hx_dyncast_170 any
								_ = hx_dyncast_170
								hx_dyncast_170 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_lorscript_completion")
								var _hx_tmp_19 *Hx_Obj_loreline_lorscript_completion
								_ = _hx_tmp_19
								if hx_dyncast_170 != nil {
									_hx_tmp_19 = hx_dyncast_170.(interface {
										Hx_As_Hx_Obj_loreline_lorscript_completion() *Hx_Obj_loreline_lorscript_completion
									}).Hx_As_Hx_Obj_loreline_lorscript_completion()
								} else {
									_hx_tmp_19 = nil
								}

								completion = _hx_tmp_19
							}
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				checker.Hx_Field_types.Hx_Field_parser.Hx_Field_allowJSON = true
				checker.Hx_Field_types.Hx_Field_parser.Hx_Field_allowTypes = true
				checker.Hx_Field_types.Hx_Field_parser.Hx_Field_resumeErrors = true
				var _hx_tmp_15 string
				_ = _hx_tmp_15
				_hx_tmp_15 = truncated
				var _hx_tmp_16 any
				_ = _hx_tmp_16
				_hx_tmp_16 = nil
				var expr any
				_ = expr
				expr = checker.Hx_Field_types.Hx_Field_parser.VTable.Hx_Field_parseString(_hx_tmp_15, _hx_tmp_16, nil)
				var _hx_tmp_17 any
				_ = _hx_tmp_17
				_hx_tmp_17 = expr
				var _hx_tmp_18 any
				_ = _hx_tmp_18
				_hx_tmp_18 = nil
				checker.VTable.Hx_Field_check(_hx_tmp_17, _hx_tmp_18, HxNullable[bool]{Value: true, Valid: true})
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	var locals Hx_Typedef_map
	_ = locals
	locals = checker.Hx_Field_locals
	return Hx_Obj_loreline_lorscriptcompletion_CreateInstance(locals, completion)
}

func (this *Hx_Obj_loreline_lens) Hx_Field_resolveAccessInFunction(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node {
	var info *Hx_Obj_loreline_funclorscript
	_ = info
	info = this.VTable.Hx_Field_getFuncLorscript(_hx_reserved_func)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(info.Hx_Field_expr, nil))) {
		return nil
	}

	var expr any
	_ = expr
	expr = this.VTable.Hx_Field_getLorscriptExpr(_hx_reserved_func, pos)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(expr, nil))) {
		return this.VTable.Hx_Field_resolveLorscriptAccess(_hx_reserved_func, expr)
	}

	return nil
}

func Hx_Field_loreline_lens_leftmostOperand(expr *Hx_Obj_loreline_nexpr) *Hx_Obj_loreline_nexpr {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if expr.Hx_Field_parens == 0 {
		_hx_tmp_0 = Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nbinary_RTTI)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		return Hx_Field_loreline_lens_leftmostOperand(expr.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
		}).Hx_As_Hx_Obj_loreline_nbinary().Hx_Field_left)
	}

	return expr
}

func Hx_Field_loreline_lens_topLevelClauses(expr *Hx_Obj_loreline_nexpr) HxArray[*Hx_Obj_loreline_nexpr] {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if expr.Hx_Field_parens == 0 {
		_hx_tmp_0 = Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nbinary_RTTI)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = expr.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
		}).Hx_As_Hx_Obj_loreline_nbinary()
		if binary.Hx_Field_op.Hx_Field_enumIndex() == 38 {
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_loreline_lens_topLevelClauses(binary.Hx_Field_left)
			return (Hx_Array_Concat[*Hx_Obj_loreline_nexpr](_hx_tmp_1, Hx_Field_loreline_lens_topLevelClauses(binary.Hx_Field_right)))
		}
	}

	return HxMakeArray[*Hx_Obj_loreline_nexpr](expr)
}

func Hx_Field_loreline_lens_comparedValue(expr *Hx_Obj_loreline_nexpr) HxNullable[string] {
	if !Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nbinary_RTTI) {
		return HxNullable[string]{Value: HxStringNull}
	}

	var binary *Hx_Obj_loreline_nbinary
	_ = binary
	binary = expr.VTable.(interface {
		Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
	}).Hx_As_Hx_Obj_loreline_nbinary()
	switch binary.Hx_Field_op.Hx_Field_enumIndex() {
	case 34, 35, 36, 37:
		{
		}

	default:
		{
			return HxNullable[string]{Value: HxStringNull}
		}
	}

	var _hx_tmp_0 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_0
	if Hx_Field_loreline_lens_isNumberLiteral(binary.Hx_Field_right) {
		_hx_tmp_0 = binary.Hx_Field_left
	} else {
		var _hx_tmp_1 *Hx_Obj_loreline_nexpr
		_ = _hx_tmp_1
		if Hx_Field_loreline_lens_isNumberLiteral(binary.Hx_Field_left) {
			_hx_tmp_1 = binary.Hx_Field_right
		} else {
			_hx_tmp_1 = nil
		}

		_hx_tmp_0 = _hx_tmp_1
	}

	var value *Hx_Obj_loreline_nexpr
	_ = value
	value = _hx_tmp_0
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if value == nil {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if Hx_Field_std_isOfType(value, Hx_Obj_loreline_naccess_RTTI) {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = Hx_Field_std_isOfType(value, Hx_Obj_loreline_narrayaccess_RTTI)
		}

		_hx_tmp_2 = !(_hx_tmp_3)
	}

	if _hx_tmp_2 {
		return HxNullable[string]{Value: HxStringNull}
	}

	var printer *Hx_Obj_loreline_printer
	_ = printer
	printer = Hx_Obj_loreline_printer_CreateInstance(nil, nil)
	printer.Hx_Field_enableComments = false
	return HxNullableString(Hx_Field_stringtools_trim(printer.VTable.Hx_Field_print((func() *Hx_Obj_loreline_node {
		_hx_up := value
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())))
}

func Hx_Field_loreline_lens_isNumberLiteral(expr *Hx_Obj_loreline_nexpr) bool {
	if !Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nliteral_RTTI) {
		return false
	}

	return (expr.VTable.(interface {
		Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
	}).Hx_As_Hx_Obj_loreline_nliteral().Hx_Field_literalType.Hx_Field_enumIndex() == 0)
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatdecl]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatparam]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_ncharacterdecl]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nfunctiondecl]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nimportstatement]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_node]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_reference]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
