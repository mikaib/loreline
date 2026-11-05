package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_printer() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_printer() {
}

type Hx_Obj_VTable_loreline_printer interface {
	Hx_Field_printFunctionDecl(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl)
	Hx_Field_printTextStatement(text *Hx_Obj_loreline_ntextstatement)
	Hx_Field_printChoiceStatement(choice *Hx_Obj_loreline_nchoicestatement)
	Hx_Field_printImportStatement(imp *Hx_Obj_loreline_nimportstatement)
	Hx_Field_needsEmptyLines(node *Hx_Obj_loreline_astnode) bool
	Hx_Field_printTransition(trans *Hx_Obj_loreline_ntransition, sameLine bool)
	Hx_Field_clear()
	Hx_Field_space() *Hx_Obj_loreline_printer
	Hx_Field_printNode(node *Hx_Obj_loreline_node, sameLine__isOptional any)
	Hx_Field_printDialogueStatement(dialogue *Hx_Obj_loreline_ndialoguestatement)
	Hx_Field_printInsertion(insert *Hx_Obj_loreline_ninsertion)
	Hx_Field_printTargetExprAndArgs(targetExpr *Hx_Obj_loreline_nexpr, args HxArray[*Hx_Obj_loreline_nexpr])
	Hx_Field_writeUnquotedRaw(text string)
	Hx_Field_printArrayAccess(access *Hx_Obj_loreline_narrayaccess)
	Hx_Field__writeln(s__isOptional any) *Hx_Obj_loreline_printer
	Hx_Field_toString() string
	Hx_Field_printScript(script *Hx_Obj_loreline_script)
	Hx_Field_printIfStatement(ifStmt *Hx_Obj_loreline_nifstatement, isElseIf__isOptional any)
	Hx_Field_printAlternative(alt *Hx_Obj_loreline_nalternative)
	Hx_Field_printBeatCall(call *Hx_Obj_loreline_nbeatcall)
	Hx_Field_printAccess(access *Hx_Obj_loreline_naccess)
	Hx_Field_tab() *Hx_Obj_loreline_printer
	Hx_Field__line(s string) *Hx_Obj_loreline_printer
	Hx_Field_print(node *Hx_Obj_loreline_node) string
	Hx_Field_printStateDecl(state *Hx_Obj_loreline_nstatedecl)
	Hx_Field_printBeatDecl(beat *Hx_Obj_loreline_nbeatdecl)
	Hx_Field_printBinary(binary *Hx_Obj_loreline_nbinary)
	Hx_Field_printUnary(unary *Hx_Obj_loreline_nunary)
	Hx_Field_newline() *Hx_Obj_loreline_printer
	Hx_Field_printOperand(expr *Hx_Obj_loreline_nexpr, minPrecedence int32)
	Hx_Field_printCharacterDecl(char *Hx_Obj_loreline_ncharacterdecl)
	Hx_Field_printWhenStatement(when *Hx_Obj_loreline_nwhenstatement)
	Hx_Field_printLiteral(lit *Hx_Obj_loreline_nliteral)
	Hx_Field_printAssignment(assign *Hx_Obj_loreline_nassign)
	Hx_Field_getOperator(op Hx_Enum_loreline_tokentype) string
	Hx_Field_unindent()
	Hx_Field_printStringLiteralAsReference(str *Hx_Obj_loreline_nstringliteral) string
	Hx_Field_printImport(imp *Hx_Obj_loreline_nimportstatement)
	Hx_Field_printWhenRule(rule *Hx_Obj_loreline_nwhenrule)
	Hx_Field_printChoiceOption(option *Hx_Obj_loreline_nchoiceoption)
	Hx_Field_hasRealNewlines(node *Hx_Obj_loreline_node) bool
	Hx_Field_printTrailingComments(node *Hx_Obj_loreline_astnode)
	Hx_Field_writeRawWithContinuation(text string, stripIndent bool)
	Hx_Field_printInLineExpression(expr *Hx_Obj_loreline_nexpr, parens bool)
	Hx_Field_printStringLiteralAsText(str *Hx_Obj_loreline_nstringliteral) string
	Hx_Field_printLeadingComments(node *Hx_Obj_loreline_astnode)
	Hx_Field_printHashTrailingComments(node *Hx_Obj_loreline_astnode)
	Hx_Field_printStringLiteral(str *Hx_Obj_loreline_nstringliteral)
	Hx_Field_writeQuotedRaw(text string)
	Hx_Field_printCall(call *Hx_Obj_loreline_ncall)
	Hx_Field_printTernary(ternary *Hx_Obj_loreline_nternary)
	Hx_Field_isSimpleInterpolationExpr(expr *Hx_Obj_loreline_nexpr) bool
	Hx_Field_indent()
	Hx_Field_write(s string) *Hx_Obj_loreline_printer
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_printer struct {
	VTable                    Hx_Obj_VTable_loreline_printer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field__level           int32
	Hx_Field__prevLevel       int32
	Hx_Field__buf             Hx_Typedef_loreline_utf8buf
	Hx_Field__beginLine       int32
	Hx_Field__lastChar        int32
	Hx_Field__lastVisibleChar int32
	Hx_Field__noLn            int32
	Hx_Field__pendingSpace    bool
	Hx_Field__numEmptyLines   int32
	Hx_Field__indent          string
	Hx_Field__newline         string
	Hx_Field_enableComments   bool
}

func Hx_Obj_loreline_printer_CreateEmptyInstance() *Hx_Obj_loreline_printer {
	obj := &Hx_Obj_loreline_printer{}
	obj.VTable = obj
	obj.Hx_Field__level = HxDefault[int32]()
	obj.Hx_Field__prevLevel = HxDefault[int32]()
	obj.Hx_Field__buf = HxDefault[Hx_Typedef_loreline_utf8buf]()
	obj.Hx_Field__beginLine = HxDefault[int32]()
	obj.Hx_Field__lastChar = HxDefault[int32]()
	obj.Hx_Field__lastVisibleChar = HxDefault[int32]()
	obj.Hx_Field__noLn = HxDefault[int32]()
	obj.Hx_Field__pendingSpace = HxDefault[bool]()
	obj.Hx_Field__numEmptyLines = HxDefault[int32]()
	obj.Hx_Field__indent = HxDefault[string]()
	obj.Hx_Field__newline = HxDefault[string]()
	obj.Hx_Field_enableComments = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_printer_CreateInstance(indent__isOptional any, newline__isOptional any) *Hx_Obj_loreline_printer {
	obj := Hx_Obj_loreline_printer_CreateEmptyInstance()
	obj.Hx_New(indent__isOptional, newline__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_printer) Hx_New(indent__isOptional any, newline__isOptional any) {
	{
		indent := HxDefault[string]()
		_ = indent

		if indent__isOptional != nil {
			indent = HxConvert[string](indent__isOptional)
		} else {
			indent = "  "
		}
		newline := HxDefault[string]()
		_ = newline

		if newline__isOptional != nil {
			newline = HxConvert[string](newline__isOptional)
		} else {
			newline = "\n"
		}
		{
			this.Hx_Field__pendingSpace = false
		}
		{
			this.Hx_Field_enableComments = true
		}
		{
			this.Hx_Field__indent = indent
			this.Hx_Field__newline = newline
			this.Hx_Field__beginLine = 0
			this.Hx_Field__lastChar = -1
			{
				this.Hx_Field__noLn = 0
				this.Hx_Field__pendingSpace = false
				this.Hx_Field__level = 0
				this.Hx_Field__prevLevel = 0
				this.Hx_Field__numEmptyLines = 0
				this.Hx_Field__buf = Hx_Obj_stringbuf_CreateInstance()
			}
		}

	}
}

var Hx_Obj_loreline_printer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Printer",
	HxMakeArray[string]("precedence", "isWordNotAfterIs"),
	HxMakeArray[string]("_level", "_prevLevel", "_buf", "_beginLine", "_lastChar", "_lastVisibleChar", "_noLn", "_pendingSpace", "_numEmptyLines", "_indent", "_newline", "enableComments", "clear", "indent", "unindent", "write", "writeln", "_writeln", "newline", "space", "tab", "line", "_line", "toString", "printStringLiteralAsReference", "printStringLiteralAsText", "print", "printNode", "printOperand", "printLeadingComments", "printHashTrailingComments", "printTrailingComments", "printImport", "printScript", "printStateDecl", "printCharacterDecl", "printBeatDecl", "printFunctionDecl", "printTextStatement", "printDialogueStatement", "printChoiceStatement", "printWhenStatement", "printWhenRule", "printImportStatement", "printChoiceOption", "printIfStatement", "printAlternative", "needsEmptyLines", "printTransition", "printInsertion", "printBeatCall", "printTargetExprAndArgs", "printStringLiteral", "writeUnquotedRaw", "writeQuotedRaw", "writeRawWithContinuation", "printLiteral", "printAccess", "printArrayAccess", "printCall", "printBinary", "printUnary", "printTernary", "printAssignment", "printInLineExpression", "isSimpleInterpolationExpr", "hasRealNewlines", "getOperator"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_printer_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_printer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_printer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "precedence":
			return Hx_Field_loreline_printer_precedence
		case "isWordNotAfterIs":
			return Hx_Field_loreline_printer_isWordNotAfterIs
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_printer_RTTI
}

func (this *Hx_Obj_loreline_printer) Hx_As_Hx_Obj_loreline_printer() *Hx_Obj_loreline_printer {
	return this
}

func (this *Hx_Obj_loreline_printer) Hx_Field_clear() {
	this.Hx_Field__noLn = 0
	this.Hx_Field__pendingSpace = false
	this.Hx_Field__level = 0
	this.Hx_Field__prevLevel = 0
	this.Hx_Field__numEmptyLines = 0
	this.Hx_Field__buf = Hx_Obj_stringbuf_CreateInstance()
}

func (this *Hx_Obj_loreline_printer) Hx_Field_indent() {
	this.Hx_Field__level++
}

func (this *Hx_Obj_loreline_printer) Hx_Field_unindent() {
	this.Hx_Field__level--
}

func (this *Hx_Obj_loreline_printer) Hx_Field_write(s string) *Hx_Obj_loreline_printer {
	if int32(utf8.RuneCountInString(s)) > 0 {
		if this.Hx_Field__beginLine > 0 {
			this.VTable.Hx_Field_tab()
			this.Hx_Field__beginLine = 0
		}

		if this.Hx_Field__pendingSpace {
			this.Hx_Field__pendingSpace = false
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(s, 0))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			var first int32
			_ = first
			first = _hx_tmp_0
			if ((first != 32) && (first != 10)) && (first != 13) {
				func() func(x string) {
					var _hx_closure_src_438 func(x any)
					_ = _hx_closure_src_438
					_hx_closure_src_438 = this.Hx_Field__buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_438(((any)(x)))
					}
				}()(" ")
			}
		}

		func() func(x string) {
			var _hx_closure_src_439 func(x any)
			_ = _hx_closure_src_439
			_hx_closure_src_439 = this.Hx_Field__buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_439(((any)(x)))
			}
		}()(s)
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, (int32(utf8.RuneCountInString(s)) - 1)))
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (c).Valid == false {
			_hx_tmp_2 = -1
		} else {
			_hx_tmp_2 = (c).Value
		}

		this.Hx_Field__lastChar = _hx_tmp_2
		var i int32
		_ = i
		i = (int32(utf8.RuneCountInString(s)) - 1)
		for i >= 0 {
			var tmp_c_1 HxNullable[int32]
			_ = tmp_c_1
			tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(s, i))
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
			if (((c1 != 32) && (c1 != 10)) && (c1 != 13)) && (c1 != 9) {
				this.Hx_Field__lastVisibleChar = c1
				break
			}

			i--
		}
	}

	return this
}

