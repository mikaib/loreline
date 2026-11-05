package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_lexererror() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lexererror() {
}

type Hx_Obj_VTable_loreline_lexererror interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lexererror struct {
	Hx_Obj_loreline_error
	VTable Hx_Obj_VTable_loreline_lexererror `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_lexererror_CreateEmptyInstance() *Hx_Obj_loreline_lexererror {
	obj := &Hx_Obj_loreline_lexererror{}
	obj.Hx_Obj_loreline_error = *Hx_Obj_loreline_error_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_error.VTable = obj
	return obj
}

func Hx_Obj_loreline_lexererror_CreateInstance(message string, pos__isOptional any) *Hx_Obj_loreline_lexererror {
	obj := Hx_Obj_loreline_lexererror_CreateEmptyInstance()
	obj.Hx_New(message, pos__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_lexererror) Hx_New(message string, pos__isOptional any) {
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

var Hx_Obj_loreline_lexererror_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.LexerError",
	HxMakeArray[string](),
	HxMakeArray[string]("message", "pos", "stack", "toString"),
	Hx_Obj_loreline_error_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lexererror_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_lexererror_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lexererror_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lexererror) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lexererror_RTTI
}

func (this *Hx_Obj_loreline_lexererror) Hx_As_Hx_Obj_loreline_lexererror() *Hx_Obj_loreline_lexererror {
	return this
}

type Hx_Enum_loreline_lstringattachment interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lstringattachment()
}

type Hx_Enum_loreline_lstringattachment_Tag struct {
	Hx_Field_closing bool
	Hx_Field_start   int32
	Hx_Field_length  int32
}

func (this Hx_Enum_loreline_lstringattachment_Tag) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lstringattachment_Tag) M_Hx_Enum_loreline_lstringattachment() {}
func (this Hx_Enum_loreline_lstringattachment_Tag) Hx_Field_enumIndex() int32             { return 1 }
func (this Hx_Enum_loreline_lstringattachment_Tag) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lstringattachment_RTTI
}
func (this Hx_Enum_loreline_lstringattachment_Tag) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_closing), any(this.Hx_Field_start), any(this.Hx_Field_length))
}
func (this Hx_Enum_loreline_lstringattachment_Tag) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_closing)
	case 1:
		return any(this.Hx_Field_start)
	case 2:
		return any(this.Hx_Field_length)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lstringattachment_Interpolation struct {
	Hx_Field_braces bool
	Hx_Field_inTag  bool
	Hx_Field_expr   Hx_Typedef_loreline_tokens
	Hx_Field_start  int32
	Hx_Field_length int32
}

func (this Hx_Enum_loreline_lstringattachment_Interpolation) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lstringattachment_Interpolation) M_Hx_Enum_loreline_lstringattachment() {}
func (this Hx_Enum_loreline_lstringattachment_Interpolation) Hx_Field_enumIndex() int32             { return 0 }
func (this Hx_Enum_loreline_lstringattachment_Interpolation) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lstringattachment_RTTI
}
func (this Hx_Enum_loreline_lstringattachment_Interpolation) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_braces), any(this.Hx_Field_inTag), any(this.Hx_Field_expr), any(this.Hx_Field_start), any(this.Hx_Field_length))
}
func (this Hx_Enum_loreline_lstringattachment_Interpolation) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_braces)
	case 1:
		return any(this.Hx_Field_inTag)
	case 2:
		return any(this.Hx_Field_expr)
	case 3:
		return any(this.Hx_Field_start)
	case 4:
		return any(this.Hx_Field_length)
	default:
		return nil
	}
}

var Hx_Enum_loreline_lstringattachment_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.LStringAttachment",
	HxMakeArray[string]("Interpolation", "Tag"),
	HxMakeArray[int32](5, 3),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lstringattachment_Interpolation{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0)))), ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(1)))), params.Get_Dyn(2).(Hx_Typedef_loreline_tokens), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(3)))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(4))))}
		case 1:
			return Hx_Enum_loreline_lstringattachment_Tag{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0)))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(1)))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(2))))}
		default:
			return nil
		}
	},
	nil,
)

var Hx_Abstract_loreline_strictexprtype_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.StrictExprType",
	HxMakeArray[string]("Loose", "Strict", "Inherit"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__lexer_strictexprtype_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__lexer_strictexprtype_impl_() {
	Hx_Field_loreline__lexer_strictexprtype_impl__Loose = Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Loose()
	Hx_Field_loreline__lexer_strictexprtype_impl__Strict = Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Strict()
	Hx_Field_loreline__lexer_strictexprtype_impl__Inherit = Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Inherit()
}

type Hx_Obj_VTable_loreline__lexer_strictexprtype_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__lexer_strictexprtype_impl_ struct {
	VTable Hx_Obj_VTable_loreline__lexer_strictexprtype_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__lexer_strictexprtype_impl__CreateEmptyInstance() *Hx_Obj_loreline__lexer_strictexprtype_impl_ {
	obj := &Hx_Obj_loreline__lexer_strictexprtype_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__lexer_strictexprtype_impl__CreateInstance() *Hx_Obj_loreline__lexer_strictexprtype_impl_ {
	obj := Hx_Obj_loreline__lexer_strictexprtype_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__lexer_strictexprtype_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Lexer.StrictExprType_Impl_",
	HxMakeArray[string]("Loose", "Strict", "Inherit"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__lexer_strictexprtype_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__lexer_strictexprtype_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__lexer_strictexprtype_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Loose":
			return &Hx_Field_loreline__lexer_strictexprtype_impl__Loose
		case "Strict":
			return &Hx_Field_loreline__lexer_strictexprtype_impl__Strict
		case "Inherit":
			return &Hx_Field_loreline__lexer_strictexprtype_impl__Inherit
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__lexer_strictexprtype_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__lexer_strictexprtype_impl__RTTI
}

func (this *Hx_Obj_loreline__lexer_strictexprtype_impl_) Hx_As_Hx_Obj_loreline__lexer_strictexprtype_impl_() *Hx_Obj_loreline__lexer_strictexprtype_impl_ {
	return this
}

var Hx_Field_loreline__lexer_strictexprtype_impl__Loose int32

func Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Loose() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__lexer_strictexprtype_impl__Strict int32

func Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Strict() int32 {
	return ((int32)(1))
}

var Hx_Field_loreline__lexer_strictexprtype_impl__Inherit int32

func Hx_Init_Hx_Field_loreline__lexer_strictexprtype_impl__Inherit() int32 {
	return ((int32)(2))
}

var Hx_Abstract_loreline_tokenstacktype_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.TokenStackType",
	HxMakeArray[string]("ChoiceBrace", "ChoiceIndent", "WhenBrace", "WhenIndent", "StateBrace", "StateIndent", "CharacterBrace", "CharacterIndent", "BeatBrace", "BeatIndent", "AlternativeBrace", "AlternativeIndent", "Brace", "Indent", "Bracket", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__lexer_tokenstacktype_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__lexer_tokenstacktype_impl_() {
	Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__WhenBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__WhenBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__WhenIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__WhenIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__StateBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__StateBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__StateIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__StateIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__BeatBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__BeatBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__BeatIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__BeatIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeBrace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeBrace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeIndent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeIndent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__Brace = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Brace()
	Hx_Field_loreline__lexer_tokenstacktype_impl__Indent = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Indent()
	Hx_Field_loreline__lexer_tokenstacktype_impl__Bracket = Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Bracket()
}

type Hx_Obj_VTable_loreline__lexer_tokenstacktype_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__lexer_tokenstacktype_impl_ struct {
	VTable Hx_Obj_VTable_loreline__lexer_tokenstacktype_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__lexer_tokenstacktype_impl__CreateEmptyInstance() *Hx_Obj_loreline__lexer_tokenstacktype_impl_ {
	obj := &Hx_Obj_loreline__lexer_tokenstacktype_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__lexer_tokenstacktype_impl__CreateInstance() *Hx_Obj_loreline__lexer_tokenstacktype_impl_ {
	obj := Hx_Obj_loreline__lexer_tokenstacktype_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__lexer_tokenstacktype_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Lexer.TokenStackType_Impl_",
	HxMakeArray[string]("ChoiceBrace", "ChoiceIndent", "WhenBrace", "WhenIndent", "StateBrace", "StateIndent", "CharacterBrace", "CharacterIndent", "BeatBrace", "BeatIndent", "AlternativeBrace", "AlternativeIndent", "Brace", "Indent", "Bracket", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__lexer_tokenstacktype_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__lexer_tokenstacktype_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__lexer_tokenstacktype_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "ChoiceBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceBrace
		case "ChoiceIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceIndent
		case "WhenBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__WhenBrace
		case "WhenIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__WhenIndent
		case "StateBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__StateBrace
		case "StateIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__StateIndent
		case "CharacterBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterBrace
		case "CharacterIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterIndent
		case "BeatBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__BeatBrace
		case "BeatIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__BeatIndent
		case "AlternativeBrace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeBrace
		case "AlternativeIndent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeIndent
		case "Brace":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__Brace
		case "Indent":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__Indent
		case "Bracket":
			return &Hx_Field_loreline__lexer_tokenstacktype_impl__Bracket
		case "toString":
			return Hx_Field_loreline__lexer_tokenstacktype_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__lexer_tokenstacktype_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__lexer_tokenstacktype_impl__RTTI
}

func (this *Hx_Obj_loreline__lexer_tokenstacktype_impl_) Hx_As_Hx_Obj_loreline__lexer_tokenstacktype_impl_() *Hx_Obj_loreline__lexer_tokenstacktype_impl_ {
	return this
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceBrace() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__ChoiceIndent() int32 {
	return ((int32)(1))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__WhenBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__WhenBrace() int32 {
	return ((int32)(2))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__WhenIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__WhenIndent() int32 {
	return ((int32)(3))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__StateBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__StateBrace() int32 {
	return ((int32)(4))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__StateIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__StateIndent() int32 {
	return ((int32)(5))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterBrace() int32 {
	return ((int32)(6))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__CharacterIndent() int32 {
	return ((int32)(7))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__BeatBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__BeatBrace() int32 {
	return ((int32)(8))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__BeatIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__BeatIndent() int32 {
	return ((int32)(9))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeBrace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeBrace() int32 {
	return ((int32)(10))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeIndent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__AlternativeIndent() int32 {
	return ((int32)(11))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__Brace int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Brace() int32 {
	return ((int32)(12))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__Indent int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Indent() int32 {
	return ((int32)(13))
}

var Hx_Field_loreline__lexer_tokenstacktype_impl__Bracket int32

func Hx_Init_Hx_Field_loreline__lexer_tokenstacktype_impl__Bracket() int32 {
	return ((int32)(14))
}

func Hx_Field_loreline__lexer_tokenstacktype_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "ChoiceBrace"
		}

	case 1:
		{
			return "ChoiceIndent"
		}

	case 2:
		{
			return "WhenBrace"
		}

	case 3:
		{
			return "WhenIndent"
		}

	case 4:
		{
			return "StateBrace"
		}

	case 5:
		{
			return "StateIndent"
		}

	case 6:
		{
			return "CharacterBrace"
		}

	case 7:
		{
			return "CharacterIndent"
		}

	case 8:
		{
			return "BeatBrace"
		}

	case 9:
		{
			return "BeatIndent"
		}

	case 10:
		{
			return "AlternativeBrace"
		}

	case 11:
		{
			return "AlternativeIndent"
		}

	case 12:
		{
			return "Brace"
		}

	case 13:
		{
			return "Indent"
		}

	case 14:
		{
			return "Bracket"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

type Hx_Enum_loreline_tokentype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_tokentype()
}

type Hx_Enum_loreline_tokentype_Unindent struct{}

func (this Hx_Enum_loreline_tokentype_Unindent) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Unindent) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Unindent) Hx_Field_enumIndex() int32               { return 56 }
func (this Hx_Enum_loreline_tokentype_Unindent) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Unindent) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Unindent) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Separator struct{}

func (this Hx_Enum_loreline_tokentype_Separator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Separator) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Separator) Hx_Field_enumIndex() int32               { return 14 }
func (this Hx_Enum_loreline_tokentype_Separator) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Separator) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Separator) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_RParen struct{}

func (this Hx_Enum_loreline_tokentype_RParen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_RParen) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_RParen) Hx_Field_enumIndex() int32               { return 49 }
func (this Hx_Enum_loreline_tokentype_RParen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_RParen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_RParen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_RBracket struct{}

func (this Hx_Enum_loreline_tokentype_RBracket) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_RBracket) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_RBracket) Hx_Field_enumIndex() int32               { return 51 }
func (this Hx_Enum_loreline_tokentype_RBracket) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_RBracket) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_RBracket) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_RBrace struct{}

func (this Hx_Enum_loreline_tokentype_RBrace) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_RBrace) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_RBrace) Hx_Field_enumIndex() int32               { return 47 }
func (this Hx_Enum_loreline_tokentype_RBrace) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_RBrace) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_RBrace) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Question struct{}

func (this Hx_Enum_loreline_tokentype_Question) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Question) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Question) Hx_Field_enumIndex() int32               { return 43 }
func (this Hx_Enum_loreline_tokentype_Question) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Question) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Question) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpUnquotedAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) Hx_Field_enumIndex() int32     { return 26 }
func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpUnquotedAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpPlusAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpPlusAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpPlusAssign) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpPlusAssign) Hx_Field_enumIndex() int32     { return 22 }
func (this Hx_Enum_loreline_tokentype_OpPlusAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpPlusAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpPlusAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpPlus struct{}

func (this Hx_Enum_loreline_tokentype_OpPlus) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpPlus) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpPlus) Hx_Field_enumIndex() int32               { return 27 }
func (this Hx_Enum_loreline_tokentype_OpPlus) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpPlus) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpPlus) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpOr struct {
	Hx_Field_word bool
}

func (this Hx_Enum_loreline_tokentype_OpOr) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpOr) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpOr) Hx_Field_enumIndex() int32               { return 39 }
func (this Hx_Enum_loreline_tokentype_OpOr) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpOr) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_word))
}
func (this Hx_Enum_loreline_tokentype_OpOr) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_word)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpNotEquals struct {
	Hx_Field_word bool
}

func (this Hx_Enum_loreline_tokentype_OpNotEquals) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpNotEquals) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpNotEquals) Hx_Field_enumIndex() int32     { return 33 }
func (this Hx_Enum_loreline_tokentype_OpNotEquals) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpNotEquals) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_word))
}
func (this Hx_Enum_loreline_tokentype_OpNotEquals) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_word)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpNot struct {
	Hx_Field_word bool
}

func (this Hx_Enum_loreline_tokentype_OpNot) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpNot) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpNot) Hx_Field_enumIndex() int32               { return 40 }
func (this Hx_Enum_loreline_tokentype_OpNot) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpNot) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_word))
}
func (this Hx_Enum_loreline_tokentype_OpNot) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_word)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpMultiplyAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) Hx_Field_enumIndex() int32     { return 24 }
func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpMultiplyAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpMultiply struct{}

func (this Hx_Enum_loreline_tokentype_OpMultiply) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpMultiply) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpMultiply) Hx_Field_enumIndex() int32     { return 29 }
func (this Hx_Enum_loreline_tokentype_OpMultiply) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpMultiply) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpMultiply) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpModulo struct{}

func (this Hx_Enum_loreline_tokentype_OpModulo) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpModulo) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpModulo) Hx_Field_enumIndex() int32               { return 31 }
func (this Hx_Enum_loreline_tokentype_OpModulo) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpModulo) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpModulo) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpMinusAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpMinusAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpMinusAssign) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpMinusAssign) Hx_Field_enumIndex() int32     { return 23 }
func (this Hx_Enum_loreline_tokentype_OpMinusAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpMinusAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpMinusAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpMinus struct{}

func (this Hx_Enum_loreline_tokentype_OpMinus) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpMinus) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpMinus) Hx_Field_enumIndex() int32               { return 28 }
func (this Hx_Enum_loreline_tokentype_OpMinus) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpMinus) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpMinus) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpLessEq struct{}

func (this Hx_Enum_loreline_tokentype_OpLessEq) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpLessEq) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpLessEq) Hx_Field_enumIndex() int32               { return 37 }
func (this Hx_Enum_loreline_tokentype_OpLessEq) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpLessEq) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpLessEq) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpLess struct{}

func (this Hx_Enum_loreline_tokentype_OpLess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpLess) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpLess) Hx_Field_enumIndex() int32               { return 35 }
func (this Hx_Enum_loreline_tokentype_OpLess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpLess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpLess) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpGreaterEq struct{}

func (this Hx_Enum_loreline_tokentype_OpGreaterEq) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpGreaterEq) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpGreaterEq) Hx_Field_enumIndex() int32     { return 36 }
func (this Hx_Enum_loreline_tokentype_OpGreaterEq) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpGreaterEq) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpGreaterEq) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpGreater struct{}

func (this Hx_Enum_loreline_tokentype_OpGreater) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpGreater) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpGreater) Hx_Field_enumIndex() int32               { return 34 }
func (this Hx_Enum_loreline_tokentype_OpGreater) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpGreater) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpGreater) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpEquals struct {
	Hx_Field_word bool
}

func (this Hx_Enum_loreline_tokentype_OpEquals) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpEquals) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpEquals) Hx_Field_enumIndex() int32               { return 32 }
func (this Hx_Enum_loreline_tokentype_OpEquals) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpEquals) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_word))
}
func (this Hx_Enum_loreline_tokentype_OpEquals) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_word)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpDivideAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpDivideAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_OpDivideAssign) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_OpDivideAssign) Hx_Field_enumIndex() int32     { return 25 }
func (this Hx_Enum_loreline_tokentype_OpDivideAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpDivideAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpDivideAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpDivide struct{}

func (this Hx_Enum_loreline_tokentype_OpDivide) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpDivide) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpDivide) Hx_Field_enumIndex() int32               { return 30 }
func (this Hx_Enum_loreline_tokentype_OpDivide) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpDivide) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpDivide) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpAssign struct{}

func (this Hx_Enum_loreline_tokentype_OpAssign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpAssign) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpAssign) Hx_Field_enumIndex() int32               { return 21 }
func (this Hx_Enum_loreline_tokentype_OpAssign) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpAssign) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_OpAssign) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_OpAnd struct {
	Hx_Field_word bool
}

func (this Hx_Enum_loreline_tokentype_OpAnd) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_OpAnd) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_OpAnd) Hx_Field_enumIndex() int32               { return 38 }
func (this Hx_Enum_loreline_tokentype_OpAnd) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_OpAnd) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_word))
}
func (this Hx_Enum_loreline_tokentype_OpAnd) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_word)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LineBreak struct{}

func (this Hx_Enum_loreline_tokentype_LineBreak) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LineBreak) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LineBreak) Hx_Field_enumIndex() int32               { return 57 }
func (this Hx_Enum_loreline_tokentype_LineBreak) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LineBreak) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_LineBreak) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LString struct {
	Hx_Field_quotes      int32
	Hx_Field_s           string
	Hx_Field_attachments HxArray[Hx_Enum_loreline_lstringattachment]
}

func (this Hx_Enum_loreline_tokentype_LString) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LString) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LString) Hx_Field_enumIndex() int32               { return 16 }
func (this Hx_Enum_loreline_tokentype_LString) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LString) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_quotes), any(this.Hx_Field_s), any(this.Hx_Field_attachments))
}
func (this Hx_Enum_loreline_tokentype_LString) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_quotes)
	case 1:
		return any(this.Hx_Field_s)
	case 2:
		return any(this.Hx_Field_attachments)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LParen struct{}

func (this Hx_Enum_loreline_tokentype_LParen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LParen) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LParen) Hx_Field_enumIndex() int32               { return 48 }
func (this Hx_Enum_loreline_tokentype_LParen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LParen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_LParen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LNumber struct {
	Hx_Field_n float64
}

func (this Hx_Enum_loreline_tokentype_LNumber) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LNumber) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LNumber) Hx_Field_enumIndex() int32               { return 17 }
func (this Hx_Enum_loreline_tokentype_LNumber) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LNumber) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_n))
}
func (this Hx_Enum_loreline_tokentype_LNumber) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_n)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LNull struct{}

func (this Hx_Enum_loreline_tokentype_LNull) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LNull) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LNull) Hx_Field_enumIndex() int32               { return 19 }
func (this Hx_Enum_loreline_tokentype_LNull) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LNull) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_LNull) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LBracket struct{}

func (this Hx_Enum_loreline_tokentype_LBracket) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LBracket) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LBracket) Hx_Field_enumIndex() int32               { return 50 }
func (this Hx_Enum_loreline_tokentype_LBracket) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LBracket) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_LBracket) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LBrace struct{}

func (this Hx_Enum_loreline_tokentype_LBrace) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LBrace) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LBrace) Hx_Field_enumIndex() int32               { return 46 }
func (this Hx_Enum_loreline_tokentype_LBrace) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LBrace) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_LBrace) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_LBoolean struct {
	Hx_Field_b bool
}

func (this Hx_Enum_loreline_tokentype_LBoolean) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_LBoolean) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_LBoolean) Hx_Field_enumIndex() int32               { return 18 }
func (this Hx_Enum_loreline_tokentype_LBoolean) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_LBoolean) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_b))
}
func (this Hx_Enum_loreline_tokentype_LBoolean) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_b)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwWhen struct{}

func (this Hx_Enum_loreline_tokentype_KwWhen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwWhen) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwWhen) Hx_Field_enumIndex() int32               { return 5 }
func (this Hx_Enum_loreline_tokentype_KwWhen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwWhen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwWhen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwState struct{}

func (this Hx_Enum_loreline_tokentype_KwState) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwState) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwState) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_loreline_tokentype_KwState) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwState) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwState) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwShuffle struct{}

func (this Hx_Enum_loreline_tokentype_KwShuffle) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwShuffle) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwShuffle) Hx_Field_enumIndex() int32               { return 13 }
func (this Hx_Enum_loreline_tokentype_KwShuffle) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwShuffle) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwShuffle) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwSequence struct{}

func (this Hx_Enum_loreline_tokentype_KwSequence) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_KwSequence) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_KwSequence) Hx_Field_enumIndex() int32     { return 9 }
func (this Hx_Enum_loreline_tokentype_KwSequence) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwSequence) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwSequence) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwPick struct{}

func (this Hx_Enum_loreline_tokentype_KwPick) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwPick) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwPick) Hx_Field_enumIndex() int32               { return 12 }
func (this Hx_Enum_loreline_tokentype_KwPick) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwPick) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwPick) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwOnce struct{}

func (this Hx_Enum_loreline_tokentype_KwOnce) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwOnce) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwOnce) Hx_Field_enumIndex() int32               { return 11 }
func (this Hx_Enum_loreline_tokentype_KwOnce) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwOnce) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwOnce) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwNew struct{}

func (this Hx_Enum_loreline_tokentype_KwNew) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwNew) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwNew) Hx_Field_enumIndex() int32               { return 8 }
func (this Hx_Enum_loreline_tokentype_KwNew) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwNew) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwNew) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwImport struct{}

func (this Hx_Enum_loreline_tokentype_KwImport) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwImport) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwImport) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_loreline_tokentype_KwImport) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwImport) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwImport) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwIf struct{}

func (this Hx_Enum_loreline_tokentype_KwIf) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwIf) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwIf) Hx_Field_enumIndex() int32               { return 6 }
func (this Hx_Enum_loreline_tokentype_KwIf) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwIf) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwIf) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwElse struct{}

func (this Hx_Enum_loreline_tokentype_KwElse) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwElse) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwElse) Hx_Field_enumIndex() int32               { return 7 }
func (this Hx_Enum_loreline_tokentype_KwElse) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwElse) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwElse) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwCycle struct{}

func (this Hx_Enum_loreline_tokentype_KwCycle) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwCycle) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwCycle) Hx_Field_enumIndex() int32               { return 10 }
func (this Hx_Enum_loreline_tokentype_KwCycle) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwCycle) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwCycle) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwChoice struct{}

func (this Hx_Enum_loreline_tokentype_KwChoice) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwChoice) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwChoice) Hx_Field_enumIndex() int32               { return 4 }
func (this Hx_Enum_loreline_tokentype_KwChoice) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwChoice) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwChoice) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwCharacter struct{}

func (this Hx_Enum_loreline_tokentype_KwCharacter) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_KwCharacter) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_KwCharacter) Hx_Field_enumIndex() int32     { return 3 }
func (this Hx_Enum_loreline_tokentype_KwCharacter) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwCharacter) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwCharacter) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_KwBeat struct{}

func (this Hx_Enum_loreline_tokentype_KwBeat) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_KwBeat) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_KwBeat) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_loreline_tokentype_KwBeat) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_KwBeat) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_KwBeat) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Indent struct{}

func (this Hx_Enum_loreline_tokentype_Indent) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Indent) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Indent) Hx_Field_enumIndex() int32               { return 55 }
func (this Hx_Enum_loreline_tokentype_Indent) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Indent) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Indent) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Identifier struct {
	Hx_Field_name string
}

func (this Hx_Enum_loreline_tokentype_Identifier) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_Identifier) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_Identifier) Hx_Field_enumIndex() int32     { return 20 }
func (this Hx_Enum_loreline_tokentype_Identifier) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Identifier) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_tokentype_Identifier) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Function struct {
	Hx_Field_name     HxNullable[string]
	Hx_Field_args     HxArray[string]
	Hx_Field_code     string
	Hx_Field_external bool
}

func (this Hx_Enum_loreline_tokentype_Function) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Function) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Function) Hx_Field_enumIndex() int32               { return 15 }
func (this Hx_Enum_loreline_tokentype_Function) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Function) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_name), any(this.Hx_Field_args), any(this.Hx_Field_code), any(this.Hx_Field_external))
}
func (this Hx_Enum_loreline_tokentype_Function) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_name)
	case 1:
		return any(this.Hx_Field_args)
	case 2:
		return any(this.Hx_Field_code)
	case 3:
		return any(this.Hx_Field_external)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Eof struct{}

func (this Hx_Enum_loreline_tokentype_Eof) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Eof) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Eof) Hx_Field_enumIndex() int32               { return 58 }
func (this Hx_Enum_loreline_tokentype_Eof) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Eof) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Eof) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Dot struct{}

func (this Hx_Enum_loreline_tokentype_Dot) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Dot) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Dot) Hx_Field_enumIndex() int32               { return 45 }
func (this Hx_Enum_loreline_tokentype_Dot) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Dot) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Dot) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_CommentMultiLine struct {
	Hx_Field_content string
}

func (this Hx_Enum_loreline_tokentype_CommentMultiLine) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_CommentMultiLine) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_CommentMultiLine) Hx_Field_enumIndex() int32     { return 53 }
func (this Hx_Enum_loreline_tokentype_CommentMultiLine) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_CommentMultiLine) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_content))
}
func (this Hx_Enum_loreline_tokentype_CommentMultiLine) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_content)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_CommentLine struct {
	Hx_Field_content string
}

func (this Hx_Enum_loreline_tokentype_CommentLine) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_CommentLine) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_CommentLine) Hx_Field_enumIndex() int32     { return 52 }
func (this Hx_Enum_loreline_tokentype_CommentLine) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_CommentLine) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_content))
}
func (this Hx_Enum_loreline_tokentype_CommentLine) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_content)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_CommentHash struct {
	Hx_Field_content string
}

func (this Hx_Enum_loreline_tokentype_CommentHash) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_tokentype_CommentHash) M_Hx_Enum_loreline_tokentype() {}
func (this Hx_Enum_loreline_tokentype_CommentHash) Hx_Field_enumIndex() int32     { return 54 }
func (this Hx_Enum_loreline_tokentype_CommentHash) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_CommentHash) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_content))
}
func (this Hx_Enum_loreline_tokentype_CommentHash) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_content)
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Comma struct{}

func (this Hx_Enum_loreline_tokentype_Comma) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Comma) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Comma) Hx_Field_enumIndex() int32               { return 44 }
func (this Hx_Enum_loreline_tokentype_Comma) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Comma) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Comma) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Colon struct{}

func (this Hx_Enum_loreline_tokentype_Colon) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Colon) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Colon) Hx_Field_enumIndex() int32               { return 42 }
func (this Hx_Enum_loreline_tokentype_Colon) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Colon) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Colon) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_tokentype_Arrow struct{}

func (this Hx_Enum_loreline_tokentype_Arrow) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_tokentype_Arrow) M_Hx_Enum_loreline_tokentype()           {}
func (this Hx_Enum_loreline_tokentype_Arrow) Hx_Field_enumIndex() int32               { return 41 }
func (this Hx_Enum_loreline_tokentype_Arrow) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_tokentype_RTTI
}
func (this Hx_Enum_loreline_tokentype_Arrow) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_tokentype_Arrow) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_tokentype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.TokenType",
	HxMakeArray[string]("KwImport", "KwState", "KwBeat", "KwCharacter", "KwChoice", "KwWhen", "KwIf", "KwElse", "KwNew", "KwSequence", "KwCycle", "KwOnce", "KwPick", "KwShuffle", "Separator", "Function", "LString", "LNumber", "LBoolean", "LNull", "Identifier", "OpAssign", "OpPlusAssign", "OpMinusAssign", "OpMultiplyAssign", "OpDivideAssign", "OpUnquotedAssign", "OpPlus", "OpMinus", "OpMultiply", "OpDivide", "OpModulo", "OpEquals", "OpNotEquals", "OpGreater", "OpLess", "OpGreaterEq", "OpLessEq", "OpAnd", "OpOr", "OpNot", "Arrow", "Colon", "Question", "Comma", "Dot", "LBrace", "RBrace", "LParen", "RParen", "LBracket", "RBracket", "CommentLine", "CommentMultiLine", "CommentHash", "Indent", "Unindent", "LineBreak", "Eof"),
	HxMakeArray[int32](0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 3, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_tokentype_KwImport{}
		case 1:
			return Hx_Enum_loreline_tokentype_KwState{}
		case 2:
			return Hx_Enum_loreline_tokentype_KwBeat{}
		case 3:
			return Hx_Enum_loreline_tokentype_KwCharacter{}
		case 4:
			return Hx_Enum_loreline_tokentype_KwChoice{}
		case 5:
			return Hx_Enum_loreline_tokentype_KwWhen{}
		case 6:
			return Hx_Enum_loreline_tokentype_KwIf{}
		case 7:
			return Hx_Enum_loreline_tokentype_KwElse{}
		case 8:
			return Hx_Enum_loreline_tokentype_KwNew{}
		case 9:
			return Hx_Enum_loreline_tokentype_KwSequence{}
		case 10:
			return Hx_Enum_loreline_tokentype_KwCycle{}
		case 11:
			return Hx_Enum_loreline_tokentype_KwOnce{}
		case 12:
			return Hx_Enum_loreline_tokentype_KwPick{}
		case 13:
			return Hx_Enum_loreline_tokentype_KwShuffle{}
		case 14:
			return Hx_Enum_loreline_tokentype_Separator{}
		case 15:
			return Hx_Enum_loreline_tokentype_Function{HxNullable[string]{Value: ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0)))), Valid: params.Get_Dyn(0) != nil}, HxAnyToTypedArray[string](params.Get_Dyn(1)), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(2)))), ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(3))))}
		case 16:
			return Hx_Enum_loreline_tokentype_LString{params.Get_Dyn(0).(int32), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1)))), params.Get_Dyn(2).(HxArray[Hx_Enum_loreline_lstringattachment])}
		case 17:
			return Hx_Enum_loreline_tokentype_LNumber{((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(params.Get_Dyn(0))))}
		case 18:
			return Hx_Enum_loreline_tokentype_LBoolean{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 19:
			return Hx_Enum_loreline_tokentype_LNull{}
		case 20:
			return Hx_Enum_loreline_tokentype_Identifier{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 21:
			return Hx_Enum_loreline_tokentype_OpAssign{}
		case 22:
			return Hx_Enum_loreline_tokentype_OpPlusAssign{}
		case 23:
			return Hx_Enum_loreline_tokentype_OpMinusAssign{}
		case 24:
			return Hx_Enum_loreline_tokentype_OpMultiplyAssign{}
		case 25:
			return Hx_Enum_loreline_tokentype_OpDivideAssign{}
		case 26:
			return Hx_Enum_loreline_tokentype_OpUnquotedAssign{}
		case 27:
			return Hx_Enum_loreline_tokentype_OpPlus{}
		case 28:
			return Hx_Enum_loreline_tokentype_OpMinus{}
		case 29:
			return Hx_Enum_loreline_tokentype_OpMultiply{}
		case 30:
			return Hx_Enum_loreline_tokentype_OpDivide{}
		case 31:
			return Hx_Enum_loreline_tokentype_OpModulo{}
		case 32:
			return Hx_Enum_loreline_tokentype_OpEquals{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 33:
			return Hx_Enum_loreline_tokentype_OpNotEquals{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 34:
			return Hx_Enum_loreline_tokentype_OpGreater{}
		case 35:
			return Hx_Enum_loreline_tokentype_OpLess{}
		case 36:
			return Hx_Enum_loreline_tokentype_OpGreaterEq{}
		case 37:
			return Hx_Enum_loreline_tokentype_OpLessEq{}
		case 38:
			return Hx_Enum_loreline_tokentype_OpAnd{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 39:
			return Hx_Enum_loreline_tokentype_OpOr{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 40:
			return Hx_Enum_loreline_tokentype_OpNot{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0))))}
		case 41:
			return Hx_Enum_loreline_tokentype_Arrow{}
		case 42:
			return Hx_Enum_loreline_tokentype_Colon{}
		case 43:
			return Hx_Enum_loreline_tokentype_Question{}
		case 44:
			return Hx_Enum_loreline_tokentype_Comma{}
		case 45:
			return Hx_Enum_loreline_tokentype_Dot{}
		case 46:
			return Hx_Enum_loreline_tokentype_LBrace{}
		case 47:
			return Hx_Enum_loreline_tokentype_RBrace{}
		case 48:
			return Hx_Enum_loreline_tokentype_LParen{}
		case 49:
			return Hx_Enum_loreline_tokentype_RParen{}
		case 50:
			return Hx_Enum_loreline_tokentype_LBracket{}
		case 51:
			return Hx_Enum_loreline_tokentype_RBracket{}
		case 52:
			return Hx_Enum_loreline_tokentype_CommentLine{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 53:
			return Hx_Enum_loreline_tokentype_CommentMultiLine{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 54:
			return Hx_Enum_loreline_tokentype_CommentHash{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 55:
			return Hx_Enum_loreline_tokentype_Indent{}
		case 56:
			return Hx_Enum_loreline_tokentype_Unindent{}
		case 57:
			return Hx_Enum_loreline_tokentype_LineBreak{}
		case 58:
			return Hx_Enum_loreline_tokentype_Eof{}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_tokentypehelpers() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_tokentypehelpers() {
}

type Hx_Obj_VTable_loreline_tokentypehelpers interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_tokentypehelpers struct {
	VTable Hx_Obj_VTable_loreline_tokentypehelpers `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_tokentypehelpers_CreateEmptyInstance() *Hx_Obj_loreline_tokentypehelpers {
	obj := &Hx_Obj_loreline_tokentypehelpers{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_tokentypehelpers_CreateInstance() *Hx_Obj_loreline_tokentypehelpers {
	obj := Hx_Obj_loreline_tokentypehelpers_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_tokentypehelpers_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.TokenTypeHelpers",
	HxMakeArray[string]("equals", "isComment", "isAssignOp", "isIdentifier", "isBlockStart", "toCodeString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_tokentypehelpers_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_tokentypehelpers_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_tokentypehelpers_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "equals":
			return Hx_Field_loreline_tokentypehelpers_equals
		case "isComment":
			return Hx_Field_loreline_tokentypehelpers_isComment
		case "isAssignOp":
			return Hx_Field_loreline_tokentypehelpers_isAssignOp
		case "isIdentifier":
			return Hx_Field_loreline_tokentypehelpers_isIdentifier
		case "isBlockStart":
			return Hx_Field_loreline_tokentypehelpers_isBlockStart
		case "toCodeString":
			return Hx_Field_loreline_tokentypehelpers_toCodeString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_tokentypehelpers) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_tokentypehelpers_RTTI
}

func (this *Hx_Obj_loreline_tokentypehelpers) Hx_As_Hx_Obj_loreline_tokentypehelpers() *Hx_Obj_loreline_tokentypehelpers {
	return this
}

func Hx_Field_loreline_tokentypehelpers_equals(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
	switch a.Hx_Field_enumIndex() {
	case 1:
		{
			if b.Hx_Field_enumIndex() == 1 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_79 func(a any, b any) bool
					_ = _hx_closure_src_79
					_hx_closure_src_79 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_79(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 2:
		{
			if b.Hx_Field_enumIndex() == 2 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_80 func(a any, b any) bool
					_ = _hx_closure_src_80
					_hx_closure_src_80 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_80(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 3:
		{
			if b.Hx_Field_enumIndex() == 3 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_81 func(a any, b any) bool
					_ = _hx_closure_src_81
					_hx_closure_src_81 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_81(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 4:
		{
			if b.Hx_Field_enumIndex() == 4 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_82 func(a any, b any) bool
					_ = _hx_closure_src_82
					_hx_closure_src_82 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_82(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 5:
		{
			if b.Hx_Field_enumIndex() == 5 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_83 func(a any, b any) bool
					_ = _hx_closure_src_83
					_hx_closure_src_83 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_83(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 6:
		{
			if b.Hx_Field_enumIndex() == 6 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_84 func(a any, b any) bool
					_ = _hx_closure_src_84
					_hx_closure_src_84 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_84(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 7:
		{
			if b.Hx_Field_enumIndex() == 7 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_85 func(a any, b any) bool
					_ = _hx_closure_src_85
					_hx_closure_src_85 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_85(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 8:
		{
			if b.Hx_Field_enumIndex() == 8 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_86 func(a any, b any) bool
					_ = _hx_closure_src_86
					_hx_closure_src_86 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_86(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 9:
		{
			if b.Hx_Field_enumIndex() == 9 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_87 func(a any, b any) bool
					_ = _hx_closure_src_87
					_hx_closure_src_87 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_87(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 10:
		{
			if b.Hx_Field_enumIndex() == 10 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_88 func(a any, b any) bool
					_ = _hx_closure_src_88
					_hx_closure_src_88 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_88(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 11:
		{
			if b.Hx_Field_enumIndex() == 11 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_89 func(a any, b any) bool
					_ = _hx_closure_src_89
					_hx_closure_src_89 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_89(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 12:
		{
			if b.Hx_Field_enumIndex() == 12 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_90 func(a any, b any) bool
					_ = _hx_closure_src_90
					_hx_closure_src_90 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_90(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 13:
		{
			if b.Hx_Field_enumIndex() == 13 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_91 func(a any, b any) bool
					_ = _hx_closure_src_91
					_hx_closure_src_91 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_91(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 14:
		{
			if b.Hx_Field_enumIndex() == 14 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_92 func(a any, b any) bool
					_ = _hx_closure_src_92
					_hx_closure_src_92 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_92(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 16:
		{
			if b.Hx_Field_enumIndex() == 16 {
				var s2 int32
				_ = s2
				s2 = b.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
				var s1 int32
				_ = s1
				s1 = a.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
				return (s1 == s2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_93 func(a any, b any) bool
					_ = _hx_closure_src_93
					_hx_closure_src_93 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_93(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 17:
		{
			if b.Hx_Field_enumIndex() == 17 {
				var n2 float64
				_ = n2
				n2 = b.(Hx_Enum_loreline_tokentype_LNumber).Hx_Field_n
				var n1 float64
				_ = n1
				n1 = a.(Hx_Enum_loreline_tokentype_LNumber).Hx_Field_n
				return (n1 == n2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_94 func(a any, b any) bool
					_ = _hx_closure_src_94
					_hx_closure_src_94 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_94(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 18:
		{
			if b.Hx_Field_enumIndex() == 18 {
				var b2 bool
				_ = b2
				b2 = b.(Hx_Enum_loreline_tokentype_LBoolean).Hx_Field_b
				var b1 bool
				_ = b1
				b1 = a.(Hx_Enum_loreline_tokentype_LBoolean).Hx_Field_b
				return (b1 == b2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_95 func(a any, b any) bool
					_ = _hx_closure_src_95
					_hx_closure_src_95 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_95(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 19:
		{
			if b.Hx_Field_enumIndex() == 19 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_96 func(a any, b any) bool
					_ = _hx_closure_src_96
					_hx_closure_src_96 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_96(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 20:
		{
			if b.Hx_Field_enumIndex() == 20 {
				var n2 string
				_ = n2
				n2 = b.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
				var n1 string
				_ = n1
				n1 = a.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
				return (n1 == n2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_97 func(a any, b any) bool
					_ = _hx_closure_src_97
					_hx_closure_src_97 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_97(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 21:
		{
			if b.Hx_Field_enumIndex() == 21 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_98 func(a any, b any) bool
					_ = _hx_closure_src_98
					_hx_closure_src_98 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_98(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 22:
		{
			if b.Hx_Field_enumIndex() == 22 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_99 func(a any, b any) bool
					_ = _hx_closure_src_99
					_hx_closure_src_99 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_99(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 23:
		{
			if b.Hx_Field_enumIndex() == 23 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_100 func(a any, b any) bool
					_ = _hx_closure_src_100
					_hx_closure_src_100 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_100(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 24:
		{
			if b.Hx_Field_enumIndex() == 24 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_101 func(a any, b any) bool
					_ = _hx_closure_src_101
					_hx_closure_src_101 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_101(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 25:
		{
			if b.Hx_Field_enumIndex() == 25 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_102 func(a any, b any) bool
					_ = _hx_closure_src_102
					_hx_closure_src_102 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_102(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 26:
		{
			if b.Hx_Field_enumIndex() == 26 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_103 func(a any, b any) bool
					_ = _hx_closure_src_103
					_hx_closure_src_103 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_103(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 27:
		{
			if b.Hx_Field_enumIndex() == 27 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_104 func(a any, b any) bool
					_ = _hx_closure_src_104
					_hx_closure_src_104 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_104(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 28:
		{
			if b.Hx_Field_enumIndex() == 28 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_105 func(a any, b any) bool
					_ = _hx_closure_src_105
					_hx_closure_src_105 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_105(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 29:
		{
			if b.Hx_Field_enumIndex() == 29 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_106 func(a any, b any) bool
					_ = _hx_closure_src_106
					_hx_closure_src_106 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_106(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 30:
		{
			if b.Hx_Field_enumIndex() == 30 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_107 func(a any, b any) bool
					_ = _hx_closure_src_107
					_hx_closure_src_107 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_107(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 31:
		{
			if b.Hx_Field_enumIndex() == 31 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_108 func(a any, b any) bool
					_ = _hx_closure_src_108
					_hx_closure_src_108 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_108(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 32:
		{
			if b.Hx_Field_enumIndex() == 32 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_109 func(a any, b any) bool
					_ = _hx_closure_src_109
					_hx_closure_src_109 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_109(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 33:
		{
			if b.Hx_Field_enumIndex() == 33 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_110 func(a any, b any) bool
					_ = _hx_closure_src_110
					_hx_closure_src_110 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_110(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 34:
		{
			if b.Hx_Field_enumIndex() == 34 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_111 func(a any, b any) bool
					_ = _hx_closure_src_111
					_hx_closure_src_111 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_111(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 35:
		{
			if b.Hx_Field_enumIndex() == 35 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_112 func(a any, b any) bool
					_ = _hx_closure_src_112
					_hx_closure_src_112 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_112(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 36:
		{
			if b.Hx_Field_enumIndex() == 36 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_113 func(a any, b any) bool
					_ = _hx_closure_src_113
					_hx_closure_src_113 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_113(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 37:
		{
			if b.Hx_Field_enumIndex() == 37 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_114 func(a any, b any) bool
					_ = _hx_closure_src_114
					_hx_closure_src_114 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_114(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 38:
		{
			if b.Hx_Field_enumIndex() == 38 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_115 func(a any, b any) bool
					_ = _hx_closure_src_115
					_hx_closure_src_115 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_115(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 39:
		{
			if b.Hx_Field_enumIndex() == 39 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_116 func(a any, b any) bool
					_ = _hx_closure_src_116
					_hx_closure_src_116 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_116(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 40:
		{
			if b.Hx_Field_enumIndex() == 40 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_117 func(a any, b any) bool
					_ = _hx_closure_src_117
					_hx_closure_src_117 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_117(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 41:
		{
			if b.Hx_Field_enumIndex() == 41 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_118 func(a any, b any) bool
					_ = _hx_closure_src_118
					_hx_closure_src_118 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_118(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 42:
		{
			if b.Hx_Field_enumIndex() == 42 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_119 func(a any, b any) bool
					_ = _hx_closure_src_119
					_hx_closure_src_119 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_119(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 43:
		{
			if b.Hx_Field_enumIndex() == 43 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_120 func(a any, b any) bool
					_ = _hx_closure_src_120
					_hx_closure_src_120 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_120(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 45:
		{
			if b.Hx_Field_enumIndex() == 45 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_121 func(a any, b any) bool
					_ = _hx_closure_src_121
					_hx_closure_src_121 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_121(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 46:
		{
			if b.Hx_Field_enumIndex() == 46 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_122 func(a any, b any) bool
					_ = _hx_closure_src_122
					_hx_closure_src_122 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_122(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 47:
		{
			if b.Hx_Field_enumIndex() == 47 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_123 func(a any, b any) bool
					_ = _hx_closure_src_123
					_hx_closure_src_123 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_123(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 48:
		{
			if b.Hx_Field_enumIndex() == 48 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_124 func(a any, b any) bool
					_ = _hx_closure_src_124
					_hx_closure_src_124 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_124(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 49:
		{
			if b.Hx_Field_enumIndex() == 49 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_125 func(a any, b any) bool
					_ = _hx_closure_src_125
					_hx_closure_src_125 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_125(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 50:
		{
			if b.Hx_Field_enumIndex() == 50 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_126 func(a any, b any) bool
					_ = _hx_closure_src_126
					_hx_closure_src_126 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_126(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 51:
		{
			if b.Hx_Field_enumIndex() == 51 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_127 func(a any, b any) bool
					_ = _hx_closure_src_127
					_hx_closure_src_127 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_127(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 52:
		{
			if b.Hx_Field_enumIndex() == 52 {
				var c2 string
				_ = c2
				c2 = b.(Hx_Enum_loreline_tokentype_CommentLine).Hx_Field_content
				var c1 string
				_ = c1
				c1 = a.(Hx_Enum_loreline_tokentype_CommentLine).Hx_Field_content
				return (c1 == c2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_128 func(a any, b any) bool
					_ = _hx_closure_src_128
					_hx_closure_src_128 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_128(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 53:
		{
			if b.Hx_Field_enumIndex() == 53 {
				var c2 string
				_ = c2
				c2 = b.(Hx_Enum_loreline_tokentype_CommentMultiLine).Hx_Field_content
				var c1 string
				_ = c1
				c1 = a.(Hx_Enum_loreline_tokentype_CommentMultiLine).Hx_Field_content
				return (c1 == c2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_129 func(a any, b any) bool
					_ = _hx_closure_src_129
					_hx_closure_src_129 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_129(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 54:
		{
			if b.Hx_Field_enumIndex() == 54 {
				var c2 string
				_ = c2
				c2 = b.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
				var c1 string
				_ = c1
				c1 = a.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
				return (c1 == c2)
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_130 func(a any, b any) bool
					_ = _hx_closure_src_130
					_hx_closure_src_130 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_130(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 55:
		{
			if b.Hx_Field_enumIndex() == 55 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_131 func(a any, b any) bool
					_ = _hx_closure_src_131
					_hx_closure_src_131 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_131(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 56:
		{
			if b.Hx_Field_enumIndex() == 56 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_132 func(a any, b any) bool
					_ = _hx_closure_src_132
					_hx_closure_src_132 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_132(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 57:
		{
			if b.Hx_Field_enumIndex() == 57 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_133 func(a any, b any) bool
					_ = _hx_closure_src_133
					_hx_closure_src_133 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_133(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	case 58:
		{
			if b.Hx_Field_enumIndex() == 58 {
				return true
			} else {
				return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					var _hx_closure_src_134 func(a any, b any) bool
					_ = _hx_closure_src_134
					_hx_closure_src_134 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
						return _hx_closure_src_134(((any)(a)), ((any)(b)))
					}
				}()(a, b)
			}
		}

	default:
		{
			return func() func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
				var _hx_closure_src_135 func(a any, b any) bool
				_ = _hx_closure_src_135
				_hx_closure_src_135 = Hx_Field_type_enumEq
				return func(a Hx_Enum_loreline_tokentype, b Hx_Enum_loreline_tokentype) bool {
					return _hx_closure_src_135(((any)(a)), ((any)(b)))
				}
			}()(a, b)
		}
	}
}

func Hx_Field_loreline_tokentypehelpers_isComment(a Hx_Enum_loreline_tokentype) bool {
	switch a.Hx_Field_enumIndex() {
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

func Hx_Field_loreline_tokentypehelpers_isAssignOp(a Hx_Enum_loreline_tokentype) bool {
	switch a.Hx_Field_enumIndex() {
	case 21, 22, 23, 24, 25, 26:
		{
			return true
		}

	default:
		{
			return false
		}
	}
}

func Hx_Field_loreline_tokentypehelpers_isIdentifier(a Hx_Enum_loreline_tokentype) bool {
	return (a.Hx_Field_enumIndex() == 20)
}

func Hx_Field_loreline_tokentypehelpers_isBlockStart(a Hx_Enum_loreline_tokentype) bool {
	switch a.Hx_Field_enumIndex() {
	case 1, 2, 3, 4, 5, 6:
		{
			return true
		}

	case 9, 10, 11, 12, 13, 14:
		{
			return true
		}

	default:
		{
			return false
		}
	}
}

func Hx_Field_loreline_tokentypehelpers_toCodeString(a Hx_Enum_loreline_tokentype) string {
	switch a.Hx_Field_enumIndex() {
	case 0:
		{
			return "import"
		}

	case 1:
		{
			return "state"
		}

	case 2:
		{
			return "beat"
		}

	case 3:
		{
			return "character"
		}

	case 4:
		{
			return "choice"
		}

	case 5:
		{
			return "when"
		}

	case 6:
		{
			return "if"
		}

	case 7:
		{
			return "else"
		}

	case 8:
		{
			return "new"
		}

	case 9:
		{
			return "sequence"
		}

	case 10:
		{
			return "cycle"
		}

	case 11:
		{
			return "once"
		}

	case 12:
		{
			return "pick"
		}

	case 13:
		{
			return "shuffle"
		}

	case 14:
		{
			return "--"
		}

	case 15:
		{
			return "function"
		}

	case 16:
		{
			return "string"
		}

	case 17:
		{
			return "number"
		}

	case 18:
		{
			return "boolean"
		}

	case 19:
		{
			return "null"
		}

	case 20:
		{
			var name string
			_ = name
			name = a.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
			return "identifier"
		}

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
			word = a.(Hx_Enum_loreline_tokentype_OpEquals).Hx_Field_word
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
			word = a.(Hx_Enum_loreline_tokentype_OpNotEquals).Hx_Field_word
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
			var word bool
			_ = word
			word = a.(Hx_Enum_loreline_tokentype_OpAnd).Hx_Field_word
			if word {
				return "and"
			} else {
				return "&&"
			}
		}

	case 39:
		{
			var word bool
			_ = word
			word = a.(Hx_Enum_loreline_tokentype_OpOr).Hx_Field_word
			if word {
				return "or"
			} else {
				return "||"
			}
		}

	case 40:
		{
			var word bool
			_ = word
			word = a.(Hx_Enum_loreline_tokentype_OpNot).Hx_Field_word
			if word {
				return "not"
			} else {
				return "!"
			}
		}

	case 41:
		{
			return "->"
		}

	case 42:
		{
			return ":"
		}

	case 43:
		{
			return "?"
		}

	case 44:
		{
			return ","
		}

	case 45:
		{
			return "."
		}

	case 46:
		{
			return "{"
		}

	case 47:
		{
			return "}"
		}

	case 48:
		{
			return "("
		}

	case 49:
		{
			return ")"
		}

	case 50:
		{
			return "["
		}

	case 51:
		{
			return "]"
		}

	case 52:
		{
			var content string
			_ = content
			content = a.(Hx_Enum_loreline_tokentype_CommentLine).Hx_Field_content
			return "comment"
		}

	case 53:
		{
			var content string
			_ = content
			content = a.(Hx_Enum_loreline_tokentype_CommentMultiLine).Hx_Field_content
			return "multiline comment"
		}

	case 54:
		{
			var content string
			_ = content
			content = a.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
			return "hash comment"
		}

	case 55:
		{
			return "indent"
		}

	case 56:
		{
			return "unindent"
		}

	case 57:
		{
			return "line break"
		}

	case 58:
		{
			return "end of file"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

type Hx_Typedef_loreline_tokens = HxArray[*Hx_Obj_loreline_token]

func Hx_ClassInit_Hx_Obj_VTable_loreline_token() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_token() {
}

type Hx_Obj_VTable_loreline_token interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_token struct {
	VTable        Hx_Obj_VTable_loreline_token `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_type Hx_Enum_loreline_tokentype
	Hx_Field_pos  *Hx_Obj_loreline_position
}

func Hx_Obj_loreline_token_CreateEmptyInstance() *Hx_Obj_loreline_token {
	obj := &Hx_Obj_loreline_token{}
	obj.VTable = obj
	obj.Hx_Field_type = HxDefault[Hx_Enum_loreline_tokentype]()
	obj.Hx_Field_pos = HxDefault[*Hx_Obj_loreline_position]()
	return obj
}

func Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type Hx_Enum_loreline_tokentype, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	obj := Hx_Obj_loreline_token_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_type, pos)
	return obj
}

func (this *Hx_Obj_loreline_token) Hx_New(_hx_reserved_type Hx_Enum_loreline_tokentype, pos *Hx_Obj_loreline_position) {
	{
		this.Hx_Field_type = _hx_reserved_type
		this.Hx_Field_pos = pos
	}

}

var Hx_Obj_loreline_token_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Token",
	HxMakeArray[string](),
	HxMakeArray[string]("type", "pos", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_token_CreateInstance(HxConvert[Hx_Enum_loreline_tokentype](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_token_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_token_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_token) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_token_RTTI
}

func (this *Hx_Obj_loreline_token) Hx_As_Hx_Obj_loreline_token() *Hx_Obj_loreline_token {
	return this
}

func (this *Hx_Obj_loreline_token) Hx_Field_toString() string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("", Hx_Field_std_string(this.Hx_Field_type)), " at ")
	return Hx_Field_go_haxe_hxstring_concat(_hx_tmp_0, this.Hx_Field_pos.VTable.Hx_Field_toString())
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_lexer() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lexer() {
	Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars = Hx_Init_Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars()
	Hx_Field_loreline_lexer_KEYWORDS = Hx_Init_Hx_Field_loreline_lexer_KEYWORDS()
}

type Hx_Obj_VTable_loreline_lexer interface {
	Hx_Field_isWhenStart(pos int32) bool
	Hx_Field_conditionGroupEnd(pos int32) int32
	Hx_Field_isAfterComma() bool
	Hx_Field_isAfterLBracket() bool
	Hx_Field_isStrict() bool
	Hx_Field_makePosition() *Hx_Obj_loreline_position
	Hx_Field_reset()
	Hx_Field_isConditionUntilLineEnd(pos int32) bool
	Hx_Field_isIdentifierExpressionStart(pos int32, lowercaseIdentOnly bool) bool
	Hx_Field_isTransitionStart(pos int32) bool
	Hx_Field_isColon(pos int32, skipWhitespaces__isOptional any) bool
	Hx_Field_followsAssignStart() bool
	Hx_Field_advance(count__isOptional any)
	Hx_Field_isIfStart(pos int32) bool
	Hx_Field_followsUnquotedAssign() bool
	Hx_Field_followsOnlyLabelOrCommentsInLine() bool
	Hx_Field_skipBalancedParens(p int32, allowTags bool) int32
	Hx_Field_makePositionRelativeTo(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_position
	Hx_Field_isIdentifierStart(c int32) bool
	Hx_Field_isLowerCase(c int32) bool
	Hx_Field_isNumber(value string) bool
	Hx_Field_isIdentifierPart(c int32) bool
	Hx_Field_tokenize() Hx_Typedef_loreline_tokens
	Hx_Field_parentBlockType() Hx_Enum_loreline_tokentype
	Hx_Field_isWhitespace(c int32) bool
	Hx_Field_matchIdentifier(pos int32) HxNullable[string]
	Hx_Field_isAssignStart(pos int32, strict bool) bool
	Hx_Field_followsCallableOrIndexable() bool
	Hx_Field_scanStringEnd(startPos int32, allowTags bool) int32
	Hx_Field_isCallStart(pos int32) bool
	Hx_Field_isInsideBrackets() bool
	Hx_Field_readHashComment() *Hx_Obj_loreline_token
	Hx_Field_isDigit(c int32) bool
	Hx_Field_nextToken() *Hx_Obj_loreline_token
	Hx_Field_inChoiceRoot() bool
	Hx_Field_conditionOperandEnd(pos int32) int32
	Hx_Field_followsOnlyWhitespacesOrCommentsInLine() bool
	Hx_Field_tryReadUnquotedString() *Hx_Obj_loreline_token
	Hx_Field_readFieldAccessInterpolation(interpStart *Hx_Obj_loreline_position) Hx_Typedef_loreline_tokens
	Hx_Field_isLabelStart(pos int32) bool
	Hx_Field_isStartingMultilineDialogue(labelIdentifierTokenIndex int32) int32
	Hx_Field_readMultiLineComment() *Hx_Obj_loreline_token
	Hx_Field_skipQuotedString()
	Hx_Field_countIndentation() int32
	Hx_Field_conditionOperatorEnd(pos int32) int32
	Hx_Field_isInsertionStart(pos int32) bool
	Hx_Field_followsChoiceOncePrefix() bool
	Hx_Field_scanTagEnd(startPos int32) int32
	Hx_Field_wordEnd(p int32, word string) int32
	Hx_Field_makeToken(_hx_reserved_type Hx_Enum_loreline_tokentype, position__isOptional any) *Hx_Obj_loreline_token
	Hx_Field_skipWhitespace()
	Hx_Field_nextLineOpensBlock(p int32, indent int32) bool
	Hx_Field_isConditionEnd(pos int32) bool
	Hx_Field_skipInlineSpacesAndComments(pos int32) int32
	Hx_Field_isChoiceOnceStart(pos int32) bool
	Hx_Field_hasNonSpecialChar(str string) bool
	Hx_Field_makeLooseOrStrictAfterBracket()
	Hx_Field_readNumber() *Hx_Obj_loreline_token
	Hx_Field_readFunction(start *Hx_Obj_loreline_position) *Hx_Obj_loreline_token
	Hx_Field_readLineComment() *Hx_Obj_loreline_token
	Hx_Field_getErrors() HxArray[*Hx_Obj_loreline_lexererror]
	Hx_Field_readLineBreak() *Hx_Obj_loreline_token
	Hx_Field__skipWhitespaceAndCommentsStopNextLine(pos int32) int32
	Hx_Field_afterLabelIdentifierToken(inSameLine__isOptional any) int32
	Hx_Field_readImportValue(c int32, startPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_token
	Hx_Field_readString(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_token
	Hx_Field_readComplexInterpolation(interpStart *Hx_Obj_loreline_position) Hx_Typedef_loreline_tokens
	Hx_Field_mergeIsNot(tokenType Hx_Enum_loreline_tokentype) Hx_Enum_loreline_tokentype
	Hx_Field_skipSpacesAndTabs(p int32) int32
	Hx_Field_inStateOrCharacter() bool
	Hx_Field_makeStrictIfFollowingCallable()
	Hx_Field_nextTokenWithPosition(startPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_token
	Hx_Field_readIdentifierTokenInInterpolation(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_token
	Hx_Field_readIdentifier() *Hx_Obj_loreline_token
	Hx_Field_error(message string, fatal bool) *Hx_Obj_loreline_lexererror
	Hx_Field_inWhenRoot() bool
	Hx_Field_isContinuingMultilineText(pos int32, indent int32) bool
	Hx_Field_isTagStart(pos int32) bool
	Hx_Field_scanAccessorChain(p int32, allowTags bool) int32
	Hx_Field_inBeat() bool
	Hx_Field__skipWhitespaceAndComments(pos int32, stopNextLine__isOptional any) int32
	Hx_Field_peek(offset__isOptional any) int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lexer struct {
	VTable                         Hx_Obj_VTable_loreline_lexer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_input                 HxArray[int32]
	Hx_Field_pos                   int32
	Hx_Field_length                int32
	Hx_Field_line                  int32
	Hx_Field_column                int32
	Hx_Field_startLine             int32
	Hx_Field_startColumn           int32
	Hx_Field_previous              *Hx_Obj_loreline_token
	Hx_Field_tokenized             Hx_Typedef_loreline_tokens
	Hx_Field_stack                 HxArray[int32]
	Hx_Field_nextBlock             int32
	Hx_Field_strictExprs           HxArray[int32]
	Hx_Field_indentLevel           int32
	Hx_Field_indentStack           HxArray[int32]
	Hx_Field_indentTokens          Hx_Typedef_loreline_tokens
	Hx_Field_indentSize            int32
	Hx_Field_detectedIndentSize    int32
	Hx_Field_allowTabs             bool
	Hx_Field_errors                HxArray[*Hx_Obj_loreline_lexererror]
	Hx_Field_keepOrphanExpressions bool
}

func Hx_Obj_loreline_lexer_CreateEmptyInstance() *Hx_Obj_loreline_lexer {
	obj := &Hx_Obj_loreline_lexer{}
	obj.VTable = obj
	obj.Hx_Field_input = HxDefault[HxArray[int32]]()
	obj.Hx_Field_pos = HxDefault[int32]()
	obj.Hx_Field_length = HxDefault[int32]()
	obj.Hx_Field_line = HxDefault[int32]()
	obj.Hx_Field_column = HxDefault[int32]()
	obj.Hx_Field_startLine = HxDefault[int32]()
	obj.Hx_Field_startColumn = HxDefault[int32]()
	obj.Hx_Field_previous = HxDefault[*Hx_Obj_loreline_token]()
	obj.Hx_Field_tokenized = HxDefault[Hx_Typedef_loreline_tokens]()
	obj.Hx_Field_stack = HxDefault[HxArray[int32]]()
	obj.Hx_Field_nextBlock = HxDefault[int32]()
	obj.Hx_Field_strictExprs = HxDefault[HxArray[int32]]()
	obj.Hx_Field_indentLevel = HxDefault[int32]()
	obj.Hx_Field_indentStack = HxDefault[HxArray[int32]]()
	obj.Hx_Field_indentTokens = HxDefault[Hx_Typedef_loreline_tokens]()
	obj.Hx_Field_indentSize = HxDefault[int32]()
	obj.Hx_Field_detectedIndentSize = HxDefault[int32]()
	obj.Hx_Field_allowTabs = HxDefault[bool]()
	obj.Hx_Field_errors = HxDefault[HxArray[*Hx_Obj_loreline_lexererror]]()
	obj.Hx_Field_keepOrphanExpressions = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_lexer_CreateInstance(input string) *Hx_Obj_loreline_lexer {
	obj := Hx_Obj_loreline_lexer_CreateEmptyInstance()
	obj.Hx_New(input)
	return obj
}

func (this *Hx_Obj_loreline_lexer) Hx_New(input string) {
	{
		this.Hx_Field_indentLevel = 0
	}
	{
		this.Hx_Field_indentStack = HxMakeArray[int32]()
	}
	{
		this.Hx_Field_indentTokens = ((Hx_Typedef_loreline_tokens)(HxMakeArray[*Hx_Obj_loreline_token]()))
	}
	{
		this.Hx_Field_indentSize = 4
	}
	{
		this.Hx_Field_detectedIndentSize = 2
	}
	{
		this.Hx_Field_allowTabs = true
	}
	{
		this.Hx_Field_errors = nil
	}
	{
		this.Hx_Field_keepOrphanExpressions = false
	}
	{
		this.Hx_Field_input = Hx_Field_loreline_utf8_toCodes(input)
		this.Hx_Field_length = HxMakeArrayView[any](this.Hx_Field_input).Len()
		this.VTable.Hx_Field_reset()
	}

}

var Hx_Obj_loreline_lexer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Lexer",
	HxMakeArray[string]("hasNonSpecialChar_specialChars", "KEYWORDS", "compareAttachments"),
	HxMakeArray[string]("input", "pos", "length", "line", "column", "startLine", "startColumn", "previous", "tokenized", "stack", "nextBlock", "strictExprs", "indentLevel", "indentStack", "indentTokens", "indentSize", "detectedIndentSize", "allowTabs", "errors", "keepOrphanExpressions", "reset", "tokenize", "nextToken", "countIndentation", "parentBlockType", "inBeat", "inChoiceRoot", "inWhenRoot", "isWhenStart", "nextLineOpensBlock", "inStateOrCharacter", "readLineBreak", "isWhitespace", "matchIdentifier", "skipWhitespaceAndComments", "_skipWhitespaceAndComments", "_skipWhitespaceAndCommentsStopNextLine", "isIfStart", "isConditionUntilLineEnd", "isConditionEnd", "skipInlineSpacesAndComments", "conditionOperandEnd", "conditionGroupEnd", "conditionOperatorEnd", "isIdentifierExpressionStart", "isTransitionStart", "isInsertionStart", "isChoiceOnceStart", "isLabelStart", "isCallStart", "isAssignStart", "isColon", "afterLabelIdentifierToken", "isStartingMultilineDialogue", "isContinuingMultilineText", "isAfterComma", "isAfterLBracket", "isInsideBrackets", "followsAssignStart", "followsUnquotedAssign", "followsOnlyLabelOrCommentsInLine", "followsOnlyWhitespacesOrCommentsInLine", "followsChoiceOncePrefix", "hasNonSpecialChar", "isNumber", "makeLooseOrStrictAfterBracket", "makeStrictIfFollowingCallable", "followsCallableOrIndexable", "isStrict", "readImportValue", "isTagStart", "scanTagEnd", "scanStringEnd", "scanAccessorChain", "skipBalancedParens", "tryReadUnquotedString", "readString", "readComplexInterpolation", "nextTokenWithPosition", "readFieldAccessInterpolation", "readIdentifierTokenInInterpolation", "makePositionRelativeTo", "readLineComment", "readMultiLineComment", "readHashComment", "readNumber", "mergeIsNot", "wordEnd", "skipSpacesAndTabs", "readIdentifier", "readFunction", "skipQuotedString", "makePosition", "makeToken", "advance", "peek", "error", "getErrors", "skipWhitespace", "isDigit", "isIdentifierStart", "isLowerCase", "isIdentifierPart"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lexer_CreateInstance(HxConvert[string](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lexer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lexer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "hasNonSpecialChar_specialChars":
			return &Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars
		case "KEYWORDS":
			return &Hx_Field_loreline_lexer_KEYWORDS
		case "compareAttachments":
			return Hx_Field_loreline_lexer_compareAttachments
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lexer_RTTI
}

func (this *Hx_Obj_loreline_lexer) Hx_As_Hx_Obj_loreline_lexer() *Hx_Obj_loreline_lexer {
	return this
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_reset() {
	this.Hx_Field_pos = 0
	this.Hx_Field_line = 1
	this.Hx_Field_column = 1
	this.Hx_Field_startLine = 1
	this.Hx_Field_startColumn = 1
	this.Hx_Field_previous = nil
	this.Hx_Field_stack = HxMakeArray[int32]()
	this.Hx_Field_nextBlock = ((int32)(12))
	this.Hx_Field_tokenized = nil
	this.Hx_Field_strictExprs = HxMakeArray[int32]()
	this.Hx_Field_indentLevel = 0
	this.Hx_Field_indentStack = HxMakeArray[int32](0)
	this.Hx_Field_indentTokens = ((Hx_Typedef_loreline_tokens)(HxMakeArray[*Hx_Obj_loreline_token]()))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_tokenize() Hx_Typedef_loreline_tokens {
	var tokens HxArray[*Hx_Obj_loreline_token]
	_ = tokens
	tokens = HxMakeArray[*Hx_Obj_loreline_token]()
	this.Hx_Field_tokenized = ((Hx_Typedef_loreline_tokens)(tokens))
	for {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.VTable.Hx_Field_nextToken()
		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Eof{}) {
			if this.Hx_Field_indentStack.Len() > 1 {
				var count int32
				_ = count
				count = (this.Hx_Field_indentStack.Len() - 1)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = count
					for _g < _g1 {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_0
						var position *Hx_Obj_loreline_position
						_ = position
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var tmp_token_1 *Hx_Obj_loreline_token
						_ = tmp_token_1
						tmp_token_1 = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Unindent{}, position)
						this.Hx_Field_previous = tmp_token_1
						_ = (Hx_Array_Push[*Hx_Obj_loreline_token](tokens, tmp_token_1))
					}
				}
			}

			break
		}

		_ = (Hx_Array_Push[*Hx_Obj_loreline_token](tokens, token))
		switch token.Hx_Field_type.Hx_Field_enumIndex() {
		case 1:
			{
				this.Hx_Field_nextBlock = ((int32)(5))
			}

		case 2:
			{
				var prevIndex int32
				_ = prevIndex
				prevIndex = (tokens.Len() - 2)
			hx_label_0:
				for prevIndex >= 0 {
					switch tokens.Get(prevIndex).Hx_Field_type.Hx_Field_enumIndex() {
					case 52:
						{
							prevIndex--
						}

					case 53:
						{
							prevIndex--
						}

					default:
						{
							break hx_label_0
						}
					}
				}

				var isDynamicTargetKeyword bool
				_ = isDynamicTargetKeyword
				if prevIndex >= 0 {
					switch tokens.Get(prevIndex).Hx_Field_type.Hx_Field_enumIndex() {
					case 27, 41:
						{
							isDynamicTargetKeyword = true
						}

					default:
						{
							isDynamicTargetKeyword = false
						}
					}
				} else {
					isDynamicTargetKeyword = false
				}

				var lookahead int32
				_ = lookahead
				lookahead = this.Hx_Field_pos
				for lookahead < this.Hx_Field_length {
					var v HxArray[int32]
					_ = v
					v = this.Hx_Field_input
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					if lookahead < 0 {
						_hx_tmp_2 = true
					} else {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = lookahead
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _hx_tmp_3
						_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
					}

					if _hx_tmp_2 {
						_hx_tmp_1 = -1
					} else {
						_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(lookahead))))
					}

					var c int32
					_ = c
					c = _hx_tmp_1
					if (c == 32) || (c == 9) {
						lookahead++
					} else {
						break
					}
				}

				var isBeatCallStatement bool
				_ = isBeatCallStatement
				if lookahead < this.Hx_Field_length {
					var v HxArray[int32]
					_ = v
					v = this.Hx_Field_input
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					var _hx_tmp_6 bool
					_ = _hx_tmp_6
					if lookahead < 0 {
						_hx_tmp_6 = true
					} else {
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = lookahead
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						_hx_tmp_8 = _hx_tmp_7
						_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
					}

					if _hx_tmp_6 {
						_hx_tmp_5 = -1
					} else {
						_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(lookahead))))
					}

					isBeatCallStatement = ((_hx_tmp_5) == 40)
				} else {
					isBeatCallStatement = false
				}

				if !isDynamicTargetKeyword && !isBeatCallStatement {
					this.Hx_Field_nextBlock = ((int32)(9))
				}
			}

		case 3:
			{
				this.Hx_Field_nextBlock = ((int32)(7))
			}

		case 4:
			{
				this.Hx_Field_nextBlock = ((int32)(1))
			}

		case 5:
			{
				this.Hx_Field_nextBlock = ((int32)(3))
			}

		case 9, 10, 11, 12, 13, 14:
			{
				this.Hx_Field_nextBlock = ((int32)(11))
			}

		case 46:
			{
				var tmp int32
				_ = tmp
				switch this.Hx_Field_nextBlock {
				case 0, 1:
					{
						tmp = ((int32)(0))
					}

				case 2, 3:
					{
						tmp = ((int32)(2))
					}

				case 4, 5:
					{
						tmp = ((int32)(4))
					}

				case 6, 7:
					{
						tmp = ((int32)(6))
					}

				case 8, 9:
					{
						tmp = ((int32)(9))
					}

				case 10, 11:
					{
						tmp = ((int32)(10))
					}

				case 12, 13, 14:
					{
						tmp = ((int32)(12))
					}

				}

				_ = (Hx_Array_Push[int32](this.Hx_Field_stack, tmp))
				this.Hx_Field_nextBlock = ((int32)(12))
			}

		case 50:
			{
				_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(14))))
				this.Hx_Field_nextBlock = ((int32)(12))
			}

		case 47, 51, 56:
			{
				_ = (Hx_Array_Pop[int32](this.Hx_Field_stack))
				this.Hx_Field_nextBlock = ((int32)(12))
			}

		case 55:
			{
				var tmp int32
				_ = tmp
				switch this.Hx_Field_nextBlock {
				case 0, 1:
					{
						tmp = ((int32)(1))
					}

				case 2, 3:
					{
						tmp = ((int32)(3))
					}

				case 4, 5:
					{
						tmp = ((int32)(5))
					}

				case 6, 7:
					{
						tmp = ((int32)(7))
					}

				case 8, 9:
					{
						tmp = ((int32)(9))
					}

				case 10, 11:
					{
						tmp = ((int32)(11))
					}

				case 12, 13, 14:
					{
						tmp = ((int32)(13))
					}

				}

				_ = (Hx_Array_Push[int32](this.Hx_Field_stack, tmp))
				this.Hx_Field_nextBlock = ((int32)(12))
			}

		default:
			{
			}
		}
	}

	return ((Hx_Typedef_loreline_tokens)(tokens))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_nextToken() *Hx_Obj_loreline_token {
	if this.Hx_Field_indentTokens.Len() > 0 {
		return (Hx_Array_Shift[*Hx_Obj_loreline_token](this.Hx_Field_indentTokens).Value)
	}

	this.VTable.Hx_Field_skipWhitespace()
	if this.Hx_Field_pos >= this.Hx_Field_length {
		var position *Hx_Obj_loreline_position
		_ = position
		if position == nil {
			position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
		}

		position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
		var token *Hx_Obj_loreline_token
		_ = token
		token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Eof{}, position)
		this.Hx_Field_previous = token
		return token
	}

	this.Hx_Field_startLine = this.Hx_Field_line
	this.Hx_Field_startColumn = this.Hx_Field_column
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
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
	if (c == 10) || (c == 13) {
		var lineBreakToken *Hx_Obj_loreline_token
		_ = lineBreakToken
		lineBreakToken = this.VTable.Hx_Field_readLineBreak()
		var currentIndent int32
		_ = currentIndent
		currentIndent = this.VTable.Hx_Field_countIndentation()
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = currentIndent
		var _hx_tmp_5 HxArray[int32]
		_ = _hx_tmp_5
		_hx_tmp_5 = this.Hx_Field_indentStack
		if _hx_tmp_4 > _hx_tmp_5.Get((this.Hx_Field_indentStack.Len() - 1)) {
			if this.Hx_Field_indentStack.Len() == 1 {
				this.Hx_Field_detectedIndentSize = (currentIndent - this.Hx_Field_indentStack.Get(0))
			}

			_ = (Hx_Array_Push[int32](this.Hx_Field_indentStack, currentIndent))
			var tmp Hx_Typedef_loreline_tokens
			_ = tmp
			tmp = this.Hx_Field_indentTokens
			var position *Hx_Obj_loreline_position
			_ = position
			if position == nil {
				position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
			}

			position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
			var token *Hx_Obj_loreline_token
			_ = token
			token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Indent{}, position)
			this.Hx_Field_previous = token
			_ = (Hx_Array_Push[*Hx_Obj_loreline_token](tmp, token))
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = currentIndent
			var _hx_tmp_7 HxArray[int32]
			_ = _hx_tmp_7
			_hx_tmp_7 = this.Hx_Field_indentStack
			if _hx_tmp_6 < _hx_tmp_7.Get((this.Hx_Field_indentStack.Len() - 1)) {
				for {
					var _hx_tmp_8 bool
					_ = _hx_tmp_8
					if this.Hx_Field_indentStack.Len() > 0 {
						var _hx_tmp_9 int32
						_ = _hx_tmp_9
						_hx_tmp_9 = currentIndent
						var _hx_tmp_10 HxArray[int32]
						_ = _hx_tmp_10
						_hx_tmp_10 = this.Hx_Field_indentStack
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = _hx_tmp_9
						var _hx_tmp_12 HxArray[int32]
						_ = _hx_tmp_12
						_hx_tmp_12 = _hx_tmp_10
						_hx_tmp_8 = (_hx_tmp_11 < _hx_tmp_12.Get((this.Hx_Field_indentStack.Len() - 1)))
					} else {
						_hx_tmp_8 = false
					}

					if !(_hx_tmp_8) {
						break
					}

					_ = (Hx_Array_Pop[int32](this.Hx_Field_indentStack))
					var tmp Hx_Typedef_loreline_tokens
					_ = tmp
					tmp = this.Hx_Field_indentTokens
					var position *Hx_Obj_loreline_position
					_ = position
					if position == nil {
						position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
					}

					position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
					var token *Hx_Obj_loreline_token
					_ = token
					token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Unindent{}, position)
					this.Hx_Field_previous = token
					_ = (Hx_Array_Push[*Hx_Obj_loreline_token](tmp, token))
				}
			}
		}

		return lineBreakToken
	}

	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	var _hx_tmp_13 bool
	_ = _hx_tmp_13
	if this.Hx_Field_tokenized.Len() > 0 {
		var _hx_tmp_14 Hx_Typedef_loreline_tokens
		_ = _hx_tmp_14
		_hx_tmp_14 = this.Hx_Field_tokenized
		var _hx_tmp_15 Hx_Typedef_loreline_tokens
		_ = _hx_tmp_15
		_hx_tmp_15 = _hx_tmp_14
		_hx_tmp_13 = (_hx_tmp_15.Get((this.Hx_Field_tokenized.Len() - 1)).Hx_Field_type == Hx_Enum_loreline_tokentype_KwImport{})
	} else {
		_hx_tmp_13 = false
	}

	if _hx_tmp_13 {
		return this.VTable.Hx_Field_readImportValue(c, startPos)
	}

	switch c {
	case 34:
		{
			var tmp *Hx_Obj_loreline_token
			_ = tmp
			tmp = this.VTable.Hx_Field_tryReadUnquotedString()
			if tmp != nil {
				return tmp
			} else {
				return this.VTable.Hx_Field_readString(startPos)
			}
		}

	case 91:
		{
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					_hx_tmp_16 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_16
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					var _hx_tmp_18 bool
					_ = _hx_tmp_18
					if tmp_pos_1 < 0 {
						_hx_tmp_18 = true
					} else {
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = tmp_pos_1
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						_hx_tmp_20 = _hx_tmp_19
						_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_18 {
						_hx_tmp_17 = -1
					} else {
						_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_17) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

			this.VTable.Hx_Field_makeLooseOrStrictAfterBracket()
			var position *Hx_Obj_loreline_position
			_ = position
			position = startPos
			if position == nil {
				position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
			}

			position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
			var token *Hx_Obj_loreline_token
			_ = token
			token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LBracket{}, position)
			this.Hx_Field_previous = token
			return token
		}

	case 123:
		{
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					_hx_tmp_21 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_21
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_22 int32
					_ = _hx_tmp_22
					var _hx_tmp_23 bool
					_ = _hx_tmp_23
					if tmp_pos_1 < 0 {
						_hx_tmp_23 = true
					} else {
						var _hx_tmp_24 int32
						_ = _hx_tmp_24
						_hx_tmp_24 = tmp_pos_1
						var _hx_tmp_25 int32
						_ = _hx_tmp_25
						_hx_tmp_25 = _hx_tmp_24
						_hx_tmp_23 = (_hx_tmp_25 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_23 {
						_hx_tmp_22 = -1
					} else {
						_hx_tmp_22 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_22) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

			var position *Hx_Obj_loreline_position
			_ = position
			position = startPos
			if position == nil {
				position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
			}

			position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
			var token *Hx_Obj_loreline_token
			_ = token
			token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LBrace{}, position)
			this.Hx_Field_previous = token
			return token
		}

	case 125:
		{
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_26 int32
					_ = _hx_tmp_26
					_hx_tmp_26 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_26
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					var _hx_tmp_28 bool
					_ = _hx_tmp_28
					if tmp_pos_1 < 0 {
						_hx_tmp_28 = true
					} else {
						var _hx_tmp_29 int32
						_ = _hx_tmp_29
						_hx_tmp_29 = tmp_pos_1
						var _hx_tmp_30 int32
						_ = _hx_tmp_30
						_hx_tmp_30 = _hx_tmp_29
						_hx_tmp_28 = (_hx_tmp_30 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_28 {
						_hx_tmp_27 = -1
					} else {
						_hx_tmp_27 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_27) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

			var position *Hx_Obj_loreline_position
			_ = position
			position = startPos
			if position == nil {
				position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
			}

			position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
			var token *Hx_Obj_loreline_token
			_ = token
			token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_RBrace{}, position)
			this.Hx_Field_previous = token
			return token
		}

	default:
		{
			var tmp *Hx_Obj_loreline_token
			_ = tmp
			tmp = this.VTable.Hx_Field_tryReadUnquotedString()
			if tmp != nil {
				return tmp
			} else {
				switch c {
				case 33:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_31 int32
									_ = _hx_tmp_31
									var _hx_tmp_32 bool
									_ = _hx_tmp_32
									if tmp_pos_1 < 0 {
										_hx_tmp_32 = true
									} else {
										var _hx_tmp_33 int32
										_ = _hx_tmp_33
										_hx_tmp_33 = tmp_pos_1
										var _hx_tmp_34 int32
										_ = _hx_tmp_34
										_hx_tmp_34 = _hx_tmp_33
										_hx_tmp_32 = (_hx_tmp_34 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_32 {
										_hx_tmp_31 = -1
									} else {
										_hx_tmp_31 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_31
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_35 int32
											_ = _hx_tmp_35
											_hx_tmp_35 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_35
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_36 int32
											_ = _hx_tmp_36
											var _hx_tmp_37 bool
											_ = _hx_tmp_37
											if tmp_pos_1 < 0 {
												_hx_tmp_37 = true
											} else {
												var _hx_tmp_38 int32
												_ = _hx_tmp_38
												_hx_tmp_38 = tmp_pos_1
												var _hx_tmp_39 int32
												_ = _hx_tmp_39
												_hx_tmp_39 = _hx_tmp_38
												_hx_tmp_37 = (_hx_tmp_39 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_37 {
												_hx_tmp_36 = -1
											} else {
												_hx_tmp_36 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_36) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpNotEquals{false}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_40 int32
											_ = _hx_tmp_40
											_hx_tmp_40 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_40
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_41 int32
											_ = _hx_tmp_41
											var _hx_tmp_42 bool
											_ = _hx_tmp_42
											if tmp_pos_1 < 0 {
												_hx_tmp_42 = true
											} else {
												var _hx_tmp_43 int32
												_ = _hx_tmp_43
												_hx_tmp_43 = tmp_pos_1
												var _hx_tmp_44 int32
												_ = _hx_tmp_44
												_hx_tmp_44 = _hx_tmp_43
												_hx_tmp_42 = (_hx_tmp_44 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_42 {
												_hx_tmp_41 = -1
											} else {
												_hx_tmp_41 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_41) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpNot{false}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 35:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var nextC int32
								_ = nextC
								if (this.Hx_Field_pos + 1) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + 1)
									var _hx_tmp_45 int32
									_ = _hx_tmp_45
									var _hx_tmp_46 bool
									_ = _hx_tmp_46
									if tmp_pos_1 < 0 {
										_hx_tmp_46 = true
									} else {
										var _hx_tmp_47 int32
										_ = _hx_tmp_47
										_hx_tmp_47 = tmp_pos_1
										var _hx_tmp_48 int32
										_ = _hx_tmp_48
										_hx_tmp_48 = _hx_tmp_47
										_hx_tmp_46 = (_hx_tmp_48 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_46 {
										_hx_tmp_45 = -1
									} else {
										_hx_tmp_45 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									nextC = _hx_tmp_45
								} else {
									nextC = 0
								}

								if (((((nextC >= 97) && (nextC <= 122)) || ((nextC >= 65) && (nextC <= 90))) || (nextC == 95)) || ((nextC >= 48) && (nextC <= 57))) || (nextC == 45) {
									return this.VTable.Hx_Field_readHashComment()
								} else {
									this.VTable.Hx_Field_error("Unexpected character: #", true)
									return nil
								}
							}
						}
					}

				case 37:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								{
									var _g int32
									_ = _g
									_g = 0
									var _g1 int32
									_ = _g1
									_g1 = 1
									for _g < _g1 {
										var _hx_tmp_49 int32
										_ = _hx_tmp_49
										_hx_tmp_49 = _g
										_g = (_g + 1)
										var _hx_reserved__ int32
										_ = _hx_reserved__
										_hx_reserved__ = _hx_tmp_49
										if this.Hx_Field_pos >= this.Hx_Field_length {
											break
										}

										var tmp_v_1 HxArray[int32]
										_ = tmp_v_1
										tmp_v_1 = this.Hx_Field_input
										var tmp_pos_1 int32
										_ = tmp_pos_1
										tmp_pos_1 = this.Hx_Field_pos
										var _hx_tmp_50 int32
										_ = _hx_tmp_50
										var _hx_tmp_51 bool
										_ = _hx_tmp_51
										if tmp_pos_1 < 0 {
											_hx_tmp_51 = true
										} else {
											var _hx_tmp_52 int32
											_ = _hx_tmp_52
											_hx_tmp_52 = tmp_pos_1
											var _hx_tmp_53 int32
											_ = _hx_tmp_53
											_hx_tmp_53 = _hx_tmp_52
											_hx_tmp_51 = (_hx_tmp_53 >= HxMakeArrayView[any](tmp_v_1).Len())
										}

										if _hx_tmp_51 {
											_hx_tmp_50 = -1
										} else {
											_hx_tmp_50 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
										}

										if (_hx_tmp_50) == 10 {
											this.Hx_Field_line++
											this.Hx_Field_column = 1
										} else {
											this.Hx_Field_column++
										}

										this.Hx_Field_pos++
									}
								}

								var position *Hx_Obj_loreline_position
								_ = position
								position = startPos
								if position == nil {
									position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
								}

								position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
								var token *Hx_Obj_loreline_token
								_ = token
								token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpModulo{}, position)
								this.Hx_Field_previous = token
								return token
							}
						}
					}

				case 38:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
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
										_hx_tmp_55 = (_hx_tmp_57 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_55 {
										_hx_tmp_54 = -1
									} else {
										_hx_tmp_54 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_54
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 38 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_58 int32
											_ = _hx_tmp_58
											_hx_tmp_58 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_58
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_59 int32
											_ = _hx_tmp_59
											var _hx_tmp_60 bool
											_ = _hx_tmp_60
											if tmp_pos_1 < 0 {
												_hx_tmp_60 = true
											} else {
												var _hx_tmp_61 int32
												_ = _hx_tmp_61
												_hx_tmp_61 = tmp_pos_1
												var _hx_tmp_62 int32
												_ = _hx_tmp_62
												_hx_tmp_62 = _hx_tmp_61
												_hx_tmp_60 = (_hx_tmp_62 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_60 {
												_hx_tmp_59 = -1
											} else {
												_hx_tmp_59 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_59) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpAnd{false}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_63 int32
											_ = _hx_tmp_63
											_hx_tmp_63 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_63
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_64 int32
											_ = _hx_tmp_64
											var _hx_tmp_65 bool
											_ = _hx_tmp_65
											if tmp_pos_1 < 0 {
												_hx_tmp_65 = true
											} else {
												var _hx_tmp_66 int32
												_ = _hx_tmp_66
												_hx_tmp_66 = tmp_pos_1
												var _hx_tmp_67 int32
												_ = _hx_tmp_67
												_hx_tmp_67 = _hx_tmp_66
												_hx_tmp_65 = (_hx_tmp_67 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_65 {
												_hx_tmp_64 = -1
											} else {
												_hx_tmp_64 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_64) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									this.VTable.Hx_Field_error("Expected &", false)
									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpAnd{false}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 40:
					{
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_68 int32
								_ = _hx_tmp_68
								_hx_tmp_68 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_68
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_69 int32
								_ = _hx_tmp_69
								var _hx_tmp_70 bool
								_ = _hx_tmp_70
								if tmp_pos_1 < 0 {
									_hx_tmp_70 = true
								} else {
									var _hx_tmp_71 int32
									_ = _hx_tmp_71
									_hx_tmp_71 = tmp_pos_1
									var _hx_tmp_72 int32
									_ = _hx_tmp_72
									_hx_tmp_72 = _hx_tmp_71
									_hx_tmp_70 = (_hx_tmp_72 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_70 {
									_hx_tmp_69 = -1
								} else {
									_hx_tmp_69 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								if (_hx_tmp_69) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						this.VTable.Hx_Field_makeStrictIfFollowingCallable()
						var position *Hx_Obj_loreline_position
						_ = position
						position = startPos
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var token *Hx_Obj_loreline_token
						_ = token
						token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LParen{}, position)
						this.Hx_Field_previous = token
						return token
					}

				case 41:
					{
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_73 int32
								_ = _hx_tmp_73
								_hx_tmp_73 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_73
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
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

								if (_hx_tmp_74) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
						var position *Hx_Obj_loreline_position
						_ = position
						position = startPos
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var token *Hx_Obj_loreline_token
						_ = token
						token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_RParen{}, position)
						this.Hx_Field_previous = token
						return token
					}

				case 42:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_78 int32
									_ = _hx_tmp_78
									var _hx_tmp_79 bool
									_ = _hx_tmp_79
									if tmp_pos_1 < 0 {
										_hx_tmp_79 = true
									} else {
										var _hx_tmp_80 int32
										_ = _hx_tmp_80
										_hx_tmp_80 = tmp_pos_1
										var _hx_tmp_81 int32
										_ = _hx_tmp_81
										_hx_tmp_81 = _hx_tmp_80
										_hx_tmp_79 = (_hx_tmp_81 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_79 {
										_hx_tmp_78 = -1
									} else {
										_hx_tmp_78 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_78
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_82 int32
											_ = _hx_tmp_82
											_hx_tmp_82 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_82
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_83 int32
											_ = _hx_tmp_83
											var _hx_tmp_84 bool
											_ = _hx_tmp_84
											if tmp_pos_1 < 0 {
												_hx_tmp_84 = true
											} else {
												var _hx_tmp_85 int32
												_ = _hx_tmp_85
												_hx_tmp_85 = tmp_pos_1
												var _hx_tmp_86 int32
												_ = _hx_tmp_86
												_hx_tmp_86 = _hx_tmp_85
												_hx_tmp_84 = (_hx_tmp_86 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_84 {
												_hx_tmp_83 = -1
											} else {
												_hx_tmp_83 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_83) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpMultiplyAssign{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_87 int32
											_ = _hx_tmp_87
											_hx_tmp_87 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_87
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_88 int32
											_ = _hx_tmp_88
											var _hx_tmp_89 bool
											_ = _hx_tmp_89
											if tmp_pos_1 < 0 {
												_hx_tmp_89 = true
											} else {
												var _hx_tmp_90 int32
												_ = _hx_tmp_90
												_hx_tmp_90 = tmp_pos_1
												var _hx_tmp_91 int32
												_ = _hx_tmp_91
												_hx_tmp_91 = _hx_tmp_90
												_hx_tmp_89 = (_hx_tmp_91 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_89 {
												_hx_tmp_88 = -1
											} else {
												_hx_tmp_88 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_88) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpMultiply{}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 43:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_92 int32
									_ = _hx_tmp_92
									var _hx_tmp_93 bool
									_ = _hx_tmp_93
									if tmp_pos_1 < 0 {
										_hx_tmp_93 = true
									} else {
										var _hx_tmp_94 int32
										_ = _hx_tmp_94
										_hx_tmp_94 = tmp_pos_1
										var _hx_tmp_95 int32
										_ = _hx_tmp_95
										_hx_tmp_95 = _hx_tmp_94
										_hx_tmp_93 = (_hx_tmp_95 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_93 {
										_hx_tmp_92 = -1
									} else {
										_hx_tmp_92 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_92
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_96 int32
											_ = _hx_tmp_96
											_hx_tmp_96 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_96
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_97 int32
											_ = _hx_tmp_97
											var _hx_tmp_98 bool
											_ = _hx_tmp_98
											if tmp_pos_1 < 0 {
												_hx_tmp_98 = true
											} else {
												var _hx_tmp_99 int32
												_ = _hx_tmp_99
												_hx_tmp_99 = tmp_pos_1
												var _hx_tmp_100 int32
												_ = _hx_tmp_100
												_hx_tmp_100 = _hx_tmp_99
												_hx_tmp_98 = (_hx_tmp_100 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_98 {
												_hx_tmp_97 = -1
											} else {
												_hx_tmp_97 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_97) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpPlusAssign{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_101 int32
											_ = _hx_tmp_101
											_hx_tmp_101 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_101
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_102 int32
											_ = _hx_tmp_102
											var _hx_tmp_103 bool
											_ = _hx_tmp_103
											if tmp_pos_1 < 0 {
												_hx_tmp_103 = true
											} else {
												var _hx_tmp_104 int32
												_ = _hx_tmp_104
												_hx_tmp_104 = tmp_pos_1
												var _hx_tmp_105 int32
												_ = _hx_tmp_105
												_hx_tmp_105 = _hx_tmp_104
												_hx_tmp_103 = (_hx_tmp_105 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_103 {
												_hx_tmp_102 = -1
											} else {
												_hx_tmp_102 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_102) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpPlus{}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 44:
					{
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_106 int32
								_ = _hx_tmp_106
								_hx_tmp_106 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_106
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_107 int32
								_ = _hx_tmp_107
								var _hx_tmp_108 bool
								_ = _hx_tmp_108
								if tmp_pos_1 < 0 {
									_hx_tmp_108 = true
								} else {
									var _hx_tmp_109 int32
									_ = _hx_tmp_109
									_hx_tmp_109 = tmp_pos_1
									var _hx_tmp_110 int32
									_ = _hx_tmp_110
									_hx_tmp_110 = _hx_tmp_109
									_hx_tmp_108 = (_hx_tmp_110 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_108 {
									_hx_tmp_107 = -1
								} else {
									_hx_tmp_107 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								if (_hx_tmp_107) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						var position *Hx_Obj_loreline_position
						_ = position
						position = startPos
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var token *Hx_Obj_loreline_token
						_ = token
						token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Comma{}, position)
						this.Hx_Field_previous = token
						return token
					}

				case 45:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_111 int32
									_ = _hx_tmp_111
									var _hx_tmp_112 bool
									_ = _hx_tmp_112
									if tmp_pos_1 < 0 {
										_hx_tmp_112 = true
									} else {
										var _hx_tmp_113 int32
										_ = _hx_tmp_113
										_hx_tmp_113 = tmp_pos_1
										var _hx_tmp_114 int32
										_ = _hx_tmp_114
										_hx_tmp_114 = _hx_tmp_113
										_hx_tmp_112 = (_hx_tmp_114 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_112 {
										_hx_tmp_111 = -1
									} else {
										_hx_tmp_111 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_111
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 62 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_115 int32
											_ = _hx_tmp_115
											_hx_tmp_115 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_115
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_116 int32
											_ = _hx_tmp_116
											var _hx_tmp_117 bool
											_ = _hx_tmp_117
											if tmp_pos_1 < 0 {
												_hx_tmp_117 = true
											} else {
												var _hx_tmp_118 int32
												_ = _hx_tmp_118
												_hx_tmp_118 = tmp_pos_1
												var _hx_tmp_119 int32
												_ = _hx_tmp_119
												_hx_tmp_119 = _hx_tmp_118
												_hx_tmp_117 = (_hx_tmp_119 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_117 {
												_hx_tmp_116 = -1
											} else {
												_hx_tmp_116 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_116) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Arrow{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									var tmp_offset_1 int32
									_ = tmp_offset_1
									tmp_offset_1 = 1
									var tmp_tmp_2 int32
									_ = tmp_tmp_2
									if (this.Hx_Field_pos + tmp_offset_1) < this.Hx_Field_length {
										var tmp_v_1 HxArray[int32]
										_ = tmp_v_1
										tmp_v_1 = this.Hx_Field_input
										var tmp_pos_1 int32
										_ = tmp_pos_1
										tmp_pos_1 = (this.Hx_Field_pos + tmp_offset_1)
										var _hx_tmp_120 int32
										_ = _hx_tmp_120
										var _hx_tmp_121 bool
										_ = _hx_tmp_121
										if tmp_pos_1 < 0 {
											_hx_tmp_121 = true
										} else {
											var _hx_tmp_122 int32
											_ = _hx_tmp_122
											_hx_tmp_122 = tmp_pos_1
											var _hx_tmp_123 int32
											_ = _hx_tmp_123
											_hx_tmp_123 = _hx_tmp_122
											_hx_tmp_121 = (_hx_tmp_123 >= HxMakeArrayView[any](tmp_v_1).Len())
										}

										if _hx_tmp_121 {
											_hx_tmp_120 = -1
										} else {
											_hx_tmp_120 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
										}

										tmp_tmp_2 = _hx_tmp_120
									} else {
										tmp_tmp_2 = 0
									}

									var _hx_tmp_124 bool
									_ = _hx_tmp_124
									if tmp_tmp_2 == 45 {
										_hx_tmp_124 = (this.Hx_Field_strictExprs.Len() == 0)
									} else {
										_hx_tmp_124 = false
									}

									if _hx_tmp_124 {
										{
											var _g int32
											_ = _g
											_g = 0
											var _g1 int32
											_ = _g1
											_g1 = 2
											for _g < _g1 {
												var _hx_tmp_125 int32
												_ = _hx_tmp_125
												_hx_tmp_125 = _g
												_g = (_g + 1)
												var _hx_reserved__ int32
												_ = _hx_reserved__
												_hx_reserved__ = _hx_tmp_125
												if this.Hx_Field_pos >= this.Hx_Field_length {
													break
												}

												var tmp_v_1 HxArray[int32]
												_ = tmp_v_1
												tmp_v_1 = this.Hx_Field_input
												var tmp_pos_1 int32
												_ = tmp_pos_1
												tmp_pos_1 = this.Hx_Field_pos
												var _hx_tmp_126 int32
												_ = _hx_tmp_126
												var _hx_tmp_127 bool
												_ = _hx_tmp_127
												if tmp_pos_1 < 0 {
													_hx_tmp_127 = true
												} else {
													var _hx_tmp_128 int32
													_ = _hx_tmp_128
													_hx_tmp_128 = tmp_pos_1
													var _hx_tmp_129 int32
													_ = _hx_tmp_129
													_hx_tmp_129 = _hx_tmp_128
													_hx_tmp_127 = (_hx_tmp_129 >= HxMakeArrayView[any](tmp_v_1).Len())
												}

												if _hx_tmp_127 {
													_hx_tmp_126 = -1
												} else {
													_hx_tmp_126 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
												}

												if (_hx_tmp_126) == 10 {
													this.Hx_Field_line++
													this.Hx_Field_column = 1
												} else {
													this.Hx_Field_column++
												}

												this.Hx_Field_pos++
											}
										}

										var position *Hx_Obj_loreline_position
										_ = position
										position = startPos
										if position == nil {
											position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
										}

										position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
										var token *Hx_Obj_loreline_token
										_ = token
										token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Separator{}, position)
										this.Hx_Field_previous = token
										return token
									} else {
										var tmp_offset_2 int32
										_ = tmp_offset_2
										tmp_offset_2 = 1
										var tmp_tmp_3 int32
										_ = tmp_tmp_3
										if (this.Hx_Field_pos + tmp_offset_2) < this.Hx_Field_length {
											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = (this.Hx_Field_pos + tmp_offset_2)
											var _hx_tmp_130 int32
											_ = _hx_tmp_130
											var _hx_tmp_131 bool
											_ = _hx_tmp_131
											if tmp_pos_1 < 0 {
												_hx_tmp_131 = true
											} else {
												var _hx_tmp_132 int32
												_ = _hx_tmp_132
												_hx_tmp_132 = tmp_pos_1
												var _hx_tmp_133 int32
												_ = _hx_tmp_133
												_hx_tmp_133 = _hx_tmp_132
												_hx_tmp_131 = (_hx_tmp_133 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_131 {
												_hx_tmp_130 = -1
											} else {
												_hx_tmp_130 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											tmp_tmp_3 = _hx_tmp_130
										} else {
											tmp_tmp_3 = 0
										}

										if tmp_tmp_3 == 61 {
											{
												var _g int32
												_ = _g
												_g = 0
												var _g1 int32
												_ = _g1
												_g1 = 2
												for _g < _g1 {
													var _hx_tmp_134 int32
													_ = _hx_tmp_134
													_hx_tmp_134 = _g
													_g = (_g + 1)
													var _hx_reserved__ int32
													_ = _hx_reserved__
													_hx_reserved__ = _hx_tmp_134
													if this.Hx_Field_pos >= this.Hx_Field_length {
														break
													}

													var tmp_v_1 HxArray[int32]
													_ = tmp_v_1
													tmp_v_1 = this.Hx_Field_input
													var tmp_pos_1 int32
													_ = tmp_pos_1
													tmp_pos_1 = this.Hx_Field_pos
													var _hx_tmp_135 int32
													_ = _hx_tmp_135
													var _hx_tmp_136 bool
													_ = _hx_tmp_136
													if tmp_pos_1 < 0 {
														_hx_tmp_136 = true
													} else {
														var _hx_tmp_137 int32
														_ = _hx_tmp_137
														_hx_tmp_137 = tmp_pos_1
														var _hx_tmp_138 int32
														_ = _hx_tmp_138
														_hx_tmp_138 = _hx_tmp_137
														_hx_tmp_136 = (_hx_tmp_138 >= HxMakeArrayView[any](tmp_v_1).Len())
													}

													if _hx_tmp_136 {
														_hx_tmp_135 = -1
													} else {
														_hx_tmp_135 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
													}

													if (_hx_tmp_135) == 10 {
														this.Hx_Field_line++
														this.Hx_Field_column = 1
													} else {
														this.Hx_Field_column++
													}

													this.Hx_Field_pos++
												}
											}

											var position *Hx_Obj_loreline_position
											_ = position
											position = startPos
											if position == nil {
												position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
											}

											position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
											var token *Hx_Obj_loreline_token
											_ = token
											token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpMinusAssign{}, position)
											this.Hx_Field_previous = token
											return token
										} else {
											{
												var _g int32
												_ = _g
												_g = 0
												var _g1 int32
												_ = _g1
												_g1 = 1
												for _g < _g1 {
													var _hx_tmp_139 int32
													_ = _hx_tmp_139
													_hx_tmp_139 = _g
													_g = (_g + 1)
													var _hx_reserved__ int32
													_ = _hx_reserved__
													_hx_reserved__ = _hx_tmp_139
													if this.Hx_Field_pos >= this.Hx_Field_length {
														break
													}

													var tmp_v_1 HxArray[int32]
													_ = tmp_v_1
													tmp_v_1 = this.Hx_Field_input
													var tmp_pos_1 int32
													_ = tmp_pos_1
													tmp_pos_1 = this.Hx_Field_pos
													var _hx_tmp_140 int32
													_ = _hx_tmp_140
													var _hx_tmp_141 bool
													_ = _hx_tmp_141
													if tmp_pos_1 < 0 {
														_hx_tmp_141 = true
													} else {
														var _hx_tmp_142 int32
														_ = _hx_tmp_142
														_hx_tmp_142 = tmp_pos_1
														var _hx_tmp_143 int32
														_ = _hx_tmp_143
														_hx_tmp_143 = _hx_tmp_142
														_hx_tmp_141 = (_hx_tmp_143 >= HxMakeArrayView[any](tmp_v_1).Len())
													}

													if _hx_tmp_141 {
														_hx_tmp_140 = -1
													} else {
														_hx_tmp_140 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
													}

													if (_hx_tmp_140) == 10 {
														this.Hx_Field_line++
														this.Hx_Field_column = 1
													} else {
														this.Hx_Field_column++
													}

													this.Hx_Field_pos++
												}
											}

											var position *Hx_Obj_loreline_position
											_ = position
											position = startPos
											if position == nil {
												position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
											}

											position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
											var token *Hx_Obj_loreline_token
											_ = token
											token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpMinus{}, position)
											this.Hx_Field_previous = token
											return token
										}
									}
								}
							}
						}
					}

				case 46:
					{
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_144 int32
								_ = _hx_tmp_144
								_hx_tmp_144 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_144
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_145 int32
								_ = _hx_tmp_145
								var _hx_tmp_146 bool
								_ = _hx_tmp_146
								if tmp_pos_1 < 0 {
									_hx_tmp_146 = true
								} else {
									var _hx_tmp_147 int32
									_ = _hx_tmp_147
									_hx_tmp_147 = tmp_pos_1
									var _hx_tmp_148 int32
									_ = _hx_tmp_148
									_hx_tmp_148 = _hx_tmp_147
									_hx_tmp_146 = (_hx_tmp_148 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_146 {
									_hx_tmp_145 = -1
								} else {
									_hx_tmp_145 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								if (_hx_tmp_145) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						var position *Hx_Obj_loreline_position
						_ = position
						position = startPos
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var token *Hx_Obj_loreline_token
						_ = token
						token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Dot{}, position)
						this.Hx_Field_previous = token
						return token
					}

				case 47:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_149 int32
									_ = _hx_tmp_149
									var _hx_tmp_150 bool
									_ = _hx_tmp_150
									if tmp_pos_1 < 0 {
										_hx_tmp_150 = true
									} else {
										var _hx_tmp_151 int32
										_ = _hx_tmp_151
										_hx_tmp_151 = tmp_pos_1
										var _hx_tmp_152 int32
										_ = _hx_tmp_152
										_hx_tmp_152 = _hx_tmp_151
										_hx_tmp_150 = (_hx_tmp_152 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_150 {
										_hx_tmp_149 = -1
									} else {
										_hx_tmp_149 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_149
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_153 int32
											_ = _hx_tmp_153
											_hx_tmp_153 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_153
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_154 int32
											_ = _hx_tmp_154
											var _hx_tmp_155 bool
											_ = _hx_tmp_155
											if tmp_pos_1 < 0 {
												_hx_tmp_155 = true
											} else {
												var _hx_tmp_156 int32
												_ = _hx_tmp_156
												_hx_tmp_156 = tmp_pos_1
												var _hx_tmp_157 int32
												_ = _hx_tmp_157
												_hx_tmp_157 = _hx_tmp_156
												_hx_tmp_155 = (_hx_tmp_157 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_155 {
												_hx_tmp_154 = -1
											} else {
												_hx_tmp_154 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_154) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpDivideAssign{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									var tmp_offset_1 int32
									_ = tmp_offset_1
									tmp_offset_1 = 1
									var tmp_tmp_2 int32
									_ = tmp_tmp_2
									if (this.Hx_Field_pos + tmp_offset_1) < this.Hx_Field_length {
										var tmp_v_1 HxArray[int32]
										_ = tmp_v_1
										tmp_v_1 = this.Hx_Field_input
										var tmp_pos_1 int32
										_ = tmp_pos_1
										tmp_pos_1 = (this.Hx_Field_pos + tmp_offset_1)
										var _hx_tmp_158 int32
										_ = _hx_tmp_158
										var _hx_tmp_159 bool
										_ = _hx_tmp_159
										if tmp_pos_1 < 0 {
											_hx_tmp_159 = true
										} else {
											var _hx_tmp_160 int32
											_ = _hx_tmp_160
											_hx_tmp_160 = tmp_pos_1
											var _hx_tmp_161 int32
											_ = _hx_tmp_161
											_hx_tmp_161 = _hx_tmp_160
											_hx_tmp_159 = (_hx_tmp_161 >= HxMakeArrayView[any](tmp_v_1).Len())
										}

										if _hx_tmp_159 {
											_hx_tmp_158 = -1
										} else {
											_hx_tmp_158 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
										}

										tmp_tmp_2 = _hx_tmp_158
									} else {
										tmp_tmp_2 = 0
									}

									if tmp_tmp_2 == 47 {
										return this.VTable.Hx_Field_readLineComment()
									} else {
										var tmp_offset_2 int32
										_ = tmp_offset_2
										tmp_offset_2 = 1
										var tmp_tmp_3 int32
										_ = tmp_tmp_3
										if (this.Hx_Field_pos + tmp_offset_2) < this.Hx_Field_length {
											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = (this.Hx_Field_pos + tmp_offset_2)
											var _hx_tmp_162 int32
											_ = _hx_tmp_162
											var _hx_tmp_163 bool
											_ = _hx_tmp_163
											if tmp_pos_1 < 0 {
												_hx_tmp_163 = true
											} else {
												var _hx_tmp_164 int32
												_ = _hx_tmp_164
												_hx_tmp_164 = tmp_pos_1
												var _hx_tmp_165 int32
												_ = _hx_tmp_165
												_hx_tmp_165 = _hx_tmp_164
												_hx_tmp_163 = (_hx_tmp_165 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_163 {
												_hx_tmp_162 = -1
											} else {
												_hx_tmp_162 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											tmp_tmp_3 = _hx_tmp_162
										} else {
											tmp_tmp_3 = 0
										}

										if tmp_tmp_3 == 42 {
											return this.VTable.Hx_Field_readMultiLineComment()
										} else {
											{
												var _g int32
												_ = _g
												_g = 0
												var _g1 int32
												_ = _g1
												_g1 = 1
												for _g < _g1 {
													var _hx_tmp_166 int32
													_ = _hx_tmp_166
													_hx_tmp_166 = _g
													_g = (_g + 1)
													var _hx_reserved__ int32
													_ = _hx_reserved__
													_hx_reserved__ = _hx_tmp_166
													if this.Hx_Field_pos >= this.Hx_Field_length {
														break
													}

													var tmp_v_1 HxArray[int32]
													_ = tmp_v_1
													tmp_v_1 = this.Hx_Field_input
													var tmp_pos_1 int32
													_ = tmp_pos_1
													tmp_pos_1 = this.Hx_Field_pos
													var _hx_tmp_167 int32
													_ = _hx_tmp_167
													var _hx_tmp_168 bool
													_ = _hx_tmp_168
													if tmp_pos_1 < 0 {
														_hx_tmp_168 = true
													} else {
														var _hx_tmp_169 int32
														_ = _hx_tmp_169
														_hx_tmp_169 = tmp_pos_1
														var _hx_tmp_170 int32
														_ = _hx_tmp_170
														_hx_tmp_170 = _hx_tmp_169
														_hx_tmp_168 = (_hx_tmp_170 >= HxMakeArrayView[any](tmp_v_1).Len())
													}

													if _hx_tmp_168 {
														_hx_tmp_167 = -1
													} else {
														_hx_tmp_167 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
													}

													if (_hx_tmp_167) == 10 {
														this.Hx_Field_line++
														this.Hx_Field_column = 1
													} else {
														this.Hx_Field_column++
													}

													this.Hx_Field_pos++
												}
											}

											var position *Hx_Obj_loreline_position
											_ = position
											position = startPos
											if position == nil {
												position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
											}

											position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
											var token *Hx_Obj_loreline_token
											_ = token
											token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpDivide{}, position)
											this.Hx_Field_previous = token
											return token
										}
									}
								}
							}
						}
					}

				case 58:
					{
						var offset int32
						_ = offset
						offset = 1
						var tmp_tmp_1 int32
						_ = tmp_tmp_1
						if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = (this.Hx_Field_pos + offset)
							var _hx_tmp_171 int32
							_ = _hx_tmp_171
							var _hx_tmp_172 bool
							_ = _hx_tmp_172
							if tmp_pos_1 < 0 {
								_hx_tmp_172 = true
							} else {
								var _hx_tmp_173 int32
								_ = _hx_tmp_173
								_hx_tmp_173 = tmp_pos_1
								var _hx_tmp_174 int32
								_ = _hx_tmp_174
								_hx_tmp_174 = _hx_tmp_173
								_hx_tmp_172 = (_hx_tmp_174 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_172 {
								_hx_tmp_171 = -1
							} else {
								_hx_tmp_171 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							tmp_tmp_1 = _hx_tmp_171
						} else {
							tmp_tmp_1 = 0
						}

						if tmp_tmp_1 == 61 {
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 2
								for _g < _g1 {
									var _hx_tmp_175 int32
									_ = _hx_tmp_175
									_hx_tmp_175 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_175
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = this.Hx_Field_pos
									var _hx_tmp_176 int32
									_ = _hx_tmp_176
									var _hx_tmp_177 bool
									_ = _hx_tmp_177
									if tmp_pos_1 < 0 {
										_hx_tmp_177 = true
									} else {
										var _hx_tmp_178 int32
										_ = _hx_tmp_178
										_hx_tmp_178 = tmp_pos_1
										var _hx_tmp_179 int32
										_ = _hx_tmp_179
										_hx_tmp_179 = _hx_tmp_178
										_hx_tmp_177 = (_hx_tmp_179 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_177 {
										_hx_tmp_176 = -1
									} else {
										_hx_tmp_176 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									if (_hx_tmp_176) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							var position *Hx_Obj_loreline_position
							_ = position
							position = startPos
							if position == nil {
								position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
							}

							position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
							var token *Hx_Obj_loreline_token
							_ = token
							token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpUnquotedAssign{}, position)
							this.Hx_Field_previous = token
							return token
						} else {
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_180 int32
									_ = _hx_tmp_180
									_hx_tmp_180 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_180
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = this.Hx_Field_pos
									var _hx_tmp_181 int32
									_ = _hx_tmp_181
									var _hx_tmp_182 bool
									_ = _hx_tmp_182
									if tmp_pos_1 < 0 {
										_hx_tmp_182 = true
									} else {
										var _hx_tmp_183 int32
										_ = _hx_tmp_183
										_hx_tmp_183 = tmp_pos_1
										var _hx_tmp_184 int32
										_ = _hx_tmp_184
										_hx_tmp_184 = _hx_tmp_183
										_hx_tmp_182 = (_hx_tmp_184 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_182 {
										_hx_tmp_181 = -1
									} else {
										_hx_tmp_181 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									if (_hx_tmp_181) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							var position *Hx_Obj_loreline_position
							_ = position
							position = startPos
							if position == nil {
								position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
							}

							position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
							var token *Hx_Obj_loreline_token
							_ = token
							token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Colon{}, position)
							this.Hx_Field_previous = token
							return token
						}
					}

				case 60:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_185 int32
									_ = _hx_tmp_185
									var _hx_tmp_186 bool
									_ = _hx_tmp_186
									if tmp_pos_1 < 0 {
										_hx_tmp_186 = true
									} else {
										var _hx_tmp_187 int32
										_ = _hx_tmp_187
										_hx_tmp_187 = tmp_pos_1
										var _hx_tmp_188 int32
										_ = _hx_tmp_188
										_hx_tmp_188 = _hx_tmp_187
										_hx_tmp_186 = (_hx_tmp_188 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_186 {
										_hx_tmp_185 = -1
									} else {
										_hx_tmp_185 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_185
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_189 int32
											_ = _hx_tmp_189
											_hx_tmp_189 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_189
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_190 int32
											_ = _hx_tmp_190
											var _hx_tmp_191 bool
											_ = _hx_tmp_191
											if tmp_pos_1 < 0 {
												_hx_tmp_191 = true
											} else {
												var _hx_tmp_192 int32
												_ = _hx_tmp_192
												_hx_tmp_192 = tmp_pos_1
												var _hx_tmp_193 int32
												_ = _hx_tmp_193
												_hx_tmp_193 = _hx_tmp_192
												_hx_tmp_191 = (_hx_tmp_193 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_191 {
												_hx_tmp_190 = -1
											} else {
												_hx_tmp_190 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_190) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpLessEq{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_194 int32
											_ = _hx_tmp_194
											_hx_tmp_194 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_194
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_195 int32
											_ = _hx_tmp_195
											var _hx_tmp_196 bool
											_ = _hx_tmp_196
											if tmp_pos_1 < 0 {
												_hx_tmp_196 = true
											} else {
												var _hx_tmp_197 int32
												_ = _hx_tmp_197
												_hx_tmp_197 = tmp_pos_1
												var _hx_tmp_198 int32
												_ = _hx_tmp_198
												_hx_tmp_198 = _hx_tmp_197
												_hx_tmp_196 = (_hx_tmp_198 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_196 {
												_hx_tmp_195 = -1
											} else {
												_hx_tmp_195 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_195) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpLess{}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 61:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_199 int32
									_ = _hx_tmp_199
									var _hx_tmp_200 bool
									_ = _hx_tmp_200
									if tmp_pos_1 < 0 {
										_hx_tmp_200 = true
									} else {
										var _hx_tmp_201 int32
										_ = _hx_tmp_201
										_hx_tmp_201 = tmp_pos_1
										var _hx_tmp_202 int32
										_ = _hx_tmp_202
										_hx_tmp_202 = _hx_tmp_201
										_hx_tmp_200 = (_hx_tmp_202 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_200 {
										_hx_tmp_199 = -1
									} else {
										_hx_tmp_199 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_199
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_203 int32
											_ = _hx_tmp_203
											_hx_tmp_203 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_203
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_204 int32
											_ = _hx_tmp_204
											var _hx_tmp_205 bool
											_ = _hx_tmp_205
											if tmp_pos_1 < 0 {
												_hx_tmp_205 = true
											} else {
												var _hx_tmp_206 int32
												_ = _hx_tmp_206
												_hx_tmp_206 = tmp_pos_1
												var _hx_tmp_207 int32
												_ = _hx_tmp_207
												_hx_tmp_207 = _hx_tmp_206
												_hx_tmp_205 = (_hx_tmp_207 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_205 {
												_hx_tmp_204 = -1
											} else {
												_hx_tmp_204 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_204) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpEquals{false}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_208 int32
											_ = _hx_tmp_208
											_hx_tmp_208 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_208
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_209 int32
											_ = _hx_tmp_209
											var _hx_tmp_210 bool
											_ = _hx_tmp_210
											if tmp_pos_1 < 0 {
												_hx_tmp_210 = true
											} else {
												var _hx_tmp_211 int32
												_ = _hx_tmp_211
												_hx_tmp_211 = tmp_pos_1
												var _hx_tmp_212 int32
												_ = _hx_tmp_212
												_hx_tmp_212 = _hx_tmp_211
												_hx_tmp_210 = (_hx_tmp_212 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_210 {
												_hx_tmp_209 = -1
											} else {
												_hx_tmp_209 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_209) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpAssign{}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 62:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_213 int32
									_ = _hx_tmp_213
									var _hx_tmp_214 bool
									_ = _hx_tmp_214
									if tmp_pos_1 < 0 {
										_hx_tmp_214 = true
									} else {
										var _hx_tmp_215 int32
										_ = _hx_tmp_215
										_hx_tmp_215 = tmp_pos_1
										var _hx_tmp_216 int32
										_ = _hx_tmp_216
										_hx_tmp_216 = _hx_tmp_215
										_hx_tmp_214 = (_hx_tmp_216 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_214 {
										_hx_tmp_213 = -1
									} else {
										_hx_tmp_213 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_213
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 61 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_217 int32
											_ = _hx_tmp_217
											_hx_tmp_217 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_217
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_218 int32
											_ = _hx_tmp_218
											var _hx_tmp_219 bool
											_ = _hx_tmp_219
											if tmp_pos_1 < 0 {
												_hx_tmp_219 = true
											} else {
												var _hx_tmp_220 int32
												_ = _hx_tmp_220
												_hx_tmp_220 = tmp_pos_1
												var _hx_tmp_221 int32
												_ = _hx_tmp_221
												_hx_tmp_221 = _hx_tmp_220
												_hx_tmp_219 = (_hx_tmp_221 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_219 {
												_hx_tmp_218 = -1
											} else {
												_hx_tmp_218 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_218) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpGreaterEq{}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_222 int32
											_ = _hx_tmp_222
											_hx_tmp_222 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_222
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_223 int32
											_ = _hx_tmp_223
											var _hx_tmp_224 bool
											_ = _hx_tmp_224
											if tmp_pos_1 < 0 {
												_hx_tmp_224 = true
											} else {
												var _hx_tmp_225 int32
												_ = _hx_tmp_225
												_hx_tmp_225 = tmp_pos_1
												var _hx_tmp_226 int32
												_ = _hx_tmp_226
												_hx_tmp_226 = _hx_tmp_225
												_hx_tmp_224 = (_hx_tmp_226 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_224 {
												_hx_tmp_223 = -1
											} else {
												_hx_tmp_223 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_223) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpGreater{}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				case 63:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								this.Hx_Field_pos++
								var position *Hx_Obj_loreline_position
								_ = position
								if position == nil {
									position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
								}

								position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
								var token *Hx_Obj_loreline_token
								_ = token
								token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Question{}, position)
								this.Hx_Field_previous = token
								return token
							}
						}
					}

				case 93:
					{
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_227 int32
								_ = _hx_tmp_227
								_hx_tmp_227 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_227
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_228 int32
								_ = _hx_tmp_228
								var _hx_tmp_229 bool
								_ = _hx_tmp_229
								if tmp_pos_1 < 0 {
									_hx_tmp_229 = true
								} else {
									var _hx_tmp_230 int32
									_ = _hx_tmp_230
									_hx_tmp_230 = tmp_pos_1
									var _hx_tmp_231 int32
									_ = _hx_tmp_231
									_hx_tmp_231 = _hx_tmp_230
									_hx_tmp_229 = (_hx_tmp_231 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_229 {
									_hx_tmp_228 = -1
								} else {
									_hx_tmp_228 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								if (_hx_tmp_228) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
						var position *Hx_Obj_loreline_position
						_ = position
						position = startPos
						if position == nil {
							position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
						}

						position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
						var token *Hx_Obj_loreline_token
						_ = token
						token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_RBracket{}, position)
						this.Hx_Field_previous = token
						return token
					}

				case 124:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								var offset int32
								_ = offset
								offset = 1
								var tmp_tmp_1 int32
								_ = tmp_tmp_1
								if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = (this.Hx_Field_pos + offset)
									var _hx_tmp_232 int32
									_ = _hx_tmp_232
									var _hx_tmp_233 bool
									_ = _hx_tmp_233
									if tmp_pos_1 < 0 {
										_hx_tmp_233 = true
									} else {
										var _hx_tmp_234 int32
										_ = _hx_tmp_234
										_hx_tmp_234 = tmp_pos_1
										var _hx_tmp_235 int32
										_ = _hx_tmp_235
										_hx_tmp_235 = _hx_tmp_234
										_hx_tmp_233 = (_hx_tmp_235 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_233 {
										_hx_tmp_232 = -1
									} else {
										_hx_tmp_232 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									tmp_tmp_1 = _hx_tmp_232
								} else {
									tmp_tmp_1 = 0
								}

								if tmp_tmp_1 == 124 {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 2
										for _g < _g1 {
											var _hx_tmp_236 int32
											_ = _hx_tmp_236
											_hx_tmp_236 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_236
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_237 int32
											_ = _hx_tmp_237
											var _hx_tmp_238 bool
											_ = _hx_tmp_238
											if tmp_pos_1 < 0 {
												_hx_tmp_238 = true
											} else {
												var _hx_tmp_239 int32
												_ = _hx_tmp_239
												_hx_tmp_239 = tmp_pos_1
												var _hx_tmp_240 int32
												_ = _hx_tmp_240
												_hx_tmp_240 = _hx_tmp_239
												_hx_tmp_238 = (_hx_tmp_240 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_238 {
												_hx_tmp_237 = -1
											} else {
												_hx_tmp_237 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_237) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpOr{false}, position)
									this.Hx_Field_previous = token
									return token
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_241 int32
											_ = _hx_tmp_241
											_hx_tmp_241 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_241
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_242 int32
											_ = _hx_tmp_242
											var _hx_tmp_243 bool
											_ = _hx_tmp_243
											if tmp_pos_1 < 0 {
												_hx_tmp_243 = true
											} else {
												var _hx_tmp_244 int32
												_ = _hx_tmp_244
												_hx_tmp_244 = tmp_pos_1
												var _hx_tmp_245 int32
												_ = _hx_tmp_245
												_hx_tmp_245 = _hx_tmp_244
												_hx_tmp_243 = (_hx_tmp_245 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_243 {
												_hx_tmp_242 = -1
											} else {
												_hx_tmp_242 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_242) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									this.VTable.Hx_Field_error("Expected |", false)
									var position *Hx_Obj_loreline_position
									_ = position
									position = startPos
									if position == nil {
										position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
									}

									position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
									var token *Hx_Obj_loreline_token
									_ = token
									token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_OpOr{false}, position)
									this.Hx_Field_previous = token
									return token
								}
							}
						}
					}

				default:
					{
						var c1 int32
						_ = c1
						c1 = c
						if (c1 >= 48) && (c1 <= 57) {
							return this.VTable.Hx_Field_readNumber()
						} else {
							var tmp_c1_1 int32
							_ = tmp_c1_1
							tmp_c1_1 = c
							if (((tmp_c1_1 >= 97) && (tmp_c1_1 <= 122)) || ((tmp_c1_1 >= 65) && (tmp_c1_1 <= 90))) || (tmp_c1_1 == 95) {
								return this.VTable.Hx_Field_readIdentifier()
							} else {
								this.VTable.Hx_Field_error(Hx_Field_go_haxe_hxstring_concat("Unexpected character: ", Hx_Field_go_haxe_hxstring_fromCharCode(c)), true)
								return nil
							}
						}
					}
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_countIndentation() int32 {
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
	var spaces int32
	_ = spaces
	spaces = 0
	for pos < this.Hx_Field_length {
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
		if c == 32 {
			spaces++
		} else {
			if c == 9 {
				spaces++
			} else {
				break
			}
		}

		pos++
	}

	var tmp bool
	_ = tmp
	var tmp1 bool
	_ = tmp1
	if pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp1 = ((_hx_tmp_4) == 10)
	} else {
		tmp1 = true
	}

	if !tmp1 {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_9 {
			_hx_tmp_8 = -1
		} else {
			_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp = ((_hx_tmp_8) == 13)
	} else {
		tmp = true
	}

	if tmp {
		var _hx_tmp_12 HxArray[int32]
		_ = _hx_tmp_12
		_hx_tmp_12 = this.Hx_Field_indentStack
		return _hx_tmp_12.Get((this.Hx_Field_indentStack.Len() - 1))
	}

	return spaces
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_parentBlockType() Hx_Enum_loreline_tokentype {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		if ((this.Hx_Field_stack.Get(i) != ((int32)(12))) && (this.Hx_Field_stack.Get(i) != ((int32)(13)))) && (this.Hx_Field_stack.Get(i) != ((int32)(14))) {
			switch this.Hx_Field_stack.Get(i) {
			case 0, 1:
				{
					return Hx_Enum_loreline_tokentype_KwBeat{}
				}

			case 2, 3:
				{
					return Hx_Enum_loreline_tokentype_KwBeat{}
				}

			case 4, 5:
				{
					return Hx_Enum_loreline_tokentype_KwState{}
				}

			case 6, 7:
				{
					return Hx_Enum_loreline_tokentype_KwCharacter{}
				}

			case 8, 9:
				{
					return Hx_Enum_loreline_tokentype_KwBeat{}
				}

			case 10, 11:
				{
					return Hx_Enum_loreline_tokentype_KwBeat{}
				}

			case 12:
				{
					return Hx_Enum_loreline_tokentype_LBrace{}
				}

			case 13:
				{
					return Hx_Enum_loreline_tokentype_Indent{}
				}

			case 14:
				{
					return Hx_Enum_loreline_tokentype_LBracket{}
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown("unreachable"))
				}
			}
		}

		i--
	}

	return Hx_Enum_loreline_tokentype_KwBeat{}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_inBeat() bool {
	return (this.VTable.Hx_Field_parentBlockType() == Hx_Enum_loreline_tokentype_KwBeat{})
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_inChoiceRoot() bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if this.Hx_Field_stack.Len() > 0 {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 HxArray[int32]
		_ = _hx_tmp_2
		_hx_tmp_2 = this.Hx_Field_stack
		if _hx_tmp_2.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(1)) {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_3 HxArray[int32]
			_ = _hx_tmp_3
			_hx_tmp_3 = this.Hx_Field_stack
			var _hx_tmp_4 HxArray[int32]
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_1 = (_hx_tmp_4.Get((this.Hx_Field_stack.Len() - 1)) == ((int32)(0)))
		}

		_hx_tmp_0 = (_hx_tmp_1)
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_inWhenRoot() bool {
	if this.Hx_Field_stack.Len() == 0 {
		return false
	}

	var _hx_tmp_0 HxArray[int32]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_stack
	var top int32
	_ = top
	top = _hx_tmp_0.Get((this.Hx_Field_stack.Len() - 1))
	if (top == ((int32)(3))) || (top == ((int32)(2))) {
		return true
	}

	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if top == ((int32)(13)) {
		_hx_tmp_2 = (this.Hx_Field_stack.Len() > 1)
	} else {
		_hx_tmp_2 = false
	}

	if _hx_tmp_2 {
		var _hx_tmp_3 HxArray[int32]
		_ = _hx_tmp_3
		_hx_tmp_3 = this.Hx_Field_stack
		var _hx_tmp_4 HxArray[int32]
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_3
		_hx_tmp_1 = (_hx_tmp_4.Get((this.Hx_Field_stack.Len() - 2)) == ((int32)(2)))
	} else {
		_hx_tmp_1 = false
	}

	return _hx_tmp_1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isWhenStart(pos int32) bool {
	var end int32
	_ = end
	end = this.VTable.Hx_Field_wordEnd(pos, "when")
	if end == -1 {
		return false
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (this.VTable.Hx_Field_parentBlockType() != Hx_Enum_loreline_tokentype_KwBeat{}) {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = this.VTable.Hx_Field_inChoiceRoot()
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = this.VTable.Hx_Field_inWhenRoot()
	}

	if _hx_tmp_0 {
		return false
	}

	var lineStart int32
	_ = lineStart
	lineStart = pos
	for {
		var tmp bool
		_ = tmp
		if lineStart > 0 {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (lineStart - 1)
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if pos < 0 {
				_hx_tmp_3 = true
			} else {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = pos
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_4
				_hx_tmp_3 = (_hx_tmp_5 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_3 {
				_hx_tmp_2 = -1
			} else {
				_hx_tmp_2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_2) != 32 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (lineStart - 1)
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				if tmp_pos_1 < 0 {
					_hx_tmp_7 = true
				} else {
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = tmp_pos_1
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _hx_tmp_8
					_hx_tmp_7 = (_hx_tmp_9 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_7 {
					_hx_tmp_6 = -1
				} else {
					_hx_tmp_6 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				tmp = ((_hx_tmp_6) == 9)
			} else {
				tmp = true
			}
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		lineStart--
	}

	var tmp bool
	_ = tmp
	var tmp1 bool
	_ = tmp1
	if lineStart > 0 {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = (lineStart - 1)
		var _hx_tmp_10 int32
		_ = _hx_tmp_10
		var _hx_tmp_11 bool
		_ = _hx_tmp_11
		if pos < 0 {
			_hx_tmp_11 = true
		} else {
			var _hx_tmp_12 int32
			_ = _hx_tmp_12
			_hx_tmp_12 = pos
			var _hx_tmp_13 int32
			_ = _hx_tmp_13
			_hx_tmp_13 = _hx_tmp_12
			_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_11 {
			_hx_tmp_10 = -1
		} else {
			_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp1 = ((_hx_tmp_10) != 10)
	} else {
		tmp1 = false
	}

	if tmp1 {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = (lineStart - 1)
		var _hx_tmp_14 int32
		_ = _hx_tmp_14
		var _hx_tmp_15 bool
		_ = _hx_tmp_15
		if pos < 0 {
			_hx_tmp_15 = true
		} else {
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = pos
			var _hx_tmp_17 int32
			_ = _hx_tmp_17
			_hx_tmp_17 = _hx_tmp_16
			_hx_tmp_15 = (_hx_tmp_17 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_15 {
			_hx_tmp_14 = -1
		} else {
			_hx_tmp_14 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp = ((_hx_tmp_14) != 13)
	} else {
		tmp = false
	}

	if tmp {
		return false
	}

	var indent int32
	_ = indent
	indent = (pos - lineStart)
	var p int32
	_ = p
	p = this.VTable.Hx_Field_skipInlineSpacesAndComments(end)
	var tmp2 bool
	_ = tmp2
	if p < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_18 int32
		_ = _hx_tmp_18
		var _hx_tmp_19 bool
		_ = _hx_tmp_19
		if p < 0 {
			_hx_tmp_19 = true
		} else {
			var _hx_tmp_20 int32
			_ = _hx_tmp_20
			_hx_tmp_20 = p
			var _hx_tmp_21 int32
			_ = _hx_tmp_21
			_hx_tmp_21 = _hx_tmp_20
			_hx_tmp_19 = (_hx_tmp_21 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_19 {
			_hx_tmp_18 = -1
		} else {
			_hx_tmp_18 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		var c int32
		_ = c
		c = _hx_tmp_18
		tmp2 = ((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95))
	} else {
		tmp2 = false
	}

	if tmp2 {
		for {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if p < this.Hx_Field_length {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_22 int32
				_ = _hx_tmp_22
				var _hx_tmp_23 bool
				_ = _hx_tmp_23
				if p < 0 {
					_hx_tmp_23 = true
				} else {
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = p
					var _hx_tmp_25 int32
					_ = _hx_tmp_25
					_hx_tmp_25 = _hx_tmp_24
					_hx_tmp_23 = (_hx_tmp_25 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_23 {
					_hx_tmp_22 = -1
				} else {
					_hx_tmp_22 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
				}

				var c int32
				_ = c
				c = _hx_tmp_22
				tmp_tmp_1 = (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)))
			} else {
				tmp_tmp_1 = false
			}

			if !tmp_tmp_1 {
				break
			}

			p++
		}

		p = this.VTable.Hx_Field_skipInlineSpacesAndComments(p)
	}

	var tmp3 bool
	_ = tmp3
	if p < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_26 int32
		_ = _hx_tmp_26
		var _hx_tmp_27 bool
		_ = _hx_tmp_27
		if p < 0 {
			_hx_tmp_27 = true
		} else {
			var _hx_tmp_28 int32
			_ = _hx_tmp_28
			_hx_tmp_28 = p
			var _hx_tmp_29 int32
			_ = _hx_tmp_29
			_hx_tmp_29 = _hx_tmp_28
			_hx_tmp_27 = (_hx_tmp_29 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_27 {
			_hx_tmp_26 = -1
		} else {
			_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		tmp3 = ((_hx_tmp_26) == 123)
	} else {
		tmp3 = false
	}

	if tmp3 {
		return true
	}

	var tmp4 bool
	_ = tmp4
	var tmp5 bool
	_ = tmp5
	if p < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_30 int32
		_ = _hx_tmp_30
		var _hx_tmp_31 bool
		_ = _hx_tmp_31
		if p < 0 {
			_hx_tmp_31 = true
		} else {
			var _hx_tmp_32 int32
			_ = _hx_tmp_32
			_hx_tmp_32 = p
			var _hx_tmp_33 int32
			_ = _hx_tmp_33
			_hx_tmp_33 = _hx_tmp_32
			_hx_tmp_31 = (_hx_tmp_33 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_31 {
			_hx_tmp_30 = -1
		} else {
			_hx_tmp_30 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		tmp5 = ((_hx_tmp_30) == 47)
	} else {
		tmp5 = false
	}

	if tmp5 && ((p + 1) < this.Hx_Field_length) {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = (p + 1)
		var _hx_tmp_34 int32
		_ = _hx_tmp_34
		var _hx_tmp_35 bool
		_ = _hx_tmp_35
		if pos < 0 {
			_hx_tmp_35 = true
		} else {
			var _hx_tmp_36 int32
			_ = _hx_tmp_36
			_hx_tmp_36 = pos
			var _hx_tmp_37 int32
			_ = _hx_tmp_37
			_hx_tmp_37 = _hx_tmp_36
			_hx_tmp_35 = (_hx_tmp_37 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_35 {
			_hx_tmp_34 = -1
		} else {
			_hx_tmp_34 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp4 = ((_hx_tmp_34) == 47)
	} else {
		tmp4 = false
	}

	if tmp4 {
		for {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			var tmp_tmp1_1 bool
			_ = tmp_tmp1_1
			if p < this.Hx_Field_length {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_38 int32
				_ = _hx_tmp_38
				var _hx_tmp_39 bool
				_ = _hx_tmp_39
				if p < 0 {
					_hx_tmp_39 = true
				} else {
					var _hx_tmp_40 int32
					_ = _hx_tmp_40
					_hx_tmp_40 = p
					var _hx_tmp_41 int32
					_ = _hx_tmp_41
					_hx_tmp_41 = _hx_tmp_40
					_hx_tmp_39 = (_hx_tmp_41 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_39 {
					_hx_tmp_38 = -1
				} else {
					_hx_tmp_38 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
				}

				tmp_tmp1_1 = ((_hx_tmp_38) != 10)
			} else {
				tmp_tmp1_1 = false
			}

			if tmp_tmp1_1 {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
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
					_hx_tmp_43 = (_hx_tmp_45 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_43 {
					_hx_tmp_42 = -1
				} else {
					_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
				}

				tmp_tmp_1 = ((_hx_tmp_42) != 13)
			} else {
				tmp_tmp_1 = false
			}

			if !tmp_tmp_1 {
				break
			}

			p++
		}
	}

	var tmp6 bool
	_ = tmp6
	var tmp7 bool
	_ = tmp7
	if p < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_47 = (_hx_tmp_49 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_47 {
			_hx_tmp_46 = -1
		} else {
			_hx_tmp_46 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		tmp7 = ((_hx_tmp_46) != 10)
	} else {
		tmp7 = false
	}

	if tmp7 {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_51 = (_hx_tmp_53 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_51 {
			_hx_tmp_50 = -1
		} else {
			_hx_tmp_50 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		tmp6 = ((_hx_tmp_50) != 13)
	} else {
		tmp6 = false
	}

	if tmp6 {
		return false
	}

	return this.VTable.Hx_Field_nextLineOpensBlock(p, indent)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_nextLineOpensBlock(p int32, indent int32) bool {
	for p < this.Hx_Field_length {
		for {
			var tmp bool
			_ = tmp
			var tmp1 bool
			_ = tmp1
			if p < this.Hx_Field_length {
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

				tmp1 = ((_hx_tmp_0) != 10)
			} else {
				tmp1 = false
			}

			if tmp1 {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
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
					_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_5 {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
				}

				tmp = ((_hx_tmp_4) != 13)
			} else {
				tmp = false
			}

			if !tmp {
				break
			}

			p++
		}

		for {
			var tmp bool
			_ = tmp
			if p < this.Hx_Field_length {
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

				if (_hx_tmp_8) != 10 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					var _hx_tmp_13 bool
					_ = _hx_tmp_13
					if p < 0 {
						_hx_tmp_13 = true
					} else {
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = p
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = _hx_tmp_14
						_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_13 {
						_hx_tmp_12 = -1
					} else {
						_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
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

			p++
		}

		var lineStart int32
		_ = lineStart
		lineStart = p
		for {
			var tmp bool
			_ = tmp
			if p < this.Hx_Field_length {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_16 int32
				_ = _hx_tmp_16
				var _hx_tmp_17 bool
				_ = _hx_tmp_17
				if p < 0 {
					_hx_tmp_17 = true
				} else {
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					_hx_tmp_18 = p
					var _hx_tmp_19 int32
					_ = _hx_tmp_19
					_hx_tmp_19 = _hx_tmp_18
					_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_17 {
					_hx_tmp_16 = -1
				} else {
					_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
				}

				if (_hx_tmp_16) != 32 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					var _hx_tmp_21 bool
					_ = _hx_tmp_21
					if p < 0 {
						_hx_tmp_21 = true
					} else {
						var _hx_tmp_22 int32
						_ = _hx_tmp_22
						_hx_tmp_22 = p
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = _hx_tmp_22
						_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_21 {
						_hx_tmp_20 = -1
					} else {
						_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
					}

					tmp = ((_hx_tmp_20) == 9)
				} else {
					tmp = true
				}
			} else {
				tmp = false
			}

			if !tmp {
				break
			}

			p++
		}

		if p >= this.Hx_Field_length {
			return false
		}

		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_24 int32
		_ = _hx_tmp_24
		var _hx_tmp_25 bool
		_ = _hx_tmp_25
		if p < 0 {
			_hx_tmp_25 = true
		} else {
			var _hx_tmp_26 int32
			_ = _hx_tmp_26
			_hx_tmp_26 = p
			var _hx_tmp_27 int32
			_ = _hx_tmp_27
			_hx_tmp_27 = _hx_tmp_26
			_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_25 {
			_hx_tmp_24 = -1
		} else {
			_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		var c int32
		_ = c
		c = _hx_tmp_24
		if (c == 10) || (c == 13) {
			continue
		}

		if c == 123 {
			return true
		}

		return ((p - lineStart) > indent)
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_inStateOrCharacter() bool {
	var parent Hx_Enum_loreline_tokentype
	_ = parent
	parent = this.VTable.Hx_Field_parentBlockType()
	return ((parent == Hx_Enum_loreline_tokentype_KwState{}) || (parent == Hx_Enum_loreline_tokentype_KwCharacter{}))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readLineBreak() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
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

	if (_hx_tmp_0) == 13 {
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_4
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if tmp_pos_1 < 0 {
					_hx_tmp_6 = true
				} else {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = tmp_pos_1
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_6 {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_5) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}

		var tmp bool
		_ = tmp
		if this.Hx_Field_pos < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
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

			tmp = ((_hx_tmp_9) == 10)
		} else {
			tmp = false
		}

		if tmp {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_13 int32
				_ = _hx_tmp_13
				_hx_tmp_13 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_13
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_14 int32
				_ = _hx_tmp_14
				var _hx_tmp_15 bool
				_ = _hx_tmp_15
				if tmp_pos_1 < 0 {
					_hx_tmp_15 = true
				} else {
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					_hx_tmp_16 = tmp_pos_1
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					_hx_tmp_17 = _hx_tmp_16
					_hx_tmp_15 = (_hx_tmp_17 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_15 {
					_hx_tmp_14 = -1
				} else {
					_hx_tmp_14 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_14) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	} else {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 1
		for _g < _g1 {
			var _hx_tmp_18 int32
			_ = _hx_tmp_18
			_hx_tmp_18 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_18
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
			var _hx_tmp_19 int32
			_ = _hx_tmp_19
			var _hx_tmp_20 bool
			_ = _hx_tmp_20
			if tmp_pos_1 < 0 {
				_hx_tmp_20 = true
			} else {
				var _hx_tmp_21 int32
				_ = _hx_tmp_21
				_hx_tmp_21 = tmp_pos_1
				var _hx_tmp_22 int32
				_ = _hx_tmp_22
				_hx_tmp_22 = _hx_tmp_21
				_hx_tmp_20 = (_hx_tmp_22 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_20 {
				_hx_tmp_19 = -1
			} else {
				_hx_tmp_19 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			if (_hx_tmp_19) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LineBreak{}, position)
	this.Hx_Field_previous = token
	return token
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isWhitespace(c int32) bool {
	return ((c == 32) || (c == 9))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_matchIdentifier(pos int32) HxNullable[string] {
	if this.Hx_Field_length == 0 {
		return HxNullable[string]{Value: HxStringNull}
	}

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
		return HxNullable[string]{Value: HxStringNull}
	}

	var identifierLength int32
	_ = identifierLength
	identifierLength = 1
	for identifierLength < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + identifierLength)
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if pos1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
		}

		var tmp_c_1 int32
		_ = tmp_c_1
		tmp_c_1 = _hx_tmp_4
		if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
			break
		}

		identifierLength++
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
	var _hx_tmp_8 HxArray[int32]
	_ = _hx_tmp_8
	_hx_tmp_8 = v1
	var _hx_tmp_9 int32
	_ = _hx_tmp_9
	_hx_tmp_9 = pos
	var _hx_tmp_10 int32
	_ = _hx_tmp_10
	if (HxNullable[int32]{Value: identifierLength, Valid: true}).Valid != false {
		_hx_tmp_10 = (pos + identifierLength)
	} else {
		_hx_tmp_10 = HxMakeArrayView[any](v1).Len()
	}

	return HxNullableString(Hx_Field_loreline_utf8vector_slice(_hx_tmp_8, _hx_tmp_9, _hx_tmp_10))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field__skipWhitespaceAndComments(pos int32, stopNextLine__isOptional any) int32 {
	stopNextLine := HxDefault[bool]()
	_ = stopNextLine

	if stopNextLine__isOptional != nil {
		stopNextLine = HxConvert[bool](stopNextLine__isOptional)
	} else {
		stopNextLine = false
	}
	{
		if stopNextLine {
			return this.VTable.Hx_Field__skipWhitespaceAndCommentsStopNextLine(pos)
		} else {
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

					pos++
					foundContent = true
				}

				if pos < (this.Hx_Field_length - 1) {
					var v HxArray[int32]
					_ = v
					v = this.Hx_Field_input
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
						_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
					}

					if _hx_tmp_9 {
						_hx_tmp_8 = -1
					} else {
						_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
					}

					if (_hx_tmp_8) == 47 {
						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var pos1 int32
						_ = pos1
						pos1 = (pos + 1)
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						var _hx_tmp_13 bool
						_ = _hx_tmp_13
						if pos1 < 0 {
							_hx_tmp_13 = true
						} else {
							var _hx_tmp_14 int32
							_ = _hx_tmp_14
							_hx_tmp_14 = pos1
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							_hx_tmp_15 = _hx_tmp_14
							_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_13 {
							_hx_tmp_12 = -1
						} else {
							_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
						}

						if (_hx_tmp_12) == 47 {
							pos = startPos
							return pos
						} else {
							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var tmp_pos1_1 int32
							_ = tmp_pos1_1
							tmp_pos1_1 = (pos + 1)
							var _hx_tmp_16 int32
							_ = _hx_tmp_16
							var _hx_tmp_17 bool
							_ = _hx_tmp_17
							if tmp_pos1_1 < 0 {
								_hx_tmp_17 = true
							} else {
								var _hx_tmp_18 int32
								_ = _hx_tmp_18
								_hx_tmp_18 = tmp_pos1_1
								var _hx_tmp_19 int32
								_ = _hx_tmp_19
								_hx_tmp_19 = _hx_tmp_18
								_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_17 {
								_hx_tmp_16 = -1
							} else {
								_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos1_1))))
							}

							if (_hx_tmp_16) == 42 {
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
									var _hx_tmp_20 int32
									_ = _hx_tmp_20
									var _hx_tmp_21 bool
									_ = _hx_tmp_21
									if pos < 0 {
										_hx_tmp_21 = true
									} else {
										var _hx_tmp_22 int32
										_ = _hx_tmp_22
										_hx_tmp_22 = pos
										var _hx_tmp_23 int32
										_ = _hx_tmp_23
										_hx_tmp_23 = _hx_tmp_22
										_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_3).Len())
									}

									if _hx_tmp_21 {
										_hx_tmp_20 = -1
									} else {
										_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(pos))))
									}

									if (_hx_tmp_20) == 42 {
										var tmp_v_4 HxArray[int32]
										_ = tmp_v_4
										tmp_v_4 = this.Hx_Field_input
										var tmp_pos1_2 int32
										_ = tmp_pos1_2
										tmp_pos1_2 = (pos + 1)
										var _hx_tmp_24 int32
										_ = _hx_tmp_24
										var _hx_tmp_25 bool
										_ = _hx_tmp_25
										if tmp_pos1_2 < 0 {
											_hx_tmp_25 = true
										} else {
											var _hx_tmp_26 int32
											_ = _hx_tmp_26
											_hx_tmp_26 = tmp_pos1_2
											var _hx_tmp_27 int32
											_ = _hx_tmp_27
											_hx_tmp_27 = _hx_tmp_26
											_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_4).Len())
										}

										if _hx_tmp_25 {
											_hx_tmp_24 = -1
										} else {
											_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos1_2))))
										}

										tmp = ((_hx_tmp_24) == 47)
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
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field__skipWhitespaceAndCommentsStopNextLine(pos int32) int32 {
	var startPos int32
	_ = startPos
	startPos = pos
	var foundContent bool
	_ = foundContent
	foundContent = false
	var isNextLine bool
	_ = isNextLine
	isNextLine = false
	for pos < this.Hx_Field_length {
		for pos < this.Hx_Field_length {
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
			if (((c == 32) || (c == 9)) || (c == 13)) || (!isNextLine && (c == 10)) {
				if c == 10 {
					isNextLine = true
				}

				pos++
				foundContent = true
			} else {
				break
			}
		}

		if (pos < (this.Hx_Field_length - 1)) && !isNextLine {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
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
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_4) == 47 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				var _hx_tmp_9 bool
				_ = _hx_tmp_9
				if pos1 < 0 {
					_hx_tmp_9 = true
				} else {
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					_hx_tmp_10 = pos1
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					_hx_tmp_11 = _hx_tmp_10
					_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_9 {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
				}

				if (_hx_tmp_8) == 47 {
					pos = (pos + 2)
					foundContent = true
					for {
						var tmp bool
						_ = tmp
						var tmp1 bool
						_ = tmp1
						if pos < this.Hx_Field_length {
							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
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
								_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_13 {
								_hx_tmp_12 = -1
							} else {
								_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
							}

							tmp1 = ((_hx_tmp_12) != 10)
						} else {
							tmp1 = false
						}

						if tmp1 {
							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
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
								_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_17 {
								_hx_tmp_16 = -1
							} else {
								_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
							}

							tmp = ((_hx_tmp_16) != 13)
						} else {
							tmp = false
						}

						if !tmp {
							break
						}

						pos++
					}

					continue
				} else {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos1_1 int32
					_ = tmp_pos1_1
					tmp_pos1_1 = (pos + 1)
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					var _hx_tmp_21 bool
					_ = _hx_tmp_21
					if tmp_pos1_1 < 0 {
						_hx_tmp_21 = true
					} else {
						var _hx_tmp_22 int32
						_ = _hx_tmp_22
						_hx_tmp_22 = tmp_pos1_1
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = _hx_tmp_22
						_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_21 {
						_hx_tmp_20 = -1
					} else {
						_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos1_1))))
					}

					if (_hx_tmp_20) == 42 {
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
							var _hx_tmp_24 int32
							_ = _hx_tmp_24
							var _hx_tmp_25 bool
							_ = _hx_tmp_25
							if pos < 0 {
								_hx_tmp_25 = true
							} else {
								var _hx_tmp_26 int32
								_ = _hx_tmp_26
								_hx_tmp_26 = pos
								var _hx_tmp_27 int32
								_ = _hx_tmp_27
								_hx_tmp_27 = _hx_tmp_26
								_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_3).Len())
							}

							if _hx_tmp_25 {
								_hx_tmp_24 = -1
							} else {
								_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(pos))))
							}

							if (_hx_tmp_24) == 42 {
								var tmp_v_4 HxArray[int32]
								_ = tmp_v_4
								tmp_v_4 = this.Hx_Field_input
								var tmp_pos1_2 int32
								_ = tmp_pos1_2
								tmp_pos1_2 = (pos + 1)
								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								var _hx_tmp_29 bool
								_ = _hx_tmp_29
								if tmp_pos1_2 < 0 {
									_hx_tmp_29 = true
								} else {
									var _hx_tmp_30 int32
									_ = _hx_tmp_30
									_hx_tmp_30 = tmp_pos1_2
									var _hx_tmp_31 int32
									_ = _hx_tmp_31
									_hx_tmp_31 = _hx_tmp_30
									_hx_tmp_29 = (_hx_tmp_31 >= HxMakeArrayView[any](tmp_v_4).Len())
								}

								if _hx_tmp_29 {
									_hx_tmp_28 = -1
								} else {
									_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos1_2))))
								}

								tmp = ((_hx_tmp_28) == 47)
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

func (this *Hx_Obj_loreline_lexer) Hx_Field_isIfStart(pos int32) bool {
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
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

	if (_hx_tmp_0) != 105 {
		return false
	}

	pos++
	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
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
		_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v1).Len())
	}

	if _hx_tmp_5 {
		_hx_tmp_4 = -1
	} else {
		_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos))))
	}

	if (_hx_tmp_4) != 102 {
		return false
	}

	pos++
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

		var c int32
		_ = c
		c = _hx_tmp_8
		if !((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) {
			var tmp_v_2 HxArray[int32]
			_ = tmp_v_2
			tmp_v_2 = this.Hx_Field_input
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
				_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
			}

			if _hx_tmp_13 {
				_hx_tmp_12 = -1
			} else {
				_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_12
			tmp = ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))
		} else {
			tmp = true
		}
	} else {
		tmp = false
	}

	if tmp {
		return false
	}

	return this.VTable.Hx_Field_isConditionUntilLineEnd(pos)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isConditionUntilLineEnd(pos int32) bool {
	var expectOperand bool
	_ = expectOperand
	expectOperand = true
	for {
		pos = this.VTable.Hx_Field_skipInlineSpacesAndComments(pos)
		if expectOperand {
			if pos >= this.Hx_Field_length {
				return false
			}

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
			var next int32
			_ = next
			if (pos + 1) < this.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if pos1 < 0 {
					_hx_tmp_5 = true
				} else {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = pos1
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _hx_tmp_6
					_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_5 {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
				}

				next = _hx_tmp_4
			} else {
				next = 0
			}

			if (c == 33) && (next != 61) {
				pos++
				continue
			}

			if (c == 45) && (next != 62) {
				pos++
				continue
			}

			var notEnd int32
			_ = notEnd
			notEnd = this.VTable.Hx_Field_wordEnd(pos, "not")
			if notEnd != -1 {
				pos = notEnd
				continue
			}

			pos = this.VTable.Hx_Field_conditionOperandEnd(pos)
			if pos == -1 {
				return false
			}

			expectOperand = false
		} else {
			if this.VTable.Hx_Field_isConditionEnd(pos) {
				return true
			}

			pos = this.VTable.Hx_Field_conditionOperatorEnd(pos)
			if pos == -1 {
				return false
			}

			expectOperand = true
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isConditionEnd(pos int32) bool {
	if pos >= this.Hx_Field_length {
		return true
	}

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
	var next int32
	_ = next
	if (pos + 1) < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + 1)
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if pos1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
		}

		next = _hx_tmp_4
	} else {
		next = 0
	}

	return ((((((c == 10) || (c == 13)) || (c == 123)) || (c == 35)) || ((c == 47) && (next == 47))) || ((c == 45) && (next == 62)))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_skipInlineSpacesAndComments(pos int32) int32 {
	for pos < this.Hx_Field_length {
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
		if (c == 32) || (c == 9) {
			pos++
		} else {
			var tmp bool
			_ = tmp
			if (c == 47) && ((pos + 1) < this.Hx_Field_length) {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if pos1 < 0 {
					_hx_tmp_5 = true
				} else {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = pos1
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _hx_tmp_6
					_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_5 {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
				}

				tmp = ((_hx_tmp_4) == 42)
			} else {
				tmp = false
			}

			if tmp {
				var p int32
				_ = p
				p = (pos + 2)
				for {
					var tmp_tmp_1 bool
					_ = tmp_tmp_1
					if (p + 1) < this.Hx_Field_length {
						var tmp1 bool
						_ = tmp1
						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
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
							_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_9 {
							_hx_tmp_8 = -1
						} else {
							_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
						}

						if (_hx_tmp_8) == 42 {
							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var pos int32
							_ = pos
							pos = (p + 1)
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
								_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_13 {
								_hx_tmp_12 = -1
							} else {
								_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
							}

							tmp1 = ((_hx_tmp_12) == 47)
						} else {
							tmp1 = false
						}

						tmp_tmp_1 = !tmp1
					} else {
						tmp_tmp_1 = false
					}

					if !tmp_tmp_1 {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					var _hx_tmp_17 bool
					_ = _hx_tmp_17
					if p < 0 {
						_hx_tmp_17 = true
					} else {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = p
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = _hx_tmp_18
						_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_17 {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
					}

					if (_hx_tmp_16) == 10 {
						return pos
					}

					p++
				}

				if (p + 1) >= this.Hx_Field_length {
					return pos
				}

				pos = (p + 2)
			} else {
				break
			}
		}
	}

	return pos
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_conditionOperandEnd(pos int32) int32 {
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
	if (c >= 48) && (c <= 57) {
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
				tmp = ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))
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
		var tmp1 bool
		_ = tmp1
		if (pos + 1) < this.Hx_Field_length {
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

			tmp1 = ((_hx_tmp_8) == 46)
		} else {
			tmp1 = false
		}

		if tmp1 {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos1 int32
			_ = pos1
			pos1 = (pos + 1)
			var _hx_tmp_12 int32
			_ = _hx_tmp_12
			var _hx_tmp_13 bool
			_ = _hx_tmp_13
			if pos1 < 0 {
				_hx_tmp_13 = true
			} else {
				var _hx_tmp_14 int32
				_ = _hx_tmp_14
				_hx_tmp_14 = pos1
				var _hx_tmp_15 int32
				_ = _hx_tmp_15
				_hx_tmp_15 = _hx_tmp_14
				_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_13 {
				_hx_tmp_12 = -1
			} else {
				_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_12
			tmp = ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))
		} else {
			tmp = false
		}

		if tmp {
			pos++
			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if pos < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
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
						_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_17 {
						_hx_tmp_16 = -1
					} else {
						_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_16
					tmp_tmp_1 = ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				pos++
			}
		}

		return pos
	}

	if c == 34 {
		return this.VTable.Hx_Field_conditionGroupEnd(pos)
	}

	if (c == 40) || (c == 91) {
		pos = this.VTable.Hx_Field_conditionGroupEnd(pos)
	} else {
		if (((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95) {
			var start int32
			_ = start
			start = pos
			for {
				var tmp bool
				_ = tmp
				if pos < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					var _hx_tmp_21 bool
					_ = _hx_tmp_21
					if pos < 0 {
						_hx_tmp_21 = true
					} else {
						var _hx_tmp_22 int32
						_ = _hx_tmp_22
						_hx_tmp_22 = pos
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = _hx_tmp_22
						_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_21 {
						_hx_tmp_20 = -1
					} else {
						_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_20
					tmp = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
				} else {
					tmp = false
				}

				if !tmp {
					break
				}

				pos++
			}

			{
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (pos - start), Valid: true}

				var _hx_tmp_24 HxArray[int32]
				_ = _hx_tmp_24
				_hx_tmp_24 = tmp_v_1
				var _hx_tmp_25 int32
				_ = _hx_tmp_25
				_hx_tmp_25 = start
				var _hx_tmp_26 int32
				_ = _hx_tmp_26
				if (_hx_reserved_len).Valid != false {
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					_hx_tmp_27 = start
					var _hx_tmp_28 int32
					_ = _hx_tmp_28
					_hx_tmp_28 = _hx_tmp_27
					_hx_tmp_26 = (_hx_tmp_28 + (_hx_reserved_len).Value)
				} else {
					_hx_tmp_26 = HxMakeArrayView[any](tmp_v_1).Len()
				}

				switch Hx_Field_loreline_utf8vector_slice(_hx_tmp_24, _hx_tmp_25, _hx_tmp_26) {
				case "and", "is", "not", "or":
					{
						return -1
					}

				default:
					{
					}
				}
			}
		} else {
			return -1
		}
	}

	for (pos != -1) && (pos < this.Hx_Field_length) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_29 int32
		_ = _hx_tmp_29
		var _hx_tmp_30 bool
		_ = _hx_tmp_30
		if pos < 0 {
			_hx_tmp_30 = true
		} else {
			var _hx_tmp_31 int32
			_ = _hx_tmp_31
			_hx_tmp_31 = pos
			var _hx_tmp_32 int32
			_ = _hx_tmp_32
			_hx_tmp_32 = _hx_tmp_31
			_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_30 {
			_hx_tmp_29 = -1
		} else {
			_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		var p int32
		_ = p
		p = _hx_tmp_29
		if p == 46 {
			pos++
			var tmp bool
			_ = tmp
			if pos < this.Hx_Field_length {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var _hx_tmp_33 int32
				_ = _hx_tmp_33
				var _hx_tmp_34 bool
				_ = _hx_tmp_34
				if pos < 0 {
					_hx_tmp_34 = true
				} else {
					var _hx_tmp_35 int32
					_ = _hx_tmp_35
					_hx_tmp_35 = pos
					var _hx_tmp_36 int32
					_ = _hx_tmp_36
					_hx_tmp_36 = _hx_tmp_35
					_hx_tmp_34 = (_hx_tmp_36 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_34 {
					_hx_tmp_33 = -1
				} else {
					_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				var tmp_c_1 int32
				_ = tmp_c_1
				tmp_c_1 = _hx_tmp_33
				tmp = !((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95))
			} else {
				tmp = true
			}

			if tmp {
				return -1
			}

			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if pos < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_37 int32
					_ = _hx_tmp_37
					var _hx_tmp_38 bool
					_ = _hx_tmp_38
					if pos < 0 {
						_hx_tmp_38 = true
					} else {
						var _hx_tmp_39 int32
						_ = _hx_tmp_39
						_hx_tmp_39 = pos
						var _hx_tmp_40 int32
						_ = _hx_tmp_40
						_hx_tmp_40 = _hx_tmp_39
						_hx_tmp_38 = (_hx_tmp_40 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_38 {
						_hx_tmp_37 = -1
					} else {
						_hx_tmp_37 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_37
					tmp_tmp_1 = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				pos++
			}
		} else {
			if (p == 40) || (p == 91) {
				pos = this.VTable.Hx_Field_conditionGroupEnd(pos)
			} else {
				break
			}
		}
	}

	return pos
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_conditionGroupEnd(pos int32) int32 {
	var closers HxArray[int32]
	_ = closers
	closers = HxMakeArray[int32]()
	for pos < this.Hx_Field_length {
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
		if (c == 10) || (c == 13) {
			return -1
		}

		if c == 34 {
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

					tmp = ((_hx_tmp_4) != 34)
				} else {
					tmp = false
				}

				if !tmp {
					break
				}

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

				var sc int32
				_ = sc
				sc = _hx_tmp_8
				if (sc == 10) || (sc == 13) {
					return -1
				}

				if sc == 92 {
					pos++
				}

				pos++
			}

			if pos >= this.Hx_Field_length {
				return -1
			}

			pos++
			if closers.Len() == 0 {
				return pos
			}

			continue
		}

		if c == 40 {
			_ = (Hx_Array_Push[int32](closers, 41))
		} else {
			if c == 91 {
				_ = (Hx_Array_Push[int32](closers, 93))
			} else {
				if (c == 41) || (c == 93) {
					var _hx_tmp_12 bool
					_ = _hx_tmp_12
					if closers.Len() == 0 {
						_hx_tmp_12 = true
					} else {
						var hx_cmp_0 HxNullable[int32]
						_ = hx_cmp_0
						hx_cmp_0 = (Hx_Array_Pop[int32](closers))
						var _hx_tmp_13 bool
						_ = _hx_tmp_13
						if (hx_cmp_0).Valid == false {
							_hx_tmp_13 = true
						} else {
							_hx_tmp_13 = ((hx_cmp_0).Value != c)
						}

						_hx_tmp_12 = (_hx_tmp_13)
					}

					if _hx_tmp_12 {
						return -1
					}

					if closers.Len() == 0 {
						return (pos + 1)
					}
				}
			}
		}

		pos++
	}

	return -1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_conditionOperatorEnd(pos int32) int32 {
	{
		{
			var word string
			_ = word
			word = "and"
			var end int32
			_ = end
			end = this.VTable.Hx_Field_wordEnd(pos, word)
			if end != -1 {
				return end
			}
		}

		{
			var word1 string
			_ = word1
			word1 = "or"
			var end1 int32
			_ = end1
			end1 = this.VTable.Hx_Field_wordEnd(pos, word1)
			if end1 != -1 {
				return end1
			}
		}

		{
			var word2 string
			_ = word2
			word2 = "is"
			var end2 int32
			_ = end2
			end2 = this.VTable.Hx_Field_wordEnd(pos, word2)
			if end2 != -1 {
				return end2
			}
		}
	}

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
	var next int32
	_ = next
	if (pos + 1) < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + 1)
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if pos1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
		}

		next = _hx_tmp_4
	} else {
		next = 0
	}

	if (((c == 38) && (next == 38)) || ((c == 124) && (next == 124))) || (((((c == 61) || (c == 33)) || (c == 60)) || (c == 62)) && (next == 61)) {
		return (pos + 2)
	}

	if ((((c == 60) || (c == 62)) || (c == 43)) || (c == 42)) || (c == 37) {
		return (pos + 1)
	}

	if ((c == 47) && (next != 47)) && (next != 42) {
		return (pos + 1)
	}

	if (c == 45) && (next != 62) {
		return (pos + 1)
	}

	return -1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isIdentifierExpressionStart(pos int32, lowercaseIdentOnly bool) bool {
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	if pos >= this.Hx_Field_length {
		return false
	}

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

	if (_hx_tmp_0) == 40 {
		return true
	}

	var _gthis *Hx_Obj_loreline_lexer
	_ = _gthis
	_gthis = this
	var result bool
	_ = result
	result = true
	if pos >= _gthis.Hx_Field_length {
		result = false
	} else {
		var isUnderscore bool
		_ = isUnderscore
		isUnderscore = false
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = _gthis.Hx_Field_input
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

		var c int32
		_ = c
		c = _hx_tmp_4
		isUnderscore = (c == 95)
		if !((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((lowercaseIdentOnly && !isUnderscore) && !((c >= 97) && (c <= 122))) {
			result = false
		} else {
			pos++
			for pos < _gthis.Hx_Field_length {
				var wasUnderscore bool
				_ = wasUnderscore
				wasUnderscore = isUnderscore
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = _gthis.Hx_Field_input
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
					_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_9 {
					_hx_tmp_8 = -1
				} else {
					_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				c = _hx_tmp_8
				isUnderscore = (c == 95)
				if !(((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57))) {
					break
				}

				if ((lowercaseIdentOnly && wasUnderscore) && !isUnderscore) && !((c >= 97) && (c <= 122)) {
					result = false
					break
				}

				pos++
			}
		}
	}

	if !result {
		return false
	}

	for pos < this.Hx_Field_length {
		pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
		if pos >= this.Hx_Field_length {
			return true
		}

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

		var c int32
		_ = c
		c = _hx_tmp_12
		if ((((!((c == 32) || (c == 9)) && (c != 46)) && (c != 91)) && (c != 10)) && (c != 13)) && (c != 47) {
			return false
		}

		var tmp bool
		_ = tmp
		if !((c == 10) || (c == 13)) {
			if (c == 47) && ((pos + 1) < this.Hx_Field_length) {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_16 int32
				_ = _hx_tmp_16
				var _hx_tmp_17 bool
				_ = _hx_tmp_17
				if pos1 < 0 {
					_hx_tmp_17 = true
				} else {
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					_hx_tmp_18 = pos1
					var _hx_tmp_19 int32
					_ = _hx_tmp_19
					_hx_tmp_19 = _hx_tmp_18
					_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_17 {
					_hx_tmp_16 = -1
				} else {
					_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos1))))
				}

				if (_hx_tmp_16) != 47 {
					var tmp_v_3 HxArray[int32]
					_ = tmp_v_3
					tmp_v_3 = this.Hx_Field_input
					var tmp_pos1_1 int32
					_ = tmp_pos1_1
					tmp_pos1_1 = (pos + 1)
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					var _hx_tmp_21 bool
					_ = _hx_tmp_21
					if tmp_pos1_1 < 0 {
						_hx_tmp_21 = true
					} else {
						var _hx_tmp_22 int32
						_ = _hx_tmp_22
						_hx_tmp_22 = tmp_pos1_1
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = _hx_tmp_22
						_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_3).Len())
					}

					if _hx_tmp_21 {
						_hx_tmp_20 = -1
					} else {
						_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos1_1))))
					}

					tmp = ((_hx_tmp_20) == 42)
				} else {
					tmp = true
				}
			} else {
				tmp = false
			}
		} else {
			tmp = true
		}

		if tmp {
			return true
		}

		if c == 46 {
			pos++
			pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
			if pos >= this.Hx_Field_length {
				return true
			}

			var lowercaseIdentOnly bool
			_ = lowercaseIdentOnly
			lowercaseIdentOnly = false
			var tmp_result_1 bool
			_ = tmp_result_1
			tmp_result_1 = true
			if pos >= _gthis.Hx_Field_length {
				tmp_result_1 = false
			} else {
				var isUnderscore bool
				_ = isUnderscore
				isUnderscore = false
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = _gthis.Hx_Field_input
				var _hx_tmp_24 int32
				_ = _hx_tmp_24
				var _hx_tmp_25 bool
				_ = _hx_tmp_25
				if pos < 0 {
					_hx_tmp_25 = true
				} else {
					var _hx_tmp_26 int32
					_ = _hx_tmp_26
					_hx_tmp_26 = pos
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					_hx_tmp_27 = _hx_tmp_26
					_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_25 {
					_hx_tmp_24 = -1
				} else {
					_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				var tmp_c_1 int32
				_ = tmp_c_1
				tmp_c_1 = _hx_tmp_24
				isUnderscore = (tmp_c_1 == 95)
				if !((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((lowercaseIdentOnly && !isUnderscore) && !((tmp_c_1 >= 97) && (tmp_c_1 <= 122))) {
					tmp_result_1 = false
				} else {
					pos++
					for pos < _gthis.Hx_Field_length {
						var wasUnderscore bool
						_ = wasUnderscore
						wasUnderscore = isUnderscore
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = _gthis.Hx_Field_input
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

						tmp_c_1 = _hx_tmp_28
						isUnderscore = (tmp_c_1 == 95)
						if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
							break
						}

						if ((lowercaseIdentOnly && wasUnderscore) && !isUnderscore) && !((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) {
							tmp_result_1 = false
							break
						}

						pos++
					}
				}
			}

			if !tmp_result_1 {
				return true
			}

			continue
		}

		if c == 91 {
			pos++
			var bracketLevel int32
			_ = bracketLevel
			bracketLevel = 1
			for (pos < this.Hx_Field_length) && (bracketLevel > 0) {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var _hx_tmp_32 int32
				_ = _hx_tmp_32
				var _hx_tmp_33 bool
				_ = _hx_tmp_33
				if pos < 0 {
					_hx_tmp_33 = true
				} else {
					var _hx_tmp_34 int32
					_ = _hx_tmp_34
					_hx_tmp_34 = pos
					var _hx_tmp_35 int32
					_ = _hx_tmp_35
					_hx_tmp_35 = _hx_tmp_34
					_hx_tmp_33 = (_hx_tmp_35 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_33 {
					_hx_tmp_32 = -1
				} else {
					_hx_tmp_32 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				c = _hx_tmp_32
				if c == 34 {
					pos = this.VTable.Hx_Field_scanStringEnd(pos, false)
					if pos == -1 {
						return false
					}

					continue
				}

				if c == 91 {
					bracketLevel++
				}

				if c == 93 {
					bracketLevel--
				}

				pos++
			}

			continue
		}

		if (c == 32) || (c == 9) {
			pos++
			continue
		}

		break
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isTransitionStart(pos int32) bool {
	var tmp bool
	_ = tmp
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if pos < 0 {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = pos
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = _hx_tmp_3
		_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = -1
	} else {
		_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	if (_hx_tmp_1) != 45 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (pos >= (this.Hx_Field_length - 1))
	}

	if !(_hx_tmp_0) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + 1)
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos1 < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos1
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
		}

		tmp = ((_hx_tmp_5) != 62)
	} else {
		tmp = true
	}

	if tmp {
		return false
	}

	pos = (pos + 2)
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	if pos >= this.Hx_Field_length {
		return false
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
	var _hx_tmp_9 int32
	_ = _hx_tmp_9
	var _hx_tmp_10 bool
	_ = _hx_tmp_10
	if pos < 0 {
		_hx_tmp_10 = true
	} else {
		var _hx_tmp_11 int32
		_ = _hx_tmp_11
		_hx_tmp_11 = pos
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		_hx_tmp_12 = _hx_tmp_11
		_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](v1).Len())
	}

	if _hx_tmp_10 {
		_hx_tmp_9 = -1
	} else {
		_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos))))
	}

	var char int32
	_ = char
	char = _hx_tmp_9
	if char == 46 {
		pos++
	} else {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		var _hx_tmp_14 bool
		_ = _hx_tmp_14
		if pos < 0 {
			_hx_tmp_14 = true
		} else {
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = pos
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = _hx_tmp_15
			_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_14 {
			_hx_tmp_13 = -1
		} else {
			_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_13
		if ((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)) {
			pos++
			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if pos < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					var _hx_tmp_18 bool
					_ = _hx_tmp_18
					if pos < 0 {
						_hx_tmp_18 = true
					} else {
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = pos
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						_hx_tmp_20 = _hx_tmp_19
						_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_18 {
						_hx_tmp_17 = -1
					} else {
						_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_17
					tmp_tmp_1 = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				pos++
			}

			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if pos < this.Hx_Field_length {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var _hx_tmp_21 int32
				_ = _hx_tmp_21
				var _hx_tmp_22 bool
				_ = _hx_tmp_22
				if pos < 0 {
					_hx_tmp_22 = true
				} else {
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = pos
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = _hx_tmp_23
					_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_22 {
					_hx_tmp_21 = -1
				} else {
					_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				tmp_tmp_1 = ((_hx_tmp_21) == 40)
			} else {
				tmp_tmp_1 = false
			}

			if tmp_tmp_1 {
				var _hx_tmp_25 int32
				_ = _hx_tmp_25
				_hx_tmp_25 = pos
				pos = this.VTable.Hx_Field_skipBalancedParens(_hx_tmp_25, (this.VTable.Hx_Field_parentBlockType() == Hx_Enum_loreline_tokentype_KwBeat{}))
				if pos == -1 {
					return false
				}
			}
		} else {
			return false
		}
	}

	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	if pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
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

		var c int32
		_ = c
		c = _hx_tmp_26
		if (((((c != 10) && (c != 13)) && (c != 32)) && (c != 9)) && (c != 47)) && (c != 35) {
			return false
		}
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isInsertionStart(pos int32) bool {
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

	if (_hx_tmp_0) != 43 {
		return false
	}

	pos++
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	if pos >= this.Hx_Field_length {
		return false
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
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
		_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v1).Len())
	}

	if _hx_tmp_5 {
		_hx_tmp_4 = -1
	} else {
		_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos))))
	}

	var char int32
	_ = char
	char = _hx_tmp_4
	if char == 46 {
		pos++
	} else {
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

		var c int32
		_ = c
		c = _hx_tmp_8
		if ((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)) {
			pos++
			for {
				var tmp bool
				_ = tmp
				if pos < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
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
						_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_13 {
						_hx_tmp_12 = -1
					} else {
						_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_12
					tmp = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
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
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
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
					_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_17 {
					_hx_tmp_16 = -1
				} else {
					_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
				}

				tmp = ((_hx_tmp_16) == 40)
			} else {
				tmp = false
			}

			if tmp {
				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				_hx_tmp_20 = pos
				pos = this.VTable.Hx_Field_skipBalancedParens(_hx_tmp_20, (this.VTable.Hx_Field_parentBlockType() == Hx_Enum_loreline_tokentype_KwBeat{}))
				if pos == -1 {
					return false
				}
			}
		} else {
			return false
		}
	}

	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	if pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_21 int32
		_ = _hx_tmp_21
		var _hx_tmp_22 bool
		_ = _hx_tmp_22
		if pos < 0 {
			_hx_tmp_22 = true
		} else {
			var _hx_tmp_23 int32
			_ = _hx_tmp_23
			_hx_tmp_23 = pos
			var _hx_tmp_24 int32
			_ = _hx_tmp_24
			_hx_tmp_24 = _hx_tmp_23
			_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_22 {
			_hx_tmp_21 = -1
		} else {
			_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_21
		var tmp bool
		_ = tmp
		if (c == 105) && ((pos + 1) < this.Hx_Field_length) {
			var tmp_v_2 HxArray[int32]
			_ = tmp_v_2
			tmp_v_2 = this.Hx_Field_input
			var pos1 int32
			_ = pos1
			pos1 = (pos + 1)
			var _hx_tmp_25 int32
			_ = _hx_tmp_25
			var _hx_tmp_26 bool
			_ = _hx_tmp_26
			if pos1 < 0 {
				_hx_tmp_26 = true
			} else {
				var _hx_tmp_27 int32
				_ = _hx_tmp_27
				_hx_tmp_27 = pos1
				var _hx_tmp_28 int32
				_ = _hx_tmp_28
				_hx_tmp_28 = _hx_tmp_27
				_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](tmp_v_2).Len())
			}

			if _hx_tmp_26 {
				_hx_tmp_25 = -1
			} else {
				_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos1))))
			}

			tmp = ((_hx_tmp_25) == 102)
		} else {
			tmp = false
		}

		if tmp {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if (pos + 2) < this.Hx_Field_length {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 2)
				var _hx_tmp_29 int32
				_ = _hx_tmp_29
				var _hx_tmp_30 bool
				_ = _hx_tmp_30
				if pos1 < 0 {
					_hx_tmp_30 = true
				} else {
					var _hx_tmp_31 int32
					_ = _hx_tmp_31
					_hx_tmp_31 = pos1
					var _hx_tmp_32 int32
					_ = _hx_tmp_32
					_hx_tmp_32 = _hx_tmp_31
					_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_30 {
					_hx_tmp_29 = -1
				} else {
					_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos1))))
				}

				var tmp_c_1 int32
				_ = tmp_c_1
				tmp_c_1 = _hx_tmp_29
				tmp_tmp_1 = !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
			} else {
				tmp_tmp_1 = true
			}

			if tmp_tmp_1 {
				return true
			}
		}

		if ((((c != 10) && (c != 13)) && (c != 32)) && (c != 9)) && (c != 47) {
			return false
		}
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isChoiceOnceStart(pos int32) bool {
	if !this.VTable.Hx_Field_inChoiceRoot() {
		return false
	}

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

	if (_hx_tmp_0) != 45 {
		return false
	}

	pos++
	if pos >= this.Hx_Field_length {
		return false
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
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
		_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v1).Len())
	}

	if _hx_tmp_5 {
		_hx_tmp_4 = -1
	} else {
		_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos))))
	}

	if (_hx_tmp_4) != 32 {
		return false
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isLabelStart(pos int32) bool {
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
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

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_8
			tmp = ((tmp_c_1 == 32) || (tmp_c_1 == 9))
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
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_136 func(_hx_param_key any) bool
		_ = _hx_closure_src_136
		_hx_closure_src_136 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_136(((any)(_hx_param_key)))
		}
	}()(word) {
		return false
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isCallStart(pos int32) bool {
	var startPos int32
	_ = startPos
	startPos = pos
	var _gthis *Hx_Obj_loreline_lexer
	_ = _gthis
	_gthis = this
	var result bool
	_ = result
	result = true
	if pos >= _gthis.Hx_Field_length {
		result = false
	} else {
		var v HxArray[int32]
		_ = v
		v = _gthis.Hx_Field_input
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
			result = false
		} else {
			pos++
			for pos < _gthis.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = _gthis.Hx_Field_input
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

				c = _hx_tmp_4
				if !(((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57))) {
					break
				}

				pos++
			}
		}
	}

	if !result {
		pos = startPos
		return false
	}

	for pos < this.Hx_Field_length {
		pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
		if pos >= this.Hx_Field_length {
			pos = startPos
			return false
		}

		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_9 {
			_hx_tmp_8 = -1
		} else {
			_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_8
		if c == 40 {
			pos = startPos
			return true
		}

		if c == 46 {
			pos++
			pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
			var tmp_result_1 bool
			_ = tmp_result_1
			tmp_result_1 = true
			if pos >= _gthis.Hx_Field_length {
				tmp_result_1 = false
			} else {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = _gthis.Hx_Field_input
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

				var tmp_c_1 int32
				_ = tmp_c_1
				tmp_c_1 = _hx_tmp_12
				if !((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) {
					tmp_result_1 = false
				} else {
					pos++
					for pos < _gthis.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = _gthis.Hx_Field_input
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
							_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_17 {
							_hx_tmp_16 = -1
						} else {
							_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
						}

						tmp_c_1 = _hx_tmp_16
						if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
							break
						}

						pos++
					}
				}
			}

			if !tmp_result_1 {
				pos = startPos
				return false
			}

			continue
		}

		if c == 91 {
			pos++
			for pos < this.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				var _hx_tmp_21 bool
				_ = _hx_tmp_21
				if pos < 0 {
					_hx_tmp_21 = true
				} else {
					var _hx_tmp_22 int32
					_ = _hx_tmp_22
					_hx_tmp_22 = pos
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = _hx_tmp_22
					_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_21 {
					_hx_tmp_20 = -1
				} else {
					_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
				}

				var bc int32
				_ = bc
				bc = _hx_tmp_20
				if bc == 34 {
					pos = this.VTable.Hx_Field_scanStringEnd(pos, false)
					if pos == -1 {
						pos = startPos
						return false
					}

					continue
				}

				if bc == 93 {
					pos++
					break
				}

				pos++
			}

			continue
		}

		pos = startPos
		return false
	}

	pos = startPos
	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isAssignStart(pos int32, strict bool) bool {
	var _gthis *Hx_Obj_loreline_lexer
	_ = _gthis
	_gthis = this
	var tmp bool
	_ = tmp
	if strict {
		var result bool
		_ = result
		result = true
		var startPos int32
		_ = startPos
		startPos = pos
		if pos >= _gthis.Hx_Field_length {
			result = false
		} else {
			var v HxArray[int32]
			_ = v
			v = _gthis.Hx_Field_input
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
				result = false
			} else {
				pos++
				for pos < _gthis.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = _gthis.Hx_Field_input
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

					c = _hx_tmp_4
					if !(((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57))) {
						break
					}

					pos++
				}
			}
		}

		var tmp1 bool
		_ = tmp1
		var tmp2 bool
		_ = tmp2
		if pos == (startPos + 2) {
			var v HxArray[int32]
			_ = v
			v = _gthis.Hx_Field_input
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			var _hx_tmp_9 bool
			_ = _hx_tmp_9
			if startPos < 0 {
				_hx_tmp_9 = true
			} else {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = startPos
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_10
				_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_9 {
				_hx_tmp_8 = -1
			} else {
				_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(startPos))))
			}

			tmp2 = ((_hx_tmp_8) == 105)
		} else {
			tmp2 = false
		}

		if tmp2 {
			var v HxArray[int32]
			_ = v
			v = _gthis.Hx_Field_input
			var pos int32
			_ = pos
			pos = (startPos + 1)
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
				_hx_tmp_13 = (_hx_tmp_15 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_13 {
				_hx_tmp_12 = -1
			} else {
				_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			tmp1 = ((_hx_tmp_12) == 102)
		} else {
			tmp1 = false
		}

		var _hx_tmp_16 bool
		_ = _hx_tmp_16
		if tmp1 {
			_hx_tmp_16 = false
		} else {
			_hx_tmp_16 = result
		}

		tmp = !(_hx_tmp_16)
	} else {
		tmp = false
	}

	if tmp {
		return false
	}

	var isEscape bool
	_ = isEscape
	isEscape = false
	for pos < this.Hx_Field_length {
		pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
		if pos >= this.Hx_Field_length {
			return false
		}

		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_17 int32
		_ = _hx_tmp_17
		var _hx_tmp_18 bool
		_ = _hx_tmp_18
		if pos < 0 {
			_hx_tmp_18 = true
		} else {
			var _hx_tmp_19 int32
			_ = _hx_tmp_19
			_hx_tmp_19 = pos
			var _hx_tmp_20 int32
			_ = _hx_tmp_20
			_hx_tmp_20 = _hx_tmp_19
			_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_18 {
			_hx_tmp_17 = -1
		} else {
			_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_17
		var tmp_tmp_1 bool
		_ = tmp_tmp_1
		if !isEscape {
			var tmp1 bool
			_ = tmp1
			if c == 61 {
				if (pos + 1) < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var pos1 int32
					_ = pos1
					pos1 = (pos + 1)
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					var _hx_tmp_22 bool
					_ = _hx_tmp_22
					if pos1 < 0 {
						_hx_tmp_22 = true
					} else {
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = pos1
						var _hx_tmp_24 int32
						_ = _hx_tmp_24
						_hx_tmp_24 = _hx_tmp_23
						_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_22 {
						_hx_tmp_21 = -1
					} else {
						_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
					}

					tmp1 = ((_hx_tmp_21) != 61)
				} else {
					tmp1 = true
				}
			} else {
				tmp1 = false
			}

			if !tmp1 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var pos1 int32
				_ = pos1
				pos1 = (pos + 1)
				var _hx_tmp_25 bool
				_ = _hx_tmp_25
				var _hx_tmp_26 int32
				_ = _hx_tmp_26
				var _hx_tmp_27 bool
				_ = _hx_tmp_27
				if pos1 < 0 {
					_hx_tmp_27 = true
				} else {
					var _hx_tmp_28 int32
					_ = _hx_tmp_28
					_hx_tmp_28 = pos1
					var _hx_tmp_29 int32
					_ = _hx_tmp_29
					_hx_tmp_29 = _hx_tmp_28
					_hx_tmp_27 = (_hx_tmp_29 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_27 {
					_hx_tmp_26 = -1
				} else {
					_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
				}

				if (_hx_tmp_26) == 61 {
					_hx_tmp_25 = (((((c == 43) || (c == 45)) || (c == 42)) || (c == 47)) || (c == 58))
				} else {
					_hx_tmp_25 = false
				}

				tmp_tmp_1 = _hx_tmp_25
			} else {
				tmp_tmp_1 = true
			}
		} else {
			tmp_tmp_1 = false
		}

		if tmp_tmp_1 {
			return true
		}

		if strict {
			if c == 46 {
				pos++
				pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
				var result bool
				_ = result
				result = true
				var startPos int32
				_ = startPos
				startPos = pos
				if pos >= _gthis.Hx_Field_length {
					result = false
				} else {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = _gthis.Hx_Field_input
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

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_30
					if !((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) {
						result = false
					} else {
						pos++
						for pos < _gthis.Hx_Field_length {
							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = _gthis.Hx_Field_input
							var _hx_tmp_34 int32
							_ = _hx_tmp_34
							var _hx_tmp_35 bool
							_ = _hx_tmp_35
							if pos < 0 {
								_hx_tmp_35 = true
							} else {
								var _hx_tmp_36 int32
								_ = _hx_tmp_36
								_hx_tmp_36 = pos
								var _hx_tmp_37 int32
								_ = _hx_tmp_37
								_hx_tmp_37 = _hx_tmp_36
								_hx_tmp_35 = (_hx_tmp_37 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_35 {
								_hx_tmp_34 = -1
							} else {
								_hx_tmp_34 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
							}

							tmp_c_1 = _hx_tmp_34
							if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
								break
							}

							pos++
						}
					}
				}

				var tmp_tmp_2 bool
				_ = tmp_tmp_2
				var tmp1 bool
				_ = tmp1
				if pos == (startPos + 2) {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = _gthis.Hx_Field_input
					var _hx_tmp_38 int32
					_ = _hx_tmp_38
					var _hx_tmp_39 bool
					_ = _hx_tmp_39
					if startPos < 0 {
						_hx_tmp_39 = true
					} else {
						var _hx_tmp_40 int32
						_ = _hx_tmp_40
						_hx_tmp_40 = startPos
						var _hx_tmp_41 int32
						_ = _hx_tmp_41
						_hx_tmp_41 = _hx_tmp_40
						_hx_tmp_39 = (_hx_tmp_41 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_39 {
						_hx_tmp_38 = -1
					} else {
						_hx_tmp_38 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(startPos))))
					}

					tmp1 = ((_hx_tmp_38) == 105)
				} else {
					tmp1 = false
				}

				if tmp1 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = _gthis.Hx_Field_input
					var pos int32
					_ = pos
					pos = (startPos + 1)
					var _hx_tmp_42 int32
					_ = _hx_tmp_42
					var _hx_tmp_43 bool
					_ = _hx_tmp_43
					if pos < 0 {
						_hx_tmp_43 = true
					} else {
						var _hx_tmp_44 int32
						_ = _hx_tmp_44
						_hx_tmp_44 = pos
						var _hx_tmp_45 int32
						_ = _hx_tmp_45
						_hx_tmp_45 = _hx_tmp_44
						_hx_tmp_43 = (_hx_tmp_45 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_43 {
						_hx_tmp_42 = -1
					} else {
						_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					tmp_tmp_2 = ((_hx_tmp_42) == 102)
				} else {
					tmp_tmp_2 = false
				}

				var _hx_tmp_46 bool
				_ = _hx_tmp_46
				if tmp_tmp_2 {
					_hx_tmp_46 = false
				} else {
					_hx_tmp_46 = result
				}

				if !(_hx_tmp_46) {
					return false
				}

				continue
			}

			if c == 91 {
				pos++
				for pos < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_47 int32
					_ = _hx_tmp_47
					var _hx_tmp_48 bool
					_ = _hx_tmp_48
					if pos < 0 {
						_hx_tmp_48 = true
					} else {
						var _hx_tmp_49 int32
						_ = _hx_tmp_49
						_hx_tmp_49 = pos
						var _hx_tmp_50 int32
						_ = _hx_tmp_50
						_hx_tmp_50 = _hx_tmp_49
						_hx_tmp_48 = (_hx_tmp_50 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_48 {
						_hx_tmp_47 = -1
					} else {
						_hx_tmp_47 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
					}

					var bc int32
					_ = bc
					bc = _hx_tmp_47
					if bc == 34 {
						pos = this.VTable.Hx_Field_scanStringEnd(pos, false)
						if pos == -1 {
							return false
						}

						continue
					}

					if bc == 93 {
						pos++
						break
					}

					pos++
				}

				continue
			}

			return false
		} else {
			if c == 92 {
				isEscape = true
				pos++
			} else {
				if (c == 13) || (c == 10) {
					return false
				} else {
					if c == 34 {
						pos = this.VTable.Hx_Field_scanStringEnd(pos, false)
						if pos == -1 {
							return false
						}
					} else {
						if (((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95) {
							var result bool
							_ = result
							result = true
							var startPos int32
							_ = startPos
							startPos = pos
							if pos >= _gthis.Hx_Field_length {
								result = false
							} else {
								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = _gthis.Hx_Field_input
								var _hx_tmp_51 int32
								_ = _hx_tmp_51
								var _hx_tmp_52 bool
								_ = _hx_tmp_52
								if pos < 0 {
									_hx_tmp_52 = true
								} else {
									var _hx_tmp_53 int32
									_ = _hx_tmp_53
									_hx_tmp_53 = pos
									var _hx_tmp_54 int32
									_ = _hx_tmp_54
									_hx_tmp_54 = _hx_tmp_53
									_hx_tmp_52 = (_hx_tmp_54 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_52 {
									_hx_tmp_51 = -1
								} else {
									_hx_tmp_51 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
								}

								var tmp_c_1 int32
								_ = tmp_c_1
								tmp_c_1 = _hx_tmp_51
								if !((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) {
									result = false
								} else {
									pos++
									for pos < _gthis.Hx_Field_length {
										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = _gthis.Hx_Field_input
										var _hx_tmp_55 int32
										_ = _hx_tmp_55
										var _hx_tmp_56 bool
										_ = _hx_tmp_56
										if pos < 0 {
											_hx_tmp_56 = true
										} else {
											var _hx_tmp_57 int32
											_ = _hx_tmp_57
											_hx_tmp_57 = pos
											var _hx_tmp_58 int32
											_ = _hx_tmp_58
											_hx_tmp_58 = _hx_tmp_57
											_hx_tmp_56 = (_hx_tmp_58 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_56 {
											_hx_tmp_55 = -1
										} else {
											_hx_tmp_55 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
										}

										tmp_c_1 = _hx_tmp_55
										if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
											break
										}

										pos++
									}
								}
							}

							var tmp_tmp_2 bool
							_ = tmp_tmp_2
							var tmp1 bool
							_ = tmp1
							if pos == (startPos + 2) {
								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = _gthis.Hx_Field_input
								var _hx_tmp_59 int32
								_ = _hx_tmp_59
								var _hx_tmp_60 bool
								_ = _hx_tmp_60
								if startPos < 0 {
									_hx_tmp_60 = true
								} else {
									var _hx_tmp_61 int32
									_ = _hx_tmp_61
									_hx_tmp_61 = startPos
									var _hx_tmp_62 int32
									_ = _hx_tmp_62
									_hx_tmp_62 = _hx_tmp_61
									_hx_tmp_60 = (_hx_tmp_62 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_60 {
									_hx_tmp_59 = -1
								} else {
									_hx_tmp_59 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(startPos))))
								}

								tmp1 = ((_hx_tmp_59) == 105)
							} else {
								tmp1 = false
							}

							if tmp1 {
								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = _gthis.Hx_Field_input
								var pos int32
								_ = pos
								pos = (startPos + 1)
								var _hx_tmp_63 int32
								_ = _hx_tmp_63
								var _hx_tmp_64 bool
								_ = _hx_tmp_64
								if pos < 0 {
									_hx_tmp_64 = true
								} else {
									var _hx_tmp_65 int32
									_ = _hx_tmp_65
									_hx_tmp_65 = pos
									var _hx_tmp_66 int32
									_ = _hx_tmp_66
									_hx_tmp_66 = _hx_tmp_65
									_hx_tmp_64 = (_hx_tmp_66 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_64 {
									_hx_tmp_63 = -1
								} else {
									_hx_tmp_63 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
								}

								tmp_tmp_2 = ((_hx_tmp_63) == 102)
							} else {
								tmp_tmp_2 = false
							}

							var _hx_tmp_67 bool
							_ = _hx_tmp_67
							if tmp_tmp_2 {
								_hx_tmp_67 = false
							} else {
								_hx_tmp_67 = result
							}

							if !(_hx_tmp_67) {
								return false
							}
						} else {
							var _hx_tmp_68 bool
							_ = _hx_tmp_68
							if c == 60 {
								_hx_tmp_68 = (this.VTable.Hx_Field_scanTagEnd(pos) != -1)
							} else {
								_hx_tmp_68 = false
							}

							if _hx_tmp_68 {
								return false
							} else {
								var tmp_tmp_2 bool
								_ = tmp_tmp_2
								if ((((c == 61) || (c == 33)) || (c == 60)) || (c == 62)) && ((pos + 1) < this.Hx_Field_length) {
									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var pos1 int32
									_ = pos1
									pos1 = (pos + 1)
									var _hx_tmp_69 int32
									_ = _hx_tmp_69
									var _hx_tmp_70 bool
									_ = _hx_tmp_70
									if pos1 < 0 {
										_hx_tmp_70 = true
									} else {
										var _hx_tmp_71 int32
										_ = _hx_tmp_71
										_hx_tmp_71 = pos1
										var _hx_tmp_72 int32
										_ = _hx_tmp_72
										_hx_tmp_72 = _hx_tmp_71
										_hx_tmp_70 = (_hx_tmp_72 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_70 {
										_hx_tmp_69 = -1
									} else {
										_hx_tmp_69 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
									}

									tmp_tmp_2 = ((_hx_tmp_69) == 61)
								} else {
									tmp_tmp_2 = false
								}

								if tmp_tmp_2 {
									pos = (pos + 2)
								} else {
									pos++
								}
							}
						}
					}
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isColon(pos int32, skipWhitespaces__isOptional any) bool {
	skipWhitespaces := HxDefault[bool]()
	_ = skipWhitespaces

	if skipWhitespaces__isOptional != nil {
		skipWhitespaces = HxConvert[bool](skipWhitespaces__isOptional)
	} else {
		skipWhitespaces = true
	}
	{
		if skipWhitespaces {
			for {
				var tmp bool
				_ = tmp
				if pos < this.Hx_Field_length {
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

				pos++
			}
		}

		if pos < this.Hx_Field_length {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
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
				_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_9 {
				_hx_tmp_8 = -1
			} else {
				_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			return ((_hx_tmp_8) == 58)
		} else {
			return false
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_afterLabelIdentifierToken(inSameLine__isOptional any) int32 {
	inSameLine := HxDefault[bool]()
	_ = inSameLine

	if inSameLine__isOptional != nil {
		inSameLine = HxConvert[bool](inSameLine__isOptional)
	} else {
		inSameLine = true
	}
	{
		var i int32
		_ = i
		i = (this.Hx_Field_tokenized.Len() - 1)
		for i >= 0 {
			var token *Hx_Obj_loreline_token
			_ = token
			token = this.Hx_Field_tokenized.Get(i)
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if !Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
				_hx_tmp_0 = (inSameLine || (((token.Hx_Field_type != Hx_Enum_loreline_tokentype_LineBreak{}) && (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Indent{})) && (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Unindent{})))
			} else {
				_hx_tmp_0 = false
			}

			if _hx_tmp_0 {
				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) && (i > 0) {
					_hx_tmp_1 = Hx_Field_loreline_tokentypehelpers_isIdentifier(this.Hx_Field_tokenized.Get((i - 1)).Hx_Field_type)
				} else {
					_hx_tmp_1 = false
				}

				if _hx_tmp_1 {
					return (i - 1)
				}

				return -1
			}

			i--
		}

		return -1
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isStartingMultilineDialogue(labelIdentifierTokenIndex int32) int32 {
	var labelIdentifierToken *Hx_Obj_loreline_token
	_ = labelIdentifierToken
	labelIdentifierToken = this.Hx_Field_tokenized.Get(labelIdentifierTokenIndex)
	var i int32
	_ = i
	i = (labelIdentifierTokenIndex + 1)
	for (this.Hx_Field_tokenized.Get(i).Hx_Field_type != Hx_Enum_loreline_tokentype_Colon{}) {
		i++
	}

	var pos int32
	_ = pos
	pos = (this.Hx_Field_tokenized.Get(i).Hx_Field_pos.Hx_Field_offset + 1)
	var labelColumn int32
	_ = labelColumn
	labelColumn = labelIdentifierToken.Hx_Field_pos.Hx_Field_column
	var onlyTagsOnFirstLine bool
	_ = onlyTagsOnFirstLine
	onlyTagsOnFirstLine = true
	var insideTag bool
	_ = insideTag
	insideTag = false
	for pos < this.Hx_Field_length {
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
		if (c == 32) || (c == 9) {
			pos++
		} else {
			break
		}
	}

	for pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_4
		if (c == 10) || (c == 13) {
			break
		}

		if (c == 47) && ((pos + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos1 int32
			_ = pos1
			pos1 = (pos + 1)
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			var _hx_tmp_9 bool
			_ = _hx_tmp_9
			if pos1 < 0 {
				_hx_tmp_9 = true
			} else {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = pos1
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_10
				_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_9 {
				_hx_tmp_8 = -1
			} else {
				_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
			}

			var next int32
			_ = next
			next = _hx_tmp_8
			if next == 47 {
				break
			} else {
				if next == 42 {
					pos = (pos + 2)
					for pos < this.Hx_Field_length {
						var tmp bool
						_ = tmp
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_12 bool
						_ = _hx_tmp_12
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						var _hx_tmp_14 bool
						_ = _hx_tmp_14
						if pos < 0 {
							_hx_tmp_14 = true
						} else {
							var _hx_tmp_15 int32
							_ = _hx_tmp_15
							_hx_tmp_15 = pos
							var _hx_tmp_16 int32
							_ = _hx_tmp_16
							_hx_tmp_16 = _hx_tmp_15
							_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_14 {
							_hx_tmp_13 = -1
						} else {
							_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
						}

						if (_hx_tmp_13) == 42 {
							_hx_tmp_12 = ((pos + 1) < this.Hx_Field_length)
						} else {
							_hx_tmp_12 = false
						}

						if _hx_tmp_12 {
							var tmp_v_3 HxArray[int32]
							_ = tmp_v_3
							tmp_v_3 = this.Hx_Field_input
							var tmp_pos1_1 int32
							_ = tmp_pos1_1
							tmp_pos1_1 = (pos + 1)
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							var _hx_tmp_18 bool
							_ = _hx_tmp_18
							if tmp_pos1_1 < 0 {
								_hx_tmp_18 = true
							} else {
								var _hx_tmp_19 int32
								_ = _hx_tmp_19
								_hx_tmp_19 = tmp_pos1_1
								var _hx_tmp_20 int32
								_ = _hx_tmp_20
								_hx_tmp_20 = _hx_tmp_19
								_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_3).Len())
							}

							if _hx_tmp_18 {
								_hx_tmp_17 = -1
							} else {
								_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos1_1))))
							}

							tmp = ((_hx_tmp_17) == 47)
						} else {
							tmp = false
						}

						if tmp {
							pos = (pos + 2)
							break
						}

						pos++
					}

					continue
				}
			}
		}

		if c == 60 {
			insideTag = true
		} else {
			if c == 62 {
				insideTag = false
			} else {
				if !insideTag && !((c == 32) || (c == 9)) {
					onlyTagsOnFirstLine = false
					break
				}
			}
		}

		pos++
	}

	if !onlyTagsOnFirstLine {
		for {
			var tmp bool
			_ = tmp
			var tmp1 bool
			_ = tmp1
			if pos < this.Hx_Field_length {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_21 int32
				_ = _hx_tmp_21
				var _hx_tmp_22 bool
				_ = _hx_tmp_22
				if pos < 0 {
					_hx_tmp_22 = true
				} else {
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = pos
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = _hx_tmp_23
					_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_22 {
					_hx_tmp_21 = -1
				} else {
					_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				tmp1 = ((_hx_tmp_21) != 10)
			} else {
				tmp1 = false
			}

			if tmp1 {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var _hx_tmp_25 int32
				_ = _hx_tmp_25
				var _hx_tmp_26 bool
				_ = _hx_tmp_26
				if pos < 0 {
					_hx_tmp_26 = true
				} else {
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					_hx_tmp_27 = pos
					var _hx_tmp_28 int32
					_ = _hx_tmp_28
					_hx_tmp_28 = _hx_tmp_27
					_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_26 {
					_hx_tmp_25 = -1
				} else {
					_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				tmp = ((_hx_tmp_25) != 13)
			} else {
				tmp = false
			}

			if !tmp {
				break
			}

			pos++
		}
	}

	var prevPos int32
	_ = prevPos
	prevPos = pos
	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, true)
	if pos <= prevPos {
		return -1
	}

	var indent int32
	_ = indent
	indent = 0
	var tmpPos int32
	_ = tmpPos
	tmpPos = pos
	for {
		var tmp bool
		_ = tmp
		if tmpPos > 0 {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = (tmpPos - 1)
			var _hx_tmp_29 int32
			_ = _hx_tmp_29
			var _hx_tmp_30 bool
			_ = _hx_tmp_30
			if tmp_pos_1 < 0 {
				_hx_tmp_30 = true
			} else {
				var _hx_tmp_31 int32
				_ = _hx_tmp_31
				_hx_tmp_31 = tmp_pos_1
				var _hx_tmp_32 int32
				_ = _hx_tmp_32
				_hx_tmp_32 = _hx_tmp_31
				_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_30 {
				_hx_tmp_29 = -1
			} else {
				_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(tmp_pos_1))))
			}

			if (_hx_tmp_29) != 32 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_2 int32
				_ = tmp_pos_2
				tmp_pos_2 = (tmpPos - 1)
				var _hx_tmp_33 int32
				_ = _hx_tmp_33
				var _hx_tmp_34 bool
				_ = _hx_tmp_34
				if tmp_pos_2 < 0 {
					_hx_tmp_34 = true
				} else {
					var _hx_tmp_35 int32
					_ = _hx_tmp_35
					_hx_tmp_35 = tmp_pos_2
					var _hx_tmp_36 int32
					_ = _hx_tmp_36
					_hx_tmp_36 = _hx_tmp_35
					_hx_tmp_34 = (_hx_tmp_36 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_34 {
					_hx_tmp_33 = -1
				} else {
					_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_2))))
				}

				tmp = ((_hx_tmp_33) == 9)
			} else {
				tmp = true
			}
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		indent++
		tmpPos--
	}

	if indent <= (labelColumn - 1) {
		return -1
	}

	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	var tmp bool
	_ = tmp
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_tmp_37 int32
	_ = _hx_tmp_37
	var _hx_tmp_38 bool
	_ = _hx_tmp_38
	if pos < 0 {
		_hx_tmp_38 = true
	} else {
		var _hx_tmp_39 int32
		_ = _hx_tmp_39
		_hx_tmp_39 = pos
		var _hx_tmp_40 int32
		_ = _hx_tmp_40
		_hx_tmp_40 = _hx_tmp_39
		_hx_tmp_38 = (_hx_tmp_40 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_38 {
		_hx_tmp_37 = -1
	} else {
		_hx_tmp_37 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_37
	if !((c == 32) || (c == 9)) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_41 int32
		_ = _hx_tmp_41
		var _hx_tmp_42 bool
		_ = _hx_tmp_42
		if pos < 0 {
			_hx_tmp_42 = true
		} else {
			var _hx_tmp_43 int32
			_ = _hx_tmp_43
			_hx_tmp_43 = pos
			var _hx_tmp_44 int32
			_ = _hx_tmp_44
			_hx_tmp_44 = _hx_tmp_43
			_hx_tmp_42 = (_hx_tmp_44 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_42 {
			_hx_tmp_41 = -1
		} else {
			_hx_tmp_41 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		if (_hx_tmp_41) == 47 {
			var tmp_v_2 HxArray[int32]
			_ = tmp_v_2
			tmp_v_2 = this.Hx_Field_input
			var pos1 int32
			_ = pos1
			pos1 = (pos + 1)
			var _hx_tmp_45 int32
			_ = _hx_tmp_45
			var _hx_tmp_46 bool
			_ = _hx_tmp_46
			if pos1 < 0 {
				_hx_tmp_46 = true
			} else {
				var _hx_tmp_47 int32
				_ = _hx_tmp_47
				_hx_tmp_47 = pos1
				var _hx_tmp_48 int32
				_ = _hx_tmp_48
				_hx_tmp_48 = _hx_tmp_47
				_hx_tmp_46 = (_hx_tmp_48 >= HxMakeArrayView[any](tmp_v_2).Len())
			}

			if _hx_tmp_46 {
				_hx_tmp_45 = -1
			} else {
				_hx_tmp_45 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos1))))
			}

			tmp = ((_hx_tmp_45) == 47)
		} else {
			tmp = false
		}
	} else {
		tmp = true
	}

	if tmp {
		return -1
	}

	return indent
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isContinuingMultilineText(pos int32, indent int32) bool {
	var tmp bool
	_ = tmp
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

	if (_hx_tmp_0) == 47 {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + 1)
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if pos1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos1))))
		}

		tmp = ((_hx_tmp_4) == 47)
	} else {
		tmp = false
	}

	if tmp {
		for {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			var tmp1 bool
			_ = tmp1
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

				tmp1 = ((_hx_tmp_8) != 13)
			} else {
				tmp1 = false
			}

			if tmp1 {
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

				tmp_tmp_1 = ((_hx_tmp_12) != 10)
			} else {
				tmp_tmp_1 = false
			}

			if !tmp_tmp_1 {
				break
			}

			pos++
		}
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
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
		_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](v1).Len())
	}

	if _hx_tmp_17 {
		_hx_tmp_16 = -1
	} else {
		_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos))))
	}

	if (_hx_tmp_16) == 13 {
		pos++
	}

	var v2 HxArray[int32]
	_ = v2
	v2 = this.Hx_Field_input
	var _hx_tmp_20 int32
	_ = _hx_tmp_20
	var _hx_tmp_21 bool
	_ = _hx_tmp_21
	if pos < 0 {
		_hx_tmp_21 = true
	} else {
		var _hx_tmp_22 int32
		_ = _hx_tmp_22
		_hx_tmp_22 = pos
		var _hx_tmp_23 int32
		_ = _hx_tmp_23
		_hx_tmp_23 = _hx_tmp_22
		_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](v2).Len())
	}

	if _hx_tmp_21 {
		_hx_tmp_20 = -1
	} else {
		_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v2).Get(pos))))
	}

	if (_hx_tmp_20) == 10 {
		pos++
	} else {
		return false
	}

	var computedIndent int32
	_ = computedIndent
	computedIndent = 0
	for {
		var tmp_tmp_1 bool
		_ = tmp_tmp_1
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_24 int32
		_ = _hx_tmp_24
		var _hx_tmp_25 bool
		_ = _hx_tmp_25
		if pos < 0 {
			_hx_tmp_25 = true
		} else {
			var _hx_tmp_26 int32
			_ = _hx_tmp_26
			_hx_tmp_26 = pos
			var _hx_tmp_27 int32
			_ = _hx_tmp_27
			_hx_tmp_27 = _hx_tmp_26
			_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_25 {
			_hx_tmp_24 = -1
		} else {
			_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		if (_hx_tmp_24) != 32 {
			var tmp_v_2 HxArray[int32]
			_ = tmp_v_2
			tmp_v_2 = this.Hx_Field_input
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
				_hx_tmp_29 = (_hx_tmp_31 >= HxMakeArrayView[any](tmp_v_2).Len())
			}

			if _hx_tmp_29 {
				_hx_tmp_28 = -1
			} else {
				_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
			}

			tmp_tmp_1 = ((_hx_tmp_28) == 9)
		} else {
			tmp_tmp_1 = true
		}

		if !tmp_tmp_1 {
			break
		}

		computedIndent++
		pos++
	}

	if computedIndent != indent {
		return false
	}

	pos = this.VTable.Hx_Field__skipWhitespaceAndComments(pos, false)
	var tmp1 bool
	_ = tmp1
	var tmp2 bool
	_ = tmp2
	var tmp3 bool
	_ = tmp3
	var v3 HxArray[int32]
	_ = v3
	v3 = this.Hx_Field_input
	var _hx_tmp_32 int32
	_ = _hx_tmp_32
	var _hx_tmp_33 bool
	_ = _hx_tmp_33
	if pos < 0 {
		_hx_tmp_33 = true
	} else {
		var _hx_tmp_34 int32
		_ = _hx_tmp_34
		_hx_tmp_34 = pos
		var _hx_tmp_35 int32
		_ = _hx_tmp_35
		_hx_tmp_35 = _hx_tmp_34
		_hx_tmp_33 = (_hx_tmp_35 >= HxMakeArrayView[any](v3).Len())
	}

	if _hx_tmp_33 {
		_hx_tmp_32 = -1
	} else {
		_hx_tmp_32 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v3).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_32
	if !((c == 32) || (c == 9)) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_36 int32
		_ = _hx_tmp_36
		var _hx_tmp_37 bool
		_ = _hx_tmp_37
		if pos < 0 {
			_hx_tmp_37 = true
		} else {
			var _hx_tmp_38 int32
			_ = _hx_tmp_38
			_hx_tmp_38 = pos
			var _hx_tmp_39 int32
			_ = _hx_tmp_39
			_hx_tmp_39 = _hx_tmp_38
			_hx_tmp_37 = (_hx_tmp_39 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_37 {
			_hx_tmp_36 = -1
		} else {
			_hx_tmp_36 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		tmp3 = ((_hx_tmp_36) == 13)
	} else {
		tmp3 = true
	}

	if !tmp3 {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_40 int32
		_ = _hx_tmp_40
		var _hx_tmp_41 bool
		_ = _hx_tmp_41
		if pos < 0 {
			_hx_tmp_41 = true
		} else {
			var _hx_tmp_42 int32
			_ = _hx_tmp_42
			_hx_tmp_42 = pos
			var _hx_tmp_43 int32
			_ = _hx_tmp_43
			_hx_tmp_43 = _hx_tmp_42
			_hx_tmp_41 = (_hx_tmp_43 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_41 {
			_hx_tmp_40 = -1
		} else {
			_hx_tmp_40 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		tmp2 = ((_hx_tmp_40) == 10)
	} else {
		tmp2 = true
	}

	if !tmp2 {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_tmp_44 int32
		_ = _hx_tmp_44
		var _hx_tmp_45 bool
		_ = _hx_tmp_45
		if pos < 0 {
			_hx_tmp_45 = true
		} else {
			var _hx_tmp_46 int32
			_ = _hx_tmp_46
			_hx_tmp_46 = pos
			var _hx_tmp_47 int32
			_ = _hx_tmp_47
			_hx_tmp_47 = _hx_tmp_46
			_hx_tmp_45 = (_hx_tmp_47 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_45 {
			_hx_tmp_44 = -1
		} else {
			_hx_tmp_44 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
		}

		if (_hx_tmp_44) == 47 {
			var tmp_v_2 HxArray[int32]
			_ = tmp_v_2
			tmp_v_2 = this.Hx_Field_input
			var pos1 int32
			_ = pos1
			pos1 = (pos + 1)
			var _hx_tmp_48 int32
			_ = _hx_tmp_48
			var _hx_tmp_49 bool
			_ = _hx_tmp_49
			if pos1 < 0 {
				_hx_tmp_49 = true
			} else {
				var _hx_tmp_50 int32
				_ = _hx_tmp_50
				_hx_tmp_50 = pos1
				var _hx_tmp_51 int32
				_ = _hx_tmp_51
				_hx_tmp_51 = _hx_tmp_50
				_hx_tmp_49 = (_hx_tmp_51 >= HxMakeArrayView[any](tmp_v_2).Len())
			}

			if _hx_tmp_49 {
				_hx_tmp_48 = -1
			} else {
				_hx_tmp_48 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos1))))
			}

			tmp1 = ((_hx_tmp_48) == 47)
		} else {
			tmp1 = false
		}
	} else {
		tmp1 = true
	}

	if tmp1 {
		return false
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isAfterComma() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if !Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_2 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_LineBreak{})
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Indent{})
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Unindent{})
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			return (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Comma{})
		}

		i--
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isAfterLBracket() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if !Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_2 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_LineBreak{})
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Indent{})
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = (token.Hx_Field_type != Hx_Enum_loreline_tokentype_Unindent{})
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			return (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LBracket{})
		}

		i--
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isInsideBrackets() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	for (i >= 0) && (this.Hx_Field_stack.Get(i) == ((int32)(13))) {
		i--
	}

	return ((i >= 0) && (this.Hx_Field_stack.Get(i) == ((int32)(14))))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsAssignStart() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
		}

		if _hx_tmp_0 {
			i--
		} else {
			return Hx_Field_loreline_tokentypehelpers_isAssignOp(token.Hx_Field_type)
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsUnquotedAssign() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
		}

		if _hx_tmp_0 {
			i--
		} else {
			return (token.Hx_Field_type == Hx_Enum_loreline_tokentype_OpUnquotedAssign{})
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsOnlyLabelOrCommentsInLine() bool {
	var foundLabel bool
	_ = foundLabel
	foundLabel = false
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
		}

		if _hx_tmp_0 {
			i--
		} else {
			if !foundLabel && (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Colon{}) {
				i--
				if i >= 0 {
					if Hx_Field_loreline_tokentypehelpers_isIdentifier(this.Hx_Field_tokenized.Get(i).Hx_Field_type) {
						foundLabel = true
						i--
					}
				} else {
					return false
				}
			} else {
				return (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{})
			}
		}
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsOnlyWhitespacesOrCommentsInLine() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
		}

		if _hx_tmp_0 {
			i--
		} else {
			return (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{})
		}
	}

	return true
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsChoiceOncePrefix() bool {
	if !this.VTable.Hx_Field_inChoiceRoot() {
		return false
	}

	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if Hx_Field_loreline_tokentypehelpers_isComment(token.Hx_Field_type) {
			_hx_tmp_1 = true
		} else {
			_hx_tmp_1 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
		}

		if _hx_tmp_0 {
			i--
		} else {
			if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_OpMinus{}) {
				i--
				for {
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					if i >= 0 {
						var _hx_tmp_3 bool
						_ = _hx_tmp_3
						var _hx_tmp_4 bool
						_ = _hx_tmp_4
						if Hx_Field_loreline_tokentypehelpers_isComment(this.Hx_Field_tokenized.Get(i).Hx_Field_type) {
							_hx_tmp_4 = true
						} else {
							_hx_tmp_4 = (this.Hx_Field_tokenized.Get(i).Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{})
						}

						if _hx_tmp_4 {
							_hx_tmp_3 = true
						} else {
							_hx_tmp_3 = (this.Hx_Field_tokenized.Get(i).Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})
						}

						_hx_tmp_2 = (_hx_tmp_3)
					} else {
						_hx_tmp_2 = false
					}

					if !(_hx_tmp_2) {
						break
					}

					i--
				}

				return ((i < 0) || (this.Hx_Field_tokenized.Get(i).Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{}))
			} else {
				return false
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_hasNonSpecialChar(str string) bool {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(str))
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var found bool
			_ = found
			found = false
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(str, i))
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			if (c).Valid == false {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = (c).Value
			}

			var code int32
			_ = code
			code = _hx_tmp_1
			{
				var tmp__g1_1 int32
				_ = tmp__g1_1
				tmp__g1_1 = 0
				var _g2 int32
				_ = _g2
				_g2 = int32(utf8.RuneCountInString(Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars))
				for tmp__g1_1 < _g2 {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = tmp__g1_1
					tmp__g1_1 = (tmp__g1_1 + 1)
					var j int32
					_ = j
					j = _hx_tmp_2
					var tmp_c_1 HxNullable[int32]
					_ = tmp_c_1
					tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars, j))
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = code
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					if (tmp_c_1).Valid == false {
						_hx_tmp_4 = -1
					} else {
						_hx_tmp_4 = (tmp_c_1).Value
					}

					if _hx_tmp_3 == (_hx_tmp_4) {
						found = true
						break
					}
				}
			}

			if !found {
				return true
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isNumber(value string) bool {
	var pos int32
	_ = pos
	pos = 0
	var length int32
	_ = length
	length = int32(utf8.RuneCountInString(value))
	var tmp bool
	_ = tmp
	if pos < length {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(value, pos))
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		if (c).Valid == false {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = (c).Value
		}

		tmp = ((_hx_tmp_0) == 45)
	} else {
		tmp = false
	}

	if tmp {
		pos++
	}

	var digitStart int32
	_ = digitStart
	digitStart = pos
	for {
		var tmp_tmp_1 bool
		_ = tmp_tmp_1
		if pos < length {
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(value, pos))
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			if (c).Valid == false {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = (c).Value
			}

			var c1 int32
			_ = c1
			c1 = _hx_tmp_1
			tmp_tmp_1 = ((c1 >= 48) && (c1 <= 57))
		} else {
			tmp_tmp_1 = false
		}

		if !tmp_tmp_1 {
			break
		}

		pos++
	}

	if pos == digitStart {
		return false
	}

	var tmp1 bool
	_ = tmp1
	var tmp2 bool
	_ = tmp2
	if pos < length {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(value, pos))
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (c).Valid == false {
			_hx_tmp_2 = -1
		} else {
			_hx_tmp_2 = (c).Value
		}

		tmp2 = ((_hx_tmp_2) == 46)
	} else {
		tmp2 = false
	}

	if tmp2 && ((pos + 1) < length) {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(value, (pos + 1)))
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (c).Valid == false {
			_hx_tmp_3 = -1
		} else {
			_hx_tmp_3 = (c).Value
		}

		var c1 int32
		_ = c1
		c1 = _hx_tmp_3
		tmp1 = ((c1 >= 48) && (c1 <= 57))
	} else {
		tmp1 = false
	}

	if tmp1 {
		pos++
		for {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if pos < length {
				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(value, pos))
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				if (c).Valid == false {
					_hx_tmp_4 = -1
				} else {
					_hx_tmp_4 = (c).Value
				}

				var c1 int32
				_ = c1
				c1 = _hx_tmp_4
				tmp_tmp_1 = ((c1 >= 48) && (c1 <= 57))
			} else {
				tmp_tmp_1 = false
			}

			if !tmp_tmp_1 {
				break
			}

			pos++
		}
	}

	return (pos == length)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_makeLooseOrStrictAfterBracket() {
	if this.VTable.Hx_Field_isStrict() {
		_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
		return
	}

	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		switch token.Hx_Field_type.Hx_Field_enumIndex() {
		case 20:
			{
				_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
				return
			}

		case 49, 51:
			{
				_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
				return
			}

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
				_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(0))))
				return
			}
		}

		i--
	}

	_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(2))))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_makeStrictIfFollowingCallable() {
	if this.VTable.Hx_Field_followsCallableOrIndexable() {
		_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
	} else {
		_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(2))))
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_followsCallableOrIndexable() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_tokenized.Len() - 1)
	for i >= 0 {
		var token *Hx_Obj_loreline_token
		_ = token
		token = this.Hx_Field_tokenized.Get(i)
		switch token.Hx_Field_type.Hx_Field_enumIndex() {
		case 2, 49, 51:
			{
				return true
			}

		case 20:
			{
				return true
			}

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
				return false
			}
		}

		i--
	}

	return false
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isStrict() bool {
	var i int32
	_ = i
	i = (this.Hx_Field_strictExprs.Len() - 1)
	for (i >= 0) && (this.Hx_Field_strictExprs.Get(i) == ((int32)(2))) {
		i--
	}

	return ((i >= 0) && (this.Hx_Field_strictExprs.Get(i) == ((int32)(1))))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readImportValue(c int32, startPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	if c == 34 {
		return this.VTable.Hx_Field_readString(startPos)
	} else {
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		for this.Hx_Field_pos < this.Hx_Field_length {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
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

			var cc int32
			_ = cc
			cc = _hx_tmp_0
			if cc == 10 {
				break
			}

			if (cc == 47) && ((this.Hx_Field_pos + 1) < this.Hx_Field_length) {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (this.Hx_Field_pos + 1)
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

				var next int32
				_ = next
				next = _hx_tmp_4
				if (next == 47) || (next == 42) {
					break
				}
			}

			if (cc == 35) && ((this.Hx_Field_pos + 1) < this.Hx_Field_length) {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (this.Hx_Field_pos + 1)
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

				var next int32
				_ = next
				next = _hx_tmp_8
				if (((((next >= 97) && (next <= 122)) || ((next >= 65) && (next <= 90))) || (next == 95)) || ((next >= 48) && (next <= 57))) || (next == 45) {
					break
				}
			}

			buf.VTable.Hx_Field_addChar(cc)
			this.Hx_Field_pos++
		}

		var s string
		_ = s
		s = buf.VTable.Hx_Field_toString()
		var l int32
		_ = l
		l = int32(utf8.RuneCountInString(s))
		var r int32
		_ = r
		r = 0
		for {
			var _hx_tmp_12 bool
			_ = _hx_tmp_12
			if r < l {
				_hx_tmp_12 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
			} else {
				_hx_tmp_12 = false
			}

			if !(_hx_tmp_12) {
				break
			}

			r++
		}

		var _hx_tmp_13 string
		_ = _hx_tmp_13
		if r > 0 {
			var _hx_tmp_14 string
			_ = _hx_tmp_14
			_hx_tmp_14 = s
			var _hx_tmp_15 string
			_ = _hx_tmp_15
			_hx_tmp_15 = _hx_tmp_14
			_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
		} else {
			_hx_tmp_13 = s
		}

		var value string
		_ = value
		value = _hx_tmp_13
		var position *Hx_Obj_loreline_position
		_ = position
		position = Hx_Obj_loreline_position_CreateInstance(startPos.Hx_Field_line, startPos.Hx_Field_column, startPos.Hx_Field_offset, int32(utf8.RuneCountInString(value)))
		if position == nil {
			position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
		}

		position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
		var _hx_tmp_16 string
		_ = _hx_tmp_16
		_hx_tmp_16 = value
		var token *Hx_Obj_loreline_token
		_ = token
		token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LString{((int32)(0)), _hx_tmp_16, HxMakeArray[Hx_Enum_loreline_lstringattachment]()}, position)
		this.Hx_Field_previous = token
		return token
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isTagStart(pos int32) bool {
	var tmp bool
	_ = tmp
	if pos < this.Hx_Field_length {
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

		tmp = ((_hx_tmp_0) != 60)
	} else {
		tmp = true
	}

	if tmp {
		return false
	}

	var isClosing bool
	_ = isClosing
	if (pos + 1) < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos1 int32
		_ = pos1
		pos1 = (pos + 1)
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if pos1 < 0 {
			_hx_tmp_5 = true
		} else {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos1
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _hx_tmp_6
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos1))))
		}

		isClosing = ((_hx_tmp_4) == 47)
	} else {
		isClosing = false
	}

	var _hx_tmp_8 int32
	_ = _hx_tmp_8
	_hx_tmp_8 = pos
	var _hx_tmp_9 int32
	_ = _hx_tmp_9
	if isClosing {
		_hx_tmp_9 = 2
	} else {
		_hx_tmp_9 = 1
	}

	var checkPos int32
	_ = checkPos
	checkPos = (_hx_tmp_8 + (_hx_tmp_9))
	if checkPos >= this.Hx_Field_length {
		return false
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_tmp_10 int32
	_ = _hx_tmp_10
	var _hx_tmp_11 bool
	_ = _hx_tmp_11
	if checkPos < 0 {
		_hx_tmp_11 = true
	} else {
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		_hx_tmp_12 = checkPos
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_12
		_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_11 {
		_hx_tmp_10 = -1
	} else {
		_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(checkPos))))
	}

	var nameStart int32
	_ = nameStart
	nameStart = _hx_tmp_10
	return (((((((nameStart >= 97) && (nameStart <= 122)) || ((nameStart >= 65) && (nameStart <= 90))) || (nameStart == 95)) || (nameStart == 95)) || (nameStart == 36)) || (isClosing && (nameStart == 62)))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_scanTagEnd(startPos int32) int32 {
	if !this.VTable.Hx_Field_isTagStart(startPos) {
		return -1
	}

	var p int32
	_ = p
	p = (startPos + 1)
	for p < this.Hx_Field_length {
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

		var c int32
		_ = c
		c = _hx_tmp_0
		if c == 92 {
			p = (p + 2)
			continue
		}

		if (c == 10) || (c == 13) {
			return -1
		}

		if c == 62 {
			return (p + 1)
		}

		p++
	}

	return -1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_scanStringEnd(startPos int32, allowTags bool) int32 {
	var p int32
	_ = p
	p = (startPos + 1)
	var inTag bool
	_ = inTag
	inTag = false
	for p < this.Hx_Field_length {
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

		var c int32
		_ = c
		c = _hx_tmp_0
		if c == 92 {
			p = (p + 2)
			continue
		}

		if c == 10 {
			return -1
		}

		if (c == 34) && !inTag {
			return (p + 1)
		}

		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		if (allowTags && (c == 60)) && !inTag {
			_hx_tmp_4 = this.VTable.Hx_Field_isTagStart(p)
		} else {
			_hx_tmp_4 = false
		}

		if _hx_tmp_4 {
			inTag = true
		}

		if (c == 62) && inTag {
			inTag = false
		}

		if (c == 36) && ((p + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (p + 1)
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			if pos < 0 {
				_hx_tmp_6 = true
			} else {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = pos
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _hx_tmp_7
				_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_6 {
				_hx_tmp_5 = -1
			} else {
				_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(pos))))
			}

			var next int32
			_ = next
			next = _hx_tmp_5
			if next == 123 {
				p = (p + 2)
				var depth int32
				_ = depth
				depth = 1
				for (p < this.Hx_Field_length) && (depth > 0) {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					var _hx_tmp_10 bool
					_ = _hx_tmp_10
					if p < 0 {
						_hx_tmp_10 = true
					} else {
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = p
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						_hx_tmp_12 = _hx_tmp_11
						_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_10 {
						_hx_tmp_9 = -1
					} else {
						_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(p))))
					}

					var ic int32
					_ = ic
					ic = _hx_tmp_9
					if ic == 92 {
						p = (p + 2)
						continue
					}

					if ic == 10 {
						return -1
					}

					var tmp bool
					_ = tmp
					if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = (p + 1)
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
							_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_14 {
							_hx_tmp_13 = -1
						} else {
							_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_1))))
						}

						tmp = ((_hx_tmp_13) == 47)
					} else {
						tmp = false
					}

					if tmp {
						return -1
					}

					var tmp1 bool
					_ = tmp1
					if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = (p + 1)
						var _hx_tmp_17 int32
						_ = _hx_tmp_17
						var _hx_tmp_18 bool
						_ = _hx_tmp_18
						if tmp_pos_1 < 0 {
							_hx_tmp_18 = true
						} else {
							var _hx_tmp_19 int32
							_ = _hx_tmp_19
							_hx_tmp_19 = tmp_pos_1
							var _hx_tmp_20 int32
							_ = _hx_tmp_20
							_hx_tmp_20 = _hx_tmp_19
							_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_18 {
							_hx_tmp_17 = -1
						} else {
							_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_1))))
						}

						tmp1 = ((_hx_tmp_17) == 42)
					} else {
						tmp1 = false
					}

					if tmp1 {
						p = (p + 2)
						for {
							if p >= this.Hx_Field_length {
								return -1
							}

							var tmp_v_3 HxArray[int32]
							_ = tmp_v_3
							tmp_v_3 = this.Hx_Field_input
							var _hx_tmp_21 int32
							_ = _hx_tmp_21
							var _hx_tmp_22 bool
							_ = _hx_tmp_22
							if p < 0 {
								_hx_tmp_22 = true
							} else {
								var _hx_tmp_23 int32
								_ = _hx_tmp_23
								_hx_tmp_23 = p
								var _hx_tmp_24 int32
								_ = _hx_tmp_24
								_hx_tmp_24 = _hx_tmp_23
								_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_3).Len())
							}

							if _hx_tmp_22 {
								_hx_tmp_21 = -1
							} else {
								_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(p))))
							}

							if (_hx_tmp_21) == 10 {
								return -1
							}

							var tmp_tmp_1 bool
							_ = tmp_tmp_1
							var tmp_tmp1_1 bool
							_ = tmp_tmp1_1
							if (p + 1) < this.Hx_Field_length {
								var tmp_v_4 HxArray[int32]
								_ = tmp_v_4
								tmp_v_4 = this.Hx_Field_input
								var _hx_tmp_25 int32
								_ = _hx_tmp_25
								var _hx_tmp_26 bool
								_ = _hx_tmp_26
								if p < 0 {
									_hx_tmp_26 = true
								} else {
									var _hx_tmp_27 int32
									_ = _hx_tmp_27
									_hx_tmp_27 = p
									var _hx_tmp_28 int32
									_ = _hx_tmp_28
									_hx_tmp_28 = _hx_tmp_27
									_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](tmp_v_4).Len())
								}

								if _hx_tmp_26 {
									_hx_tmp_25 = -1
								} else {
									_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(p))))
								}

								tmp_tmp1_1 = ((_hx_tmp_25) == 42)
							} else {
								tmp_tmp1_1 = false
							}

							if tmp_tmp1_1 {
								var tmp_v_4 HxArray[int32]
								_ = tmp_v_4
								tmp_v_4 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = (p + 1)
								var _hx_tmp_29 int32
								_ = _hx_tmp_29
								var _hx_tmp_30 bool
								_ = _hx_tmp_30
								if tmp_pos_1 < 0 {
									_hx_tmp_30 = true
								} else {
									var _hx_tmp_31 int32
									_ = _hx_tmp_31
									_hx_tmp_31 = tmp_pos_1
									var _hx_tmp_32 int32
									_ = _hx_tmp_32
									_hx_tmp_32 = _hx_tmp_31
									_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_4).Len())
								}

								if _hx_tmp_30 {
									_hx_tmp_29 = -1
								} else {
									_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos_1))))
								}

								tmp_tmp_1 = ((_hx_tmp_29) == 47)
							} else {
								tmp_tmp_1 = false
							}

							if tmp_tmp_1 {
								p = (p + 2)
								break
							}

							p++
						}

						continue
					}

					if ic == 34 {
						p = this.VTable.Hx_Field_scanStringEnd(p, allowTags)
						if p == -1 {
							return -1
						}

						continue
					}

					if ic == 123 {
						depth++
					} else {
						if ic == 125 {
							depth--
						}
					}

					p++
				}

				if depth != 0 {
					return -1
				}

				continue
			}

			if next == 36 {
				p = (p + 2)
				continue
			}

			if (((next >= 97) && (next <= 122)) || ((next >= 65) && (next <= 90))) || (next == 95) {
				p++
				for {
					var tmp bool
					_ = tmp
					if p < this.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_33 int32
						_ = _hx_tmp_33
						var _hx_tmp_34 bool
						_ = _hx_tmp_34
						if p < 0 {
							_hx_tmp_34 = true
						} else {
							var _hx_tmp_35 int32
							_ = _hx_tmp_35
							_hx_tmp_35 = p
							var _hx_tmp_36 int32
							_ = _hx_tmp_36
							_hx_tmp_36 = _hx_tmp_35
							_hx_tmp_34 = (_hx_tmp_36 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_34 {
							_hx_tmp_33 = -1
						} else {
							_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(p))))
						}

						var tmp_c_1 int32
						_ = tmp_c_1
						tmp_c_1 = _hx_tmp_33
						tmp = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
					} else {
						tmp = false
					}

					if !tmp {
						break
					}

					p++
				}

				p = this.VTable.Hx_Field_scanAccessorChain(p, allowTags)
				if p == -1 {
					return -1
				}

				continue
			}
		}

		p++
	}

	return -1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_scanAccessorChain(p int32, allowTags bool) int32 {
	for p < this.Hx_Field_length {
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

		var c int32
		_ = c
		c = _hx_tmp_0
		var tmp bool
		_ = tmp
		if (c == 46) && ((p + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (p + 1)
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
			tmp = ((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95))
		} else {
			tmp = false
		}

		if tmp {
			p++
			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if p < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
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
						_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_9 {
						_hx_tmp_8 = -1
					} else {
						_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_8
					tmp_tmp_1 = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				p++
			}
		} else {
			if (c == 91) || (c == 40) {
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				if c == 91 {
					_hx_tmp_12 = 93
				} else {
					_hx_tmp_12 = 41
				}

				var _hx_reserved_close int32
				_ = _hx_reserved_close
				_hx_reserved_close = _hx_tmp_12
				p++
				var depth int32
				_ = depth
				depth = 1
				for (p < this.Hx_Field_length) && (depth > 0) {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					var _hx_tmp_14 bool
					_ = _hx_tmp_14
					if p < 0 {
						_hx_tmp_14 = true
					} else {
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = p
						var _hx_tmp_16 int32
						_ = _hx_tmp_16
						_hx_tmp_16 = _hx_tmp_15
						_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_14 {
						_hx_tmp_13 = -1
					} else {
						_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
					}

					var ic int32
					_ = ic
					ic = _hx_tmp_13
					if ic == 92 {
						p = (p + 2)
						continue
					}

					if ic == 10 {
						return -1
					}

					var tmp_tmp_1 bool
					_ = tmp_tmp_1
					if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var pos int32
						_ = pos
						pos = (p + 1)
						var _hx_tmp_17 int32
						_ = _hx_tmp_17
						var _hx_tmp_18 bool
						_ = _hx_tmp_18
						if pos < 0 {
							_hx_tmp_18 = true
						} else {
							var _hx_tmp_19 int32
							_ = _hx_tmp_19
							_hx_tmp_19 = pos
							var _hx_tmp_20 int32
							_ = _hx_tmp_20
							_hx_tmp_20 = _hx_tmp_19
							_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_18 {
							_hx_tmp_17 = -1
						} else {
							_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
						}

						tmp_tmp_1 = ((_hx_tmp_17) == 47)
					} else {
						tmp_tmp_1 = false
					}

					if tmp_tmp_1 {
						return -1
					}

					var tmp1 bool
					_ = tmp1
					if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var pos int32
						_ = pos
						pos = (p + 1)
						var _hx_tmp_21 int32
						_ = _hx_tmp_21
						var _hx_tmp_22 bool
						_ = _hx_tmp_22
						if pos < 0 {
							_hx_tmp_22 = true
						} else {
							var _hx_tmp_23 int32
							_ = _hx_tmp_23
							_hx_tmp_23 = pos
							var _hx_tmp_24 int32
							_ = _hx_tmp_24
							_hx_tmp_24 = _hx_tmp_23
							_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_22 {
							_hx_tmp_21 = -1
						} else {
							_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
						}

						tmp1 = ((_hx_tmp_21) == 42)
					} else {
						tmp1 = false
					}

					if tmp1 {
						p = (p + 2)
						for {
							if p >= this.Hx_Field_length {
								return -1
							}

							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var _hx_tmp_25 int32
							_ = _hx_tmp_25
							var _hx_tmp_26 bool
							_ = _hx_tmp_26
							if p < 0 {
								_hx_tmp_26 = true
							} else {
								var _hx_tmp_27 int32
								_ = _hx_tmp_27
								_hx_tmp_27 = p
								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								_hx_tmp_28 = _hx_tmp_27
								_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_26 {
								_hx_tmp_25 = -1
							} else {
								_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(p))))
							}

							if (_hx_tmp_25) == 10 {
								return -1
							}

							var tmp_tmp_2 bool
							_ = tmp_tmp_2
							var tmp_tmp1_1 bool
							_ = tmp_tmp1_1
							if (p + 1) < this.Hx_Field_length {
								var tmp_v_3 HxArray[int32]
								_ = tmp_v_3
								tmp_v_3 = this.Hx_Field_input
								var _hx_tmp_29 int32
								_ = _hx_tmp_29
								var _hx_tmp_30 bool
								_ = _hx_tmp_30
								if p < 0 {
									_hx_tmp_30 = true
								} else {
									var _hx_tmp_31 int32
									_ = _hx_tmp_31
									_hx_tmp_31 = p
									var _hx_tmp_32 int32
									_ = _hx_tmp_32
									_hx_tmp_32 = _hx_tmp_31
									_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_3).Len())
								}

								if _hx_tmp_30 {
									_hx_tmp_29 = -1
								} else {
									_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(p))))
								}

								tmp_tmp1_1 = ((_hx_tmp_29) == 42)
							} else {
								tmp_tmp1_1 = false
							}

							if tmp_tmp1_1 {
								var tmp_v_3 HxArray[int32]
								_ = tmp_v_3
								tmp_v_3 = this.Hx_Field_input
								var pos int32
								_ = pos
								pos = (p + 1)
								var _hx_tmp_33 int32
								_ = _hx_tmp_33
								var _hx_tmp_34 bool
								_ = _hx_tmp_34
								if pos < 0 {
									_hx_tmp_34 = true
								} else {
									var _hx_tmp_35 int32
									_ = _hx_tmp_35
									_hx_tmp_35 = pos
									var _hx_tmp_36 int32
									_ = _hx_tmp_36
									_hx_tmp_36 = _hx_tmp_35
									_hx_tmp_34 = (_hx_tmp_36 >= HxMakeArrayView[any](tmp_v_3).Len())
								}

								if _hx_tmp_34 {
									_hx_tmp_33 = -1
								} else {
									_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(pos))))
								}

								tmp_tmp_2 = ((_hx_tmp_33) == 47)
							} else {
								tmp_tmp_2 = false
							}

							if tmp_tmp_2 {
								p = (p + 2)
								break
							}

							p++
						}

						continue
					}

					if ic == 34 {
						p = this.VTable.Hx_Field_scanStringEnd(p, allowTags)
						if p == -1 {
							return -1
						}

						continue
					}

					if ic == c {
						depth++
					} else {
						if ic == _hx_reserved_close {
							depth--
						}
					}

					p++
				}

				if depth != 0 {
					return -1
				}
			} else {
				break
			}
		}
	}

	return p
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_skipBalancedParens(p int32, allowTags bool) int32 {
	var tmp bool
	_ = tmp
	if p < this.Hx_Field_length {
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

		tmp = ((_hx_tmp_0) != 40)
	} else {
		tmp = true
	}

	if tmp {
		return -1
	}

	p++
	var depth int32
	_ = depth
	depth = 1
	for (p < this.Hx_Field_length) && (depth > 0) {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
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
			_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = -1
		} else {
			_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(p))))
		}

		var ic int32
		_ = ic
		ic = _hx_tmp_4
		if ic == 92 {
			p = (p + 2)
			continue
		}

		if ic == 10 {
			return -1
		}

		var tmp_tmp_1 bool
		_ = tmp_tmp_1
		if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (p + 1)
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

			tmp_tmp_1 = ((_hx_tmp_8) == 47)
		} else {
			tmp_tmp_1 = false
		}

		if tmp_tmp_1 {
			return -1
		}

		var tmp1 bool
		_ = tmp1
		if (ic == 47) && ((p + 1) < this.Hx_Field_length) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (p + 1)
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

			tmp1 = ((_hx_tmp_12) == 42)
		} else {
			tmp1 = false
		}

		if tmp1 {
			p = (p + 2)
			for {
				if p >= this.Hx_Field_length {
					return -1
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var _hx_tmp_16 int32
				_ = _hx_tmp_16
				var _hx_tmp_17 bool
				_ = _hx_tmp_17
				if p < 0 {
					_hx_tmp_17 = true
				} else {
					var _hx_tmp_18 int32
					_ = _hx_tmp_18
					_hx_tmp_18 = p
					var _hx_tmp_19 int32
					_ = _hx_tmp_19
					_hx_tmp_19 = _hx_tmp_18
					_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_17 {
					_hx_tmp_16 = -1
				} else {
					_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(p))))
				}

				if (_hx_tmp_16) == 10 {
					return -1
				}

				var tmp_tmp_2 bool
				_ = tmp_tmp_2
				var tmp_tmp1_1 bool
				_ = tmp_tmp1_1
				if (p + 1) < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					var _hx_tmp_21 bool
					_ = _hx_tmp_21
					if p < 0 {
						_hx_tmp_21 = true
					} else {
						var _hx_tmp_22 int32
						_ = _hx_tmp_22
						_hx_tmp_22 = p
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = _hx_tmp_22
						_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_21 {
						_hx_tmp_20 = -1
					} else {
						_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(p))))
					}

					tmp_tmp1_1 = ((_hx_tmp_20) == 42)
				} else {
					tmp_tmp1_1 = false
				}

				if tmp_tmp1_1 {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var pos int32
					_ = pos
					pos = (p + 1)
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					var _hx_tmp_25 bool
					_ = _hx_tmp_25
					if pos < 0 {
						_hx_tmp_25 = true
					} else {
						var _hx_tmp_26 int32
						_ = _hx_tmp_26
						_hx_tmp_26 = pos
						var _hx_tmp_27 int32
						_ = _hx_tmp_27
						_hx_tmp_27 = _hx_tmp_26
						_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_25 {
						_hx_tmp_24 = -1
					} else {
						_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(pos))))
					}

					tmp_tmp_2 = ((_hx_tmp_24) == 47)
				} else {
					tmp_tmp_2 = false
				}

				if tmp_tmp_2 {
					p = (p + 2)
					break
				}

				p++
			}

			continue
		}

		if ic == 34 {
			p = this.VTable.Hx_Field_scanStringEnd(p, allowTags)
			if p == -1 {
				return -1
			}

			continue
		}

		if ic == 40 {
			depth++
		} else {
			if ic == 41 {
				depth--
			}
		}

		p++
	}

	if depth != 0 {
		return -1
	}

	return p
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_tryReadUnquotedString() *Hx_Obj_loreline_token {
	if this.VTable.Hx_Field_isStrict() {
		return nil
	}

	if this.VTable.Hx_Field_inWhenRoot() {
		return nil
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
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
	var offset int32
	_ = offset
	offset = 1
	var cc int32
	_ = cc
	if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = (this.Hx_Field_pos + offset)
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

		cc = _hx_tmp_4
	} else {
		cc = 0
	}

	if (c == 47) && (this.Hx_Field_pos < (this.Hx_Field_length - 1)) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = (this.Hx_Field_pos + 1)
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

		var next int32
		_ = next
		next = _hx_tmp_8
		if (next == 47) || (next == 42) {
			return nil
		}
	}

	if ((((((c == 123) || (c == 125)) || (c == 91)) || (c == 93)) || (c == 58)) || (c == 35)) || ((c == 32) || (c == 9)) {
		return nil
	}

	var parent Hx_Enum_loreline_tokentype
	_ = parent
	parent = this.VTable.Hx_Field_parentBlockType()
	if ((parent != Hx_Enum_loreline_tokentype_KwBeat{}) && (parent != Hx_Enum_loreline_tokentype_KwState{})) && (parent != Hx_Enum_loreline_tokentype_KwCharacter{}) {
		return nil
	}

	var inBrackets bool
	_ = inBrackets
	inBrackets = this.VTable.Hx_Field_isInsideBrackets()
	var isAssignValue bool
	_ = isAssignValue
	isAssignValue = this.VTable.Hx_Field_followsUnquotedAssign()
	var labelIdentifierIndex int32
	_ = labelIdentifierIndex
	labelIdentifierIndex = this.VTable.Hx_Field_afterLabelIdentifierToken(false)
	if (labelIdentifierIndex != -1) && (this.Hx_Field_tokenized.Get(labelIdentifierIndex).Hx_Field_pos.Hx_Field_line < this.Hx_Field_line) {
		if this.Hx_Field_tokenized.Get(labelIdentifierIndex).Hx_Field_pos.Hx_Field_column >= this.Hx_Field_column {
			return nil
		}

		if (this.Hx_Field_line - this.Hx_Field_tokenized.Get(labelIdentifierIndex).Hx_Field_pos.Hx_Field_line) > 1 {
			return nil
		}
	}

	var _hx_tmp_12 int32
	_ = _hx_tmp_12
	if labelIdentifierIndex != -1 {
		_hx_tmp_12 = this.VTable.Hx_Field_isStartingMultilineDialogue(labelIdentifierIndex)
	} else {
		_hx_tmp_12 = -1
	}

	var multilineIndent int32
	_ = multilineIndent
	multilineIndent = _hx_tmp_12
	var isAfterLabel bool
	_ = isAfterLabel
	isAfterLabel = (labelIdentifierIndex != -1)
	var isValue bool
	_ = isValue
	isValue = ((((parent == Hx_Enum_loreline_tokentype_KwState{}) || (parent == Hx_Enum_loreline_tokentype_KwCharacter{})) || inBrackets) || isAssignValue)
	var identifier HxNullable[string]
	_ = identifier
	identifier = this.VTable.Hx_Field_matchIdentifier(this.Hx_Field_pos)
	if isValue {
		var _hx_tmp_13 bool
		_ = _hx_tmp_13
		if this.VTable.Hx_Field_isCallStart(this.Hx_Field_pos) {
			_hx_tmp_13 = true
		} else {
			_hx_tmp_13 = this.VTable.Hx_Field_isLabelStart(this.Hx_Field_pos)
		}

		if _hx_tmp_13 {
			return nil
		}
	} else {
		var _hx_tmp_14 bool
		_ = _hx_tmp_14
		var _hx_tmp_15 bool
		_ = _hx_tmp_15
		var _hx_tmp_16 bool
		_ = _hx_tmp_16
		var _hx_tmp_17 bool
		_ = _hx_tmp_17
		if !isAfterLabel {
			var _hx_tmp_18 bool
			_ = _hx_tmp_18
			if (identifier).Valid == false {
				_hx_tmp_18 = true
			} else {
				_hx_tmp_18 = ((identifier).Value != "true")
			}

			_hx_tmp_17 = (_hx_tmp_18)
		} else {
			_hx_tmp_17 = false
		}

		if _hx_tmp_17 {
			var _hx_tmp_19 bool
			_ = _hx_tmp_19
			if (identifier).Valid == false {
				_hx_tmp_19 = true
			} else {
				_hx_tmp_19 = ((identifier).Value != "false")
			}

			_hx_tmp_16 = (_hx_tmp_19)
		} else {
			_hx_tmp_16 = false
		}

		if _hx_tmp_16 {
			var _hx_tmp_20 bool
			_ = _hx_tmp_20
			if (identifier).Valid == false {
				_hx_tmp_20 = true
			} else {
				_hx_tmp_20 = ((identifier).Value != "null")
			}

			_hx_tmp_15 = (_hx_tmp_20)
		} else {
			_hx_tmp_15 = false
		}

		if _hx_tmp_15 {
			var _hx_tmp_21 bool
			_ = _hx_tmp_21
			var _hx_tmp_22 bool
			_ = _hx_tmp_22
			var _hx_tmp_23 bool
			_ = _hx_tmp_23
			var _hx_tmp_24 bool
			_ = _hx_tmp_24
			if this.Hx_Field_keepOrphanExpressions {
				_hx_tmp_24 = this.VTable.Hx_Field_isIdentifierExpressionStart(this.Hx_Field_pos, true)
			} else {
				_hx_tmp_24 = false
			}

			if _hx_tmp_24 {
				_hx_tmp_23 = true
			} else {
				_hx_tmp_23 = this.VTable.Hx_Field_isIfStart(this.Hx_Field_pos)
			}

			if _hx_tmp_23 {
				_hx_tmp_22 = true
			} else {
				_hx_tmp_22 = this.VTable.Hx_Field_isCallStart(this.Hx_Field_pos)
			}

			if _hx_tmp_22 {
				_hx_tmp_21 = true
			} else {
				_hx_tmp_21 = this.VTable.Hx_Field_isAssignStart(this.Hx_Field_pos, false)
			}

			_hx_tmp_14 = (_hx_tmp_21)
		} else {
			_hx_tmp_14 = false
		}

		if _hx_tmp_14 {
			return nil
		}
	}

	if !isValue {
		if (c == 61) || ((cc == 61) && ((((c == 43) || (c == 45)) || (c == 42)) || (c == 47))) {
			return nil
		}
	}

	if !isValue && !isAfterLabel {
		if this.VTable.Hx_Field_isInsertionStart(this.Hx_Field_pos) {
			return nil
		}

		if this.VTable.Hx_Field_isWhenStart(this.Hx_Field_pos) {
			return nil
		}

		if this.VTable.Hx_Field_isChoiceOnceStart(this.Hx_Field_pos) {
			return nil
		}
	}

	if (identifier).Valid != false {
		var _hx_tmp_25 bool
		_ = _hx_tmp_25
		if (identifier).Valid == true {
			_hx_tmp_25 = ((identifier).Value == "function")
		} else {
			_hx_tmp_25 = false
		}

		if _hx_tmp_25 {
			return nil
		}

		if !isValue {
			var _hx_tmp_26 bool
			_ = _hx_tmp_26
			var _hx_tmp_27 bool
			_ = _hx_tmp_27
			var _hx_tmp_28 bool
			_ = _hx_tmp_28
			var _hx_tmp_29 bool
			_ = _hx_tmp_29
			var _hx_tmp_30 bool
			_ = _hx_tmp_30
			var _hx_tmp_31 bool
			_ = _hx_tmp_31
			var _hx_tmp_32 bool
			_ = _hx_tmp_32
			var _hx_tmp_33 bool
			_ = _hx_tmp_33
			var _hx_tmp_34 bool
			_ = _hx_tmp_34
			if (identifier).Valid == false {
				_hx_tmp_34 = true
			} else {
				_hx_tmp_34 = ((identifier).Value != "if")
			}

			if _hx_tmp_34 {
				var _hx_tmp_35 bool
				_ = _hx_tmp_35
				if (identifier).Valid == false {
					_hx_tmp_35 = true
				} else {
					_hx_tmp_35 = ((identifier).Value != "null")
				}

				_hx_tmp_33 = (_hx_tmp_35)
			} else {
				_hx_tmp_33 = false
			}

			if _hx_tmp_33 {
				var _hx_tmp_36 bool
				_ = _hx_tmp_36
				if (identifier).Valid == false {
					_hx_tmp_36 = true
				} else {
					_hx_tmp_36 = ((identifier).Value != "true")
				}

				_hx_tmp_32 = (_hx_tmp_36)
			} else {
				_hx_tmp_32 = false
			}

			if _hx_tmp_32 {
				var _hx_tmp_37 bool
				_ = _hx_tmp_37
				if (identifier).Valid == false {
					_hx_tmp_37 = true
				} else {
					_hx_tmp_37 = ((identifier).Value != "false")
				}

				_hx_tmp_31 = (_hx_tmp_37)
			} else {
				_hx_tmp_31 = false
			}

			if _hx_tmp_31 {
				var _hx_tmp_38 bool
				_ = _hx_tmp_38
				if (identifier).Valid == false {
					_hx_tmp_38 = true
				} else {
					_hx_tmp_38 = ((identifier).Value != "and")
				}

				_hx_tmp_30 = (_hx_tmp_38)
			} else {
				_hx_tmp_30 = false
			}

			if _hx_tmp_30 {
				var _hx_tmp_39 bool
				_ = _hx_tmp_39
				if (identifier).Valid == false {
					_hx_tmp_39 = true
				} else {
					_hx_tmp_39 = ((identifier).Value != "or")
				}

				_hx_tmp_29 = (_hx_tmp_39)
			} else {
				_hx_tmp_29 = false
			}

			if _hx_tmp_29 {
				var _hx_tmp_40 bool
				_ = _hx_tmp_40
				if (identifier).Valid == false {
					_hx_tmp_40 = true
				} else {
					_hx_tmp_40 = ((identifier).Value != "is")
				}

				_hx_tmp_28 = (_hx_tmp_40)
			} else {
				_hx_tmp_28 = false
			}

			if _hx_tmp_28 {
				var _hx_tmp_41 bool
				_ = _hx_tmp_41
				if (identifier).Valid == false {
					_hx_tmp_41 = true
				} else {
					_hx_tmp_41 = ((identifier).Value != "not")
				}

				_hx_tmp_27 = (_hx_tmp_41)
			} else {
				_hx_tmp_27 = false
			}

			if _hx_tmp_27 {
				_hx_tmp_26 = func() func(_hx_param_key string) bool {
					var _hx_closure_src_137 func(_hx_param_key any) bool
					_ = _hx_closure_src_137
					_hx_closure_src_137 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_137(((any)(_hx_param_key)))
					}
				}()((identifier).Value)
			} else {
				_hx_tmp_26 = false
			}

			if _hx_tmp_26 {
				return nil
			}

			var _hx_tmp_43 int32
			_ = _hx_tmp_43
			_hx_tmp_43 = this.Hx_Field_pos
			var _hx_tmp_42 int32
			_ = _hx_tmp_42
			_hx_tmp_42 = (_hx_tmp_43 + int32(utf8.RuneCountInString(identifier.Value)))
			if this.VTable.Hx_Field_isColon(_hx_tmp_42, nil) {
				return nil
			}
		}
	}

	var allowTags bool
	_ = allowTags
	allowTags = !isValue
	var isDialogue bool
	_ = isDialogue
	isDialogue = false
	if isValue {
		if inBrackets {
			var _hx_tmp_44 bool
			_ = _hx_tmp_44
			var _hx_tmp_45 bool
			_ = _hx_tmp_45
			var _hx_tmp_46 bool
			_ = _hx_tmp_46
			if !isAfterLabel {
				_hx_tmp_46 = !this.VTable.Hx_Field_followsOnlyWhitespacesOrCommentsInLine()
			} else {
				_hx_tmp_46 = false
			}

			if _hx_tmp_46 {
				_hx_tmp_45 = !this.VTable.Hx_Field_isAfterComma()
			} else {
				_hx_tmp_45 = false
			}

			if _hx_tmp_45 {
				_hx_tmp_44 = !this.VTable.Hx_Field_isAfterLBracket()
			} else {
				_hx_tmp_44 = false
			}

			if _hx_tmp_44 {
				return nil
			}
		} else {
			if !isAssignValue && !isAfterLabel {
				return nil
			}
		}
	} else {
		isDialogue = this.VTable.Hx_Field_followsOnlyLabelOrCommentsInLine()
		var _hx_tmp_47 bool
		_ = _hx_tmp_47
		var _hx_tmp_48 bool
		_ = _hx_tmp_48
		if !isDialogue {
			_hx_tmp_48 = !this.VTable.Hx_Field_followsOnlyWhitespacesOrCommentsInLine()
		} else {
			_hx_tmp_48 = false
		}

		if _hx_tmp_48 {
			_hx_tmp_47 = !this.VTable.Hx_Field_followsChoiceOncePrefix()
		} else {
			_hx_tmp_47 = false
		}

		if _hx_tmp_47 {
			return nil
		}

		if !isDialogue {
			if (cc == 62) && (c == 45) {
				return nil
			}
		}
	}

	if c == 34 {
		var scanEndPos int32
		_ = scanEndPos
		scanEndPos = this.VTable.Hx_Field_scanStringEnd(this.Hx_Field_pos, (parent == Hx_Enum_loreline_tokentype_KwBeat{}))
		if scanEndPos == -1 {
			return nil
		}

		var hasSameLineContent bool
		_ = hasSameLineContent
		hasSameLineContent = false
		var eolPos int32
		_ = eolPos
		eolPos = scanEndPos
		for eolPos < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var _hx_tmp_49 int32
			_ = _hx_tmp_49
			var _hx_tmp_50 bool
			_ = _hx_tmp_50
			if eolPos < 0 {
				_hx_tmp_50 = true
			} else {
				var _hx_tmp_51 int32
				_ = _hx_tmp_51
				_hx_tmp_51 = eolPos
				var _hx_tmp_52 int32
				_ = _hx_tmp_52
				_hx_tmp_52 = _hx_tmp_51
				_hx_tmp_50 = (_hx_tmp_52 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_50 {
				_hx_tmp_49 = -1
			} else {
				_hx_tmp_49 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(eolPos))))
			}

			var sc int32
			_ = sc
			sc = _hx_tmp_49
			if (sc == 32) || (sc == 9) {
				eolPos++
				continue
			}

			if (sc == 10) || (sc == 13) {
				break
			}

			var tmp bool
			_ = tmp
			if (sc == 47) && ((eolPos + 1) < this.Hx_Field_length) {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (eolPos + 1)
				var _hx_tmp_53 int32
				_ = _hx_tmp_53
				var _hx_tmp_54 bool
				_ = _hx_tmp_54
				if tmp_pos_1 < 0 {
					_hx_tmp_54 = true
				} else {
					var _hx_tmp_55 int32
					_ = _hx_tmp_55
					_hx_tmp_55 = tmp_pos_1
					var _hx_tmp_56 int32
					_ = _hx_tmp_56
					_hx_tmp_56 = _hx_tmp_55
					_hx_tmp_54 = (_hx_tmp_56 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_54 {
					_hx_tmp_53 = -1
				} else {
					_hx_tmp_53 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
				}

				if (_hx_tmp_53) != 47 {
					var tmp_v_3 HxArray[int32]
					_ = tmp_v_3
					tmp_v_3 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = (eolPos + 1)
					var _hx_tmp_57 int32
					_ = _hx_tmp_57
					var _hx_tmp_58 bool
					_ = _hx_tmp_58
					if tmp_pos_2 < 0 {
						_hx_tmp_58 = true
					} else {
						var _hx_tmp_59 int32
						_ = _hx_tmp_59
						_hx_tmp_59 = tmp_pos_2
						var _hx_tmp_60 int32
						_ = _hx_tmp_60
						_hx_tmp_60 = _hx_tmp_59
						_hx_tmp_58 = (_hx_tmp_60 >= HxMakeArrayView[any](tmp_v_3).Len())
					}

					if _hx_tmp_58 {
						_hx_tmp_57 = -1
					} else {
						_hx_tmp_57 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_2))))
					}

					tmp = ((_hx_tmp_57) == 42)
				} else {
					tmp = true
				}
			} else {
				tmp = false
			}

			if tmp {
				break
			}

			if (((sc == 93) || (sc == 125)) || (sc == 41)) || (sc == 44) {
				break
			}

			if sc == 35 {
				break
			}

			hasSameLineContent = true
			break
		}

		if !hasSameLineContent && !isValue {
			var nextLinePos int32
			_ = nextLinePos
			nextLinePos = this.VTable.Hx_Field__skipWhitespaceAndComments(eolPos, true)
			if (nextLinePos <= eolPos) || (nextLinePos >= this.Hx_Field_length) {
				return nil
			}

			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var _hx_tmp_61 int32
			_ = _hx_tmp_61
			var _hx_tmp_62 bool
			_ = _hx_tmp_62
			if nextLinePos < 0 {
				_hx_tmp_62 = true
			} else {
				var _hx_tmp_63 int32
				_ = _hx_tmp_63
				_hx_tmp_63 = nextLinePos
				var _hx_tmp_64 int32
				_ = _hx_tmp_64
				_hx_tmp_64 = _hx_tmp_63
				_hx_tmp_62 = (_hx_tmp_64 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_62 {
				_hx_tmp_61 = -1
			} else {
				_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(nextLinePos))))
			}

			var nextChar int32
			_ = nextChar
			nextChar = _hx_tmp_61
			var tmp bool
			_ = tmp
			if !(((nextChar == 10) || (nextChar == 13)) || (nextChar == 34)) {
				if (nextChar == 47) && ((nextLinePos + 1) < this.Hx_Field_length) {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = (nextLinePos + 1)
					var _hx_tmp_65 int32
					_ = _hx_tmp_65
					var _hx_tmp_66 bool
					_ = _hx_tmp_66
					if tmp_pos_1 < 0 {
						_hx_tmp_66 = true
					} else {
						var _hx_tmp_67 int32
						_ = _hx_tmp_67
						_hx_tmp_67 = tmp_pos_1
						var _hx_tmp_68 int32
						_ = _hx_tmp_68
						_hx_tmp_68 = _hx_tmp_67
						_hx_tmp_66 = (_hx_tmp_68 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_66 {
						_hx_tmp_65 = -1
					} else {
						_hx_tmp_65 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
					}

					if (_hx_tmp_65) != 47 {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = (nextLinePos + 1)
						var _hx_tmp_69 int32
						_ = _hx_tmp_69
						var _hx_tmp_70 bool
						_ = _hx_tmp_70
						if tmp_pos_2 < 0 {
							_hx_tmp_70 = true
						} else {
							var _hx_tmp_71 int32
							_ = _hx_tmp_71
							_hx_tmp_71 = tmp_pos_2
							var _hx_tmp_72 int32
							_ = _hx_tmp_72
							_hx_tmp_72 = _hx_tmp_71
							_hx_tmp_70 = (_hx_tmp_72 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_70 {
							_hx_tmp_69 = -1
						} else {
							_hx_tmp_69 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_2))))
						}

						tmp = ((_hx_tmp_69) == 42)
					} else {
						tmp = true
					}
				} else {
					tmp = false
				}
			} else {
				tmp = true
			}

			if tmp {
				return nil
			}

			var nextIndent int32
			_ = nextIndent
			nextIndent = 0
			var tmpPos int32
			_ = tmpPos
			tmpPos = nextLinePos
			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if tmpPos > 0 {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = (tmpPos - 1)
					var _hx_tmp_73 int32
					_ = _hx_tmp_73
					var _hx_tmp_74 bool
					_ = _hx_tmp_74
					if tmp_pos_1 < 0 {
						_hx_tmp_74 = true
					} else {
						var _hx_tmp_75 int32
						_ = _hx_tmp_75
						_hx_tmp_75 = tmp_pos_1
						var _hx_tmp_76 int32
						_ = _hx_tmp_76
						_hx_tmp_76 = _hx_tmp_75
						_hx_tmp_74 = (_hx_tmp_76 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_74 {
						_hx_tmp_73 = -1
					} else {
						_hx_tmp_73 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
					}

					if (_hx_tmp_73) != 32 {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = (tmpPos - 1)
						var _hx_tmp_77 int32
						_ = _hx_tmp_77
						var _hx_tmp_78 bool
						_ = _hx_tmp_78
						if tmp_pos_2 < 0 {
							_hx_tmp_78 = true
						} else {
							var _hx_tmp_79 int32
							_ = _hx_tmp_79
							_hx_tmp_79 = tmp_pos_2
							var _hx_tmp_80 int32
							_ = _hx_tmp_80
							_hx_tmp_80 = _hx_tmp_79
							_hx_tmp_78 = (_hx_tmp_80 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_78 {
							_hx_tmp_77 = -1
						} else {
							_hx_tmp_77 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_2))))
						}

						tmp_tmp_1 = ((_hx_tmp_77) == 9)
					} else {
						tmp_tmp_1 = true
					}
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				nextIndent++
				tmpPos--
			}

			var _hx_tmp_81 int32
			_ = _hx_tmp_81
			if multilineIndent != -1 {
				_hx_tmp_81 = multilineIndent
			} else {
				_hx_tmp_81 = (this.Hx_Field_column - 1)
			}

			var expectedIndent int32
			_ = expectedIndent
			expectedIndent = _hx_tmp_81
			if nextIndent != expectedIndent {
				return nil
			}

			var tmp1 bool
			_ = tmp1
			if (nextChar == 45) && ((nextLinePos + 1) < this.Hx_Field_length) {
				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (nextLinePos + 1)
				var _hx_tmp_82 int32
				_ = _hx_tmp_82
				var _hx_tmp_83 bool
				_ = _hx_tmp_83
				if tmp_pos_1 < 0 {
					_hx_tmp_83 = true
				} else {
					var _hx_tmp_84 int32
					_ = _hx_tmp_84
					_hx_tmp_84 = tmp_pos_1
					var _hx_tmp_85 int32
					_ = _hx_tmp_85
					_hx_tmp_85 = _hx_tmp_84
					_hx_tmp_83 = (_hx_tmp_85 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_83 {
					_hx_tmp_82 = -1
				} else {
					_hx_tmp_82 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
				}

				tmp1 = ((_hx_tmp_82) == 62)
			} else {
				tmp1 = false
			}

			if tmp1 {
				return nil
			}

			if nextChar == 43 {
				return nil
			}

			if this.VTable.Hx_Field_isAssignStart(nextLinePos, false) {
				return nil
			}

			if this.VTable.Hx_Field_isCallStart(nextLinePos) {
				return nil
			}

			if this.VTable.Hx_Field_isLabelStart(nextLinePos) {
				return nil
			}

			if this.VTable.Hx_Field_isIfStart(nextLinePos) {
				return nil
			}

			if this.VTable.Hx_Field_isWhenStart(nextLinePos) {
				return nil
			}

			if this.VTable.Hx_Field_isIdentifierExpressionStart(nextLinePos, true) {
				return nil
			}

			if (((nextChar >= 97) && (nextChar <= 122)) || ((nextChar >= 65) && (nextChar <= 90))) || (nextChar == 95) {
				var wordEnd int32
				_ = wordEnd
				wordEnd = nextLinePos
				for {
					var tmp_tmp_1 bool
					_ = tmp_tmp_1
					if wordEnd < this.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var _hx_tmp_86 int32
						_ = _hx_tmp_86
						var _hx_tmp_87 bool
						_ = _hx_tmp_87
						if wordEnd < 0 {
							_hx_tmp_87 = true
						} else {
							var _hx_tmp_88 int32
							_ = _hx_tmp_88
							_hx_tmp_88 = wordEnd
							var _hx_tmp_89 int32
							_ = _hx_tmp_89
							_hx_tmp_89 = _hx_tmp_88
							_hx_tmp_87 = (_hx_tmp_89 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_87 {
							_hx_tmp_86 = -1
						} else {
							_hx_tmp_86 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(wordEnd))))
						}

						var tmp_c_1 int32
						_ = tmp_c_1
						tmp_c_1 = _hx_tmp_86
						tmp_tmp_1 = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
					} else {
						tmp_tmp_1 = false
					}

					if !tmp_tmp_1 {
						break
					}

					wordEnd++
				}

				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (wordEnd - nextLinePos), Valid: true}

				var _hx_tmp_90 HxArray[int32]
				_ = _hx_tmp_90
				_hx_tmp_90 = tmp_v_2
				var _hx_tmp_91 int32
				_ = _hx_tmp_91
				_hx_tmp_91 = nextLinePos
				var _hx_tmp_92 int32
				_ = _hx_tmp_92
				if (_hx_reserved_len).Valid != false {
					var _hx_tmp_93 int32
					_ = _hx_tmp_93
					_hx_tmp_93 = nextLinePos
					var _hx_tmp_94 int32
					_ = _hx_tmp_94
					_hx_tmp_94 = _hx_tmp_93
					_hx_tmp_92 = (_hx_tmp_94 + (_hx_reserved_len).Value)
				} else {
					_hx_tmp_92 = HxMakeArrayView[any](tmp_v_2).Len()
				}

				var kw string
				_ = kw
				kw = Hx_Field_loreline_utf8vector_slice(_hx_tmp_90, _hx_tmp_91, _hx_tmp_92)
				var _hx_tmp_95 bool
				_ = _hx_tmp_95
				if ((((kw != "null") && (kw != "true")) && (kw != "false")) && (kw != "is")) && (kw != "not") {
					_hx_tmp_95 = func() func(_hx_param_key string) bool {
						var _hx_closure_src_138 func(_hx_param_key any) bool
						_ = _hx_closure_src_138
						_hx_closure_src_138 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_138(((any)(_hx_param_key)))
						}
					}()(kw)
				} else {
					_hx_tmp_95 = false
				}

				if _hx_tmp_95 {
					return nil
				}
			}
		} else {
			if !hasSameLineContent {
				return nil
			}
		}
	}

	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var attachments HxArray[Hx_Enum_loreline_lstringattachment]
	_ = attachments
	attachments = HxMakeArray[Hx_Enum_loreline_lstringattachment]()
	var startLine int32
	_ = startLine
	startLine = this.Hx_Field_line
	var startColumn int32
	_ = startColumn
	startColumn = this.Hx_Field_column
	var startPos int32
	_ = startPos
	startPos = this.Hx_Field_pos
	var escaped bool
	_ = escaped
	escaped = false
	var tagStart int32
	_ = tagStart
	tagStart = -1
	var tagIsClosing bool
	_ = tagIsClosing
	tagIsClosing = false
	var currentColumn int32
	_ = currentColumn
	currentColumn = this.Hx_Field_column
	var currentLine int32
	_ = currentLine
	currentLine = this.Hx_Field_line
	var valid bool
	_ = valid
	valid = true
	var hasContent bool
	_ = hasContent
	hasContent = false
	for this.Hx_Field_pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
		var _hx_tmp_96 int32
		_ = _hx_tmp_96
		var _hx_tmp_97 bool
		_ = _hx_tmp_97
		if tmp_pos_1 < 0 {
			_hx_tmp_97 = true
		} else {
			var _hx_tmp_98 int32
			_ = _hx_tmp_98
			_hx_tmp_98 = tmp_pos_1
			var _hx_tmp_99 int32
			_ = _hx_tmp_99
			_hx_tmp_99 = _hx_tmp_98
			_hx_tmp_97 = (_hx_tmp_99 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_97 {
			_hx_tmp_96 = -1
		} else {
			_hx_tmp_96 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
		}

		var tmp_c_1 int32
		_ = tmp_c_1
		tmp_c_1 = _hx_tmp_96
		var isSpace bool
		_ = isSpace
		isSpace = ((tmp_c_1 == 32) || (tmp_c_1 == 9))
		if !hasContent {
			if (!isSpace && (tagStart == -1)) && ((tmp_c_1 != 60) || !allowTags) {
				hasContent = true
			}
		}

		if escaped {
			buf.VTable.Hx_Field_addChar(92)
			buf.VTable.Hx_Field_addChar(tmp_c_1)
			escaped = false
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_100 int32
					_ = _hx_tmp_100
					_hx_tmp_100 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_100
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_101 int32
					_ = _hx_tmp_101
					var _hx_tmp_102 bool
					_ = _hx_tmp_102
					if tmp_pos_2 < 0 {
						_hx_tmp_102 = true
					} else {
						var _hx_tmp_103 int32
						_ = _hx_tmp_103
						_hx_tmp_103 = tmp_pos_2
						var _hx_tmp_104 int32
						_ = _hx_tmp_104
						_hx_tmp_104 = _hx_tmp_103
						_hx_tmp_102 = (_hx_tmp_104 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_102 {
						_hx_tmp_101 = -1
					} else {
						_hx_tmp_101 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_101) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

			currentColumn = (currentColumn + 2)
		} else {
			if tmp_c_1 == 92 {
				escaped = true
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_105 int32
						_ = _hx_tmp_105
						_hx_tmp_105 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_105
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_106 int32
						_ = _hx_tmp_106
						var _hx_tmp_107 bool
						_ = _hx_tmp_107
						if tmp_pos_2 < 0 {
							_hx_tmp_107 = true
						} else {
							var _hx_tmp_108 int32
							_ = _hx_tmp_108
							_hx_tmp_108 = tmp_pos_2
							var _hx_tmp_109 int32
							_ = _hx_tmp_109
							_hx_tmp_109 = _hx_tmp_108
							_hx_tmp_107 = (_hx_tmp_109 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_107 {
							_hx_tmp_106 = -1
						} else {
							_hx_tmp_106 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_106) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}
			} else {
				var _hx_tmp_110 bool
				_ = _hx_tmp_110
				if ((tagStart == -1) && isSpace) && !hasContent {
					_hx_tmp_110 = (attachments.Len() == 0)
				} else {
					_hx_tmp_110 = false
				}

				if _hx_tmp_110 {
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_111 int32
						_ = _hx_tmp_111
						_hx_tmp_111 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_111
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_112 int32
						_ = _hx_tmp_112
						var _hx_tmp_113 bool
						_ = _hx_tmp_113
						if tmp_pos_2 < 0 {
							_hx_tmp_113 = true
						} else {
							var _hx_tmp_114 int32
							_ = _hx_tmp_114
							_hx_tmp_114 = tmp_pos_2
							var _hx_tmp_115 int32
							_ = _hx_tmp_115
							_hx_tmp_115 = _hx_tmp_114
							_hx_tmp_113 = (_hx_tmp_115 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_113 {
							_hx_tmp_112 = -1
						} else {
							_hx_tmp_112 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_112) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				} else {
					if (tagStart == -1) && (tmp_c_1 == 123) {
						break
					} else {
						var tmp bool
						_ = tmp
						if tagStart == -1 {
							if !((tmp_c_1 == 10) || (tmp_c_1 == 13)) {
								if (tmp_c_1 == 47) && (this.Hx_Field_pos < (this.Hx_Field_length - 1)) {
									var tmp_v_2 HxArray[int32]
									_ = tmp_v_2
									tmp_v_2 = this.Hx_Field_input
									var tmp_pos_2 int32
									_ = tmp_pos_2
									tmp_pos_2 = (this.Hx_Field_pos + 1)
									var _hx_tmp_116 int32
									_ = _hx_tmp_116
									var _hx_tmp_117 bool
									_ = _hx_tmp_117
									if tmp_pos_2 < 0 {
										_hx_tmp_117 = true
									} else {
										var _hx_tmp_118 int32
										_ = _hx_tmp_118
										_hx_tmp_118 = tmp_pos_2
										var _hx_tmp_119 int32
										_ = _hx_tmp_119
										_hx_tmp_119 = _hx_tmp_118
										_hx_tmp_117 = (_hx_tmp_119 >= HxMakeArrayView[any](tmp_v_2).Len())
									}

									if _hx_tmp_117 {
										_hx_tmp_116 = -1
									} else {
										_hx_tmp_116 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
									}

									tmp = ((_hx_tmp_116) == 47)
								} else {
									tmp = false
								}
							} else {
								tmp = true
							}
						} else {
							tmp = false
						}

						if tmp {
							if multilineIndent != -1 {
								if this.VTable.Hx_Field_isContinuingMultilineText(this.Hx_Field_pos, multilineIndent) {
									buf.VTable.Hx_Field_addChar(tmp_c_1)
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_120 int32
											_ = _hx_tmp_120
											_hx_tmp_120 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_120
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_2 HxArray[int32]
											_ = tmp_v_2
											tmp_v_2 = this.Hx_Field_input
											var tmp_pos_2 int32
											_ = tmp_pos_2
											tmp_pos_2 = this.Hx_Field_pos
											var _hx_tmp_121 int32
											_ = _hx_tmp_121
											var _hx_tmp_122 bool
											_ = _hx_tmp_122
											if tmp_pos_2 < 0 {
												_hx_tmp_122 = true
											} else {
												var _hx_tmp_123 int32
												_ = _hx_tmp_123
												_hx_tmp_123 = tmp_pos_2
												var _hx_tmp_124 int32
												_ = _hx_tmp_124
												_hx_tmp_124 = _hx_tmp_123
												_hx_tmp_122 = (_hx_tmp_124 >= HxMakeArrayView[any](tmp_v_2).Len())
											}

											if _hx_tmp_122 {
												_hx_tmp_121 = -1
											} else {
												_hx_tmp_121 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
											}

											if (_hx_tmp_121) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									var tmp_tmp_1 bool
									_ = tmp_tmp_1
									if tmp_c_1 == 13 {
										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = this.Hx_Field_input
										var tmp_pos_2 int32
										_ = tmp_pos_2
										tmp_pos_2 = this.Hx_Field_pos
										var _hx_tmp_125 int32
										_ = _hx_tmp_125
										var _hx_tmp_126 bool
										_ = _hx_tmp_126
										if tmp_pos_2 < 0 {
											_hx_tmp_126 = true
										} else {
											var _hx_tmp_127 int32
											_ = _hx_tmp_127
											_hx_tmp_127 = tmp_pos_2
											var _hx_tmp_128 int32
											_ = _hx_tmp_128
											_hx_tmp_128 = _hx_tmp_127
											_hx_tmp_126 = (_hx_tmp_128 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_126 {
											_hx_tmp_125 = -1
										} else {
											_hx_tmp_125 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
										}

										tmp_tmp_1 = ((_hx_tmp_125) == 10)
									} else {
										tmp_tmp_1 = false
									}

									if tmp_tmp_1 {
										buf.VTable.Hx_Field_addChar(10)
										{
											var _g int32
											_ = _g
											_g = 0
											var _g1 int32
											_ = _g1
											_g1 = 1
											for _g < _g1 {
												var _hx_tmp_129 int32
												_ = _hx_tmp_129
												_hx_tmp_129 = _g
												_g = (_g + 1)
												var _hx_reserved__ int32
												_ = _hx_reserved__
												_hx_reserved__ = _hx_tmp_129
												if this.Hx_Field_pos >= this.Hx_Field_length {
													break
												}

												var tmp_v_2 HxArray[int32]
												_ = tmp_v_2
												tmp_v_2 = this.Hx_Field_input
												var tmp_pos_2 int32
												_ = tmp_pos_2
												tmp_pos_2 = this.Hx_Field_pos
												var _hx_tmp_130 int32
												_ = _hx_tmp_130
												var _hx_tmp_131 bool
												_ = _hx_tmp_131
												if tmp_pos_2 < 0 {
													_hx_tmp_131 = true
												} else {
													var _hx_tmp_132 int32
													_ = _hx_tmp_132
													_hx_tmp_132 = tmp_pos_2
													var _hx_tmp_133 int32
													_ = _hx_tmp_133
													_hx_tmp_133 = _hx_tmp_132
													_hx_tmp_131 = (_hx_tmp_133 >= HxMakeArrayView[any](tmp_v_2).Len())
												}

												if _hx_tmp_131 {
													_hx_tmp_130 = -1
												} else {
													_hx_tmp_130 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
												}

												if (_hx_tmp_130) == 10 {
													this.Hx_Field_line++
													this.Hx_Field_column = 1
												} else {
													this.Hx_Field_column++
												}

												this.Hx_Field_pos++
											}
										}
									}

									currentColumn = this.Hx_Field_column
									currentLine = this.Hx_Field_line
								} else {
									break
								}
							} else {
								break
							}
						} else {
							var tmp_tmp_1 bool
							_ = tmp_tmp_1
							if tagStart == -1 {
								if (tmp_c_1 == 47) && (this.Hx_Field_pos < (this.Hx_Field_length - 1)) {
									var tmp_v_2 HxArray[int32]
									_ = tmp_v_2
									tmp_v_2 = this.Hx_Field_input
									var tmp_pos_2 int32
									_ = tmp_pos_2
									tmp_pos_2 = (this.Hx_Field_pos + 1)
									var _hx_tmp_134 int32
									_ = _hx_tmp_134
									var _hx_tmp_135 bool
									_ = _hx_tmp_135
									if tmp_pos_2 < 0 {
										_hx_tmp_135 = true
									} else {
										var _hx_tmp_136 int32
										_ = _hx_tmp_136
										_hx_tmp_136 = tmp_pos_2
										var _hx_tmp_137 int32
										_ = _hx_tmp_137
										_hx_tmp_137 = _hx_tmp_136
										_hx_tmp_135 = (_hx_tmp_137 >= HxMakeArrayView[any](tmp_v_2).Len())
									}

									if _hx_tmp_135 {
										_hx_tmp_134 = -1
									} else {
										_hx_tmp_134 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
									}

									tmp_tmp_1 = ((_hx_tmp_134) == 42)
								} else {
									tmp_tmp_1 = false
								}
							} else {
								tmp_tmp_1 = false
							}

							if tmp_tmp_1 {
								buf.VTable.Hx_Field_addChar(47)
								buf.VTable.Hx_Field_addChar(42)
								{
									var _g int32
									_ = _g
									_g = 0
									var _g1 int32
									_ = _g1
									_g1 = 2
									for _g < _g1 {
										var _hx_tmp_138 int32
										_ = _hx_tmp_138
										_hx_tmp_138 = _g
										_g = (_g + 1)
										var _hx_reserved__ int32
										_ = _hx_reserved__
										_hx_reserved__ = _hx_tmp_138
										if this.Hx_Field_pos >= this.Hx_Field_length {
											break
										}

										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = this.Hx_Field_input
										var tmp_pos_2 int32
										_ = tmp_pos_2
										tmp_pos_2 = this.Hx_Field_pos
										var _hx_tmp_139 int32
										_ = _hx_tmp_139
										var _hx_tmp_140 bool
										_ = _hx_tmp_140
										if tmp_pos_2 < 0 {
											_hx_tmp_140 = true
										} else {
											var _hx_tmp_141 int32
											_ = _hx_tmp_141
											_hx_tmp_141 = tmp_pos_2
											var _hx_tmp_142 int32
											_ = _hx_tmp_142
											_hx_tmp_142 = _hx_tmp_141
											_hx_tmp_140 = (_hx_tmp_142 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_140 {
											_hx_tmp_139 = -1
										} else {
											_hx_tmp_139 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
										}

										if (_hx_tmp_139) == 10 {
											this.Hx_Field_line++
											this.Hx_Field_column = 1
										} else {
											this.Hx_Field_column++
										}

										this.Hx_Field_pos++
									}
								}

								var commentClosed bool
								_ = commentClosed
								commentClosed = false
								for this.Hx_Field_pos < this.Hx_Field_length {
									var tmp_tmp_2 bool
									_ = tmp_tmp_2
									var tmp_v_2 HxArray[int32]
									_ = tmp_v_2
									tmp_v_2 = this.Hx_Field_input
									var tmp_pos_2 int32
									_ = tmp_pos_2
									tmp_pos_2 = this.Hx_Field_pos
									var _hx_tmp_143 bool
									_ = _hx_tmp_143
									var _hx_tmp_144 int32
									_ = _hx_tmp_144
									var _hx_tmp_145 bool
									_ = _hx_tmp_145
									if tmp_pos_2 < 0 {
										_hx_tmp_145 = true
									} else {
										var _hx_tmp_146 int32
										_ = _hx_tmp_146
										_hx_tmp_146 = tmp_pos_2
										var _hx_tmp_147 int32
										_ = _hx_tmp_147
										_hx_tmp_147 = _hx_tmp_146
										_hx_tmp_145 = (_hx_tmp_147 >= HxMakeArrayView[any](tmp_v_2).Len())
									}

									if _hx_tmp_145 {
										_hx_tmp_144 = -1
									} else {
										_hx_tmp_144 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
									}

									if (_hx_tmp_144) == 42 {
										_hx_tmp_143 = ((this.Hx_Field_pos + 1) < this.Hx_Field_length)
									} else {
										_hx_tmp_143 = false
									}

									if _hx_tmp_143 {
										var tmp_v_3 HxArray[int32]
										_ = tmp_v_3
										tmp_v_3 = this.Hx_Field_input
										var tmp_pos_3 int32
										_ = tmp_pos_3
										tmp_pos_3 = (this.Hx_Field_pos + 1)
										var _hx_tmp_148 int32
										_ = _hx_tmp_148
										var _hx_tmp_149 bool
										_ = _hx_tmp_149
										if tmp_pos_3 < 0 {
											_hx_tmp_149 = true
										} else {
											var _hx_tmp_150 int32
											_ = _hx_tmp_150
											_hx_tmp_150 = tmp_pos_3
											var _hx_tmp_151 int32
											_ = _hx_tmp_151
											_hx_tmp_151 = _hx_tmp_150
											_hx_tmp_149 = (_hx_tmp_151 >= HxMakeArrayView[any](tmp_v_3).Len())
										}

										if _hx_tmp_149 {
											_hx_tmp_148 = -1
										} else {
											_hx_tmp_148 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
										}

										tmp_tmp_2 = ((_hx_tmp_148) == 47)
									} else {
										tmp_tmp_2 = false
									}

									if tmp_tmp_2 {
										buf.VTable.Hx_Field_addChar(42)
										buf.VTable.Hx_Field_addChar(47)
										{
											var _g int32
											_ = _g
											_g = 0
											var _g1 int32
											_ = _g1
											_g1 = 2
											for _g < _g1 {
												var _hx_tmp_152 int32
												_ = _hx_tmp_152
												_hx_tmp_152 = _g
												_g = (_g + 1)
												var _hx_reserved__ int32
												_ = _hx_reserved__
												_hx_reserved__ = _hx_tmp_152
												if this.Hx_Field_pos >= this.Hx_Field_length {
													break
												}

												var tmp_v_3 HxArray[int32]
												_ = tmp_v_3
												tmp_v_3 = this.Hx_Field_input
												var tmp_pos_3 int32
												_ = tmp_pos_3
												tmp_pos_3 = this.Hx_Field_pos
												var _hx_tmp_153 int32
												_ = _hx_tmp_153
												var _hx_tmp_154 bool
												_ = _hx_tmp_154
												if tmp_pos_3 < 0 {
													_hx_tmp_154 = true
												} else {
													var _hx_tmp_155 int32
													_ = _hx_tmp_155
													_hx_tmp_155 = tmp_pos_3
													var _hx_tmp_156 int32
													_ = _hx_tmp_156
													_hx_tmp_156 = _hx_tmp_155
													_hx_tmp_154 = (_hx_tmp_156 >= HxMakeArrayView[any](tmp_v_3).Len())
												}

												if _hx_tmp_154 {
													_hx_tmp_153 = -1
												} else {
													_hx_tmp_153 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
												}

												if (_hx_tmp_153) == 10 {
													this.Hx_Field_line++
													this.Hx_Field_column = 1
												} else {
													this.Hx_Field_column++
												}

												this.Hx_Field_pos++
											}
										}

										commentClosed = true
										break
									}

									var v1 HxArray[int32]
									_ = v1
									v1 = this.Hx_Field_input
									var pos1 int32
									_ = pos1
									pos1 = this.Hx_Field_pos
									var _hx_tmp_157 int32
									_ = _hx_tmp_157
									var _hx_tmp_158 bool
									_ = _hx_tmp_158
									if pos1 < 0 {
										_hx_tmp_158 = true
									} else {
										var _hx_tmp_159 int32
										_ = _hx_tmp_159
										_hx_tmp_159 = pos1
										var _hx_tmp_160 int32
										_ = _hx_tmp_160
										_hx_tmp_160 = _hx_tmp_159
										_hx_tmp_158 = (_hx_tmp_160 >= HxMakeArrayView[any](v1).Len())
									}

									if _hx_tmp_158 {
										_hx_tmp_157 = -1
									} else {
										_hx_tmp_157 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v1).Get(pos1))))
									}

									buf.VTable.Hx_Field_addChar(_hx_tmp_157)
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_161 int32
											_ = _hx_tmp_161
											_hx_tmp_161 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_161
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_3 HxArray[int32]
											_ = tmp_v_3
											tmp_v_3 = this.Hx_Field_input
											var tmp_pos_3 int32
											_ = tmp_pos_3
											tmp_pos_3 = this.Hx_Field_pos
											var _hx_tmp_162 int32
											_ = _hx_tmp_162
											var _hx_tmp_163 bool
											_ = _hx_tmp_163
											if tmp_pos_3 < 0 {
												_hx_tmp_163 = true
											} else {
												var _hx_tmp_164 int32
												_ = _hx_tmp_164
												_hx_tmp_164 = tmp_pos_3
												var _hx_tmp_165 int32
												_ = _hx_tmp_165
												_hx_tmp_165 = _hx_tmp_164
												_hx_tmp_163 = (_hx_tmp_165 >= HxMakeArrayView[any](tmp_v_3).Len())
											}

											if _hx_tmp_163 {
												_hx_tmp_162 = -1
											} else {
												_hx_tmp_162 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
											}

											if (_hx_tmp_162) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}
								}

								currentColumn = this.Hx_Field_column
								currentLine = this.Hx_Field_line
								if !commentClosed {
									this.VTable.Hx_Field_error("Unterminated multiline comment", false)
								}
							} else {
								var _hx_tmp_166 bool
								_ = _hx_tmp_166
								if (tagStart == -1) && !isValue {
									_hx_tmp_166 = this.VTable.Hx_Field_isIfStart(this.Hx_Field_pos)
								} else {
									_hx_tmp_166 = false
								}

								if _hx_tmp_166 {
									break
								} else {
									var tmp_tmp_2 bool
									_ = tmp_tmp_2
									if ((tagStart == -1) && (tmp_c_1 == 45)) && (this.Hx_Field_pos < (this.Hx_Field_length - 1)) {
										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = this.Hx_Field_input
										var tmp_pos_2 int32
										_ = tmp_pos_2
										tmp_pos_2 = (this.Hx_Field_pos + 1)
										var _hx_tmp_167 int32
										_ = _hx_tmp_167
										var _hx_tmp_168 bool
										_ = _hx_tmp_168
										if tmp_pos_2 < 0 {
											_hx_tmp_168 = true
										} else {
											var _hx_tmp_169 int32
											_ = _hx_tmp_169
											_hx_tmp_169 = tmp_pos_2
											var _hx_tmp_170 int32
											_ = _hx_tmp_170
											_hx_tmp_170 = _hx_tmp_169
											_hx_tmp_168 = (_hx_tmp_170 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_168 {
											_hx_tmp_167 = -1
										} else {
											_hx_tmp_167 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
										}

										tmp_tmp_2 = ((_hx_tmp_167) == 62)
									} else {
										tmp_tmp_2 = false
									}

									var _hx_tmp_171 bool
									_ = _hx_tmp_171
									if tmp_tmp_2 {
										_hx_tmp_171 = this.VTable.Hx_Field_isTransitionStart(this.Hx_Field_pos)
									} else {
										_hx_tmp_171 = false
									}

									if _hx_tmp_171 {
										break
									} else {
										if ((tagStart == -1) && isValue) && (((tmp_c_1 == 44) || (tmp_c_1 == 93)) || (tmp_c_1 == 125)) {
											break
										} else {
											if allowTags && (tmp_c_1 == 60) {
												if tagStart != -1 {
													this.VTable.Hx_Field_error("Unexpected < inside tag", true)
												}

												var nextChar int32
												_ = nextChar
												if (this.Hx_Field_pos + 1) < this.Hx_Field_length {
													var tmp_v_2 HxArray[int32]
													_ = tmp_v_2
													tmp_v_2 = this.Hx_Field_input
													var tmp_pos_2 int32
													_ = tmp_pos_2
													tmp_pos_2 = (this.Hx_Field_pos + 1)
													var _hx_tmp_172 int32
													_ = _hx_tmp_172
													var _hx_tmp_173 bool
													_ = _hx_tmp_173
													if tmp_pos_2 < 0 {
														_hx_tmp_173 = true
													} else {
														var _hx_tmp_174 int32
														_ = _hx_tmp_174
														_hx_tmp_174 = tmp_pos_2
														var _hx_tmp_175 int32
														_ = _hx_tmp_175
														_hx_tmp_175 = _hx_tmp_174
														_hx_tmp_173 = (_hx_tmp_175 >= HxMakeArrayView[any](tmp_v_2).Len())
													}

													if _hx_tmp_173 {
														_hx_tmp_172 = -1
													} else {
														_hx_tmp_172 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
													}

													nextChar = _hx_tmp_172
												} else {
													nextChar = 0
												}

												tagIsClosing = (nextChar == 47)
												if this.VTable.Hx_Field_isTagStart(this.Hx_Field_pos) {
													tagStart = buf.VTable.Hx_Field_get_length()
												}

												buf.VTable.Hx_Field_addChar(tmp_c_1)
												{
													var _g int32
													_ = _g
													_g = 0
													var _g1 int32
													_ = _g1
													_g1 = 1
													for _g < _g1 {
														var _hx_tmp_176 int32
														_ = _hx_tmp_176
														_hx_tmp_176 = _g
														_g = (_g + 1)
														var _hx_reserved__ int32
														_ = _hx_reserved__
														_hx_reserved__ = _hx_tmp_176
														if this.Hx_Field_pos >= this.Hx_Field_length {
															break
														}

														var tmp_v_2 HxArray[int32]
														_ = tmp_v_2
														tmp_v_2 = this.Hx_Field_input
														var tmp_pos_2 int32
														_ = tmp_pos_2
														tmp_pos_2 = this.Hx_Field_pos
														var _hx_tmp_177 int32
														_ = _hx_tmp_177
														var _hx_tmp_178 bool
														_ = _hx_tmp_178
														if tmp_pos_2 < 0 {
															_hx_tmp_178 = true
														} else {
															var _hx_tmp_179 int32
															_ = _hx_tmp_179
															_hx_tmp_179 = tmp_pos_2
															var _hx_tmp_180 int32
															_ = _hx_tmp_180
															_hx_tmp_180 = _hx_tmp_179
															_hx_tmp_178 = (_hx_tmp_180 >= HxMakeArrayView[any](tmp_v_2).Len())
														}

														if _hx_tmp_178 {
															_hx_tmp_177 = -1
														} else {
															_hx_tmp_177 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
														}

														if (_hx_tmp_177) == 10 {
															this.Hx_Field_line++
															this.Hx_Field_column = 1
														} else {
															this.Hx_Field_column++
														}

														this.Hx_Field_pos++
													}
												}

												currentColumn++
											} else {
												if allowTags && (tmp_c_1 == 62) {
													buf.VTable.Hx_Field_addChar(tmp_c_1)
													{
														var _g int32
														_ = _g
														_g = 0
														var _g1 int32
														_ = _g1
														_g1 = 1
														for _g < _g1 {
															var _hx_tmp_181 int32
															_ = _hx_tmp_181
															_hx_tmp_181 = _g
															_g = (_g + 1)
															var _hx_reserved__ int32
															_ = _hx_reserved__
															_hx_reserved__ = _hx_tmp_181
															if this.Hx_Field_pos >= this.Hx_Field_length {
																break
															}

															var tmp_v_2 HxArray[int32]
															_ = tmp_v_2
															tmp_v_2 = this.Hx_Field_input
															var tmp_pos_2 int32
															_ = tmp_pos_2
															tmp_pos_2 = this.Hx_Field_pos
															var _hx_tmp_182 int32
															_ = _hx_tmp_182
															var _hx_tmp_183 bool
															_ = _hx_tmp_183
															if tmp_pos_2 < 0 {
																_hx_tmp_183 = true
															} else {
																var _hx_tmp_184 int32
																_ = _hx_tmp_184
																_hx_tmp_184 = tmp_pos_2
																var _hx_tmp_185 int32
																_ = _hx_tmp_185
																_hx_tmp_185 = _hx_tmp_184
																_hx_tmp_183 = (_hx_tmp_185 >= HxMakeArrayView[any](tmp_v_2).Len())
															}

															if _hx_tmp_183 {
																_hx_tmp_182 = -1
															} else {
																_hx_tmp_182 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
															}

															if (_hx_tmp_182) == 10 {
																this.Hx_Field_line++
																this.Hx_Field_column = 1
															} else {
																this.Hx_Field_column++
															}

															this.Hx_Field_pos++
														}
													}

													currentColumn++
													if tagStart != -1 {
														var _hx_tmp_186 HxArray[Hx_Enum_loreline_lstringattachment]
														_ = _hx_tmp_186
														_hx_tmp_186 = attachments
														var _hx_tmp_187 bool
														_ = _hx_tmp_187
														_hx_tmp_187 = tagIsClosing
														var _hx_tmp_188 int32
														_ = _hx_tmp_188
														_hx_tmp_188 = tagStart
														_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_186, Hx_Enum_loreline_lstringattachment_Tag{_hx_tmp_187, _hx_tmp_188, (buf.VTable.Hx_Field_get_length() - tagStart)}))
														tagStart = -1
													}
												} else {
													if ((tagStart == -1) && (tmp_c_1 == 35)) && !escaped {
														var nextChar int32
														_ = nextChar
														if (this.Hx_Field_pos + 1) < this.Hx_Field_length {
															var tmp_v_2 HxArray[int32]
															_ = tmp_v_2
															tmp_v_2 = this.Hx_Field_input
															var tmp_pos_2 int32
															_ = tmp_pos_2
															tmp_pos_2 = (this.Hx_Field_pos + 1)
															var _hx_tmp_189 int32
															_ = _hx_tmp_189
															var _hx_tmp_190 bool
															_ = _hx_tmp_190
															if tmp_pos_2 < 0 {
																_hx_tmp_190 = true
															} else {
																var _hx_tmp_191 int32
																_ = _hx_tmp_191
																_hx_tmp_191 = tmp_pos_2
																var _hx_tmp_192 int32
																_ = _hx_tmp_192
																_hx_tmp_192 = _hx_tmp_191
																_hx_tmp_190 = (_hx_tmp_192 >= HxMakeArrayView[any](tmp_v_2).Len())
															}

															if _hx_tmp_190 {
																_hx_tmp_189 = -1
															} else {
																_hx_tmp_189 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
															}

															nextChar = _hx_tmp_189
														} else {
															nextChar = 0
														}

														if nextChar == 35 {
															buf.VTable.Hx_Field_addChar(35)
															buf.VTable.Hx_Field_addChar(35)
															{
																var _g int32
																_ = _g
																_g = 0
																var _g1 int32
																_ = _g1
																_g1 = 1
																for _g < _g1 {
																	var _hx_tmp_193 int32
																	_ = _hx_tmp_193
																	_hx_tmp_193 = _g
																	_g = (_g + 1)
																	var _hx_reserved__ int32
																	_ = _hx_reserved__
																	_hx_reserved__ = _hx_tmp_193
																	if this.Hx_Field_pos >= this.Hx_Field_length {
																		break
																	}

																	var tmp_v_2 HxArray[int32]
																	_ = tmp_v_2
																	tmp_v_2 = this.Hx_Field_input
																	var tmp_pos_2 int32
																	_ = tmp_pos_2
																	tmp_pos_2 = this.Hx_Field_pos
																	var _hx_tmp_194 int32
																	_ = _hx_tmp_194
																	var _hx_tmp_195 bool
																	_ = _hx_tmp_195
																	if tmp_pos_2 < 0 {
																		_hx_tmp_195 = true
																	} else {
																		var _hx_tmp_196 int32
																		_ = _hx_tmp_196
																		_hx_tmp_196 = tmp_pos_2
																		var _hx_tmp_197 int32
																		_ = _hx_tmp_197
																		_hx_tmp_197 = _hx_tmp_196
																		_hx_tmp_195 = (_hx_tmp_197 >= HxMakeArrayView[any](tmp_v_2).Len())
																	}

																	if _hx_tmp_195 {
																		_hx_tmp_194 = -1
																	} else {
																		_hx_tmp_194 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																	}

																	if (_hx_tmp_194) == 10 {
																		this.Hx_Field_line++
																		this.Hx_Field_column = 1
																	} else {
																		this.Hx_Field_column++
																	}

																	this.Hx_Field_pos++
																}
															}

															{
																var _g2 int32
																_ = _g2
																_g2 = 0
																var _g3 int32
																_ = _g3
																_g3 = 1
																for _g2 < _g3 {
																	var _hx_tmp_198 int32
																	_ = _hx_tmp_198
																	_hx_tmp_198 = _g2
																	_g2 = (_g2 + 1)
																	var _hx_reserved__ int32
																	_ = _hx_reserved__
																	_hx_reserved__ = _hx_tmp_198
																	if this.Hx_Field_pos >= this.Hx_Field_length {
																		break
																	}

																	var tmp_v_2 HxArray[int32]
																	_ = tmp_v_2
																	tmp_v_2 = this.Hx_Field_input
																	var tmp_pos_2 int32
																	_ = tmp_pos_2
																	tmp_pos_2 = this.Hx_Field_pos
																	var _hx_tmp_199 int32
																	_ = _hx_tmp_199
																	var _hx_tmp_200 bool
																	_ = _hx_tmp_200
																	if tmp_pos_2 < 0 {
																		_hx_tmp_200 = true
																	} else {
																		var _hx_tmp_201 int32
																		_ = _hx_tmp_201
																		_hx_tmp_201 = tmp_pos_2
																		var _hx_tmp_202 int32
																		_ = _hx_tmp_202
																		_hx_tmp_202 = _hx_tmp_201
																		_hx_tmp_200 = (_hx_tmp_202 >= HxMakeArrayView[any](tmp_v_2).Len())
																	}

																	if _hx_tmp_200 {
																		_hx_tmp_199 = -1
																	} else {
																		_hx_tmp_199 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																	}

																	if (_hx_tmp_199) == 10 {
																		this.Hx_Field_line++
																		this.Hx_Field_column = 1
																	} else {
																		this.Hx_Field_column++
																	}

																	this.Hx_Field_pos++
																}
															}

															currentColumn = (currentColumn + 2)
														} else {
															if (((((nextChar >= 97) && (nextChar <= 122)) || ((nextChar >= 65) && (nextChar <= 90))) || (nextChar == 95)) || ((nextChar >= 48) && (nextChar <= 57))) || (nextChar == 45) {
																break
															} else {
																buf.VTable.Hx_Field_addChar(tmp_c_1)
																{
																	var _g int32
																	_ = _g
																	_g = 0
																	var _g1 int32
																	_ = _g1
																	_g1 = 1
																	for _g < _g1 {
																		var _hx_tmp_203 int32
																		_ = _hx_tmp_203
																		_hx_tmp_203 = _g
																		_g = (_g + 1)
																		var _hx_reserved__ int32
																		_ = _hx_reserved__
																		_hx_reserved__ = _hx_tmp_203
																		if this.Hx_Field_pos >= this.Hx_Field_length {
																			break
																		}

																		var tmp_v_2 HxArray[int32]
																		_ = tmp_v_2
																		tmp_v_2 = this.Hx_Field_input
																		var tmp_pos_2 int32
																		_ = tmp_pos_2
																		tmp_pos_2 = this.Hx_Field_pos
																		var _hx_tmp_204 int32
																		_ = _hx_tmp_204
																		var _hx_tmp_205 bool
																		_ = _hx_tmp_205
																		if tmp_pos_2 < 0 {
																			_hx_tmp_205 = true
																		} else {
																			var _hx_tmp_206 int32
																			_ = _hx_tmp_206
																			_hx_tmp_206 = tmp_pos_2
																			var _hx_tmp_207 int32
																			_ = _hx_tmp_207
																			_hx_tmp_207 = _hx_tmp_206
																			_hx_tmp_205 = (_hx_tmp_207 >= HxMakeArrayView[any](tmp_v_2).Len())
																		}

																		if _hx_tmp_205 {
																			_hx_tmp_204 = -1
																		} else {
																			_hx_tmp_204 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																		}

																		if (_hx_tmp_204) == 10 {
																			this.Hx_Field_line++
																			this.Hx_Field_column = 1
																		} else {
																			this.Hx_Field_column++
																		}

																		this.Hx_Field_pos++
																	}
																}

																currentColumn++
															}
														}
													} else {
														if (tmp_c_1 == 36) && !escaped {
															var interpStart int32
															_ = interpStart
															interpStart = buf.VTable.Hx_Field_get_length()
															var interpLine int32
															_ = interpLine
															interpLine = currentLine
															var interpColumn int32
															_ = interpColumn
															interpColumn = currentColumn
															var tokenStartPos int32
															_ = tokenStartPos
															tokenStartPos = this.Hx_Field_pos
															{
																var hx_try_return *Hx_Obj_loreline_token
																_ = hx_try_return
																var hx_try_state int
																_ = hx_try_state

																func() {
																	defer func() {
																		if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
																			if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_lexererror").(*Hx_Obj_loreline_lexererror); _hx_ok {
																				_ = e
																				{
																					if e.Hx_Field_pos == nil {
																						e.Hx_Field_pos = Hx_Obj_loreline_position_CreateInstance(interpLine, interpColumn, this.Hx_Field_pos, nil)
																					}

																					panic(Hx_Field_haxe_exception_thrown(((any)(e))))
																				}
																			} else {
																				panic(_hx_ex)
																			}
																		}
																	}()
																	{
																		{
																			var _g int32
																			_ = _g
																			_g = 0
																			var _g1 int32
																			_ = _g1
																			_g1 = 1
																			for _g < _g1 {
																				var _hx_tmp_208 int32
																				_ = _hx_tmp_208
																				_hx_tmp_208 = _g
																				_g = (_g + 1)
																				var _hx_reserved__ int32
																				_ = _hx_reserved__
																				_hx_reserved__ = _hx_tmp_208
																				if this.Hx_Field_pos >= this.Hx_Field_length {
																					break
																				}

																				var tmp_v_2 HxArray[int32]
																				_ = tmp_v_2
																				tmp_v_2 = this.Hx_Field_input
																				var tmp_pos_2 int32
																				_ = tmp_pos_2
																				tmp_pos_2 = this.Hx_Field_pos
																				var _hx_tmp_209 int32
																				_ = _hx_tmp_209
																				var _hx_tmp_210 bool
																				_ = _hx_tmp_210
																				if tmp_pos_2 < 0 {
																					_hx_tmp_210 = true
																				} else {
																					var _hx_tmp_211 int32
																					_ = _hx_tmp_211
																					_hx_tmp_211 = tmp_pos_2
																					var _hx_tmp_212 int32
																					_ = _hx_tmp_212
																					_hx_tmp_212 = _hx_tmp_211
																					_hx_tmp_210 = (_hx_tmp_212 >= HxMakeArrayView[any](tmp_v_2).Len())
																				}

																				if _hx_tmp_210 {
																					_hx_tmp_209 = -1
																				} else {
																					_hx_tmp_209 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																				}

																				if (_hx_tmp_209) == 10 {
																					this.Hx_Field_line++
																					this.Hx_Field_column = 1
																				} else {
																					this.Hx_Field_column++
																				}

																				this.Hx_Field_pos++
																			}
																		}

																		currentColumn++
																		var tmp_v_2 HxArray[int32]
																		_ = tmp_v_2
																		tmp_v_2 = this.Hx_Field_input
																		var tmp_pos_2 int32
																		_ = tmp_pos_2
																		tmp_pos_2 = this.Hx_Field_pos
																		var _hx_tmp_213 int32
																		_ = _hx_tmp_213
																		var _hx_tmp_214 bool
																		_ = _hx_tmp_214
																		if tmp_pos_2 < 0 {
																			_hx_tmp_214 = true
																		} else {
																			var _hx_tmp_215 int32
																			_ = _hx_tmp_215
																			_hx_tmp_215 = tmp_pos_2
																			var _hx_tmp_216 int32
																			_ = _hx_tmp_216
																			_hx_tmp_216 = _hx_tmp_215
																			_hx_tmp_214 = (_hx_tmp_216 >= HxMakeArrayView[any](tmp_v_2).Len())
																		}

																		if _hx_tmp_214 {
																			_hx_tmp_213 = -1
																		} else {
																			_hx_tmp_213 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																		}

																		if (_hx_tmp_213) == 123 {
																			{
																				var _g int32
																				_ = _g
																				_g = 0
																				var _g1 int32
																				_ = _g1
																				_g1 = 1
																				for _g < _g1 {
																					var _hx_tmp_217 int32
																					_ = _hx_tmp_217
																					_hx_tmp_217 = _g
																					_g = (_g + 1)
																					var _hx_reserved__ int32
																					_ = _hx_reserved__
																					_hx_reserved__ = _hx_tmp_217
																					if this.Hx_Field_pos >= this.Hx_Field_length {
																						break
																					}

																					var tmp_v_3 HxArray[int32]
																					_ = tmp_v_3
																					tmp_v_3 = this.Hx_Field_input
																					var tmp_pos_3 int32
																					_ = tmp_pos_3
																					tmp_pos_3 = this.Hx_Field_pos
																					var _hx_tmp_218 int32
																					_ = _hx_tmp_218
																					var _hx_tmp_219 bool
																					_ = _hx_tmp_219
																					if tmp_pos_3 < 0 {
																						_hx_tmp_219 = true
																					} else {
																						var _hx_tmp_220 int32
																						_ = _hx_tmp_220
																						_hx_tmp_220 = tmp_pos_3
																						var _hx_tmp_221 int32
																						_ = _hx_tmp_221
																						_hx_tmp_221 = _hx_tmp_220
																						_hx_tmp_219 = (_hx_tmp_221 >= HxMakeArrayView[any](tmp_v_3).Len())
																					}

																					if _hx_tmp_219 {
																						_hx_tmp_218 = -1
																					} else {
																						_hx_tmp_218 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
																					}

																					if (_hx_tmp_218) == 10 {
																						this.Hx_Field_line++
																						this.Hx_Field_column = 1
																					} else {
																						this.Hx_Field_column++
																					}

																					this.Hx_Field_pos++
																				}
																			}

																			currentColumn++
																			var interpPos *Hx_Obj_loreline_position
																			_ = interpPos
																			interpPos = Hx_Obj_loreline_position_CreateInstance(interpLine, (interpColumn + 2), this.Hx_Field_pos, nil)
																			var tokens Hx_Typedef_loreline_tokens
																			_ = tokens
																			tokens = this.VTable.Hx_Field_readComplexInterpolation(interpPos)
																			var interpLength int32
																			_ = interpLength
																			interpLength = (this.Hx_Field_pos - tokenStartPos)
																			var _hx_tmp_222 HxArray[Hx_Enum_loreline_lstringattachment]
																			_ = _hx_tmp_222
																			_hx_tmp_222 = attachments
																			_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_222, Hx_Enum_loreline_lstringattachment_Interpolation{true, (tagStart != -1), tokens, interpStart, interpLength}))
																			var tmp_v_3 HxArray[int32]
																			_ = tmp_v_3
																			tmp_v_3 = this.Hx_Field_input
																			var _hx_tmp_223 HxArray[int32]
																			_ = _hx_tmp_223
																			_hx_tmp_223 = tmp_v_3
																			var _hx_tmp_224 int32
																			_ = _hx_tmp_224
																			_hx_tmp_224 = tokenStartPos
																			var _hx_tmp_225 int32
																			_ = _hx_tmp_225
																			if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
																				_hx_tmp_225 = (tokenStartPos + interpLength)
																			} else {
																				_hx_tmp_225 = HxMakeArrayView[any](tmp_v_3).Len()
																			}

																			func() func(x string) {
																				var _hx_closure_src_139 func(x any)
																				_ = _hx_closure_src_139
																				_hx_closure_src_139 = buf.VTable.Hx_Field_add
																				return func(x string) {
																					_hx_closure_src_139(((any)(x)))
																				}
																			}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_223, _hx_tmp_224, _hx_tmp_225))
																		} else {
																			var tmp_v_3 HxArray[int32]
																			_ = tmp_v_3
																			tmp_v_3 = this.Hx_Field_input
																			var tmp_pos_3 int32
																			_ = tmp_pos_3
																			tmp_pos_3 = this.Hx_Field_pos
																			var _hx_tmp_226 int32
																			_ = _hx_tmp_226
																			var _hx_tmp_227 bool
																			_ = _hx_tmp_227
																			if tmp_pos_3 < 0 {
																				_hx_tmp_227 = true
																			} else {
																				var _hx_tmp_228 int32
																				_ = _hx_tmp_228
																				_hx_tmp_228 = tmp_pos_3
																				var _hx_tmp_229 int32
																				_ = _hx_tmp_229
																				_hx_tmp_229 = _hx_tmp_228
																				_hx_tmp_227 = (_hx_tmp_229 >= HxMakeArrayView[any](tmp_v_3).Len())
																			}

																			if _hx_tmp_227 {
																				_hx_tmp_226 = -1
																			} else {
																				_hx_tmp_226 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
																			}

																			var tmp_c_2 int32
																			_ = tmp_c_2
																			tmp_c_2 = _hx_tmp_226
																			if (((tmp_c_2 >= 97) && (tmp_c_2 <= 122)) || ((tmp_c_2 >= 65) && (tmp_c_2 <= 90))) || (tmp_c_2 == 95) {
																				var interpPos *Hx_Obj_loreline_position
																				_ = interpPos
																				interpPos = Hx_Obj_loreline_position_CreateInstance(interpLine, (interpColumn + 1), this.Hx_Field_pos, nil)
																				var tokens Hx_Typedef_loreline_tokens
																				_ = tokens
																				tokens = this.VTable.Hx_Field_readFieldAccessInterpolation(interpPos)
																				var interpLength int32
																				_ = interpLength
																				interpLength = (this.Hx_Field_pos - tokenStartPos)
																				var _hx_tmp_230 HxArray[Hx_Enum_loreline_lstringattachment]
																				_ = _hx_tmp_230
																				_hx_tmp_230 = attachments
																				_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_230, Hx_Enum_loreline_lstringattachment_Interpolation{false, (tagStart != -1), tokens, interpStart, interpLength}))
																				var tmp_v_4 HxArray[int32]
																				_ = tmp_v_4
																				tmp_v_4 = this.Hx_Field_input
																				var _hx_tmp_231 HxArray[int32]
																				_ = _hx_tmp_231
																				_hx_tmp_231 = tmp_v_4
																				var _hx_tmp_232 int32
																				_ = _hx_tmp_232
																				_hx_tmp_232 = tokenStartPos
																				var _hx_tmp_233 int32
																				_ = _hx_tmp_233
																				if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
																					_hx_tmp_233 = (tokenStartPos + interpLength)
																				} else {
																					_hx_tmp_233 = HxMakeArrayView[any](tmp_v_4).Len()
																				}

																				func() func(x string) {
																					var _hx_closure_src_140 func(x any)
																					_ = _hx_closure_src_140
																					_hx_closure_src_140 = buf.VTable.Hx_Field_add
																					return func(x string) {
																						_hx_closure_src_140(((any)(x)))
																					}
																				}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_231, _hx_tmp_232, _hx_tmp_233))
																			} else {
																				var tmp_v_4 HxArray[int32]
																				_ = tmp_v_4
																				tmp_v_4 = this.Hx_Field_input
																				var tmp_pos_4 int32
																				_ = tmp_pos_4
																				tmp_pos_4 = this.Hx_Field_pos
																				var _hx_tmp_234 int32
																				_ = _hx_tmp_234
																				var _hx_tmp_235 bool
																				_ = _hx_tmp_235
																				if tmp_pos_4 < 0 {
																					_hx_tmp_235 = true
																				} else {
																					var _hx_tmp_236 int32
																					_ = _hx_tmp_236
																					_hx_tmp_236 = tmp_pos_4
																					var _hx_tmp_237 int32
																					_ = _hx_tmp_237
																					_hx_tmp_237 = _hx_tmp_236
																					_hx_tmp_235 = (_hx_tmp_237 >= HxMakeArrayView[any](tmp_v_4).Len())
																				}

																				if _hx_tmp_235 {
																					_hx_tmp_234 = -1
																				} else {
																					_hx_tmp_234 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos_4))))
																				}

																				if (_hx_tmp_234) == 36 {
																					buf.VTable.Hx_Field_addChar(36)
																					buf.VTable.Hx_Field_addChar(36)
																					{
																						var _g int32
																						_ = _g
																						_g = 0
																						var _g1 int32
																						_ = _g1
																						_g1 = 1
																						for _g < _g1 {
																							var _hx_tmp_238 int32
																							_ = _hx_tmp_238
																							_hx_tmp_238 = _g
																							_g = (_g + 1)
																							var _hx_reserved__ int32
																							_ = _hx_reserved__
																							_hx_reserved__ = _hx_tmp_238
																							if this.Hx_Field_pos >= this.Hx_Field_length {
																								break
																							}

																							var tmp_v_5 HxArray[int32]
																							_ = tmp_v_5
																							tmp_v_5 = this.Hx_Field_input
																							var tmp_pos_5 int32
																							_ = tmp_pos_5
																							tmp_pos_5 = this.Hx_Field_pos
																							var _hx_tmp_239 int32
																							_ = _hx_tmp_239
																							var _hx_tmp_240 bool
																							_ = _hx_tmp_240
																							if tmp_pos_5 < 0 {
																								_hx_tmp_240 = true
																							} else {
																								var _hx_tmp_241 int32
																								_ = _hx_tmp_241
																								_hx_tmp_241 = tmp_pos_5
																								var _hx_tmp_242 int32
																								_ = _hx_tmp_242
																								_hx_tmp_242 = _hx_tmp_241
																								_hx_tmp_240 = (_hx_tmp_242 >= HxMakeArrayView[any](tmp_v_5).Len())
																							}

																							if _hx_tmp_240 {
																								_hx_tmp_239 = -1
																							} else {
																								_hx_tmp_239 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_5).Get(tmp_pos_5))))
																							}

																							if (_hx_tmp_239) == 10 {
																								this.Hx_Field_line++
																								this.Hx_Field_column = 1
																							} else {
																								this.Hx_Field_column++
																							}

																							this.Hx_Field_pos++
																						}
																					}

																					currentColumn++
																				} else {
																					this.VTable.Hx_Field_error("Expected identifier or { after $", false)
																					var interpLength int32
																					_ = interpLength
																					interpLength = (this.Hx_Field_pos - tokenStartPos)
																					var _hx_tmp_243 HxArray[Hx_Enum_loreline_lstringattachment]
																					_ = _hx_tmp_243
																					_hx_tmp_243 = attachments
																					var _hx_tmp_244 bool
																					_ = _hx_tmp_244
																					_hx_tmp_244 = (tagStart != -1)
																					_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_243, Hx_Enum_loreline_lstringattachment_Interpolation{false, _hx_tmp_244, ((Hx_Typedef_loreline_tokens)(HxMakeArray[*Hx_Obj_loreline_token]())), interpStart, interpLength}))
																					var tmp_v_5 HxArray[int32]
																					_ = tmp_v_5
																					tmp_v_5 = this.Hx_Field_input
																					var _hx_tmp_245 HxArray[int32]
																					_ = _hx_tmp_245
																					_hx_tmp_245 = tmp_v_5
																					var _hx_tmp_246 int32
																					_ = _hx_tmp_246
																					_hx_tmp_246 = tokenStartPos
																					var _hx_tmp_247 int32
																					_ = _hx_tmp_247
																					if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
																						_hx_tmp_247 = (tokenStartPos + interpLength)
																					} else {
																						_hx_tmp_247 = HxMakeArrayView[any](tmp_v_5).Len()
																					}

																					func() func(x string) {
																						var _hx_closure_src_141 func(x any)
																						_ = _hx_closure_src_141
																						_hx_closure_src_141 = buf.VTable.Hx_Field_add
																						return func(x string) {
																							_hx_closure_src_141(((any)(x)))
																						}
																					}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_245, _hx_tmp_246, _hx_tmp_247))
																				}
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

															currentColumn = (interpColumn + (this.Hx_Field_pos - tokenStartPos))
														} else {
															if tmp_c_1 == 10 {
																currentLine++
																currentColumn = 1
															} else {
																currentColumn++
															}

															buf.VTable.Hx_Field_addChar(tmp_c_1)
															{
																var _g int32
																_ = _g
																_g = 0
																var _g1 int32
																_ = _g1
																_g1 = 1
																for _g < _g1 {
																	var _hx_tmp_248 int32
																	_ = _hx_tmp_248
																	_hx_tmp_248 = _g
																	_g = (_g + 1)
																	var _hx_reserved__ int32
																	_ = _hx_reserved__
																	_hx_reserved__ = _hx_tmp_248
																	if this.Hx_Field_pos >= this.Hx_Field_length {
																		break
																	}

																	var tmp_v_2 HxArray[int32]
																	_ = tmp_v_2
																	tmp_v_2 = this.Hx_Field_input
																	var tmp_pos_2 int32
																	_ = tmp_pos_2
																	tmp_pos_2 = this.Hx_Field_pos
																	var _hx_tmp_249 int32
																	_ = _hx_tmp_249
																	var _hx_tmp_250 bool
																	_ = _hx_tmp_250
																	if tmp_pos_2 < 0 {
																		_hx_tmp_250 = true
																	} else {
																		var _hx_tmp_251 int32
																		_ = _hx_tmp_251
																		_hx_tmp_251 = tmp_pos_2
																		var _hx_tmp_252 int32
																		_ = _hx_tmp_252
																		_hx_tmp_252 = _hx_tmp_251
																		_hx_tmp_250 = (_hx_tmp_252 >= HxMakeArrayView[any](tmp_v_2).Len())
																	}

																	if _hx_tmp_250 {
																		_hx_tmp_249 = -1
																	} else {
																		_hx_tmp_249 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
																	}

																	if (_hx_tmp_249) == 10 {
																		this.Hx_Field_line++
																		this.Hx_Field_column = 1
																	} else {
																		this.Hx_Field_column++
																	}

																	this.Hx_Field_pos++
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
			}
		}
	}

	if valid {
		var content string
		_ = content
		content = buf.VTable.Hx_Field_toString()
		var rawContentLength int32
		_ = rawContentLength
		rawContentLength = int32(utf8.RuneCountInString(content))
		var l int32
		_ = l
		l = int32(utf8.RuneCountInString(content))
		var r int32
		_ = r
		r = 0
		for {
			var _hx_tmp_253 bool
			_ = _hx_tmp_253
			if r < l {
				_hx_tmp_253 = Hx_Field_stringtools_isSpace(content, ((l - r) - 1))
			} else {
				_hx_tmp_253 = false
			}

			if !(_hx_tmp_253) {
				break
			}

			r++
		}

		if r > 0 {
			var _hx_tmp_254 string
			_ = _hx_tmp_254
			_hx_tmp_254 = content
			content = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_254, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
		}

		var contentLength int32
		_ = contentLength
		contentLength = int32(utf8.RuneCountInString(content))
		var rtrimmedOffset int32
		_ = rtrimmedOffset
		rtrimmedOffset = (rawContentLength - contentLength)
		var _hx_tmp_255 bool
		_ = _hx_tmp_255
		var _hx_tmp_256 bool
		_ = _hx_tmp_256
		if contentLength > 0 {
			_hx_tmp_256 = this.VTable.Hx_Field_hasNonSpecialChar(content)
		} else {
			_hx_tmp_256 = false
		}

		if _hx_tmp_256 {
			var _hx_tmp_257 bool
			_ = _hx_tmp_257
			if isValue {
				var _hx_tmp_258 bool
				_ = _hx_tmp_258
				var _hx_tmp_259 bool
				_ = _hx_tmp_259
				var _hx_tmp_260 bool
				_ = _hx_tmp_260
				if !this.VTable.Hx_Field_isNumber(content) {
					_hx_tmp_260 = (content != "null")
				} else {
					_hx_tmp_260 = false
				}

				if _hx_tmp_260 {
					_hx_tmp_259 = (content != "true")
				} else {
					_hx_tmp_259 = false
				}

				if _hx_tmp_259 {
					_hx_tmp_258 = (content != "false")
				} else {
					_hx_tmp_258 = false
				}

				_hx_tmp_257 = _hx_tmp_258
			} else {
				_hx_tmp_257 = true
			}

			_hx_tmp_255 = (_hx_tmp_257)
		} else {
			_hx_tmp_255 = false
		}

		if _hx_tmp_255 {
			var _hx_tmp_261 bool
			_ = _hx_tmp_261
			if (multilineIndent == -1) && !isAfterLabel {
				_hx_tmp_261 = !this.VTable.Hx_Field_inChoiceRoot()
			} else {
				_hx_tmp_261 = false
			}

			if _hx_tmp_261 {
				var savedLine int32
				_ = savedLine
				savedLine = this.Hx_Field_line
				var savedColumn int32
				_ = savedColumn
				savedColumn = this.Hx_Field_column
				var savedPos int32
				_ = savedPos
				savedPos = this.Hx_Field_pos
				{
					var newPos int32
					_ = newPos
					newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, true)
					for this.Hx_Field_pos < newPos {
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_262 int32
							_ = _hx_tmp_262
							_hx_tmp_262 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_262
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
							var _hx_tmp_263 int32
							_ = _hx_tmp_263
							var _hx_tmp_264 bool
							_ = _hx_tmp_264
							if tmp_pos_1 < 0 {
								_hx_tmp_264 = true
							} else {
								var _hx_tmp_265 int32
								_ = _hx_tmp_265
								_hx_tmp_265 = tmp_pos_1
								var _hx_tmp_266 int32
								_ = _hx_tmp_266
								_hx_tmp_266 = _hx_tmp_265
								_hx_tmp_264 = (_hx_tmp_266 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_264 {
								_hx_tmp_263 = -1
							} else {
								_hx_tmp_263 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							if (_hx_tmp_263) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}
				}

				var afterWhitespaceAndCommentsPos int32
				_ = afterWhitespaceAndCommentsPos
				afterWhitespaceAndCommentsPos = this.Hx_Field_pos
				if this.Hx_Field_column == startColumn {
					var followingText *Hx_Obj_loreline_token
					_ = followingText
					followingText = this.VTable.Hx_Field_tryReadUnquotedString()
					if followingText != nil {
						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var _hx_tmp_267 HxArray[int32]
						_ = _hx_tmp_267
						_hx_tmp_267 = tmp_v_1
						var _hx_tmp_268 int32
						_ = _hx_tmp_268
						_hx_tmp_268 = (savedPos - rtrimmedOffset)
						var _hx_tmp_269 int32
						_ = _hx_tmp_269
						if (HxNullable[int32]{Value: afterWhitespaceAndCommentsPos, Valid: true}).Valid != false {
							_hx_tmp_269 = afterWhitespaceAndCommentsPos
						} else {
							_hx_tmp_269 = HxMakeArrayView[any](tmp_v_1).Len()
						}

						var between string
						_ = between
						between = Hx_Field_loreline_utf8vector_slice(_hx_tmp_267, _hx_tmp_268, _hx_tmp_269)
						{
							var _g Hx_Enum_loreline_tokentype
							_ = _g
							_g = followingText.Hx_Field_type
							if _g.Hx_Field_enumIndex() == 16 {
								var s_ string
								_ = s_
								s_ = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
								var attachments_ HxArray[Hx_Enum_loreline_lstringattachment]
								_ = attachments_
								attachments_ = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_attachments
								{
									if attachments_ != nil {
										var tmp__g_1 int32
										_ = tmp__g_1
										tmp__g_1 = 0
										for {
											var _hx_tmp_270 int32
											_ = _hx_tmp_270
											_hx_tmp_270 = tmp__g_1
											if !(_hx_tmp_270 < attachments_.Len()) {
												break
											}

											var attachment Hx_Enum_loreline_lstringattachment
											_ = attachment
											attachment = attachments_.Get(tmp__g_1)
											tmp__g_1++
											switch attachment.Hx_Field_enumIndex() {
											case 0:
												{
													var braces bool
													_ = braces
													braces = attachment.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_braces
													var inTag bool
													_ = inTag
													inTag = attachment.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_inTag
													var expr Hx_Typedef_loreline_tokens
													_ = expr
													expr = attachment.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_expr
													var tmp_start_1 int32
													_ = tmp_start_1
													tmp_start_1 = attachment.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
													var length int32
													_ = length
													length = attachment.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_length
													var _hx_tmp_271 HxArray[Hx_Enum_loreline_lstringattachment]
													_ = _hx_tmp_271
													_hx_tmp_271 = attachments
													_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_271, Hx_Enum_loreline_lstringattachment_Interpolation{braces, inTag, expr, ((tmp_start_1 + afterWhitespaceAndCommentsPos) - startPos), length}))
												}

											case 1:
												{
													var closing bool
													_ = closing
													closing = attachment.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_closing
													var tmp_start_1 int32
													_ = tmp_start_1
													tmp_start_1 = attachment.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_start
													var length int32
													_ = length
													length = attachment.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_length
													var _hx_tmp_272 HxArray[Hx_Enum_loreline_lstringattachment]
													_ = _hx_tmp_272
													_hx_tmp_272 = attachments
													_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_272, Hx_Enum_loreline_lstringattachment_Tag{closing, ((tmp_start_1 + afterWhitespaceAndCommentsPos) - startPos), length}))
												}

											}
										}
									}

									content = Hx_Field_go_haxe_hxstring_concat(content, between)
									content = Hx_Field_go_haxe_hxstring_concat(content, s_)
								}
							}
						}

						rtrimmedOffset = 0
						var n int32
						_ = n
						n = (this.Hx_Field_pos - 1)
						for {
							var tmp bool
							_ = tmp
							if n >= 0 {
								var tmp_v_2 HxArray[int32]
								_ = tmp_v_2
								tmp_v_2 = this.Hx_Field_input
								var _hx_tmp_273 int32
								_ = _hx_tmp_273
								var _hx_tmp_274 bool
								_ = _hx_tmp_274
								if n < 0 {
									_hx_tmp_274 = true
								} else {
									var _hx_tmp_275 int32
									_ = _hx_tmp_275
									_hx_tmp_275 = n
									var _hx_tmp_276 int32
									_ = _hx_tmp_276
									_hx_tmp_276 = _hx_tmp_275
									_hx_tmp_274 = (_hx_tmp_276 >= HxMakeArrayView[any](tmp_v_2).Len())
								}

								if _hx_tmp_274 {
									_hx_tmp_273 = -1
								} else {
									_hx_tmp_273 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(n))))
								}

								if (_hx_tmp_273) != 32 {
									var tmp_v_3 HxArray[int32]
									_ = tmp_v_3
									tmp_v_3 = this.Hx_Field_input
									var _hx_tmp_277 int32
									_ = _hx_tmp_277
									var _hx_tmp_278 bool
									_ = _hx_tmp_278
									if n < 0 {
										_hx_tmp_278 = true
									} else {
										var _hx_tmp_279 int32
										_ = _hx_tmp_279
										_hx_tmp_279 = n
										var _hx_tmp_280 int32
										_ = _hx_tmp_280
										_hx_tmp_280 = _hx_tmp_279
										_hx_tmp_278 = (_hx_tmp_280 >= HxMakeArrayView[any](tmp_v_3).Len())
									}

									if _hx_tmp_278 {
										_hx_tmp_277 = -1
									} else {
										_hx_tmp_277 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(n))))
									}

									tmp = ((_hx_tmp_277) == 9)
								} else {
									tmp = true
								}
							} else {
								tmp = false
							}

							if !tmp {
								break
							}

							rtrimmedOffset++
							n--
						}
					} else {
						this.Hx_Field_line = savedLine
						this.Hx_Field_column = savedColumn
						this.Hx_Field_pos = savedPos
					}
				} else {
					this.Hx_Field_line = savedLine
					this.Hx_Field_column = savedColumn
					this.Hx_Field_pos = savedPos
				}
			}

			var _hx_tmp_281 HxArray[Hx_Enum_loreline_lstringattachment]
			_ = _hx_tmp_281
			_hx_tmp_281 = attachments
			(Hx_Array_Sort[Hx_Enum_loreline_lstringattachment](_hx_tmp_281, func(a Hx_Enum_loreline_lstringattachment, b Hx_Enum_loreline_lstringattachment) int32 {
				return Hx_Field_loreline_lexer_compareAttachments(a, b)
			}))
			var _hx_tmp_282 HxArray[Hx_Enum_loreline_lstringattachment]
			_ = _hx_tmp_282
			if attachments.Len() > 0 {
				_hx_tmp_282 = attachments
			} else {
				_hx_tmp_282 = nil
			}

			var _hx_reserved_type HxArray[Hx_Enum_loreline_lstringattachment]
			_ = _hx_reserved_type
			_hx_reserved_type = _hx_tmp_282
			var position *Hx_Obj_loreline_position
			_ = position
			position = start
			if position == nil {
				position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
			}

			position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
			var token *Hx_Obj_loreline_token
			_ = token
			token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LString{((int32)(0)), content, _hx_reserved_type}, position)
			this.Hx_Field_previous = token
			var token1 *Hx_Obj_loreline_token
			_ = token1
			token1 = token
			token1.Hx_Field_pos.Hx_Field_length = (token1.Hx_Field_pos.Hx_Field_length - rtrimmedOffset)
			return token1
		}
	}

	this.Hx_Field_line = startLine
	this.Hx_Field_column = startColumn
	this.Hx_Field_pos = startPos
	return nil
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readString(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 1
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var attachments HxArray[Hx_Enum_loreline_lstringattachment]
	_ = attachments
	attachments = HxMakeArray[Hx_Enum_loreline_lstringattachment]()
	var escaped bool
	_ = escaped
	escaped = false
	var tagStart int32
	_ = tagStart
	tagStart = -1
	var tagIsClosing bool
	_ = tagIsClosing
	tagIsClosing = false
	var currentColumn int32
	_ = currentColumn
	currentColumn = (stringStart.Hx_Field_column + 1)
	var currentLine int32
	_ = currentLine
	currentLine = stringStart.Hx_Field_line
	var allowTags bool
	_ = allowTags
	allowTags = (this.VTable.Hx_Field_parentBlockType() == Hx_Enum_loreline_tokentype_KwBeat{})
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_5
		if escaped {
			buf.VTable.Hx_Field_addChar(92)
			buf.VTable.Hx_Field_addChar(c)
			escaped = false
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_9
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					var _hx_tmp_11 bool
					_ = _hx_tmp_11
					if tmp_pos_1 < 0 {
						_hx_tmp_11 = true
					} else {
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						_hx_tmp_12 = tmp_pos_1
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = _hx_tmp_12
						_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_11 {
						_hx_tmp_10 = -1
					} else {
						_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_10) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

			currentColumn = (currentColumn + 2)
		} else {
			if c == 92 {
				escaped = true
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_14
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = this.Hx_Field_pos
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						var _hx_tmp_16 bool
						_ = _hx_tmp_16
						if tmp_pos_1 < 0 {
							_hx_tmp_16 = true
						} else {
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							_hx_tmp_17 = tmp_pos_1
							var _hx_tmp_18 int32
							_ = _hx_tmp_18
							_hx_tmp_18 = _hx_tmp_17
							_hx_tmp_16 = (_hx_tmp_18 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_16 {
							_hx_tmp_15 = -1
						} else {
							_hx_tmp_15 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
						}

						if (_hx_tmp_15) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}
			} else {
				if (c == 34) && (tagStart == -1) {
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_19 int32
							_ = _hx_tmp_19
							_hx_tmp_19 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_19
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
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
								_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_21 {
								_hx_tmp_20 = -1
							} else {
								_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							if (_hx_tmp_20) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}

					var _hx_tmp_24 HxArray[Hx_Enum_loreline_lstringattachment]
					_ = _hx_tmp_24
					_hx_tmp_24 = attachments
					(Hx_Array_Sort[Hx_Enum_loreline_lstringattachment](_hx_tmp_24, func(a Hx_Enum_loreline_lstringattachment, b Hx_Enum_loreline_lstringattachment) int32 {
						return Hx_Field_loreline_lexer_compareAttachments(a, b)
					}))
					var _hx_reserved_type string
					_ = _hx_reserved_type
					_hx_reserved_type = buf.VTable.Hx_Field_toString()
					var _hx_tmp_25 HxArray[Hx_Enum_loreline_lstringattachment]
					_ = _hx_tmp_25
					if attachments.Len() > 0 {
						_hx_tmp_25 = attachments
					} else {
						_hx_tmp_25 = nil
					}

					var type1 HxArray[Hx_Enum_loreline_lstringattachment]
					_ = type1
					type1 = _hx_tmp_25
					var position *Hx_Obj_loreline_position
					_ = position
					position = stringStart
					if position == nil {
						position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
					}

					position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
					var token *Hx_Obj_loreline_token
					_ = token
					token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LString{((int32)(1)), _hx_reserved_type, type1}, position)
					this.Hx_Field_previous = token
					return token
				} else {
					if allowTags && (c == 60) {
						if tagStart != -1 {
							this.VTable.Hx_Field_error("Unexpected < inside tag", true)
						}

						var nextChar int32
						_ = nextChar
						if (this.Hx_Field_pos + 1) < this.Hx_Field_length {
							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = (this.Hx_Field_pos + 1)
							var _hx_tmp_26 int32
							_ = _hx_tmp_26
							var _hx_tmp_27 bool
							_ = _hx_tmp_27
							if tmp_pos_1 < 0 {
								_hx_tmp_27 = true
							} else {
								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								_hx_tmp_28 = tmp_pos_1
								var _hx_tmp_29 int32
								_ = _hx_tmp_29
								_hx_tmp_29 = _hx_tmp_28
								_hx_tmp_27 = (_hx_tmp_29 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_27 {
								_hx_tmp_26 = -1
							} else {
								_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							nextChar = _hx_tmp_26
						} else {
							nextChar = 0
						}

						tagIsClosing = (nextChar == 47)
						if this.VTable.Hx_Field_isTagStart(this.Hx_Field_pos) {
							tagStart = buf.VTable.Hx_Field_get_length()
						}

						buf.VTable.Hx_Field_addChar(c)
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_30 int32
								_ = _hx_tmp_30
								_hx_tmp_30 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_30
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_31 int32
								_ = _hx_tmp_31
								var _hx_tmp_32 bool
								_ = _hx_tmp_32
								if tmp_pos_1 < 0 {
									_hx_tmp_32 = true
								} else {
									var _hx_tmp_33 int32
									_ = _hx_tmp_33
									_hx_tmp_33 = tmp_pos_1
									var _hx_tmp_34 int32
									_ = _hx_tmp_34
									_hx_tmp_34 = _hx_tmp_33
									_hx_tmp_32 = (_hx_tmp_34 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_32 {
									_hx_tmp_31 = -1
								} else {
									_hx_tmp_31 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								if (_hx_tmp_31) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						currentColumn++
					} else {
						if c == 62 {
							buf.VTable.Hx_Field_addChar(c)
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_35 int32
									_ = _hx_tmp_35
									_hx_tmp_35 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_35
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = this.Hx_Field_pos
									var _hx_tmp_36 int32
									_ = _hx_tmp_36
									var _hx_tmp_37 bool
									_ = _hx_tmp_37
									if tmp_pos_1 < 0 {
										_hx_tmp_37 = true
									} else {
										var _hx_tmp_38 int32
										_ = _hx_tmp_38
										_hx_tmp_38 = tmp_pos_1
										var _hx_tmp_39 int32
										_ = _hx_tmp_39
										_hx_tmp_39 = _hx_tmp_38
										_hx_tmp_37 = (_hx_tmp_39 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_37 {
										_hx_tmp_36 = -1
									} else {
										_hx_tmp_36 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									if (_hx_tmp_36) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							currentColumn++
							if tagStart != -1 {
								var _hx_tmp_40 HxArray[Hx_Enum_loreline_lstringattachment]
								_ = _hx_tmp_40
								_hx_tmp_40 = attachments
								var _hx_tmp_41 bool
								_ = _hx_tmp_41
								_hx_tmp_41 = tagIsClosing
								var _hx_tmp_42 int32
								_ = _hx_tmp_42
								_hx_tmp_42 = tagStart
								_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_40, Hx_Enum_loreline_lstringattachment_Tag{_hx_tmp_41, _hx_tmp_42, (buf.VTable.Hx_Field_get_length() - tagStart)}))
								tagStart = -1
							}
						} else {
							if (c == 36) && !escaped {
								var interpStart int32
								_ = interpStart
								interpStart = buf.VTable.Hx_Field_get_length()
								var interpLine int32
								_ = interpLine
								interpLine = currentLine
								var interpColumn int32
								_ = interpColumn
								interpColumn = currentColumn
								var tokenStartPos int32
								_ = tokenStartPos
								tokenStartPos = this.Hx_Field_pos
								{
									var hx_try_return *Hx_Obj_loreline_token
									_ = hx_try_return
									var hx_try_state int
									_ = hx_try_state

									func() {
										defer func() {
											if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
												if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "*Hx_Obj_loreline_lexererror").(*Hx_Obj_loreline_lexererror); _hx_ok {
													_ = e
													{
														if e.Hx_Field_pos == nil {
															e.Hx_Field_pos = Hx_Obj_loreline_position_CreateInstance(interpLine, interpColumn, this.Hx_Field_pos, nil)
														}

														panic(Hx_Field_haxe_exception_thrown(((any)(e))))
													}
												} else {
													panic(_hx_ex)
												}
											}
										}()
										{
											{
												var _g int32
												_ = _g
												_g = 0
												var _g1 int32
												_ = _g1
												_g1 = 1
												for _g < _g1 {
													var _hx_tmp_43 int32
													_ = _hx_tmp_43
													_hx_tmp_43 = _g
													_g = (_g + 1)
													var _hx_reserved__ int32
													_ = _hx_reserved__
													_hx_reserved__ = _hx_tmp_43
													if this.Hx_Field_pos >= this.Hx_Field_length {
														break
													}

													var tmp_v_1 HxArray[int32]
													_ = tmp_v_1
													tmp_v_1 = this.Hx_Field_input
													var tmp_pos_1 int32
													_ = tmp_pos_1
													tmp_pos_1 = this.Hx_Field_pos
													var _hx_tmp_44 int32
													_ = _hx_tmp_44
													var _hx_tmp_45 bool
													_ = _hx_tmp_45
													if tmp_pos_1 < 0 {
														_hx_tmp_45 = true
													} else {
														var _hx_tmp_46 int32
														_ = _hx_tmp_46
														_hx_tmp_46 = tmp_pos_1
														var _hx_tmp_47 int32
														_ = _hx_tmp_47
														_hx_tmp_47 = _hx_tmp_46
														_hx_tmp_45 = (_hx_tmp_47 >= HxMakeArrayView[any](tmp_v_1).Len())
													}

													if _hx_tmp_45 {
														_hx_tmp_44 = -1
													} else {
														_hx_tmp_44 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
													}

													if (_hx_tmp_44) == 10 {
														this.Hx_Field_line++
														this.Hx_Field_column = 1
													} else {
														this.Hx_Field_column++
													}

													this.Hx_Field_pos++
												}
											}

											currentColumn++
											var tmp_v_1 HxArray[int32]
											_ = tmp_v_1
											tmp_v_1 = this.Hx_Field_input
											var tmp_pos_1 int32
											_ = tmp_pos_1
											tmp_pos_1 = this.Hx_Field_pos
											var _hx_tmp_48 int32
											_ = _hx_tmp_48
											var _hx_tmp_49 bool
											_ = _hx_tmp_49
											if tmp_pos_1 < 0 {
												_hx_tmp_49 = true
											} else {
												var _hx_tmp_50 int32
												_ = _hx_tmp_50
												_hx_tmp_50 = tmp_pos_1
												var _hx_tmp_51 int32
												_ = _hx_tmp_51
												_hx_tmp_51 = _hx_tmp_50
												_hx_tmp_49 = (_hx_tmp_51 >= HxMakeArrayView[any](tmp_v_1).Len())
											}

											if _hx_tmp_49 {
												_hx_tmp_48 = -1
											} else {
												_hx_tmp_48 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
											}

											if (_hx_tmp_48) == 123 {
												{
													var _g int32
													_ = _g
													_g = 0
													var _g1 int32
													_ = _g1
													_g1 = 1
													for _g < _g1 {
														var _hx_tmp_52 int32
														_ = _hx_tmp_52
														_hx_tmp_52 = _g
														_g = (_g + 1)
														var _hx_reserved__ int32
														_ = _hx_reserved__
														_hx_reserved__ = _hx_tmp_52
														if this.Hx_Field_pos >= this.Hx_Field_length {
															break
														}

														var tmp_v_2 HxArray[int32]
														_ = tmp_v_2
														tmp_v_2 = this.Hx_Field_input
														var tmp_pos_2 int32
														_ = tmp_pos_2
														tmp_pos_2 = this.Hx_Field_pos
														var _hx_tmp_53 int32
														_ = _hx_tmp_53
														var _hx_tmp_54 bool
														_ = _hx_tmp_54
														if tmp_pos_2 < 0 {
															_hx_tmp_54 = true
														} else {
															var _hx_tmp_55 int32
															_ = _hx_tmp_55
															_hx_tmp_55 = tmp_pos_2
															var _hx_tmp_56 int32
															_ = _hx_tmp_56
															_hx_tmp_56 = _hx_tmp_55
															_hx_tmp_54 = (_hx_tmp_56 >= HxMakeArrayView[any](tmp_v_2).Len())
														}

														if _hx_tmp_54 {
															_hx_tmp_53 = -1
														} else {
															_hx_tmp_53 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
														}

														if (_hx_tmp_53) == 10 {
															this.Hx_Field_line++
															this.Hx_Field_column = 1
														} else {
															this.Hx_Field_column++
														}

														this.Hx_Field_pos++
													}
												}

												currentColumn++
												var interpPos *Hx_Obj_loreline_position
												_ = interpPos
												interpPos = Hx_Obj_loreline_position_CreateInstance(interpLine, (interpColumn + 2), this.Hx_Field_pos, nil)
												var tokens Hx_Typedef_loreline_tokens
												_ = tokens
												tokens = this.VTable.Hx_Field_readComplexInterpolation(interpPos)
												var interpLength int32
												_ = interpLength
												interpLength = (this.Hx_Field_pos - tokenStartPos)
												var _hx_tmp_57 HxArray[Hx_Enum_loreline_lstringattachment]
												_ = _hx_tmp_57
												_hx_tmp_57 = attachments
												_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_57, Hx_Enum_loreline_lstringattachment_Interpolation{true, (tagStart != -1), tokens, interpStart, interpLength}))
												var tmp_v_2 HxArray[int32]
												_ = tmp_v_2
												tmp_v_2 = this.Hx_Field_input
												var _hx_tmp_58 HxArray[int32]
												_ = _hx_tmp_58
												_hx_tmp_58 = tmp_v_2
												var _hx_tmp_59 int32
												_ = _hx_tmp_59
												_hx_tmp_59 = tokenStartPos
												var _hx_tmp_60 int32
												_ = _hx_tmp_60
												if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
													_hx_tmp_60 = (tokenStartPos + interpLength)
												} else {
													_hx_tmp_60 = HxMakeArrayView[any](tmp_v_2).Len()
												}

												func() func(x string) {
													var _hx_closure_src_142 func(x any)
													_ = _hx_closure_src_142
													_hx_closure_src_142 = buf.VTable.Hx_Field_add
													return func(x string) {
														_hx_closure_src_142(((any)(x)))
													}
												}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_58, _hx_tmp_59, _hx_tmp_60))
											} else {
												var tmp_v_2 HxArray[int32]
												_ = tmp_v_2
												tmp_v_2 = this.Hx_Field_input
												var tmp_pos_2 int32
												_ = tmp_pos_2
												tmp_pos_2 = this.Hx_Field_pos
												var _hx_tmp_61 int32
												_ = _hx_tmp_61
												var _hx_tmp_62 bool
												_ = _hx_tmp_62
												if tmp_pos_2 < 0 {
													_hx_tmp_62 = true
												} else {
													var _hx_tmp_63 int32
													_ = _hx_tmp_63
													_hx_tmp_63 = tmp_pos_2
													var _hx_tmp_64 int32
													_ = _hx_tmp_64
													_hx_tmp_64 = _hx_tmp_63
													_hx_tmp_62 = (_hx_tmp_64 >= HxMakeArrayView[any](tmp_v_2).Len())
												}

												if _hx_tmp_62 {
													_hx_tmp_61 = -1
												} else {
													_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
												}

												var tmp_c_1 int32
												_ = tmp_c_1
												tmp_c_1 = _hx_tmp_61
												if (((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95) {
													var interpPos *Hx_Obj_loreline_position
													_ = interpPos
													interpPos = Hx_Obj_loreline_position_CreateInstance(interpLine, (interpColumn + 1), this.Hx_Field_pos, nil)
													var tokens Hx_Typedef_loreline_tokens
													_ = tokens
													tokens = this.VTable.Hx_Field_readFieldAccessInterpolation(interpPos)
													var interpLength int32
													_ = interpLength
													interpLength = (this.Hx_Field_pos - tokenStartPos)
													var _hx_tmp_65 HxArray[Hx_Enum_loreline_lstringattachment]
													_ = _hx_tmp_65
													_hx_tmp_65 = attachments
													_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_65, Hx_Enum_loreline_lstringattachment_Interpolation{false, (tagStart != -1), tokens, interpStart, interpLength}))
													var tmp_v_3 HxArray[int32]
													_ = tmp_v_3
													tmp_v_3 = this.Hx_Field_input
													var _hx_tmp_66 HxArray[int32]
													_ = _hx_tmp_66
													_hx_tmp_66 = tmp_v_3
													var _hx_tmp_67 int32
													_ = _hx_tmp_67
													_hx_tmp_67 = tokenStartPos
													var _hx_tmp_68 int32
													_ = _hx_tmp_68
													if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
														_hx_tmp_68 = (tokenStartPos + interpLength)
													} else {
														_hx_tmp_68 = HxMakeArrayView[any](tmp_v_3).Len()
													}

													func() func(x string) {
														var _hx_closure_src_143 func(x any)
														_ = _hx_closure_src_143
														_hx_closure_src_143 = buf.VTable.Hx_Field_add
														return func(x string) {
															_hx_closure_src_143(((any)(x)))
														}
													}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_66, _hx_tmp_67, _hx_tmp_68))
												} else {
													var tmp_v_3 HxArray[int32]
													_ = tmp_v_3
													tmp_v_3 = this.Hx_Field_input
													var tmp_pos_3 int32
													_ = tmp_pos_3
													tmp_pos_3 = this.Hx_Field_pos
													var _hx_tmp_69 int32
													_ = _hx_tmp_69
													var _hx_tmp_70 bool
													_ = _hx_tmp_70
													if tmp_pos_3 < 0 {
														_hx_tmp_70 = true
													} else {
														var _hx_tmp_71 int32
														_ = _hx_tmp_71
														_hx_tmp_71 = tmp_pos_3
														var _hx_tmp_72 int32
														_ = _hx_tmp_72
														_hx_tmp_72 = _hx_tmp_71
														_hx_tmp_70 = (_hx_tmp_72 >= HxMakeArrayView[any](tmp_v_3).Len())
													}

													if _hx_tmp_70 {
														_hx_tmp_69 = -1
													} else {
														_hx_tmp_69 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
													}

													if (_hx_tmp_69) == 36 {
														buf.VTable.Hx_Field_addChar(36)
														buf.VTable.Hx_Field_addChar(36)
														{
															var _g int32
															_ = _g
															_g = 0
															var _g1 int32
															_ = _g1
															_g1 = 1
															for _g < _g1 {
																var _hx_tmp_73 int32
																_ = _hx_tmp_73
																_hx_tmp_73 = _g
																_g = (_g + 1)
																var _hx_reserved__ int32
																_ = _hx_reserved__
																_hx_reserved__ = _hx_tmp_73
																if this.Hx_Field_pos >= this.Hx_Field_length {
																	break
																}

																var tmp_v_4 HxArray[int32]
																_ = tmp_v_4
																tmp_v_4 = this.Hx_Field_input
																var tmp_pos_4 int32
																_ = tmp_pos_4
																tmp_pos_4 = this.Hx_Field_pos
																var _hx_tmp_74 int32
																_ = _hx_tmp_74
																var _hx_tmp_75 bool
																_ = _hx_tmp_75
																if tmp_pos_4 < 0 {
																	_hx_tmp_75 = true
																} else {
																	var _hx_tmp_76 int32
																	_ = _hx_tmp_76
																	_hx_tmp_76 = tmp_pos_4
																	var _hx_tmp_77 int32
																	_ = _hx_tmp_77
																	_hx_tmp_77 = _hx_tmp_76
																	_hx_tmp_75 = (_hx_tmp_77 >= HxMakeArrayView[any](tmp_v_4).Len())
																}

																if _hx_tmp_75 {
																	_hx_tmp_74 = -1
																} else {
																	_hx_tmp_74 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_4).Get(tmp_pos_4))))
																}

																if (_hx_tmp_74) == 10 {
																	this.Hx_Field_line++
																	this.Hx_Field_column = 1
																} else {
																	this.Hx_Field_column++
																}

																this.Hx_Field_pos++
															}
														}

														currentColumn++
													} else {
														this.VTable.Hx_Field_error("Expected identifier or { after $", false)
														var interpLength int32
														_ = interpLength
														interpLength = (this.Hx_Field_pos - tokenStartPos)
														var _hx_tmp_78 HxArray[Hx_Enum_loreline_lstringattachment]
														_ = _hx_tmp_78
														_hx_tmp_78 = attachments
														var _hx_tmp_79 bool
														_ = _hx_tmp_79
														_hx_tmp_79 = (tagStart != -1)
														_ = (Hx_Array_Push[Hx_Enum_loreline_lstringattachment](_hx_tmp_78, Hx_Enum_loreline_lstringattachment_Interpolation{false, _hx_tmp_79, ((Hx_Typedef_loreline_tokens)(HxMakeArray[*Hx_Obj_loreline_token]())), interpStart, interpLength}))
														var tmp_v_4 HxArray[int32]
														_ = tmp_v_4
														tmp_v_4 = this.Hx_Field_input
														var _hx_tmp_80 HxArray[int32]
														_ = _hx_tmp_80
														_hx_tmp_80 = tmp_v_4
														var _hx_tmp_81 int32
														_ = _hx_tmp_81
														_hx_tmp_81 = tokenStartPos
														var _hx_tmp_82 int32
														_ = _hx_tmp_82
														if (HxNullable[int32]{Value: interpLength, Valid: true}).Valid != false {
															_hx_tmp_82 = (tokenStartPos + interpLength)
														} else {
															_hx_tmp_82 = HxMakeArrayView[any](tmp_v_4).Len()
														}

														func() func(x string) {
															var _hx_closure_src_144 func(x any)
															_ = _hx_closure_src_144
															_hx_closure_src_144 = buf.VTable.Hx_Field_add
															return func(x string) {
																_hx_closure_src_144(((any)(x)))
															}
														}()(Hx_Field_loreline_utf8vector_slice(_hx_tmp_80, _hx_tmp_81, _hx_tmp_82))
													}
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

								currentColumn = (interpColumn + (this.Hx_Field_pos - tokenStartPos))
							} else {
								if c == 10 {
									currentLine++
									currentColumn = 1
								} else {
									currentColumn++
								}

								buf.VTable.Hx_Field_addChar(c)
								{
									var _g int32
									_ = _g
									_g = 0
									var _g1 int32
									_ = _g1
									_g1 = 1
									for _g < _g1 {
										var _hx_tmp_83 int32
										_ = _hx_tmp_83
										_hx_tmp_83 = _g
										_g = (_g + 1)
										var _hx_reserved__ int32
										_ = _hx_reserved__
										_hx_reserved__ = _hx_tmp_83
										if this.Hx_Field_pos >= this.Hx_Field_length {
											break
										}

										var tmp_v_1 HxArray[int32]
										_ = tmp_v_1
										tmp_v_1 = this.Hx_Field_input
										var tmp_pos_1 int32
										_ = tmp_pos_1
										tmp_pos_1 = this.Hx_Field_pos
										var _hx_tmp_84 int32
										_ = _hx_tmp_84
										var _hx_tmp_85 bool
										_ = _hx_tmp_85
										if tmp_pos_1 < 0 {
											_hx_tmp_85 = true
										} else {
											var _hx_tmp_86 int32
											_ = _hx_tmp_86
											_hx_tmp_86 = tmp_pos_1
											var _hx_tmp_87 int32
											_ = _hx_tmp_87
											_hx_tmp_87 = _hx_tmp_86
											_hx_tmp_85 = (_hx_tmp_87 >= HxMakeArrayView[any](tmp_v_1).Len())
										}

										if _hx_tmp_85 {
											_hx_tmp_84 = -1
										} else {
											_hx_tmp_84 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
										}

										if (_hx_tmp_84) == 10 {
											this.Hx_Field_line++
											this.Hx_Field_column = 1
										} else {
											this.Hx_Field_column++
										}

										this.Hx_Field_pos++
									}
								}
							}
						}
					}
				}
			}
		}
	}

	this.VTable.Hx_Field_error("Unterminated string", true)
	return nil
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readComplexInterpolation(interpStart *Hx_Obj_loreline_position) Hx_Typedef_loreline_tokens {
	_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
	var tokens HxArray[*Hx_Obj_loreline_token]
	_ = tokens
	tokens = HxMakeArray[*Hx_Obj_loreline_token]()
	var braceLevel int32
	_ = braceLevel
	braceLevel = 1
	var currentColumn int32
	_ = currentColumn
	currentColumn = interpStart.Hx_Field_column
	var currentLine int32
	_ = currentLine
	currentLine = interpStart.Hx_Field_line
	for (this.Hx_Field_pos < this.Hx_Field_length) && (braceLevel > 0) {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
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

		if (_hx_tmp_0) == 34 {
			var stringPos *Hx_Obj_loreline_position
			_ = stringPos
			stringPos = Hx_Obj_loreline_position_CreateInstance(currentLine, currentColumn, this.Hx_Field_pos, nil)
			var _hx_tmp_4 HxArray[*Hx_Obj_loreline_token]
			_ = _hx_tmp_4
			_hx_tmp_4 = tokens
			_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_4, this.VTable.Hx_Field_readString(stringPos)))
			currentColumn = (currentColumn + (this.Hx_Field_pos - stringPos.Hx_Field_offset))
			continue
		}

		var token *Hx_Obj_loreline_token
		_ = token
		token = this.VTable.Hx_Field_nextToken()
		_ = (Hx_Array_Push[*Hx_Obj_loreline_token](tokens, token))
		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_LineBreak{}) {
			currentLine++
			currentColumn = 1
		} else {
			if !((token.Hx_Field_type == Hx_Enum_loreline_tokentype_Indent{}) || (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Unindent{})) {
				var tokenLength int32
				_ = tokenLength
				var _g Hx_Enum_loreline_tokentype
				_ = _g
				_g = token.Hx_Field_type
				switch _g.Hx_Field_enumIndex() {
				case 16:
					{
						var q int32
						_ = q
						q = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_quotes
						var s string
						_ = s
						s = _g.(Hx_Enum_loreline_tokentype_LString).Hx_Field_s
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = int32(utf8.RuneCountInString(s))
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						if q != ((int32)(0)) {
							_hx_tmp_6 = 2
						} else {
							_hx_tmp_6 = 0
						}

						tokenLength = (_hx_tmp_5 + (_hx_tmp_6))
					}

				case 17:
					{
						var n float64
						_ = n
						n = _g.(Hx_Enum_loreline_tokentype_LNumber).Hx_Field_n
						tokenLength = int32(utf8.RuneCountInString(Hx_Field_std_string(n)))
					}

				case 20:
					{
						var name string
						_ = name
						name = _g.(Hx_Enum_loreline_tokentype_Identifier).Hx_Field_name
						tokenLength = int32(utf8.RuneCountInString(name))
					}

				default:
					{
						tokenLength = 1
					}
				}

				currentColumn = (currentColumn + tokenLength)
			}
		}

		switch token.Hx_Field_type.Hx_Field_enumIndex() {
		case 46:
			{
				braceLevel++
			}

		case 47:
			{
				braceLevel--
			}

		default:
			{
			}
		}

		if (token.Hx_Field_type == Hx_Enum_loreline_tokentype_Eof{}) {
			this.VTable.Hx_Field_error("Unterminated interpolation expression", true)
		}
	}

	_ = (Hx_Array_Pop[*Hx_Obj_loreline_token](tokens).Value)
	_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
	return ((Hx_Typedef_loreline_tokens)(tokens))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_nextTokenWithPosition(startPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	var savedLine int32
	_ = savedLine
	savedLine = this.Hx_Field_line
	var savedColumn int32
	_ = savedColumn
	savedColumn = this.Hx_Field_column
	var savedPos int32
	_ = savedPos
	savedPos = this.Hx_Field_pos
	var token *Hx_Obj_loreline_token
	_ = token
	token = this.VTable.Hx_Field_nextToken()
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = token.Hx_Field_pos.Hx_Field_length
	var tokenOffset int32
	_ = tokenOffset
	tokenOffset = (this.Hx_Field_pos - savedPos)
	token.Hx_Field_pos = Hx_Obj_loreline_position_CreateInstance(startPos.Hx_Field_line, (startPos.Hx_Field_column + tokenOffset), (startPos.Hx_Field_offset + tokenOffset), _hx_reserved_len)
	this.Hx_Field_line = savedLine
	this.Hx_Field_column = savedColumn
	return token
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readFieldAccessInterpolation(interpStart *Hx_Obj_loreline_position) Hx_Typedef_loreline_tokens {
	_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
	var tokens HxArray[*Hx_Obj_loreline_token]
	_ = tokens
	tokens = HxMakeArray[*Hx_Obj_loreline_token]()
	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
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
		this.VTable.Hx_Field_error("Expected identifier in field access", true)
	}

	var idStartPos int32
	_ = idStartPos
	idStartPos = this.Hx_Field_pos
	for this.Hx_Field_pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
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

		var tmp_c_1 int32
		_ = tmp_c_1
		tmp_c_1 = _hx_tmp_4
		if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
			break
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_8
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_2 HxArray[int32]
				_ = tmp_v_2
				tmp_v_2 = this.Hx_Field_input
				var tmp_pos_2 int32
				_ = tmp_pos_2
				tmp_pos_2 = this.Hx_Field_pos
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				if tmp_pos_2 < 0 {
					_hx_tmp_10 = true
				} else {
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					_hx_tmp_11 = tmp_pos_2
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					_hx_tmp_12 = _hx_tmp_11
					_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](tmp_v_2).Len())
				}

				if _hx_tmp_10 {
					_hx_tmp_9 = -1
				} else {
					_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
				}

				if (_hx_tmp_9) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - idStartPos), Valid: true}

	var _hx_tmp_13 HxArray[int32]
	_ = _hx_tmp_13
	_hx_tmp_13 = v1
	var _hx_tmp_14 int32
	_ = _hx_tmp_14
	_hx_tmp_14 = idStartPos
	var _hx_tmp_15 int32
	_ = _hx_tmp_15
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_16 int32
		_ = _hx_tmp_16
		_hx_tmp_16 = idStartPos
		var _hx_tmp_17 int32
		_ = _hx_tmp_17
		_hx_tmp_17 = _hx_tmp_16
		_hx_tmp_15 = (_hx_tmp_17 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_15 = HxMakeArrayView[any](v1).Len()
	}

	var name string
	_ = name
	name = Hx_Field_loreline_utf8vector_slice(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15)
	var _hx_tmp_18 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_18
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_145 func(_hx_param_key any) bool
		_ = _hx_closure_src_145
		_hx_closure_src_145 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_145(((any)(_hx_param_key)))
		}
	}()(name) {
		_hx_tmp_18 = func() func(_hx_param_key string) Hx_Enum_loreline_tokentype {
			var _hx_closure_src_146 func(_hx_param_key any) any
			_ = _hx_closure_src_146
			_hx_closure_src_146 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) Hx_Enum_loreline_tokentype {
				return func() Hx_Enum_loreline_tokentype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_146(((any)(_hx_param_key))))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_tokentype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_tokentype)
				}()
			}
		}()(name)
	} else {
		_hx_tmp_18 = Hx_Enum_loreline_tokentype_Identifier{name}
	}

	var tokenType Hx_Enum_loreline_tokentype
	_ = tokenType
	tokenType = _hx_tmp_18
	var _hx_tmp_19 HxArray[*Hx_Obj_loreline_token]
	_ = _hx_tmp_19
	_hx_tmp_19 = tokens
	_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_19, Hx_Obj_loreline_token_CreateInstance(tokenType, Hx_Obj_loreline_position_CreateInstance(interpStart.Hx_Field_line, interpStart.Hx_Field_column, idStartPos, (this.Hx_Field_pos-idStartPos)))))
hx_label_0:
	for this.Hx_Field_pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
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
			_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_21 {
			_hx_tmp_20 = -1
		} else {
			_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
		}

		switch _hx_tmp_20 {
		case 40:
			{
				var _hx_tmp_24 HxArray[*Hx_Obj_loreline_token]
				_ = _hx_tmp_24
				_hx_tmp_24 = tokens
				_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_24, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LParen{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_25 int32
						_ = _hx_tmp_25
						_hx_tmp_25 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_25
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_26 int32
						_ = _hx_tmp_26
						var _hx_tmp_27 bool
						_ = _hx_tmp_27
						if tmp_pos_2 < 0 {
							_hx_tmp_27 = true
						} else {
							var _hx_tmp_28 int32
							_ = _hx_tmp_28
							_hx_tmp_28 = tmp_pos_2
							var _hx_tmp_29 int32
							_ = _hx_tmp_29
							_hx_tmp_29 = _hx_tmp_28
							_hx_tmp_27 = (_hx_tmp_29 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_27 {
							_hx_tmp_26 = -1
						} else {
							_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_26) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(12))))
				_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
				var parenLevel int32
				_ = parenLevel
				parenLevel = 1
				for (this.Hx_Field_pos < this.Hx_Field_length) && (parenLevel > 0) {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_30 int32
					_ = _hx_tmp_30
					var _hx_tmp_31 bool
					_ = _hx_tmp_31
					if tmp_pos_2 < 0 {
						_hx_tmp_31 = true
					} else {
						var _hx_tmp_32 int32
						_ = _hx_tmp_32
						_hx_tmp_32 = tmp_pos_2
						var _hx_tmp_33 int32
						_ = _hx_tmp_33
						_hx_tmp_33 = _hx_tmp_32
						_hx_tmp_31 = (_hx_tmp_33 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_31 {
						_hx_tmp_30 = -1
					} else {
						_hx_tmp_30 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_30) == 41 {
						parenLevel--
						if parenLevel == 0 {
							var _hx_tmp_34 HxArray[*Hx_Obj_loreline_token]
							_ = _hx_tmp_34
							_hx_tmp_34 = tokens
							_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_34, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_RParen{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_35 int32
									_ = _hx_tmp_35
									_hx_tmp_35 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_35
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_3 HxArray[int32]
									_ = tmp_v_3
									tmp_v_3 = this.Hx_Field_input
									var tmp_pos_3 int32
									_ = tmp_pos_3
									tmp_pos_3 = this.Hx_Field_pos
									var _hx_tmp_36 int32
									_ = _hx_tmp_36
									var _hx_tmp_37 bool
									_ = _hx_tmp_37
									if tmp_pos_3 < 0 {
										_hx_tmp_37 = true
									} else {
										var _hx_tmp_38 int32
										_ = _hx_tmp_38
										_hx_tmp_38 = tmp_pos_3
										var _hx_tmp_39 int32
										_ = _hx_tmp_39
										_hx_tmp_39 = _hx_tmp_38
										_hx_tmp_37 = (_hx_tmp_39 >= HxMakeArrayView[any](tmp_v_3).Len())
									}

									if _hx_tmp_37 {
										_hx_tmp_36 = -1
									} else {
										_hx_tmp_36 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
									}

									if (_hx_tmp_36) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							_ = (Hx_Array_Pop[int32](this.Hx_Field_stack))
							_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
							break
						}
					} else {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_3 int32
						_ = tmp_pos_3
						tmp_pos_3 = this.Hx_Field_pos
						var _hx_tmp_40 int32
						_ = _hx_tmp_40
						var _hx_tmp_41 bool
						_ = _hx_tmp_41
						if tmp_pos_3 < 0 {
							_hx_tmp_41 = true
						} else {
							var _hx_tmp_42 int32
							_ = _hx_tmp_42
							_hx_tmp_42 = tmp_pos_3
							var _hx_tmp_43 int32
							_ = _hx_tmp_43
							_hx_tmp_43 = _hx_tmp_42
							_hx_tmp_41 = (_hx_tmp_43 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_41 {
							_hx_tmp_40 = -1
						} else {
							_hx_tmp_40 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
						}

						if (_hx_tmp_40) == 40 {
							parenLevel++
						}
					}

					var tmp_v1_1 HxArray[int32]
					_ = tmp_v1_1
					tmp_v1_1 = this.Hx_Field_input
					var pos1 int32
					_ = pos1
					pos1 = this.Hx_Field_pos
					var _hx_tmp_44 int32
					_ = _hx_tmp_44
					var _hx_tmp_45 bool
					_ = _hx_tmp_45
					if pos1 < 0 {
						_hx_tmp_45 = true
					} else {
						var _hx_tmp_46 int32
						_ = _hx_tmp_46
						_hx_tmp_46 = pos1
						var _hx_tmp_47 int32
						_ = _hx_tmp_47
						_hx_tmp_47 = _hx_tmp_46
						_hx_tmp_45 = (_hx_tmp_47 >= HxMakeArrayView[any](tmp_v1_1).Len())
					}

					if _hx_tmp_45 {
						_hx_tmp_44 = -1
					} else {
						_hx_tmp_44 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v1_1).Get(pos1))))
					}

					if (_hx_tmp_44) == 44 {
						var _hx_tmp_48 HxArray[*Hx_Obj_loreline_token]
						_ = _hx_tmp_48
						_hx_tmp_48 = tokens
						_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_48, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Comma{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_49 int32
								_ = _hx_tmp_49
								_hx_tmp_49 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_49
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_3 HxArray[int32]
								_ = tmp_v_3
								tmp_v_3 = this.Hx_Field_input
								var tmp_pos_3 int32
								_ = tmp_pos_3
								tmp_pos_3 = this.Hx_Field_pos
								var _hx_tmp_50 int32
								_ = _hx_tmp_50
								var _hx_tmp_51 bool
								_ = _hx_tmp_51
								if tmp_pos_3 < 0 {
									_hx_tmp_51 = true
								} else {
									var _hx_tmp_52 int32
									_ = _hx_tmp_52
									_hx_tmp_52 = tmp_pos_3
									var _hx_tmp_53 int32
									_ = _hx_tmp_53
									_hx_tmp_53 = _hx_tmp_52
									_hx_tmp_51 = (_hx_tmp_53 >= HxMakeArrayView[any](tmp_v_3).Len())
								}

								if _hx_tmp_51 {
									_hx_tmp_50 = -1
								} else {
									_hx_tmp_50 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
								}

								if (_hx_tmp_50) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}
					} else {
						var _hx_tmp_54 HxArray[*Hx_Obj_loreline_token]
						_ = _hx_tmp_54
						_hx_tmp_54 = tokens
						_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_54, this.VTable.Hx_Field_nextToken()))
					}
				}

				if parenLevel > 0 {
					this.VTable.Hx_Field_error("Unterminated function call in interpolation", true)
				}
			}

		case 46:
			{
				var tmp bool
				_ = tmp
				if (this.Hx_Field_pos + 1) < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = (this.Hx_Field_pos + 1)
					var _hx_tmp_55 int32
					_ = _hx_tmp_55
					var _hx_tmp_56 bool
					_ = _hx_tmp_56
					if tmp_pos_2 < 0 {
						_hx_tmp_56 = true
					} else {
						var _hx_tmp_57 int32
						_ = _hx_tmp_57
						_hx_tmp_57 = tmp_pos_2
						var _hx_tmp_58 int32
						_ = _hx_tmp_58
						_hx_tmp_58 = _hx_tmp_57
						_hx_tmp_56 = (_hx_tmp_58 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_56 {
						_hx_tmp_55 = -1
					} else {
						_hx_tmp_55 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_55
					tmp = ((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95))
				} else {
					tmp = false
				}

				if tmp {
					var _hx_tmp_59 HxArray[*Hx_Obj_loreline_token]
					_ = _hx_tmp_59
					_hx_tmp_59 = tokens
					_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_59, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Dot{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_60 int32
							_ = _hx_tmp_60
							_hx_tmp_60 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_60
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var tmp_pos_2 int32
							_ = tmp_pos_2
							tmp_pos_2 = this.Hx_Field_pos
							var _hx_tmp_61 int32
							_ = _hx_tmp_61
							var _hx_tmp_62 bool
							_ = _hx_tmp_62
							if tmp_pos_2 < 0 {
								_hx_tmp_62 = true
							} else {
								var _hx_tmp_63 int32
								_ = _hx_tmp_63
								_hx_tmp_63 = tmp_pos_2
								var _hx_tmp_64 int32
								_ = _hx_tmp_64
								_hx_tmp_64 = _hx_tmp_63
								_hx_tmp_62 = (_hx_tmp_64 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_62 {
								_hx_tmp_61 = -1
							} else {
								_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
							}

							if (_hx_tmp_61) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}

					var tmp_idStartPos_1 int32
					_ = tmp_idStartPos_1
					tmp_idStartPos_1 = this.Hx_Field_pos
					for this.Hx_Field_pos < this.Hx_Field_length {
						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_65 int32
						_ = _hx_tmp_65
						var _hx_tmp_66 bool
						_ = _hx_tmp_66
						if tmp_pos_2 < 0 {
							_hx_tmp_66 = true
						} else {
							var _hx_tmp_67 int32
							_ = _hx_tmp_67
							_hx_tmp_67 = tmp_pos_2
							var _hx_tmp_68 int32
							_ = _hx_tmp_68
							_hx_tmp_68 = _hx_tmp_67
							_hx_tmp_66 = (_hx_tmp_68 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_66 {
							_hx_tmp_65 = -1
						} else {
							_hx_tmp_65 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						var tmp_c_1 int32
						_ = tmp_c_1
						tmp_c_1 = _hx_tmp_65
						if !(((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57))) {
							break
						}

						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_69 int32
								_ = _hx_tmp_69
								_hx_tmp_69 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_69
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_3 HxArray[int32]
								_ = tmp_v_3
								tmp_v_3 = this.Hx_Field_input
								var tmp_pos_3 int32
								_ = tmp_pos_3
								tmp_pos_3 = this.Hx_Field_pos
								var _hx_tmp_70 int32
								_ = _hx_tmp_70
								var _hx_tmp_71 bool
								_ = _hx_tmp_71
								if tmp_pos_3 < 0 {
									_hx_tmp_71 = true
								} else {
									var _hx_tmp_72 int32
									_ = _hx_tmp_72
									_hx_tmp_72 = tmp_pos_3
									var _hx_tmp_73 int32
									_ = _hx_tmp_73
									_hx_tmp_73 = _hx_tmp_72
									_hx_tmp_71 = (_hx_tmp_73 >= HxMakeArrayView[any](tmp_v_3).Len())
								}

								if _hx_tmp_71 {
									_hx_tmp_70 = -1
								} else {
									_hx_tmp_70 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
								}

								if (_hx_tmp_70) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp__hx_reserved_len_1 HxNullable[int32]
					_ = tmp__hx_reserved_len_1
					tmp__hx_reserved_len_1 = HxNullable[int32]{Value: (this.Hx_Field_pos - tmp_idStartPos_1), Valid: true}

					var _hx_tmp_74 HxArray[int32]
					_ = _hx_tmp_74
					_hx_tmp_74 = tmp_v_2
					var _hx_tmp_75 int32
					_ = _hx_tmp_75
					_hx_tmp_75 = tmp_idStartPos_1
					var _hx_tmp_76 int32
					_ = _hx_tmp_76
					if (tmp__hx_reserved_len_1).Valid != false {
						var _hx_tmp_77 int32
						_ = _hx_tmp_77
						_hx_tmp_77 = tmp_idStartPos_1
						var _hx_tmp_78 int32
						_ = _hx_tmp_78
						_hx_tmp_78 = _hx_tmp_77
						_hx_tmp_76 = (_hx_tmp_78 + (tmp__hx_reserved_len_1).Value)
					} else {
						_hx_tmp_76 = HxMakeArrayView[any](tmp_v_2).Len()
					}

					var tmp_name_1 string
					_ = tmp_name_1
					tmp_name_1 = Hx_Field_loreline_utf8vector_slice(_hx_tmp_74, _hx_tmp_75, _hx_tmp_76)
					var _hx_tmp_79 Hx_Enum_loreline_tokentype
					_ = _hx_tmp_79
					if func() func(_hx_param_key string) bool {
						var _hx_closure_src_147 func(_hx_param_key any) bool
						_ = _hx_closure_src_147
						_hx_closure_src_147 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_147(((any)(_hx_param_key)))
						}
					}()(tmp_name_1) {
						_hx_tmp_79 = func() func(_hx_param_key string) Hx_Enum_loreline_tokentype {
							var _hx_closure_src_148 func(_hx_param_key any) any
							_ = _hx_closure_src_148
							_hx_closure_src_148 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
							return func(_hx_param_key string) Hx_Enum_loreline_tokentype {
								return func() Hx_Enum_loreline_tokentype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_148(((any)(_hx_param_key))))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_tokentype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_tokentype)
								}()
							}
						}()(tmp_name_1)
					} else {
						_hx_tmp_79 = Hx_Enum_loreline_tokentype_Identifier{tmp_name_1}
					}

					var tmp_tokenType_1 Hx_Enum_loreline_tokentype
					_ = tmp_tokenType_1
					tmp_tokenType_1 = _hx_tmp_79
					var _hx_tmp_80 HxArray[*Hx_Obj_loreline_token]
					_ = _hx_tmp_80
					_hx_tmp_80 = tokens
					_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_80, Hx_Obj_loreline_token_CreateInstance(tmp_tokenType_1, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, (this.Hx_Field_column-(this.Hx_Field_pos-tmp_idStartPos_1)), tmp_idStartPos_1, nil))))
				} else {
					break hx_label_0
				}
			}

		case 91:
			{
				var _hx_tmp_81 HxArray[*Hx_Obj_loreline_token]
				_ = _hx_tmp_81
				_hx_tmp_81 = tokens
				_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_81, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_LBracket{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_82 int32
						_ = _hx_tmp_82
						_hx_tmp_82 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_82
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_83 int32
						_ = _hx_tmp_83
						var _hx_tmp_84 bool
						_ = _hx_tmp_84
						if tmp_pos_2 < 0 {
							_hx_tmp_84 = true
						} else {
							var _hx_tmp_85 int32
							_ = _hx_tmp_85
							_hx_tmp_85 = tmp_pos_2
							var _hx_tmp_86 int32
							_ = _hx_tmp_86
							_hx_tmp_86 = _hx_tmp_85
							_hx_tmp_84 = (_hx_tmp_86 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_84 {
							_hx_tmp_83 = -1
						} else {
							_hx_tmp_83 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_83) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				_ = (Hx_Array_Push[int32](this.Hx_Field_stack, ((int32)(14))))
				_ = (Hx_Array_Push[int32](this.Hx_Field_strictExprs, ((int32)(1))))
				var bracketLevel int32
				_ = bracketLevel
				bracketLevel = 1
				for (this.Hx_Field_pos < this.Hx_Field_length) && (bracketLevel > 0) {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_87 int32
					_ = _hx_tmp_87
					var _hx_tmp_88 bool
					_ = _hx_tmp_88
					if tmp_pos_2 < 0 {
						_hx_tmp_88 = true
					} else {
						var _hx_tmp_89 int32
						_ = _hx_tmp_89
						_hx_tmp_89 = tmp_pos_2
						var _hx_tmp_90 int32
						_ = _hx_tmp_90
						_hx_tmp_90 = _hx_tmp_89
						_hx_tmp_88 = (_hx_tmp_90 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_88 {
						_hx_tmp_87 = -1
					} else {
						_hx_tmp_87 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_87) == 93 {
						bracketLevel--
						if bracketLevel == 0 {
							var _hx_tmp_91 HxArray[*Hx_Obj_loreline_token]
							_ = _hx_tmp_91
							_hx_tmp_91 = tokens
							_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_91, Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_RBracket{}, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_pos, 1))))
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_92 int32
									_ = _hx_tmp_92
									_hx_tmp_92 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_92
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_3 HxArray[int32]
									_ = tmp_v_3
									tmp_v_3 = this.Hx_Field_input
									var tmp_pos_3 int32
									_ = tmp_pos_3
									tmp_pos_3 = this.Hx_Field_pos
									var _hx_tmp_93 int32
									_ = _hx_tmp_93
									var _hx_tmp_94 bool
									_ = _hx_tmp_94
									if tmp_pos_3 < 0 {
										_hx_tmp_94 = true
									} else {
										var _hx_tmp_95 int32
										_ = _hx_tmp_95
										_hx_tmp_95 = tmp_pos_3
										var _hx_tmp_96 int32
										_ = _hx_tmp_96
										_hx_tmp_96 = _hx_tmp_95
										_hx_tmp_94 = (_hx_tmp_96 >= HxMakeArrayView[any](tmp_v_3).Len())
									}

									if _hx_tmp_94 {
										_hx_tmp_93 = -1
									} else {
										_hx_tmp_93 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
									}

									if (_hx_tmp_93) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							_ = (Hx_Array_Pop[int32](this.Hx_Field_stack))
							_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
							break
						}
					} else {
						var tmp_v_3 HxArray[int32]
						_ = tmp_v_3
						tmp_v_3 = this.Hx_Field_input
						var tmp_pos_3 int32
						_ = tmp_pos_3
						tmp_pos_3 = this.Hx_Field_pos
						var _hx_tmp_97 int32
						_ = _hx_tmp_97
						var _hx_tmp_98 bool
						_ = _hx_tmp_98
						if tmp_pos_3 < 0 {
							_hx_tmp_98 = true
						} else {
							var _hx_tmp_99 int32
							_ = _hx_tmp_99
							_hx_tmp_99 = tmp_pos_3
							var _hx_tmp_100 int32
							_ = _hx_tmp_100
							_hx_tmp_100 = _hx_tmp_99
							_hx_tmp_98 = (_hx_tmp_100 >= HxMakeArrayView[any](tmp_v_3).Len())
						}

						if _hx_tmp_98 {
							_hx_tmp_97 = -1
						} else {
							_hx_tmp_97 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_3).Get(tmp_pos_3))))
						}

						if (_hx_tmp_97) == 91 {
							bracketLevel++
						}
					}

					var _hx_tmp_101 HxArray[*Hx_Obj_loreline_token]
					_ = _hx_tmp_101
					_hx_tmp_101 = tokens
					_ = (Hx_Array_Push[*Hx_Obj_loreline_token](_hx_tmp_101, this.VTable.Hx_Field_nextToken()))
				}

				if bracketLevel > 0 {
					this.VTable.Hx_Field_error("Unterminated array access in interpolation", true)
				}
			}

		default:
			{
				break hx_label_0
			}
		}
	}

	_ = (Hx_Array_Pop[int32](this.Hx_Field_strictExprs))
	return ((Hx_Typedef_loreline_tokens)(tokens))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readIdentifierTokenInInterpolation(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	var startPos *Hx_Obj_loreline_position
	_ = startPos
	startPos = this.VTable.Hx_Field_makePositionRelativeTo(stringStart)
	var startOffset int32
	_ = startOffset
	startOffset = this.Hx_Field_pos
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
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

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_4
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if tmp_pos_1 < 0 {
					_hx_tmp_6 = true
				} else {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = tmp_pos_1
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_6 {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_5) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - startOffset), Valid: true}

	var _hx_tmp_9 HxArray[int32]
	_ = _hx_tmp_9
	_hx_tmp_9 = v
	var _hx_tmp_10 int32
	_ = _hx_tmp_10
	_hx_tmp_10 = startOffset
	var _hx_tmp_11 int32
	_ = _hx_tmp_11
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		_hx_tmp_12 = startOffset
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_12
		_hx_tmp_11 = (_hx_tmp_13 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_11 = HxMakeArrayView[any](v).Len()
	}

	var name string
	_ = name
	name = Hx_Field_loreline_utf8vector_slice(_hx_tmp_9, _hx_tmp_10, _hx_tmp_11)
	var _hx_tmp_14 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_14
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_149 func(_hx_param_key any) bool
		_ = _hx_closure_src_149
		_hx_closure_src_149 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_149(((any)(_hx_param_key)))
		}
	}()(name) {
		_hx_tmp_14 = func() func(_hx_param_key string) Hx_Enum_loreline_tokentype {
			var _hx_closure_src_150 func(_hx_param_key any) any
			_ = _hx_closure_src_150
			_hx_closure_src_150 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) Hx_Enum_loreline_tokentype {
				return func() Hx_Enum_loreline_tokentype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_150(((any)(_hx_param_key))))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_tokentype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_tokentype)
				}()
			}
		}()(name)
	} else {
		_hx_tmp_14 = Hx_Enum_loreline_tokentype_Identifier{name}
	}

	var tokenType Hx_Enum_loreline_tokentype
	_ = tokenType
	tokenType = this.VTable.Hx_Field_mergeIsNot(_hx_tmp_14)
	return Hx_Obj_loreline_token_CreateInstance(tokenType, startPos)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_makePositionRelativeTo(stringStart *Hx_Obj_loreline_position) *Hx_Obj_loreline_position {
	var line int32
	_ = line
	line = stringStart.Hx_Field_line
	var column int32
	_ = column
	column = stringStart.Hx_Field_column
	var i int32
	_ = i
	i = stringStart.Hx_Field_offset
	for i < this.Hx_Field_pos {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if i < 0 {
			_hx_tmp_1 = true
		} else {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = i
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(i))))
		}

		if (_hx_tmp_0) == 10 {
			line++
			column = 1
		} else {
			column++
		}

		i++
	}

	return Hx_Obj_loreline_position_CreateInstance(line, column, this.Hx_Field_pos, stringStart.Hx_Field_length)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readLineComment() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 2
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	var contentStart int32
	_ = contentStart
	contentStart = this.Hx_Field_pos
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_5
		if (c == 10) || (c == 13) {
			break
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_9
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				if tmp_pos_1 < 0 {
					_hx_tmp_11 = true
				} else {
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					_hx_tmp_12 = tmp_pos_1
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = _hx_tmp_12
					_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_11 {
					_hx_tmp_10 = -1
				} else {
					_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_10) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - contentStart), Valid: true}

	var _hx_tmp_14 HxArray[int32]
	_ = _hx_tmp_14
	_hx_tmp_14 = v
	var _hx_tmp_15 int32
	_ = _hx_tmp_15
	_hx_tmp_15 = contentStart
	var _hx_tmp_16 int32
	_ = _hx_tmp_16
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_17 int32
		_ = _hx_tmp_17
		_hx_tmp_17 = contentStart
		var _hx_tmp_18 int32
		_ = _hx_tmp_18
		_hx_tmp_18 = _hx_tmp_17
		_hx_tmp_16 = (_hx_tmp_18 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_16 = HxMakeArrayView[any](v).Len()
	}

	var _hx_reserved_type Hx_Enum_loreline_tokentype
	_ = _hx_reserved_type
	_hx_reserved_type = Hx_Enum_loreline_tokentype_CommentLine{Hx_Field_loreline_utf8vector_slice(_hx_tmp_14, _hx_tmp_15, _hx_tmp_16)}

	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type, position)
	this.Hx_Field_previous = token
	return token
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readMultiLineComment() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 2
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	var contentStart int32
	_ = contentStart
	contentStart = this.Hx_Field_pos
	var nestLevel int32
	_ = nestLevel
	nestLevel = 1
	for (this.Hx_Field_pos < this.Hx_Field_length) && (nestLevel > 0) {
		var tmp bool
		_ = tmp
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		if (_hx_tmp_5) == 42 {
			var offset int32
			_ = offset
			offset = 1
			var tmp1 int32
			_ = tmp1
			if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = (this.Hx_Field_pos + offset)
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

				tmp1 = _hx_tmp_9
			} else {
				tmp1 = 0
			}

			tmp = (tmp1 == 47)
		} else {
			tmp = false
		}

		if tmp {
			nestLevel--
			if nestLevel == 0 {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - contentStart), Valid: true}

				var _hx_tmp_13 HxArray[int32]
				_ = _hx_tmp_13
				_hx_tmp_13 = tmp_v_1
				var _hx_tmp_14 int32
				_ = _hx_tmp_14
				_hx_tmp_14 = contentStart
				var _hx_tmp_15 int32
				_ = _hx_tmp_15
				if (_hx_reserved_len).Valid != false {
					var _hx_tmp_16 int32
					_ = _hx_tmp_16
					_hx_tmp_16 = contentStart
					var _hx_tmp_17 int32
					_ = _hx_tmp_17
					_hx_tmp_17 = _hx_tmp_16
					_hx_tmp_15 = (_hx_tmp_17 + (_hx_reserved_len).Value)
				} else {
					_hx_tmp_15 = HxMakeArrayView[any](tmp_v_1).Len()
				}

				var content string
				_ = content
				content = Hx_Field_loreline_utf8vector_slice(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 2
					for _g < _g1 {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_18
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = this.Hx_Field_pos
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						var _hx_tmp_20 bool
						_ = _hx_tmp_20
						if tmp_pos_1 < 0 {
							_hx_tmp_20 = true
						} else {
							var _hx_tmp_21 int32
							_ = _hx_tmp_21
							_hx_tmp_21 = tmp_pos_1
							var _hx_tmp_22 int32
							_ = _hx_tmp_22
							_hx_tmp_22 = _hx_tmp_21
							_hx_tmp_20 = (_hx_tmp_22 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_20 {
							_hx_tmp_19 = -1
						} else {
							_hx_tmp_19 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_1))))
						}

						if (_hx_tmp_19) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				var position *Hx_Obj_loreline_position
				_ = position
				position = start
				if position == nil {
					position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
				}

				position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
				var token *Hx_Obj_loreline_token
				_ = token
				token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_CommentMultiLine{content}, position)
				this.Hx_Field_previous = token
				return token
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 2
				for _g < _g1 {
					var _hx_tmp_23 int32
					_ = _hx_tmp_23
					_hx_tmp_23 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_23
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					var _hx_tmp_25 bool
					_ = _hx_tmp_25
					if tmp_pos_1 < 0 {
						_hx_tmp_25 = true
					} else {
						var _hx_tmp_26 int32
						_ = _hx_tmp_26
						_hx_tmp_26 = tmp_pos_1
						var _hx_tmp_27 int32
						_ = _hx_tmp_27
						_hx_tmp_27 = _hx_tmp_26
						_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_25 {
						_hx_tmp_24 = -1
					} else {
						_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_24) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		} else {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
			var _hx_tmp_28 int32
			_ = _hx_tmp_28
			var _hx_tmp_29 bool
			_ = _hx_tmp_29
			if tmp_pos_1 < 0 {
				_hx_tmp_29 = true
			} else {
				var _hx_tmp_30 int32
				_ = _hx_tmp_30
				_hx_tmp_30 = tmp_pos_1
				var _hx_tmp_31 int32
				_ = _hx_tmp_31
				_hx_tmp_31 = _hx_tmp_30
				_hx_tmp_29 = (_hx_tmp_31 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_29 {
				_hx_tmp_28 = -1
			} else {
				_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			if (_hx_tmp_28) == 47 {
				var offset int32
				_ = offset
				offset = 1
				var tmp1 int32
				_ = tmp1
				if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = (this.Hx_Field_pos + offset)
					var _hx_tmp_32 int32
					_ = _hx_tmp_32
					var _hx_tmp_33 bool
					_ = _hx_tmp_33
					if tmp_pos_2 < 0 {
						_hx_tmp_33 = true
					} else {
						var _hx_tmp_34 int32
						_ = _hx_tmp_34
						_hx_tmp_34 = tmp_pos_2
						var _hx_tmp_35 int32
						_ = _hx_tmp_35
						_hx_tmp_35 = _hx_tmp_34
						_hx_tmp_33 = (_hx_tmp_35 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_33 {
						_hx_tmp_32 = -1
					} else {
						_hx_tmp_32 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					tmp1 = _hx_tmp_32
				} else {
					tmp1 = 0
				}

				tmp_tmp_1 = (tmp1 == 42)
			} else {
				tmp_tmp_1 = false
			}

			if tmp_tmp_1 {
				nestLevel++
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 2
					for _g < _g1 {
						var _hx_tmp_36 int32
						_ = _hx_tmp_36
						_hx_tmp_36 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_36
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_37 int32
						_ = _hx_tmp_37
						var _hx_tmp_38 bool
						_ = _hx_tmp_38
						if tmp_pos_2 < 0 {
							_hx_tmp_38 = true
						} else {
							var _hx_tmp_39 int32
							_ = _hx_tmp_39
							_hx_tmp_39 = tmp_pos_2
							var _hx_tmp_40 int32
							_ = _hx_tmp_40
							_hx_tmp_40 = _hx_tmp_39
							_hx_tmp_38 = (_hx_tmp_40 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_38 {
							_hx_tmp_37 = -1
						} else {
							_hx_tmp_37 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_37) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}
			} else {
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_41 int32
					_ = _hx_tmp_41
					_hx_tmp_41 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_41
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_42 int32
					_ = _hx_tmp_42
					var _hx_tmp_43 bool
					_ = _hx_tmp_43
					if tmp_pos_2 < 0 {
						_hx_tmp_43 = true
					} else {
						var _hx_tmp_44 int32
						_ = _hx_tmp_44
						_hx_tmp_44 = tmp_pos_2
						var _hx_tmp_45 int32
						_ = _hx_tmp_45
						_hx_tmp_45 = _hx_tmp_44
						_hx_tmp_43 = (_hx_tmp_45 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_43 {
						_hx_tmp_42 = -1
					} else {
						_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_42) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}
	}

	this.VTable.Hx_Field_error("Unterminated multi-line comment", true)
	return nil
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readHashComment() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 1
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	var contentStart int32
	_ = contentStart
	contentStart = this.Hx_Field_pos
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_5
		if (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57))) || (c == 45) {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_9
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				if tmp_pos_1 < 0 {
					_hx_tmp_11 = true
				} else {
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					_hx_tmp_12 = tmp_pos_1
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = _hx_tmp_12
					_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_11 {
					_hx_tmp_10 = -1
				} else {
					_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_10) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		} else {
			break
		}
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - contentStart), Valid: true}

	var _hx_tmp_14 HxArray[int32]
	_ = _hx_tmp_14
	_hx_tmp_14 = v
	var _hx_tmp_15 int32
	_ = _hx_tmp_15
	_hx_tmp_15 = contentStart
	var _hx_tmp_16 int32
	_ = _hx_tmp_16
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_17 int32
		_ = _hx_tmp_17
		_hx_tmp_17 = contentStart
		var _hx_tmp_18 int32
		_ = _hx_tmp_18
		_hx_tmp_18 = _hx_tmp_17
		_hx_tmp_16 = (_hx_tmp_18 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_16 = HxMakeArrayView[any](v).Len()
	}

	var _hx_reserved_type Hx_Enum_loreline_tokentype
	_ = _hx_reserved_type
	_hx_reserved_type = Hx_Enum_loreline_tokentype_CommentHash{Hx_Field_loreline_utf8vector_slice(_hx_tmp_14, _hx_tmp_15, _hx_tmp_16)}

	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type, position)
	this.Hx_Field_previous = token
	return token
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readNumber() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	var startPos int32
	_ = startPos
	startPos = this.Hx_Field_pos
	for {
		var tmp bool
		_ = tmp
		if this.Hx_Field_pos < this.Hx_Field_length {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
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
			tmp = ((c >= 48) && (c <= 57))
		} else {
			tmp = false
		}

		if !tmp {
			break
		}

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_4
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var pos int32
				_ = pos
				pos = this.Hx_Field_pos
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if pos < 0 {
					_hx_tmp_6 = true
				} else {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = pos
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_6 {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				if (_hx_tmp_5) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var tmp bool
	_ = tmp
	var tmp1 bool
	_ = tmp1
	if this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_9 int32
		_ = _hx_tmp_9
		var _hx_tmp_10 bool
		_ = _hx_tmp_10
		if pos < 0 {
			_hx_tmp_10 = true
		} else {
			var _hx_tmp_11 int32
			_ = _hx_tmp_11
			_hx_tmp_11 = pos
			var _hx_tmp_12 int32
			_ = _hx_tmp_12
			_hx_tmp_12 = _hx_tmp_11
			_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_10 {
			_hx_tmp_9 = -1
		} else {
			_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		tmp1 = ((_hx_tmp_9) == 46)
	} else {
		tmp1 = false
	}

	if tmp1 && ((this.Hx_Field_pos + 1) < this.Hx_Field_length) {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = (this.Hx_Field_pos + 1)
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		var _hx_tmp_14 bool
		_ = _hx_tmp_14
		if pos < 0 {
			_hx_tmp_14 = true
		} else {
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = pos
			var _hx_tmp_16 int32
			_ = _hx_tmp_16
			_hx_tmp_16 = _hx_tmp_15
			_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_14 {
			_hx_tmp_13 = -1
		} else {
			_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_13
		tmp = ((c >= 48) && (c <= 57))
	} else {
		tmp = false
	}

	if tmp {
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_17 int32
				_ = _hx_tmp_17
				_hx_tmp_17 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_17
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var pos int32
				_ = pos
				pos = this.Hx_Field_pos
				var _hx_tmp_18 int32
				_ = _hx_tmp_18
				var _hx_tmp_19 bool
				_ = _hx_tmp_19
				if pos < 0 {
					_hx_tmp_19 = true
				} else {
					var _hx_tmp_20 int32
					_ = _hx_tmp_20
					_hx_tmp_20 = pos
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					_hx_tmp_21 = _hx_tmp_20
					_hx_tmp_19 = (_hx_tmp_21 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_19 {
					_hx_tmp_18 = -1
				} else {
					_hx_tmp_18 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				if (_hx_tmp_18) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}

		for {
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if this.Hx_Field_pos < this.Hx_Field_length {
				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var pos int32
				_ = pos
				pos = this.Hx_Field_pos
				var _hx_tmp_22 int32
				_ = _hx_tmp_22
				var _hx_tmp_23 bool
				_ = _hx_tmp_23
				if pos < 0 {
					_hx_tmp_23 = true
				} else {
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = pos
					var _hx_tmp_25 int32
					_ = _hx_tmp_25
					_hx_tmp_25 = _hx_tmp_24
					_hx_tmp_23 = (_hx_tmp_25 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_23 {
					_hx_tmp_22 = -1
				} else {
					_hx_tmp_22 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				var c int32
				_ = c
				c = _hx_tmp_22
				tmp_tmp_1 = ((c >= 48) && (c <= 57))
			} else {
				tmp_tmp_1 = false
			}

			if !tmp_tmp_1 {
				break
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_26 int32
					_ = _hx_tmp_26
					_hx_tmp_26 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_26
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var v HxArray[int32]
					_ = v
					v = this.Hx_Field_input
					var pos int32
					_ = pos
					pos = this.Hx_Field_pos
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					var _hx_tmp_28 bool
					_ = _hx_tmp_28
					if pos < 0 {
						_hx_tmp_28 = true
					} else {
						var _hx_tmp_29 int32
						_ = _hx_tmp_29
						_hx_tmp_29 = pos
						var _hx_tmp_30 int32
						_ = _hx_tmp_30
						_hx_tmp_30 = _hx_tmp_29
						_hx_tmp_28 = (_hx_tmp_30 >= HxMakeArrayView[any](v).Len())
					}

					if _hx_tmp_28 {
						_hx_tmp_27 = -1
					} else {
						_hx_tmp_27 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
					}

					if (_hx_tmp_27) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - startPos), Valid: true}

	var _hx_tmp_31 HxArray[int32]
	_ = _hx_tmp_31
	_hx_tmp_31 = v
	var _hx_tmp_32 int32
	_ = _hx_tmp_32
	_hx_tmp_32 = startPos
	var _hx_tmp_33 int32
	_ = _hx_tmp_33
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_34 int32
		_ = _hx_tmp_34
		_hx_tmp_34 = startPos
		var _hx_tmp_35 int32
		_ = _hx_tmp_35
		_hx_tmp_35 = _hx_tmp_34
		_hx_tmp_33 = (_hx_tmp_35 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_33 = HxMakeArrayView[any](v).Len()
	}

	var _hx_reserved_type Hx_Enum_loreline_tokentype
	_ = _hx_reserved_type
	_hx_reserved_type = Hx_Enum_loreline_tokentype_LNumber{Hx_Field_std_parseFloat(Hx_Field_loreline_utf8vector_slice(_hx_tmp_31, _hx_tmp_32, _hx_tmp_33))}

	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type, position)
	this.Hx_Field_previous = token
	var token1 *Hx_Obj_loreline_token
	_ = token1
	token1 = token
	return token1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_mergeIsNot(tokenType Hx_Enum_loreline_tokentype) Hx_Enum_loreline_tokentype {
	if tokenType.Hx_Field_enumIndex() == 32 {
		if tokenType.(Hx_Enum_loreline_tokentype_OpEquals).Hx_Field_word == true {
			var end int32
			_ = end
			end = this.VTable.Hx_Field_wordEnd(this.VTable.Hx_Field_skipSpacesAndTabs(this.Hx_Field_pos), "not")
			if end != -1 {
				for this.Hx_Field_pos < end {
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_0
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var v HxArray[int32]
						_ = v
						v = this.Hx_Field_input
						var pos int32
						_ = pos
						pos = this.Hx_Field_pos
						var _hx_tmp_1 int32
						_ = _hx_tmp_1
						var _hx_tmp_2 bool
						_ = _hx_tmp_2
						if pos < 0 {
							_hx_tmp_2 = true
						} else {
							var _hx_tmp_3 int32
							_ = _hx_tmp_3
							_hx_tmp_3 = pos
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = _hx_tmp_3
							_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
						}

						if _hx_tmp_2 {
							_hx_tmp_1 = -1
						} else {
							_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
						}

						if (_hx_tmp_1) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				return Hx_Enum_loreline_tokentype_OpNotEquals{true}
			}
		}
	}

	return tokenType
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_wordEnd(p int32, word string) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = p
	var end int32
	_ = end
	end = (_hx_tmp_0 + int32(utf8.RuneCountInString(word)))
	if end > this.Hx_Field_length {
		return -1
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(word))
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (p + i)
			var hx_cmp_1 HxNullable[int32]
			_ = hx_cmp_1
			hx_cmp_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(word, i))
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if (hx_cmp_1).Valid == false {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if pos < 0 {
					_hx_tmp_4 = true
				} else {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = pos
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_5
					_hx_tmp_4 = (_hx_tmp_6 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_4 {
					_hx_tmp_3 = -1
				} else {
					_hx_tmp_3 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				_hx_tmp_2 = ((_hx_tmp_3) != (hx_cmp_1).Value)
			}

			if _hx_tmp_2 {
				return -1
			}
		}
	}

	var tmp bool
	_ = tmp
	if end < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		var _hx_tmp_8 bool
		_ = _hx_tmp_8
		if end < 0 {
			_hx_tmp_8 = true
		} else {
			var _hx_tmp_9 int32
			_ = _hx_tmp_9
			_hx_tmp_9 = end
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = _hx_tmp_9
			_hx_tmp_8 = (_hx_tmp_10 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_8 {
			_hx_tmp_7 = -1
		} else {
			_hx_tmp_7 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(end))))
		}

		var c int32
		_ = c
		c = _hx_tmp_7
		tmp = (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)))
	} else {
		tmp = false
	}

	if tmp {
		return -1
	}

	return end
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_skipSpacesAndTabs(p int32) int32 {
	for {
		var tmp bool
		_ = tmp
		if p < this.Hx_Field_length {
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

		p++
	}

	return p
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readIdentifier() *Hx_Obj_loreline_token {
	var start *Hx_Obj_loreline_position
	_ = start
	start = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	var startPos int32
	_ = startPos
	startPos = this.Hx_Field_pos
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
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

		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_4
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if tmp_pos_1 < 0 {
					_hx_tmp_6 = true
				} else {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = tmp_pos_1
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_6 {
					_hx_tmp_5 = -1
				} else {
					_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_5) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - startPos), Valid: true}

	var _hx_tmp_9 HxArray[int32]
	_ = _hx_tmp_9
	_hx_tmp_9 = v
	var _hx_tmp_10 int32
	_ = _hx_tmp_10
	_hx_tmp_10 = startPos
	var _hx_tmp_11 int32
	_ = _hx_tmp_11
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_12 int32
		_ = _hx_tmp_12
		_hx_tmp_12 = startPos
		var _hx_tmp_13 int32
		_ = _hx_tmp_13
		_hx_tmp_13 = _hx_tmp_12
		_hx_tmp_11 = (_hx_tmp_13 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_11 = HxMakeArrayView[any](v).Len()
	}

	var word string
	_ = word
	word = Hx_Field_loreline_utf8vector_slice(_hx_tmp_9, _hx_tmp_10, _hx_tmp_11)
	if word == "function" {
		return this.VTable.Hx_Field_readFunction(start)
	}

	var _hx_tmp_14 Hx_Enum_loreline_tokentype
	_ = _hx_tmp_14
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_151 func(_hx_param_key any) bool
		_ = _hx_closure_src_151
		_hx_closure_src_151 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_151(((any)(_hx_param_key)))
		}
	}()(word) {
		_hx_tmp_14 = func() func(_hx_param_key string) Hx_Enum_loreline_tokentype {
			var _hx_closure_src_152 func(_hx_param_key any) any
			_ = _hx_closure_src_152
			_hx_closure_src_152 = Hx_Field_loreline_lexer_KEYWORDS.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) Hx_Enum_loreline_tokentype {
				return func() Hx_Enum_loreline_tokentype {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(_hx_closure_src_152(((any)(_hx_param_key))))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_tokentype
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_tokentype)
				}()
			}
		}()(word)
	} else {
		_hx_tmp_14 = Hx_Enum_loreline_tokentype_Identifier{word}
	}

	var tokenType Hx_Enum_loreline_tokentype
	_ = tokenType
	tokenType = _hx_tmp_14
	var _hx_tmp_15 bool
	_ = _hx_tmp_15
	if word == "when" {
		_hx_tmp_15 = this.VTable.Hx_Field_isWhenStart(startPos)
	} else {
		_hx_tmp_15 = false
	}

	if _hx_tmp_15 {
		tokenType = Hx_Enum_loreline_tokentype_KwWhen{}
	}

	if (this.Hx_Field_previous != nil) && (this.Hx_Field_previous.Hx_Field_type == Hx_Enum_loreline_tokentype_Dot{}) {
		if tokenType != nil {
			switch tokenType.Hx_Field_enumIndex() {
			case 9, 10, 11, 12, 13:
				{
					tokenType = Hx_Enum_loreline_tokentype_Identifier{word}
				}

			case 32:
				{
					if tokenType.(Hx_Enum_loreline_tokentype_OpEquals).Hx_Field_word == true {
						tokenType = Hx_Enum_loreline_tokentype_Identifier{word}
					}
				}

			case 40:
				{
					if tokenType.(Hx_Enum_loreline_tokentype_OpNot).Hx_Field_word == true {
						tokenType = Hx_Enum_loreline_tokentype_Identifier{word}
					}
				}

			default:
				{
				}
			}
		}
	}

	if ((this.Hx_Field_previous != nil) && (this.Hx_Field_previous.Hx_Field_type == Hx_Enum_loreline_tokentype_KwWhen{})) && (this.Hx_Field_previous.Hx_Field_pos.Hx_Field_line == start.Hx_Field_line) {
		tokenType = Hx_Enum_loreline_tokentype_Identifier{word}
	}

	tokenType = this.VTable.Hx_Field_mergeIsNot(tokenType)
	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(tokenType, position)
	this.Hx_Field_previous = token
	return token
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_readFunction(start *Hx_Obj_loreline_position) *Hx_Obj_loreline_token {
	{
		var newPos int32
		_ = newPos
		newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
		for this.Hx_Field_pos < newPos {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_0
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var v HxArray[int32]
				_ = v
				v = this.Hx_Field_input
				var pos int32
				_ = pos
				pos = this.Hx_Field_pos
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				var _hx_tmp_2 bool
				_ = _hx_tmp_2
				if pos < 0 {
					_hx_tmp_2 = true
				} else {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = pos
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _hx_tmp_3
					_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
				}

				if _hx_tmp_2 {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
				}

				if (_hx_tmp_1) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var minIndent int32
	_ = minIndent
	minIndent = start.Hx_Field_column
	var name HxNullable[string]
	_ = name
	name = HxNullable[string]{Value: HxStringNull}

	var v HxArray[int32]
	_ = v
	v = this.Hx_Field_input
	var pos int32
	_ = pos
	pos = this.Hx_Field_pos
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	if pos < 0 {
		_hx_tmp_6 = true
	} else {
		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		_hx_tmp_7 = pos
		var _hx_tmp_8 int32
		_ = _hx_tmp_8
		_hx_tmp_8 = _hx_tmp_7
		_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
	}

	if _hx_tmp_6 {
		_hx_tmp_5 = -1
	} else {
		_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
	}

	var c int32
	_ = c
	c = _hx_tmp_5
	if (((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95) {
		var nameStart int32
		_ = nameStart
		nameStart = this.Hx_Field_pos
		for {
			var tmp bool
			_ = tmp
			if this.Hx_Field_pos < this.Hx_Field_length {
				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
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

				var tmp_c_1 int32
				_ = tmp_c_1
				tmp_c_1 = _hx_tmp_9
				tmp = (((((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95)) || ((tmp_c_1 >= 48) && (tmp_c_1 <= 57)))
			} else {
				tmp = false
			}

			if !tmp {
				break
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_13 int32
					_ = _hx_tmp_13
					_hx_tmp_13 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_13
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_14 int32
					_ = _hx_tmp_14
					var _hx_tmp_15 bool
					_ = _hx_tmp_15
					if tmp_pos_1 < 0 {
						_hx_tmp_15 = true
					} else {
						var _hx_tmp_16 int32
						_ = _hx_tmp_16
						_hx_tmp_16 = tmp_pos_1
						var _hx_tmp_17 int32
						_ = _hx_tmp_17
						_hx_tmp_17 = _hx_tmp_16
						_hx_tmp_15 = (_hx_tmp_17 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_15 {
						_hx_tmp_14 = -1
					} else {
						_hx_tmp_14 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_14) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}

		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var _hx_reserved_len HxNullable[int32]
		_ = _hx_reserved_len
		_hx_reserved_len = HxNullable[int32]{Value: (this.Hx_Field_pos - nameStart), Valid: true}

		var _hx_tmp_18 HxArray[int32]
		_ = _hx_tmp_18
		_hx_tmp_18 = tmp_v_1
		var _hx_tmp_19 int32
		_ = _hx_tmp_19
		_hx_tmp_19 = nameStart
		var _hx_tmp_20 int32
		_ = _hx_tmp_20
		if (_hx_reserved_len).Valid != false {
			var _hx_tmp_21 int32
			_ = _hx_tmp_21
			_hx_tmp_21 = nameStart
			var _hx_tmp_22 int32
			_ = _hx_tmp_22
			_hx_tmp_22 = _hx_tmp_21
			_hx_tmp_20 = (_hx_tmp_22 + (_hx_reserved_len).Value)
		} else {
			_hx_tmp_20 = HxMakeArrayView[any](tmp_v_1).Len()
		}

		name = HxNullableString(Hx_Field_loreline_utf8vector_slice(_hx_tmp_18, _hx_tmp_19, _hx_tmp_20))
	}

	{
		var newPos1 int32
		_ = newPos1
		newPos1 = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
		for this.Hx_Field_pos < newPos1 {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_23 int32
				_ = _hx_tmp_23
				_hx_tmp_23 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_23
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_24 int32
				_ = _hx_tmp_24
				var _hx_tmp_25 bool
				_ = _hx_tmp_25
				if tmp_pos_1 < 0 {
					_hx_tmp_25 = true
				} else {
					var _hx_tmp_26 int32
					_ = _hx_tmp_26
					_hx_tmp_26 = tmp_pos_1
					var _hx_tmp_27 int32
					_ = _hx_tmp_27
					_hx_tmp_27 = _hx_tmp_26
					_hx_tmp_25 = (_hx_tmp_27 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_25 {
					_hx_tmp_24 = -1
				} else {
					_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_24) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var tmp bool
	_ = tmp
	if this.Hx_Field_pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
		var _hx_tmp_28 int32
		_ = _hx_tmp_28
		var _hx_tmp_29 bool
		_ = _hx_tmp_29
		if tmp_pos_1 < 0 {
			_hx_tmp_29 = true
		} else {
			var _hx_tmp_30 int32
			_ = _hx_tmp_30
			_hx_tmp_30 = tmp_pos_1
			var _hx_tmp_31 int32
			_ = _hx_tmp_31
			_hx_tmp_31 = _hx_tmp_30
			_hx_tmp_29 = (_hx_tmp_31 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_29 {
			_hx_tmp_28 = -1
		} else {
			_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
		}

		tmp = ((_hx_tmp_28) != 40)
	} else {
		tmp = true
	}

	if tmp {
		this.VTable.Hx_Field_error("Expected opening parenthesis after function name", true)
	}

	var parenLevel int32
	_ = parenLevel
	parenLevel = 0
	var args HxArray[string]
	_ = args
	args = HxMakeArray[string]()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 1
		for _g < _g1 {
			var _hx_tmp_32 int32
			_ = _hx_tmp_32
			_hx_tmp_32 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_32
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
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

			if (_hx_tmp_33) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	parenLevel = 1
	var currentArg *Hx_Obj_stringbuf
	_ = currentArg
	currentArg = Hx_Obj_stringbuf_CreateInstance()
	for (this.Hx_Field_pos < this.Hx_Field_length) && (parenLevel > 0) {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
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

		var tmp_c_1 int32
		_ = tmp_c_1
		tmp_c_1 = _hx_tmp_37
		if tmp_c_1 == 47 {
			var prevPos int32
			_ = prevPos
			prevPos = this.Hx_Field_pos
			{
				var newPos int32
				_ = newPos
				newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
				for this.Hx_Field_pos < newPos {
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_41 int32
						_ = _hx_tmp_41
						_hx_tmp_41 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_41
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_42 int32
						_ = _hx_tmp_42
						var _hx_tmp_43 bool
						_ = _hx_tmp_43
						if tmp_pos_2 < 0 {
							_hx_tmp_43 = true
						} else {
							var _hx_tmp_44 int32
							_ = _hx_tmp_44
							_hx_tmp_44 = tmp_pos_2
							var _hx_tmp_45 int32
							_ = _hx_tmp_45
							_hx_tmp_45 = _hx_tmp_44
							_hx_tmp_43 = (_hx_tmp_45 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_43 {
							_hx_tmp_42 = -1
						} else {
							_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_42) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}
			}

			if this.Hx_Field_pos == prevPos {
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_46 int32
						_ = _hx_tmp_46
						_hx_tmp_46 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_46
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_47 int32
						_ = _hx_tmp_47
						var _hx_tmp_48 bool
						_ = _hx_tmp_48
						if tmp_pos_2 < 0 {
							_hx_tmp_48 = true
						} else {
							var _hx_tmp_49 int32
							_ = _hx_tmp_49
							_hx_tmp_49 = tmp_pos_2
							var _hx_tmp_50 int32
							_ = _hx_tmp_50
							_hx_tmp_50 = _hx_tmp_49
							_hx_tmp_48 = (_hx_tmp_50 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_48 {
							_hx_tmp_47 = -1
						} else {
							_hx_tmp_47 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_47) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				this.VTable.Hx_Field_error("Invalid character \"/\"", false)
			}
		} else {
			if tmp_c_1 == 40 {
				parenLevel++
				currentArg.VTable.Hx_Field_addChar(tmp_c_1)
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_51 int32
						_ = _hx_tmp_51
						_hx_tmp_51 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_51
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_52 int32
						_ = _hx_tmp_52
						var _hx_tmp_53 bool
						_ = _hx_tmp_53
						if tmp_pos_2 < 0 {
							_hx_tmp_53 = true
						} else {
							var _hx_tmp_54 int32
							_ = _hx_tmp_54
							_hx_tmp_54 = tmp_pos_2
							var _hx_tmp_55 int32
							_ = _hx_tmp_55
							_hx_tmp_55 = _hx_tmp_54
							_hx_tmp_53 = (_hx_tmp_55 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_53 {
							_hx_tmp_52 = -1
						} else {
							_hx_tmp_52 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_52) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				{
					var newPos int32
					_ = newPos
					newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
					for this.Hx_Field_pos < newPos {
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_56 int32
							_ = _hx_tmp_56
							_hx_tmp_56 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_56
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var tmp_pos_2 int32
							_ = tmp_pos_2
							tmp_pos_2 = this.Hx_Field_pos
							var _hx_tmp_57 int32
							_ = _hx_tmp_57
							var _hx_tmp_58 bool
							_ = _hx_tmp_58
							if tmp_pos_2 < 0 {
								_hx_tmp_58 = true
							} else {
								var _hx_tmp_59 int32
								_ = _hx_tmp_59
								_hx_tmp_59 = tmp_pos_2
								var _hx_tmp_60 int32
								_ = _hx_tmp_60
								_hx_tmp_60 = _hx_tmp_59
								_hx_tmp_58 = (_hx_tmp_60 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_58 {
								_hx_tmp_57 = -1
							} else {
								_hx_tmp_57 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
							}

							if (_hx_tmp_57) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}
				}
			} else {
				if tmp_c_1 == 41 {
					parenLevel--
					if parenLevel > 0 {
						currentArg.VTable.Hx_Field_addChar(tmp_c_1)
					} else {
						var argStr string
						_ = argStr
						argStr = Hx_Field_stringtools_trim(currentArg.VTable.Hx_Field_toString())
						if int32(utf8.RuneCountInString(argStr)) > 0 {
							_ = (Hx_Array_Push[string](args, argStr))
						}
					}

					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_61 int32
							_ = _hx_tmp_61
							_hx_tmp_61 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_61
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var tmp_pos_2 int32
							_ = tmp_pos_2
							tmp_pos_2 = this.Hx_Field_pos
							var _hx_tmp_62 int32
							_ = _hx_tmp_62
							var _hx_tmp_63 bool
							_ = _hx_tmp_63
							if tmp_pos_2 < 0 {
								_hx_tmp_63 = true
							} else {
								var _hx_tmp_64 int32
								_ = _hx_tmp_64
								_hx_tmp_64 = tmp_pos_2
								var _hx_tmp_65 int32
								_ = _hx_tmp_65
								_hx_tmp_65 = _hx_tmp_64
								_hx_tmp_63 = (_hx_tmp_65 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_63 {
								_hx_tmp_62 = -1
							} else {
								_hx_tmp_62 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
							}

							if (_hx_tmp_62) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}

					{
						var newPos int32
						_ = newPos
						newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
						for this.Hx_Field_pos < newPos {
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_66 int32
								_ = _hx_tmp_66
								_hx_tmp_66 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_66
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_2 HxArray[int32]
								_ = tmp_v_2
								tmp_v_2 = this.Hx_Field_input
								var tmp_pos_2 int32
								_ = tmp_pos_2
								tmp_pos_2 = this.Hx_Field_pos
								var _hx_tmp_67 int32
								_ = _hx_tmp_67
								var _hx_tmp_68 bool
								_ = _hx_tmp_68
								if tmp_pos_2 < 0 {
									_hx_tmp_68 = true
								} else {
									var _hx_tmp_69 int32
									_ = _hx_tmp_69
									_hx_tmp_69 = tmp_pos_2
									var _hx_tmp_70 int32
									_ = _hx_tmp_70
									_hx_tmp_70 = _hx_tmp_69
									_hx_tmp_68 = (_hx_tmp_70 >= HxMakeArrayView[any](tmp_v_2).Len())
								}

								if _hx_tmp_68 {
									_hx_tmp_67 = -1
								} else {
									_hx_tmp_67 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
								}

								if (_hx_tmp_67) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}
					}
				} else {
					if (tmp_c_1 == 44) && (parenLevel == 1) {
						var argStr string
						_ = argStr
						argStr = Hx_Field_stringtools_trim(currentArg.VTable.Hx_Field_toString())
						if int32(utf8.RuneCountInString(argStr)) > 0 {
							_ = (Hx_Array_Push[string](args, argStr))
						}

						currentArg = Hx_Obj_stringbuf_CreateInstance()
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_71 int32
								_ = _hx_tmp_71
								_hx_tmp_71 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_71
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_2 HxArray[int32]
								_ = tmp_v_2
								tmp_v_2 = this.Hx_Field_input
								var tmp_pos_2 int32
								_ = tmp_pos_2
								tmp_pos_2 = this.Hx_Field_pos
								var _hx_tmp_72 int32
								_ = _hx_tmp_72
								var _hx_tmp_73 bool
								_ = _hx_tmp_73
								if tmp_pos_2 < 0 {
									_hx_tmp_73 = true
								} else {
									var _hx_tmp_74 int32
									_ = _hx_tmp_74
									_hx_tmp_74 = tmp_pos_2
									var _hx_tmp_75 int32
									_ = _hx_tmp_75
									_hx_tmp_75 = _hx_tmp_74
									_hx_tmp_73 = (_hx_tmp_75 >= HxMakeArrayView[any](tmp_v_2).Len())
								}

								if _hx_tmp_73 {
									_hx_tmp_72 = -1
								} else {
									_hx_tmp_72 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
								}

								if (_hx_tmp_72) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						{
							var newPos int32
							_ = newPos
							newPos = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
							for this.Hx_Field_pos < newPos {
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_76 int32
									_ = _hx_tmp_76
									_hx_tmp_76 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_76
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_2 HxArray[int32]
									_ = tmp_v_2
									tmp_v_2 = this.Hx_Field_input
									var tmp_pos_2 int32
									_ = tmp_pos_2
									tmp_pos_2 = this.Hx_Field_pos
									var _hx_tmp_77 int32
									_ = _hx_tmp_77
									var _hx_tmp_78 bool
									_ = _hx_tmp_78
									if tmp_pos_2 < 0 {
										_hx_tmp_78 = true
									} else {
										var _hx_tmp_79 int32
										_ = _hx_tmp_79
										_hx_tmp_79 = tmp_pos_2
										var _hx_tmp_80 int32
										_ = _hx_tmp_80
										_hx_tmp_80 = _hx_tmp_79
										_hx_tmp_78 = (_hx_tmp_80 >= HxMakeArrayView[any](tmp_v_2).Len())
									}

									if _hx_tmp_78 {
										_hx_tmp_77 = -1
									} else {
										_hx_tmp_77 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
									}

									if (_hx_tmp_77) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}
						}
					} else {
						currentArg.VTable.Hx_Field_addChar(tmp_c_1)
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_81 int32
								_ = _hx_tmp_81
								_hx_tmp_81 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_81
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_2 HxArray[int32]
								_ = tmp_v_2
								tmp_v_2 = this.Hx_Field_input
								var tmp_pos_2 int32
								_ = tmp_pos_2
								tmp_pos_2 = this.Hx_Field_pos
								var _hx_tmp_82 int32
								_ = _hx_tmp_82
								var _hx_tmp_83 bool
								_ = _hx_tmp_83
								if tmp_pos_2 < 0 {
									_hx_tmp_83 = true
								} else {
									var _hx_tmp_84 int32
									_ = _hx_tmp_84
									_hx_tmp_84 = tmp_pos_2
									var _hx_tmp_85 int32
									_ = _hx_tmp_85
									_hx_tmp_85 = _hx_tmp_84
									_hx_tmp_83 = (_hx_tmp_85 >= HxMakeArrayView[any](tmp_v_2).Len())
								}

								if _hx_tmp_83 {
									_hx_tmp_82 = -1
								} else {
									_hx_tmp_82 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
								}

								if (_hx_tmp_82) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}
					}
				}
			}
		}
	}

	if parenLevel > 0 {
		this.VTable.Hx_Field_error("Unclosed parentheses in function declaration", false)
	}

	{
		var newPos2 int32
		_ = newPos2
		newPos2 = this.VTable.Hx_Field__skipWhitespaceAndComments(this.Hx_Field_pos, false)
		for this.Hx_Field_pos < newPos2 {
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_86 int32
				_ = _hx_tmp_86
				_hx_tmp_86 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_86
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_87 int32
				_ = _hx_tmp_87
				var _hx_tmp_88 bool
				_ = _hx_tmp_88
				if tmp_pos_1 < 0 {
					_hx_tmp_88 = true
				} else {
					var _hx_tmp_89 int32
					_ = _hx_tmp_89
					_hx_tmp_89 = tmp_pos_1
					var _hx_tmp_90 int32
					_ = _hx_tmp_90
					_hx_tmp_90 = _hx_tmp_89
					_hx_tmp_88 = (_hx_tmp_90 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_88 {
					_hx_tmp_87 = -1
				} else {
					_hx_tmp_87 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_87) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}
	}

	var usesBraces bool
	_ = usesBraces
	if this.Hx_Field_pos < this.Hx_Field_length {
		var tmp_v_1 HxArray[int32]
		_ = tmp_v_1
		tmp_v_1 = this.Hx_Field_input
		var tmp_pos_1 int32
		_ = tmp_pos_1
		tmp_pos_1 = this.Hx_Field_pos
		var _hx_tmp_91 int32
		_ = _hx_tmp_91
		var _hx_tmp_92 bool
		_ = _hx_tmp_92
		if tmp_pos_1 < 0 {
			_hx_tmp_92 = true
		} else {
			var _hx_tmp_93 int32
			_ = _hx_tmp_93
			_hx_tmp_93 = tmp_pos_1
			var _hx_tmp_94 int32
			_ = _hx_tmp_94
			_hx_tmp_94 = _hx_tmp_93
			_hx_tmp_92 = (_hx_tmp_94 >= HxMakeArrayView[any](tmp_v_1).Len())
		}

		if _hx_tmp_92 {
			_hx_tmp_91 = -1
		} else {
			_hx_tmp_91 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
		}

		usesBraces = ((_hx_tmp_91) == 123)
	} else {
		usesBraces = false
	}

	var lastLineBreakPos int32
	_ = lastLineBreakPos
	lastLineBreakPos = this.Hx_Field_pos
	var lastLineBreakLine int32
	_ = lastLineBreakLine
	lastLineBreakLine = this.Hx_Field_line
	var lastLineBreakColumn int32
	_ = lastLineBreakColumn
	lastLineBreakColumn = this.Hx_Field_column
	if usesBraces {
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = 1
			for _g < _g1 {
				var _hx_tmp_95 int32
				_ = _hx_tmp_95
				_hx_tmp_95 = _g
				_g = (_g + 1)
				var _hx_reserved__ int32
				_ = _hx_reserved__
				_hx_reserved__ = _hx_tmp_95
				if this.Hx_Field_pos >= this.Hx_Field_length {
					break
				}

				var tmp_v_1 HxArray[int32]
				_ = tmp_v_1
				tmp_v_1 = this.Hx_Field_input
				var tmp_pos_1 int32
				_ = tmp_pos_1
				tmp_pos_1 = this.Hx_Field_pos
				var _hx_tmp_96 int32
				_ = _hx_tmp_96
				var _hx_tmp_97 bool
				_ = _hx_tmp_97
				if tmp_pos_1 < 0 {
					_hx_tmp_97 = true
				} else {
					var _hx_tmp_98 int32
					_ = _hx_tmp_98
					_hx_tmp_98 = tmp_pos_1
					var _hx_tmp_99 int32
					_ = _hx_tmp_99
					_hx_tmp_99 = _hx_tmp_98
					_hx_tmp_97 = (_hx_tmp_99 >= HxMakeArrayView[any](tmp_v_1).Len())
				}

				if _hx_tmp_97 {
					_hx_tmp_96 = -1
				} else {
					_hx_tmp_96 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
				}

				if (_hx_tmp_96) == 10 {
					this.Hx_Field_line++
					this.Hx_Field_column = 1
				} else {
					this.Hx_Field_column++
				}

				this.Hx_Field_pos++
			}
		}

		var braceLevel int32
		_ = braceLevel
		braceLevel = 1
		for (this.Hx_Field_pos < this.Hx_Field_length) && (braceLevel > 0) {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
			var _hx_tmp_100 int32
			_ = _hx_tmp_100
			var _hx_tmp_101 bool
			_ = _hx_tmp_101
			if tmp_pos_1 < 0 {
				_hx_tmp_101 = true
			} else {
				var _hx_tmp_102 int32
				_ = _hx_tmp_102
				_hx_tmp_102 = tmp_pos_1
				var _hx_tmp_103 int32
				_ = _hx_tmp_103
				_hx_tmp_103 = _hx_tmp_102
				_hx_tmp_101 = (_hx_tmp_103 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_101 {
				_hx_tmp_100 = -1
			} else {
				_hx_tmp_100 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_100
			if tmp_c_1 == 34 {
				this.VTable.Hx_Field_skipQuotedString()
				continue
			}

			if tmp_c_1 == 123 {
				braceLevel++
			} else {
				if tmp_c_1 == 125 {
					braceLevel--
				}
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_104 int32
					_ = _hx_tmp_104
					_hx_tmp_104 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_104
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_105 int32
					_ = _hx_tmp_105
					var _hx_tmp_106 bool
					_ = _hx_tmp_106
					if tmp_pos_2 < 0 {
						_hx_tmp_106 = true
					} else {
						var _hx_tmp_107 int32
						_ = _hx_tmp_107
						_hx_tmp_107 = tmp_pos_2
						var _hx_tmp_108 int32
						_ = _hx_tmp_108
						_hx_tmp_108 = _hx_tmp_107
						_hx_tmp_106 = (_hx_tmp_108 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_106 {
						_hx_tmp_105 = -1
					} else {
						_hx_tmp_105 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_105) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}

		if braceLevel > 0 {
			this.VTable.Hx_Field_error("Unclosed braces in function body", false)
		}

		lastLineBreakPos = this.Hx_Field_pos
		lastLineBreakLine = this.Hx_Field_line
		lastLineBreakColumn = this.Hx_Field_column
	} else {
		var functionIndentLevel int32
		_ = functionIndentLevel
		functionIndentLevel = -1
		var currentLine bool
		_ = currentLine
		currentLine = true
		for this.Hx_Field_pos < this.Hx_Field_length {
			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
			var _hx_tmp_109 int32
			_ = _hx_tmp_109
			var _hx_tmp_110 bool
			_ = _hx_tmp_110
			if tmp_pos_1 < 0 {
				_hx_tmp_110 = true
			} else {
				var _hx_tmp_111 int32
				_ = _hx_tmp_111
				_hx_tmp_111 = tmp_pos_1
				var _hx_tmp_112 int32
				_ = _hx_tmp_112
				_hx_tmp_112 = _hx_tmp_111
				_hx_tmp_110 = (_hx_tmp_112 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_110 {
				_hx_tmp_109 = -1
			} else {
				_hx_tmp_109 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_109
			if (tmp_c_1 == 10) || (tmp_c_1 == 13) {
				lastLineBreakPos = this.Hx_Field_pos
				lastLineBreakLine = this.Hx_Field_line
				lastLineBreakColumn = this.Hx_Field_column
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_113 int32
						_ = _hx_tmp_113
						_hx_tmp_113 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_113
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_2 HxArray[int32]
						_ = tmp_v_2
						tmp_v_2 = this.Hx_Field_input
						var tmp_pos_2 int32
						_ = tmp_pos_2
						tmp_pos_2 = this.Hx_Field_pos
						var _hx_tmp_114 int32
						_ = _hx_tmp_114
						var _hx_tmp_115 bool
						_ = _hx_tmp_115
						if tmp_pos_2 < 0 {
							_hx_tmp_115 = true
						} else {
							var _hx_tmp_116 int32
							_ = _hx_tmp_116
							_hx_tmp_116 = tmp_pos_2
							var _hx_tmp_117 int32
							_ = _hx_tmp_117
							_hx_tmp_117 = _hx_tmp_116
							_hx_tmp_115 = (_hx_tmp_117 >= HxMakeArrayView[any](tmp_v_2).Len())
						}

						if _hx_tmp_115 {
							_hx_tmp_114 = -1
						} else {
							_hx_tmp_114 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
						}

						if (_hx_tmp_114) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}

				currentLine = true
				break
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_118 int32
					_ = _hx_tmp_118
					_hx_tmp_118 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_118
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_119 int32
					_ = _hx_tmp_119
					var _hx_tmp_120 bool
					_ = _hx_tmp_120
					if tmp_pos_2 < 0 {
						_hx_tmp_120 = true
					} else {
						var _hx_tmp_121 int32
						_ = _hx_tmp_121
						_hx_tmp_121 = tmp_pos_2
						var _hx_tmp_122 int32
						_ = _hx_tmp_122
						_hx_tmp_122 = _hx_tmp_121
						_hx_tmp_120 = (_hx_tmp_122 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_120 {
						_hx_tmp_119 = -1
					} else {
						_hx_tmp_119 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_119) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}

		for this.Hx_Field_pos < this.Hx_Field_length {
			if currentLine {
				var indent int32
				_ = indent
				indent = 0
				var indentStart int32
				_ = indentStart
				indentStart = this.Hx_Field_pos
				for this.Hx_Field_pos < this.Hx_Field_length {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_123 int32
					_ = _hx_tmp_123
					var _hx_tmp_124 bool
					_ = _hx_tmp_124
					if tmp_pos_1 < 0 {
						_hx_tmp_124 = true
					} else {
						var _hx_tmp_125 int32
						_ = _hx_tmp_125
						_hx_tmp_125 = tmp_pos_1
						var _hx_tmp_126 int32
						_ = _hx_tmp_126
						_hx_tmp_126 = _hx_tmp_125
						_hx_tmp_124 = (_hx_tmp_126 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_124 {
						_hx_tmp_123 = -1
					} else {
						_hx_tmp_123 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					var tmp_c_1 int32
					_ = tmp_c_1
					tmp_c_1 = _hx_tmp_123
					if tmp_c_1 == 32 {
						indent++
					} else {
						if tmp_c_1 == 9 {
							indent++
						} else {
							break
						}
					}

					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_127 int32
							_ = _hx_tmp_127
							_hx_tmp_127 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_127
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_2 HxArray[int32]
							_ = tmp_v_2
							tmp_v_2 = this.Hx_Field_input
							var tmp_pos_2 int32
							_ = tmp_pos_2
							tmp_pos_2 = this.Hx_Field_pos
							var _hx_tmp_128 int32
							_ = _hx_tmp_128
							var _hx_tmp_129 bool
							_ = _hx_tmp_129
							if tmp_pos_2 < 0 {
								_hx_tmp_129 = true
							} else {
								var _hx_tmp_130 int32
								_ = _hx_tmp_130
								_hx_tmp_130 = tmp_pos_2
								var _hx_tmp_131 int32
								_ = _hx_tmp_131
								_hx_tmp_131 = _hx_tmp_130
								_hx_tmp_129 = (_hx_tmp_131 >= HxMakeArrayView[any](tmp_v_2).Len())
							}

							if _hx_tmp_129 {
								_hx_tmp_128 = -1
							} else {
								_hx_tmp_128 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
							}

							if (_hx_tmp_128) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}
				}

				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				var tmp1 bool
				_ = tmp1
				if (functionIndentLevel == -1) && (this.Hx_Field_pos < this.Hx_Field_length) {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_132 int32
					_ = _hx_tmp_132
					var _hx_tmp_133 bool
					_ = _hx_tmp_133
					if tmp_pos_1 < 0 {
						_hx_tmp_133 = true
					} else {
						var _hx_tmp_134 int32
						_ = _hx_tmp_134
						_hx_tmp_134 = tmp_pos_1
						var _hx_tmp_135 int32
						_ = _hx_tmp_135
						_hx_tmp_135 = _hx_tmp_134
						_hx_tmp_133 = (_hx_tmp_135 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_133 {
						_hx_tmp_132 = -1
					} else {
						_hx_tmp_132 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					tmp1 = ((_hx_tmp_132) != 10)
				} else {
					tmp1 = false
				}

				if tmp1 {
					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_136 int32
					_ = _hx_tmp_136
					var _hx_tmp_137 bool
					_ = _hx_tmp_137
					if tmp_pos_1 < 0 {
						_hx_tmp_137 = true
					} else {
						var _hx_tmp_138 int32
						_ = _hx_tmp_138
						_hx_tmp_138 = tmp_pos_1
						var _hx_tmp_139 int32
						_ = _hx_tmp_139
						_hx_tmp_139 = _hx_tmp_138
						_hx_tmp_137 = (_hx_tmp_139 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_137 {
						_hx_tmp_136 = -1
					} else {
						_hx_tmp_136 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					tmp_tmp_1 = ((_hx_tmp_136) != 13)
				} else {
					tmp_tmp_1 = false
				}

				if tmp_tmp_1 {
					functionIndentLevel = indent
					if functionIndentLevel < minIndent {
						this.Hx_Field_pos = indentStart
						break
					}
				} else {
					var tmp_tmp_2 bool
					_ = tmp_tmp_2
					if (functionIndentLevel != -1) && (indent < functionIndentLevel) {
						if this.Hx_Field_pos < this.Hx_Field_length {
							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
							var _hx_tmp_140 int32
							_ = _hx_tmp_140
							var _hx_tmp_141 bool
							_ = _hx_tmp_141
							if tmp_pos_1 < 0 {
								_hx_tmp_141 = true
							} else {
								var _hx_tmp_142 int32
								_ = _hx_tmp_142
								_hx_tmp_142 = tmp_pos_1
								var _hx_tmp_143 int32
								_ = _hx_tmp_143
								_hx_tmp_143 = _hx_tmp_142
								_hx_tmp_141 = (_hx_tmp_143 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_141 {
								_hx_tmp_140 = -1
							} else {
								_hx_tmp_140 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							if (_hx_tmp_140) != 10 {
								var tmp_v_2 HxArray[int32]
								_ = tmp_v_2
								tmp_v_2 = this.Hx_Field_input
								var tmp_pos_2 int32
								_ = tmp_pos_2
								tmp_pos_2 = this.Hx_Field_pos
								var _hx_tmp_144 int32
								_ = _hx_tmp_144
								var _hx_tmp_145 bool
								_ = _hx_tmp_145
								if tmp_pos_2 < 0 {
									_hx_tmp_145 = true
								} else {
									var _hx_tmp_146 int32
									_ = _hx_tmp_146
									_hx_tmp_146 = tmp_pos_2
									var _hx_tmp_147 int32
									_ = _hx_tmp_147
									_hx_tmp_147 = _hx_tmp_146
									_hx_tmp_145 = (_hx_tmp_147 >= HxMakeArrayView[any](tmp_v_2).Len())
								}

								if _hx_tmp_145 {
									_hx_tmp_144 = -1
								} else {
									_hx_tmp_144 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
								}

								tmp_tmp_2 = ((_hx_tmp_144) != 13)
							} else {
								tmp_tmp_2 = false
							}
						} else {
							tmp_tmp_2 = true
						}
					} else {
						tmp_tmp_2 = false
					}

					if tmp_tmp_2 {
						this.Hx_Field_pos = indentStart
						break
					}
				}

				currentLine = false
			}

			var tmp_v_1 HxArray[int32]
			_ = tmp_v_1
			tmp_v_1 = this.Hx_Field_input
			var tmp_pos_1 int32
			_ = tmp_pos_1
			tmp_pos_1 = this.Hx_Field_pos
			var _hx_tmp_148 int32
			_ = _hx_tmp_148
			var _hx_tmp_149 bool
			_ = _hx_tmp_149
			if tmp_pos_1 < 0 {
				_hx_tmp_149 = true
			} else {
				var _hx_tmp_150 int32
				_ = _hx_tmp_150
				_hx_tmp_150 = tmp_pos_1
				var _hx_tmp_151 int32
				_ = _hx_tmp_151
				_hx_tmp_151 = _hx_tmp_150
				_hx_tmp_149 = (_hx_tmp_151 >= HxMakeArrayView[any](tmp_v_1).Len())
			}

			if _hx_tmp_149 {
				_hx_tmp_148 = -1
			} else {
				_hx_tmp_148 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
			}

			var tmp_c_1 int32
			_ = tmp_c_1
			tmp_c_1 = _hx_tmp_148
			if tmp_c_1 == 34 {
				this.VTable.Hx_Field_skipQuotedString()
				continue
			}

			if (tmp_c_1 == 10) || (tmp_c_1 == 13) {
				lastLineBreakPos = this.Hx_Field_pos
				lastLineBreakLine = this.Hx_Field_line
				lastLineBreakColumn = this.Hx_Field_column
				currentLine = true
			}

			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_152 int32
					_ = _hx_tmp_152
					_hx_tmp_152 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_152
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_2 HxArray[int32]
					_ = tmp_v_2
					tmp_v_2 = this.Hx_Field_input
					var tmp_pos_2 int32
					_ = tmp_pos_2
					tmp_pos_2 = this.Hx_Field_pos
					var _hx_tmp_153 int32
					_ = _hx_tmp_153
					var _hx_tmp_154 bool
					_ = _hx_tmp_154
					if tmp_pos_2 < 0 {
						_hx_tmp_154 = true
					} else {
						var _hx_tmp_155 int32
						_ = _hx_tmp_155
						_hx_tmp_155 = tmp_pos_2
						var _hx_tmp_156 int32
						_ = _hx_tmp_156
						_hx_tmp_156 = _hx_tmp_155
						_hx_tmp_154 = (_hx_tmp_156 >= HxMakeArrayView[any](tmp_v_2).Len())
					}

					if _hx_tmp_154 {
						_hx_tmp_153 = -1
					} else {
						_hx_tmp_153 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
					}

					if (_hx_tmp_153) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		}
	}

	this.Hx_Field_pos = lastLineBreakPos
	this.Hx_Field_line = lastLineBreakLine
	this.Hx_Field_column = lastLineBreakColumn
	var bodyEnd int32
	_ = bodyEnd
	bodyEnd = this.Hx_Field_pos
	var v1 HxArray[int32]
	_ = v1
	v1 = this.Hx_Field_input
	var pos1 int32
	_ = pos1
	pos1 = start.Hx_Field_offset
	var _hx_reserved_len HxNullable[int32]
	_ = _hx_reserved_len
	_hx_reserved_len = HxNullable[int32]{Value: (bodyEnd - start.Hx_Field_offset), Valid: true}

	var _hx_tmp_157 HxArray[int32]
	_ = _hx_tmp_157
	_hx_tmp_157 = v1
	var _hx_tmp_158 int32
	_ = _hx_tmp_158
	_hx_tmp_158 = pos1
	var _hx_tmp_159 int32
	_ = _hx_tmp_159
	if (_hx_reserved_len).Valid != false {
		var _hx_tmp_160 int32
		_ = _hx_tmp_160
		_hx_tmp_160 = pos1
		var _hx_tmp_161 int32
		_ = _hx_tmp_161
		_hx_tmp_161 = _hx_tmp_160
		_hx_tmp_159 = (_hx_tmp_161 + (_hx_reserved_len).Value)
	} else {
		_hx_tmp_159 = HxMakeArrayView[any](v1).Len()
	}

	var s string
	_ = s
	s = Hx_Field_loreline_utf8vector_slice(_hx_tmp_157, _hx_tmp_158, _hx_tmp_159)
	var l int32
	_ = l
	l = int32(utf8.RuneCountInString(s))
	var r int32
	_ = r
	r = 0
	for {
		var _hx_tmp_162 bool
		_ = _hx_tmp_162
		if r < l {
			_hx_tmp_162 = Hx_Field_stringtools_isSpace(s, ((l - r) - 1))
		} else {
			_hx_tmp_162 = false
		}

		if !(_hx_tmp_162) {
			break
		}

		r++
	}

	var _hx_tmp_163 string
	_ = _hx_tmp_163
	if r > 0 {
		var _hx_tmp_164 string
		_ = _hx_tmp_164
		_hx_tmp_164 = s
		var _hx_tmp_165 string
		_ = _hx_tmp_165
		_hx_tmp_165 = _hx_tmp_164
		_hx_tmp_163 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_165, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
	} else {
		_hx_tmp_163 = s
	}

	var code string
	_ = code
	code = _hx_tmp_163
	var _hx_tmp_166 bool
	_ = _hx_tmp_166
	if !usesBraces {
		var _hx_tmp_167 string
		_ = _hx_tmp_167
		_hx_tmp_167 = code
		var _hx_tmp_168 string
		_ = _hx_tmp_168
		_hx_tmp_168 = _hx_tmp_167
		_hx_tmp_166 = ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_168, "\n", HxNullable[int32]{})) == -1)
	} else {
		_hx_tmp_166 = false
	}

	var external bool
	_ = external
	external = _hx_tmp_166
	if !external {
		code = Hx_Field_go_haxe_hxstring_concat(code, "\n")
	}

	var position *Hx_Obj_loreline_position
	_ = position
	position = start
	if position == nil {
		position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
	}

	position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
	var token *Hx_Obj_loreline_token
	_ = token
	token = Hx_Obj_loreline_token_CreateInstance(Hx_Enum_loreline_tokentype_Function{name, args, code, external}, position)
	this.Hx_Field_previous = token
	var token1 *Hx_Obj_loreline_token
	_ = token1
	token1 = token
	token1.Hx_Field_pos.Hx_Field_length = int32(utf8.RuneCountInString(code))
	return token1
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_skipQuotedString() {
	var escaped bool
	_ = escaped
	escaped = false
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = 1
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}

	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if pos < 0 {
			_hx_tmp_6 = true
		} else {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = pos
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _hx_tmp_7
			_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](v).Len())
		}

		if _hx_tmp_6 {
			_hx_tmp_5 = -1
		} else {
			_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
		}

		var c int32
		_ = c
		c = _hx_tmp_5
		if escaped {
			escaped = false
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_9
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_10 int32
					_ = _hx_tmp_10
					var _hx_tmp_11 bool
					_ = _hx_tmp_11
					if tmp_pos_1 < 0 {
						_hx_tmp_11 = true
					} else {
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						_hx_tmp_12 = tmp_pos_1
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = _hx_tmp_12
						_hx_tmp_11 = (_hx_tmp_13 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_11 {
						_hx_tmp_10 = -1
					} else {
						_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_10) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}
		} else {
			if c == 92 {
				escaped = true
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = 1
					for _g < _g1 {
						var _hx_tmp_14 int32
						_ = _hx_tmp_14
						_hx_tmp_14 = _g
						_g = (_g + 1)
						var _hx_reserved__ int32
						_ = _hx_reserved__
						_hx_reserved__ = _hx_tmp_14
						if this.Hx_Field_pos >= this.Hx_Field_length {
							break
						}

						var tmp_v_1 HxArray[int32]
						_ = tmp_v_1
						tmp_v_1 = this.Hx_Field_input
						var tmp_pos_1 int32
						_ = tmp_pos_1
						tmp_pos_1 = this.Hx_Field_pos
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						var _hx_tmp_16 bool
						_ = _hx_tmp_16
						if tmp_pos_1 < 0 {
							_hx_tmp_16 = true
						} else {
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							_hx_tmp_17 = tmp_pos_1
							var _hx_tmp_18 int32
							_ = _hx_tmp_18
							_hx_tmp_18 = _hx_tmp_17
							_hx_tmp_16 = (_hx_tmp_18 >= HxMakeArrayView[any](tmp_v_1).Len())
						}

						if _hx_tmp_16 {
							_hx_tmp_15 = -1
						} else {
							_hx_tmp_15 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
						}

						if (_hx_tmp_15) == 10 {
							this.Hx_Field_line++
							this.Hx_Field_column = 1
						} else {
							this.Hx_Field_column++
						}

						this.Hx_Field_pos++
					}
				}
			} else {
				if c == 34 {
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_19 int32
							_ = _hx_tmp_19
							_hx_tmp_19 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_19
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
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
								_hx_tmp_21 = (_hx_tmp_23 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_21 {
								_hx_tmp_20 = -1
							} else {
								_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							if (_hx_tmp_20) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}

					break
				} else {
					if (c == 36) && ((this.Hx_Field_pos + 1) < this.Hx_Field_length) {
						{
							var _g int32
							_ = _g
							_g = 0
							var _g1 int32
							_ = _g1
							_g1 = 1
							for _g < _g1 {
								var _hx_tmp_24 int32
								_ = _hx_tmp_24
								_hx_tmp_24 = _g
								_g = (_g + 1)
								var _hx_reserved__ int32
								_ = _hx_reserved__
								_hx_reserved__ = _hx_tmp_24
								if this.Hx_Field_pos >= this.Hx_Field_length {
									break
								}

								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
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

								if (_hx_tmp_25) == 10 {
									this.Hx_Field_line++
									this.Hx_Field_column = 1
								} else {
									this.Hx_Field_column++
								}

								this.Hx_Field_pos++
							}
						}

						var tmp bool
						_ = tmp
						if this.Hx_Field_pos < this.Hx_Field_length {
							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
							var _hx_tmp_29 int32
							_ = _hx_tmp_29
							var _hx_tmp_30 bool
							_ = _hx_tmp_30
							if tmp_pos_1 < 0 {
								_hx_tmp_30 = true
							} else {
								var _hx_tmp_31 int32
								_ = _hx_tmp_31
								_hx_tmp_31 = tmp_pos_1
								var _hx_tmp_32 int32
								_ = _hx_tmp_32
								_hx_tmp_32 = _hx_tmp_31
								_hx_tmp_30 = (_hx_tmp_32 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_30 {
								_hx_tmp_29 = -1
							} else {
								_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							tmp = ((_hx_tmp_29) == 123)
						} else {
							tmp = false
						}

						if tmp {
							{
								var _g int32
								_ = _g
								_g = 0
								var _g1 int32
								_ = _g1
								_g1 = 1
								for _g < _g1 {
									var _hx_tmp_33 int32
									_ = _hx_tmp_33
									_hx_tmp_33 = _g
									_g = (_g + 1)
									var _hx_reserved__ int32
									_ = _hx_reserved__
									_hx_reserved__ = _hx_tmp_33
									if this.Hx_Field_pos >= this.Hx_Field_length {
										break
									}

									var tmp_v_1 HxArray[int32]
									_ = tmp_v_1
									tmp_v_1 = this.Hx_Field_input
									var tmp_pos_1 int32
									_ = tmp_pos_1
									tmp_pos_1 = this.Hx_Field_pos
									var _hx_tmp_34 int32
									_ = _hx_tmp_34
									var _hx_tmp_35 bool
									_ = _hx_tmp_35
									if tmp_pos_1 < 0 {
										_hx_tmp_35 = true
									} else {
										var _hx_tmp_36 int32
										_ = _hx_tmp_36
										_hx_tmp_36 = tmp_pos_1
										var _hx_tmp_37 int32
										_ = _hx_tmp_37
										_hx_tmp_37 = _hx_tmp_36
										_hx_tmp_35 = (_hx_tmp_37 >= HxMakeArrayView[any](tmp_v_1).Len())
									}

									if _hx_tmp_35 {
										_hx_tmp_34 = -1
									} else {
										_hx_tmp_34 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
									}

									if (_hx_tmp_34) == 10 {
										this.Hx_Field_line++
										this.Hx_Field_column = 1
									} else {
										this.Hx_Field_column++
									}

									this.Hx_Field_pos++
								}
							}

							var interpBraceLevel int32
							_ = interpBraceLevel
							interpBraceLevel = 1
							for (this.Hx_Field_pos < this.Hx_Field_length) && (interpBraceLevel > 0) {
								var tmp_v_1 HxArray[int32]
								_ = tmp_v_1
								tmp_v_1 = this.Hx_Field_input
								var tmp_pos_1 int32
								_ = tmp_pos_1
								tmp_pos_1 = this.Hx_Field_pos
								var _hx_tmp_38 int32
								_ = _hx_tmp_38
								var _hx_tmp_39 bool
								_ = _hx_tmp_39
								if tmp_pos_1 < 0 {
									_hx_tmp_39 = true
								} else {
									var _hx_tmp_40 int32
									_ = _hx_tmp_40
									_hx_tmp_40 = tmp_pos_1
									var _hx_tmp_41 int32
									_ = _hx_tmp_41
									_hx_tmp_41 = _hx_tmp_40
									_hx_tmp_39 = (_hx_tmp_41 >= HxMakeArrayView[any](tmp_v_1).Len())
								}

								if _hx_tmp_39 {
									_hx_tmp_38 = -1
								} else {
									_hx_tmp_38 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
								}

								var ic int32
								_ = ic
								ic = _hx_tmp_38
								if ic == 34 {
									this.VTable.Hx_Field_skipQuotedString()
									continue
								} else {
									if ic == 123 {
										interpBraceLevel++
									} else {
										if ic == 125 {
											interpBraceLevel--
										}
									}
								}

								if (interpBraceLevel > 0) || (ic != 125) {
									var _g int32
									_ = _g
									_g = 0
									var _g1 int32
									_ = _g1
									_g1 = 1
									for _g < _g1 {
										var _hx_tmp_42 int32
										_ = _hx_tmp_42
										_hx_tmp_42 = _g
										_g = (_g + 1)
										var _hx_reserved__ int32
										_ = _hx_reserved__
										_hx_reserved__ = _hx_tmp_42
										if this.Hx_Field_pos >= this.Hx_Field_length {
											break
										}

										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = this.Hx_Field_input
										var tmp_pos_2 int32
										_ = tmp_pos_2
										tmp_pos_2 = this.Hx_Field_pos
										var _hx_tmp_43 int32
										_ = _hx_tmp_43
										var _hx_tmp_44 bool
										_ = _hx_tmp_44
										if tmp_pos_2 < 0 {
											_hx_tmp_44 = true
										} else {
											var _hx_tmp_45 int32
											_ = _hx_tmp_45
											_hx_tmp_45 = tmp_pos_2
											var _hx_tmp_46 int32
											_ = _hx_tmp_46
											_hx_tmp_46 = _hx_tmp_45
											_hx_tmp_44 = (_hx_tmp_46 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_44 {
											_hx_tmp_43 = -1
										} else {
											_hx_tmp_43 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
										}

										if (_hx_tmp_43) == 10 {
											this.Hx_Field_line++
											this.Hx_Field_column = 1
										} else {
											this.Hx_Field_column++
										}

										this.Hx_Field_pos++
									}
								} else {
									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_47 int32
											_ = _hx_tmp_47
											_hx_tmp_47 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_47
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_2 HxArray[int32]
											_ = tmp_v_2
											tmp_v_2 = this.Hx_Field_input
											var tmp_pos_2 int32
											_ = tmp_pos_2
											tmp_pos_2 = this.Hx_Field_pos
											var _hx_tmp_48 int32
											_ = _hx_tmp_48
											var _hx_tmp_49 bool
											_ = _hx_tmp_49
											if tmp_pos_2 < 0 {
												_hx_tmp_49 = true
											} else {
												var _hx_tmp_50 int32
												_ = _hx_tmp_50
												_hx_tmp_50 = tmp_pos_2
												var _hx_tmp_51 int32
												_ = _hx_tmp_51
												_hx_tmp_51 = _hx_tmp_50
												_hx_tmp_49 = (_hx_tmp_51 >= HxMakeArrayView[any](tmp_v_2).Len())
											}

											if _hx_tmp_49 {
												_hx_tmp_48 = -1
											} else {
												_hx_tmp_48 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
											}

											if (_hx_tmp_48) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}

									break
								}
							}
						} else {
							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
							var _hx_tmp_52 int32
							_ = _hx_tmp_52
							var _hx_tmp_53 bool
							_ = _hx_tmp_53
							if tmp_pos_1 < 0 {
								_hx_tmp_53 = true
							} else {
								var _hx_tmp_54 int32
								_ = _hx_tmp_54
								_hx_tmp_54 = tmp_pos_1
								var _hx_tmp_55 int32
								_ = _hx_tmp_55
								_hx_tmp_55 = _hx_tmp_54
								_hx_tmp_53 = (_hx_tmp_55 >= HxMakeArrayView[any](tmp_v_1).Len())
							}

							if _hx_tmp_53 {
								_hx_tmp_52 = -1
							} else {
								_hx_tmp_52 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
							}

							var tmp_c_1 int32
							_ = tmp_c_1
							tmp_c_1 = _hx_tmp_52
							if (((tmp_c_1 >= 97) && (tmp_c_1 <= 122)) || ((tmp_c_1 >= 65) && (tmp_c_1 <= 90))) || (tmp_c_1 == 95) {
								for {
									var tmp_tmp_1 bool
									_ = tmp_tmp_1
									if this.Hx_Field_pos < this.Hx_Field_length {
										var tmp_v_2 HxArray[int32]
										_ = tmp_v_2
										tmp_v_2 = this.Hx_Field_input
										var tmp_pos_2 int32
										_ = tmp_pos_2
										tmp_pos_2 = this.Hx_Field_pos
										var _hx_tmp_56 int32
										_ = _hx_tmp_56
										var _hx_tmp_57 bool
										_ = _hx_tmp_57
										if tmp_pos_2 < 0 {
											_hx_tmp_57 = true
										} else {
											var _hx_tmp_58 int32
											_ = _hx_tmp_58
											_hx_tmp_58 = tmp_pos_2
											var _hx_tmp_59 int32
											_ = _hx_tmp_59
											_hx_tmp_59 = _hx_tmp_58
											_hx_tmp_57 = (_hx_tmp_59 >= HxMakeArrayView[any](tmp_v_2).Len())
										}

										if _hx_tmp_57 {
											_hx_tmp_56 = -1
										} else {
											_hx_tmp_56 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
										}

										var tmp_c_2 int32
										_ = tmp_c_2
										tmp_c_2 = _hx_tmp_56
										tmp_tmp_1 = (((((tmp_c_2 >= 97) && (tmp_c_2 <= 122)) || ((tmp_c_2 >= 65) && (tmp_c_2 <= 90))) || (tmp_c_2 == 95)) || ((tmp_c_2 >= 48) && (tmp_c_2 <= 57)))
									} else {
										tmp_tmp_1 = false
									}

									if !tmp_tmp_1 {
										break
									}

									{
										var _g int32
										_ = _g
										_g = 0
										var _g1 int32
										_ = _g1
										_g1 = 1
										for _g < _g1 {
											var _hx_tmp_60 int32
											_ = _hx_tmp_60
											_hx_tmp_60 = _g
											_g = (_g + 1)
											var _hx_reserved__ int32
											_ = _hx_reserved__
											_hx_reserved__ = _hx_tmp_60
											if this.Hx_Field_pos >= this.Hx_Field_length {
												break
											}

											var tmp_v_2 HxArray[int32]
											_ = tmp_v_2
											tmp_v_2 = this.Hx_Field_input
											var tmp_pos_2 int32
											_ = tmp_pos_2
											tmp_pos_2 = this.Hx_Field_pos
											var _hx_tmp_61 int32
											_ = _hx_tmp_61
											var _hx_tmp_62 bool
											_ = _hx_tmp_62
											if tmp_pos_2 < 0 {
												_hx_tmp_62 = true
											} else {
												var _hx_tmp_63 int32
												_ = _hx_tmp_63
												_hx_tmp_63 = tmp_pos_2
												var _hx_tmp_64 int32
												_ = _hx_tmp_64
												_hx_tmp_64 = _hx_tmp_63
												_hx_tmp_62 = (_hx_tmp_64 >= HxMakeArrayView[any](tmp_v_2).Len())
											}

											if _hx_tmp_62 {
												_hx_tmp_61 = -1
											} else {
												_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_2).Get(tmp_pos_2))))
											}

											if (_hx_tmp_61) == 10 {
												this.Hx_Field_line++
												this.Hx_Field_column = 1
											} else {
												this.Hx_Field_column++
											}

											this.Hx_Field_pos++
										}
									}
								}
							}
						}
					} else {
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = 1
						for _g < _g1 {
							var _hx_tmp_65 int32
							_ = _hx_tmp_65
							_hx_tmp_65 = _g
							_g = (_g + 1)
							var _hx_reserved__ int32
							_ = _hx_reserved__
							_hx_reserved__ = _hx_tmp_65
							if this.Hx_Field_pos >= this.Hx_Field_length {
								break
							}

							var tmp_v_1 HxArray[int32]
							_ = tmp_v_1
							tmp_v_1 = this.Hx_Field_input
							var tmp_pos_1 int32
							_ = tmp_pos_1
							tmp_pos_1 = this.Hx_Field_pos
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

							if (_hx_tmp_66) == 10 {
								this.Hx_Field_line++
								this.Hx_Field_column = 1
							} else {
								this.Hx_Field_column++
							}

							this.Hx_Field_pos++
						}
					}
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_makePosition() *Hx_Obj_loreline_position {
	return Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_makeToken(_hx_reserved_type Hx_Enum_loreline_tokentype, position__isOptional any) *Hx_Obj_loreline_token {
	position := HxDefault[*Hx_Obj_loreline_position]()
	_ = position

	if position__isOptional != nil {
		position = HxConvert[*Hx_Obj_loreline_position](position__isOptional)
	} else {
		position = nil
	}
	{
		if position == nil {
			position = Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil)
		}

		position.Hx_Field_length = (this.Hx_Field_pos - position.Hx_Field_offset)
		var token *Hx_Obj_loreline_token
		_ = token
		token = Hx_Obj_loreline_token_CreateInstance(_hx_reserved_type, position)
		this.Hx_Field_previous = token
		return token
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_advance(count__isOptional any) {
	count := HxDefault[int32]()
	_ = count

	if count__isOptional != nil {
		count = HxConvert[int32](count__isOptional)
	} else {
		count = 1
	}
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = count
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			if this.Hx_Field_pos >= this.Hx_Field_length {
				break
			}

			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = this.Hx_Field_pos
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if pos < 0 {
				_hx_tmp_2 = true
			} else {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_2 {
				_hx_tmp_1 = -1
			} else {
				_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}

			if (_hx_tmp_1) == 10 {
				this.Hx_Field_line++
				this.Hx_Field_column = 1
			} else {
				this.Hx_Field_column++
			}

			this.Hx_Field_pos++
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_peek(offset__isOptional any) int32 {
	offset := HxDefault[int32]()
	_ = offset

	if offset__isOptional != nil {
		offset = HxConvert[int32](offset__isOptional)
	} else {
		offset = 1
	}
	{
		if (this.Hx_Field_pos + offset) < this.Hx_Field_length {
			var v HxArray[int32]
			_ = v
			v = this.Hx_Field_input
			var pos int32
			_ = pos
			pos = (this.Hx_Field_pos + offset)
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if pos < 0 {
				_hx_tmp_0 = true
			} else {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = pos
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _hx_tmp_1
				_hx_tmp_0 = (_hx_tmp_2 >= HxMakeArrayView[any](v).Len())
			}

			if _hx_tmp_0 {
				return -1
			} else {
				return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](v).Get(pos))))
			}
		} else {
			return 0
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_error(message string, fatal bool) *Hx_Obj_loreline_lexererror {
	var err *Hx_Obj_loreline_lexererror
	_ = err
	err = Hx_Obj_loreline_lexererror_CreateInstance(message, Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_startLine, this.Hx_Field_startColumn, this.Hx_Field_pos, nil))
	if this.Hx_Field_errors == nil {
		this.Hx_Field_errors = HxMakeArray[*Hx_Obj_loreline_lexererror]()
	}

	_ = (Hx_Array_Push[*Hx_Obj_loreline_lexererror](this.Hx_Field_errors, err))
	if fatal {
		panic(Hx_Field_haxe_exception_thrown(((any)(err))))
	}

	return err
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_getErrors() HxArray[*Hx_Obj_loreline_lexererror] {
	if this.Hx_Field_errors == nil {
		this.Hx_Field_errors = HxMakeArray[*Hx_Obj_loreline_lexererror]()
	}

	return this.Hx_Field_errors
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_skipWhitespace() {
	for this.Hx_Field_pos < this.Hx_Field_length {
		var v HxArray[int32]
		_ = v
		v = this.Hx_Field_input
		var pos int32
		_ = pos
		pos = this.Hx_Field_pos
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

		switch _hx_tmp_0 {
		case 9, 32:
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 int32
				_ = _g1
				_g1 = 1
				for _g < _g1 {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g
					_g = (_g + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_4
					if this.Hx_Field_pos >= this.Hx_Field_length {
						break
					}

					var tmp_v_1 HxArray[int32]
					_ = tmp_v_1
					tmp_v_1 = this.Hx_Field_input
					var tmp_pos_1 int32
					_ = tmp_pos_1
					tmp_pos_1 = this.Hx_Field_pos
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					var _hx_tmp_6 bool
					_ = _hx_tmp_6
					if tmp_pos_1 < 0 {
						_hx_tmp_6 = true
					} else {
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = tmp_pos_1
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						_hx_tmp_8 = _hx_tmp_7
						_hx_tmp_6 = (_hx_tmp_8 >= HxMakeArrayView[any](tmp_v_1).Len())
					}

					if _hx_tmp_6 {
						_hx_tmp_5 = -1
					} else {
						_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](tmp_v_1).Get(tmp_pos_1))))
					}

					if (_hx_tmp_5) == 10 {
						this.Hx_Field_line++
						this.Hx_Field_column = 1
					} else {
						this.Hx_Field_column++
					}

					this.Hx_Field_pos++
				}
			}

		default:
			{
				return
			}
		}
	}
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isDigit(c int32) bool {
	return ((c >= 48) && (c <= 57))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isIdentifierStart(c int32) bool {
	return ((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isLowerCase(c int32) bool {
	return ((c >= 97) && (c <= 122))
}

func (this *Hx_Obj_loreline_lexer) Hx_Field_isIdentifierPart(c int32) bool {
	return (((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || (c == 95)) || ((c >= 48) && (c <= 57)))
}

var Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars string

func Hx_Init_Hx_Field_loreline_lexer_hasNonSpecialChar_specialChars() string {
	return " \t\"+=*/-(){}[]:\n\r"
}

var Hx_Field_loreline_lexer_KEYWORDS *Hx_Obj_haxe_imap

func Hx_Init_Hx_Field_loreline_lexer_KEYWORDS() *Hx_Obj_haxe_imap {
	var _g *Hx_Obj_haxe_imap
	_ = _g
	_g = (func() *Hx_Obj_haxe_imap {
		_hx_icast_28 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_28 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_28, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_28.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_153 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_153
		_hx_closure_src_153 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_153(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("import", Hx_Enum_loreline_tokentype_KwImport{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_154 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_154
		_hx_closure_src_154 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_154(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("state", Hx_Enum_loreline_tokentype_KwState{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_155 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_155
		_hx_closure_src_155 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_155(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("beat", Hx_Enum_loreline_tokentype_KwBeat{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_156 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_156
		_hx_closure_src_156 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_156(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("character", Hx_Enum_loreline_tokentype_KwCharacter{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_157 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_157
		_hx_closure_src_157 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_157(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("choice", Hx_Enum_loreline_tokentype_KwChoice{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_158 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_158
		_hx_closure_src_158 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_158(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("if", Hx_Enum_loreline_tokentype_KwIf{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_159 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_159
		_hx_closure_src_159 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_159(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("else", Hx_Enum_loreline_tokentype_KwElse{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_160 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_160
		_hx_closure_src_160 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_160(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("new", Hx_Enum_loreline_tokentype_KwNew{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_161 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_161
		_hx_closure_src_161 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_161(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("sequence", Hx_Enum_loreline_tokentype_KwSequence{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_162 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_162
		_hx_closure_src_162 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_162(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("cycle", Hx_Enum_loreline_tokentype_KwCycle{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_163 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_163
		_hx_closure_src_163 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_163(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("once", Hx_Enum_loreline_tokentype_KwOnce{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_164 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_164
		_hx_closure_src_164 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_164(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("pick", Hx_Enum_loreline_tokentype_KwPick{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_165 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_165
		_hx_closure_src_165 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_165(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("shuffle", Hx_Enum_loreline_tokentype_KwShuffle{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_166 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_166
		_hx_closure_src_166 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_166(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("true", Hx_Enum_loreline_tokentype_LBoolean{true})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_167 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_167
		_hx_closure_src_167 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_167(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("false", Hx_Enum_loreline_tokentype_LBoolean{false})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_168 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_168
		_hx_closure_src_168 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_168(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("null", Hx_Enum_loreline_tokentype_LNull{})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_169 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_169
		_hx_closure_src_169 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_169(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("and", Hx_Enum_loreline_tokentype_OpAnd{true})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_170 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_170
		_hx_closure_src_170 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_170(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("or", Hx_Enum_loreline_tokentype_OpOr{true})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_171 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_171
		_hx_closure_src_171 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_171(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("is", Hx_Enum_loreline_tokentype_OpEquals{true})
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
		var _hx_closure_src_172 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_172
		_hx_closure_src_172 = _g.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_loreline_tokentype) {
			_hx_closure_src_172(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("not", Hx_Enum_loreline_tokentype_OpNot{true})
	var hx_dyncast_27 any
	_ = hx_dyncast_27
	hx_dyncast_27 = Hx_Field_go_haxe_hxdynamic_castClass(_g, "Hx_Obj_haxe_imap")
	var _hx_tmp_0 *Hx_Obj_haxe_imap
	_ = _hx_tmp_0
	if hx_dyncast_27 != nil {
		_hx_tmp_0 = hx_dyncast_27.(*Hx_Obj_haxe_imap)
	} else {
		_hx_tmp_0 = nil
	}

	return _hx_tmp_0
}

func Hx_Field_loreline_lexer_compareAttachments(l Hx_Enum_loreline_lstringattachment, r Hx_Enum_loreline_lstringattachment) int32 {
	var lStart int32
	_ = lStart
	switch l.Hx_Field_enumIndex() {
	case 0:
		{
			var start int32
			_ = start
			start = l.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
			lStart = start
		}

	case 1:
		{
			var start int32
			_ = start
			start = l.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_start
			lStart = start
		}

	}

	var rStart int32
	_ = rStart
	switch r.Hx_Field_enumIndex() {
	case 0:
		{
			var start int32
			_ = start
			start = r.(Hx_Enum_loreline_lstringattachment_Interpolation).Hx_Field_start
			rStart = start
		}

	case 1:
		{
			var start int32
			_ = start
			start = r.(Hx_Enum_loreline_lstringattachment_Tag).Hx_Field_start
			rStart = start
		}

	}

	return (lStart - rStart)
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_lexererror]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_token]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lstringattachment]()
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
