package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_jsontoast() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_jsontoast() {
}

type Hx_Obj_VTable_loreline_jsontoast interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_jsontoast struct {
	VTable Hx_Obj_VTable_loreline_jsontoast `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_jsontoast_CreateEmptyInstance() *Hx_Obj_loreline_jsontoast {
	obj := &Hx_Obj_loreline_jsontoast{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_jsontoast_CreateInstance() *Hx_Obj_loreline_jsontoast {
	obj := Hx_Obj_loreline_jsontoast_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_jsontoast_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.JsonToAst",
	HxMakeArray[string]("scriptFromJson", "nodeFromJson", "nodeFromJsonOfType", "commentFromJson", "stateFromJson", "objectFieldFromJson", "characterFromJson", "beatFromJson", "stringPartFromJson", "stringLiteralFromJson", "textStatementFromJson", "dialogueStatementFromJson", "choiceStatementFromJson", "choiceOptionFromJson", "whenStatementFromJson", "whenRuleFromJson", "blockFromJson", "ifStatementFromJson", "alternativeFromJson", "callFromJson", "transitionFromJson", "beatCallFromJson", "insertionFromJson", "functionDeclFromJson", "literalFromJson", "accessFromJson", "assignFromJson", "arrayAccessFromJson", "binaryFromJson", "unaryFromJson", "ternaryFromJson", "importStatementFromJson", "idFromJson", "positionFromJson", "extractComments", "applyComments", "blockStyleFromString", "conditionStyleFromString", "alternativeModeFromString", "quotesFromString", "tokenTypeFromString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_jsontoast_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_jsontoast_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_jsontoast_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "scriptFromJson":
			return Hx_Field_loreline_jsontoast_scriptFromJson
		case "nodeFromJson":
			return Hx_Field_loreline_jsontoast_nodeFromJson
		case "nodeFromJsonOfType":
			return Hx_Field_loreline_jsontoast_nodeFromJsonOfType
		case "commentFromJson":
			return Hx_Field_loreline_jsontoast_commentFromJson
		case "stateFromJson":
			return Hx_Field_loreline_jsontoast_stateFromJson
		case "objectFieldFromJson":
			return Hx_Field_loreline_jsontoast_objectFieldFromJson
		case "characterFromJson":
			return Hx_Field_loreline_jsontoast_characterFromJson
		case "beatFromJson":
			return Hx_Field_loreline_jsontoast_beatFromJson
		case "stringPartFromJson":
			return Hx_Field_loreline_jsontoast_stringPartFromJson
		case "stringLiteralFromJson":
			return Hx_Field_loreline_jsontoast_stringLiteralFromJson
		case "textStatementFromJson":
			return Hx_Field_loreline_jsontoast_textStatementFromJson
		case "dialogueStatementFromJson":
			return Hx_Field_loreline_jsontoast_dialogueStatementFromJson
		case "choiceStatementFromJson":
			return Hx_Field_loreline_jsontoast_choiceStatementFromJson
		case "choiceOptionFromJson":
			return Hx_Field_loreline_jsontoast_choiceOptionFromJson
		case "whenStatementFromJson":
			return Hx_Field_loreline_jsontoast_whenStatementFromJson
		case "whenRuleFromJson":
			return Hx_Field_loreline_jsontoast_whenRuleFromJson
		case "blockFromJson":
			return Hx_Field_loreline_jsontoast_blockFromJson
		case "ifStatementFromJson":
			return Hx_Field_loreline_jsontoast_ifStatementFromJson
		case "alternativeFromJson":
			return Hx_Field_loreline_jsontoast_alternativeFromJson
		case "callFromJson":
			return Hx_Field_loreline_jsontoast_callFromJson
		case "transitionFromJson":
			return Hx_Field_loreline_jsontoast_transitionFromJson
		case "beatCallFromJson":
			return Hx_Field_loreline_jsontoast_beatCallFromJson
		case "insertionFromJson":
			return Hx_Field_loreline_jsontoast_insertionFromJson
		case "functionDeclFromJson":
			return Hx_Field_loreline_jsontoast_functionDeclFromJson
		case "literalFromJson":
			return Hx_Field_loreline_jsontoast_literalFromJson
		case "accessFromJson":
			return Hx_Field_loreline_jsontoast_accessFromJson
		case "assignFromJson":
			return Hx_Field_loreline_jsontoast_assignFromJson
		case "arrayAccessFromJson":
			return Hx_Field_loreline_jsontoast_arrayAccessFromJson
		case "binaryFromJson":
			return Hx_Field_loreline_jsontoast_binaryFromJson
		case "unaryFromJson":
			return Hx_Field_loreline_jsontoast_unaryFromJson
		case "ternaryFromJson":
			return Hx_Field_loreline_jsontoast_ternaryFromJson
		case "importStatementFromJson":
			return Hx_Field_loreline_jsontoast_importStatementFromJson
		case "idFromJson":
			return Hx_Field_loreline_jsontoast_idFromJson
		case "positionFromJson":
			return Hx_Field_loreline_jsontoast_positionFromJson
		case "extractComments":
			return Hx_Field_loreline_jsontoast_extractComments
		case "applyComments":
			return Hx_Field_loreline_jsontoast_applyComments
		case "blockStyleFromString":
			return Hx_Field_loreline_jsontoast_blockStyleFromString
		case "conditionStyleFromString":
			return Hx_Field_loreline_jsontoast_conditionStyleFromString
		case "alternativeModeFromString":
			return Hx_Field_loreline_jsontoast_alternativeModeFromString
		case "quotesFromString":
			return Hx_Field_loreline_jsontoast_quotesFromString
		case "tokenTypeFromString":
			return Hx_Field_loreline_jsontoast_tokenTypeFromString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_jsontoast) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_jsontoast_RTTI
}

func (this *Hx_Obj_loreline_jsontoast) Hx_As_Hx_Obj_loreline_jsontoast() *Hx_Obj_loreline_jsontoast {
	return this
}

func Hx_Field_loreline_jsontoast_scriptFromJson(json any) *Hx_Obj_loreline_script {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(json, nil))) {
		panic(Hx_Field_haxe_exception_thrown(((any)("Cannot create Script from null JSON"))))
	}

	var nodeType string
	_ = nodeType
	nodeType = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "type"))
	if nodeType != "Script" {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Expected Script node, got: ", nodeType)))))
	}

	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var bodyArr HxArray[any]
	_ = bodyArr
	bodyArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "body"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(bodyArr)) {
				break
			}

			var item any
			_ = item
			item = bodyArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	body = _g
	var script *Hx_Obj_loreline_script
	_ = script
	script = Hx_Obj_loreline_script_CreateInstance(id, pos, body)
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "indentSize"), nil))) {
		_hx_tmp_2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(json, "indentSize"))))
	} else {
		_hx_tmp_2 = 2
	}

	script.Hx_Field_indentSize = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "filePath"), nil))) {
		script.Hx_Field_filePath = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "filePath"))
	}

	Hx_Field_loreline_jsontoast_applyComments((func() *Hx_Obj_loreline_astnode {
		_hx_up := script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), json)
	return script
}

func Hx_Field_loreline_jsontoast_nodeFromJson(json any) *Hx_Obj_loreline_node {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(json, nil))) {
		return nil
	}

	var node *Hx_Obj_loreline_node
	_ = node
	node = Hx_Field_loreline_jsontoast_nodeFromJsonOfType(json)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "parens"), nil))) {
		_hx_tmp_0 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_nexpr_RTTI)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		node.VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr().Hx_Field_parens = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_getField(json, "parens")))))
	}

	return node
}

func Hx_Field_loreline_jsontoast_nodeFromJsonOfType(json any) *Hx_Obj_loreline_node {
	var nodeType string
	_ = nodeType
	nodeType = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "type"))
	switch nodeType {
	case "Access":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_accessFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Alternative":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_alternativeFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "ArrayAccess":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_arrayAccessFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Assign":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_assignFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Beat":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_beatFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "BeatCall":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_beatCallFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Binary":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_binaryFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Block":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_blockFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Call":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_callFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Character":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_characterFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Choice":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_choiceStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Comment":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_commentFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Dialogue":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_dialogueStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Field":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_objectFieldFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Function":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_functionDeclFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "If":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_ifStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Import":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_importStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Insertion":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_insertionFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Literal":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_literalFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Option":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_choiceOptionFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Part":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_stringPartFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Script":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_scriptFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "State":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_stateFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "String":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_stringLiteralFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Ternary":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_ternaryFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Text":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_textStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Transition":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_transitionFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "Unary":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_unaryFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "When":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_whenStatementFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	case "WhenRule":
		{
			return (func() *Hx_Obj_loreline_node {
				_hx_up := Hx_Field_loreline_jsontoast_whenRuleFromJson(json)
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unknown node type: ", nodeType)))))
		}
	}
}

func Hx_Field_loreline_jsontoast_commentFromJson(json any) *Hx_Obj_loreline_comment {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "isHash"), nil))) {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "isHash"))))
	} else {
		_hx_tmp_0 = false
	}

	var isHash bool
	_ = isHash
	isHash = _hx_tmp_0
	return Hx_Obj_loreline_comment_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "content")), ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "multiline")))), isHash)
}

func Hx_Field_loreline_jsontoast_stateFromJson(json any) *Hx_Obj_loreline_nstatedecl {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var fieldsArr HxArray[any]
	_ = fieldsArr
	fieldsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "fields"))
	var _g HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nobjectfield]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fieldsArr)) {
				break
			}

			var f any
			_ = f
			f = fieldsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nobjectfield]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(f).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()))
		}
	}

	var fields HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = fields
	fields = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nstatedecl
	_ = node
	node = Hx_Obj_loreline_nstatedecl_CreateInstance(id, pos, ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "temporary")))), fields, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	return node
}

func Hx_Field_loreline_jsontoast_objectFieldFromJson(json any) *Hx_Obj_loreline_nobjectfield {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nobjectfield_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "name")), Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "value")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr(), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_characterFromJson(json any) *Hx_Obj_loreline_ncharacterdecl {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var namePos *Hx_Obj_loreline_position
	_ = namePos
	namePos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "namePos"))
	var fieldsArr HxArray[any]
	_ = fieldsArr
	fieldsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "fields"))
	var _g HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nobjectfield]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(fieldsArr)) {
				break
			}

			var f any
			_ = f
			f = fieldsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nobjectfield]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(f).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
			}).Hx_As_Hx_Obj_loreline_nobjectfield()))
		}
	}

	var fields HxArray[*Hx_Obj_loreline_nobjectfield]
	_ = fields
	fields = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_ncharacterdecl
	_ = node
	node = Hx_Obj_loreline_ncharacterdecl_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "name")), namePos, fields, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	return node
}

func Hx_Field_loreline_jsontoast_beatFromJson(json any) *Hx_Obj_loreline_nbeatdecl {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var bodyArr HxArray[any]
	_ = bodyArr
	bodyArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "body"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(bodyArr)) {
				break
			}

			var item any
			_ = item
			item = bodyArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	body = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nbeatdecl
	_ = node
	node = Hx_Obj_loreline_nbeatdecl_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "name")), body, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "params"), nil))) {
		var paramsArr HxArray[any]
		_ = paramsArr
		paramsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "params"))
		var tmp__g_1 HxArray[*Hx_Obj_loreline_nbeatparam]
		_ = tmp__g_1
		tmp__g_1 = HxMakeArray[*Hx_Obj_loreline_nbeatparam]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(paramsArr)) {
					break
				}

				var p any
				_ = p
				p = paramsArr.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_nbeatparam]
				_ = _hx_tmp_3
				_hx_tmp_3 = tmp__g_1
				var _hx_tmp_4 *Hx_Obj_loreline_nexpr
				_ = _hx_tmp_4
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(p, "defaultValue"), nil)))) {
					_hx_tmp_4 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(p, "defaultValue")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
				} else {
					_hx_tmp_4 = nil
				}

				_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatparam](_hx_tmp_3, Hx_Obj_loreline_nbeatparam_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(p, "name")), Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(p, "namePos")), _hx_tmp_4)))
			}
		}

		node.Hx_Field_params = tmp__g_1
	}

	return node
}

func Hx_Field_loreline_jsontoast_stringPartFromJson(json any) *Hx_Obj_loreline_nstringpart {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var partStr string
	_ = partStr
	partStr = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "part"))
	var partType Hx_Enum_loreline_stringparttype
	_ = partType
	switch partStr {
	case "Expr":
		{
			partType = Hx_Enum_loreline_stringparttype_Expr{Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "expression")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()}
		}

	case "Raw":
		{
			partType = Hx_Enum_loreline_stringparttype_Raw{Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "text"))}
		}

	case "Tag":
		{
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "closing"))))
			partType = Hx_Enum_loreline_stringparttype_Tag{_hx_tmp_0, Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "content")).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
			}).Hx_As_Hx_Obj_loreline_nstringliteral()}
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unknown string part type: ", partStr)))))
		}
	}

	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nstringpart_CreateInstance(id, pos, partType, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_stringLiteralFromJson(json any) *Hx_Obj_loreline_nstringliteral {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var quotes int32
	_ = quotes
	quotes = Hx_Field_loreline_jsontoast_quotesFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "quotes")))
	var partsArr HxArray[any]
	_ = partsArr
	partsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "parts"))
	var _g HxArray[*Hx_Obj_loreline_nstringpart]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nstringpart]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(partsArr)) {
				break
			}

			var p any
			_ = p
			p = partsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nstringpart]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(p).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstringpart() *Hx_Obj_loreline_nstringpart
			}).Hx_As_Hx_Obj_loreline_nstringpart()))
		}
	}

	var parts HxArray[*Hx_Obj_loreline_nstringpart]
	_ = parts
	parts = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nstringliteral_CreateInstance(id, pos, quotes, parts, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_textStatementFromJson(json any) *Hx_Obj_loreline_ntextstatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var content *Hx_Obj_loreline_nstringliteral
	_ = content
	content = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "content")).VTable.(interface {
		Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
	}).Hx_As_Hx_Obj_loreline_nstringliteral()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_ntextstatement
	_ = node
	node = Hx_Obj_loreline_ntextstatement_CreateInstance(id, pos, content, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "condition"), nil))) {
		node.Hx_Field_condition = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
		node.Hx_Field_conditionStyle = ((int32)(Hx_Field_loreline_jsontoast_conditionStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionStyle")))))
		node.Hx_Field_conditionPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionPos"))
	}

	return node
}

func Hx_Field_loreline_jsontoast_dialogueStatementFromJson(json any) *Hx_Obj_loreline_ndialoguestatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var characterPos *Hx_Obj_loreline_position
	_ = characterPos
	characterPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "characterPos"))
	var content *Hx_Obj_loreline_nstringliteral
	_ = content
	content = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "content")).VTable.(interface {
		Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
	}).Hx_As_Hx_Obj_loreline_nstringliteral()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_ndialoguestatement
	_ = node
	node = Hx_Obj_loreline_ndialoguestatement_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "character")), characterPos, content, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "condition"), nil))) {
		node.Hx_Field_condition = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
		node.Hx_Field_conditionStyle = ((int32)(Hx_Field_loreline_jsontoast_conditionStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionStyle")))))
		node.Hx_Field_conditionPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionPos"))
	}

	return node
}

func Hx_Field_loreline_jsontoast_choiceStatementFromJson(json any) *Hx_Obj_loreline_nchoicestatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var optionsArr HxArray[any]
	_ = optionsArr
	optionsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "options"))
	var _g HxArray[*Hx_Obj_loreline_nchoiceoption]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nchoiceoption]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(optionsArr)) {
				break
			}

			var o any
			_ = o
			o = optionsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nchoiceoption]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nchoiceoption](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(o).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
			}).Hx_As_Hx_Obj_loreline_nchoiceoption()))
		}
	}

	var options HxArray[*Hx_Obj_loreline_nchoiceoption]
	_ = options
	options = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nchoicestatement
	_ = node
	node = Hx_Obj_loreline_nchoicestatement_CreateInstance(id, pos, options, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	return node
}

func Hx_Field_loreline_jsontoast_choiceOptionFromJson(json any) *Hx_Obj_loreline_nchoiceoption {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var _hx_tmp_0 *Hx_Obj_loreline_nstringliteral
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "text"), nil))) {
		_hx_tmp_0 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "text")).VTable.(interface {
			Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
		}).Hx_As_Hx_Obj_loreline_nstringliteral()
	} else {
		_hx_tmp_0 = nil
	}

	var text *Hx_Obj_loreline_nstringliteral
	_ = text
	text = _hx_tmp_0
	var _hx_tmp_1 *Hx_Obj_loreline_ninsertion
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "insertion"), nil))) {
		_hx_tmp_1 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "insertion")).VTable.(interface {
			Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
		}).Hx_As_Hx_Obj_loreline_ninsertion()
	} else {
		_hx_tmp_1 = nil
	}

	var insertion *Hx_Obj_loreline_ninsertion
	_ = insertion
	insertion = _hx_tmp_1
	var _hx_tmp_2 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "condition"), nil))) {
		_hx_tmp_2 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	} else {
		_hx_tmp_2 = nil
	}

	var condition *Hx_Obj_loreline_nexpr
	_ = condition
	condition = _hx_tmp_2
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionStyle"), nil))) {
		_hx_tmp_3 = Hx_Field_loreline_jsontoast_conditionStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionStyle")))
	} else {
		_hx_tmp_3 = ((int32)(0))
	}

	var condStyle int32
	_ = condStyle
	condStyle = _hx_tmp_3
	var bodyArr HxArray[any]
	_ = bodyArr
	bodyArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "body"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g1
			if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(bodyArr)) {
				break
			}

			var item any
			_ = item
			item = bodyArr.Get(_g1)
			_g1++
			var _hx_tmp_5 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_5, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	body = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nchoiceoption
	_ = node
	node = Hx_Obj_loreline_nchoiceoption_CreateInstance(id, pos, text, insertion, condition, condStyle, body, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "once"), nil))) {
		_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "once"))))
	} else {
		_hx_tmp_6 = false
	}

	if _hx_tmp_6 {
		node.Hx_Field_once = true
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionPos"), nil))) {
		node.Hx_Field_conditionPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionPos"))
	}

	return node
}

func Hx_Field_loreline_jsontoast_whenStatementFromJson(json any) *Hx_Obj_loreline_nwhenstatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var rulesArr HxArray[any]
	_ = rulesArr
	rulesArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "rules"))
	var _g HxArray[*Hx_Obj_loreline_nwhenrule]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nwhenrule]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(rulesArr)) {
				break
			}

			var r any
			_ = r
			r = rulesArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nwhenrule]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nwhenrule](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(r).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
			}).Hx_As_Hx_Obj_loreline_nwhenrule()))
		}
	}

	var rules HxArray[*Hx_Obj_loreline_nwhenrule]
	_ = rules
	rules = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var hx_nullable_t29 any
	_ = hx_nullable_t29
	hx_nullable_t29 = Hx_Field_go_haxe_hxdynamic_getField(json, "strategy")
	var _hx_tmp_2 HxNullable[string]
	_ = _hx_tmp_2
	if hx_nullable_t29 == nil {
		_hx_tmp_2 = HxNullable[string]{Value: HxStringNull}
	} else {
		_hx_tmp_2 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t29), Valid: true}
	}

	var node *Hx_Obj_loreline_nwhenstatement
	_ = node
	node = Hx_Obj_loreline_nwhenstatement_CreateInstance(id, pos, _hx_tmp_2, rules, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "strategyPos"), nil))) {
		node.Hx_Field_strategyPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "strategyPos"))
	}

	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	return node
}

func Hx_Field_loreline_jsontoast_whenRuleFromJson(json any) *Hx_Obj_loreline_nwhenrule {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var _hx_tmp_0 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "condition"), nil))) {
		_hx_tmp_0 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	} else {
		_hx_tmp_0 = nil
	}

	var condition *Hx_Obj_loreline_nexpr
	_ = condition
	condition = _hx_tmp_0
	var bodyArr HxArray[any]
	_ = bodyArr
	bodyArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "body"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(bodyArr)) {
				break
			}

			var item any
			_ = item
			item = bodyArr.Get(_g1)
			_g1++
			var _hx_tmp_2 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_2, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	body = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nwhenrule
	_ = node
	node = Hx_Obj_loreline_nwhenrule_CreateInstance(id, pos, condition, body, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "insertion"), nil))) {
		node.Hx_Field_insertion = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "insertion")).VTable.(interface {
			Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
		}).Hx_As_Hx_Obj_loreline_ninsertion()
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "insertionCondition"), nil))) {
		node.Hx_Field_insertionCondition = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "insertionCondition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "insertionConditionStyle"), nil))) {
			_hx_tmp_3 = Hx_Field_loreline_jsontoast_conditionStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "insertionConditionStyle")))
		} else {
			_hx_tmp_3 = ((int32)(0))
		}

		node.Hx_Field_insertionConditionStyle = _hx_tmp_3
	}

	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "once"), nil))) {
		_hx_tmp_4 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "once"))))
	} else {
		_hx_tmp_4 = false
	}

	if _hx_tmp_4 {
		node.Hx_Field_once = true
	}

	return node
}

func Hx_Field_loreline_jsontoast_blockFromJson(json any) *Hx_Obj_loreline_nblock {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var bodyArr HxArray[any]
	_ = bodyArr
	bodyArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "body"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(bodyArr)) {
				break
			}

			var item any
			_ = item
			item = bodyArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var body HxArray[*Hx_Obj_loreline_astnode]
	_ = body
	body = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nblock
	_ = node
	node = Hx_Obj_loreline_nblock_CreateInstance(id, pos, body, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	return node
}

func Hx_Field_loreline_jsontoast_ifStatementFromJson(json any) *Hx_Obj_loreline_nifstatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var condition *Hx_Obj_loreline_nexpr
	_ = condition
	condition = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var condStyle int32
	_ = condStyle
	condStyle = Hx_Field_loreline_jsontoast_conditionStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "conditionStyle")))
	var thenArr HxArray[any]
	_ = thenArr
	thenArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "thenBranch"))
	var _g HxArray[*Hx_Obj_loreline_astnode]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_astnode]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(thenArr)) {
				break
			}

			var item any
			_ = item
			item = thenArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
		}
	}

	var thenBody HxArray[*Hx_Obj_loreline_astnode]
	_ = thenBody
	thenBody = _g
	var thenBlock *Hx_Obj_loreline_nblock
	_ = thenBlock
	thenBlock = Hx_Obj_loreline_nblock_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, thenBody, nil, nil)
	thenBlock.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "thenStyle")))))
	var elseBlock *Hx_Obj_loreline_nblock
	_ = elseBlock
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "elseBranch"), nil))) {
		var elseArr HxArray[any]
		_ = elseArr
		elseArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "elseBranch"))
		var tmp__g_1 HxArray[*Hx_Obj_loreline_astnode]
		_ = tmp__g_1
		tmp__g_1 = HxMakeArray[*Hx_Obj_loreline_astnode]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(elseArr)) {
					break
				}

				var item any
				_ = item
				item = elseArr.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
				_ = _hx_tmp_3
				_hx_tmp_3 = tmp__g_1
				_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](_hx_tmp_3, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()))
			}
		}

		var elseBody HxArray[*Hx_Obj_loreline_astnode]
		_ = elseBody
		elseBody = tmp__g_1
		elseBlock = Hx_Obj_loreline_nblock_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, elseBody, nil, nil)
		elseBlock.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "elseStyle")))))
	}

	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var elseLeading HxArray[*Hx_Obj_loreline_comment]
	_ = elseLeading
	var elseTrailing HxArray[*Hx_Obj_loreline_comment]
	_ = elseTrailing
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), nil))) {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "elseLeading"), nil))) {
			var arr HxArray[any]
			_ = arr
			arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "elseLeading"))
			var tmp__g_1 HxArray[*Hx_Obj_loreline_comment]
			_ = tmp__g_1
			tmp__g_1 = HxMakeArray[*Hx_Obj_loreline_comment]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g1
					if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
						break
					}

					var c any
					_ = c
					c = arr.Get(_g1)
					_g1++
					var _hx_tmp_5 HxArray[*Hx_Obj_loreline_comment]
					_ = _hx_tmp_5
					_hx_tmp_5 = tmp__g_1
					_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_5, Hx_Field_loreline_jsontoast_commentFromJson(c)))
				}
			}

			elseLeading = tmp__g_1
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "elseTrailing"), nil))) {
			var arr HxArray[any]
			_ = arr
			arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "elseTrailing"))
			var tmp__g_1 HxArray[*Hx_Obj_loreline_comment]
			_ = tmp__g_1
			tmp__g_1 = HxMakeArray[*Hx_Obj_loreline_comment]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _g1
					if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
						break
					}

					var c any
					_ = c
					c = arr.Get(_g1)
					_g1++
					var _hx_tmp_7 HxArray[*Hx_Obj_loreline_comment]
					_ = _hx_tmp_7
					_hx_tmp_7 = tmp__g_1
					_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_7, Hx_Field_loreline_jsontoast_commentFromJson(c)))
				}
			}

			elseTrailing = tmp__g_1
		}
	}

	return Hx_Obj_loreline_nifstatement_CreateInstance(id, pos, condition, condStyle, thenBlock, elseBlock, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")), elseLeading, elseTrailing)
}

func Hx_Field_loreline_jsontoast_alternativeFromJson(json any) *Hx_Obj_loreline_nalternative {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var mode int32
	_ = mode
	mode = Hx_Field_loreline_jsontoast_alternativeModeFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "mode")))
	var itemsArr HxArray[any]
	_ = itemsArr
	itemsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "items"))
	var _g HxArray[*Hx_Obj_loreline_nblock]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nblock]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(itemsArr)) {
				break
			}

			var item any
			_ = item
			item = itemsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nblock]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nblock](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(item).VTable.(interface {
				Hx_As_Hx_Obj_loreline_nblock() *Hx_Obj_loreline_nblock
			}).Hx_As_Hx_Obj_loreline_nblock()))
		}
	}

	var items HxArray[*Hx_Obj_loreline_nblock]
	_ = items
	items = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nalternative
	_ = node
	node = Hx_Obj_loreline_nalternative_CreateInstance(id, pos, mode, items, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	node.Hx_Field_style = ((int32)(Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "separatorComments"), nil))) {
		var sepArr HxArray[any]
		_ = sepArr
		sepArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "separatorComments"))
		var tmp__g_1 HxArray[HxArray[*Hx_Obj_loreline_comment]]
		_ = tmp__g_1
		tmp__g_1 = HxMakeArray[HxArray[*Hx_Obj_loreline_comment]]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(sepArr)) {
					break
				}

				var commentsGroup any
				_ = commentsGroup
				commentsGroup = sepArr.Get(_g1)
				_g1++
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(commentsGroup, nil))) {
					var group HxArray[any]
					_ = group
					group = HxAnyToArray(commentsGroup)
					var tmp__g1_1 HxArray[*Hx_Obj_loreline_comment]
					_ = tmp__g1_1
					tmp__g1_1 = HxMakeArray[*Hx_Obj_loreline_comment]()
					{
						var _g2 int32
						_ = _g2
						_g2 = 0
						for {
							var _hx_tmp_3 int32
							_ = _hx_tmp_3
							_hx_tmp_3 = _g2
							if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(group)) {
								break
							}

							var c any
							_ = c
							c = group.Get(_g2)
							_g2++
							var _hx_tmp_4 HxArray[*Hx_Obj_loreline_comment]
							_ = _hx_tmp_4
							_hx_tmp_4 = tmp__g1_1
							_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_4, Hx_Field_loreline_jsontoast_commentFromJson(c)))
						}
					}

					_ = (Hx_Array_Push[HxArray[*Hx_Obj_loreline_comment]](tmp__g_1, tmp__g1_1))
				} else {
					_ = (Hx_Array_Push[HxArray[*Hx_Obj_loreline_comment]](tmp__g_1, nil))
				}
			}
		}

		node.Hx_Field_separatorComments = tmp__g_1
	}

	return node
}

func Hx_Field_loreline_jsontoast_callFromJson(json any) *Hx_Obj_loreline_ncall {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var target *Hx_Obj_loreline_nexpr
	_ = target
	target = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "target")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var argsArr HxArray[any]
	_ = argsArr
	argsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "args"))
	var _g HxArray[*Hx_Obj_loreline_nexpr]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_loreline_nexpr]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(argsArr)) {
				break
			}

			var a any
			_ = a
			a = argsArr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(a).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()))
		}
	}

	var args HxArray[*Hx_Obj_loreline_nexpr]
	_ = args
	args = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_ncall_CreateInstance(id, pos, target, args, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_transitionFromJson(json any) *Hx_Obj_loreline_ntransition {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var targetPos *Hx_Obj_loreline_position
	_ = targetPos
	targetPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetPos"))
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_ntransition
	_ = node
	node = Hx_Obj_loreline_ntransition_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "target")), targetPos, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "targetExpr"), nil))) {
		node.Hx_Field_targetExpr = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetExpr")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "args"), nil))) {
		var argsArr HxArray[any]
		_ = argsArr
		argsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "args"))
		var _g HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_nexpr]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(argsArr)) {
					break
				}

				var a any
				_ = a
				a = argsArr.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(a).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()))
			}
		}

		node.Hx_Field_args = _g
	}

	return node
}

func Hx_Field_loreline_jsontoast_beatCallFromJson(json any) *Hx_Obj_loreline_nbeatcall {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var targetPos *Hx_Obj_loreline_position
	_ = targetPos
	targetPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetPos"))
	var targetExpr *Hx_Obj_loreline_nexpr
	_ = targetExpr
	targetExpr = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetExpr")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_nbeatcall
	_ = node
	node = Hx_Obj_loreline_nbeatcall_CreateInstance(id, pos, targetExpr, targetPos, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "args"), nil))) {
		var argsArr HxArray[any]
		_ = argsArr
		argsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "args"))
		var _g HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_nexpr]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(argsArr)) {
					break
				}

				var a any
				_ = a
				a = argsArr.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(a).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()))
			}
		}

		node.Hx_Field_args = _g
	}

	return node
}

func Hx_Field_loreline_jsontoast_insertionFromJson(json any) *Hx_Obj_loreline_ninsertion {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var targetPos *Hx_Obj_loreline_position
	_ = targetPos
	targetPos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetPos"))
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var node *Hx_Obj_loreline_ninsertion
	_ = node
	node = Hx_Obj_loreline_ninsertion_CreateInstance(id, pos, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "target")), targetPos, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "targetExpr"), nil))) {
		node.Hx_Field_targetExpr = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "targetExpr")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "args"), nil))) {
		var argsArr HxArray[any]
		_ = argsArr
		argsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "args"))
		var _g HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_nexpr]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(argsArr)) {
					break
				}

				var a any
				_ = a
				a = argsArr.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_nexpr]
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nexpr](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(a).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()))
			}
		}

		node.Hx_Field_args = _g
	}

	return node
}

func Hx_Field_loreline_jsontoast_functionDeclFromJson(json any) *Hx_Obj_loreline_nfunctiondecl {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var hx_nullable_t30 any
	_ = hx_nullable_t30
	hx_nullable_t30 = Hx_Field_go_haxe_hxdynamic_getField(json, "name")
	var _hx_tmp_0 HxNullable[string]
	_ = _hx_tmp_0
	if hx_nullable_t30 == nil {
		_hx_tmp_0 = HxNullable[string]{Value: HxStringNull}
	} else {
		_hx_tmp_0 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t30), Valid: true}
	}

	var name HxNullable[string]
	_ = name
	name = _hx_tmp_0
	var argsArr HxArray[any]
	_ = argsArr
	argsArr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "args"))
	var _g HxArray[string]
	_ = _g
	_g = HxMakeArray[string]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(argsArr)) {
				break
			}

			var a any
			_ = a
			a = argsArr.Get(_g1)
			_g1++
			var _hx_tmp_2 HxArray[string]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[string](_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_toString(a)))
		}
	}

	var args HxArray[string]
	_ = args
	args = _g
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nfunctiondecl_CreateInstance(id, pos, name, args, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "code")), ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(json, "external")))), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_literalFromJson(json any) *Hx_Obj_loreline_nliteral {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var litStr string
	_ = litStr
	litStr = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "literal"))
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	var literalType Hx_Enum_loreline_literaltype
	_ = literalType
	var value any
	_ = value
	switch litStr {
	case "Array":
		{
			literalType = Hx_Enum_loreline_literaltype_Array{}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "value"), nil))) {
				var arr HxArray[any]
				_ = arr
				arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "value"))
				var _g HxArray[any]
				_ = _g
				_g = HxMakeArray[any]().Dyn()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = _g1
						if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
							break
						}

						var elem any
						_ = elem
						elem = arr.Get(_g1)
						_g1++
						if Hx_Field_reflect_hasField(elem, "type") {
							var _hx_tmp_1 HxArray[any]
							_ = _hx_tmp_1
							_hx_tmp_1 = _g
							_ = (Hx_Array_Push[any](_hx_tmp_1, Hx_Field_loreline_jsontoast_nodeFromJson(elem)))
						} else {
							_ = (Hx_Array_Push[any](_g, elem))
						}
					}
				}

				value = ((any)(_g))
			} else {
				value = nil
			}
		}

	case "Boolean":
		{
			literalType = Hx_Enum_loreline_literaltype_Boolean{}

			value = Hx_Field_go_haxe_hxdynamic_getField(json, "value")
		}

	case "Null":
		{
			literalType = Hx_Enum_loreline_literaltype_Null{}

			value = nil
		}

	case "Number":
		{
			literalType = Hx_Enum_loreline_literaltype_Number{}

			value = Hx_Field_go_haxe_hxdynamic_getField(json, "value")
		}

	case "Object":
		{
			var style int32
			_ = style
			style = Hx_Field_loreline_jsontoast_blockStyleFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "style")))
			literalType = Hx_Enum_loreline_literaltype_Object{style}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "value"), nil))) {
				var arr HxArray[any]
				_ = arr
				arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(json, "value"))
				var _g HxArray[*Hx_Obj_loreline_nobjectfield]
				_ = _g
				_g = HxMakeArray[*Hx_Obj_loreline_nobjectfield]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					for {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = _g1
						if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
							break
						}

						var f any
						_ = f
						f = arr.Get(_g1)
						_g1++
						var _hx_tmp_3 HxArray[*Hx_Obj_loreline_nobjectfield]
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						_ = (Hx_Array_Push[*Hx_Obj_loreline_nobjectfield](_hx_tmp_3, Hx_Field_loreline_jsontoast_nodeFromJson(f).VTable.(interface {
							Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield
						}).Hx_As_Hx_Obj_loreline_nobjectfield()))
					}
				}

				value = ((any)(_g))
			} else {
				value = ((any)(HxMakeArray[*Hx_Obj_loreline_nobjectfield]()))
			}
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unknown literal type: ", litStr)))))
		}
	}

	return Hx_Obj_loreline_nliteral_CreateInstance(id, pos, value, literalType, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_accessFromJson(json any) *Hx_Obj_loreline_naccess {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var _hx_tmp_0 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "target"), nil))) {
		_hx_tmp_0 = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "target")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	} else {
		_hx_tmp_0 = nil
	}

	var target *Hx_Obj_loreline_nexpr
	_ = target
	target = _hx_tmp_0
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_naccess_CreateInstance(id, pos, target, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "name")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_assignFromJson(json any) *Hx_Obj_loreline_nassign {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var target *Hx_Obj_loreline_nexpr
	_ = target
	target = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "target")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var op Hx_Enum_loreline_tokentype
	_ = op
	op = Hx_Field_loreline_jsontoast_tokenTypeFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "op")))
	var value *Hx_Obj_loreline_nexpr
	_ = value
	value = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "value")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nassign_CreateInstance(id, pos, target, op, value, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_arrayAccessFromJson(json any) *Hx_Obj_loreline_narrayaccess {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var target *Hx_Obj_loreline_nexpr
	_ = target
	target = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "target")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var index *Hx_Obj_loreline_nexpr
	_ = index
	index = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "index")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_narrayaccess_CreateInstance(id, pos, target, index, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_binaryFromJson(json any) *Hx_Obj_loreline_nbinary {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var left *Hx_Obj_loreline_nexpr
	_ = left
	left = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "left")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var op Hx_Enum_loreline_tokentype
	_ = op
	op = Hx_Field_loreline_jsontoast_tokenTypeFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "op")))
	var right *Hx_Obj_loreline_nexpr
	_ = right
	right = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "right")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nbinary_CreateInstance(id, pos, left, op, right, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_unaryFromJson(json any) *Hx_Obj_loreline_nunary {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var op Hx_Enum_loreline_tokentype
	_ = op
	op = Hx_Field_loreline_jsontoast_tokenTypeFromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "op")))
	var operand *Hx_Obj_loreline_nexpr
	_ = operand
	operand = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "operand")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nunary_CreateInstance(id, pos, op, operand, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_ternaryFromJson(json any) *Hx_Obj_loreline_nternary {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var condition *Hx_Obj_loreline_nexpr
	_ = condition
	condition = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "condition")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var trueExpr *Hx_Obj_loreline_nexpr
	_ = trueExpr
	trueExpr = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "trueExpr")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var falseExpr *Hx_Obj_loreline_nexpr
	_ = falseExpr
	falseExpr = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "falseExpr")).VTable.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nternary_CreateInstance(id, pos, condition, trueExpr, falseExpr, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_importStatementFromJson(json any) *Hx_Obj_loreline_nimportstatement {
	var id Hx_Typedef_haxe__int64___int64
	_ = id
	id = Hx_Field_loreline_jsontoast_idFromJson(json)
	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = Hx_Field_loreline_jsontoast_positionFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "pos"))
	var path *Hx_Obj_loreline_nstringliteral
	_ = path
	path = Hx_Field_loreline_jsontoast_nodeFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "path")).VTable.(interface {
		Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
	}).Hx_As_Hx_Obj_loreline_nstringliteral()
	var _hx_tmp_0 *Hx_Obj_loreline_script
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "script"), nil))) {
		_hx_tmp_0 = Hx_Field_loreline_jsontoast_scriptFromJson(Hx_Field_go_haxe_hxdynamic_getField(json, "script"))
	} else {
		_hx_tmp_0 = nil
	}

	var script *Hx_Obj_loreline_script
	_ = script
	script = _hx_tmp_0
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	return Hx_Obj_loreline_nimportstatement_CreateInstance(id, pos, path, script, HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading")), HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing")))
}

func Hx_Field_loreline_jsontoast_idFromJson(json any) Hx_Typedef_haxe__int64___int64 {
	var idStr string
	_ = idStr
	idStr = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(json, "id"))
	if idStr != HxStringNull {
		return ((Hx_Typedef_haxe__int64___int64)(Hx_Field_loreline__node_nodeid_impl__fromString(idStr)))
	} else {
		return Hx_Field_loreline__node_nodeid_impl__UNDEFINED
	}
}

func Hx_Field_loreline_jsontoast_positionFromJson(json any) *Hx_Obj_loreline_position {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(json, nil))) {
		return Hx_Obj_loreline_position_CreateInstance(0, 0, 0, 0)
	}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(json, "length"), nil))) {
		_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(json, "length"))))
	} else {
		_hx_tmp_0 = 0
	}

	var length int32
	_ = length
	length = _hx_tmp_0
	return Hx_Obj_loreline_position_CreateInstance(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(json, "line")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(json, "column")))), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(json, "offset")))), length)
}

func Hx_Field_loreline_jsontoast_extractComments(json any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), nil))) {
		return any(map[string]any{"leading": ((any)(nil)), "trailing": ((any)(nil))})
	}

	var leading HxArray[*Hx_Obj_loreline_comment]
	_ = leading
	var trailing HxArray[*Hx_Obj_loreline_comment]
	_ = trailing
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "leading"), nil))) {
		var arr HxArray[any]
		_ = arr
		arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "leading"))
		var _g HxArray[*Hx_Obj_loreline_comment]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_comment]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
					break
				}

				var c any
				_ = c
				c = arr.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_comment]
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_1, Hx_Field_loreline_jsontoast_commentFromJson(c)))
			}
		}

		leading = _g
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "trailing"), nil))) {
		var arr HxArray[any]
		_ = arr
		arr = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(json, "comments"), "trailing"))
		var _g HxArray[*Hx_Obj_loreline_comment]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_comment]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
					break
				}

				var c any
				_ = c
				c = arr.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_comment]
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_comment](_hx_tmp_3, Hx_Field_loreline_jsontoast_commentFromJson(c)))
			}
		}

		trailing = _g
	}

	return any(map[string]any{"leading": ((any)(leading)), "trailing": ((any)(trailing))})
}

func Hx_Field_loreline_jsontoast_applyComments(node *Hx_Obj_loreline_astnode, json any) {
	var comments any
	_ = comments
	comments = Hx_Field_loreline_jsontoast_extractComments(json)
	node.Hx_Field_leadingComments = HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "leading"))
	node.Hx_Field_trailingComments = HxAnyToTypedArray[*Hx_Obj_loreline_comment](Hx_Field_go_haxe_hxdynamic_getField(comments, "trailing"))
}

func Hx_Field_loreline_jsontoast_blockStyleFromString(s string) int32 {
	if s == "Braces" {
		return ((int32)(1))
	} else {
		return ((int32)(0))
	}
}

func Hx_Field_loreline_jsontoast_conditionStyleFromString(s string) int32 {
	if s == "Parens" {
		return ((int32)(2))
	} else {
		return ((int32)(0))
	}
}

func Hx_Field_loreline_jsontoast_alternativeModeFromString(s string) int32 {
	switch s {
	case "Cycle":
		{
			return ((int32)(1))
		}

	case "Once":
		{
			return ((int32)(2))
		}

	case "Pick":
		{
			return ((int32)(3))
		}

	case "Shuffle":
		{
			return ((int32)(4))
		}

	default:
		{
			return ((int32)(0))
		}
	}
}

func Hx_Field_loreline_jsontoast_quotesFromString(s string) int32 {
	if s == "DoubleQuotes" {
		return ((int32)(1))
	} else {
		return ((int32)(0))
	}
}

func Hx_Field_loreline_jsontoast_tokenTypeFromString(s string) Hx_Enum_loreline_tokentype {
	switch s {
	case "OpAssign":
		{
			return Hx_Enum_loreline_tokentype_OpAssign{}
		}

	case "OpDivide":
		{
			return Hx_Enum_loreline_tokentype_OpDivide{}
		}

	case "OpDivideAssign":
		{
			return Hx_Enum_loreline_tokentype_OpDivideAssign{}
		}

	case "OpGreater":
		{
			return Hx_Enum_loreline_tokentype_OpGreater{}
		}

	case "OpGreaterEq":
		{
			return Hx_Enum_loreline_tokentype_OpGreaterEq{}
		}

	case "OpLess":
		{
			return Hx_Enum_loreline_tokentype_OpLess{}
		}

	case "OpLessEq":
		{
			return Hx_Enum_loreline_tokentype_OpLessEq{}
		}

	case "OpMinus":
		{
			return Hx_Enum_loreline_tokentype_OpMinus{}
		}

	case "OpMinusAssign":
		{
			return Hx_Enum_loreline_tokentype_OpMinusAssign{}
		}

	case "OpModulo":
		{
			return Hx_Enum_loreline_tokentype_OpModulo{}
		}

	case "OpMultiply":
		{
			return Hx_Enum_loreline_tokentype_OpMultiply{}
		}

	case "OpMultiplyAssign":
		{
			return Hx_Enum_loreline_tokentype_OpMultiplyAssign{}
		}

	case "OpPlus":
		{
			return Hx_Enum_loreline_tokentype_OpPlus{}
		}

	case "OpPlusAssign":
		{
			return Hx_Enum_loreline_tokentype_OpPlusAssign{}
		}

	case "OpUnquotedAssign":
		{
			return Hx_Enum_loreline_tokentype_OpUnquotedAssign{}
		}

	default:
		{
			if Hx_Field_stringtools_startsWith(s, "OpNotEquals") {
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = s
				return Hx_Enum_loreline_tokentype_OpNotEquals{((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "true", HxNullable[int32]{})) != -1)}
			} else {
				if Hx_Field_stringtools_startsWith(s, "OpEquals") {
					var _hx_tmp_1 string
					_ = _hx_tmp_1
					_hx_tmp_1 = s
					return Hx_Enum_loreline_tokentype_OpEquals{((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_1, "true", HxNullable[int32]{})) != -1)}
				} else {
					if Hx_Field_stringtools_startsWith(s, "OpNot") {
						var _hx_tmp_2 string
						_ = _hx_tmp_2
						_hx_tmp_2 = s
						return Hx_Enum_loreline_tokentype_OpNot{((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_2, "true", HxNullable[int32]{})) != -1)}
					} else {
						if Hx_Field_stringtools_startsWith(s, "OpAnd") {
							var _hx_tmp_3 string
							_ = _hx_tmp_3
							_hx_tmp_3 = s
							return Hx_Enum_loreline_tokentype_OpAnd{((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_3, "true", HxNullable[int32]{})) != -1)}
						} else {
							if Hx_Field_stringtools_startsWith(s, "OpOr") {
								var _hx_tmp_4 string
								_ = _hx_tmp_4
								_hx_tmp_4 = s
								return Hx_Enum_loreline_tokentype_OpOr{((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_4, "true", HxNullable[int32]{})) != -1)}
							} else {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Unknown operator: ", s)))))
							}
						}
					}
				}
			}
		}
	}
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatparam]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nblock]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nchoiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[HxArray[*Hx_Obj_loreline_comment]]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