func (this *Hx_Obj_loreline_printer) Hx_Field__writeln(s__isOptional any) *Hx_Obj_loreline_printer {
	s := HxDefault[string]()
	_ = s

	if s__isOptional != nil {
		s = HxConvert[string](s__isOptional)
	} else {
		s = ""
	}
	{
		this.VTable.Hx_Field_write(s)
		return this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_newline() *Hx_Obj_loreline_printer {
	if this.Hx_Field__noLn == 0 {
		this.Hx_Field__pendingSpace = false
	}

	if (this.Hx_Field__beginLine < 2) && (this.Hx_Field__noLn == 0) {
		if this.Hx_Field__beginLine == 1 {
			this.Hx_Field__numEmptyLines++
		}

		func() func(x string) {
			var _hx_closure_src_440 func(x any)
			_ = _hx_closure_src_440
			_hx_closure_src_440 = this.Hx_Field__buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_440(((any)(x)))
			}
		}()(this.Hx_Field__newline)
		this.Hx_Field__beginLine++
	}

	return this
}

func (this *Hx_Obj_loreline_printer) Hx_Field_space() *Hx_Obj_loreline_printer {
	this.Hx_Field__pendingSpace = true
	return this
}

func (this *Hx_Obj_loreline_printer) Hx_Field_tab() *Hx_Obj_loreline_printer {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field__level
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			func() func(x string) {
				var _hx_closure_src_441 func(x any)
				_ = _hx_closure_src_441
				_hx_closure_src_441 = this.Hx_Field__buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_441(((any)(x)))
				}
			}()(this.Hx_Field__indent)
		}
	}

	return this
}

func (this *Hx_Obj_loreline_printer) Hx_Field__line(s string) *Hx_Obj_loreline_printer {
	this.VTable.Hx_Field_tab()
	func() func(x string) {
		var _hx_closure_src_442 func(x any)
		_ = _hx_closure_src_442
		_hx_closure_src_442 = this.Hx_Field__buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_442(((any)(x)))
		}
	}()(s)
	return this.VTable.Hx_Field_newline()
}

func (this *Hx_Obj_loreline_printer) Hx_Field_toString() string {
	return this.Hx_Field__buf.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printStringLiteralAsReference(str *Hx_Obj_loreline_nstringliteral) string {
	var saved bool
	_ = saved
	saved = this.Hx_Field_enableComments
	this.Hx_Field_enableComments = false
	{
		this.Hx_Field__noLn = 0
		this.Hx_Field__pendingSpace = false
		this.Hx_Field__level = 0
		this.Hx_Field__prevLevel = 0
		this.Hx_Field__numEmptyLines = 0
		this.Hx_Field__buf = Hx_Obj_stringbuf_CreateInstance()
	}

	this.Hx_Field__beginLine = 1
	this.VTable.Hx_Field_printStringLiteral(str)
	this.Hx_Field_enableComments = saved
	var result string
	_ = result
	result = Hx_Field_stringtools_trim(this.VTable.Hx_Field_toString())
	result = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(result, "\n")), "\\n"))
	return result
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printStringLiteralAsText(str *Hx_Obj_loreline_nstringliteral) string {
	var saved bool
	_ = saved
	saved = this.Hx_Field_enableComments
	this.Hx_Field_enableComments = false
	{
		this.Hx_Field__noLn = 0
		this.Hx_Field__pendingSpace = false
		this.Hx_Field__level = 0
		this.Hx_Field__prevLevel = 0
		this.Hx_Field__numEmptyLines = 0
		this.Hx_Field__buf = Hx_Obj_stringbuf_CreateInstance()
	}

	this.Hx_Field__beginLine = 1
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nstringpart]
		_ = _g1
		_g1 = str.Hx_Field_parts
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = _g1.Get(_g)
			_g++
			{
				var _g2 Hx_Enum_loreline_stringparttype
				_ = _g2
				_g2 = part.Hx_Field_partType
				switch _g2.Hx_Field_enumIndex() {
				case 0:
					{
						var text string
						_ = text
						text = _g2.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						if str.Hx_Field_quotes == ((int32)(1)) {
							this.VTable.Hx_Field_writeQuotedRaw(text)
						} else {
							this.VTable.Hx_Field_writeUnquotedRaw(text)
						}
					}

				case 1:
					{
						var expr *Hx_Obj_loreline_nexpr
						_ = expr
						expr = _g2.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
						{
							var canBeSimple bool
							_ = canBeSimple
							canBeSimple = this.VTable.Hx_Field_isSimpleInterpolationExpr(expr)
							this.VTable.Hx_Field_write("$")
							if !canBeSimple {
								this.VTable.Hx_Field_write("{")
							}

							var _hx_tmp_1 *Hx_Obj_loreline_node
							_ = _hx_tmp_1
							_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
								_hx_up := expr
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
							this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
							if !canBeSimple {
								this.VTable.Hx_Field_write("}")
							}
						}
					}

				case 2:
					{
						var closing bool
						_ = closing
						closing = _g2.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_closing
						var content *Hx_Obj_loreline_nstringliteral
						_ = content
						content = _g2.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_expr
						{
							var _hx_tmp_2 string
							_ = _hx_tmp_2
							if closing {
								_hx_tmp_2 = "</"
							} else {
								_hx_tmp_2 = "<"
							}

							this.VTable.Hx_Field_write(_hx_tmp_2)
							this.VTable.Hx_Field_printStringLiteral(content)
							this.VTable.Hx_Field_write(">")
						}
					}

				}
			}
		}
	}

	this.Hx_Field_enableComments = saved
	var result string
	_ = result
	result = Hx_Field_stringtools_trim(this.VTable.Hx_Field_toString())
	result = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(result, "\n")), "\\n"))
	return result
}

