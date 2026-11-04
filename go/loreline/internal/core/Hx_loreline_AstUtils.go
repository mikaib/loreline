package core

import "unicode/utf8"
import "math"

func Hx_ClassInit_Hx_Obj_VTable_loreline_astutils() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_astutils() {
	Hx_Field_loreline_astutils_randomId_chars = Hx_Init_Hx_Field_loreline_astutils_randomId_chars()
	Hx_Field_loreline_astutils_randomId_nonHexChars = Hx_Init_Hx_Field_loreline_astutils_randomId_nonHexChars()
}

type Hx_Obj_VTable_loreline_astutils interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_astutils struct {
	VTable Hx_Obj_VTable_loreline_astutils `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_astutils_CreateEmptyInstance() *Hx_Obj_loreline_astutils {
	obj := &Hx_Obj_loreline_astutils{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_astutils_CreateInstance() *Hx_Obj_loreline_astutils {
	obj := Hx_Obj_loreline_astutils_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_astutils_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.AstUtils",
	HxMakeArray[string]("randomId_chars", "randomId_nonHexChars", "addQuotes", "removeQuotes", "toBraceStyle", "toPlainStyle", "useWordOperators", "useSymbolOperators", "setWordOperators", "whenRuleScore", "andClauseCount", "isLiteralTrue", "useParenConditions", "usePlainConditions", "reindentFunctions", "insertLocalizationKeys", "removeLocalizationKeys", "lineHasHashComment", "extractTranslations", "extractTranslatableEntries", "hasUntaggedTranslatableStrings", "collectHashIds", "generateTranslationFile", "unquotedRawToQuotedRaw", "quotedRawToUnquotedRaw", "canSafelyRemoveQuotes", "isNumberLiteral", "reindentCode", "getStringEndLine", "findHashComment", "randomId"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_astutils_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_astutils_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_astutils_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "randomId_chars":
			return &Hx_Field_loreline_astutils_randomId_chars
		case "randomId_nonHexChars":
			return &Hx_Field_loreline_astutils_randomId_nonHexChars
		case "addQuotes":
			return Hx_Field_loreline_astutils_addQuotes
		case "removeQuotes":
			return Hx_Field_loreline_astutils_removeQuotes
		case "toBraceStyle":
			return Hx_Field_loreline_astutils_toBraceStyle
		case "toPlainStyle":
			return Hx_Field_loreline_astutils_toPlainStyle
		case "useWordOperators":
			return Hx_Field_loreline_astutils_useWordOperators
		case "useSymbolOperators":
			return Hx_Field_loreline_astutils_useSymbolOperators
		case "setWordOperators":
			return Hx_Field_loreline_astutils_setWordOperators
		case "whenRuleScore":
			return Hx_Field_loreline_astutils_whenRuleScore
		case "andClauseCount":
			return Hx_Field_loreline_astutils_andClauseCount
		case "isLiteralTrue":
			return Hx_Field_loreline_astutils_isLiteralTrue
		case "useParenConditions":
			return Hx_Field_loreline_astutils_useParenConditions
		case "usePlainConditions":
			return Hx_Field_loreline_astutils_usePlainConditions
		case "reindentFunctions":
			return Hx_Field_loreline_astutils_reindentFunctions
		case "insertLocalizationKeys":
			return Hx_Field_loreline_astutils_insertLocalizationKeys
		case "removeLocalizationKeys":
			return Hx_Field_loreline_astutils_removeLocalizationKeys
		case "lineHasHashComment":
			return Hx_Field_loreline_astutils_lineHasHashComment
		case "extractTranslations":
			return Hx_Field_loreline_astutils_extractTranslations
		case "extractTranslatableEntries":
			return Hx_Field_loreline_astutils_extractTranslatableEntries
		case "hasUntaggedTranslatableStrings":
			return Hx_Field_loreline_astutils_hasUntaggedTranslatableStrings
		case "collectHashIds":
			return Hx_Field_loreline_astutils_collectHashIds
		case "generateTranslationFile":
			return Hx_Field_loreline_astutils_generateTranslationFile
		case "unquotedRawToQuotedRaw":
			return Hx_Field_loreline_astutils_unquotedRawToQuotedRaw
		case "quotedRawToUnquotedRaw":
			return Hx_Field_loreline_astutils_quotedRawToUnquotedRaw
		case "canSafelyRemoveQuotes":
			return Hx_Field_loreline_astutils_canSafelyRemoveQuotes
		case "isNumberLiteral":
			return Hx_Field_loreline_astutils_isNumberLiteral
		case "reindentCode":
			return Hx_Field_loreline_astutils_reindentCode
		case "getStringEndLine":
			return Hx_Field_loreline_astutils_getStringEndLine
		case "findHashComment":
			return Hx_Field_loreline_astutils_findHashComment
		case "randomId":
			return Hx_Field_loreline_astutils_randomId
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_astutils) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_astutils_RTTI
}

func (this *Hx_Obj_loreline_astutils) Hx_As_Hx_Obj_loreline_astutils() *Hx_Obj_loreline_astutils {
	return this
}

var Hx_Field_loreline_astutils_randomId_chars string

func Hx_Init_Hx_Field_loreline_astutils_randomId_chars() string {
	return "0123456789abcdefghijklmnopqrstuvwxyz"
}

var Hx_Field_loreline_astutils_randomId_nonHexChars string

func Hx_Init_Hx_Field_loreline_astutils_randomId_nonHexChars() string {
	return "ghijklmnopqrstuvwxyz"
}

func Hx_Field_loreline_astutils_addQuotes(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nstringliteral_RTTI) {
			var str *Hx_Obj_loreline_nstringliteral
			_ = str
			str = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
			}).Hx_As_Hx_Obj_loreline_nstringliteral()
			if str.Hx_Field_quotes == ((int32)(0)) {
				str.Hx_Field_quotes = ((int32)(1))
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = str.Hx_Field_parts.Len()
					for _g < _g1 {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g
						_g = (_g + 1)
						var i int32
						_ = i
						i = _hx_tmp_0
						{
							var tmp__g1_1 Hx_Enum_loreline_stringparttype
							_ = tmp__g1_1
							tmp__g1_1 = str.Hx_Field_parts.Get(i).Hx_Field_partType
							if tmp__g1_1.Hx_Field_enumIndex() == 0 {
								var text string
								_ = text
								text = tmp__g1_1.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
								var tmp string
								_ = tmp
								tmp = Hx_Field_loreline_astutils_unquotedRawToQuotedRaw(text)
								str.Hx_Field_parts.Get(i).Hx_Field_partType = Hx_Enum_loreline_stringparttype_Raw{tmp}
							}
						}
					}
				}
			}
		}
	})
}

func Hx_Field_loreline_astutils_removeQuotes(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nstringliteral_RTTI) {
			var str *Hx_Obj_loreline_nstringliteral
			_ = str
			str = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
			}).Hx_As_Hx_Obj_loreline_nstringliteral()
			if str.Hx_Field_quotes == ((int32)(1)) {
				var fullText string
				_ = fullText
				fullText = ""
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
								fullText = Hx_Field_go_haxe_hxdynamic_stringConcat(fullText, text)
							}
						}
					}
				}

				if Hx_Field_loreline_astutils_canSafelyRemoveQuotes(fullText) {
					str.Hx_Field_quotes = ((int32)(0))
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = str.Hx_Field_parts.Len()
						for _g < _g1 {
							var _hx_tmp_1 int32
							_ = _hx_tmp_1
							_hx_tmp_1 = _g
							_g = (_g + 1)
							var i int32
							_ = i
							i = _hx_tmp_1
							{
								var tmp__g1_1 Hx_Enum_loreline_stringparttype
								_ = tmp__g1_1
								tmp__g1_1 = str.Hx_Field_parts.Get(i).Hx_Field_partType
								if tmp__g1_1.Hx_Field_enumIndex() == 0 {
									var text string
									_ = text
									text = tmp__g1_1.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
									var tmp string
									_ = tmp
									tmp = Hx_Field_loreline_astutils_quotedRawToUnquotedRaw(text)
									str.Hx_Field_parts.Get(i).Hx_Field_partType = Hx_Enum_loreline_stringparttype_Raw{tmp}
								}
							}
						}
					}
				}
			}
		}
	})
}

func Hx_Field_loreline_astutils_toBraceStyle(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nstatedecl_RTTI) {
			var hx_cast_ex_96 *Hx_Obj_loreline_node
			_ = hx_cast_ex_96
			hx_cast_ex_96 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_96, Hx_Obj_loreline_nstatedecl_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			(hx_cast_ex_96.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
			}).Hx_As_Hx_Obj_loreline_nstatedecl()).Hx_Field_style = ((int32)(1))
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ncharacterdecl_RTTI) {
				var hx_cast_ex_97 *Hx_Obj_loreline_node
				_ = hx_cast_ex_97
				hx_cast_ex_97 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_97, Hx_Obj_loreline_ncharacterdecl_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				(hx_cast_ex_97.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
				}).Hx_As_Hx_Obj_loreline_ncharacterdecl()).Hx_Field_style = ((int32)(1))
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nbeatdecl_RTTI) {
					var hx_cast_ex_98 *Hx_Obj_loreline_node
					_ = hx_cast_ex_98
					hx_cast_ex_98 = child
					if !Hx_Field_std_isOfType(hx_cast_ex_98, Hx_Obj_loreline_nbeatdecl_RTTI) {
						panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
					}

					(hx_cast_ex_98.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()).Hx_Field_style = ((int32)(1))
				} else {
					if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoicestatement_RTTI) {
						var hx_cast_ex_99 *Hx_Obj_loreline_node
						_ = hx_cast_ex_99
						hx_cast_ex_99 = child
						if !Hx_Field_std_isOfType(hx_cast_ex_99, Hx_Obj_loreline_nchoicestatement_RTTI) {
							panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
						}

						(hx_cast_ex_99.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
						}).Hx_As_Hx_Obj_loreline_nchoicestatement()).Hx_Field_style = ((int32)(1))
					} else {
						if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
							var hx_cast_ex_100 *Hx_Obj_loreline_node
							_ = hx_cast_ex_100
							hx_cast_ex_100 = child
							if !Hx_Field_std_isOfType(hx_cast_ex_100, Hx_Obj_loreline_nchoiceoption_RTTI) {
								panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
							}

							(hx_cast_ex_100.VTable.(interface {
								Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
							}).Hx_As_Hx_Obj_loreline_nchoiceoption()).Hx_Field_style = ((int32)(1))
						} else {
							if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nblock_RTTI) {
								var hx_cast_ex_101 *Hx_Obj_loreline_node
								_ = hx_cast_ex_101
								hx_cast_ex_101 = child
								if !Hx_Field_std_isOfType(hx_cast_ex_101, Hx_Obj_loreline_nblock_RTTI) {
									panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
								}

								(hx_cast_ex_101.VTable.(interface {
									Hx_As_Hx_Obj_loreline_nblock() *Hx_Obj_loreline_nblock
								}).Hx_As_Hx_Obj_loreline_nblock()).Hx_Field_style = ((int32)(1))
							} else {
								if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nliteral_RTTI) {
									var lit *Hx_Obj_loreline_nliteral
									_ = lit
									lit = child.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
									}).Hx_As_Hx_Obj_loreline_nliteral()
									if lit.Hx_Field_literalType.Hx_Field_enumIndex() == 4 {
										lit.Hx_Field_literalType = Hx_Enum_loreline_literaltype_Object{((int32)(1))}
									}
								}
							}
						}
					}
				}
			}
		}
	})
}

func Hx_Field_loreline_astutils_toPlainStyle(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nstatedecl_RTTI) {
			var hx_cast_ex_102 *Hx_Obj_loreline_node
			_ = hx_cast_ex_102
			hx_cast_ex_102 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_102, Hx_Obj_loreline_nstatedecl_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			(hx_cast_ex_102.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
			}).Hx_As_Hx_Obj_loreline_nstatedecl()).Hx_Field_style = ((int32)(0))
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ncharacterdecl_RTTI) {
				var hx_cast_ex_103 *Hx_Obj_loreline_node
				_ = hx_cast_ex_103
				hx_cast_ex_103 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_103, Hx_Obj_loreline_ncharacterdecl_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				(hx_cast_ex_103.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
				}).Hx_As_Hx_Obj_loreline_ncharacterdecl()).Hx_Field_style = ((int32)(0))
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nbeatdecl_RTTI) {
					var hx_cast_ex_104 *Hx_Obj_loreline_node
					_ = hx_cast_ex_104
					hx_cast_ex_104 = child
					if !Hx_Field_std_isOfType(hx_cast_ex_104, Hx_Obj_loreline_nbeatdecl_RTTI) {
						panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
					}

					(hx_cast_ex_104.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()).Hx_Field_style = ((int32)(0))
				} else {
					if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoicestatement_RTTI) {
						var hx_cast_ex_105 *Hx_Obj_loreline_node
						_ = hx_cast_ex_105
						hx_cast_ex_105 = child
						if !Hx_Field_std_isOfType(hx_cast_ex_105, Hx_Obj_loreline_nchoicestatement_RTTI) {
							panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
						}

						(hx_cast_ex_105.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
						}).Hx_As_Hx_Obj_loreline_nchoicestatement()).Hx_Field_style = ((int32)(0))
					} else {
						if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
							var hx_cast_ex_106 *Hx_Obj_loreline_node
							_ = hx_cast_ex_106
							hx_cast_ex_106 = child
							if !Hx_Field_std_isOfType(hx_cast_ex_106, Hx_Obj_loreline_nchoiceoption_RTTI) {
								panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
							}

							(hx_cast_ex_106.VTable.(interface {
								Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
							}).Hx_As_Hx_Obj_loreline_nchoiceoption()).Hx_Field_style = ((int32)(0))
						} else {
							if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nblock_RTTI) {
								var hx_cast_ex_107 *Hx_Obj_loreline_node
								_ = hx_cast_ex_107
								hx_cast_ex_107 = child
								if !Hx_Field_std_isOfType(hx_cast_ex_107, Hx_Obj_loreline_nblock_RTTI) {
									panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
								}

								(hx_cast_ex_107.VTable.(interface {
									Hx_As_Hx_Obj_loreline_nblock() *Hx_Obj_loreline_nblock
								}).Hx_As_Hx_Obj_loreline_nblock()).Hx_Field_style = ((int32)(0))
							} else {
								if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nliteral_RTTI) {
									var lit *Hx_Obj_loreline_nliteral
									_ = lit
									lit = child.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
									}).Hx_As_Hx_Obj_loreline_nliteral()
									if lit.Hx_Field_literalType.Hx_Field_enumIndex() == 4 {
										lit.Hx_Field_literalType = Hx_Enum_loreline_literaltype_Object{((int32)(0))}
									}
								}
							}
						}
					}
				}
			}
		}
	})
}

func Hx_Field_loreline_astutils_useWordOperators(node *Hx_Obj_loreline_astnode) {
	Hx_Field_loreline_astutils_setWordOperators(node, true)
}

func Hx_Field_loreline_astutils_useSymbolOperators(node *Hx_Obj_loreline_astnode) {
	Hx_Field_loreline_astutils_setWordOperators(node, false)
}

func Hx_Field_loreline_astutils_setWordOperators(node *Hx_Obj_loreline_astnode, word bool) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nbinary_RTTI) {
			var bin *Hx_Obj_loreline_nbinary
			_ = bin
			bin = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
			}).Hx_As_Hx_Obj_loreline_nbinary()
			switch bin.Hx_Field_op.Hx_Field_enumIndex() {
			case 32:
				{
					bin.Hx_Field_op = Hx_Enum_loreline_tokentype_OpEquals{word}
				}

			case 33:
				{
					bin.Hx_Field_op = Hx_Enum_loreline_tokentype_OpNotEquals{word}
				}

			case 38:
				{
					bin.Hx_Field_op = Hx_Enum_loreline_tokentype_OpAnd{word}
				}

			case 39:
				{
					bin.Hx_Field_op = Hx_Enum_loreline_tokentype_OpOr{word}
				}

			default:
				{
				}
			}
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nunary_RTTI) {
				var unary *Hx_Obj_loreline_nunary
				_ = unary
				unary = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
				}).Hx_As_Hx_Obj_loreline_nunary()
				if unary.Hx_Field_op.Hx_Field_enumIndex() == 40 {
					unary.Hx_Field_op = Hx_Enum_loreline_tokentype_OpNot{word}
				}
			}
		}
	})
}

func Hx_Field_loreline_astutils_whenRuleScore(rule *Hx_Obj_loreline_nwhenrule) int32 {
	if rule.Hx_Field_condition == nil {
		return 0
	} else {
		return Hx_Field_loreline_astutils_andClauseCount(rule.Hx_Field_condition)
	}
}

func Hx_Field_loreline_astutils_andClauseCount(expr *Hx_Obj_loreline_nexpr) int32 {
	if Hx_Field_loreline_astutils_isLiteralTrue(expr) {
		return 0
	}

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
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_loreline_astutils_andClauseCount(binary.Hx_Field_left)
			return (_hx_tmp_1 + Hx_Field_loreline_astutils_andClauseCount(binary.Hx_Field_right))
		}
	}

	return 1
}

func Hx_Field_loreline_astutils_isLiteralTrue(expr *Hx_Obj_loreline_nexpr) bool {
	if Hx_Field_std_isOfType(expr, Hx_Obj_loreline_nliteral_RTTI) {
		var literal *Hx_Obj_loreline_nliteral
		_ = literal
		literal = expr.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
		}).Hx_As_Hx_Obj_loreline_nliteral()
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (literal.Hx_Field_literalType == Hx_Enum_loreline_literaltype_Boolean{}) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(literal.Hx_Field_value, ((any)(true))))))
		} else {
			_hx_tmp_0 = false
		}

		return _hx_tmp_0
	}

	return false
}

func Hx_Field_loreline_astutils_useParenConditions(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nifstatement_RTTI) {
			var hx_cast_ex_108 *Hx_Obj_loreline_node
			_ = hx_cast_ex_108
			hx_cast_ex_108 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_108, Hx_Obj_loreline_nifstatement_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			(hx_cast_ex_108.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
			}).Hx_As_Hx_Obj_loreline_nifstatement()).Hx_Field_conditionStyle = ((int32)(2))
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
				var hx_cast_ex_109 *Hx_Obj_loreline_node
				_ = hx_cast_ex_109
				hx_cast_ex_109 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_109, Hx_Obj_loreline_nchoiceoption_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				(hx_cast_ex_109.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
				}).Hx_As_Hx_Obj_loreline_nchoiceoption()).Hx_Field_conditionStyle = ((int32)(2))
			}
		}
	})
}

func Hx_Field_loreline_astutils_usePlainConditions(node *Hx_Obj_loreline_astnode) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nifstatement_RTTI) {
			var hx_cast_ex_110 *Hx_Obj_loreline_node
			_ = hx_cast_ex_110
			hx_cast_ex_110 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_110, Hx_Obj_loreline_nifstatement_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			(hx_cast_ex_110.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
			}).Hx_As_Hx_Obj_loreline_nifstatement()).Hx_Field_conditionStyle = ((int32)(0))
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
				var hx_cast_ex_111 *Hx_Obj_loreline_node
				_ = hx_cast_ex_111
				hx_cast_ex_111 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_111, Hx_Obj_loreline_nchoiceoption_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				(hx_cast_ex_111.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
				}).Hx_As_Hx_Obj_loreline_nchoiceoption()).Hx_Field_conditionStyle = ((int32)(0))
			}
		}
	})
}

func Hx_Field_loreline_astutils_reindentFunctions(node *Hx_Obj_loreline_astnode, oldIndent string, newIndent string) {
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nfunctiondecl_RTTI) {
			var _hx_reserved_func *Hx_Obj_loreline_nfunctiondecl
			_ = _hx_reserved_func
			_hx_reserved_func = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
			}).Hx_As_Hx_Obj_loreline_nfunctiondecl()
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if _hx_reserved_func.Hx_Field_code != HxStringNull {
				_hx_tmp_1 = (int32(utf8.RuneCountInString(_hx_reserved_func.Hx_Field_code)) > 0)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				_hx_tmp_0 = !_hx_reserved_func.Hx_Field_external
			} else {
				_hx_tmp_0 = false
			}

			if _hx_tmp_0 {
				_hx_reserved_func.Hx_Field_code = Hx_Field_loreline_astutils_reindentCode(_hx_reserved_func.Hx_Field_code, oldIndent, newIndent)
			}
		}
	})
}

func Hx_Field_loreline_astutils_insertLocalizationKeys(content string, node *Hx_Obj_loreline_astnode, includeImports__isOptional any, reservedIds__isOptional any, rng__isOptional any) string {
	includeImports := HxDefault[bool]()
	_ = includeImports

	if includeImports__isOptional != nil {
		includeImports = HxConvert[bool](includeImports__isOptional)
	} else {
		includeImports = true
	}
	reservedIds := HxDefault[Hx_Typedef_map]()
	_ = reservedIds

	if reservedIds__isOptional != nil {
		reservedIds = HxConvert[Hx_Typedef_map](reservedIds__isOptional)
	} else {
		reservedIds = nil
	}
	rng := HxDefault[*Hx_Obj_loreline_random]()
	_ = rng

	if rng__isOptional != nil {
		rng = HxConvert[*Hx_Obj_loreline_random](rng__isOptional)
	} else {
		rng = nil
	}
	{
		if rng == nil {
			rng = Hx_Obj_loreline_random_CreateInstance(nil)
		}

		var sourceLines HxArray[string]
		_ = sourceLines
		sourceLines = (Hx_Field_go_haxe_hxstring_split(content, "\n"))
		var _hx_tmp_0 func(_hx_carg0 func(_hx_carg0 *Hx_Obj_loreline_node, _hx_carg1 *Hx_Obj_loreline_node))
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if !includeImports {
			_hx_tmp_1 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_script_RTTI)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var hx_cast_ex_112 *Hx_Obj_loreline_astnode
			_ = hx_cast_ex_112
			hx_cast_ex_112 = node
			if !Hx_Field_std_isOfType(hx_cast_ex_112, Hx_Obj_loreline_script_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			_hx_tmp_0 = (hx_cast_ex_112.VTable.(interface {
				Hx_As_Hx_Obj_loreline_script() *Hx_Obj_loreline_script
			}).Hx_As_Hx_Obj_loreline_script()).VTable.Hx_Field_eachExcludingImported
		} else {
			_hx_tmp_0 = node.VTable.Hx_Field_each
		}

		var eachFn func(_hx_carg0 func(_hx_carg0 *Hx_Obj_loreline_node, _hx_carg1 *Hx_Obj_loreline_node))
		_ = eachFn
		eachFn = _hx_tmp_0
		var _hx_tmp_2 Hx_Typedef_map
		_ = _hx_tmp_2
		if reservedIds != nil {
			_hx_tmp_2 = reservedIds
		} else {
			_hx_tmp_2 = (func() *Hx_Obj_haxe_imap {
				_hx_icast_116 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_116 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_116, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_116.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}

		var existingIds Hx_Typedef_map
		_ = existingIds
		existingIds = _hx_tmp_2
		eachFn(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_astnode_RTTI) {
				var astChild *Hx_Obj_loreline_astnode
				_ = astChild
				astChild = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
				if astChild.Hx_Field_trailingComments != nil {
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_comment]
					_ = _g1
					_g1 = astChild.Hx_Field_trailingComments
					for {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						if !(_hx_tmp_3 < _g1.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = _g1.Get(_g)
						_g++
						if c.Hx_Field_isHash {
							func() func(_hx_param_key string, _hx_param_value bool) {
								var _hx_closure_src_451 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_451
								_hx_closure_src_451 = existingIds.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value bool) {
									_hx_closure_src_451(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(c.Hx_Field_content, true)
						}
					}
				}

				if astChild.Hx_Field_leadingComments != nil {
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_comment]
					_ = _g1
					_g1 = astChild.Hx_Field_leadingComments
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g
						if !(_hx_tmp_4 < _g1.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = _g1.Get(_g)
						_g++
						if c.Hx_Field_isHash {
							func() func(_hx_param_key string, _hx_param_value bool) {
								var _hx_closure_src_452 func(_hx_param_key any, _hx_param_value any)
								_ = _hx_closure_src_452
								_hx_closure_src_452 = existingIds.VTable.Hx_Field_set
								return func(_hx_param_key string, _hx_param_value bool) {
									_hx_closure_src_452(((any)(_hx_param_key)), ((any)(_hx_param_value)))
								}
							}()(c.Hx_Field_content, true)
						}
					}
				}
			}
		})
		var insertions HxArray[any]
		_ = insertions
		insertions = HxMakeArray[any]().Dyn()
		eachFn(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
			var str *Hx_Obj_loreline_nstringliteral
			_ = str
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ntextstatement_RTTI) {
				var hx_cast_ex_113 *Hx_Obj_loreline_node
				_ = hx_cast_ex_113
				hx_cast_ex_113 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_113, Hx_Obj_loreline_ntextstatement_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				str = (hx_cast_ex_113.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
				}).Hx_As_Hx_Obj_loreline_ntextstatement()).Hx_Field_content
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ndialoguestatement_RTTI) {
					var hx_cast_ex_114 *Hx_Obj_loreline_node
					_ = hx_cast_ex_114
					hx_cast_ex_114 = child
					if !Hx_Field_std_isOfType(hx_cast_ex_114, Hx_Obj_loreline_ndialoguestatement_RTTI) {
						panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
					}

					str = (hx_cast_ex_114.VTable.(interface {
						Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
					}).Hx_As_Hx_Obj_loreline_ndialoguestatement()).Hx_Field_content
				} else {
					if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
						var hx_cast_ex_115 *Hx_Obj_loreline_node
						_ = hx_cast_ex_115
						hx_cast_ex_115 = child
						if !Hx_Field_std_isOfType(hx_cast_ex_115, Hx_Obj_loreline_nchoiceoption_RTTI) {
							panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
						}

						str = (hx_cast_ex_115.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
						}).Hx_As_Hx_Obj_loreline_nchoiceoption()).Hx_Field_text
					}
				}
			}

			if str != nil {
				var endLine int32
				_ = endLine
				endLine = Hx_Field_loreline_astutils_getStringEndLine(str)
				var lineIdx int32
				_ = lineIdx
				lineIdx = (endLine - 1)
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if lineIdx >= 0 {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = lineIdx
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _hx_tmp_7
					_hx_tmp_6 = (_hx_tmp_8 < sourceLines.Len())
				} else {
					_hx_tmp_6 = false
				}

				if _hx_tmp_6 {
					_hx_tmp_5 = Hx_Field_loreline_astutils_lineHasHashComment(sourceLines.Get(lineIdx))
				} else {
					_hx_tmp_5 = false
				}

				if _hx_tmp_5 {
					return
				}

				var id string
				_ = id
				var iterations int32
				_ = iterations
				iterations = 0
				for {
					var _hx_tmp_9 *Hx_Obj_loreline_random
					_ = _hx_tmp_9
					_hx_tmp_9 = rng
					id = Hx_Field_loreline_astutils_randomId(_hx_tmp_9, (7 + Hx_Field_std_int((((float64)(iterations)) * float64(0.01)))))
					iterations++
					if !(func() func(_hx_param_key string) bool {
						var _hx_closure_src_453 func(_hx_param_key any) bool
						_ = _hx_closure_src_453
						_hx_closure_src_453 = existingIds.VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_453(((any)(_hx_param_key)))
						}
					}()(id)) {
						break
					}
				}

				func() func(_hx_param_key string, _hx_param_value bool) {
					var _hx_closure_src_454 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_454
					_hx_closure_src_454 = existingIds.VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value bool) {
						_hx_closure_src_454(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(id, true)
				var _hx_tmp_10 HxArray[any]
				_ = _hx_tmp_10
				_hx_tmp_10 = HxMakeArrayView[any](insertions)
				var _hx_tmp_11 any
				_ = _hx_tmp_11
				_hx_tmp_11 = ((any)((str.Hx_Field_pos.Hx_Field_offset + str.Hx_Field_pos.Hx_Field_length)))
				_ = (Hx_Array_Push[any](_hx_tmp_10, any(map[string]any{"offset": _hx_tmp_11, "text": ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(" #", id)))})))
			}
		})
		if Hx_Field_go_haxe_hxdynamic_getArrayLength(insertions) == 0 {
			return content
		}

		var _hx_tmp_12 HxArray[any]
		_ = _hx_tmp_12
		_hx_tmp_12 = HxMakeArrayView[any](insertions)
		(Hx_Array_Sort[any](_hx_tmp_12, func(a any, b any) int32 {
			var _hx_tmp_13 any
			_ = _hx_tmp_13
			_hx_tmp_13 = Hx_Field_go_haxe_hxdynamic_getField(b, "offset")
			return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_subtract(_hx_tmp_13, Hx_Field_go_haxe_hxdynamic_getField(a, "offset")))))
		}))
		var result string
		_ = result
		result = content
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_14 int32
				_ = _hx_tmp_14
				_hx_tmp_14 = _g
				if !(_hx_tmp_14 < Hx_Field_go_haxe_hxdynamic_getArrayLength(insertions)) {
					break
				}

				var ins any
				_ = ins
				ins = insertions.Get(_g)
				_g++
				var hx_nullable_t54 any
				_ = hx_nullable_t54
				hx_nullable_t54 = Hx_Field_go_haxe_hxdynamic_getField(ins, "offset")
				var _hx_tmp_15 HxNullable[int32]
				_ = _hx_tmp_15
				if hx_nullable_t54 == nil {
					_hx_tmp_15 = HxNullable[int32]{}
				} else {
					_hx_tmp_15 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t54))), Valid: true}
				}

				var _hx_reserved_len HxNullable[int32]
				_ = _hx_reserved_len
				_hx_reserved_len = _hx_tmp_15
				var _hx_tmp_16 string
				_ = _hx_tmp_16
				if (_hx_reserved_len).Valid != false {
					_hx_tmp_16 = (Hx_Field_go_haxe_hxstring_substr(result, 0, _hx_reserved_len))
				} else {
					var _hx_tmp_17 string
					_ = _hx_tmp_17
					_hx_tmp_17 = result
					var _hx_tmp_18 string
					_ = _hx_tmp_18
					_hx_tmp_18 = _hx_tmp_17
					_hx_tmp_16 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_18, 0, HxNullable[int32]{}))
				}

				var result1 string
				_ = result1
				result1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(((any)((_hx_tmp_16))), Hx_Field_go_haxe_hxdynamic_getField(ins, "text")))
				var pos int32
				_ = pos
				pos = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(ins, "offset"))))
				var len1 HxNullable[int32]
				_ = len1
				len1 = HxNullable[int32]{}

				var _hx_tmp_19 string
				_ = _hx_tmp_19
				_hx_tmp_19 = result1
				var _hx_tmp_20 string
				_ = _hx_tmp_20
				if (len1).Valid != false {
					_hx_tmp_20 = (Hx_Field_go_haxe_hxstring_substr(result, pos, len1))
				} else {
					var _hx_tmp_21 string
					_ = _hx_tmp_21
					_hx_tmp_21 = result
					var _hx_tmp_22 int32
					_ = _hx_tmp_22
					_hx_tmp_22 = pos
					var _hx_tmp_23 string
					_ = _hx_tmp_23
					_hx_tmp_23 = _hx_tmp_21
					var _hx_tmp_24 int32
					_ = _hx_tmp_24
					_hx_tmp_24 = _hx_tmp_22
					_hx_tmp_20 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_23, _hx_tmp_24, HxNullable[int32]{}))
				}

				result = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_19, (_hx_tmp_20))
			}
		}

		return result
	}
}

func Hx_Field_loreline_astutils_removeLocalizationKeys(content string, node *Hx_Obj_loreline_astnode) string {
	var removals HxArray[any]
	_ = removals
	removals = HxMakeArray[any]().Dyn()
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_astnode_RTTI) {
			var astChild *Hx_Obj_loreline_astnode
			_ = astChild
			astChild = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
			{
				var comments HxArray[*Hx_Obj_loreline_comment]
				_ = comments
				comments = astChild.Hx_Field_trailingComments
				if comments != nil {
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g
						if !(_hx_tmp_0 < comments.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = comments.Get(_g)
						_g++
						if c.Hx_Field_isHash {
							var start int32
							_ = start
							start = c.Hx_Field_pos.Hx_Field_offset
							var end int32
							_ = end
							end = (c.Hx_Field_pos.Hx_Field_offset + c.Hx_Field_pos.Hx_Field_length)
							for start > 0 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(content, (start - 1)))
								var _hx_tmp_1 int32
								_ = _hx_tmp_1
								if (tmp_c_1).Valid == false {
									_hx_tmp_1 = -1
								} else {
									_hx_tmp_1 = (tmp_c_1).Value
								}

								var ch int32
								_ = ch
								ch = _hx_tmp_1
								if (ch == 32) || (ch == 9) {
									start--
								} else {
									break
								}
							}

							var _hx_tmp_2 HxArray[any]
							_ = _hx_tmp_2
							_hx_tmp_2 = HxMakeArrayView[any](removals)
							_ = (Hx_Array_Push[any](_hx_tmp_2, any(map[string]any{"start": ((any)(start)), "end": ((any)(end))})))
						}
					}
				}
			}

			{
				var comments1 HxArray[*Hx_Obj_loreline_comment]
				_ = comments1
				comments1 = astChild.Hx_Field_leadingComments
				if comments1 != nil {
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						if !(_hx_tmp_3 < comments1.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = comments1.Get(_g)
						_g++
						if c.Hx_Field_isHash {
							var start int32
							_ = start
							start = c.Hx_Field_pos.Hx_Field_offset
							var end int32
							_ = end
							end = (c.Hx_Field_pos.Hx_Field_offset + c.Hx_Field_pos.Hx_Field_length)
							for start > 0 {
								var tmp_c_1 HxNullable[int32]
								_ = tmp_c_1
								tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(content, (start - 1)))
								var _hx_tmp_4 int32
								_ = _hx_tmp_4
								if (tmp_c_1).Valid == false {
									_hx_tmp_4 = -1
								} else {
									_hx_tmp_4 = (tmp_c_1).Value
								}

								var ch int32
								_ = ch
								ch = _hx_tmp_4
								if (ch == 32) || (ch == 9) {
									start--
								} else {
									break
								}
							}

							var _hx_tmp_5 HxArray[any]
							_ = _hx_tmp_5
							_hx_tmp_5 = HxMakeArrayView[any](removals)
							_ = (Hx_Array_Push[any](_hx_tmp_5, any(map[string]any{"start": ((any)(start)), "end": ((any)(end))})))
						}
					}
				}
			}
		}
	})
	if Hx_Field_go_haxe_hxdynamic_getArrayLength(removals) == 0 {
		return content
	}

	var _hx_tmp_6 HxArray[any]
	_ = _hx_tmp_6
	_hx_tmp_6 = HxMakeArrayView[any](removals)
	(Hx_Array_Sort[any](_hx_tmp_6, func(a any, b any) int32 {
		var _hx_tmp_7 any
		_ = _hx_tmp_7
		_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(b, "start")
		return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_subtract(_hx_tmp_7, Hx_Field_go_haxe_hxdynamic_getField(a, "start")))))
	}))
	var result string
	_ = result
	result = content
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = _g
			if !(_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_getArrayLength(removals)) {
				break
			}

			var r any
			_ = r
			r = removals.Get(_g)
			_g++
			var hx_nullable_t55 any
			_ = hx_nullable_t55
			hx_nullable_t55 = Hx_Field_go_haxe_hxdynamic_getField(r, "start")
			var _hx_tmp_9 HxNullable[int32]
			_ = _hx_tmp_9
			if hx_nullable_t55 == nil {
				_hx_tmp_9 = HxNullable[int32]{}
			} else {
				_hx_tmp_9 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(hx_nullable_t55))), Valid: true}
			}

			var _hx_reserved_len HxNullable[int32]
			_ = _hx_reserved_len
			_hx_reserved_len = _hx_tmp_9
			var _hx_tmp_10 string
			_ = _hx_tmp_10
			if (_hx_reserved_len).Valid != false {
				_hx_tmp_10 = (Hx_Field_go_haxe_hxstring_substr(result, 0, _hx_reserved_len))
			} else {
				var _hx_tmp_11 string
				_ = _hx_tmp_11
				_hx_tmp_11 = result
				var _hx_tmp_12 string
				_ = _hx_tmp_12
				_hx_tmp_12 = _hx_tmp_11
				_hx_tmp_10 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_12, 0, HxNullable[int32]{}))
			}

			var result1 string
			_ = result1
			result1 = _hx_tmp_10
			var pos int32
			_ = pos
			pos = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(r, "end"))))
			var len1 HxNullable[int32]
			_ = len1
			len1 = HxNullable[int32]{}

			var _hx_tmp_13 string
			_ = _hx_tmp_13
			_hx_tmp_13 = result1
			var _hx_tmp_14 string
			_ = _hx_tmp_14
			if (len1).Valid != false {
				_hx_tmp_14 = (Hx_Field_go_haxe_hxstring_substr(result, pos, len1))
			} else {
				var _hx_tmp_15 string
				_ = _hx_tmp_15
				_hx_tmp_15 = result
				var _hx_tmp_16 int32
				_ = _hx_tmp_16
				_hx_tmp_16 = pos
				var _hx_tmp_17 string
				_ = _hx_tmp_17
				_hx_tmp_17 = _hx_tmp_15
				var _hx_tmp_18 int32
				_ = _hx_tmp_18
				_hx_tmp_18 = _hx_tmp_16
				_hx_tmp_14 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_17, _hx_tmp_18, HxNullable[int32]{}))
			}

			result = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_13, (_hx_tmp_14))
		}
	}

	return result
}

func Hx_Field_loreline_astutils_lineHasHashComment(line string) bool {
	var i int32
	_ = i
	i = 0
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < int32(utf8.RuneCountInString(line))) {
			break
		}

		var hx_cmp_14 HxNullable[int32]
		_ = hx_cmp_14
		hx_cmp_14 = (Hx_Field_go_haxe_hxstring_charCodeAt(line, i))
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (hx_cmp_14).Valid == true {
			_hx_tmp_1 = ((hx_cmp_14).Value == 35)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = (i + 1)
			if _hx_tmp_3 < int32(utf8.RuneCountInString(line)) {
				var hx_cmp_15 HxNullable[int32]
				_ = hx_cmp_15
				hx_cmp_15 = (Hx_Field_go_haxe_hxstring_charCodeAt(line, (i + 1)))
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if (hx_cmp_15).Valid == true {
					_hx_tmp_4 = ((hx_cmp_15).Value == 35)
				} else {
					_hx_tmp_4 = false
				}

				_hx_tmp_2 = (_hx_tmp_4)
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				i = (i + 2)
				continue
			}

			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if i > 0 {
				var hx_cmp_16 HxNullable[int32]
				_ = hx_cmp_16
				hx_cmp_16 = (Hx_Field_go_haxe_hxstring_charCodeAt(line, (i - 1)))
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (hx_cmp_16).Valid == true {
					_hx_tmp_6 = ((hx_cmp_16).Value == 92)
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_5 = (_hx_tmp_6)
			} else {
				_hx_tmp_5 = false
			}

			if _hx_tmp_5 {
				i++
				continue
			}

			var j int32
			_ = j
			j = (i + 1)
			for {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = j
				if !(_hx_tmp_7 < int32(utf8.RuneCountInString(line))) {
					break
				}

				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, j))
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				var _hx_tmp_9 bool
				_ = _hx_tmp_9
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if (c).Valid && true && ((c).Value >= 97) || (!(c).Valid && !true) {
					_hx_tmp_12 = ((c).Valid && true && ((c).Value <= 122) || (!(c).Valid && !true))
				} else {
					_hx_tmp_12 = false
				}

				if _hx_tmp_12 {
					_hx_tmp_11 = true
				} else {
					var _hx_tmp_13 bool
					_ = _hx_tmp_13
					if (c).Valid && true && ((c).Value >= 65) || (!(c).Valid && !true) {
						_hx_tmp_13 = ((c).Valid && true && ((c).Value <= 90) || (!(c).Valid && !true))
					} else {
						_hx_tmp_13 = false
					}

					_hx_tmp_11 = _hx_tmp_13
				}

				if _hx_tmp_11 {
					_hx_tmp_10 = true
				} else {
					var _hx_tmp_14 bool
					_ = _hx_tmp_14
					if (c).Valid && true && ((c).Value >= 48) || (!(c).Valid && !true) {
						_hx_tmp_14 = ((c).Valid && true && ((c).Value <= 57) || (!(c).Valid && !true))
					} else {
						_hx_tmp_14 = false
					}

					_hx_tmp_10 = _hx_tmp_14
				}

				if _hx_tmp_10 {
					_hx_tmp_9 = true
				} else {
					var _hx_tmp_15 bool
					_ = _hx_tmp_15
					if (c).Valid == true {
						_hx_tmp_15 = ((c).Value == 95)
					} else {
						_hx_tmp_15 = false
					}

					_hx_tmp_9 = (_hx_tmp_15)
				}

				if _hx_tmp_9 {
					_hx_tmp_8 = true
				} else {
					var _hx_tmp_16 bool
					_ = _hx_tmp_16
					if (c).Valid == true {
						_hx_tmp_16 = ((c).Value == 45)
					} else {
						_hx_tmp_16 = false
					}

					_hx_tmp_8 = (_hx_tmp_16)
				}

				if _hx_tmp_8 {
					j++
				} else {
					break
				}
			}

			if j > (i + 1) {
				return true
			}
		}

		i++
	}

	return false
}

func Hx_Field_loreline_astutils_extractTranslations(node *Hx_Obj_loreline_astnode) Hx_Typedef_map {
	var result *Hx_Obj_haxe_imap
	_ = result
	result = (func() *Hx_Obj_haxe_imap {
		_hx_icast_119 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_119 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_119, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_119.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		var str *Hx_Obj_loreline_nstringliteral
		_ = str
		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ntextstatement_RTTI) {
			astNode = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
			var hx_cast_ex_117 *Hx_Obj_loreline_node
			_ = hx_cast_ex_117
			hx_cast_ex_117 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_117, Hx_Obj_loreline_ntextstatement_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			str = (hx_cast_ex_117.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
			}).Hx_As_Hx_Obj_loreline_ntextstatement()).Hx_Field_content
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ndialoguestatement_RTTI) {
				astNode = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
				var hx_cast_ex_118 *Hx_Obj_loreline_node
				_ = hx_cast_ex_118
				hx_cast_ex_118 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_118, Hx_Obj_loreline_ndialoguestatement_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				str = (hx_cast_ex_118.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
				}).Hx_As_Hx_Obj_loreline_ndialoguestatement()).Hx_Field_content
			}
		}

		if (str != nil) && (astNode != nil) {
			var hashId HxNullable[string]
			_ = hashId
			hashId = Hx_Field_loreline_astutils_findHashComment(astNode, str)
			if (hashId).Valid != false {
				func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nstringliteral) {
					var _hx_closure_src_455 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_455
					_hx_closure_src_455 = result.VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nstringliteral) {
						_hx_closure_src_455(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()((hashId).Value, str)
			}
		}
	})
	return result
}

func Hx_Field_loreline_astutils_extractTranslatableEntries(node *Hx_Obj_loreline_astnode) HxArray[any] {
	var result HxArray[any]
	_ = result
	result = HxMakeArray[any]().Dyn()
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		var str *Hx_Obj_loreline_nstringliteral
		_ = str
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ntextstatement_RTTI) {
			astNode = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
			var hx_cast_ex_120 *Hx_Obj_loreline_node
			_ = hx_cast_ex_120
			hx_cast_ex_120 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_120, Hx_Obj_loreline_ntextstatement_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			str = (hx_cast_ex_120.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
			}).Hx_As_Hx_Obj_loreline_ntextstatement()).Hx_Field_content
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ndialoguestatement_RTTI) {
				astNode = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
				var hx_cast_ex_121 *Hx_Obj_loreline_node
				_ = hx_cast_ex_121
				hx_cast_ex_121 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_121, Hx_Obj_loreline_ndialoguestatement_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				str = (hx_cast_ex_121.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
				}).Hx_As_Hx_Obj_loreline_ndialoguestatement()).Hx_Field_content
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
					var opt *Hx_Obj_loreline_nchoiceoption
					_ = opt
					opt = child.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
					}).Hx_As_Hx_Obj_loreline_nchoiceoption()
					if opt.Hx_Field_text != nil {
						astNode = child.VTable.(interface {
							Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
						}).Hx_As_Hx_Obj_loreline_astnode()
						str = opt.Hx_Field_text
					}
				}
			}
		}

		if (str != nil) && (astNode != nil) {
			var hashId HxNullable[string]
			_ = hashId
			hashId = Hx_Field_loreline_astutils_findHashComment(astNode, str)
			if (hashId).Valid != false {
				var _hx_tmp_0 HxArray[any]
				_ = _hx_tmp_0
				_hx_tmp_0 = HxMakeArrayView[any](result)
				var hx_nullable_f201 HxNullable[string]
				_ = hx_nullable_f201
				hx_nullable_f201 = hashId
				var _hx_tmp_1 any
				_ = _hx_tmp_1
				if (hx_nullable_f201).Valid == true {
					_hx_tmp_1 = (hx_nullable_f201).Value
				} else {
					_hx_tmp_1 = nil
				}

				_ = (Hx_Array_Push[any](_hx_tmp_0, any(map[string]any{"id": _hx_tmp_1, "str": ((any)(str))})))
			}
		}
	})
	return result
}

func Hx_Field_loreline_astutils_hasUntaggedTranslatableStrings(node *Hx_Obj_loreline_astnode) bool {
	var found bool
	_ = found
	found = false
	node.VTable.Hx_Field_each(func(child *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if found {
			return
		}

		var astNode *Hx_Obj_loreline_astnode
		_ = astNode
		var str *Hx_Obj_loreline_nstringliteral
		_ = str
		if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ntextstatement_RTTI) {
			astNode = child.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
			var hx_cast_ex_122 *Hx_Obj_loreline_node
			_ = hx_cast_ex_122
			hx_cast_ex_122 = child
			if !Hx_Field_std_isOfType(hx_cast_ex_122, Hx_Obj_loreline_ntextstatement_RTTI) {
				panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
			}

			str = (hx_cast_ex_122.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
			}).Hx_As_Hx_Obj_loreline_ntextstatement()).Hx_Field_content
		} else {
			if Hx_Field_std_isOfType(child, Hx_Obj_loreline_ndialoguestatement_RTTI) {
				astNode = child.VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
				var hx_cast_ex_123 *Hx_Obj_loreline_node
				_ = hx_cast_ex_123
				hx_cast_ex_123 = child
				if !Hx_Field_std_isOfType(hx_cast_ex_123, Hx_Obj_loreline_ndialoguestatement_RTTI) {
					panic(Hx_Field_haxe_exception_thrown(((any)("Class cast error"))))
				}

				str = (hx_cast_ex_123.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
				}).Hx_As_Hx_Obj_loreline_ndialoguestatement()).Hx_Field_content
			} else {
				if Hx_Field_std_isOfType(child, Hx_Obj_loreline_nchoiceoption_RTTI) {
					var opt *Hx_Obj_loreline_nchoiceoption
					_ = opt
					opt = child.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
					}).Hx_As_Hx_Obj_loreline_nchoiceoption()
					if opt.Hx_Field_text != nil {
						astNode = child.VTable.(interface {
							Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
						}).Hx_As_Hx_Obj_loreline_astnode()
						str = opt.Hx_Field_text
					}
				}
			}
		}

		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (str != nil) && (astNode != nil) {
			_hx_tmp_0 = ((Hx_Field_loreline_astutils_findHashComment(astNode, str)).Valid == false)
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			found = true
		}
	})
	return found
}

func Hx_Field_loreline_astutils_collectHashIds(content string, out__isOptional any) Hx_Typedef_map {
	out := HxDefault[Hx_Typedef_map]()
	_ = out

	if out__isOptional != nil {
		out = HxConvert[Hx_Typedef_map](out__isOptional)
	} else {
		out = nil
	}
	{
		if out == nil {
			out = (func() *Hx_Obj_haxe_imap {
				_hx_icast_124 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_124 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_124, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_124.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}

		if content == HxStringNull {
			return out
		}

		var lexer *Hx_Obj_loreline_lexer
		_ = lexer
		lexer = Hx_Obj_loreline_lexer_CreateInstance(content)
		var tokens Hx_Typedef_loreline_tokens
		_ = tokens
		tokens = lexer.VTable.Hx_Field_tokenize()
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < tokens.Len()) {
					break
				}

				var tok *Hx_Obj_loreline_token
				_ = tok
				tok = tokens.Get(_g)
				_g++
				{
					var _g1 Hx_Enum_loreline_tokentype
					_ = _g1
					_g1 = tok.Hx_Field_type
					if _g1.Hx_Field_enumIndex() == 54 {
						var c string
						_ = c
						c = _g1.(Hx_Enum_loreline_tokentype_CommentHash).Hx_Field_content
						func() func(_hx_param_key string, _hx_param_value bool) {
							var _hx_closure_src_456 func(_hx_param_key any, _hx_param_value any)
							_ = _hx_closure_src_456
							_hx_closure_src_456 = out.VTable.Hx_Field_set
							return func(_hx_param_key string, _hx_param_value bool) {
								_hx_closure_src_456(((any)(_hx_param_key)), ((any)(_hx_param_value)))
							}
						}()(c, true)
					}
				}
			}
		}

		return out
	}
}

func Hx_Field_loreline_astutils_generateTranslationFile(sourceScript *Hx_Obj_loreline_astnode, existingTranslations Hx_Typedef_map, printer *Hx_Obj_loreline_printer) string {
	var entries HxArray[any]
	_ = entries
	entries = Hx_Field_loreline_astutils_extractTranslatableEntries(sourceScript)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var first bool
	_ = first
	first = true
	func() func(x string) {
		var _hx_closure_src_457 func(x any)
		_ = _hx_closure_src_457
		_hx_closure_src_457 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_457(((any)(x)))
		}
	}()("\n")
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(entries)) {
				break
			}

			var entry any
			_ = entry
			entry = entries.Get(_g)
			_g++
			if !first {
				func() func(x string) {
					var _hx_closure_src_458 func(x any)
					_ = _hx_closure_src_458
					_hx_closure_src_458 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_458(((any)(x)))
					}
				}()("\n")
			}

			first = false
			var hx_dyncast_125 any
			_ = hx_dyncast_125
			hx_dyncast_125 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(entry, "str"), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_1 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_1
			if hx_dyncast_125 != nil {
				_hx_tmp_1 = hx_dyncast_125.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
			} else {
				_hx_tmp_1 = nil
			}

			var refText string
			_ = refText
			refText = printer.VTable.Hx_Field_printStringLiteralAsReference(_hx_tmp_1)
			var hx_dyncast_126 any
			_ = hx_dyncast_126
			hx_dyncast_126 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(entry, "str"), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_2 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_2
			if hx_dyncast_126 != nil {
				_hx_tmp_2 = hx_dyncast_126.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
			} else {
				_hx_tmp_2 = nil
			}

			var plainText string
			_ = plainText
			plainText = printer.VTable.Hx_Field_printStringLiteralAsText(_hx_tmp_2)
			func() func(x string) {
				var _hx_closure_src_459 func(x any)
				_ = _hx_closure_src_459
				_hx_closure_src_459 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_459(((any)(x)))
				}
			}()("#")
			func() func(x string) {
				var _hx_closure_src_460 func(x any)
				_ = _hx_closure_src_460
				_hx_closure_src_460 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_460(((any)(x)))
				}
			}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))
			func() func(x string) {
				var _hx_closure_src_461 func(x any)
				_ = _hx_closure_src_461
				_hx_closure_src_461 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_461(((any)(x)))
				}
			}()(" // ")
			func() func(x string) {
				var _hx_closure_src_462 func(x any)
				_ = _hx_closure_src_462
				_hx_closure_src_462 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_462(((any)(x)))
				}
			}()(refText)
			func() func(x string) {
				var _hx_closure_src_463 func(x any)
				_ = _hx_closure_src_463
				_hx_closure_src_463 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_463(((any)(x)))
				}
			}()("\n")
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if existingTranslations != nil {
				_hx_tmp_3 = func() func(_hx_param_key string) bool {
					var _hx_closure_src_464 func(_hx_param_key any) bool
					_ = _hx_closure_src_464
					_hx_closure_src_464 = existingTranslations.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_464(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				func() func(x string) {
					var _hx_closure_src_465 func(x any)
					_ = _hx_closure_src_465
					_hx_closure_src_465 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_465(((any)(x)))
					}
				}()(printer.VTable.Hx_Field_printStringLiteralAsText(func() func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
					var _hx_closure_src_466 func(_hx_param_key any) any
					_ = _hx_closure_src_466
					_hx_closure_src_466 = existingTranslations.VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
						var hx_dyncast_127 any
						_ = hx_dyncast_127
						hx_dyncast_127 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_466(((any)(_hx_param_key))), "Hx_Obj_loreline_nstringliteral")
						var _hx_tmp_4 *Hx_Obj_loreline_nstringliteral
						_ = _hx_tmp_4
						if hx_dyncast_127 != nil {
							_hx_tmp_4 = hx_dyncast_127.(interface {
								Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
							}).Hx_As_Hx_Obj_loreline_nstringliteral()
						} else {
							_hx_tmp_4 = nil
						}

						return _hx_tmp_4
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))))
			} else {
				func() func(x string) {
					var _hx_closure_src_467 func(x any)
					_ = _hx_closure_src_467
					_hx_closure_src_467 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_467(((any)(x)))
					}
				}()(plainText)
			}

			func() func(x string) {
				var _hx_closure_src_468 func(x any)
				_ = _hx_closure_src_468
				_hx_closure_src_468 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_468(((any)(x)))
				}
			}()("\n")
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_astutils_unquotedRawToQuotedRaw(text string) string {
	var result *Hx_Obj_stringbuf
	_ = result
	result = Hx_Obj_stringbuf_CreateInstance()
	var i int32
	_ = i
	i = 0
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < int32(utf8.RuneCountInString(text))) {
			break
		}

		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if (c).Valid == true {
			_hx_tmp_2 = ((c).Value == 92)
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = (i + 1)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_1 = (_hx_tmp_4 < int32(utf8.RuneCountInString(text)))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var next HxNullable[int32]
			_ = next
			next = (Hx_Field_go_haxe_hxstring_charCodeAt(text, (i + 1)))
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			var _hx_tmp_7 bool
			_ = _hx_tmp_7
			var _hx_tmp_8 bool
			_ = _hx_tmp_8
			var _hx_tmp_9 bool
			_ = _hx_tmp_9
			if (next).Valid == true {
				_hx_tmp_9 = ((next).Value == 110)
			} else {
				_hx_tmp_9 = false
			}

			if _hx_tmp_9 {
				_hx_tmp_8 = true
			} else {
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				if (next).Valid == true {
					_hx_tmp_10 = ((next).Value == 116)
				} else {
					_hx_tmp_10 = false
				}

				_hx_tmp_8 = (_hx_tmp_10)
			}

			if _hx_tmp_8 {
				_hx_tmp_7 = true
			} else {
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				if (next).Valid == true {
					_hx_tmp_11 = ((next).Value == 114)
				} else {
					_hx_tmp_11 = false
				}

				_hx_tmp_7 = (_hx_tmp_11)
			}

			if _hx_tmp_7 {
				_hx_tmp_6 = true
			} else {
				var _hx_tmp_12 bool
				_ = _hx_tmp_12
				if (next).Valid == true {
					_hx_tmp_12 = ((next).Value == 92)
				} else {
					_hx_tmp_12 = false
				}

				_hx_tmp_6 = (_hx_tmp_12)
			}

			if _hx_tmp_6 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_13 bool
				_ = _hx_tmp_13
				if (next).Valid == true {
					_hx_tmp_13 = ((next).Value == 60)
				} else {
					_hx_tmp_13 = false
				}

				_hx_tmp_5 = (_hx_tmp_13)
			}

			if _hx_tmp_5 {
				result.VTable.Hx_Field_addChar((c).Value)
				result.VTable.Hx_Field_addChar((next).Value)
				i = (i + 2)
			} else {
				result.VTable.Hx_Field_addChar((next).Value)
				i = (i + 2)
			}
		} else {
			var _hx_tmp_14 bool
			_ = _hx_tmp_14
			if (c).Valid == true {
				_hx_tmp_14 = ((c).Value == 34)
			} else {
				_hx_tmp_14 = false
			}

			if _hx_tmp_14 {
				func() func(x string) {
					var _hx_closure_src_469 func(x any)
					_ = _hx_closure_src_469
					_hx_closure_src_469 = result.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_469(((any)(x)))
					}
				}()("\\\"")
				i++
			} else {
				result.VTable.Hx_Field_addChar((c).Value)
				i++
			}
		}
	}

	return result.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_astutils_quotedRawToUnquotedRaw(text string) string {
	var result *Hx_Obj_stringbuf
	_ = result
	result = Hx_Obj_stringbuf_CreateInstance()
	var i int32
	_ = i
	i = 0
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 < int32(utf8.RuneCountInString(text))) {
			break
		}

		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if (c).Valid == true {
			_hx_tmp_2 = ((c).Value == 92)
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = (i + 1)
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_1 = (_hx_tmp_4 < int32(utf8.RuneCountInString(text)))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			var next HxNullable[int32]
			_ = next
			next = (Hx_Field_go_haxe_hxstring_charCodeAt(text, (i + 1)))
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if (next).Valid == true {
				_hx_tmp_5 = ((next).Value == 34)
			} else {
				_hx_tmp_5 = false
			}

			if _hx_tmp_5 {
				result.VTable.Hx_Field_addChar((next).Value)
				i = (i + 2)
			} else {
				result.VTable.Hx_Field_addChar((c).Value)
				result.VTable.Hx_Field_addChar((next).Value)
				i = (i + 2)
			}
		} else {
			result.VTable.Hx_Field_addChar((c).Value)
			i++
		}
	}

	return result.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_astutils_canSafelyRemoveQuotes(text string) bool {
	if int32(utf8.RuneCountInString(text)) == 0 {
		return false
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "\n", HxNullable[int32]{})) != -1 {
		return false
	}

	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = text
	if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, "{", HxNullable[int32]{})) != -1 {
		return false
	}

	var firstChar HxNullable[int32]
	_ = firstChar
	firstChar = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 0))
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	var _hx_tmp_7 bool
	_ = _hx_tmp_7
	var _hx_tmp_8 bool
	_ = _hx_tmp_8
	if (firstChar).Valid == true {
		_hx_tmp_8 = ((firstChar).Value == 40)
	} else {
		_hx_tmp_8 = false
	}

	if _hx_tmp_8 {
		_hx_tmp_7 = true
	} else {
		var _hx_tmp_9 bool
		_ = _hx_tmp_9
		if (firstChar).Valid == true {
			_hx_tmp_9 = ((firstChar).Value == 91)
		} else {
			_hx_tmp_9 = false
		}

		_hx_tmp_7 = (_hx_tmp_9)
	}

	if _hx_tmp_7 {
		_hx_tmp_6 = true
	} else {
		var _hx_tmp_10 bool
		_ = _hx_tmp_10
		if (firstChar).Valid == true {
			_hx_tmp_10 = ((firstChar).Value == 123)
		} else {
			_hx_tmp_10 = false
		}

		_hx_tmp_6 = (_hx_tmp_10)
	}

	if _hx_tmp_6 {
		_hx_tmp_5 = true
	} else {
		var _hx_tmp_11 bool
		_ = _hx_tmp_11
		if (firstChar).Valid == true {
			_hx_tmp_11 = ((firstChar).Value == 125)
		} else {
			_hx_tmp_11 = false
		}

		_hx_tmp_5 = (_hx_tmp_11)
	}

	if _hx_tmp_5 {
		_hx_tmp_4 = true
	} else {
		var _hx_tmp_12 bool
		_ = _hx_tmp_12
		if (firstChar).Valid == true {
			_hx_tmp_12 = ((firstChar).Value == 93)
		} else {
			_hx_tmp_12 = false
		}

		_hx_tmp_4 = (_hx_tmp_12)
	}

	if _hx_tmp_4 {
		_hx_tmp_3 = true
	} else {
		var _hx_tmp_13 bool
		_ = _hx_tmp_13
		if (firstChar).Valid == true {
			_hx_tmp_13 = ((firstChar).Value == 58)
		} else {
			_hx_tmp_13 = false
		}

		_hx_tmp_3 = (_hx_tmp_13)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		var _hx_tmp_14 bool
		_ = _hx_tmp_14
		if (firstChar).Valid == true {
			_hx_tmp_14 = ((firstChar).Value == 61)
		} else {
			_hx_tmp_14 = false
		}

		_hx_tmp_2 = (_hx_tmp_14)
	}

	if _hx_tmp_2 {
		return false
	}

	var _hx_tmp_15 bool
	_ = _hx_tmp_15
	if int32(utf8.RuneCountInString(text)) >= 2 {
		var _hx_tmp_16 bool
		_ = _hx_tmp_16
		if (firstChar).Valid == true {
			_hx_tmp_16 = ((firstChar).Value == 47)
		} else {
			_hx_tmp_16 = false
		}

		_hx_tmp_15 = (_hx_tmp_16)
	} else {
		_hx_tmp_15 = false
	}

	if _hx_tmp_15 {
		var second HxNullable[int32]
		_ = second
		second = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 1))
		var _hx_tmp_17 bool
		_ = _hx_tmp_17
		var _hx_tmp_18 bool
		_ = _hx_tmp_18
		if (second).Valid == true {
			_hx_tmp_18 = ((second).Value == 47)
		} else {
			_hx_tmp_18 = false
		}

		if _hx_tmp_18 {
			_hx_tmp_17 = true
		} else {
			var _hx_tmp_19 bool
			_ = _hx_tmp_19
			if (second).Valid == true {
				_hx_tmp_19 = ((second).Value == 42)
			} else {
				_hx_tmp_19 = false
			}

			_hx_tmp_17 = (_hx_tmp_19)
		}

		if _hx_tmp_17 {
			return false
		}
	}

	var _hx_tmp_20 bool
	_ = _hx_tmp_20
	var _hx_tmp_21 bool
	_ = _hx_tmp_21
	if int32(utf8.RuneCountInString(text)) >= 2 {
		var _hx_tmp_22 bool
		_ = _hx_tmp_22
		if (firstChar).Valid == true {
			_hx_tmp_22 = ((firstChar).Value == 45)
		} else {
			_hx_tmp_22 = false
		}

		_hx_tmp_21 = (_hx_tmp_22)
	} else {
		_hx_tmp_21 = false
	}

	if _hx_tmp_21 {
		var hx_cmp_17 HxNullable[int32]
		_ = hx_cmp_17
		hx_cmp_17 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 1))
		var _hx_tmp_23 bool
		_ = _hx_tmp_23
		if (hx_cmp_17).Valid == true {
			_hx_tmp_23 = ((hx_cmp_17).Value == 62)
		} else {
			_hx_tmp_23 = false
		}

		_hx_tmp_20 = (_hx_tmp_23)
	} else {
		_hx_tmp_20 = false
	}

	if _hx_tmp_20 {
		return false
	}

	var _hx_tmp_24 bool
	_ = _hx_tmp_24
	var _hx_tmp_25 bool
	_ = _hx_tmp_25
	var _hx_tmp_26 bool
	_ = _hx_tmp_26
	if (firstChar).Valid == true {
		_hx_tmp_26 = ((firstChar).Value == 43)
	} else {
		_hx_tmp_26 = false
	}

	if _hx_tmp_26 {
		_hx_tmp_25 = (int32(utf8.RuneCountInString(text)) >= 3)
	} else {
		_hx_tmp_25 = false
	}

	if _hx_tmp_25 {
		var hx_cmp_18 HxNullable[int32]
		_ = hx_cmp_18
		hx_cmp_18 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 1))
		var _hx_tmp_27 bool
		_ = _hx_tmp_27
		if (hx_cmp_18).Valid == true {
			_hx_tmp_27 = ((hx_cmp_18).Value == 32)
		} else {
			_hx_tmp_27 = false
		}

		_hx_tmp_24 = (_hx_tmp_27)
	} else {
		_hx_tmp_24 = false
	}

	if _hx_tmp_24 {
		return false
	}

	var keywords HxArray[string]
	_ = keywords
	keywords = HxMakeArray[string]("beat ", "state ", "character ", "choice ", "import ", "new ", "function ", "if ", "else ")
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_28 int32
			_ = _hx_tmp_28
			_hx_tmp_28 = _g
			if !(_hx_tmp_28 < keywords.Len()) {
				break
			}

			var kw string
			_ = kw
			kw = keywords.Get(_g)
			_g++
			if Hx_Field_stringtools_startsWith(text, kw) {
				return false
			}
		}
	}

	var _hx_tmp_29 string
	_ = _hx_tmp_29
	_hx_tmp_29 = text
	var colonIdx int32
	_ = colonIdx
	colonIdx = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_29, ":", HxNullable[int32]{}))
	if colonIdx > 0 {
		var allIdentChars bool
		_ = allIdentChars
		allIdentChars = true
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = colonIdx
			for _g < _g1 {
				var _hx_tmp_30 int32
				_ = _hx_tmp_30
				_hx_tmp_30 = _g
				_g = (_g + 1)
				var j int32
				_ = j
				j = _hx_tmp_30
				var ch HxNullable[int32]
				_ = ch
				ch = (Hx_Field_go_haxe_hxstring_charCodeAt(text, j))
				var _hx_tmp_31 bool
				_ = _hx_tmp_31
				var _hx_tmp_32 bool
				_ = _hx_tmp_32
				var _hx_tmp_33 bool
				_ = _hx_tmp_33
				var _hx_tmp_34 bool
				_ = _hx_tmp_34
				if (ch).Valid && true && ((ch).Value >= 97) || (!(ch).Valid && !true) {
					_hx_tmp_34 = ((ch).Valid && true && ((ch).Value <= 122) || (!(ch).Valid && !true))
				} else {
					_hx_tmp_34 = false
				}

				if _hx_tmp_34 {
					_hx_tmp_33 = true
				} else {
					var _hx_tmp_35 bool
					_ = _hx_tmp_35
					if (ch).Valid && true && ((ch).Value >= 65) || (!(ch).Valid && !true) {
						_hx_tmp_35 = ((ch).Valid && true && ((ch).Value <= 90) || (!(ch).Valid && !true))
					} else {
						_hx_tmp_35 = false
					}

					_hx_tmp_33 = _hx_tmp_35
				}

				if _hx_tmp_33 {
					_hx_tmp_32 = true
				} else {
					var _hx_tmp_36 bool
					_ = _hx_tmp_36
					if (ch).Valid && true && ((ch).Value >= 48) || (!(ch).Valid && !true) {
						_hx_tmp_36 = ((ch).Valid && true && ((ch).Value <= 57) || (!(ch).Valid && !true))
					} else {
						_hx_tmp_36 = false
					}

					_hx_tmp_32 = _hx_tmp_36
				}

				if _hx_tmp_32 {
					_hx_tmp_31 = true
				} else {
					var _hx_tmp_37 bool
					_ = _hx_tmp_37
					if (ch).Valid == true {
						_hx_tmp_37 = ((ch).Value == 95)
					} else {
						_hx_tmp_37 = false
					}

					_hx_tmp_31 = (_hx_tmp_37)
				}

				if !(_hx_tmp_31) {
					allIdentChars = false
					break
				}
			}
		}

		if allIdentChars {
			return false
		}
	}

	if ((text == "null") || (text == "true")) || (text == "false") {
		return false
	}

	if Hx_Field_loreline_astutils_isNumberLiteral(text) {
		return false
	}

	if int32(utf8.RuneCountInString(text)) >= 2 {
		var second HxNullable[int32]
		_ = second
		second = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 1))
		var _hx_tmp_38 bool
		_ = _hx_tmp_38
		var _hx_tmp_39 bool
		_ = _hx_tmp_39
		if (second).Valid == true {
			_hx_tmp_39 = ((second).Value == 61)
		} else {
			_hx_tmp_39 = false
		}

		if _hx_tmp_39 {
			var _hx_tmp_40 bool
			_ = _hx_tmp_40
			var _hx_tmp_41 bool
			_ = _hx_tmp_41
			var _hx_tmp_42 bool
			_ = _hx_tmp_42
			var _hx_tmp_43 bool
			_ = _hx_tmp_43
			if (firstChar).Valid == true {
				_hx_tmp_43 = ((firstChar).Value == 43)
			} else {
				_hx_tmp_43 = false
			}

			if _hx_tmp_43 {
				_hx_tmp_42 = true
			} else {
				var _hx_tmp_44 bool
				_ = _hx_tmp_44
				if (firstChar).Valid == true {
					_hx_tmp_44 = ((firstChar).Value == 45)
				} else {
					_hx_tmp_44 = false
				}

				_hx_tmp_42 = (_hx_tmp_44)
			}

			if _hx_tmp_42 {
				_hx_tmp_41 = true
			} else {
				var _hx_tmp_45 bool
				_ = _hx_tmp_45
				if (firstChar).Valid == true {
					_hx_tmp_45 = ((firstChar).Value == 42)
				} else {
					_hx_tmp_45 = false
				}

				_hx_tmp_41 = (_hx_tmp_45)
			}

			if _hx_tmp_41 {
				_hx_tmp_40 = true
			} else {
				var _hx_tmp_46 bool
				_ = _hx_tmp_46
				if (firstChar).Valid == true {
					_hx_tmp_46 = ((firstChar).Value == 47)
				} else {
					_hx_tmp_46 = false
				}

				_hx_tmp_40 = (_hx_tmp_46)
			}

			_hx_tmp_38 = (_hx_tmp_40)
		} else {
			_hx_tmp_38 = false
		}

		if _hx_tmp_38 {
			return false
		}
	}

	return true
}

func Hx_Field_loreline_astutils_isNumberLiteral(text string) bool {
	if int32(utf8.RuneCountInString(text)) == 0 {
		return false
	}

	var i int32
	_ = i
	i = 0
	var hx_cmp_19 HxNullable[int32]
	_ = hx_cmp_19
	hx_cmp_19 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, 0))
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (hx_cmp_19).Valid == true {
		_hx_tmp_0 = ((hx_cmp_19).Value == 45)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		i++
	}

	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = i
	if _hx_tmp_1 >= int32(utf8.RuneCountInString(text)) {
		return false
	}

	var hasDigit bool
	_ = hasDigit
	hasDigit = false
	var hasDot bool
	_ = hasDot
	hasDot = false
	for {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = i
		if !(_hx_tmp_2 < int32(utf8.RuneCountInString(text))) {
			break
		}

		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if (c).Valid && true && ((c).Value >= 48) || (!(c).Valid && !true) {
			_hx_tmp_3 = ((c).Valid && true && ((c).Value <= 57) || (!(c).Valid && !true))
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			hasDigit = true
		} else {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if (c).Valid == true {
				_hx_tmp_5 = ((c).Value == 46)
			} else {
				_hx_tmp_5 = false
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = !hasDot
			} else {
				_hx_tmp_4 = false
			}

			if _hx_tmp_4 {
				hasDot = true
			} else {
				return false
			}
		}

		i++
	}

	return hasDigit
}

func Hx_Field_loreline_astutils_reindentCode(code string, oldIndent string, newIndent string) string {
	if oldIndent == newIndent {
		return code
	}

	var lines HxArray[string]
	_ = lines
	lines = (Hx_Field_go_haxe_hxstring_split(code, "\n"))
	var result *Hx_Obj_stringbuf
	_ = result
	result = Hx_Obj_stringbuf_CreateInstance()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = lines.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			if i > 0 {
				func() func(x string) {
					var _hx_closure_src_470 func(x any)
					_ = _hx_closure_src_470
					_hx_closure_src_470 = result.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_470(((any)(x)))
					}
				}()("\n")
			}

			var line string
			_ = line
			line = lines.Get(i)
			var level int32
			_ = level
			level = 0
			var pos int32
			_ = pos
			pos = 0
			for {
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				_hx_tmp_1 = line
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = pos
				if !(Hx_Field_stringtools_startsWith((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{})), oldIndent)) {
					break
				}

				level++
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = pos
				pos = (_hx_tmp_3 + int32(utf8.RuneCountInString(oldIndent)))
			}

			{
				var tmp__g1_1 int32
				_ = tmp__g1_1
				tmp__g1_1 = 0
				var _g2 int32
				_ = _g2
				_g2 = level
				for tmp__g1_1 < _g2 {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = tmp__g1_1
					tmp__g1_1 = (tmp__g1_1 + 1)
					var _hx_reserved__ int32
					_ = _hx_reserved__
					_hx_reserved__ = _hx_tmp_4
					func() func(x string) {
						var _hx_closure_src_471 func(x any)
						_ = _hx_closure_src_471
						_hx_closure_src_471 = result.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_471(((any)(x)))
						}
					}()(newIndent)
				}
			}

			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = line
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = pos
			func() func(x string) {
				var _hx_closure_src_472 func(x any)
				_ = _hx_closure_src_472
				_hx_closure_src_472 = result.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_472(((any)(x)))
				}
			}()((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, _hx_tmp_6, HxNullable[int32]{})))
		}
	}

	return result.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_astutils_getStringEndLine(str *Hx_Obj_loreline_nstringliteral) int32 {
	var endLine int32
	_ = endLine
	endLine = str.Hx_Field_pos.Hx_Field_line
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
			if part.Hx_Field_pos != nil {
				var partEndLine int32
				_ = partEndLine
				partEndLine = part.Hx_Field_pos.Hx_Field_line
				{
					var tmp__g_1 Hx_Enum_loreline_stringparttype
					_ = tmp__g_1
					tmp__g_1 = part.Hx_Field_partType
					if tmp__g_1.Hx_Field_enumIndex() == 0 {
						var text string
						_ = text
						text = tmp__g_1.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						{
							var i int32
							_ = i
							i = 0
							for {
								var _hx_tmp_1 int32
								_ = _hx_tmp_1
								_hx_tmp_1 = i
								if !(_hx_tmp_1 < int32(utf8.RuneCountInString(text))) {
									break
								}

								var hx_cmp_20 HxNullable[int32]
								_ = hx_cmp_20
								hx_cmp_20 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
								var _hx_tmp_2 bool
								_ = _hx_tmp_2
								if (hx_cmp_20).Valid == true {
									_hx_tmp_2 = ((hx_cmp_20).Value == 10)
								} else {
									_hx_tmp_2 = false
								}

								if _hx_tmp_2 {
									partEndLine++
								}

								i++
							}
						}
					}
				}

				if partEndLine > endLine {
					endLine = partEndLine
				}
			}
		}
	}

	return endLine
}

func Hx_Field_loreline_astutils_findHashComment(node *Hx_Obj_loreline_astnode, str__isOptional any) HxNullable[string] {
	str := HxDefault[*Hx_Obj_loreline_nstringliteral]()
	_ = str

	if str__isOptional != nil {
		str = HxConvert[*Hx_Obj_loreline_nstringliteral](str__isOptional)
	} else {
		str = nil
	}
	{
		if node.Hx_Field_trailingComments != nil {
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

				var c *Hx_Obj_loreline_comment
				_ = c
				c = _g1.Get(_g)
				_g++
				if c.Hx_Field_isHash {
					return HxNullableString(c.Hx_Field_content)
				}
			}
		}

		if node.Hx_Field_leadingComments != nil {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_comment]
			_ = _g1
			_g1 = node.Hx_Field_leadingComments
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < _g1.Len()) {
					break
				}

				var c *Hx_Obj_loreline_comment
				_ = c
				c = _g1.Get(_g)
				_g++
				if c.Hx_Field_isHash {
					return HxNullableString(c.Hx_Field_content)
				}
			}
		}

		if str != nil {
			if str.Hx_Field_trailingComments != nil {
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = str.Hx_Field_trailingComments
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g
					if !(_hx_tmp_2 < _g1.Len()) {
						break
					}

					var c *Hx_Obj_loreline_comment
					_ = c
					c = _g1.Get(_g)
					_g++
					if c.Hx_Field_isHash {
						return HxNullableString(c.Hx_Field_content)
					}
				}
			}

			if str.Hx_Field_leadingComments != nil {
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = str.Hx_Field_leadingComments
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					if !(_hx_tmp_3 < _g1.Len()) {
						break
					}

					var c *Hx_Obj_loreline_comment
					_ = c
					c = _g1.Get(_g)
					_g++
					if c.Hx_Field_isHash {
						return HxNullableString(c.Hx_Field_content)
					}
				}
			}
		}

		return HxNullable[string]{Value: HxStringNull}
	}
}

func Hx_Field_loreline_astutils_randomId(rng *Hx_Obj_loreline_random, length int32) string {
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = length
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var _hx_reserved__ int32
			_ = _hx_reserved__
			_hx_reserved__ = _hx_tmp_0
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_loreline_astutils_randomId_chars
			rng.Hx_Field_seed = math.Mod((rng.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
			var _hx_tmp_2 float64
			_ = _hx_tmp_2
			_hx_tmp_2 = rng.Hx_Field_seed
			buf.VTable.Hx_Field_addChar((Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_1, Hx_Field_math_floor((((float64)(36)) * (((float64)(((_hx_tmp_2) / ((float64)(2147483647))))) + float64(0.000000000233)))))).Value)
		}
	}

	var result string
	_ = result
	result = buf.VTable.Hx_Field_toString()
	var allHex bool
	_ = allHex
	allHex = true
	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = int32(utf8.RuneCountInString(result))
		for _g2 < _g3 {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_3
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(result, i))
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if (c).Valid && true && ((c).Value >= 48) || (!(c).Valid && !true) {
				_hx_tmp_5 = ((c).Valid && true && ((c).Value <= 57) || (!(c).Valid && !true))
			} else {
				_hx_tmp_5 = false
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = true
			} else {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if (c).Valid && true && ((c).Value >= 97) || (!(c).Valid && !true) {
					_hx_tmp_6 = ((c).Valid && true && ((c).Value <= 102) || (!(c).Valid && !true))
				} else {
					_hx_tmp_6 = false
				}

				_hx_tmp_4 = _hx_tmp_6
			}

			if !(_hx_tmp_4) {
				allHex = false
				break
			}
		}
	}

	if allHex {
		var _hx_tmp_7 float64
		_ = _hx_tmp_7
		_hx_tmp_7 = ((float64)(length))
		rng.Hx_Field_seed = math.Mod((rng.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
		var _hx_tmp_8 float64
		_ = _hx_tmp_8
		_hx_tmp_8 = rng.Hx_Field_seed
		var pos int32
		_ = pos
		pos = Hx_Field_math_floor((_hx_tmp_7 * (((float64)(((_hx_tmp_8) / ((float64)(2147483647))))) + float64(0.000000000233))))
		var _hx_tmp_9 string
		_ = _hx_tmp_9
		_hx_tmp_9 = Hx_Field_loreline_astutils_randomId_nonHexChars
		var _hx_tmp_10 float64
		_ = _hx_tmp_10
		_hx_tmp_10 = ((float64)(((float64)(int32(utf8.RuneCountInString(Hx_Field_loreline_astutils_randomId_nonHexChars))))))
		rng.Hx_Field_seed = math.Mod((rng.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
		var _hx_tmp_11 float64
		_ = _hx_tmp_11
		_hx_tmp_11 = rng.Hx_Field_seed
		var replacement HxNullable[int32]
		_ = replacement
		replacement = (Hx_Field_go_haxe_hxstring_charCodeAt(_hx_tmp_9, Hx_Field_math_floor((_hx_tmp_10 * (((float64)(((_hx_tmp_11) / ((float64)(2147483647))))) + float64(0.000000000233))))))
		var _hx_tmp_14 string
		_ = _hx_tmp_14
		_hx_tmp_14 = result
		var _hx_tmp_13 string
		_ = _hx_tmp_13
		_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_14, 0, HxNullable[int32]{Value: pos, Valid: true}))
		var _hx_tmp_12 string
		_ = _hx_tmp_12
		_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_13, Hx_Field_go_haxe_hxstring_fromCharCode((replacement).Value))
		var _hx_tmp_15 string
		_ = _hx_tmp_15
		_hx_tmp_15 = result
		var _hx_tmp_16 int32
		_ = _hx_tmp_16
		_hx_tmp_16 = (pos + 1)
		result = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_12, (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, _hx_tmp_16, HxNullable[int32]{})))
	}

	return result
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
