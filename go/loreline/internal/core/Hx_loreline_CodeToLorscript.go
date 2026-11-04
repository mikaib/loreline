package core

import "unicode/utf8"

var Hx_Abstract_loreline_codetolorscriptstacktype_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.CodeToLorscriptStackType",
	HxMakeArray[string]("ObjectBrace", "ArrayBracket", "Brace", "Indent", "CaseIndent", "Bracket", "Paren", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__codetolorscript_codetolorscriptstacktype_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__codetolorscript_codetolorscriptstacktype_impl_() {
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ObjectBrace = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ObjectBrace()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ArrayBracket = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ArrayBracket()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Brace = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Brace()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Indent = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Indent()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__CaseIndent = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__CaseIndent()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Bracket = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Bracket()
	Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Paren = Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Paren()
}

type Hx_Obj_VTable_loreline__codetolorscript_codetolorscriptstacktype_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_ struct {
	VTable Hx_Obj_VTable_loreline__codetolorscript_codetolorscriptstacktype_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__CreateEmptyInstance() *Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_ {
	obj := &Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__CreateInstance() *Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_ {
	obj := Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._CodeToLorscript.CodeToLorscriptStackType_Impl_",
	HxMakeArray[string]("ObjectBrace", "ArrayBracket", "Brace", "Indent", "CaseIndent", "Bracket", "Paren", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "ObjectBrace":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ObjectBrace
		case "ArrayBracket":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ArrayBracket
		case "Brace":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Brace
		case "Indent":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Indent
		case "CaseIndent":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__CaseIndent
		case "Bracket":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Bracket
		case "Paren":
			return &Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Paren
		case "toString":
			return Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl__RTTI
}

func (this *Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_) Hx_As_Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_() *Hx_Obj_loreline__codetolorscript_codetolorscriptstacktype_impl_ {
	return this
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ObjectBrace int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ObjectBrace() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ArrayBracket int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__ArrayBracket() int32 {
	return ((int32)(1))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Brace int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Brace() int32 {
	return ((int32)(2))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Indent int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Indent() int32 {
	return ((int32)(3))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__CaseIndent int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__CaseIndent() int32 {
	return ((int32)(4))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Bracket int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Bracket() int32 {
	return ((int32)(5))
}

var Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Paren int32

func Hx_Init_Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__Paren() int32 {
	return ((int32)(6))
}

func Hx_Field_loreline__codetolorscript_codetolorscriptstacktype_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "ObjectBrace"
		}

	case 1:
		{
			return "ArrayBracket"
		}

	case 2:
		{
			return "Brace"
		}

	case 3:
		{
			return "Indent"
		}

	case 4:
		{
			return "CaseIndent"
		}

	case 5:
		{
			return "Bracket"
		}

	case 6:
		{
			return "Paren"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_codetolorscript() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_codetolorscript() {
	Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS = Hx_Init_Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS()
}

type Hx_Obj_VTable_loreline_codetolorscript interface {
	Hx_Field_processedPosFromInputPos(pos int32) int32
	Hx_Field_processFieldAccessInterpolation()
	Hx_Field_endsOrFollowsWithChar(line string, c int32, pos int32, hxpos__isOptional any) bool
	Hx_Field_nextLineIndentOffset(line string, pos int32) int32
	Hx_Field_stackPush(item int32)
	Hx_Field_inputPosFromProcessedPos(pos int32) int32
	Hx_Field_followsWithChar(c int32, pos int32) bool
	Hx_Field_isWhitespace(c int32) bool
	Hx_Field_stackPop() HxNullable[int32]
	Hx_Field_endsWithWord(line string, word string) bool
	Hx_Field_endsWithChar(line string, c int32) bool
	Hx_Field_toLorelinePos(funcPos *Hx_Obj_loreline_position, pmin int32, pmax int32) *Hx_Obj_loreline_position
	Hx_Field_processInput(until__isOptional any)
	Hx_Field_isDigit(c int32) bool
	Hx_Field_followsWithIf(pos int32) bool
	Hx_Field_processString()
	Hx_Field_isIdentifierPart(c int32) bool
	Hx_Field_inStatementsBlock() bool
	Hx_Field_inArrayBlock() bool
	Hx_Field_isIdentifierStart(c int32) bool
	Hx_Field_processComment()
	Hx_Field_processComplexInterpolation()
	Hx_Field_isAlphaNumeric(c int32) bool
	Hx_Field_isLabelStart(pos int32) bool
	Hx_Field_isAfterDot() bool
	Hx_Field_endsWithControlKeyword(line string, pos int32) bool
	Hx_Field_endsWithArrayIndexable(line string) bool
	Hx_Field_process(input string) string
	Hx_Field_processIdentifier()
	Hx_Field_followsWithWord(word string, pos int32) bool
	Hx_Field__add(c int32, incrementIndex bool)
	Hx_Field_error(message string)
	Hx_Field_inObjectBlock() bool
	Hx_Field_skipWhitespaceAndComments(pos int32) int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_codetolorscript struct {
	VTable                          Hx_Obj_VTable_loreline_codetolorscript `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_index                  int32
	Hx_Field_input                  HxArray[int32]
	Hx_Field_inputText              string
	Hx_Field_output                 Hx_Typedef_loreline_utf8buf
	Hx_Field_lineOutput             Hx_Typedef_loreline_utf8buf
	Hx_Field_length                 int32
	Hx_Field_indent                 int32
	Hx_Field_inComment              bool
	Hx_Field_inString               bool
	Hx_Field_posOffsets             HxArray[int32]
	Hx_Field_currentPosOffset       int32
	Hx_Field_inControl              bool
	Hx_Field_inControlWithoutParens bool
	Hx_Field_indentLevel            int32
	Hx_Field_indentStack            HxArray[int32]
	Hx_Field_stack                  HxArray[int32]
	Hx_Field_doBraces               HxArray[int32]
	Hx_Field_closedDoBrace          bool
}

func Hx_Obj_loreline_codetolorscript_CreateEmptyInstance() *Hx_Obj_loreline_codetolorscript {
	obj := &Hx_Obj_loreline_codetolorscript{}
	obj.VTable = obj
	obj.Hx_Field_index = HxDefault[int32]()
	obj.Hx_Field_input = HxDefault[HxArray[int32]]()
	obj.Hx_Field_inputText = HxDefault[string]()
	obj.Hx_Field_output = HxDefault[Hx_Typedef_loreline_utf8buf]()
	obj.Hx_Field_lineOutput = HxDefault[Hx_Typedef_loreline_utf8buf]()
	obj.Hx_Field_length = HxDefault[int32]()
	obj.Hx_Field_indent = HxDefault[int32]()
	obj.Hx_Field_inComment = HxDefault[bool]()
	obj.Hx_Field_inString = HxDefault[bool]()
	obj.Hx_Field_posOffsets = HxDefault[HxArray[int32]]()
	obj.Hx_Field_currentPosOffset = HxDefault[int32]()
	obj.Hx_Field_inControl = HxDefault[bool]()
	obj.Hx_Field_inControlWithoutParens = HxDefault[bool]()
	obj.Hx_Field_indentLevel = HxDefault[int32]()
	obj.Hx_Field_indentStack = HxDefault[HxArray[int32]]()
	obj.Hx_Field_stack = HxDefault[HxArray[int32]]()
	obj.Hx_Field_doBraces = HxDefault[HxArray[int32]]()
	obj.Hx_Field_closedDoBrace = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_codetolorscript_CreateInstance() *Hx_Obj_loreline_codetolorscript {
	obj := Hx_Obj_loreline_codetolorscript_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_New() {
	{
		this.Hx_Field_index = 0
	}
	{
		this.Hx_Field_input = nil
	}
	{
		this.Hx_Field_inputText = HxStringNull
	}
	{
		this.Hx_Field_output = nil
	}
	{
		this.Hx_Field_lineOutput = nil
	}
	{
		this.Hx_Field_length = 0
	}
	{
		this.Hx_Field_indent = 0
	}
	{
		this.Hx_Field_inComment = false
	}
	{
		this.Hx_Field_inString = false
	}
	{
		this.Hx_Field_posOffsets = nil
	}
	{
		this.Hx_Field_currentPosOffset = 0
	}
	{
		this.Hx_Field_inControl = false
	}
	{
		this.Hx_Field_inControlWithoutParens = false
	}
	{
		this.Hx_Field_indentLevel = 0
	}
	{
		this.Hx_Field_indentStack = nil
	}
	{
		this.Hx_Field_closedDoBrace = false
	}
	{
	}

}

var Hx_Obj_loreline_codetolorscript_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.CodeToLorscript",
	HxMakeArray[string]("CONTROL_KEYWORDS"),
	HxMakeArray[string]("index", "input", "inputText", "output", "lineOutput", "length", "indent", "inComment", "inString", "posOffsets", "currentPosOffset", "inControl", "inControlWithoutParens", "indentLevel", "indentStack", "stack", "doBraces", "closedDoBrace", "process", "toLorelinePos", "inputPosFromProcessedPos", "processedPosFromInputPos", "processInput", "processString", "processComplexInterpolation", "processFieldAccessInterpolation", "processIdentifier", "isDigit", "isIdentifierStart", "isIdentifierPart", "processComment", "endsWithChar", "followsWithChar", "endsOrFollowsWithChar", "nextLineIndentOffset", "followsWithIf", "followsWithWord", "endsWithControlKeyword", "endsWithWord", "endsWithArrayIndexable", "isWhitespace", "isAfterDot", "isAlphaNumeric", "addExtra", "add", "inStatementsBlock", "inObjectBlock", "inArrayBlock", "isLabelStart", "skipWhitespaceAndComments", "stackPush", "stackPop", "_add", "error"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_codetolorscript_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_codetolorscript_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_codetolorscript_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "CONTROL_KEYWORDS":
			return &Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_codetolorscript_RTTI
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_As_Hx_Obj_loreline_codetolorscript() *Hx_Obj_loreline_codetolorscript {
	return this
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_process(input string) string {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(input))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(input, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = input
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = input
	}

	input = Hx_Field_go_haxe_hxdynamic_stringConcat((_hx_tmp_1), "\n//<END>")
	this.Hx_Field_inputText = input
	this.Hx_Field_input = Hx_Field_loreline_utf8_toCodes(input)
	this.Hx_Field_index = 0
	this.Hx_Field_output = Hx_Obj_stringbuf_CreateInstance()
	this.Hx_Field_lineOutput = Hx_Obj_stringbuf_CreateInstance()
	this.Hx_Field_length = int32(utf8.RuneCountInString(input))
	this.Hx_Field_indent = 0
	this.Hx_Field_inComment = false
	this.Hx_Field_inString = false
	this.Hx_Field_posOffsets = HxMakeArray[int32]()
	this.Hx_Field_currentPosOffset = 0
	this.Hx_Field_inControl = false
	this.Hx_Field_inControlWithoutParens = false
	this.Hx_Field_indentStack = HxMakeArray[int32]()
	this.Hx_Field_indentLevel = 0
	this.Hx_Field_stack = HxMakeArray[int32]()
	this.Hx_Field_doBraces = HxMakeArray[int32]()
	this.Hx_Field_closedDoBrace = false
	this.VTable.Hx_Field_processInput(nil)
	var s string
	_ = s
	s = this.Hx_Field_output.VTable.Hx_Field_toString()
	var l1 int32
	_ = l1
	l1 = int32(utf8.RuneCountInString(s))
	var r1 int32
	_ = r1
	r1 = 0
	for {
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if r1 < l1 {
			_hx_tmp_4 = Hx_Field_stringtools_isSpace(s, ((l1 - r1) - 1))
		} else {
			_hx_tmp_4 = false
		}

		if !(_hx_tmp_4) {
			break
		}

		r1++
	}

	var _hx_tmp_5 string
	_ = _hx_tmp_5
	if r1 > 0 {
		var _hx_tmp_6 string
		_ = _hx_tmp_6
		_hx_tmp_6 = s
		var _hx_tmp_7 string
		_ = _hx_tmp_7
		_hx_tmp_7 = _hx_tmp_6
		_hx_tmp_5 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, 0, HxNullable[int32]{Value: (l1 - r1), Valid: true}))
	} else {
		_hx_tmp_5 = s
	}

	return Hx_Field_go_haxe_hxdynamic_stringConcat((_hx_tmp_5), "\n")
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_toLorelinePos(funcPos *Hx_Obj_loreline_position, pmin int32, pmax int32) *Hx_Obj_loreline_position {
	var min int32
	_ = min
	min = this.VTable.Hx_Field_inputPosFromProcessedPos(pmin)
	var max int32
	_ = max
	max = this.VTable.Hx_Field_inputPosFromProcessedPos(pmax)
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = ((max + 1) - min)
	return funcPos.VTable.Hx_Field_withOffset(this.Hx_Field_inputText, min, _hx_reserved_len, funcPos.Hx_Field_offset)
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_inputPosFromProcessedPos(pos int32) int32 {
	if pos < 0 {
		return 0
	}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = pos
	if _hx_tmp_0 >= this.Hx_Field_posOffsets.Len() {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = (HxMakeArrayView[any](this.Hx_Field_input).Len() - 1)
		var _hx_tmp_2 HxArray[int32]
		_ = _hx_tmp_2
		_hx_tmp_2 = this.Hx_Field_posOffsets
		return (_hx_tmp_1 - _hx_tmp_2.Get((HxMakeArrayView[any](this.Hx_Field_input).Len() - 1)))
	}

	return (pos - this.Hx_Field_posOffsets.Get(pos))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processedPosFromInputPos(pos int32) int32 {
	if pos < 0 {
		return 0
	}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = pos
	if _hx_tmp_0 >= HxMakeArrayView[any](this.Hx_Field_input).Len() {
		return (this.Hx_Field_output.VTable.Hx_Field_get_length() - 1)
	}

	var count int32
	_ = count
	count = 0
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_posOffsets.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			if (i - this.Hx_Field_posOffsets.Get(i)) > pos {
				return (i - 1)
			} else {
				if (i - this.Hx_Field_posOffsets.Get(i)) == pos {
					return i
				}
			}

			count = i
		}
	}

	return count
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processInput(until__isOptional any) {
	until := HxDefault[int32]()
	_ = until

	if until__isOptional != nil {
		until = HxConvert[int32](until__isOptional)
	} else {
		until = -1
	}
	{
		var braceLevel int32
		_ = braceLevel
		braceLevel = 0
		var bracketLevel int32
		_ = bracketLevel
		bracketLevel = 0
		var parenLevel int32
		_ = parenLevel
		parenLevel = 0
		for this.Hx_Field_index < this.Hx_Field_length {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_index
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if pos < 0 {
				_hx_tmp_1 = true
			} else {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = pos
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _hx_tmp_2
				_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_1 {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			var c int32
			_ = c
			c = _hx_tmp_0
			if c == 34 {
				this.VTable.Hx_Field_processString()
			} else {
				if c == 39 {
					this.VTable.Hx_Field_error("Unexpected single quote")
				} else {
					if c == 47 {
						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = (this.Hx_Field_index + 1)
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						var _hx_tmp_5 bool
						_ = _hx_tmp_5
						if tmp_pos_1 < 0 {
							_hx_tmp_5 = true
						} else {
							var _hx_tmp_6 int32
							_ = _hx_tmp_6
							_hx_tmp_6 = tmp_pos_1
							var _hx_tmp_7 int32
							_ = _hx_tmp_7
							_hx_tmp_7 = _hx_tmp_6
							_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_5 {
							_hx_tmp_4 = -1
						} else {
							_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
						}

						var cc int32
						_ = cc
						cc = _hx_tmp_4
						if (cc == 47) || (cc == 42) {
							this.VTable.Hx_Field_processComment()
						} else {
							this.VTable.Hx_Field__add(c, true)
						}
					} else {
						if c == 13 {
							this.VTable.Hx_Field__add(c, true)
						} else {
							if c == 10 {
								this.VTable.Hx_Field__add(c, true)
							} else {
								if c == 123 {
									braceLevel++
									this.VTable.Hx_Field__add(c, true)
								} else {
									if c == 125 {
										braceLevel--
										if (c == until) && (braceLevel < 0) {
											return
										} else {
											this.VTable.Hx_Field__add(c, true)
										}
									} else {
										if c == 91 {
											bracketLevel++
											this.VTable.Hx_Field__add(c, true)
										} else {
											if c == 93 {
												bracketLevel--
												if (c == until) && (bracketLevel < 0) {
													return
												} else {
													this.VTable.Hx_Field__add(c, true)
												}
											} else {
												if c == 40 {
													parenLevel++
													this.VTable.Hx_Field__add(c, true)
												} else {
													if c == 41 {
														parenLevel--
														if (c == until) && (parenLevel < 0) {
															return
														} else {
															this.VTable.Hx_Field__add(c, true)
														}
													} else {
														var tmp bool
														_ = tmp
														var _hx_tmp_8 bool
														_ = _hx_tmp_8
														if this.VTable.Hx_Field_isAlphaNumeric(c) {
															_hx_tmp_8 = (this.Hx_Field_index > 0)
														} else {
															_hx_tmp_8 = false
														}

														if _hx_tmp_8 {
															var tmp_v_1 HxArray[int32]
															_ = tmp_v_1
															tmp_v_1 = this.Hx_Field_input
															var tmp_pos_1 int32
															_ = tmp_pos_1
															tmp_pos_1 = (this.Hx_Field_index - 1)
															var _hx_tmp_9 int32
															_ = _hx_tmp_9
															var _hx_tmp_10 bool
															_ = _hx_tmp_10
															if tmp_pos_1 < 0 {
																_hx_tmp_10 = true
															} else {
																var _hx_tmp_11 int32
																_ = _hx_tmp_11
																_hx_tmp_11 = tmp_pos_1
																var _hx_tmp_12 int32
																_ = _hx_tmp_12
																_hx_tmp_12 = _hx_tmp_11
																_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](tmp_v_1).Len())
															}

															if _hx_tmp_10 {
																_hx_tmp_9 = -1
															} else {
																_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
															}

															tmp = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_9)
														} else {
															tmp = false
														}

														if tmp {
															if c == 97 {
																var tmp_tmp_1 bool
																_ = tmp_tmp_1
																var tmp1 bool
																_ = tmp1
																var tmp_v_1 HxArray[int32]
																_ = tmp_v_1
																tmp_v_1 = this.Hx_Field_input
																var tmp_pos_1 int32
																_ = tmp_pos_1
																tmp_pos_1 = (this.Hx_Field_index + 1)
																var _hx_tmp_13 int32
																_ = _hx_tmp_13
																var _hx_tmp_14 bool
																_ = _hx_tmp_14
																if tmp_pos_1 < 0 {
																	_hx_tmp_14 = true
																} else {
																	var _hx_tmp_15 int32
																	_ = _hx_tmp_15
																	_hx_tmp_15 = tmp_pos_1
																	var _hx_tmp_16 int32
																	_ = _hx_tmp_16
																	_hx_tmp_16 = _hx_tmp_15
																	_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](tmp_v_1).Len())
																}

																if _hx_tmp_14 {
																	_hx_tmp_13 = -1
																} else {
																	_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																}

																if (_hx_tmp_13) == 110 {
																	var tmp_v_2 HxArray[int32]
																	_ = tmp_v_2
																	tmp_v_2 = this.Hx_Field_input
																	var tmp_pos_2 int32
																	_ = tmp_pos_2
																	tmp_pos_2 = (this.Hx_Field_index + 2)
																	var _hx_tmp_17 int32
																	_ = _hx_tmp_17
																	var _hx_tmp_18 bool
																	_ = _hx_tmp_18
																	if tmp_pos_2 < 0 {
																		_hx_tmp_18 = true
																	} else {
																		var _hx_tmp_19 int32
																		_ = _hx_tmp_19
																		_hx_tmp_19 = tmp_pos_2
																		var _hx_tmp_20 int32
																		_ = _hx_tmp_20
																		_hx_tmp_20 = _hx_tmp_19
																		_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_2).Len())
																	}

																	if _hx_tmp_18 {
																		_hx_tmp_17 = -1
																	} else {
																		_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																	}

																	tmp1 = ((_hx_tmp_17) == 100)
																} else {
																	tmp1 = false
																}

																if tmp1 {
																	var tmp_v_2 HxArray[int32]
																	_ = tmp_v_2
																	tmp_v_2 = this.Hx_Field_input
																	var tmp_pos_2 int32
																	_ = tmp_pos_2
																	tmp_pos_2 = (this.Hx_Field_index + 3)
																	var _hx_tmp_21 int32
																	_ = _hx_tmp_21
																	var _hx_tmp_22 bool
																	_ = _hx_tmp_22
																	if tmp_pos_2 < 0 {
																		_hx_tmp_22 = true
																	} else {
																		var _hx_tmp_23 int32
																		_ = _hx_tmp_23
																		_hx_tmp_23 = tmp_pos_2
																		var _hx_tmp_24 int32
																		_ = _hx_tmp_24
																		_hx_tmp_24 = _hx_tmp_23
																		_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_2).Len())
																	}

																	if _hx_tmp_22 {
																		_hx_tmp_21 = -1
																	} else {
																		_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																	}

																	tmp_tmp_1 = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_21)
																} else {
																	tmp_tmp_1 = false
																}

																if tmp_tmp_1 {
																	this.VTable.Hx_Field__add(38, true)
																	this.VTable.Hx_Field__add(38, true)
																	this.VTable.Hx_Field__add(32, true)
																} else {
																	this.VTable.Hx_Field__add(c, true)
																}
															} else {
																if c == 111 {
																	var tmp_tmp_1 bool
																	_ = tmp_tmp_1
																	var tmp_v_1 HxArray[int32]
																	_ = tmp_v_1
																	tmp_v_1 = this.Hx_Field_input
																	var tmp_pos_1 int32
																	_ = tmp_pos_1
																	tmp_pos_1 = (this.Hx_Field_index + 1)
																	var _hx_tmp_25 int32
																	_ = _hx_tmp_25
																	var _hx_tmp_26 bool
																	_ = _hx_tmp_26
																	if tmp_pos_1 < 0 {
																		_hx_tmp_26 = true
																	} else {
																		var _hx_tmp_27 int32
																		_ = _hx_tmp_27
																		_hx_tmp_27 = tmp_pos_1
																		var _hx_tmp_28 int32
																		_ = _hx_tmp_28
																		_hx_tmp_28 = _hx_tmp_27
																		_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](tmp_v_1).Len())
																	}

																	if _hx_tmp_26 {
																		_hx_tmp_25 = -1
																	} else {
																		_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																	}

																	if (_hx_tmp_25) == 114 {
																		var tmp_v_2 HxArray[int32]
																		_ = tmp_v_2
																		tmp_v_2 = this.Hx_Field_input
																		var tmp_pos_2 int32
																		_ = tmp_pos_2
																		tmp_pos_2 = (this.Hx_Field_index + 2)
																		var _hx_tmp_29 int32
																		_ = _hx_tmp_29
																		var _hx_tmp_30 bool
																		_ = _hx_tmp_30
																		if tmp_pos_2 < 0 {
																			_hx_tmp_30 = true
																		} else {
																			var _hx_tmp_31 int32
																			_ = _hx_tmp_31
																			_hx_tmp_31 = tmp_pos_2
																			var _hx_tmp_32 int32
																			_ = _hx_tmp_32
																			_hx_tmp_32 = _hx_tmp_31
																			_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_2).Len())
																		}

																		if _hx_tmp_30 {
																			_hx_tmp_29 = -1
																		} else {
																			_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																		}

																		tmp_tmp_1 = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_29)
																	} else {
																		tmp_tmp_1 = false
																	}

																	if tmp_tmp_1 {
																		this.VTable.Hx_Field__add(124, true)
																		this.VTable.Hx_Field__add(124, true)
																		this.VTable.Hx_Field__add(32, true)
																	} else {
																		this.VTable.Hx_Field__add(c, true)
																	}
																} else {
																	var tmp_tmp_1 bool
																	_ = tmp_tmp_1
																	var tmp1 bool
																	_ = tmp1
																	if c == 105 {
																		var tmp_v_1 HxArray[int32]
																		_ = tmp_v_1
																		tmp_v_1 = this.Hx_Field_input
																		var tmp_pos_1 int32
																		_ = tmp_pos_1
																		tmp_pos_1 = (this.Hx_Field_index + 1)
																		var _hx_tmp_33 int32
																		_ = _hx_tmp_33
																		var _hx_tmp_34 bool
																		_ = _hx_tmp_34
																		if tmp_pos_1 < 0 {
																			_hx_tmp_34 = true
																		} else {
																			var _hx_tmp_35 int32
																			_ = _hx_tmp_35
																			_hx_tmp_35 = tmp_pos_1
																			var _hx_tmp_36 int32
																			_ = _hx_tmp_36
																			_hx_tmp_36 = _hx_tmp_35
																			_hx_tmp_34 = (_hx_tmp_36 >= HxMakeArrayView[any](tmp_v_1).Len())
																		}

																		if _hx_tmp_34 {
																			_hx_tmp_33 = -1
																		} else {
																			_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																		}

																		tmp1 = ((_hx_tmp_33) == 115)
																	} else {
																		tmp1 = false
																	}

																	if tmp1 {
																		var tmp_v_1 HxArray[int32]
																		_ = tmp_v_1
																		tmp_v_1 = this.Hx_Field_input
																		var tmp_pos_1 int32
																		_ = tmp_pos_1
																		tmp_pos_1 = (this.Hx_Field_index + 2)
																		var _hx_tmp_37 int32
																		_ = _hx_tmp_37
																		var _hx_tmp_38 bool
																		_ = _hx_tmp_38
																		if tmp_pos_1 < 0 {
																			_hx_tmp_38 = true
																		} else {
																			var _hx_tmp_39 int32
																			_ = _hx_tmp_39
																			_hx_tmp_39 = tmp_pos_1
																			var _hx_tmp_40 int32
																			_ = _hx_tmp_40
																			_hx_tmp_40 = _hx_tmp_39
																			_hx_tmp_38 = (_hx_tmp_40 >= HxMakeArrayView[any](tmp_v_1).Len())
																		}

																		if _hx_tmp_38 {
																			_hx_tmp_37 = -1
																		} else {
																			_hx_tmp_37 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																		}

																		tmp_tmp_1 = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_37)
																	} else {
																		tmp_tmp_1 = false
																	}

																	var _hx_tmp_41 bool
																	_ = _hx_tmp_41
																	if tmp_tmp_1 {
																		_hx_tmp_41 = !this.VTable.Hx_Field_isAfterDot()
																	} else {
																		_hx_tmp_41 = false
																	}

																	if _hx_tmp_41 {
																		var p int32
																		_ = p
																		p = (this.Hx_Field_index + 2)
																		for {
																			var tmp_tmp_2 bool
																			_ = tmp_tmp_2
																			var tmp_v_1 HxArray[int32]
																			_ = tmp_v_1
																			tmp_v_1 = this.Hx_Field_input
																			var _hx_tmp_42 int32
																			_ = _hx_tmp_42
																			var _hx_tmp_43 bool
																			_ = _hx_tmp_43
																			if p < 0 {
																				_hx_tmp_43 = true
																			} else {
																				var _hx_tmp_44 int32
																				_ = _hx_tmp_44
																				_hx_tmp_44 = p
																				var _hx_tmp_45 int32
																				_ = _hx_tmp_45
																				_hx_tmp_45 = _hx_tmp_44
																				_hx_tmp_43 = (_hx_tmp_45 >= HxMakeArrayView[any](tmp_v_1).Len())
																			}

																			if _hx_tmp_43 {
																				_hx_tmp_42 = -1
																			} else {
																				_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
																			}

																			if (_hx_tmp_42) != 32 {
																				var tmp_v_2 HxArray[int32]
																				_ = tmp_v_2
																				tmp_v_2 = this.Hx_Field_input
																				var _hx_tmp_46 int32
																				_ = _hx_tmp_46
																				var _hx_tmp_47 bool
																				_ = _hx_tmp_47
																				if p < 0 {
																					_hx_tmp_47 = true
																				} else {
																					var _hx_tmp_48 int32
																					_ = _hx_tmp_48
																					_hx_tmp_48 = p
																					var _hx_tmp_49 int32
																					_ = _hx_tmp_49
																					_hx_tmp_49 = _hx_tmp_48
																					_hx_tmp_47 = (_hx_tmp_49 >= HxMakeArrayView[any](tmp_v_2).Len())
																				}

																				if _hx_tmp_47 {
																					_hx_tmp_46 = -1
																				} else {
																					_hx_tmp_46 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(p))))
																				}

																				tmp_tmp_2 = ((_hx_tmp_46) == 9)
																			} else {
																				tmp_tmp_2 = true
																			}

																			if !tmp_tmp_2 {
																				break
																			}

																			p++
																		}

																		var isNot bool
																		_ = isNot
																		var isNot1 bool
																		_ = isNot1
																		var isNot2 bool
																		_ = isNot2
																		var tmp_v_1 HxArray[int32]
																		_ = tmp_v_1
																		tmp_v_1 = this.Hx_Field_input
																		var _hx_tmp_50 int32
																		_ = _hx_tmp_50
																		var _hx_tmp_51 bool
																		_ = _hx_tmp_51
																		if p < 0 {
																			_hx_tmp_51 = true
																		} else {
																			var _hx_tmp_52 int32
																			_ = _hx_tmp_52
																			_hx_tmp_52 = p
																			var _hx_tmp_53 int32
																			_ = _hx_tmp_53
																			_hx_tmp_53 = _hx_tmp_52
																			_hx_tmp_51 = (_hx_tmp_53 >= HxMakeArrayView[any](tmp_v_1).Len())
																		}

																		if _hx_tmp_51 {
																			_hx_tmp_50 = -1
																		} else {
																			_hx_tmp_50 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
																		}

																		if (_hx_tmp_50) == 110 {
																			var tmp_v_2 HxArray[int32]
																			_ = tmp_v_2
																			tmp_v_2 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (p + 1)
																			var _hx_tmp_54 int32
																			_ = _hx_tmp_54
																			var _hx_tmp_55 bool
																			_ = _hx_tmp_55
																			if tmp_pos_1 < 0 {
																				_hx_tmp_55 = true
																			} else {
																				var _hx_tmp_56 int32
																				_ = _hx_tmp_56
																				_hx_tmp_56 = tmp_pos_1
																				var _hx_tmp_57 int32
																				_ = _hx_tmp_57
																				_hx_tmp_57 = _hx_tmp_56
																				_hx_tmp_55 = (_hx_tmp_57 >= HxMakeArrayView[any](tmp_v_2).Len())
																			}

																			if _hx_tmp_55 {
																				_hx_tmp_54 = -1
																			} else {
																				_hx_tmp_54 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
																			}

																			isNot2 = ((_hx_tmp_54) == 111)
																		} else {
																			isNot2 = false
																		}

																		if isNot2 {
																			var tmp_v_2 HxArray[int32]
																			_ = tmp_v_2
																			tmp_v_2 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (p + 2)
																			var _hx_tmp_58 int32
																			_ = _hx_tmp_58
																			var _hx_tmp_59 bool
																			_ = _hx_tmp_59
																			if tmp_pos_1 < 0 {
																				_hx_tmp_59 = true
																			} else {
																				var _hx_tmp_60 int32
																				_ = _hx_tmp_60
																				_hx_tmp_60 = tmp_pos_1
																				var _hx_tmp_61 int32
																				_ = _hx_tmp_61
																				_hx_tmp_61 = _hx_tmp_60
																				_hx_tmp_59 = (_hx_tmp_61 >= HxMakeArrayView[any](tmp_v_2).Len())
																			}

																			if _hx_tmp_59 {
																				_hx_tmp_58 = -1
																			} else {
																				_hx_tmp_58 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
																			}

																			isNot1 = ((_hx_tmp_58) == 116)
																		} else {
																			isNot1 = false
																		}

																		if isNot1 {
																			var tmp_v_2 HxArray[int32]
																			_ = tmp_v_2
																			tmp_v_2 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (p + 3)
																			var _hx_tmp_62 int32
																			_ = _hx_tmp_62
																			var _hx_tmp_63 bool
																			_ = _hx_tmp_63
																			if tmp_pos_1 < 0 {
																				_hx_tmp_63 = true
																			} else {
																				var _hx_tmp_64 int32
																				_ = _hx_tmp_64
																				_hx_tmp_64 = tmp_pos_1
																				var _hx_tmp_65 int32
																				_ = _hx_tmp_65
																				_hx_tmp_65 = _hx_tmp_64
																				_hx_tmp_63 = (_hx_tmp_65 >= HxMakeArrayView[any](tmp_v_2).Len())
																			}

																			if _hx_tmp_63 {
																				_hx_tmp_62 = -1
																			} else {
																				_hx_tmp_62 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
																			}

																			isNot = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_62)
																		} else {
																			isNot = false
																		}

																		if isNot {
																			this.VTable.Hx_Field__add(33, true)
																			this.VTable.Hx_Field__add(61, true)
																			for this.Hx_Field_index < (p + 3) {
																				this.VTable.Hx_Field__add(32, true)
																			}
																		} else {
																			this.VTable.Hx_Field__add(61, true)
																			this.VTable.Hx_Field__add(61, true)
																		}
																	} else {
																		var tmp_tmp_2 bool
																		_ = tmp_tmp_2
																		var tmp_tmp1_1 bool
																		_ = tmp_tmp1_1
																		var tmp2 bool
																		_ = tmp2
																		if c == 110 {
																			var tmp_v_1 HxArray[int32]
																			_ = tmp_v_1
																			tmp_v_1 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (this.Hx_Field_index + 1)
																			var _hx_tmp_66 int32
																			_ = _hx_tmp_66
																			var _hx_tmp_67 bool
																			_ = _hx_tmp_67
																			if tmp_pos_1 < 0 {
																				_hx_tmp_67 = true
																			} else {
																				var _hx_tmp_68 int32
																				_ = _hx_tmp_68
																				_hx_tmp_68 = tmp_pos_1
																				var _hx_tmp_69 int32
																				_ = _hx_tmp_69
																				_hx_tmp_69 = _hx_tmp_68
																				_hx_tmp_67 = (_hx_tmp_69 >= HxMakeArrayView[any](tmp_v_1).Len())
																			}

																			if _hx_tmp_67 {
																				_hx_tmp_66 = -1
																			} else {
																				_hx_tmp_66 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																			}

																			tmp2 = ((_hx_tmp_66) == 111)
																		} else {
																			tmp2 = false
																		}

																		if tmp2 {
																			var tmp_v_1 HxArray[int32]
																			_ = tmp_v_1
																			tmp_v_1 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (this.Hx_Field_index + 2)
																			var _hx_tmp_70 int32
																			_ = _hx_tmp_70
																			var _hx_tmp_71 bool
																			_ = _hx_tmp_71
																			if tmp_pos_1 < 0 {
																				_hx_tmp_71 = true
																			} else {
																				var _hx_tmp_72 int32
																				_ = _hx_tmp_72
																				_hx_tmp_72 = tmp_pos_1
																				var _hx_tmp_73 int32
																				_ = _hx_tmp_73
																				_hx_tmp_73 = _hx_tmp_72
																				_hx_tmp_71 = (_hx_tmp_73 >= HxMakeArrayView[any](tmp_v_1).Len())
																			}

																			if _hx_tmp_71 {
																				_hx_tmp_70 = -1
																			} else {
																				_hx_tmp_70 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																			}

																			tmp_tmp1_1 = ((_hx_tmp_70) == 116)
																		} else {
																			tmp_tmp1_1 = false
																		}

																		if tmp_tmp1_1 {
																			var tmp_v_1 HxArray[int32]
																			_ = tmp_v_1
																			tmp_v_1 = this.Hx_Field_input
																			var tmp_pos_1 int32
																			_ = tmp_pos_1
																			tmp_pos_1 = (this.Hx_Field_index + 3)
																			var _hx_tmp_74 int32
																			_ = _hx_tmp_74
																			var _hx_tmp_75 bool
																			_ = _hx_tmp_75
																			if tmp_pos_1 < 0 {
																				_hx_tmp_75 = true
																			} else {
																				var _hx_tmp_76 int32
																				_ = _hx_tmp_76
																				_hx_tmp_76 = tmp_pos_1
																				var _hx_tmp_77 int32
																				_ = _hx_tmp_77
																				_hx_tmp_77 = _hx_tmp_76
																				_hx_tmp_75 = (_hx_tmp_77 >= HxMakeArrayView[any](tmp_v_1).Len())
																			}

																			if _hx_tmp_75 {
																				_hx_tmp_74 = -1
																			} else {
																				_hx_tmp_74 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
																			}

																			tmp_tmp_2 = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_74)
																		} else {
																			tmp_tmp_2 = false
																		}

																		var _hx_tmp_78 bool
																		_ = _hx_tmp_78
																		if tmp_tmp_2 {
																			_hx_tmp_78 = !this.VTable.Hx_Field_isAfterDot()
																		} else {
																			_hx_tmp_78 = false
																		}

																		if _hx_tmp_78 {
																			this.VTable.Hx_Field__add(33, true)
																			this.VTable.Hx_Field__add(32, true)
																			this.VTable.Hx_Field__add(32, true)
																		} else {
																			this.VTable.Hx_Field__add(c, true)
																		}
																	}
																}
															}
														} else {
															this.VTable.Hx_Field__add(c, true)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		if until == -1 {
			for {
				var _hx_tmp_79 bool
				_ = _hx_tmp_79
				if this.Hx_Field_stack.Len() > 0 {
					var _hx_tmp_80 bool
					_ = _hx_tmp_80
					var _hx_tmp_81 HxArray[int32]
					_ = _hx_tmp_81
					_hx_tmp_81 = this.Hx_Field_stack
					if _hx_tmp_81.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(3)) {
						_hx_tmp_80 = true
					} else {
						var _hx_tmp_82 HxArray[int32]
						_ = _hx_tmp_82
						_hx_tmp_82 = this.Hx_Field_stack
						var _hx_tmp_83 HxArray[int32]
						_ = _hx_tmp_83
						_hx_tmp_83 = _hx_tmp_82
						_hx_tmp_80 = (_hx_tmp_83.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(4)))
					}

					_hx_tmp_79 = (_hx_tmp_80)
				} else {
					_hx_tmp_79 = false
				}

				if !(_hx_tmp_79) {
					break
				}

				var popped HxNullable[int32]
				_ = popped
				popped = (Hx_Array_Pop[int32](this.Hx_Field_stack))
				_ = (Hx_Array_Pop[int32](this.Hx_Field_indentStack))
				var _hx_tmp_84 bool
				_ = _hx_tmp_84
				if (popped).Valid == true {
					_hx_tmp_84 = ((popped).Value == ((int32)(3)))
				} else {
					_hx_tmp_84 = false
				}

				if _hx_tmp_84 {
					this.Hx_Field_currentPosOffset++
					this.Hx_Field_output.VTable.Hx_Field_addChar(32)
					_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
					this.Hx_Field_currentPosOffset++
					this.Hx_Field_output.VTable.Hx_Field_addChar(125)
					_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processString() {
	this.Hx_Field_inString = true
	this.VTable.Hx_Field__add(34, true)
	var escaped bool
	_ = escaped
	escaped = false
	for this.Hx_Field_index < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_index
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if pos < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = pos
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_0
		if escaped {
			this.VTable.Hx_Field__add(c, true)
			escaped = false
		} else {
			if c == 92 {
				escaped = true
				this.VTable.Hx_Field__add(92, true)
			} else {
				if c == 34 {
					this.VTable.Hx_Field__add(c, true)
					this.Hx_Field_inString = false
					return
				} else {
					if (c == 36) && !escaped {
						this.Hx_Field_index++
						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = this.Hx_Field_index
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						var _hx_tmp_5 bool
						_ = _hx_tmp_5
						if tmp_pos_1 < 0 {
							_hx_tmp_5 = true
						} else {
							var _hx_tmp_6 int32
							_ = _hx_tmp_6
							_hx_tmp_6 = tmp_pos_1
							var _hx_tmp_7 int32
							_ = _hx_tmp_7
							_hx_tmp_7 = _hx_tmp_6
							_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_5 {
							_hx_tmp_4 = -1
						} else {
							_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
						}

						c = _hx_tmp_4
						if c == 123 {
							this.Hx_Field_currentPosOffset--
							this.VTable.Hx_Field__add(34, false)
							this.VTable.Hx_Field__add(43, false)
							this.VTable.Hx_Field__add(40, true)
							this.Hx_Field_inString = false
							this.VTable.Hx_Field_processComplexInterpolation()
							this.Hx_Field_inString = true
							this.VTable.Hx_Field__add(41, true)
							this.Hx_Field_currentPosOffset--
							this.VTable.Hx_Field__add(43, false)
							this.VTable.Hx_Field__add(34, false)
							this.Hx_Field_currentPosOffset++
						} else {
							if (((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95) {
								this.Hx_Field_currentPosOffset--
								this.VTable.Hx_Field__add(34, false)
								this.VTable.Hx_Field__add(43, false)
								this.VTable.Hx_Field_processFieldAccessInterpolation()
								this.Hx_Field_currentPosOffset--
								this.VTable.Hx_Field__add(43, false)
								this.VTable.Hx_Field__add(34, false)
								this.Hx_Field_currentPosOffset++
							} else {
								if c == 36 {
									this.VTable.Hx_Field__add(36, true)
								} else {
									this.VTable.Hx_Field_error("Expected identifier or { after $")
								}
							}
						}
					} else {
						this.VTable.Hx_Field__add(c, true)
					}
				}
			}
		}
	}

	this.VTable.Hx_Field_error("Unterminated string")
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processComplexInterpolation() {
	this.VTable.Hx_Field_processInput(125)
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processFieldAccessInterpolation() {
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_index
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if pos < 0 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = pos
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = -1
	} else {
		_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_0
	if !((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) {
		this.VTable.Hx_Field_error("Expected identifier in field access")
	}

	this.VTable.Hx_Field_processIdentifier()
hx_label_0:
	for this.Hx_Field_index < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_index
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if tmp_pos_1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = tmp_pos_1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
		}

		switch _hx_tmp_4 {
		case 40:
			{
				this.VTable.Hx_Field__add(40, true)
				this.VTable.Hx_Field_processInput(41)
				this.VTable.Hx_Field__add(41, true)
			}

		case 46:
			{
				var tmp bool
				_ = tmp
				if (this.Hx_Field_index + 1) < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = (this.Hx_Field_index + 1)
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					var _hx_tmp_9 bool
					_ = _hx_tmp_9
					if tmp_pos_2 < 0 {
						_hx_tmp_9 = true
					} else {
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						_hx_tmp_10 = tmp_pos_2
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = _hx_tmp_10
						_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_9 {
						_hx_tmp_8 = -1
					} else {
						_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_8
					tmp = ((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95))
				} else {
					tmp = false
				}

				if tmp {
					this.VTable.Hx_Field__add(46, true)
					this.VTable.Hx_Field_processIdentifier()
				} else {
					break hx_label_0
				}
			}

		case 91:
			{
				this.VTable.Hx_Field__add(91, true)
				this.VTable.Hx_Field_processInput(93)
			}

		default:
			{
				break hx_label_0
			}
		}
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processIdentifier() {
	for this.Hx_Field_index < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_index
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if pos < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = pos
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_0
		if !(((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57))) {
			break
		}

		this.VTable.Hx_Field__add(c, true)
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isDigit(c int32) bool {
	return ((c >= 48) && (c <= 57))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isIdentifierStart(c int32) bool {
	return ((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isIdentifierPart(c int32) bool {
	return (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_processComment() {
	this.Hx_Field_inComment = true
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = (this.Hx_Field_index + 1)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if pos < 0 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = pos
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = -1
	} else {
		_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_0
	if c == 47 {
		this.VTable.Hx_Field__add(32, true)
		this.VTable.Hx_Field__add(32, true)
		for this.Hx_Field_index < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_index
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if tmp_pos_1 < 0 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = tmp_pos_1
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _hx_tmp_6
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			var cc int32
			_ = cc
			cc = _hx_tmp_4
			if (cc == 13) || (cc == 10) {
				break
			}

			this.VTable.Hx_Field__add(32, true)
		}
	} else {
		if c == 42 {
			this.VTable.Hx_Field__add(32, true)
			this.VTable.Hx_Field__add(32, true)
			for this.Hx_Field_index < this.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_index
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				var _hx_tmp_9 bool
				_ = _hx_tmp_9
				if tmp_pos_1 < 0 {
					_hx_tmp_9 = true
				} else {
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					_hx_tmp_10 = tmp_pos_1
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					_hx_tmp_11 = _hx_tmp_10
					_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_9 {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				var cc int32
				_ = cc
				cc = _hx_tmp_8
				var tmp bool
				_ = tmp
				if (cc == 42) && ((this.Hx_Field_index + 1) < this.Hx_Field_length) {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = (this.Hx_Field_index + 1)
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					var _hx_tmp_13 bool
					_ = _hx_tmp_13
					if tmp_pos_2 < 0 {
						_hx_tmp_13 = true
					} else {
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = tmp_pos_2
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = _hx_tmp_14
						_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_13 {
						_hx_tmp_12 = -1
					} else {
						_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					tmp = ((_hx_tmp_12) == 47)
				} else {
					tmp = false
				}

				if tmp {
					this.VTable.Hx_Field__add(32, true)
					this.VTable.Hx_Field__add(32, true)
					break
				} else {
					if (cc == 13) || (cc == 10) {
						this.VTable.Hx_Field__add(cc, true)
					} else {
						this.VTable.Hx_Field__add(32, true)
					}
				}
			}
		}
	}

	this.Hx_Field_inComment = false
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_endsWithChar(line string, c int32) bool {
	if int32(utf8.RuneCountInString(line)) == 0 {
		return false
	}

	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(line))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(line, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = line
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = line
	}

	var trimmed string
	_ = trimmed
	trimmed = _hx_tmp_1
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(trimmed))
	var result bool
	_ = result
	if _hx_reserved_len > 0 {
		var c1 HxNullable[int32]
		_ = c1
		c1 = (Hx_Field_go_haxe_hxstring_charCodeAt(trimmed, (_hx_reserved_len - 1)))
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		if (c1).Valid == false {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = (c1).Value
		}

		result = ((_hx_tmp_4) == c)
	} else {
		result = false
	}

	return result
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_followsWithChar(c int32, pos int32) bool {
	var tempIndex int32
	_ = tempIndex
	tempIndex = pos
	var result bool
	_ = result
	result = false
	for tempIndex < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if tempIndex < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = tempIndex
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(tempIndex))))
		}

		var cc int32
		_ = cc
		cc = _hx_tmp_0
		if (((cc == 32) || (cc == 9)) || (cc == 10)) || (cc == 13) {
			tempIndex++
			continue
		}

		if (cc == 47) && ((tempIndex + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (tempIndex + 1)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if pos < 0 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = pos
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _hx_tmp_6
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
			}

			var nextChar int32
			_ = nextChar
			nextChar = _hx_tmp_4
			if nextChar == 47 {
				tempIndex = (tempIndex + 2)
				for {
					var tmp bool
					_ = tmp
					var tmp1 bool
					_ = tmp1
					if tempIndex < this.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						var _hx_tmp_9 bool
						_ = _hx_tmp_9
						if tempIndex < 0 {
							_hx_tmp_9 = true
						} else {
							var _hx_tmp_10 int32
							_ = _hx_tmp_10
							_hx_tmp_10 = tempIndex
							var _hx_tmp_11 int32
							_ = _hx_tmp_11
							_hx_tmp_11 = _hx_tmp_10
							_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_9 {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
						}

						tmp1 = ((_hx_tmp_8) != 13)
					} else {
						tmp1 = false
					}

					if tmp1 {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						var _hx_tmp_13 bool
						_ = _hx_tmp_13
						if tempIndex < 0 {
							_hx_tmp_13 = true
						} else {
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							_hx_tmp_14 = tempIndex
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							_hx_tmp_15 = _hx_tmp_14
							_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_13 {
							_hx_tmp_12 = -1
						} else {
							_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
						}

						tmp = ((_hx_tmp_12) != 10)
					} else {
						tmp = false
					}

					if !tmp {
						break
					}

					tempIndex++
				}

				if tempIndex < this.Hx_Field_length {
					tempIndex++
				}

				continue
			}

			if nextChar == 42 {
				tempIndex = (tempIndex + 2)
				for (tempIndex + 1) < this.Hx_Field_length {
					var tmp bool
					_ = tmp
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					var _hx_tmp_17 bool
					_ = _hx_tmp_17
					if tempIndex < 0 {
						_hx_tmp_17 = true
					} else {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = tempIndex
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = _hx_tmp_18
						_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_17 {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
					}

					if (_hx_tmp_16) == 42 {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = (tempIndex + 1)
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						var _hx_tmp_21 bool
						_ = _hx_tmp_21
						if tmp_pos_1 < 0 {
							_hx_tmp_21 = true
						} else {
							var _hx_tmp_22 int32
							_ = _hx_tmp_22
							_hx_tmp_22 = tmp_pos_1
							var _hx_tmp_23 int32
							_ = _hx_tmp_23
							_hx_tmp_23 = _hx_tmp_22
							_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_21 {
							_hx_tmp_20 = -1
						} else {
							_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_1))))
						}

						tmp = ((_hx_tmp_20) == 47)
					} else {
						tmp = false
					}

					if tmp {
						tempIndex = (tempIndex + 2)
						break
					}

					tempIndex++
				}

				continue
			}
		}

		result = (cc == c)
		break
	}

	return result
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_endsOrFollowsWithChar(line string, c int32, pos int32, hxpos__isOptional any) bool {
	hxpos := HxDefault[any]()
	_ = hxpos

	if hxpos__isOptional != nil {
		hxpos = HxConvert[any](hxpos__isOptional)
	} else {
		hxpos = nil
	}
	{
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if this.VTable.Hx_Field_endsWithChar(line, c) {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = this.VTable.Hx_Field_followsWithChar(c, pos)
		}

		return _hx_tmp_0
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_nextLineIndentOffset(line string, pos int32) int32 {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(line))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(line, r)
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = line
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = r
		var _hx_tmp_4 string
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_2
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		_hx_tmp_5 = _hx_tmp_3
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = line
	}

	var currentLine string
	_ = currentLine
	currentLine = _hx_tmp_1
	if int32(utf8.RuneCountInString(currentLine)) == 0 {
		return 0
	}

	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	_hx_tmp_6 = int32(utf8.RuneCountInString(line))
	var currentIndent int32
	_ = currentIndent
	currentIndent = (_hx_tmp_6 - int32(utf8.RuneCountInString(currentLine)))
	var firstSkippedCommentIndent int32
	_ = firstSkippedCommentIndent
	firstSkippedCommentIndent = -1
	for pos < this.Hx_Field_length {
		var _hx_tmp_7 HxArray[int32]
		_ = _hx_tmp_7
		_hx_tmp_7 = this.Hx_Field_input
		var endLine int32
		_ = endLine
		endLine = Hx_Field_loreline_utf8vector_uIndexOf(_hx_tmp_7, "\n", HxNullable[int32]{Value: pos, Valid: true})
		if endLine == -1 {
			endLine = this.Hx_Field_length
		}

		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_8 HxArray[int32]
		_ = _hx_tmp_8
		_hx_tmp_8 = v
		var _hx_tmp_9 int32
		_ = _hx_tmp_9
		_hx_tmp_9 = pos
		var _hx_tmp_10 int32
		_ = _hx_tmp_10
		if (HxNullable[int32]{Value: endLine, Valid: true}).Valid != false {
			_hx_tmp_10 = endLine
		} else {
			_hx_tmp_10 = HxMakeArrayView[any](v).Len()
		}

		var nextLine string
		_ = nextLine
		nextLine = Hx_Field_loreline_utf8vector_slice(_hx_tmp_8, _hx_tmp_9, _hx_tmp_10)
		var tmp_l_1 int32
		_ = tmp_l_1
		tmp_l_1 = int32(utf8.RuneCountInString(nextLine))
		var tmp_r_1 int32
		_ = tmp_r_1
		tmp_r_1 = 0
		for {
			var _hx_tmp_11 bool
			_ = _hx_tmp_11
			if tmp_r_1 < tmp_l_1 {
				_hx_tmp_11 = Hx_Field_stringtools_isSpace(nextLine, tmp_r_1)
			} else {
				_hx_tmp_11 = false
			}

			if !(_hx_tmp_11) {
				break
			}

			tmp_r_1++
		}

		var _hx_tmp_12 string
		_ = _hx_tmp_12
		if tmp_r_1 > 0 {
			var _hx_tmp_13 string
			_ = _hx_tmp_13
			_hx_tmp_13 = nextLine
			var _hx_tmp_14 int32
			_ = _hx_tmp_14
			_hx_tmp_14 = tmp_r_1
			var _hx_tmp_15 string
			_ = _hx_tmp_15
			_hx_tmp_15 = _hx_tmp_13
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = _hx_tmp_14
			_hx_tmp_12 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, _hx_tmp_16, HxNullable[int32]{Value: (tmp_l_1 - tmp_r_1), Valid: true}))
		} else {
			_hx_tmp_12 = nextLine
		}

		var trimmed string
		_ = trimmed
		trimmed = _hx_tmp_12
		var _hx_tmp_17 bool
		_ = _hx_tmp_17
		var _hx_tmp_18 bool
		_ = _hx_tmp_18
		if int32(utf8.RuneCountInString(trimmed)) > 0 {
			var _hx_tmp_19 bool
			_ = _hx_tmp_19
			if !Hx_Field_stringtools_startsWith(trimmed, "//") {
				_hx_tmp_19 = true
			} else {
				_hx_tmp_19 = Hx_Field_stringtools_startsWith(trimmed, "//<END>")
			}

			_hx_tmp_18 = (_hx_tmp_19)
		} else {
			_hx_tmp_18 = false
		}

		if _hx_tmp_18 {
			_hx_tmp_17 = !Hx_Field_stringtools_startsWith(trimmed, "/*")
		} else {
			_hx_tmp_17 = false
		}

		if _hx_tmp_17 {
			var _hx_tmp_20 int32
			_ = _hx_tmp_20
			_hx_tmp_20 = int32(utf8.RuneCountInString(nextLine))
			var nextIndent int32
			_ = nextIndent
			nextIndent = (_hx_tmp_20 - int32(utf8.RuneCountInString(trimmed)))
			var _hx_tmp_21 bool
			_ = _hx_tmp_21
			var _hx_tmp_22 bool
			_ = _hx_tmp_22
			if Hx_Field_stringtools_startsWith(trimmed, "//<END>") {
				_hx_tmp_22 = (currentIndent == 0)
			} else {
				_hx_tmp_22 = false
			}

			if _hx_tmp_22 {
				_hx_tmp_21 = (firstSkippedCommentIndent > 0)
			} else {
				_hx_tmp_21 = false
			}

			if _hx_tmp_21 {
				return (firstSkippedCommentIndent - currentIndent)
			}

			return (nextIndent - currentIndent)
		}

		var _hx_tmp_23 bool
		_ = _hx_tmp_23
		if int32(utf8.RuneCountInString(trimmed)) > 0 {
			_hx_tmp_23 = (firstSkippedCommentIndent == -1)
		} else {
			_hx_tmp_23 = false
		}

		if _hx_tmp_23 {
			var _hx_tmp_24 int32
			_ = _hx_tmp_24
			_hx_tmp_24 = int32(utf8.RuneCountInString(nextLine))
			firstSkippedCommentIndent = (_hx_tmp_24 - int32(utf8.RuneCountInString(trimmed)))
		}

		pos = (endLine + 1)
	}

	if (currentIndent == 0) && (firstSkippedCommentIndent > 0) {
		return (firstSkippedCommentIndent - currentIndent)
	}

	return 0
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_followsWithIf(pos int32) bool {
	return this.VTable.Hx_Field_followsWithWord("if", pos)
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_followsWithWord(word string, pos int32) bool {
	var tempIndex int32
	_ = tempIndex
	tempIndex = pos
	for tempIndex < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if tempIndex < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = tempIndex
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(tempIndex))))
		}

		var cc int32
		_ = cc
		cc = _hx_tmp_0
		if (((cc == 32) || (cc == 9)) || (cc == 10)) || (cc == 13) {
			tempIndex++
			continue
		}

		if (cc == 47) && ((tempIndex + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (tempIndex + 1)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if pos < 0 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = pos
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _hx_tmp_6
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
			}

			var nextChar int32
			_ = nextChar
			nextChar = _hx_tmp_4
			if nextChar == 47 {
				tempIndex = (tempIndex + 2)
				for {
					var tmp bool
					_ = tmp
					var tmp1 bool
					_ = tmp1
					if tempIndex < this.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						var _hx_tmp_9 bool
						_ = _hx_tmp_9
						if tempIndex < 0 {
							_hx_tmp_9 = true
						} else {
							var _hx_tmp_10 int32
							_ = _hx_tmp_10
							_hx_tmp_10 = tempIndex
							var _hx_tmp_11 int32
							_ = _hx_tmp_11
							_hx_tmp_11 = _hx_tmp_10
							_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_9 {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
						}

						tmp1 = ((_hx_tmp_8) != 13)
					} else {
						tmp1 = false
					}

					if tmp1 {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						var _hx_tmp_13 bool
						_ = _hx_tmp_13
						if tempIndex < 0 {
							_hx_tmp_13 = true
						} else {
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							_hx_tmp_14 = tempIndex
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							_hx_tmp_15 = _hx_tmp_14
							_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_13 {
							_hx_tmp_12 = -1
						} else {
							_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
						}

						tmp = ((_hx_tmp_12) != 10)
					} else {
						tmp = false
					}

					if !tmp {
						break
					}

					tempIndex++
				}

				continue
			}

			if nextChar == 42 {
				tempIndex = (tempIndex + 2)
				for (tempIndex + 1) < this.Hx_Field_length {
					var tmp bool
					_ = tmp
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					var _hx_tmp_17 bool
					_ = _hx_tmp_17
					if tempIndex < 0 {
						_hx_tmp_17 = true
					} else {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = tempIndex
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = _hx_tmp_18
						_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_17 {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tempIndex))))
					}

					if (_hx_tmp_16) == 42 {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = (tempIndex + 1)
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						var _hx_tmp_21 bool
						_ = _hx_tmp_21
						if tmp_pos_1 < 0 {
							_hx_tmp_21 = true
						} else {
							var _hx_tmp_22 int32
							_ = _hx_tmp_22
							_hx_tmp_22 = tmp_pos_1
							var _hx_tmp_23 int32
							_ = _hx_tmp_23
							_hx_tmp_23 = _hx_tmp_22
							_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_21 {
							_hx_tmp_20 = -1
						} else {
							_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_1))))
						}

						tmp = ((_hx_tmp_20) == 47)
					} else {
						tmp = false
					}

					if tmp {
						tempIndex = (tempIndex + 2)
						break
					}

					tempIndex++
				}

				continue
			}
		}

		var wordLength int32
		_ = wordLength
		wordLength = int32(utf8.RuneCountInString(word))
		if (tempIndex + wordLength) <= this.Hx_Field_length {
			var matches bool
			_ = matches
			matches = true
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = wordLength
				for _g < _g1 {
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = _g
					_g = (_g + 1)
					var i int32
					_ = i
					i = _hx_tmp_24
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var pos int32
					_ = pos
					pos = (tempIndex + i)
					var hx_cmp_2 HxNullable[int32]
					_ = hx_cmp_2
					hx_cmp_2 = (Hx_Field_go_haxe_hxstring_charCodeAt(word, i))
					var _hx_tmp_25 bool
					_ = _hx_tmp_25
					if (hx_cmp_2).Valid == false {
						_hx_tmp_25 = true
					} else {
						var _hx_tmp_26 int32
						_ = _hx_tmp_26
						var _hx_tmp_27 bool
						_ = _hx_tmp_27
						if pos < 0 {
							_hx_tmp_27 = true
						} else {
							var _hx_tmp_28 int32
							_ = _hx_tmp_28
							_hx_tmp_28 = pos
							var _hx_tmp_29 int32
							_ = _hx_tmp_29
							_hx_tmp_29 = _hx_tmp_28
							_hx_tmp_27 = (_hx_tmp_29 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_27 {
							_hx_tmp_26 = -1
						} else {
							_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
						}

						_hx_tmp_25 = ((_hx_tmp_26) != (hx_cmp_2).Value)
					}

					if _hx_tmp_25 {
						matches = false
						break
					}
				}
			}

			var tmp bool
			_ = tmp
			if matches {
				if (tempIndex + wordLength) < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var pos int32
					_ = pos
					pos = (tempIndex + wordLength)
					var _hx_tmp_30 int32
					_ = _hx_tmp_30
					var _hx_tmp_31 bool
					_ = _hx_tmp_31
					if pos < 0 {
						_hx_tmp_31 = true
					} else {
						var _hx_tmp_32 int32
						_ = _hx_tmp_32
						_hx_tmp_32 = pos
						var _hx_tmp_33 int32
						_ = _hx_tmp_33
						_hx_tmp_33 = _hx_tmp_32
						_hx_tmp_31 = (_hx_tmp_33 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_31 {
						_hx_tmp_30 = -1
					} else {
						_hx_tmp_30 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					tmp = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_30)
				} else {
					tmp = true
				}
			} else {
				tmp = false
			}

			if tmp {
				return true
			}
		}

		return false
	}

	return false
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_endsWithControlKeyword(line string, pos int32) bool {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(line))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(line, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = line
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = line
	}

	var trimmed string
	_ = trimmed
	trimmed = _hx_tmp_1
	if int32(utf8.RuneCountInString(trimmed)) == 0 {
		return false
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS
		for {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g
			if !(_hx_tmp_4 < _g1.Len()) {
				break
			}

			var keyword string
			_ = keyword
			keyword = _g1.Get(_g)
			_g++
			if Hx_Field_stringtools_endsWith(trimmed, keyword) {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = int32(utf8.RuneCountInString(line))
				var trimmedSize int32
				_ = trimmedSize
				trimmedSize = (_hx_tmp_5 - int32(utf8.RuneCountInString(trimmed)))
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = int32(utf8.RuneCountInString(trimmed))
				var keywordPos int32
				_ = keywordPos
				keywordPos = (_hx_tmp_6 - int32(utf8.RuneCountInString(keyword)))
				var hasValidPrefix bool
				_ = hasValidPrefix
				if keywordPos != 0 {
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(trimmed, (keywordPos - 1)))
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					if (c).Valid == false {
						_hx_tmp_7 = -1
					} else {
						_hx_tmp_7 = (c).Value
					}

					hasValidPrefix = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_7)
				} else {
					hasValidPrefix = true
				}

				var nextPos int32
				_ = nextPos
				nextPos = pos
				var hasValidSuffix bool
				_ = hasValidSuffix
				if nextPos < this.Hx_Field_length {
					if trimmedSize <= 0 {
						var v HxArray[int32]
						_ = v
						v = this.Hx_Field_input
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						var _hx_tmp_9 bool
						_ = _hx_tmp_9
						if nextPos < 0 {
							_hx_tmp_9 = true
						} else {
							var _hx_tmp_10 int32
							_ = _hx_tmp_10
							_hx_tmp_10 = nextPos
							var _hx_tmp_11 int32
							_ = _hx_tmp_11
							_hx_tmp_11 = _hx_tmp_10
							_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
						}

						if _hx_tmp_9 {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(nextPos))))
						}

						hasValidSuffix = !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_8)
					} else {
						hasValidSuffix = true
					}
				} else {
					hasValidSuffix = false
				}

				if hasValidPrefix && hasValidSuffix {
					var _hx_tmp_12 bool
					_ = _hx_tmp_12
					if keyword != "else" {
						_hx_tmp_12 = true
					} else {
						_hx_tmp_12 = !this.VTable.Hx_Field_followsWithIf(pos)
					}

					if _hx_tmp_12 {
						return true
					}
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_endsWithWord(line string, word string) bool {
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(line))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if r < l {
			_hx_tmp_0 = Hx_Field_stringtools_isSpace(line, ((l - r) - 1))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		r++
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if r > 0 {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = line
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_1 = line
	}

	var trimmed string
	_ = trimmed
	trimmed = _hx_tmp_1
	if !Hx_Field_stringtools_endsWith(trimmed, word) {
		return false
	}

	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = int32(utf8.RuneCountInString(trimmed))
	var wordPos int32
	_ = wordPos
	wordPos = (_hx_tmp_4 - int32(utf8.RuneCountInString(word)))
	if wordPos != 0 {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(trimmed, (wordPos - 1)))
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		if (c).Valid == false {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = (c).Value
		}

		return !this.VTable.Hx_Field_isAlphaNumeric(_hx_tmp_5)
	} else {
		return true
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_endsWithArrayIndexable(line string) bool {
	var length int32
	_ = length
	length = int32(utf8.RuneCountInString(line))
	if length == 0 {
		return false
	}

	var lastNonWhitespacePos int32
	_ = lastNonWhitespacePos
	lastNonWhitespacePos = (length - 1)
	for lastNonWhitespacePos >= 0 {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, lastNonWhitespacePos))
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (c).Valid == false {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_0
		if !this.VTable.Hx_Field_isWhitespace(c1) {
			break
		}

		lastNonWhitespacePos--
	}

	if lastNonWhitespacePos < 0 {
		return false
	}

	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, lastNonWhitespacePos))
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if (c).Valid == false {
		_hx_tmp_1 = -1
	} else {
		_hx_tmp_1 = (c).Value
	}

	var lastChar int32
	_ = lastChar
	lastChar = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if ((lastChar == 41) || (lastChar == 125)) || (lastChar == 93) {
		_hx_tmp_2 = true
	} else {
		_hx_tmp_2 = this.VTable.Hx_Field_isAlphaNumeric(lastChar)
	}

	return _hx_tmp_2
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isWhitespace(c int32) bool {
	return ((((c == 32) || (c == 10)) || (c == 9)) || (c == 13))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isAfterDot() bool {
	var p int32
	_ = p
	p = (this.Hx_Field_index - 1)
	for {
		var tmp bool
		_ = tmp
		if p >= 0 {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if p < 0 {
				_hx_tmp_1 = true
			} else {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = p
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _hx_tmp_2
				_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_1 {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
			}

			if (_hx_tmp_0) != 32 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if p < 0 {
					_hx_tmp_5 = true
				} else {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = p
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _hx_tmp_6
					_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_5 {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
				}

				tmp = ((_hx_tmp_4) == 9)
			} else {
				tmp = true
			}
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		p--
	}

	if p >= 0 {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_8 int32
		_ = _hx_tmp_8
		var _hx_tmp_9 bool
		_ = _hx_tmp_9
		if p < 0 {
			_hx_tmp_9 = true
		} else {
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = p
			var _hx_tmp_11 int32
			_ = _hx_tmp_11
			_hx_tmp_11 = _hx_tmp_10
			_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_9 {
			_hx_tmp_8 = -1
		} else {
			_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		return ((_hx_tmp_8) == 46)
	} else {
		return false
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isAlphaNumeric(c int32) bool {
	return (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 95))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_inStatementsBlock() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	if i >= 0 {
		return (((this.Hx_Field_stack.Get(i) == ((int32)(2))) || (this.Hx_Field_stack.Get(i) == ((int32)(3)))) || (this.Hx_Field_stack.Get(i) == ((int32)(4))))
	}

	return true
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_inObjectBlock() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		if (this.Hx_Field_stack.Get(i) != ((int32)(3))) && (this.Hx_Field_stack.Get(i) != ((int32)(4))) {
			var res bool
			_ = res
			res = (this.Hx_Field_stack.Get(i) == ((int32)(0)))
			return res
		}

		i--
	}

	return false
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_inArrayBlock() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		if (this.Hx_Field_stack.Get(i) != ((int32)(3))) && (this.Hx_Field_stack.Get(i) != ((int32)(4))) {
			var res bool
			_ = res
			res = (this.Hx_Field_stack.Get(i) == ((int32)(1)))
			return res
		}

		i--
	}

	return false
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_isLabelStart(pos int32) bool {
	pos = this.VTable.Hx_Field_skipWhitespaceAndComments(pos)
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if pos < 0 {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = pos
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = -1
	} else {
		_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_0
	if !((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) {
		return false
	}

	var startPos int32
	_ = startPos
	startPos = pos
	pos++
	for {
		var tmp bool
		_ = tmp
		if pos < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if pos < 0 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = pos
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _hx_tmp_6
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_4
			tmp = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		pos++
	}

	for {
		var tmp bool
		_ = tmp
		if pos < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			var _hx_tmp_9 bool
			_ = _hx_tmp_9
			if pos < 0 {
				_hx_tmp_9 = true
			} else {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = pos
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_10
				_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_9 {
				_hx_tmp_8 = -1
			} else {
				_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
			}

			tmp = this.VTable.Hx_Field_isWhitespace(_hx_tmp_8)
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		pos++
	}

	var tmp bool
	_ = tmp
	if pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		var _hx_tmp_13 bool
		_ = _hx_tmp_13
		if pos < 0 {
			_hx_tmp_13 = true
		} else {
			var _hx_tmp_14 int32
			_ = _hx_tmp_14
			_hx_tmp_14 = pos
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = _hx_tmp_14
			_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_13 {
			_hx_tmp_12 = -1
		} else {
			_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		tmp = ((_hx_tmp_12) != 58)
	} else {
		tmp = true
	}

	if tmp {
		return false
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (pos - startPos), Valid: true}

	var _hx_tmp_16 HxArray[int32]
	_ = _hx_tmp_16
	_hx_tmp_16 = v1
	var _hx_tmp_17 int32
	_ = _hx_tmp_17
	_hx_tmp_17 = startPos
	var _hx_tmp_18 int32
	_ = _hx_tmp_18
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_19 int32
		_ = _hx_tmp_19
		_hx_tmp_19 = startPos
		var _hx_tmp_20 int32
		_ = _hx_tmp_20
		_hx_tmp_20 = _hx_tmp_19
		_hx_tmp_18 = (_hx_tmp_20 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_18 = HxMakeArrayView[any](v1).Len()
	}

	var word string
	_ = word
	word = Hx_Field_loreline_utf8vector_slice(_hx_tmp_16, _hx_tmp_17, _hx_tmp_18)
	if word == "case" {
		return false
	}

	return true
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_skipWhitespaceAndComments(pos int32) int32 {
	var startPos int32
	_ = startPos
	startPos = pos
	var foundContent bool
	_ = foundContent
	foundContent = false
	for pos < this.Hx_Field_length {
		for {
			var tmp bool
			_ = tmp
			if pos < this.Hx_Field_length {
				var tmp1 bool
				_ = tmp1
				var tmp2 bool
				_ = tmp2
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if pos < 0 {
					_hx_tmp_1 = true
				} else {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = pos
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _hx_tmp_2
					_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_1 {
					_hx_tmp_0 = -1
				} else {
					_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				if (_hx_tmp_0) != 32 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					var _hx_tmp_5 bool
					_ = _hx_tmp_5
					if pos < 0 {
						_hx_tmp_5 = true
					} else {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = pos
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = _hx_tmp_6
						_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_5 {
						_hx_tmp_4 = -1
					} else {
						_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					tmp2 = ((_hx_tmp_4) == 9)
				} else {
					tmp2 = true
				}

				if !tmp2 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					var _hx_tmp_9 bool
					_ = _hx_tmp_9
					if pos < 0 {
						_hx_tmp_9 = true
					} else {
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						_hx_tmp_10 = pos
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = _hx_tmp_10
						_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_9 {
						_hx_tmp_8 = -1
					} else {
						_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					tmp1 = ((_hx_tmp_8) == 10)
				} else {
					tmp1 = true
				}

				if !tmp1 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					var _hx_tmp_13 bool
					_ = _hx_tmp_13
					if pos < 0 {
						_hx_tmp_13 = true
					} else {
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = pos
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = _hx_tmp_14
						_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_13 {
						_hx_tmp_12 = -1
					} else {
						_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					tmp = ((_hx_tmp_12) == 13)
				} else {
					tmp = true
				}
			} else {
				tmp = false
			}

			if !tmp {
				break
			}

			pos++
			foundContent = true
		}

		if pos < (this.Hx_Field_length - 1) {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			var _hx_tmp_17 bool
			_ = _hx_tmp_17
			if pos < 0 {
				_hx_tmp_17 = true
			} else {
				var _hx_tmp_18 int32
				_ = _hx_tmp_18
				_hx_tmp_18 = pos
				var _hx_tmp_19 int32
				_ = _hx_tmp_19
				_hx_tmp_19 = _hx_tmp_18
				_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_17 {
				_hx_tmp_16 = -1
			} else {
				_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_16) == 47 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				var _hx_tmp_21 bool
				_ = _hx_tmp_21
				if pos1 < 0 {
					_hx_tmp_21 = true
				} else {
					var _hx_tmp_22 int32
					_ = _hx_tmp_22
					_hx_tmp_22 = pos1
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = _hx_tmp_22
					_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_21 {
					_hx_tmp_20 = -1
				} else {
					_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
				}

				if (_hx_tmp_20) == 47 {
					pos = startPos
					return pos
				} else {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos1_1 int32
					_ = tmp_pos1_1
					tmp_pos1_1 = (pos + 1)
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					var _hx_tmp_25 bool
					_ = _hx_tmp_25
					if tmp_pos1_1 < 0 {
						_hx_tmp_25 = true
					} else {
						var _hx_tmp_26 int32
						_ = _hx_tmp_26
						_hx_tmp_26 = tmp_pos1_1
						var _hx_tmp_27 int32
						_ = _hx_tmp_27
						_hx_tmp_27 = _hx_tmp_26
						_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_25 {
						_hx_tmp_24 = -1
					} else {
						_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos1_1))))
					}

					if (_hx_tmp_24) == 42 {
						pos = (pos + 2)
						foundContent = true
						var commentClosed bool
						_ = commentClosed
						commentClosed = false
						for pos < (this.Hx_Field_length - 1) {
							var tmp bool
							_ = tmp
							var tmp_v_3 HxArray[int32]
							_ = tmp_v_3
							tmp_v_3 = this.Hx_Field_input
							var _hx_tmp_28 int32
							_ = _hx_tmp_28
							var _hx_tmp_29 bool
							_ = _hx_tmp_29
							if pos < 0 {
								_hx_tmp_29 = true
							} else {
								var _hx_tmp_30 int32
								_ = _hx_tmp_30
								_hx_tmp_30 = pos
								var _hx_tmp_31 int32
								_ = _hx_tmp_31
								_hx_tmp_31 = _hx_tmp_30
								_hx_tmp_29 = (_hx_tmp_31 >= HxMakeArrayView[any](tmp_v_3).Len())
							}

							if _hx_tmp_29 {
								_hx_tmp_28 = -1
							} else {
								_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(pos))))
							}

							if (_hx_tmp_28) == 42 {
								var tmp_v_4 HxArray[int32]
								_ = tmp_v_4
								tmp_v_4 = this.Hx_Field_input
								var tmp_pos1_2 int32
								_ = tmp_pos1_2
								tmp_pos1_2 = (pos + 1)
								var _hx_tmp_32 int32
								_ = _hx_tmp_32
								var _hx_tmp_33 bool
								_ = _hx_tmp_33
								if tmp_pos1_2 < 0 {
									_hx_tmp_33 = true
								} else {
									var _hx_tmp_34 int32
									_ = _hx_tmp_34
									_hx_tmp_34 = tmp_pos1_2
									var _hx_tmp_35 int32
									_ = _hx_tmp_35
									_hx_tmp_35 = _hx_tmp_34
									_hx_tmp_33 = (_hx_tmp_35 >= HxMakeArrayView[any](tmp_v_4).Len())
								}

								if _hx_tmp_33 {
									_hx_tmp_32 = -1
								} else {
									_hx_tmp_32 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos1_2))))
								}

								tmp = ((_hx_tmp_32) == 47)
							} else {
								tmp = false
							}

							if tmp {
								pos = (pos + 2)
								commentClosed = true
								break
							}

							pos++
						}

						if !commentClosed {
							pos = startPos
							return pos
						}

						continue
					}
				}
			}
		}

		break
	}

	if foundContent {
		return pos
	} else {
		return startPos
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_stackPush(item int32) {
	_ = (Hx_Array_Push[int32](this.Hx_Field_stack, item))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_stackPop() HxNullable[int32] {
	return (Hx_Array_Pop[int32](this.Hx_Field_stack))
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field__add(c int32, incrementIndex bool) {
	if incrementIndex {
		this.Hx_Field_index++
	} else {
		this.Hx_Field_currentPosOffset++
	}

	if this.Hx_Field_inString {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if c == 34 {
			_hx_tmp_0 = c
		} else {
			_hx_tmp_0 = 32
		}

		this.Hx_Field_lineOutput.VTable.Hx_Field_addChar(_hx_tmp_0)
		this.Hx_Field_output.VTable.Hx_Field_addChar(c)
		_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
	} else {
		if (c == 13) || (c == 10) {
			if this.Hx_Field_inControlWithoutParens {
				this.Hx_Field_inControlWithoutParens = false
				var hx_cmp_3 HxNullable[int32]
				_ = hx_cmp_3
				hx_cmp_3 = (Hx_Array_Pop[int32](this.Hx_Field_stack))
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if (hx_cmp_3).Valid == false {
					_hx_tmp_1 = true
				} else {
					_hx_tmp_1 = ((hx_cmp_3).Value != ((int32)(6)))
				}

				if _hx_tmp_1 {
					this.VTable.Hx_Field_error("Unexpected end of line")
				}

				this.Hx_Field_currentPosOffset++
				this.Hx_Field_lineOutput.VTable.Hx_Field_addChar(41)
				this.Hx_Field_output.VTable.Hx_Field_addChar(41)
				_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
			}

			this.Hx_Field_inControl = false
			var line string
			_ = line
			line = this.Hx_Field_lineOutput.VTable.Hx_Field_toString()
			this.Hx_Field_lineOutput = Hx_Obj_stringbuf_CreateInstance()
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if this.VTable.Hx_Field_endsWithChar(line, 125) {
				var _hx_tmp_3 bool
				_ = _hx_tmp_3
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if this.VTable.Hx_Field_followsWithWord("else", this.Hx_Field_index) {
					_hx_tmp_4 = true
				} else {
					_hx_tmp_4 = this.VTable.Hx_Field_followsWithWord("catch", this.Hx_Field_index)
				}

				if _hx_tmp_4 {
					_hx_tmp_3 = true
				} else {
					var _hx_tmp_5 bool
					_ = _hx_tmp_5
					if this.Hx_Field_closedDoBrace {
						_hx_tmp_5 = this.VTable.Hx_Field_followsWithWord("while", this.Hx_Field_index)
					} else {
						_hx_tmp_5 = false
					}

					_hx_tmp_3 = _hx_tmp_5
				}

				_hx_tmp_2 = (_hx_tmp_3)
			} else {
				_hx_tmp_2 = false
			}

			var continuedAfterBrace bool
			_ = continuedAfterBrace
			continuedAfterBrace = _hx_tmp_2
			this.Hx_Field_closedDoBrace = false
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			var _hx_tmp_7 bool
			_ = _hx_tmp_7
			if this.VTable.Hx_Field_inStatementsBlock() {
				_hx_tmp_7 = true
			} else {
				_hx_tmp_7 = this.VTable.Hx_Field_inObjectBlock()
			}

			if _hx_tmp_7 {
				_hx_tmp_6 = true
			} else {
				_hx_tmp_6 = this.VTable.Hx_Field_inArrayBlock()
			}

			if _hx_tmp_6 {
				var indent int32
				_ = indent
				indent = this.VTable.Hx_Field_nextLineIndentOffset(line, this.Hx_Field_index)
				if int32(utf8.RuneCountInString(Hx_Field_stringtools_trim(line))) != 0 {
					var _hx_tmp_8 bool
					_ = _hx_tmp_8
					var _hx_tmp_9 bool
					_ = _hx_tmp_9
					if indent > 0 {
						var _hx_tmp_10 string
						_ = _hx_tmp_10
						_hx_tmp_10 = line
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = this.Hx_Field_index
						var _hx_tmp_12 string
						_ = _hx_tmp_12
						_hx_tmp_12 = _hx_tmp_10
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = _hx_tmp_11
						_hx_tmp_9 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_12, 123, _hx_tmp_13, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1150)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
					} else {
						_hx_tmp_9 = false
					}

					if _hx_tmp_9 {
						var _hx_tmp_14 string
						_ = _hx_tmp_14
						_hx_tmp_14 = line
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = this.Hx_Field_index
						var _hx_tmp_16 string
						_ = _hx_tmp_16
						_hx_tmp_16 = _hx_tmp_14
						var _hx_tmp_17 int32
						_ = _hx_tmp_17
						_hx_tmp_17 = _hx_tmp_15
						_hx_tmp_8 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_16, 91, _hx_tmp_17, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1150)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
					} else {
						_hx_tmp_8 = false
					}

					if _hx_tmp_8 {
						var l int32
						_ = l
						l = int32(utf8.RuneCountInString(line))
						var r int32
						_ = r
						r = 0
						for {
							var _hx_tmp_18 bool
							_ = _hx_tmp_18
							if r < l {
								_hx_tmp_18 = Hx_Field_stringtools_isSpace(line, r)
							} else {
								_hx_tmp_18 = false
							}

							if !(_hx_tmp_18) {
								break
							}

							r++
						}

						var _hx_tmp_19 string
						_ = _hx_tmp_19
						if r > 0 {
							var _hx_tmp_20 string
							_ = _hx_tmp_20
							_hx_tmp_20 = line
							var _hx_tmp_21 int32
							_ = _hx_tmp_21
							_hx_tmp_21 = r
							var _hx_tmp_22 string
							_ = _hx_tmp_22
							_hx_tmp_22 = _hx_tmp_20
							var _hx_tmp_23 int32
							_ = _hx_tmp_23
							_hx_tmp_23 = _hx_tmp_21
							_hx_tmp_19 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_22, _hx_tmp_23, HxNullable[int32]{Value: (l - r), Valid: true}))
						} else {
							_hx_tmp_19 = line
						}

						var trimmedLine string
						_ = trimmedLine
						trimmedLine = _hx_tmp_19
						var _hx_tmp_24 bool
						_ = _hx_tmp_24
						var _hx_tmp_25 bool
						_ = _hx_tmp_25
						if Hx_Field_stringtools_startsWith(trimmedLine, "case ") {
							_hx_tmp_25 = true
						} else {
							_hx_tmp_25 = Hx_Field_stringtools_startsWith(trimmedLine, "case\t")
						}

						if _hx_tmp_25 {
							_hx_tmp_24 = true
						} else {
							var _hx_tmp_26 bool
							_ = _hx_tmp_26
							var _hx_tmp_27 bool
							_ = _hx_tmp_27
							if trimmedLine == "default" {
								_hx_tmp_27 = true
							} else {
								_hx_tmp_27 = Hx_Field_stringtools_startsWith(trimmedLine, "default ")
							}

							if _hx_tmp_27 {
								_hx_tmp_26 = true
							} else {
								_hx_tmp_26 = Hx_Field_stringtools_startsWith(trimmedLine, "default\t")
							}

							_hx_tmp_24 = (_hx_tmp_26)
						}

						var isCaseLabel bool
						_ = isCaseLabel
						isCaseLabel = _hx_tmp_24
						if isCaseLabel {
							_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(4))))
							this.Hx_Field_indentLevel = (this.Hx_Field_indentLevel + indent)
							_ = (Hx_Array_Push[int32](this.Hx_Field_indentStack, this.Hx_Field_indentLevel))
							this.Hx_Field_currentPosOffset++
							this.Hx_Field_output.VTable.Hx_Field_addChar(58)
							_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
						} else {
							_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(3))))
							this.Hx_Field_indentLevel = (this.Hx_Field_indentLevel + indent)
							_ = (Hx_Array_Push[int32](this.Hx_Field_indentStack, this.Hx_Field_indentLevel))
							this.Hx_Field_currentPosOffset++
							this.Hx_Field_output.VTable.Hx_Field_addChar(32)
							_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
							this.Hx_Field_currentPosOffset++
							this.Hx_Field_output.VTable.Hx_Field_addChar(123)
							_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
						}
					} else {
						var _hx_tmp_28 bool
						_ = _hx_tmp_28
						var _hx_tmp_29 bool
						_ = _hx_tmp_29
						if indent < 0 {
							_hx_tmp_29 = (this.Hx_Field_stack.Len() > 0)
						} else {
							_hx_tmp_29 = false
						}

						if _hx_tmp_29 {
							var _hx_tmp_30 bool
							_ = _hx_tmp_30
							var _hx_tmp_31 HxArray[int32]
							_ = _hx_tmp_31
							_hx_tmp_31 = this.Hx_Field_stack
							if _hx_tmp_31.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(3)) {
								_hx_tmp_30 = true
							} else {
								var _hx_tmp_32 HxArray[int32]
								_ = _hx_tmp_32
								_hx_tmp_32 = this.Hx_Field_stack
								var _hx_tmp_33 HxArray[int32]
								_ = _hx_tmp_33
								_hx_tmp_33 = _hx_tmp_32
								_hx_tmp_30 = (_hx_tmp_33.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(4)))
							}

							_hx_tmp_28 = (_hx_tmp_30)
						} else {
							_hx_tmp_28 = false
						}

						if _hx_tmp_28 {
							var _hx_tmp_34 bool
							_ = _hx_tmp_34
							if !this.VTable.Hx_Field_inObjectBlock() {
								var _hx_tmp_35 string
								_ = _hx_tmp_35
								_hx_tmp_35 = line
								var _hx_tmp_36 int32
								_ = _hx_tmp_36
								_hx_tmp_36 = this.Hx_Field_index
								var _hx_tmp_37 string
								_ = _hx_tmp_37
								_hx_tmp_37 = _hx_tmp_35
								var _hx_tmp_38 int32
								_ = _hx_tmp_38
								_hx_tmp_38 = _hx_tmp_36
								_hx_tmp_34 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_37, 59, _hx_tmp_38, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1179)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
							} else {
								_hx_tmp_34 = false
							}

							if _hx_tmp_34 {
								this.Hx_Field_currentPosOffset++
								this.Hx_Field_output.VTable.Hx_Field_addChar(59)
								_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
							}

							this.Hx_Field_indentLevel = (this.Hx_Field_indentLevel + indent)
							var first bool
							_ = first
							first = true
							for {
								var _hx_tmp_39 bool
								_ = _hx_tmp_39
								var _hx_tmp_40 bool
								_ = _hx_tmp_40
								var _hx_tmp_41 bool
								_ = _hx_tmp_41
								if this.Hx_Field_indentStack.Len() > 0 {
									var _hx_tmp_42 HxArray[int32]
									_ = _hx_tmp_42
									_hx_tmp_42 = this.Hx_Field_indentStack
									var _hx_tmp_43 HxArray[int32]
									_ = _hx_tmp_43
									_hx_tmp_43 = _hx_tmp_42
									_hx_tmp_41 = (_hx_tmp_43.Get((this.Hx_Field_indentStack.Len() - 1)) > this.Hx_Field_indentLevel)
								} else {
									_hx_tmp_41 = false
								}

								if _hx_tmp_41 {
									_hx_tmp_40 = (this.Hx_Field_stack.Len() > 0)
								} else {
									_hx_tmp_40 = false
								}

								if _hx_tmp_40 {
									var _hx_tmp_44 bool
									_ = _hx_tmp_44
									var _hx_tmp_45 HxArray[int32]
									_ = _hx_tmp_45
									_hx_tmp_45 = this.Hx_Field_stack
									if _hx_tmp_45.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(3)) {
										_hx_tmp_44 = true
									} else {
										var _hx_tmp_46 HxArray[int32]
										_ = _hx_tmp_46
										_hx_tmp_46 = this.Hx_Field_stack
										var _hx_tmp_47 HxArray[int32]
										_ = _hx_tmp_47
										_hx_tmp_47 = _hx_tmp_46
										_hx_tmp_44 = (_hx_tmp_47.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(4)))
									}

									_hx_tmp_39 = (_hx_tmp_44)
								} else {
									_hx_tmp_39 = false
								}

								if !(_hx_tmp_39) {
									break
								}

								var popped HxNullable[int32]
								_ = popped
								popped = (Hx_Array_Pop[int32](this.Hx_Field_stack))
								_ = (Hx_Array_Pop[int32](this.Hx_Field_indentStack))
								var _hx_tmp_48 bool
								_ = _hx_tmp_48
								if (popped).Valid == true {
									_hx_tmp_48 = ((popped).Value == ((int32)(3)))
								} else {
									_hx_tmp_48 = false
								}

								if _hx_tmp_48 {
									if first {
										first = false
										this.Hx_Field_currentPosOffset++
										this.Hx_Field_output.VTable.Hx_Field_addChar(32)
										_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
									}

									this.Hx_Field_currentPosOffset++
									this.Hx_Field_output.VTable.Hx_Field_addChar(125)
									_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
								}
							}
						} else {
							var _hx_tmp_49 bool
							_ = _hx_tmp_49
							var _hx_tmp_50 bool
							_ = _hx_tmp_50
							var _hx_tmp_51 bool
							_ = _hx_tmp_51
							var _hx_tmp_52 bool
							_ = _hx_tmp_52
							if indent < 0 {
								_hx_tmp_52 = (this.Hx_Field_stack.Len() > 0)
							} else {
								_hx_tmp_52 = false
							}

							if _hx_tmp_52 {
								var _hx_tmp_53 HxArray[int32]
								_ = _hx_tmp_53
								_hx_tmp_53 = this.Hx_Field_stack
								var _hx_tmp_54 HxArray[int32]
								_ = _hx_tmp_54
								_hx_tmp_54 = _hx_tmp_53
								_hx_tmp_51 = (_hx_tmp_54.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(2)))
							} else {
								_hx_tmp_51 = false
							}

							if _hx_tmp_51 {
								_hx_tmp_50 = !continuedAfterBrace
							} else {
								_hx_tmp_50 = false
							}

							if _hx_tmp_50 {
								var _hx_tmp_55 string
								_ = _hx_tmp_55
								_hx_tmp_55 = line
								var _hx_tmp_56 int32
								_ = _hx_tmp_56
								_hx_tmp_56 = this.Hx_Field_index
								var _hx_tmp_57 string
								_ = _hx_tmp_57
								_hx_tmp_57 = _hx_tmp_55
								var _hx_tmp_58 int32
								_ = _hx_tmp_58
								_hx_tmp_58 = _hx_tmp_56
								_hx_tmp_49 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_57, 59, _hx_tmp_58, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1204)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
							} else {
								_hx_tmp_49 = false
							}

							if _hx_tmp_49 {
								this.Hx_Field_currentPosOffset++
								this.Hx_Field_output.VTable.Hx_Field_addChar(59)
								_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
							} else {
								var _hx_tmp_59 bool
								_ = _hx_tmp_59
								var _hx_tmp_60 bool
								_ = _hx_tmp_60
								if (indent == 0) && !continuedAfterBrace {
									var _hx_tmp_61 string
									_ = _hx_tmp_61
									_hx_tmp_61 = line
									var _hx_tmp_62 int32
									_ = _hx_tmp_62
									_hx_tmp_62 = this.Hx_Field_index
									var _hx_tmp_63 string
									_ = _hx_tmp_63
									_hx_tmp_63 = _hx_tmp_61
									var _hx_tmp_64 int32
									_ = _hx_tmp_64
									_hx_tmp_64 = _hx_tmp_62
									_hx_tmp_60 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_63, 59, _hx_tmp_64, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1210)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
								} else {
									_hx_tmp_60 = false
								}

								if _hx_tmp_60 {
									var _hx_tmp_65 string
									_ = _hx_tmp_65
									_hx_tmp_65 = line
									var _hx_tmp_66 int32
									_ = _hx_tmp_66
									_hx_tmp_66 = this.Hx_Field_index
									var _hx_tmp_67 string
									_ = _hx_tmp_67
									_hx_tmp_67 = _hx_tmp_65
									var _hx_tmp_68 int32
									_ = _hx_tmp_68
									_hx_tmp_68 = _hx_tmp_66
									_hx_tmp_59 = !this.VTable.Hx_Field_endsOrFollowsWithChar(_hx_tmp_67, 44, _hx_tmp_68, any(map[string]any{"fileName": ((any)("src/loreline/CodeToLorscript.hx")), "lineNumber": ((int32)(1210)), "className": ((any)("loreline.CodeToLorscript")), "methodName": ((any)("_add"))}))
								} else {
									_hx_tmp_59 = false
								}

								if _hx_tmp_59 {
									if this.VTable.Hx_Field_inObjectBlock() {
										this.Hx_Field_currentPosOffset++
										this.Hx_Field_output.VTable.Hx_Field_addChar(44)
										_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
									} else {
										if this.VTable.Hx_Field_inArrayBlock() {
											this.Hx_Field_currentPosOffset++
											this.Hx_Field_output.VTable.Hx_Field_addChar(44)
											_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
										} else {
											this.Hx_Field_currentPosOffset++
											this.Hx_Field_output.VTable.Hx_Field_addChar(59)
											_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
										}
									}
								}
							}
						}
					}
				}
			}

			this.Hx_Field_output.VTable.Hx_Field_addChar(c)
			_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
		} else {
			var _hx_tmp_69 bool
			_ = _hx_tmp_69
			var _hx_tmp_70 bool
			_ = _hx_tmp_70
			if !this.Hx_Field_inControl {
				_hx_tmp_70 = !this.VTable.Hx_Field_isWhitespace(c)
			} else {
				_hx_tmp_70 = false
			}

			if _hx_tmp_70 {
				_hx_tmp_69 = this.VTable.Hx_Field_endsWithControlKeyword(this.Hx_Field_lineOutput.VTable.Hx_Field_toString(), (this.Hx_Field_index - 1))
			} else {
				_hx_tmp_69 = false
			}

			if _hx_tmp_69 {
				this.Hx_Field_inControl = true
				if !this.VTable.Hx_Field_followsWithChar(40, (this.Hx_Field_index - 1)) {
					this.Hx_Field_inControlWithoutParens = true
					_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(6))))
					this.Hx_Field_currentPosOffset++
					this.Hx_Field_lineOutput.VTable.Hx_Field_addChar(c)
					this.Hx_Field_output.VTable.Hx_Field_addChar(40)
					_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
				} else {
					if c == 40 {
						_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(6))))
					}
				}

				this.Hx_Field_lineOutput.VTable.Hx_Field_addChar(c)
				this.Hx_Field_output.VTable.Hx_Field_addChar(c)
				_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
			} else {
				if c == 40 {
					_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(6))))
				} else {
					if c == 41 {
						var hx_cmp_4 HxNullable[int32]
						_ = hx_cmp_4
						hx_cmp_4 = (Hx_Array_Pop[int32](this.Hx_Field_stack))
						var _hx_tmp_71 bool
						_ = _hx_tmp_71
						if (hx_cmp_4).Valid == false {
							_hx_tmp_71 = true
						} else {
							_hx_tmp_71 = ((hx_cmp_4).Value != ((int32)(6)))
						}

						if _hx_tmp_71 {
							this.VTable.Hx_Field_error("Unexpected: )")
						}
					} else {
						if c == 91 {
							if !this.VTable.Hx_Field_endsWithArrayIndexable(this.Hx_Field_lineOutput.VTable.Hx_Field_toString()) {
								_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(1))))
							} else {
								_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(5))))
							}
						} else {
							if c == 93 {
								var popped HxNullable[int32]
								_ = popped
								popped = (Hx_Array_Pop[int32](this.Hx_Field_stack))
								var _hx_tmp_72 bool
								_ = _hx_tmp_72
								var _hx_tmp_73 bool
								_ = _hx_tmp_73
								if (popped).Valid == false {
									_hx_tmp_73 = true
								} else {
									_hx_tmp_73 = ((popped).Value != ((int32)(5)))
								}

								if _hx_tmp_73 {
									var _hx_tmp_74 bool
									_ = _hx_tmp_74
									if (popped).Valid == false {
										_hx_tmp_74 = true
									} else {
										_hx_tmp_74 = ((popped).Value != ((int32)(1)))
									}

									_hx_tmp_72 = (_hx_tmp_74)
								} else {
									_hx_tmp_72 = false
								}

								if _hx_tmp_72 {
									this.VTable.Hx_Field_error("Unexpected: ]")
								}
							} else {
								if c == 123 {
									if this.VTable.Hx_Field_isLabelStart(this.Hx_Field_index) {
										_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(0))))
									} else {
										_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(2))))
										if this.VTable.Hx_Field_endsWithWord(this.Hx_Field_lineOutput.VTable.Hx_Field_toString(), "do") {
											var _hx_tmp_75 HxArray[int32]
											_ = _hx_tmp_75
											_hx_tmp_75 = this.Hx_Field_doBraces
											_ = (Hx_Array_Push[int32](_hx_tmp_75, this.Hx_Field_stack.Len()))
										}
									}
								} else {
									if c == 125 {
										var s string
										_ = s
										s = this.Hx_Field_lineOutput.VTable.Hx_Field_toString()
										var l int32
										_ = l
										l = int32(utf8.RuneCountInString(s))
										var r int32
										_ = r
										r = 0
										for {
											var _hx_tmp_76 bool
											_ = _hx_tmp_76
											if r < l {
												_hx_tmp_76 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
											} else {
												_hx_tmp_76 = false
											}

											if !(_hx_tmp_76) {
												break
											}

											r++
										}

										var _hx_tmp_77 string
										_ = _hx_tmp_77
										if r > 0 {
											var _hx_tmp_78 string
											_ = _hx_tmp_78
											_hx_tmp_78 = s
											var _hx_tmp_79 string
											_ = _hx_tmp_79
											_hx_tmp_79 = _hx_tmp_78
											_hx_tmp_77 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_79, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
										} else {
											_hx_tmp_77 = s
										}

										var before string
										_ = before
										before = _hx_tmp_77
										var _hx_tmp_80 bool
										_ = _hx_tmp_80
										var _hx_tmp_81 bool
										_ = _hx_tmp_81
										if this.Hx_Field_stack.Len() > 0 {
											var _hx_tmp_82 HxArray[int32]
											_ = _hx_tmp_82
											_hx_tmp_82 = this.Hx_Field_stack
											var _hx_tmp_83 HxArray[int32]
											_ = _hx_tmp_83
											_hx_tmp_83 = _hx_tmp_82
											_hx_tmp_81 = (_hx_tmp_83.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(2)))
										} else {
											_hx_tmp_81 = false
										}

										if _hx_tmp_81 {
											_hx_tmp_80 = (int32(utf8.RuneCountInString(before)) > 0)
										} else {
											_hx_tmp_80 = false
										}

										if _hx_tmp_80 {
											var _hx_tmp_84 string
											_ = _hx_tmp_84
											_hx_tmp_84 = before
											var c HxNullable[int32]
											_ = c
											c = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_84, (int32(utf8.RuneCountInString(before)) - 1)))
											var _hx_tmp_85 int32
											_ = _hx_tmp_85
											if (c).Valid == false {
												_hx_tmp_85 = -1
											} else {
												_hx_tmp_85 = (c).Value
											}

											var last int32
											_ = last
											last = _hx_tmp_85
											if ((last != 59) && (last != 123)) && (last != 125) {
												this.Hx_Field_currentPosOffset++
												this.Hx_Field_output.VTable.Hx_Field_addChar(59)
												_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
											}
										}

										var _hx_tmp_86 bool
										_ = _hx_tmp_86
										if this.Hx_Field_doBraces.Len() > 0 {
											var _hx_tmp_88 HxArray[int32]
											_ = _hx_tmp_88
											_hx_tmp_88 = this.Hx_Field_doBraces
											var _hx_tmp_87 int32
											_ = _hx_tmp_87
											_hx_tmp_87 = _hx_tmp_88.Get((this.Hx_Field_doBraces.Len() - 1))
											var _hx_tmp_89 int32
											_ = _hx_tmp_89
											_hx_tmp_89 = _hx_tmp_87
											_hx_tmp_86 = (_hx_tmp_89 == this.Hx_Field_stack.Len())
										} else {
											_hx_tmp_86 = false
										}

										this.Hx_Field_closedDoBrace = _hx_tmp_86
										if this.Hx_Field_closedDoBrace {
											_ = (Hx_Array_Pop[int32](this.Hx_Field_doBraces))
										}

										var popped HxNullable[int32]
										_ = popped
										popped = (Hx_Array_Pop[int32](this.Hx_Field_stack))
										var _hx_tmp_90 bool
										_ = _hx_tmp_90
										var _hx_tmp_91 bool
										_ = _hx_tmp_91
										if (popped).Valid == false {
											_hx_tmp_91 = true
										} else {
											_hx_tmp_91 = ((popped).Value != ((int32)(2)))
										}

										if _hx_tmp_91 {
											var _hx_tmp_92 bool
											_ = _hx_tmp_92
											if (popped).Valid == false {
												_hx_tmp_92 = true
											} else {
												_hx_tmp_92 = ((popped).Value != ((int32)(0)))
											}

											_hx_tmp_90 = (_hx_tmp_92)
										} else {
											_hx_tmp_90 = false
										}

										if _hx_tmp_90 {
											this.VTable.Hx_Field_error("Unexpected: }")
										}
									}
								}
							}
						}
					}
				}

				this.Hx_Field_lineOutput.VTable.Hx_Field_addChar(c)
				this.Hx_Field_output.VTable.Hx_Field_addChar(c)
				_ = (Hx_Array_Push[int32](this.Hx_Field_posOffsets, this.Hx_Field_currentPosOffset))
			}
		}
	}
}

func (this *Hx_Obj_loreline_codetolorscript) Hx_Field_error(message string) {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_inputText
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_index
	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance(message, Hx_Field_loreline_position_fromContentAndIndex(_hx_tmp_0, _hx_tmp_1, nil))))))
}

var Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS HxArray[string]

func Hx_Init_Hx_Field_loreline_codetolorscript_CONTROL_KEYWORDS() HxArray[string] {
	return HxMakeArray[string]("for", "while", "if", "switch", "catch")
}

func init() {
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