func (this *Hx_Obj_loreline_printer) Hx_Field_print(node *Hx_Obj_loreline_node) string {
	{
		this.Hx_Field__noLn = 0
		this.Hx_Field__pendingSpace = false
		this.Hx_Field__level = 0
		this.Hx_Field__prevLevel = 0
		this.Hx_Field__numEmptyLines = 0
		this.Hx_Field__buf = Hx_Obj_stringbuf_CreateInstance()
	}

	this.Hx_Field__beginLine = 1
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = node
	this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	return this.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printNode(node *Hx_Obj_loreline_node, sameLine__isOptional any) {
	sameLine := HxDefault[bool]()
	_ = sameLine

	if sameLine__isOptional != nil {
		sameLine = HxConvert[bool](sameLine__isOptional)
	} else {
		sameLine = false
	}
	{
		if node == nil {
			return
		}

		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nexpr_RTTI) {
			_hx_tmp_0 = node.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr().Hx_Field_parens
		} else {
			_hx_tmp_0 = 0
		}

		var parens int32
		_ = parens
		parens = _hx_tmp_0
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = parens
			for _g < _g1 {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_1
				this.VTable.Hx_Field_write("(")
			}
		}

		{
			var _g2 *Hx_Obj_go_haxe_hxclass
			_ = _g2
			_g2 = func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_443 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_443
				_hx_closure_src_443 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_443(((any)(o)))))
				}
			}()(node)
			if _g2 == nil {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unsupported node type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
					var _hx_closure_src_446 func(o any) *Hx_Obj_go_haxe_hxclass
					_ = _hx_closure_src_446
					_hx_closure_src_446 = Hx_Field_type_getClass
					return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
						return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_446(((any)(o)))))
					}
				}()(node)))))))))
			} else {
				switch _g2 {
				case Hx_Obj_loreline_naccess_RTTI:
					{
						this.VTable.Hx_Field_printAccess(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
						}).Hx_As_Hx_Obj_loreline_naccess())
					}

				case Hx_Obj_loreline_nalternative_RTTI:
					{
						this.VTable.Hx_Field_printAlternative(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
						}).Hx_As_Hx_Obj_loreline_nalternative())
					}

				case Hx_Obj_loreline_narrayaccess_RTTI:
					{
						this.VTable.Hx_Field_printArrayAccess(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
						}).Hx_As_Hx_Obj_loreline_narrayaccess())
					}

				case Hx_Obj_loreline_nassign_RTTI:
					{
						this.VTable.Hx_Field_printAssignment(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
						}).Hx_As_Hx_Obj_loreline_nassign())
					}

				case Hx_Obj_loreline_nbeatcall_RTTI:
					{
						this.VTable.Hx_Field_printBeatCall(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
						}).Hx_As_Hx_Obj_loreline_nbeatcall())
					}

				case Hx_Obj_loreline_nbeatdecl_RTTI:
					{
						this.VTable.Hx_Field_printBeatDecl(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl())
					}

				case Hx_Obj_loreline_nbinary_RTTI:
					{
						this.VTable.Hx_Field_printBinary(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
						}).Hx_As_Hx_Obj_loreline_nbinary())
					}

				case Hx_Obj_loreline_ncall_RTTI:
					{
						this.VTable.Hx_Field_printCall(node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall())
					}

				case Hx_Obj_loreline_ncharacterdecl_RTTI:
					{
						this.VTable.Hx_Field_printCharacterDecl(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
						}).Hx_As_Hx_Obj_loreline_ncharacterdecl())
					}

				case Hx_Obj_loreline_nchoiceoption_RTTI:
					{
						this.VTable.Hx_Field_printChoiceOption(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
						}).Hx_As_Hx_Obj_loreline_nchoiceoption())
					}

				case Hx_Obj_loreline_nchoicestatement_RTTI:
					{
						this.VTable.Hx_Field_printChoiceStatement(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
						}).Hx_As_Hx_Obj_loreline_nchoicestatement())
					}

				case Hx_Obj_loreline_ndialoguestatement_RTTI:
					{
						this.VTable.Hx_Field_printDialogueStatement(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
						}).Hx_As_Hx_Obj_loreline_ndialoguestatement())
					}

				case Hx_Obj_loreline_nfunctiondecl_RTTI:
					{
						this.VTable.Hx_Field_printFunctionDecl(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
						}).Hx_As_Hx_Obj_loreline_nfunctiondecl())
					}

				case Hx_Obj_loreline_nifstatement_RTTI:
					{
						var _hx_tmp_2 *Hx_Obj_loreline_nifstatement
						_ = _hx_tmp_2
						_hx_tmp_2 = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
						}).Hx_As_Hx_Obj_loreline_nifstatement()
						this.VTable.Hx_Field_printIfStatement(_hx_tmp_2, nil)
					}

				case Hx_Obj_loreline_nimportstatement_RTTI:
					{
						this.VTable.Hx_Field_printImportStatement(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
						}).Hx_As_Hx_Obj_loreline_nimportstatement())
					}

				case Hx_Obj_loreline_ninsertion_RTTI:
					{
						this.VTable.Hx_Field_printInsertion(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
						}).Hx_As_Hx_Obj_loreline_ninsertion())
					}

				case Hx_Obj_loreline_nliteral_RTTI:
					{
						this.VTable.Hx_Field_printLiteral(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
						}).Hx_As_Hx_Obj_loreline_nliteral())
					}

				case Hx_Obj_loreline_nstatedecl_RTTI:
					{
						this.VTable.Hx_Field_printStateDecl(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
						}).Hx_As_Hx_Obj_loreline_nstatedecl())
					}

				case Hx_Obj_loreline_nstringliteral_RTTI:
					{
						this.VTable.Hx_Field_printStringLiteral(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
						}).Hx_As_Hx_Obj_loreline_nstringliteral())
					}

				case Hx_Obj_loreline_nternary_RTTI:
					{
						this.VTable.Hx_Field_printTernary(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nternary() *Hx_Obj_loreline_nternary
						}).Hx_As_Hx_Obj_loreline_nternary())
					}

				case Hx_Obj_loreline_ntextstatement_RTTI:
					{
						this.VTable.Hx_Field_printTextStatement(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
						}).Hx_As_Hx_Obj_loreline_ntextstatement())
					}

				case Hx_Obj_loreline_ntransition_RTTI:
					{
						this.VTable.Hx_Field_printTransition(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
						}).Hx_As_Hx_Obj_loreline_ntransition(), sameLine)
					}

				case Hx_Obj_loreline_nunary_RTTI:
					{
						this.VTable.Hx_Field_printUnary(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
						}).Hx_As_Hx_Obj_loreline_nunary())
					}

				case Hx_Obj_loreline_nwhenrule_RTTI:
					{
						this.VTable.Hx_Field_printWhenRule(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
						}).Hx_As_Hx_Obj_loreline_nwhenrule())
					}

				case Hx_Obj_loreline_nwhenstatement_RTTI:
					{
						this.VTable.Hx_Field_printWhenStatement(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nwhenstatement() *Hx_Obj_loreline_nwhenstatement
						}).Hx_As_Hx_Obj_loreline_nwhenstatement())
					}

				case Hx_Obj_loreline_script_RTTI:
					{
						this.VTable.Hx_Field_printScript(node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_script() *Hx_Obj_loreline_script
						}).Hx_As_Hx_Obj_loreline_script())
					}

				default:
					{
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unsupported node type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
							var _hx_closure_src_447 func(o any) *Hx_Obj_go_haxe_hxclass
							_ = _hx_closure_src_447
							_hx_closure_src_447 = Hx_Field_type_getClass
							return func(o *Hx_Obj_loreline_node) *Hx_Obj_go_haxe_hxclass {
								return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_447(((any)(o)))))
							}
						}()(node)))))))))
					}
				}
			}
		}

		{
			var _g3 int32
			_ = _g3
			_g3 = 0
			var _g4 int32
			_ = _g4
			_g4 = parens
			for _g3 < _g4 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g3
				_g3 = (_g3 + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_3
				this.VTable.Hx_Field_write(")")
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printOperand(expr *Hx_Obj_loreline_nexpr, minPrecedence int32) {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (expr != nil) && (expr.Hx_Field_parens == 0) {
		_hx_tmp_0 = (Hx_Field_loreline_printer_precedence(expr) < minPrecedence)
	} else {
		_hx_tmp_0 = false
	}

	var extra bool
	_ = extra
	extra = _hx_tmp_0
	if extra {
		this.VTable.Hx_Field_write("(")
	}

	var _hx_tmp_1 *Hx_Obj_loreline_node
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
		_hx_up := expr
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
	if extra {
		this.VTable.Hx_Field_write(")")
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printLeadingComments(node *Hx_Obj_loreline_astnode) {
	if this.Hx_Field_enableComments && (node.Hx_Field_leadingComments != nil) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = node.Hx_Field_leadingComments
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
			if comment.Hx_Field_multiline {
				this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("/*", comment.Hx_Field_content), "*/"))
				this.VTable.Hx_Field_newline()
			} else {
				if comment.Hx_Field_isHash {
					this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("#", comment.Hx_Field_content))
					this.VTable.Hx_Field_newline()
				} else {
					this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("//", comment.Hx_Field_content))
					this.VTable.Hx_Field_newline()
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printHashTrailingComments(node *Hx_Obj_loreline_astnode) {
	if this.Hx_Field_enableComments && (node.Hx_Field_trailingComments != nil) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = node.Hx_Field_trailingComments
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
			if comment.Hx_Field_isHash {
				if (this.Hx_Field__lastChar != 32) && (this.Hx_Field__beginLine == 0) {
					this.VTable.Hx_Field_write(" ")
				}

				{
					this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("#", comment.Hx_Field_content))
					this.VTable.Hx_Field_newline()
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printTrailingComments(node *Hx_Obj_loreline_astnode) {
	if this.Hx_Field_enableComments && (node.Hx_Field_trailingComments != nil) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = node.Hx_Field_trailingComments
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
			if (this.Hx_Field__lastChar != 32) && (this.Hx_Field__beginLine == 0) {
				this.VTable.Hx_Field_write(" ")
			}

			if comment.Hx_Field_multiline {
				this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("/*", comment.Hx_Field_content), "*/"))
				this.VTable.Hx_Field_space()
			} else {
				if comment.Hx_Field_isHash {
					this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("#", comment.Hx_Field_content))
					this.VTable.Hx_Field_newline()
				} else {
					this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("//", comment.Hx_Field_content))
					this.VTable.Hx_Field_newline()
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printImport(imp *Hx_Obj_loreline_nimportstatement) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := imp
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	{
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("import \"", Hx_Field_std_string(imp.Hx_Field_path)), "\""))
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := imp
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printScript(script *Hx_Obj_loreline_script) {
	var _g int32
	_ = _g
	_g = 0
	var _g1 HxArray[*Hx_Obj_loreline_astnode]
	_ = _g1
	_g1 = script.Hx_Field_body
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = _g
		if !(_hx_tmp_0 < _g1.Len()) {
			break
		}

		var decl *Hx_Obj_loreline_astnode
		_ = decl
		decl = _g1.Get(_g)
		_g++
		this.Hx_Field__prevLevel = this.Hx_Field__level
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := decl
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printStateDecl(state *Hx_Obj_loreline_nstatedecl) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := state
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if state.Hx_Field_temporary {
		this.VTable.Hx_Field_write("new ")
	}

	this.VTable.Hx_Field_write("state")
	this.VTable.Hx_Field_space()
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := state
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if state.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("{")
		this.VTable.Hx_Field_newline()
	} else {
		this.VTable.Hx_Field_newline()
	}

	this.Hx_Field__level++
	var first bool
	_ = first
	first = true
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = state.Hx_Field_fields
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var field *Hx_Obj_loreline_nobjectfield
			_ = field
			field = _g1.Get(_g)
			_g++
			if !first {
				if this.Hx_Field__beginLine == 0 {
					this.VTable.Hx_Field_newline()
				}
			}

			first = false
			this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
				_hx_up := field
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())
			this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", field.Hx_Field_name), ":"))
			this.VTable.Hx_Field_space()
			this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
				_hx_up := field
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := field.Hx_Field_value
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
		}
	}

	this.VTable.Hx_Field_newline()
	this.Hx_Field__level--
	if state.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printCharacterDecl(char *Hx_Obj_loreline_ncharacterdecl) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := char
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("character ", char.Hx_Field_name))
	this.VTable.Hx_Field_space()
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := char
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if char.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("{")
		this.VTable.Hx_Field_newline()
	} else {
		this.VTable.Hx_Field_newline()
	}

	this.Hx_Field__level++
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = char.Hx_Field_fields
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var prop *Hx_Obj_loreline_nobjectfield
			_ = prop
			prop = _g1.Get(_g)
			_g++
			this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
				_hx_up := prop
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())
			this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", prop.Hx_Field_name), ":"))
			this.VTable.Hx_Field_space()
			this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
				_hx_up := prop
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := prop.Hx_Field_value
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
			if this.Hx_Field__beginLine == 0 {
				this.VTable.Hx_Field_newline()
			}
		}
	}

	this.Hx_Field__level--
	if char.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printBeatDecl(beat *Hx_Obj_loreline_nbeatdecl) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := beat
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("beat ", beat.Hx_Field_name))
	if beat.Hx_Field_params != nil {
		var savedComments bool
		_ = savedComments
		savedComments = this.Hx_Field_enableComments
		this.Hx_Field_enableComments = false
		this.VTable.Hx_Field_write("(")
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = beat.Hx_Field_params.Len()
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				if i > 0 {
					this.VTable.Hx_Field_write(", ")
				}

				var param *Hx_Obj_loreline_nbeatparam
				_ = param
				param = beat.Hx_Field_params.Get(i)
				this.VTable.Hx_Field_write(param.Hx_Field_name)
				if param.Hx_Field_defaultValue != nil {
					this.VTable.Hx_Field_write(" = ")
					var _hx_tmp_1 *Hx_Obj_loreline_node
					_ = _hx_tmp_1
					_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
						_hx_up := param.Hx_Field_defaultValue
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
				}
			}
		}

		this.VTable.Hx_Field_write(")")
		this.Hx_Field_enableComments = savedComments
	}

	this.VTable.Hx_Field_space()
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := beat
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if beat.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("{")
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.Hx_Field__level++
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = beat.Hx_Field_body.Len()
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			this.Hx_Field__prevLevel = this.Hx_Field__level
			var _hx_tmp_3 *Hx_Obj_loreline_node
			_ = _hx_tmp_3
			_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
				_hx_up := beat.Hx_Field_body.Get(i)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_3, nil)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = i
			if _hx_tmp_4 < (beat.Hx_Field_body.Len() - 1) {
				this.VTable.Hx_Field_newline()
			}
		}
	}

	this.Hx_Field__level--
	if this.Hx_Field__beginLine == 0 {
		this.VTable.Hx_Field_newline()
	}

	if this.Hx_Field__lastVisibleChar != 125 {
		this.VTable.Hx_Field_newline()
	}

	if beat.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printFunctionDecl(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := _hx_reserved_func
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write(_hx_reserved_func.Hx_Field_code)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := _hx_reserved_func
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printTextStatement(text *Hx_Obj_loreline_ntextstatement) {
	if this.Hx_Field__level == this.Hx_Field__prevLevel {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := text
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := text.Hx_Field_content
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	if text.Hx_Field_condition != nil {
		this.VTable.Hx_Field_write(" if ")
		this.VTable.Hx_Field_printInLineExpression(text.Hx_Field_condition, (text.Hx_Field_conditionStyle == ((int32)(2))))
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := text
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printDialogueStatement(dialogue *Hx_Obj_loreline_ndialoguestatement) {
	if this.Hx_Field__level == this.Hx_Field__prevLevel {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := dialogue
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dialogue.Hx_Field_content != nil {
		_hx_tmp_0 = this.VTable.Hx_Field_hasRealNewlines((func() *Hx_Obj_loreline_node {
			_hx_up := dialogue.Hx_Field_content
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", dialogue.Hx_Field_character), ":"))
		this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := dialogue
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		this.VTable.Hx_Field_newline()
		this.Hx_Field__level++
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := dialogue.Hx_Field_content
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
		this.Hx_Field__level--
	} else {
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", dialogue.Hx_Field_character), ": "))
		this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := dialogue
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		var _hx_tmp_2 *Hx_Obj_loreline_node
		_ = _hx_tmp_2
		_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
			_hx_up := dialogue.Hx_Field_content
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_2, nil)
	}

	if dialogue.Hx_Field_condition != nil {
		this.VTable.Hx_Field_write(" if ")
		this.VTable.Hx_Field_printInLineExpression(dialogue.Hx_Field_condition, (dialogue.Hx_Field_conditionStyle == ((int32)(2))))
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printChoiceStatement(choice *Hx_Obj_loreline_nchoicestatement) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := choice
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write("choice")
	this.VTable.Hx_Field_space()
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := choice
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if choice.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("{")
		this.VTable.Hx_Field_newline()
	} else {
		this.VTable.Hx_Field_newline()
	}

	this.Hx_Field__level++
	if choice.Hx_Field_options.Len() > 0 {
		this.Hx_Field__noLn++
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_nchoiceoption]
			_ = _g1
			_g1 = choice.Hx_Field_options
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var option *Hx_Obj_loreline_nchoiceoption
				_ = option
				option = _g1.Get(_g)
				_g++
				var _hx_tmp_1 *Hx_Obj_loreline_node
				_ = _hx_tmp_1
				_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
					_hx_up := option
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})()
				this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
				this.VTable.Hx_Field_newline()
			}
		}
	}

	this.Hx_Field__level--
	if choice.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printWhenStatement(when *Hx_Obj_loreline_nwhenstatement) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := when
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write("when")
	if (when.Hx_Field_strategy).Valid != false {
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(" ", (when.Hx_Field_strategy).Value))
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := when
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if when.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write(" {")
		this.VTable.Hx_Field_newline()
	} else {
		if this.Hx_Field__beginLine == 0 {
			this.VTable.Hx_Field_newline()
		}
	}

	this.Hx_Field__level++
	if when.Hx_Field_rules.Len() > 0 {
		this.Hx_Field__noLn++
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_nwhenrule]
			_ = _g1
			_g1 = when.Hx_Field_rules
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var rule *Hx_Obj_loreline_nwhenrule
				_ = rule
				rule = _g1.Get(_g)
				_g++
				var _hx_tmp_1 *Hx_Obj_loreline_node
				_ = _hx_tmp_1
				_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
					_hx_up := rule
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})()
				this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
				this.VTable.Hx_Field_newline()
			}
		}
	}

	this.Hx_Field__level--
	if when.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printWhenRule(rule *Hx_Obj_loreline_nwhenrule) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.Hx_Field__noLn = 0
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if rule.Hx_Field_insertion != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := rule.Hx_Field_insertion
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
		if rule.Hx_Field_insertionCondition != nil {
			this.VTable.Hx_Field_write(" if ")
			this.VTable.Hx_Field_printInLineExpression(rule.Hx_Field_insertionCondition, (rule.Hx_Field_insertionConditionStyle == ((int32)(2))))
		}

		this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := rule
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		this.VTable.Hx_Field_newline()
		return
	}

	if rule.Hx_Field_once {
		this.VTable.Hx_Field_write("- ")
	}

	if rule.Hx_Field_condition == nil {
		this.VTable.Hx_Field_write("true")
	} else {
		this.VTable.Hx_Field_printInLineExpression(rule.Hx_Field_condition, false)
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if rule.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write(" {")
		this.VTable.Hx_Field_newline()
	} else {
		if this.Hx_Field__beginLine == 0 {
			this.VTable.Hx_Field_newline()
		}
	}

	this.Hx_Field__prevLevel = this.Hx_Field__level
	this.Hx_Field__level++
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g1
		_g1 = rule.Hx_Field_body
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g1.Get(_g)
			_g++
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := node
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_2, nil)
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if this.Hx_Field__beginLine == 0 {
				var _hx_tmp_4 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_4
				_hx_tmp_4 = node
				var _hx_tmp_5 HxArray[*Hx_Obj_loreline_astnode]
				_ = _hx_tmp_5
				_hx_tmp_5 = rule.Hx_Field_body
				var _hx_tmp_6 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_6
				_hx_tmp_6 = _hx_tmp_4
				var _hx_tmp_7 HxArray[*Hx_Obj_loreline_astnode]
				_ = _hx_tmp_7
				_hx_tmp_7 = _hx_tmp_5
				_hx_tmp_3 = (_hx_tmp_6 != _hx_tmp_7.Get((rule.Hx_Field_body.Len() - 1)))
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				this.VTable.Hx_Field_newline()
			}

			this.Hx_Field__prevLevel = this.Hx_Field__level
		}
	}

	this.Hx_Field__level--
	this.VTable.Hx_Field_newline()
	if rule.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write("}")
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printImportStatement(imp *Hx_Obj_loreline_nimportstatement) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := imp
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write("import ")
	this.VTable.Hx_Field_printStringLiteral(imp.Hx_Field_path)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := imp
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printChoiceOption(option *Hx_Obj_loreline_nchoiceoption) {
	this.VTable.Hx_Field_newline()
	this.VTable.Hx_Field_newline()
	this.Hx_Field__noLn = 0
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := option
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if option.Hx_Field_once {
		this.VTable.Hx_Field_write("- ")
	}

	if option.Hx_Field_insertion != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := option.Hx_Field_insertion
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	} else {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := option.Hx_Field_text
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := option
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if option.Hx_Field_condition != nil {
		this.VTable.Hx_Field_write(" if ")
		this.VTable.Hx_Field_printInLineExpression(option.Hx_Field_condition, (option.Hx_Field_conditionStyle == ((int32)(2))))
	}

	if option.Hx_Field_body.Len() > 0 {
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if option.Hx_Field_condition == nil {
			_hx_tmp_3 = (option.Hx_Field_body.Len() == 1)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			_hx_tmp_2 = Hx_Field_std_isOfType(option.Hx_Field_body.Get(0), Hx_Obj_loreline_ntransition_RTTI)
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			this.Hx_Field__prevLevel = this.Hx_Field__level
			this.VTable.Hx_Field_write(" ")
			this.VTable.Hx_Field_printNode((func() *Hx_Obj_loreline_node {
				_hx_up := option.Hx_Field_body.Get(0)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})(), true)
			if this.Hx_Field__beginLine == 0 {
				this.VTable.Hx_Field_newline()
			}
		} else {
			if option.Hx_Field_style == ((int32)(1)) {
				this.VTable.Hx_Field_write(" {")
				this.VTable.Hx_Field_newline()
			} else {
				this.VTable.Hx_Field_newline()
			}

			this.Hx_Field__prevLevel = this.Hx_Field__level
			this.Hx_Field__level++
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_astnode]
				_ = _g1
				_g1 = option.Hx_Field_body
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					if !(_hx_tmp_4 < _g1.Len()) {
						break
					}

					var node *Hx_Obj_loreline_astnode
					_ = node
					node = _g1.Get(_g)
					_g++
					var _hx_tmp_5 *Hx_Obj_loreline_node
					_ = _hx_tmp_5
					_hx_tmp_5 = (func() *Hx_Obj_loreline_node {
						_hx_up := node
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					this.VTable.Hx_Field_printNode(_hx_tmp_5, nil)
					var _hx_tmp_6 bool
					_ = _hx_tmp_6
					if this.Hx_Field__beginLine == 0 {
						var _hx_tmp_7 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_7
						_hx_tmp_7 = node
						var _hx_tmp_8 HxArray[*Hx_Obj_loreline_astnode]
						_ = _hx_tmp_8
						_hx_tmp_8 = option.Hx_Field_body
						var _hx_tmp_9 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_9
						_hx_tmp_9 = _hx_tmp_7
						var _hx_tmp_10 HxArray[*Hx_Obj_loreline_astnode]
						_ = _hx_tmp_10
						_hx_tmp_10 = _hx_tmp_8
						_hx_tmp_6 = (_hx_tmp_9 != _hx_tmp_10.Get((option.Hx_Field_body.Len() - 1)))
					} else {
						_hx_tmp_6 = false
					}

					if _hx_tmp_6 {
						this.VTable.Hx_Field_newline()
					}

					this.Hx_Field__prevLevel = this.Hx_Field__level
				}
			}

			this.Hx_Field__level--
			this.VTable.Hx_Field_newline()
			if option.Hx_Field_style == ((int32)(1)) {
				this.VTable.Hx_Field_write("}")
				this.VTable.Hx_Field_newline()
			}
		}
	} else {
		this.VTable.Hx_Field_newline()
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printIfStatement(ifStmt *Hx_Obj_loreline_nifstatement, isElseIf__isOptional any) {
	isElseIf := HxDefault[bool]()
	_ = isElseIf

	if isElseIf__isOptional != nil {
		isElseIf = HxConvert[bool](isElseIf__isOptional)
	} else {
		isElseIf = false
	}
	{
		if this.Hx_Field__level == this.Hx_Field__prevLevel {
			this.VTable.Hx_Field_newline()
		}

		this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := ifStmt
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		this.VTable.Hx_Field_write("if ")
		this.VTable.Hx_Field_printInLineExpression(ifStmt.Hx_Field_condition, (ifStmt.Hx_Field_conditionStyle == ((int32)(2))))
		this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := ifStmt
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		if ifStmt.Hx_Field_thenBranch.Hx_Field_style == ((int32)(1)) {
			this.VTable.Hx_Field_write(" {")
		}

		this.VTable.Hx_Field_newline()
		this.Hx_Field__prevLevel = this.Hx_Field__level
		this.Hx_Field__level++
		var emptyLinesBeforeIf int32
		_ = emptyLinesBeforeIf
		emptyLinesBeforeIf = this.Hx_Field__numEmptyLines
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _g1
			_g1 = ifStmt.Hx_Field_thenBranch.Hx_Field_body
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var node *Hx_Obj_loreline_astnode
				_ = node
				node = _g1.Get(_g)
				_g++
				var _hx_tmp_1 *Hx_Obj_loreline_node
				_ = _hx_tmp_1
				_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
					_hx_up := node
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})()
				this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
				this.VTable.Hx_Field_newline()
				this.Hx_Field__prevLevel = this.Hx_Field__level
			}
		}

		this.Hx_Field__level--
		if ifStmt.Hx_Field_thenBranch.Hx_Field_style == ((int32)(1)) {
			this.VTable.Hx_Field_write("}")
		}

		if ifStmt.Hx_Field_elseBranch != nil {
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if ifStmt.Hx_Field_elseBranch.Hx_Field_body.Len() == 1 {
				_hx_tmp_2 = Hx_Field_std_isOfType(ifStmt.Hx_Field_elseBranch.Hx_Field_body.Get(0), Hx_Obj_loreline_nifstatement_RTTI)
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				var subIf *Hx_Obj_loreline_nifstatement
				_ = subIf
				subIf = ifStmt.Hx_Field_elseBranch.Hx_Field_body.Get(0).VTable.(interface {
					Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
				}).Hx_As_Hx_Obj_loreline_nifstatement()
				var _hx_tmp_3 bool
				_ = _hx_tmp_3
				if (this.Hx_Field__numEmptyLines - emptyLinesBeforeIf) > 0 {
					_hx_tmp_3 = true
				} else {
					_hx_tmp_3 = this.VTable.Hx_Field_needsEmptyLines((func() *Hx_Obj_loreline_astnode {
						_hx_up := subIf
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})())
				}

				if _hx_tmp_3 {
					this.VTable.Hx_Field_newline()
				}

				this.VTable.Hx_Field_write("else ")
				this.VTable.Hx_Field_printIfStatement(ifStmt.Hx_Field_elseBranch.Hx_Field_body.Get(0).VTable.(interface {
					Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
				}).Hx_As_Hx_Obj_loreline_nifstatement(), true)
			} else {
				if (this.Hx_Field__numEmptyLines - emptyLinesBeforeIf) > 0 {
					this.VTable.Hx_Field_newline()
				}

				if ifStmt.Hx_Field_elseBranch.Hx_Field_style == ((int32)(1)) {
					this.VTable.Hx_Field_write("else {")
					this.VTable.Hx_Field_newline()
				} else {
					this.VTable.Hx_Field_write("else")
					this.VTable.Hx_Field_newline()
				}

				this.Hx_Field__prevLevel = this.Hx_Field__level
				this.Hx_Field__level++
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_astnode]
					_ = _g1
					_g1 = ifStmt.Hx_Field_elseBranch.Hx_Field_body
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g
						if !(_hx_tmp_4 < _g1.Len()) {
							break
						}

						var node *Hx_Obj_loreline_astnode
						_ = node
						node = _g1.Get(_g)
						_g++
						var _hx_tmp_5 *Hx_Obj_loreline_node
						_ = _hx_tmp_5
						_hx_tmp_5 = (func() *Hx_Obj_loreline_node {
							_hx_up := node
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})()
						this.VTable.Hx_Field_printNode(_hx_tmp_5, nil)
						this.VTable.Hx_Field_newline()
						this.Hx_Field__prevLevel = this.Hx_Field__level
					}
				}

				this.Hx_Field__level--
				if ifStmt.Hx_Field_elseBranch.Hx_Field_style == ((int32)(1)) {
					this.VTable.Hx_Field_write("}")
					this.VTable.Hx_Field_newline()
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printAlternative(alt *Hx_Obj_loreline_nalternative) {
	if this.Hx_Field__level == this.Hx_Field__prevLevel {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := alt
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var keyword string
	_ = keyword
	switch alt.Hx_Field_mode {
	case 0:
		{
			keyword = "sequence"
		}

	case 1:
		{
			keyword = "cycle"
		}

	case 2:
		{
			keyword = "once"
		}

	case 3:
		{
			keyword = "pick"
		}

	case 4:
		{
			keyword = "shuffle"
		}

	}

	this.VTable.Hx_Field_write(keyword)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := alt
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if alt.Hx_Field_style == ((int32)(1)) {
		this.VTable.Hx_Field_write(" {")
	}

	this.VTable.Hx_Field_newline()
	this.Hx_Field__prevLevel = this.Hx_Field__level
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = alt.Hx_Field_items.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var item *Hx_Obj_loreline_nblock
			_ = item
			item = alt.Hx_Field_items.Get(i)
			this.Hx_Field__level++
			{
				var tmp__g1_1 int32
				_ = tmp__g1_1
				tmp__g1_1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_astnode]
				_ = _g2
				_g2 = item.Hx_Field_body
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = tmp__g1_1
					if !(_hx_tmp_1 < _g2.Len()) {
						break
					}

					var node *Hx_Obj_loreline_astnode
					_ = node
					node = _g2.Get(tmp__g1_1)
					tmp__g1_1++
					var _hx_tmp_2 *Hx_Obj_loreline_node
					_ = _hx_tmp_2
					_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
						_hx_up := node
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					this.VTable.Hx_Field_printNode(_hx_tmp_2, nil)
					this.VTable.Hx_Field_newline()
					this.Hx_Field__prevLevel = this.Hx_Field__level
				}
			}

			this.Hx_Field__level--
			if alt.Hx_Field_style == ((int32)(1)) {
				this.VTable.Hx_Field_write("}")
			}

			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = i
			if _hx_tmp_3 < (alt.Hx_Field_items.Len() - 1) {
				{
					this.VTable.Hx_Field_write("--")
					this.VTable.Hx_Field_newline()
				}

				if alt.Hx_Field_style == ((int32)(1)) {
					this.VTable.Hx_Field_write(" {")
				}

				this.VTable.Hx_Field_newline()
				this.Hx_Field__prevLevel = this.Hx_Field__level
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_needsEmptyLines(node *Hx_Obj_loreline_astnode) bool {
	var printer *Hx_Obj_loreline_printer
	_ = printer
	printer = Hx_Obj_loreline_printer_CreateInstance(this.Hx_Field__indent, "\n")
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = printer.VTable.Hx_Field_print((func() *Hx_Obj_loreline_node {
		_hx_up := node
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
	return ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "\n\n", HxNullable[int32]{})) != -1)
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printTransition(trans *Hx_Obj_loreline_ntransition, sameLine bool) {
	if !sameLine && (this.Hx_Field__prevLevel == this.Hx_Field__level) {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := trans
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if trans.Hx_Field_targetExpr != nil {
		this.VTable.Hx_Field_write("-> beat(")
		this.VTable.Hx_Field_printTargetExprAndArgs(trans.Hx_Field_targetExpr, trans.Hx_Field_args)
		this.VTable.Hx_Field_write(")")
	} else {
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("-> ", trans.Hx_Field_target))
		if trans.Hx_Field_args != nil {
			this.VTable.Hx_Field_write("(")
			this.VTable.Hx_Field_printTargetExprAndArgs(nil, trans.Hx_Field_args)
			this.VTable.Hx_Field_write(")")
		}
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := trans
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printInsertion(insert *Hx_Obj_loreline_ninsertion) {
	if this.Hx_Field__prevLevel == this.Hx_Field__level {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := insert
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if insert.Hx_Field_targetExpr != nil {
		this.VTable.Hx_Field_write("+ beat(")
		this.VTable.Hx_Field_printTargetExprAndArgs(insert.Hx_Field_targetExpr, insert.Hx_Field_args)
		this.VTable.Hx_Field_write(")")
	} else {
		this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat("+ ", insert.Hx_Field_target))
		if insert.Hx_Field_args != nil {
			this.VTable.Hx_Field_write("(")
			this.VTable.Hx_Field_printTargetExprAndArgs(nil, insert.Hx_Field_args)
			this.VTable.Hx_Field_write(")")
		}
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := insert
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printBeatCall(call *Hx_Obj_loreline_nbeatcall) {
	if this.Hx_Field__prevLevel == this.Hx_Field__level {
		this.VTable.Hx_Field_newline()
	}

	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := call
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write("beat(")
	this.VTable.Hx_Field_printTargetExprAndArgs(call.Hx_Field_targetExpr, call.Hx_Field_args)
	this.VTable.Hx_Field_write(")")
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := call
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printTargetExprAndArgs(targetExpr *Hx_Obj_loreline_nexpr, args HxArray[*Hx_Obj_loreline_nexpr]) {
	var savedComments bool
	_ = savedComments
	savedComments = this.Hx_Field_enableComments
	this.Hx_Field_enableComments = false
	var first bool
	_ = first
	first = true
	if targetExpr != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := targetExpr
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
		first = false
	}

	if args != nil {
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < args.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = args.Get(_g)
			_g++
			if !first {
				this.VTable.Hx_Field_write(", ")
			}

			first = false
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_2, nil)
		}
	}

	this.Hx_Field_enableComments = savedComments
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printStringLiteral(str *Hx_Obj_loreline_nstringliteral) {
	var surroundWithQuotes bool
	_ = surroundWithQuotes
	surroundWithQuotes = (str.Hx_Field_quotes == ((int32)(1)))
	if surroundWithQuotes {
		this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := str
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
		this.VTable.Hx_Field_write("\"")
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nstringpart]
		_ = _g1
		_g1 = str.Hx_Field_parts
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = _g1.Get(_g)
			_g++
			{
				var _g2 Hx_Enum_loreline_stringparttype
				_ = _g2
				_g2 = part.Hx_Field_partType
				switch _g2.Hx_Field_enumIndex() {
				case 0:
					{
						var text string
						_ = text
						text = _g2.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						if surroundWithQuotes {
							this.VTable.Hx_Field_writeQuotedRaw(text)
						} else {
							this.VTable.Hx_Field_writeUnquotedRaw(text)
						}
					}

				case 1:
					{
						var expr *Hx_Obj_loreline_nexpr
						_ = expr
						expr = _g2.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
						{
							var canBeSimple bool
							_ = canBeSimple
							canBeSimple = this.VTable.Hx_Field_isSimpleInterpolationExpr(expr)
							this.VTable.Hx_Field_write("$")
							if !canBeSimple {
								this.VTable.Hx_Field_write("{")
							}

							var _hx_tmp_1 *Hx_Obj_loreline_node
							_ = _hx_tmp_1
							_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
								_hx_up := expr
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
							this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
							if !canBeSimple {
								this.VTable.Hx_Field_write("}")
							}
						}
					}

				case 2:
					{
						var closing bool
						_ = closing
						closing = _g2.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_closing
						var content *Hx_Obj_loreline_nstringliteral
						_ = content
						content = _g2.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_expr
						{
							var _hx_tmp_2 string
							_ = _hx_tmp_2
							if closing {
								_hx_tmp_2 = "</"
							} else {
								_hx_tmp_2 = "<"
							}

							this.VTable.Hx_Field_write(_hx_tmp_2)
							this.VTable.Hx_Field_printStringLiteral(content)
							this.VTable.Hx_Field_write(">")
						}
					}

				}
			}
		}
	}

	if surroundWithQuotes {
		this.VTable.Hx_Field_write("\"")
		this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := str
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
	} else {
		this.VTable.Hx_Field_printHashTrailingComments((func() *Hx_Obj_loreline_astnode {
			_hx_up := str
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_writeUnquotedRaw(text string) {
	this.VTable.Hx_Field_writeRawWithContinuation(text, true)
}

func (this *Hx_Obj_loreline_printer) Hx_Field_writeQuotedRaw(text string) {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "\n", HxNullable[int32]{})) == -1 {
		this.VTable.Hx_Field_write(text)
		return
	}

	var lines HxArray[string]
	_ = lines
	lines = (Hx_Field_go_haxe_hxstring_split(text, "\n"))
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = lines.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_448 func(x any)
					_ = _hx_closure_src_448
					_hx_closure_src_448 = this.Hx_Field__buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_448(((any)(x)))
					}
				}()(this.Hx_Field__newline)
			}

			var line string
			_ = line
			line = lines.Get(i)
			var tmp bool
			_ = tmp
			if int32(utf8.RuneCountInString(line)) > 0 {
				var _hx_tmp_2 string
				_ = _hx_tmp_2
				_hx_tmp_2 = line
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_2, (int32(utf8.RuneCountInString(line)) - 1)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (c).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (c).Value
				}

				tmp = ((_hx_tmp_3) == 13)
			} else {
				tmp = false
			}

			if tmp {
				var endIndex HxNullable[int32]
				_ = endIndex
				endIndex = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(line)) - 1), Valid: true}

				var _hx_tmp_4 string
				_ = _hx_tmp_4
				if (endIndex).Valid != false {
					_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substring(line, 0, endIndex))
				} else {
					var _hx_tmp_5 string
					_ = _hx_tmp_5
					_hx_tmp_5 = line
					var _hx_tmp_6 string
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_5
					_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_6, 0, HxNullable[int32]{}))
				}

				line = _hx_tmp_4
			}

			if int32(utf8.RuneCountInString(line)) > 0 {
				func() func(x string) {
					var _hx_closure_src_449 func(x any)
					_ = _hx_closure_src_449
					_hx_closure_src_449 = this.Hx_Field__buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_449(((any)(x)))
					}
				}()(line)
				var _hx_tmp_7 string
				_ = _hx_tmp_7
				_hx_tmp_7 = line
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_7, (int32(utf8.RuneCountInString(line)) - 1)))
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				if (c).Valid == false {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = (c).Value
				}

				this.Hx_Field__lastChar = _hx_tmp_8
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_writeRawWithContinuation(text string, stripIndent bool) {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "\n", HxNullable[int32]{})) == -1 {
		this.VTable.Hx_Field_write(text)
		return
	}

	var lines HxArray[string]
	_ = lines
	lines = (Hx_Field_go_haxe_hxstring_split(text, "\n"))
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = lines.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_450 func(x any)
					_ = _hx_closure_src_450
					_hx_closure_src_450 = this.Hx_Field__buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_450(((any)(x)))
					}
				}()(this.Hx_Field__newline)
				this.Hx_Field__beginLine = 1
			}

			var line string
			_ = line
			line = lines.Get(i)
			var tmp bool
			_ = tmp
			if int32(utf8.RuneCountInString(line)) > 0 {
				var _hx_tmp_2 string
				_ = _hx_tmp_2
				_hx_tmp_2 = line
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_2, (int32(utf8.RuneCountInString(line)) - 1)))
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if (c).Valid == false {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = (c).Value
				}

				tmp = ((_hx_tmp_3) == 13)
			} else {
				tmp = false
			}

			if tmp {
				var endIndex HxNullable[int32]
				_ = endIndex
				endIndex = HxNullable[int32]{Value: (int32(utf8.RuneCountInString(line)) - 1), Valid: true}

				var _hx_tmp_4 string
				_ = _hx_tmp_4
				if (endIndex).Valid != false {
					_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substring(line, 0, endIndex))
				} else {
					var _hx_tmp_5 string
					_ = _hx_tmp_5
					_hx_tmp_5 = line
					var _hx_tmp_6 string
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_5
					_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_6, 0, HxNullable[int32]{}))
				}

				line = _hx_tmp_4
			}

			if stripIndent && (i > 0) {
				var j int32
				_ = j
				j = 0
				for {
					var tmp_tmp_1 bool
					_ = tmp_tmp_1
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = j
					if _hx_tmp_7 < int32(utf8.RuneCountInString(line)) {
						var c HxNullable[int32]
						_ = c
						c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, j))
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						if (c).Valid == false {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = (c).Value
						}

						if (_hx_tmp_8) != 32 {
							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(line, j))
							var _hx_tmp_9 int32
							_ = _hx_tmp_9
							if (tmp_c_1).Valid == false {
								_hx_tmp_9 = -1
							} else {
								_hx_tmp_9 = (tmp_c_1).Value
							}

							tmp_tmp_1 = ((_hx_tmp_9) == 9)
						} else {
							tmp_tmp_1 = true
						}
					} else {
						tmp_tmp_1 = false
					}

					if !tmp_tmp_1 {
						break
					}

					j++
				}

				if j > 0 {
					var endIndex HxNullable[int32]
					_ = endIndex
					endIndex = HxNullable[int32]{}

					var _hx_tmp_10 string
					_ = _hx_tmp_10
					if (endIndex).Valid != false {
						_hx_tmp_10 = (Hx_Field_go_haxe_hxstring_substring(line, j, endIndex))
					} else {
						var _hx_tmp_11 string
						_ = _hx_tmp_11
						_hx_tmp_11 = line
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						_hx_tmp_12 = j
						var _hx_tmp_13 string
						_ = _hx_tmp_13
						_hx_tmp_13 = _hx_tmp_11
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = _hx_tmp_12
						_hx_tmp_10 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_13, _hx_tmp_14, HxNullable[int32]{}))
					}

					line = _hx_tmp_10
				}
			}

			if int32(utf8.RuneCountInString(line)) > 0 {
				this.VTable.Hx_Field_write(line)
			}
		}
	}
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printLiteral(lit *Hx_Obj_loreline_nliteral) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := lit
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	{
		var _g Hx_Enum_loreline_literaltype
		_ = _g
		_g = lit.Hx_Field_literalType
		switch _g.Hx_Field_enumIndex() {
		case 0:
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(lit.Hx_Field_value, nil))) {
					_hx_tmp_0 = "null"
				} else {
					_hx_tmp_0 = Hx_Field_std_string(lit.Hx_Field_value)
				}

				this.VTable.Hx_Field_write(_hx_tmp_0)
			}

		case 1:
			{
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((lit.Hx_Field_value))) {
					_hx_tmp_1 = "true"
				} else {
					_hx_tmp_1 = "false"
				}

				this.VTable.Hx_Field_write(_hx_tmp_1)
			}

		case 2:
			{
				this.VTable.Hx_Field_write("null")
			}

		case 3:
			{
				var hasMultilineUnquoted bool
				_ = hasMultilineUnquoted
				hasMultilineUnquoted = false
				{
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var _g1 HxArray[any]
					_ = _g1
					_g1 = HxAnyToArray(lit.Hx_Field_value)
					for {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = tmp__g_1
						if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
							break
						}

						var elem any
						_ = elem
						elem = _g1.Get(tmp__g_1)
						tmp__g_1++
						var hx_dyncast_91 any
						_ = hx_dyncast_91
						hx_dyncast_91 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_node")
						var _hx_tmp_3 *Hx_Obj_loreline_node
						_ = _hx_tmp_3
						if hx_dyncast_91 != nil {
							_hx_tmp_3 = hx_dyncast_91.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
						} else {
							_hx_tmp_3 = nil
						}

						if this.VTable.Hx_Field_hasRealNewlines(_hx_tmp_3) {
							hasMultilineUnquoted = true
							break
						}
					}
				}

				if hasMultilineUnquoted {
					this.VTable.Hx_Field_write("[")
					this.VTable.Hx_Field_newline()
					this.Hx_Field__level++
					var first bool
					_ = first
					first = true
					{
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						var _g1 HxArray[any]
						_ = _g1
						_g1 = HxAnyToArray(lit.Hx_Field_value)
						for {
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = tmp__g_1
							if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
								break
							}

							var elem any
							_ = elem
							elem = _g1.Get(tmp__g_1)
							tmp__g_1++
							if !first {
								this.VTable.Hx_Field_write(",")
								this.VTable.Hx_Field_newline()
							}

							first = false
							var hx_dyncast_92 any
							_ = hx_dyncast_92
							hx_dyncast_92 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_node")
							var _hx_tmp_5 *Hx_Obj_loreline_node
							_ = _hx_tmp_5
							if hx_dyncast_92 != nil {
								_hx_tmp_5 = hx_dyncast_92.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
							} else {
								_hx_tmp_5 = nil
							}

							this.VTable.Hx_Field_printNode(_hx_tmp_5, nil)
						}
					}

					this.VTable.Hx_Field_newline()
					this.Hx_Field__level--
					this.VTable.Hx_Field_write("]")
				} else {
					var first bool
					_ = first
					first = true
					this.VTable.Hx_Field_write("[")
					{
						var tmp__g_1 int32
						_ = tmp__g_1
						tmp__g_1 = 0
						var _g1 HxArray[any]
						_ = _g1
						_g1 = HxAnyToArray(lit.Hx_Field_value)
						for {
							var _hx_tmp_6 int32
							_ = _hx_tmp_6
							_hx_tmp_6 = tmp__g_1
							if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
								break
							}

							var elem any
							_ = elem
							elem = _g1.Get(tmp__g_1)
							tmp__g_1++
							if !first {
								this.VTable.Hx_Field_write(", ")
							}

							first = false
							var hx_dyncast_93 any
							_ = hx_dyncast_93
							hx_dyncast_93 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_node")
							var _hx_tmp_7 *Hx_Obj_loreline_node
							_ = _hx_tmp_7
							if hx_dyncast_93 != nil {
								_hx_tmp_7 = hx_dyncast_93.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
							} else {
								_hx_tmp_7 = nil
							}

							this.VTable.Hx_Field_printNode(_hx_tmp_7, nil)
						}
					}

					this.VTable.Hx_Field_write("]")
				}
			}

		case 4:
			{
				var style int32
				_ = style
				style = _g.(Hx_Enum_loreline_literaltype_Object).Hx_Field_style
				{
					if style == ((int32)(1)) {
						this.VTable.Hx_Field_write("{")
						this.VTable.Hx_Field_newline()
					} else {
						this.VTable.Hx_Field_newline()
					}

					this.Hx_Field__level++
					var first bool
					_ = first
					first = true
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[*Hx_Obj_loreline_nobjectfield]
						_ = _g2
						_g2 = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](lit.Hx_Field_value)
						for {
							var _hx_tmp_8 int32
							_ = _hx_tmp_8
							_hx_tmp_8 = _g1
							if !(_hx_tmp_8 < _g2.Len()) {
								break
							}

							var field *Hx_Obj_loreline_nobjectfield
							_ = field
							field = _g2.Get(_g1)
							_g1++
							if !first {
								if this.Hx_Field__beginLine == 0 {
									this.VTable.Hx_Field_newline()
								}
							}

							first = false
							this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
								_hx_up := field
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_astnode
								}
								return nil
							})())
							this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", field.Hx_Field_name), ":"))
							this.VTable.Hx_Field_space()
							this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
								_hx_up := field
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_astnode
								}
								return nil
							})())
							var _hx_tmp_9 *Hx_Obj_loreline_node
							_ = _hx_tmp_9
							_hx_tmp_9 = (func() *Hx_Obj_loreline_node {
								_hx_up := field.Hx_Field_value
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
							this.VTable.Hx_Field_printNode(_hx_tmp_9, nil)
						}
					}

					this.VTable.Hx_Field_newline()
					this.Hx_Field__level--
					if style == ((int32)(1)) {
						this.VTable.Hx_Field_write("}")
					}
				}
			}

		}
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := lit
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printAccess(access *Hx_Obj_loreline_naccess) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	if access.Hx_Field_target != nil {
		this.VTable.Hx_Field_printOperand(access.Hx_Field_target, 10)
		this.VTable.Hx_Field_write(".")
	}

	this.VTable.Hx_Field_write(access.Hx_Field_name)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printArrayAccess(access *Hx_Obj_loreline_narrayaccess) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_printOperand(access.Hx_Field_target, 10)
	this.VTable.Hx_Field_write("[")
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := access.Hx_Field_index
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	this.VTable.Hx_Field_write("]")
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := access
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printCall(call *Hx_Obj_loreline_ncall) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := call
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_printOperand(call.Hx_Field_target, 10)
	this.VTable.Hx_Field_write("(")
	var first bool
	_ = first
	first = true
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g1
		_g1 = call.Hx_Field_args
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = _g1.Get(_g)
			_g++
			if !first {
				this.VTable.Hx_Field_write(", ")
			}

			first = false
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
		}
	}

	this.VTable.Hx_Field_write(")")
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := call
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printBinary(binary *Hx_Obj_loreline_nbinary) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := binary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var level int32
	_ = level
	level = Hx_Field_loreline_printer_precedence((func() *Hx_Obj_loreline_nexpr {
		_hx_up := binary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})())
	this.VTable.Hx_Field_printOperand(binary.Hx_Field_left, level)
	this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(" ", this.VTable.Hx_Field_getOperator(binary.Hx_Field_op)), " "))
	if Hx_Field_loreline_printer_isWordNotAfterIs(binary) {
		this.VTable.Hx_Field_write("(")
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := binary.Hx_Field_right
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
		this.VTable.Hx_Field_write(")")
	} else {
		this.VTable.Hx_Field_printOperand(binary.Hx_Field_right, (level + 1))
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := binary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printUnary(unary *Hx_Obj_loreline_nunary) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := unary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write(this.VTable.Hx_Field_getOperator(unary.Hx_Field_op))
	{
		var _g Hx_Enum_loreline_tokentype
		_ = _g
		_g = unary.Hx_Field_op
		if _g.Hx_Field_enumIndex() == 40 {
			if _g.(Hx_Enum_loreline_tokentype_OpNot).Hx_Field_word == true {
				this.VTable.Hx_Field_write(" ")
			}
		}
	}

	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := unary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_printOperand(unary.Hx_Field_operand, 9)
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printTernary(ternary *Hx_Obj_loreline_nternary) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := ternary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_printOperand(ternary.Hx_Field_condition, 3)
	this.VTable.Hx_Field_write(" ? ")
	this.VTable.Hx_Field_printOperand(ternary.Hx_Field_trueExpr, 2)
	this.VTable.Hx_Field_write(" : ")
	this.VTable.Hx_Field_printOperand(ternary.Hx_Field_falseExpr, 2)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := ternary
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printAssignment(assign *Hx_Obj_loreline_nassign) {
	this.VTable.Hx_Field_printLeadingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := assign
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := assign.Hx_Field_target
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	this.VTable.Hx_Field_printTrailingComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := assign
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	this.VTable.Hx_Field_write(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(" ", this.VTable.Hx_Field_getOperator(assign.Hx_Field_op)), " "))
	var _hx_tmp_1 *Hx_Obj_loreline_node
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
		_hx_up := assign.Hx_Field_value
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_1, nil)
}

