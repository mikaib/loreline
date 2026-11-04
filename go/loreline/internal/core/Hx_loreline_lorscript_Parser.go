package core

import "unicode/utf8"

type Hx_Enum_loreline_lorscript_token interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_lorscript_token()
}

type Hx_Enum_loreline_lorscript_token_TSemicolon struct{}

func (this Hx_Enum_loreline_lorscript_token_TSemicolon) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TSemicolon) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TSemicolon) Hx_Field_enumIndex() int32           { return 11 }
func (this Hx_Enum_loreline_lorscript_token_TSemicolon) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TSemicolon) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TSemicolon) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TQuestionDot struct{}

func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) Hx_Field_enumIndex() int32           { return 9 }
func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TQuestionDot) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TQuestion struct{}

func (this Hx_Enum_loreline_lorscript_token_TQuestion) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TQuestion) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TQuestion) Hx_Field_enumIndex() int32           { return 14 }
func (this Hx_Enum_loreline_lorscript_token_TQuestion) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TQuestion) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TQuestion) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TPrepro struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_token_TPrepro) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TPrepro) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TPrepro) Hx_Field_enumIndex() int32           { return 17 }
func (this Hx_Enum_loreline_lorscript_token_TPrepro) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TPrepro) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_token_TPrepro) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TPOpen struct{}

func (this Hx_Enum_loreline_lorscript_token_TPOpen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TPOpen) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TPOpen) Hx_Field_enumIndex() int32           { return 4 }
func (this Hx_Enum_loreline_lorscript_token_TPOpen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TPOpen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TPOpen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TPClose struct{}

func (this Hx_Enum_loreline_lorscript_token_TPClose) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TPClose) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TPClose) Hx_Field_enumIndex() int32           { return 5 }
func (this Hx_Enum_loreline_lorscript_token_TPClose) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TPClose) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TPClose) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TOp struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_token_TOp) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_lorscript_token_TOp) M_Hx_Enum_loreline_lorscript_token()     {}
func (this Hx_Enum_loreline_lorscript_token_TOp) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_loreline_lorscript_token_TOp) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TOp) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_token_TOp) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TMeta struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_token_TMeta) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TMeta) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TMeta) Hx_Field_enumIndex() int32           { return 16 }
func (this Hx_Enum_loreline_lorscript_token_TMeta) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TMeta) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_token_TMeta) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TId struct {
	Hx_Field_s string
}

func (this Hx_Enum_loreline_lorscript_token_TId) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_lorscript_token_TId) M_Hx_Enum_loreline_lorscript_token()     {}
func (this Hx_Enum_loreline_lorscript_token_TId) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_loreline_lorscript_token_TId) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TId) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_s))
}
func (this Hx_Enum_loreline_lorscript_token_TId) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_s)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TEof struct{}

func (this Hx_Enum_loreline_lorscript_token_TEof) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TEof) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TEof) Hx_Field_enumIndex() int32           { return 0 }
func (this Hx_Enum_loreline_lorscript_token_TEof) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TEof) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TEof) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TDoubleDot struct{}

func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) Hx_Field_enumIndex() int32           { return 15 }
func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TDoubleDot) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TDot struct{}

func (this Hx_Enum_loreline_lorscript_token_TDot) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TDot) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TDot) Hx_Field_enumIndex() int32           { return 8 }
func (this Hx_Enum_loreline_lorscript_token_TDot) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TDot) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TDot) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TConst struct {
	Hx_Field_c Hx_Enum_loreline_lorscript_const
}

func (this Hx_Enum_loreline_lorscript_token_TConst) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TConst) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TConst) Hx_Field_enumIndex() int32           { return 1 }
func (this Hx_Enum_loreline_lorscript_token_TConst) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TConst) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_c))
}
func (this Hx_Enum_loreline_lorscript_token_TConst) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_c)
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TComma struct{}

func (this Hx_Enum_loreline_lorscript_token_TComma) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TComma) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TComma) Hx_Field_enumIndex() int32           { return 10 }
func (this Hx_Enum_loreline_lorscript_token_TComma) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TComma) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TComma) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TBrOpen struct{}

func (this Hx_Enum_loreline_lorscript_token_TBrOpen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TBrOpen) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TBrOpen) Hx_Field_enumIndex() int32           { return 6 }
func (this Hx_Enum_loreline_lorscript_token_TBrOpen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TBrOpen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TBrOpen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TBrClose struct{}

func (this Hx_Enum_loreline_lorscript_token_TBrClose) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TBrClose) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TBrClose) Hx_Field_enumIndex() int32           { return 7 }
func (this Hx_Enum_loreline_lorscript_token_TBrClose) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TBrClose) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TBrClose) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TBkOpen struct{}

func (this Hx_Enum_loreline_lorscript_token_TBkOpen) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TBkOpen) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TBkOpen) Hx_Field_enumIndex() int32           { return 12 }
func (this Hx_Enum_loreline_lorscript_token_TBkOpen) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TBkOpen) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TBkOpen) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_lorscript_token_TBkClose struct{}

