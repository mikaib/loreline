package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_parseerror() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_parseerror() {
}

type Hx_Obj_VTable_loreline_parseerror interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_parseerror struct {
	Hx_Obj_loreline_error
	VTable Hx_Obj_VTable_loreline_parseerror `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_parseerror_CreateEmptyInstance() *Hx_Obj_loreline_parseerror {
	obj := &Hx_Obj_loreline_parseerror{}
	obj.Hx_Obj_loreline_error = *Hx_Obj_loreline_error_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_error.VTable = obj
	return obj
}

func Hx_Obj_loreline_parseerror_CreateInstance(message string, pos__isOptional any) *Hx_Obj_loreline_parseerror {
	obj := Hx_Obj_loreline_parseerror_CreateEmptyInstance()
	obj.Hx_New(message, pos__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_parseerror) Hx_New(message string, pos__isOptional any) {
	{
		pos := HxDefault[*Hx_Obj_loreline_position]()
		_ = pos

		if pos__isOptional != nil {
			pos = HxConvert[*Hx_Obj_loreline_position](pos__isOptional)
		} else {
			pos = nil
		}
		{
			this.Hx_Obj_loreline_error.Hx_New(message, pos)
		}

	}
}

var Hx_Obj_loreline_parseerror_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.ParseError",
	HxMakeArray[string](),
	HxMakeArray[string]("message", "pos", "stack", "toString"),
	Hx_Obj_loreline_error_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_parseerror_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_parseerror_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_parseerror_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_parseerror) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_parseerror_RTTI
}

func (this *Hx_Obj_loreline_parseerror) Hx_As_Hx_Obj_loreline_parseerror() *Hx_Obj_loreline_parseerror {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_parsercontext() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_parsercontext() {
}

type Hx_Obj_VTable_loreline_parsercontext interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_parsercontext struct {
	VTable            Hx_Obj_VTable_loreline_parsercontext `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_rootPath string
	Hx_Field_path     string
	Hx_Field_imports  Hx_Typedef_map
	Hx_Field_imported Hx_Typedef_map
}

func Hx_Obj_loreline_parsercontext_CreateEmptyInstance() *Hx_Obj_loreline_parsercontext {
	obj := &Hx_Obj_loreline_parsercontext{}
	obj.VTable = obj
	obj.Hx_Field_rootPath = HxDefault[string]()
	obj.Hx_Field_path = HxDefault[string]()
	obj.Hx_Field_imports = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_imported = HxDefault[Hx_Typedef_map]()
	return obj
}