func (this *Hx_Obj_loreline_printer) Hx_Field_printInLineExpression(expr *Hx_Obj_loreline_nexpr, parens bool) {
	var savedComments bool
	_ = savedComments
	savedComments = this.Hx_Field_enableComments
	this.Hx_Field_enableComments = false
	if parens {
		this.VTable.Hx_Field_write("(")
	}

	var _hx_tmp_0 *Hx_Obj_loreline_node
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
		_hx_up := expr
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()
	this.VTable.Hx_Field_printNode(_hx_tmp_0, nil)
	if parens {
		this.VTable.Hx_Field_write(")")
	}

	this.Hx_Field_enableComments = savedComments
}

func (this *Hx_Obj_loreline_printer) Hx_Field_isSimpleInterpolationExpr(expr *Hx_Obj_loreline_nexpr) bool {
	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_naccess_RTTI) {
		var access *Hx_Obj_loreline_naccess
		_ = access
		access = expr.VTable.(interface {
			Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
		}).Hx_As_Hx_Obj_loreline_naccess()
		if access.Hx_Field_target == nil {
			return true
		}

		return this.VTable.Hx_Field_isSimpleInterpolationExpr(access.Hx_Field_target)
	}

	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_narrayaccess_RTTI) {
		var hx_cast_ex_94 *Hx_Obj_loreline_nexpr
		_ = hx_cast_ex_94
		hx_cast_ex_94 = expr
		if !Hx_Field_std_isOfType(hx_cast_ex_94, Hx_Obj_loreline_narrayaccess_RTTI) {
			panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
		}

		return this.VTable.Hx_Field_isSimpleInterpolationExpr((hx_cast_ex_94.VTable.(interface {
			Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
		}).Hx_As_Hx_Obj_loreline_narrayaccess()).Hx_Field_target)
	}

	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_ncall_RTTI) {
		var hx_cast_ex_95 *Hx_Obj_loreline_nexpr
		_ = hx_cast_ex_95
		hx_cast_ex_95 = expr
		if !Hx_Field_std_isOfType(hx_cast_ex_95, Hx_Obj_loreline_ncall_RTTI) {
			panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
		}

		return this.VTable.Hx_Field_isSimpleInterpolationExpr((hx_cast_ex_95.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()).Hx_Field_target)
	}

	return false
}