func (this Hx_Enum_loreline_lorscript_token_TBkClose) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_lorscript_token_TBkClose) M_Hx_Enum_loreline_lorscript_token() {}
func (this Hx_Enum_loreline_lorscript_token_TBkClose) Hx_Field_enumIndex() int32           { return 13 }
func (this Hx_Enum_loreline_lorscript_token_TBkClose) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_lorscript_token_RTTI
}
func (this Hx_Enum_loreline_lorscript_token_TBkClose) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_lorscript_token_TBkClose) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_lorscript_token_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.lorscript.Token",
	HxMakeArray[string]("TEof", "TConst", "TId", "TOp", "TPOpen", "TPClose", "TBrOpen", "TBrClose", "TDot", "TQuestionDot", "TComma", "TSemicolon", "TBkOpen", "TBkClose", "TQuestion", "TDoubleDot", "TMeta", "TPrepro"),
	HxMakeArray[int32](0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_lorscript_token_TEof{}
		case 1:
			return Hx_Enum_loreline_lorscript_token_TConst{params.Get_Dyn(0).(Hx_Enum_loreline_lorscript_const)}
		case 2:
			return Hx_Enum_loreline_lorscript_token_TId{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 3:
			return Hx_Enum_loreline_lorscript_token_TOp{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 4:
			return Hx_Enum_loreline_lorscript_token_TPOpen{}
		case 5:
			return Hx_Enum_loreline_lorscript_token_TPClose{}
		case 6:
			return Hx_Enum_loreline_lorscript_token_TBrOpen{}
		case 7:
			return Hx_Enum_loreline_lorscript_token_TBrClose{}
		case 8:
			return Hx_Enum_loreline_lorscript_token_TDot{}
		case 9:
			return Hx_Enum_loreline_lorscript_token_TQuestionDot{}
		case 10:
			return Hx_Enum_loreline_lorscript_token_TComma{}
		case 11:
			return Hx_Enum_loreline_lorscript_token_TSemicolon{}
		case 12:
			return Hx_Enum_loreline_lorscript_token_TBkOpen{}
		case 13:
			return Hx_Enum_loreline_lorscript_token_TBkClose{}
		case 14:
			return Hx_Enum_loreline_lorscript_token_TQuestion{}
		case 15:
			return Hx_Enum_loreline_lorscript_token_TDoubleDot{}
		case 16:
			return Hx_Enum_loreline_lorscript_token_TMeta{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 17:
			return Hx_Enum_loreline_lorscript_token_TPrepro{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_lorscript_parser() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lorscript_parser() {
}

type Hx_Obj_VTable_loreline_lorscript_parser interface {
	Hx_Field_error(err Hx_Enum_loreline_lorscript_errordef, pmin int32, pmax int32)
	Hx_Field_pmin(e any) int32
	Hx_Field_parseLambda(args HxArray[any], pmin HxNullable[int32]) any
	Hx_Field_parseMetaArgs() HxArray[any]
	Hx_Field_pmax(e any) int32
	Hx_Field_parseType() Hx_Enum_loreline_lorscript_ctype
	Hx_Field_readChar() int32
	Hx_Field_readString(until int32) string
	Hx_Field_initParser(origin string, pos int32)
	Hx_Field_maybe(tk Hx_Enum_loreline_lorscript_token) bool
	Hx_Field_getIdent() string
	Hx_Field_makeUnop(op string, e any) any
	Hx_Field_parseTypeNext(t Hx_Enum_loreline_lorscript_ctype) Hx_Enum_loreline_lorscript_ctype
	Hx_Field_parseParams() any
	Hx_Field_token() Hx_Enum_loreline_lorscript_token
	Hx_Field_skipTokens()
	Hx_Field_unexpected(tk Hx_Enum_loreline_lorscript_token) any
	Hx_Field_ensureToken(tk Hx_Enum_loreline_lorscript_token)
	Hx_Field_parseExpr() any
	Hx_Field_parseModule(content string, origin__isOptional any, position__isOptional any) HxArray[Hx_Enum_loreline_lorscript_moduledecl]
	Hx_Field__token() Hx_Enum_loreline_lorscript_token
	Hx_Field_parsePreproCond() any
	Hx_Field_invalidChar(c int32)
	Hx_Field_parseStructure(id string) any
	Hx_Field_parsePath() HxArray[string]
	Hx_Field_parseField() any
	Hx_Field_evalPreproCond(e any) bool
	Hx_Field_get_currentPos() int32
	Hx_Field_parseString(s string, origin__isOptional any, position__isOptional any) any
	Hx_Field_push(tk Hx_Enum_loreline_lorscript_token)
	Hx_Field_expr(e any) Hx_Enum_loreline_lorscript_exprdef
	Hx_Field_makeBinop(op string, e1 any, e any) any
	Hx_Field_parseExprNext(e1 any) any
	Hx_Field_parseFunctionArgs() HxArray[any]
	Hx_Field_parseExprList(etk Hx_Enum_loreline_lorscript_token) HxArray[any]
	Hx_Field_mk(e Hx_Enum_loreline_lorscript_exprdef, pmin__isOptional any, pmax__isOptional any) any
	Hx_Field_parseFullExpr(exprs HxArray[any])
	Hx_Field_parseObject(p1 HxNullable[int32]) any
	Hx_Field_mapCompr(tmp string, e any) any
	Hx_Field_parseFunctionDecl() any
	Hx_Field_parseMetadata() HxArray[any]
	Hx_Field_parseModuleDecl() Hx_Enum_loreline_lorscript_moduledecl
	Hx_Field_preprocValue(id string) any
	Hx_Field_ensure(tk Hx_Enum_loreline_lorscript_token)
	Hx_Field_isBlock(e any) bool
	Hx_Field_preprocess(id string) Hx_Enum_loreline_lorscript_token
	Hx_Field_tokenComment(op string, char int32) Hx_Enum_loreline_lorscript_token
	Hx_Field_constString(c Hx_Enum_loreline_lorscript_const) string
	Hx_Field_tokenString(t Hx_Enum_loreline_lorscript_token) string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lorscript_parser struct {
	VTable                     Hx_Obj_VTable_loreline_lorscript_parser `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_line              int32
	Hx_Field_opChars           string
	Hx_Field_identChars        string
	Hx_Field_opPriority        Hx_Typedef_map
	Hx_Field_opRightAssoc      Hx_Typedef_map
	Hx_Field_preprocesorValues Hx_Typedef_map
	Hx_Field_allowJSON         bool
	Hx_Field_allowTypes        bool
	Hx_Field_allowMetadata     bool
	Hx_Field_resumeErrors      bool
	Hx_Field_input             string
	Hx_Field_readPos           int32
	Hx_Field_offset            int32
	Hx_Field_char              int32
	Hx_Field_ops               HxArray[bool]
	Hx_Field_idents            HxArray[bool]
	Hx_Field_uid               int32
	Hx_Field_origin            string
	Hx_Field_tokenMin          int32
	Hx_Field_tokenMax          int32
	Hx_Field_oldTokenMin       int32
	Hx_Field_oldTokenMax       int32
	Hx_Field_tokens            Hx_Typedef_list
	Hx_Field_preprocStack      HxArray[any]
}

func Hx_Obj_loreline_lorscript_parser_CreateEmptyInstance() *Hx_Obj_loreline_lorscript_parser {
	obj := &Hx_Obj_loreline_lorscript_parser{}
	obj.VTable = obj
	obj.Hx_Field_line = HxDefault[int32]()
	obj.Hx_Field_opChars = HxDefault[string]()
	obj.Hx_Field_identChars = HxDefault[string]()
	obj.Hx_Field_opPriority = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_opRightAssoc = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_preprocesorValues = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_allowJSON = HxDefault[bool]()
	obj.Hx_Field_allowTypes = HxDefault[bool]()
	obj.Hx_Field_allowMetadata = HxDefault[bool]()
	obj.Hx_Field_resumeErrors = HxDefault[bool]()
	obj.Hx_Field_input = HxDefault[string]()
	obj.Hx_Field_readPos = HxDefault[int32]()
	obj.Hx_Field_offset = HxDefault[int32]()
	obj.Hx_Field_char = HxDefault[int32]()
	obj.Hx_Field_ops = HxDefault[HxArray[bool]]()
	obj.Hx_Field_idents = HxDefault[HxArray[bool]]()
	obj.Hx_Field_uid = HxDefault[int32]()
	obj.Hx_Field_origin = HxDefault[string]()
	obj.Hx_Field_tokenMin = HxDefault[int32]()
	obj.Hx_Field_tokenMax = HxDefault[int32]()
	obj.Hx_Field_oldTokenMin = HxDefault[int32]()
	obj.Hx_Field_oldTokenMax = HxDefault[int32]()
	obj.Hx_Field_tokens = HxDefault[Hx_Typedef_list]()
	obj.Hx_Field_preprocStack = HxDefault[HxArray[any]]()
	return obj
}

func Hx_Obj_loreline_lorscript_parser_CreateInstance() *Hx_Obj_loreline_lorscript_parser {
	obj := Hx_Obj_loreline_lorscript_parser_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_New() {
	{
		this.Hx_Field_preprocesorValues = (func() *Hx_Obj_haxe_imap {
			_hx_icast_45 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_45 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_45, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_45.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_uid = 0
	}
	{
		this.Hx_Field_line = 1
		this.Hx_Field_opChars = "+*/-=!><&|^%~"
		this.Hx_Field_identChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
		var _hx_tmp_0 HxArray[string]
		_ = _hx_tmp_0
		_hx_tmp_0 = HxMakeArray[string]("%")
		var _hx_tmp_1 HxArray[string]
		_ = _hx_tmp_1
		_hx_tmp_1 = HxMakeArray[string]("*", "/")
		var _hx_tmp_2 HxArray[string]
		_ = _hx_tmp_2
		_hx_tmp_2 = HxMakeArray[string]("+", "-")
		var _hx_tmp_3 HxArray[string]
		_ = _hx_tmp_3
		_hx_tmp_3 = HxMakeArray[string]("<<", ">>", ">>>")
		var _hx_tmp_4 HxArray[string]
		_ = _hx_tmp_4
		_hx_tmp_4 = HxMakeArray[string]("|", "&", "^")
		var _hx_tmp_5 HxArray[string]
		_ = _hx_tmp_5
		_hx_tmp_5 = HxMakeArray[string]("==", "!=", ">", "<", ">=", "<=")
		var _hx_tmp_6 HxArray[string]
		_ = _hx_tmp_6
		_hx_tmp_6 = HxMakeArray[string]("...")
		var _hx_tmp_7 HxArray[string]
		_ = _hx_tmp_7
		_hx_tmp_7 = HxMakeArray[string]("&&")
		var _hx_tmp_8 HxArray[string]
		_ = _hx_tmp_8
		_hx_tmp_8 = HxMakeArray[string]("||")
		var _hx_tmp_9 HxArray[string]
		_ = _hx_tmp_9
		_hx_tmp_9 = HxMakeArray[string]("=", "+=", "-=", "*=", "/=", "%=", "<<=", ">>=", ">>>=", "|=", "&=", "^=", "=>")
		var priorities HxArray[HxArray[string]]
		_ = priorities
		priorities = HxMakeArray[HxArray[string]](_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, _hx_tmp_7, _hx_tmp_8, _hx_tmp_9, HxMakeArray[string]("->"))
		this.Hx_Field_opPriority = (func() *Hx_Obj_haxe_imap {
			_hx_icast_46 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_46 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_46, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_46.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		this.Hx_Field_opRightAssoc = (func() *Hx_Obj_haxe_imap {
			_hx_icast_47 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_47 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_47, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_47.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = priorities.Len()
			for _g < _g1 {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_10
				{
					var tmp__g1_1 int32
					_ = tmp__g1_1
					tmp__g1_1 = 0
					var _g2 HxArray[string]
					_ = _g2
					_g2 = priorities.Get(i)
					for {
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = tmp__g1_1
						if !(_hx_tmp_11 < _g2.Len()) {
							break
						}

						var x string
						_ = x
						x = _g2.Get(tmp__g1_1)
						tmp__g1_1++
						func() func(_hx_param_key string, _hx_param_value int32) {
							var _hx_closure_src_327 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_327
							_hx_closure_src_327 = this.Hx_Field_opPriority.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value int32) {
								_hx_closure_src_327(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(x, i)
						if i == 9 {
							func() func(_hx_param_key string, _hx_param_value bool) {
								var _hx_closure_src_328 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_328
								_hx_closure_src_328 = this.Hx_Field_opRightAssoc.VTable.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value bool) {
									_hx_closure_src_328(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(x, true)
						}
					}
				}
			}
		}

		{
			{
				var x string
				_ = x
				x = "!"
				var _hx_tmp_12 string
				_ = _hx_tmp_12
				_hx_tmp_12 = x
				var _hx_tmp_13 int32
				_ = _hx_tmp_13
				if (x == "++") || (x == "--") {
					_hx_tmp_13 = -1
				} else {
					_hx_tmp_13 = -2
				}

				func() func(_hx_param_key string, _hx_param_value int32) {
					var _hx_closure_src_329 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_329
					_hx_closure_src_329 = this.Hx_Field_opPriority.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value int32) {
						_hx_closure_src_329(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_12, _hx_tmp_13)
			}

			{
				var x1 string
				_ = x1
				x1 = "++"
				var _hx_tmp_14 string
				_ = _hx_tmp_14
				_hx_tmp_14 = x1
				var _hx_tmp_15 int32
				_ = _hx_tmp_15
				if (x1 == "++") || (x1 == "--") {
					_hx_tmp_15 = -1
				} else {
					_hx_tmp_15 = -2
				}

				func() func(_hx_param_key string, _hx_param_value int32) {
					var _hx_closure_src_330 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_330
					_hx_closure_src_330 = this.Hx_Field_opPriority.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value int32) {
						_hx_closure_src_330(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_14, _hx_tmp_15)
			}

			{
				var x2 string
				_ = x2
				x2 = "--"
				var _hx_tmp_16 string
				_ = _hx_tmp_16
				_hx_tmp_16 = x2
				var _hx_tmp_17 int32
				_ = _hx_tmp_17
				if (x2 == "++") || (x2 == "--") {
					_hx_tmp_17 = -1
				} else {
					_hx_tmp_17 = -2
				}

				func() func(_hx_param_key string, _hx_param_value int32) {
					var _hx_closure_src_331 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_331
					_hx_closure_src_331 = this.Hx_Field_opPriority.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value int32) {
						_hx_closure_src_331(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_16, _hx_tmp_17)
			}

			{
				var x3 string
				_ = x3
				x3 = "~"
				var _hx_tmp_18 string
				_ = _hx_tmp_18
				_hx_tmp_18 = x3
				var _hx_tmp_19 int32
				_ = _hx_tmp_19
				if (x3 == "++") || (x3 == "--") {
					_hx_tmp_19 = -1
				} else {
					_hx_tmp_19 = -2
				}

				func() func(_hx_param_key string, _hx_param_value int32) {
					var _hx_closure_src_332 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_332
					_hx_closure_src_332 = this.Hx_Field_opPriority.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value int32) {
						_hx_closure_src_332(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_18, _hx_tmp_19)
			}
		}
	}

}

var Hx_Obj_loreline_lorscript_parser_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lorscript.Parser",
	HxMakeArray[string](),
	HxMakeArray[string]("line", "opChars", "identChars", "opPriority", "opRightAssoc", "preprocesorValues", "allowJSON", "allowTypes", "allowMetadata", "resumeErrors", "input", "readPos", "offset", "char", "ops", "idents", "uid", "origin", "tokenMin", "tokenMax", "oldTokenMin", "oldTokenMax", "tokens", "get_currentPos", "error", "invalidChar", "initParser", "parseString", "unexpected", "push", "ensure", "ensureToken", "maybe", "getIdent", "expr", "pmin", "pmax", "mk", "isBlock", "parseFullExpr", "parseObject", "parseExpr", "parseLambda", "parseMetaArgs", "mapCompr", "makeUnop", "makeBinop", "parseStructure", "parseExprNext", "parseFunctionArgs", "parseFunctionDecl", "parsePath", "parseType", "parseTypeNext", "parseExprList", "parseModule", "parseMetadata", "parseParams", "parseModuleDecl", "parseField", "readChar", "readString", "token", "_token", "preprocValue", "preprocStack", "parsePreproCond", "evalPreproCond", "preprocess", "skipTokens", "tokenComment", "constString", "tokenString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lorscript_parser_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_lorscript_parser_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lorscript_parser_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lorscript_parser_RTTI
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_As_Hx_Obj_loreline_lorscript_parser() *Hx_Obj_loreline_lorscript_parser {
	return this
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_get_currentPos() int32 {
	return (this.Hx_Field_readPos + this.Hx_Field_offset)
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_error(err Hx_Enum_loreline_lorscript_errordef, pmin int32, pmax int32) {
	if !this.Hx_Field_resumeErrors {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(err, pmin, pmax, this.Hx_Field_origin, this.Hx_Field_line)))))
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_invalidChar(c int32) {
	if !this.Hx_Field_resumeErrors {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidChar{c}, (this.Hx_Field_readPos - 1), (this.Hx_Field_readPos - 1), this.Hx_Field_origin, this.Hx_Field_line)))))
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_initParser(origin string, pos int32) {
	this.Hx_Field_preprocStack = HxMakeArray[any]().Dyn()
	this.Hx_Field_origin = origin
	this.Hx_Field_readPos = 0
	this.Hx_Field_oldTokenMin = pos
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_oldTokenMin
	this.Hx_Field_tokenMin = _hx_tmp_0
	this.Hx_Field_oldTokenMax = pos
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_oldTokenMax
	this.Hx_Field_tokenMax = _hx_tmp_1
	this.Hx_Field_tokens = Hx_Obj_haxe_ds_list_CreateInstance()
	this.Hx_Field_offset = pos
	this.Hx_Field_char = -1
	this.Hx_Field_ops = HxMakeArray[bool]()
	this.Hx_Field_idents = HxMakeArray[bool]()
	this.Hx_Field_uid = 0
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = int32(utf8.RuneCountInString(this.Hx_Field_opChars))
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			var _hx_tmp_3 HxArray[bool]
			_ = _hx_tmp_3
			_hx_tmp_3 = this.Hx_Field_ops
			_hx_tmp_3.Set(((int32)((Hx_Field_go_haxe_hxstring_charCodeAt(this.Hx_Field_opChars, i)).Value)), true)
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = int32(utf8.RuneCountInString(this.Hx_Field_identChars))
		for _g2 < _g3 {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_4
			var _hx_tmp_5 HxArray[bool]
			_ = _hx_tmp_5
			_hx_tmp_5 = this.Hx_Field_idents
			_hx_tmp_5.Set(((int32)((Hx_Field_go_haxe_hxstring_charCodeAt(this.Hx_Field_identChars, i)).Value)), true)
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseString(s string, origin__isOptional any, position__isOptional any) any {
	origin := HxDefault[HxNullable[string]]()
	_ = origin

	if origin__isOptional != nil {
		origin = HxConvert[HxNullable[string]](origin__isOptional)
	} else {

	}
	position := HxDefault[HxNullable[int32]]()
	_ = position

	if position__isOptional != nil {
		position = HxConvert[HxNullable[int32]](position__isOptional)
	} else {

	}
	{
		if (origin).Valid == false {
			origin = HxNullableString("loreline.lorscript")
		}

		if (position).Valid == false {
			position = HxNullable[int32]{Value: 0, Valid: true}
		}

		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = (origin).Value
		this.VTable.Hx_Field_initParser(_hx_tmp_0, (position).Value)
		this.Hx_Field_input = s
		this.Hx_Field_readPos = 0
		var a HxArray[any]
		_ = a
		a = HxMakeArray[any]().Dyn()
		for {
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			if (tk == Hx_Enum_loreline_lorscript_token_TEof{}) {
				break
			}

			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			this.VTable.Hx_Field_parseFullExpr(a)
		}

		if Hx_Field_go_haxe_hxdynamic_getArrayLength(a) == 1 {
			return a.Get(0)
		} else {
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EBlock{a}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: 0, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = ((any)(e))
				var hx_nullable_f10 HxNullable[int32]
				_ = hx_nullable_f10
				hx_nullable_f10 = pmin
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				if (hx_nullable_f10).Valid == true {
					_hx_tmp_2 = (hx_nullable_f10).Value
				} else {
					_hx_tmp_2 = nil
				}

				var hx_nullable_f11 HxNullable[int32]
				_ = hx_nullable_f11
				hx_nullable_f11 = pmax
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				if (hx_nullable_f11).Valid == true {
					_hx_tmp_3 = (hx_nullable_f11).Value
				} else {
					_hx_tmp_3 = nil
				}

				return any(map[string]any{"e": _hx_tmp_1, "pmin": _hx_tmp_2, "pmax": _hx_tmp_3, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_unexpected(tk Hx_Enum_loreline_lorscript_token) any {
	{
		var err Hx_Enum_loreline_lorscript_errordef
		_ = err
		err = Hx_Enum_loreline_lorscript_errordef_EUnexpected{this.VTable.Hx_Field_tokenString(tk)}

		if !this.Hx_Field_resumeErrors {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(err, this.Hx_Field_tokenMin, this.Hx_Field_tokenMax, this.Hx_Field_origin, this.Hx_Field_line)))))
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_push(tk Hx_Enum_loreline_lorscript_token) {
	this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
	this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
	this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_ensure(tk Hx_Enum_loreline_lorscript_token) {
	var t Hx_Enum_loreline_lorscript_token
	_ = t
	t = this.VTable.Hx_Field_token()
	if t != tk {
		this.VTable.Hx_Field_unexpected(t)
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_ensureToken(tk Hx_Enum_loreline_lorscript_token) {
	var t Hx_Enum_loreline_lorscript_token
	_ = t
	t = this.VTable.Hx_Field_token()
	if !func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
		var _hx_closure_src_306 func(a any, b any) bool
		_ = _hx_closure_src_306
		_hx_closure_src_306 = Hx_Field_type_enumEq
		return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
			return _hx_closure_src_306(((any)(a)), ((any)(b)))
		}
	}()(t, tk) {
		this.VTable.Hx_Field_unexpected(t)
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_maybe(tk Hx_Enum_loreline_lorscript_token) bool {
	var t Hx_Enum_loreline_lorscript_token
	_ = t
	t = this.VTable.Hx_Field_token()
	if func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
		var _hx_closure_src_307 func(a any, b any) bool
		_ = _hx_closure_src_307
		_hx_closure_src_307 = Hx_Field_type_enumEq
		return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
			return _hx_closure_src_307(((any)(a)), ((any)(b)))
		}
	}()(t, tk) {
		return true
	}

	{
		this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
		this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
		this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
	}

	return false
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_getIdent() string {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	if tk.Hx_Field_enumIndex() == 2 {
		var id string
		_ = id
		id = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
		return id
	} else {
		this.VTable.Hx_Field_unexpected(tk)
		return HxStringNull
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_expr(e any) Hx_Enum_loreline_lorscript_exprdef {
	return func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_pmin(e any) int32 {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		return 0
	} else {
		return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_pmax(e any) int32 {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		return 0
	} else {
		return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_mk(e Hx_Enum_loreline_lorscript_exprdef, pmin__isOptional any, pmax__isOptional any) any {
	pmin := HxDefault[HxNullable[int32]]()
	_ = pmin

	if pmin__isOptional != nil {
		pmin = HxConvert[HxNullable[int32]](pmin__isOptional)
	} else {

	}
	pmax := HxDefault[HxNullable[int32]]()
	_ = pmax

	if pmax__isOptional != nil {
		pmax = HxConvert[HxNullable[int32]](pmax__isOptional)
	} else {

	}
	{
		if (pmin).Valid == false {
			var hx_nullable_t35 HxNullable[int32]
			_ = hx_nullable_t35
			hx_nullable_t35 = HxNullable[int32]{}

			pmin = HxNullable[int32]{Value: (hx_nullable_t35).Value, Valid: (hx_nullable_t35).Valid}
		}

		if (pmax).Valid == false {
			var hx_nullable_t34 HxNullable[int32]
			_ = hx_nullable_t34
			hx_nullable_t34 = HxNullable[int32]{}

			pmax = HxNullable[int32]{Value: (hx_nullable_t34).Value, Valid: (hx_nullable_t34).Valid}
		}

		if e == nil {
			return nil
		}

		if (pmin).Valid == false {
			pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
		}

		if (pmax).Valid == false {
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
		}

		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = ((any)(e))
		var hx_nullable_f12 HxNullable[int32]
		_ = hx_nullable_f12
		hx_nullable_f12 = pmin
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		if (hx_nullable_f12).Valid == true {
			_hx_tmp_1 = (hx_nullable_f12).Value
		} else {
			_hx_tmp_1 = nil
		}

		var hx_nullable_f13 HxNullable[int32]
		_ = hx_nullable_f13
		hx_nullable_f13 = pmax
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		if (hx_nullable_f13).Valid == true {
			_hx_tmp_2 = (hx_nullable_f13).Value
		} else {
			_hx_tmp_2 = nil
		}

		return any(map[string]any{"e": _hx_tmp_0, "pmin": _hx_tmp_1, "pmax": _hx_tmp_2, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_isBlock(e any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		return false
	}

	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 2:
		{
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			t = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_t
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EVar).Hx_Field_e
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				return this.VTable.Hx_Field_isBlock(e)
			} else {
				if t != nil {
					if t == nil {
						return false
					} else {
						return (t.Hx_Field_enumIndex() == 2)
					}
				} else {
					return false
				}
			}
		}

	case 4:
		{
			return true
		}

	case 6:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 7:
		{
			var prefix bool
			_ = prefix
			prefix = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_prefix
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if !prefix {
				_hx_tmp_0 = this.VTable.Hx_Field_isBlock(e)
			} else {
				_hx_tmp_0 = false
			}

			return _hx_tmp_0
		}

	case 9:
		{
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e2, nil))) {
				return this.VTable.Hx_Field_isBlock(e2)
			} else {
				return this.VTable.Hx_Field_isBlock(e1)
			}
		}

	case 10:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 11:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_e
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 14:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EFunction).Hx_Field_e
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 15:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EReturn).Hx_Field_e
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))) {
				_hx_tmp_1 = this.VTable.Hx_Field_isBlock(e)
			} else {
				_hx_tmp_1 = false
			}

			return _hx_tmp_1
		}

	case 20:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_ETry).Hx_Field_ecatch
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 21:
		{
			return true
		}

	case 23:
		{
			return true
		}

	case 24:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
			return this.VTable.Hx_Field_isBlock(e)
		}

	case 25:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EMeta).Hx_Field_e
			return this.VTable.Hx_Field_isBlock(e)
		}

	default:
		{
			return false
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseFullExpr(exprs HxArray[any]) {
	var e any
	_ = e
	e = this.VTable.Hx_Field_parseExpr()
	_ = (Hx_Array_Push[any](HxMakeArrayView[any](exprs), e))
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (tk == Hx_Enum_loreline_lorscript_token_TComma{}) {
			_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(e, nil))))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = (Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() == 2)
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		e = this.VTable.Hx_Field_parseStructure("var")
		_ = (Hx_Array_Push[any](HxMakeArrayView[any](exprs), e))
		tk = this.VTable.Hx_Field_token()
	}

	if (tk != Hx_Enum_loreline_lorscript_token_TSemicolon{}) && (tk != Hx_Enum_loreline_lorscript_token_TEof{}) {
		if this.VTable.Hx_Field_isBlock(e) {
			this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
			this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
			this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
		} else {
			this.VTable.Hx_Field_unexpected(tk)
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseObject(p1 HxNullable[int32]) any {
	var fl HxArray[any]
	_ = fl
	fl = HxMakeArray[any]().Dyn()
hx_label_0:
	for {
		var tk Hx_Enum_loreline_lorscript_token
		_ = tk
		tk = this.VTable.Hx_Field_token()
		var id HxNullable[string]
		_ = id
		id = HxNullable[string]{Value: HxStringNull}

		switch tk.Hx_Field_enumIndex() {
		case 1:
			{
				var c Hx_Enum_loreline_lorscript_const
				_ = c
				c = tk.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
				{
					if !this.Hx_Field_allowJSON {
						this.VTable.Hx_Field_unexpected(tk)
					}

					if c.Hx_Field_enumIndex() == 2 {
						var s string
						_ = s
						s = c.(Hx_Enum_loreline_lorscript_const_CString).Hx_Field_s
						id = HxNullableString(s)
					} else {
						this.VTable.Hx_Field_unexpected(tk)
					}
				}
			}

		case 2:
			{
				var i string
				_ = i
				i = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
				id = HxNullableString(i)
			}

		case 7:
			{
				break hx_label_0
			}

		default:
			{
				this.VTable.Hx_Field_unexpected(tk)
				break hx_label_0
			}
		}

		{
			var t Hx_Enum_loreline_lorscript_token
			_ = t
			t = this.VTable.Hx_Field_token()
			if (t != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
				this.VTable.Hx_Field_unexpected(t)
			}
		}

		var _hx_tmp_0 HxArray[any]
		_ = _hx_tmp_0
		_hx_tmp_0 = HxMakeArrayView[any](fl)
		var hx_nullable_f14 HxNullable[string]
		_ = hx_nullable_f14
		hx_nullable_f14 = id
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		if (hx_nullable_f14).Valid == true {
			_hx_tmp_1 = (hx_nullable_f14).Value
		} else {
			_hx_tmp_1 = nil
		}

		_ = (Hx_Array_Push[any](_hx_tmp_0, any(map[string]any{"name": _hx_tmp_1, "e": this.VTable.Hx_Field_parseExpr()})))
		tk = this.VTable.Hx_Field_token()
		switch tk.Hx_Field_enumIndex() {
		case 7:
			{
				break hx_label_0
			}

		case 10:
			{
			}

		default:
			{
				this.VTable.Hx_Field_unexpected(tk)
			}
		}
	}

	var e Hx_Enum_loreline_lorscript_exprdef
	_ = e
	e = Hx_Enum_loreline_lorscript_exprdef_EObject{fl}

	var pmin HxNullable[int32]
	_ = pmin
	pmin = p1
	var pmax HxNullable[int32]
	_ = pmax
	pmax = HxNullable[int32]{}

	var tmp any
	_ = tmp
	if e == nil {
		tmp = nil
	} else {
		if (pmin).Valid == false {
			pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
		}

		if (pmax).Valid == false {
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
		}

		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = ((any)(e))
		var hx_nullable_f15 HxNullable[int32]
		_ = hx_nullable_f15
		hx_nullable_f15 = pmin
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		if (hx_nullable_f15).Valid == true {
			_hx_tmp_3 = (hx_nullable_f15).Value
		} else {
			_hx_tmp_3 = nil
		}

		var hx_nullable_f16 HxNullable[int32]
		_ = hx_nullable_f16
		hx_nullable_f16 = pmax
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		if (hx_nullable_f16).Valid == true {
			_hx_tmp_4 = (hx_nullable_f16).Value
		} else {
			_hx_tmp_4 = nil
		}

		tmp = any(map[string]any{"e": _hx_tmp_2, "pmin": _hx_tmp_3, "pmax": _hx_tmp_4, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
	}

	return this.VTable.Hx_Field_parseExprNext(tmp)
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseExpr() any {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	var p1 int32
	_ = p1
	p1 = this.Hx_Field_tokenMin
	switch tk.Hx_Field_enumIndex() {
	case 1:
		{
			var c Hx_Enum_loreline_lorscript_const
			_ = c
			c = tk.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EConst{c}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			var tmp any
			_ = tmp
			if e == nil {
				tmp = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_0 any
				_ = _hx_tmp_0
				_hx_tmp_0 = ((any)(e))
				var hx_nullable_f17 HxNullable[int32]
				_ = hx_nullable_f17
				hx_nullable_f17 = pmin
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				if (hx_nullable_f17).Valid == true {
					_hx_tmp_1 = (hx_nullable_f17).Value
				} else {
					_hx_tmp_1 = nil
				}

				var hx_nullable_f18 HxNullable[int32]
				_ = hx_nullable_f18
				hx_nullable_f18 = pmax
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				if (hx_nullable_f18).Valid == true {
					_hx_tmp_2 = (hx_nullable_f18).Value
				} else {
					_hx_tmp_2 = nil
				}

				tmp = any(map[string]any{"e": _hx_tmp_0, "pmin": _hx_tmp_1, "pmax": _hx_tmp_2, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(tmp)
		}

	case 2:
		{
			var id string
			_ = id
			id = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseStructure(id)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				var e1 Hx_Enum_loreline_lorscript_exprdef
				_ = e1
				e1 = Hx_Enum_loreline_lorscript_exprdef_EIdent{id}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{}

				if e1 == nil {
					e = nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_3 any
					_ = _hx_tmp_3
					_hx_tmp_3 = ((any)(e1))
					var hx_nullable_f19 HxNullable[int32]
					_ = hx_nullable_f19
					hx_nullable_f19 = pmin
					var _hx_tmp_4 any
					_ = _hx_tmp_4
					if (hx_nullable_f19).Valid == true {
						_hx_tmp_4 = (hx_nullable_f19).Value
					} else {
						_hx_tmp_4 = nil
					}

					var hx_nullable_f20 HxNullable[int32]
					_ = hx_nullable_f20
					hx_nullable_f20 = pmax
					var _hx_tmp_5 any
					_ = _hx_tmp_5
					if (hx_nullable_f20).Valid == true {
						_hx_tmp_5 = (hx_nullable_f20).Value
					} else {
						_hx_tmp_5 = nil
					}

					e = any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			}

			return this.VTable.Hx_Field_parseExprNext(e)
		}

	case 3:
		{
			var op string
			_ = op
			op = tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
			if op == "-" {
				var start int32
				_ = start
				start = this.Hx_Field_tokenMin
				var e any
				_ = e
				e = this.VTable.Hx_Field_parseExpr()
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					return this.VTable.Hx_Field_makeUnop(op, e)
				}

				var _g Hx_Enum_loreline_lorscript_exprdef
				_ = _g
				_g = func() Hx_Enum_loreline_lorscript_exprdef {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
					if _hx_d == nil {
						var _hx_z Hx_Enum_loreline_lorscript_exprdef
						return _hx_z
					}
					return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
				}()
				if _g.Hx_Field_enumIndex() == 0 {
					var _gc Hx_Enum_loreline_lorscript_const
					_ = _gc
					_gc = _g.(Hx_Enum_loreline_lorscript_exprdef_EConst).Hx_Field_c
					switch _gc.Hx_Field_enumIndex() {
					case 0:
						{
							var i int32
							_ = i
							i = _gc.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
							var e1 Hx_Enum_loreline_lorscript_exprdef
							_ = e1
							e1 = Hx_Enum_loreline_lorscript_exprdef_EConst{Hx_Enum_loreline_lorscript_const_CInt{-i}}

							var pmin HxNullable[int32]
							_ = pmin
							pmin = HxNullable[int32]{Value: start, Valid: true}

							var _hx_tmp_6 int32
							_ = _hx_tmp_6
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
								_hx_tmp_6 = 0
							} else {
								_hx_tmp_6 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
							}

							var pmax HxNullable[int32]
							_ = pmax
							pmax = HxNullable[int32]{Value: _hx_tmp_6, Valid: true}

							if e1 == nil {
								return nil
							} else {
								if (pmin).Valid == false {
									pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
								}

								if (pmax).Valid == false {
									pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
								}

								var _hx_tmp_7 any
								_ = _hx_tmp_7
								_hx_tmp_7 = ((any)(e1))
								var hx_nullable_f21 HxNullable[int32]
								_ = hx_nullable_f21
								hx_nullable_f21 = pmin
								var _hx_tmp_8 any
								_ = _hx_tmp_8
								if (hx_nullable_f21).Valid == true {
									_hx_tmp_8 = (hx_nullable_f21).Value
								} else {
									_hx_tmp_8 = nil
								}

								var hx_nullable_f22 HxNullable[int32]
								_ = hx_nullable_f22
								hx_nullable_f22 = pmax
								var _hx_tmp_9 any
								_ = _hx_tmp_9
								if (hx_nullable_f22).Valid == true {
									_hx_tmp_9 = (hx_nullable_f22).Value
								} else {
									_hx_tmp_9 = nil
								}

								return any(map[string]any{"e": _hx_tmp_7, "pmin": _hx_tmp_8, "pmax": _hx_tmp_9, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
							}
						}

					case 1:
						{
							var f float64
							_ = f
							f = _gc.(Hx_Enum_loreline_lorscript_const_CFloat).Hx_Field_f
							var e1 Hx_Enum_loreline_lorscript_exprdef
							_ = e1
							e1 = Hx_Enum_loreline_lorscript_exprdef_EConst{Hx_Enum_loreline_lorscript_const_CFloat{-f}}

							var pmin HxNullable[int32]
							_ = pmin
							pmin = HxNullable[int32]{Value: start, Valid: true}

							var _hx_tmp_10 int32
							_ = _hx_tmp_10
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
								_hx_tmp_10 = 0
							} else {
								_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
							}

							var pmax HxNullable[int32]
							_ = pmax
							pmax = HxNullable[int32]{Value: _hx_tmp_10, Valid: true}

							if e1 == nil {
								return nil
							} else {
								if (pmin).Valid == false {
									pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
								}

								if (pmax).Valid == false {
									pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
								}

								var _hx_tmp_11 any
								_ = _hx_tmp_11
								_hx_tmp_11 = ((any)(e1))
								var hx_nullable_f23 HxNullable[int32]
								_ = hx_nullable_f23
								hx_nullable_f23 = pmin
								var _hx_tmp_12 any
								_ = _hx_tmp_12
								if (hx_nullable_f23).Valid == true {
									_hx_tmp_12 = (hx_nullable_f23).Value
								} else {
									_hx_tmp_12 = nil
								}

								var hx_nullable_f24 HxNullable[int32]
								_ = hx_nullable_f24
								hx_nullable_f24 = pmax
								var _hx_tmp_13 any
								_ = _hx_tmp_13
								if (hx_nullable_f24).Valid == true {
									_hx_tmp_13 = (hx_nullable_f24).Value
								} else {
									_hx_tmp_13 = nil
								}

								return any(map[string]any{"e": _hx_tmp_11, "pmin": _hx_tmp_12, "pmax": _hx_tmp_13, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
							}
						}

					default:
						{
							return this.VTable.Hx_Field_makeUnop(op, e)
						}
					}
				} else {
					return this.VTable.Hx_Field_makeUnop(op, e)
				}
			}

			var hx_cmp_5 HxNullable[int32]
			_ = hx_cmp_5
			hx_cmp_5 = func() func(_hx_param_key string) HxNullable[int32] {
				var _hx_closure_src_310 func(_hx_param_key any) any
				_ = _hx_closure_src_310
				_hx_closure_src_310 = this.Hx_Field_opPriority.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) HxNullable[int32] {
					var hx_nullable_t37 any
					_ = hx_nullable_t37
					hx_nullable_t37 = _hx_closure_src_310(((any)(_hx_param_key)))
					var _hx_tmp_14 HxNullable[int32]
					_ = _hx_tmp_14
					if hx_nullable_t37 == nil {
						_hx_tmp_14 = HxNullable[int32]{}
					} else {
						_hx_tmp_14 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t37))), Valid: true}
					}

					return _hx_tmp_14
				}
			}()(op)
			if (hx_cmp_5).Valid && true && ((hx_cmp_5).Value < 0) {
				var _hx_tmp_15 string
				_ = _hx_tmp_15
				_hx_tmp_15 = op
				return this.VTable.Hx_Field_makeUnop(_hx_tmp_15, this.VTable.Hx_Field_parseExpr())
			}

			return this.VTable.Hx_Field_unexpected(tk)
		}

	case 4:
		{
			tk = this.VTable.Hx_Field_token()
			if (tk == Hx_Enum_loreline_lorscript_token_TPClose{}) {
				{
					var t Hx_Enum_loreline_lorscript_token
					_ = t
					t = this.VTable.Hx_Field_token()
					var _hx_tmp_16 Hx_Enum_loreline_lorscript_token
					_ = _hx_tmp_16
					_hx_tmp_16 = t
					if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
						var _hx_closure_src_309 func(a any, b any) bool
						_ = _hx_closure_src_309
						_hx_closure_src_309 = Hx_Field_type_enumEq
						return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
							return _hx_closure_src_309(((any)(a)), ((any)(b)))
						}
					}()(_hx_tmp_16, Hx_Enum_loreline_lorscript_token_TOp{"->"})) {
						this.VTable.Hx_Field_unexpected(t)
					}
				}

				var eret any
				_ = eret
				eret = this.VTable.Hx_Field_parseExpr()
				var e Hx_Enum_loreline_lorscript_exprdef
				_ = e
				e = Hx_Enum_loreline_lorscript_exprdef_EReturn{eret}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: p1, Valid: true}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{}

				var e1 any
				_ = e1
				if e == nil {
					e1 = nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_17 any
					_ = _hx_tmp_17
					_hx_tmp_17 = ((any)(e))
					var hx_nullable_f25 HxNullable[int32]
					_ = hx_nullable_f25
					hx_nullable_f25 = pmin
					var _hx_tmp_18 any
					_ = _hx_tmp_18
					if (hx_nullable_f25).Valid == true {
						_hx_tmp_18 = (hx_nullable_f25).Value
					} else {
						_hx_tmp_18 = nil
					}

					var hx_nullable_f26 HxNullable[int32]
					_ = hx_nullable_f26
					hx_nullable_f26 = pmax
					var _hx_tmp_19 any
					_ = _hx_tmp_19
					if (hx_nullable_f26).Valid == true {
						_hx_tmp_19 = (hx_nullable_f26).Value
					} else {
						_hx_tmp_19 = nil
					}

					e1 = any(map[string]any{"e": _hx_tmp_17, "pmin": _hx_tmp_18, "pmax": _hx_tmp_19, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				var _hx_tmp_20 HxArray[any]
				_ = _hx_tmp_20
				_hx_tmp_20 = HxMakeArray[any]().Dyn()
				var _hx_tmp_21 any
				_ = _hx_tmp_21
				_hx_tmp_21 = e1
				var e2 Hx_Enum_loreline_lorscript_exprdef
				_ = e2
				e2 = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_20, _hx_tmp_21, HxNullable[string]{Value: HxStringNull}, nil}

				var pmin1 HxNullable[int32]
				_ = pmin1
				pmin1 = HxNullable[int32]{Value: p1, Valid: true}

				var pmax1 HxNullable[int32]
				_ = pmax1
				pmax1 = HxNullable[int32]{}

				if e2 == nil {
					return nil
				} else {
					if (pmin1).Valid == false {
						pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax1).Valid == false {
						pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_22 any
					_ = _hx_tmp_22
					_hx_tmp_22 = ((any)(e2))
					var hx_nullable_f27 HxNullable[int32]
					_ = hx_nullable_f27
					hx_nullable_f27 = pmin1
					var _hx_tmp_23 any
					_ = _hx_tmp_23
					if (hx_nullable_f27).Valid == true {
						_hx_tmp_23 = (hx_nullable_f27).Value
					} else {
						_hx_tmp_23 = nil
					}

					var hx_nullable_f28 HxNullable[int32]
					_ = hx_nullable_f28
					hx_nullable_f28 = pmax1
					var _hx_tmp_24 any
					_ = _hx_tmp_24
					if (hx_nullable_f28).Valid == true {
						_hx_tmp_24 = (hx_nullable_f28).Value
					} else {
						_hx_tmp_24 = nil
					}

					return any(map[string]any{"e": _hx_tmp_22, "pmin": _hx_tmp_23, "pmax": _hx_tmp_24, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			}

			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			tk = this.VTable.Hx_Field_token()
			switch tk.Hx_Field_enumIndex() {
			case 5:
				{
					var e1 Hx_Enum_loreline_lorscript_exprdef
					_ = e1
					e1 = Hx_Enum_loreline_lorscript_exprdef_EParent{e}

					var pmin HxNullable[int32]
					_ = pmin
					pmin = HxNullable[int32]{Value: p1, Valid: true}

					var pmax HxNullable[int32]
					_ = pmax
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}

					var tmp any
					_ = tmp
					if e1 == nil {
						tmp = nil
					} else {
						if (pmin).Valid == false {
							pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
						}

						if (pmax).Valid == false {
							pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
						}

						var _hx_tmp_25 any
						_ = _hx_tmp_25
						_hx_tmp_25 = ((any)(e1))
						var hx_nullable_f29 HxNullable[int32]
						_ = hx_nullable_f29
						hx_nullable_f29 = pmin
						var _hx_tmp_26 any
						_ = _hx_tmp_26
						if (hx_nullable_f29).Valid == true {
							_hx_tmp_26 = (hx_nullable_f29).Value
						} else {
							_hx_tmp_26 = nil
						}

						var hx_nullable_f30 HxNullable[int32]
						_ = hx_nullable_f30
						hx_nullable_f30 = pmax
						var _hx_tmp_27 any
						_ = _hx_tmp_27
						if (hx_nullable_f30).Valid == true {
							_hx_tmp_27 = (hx_nullable_f30).Value
						} else {
							_hx_tmp_27 = nil
						}

						tmp = any(map[string]any{"e": _hx_tmp_25, "pmin": _hx_tmp_26, "pmax": _hx_tmp_27, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
					}

					return this.VTable.Hx_Field_parseExprNext(tmp)
				}

			case 10:
				{
					var _g Hx_Enum_loreline_lorscript_exprdef
					_ = _g
					_g = func() Hx_Enum_loreline_lorscript_exprdef {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_exprdef
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
					}()
					if _g.Hx_Field_enumIndex() == 1 {
						var v string
						_ = v
						v = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
						var _hx_tmp_28 HxArray[any]
						_ = _hx_tmp_28
						_hx_tmp_28 = HxMakeArray[any](any(map[string]any{"name": ((any)(v))})).Dyn()
						var _hx_tmp_29 int32
						_ = _hx_tmp_29
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
							_hx_tmp_29 = 0
						} else {
							_hx_tmp_29 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
						}

						return this.VTable.Hx_Field_parseLambda(_hx_tmp_28, HxNullable[int32]{Value: _hx_tmp_29, Valid: true})
					}
				}

			case 15:
				{
					var t Hx_Enum_loreline_lorscript_ctype
					_ = t
					t = this.VTable.Hx_Field_parseType()
					tk = this.VTable.Hx_Field_token()
					switch tk.Hx_Field_enumIndex() {
					case 5:
						{
							var e1 Hx_Enum_loreline_lorscript_exprdef
							_ = e1
							e1 = Hx_Enum_loreline_lorscript_exprdef_ECheckType{e, t}

							var pmin HxNullable[int32]
							_ = pmin
							pmin = HxNullable[int32]{Value: p1, Valid: true}

							var pmax HxNullable[int32]
							_ = pmax
							pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}

							var tmp any
							_ = tmp
							if e1 == nil {
								tmp = nil
							} else {
								if (pmin).Valid == false {
									pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
								}

								if (pmax).Valid == false {
									pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
								}

								var _hx_tmp_30 any
								_ = _hx_tmp_30
								_hx_tmp_30 = ((any)(e1))
								var hx_nullable_f31 HxNullable[int32]
								_ = hx_nullable_f31
								hx_nullable_f31 = pmin
								var _hx_tmp_31 any
								_ = _hx_tmp_31
								if (hx_nullable_f31).Valid == true {
									_hx_tmp_31 = (hx_nullable_f31).Value
								} else {
									_hx_tmp_31 = nil
								}

								var hx_nullable_f32 HxNullable[int32]
								_ = hx_nullable_f32
								hx_nullable_f32 = pmax
								var _hx_tmp_32 any
								_ = _hx_tmp_32
								if (hx_nullable_f32).Valid == true {
									_hx_tmp_32 = (hx_nullable_f32).Value
								} else {
									_hx_tmp_32 = nil
								}

								tmp = any(map[string]any{"e": _hx_tmp_30, "pmin": _hx_tmp_31, "pmax": _hx_tmp_32, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
							}

							return this.VTable.Hx_Field_parseExprNext(tmp)
						}

					case 10:
						{
							var _g Hx_Enum_loreline_lorscript_exprdef
							_ = _g
							_g = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if _g.Hx_Field_enumIndex() == 1 {
								var v string
								_ = v
								v = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
								var _hx_tmp_33 HxArray[any]
								_ = _hx_tmp_33
								_hx_tmp_33 = HxMakeArray[any](any(map[string]any{"name": ((any)(v)), "t": ((any)(t))})).Dyn()
								var _hx_tmp_34 int32
								_ = _hx_tmp_34
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
									_hx_tmp_34 = 0
								} else {
									_hx_tmp_34 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
								}

								return this.VTable.Hx_Field_parseLambda(_hx_tmp_33, HxNullable[int32]{Value: _hx_tmp_34, Valid: true})
							}
						}

					default:
						{
						}
					}
				}

			default:
				{
				}
			}

			return this.VTable.Hx_Field_unexpected(tk)
		}

	case 6:
		{
			tk = this.VTable.Hx_Field_token()
			switch tk.Hx_Field_enumIndex() {
			case 1:
				{
					var c Hx_Enum_loreline_lorscript_const
					_ = c
					c = tk.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
					if this.Hx_Field_allowJSON {
						if c.Hx_Field_enumIndex() == 2 {
							var tk2 Hx_Enum_loreline_lorscript_token
							_ = tk2
							tk2 = this.VTable.Hx_Field_token()
							{
								this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk2)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
								this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
								this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
							}

							{
								this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
								this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
								this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
							}

							if tk2.Hx_Field_enumIndex() == 15 {
								return this.VTable.Hx_Field_parseExprNext(this.VTable.Hx_Field_parseObject(HxNullable[int32]{Value: p1, Valid: true}))
							}
						} else {
							this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
							this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
							this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
						}
					} else {
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}
				}

			case 2:
				{
					var tk2 Hx_Enum_loreline_lorscript_token
					_ = tk2
					tk2 = this.VTable.Hx_Field_token()
					{
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk2)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}

					{
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}

					if tk2.Hx_Field_enumIndex() == 15 {
						return this.VTable.Hx_Field_parseExprNext(this.VTable.Hx_Field_parseObject(HxNullable[int32]{Value: p1, Valid: true}))
					}
				}

			case 7:
				{
					var e Hx_Enum_loreline_lorscript_exprdef
					_ = e
					e = Hx_Enum_loreline_lorscript_exprdef_EObject{HxMakeArray[any]().Dyn()}

					var pmin HxNullable[int32]
					_ = pmin
					pmin = HxNullable[int32]{Value: p1, Valid: true}

					var pmax HxNullable[int32]
					_ = pmax
					pmax = HxNullable[int32]{}

					var tmp any
					_ = tmp
					if e == nil {
						tmp = nil
					} else {
						if (pmin).Valid == false {
							pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
						}

						if (pmax).Valid == false {
							pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
						}

						var _hx_tmp_35 any
						_ = _hx_tmp_35
						_hx_tmp_35 = ((any)(e))
						var hx_nullable_f33 HxNullable[int32]
						_ = hx_nullable_f33
						hx_nullable_f33 = pmin
						var _hx_tmp_36 any
						_ = _hx_tmp_36
						if (hx_nullable_f33).Valid == true {
							_hx_tmp_36 = (hx_nullable_f33).Value
						} else {
							_hx_tmp_36 = nil
						}

						var hx_nullable_f34 HxNullable[int32]
						_ = hx_nullable_f34
						hx_nullable_f34 = pmax
						var _hx_tmp_37 any
						_ = _hx_tmp_37
						if (hx_nullable_f34).Valid == true {
							_hx_tmp_37 = (hx_nullable_f34).Value
						} else {
							_hx_tmp_37 = nil
						}

						tmp = any(map[string]any{"e": _hx_tmp_35, "pmin": _hx_tmp_36, "pmax": _hx_tmp_37, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
					}

					return this.VTable.Hx_Field_parseExprNext(tmp)
				}

			default:
				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}
			}

			var a HxArray[any]
			_ = a
			a = HxMakeArray[any]().Dyn()
			for {
				this.VTable.Hx_Field_parseFullExpr(a)
				tk = this.VTable.Hx_Field_token()
				if (tk == Hx_Enum_loreline_lorscript_token_TBrClose{}) || (this.Hx_Field_resumeErrors && (tk == Hx_Enum_loreline_lorscript_token_TEof{})) {
					break
				}

				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}
			}

			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EBlock{a}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_38 any
				_ = _hx_tmp_38
				_hx_tmp_38 = ((any)(e))
				var hx_nullable_f35 HxNullable[int32]
				_ = hx_nullable_f35
				hx_nullable_f35 = pmin
				var _hx_tmp_39 any
				_ = _hx_tmp_39
				if (hx_nullable_f35).Valid == true {
					_hx_tmp_39 = (hx_nullable_f35).Value
				} else {
					_hx_tmp_39 = nil
				}

				var hx_nullable_f36 HxNullable[int32]
				_ = hx_nullable_f36
				hx_nullable_f36 = pmax
				var _hx_tmp_40 any
				_ = _hx_tmp_40
				if (hx_nullable_f36).Valid == true {
					_hx_tmp_40 = (hx_nullable_f36).Value
				} else {
					_hx_tmp_40 = nil
				}

				return any(map[string]any{"e": _hx_tmp_38, "pmin": _hx_tmp_39, "pmax": _hx_tmp_40, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case 12:
		{
			var a HxArray[any]
			_ = a
			a = HxMakeArray[any]().Dyn()
			tk = this.VTable.Hx_Field_token()
			var first bool
			_ = first
			first = true
			for (tk != Hx_Enum_loreline_lorscript_token_TBkClose{}) && (!this.Hx_Field_resumeErrors || (tk != Hx_Enum_loreline_lorscript_token_TEof{})) {
				if !first {
					if (tk != Hx_Enum_loreline_lorscript_token_TComma{}) {
						this.VTable.Hx_Field_unexpected(tk)
					} else {
						tk = this.VTable.Hx_Field_token()
					}
				}

				first = false
				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}

				var _hx_tmp_41 HxArray[any]
				_ = _hx_tmp_41
				_hx_tmp_41 = HxMakeArrayView[any](a)
				_ = (Hx_Array_Push[any](_hx_tmp_41, this.VTable.Hx_Field_parseExpr()))
				tk = this.VTable.Hx_Field_token()
			}

			var _hx_tmp_42 bool
			_ = _hx_tmp_42
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(a) == 1 {
				_hx_tmp_42 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(a.Get(0), nil))))
			} else {
				_hx_tmp_42 = false
			}

			if _hx_tmp_42 {
				switch Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(a.Get(0), "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() {
				case 10:
					{
						var _hx_tmp_43 int32
						_ = _hx_tmp_43
						_hx_tmp_43 = this.Hx_Field_uid
						this.Hx_Field_uid = (this.Hx_Field_uid + 1)
						var tmp string
						_ = tmp
						tmp = ("__a_" + Hx_Field_go_haxe_hxdynamic_toString(_hx_tmp_43))
						var e Hx_Enum_loreline_lorscript_exprdef
						_ = e
						e = Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{HxMakeArray[any]().Dyn()}

						var pmin HxNullable[int32]
						_ = pmin
						pmin = HxNullable[int32]{Value: p1, Valid: true}

						var pmax HxNullable[int32]
						_ = pmax
						pmax = HxNullable[int32]{}

						var e1 any
						_ = e1
						if e == nil {
							e1 = nil
						} else {
							if (pmin).Valid == false {
								pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax).Valid == false {
								pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_44 any
							_ = _hx_tmp_44
							_hx_tmp_44 = ((any)(e))
							var hx_nullable_f37 HxNullable[int32]
							_ = hx_nullable_f37
							hx_nullable_f37 = pmin
							var _hx_tmp_45 any
							_ = _hx_tmp_45
							if (hx_nullable_f37).Valid == true {
								_hx_tmp_45 = (hx_nullable_f37).Value
							} else {
								_hx_tmp_45 = nil
							}

							var hx_nullable_f38 HxNullable[int32]
							_ = hx_nullable_f38
							hx_nullable_f38 = pmax
							var _hx_tmp_46 any
							_ = _hx_tmp_46
							if (hx_nullable_f38).Valid == true {
								_hx_tmp_46 = (hx_nullable_f38).Value
							} else {
								_hx_tmp_46 = nil
							}

							e1 = any(map[string]any{"e": _hx_tmp_44, "pmin": _hx_tmp_45, "pmax": _hx_tmp_46, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e2 Hx_Enum_loreline_lorscript_exprdef
						_ = e2
						e2 = Hx_Enum_loreline_lorscript_exprdef_EVar{tmp, nil, e1}

						var pmin1 HxNullable[int32]
						_ = pmin1
						pmin1 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax1 HxNullable[int32]
						_ = pmax1
						pmax1 = HxNullable[int32]{}

						var e3 any
						_ = e3
						if e2 == nil {
							e3 = nil
						} else {
							if (pmin1).Valid == false {
								pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax1).Valid == false {
								pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_47 any
							_ = _hx_tmp_47
							_hx_tmp_47 = ((any)(e2))
							var hx_nullable_f39 HxNullable[int32]
							_ = hx_nullable_f39
							hx_nullable_f39 = pmin1
							var _hx_tmp_48 any
							_ = _hx_tmp_48
							if (hx_nullable_f39).Valid == true {
								_hx_tmp_48 = (hx_nullable_f39).Value
							} else {
								_hx_tmp_48 = nil
							}

							var hx_nullable_f40 HxNullable[int32]
							_ = hx_nullable_f40
							hx_nullable_f40 = pmax1
							var _hx_tmp_49 any
							_ = _hx_tmp_49
							if (hx_nullable_f40).Valid == true {
								_hx_tmp_49 = (hx_nullable_f40).Value
							} else {
								_hx_tmp_49 = nil
							}

							e3 = any(map[string]any{"e": _hx_tmp_47, "pmin": _hx_tmp_48, "pmax": _hx_tmp_49, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e4 any
						_ = e4
						e4 = this.VTable.Hx_Field_mapCompr(tmp, a.Get(0))
						var e5 Hx_Enum_loreline_lorscript_exprdef
						_ = e5
						e5 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

						var pmin2 HxNullable[int32]
						_ = pmin2
						pmin2 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax2 HxNullable[int32]
						_ = pmax2
						pmax2 = HxNullable[int32]{}

						var e6 any
						_ = e6
						if e5 == nil {
							e6 = nil
						} else {
							if (pmin2).Valid == false {
								pmin2 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax2).Valid == false {
								pmax2 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_50 any
							_ = _hx_tmp_50
							_hx_tmp_50 = ((any)(e5))
							var hx_nullable_f41 HxNullable[int32]
							_ = hx_nullable_f41
							hx_nullable_f41 = pmin2
							var _hx_tmp_51 any
							_ = _hx_tmp_51
							if (hx_nullable_f41).Valid == true {
								_hx_tmp_51 = (hx_nullable_f41).Value
							} else {
								_hx_tmp_51 = nil
							}

							var hx_nullable_f42 HxNullable[int32]
							_ = hx_nullable_f42
							hx_nullable_f42 = pmax2
							var _hx_tmp_52 any
							_ = _hx_tmp_52
							if (hx_nullable_f42).Valid == true {
								_hx_tmp_52 = (hx_nullable_f42).Value
							} else {
								_hx_tmp_52 = nil
							}

							e6 = any(map[string]any{"e": _hx_tmp_50, "pmin": _hx_tmp_51, "pmax": _hx_tmp_52, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e7 Hx_Enum_loreline_lorscript_exprdef
						_ = e7
						e7 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](e3, e4, e6).Dyn()}

						var pmin3 HxNullable[int32]
						_ = pmin3
						pmin3 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax3 HxNullable[int32]
						_ = pmax3
						pmax3 = HxNullable[int32]{}

						var e8 any
						_ = e8
						if e7 == nil {
							e8 = nil
						} else {
							if (pmin3).Valid == false {
								pmin3 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax3).Valid == false {
								pmax3 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_53 any
							_ = _hx_tmp_53
							_hx_tmp_53 = ((any)(e7))
							var hx_nullable_f43 HxNullable[int32]
							_ = hx_nullable_f43
							hx_nullable_f43 = pmin3
							var _hx_tmp_54 any
							_ = _hx_tmp_54
							if (hx_nullable_f43).Valid == true {
								_hx_tmp_54 = (hx_nullable_f43).Value
							} else {
								_hx_tmp_54 = nil
							}

							var hx_nullable_f44 HxNullable[int32]
							_ = hx_nullable_f44
							hx_nullable_f44 = pmax3
							var _hx_tmp_55 any
							_ = _hx_tmp_55
							if (hx_nullable_f44).Valid == true {
								_hx_tmp_55 = (hx_nullable_f44).Value
							} else {
								_hx_tmp_55 = nil
							}

							e8 = any(map[string]any{"e": _hx_tmp_53, "pmin": _hx_tmp_54, "pmax": _hx_tmp_55, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						return this.VTable.Hx_Field_parseExprNext(e8)
					}

				case 11:
					{
						var _hx_tmp_56 int32
						_ = _hx_tmp_56
						_hx_tmp_56 = this.Hx_Field_uid
						this.Hx_Field_uid = (this.Hx_Field_uid + 1)
						var tmp string
						_ = tmp
						tmp = ("__a_" + Hx_Field_go_haxe_hxdynamic_toString(_hx_tmp_56))
						var e Hx_Enum_loreline_lorscript_exprdef
						_ = e
						e = Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{HxMakeArray[any]().Dyn()}

						var pmin HxNullable[int32]
						_ = pmin
						pmin = HxNullable[int32]{Value: p1, Valid: true}

						var pmax HxNullable[int32]
						_ = pmax
						pmax = HxNullable[int32]{}

						var e1 any
						_ = e1
						if e == nil {
							e1 = nil
						} else {
							if (pmin).Valid == false {
								pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax).Valid == false {
								pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_57 any
							_ = _hx_tmp_57
							_hx_tmp_57 = ((any)(e))
							var hx_nullable_f45 HxNullable[int32]
							_ = hx_nullable_f45
							hx_nullable_f45 = pmin
							var _hx_tmp_58 any
							_ = _hx_tmp_58
							if (hx_nullable_f45).Valid == true {
								_hx_tmp_58 = (hx_nullable_f45).Value
							} else {
								_hx_tmp_58 = nil
							}

							var hx_nullable_f46 HxNullable[int32]
							_ = hx_nullable_f46
							hx_nullable_f46 = pmax
							var _hx_tmp_59 any
							_ = _hx_tmp_59
							if (hx_nullable_f46).Valid == true {
								_hx_tmp_59 = (hx_nullable_f46).Value
							} else {
								_hx_tmp_59 = nil
							}

							e1 = any(map[string]any{"e": _hx_tmp_57, "pmin": _hx_tmp_58, "pmax": _hx_tmp_59, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e2 Hx_Enum_loreline_lorscript_exprdef
						_ = e2
						e2 = Hx_Enum_loreline_lorscript_exprdef_EVar{tmp, nil, e1}

						var pmin1 HxNullable[int32]
						_ = pmin1
						pmin1 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax1 HxNullable[int32]
						_ = pmax1
						pmax1 = HxNullable[int32]{}

						var e3 any
						_ = e3
						if e2 == nil {
							e3 = nil
						} else {
							if (pmin1).Valid == false {
								pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax1).Valid == false {
								pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_60 any
							_ = _hx_tmp_60
							_hx_tmp_60 = ((any)(e2))
							var hx_nullable_f47 HxNullable[int32]
							_ = hx_nullable_f47
							hx_nullable_f47 = pmin1
							var _hx_tmp_61 any
							_ = _hx_tmp_61
							if (hx_nullable_f47).Valid == true {
								_hx_tmp_61 = (hx_nullable_f47).Value
							} else {
								_hx_tmp_61 = nil
							}

							var hx_nullable_f48 HxNullable[int32]
							_ = hx_nullable_f48
							hx_nullable_f48 = pmax1
							var _hx_tmp_62 any
							_ = _hx_tmp_62
							if (hx_nullable_f48).Valid == true {
								_hx_tmp_62 = (hx_nullable_f48).Value
							} else {
								_hx_tmp_62 = nil
							}

							e3 = any(map[string]any{"e": _hx_tmp_60, "pmin": _hx_tmp_61, "pmax": _hx_tmp_62, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e4 any
						_ = e4
						e4 = this.VTable.Hx_Field_mapCompr(tmp, a.Get(0))
						var e5 Hx_Enum_loreline_lorscript_exprdef
						_ = e5
						e5 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

						var pmin2 HxNullable[int32]
						_ = pmin2
						pmin2 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax2 HxNullable[int32]
						_ = pmax2
						pmax2 = HxNullable[int32]{}

						var e6 any
						_ = e6
						if e5 == nil {
							e6 = nil
						} else {
							if (pmin2).Valid == false {
								pmin2 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax2).Valid == false {
								pmax2 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_63 any
							_ = _hx_tmp_63
							_hx_tmp_63 = ((any)(e5))
							var hx_nullable_f49 HxNullable[int32]
							_ = hx_nullable_f49
							hx_nullable_f49 = pmin2
							var _hx_tmp_64 any
							_ = _hx_tmp_64
							if (hx_nullable_f49).Valid == true {
								_hx_tmp_64 = (hx_nullable_f49).Value
							} else {
								_hx_tmp_64 = nil
							}

							var hx_nullable_f50 HxNullable[int32]
							_ = hx_nullable_f50
							hx_nullable_f50 = pmax2
							var _hx_tmp_65 any
							_ = _hx_tmp_65
							if (hx_nullable_f50).Valid == true {
								_hx_tmp_65 = (hx_nullable_f50).Value
							} else {
								_hx_tmp_65 = nil
							}

							e6 = any(map[string]any{"e": _hx_tmp_63, "pmin": _hx_tmp_64, "pmax": _hx_tmp_65, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e7 Hx_Enum_loreline_lorscript_exprdef
						_ = e7
						e7 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](e3, e4, e6).Dyn()}

						var pmin3 HxNullable[int32]
						_ = pmin3
						pmin3 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax3 HxNullable[int32]
						_ = pmax3
						pmax3 = HxNullable[int32]{}

						var e8 any
						_ = e8
						if e7 == nil {
							e8 = nil
						} else {
							if (pmin3).Valid == false {
								pmin3 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax3).Valid == false {
								pmax3 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_66 any
							_ = _hx_tmp_66
							_hx_tmp_66 = ((any)(e7))
							var hx_nullable_f51 HxNullable[int32]
							_ = hx_nullable_f51
							hx_nullable_f51 = pmin3
							var _hx_tmp_67 any
							_ = _hx_tmp_67
							if (hx_nullable_f51).Valid == true {
								_hx_tmp_67 = (hx_nullable_f51).Value
							} else {
								_hx_tmp_67 = nil
							}

							var hx_nullable_f52 HxNullable[int32]
							_ = hx_nullable_f52
							hx_nullable_f52 = pmax3
							var _hx_tmp_68 any
							_ = _hx_tmp_68
							if (hx_nullable_f52).Valid == true {
								_hx_tmp_68 = (hx_nullable_f52).Value
							} else {
								_hx_tmp_68 = nil
							}

							e8 = any(map[string]any{"e": _hx_tmp_66, "pmin": _hx_tmp_67, "pmax": _hx_tmp_68, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						return this.VTable.Hx_Field_parseExprNext(e8)
					}

				case 24:
					{
						var _hx_tmp_69 int32
						_ = _hx_tmp_69
						_hx_tmp_69 = this.Hx_Field_uid
						this.Hx_Field_uid = (this.Hx_Field_uid + 1)
						var tmp string
						_ = tmp
						tmp = ("__a_" + Hx_Field_go_haxe_hxdynamic_toString(_hx_tmp_69))
						var e Hx_Enum_loreline_lorscript_exprdef
						_ = e
						e = Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{HxMakeArray[any]().Dyn()}

						var pmin HxNullable[int32]
						_ = pmin
						pmin = HxNullable[int32]{Value: p1, Valid: true}

						var pmax HxNullable[int32]
						_ = pmax
						pmax = HxNullable[int32]{}

						var e1 any
						_ = e1
						if e == nil {
							e1 = nil
						} else {
							if (pmin).Valid == false {
								pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax).Valid == false {
								pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_70 any
							_ = _hx_tmp_70
							_hx_tmp_70 = ((any)(e))
							var hx_nullable_f53 HxNullable[int32]
							_ = hx_nullable_f53
							hx_nullable_f53 = pmin
							var _hx_tmp_71 any
							_ = _hx_tmp_71
							if (hx_nullable_f53).Valid == true {
								_hx_tmp_71 = (hx_nullable_f53).Value
							} else {
								_hx_tmp_71 = nil
							}

							var hx_nullable_f54 HxNullable[int32]
							_ = hx_nullable_f54
							hx_nullable_f54 = pmax
							var _hx_tmp_72 any
							_ = _hx_tmp_72
							if (hx_nullable_f54).Valid == true {
								_hx_tmp_72 = (hx_nullable_f54).Value
							} else {
								_hx_tmp_72 = nil
							}

							e1 = any(map[string]any{"e": _hx_tmp_70, "pmin": _hx_tmp_71, "pmax": _hx_tmp_72, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e2 Hx_Enum_loreline_lorscript_exprdef
						_ = e2
						e2 = Hx_Enum_loreline_lorscript_exprdef_EVar{tmp, nil, e1}

						var pmin1 HxNullable[int32]
						_ = pmin1
						pmin1 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax1 HxNullable[int32]
						_ = pmax1
						pmax1 = HxNullable[int32]{}

						var e3 any
						_ = e3
						if e2 == nil {
							e3 = nil
						} else {
							if (pmin1).Valid == false {
								pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax1).Valid == false {
								pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_73 any
							_ = _hx_tmp_73
							_hx_tmp_73 = ((any)(e2))
							var hx_nullable_f55 HxNullable[int32]
							_ = hx_nullable_f55
							hx_nullable_f55 = pmin1
							var _hx_tmp_74 any
							_ = _hx_tmp_74
							if (hx_nullable_f55).Valid == true {
								_hx_tmp_74 = (hx_nullable_f55).Value
							} else {
								_hx_tmp_74 = nil
							}

							var hx_nullable_f56 HxNullable[int32]
							_ = hx_nullable_f56
							hx_nullable_f56 = pmax1
							var _hx_tmp_75 any
							_ = _hx_tmp_75
							if (hx_nullable_f56).Valid == true {
								_hx_tmp_75 = (hx_nullable_f56).Value
							} else {
								_hx_tmp_75 = nil
							}

							e3 = any(map[string]any{"e": _hx_tmp_73, "pmin": _hx_tmp_74, "pmax": _hx_tmp_75, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e4 any
						_ = e4
						e4 = this.VTable.Hx_Field_mapCompr(tmp, a.Get(0))
						var e5 Hx_Enum_loreline_lorscript_exprdef
						_ = e5
						e5 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

						var pmin2 HxNullable[int32]
						_ = pmin2
						pmin2 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax2 HxNullable[int32]
						_ = pmax2
						pmax2 = HxNullable[int32]{}

						var e6 any
						_ = e6
						if e5 == nil {
							e6 = nil
						} else {
							if (pmin2).Valid == false {
								pmin2 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax2).Valid == false {
								pmax2 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_76 any
							_ = _hx_tmp_76
							_hx_tmp_76 = ((any)(e5))
							var hx_nullable_f57 HxNullable[int32]
							_ = hx_nullable_f57
							hx_nullable_f57 = pmin2
							var _hx_tmp_77 any
							_ = _hx_tmp_77
							if (hx_nullable_f57).Valid == true {
								_hx_tmp_77 = (hx_nullable_f57).Value
							} else {
								_hx_tmp_77 = nil
							}

							var hx_nullable_f58 HxNullable[int32]
							_ = hx_nullable_f58
							hx_nullable_f58 = pmax2
							var _hx_tmp_78 any
							_ = _hx_tmp_78
							if (hx_nullable_f58).Valid == true {
								_hx_tmp_78 = (hx_nullable_f58).Value
							} else {
								_hx_tmp_78 = nil
							}

							e6 = any(map[string]any{"e": _hx_tmp_76, "pmin": _hx_tmp_77, "pmax": _hx_tmp_78, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						var e7 Hx_Enum_loreline_lorscript_exprdef
						_ = e7
						e7 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](e3, e4, e6).Dyn()}

						var pmin3 HxNullable[int32]
						_ = pmin3
						pmin3 = HxNullable[int32]{Value: p1, Valid: true}

						var pmax3 HxNullable[int32]
						_ = pmax3
						pmax3 = HxNullable[int32]{}

						var e8 any
						_ = e8
						if e7 == nil {
							e8 = nil
						} else {
							if (pmin3).Valid == false {
								pmin3 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
							}

							if (pmax3).Valid == false {
								pmax3 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
							}

							var _hx_tmp_79 any
							_ = _hx_tmp_79
							_hx_tmp_79 = ((any)(e7))
							var hx_nullable_f59 HxNullable[int32]
							_ = hx_nullable_f59
							hx_nullable_f59 = pmin3
							var _hx_tmp_80 any
							_ = _hx_tmp_80
							if (hx_nullable_f59).Valid == true {
								_hx_tmp_80 = (hx_nullable_f59).Value
							} else {
								_hx_tmp_80 = nil
							}

							var hx_nullable_f60 HxNullable[int32]
							_ = hx_nullable_f60
							hx_nullable_f60 = pmax3
							var _hx_tmp_81 any
							_ = _hx_tmp_81
							if (hx_nullable_f60).Valid == true {
								_hx_tmp_81 = (hx_nullable_f60).Value
							} else {
								_hx_tmp_81 = nil
							}

							e8 = any(map[string]any{"e": _hx_tmp_79, "pmin": _hx_tmp_80, "pmax": _hx_tmp_81, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
						}

						return this.VTable.Hx_Field_parseExprNext(e8)
					}

				default:
					{
					}
				}
			}

			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EArrayDecl{a}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			var tmp any
			_ = tmp
			if e == nil {
				tmp = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_82 any
				_ = _hx_tmp_82
				_hx_tmp_82 = ((any)(e))
				var hx_nullable_f61 HxNullable[int32]
				_ = hx_nullable_f61
				hx_nullable_f61 = pmin
				var _hx_tmp_83 any
				_ = _hx_tmp_83
				if (hx_nullable_f61).Valid == true {
					_hx_tmp_83 = (hx_nullable_f61).Value
				} else {
					_hx_tmp_83 = nil
				}

				var hx_nullable_f62 HxNullable[int32]
				_ = hx_nullable_f62
				hx_nullable_f62 = pmax
				var _hx_tmp_84 any
				_ = _hx_tmp_84
				if (hx_nullable_f62).Valid == true {
					_hx_tmp_84 = (hx_nullable_f62).Value
				} else {
					_hx_tmp_84 = nil
				}

				tmp = any(map[string]any{"e": _hx_tmp_82, "pmin": _hx_tmp_83, "pmax": _hx_tmp_84, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(tmp)
		}

	case 16:
		{
			var id string
			_ = id
			id = tk.(Hx_Enum_loreline_lorscript_token_TMeta).Hx_Field_s
			if this.Hx_Field_allowMetadata {
				var args HxArray[any]
				_ = args
				args = this.VTable.Hx_Field_parseMetaArgs()
				var _hx_tmp_85 string
				_ = _hx_tmp_85
				_hx_tmp_85 = id
				var _hx_tmp_86 HxArray[any]
				_ = _hx_tmp_86
				_hx_tmp_86 = args
				var e Hx_Enum_loreline_lorscript_exprdef
				_ = e
				e = Hx_Enum_loreline_lorscript_exprdef_EMeta{_hx_tmp_85, _hx_tmp_86, this.VTable.Hx_Field_parseExpr()}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: p1, Valid: true}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{}

				if e == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_87 any
					_ = _hx_tmp_87
					_hx_tmp_87 = ((any)(e))
					var hx_nullable_f63 HxNullable[int32]
					_ = hx_nullable_f63
					hx_nullable_f63 = pmin
					var _hx_tmp_88 any
					_ = _hx_tmp_88
					if (hx_nullable_f63).Valid == true {
						_hx_tmp_88 = (hx_nullable_f63).Value
					} else {
						_hx_tmp_88 = nil
					}

					var hx_nullable_f64 HxNullable[int32]
					_ = hx_nullable_f64
					hx_nullable_f64 = pmax
					var _hx_tmp_89 any
					_ = _hx_tmp_89
					if (hx_nullable_f64).Valid == true {
						_hx_tmp_89 = (hx_nullable_f64).Value
					} else {
						_hx_tmp_89 = nil
					}

					return any(map[string]any{"e": _hx_tmp_87, "pmin": _hx_tmp_88, "pmax": _hx_tmp_89, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			} else {
				return this.VTable.Hx_Field_unexpected(tk)
			}
		}

	default:
		{
			return this.VTable.Hx_Field_unexpected(tk)
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseLambda(args HxArray[any], pmin HxNullable[int32]) any {
hx_label_0:
	for {
		var id string
		_ = id
		id = this.VTable.Hx_Field_getIdent()
		var _hx_tmp_0 Hx_Enum_loreline_lorscript_ctype
		_ = _hx_tmp_0
		if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TDoubleDot{})) {
			_hx_tmp_0 = this.VTable.Hx_Field_parseType()
		} else {
			_hx_tmp_0 = nil
		}

		var t Hx_Enum_loreline_lorscript_ctype
		_ = t
		t = _hx_tmp_0
		var _hx_tmp_1 HxArray[any]
		_ = _hx_tmp_1
		_hx_tmp_1 = HxMakeArrayView[any](args)
		_ = (Hx_Array_Push[any](_hx_tmp_1, any(map[string]any{"name": ((any)(id)), "t": ((any)(t))})))
		var tk Hx_Enum_loreline_lorscript_token
		_ = tk
		tk = this.VTable.Hx_Field_token()
		switch tk.Hx_Field_enumIndex() {
		case 5:
			{
				break hx_label_0
			}

		case 10:
			{
			}

		default:
			{
				this.VTable.Hx_Field_unexpected(tk)
				break hx_label_0
			}
		}
	}

	{
		var t Hx_Enum_loreline_lorscript_token
		_ = t
		t = this.VTable.Hx_Field_token()
		var _hx_tmp_2 Hx_Enum_loreline_lorscript_token
		_ = _hx_tmp_2
		_hx_tmp_2 = t
		if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
			var _hx_closure_src_311 func(a any, b any) bool
			_ = _hx_closure_src_311
			_hx_closure_src_311 = Hx_Field_type_enumEq
			return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
				return _hx_closure_src_311(((any)(a)), ((any)(b)))
			}
		}()(_hx_tmp_2, Hx_Enum_loreline_lorscript_token_TOp{"->"})) {
			this.VTable.Hx_Field_unexpected(t)
		}
	}

	var eret any
	_ = eret
	eret = this.VTable.Hx_Field_parseExpr()
	var e Hx_Enum_loreline_lorscript_exprdef
	_ = e
	e = Hx_Enum_loreline_lorscript_exprdef_EReturn{eret}

	var pmin1 HxNullable[int32]
	_ = pmin1
	pmin1 = pmin
	var pmax HxNullable[int32]
	_ = pmax
	pmax = HxNullable[int32]{}

	var e1 any
	_ = e1
	if e == nil {
		e1 = nil
	} else {
		if (pmin1).Valid == false {
			pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
		}

		if (pmax).Valid == false {
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
		}

		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = ((any)(e))
		var hx_nullable_f65 HxNullable[int32]
		_ = hx_nullable_f65
		hx_nullable_f65 = pmin1
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		if (hx_nullable_f65).Valid == true {
			_hx_tmp_4 = (hx_nullable_f65).Value
		} else {
			_hx_tmp_4 = nil
		}

		var hx_nullable_f66 HxNullable[int32]
		_ = hx_nullable_f66
		hx_nullable_f66 = pmax
		var _hx_tmp_5 any
		_ = _hx_tmp_5
		if (hx_nullable_f66).Valid == true {
			_hx_tmp_5 = (hx_nullable_f66).Value
		} else {
			_hx_tmp_5 = nil
		}

		e1 = any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
	}

	var _hx_tmp_6 HxArray[any]
	_ = _hx_tmp_6
	_hx_tmp_6 = args
	var _hx_tmp_7 any
	_ = _hx_tmp_7
	_hx_tmp_7 = e1
	var e2 Hx_Enum_loreline_lorscript_exprdef
	_ = e2
	e2 = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_6, _hx_tmp_7, HxNullable[string]{Value: HxStringNull}, nil}

	var pmin2 HxNullable[int32]
	_ = pmin2
	pmin2 = pmin
	var pmax1 HxNullable[int32]
	_ = pmax1
	pmax1 = HxNullable[int32]{}

	if e2 == nil {
		return nil
	} else {
		if (pmin2).Valid == false {
			pmin2 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
		}

		if (pmax1).Valid == false {
			pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
		}

		var _hx_tmp_8 any
		_ = _hx_tmp_8
		_hx_tmp_8 = ((any)(e2))
		var hx_nullable_f67 HxNullable[int32]
		_ = hx_nullable_f67
		hx_nullable_f67 = pmin2
		var _hx_tmp_9 any
		_ = _hx_tmp_9
		if (hx_nullable_f67).Valid == true {
			_hx_tmp_9 = (hx_nullable_f67).Value
		} else {
			_hx_tmp_9 = nil
		}

		var hx_nullable_f68 HxNullable[int32]
		_ = hx_nullable_f68
		hx_nullable_f68 = pmax1
		var _hx_tmp_10 any
		_ = _hx_tmp_10
		if (hx_nullable_f68).Valid == true {
			_hx_tmp_10 = (hx_nullable_f68).Value
		} else {
			_hx_tmp_10 = nil
		}

		return any(map[string]any{"e": _hx_tmp_8, "pmin": _hx_tmp_9, "pmax": _hx_tmp_10, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseMetaArgs() HxArray[any] {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	if (tk != Hx_Enum_loreline_lorscript_token_TPOpen{}) {
		{
			this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
			this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
			this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
		}

		return nil
	}

	var args HxArray[any]
	_ = args
	args = HxMakeArray[any]().Dyn()
	tk = this.VTable.Hx_Field_token()
	if (tk != Hx_Enum_loreline_lorscript_token_TPClose{}) {
		{
			this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
			this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
			this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
		}

	hx_label_0:
		for {
			var _hx_tmp_0 HxArray[any]
			_ = _hx_tmp_0
			_hx_tmp_0 = HxMakeArrayView[any](args)
			_ = (Hx_Array_Push[any](_hx_tmp_0, this.VTable.Hx_Field_parseExpr()))
			{
				var _g Hx_Enum_loreline_lorscript_token
				_ = _g
				_g = this.VTable.Hx_Field_token()
				switch _g.Hx_Field_enumIndex() {
				case 5:
					{
						break hx_label_0
					}

				case 10:
					{
					}

				default:
					{
						var tmp_tk_1 Hx_Enum_loreline_lorscript_token
						_ = tmp_tk_1
						tmp_tk_1 = _g
						this.VTable.Hx_Field_unexpected(tmp_tk_1)
					}
				}
			}
		}
	}

	return args
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_mapCompr(tmp string, e any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		return nil
	}

	var edef Hx_Enum_loreline_lorscript_exprdef
	_ = edef
	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 3:
		{
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
			edef = Hx_Enum_loreline_lorscript_exprdef_EParent{this.VTable.Hx_Field_mapCompr(tmp, e2)}
		}

	case 4:
		{
			var _ge HxArray[any]
			_ = _ge
			_ge = _g.(Hx_Enum_loreline_lorscript_exprdef_EBlock).Hx_Field_e
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(_ge) == 1 {
				var e any
				_ = e
				e = _ge.Get(0)
				edef = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](this.VTable.Hx_Field_mapCompr(tmp, e)).Dyn()}
			} else {
				var e1 Hx_Enum_loreline_lorscript_exprdef
				_ = e1
				e1 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_0 = 0
				} else {
					_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_0, Valid: true}

				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_1 = 0
				} else {
					_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_1, Valid: true}

				var e2 any
				_ = e2
				if e1 == nil {
					e2 = nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = ((any)(e1))
					var hx_nullable_f69 HxNullable[int32]
					_ = hx_nullable_f69
					hx_nullable_f69 = pmin
					var _hx_tmp_3 any
					_ = _hx_tmp_3
					if (hx_nullable_f69).Valid == true {
						_hx_tmp_3 = (hx_nullable_f69).Value
					} else {
						_hx_tmp_3 = nil
					}

					var hx_nullable_f70 HxNullable[int32]
					_ = hx_nullable_f70
					hx_nullable_f70 = pmax
					var _hx_tmp_4 any
					_ = _hx_tmp_4
					if (hx_nullable_f70).Valid == true {
						_hx_tmp_4 = (hx_nullable_f70).Value
					} else {
						_hx_tmp_4 = nil
					}

					e2 = any(map[string]any{"e": _hx_tmp_2, "pmin": _hx_tmp_3, "pmax": _hx_tmp_4, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				var e3 Hx_Enum_loreline_lorscript_exprdef
				_ = e3
				e3 = Hx_Enum_loreline_lorscript_exprdef_EField{e2, "push"}

				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_5 = 0
				} else {
					_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
				}

				var pmin1 HxNullable[int32]
				_ = pmin1
				pmin1 = HxNullable[int32]{Value: _hx_tmp_5, Valid: true}

				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_6 = 0
				} else {
					_hx_tmp_6 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax1 HxNullable[int32]
				_ = pmax1
				pmax1 = HxNullable[int32]{Value: _hx_tmp_6, Valid: true}

				var edef1 any
				_ = edef1
				if e3 == nil {
					edef1 = nil
				} else {
					if (pmin1).Valid == false {
						pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax1).Valid == false {
						pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = ((any)(e3))
					var hx_nullable_f71 HxNullable[int32]
					_ = hx_nullable_f71
					hx_nullable_f71 = pmin1
					var _hx_tmp_8 any
					_ = _hx_tmp_8
					if (hx_nullable_f71).Valid == true {
						_hx_tmp_8 = (hx_nullable_f71).Value
					} else {
						_hx_tmp_8 = nil
					}

					var hx_nullable_f72 HxNullable[int32]
					_ = hx_nullable_f72
					hx_nullable_f72 = pmax1
					var _hx_tmp_9 any
					_ = _hx_tmp_9
					if (hx_nullable_f72).Valid == true {
						_hx_tmp_9 = (hx_nullable_f72).Value
					} else {
						_hx_tmp_9 = nil
					}

					edef1 = any(map[string]any{"e": _hx_tmp_7, "pmin": _hx_tmp_8, "pmax": _hx_tmp_9, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				var _hx_tmp_10 any
				_ = _hx_tmp_10
				_hx_tmp_10 = edef1
				edef = Hx_Enum_loreline_lorscript_exprdef_ECall{_hx_tmp_10, HxMakeArray[any](e).Dyn()}
			}
		}

	case 9:
		{
			var cond any
			_ = cond
			cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_cond
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EIf).Hx_Field_e2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
				var _hx_tmp_11 any
				_ = _hx_tmp_11
				_hx_tmp_11 = cond
				edef = Hx_Enum_loreline_lorscript_exprdef_EIf{_hx_tmp_11, this.VTable.Hx_Field_mapCompr(tmp, e1), nil}
			} else {
				var tmp_e1_1 Hx_Enum_loreline_lorscript_exprdef
				_ = tmp_e1_1
				tmp_e1_1 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_12 = 0
				} else {
					_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_12, Valid: true}

				var _hx_tmp_13 int32
				_ = _hx_tmp_13
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_13 = 0
				} else {
					_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_13, Valid: true}

				var tmp_e2_1 any
				_ = tmp_e2_1
				if tmp_e1_1 == nil {
					tmp_e2_1 = nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_14 any
					_ = _hx_tmp_14
					_hx_tmp_14 = ((any)(tmp_e1_1))
					var hx_nullable_f73 HxNullable[int32]
					_ = hx_nullable_f73
					hx_nullable_f73 = pmin
					var _hx_tmp_15 any
					_ = _hx_tmp_15
					if (hx_nullable_f73).Valid == true {
						_hx_tmp_15 = (hx_nullable_f73).Value
					} else {
						_hx_tmp_15 = nil
					}

					var hx_nullable_f74 HxNullable[int32]
					_ = hx_nullable_f74
					hx_nullable_f74 = pmax
					var _hx_tmp_16 any
					_ = _hx_tmp_16
					if (hx_nullable_f74).Valid == true {
						_hx_tmp_16 = (hx_nullable_f74).Value
					} else {
						_hx_tmp_16 = nil
					}

					tmp_e2_1 = any(map[string]any{"e": _hx_tmp_14, "pmin": _hx_tmp_15, "pmax": _hx_tmp_16, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				var e3 Hx_Enum_loreline_lorscript_exprdef
				_ = e3
				e3 = Hx_Enum_loreline_lorscript_exprdef_EField{tmp_e2_1, "push"}

				var _hx_tmp_17 int32
				_ = _hx_tmp_17
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_17 = 0
				} else {
					_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
				}

				var pmin1 HxNullable[int32]
				_ = pmin1
				pmin1 = HxNullable[int32]{Value: _hx_tmp_17, Valid: true}

				var _hx_tmp_18 int32
				_ = _hx_tmp_18
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_18 = 0
				} else {
					_hx_tmp_18 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax1 HxNullable[int32]
				_ = pmax1
				pmax1 = HxNullable[int32]{Value: _hx_tmp_18, Valid: true}

				var edef1 any
				_ = edef1
				if e3 == nil {
					edef1 = nil
				} else {
					if (pmin1).Valid == false {
						pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax1).Valid == false {
						pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_19 any
					_ = _hx_tmp_19
					_hx_tmp_19 = ((any)(e3))
					var hx_nullable_f75 HxNullable[int32]
					_ = hx_nullable_f75
					hx_nullable_f75 = pmin1
					var _hx_tmp_20 any
					_ = _hx_tmp_20
					if (hx_nullable_f75).Valid == true {
						_hx_tmp_20 = (hx_nullable_f75).Value
					} else {
						_hx_tmp_20 = nil
					}

					var hx_nullable_f76 HxNullable[int32]
					_ = hx_nullable_f76
					hx_nullable_f76 = pmax1
					var _hx_tmp_21 any
					_ = _hx_tmp_21
					if (hx_nullable_f76).Valid == true {
						_hx_tmp_21 = (hx_nullable_f76).Value
					} else {
						_hx_tmp_21 = nil
					}

					edef1 = any(map[string]any{"e": _hx_tmp_19, "pmin": _hx_tmp_20, "pmax": _hx_tmp_21, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				var _hx_tmp_22 any
				_ = _hx_tmp_22
				_hx_tmp_22 = edef1
				edef = Hx_Enum_loreline_lorscript_exprdef_ECall{_hx_tmp_22, HxMakeArray[any](e).Dyn()}
			}
		}

	case 10:
		{
			var cond any
			_ = cond
			cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_cond
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EWhile).Hx_Field_e
			var _hx_tmp_23 any
			_ = _hx_tmp_23
			_hx_tmp_23 = cond
			edef = Hx_Enum_loreline_lorscript_exprdef_EWhile{_hx_tmp_23, this.VTable.Hx_Field_mapCompr(tmp, e2)}
		}

	case 11:
		{
			var v string
			_ = v
			v = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_v
			var it any
			_ = it
			it = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_it
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EFor).Hx_Field_e
			var _hx_tmp_24 string
			_ = _hx_tmp_24
			_hx_tmp_24 = v
			var _hx_tmp_25 any
			_ = _hx_tmp_25
			_hx_tmp_25 = it
			edef = Hx_Enum_loreline_lorscript_exprdef_EFor{_hx_tmp_24, _hx_tmp_25, this.VTable.Hx_Field_mapCompr(tmp, e2)}
		}

	case 24:
		{
			var cond any
			_ = cond
			cond = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_cond
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EDoWhile).Hx_Field_e
			var _hx_tmp_26 any
			_ = _hx_tmp_26
			_hx_tmp_26 = cond
			edef = Hx_Enum_loreline_lorscript_exprdef_EDoWhile{_hx_tmp_26, this.VTable.Hx_Field_mapCompr(tmp, e2)}
		}

	default:
		{
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

			var _hx_tmp_27 int32
			_ = _hx_tmp_27
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_27 = 0
			} else {
				_hx_tmp_27 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_27, Valid: true}

			var _hx_tmp_28 int32
			_ = _hx_tmp_28
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_28 = 0
			} else {
				_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_28, Valid: true}

			var e2 any
			_ = e2
			if e1 == nil {
				e2 = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_29 any
				_ = _hx_tmp_29
				_hx_tmp_29 = ((any)(e1))
				var hx_nullable_f77 HxNullable[int32]
				_ = hx_nullable_f77
				hx_nullable_f77 = pmin
				var _hx_tmp_30 any
				_ = _hx_tmp_30
				if (hx_nullable_f77).Valid == true {
					_hx_tmp_30 = (hx_nullable_f77).Value
				} else {
					_hx_tmp_30 = nil
				}

				var hx_nullable_f78 HxNullable[int32]
				_ = hx_nullable_f78
				hx_nullable_f78 = pmax
				var _hx_tmp_31 any
				_ = _hx_tmp_31
				if (hx_nullable_f78).Valid == true {
					_hx_tmp_31 = (hx_nullable_f78).Value
				} else {
					_hx_tmp_31 = nil
				}

				e2 = any(map[string]any{"e": _hx_tmp_29, "pmin": _hx_tmp_30, "pmax": _hx_tmp_31, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e3 Hx_Enum_loreline_lorscript_exprdef
			_ = e3
			e3 = Hx_Enum_loreline_lorscript_exprdef_EField{e2, "push"}

			var _hx_tmp_32 int32
			_ = _hx_tmp_32
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_32 = 0
			} else {
				_hx_tmp_32 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
			}

			var pmin1 HxNullable[int32]
			_ = pmin1
			pmin1 = HxNullable[int32]{Value: _hx_tmp_32, Valid: true}

			var _hx_tmp_33 int32
			_ = _hx_tmp_33
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_33 = 0
			} else {
				_hx_tmp_33 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax1 HxNullable[int32]
			_ = pmax1
			pmax1 = HxNullable[int32]{Value: _hx_tmp_33, Valid: true}

			var edef1 any
			_ = edef1
			if e3 == nil {
				edef1 = nil
			} else {
				if (pmin1).Valid == false {
					pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax1).Valid == false {
					pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_34 any
				_ = _hx_tmp_34
				_hx_tmp_34 = ((any)(e3))
				var hx_nullable_f79 HxNullable[int32]
				_ = hx_nullable_f79
				hx_nullable_f79 = pmin1
				var _hx_tmp_35 any
				_ = _hx_tmp_35
				if (hx_nullable_f79).Valid == true {
					_hx_tmp_35 = (hx_nullable_f79).Value
				} else {
					_hx_tmp_35 = nil
				}

				var hx_nullable_f80 HxNullable[int32]
				_ = hx_nullable_f80
				hx_nullable_f80 = pmax1
				var _hx_tmp_36 any
				_ = _hx_tmp_36
				if (hx_nullable_f80).Valid == true {
					_hx_tmp_36 = (hx_nullable_f80).Value
				} else {
					_hx_tmp_36 = nil
				}

				edef1 = any(map[string]any{"e": _hx_tmp_34, "pmin": _hx_tmp_35, "pmax": _hx_tmp_36, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var _hx_tmp_37 any
			_ = _hx_tmp_37
			_hx_tmp_37 = edef1
			edef = Hx_Enum_loreline_lorscript_exprdef_ECall{_hx_tmp_37, HxMakeArray[any](e).Dyn()}
		}
	}

	var _hx_tmp_38 int32
	_ = _hx_tmp_38
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		_hx_tmp_38 = 0
	} else {
		_hx_tmp_38 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
	}

	var pmin HxNullable[int32]
	_ = pmin
	pmin = HxNullable[int32]{Value: _hx_tmp_38, Valid: true}

	var _hx_tmp_39 int32
	_ = _hx_tmp_39
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		_hx_tmp_39 = 0
	} else {
		_hx_tmp_39 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
	}

	var pmax HxNullable[int32]
	_ = pmax
	pmax = HxNullable[int32]{Value: _hx_tmp_39, Valid: true}

	if edef == nil {
		return nil
	} else {
		if (pmin).Valid == false {
			pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
		}

		if (pmax).Valid == false {
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
		}

		var _hx_tmp_40 any
		_ = _hx_tmp_40
		_hx_tmp_40 = ((any)(edef))
		var hx_nullable_f81 HxNullable[int32]
		_ = hx_nullable_f81
		hx_nullable_f81 = pmin
		var _hx_tmp_41 any
		_ = _hx_tmp_41
		if (hx_nullable_f81).Valid == true {
			_hx_tmp_41 = (hx_nullable_f81).Value
		} else {
			_hx_tmp_41 = nil
		}

		var hx_nullable_f82 HxNullable[int32]
		_ = hx_nullable_f82
		hx_nullable_f82 = pmax
		var _hx_tmp_42 any
		_ = _hx_tmp_42
		if (hx_nullable_f82).Valid == true {
			_hx_tmp_42 = (hx_nullable_f82).Value
		} else {
			_hx_tmp_42 = nil
		}

		return any(map[string]any{"e": _hx_tmp_40, "pmin": _hx_tmp_41, "pmax": _hx_tmp_42, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_makeUnop(op string, e any) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		_hx_tmp_0 = this.Hx_Field_resumeErrors
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		return nil
	}

	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 6:
		{
			var bop string
			_ = bop
			bop = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = bop
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EBinop{_hx_tmp_1, this.VTable.Hx_Field_makeUnop(op, e1), e2}

			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_2 = 0
			} else {
				_hx_tmp_2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_2, Valid: true}

			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
				_hx_tmp_3 = 0
			} else {
				_hx_tmp_3 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e2, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_3, Valid: true}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = ((any)(e))
				var hx_nullable_f83 HxNullable[int32]
				_ = hx_nullable_f83
				hx_nullable_f83 = pmin
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				if (hx_nullable_f83).Valid == true {
					_hx_tmp_5 = (hx_nullable_f83).Value
				} else {
					_hx_tmp_5 = nil
				}

				var hx_nullable_f84 HxNullable[int32]
				_ = hx_nullable_f84
				hx_nullable_f84 = pmax
				var _hx_tmp_6 any
				_ = _hx_tmp_6
				if (hx_nullable_f84).Valid == true {
					_hx_tmp_6 = (hx_nullable_f84).Value
				} else {
					_hx_tmp_6 = nil
				}

				return any(map[string]any{"e": _hx_tmp_4, "pmin": _hx_tmp_5, "pmax": _hx_tmp_6, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case 22:
		{
			var e1 any
			_ = e1
			e1 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
			var e3 any
			_ = e3
			e3 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_ETernary{this.VTable.Hx_Field_makeUnop(op, e1), e2, e3}

			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_7 = 0
			} else {
				_hx_tmp_7 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_7, Valid: true}

			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e3, nil))) {
				_hx_tmp_8 = 0
			} else {
				_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e3, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_8, Valid: true}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_9 any
				_ = _hx_tmp_9
				_hx_tmp_9 = ((any)(e))
				var hx_nullable_f85 HxNullable[int32]
				_ = hx_nullable_f85
				hx_nullable_f85 = pmin
				var _hx_tmp_10 any
				_ = _hx_tmp_10
				if (hx_nullable_f85).Valid == true {
					_hx_tmp_10 = (hx_nullable_f85).Value
				} else {
					_hx_tmp_10 = nil
				}

				var hx_nullable_f86 HxNullable[int32]
				_ = hx_nullable_f86
				hx_nullable_f86 = pmax
				var _hx_tmp_11 any
				_ = _hx_tmp_11
				if (hx_nullable_f86).Valid == true {
					_hx_tmp_11 = (hx_nullable_f86).Value
				} else {
					_hx_tmp_11 = nil
				}

				return any(map[string]any{"e": _hx_tmp_9, "pmin": _hx_tmp_10, "pmax": _hx_tmp_11, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	default:
		{
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EUnop{op, true, e}

			var _hx_tmp_12 int32
			_ = _hx_tmp_12
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_12 = 0
			} else {
				_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_12, Valid: true}

			var _hx_tmp_13 int32
			_ = _hx_tmp_13
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_13 = 0
			} else {
				_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_13, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_14 any
				_ = _hx_tmp_14
				_hx_tmp_14 = ((any)(e1))
				var hx_nullable_f87 HxNullable[int32]
				_ = hx_nullable_f87
				hx_nullable_f87 = pmin
				var _hx_tmp_15 any
				_ = _hx_tmp_15
				if (hx_nullable_f87).Valid == true {
					_hx_tmp_15 = (hx_nullable_f87).Value
				} else {
					_hx_tmp_15 = nil
				}

				var hx_nullable_f88 HxNullable[int32]
				_ = hx_nullable_f88
				hx_nullable_f88 = pmax
				var _hx_tmp_16 any
				_ = _hx_tmp_16
				if (hx_nullable_f88).Valid == true {
					_hx_tmp_16 = (hx_nullable_f88).Value
				} else {
					_hx_tmp_16 = nil
				}

				return any(map[string]any{"e": _hx_tmp_14, "pmin": _hx_tmp_15, "pmax": _hx_tmp_16, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_makeBinop(op string, e1 any, e any) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
		_hx_tmp_0 = this.Hx_Field_resumeErrors
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var e2 Hx_Enum_loreline_lorscript_exprdef
		_ = e2
		e2 = Hx_Enum_loreline_lorscript_exprdef_EBinop{op, e1, e}

		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
			_hx_tmp_1 = 0
		} else {
			_hx_tmp_1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
		}

		var pmin HxNullable[int32]
		_ = pmin
		pmin = HxNullable[int32]{Value: _hx_tmp_1, Valid: true}

		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
			_hx_tmp_2 = 0
		} else {
			_hx_tmp_2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
		}

		var pmax HxNullable[int32]
		_ = pmax
		pmax = HxNullable[int32]{Value: _hx_tmp_2, Valid: true}

		if e2 == nil {
			return nil
		} else {
			if (pmin).Valid == false {
				pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
			}

			if (pmax).Valid == false {
				pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
			}

			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = ((any)(e2))
			var hx_nullable_f89 HxNullable[int32]
			_ = hx_nullable_f89
			hx_nullable_f89 = pmin
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			if (hx_nullable_f89).Valid == true {
				_hx_tmp_4 = (hx_nullable_f89).Value
			} else {
				_hx_tmp_4 = nil
			}

			var hx_nullable_f90 HxNullable[int32]
			_ = hx_nullable_f90
			hx_nullable_f90 = pmax
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			if (hx_nullable_f90).Valid == true {
				_hx_tmp_5 = (hx_nullable_f90).Value
			} else {
				_hx_tmp_5 = nil
			}

			return any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
		}
	}

	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 6:
		{
			var op2 string
			_ = op2
			op2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
			var e3 any
			_ = e3
			e3 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			var hx_cmp_6 HxNullable[int32]
			_ = hx_cmp_6
			hx_cmp_6 = func() func(_hx_param_key string) HxNullable[int32] {
				var _hx_closure_src_316 func(_hx_param_key any) any
				_ = _hx_closure_src_316
				_hx_closure_src_316 = this.Hx_Field_opPriority.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) HxNullable[int32] {
					var hx_nullable_t40 any
					_ = hx_nullable_t40
					hx_nullable_t40 = _hx_closure_src_316(((any)(_hx_param_key)))
					var _hx_tmp_7 HxNullable[int32]
					_ = _hx_tmp_7
					if hx_nullable_t40 == nil {
						_hx_tmp_7 = HxNullable[int32]{}
					} else {
						_hx_tmp_7 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t40))), Valid: true}
					}

					return _hx_tmp_7
				}
			}()(op)
			var hx_cmp_7 HxNullable[int32]
			_ = hx_cmp_7
			hx_cmp_7 = func() func(_hx_param_key string) HxNullable[int32] {
				var _hx_closure_src_317 func(_hx_param_key any) any
				_ = _hx_closure_src_317
				_hx_closure_src_317 = this.Hx_Field_opPriority.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) HxNullable[int32] {
					var hx_nullable_t41 any
					_ = hx_nullable_t41
					hx_nullable_t41 = _hx_closure_src_317(((any)(_hx_param_key)))
					var _hx_tmp_8 HxNullable[int32]
					_ = _hx_tmp_8
					if hx_nullable_t41 == nil {
						_hx_tmp_8 = HxNullable[int32]{}
					} else {
						_hx_tmp_8 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t41))), Valid: true}
					}

					return _hx_tmp_8
				}
			}()(op2)
			if (hx_cmp_6).Valid && (hx_cmp_7).Valid && ((hx_cmp_6).Value <= (hx_cmp_7).Value) || (!(hx_cmp_6).Valid && !(hx_cmp_7).Valid) {
				_hx_tmp_6 = !func() func(_hx_param_key string) bool {
					var _hx_closure_src_314 func(_hx_param_key any) bool
					_ = _hx_closure_src_314
					_hx_closure_src_314 = this.Hx_Field_opRightAssoc.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_314(((any)(_hx_param_key)))
					}
				}()(op)
			} else {
				_hx_tmp_6 = false
			}

			if _hx_tmp_6 {
				var _hx_tmp_9 string
				_ = _hx_tmp_9
				_hx_tmp_9 = op2
				var e Hx_Enum_loreline_lorscript_exprdef
				_ = e
				e = Hx_Enum_loreline_lorscript_exprdef_EBinop{_hx_tmp_9, this.VTable.Hx_Field_makeBinop(op, e1, e2), e3}

				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
					_hx_tmp_10 = 0
				} else {
					_hx_tmp_10 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_10, Valid: true}

				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e3, nil))) {
					_hx_tmp_11 = 0
				} else {
					_hx_tmp_11 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e3, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_11, Valid: true}

				if e == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_12 any
					_ = _hx_tmp_12
					_hx_tmp_12 = ((any)(e))
					var hx_nullable_f91 HxNullable[int32]
					_ = hx_nullable_f91
					hx_nullable_f91 = pmin
					var _hx_tmp_13 any
					_ = _hx_tmp_13
					if (hx_nullable_f91).Valid == true {
						_hx_tmp_13 = (hx_nullable_f91).Value
					} else {
						_hx_tmp_13 = nil
					}

					var hx_nullable_f92 HxNullable[int32]
					_ = hx_nullable_f92
					hx_nullable_f92 = pmax
					var _hx_tmp_14 any
					_ = _hx_tmp_14
					if (hx_nullable_f92).Valid == true {
						_hx_tmp_14 = (hx_nullable_f92).Value
					} else {
						_hx_tmp_14 = nil
					}

					return any(map[string]any{"e": _hx_tmp_12, "pmin": _hx_tmp_13, "pmax": _hx_tmp_14, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			} else {
				var tmp_e2_1 Hx_Enum_loreline_lorscript_exprdef
				_ = tmp_e2_1
				tmp_e2_1 = Hx_Enum_loreline_lorscript_exprdef_EBinop{op, e1, e}

				var _hx_tmp_15 int32
				_ = _hx_tmp_15
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
					_hx_tmp_15 = 0
				} else {
					_hx_tmp_15 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_15, Valid: true}

				var _hx_tmp_16 int32
				_ = _hx_tmp_16
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_16 = 0
				} else {
					_hx_tmp_16 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_16, Valid: true}

				if tmp_e2_1 == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_17 any
					_ = _hx_tmp_17
					_hx_tmp_17 = ((any)(tmp_e2_1))
					var hx_nullable_f93 HxNullable[int32]
					_ = hx_nullable_f93
					hx_nullable_f93 = pmin
					var _hx_tmp_18 any
					_ = _hx_tmp_18
					if (hx_nullable_f93).Valid == true {
						_hx_tmp_18 = (hx_nullable_f93).Value
					} else {
						_hx_tmp_18 = nil
					}

					var hx_nullable_f94 HxNullable[int32]
					_ = hx_nullable_f94
					hx_nullable_f94 = pmax
					var _hx_tmp_19 any
					_ = _hx_tmp_19
					if (hx_nullable_f94).Valid == true {
						_hx_tmp_19 = (hx_nullable_f94).Value
					} else {
						_hx_tmp_19 = nil
					}

					return any(map[string]any{"e": _hx_tmp_17, "pmin": _hx_tmp_18, "pmax": _hx_tmp_19, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			}
		}

	case 22:
		{
			var e2 any
			_ = e2
			e2 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_cond
			var e3 any
			_ = e3
			e3 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e1
			var e4 any
			_ = e4
			e4 = _g.(Hx_Enum_loreline_lorscript_exprdef_ETernary).Hx_Field_e2
			if func() func(_hx_param_key string) bool {
				var _hx_closure_src_315 func(_hx_param_key any) bool
				_ = _hx_closure_src_315
				_hx_closure_src_315 = this.Hx_Field_opRightAssoc.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_315(((any)(_hx_param_key)))
				}
			}()(op) {
				var tmp_e2_1 Hx_Enum_loreline_lorscript_exprdef
				_ = tmp_e2_1
				tmp_e2_1 = Hx_Enum_loreline_lorscript_exprdef_EBinop{op, e1, e}

				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
					_hx_tmp_20 = 0
				} else {
					_hx_tmp_20 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_20, Valid: true}

				var _hx_tmp_21 int32
				_ = _hx_tmp_21
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_21 = 0
				} else {
					_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_21, Valid: true}

				if tmp_e2_1 == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_22 any
					_ = _hx_tmp_22
					_hx_tmp_22 = ((any)(tmp_e2_1))
					var hx_nullable_f95 HxNullable[int32]
					_ = hx_nullable_f95
					hx_nullable_f95 = pmin
					var _hx_tmp_23 any
					_ = _hx_tmp_23
					if (hx_nullable_f95).Valid == true {
						_hx_tmp_23 = (hx_nullable_f95).Value
					} else {
						_hx_tmp_23 = nil
					}

					var hx_nullable_f96 HxNullable[int32]
					_ = hx_nullable_f96
					hx_nullable_f96 = pmax
					var _hx_tmp_24 any
					_ = _hx_tmp_24
					if (hx_nullable_f96).Valid == true {
						_hx_tmp_24 = (hx_nullable_f96).Value
					} else {
						_hx_tmp_24 = nil
					}

					return any(map[string]any{"e": _hx_tmp_22, "pmin": _hx_tmp_23, "pmax": _hx_tmp_24, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			} else {
				var e5 Hx_Enum_loreline_lorscript_exprdef
				_ = e5
				e5 = Hx_Enum_loreline_lorscript_exprdef_ETernary{this.VTable.Hx_Field_makeBinop(op, e1, e2), e3, e4}

				var _hx_tmp_25 int32
				_ = _hx_tmp_25
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
					_hx_tmp_25 = 0
				} else {
					_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_25, Valid: true}

				var _hx_tmp_26 int32
				_ = _hx_tmp_26
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_26 = 0
				} else {
					_hx_tmp_26 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: _hx_tmp_26, Valid: true}

				if e5 == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_27 any
					_ = _hx_tmp_27
					_hx_tmp_27 = ((any)(e5))
					var hx_nullable_f97 HxNullable[int32]
					_ = hx_nullable_f97
					hx_nullable_f97 = pmin
					var _hx_tmp_28 any
					_ = _hx_tmp_28
					if (hx_nullable_f97).Valid == true {
						_hx_tmp_28 = (hx_nullable_f97).Value
					} else {
						_hx_tmp_28 = nil
					}

					var hx_nullable_f98 HxNullable[int32]
					_ = hx_nullable_f98
					hx_nullable_f98 = pmax
					var _hx_tmp_29 any
					_ = _hx_tmp_29
					if (hx_nullable_f98).Valid == true {
						_hx_tmp_29 = (hx_nullable_f98).Value
					} else {
						_hx_tmp_29 = nil
					}

					return any(map[string]any{"e": _hx_tmp_27, "pmin": _hx_tmp_28, "pmax": _hx_tmp_29, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			}
		}

	default:
		{
			var e2 Hx_Enum_loreline_lorscript_exprdef
			_ = e2
			e2 = Hx_Enum_loreline_lorscript_exprdef_EBinop{op, e1, e}

			var _hx_tmp_30 int32
			_ = _hx_tmp_30
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_30 = 0
			} else {
				_hx_tmp_30 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_30, Valid: true}

			var _hx_tmp_31 int32
			_ = _hx_tmp_31
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_31 = 0
			} else {
				_hx_tmp_31 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_31, Valid: true}

			if e2 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_32 any
				_ = _hx_tmp_32
				_hx_tmp_32 = ((any)(e2))
				var hx_nullable_f99 HxNullable[int32]
				_ = hx_nullable_f99
				hx_nullable_f99 = pmin
				var _hx_tmp_33 any
				_ = _hx_tmp_33
				if (hx_nullable_f99).Valid == true {
					_hx_tmp_33 = (hx_nullable_f99).Value
				} else {
					_hx_tmp_33 = nil
				}

				var hx_nullable_f100 HxNullable[int32]
				_ = hx_nullable_f100
				hx_nullable_f100 = pmax
				var _hx_tmp_34 any
				_ = _hx_tmp_34
				if (hx_nullable_f100).Valid == true {
					_hx_tmp_34 = (hx_nullable_f100).Value
				} else {
					_hx_tmp_34 = nil
				}

				return any(map[string]any{"e": _hx_tmp_32, "pmin": _hx_tmp_33, "pmax": _hx_tmp_34, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseStructure(id string) any {
	var p1 int32
	_ = p1
	p1 = this.Hx_Field_tokenMin
	switch id {
	case "break":
		{
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EBreak{}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_0 any
				_ = _hx_tmp_0
				_hx_tmp_0 = ((any)(e))
				var hx_nullable_f101 HxNullable[int32]
				_ = hx_nullable_f101
				hx_nullable_f101 = pmin
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				if (hx_nullable_f101).Valid == true {
					_hx_tmp_1 = (hx_nullable_f101).Value
				} else {
					_hx_tmp_1 = nil
				}

				var hx_nullable_f102 HxNullable[int32]
				_ = hx_nullable_f102
				hx_nullable_f102 = pmax
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				if (hx_nullable_f102).Valid == true {
					_hx_tmp_2 = (hx_nullable_f102).Value
				} else {
					_hx_tmp_2 = nil
				}

				return any(map[string]any{"e": _hx_tmp_0, "pmin": _hx_tmp_1, "pmax": _hx_tmp_2, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "continue":
		{
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EContinue{}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = ((any)(e))
				var hx_nullable_f103 HxNullable[int32]
				_ = hx_nullable_f103
				hx_nullable_f103 = pmin
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				if (hx_nullable_f103).Valid == true {
					_hx_tmp_4 = (hx_nullable_f103).Value
				} else {
					_hx_tmp_4 = nil
				}

				var hx_nullable_f104 HxNullable[int32]
				_ = hx_nullable_f104
				hx_nullable_f104 = pmax
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				if (hx_nullable_f104).Valid == true {
					_hx_tmp_5 = (hx_nullable_f104).Value
				} else {
					_hx_tmp_5 = nil
				}

				return any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "do":
		{
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			if tk.Hx_Field_enumIndex() == 2 {
				if tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s != "while" {
					this.VTable.Hx_Field_unexpected(tk)
				}
			} else {
				this.VTable.Hx_Field_unexpected(tk)
			}

			var econd any
			_ = econd
			econd = this.VTable.Hx_Field_parseExpr()
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EDoWhile{econd, e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(econd, nil))) {
				_hx_tmp_6 = 0
			} else {
				_hx_tmp_6 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(econd, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_6, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_7 any
				_ = _hx_tmp_7
				_hx_tmp_7 = ((any)(e1))
				var hx_nullable_f105 HxNullable[int32]
				_ = hx_nullable_f105
				hx_nullable_f105 = pmin
				var _hx_tmp_8 any
				_ = _hx_tmp_8
				if (hx_nullable_f105).Valid == true {
					_hx_tmp_8 = (hx_nullable_f105).Value
				} else {
					_hx_tmp_8 = nil
				}

				var hx_nullable_f106 HxNullable[int32]
				_ = hx_nullable_f106
				hx_nullable_f106 = pmax
				var _hx_tmp_9 any
				_ = _hx_tmp_9
				if (hx_nullable_f106).Valid == true {
					_hx_tmp_9 = (hx_nullable_f106).Value
				} else {
					_hx_tmp_9 = nil
				}

				return any(map[string]any{"e": _hx_tmp_7, "pmin": _hx_tmp_8, "pmax": _hx_tmp_9, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "else":
		{
			return this.VTable.Hx_Field_unexpected(Hx_Enum_loreline_lorscript_token_TId{id})
		}

	case "for":
		{
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TPOpen{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var vname string
			_ = vname
			vname = this.VTable.Hx_Field_getIdent()
			{
				var t1 Hx_Enum_loreline_lorscript_token
				_ = t1
				t1 = this.VTable.Hx_Field_token()
				var _hx_tmp_10 Hx_Enum_loreline_lorscript_token
				_ = _hx_tmp_10
				_hx_tmp_10 = t1
				if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
					var _hx_closure_src_318 func(a any, b any) bool
					_ = _hx_closure_src_318
					_hx_closure_src_318 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
						return _hx_closure_src_318(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_10, Hx_Enum_loreline_lorscript_token_TId{"in"})) {
					this.VTable.Hx_Field_unexpected(t1)
				}
			}

			var eiter any
			_ = eiter
			eiter = this.VTable.Hx_Field_parseExpr()
			{
				var t2 Hx_Enum_loreline_lorscript_token
				_ = t2
				t2 = this.VTable.Hx_Field_token()
				if (t2 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
					this.VTable.Hx_Field_unexpected(t2)
				}
			}

			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EFor{vname, eiter, e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_11 int32
			_ = _hx_tmp_11
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_11 = 0
			} else {
				_hx_tmp_11 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_11, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_12 any
				_ = _hx_tmp_12
				_hx_tmp_12 = ((any)(e1))
				var hx_nullable_f107 HxNullable[int32]
				_ = hx_nullable_f107
				hx_nullable_f107 = pmin
				var _hx_tmp_13 any
				_ = _hx_tmp_13
				if (hx_nullable_f107).Valid == true {
					_hx_tmp_13 = (hx_nullable_f107).Value
				} else {
					_hx_tmp_13 = nil
				}

				var hx_nullable_f108 HxNullable[int32]
				_ = hx_nullable_f108
				hx_nullable_f108 = pmax
				var _hx_tmp_14 any
				_ = _hx_tmp_14
				if (hx_nullable_f108).Valid == true {
					_hx_tmp_14 = (hx_nullable_f108).Value
				} else {
					_hx_tmp_14 = nil
				}

				return any(map[string]any{"e": _hx_tmp_12, "pmin": _hx_tmp_13, "pmax": _hx_tmp_14, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "function":
		{
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			var name HxNullable[string]
			_ = name
			name = HxNullable[string]{Value: HxStringNull}

			if tk.Hx_Field_enumIndex() == 2 {
				var id string
				_ = id
				id = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
				name = HxNullableString(id)
			} else {
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var inf any
			_ = inf
			inf = this.VTable.Hx_Field_parseFunctionDecl()
			var _hx_tmp_15 HxArray[any]
			_ = _hx_tmp_15
			_hx_tmp_15 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(inf, "args"))
			var _hx_tmp_16 any
			_ = _hx_tmp_16
			_hx_tmp_16 = Hx_Field_go_haxe_hxdynamic_getField(inf, "body")
			var _hx_tmp_17 HxNullable[string]
			_ = _hx_tmp_17
			_hx_tmp_17 = name
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_15, _hx_tmp_16, _hx_tmp_17, func() Hx_Enum_loreline_lorscript_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(inf, "ret"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
			}()}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var e1 any
			_ = e1
			e1 = Hx_Field_go_haxe_hxdynamic_getField(inf, "body")
			var _hx_tmp_18 int32
			_ = _hx_tmp_18
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_18 = 0
			} else {
				_hx_tmp_18 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_18, Valid: true}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_19 any
				_ = _hx_tmp_19
				_hx_tmp_19 = ((any)(e))
				var hx_nullable_f109 HxNullable[int32]
				_ = hx_nullable_f109
				hx_nullable_f109 = pmin
				var _hx_tmp_20 any
				_ = _hx_tmp_20
				if (hx_nullable_f109).Valid == true {
					_hx_tmp_20 = (hx_nullable_f109).Value
				} else {
					_hx_tmp_20 = nil
				}

				var hx_nullable_f110 HxNullable[int32]
				_ = hx_nullable_f110
				hx_nullable_f110 = pmax
				var _hx_tmp_21 any
				_ = _hx_tmp_21
				if (hx_nullable_f110).Valid == true {
					_hx_tmp_21 = (hx_nullable_f110).Value
				} else {
					_hx_tmp_21 = nil
				}

				return any(map[string]any{"e": _hx_tmp_19, "pmin": _hx_tmp_20, "pmax": _hx_tmp_21, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "if":
		{
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TPOpen{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var cond any
			_ = cond
			cond = this.VTable.Hx_Field_parseExpr()
			{
				var t1 Hx_Enum_loreline_lorscript_token
				_ = t1
				t1 = this.VTable.Hx_Field_token()
				if (t1 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
					this.VTable.Hx_Field_unexpected(t1)
				}
			}

			var e1 any
			_ = e1
			e1 = this.VTable.Hx_Field_parseExpr()
			var e2 any
			_ = e2
			var semic bool
			_ = semic
			semic = false
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			if (tk == Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
				semic = true
				tk = this.VTable.Hx_Field_token()
			}

			var _hx_tmp_22 Hx_Enum_loreline_lorscript_token
			_ = _hx_tmp_22
			_hx_tmp_22 = tk
			if (func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
				var _hx_closure_src_319 func(a any, b any) bool
				_ = _hx_closure_src_319
				_hx_closure_src_319 = Hx_Field_type_enumEq
				return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
					return _hx_closure_src_319(((any)(a)), ((any)(b)))
				}
			}()(_hx_tmp_22, Hx_Enum_loreline_lorscript_token_TId{"else"})) {
				e2 = this.VTable.Hx_Field_parseExpr()
			} else {
				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}

				if semic {
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(Hx_Enum_loreline_lorscript_token_TSemicolon{})), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}
			}

			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EIf{cond, e1, e2}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_23 HxNullable[int32]
			_ = _hx_tmp_23
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
				_hx_tmp_23 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
			} else {
				var _hx_tmp_24 int32
				_ = _hx_tmp_24
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
					_hx_tmp_24 = 0
				} else {
					_hx_tmp_24 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e2, "pmax"))))
				}

				_hx_tmp_23 = HxNullable[int32]{Value: _hx_tmp_24, Valid: true}
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = _hx_tmp_23
			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_25 any
				_ = _hx_tmp_25
				_hx_tmp_25 = ((any)(e))
				var hx_nullable_f111 HxNullable[int32]
				_ = hx_nullable_f111
				hx_nullable_f111 = pmin
				var _hx_tmp_26 any
				_ = _hx_tmp_26
				if (hx_nullable_f111).Valid == true {
					_hx_tmp_26 = (hx_nullable_f111).Value
				} else {
					_hx_tmp_26 = nil
				}

				var hx_nullable_f112 HxNullable[int32]
				_ = hx_nullable_f112
				hx_nullable_f112 = pmax
				var _hx_tmp_27 any
				_ = _hx_tmp_27
				if (hx_nullable_f112).Valid == true {
					_hx_tmp_27 = (hx_nullable_f112).Value
				} else {
					_hx_tmp_27 = nil
				}

				return any(map[string]any{"e": _hx_tmp_25, "pmin": _hx_tmp_26, "pmax": _hx_tmp_27, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "inline":
		{
			if (!this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TId{"function"})) {
				this.VTable.Hx_Field_unexpected(Hx_Enum_loreline_lorscript_token_TId{"inline"})
			}

			return this.VTable.Hx_Field_parseStructure("function")
		}

	case "new":
		{
			var a HxArray[string]
			_ = a
			a = HxMakeArray[string]()
			var _hx_tmp_28 HxArray[string]
			_ = _hx_tmp_28
			_hx_tmp_28 = a
			_ = (Hx_Array_Push[string](_hx_tmp_28, this.VTable.Hx_Field_getIdent()))
		hx_label_0:
			for {
				var tk Hx_Enum_loreline_lorscript_token
				_ = tk
				tk = this.VTable.Hx_Field_token()
				switch tk.Hx_Field_enumIndex() {
				case 4:
					{
						break hx_label_0
					}

				case 8:
					{
						var _hx_tmp_29 HxArray[string]
						_ = _hx_tmp_29
						_hx_tmp_29 = a
						_ = (Hx_Array_Push[string](_hx_tmp_29, this.VTable.Hx_Field_getIdent()))
					}

				default:
					{
						this.VTable.Hx_Field_unexpected(tk)
						break hx_label_0
					}
				}
			}

			var args HxArray[any]
			_ = args
			args = this.VTable.Hx_Field_parseExprList(Hx_Enum_loreline_lorscript_token_TPClose{})
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_ENew{(Hx_Array_Join[string](a, ".")), args}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_30 any
				_ = _hx_tmp_30
				_hx_tmp_30 = ((any)(e))
				var hx_nullable_f113 HxNullable[int32]
				_ = hx_nullable_f113
				hx_nullable_f113 = pmin
				var _hx_tmp_31 any
				_ = _hx_tmp_31
				if (hx_nullable_f113).Valid == true {
					_hx_tmp_31 = (hx_nullable_f113).Value
				} else {
					_hx_tmp_31 = nil
				}

				var hx_nullable_f114 HxNullable[int32]
				_ = hx_nullable_f114
				hx_nullable_f114 = pmax
				var _hx_tmp_32 any
				_ = _hx_tmp_32
				if (hx_nullable_f114).Valid == true {
					_hx_tmp_32 = (hx_nullable_f114).Value
				} else {
					_hx_tmp_32 = nil
				}

				return any(map[string]any{"e": _hx_tmp_30, "pmin": _hx_tmp_31, "pmax": _hx_tmp_32, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "return":
		{
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var _hx_tmp_33 any
			_ = _hx_tmp_33
			if (tk == Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
				_hx_tmp_33 = nil
			} else {
				_hx_tmp_33 = this.VTable.Hx_Field_parseExpr()
			}

			var e any
			_ = e
			e = _hx_tmp_33
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EReturn{e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_34 HxNullable[int32]
			_ = _hx_tmp_34
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_34 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
			} else {
				var _hx_tmp_35 int32
				_ = _hx_tmp_35
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_35 = 0
				} else {
					_hx_tmp_35 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				_hx_tmp_34 = HxNullable[int32]{Value: _hx_tmp_35, Valid: true}
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = _hx_tmp_34
			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_36 any
				_ = _hx_tmp_36
				_hx_tmp_36 = ((any)(e1))
				var hx_nullable_f115 HxNullable[int32]
				_ = hx_nullable_f115
				hx_nullable_f115 = pmin
				var _hx_tmp_37 any
				_ = _hx_tmp_37
				if (hx_nullable_f115).Valid == true {
					_hx_tmp_37 = (hx_nullable_f115).Value
				} else {
					_hx_tmp_37 = nil
				}

				var hx_nullable_f116 HxNullable[int32]
				_ = hx_nullable_f116
				hx_nullable_f116 = pmax
				var _hx_tmp_38 any
				_ = _hx_tmp_38
				if (hx_nullable_f116).Valid == true {
					_hx_tmp_38 = (hx_nullable_f116).Value
				} else {
					_hx_tmp_38 = nil
				}

				return any(map[string]any{"e": _hx_tmp_36, "pmin": _hx_tmp_37, "pmax": _hx_tmp_38, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "switch":
		{
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			var def any
			_ = def
			var cases HxArray[any]
			_ = cases
			cases = HxMakeArray[any]().Dyn()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TBrOpen{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

		hx_label_4:
			for {
				var tk Hx_Enum_loreline_lorscript_token
				_ = tk
				tk = this.VTable.Hx_Field_token()
				switch tk.Hx_Field_enumIndex() {
				case 2:
					{
						switch tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s {
						case "case":
							{
								var c any
								_ = c
								c = any(map[string]any{"values": ((any)(HxMakeArray[any]().Dyn())), "expr": nil})
								_ = (Hx_Array_Push[any](HxMakeArrayView[any](cases), c))
							hx_label_1:
								for {
									var tmp_e_1 any
									_ = tmp_e_1
									tmp_e_1 = this.VTable.Hx_Field_parseExpr()
									var _hx_tmp_39 any
									_ = _hx_tmp_39
									_hx_tmp_39 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "values"), "push")
									Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_39, HxMakeArray[any](tmp_e_1).Dyn())
									tk = this.VTable.Hx_Field_token()
									switch tk.Hx_Field_enumIndex() {
									case 10:
										{
										}

									case 15:
										{
											break hx_label_1
										}

									default:
										{
											this.VTable.Hx_Field_unexpected(tk)
											break hx_label_1
										}
									}
								}

								var exprs HxArray[any]
								_ = exprs
								exprs = HxMakeArray[any]().Dyn()
							hx_label_2:
								for {
									tk = this.VTable.Hx_Field_token()
									{
										this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
										this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
										this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
									}

									switch tk.Hx_Field_enumIndex() {
									case 0:
										{
											if this.Hx_Field_resumeErrors {
												break hx_label_2
											} else {
												this.VTable.Hx_Field_parseFullExpr(exprs)
											}
										}

									case 2:
										{
											switch tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s {
											case "case", "default":
												{
													break hx_label_2
												}

											default:
												{
													this.VTable.Hx_Field_parseFullExpr(exprs)
												}
											}
										}

									case 7:
										{
											break hx_label_2
										}

									default:
										{
											this.VTable.Hx_Field_parseFullExpr(exprs)
										}
									}
								}

								var tmp any
								_ = tmp
								if Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) == 1 {
									tmp = exprs.Get(0)
								} else {
									if Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) == 0 {
										var tmp_e_1 Hx_Enum_loreline_lorscript_exprdef
										_ = tmp_e_1
										tmp_e_1 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any]().Dyn()}

										var pmin HxNullable[int32]
										_ = pmin
										pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

										var pmax HxNullable[int32]
										_ = pmax
										pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

										if tmp_e_1 == nil {
											tmp = nil
										} else {
											if (pmin).Valid == false {
												pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
											}

											if (pmax).Valid == false {
												pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
											}

											var _hx_tmp_40 any
											_ = _hx_tmp_40
											_hx_tmp_40 = ((any)(tmp_e_1))
											var hx_nullable_f117 HxNullable[int32]
											_ = hx_nullable_f117
											hx_nullable_f117 = pmin
											var _hx_tmp_41 any
											_ = _hx_tmp_41
											if (hx_nullable_f117).Valid == true {
												_hx_tmp_41 = (hx_nullable_f117).Value
											} else {
												_hx_tmp_41 = nil
											}

											var hx_nullable_f118 HxNullable[int32]
											_ = hx_nullable_f118
											hx_nullable_f118 = pmax
											var _hx_tmp_42 any
											_ = _hx_tmp_42
											if (hx_nullable_f118).Valid == true {
												_hx_tmp_42 = (hx_nullable_f118).Value
											} else {
												_hx_tmp_42 = nil
											}

											tmp = any(map[string]any{"e": _hx_tmp_40, "pmin": _hx_tmp_41, "pmax": _hx_tmp_42, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
										}
									} else {
										var tmp_e_1 Hx_Enum_loreline_lorscript_exprdef
										_ = tmp_e_1
										tmp_e_1 = Hx_Enum_loreline_lorscript_exprdef_EBlock{exprs}

										var e1 any
										_ = e1
										e1 = exprs.Get(0)
										var _hx_tmp_43 int32
										_ = _hx_tmp_43
										if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
											_hx_tmp_43 = 0
										} else {
											_hx_tmp_43 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
										}

										var pmin HxNullable[int32]
										_ = pmin
										pmin = HxNullable[int32]{Value: _hx_tmp_43, Valid: true}

										var _hx_tmp_44 HxArray[any]
										_ = _hx_tmp_44
										_hx_tmp_44 = exprs
										var e2 any
										_ = e2
										e2 = _hx_tmp_44.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) - 1))
										var _hx_tmp_45 int32
										_ = _hx_tmp_45
										if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
											_hx_tmp_45 = 0
										} else {
											_hx_tmp_45 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e2, "pmax"))))
										}

										var pmax HxNullable[int32]
										_ = pmax
										pmax = HxNullable[int32]{Value: _hx_tmp_45, Valid: true}

										if tmp_e_1 == nil {
											tmp = nil
										} else {
											if (pmin).Valid == false {
												pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
											}

											if (pmax).Valid == false {
												pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
											}

											var _hx_tmp_46 any
											_ = _hx_tmp_46
											_hx_tmp_46 = ((any)(tmp_e_1))
											var hx_nullable_f119 HxNullable[int32]
											_ = hx_nullable_f119
											hx_nullable_f119 = pmin
											var _hx_tmp_47 any
											_ = _hx_tmp_47
											if (hx_nullable_f119).Valid == true {
												_hx_tmp_47 = (hx_nullable_f119).Value
											} else {
												_hx_tmp_47 = nil
											}

											var hx_nullable_f120 HxNullable[int32]
											_ = hx_nullable_f120
											hx_nullable_f120 = pmax
											var _hx_tmp_48 any
											_ = _hx_tmp_48
											if (hx_nullable_f120).Valid == true {
												_hx_tmp_48 = (hx_nullable_f120).Value
											} else {
												_hx_tmp_48 = nil
											}

											tmp = any(map[string]any{"e": _hx_tmp_46, "pmin": _hx_tmp_47, "pmax": _hx_tmp_48, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
										}
									}
								}

								Hx_Field_go_haxe_hxdynamic_setField(c, "expr", tmp)
							}

						case "default":
							{
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(def, nil))) {
									this.VTable.Hx_Field_unexpected(tk)
								}

								{
									var t Hx_Enum_loreline_lorscript_token
									_ = t
									t = this.VTable.Hx_Field_token()
									if (t != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
										this.VTable.Hx_Field_unexpected(t)
									}
								}

								var exprs HxArray[any]
								_ = exprs
								exprs = HxMakeArray[any]().Dyn()
							hx_label_3:
								for {
									tk = this.VTable.Hx_Field_token()
									{
										this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
										this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
										this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
									}

									switch tk.Hx_Field_enumIndex() {
									case 0:
										{
											if this.Hx_Field_resumeErrors {
												break hx_label_3
											} else {
												this.VTable.Hx_Field_parseFullExpr(exprs)
											}
										}

									case 2:
										{
											switch tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s {
											case "case", "default":
												{
													break hx_label_3
												}

											default:
												{
													this.VTable.Hx_Field_parseFullExpr(exprs)
												}
											}
										}

									case 7:
										{
											break hx_label_3
										}

									default:
										{
											this.VTable.Hx_Field_parseFullExpr(exprs)
										}
									}
								}

								if Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) == 1 {
									def = exprs.Get(0)
								} else {
									if Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) == 0 {
										var tmp_e_1 Hx_Enum_loreline_lorscript_exprdef
										_ = tmp_e_1
										tmp_e_1 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any]().Dyn()}

										var pmin HxNullable[int32]
										_ = pmin
										pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

										var pmax HxNullable[int32]
										_ = pmax
										pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

										if tmp_e_1 == nil {
											def = nil
										} else {
											if (pmin).Valid == false {
												pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
											}

											if (pmax).Valid == false {
												pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
											}

											var _hx_tmp_49 any
											_ = _hx_tmp_49
											_hx_tmp_49 = ((any)(tmp_e_1))
											var hx_nullable_f121 HxNullable[int32]
											_ = hx_nullable_f121
											hx_nullable_f121 = pmin
											var _hx_tmp_50 any
											_ = _hx_tmp_50
											if (hx_nullable_f121).Valid == true {
												_hx_tmp_50 = (hx_nullable_f121).Value
											} else {
												_hx_tmp_50 = nil
											}

											var hx_nullable_f122 HxNullable[int32]
											_ = hx_nullable_f122
											hx_nullable_f122 = pmax
											var _hx_tmp_51 any
											_ = _hx_tmp_51
											if (hx_nullable_f122).Valid == true {
												_hx_tmp_51 = (hx_nullable_f122).Value
											} else {
												_hx_tmp_51 = nil
											}

											def = any(map[string]any{"e": _hx_tmp_49, "pmin": _hx_tmp_50, "pmax": _hx_tmp_51, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
										}
									} else {
										var tmp_e_1 Hx_Enum_loreline_lorscript_exprdef
										_ = tmp_e_1
										tmp_e_1 = Hx_Enum_loreline_lorscript_exprdef_EBlock{exprs}

										var e1 any
										_ = e1
										e1 = exprs.Get(0)
										var _hx_tmp_52 int32
										_ = _hx_tmp_52
										if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
											_hx_tmp_52 = 0
										} else {
											_hx_tmp_52 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
										}

										var pmin HxNullable[int32]
										_ = pmin
										pmin = HxNullable[int32]{Value: _hx_tmp_52, Valid: true}

										var _hx_tmp_53 HxArray[any]
										_ = _hx_tmp_53
										_hx_tmp_53 = exprs
										var e2 any
										_ = e2
										e2 = _hx_tmp_53.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(exprs) - 1))
										var _hx_tmp_54 int32
										_ = _hx_tmp_54
										if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e2, nil))) {
											_hx_tmp_54 = 0
										} else {
											_hx_tmp_54 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e2, "pmax"))))
										}

										var pmax HxNullable[int32]
										_ = pmax
										pmax = HxNullable[int32]{Value: _hx_tmp_54, Valid: true}

										if tmp_e_1 == nil {
											def = nil
										} else {
											if (pmin).Valid == false {
												pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
											}

											if (pmax).Valid == false {
												pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
											}

											var _hx_tmp_55 any
											_ = _hx_tmp_55
											_hx_tmp_55 = ((any)(tmp_e_1))
											var hx_nullable_f123 HxNullable[int32]
											_ = hx_nullable_f123
											hx_nullable_f123 = pmin
											var _hx_tmp_56 any
											_ = _hx_tmp_56
											if (hx_nullable_f123).Valid == true {
												_hx_tmp_56 = (hx_nullable_f123).Value
											} else {
												_hx_tmp_56 = nil
											}

											var hx_nullable_f124 HxNullable[int32]
											_ = hx_nullable_f124
											hx_nullable_f124 = pmax
											var _hx_tmp_57 any
											_ = _hx_tmp_57
											if (hx_nullable_f124).Valid == true {
												_hx_tmp_57 = (hx_nullable_f124).Value
											} else {
												_hx_tmp_57 = nil
											}

											def = any(map[string]any{"e": _hx_tmp_55, "pmin": _hx_tmp_56, "pmax": _hx_tmp_57, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
										}
									}
								}
							}

						default:
							{
								this.VTable.Hx_Field_unexpected(tk)
								break hx_label_4
							}
						}
					}

				case 7:
					{
						break hx_label_4
					}

				default:
					{
						this.VTable.Hx_Field_unexpected(tk)
						break hx_label_4
					}
				}
			}

			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_ESwitch{e, cases, def}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_58 any
				_ = _hx_tmp_58
				_hx_tmp_58 = ((any)(e1))
				var hx_nullable_f125 HxNullable[int32]
				_ = hx_nullable_f125
				hx_nullable_f125 = pmin
				var _hx_tmp_59 any
				_ = _hx_tmp_59
				if (hx_nullable_f125).Valid == true {
					_hx_tmp_59 = (hx_nullable_f125).Value
				} else {
					_hx_tmp_59 = nil
				}

				var hx_nullable_f126 HxNullable[int32]
				_ = hx_nullable_f126
				hx_nullable_f126 = pmax
				var _hx_tmp_60 any
				_ = _hx_tmp_60
				if (hx_nullable_f126).Valid == true {
					_hx_tmp_60 = (hx_nullable_f126).Value
				} else {
					_hx_tmp_60 = nil
				}

				return any(map[string]any{"e": _hx_tmp_58, "pmin": _hx_tmp_59, "pmax": _hx_tmp_60, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "throw":
		{
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EThrow{e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_61 int32
			_ = _hx_tmp_61
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_61 = 0
			} else {
				_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_61, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_62 any
				_ = _hx_tmp_62
				_hx_tmp_62 = ((any)(e1))
				var hx_nullable_f127 HxNullable[int32]
				_ = hx_nullable_f127
				hx_nullable_f127 = pmin
				var _hx_tmp_63 any
				_ = _hx_tmp_63
				if (hx_nullable_f127).Valid == true {
					_hx_tmp_63 = (hx_nullable_f127).Value
				} else {
					_hx_tmp_63 = nil
				}

				var hx_nullable_f128 HxNullable[int32]
				_ = hx_nullable_f128
				hx_nullable_f128 = pmax
				var _hx_tmp_64 any
				_ = _hx_tmp_64
				if (hx_nullable_f128).Valid == true {
					_hx_tmp_64 = (hx_nullable_f128).Value
				} else {
					_hx_tmp_64 = nil
				}

				return any(map[string]any{"e": _hx_tmp_62, "pmin": _hx_tmp_63, "pmax": _hx_tmp_64, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "try":
		{
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				var _hx_tmp_65 Hx_Enum_loreline_lorscript_token
				_ = _hx_tmp_65
				_hx_tmp_65 = t
				if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
					var _hx_closure_src_320 func(a any, b any) bool
					_ = _hx_closure_src_320
					_hx_closure_src_320 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
						return _hx_closure_src_320(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_65, Hx_Enum_loreline_lorscript_token_TId{"catch"})) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			{
				var t1 Hx_Enum_loreline_lorscript_token
				_ = t1
				t1 = this.VTable.Hx_Field_token()
				if (t1 != Hx_Enum_loreline_lorscript_token_TPOpen{}) {
					this.VTable.Hx_Field_unexpected(t1)
				}
			}

			var vname string
			_ = vname
			vname = this.VTable.Hx_Field_getIdent()
			{
				var t2 Hx_Enum_loreline_lorscript_token
				_ = t2
				t2 = this.VTable.Hx_Field_token()
				if (t2 != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
					this.VTable.Hx_Field_unexpected(t2)
				}
			}

			var t3 Hx_Enum_loreline_lorscript_ctype
			_ = t3
			if this.Hx_Field_allowTypes {
				t3 = this.VTable.Hx_Field_parseType()
			} else {
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				var _hx_tmp_66 Hx_Enum_loreline_lorscript_token
				_ = _hx_tmp_66
				_hx_tmp_66 = t
				if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
					var _hx_closure_src_321 func(a any, b any) bool
					_ = _hx_closure_src_321
					_hx_closure_src_321 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
						return _hx_closure_src_321(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_66, Hx_Enum_loreline_lorscript_token_TId{"Dynamic"})) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			{
				var t4 Hx_Enum_loreline_lorscript_token
				_ = t4
				t4 = this.VTable.Hx_Field_token()
				if (t4 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
					this.VTable.Hx_Field_unexpected(t4)
				}
			}

			var ec any
			_ = ec
			ec = this.VTable.Hx_Field_parseExpr()
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_ETry{e, vname, t3, ec}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_67 int32
			_ = _hx_tmp_67
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(ec, nil))) {
				_hx_tmp_67 = 0
			} else {
				_hx_tmp_67 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(ec, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_67, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_68 any
				_ = _hx_tmp_68
				_hx_tmp_68 = ((any)(e1))
				var hx_nullable_f129 HxNullable[int32]
				_ = hx_nullable_f129
				hx_nullable_f129 = pmin
				var _hx_tmp_69 any
				_ = _hx_tmp_69
				if (hx_nullable_f129).Valid == true {
					_hx_tmp_69 = (hx_nullable_f129).Value
				} else {
					_hx_tmp_69 = nil
				}

				var hx_nullable_f130 HxNullable[int32]
				_ = hx_nullable_f130
				hx_nullable_f130 = pmax
				var _hx_tmp_70 any
				_ = _hx_tmp_70
				if (hx_nullable_f130).Valid == true {
					_hx_tmp_70 = (hx_nullable_f130).Value
				} else {
					_hx_tmp_70 = nil
				}

				return any(map[string]any{"e": _hx_tmp_68, "pmin": _hx_tmp_69, "pmax": _hx_tmp_70, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "var":
		{
			var ident string
			_ = ident
			ident = this.VTable.Hx_Field_getIdent()
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			var t Hx_Enum_loreline_lorscript_ctype
			_ = t
			if (tk == Hx_Enum_loreline_lorscript_token_TDoubleDot{}) && this.Hx_Field_allowTypes {
				t = this.VTable.Hx_Field_parseType()
				tk = this.VTable.Hx_Field_token()
			}

			var e any
			_ = e
			switch tk.Hx_Field_enumIndex() {
			case 3:
				{
					if tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s == "=" {
						e = this.VTable.Hx_Field_parseExpr()
					} else {
						this.VTable.Hx_Field_unexpected(tk)
					}
				}

			case 10, 11:
				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}

			default:
				{
					if t != nil {
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					} else {
						this.VTable.Hx_Field_unexpected(tk)
					}
				}
			}

			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EVar{ident, t, e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_71 HxNullable[int32]
			_ = _hx_tmp_71
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_71 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
			} else {
				var _hx_tmp_72 int32
				_ = _hx_tmp_72
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
					_hx_tmp_72 = 0
				} else {
					_hx_tmp_72 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
				}

				_hx_tmp_71 = HxNullable[int32]{Value: _hx_tmp_72, Valid: true}
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = _hx_tmp_71
			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_73 any
				_ = _hx_tmp_73
				_hx_tmp_73 = ((any)(e1))
				var hx_nullable_f131 HxNullable[int32]
				_ = hx_nullable_f131
				hx_nullable_f131 = pmin
				var _hx_tmp_74 any
				_ = _hx_tmp_74
				if (hx_nullable_f131).Valid == true {
					_hx_tmp_74 = (hx_nullable_f131).Value
				} else {
					_hx_tmp_74 = nil
				}

				var hx_nullable_f132 HxNullable[int32]
				_ = hx_nullable_f132
				hx_nullable_f132 = pmax
				var _hx_tmp_75 any
				_ = _hx_tmp_75
				if (hx_nullable_f132).Valid == true {
					_hx_tmp_75 = (hx_nullable_f132).Value
				} else {
					_hx_tmp_75 = nil
				}

				return any(map[string]any{"e": _hx_tmp_73, "pmin": _hx_tmp_74, "pmax": _hx_tmp_75, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case "while":
		{
			var econd any
			_ = econd
			econd = this.VTable.Hx_Field_parseExpr()
			var e any
			_ = e
			e = this.VTable.Hx_Field_parseExpr()
			var e1 Hx_Enum_loreline_lorscript_exprdef
			_ = e1
			e1 = Hx_Enum_loreline_lorscript_exprdef_EWhile{econd, e}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: p1, Valid: true}

			var _hx_tmp_76 int32
			_ = _hx_tmp_76
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
				_hx_tmp_76 = 0
			} else {
				_hx_tmp_76 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_76, Valid: true}

			if e1 == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_77 any
				_ = _hx_tmp_77
				_hx_tmp_77 = ((any)(e1))
				var hx_nullable_f133 HxNullable[int32]
				_ = hx_nullable_f133
				hx_nullable_f133 = pmin
				var _hx_tmp_78 any
				_ = _hx_tmp_78
				if (hx_nullable_f133).Valid == true {
					_hx_tmp_78 = (hx_nullable_f133).Value
				} else {
					_hx_tmp_78 = nil
				}

				var hx_nullable_f134 HxNullable[int32]
				_ = hx_nullable_f134
				hx_nullable_f134 = pmax
				var _hx_tmp_79 any
				_ = _hx_tmp_79
				if (hx_nullable_f134).Valid == true {
					_hx_tmp_79 = (hx_nullable_f134).Value
				} else {
					_hx_tmp_79 = nil
				}

				return any(map[string]any{"e": _hx_tmp_77, "pmin": _hx_tmp_78, "pmax": _hx_tmp_79, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	default:
		{
			return nil
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseExprNext(e1 any) any {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	switch tk.Hx_Field_enumIndex() {
	case 2:
		{
			var op string
			_ = op
			op = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
			if op == "is" {
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = op
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = e1
				return this.VTable.Hx_Field_makeBinop(_hx_tmp_0, _hx_tmp_1, this.VTable.Hx_Field_parseExpr())
			} else {
				{
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}

				return e1
			}
		}

	case 3:
		{
			var op string
			_ = op
			op = tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
			if op == "->" {
				{
					var _g Hx_Enum_loreline_lorscript_exprdef
					_ = _g
					_g = func() Hx_Enum_loreline_lorscript_exprdef {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e"))
						if _hx_d == nil {
							var _hx_z Hx_Enum_loreline_lorscript_exprdef
							return _hx_z
						}
						return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
					}()
					switch _g.Hx_Field_enumIndex() {
					case 1:
						{
							var i string
							_ = i
							i = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
							var eret any
							_ = eret
							eret = this.VTable.Hx_Field_parseExpr()
							var e Hx_Enum_loreline_lorscript_exprdef
							_ = e
							e = Hx_Enum_loreline_lorscript_exprdef_EReturn{eret}

							var _hx_tmp_2 int32
							_ = _hx_tmp_2
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(eret, nil))) {
								_hx_tmp_2 = 0
							} else {
								_hx_tmp_2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(eret, "pmin"))))
							}

							var pmin HxNullable[int32]
							_ = pmin
							pmin = HxNullable[int32]{Value: _hx_tmp_2, Valid: true}

							var pmax HxNullable[int32]
							_ = pmax
							pmax = HxNullable[int32]{}

							var e2 any
							_ = e2
							if e == nil {
								e2 = nil
							} else {
								if (pmin).Valid == false {
									pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
								}

								if (pmax).Valid == false {
									pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
								}

								var _hx_tmp_3 any
								_ = _hx_tmp_3
								_hx_tmp_3 = ((any)(e))
								var hx_nullable_f135 HxNullable[int32]
								_ = hx_nullable_f135
								hx_nullable_f135 = pmin
								var _hx_tmp_4 any
								_ = _hx_tmp_4
								if (hx_nullable_f135).Valid == true {
									_hx_tmp_4 = (hx_nullable_f135).Value
								} else {
									_hx_tmp_4 = nil
								}

								var hx_nullable_f136 HxNullable[int32]
								_ = hx_nullable_f136
								hx_nullable_f136 = pmax
								var _hx_tmp_5 any
								_ = _hx_tmp_5
								if (hx_nullable_f136).Valid == true {
									_hx_tmp_5 = (hx_nullable_f136).Value
								} else {
									_hx_tmp_5 = nil
								}

								e2 = any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
							}

							var _hx_tmp_6 HxArray[any]
							_ = _hx_tmp_6
							_hx_tmp_6 = HxMakeArray[any](any(map[string]any{"name": ((any)(i))})).Dyn()
							var _hx_tmp_7 any
							_ = _hx_tmp_7
							_hx_tmp_7 = e2
							var e3 Hx_Enum_loreline_lorscript_exprdef
							_ = e3
							e3 = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_6, _hx_tmp_7, HxNullable[string]{Value: HxStringNull}, nil}

							var _hx_tmp_8 int32
							_ = _hx_tmp_8
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
								_hx_tmp_8 = 0
							} else {
								_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
							}

							var pmin1 HxNullable[int32]
							_ = pmin1
							pmin1 = HxNullable[int32]{Value: _hx_tmp_8, Valid: true}

							var pmax1 HxNullable[int32]
							_ = pmax1
							pmax1 = HxNullable[int32]{}

							if e3 == nil {
								return nil
							} else {
								if (pmin1).Valid == false {
									pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
								}

								if (pmax1).Valid == false {
									pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
								}

								var _hx_tmp_9 any
								_ = _hx_tmp_9
								_hx_tmp_9 = ((any)(e3))
								var hx_nullable_f137 HxNullable[int32]
								_ = hx_nullable_f137
								hx_nullable_f137 = pmin1
								var _hx_tmp_10 any
								_ = _hx_tmp_10
								if (hx_nullable_f137).Valid == true {
									_hx_tmp_10 = (hx_nullable_f137).Value
								} else {
									_hx_tmp_10 = nil
								}

								var hx_nullable_f138 HxNullable[int32]
								_ = hx_nullable_f138
								hx_nullable_f138 = pmax1
								var _hx_tmp_11 any
								_ = _hx_tmp_11
								if (hx_nullable_f138).Valid == true {
									_hx_tmp_11 = (hx_nullable_f138).Value
								} else {
									_hx_tmp_11 = nil
								}

								return any(map[string]any{"e": _hx_tmp_9, "pmin": _hx_tmp_10, "pmax": _hx_tmp_11, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
							}
						}

					case 3:
						{
							var _hx_tmp Hx_Enum_loreline_lorscript_exprdef
							_ = _hx_tmp
							_hx_tmp = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(_g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if _hx_tmp.Hx_Field_enumIndex() == 1 {
								var i string
								_ = i
								i = _hx_tmp.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
								var eret any
								_ = eret
								eret = this.VTable.Hx_Field_parseExpr()
								var e Hx_Enum_loreline_lorscript_exprdef
								_ = e
								e = Hx_Enum_loreline_lorscript_exprdef_EReturn{eret}

								var _hx_tmp_12 int32
								_ = _hx_tmp_12
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(eret, nil))) {
									_hx_tmp_12 = 0
								} else {
									_hx_tmp_12 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(eret, "pmin"))))
								}

								var pmin HxNullable[int32]
								_ = pmin
								pmin = HxNullable[int32]{Value: _hx_tmp_12, Valid: true}

								var pmax HxNullable[int32]
								_ = pmax
								pmax = HxNullable[int32]{}

								var e2 any
								_ = e2
								if e == nil {
									e2 = nil
								} else {
									if (pmin).Valid == false {
										pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
									}

									if (pmax).Valid == false {
										pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
									}

									var _hx_tmp_13 any
									_ = _hx_tmp_13
									_hx_tmp_13 = ((any)(e))
									var hx_nullable_f139 HxNullable[int32]
									_ = hx_nullable_f139
									hx_nullable_f139 = pmin
									var _hx_tmp_14 any
									_ = _hx_tmp_14
									if (hx_nullable_f139).Valid == true {
										_hx_tmp_14 = (hx_nullable_f139).Value
									} else {
										_hx_tmp_14 = nil
									}

									var hx_nullable_f140 HxNullable[int32]
									_ = hx_nullable_f140
									hx_nullable_f140 = pmax
									var _hx_tmp_15 any
									_ = _hx_tmp_15
									if (hx_nullable_f140).Valid == true {
										_hx_tmp_15 = (hx_nullable_f140).Value
									} else {
										_hx_tmp_15 = nil
									}

									e2 = any(map[string]any{"e": _hx_tmp_13, "pmin": _hx_tmp_14, "pmax": _hx_tmp_15, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
								}

								var _hx_tmp_16 HxArray[any]
								_ = _hx_tmp_16
								_hx_tmp_16 = HxMakeArray[any](any(map[string]any{"name": ((any)(i))})).Dyn()
								var _hx_tmp_17 any
								_ = _hx_tmp_17
								_hx_tmp_17 = e2
								var e3 Hx_Enum_loreline_lorscript_exprdef
								_ = e3
								e3 = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_16, _hx_tmp_17, HxNullable[string]{Value: HxStringNull}, nil}

								var _hx_tmp_18 int32
								_ = _hx_tmp_18
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
									_hx_tmp_18 = 0
								} else {
									_hx_tmp_18 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
								}

								var pmin1 HxNullable[int32]
								_ = pmin1
								pmin1 = HxNullable[int32]{Value: _hx_tmp_18, Valid: true}

								var pmax1 HxNullable[int32]
								_ = pmax1
								pmax1 = HxNullable[int32]{}

								if e3 == nil {
									return nil
								} else {
									if (pmin1).Valid == false {
										pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
									}

									if (pmax1).Valid == false {
										pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
									}

									var _hx_tmp_19 any
									_ = _hx_tmp_19
									_hx_tmp_19 = ((any)(e3))
									var hx_nullable_f141 HxNullable[int32]
									_ = hx_nullable_f141
									hx_nullable_f141 = pmin1
									var _hx_tmp_20 any
									_ = _hx_tmp_20
									if (hx_nullable_f141).Valid == true {
										_hx_tmp_20 = (hx_nullable_f141).Value
									} else {
										_hx_tmp_20 = nil
									}

									var hx_nullable_f142 HxNullable[int32]
									_ = hx_nullable_f142
									hx_nullable_f142 = pmax1
									var _hx_tmp_21 any
									_ = _hx_tmp_21
									if (hx_nullable_f142).Valid == true {
										_hx_tmp_21 = (hx_nullable_f142).Value
									} else {
										_hx_tmp_21 = nil
									}

									return any(map[string]any{"e": _hx_tmp_19, "pmin": _hx_tmp_20, "pmax": _hx_tmp_21, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
								}
							}
						}

					case 26:
						{
							var _hx_tmp Hx_Enum_loreline_lorscript_exprdef
							_ = _hx_tmp
							_hx_tmp = func() Hx_Enum_loreline_lorscript_exprdef {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(_g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_e, "e"))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_exprdef
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
							}()
							if _hx_tmp.Hx_Field_enumIndex() == 1 {
								var i string
								_ = i
								i = _hx_tmp.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
								var t Hx_Enum_loreline_lorscript_ctype
								_ = t
								t = _g.(Hx_Enum_loreline_lorscript_exprdef_ECheckType).Hx_Field_t
								var eret any
								_ = eret
								eret = this.VTable.Hx_Field_parseExpr()
								var e Hx_Enum_loreline_lorscript_exprdef
								_ = e
								e = Hx_Enum_loreline_lorscript_exprdef_EReturn{eret}

								var _hx_tmp_22 int32
								_ = _hx_tmp_22
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(eret, nil))) {
									_hx_tmp_22 = 0
								} else {
									_hx_tmp_22 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(eret, "pmin"))))
								}

								var pmin HxNullable[int32]
								_ = pmin
								pmin = HxNullable[int32]{Value: _hx_tmp_22, Valid: true}

								var pmax HxNullable[int32]
								_ = pmax
								pmax = HxNullable[int32]{}

								var e2 any
								_ = e2
								if e == nil {
									e2 = nil
								} else {
									if (pmin).Valid == false {
										pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
									}

									if (pmax).Valid == false {
										pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
									}

									var _hx_tmp_23 any
									_ = _hx_tmp_23
									_hx_tmp_23 = ((any)(e))
									var hx_nullable_f143 HxNullable[int32]
									_ = hx_nullable_f143
									hx_nullable_f143 = pmin
									var _hx_tmp_24 any
									_ = _hx_tmp_24
									if (hx_nullable_f143).Valid == true {
										_hx_tmp_24 = (hx_nullable_f143).Value
									} else {
										_hx_tmp_24 = nil
									}

									var hx_nullable_f144 HxNullable[int32]
									_ = hx_nullable_f144
									hx_nullable_f144 = pmax
									var _hx_tmp_25 any
									_ = _hx_tmp_25
									if (hx_nullable_f144).Valid == true {
										_hx_tmp_25 = (hx_nullable_f144).Value
									} else {
										_hx_tmp_25 = nil
									}

									e2 = any(map[string]any{"e": _hx_tmp_23, "pmin": _hx_tmp_24, "pmax": _hx_tmp_25, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
								}

								var _hx_tmp_26 HxArray[any]
								_ = _hx_tmp_26
								_hx_tmp_26 = HxMakeArray[any](any(map[string]any{"name": ((any)(i)), "t": ((any)(t))})).Dyn()
								var _hx_tmp_27 any
								_ = _hx_tmp_27
								_hx_tmp_27 = e2
								var e3 Hx_Enum_loreline_lorscript_exprdef
								_ = e3
								e3 = Hx_Enum_loreline_lorscript_exprdef_EFunction{_hx_tmp_26, _hx_tmp_27, HxNullable[string]{Value: HxStringNull}, nil}

								var _hx_tmp_28 int32
								_ = _hx_tmp_28
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
									_hx_tmp_28 = 0
								} else {
									_hx_tmp_28 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
								}

								var pmin1 HxNullable[int32]
								_ = pmin1
								pmin1 = HxNullable[int32]{Value: _hx_tmp_28, Valid: true}

								var pmax1 HxNullable[int32]
								_ = pmax1
								pmax1 = HxNullable[int32]{}

								if e3 == nil {
									return nil
								} else {
									if (pmin1).Valid == false {
										pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
									}

									if (pmax1).Valid == false {
										pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
									}

									var _hx_tmp_29 any
									_ = _hx_tmp_29
									_hx_tmp_29 = ((any)(e3))
									var hx_nullable_f145 HxNullable[int32]
									_ = hx_nullable_f145
									hx_nullable_f145 = pmin1
									var _hx_tmp_30 any
									_ = _hx_tmp_30
									if (hx_nullable_f145).Valid == true {
										_hx_tmp_30 = (hx_nullable_f145).Value
									} else {
										_hx_tmp_30 = nil
									}

									var hx_nullable_f146 HxNullable[int32]
									_ = hx_nullable_f146
									hx_nullable_f146 = pmax1
									var _hx_tmp_31 any
									_ = _hx_tmp_31
									if (hx_nullable_f146).Valid == true {
										_hx_tmp_31 = (hx_nullable_f146).Value
									} else {
										_hx_tmp_31 = nil
									}

									return any(map[string]any{"e": _hx_tmp_29, "pmin": _hx_tmp_30, "pmax": _hx_tmp_31, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
								}
							}
						}

					default:
						{
						}
					}
				}

				this.VTable.Hx_Field_unexpected(tk)
			}

			var hx_cmp_8 HxNullable[int32]
			_ = hx_cmp_8
			hx_cmp_8 = func() func(_hx_param_key string) HxNullable[int32] {
				var _hx_closure_src_323 func(_hx_param_key any) any
				_ = _hx_closure_src_323
				_hx_closure_src_323 = this.Hx_Field_opPriority.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) HxNullable[int32] {
					var hx_nullable_t43 any
					_ = hx_nullable_t43
					hx_nullable_t43 = _hx_closure_src_323(((any)(_hx_param_key)))
					var _hx_tmp_32 HxNullable[int32]
					_ = _hx_tmp_32
					if hx_nullable_t43 == nil {
						_hx_tmp_32 = HxNullable[int32]{}
					} else {
						_hx_tmp_32 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t43))), Valid: true}
					}

					return _hx_tmp_32
				}
			}()(op)
			var _hx_tmp_33 bool
			_ = _hx_tmp_33
			if (hx_cmp_8).Valid == true {
				_hx_tmp_33 = ((hx_cmp_8).Value == -1)
			} else {
				_hx_tmp_33 = false
			}

			if _hx_tmp_33 {
				var _hx_tmp_34 bool
				_ = _hx_tmp_34
				if this.VTable.Hx_Field_isBlock(e1) {
					_hx_tmp_34 = true
				} else {
					_hx_tmp_34 = (Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e1, "e")).(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue).Hx_Field_enumIndex() == 3)
				}

				if _hx_tmp_34 {
					{
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}

					return e1
				}

				var e Hx_Enum_loreline_lorscript_exprdef
				_ = e
				e = Hx_Enum_loreline_lorscript_exprdef_EUnop{op, false, e1}

				var _hx_tmp_35 int32
				_ = _hx_tmp_35
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
					_hx_tmp_35 = 0
				} else {
					_hx_tmp_35 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
				}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: _hx_tmp_35, Valid: true}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{}

				var tmp any
				_ = tmp
				if e == nil {
					tmp = nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_36 any
					_ = _hx_tmp_36
					_hx_tmp_36 = ((any)(e))
					var hx_nullable_f147 HxNullable[int32]
					_ = hx_nullable_f147
					hx_nullable_f147 = pmin
					var _hx_tmp_37 any
					_ = _hx_tmp_37
					if (hx_nullable_f147).Valid == true {
						_hx_tmp_37 = (hx_nullable_f147).Value
					} else {
						_hx_tmp_37 = nil
					}

					var hx_nullable_f148 HxNullable[int32]
					_ = hx_nullable_f148
					hx_nullable_f148 = pmax
					var _hx_tmp_38 any
					_ = _hx_tmp_38
					if (hx_nullable_f148).Valid == true {
						_hx_tmp_38 = (hx_nullable_f148).Value
					} else {
						_hx_tmp_38 = nil
					}

					tmp = any(map[string]any{"e": _hx_tmp_36, "pmin": _hx_tmp_37, "pmax": _hx_tmp_38, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}

				return this.VTable.Hx_Field_parseExprNext(tmp)
			}

			var _hx_tmp_39 string
			_ = _hx_tmp_39
			_hx_tmp_39 = op
			var _hx_tmp_40 any
			_ = _hx_tmp_40
			_hx_tmp_40 = e1
			return this.VTable.Hx_Field_makeBinop(_hx_tmp_39, _hx_tmp_40, this.VTable.Hx_Field_parseExpr())
		}

	case 4:
		{
			var _hx_tmp_41 any
			_ = _hx_tmp_41
			_hx_tmp_41 = e1
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_ECall{_hx_tmp_41, this.VTable.Hx_Field_parseExprList(Hx_Enum_loreline_lorscript_token_TPClose{})}

			var _hx_tmp_42 int32
			_ = _hx_tmp_42
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_42 = 0
			} else {
				_hx_tmp_42 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_42, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			var tmp any
			_ = tmp
			if e == nil {
				tmp = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_43 any
				_ = _hx_tmp_43
				_hx_tmp_43 = ((any)(e))
				var hx_nullable_f149 HxNullable[int32]
				_ = hx_nullable_f149
				hx_nullable_f149 = pmin
				var _hx_tmp_44 any
				_ = _hx_tmp_44
				if (hx_nullable_f149).Valid == true {
					_hx_tmp_44 = (hx_nullable_f149).Value
				} else {
					_hx_tmp_44 = nil
				}

				var hx_nullable_f150 HxNullable[int32]
				_ = hx_nullable_f150
				hx_nullable_f150 = pmax
				var _hx_tmp_45 any
				_ = _hx_tmp_45
				if (hx_nullable_f150).Valid == true {
					_hx_tmp_45 = (hx_nullable_f150).Value
				} else {
					_hx_tmp_45 = nil
				}

				tmp = any(map[string]any{"e": _hx_tmp_43, "pmin": _hx_tmp_44, "pmax": _hx_tmp_45, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(tmp)
		}

	case 8:
		{
			var field string
			_ = field
			field = this.VTable.Hx_Field_getIdent()
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EField{e1, field}

			var _hx_tmp_46 int32
			_ = _hx_tmp_46
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_46 = 0
			} else {
				_hx_tmp_46 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_46, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			var tmp any
			_ = tmp
			if e == nil {
				tmp = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_47 any
				_ = _hx_tmp_47
				_hx_tmp_47 = ((any)(e))
				var hx_nullable_f151 HxNullable[int32]
				_ = hx_nullable_f151
				hx_nullable_f151 = pmin
				var _hx_tmp_48 any
				_ = _hx_tmp_48
				if (hx_nullable_f151).Valid == true {
					_hx_tmp_48 = (hx_nullable_f151).Value
				} else {
					_hx_tmp_48 = nil
				}

				var hx_nullable_f152 HxNullable[int32]
				_ = hx_nullable_f152
				hx_nullable_f152 = pmax
				var _hx_tmp_49 any
				_ = _hx_tmp_49
				if (hx_nullable_f152).Valid == true {
					_hx_tmp_49 = (hx_nullable_f152).Value
				} else {
					_hx_tmp_49 = nil
				}

				tmp = any(map[string]any{"e": _hx_tmp_47, "pmin": _hx_tmp_48, "pmax": _hx_tmp_49, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(tmp)
		}

	case 9:
		{
			var field string
			_ = field
			field = this.VTable.Hx_Field_getIdent()
			var _hx_tmp_50 int32
			_ = _hx_tmp_50
			_hx_tmp_50 = this.Hx_Field_uid
			this.Hx_Field_uid = (this.Hx_Field_uid + 1)
			var tmp string
			_ = tmp
			tmp = ("__a_" + Hx_Field_go_haxe_hxdynamic_toString(_hx_tmp_50))
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EVar{tmp, nil, e1}

			var _hx_tmp_51 int32
			_ = _hx_tmp_51
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_51 = 0
			} else {
				_hx_tmp_51 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_51, Valid: true}

			var _hx_tmp_52 int32
			_ = _hx_tmp_52
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_52 = 0
			} else {
				_hx_tmp_52 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_52, Valid: true}

			var e2 any
			_ = e2
			if e == nil {
				e2 = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_53 any
				_ = _hx_tmp_53
				_hx_tmp_53 = ((any)(e))
				var hx_nullable_f153 HxNullable[int32]
				_ = hx_nullable_f153
				hx_nullable_f153 = pmin
				var _hx_tmp_54 any
				_ = _hx_tmp_54
				if (hx_nullable_f153).Valid == true {
					_hx_tmp_54 = (hx_nullable_f153).Value
				} else {
					_hx_tmp_54 = nil
				}

				var hx_nullable_f154 HxNullable[int32]
				_ = hx_nullable_f154
				hx_nullable_f154 = pmax
				var _hx_tmp_55 any
				_ = _hx_tmp_55
				if (hx_nullable_f154).Valid == true {
					_hx_tmp_55 = (hx_nullable_f154).Value
				} else {
					_hx_tmp_55 = nil
				}

				e2 = any(map[string]any{"e": _hx_tmp_53, "pmin": _hx_tmp_54, "pmax": _hx_tmp_55, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e3 Hx_Enum_loreline_lorscript_exprdef
			_ = e3
			e3 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

			var _hx_tmp_56 int32
			_ = _hx_tmp_56
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_56 = 0
			} else {
				_hx_tmp_56 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin1 HxNullable[int32]
			_ = pmin1
			pmin1 = HxNullable[int32]{Value: _hx_tmp_56, Valid: true}

			var _hx_tmp_57 int32
			_ = _hx_tmp_57
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_57 = 0
			} else {
				_hx_tmp_57 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax1 HxNullable[int32]
			_ = pmax1
			pmax1 = HxNullable[int32]{Value: _hx_tmp_57, Valid: true}

			var e4 any
			_ = e4
			if e3 == nil {
				e4 = nil
			} else {
				if (pmin1).Valid == false {
					pmin1 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax1).Valid == false {
					pmax1 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_58 any
				_ = _hx_tmp_58
				_hx_tmp_58 = ((any)(e3))
				var hx_nullable_f155 HxNullable[int32]
				_ = hx_nullable_f155
				hx_nullable_f155 = pmin1
				var _hx_tmp_59 any
				_ = _hx_tmp_59
				if (hx_nullable_f155).Valid == true {
					_hx_tmp_59 = (hx_nullable_f155).Value
				} else {
					_hx_tmp_59 = nil
				}

				var hx_nullable_f156 HxNullable[int32]
				_ = hx_nullable_f156
				hx_nullable_f156 = pmax1
				var _hx_tmp_60 any
				_ = _hx_tmp_60
				if (hx_nullable_f156).Valid == true {
					_hx_tmp_60 = (hx_nullable_f156).Value
				} else {
					_hx_tmp_60 = nil
				}

				e4 = any(map[string]any{"e": _hx_tmp_58, "pmin": _hx_tmp_59, "pmax": _hx_tmp_60, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e5 Hx_Enum_loreline_lorscript_exprdef
			_ = e5
			e5 = Hx_Enum_loreline_lorscript_exprdef_EIdent{"null"}

			var _hx_tmp_61 int32
			_ = _hx_tmp_61
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_61 = 0
			} else {
				_hx_tmp_61 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin2 HxNullable[int32]
			_ = pmin2
			pmin2 = HxNullable[int32]{Value: _hx_tmp_61, Valid: true}

			var _hx_tmp_62 int32
			_ = _hx_tmp_62
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_62 = 0
			} else {
				_hx_tmp_62 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax2 HxNullable[int32]
			_ = pmax2
			pmax2 = HxNullable[int32]{Value: _hx_tmp_62, Valid: true}

			var e6 any
			_ = e6
			if e5 == nil {
				e6 = nil
			} else {
				if (pmin2).Valid == false {
					pmin2 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax2).Valid == false {
					pmax2 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_63 any
				_ = _hx_tmp_63
				_hx_tmp_63 = ((any)(e5))
				var hx_nullable_f157 HxNullable[int32]
				_ = hx_nullable_f157
				hx_nullable_f157 = pmin2
				var _hx_tmp_64 any
				_ = _hx_tmp_64
				if (hx_nullable_f157).Valid == true {
					_hx_tmp_64 = (hx_nullable_f157).Value
				} else {
					_hx_tmp_64 = nil
				}

				var hx_nullable_f158 HxNullable[int32]
				_ = hx_nullable_f158
				hx_nullable_f158 = pmax2
				var _hx_tmp_65 any
				_ = _hx_tmp_65
				if (hx_nullable_f158).Valid == true {
					_hx_tmp_65 = (hx_nullable_f158).Value
				} else {
					_hx_tmp_65 = nil
				}

				e6 = any(map[string]any{"e": _hx_tmp_63, "pmin": _hx_tmp_64, "pmax": _hx_tmp_65, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e7 Hx_Enum_loreline_lorscript_exprdef
			_ = e7
			e7 = Hx_Enum_loreline_lorscript_exprdef_EBinop{"==", e4, e6}

			var pmin3 HxNullable[int32]
			_ = pmin3
			pmin3 = HxNullable[int32]{}

			var pmax3 HxNullable[int32]
			_ = pmax3
			pmax3 = HxNullable[int32]{}

			var e8 any
			_ = e8
			if e7 == nil {
				e8 = nil
			} else {
				if (pmin3).Valid == false {
					pmin3 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax3).Valid == false {
					pmax3 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_66 any
				_ = _hx_tmp_66
				_hx_tmp_66 = ((any)(e7))
				var hx_nullable_f159 HxNullable[int32]
				_ = hx_nullable_f159
				hx_nullable_f159 = pmin3
				var _hx_tmp_67 any
				_ = _hx_tmp_67
				if (hx_nullable_f159).Valid == true {
					_hx_tmp_67 = (hx_nullable_f159).Value
				} else {
					_hx_tmp_67 = nil
				}

				var hx_nullable_f160 HxNullable[int32]
				_ = hx_nullable_f160
				hx_nullable_f160 = pmax3
				var _hx_tmp_68 any
				_ = _hx_tmp_68
				if (hx_nullable_f160).Valid == true {
					_hx_tmp_68 = (hx_nullable_f160).Value
				} else {
					_hx_tmp_68 = nil
				}

				e8 = any(map[string]any{"e": _hx_tmp_66, "pmin": _hx_tmp_67, "pmax": _hx_tmp_68, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e9 Hx_Enum_loreline_lorscript_exprdef
			_ = e9
			e9 = Hx_Enum_loreline_lorscript_exprdef_EIdent{"null"}

			var _hx_tmp_69 int32
			_ = _hx_tmp_69
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_69 = 0
			} else {
				_hx_tmp_69 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin4 HxNullable[int32]
			_ = pmin4
			pmin4 = HxNullable[int32]{Value: _hx_tmp_69, Valid: true}

			var _hx_tmp_70 int32
			_ = _hx_tmp_70
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_70 = 0
			} else {
				_hx_tmp_70 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax4 HxNullable[int32]
			_ = pmax4
			pmax4 = HxNullable[int32]{Value: _hx_tmp_70, Valid: true}

			var e10 any
			_ = e10
			if e9 == nil {
				e10 = nil
			} else {
				if (pmin4).Valid == false {
					pmin4 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax4).Valid == false {
					pmax4 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_71 any
				_ = _hx_tmp_71
				_hx_tmp_71 = ((any)(e9))
				var hx_nullable_f161 HxNullable[int32]
				_ = hx_nullable_f161
				hx_nullable_f161 = pmin4
				var _hx_tmp_72 any
				_ = _hx_tmp_72
				if (hx_nullable_f161).Valid == true {
					_hx_tmp_72 = (hx_nullable_f161).Value
				} else {
					_hx_tmp_72 = nil
				}

				var hx_nullable_f162 HxNullable[int32]
				_ = hx_nullable_f162
				hx_nullable_f162 = pmax4
				var _hx_tmp_73 any
				_ = _hx_tmp_73
				if (hx_nullable_f162).Valid == true {
					_hx_tmp_73 = (hx_nullable_f162).Value
				} else {
					_hx_tmp_73 = nil
				}

				e10 = any(map[string]any{"e": _hx_tmp_71, "pmin": _hx_tmp_72, "pmax": _hx_tmp_73, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e11 Hx_Enum_loreline_lorscript_exprdef
			_ = e11
			e11 = Hx_Enum_loreline_lorscript_exprdef_EIdent{tmp}

			var _hx_tmp_74 int32
			_ = _hx_tmp_74
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_74 = 0
			} else {
				_hx_tmp_74 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin5 HxNullable[int32]
			_ = pmin5
			pmin5 = HxNullable[int32]{Value: _hx_tmp_74, Valid: true}

			var _hx_tmp_75 int32
			_ = _hx_tmp_75
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_75 = 0
			} else {
				_hx_tmp_75 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmax"))))
			}

			var pmax5 HxNullable[int32]
			_ = pmax5
			pmax5 = HxNullable[int32]{Value: _hx_tmp_75, Valid: true}

			var e12 any
			_ = e12
			if e11 == nil {
				e12 = nil
			} else {
				if (pmin5).Valid == false {
					pmin5 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax5).Valid == false {
					pmax5 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_76 any
				_ = _hx_tmp_76
				_hx_tmp_76 = ((any)(e11))
				var hx_nullable_f163 HxNullable[int32]
				_ = hx_nullable_f163
				hx_nullable_f163 = pmin5
				var _hx_tmp_77 any
				_ = _hx_tmp_77
				if (hx_nullable_f163).Valid == true {
					_hx_tmp_77 = (hx_nullable_f163).Value
				} else {
					_hx_tmp_77 = nil
				}

				var hx_nullable_f164 HxNullable[int32]
				_ = hx_nullable_f164
				hx_nullable_f164 = pmax5
				var _hx_tmp_78 any
				_ = _hx_tmp_78
				if (hx_nullable_f164).Valid == true {
					_hx_tmp_78 = (hx_nullable_f164).Value
				} else {
					_hx_tmp_78 = nil
				}

				e12 = any(map[string]any{"e": _hx_tmp_76, "pmin": _hx_tmp_77, "pmax": _hx_tmp_78, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e13 Hx_Enum_loreline_lorscript_exprdef
			_ = e13
			e13 = Hx_Enum_loreline_lorscript_exprdef_EField{e12, field}

			var _hx_tmp_79 int32
			_ = _hx_tmp_79
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_79 = 0
			} else {
				_hx_tmp_79 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin6 HxNullable[int32]
			_ = pmin6
			pmin6 = HxNullable[int32]{Value: _hx_tmp_79, Valid: true}

			var pmax6 HxNullable[int32]
			_ = pmax6
			pmax6 = HxNullable[int32]{}

			var e14 any
			_ = e14
			if e13 == nil {
				e14 = nil
			} else {
				if (pmin6).Valid == false {
					pmin6 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax6).Valid == false {
					pmax6 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_80 any
				_ = _hx_tmp_80
				_hx_tmp_80 = ((any)(e13))
				var hx_nullable_f165 HxNullable[int32]
				_ = hx_nullable_f165
				hx_nullable_f165 = pmin6
				var _hx_tmp_81 any
				_ = _hx_tmp_81
				if (hx_nullable_f165).Valid == true {
					_hx_tmp_81 = (hx_nullable_f165).Value
				} else {
					_hx_tmp_81 = nil
				}

				var hx_nullable_f166 HxNullable[int32]
				_ = hx_nullable_f166
				hx_nullable_f166 = pmax6
				var _hx_tmp_82 any
				_ = _hx_tmp_82
				if (hx_nullable_f166).Valid == true {
					_hx_tmp_82 = (hx_nullable_f166).Value
				} else {
					_hx_tmp_82 = nil
				}

				e14 = any(map[string]any{"e": _hx_tmp_80, "pmin": _hx_tmp_81, "pmax": _hx_tmp_82, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e15 Hx_Enum_loreline_lorscript_exprdef
			_ = e15
			e15 = Hx_Enum_loreline_lorscript_exprdef_ETernary{e8, e10, e14}

			var pmin7 HxNullable[int32]
			_ = pmin7
			pmin7 = HxNullable[int32]{}

			var pmax7 HxNullable[int32]
			_ = pmax7
			pmax7 = HxNullable[int32]{}

			var e16 any
			_ = e16
			if e15 == nil {
				e16 = nil
			} else {
				if (pmin7).Valid == false {
					pmin7 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax7).Valid == false {
					pmax7 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_83 any
				_ = _hx_tmp_83
				_hx_tmp_83 = ((any)(e15))
				var hx_nullable_f167 HxNullable[int32]
				_ = hx_nullable_f167
				hx_nullable_f167 = pmin7
				var _hx_tmp_84 any
				_ = _hx_tmp_84
				if (hx_nullable_f167).Valid == true {
					_hx_tmp_84 = (hx_nullable_f167).Value
				} else {
					_hx_tmp_84 = nil
				}

				var hx_nullable_f168 HxNullable[int32]
				_ = hx_nullable_f168
				hx_nullable_f168 = pmax7
				var _hx_tmp_85 any
				_ = _hx_tmp_85
				if (hx_nullable_f168).Valid == true {
					_hx_tmp_85 = (hx_nullable_f168).Value
				} else {
					_hx_tmp_85 = nil
				}

				e16 = any(map[string]any{"e": _hx_tmp_83, "pmin": _hx_tmp_84, "pmax": _hx_tmp_85, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			var e17 Hx_Enum_loreline_lorscript_exprdef
			_ = e17
			e17 = Hx_Enum_loreline_lorscript_exprdef_EBlock{HxMakeArray[any](e2, e16).Dyn()}

			var _hx_tmp_86 int32
			_ = _hx_tmp_86
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_86 = 0
			} else {
				_hx_tmp_86 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin8 HxNullable[int32]
			_ = pmin8
			pmin8 = HxNullable[int32]{Value: _hx_tmp_86, Valid: true}

			var pmax8 HxNullable[int32]
			_ = pmax8
			pmax8 = HxNullable[int32]{}

			var e18 any
			_ = e18
			if e17 == nil {
				e18 = nil
			} else {
				if (pmin8).Valid == false {
					pmin8 = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax8).Valid == false {
					pmax8 = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_87 any
				_ = _hx_tmp_87
				_hx_tmp_87 = ((any)(e17))
				var hx_nullable_f169 HxNullable[int32]
				_ = hx_nullable_f169
				hx_nullable_f169 = pmin8
				var _hx_tmp_88 any
				_ = _hx_tmp_88
				if (hx_nullable_f169).Valid == true {
					_hx_tmp_88 = (hx_nullable_f169).Value
				} else {
					_hx_tmp_88 = nil
				}

				var hx_nullable_f170 HxNullable[int32]
				_ = hx_nullable_f170
				hx_nullable_f170 = pmax8
				var _hx_tmp_89 any
				_ = _hx_tmp_89
				if (hx_nullable_f170).Valid == true {
					_hx_tmp_89 = (hx_nullable_f170).Value
				} else {
					_hx_tmp_89 = nil
				}

				e18 = any(map[string]any{"e": _hx_tmp_87, "pmin": _hx_tmp_88, "pmax": _hx_tmp_89, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(e18)
		}

	case 12:
		{
			var e2 any
			_ = e2
			e2 = this.VTable.Hx_Field_parseExpr()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TBkClose{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EArray{e1, e2}

			var _hx_tmp_90 int32
			_ = _hx_tmp_90
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_90 = 0
			} else {
				_hx_tmp_90 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_90, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{}

			var tmp any
			_ = tmp
			if e == nil {
				tmp = nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_91 any
				_ = _hx_tmp_91
				_hx_tmp_91 = ((any)(e))
				var hx_nullable_f171 HxNullable[int32]
				_ = hx_nullable_f171
				hx_nullable_f171 = pmin
				var _hx_tmp_92 any
				_ = _hx_tmp_92
				if (hx_nullable_f171).Valid == true {
					_hx_tmp_92 = (hx_nullable_f171).Value
				} else {
					_hx_tmp_92 = nil
				}

				var hx_nullable_f172 HxNullable[int32]
				_ = hx_nullable_f172
				hx_nullable_f172 = pmax
				var _hx_tmp_93 any
				_ = _hx_tmp_93
				if (hx_nullable_f172).Valid == true {
					_hx_tmp_93 = (hx_nullable_f172).Value
				} else {
					_hx_tmp_93 = nil
				}

				tmp = any(map[string]any{"e": _hx_tmp_91, "pmin": _hx_tmp_92, "pmax": _hx_tmp_93, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}

			return this.VTable.Hx_Field_parseExprNext(tmp)
		}

	case 14:
		{
			var e2 any
			_ = e2
			e2 = this.VTable.Hx_Field_parseExpr()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var e3 any
			_ = e3
			e3 = this.VTable.Hx_Field_parseExpr()
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_ETernary{e1, e2, e3}

			var _hx_tmp_94 int32
			_ = _hx_tmp_94
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e1, nil))) {
				_hx_tmp_94 = 0
			} else {
				_hx_tmp_94 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e1, "pmin"))))
			}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: _hx_tmp_94, Valid: true}

			var _hx_tmp_95 int32
			_ = _hx_tmp_95
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e3, nil))) {
				_hx_tmp_95 = 0
			} else {
				_hx_tmp_95 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(e3, "pmax"))))
			}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: _hx_tmp_95, Valid: true}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_96 any
				_ = _hx_tmp_96
				_hx_tmp_96 = ((any)(e))
				var hx_nullable_f173 HxNullable[int32]
				_ = hx_nullable_f173
				hx_nullable_f173 = pmin
				var _hx_tmp_97 any
				_ = _hx_tmp_97
				if (hx_nullable_f173).Valid == true {
					_hx_tmp_97 = (hx_nullable_f173).Value
				} else {
					_hx_tmp_97 = nil
				}

				var hx_nullable_f174 HxNullable[int32]
				_ = hx_nullable_f174
				hx_nullable_f174 = pmax
				var _hx_tmp_98 any
				_ = _hx_tmp_98
				if (hx_nullable_f174).Valid == true {
					_hx_tmp_98 = (hx_nullable_f174).Value
				} else {
					_hx_tmp_98 = nil
				}

				return any(map[string]any{"e": _hx_tmp_96, "pmin": _hx_tmp_97, "pmax": _hx_tmp_98, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	default:
		{
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			return e1
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseFunctionArgs() HxArray[any] {
	var args HxArray[any]
	_ = args
	args = HxMakeArray[any]().Dyn()
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	if (tk != Hx_Enum_loreline_lorscript_token_TPClose{}) {
		var done bool
		_ = done
		done = false
		for !done {
			var name HxNullable[string]
			_ = name
			name = HxNullable[string]{Value: HxStringNull}

			var opt bool
			_ = opt
			opt = false
			if tk.Hx_Field_enumIndex() == 14 {
				opt = true
				tk = this.VTable.Hx_Field_token()
			}

			if tk.Hx_Field_enumIndex() == 2 {
				var id string
				_ = id
				id = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
				name = HxNullableString(id)
			} else {
				this.VTable.Hx_Field_unexpected(tk)
				break
			}

			var hx_nullable_f175 HxNullable[string]
			_ = hx_nullable_f175
			hx_nullable_f175 = name
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			if (hx_nullable_f175).Valid == true {
				_hx_tmp_0 = (hx_nullable_f175).Value
			} else {
				_hx_tmp_0 = nil
			}

			var arg any
			_ = arg
			arg = any(map[string]any{"name": _hx_tmp_0})
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](args), arg))
			if opt {
				Hx_Field_go_haxe_hxdynamic_setField(arg, "opt", true)
			}

			if this.Hx_Field_allowTypes {
				if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TDoubleDot{})) {
					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = arg
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "t", this.VTable.Hx_Field_parseType())
				}

				if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TOp{"="})) {
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = arg
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "value", this.VTable.Hx_Field_parseExpr())
				}
			}

			tk = this.VTable.Hx_Field_token()
			switch tk.Hx_Field_enumIndex() {
			case 5:
				{
					done = true
				}

			case 10:
				{
					tk = this.VTable.Hx_Field_token()
				}

			default:
				{
					this.VTable.Hx_Field_unexpected(tk)
				}
			}
		}
	}

	return args
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseFunctionDecl() any {
	{
		var t Hx_Enum_loreline_lorscript_token
		_ = t
		t = this.VTable.Hx_Field_token()
		if (t != Hx_Enum_loreline_lorscript_token_TPOpen{}) {
			this.VTable.Hx_Field_unexpected(t)
		}
	}

	var args HxArray[any]
	_ = args
	args = this.VTable.Hx_Field_parseFunctionArgs()
	var ret Hx_Enum_loreline_lorscript_ctype
	_ = ret
	if this.Hx_Field_allowTypes {
		var tk Hx_Enum_loreline_lorscript_token
		_ = tk
		tk = this.VTable.Hx_Field_token()
		if (tk != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
			this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
			this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
			this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
		} else {
			ret = this.VTable.Hx_Field_parseType()
		}
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(args))
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = ((any)(ret))
	return any(map[string]any{"args": _hx_tmp_0, "ret": _hx_tmp_1, "body": this.VTable.Hx_Field_parseExpr()})
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parsePath() HxArray[string] {
	var path HxArray[string]
	_ = path
	path = HxMakeArray[string](this.VTable.Hx_Field_getIdent())
	for {
		var t Hx_Enum_loreline_lorscript_token
		_ = t
		t = this.VTable.Hx_Field_token()
		if (t != Hx_Enum_loreline_lorscript_token_TDot{}) {
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			break
		}

		var _hx_tmp_0 HxArray[string]
		_ = _hx_tmp_0
		_hx_tmp_0 = path
		_ = (Hx_Array_Push[string](_hx_tmp_0, this.VTable.Hx_Field_getIdent()))
	}

	return path
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseType() Hx_Enum_loreline_lorscript_ctype {
	var t Hx_Enum_loreline_lorscript_token
	_ = t
	t = this.VTable.Hx_Field_token()
	var _gthis *Hx_Obj_loreline_lorscript_parser
	_ = _gthis
	_gthis = this
	switch t.Hx_Field_enumIndex() {
	case 2:
		{
			var v string
			_ = v
			v = t.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var path HxArray[string]
			_ = path
			path = this.VTable.Hx_Field_parsePath()
			var params HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = params
			t = this.VTable.Hx_Field_token()
			if t.Hx_Field_enumIndex() == 3 {
				var op string
				_ = op
				op = t.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
				if op == "<" {
					params = HxMakeArray[Hx_Enum_loreline_lorscript_ctype]()
				hx_label_0:
					for {
						var _hx_tmp_0 HxArray[Hx_Enum_loreline_lorscript_ctype]
						_ = _hx_tmp_0
						_hx_tmp_0 = params
						_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_0, this.VTable.Hx_Field_parseType()))
						t = this.VTable.Hx_Field_token()
						switch t.Hx_Field_enumIndex() {
						case 3:
							{
								var tmp_op_1 string
								_ = tmp_op_1
								tmp_op_1 = t.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
								{
									if tmp_op_1 == ">" {
										break hx_label_0
									}

									var hx_cmp_9 HxNullable[int32]
									_ = hx_cmp_9
									hx_cmp_9 = (Hx_Field_go_haxe_hxstring_charCodeAt(tmp_op_1, 0))
									var _hx_tmp_1 bool
									_ = _hx_tmp_1
									if (hx_cmp_9).Valid == true {
										_hx_tmp_1 = ((hx_cmp_9).Value == 62)
									} else {
										_hx_tmp_1 = false
									}

									if _hx_tmp_1 {
										var _hx_tmp_3 string
										_ = _hx_tmp_3
										_hx_tmp_3 = tmp_op_1
										var _hx_tmp_2 any
										_ = _hx_tmp_2
										_hx_tmp_2 = ((any)(Hx_Enum_loreline_lorscript_token_TOp{(Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 1, HxNullable[int32]{}))}))
										var _hx_tmp_4 int32
										_ = _hx_tmp_4
										_hx_tmp_4 = this.Hx_Field_tokenMax
										this.Hx_Field_tokens.VTable.Hx_Field_add(any(map[string]any{"t": _hx_tmp_2, "min": ((any)(((_hx_tmp_4 - int32(utf8.RuneCountInString(tmp_op_1))) - 1))), "max": ((any)(this.Hx_Field_tokenMax))}))
										break hx_label_0
									}
								}
							}

						case 10:
							{
								continue
							}

						default:
							{
							}
						}

						this.VTable.Hx_Field_unexpected(t)
						break
					}
				} else {
					this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
					this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
					this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
				}
			} else {
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			return this.VTable.Hx_Field_parseTypeNext(Hx_Enum_loreline_lorscript_ctype_CTPath{path, params})
		}

	case 4:
		{
			var a Hx_Enum_loreline_lorscript_token
			_ = a
			a = this.VTable.Hx_Field_token()
			var b Hx_Enum_loreline_lorscript_token
			_ = b
			b = this.VTable.Hx_Field_token()
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(b)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(a)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var withReturn func(args HxArray[Hx_Enum_loreline_lorscript_ctype]) Hx_Enum_loreline_lorscript_ctype
			_ = withReturn
			withReturn = func(args HxArray[Hx_Enum_loreline_lorscript_ctype]) Hx_Enum_loreline_lorscript_ctype {
				{
					var _g Hx_Enum_loreline_lorscript_token
					_ = _g
					_g = _gthis.VTable.Hx_Field_token()
					if _g.Hx_Field_enumIndex() == 3 {
						if _g.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s != "->" {
							var tmp_t_1 Hx_Enum_loreline_lorscript_token
							_ = tmp_t_1
							tmp_t_1 = _g
							_gthis.VTable.Hx_Field_unexpected(tmp_t_1)
						}
					} else {
						var tmp_t_1 Hx_Enum_loreline_lorscript_token
						_ = tmp_t_1
						tmp_t_1 = _g
						_gthis.VTable.Hx_Field_unexpected(tmp_t_1)
					}
				}

				var _hx_tmp_5 HxArray[Hx_Enum_loreline_lorscript_ctype]
				_ = _hx_tmp_5
				_hx_tmp_5 = args
				return Hx_Enum_loreline_lorscript_ctype_CTFun{_hx_tmp_5, _gthis.VTable.Hx_Field_parseType()}
			}

			switch a.Hx_Field_enumIndex() {
			case 2:
				{
					if b.Hx_Field_enumIndex() == 15 {
						var _g HxArray[Hx_Enum_loreline_lorscript_ctype]
						_ = _g
						_g = HxMakeArray[Hx_Enum_loreline_lorscript_ctype]()
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							var _g2 HxArray[any]
							_ = _g2
							_g2 = this.VTable.Hx_Field_parseFunctionArgs()
							for {
								var _hx_tmp_6 int32
								_ = _hx_tmp_6
								_hx_tmp_6 = _g1
								if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
									break
								}

								var arg any
								_ = arg
								arg = _g2.Get(_g1)
								_g1++
								{
									var _g3 any
									_ = _g3
									_g3 = Hx_Field_go_haxe_hxdynamic_getField(arg, "value")
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_g3, nil))) {
										var v any
										_ = v
										v = _g3
										if !this.Hx_Field_resumeErrors {
											panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"Default values not allowed in function types"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(v, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(v, "pmax")))), this.Hx_Field_origin, this.Hx_Field_line)))))
										}
									}
								}

								var _hx_tmp_7 HxArray[Hx_Enum_loreline_lorscript_ctype]
								_ = _hx_tmp_7
								_hx_tmp_7 = _g
								var _hx_tmp_8 string
								_ = _hx_tmp_8
								_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(arg, "name"))
								var _hx_tmp_9 Hx_Enum_loreline_lorscript_ctype
								_ = _hx_tmp_9
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(arg, "opt")))) {
									_hx_tmp_9 = Hx_Enum_loreline_lorscript_ctype_CTOpt{func() Hx_Enum_loreline_lorscript_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
										if _hx_d == nil {
											var _hx_z Hx_Enum_loreline_lorscript_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
									}()}
								} else {
									_hx_tmp_9 = func() Hx_Enum_loreline_lorscript_ctype {
										_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
										if _hx_d == nil {
											var _hx_z Hx_Enum_loreline_lorscript_ctype
											return _hx_z
										}
										return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
									}()
								}

								_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_7, Hx_Enum_loreline_lorscript_ctype_CTNamed{_hx_tmp_8, _hx_tmp_9}))
							}
						}

						var args HxArray[Hx_Enum_loreline_lorscript_ctype]
						_ = args
						args = _g
						return withReturn(args)
					} else {
						var tmp_t_1 Hx_Enum_loreline_lorscript_ctype
						_ = tmp_t_1
						tmp_t_1 = this.VTable.Hx_Field_parseType()
						var _g Hx_Enum_loreline_lorscript_token
						_ = _g
						_g = this.VTable.Hx_Field_token()
						switch _g.Hx_Field_enumIndex() {
						case 5:
							{
								return this.VTable.Hx_Field_parseTypeNext(Hx_Enum_loreline_lorscript_ctype_CTParent{tmp_t_1})
							}

						case 10:
							{
								var args HxArray[Hx_Enum_loreline_lorscript_ctype]
								_ = args
								args = HxMakeArray[Hx_Enum_loreline_lorscript_ctype](tmp_t_1)
								for {
									var _hx_tmp_10 HxArray[Hx_Enum_loreline_lorscript_ctype]
									_ = _hx_tmp_10
									_hx_tmp_10 = args
									_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_10, this.VTable.Hx_Field_parseType()))
									if !(this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TComma{})) {
										break
									}
								}

								{
									var t1 Hx_Enum_loreline_lorscript_token
									_ = t1
									t1 = this.VTable.Hx_Field_token()
									if (t1 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
										this.VTable.Hx_Field_unexpected(t1)
									}
								}

								return withReturn(args)
							}

						default:
							{
								var tmp_t_2 Hx_Enum_loreline_lorscript_token
								_ = tmp_t_2
								tmp_t_2 = _g
								return func() Hx_Enum_loreline_lorscript_ctype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_unexpected(tmp_t_2))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ctype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
								}()
							}
						}
					}
				}

			case 5:
				{
					var _g HxArray[Hx_Enum_loreline_lorscript_ctype]
					_ = _g
					_g = HxMakeArray[Hx_Enum_loreline_lorscript_ctype]()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[any]
						_ = _g2
						_g2 = this.VTable.Hx_Field_parseFunctionArgs()
						for {
							var _hx_tmp_11 int32
							_ = _hx_tmp_11
							_hx_tmp_11 = _g1
							if !(_hx_tmp_11 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
								break
							}

							var arg any
							_ = arg
							arg = _g2.Get(_g1)
							_g1++
							{
								var _g3 any
								_ = _g3
								_g3 = Hx_Field_go_haxe_hxdynamic_getField(arg, "value")
								if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_g3, nil))) {
									var v any
									_ = v
									v = _g3
									if !this.Hx_Field_resumeErrors {
										panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_ECustom{"Default values not allowed in function types"}, ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(v, "pmin")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(v, "pmax")))), this.Hx_Field_origin, this.Hx_Field_line)))))
									}
								}
							}

							var _hx_tmp_12 HxArray[Hx_Enum_loreline_lorscript_ctype]
							_ = _hx_tmp_12
							_hx_tmp_12 = _g
							var _hx_tmp_13 string
							_ = _hx_tmp_13
							_hx_tmp_13 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(arg, "name"))
							var _hx_tmp_14 Hx_Enum_loreline_lorscript_ctype
							_ = _hx_tmp_14
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(arg, "opt")))) {
								_hx_tmp_14 = Hx_Enum_loreline_lorscript_ctype_CTOpt{func() Hx_Enum_loreline_lorscript_ctype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ctype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
								}()}
							} else {
								_hx_tmp_14 = func() Hx_Enum_loreline_lorscript_ctype {
									_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(arg, "t"))
									if _hx_d == nil {
										var _hx_z Hx_Enum_loreline_lorscript_ctype
										return _hx_z
									}
									return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
								}()
							}

							_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_12, Hx_Enum_loreline_lorscript_ctype_CTNamed{_hx_tmp_13, _hx_tmp_14}))
						}
					}

					var args HxArray[Hx_Enum_loreline_lorscript_ctype]
					_ = args
					args = _g
					return withReturn(args)
				}

			default:
				{
					var tmp_t_1 Hx_Enum_loreline_lorscript_ctype
					_ = tmp_t_1
					tmp_t_1 = this.VTable.Hx_Field_parseType()
					var _g Hx_Enum_loreline_lorscript_token
					_ = _g
					_g = this.VTable.Hx_Field_token()
					switch _g.Hx_Field_enumIndex() {
					case 5:
						{
							return this.VTable.Hx_Field_parseTypeNext(Hx_Enum_loreline_lorscript_ctype_CTParent{tmp_t_1})
						}

					case 10:
						{
							var args HxArray[Hx_Enum_loreline_lorscript_ctype]
							_ = args
							args = HxMakeArray[Hx_Enum_loreline_lorscript_ctype](tmp_t_1)
							for {
								var _hx_tmp_15 HxArray[Hx_Enum_loreline_lorscript_ctype]
								_ = _hx_tmp_15
								_hx_tmp_15 = args
								_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_15, this.VTable.Hx_Field_parseType()))
								if !(this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TComma{})) {
									break
								}
							}

							{
								var t1 Hx_Enum_loreline_lorscript_token
								_ = t1
								t1 = this.VTable.Hx_Field_token()
								if (t1 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
									this.VTable.Hx_Field_unexpected(t1)
								}
							}

							return withReturn(args)
						}

					default:
						{
							var tmp_t_2 Hx_Enum_loreline_lorscript_token
							_ = tmp_t_2
							tmp_t_2 = _g
							return func() Hx_Enum_loreline_lorscript_ctype {
								_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_unexpected(tmp_t_2))
								if _hx_d == nil {
									var _hx_z Hx_Enum_loreline_lorscript_ctype
									return _hx_z
								}
								return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
							}()
						}
					}
				}
			}
		}

	case 6:
		{
			var fields HxArray[any]
			_ = fields
			fields = HxMakeArray[any]().Dyn()
			var meta HxArray[any]
			_ = meta
		hx_label_1:
			for {
				t = this.VTable.Hx_Field_token()
				switch t.Hx_Field_enumIndex() {
				case 2:
					{
						var _gs string
						_ = _gs
						_gs = t.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
						if _gs == "var" {
							var name string
							_ = name
							name = this.VTable.Hx_Field_getIdent()
							{
								var tmp_t_1 Hx_Enum_loreline_lorscript_token
								_ = tmp_t_1
								tmp_t_1 = this.VTable.Hx_Field_token()
								if (tmp_t_1 != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
									this.VTable.Hx_Field_unexpected(tmp_t_1)
								}
							}

							var _hx_tmp_16 HxArray[any]
							_ = _hx_tmp_16
							_hx_tmp_16 = HxMakeArrayView[any](fields)
							var _hx_tmp_17 any
							_ = _hx_tmp_17
							_hx_tmp_17 = ((any)(name))
							_ = (Hx_Array_Push[any](_hx_tmp_16, any(map[string]any{"name": _hx_tmp_17, "t": ((any)(this.VTable.Hx_Field_parseType())), "meta": ((any)(meta))})))
							meta = nil
							{
								var t1 Hx_Enum_loreline_lorscript_token
								_ = t1
								t1 = this.VTable.Hx_Field_token()
								if (t1 != Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
									this.VTable.Hx_Field_unexpected(t1)
								}
							}
						} else {
							var name string
							_ = name
							name = _gs
							{
								{
									var t1 Hx_Enum_loreline_lorscript_token
									_ = t1
									t1 = this.VTable.Hx_Field_token()
									if (t1 != Hx_Enum_loreline_lorscript_token_TDoubleDot{}) {
										this.VTable.Hx_Field_unexpected(t1)
									}
								}

								var _hx_tmp_18 HxArray[any]
								_ = _hx_tmp_18
								_hx_tmp_18 = HxMakeArrayView[any](fields)
								var _hx_tmp_19 any
								_ = _hx_tmp_19
								_hx_tmp_19 = ((any)(name))
								_ = (Hx_Array_Push[any](_hx_tmp_18, any(map[string]any{"name": _hx_tmp_19, "t": ((any)(this.VTable.Hx_Field_parseType())), "meta": ((any)(meta))})))
								t = this.VTable.Hx_Field_token()
								switch t.Hx_Field_enumIndex() {
								case 7:
									{
										break hx_label_1
									}

								case 10:
									{
									}

								default:
									{
										this.VTable.Hx_Field_unexpected(t)
									}
								}
							}
						}
					}

				case 7:
					{
						break hx_label_1
					}

				case 16:
					{
						var name string
						_ = name
						name = t.(Hx_Enum_loreline_lorscript_token_TMeta).Hx_Field_s
						{
							if meta == nil {
								meta = HxMakeArray[any]().Dyn()
							}

							var _hx_tmp_20 HxArray[any]
							_ = _hx_tmp_20
							_hx_tmp_20 = HxMakeArrayView[any](meta)
							var _hx_tmp_21 any
							_ = _hx_tmp_21
							_hx_tmp_21 = ((any)(name))
							_ = (Hx_Array_Push[any](_hx_tmp_20, any(map[string]any{"name": _hx_tmp_21, "params": ((any)(this.VTable.Hx_Field_parseMetaArgs()))})))
						}
					}

				default:
					{
						this.VTable.Hx_Field_unexpected(t)
						break hx_label_1
					}
				}
			}

			return this.VTable.Hx_Field_parseTypeNext(Hx_Enum_loreline_lorscript_ctype_CTAnon{fields})
		}

	default:
		{
			return func() Hx_Enum_loreline_lorscript_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_unexpected(t))
				if _hx_d == nil {
					var _hx_z Hx_Enum_loreline_lorscript_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_loreline_lorscript_ctype)
			}()
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseTypeNext(t Hx_Enum_loreline_lorscript_ctype) Hx_Enum_loreline_lorscript_ctype {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	if tk.Hx_Field_enumIndex() == 3 {
		var op string
		_ = op
		op = tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
		if op != "->" {
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			return t
		}
	} else {
		{
			this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
			this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
			this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
		}

		return t
	}

	var t2 Hx_Enum_loreline_lorscript_ctype
	_ = t2
	t2 = this.VTable.Hx_Field_parseType()
	if t2.Hx_Field_enumIndex() == 1 {
		var args HxArray[Hx_Enum_loreline_lorscript_ctype]
		_ = args
		args = t2.(Hx_Enum_loreline_lorscript_ctype_CTFun).Hx_Field_args
		(Hx_Array_Unshift[Hx_Enum_loreline_lorscript_ctype](args, t))
		return t2
	} else {
		return Hx_Enum_loreline_lorscript_ctype_CTFun{HxMakeArray[Hx_Enum_loreline_lorscript_ctype](t), t2}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseExprList(etk Hx_Enum_loreline_lorscript_token) HxArray[any] {
	var args HxArray[any]
	_ = args
	args = HxMakeArray[any]().Dyn()
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	if tk == etk {
		return args
	}

	{
		this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
		this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
		this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
	}

	for {
		var _hx_tmp_0 HxArray[any]
		_ = _hx_tmp_0
		_hx_tmp_0 = HxMakeArrayView[any](args)
		_ = (Hx_Array_Push[any](_hx_tmp_0, this.VTable.Hx_Field_parseExpr()))
		tk = this.VTable.Hx_Field_token()
		if tk.Hx_Field_enumIndex() != 10 {
			if tk == etk {
				break
			}

			this.VTable.Hx_Field_unexpected(tk)
			break
		}
	}

	return args
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseModule(content string, origin__isOptional any, position__isOptional any) HxArray[Hx_Enum_loreline_lorscript_moduledecl] {
	origin := HxDefault[HxNullable[string]]()
	_ = origin

	if origin__isOptional != nil {
		origin = HxConvert[HxNullable[string]](origin__isOptional)
	} else {

	}
	position := HxDefault[HxNullable[int32]]()
	_ = position

	if position__isOptional != nil {
		position = HxConvert[HxNullable[int32]](position__isOptional)
	} else {

	}
	{
		if (origin).Valid == false {
			origin = HxNullableString("loreline.lorscript")
		}

		if (position).Valid == false {
			position = HxNullable[int32]{Value: 0, Valid: true}
		}

		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = (origin).Value
		this.VTable.Hx_Field_initParser(_hx_tmp_0, (position).Value)
		this.Hx_Field_input = content
		this.Hx_Field_readPos = 0
		this.Hx_Field_allowTypes = true
		this.Hx_Field_allowMetadata = true
		var decls HxArray[Hx_Enum_loreline_lorscript_moduledecl]
		_ = decls
		decls = HxMakeArray[Hx_Enum_loreline_lorscript_moduledecl]()
		for {
			var tk Hx_Enum_loreline_lorscript_token
			_ = tk
			tk = this.VTable.Hx_Field_token()
			if (tk == Hx_Enum_loreline_lorscript_token_TEof{}) {
				break
			}

			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			var _hx_tmp_1 HxArray[Hx_Enum_loreline_lorscript_moduledecl]
			_ = _hx_tmp_1
			_hx_tmp_1 = decls
			_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_moduledecl](_hx_tmp_1, this.VTable.Hx_Field_parseModuleDecl()))
		}

		return decls
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseMetadata() HxArray[any] {
	var meta HxArray[any]
	_ = meta
	meta = HxMakeArray[any]().Dyn()
	for {
		var tk Hx_Enum_loreline_lorscript_token
		_ = tk
		tk = this.VTable.Hx_Field_token()
		if tk.Hx_Field_enumIndex() == 16 {
			var name string
			_ = name
			name = tk.(Hx_Enum_loreline_lorscript_token_TMeta).Hx_Field_s
			var _hx_tmp_0 HxArray[any]
			_ = _hx_tmp_0
			_hx_tmp_0 = HxMakeArrayView[any](meta)
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = ((any)(name))
			_ = (Hx_Array_Push[any](_hx_tmp_0, any(map[string]any{"name": _hx_tmp_1, "params": ((any)(this.VTable.Hx_Field_parseMetaArgs()))})))
		} else {
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			break
		}
	}

	return meta
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseParams() any {
	if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TOp{"<"})) {
		if !this.Hx_Field_resumeErrors {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidOp{"Unsupported class type parameters"}, (this.Hx_Field_readPos + this.Hx_Field_offset), (this.Hx_Field_readPos + this.Hx_Field_offset), this.Hx_Field_origin, this.Hx_Field_line)))))
		}
	}

	return any(map[string]any{})
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseModuleDecl() Hx_Enum_loreline_lorscript_moduledecl {
	var meta HxArray[any]
	_ = meta
	meta = this.VTable.Hx_Field_parseMetadata()
	var ident string
	_ = ident
	ident = this.VTable.Hx_Field_getIdent()
	var isPrivate bool
	_ = isPrivate
	isPrivate = false
	var isExtern bool
	_ = isExtern
	isExtern = false
hx_label_0:
	for {
		switch ident {
		case "extern":
			{
				isExtern = true
			}

		case "private":
			{
				isPrivate = true
			}

		default:
			{
				break hx_label_0
			}
		}

		ident = this.VTable.Hx_Field_getIdent()
	}

	switch ident {
	case "class":
		{
			var name string
			_ = name
			name = this.VTable.Hx_Field_getIdent()
			var params any
			_ = params
			params = this.VTable.Hx_Field_parseParams()
			var extend Hx_Enum_loreline_lorscript_ctype
			_ = extend
			var implement HxArray[Hx_Enum_loreline_lorscript_ctype]
			_ = implement
			implement = HxMakeArray[Hx_Enum_loreline_lorscript_ctype]()
		hx_label_1:
			for {
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if t.Hx_Field_enumIndex() == 2 {
					switch t.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s {
					case "extends":
						{
							extend = this.VTable.Hx_Field_parseType()
						}

					case "implements":
						{
							var _hx_tmp_0 HxArray[Hx_Enum_loreline_lorscript_ctype]
							_ = _hx_tmp_0
							_hx_tmp_0 = implement
							_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_ctype](_hx_tmp_0, this.VTable.Hx_Field_parseType()))
						}

					default:
						{
							{
								this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
								this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
								this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
							}

							break hx_label_1
						}
					}
				} else {
					{
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}

					break
				}
			}

			var fields HxArray[any]
			_ = fields
			fields = HxMakeArray[any]().Dyn()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TBrOpen{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			for {
				if !(!this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TBrClose{})) {
					break
				}

				var _hx_tmp_1 HxArray[any]
				_ = _hx_tmp_1
				_hx_tmp_1 = HxMakeArrayView[any](fields)
				_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_parseField()))
			}

			return Hx_Enum_loreline_lorscript_moduledecl_DClass{any(map[string]any{"name": ((any)(name)), "meta": ((any)(meta)), "params": params, "extend": ((any)(extend)), "implement": ((any)(implement)), "fields": ((any)(fields)), "isPrivate": ((any)(isPrivate)), "isExtern": ((any)(isExtern))})}
		}

	case "import":
		{
			var path HxArray[string]
			_ = path
			path = HxMakeArray[string](this.VTable.Hx_Field_getIdent())
			var star bool
			_ = star
			star = false
			for {
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TDot{}) {
					{
						this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(t)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
						this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
						this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
					}

					break
				}

				t = this.VTable.Hx_Field_token()
				switch t.Hx_Field_enumIndex() {
				case 2:
					{
						var id string
						_ = id
						id = t.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
						_ = (Hx_Array_Push[string](path, id))
					}

				case 3:
					{
						if t.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s == "*" {
							star = true
						} else {
							this.VTable.Hx_Field_unexpected(t)
						}
					}

				default:
					{
						this.VTable.Hx_Field_unexpected(t)
					}
				}
			}

			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var _hx_tmp_2 HxArray[string]
			_ = _hx_tmp_2
			_hx_tmp_2 = path
			return Hx_Enum_loreline_lorscript_moduledecl_DImport{_hx_tmp_2, HxNullable[bool]{Value: star, Valid: true}}
		}

	case "package":
		{
			var path HxArray[string]
			_ = path
			path = this.VTable.Hx_Field_parsePath()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				if (t != Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			return Hx_Enum_loreline_lorscript_moduledecl_DPackage{path}
		}

	case "typedef":
		{
			var name string
			_ = name
			name = this.VTable.Hx_Field_getIdent()
			var params any
			_ = params
			params = this.VTable.Hx_Field_parseParams()
			{
				var t Hx_Enum_loreline_lorscript_token
				_ = t
				t = this.VTable.Hx_Field_token()
				var _hx_tmp_3 Hx_Enum_loreline_lorscript_token
				_ = _hx_tmp_3
				_hx_tmp_3 = t
				if (!func() func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
					var _hx_closure_src_324 func(a any, b any) bool
					_ = _hx_closure_src_324
					_hx_closure_src_324 = Hx_Field_type_enumEq
					return func(a Hx_Enum_loreline_lorscript_token, b Hx_Enum_loreline_lorscript_token) bool {
						return _hx_closure_src_324(((any)(a)), ((any)(b)))
					}
				}()(_hx_tmp_3, Hx_Enum_loreline_lorscript_token_TOp{"="})) {
					this.VTable.Hx_Field_unexpected(t)
				}
			}

			var t1 Hx_Enum_loreline_lorscript_ctype
			_ = t1
			t1 = this.VTable.Hx_Field_parseType()
			return Hx_Enum_loreline_lorscript_moduledecl_DTypedef{any(map[string]any{"name": ((any)(name)), "meta": ((any)(meta)), "params": params, "isPrivate": ((any)(isPrivate)), "t": ((any)(t1))})}
		}

	default:
		{
			this.VTable.Hx_Field_unexpected(Hx_Enum_loreline_lorscript_token_TId{ident})
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parseField() any {
	var meta HxArray[any]
	_ = meta
	meta = this.VTable.Hx_Field_parseMetadata()
	var access HxArray[Hx_Enum_loreline_lorscript_fieldaccess]
	_ = access
	access = HxMakeArray[Hx_Enum_loreline_lorscript_fieldaccess]()
hx_label_0:
	for {
		var id string
		_ = id
		id = this.VTable.Hx_Field_getIdent()
		switch id {
		case "function":
			{
				var name string
				_ = name
				name = this.VTable.Hx_Field_getIdent()
				var inf any
				_ = inf
				inf = this.VTable.Hx_Field_parseFunctionDecl()
				var _hx_tmp_0 any
				_ = _hx_tmp_0
				_hx_tmp_0 = ((any)(name))
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				_hx_tmp_1 = ((any)(meta))
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = ((any)(access))
				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(inf, "args")
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(inf, "body")
				return any(map[string]any{"name": _hx_tmp_0, "meta": _hx_tmp_1, "access": _hx_tmp_2, "kind": ((any)(Hx_Enum_loreline_lorscript_fieldkind_KFunction{any(map[string]any{"args": _hx_tmp_3, "expr": _hx_tmp_4, "ret": Hx_Field_go_haxe_hxdynamic_getField(inf, "ret")})}))})
			}

		case "inline":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_AInline{}))
			}

		case "macro":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_AMacro{}))
			}

		case "override":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_AOverride{}))
			}

		case "private":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_APrivate{}))
			}

		case "public":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_APublic{}))
			}

		case "static":
			{
				_ = (Hx_Array_Push[Hx_Enum_loreline_lorscript_fieldaccess](access, Hx_Enum_loreline_lorscript_fieldaccess_AStatic{}))
			}

		case "var":
			{
				var name string
				_ = name
				name = this.VTable.Hx_Field_getIdent()
				var get HxNullable[string]
				_ = get
				get = HxNullable[string]{Value: HxStringNull}

				var set HxNullable[string]
				_ = set
				set = HxNullable[string]{Value: HxStringNull}

				if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TPOpen{})) {
					get = HxNullableString(this.VTable.Hx_Field_getIdent())
					{
						var t Hx_Enum_loreline_lorscript_token
						_ = t
						t = this.VTable.Hx_Field_token()
						if (t != Hx_Enum_loreline_lorscript_token_TComma{}) {
							this.VTable.Hx_Field_unexpected(t)
						}
					}

					set = HxNullableString(this.VTable.Hx_Field_getIdent())
					{
						var t1 Hx_Enum_loreline_lorscript_token
						_ = t1
						t1 = this.VTable.Hx_Field_token()
						if (t1 != Hx_Enum_loreline_lorscript_token_TPClose{}) {
							this.VTable.Hx_Field_unexpected(t1)
						}
					}
				}

				var _hx_tmp_5 Hx_Enum_loreline_lorscript_ctype
				_ = _hx_tmp_5
				if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TDoubleDot{})) {
					_hx_tmp_5 = this.VTable.Hx_Field_parseType()
				} else {
					_hx_tmp_5 = nil
				}

				var _hx_reserved_type Hx_Enum_loreline_lorscript_ctype
				_ = _hx_reserved_type
				_hx_reserved_type = _hx_tmp_5
				var _hx_tmp_6 any
				_ = _hx_tmp_6
				if (this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TOp{"="})) {
					_hx_tmp_6 = this.VTable.Hx_Field_parseExpr()
				} else {
					_hx_tmp_6 = nil
				}

				var expr any
				_ = expr
				expr = _hx_tmp_6
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(expr, nil))) {
					if this.VTable.Hx_Field_isBlock(expr) {
						this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TSemicolon{})
					} else {
						var t Hx_Enum_loreline_lorscript_token
						_ = t
						t = this.VTable.Hx_Field_token()
						if (t != Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
							this.VTable.Hx_Field_unexpected(t)
						}
					}
				} else {
					var _hx_tmp_7 bool
					_ = _hx_tmp_7
					if _hx_reserved_type != nil {
						var _hx_tmp_8 bool
						_ = _hx_tmp_8
						if _hx_reserved_type == nil {
							_hx_tmp_8 = false
						} else {
							_hx_tmp_8 = (_hx_reserved_type.Hx_Field_enumIndex() == 2)
						}

						_hx_tmp_7 = (_hx_tmp_8)
					} else {
						_hx_tmp_7 = false
					}

					if _hx_tmp_7 {
						this.VTable.Hx_Field_maybe(Hx_Enum_loreline_lorscript_token_TSemicolon{})
					} else {
						var t Hx_Enum_loreline_lorscript_token
						_ = t
						t = this.VTable.Hx_Field_token()
						if (t != Hx_Enum_loreline_lorscript_token_TSemicolon{}) {
							this.VTable.Hx_Field_unexpected(t)
						}
					}
				}

				var _hx_tmp_9 any
				_ = _hx_tmp_9
				_hx_tmp_9 = ((any)(name))
				var _hx_tmp_10 any
				_ = _hx_tmp_10
				_hx_tmp_10 = ((any)(meta))
				var _hx_tmp_11 any
				_ = _hx_tmp_11
				_hx_tmp_11 = ((any)(access))
				var hx_nullable_f178 HxNullable[string]
				_ = hx_nullable_f178
				hx_nullable_f178 = get
				var _hx_tmp_12 any
				_ = _hx_tmp_12
				if (hx_nullable_f178).Valid == true {
					_hx_tmp_12 = (hx_nullable_f178).Value
				} else {
					_hx_tmp_12 = nil
				}

				var hx_nullable_f179 HxNullable[string]
				_ = hx_nullable_f179
				hx_nullable_f179 = set
				var _hx_tmp_13 any
				_ = _hx_tmp_13
				if (hx_nullable_f179).Valid == true {
					_hx_tmp_13 = (hx_nullable_f179).Value
				} else {
					_hx_tmp_13 = nil
				}

				return any(map[string]any{"name": _hx_tmp_9, "meta": _hx_tmp_10, "access": _hx_tmp_11, "kind": ((any)(Hx_Enum_loreline_lorscript_fieldkind_KVar{any(map[string]any{"get": _hx_tmp_12, "set": _hx_tmp_13, "type": ((any)(_hx_reserved_type)), "expr": expr})}))})
			}

		default:
			{
				this.VTable.Hx_Field_unexpected(Hx_Enum_loreline_lorscript_token_TId{id})
				break hx_label_0
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_readChar() int32 {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_input
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_readPos
	this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
	return (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, _hx_tmp_1)).Value
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_readString(until int32) string {
	var c int32
	_ = c
	c = 0
	var b *Hx_Obj_stringbuf
	_ = b
	b = Hx_Obj_stringbuf_CreateInstance()
	var esc bool
	_ = esc
	esc = false
	var old int32
	_ = old
	old = this.Hx_Field_line
	var s string
	_ = s
	s = this.Hx_Field_input
	var p1 int32
	_ = p1
	p1 = ((this.Hx_Field_readPos + this.Hx_Field_offset) - 1)
	for {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_input
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_readPos
		this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
		var tmp_c_1 int32
		_ = tmp_c_1
		tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, _hx_tmp_1)).Value
		if tmp_c_1 == 0 {
			this.Hx_Field_line = old
			if !this.Hx_Field_resumeErrors {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EUnterminatedString{}, p1, p1, this.Hx_Field_origin, this.Hx_Field_line)))))
			}

			break
		}

		if esc {
			esc = false
			switch tmp_c_1 {
			case 47:
				{
					if this.Hx_Field_allowJSON {
						b.VTable.Hx_Field_addChar(tmp_c_1)
					} else {
						this.VTable.Hx_Field_invalidChar(tmp_c_1)
					}
				}

			case 34, 39, 92:
				{
					b.VTable.Hx_Field_addChar(tmp_c_1)
				}

			case 110:
				{
					b.VTable.Hx_Field_addChar(10)
				}

			case 114:
				{
					b.VTable.Hx_Field_addChar(13)
				}

			case 116:
				{
					b.VTable.Hx_Field_addChar(9)
				}

			case 117:
				{
					if !this.Hx_Field_allowJSON {
						this.VTable.Hx_Field_invalidChar(tmp_c_1)
					}

					var k int32
					_ = k
					k = 0
					{
						var _g int32
						_ = _g
						_g = 0
						for _g < 4 {
							var _hx_tmp_2 int32
							_ = _hx_tmp_2
							_hx_tmp_2 = _g
							_g = (_g + 1)
							var i int32
							_ = i
							i = _hx_tmp_2
							k = (k << 4)
							var _hx_tmp_3 string
							_ = _hx_tmp_3
							_hx_tmp_3 = this.Hx_Field_input
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = this.Hx_Field_readPos
							this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
							var char int32
							_ = char
							char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_3, _hx_tmp_4)).Value
							switch char {
							case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
								{
									k = (k + (char - 48))
								}

							case 65, 66, 67, 68, 69, 70:
								{
									k = (k + (char - 55))
								}

							case 97, 98, 99, 100, 101, 102:
								{
									k = (k + (char - 87))
								}

							default:
								{
									if char == 0 {
										this.Hx_Field_line = old
										if !this.Hx_Field_resumeErrors {
											panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EUnterminatedString{}, p1, p1, this.Hx_Field_origin, this.Hx_Field_line)))))
										}
									}

									this.VTable.Hx_Field_invalidChar(char)
								}
							}
						}
					}

					b.VTable.Hx_Field_addChar(k)
				}

			default:
				{
					this.VTable.Hx_Field_invalidChar(tmp_c_1)
				}
			}
		} else {
			if tmp_c_1 == 92 {
				esc = true
			} else {
				if tmp_c_1 == until {
					break
				} else {
					if tmp_c_1 == 10 {
						this.Hx_Field_line++
					}

					b.VTable.Hx_Field_addChar(tmp_c_1)
				}
			}
		}
	}

	return b.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_token() Hx_Enum_loreline_lorscript_token {
	var t any
	_ = t
	t = this.Hx_Field_tokens.VTable.Hx_Field_pop()
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(t, nil))) {
		this.Hx_Field_tokenMin = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(t, "min"))))
		this.Hx_Field_tokenMax = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(t, "max"))))
		return func() Hx_Enum_loreline_lorscript_token {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t, "t"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_loreline_lorscript_token
				return _hx_z
			}
			return _hx_d.(Hx_Enum_loreline_lorscript_token)
		}()
	}

	this.Hx_Field_oldTokenMin = this.Hx_Field_tokenMin
	this.Hx_Field_oldTokenMax = this.Hx_Field_tokenMax
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if this.Hx_Field_char < 0 {
		_hx_tmp_0 = (this.Hx_Field_readPos + this.Hx_Field_offset)
	} else {
		_hx_tmp_0 = ((this.Hx_Field_readPos + this.Hx_Field_offset) - 1)
	}

	this.Hx_Field_tokenMin = _hx_tmp_0
	var t1 Hx_Enum_loreline_lorscript_token
	_ = t1
	t1 = this.VTable.Hx_Field__token()
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if this.Hx_Field_char < 0 {
		_hx_tmp_1 = ((this.Hx_Field_readPos + this.Hx_Field_offset) - 1)
	} else {
		_hx_tmp_1 = ((this.Hx_Field_readPos + this.Hx_Field_offset) - 2)
	}

	this.Hx_Field_tokenMax = _hx_tmp_1
	return t1
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field__token() Hx_Enum_loreline_lorscript_token {
	var char int32
	_ = char
	if this.Hx_Field_char < 0 {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_input
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_readPos
		this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
		char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_0, _hx_tmp_1)).Value
	} else {
		char = this.Hx_Field_char
		this.Hx_Field_char = -1
	}

	for {
		if char == 0 {
			this.Hx_Field_char = char
			return Hx_Enum_loreline_lorscript_token_TEof{}
		}

		switch char {
		case 0:
			{
				return Hx_Enum_loreline_lorscript_token_TEof{}
			}

		case 10:
			{
				this.Hx_Field_line++
				this.Hx_Field_tokenMin++
			}

		case 9, 13, 32:
			{
				this.Hx_Field_tokenMin++
			}

		case 34, 39:
			{
				return Hx_Enum_loreline_lorscript_token_TConst{Hx_Enum_loreline_lorscript_const_CString{this.VTable.Hx_Field_readString(char)}}
			}

		case 35:
			{
				var _hx_tmp_2 string
				_ = _hx_tmp_2
				_hx_tmp_2 = this.Hx_Field_input
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_2, _hx_tmp_3)).Value
				if this.Hx_Field_idents.Get(char) {
					var id string
					_ = id
					id = Hx_Field_go_haxe_hxstring_fromCharCode(char)
					for {
						var _hx_tmp_4 string
						_ = _hx_tmp_4
						_hx_tmp_4 = this.Hx_Field_input
						var _hx_tmp_5 int32
						_ = _hx_tmp_5
						_hx_tmp_5 = this.Hx_Field_readPos
						this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
						char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_4, _hx_tmp_5)).Value
						if !this.Hx_Field_idents.Get(char) {
							this.Hx_Field_char = char
							return this.VTable.Hx_Field_preprocess(id)
						}

						var _hx_tmp_6 string
						_ = _hx_tmp_6
						_hx_tmp_6 = id
						id = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_6, Hx_Field_go_haxe_hxstring_fromCharCode(char))
					}
				}

				this.VTable.Hx_Field_invalidChar(char)
			}

		case 40:
			{
				return Hx_Enum_loreline_lorscript_token_TPOpen{}
			}

		case 41:
			{
				return Hx_Enum_loreline_lorscript_token_TPClose{}
			}

		case 44:
			{
				return Hx_Enum_loreline_lorscript_token_TComma{}
			}

		case 46:
			{
				var _hx_tmp_7 string
				_ = _hx_tmp_7
				_hx_tmp_7 = this.Hx_Field_input
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_7, _hx_tmp_8)).Value
				switch char {
				case 46:
					{
						var _hx_tmp_9 string
						_ = _hx_tmp_9
						_hx_tmp_9 = this.Hx_Field_input
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						_hx_tmp_10 = this.Hx_Field_readPos
						this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
						char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_9, _hx_tmp_10)).Value
						if char != 46 {
							this.VTable.Hx_Field_invalidChar(char)
						}

						return Hx_Enum_loreline_lorscript_token_TOp{"..."}
					}

				case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
					{
						var n int32
						_ = n
						n = (char - 48)
						var exp int32
						_ = exp
						exp = 1
						for {
							var _hx_tmp_11 string
							_ = _hx_tmp_11
							_hx_tmp_11 = this.Hx_Field_input
							var _hx_tmp_12 int32
							_ = _hx_tmp_12
							_hx_tmp_12 = this.Hx_Field_readPos
							this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
							char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_11, _hx_tmp_12)).Value
							exp = (exp * 10)
							switch char {
							case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
								{
									n = ((n * 10) + (char - 48))
								}

							default:
								{
									this.Hx_Field_char = char
									return Hx_Enum_loreline_lorscript_token_TConst{Hx_Enum_loreline_lorscript_const_CFloat{((float64)((((float64)(n)) / ((float64)(exp)))))}}
								}
							}
						}
					}

				default:
					{
						this.Hx_Field_char = char
						return Hx_Enum_loreline_lorscript_token_TDot{}
					}
				}
			}

		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			{
				var n float64
				_ = n
				n = (((float64)((char - 48))) * float64(1.0))
				var exp float64
				_ = exp
				exp = float64(0.)
				for {
					var _hx_tmp_13 string
					_ = _hx_tmp_13
					_hx_tmp_13 = this.Hx_Field_input
					var _hx_tmp_14 int32
					_ = _hx_tmp_14
					_hx_tmp_14 = this.Hx_Field_readPos
					this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
					char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_13, _hx_tmp_14)).Value
					exp = (exp * ((float64)(10)))
					switch char {
					case 46:
						{
							if exp > ((float64)(0)) {
								var _hx_tmp_15 bool
								_ = _hx_tmp_15
								if exp == ((float64)(10)) {
									var _hx_tmp_16 string
									_ = _hx_tmp_16
									_hx_tmp_16 = this.Hx_Field_input
									var _hx_tmp_17 int32
									_ = _hx_tmp_17
									_hx_tmp_17 = this.Hx_Field_readPos
									this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
									_hx_tmp_15 = ((Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_16, _hx_tmp_17)).Value == 46)
								} else {
									_hx_tmp_15 = false
								}

								if _hx_tmp_15 {
									{
										this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(Hx_Enum_loreline_lorscript_token_TOp{"..."})), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
										this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
										this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
									}

									var i int32
									_ = i
									i = Hx_Field_std_int(n)
									var _hx_tmp_18 Hx_Enum_loreline_lorscript_const
									_ = _hx_tmp_18
									if ((float64)(i)) == n {
										_hx_tmp_18 = Hx_Enum_loreline_lorscript_const_CInt{i}
									} else {
										_hx_tmp_18 = Hx_Enum_loreline_lorscript_const_CFloat{n}
									}

									return Hx_Enum_loreline_lorscript_token_TConst{_hx_tmp_18}
								}

								this.VTable.Hx_Field_invalidChar(char)
							}

							exp = float64(1.)
						}

					case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
						{
							n = ((n * ((float64)(10))) + ((float64)((char - 48))))
						}

					case 69, 101:
						{
							var tk Hx_Enum_loreline_lorscript_token
							_ = tk
							tk = this.VTable.Hx_Field_token()
							var pow HxNullable[int32]
							_ = pow
							pow = HxNullable[int32]{}

							switch tk.Hx_Field_enumIndex() {
							case 1:
								{
									var _gc Hx_Enum_loreline_lorscript_const
									_ = _gc
									_gc = tk.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
									if _gc.Hx_Field_enumIndex() == 0 {
										var e int32
										_ = e
										e = _gc.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
										pow = HxNullable[int32]{Value: e, Valid: true}
									} else {
										this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
										this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
										this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
									}
								}

							case 3:
								{
									if tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s == "-" {
										tk = this.VTable.Hx_Field_token()
										if tk.Hx_Field_enumIndex() == 1 {
											var _gc Hx_Enum_loreline_lorscript_const
											_ = _gc
											_gc = tk.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
											if _gc.Hx_Field_enumIndex() == 0 {
												var e int32
												_ = e
												e = _gc.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
												pow = HxNullable[int32]{Value: -e, Valid: true}
											} else {
												this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
												this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
												this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
											}
										} else {
											this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
											this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
											this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
										}
									} else {
										this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
										this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
										this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
									}
								}

							default:
								{
									this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
									this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
									this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
								}
							}

							if (pow).Valid == false {
								this.VTable.Hx_Field_invalidChar(char)
							}

							if exp == ((float64)(0)) {
								exp = ((float64)(10))
							}

							return Hx_Enum_loreline_lorscript_token_TConst{Hx_Enum_loreline_lorscript_const_CFloat{((((float64)((Hx_Field_math_pow(((float64)(10)), ((float64)((pow).Value))) / exp))) * n) * ((float64)(10)))}}
						}

					case 120:
						{
							if (n > ((float64)(0))) || (exp > ((float64)(0))) {
								this.VTable.Hx_Field_invalidChar(char)
							}

							var n1 int32
							_ = n1
							n1 = 0
							for {
								var _hx_tmp_19 string
								_ = _hx_tmp_19
								_hx_tmp_19 = this.Hx_Field_input
								var _hx_tmp_20 int32
								_ = _hx_tmp_20
								_hx_tmp_20 = this.Hx_Field_readPos
								this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
								char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_19, _hx_tmp_20)).Value
								switch char {
								case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
									{
										n1 = (((n1 << 4) + char) - 48)
									}

								case 65, 66, 67, 68, 69, 70:
									{
										n1 = ((n1 << 4) + (char - 55))
									}

								case 97, 98, 99, 100, 101, 102:
									{
										n1 = ((n1 << 4) + (char - 87))
									}

								default:
									{
										this.Hx_Field_char = char
										return Hx_Enum_loreline_lorscript_token_TConst{Hx_Enum_loreline_lorscript_const_CInt{n1}}
									}
								}
							}
						}

					default:
						{
							this.Hx_Field_char = char
							var i int32
							_ = i
							i = Hx_Field_std_int(n)
							var _hx_tmp_21 Hx_Enum_loreline_lorscript_const
							_ = _hx_tmp_21
							if exp > ((float64)(0)) {
								_hx_tmp_21 = Hx_Enum_loreline_lorscript_const_CFloat{((float64)(((n * ((float64)(10))) / exp)))}
							} else {
								var _hx_tmp_22 Hx_Enum_loreline_lorscript_const
								_ = _hx_tmp_22
								if ((float64)(i)) == n {
									_hx_tmp_22 = Hx_Enum_loreline_lorscript_const_CInt{i}
								} else {
									_hx_tmp_22 = Hx_Enum_loreline_lorscript_const_CFloat{n}
								}

								_hx_tmp_21 = _hx_tmp_22
							}

							return Hx_Enum_loreline_lorscript_token_TConst{_hx_tmp_21}
						}
					}
				}
			}

		case 58:
			{
				return Hx_Enum_loreline_lorscript_token_TDoubleDot{}
			}

		case 59:
			{
				return Hx_Enum_loreline_lorscript_token_TSemicolon{}
			}

		case 61:
			{
				var _hx_tmp_23 string
				_ = _hx_tmp_23
				_hx_tmp_23 = this.Hx_Field_input
				var _hx_tmp_24 int32
				_ = _hx_tmp_24
				_hx_tmp_24 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_23, _hx_tmp_24)).Value
				if char == 61 {
					return Hx_Enum_loreline_lorscript_token_TOp{"=="}
				} else {
					if char == 62 {
						return Hx_Enum_loreline_lorscript_token_TOp{"=>"}
					}
				}

				this.Hx_Field_char = char
				return Hx_Enum_loreline_lorscript_token_TOp{"="}
			}

		case 63:
			{
				var _hx_tmp_25 string
				_ = _hx_tmp_25
				_hx_tmp_25 = this.Hx_Field_input
				var _hx_tmp_26 int32
				_ = _hx_tmp_26
				_hx_tmp_26 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_25, _hx_tmp_26)).Value
				if char == 46 {
					return Hx_Enum_loreline_lorscript_token_TQuestionDot{}
				}

				this.Hx_Field_char = char
				return Hx_Enum_loreline_lorscript_token_TQuestion{}
			}

		case 64:
			{
				var _hx_tmp_27 string
				_ = _hx_tmp_27
				_hx_tmp_27 = this.Hx_Field_input
				var _hx_tmp_28 int32
				_ = _hx_tmp_28
				_hx_tmp_28 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_27, _hx_tmp_28)).Value
				if this.Hx_Field_idents.Get(char) || (char == 58) {
					var id string
					_ = id
					id = Hx_Field_go_haxe_hxstring_fromCharCode(char)
					for {
						var _hx_tmp_29 string
						_ = _hx_tmp_29
						_hx_tmp_29 = this.Hx_Field_input
						var _hx_tmp_30 int32
						_ = _hx_tmp_30
						_hx_tmp_30 = this.Hx_Field_readPos
						this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
						char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_29, _hx_tmp_30)).Value
						if !this.Hx_Field_idents.Get(char) {
							this.Hx_Field_char = char
							return Hx_Enum_loreline_lorscript_token_TMeta{id}
						}

						var _hx_tmp_31 string
						_ = _hx_tmp_31
						_hx_tmp_31 = id
						id = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_31, Hx_Field_go_haxe_hxstring_fromCharCode(char))
					}
				}

				this.VTable.Hx_Field_invalidChar(char)
			}

		case 91:
			{
				return Hx_Enum_loreline_lorscript_token_TBkOpen{}
			}

		case 93:
			{
				return Hx_Enum_loreline_lorscript_token_TBkClose{}
			}

		case 123:
			{
				return Hx_Enum_loreline_lorscript_token_TBrOpen{}
			}

		case 125:
			{
				return Hx_Enum_loreline_lorscript_token_TBrClose{}
			}

		default:
			{
				if this.Hx_Field_ops.Get(char) {
					var op string
					_ = op
					op = Hx_Field_go_haxe_hxstring_fromCharCode(char)
					for {
						var _hx_tmp_32 string
						_ = _hx_tmp_32
						_hx_tmp_32 = this.Hx_Field_input
						var _hx_tmp_33 int32
						_ = _hx_tmp_33
						_hx_tmp_33 = this.Hx_Field_readPos
						this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
						char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_32, _hx_tmp_33)).Value
						if char == 0 {
							char = 0
						}

						if !this.Hx_Field_ops.Get(char) {
							this.Hx_Field_char = char
							return Hx_Enum_loreline_lorscript_token_TOp{op}
						}

						var pop string
						_ = pop
						pop = op
						var _hx_tmp_34 string
						_ = _hx_tmp_34
						_hx_tmp_34 = op
						op = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_34, Hx_Field_go_haxe_hxstring_fromCharCode(char))
						var _hx_tmp_35 bool
						_ = _hx_tmp_35
						if !func() func(_hx_param_key string) bool {
							var _hx_closure_src_325 func(_hx_param_key any) bool
							_ = _hx_closure_src_325
							_hx_closure_src_325 = this.Hx_Field_opPriority.VTable.(interface {
								Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
							}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
							return func(_hx_param_key string) bool {
								return _hx_closure_src_325(((any)(_hx_param_key)))
							}
						}()(op) {
							_hx_tmp_35 = func() func(_hx_param_key string) bool {
								var _hx_closure_src_326 func(_hx_param_key any) bool
								_ = _hx_closure_src_326
								_hx_closure_src_326 = this.Hx_Field_opPriority.VTable.(interface {
									Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
								}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
								return func(_hx_param_key string) bool {
									return _hx_closure_src_326(((any)(_hx_param_key)))
								}
							}()(pop)
						} else {
							_hx_tmp_35 = false
						}

						if _hx_tmp_35 {
							if (op == "//") || (op == "/*") {
								return this.VTable.Hx_Field_tokenComment(op, char)
							}

							this.Hx_Field_char = char
							return Hx_Enum_loreline_lorscript_token_TOp{pop}
						}
					}
				}

				if this.Hx_Field_idents.Get(char) {
					var id string
					_ = id
					id = Hx_Field_go_haxe_hxstring_fromCharCode(char)
					for {
						var _hx_tmp_36 string
						_ = _hx_tmp_36
						_hx_tmp_36 = this.Hx_Field_input
						var _hx_tmp_37 int32
						_ = _hx_tmp_37
						_hx_tmp_37 = this.Hx_Field_readPos
						this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
						char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_36, _hx_tmp_37)).Value
						if char == 0 {
							char = 0
						}

						if !this.Hx_Field_idents.Get(char) {
							this.Hx_Field_char = char
							return Hx_Enum_loreline_lorscript_token_TId{id}
						}

						var _hx_tmp_38 string
						_ = _hx_tmp_38
						_hx_tmp_38 = id
						id = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_38, Hx_Field_go_haxe_hxstring_fromCharCode(char))
					}
				}

				this.VTable.Hx_Field_invalidChar(char)
			}
		}

		var _hx_tmp_39 string
		_ = _hx_tmp_39
		_hx_tmp_39 = this.Hx_Field_input
		var _hx_tmp_40 int32
		_ = _hx_tmp_40
		_hx_tmp_40 = this.Hx_Field_readPos
		this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
		char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_39, _hx_tmp_40)).Value
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_preprocValue(id string) any {
	return this.Hx_Field_preprocesorValues.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(id)
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_parsePreproCond() any {
	var tk Hx_Enum_loreline_lorscript_token
	_ = tk
	tk = this.VTable.Hx_Field_token()
	switch tk.Hx_Field_enumIndex() {
	case 2:
		{
			var id string
			_ = id
			id = tk.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
			var e Hx_Enum_loreline_lorscript_exprdef
			_ = e
			e = Hx_Enum_loreline_lorscript_exprdef_EIdent{id}

			var pmin HxNullable[int32]
			_ = pmin
			pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

			var pmax HxNullable[int32]
			_ = pmax
			pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}

			if e == nil {
				return nil
			} else {
				if (pmin).Valid == false {
					pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
				}

				if (pmax).Valid == false {
					pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
				}

				var _hx_tmp_0 any
				_ = _hx_tmp_0
				_hx_tmp_0 = ((any)(e))
				var hx_nullable_f180 HxNullable[int32]
				_ = hx_nullable_f180
				hx_nullable_f180 = pmin
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				if (hx_nullable_f180).Valid == true {
					_hx_tmp_1 = (hx_nullable_f180).Value
				} else {
					_hx_tmp_1 = nil
				}

				var hx_nullable_f181 HxNullable[int32]
				_ = hx_nullable_f181
				hx_nullable_f181 = pmax
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				if (hx_nullable_f181).Valid == true {
					_hx_tmp_2 = (hx_nullable_f181).Value
				} else {
					_hx_tmp_2 = nil
				}

				return any(map[string]any{"e": _hx_tmp_0, "pmin": _hx_tmp_1, "pmax": _hx_tmp_2, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
			}
		}

	case 3:
		{
			if tk.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s == "!" {
				var e Hx_Enum_loreline_lorscript_exprdef
				_ = e
				e = Hx_Enum_loreline_lorscript_exprdef_EUnop{"!", true, this.VTable.Hx_Field_parsePreproCond()}

				var pmin HxNullable[int32]
				_ = pmin
				pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}

				var pmax HxNullable[int32]
				_ = pmax
				pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}

				if e == nil {
					return nil
				} else {
					if (pmin).Valid == false {
						pmin = HxNullable[int32]{Value: this.Hx_Field_tokenMin, Valid: true}
					}

					if (pmax).Valid == false {
						pmax = HxNullable[int32]{Value: this.Hx_Field_tokenMax, Valid: true}
					}

					var _hx_tmp_3 any
					_ = _hx_tmp_3
					_hx_tmp_3 = ((any)(e))
					var hx_nullable_f182 HxNullable[int32]
					_ = hx_nullable_f182
					hx_nullable_f182 = pmin
					var _hx_tmp_4 any
					_ = _hx_tmp_4
					if (hx_nullable_f182).Valid == true {
						_hx_tmp_4 = (hx_nullable_f182).Value
					} else {
						_hx_tmp_4 = nil
					}

					var hx_nullable_f183 HxNullable[int32]
					_ = hx_nullable_f183
					hx_nullable_f183 = pmax
					var _hx_tmp_5 any
					_ = _hx_tmp_5
					if (hx_nullable_f183).Valid == true {
						_hx_tmp_5 = (hx_nullable_f183).Value
					} else {
						_hx_tmp_5 = nil
					}

					return any(map[string]any{"e": _hx_tmp_3, "pmin": _hx_tmp_4, "pmax": _hx_tmp_5, "origin": ((any)(this.Hx_Field_origin)), "line": ((any)(this.Hx_Field_line))})
				}
			} else {
				return this.VTable.Hx_Field_unexpected(tk)
			}
		}

	case 4:
		{
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(Hx_Enum_loreline_lorscript_token_TPOpen{})), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			return this.VTable.Hx_Field_parseExpr()
		}

	default:
		{
			return this.VTable.Hx_Field_unexpected(tk)
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_evalPreproCond(e any) bool {
	var _g Hx_Enum_loreline_lorscript_exprdef
	_ = _g
	_g = func() Hx_Enum_loreline_lorscript_exprdef {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_loreline_lorscript_exprdef
			return _hx_z
		}
		return _hx_d.(Hx_Enum_loreline_lorscript_exprdef)
	}()
	switch _g.Hx_Field_enumIndex() {
	case 1:
		{
			var id string
			_ = id
			id = _g.(Hx_Enum_loreline_lorscript_exprdef_EIdent).Hx_Field_v
			return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.VTable.Hx_Field_preprocValue(id), nil))))
		}

	case 3:
		{
			var e any
			_ = e
			e = _g.(Hx_Enum_loreline_lorscript_exprdef_EParent).Hx_Field_e
			return this.VTable.Hx_Field_evalPreproCond(e)
		}

	case 6:
		{
			var _ge any
			_ = _ge
			_ge = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e1
			var _ge1 any
			_ = _ge1
			_ge1 = _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_e2
			switch _g.(Hx_Enum_loreline_lorscript_exprdef_EBinop).Hx_Field_op {
			case "&&":
				{
					var e1 any
					_ = e1
					e1 = _ge
					var e2 any
					_ = e2
					e2 = _ge1
					var _hx_tmp_0 bool
					_ = _hx_tmp_0
					if this.VTable.Hx_Field_evalPreproCond(e1) {
						_hx_tmp_0 = this.VTable.Hx_Field_evalPreproCond(e2)
					} else {
						_hx_tmp_0 = false
					}

					return _hx_tmp_0
				}

			case "||":
				{
					var e1 any
					_ = e1
					e1 = _ge
					var e2 any
					_ = e2
					e2 = _ge1
					var _hx_tmp_1 bool
					_ = _hx_tmp_1
					if this.VTable.Hx_Field_evalPreproCond(e1) {
						_hx_tmp_1 = true
					} else {
						_hx_tmp_1 = this.VTable.Hx_Field_evalPreproCond(e2)
					}

					return _hx_tmp_1
				}

			default:
				{
					{
						var err Hx_Enum_loreline_lorscript_errordef
						_ = err
						err = Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor{Hx_Field_go_haxe_hxdynamic_stringConcat("Can't eval ", Hx_Field_type_enumConstructor(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
							_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
							if _hx_d == nil {
								var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
								return _hx_z
							}
							return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
						}()))}

						if !this.Hx_Field_resumeErrors {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(err, (this.Hx_Field_readPos + this.Hx_Field_offset), (this.Hx_Field_readPos + this.Hx_Field_offset), this.Hx_Field_origin, this.Hx_Field_line)))))
						}
					}

					return false
				}
			}
		}

	case 7:
		{
			if _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_op == "!" {
				var e any
				_ = e
				e = _g.(Hx_Enum_loreline_lorscript_exprdef_EUnop).Hx_Field_e
				return !this.VTable.Hx_Field_evalPreproCond(e)
			} else {
				{
					var err Hx_Enum_loreline_lorscript_errordef
					_ = err
					err = Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor{Hx_Field_go_haxe_hxdynamic_stringConcat("Can't eval ", Hx_Field_type_enumConstructor(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
						_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
						if _hx_d == nil {
							var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
							return _hx_z
						}
						return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
					}()))}

					if !this.Hx_Field_resumeErrors {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(err, (this.Hx_Field_readPos + this.Hx_Field_offset), (this.Hx_Field_readPos + this.Hx_Field_offset), this.Hx_Field_origin, this.Hx_Field_line)))))
					}
				}

				return false
			}
		}

	default:
		{
			{
				var err Hx_Enum_loreline_lorscript_errordef
				_ = err
				err = Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor{Hx_Field_go_haxe_hxdynamic_stringConcat("Can't eval ", Hx_Field_type_enumConstructor(func() Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
					_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(e, "e"))
					if _hx_d == nil {
						var _hx_z Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
						return _hx_z
					}
					return _hx_d.(Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)
				}()))}

				if !this.Hx_Field_resumeErrors {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(err, (this.Hx_Field_readPos + this.Hx_Field_offset), (this.Hx_Field_readPos + this.Hx_Field_offset), this.Hx_Field_origin, this.Hx_Field_line)))))
				}
			}

			return false
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_preprocess(id string) Hx_Enum_loreline_lorscript_token {
	switch id {
	case "else", "elseif":
		{
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_preprocStack) > 0 {
				var _hx_tmp_0 HxArray[any]
				_ = _hx_tmp_0
				_hx_tmp_0 = this.Hx_Field_preprocStack
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(_hx_tmp_0.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_preprocStack) - 1)), "r")))) {
					var _hx_tmp_1 HxArray[any]
					_ = _hx_tmp_1
					_hx_tmp_1 = this.Hx_Field_preprocStack
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_preprocStack) - 1)), "r", false)
					this.VTable.Hx_Field_skipTokens()
					return this.VTable.Hx_Field_token()
				} else {
					if id == "else" {
						_ = (Hx_Array_Pop[any](HxMakeArrayView[any](this.Hx_Field_preprocStack)).Value)
						var _hx_tmp_2 HxArray[any]
						_ = _hx_tmp_2
						_hx_tmp_2 = HxMakeArrayView[any](this.Hx_Field_preprocStack)
						_ = (Hx_Array_Push[any](_hx_tmp_2, any(map[string]any{"r": ((any)(true))})))
						return this.VTable.Hx_Field_token()
					} else {
						_ = (Hx_Array_Pop[any](HxMakeArrayView[any](this.Hx_Field_preprocStack)).Value)
						return this.VTable.Hx_Field_preprocess("if")
					}
				}
			} else {
				return Hx_Enum_loreline_lorscript_token_TPrepro{id}
			}
		}

	case "end":
		{
			if Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_preprocStack) > 0 {
				_ = (Hx_Array_Pop[any](HxMakeArrayView[any](this.Hx_Field_preprocStack)).Value)
				return this.VTable.Hx_Field_token()
			} else {
				return Hx_Enum_loreline_lorscript_token_TPrepro{id}
			}
		}

	case "if":
		{
			var e any
			_ = e
			e = this.VTable.Hx_Field_parsePreproCond()
			if this.VTable.Hx_Field_evalPreproCond(e) {
				var _hx_tmp_3 HxArray[any]
				_ = _hx_tmp_3
				_hx_tmp_3 = HxMakeArrayView[any](this.Hx_Field_preprocStack)
				_ = (Hx_Array_Push[any](_hx_tmp_3, any(map[string]any{"r": ((any)(true))})))
				return this.VTable.Hx_Field_token()
			}

			var _hx_tmp_4 HxArray[any]
			_ = _hx_tmp_4
			_hx_tmp_4 = HxMakeArrayView[any](this.Hx_Field_preprocStack)
			_ = (Hx_Array_Push[any](_hx_tmp_4, any(map[string]any{"r": ((any)(false))})))
			this.VTable.Hx_Field_skipTokens()
			return this.VTable.Hx_Field_token()
		}

	default:
		{
			return Hx_Enum_loreline_lorscript_token_TPrepro{id}
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_skipTokens() {
	var spos int32
	_ = spos
	spos = (Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_preprocStack) - 1)
	var obj any
	_ = obj
	obj = this.Hx_Field_preprocStack.Get(spos)
	var pos int32
	_ = pos
	pos = (this.Hx_Field_readPos + this.Hx_Field_offset)
	for {
		var tk Hx_Enum_loreline_lorscript_token
		_ = tk
		tk = this.VTable.Hx_Field_token()
		if (tk == Hx_Enum_loreline_lorscript_token_TEof{}) {
			if !this.Hx_Field_resumeErrors {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EInvalidPreprocessor{"Unclosed"}, pos, pos, this.Hx_Field_origin, this.Hx_Field_line)))))
			}
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_preprocStack.Get(spos), obj))) {
			{
				this.Hx_Field_tokens.VTable.Hx_Field_push(any(map[string]any{"t": ((any)(tk)), "min": ((any)(this.Hx_Field_tokenMin)), "max": ((any)(this.Hx_Field_tokenMax))}))
				this.Hx_Field_tokenMin = this.Hx_Field_oldTokenMin
				this.Hx_Field_tokenMax = this.Hx_Field_oldTokenMax
			}

			break
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_tokenComment(op string, char int32) Hx_Enum_loreline_lorscript_token {
	var c HxNullable[int32]
	_ = c
	c = (Hx_Field_go_haxe_hxstring_charCodeAt(op, 1))
	var s string
	_ = s
	s = this.Hx_Field_input
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (c).Valid == true {
		_hx_tmp_0 = ((c).Value == 47)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		for (char != 13) && (char != 10) {
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_input
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = this.Hx_Field_readPos
			this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
			char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, _hx_tmp_2)).Value
			if char == 0 {
				break
			}
		}

		this.Hx_Field_char = char
		return this.VTable.Hx_Field_token()
	}

	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if (c).Valid == true {
		_hx_tmp_3 = ((c).Value == 42)
	} else {
		_hx_tmp_3 = false
	}

	if _hx_tmp_3 {
		var old int32
		_ = old
		old = this.Hx_Field_line
		if op == "/**/" {
			this.Hx_Field_char = char
			return this.VTable.Hx_Field_token()
		}

		for {
			for char != 42 {
				if char == 10 {
					this.Hx_Field_line++
				}

				var _hx_tmp_4 string
				_ = _hx_tmp_4
				_hx_tmp_4 = this.Hx_Field_input
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = this.Hx_Field_readPos
				this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
				char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_4, _hx_tmp_5)).Value
				if char == 0 {
					this.Hx_Field_line = old
					if !this.Hx_Field_resumeErrors {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment{}, this.Hx_Field_tokenMin, this.Hx_Field_tokenMin, this.Hx_Field_origin, this.Hx_Field_line)))))
					}

					break
				}
			}

			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = this.Hx_Field_input
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = this.Hx_Field_readPos
			this.Hx_Field_readPos = (this.Hx_Field_readPos + 1)
			char = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_6, _hx_tmp_7)).Value
			if char == 0 {
				this.Hx_Field_line = old
				if !this.Hx_Field_resumeErrors {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_lorscript_error_CreateInstance(Hx_Enum_loreline_lorscript_errordef_EUnterminatedComment{}, this.Hx_Field_tokenMin, this.Hx_Field_tokenMin, this.Hx_Field_origin, this.Hx_Field_line)))))
				}

				break
			}

			if !(char != 47) {
				break
			}
		}

		return this.VTable.Hx_Field_token()
	}

	this.Hx_Field_char = char
	return Hx_Enum_loreline_lorscript_token_TOp{op}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_constString(c Hx_Enum_loreline_lorscript_const) string {
	switch c.Hx_Field_enumIndex() {
	case 0:
		{
			var v int32
			_ = v
			v = c.(Hx_Enum_loreline_lorscript_const_CInt).Hx_Field_v
			return Hx_Field_std_string(v)
		}

	case 1:
		{
			var f float64
			_ = f
			f = c.(Hx_Enum_loreline_lorscript_const_CFloat).Hx_Field_f
			return Hx_Field_std_string(f)
		}

	case 2:
		{
			var s string
			_ = s
			s = c.(Hx_Enum_loreline_lorscript_const_CString).Hx_Field_s
			return s
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_lorscript_parser) Hx_Field_tokenString(t Hx_Enum_loreline_lorscript_token) string {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			return "<eof>"
		}

	case 1:
		{
			var c Hx_Enum_loreline_lorscript_const
			_ = c
			c = t.(Hx_Enum_loreline_lorscript_token_TConst).Hx_Field_c
			return this.VTable.Hx_Field_constString(c)
		}

	case 2:
		{
			var s string
			_ = s
			s = t.(Hx_Enum_loreline_lorscript_token_TId).Hx_Field_s
			return s
		}

	case 3:
		{
			var s string
			_ = s
			s = t.(Hx_Enum_loreline_lorscript_token_TOp).Hx_Field_s
			return s
		}

	case 4:
		{
			return "("
		}

	case 5:
		{
			return ")"
		}

	case 6:
		{
			return "{"
		}

	case 7:
		{
			return "}"
		}

	case 8:
		{
			return "."
		}

	case 9:
		{
			return "?."
		}

	case 10:
		{
			return ","
		}

	case 11:
		{
			return ";"
		}

	case 12:
		{
			return "["
		}

	case 13:
		{
			return "]"
		}

	case 14:
		{
			return "?"
		}

	case 15:
		{
			return ":"
		}

	case 16:
		{
			var id string
			_ = id
			id = t.(Hx_Enum_loreline_lorscript_token_TMeta).Hx_Field_s
			return Hx_Field_go_haxe_hxdynamic_stringConcat("@", id)
		}

	case 17:
		{
			var id string
			_ = id
			id = t.(Hx_Enum_loreline_lorscript_token_TPrepro).Hx_Field_s
			return Hx_Field_go_haxe_hxdynamic_stringConcat("#", id)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func init() {
	hxRegisterArrayConstructor[HxArray[string]]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_ctype]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_fieldaccess]()
	hxRegisterArrayConstructor[Hx_Enum_loreline_lorscript_moduledecl]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[bool]()
	hxRegisterArrayConstructor[string]()
}