func Hx_Obj_loreline_parsercontext_CreateInstance(rootPath string, path string, imports Hx_Typedef_map, imported__isOptional any) *Hx_Obj_loreline_parsercontext {
	obj := Hx_Obj_loreline_parsercontext_CreateEmptyInstance()
	obj.Hx_New(rootPath, path, imports, imported__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_parsercontext) Hx_New(rootPath string, path string, imports Hx_Typedef_map, imported__isOptional any) {
	{
		imported := HxDefault[Hx_Typedef_map]()
		_ = imported

		if imported__isOptional != nil {
			imported = HxConvert[Hx_Typedef_map](imported__isOptional)
		} else {
		}
		{
			this.Hx_Field_imported = (func() *Hx_Obj_haxe_imap {
				_hx_icast_292 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_292 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_292, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_292.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_rootPath = rootPath
			this.Hx_Field_path = path
			this.Hx_Field_imports = imports
			if imported != nil {
				this.Hx_Field_imported = imported
			}
		}

	}
}

var Hx_Obj_loreline_parsercontext_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.ParserContext",
	HxMakeArray[string](),
	HxMakeArray[string]("rootPath", "path", "imports", "imported"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_parsercontext_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[string](params.Get(1)), HxConvert[Hx_Typedef_map](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_parsercontext_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_parsercontext_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_parsercontext) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_parsercontext_RTTI
}

func (this *Hx_Obj_loreline_parsercontext) Hx_As_Hx_Obj_loreline_parsercontext() *Hx_Obj_loreline_parsercontext {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_parser() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_parser() {
}

type Hx_Obj_VTable_loreline_parser interface {
	Hx_Field_parseAdditive() *Hx_Obj_loreline_nexpr
	Hx_Field_parseStringInterpolation(braces bool, inTag bool, tokens Hx_Typedef_loreline_tokens, start int32, length int32, content string) *Hx_Obj_loreline_nstringpart
	Hx_Field_isComment(_hx_reserved_type Hx_Enum_loreline_tokentype) bool
	Hx_Field_parseWhenStatement() *Hx_Obj_loreline_nwhenstatement
	Hx_Field_expect(_hx_reserved_type Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_token
	Hx_Field_nextNonLineBreakOrComment() *Hx_Obj_loreline_token
	Hx_Field_parseImport() *Hx_Obj_loreline_nimportstatement
	Hx_Field_parseStateDecl(temporary bool) *Hx_Obj_loreline_nstatedecl
	Hx_Field_parseChoiceStatement() *Hx_Obj_loreline_nchoicestatement
	Hx_Field_parseChoiceOption(blockEnd Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_nchoiceoption
	Hx_Field_parseFunction() *Hx_Obj_loreline_nfunctiondecl
	Hx_Field_parsePrimary() *Hx_Obj_loreline_nexpr
	Hx_Field_parseStringLiteral() *Hx_Obj_loreline_nstringliteral
	Hx_Field_expectIdentifier() string
	Hx_Field_currentPos() *Hx_Obj_loreline_position
	Hx_Field_wrapInRootBeat(node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_parseDialogueStatement() *Hx_Obj_loreline_ndialoguestatement
	Hx_Field_isExpressionStart() bool
	Hx_Field_parseAlternative() *Hx_Obj_loreline_nalternative
	Hx_Field_parseLogicalOr() *Hx_Obj_loreline_nexpr
	Hx_Field_synchronize()
	Hx_Field_prevWithType(_hx_reserved_type Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_token
	Hx_Field_checkString() bool
	Hx_Field_checkBlockStart() bool
	Hx_Field_parseLogicalAnd() *Hx_Obj_loreline_nexpr
	Hx_Field_parseComparison() *Hx_Obj_loreline_nexpr
	Hx_Field_parse() *Hx_Obj_loreline_script
	Hx_Field_advance(advanceLineBreaks__isOptional any) *Hx_Obj_loreline_token
	Hx_Field_parseNode(topLevel__isOptional any) *Hx_Obj_loreline_astnode
	Hx_Field_parseConditionExpression() *Hx_Obj_loreline_nexpr
	Hx_Field_takeConditionParens(condition *Hx_Obj_loreline_nexpr) int32
	Hx_Field_addError(_hx_reserved_error *Hx_Obj_loreline_parseerror) *Hx_Obj_loreline_parseerror
	Hx_Field_parseTernary() *Hx_Obj_loreline_nexpr
	Hx_Field_previous() *Hx_Obj_loreline_token
	Hx_Field_prevIdentifier() *Hx_Obj_loreline_token
	Hx_Field_parseBeatDecl() *Hx_Obj_loreline_nbeatdecl
	Hx_Field_parseExpression() *Hx_Obj_loreline_nexpr
	Hx_Field_parseArrayLiteral() *Hx_Obj_loreline_nexpr
	Hx_Field_makeStringPartPosition(stringStart *Hx_Obj_loreline_position, content string, offset int32) *Hx_Obj_loreline_position
	Hx_Field_currentTokenTypeToString() string
	Hx_Field_check(_hx_reserved_type Hx_Enum_loreline_tokentype) bool
	Hx_Field_isAtEnd() bool
	Hx_Field_isKnownNodeStart() bool
	Hx_Field_parseBeatCall() *Hx_Obj_loreline_nbeatcall
	Hx_Field_parseUnary() *Hx_Obj_loreline_nexpr
	Hx_Field_parseStringTag(closing bool, start int32, length int32, content string, quotes int32, attachments HxArray[Hx_Enum_loreline_lstringattachment]) *Hx_Obj_loreline_nstringpart
	Hx_Field_nextNonWhitespaceOrComment() *Hx_Obj_loreline_token
	Hx_Field_parseIfStatement() *Hx_Obj_loreline_nifstatement
	Hx_Field_parseObjectLiteral() *Hx_Obj_loreline_nexpr
	Hx_Field_requiresNewLine() bool
	Hx_Field_parseBlockStart() *Hx_Obj_loreline_token
	Hx_Field_parseCharacterDecl() *Hx_Obj_loreline_ncharacterdecl
	Hx_Field_nextNodeId(step int32) Hx_Typedef_haxe__int64___int64
	Hx_Field_prevNonWhitespaceOrComment() *Hx_Obj_loreline_token
	Hx_Field_parseExpressionStatement() *Hx_Obj_loreline_astnode
	Hx_Field_parseEquality() *Hx_Obj_loreline_nexpr
	Hx_Field_parseMultiplicative() *Hx_Obj_loreline_nexpr
	Hx_Field_match(_hx_reserved_type Hx_Enum_loreline_tokentype, advanceLineBreaks__isOptional any) bool
	Hx_Field_attachElseComments(node any, elseToken *Hx_Obj_loreline_token) any
	Hx_Field_parseObjectField() *Hx_Obj_loreline_nobjectfield
	Hx_Field_parseTextStatement() *Hx_Obj_loreline_ntextstatement
	Hx_Field_parseWhenRule() *Hx_Obj_loreline_nwhenrule
	Hx_Field_parseCallArguments() HxArray[*Hx_Obj_loreline_nexpr]
	Hx_Field_isAtLineBreak() bool
	Hx_Field_getErrors() HxArray[*Hx_Obj_loreline_parseerror]
	Hx_Field_parseTransition() *Hx_Obj_loreline_ntransition
	Hx_Field_makeAccess(pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, name string, namePos *Hx_Obj_loreline_position) *Hx_Obj_loreline_naccess
	Hx_Field_checkIdentifier() bool
	Hx_Field_attachComments(node any) any
	Hx_Field_peek() *Hx_Obj_loreline_token
	Hx_Field_parseStatementBlock(statements HxArray[*Hx_Obj_loreline_astnode]) int32
	Hx_Field_parseInsertion() *Hx_Obj_loreline_ninsertion
	Hx_Field_parseIdentifierExpression(startPos *Hx_Obj_loreline_position, name string) *Hx_Obj_loreline_nexpr
	Hx_Field_parsePostfix(startPos *Hx_Obj_loreline_position, expr *Hx_Obj_loreline_nexpr) *Hx_Obj_loreline_nexpr
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_parser struct {
	VTable                       Hx_Obj_VTable_loreline_parser `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_tokens              Hx_Typedef_loreline_tokens
	Hx_Field_current             int32
	Hx_Field_errors              HxArray[*Hx_Obj_loreline_parseerror]
	Hx_Field_pendingComments     HxArray[*Hx_Obj_loreline_comment]
	Hx_Field_lastTokenEnd        *Hx_Obj_loreline_position
	Hx_Field_lastLineBreak       *Hx_Obj_loreline_position
	Hx_Field_lineBreakAfterToken bool
	Hx_Field_parsingCondition    bool
	Hx_Field_currentNodeId       Hx_Typedef_haxe__int64___int64
	Hx_Field_rootBeat            *Hx_Obj_loreline_nbeatdecl
	Hx_Field_context             *Hx_Obj_loreline_parsercontext
}

func Hx_Obj_loreline_parser_CreateEmptyInstance() *Hx_Obj_loreline_parser {
	obj := &Hx_Obj_loreline_parser{}
	obj.VTable = obj
	obj.Hx_Field_tokens = HxDefault[Hx_Typedef_loreline_tokens]()
	obj.Hx_Field_current = HxDefault[int32]()
	obj.Hx_Field_errors = HxDefault[HxArray[*Hx_Obj_loreline_parseerror]]()
	obj.Hx_Field_pendingComments = HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
	obj.Hx_Field_lastTokenEnd = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_lastLineBreak = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_lineBreakAfterToken = HxDefault[bool]()
	obj.Hx_Field_parsingCondition = HxDefault[bool]()
	obj.Hx_Field_currentNodeId = HxDefault[Hx_Typedef_haxe__int64___int64]()
	obj.Hx_Field_rootBeat = HxDefault[*Hx_Obj_loreline_nbeatdecl]()
	obj.Hx_Field_context = HxDefault[*Hx_Obj_loreline_parsercontext]()
	return obj
}

func Hx_Obj_loreline_parser_CreateInstance(tokens Hx_Typedef_loreline_tokens, context__isOptional any) *Hx_Obj_loreline_parser {
	obj := Hx_Obj_loreline_parser_CreateEmptyInstance()
	obj.Hx_New(tokens, context__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_parser) Hx_New(tokens Hx_Typedef_loreline_tokens, context__isOptional any) {
	{
		context := HxDefault[*Hx_Obj_loreline_parsercontext]()
		_ = context

		if context__isOptional != nil {
			context = HxConvert[*Hx_Obj_loreline_parsercontext](context__isOptional)
		} else {
			context = nil
		}
		{
			this.Hx_Field_parsingCondition = false
		}
		{
			this.Hx_Field_tokens = tokens
			this.Hx_Field_current = 0
			this.Hx_Field_errors = nil
			this.Hx_Field_pendingComments = nil
			this.Hx_Field_lastTokenEnd = Hx_Obj_loreline_position_CreateInstance(1, 1, 0, nil)
			this.Hx_Field_lastLineBreak = nil
			this.Hx_Field_lineBreakAfterToken = false
			this.Hx_Field_currentNodeId = Hx_Field_loreline__node_nodeid_impl__UNDEFINED
			this.Hx_Field_rootBeat = nil
			this.Hx_Field_context = context
		}

	}
}

var Hx_Obj_loreline_parser_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Parser",
	HxMakeArray[string](),
	HxMakeArray[string]("tokens", "current", "errors", "pendingComments", "lastTokenEnd", "lastLineBreak", "lineBreakAfterToken", "parsingCondition", "currentNodeId", "rootBeat", "context", "parse", "getErrors", "nextNodeId", "peek", "advance", "previous", "prevWithType", "prevIdentifier", "prevNonWhitespaceOrComment", "nextNonWhitespaceOrComment", "nextNonLineBreakOrComment", "currentPos", "currentTokenTypeToString", "check", "checkString", "isAtEnd", "isComment", "parseNode", "wrapInRootBeat", "parseImport", "parseDialogueStatement", "parseStatementBlock", "parseStateDecl", "parseObjectField", "parseBeatDecl", "checkBlockStart", "parseBlockStart", "parseCharacterDecl", "parseTextStatement", "parseChoiceStatement", "parseChoiceOption", "parseWhenStatement", "parseWhenRule", "isExpressionStart", "parseExpressionStatement", "isKnownNodeStart", "parseIfStatement", "parseAlternative", "parseTransition", "parseInsertion", "parseBeatCall", "parseFunction", "parseExpression", "parseTernary", "parseLogicalOr", "parseLogicalAnd", "parseEquality", "parseComparison", "parseAdditive", "parseMultiplicative", "parseUnary", "parsePrimary", "parseStringLiteral", "makeStringPartPosition", "makeAccess", "parseStringInterpolation", "parseStringTag", "parseIdentifierExpression", "parsePostfix", "parseArrayLiteral", "parseObjectLiteral", "parseCallArguments", "parseConditionExpression", "takeConditionParens", "match", "isAtLineBreak", "expect", "checkIdentifier", "expectIdentifier", "attachComments", "attachElseComments", "synchronize", "requiresNewLine", "addError"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_parser_CreateInstance(HxConvert[Hx_Typedef_loreline_tokens](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_parser_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_parser_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_parser_RTTI
}

func (this *Hx_Obj_loreline_parser) Hx_As_Hx_Obj_loreline_parser() *Hx_Obj_loreline_parser {
	return this
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parse() *Hx_Obj_loreline_script {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var nodes HxArray[*Hx_Obj_loreline_astnode]
	_ = nodes
	nodes = HxMakeArray[*Hx_Obj_loreline_astnode]()
	var script *Hx_Obj_loreline_script
	_ = script
	script = Hx_Obj_loreline_script_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, nodes)
	if this.Hx_Field_context != nil {
		script.Hx_Field_filePath = this.Hx_Field_context.Hx_Field_path
	}

	for {
		if !(!this.VTable.Hx_Field_isAtEnd()) {
			break
		}

		{
			var hx_try_return *Hx_Obj_loreline_script
			_ = hx_try_return
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
							_ = e
							{
								this.VTable.Hx_Field_addError(e)
								this.VTable.Hx_Field_synchronize()
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					var node *Hx_Obj_loreline_astnode
					_ = node
					node = this.VTable.Hx_Field_parseNode(true)
					if node != nil {
						_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](nodes, node))
					}

					for {
						var _hx_tmp_0 Hx_Enum_loreline_tokentype
						_ = _hx_tmp_0
						_hx_tmp_0 = Hx_Enum_loreline_tokentype_LineBreak{}

						if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
							break
						}
					}
				}
			}()
			if hx_try_state == 1 {
				return hx_try_return
			}
			if hx_try_state == 2 {
				break
			}
			if hx_try_state == 3 {
				continue
			}
		}
	}

	func() func(node *Hx_Obj_loreline_script) *Hx_Obj_loreline_script {
		var _hx_closure_src_691 func(node any) any
		_ = _hx_closure_src_691
		_hx_closure_src_691 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_script) *Hx_Obj_loreline_script {
			var hx_dyncast_293 any
			_ = hx_dyncast_293
			hx_dyncast_293 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_691(((any)(node))), "Hx_Obj_loreline_script")
			var _hx_tmp_1 *Hx_Obj_loreline_script
			_ = _hx_tmp_1
			if hx_dyncast_293 != nil {
				_hx_tmp_1 = hx_dyncast_293.(interface {
					Hx_As_Hx_Obj_loreline_script() *Hx_Obj_loreline_script
				}).Hx_As_Hx_Obj_loreline_script()
			} else {
				_hx_tmp_1 = nil
			}

			return _hx_tmp_1
		}
	}()(script)
	return script
}

func (this *Hx_Obj_loreline_parser) Hx_Field_getErrors() HxArray[*Hx_Obj_loreline_parseerror] {
	if this.Hx_Field_errors == nil {
		this.Hx_Field_errors = HxMakeArray[*Hx_Obj_loreline_parseerror]()
	}

	return this.Hx_Field_errors
}

func (this *Hx_Obj_loreline_parser) Hx_Field_nextNodeId(step int32) Hx_Typedef_haxe__int64___int64 {
	var tmp Hx_Typedef_haxe__int64___int64
	_ = tmp
	switch step {
	case 1:
		{
			tmp = Hx_Field_loreline__node_nodeid_impl__nextSection(this.Hx_Field_currentNodeId)
		}

	case 2:
		{
			tmp = Hx_Field_loreline__node_nodeid_impl__nextBranch(this.Hx_Field_currentNodeId)
		}

	case 3:
		{
			tmp = Hx_Field_loreline__node_nodeid_impl__nextBlock(this.Hx_Field_currentNodeId)
		}

	case 4:
		{
			tmp = Hx_Field_loreline__node_nodeid_impl__nextNode(this.Hx_Field_currentNodeId)
		}

	}

	this.Hx_Field_currentNodeId = tmp
	return this.Hx_Field_currentNodeId
}

func (this *Hx_Obj_loreline_parser) Hx_Field_peek() *Hx_Obj_loreline_token {
	var i int32
	_ = i
	i = (this.Hx_Field_current + 1)
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < this.Hx_Field_tokens.Len()) {
			break
		}

		switch this.Hx_Field_tokens.Get(i).Hx_Field_type.Hx_Field_enumIndex() {
		case 52:
			{
				i++
			}

		case 53:
			{
				i++
			}

		case 54:
			{
				i++
			}

		case 57:
			{
				i++
			}

		default:
			{
				return this.Hx_Field_tokens.Get(i)
			}
		}
	}

	var _hx_tmp_1 Hx_Typedef_loreline_tokens
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_tokens
	return _hx_tmp_1.Get((this.Hx_Field_tokens.Len() - 1))
}

func (this *Hx_Obj_loreline_parser) Hx_Field_advance(advanceLineBreaks__isOptional any) *Hx_Obj_loreline_token {
	advanceLineBreaks := HxDefault[bool]()
	_ = advanceLineBreaks

	if advanceLineBreaks__isOptional != nil {
		advanceLineBreaks = HxConvert[bool](advanceLineBreaks__isOptional)
	} else {
		advanceLineBreaks = true
	}
	{
		if this.Hx_Field_parsingCondition {
			advanceLineBreaks = false
		}

		var prev *Hx_Obj_loreline_token
		_ = prev
		prev = this.Hx_Field_tokens.Get(this.Hx_Field_current)
		if !this.VTable.Hx_Field_isAtEnd() {
			this.Hx_Field_lastTokenEnd = prev.Hx_Field_pos
			this.Hx_Field_lineBreakAfterToken = false
			for {
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if !this.VTable.Hx_Field_isAtEnd() {
					_hx_tmp_1 = (this.Hx_Field_tokens.Get((this.Hx_Field_current + 1)) != nil)
				} else {
					_hx_tmp_1 = false
				}

				if _hx_tmp_1 {
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					if this.VTable.Hx_Field_isComment(this.Hx_Field_tokens.Get((this.Hx_Field_current + 1)).Hx_Field_type) {
						_hx_tmp_2 = true
					} else {
						_hx_tmp_2 = (advanceLineBreaks && (this.Hx_Field_tokens.Get((this.Hx_Field_current + 1)).Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{}))
					}

					_hx_tmp_0 = (_hx_tmp_2)
				} else {
					_hx_tmp_0 = false
				}

				if !(_hx_tmp_0) {
					break
				}

				this.Hx_Field_current++
				{
					var _g Hx_Enum_loreline_tokentype
					_ = _g
					_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
					switch _g.Hx_Field_enumIndex() {
					case 52:
						{
							var content string
							_ = content
							content = _g.(Hx_Enum_loreline_tokentype_CommentLine).Hx_Field_content
							{
								if this.Hx_Field_pendingComments == nil {
									this.Hx_Field_pendingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
								}

								var _hx_tmp_3 HxArray[*Hx_Obj_loreline_comment]
								_ = _hx_tmp_3
								_hx_tmp_3 = this.Hx_Field_pendingComments
								_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_3, Hx_Obj_loreline_comment_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), content, false, nil)))
							}
						}

					case 53:
						{
							var content string
							_ = content
							content = _g.(Hx_Enum_loreline_tokentype_CommentMultiLine).Hx_Field_content
							{
								if this.Hx_Field_pendingComments == nil {
									this.Hx_Field_pendingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
								}

								var _hx_tmp_4 HxArray[*Hx_Obj_loreline_comment]
								_ = _hx_tmp_4
								_hx_tmp_4 = this.Hx_Field_pendingComments
								_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_4, Hx_Obj_loreline_comment_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), content, true, nil)))
							}
						}

					case 54:
						{
							var content string
							_ = content
							content = _g.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
							{
								if this.Hx_Field_pendingComments == nil {
									this.Hx_Field_pendingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
								}

								var _hx_tmp_5 HxArray[*Hx_Obj_loreline_comment]
								_ = _hx_tmp_5
								_hx_tmp_5 = this.Hx_Field_pendingComments
								_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_5, Hx_Obj_loreline_comment_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), content, false, true)))
							}
						}

					case 57:
						{
							this.Hx_Field_lastLineBreak = this.VTable.Hx_Field_currentPos()
							this.Hx_Field_lineBreakAfterToken = true
						}

					default:
						{
						}
					}
				}
			}

			this.Hx_Field_current++
		}

		return prev
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_previous() *Hx_Obj_loreline_token {
	return this.Hx_Field_tokens.Get((this.Hx_Field_current - 1))
}

func (this *Hx_Obj_loreline_parser) Hx_Field_prevWithType(_hx_reserved_type Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_token {
	var n int32
	_ = n
	n = (this.Hx_Field_current - 1)
	for n >= 0 {
		if Hx_Field_loreline_tokentypehelpers_equals(this.Hx_Field_tokens.Get(n).Hx_Field_type, _hx_reserved_type) {
			return this.Hx_Field_tokens.Get(n)
		}

		n--
	}

	return nil
}

func (this *Hx_Obj_loreline_parser) Hx_Field_prevIdentifier() *Hx_Obj_loreline_token {
	var n int32
	_ = n
	n = (this.Hx_Field_current - 1)
	for n >= 0 {
		if this.Hx_Field_tokens.Get(n).Hx_Field_type.Hx_Field_enumIndex() == 20 {
			return this.Hx_Field_tokens.Get(n)
		}

		n--
	}

	return nil
}

func (this *Hx_Obj_loreline_parser) Hx_Field_prevNonWhitespaceOrComment() *Hx_Obj_loreline_token {
	var n int32
	_ = n
	n = (this.Hx_Field_current - 1)
	for n >= 0 {
		switch this.Hx_Field_tokens.Get(n).Hx_Field_type.Hx_Field_enumIndex() {
		case 52:
			{
			}

		case 53:
			{
			}

		case 54:
			{
			}

		case 55, 56, 57:
			{
			}

		default:
			{
				return this.Hx_Field_tokens.Get(n)
			}
		}

		n--
	}

	return nil
}

func (this *Hx_Obj_loreline_parser) Hx_Field_nextNonWhitespaceOrComment() *Hx_Obj_loreline_token {
	var n int32
	_ = n
	n = this.Hx_Field_current
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = n
		if !(_hx_tmp_0 < this.Hx_Field_tokens.Len()) {
			break
		}

		switch this.Hx_Field_tokens.Get(n).Hx_Field_type.Hx_Field_enumIndex() {
		case 52:
			{
			}

		case 53:
			{
			}

		case 54:
			{
			}

		case 55, 56, 57:
			{
			}

		default:
			{
				return this.Hx_Field_tokens.Get(n)
			}
		}

		n++
	}

	return nil
}

func (this *Hx_Obj_loreline_parser) Hx_Field_nextNonLineBreakOrComment() *Hx_Obj_loreline_token {
	var n int32
	_ = n
	n = this.Hx_Field_current
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = n
		if !(_hx_tmp_0 < this.Hx_Field_tokens.Len()) {
			break
		}

		switch this.Hx_Field_tokens.Get(n).Hx_Field_type.Hx_Field_enumIndex() {
		case 52:
			{
			}

		case 53:
			{
			}

		case 54:
			{
			}

		case 57:
			{
			}

		default:
			{
				return this.Hx_Field_tokens.Get(n)
			}
		}

		n++
	}

	return nil
}

func (this *Hx_Obj_loreline_parser) Hx_Field_currentPos() *Hx_Obj_loreline_position {
	var tmp *Hx_Obj_loreline_token
	_ = tmp
	tmp = this.Hx_Field_tokens.Get(this.Hx_Field_current)
	var _hx_tmp_0 *Hx_Obj_loreline_position
	_ = _hx_tmp_0
	if tmp != nil {
		_hx_tmp_0 = tmp.Hx_Field_pos
	} else {
		_hx_tmp_0 = nil
	}

	var tmp1 *Hx_Obj_loreline_position
	_ = tmp1
	tmp1 = _hx_tmp_0
	if tmp1 != nil {
		return tmp1
	} else {
		return Hx_Obj_loreline_position_CreateInstance(1, 1, 0, 0)
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_currentTokenTypeToString() string {
	if (this.Hx_Field_tokens != nil) && (this.Hx_Field_current >= 0) {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokens.Get(this.Hx_Field_current)
		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		if token != nil {
			_hx_tmp_0 = token.Hx_Field_type
		} else {
			_hx_tmp_0 = nil
		}

		var tmp Hx_Enum_loreline_tokentype
		_ = tmp
		tmp = _hx_tmp_0
		if tmp != nil {
			return Hx_Field_loreline_tokentypehelpers_toCodeString(tmp)
		} else {
			return (HxNullable[string]{Value: HxStringNull}).Value
		}
	}

	return HxStringNull
}

func (this *Hx_Obj_loreline_parser) Hx_Field_check(_hx_reserved_type Hx_Enum_loreline_tokentype) bool {
	if this.VTable.Hx_Field_isAtEnd() {
		return false
	}

	if _hx_reserved_type.Hx_Field_enumIndex() == 42 {
		return (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{})
	} else {
		return Hx_Field_loreline_tokentypehelpers_equals(this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type, _hx_reserved_type)
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_checkString() bool {
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	if _g.Hx_Field_enumIndex() == 16 {
		var s int32
		_ = s
		s = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
		var attachments string
		_ = attachments
		attachments = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
		return true
	} else {
		return false
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_isAtEnd() bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_current
	if _hx_tmp_1 >= this.Hx_Field_tokens.Len() {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type == Hx_Enum_loreline_tokentype_Eof{})
	}

	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_parser) Hx_Field_isComment(_hx_reserved_type Hx_Enum_loreline_tokentype) bool {
	switch _hx_reserved_type.Hx_Field_enumIndex() {
	case 52:
		{
			return true
		}

	case 53:
		{
			return true
		}

	case 54:
		{
			return true
		}

	default:
		{
			return false
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseNode(topLevel__isOptional any) *Hx_Obj_loreline_astnode {
	topLevel := HxDefault[bool]()
	_ = topLevel

	if topLevel__isOptional != nil {
		topLevel = HxConvert[bool](topLevel__isOptional)
	} else {
		topLevel = false
	}
	{
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if this.VTable.Hx_Field_isComment(this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type) {
				_hx_tmp_0 = true
			} else {
				_hx_tmp_0 = (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{})
			}

			if !(_hx_tmp_0) {
				break
			}

			if this.VTable.Hx_Field_isComment(this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type) {
				if this.Hx_Field_pendingComments == nil {
					this.Hx_Field_pendingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
				}

				var tmp HxArray[*Hx_Obj_loreline_comment]
				_ = tmp
				tmp = this.Hx_Field_pendingComments
				var tmp1 Hx_Typedef_haxe__int64___int64
				_ = tmp1
				tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
				var tmp2 *Hx_Obj_loreline_position
				_ = tmp2
				tmp2 = this.VTable.Hx_Field_currentPos()
				var _g Hx_Enum_loreline_tokentype
				_ = _g
				_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
				var tmp3 string
				_ = tmp3
				switch _g.Hx_Field_enumIndex() {
				case 52:
					{
						var content string
						_ = content
						content = _g.(Hx_Enum_loreline_tokentype_CommentLine).Hx_Field_content
						tmp3 = content
					}

				case 53:
					{
						var content string
						_ = content
						content = _g.(Hx_Enum_loreline_tokentype_CommentMultiLine).Hx_Field_content
						tmp3 = content
					}

				case 54:
					{
						var content string
						_ = content
						content = _g.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
						tmp3 = content
					}

				default:
					{
						tmp3 = ""
					}
				}

				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_comment]
				_ = _hx_tmp_1
				_hx_tmp_1 = tmp
				_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_1, Hx_Obj_loreline_comment_CreateInstance(tmp1, tmp2, tmp3, (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() == 53), (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() == 54))))
			}

			this.VTable.Hx_Field_advance(nil)
			if this.VTable.Hx_Field_isAtEnd() {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Unexpected end of file", this.VTable.Hx_Field_currentPos())))))
			}
		}

		if this.VTable.Hx_Field_isAtEnd() {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Unexpected end of file", this.VTable.Hx_Field_currentPos())))))
		}

		var _gthis *Hx_Obj_loreline_parser
		_ = _gthis
		_gthis = this
		switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
		case 0:
			{
				return (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseImport()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
			}

		case 1:
			{
				return (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseStateDecl(false)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
			}

		case 2:
			{
				if (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_LParen{}) {
					var node *Hx_Obj_loreline_astnode
					_ = node
					node = (func() *Hx_Obj_loreline_astnode {
						_hx_up := this.VTable.Hx_Field_parseBeatCall()
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					if topLevel {
						return (func() *Hx_Obj_loreline_astnode {
							_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_astnode
							}
							return nil
						})()
					} else {
						return node
					}
				} else {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := this.VTable.Hx_Field_parseBeatDecl()
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				}
			}

		case 3:
			{
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := this.VTable.Hx_Field_parseCharacterDecl()
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unexpected: ", this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos()))
					this.VTable.Hx_Field_advance(nil)
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				}
			}

		case 4:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseChoiceStatement()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 5:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseWhenStatement()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 6:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseIfStatement()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 8:
			{
				this.VTable.Hx_Field_advance(nil)
				if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_KwState{})) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected 'state' after 'new'", this.VTable.Hx_Field_currentPos())))))
				}

				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseStateDecl(true)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 9, 10, 11, 12, 13:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseAlternative()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 15:
			{
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := this.VTable.Hx_Field_parseFunction()
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unexpected: ", this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos()))
					this.VTable.Hx_Field_advance(nil)
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				}
			}

		case 16:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseTextStatement()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 17:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = this.VTable.Hx_Field_parseExpressionStatement()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 18:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = this.VTable.Hx_Field_parseExpressionStatement()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 20:
			{
				if (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) {
					var node *Hx_Obj_loreline_astnode
					_ = node
					node = (func() *Hx_Obj_loreline_astnode {
						_hx_up := this.VTable.Hx_Field_parseDialogueStatement()
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					if topLevel {
						return (func() *Hx_Obj_loreline_astnode {
							_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_astnode
							}
							return nil
						})()
					} else {
						return node
					}
				} else {
					var node *Hx_Obj_loreline_astnode
					_ = node
					node = this.VTable.Hx_Field_parseExpressionStatement()
					if topLevel {
						return (func() *Hx_Obj_loreline_astnode {
							_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_astnode
							}
							return nil
						})()
					} else {
						return node
					}
				}
			}

		case 27:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseInsertion()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 19, 28, 46, 48, 50:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = this.VTable.Hx_Field_parseExpressionStatement()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 40:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = this.VTable.Hx_Field_parseExpressionStatement()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		case 41:
			{
				var node *Hx_Obj_loreline_astnode
				_ = node
				node = (func() *Hx_Obj_loreline_astnode {
					_hx_up := this.VTable.Hx_Field_parseTransition()
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				if topLevel {
					return (func() *Hx_Obj_loreline_astnode {
						_hx_up := _gthis.VTable.Hx_Field_wrapInRootBeat(node)
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
				} else {
					return node
				}
			}

		default:
			{
				this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unexpected: ", this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos()))
				this.VTable.Hx_Field_advance(nil)
				return (func() *Hx_Obj_loreline_astnode {
					_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
			}
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_wrapInRootBeat(node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_nbeatdecl {
	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	var result *Hx_Obj_loreline_nbeatdecl
	_ = result
	if this.Hx_Field_rootBeat == nil {
		var startPos *Hx_Obj_loreline_position
		_ = startPos
		startPos = this.VTable.Hx_Field_currentPos()
		body = HxMakeArray[*Hx_Obj_loreline_astnode]()
		this.Hx_Field_rootBeat = Hx_Obj_loreline_nbeatdecl_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(1))), startPos, "_", body, nil, nil)
		result = this.Hx_Field_rootBeat
	} else {
		body = this.Hx_Field_rootBeat.Hx_Field_body
	}

	_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](body, node))
	return result
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseImport() *Hx_Obj_loreline_nimportstatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwImport{})
	if this.Hx_Field_context == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Cannot import without a context", this.VTable.Hx_Field_currentPos())))))
	}

	var quotes int32
	_ = quotes
	quotes = ((int32)(0))
	var pathToken *Hx_Obj_loreline_token
	_ = pathToken
	pathToken = this.Hx_Field_tokens.Get(this.Hx_Field_current)
	var rawImportPath string
	_ = rawImportPath
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = pathToken.Hx_Field_type
	if _g.Hx_Field_enumIndex() == 16 {
		var q int32
		_ = q
		q = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
		var s string
		_ = s
		s = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
		quotes = q
		rawImportPath = s
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected string literal for import path", this.VTable.Hx_Field_currentPos())))))
	}

	var importPath string
	_ = importPath
	importPath = rawImportPath
	if !Hx_Field_haxe_io_path_isAbsolute(importPath) {
		importPath = Hx_Field_haxe_io_path_join(HxMakeArray[string](Hx_Field_haxe_io_path_directory(this.Hx_Field_context.Hx_Field_path), importPath))
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
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 7)))
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
				tmp6 = ((c1 == 108) || (c1 == 76))
			} else {
				tmp6 = false
			}

			if tmp6 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 6)))
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
				tmp5 = ((c1 == 111) || (c1 == 79))
			} else {
				tmp5 = false
			}

			if tmp5 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 5)))
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
				tmp4 = ((c1 == 114) || (c1 == 82))
			} else {
				tmp4 = false
			}

			if tmp4 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 4)))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (tmp_c_1).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (tmp_c_1).Value
				}

				tmp3 = ((_hx_tmp_4) == 46)
			} else {
				tmp3 = false
			}

			if tmp3 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
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
				tmp2 = ((c1 == 116) || (c1 == 84))
			} else {
				tmp2 = false
			}

			if tmp2 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
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
				tmp_tmp_1 = ((c1 == 120) || (c1 == 88))
			} else {
				tmp_tmp_1 = false
			}

			if tmp_tmp_1 {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
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
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				if (c).Valid == false {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = (c).Value
				}

				if (_hx_tmp_8) == 46 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 3)))
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
					tmp2 = ((c1 == 108) || (c1 == 76))
				} else {
					tmp2 = false
				}

				if tmp2 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 2)))
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					if (tmp_c_1).Valid == false {
						_hx_tmp_10 = -1
					} else {
						_hx_tmp_10 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_10
					tmp_tmp1_1 = ((c1 == 111) || (c1 == 79))
				} else {
					tmp_tmp1_1 = false
				}

				if tmp_tmp1_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(importPath, (n - 1)))
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
		var path string
		_ = path
		path = this.Hx_Field_context.Hx_Field_rootPath
		var importPath1 bool
		_ = importPath1
		if path == HxStringNull {
			importPath1 = false
		} else {
			var n int32
			_ = n
			n = int32(utf8.RuneCountInString(path))
			if n < 8 {
				importPath1 = false
			} else {
				var tmp_importPath_1 bool
				_ = tmp_importPath_1
				var importPath2 bool
				_ = importPath2
				var importPath3 bool
				_ = importPath3
				var importPath4 bool
				_ = importPath4
				var importPath5 bool
				_ = importPath5
				var importPath6 bool
				_ = importPath6
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 8)))
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
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 7)))
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
					importPath6 = ((c1 == 108) || (c1 == 76))
				} else {
					importPath6 = false
				}

				if importPath6 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 6)))
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
					importPath5 = ((c1 == 111) || (c1 == 79))
				} else {
					importPath5 = false
				}

				if importPath5 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 5)))
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
					importPath4 = ((c1 == 114) || (c1 == 82))
				} else {
					importPath4 = false
				}

				if importPath4 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 4)))
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					if (tmp_c_1).Valid == false {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = (tmp_c_1).Value
					}

					importPath3 = ((_hx_tmp_16) == 46)
				} else {
					importPath3 = false
				}

				if importPath3 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 3)))
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
					importPath2 = ((c1 == 116) || (c1 == 84))
				} else {
					importPath2 = false
				}

				if importPath2 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 2)))
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					if (tmp_c_1).Valid == false {
						_hx_tmp_18 = -1
					} else {
						_hx_tmp_18 = (tmp_c_1).Value
					}

					var c1 int32
					_ = c1
					c1 = _hx_tmp_18
					tmp_importPath_1 = ((c1 == 120) || (c1 == 88))
				} else {
					tmp_importPath_1 = false
				}

				if tmp_importPath_1 {
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(path, (n - 1)))
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
					importPath1 = ((c1 == 116) || (c1 == 84))
				} else {
					importPath1 = false
				}
			}
		}

		var _hx_tmp_20 string
		_ = _hx_tmp_20
		_hx_tmp_20 = importPath
		var _hx_tmp_21 string
		_ = _hx_tmp_21
		if importPath1 {
			_hx_tmp_21 = ".lor.txt"
		} else {
			_hx_tmp_21 = ".lor"
		}

		importPath = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_20, _hx_tmp_21)
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_692 func(_hx_param_key any) bool
		_ = _hx_closure_src_692
		_hx_closure_src_692 = this.Hx_Field_context.Hx_Field_imported.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_692(((any)(_hx_param_key)))
		}
	}()(importPath) {
		this.VTable.Hx_Field_advance(nil)
		var node *Hx_Obj_loreline_nimportstatement
		_ = node
		node = Hx_Obj_loreline_nimportstatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(1))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), nil, nil, nil, nil)
		node.Hx_Field_path = Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pathToken.Hx_Field_pos, quotes, HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pathToken.Hx_Field_pos, Hx_Enum_loreline_stringparttype_Raw{rawImportPath}, nil, nil)), nil, nil)
		return func() func(node *Hx_Obj_loreline_nimportstatement) *Hx_Obj_loreline_nimportstatement {
			var _hx_closure_src_693 func(node any) any
			_ = _hx_closure_src_693
			_hx_closure_src_693 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nimportstatement) *Hx_Obj_loreline_nimportstatement {
				var hx_dyncast_294 any
				_ = hx_dyncast_294
				hx_dyncast_294 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_693(((any)(node))), "Hx_Obj_loreline_nimportstatement")
				var _hx_tmp_22 *Hx_Obj_loreline_nimportstatement
				_ = _hx_tmp_22
				if hx_dyncast_294 != nil {
					_hx_tmp_22 = hx_dyncast_294.(interface {
						Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
					}).Hx_As_Hx_Obj_loreline_nimportstatement()
				} else {
					_hx_tmp_22 = nil
				}

				return _hx_tmp_22
			}
		}()(node)
	}

	var importedTokens Hx_Typedef_loreline_tokens
	_ = importedTokens
	importedTokens = func() func(_hx_param_key string) Hx_Typedef_loreline_tokens {
		var _hx_closure_src_694 func(_hx_param_key any) any
		_ = _hx_closure_src_694
		_hx_closure_src_694 = this.Hx_Field_context.Hx_Field_imports.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) Hx_Typedef_loreline_tokens {
			return HxAnyToTypedArray[*Hx_Obj_loreline_token](_hx_closure_src_694(((any)(_hx_param_key))))
		}
	}()(importPath)
	if importedTokens == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Failed to import file at path ", importPath), " ("), rawImportPath), ")"), this.VTable.Hx_Field_currentPos())))))
	}

	func() func(_hx_param_key string, _hx_param_value bool) {
		var _hx_closure_src_695 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_695
		_hx_closure_src_695 = this.Hx_Field_context.Hx_Field_imported.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value bool) {
			_hx_closure_src_695(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(importPath, true)
	var tempParser *Hx_Obj_loreline_parser
	_ = tempParser
	tempParser = Hx_Obj_loreline_parser_CreateInstance(importedTokens, Hx_Obj_loreline_parsercontext_CreateInstance(this.Hx_Field_context.Hx_Field_rootPath, importPath, this.Hx_Field_context.Hx_Field_imports, this.Hx_Field_context.Hx_Field_imported))
	tempParser.Hx_Field_currentNodeId = this.Hx_Field_currentNodeId
	var importedScript *Hx_Obj_loreline_script
	_ = importedScript
	importedScript = tempParser.VTable.Hx_Field_parse()
	this.Hx_Field_currentNodeId = tempParser.Hx_Field_currentNodeId
	this.VTable.Hx_Field_advance(nil)
	var node *Hx_Obj_loreline_nimportstatement
	_ = node
	node = Hx_Obj_loreline_nimportstatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(1))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), nil, importedScript, nil, nil)
	node.Hx_Field_path = Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pathToken.Hx_Field_pos, quotes, HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pathToken.Hx_Field_pos, Hx_Enum_loreline_stringparttype_Raw{rawImportPath}, nil, nil)), nil, nil)
	return func() func(node *Hx_Obj_loreline_nimportstatement) *Hx_Obj_loreline_nimportstatement {
		var _hx_closure_src_696 func(node any) any
		_ = _hx_closure_src_696
		_hx_closure_src_696 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nimportstatement) *Hx_Obj_loreline_nimportstatement {
			var hx_dyncast_295 any
			_ = hx_dyncast_295
			hx_dyncast_295 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_696(((any)(node))), "Hx_Obj_loreline_nimportstatement")
			var _hx_tmp_23 *Hx_Obj_loreline_nimportstatement
			_ = _hx_tmp_23
			if hx_dyncast_295 != nil {
				_hx_tmp_23 = hx_dyncast_295.(interface {
					Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
				}).Hx_As_Hx_Obj_loreline_nimportstatement()
			} else {
				_hx_tmp_23 = nil
			}

			return _hx_tmp_23
		}
	}()(node)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseDialogueStatement() *Hx_Obj_loreline_ndialoguestatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var dialogue *Hx_Obj_loreline_ndialoguestatement
	_ = dialogue
	dialogue = Hx_Obj_loreline_ndialoguestatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, HxStringNull, nil, nil, nil, nil)
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	var tmp string
	_ = tmp
	if _g.Hx_Field_enumIndex() == 20 {
		var name string
		_ = name
		name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
		tmp = name
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected character name", this.VTable.Hx_Field_currentPos())))))
	}

	dialogue.Hx_Field_character = tmp
	dialogue.Hx_Field_characterPos = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_pos
	this.VTable.Hx_Field_advance(nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_Colon{})
	func() func(node *Hx_Obj_loreline_ndialoguestatement) *Hx_Obj_loreline_ndialoguestatement {
		var _hx_closure_src_697 func(node any) any
		_ = _hx_closure_src_697
		_hx_closure_src_697 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_ndialoguestatement) *Hx_Obj_loreline_ndialoguestatement {
			var hx_dyncast_296 any
			_ = hx_dyncast_296
			hx_dyncast_296 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_697(((any)(node))), "Hx_Obj_loreline_ndialoguestatement")
			var _hx_tmp_0 *Hx_Obj_loreline_ndialoguestatement
			_ = _hx_tmp_0
			if hx_dyncast_296 != nil {
				_hx_tmp_0 = hx_dyncast_296.(interface {
					Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
				}).Hx_As_Hx_Obj_loreline_ndialoguestatement()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(dialogue)
	var indented bool
	_ = indented
	indented = false
	if this.VTable.Hx_Field_checkBlockStart() {
		var blockToken *Hx_Obj_loreline_token
		_ = blockToken
		blockToken = this.VTable.Hx_Field_parseBlockStart()
		if (blockToken.Hx_Field_type != Hx_Enum_loreline_tokentype_Indent{}) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Expected indent, got ", Hx_Field_loreline_tokentypehelpers_toCodeString(blockToken.Hx_Field_type)), blockToken.Hx_Field_pos)))))
		}

		indented = true
	}

	{
		var hx_try_return *Hx_Obj_loreline_ndialoguestatement
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
						_ = e
						{
							this.VTable.Hx_Field_addError(e)
							dialogue.Hx_Field_content = Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), ((int32)(0)), HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), Hx_Enum_loreline_stringparttype_Raw{"?"}, nil, nil)), nil, nil)
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				dialogue.Hx_Field_content = this.VTable.Hx_Field_parseStringLiteral()
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if !this.Hx_Field_lineBreakAfterToken {
		var _hx_tmp_2 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Enum_loreline_tokentype_KwIf{}

		var _hx_tmp_3 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = this.VTable.Hx_Field_match(_hx_tmp_3, nil)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var ifPos *Hx_Obj_loreline_position
		_ = ifPos
		ifPos = this.VTable.Hx_Field_previous().Hx_Field_pos
		dialogue.Hx_Field_condition = this.VTable.Hx_Field_parseConditionExpression()
		dialogue.Hx_Field_conditionStyle = this.VTable.Hx_Field_takeConditionParens(dialogue.Hx_Field_condition)
		dialogue.Hx_Field_conditionPos = ifPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	}

	if indented {
		for {
			var _hx_tmp_4 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Enum_loreline_tokentype_LineBreak{}

			if !(this.VTable.Hx_Field_match(_hx_tmp_4, nil)) {
				break
			}
		}

		this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_Unindent{})
	}

	dialogue.Hx_Field_pos = dialogue.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return dialogue
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseStatementBlock(statements HxArray[*Hx_Obj_loreline_astnode]) int32 {
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	for {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_1 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_1 = false
		}

		if !(_hx_tmp_1) {
			break
		}

		for {
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_3, nil) {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_5 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_5
					_hx_tmp_5 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_6 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_5
					_hx_tmp_4 = this.VTable.Hx_Field_match(_hx_tmp_6, nil)
				} else {
					_hx_tmp_4 = false
				}

				_hx_tmp_2 = _hx_tmp_4
			}

			if !(_hx_tmp_2) {
				break
			}
		}

		if this.VTable.Hx_Field_check(blockEnd) {
			break
		}

		{
			var hx_try_return int32
			_ = hx_try_return
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
							_ = e
							{
								if this.Hx_Field_errors == nil {
									this.Hx_Field_errors = HxMakeArray[*Hx_Obj_loreline_parseerror]()
								}

								_ = (Hx_Array_Push[*Hx_Obj_loreline_parseerror](this.Hx_Field_errors, e))
								this.VTable.Hx_Field_synchronize()
								if this.VTable.Hx_Field_check(blockEnd) {
									hx_try_state = 2
									return
								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					var _hx_tmp_7 HxArray[*Hx_Obj_loreline_astnode]
					_ = _hx_tmp_7
					_hx_tmp_7 = statements
					_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_7, this.VTable.Hx_Field_parseNode(nil)))
				}
			}()
			if hx_try_state == 1 {
				return hx_try_return
			}
			if hx_try_state == 2 {
				break
			}
			if hx_try_state == 3 {
				continue
			}
		}

		for {
			var _hx_tmp_8 bool
			_ = _hx_tmp_8
			var _hx_tmp_9 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_9
			_hx_tmp_9 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_9, nil) {
				_hx_tmp_8 = true
			} else {
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_11 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_11
					_hx_tmp_11 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_12 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_12
					_hx_tmp_12 = _hx_tmp_11
					_hx_tmp_10 = this.VTable.Hx_Field_match(_hx_tmp_12, nil)
				} else {
					_hx_tmp_10 = false
				}

				_hx_tmp_8 = _hx_tmp_10
			}

			if !(_hx_tmp_8) {
				break
			}
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		return ((int32)(1))
	} else {
		return ((int32)(0))
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseStateDecl(temporary bool) *Hx_Obj_loreline_nstatedecl {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var stateNode *Hx_Obj_loreline_nstatedecl
	_ = stateNode
	stateNode = Hx_Obj_loreline_nstatedecl_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, temporary, HxMakeArray[*Hx_Obj_loreline_nobjectfield](), nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwState{})
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_1 = ((int32)(1))
	} else {
		_hx_tmp_1 = ((int32)(0))
	}

	stateNode.Hx_Field_style = _hx_tmp_1
	func() func(node *Hx_Obj_loreline_nstatedecl) *Hx_Obj_loreline_nstatedecl {
		var _hx_closure_src_698 func(node any) any
		_ = _hx_closure_src_698
		_hx_closure_src_698 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nstatedecl) *Hx_Obj_loreline_nstatedecl {
			var hx_dyncast_297 any
			_ = hx_dyncast_297
			hx_dyncast_297 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_698(((any)(node))), "Hx_Obj_loreline_nstatedecl")
			var _hx_tmp_2 *Hx_Obj_loreline_nstatedecl
			_ = _hx_tmp_2
			if hx_dyncast_297 != nil {
				_hx_tmp_2 = hx_dyncast_297.(interface {
					Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
				}).Hx_As_Hx_Obj_loreline_nstatedecl()
			} else {
				_hx_tmp_2 = nil
			}

			return _hx_tmp_2
		}
	}()(stateNode)
	for {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_3 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_3 = false
		}

		if !(_hx_tmp_3) {
			break
		}

		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_7 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_8 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = this.VTable.Hx_Field_match(_hx_tmp_8, nil)
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_4 = _hx_tmp_6
			}

			if !(_hx_tmp_4) {
				break
			}
		}

		var _hx_tmp_9 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _hx_tmp_9
		_hx_tmp_9 = stateNode.Hx_Field_fields
		_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_9, this.VTable.Hx_Field_parseObjectField()))
		for {
			var _hx_tmp_10 bool
			_ = _hx_tmp_10
			var _hx_tmp_11 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_11
			_hx_tmp_11 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_11, nil) {
				_hx_tmp_10 = true
			} else {
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_13 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_13
					_hx_tmp_13 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = _hx_tmp_13
					_hx_tmp_12 = this.VTable.Hx_Field_match(_hx_tmp_14, nil)
				} else {
					_hx_tmp_12 = false
				}

				_hx_tmp_10 = _hx_tmp_12
			}

			if !(_hx_tmp_10) {
				break
			}
		}
	}

	for {
		var _hx_tmp_15 bool
		_ = _hx_tmp_15
		var _hx_tmp_16 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_16
		_hx_tmp_16 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_16, nil) {
			_hx_tmp_15 = true
		} else {
			var _hx_tmp_17 bool
			_ = _hx_tmp_17
			if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
				var _hx_tmp_18 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_18
				_hx_tmp_18 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_19 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_19
				_hx_tmp_19 = _hx_tmp_18
				_hx_tmp_17 = this.VTable.Hx_Field_match(_hx_tmp_19, nil)
			} else {
				_hx_tmp_17 = false
			}

			_hx_tmp_15 = _hx_tmp_17
		}

		if !(_hx_tmp_15) {
			break
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	stateNode.Hx_Field_pos = stateNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return stateNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseObjectField() *Hx_Obj_loreline_nobjectfield {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var name string
	_ = name
	name = this.VTable.Hx_Field_expectIdentifier()
	var objectField *Hx_Obj_loreline_nobjectfield
	_ = objectField
	objectField = Hx_Obj_loreline_nobjectfield_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, name, nil, nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_Colon{})
	func() func(node *Hx_Obj_loreline_nobjectfield) *Hx_Obj_loreline_nobjectfield {
		var _hx_closure_src_699 func(node any) any
		_ = _hx_closure_src_699
		_hx_closure_src_699 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nobjectfield) *Hx_Obj_loreline_nobjectfield {
			var hx_dyncast_298 any
			_ = hx_dyncast_298
			hx_dyncast_298 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_699(((any)(node))), "Hx_Obj_loreline_nobjectfield")
			var _hx_tmp_0 *Hx_Obj_loreline_nobjectfield
			_ = _hx_tmp_0
			if hx_dyncast_298 != nil {
				_hx_tmp_0 = hx_dyncast_298.(interface {
					Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
				}).Hx_As_Hx_Obj_loreline_nobjectfield()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(objectField)
	if this.VTable.Hx_Field_checkBlockStart() {
		objectField.Hx_Field_value = this.VTable.Hx_Field_parseObjectLiteral()
	} else {
		objectField.Hx_Field_value = this.VTable.Hx_Field_parseExpression()
	}

	objectField.Hx_Field_pos = objectField.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return objectField
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseBeatDecl() *Hx_Obj_loreline_nbeatdecl {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var beatNode *Hx_Obj_loreline_nbeatdecl
	_ = beatNode
	beatNode = Hx_Obj_loreline_nbeatdecl_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(1))), startPos, HxStringNull, HxMakeArray[*Hx_Obj_loreline_astnode](), HxMakeArray[*Hx_Obj_loreline_comment](), nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwBeat{})
	beatNode.Hx_Field_pos = startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_currentPos())
	beatNode.Hx_Field_name = this.VTable.Hx_Field_expectIdentifier()
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Enum_loreline_tokentype_LParen{}

	if this.VTable.Hx_Field_match(_hx_tmp_0, nil) {
		var params HxArray[*Hx_Obj_loreline_nbeatparam]
		_ = params
		params = HxMakeArray[*Hx_Obj_loreline_nbeatparam]()
		if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_RParen{})) {
			for {
				for {
					var _hx_tmp_1 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Enum_loreline_tokentype_LineBreak{}

					if !(this.VTable.Hx_Field_match(_hx_tmp_1, nil)) {
						break
					}
				}

				var paramPos *Hx_Obj_loreline_position
				_ = paramPos
				paramPos = this.VTable.Hx_Field_currentPos()
				var paramName string
				_ = paramName
				paramName = this.VTable.Hx_Field_expectIdentifier()
				var defaultValue *Hx_Obj_loreline_nexpr
				_ = defaultValue
				var _hx_tmp_2 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpAssign{}

				if this.VTable.Hx_Field_match(_hx_tmp_2, nil) {
					defaultValue = this.VTable.Hx_Field_parseExpression()
				}

				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						if !(_hx_tmp_3 < params.Len()) {
							break
						}

						var existing *Hx_Obj_loreline_nbeatparam
						_ = existing
						existing = params.Get(_g)
						_g++
						if existing.Hx_Field_name == paramName {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Duplicate beat parameter: ", paramName), paramPos))
						}
					}
				}

				var _hx_tmp_4 HxArray[*Hx_Obj_loreline_nbeatparam]
				_ = _hx_tmp_4
				_hx_tmp_4 = params
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatparam](_hx_tmp_4, Hx_Obj_loreline_nbeatparam_CreateInstance(paramName, paramPos, defaultValue)))
				for {
					var _hx_tmp_5 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_5
					_hx_tmp_5 = Hx_Enum_loreline_tokentype_LineBreak{}

					if !(this.VTable.Hx_Field_match(_hx_tmp_5, nil)) {
						break
					}
				}

				var _hx_tmp_6 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_6
				_hx_tmp_6 = Hx_Enum_loreline_tokentype_Comma{}

				if !(this.VTable.Hx_Field_match(_hx_tmp_6, nil)) {
					break
				}
			}
		}

		this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_RParen{})
		beatNode.Hx_Field_params = params
	}

	var blockStart *Hx_Obj_loreline_token
	_ = blockStart
	blockStart = this.VTable.Hx_Field_parseBlockStart()
	var _hx_tmp_7 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_7
	if (blockStart.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_7 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_7 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_7
	var _hx_tmp_8 int32
	_ = _hx_tmp_8
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_8 = ((int32)(1))
	} else {
		_hx_tmp_8 = ((int32)(0))
	}

	beatNode.Hx_Field_style = _hx_tmp_8
	func() func(node *Hx_Obj_loreline_nbeatdecl) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_700 func(node any) any
		_ = _hx_closure_src_700
		_hx_closure_src_700 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nbeatdecl) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_299 any
			_ = hx_dyncast_299
			hx_dyncast_299 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_700(((any)(node))), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_9 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_9
			if hx_dyncast_299 != nil {
				_hx_tmp_9 = hx_dyncast_299.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_9 = nil
			}

			return _hx_tmp_9
		}
	}()(beatNode)
	for {
		var _hx_tmp_10 bool
		_ = _hx_tmp_10
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_10 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_10 = false
		}

		if !(_hx_tmp_10) {
			break
		}

		for {
			var _hx_tmp_11 bool
			_ = _hx_tmp_11
			var _hx_tmp_12 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_12
			_hx_tmp_12 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_12, nil) {
				_hx_tmp_11 = true
			} else {
				var _hx_tmp_13 bool
				_ = _hx_tmp_13
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_15 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_15
					_hx_tmp_15 = _hx_tmp_14
					_hx_tmp_13 = this.VTable.Hx_Field_match(_hx_tmp_15, nil)
				} else {
					_hx_tmp_13 = false
				}

				_hx_tmp_11 = _hx_tmp_13
			}

			if !(_hx_tmp_11) {
				break
			}
		}

		var _hx_tmp_16 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_16
		_hx_tmp_16 = beatNode.Hx_Field_body
		_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_16, this.VTable.Hx_Field_parseNode(nil)))
		for {
			var _hx_tmp_17 bool
			_ = _hx_tmp_17
			var _hx_tmp_18 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_18
			_hx_tmp_18 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_18, nil) {
				_hx_tmp_17 = true
			} else {
				var _hx_tmp_19 bool
				_ = _hx_tmp_19
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_20 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_20
					_hx_tmp_20 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_21 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_21
					_hx_tmp_21 = _hx_tmp_20
					_hx_tmp_19 = this.VTable.Hx_Field_match(_hx_tmp_21, nil)
				} else {
					_hx_tmp_19 = false
				}

				_hx_tmp_17 = _hx_tmp_19
			}

			if !(_hx_tmp_17) {
				break
			}
		}
	}

	for {
		var _hx_tmp_22 bool
		_ = _hx_tmp_22
		var _hx_tmp_23 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_23
		_hx_tmp_23 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_23, nil) {
			_hx_tmp_22 = true
		} else {
			var _hx_tmp_24 bool
			_ = _hx_tmp_24
			if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
				var _hx_tmp_25 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_25
				_hx_tmp_25 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_26 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_26
				_hx_tmp_26 = _hx_tmp_25
				_hx_tmp_24 = this.VTable.Hx_Field_match(_hx_tmp_26, nil)
			} else {
				_hx_tmp_24 = false
			}

			_hx_tmp_22 = _hx_tmp_24
		}

		if !(_hx_tmp_22) {
			break
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	beatNode.Hx_Field_pos = beatNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return beatNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_checkBlockStart() bool {
	var indentToken *Hx_Obj_loreline_token
	_ = indentToken
	var braceToken *Hx_Obj_loreline_token
	_ = braceToken
	var numIndents int32
	_ = numIndents
	numIndents = 0
	var i int32
	_ = i
	i = 0
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = (this.Hx_Field_current + i)
		if !(_hx_tmp_0 < this.Hx_Field_tokens.Len()) {
			break
		}

		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokens.Get((this.Hx_Field_current + i))
		i++
		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{}) {
			continue
		}

		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
			numIndents++
			indentToken = token
			continue
		}

		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LBrace{}) {
			if braceToken == nil {
				braceToken = token
			}

			continue
		}

		break
	}

	if braceToken != nil {
		return true
	} else {
		if indentToken != nil {
			if numIndents > 1 {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Invalid indentation level", indentToken.Hx_Field_pos)))))
			}

			return true
		} else {
			return false
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseBlockStart() *Hx_Obj_loreline_token {
	var indentToken *Hx_Obj_loreline_token
	_ = indentToken
	var braceToken *Hx_Obj_loreline_token
	_ = braceToken
	var numIndents int32
	_ = numIndents
	numIndents = 0
	for {
		if !(!this.VTable.Hx_Field_isAtEnd()) {
			break
		}

		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_0, nil) {
			continue
		}

		var _hx_tmp_1 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Enum_loreline_tokentype_Indent{}

		if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
			numIndents++
			indentToken = this.VTable.Hx_Field_prevWithType(Hx_Enum_loreline_tokentype_Indent{})
			continue
		}

		var _hx_tmp_2 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Enum_loreline_tokentype_LBrace{}

		if this.VTable.Hx_Field_match(_hx_tmp_2, nil) {
			if braceToken == nil {
				braceToken = this.VTable.Hx_Field_prevWithType(Hx_Enum_loreline_tokentype_LBrace{})
			}

			continue
		}

		break
	}

	if braceToken != nil {
		return braceToken
	} else {
		if indentToken != nil {
			if numIndents > 1 {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Invalid indentation level", indentToken.Hx_Field_pos)))))
			}

			return indentToken
		} else {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Expected ", Hx_Field_loreline_tokentypehelpers_toCodeString(Hx_Enum_loreline_tokentype_LBrace{})), " or ")
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_4, Hx_Field_loreline_tokentypehelpers_toCodeString(Hx_Enum_loreline_tokentype_Indent{})), ", got ")
			this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_3, this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos()))
			return Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Indent{}, this.VTable.Hx_Field_currentPos())
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseCharacterDecl() *Hx_Obj_loreline_ncharacterdecl {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var characterNode *Hx_Obj_loreline_ncharacterdecl
	_ = characterNode
	characterNode = Hx_Obj_loreline_ncharacterdecl_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, HxStringNull, nil, HxMakeArray[*Hx_Obj_loreline_nobjectfield](), nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwCharacter{})
	characterNode.Hx_Field_name = this.VTable.Hx_Field_expectIdentifier()
	characterNode.Hx_Field_namePos = this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_1 = ((int32)(1))
	} else {
		_hx_tmp_1 = ((int32)(0))
	}

	characterNode.Hx_Field_style = _hx_tmp_1
	func() func(node *Hx_Obj_loreline_ncharacterdecl) *Hx_Obj_loreline_ncharacterdecl {
		var _hx_closure_src_701 func(node any) any
		_ = _hx_closure_src_701
		_hx_closure_src_701 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_ncharacterdecl) *Hx_Obj_loreline_ncharacterdecl {
			var hx_dyncast_300 any
			_ = hx_dyncast_300
			hx_dyncast_300 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_701(((any)(node))), "Hx_Obj_loreline_ncharacterdecl")
			var _hx_tmp_2 *Hx_Obj_loreline_ncharacterdecl
			_ = _hx_tmp_2
			if hx_dyncast_300 != nil {
				_hx_tmp_2 = hx_dyncast_300.(interface {
					Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
				}).Hx_As_Hx_Obj_loreline_ncharacterdecl()
			} else {
				_hx_tmp_2 = nil
			}

			return _hx_tmp_2
		}
	}()(characterNode)
	for {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_3 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_3 = false
		}

		if !(_hx_tmp_3) {
			break
		}

		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_7 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_8 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = this.VTable.Hx_Field_match(_hx_tmp_8, nil)
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_4 = _hx_tmp_6
			}

			if !(_hx_tmp_4) {
				break
			}
		}

		var _hx_tmp_9 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _hx_tmp_9
		_hx_tmp_9 = characterNode.Hx_Field_fields
		_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_9, this.VTable.Hx_Field_parseObjectField()))
		for {
			var _hx_tmp_10 bool
			_ = _hx_tmp_10
			var _hx_tmp_11 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_11
			_hx_tmp_11 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_11, nil) {
				_hx_tmp_10 = true
			} else {
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_13 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_13
					_hx_tmp_13 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = _hx_tmp_13
					_hx_tmp_12 = this.VTable.Hx_Field_match(_hx_tmp_14, nil)
				} else {
					_hx_tmp_12 = false
				}

				_hx_tmp_10 = _hx_tmp_12
			}

			if !(_hx_tmp_10) {
				break
			}
		}
	}

	for {
		var _hx_tmp_15 bool
		_ = _hx_tmp_15
		var _hx_tmp_16 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_16
		_hx_tmp_16 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_16, nil) {
			_hx_tmp_15 = true
		} else {
			var _hx_tmp_17 bool
			_ = _hx_tmp_17
			if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
				var _hx_tmp_18 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_18
				_hx_tmp_18 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_19 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_19
				_hx_tmp_19 = _hx_tmp_18
				_hx_tmp_17 = this.VTable.Hx_Field_match(_hx_tmp_19, nil)
			} else {
				_hx_tmp_17 = false
			}

			_hx_tmp_15 = _hx_tmp_17
		}

		if !(_hx_tmp_15) {
			break
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	characterNode.Hx_Field_pos = characterNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return characterNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseTextStatement() *Hx_Obj_loreline_ntextstatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var statement *Hx_Obj_loreline_ntextstatement
	_ = statement
	statement = func() func(node *Hx_Obj_loreline_ntextstatement) *Hx_Obj_loreline_ntextstatement {
		var _hx_closure_src_702 func(node any) any
		_ = _hx_closure_src_702
		_hx_closure_src_702 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_ntextstatement) *Hx_Obj_loreline_ntextstatement {
			var hx_dyncast_301 any
			_ = hx_dyncast_301
			hx_dyncast_301 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_702(((any)(node))), "Hx_Obj_loreline_ntextstatement")
			var _hx_tmp_0 *Hx_Obj_loreline_ntextstatement
			_ = _hx_tmp_0
			if hx_dyncast_301 != nil {
				_hx_tmp_0 = hx_dyncast_301.(interface {
					Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
				}).Hx_As_Hx_Obj_loreline_ntextstatement()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(Hx_Obj_loreline_ntextstatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, nil, nil, nil))
	statement.Hx_Field_content = this.VTable.Hx_Field_parseStringLiteral()
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if !this.Hx_Field_lineBreakAfterToken {
		var _hx_tmp_2 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Enum_loreline_tokentype_KwIf{}

		var _hx_tmp_3 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = this.VTable.Hx_Field_match(_hx_tmp_3, nil)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var ifPos *Hx_Obj_loreline_position
		_ = ifPos
		ifPos = this.VTable.Hx_Field_previous().Hx_Field_pos
		statement.Hx_Field_condition = this.VTable.Hx_Field_parseConditionExpression()
		statement.Hx_Field_conditionStyle = this.VTable.Hx_Field_takeConditionParens(statement.Hx_Field_condition)
		statement.Hx_Field_conditionPos = ifPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	}

	statement.Hx_Field_pos = statement.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return statement
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseChoiceStatement() *Hx_Obj_loreline_nchoicestatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var choiceNode *Hx_Obj_loreline_nchoicestatement
	_ = choiceNode
	choiceNode = Hx_Obj_loreline_nchoicestatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(2))), startPos, HxMakeArray[*Hx_Obj_loreline_nchoiceoption](), nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwChoice{})
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_1 = ((int32)(1))
	} else {
		_hx_tmp_1 = ((int32)(0))
	}

	choiceNode.Hx_Field_style = _hx_tmp_1
	func() func(node *Hx_Obj_loreline_nchoicestatement) *Hx_Obj_loreline_nchoicestatement {
		var _hx_closure_src_703 func(node any) any
		_ = _hx_closure_src_703
		_hx_closure_src_703 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nchoicestatement) *Hx_Obj_loreline_nchoicestatement {
			var hx_dyncast_302 any
			_ = hx_dyncast_302
			hx_dyncast_302 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_703(((any)(node))), "Hx_Obj_loreline_nchoicestatement")
			var _hx_tmp_2 *Hx_Obj_loreline_nchoicestatement
			_ = _hx_tmp_2
			if hx_dyncast_302 != nil {
				_hx_tmp_2 = hx_dyncast_302.(interface {
					Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
				}).Hx_As_Hx_Obj_loreline_nchoicestatement()
			} else {
				_hx_tmp_2 = nil
			}

			return _hx_tmp_2
		}
	}()(choiceNode)
	for {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_3 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_3 = false
		}

		if !(_hx_tmp_3) {
			break
		}

		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_7 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_8 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = this.VTable.Hx_Field_match(_hx_tmp_8, nil)
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_4 = _hx_tmp_6
			}

			if !(_hx_tmp_4) {
				break
			}
		}

		var _hx_tmp_9 HxArray[*Hx_Obj_loreline_nchoiceoption]
		_ = _hx_tmp_9
		_hx_tmp_9 = choiceNode.Hx_Field_options
		_ = (Hx_Array_Push[*Hx_Obj_loreline_nchoiceoption](_hx_tmp_9, this.VTable.Hx_Field_parseChoiceOption(blockEnd)))
		for {
			var _hx_tmp_10 bool
			_ = _hx_tmp_10
			var _hx_tmp_11 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_11
			_hx_tmp_11 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_11, nil) {
				_hx_tmp_10 = true
			} else {
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_13 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_13
					_hx_tmp_13 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = _hx_tmp_13
					_hx_tmp_12 = this.VTable.Hx_Field_match(_hx_tmp_14, nil)
				} else {
					_hx_tmp_12 = false
				}

				_hx_tmp_10 = _hx_tmp_12
			}

			if !(_hx_tmp_10) {
				break
			}
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	choiceNode.Hx_Field_pos = choiceNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return choiceNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseChoiceOption(blockEnd Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_nchoiceoption {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var choiceOption *Hx_Obj_loreline_nchoiceoption
	_ = choiceOption
	choiceOption = func() func(node *Hx_Obj_loreline_nchoiceoption) *Hx_Obj_loreline_nchoiceoption {
		var _hx_closure_src_704 func(node any) any
		_ = _hx_closure_src_704
		_hx_closure_src_704 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nchoiceoption) *Hx_Obj_loreline_nchoiceoption {
			var hx_dyncast_303 any
			_ = hx_dyncast_303
			hx_dyncast_303 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_704(((any)(node))), "Hx_Obj_loreline_nchoiceoption")
			var _hx_tmp_0 *Hx_Obj_loreline_nchoiceoption
			_ = _hx_tmp_0
			if hx_dyncast_303 != nil {
				_hx_tmp_0 = hx_dyncast_303.(interface {
					Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
				}).Hx_As_Hx_Obj_loreline_nchoiceoption()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(Hx_Obj_loreline_nchoiceoption_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, nil, nil, nil, ((int32)(0)), HxMakeArray[*Hx_Obj_loreline_astnode](), nil, nil))
	var errorPos *Hx_Obj_loreline_position
	_ = errorPos
	var isInsertion bool
	_ = isInsertion
	isInsertion = false
	{
		var hx_try_return *Hx_Obj_loreline_nchoiceoption
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
						_ = e
						{
							this.VTable.Hx_Field_addError(e)
							errorPos = this.VTable.Hx_Field_currentPos()
							if this.VTable.Hx_Field_currentPos().Hx_Field_offset == startPos.Hx_Field_offset {
								this.VTable.Hx_Field_advance(nil)
							}

							choiceOption.Hx_Field_text = Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), ((int32)(0)), HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), Hx_Enum_loreline_stringparttype_Raw{"?"}, nil, nil)), nil, nil)
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpPlus{})) {
					isInsertion = true
					choiceOption.Hx_Field_insertion = this.VTable.Hx_Field_parseInsertion()
				} else {
					var _hx_tmp_1 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Enum_loreline_tokentype_OpMinus{}

					if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
						choiceOption.Hx_Field_once = true
					}

					choiceOption.Hx_Field_text = this.VTable.Hx_Field_parseStringLiteral()
				}
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	var _hx_tmp_2 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Enum_loreline_tokentype_KwIf{}

	if this.VTable.Hx_Field_match(_hx_tmp_2, nil) {
		var ifPos *Hx_Obj_loreline_position
		_ = ifPos
		ifPos = this.VTable.Hx_Field_previous().Hx_Field_pos
		var offset int32
		_ = offset
		offset = this.VTable.Hx_Field_currentPos().Hx_Field_offset
		{
			var hx_try_return *Hx_Obj_loreline_nchoiceoption
			_ = hx_try_return
			var hx_try_state int
			_ = hx_try_state

			func() {
				defer func() {
					if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
						if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
							_ = e
							{
								this.VTable.Hx_Field_addError(e)
								errorPos = this.VTable.Hx_Field_currentPos()
								if this.VTable.Hx_Field_currentPos().Hx_Field_offset == offset {
									this.VTable.Hx_Field_advance(nil)
								}
							}
						} else {
							panic(_hx_ex)
						}
					}
				}()
				{
					choiceOption.Hx_Field_condition = this.VTable.Hx_Field_parseConditionExpression()
					choiceOption.Hx_Field_conditionStyle = this.VTable.Hx_Field_takeConditionParens(choiceOption.Hx_Field_condition)
					choiceOption.Hx_Field_conditionPos = ifPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
				}
			}()
			if hx_try_state == 1 {
				return hx_try_return
			}
		}
	}

	if !isInsertion {
		if this.VTable.Hx_Field_checkBlockStart() {
			choiceOption.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
			var offset int32
			_ = offset
			offset = this.VTable.Hx_Field_currentPos().Hx_Field_offset
			{
				var hx_try_return *Hx_Obj_loreline_nchoiceoption
				_ = hx_try_return
				var hx_try_state int
				_ = hx_try_state

				func() {
					defer func() {
						if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
							if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
								_ = e
								{
									this.VTable.Hx_Field_addError(e)
									errorPos = this.VTable.Hx_Field_currentPos()
									if this.VTable.Hx_Field_currentPos().Hx_Field_offset == offset {
										this.VTable.Hx_Field_advance(nil)
									}

									if choiceOption.Hx_Field_body.Len() == 0 {
										choiceOption.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]((func() *Hx_Obj_loreline_astnode {
											_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
											if _hx_up != nil {
												return &_hx_up.Hx_Obj_loreline_astnode
											}
											return nil
										})())
									}
								}
							} else {
								panic(_hx_ex)
							}
						}
					}()
					{
						choiceOption.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(choiceOption.Hx_Field_body)
					}
				}()
				if hx_try_state == 1 {
					return hx_try_return
				}
			}
		} else {
			if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_Arrow{})) {
				var offset int32
				_ = offset
				offset = this.VTable.Hx_Field_currentPos().Hx_Field_offset
				{
					var hx_try_return *Hx_Obj_loreline_nchoiceoption
					_ = hx_try_return
					var hx_try_state int
					_ = hx_try_state

					func() {
						defer func() {
							if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
								if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
									_ = e
									{
										this.VTable.Hx_Field_addError(e)
										errorPos = this.VTable.Hx_Field_currentPos()
										if this.VTable.Hx_Field_currentPos().Hx_Field_offset == offset {
											this.VTable.Hx_Field_advance(nil)
										}

										if choiceOption.Hx_Field_body.Len() == 0 {
											choiceOption.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]((func() *Hx_Obj_loreline_astnode {
												_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
												if _hx_up != nil {
													return &_hx_up.Hx_Obj_loreline_astnode
												}
												return nil
											})())
										}
									}
								} else {
									panic(_hx_ex)
								}
							}
						}()
						{
							choiceOption.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
							var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
							_ = _hx_tmp_3
							_hx_tmp_3 = choiceOption.Hx_Field_body
							_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_3, this.VTable.Hx_Field_parseNode(nil)))
						}
					}()
					if hx_try_state == 1 {
						return hx_try_return
					}
				}

				choiceOption.Hx_Field_style = ((int32)(0))
			}
		}
	}

	choiceOption.Hx_Field_pos = choiceOption.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	if errorPos != nil {
		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if !this.VTable.Hx_Field_isAtEnd() {
				_hx_tmp_4 = (this.VTable.Hx_Field_currentPos().Hx_Field_line <= (errorPos.Hx_Field_line + 1))
			} else {
				_hx_tmp_4 = false
			}

			if !(_hx_tmp_4) {
				break
			}

			this.VTable.Hx_Field_advance(nil)
		}
	}

	return choiceOption
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseWhenStatement() *Hx_Obj_loreline_nwhenstatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var whenNode *Hx_Obj_loreline_nwhenstatement
	_ = whenNode
	whenNode = Hx_Obj_loreline_nwhenstatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(2))), startPos, HxNullable[string]{Value: HxStringNull}, HxMakeArray[*Hx_Obj_loreline_nwhenrule](), nil, nil)
	var whenToken *Hx_Obj_loreline_token
	_ = whenToken
	whenToken = this.Hx_Field_tokens.Get(this.Hx_Field_current)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwWhen{})
	{
		var _g Hx_Enum_loreline_tokentype
		_ = _g
		_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
		if _g.Hx_Field_enumIndex() == 20 {
			var name string
			_ = name
			name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
			if this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_pos.Hx_Field_line == whenToken.Hx_Field_pos.Hx_Field_line {
				whenNode.Hx_Field_strategy = HxNullableString(name)
				whenNode.Hx_Field_strategyPos = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_pos
				this.VTable.Hx_Field_advance(nil)
			}
		}
	}

	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_1 = ((int32)(1))
	} else {
		_hx_tmp_1 = ((int32)(0))
	}

	whenNode.Hx_Field_style = _hx_tmp_1
	func() func(node *Hx_Obj_loreline_nwhenstatement) *Hx_Obj_loreline_nwhenstatement {
		var _hx_closure_src_705 func(node any) any
		_ = _hx_closure_src_705
		_hx_closure_src_705 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nwhenstatement) *Hx_Obj_loreline_nwhenstatement {
			var hx_dyncast_304 any
			_ = hx_dyncast_304
			hx_dyncast_304 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_705(((any)(node))), "Hx_Obj_loreline_nwhenstatement")
			var _hx_tmp_2 *Hx_Obj_loreline_nwhenstatement
			_ = _hx_tmp_2
			if hx_dyncast_304 != nil {
				_hx_tmp_2 = hx_dyncast_304.(interface {
					Hx_As_Hx_Obj_loreline_nwhenstatement() *Hx_Obj_loreline_nwhenstatement
				}).Hx_As_Hx_Obj_loreline_nwhenstatement()
			} else {
				_hx_tmp_2 = nil
			}

			return _hx_tmp_2
		}
	}()(whenNode)
	for {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_3 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_3 = false
		}

		if !(_hx_tmp_3) {
			break
		}

		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_7 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_8 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = this.VTable.Hx_Field_match(_hx_tmp_8, nil)
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_4 = _hx_tmp_6
			}

			if !(_hx_tmp_4) {
				break
			}
		}

		var _hx_tmp_9 bool
		_ = _hx_tmp_9
		if this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_9 = true
		} else {
			_hx_tmp_9 = this.VTable.Hx_Field_isAtEnd()
		}

		if _hx_tmp_9 {
			break
		}

		var _hx_tmp_10 HxArray[*Hx_Obj_loreline_nwhenrule]
		_ = _hx_tmp_10
		_hx_tmp_10 = whenNode.Hx_Field_rules
		_ = (Hx_Array_Push[*Hx_Obj_loreline_nwhenrule](_hx_tmp_10, this.VTable.Hx_Field_parseWhenRule()))
		for {
			var _hx_tmp_11 bool
			_ = _hx_tmp_11
			var _hx_tmp_12 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_12
			_hx_tmp_12 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_12, nil) {
				_hx_tmp_11 = true
			} else {
				var _hx_tmp_13 bool
				_ = _hx_tmp_13
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_15 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_15
					_hx_tmp_15 = _hx_tmp_14
					_hx_tmp_13 = this.VTable.Hx_Field_match(_hx_tmp_15, nil)
				} else {
					_hx_tmp_13 = false
				}

				_hx_tmp_11 = _hx_tmp_13
			}

			if !(_hx_tmp_11) {
				break
			}
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	whenNode.Hx_Field_pos = whenNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return whenNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseWhenRule() *Hx_Obj_loreline_nwhenrule {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var rule *Hx_Obj_loreline_nwhenrule
	_ = rule
	rule = func() func(node *Hx_Obj_loreline_nwhenrule) *Hx_Obj_loreline_nwhenrule {
		var _hx_closure_src_706 func(node any) any
		_ = _hx_closure_src_706
		_hx_closure_src_706 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nwhenrule) *Hx_Obj_loreline_nwhenrule {
			var hx_dyncast_305 any
			_ = hx_dyncast_305
			hx_dyncast_305 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_706(((any)(node))), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_0 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_0
			if hx_dyncast_305 != nil {
				_hx_tmp_0 = hx_dyncast_305.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(Hx_Obj_loreline_nwhenrule_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, nil, HxMakeArray[*Hx_Obj_loreline_astnode](), nil, nil))
	var errorPos *Hx_Obj_loreline_position
	_ = errorPos
	{
		var hx_try_return *Hx_Obj_loreline_nwhenrule
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
						_ = e
						{
							this.VTable.Hx_Field_addError(e)
							errorPos = this.VTable.Hx_Field_currentPos()
							if this.VTable.Hx_Field_currentPos().Hx_Field_offset == startPos.Hx_Field_offset {
								this.VTable.Hx_Field_advance(nil)
							}
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpPlus{})) {
					rule.Hx_Field_insertion = this.VTable.Hx_Field_parseInsertion()
					var _hx_tmp_1 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Enum_loreline_tokentype_KwIf{}

					if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
						rule.Hx_Field_insertionCondition = this.VTable.Hx_Field_parseConditionExpression()
						rule.Hx_Field_insertionConditionStyle = this.VTable.Hx_Field_takeConditionParens(rule.Hx_Field_insertionCondition)
					}
				} else {
					var minus *Hx_Obj_loreline_token
					_ = minus
					minus = this.Hx_Field_tokens.Get(this.Hx_Field_current)
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					var _hx_tmp_3 bool
					_ = _hx_tmp_3
					if (minus.Hx_Field_type == Hx_Enum_loreline_tokentype_OpMinus{}) {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = (this.Hx_Field_current + 1)
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = _hx_tmp_4
						_hx_tmp_3 = (_hx_tmp_5 < this.Hx_Field_tokens.Len())
					} else {
						_hx_tmp_3 = false
					}

					if _hx_tmp_3 {
						_hx_tmp_2 = (this.Hx_Field_tokens.Get((this.Hx_Field_current + 1)).Hx_Field_pos.Hx_Field_offset > (minus.Hx_Field_pos.Hx_Field_offset + minus.Hx_Field_pos.Hx_Field_length))
					} else {
						_hx_tmp_2 = false
					}

					if _hx_tmp_2 {
						this.VTable.Hx_Field_advance(nil)
						rule.Hx_Field_once = true
					}

					rule.Hx_Field_condition = this.VTable.Hx_Field_parseConditionExpression()
					rule.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(rule.Hx_Field_body)
				}
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	rule.Hx_Field_pos = rule.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	if errorPos != nil {
		for {
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			if !this.VTable.Hx_Field_isAtEnd() {
				_hx_tmp_6 = (this.VTable.Hx_Field_currentPos().Hx_Field_line <= (errorPos.Hx_Field_line + 1))
			} else {
				_hx_tmp_6 = false
			}

			if !(_hx_tmp_6) {
				break
			}

			this.VTable.Hx_Field_advance(nil)
		}
	}

	return rule
}

func (this *Hx_Obj_loreline_parser) Hx_Field_isExpressionStart() bool {
	switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
	case 16:
		{
			return true
		}

	case 17:
		{
			return true
		}

	case 18:
		{
			return true
		}

	case 20:
		{
			if (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) {
				return false
			} else {
				return true
			}
		}

	case 19, 28, 46, 48, 50:
		{
			return true
		}

	case 40:
		{
			return true
		}

	default:
		{
			return false
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseExpressionStatement() *Hx_Obj_loreline_astnode {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseExpression()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	var _hx_tmp_5 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_5
	_hx_tmp_5 = Hx_Enum_loreline_tokentype_OpAssign{}

	if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
		_hx_tmp_4 = true
	} else {
		var _hx_tmp_6 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_6
		_hx_tmp_6 = Hx_Enum_loreline_tokentype_OpPlusAssign{}

		var _hx_tmp_7 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_6
		_hx_tmp_4 = this.VTable.Hx_Field_match(_hx_tmp_7, nil)
	}

	if _hx_tmp_4 {
		_hx_tmp_3 = true
	} else {
		var _hx_tmp_8 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_8
		_hx_tmp_8 = Hx_Enum_loreline_tokentype_OpMinusAssign{}

		var _hx_tmp_9 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_9
		_hx_tmp_9 = _hx_tmp_8
		_hx_tmp_3 = this.VTable.Hx_Field_match(_hx_tmp_9, nil)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_10 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_10
		_hx_tmp_10 = Hx_Enum_loreline_tokentype_OpMultiplyAssign{}

		var _hx_tmp_11 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_11
		_hx_tmp_11 = _hx_tmp_10
		_hx_tmp_2 = this.VTable.Hx_Field_match(_hx_tmp_11, nil)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_12 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_12
		_hx_tmp_12 = Hx_Enum_loreline_tokentype_OpDivideAssign{}

		var _hx_tmp_13 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_12
		_hx_tmp_1 = this.VTable.Hx_Field_match(_hx_tmp_13, nil)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_14 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_14
		_hx_tmp_14 = Hx_Enum_loreline_tokentype_OpUnquotedAssign{}

		var _hx_tmp_15 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_15
		_hx_tmp_15 = _hx_tmp_14
		_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_15, nil)
	}

	if _hx_tmp_0 {
		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var assignment *Hx_Obj_loreline_nassign
		_ = assignment
		assignment = func() func(node *Hx_Obj_loreline_nassign) *Hx_Obj_loreline_nassign {
			var _hx_closure_src_707 func(node any) any
			_ = _hx_closure_src_707
			_hx_closure_src_707 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nassign) *Hx_Obj_loreline_nassign {
				var hx_dyncast_306 any
				_ = hx_dyncast_306
				hx_dyncast_306 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_707(((any)(node))), "Hx_Obj_loreline_nassign")
				var _hx_tmp_16 *Hx_Obj_loreline_nassign
				_ = _hx_tmp_16
				if hx_dyncast_306 != nil {
					_hx_tmp_16 = hx_dyncast_306.(interface {
						Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
					}).Hx_As_Hx_Obj_loreline_nassign()
				} else {
					_hx_tmp_16 = nil
				}

				return _hx_tmp_16
			}
		}()(Hx_Obj_loreline_nassign_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		assignment.Hx_Field_value = this.VTable.Hx_Field_parseExpression()
		assignment.Hx_Field_pos = assignment.Hx_Field_pos.VTable.Hx_Field_extendedTo(assignment.Hx_Field_value.Hx_Field_pos)
		return (func() *Hx_Obj_loreline_astnode {
			_hx_up := assignment
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
	}

	return (func() *Hx_Obj_loreline_astnode {
		_hx_up := func() func(node *Hx_Obj_loreline_nexpr) *Hx_Obj_loreline_nexpr {
			var _hx_closure_src_709 func(node any) any
			_ = _hx_closure_src_709
			_hx_closure_src_709 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nexpr) *Hx_Obj_loreline_nexpr {
				var hx_dyncast_308 any
				_ = hx_dyncast_308
				hx_dyncast_308 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_709(((any)(node))), "Hx_Obj_loreline_nexpr")
				var _hx_tmp_17 *Hx_Obj_loreline_nexpr
				_ = _hx_tmp_17
				if hx_dyncast_308 != nil {
					_hx_tmp_17 = hx_dyncast_308.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
				} else {
					_hx_tmp_17 = nil
				}

				return _hx_tmp_17
			}
		}()(expr)
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
}

func (this *Hx_Obj_loreline_parser) Hx_Field_isKnownNodeStart() bool {
	switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
	case 16:
		{
			return true
		}

	case 20:
		{
			if (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) {
				return true
			} else {
				return (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_Arrow{})
			}
		}

	case 1, 2, 3, 4, 5, 6, 41:
		{
			return true
		}

	default:
		{
			return false
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseIfStatement() *Hx_Obj_loreline_nifstatement {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var ifNode *Hx_Obj_loreline_nifstatement
	_ = ifNode
	ifNode = Hx_Obj_loreline_nifstatement_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, nil, ((int32)(0)), nil, nil, nil, nil, nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwIf{})
	ifNode.Hx_Field_condition = this.VTable.Hx_Field_parseConditionExpression()
	ifNode.Hx_Field_conditionStyle = this.VTable.Hx_Field_takeConditionParens(ifNode.Hx_Field_condition)
	for {
		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_LineBreak{}

		if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
			break
		}
	}

	func() func(node *Hx_Obj_loreline_nifstatement) *Hx_Obj_loreline_nifstatement {
		var _hx_closure_src_710 func(node any) any
		_ = _hx_closure_src_710
		_hx_closure_src_710 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nifstatement) *Hx_Obj_loreline_nifstatement {
			var hx_dyncast_309 any
			_ = hx_dyncast_309
			hx_dyncast_309 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_710(((any)(node))), "Hx_Obj_loreline_nifstatement")
			var _hx_tmp_1 *Hx_Obj_loreline_nifstatement
			_ = _hx_tmp_1
			if hx_dyncast_309 != nil {
				_hx_tmp_1 = hx_dyncast_309.(interface {
					Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
				}).Hx_As_Hx_Obj_loreline_nifstatement()
			} else {
				_hx_tmp_1 = nil
			}

			return _hx_tmp_1
		}
	}()(ifNode)
	ifNode.Hx_Field_thenBranch = Hx_Obj_loreline_nblock_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), this.VTable.Hx_Field_currentPos(), nil, nil, nil)
	ifNode.Hx_Field_thenBranch.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
	ifNode.Hx_Field_thenBranch.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(ifNode.Hx_Field_thenBranch.Hx_Field_body)
	var elseToken *Hx_Obj_loreline_token
	_ = elseToken
	elseToken = this.Hx_Field_tokens.Get(this.Hx_Field_current)
	if (elseToken != nil) && (elseToken.Hx_Field_type == Hx_Enum_loreline_tokentype_KwElse{}) {
		this.VTable.Hx_Field_advance(nil)
		for {
			var _hx_tmp_2 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Enum_loreline_tokentype_LineBreak{}

			if !(this.VTable.Hx_Field_match(_hx_tmp_2, nil)) {
				break
			}
		}

		func() func(node *Hx_Obj_loreline_nifstatement, elseToken *Hx_Obj_loreline_token) *Hx_Obj_loreline_nifstatement {
			var _hx_closure_src_711 func(node any, elseToken *Hx_Obj_loreline_token) any
			_ = _hx_closure_src_711
			_hx_closure_src_711 = this.VTable.Hx_Field_attachElseComments
			return func(node *Hx_Obj_loreline_nifstatement, elseToken *Hx_Obj_loreline_token) *Hx_Obj_loreline_nifstatement {
				var hx_dyncast_310 any
				_ = hx_dyncast_310
				hx_dyncast_310 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_711(((any)(node)), elseToken), "Hx_Obj_loreline_nifstatement")
				var _hx_tmp_3 *Hx_Obj_loreline_nifstatement
				_ = _hx_tmp_3
				if hx_dyncast_310 != nil {
					_hx_tmp_3 = hx_dyncast_310.(interface {
						Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
					}).Hx_As_Hx_Obj_loreline_nifstatement()
				} else {
					_hx_tmp_3 = nil
				}

				return _hx_tmp_3
			}
		}()(ifNode, elseToken)
		if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_KwIf{})) {
			ifNode.Hx_Field_elseBranch = Hx_Obj_loreline_nblock_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), this.VTable.Hx_Field_currentPos(), nil, nil, nil)
			var tmp *Hx_Obj_loreline_nifstatement
			_ = tmp
			tmp = this.VTable.Hx_Field_parseIfStatement()
			ifNode.Hx_Field_elseBranch.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]((func() *Hx_Obj_loreline_astnode {
				_hx_up := tmp
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())
			ifNode.Hx_Field_elseBranch.Hx_Field_style = ((int32)(0))
		} else {
			ifNode.Hx_Field_elseBranch = Hx_Obj_loreline_nblock_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), this.VTable.Hx_Field_currentPos(), nil, nil, nil)
			ifNode.Hx_Field_elseBranch.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
			ifNode.Hx_Field_elseBranch.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(ifNode.Hx_Field_elseBranch.Hx_Field_body)
		}
	}

	ifNode.Hx_Field_pos = ifNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return ifNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseAlternative() *Hx_Obj_loreline_nalternative {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var mode int32
	_ = mode
	switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
	case 9:
		{
			mode = ((int32)(0))
		}

	case 10:
		{
			mode = ((int32)(1))
		}

	case 11:
		{
			mode = ((int32)(2))
		}

	case 12:
		{
			mode = ((int32)(3))
		}

	case 13:
		{
			mode = ((int32)(4))
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected alternative keyword", this.VTable.Hx_Field_currentPos())))))
		}
	}

	var altNode *Hx_Obj_loreline_nalternative
	_ = altNode
	altNode = Hx_Obj_loreline_nalternative_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(2))), startPos, mode, HxMakeArray[*Hx_Obj_loreline_nblock](), nil, nil)
	this.VTable.Hx_Field_advance(nil)
	for {
		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_LineBreak{}

		if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
			break
		}
	}

	func() func(node *Hx_Obj_loreline_nalternative) *Hx_Obj_loreline_nalternative {
		var _hx_closure_src_712 func(node any) any
		_ = _hx_closure_src_712
		_hx_closure_src_712 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nalternative) *Hx_Obj_loreline_nalternative {
			var hx_dyncast_311 any
			_ = hx_dyncast_311
			hx_dyncast_311 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_712(((any)(node))), "Hx_Obj_loreline_nalternative")
			var _hx_tmp_1 *Hx_Obj_loreline_nalternative
			_ = _hx_tmp_1
			if hx_dyncast_311 != nil {
				_hx_tmp_1 = hx_dyncast_311.(interface {
					Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
				}).Hx_As_Hx_Obj_loreline_nalternative()
			} else {
				_hx_tmp_1 = nil
			}

			return _hx_tmp_1
		}
	}()(altNode)
	var firstBlock *Hx_Obj_loreline_nblock
	_ = firstBlock
	firstBlock = Hx_Obj_loreline_nblock_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), this.VTable.Hx_Field_currentPos(), nil, nil, nil)
	firstBlock.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
	firstBlock.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(firstBlock.Hx_Field_body)
	_ = (Hx_Array_Push[*Hx_Obj_loreline_nblock](altNode.Hx_Field_items, firstBlock))
	altNode.Hx_Field_style = firstBlock.Hx_Field_style
	for {
		if !(this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_Separator{})) {
			break
		}

		var separatorToken *Hx_Obj_loreline_token
		_ = separatorToken
		separatorToken = this.Hx_Field_tokens.Get(this.Hx_Field_current)
		this.VTable.Hx_Field_advance(nil)
		for {
			var _hx_tmp_2 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Enum_loreline_tokentype_LineBreak{}

			if !(this.VTable.Hx_Field_match(_hx_tmp_2, nil)) {
				break
			}
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if this.Hx_Field_pendingComments != nil {
			_hx_tmp_3 = (this.Hx_Field_pendingComments.Len() > 0)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			if altNode.Hx_Field_separatorComments == nil {
				altNode.Hx_Field_separatorComments = HxMakeArray[HxArray[*Hx_Obj_loreline_comment]]()
			}

			var sepComments HxArray[*Hx_Obj_loreline_comment]
			_ = sepComments
			sepComments = HxMakeArray[*Hx_Obj_loreline_comment]()
			var remainingComments HxArray[*Hx_Obj_loreline_comment]
			_ = remainingComments
			remainingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = this.Hx_Field_pendingComments
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					if !(_hx_tmp_4 < _g1.Len()) {
						break
					}

					var comment *Hx_Obj_loreline_comment
					_ = comment
					comment = _g1.Get(_g)
					_g++
					if comment.Hx_Field_pos.Hx_Field_line <= separatorToken.Hx_Field_pos.Hx_Field_line {
						_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](sepComments, comment))
					} else {
						_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](remainingComments, comment))
					}
				}
			}

			_ = (Hx_Array_Push[HxArray[*Hx_Obj_loreline_comment]](altNode.Hx_Field_separatorComments, sepComments))
			this.Hx_Field_pendingComments = remainingComments
		}

		var itemBlock *Hx_Obj_loreline_nblock
		_ = itemBlock
		itemBlock = Hx_Obj_loreline_nblock_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), this.VTable.Hx_Field_currentPos(), nil, nil, nil)
		itemBlock.Hx_Field_body = HxMakeArray[*Hx_Obj_loreline_astnode]()
		itemBlock.Hx_Field_style = this.VTable.Hx_Field_parseStatementBlock(itemBlock.Hx_Field_body)
		_ = (Hx_Array_Push[*Hx_Obj_loreline_nblock](altNode.Hx_Field_items, itemBlock))
	}

	altNode.Hx_Field_pos = altNode.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return altNode
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseTransition() *Hx_Obj_loreline_ntransition {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_Arrow{})
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Enum_loreline_tokentype_Dot{}

	if this.VTable.Hx_Field_match(_hx_tmp_0, nil) {
		return func() func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
			var _hx_closure_src_713 func(node any) any
			_ = _hx_closure_src_713
			_hx_closure_src_713 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
				var hx_dyncast_312 any
				_ = hx_dyncast_312
				hx_dyncast_312 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_713(((any)(node))), "Hx_Obj_loreline_ntransition")
				var _hx_tmp_1 *Hx_Obj_loreline_ntransition
				_ = _hx_tmp_1
				if hx_dyncast_312 != nil {
					_hx_tmp_1 = hx_dyncast_312.(interface {
						Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
					}).Hx_As_Hx_Obj_loreline_ntransition()
				} else {
					_hx_tmp_1 = nil
				}

				return _hx_tmp_1
			}
		}()(Hx_Obj_loreline_ntransition_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), ".", this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos, nil, nil))
	}

	if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_KwBeat{})) {
		var targetStart *Hx_Obj_loreline_position
		_ = targetStart
		targetStart = this.VTable.Hx_Field_currentPos()
		this.VTable.Hx_Field_advance(nil)
		this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_LParen{})
		var callArgs HxArray[*Hx_Obj_loreline_nexpr]
		_ = callArgs
		callArgs = this.VTable.Hx_Field_parseCallArguments()
		if callArgs.Len() == 0 {
			this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("beat() requires a target expression", targetStart))
		}

		var node *Hx_Obj_loreline_ntransition
		_ = node
		node = Hx_Obj_loreline_ntransition_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), HxStringNull, targetStart.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), nil, nil)
		var _hx_tmp_2 *Hx_Obj_loreline_nexpr
		_ = _hx_tmp_2
		if callArgs.Len() > 0 {
			_hx_tmp_2 = callArgs.Get(0)
		} else {
			_hx_tmp_2 = nil
		}

		node.Hx_Field_targetExpr = _hx_tmp_2
		if callArgs.Len() > 1 {
			var _hx_tmp_3 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_3
			_hx_tmp_3 = callArgs
			node.Hx_Field_args = (Hx_Array_Slice[*Hx_Obj_loreline_nexpr](_hx_tmp_3, 1, HxNullable[int32]{}))
		}

		return func() func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
			var _hx_closure_src_714 func(node any) any
			_ = _hx_closure_src_714
			_hx_closure_src_714 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
				var hx_dyncast_313 any
				_ = hx_dyncast_313
				hx_dyncast_313 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_714(((any)(node))), "Hx_Obj_loreline_ntransition")
				var _hx_tmp_4 *Hx_Obj_loreline_ntransition
				_ = _hx_tmp_4
				if hx_dyncast_313 != nil {
					_hx_tmp_4 = hx_dyncast_313.(interface {
						Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
					}).Hx_As_Hx_Obj_loreline_ntransition()
				} else {
					_hx_tmp_4 = nil
				}

				return _hx_tmp_4
			}
		}()(node)
	}

	var target string
	_ = target
	target = this.VTable.Hx_Field_expectIdentifier()
	var targetPos *Hx_Obj_loreline_position
	_ = targetPos
	targetPos = this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos
	var args HxArray[*Hx_Obj_loreline_nexpr]
	_ = args
	var _hx_tmp_5 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_5
	_hx_tmp_5 = Hx_Enum_loreline_tokentype_LParen{}

	if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
		args = this.VTable.Hx_Field_parseCallArguments()
	}

	var node *Hx_Obj_loreline_ntransition
	_ = node
	node = Hx_Obj_loreline_ntransition_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), target, targetPos, nil, nil)
	node.Hx_Field_args = args
	return func() func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
		var _hx_closure_src_715 func(node any) any
		_ = _hx_closure_src_715
		_hx_closure_src_715 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_ntransition) *Hx_Obj_loreline_ntransition {
			var hx_dyncast_314 any
			_ = hx_dyncast_314
			hx_dyncast_314 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_715(((any)(node))), "Hx_Obj_loreline_ntransition")
			var _hx_tmp_6 *Hx_Obj_loreline_ntransition
			_ = _hx_tmp_6
			if hx_dyncast_314 != nil {
				_hx_tmp_6 = hx_dyncast_314.(interface {
					Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
				}).Hx_As_Hx_Obj_loreline_ntransition()
			} else {
				_hx_tmp_6 = nil
			}

			return _hx_tmp_6
		}
	}()(node)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseInsertion() *Hx_Obj_loreline_ninsertion {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_OpPlus{})
	if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_KwBeat{})) {
		var targetStart *Hx_Obj_loreline_position
		_ = targetStart
		targetStart = this.VTable.Hx_Field_currentPos()
		this.VTable.Hx_Field_advance(nil)
		this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_LParen{})
		var callArgs HxArray[*Hx_Obj_loreline_nexpr]
		_ = callArgs
		callArgs = this.VTable.Hx_Field_parseCallArguments()
		if callArgs.Len() == 0 {
			this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("beat() requires a target expression", targetStart))
		}

		var node *Hx_Obj_loreline_ninsertion
		_ = node
		node = Hx_Obj_loreline_ninsertion_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), HxStringNull, targetStart.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), nil, nil)
		var _hx_tmp_0 *Hx_Obj_loreline_nexpr
		_ = _hx_tmp_0
		if callArgs.Len() > 0 {
			_hx_tmp_0 = callArgs.Get(0)
		} else {
			_hx_tmp_0 = nil
		}

		node.Hx_Field_targetExpr = _hx_tmp_0
		if callArgs.Len() > 1 {
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_1
			_hx_tmp_1 = callArgs
			node.Hx_Field_args = (Hx_Array_Slice[*Hx_Obj_loreline_nexpr](_hx_tmp_1, 1, HxNullable[int32]{}))
		}

		return func() func(node *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_ninsertion {
			var _hx_closure_src_716 func(node any) any
			_ = _hx_closure_src_716
			_hx_closure_src_716 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_ninsertion {
				var hx_dyncast_315 any
				_ = hx_dyncast_315
				hx_dyncast_315 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_716(((any)(node))), "Hx_Obj_loreline_ninsertion")
				var _hx_tmp_2 *Hx_Obj_loreline_ninsertion
				_ = _hx_tmp_2
				if hx_dyncast_315 != nil {
					_hx_tmp_2 = hx_dyncast_315.(interface {
						Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
					}).Hx_As_Hx_Obj_loreline_ninsertion()
				} else {
					_hx_tmp_2 = nil
				}

				return _hx_tmp_2
			}
		}()(node)
	}

	var target string
	_ = target
	target = this.VTable.Hx_Field_expectIdentifier()
	var targetPos *Hx_Obj_loreline_position
	_ = targetPos
	targetPos = this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos
	var args HxArray[*Hx_Obj_loreline_nexpr]
	_ = args
	var _hx_tmp_3 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Enum_loreline_tokentype_LParen{}

	if this.VTable.Hx_Field_match(_hx_tmp_3, nil) {
		args = this.VTable.Hx_Field_parseCallArguments()
	}

	var node *Hx_Obj_loreline_ninsertion
	_ = node
	node = Hx_Obj_loreline_ninsertion_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), target, targetPos, nil, nil)
	node.Hx_Field_args = args
	return func() func(node *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_ninsertion {
		var _hx_closure_src_717 func(node any) any
		_ = _hx_closure_src_717
		_hx_closure_src_717 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_ninsertion {
			var hx_dyncast_316 any
			_ = hx_dyncast_316
			hx_dyncast_316 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_717(((any)(node))), "Hx_Obj_loreline_ninsertion")
			var _hx_tmp_4 *Hx_Obj_loreline_ninsertion
			_ = _hx_tmp_4
			if hx_dyncast_316 != nil {
				_hx_tmp_4 = hx_dyncast_316.(interface {
					Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
				}).Hx_As_Hx_Obj_loreline_ninsertion()
			} else {
				_hx_tmp_4 = nil
			}

			return _hx_tmp_4
		}
	}()(node)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseBeatCall() *Hx_Obj_loreline_nbeatcall {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_KwBeat{})
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_LParen{})
	var callArgs HxArray[*Hx_Obj_loreline_nexpr]
	_ = callArgs
	callArgs = this.VTable.Hx_Field_parseCallArguments()
	if callArgs.Len() == 0 {
		this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("beat() requires a target expression", startPos))
	}

	var _hx_tmp_0 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_0
	if callArgs.Len() > 0 {
		_hx_tmp_0 = callArgs.Get(0)
	} else {
		_hx_tmp_0 = nil
	}

	var node *Hx_Obj_loreline_nbeatcall
	_ = node
	node = Hx_Obj_loreline_nbeatcall_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), _hx_tmp_0, startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), nil, nil)
	if callArgs.Len() > 1 {
		var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _hx_tmp_1
		_hx_tmp_1 = callArgs
		node.Hx_Field_args = (Hx_Array_Slice[*Hx_Obj_loreline_nexpr](_hx_tmp_1, 1, HxNullable[int32]{}))
	}

	return func() func(node *Hx_Obj_loreline_nbeatcall) *Hx_Obj_loreline_nbeatcall {
		var _hx_closure_src_718 func(node any) any
		_ = _hx_closure_src_718
		_hx_closure_src_718 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nbeatcall) *Hx_Obj_loreline_nbeatcall {
			var hx_dyncast_317 any
			_ = hx_dyncast_317
			hx_dyncast_317 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_718(((any)(node))), "Hx_Obj_loreline_nbeatcall")
			var _hx_tmp_2 *Hx_Obj_loreline_nbeatcall
			_ = _hx_tmp_2
			if hx_dyncast_317 != nil {
				_hx_tmp_2 = hx_dyncast_317.(interface {
					Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
				}).Hx_As_Hx_Obj_loreline_nbeatcall()
			} else {
				_hx_tmp_2 = nil
			}

			return _hx_tmp_2
		}
	}()(node)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseFunction() *Hx_Obj_loreline_nfunctiondecl {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	if _g.Hx_Field_enumIndex() == 15 {
		var name HxNullable[string]
		_ = name
		name = _g.(Hx_Enum_loreline_tokentype_Function).Hx_Field_name
		var args HxArray[string]
		_ = args
		args = _g.(Hx_Enum_loreline_tokentype_Function).Hx_Field_args
		var code string
		_ = code
		code = _g.(Hx_Enum_loreline_tokentype_Function).Hx_Field_code
		var external bool
		_ = external
		external = _g.(Hx_Enum_loreline_tokentype_Function).Hx_Field_external
		this.VTable.Hx_Field_advance(nil)
		return func() func(node *Hx_Obj_loreline_nfunctiondecl) *Hx_Obj_loreline_nfunctiondecl {
			var _hx_closure_src_719 func(node any) any
			_ = _hx_closure_src_719
			_hx_closure_src_719 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nfunctiondecl) *Hx_Obj_loreline_nfunctiondecl {
				var hx_dyncast_318 any
				_ = hx_dyncast_318
				hx_dyncast_318 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_719(((any)(node))), "Hx_Obj_loreline_nfunctiondecl")
				var _hx_tmp_0 *Hx_Obj_loreline_nfunctiondecl
				_ = _hx_tmp_0
				if hx_dyncast_318 != nil {
					_hx_tmp_0 = hx_dyncast_318.(interface {
						Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
					}).Hx_As_Hx_Obj_loreline_nfunctiondecl()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(Hx_Obj_loreline_nfunctiondecl_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(3))), startPos, name, (Hx_Array_Concat[string](HxMakeArray[string](), args)), code, external, nil, nil))
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Unexpected token in expression", this.VTable.Hx_Field_currentPos())))))
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseExpression() *Hx_Obj_loreline_nexpr {
	{
		var hx_try_return *Hx_Obj_loreline_nexpr
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							if Hx_Field_std_isOfType(e, Hx_Obj_loreline_parseerror_RTTI) {
								var hx_dyncast_319 any
								_ = hx_dyncast_319
								hx_dyncast_319 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_parseerror")
								var _hx_tmp_6 *Hx_Obj_loreline_parseerror
								_ = _hx_tmp_6
								if hx_dyncast_319 != nil {
									_hx_tmp_6 = hx_dyncast_319.(interface {
										Hx_As_Hx_Obj_loreline_parseerror() *Hx_Obj_loreline_parseerror
									}).Hx_As_Hx_Obj_loreline_parseerror()
								} else {
									_hx_tmp_6 = nil
								}

								this.VTable.Hx_Field_addError(_hx_tmp_6)
							}

							hx_try_state = 1
							hx_try_return = (func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var expr *Hx_Obj_loreline_nexpr
				_ = expr
				expr = this.VTable.Hx_Field_parseTernary()
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				var _hx_tmp_2 bool
				_ = _hx_tmp_2
				var _hx_tmp_3 bool
				_ = _hx_tmp_3
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpAssign{})) {
					_hx_tmp_4 = true
				} else {
					_hx_tmp_4 = this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpPlusAssign{})
				}

				if _hx_tmp_4 {
					_hx_tmp_3 = true
				} else {
					_hx_tmp_3 = this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpMinusAssign{})
				}

				if _hx_tmp_3 {
					_hx_tmp_2 = true
				} else {
					_hx_tmp_2 = this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpMultiplyAssign{})
				}

				if _hx_tmp_2 {
					_hx_tmp_1 = true
				} else {
					_hx_tmp_1 = this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpDivideAssign{})
				}

				if _hx_tmp_1 {
					_hx_tmp_0 = true
				} else {
					_hx_tmp_0 = this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_OpUnquotedAssign{})
				}

				if _hx_tmp_0 {
					var op Hx_Enum_loreline_tokentype
					_ = op
					op = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
					this.VTable.Hx_Field_advance(nil)
					var assignment *Hx_Obj_loreline_nassign
					_ = assignment
					assignment = func() func(node *Hx_Obj_loreline_nassign) *Hx_Obj_loreline_nassign {
						var _hx_closure_src_720 func(node any) any
						_ = _hx_closure_src_720
						_hx_closure_src_720 = this.VTable.Hx_Field_attachComments
						return func(node *Hx_Obj_loreline_nassign) *Hx_Obj_loreline_nassign {
							var hx_dyncast_320 any
							_ = hx_dyncast_320
							hx_dyncast_320 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_720(((any)(node))), "Hx_Obj_loreline_nassign")
							var _hx_tmp_5 *Hx_Obj_loreline_nassign
							_ = _hx_tmp_5
							if hx_dyncast_320 != nil {
								_hx_tmp_5 = hx_dyncast_320.(interface {
									Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
								}).Hx_As_Hx_Obj_loreline_nassign()
							} else {
								_hx_tmp_5 = nil
							}

							return _hx_tmp_5
						}
					}()(Hx_Obj_loreline_nassign_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
					assignment.Hx_Field_value = this.VTable.Hx_Field_parseExpression()
					assignment.Hx_Field_pos = assignment.Hx_Field_pos.VTable.Hx_Field_extendedTo(assignment.Hx_Field_value.Hx_Field_pos)
					hx_try_state = 1
					hx_try_return = (func() *Hx_Obj_loreline_nexpr {
						_hx_up := assignment
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_nexpr
						}
						return nil
					})()
					return
				}

				hx_try_state = 1
				hx_try_return = expr
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseTernary() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseLogicalOr()
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Enum_loreline_tokentype_Question{}

	if this.VTable.Hx_Field_match(_hx_tmp_0, nil) {
		var trueExpr *Hx_Obj_loreline_nexpr
		_ = trueExpr
		trueExpr = this.VTable.Hx_Field_parseTernary()
		this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_Colon{})
		var falseExpr *Hx_Obj_loreline_nexpr
		_ = falseExpr
		falseExpr = this.VTable.Hx_Field_parseTernary()
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := func() func(node *Hx_Obj_loreline_nternary) *Hx_Obj_loreline_nternary {
				var _hx_closure_src_722 func(node any) any
				_ = _hx_closure_src_722
				_hx_closure_src_722 = this.VTable.Hx_Field_attachComments
				return func(node *Hx_Obj_loreline_nternary) *Hx_Obj_loreline_nternary {
					var hx_dyncast_322 any
					_ = hx_dyncast_322
					hx_dyncast_322 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_722(((any)(node))), "Hx_Obj_loreline_nternary")
					var _hx_tmp_1 *Hx_Obj_loreline_nternary
					_ = _hx_tmp_1
					if hx_dyncast_322 != nil {
						_hx_tmp_1 = hx_dyncast_322.(interface {
							Hx_As_Hx_Obj_loreline_nternary() *Hx_Obj_loreline_nternary
						}).Hx_As_Hx_Obj_loreline_nternary()
					} else {
						_hx_tmp_1 = nil
					}

					return _hx_tmp_1
				}
			}()(Hx_Obj_loreline_nternary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, trueExpr, falseExpr, nil, nil))
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
		expr.Hx_Field_pos = expr.Hx_Field_pos.VTable.Hx_Field_extendedTo(falseExpr.Hx_Field_pos)
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseLogicalOr() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseLogicalAnd()
	for {
		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_OpOr{false}

		if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_723 func(node any) any
			_ = _hx_closure_src_723
			_hx_closure_src_723 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_323 any
				_ = hx_dyncast_323
				hx_dyncast_323 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_723(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_1 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_1
				if hx_dyncast_323 != nil {
					_hx_tmp_1 = hx_dyncast_323.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_1 = nil
				}

				return _hx_tmp_1
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseLogicalAnd()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseLogicalAnd() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseEquality()
	for {
		var _hx_tmp_0 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_OpAnd{false}

		if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_724 func(node any) any
			_ = _hx_closure_src_724
			_hx_closure_src_724 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_324 any
				_ = hx_dyncast_324
				hx_dyncast_324 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_724(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_1 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_1
				if hx_dyncast_324 != nil {
					_hx_tmp_1 = hx_dyncast_324.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_1 = nil
				}

				return _hx_tmp_1
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseEquality()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseEquality() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseComparison()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Enum_loreline_tokentype_OpEquals{false}

		if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_2 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpNotEquals{false}

			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_3, nil)
		}

		if !(_hx_tmp_0) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_725 func(node any) any
			_ = _hx_closure_src_725
			_hx_closure_src_725 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_325 any
				_ = hx_dyncast_325
				hx_dyncast_325 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_725(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_4 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_4
				if hx_dyncast_325 != nil {
					_hx_tmp_4 = hx_dyncast_325.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_4 = nil
				}

				return _hx_tmp_4
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseComparison()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseComparison() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseAdditive()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_3
		_hx_tmp_3 = Hx_Enum_loreline_tokentype_OpGreater{}

		if this.VTable.Hx_Field_match(_hx_tmp_3, nil) {
			_hx_tmp_2 = true
		} else {
			var _hx_tmp_4 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_4
			_hx_tmp_4 = Hx_Enum_loreline_tokentype_OpGreaterEq{}

			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = _hx_tmp_4
			_hx_tmp_2 = this.VTable.Hx_Field_match(_hx_tmp_5, nil)
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_6 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Enum_loreline_tokentype_OpLess{}

			var _hx_tmp_7 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_1 = this.VTable.Hx_Field_match(_hx_tmp_7, nil)
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_8 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_8
			_hx_tmp_8 = Hx_Enum_loreline_tokentype_OpLessEq{}

			var _hx_tmp_9 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_9
			_hx_tmp_9 = _hx_tmp_8
			_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_9, nil)
		}

		if !(_hx_tmp_0) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_726 func(node any) any
			_ = _hx_closure_src_726
			_hx_closure_src_726 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_326 any
				_ = hx_dyncast_326
				hx_dyncast_326 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_726(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_10 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_10
				if hx_dyncast_326 != nil {
					_hx_tmp_10 = hx_dyncast_326.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_10 = nil
				}

				return _hx_tmp_10
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseAdditive()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseAdditive() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseMultiplicative()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Enum_loreline_tokentype_OpPlus{}

		if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_2 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpMinus{}

			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_3, nil)
		}

		if !(_hx_tmp_0) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_727 func(node any) any
			_ = _hx_closure_src_727
			_hx_closure_src_727 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_327 any
				_ = hx_dyncast_327
				hx_dyncast_327 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_727(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_4 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_4
				if hx_dyncast_327 != nil {
					_hx_tmp_4 = hx_dyncast_327.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_4 = nil
				}

				return _hx_tmp_4
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseMultiplicative()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseMultiplicative() *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = this.VTable.Hx_Field_parseUnary()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpMultiply{}

		if this.VTable.Hx_Field_match(_hx_tmp_2, nil) {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_OpDivide{}

			var _hx_tmp_4 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_1 = this.VTable.Hx_Field_match(_hx_tmp_4, nil)
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			var _hx_tmp_5 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_5
			_hx_tmp_5 = Hx_Enum_loreline_tokentype_OpModulo{}

			var _hx_tmp_6 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_6
			_hx_tmp_6 = _hx_tmp_5
			_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_6, nil)
		}

		if !(_hx_tmp_0) {
			break
		}

		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var binary *Hx_Obj_loreline_nbinary
		_ = binary
		binary = func() func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
			var _hx_closure_src_728 func(node any) any
			_ = _hx_closure_src_728
			_hx_closure_src_728 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nbinary) *Hx_Obj_loreline_nbinary {
				var hx_dyncast_328 any
				_ = hx_dyncast_328
				hx_dyncast_328 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_728(((any)(node))), "Hx_Obj_loreline_nbinary")
				var _hx_tmp_7 *Hx_Obj_loreline_nbinary
				_ = _hx_tmp_7
				if hx_dyncast_328 != nil {
					_hx_tmp_7 = hx_dyncast_328.(interface {
						Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
					}).Hx_As_Hx_Obj_loreline_nbinary()
				} else {
					_hx_tmp_7 = nil
				}

				return _hx_tmp_7
			}
		}()(Hx_Obj_loreline_nbinary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), expr.Hx_Field_pos, expr, op, nil, nil, nil))
		binary.Hx_Field_right = this.VTable.Hx_Field_parseUnary()
		binary.Hx_Field_pos = binary.Hx_Field_pos.VTable.Hx_Field_extendedTo(binary.Hx_Field_right.Hx_Field_pos)
		expr = (func() *Hx_Obj_loreline_nexpr {
			_hx_up := binary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseUnary() *Hx_Obj_loreline_nexpr {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Enum_loreline_tokentype_OpNot{false}

	if this.VTable.Hx_Field_match(_hx_tmp_1, nil) {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpMinus{}

		var _hx_tmp_3 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = this.VTable.Hx_Field_match(_hx_tmp_3, nil)
	}

	if _hx_tmp_0 {
		var op Hx_Enum_loreline_tokentype
		_ = op
		op = this.VTable.Hx_Field_previous().Hx_Field_type
		var unary *Hx_Obj_loreline_nunary
		_ = unary
		unary = func() func(node *Hx_Obj_loreline_nunary) *Hx_Obj_loreline_nunary {
			var _hx_closure_src_729 func(node any) any
			_ = _hx_closure_src_729
			_hx_closure_src_729 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nunary) *Hx_Obj_loreline_nunary {
				var hx_dyncast_329 any
				_ = hx_dyncast_329
				hx_dyncast_329 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_729(((any)(node))), "Hx_Obj_loreline_nunary")
				var _hx_tmp_4 *Hx_Obj_loreline_nunary
				_ = _hx_tmp_4
				if hx_dyncast_329 != nil {
					_hx_tmp_4 = hx_dyncast_329.(interface {
						Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
					}).Hx_As_Hx_Obj_loreline_nunary()
				} else {
					_hx_tmp_4 = nil
				}

				return _hx_tmp_4
			}
		}()(Hx_Obj_loreline_nunary_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_previous().Hx_Field_pos, op, nil, nil, nil))
		unary.Hx_Field_operand = this.VTable.Hx_Field_parseUnary()
		unary.Hx_Field_pos = unary.Hx_Field_pos.VTable.Hx_Field_extendedTo(unary.Hx_Field_operand.Hx_Field_pos)
		return (func() *Hx_Obj_loreline_nexpr {
			_hx_up := unary
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()
	}

	return this.VTable.Hx_Field_parsePrimary()
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parsePrimary() *Hx_Obj_loreline_nexpr {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	switch _g.Hx_Field_enumIndex() {
	case 16:
		{
			var _hx_tmp_0 *Hx_Obj_loreline_position
			_ = _hx_tmp_0
			_hx_tmp_0 = startPos
			return this.VTable.Hx_Field_parsePostfix(_hx_tmp_0, (func() *Hx_Obj_loreline_nexpr {
				_hx_up := this.VTable.Hx_Field_parseStringLiteral()
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_nexpr
				}
				return nil
			})())
		}

	case 17:
		{
			var n float64
			_ = n
			n = _g.(Hx_Enum_loreline_tokentype_LNumber).Hx_Field_n
			this.VTable.Hx_Field_advance(nil)
			return (func() *Hx_Obj_loreline_nexpr {
				_hx_up := func() func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
					var _hx_closure_src_733 func(node any) any
					_ = _hx_closure_src_733
					_hx_closure_src_733 = this.VTable.Hx_Field_attachComments
					return func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
						var hx_dyncast_333 any
						_ = hx_dyncast_333
						hx_dyncast_333 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_733(((any)(node))), "Hx_Obj_loreline_nliteral")
						var _hx_tmp_1 *Hx_Obj_loreline_nliteral
						_ = _hx_tmp_1
						if hx_dyncast_333 != nil {
							_hx_tmp_1 = hx_dyncast_333.(interface {
								Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
							}).Hx_As_Hx_Obj_loreline_nliteral()
						} else {
							_hx_tmp_1 = nil
						}

						return _hx_tmp_1
					}
				}()(Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, n, Hx_Enum_loreline_literaltype_Number{}, nil, nil))
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_nexpr
				}
				return nil
			})()
		}

	case 18:
		{
			var b bool
			_ = b
			b = _g.(Hx_Enum_loreline_tokentype_LBoolean).Hx_Field_b
			this.VTable.Hx_Field_advance(nil)
			return (func() *Hx_Obj_loreline_nexpr {
				_hx_up := func() func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
					var _hx_closure_src_734 func(node any) any
					_ = _hx_closure_src_734
					_hx_closure_src_734 = this.VTable.Hx_Field_attachComments
					return func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
						var hx_dyncast_334 any
						_ = hx_dyncast_334
						hx_dyncast_334 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_734(((any)(node))), "Hx_Obj_loreline_nliteral")
						var _hx_tmp_2 *Hx_Obj_loreline_nliteral
						_ = _hx_tmp_2
						if hx_dyncast_334 != nil {
							_hx_tmp_2 = hx_dyncast_334.(interface {
								Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
							}).Hx_As_Hx_Obj_loreline_nliteral()
						} else {
							_hx_tmp_2 = nil
						}

						return _hx_tmp_2
					}
				}()(Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, b, Hx_Enum_loreline_literaltype_Boolean{}, nil, nil))
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_nexpr
				}
				return nil
			})()
		}

	case 19:
		{
			this.VTable.Hx_Field_advance(nil)
			return (func() *Hx_Obj_loreline_nexpr {
				_hx_up := func() func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
					var _hx_closure_src_735 func(node any) any
					_ = _hx_closure_src_735
					_hx_closure_src_735 = this.VTable.Hx_Field_attachComments
					return func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
						var hx_dyncast_335 any
						_ = hx_dyncast_335
						hx_dyncast_335 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_735(((any)(node))), "Hx_Obj_loreline_nliteral")
						var _hx_tmp_3 *Hx_Obj_loreline_nliteral
						_ = _hx_tmp_3
						if hx_dyncast_335 != nil {
							_hx_tmp_3 = hx_dyncast_335.(interface {
								Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
							}).Hx_As_Hx_Obj_loreline_nliteral()
						} else {
							_hx_tmp_3 = nil
						}

						return _hx_tmp_3
					}
				}()(Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil))
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_nexpr
				}
				return nil
			})()
		}

	case 20:
		{
			var name string
			_ = name
			name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
			if (this.VTable.Hx_Field_peek().Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) {
				var fields HxArray[*Hx_Obj_loreline_nobjectfield]
				_ = fields
				fields = HxMakeArray[*Hx_Obj_loreline_nobjectfield](this.VTable.Hx_Field_parseObjectField())
				return (func() *Hx_Obj_loreline_nexpr {
					_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos), fields, Hx_Enum_loreline_literaltype_Object{((int32)(0))}, nil, nil)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_nexpr
					}
					return nil
				})()
			} else {
				this.VTable.Hx_Field_advance(nil)
				return this.VTable.Hx_Field_parseIdentifierExpression(startPos, name)
			}
		}

	case 46:
		{
			var _hx_tmp_4 *Hx_Obj_loreline_position
			_ = _hx_tmp_4
			_hx_tmp_4 = startPos
			return this.VTable.Hx_Field_parsePostfix(_hx_tmp_4, this.VTable.Hx_Field_parseObjectLiteral())
		}

	case 48:
		{
			this.VTable.Hx_Field_advance(nil)
			var expr *Hx_Obj_loreline_nexpr
			_ = expr
			expr = this.VTable.Hx_Field_parseExpression()
			this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_RParen{})
			expr.Hx_Field_parens++
			return this.VTable.Hx_Field_parsePostfix(startPos, expr)
		}

	case 50:
		{
			var _hx_tmp_5 *Hx_Obj_loreline_position
			_ = _hx_tmp_5
			_hx_tmp_5 = startPos
			return this.VTable.Hx_Field_parsePostfix(_hx_tmp_5, this.VTable.Hx_Field_parseArrayLiteral())
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Unexpected token in expression", this.VTable.Hx_Field_currentPos())))))
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseStringLiteral() *Hx_Obj_loreline_nstringliteral {
	var stringLiteralPos *Hx_Obj_loreline_position
	_ = stringLiteralPos
	stringLiteralPos = this.VTable.Hx_Field_currentPos()
	var parts HxArray[*Hx_Obj_loreline_nstringpart]
	_ = parts
	parts = HxMakeArray[*Hx_Obj_loreline_nstringpart]()
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	if _g.Hx_Field_enumIndex() == 16 {
		var quotes int32
		_ = quotes
		quotes = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
		var content string
		_ = content
		content = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
		var attachments HxArray[Hx_Enum_loreline_lstringattachment]
		_ = attachments
		attachments = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_attachments
		var _hx_tmp_0 *Hx_Obj_loreline_position
		_ = _hx_tmp_0
		if quotes != ((int32)(0)) {
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = content
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = (stringLiteralPos.Hx_Field_length - 2)
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_1
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_2
			_hx_tmp_0 = stringLiteralPos.VTable.Hx_Field_withOffset(_hx_tmp_3, 1, _hx_tmp_4, nil)
		} else {
			_hx_tmp_0 = stringLiteralPos
		}

		var startPos *Hx_Obj_loreline_position
		_ = startPos
		startPos = _hx_tmp_0
		var currentPos int32
		_ = currentPos
		currentPos = 0
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if attachments == nil {
			_hx_tmp_5 = true
		} else {
			_hx_tmp_5 = (attachments.Len() == 0)
		}

		if _hx_tmp_5 {
			var partPos *Hx_Obj_loreline_position
			_ = partPos
			partPos = this.VTable.Hx_Field_makeStringPartPosition(startPos, content, 0)
			partPos.Hx_Field_length = startPos.Hx_Field_length
			var literalId Hx_Typedef_haxe__int64___int64
			_ = literalId
			literalId = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
			var partId Hx_Typedef_haxe__int64___int64
			_ = partId
			partId = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
			var _hx_tmp_6 HxArray[*Hx_Obj_loreline_nstringpart]
			_ = _hx_tmp_6
			_hx_tmp_6 = parts
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_6, Hx_Obj_loreline_nstringpart_CreateInstance(partId, partPos, Hx_Enum_loreline_stringparttype_Raw{content}, nil, nil)))
			this.VTable.Hx_Field_advance(nil)
			return func() func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
				var _hx_closure_src_736 func(node any) any
				_ = _hx_closure_src_736
				_hx_closure_src_736 = this.VTable.Hx_Field_attachComments
				return func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
					var hx_dyncast_336 any
					_ = hx_dyncast_336
					hx_dyncast_336 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_736(((any)(node))), "Hx_Obj_loreline_nstringliteral")
					var _hx_tmp_7 *Hx_Obj_loreline_nstringliteral
					_ = _hx_tmp_7
					if hx_dyncast_336 != nil {
						_hx_tmp_7 = hx_dyncast_336.(interface {
							Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
						}).Hx_As_Hx_Obj_loreline_nstringliteral()
					} else {
						_hx_tmp_7 = nil
					}

					return _hx_tmp_7
				}
			}()(Hx_Obj_loreline_nstringliteral_CreateInstance(literalId, stringLiteralPos, quotes, parts, nil, nil))
		}

		var stringLiteral *Hx_Obj_loreline_nstringliteral
		_ = stringLiteral
		stringLiteral = func() func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
			var _hx_closure_src_737 func(node any) any
			_ = _hx_closure_src_737
			_hx_closure_src_737 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
				var hx_dyncast_337 any
				_ = hx_dyncast_337
				hx_dyncast_337 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_737(((any)(node))), "Hx_Obj_loreline_nstringliteral")
				var _hx_tmp_8 *Hx_Obj_loreline_nstringliteral
				_ = _hx_tmp_8
				if hx_dyncast_337 != nil {
					_hx_tmp_8 = hx_dyncast_337.(interface {
						Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
					}).Hx_As_Hx_Obj_loreline_nstringliteral()
				} else {
					_hx_tmp_8 = nil
				}

				return _hx_tmp_8
			}
		}()(Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), stringLiteralPos, quotes, parts, nil, nil))
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 int32
			_ = _g2
			_g2 = attachments.Len()
			for _g1 < _g2 {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g1
				_g1 = (_g1 + 1)
				var i int32
				_ = i
				i = _hx_tmp_9
				{
					var tmp__g_1 Hx_Enum_loreline_lstringattachment
					_ = tmp__g_1
					tmp__g_1 = attachments.Get(i)
					switch tmp__g_1.Hx_Field_enumIndex() {
					case 0:
						{
							var braces bool
							_ = braces
							braces = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_braces
							var inTag bool
							_ = inTag
							inTag = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_inTag
							var tokens Hx_Typedef_loreline_tokens
							_ = tokens
							tokens = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_expr
							var start int32
							_ = start
							start = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
							var length int32
							_ = length
							length = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_length
							{
								if currentPos >= (start + length) {
									continue
								}

								if start > currentPos {
									var partPos *Hx_Obj_loreline_position
									_ = partPos
									partPos = this.VTable.Hx_Field_makeStringPartPosition(startPos, content, currentPos)
									partPos.Hx_Field_length = (start - currentPos)
									var _hx_reserved_len HxNullable[int32]
									_ = _hx_reserved_len
									_hx_reserved_len = HxNullable[int32]{Value: (start - currentPos), Valid: true}

									var _hx_tmp_10 HxArray[*Hx_Obj_loreline_nstringpart]
									_ = _hx_tmp_10
									_hx_tmp_10 = parts
									var _hx_tmp_11 string
									_ = _hx_tmp_11
									if (_hx_reserved_len).Valid != false {
										_hx_tmp_11 = (Hx_Field_go_haxe_hxstring_substr(content, currentPos, _hx_reserved_len))
									} else {
										var _hx_tmp_12 string
										_ = _hx_tmp_12
										_hx_tmp_12 = content
										var _hx_tmp_13 int32
										_ = _hx_tmp_13
										_hx_tmp_13 = currentPos
										var _hx_tmp_14 string
										_ = _hx_tmp_14
										_hx_tmp_14 = _hx_tmp_12
										var _hx_tmp_15 int32
										_ = _hx_tmp_15
										_hx_tmp_15 = _hx_tmp_13
										_hx_tmp_11 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_14, _hx_tmp_15, HxNullable[int32]{}))
									}

									_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_10, Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_11}, nil, nil)))
								}

								var _hx_tmp_16 HxArray[*Hx_Obj_loreline_nstringpart]
								_ = _hx_tmp_16
								_hx_tmp_16 = parts
								_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_16, this.VTable.Hx_Field_parseStringInterpolation(braces, inTag, tokens, start, length, content)))
								currentPos = (start + length)
							}
						}

					case 1:
						{
							var closing bool
							_ = closing
							closing = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_closing
							var start int32
							_ = start
							start = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_start
							var length int32
							_ = length
							length = tmp__g_1.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_length
							{
								if currentPos >= (start + length) {
									continue
								}

								if start > currentPos {
									var partPos *Hx_Obj_loreline_position
									_ = partPos
									partPos = this.VTable.Hx_Field_makeStringPartPosition(startPos, content, currentPos)
									partPos.Hx_Field_length = (start - currentPos)
									var _hx_reserved_len HxNullable[int32]
									_ = _hx_reserved_len
									_hx_reserved_len = HxNullable[int32]{Value: (start - currentPos), Valid: true}

									var _hx_tmp_17 HxArray[*Hx_Obj_loreline_nstringpart]
									_ = _hx_tmp_17
									_hx_tmp_17 = parts
									var _hx_tmp_18 string
									_ = _hx_tmp_18
									if (_hx_reserved_len).Valid != false {
										_hx_tmp_18 = (Hx_Field_go_haxe_hxstring_substr(content, currentPos, _hx_reserved_len))
									} else {
										var _hx_tmp_19 string
										_ = _hx_tmp_19
										_hx_tmp_19 = content
										var _hx_tmp_20 int32
										_ = _hx_tmp_20
										_hx_tmp_20 = currentPos
										var _hx_tmp_21 string
										_ = _hx_tmp_21
										_hx_tmp_21 = _hx_tmp_19
										var _hx_tmp_22 int32
										_ = _hx_tmp_22
										_hx_tmp_22 = _hx_tmp_20
										_hx_tmp_18 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_21, _hx_tmp_22, HxNullable[int32]{}))
									}

									_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_17, Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_18}, nil, nil)))
								}

								var _hx_tmp_23 HxArray[*Hx_Obj_loreline_nstringpart]
								_ = _hx_tmp_23
								_hx_tmp_23 = parts
								_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_23, this.VTable.Hx_Field_parseStringTag(closing, start, length, content, quotes, attachments)))
								currentPos = (start + length)
							}
						}

					}
				}
			}
		}

		var _hx_tmp_24 int32
		_ = _hx_tmp_24
		_hx_tmp_24 = currentPos
		if _hx_tmp_24 < int32(utf8.RuneCountInString(content)) {
			var partPos *Hx_Obj_loreline_position
			_ = partPos
			partPos = this.VTable.Hx_Field_makeStringPartPosition(startPos, content, currentPos)
			partPos.Hx_Field_length = (int32(utf8.RuneCountInString(content)) - currentPos)
			var _hx_reserved_len HxNullable[int32]
			_ = _hx_reserved_len
			_hx_reserved_len = HxNullable[int32]{}

			var _hx_tmp_25 HxArray[*Hx_Obj_loreline_nstringpart]
			_ = _hx_tmp_25
			_hx_tmp_25 = parts
			var _hx_tmp_26 string
			_ = _hx_tmp_26
			if (_hx_reserved_len).Valid != false {
				_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(content, currentPos, _hx_reserved_len))
			} else {
				var _hx_tmp_27 string
				_ = _hx_tmp_27
				_hx_tmp_27 = content
				var _hx_tmp_28 int32
				_ = _hx_tmp_28
				_hx_tmp_28 = currentPos
				var _hx_tmp_29 string
				_ = _hx_tmp_29
				_hx_tmp_29 = _hx_tmp_27
				var _hx_tmp_30 int32
				_ = _hx_tmp_30
				_hx_tmp_30 = _hx_tmp_28
				_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_29, _hx_tmp_30, HxNullable[int32]{}))
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_25, Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_26}, nil, nil)))
		}

		this.VTable.Hx_Field_advance(nil)
		if this.Hx_Field_pendingComments != nil {
			var remaining HxArray[*Hx_Obj_loreline_comment]
			_ = remaining
			{
				var tmp__g_1 int32
				_ = tmp__g_1
				tmp__g_1 = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = this.Hx_Field_pendingComments
				for {
					var _hx_tmp_31 int32
					_ = _hx_tmp_31
					_hx_tmp_31 = tmp__g_1
					if !(_hx_tmp_31 < _g1.Len()) {
						break
					}

					var c *Hx_Obj_loreline_comment
					_ = c
					c = _g1.Get(tmp__g_1)
					tmp__g_1++
					if c.Hx_Field_pos.Hx_Field_line == stringLiteralPos.Hx_Field_line {
						if stringLiteral.Hx_Field_trailingComments == nil {
							stringLiteral.Hx_Field_trailingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
						}

						_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](stringLiteral.Hx_Field_trailingComments, c))
					} else {
						if remaining == nil {
							remaining = HxMakeArray[*Hx_Obj_loreline_comment]()
						}

						_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](remaining, c))
					}
				}
			}

			this.Hx_Field_pendingComments = remaining
		}

		return stringLiteral
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Expected text, got ", this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos())))))
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_makeStringPartPosition(stringStart *Hx_Obj_loreline_position, content string, offset int32) *Hx_Obj_loreline_position {
	var line int32
	_ = line
	line = stringStart.Hx_Field_line
	var column int32
	_ = column
	column = stringStart.Hx_Field_column
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = offset
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(content, i))
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			if (c).Valid == false {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = (c).Value
			}

			if (_hx_tmp_1) == 10 {
				line++
				column = 1
			} else {
				column++
			}
		}
	}

	return Hx_Obj_loreline_position_CreateInstance(line, column, (stringStart.Hx_Field_offset + offset), nil)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_makeAccess(pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, name string, namePos *Hx_Obj_loreline_position) *Hx_Obj_loreline_naccess {
	if name != HxStringNull {
		if int32(utf8.RuneCountInString(name)) == 0 {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if this.Hx_Field_errors == nil {
				_hx_tmp_0 = true
			} else {
				_hx_tmp_0 = (this.Hx_Field_errors.Len() == 0)
			}

			if _hx_tmp_0 {
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				if name != HxStringNull {
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("'", name), "'")
				} else {
					_hx_tmp_1 = "null"
				}

				this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid access: ", (_hx_tmp_1)), pos))
			}
		}

		if target != nil {
			if namePos != nil {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = namePos.Hx_Field_offset
				pos = Hx_Obj_loreline_position_CreateInstance(target.Hx_Field_pos.Hx_Field_line, target.Hx_Field_pos.Hx_Field_column, target.Hx_Field_pos.Hx_Field_offset, ((_hx_tmp_2 + int32(utf8.RuneCountInString(name))) - target.Hx_Field_pos.Hx_Field_offset))
			} else {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Invalid access (missing name position)", pos)))))
			}
		} else {
			if pos.Hx_Field_length == 0 {
				var _hx_tmp_3 HxNullable[int32]
				_ = _hx_tmp_3
				if namePos != nil {
					_hx_tmp_3 = HxNullable[int32]{Value: namePos.Hx_Field_length, Valid: true}
				} else {
					_hx_tmp_3 = HxNullable[int32]{}
				}

				var length HxNullable[int32]
				_ = length
				length = _hx_tmp_3
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (length).Valid != false {
					_hx_tmp_4 = (length).Value
				} else {
					_hx_tmp_4 = int32(utf8.RuneCountInString(name))
				}

				pos = Hx_Obj_loreline_position_CreateInstance(pos.Hx_Field_line, pos.Hx_Field_column, pos.Hx_Field_offset, _hx_tmp_4)
			}
		}
	}

	return Hx_Obj_loreline_naccess_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pos, target, name, nil, nil)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseStringInterpolation(braces bool, inTag bool, tokens Hx_Typedef_loreline_tokens, start int32, length int32, content string) *Hx_Obj_loreline_nstringpart {
	var tmp *Hx_Obj_loreline_token
	_ = tmp
	tmp = tokens.Get(0)
	var _hx_tmp_0 *Hx_Obj_loreline_position
	_ = _hx_tmp_0
	if tmp != nil {
		_hx_tmp_0 = tmp.Hx_Field_pos
	} else {
		_hx_tmp_0 = nil
	}

	var stringStart *Hx_Obj_loreline_position
	_ = stringStart
	stringStart = _hx_tmp_0
	var _hx_tmp_1 *Hx_Obj_loreline_position
	_ = _hx_tmp_1
	if stringStart != nil {
		_hx_tmp_1 = stringStart
	} else {
		_hx_tmp_1 = this.VTable.Hx_Field_currentPos()
	}

	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = this.VTable.Hx_Field_makeStringPartPosition(_hx_tmp_1, content, start)
	pos.Hx_Field_length = length
	if tokens.Len() == 0 {
		var tmp_tmp_1 *Hx_Obj_loreline_token
		_ = tmp_tmp_1
		tmp_tmp_1 = tokens.Get(0)
		var _hx_tmp_2 *Hx_Obj_loreline_position
		_ = _hx_tmp_2
		if tmp_tmp_1 != nil {
			_hx_tmp_2 = tmp_tmp_1.Hx_Field_pos
		} else {
			_hx_tmp_2 = nil
		}

		var pos1 *Hx_Obj_loreline_position
		_ = pos1
		pos1 = _hx_tmp_2
		var _hx_tmp_3 *Hx_Obj_loreline_position
		_ = _hx_tmp_3
		if pos1 != nil {
			_hx_tmp_3 = pos1
		} else {
			_hx_tmp_3 = this.VTable.Hx_Field_currentPos()
		}

		this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Empty interpolation", _hx_tmp_3))
		var tmp1 Hx_Typedef_haxe__int64___int64
		_ = tmp1
		tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
		var tmp2 Hx_Typedef_haxe__int64___int64
		_ = tmp2
		tmp2 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
		var tmp3 *Hx_Obj_loreline_token
		_ = tmp3
		tmp3 = tokens.Get(0)
		var _hx_tmp_4 *Hx_Obj_loreline_position
		_ = _hx_tmp_4
		if tmp3 != nil {
			_hx_tmp_4 = tmp3.Hx_Field_pos
		} else {
			_hx_tmp_4 = nil
		}

		var pos2 *Hx_Obj_loreline_position
		_ = pos2
		pos2 = _hx_tmp_4
		var _hx_tmp_5 *Hx_Obj_loreline_position
		_ = _hx_tmp_5
		if pos2 != nil {
			_hx_tmp_5 = pos2
		} else {
			_hx_tmp_5 = this.VTable.Hx_Field_currentPos()
		}

		return Hx_Obj_loreline_nstringpart_CreateInstance(tmp1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
			_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp2, _hx_tmp_5, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_nexpr
			}
			return nil
		})()}, nil, nil)
	}

	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	if !braces {
		_hx_tmp_6 = (tokens.Len() > 0)
	} else {
		_hx_tmp_6 = false
	}

	if _hx_tmp_6 {
		var target *Hx_Obj_loreline_nexpr
		_ = target
		var i int32
		_ = i
		i = 0
		var prevIsDot bool
		_ = prevIsDot
		prevIsDot = false
		for {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = i
			if !(_hx_tmp_7 < tokens.Len()) {
				break
			}

			var token *Hx_Obj_loreline_token
			_ = token
			token = tokens.Get(i)
			i++
			{
				var _g Hx_Enum_loreline_tokentype
				_ = _g
				_g = token.Hx_Field_type
				switch _g.Hx_Field_enumIndex() {
				case 20:
					{
						var name string
						_ = name
						name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
						{
							if target == nil {
								target = (func() *Hx_Obj_loreline_nexpr {
									_hx_up := func() func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
										var _hx_closure_src_742 func(node any) any
										_ = _hx_closure_src_742
										_hx_closure_src_742 = this.VTable.Hx_Field_attachComments
										return func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
											var hx_dyncast_342 any
											_ = hx_dyncast_342
											hx_dyncast_342 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_742(((any)(node))), "Hx_Obj_loreline_naccess")
											var _hx_tmp_8 *Hx_Obj_loreline_naccess
											_ = _hx_tmp_8
											if hx_dyncast_342 != nil {
												_hx_tmp_8 = hx_dyncast_342.(interface {
													Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
												}).Hx_As_Hx_Obj_loreline_naccess()
											} else {
												_hx_tmp_8 = nil
											}

											return _hx_tmp_8
										}
									}()(this.VTable.Hx_Field_makeAccess(token.Hx_Field_pos, nil, name, nil))
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_nexpr
									}
									return nil
								})()
							} else {
								if prevIsDot {
									target = (func() *Hx_Obj_loreline_nexpr {
										_hx_up := func() func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
											var _hx_closure_src_743 func(node any) any
											_ = _hx_closure_src_743
											_hx_closure_src_743 = this.VTable.Hx_Field_attachComments
											return func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
												var hx_dyncast_343 any
												_ = hx_dyncast_343
												hx_dyncast_343 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_743(((any)(node))), "Hx_Obj_loreline_naccess")
												var _hx_tmp_9 *Hx_Obj_loreline_naccess
												_ = _hx_tmp_9
												if hx_dyncast_343 != nil {
													_hx_tmp_9 = hx_dyncast_343.(interface {
														Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
													}).Hx_As_Hx_Obj_loreline_naccess()
												} else {
													_hx_tmp_9 = nil
												}

												return _hx_tmp_9
											}
										}()(this.VTable.Hx_Field_makeAccess(token.Hx_Field_pos, target, name, token.Hx_Field_pos))
										if _hx_up != nil {
											return &_hx_up.Hx_Obj_loreline_nexpr
										}
										return nil
									})()
								} else {
									this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Missing dot in field access", token.Hx_Field_pos))
									var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
									_ = tmp_tmp_1
									tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
									var tmp1 Hx_Typedef_haxe__int64___int64
									_ = tmp1
									tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
									var tmp2 *Hx_Obj_loreline_token
									_ = tmp2
									tmp2 = tokens.Get(0)
									var _hx_tmp_10 *Hx_Obj_loreline_position
									_ = _hx_tmp_10
									if tmp2 != nil {
										_hx_tmp_10 = tmp2.Hx_Field_pos
									} else {
										_hx_tmp_10 = nil
									}

									var pos1 *Hx_Obj_loreline_position
									_ = pos1
									pos1 = _hx_tmp_10
									var _hx_tmp_11 *Hx_Obj_loreline_position
									_ = _hx_tmp_11
									if pos1 != nil {
										_hx_tmp_11 = pos1
									} else {
										_hx_tmp_11 = this.VTable.Hx_Field_currentPos()
									}

									return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
										_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_11, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
										if _hx_up != nil {
											return &_hx_up.Hx_Obj_loreline_nexpr
										}
										return nil
									})()}, nil, nil)
								}
							}

							prevIsDot = false
						}
					}

				case 45:
					{
						if target == nil {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Leading dot in field access", token.Hx_Field_pos))
							var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp1 Hx_Typedef_haxe__int64___int64
							_ = tmp1
							tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp2 *Hx_Obj_loreline_token
							_ = tmp2
							tmp2 = tokens.Get(0)
							var _hx_tmp_12 *Hx_Obj_loreline_position
							_ = _hx_tmp_12
							if tmp2 != nil {
								_hx_tmp_12 = tmp2.Hx_Field_pos
							} else {
								_hx_tmp_12 = nil
							}

							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							pos1 = _hx_tmp_12
							var _hx_tmp_13 *Hx_Obj_loreline_position
							_ = _hx_tmp_13
							if pos1 != nil {
								_hx_tmp_13 = pos1
							} else {
								_hx_tmp_13 = this.VTable.Hx_Field_currentPos()
							}

							return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_13, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()}, nil, nil)
						}

						prevIsDot = true
					}

				case 48:
					{
						if target == nil {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Function call without target", token.Hx_Field_pos))
							var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp1 Hx_Typedef_haxe__int64___int64
							_ = tmp1
							tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp2 *Hx_Obj_loreline_token
							_ = tmp2
							tmp2 = tokens.Get(0)
							var _hx_tmp_14 *Hx_Obj_loreline_position
							_ = _hx_tmp_14
							if tmp2 != nil {
								_hx_tmp_14 = tmp2.Hx_Field_pos
							} else {
								_hx_tmp_14 = nil
							}

							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							pos1 = _hx_tmp_14
							var _hx_tmp_15 *Hx_Obj_loreline_position
							_ = _hx_tmp_15
							if pos1 != nil {
								_hx_tmp_15 = pos1
							} else {
								_hx_tmp_15 = this.VTable.Hx_Field_currentPos()
							}

							return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_15, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()}, nil, nil)
						}

						var callStart *Hx_Obj_loreline_position
						_ = callStart
						callStart = token.Hx_Field_pos
						var argTokens HxArray[HxArray[*Hx_Obj_loreline_token]]
						_ = argTokens
						argTokens = HxMakeArray[HxArray[*Hx_Obj_loreline_token]]()
						var currentArgTokens HxArray[*Hx_Obj_loreline_token]
						_ = currentArgTokens
						currentArgTokens = HxMakeArray[*Hx_Obj_loreline_token]()
						var parenLevel int32
						_ = parenLevel
						parenLevel = 1
						for {
							var _hx_tmp_16 bool
							_ = _hx_tmp_16
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							_hx_tmp_17 = i
							if _hx_tmp_17 < tokens.Len() {
								_hx_tmp_16 = (parenLevel > 0)
							} else {
								_hx_tmp_16 = false
							}

							if !(_hx_tmp_16) {
								break
							}

							var t *Hx_Obj_loreline_token
							_ = t
							t = tokens.Get(i)
							i++
							switch t.Hx_Field_type.Hx_Field_enumIndex() {
							case 44:
								{
									if parenLevel == 1 {
										if currentArgTokens.Len() > 0 {
											_ = (Hx_Array_Push[HxArray[*Hx_Obj_loreline_token]](argTokens, currentArgTokens))
											currentArgTokens = HxMakeArray[*Hx_Obj_loreline_token]()
										}
									} else {
										_ = (Hx_Array_Push[*Hx_Obj_loreline_token](currentArgTokens, t))
									}
								}

							case 48:
								{
									parenLevel++
									_ = (Hx_Array_Push[*Hx_Obj_loreline_token](currentArgTokens, t))
								}

							case 49:
								{
									parenLevel--
									if parenLevel > 0 {
										_ = (Hx_Array_Push[*Hx_Obj_loreline_token](currentArgTokens, t))
									} else {
										if currentArgTokens.Len() > 0 {
											_ = (Hx_Array_Push[HxArray[*Hx_Obj_loreline_token]](argTokens, currentArgTokens))
										}
									}
								}

							default:
								{
									_ = (Hx_Array_Push[*Hx_Obj_loreline_token](currentArgTokens, t))
								}
							}
						}

						if parenLevel > 0 {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Unterminated function call", callStart))
							var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp1 Hx_Typedef_haxe__int64___int64
							_ = tmp1
							tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp2 *Hx_Obj_loreline_token
							_ = tmp2
							tmp2 = tokens.Get(0)
							var _hx_tmp_18 *Hx_Obj_loreline_position
							_ = _hx_tmp_18
							if tmp2 != nil {
								_hx_tmp_18 = tmp2.Hx_Field_pos
							} else {
								_hx_tmp_18 = nil
							}

							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							pos1 = _hx_tmp_18
							var _hx_tmp_19 *Hx_Obj_loreline_position
							_ = _hx_tmp_19
							if pos1 != nil {
								_hx_tmp_19 = pos1
							} else {
								_hx_tmp_19 = this.VTable.Hx_Field_currentPos()
							}

							return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_19, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()}, nil, nil)
						}

						var args HxArray[*Hx_Obj_loreline_nexpr]
						_ = args
						args = HxMakeArray[*Hx_Obj_loreline_nexpr]()
						{
							var tmp__g_1 int32
							_ = tmp__g_1
							tmp__g_1 = 0
							for {
								var _hx_tmp_20 int32
								_ = _hx_tmp_20
								_hx_tmp_20 = tmp__g_1
								if !(_hx_tmp_20 < argTokens.Len()) {
									break
								}

								var argTokenGroup HxArray[*Hx_Obj_loreline_token]
								_ = argTokenGroup
								argTokenGroup = argTokens.Get(tmp__g_1)
								tmp__g_1++
								var tempParser *Hx_Obj_loreline_parser
								_ = tempParser
								tempParser = Hx_Obj_loreline_parser_CreateInstance(((Hx_Typedef_loreline_tokens)(argTokenGroup)), nil)
								tempParser.Hx_Field_currentNodeId = this.Hx_Field_currentNodeId
								var arg *Hx_Obj_loreline_nexpr
								_ = arg
								arg = tempParser.VTable.Hx_Field_parseExpression()
								_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](args, arg))
								this.Hx_Field_currentNodeId = tempParser.Hx_Field_currentNodeId
							}
						}

						target = (func() *Hx_Obj_loreline_nexpr {
							_hx_up := func() func(node *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_ncall {
								var _hx_closure_src_744 func(node any) any
								_ = _hx_closure_src_744
								_hx_closure_src_744 = this.VTable.Hx_Field_attachComments
								return func(node *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_ncall {
									var hx_dyncast_344 any
									_ = hx_dyncast_344
									hx_dyncast_344 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_744(((any)(node))), "Hx_Obj_loreline_ncall")
									var _hx_tmp_21 *Hx_Obj_loreline_ncall
									_ = _hx_tmp_21
									if hx_dyncast_344 != nil {
										_hx_tmp_21 = hx_dyncast_344.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()
									} else {
										_hx_tmp_21 = nil
									}

									return _hx_tmp_21
								}
							}()(Hx_Obj_loreline_ncall_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), callStart.VTable.Hx_Field_extendedTo(tokens.Get((i - 1)).Hx_Field_pos), target, args, nil, nil))
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_nexpr
							}
							return nil
						})()
						prevIsDot = false
					}

				case 50:
					{
						if target == nil {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Array access without target", token.Hx_Field_pos))
							var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp1 Hx_Typedef_haxe__int64___int64
							_ = tmp1
							tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp2 *Hx_Obj_loreline_token
							_ = tmp2
							tmp2 = tokens.Get(0)
							var _hx_tmp_22 *Hx_Obj_loreline_position
							_ = _hx_tmp_22
							if tmp2 != nil {
								_hx_tmp_22 = tmp2.Hx_Field_pos
							} else {
								_hx_tmp_22 = nil
							}

							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							pos1 = _hx_tmp_22
							var _hx_tmp_23 *Hx_Obj_loreline_position
							_ = _hx_tmp_23
							if pos1 != nil {
								_hx_tmp_23 = pos1
							} else {
								_hx_tmp_23 = this.VTable.Hx_Field_currentPos()
							}

							return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_23, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()}, nil, nil)
						}

						var arrayStart *Hx_Obj_loreline_position
						_ = arrayStart
						arrayStart = token.Hx_Field_pos
						var arrayTokens HxArray[*Hx_Obj_loreline_token]
						_ = arrayTokens
						arrayTokens = HxMakeArray[*Hx_Obj_loreline_token]()
						var bracketLevel int32
						_ = bracketLevel
						bracketLevel = 1
						var lastRBracket *Hx_Obj_loreline_token
						_ = lastRBracket
						for {
							var _hx_tmp_24 bool
							_ = _hx_tmp_24
							var _hx_tmp_25 int32
							_ = _hx_tmp_25
							_hx_tmp_25 = i
							if _hx_tmp_25 < tokens.Len() {
								_hx_tmp_24 = (bracketLevel > 0)
							} else {
								_hx_tmp_24 = false
							}

							if !(_hx_tmp_24) {
								break
							}

							var t *Hx_Obj_loreline_token
							_ = t
							t = tokens.Get(i)
							i++
							switch t.Hx_Field_type.Hx_Field_enumIndex() {
							case 50:
								{
									bracketLevel++
									_ = (Hx_Array_Push[*Hx_Obj_loreline_token](arrayTokens, t))
								}

							case 51:
								{
									bracketLevel--
									if bracketLevel > 0 {
										_ = (Hx_Array_Push[*Hx_Obj_loreline_token](arrayTokens, t))
									}

									lastRBracket = t
								}

							default:
								{
									_ = (Hx_Array_Push[*Hx_Obj_loreline_token](arrayTokens, t))
								}
							}
						}

						if bracketLevel > 0 {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Unterminated array access", arrayStart))
							var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp1 Hx_Typedef_haxe__int64___int64
							_ = tmp1
							tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
							var tmp2 *Hx_Obj_loreline_token
							_ = tmp2
							tmp2 = tokens.Get(0)
							var _hx_tmp_26 *Hx_Obj_loreline_position
							_ = _hx_tmp_26
							if tmp2 != nil {
								_hx_tmp_26 = tmp2.Hx_Field_pos
							} else {
								_hx_tmp_26 = nil
							}

							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							pos1 = _hx_tmp_26
							var _hx_tmp_27 *Hx_Obj_loreline_position
							_ = _hx_tmp_27
							if pos1 != nil {
								_hx_tmp_27 = pos1
							} else {
								_hx_tmp_27 = this.VTable.Hx_Field_currentPos()
							}

							return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_27, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()}, nil, nil)
						}

						var tempParser *Hx_Obj_loreline_parser
						_ = tempParser
						tempParser = Hx_Obj_loreline_parser_CreateInstance(((Hx_Typedef_loreline_tokens)(arrayTokens)), nil)
						tempParser.Hx_Field_currentNodeId = this.Hx_Field_currentNodeId
						var indexExpr *Hx_Obj_loreline_nexpr
						_ = indexExpr
						indexExpr = tempParser.VTable.Hx_Field_parseExpression()
						this.Hx_Field_currentNodeId = tempParser.Hx_Field_currentNodeId
						var accessPos *Hx_Obj_loreline_position
						_ = accessPos
						accessPos = pos.VTable.Hx_Field_extendedTo(tokens.Get((i - 1)).Hx_Field_pos)
						accessPos.Hx_Field_length = (accessPos.Hx_Field_length + 1)
						target = (func() *Hx_Obj_loreline_nexpr {
							_hx_up := func() func(node *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_narrayaccess {
								var _hx_closure_src_745 func(node any) any
								_ = _hx_closure_src_745
								_hx_closure_src_745 = this.VTable.Hx_Field_attachComments
								return func(node *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_narrayaccess {
									var hx_dyncast_345 any
									_ = hx_dyncast_345
									hx_dyncast_345 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_745(((any)(node))), "Hx_Obj_loreline_narrayaccess")
									var _hx_tmp_28 *Hx_Obj_loreline_narrayaccess
									_ = _hx_tmp_28
									if hx_dyncast_345 != nil {
										_hx_tmp_28 = hx_dyncast_345.(interface {
											Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
										}).Hx_As_Hx_Obj_loreline_narrayaccess()
									} else {
										_hx_tmp_28 = nil
									}

									return _hx_tmp_28
								}
							}()(Hx_Obj_loreline_narrayaccess_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), accessPos, target, indexExpr, nil, nil))
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_nexpr
							}
							return nil
						})()
						prevIsDot = false
					}

				default:
					{
						this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unexpected token in field access: ", Hx_Field_loreline_tokentypehelpers_toCodeString(token.Hx_Field_type)), token.Hx_Field_pos))
						var tmp_tmp_1 Hx_Typedef_haxe__int64___int64
						_ = tmp_tmp_1
						tmp_tmp_1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
						var tmp1 Hx_Typedef_haxe__int64___int64
						_ = tmp1
						tmp1 = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
						var tmp2 *Hx_Obj_loreline_token
						_ = tmp2
						tmp2 = tokens.Get(0)
						var _hx_tmp_29 *Hx_Obj_loreline_position
						_ = _hx_tmp_29
						if tmp2 != nil {
							_hx_tmp_29 = tmp2.Hx_Field_pos
						} else {
							_hx_tmp_29 = nil
						}

						var pos1 *Hx_Obj_loreline_position
						_ = pos1
						pos1 = _hx_tmp_29
						var _hx_tmp_30 *Hx_Obj_loreline_position
						_ = _hx_tmp_30
						if pos1 != nil {
							_hx_tmp_30 = pos1
						} else {
							_hx_tmp_30 = this.VTable.Hx_Field_currentPos()
						}

						return Hx_Obj_loreline_nstringpart_CreateInstance(tmp_tmp_1, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
							_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(tmp1, _hx_tmp_30, nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_nexpr
							}
							return nil
						})()}, nil, nil)
					}
				}
			}
		}

		if prevIsDot {
			var _hx_tmp_31 Hx_Typedef_loreline_tokens
			_ = _hx_tmp_31
			_hx_tmp_31 = tokens
			this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Trailing dot in field access", _hx_tmp_31.Get((tokens.Len() - 1)).Hx_Field_pos))
		}

		expr = target
	} else {
		var tempParser *Hx_Obj_loreline_parser
		_ = tempParser
		tempParser = Hx_Obj_loreline_parser_CreateInstance(tokens, nil)
		tempParser.Hx_Field_currentNodeId = this.Hx_Field_currentNodeId
		expr = tempParser.VTable.Hx_Field_parseExpression()
		this.Hx_Field_currentNodeId = tempParser.Hx_Field_currentNodeId
		if !tempParser.VTable.Hx_Field_isAtEnd() {
			this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Unexpected tokens after interpolation expression", tempParser.Hx_Field_tokens.Get(tempParser.Hx_Field_current).Hx_Field_pos))
		}
	}

	var _hx_tmp_32 int32
	_ = _hx_tmp_32
	_hx_tmp_32 = pos.Hx_Field_column
	var _hx_tmp_33 int32
	_ = _hx_tmp_33
	if braces {
		_hx_tmp_33 = 1
	} else {
		_hx_tmp_33 = 0
	}

	var _hx_tmp_34 int32
	_ = _hx_tmp_34
	_hx_tmp_34 = pos.Hx_Field_offset
	var _hx_tmp_35 int32
	_ = _hx_tmp_35
	if braces {
		_hx_tmp_35 = 2
	} else {
		_hx_tmp_35 = 0
	}

	var partPos *Hx_Obj_loreline_position
	_ = partPos
	partPos = Hx_Obj_loreline_position_CreateInstance(pos.Hx_Field_line, (_hx_tmp_32 + (_hx_tmp_33)), (_hx_tmp_34 - (_hx_tmp_35)), length)
	return Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Expr{expr}, nil, nil)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseStringTag(closing bool, start int32, length int32, content string, quotes int32, attachments HxArray[Hx_Enum_loreline_lstringattachment]) *Hx_Obj_loreline_nstringpart {
	var strPos *Hx_Obj_loreline_position
	_ = strPos
	strPos = this.VTable.Hx_Field_currentPos()
	if quotes != ((int32)(0)) {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = content
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = (strPos.Hx_Field_length - 2)
		strPos = strPos.VTable.Hx_Field_withOffset(_hx_tmp_0, 1, _hx_tmp_1, nil)
	}

	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = this.VTable.Hx_Field_makeStringPartPosition(strPos, content, start)
	pos.Hx_Field_length = length
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	if closing {
		_hx_tmp_2 = 2
	} else {
		_hx_tmp_2 = 1
	}

	var offsetStart int32
	_ = offsetStart
	offsetStart = _hx_tmp_2
	var innerStart int32
	_ = innerStart
	innerStart = (start + offsetStart)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = length
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	if closing {
		_hx_tmp_4 = 3
	} else {
		_hx_tmp_4 = 2
	}

	var innerLength int32
	_ = innerLength
	innerLength = (_hx_tmp_3 - (_hx_tmp_4))
	var innerEnd int32
	_ = innerEnd
	innerEnd = (innerStart + innerLength)
	var tagId Hx_Typedef_haxe__int64___int64
	_ = tagId
	tagId = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
	var hasAttachmentsInRange bool
	_ = hasAttachmentsInRange
	hasAttachmentsInRange = false
	if attachments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = attachments.Len()
	hx_label_0:
		for _g < _g1 {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_5
			{
				var tmp__g1_1 Hx_Enum_loreline_lstringattachment
				_ = tmp__g1_1
				tmp__g1_1 = attachments.Get(i)
				switch tmp__g1_1.Hx_Field_enumIndex() {
				case 0:
					{
						var aStart int32
						_ = aStart
						aStart = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
						var aLength int32
						_ = aLength
						aLength = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_length
						{
							var aEnd int32
							_ = aEnd
							aEnd = (aStart + aLength)
							if (aStart >= innerStart) && (aEnd <= innerEnd) {
								hasAttachmentsInRange = true
								break hx_label_0
							}
						}
					}

				case 1:
					{
						continue
					}

				}
			}
		}
	}

	if !hasAttachmentsInRange {
		var partPos *Hx_Obj_loreline_position
		_ = partPos
		partPos = this.VTable.Hx_Field_makeStringPartPosition(strPos, content, innerStart)
		partPos.Hx_Field_length = innerLength
		var literalId Hx_Typedef_haxe__int64___int64
		_ = literalId
		literalId = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
		var partId Hx_Typedef_haxe__int64___int64
		_ = partId
		partId = this.VTable.Hx_Field_nextNodeId(((int32)(4)))
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		_hx_tmp_6 = closing
		var _hx_tmp_8 string
		_ = _hx_tmp_8
		if (HxNullable[int32]{Value: innerLength, Valid: true}).Valid != false {
			var _hx_tmp_9 string
			_ = _hx_tmp_9
			_hx_tmp_9 = content
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = innerStart
			var _hx_tmp_11 string
			_ = _hx_tmp_11
			_hx_tmp_11 = _hx_tmp_9
			var _hx_tmp_12 int32
			_ = _hx_tmp_12
			_hx_tmp_12 = _hx_tmp_10
			_hx_tmp_8 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_11, _hx_tmp_12, HxNullable[int32]{Value: innerLength, Valid: true}))
		} else {
			var _hx_tmp_13 string
			_ = _hx_tmp_13
			_hx_tmp_13 = content
			var _hx_tmp_14 int32
			_ = _hx_tmp_14
			_hx_tmp_14 = innerStart
			var _hx_tmp_15 string
			_ = _hx_tmp_15
			_hx_tmp_15 = _hx_tmp_13
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = _hx_tmp_14
			_hx_tmp_8 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, _hx_tmp_16, HxNullable[int32]{}))
		}

		return Hx_Obj_loreline_nstringpart_CreateInstance(tagId, pos, Hx_Enum_loreline_stringparttype_Tag{_hx_tmp_6, func() func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
			var _hx_closure_src_746 func(node any) any
			_ = _hx_closure_src_746
			_hx_closure_src_746 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
				var hx_dyncast_346 any
				_ = hx_dyncast_346
				hx_dyncast_346 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_746(((any)(node))), "Hx_Obj_loreline_nstringliteral")
				var _hx_tmp_7 *Hx_Obj_loreline_nstringliteral
				_ = _hx_tmp_7
				if hx_dyncast_346 != nil {
					_hx_tmp_7 = hx_dyncast_346.(interface {
						Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
					}).Hx_As_Hx_Obj_loreline_nstringliteral()
				} else {
					_hx_tmp_7 = nil
				}

				return _hx_tmp_7
			}
		}()(Hx_Obj_loreline_nstringliteral_CreateInstance(literalId, partPos, ((int32)(0)), HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(partId, partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_8}, nil, nil)), nil, nil))}, nil, nil)
	}

	var parts HxArray[*Hx_Obj_loreline_nstringpart]
	_ = parts
	parts = HxMakeArray[*Hx_Obj_loreline_nstringpart]()
	var stringLiteral *Hx_Obj_loreline_nstringliteral
	_ = stringLiteral
	stringLiteral = func() func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
		var _hx_closure_src_747 func(node any) any
		_ = _hx_closure_src_747
		_hx_closure_src_747 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
			var hx_dyncast_347 any
			_ = hx_dyncast_347
			hx_dyncast_347 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_747(((any)(node))), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_17 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_17
			if hx_dyncast_347 != nil {
				_hx_tmp_17 = hx_dyncast_347.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
			} else {
				_hx_tmp_17 = nil
			}

			return _hx_tmp_17
		}
	}()(Hx_Obj_loreline_nstringliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), pos, ((int32)(0)), parts, nil, nil))
	var currentPos int32
	_ = currentPos
	currentPos = innerStart
	if attachments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = attachments.Len()
		for _g < _g1 {
			var _hx_tmp_18 int32
			_ = _hx_tmp_18
			_hx_tmp_18 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_18
			{
				var tmp__g1_1 Hx_Enum_loreline_lstringattachment
				_ = tmp__g1_1
				tmp__g1_1 = attachments.Get(i)
				switch tmp__g1_1.Hx_Field_enumIndex() {
				case 0:
					{
						var braces bool
						_ = braces
						braces = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_braces
						var tokens Hx_Typedef_loreline_tokens
						_ = tokens
						tokens = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_expr
						var aStart int32
						_ = aStart
						aStart = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
						var aLength int32
						_ = aLength
						aLength = tmp__g1_1.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_length
						{
							var aEnd int32
							_ = aEnd
							aEnd = (aStart + aLength)
							if (aStart >= innerStart) && (aEnd <= innerEnd) {
								if aStart > currentPos {
									var partPos *Hx_Obj_loreline_position
									_ = partPos
									partPos = this.VTable.Hx_Field_makeStringPartPosition(strPos, content, innerStart)
									partPos.Hx_Field_length = (aStart - innerStart)
									var _hx_reserved_len HxNullable[int32]
									_ = _hx_reserved_len
									_hx_reserved_len = HxNullable[int32]{Value: (aStart - currentPos), Valid: true}

									var _hx_tmp_19 HxArray[*Hx_Obj_loreline_nstringpart]
									_ = _hx_tmp_19
									_hx_tmp_19 = parts
									var _hx_tmp_20 string
									_ = _hx_tmp_20
									if (_hx_reserved_len).Valid != false {
										_hx_tmp_20 = (Hx_Field_go_haxe_hxstring_substr(content, currentPos, _hx_reserved_len))
									} else {
										var _hx_tmp_21 string
										_ = _hx_tmp_21
										_hx_tmp_21 = content
										var _hx_tmp_22 int32
										_ = _hx_tmp_22
										_hx_tmp_22 = currentPos
										var _hx_tmp_23 string
										_ = _hx_tmp_23
										_hx_tmp_23 = _hx_tmp_21
										var _hx_tmp_24 int32
										_ = _hx_tmp_24
										_hx_tmp_24 = _hx_tmp_22
										_hx_tmp_20 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_23, _hx_tmp_24, HxNullable[int32]{}))
									}

									_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_19, Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_20}, nil, nil)))
								}

								var interpPart *Hx_Obj_loreline_nstringpart
								_ = interpPart
								interpPart = this.VTable.Hx_Field_parseStringInterpolation(braces, true, tokens, aStart, aLength, content)
								interpPart.Hx_Field_pos = this.VTable.Hx_Field_makeStringPartPosition(strPos, content, aStart)
								interpPart.Hx_Field_pos.Hx_Field_length = aLength
								_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](parts, interpPart))
								currentPos = aEnd
							}
						}
					}

				case 1:
					{
						continue
					}

				}
			}
		}
	}

	if currentPos < innerEnd {
		var _hx_reserved_len HxNullable[int32]
		_ = _hx_reserved_len
		_hx_reserved_len = HxNullable[int32]{}

		var _hx_tmp_25 *Hx_Obj_loreline_position
		_ = _hx_tmp_25
		_hx_tmp_25 = pos
		var _hx_tmp_26 string
		_ = _hx_tmp_26
		if (_hx_reserved_len).Valid != false {
			_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(content, start, _hx_reserved_len))
		} else {
			var _hx_tmp_27 string
			_ = _hx_tmp_27
			_hx_tmp_27 = content
			var _hx_tmp_28 int32
			_ = _hx_tmp_28
			_hx_tmp_28 = start
			var _hx_tmp_29 string
			_ = _hx_tmp_29
			_hx_tmp_29 = _hx_tmp_27
			var _hx_tmp_30 int32
			_ = _hx_tmp_30
			_hx_tmp_30 = _hx_tmp_28
			_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_29, _hx_tmp_30, HxNullable[int32]{}))
		}

		var partPos *Hx_Obj_loreline_position
		_ = partPos
		partPos = this.VTable.Hx_Field_makeStringPartPosition(_hx_tmp_25, _hx_tmp_26, ((currentPos - start) + offsetStart))
		partPos.Hx_Field_length = ((innerStart + innerEnd) - currentPos)
		var len1 HxNullable[int32]
		_ = len1
		len1 = HxNullable[int32]{Value: (innerEnd - currentPos), Valid: true}

		var _hx_tmp_31 HxArray[*Hx_Obj_loreline_nstringpart]
		_ = _hx_tmp_31
		_hx_tmp_31 = parts
		var _hx_tmp_32 string
		_ = _hx_tmp_32
		if (len1).Valid != false {
			_hx_tmp_32 = (Hx_Field_go_haxe_hxstring_substr(content, currentPos, len1))
		} else {
			var _hx_tmp_33 string
			_ = _hx_tmp_33
			_hx_tmp_33 = content
			var _hx_tmp_34 int32
			_ = _hx_tmp_34
			_hx_tmp_34 = currentPos
			var _hx_tmp_35 string
			_ = _hx_tmp_35
			_hx_tmp_35 = _hx_tmp_33
			var _hx_tmp_36 int32
			_ = _hx_tmp_36
			_hx_tmp_36 = _hx_tmp_34
			_hx_tmp_32 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_35, _hx_tmp_36, HxNullable[int32]{}))
		}

		_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_31, Hx_Obj_loreline_nstringpart_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), partPos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_32}, nil, nil)))
	}

	return Hx_Obj_loreline_nstringpart_CreateInstance(tagId, pos, Hx_Enum_loreline_stringparttype_Tag{closing, stringLiteral}, nil, nil)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseIdentifierExpression(startPos *Hx_Obj_loreline_position, name string) *Hx_Obj_loreline_nexpr {
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	expr = (func() *Hx_Obj_loreline_nexpr {
		_hx_up := func() func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
			var _hx_closure_src_749 func(node any) any
			_ = _hx_closure_src_749
			_hx_closure_src_749 = this.VTable.Hx_Field_attachComments
			return func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
				var hx_dyncast_349 any
				_ = hx_dyncast_349
				hx_dyncast_349 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_749(((any)(node))), "Hx_Obj_loreline_naccess")
				var _hx_tmp_0 *Hx_Obj_loreline_naccess
				_ = _hx_tmp_0
				if hx_dyncast_349 != nil {
					_hx_tmp_0 = hx_dyncast_349.(interface {
						Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
					}).Hx_As_Hx_Obj_loreline_naccess()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(this.VTable.Hx_Field_makeAccess(startPos, nil, name, nil))
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})()
	return this.VTable.Hx_Field_parsePostfix(startPos, expr)
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parsePostfix(startPos *Hx_Obj_loreline_position, expr *Hx_Obj_loreline_nexpr) *Hx_Obj_loreline_nexpr {
	for {
		if (this.VTable.Hx_Field_match(Hx_Enum_loreline_tokentype_Dot{}, false)) {
			var prop HxNullable[string]
			_ = prop
			prop = HxNullable[string]{Value: HxStringNull}

			var propPos *Hx_Obj_loreline_position
			_ = propPos
			propPos = this.VTable.Hx_Field_currentPos()
			if this.VTable.Hx_Field_checkIdentifier() {
				prop = HxNullableString(this.VTable.Hx_Field_expectIdentifier())
			} else {
				this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Expected identifier after '.'", this.VTable.Hx_Field_currentPos()))
				prop = HxNullableString("")
			}

			var _hx_tmp_1 *Hx_Obj_loreline_position
			_ = _hx_tmp_1
			_hx_tmp_1 = startPos
			var _hx_tmp_2 *Hx_Obj_loreline_nexpr
			_ = _hx_tmp_2
			_hx_tmp_2 = expr
			expr = (func() *Hx_Obj_loreline_nexpr {
				_hx_up := func() func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
					var _hx_closure_src_753 func(node any) any
					_ = _hx_closure_src_753
					_hx_closure_src_753 = this.VTable.Hx_Field_attachComments
					return func(node *Hx_Obj_loreline_naccess) *Hx_Obj_loreline_naccess {
						var hx_dyncast_353 any
						_ = hx_dyncast_353
						hx_dyncast_353 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_753(((any)(node))), "Hx_Obj_loreline_naccess")
						var _hx_tmp_0 *Hx_Obj_loreline_naccess
						_ = _hx_tmp_0
						if hx_dyncast_353 != nil {
							_hx_tmp_0 = hx_dyncast_353.(interface {
								Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
							}).Hx_As_Hx_Obj_loreline_naccess()
						} else {
							_hx_tmp_0 = nil
						}

						return _hx_tmp_0
					}
				}()(this.VTable.Hx_Field_makeAccess(_hx_tmp_1, _hx_tmp_2, (prop).Value, propPos))
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_nexpr
				}
				return nil
			})()
		} else {
			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_LBracket{}

			if this.VTable.Hx_Field_match(_hx_tmp_3, nil) {
				var index *Hx_Obj_loreline_nexpr
				_ = index
				index = this.VTable.Hx_Field_parseExpression()
				this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_RBracket{})
				var accessPos *Hx_Obj_loreline_position
				_ = accessPos
				accessPos = startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_previous().Hx_Field_pos)
				expr = (func() *Hx_Obj_loreline_nexpr {
					_hx_up := func() func(node *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_narrayaccess {
						var _hx_closure_src_754 func(node any) any
						_ = _hx_closure_src_754
						_hx_closure_src_754 = this.VTable.Hx_Field_attachComments
						return func(node *Hx_Obj_loreline_narrayaccess) *Hx_Obj_loreline_narrayaccess {
							var hx_dyncast_354 any
							_ = hx_dyncast_354
							hx_dyncast_354 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_754(((any)(node))), "Hx_Obj_loreline_narrayaccess")
							var _hx_tmp_4 *Hx_Obj_loreline_narrayaccess
							_ = _hx_tmp_4
							if hx_dyncast_354 != nil {
								_hx_tmp_4 = hx_dyncast_354.(interface {
									Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
								}).Hx_As_Hx_Obj_loreline_narrayaccess()
							} else {
								_hx_tmp_4 = nil
							}

							return _hx_tmp_4
						}
					}()(Hx_Obj_loreline_narrayaccess_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), accessPos, expr, index, nil, nil))
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_nexpr
					}
					return nil
				})()
			} else {
				var _hx_tmp_5 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_5
				_hx_tmp_5 = Hx_Enum_loreline_tokentype_LParen{}

				if this.VTable.Hx_Field_match(_hx_tmp_5, nil) {
					var args HxArray[*Hx_Obj_loreline_nexpr]
					_ = args
					args = this.VTable.Hx_Field_parseCallArguments()
					expr = (func() *Hx_Obj_loreline_nexpr {
						_hx_up := func() func(node *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_ncall {
							var _hx_closure_src_755 func(node any) any
							_ = _hx_closure_src_755
							_hx_closure_src_755 = this.VTable.Hx_Field_attachComments
							return func(node *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_ncall {
								var hx_dyncast_355 any
								_ = hx_dyncast_355
								hx_dyncast_355 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_755(((any)(node))), "Hx_Obj_loreline_ncall")
								var _hx_tmp_6 *Hx_Obj_loreline_ncall
								_ = _hx_tmp_6
								if hx_dyncast_355 != nil {
									_hx_tmp_6 = hx_dyncast_355.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()
								} else {
									_hx_tmp_6 = nil
								}

								return _hx_tmp_6
							}
						}()(Hx_Obj_loreline_ncall_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_previous().Hx_Field_pos), expr, args, nil, nil))
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_nexpr
						}
						return nil
					})()
				} else {
					break
				}
			}
		}
	}

	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseArrayLiteral() *Hx_Obj_loreline_nexpr {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var elements HxArray[*Hx_Obj_loreline_nexpr]
	_ = elements
	elements = HxMakeArray[*Hx_Obj_loreline_nexpr]()
	var literal *Hx_Obj_loreline_nliteral
	_ = literal
	literal = Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), startPos, elements, Hx_Enum_loreline_literaltype_Array{}, nil, nil)
	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_LBracket{})
	func() func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
		var _hx_closure_src_756 func(node any) any
		_ = _hx_closure_src_756
		_hx_closure_src_756 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
			var hx_dyncast_356 any
			_ = hx_dyncast_356
			hx_dyncast_356 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_756(((any)(node))), "Hx_Obj_loreline_nliteral")
			var _hx_tmp_0 *Hx_Obj_loreline_nliteral
			_ = _hx_tmp_0
			if hx_dyncast_356 != nil {
				_hx_tmp_0 = hx_dyncast_356.(interface {
					Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
				}).Hx_As_Hx_Obj_loreline_nliteral()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(literal)
	var needsSeparator bool
	_ = needsSeparator
	needsSeparator = false
	for {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_RBracket{})) {
			_hx_tmp_1 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_1 = false
		}

		if !(_hx_tmp_1) {
			break
		}

		if needsSeparator {
			for {
				var _hx_tmp_2 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Enum_loreline_tokentype_LineBreak{}

				if !(this.VTable.Hx_Field_match(_hx_tmp_2, nil)) {
					break
				}

				needsSeparator = false
			}

			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_Comma{}

			if this.VTable.Hx_Field_match(_hx_tmp_3, nil) {
				needsSeparator = false
			}
		}

		for {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			var _hx_tmp_6 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_6, nil) {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_7 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_7
				_hx_tmp_7 = Hx_Enum_loreline_tokentype_Indent{}

				var _hx_tmp_8 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_8
				_hx_tmp_8 = _hx_tmp_7
				_hx_tmp_5 = this.VTable.Hx_Field_match(_hx_tmp_8, nil)
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_9 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_9
				_hx_tmp_9 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_10 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_10
				_hx_tmp_10 = _hx_tmp_9
				_hx_tmp_4 = this.VTable.Hx_Field_match(_hx_tmp_10, nil)
			}

			if !(_hx_tmp_4) {
				break
			}

			needsSeparator = false
		}

		var _hx_tmp_11 bool
		_ = _hx_tmp_11
		if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_RBracket{})) {
			_hx_tmp_11 = needsSeparator
		} else {
			_hx_tmp_11 = false
		}

		if _hx_tmp_11 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected comma or line break between elements", this.VTable.Hx_Field_currentPos())))))
		}

		for {
			var _hx_tmp_12 bool
			_ = _hx_tmp_12
			var _hx_tmp_13 bool
			_ = _hx_tmp_13
			var _hx_tmp_14 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_14
			_hx_tmp_14 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_14, nil) {
				_hx_tmp_13 = true
			} else {
				var _hx_tmp_15 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_15
				_hx_tmp_15 = Hx_Enum_loreline_tokentype_Indent{}

				var _hx_tmp_16 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_16
				_hx_tmp_16 = _hx_tmp_15
				_hx_tmp_13 = this.VTable.Hx_Field_match(_hx_tmp_16, nil)
			}

			if _hx_tmp_13 {
				_hx_tmp_12 = true
			} else {
				var _hx_tmp_17 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_17
				_hx_tmp_17 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_18 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_18
				_hx_tmp_18 = _hx_tmp_17
				_hx_tmp_12 = this.VTable.Hx_Field_match(_hx_tmp_18, nil)
			}

			if !(_hx_tmp_12) {
				break
			}
		}

		if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_RBracket{})) {
			var _hx_tmp_19 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_19
			_hx_tmp_19 = elements
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_19, this.VTable.Hx_Field_parseExpression()))
		}

		var prev *Hx_Obj_loreline_token
		_ = prev
		prev = this.VTable.Hx_Field_previous()
		needsSeparator = (((prev.Hx_Field_type != Hx_Enum_loreline_tokentype_Colon{}) && (prev.Hx_Field_type != Hx_Enum_loreline_tokentype_LineBreak{})) && (prev.Hx_Field_type != Hx_Enum_loreline_tokentype_Unindent{}))
	}

	for {
		var _hx_tmp_20 bool
		_ = _hx_tmp_20
		var _hx_tmp_21 bool
		_ = _hx_tmp_21
		var _hx_tmp_22 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_22
		_hx_tmp_22 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_22, nil) {
			_hx_tmp_21 = true
		} else {
			var _hx_tmp_23 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_23
			_hx_tmp_23 = Hx_Enum_loreline_tokentype_Indent{}

			var _hx_tmp_24 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_24
			_hx_tmp_24 = _hx_tmp_23
			_hx_tmp_21 = this.VTable.Hx_Field_match(_hx_tmp_24, nil)
		}

		if _hx_tmp_21 {
			_hx_tmp_20 = true
		} else {
			var _hx_tmp_25 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_25
			_hx_tmp_25 = Hx_Enum_loreline_tokentype_Unindent{}

			var _hx_tmp_26 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_26
			_hx_tmp_26 = _hx_tmp_25
			_hx_tmp_20 = this.VTable.Hx_Field_match(_hx_tmp_26, nil)
		}

		if !(_hx_tmp_20) {
			break
		}
	}

	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_RBracket{})
	literal.Hx_Field_pos = literal.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return (func() *Hx_Obj_loreline_nexpr {
		_hx_up := literal
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})()
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseObjectLiteral() *Hx_Obj_loreline_nexpr {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_currentPos()
	var fields HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = fields
	fields = HxMakeArray[*Hx_Obj_loreline_nobjectfield]()
	var _hx_tmp_0 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_0
	if (this.VTable.Hx_Field_parseBlockStart().Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_Unindent{}
	} else {
		_hx_tmp_0 = Hx_Enum_loreline_tokentype_RBrace{}
	}

	var blockEnd Hx_Enum_loreline_tokentype
	_ = blockEnd
	blockEnd = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (blockEnd == Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_1 = ((int32)(1))
	} else {
		_hx_tmp_1 = ((int32)(0))
	}

	var style int32
	_ = style
	style = _hx_tmp_1
	var _hx_tmp_2 *Hx_Obj_loreline_position
	_ = _hx_tmp_2
	if (blockEnd != Hx_Enum_loreline_tokentype_RBrace{}) {
		_hx_tmp_2 = this.VTable.Hx_Field_nextNonWhitespaceOrComment().Hx_Field_pos
	} else {
		_hx_tmp_2 = startPos
	}

	var literal *Hx_Obj_loreline_nliteral
	_ = literal
	literal = Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), _hx_tmp_2, fields, Hx_Enum_loreline_literaltype_Object{style}, nil, nil)
	func() func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
		var _hx_closure_src_757 func(node any) any
		_ = _hx_closure_src_757
		_hx_closure_src_757 = this.VTable.Hx_Field_attachComments
		return func(node *Hx_Obj_loreline_nliteral) *Hx_Obj_loreline_nliteral {
			var hx_dyncast_357 any
			_ = hx_dyncast_357
			hx_dyncast_357 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_757(((any)(node))), "Hx_Obj_loreline_nliteral")
			var _hx_tmp_3 *Hx_Obj_loreline_nliteral
			_ = _hx_tmp_3
			if hx_dyncast_357 != nil {
				_hx_tmp_3 = hx_dyncast_357.(interface {
					Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
				}).Hx_As_Hx_Obj_loreline_nliteral()
			} else {
				_hx_tmp_3 = nil
			}

			return _hx_tmp_3
		}
	}()(literal)
	var needsSeparator bool
	_ = needsSeparator
	needsSeparator = false
	for {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_4 = !this.VTable.Hx_Field_isAtEnd()
		} else {
			_hx_tmp_4 = false
		}

		if !(_hx_tmp_4) {
			break
		}

		if needsSeparator {
			for {
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				var _hx_tmp_6 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_6
				_hx_tmp_6 = Hx_Enum_loreline_tokentype_LineBreak{}

				if this.VTable.Hx_Field_match(_hx_tmp_6, nil) {
					_hx_tmp_5 = true
				} else {
					var _hx_tmp_7 bool
					_ = _hx_tmp_7
					if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
						var _hx_tmp_8 Hx_Enum_loreline_tokentype
						_ = _hx_tmp_8
						_hx_tmp_8 = Hx_Enum_loreline_tokentype_Unindent{}

						var _hx_tmp_9 Hx_Enum_loreline_tokentype
						_ = _hx_tmp_9
						_hx_tmp_9 = _hx_tmp_8
						_hx_tmp_7 = this.VTable.Hx_Field_match(_hx_tmp_9, nil)
					} else {
						_hx_tmp_7 = false
					}

					_hx_tmp_5 = _hx_tmp_7
				}

				if !(_hx_tmp_5) {
					break
				}

				needsSeparator = false
			}

			var _hx_tmp_10 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_10
			_hx_tmp_10 = Hx_Enum_loreline_tokentype_Comma{}

			if this.VTable.Hx_Field_match(_hx_tmp_10, nil) {
				needsSeparator = false
			}
		}

		for {
			var _hx_tmp_11 bool
			_ = _hx_tmp_11
			var _hx_tmp_12 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_12
			_hx_tmp_12 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_12, nil) {
				_hx_tmp_11 = true
			} else {
				var _hx_tmp_13 bool
				_ = _hx_tmp_13
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_14 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_14
					_hx_tmp_14 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_15 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_15
					_hx_tmp_15 = _hx_tmp_14
					_hx_tmp_13 = this.VTable.Hx_Field_match(_hx_tmp_15, nil)
				} else {
					_hx_tmp_13 = false
				}

				_hx_tmp_11 = _hx_tmp_13
			}

			if !(_hx_tmp_11) {
				break
			}

			needsSeparator = false
		}

		var _hx_tmp_16 bool
		_ = _hx_tmp_16
		if !this.VTable.Hx_Field_check(blockEnd) {
			_hx_tmp_16 = needsSeparator
		} else {
			_hx_tmp_16 = false
		}

		if _hx_tmp_16 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance("Expected comma or line break between fields", this.VTable.Hx_Field_currentPos())))))
		}

		for {
			var _hx_tmp_17 bool
			_ = _hx_tmp_17
			var _hx_tmp_18 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_18
			_hx_tmp_18 = Hx_Enum_loreline_tokentype_LineBreak{}

			if this.VTable.Hx_Field_match(_hx_tmp_18, nil) {
				_hx_tmp_17 = true
			} else {
				var _hx_tmp_19 bool
				_ = _hx_tmp_19
				if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
					var _hx_tmp_20 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_20
					_hx_tmp_20 = Hx_Enum_loreline_tokentype_Unindent{}

					var _hx_tmp_21 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_21
					_hx_tmp_21 = _hx_tmp_20
					_hx_tmp_19 = this.VTable.Hx_Field_match(_hx_tmp_21, nil)
				} else {
					_hx_tmp_19 = false
				}

				_hx_tmp_17 = _hx_tmp_19
			}

			if !(_hx_tmp_17) {
				break
			}
		}

		if !this.VTable.Hx_Field_check(blockEnd) {
			var _hx_tmp_22 HxArray[*Hx_Obj_loreline_nobjectfield]
			_ = _hx_tmp_22
			_hx_tmp_22 = fields
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_22, this.VTable.Hx_Field_parseObjectField()))
		}

		var prev *Hx_Obj_loreline_token
		_ = prev
		prev = this.VTable.Hx_Field_previous()
		needsSeparator = (((prev.Hx_Field_type != Hx_Enum_loreline_tokentype_Colon{}) && (prev.Hx_Field_type != Hx_Enum_loreline_tokentype_LineBreak{})) && (prev.Hx_Field_type != Hx_Enum_loreline_tokentype_Unindent{}))
	}

	for {
		var _hx_tmp_23 bool
		_ = _hx_tmp_23
		var _hx_tmp_24 Hx_Enum_loreline_tokentype
		_ = _hx_tmp_24
		_hx_tmp_24 = Hx_Enum_loreline_tokentype_LineBreak{}

		if this.VTable.Hx_Field_match(_hx_tmp_24, nil) {
			_hx_tmp_23 = true
		} else {
			var _hx_tmp_25 bool
			_ = _hx_tmp_25
			if (blockEnd != Hx_Enum_loreline_tokentype_Unindent{}) {
				var _hx_tmp_26 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_26
				_hx_tmp_26 = Hx_Enum_loreline_tokentype_Unindent{}

				var _hx_tmp_27 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_27
				_hx_tmp_27 = _hx_tmp_26
				_hx_tmp_25 = this.VTable.Hx_Field_match(_hx_tmp_27, nil)
			} else {
				_hx_tmp_25 = false
			}

			_hx_tmp_23 = _hx_tmp_25
		}

		if !(_hx_tmp_23) {
			break
		}
	}

	this.VTable.Hx_Field_expect(blockEnd)
	literal.Hx_Field_pos = literal.Hx_Field_pos.VTable.Hx_Field_extendedTo(this.VTable.Hx_Field_prevNonWhitespaceOrComment().Hx_Field_pos)
	return (func() *Hx_Obj_loreline_nexpr {
		_hx_up := literal
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})()
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseCallArguments() HxArray[*Hx_Obj_loreline_nexpr] {
	var args HxArray[*Hx_Obj_loreline_nexpr]
	_ = args
	args = HxMakeArray[*Hx_Obj_loreline_nexpr]()
	if (!this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_RParen{})) {
		for {
			for {
				var _hx_tmp_0 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_0
				_hx_tmp_0 = Hx_Enum_loreline_tokentype_LineBreak{}

				if !(this.VTable.Hx_Field_match(_hx_tmp_0, nil)) {
					break
				}
			}

			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_1
			_hx_tmp_1 = args
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_1, this.VTable.Hx_Field_parseExpression()))
			for {
				var _hx_tmp_2 Hx_Enum_loreline_tokentype
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Enum_loreline_tokentype_LineBreak{}

				if !(this.VTable.Hx_Field_match(_hx_tmp_2, nil)) {
					break
				}
			}

			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_Comma{}

			if !(this.VTable.Hx_Field_match(_hx_tmp_3, nil)) {
				break
			}
		}
	}

	this.VTable.Hx_Field_expect(Hx_Enum_loreline_tokentype_RParen{})
	return args
}

func (this *Hx_Obj_loreline_parser) Hx_Field_parseConditionExpression() *Hx_Obj_loreline_nexpr {
	var prevParsingCondition bool
	_ = prevParsingCondition
	prevParsingCondition = this.Hx_Field_parsingCondition
	this.Hx_Field_parsingCondition = true
	var expr *Hx_Obj_loreline_nexpr
	_ = expr
	{
		var hx_try_return *Hx_Obj_loreline_nexpr
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_parseerror").(*Hx_Obj_loreline_parseerror); _hx_ok {
						_ = e
						{
							this.VTable.Hx_Field_addError(e)
							expr = (func() *Hx_Obj_loreline_nexpr {
								_hx_up := Hx_Obj_loreline_nliteral_CreateInstance(this.VTable.Hx_Field_nextNodeId(((int32)(4))), this.VTable.Hx_Field_currentPos(), nil, Hx_Enum_loreline_literaltype_Null{}, nil, nil)
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_nexpr
								}
								return nil
							})()
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				expr = this.VTable.Hx_Field_parseExpression()
				if expr != nil {
					var _g *Hx_Obj_go_haxe_hxclass
					_ = _g
					_g = func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
						var _hx_closure_src_758 func(o any) *Hx_Obj_go_haxe_hxclass
						_ = _hx_closure_src_758
						_hx_closure_src_758 = Hx_Field_type_getClass
						return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
							return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_758(((any)(o)))))
						}
					}()(expr)
					if _g != nil {
						if _g == Hx_Obj_loreline_nassign_RTTI {
							this.VTable.Hx_Field_addError(Hx_Obj_loreline_parseerror_CreateInstance("Invalid condition expression: can't assign a variable here. Did you want to use '==' instead?", expr.Hx_Field_pos))
						}
					}
				}
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	this.Hx_Field_parsingCondition = prevParsingCondition
	return expr
}

func (this *Hx_Obj_loreline_parser) Hx_Field_takeConditionParens(condition *Hx_Obj_loreline_nexpr) int32 {
	if (condition != nil) && (condition.Hx_Field_parens > 0) {
		condition.Hx_Field_parens--
		return ((int32)(2))
	}

	return ((int32)(0))
}

func (this *Hx_Obj_loreline_parser) Hx_Field_match(_hx_reserved_type Hx_Enum_loreline_tokentype, advanceLineBreaks__isOptional any) bool {
	advanceLineBreaks := HxDefault[bool]()
	_ = advanceLineBreaks

	if advanceLineBreaks__isOptional != nil {
		advanceLineBreaks = HxConvert[bool](advanceLineBreaks__isOptional)
	} else {
		advanceLineBreaks = true
	}
	{
		if this.VTable.Hx_Field_check(_hx_reserved_type) {
			this.VTable.Hx_Field_advance(advanceLineBreaks)
			return true
		}

		return false
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_isAtLineBreak() bool {
	return (this.Hx_Field_lineBreakAfterToken || (this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{}))
}

func (this *Hx_Obj_loreline_parser) Hx_Field_expect(_hx_reserved_type Hx_Enum_loreline_tokentype) *Hx_Obj_loreline_token {
	if this.VTable.Hx_Field_check(_hx_reserved_type) {
		return this.VTable.Hx_Field_advance(nil)
	} else {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Expected ", Hx_Field_loreline_tokentypehelpers_toCodeString(_hx_reserved_type)), ", got ")
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		if this.VTable.Hx_Field_isAtEnd() {
			_hx_tmp_1 = "end of file"
		} else {
			_hx_tmp_1 = this.VTable.Hx_Field_currentTokenTypeToString()
		}

		var _hx_tmp_2 Hx_Typedef_loreline_tokens
		_ = _hx_tmp_2
		_hx_tmp_2 = this.Hx_Field_tokens
		var _hx_tmp_3 float64
		_ = _hx_tmp_3
		_hx_tmp_3 = ((float64)(this.Hx_Field_current))
		var _hx_reserved_error *Hx_Obj_loreline_parseerror
		_ = _hx_reserved_error
		_hx_reserved_error = Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_0, (_hx_tmp_1)), _hx_tmp_2.Get(Hx_Field_std_int(Hx_Field_math_min(_hx_tmp_3, ((float64)((this.Hx_Field_tokens.Len() - 1)))))).Hx_Field_pos)
		switch _hx_reserved_type.Hx_Field_enumIndex() {
		case 47, 49, 56:
			{
				this.VTable.Hx_Field_addError(_hx_reserved_error)
				return Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type, this.VTable.Hx_Field_currentPos())
			}

		default:
			{
			}
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(_hx_reserved_error))))
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_checkIdentifier() bool {
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	if _g.Hx_Field_enumIndex() == 20 {
		var name string
		_ = name
		name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
		return true
	} else {
		return false
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_expectIdentifier() string {
	var _g Hx_Enum_loreline_tokentype
	_ = _g
	_g = this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type
	if _g.Hx_Field_enumIndex() == 20 {
		var name string
		_ = name
		name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
		this.VTable.Hx_Field_advance(nil)
		return name
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_parseerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Expected identifier, got ", this.VTable.Hx_Field_currentTokenTypeToString()), this.VTable.Hx_Field_currentPos())))))
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_attachComments(node any) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if this.Hx_Field_pendingComments == nil {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (this.Hx_Field_pendingComments.Len() == 0)
	}

	if _hx_tmp_0 {
		return node
	}

	var hx_dyncast_358 any
	_ = hx_dyncast_358
	hx_dyncast_358 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(node, "pos"), "Hx_Obj_loreline_position")
	var _hx_tmp_1 *Hx_Obj_loreline_position
	_ = _hx_tmp_1
	if hx_dyncast_358 != nil {
		_hx_tmp_1 = hx_dyncast_358.(interface {
			Hx_As_Hx_Obj_loreline_position() *Hx_Obj_loreline_position
		}).Hx_As_Hx_Obj_loreline_position()
	} else {
		_hx_tmp_1 = nil
	}

	var nodeStart *Hx_Obj_loreline_position
	_ = nodeStart
	nodeStart = _hx_tmp_1
	var remainingComments HxArray[*Hx_Obj_loreline_comment]
	_ = remainingComments
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_pendingComments.Len()
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = this.Hx_Field_pendingComments.Get(i)
			if comment.Hx_Field_pos.Hx_Field_line < nodeStart.Hx_Field_line {
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(node, "leadingComments"), ((any)(nil))))) {
					var _hx_tmp_3 any
					_ = _hx_tmp_3
					_hx_tmp_3 = node
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "leadingComments", HxMakeArray[*Hx_Obj_loreline_comment]())
				}

				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(node, "leadingComments"), "push")
				Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any](comment).Dyn())
			} else {
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if comment.Hx_Field_pos.Hx_Field_line == nodeStart.Hx_Field_line {
					_hx_tmp_5 = true
				} else {
					_hx_tmp_5 = this.VTable.Hx_Field_isAtEnd()
				}

				if _hx_tmp_5 {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(node, "trailingComments"), ((any)(nil))))) {
						var _hx_tmp_6 any
						_ = _hx_tmp_6
						_hx_tmp_6 = node
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_6, "trailingComments", HxMakeArray[*Hx_Obj_loreline_comment]())
					}

					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(node, "trailingComments"), "push")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_7, HxMakeArray[any](comment).Dyn())
				} else {
					if remainingComments == nil {
						remainingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
					}

					_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](remainingComments, comment))
				}
			}
		}
	}

	this.Hx_Field_pendingComments = remainingComments
	return node
}

func (this *Hx_Obj_loreline_parser) Hx_Field_attachElseComments(node any, elseToken *Hx_Obj_loreline_token) any {
	if this.Hx_Field_pendingComments != nil {
		var nodeStart *Hx_Obj_loreline_position
		_ = nodeStart
		nodeStart = elseToken.Hx_Field_pos
		var remainingComments HxArray[*Hx_Obj_loreline_comment]
		_ = remainingComments
		remainingComments = HxMakeArray[*Hx_Obj_loreline_comment]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_comment]
			_ = _g1
			_g1 = this.Hx_Field_pendingComments
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var comment *Hx_Obj_loreline_comment
				_ = comment
				comment = _g1.Get(_g)
				_g++
				if comment.Hx_Field_pos.Hx_Field_line < nodeStart.Hx_Field_line {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(node, "elseLeadingComments"), ((any)(nil))))) {
						var _hx_tmp_1 any
						_ = _hx_tmp_1
						_hx_tmp_1 = node
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "elseLeadingComments", HxMakeArray[*Hx_Obj_loreline_comment]())
					}

					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(node, "elseLeadingComments"), "push")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any](comment).Dyn())
				} else {
					if comment.Hx_Field_pos.Hx_Field_line == nodeStart.Hx_Field_line {
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(node, "elseTrailingComments"), ((any)(nil))))) {
							var _hx_tmp_3 any
							_ = _hx_tmp_3
							_hx_tmp_3 = node
							Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "elseTrailingComments", HxMakeArray[*Hx_Obj_loreline_comment]())
						}

						var _hx_tmp_4 any
						_ = _hx_tmp_4
						_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(node, "elseTrailingComments"), "push")
						Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any](comment).Dyn())
					} else {
						_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](remainingComments, comment))
					}
				}
			}
		}

		this.Hx_Field_pendingComments = remainingComments
	}

	return node
}

func (this *Hx_Obj_loreline_parser) Hx_Field_synchronize() {
	this.VTable.Hx_Field_advance(nil)
	for {
		if !(!this.VTable.Hx_Field_isAtEnd()) {
			break
		}

		switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
		case 16:
			{
				if (this.VTable.Hx_Field_previous().Hx_Field_type == Hx_Enum_loreline_tokentype_RBrace{}) {
					return
				}

				this.VTable.Hx_Field_advance(nil)
			}

		case 41:
			{
				this.VTable.Hx_Field_advance(nil)
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				if (this.VTable.Hx_Field_check(Hx_Enum_loreline_tokentype_Dot{})) {
					_hx_tmp_0 = true
				} else {
					_hx_tmp_0 = Hx_Field_loreline_tokentypehelpers_isIdentifier(this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type)
				}

				if _hx_tmp_0 {
					this.VTable.Hx_Field_advance(nil)
				}

				return
			}

		case 1, 2, 3, 4, 5, 6, 47, 55:
			{
				return
			}

		default:
			{
				this.VTable.Hx_Field_advance(nil)
			}
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_requiresNewLine() bool {
	switch this.Hx_Field_tokens.Get(this.Hx_Field_current).Hx_Field_type.Hx_Field_enumIndex() {
	case 7, 41, 47:
		{
			return false
		}

	default:
		{
			return true
		}
	}
}

func (this *Hx_Obj_loreline_parser) Hx_Field_addError(_hx_reserved_error *Hx_Obj_loreline_parseerror) *Hx_Obj_loreline_parseerror {
	if this.Hx_Field_errors == nil {
		this.Hx_Field_errors = HxMakeArray[*Hx_Obj_loreline_parseerror]()
	}

	if !(Hx_Array_Contains[*Hx_Obj_loreline_parseerror](this.Hx_Field_errors, _hx_reserved_error)) {
		_ = (Hx_Array_Push[*Hx_Obj_loreline_parseerror](this.Hx_Field_errors, _hx_reserved_error))
	}

	return _hx_reserved_error
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatparam]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nchoiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_parseerror]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_token]()
	hxRegisterArrayConstructor[HxArray[*Hx_Obj_loreline_token]]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lstringattachment]()
	hxRegisterArrayConstructor[string]()
}