func (this *Hx_Obj_loreline_printer) Hx_Field_hasRealNewlines(node *Hx_Obj_loreline_node) bool {
	if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nstringliteral_RTTI) {
		var str *Hx_Obj_loreline_nstringliteral
		_ = str
		str = node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
		}).Hx_As_Hx_Obj_loreline_nstringliteral()
		if str.Hx_Field_quotes == ((int32)(1)) {
			return false
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_nstringpart]
			_ = _g1
			_g1 = str.Hx_Field_parts
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var part *Hx_Obj_loreline_nstringpart
				_ = part
				part = _g1.Get(_g)
				_g++
				{
					var _g2 Hx_Enum_loreline_stringparttype
					_ = _g2
					_g2 = part.Hx_Field_partType
					if _g2.Hx_Field_enumIndex() == 0 {
						var text string
						_ = text
						text = _g2.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						var _hx_tmp_1 string
						_ = _hx_tmp_1
						_hx_tmp_1 = text
						if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, "\n", HxNullable[int32]{})) != -1 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_printer) Hx_Field_getOperator(op Hx_Enum_loreline_tokentype) string {
	switch op.Hx_Field_enumIndex() {
	case 21:
		{
			return "="
		}

	case 22:
		{
			return "+="
		}

	case 23:
		{
			return "-="
		}

	case 24:
		{
			return "*="
		}

	case 25:
		{
			return "/="
		}

	case 26:
		{
			return ":="
		}

	case 27:
		{
			return "+"
		}

	case 28:
		{
			return "-"
		}

	case 29:
		{
			return "*"
		}

	case 30:
		{
			return "/"
		}

	case 31:
		{
			return "%"
		}

	case 32:
		{
			var word bool
			_ = word
			word = op.(Hx_Enum_loreline_tokentype_OpEquals).Hx_Field_word
			if word {
				return "is"
			} else {
				return "=="
			}
		}

	case 33:
		{
			var word bool
			_ = word
			word = op.(Hx_Enum_loreline_tokentype_OpNotEquals).Hx_Field_word
			if word {
				return "is not"
			} else {
				return "!="
			}
		}

	case 34:
		{
			return ">"
		}

	case 35:
		{
			return "<"
		}

	case 36:
		{
			return ">="
		}

	case 37:
		{
			return "<="
		}

	case 38:
		{
			if op.(Hx_Enum_loreline_tokentype_OpAnd).Hx_Field_word {
				return "and"
			} else {
				return "&&"
			}
		}

	case 39:
		{
			if op.(Hx_Enum_loreline_tokentype_OpOr).Hx_Field_word {
				return "or"
			} else {
				return "||"
			}
		}

	case 40:
		{
			var word bool
			_ = word
			word = op.(Hx_Enum_loreline_tokentype_OpNot).Hx_Field_word
			if word {
				return "not"
			} else {
				return "!"
			}
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unsupported operator: ", Hx_Field_std_string(op))))))
		}
	}
}

func Hx_Field_loreline_printer_precedence(expr *Hx_Obj_loreline_nexpr) int32 {
	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nassign_RTTI) {
		return 1
	}

	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nternary_RTTI) {
		return 2
	}

	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nunary_RTTI) {
		return 9
	}

	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nbinary_RTTI) {
		switch expr.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
		}).Hx_As_Hx_Obj_loreline_nbinary().Hx_Field_op.Hx_Field_enumIndex() {
		case 27, 28:
			{
				return 7
			}

		case 29, 30, 31:
			{
				return 8
			}

		case 32:
			{
				return 5
			}

		case 33:
			{
				return 5
			}

		case 34, 35, 36, 37:
			{
				return 6
			}

		case 38:
			{
				return 4
			}

		case 39:
			{
				return 3
			}

		default:
			{
				return 8
			}
		}
	}

	return 10
}

func Hx_Field_loreline_printer_isWordNotAfterIs(binary *Hx_Obj_loreline_nbinary) bool {
	{
		var _g Hx_Enum_loreline_tokentype
		_ = _g
		_g = binary.Hx_Field_op
		if _g.Hx_Field_enumIndex() == 32 {
			if _g.(Hx_Enum_loreline_tokentype_OpEquals).Hx_Field_word != true {
				return false
			}
		} else {
			return false
		}
	}

	var right *Hx_Obj_loreline_nexpr
	_ = right
	right = binary.Hx_Field_right
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (right == nil) || (right.Hx_Field_parens > 0) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = !Hx_Field_std_isOfType(right, Hx_Obj_loreline_nunary_RTTI)
	}

	if _hx_tmp_0 {
		return false
	}

	var _g1 Hx_Enum_loreline_tokentype
	_ = _g1
	_g1 = right.VTable.(interface {
		Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
	}).Hx_As_Hx_Obj_loreline_nunary().Hx_Field_op
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if _g1.Hx_Field_enumIndex() == 40 {
		_hx_tmp_1 = (_g1.(Hx_Enum_loreline_tokentype_OpNot).Hx_Field_word == true)
	} else {
		_hx_tmp_1 = false
	}

	return _hx_tmp_1
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nchoiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
