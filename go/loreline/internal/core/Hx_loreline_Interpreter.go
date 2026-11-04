package core

import "unicode/utf8"
import "math"

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimestate() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimestate() {
}

type Hx_Obj_VTable_loreline_runtimestate interface {
	Hx_Field_clear(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_node)
	Hx_Field_createOriginalFields()
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimestate struct {
	VTable                  Hx_Obj_VTable_loreline_runtimestate `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_scope          int32
	Hx_Field_fields         any
	Hx_Field_originalFields any
}

func Hx_Obj_loreline_runtimestate_CreateEmptyInstance() *Hx_Obj_loreline_runtimestate {
	obj := &Hx_Obj_loreline_runtimestate{}
	obj.VTable = obj
	obj.Hx_Field_scope = HxDefault[int32]()
	obj.Hx_Field_fields = HxDefault[any]()
	obj.Hx_Field_originalFields = HxDefault[any]()
	return obj
}

func Hx_Obj_loreline_runtimestate_CreateInstance(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_node, fields any, originalFields any) *Hx_Obj_loreline_runtimestate {
	obj := Hx_Obj_loreline_runtimestate_CreateEmptyInstance()
	obj.Hx_New(interpreter, node, fields, originalFields)
	return obj
}

func (this *Hx_Obj_loreline_runtimestate) Hx_New(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_node, fields any, originalFields any) {
	{
		this.Hx_Field_scope = -1
	}
	{
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(fields, nil))) {
			this.Hx_Field_fields = fields
		} else {
			this.VTable.Hx_Field_clear(interpreter, node)
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(originalFields, nil))) {
			this.Hx_Field_originalFields = originalFields
		} else {
			this.VTable.Hx_Field_createOriginalFields()
		}
	}

}

var Hx_Obj_loreline_runtimestate_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeState",
	HxMakeArray[string](),
	HxMakeArray[string]("scope", "fields", "originalFields", "clear", "createOriginalFields"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimestate_CreateInstance(HxConvert[*Hx_Obj_loreline_interpreter](params.Get(0)), HxConvert[*Hx_Obj_loreline_node](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_runtimestate_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimestate_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimestate) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimestate_RTTI
}

func (this *Hx_Obj_loreline_runtimestate) Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate {
	return this
}

func (this *Hx_Obj_loreline_runtimestate) Hx_Field_clear(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_node) {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = interpreter
	this.Hx_Field_fields = Hx_Field_loreline_objects_createFields(_hx_tmp_0, nil, node)
}

func (this *Hx_Obj_loreline_runtimestate) Hx_Field_createOriginalFields() {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = nil
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = nil
	this.Hx_Field_originalFields = Hx_Field_loreline_objects_createFields(_hx_tmp_0, _hx_tmp_1, nil)
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimecharacter() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimecharacter() {
}

type Hx_Obj_VTable_loreline_runtimecharacter interface {
	Hx_Field_clear(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_node)
	Hx_Field_createOriginalFields()
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimecharacter struct {
	Hx_Obj_loreline_runtimestate
	VTable                Hx_Obj_VTable_loreline_runtimecharacter `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_characterRef *Hx_Obj_loreline_runtimecharacterref
}

func Hx_Obj_loreline_runtimecharacter_CreateEmptyInstance() *Hx_Obj_loreline_runtimecharacter {
	obj := &Hx_Obj_loreline_runtimecharacter{}
	obj.Hx_Obj_loreline_runtimestate = *Hx_Obj_loreline_runtimestate_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_runtimestate.VTable = obj
	obj.Hx_Field_characterRef = HxDefault[*Hx_Obj_loreline_runtimecharacterref]()
	return obj
}

func Hx_Obj_loreline_runtimecharacter_CreateInstance(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_astnode, fields any, originalFields any) *Hx_Obj_loreline_runtimecharacter {
	obj := Hx_Obj_loreline_runtimecharacter_CreateEmptyInstance()
	obj.Hx_New(interpreter, node, fields, originalFields)
	return obj
}

func (this *Hx_Obj_loreline_runtimecharacter) Hx_New(interpreter *Hx_Obj_loreline_interpreter, node *Hx_Obj_loreline_astnode, fields any, originalFields any) {
	{
		this.Hx_Field_characterRef = nil
	}
	{
		var _hx_tmp_0 *Hx_Obj_loreline_interpreter
		_ = _hx_tmp_0
		_hx_tmp_0 = interpreter
		this.Hx_Obj_loreline_runtimestate.Hx_New(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := node
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), fields, originalFields)
	}

}

var Hx_Obj_loreline_runtimecharacter_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeCharacter",
	HxMakeArray[string](),
	HxMakeArray[string]("characterRef", "scope", "fields", "originalFields", "clear", "createOriginalFields"),
	Hx_Obj_loreline_runtimestate_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimecharacter_CreateInstance(HxConvert[*Hx_Obj_loreline_interpreter](params.Get(0)), HxConvert[*Hx_Obj_loreline_astnode](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_runtimecharacter_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimecharacter_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimecharacter) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimecharacter_RTTI
}

func (this *Hx_Obj_loreline_runtimecharacter) Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter {
	return this
}

type Hx_Enum_loreline_runtimeaccess interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_runtimeaccess()
}

type Hx_Enum_loreline_runtimeaccess_FunctionAccess struct {
	Hx_Field_pos  *Hx_Obj_loreline_position
	Hx_Field_name string
}

func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) M_Hx_Enum_loreline_runtimeaccess() {}
func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) Hx_Field_enumIndex() int32         { return 3 }
func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_runtimeaccess_RTTI
}
func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_pos), any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_runtimeaccess_FunctionAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_pos)
	case 1:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_runtimeaccess_FieldAccess struct {
	Hx_Field_pos  *Hx_Obj_loreline_position
	Hx_Field_obj  any
	Hx_Field_name string
}

func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) M_Hx_Enum_loreline_runtimeaccess() {}
func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) Hx_Field_enumIndex() int32         { return 0 }
func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_runtimeaccess_RTTI
}
func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_pos), any(this.Hx_Field_obj), any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_runtimeaccess_FieldAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_pos)
	case 1:
		return any(this.Hx_Field_obj)
	case 2:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_runtimeaccess_CharacterAccess struct {
	Hx_Field_pos  *Hx_Obj_loreline_position
	Hx_Field_name string
}

func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) M_Hx_Enum_loreline_runtimeaccess() {}
func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) Hx_Field_enumIndex() int32         { return 2 }
func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_runtimeaccess_RTTI
}
func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_pos), any(this.Hx_Field_name))
}
func (this Hx_Enum_loreline_runtimeaccess_CharacterAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_pos)
	case 1:
		return any(this.Hx_Field_name)
	default:
		return nil
	}
}

type Hx_Enum_loreline_runtimeaccess_BeatAccess struct {
	Hx_Field_pos *Hx_Obj_loreline_position
	Hx_Field_ref *Hx_Obj_loreline_runtimebeatref
}

func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) M_Hx_Enum_loreline_runtimeaccess() {}
func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) Hx_Field_enumIndex() int32         { return 4 }
func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_runtimeaccess_RTTI
}
func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_pos), any(this.Hx_Field_ref))
}
func (this Hx_Enum_loreline_runtimeaccess_BeatAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_pos)
	case 1:
		return any(this.Hx_Field_ref)
	default:
		return nil
	}
}

type Hx_Enum_loreline_runtimeaccess_ArrayAccess struct {
	Hx_Field_pos   *Hx_Obj_loreline_position
	Hx_Field_array any
	Hx_Field_index int32
}

func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return nil
}
func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) M_Hx_Enum_loreline_runtimeaccess() {}
func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) Hx_Field_enumIndex() int32         { return 1 }
func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_runtimeaccess_RTTI
}
func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_pos), any(this.Hx_Field_array), any(this.Hx_Field_index))
}
func (this Hx_Enum_loreline_runtimeaccess_ArrayAccess) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_pos)
	case 1:
		return any(this.Hx_Field_array)
	case 2:
		return any(this.Hx_Field_index)
	default:
		return nil
	}
}

var Hx_Enum_loreline_runtimeaccess_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.RuntimeAccess",
	HxMakeArray[string]("FieldAccess", "ArrayAccess", "CharacterAccess", "FunctionAccess", "BeatAccess"),
	HxMakeArray[int32](3, 3, 2, 2, 2),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_runtimeaccess_FieldAccess{params.Get_Dyn(0).(*Hx_Obj_loreline_position), params.Get_Dyn(1).(any), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(2))))}
		case 1:
			return Hx_Enum_loreline_runtimeaccess_ArrayAccess{params.Get_Dyn(0).(*Hx_Obj_loreline_position), params.Get_Dyn(1).(any), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(params.Get_Dyn(2))))}
		case 2:
			return Hx_Enum_loreline_runtimeaccess_CharacterAccess{params.Get_Dyn(0).(*Hx_Obj_loreline_position), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1))))}
		case 3:
			return Hx_Enum_loreline_runtimeaccess_FunctionAccess{params.Get_Dyn(0).(*Hx_Obj_loreline_position), ((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(1))))}
		case 4:
			return Hx_Enum_loreline_runtimeaccess_BeatAccess{params.Get_Dyn(0).(*Hx_Obj_loreline_position), params.Get_Dyn(1).(*Hx_Obj_loreline_runtimebeatref)}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimescope() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimescope() {
}

type Hx_Obj_VTable_loreline_runtimescope interface {
	Hx_Field_beatByName(name string) *Hx_Obj_loreline_nbeatdecl
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimescope struct {
	VTable             Hx_Obj_VTable_loreline_runtimescope `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_id        int32
	Hx_Field_beat      *Hx_Obj_loreline_nbeatdecl
	Hx_Field_node      *Hx_Obj_loreline_astnode
	Hx_Field_beats     HxArray[*Hx_Obj_loreline_nbeatdecl]
	Hx_Field_state     *Hx_Obj_loreline_runtimestate
	Hx_Field_head      *Hx_Obj_loreline_astnode
	Hx_Field_insertion *Hx_Obj_loreline_runtimeinsertion
	Hx_Field_captured  HxArray[*Hx_Obj_loreline_runtimescope]
	Hx_Field_shuffle   HxArray[int32]
}

func Hx_Obj_loreline_runtimescope_CreateEmptyInstance() *Hx_Obj_loreline_runtimescope {
	obj := &Hx_Obj_loreline_runtimescope{}
	obj.VTable = obj
	obj.Hx_Field_id = HxDefault[int32]()
	obj.Hx_Field_beat = HxDefault[*Hx_Obj_loreline_nbeatdecl]()
	obj.Hx_Field_node = HxDefault[*Hx_Obj_loreline_astnode]()
	obj.Hx_Field_beats = HxDefault[HxArray[*Hx_Obj_loreline_nbeatdecl]]()
	obj.Hx_Field_state = HxDefault[*Hx_Obj_loreline_runtimestate]()
	obj.Hx_Field_head = HxDefault[*Hx_Obj_loreline_astnode]()
	obj.Hx_Field_insertion = HxDefault[*Hx_Obj_loreline_runtimeinsertion]()
	obj.Hx_Field_captured = HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	obj.Hx_Field_shuffle = HxDefault[HxArray[int32]]()
	return obj
}

func Hx_Obj_loreline_runtimescope_CreateInstance(id__isOptional any, beat *Hx_Obj_loreline_nbeatdecl, node *Hx_Obj_loreline_astnode, beats__isOptional any, state__isOptional any, head__isOptional any, insertion__isOptional any, captured__isOptional any, shuffle__isOptional any) *Hx_Obj_loreline_runtimescope {
	obj := Hx_Obj_loreline_runtimescope_CreateEmptyInstance()
	obj.Hx_New(id__isOptional, beat, node, beats__isOptional, state__isOptional, head__isOptional, insertion__isOptional, captured__isOptional, shuffle__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_runtimescope) Hx_New(id__isOptional any, beat *Hx_Obj_loreline_nbeatdecl, node *Hx_Obj_loreline_astnode, beats__isOptional any, state__isOptional any, head__isOptional any, insertion__isOptional any, captured__isOptional any, shuffle__isOptional any) {
	{
		id := HxDefault[HxNullable[int32]]()
		_ = id

		if id__isOptional != nil {
			id = HxConvert[HxNullable[int32]](id__isOptional)
		} else {
		}
		beats := HxDefault[HxArray[*Hx_Obj_loreline_nbeatdecl]]()
		_ = beats

		if beats__isOptional != nil {
			beats = HxConvert[HxArray[*Hx_Obj_loreline_nbeatdecl]](beats__isOptional)
		} else {
		}
		state := HxDefault[*Hx_Obj_loreline_runtimestate]()
		_ = state

		if state__isOptional != nil {
			state = HxConvert[*Hx_Obj_loreline_runtimestate](state__isOptional)
		} else {
		}
		head := HxDefault[*Hx_Obj_loreline_astnode]()
		_ = head

		if head__isOptional != nil {
			head = HxConvert[*Hx_Obj_loreline_astnode](head__isOptional)
		} else {
		}
		insertion := HxDefault[*Hx_Obj_loreline_runtimeinsertion]()
		_ = insertion

		if insertion__isOptional != nil {
			insertion = HxConvert[*Hx_Obj_loreline_runtimeinsertion](insertion__isOptional)
		} else {
		}
		captured := HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
		_ = captured

		if captured__isOptional != nil {
			captured = HxConvert[HxArray[*Hx_Obj_loreline_runtimescope]](captured__isOptional)
		} else {
		}
		shuffle := HxDefault[HxArray[int32]]()
		_ = shuffle

		if shuffle__isOptional != nil {
			shuffle = HxConvert[HxArray[int32]](shuffle__isOptional)
		} else {
		}
		{
			this.Hx_Field_id = -1
		}
		{
			this.Hx_Field_beats = nil
		}
		{
			this.Hx_Field_state = nil
		}
		{
			this.Hx_Field_head = nil
		}
		{
			this.Hx_Field_insertion = nil
		}
		{
			this.Hx_Field_captured = nil
		}
		{
			this.Hx_Field_shuffle = nil
		}
		{
			if (id).Valid != false {
				this.Hx_Field_id = (id).Value
			}

			this.Hx_Field_beat = beat
			this.Hx_Field_node = node
			if beats != nil {
				this.Hx_Field_beats = beats
			}

			if state != nil {
				this.Hx_Field_state = state
			}

			if head != nil {
				this.Hx_Field_head = head
			}

			if insertion != nil {
				this.Hx_Field_insertion = insertion
			}

			if captured != nil {
				this.Hx_Field_captured = captured
			}

			if shuffle != nil {
				this.Hx_Field_shuffle = shuffle
			}
		}

	}
}

var Hx_Obj_loreline_runtimescope_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeScope",
	HxMakeArray[string](),
	HxMakeArray[string]("id", "beat", "node", "beats", "state", "head", "insertion", "captured", "shuffle", "beatByName"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimescope_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[*Hx_Obj_loreline_nbeatdecl](params.Get(1)), HxConvert[*Hx_Obj_loreline_astnode](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)), HxConvert[any](params.Get(7)), HxConvert[any](params.Get(8)))
	},
	func() any {
		return Hx_Obj_loreline_runtimescope_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimescope_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimescope) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimescope_RTTI
}

func (this *Hx_Obj_loreline_runtimescope) Hx_As_Hx_Obj_loreline_runtimescope() *Hx_Obj_loreline_runtimescope {
	return this
}

func (this *Hx_Obj_loreline_runtimescope) Hx_Field_beatByName(name string) *Hx_Obj_loreline_nbeatdecl {
	if this.Hx_Field_beats == nil {
		return nil
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_beats.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var beat *Hx_Obj_loreline_nbeatdecl
			_ = beat
			beat = this.Hx_Field_beats.Get(i)
			if beat.Hx_Field_name == name {
				return beat
			}
		}
	}

	return nil
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimebeatref() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimebeatref() {
}

type Hx_Obj_VTable_loreline_runtimebeatref interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimebeatref struct {
	VTable         Hx_Obj_VTable_loreline_runtimebeatref `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_beat  *Hx_Obj_loreline_nbeatdecl
	Hx_Field_stack HxArray[*Hx_Obj_loreline_runtimescope]
}

func Hx_Obj_loreline_runtimebeatref_CreateEmptyInstance() *Hx_Obj_loreline_runtimebeatref {
	obj := &Hx_Obj_loreline_runtimebeatref{}
	obj.VTable = obj
	obj.Hx_Field_beat = HxDefault[*Hx_Obj_loreline_nbeatdecl]()
	obj.Hx_Field_stack = HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	return obj
}

func Hx_Obj_loreline_runtimebeatref_CreateInstance(beat *Hx_Obj_loreline_nbeatdecl, stack HxArray[*Hx_Obj_loreline_runtimescope]) *Hx_Obj_loreline_runtimebeatref {
	obj := Hx_Obj_loreline_runtimebeatref_CreateEmptyInstance()
	obj.Hx_New(beat, stack)
	return obj
}

func (this *Hx_Obj_loreline_runtimebeatref) Hx_New(beat *Hx_Obj_loreline_nbeatdecl, stack HxArray[*Hx_Obj_loreline_runtimescope]) {
	{
		this.Hx_Field_beat = beat
		this.Hx_Field_stack = stack
	}

}

var Hx_Obj_loreline_runtimebeatref_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeBeatRef",
	HxMakeArray[string]("beatOf", "beatNameOf"),
	HxMakeArray[string]("beat", "stack"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimebeatref_CreateInstance(HxConvert[*Hx_Obj_loreline_nbeatdecl](params.Get(0)), HxConvert[HxArray[*Hx_Obj_loreline_runtimescope]](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_runtimebeatref_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimebeatref_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "beatOf":
			return Hx_Field_loreline_runtimebeatref_beatOf
		case "beatNameOf":
			return Hx_Field_loreline_runtimebeatref_beatNameOf
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimebeatref) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimebeatref_RTTI
}

func (this *Hx_Obj_loreline_runtimebeatref) Hx_As_Hx_Obj_loreline_runtimebeatref() *Hx_Obj_loreline_runtimebeatref {
	return this
}

func Hx_Field_loreline_runtimebeatref_beatOf(value any) *Hx_Obj_loreline_nbeatdecl {
	if Hx_Field_std_isOfType(value, Hx_Obj_loreline_nbeatdecl_RTTI) {
		var hx_dyncast_207 any
		_ = hx_dyncast_207
		hx_dyncast_207 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_nbeatdecl")
		var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_0
		if hx_dyncast_207 != nil {
			_hx_tmp_0 = hx_dyncast_207.(interface {
				Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
			}).Hx_As_Hx_Obj_loreline_nbeatdecl()
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_loreline_runtimebeatref_RTTI) {
		var hx_dyncast_208 any
		_ = hx_dyncast_208
		hx_dyncast_208 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_runtimebeatref")
		var _hx_tmp_1 *Hx_Obj_loreline_runtimebeatref
		_ = _hx_tmp_1
		if hx_dyncast_208 != nil {
			_hx_tmp_1 = hx_dyncast_208.(interface {
				Hx_As_Hx_Obj_loreline_runtimebeatref() *Hx_Obj_loreline_runtimebeatref
			}).Hx_As_Hx_Obj_loreline_runtimebeatref()
		} else {
			_hx_tmp_1 = nil
		}

		var ref *Hx_Obj_loreline_runtimebeatref
		_ = ref
		ref = _hx_tmp_1
		return ref.Hx_Field_beat
	}

	return nil
}

func Hx_Field_loreline_runtimebeatref_beatNameOf(value any) HxNullable[string] {
	var beat *Hx_Obj_loreline_nbeatdecl
	_ = beat
	beat = Hx_Field_loreline_runtimebeatref_beatOf(value)
	if beat != nil {
		return HxNullableString(beat.Hx_Field_name)
	} else {
		return HxNullable[string]{Value: HxStringNull}
	}
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimecharacterref() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimecharacterref() {
}

type Hx_Obj_VTable_loreline_runtimecharacterref interface {
	Hx_Field_get_fields() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimecharacterref struct {
	VTable             Hx_Obj_VTable_loreline_runtimecharacterref `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name      string
	Hx_Field_character *Hx_Obj_loreline_runtimecharacter
}

func Hx_Obj_loreline_runtimecharacterref_CreateEmptyInstance() *Hx_Obj_loreline_runtimecharacterref {
	obj := &Hx_Obj_loreline_runtimecharacterref{}
	obj.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_character = HxDefault[*Hx_Obj_loreline_runtimecharacter]()
	return obj
}

func Hx_Obj_loreline_runtimecharacterref_CreateInstance(name string, character *Hx_Obj_loreline_runtimecharacter) *Hx_Obj_loreline_runtimecharacterref {
	obj := Hx_Obj_loreline_runtimecharacterref_CreateEmptyInstance()
	obj.Hx_New(name, character)
	return obj
}

func (this *Hx_Obj_loreline_runtimecharacterref) Hx_New(name string, character *Hx_Obj_loreline_runtimecharacter) {
	{
		this.Hx_Field_name = name
		this.Hx_Field_character = character
	}

}

var Hx_Obj_loreline_runtimecharacterref_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeCharacterRef",
	HxMakeArray[string]("characterOf", "characterNameOf", "fieldsOf"),
	HxMakeArray[string]("name", "character", "get_fields"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimecharacterref_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[*Hx_Obj_loreline_runtimecharacter](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_runtimecharacterref_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimecharacterref_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "characterOf":
			return Hx_Field_loreline_runtimecharacterref_characterOf
		case "characterNameOf":
			return Hx_Field_loreline_runtimecharacterref_characterNameOf
		case "fieldsOf":
			return Hx_Field_loreline_runtimecharacterref_fieldsOf
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimecharacterref) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimecharacterref_RTTI
}

func (this *Hx_Obj_loreline_runtimecharacterref) Hx_As_Hx_Obj_loreline_runtimecharacterref() *Hx_Obj_loreline_runtimecharacterref {
	return this
}

func (this *Hx_Obj_loreline_runtimecharacterref) Hx_Field_get_fields() any {
	return this.Hx_Field_character.Hx_Field_fields
}

func Hx_Field_loreline_runtimecharacterref_characterOf(value any) *Hx_Obj_loreline_runtimecharacterref {
	if Hx_Field_std_isOfType(value, Hx_Obj_loreline_runtimecharacterref_RTTI) {
		var hx_dyncast_209 any
		_ = hx_dyncast_209
		hx_dyncast_209 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_runtimecharacterref")
		var _hx_tmp_0 *Hx_Obj_loreline_runtimecharacterref
		_ = _hx_tmp_0
		if hx_dyncast_209 != nil {
			_hx_tmp_0 = hx_dyncast_209.(interface {
				Hx_As_Hx_Obj_loreline_runtimecharacterref() *Hx_Obj_loreline_runtimecharacterref
			}).Hx_As_Hx_Obj_loreline_runtimecharacterref()
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0
	}

	return nil
}

func Hx_Field_loreline_runtimecharacterref_characterNameOf(value any) HxNullable[string] {
	var ref *Hx_Obj_loreline_runtimecharacterref
	_ = ref
	ref = Hx_Field_loreline_runtimecharacterref_characterOf(value)
	if ref != nil {
		return HxNullableString(ref.Hx_Field_name)
	} else {
		return HxNullable[string]{Value: HxStringNull}
	}
}

func Hx_Field_loreline_runtimecharacterref_fieldsOf(value any) any {
	var ref *Hx_Obj_loreline_runtimecharacterref
	_ = ref
	ref = Hx_Field_loreline_runtimecharacterref_characterOf(value)
	if ref != nil {
		return ref.VTable.Hx_Field_get_fields()
	} else {
		return value
	}
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimeinsertion() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimeinsertion() {
}

type Hx_Obj_VTable_loreline_runtimeinsertion interface {
	Hx_Field_isCollected() bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimeinsertion struct {
	VTable                         Hx_Obj_VTable_loreline_runtimeinsertion `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_id                    int32
	Hx_Field_origin                *Hx_Obj_loreline_ninsertion
	Hx_Field_options               HxArray[*Hx_Obj_loreline_choiceoption]
	Hx_Field_stack                 HxArray[*Hx_Obj_loreline_runtimescope]
	Hx_Field_parentPartialOptions  HxArray[*Hx_Obj_loreline_choiceoption]
	Hx_Field_parentNextOptionIndex int32
	Hx_Field_kind                  int32
	Hx_Field_rules                 HxArray[any]
	Hx_Field_parentPartialRules    HxArray[any]
	Hx_Field_parentNextRuleIndex   int32
	Hx_Field_whenFirst             bool
}

func Hx_Obj_loreline_runtimeinsertion_CreateEmptyInstance() *Hx_Obj_loreline_runtimeinsertion {
	obj := &Hx_Obj_loreline_runtimeinsertion{}
	obj.VTable = obj
	obj.Hx_Field_id = HxDefault[int32]()
	obj.Hx_Field_origin = HxDefault[*Hx_Obj_loreline_ninsertion]()
	obj.Hx_Field_options = HxDefault[HxArray[*Hx_Obj_loreline_choiceoption]]()
	obj.Hx_Field_stack = HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	obj.Hx_Field_parentPartialOptions = HxDefault[HxArray[*Hx_Obj_loreline_choiceoption]]()
	obj.Hx_Field_parentNextOptionIndex = HxDefault[int32]()
	obj.Hx_Field_kind = HxDefault[int32]()
	obj.Hx_Field_rules = HxDefault[HxArray[any]]()
	obj.Hx_Field_parentPartialRules = HxDefault[HxArray[any]]()
	obj.Hx_Field_parentNextRuleIndex = HxDefault[int32]()
	obj.Hx_Field_whenFirst = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_runtimeinsertion_CreateInstance(id int32, origin *Hx_Obj_loreline_ninsertion) *Hx_Obj_loreline_runtimeinsertion {
	obj := Hx_Obj_loreline_runtimeinsertion_CreateEmptyInstance()
	obj.Hx_New(id, origin)
	return obj
}

func (this *Hx_Obj_loreline_runtimeinsertion) Hx_New(id int32, origin *Hx_Obj_loreline_ninsertion) {
	{
		this.Hx_Field_options = nil
	}
	{
		this.Hx_Field_stack = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
	}
	{
		this.Hx_Field_parentPartialOptions = nil
	}
	{
		this.Hx_Field_parentNextOptionIndex = 0
	}
	{
		this.Hx_Field_kind = ((int32)(0))
	}
	{
		this.Hx_Field_rules = nil
	}
	{
		this.Hx_Field_parentPartialRules = nil
	}
	{
		this.Hx_Field_parentNextRuleIndex = 0
	}
	{
		this.Hx_Field_whenFirst = false
	}
	{
		this.Hx_Field_id = id
		this.Hx_Field_origin = origin
	}

}

var Hx_Obj_loreline_runtimeinsertion_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeInsertion",
	HxMakeArray[string](),
	HxMakeArray[string]("id", "origin", "options", "stack", "parentPartialOptions", "parentNextOptionIndex", "kind", "rules", "parentPartialRules", "parentNextRuleIndex", "whenFirst", "isCollected"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimeinsertion_CreateInstance(HxConvert[int32](params.Get(0)), HxConvert[*Hx_Obj_loreline_ninsertion](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_runtimeinsertion_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimeinsertion_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimeinsertion) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimeinsertion_RTTI
}

func (this *Hx_Obj_loreline_runtimeinsertion) Hx_As_Hx_Obj_loreline_runtimeinsertion() *Hx_Obj_loreline_runtimeinsertion {
	return this
}

func (this *Hx_Obj_loreline_runtimeinsertion) Hx_Field_isCollected() bool {
	return ((this.Hx_Field_options != nil) || (this.Hx_Field_rules != nil))
}

var Hx_Abstract_loreline_insertionkind_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.InsertionKind",
	HxMakeArray[string]("Choice", "When"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__interpreter_insertionkind_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__interpreter_insertionkind_impl_() {
	Hx_Field_loreline__interpreter_insertionkind_impl__Choice = Hx_Init_Hx_Field_loreline__interpreter_insertionkind_impl__Choice()
	Hx_Field_loreline__interpreter_insertionkind_impl__When = Hx_Init_Hx_Field_loreline__interpreter_insertionkind_impl__When()
}

type Hx_Obj_VTable_loreline__interpreter_insertionkind_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__interpreter_insertionkind_impl_ struct {
	VTable Hx_Obj_VTable_loreline__interpreter_insertionkind_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__interpreter_insertionkind_impl__CreateEmptyInstance() *Hx_Obj_loreline__interpreter_insertionkind_impl_ {
	obj := &Hx_Obj_loreline__interpreter_insertionkind_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__interpreter_insertionkind_impl__CreateInstance() *Hx_Obj_loreline__interpreter_insertionkind_impl_ {
	obj := Hx_Obj_loreline__interpreter_insertionkind_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__interpreter_insertionkind_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Interpreter.InsertionKind_Impl_",
	HxMakeArray[string]("Choice", "When"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__interpreter_insertionkind_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__interpreter_insertionkind_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__interpreter_insertionkind_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Choice":
			return &Hx_Field_loreline__interpreter_insertionkind_impl__Choice
		case "When":
			return &Hx_Field_loreline__interpreter_insertionkind_impl__When
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__interpreter_insertionkind_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__interpreter_insertionkind_impl__RTTI
}

func (this *Hx_Obj_loreline__interpreter_insertionkind_impl_) Hx_As_Hx_Obj_loreline__interpreter_insertionkind_impl_() *Hx_Obj_loreline__interpreter_insertionkind_impl_ {
	return this
}

var Hx_Field_loreline__interpreter_insertionkind_impl__Choice int32

func Hx_Init_Hx_Field_loreline__interpreter_insertionkind_impl__Choice() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__interpreter_insertionkind_impl__When int32

func Hx_Init_Hx_Field_loreline__interpreter_insertionkind_impl__When() int32 {
	return ((int32)(1))
}

type Hx_Typedef_loreline_whencandidate = any

func Hx_ClassInit_Hx_Obj_VTable_loreline_texttag() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_texttag() {
}

type Hx_Obj_VTable_loreline_texttag interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_texttag struct {
	VTable           Hx_Obj_VTable_loreline_texttag `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_closing bool
	Hx_Field_value   string
	Hx_Field_offset  int32
}

func Hx_Obj_loreline_texttag_CreateEmptyInstance() *Hx_Obj_loreline_texttag {
	obj := &Hx_Obj_loreline_texttag{}
	obj.VTable = obj
	obj.Hx_Field_closing = HxDefault[bool]()
	obj.Hx_Field_value = HxDefault[string]()
	obj.Hx_Field_offset = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_texttag_CreateInstance(closing bool, value string, offset int32) *Hx_Obj_loreline_texttag {
	obj := Hx_Obj_loreline_texttag_CreateEmptyInstance()
	obj.Hx_New(closing, value, offset)
	return obj
}

func (this *Hx_Obj_loreline_texttag) Hx_New(closing bool, value string, offset int32) {
	{
		this.Hx_Field_closing = closing
		this.Hx_Field_value = value
		this.Hx_Field_offset = offset
	}

}

var Hx_Obj_loreline_texttag_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.TextTag",
	HxMakeArray[string](),
	HxMakeArray[string]("closing", "value", "offset"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_texttag_CreateInstance(HxConvert[bool](params.Get(0)), HxConvert[string](params.Get(1)), HxConvert[int32](params.Get(2)))
	},
	func() any {
		return Hx_Obj_loreline_texttag_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_texttag_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_texttag) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_texttag_RTTI
}

func (this *Hx_Obj_loreline_texttag) Hx_As_Hx_Obj_loreline_texttag() *Hx_Obj_loreline_texttag {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_choiceoption() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_choiceoption() {
}

type Hx_Obj_VTable_loreline_choiceoption interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_choiceoption struct {
	VTable             Hx_Obj_VTable_loreline_choiceoption `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_text      string
	Hx_Field_tags      HxArray[*Hx_Obj_loreline_texttag]
	Hx_Field_enabled   bool
	Hx_Field_node      *Hx_Obj_loreline_nchoiceoption
	Hx_Field_insertion *Hx_Obj_loreline_runtimeinsertion
}

func Hx_Obj_loreline_choiceoption_CreateEmptyInstance() *Hx_Obj_loreline_choiceoption {
	obj := &Hx_Obj_loreline_choiceoption{}
	obj.VTable = obj
	obj.Hx_Field_text = HxDefault[string]()
	obj.Hx_Field_tags = HxDefault[HxArray[*Hx_Obj_loreline_texttag]]()
	obj.Hx_Field_enabled = HxDefault[bool]()
	obj.Hx_Field_node = HxDefault[*Hx_Obj_loreline_nchoiceoption]()
	obj.Hx_Field_insertion = HxDefault[*Hx_Obj_loreline_runtimeinsertion]()
	return obj
}

func Hx_Obj_loreline_choiceoption_CreateInstance(text string, tags HxArray[*Hx_Obj_loreline_texttag], enabled bool, node *Hx_Obj_loreline_nchoiceoption, insertion *Hx_Obj_loreline_runtimeinsertion) *Hx_Obj_loreline_choiceoption {
	obj := Hx_Obj_loreline_choiceoption_CreateEmptyInstance()
	obj.Hx_New(text, tags, enabled, node, insertion)
	return obj
}

func (this *Hx_Obj_loreline_choiceoption) Hx_New(text string, tags HxArray[*Hx_Obj_loreline_texttag], enabled bool, node *Hx_Obj_loreline_nchoiceoption, insertion *Hx_Obj_loreline_runtimeinsertion) {
	{
		this.Hx_Field_text = text
		this.Hx_Field_tags = tags
		this.Hx_Field_enabled = enabled
		this.Hx_Field_node = node
		this.Hx_Field_insertion = insertion
	}

}

var Hx_Obj_loreline_choiceoption_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.ChoiceOption",
	HxMakeArray[string](),
	HxMakeArray[string]("text", "tags", "enabled", "node", "insertion"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_choiceoption_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[HxArray[*Hx_Obj_loreline_texttag]](params.Get(1)), HxConvert[bool](params.Get(2)), HxConvert[*Hx_Obj_loreline_nchoiceoption](params.Get(3)), HxConvert[*Hx_Obj_loreline_runtimeinsertion](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_choiceoption_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_choiceoption_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_choiceoption) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_choiceoption_RTTI
}

func (this *Hx_Obj_loreline_choiceoption) Hx_As_Hx_Obj_loreline_choiceoption() *Hx_Obj_loreline_choiceoption {
	return this
}

type Hx_Typedef_loreline_dialoguehandler func(interpreter *Hx_Obj_loreline_interpreter, character string, text string, tags HxArray[*Hx_Obj_loreline_texttag], callback func())

type Hx_Typedef_loreline_choicehandler func(interpreter *Hx_Obj_loreline_interpreter, options HxArray[*Hx_Obj_loreline_choiceoption], callback func(index int32))

type Hx_Typedef_loreline_finishhandler func(interpreter *Hx_Obj_loreline_interpreter)

func Hx_ClassInit_Hx_Obj_VTable_loreline_runtimeerror() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_runtimeerror() {
}

type Hx_Obj_VTable_loreline_runtimeerror interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_runtimeerror struct {
	Hx_Obj_loreline_error
	VTable Hx_Obj_VTable_loreline_runtimeerror `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_runtimeerror_CreateEmptyInstance() *Hx_Obj_loreline_runtimeerror {
	obj := &Hx_Obj_loreline_runtimeerror{}
	obj.Hx_Obj_loreline_error = *Hx_Obj_loreline_error_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_error.VTable = obj
	return obj
}

func Hx_Obj_loreline_runtimeerror_CreateInstance(message string, pos__isOptional any) *Hx_Obj_loreline_runtimeerror {
	obj := Hx_Obj_loreline_runtimeerror_CreateEmptyInstance()
	obj.Hx_New(message, pos__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_runtimeerror) Hx_New(message string, pos__isOptional any) {
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

var Hx_Obj_loreline_runtimeerror_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.RuntimeError",
	HxMakeArray[string](),
	HxMakeArray[string]("message", "pos", "stack", "toString"),
	Hx_Obj_loreline_error_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_runtimeerror_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[any](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_runtimeerror_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_runtimeerror_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_runtimeerror) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_runtimeerror_RTTI
}

func (this *Hx_Obj_loreline_runtimeerror) Hx_As_Hx_Obj_loreline_runtimeerror() *Hx_Obj_loreline_runtimeerror {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_evalnext() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_evalnext() {
}

type Hx_Obj_VTable_loreline_evalnext interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_evalnext struct {
	VTable        Hx_Obj_VTable_loreline_evalnext `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_sync bool
	Hx_Field_cb   func()
}

func Hx_Obj_loreline_evalnext_CreateEmptyInstance() *Hx_Obj_loreline_evalnext {
	obj := &Hx_Obj_loreline_evalnext{}
	obj.VTable = obj
	obj.Hx_Field_sync = HxDefault[bool]()
	obj.Hx_Field_cb = HxDefault[func()]()
	return obj
}

func Hx_Obj_loreline_evalnext_CreateInstance() *Hx_Obj_loreline_evalnext {
	obj := Hx_Obj_loreline_evalnext_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_evalnext) Hx_New() {
	{
		this.Hx_Field_sync = true
	}
	{
		this.Hx_Field_cb = nil
	}
	{
	}

}

var Hx_Obj_loreline_evalnext_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.EvalNext",
	HxMakeArray[string](),
	HxMakeArray[string]("sync", "cb"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_evalnext_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_evalnext_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_evalnext_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_evalnext) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_evalnext_RTTI
}

func (this *Hx_Obj_loreline_evalnext) Hx_As_Hx_Obj_loreline_evalnext() *Hx_Obj_loreline_evalnext {
	return this
}

type Hx_Typedef_loreline_functionsmap = *Hx_Obj_haxe_imap

func Hx_ClassInit_Hx_Obj_VTable_loreline_interpreteroptions() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_interpreteroptions() {
}

type Hx_Obj_VTable_loreline_interpreteroptions interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_interpreteroptions struct {
	VTable                           Hx_Obj_VTable_loreline_interpreteroptions `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_wrapper                 any
	Hx_Field_functions               Hx_Typedef_loreline_functionsmap
	Hx_Field_strictAccess            bool
	Hx_Field_customCreateFields      func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
	Hx_Field_translations            Hx_Typedef_map
	Hx_Field_stringLiteralProcessors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
}

func Hx_Obj_loreline_interpreteroptions_CreateEmptyInstance() *Hx_Obj_loreline_interpreteroptions {
	obj := &Hx_Obj_loreline_interpreteroptions{}
	obj.VTable = obj
	obj.Hx_Field_wrapper = HxDefault[any]()
	obj.Hx_Field_functions = HxDefault[Hx_Typedef_loreline_functionsmap]()
	obj.Hx_Field_strictAccess = HxDefault[bool]()
	obj.Hx_Field_customCreateFields = HxDefault[func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any]()
	obj.Hx_Field_translations = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_stringLiteralProcessors = HxDefault[HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]]()
	return obj
}

func Hx_Obj_loreline_interpreteroptions_CreateInstance(wrapper__isOptional any, functions__isOptional any, strictAccess__isOptional any, customCreateFields__isOptional any, translations__isOptional any, stringLiteralProcessors__isOptional any) *Hx_Obj_loreline_interpreteroptions {
	obj := Hx_Obj_loreline_interpreteroptions_CreateEmptyInstance()
	obj.Hx_New(wrapper__isOptional, functions__isOptional, strictAccess__isOptional, customCreateFields__isOptional, translations__isOptional, stringLiteralProcessors__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_interpreteroptions) Hx_New(wrapper__isOptional any, functions__isOptional any, strictAccess__isOptional any, customCreateFields__isOptional any, translations__isOptional any, stringLiteralProcessors__isOptional any) {
	{
		wrapper := HxDefault[any]()
		_ = wrapper

		if wrapper__isOptional != nil {
			wrapper = HxConvert[any](wrapper__isOptional)
		} else {
		}
		functions := HxDefault[Hx_Typedef_loreline_functionsmap]()
		_ = functions

		if functions__isOptional != nil {
			functions = HxConvert[Hx_Typedef_loreline_functionsmap](functions__isOptional)
		} else {
		}
		strictAccess := HxDefault[HxNullable[bool]]()
		_ = strictAccess

		if strictAccess__isOptional != nil {
			strictAccess = HxConvert[HxNullable[bool]](strictAccess__isOptional)
		} else {
		}
		customCreateFields := HxDefault[func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any]()
		_ = customCreateFields

		if customCreateFields__isOptional != nil {
			customCreateFields = HxConvert[func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any](customCreateFields__isOptional)
		} else {
		}
		translations := HxDefault[Hx_Typedef_map]()
		_ = translations

		if translations__isOptional != nil {
			translations = HxConvert[Hx_Typedef_map](translations__isOptional)
		} else {
		}
		stringLiteralProcessors := HxDefault[HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]]()
		_ = stringLiteralProcessors

		if stringLiteralProcessors__isOptional != nil {
			stringLiteralProcessors = HxConvert[HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]](stringLiteralProcessors__isOptional)
		} else {
		}
		{
			this.Hx_Field_wrapper = nil
		}
		{
			this.Hx_Field_functions = nil
		}
		{
			this.Hx_Field_strictAccess = false
		}
		{
			this.Hx_Field_customCreateFields = nil
		}
		{
			this.Hx_Field_translations = nil
		}
		{
			this.Hx_Field_stringLiteralProcessors = nil
		}
		{
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(wrapper, nil))) {
				this.Hx_Field_wrapper = wrapper
			}

			if functions != nil {
				this.Hx_Field_functions = functions
			}

			if (strictAccess).Valid != false {
				this.Hx_Field_strictAccess = (strictAccess).Value
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(customCreateFields)), ((any)(nil))))) {
				this.Hx_Field_customCreateFields = customCreateFields
			}

			if translations != nil {
				this.Hx_Field_translations = translations
			}

			if stringLiteralProcessors != nil {
				this.Hx_Field_stringLiteralProcessors = stringLiteralProcessors
			}
		}

	}
}

var Hx_Obj_loreline_interpreteroptions_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.InterpreterOptions",
	HxMakeArray[string](),
	HxMakeArray[string]("wrapper", "functions", "strictAccess", "customCreateFields", "translations", "stringLiteralProcessors"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_interpreteroptions_CreateInstance(HxConvert[any](params.Get(0)), HxConvert[any](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_interpreteroptions_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_interpreteroptions_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_interpreteroptions) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_interpreteroptions_RTTI
}

func (this *Hx_Obj_loreline_interpreteroptions) Hx_As_Hx_Obj_loreline_interpreteroptions() *Hx_Obj_loreline_interpreteroptions {
	return this
}

type Hx_Typedef_loreline__interpreter_pendingdialogue = any

type Hx_Typedef_loreline__interpreter_scriptfunction = any

func Hx_ClassInit_Hx_Obj_VTable_loreline_interpretercontext() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_interpretercontext() {
}

type Hx_Obj_VTable_loreline_interpretercontext interface {
	Hx_Field_childByKey(key string) *Hx_Obj_loreline_interpreter
	Hx_Field_removePendingFlow(key string) any
	Hx_Field_clearPendingFlows()
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_interpretercontext struct {
	VTable                           Hx_Obj_VTable_loreline_interpretercontext `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_script                  *Hx_Obj_loreline_script
	Hx_Field_lens                    *Hx_Obj_loreline_lens
	Hx_Field_translations            Hx_Typedef_map
	Hx_Field_stringLiteralProcessors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
	Hx_Field_strictAccess            bool
	Hx_Field_customCreateFields      func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
	Hx_Field_topLevelState           *Hx_Obj_loreline_runtimestate
	Hx_Field_topLevelCharacters      Hx_Typedef_map
	Hx_Field_topLevelBeats           Hx_Typedef_map
	Hx_Field_nodeStates              *Hx_Obj_loreline_nodeidmap
	Hx_Field_random                  *Hx_Obj_loreline_random
	Hx_Field_hostFunctions           Hx_Typedef_loreline_functionsmap
	Hx_Field_scriptFunctions         HxArray[any]
	Hx_Field_nextInsertionId         int32
	Hx_Field_root                    *Hx_Obj_loreline_interpreter
	Hx_Field_children                HxArray[*Hx_Obj_loreline_interpreter]
	Hx_Field_pendingRestoredFlows    Hx_Typedef_map
	Hx_Field_pendingRestoredKeys     HxArray[string]
	Hx_Field_pendingSavedInsertions  any
	Hx_Field_pendingInsertionsCache  Hx_Typedef_map
}

func Hx_Obj_loreline_interpretercontext_CreateEmptyInstance() *Hx_Obj_loreline_interpretercontext {
	obj := &Hx_Obj_loreline_interpretercontext{}
	obj.VTable = obj
	obj.Hx_Field_script = HxDefault[*Hx_Obj_loreline_script]()
	obj.Hx_Field_lens = HxDefault[*Hx_Obj_loreline_lens]()
	obj.Hx_Field_translations = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_stringLiteralProcessors = HxDefault[HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]]()
	obj.Hx_Field_strictAccess = HxDefault[bool]()
	obj.Hx_Field_customCreateFields = HxDefault[func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any]()
	obj.Hx_Field_topLevelState = HxDefault[*Hx_Obj_loreline_runtimestate]()
	obj.Hx_Field_topLevelCharacters = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_topLevelBeats = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_nodeStates = HxDefault[*Hx_Obj_loreline_nodeidmap]()
	obj.Hx_Field_random = HxDefault[*Hx_Obj_loreline_random]()
	obj.Hx_Field_hostFunctions = HxDefault[Hx_Typedef_loreline_functionsmap]()
	obj.Hx_Field_scriptFunctions = HxDefault[HxArray[any]]()
	obj.Hx_Field_nextInsertionId = HxDefault[int32]()
	obj.Hx_Field_root = HxDefault[*Hx_Obj_loreline_interpreter]()
	obj.Hx_Field_children = HxDefault[HxArray[*Hx_Obj_loreline_interpreter]]()
	obj.Hx_Field_pendingRestoredFlows = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_pendingRestoredKeys = HxDefault[HxArray[string]]()
	obj.Hx_Field_pendingSavedInsertions = HxDefault[any]()
	obj.Hx_Field_pendingInsertionsCache = HxDefault[Hx_Typedef_map]()
	return obj
}

func Hx_Obj_loreline_interpretercontext_CreateInstance(script *Hx_Obj_loreline_script, options *Hx_Obj_loreline_interpreteroptions) *Hx_Obj_loreline_interpretercontext {
	obj := Hx_Obj_loreline_interpretercontext_CreateEmptyInstance()
	obj.Hx_New(script, options)
	return obj
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_New(script *Hx_Obj_loreline_script, options *Hx_Obj_loreline_interpreteroptions) {
	{
		this.Hx_Field_stringLiteralProcessors = HxMakeArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]()
	}
	{
		this.Hx_Field_topLevelState = nil
	}
	{
		this.Hx_Field_topLevelCharacters = (func() *Hx_Obj_haxe_imap {
			_hx_icast_210 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_210 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_210, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_210.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_topLevelBeats = (func() *Hx_Obj_haxe_imap {
			_hx_icast_211 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_211 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_211, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_211.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
	}
	{
		this.Hx_Field_nodeStates = Hx_Obj_loreline_nodeidmap_CreateInstance()
	}
	{
		this.Hx_Field_random = nil
	}
	{
		this.Hx_Field_scriptFunctions = HxMakeArray[any]().Dyn()
	}
	{
		this.Hx_Field_nextInsertionId = 1
	}
	{
		this.Hx_Field_root = nil
	}
	{
		this.Hx_Field_children = HxMakeArray[*Hx_Obj_loreline_interpreter]()
	}
	{
		this.Hx_Field_pendingRestoredFlows = nil
	}
	{
		this.Hx_Field_pendingRestoredKeys = nil
	}
	{
		this.Hx_Field_pendingSavedInsertions = nil
	}
	{
		this.Hx_Field_pendingInsertionsCache = nil
	}
	{
		this.Hx_Field_script = script
		this.Hx_Field_lens = Hx_Obj_loreline_lens_CreateInstance(script)
		var _hx_tmp_0 HxNullable[bool]
		_ = _hx_tmp_0
		if options != nil {
			_hx_tmp_0 = HxNullable[bool]{Value: options.Hx_Field_strictAccess, Valid: true}
		} else {
			_hx_tmp_0 = HxNullable[bool]{}
		}

		var tmp HxNullable[bool]
		_ = tmp
		tmp = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (tmp).Valid != false {
			_hx_tmp_1 = (tmp).Value
		} else {
			_hx_tmp_1 = false
		}

		this.Hx_Field_strictAccess = _hx_tmp_1
		var _hx_tmp_2 Hx_Typedef_map
		_ = _hx_tmp_2
		if options != nil {
			_hx_tmp_2 = options.Hx_Field_translations
		} else {
			_hx_tmp_2 = nil
		}

		this.Hx_Field_translations = _hx_tmp_2
		var _hx_tmp_3 Hx_Typedef_loreline_functionsmap
		_ = _hx_tmp_3
		if options != nil {
			_hx_tmp_3 = options.Hx_Field_functions
		} else {
			_hx_tmp_3 = nil
		}

		this.Hx_Field_hostFunctions = _hx_tmp_3
		var _hx_tmp_4 func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
		_ = _hx_tmp_4
		if options != nil {
			_hx_tmp_4 = options.Hx_Field_customCreateFields
		} else {
			_hx_tmp_4 = nil
		}

		this.Hx_Field_customCreateFields = _hx_tmp_4
	}

}

var Hx_Obj_loreline_interpretercontext_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.InterpreterContext",
	HxMakeArray[string](),
	HxMakeArray[string]("script", "lens", "translations", "stringLiteralProcessors", "strictAccess", "customCreateFields", "topLevelState", "topLevelCharacters", "topLevelBeats", "nodeStates", "random", "hostFunctions", "scriptFunctions", "nextInsertionId", "root", "children", "pendingRestoredFlows", "pendingRestoredKeys", "pendingSavedInsertions", "pendingInsertionsCache", "childByKey", "removePendingFlow", "clearPendingFlows"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_interpretercontext_CreateInstance(HxConvert[*Hx_Obj_loreline_script](params.Get(0)), HxConvert[*Hx_Obj_loreline_interpreteroptions](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_interpretercontext_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_interpretercontext_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_interpretercontext_RTTI
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_As_Hx_Obj_loreline_interpretercontext() *Hx_Obj_loreline_interpretercontext {
	return this
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_Field_childByKey(key string) *Hx_Obj_loreline_interpreter {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_interpreter]
		_ = _g1
		_g1 = this.Hx_Field_children
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var child *Hx_Obj_loreline_interpreter
			_ = child
			child = _g1.Get(_g)
			_g++
			if child.Hx_Field_key == key {
				return child
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_Field_removePendingFlow(key string) any {
	if this.Hx_Field_pendingRestoredFlows == nil {
		return nil
	}

	var flow any
	_ = flow
	flow = this.Hx_Field_pendingRestoredFlows.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(key)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(flow, nil))) {
		return nil
	}

	func() func(_hx_param_key string) bool {
		var _hx_closure_src_583 func(_hx_param_key any) bool
		_ = _hx_closure_src_583
		_hx_closure_src_583 = this.Hx_Field_pendingRestoredFlows.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_remove
		return func(_hx_param_key string) bool {
			return _hx_closure_src_583(((any)(_hx_param_key)))
		}
	}()(key)
	_ = (Hx_Array_Remove[string](this.Hx_Field_pendingRestoredKeys, key))
	if this.Hx_Field_pendingRestoredKeys.Len() == 0 {
		this.VTable.Hx_Field_clearPendingFlows()
	}

	return flow
}

func (this *Hx_Obj_loreline_interpretercontext) Hx_Field_clearPendingFlows() {
	this.Hx_Field_pendingRestoredFlows = nil
	this.Hx_Field_pendingRestoredKeys = nil
	this.Hx_Field_pendingSavedInsertions = nil
	this.Hx_Field_pendingInsertionsCache = nil
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_interpreter() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_interpreter() {
	Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH = Hx_Init_Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH()
}

type Hx_Obj_VTable_loreline_interpreter interface {
	Hx_Field_restoreChoiceOption(saved any, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_choiceoption
	Hx_Field_stripStringComments(content string) string
	Hx_Field_setStateField(name string, value any)
	Hx_Field_save() any
	Hx_Field_transitionToBeat(beat *Hx_Obj_loreline_nbeatdecl, argValues__isOptional any, capturedChain__isOptional any)
	Hx_Field_evalBeatCall(call *Hx_Obj_loreline_nbeatcall, next func())
	Hx_Field_customStrategyCandidate(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any], strategy any) any
	Hx_Field_isOriginalScriptExpression(expr *Hx_Obj_loreline_nexpr) bool
	Hx_Field_printLoreline(node *Hx_Obj_loreline_node) string
	Hx_Field_resumeSpawn(key string, handleDialogue__isOptional any, handleChoice__isOptional any, handleFinish__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter
	Hx_Field_resumeNode(node *Hx_Obj_loreline_astnode, scopeLevel int32, next func())
	Hx_Field_resumeNodeBody(node *Hx_Obj_loreline_astnode, scopeLevel int32, body HxArray[*Hx_Obj_loreline_astnode], next func())
	Hx_Field_resumeBeatCall(call *Hx_Obj_loreline_nbeatcall, scopeLevel int32, next func())
	Hx_Field_seedBeatParams(scope *Hx_Obj_loreline_runtimescope, beat *Hx_Obj_loreline_nbeatdecl, values HxArray[any])
	Hx_Field_readAccess(access Hx_Enum_loreline_runtimeaccess) any
	Hx_Field_serializeState(state *Hx_Obj_loreline_runtimestate) any
	Hx_Field_resumeCall(call *Hx_Obj_loreline_ncall, scopeLevel int32, next func())
	Hx_Field_evalNode(node *Hx_Obj_loreline_astnode, next func())
	Hx_Field_getBeatVisitCount(beat *Hx_Obj_loreline_nbeatdecl) int32
	Hx_Field_recordWhenRulePlayed(rule *Hx_Obj_loreline_nwhenrule)
	Hx_Field_isWhenRuleConsumed(rule *Hx_Obj_loreline_nwhenrule) bool
	Hx_Field_getWhenRuleLastPlayed(rule *Hx_Obj_loreline_nwhenrule) int32
	Hx_Field_setNodeStateField(node *Hx_Obj_loreline_astnode, field string, value any)
	Hx_Field_get_script() *Hx_Obj_loreline_script
	Hx_Field_get_translations() Hx_Typedef_map
	Hx_Field_set_translations(translations Hx_Typedef_map) Hx_Typedef_map
	Hx_Field_initializeState(state *Hx_Obj_loreline_nstatedecl, scope *Hx_Obj_loreline_runtimescope)
	Hx_Field_presentRestoredDialogue(node *Hx_Obj_loreline_astnode, next func()) bool
	Hx_Field_isChoiceOptionChosen(option *Hx_Obj_loreline_nchoiceoption) bool
	Hx_Field_setChoiceOptionChosen(option *Hx_Obj_loreline_nchoiceoption)
	Hx_Field_snapshotOriginalValue(value any) any
	Hx_Field_setTopLevelStateField(name string, value any)
	Hx_Field_get_currentScope() *Hx_Obj_loreline_runtimescope
	Hx_Field_collectWhenRules(when *Hx_Obj_loreline_nwhenstatement, result HxArray[any], startIndex int32, stopAtFirst bool, owner *Hx_Obj_loreline_runtimeinsertion, next func())
	Hx_Field_resolveBeatRefFromCall(call *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_runtimebeatref
	Hx_Field_evalCall(call *Hx_Obj_loreline_ncall, next func())
	Hx_Field_findHashCommentId(node *Hx_Obj_loreline_astnode, str *Hx_Obj_loreline_nstringliteral) HxNullable[string]
	Hx_Field_initializeTopLevelState(state *Hx_Obj_loreline_nstatedecl)
	Hx_Field_isBeatCall(node *Hx_Obj_loreline_astnode, scopeLevel__isOptional any) bool
	Hx_Field_evaluateString(str *Hx_Obj_loreline_nstringliteral) any
	Hx_Field_dispose()
	Hx_Field_resumeChoiceOption(option *Hx_Obj_loreline_nchoiceoption, scopeLevel int32, next func())
	Hx_Field_resumeWhen(when *Hx_Obj_loreline_nwhenstatement, scopeLevel int32, next func())
	Hx_Field_initializeTopLevelFunction(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl)
	Hx_Field_serializeChoiceEvalContext(flow *Hx_Obj_loreline_interpreter) HxArray[any]
	Hx_Field_setCharacterField(character string, name string, value any)
	Hx_Field_hostValueToRuntimeDepth(value any, depth int32) any
	Hx_Field_registerTopLevelFunction(key string, _hx_reserved_func any)
	Hx_Field_evaluateFunctionCall(call *Hx_Obj_loreline_ncall, next func()) any
	Hx_Field_get_stringLiteralProcessors() HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
	Hx_Field_seedRandom(seed__isOptional any)
	Hx_Field_eachResolutionScope(startIndex int32, visit func(scope *Hx_Obj_loreline_runtimescope) bool)
	Hx_Field_walkResolutionScopes(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32, visit func(scope *Hx_Obj_loreline_runtimescope) bool)
	Hx_Field_getStateField(name string) any
	Hx_Field_restoreStack(savedStack HxArray[any], savedInsertions__isOptional any, restoredInsertions__isOptional any) bool
	Hx_Field_evalText(text *Hx_Obj_loreline_ntextstatement, next func())
	Hx_Field_restoreShuffleOrder(saved HxArray[int32], item *Hx_Obj_loreline_nblock) HxArray[int32]
	Hx_Field_evaluateBeatArgs(beat *Hx_Obj_loreline_nbeatdecl, args HxArray[*Hx_Obj_loreline_nexpr], pos *Hx_Obj_loreline_position) HxArray[any]
	Hx_Field_evalIf(ifStmt *Hx_Obj_loreline_nifstatement, next func())
	Hx_Field_beatNotFoundError(name string, node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_runtimeerror
	Hx_Field_evaluateArrayLiteral(expr HxArray[any]) any
	Hx_Field_evaluateObjectLiteral(expr HxArray[*Hx_Obj_loreline_nobjectfield]) any
	Hx_Field_push(scope *Hx_Obj_loreline_runtimescope)
	Hx_Field_serializeWhenCandidate(candidate any, insertions any) any
	Hx_Field_isRoot() bool
	Hx_Field_get_topLevelState() *Hx_Obj_loreline_runtimestate
	Hx_Field_resumableSpawnKeys() HxArray[string]
	Hx_Field_getStateFieldRaw(name string) any
	Hx_Field_restoreInsertion(insertionId int32, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_runtimeinsertion
	Hx_Field_presentDialogue(node *Hx_Obj_loreline_astnode, character HxNullable[string], text string, tags HxArray[*Hx_Obj_loreline_texttag], next func())
	Hx_Field_insertionAtOrBelow(scopeLevel int32) *Hx_Obj_loreline_runtimeinsertion
	Hx_Field_inferImplicitCharactersFromRootBeat()
	Hx_Field_spawn(key string, handleDialogue__isOptional any, handleChoice__isOptional any, handleFinish__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter
	Hx_Field_serializeNodeStates() any
	Hx_Field_incrementBeatVisitCount(beat *Hx_Obj_loreline_nbeatdecl)
	Hx_Field_playWhenRule(rule *Hx_Obj_loreline_nwhenrule, next func())
	Hx_Field_playInsertedWhenRule(candidate any, next func())
	Hx_Field_getTranslatedString(node *Hx_Obj_loreline_astnode, str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral
	Hx_Field_restoreFlow(savedStack HxArray[any], savedPendingChoiceOptions HxArray[any], savedPendingDialogue any, savedChoiceEvalContext HxArray[any], savedInsertions any, restoredInsertions Hx_Typedef_map)
	Hx_Field_attachBeforeRun()
	Hx_Field_restoreFields(target any, savedFields any) any
	Hx_Field_presentChoice(choice *Hx_Obj_loreline_nchoicestatement, options HxArray[*Hx_Obj_loreline_choiceoption], next func())
	Hx_Field_transformPipeRaw(text string, countExpr *Hx_Obj_loreline_nexpr, pos *Hx_Obj_loreline_position, parts HxArray[*Hx_Obj_loreline_nstringpart])
	Hx_Field_resolvesExistingBinding(name string) bool
	Hx_Field_createChild(key string, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, options *Hx_Obj_loreline_interpreteroptions) *Hx_Obj_loreline_interpreter
	Hx_Field_currentNodeFilePath(rootPath string) string
	Hx_Field_serializeChoiceOption(option *Hx_Obj_loreline_choiceoption, insertions any) any
	Hx_Field_resumeBeatRun(beat *Hx_Obj_loreline_nbeatdecl, scopeLevel int32, next func())
	Hx_Field_evalBeatRun(beat *Hx_Obj_loreline_nbeatdecl, next func(), argValues__isOptional any, capturedChain__isOptional any)
	Hx_Field_evalStateDecl(state *Hx_Obj_loreline_nstatedecl, next func())
	Hx_Field_resolveBeatRefByName(name string) *Hx_Obj_loreline_runtimebeatref
	Hx_Field_checkSpawnKey(key string)
	Hx_Field_serializeCharacters() any
	Hx_Field_serializeFields(fields any, originalFields__isOptional any) any
	Hx_Field_restoreWhenCandidates(saved HxArray[any], savedInsertions any, restoredInsertions Hx_Typedef_map) HxArray[any]
	Hx_Field_pop() bool
	Hx_Field_evalChoiceOptionsAndInsertions(beat *Hx_Obj_loreline_nbeatdecl, choice *Hx_Obj_loreline_nchoicestatement, result HxArray[*Hx_Obj_loreline_choiceoption], next func(), startIndex__isOptional any, isChoiceRoot__isOptional any)
	Hx_Field_evaluateCondition(expr *Hx_Obj_loreline_nexpr) bool
	Hx_Field_resolveAssignmentTarget(target *Hx_Obj_loreline_nexpr) Hx_Enum_loreline_runtimeaccess
	Hx_Field_serializeScope(scope *Hx_Obj_loreline_runtimescope, insertions any) any
	Hx_Field_setAlternativeVisitCount(alt *Hx_Obj_loreline_nalternative, count int32)
	Hx_Field_evaluateDynamicCall(call *Hx_Obj_loreline_ncall, next func()) any
	Hx_Field_valueToString(value any) string
	Hx_Field_restore(saveData any)
	Hx_Field_get_currentInsertion() *Hx_Obj_loreline_runtimeinsertion
	Hx_Field_initializeTopLevelCharacter(character *Hx_Obj_loreline_ncharacterdecl)
	Hx_Field_evalDialogue(dialogue *Hx_Obj_loreline_ndialoguestatement, next func())
	Hx_Field_evalWhen(when *Hx_Obj_loreline_nwhenstatement, next func())
	Hx_Field_evalTransition(transition *Hx_Obj_loreline_ntransition)
	Hx_Field_get_nextInsertionId() int32
	Hx_Field_currentNode() *Hx_Obj_loreline_astnode
	Hx_Field_serializeTextTag(tag *Hx_Obj_loreline_texttag) any
	Hx_Field_start(beatName__isOptional any)
	Hx_Field_hostValueOf(value any) any
	Hx_Field_serializeCharacter(character *Hx_Obj_loreline_runtimecharacter) any
	Hx_Field_initializeTopLevelBeat(beat *Hx_Obj_loreline_nbeatdecl)
	Hx_Field_getCharacter(name string) any
	Hx_Field_getCharacterField(character string, name string) any
	Hx_Field_get_topLevelCharacters() Hx_Typedef_map
	Hx_Field_restoreBeat(beatRef any) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_restoreCharacter(character *Hx_Obj_loreline_runtimecharacter, data any) *Hx_Obj_loreline_runtimecharacter
	Hx_Field_bindFunctions(functions Hx_Typedef_loreline_functionsmap)
	Hx_Field_refreshFunctionHelpers()
	Hx_Field_evalChoice(choice *Hx_Obj_loreline_nchoicestatement, next func())
	Hx_Field_finish()
	Hx_Field_get_lens() *Hx_Obj_loreline_lens
	Hx_Field_serializePendingDialogue(flow *Hx_Obj_loreline_interpreter) any
	Hx_Field_restoreScopeItem(savedScope any, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_runtimescope
	Hx_Field_restoreNode(savedNode any, savedBeatId Hx_Typedef_haxe__int64___int64, beat *Hx_Obj_loreline_nbeatdecl) *Hx_Obj_loreline_astnode
	Hx_Field_flush()
	Hx_Field_chooseWhenCandidate(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any]) any
	Hx_Field_getAlternativeVisitCount(alt *Hx_Obj_loreline_nalternative) int32
	Hx_Field_hostValueToRuntime(value any) any
	Hx_Field_serializeValue(value any) any
	Hx_Field_resumeChoice(choice *Hx_Obj_loreline_nchoicestatement, scopeLevel int32, next func())
	Hx_Field_resumeAlternative(alt *Hx_Obj_loreline_nalternative, scopeLevel int32, next func())
	Hx_Field_restoreCharacters(data any)
	Hx_Field_wrapNext(cb func()) *Hx_Obj_loreline_evalnext
	Hx_Field_evalChoiceOption(option *Hx_Obj_loreline_nchoiceoption, next func())
	Hx_Field_resolveBeatByName(name string, scopeLevel__isOptional any) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_evalNodeBody(beat *Hx_Obj_loreline_nbeatdecl, node *Hx_Obj_loreline_astnode, body HxArray[*Hx_Obj_loreline_astnode], insertion__isOptional any, argValues__isOptional any, capturedChain__isOptional any, shuffle__isOptional any, next func())
	Hx_Field_evalInsertion(insertion *Hx_Obj_loreline_runtimeinsertion, next func())
	Hx_Field_collectingWhenInsertion(scopeLevel int32) *Hx_Obj_loreline_runtimeinsertion
	Hx_Field_resolveBeatRefFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_runtimebeatref
	Hx_Field_captureChainForBeat(beat *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_runtimescope]
	Hx_Field_resolveAccess(access *Hx_Obj_loreline_naccess, target__isOptional any, name string) Hx_Enum_loreline_runtimeaccess
	Hx_Field_getTypeName(t Hx_Enum_valuetype) string
	Hx_Field_getTopLevelStateField(name string) any
	Hx_Field_serializeNodeReference(node *Hx_Obj_loreline_astnode) any
	Hx_Field_restoreValue(value any) any
	Hx_Field_bindScriptFunction(_hx_reserved_func any)
	Hx_Field_resolveDynamicBeatTarget(value any, node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_runtimebeatref
	Hx_Field_rawTextHasUnescapedPipe(text string) bool
	Hx_Field_writeAccess(access Hx_Enum_loreline_runtimeaccess, value any)
	Hx_Field_resumeIf(ifStmt *Hx_Obj_loreline_nifstatement, scopeLevel int32, next func())
	Hx_Field_restoreBeatToResume(savedStack HxArray[any]) *Hx_Obj_loreline_nbeatdecl
	Hx_Field_get_nodeStates() *Hx_Obj_loreline_nodeidmap
	Hx_Field_set_nextInsertionId(nextInsertionId int32) int32
	Hx_Field_serializeInsertion(insertion *Hx_Obj_loreline_runtimeinsertion, insertions any) HxNullable[int32]
	Hx_Field_restoreNodeStates(data any)
	Hx_Field_restoredOptionsMatch(choice *Hx_Obj_loreline_nchoicestatement, options HxArray[*Hx_Obj_loreline_choiceoption]) bool
	Hx_Field_evalAlternative(alt *Hx_Obj_loreline_nalternative, next func())
	Hx_Field_evaluateExpression(expr *Hx_Obj_loreline_nexpr) any
	Hx_Field_restoreState(state *Hx_Obj_loreline_runtimestate, data any) *Hx_Obj_loreline_runtimestate
	Hx_Field_evalBeatDecl(beat *Hx_Obj_loreline_nbeatdecl, next func())
	Hx_Field_get_topLevelBeats() Hx_Typedef_map
	Hx_Field_serializePendingChoiceOptions(flow *Hx_Obj_loreline_interpreter, insertions any) HxArray[any]
	Hx_Field_serializeBeatReference(beat *Hx_Obj_loreline_nbeatdecl) any
	Hx_Field_getNodeStateInt(node *Hx_Obj_loreline_astnode, field string) int32
	Hx_Field_makePluralCall(countExpr *Hx_Obj_loreline_nexpr, singular string, pluralForm string, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_ncall
	Hx_Field_set_stringLiteralProcessors(processors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]) HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
	Hx_Field_get_strictAccess() bool
	Hx_Field_resumeFromLevel(scopeLevel int32, next func())
	Hx_Field_initializeStringLiteralProcessors(processors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral])
	Hx_Field_finishWhenCollection(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any], owner *Hx_Obj_loreline_runtimeinsertion, next func())
	Hx_Field_isDynamicCallForm(call *Hx_Obj_loreline_ncall) bool
	Hx_Field_pluralPipeProcess(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral
	Hx_Field_stripStringIndent(content string) string
	Hx_Field_removeCurrentInsertion()
	Hx_Field_evalAssignment(assign *Hx_Obj_loreline_nassign, next func())
	Hx_Field_callFunctionValue(fn any, args HxArray[any], pos *Hx_Obj_loreline_position, next func()) any
	Hx_Field_performOperation(op Hx_Enum_loreline_tokentype, left any, right any, pos *Hx_Obj_loreline_position) any
	Hx_Field_valueToStringImpl(value any, seen HxArray[any]) string
	Hx_Field_resume()
	Hx_Field_evalShuffledItems(alt *Hx_Obj_loreline_nalternative, indices HxArray[int32], idx int32, next func())
	Hx_Field_mostSalientCandidate(eligible HxArray[any]) any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_interpreter struct {
	VTable                           Hx_Obj_VTable_loreline_interpreter `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_context                 *Hx_Obj_loreline_interpretercontext
	Hx_Field_key                     string
	Hx_Field_disposed                bool
	Hx_Field_finished                bool
	Hx_Field_handleDialogue          Hx_Typedef_loreline_dialoguehandler
	Hx_Field_handleChoice            Hx_Typedef_loreline_choicehandler
	Hx_Field_handleFinish            Hx_Typedef_loreline_finishhandler
	Hx_Field_topLevelFunctions       Hx_Typedef_map
	Hx_Field_stringHelpers           Hx_Typedef_map
	Hx_Field_arrayHelpers            Hx_Typedef_map
	Hx_Field_mapHelpers              Hx_Typedef_map
	Hx_Field_beatHelpers             Hx_Typedef_map
	Hx_Field_builtins                *Hx_Obj_loreline_functions
	Hx_Field_stack                   HxArray[*Hx_Obj_loreline_runtimescope]
	Hx_Field_nextScopeId             int32
	Hx_Field_syncCallbacks           HxArray[func()]
	Hx_Field_flushing                bool
	Hx_Field_finishTrigger           *Hx_Obj_loreline_evalnext
	Hx_Field_beatToResume            *Hx_Obj_loreline_nbeatdecl
	Hx_Field_pendingChoiceOptions    HxArray[*Hx_Obj_loreline_choiceoption]
	Hx_Field_pendingDialogue         any
	Hx_Field__choiceEvalTexts        HxArray[string]
	Hx_Field__choiceEvalEnabled      HxArray[bool]
	Hx_Field_customCreateFields      func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
	Hx_Field_wrapper                 any
	Hx_Field__saveInsertions         any
	Hx_Field__serializingScopes      HxArray[*Hx_Obj_loreline_runtimescope]
	Hx_Field__restoreSavedInsertions any
	Hx_Field__restoreInsertionsCache Hx_Typedef_map
}

func Hx_Obj_loreline_interpreter_CreateEmptyInstance() *Hx_Obj_loreline_interpreter {
	obj := &Hx_Obj_loreline_interpreter{}
	obj.VTable = obj
	obj.Hx_Field_context = HxDefault[*Hx_Obj_loreline_interpretercontext]()
	obj.Hx_Field_key = HxDefault[string]()
	obj.Hx_Field_disposed = HxDefault[bool]()
	obj.Hx_Field_finished = HxDefault[bool]()
	obj.Hx_Field_handleDialogue = HxDefault[Hx_Typedef_loreline_dialoguehandler]()
	obj.Hx_Field_handleChoice = HxDefault[Hx_Typedef_loreline_choicehandler]()
	obj.Hx_Field_handleFinish = HxDefault[Hx_Typedef_loreline_finishhandler]()
	obj.Hx_Field_topLevelFunctions = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_stringHelpers = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_arrayHelpers = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_mapHelpers = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_beatHelpers = HxDefault[Hx_Typedef_map]()
	obj.Hx_Field_builtins = HxDefault[*Hx_Obj_loreline_functions]()
	obj.Hx_Field_stack = HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	obj.Hx_Field_nextScopeId = HxDefault[int32]()
	obj.Hx_Field_syncCallbacks = HxDefault[HxArray[func()]]()
	obj.Hx_Field_flushing = HxDefault[bool]()
	obj.Hx_Field_finishTrigger = HxDefault[*Hx_Obj_loreline_evalnext]()
	obj.Hx_Field_beatToResume = HxDefault[*Hx_Obj_loreline_nbeatdecl]()
	obj.Hx_Field_pendingChoiceOptions = HxDefault[HxArray[*Hx_Obj_loreline_choiceoption]]()
	obj.Hx_Field_pendingDialogue = HxDefault[any]()
	obj.Hx_Field__choiceEvalTexts = HxDefault[HxArray[string]]()
	obj.Hx_Field__choiceEvalEnabled = HxDefault[HxArray[bool]]()
	obj.Hx_Field_customCreateFields = HxDefault[func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any]()
	obj.Hx_Field_wrapper = HxDefault[any]()
	obj.Hx_Field__saveInsertions = HxDefault[any]()
	obj.Hx_Field__serializingScopes = HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	obj.Hx_Field__restoreSavedInsertions = HxDefault[any]()
	obj.Hx_Field__restoreInsertionsCache = HxDefault[Hx_Typedef_map]()
	return obj
}

func Hx_Obj_loreline_interpreter_CreateInstance(script *Hx_Obj_loreline_script, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, options__isOptional any, parentContext__isOptional any) *Hx_Obj_loreline_interpreter {
	obj := Hx_Obj_loreline_interpreter_CreateEmptyInstance()
	obj.Hx_New(script, handleDialogue, handleChoice, handleFinish, options__isOptional, parentContext__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_interpreter) Hx_New(script *Hx_Obj_loreline_script, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, options__isOptional any, parentContext__isOptional any) {
	{
		options := HxDefault[*Hx_Obj_loreline_interpreteroptions]()
		_ = options

		if options__isOptional != nil {
			options = HxConvert[*Hx_Obj_loreline_interpreteroptions](options__isOptional)
		} else {
			options = nil
		}
		parentContext := HxDefault[*Hx_Obj_loreline_interpretercontext]()
		_ = parentContext

		if parentContext__isOptional != nil {
			parentContext = HxConvert[*Hx_Obj_loreline_interpretercontext](parentContext__isOptional)
		} else {
			parentContext = nil
		}
		{
			this.Hx_Field_key = HxStringNull
		}
		{
			this.Hx_Field_disposed = false
		}
		{
			this.Hx_Field_finished = false
		}
		{
			this.Hx_Field_topLevelFunctions = (func() *Hx_Obj_haxe_imap {
				_hx_icast_212 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_212 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_212, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_212.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_stringHelpers = (func() *Hx_Obj_haxe_imap {
				_hx_icast_213 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_213 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_213, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_213.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_arrayHelpers = (func() *Hx_Obj_haxe_imap {
				_hx_icast_214 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_214 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_214, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_214.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_mapHelpers = (func() *Hx_Obj_haxe_imap {
				_hx_icast_215 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_215 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_215, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_215.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_beatHelpers = (func() *Hx_Obj_haxe_imap {
				_hx_icast_216 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
				if _hx_icast_216 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_216, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_216.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}
		{
			this.Hx_Field_stack = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		}
		{
			this.Hx_Field_nextScopeId = 1
		}
		{
			this.Hx_Field_syncCallbacks = HxMakeArray[func()]()
		}
		{
			this.Hx_Field_flushing = false
		}
		{
			this.Hx_Field_finishTrigger = nil
		}
		{
			this.Hx_Field_beatToResume = nil
		}
		{
			this.Hx_Field_pendingChoiceOptions = nil
		}
		{
			this.Hx_Field_pendingDialogue = nil
		}
		{
			this.Hx_Field__choiceEvalTexts = HxMakeArray[string]()
		}
		{
			this.Hx_Field__choiceEvalEnabled = HxMakeArray[bool]()
		}
		{
			this.Hx_Field_wrapper = nil
		}
		{
			this.Hx_Field__saveInsertions = nil
		}
		{
			this.Hx_Field__serializingScopes = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		}
		{
			this.Hx_Field__restoreSavedInsertions = nil
		}
		{
			this.Hx_Field__restoreInsertionsCache = nil
		}
		{
			this.Hx_Field_handleDialogue = handleDialogue
			this.Hx_Field_handleChoice = handleChoice
			this.Hx_Field_handleFinish = handleFinish
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			if options != nil {
				_hx_tmp_0 = options.Hx_Field_wrapper
			} else {
				_hx_tmp_0 = nil
			}

			this.Hx_Field_wrapper = _hx_tmp_0
			if parentContext != nil {
				this.Hx_Field_context = parentContext
				var _hx_tmp_1 func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
				_ = _hx_tmp_1
				if options != nil {
					_hx_tmp_1 = options.Hx_Field_customCreateFields
				} else {
					_hx_tmp_1 = nil
				}

				var tmp func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
				_ = tmp
				tmp = _hx_tmp_1
				var _hx_tmp_2 func(interpreter *Hx_Obj_loreline_interpreter, _hx_reserved_type string, node *Hx_Obj_loreline_node) any
				_ = _hx_tmp_2
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(tmp)), ((any)(nil))))) {
					_hx_tmp_2 = tmp
				} else {
					_hx_tmp_2 = this.Hx_Field_context.Hx_Field_customCreateFields
				}

				this.Hx_Field_customCreateFields = _hx_tmp_2
				var _hx_tmp_3 Hx_Typedef_loreline_functionsmap
				_ = _hx_tmp_3
				if options != nil {
					_hx_tmp_3 = options.Hx_Field_functions
				} else {
					_hx_tmp_3 = nil
				}

				var functions Hx_Typedef_loreline_functionsmap
				_ = functions
				functions = _hx_tmp_3
				var _hx_tmp_4 Hx_Typedef_loreline_functionsmap
				_ = _hx_tmp_4
				if functions != nil {
					_hx_tmp_4 = functions
				} else {
					_hx_tmp_4 = this.Hx_Field_context.Hx_Field_hostFunctions
				}

				this.VTable.Hx_Field_bindFunctions(_hx_tmp_4)
				return
			}

			this.Hx_Field_context = Hx_Obj_loreline_interpretercontext_CreateInstance(script, options)
			this.Hx_Field_context.Hx_Field_root = this
			this.Hx_Field_customCreateFields = this.Hx_Field_context.Hx_Field_customCreateFields
			this.Hx_Field_context.Hx_Field_topLevelState = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
				_hx_up := script
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})(), nil, nil)
			this.VTable.Hx_Field_bindFunctions(this.Hx_Field_context.Hx_Field_hostFunctions)
			var _hx_tmp_5 HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
			_ = _hx_tmp_5
			if options != nil {
				_hx_tmp_5 = options.Hx_Field_stringLiteralProcessors
			} else {
				_hx_tmp_5 = nil
			}

			this.VTable.Hx_Field_initializeStringLiteralProcessors(_hx_tmp_5)
			{
				var _gdecl any
				_ = _gdecl
				_gdecl = ((any)(script.VTable.Hx_Field_iterator()))
				for {
					var _hx_tmp_6 any
					_ = _hx_tmp_6
					_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(_gdecl, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(_gdecl, "next")
					var hx_dyncast_282 any
					_ = hx_dyncast_282
					hx_dyncast_282 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_7, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_astnode")
					var _hx_tmp_8 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_8
					if hx_dyncast_282 != nil {
						_hx_tmp_8 = hx_dyncast_282.(interface {
							Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
						}).Hx_As_Hx_Obj_loreline_astnode()
					} else {
						_hx_tmp_8 = nil
					}

					var decl *Hx_Obj_loreline_astnode
					_ = decl
					decl = _hx_tmp_8
					{
						var _g *Hx_Obj_go_haxe_hxclass
						_ = _g
						_g = func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
							var _hx_closure_src_687 func(o any) *Hx_Obj_go_haxe_hxclass
							_ = _hx_closure_src_687
							_hx_closure_src_687 = Hx_Field_type_getClass
							return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
								return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_687(((any)(o)))))
							}
						}()(decl)
						if _g != nil {
							switch _g {
							case Hx_Obj_loreline_nbeatdecl_RTTI:
								{
									this.VTable.Hx_Field_initializeTopLevelBeat(decl.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
									}).Hx_As_Hx_Obj_loreline_nbeatdecl())
								}

							case Hx_Obj_loreline_ncharacterdecl_RTTI:
								{
									this.VTable.Hx_Field_initializeTopLevelCharacter(decl.VTable.(interface {
										Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl
									}).Hx_As_Hx_Obj_loreline_ncharacterdecl())
								}

							case Hx_Obj_loreline_nfunctiondecl_RTTI:
								{
									this.VTable.Hx_Field_initializeTopLevelFunction(decl.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl
									}).Hx_As_Hx_Obj_loreline_nfunctiondecl())
								}

							case Hx_Obj_loreline_nstatedecl_RTTI:
								{
									this.VTable.Hx_Field_initializeTopLevelState(decl.VTable.(interface {
										Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
									}).Hx_As_Hx_Obj_loreline_nstatedecl())
								}

							default:
								{
								}
							}
						}
					}
				}
			}

			this.VTable.Hx_Field_inferImplicitCharactersFromRootBeat()
		}

	}
}

var Hx_Obj_loreline_interpreter_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Interpreter",
	HxMakeArray[string]("HOST_VALUE_MAX_DEPTH", "hasEligibleCandidate", "debug"),
	HxMakeArray[string]("context", "key", "isRoot", "disposed", "finished", "get_script", "get_translations", "set_translations", "handleDialogue", "handleChoice", "handleFinish", "get_topLevelState", "get_topLevelCharacters", "get_topLevelBeats", "get_nodeStates", "topLevelFunctions", "stringHelpers", "arrayHelpers", "mapHelpers", "beatHelpers", "builtins", "get_stringLiteralProcessors", "set_stringLiteralProcessors", "stack", "get_lens", "get_strictAccess", "get_currentScope", "get_currentInsertion", "insertionAtOrBelow", "removeCurrentInsertion", "nextScopeId", "get_nextInsertionId", "set_nextInsertionId", "syncCallbacks", "flushing", "finishTrigger", "beatToResume", "pendingChoiceOptions", "pendingDialogue", "_choiceEvalTexts", "_choiceEvalEnabled", "customCreateFields", "wrapper", "inferImplicitCharactersFromRootBeat", "start", "_saveInsertions", "_serializingScopes", "_restoreSavedInsertions", "_restoreInsertionsCache", "seedRandom", "save", "serializePendingChoiceOptions", "serializePendingDialogue", "serializeChoiceEvalContext", "restore", "restoreFlow", "resume", "resumeFromLevel", "spawn", "resumeSpawn", "resumableSpawnKeys", "dispose", "checkSpawnKey", "createChild", "attachBeforeRun", "getCharacter", "getCharacterField", "setCharacterField", "eachResolutionScope", "walkResolutionScopes", "hostValueOf", "hostValueToRuntime", "hostValueToRuntimeDepth", "getStateField", "getStateFieldRaw", "setStateField", "getTopLevelStateField", "setTopLevelStateField", "currentNode", "currentNodeFilePath", "serializeScope", "serializeInsertion", "serializeWhenCandidate", "serializeChoiceOption", "serializeTextTag", "serializeState", "serializeCharacters", "serializeNodeStates", "serializeBeatReference", "serializeNodeReference", "serializeCharacter", "serializeFields", "serializeValue", "resumeNode", "resumeNodeBody", "resumeBeatRun", "resumeChoiceOption", "resumeChoice", "resumeIf", "resumeWhen", "resumeAlternative", "resumeCall", "resumeBeatCall", "restoreStack", "restoreScopeItem", "restoreInsertion", "restoreWhenCandidates", "restoreChoiceOption", "restoreBeatToResume", "restoreNode", "restoreBeat", "restoreState", "restoreCharacters", "restoreNodeStates", "restoreCharacter", "restoreFields", "restoreValue", "bindFunctions", "refreshFunctionHelpers", "bindScriptFunction", "registerTopLevelFunction", "initializeStringLiteralProcessors", "initializeTopLevelFunction", "wrapNext", "flush", "pop", "push", "initializeTopLevelState", "initializeTopLevelBeat", "initializeTopLevelCharacter", "initializeState", "evaluateBeatArgs", "seedBeatParams", "finish", "transitionToBeat", "evalNode", "evalBeatDecl", "evalNodeBody", "evalBeatRun", "evalStateDecl", "evalText", "evalDialogue", "presentDialogue", "presentRestoredDialogue", "evalChoice", "restoredOptionsMatch", "presentChoice", "evalChoiceOptionsAndInsertions", "evalInsertion", "evalBeatCall", "evalChoiceOption", "evalIf", "evalAlternative", "evalShuffledItems", "restoreShuffleOrder", "getBeatVisitCount", "incrementBeatVisitCount", "evalWhen", "collectingWhenInsertion", "finishWhenCollection", "collectWhenRules", "chooseWhenCandidate", "mostSalientCandidate", "customStrategyCandidate", "playWhenRule", "recordWhenRulePlayed", "playInsertedWhenRule", "isWhenRuleConsumed", "getWhenRuleLastPlayed", "getNodeStateInt", "setNodeStateField", "getAlternativeVisitCount", "setAlternativeVisitCount", "isChoiceOptionChosen", "setChoiceOptionChosen", "evalAssignment", "snapshotOriginalValue", "isOriginalScriptExpression", "isBeatCall", "isDynamicCallForm", "resolveBeatByName", "resolveBeatRefByName", "resolveBeatRefFromNode", "captureChainForBeat", "resolveDynamicBeatTarget", "beatNotFoundError", "resolveBeatRefFromCall", "evalCall", "evalTransition", "getTranslatedString", "findHashCommentId", "evaluateString", "pluralPipeProcess", "rawTextHasUnescapedPipe", "transformPipeRaw", "makePluralCall", "stripStringIndent", "stripStringComments", "evaluateCondition", "callFunctionValue", "resolvesExistingBinding", "evaluateDynamicCall", "evaluateFunctionCall", "evaluateArrayLiteral", "evaluateObjectLiteral", "evaluateExpression", "readAccess", "writeAccess", "resolveAssignmentTarget", "resolveAccess", "getTypeName", "performOperation", "valueToString", "valueToStringImpl", "printLoreline"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_interpreter_CreateInstance(HxConvert[*Hx_Obj_loreline_script](params.Get(0)), HxConvert[Hx_Typedef_loreline_dialoguehandler](params.Get(1)), HxConvert[Hx_Typedef_loreline_choicehandler](params.Get(2)), HxConvert[Hx_Typedef_loreline_finishhandler](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_interpreter_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_interpreter_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "HOST_VALUE_MAX_DEPTH":
			return &Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH
		case "hasEligibleCandidate":
			return Hx_Field_loreline_interpreter_hasEligibleCandidate
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_interpreter_RTTI
}

func (this *Hx_Obj_loreline_interpreter) Hx_As_Hx_Obj_loreline_interpreter() *Hx_Obj_loreline_interpreter {
	return this
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isRoot() bool {
	return (this.Hx_Field_context.Hx_Field_root == this)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_script() *Hx_Obj_loreline_script {
	return this.Hx_Field_context.Hx_Field_script
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_translations() Hx_Typedef_map {
	return this.Hx_Field_context.Hx_Field_translations
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_set_translations(translations Hx_Typedef_map) Hx_Typedef_map {
	this.Hx_Field_context.Hx_Field_translations = translations
	var _hx_tmp_0 Hx_Typedef_map
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_context.Hx_Field_translations
	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_topLevelState() *Hx_Obj_loreline_runtimestate {
	return this.Hx_Field_context.Hx_Field_topLevelState
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_topLevelCharacters() Hx_Typedef_map {
	return this.Hx_Field_context.Hx_Field_topLevelCharacters
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_topLevelBeats() Hx_Typedef_map {
	return this.Hx_Field_context.Hx_Field_topLevelBeats
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_nodeStates() *Hx_Obj_loreline_nodeidmap {
	return this.Hx_Field_context.Hx_Field_nodeStates
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_stringLiteralProcessors() HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral] {
	return this.Hx_Field_context.Hx_Field_stringLiteralProcessors
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_set_stringLiteralProcessors(processors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]) HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral] {
	this.Hx_Field_context.Hx_Field_stringLiteralProcessors = processors
	var _hx_tmp_0 HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_context.Hx_Field_stringLiteralProcessors
	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_lens() *Hx_Obj_loreline_lens {
	return this.Hx_Field_context.Hx_Field_lens
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_strictAccess() bool {
	return this.Hx_Field_context.Hx_Field_strictAccess
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_currentScope() *Hx_Obj_loreline_runtimescope {
	if this.Hx_Field_stack.Len() > 0 {
		var _hx_tmp_0 HxArray[*Hx_Obj_loreline_runtimescope]
		_ = _hx_tmp_0
		_hx_tmp_0 = this.Hx_Field_stack
		return _hx_tmp_0.Get((this.Hx_Field_stack.Len() - 1))
	} else {
		return nil
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_currentInsertion() *Hx_Obj_loreline_runtimeinsertion {
	return this.VTable.Hx_Field_insertionAtOrBelow((this.Hx_Field_stack.Len() - 1))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_insertionAtOrBelow(scopeLevel int32) *Hx_Obj_loreline_runtimeinsertion {
	var i int32
	_ = i
	i = scopeLevel
	for i >= 0 {
		var scope *Hx_Obj_loreline_runtimescope
		_ = scope
		scope = this.Hx_Field_stack.Get(i)
		if scope.Hx_Field_insertion != nil {
			return scope.Hx_Field_insertion
		}

		i--
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_removeCurrentInsertion() {
	var i int32
	_ = i
	i = (this.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		var scope *Hx_Obj_loreline_runtimescope
		_ = scope
		scope = this.Hx_Field_stack.Get(i)
		if scope.Hx_Field_insertion != nil {
			scope.Hx_Field_insertion = nil
			break
		}

		i--
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_get_nextInsertionId() int32 {
	return this.Hx_Field_context.Hx_Field_nextInsertionId
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_set_nextInsertionId(nextInsertionId int32) int32 {
	this.Hx_Field_context.Hx_Field_nextInsertionId = nextInsertionId
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_context.Hx_Field_nextInsertionId
	return _hx_tmp_0
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_inferImplicitCharactersFromRootBeat() {
	var _g int32
	_ = _g
	_g = 0
	var _hx_tmp_2 *Hx_Obj_go_haxe_hxclass
	_ = _hx_tmp_2
	_hx_tmp_2 = ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI))
	var hx_nullable_f208 HxNullable[bool]
	_ = hx_nullable_f208
	hx_nullable_f208 = HxNullable[bool]{}

	var _hx_tmp_3 any
	_ = _hx_tmp_3
	if (hx_nullable_f208).Valid == true {
		_hx_tmp_3 = (hx_nullable_f208).Value
	} else {
		_hx_tmp_3 = nil
	}

	var _g1 HxArray[*Hx_Obj_loreline_nbeatdecl]
	_ = _g1
	_g1 = func() func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nbeatdecl] {
		var _hx_closure_src_584 func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any]
		_ = _hx_closure_src_584
		_hx_closure_src_584 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodesOfType
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

				return HxMakeArrayView[*Hx_Obj_loreline_nbeatdecl](_hx_closure_src_584(_hx_tmp_0, _hx_tmp_1))
			}
		}
	}()(_hx_tmp_2, _hx_tmp_3)
	for {
		var _hx_tmp_4 int32
		_ = _hx_tmp_4
		_hx_tmp_4 = _g
		if !(_hx_tmp_4 < _g1.Len()) {
			break
		}

		var beat *Hx_Obj_loreline_nbeatdecl
		_ = beat
		beat = _g1.Get(_g)
		_g++
		if beat.Hx_Field_name != "_" {
			continue
		}

		{
			var _g2 int32
			_ = _g2
			_g2 = 0
			var _g3 HxArray[string]
			_ = _g3
			_g3 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getDialogueSpeakers((func() *Hx_Obj_loreline_node {
				_hx_up := beat
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g2
				if !(_hx_tmp_5 < _g3.Len()) {
					break
				}

				var name string
				_ = name
				name = _g3.Get(_g2)
				_g2++
				if func() func(_hx_param_key string) bool {
					var _hx_closure_src_585 func(_hx_param_key any) bool
					_ = _hx_closure_src_585
					_hx_closure_src_585 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_585(((any)(_hx_param_key)))
					}
				}()(name) {
					continue
				}

				if Hx_Field_loreline_objects_fieldExists(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name) {
					continue
				}

				this.VTable.Hx_Field_initializeTopLevelCharacter(Hx_Obj_loreline_ncharacterdecl_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, beat.Hx_Field_pos, name, beat.Hx_Field_pos, HxMakeArray[*Hx_Obj_loreline_nobjectfield](), nil, nil))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_start(beatName__isOptional any) {
	beatName := HxDefault[HxNullable[string]]()
	_ = beatName

	if beatName__isOptional != nil {
		beatName = HxConvert[HxNullable[string]](beatName__isOptional)
	} else {

	}
	{
		if (beatName).Valid == false {
			var hx_nullable_t58 HxNullable[string]
			_ = hx_nullable_t58
			hx_nullable_t58 = HxNullable[string]{Value: HxStringNull}

			beatName = HxNullable[string]{Value: (hx_nullable_t58).Value, Valid: (hx_nullable_t58).Valid}
		}

		this.VTable.Hx_Field_attachBeforeRun()
		var resolvedBeat *Hx_Obj_loreline_nbeatdecl
		_ = resolvedBeat
		if (beatName).Valid != false {
			resolvedBeat = func() func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
				var _hx_closure_src_586 func(_hx_param_key any) any
				_ = _hx_closure_src_586
				_hx_closure_src_586 = this.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
					var hx_dyncast_218 any
					_ = hx_dyncast_218
					hx_dyncast_218 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_586(((any)(_hx_param_key))), "Hx_Obj_loreline_nbeatdecl")
					var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
					_ = _hx_tmp_0
					if hx_dyncast_218 != nil {
						_hx_tmp_0 = hx_dyncast_218.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					} else {
						_hx_tmp_0 = nil
					}

					return _hx_tmp_0
				}
			}()((beatName).Value)
			if resolvedBeat == nil {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Beat ", (beatName).Value), " not found"), this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
			}
		} else {
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_astnode]
				_ = _g1
				_g1 = this.Hx_Field_context.Hx_Field_script.Hx_Field_body
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
					if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nbeatdecl_RTTI) {
						var beat *Hx_Obj_loreline_nbeatdecl
						_ = beat
						beat = node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl()
						if beat.Hx_Field_name == "_" {
							resolvedBeat = beat
							break
						}

						if resolvedBeat == nil {
							resolvedBeat = beat
						}
					}
				}
			}

			if resolvedBeat == nil {
				var _gdecl any
				_ = _gdecl
				_gdecl = ((any)(this.Hx_Field_context.Hx_Field_script.VTable.Hx_Field_iterator()))
				for {
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gdecl, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_3 any
					_ = _hx_tmp_3
					_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(_gdecl, "next")
					var hx_dyncast_217 any
					_ = hx_dyncast_217
					hx_dyncast_217 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()), "Hx_Obj_loreline_astnode")
					var _hx_tmp_4 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_4
					if hx_dyncast_217 != nil {
						_hx_tmp_4 = hx_dyncast_217.(interface {
							Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
						}).Hx_As_Hx_Obj_loreline_astnode()
					} else {
						_hx_tmp_4 = nil
					}

					var decl *Hx_Obj_loreline_astnode
					_ = decl
					decl = _hx_tmp_4
					if Hx_Field_std_isOfType(decl, Hx_Obj_loreline_nbeatdecl_RTTI) {
						var beat *Hx_Obj_loreline_nbeatdecl
						_ = beat
						beat = decl.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl()
						if beat.Hx_Field_name == "_" {
							resolvedBeat = beat
							break
						}

						if resolvedBeat == nil {
							resolvedBeat = beat
						}
					}
				}
			}

			if resolvedBeat == nil {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("No beats found in script", this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
			}
		}

		var _hx_tmp_5 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_5
		_hx_tmp_5 = resolvedBeat
		var _hx_tmp_6 any
		_ = _hx_tmp_6
		_hx_tmp_6 = nil
		this.VTable.Hx_Field_transitionToBeat(_hx_tmp_5, _hx_tmp_6, nil)
		this.VTable.Hx_Field_flush()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_seedRandom(seed__isOptional any) {
	seed := HxDefault[HxNullable[float64]]()
	_ = seed

	if seed__isOptional != nil {
		seed = HxConvert[HxNullable[float64]](seed__isOptional)
	} else {

	}
	{
		if (seed).Valid == false {
			var hx_nullable_t59 HxNullable[float64]
			_ = hx_nullable_t59
			hx_nullable_t59 = HxNullable[float64]{}

			seed = HxNullable[float64]{Value: (hx_nullable_t59).Value, Valid: (hx_nullable_t59).Valid}
		}

		this.Hx_Field_context.Hx_Field_random = Hx_Obj_loreline_random_CreateInstance(seed)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_save() any {
	if !this.VTable.Hx_Field_isRoot() {
		return this.Hx_Field_context.Hx_Field_root.VTable.Hx_Field_save()
	}

	var insertions any
	_ = insertions
	insertions = any(map[string]any{})
	this.Hx_Field__saveInsertions = insertions
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_runtimescope]
		_ = _g2
		_g2 = this.Hx_Field_stack
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var scope *Hx_Obj_loreline_runtimescope
			_ = scope
			scope = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_serializeScope(scope, insertions)))
		}
	}

	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = ((any)(_g))
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = this.VTable.Hx_Field_serializeState(this.Hx_Field_context.Hx_Field_topLevelState)
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = this.VTable.Hx_Field_serializeCharacters()
	var result any
	_ = result
	result = any(map[string]any{"version": ((int32)(1)), "stack": _hx_tmp_2, "state": _hx_tmp_3, "characters": _hx_tmp_4, "nodeStates": this.VTable.Hx_Field_serializeNodeStates()})
	if this.Hx_Field_context.Hx_Field_random != nil {
		Hx_Field_go_haxe_hxdynamic_setField(result, "random", this.Hx_Field_context.Hx_Field_random.Hx_Field_seed)
	}

	var pendingOptions HxArray[any]
	_ = pendingOptions
	pendingOptions = this.VTable.Hx_Field_serializePendingChoiceOptions(this, insertions)
	if pendingOptions != nil {
		Hx_Field_go_haxe_hxdynamic_setField(result, "pendingChoiceOptions", pendingOptions)
	}

	var pendingDialogueData any
	_ = pendingDialogueData
	pendingDialogueData = this.VTable.Hx_Field_serializePendingDialogue(this)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(pendingDialogueData, nil))) {
		Hx_Field_go_haxe_hxdynamic_setField(result, "pendingDialogue", pendingDialogueData)
	}

	var choiceEvalContext HxArray[any]
	_ = choiceEvalContext
	choiceEvalContext = this.VTable.Hx_Field_serializeChoiceEvalContext(this)
	if choiceEvalContext != nil {
		Hx_Field_go_haxe_hxdynamic_setField(result, "choiceEvalContext", choiceEvalContext)
	}

	var children HxArray[any]
	_ = children
	children = HxMakeArray[any]().Dyn()
	{
		var _g3 int32
		_ = _g3
		_g3 = 0
		var _g4 HxArray[*Hx_Obj_loreline_interpreter]
		_ = _g4
		_g4 = this.Hx_Field_context.Hx_Field_children
		for {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g3
			if !(_hx_tmp_5 < _g4.Len()) {
				break
			}

			var child *Hx_Obj_loreline_interpreter
			_ = child
			child = _g4.Get(_g3)
			_g3++
			if child.Hx_Field_stack.Len() == 0 {
				continue
			}

			var child1 string
			_ = child1
			child1 = child.Hx_Field_key
			var tmp__g_1 HxArray[any]
			_ = tmp__g_1
			tmp__g_1 = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _g2
				_g2 = child.Hx_Field_stack
				for {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _g1
					if !(_hx_tmp_6 < _g2.Len()) {
						break
					}

					var scope *Hx_Obj_loreline_runtimescope
					_ = scope
					scope = _g2.Get(_g1)
					_g1++
					var _hx_tmp_7 HxArray[any]
					_ = _hx_tmp_7
					_hx_tmp_7 = HxMakeArrayView[any](tmp__g_1)
					_ = (Hx_Array_Push[any](_hx_tmp_7, this.VTable.Hx_Field_serializeScope(scope, insertions)))
				}
			}

			var flow any
			_ = flow
			flow = any(map[string]any{"key": ((any)(child1)), "stack": ((any)(tmp__g_1))})
			var childPendingOptions HxArray[any]
			_ = childPendingOptions
			childPendingOptions = this.VTable.Hx_Field_serializePendingChoiceOptions(child, insertions)
			if childPendingOptions != nil {
				Hx_Field_go_haxe_hxdynamic_setField(flow, "pendingChoiceOptions", childPendingOptions)
			}

			var childPendingDialogue any
			_ = childPendingDialogue
			childPendingDialogue = this.VTable.Hx_Field_serializePendingDialogue(child)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(childPendingDialogue, nil))) {
				Hx_Field_go_haxe_hxdynamic_setField(flow, "pendingDialogue", childPendingDialogue)
			}

			var childChoiceEvalContext HxArray[any]
			_ = childChoiceEvalContext
			childChoiceEvalContext = this.VTable.Hx_Field_serializeChoiceEvalContext(child)
			if childChoiceEvalContext != nil {
				Hx_Field_go_haxe_hxdynamic_setField(flow, "choiceEvalContext", childChoiceEvalContext)
			}

			_ = (Hx_Array_Push[any](HxMakeArrayView[any](children), flow))
		}
	}

	if Hx_Field_go_haxe_hxdynamic_getArrayLength(children) > 0 {
		Hx_Field_go_haxe_hxdynamic_setField(result, "children", children)
		var _hx_tmp_8 bool
		_ = _hx_tmp_8
		if this.Hx_Field_finished {
			_hx_tmp_8 = (this.Hx_Field_stack.Len() == 0)
		} else {
			_hx_tmp_8 = false
		}

		if _hx_tmp_8 {
			Hx_Field_go_haxe_hxdynamic_setField(result, "finished", true)
		}
	}

	if Hx_Field_reflect_fields(insertions).Len() > 0 {
		Hx_Field_go_haxe_hxdynamic_setField(result, "insertions", insertions)
	}

	this.Hx_Field__saveInsertions = nil
	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializePendingChoiceOptions(flow *Hx_Obj_loreline_interpreter, insertions any) HxArray[any] {
	if flow.Hx_Field_pendingChoiceOptions == nil {
		return nil
	}

	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_choiceoption]
		_ = _g2
		_g2 = flow.Hx_Field_pendingChoiceOptions
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var opt *Hx_Obj_loreline_choiceoption
			_ = opt
			opt = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_serializeChoiceOption(opt, insertions)))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializePendingDialogue(flow *Hx_Obj_loreline_interpreter) any {
	var pending any
	_ = pending
	pending = flow.Hx_Field_pendingDialogue
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(pending, nil))) {
		return nil
	}

	var hx_dyncast_219 any
	_ = hx_dyncast_219
	hx_dyncast_219 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(pending, "node"), "Hx_Obj_loreline_astnode")
	var _hx_tmp_0 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_0
	if hx_dyncast_219 != nil {
		_hx_tmp_0 = hx_dyncast_219.(interface {
			Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
		}).Hx_As_Hx_Obj_loreline_astnode()
	} else {
		_hx_tmp_0 = nil
	}

	var result any
	_ = result
	result = any(map[string]any{"node": this.VTable.Hx_Field_serializeNodeReference(_hx_tmp_0), "text": Hx_Field_go_haxe_hxdynamic_getField(pending, "text")})
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(pending, "character"), nil)))))), ((any)(false))))) {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "character", Hx_Field_go_haxe_hxdynamic_getField(pending, "character"))
	}

	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(pending, "tags"), ((any)(nil))))) {
		_hx_tmp_2 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(pending, "tags"), "length"), ((int32)(0))))))
	} else {
		_hx_tmp_2 = false
	}

	if _hx_tmp_2 {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_texttag]
			_ = _g2
			_g2 = HxAnyToTypedArray[*Hx_Obj_loreline_texttag](Hx_Field_go_haxe_hxdynamic_getField(pending, "tags"))
			for {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g1
				if !(_hx_tmp_3 < _g2.Len()) {
					break
				}

				var tag *Hx_Obj_loreline_texttag
				_ = tag
				tag = _g2.Get(_g1)
				_g1++
				var _hx_tmp_4 HxArray[any]
				_ = _hx_tmp_4
				_hx_tmp_4 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_4, this.VTable.Hx_Field_serializeTextTag(tag)))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(result, "tags", _g)
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeChoiceEvalContext(flow *Hx_Obj_loreline_interpreter) HxArray[any] {
	if flow.Hx_Field__choiceEvalTexts.Len() == 0 {
		return nil
	}

	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 int32
		_ = _g2
		_g2 = flow.Hx_Field__choiceEvalTexts.Len()
		for _g1 < _g2 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			_g1 = (_g1 + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var entry any
			_ = entry
			entry = any(map[string]any{"text": ((any)(flow.Hx_Field__choiceEvalTexts.Get(i)))})
			if !flow.Hx_Field__choiceEvalEnabled.Get(i) {
				Hx_Field_go_haxe_hxdynamic_setField(entry, "disabled", true)
			}

			_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), entry))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restore(saveData any) {
	if !this.VTable.Hx_Field_isRoot() {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("restore() can only be called on a root interpreter", this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saveData, "version"), ((int32)(1))))) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(((any)("Unsupported save version: ")), Hx_Field_go_haxe_hxdynamic_getField(saveData, "version"))), this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_interpreter]
		_ = _g1
		_g1 = (Hx_Array_Copy[*Hx_Obj_loreline_interpreter](this.Hx_Field_context.Hx_Field_children))
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var child *Hx_Obj_loreline_interpreter
			_ = child
			child = _g1.Get(_g)
			_g++
			child.VTable.Hx_Field_dispose()
		}
	}

	this.Hx_Field_context.VTable.Hx_Field_clearPendingFlows()
	this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field_clear()
	this.Hx_Field_finished = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(saveData, "finished"), ((any)(true))))))
	this.Hx_Field_context.Hx_Field_nextInsertionId = 1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saveData, "insertions"), nil))) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = Hx_Field_reflect_fields(Hx_Field_go_haxe_hxdynamic_getField(saveData, "insertions"))
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var idStr string
			_ = idStr
			idStr = _g1.Get(_g)
			_g++
			var id HxNullable[int32]
			_ = id
			id = Hx_Field_std_parseInt(idStr)
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if (id).Valid != false {
				_hx_tmp_2 = ((id).Valid && true && ((id).Value >= this.Hx_Field_context.Hx_Field_nextInsertionId) || (!(id).Valid && !true))
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				this.Hx_Field_context.Hx_Field_nextInsertionId = ((id).Value + 1)
			}
		}
	}

	var restoredInsertions *Hx_Obj_haxe_imap
	_ = restoredInsertions
	restoredInsertions = (func() *Hx_Obj_haxe_imap {
		_hx_icast_221 := Hx_Obj_haxe_ds_intmap_CreateInstance()
		if _hx_icast_221 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_221, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_221.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	this.Hx_Field__restoreSavedInsertions = Hx_Field_go_haxe_hxdynamic_getField(saveData, "insertions")
	this.Hx_Field__restoreInsertionsCache = restoredInsertions
	var _hx_tmp_3 *Hx_Obj_loreline_runtimestate
	_ = _hx_tmp_3
	_hx_tmp_3 = this.Hx_Field_context.Hx_Field_topLevelState
	this.VTable.Hx_Field_restoreState(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(saveData, "state"))
	this.VTable.Hx_Field_restoreCharacters(Hx_Field_go_haxe_hxdynamic_getField(saveData, "characters"))
	this.VTable.Hx_Field_restoreNodeStates(Hx_Field_go_haxe_hxdynamic_getField(saveData, "nodeStates"))
	var hx_nullable_t60 any
	_ = hx_nullable_t60
	hx_nullable_t60 = Hx_Field_go_haxe_hxdynamic_getField(saveData, "random")
	var _hx_tmp_4 HxNullable[float64]
	_ = _hx_tmp_4
	if hx_nullable_t60 == nil {
		_hx_tmp_4 = HxNullable[float64]{}
	} else {
		_hx_tmp_4 = HxNullable[float64]{Value: ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(hx_nullable_t60))), Valid: true}
	}

	var savedRandom HxNullable[float64]
	_ = savedRandom
	savedRandom = _hx_tmp_4
	if (savedRandom).Valid != false {
		this.Hx_Field_context.Hx_Field_random = Hx_Obj_loreline_random_CreateInstance(savedRandom)
	}

	var _hx_tmp_5 HxArray[any]
	_ = _hx_tmp_5
	_hx_tmp_5 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saveData, "stack"))
	var _hx_tmp_6 HxArray[any]
	_ = _hx_tmp_6
	_hx_tmp_6 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saveData, "pendingChoiceOptions"))
	var _hx_tmp_7 any
	_ = _hx_tmp_7
	_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(saveData, "pendingDialogue")
	var _hx_tmp_8 HxArray[any]
	_ = _hx_tmp_8
	_hx_tmp_8 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saveData, "choiceEvalContext"))
	this.VTable.Hx_Field_restoreFlow(_hx_tmp_5, _hx_tmp_6, _hx_tmp_7, _hx_tmp_8, Hx_Field_go_haxe_hxdynamic_getField(saveData, "insertions"), restoredInsertions)
	var _hx_tmp_9 bool
	_ = _hx_tmp_9
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saveData, "children"), ((any)(nil))))) {
		_hx_tmp_9 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_gt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(saveData, "children"), "length"), ((int32)(0))))))
	} else {
		_hx_tmp_9 = false
	}

	if _hx_tmp_9 {
		this.Hx_Field_context.Hx_Field_pendingRestoredFlows = (func() *Hx_Obj_haxe_imap {
			_hx_icast_220 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
			if _hx_icast_220 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_220, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_220.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		this.Hx_Field_context.Hx_Field_pendingRestoredKeys = HxMakeArray[string]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saveData, "children"))
			for {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = _g
				if !(_hx_tmp_10 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var flow any
				_ = flow
				flow = _g1.Get(_g)
				_g++
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(flow, "key"), ((any)(HxStringNull))))) {
					_hx_tmp_11 = true
				} else {
					_hx_tmp_11 = func() func(_hx_param_key string) bool {
						var _hx_closure_src_587 func(_hx_param_key any) bool
						_ = _hx_closure_src_587
						_hx_closure_src_587 = this.Hx_Field_context.Hx_Field_pendingRestoredFlows.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_587(((any)(_hx_param_key)))
						}
					}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(flow, "key")))
				}

				if _hx_tmp_11 {
					continue
				}

				this.Hx_Field_context.Hx_Field_pendingRestoredFlows.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_getField(flow, "key"), flow)
				var _hx_tmp_12 HxArray[string]
				_ = _hx_tmp_12
				_hx_tmp_12 = this.Hx_Field_context.Hx_Field_pendingRestoredKeys
				_ = (Hx_Array_Push[string](_hx_tmp_12, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(flow, "key"))))
			}
		}

		this.Hx_Field_context.Hx_Field_pendingSavedInsertions = Hx_Field_go_haxe_hxdynamic_getField(saveData, "insertions")
		this.Hx_Field_context.Hx_Field_pendingInsertionsCache = restoredInsertions
	}

	this.Hx_Field__restoreSavedInsertions = nil
	this.Hx_Field__restoreInsertionsCache = nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreFlow(savedStack HxArray[any], savedPendingChoiceOptions HxArray[any], savedPendingDialogue any, savedChoiceEvalContext HxArray[any], savedInsertions any, restoredInsertions Hx_Typedef_map) {
	(Hx_Array_Resize[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack, 0))
	this.Hx_Field_nextScopeId = 1
	this.Hx_Field_pendingChoiceOptions = nil
	this.Hx_Field_pendingDialogue = nil
	(Hx_Array_Resize[string](this.Hx_Field__choiceEvalTexts, 0))
	(Hx_Array_Resize[bool](this.Hx_Field__choiceEvalEnabled, 0))
	var prevSavedInsertions any
	_ = prevSavedInsertions
	prevSavedInsertions = this.Hx_Field__restoreSavedInsertions
	var prevInsertionsCache Hx_Typedef_map
	_ = prevInsertionsCache
	prevInsertionsCache = this.Hx_Field__restoreInsertionsCache
	this.Hx_Field__restoreSavedInsertions = savedInsertions
	this.Hx_Field__restoreInsertionsCache = restoredInsertions
	if !this.VTable.Hx_Field_restoreStack(savedStack, savedInsertions, restoredInsertions) {
		this.Hx_Field_beatToResume = this.VTable.Hx_Field_restoreBeatToResume(savedStack)
	}

	if savedPendingChoiceOptions != nil {
		this.Hx_Field_pendingChoiceOptions = HxMakeArray[*Hx_Obj_loreline_choiceoption]()
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(savedPendingChoiceOptions)) {
					break
				}

				var savedOpt any
				_ = savedOpt
				savedOpt = savedPendingChoiceOptions.Get(_g)
				_g++
				var opt *Hx_Obj_loreline_choiceoption
				_ = opt
				opt = this.VTable.Hx_Field_restoreChoiceOption(savedOpt, savedInsertions, restoredInsertions)
				if opt != nil {
					_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](this.Hx_Field_pendingChoiceOptions, opt))
				}
			}
		}
	}

	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedPendingDialogue, nil))) {
		_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "node"), nil))))
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var node *Hx_Obj_loreline_node
		_ = node
		node = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeById(Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "node"), "id"))))
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if node != nil {
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = ((any)(node.VTable.Hx_Field_type()))
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_2 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_4, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "node"), "type")))))
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			var hx_nullable_t61 any
			_ = hx_nullable_t61
			hx_nullable_t61 = Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "character")
			var _hx_tmp_5 HxNullable[string]
			_ = _hx_tmp_5
			if hx_nullable_t61 == nil {
				_hx_tmp_5 = HxNullable[string]{Value: HxStringNull}
			} else {
				_hx_tmp_5 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t61), Valid: true}
			}

			var savedPendingDialogue1 HxNullable[string]
			_ = savedPendingDialogue1
			savedPendingDialogue1 = _hx_tmp_5
			var savedPendingDialogue2 string
			_ = savedPendingDialogue2
			savedPendingDialogue2 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "text"))
			var tmp HxArray[*Hx_Obj_loreline_texttag]
			_ = tmp
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "tags"), ((any)(nil))))) {
				var _g HxArray[*Hx_Obj_loreline_texttag]
				_ = _g
				_g = HxMakeArray[*Hx_Obj_loreline_texttag]()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 HxArray[any]
					_ = _g2
					_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(savedPendingDialogue, "tags"))
					for {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = _g1
						if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
							break
						}

						var t any
						_ = t
						t = _g2.Get(_g1)
						_g1++
						var _hx_tmp_7 HxArray[*Hx_Obj_loreline_texttag]
						_ = _hx_tmp_7
						_hx_tmp_7 = _g
						_ = (Hx_Array_Push[*Hx_Obj_loreline_texttag](_hx_tmp_7, Hx_Obj_loreline_texttag_CreateInstance(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t, "closing"), ((any)(true)))))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(t, "value")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(t, "offset")))))))
					}
				}

				tmp = _g
			} else {
				tmp = HxMakeArray[*Hx_Obj_loreline_texttag]()
			}

			var _hx_tmp_8 any
			_ = _hx_tmp_8
			_hx_tmp_8 = ((any)(node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()))
			var hx_nullable_f209 HxNullable[string]
			_ = hx_nullable_f209
			hx_nullable_f209 = savedPendingDialogue1
			var _hx_tmp_9 any
			_ = _hx_tmp_9
			if (hx_nullable_f209).Valid == true {
				_hx_tmp_9 = (hx_nullable_f209).Value
			} else {
				_hx_tmp_9 = nil
			}

			this.Hx_Field_pendingDialogue = any(map[string]any{"node": _hx_tmp_8, "character": _hx_tmp_9, "text": ((any)(savedPendingDialogue2)), "tags": ((any)(tmp))})
		}
	}

	if savedChoiceEvalContext != nil {
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_10 int32
			_ = _hx_tmp_10
			_hx_tmp_10 = _g
			if !(_hx_tmp_10 < Hx_Field_go_haxe_hxdynamic_getArrayLength(savedChoiceEvalContext)) {
				break
			}

			var entry any
			_ = entry
			entry = savedChoiceEvalContext.Get(_g)
			_g++
			var _hx_tmp_11 HxArray[string]
			_ = _hx_tmp_11
			_hx_tmp_11 = this.Hx_Field__choiceEvalTexts
			_ = (Hx_Array_Push[string](_hx_tmp_11, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "text"))))
			var _hx_tmp_12 HxArray[bool]
			_ = _hx_tmp_12
			_hx_tmp_12 = this.Hx_Field__choiceEvalEnabled
			_ = (Hx_Array_Push[bool](_hx_tmp_12, ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(entry, "disabled"), ((any)(true))))))))
		}
	}

	this.Hx_Field__restoreSavedInsertions = prevSavedInsertions
	this.Hx_Field__restoreInsertionsCache = prevInsertionsCache
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resume() {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if this.Hx_Field_stack.Len() == 0 {
		_hx_tmp_1 = this.Hx_Field_finished
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = this.VTable.Hx_Field_isRoot()
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		this.VTable.Hx_Field_finish()
		return
	}

	if this.Hx_Field_stack.Len() == 0 {
		this.VTable.Hx_Field_start(nil)
		return
	}

	this.VTable.Hx_Field_attachBeforeRun()
	var done *Hx_Obj_loreline_evalnext
	_ = done
	done = this.VTable.Hx_Field_wrapNext(this.VTable.Hx_Field_finish)
	this.Hx_Field_finishTrigger = done
	this.VTable.Hx_Field_resumeNode(this.Hx_Field_stack.Get(0).Hx_Field_node, 0, done.Hx_Field_cb)
	done.Hx_Field_sync = false
	this.VTable.Hx_Field_flush()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeFromLevel(scopeLevel int32, next func()) {
	var done *Hx_Obj_loreline_evalnext
	_ = done
	done = this.VTable.Hx_Field_wrapNext(next)
	this.VTable.Hx_Field_resumeNode(this.Hx_Field_stack.Get(scopeLevel).Hx_Field_node, scopeLevel, done.Hx_Field_cb)
	done.Hx_Field_sync = false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_spawn(key string, handleDialogue__isOptional any, handleChoice__isOptional any, handleFinish__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter {
	handleDialogue := HxDefault[Hx_Typedef_loreline_dialoguehandler]()
	_ = handleDialogue

	if handleDialogue__isOptional != nil {
		handleDialogue = HxConvert[Hx_Typedef_loreline_dialoguehandler](handleDialogue__isOptional)
	} else {
		handleDialogue = nil
	}
	handleChoice := HxDefault[Hx_Typedef_loreline_choicehandler]()
	_ = handleChoice

	if handleChoice__isOptional != nil {
		handleChoice = HxConvert[Hx_Typedef_loreline_choicehandler](handleChoice__isOptional)
	} else {
		handleChoice = nil
	}
	handleFinish := HxDefault[Hx_Typedef_loreline_finishhandler]()
	_ = handleFinish

	if handleFinish__isOptional != nil {
		handleFinish = HxConvert[Hx_Typedef_loreline_finishhandler](handleFinish__isOptional)
	} else {
		handleFinish = nil
	}
	options := HxDefault[*Hx_Obj_loreline_interpreteroptions]()
	_ = options

	if options__isOptional != nil {
		options = HxConvert[*Hx_Obj_loreline_interpreteroptions](options__isOptional)
	} else {
		options = nil
	}
	{
		this.VTable.Hx_Field_checkSpawnKey(key)
		var existing *Hx_Obj_loreline_interpreter
		_ = existing
		existing = this.Hx_Field_context.VTable.Hx_Field_childByKey(key)
		if existing != nil {
			existing.VTable.Hx_Field_dispose()
		}

		this.Hx_Field_context.VTable.Hx_Field_removePendingFlow(key)
		return this.VTable.Hx_Field_createChild(key, handleDialogue, handleChoice, handleFinish, options)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeSpawn(key string, handleDialogue__isOptional any, handleChoice__isOptional any, handleFinish__isOptional any, options__isOptional any) *Hx_Obj_loreline_interpreter {
	handleDialogue := HxDefault[Hx_Typedef_loreline_dialoguehandler]()
	_ = handleDialogue

	if handleDialogue__isOptional != nil {
		handleDialogue = HxConvert[Hx_Typedef_loreline_dialoguehandler](handleDialogue__isOptional)
	} else {
		handleDialogue = nil
	}
	handleChoice := HxDefault[Hx_Typedef_loreline_choicehandler]()
	_ = handleChoice

	if handleChoice__isOptional != nil {
		handleChoice = HxConvert[Hx_Typedef_loreline_choicehandler](handleChoice__isOptional)
	} else {
		handleChoice = nil
	}
	handleFinish := HxDefault[Hx_Typedef_loreline_finishhandler]()
	_ = handleFinish

	if handleFinish__isOptional != nil {
		handleFinish = HxConvert[Hx_Typedef_loreline_finishhandler](handleFinish__isOptional)
	} else {
		handleFinish = nil
	}
	options := HxDefault[*Hx_Obj_loreline_interpreteroptions]()
	_ = options

	if options__isOptional != nil {
		options = HxConvert[*Hx_Obj_loreline_interpreteroptions](options__isOptional)
	} else {
		options = nil
	}
	{
		this.VTable.Hx_Field_checkSpawnKey(key)
		var savedInsertions any
		_ = savedInsertions
		savedInsertions = this.Hx_Field_context.Hx_Field_pendingSavedInsertions
		var insertionsCache Hx_Typedef_map
		_ = insertionsCache
		insertionsCache = this.Hx_Field_context.Hx_Field_pendingInsertionsCache
		var flow any
		_ = flow
		flow = this.Hx_Field_context.VTable.Hx_Field_removePendingFlow(key)
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(flow, nil))) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("No saved interpreter to resume with key: ", key), this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
		}

		var existing *Hx_Obj_loreline_interpreter
		_ = existing
		existing = this.Hx_Field_context.VTable.Hx_Field_childByKey(key)
		if existing != nil {
			existing.VTable.Hx_Field_dispose()
		}

		var child *Hx_Obj_loreline_interpreter
		_ = child
		child = this.VTable.Hx_Field_createChild(key, handleDialogue, handleChoice, handleFinish, options)
		var _hx_tmp_0 HxArray[any]
		_ = _hx_tmp_0
		_hx_tmp_0 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(flow, "stack"))
		var _hx_tmp_1 HxArray[any]
		_ = _hx_tmp_1
		_hx_tmp_1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(flow, "pendingChoiceOptions"))
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(flow, "pendingDialogue")
		child.VTable.Hx_Field_restoreFlow(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(flow, "choiceEvalContext")), savedInsertions, insertionsCache)
		return child
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumableSpawnKeys() HxArray[string] {
	if this.Hx_Field_context.Hx_Field_pendingRestoredKeys != nil {
		return (Hx_Array_Copy[string](this.Hx_Field_context.Hx_Field_pendingRestoredKeys))
	} else {
		return HxMakeArray[string]()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_dispose() {
	if this.VTable.Hx_Field_isRoot() {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot dispose a root interpreter", this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
	}

	if this.Hx_Field_disposed {
		return
	}

	this.Hx_Field_disposed = true
	(Hx_Array_Resize[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack, 0))
	this.Hx_Field_syncCallbacks = HxMakeArray[func()]()
	this.Hx_Field_finishTrigger = nil
	this.Hx_Field_beatToResume = nil
	this.Hx_Field_pendingChoiceOptions = nil
	this.Hx_Field_pendingDialogue = nil
	(Hx_Array_Resize[string](this.Hx_Field__choiceEvalTexts, 0))
	(Hx_Array_Resize[bool](this.Hx_Field__choiceEvalEnabled, 0))
	_ = (Hx_Array_Remove[*Hx_Obj_loreline_interpreter](this.Hx_Field_context.Hx_Field_children, this))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_checkSpawnKey(key string) {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if key == HxStringNull {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (int32(utf8.RuneCountInString(key)) == 0)
	}

	if _hx_tmp_0 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("A key is required to spawn an interpreter", this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_createChild(key string, handleDialogue Hx_Typedef_loreline_dialoguehandler, handleChoice Hx_Typedef_loreline_choicehandler, handleFinish Hx_Typedef_loreline_finishhandler, options *Hx_Obj_loreline_interpreteroptions) *Hx_Obj_loreline_interpreter {
	var handleDialogue1 Hx_Typedef_loreline_dialoguehandler
	_ = handleDialogue1
	handleDialogue1 = handleDialogue
	var handleChoice1 Hx_Typedef_loreline_choicehandler
	_ = handleChoice1
	handleChoice1 = handleChoice
	var handleFinish1 Hx_Typedef_loreline_finishhandler
	_ = handleFinish1
	handleFinish1 = handleFinish
	var _hx_tmp_0 Hx_Typedef_loreline_dialoguehandler
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(handleDialogue1)), ((any)(nil))))) {
		_hx_tmp_0 = handleDialogue1
	} else {
		_hx_tmp_0 = this.Hx_Field_handleDialogue
	}

	var _hx_tmp_1 Hx_Typedef_loreline_choicehandler
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(handleChoice1)), ((any)(nil))))) {
		_hx_tmp_1 = handleChoice1
	} else {
		_hx_tmp_1 = this.Hx_Field_handleChoice
	}

	var _hx_tmp_2 Hx_Typedef_loreline_finishhandler
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(handleFinish1)), ((any)(nil))))) {
		_hx_tmp_2 = handleFinish1
	} else {
		_hx_tmp_2 = this.Hx_Field_handleFinish
	}

	var child *Hx_Obj_loreline_interpreter
	_ = child
	child = Hx_Obj_loreline_interpreter_CreateInstance(this.Hx_Field_context.Hx_Field_script, _hx_tmp_0, _hx_tmp_1, _hx_tmp_2, options, this.Hx_Field_context)
	child.Hx_Field_key = key
	_ = (Hx_Array_Push[*Hx_Obj_loreline_interpreter](this.Hx_Field_context.Hx_Field_children, child))
	return child
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_attachBeforeRun() {
	if this.Hx_Field_disposed {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot run a disposed interpreter (key: ", this.Hx_Field_key), ")"), this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
	}

	if this.VTable.Hx_Field_isRoot() {
		this.Hx_Field_finished = false
		return
	}

	var _hx_tmp_0 HxArray[*Hx_Obj_loreline_interpreter]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_context.Hx_Field_children
	var _hx_tmp_1 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_1
	_hx_tmp_1 = this
	if (Hx_Array_IndexOf[*Hx_Obj_loreline_interpreter](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) == -1 {
		var other *Hx_Obj_loreline_interpreter
		_ = other
		other = this.Hx_Field_context.VTable.Hx_Field_childByKey(this.Hx_Field_key)
		if other != nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Another interpreter is already running with key: ", this.Hx_Field_key), this.Hx_Field_context.Hx_Field_script.Hx_Field_pos)))))
		}

		_ = (Hx_Array_Push[*Hx_Obj_loreline_interpreter](this.Hx_Field_context.Hx_Field_children, this))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getCharacter(name string) any {
	var tmp *Hx_Obj_loreline_runtimecharacter
	_ = tmp
	tmp = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
		var _hx_closure_src_588 func(_hx_param_key any) any
		_ = _hx_closure_src_588
		_hx_closure_src_588 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
			var hx_dyncast_222 any
			_ = hx_dyncast_222
			hx_dyncast_222 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_588(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimecharacter
			_ = _hx_tmp_0
			if hx_dyncast_222 != nil {
				_hx_tmp_0 = hx_dyncast_222.(interface {
					Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
				}).Hx_As_Hx_Obj_loreline_runtimecharacter()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(name)
	if tmp != nil {
		return tmp.Hx_Field_fields
	} else {
		return nil
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getCharacterField(character string, name string) any {
	var tmp *Hx_Obj_loreline_runtimecharacter
	_ = tmp
	tmp = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
		var _hx_closure_src_589 func(_hx_param_key any) any
		_ = _hx_closure_src_589
		_hx_closure_src_589 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
			var hx_dyncast_223 any
			_ = hx_dyncast_223
			hx_dyncast_223 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_589(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimecharacter
			_ = _hx_tmp_0
			if hx_dyncast_223 != nil {
				_hx_tmp_0 = hx_dyncast_223.(interface {
					Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
				}).Hx_As_Hx_Obj_loreline_runtimecharacter()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(character)
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	if tmp != nil {
		_hx_tmp_1 = tmp.Hx_Field_fields
	} else {
		_hx_tmp_1 = nil
	}

	var fields any
	_ = fields
	fields = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(fields, nil))) {
		return this.VTable.Hx_Field_hostValueOf(Hx_Field_loreline_objects_getField(this, fields, name))
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setCharacterField(character string, name string, value any) {
	var fields any
	_ = fields
	fields = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
		var _hx_closure_src_590 func(_hx_param_key any) any
		_ = _hx_closure_src_590
		_hx_closure_src_590 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
			var hx_dyncast_224 any
			_ = hx_dyncast_224
			hx_dyncast_224 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_590(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimecharacter
			_ = _hx_tmp_0
			if hx_dyncast_224 != nil {
				_hx_tmp_0 = hx_dyncast_224.(interface {
					Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
				}).Hx_As_Hx_Obj_loreline_runtimecharacter()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(character).Hx_Field_fields
	var _hx_tmp_1 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_1
	_hx_tmp_1 = this
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = fields
	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = name
	Hx_Field_loreline_objects_setField(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, this.VTable.Hx_Field_hostValueToRuntime(value))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_eachResolutionScope(startIndex int32, visit func(scope *Hx_Obj_loreline_runtimescope) bool) {
	var _hx_tmp_0 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_stack
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if startIndex == -1 {
		_hx_tmp_1 = (this.Hx_Field_stack.Len() - 1)
	} else {
		_hx_tmp_1 = startIndex
	}

	this.VTable.Hx_Field_walkResolutionScopes(_hx_tmp_0, _hx_tmp_1, visit)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_walkResolutionScopes(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32, visit func(scope *Hx_Obj_loreline_runtimescope) bool) {
	var i int32
	_ = i
	i = startIndex
	for i >= 0 {
		var scope *Hx_Obj_loreline_runtimescope
		_ = scope
		scope = scopes.Get(i)
		if visit(scope) {
			return
		}

		if scope.Hx_Field_captured != nil {
			var _hx_tmp_0 HxArray[*Hx_Obj_loreline_runtimescope]
			_ = _hx_tmp_0
			_hx_tmp_0 = scope.Hx_Field_captured
			this.VTable.Hx_Field_walkResolutionScopes(_hx_tmp_0, (scope.Hx_Field_captured.Len() - 1), visit)
			return
		}

		i--
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_hostValueOf(value any) any {
	var asCharacter *Hx_Obj_loreline_runtimecharacterref
	_ = asCharacter
	asCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(value)
	if asCharacter != nil {
		var marker any
		_ = marker
		marker = any(map[string]any{"type": ((any)("$characterRef")), "name": ((any)(asCharacter.Hx_Field_name))})
		return marker
	}

	var asBeat *Hx_Obj_loreline_nbeatdecl
	_ = asBeat
	asBeat = Hx_Field_loreline_runtimebeatref_beatOf(value)
	if asBeat != nil {
		var marker any
		_ = marker
		marker = any(map[string]any{"type": ((any)("$beatRef")), "beat": this.VTable.Hx_Field_serializeBeatReference(asBeat)})
		return marker
	}

	return value
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_hostValueToRuntime(value any) any {
	return this.VTable.Hx_Field_hostValueToRuntimeDepth(value, 0)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_hostValueToRuntimeDepth(value any, depth int32) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		_hx_tmp_3 = true
	} else {
		_hx_tmp_3 = Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
	}

	if _hx_tmp_0 {
		return value
	}

	if depth >= Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH {
		return value
	}

	if Hx_Field_loreline_runtimecharacterref_characterOf(value) != nil {
		return value
	}

	if Hx_Field_loreline_runtimebeatref_beatOf(value) != nil {
		return value
	}

	if Hx_Field_loreline_arrays_isArray(value) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(value)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_4
				var item any
				_ = item
				item = Hx_Field_loreline_arrays_arrayGet(value, i)
				var converted any
				_ = converted
				converted = this.VTable.Hx_Field_hostValueToRuntimeDepth(item, (depth + 1))
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(converted, item))) {
					Hx_Field_loreline_arrays_arraySet(value, i, converted)
				}
			}
		}

		return value
	}

	if Hx_Field_loreline_objects_isFields(value) {
		var markerTypeValue any
		_ = markerTypeValue
		markerTypeValue = Hx_Field_loreline_objects_getField(this, value, "type")
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(markerTypeValue, nil))) {
			_hx_tmp_5 = Hx_Field_std_isOfType(markerTypeValue, Hx_Obj_string_RTTI)
		} else {
			_hx_tmp_5 = false
		}

		if _hx_tmp_5 {
			var markerType string
			_ = markerType
			markerType = Hx_Field_go_haxe_hxdynamic_toString(markerTypeValue)
			if markerType == "$characterRef" {
				var name any
				_ = name
				name = Hx_Field_loreline_objects_getField(this, value, "name")
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(name, nil))) {
					_hx_tmp_7 = Hx_Field_std_isOfType(name, Hx_Obj_string_RTTI)
				} else {
					_hx_tmp_7 = false
				}

				if _hx_tmp_7 {
					_hx_tmp_6 = func() func(_hx_param_key string) bool {
						var _hx_closure_src_591 func(_hx_param_key any) bool
						_ = _hx_closure_src_591
						_hx_closure_src_591 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
						return func(_hx_param_key string) bool {
							return _hx_closure_src_591(((any)(_hx_param_key)))
						}
					}()(Hx_Field_go_haxe_hxdynamic_toString(name))
				} else {
					_hx_tmp_6 = false
				}

				if _hx_tmp_6 {
					var character *Hx_Obj_loreline_runtimecharacter
					_ = character
					character = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
						var _hx_closure_src_592 func(_hx_param_key any) any
						_ = _hx_closure_src_592
						_hx_closure_src_592 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
						return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
							var hx_dyncast_225 any
							_ = hx_dyncast_225
							hx_dyncast_225 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_592(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
							var _hx_tmp_8 *Hx_Obj_loreline_runtimecharacter
							_ = _hx_tmp_8
							if hx_dyncast_225 != nil {
								_hx_tmp_8 = hx_dyncast_225.(interface {
									Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
								}).Hx_As_Hx_Obj_loreline_runtimecharacter()
							} else {
								_hx_tmp_8 = nil
							}

							return _hx_tmp_8
						}
					}()(Hx_Field_go_haxe_hxdynamic_toString(name))
					if character.Hx_Field_characterRef == nil {
						character.Hx_Field_characterRef = Hx_Obj_loreline_runtimecharacterref_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(name), character)
					}

					return ((any)(character.Hx_Field_characterRef))
				}
			} else {
				if markerType == "$beatRef" {
					var path any
					_ = path
					var beatData any
					_ = beatData
					beatData = Hx_Field_loreline_objects_getField(this, value, "beat")
					var _hx_tmp_9 bool
					_ = _hx_tmp_9
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(beatData, nil))) {
						_hx_tmp_9 = Hx_Field_loreline_objects_isFields(beatData)
					} else {
						_hx_tmp_9 = false
					}

					if _hx_tmp_9 {
						path = Hx_Field_loreline_objects_getField(this, beatData, "path")
					}

					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(path, nil))) {
						path = Hx_Field_loreline_objects_getField(this, value, "name")
					}

					var _hx_tmp_10 bool
					_ = _hx_tmp_10
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(path, nil))) {
						_hx_tmp_10 = Hx_Field_std_isOfType(path, Hx_Obj_string_RTTI)
					} else {
						_hx_tmp_10 = false
					}

					if _hx_tmp_10 {
						var _hx_tmp_11 string
						_ = _hx_tmp_11
						_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_toString(path)
						var beat *Hx_Obj_loreline_nbeatdecl
						_ = beat
						beat = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_findBeatByPathFromNode(_hx_tmp_11, (func() *Hx_Obj_loreline_node {
							_hx_up := this.Hx_Field_context.Hx_Field_script
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_node
							}
							return nil
						})())
						if beat != nil {
							return ((any)(Hx_Obj_loreline_runtimebeatref_CreateInstance(beat, HxMakeArray[*Hx_Obj_loreline_runtimescope]())))
						}
					}
				}
			}
		}

		var keys HxArray[string]
		_ = keys
		keys = Hx_Field_loreline_objects_getFields(this, value)
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				_hx_tmp_12 = _g
				if !(_hx_tmp_12 < keys.Len()) {
					break
				}

				var key string
				_ = key
				key = keys.Get(_g)
				_g++
				var item any
				_ = item
				item = Hx_Field_loreline_objects_getField(this, value, key)
				var converted any
				_ = converted
				converted = this.VTable.Hx_Field_hostValueToRuntimeDepth(item, (depth + 1))
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(converted, item))) {
					Hx_Field_loreline_objects_setField(this, value, key, converted)
				}
			}
		}
	}

	return value
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getStateField(name string) any {
	return this.VTable.Hx_Field_hostValueOf(this.VTable.Hx_Field_getStateFieldRaw(name))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getStateFieldRaw(name string) any {
	var found bool
	_ = found
	found = false
	var result any
	_ = result
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	this.VTable.Hx_Field_eachResolutionScope(-1, func(scope *Hx_Obj_loreline_runtimescope) bool {
		if scope.Hx_Field_state != nil {
			if Hx_Field_loreline_objects_fieldExists(_gthis, scope.Hx_Field_state.Hx_Field_fields, name) {
				result = Hx_Field_loreline_objects_getField(_gthis, scope.Hx_Field_state.Hx_Field_fields, name)
				found = true
				return true
			}
		}

		if scope.Hx_Field_node != nil {
			var key Hx_Typedef_haxe__int64___int64
			_ = key
			key = scope.Hx_Field_node.Hx_Field_id
			var _hx_tmp_1 int64
			_ = _hx_tmp_1
			_hx_tmp_1 = ((int64)(key))
			var stateInNode *Hx_Obj_loreline_runtimestate
			_ = stateInNode
			stateInNode = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
				var _hx_closure_src_593 func(high int32, low int32) any
				_ = _hx_closure_src_593
				_hx_closure_src_593 = _gthis.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
				return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
					var hx_dyncast_226 any
					_ = hx_dyncast_226
					hx_dyncast_226 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_593(high, low), "Hx_Obj_loreline_runtimestate")
					var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
					_ = _hx_tmp_0
					if hx_dyncast_226 != nil {
						_hx_tmp_0 = hx_dyncast_226.(interface {
							Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
						}).Hx_As_Hx_Obj_loreline_runtimestate()
					} else {
						_hx_tmp_0 = nil
					}

					return _hx_tmp_0
				}
			}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
			if stateInNode != nil {
				if Hx_Field_loreline_objects_fieldExists(_gthis, stateInNode.Hx_Field_fields, name) {
					result = Hx_Field_loreline_objects_getField(_gthis, stateInNode.Hx_Field_fields, name)
					found = true
					return true
				}
			}
		}

		return false
	})
	if found {
		return result
	}

	if Hx_Field_loreline_objects_fieldExists(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name) {
		return Hx_Field_loreline_objects_getField(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name)
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setStateField(name string, value any) {
	value = this.VTable.Hx_Field_hostValueToRuntime(value)
	var done bool
	_ = done
	done = false
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	this.VTable.Hx_Field_eachResolutionScope(-1, func(scope *Hx_Obj_loreline_runtimescope) bool {
		if scope.Hx_Field_state != nil {
			if Hx_Field_loreline_objects_fieldExists(_gthis, scope.Hx_Field_state.Hx_Field_fields, name) {
				Hx_Field_loreline_objects_setField(_gthis, scope.Hx_Field_state.Hx_Field_fields, name, value)
				done = true
				return true
			}
		}

		if scope.Hx_Field_node != nil {
			var key Hx_Typedef_haxe__int64___int64
			_ = key
			key = scope.Hx_Field_node.Hx_Field_id
			var _hx_tmp_1 int64
			_ = _hx_tmp_1
			_hx_tmp_1 = ((int64)(key))
			var stateInNode *Hx_Obj_loreline_runtimestate
			_ = stateInNode
			stateInNode = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
				var _hx_closure_src_594 func(high int32, low int32) any
				_ = _hx_closure_src_594
				_hx_closure_src_594 = _gthis.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
				return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
					var hx_dyncast_227 any
					_ = hx_dyncast_227
					hx_dyncast_227 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_594(high, low), "Hx_Obj_loreline_runtimestate")
					var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
					_ = _hx_tmp_0
					if hx_dyncast_227 != nil {
						_hx_tmp_0 = hx_dyncast_227.(interface {
							Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
						}).Hx_As_Hx_Obj_loreline_runtimestate()
					} else {
						_hx_tmp_0 = nil
					}

					return _hx_tmp_0
				}
			}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
			if stateInNode != nil {
				if Hx_Field_loreline_objects_fieldExists(_gthis, stateInNode.Hx_Field_fields, name) {
					Hx_Field_loreline_objects_setField(_gthis, stateInNode.Hx_Field_fields, name, value)
					done = true
					return true
				}
			}
		}

		return false
	})
	if done {
		return
	}

	Hx_Field_loreline_objects_setField(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name, value)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getTopLevelStateField(name string) any {
	return this.VTable.Hx_Field_hostValueOf(Hx_Field_loreline_objects_getField(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setTopLevelStateField(name string, value any) {
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = name
	Hx_Field_loreline_objects_setField(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, this.VTable.Hx_Field_hostValueToRuntime(value))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_currentNode() *Hx_Obj_loreline_astnode {
	var scope *Hx_Obj_loreline_runtimescope
	_ = scope
	scope = this.VTable.Hx_Field_get_currentScope()
	if scope == nil {
		return nil
	}

	if scope.Hx_Field_head != nil {
		return scope.Hx_Field_head
	} else {
		return scope.Hx_Field_node
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_currentNodeFilePath(rootPath string) string {
	var node *Hx_Obj_loreline_astnode
	_ = node
	node = this.VTable.Hx_Field_currentNode()
	if (node == nil) || (rootPath == HxStringNull) {
		return rootPath
	}

	return this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeFilePath((func() *Hx_Obj_loreline_node {
		_hx_up := node
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})(), rootPath)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeScope(scope *Hx_Obj_loreline_runtimescope, insertions any) any {
	var result any
	_ = result
	result = any(map[string]any{"id": ((any)(scope.Hx_Field_id))})
	_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](this.Hx_Field__serializingScopes, scope))
	if scope.Hx_Field_beat != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "beat", this.VTable.Hx_Field_serializeBeatReference(scope.Hx_Field_beat))
	}

	if scope.Hx_Field_node != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "node", this.VTable.Hx_Field_serializeNodeReference(scope.Hx_Field_node))
	}

	if scope.Hx_Field_state != nil {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "state", this.VTable.Hx_Field_serializeState(scope.Hx_Field_state))
	}

	if scope.Hx_Field_beats != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_nbeatdecl]
			_ = _g2
			_g2 = scope.Hx_Field_beats
			for {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g1
				if !(_hx_tmp_3 < _g2.Len()) {
					break
				}

				var beat *Hx_Obj_loreline_nbeatdecl
				_ = beat
				beat = _g2.Get(_g1)
				_g1++
				var _hx_tmp_4 HxArray[any]
				_ = _hx_tmp_4
				_hx_tmp_4 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_4, this.VTable.Hx_Field_serializeBeatReference(beat)))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(result, "beats", _g)
	}

	if scope.Hx_Field_head != nil {
		var _hx_tmp_5 any
		_ = _hx_tmp_5
		_hx_tmp_5 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_5, "head", this.VTable.Hx_Field_serializeNodeReference(scope.Hx_Field_head))
	}

	if scope.Hx_Field_insertion != nil {
		var _hx_tmp_6 any
		_ = _hx_tmp_6
		_hx_tmp_6 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_6, "insertion", this.VTable.Hx_Field_serializeInsertion(scope.Hx_Field_insertion, insertions))
	}

	if scope.Hx_Field_captured != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_runtimescope]
			_ = _g2
			_g2 = scope.Hx_Field_captured
			for {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _g1
				if !(_hx_tmp_7 < _g2.Len()) {
					break
				}

				var captured *Hx_Obj_loreline_runtimescope
				_ = captured
				captured = _g2.Get(_g1)
				_g1++
				var _hx_tmp_8 HxArray[any]
				_ = _hx_tmp_8
				_hx_tmp_8 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_8, this.VTable.Hx_Field_serializeScope(captured, insertions)))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(result, "captured", _g)
	}

	if scope.Hx_Field_shuffle != nil {
		var _hx_tmp_9 any
		_ = _hx_tmp_9
		_hx_tmp_9 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_9, "shuffle", (Hx_Array_Copy[int32](scope.Hx_Field_shuffle)))
	}

	_ = (Hx_Array_Pop[*Hx_Obj_loreline_runtimescope](this.Hx_Field__serializingScopes).Value)
	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeInsertion(insertion *Hx_Obj_loreline_runtimeinsertion, insertions any) HxNullable[int32] {
	var key string
	_ = key
	key = Hx_Field_std_string(insertion.Hx_Field_id)
	var serialized any
	_ = serialized
	serialized = Hx_Field_reflect_field(insertions, key)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(serialized, nil))) {
		serialized = any(map[string]any{})
		Hx_Field_reflect_setField(insertions, key, serialized)
		if insertion.Hx_Field_options != nil {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_choiceoption]
				_ = _g2
				_g2 = insertion.Hx_Field_options
				for {
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = _g1
					if !(_hx_tmp_0 < _g2.Len()) {
						break
					}

					var opt *Hx_Obj_loreline_choiceoption
					_ = opt
					opt = _g2.Get(_g1)
					_g1++
					var _hx_tmp_1 HxArray[any]
					_ = _hx_tmp_1
					_hx_tmp_1 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_serializeChoiceOption(opt, insertions)))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(serialized, "options", _g)
		}

		if insertion.Hx_Field_origin != nil {
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = serialized
			Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "origin", this.VTable.Hx_Field_serializeNodeReference((func() *Hx_Obj_loreline_astnode {
				_hx_up := insertion.Hx_Field_origin
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()))
		}

		if insertion.Hx_Field_stack != nil {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _g2
				_g2 = insertion.Hx_Field_stack
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g1
					if !(_hx_tmp_3 < _g2.Len()) {
						break
					}

					var scope *Hx_Obj_loreline_runtimescope
					_ = scope
					scope = _g2.Get(_g1)
					_g1++
					var _hx_tmp_4 HxArray[any]
					_ = _hx_tmp_4
					_hx_tmp_4 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_4, this.VTable.Hx_Field_serializeScope(scope, insertions)))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(serialized, "stack", _g)
		}

		if insertion.Hx_Field_parentPartialOptions != nil {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_choiceoption]
				_ = _g2
				_g2 = insertion.Hx_Field_parentPartialOptions
				for {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = _g1
					if !(_hx_tmp_5 < _g2.Len()) {
						break
					}

					var opt *Hx_Obj_loreline_choiceoption
					_ = opt
					opt = _g2.Get(_g1)
					_g1++
					var _hx_tmp_6 HxArray[any]
					_ = _hx_tmp_6
					_hx_tmp_6 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_6, this.VTable.Hx_Field_serializeChoiceOption(opt, insertions)))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(serialized, "parentPartialOptions", _g)
			Hx_Field_go_haxe_hxdynamic_setField(serialized, "parentNextOptionIndex", insertion.Hx_Field_parentNextOptionIndex)
		}

		if insertion.Hx_Field_kind == ((int32)(1)) {
			Hx_Field_go_haxe_hxdynamic_setField(serialized, "kind", "when")
		}

		if insertion.Hx_Field_rules != nil {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[any]
				_ = _g2
				_g2 = insertion.Hx_Field_rules
				for {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _g1
					if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
						break
					}

					var candidate any
					_ = candidate
					candidate = _g2.Get(_g1)
					_g1++
					var _hx_tmp_8 HxArray[any]
					_ = _hx_tmp_8
					_hx_tmp_8 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_8, this.VTable.Hx_Field_serializeWhenCandidate(candidate, insertions)))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(serialized, "rules", _g)
		}

		if insertion.Hx_Field_parentPartialRules != nil {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[any]
				_ = _g2
				_g2 = insertion.Hx_Field_parentPartialRules
				for {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = _g1
					if !(_hx_tmp_9 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
						break
					}

					var candidate any
					_ = candidate
					candidate = _g2.Get(_g1)
					_g1++
					var _hx_tmp_10 HxArray[any]
					_ = _hx_tmp_10
					_hx_tmp_10 = HxMakeArrayView[any](_g)
					_ = (Hx_Array_Push[any](_hx_tmp_10, this.VTable.Hx_Field_serializeWhenCandidate(candidate, insertions)))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(serialized, "parentPartialRules", _g)
			Hx_Field_go_haxe_hxdynamic_setField(serialized, "parentNextRuleIndex", insertion.Hx_Field_parentNextRuleIndex)
		}

		if insertion.Hx_Field_whenFirst {
			Hx_Field_go_haxe_hxdynamic_setField(serialized, "whenFirst", true)
		}
	}

	return HxNullable[int32]{Value: insertion.Hx_Field_id, Valid: true}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeWhenCandidate(candidate any, insertions any) any {
	var hx_dyncast_228 any
	_ = hx_dyncast_228
	hx_dyncast_228 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_astnode")
	var _hx_tmp_0 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_0
	if hx_dyncast_228 != nil {
		_hx_tmp_0 = hx_dyncast_228.(interface {
			Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
		}).Hx_As_Hx_Obj_loreline_astnode()
	} else {
		_hx_tmp_0 = nil
	}

	var result any
	_ = result
	result = any(map[string]any{"node": this.VTable.Hx_Field_serializeNodeReference(_hx_tmp_0)})
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(candidate, "eligible")))) {
		Hx_Field_go_haxe_hxdynamic_setField(result, "eligible", true)
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(candidate, "insertion"), ((any)(nil))))) {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = result
		var hx_dyncast_229 any
		_ = hx_dyncast_229
		hx_dyncast_229 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "insertion"), "Hx_Obj_loreline_runtimeinsertion")
		var _hx_tmp_2 *Hx_Obj_loreline_runtimeinsertion
		_ = _hx_tmp_2
		if hx_dyncast_229 != nil {
			_hx_tmp_2 = hx_dyncast_229.(interface {
				Hx_As_Hx_Obj_loreline_runtimeinsertion() *Hx_Obj_loreline_runtimeinsertion
			}).Hx_As_Hx_Obj_loreline_runtimeinsertion()
		} else {
			_hx_tmp_2 = nil
		}

		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "insertion", this.VTable.Hx_Field_serializeInsertion(_hx_tmp_2, insertions))
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeChoiceOption(option *Hx_Obj_loreline_choiceoption, insertions any) any {
	var result any
	_ = result
	result = any(map[string]any{"text": ((any)(option.Hx_Field_text))})
	if !option.Hx_Field_enabled {
		Hx_Field_go_haxe_hxdynamic_setField(result, "disabled", true)
	}

	if option.Hx_Field_tags != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_texttag]
			_ = _g2
			_g2 = option.Hx_Field_tags
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < _g2.Len()) {
					break
				}

				var tag *Hx_Obj_loreline_texttag
				_ = tag
				tag = _g2.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[any]
				_ = _hx_tmp_1
				_hx_tmp_1 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_serializeTextTag(tag)))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(result, "tags", _g)
	}

	if option.Hx_Field_node != nil {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "node", this.VTable.Hx_Field_serializeNodeReference((func() *Hx_Obj_loreline_astnode {
			_hx_up := option.Hx_Field_node
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()))
	}

	if option.Hx_Field_insertion != nil {
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = result
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "insertion", this.VTable.Hx_Field_serializeInsertion(option.Hx_Field_insertion, insertions))
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeTextTag(tag *Hx_Obj_loreline_texttag) any {
	var result any
	_ = result
	result = any(map[string]any{"value": ((any)(tag.Hx_Field_value)), "offset": ((any)(tag.Hx_Field_offset))})
	if tag.Hx_Field_closing {
		Hx_Field_go_haxe_hxdynamic_setField(result, "closing", true)
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeState(state *Hx_Obj_loreline_runtimestate) any {
	return this.VTable.Hx_Field_serializeFields(state.Hx_Field_fields, state.Hx_Field_originalFields)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeCharacters() any {
	var result any
	_ = result
	result = any(map[string]any{})
	{
		var _hx_reserved_map *Hx_Obj_haxe_imap
		_ = _hx_reserved_map
		_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
			_hx_icast_230 := this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap()
			if _hx_icast_230 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_230, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_230.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		var _g_map *Hx_Obj_haxe_imap
		_ = _g_map
		_g_map = _hx_reserved_map
		var _g_keys any
		_ = _g_keys
		_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
			var key string
			_ = key
			key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _g_value *Hx_Obj_loreline_runtimecharacter
			_ = _g_value
			_g_value = func() func(k string) *Hx_Obj_loreline_runtimecharacter {
				var _hx_closure_src_595 func(k any) any
				_ = _hx_closure_src_595
				_hx_closure_src_595 = _g_map.VTable.Hx_Field_get
				return func(k string) *Hx_Obj_loreline_runtimecharacter {
					var hx_dyncast_231 any
					_ = hx_dyncast_231
					hx_dyncast_231 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_595(((any)(k))), "Hx_Obj_loreline_runtimecharacter")
					var _hx_tmp_2 *Hx_Obj_loreline_runtimecharacter
					_ = _hx_tmp_2
					if hx_dyncast_231 != nil {
						_hx_tmp_2 = hx_dyncast_231.(interface {
							Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
						}).Hx_As_Hx_Obj_loreline_runtimecharacter()
					} else {
						_hx_tmp_2 = nil
					}

					return _hx_tmp_2
				}
			}()(key)
			var _g_key string
			_ = _g_key
			_g_key = key
			var name string
			_ = name
			name = _g_key
			var character *Hx_Obj_loreline_runtimecharacter
			_ = character
			character = _g_value
			{
				var serialized any
				_ = serialized
				serialized = this.VTable.Hx_Field_serializeCharacter(character)
				if Hx_Field_reflect_fields(Hx_Field_go_haxe_hxdynamic_getField(serialized, "fields")).Len() > 0 {
					Hx_Field_reflect_setField(result, name, serialized)
				}
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeNodeStates() any {
	var result any
	_ = result
	result = any(map[string]any{})
	{
		var _g_map *Hx_Obj_loreline_int64map
		_ = _g_map
		_g_map = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map
		var _g_index int32
		_ = _g_index
		_g_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g_index
			if _hx_tmp_1 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Len() {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			_g_index++
		}

		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g_index
			if !(_hx_tmp_2 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Len()) {
				break
			}

			var k1 int32
			_ = k1
			k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](_g_map.Hx_Field__keys1).Get(_g_index))))
			var k2 int32
			_ = k2
			k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](_g_map.Hx_Field__keys2).Get(_g_index))))
			var hx_dyncast_232 any
			_ = hx_dyncast_232
			hx_dyncast_232 = Hx_Field_go_haxe_hxdynamic_castClass(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Get(_g_index), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_3 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_3
			if hx_dyncast_232 != nil {
				_hx_tmp_3 = hx_dyncast_232.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_3 = nil
			}

			var v *Hx_Obj_loreline_runtimestate
			_ = v
			v = _hx_tmp_3
			_g_index++
			for {
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g_index
				if _hx_tmp_5 < HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Len() {
					_hx_tmp_4 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(HxMakeArrayView[any](HxMakeArrayView[*Hx_Obj_loreline_runtimestate](_g_map.Hx_Field__values)).Get(_g_index), ((any)(nil))))))
				} else {
					_hx_tmp_4 = false
				}

				if !(_hx_tmp_4) {
					break
				}

				_g_index++
			}

			var _g_key Hx_Typedef_haxe__int64___int64
			_ = _g_key
			_g_key = Hx_Field_loreline__node_nodeid_impl__fromTwoInt32(k1, k2)
			var _g_value *Hx_Obj_loreline_runtimestate
			_ = _g_value
			_g_value = v
			var id Hx_Typedef_haxe__int64___int64
			_ = id
			id = _g_key
			var state *Hx_Obj_loreline_runtimestate
			_ = state
			state = _g_value
			{
				var serialized any
				_ = serialized
				serialized = this.VTable.Hx_Field_serializeState(state)
				var declared bool
				_ = declared
				declared = (Hx_Field_loreline_objects_getFields(this, state.Hx_Field_originalFields).Len() > 0)
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if declared {
					_hx_tmp_6 = true
				} else {
					_hx_tmp_6 = (Hx_Field_reflect_fields(Hx_Field_go_haxe_hxdynamic_getField(serialized, "fields")).Len() > 0)
				}

				if _hx_tmp_6 {
					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = result
					Hx_Field_reflect_setField(_hx_tmp_7, Hx_Field_loreline__node_nodeid_impl__toString(id), serialized)
				}
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeBeatReference(beat *Hx_Obj_loreline_nbeatdecl) any {
	var path string
	_ = path
	path = beat.Hx_Field_name
	var parentBeat *Hx_Obj_loreline_nbeatdecl
	_ = parentBeat
	parentBeat = beat
	for {
		parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_596 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
			_ = _hx_closure_src_596
			_hx_closure_src_596 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getFirstParentOfType
			return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_233 any
				_ = hx_dyncast_233
				hx_dyncast_233 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_596(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_0
				if hx_dyncast_233 != nil {
					_hx_tmp_0 = hx_dyncast_233.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()((func() *Hx_Obj_loreline_node {
			_hx_up := parentBeat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
		if parentBeat != nil {
			path = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(parentBeat.Hx_Field_name, "."), path)
		}

		if !(parentBeat != nil) {
			break
		}
	}

	return any(map[string]any{"id": ((any)(Hx_Field_loreline__node_nodeid_impl__toString(beat.Hx_Field_id))), "path": ((any)(path))})
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeNodeReference(node *Hx_Obj_loreline_astnode) any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(Hx_Field_loreline__node_nodeid_impl__toString(node.Hx_Field_id)))
	return any(map[string]any{"id": _hx_tmp_0, "type": ((any)(node.VTable.Hx_Field_type()))})
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeCharacter(character *Hx_Obj_loreline_runtimecharacter) any {
	return this.VTable.Hx_Field_serializeFields(character.Hx_Field_fields, character.Hx_Field_originalFields)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeFields(fields any, originalFields__isOptional any) any {
	originalFields := HxDefault[any]()
	_ = originalFields

	if originalFields__isOptional != nil {
		originalFields = HxConvert[any](originalFields__isOptional)
	} else {
		originalFields = nil
	}
	{
		var _hx_reserved_type string
		_ = _hx_reserved_type
		_hx_reserved_type = HxStringNull
		var result any
		_ = result
		result = any(map[string]any{})
		if Hx_Field_std_isOfType(fields, Hx_Obj_loreline_fields_RTTI) {
			var cls *Hx_Obj_go_haxe_hxclass
			_ = cls
			cls = ((*Hx_Obj_go_haxe_hxclass)(((*Hx_Obj_go_haxe_hxclass)(Hx_Field_type_getClass(fields)))))
			if cls != nil {
				_hx_reserved_type = Hx_Field_type_getClassName(cls)
			}

			var hx_dyncast_234 any
			_ = hx_dyncast_234
			hx_dyncast_234 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_loreline_fields")
			var _hx_tmp_0 *Hx_Obj_loreline_fields
			_ = _hx_tmp_0
			if hx_dyncast_234 != nil {
				_hx_tmp_0 = hx_dyncast_234.(*Hx_Obj_loreline_fields)
			} else {
				_hx_tmp_0 = nil
			}

			var fieldMap *Hx_Obj_loreline_fields
			_ = fieldMap
			fieldMap = _hx_tmp_0
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[string]
				_ = _g1
				_g1 = fieldMap.VTable.Hx_Field_lorelineFields(this)
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = _g
					if !(_hx_tmp_1 < _g1.Len()) {
						break
					}

					var key string
					_ = key
					key = _g1.Get(_g)
					_g++
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					_hx_tmp_2 = result
					var _hx_tmp_3 string
					_ = _hx_tmp_3
					_hx_tmp_3 = key
					Hx_Field_reflect_setField(_hx_tmp_2, _hx_tmp_3, this.VTable.Hx_Field_serializeValue(fieldMap.VTable.Hx_Field_lorelineGet(this, key)))
				}
			}
		} else {
			if Hx_Field_std_isOfType(fields, Hx_Obj_haxe_ds_stringmap_RTTI) {
				var hx_dyncast_235 any
				_ = hx_dyncast_235
				hx_dyncast_235 = Hx_Field_go_haxe_hxdynamic_castClass(fields, "Hx_Obj_haxe_ds_stringmap")
				var _hx_tmp_4 *Hx_Obj_haxe_ds_stringmap
				_ = _hx_tmp_4
				if hx_dyncast_235 != nil {
					_hx_tmp_4 = hx_dyncast_235.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap()
				} else {
					_hx_tmp_4 = nil
				}

				var _hx_reserved_map *Hx_Obj_haxe_ds_stringmap
				_ = _hx_reserved_map
				_hx_reserved_map = _hx_tmp_4
				{
					var _gkey any
					_ = _gkey
					_gkey = _hx_reserved_map.VTable.Hx_Field_keys()
					for {
						var _hx_tmp_5 any
						_ = _hx_tmp_5
						_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "hasNext")
						if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any]().Dyn()))) {
							break
						}

						var _hx_tmp_6 any
						_ = _hx_tmp_6
						_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(_gkey, "next")
						var key string
						_ = key
						key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any]().Dyn()))
						var value any
						_ = value
						value = _hx_reserved_map.VTable.Hx_Field_get(key)
						var _hx_tmp_7 bool
						_ = _hx_tmp_7
						var _hx_tmp_8 bool
						_ = _hx_tmp_8
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(originalFields, nil))) {
							_hx_tmp_8 = true
						} else {
							_hx_tmp_8 = !Hx_Field_loreline_objects_fieldExists(this, originalFields, key)
						}

						if _hx_tmp_8 {
							_hx_tmp_7 = true
						} else {
							var _hx_tmp_9 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_9
							_hx_tmp_9 = this
							var _hx_tmp_10 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_10
							_hx_tmp_10 = _hx_tmp_9
							_hx_tmp_7 = !Hx_Field_loreline_equal_equal(_hx_tmp_10, Hx_Field_loreline_objects_getField(this, originalFields, key), value)
						}

						if _hx_tmp_7 {
							var _hx_tmp_11 any
							_ = _hx_tmp_11
							_hx_tmp_11 = result
							var _hx_tmp_12 string
							_ = _hx_tmp_12
							_hx_tmp_12 = key
							Hx_Field_reflect_setField(_hx_tmp_11, _hx_tmp_12, this.VTable.Hx_Field_serializeValue(value))
						}
					}
				}
			} else {
				var cls *Hx_Obj_go_haxe_hxclass
				_ = cls
				cls = ((*Hx_Obj_go_haxe_hxclass)(((*Hx_Obj_go_haxe_hxclass)(Hx_Field_type_getClass(fields)))))
				if cls != nil {
					_hx_reserved_type = Hx_Field_type_getClassName(cls)
				}

				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[string]
					_ = _g1
					_g1 = Hx_Field_reflect_fields(fields)
					for {
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = _g
						if !(_hx_tmp_13 < _g1.Len()) {
							break
						}

						var field string
						_ = field
						field = _g1.Get(_g)
						_g++
						var value any
						_ = value
						value = Hx_Field_reflect_getProperty(fields, field)
						var _hx_tmp_14 bool
						_ = _hx_tmp_14
						var _hx_tmp_15 bool
						_ = _hx_tmp_15
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(originalFields, nil))) {
							_hx_tmp_15 = true
						} else {
							_hx_tmp_15 = !Hx_Field_loreline_objects_fieldExists(this, originalFields, field)
						}

						if _hx_tmp_15 {
							_hx_tmp_14 = true
						} else {
							var _hx_tmp_16 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_16
							_hx_tmp_16 = this
							var _hx_tmp_17 *Hx_Obj_loreline_interpreter
							_ = _hx_tmp_17
							_hx_tmp_17 = _hx_tmp_16
							_hx_tmp_14 = !Hx_Field_loreline_equal_equal(_hx_tmp_17, Hx_Field_loreline_objects_getField(this, originalFields, field), value)
						}

						if _hx_tmp_14 {
							var _hx_tmp_18 any
							_ = _hx_tmp_18
							_hx_tmp_18 = result
							var _hx_tmp_19 string
							_ = _hx_tmp_19
							_hx_tmp_19 = field
							Hx_Field_reflect_setField(_hx_tmp_18, _hx_tmp_19, this.VTable.Hx_Field_serializeValue(value))
						}
					}
				}
			}
		}

		if _hx_reserved_type != HxStringNull {
			return any(map[string]any{"type": ((any)(_hx_reserved_type)), "fields": result})
		} else {
			return any(map[string]any{"fields": result})
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_serializeValue(value any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return nil
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		_hx_tmp_2 = true
	} else {
		_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
	}

	if _hx_tmp_0 {
		return value
	}

	if Hx_Field_loreline_arrays_isArray(value) {
		var arr HxArray[any]
		_ = arr
		arr = HxMakeArray[any]().Dyn()
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(value)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_3
				var _hx_tmp_4 HxArray[any]
				_ = _hx_tmp_4
				_hx_tmp_4 = HxMakeArrayView[any](arr)
				_ = (Hx_Array_Push[any](_hx_tmp_4, this.VTable.Hx_Field_serializeValue(Hx_Field_loreline_arrays_arrayGet(value, i))))
			}
		}

		return ((any)(arr))
	}

	var asBeat *Hx_Obj_loreline_nbeatdecl
	_ = asBeat
	asBeat = Hx_Field_loreline_runtimebeatref_beatOf(value)
	if asBeat != nil {
		var serializedRef any
		_ = serializedRef
		serializedRef = any(map[string]any{"type": ((any)("$beatRef")), "beat": this.VTable.Hx_Field_serializeBeatReference(asBeat)})
		if Hx_Field_std_isOfType(value, Hx_Obj_loreline_runtimebeatref_RTTI) {
			var hx_dyncast_236 any
			_ = hx_dyncast_236
			hx_dyncast_236 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_runtimebeatref")
			var _hx_tmp_5 *Hx_Obj_loreline_runtimebeatref
			_ = _hx_tmp_5
			if hx_dyncast_236 != nil {
				_hx_tmp_5 = hx_dyncast_236.(interface {
					Hx_As_Hx_Obj_loreline_runtimebeatref() *Hx_Obj_loreline_runtimebeatref
				}).Hx_As_Hx_Obj_loreline_runtimebeatref()
			} else {
				_hx_tmp_5 = nil
			}

			var ref *Hx_Obj_loreline_runtimebeatref
			_ = ref
			ref = _hx_tmp_5
			if ref.Hx_Field_stack.Len() > 0 {
				var cyclic bool
				_ = cyclic
				cyclic = false
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_runtimescope]
					_ = _g1
					_g1 = ref.Hx_Field_stack
					for {
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = _g
						if !(_hx_tmp_6 < _g1.Len()) {
							break
						}

						var captured *Hx_Obj_loreline_runtimescope
						_ = captured
						captured = _g1.Get(_g)
						_g++
						var _hx_tmp_7 HxArray[*Hx_Obj_loreline_runtimescope]
						_ = _hx_tmp_7
						_hx_tmp_7 = this.Hx_Field__serializingScopes
						var _hx_tmp_8 *Hx_Obj_loreline_runtimescope
						_ = _hx_tmp_8
						_hx_tmp_8 = captured
						if (Hx_Array_IndexOf[*Hx_Obj_loreline_runtimescope](_hx_tmp_7, _hx_tmp_8, HxNullable[int32]{})) != -1 {
							cyclic = true
							break
						}
					}
				}

				if !cyclic {
					var _hx_tmp_9 any
					_ = _hx_tmp_9
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field__saveInsertions, nil))) {
						_hx_tmp_9 = this.Hx_Field__saveInsertions
					} else {
						_hx_tmp_9 = any(map[string]any{})
					}

					var insertions any
					_ = insertions
					insertions = _hx_tmp_9
					var _g HxArray[any]
					_ = _g
					_g = HxMakeArray[any]().Dyn()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[*Hx_Obj_loreline_runtimescope]
						_ = _g2
						_g2 = ref.Hx_Field_stack
						for {
							var _hx_tmp_10 int32
							_ = _hx_tmp_10
							_hx_tmp_10 = _g1
							if !(_hx_tmp_10 < _g2.Len()) {
								break
							}

							var captured *Hx_Obj_loreline_runtimescope
							_ = captured
							captured = _g2.Get(_g1)
							_g1++
							var _hx_tmp_11 HxArray[any]
							_ = _hx_tmp_11
							_hx_tmp_11 = HxMakeArrayView[any](_g)
							_ = (Hx_Array_Push[any](_hx_tmp_11, this.VTable.Hx_Field_serializeScope(captured, insertions)))
						}
					}

					Hx_Field_go_haxe_hxdynamic_setField(serializedRef, "stack", _g)
				}
			}
		}

		return serializedRef
	}

	var asCharacter *Hx_Obj_loreline_runtimecharacterref
	_ = asCharacter
	asCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(value)
	if asCharacter != nil {
		var serializedCharacterRef any
		_ = serializedCharacterRef
		serializedCharacterRef = any(map[string]any{"type": ((any)("$characterRef")), "name": ((any)(asCharacter.Hx_Field_name))})
		return serializedCharacterRef
	}

	var _hx_tmp_12 any
	_ = _hx_tmp_12
	_hx_tmp_12 = value
	return this.VTable.Hx_Field_serializeFields(_hx_tmp_12, nil)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeNode(node *Hx_Obj_loreline_astnode, scopeLevel int32, next func()) {
	var lastLevel int32
	_ = lastLevel
	lastLevel = (this.Hx_Field_stack.Len() - 1)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = scopeLevel
	var isLastNode bool
	_ = isLastNode
	isLastNode = (_hx_tmp_0 >= this.Hx_Field_stack.Len())
	if !isLastNode {
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_597 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_597
			_hx_closure_src_597 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_597(((any)(o)))))
			}
		}()(node)
		if _g == nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from node within stack: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_605 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_605
				_hx_closure_src_605 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_605(((any)(o)))))
				}
			}()(node))))), node.Hx_Field_pos)))))
		} else {
			switch _g {
			case Hx_Obj_loreline_nalternative_RTTI:
				{
					this.VTable.Hx_Field_resumeAlternative(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
					}).Hx_As_Hx_Obj_loreline_nalternative(), scopeLevel, next)
				}

			case Hx_Obj_loreline_nbeatcall_RTTI:
				{
					this.VTable.Hx_Field_resumeBeatCall(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
					}).Hx_As_Hx_Obj_loreline_nbeatcall(), scopeLevel, next)
				}

			case Hx_Obj_loreline_nbeatdecl_RTTI:
				{
					this.VTable.Hx_Field_resumeBeatRun(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl(), scopeLevel, next)
				}

			case Hx_Obj_loreline_ncall_RTTI:
				{
					if this.VTable.Hx_Field_isBeatCall(node, scopeLevel) {
						this.VTable.Hx_Field_resumeCall(node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall(), scopeLevel, next)
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from node within stack: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
							var _hx_closure_src_606 func(o any) *Hx_Obj_go_haxe_hxclass
							_ = _hx_closure_src_606
							_hx_closure_src_606 = Hx_Field_type_getClass
							return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
								return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_606(((any)(o)))))
							}
						}()(node))))), node.Hx_Field_pos)))))
					}
				}

			case Hx_Obj_loreline_nchoiceoption_RTTI:
				{
					this.VTable.Hx_Field_resumeChoiceOption(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
					}).Hx_As_Hx_Obj_loreline_nchoiceoption(), scopeLevel, next)
				}

			case Hx_Obj_loreline_nchoicestatement_RTTI:
				{
					this.VTable.Hx_Field_resumeChoice(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
					}).Hx_As_Hx_Obj_loreline_nchoicestatement(), scopeLevel, next)
				}

			case Hx_Obj_loreline_nifstatement_RTTI:
				{
					this.VTable.Hx_Field_resumeIf(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
					}).Hx_As_Hx_Obj_loreline_nifstatement(), scopeLevel, next)
				}

			case Hx_Obj_loreline_nwhenstatement_RTTI:
				{
					this.VTable.Hx_Field_resumeWhen(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nwhenstatement() *Hx_Obj_loreline_nwhenstatement
					}).Hx_As_Hx_Obj_loreline_nwhenstatement(), scopeLevel, next)
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from node within stack: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
						var _hx_closure_src_607 func(o any) *Hx_Obj_go_haxe_hxclass
						_ = _hx_closure_src_607
						_hx_closure_src_607 = Hx_Field_type_getClass
						return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
							return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_607(((any)(o)))))
						}
					}()(node))))), node.Hx_Field_pos)))))
				}
			}
		}
	} else {
		var _g *Hx_Obj_go_haxe_hxclass
		_ = _g
		_g = func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_601 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_601
			_hx_closure_src_601 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_601(((any)(o)))))
			}
		}()(node)
		if _g == nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from last node: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
				var _hx_closure_src_608 func(o any) *Hx_Obj_go_haxe_hxclass
				_ = _hx_closure_src_608
				_hx_closure_src_608 = Hx_Field_type_getClass
				return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
					return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_608(((any)(o)))))
				}
			}()(node))))), node.Hx_Field_pos)))))
		} else {
			switch _g {
			case Hx_Obj_loreline_nalternative_RTTI:
				{
					this.VTable.Hx_Field_evalAlternative(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
					}).Hx_As_Hx_Obj_loreline_nalternative(), next)
				}

			case Hx_Obj_loreline_nbeatcall_RTTI:
				{
					this.VTable.Hx_Field_evalBeatCall(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
					}).Hx_As_Hx_Obj_loreline_nbeatcall(), next)
				}

			case Hx_Obj_loreline_ncall_RTTI:
				{
					var _hx_tmp_1 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_1
					_hx_tmp_1 = node
					if !this.VTable.Hx_Field_isBeatCall(_hx_tmp_1, nil) {
						this.VTable.Hx_Field_evalCall(node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall(), next)
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from last node: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
							var _hx_closure_src_609 func(o any) *Hx_Obj_go_haxe_hxclass
							_ = _hx_closure_src_609
							_hx_closure_src_609 = Hx_Field_type_getClass
							return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
								return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_609(((any)(o)))))
							}
						}()(node))))), node.Hx_Field_pos)))))
					}
				}

			case Hx_Obj_loreline_nchoicestatement_RTTI:
				{
					this.VTable.Hx_Field_evalChoice(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
					}).Hx_As_Hx_Obj_loreline_nchoicestatement(), next)
				}

			case Hx_Obj_loreline_ndialoguestatement_RTTI:
				{
					this.VTable.Hx_Field_evalDialogue(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
					}).Hx_As_Hx_Obj_loreline_ndialoguestatement(), next)
				}

			case Hx_Obj_loreline_ntextstatement_RTTI:
				{
					this.VTable.Hx_Field_evalText(node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
					}).Hx_As_Hx_Obj_loreline_ntextstatement(), next)
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Resume execution not supported from last node: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
						var _hx_closure_src_610 func(o any) *Hx_Obj_go_haxe_hxclass
						_ = _hx_closure_src_610
						_hx_closure_src_610 = Hx_Field_type_getClass
						return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
							return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_610(((any)(o)))))
						}
					}()(node))))), node.Hx_Field_pos)))))
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeNodeBody(node *Hx_Obj_loreline_astnode, scopeLevel int32, body HxArray[*Hx_Obj_loreline_astnode], next func()) {
	var currentScope *Hx_Obj_loreline_runtimescope
	_ = currentScope
	currentScope = this.Hx_Field_stack.Get(scopeLevel)
	var index int32
	_ = index
	index = 0
	var resumeIndex int32
	_ = resumeIndex
	resumeIndex = -1
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if currentScope.Hx_Field_head == nil {
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = scopeLevel
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 < (this.Hx_Field_stack.Len() - 1))
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot resume through a body with a no-headed scope that is not at the bottom of the stack", node.Hx_Field_pos)))))
	}

	if currentScope.Hx_Field_head != nil {
		var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_3
		_hx_tmp_3 = body
		var _hx_tmp_4 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_4
		_hx_tmp_4 = currentScope.Hx_Field_head
		index = (Hx_Array_IndexOf[*Hx_Obj_loreline_astnode](_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{}))
		resumeIndex = index
		if index == -1 {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resolve head in scope when resuming through body", node.Hx_Field_pos)))))
		}
	}

	var currentInsertion *Hx_Obj_loreline_runtimeinsertion
	_ = currentInsertion
	currentInsertion = this.VTable.Hx_Field_insertionAtOrBelow(scopeLevel)
	var moveNext func()
	_ = moveNext
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	moveNext = func() {
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if currentInsertion != nil {
			_hx_tmp_5 = currentInsertion.VTable.Hx_Field_isCollected()
		} else {
			_hx_tmp_5 = false
		}

		if _hx_tmp_5 {
			_gthis.VTable.Hx_Field_pop()
			next()
		} else {
			if (index != -1) && (index == resumeIndex) {
				var childNode *Hx_Obj_loreline_astnode
				_ = childNode
				childNode = body.Get(index)
				index = (index + 1)
				var done *Hx_Obj_loreline_evalnext
				_ = done
				done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
				_gthis.VTable.Hx_Field_resumeNode(childNode, (scopeLevel + 1), done.Hx_Field_cb)
				done.Hx_Field_sync = false
			} else {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = index
				if _hx_tmp_6 < body.Len() {
					var childNode *Hx_Obj_loreline_astnode
					_ = childNode
					childNode = body.Get(index)
					_gthis.Hx_Field_stack.Get(scopeLevel).Hx_Field_head = childNode
					index = (index + 1)
					var done *Hx_Obj_loreline_evalnext
					_ = done
					done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
					_gthis.VTable.Hx_Field_evalNode(childNode, done.Hx_Field_cb)
					done.Hx_Field_sync = false
				} else {
					_gthis.VTable.Hx_Field_pop()
					next()
				}
			}
		}
	}

	moveNext()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeBeatRun(beat *Hx_Obj_loreline_nbeatdecl, scopeLevel int32, next func()) {
	this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
		_hx_up := beat
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), scopeLevel, beat.Hx_Field_body, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeChoiceOption(option *Hx_Obj_loreline_nchoiceoption, scopeLevel int32, next func()) {
	this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
		_hx_up := option
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), scopeLevel, option.Hx_Field_body, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeChoice(choice *Hx_Obj_loreline_nchoicestatement, scopeLevel int32, next func()) {
	var currentScope *Hx_Obj_loreline_runtimescope
	_ = currentScope
	currentScope = this.Hx_Field_stack.Get(scopeLevel)
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	if currentScope.Hx_Field_head == nil {
		this.VTable.Hx_Field_evalChoice(choice, next)
	} else {
		if Hx_Field_std_isOfType(currentScope.Hx_Field_head, Hx_Obj_loreline_nchoiceoption_RTTI) {
			var option *Hx_Obj_loreline_nchoiceoption
			_ = option
			option = currentScope.Hx_Field_head.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
			}).Hx_As_Hx_Obj_loreline_nchoiceoption()
			var popThenNext func()
			_ = popThenNext
			popThenNext = func() {
				_gthis.VTable.Hx_Field_pop()
				next()
			}

			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = (scopeLevel + 1)
			if _hx_tmp_0 < this.Hx_Field_stack.Len() {
				this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
					_hx_up := option
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})(), (scopeLevel + 1), option.Hx_Field_body, popThenNext)
			} else {
				var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_1
				_hx_tmp_1 = currentScope.Hx_Field_beat
				var _hx_tmp_2 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_2
				_hx_tmp_2 = (func() *Hx_Obj_loreline_astnode {
					_hx_up := option
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})()
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
				_ = _hx_tmp_3
				_hx_tmp_3 = option.Hx_Field_body
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = nil
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = nil
				var _hx_tmp_6 any
				_ = _hx_tmp_6
				_hx_tmp_6 = nil
				this.VTable.Hx_Field_evalNodeBody(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, nil, popThenNext)
			}
		} else {
			if (currentScope.Hx_Field_insertion != nil) && (currentScope.Hx_Field_insertion.Hx_Field_kind == ((int32)(0))) {
				var insertionNode *Hx_Obj_loreline_astnode
				_ = insertionNode
				insertionNode = currentScope.Hx_Field_node
				var insertion *Hx_Obj_loreline_runtimeinsertion
				_ = insertion
				insertion = currentScope.Hx_Field_insertion
				if Hx_Field_std_isOfType(insertionNode, Hx_Obj_loreline_nbeatdecl_RTTI) {
					var beat *Hx_Obj_loreline_nbeatdecl
					_ = beat
					beat = insertionNode.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					var isRootPhase1 bool
					_ = isRootPhase1
					isRootPhase1 = true
					{
						var _g int32
						_ = _g
						_g = 0
						var _g1 int32
						_ = _g1
						_g1 = scopeLevel
						for _g < _g1 {
							var _hx_tmp_7 int32
							_ = _hx_tmp_7
							_hx_tmp_7 = _g
							_g = (_g + 1)
							var i int32
							_ = i
							i = _hx_tmp_7
							if this.Hx_Field_stack.Get(i).Hx_Field_insertion != nil {
								isRootPhase1 = false
								break
							}
						}
					}

					var _hx_tmp_8 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_8
					_hx_tmp_8 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := beat
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = scopeLevel
					var _hx_tmp_10 HxArray[*Hx_Obj_loreline_astnode]
					_ = _hx_tmp_10
					_hx_tmp_10 = beat.Hx_Field_body
					this.VTable.Hx_Field_resumeNodeBody(_hx_tmp_8, _hx_tmp_9, _hx_tmp_10, func() {
						var _hx_tmp_11 HxArray[*Hx_Obj_loreline_choiceoption]
						_ = _hx_tmp_11
						if insertion.Hx_Field_parentPartialOptions != nil {
							_hx_tmp_11 = insertion.Hx_Field_parentPartialOptions
						} else {
							_hx_tmp_11 = HxMakeArray[*Hx_Obj_loreline_choiceoption]()
						}

						var options HxArray[*Hx_Obj_loreline_choiceoption]
						_ = options
						options = _hx_tmp_11
						if insertion.Hx_Field_options != nil {
							var _g int32
							_ = _g
							_g = 0
							var _g1 HxArray[*Hx_Obj_loreline_choiceoption]
							_ = _g1
							_g1 = insertion.Hx_Field_options
							for {
								var _hx_tmp_12 int32
								_ = _hx_tmp_12
								_hx_tmp_12 = _g
								if !(_hx_tmp_12 < _g1.Len()) {
									break
								}

								var opt *Hx_Obj_loreline_choiceoption
								_ = opt
								opt = _g1.Get(_g)
								_g++
								_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](options, opt))
							}
						}

						var nextIndex int32
						_ = nextIndex
						nextIndex = insertion.Hx_Field_parentNextOptionIndex
						var parentBeat *Hx_Obj_loreline_nbeatdecl
						_ = parentBeat
						parentBeat = currentScope.Hx_Field_beat
						var optionsDone *Hx_Obj_loreline_evalnext
						_ = optionsDone
						optionsDone = _gthis.VTable.Hx_Field_wrapNext(func() {
							_gthis.VTable.Hx_Field_presentChoice(choice, options, next)
						})
						var _hx_tmp_13 *Hx_Obj_loreline_nbeatdecl
						_ = _hx_tmp_13
						_hx_tmp_13 = parentBeat
						var _hx_tmp_14 *Hx_Obj_loreline_nchoicestatement
						_ = _hx_tmp_14
						_hx_tmp_14 = choice
						var _hx_tmp_15 HxArray[*Hx_Obj_loreline_choiceoption]
						_ = _hx_tmp_15
						_hx_tmp_15 = options
						var _hx_tmp_16 func()
						_ = _hx_tmp_16
						_hx_tmp_16 = optionsDone.Hx_Field_cb
						_gthis.VTable.Hx_Field_evalChoiceOptionsAndInsertions(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15, _hx_tmp_16, HxNullable[int32]{Value: nextIndex, Valid: true}, isRootPhase1)
						optionsDone.Hx_Field_sync = false
					})
				} else {
					for {
						if !(this.Hx_Field_stack.Len() > scopeLevel) {
							break
						}

						this.VTable.Hx_Field_pop()
					}

					this.VTable.Hx_Field_evalChoice(choice, next)
				}
			} else {
				if Hx_Field_std_isOfType(currentScope.Hx_Field_node, Hx_Obj_loreline_nchoiceoption_RTTI) {
					var option *Hx_Obj_loreline_nchoiceoption
					_ = option
					option = currentScope.Hx_Field_node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
					}).Hx_As_Hx_Obj_loreline_nchoiceoption()
					this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
						_hx_up := option
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})(), scopeLevel, option.Hx_Field_body, next)
				} else {
					if Hx_Field_std_isOfType(currentScope.Hx_Field_node, Hx_Obj_loreline_nbeatdecl_RTTI) {
						var beat *Hx_Obj_loreline_nbeatdecl
						_ = beat
						beat = currentScope.Hx_Field_node.VTable.(interface {
							Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
						}).Hx_As_Hx_Obj_loreline_nbeatdecl()
						this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
							_hx_up := beat
							if _hx_up != nil {
								return &_hx_up.Hx_Obj_loreline_astnode
							}
							return nil
						})(), scopeLevel, beat.Hx_Field_body, next)
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Choice head is not a choice option", currentScope.Hx_Field_head.Hx_Field_pos)))))
					}
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeIf(ifStmt *Hx_Obj_loreline_nifstatement, scopeLevel int32, next func()) {
	var currentScope *Hx_Obj_loreline_runtimescope
	_ = currentScope
	currentScope = this.Hx_Field_stack.Get(scopeLevel)
	if currentScope.Hx_Field_head == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resolve head in scope when resuming through condition", ifStmt.Hx_Field_pos)))))
	}

	var _hx_tmp_0 HxArray[*Hx_Obj_loreline_astnode]
	_ = _hx_tmp_0
	_hx_tmp_0 = ifStmt.Hx_Field_thenBranch.Hx_Field_body
	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = currentScope.Hx_Field_head
	var isTrue bool
	_ = isTrue
	isTrue = ((Hx_Array_IndexOf[*Hx_Obj_loreline_astnode](_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) != -1)
	var _hx_tmp_2 *Hx_Obj_loreline_nblock
	_ = _hx_tmp_2
	if isTrue {
		_hx_tmp_2 = ifStmt.Hx_Field_thenBranch
	} else {
		_hx_tmp_2 = ifStmt.Hx_Field_elseBranch
	}

	var branch *Hx_Obj_loreline_nblock
	_ = branch
	branch = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if branch != nil {
		_hx_tmp_3 = (branch.Hx_Field_body.Len() > 0)
	} else {
		_hx_tmp_3 = false
	}

	if _hx_tmp_3 {
		this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
			_hx_up := branch
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})(), scopeLevel, branch.Hx_Field_body, next)
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resume condition: invalid scope", ifStmt.Hx_Field_pos)))))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeWhen(when *Hx_Obj_loreline_nwhenstatement, scopeLevel int32, next func()) {
	var currentScope *Hx_Obj_loreline_runtimescope
	_ = currentScope
	currentScope = this.Hx_Field_stack.Get(scopeLevel)
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	if Hx_Field_std_isOfType(currentScope.Hx_Field_head, Hx_Obj_loreline_nwhenrule_RTTI) {
		var rule *Hx_Obj_loreline_nwhenrule
		_ = rule
		rule = currentScope.Hx_Field_head.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
		}).Hx_As_Hx_Obj_loreline_nwhenrule()
		var popThenNext func()
		_ = popThenNext
		popThenNext = func() {
			_gthis.VTable.Hx_Field_pop()
			next()
		}

		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = (scopeLevel + 1)
		if _hx_tmp_0 < this.Hx_Field_stack.Len() {
			this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
				_hx_up := rule
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})(), (scopeLevel + 1), rule.Hx_Field_body, popThenNext)
		} else {
			this.VTable.Hx_Field_recordWhenRulePlayed(rule)
			var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_1
			_hx_tmp_1 = currentScope.Hx_Field_beat
			var _hx_tmp_2 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := rule
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_3
			_hx_tmp_3 = rule.Hx_Field_body
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = nil
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = nil
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = nil
			this.VTable.Hx_Field_evalNodeBody(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, nil, popThenNext)
		}

		return
	}

	var _hx_tmp_7 bool
	_ = _hx_tmp_7
	if (currentScope.Hx_Field_insertion != nil) && (currentScope.Hx_Field_insertion.Hx_Field_kind == ((int32)(1))) {
		_hx_tmp_7 = Hx_Field_std_isOfType(currentScope.Hx_Field_node, Hx_Obj_loreline_nbeatdecl_RTTI)
	} else {
		_hx_tmp_7 = false
	}

	if _hx_tmp_7 {
		var insertion *Hx_Obj_loreline_runtimeinsertion
		_ = insertion
		insertion = currentScope.Hx_Field_insertion
		var beat *Hx_Obj_loreline_nbeatdecl
		_ = beat
		beat = currentScope.Hx_Field_node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
		}).Hx_As_Hx_Obj_loreline_nbeatdecl()
		var owner *Hx_Obj_loreline_runtimeinsertion
		_ = owner
		owner = this.VTable.Hx_Field_collectingWhenInsertion((scopeLevel - 1))
		var _hx_tmp_8 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_8
		_hx_tmp_8 = (func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
		var _hx_tmp_9 int32
		_ = _hx_tmp_9
		_hx_tmp_9 = scopeLevel
		var _hx_tmp_10 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_10
		_hx_tmp_10 = beat.Hx_Field_body
		this.VTable.Hx_Field_resumeNodeBody(_hx_tmp_8, _hx_tmp_9, _hx_tmp_10, func() {
			var _hx_tmp_11 HxArray[any]
			_ = _hx_tmp_11
			if insertion.Hx_Field_parentPartialRules != nil {
				_hx_tmp_11 = insertion.Hx_Field_parentPartialRules
			} else {
				_hx_tmp_11 = HxMakeArray[any]().Dyn()
			}

			var candidates HxArray[any]
			_ = candidates
			candidates = _hx_tmp_11
			if insertion.Hx_Field_rules != nil {
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[any]
				_ = _g1
				_g1 = insertion.Hx_Field_rules
				for {
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					_hx_tmp_12 = _g
					if !(_hx_tmp_12 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
						break
					}

					var candidate any
					_ = candidate
					candidate = _g1.Get(_g)
					_g++
					_ = (Hx_Array_Push[any](HxMakeArrayView[any](candidates), candidate))
				}
			}

			var _hx_tmp_13 bool
			_ = _hx_tmp_13
			if owner != nil {
				_hx_tmp_13 = owner.Hx_Field_whenFirst
			} else {
				var _hx_tmp_14 bool
				_ = _hx_tmp_14
				if (when.Hx_Field_strategy).Valid == true {
					_hx_tmp_14 = ((when.Hx_Field_strategy).Value == "first")
				} else {
					_hx_tmp_14 = false
				}

				_hx_tmp_13 = (_hx_tmp_14)
			}

			var stopAtFirst bool
			_ = stopAtFirst
			stopAtFirst = _hx_tmp_13
			var done *Hx_Obj_loreline_evalnext
			_ = done
			done = _gthis.VTable.Hx_Field_wrapNext(func() {
				_gthis.VTable.Hx_Field_finishWhenCollection(when, candidates, owner, next)
			})
			_gthis.VTable.Hx_Field_collectWhenRules(when, candidates, insertion.Hx_Field_parentNextRuleIndex, stopAtFirst, owner, done.Hx_Field_cb)
			done.Hx_Field_sync = false
		})
		return
	}

	if Hx_Field_std_isOfType(currentScope.Hx_Field_node, Hx_Obj_loreline_nbeatdecl_RTTI) {
		var beat *Hx_Obj_loreline_nbeatdecl
		_ = beat
		beat = currentScope.Hx_Field_node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
		}).Hx_As_Hx_Obj_loreline_nbeatdecl()
		this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})(), scopeLevel, beat.Hx_Field_body, next)
		return
	}

	if currentScope.Hx_Field_head == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resolve head when resuming when block", when.Hx_Field_pos)))))
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nwhenrule]
		_ = _g1
		_g1 = when.Hx_Field_rules
		for {
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = _g
			if !(_hx_tmp_15 < _g1.Len()) {
				break
			}

			var rule *Hx_Obj_loreline_nwhenrule
			_ = rule
			rule = _g1.Get(_g)
			_g++
			var _hx_tmp_16 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_16
			_hx_tmp_16 = rule.Hx_Field_body
			var _hx_tmp_17 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_17
			_hx_tmp_17 = currentScope.Hx_Field_head
			if (Hx_Array_IndexOf[*Hx_Obj_loreline_astnode](_hx_tmp_16, _hx_tmp_17, HxNullable[int32]{})) != -1 {
				this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
					_hx_up := rule
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})(), scopeLevel, rule.Hx_Field_body, next)
				return
			}
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resume when block: head not found in any rule", when.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeAlternative(alt *Hx_Obj_loreline_nalternative, scopeLevel int32, next func()) {
	var currentScope *Hx_Obj_loreline_runtimescope
	_ = currentScope
	currentScope = this.Hx_Field_stack.Get(scopeLevel)
	if currentScope.Hx_Field_head == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resolve head when resuming alternative", alt.Hx_Field_pos)))))
	}

	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
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
			var _hx_tmp_1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_1
			_hx_tmp_1 = item.Hx_Field_body
			var _hx_tmp_2 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_2
			_hx_tmp_2 = currentScope.Hx_Field_head
			if (Hx_Array_IndexOf[*Hx_Obj_loreline_astnode](_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{})) != -1 {
				var order HxArray[int32]
				_ = order
				order = currentScope.Hx_Field_shuffle
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				if order != nil {
					var _hx_tmp_4 HxArray[int32]
					_ = _hx_tmp_4
					_hx_tmp_4 = order
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = i
					var _hx_tmp_6 HxArray[int32]
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_4
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _hx_tmp_5
					_hx_tmp_3 = (Hx_Array_IndexOf[int32](_hx_tmp_6, _hx_tmp_7, HxNullable[int32]{}))
				} else {
					_hx_tmp_3 = -1
				}

				var position int32
				_ = position
				position = _hx_tmp_3
				if (alt.Hx_Field_mode == ((int32)(4))) && (position != -1) {
					var _hx_tmp_8 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_8
					_hx_tmp_8 = (func() *Hx_Obj_loreline_astnode {
						_hx_up := item
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})()
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = scopeLevel
					var _hx_tmp_10 HxArray[*Hx_Obj_loreline_astnode]
					_ = _hx_tmp_10
					_hx_tmp_10 = item.Hx_Field_body
					this.VTable.Hx_Field_resumeNodeBody(_hx_tmp_8, _hx_tmp_9, _hx_tmp_10, func() {
						_gthis.VTable.Hx_Field_evalShuffledItems(alt, order, (position + 1), next)
					})
				} else {
					this.VTable.Hx_Field_resumeNodeBody((func() *Hx_Obj_loreline_astnode {
						_hx_up := item
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})(), scopeLevel, item.Hx_Field_body, next)
				}

				return
			}
		}
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Failed to resume alternative: head not found in any item", alt.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeCall(call *Hx_Obj_loreline_ncall, scopeLevel int32, next func()) {
	var resolvedRef *Hx_Obj_loreline_runtimebeatref
	_ = resolvedRef
	resolvedRef = this.VTable.Hx_Field_resolveBeatRefFromCall(call)
	if resolvedRef != nil {
		this.VTable.Hx_Field_resumeBeatRun(resolvedRef.Hx_Field_beat, scopeLevel, next)
		return
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot resume through a function call that is not at the bottom of the stack", call.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resumeBeatCall(call *Hx_Obj_loreline_nbeatcall, scopeLevel int32, next func()) {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if scopeLevel >= 0 {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = scopeLevel
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (_hx_tmp_3 < this.Hx_Field_stack.Len())
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = Hx_Field_std_isOfType(this.Hx_Field_stack.Get(scopeLevel).Hx_Field_node, Hx_Obj_loreline_nbeatdecl_RTTI)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		this.VTable.Hx_Field_resumeBeatRun(this.Hx_Field_stack.Get(scopeLevel).Hx_Field_node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
		}).Hx_As_Hx_Obj_loreline_nbeatdecl(), scopeLevel, next)
		return
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot resume through a beat call that is not at the bottom of the stack", call.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreStack(savedStack HxArray[any], savedInsertions__isOptional any, restoredInsertions__isOptional any) bool {
	savedInsertions := HxDefault[any]()
	_ = savedInsertions

	if savedInsertions__isOptional != nil {
		savedInsertions = HxConvert[any](savedInsertions__isOptional)
	} else {
		savedInsertions = nil
	}
	restoredInsertions := HxDefault[Hx_Typedef_map]()
	_ = restoredInsertions

	if restoredInsertions__isOptional != nil {
		restoredInsertions = HxConvert[Hx_Typedef_map](restoredInsertions__isOptional)
	} else {
		restoredInsertions = nil
	}
	{
		var result HxArray[*Hx_Obj_loreline_runtimescope]
		_ = result
		result = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		if restoredInsertions == nil {
			restoredInsertions = (func() *Hx_Obj_haxe_imap {
				_hx_icast_237 := Hx_Obj_haxe_ds_intmap_CreateInstance()
				if _hx_icast_237 == nil {
					return nil
				}
				Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_237, "Hx_Obj_haxe_imap")
				return &Hx_Obj_haxe_imap{VTable: _hx_icast_237.VTable.(Hx_Obj_VTable_haxe_imap)}
			})()
		}

		var i int32
		_ = i
		i = (Hx_Field_go_haxe_hxdynamic_getArrayLength(savedStack) - 1)
		for i >= 0 {
			var scope *Hx_Obj_loreline_runtimescope
			_ = scope
			scope = this.VTable.Hx_Field_restoreScopeItem(savedStack.Get(i), savedInsertions, restoredInsertions)
			if scope == nil {
				return false
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](result, scope))
			i--
		}

		i = (result.Len() - 1)
		for i >= 0 {
			this.VTable.Hx_Field_push(result.Get(i))
			i--
		}

		return true
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreScopeItem(savedScope any, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_runtimescope {
	var beat *Hx_Obj_loreline_nbeatdecl
	_ = beat
	beat = this.VTable.Hx_Field_restoreBeat(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "beat"))
	if beat == nil {
		return nil
	}

	var savedBeatId Hx_Typedef_haxe__int64___int64
	_ = savedBeatId
	savedBeatId = Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "beat"), "id")))
	var savedNode any
	_ = savedNode
	savedNode = Hx_Field_go_haxe_hxdynamic_getField(savedScope, "node")
	var _hx_tmp_0 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedNode, nil))) {
		_hx_tmp_0 = this.VTable.Hx_Field_restoreNode(savedNode, savedBeatId, beat)
	} else {
		_hx_tmp_0 = nil
	}

	var node *Hx_Obj_loreline_astnode
	_ = node
	node = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedNode, nil))) {
		_hx_tmp_1 = (node == nil)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		return nil
	}

	var beats HxArray[*Hx_Obj_loreline_nbeatdecl]
	_ = beats
	beats = HxMakeArray[*Hx_Obj_loreline_nbeatdecl]()
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "beats"), ((any)(nil))))) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "beats"))
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var savedBeat any
			_ = savedBeat
			savedBeat = _g1.Get(_g)
			_g++
			var beatInScope *Hx_Obj_loreline_nbeatdecl
			_ = beatInScope
			beatInScope = this.VTable.Hx_Field_restoreBeat(savedBeat)
			if beatInScope == nil {
				return nil
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatdecl](beats, beatInScope))
		}
	}

	var savedState any
	_ = savedState
	savedState = Hx_Field_go_haxe_hxdynamic_getField(savedScope, "state")
	var _hx_tmp_3 *Hx_Obj_loreline_runtimestate
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedState, nil))) {
		_hx_tmp_3 = this.VTable.Hx_Field_restoreState(nil, Hx_Field_go_haxe_hxdynamic_getField(savedScope, "state"))
	} else {
		_hx_tmp_3 = nil
	}

	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = _hx_tmp_3
	var savedHead any
	_ = savedHead
	savedHead = Hx_Field_go_haxe_hxdynamic_getField(savedScope, "head")
	var _hx_tmp_4 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_4
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedHead, nil))) {
		_hx_tmp_4 = this.VTable.Hx_Field_restoreNode(savedHead, savedBeatId, beat)
	} else {
		_hx_tmp_4 = nil
	}

	var head *Hx_Obj_loreline_astnode
	_ = head
	head = _hx_tmp_4
	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedHead, nil))) {
		_hx_tmp_5 = (head == nil)
	} else {
		_hx_tmp_5 = false
	}

	if _hx_tmp_5 {
		return nil
	}

	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	var _hx_tmp_6 bool
	_ = _hx_tmp_6
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "insertion"), nil)))))), ((any)(false))))) {
		_hx_tmp_6 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedInsertions, nil))))
	} else {
		_hx_tmp_6 = false
	}

	if _hx_tmp_6 {
		insertion = this.VTable.Hx_Field_restoreInsertion(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "insertion")))), savedInsertions, restoredInsertions)
	}

	var captured HxArray[*Hx_Obj_loreline_runtimescope]
	_ = captured
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "captured"), ((any)(nil))))) {
		captured = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "captured"))
			for {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = _g
				if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var savedCaptured any
				_ = savedCaptured
				savedCaptured = _g1.Get(_g)
				_g++
				var capturedScope *Hx_Obj_loreline_runtimescope
				_ = capturedScope
				capturedScope = this.VTable.Hx_Field_restoreScopeItem(savedCaptured, savedInsertions, restoredInsertions)
				if capturedScope == nil {
					return nil
				}

				_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](captured, capturedScope))
			}
		}
	}

	var shuffle HxArray[int32]
	_ = shuffle
	var _hx_tmp_8 bool
	_ = _hx_tmp_8
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "shuffle"), ((any)(nil))))) {
		_hx_tmp_8 = Hx_Field_std_isOfType(node, Hx_Obj_loreline_nblock_RTTI)
	} else {
		_hx_tmp_8 = false
	}

	if _hx_tmp_8 {
		var _hx_tmp_9 HxArray[int32]
		_ = _hx_tmp_9
		_hx_tmp_9 = HxAnyToTypedArray[int32](Hx_Field_go_haxe_hxdynamic_getField(savedScope, "shuffle"))
		shuffle = this.VTable.Hx_Field_restoreShuffleOrder(_hx_tmp_9, node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_nblock() *Hx_Obj_loreline_nblock
		}).Hx_As_Hx_Obj_loreline_nblock())
	}

	return Hx_Obj_loreline_runtimescope_CreateInstance(nil, beat, node, beats, state, head, insertion, captured, shuffle)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreInsertion(insertionId int32, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_runtimeinsertion {
	if func() func(_hx_param_key int32) bool {
		var _hx_closure_src_611 func(_hx_param_key any) bool
		_ = _hx_closure_src_611
		_hx_closure_src_611 = restoredInsertions.VTable.Hx_Field_exists
		return func(_hx_param_key int32) bool {
			return _hx_closure_src_611(((any)(_hx_param_key)))
		}
	}()(insertionId) {
		return func() func(_hx_param_key int32) *Hx_Obj_loreline_runtimeinsertion {
			var _hx_closure_src_612 func(_hx_param_key any) any
			_ = _hx_closure_src_612
			_hx_closure_src_612 = restoredInsertions.VTable.Hx_Field_get
			return func(_hx_param_key int32) *Hx_Obj_loreline_runtimeinsertion {
				var hx_dyncast_238 any
				_ = hx_dyncast_238
				hx_dyncast_238 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_612(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimeinsertion")
				var _hx_tmp_0 *Hx_Obj_loreline_runtimeinsertion
				_ = _hx_tmp_0
				if hx_dyncast_238 != nil {
					_hx_tmp_0 = hx_dyncast_238.(interface {
						Hx_As_Hx_Obj_loreline_runtimeinsertion() *Hx_Obj_loreline_runtimeinsertion
					}).Hx_As_Hx_Obj_loreline_runtimeinsertion()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(insertionId)
	}

	var key string
	_ = key
	key = Hx_Field_std_string(insertionId)
	var saved any
	_ = saved
	saved = Hx_Field_reflect_field(savedInsertions, key)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(saved, nil))) {
		return nil
	}

	var origin *Hx_Obj_loreline_ninsertion
	_ = origin
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "origin"), nil))) {
		var nodeId Hx_Typedef_haxe__int64___int64
		_ = nodeId
		nodeId = Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(saved, "origin"), "id")))
		var node *Hx_Obj_loreline_node
		_ = node
		node = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeById(nodeId)
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if node != nil {
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = ((any)(node.VTable.Hx_Field_type()))
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(saved, "origin"), "type")))))
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			origin = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion
			}).Hx_As_Hx_Obj_loreline_ninsertion()
		}
	}

	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	insertion = Hx_Obj_loreline_runtimeinsertion_CreateInstance(insertionId, origin)
	func() func(_hx_param_key int32, _hx_param_value *Hx_Obj_loreline_runtimeinsertion) {
		var _hx_closure_src_613 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_613
		_hx_closure_src_613 = restoredInsertions.VTable.Hx_Field_set
		return func(_hx_param_key int32, _hx_param_value *Hx_Obj_loreline_runtimeinsertion) {
			_hx_closure_src_613(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(insertionId, insertion)
	if insertionId >= this.Hx_Field_context.Hx_Field_nextInsertionId {
		this.Hx_Field_context.Hx_Field_nextInsertionId = (insertionId + 1)
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "options"), ((any)(nil))))) {
		insertion.Hx_Field_options = HxMakeArray[*Hx_Obj_loreline_choiceoption]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "options"))
			for {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var savedOpt any
				_ = savedOpt
				savedOpt = _g1.Get(_g)
				_g++
				var opt *Hx_Obj_loreline_choiceoption
				_ = opt
				opt = this.VTable.Hx_Field_restoreChoiceOption(savedOpt, savedInsertions, restoredInsertions)
				if opt != nil {
					_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](insertion.Hx_Field_options, opt))
				}
			}
		}
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "stack"), ((any)(nil))))) {
		insertion.Hx_Field_stack = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "stack"))
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g
				if !(_hx_tmp_5 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var savedScope any
				_ = savedScope
				savedScope = _g1.Get(_g)
				_g++
				var scope *Hx_Obj_loreline_runtimescope
				_ = scope
				scope = this.VTable.Hx_Field_restoreScopeItem(savedScope, savedInsertions, restoredInsertions)
				if scope != nil {
					_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](insertion.Hx_Field_stack, scope))
				}
			}
		}
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentPartialOptions"), ((any)(nil))))) {
		insertion.Hx_Field_parentPartialOptions = HxMakeArray[*Hx_Obj_loreline_choiceoption]()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentPartialOptions"))
			for {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = _g
				if !(_hx_tmp_6 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var savedOpt any
				_ = savedOpt
				savedOpt = _g1.Get(_g)
				_g++
				var opt *Hx_Obj_loreline_choiceoption
				_ = opt
				opt = this.VTable.Hx_Field_restoreChoiceOption(savedOpt, savedInsertions, restoredInsertions)
				if opt != nil {
					_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](insertion.Hx_Field_parentPartialOptions, opt))
				}
			}
		}

		var _hx_tmp_7 int32
		_ = _hx_tmp_7
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentNextOptionIndex"), nil)))))), ((any)(false))))) {
			_hx_tmp_7 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentNextOptionIndex"))))
		} else {
			_hx_tmp_7 = 0
		}

		insertion.Hx_Field_parentNextOptionIndex = _hx_tmp_7
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(saved, "kind"), ((any)("when"))))) {
		insertion.Hx_Field_kind = ((int32)(1))
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "rules"), ((any)(nil))))) {
		insertion.Hx_Field_rules = this.VTable.Hx_Field_restoreWhenCandidates(HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "rules")), savedInsertions, restoredInsertions)
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentPartialRules"), ((any)(nil))))) {
		insertion.Hx_Field_parentPartialRules = this.VTable.Hx_Field_restoreWhenCandidates(HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentPartialRules")), savedInsertions, restoredInsertions)
		var _hx_tmp_8 int32
		_ = _hx_tmp_8
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentNextRuleIndex"), nil)))))), ((any)(false))))) {
			_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(saved, "parentNextRuleIndex"))))
		} else {
			_hx_tmp_8 = 0
		}

		insertion.Hx_Field_parentNextRuleIndex = _hx_tmp_8
	}

	insertion.Hx_Field_whenFirst = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(saved, "whenFirst"), ((any)(true))))))
	return insertion
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreWhenCandidates(saved HxArray[any], savedInsertions any, restoredInsertions Hx_Typedef_map) HxArray[any] {
	var result HxArray[any]
	_ = result
	result = HxMakeArray[any]().Dyn()
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(saved)) {
				break
			}

			var savedRule any
			_ = savedRule
			savedRule = saved.Get(_g)
			_g++
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "node"), nil))) {
				_hx_tmp_1 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeById(Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "node"), "id"))))
			} else {
				_hx_tmp_1 = nil
			}

			var node *Hx_Obj_loreline_node
			_ = node
			node = _hx_tmp_1
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if node == nil {
				_hx_tmp_3 = true
			} else {
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = ((any)(node.VTable.Hx_Field_type()))
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_4
				_hx_tmp_3 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_5, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "node"), "type")))))
			}

			if _hx_tmp_3 {
				_hx_tmp_2 = true
			} else {
				_hx_tmp_2 = !Hx_Field_std_isOfType(node, Hx_Obj_loreline_nwhenrule_RTTI)
			}

			if _hx_tmp_2 {
				continue
			}

			var _hx_tmp_6 HxArray[any]
			_ = _hx_tmp_6
			_hx_tmp_6 = HxMakeArrayView[any](result)
			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = ((any)(node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
			}).Hx_As_Hx_Obj_loreline_nwhenrule()))
			var _hx_tmp_8 any
			_ = _hx_tmp_8
			_hx_tmp_8 = ((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "eligible"), ((any)(true))))))))
			var _hx_tmp_9 *Hx_Obj_loreline_runtimeinsertion
			_ = _hx_tmp_9
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "insertion"), nil)))))), ((any)(false))))) {
				_hx_tmp_9 = this.VTable.Hx_Field_restoreInsertion(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(savedRule, "insertion")))), savedInsertions, restoredInsertions)
			} else {
				_hx_tmp_9 = nil
			}

			_ = (Hx_Array_Push[any](_hx_tmp_6, any(map[string]any{"rule": _hx_tmp_7, "eligible": _hx_tmp_8, "insertion": ((any)(_hx_tmp_9))})))
		}
	}

	return result
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreChoiceOption(saved any, savedInsertions any, restoredInsertions Hx_Typedef_map) *Hx_Obj_loreline_choiceoption {
	var node *Hx_Obj_loreline_nchoiceoption
	_ = node
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "node"), nil))) {
		var nodeId Hx_Typedef_haxe__int64___int64
		_ = nodeId
		nodeId = Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(saved, "node"), "id")))
		var astNode *Hx_Obj_loreline_node
		_ = astNode
		astNode = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeById(nodeId)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if astNode != nil {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = ((any)(astNode.VTable.Hx_Field_type()))
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = _hx_tmp_1
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(saved, "node"), "type")))))
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			node = astNode.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
			}).Hx_As_Hx_Obj_loreline_nchoiceoption()
		}
	}

	var tags HxArray[*Hx_Obj_loreline_texttag]
	_ = tags
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "tags"), ((any)(nil))))) {
		var _g HxArray[*Hx_Obj_loreline_texttag]
		_ = _g
		_g = HxMakeArray[*Hx_Obj_loreline_texttag]()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[any]
			_ = _g2
			_g2 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(saved, "tags"))
			for {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g1
				if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
					break
				}

				var t any
				_ = t
				t = _g2.Get(_g1)
				_g1++
				var _hx_tmp_4 HxArray[*Hx_Obj_loreline_texttag]
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_ = (Hx_Array_Push[*Hx_Obj_loreline_texttag](_hx_tmp_4, Hx_Obj_loreline_texttag_CreateInstance(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(t, "closing"), ((any)(true)))))), Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(t, "value")), ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(t, "offset")))))))
			}
		}

		tags = _g
	}

	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "insertion"), nil)))))), ((any)(false))))) {
		_hx_tmp_5 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(savedInsertions, nil))))
	} else {
		_hx_tmp_5 = false
	}

	if _hx_tmp_5 {
		insertion = this.VTable.Hx_Field_restoreInsertion(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(saved, "insertion")))), savedInsertions, restoredInsertions)
	}

	return Hx_Obj_loreline_choiceoption_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(saved, "text")), tags, ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(saved, "disabled"), ((any)(true)))))), node, insertion)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreBeatToResume(savedStack HxArray[any]) *Hx_Obj_loreline_nbeatdecl {
	var beat *Hx_Obj_loreline_nbeatdecl
	_ = beat
	if Hx_Field_go_haxe_hxdynamic_getArrayLength(savedStack) > 0 {
		var savedScope any
		_ = savedScope
		savedScope = savedStack.Get(0)
		beat = this.VTable.Hx_Field_restoreBeat(Hx_Field_go_haxe_hxdynamic_getField(savedScope, "beat"))
		if beat != nil {
			var parentBeat *Hx_Obj_loreline_nbeatdecl
			_ = parentBeat
			parentBeat = beat
			for {
				parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
					var _hx_closure_src_614 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
					_ = _hx_closure_src_614
					_hx_closure_src_614 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getFirstParentOfType
					return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
						var hx_dyncast_239 any
						_ = hx_dyncast_239
						hx_dyncast_239 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_614(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
						var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
						_ = _hx_tmp_0
						if hx_dyncast_239 != nil {
							_hx_tmp_0 = hx_dyncast_239.(interface {
								Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
							}).Hx_As_Hx_Obj_loreline_nbeatdecl()
						} else {
							_hx_tmp_0 = nil
						}

						return _hx_tmp_0
					}
				}()((func() *Hx_Obj_loreline_node {
					_hx_up := parentBeat
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
				if parentBeat != nil {
					beat = parentBeat
				}

				if !(parentBeat != nil) {
					break
				}
			}
		}
	}

	return beat
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreNode(savedNode any, savedBeatId Hx_Typedef_haxe__int64___int64, beat *Hx_Obj_loreline_nbeatdecl) *Hx_Obj_loreline_astnode {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(((Hx_Typedef_haxe__int64___int64)(beat.Hx_Field_id))))
	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)((_hx_tmp_2 >> int64(((int32)(32))))))
	var high int32
	_ = high
	high = ((_hx_tmp_1 >> 16) & 65535)
	var _hx_tmp_3 int64
	_ = _hx_tmp_3
	_hx_tmp_3 = ((int64)(((Hx_Typedef_haxe__int64___int64)(savedBeatId))))
	var _hx_tmp_5 int64
	_ = _hx_tmp_5
	_hx_tmp_5 = _hx_tmp_3
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = ((int32)((_hx_tmp_5 >> int64(((int32)(32))))))
	var high1 int32
	_ = high1
	high1 = ((_hx_tmp_4 >> 16) & 65535)
	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	if high >= 32768 {
		_hx_tmp_6 = (high - 65536)
	} else {
		_hx_tmp_6 = high
	}

	var _hx_tmp_7 int32
	_ = _hx_tmp_7
	if high1 >= 32768 {
		_hx_tmp_7 = (high1 - 65536)
	} else {
		_hx_tmp_7 = high1
	}

	var sectionOffset int32
	_ = sectionOffset
	sectionOffset = (((_hx_tmp_6) + 32768) - ((_hx_tmp_7) + 32768))
	var nodeId Hx_Typedef_haxe__int64___int64
	_ = nodeId
	nodeId = Hx_Field_loreline__node_nodeid_impl__fromString(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(savedNode, "id")))
	if Hx_Field_loreline__node_nodeid_impl__equals(nodeId, savedBeatId) {
		return (func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
	}

	{
		var _hx_tmp_8 int64
		_ = _hx_tmp_8
		_hx_tmp_8 = ((int64)(((Hx_Typedef_haxe__int64___int64)(nodeId))))
		var _hx_tmp_10 int64
		_ = _hx_tmp_10
		_hx_tmp_10 = _hx_tmp_8
		var _hx_tmp_9 int32
		_ = _hx_tmp_9
		_hx_tmp_9 = ((int32)((_hx_tmp_10 >> int64(((int32)(32))))))
		var high2 int32
		_ = high2
		high2 = ((_hx_tmp_9 >> 16) & 65535)
		var _hx_tmp_11 int32
		_ = _hx_tmp_11
		if high2 >= 32768 {
			_hx_tmp_11 = (high2 - 65536)
		} else {
			_hx_tmp_11 = high2
		}

		var section int32
		_ = section
		section = (((_hx_tmp_11) + 32768) + sectionOffset)
		var value int32
		_ = value
		value = (section - 32768)
		if value < 0 {
			value = ((value + 65536) & 65535)
		}

		var _hx_tmp_12 int64
		_ = _hx_tmp_12
		_hx_tmp_12 = ((int64)(((Hx_Typedef_haxe__int64___int64)(nodeId))))
		var h int64
		_ = h
		h = int64(((int32)(((((int32)(((int32)((_hx_tmp_12 >> int64(((int32)(32)))))))) & 65535) | (value << 16)))))
		var l int64
		_ = l
		l = int64(((int32)(((Hx_Typedef_haxe__int64___int64)(nodeId)))))
		var _hx_tmp_13 int64
		_ = _hx_tmp_13
		_hx_tmp_13 = h
		nodeId = ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_13 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
	}

	var node *Hx_Obj_loreline_node
	_ = node
	node = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeById(nodeId)
	var _hx_tmp_14 bool
	_ = _hx_tmp_14
	if node != nil {
		var _hx_tmp_15 any
		_ = _hx_tmp_15
		_hx_tmp_15 = ((any)(node.VTable.Hx_Field_type()))
		var _hx_tmp_16 any
		_ = _hx_tmp_16
		_hx_tmp_16 = _hx_tmp_15
		_hx_tmp_14 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_16, Hx_Field_go_haxe_hxdynamic_getField(savedNode, "type")))))
	} else {
		_hx_tmp_14 = false
	}

	if _hx_tmp_14 {
		return node.VTable.(interface {
			Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
		}).Hx_As_Hx_Obj_loreline_astnode()
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreBeat(beatRef any) *Hx_Obj_loreline_nbeatdecl {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(beatRef, "path"))
	return this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_findBeatByPathFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
		_hx_up := this.Hx_Field_context.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})())
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreState(state *Hx_Obj_loreline_runtimestate, data any) *Hx_Obj_loreline_runtimestate {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	if state != nil {
		_hx_tmp_0 = state.Hx_Field_fields
	} else {
		_hx_tmp_0 = nil
	}

	var target any
	_ = target
	target = _hx_tmp_0
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(target, nil))) {
		_hx_tmp_1 = target
	} else {
		_hx_tmp_1 = nil
	}

	var fields any
	_ = fields
	fields = this.VTable.Hx_Field_restoreFields(_hx_tmp_1, data)
	if state == nil {
		state = Hx_Obj_loreline_runtimestate_CreateInstance(this, nil, fields, nil)
	}

	return state
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreCharacters(data any) {
	var _g int32
	_ = _g
	_g = 0
	var _g1 HxArray[string]
	_ = _g1
	_g1 = Hx_Field_reflect_fields(data)
	for {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = _g
		if !(_hx_tmp_0 < _g1.Len()) {
			break
		}

		var name string
		_ = name
		name = _g1.Get(_g)
		_g++
		var characterData any
		_ = characterData
		characterData = Hx_Field_reflect_field(data, name)
		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_615 func(_hx_param_key any) bool
			_ = _hx_closure_src_615
			_hx_closure_src_615 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_615(((any)(_hx_param_key)))
			}
		}()(name) {
			this.VTable.Hx_Field_restoreCharacter(func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
				var _hx_closure_src_616 func(_hx_param_key any) any
				_ = _hx_closure_src_616
				_hx_closure_src_616 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
				return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
					var hx_dyncast_240 any
					_ = hx_dyncast_240
					hx_dyncast_240 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_616(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
					var _hx_tmp_1 *Hx_Obj_loreline_runtimecharacter
					_ = _hx_tmp_1
					if hx_dyncast_240 != nil {
						_hx_tmp_1 = hx_dyncast_240.(interface {
							Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
						}).Hx_As_Hx_Obj_loreline_runtimecharacter()
					} else {
						_hx_tmp_1 = nil
					}

					return _hx_tmp_1
				}
			}()(name), characterData)
		} else {
			var newCharacter *Hx_Obj_loreline_runtimecharacter
			_ = newCharacter
			newCharacter = this.VTable.Hx_Field_restoreCharacter(nil, characterData)
			func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_runtimecharacter) {
				var _hx_closure_src_617 func(_hx_param_key any, _hx_param_value any)
				_ = _hx_closure_src_617
				_hx_closure_src_617 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
				return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_runtimecharacter) {
					_hx_closure_src_617(((any)(_hx_param_key)), ((any)(_hx_param_value)))
				}
			}()(name, newCharacter)
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreNodeStates(data any) {
	var declarationsByNode Hx_Typedef_map
	_ = declarationsByNode
	declarationsByNode = (func() *Hx_Obj_haxe_imap {
		_hx_icast_241 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_241 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_241, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_241.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nstatedecl]
		_ = _g1
		_g1 = func() func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nstatedecl] {
			var _hx_closure_src_618 func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any]
			_ = _hx_closure_src_618
			_hx_closure_src_618 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodesOfType
			return func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nstatedecl] {
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

					return HxMakeArrayView[*Hx_Obj_loreline_nstatedecl](_hx_closure_src_618(_hx_tmp_0, _hx_tmp_1))
				}
			}
		}()(((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nstatedecl_RTTI)), true)
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < _g1.Len()) {
				break
			}

			var decl *Hx_Obj_loreline_nstatedecl
			_ = decl
			decl = _g1.Get(_g)
			_g++
			if decl.Hx_Field_temporary {
				continue
			}

			var parent *Hx_Obj_loreline_node
			_ = parent
			parent = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getParentNode((func() *Hx_Obj_loreline_node {
				_hx_up := decl
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if parent == nil {
				_hx_tmp_3 = true
			} else {
				_hx_tmp_3 = !Hx_Field_std_isOfType(parent, Hx_Obj_loreline_astnode_RTTI)
			}

			if _hx_tmp_3 {
				continue
			}

			var key string
			_ = key
			key = Hx_Field_loreline__node_nodeid_impl__toString(parent.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode().Hx_Field_id)
			if !func() func(_hx_param_key string) bool {
				var _hx_closure_src_619 func(_hx_param_key any) bool
				_ = _hx_closure_src_619
				_hx_closure_src_619 = declarationsByNode.VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_619(((any)(_hx_param_key)))
				}
			}()(key) {
				var _hx_tmp_4 string
				_ = _hx_tmp_4
				_hx_tmp_4 = key
				func() func(_hx_param_key string, _hx_param_value HxArray[*Hx_Obj_loreline_nstatedecl]) {
					var _hx_closure_src_620 func(_hx_param_key any, _hx_param_value any)
					_ = _hx_closure_src_620
					_hx_closure_src_620 = declarationsByNode.VTable.Hx_Field_set
					return func(_hx_param_key string, _hx_param_value HxArray[*Hx_Obj_loreline_nstatedecl]) {
						_hx_closure_src_620(((any)(_hx_param_key)), ((any)(_hx_param_value)))
					}
				}()(_hx_tmp_4, HxMakeArray[*Hx_Obj_loreline_nstatedecl]())
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_nstatedecl](func() func(_hx_param_key string) HxArray[*Hx_Obj_loreline_nstatedecl] {
				var _hx_closure_src_621 func(_hx_param_key any) any
				_ = _hx_closure_src_621
				_hx_closure_src_621 = declarationsByNode.VTable.Hx_Field_get
				return func(_hx_param_key string) HxArray[*Hx_Obj_loreline_nstatedecl] {
					return HxAnyToTypedArray[*Hx_Obj_loreline_nstatedecl](_hx_closure_src_621(((any)(_hx_param_key))))
				}
			}()(key), decl))
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 HxArray[string]
		_ = _g3
		_g3 = Hx_Field_reflect_fields(data)
		for {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g2
			if !(_hx_tmp_5 < _g3.Len()) {
				break
			}

			var idStr string
			_ = idStr
			idStr = _g3.Get(_g2)
			_g2++
			var id Hx_Typedef_haxe__int64___int64
			_ = id
			id = Hx_Field_loreline__node_nodeid_impl__fromString(idStr)
			var stateData any
			_ = stateData
			stateData = Hx_Field_reflect_field(data, idStr)
			var declarations HxArray[*Hx_Obj_loreline_nstatedecl]
			_ = declarations
			declarations = func() func(_hx_param_key string) HxArray[*Hx_Obj_loreline_nstatedecl] {
				var _hx_closure_src_622 func(_hx_param_key any) any
				_ = _hx_closure_src_622
				_hx_closure_src_622 = declarationsByNode.VTable.Hx_Field_get
				return func(_hx_param_key string) HxArray[*Hx_Obj_loreline_nstatedecl] {
					return HxAnyToTypedArray[*Hx_Obj_loreline_nstatedecl](_hx_closure_src_622(((any)(_hx_param_key))))
				}
			}()(idStr)
			var nodeState *Hx_Obj_loreline_runtimestate
			_ = nodeState
			if declarations != nil {
				var _hx_tmp_6 any
				_ = _hx_tmp_6
				var _hx_tmp_7 HxNullable[string]
				_ = _hx_tmp_7
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(stateData, nil))) {
					var hx_nullable_t64 any
					_ = hx_nullable_t64
					hx_nullable_t64 = Hx_Field_go_haxe_hxdynamic_getField(stateData, "type")
					var _hx_tmp_8 HxNullable[string]
					_ = _hx_tmp_8
					if hx_nullable_t64 == nil {
						_hx_tmp_8 = HxNullable[string]{Value: HxStringNull}
					} else {
						_hx_tmp_8 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t64), Valid: true}
					}

					_hx_tmp_7 = _hx_tmp_8
				} else {
					_hx_tmp_7 = HxNullable[string]{Value: HxStringNull}
				}

				if (_hx_tmp_7).Valid != false {
					var _hx_tmp_9 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_9
					_hx_tmp_9 = this
					var hx_nullable_t62 any
					_ = hx_nullable_t62
					hx_nullable_t62 = Hx_Field_go_haxe_hxdynamic_getField(stateData, "type")
					var _hx_tmp_10 HxNullable[string]
					_ = _hx_tmp_10
					if hx_nullable_t62 == nil {
						_hx_tmp_10 = HxNullable[string]{Value: HxStringNull}
					} else {
						_hx_tmp_10 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t62), Valid: true}
					}

					_hx_tmp_6 = Hx_Field_loreline_objects_createFields(_hx_tmp_9, _hx_tmp_10, ((any)(declarations.Get(0))))
				} else {
					_hx_tmp_6 = nil
				}

				var fields any
				_ = fields
				fields = _hx_tmp_6
				nodeState = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
					_hx_up := declarations.Get(0)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})(), fields, nil)
				{
					var _g int32
					_ = _g
					_g = 0
					for {
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = _g
						if !(_hx_tmp_11 < declarations.Len()) {
							break
						}

						var decl *Hx_Obj_loreline_nstatedecl
						_ = decl
						decl = declarations.Get(_g)
						_g++
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							var tmp__g2_1 HxArray[*Hx_Obj_loreline_nobjectfield]
							_ = tmp__g2_1
							tmp__g2_1 = decl.Hx_Field_fields
							for {
								var _hx_tmp_12 int32
								_ = _hx_tmp_12
								_hx_tmp_12 = _g1
								if !(_hx_tmp_12 < tmp__g2_1.Len()) {
									break
								}

								var field *Hx_Obj_loreline_nobjectfield
								_ = field
								field = tmp__g2_1.Get(_g1)
								_g1++
								if this.VTable.Hx_Field_isOriginalScriptExpression(field.Hx_Field_value) {
									var value any
									_ = value
									value = this.VTable.Hx_Field_evaluateExpression(field.Hx_Field_value)
									Hx_Field_loreline_objects_setField(this, nodeState.Hx_Field_fields, field.Hx_Field_name, value)
									var _hx_tmp_13 *Hx_Obj_loreline_interpreter
									_ = _hx_tmp_13
									_hx_tmp_13 = this
									var _hx_tmp_14 any
									_ = _hx_tmp_14
									_hx_tmp_14 = nodeState.Hx_Field_originalFields
									var _hx_tmp_15 string
									_ = _hx_tmp_15
									_hx_tmp_15 = field.Hx_Field_name
									Hx_Field_loreline_objects_setField(_hx_tmp_13, _hx_tmp_14, _hx_tmp_15, this.VTable.Hx_Field_snapshotOriginalValue(value))
								}
							}
						}
					}
				}

				this.VTable.Hx_Field_restoreState(nodeState, stateData)
			} else {
				nodeState = this.VTable.Hx_Field_restoreState(nil, stateData)
			}

			{
				var key Hx_Typedef_haxe__int64___int64
				_ = key
				key = id
				var _hx_tmp_16 int64
				_ = _hx_tmp_16
				_hx_tmp_16 = ((int64)(key))
				func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					var _hx_closure_src_623 func(high int32, low int32, value any)
					_ = _hx_closure_src_623
					_hx_closure_src_623 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
					return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
						_hx_closure_src_623(high, low, ((any)(value)))
					}
				}()(((int32)((_hx_tmp_16 >> int64(((int32)(32)))))), ((int32)(key)), nodeState)
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreCharacter(character *Hx_Obj_loreline_runtimecharacter, data any) *Hx_Obj_loreline_runtimecharacter {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	if character != nil {
		_hx_tmp_0 = character.Hx_Field_fields
	} else {
		_hx_tmp_0 = nil
	}

	var target any
	_ = target
	target = _hx_tmp_0
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(target, nil))) {
		_hx_tmp_1 = target
	} else {
		_hx_tmp_1 = nil
	}

	var fields any
	_ = fields
	fields = this.VTable.Hx_Field_restoreFields(_hx_tmp_1, data)
	if character == nil {
		character = Hx_Obj_loreline_runtimecharacter_CreateInstance(this, nil, fields, nil)
	}

	return character
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreFields(target any, savedFields any) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(savedFields, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(savedFields, "fields"), nil))))
	}

	if _hx_tmp_0 {
		return target
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(target, nil))) {
		var _hx_tmp_1 *Hx_Obj_loreline_interpreter
		_ = _hx_tmp_1
		_hx_tmp_1 = this
		var hx_nullable_t65 any
		_ = hx_nullable_t65
		hx_nullable_t65 = Hx_Field_go_haxe_hxdynamic_getField(savedFields, "type")
		var _hx_tmp_2 HxNullable[string]
		_ = _hx_tmp_2
		if hx_nullable_t65 == nil {
			_hx_tmp_2 = HxNullable[string]{Value: HxStringNull}
		} else {
			_hx_tmp_2 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t65), Valid: true}
		}

		target = Hx_Field_loreline_objects_createFields(_hx_tmp_1, _hx_tmp_2, nil)
	}

	var data any
	_ = data
	data = Hx_Field_go_haxe_hxdynamic_getField(savedFields, "fields")
	if Hx_Field_loreline_objects_isFields(target) {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = Hx_Field_reflect_fields(data)
		for {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			if !(_hx_tmp_3 < _g1.Len()) {
				break
			}

			var key string
			_ = key
			key = _g1.Get(_g)
			_g++
			var _hx_tmp_4 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_4
			_hx_tmp_4 = this
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = target
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = key
			Hx_Field_loreline_objects_setField(_hx_tmp_4, _hx_tmp_5, _hx_tmp_6, this.VTable.Hx_Field_restoreValue(Hx_Field_reflect_field(data, key)))
		}
	} else {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[string]
		_ = _g1
		_g1 = Hx_Field_reflect_fields(data)
		for {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _g
			if !(_hx_tmp_7 < _g1.Len()) {
				break
			}

			var field string
			_ = field
			field = _g1.Get(_g)
			_g++
			var _hx_tmp_8 any
			_ = _hx_tmp_8
			_hx_tmp_8 = target
			var _hx_tmp_9 string
			_ = _hx_tmp_9
			_hx_tmp_9 = field
			Hx_Field_reflect_setField(_hx_tmp_8, _hx_tmp_9, this.VTable.Hx_Field_restoreValue(Hx_Field_reflect_field(data, field)))
		}
	}

	return target
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreValue(value any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return nil
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		_hx_tmp_2 = true
	} else {
		_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
	}

	if _hx_tmp_0 {
		return value
	}

	if Hx_Field_loreline_arrays_isArray(value) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(value)
		var arr any
		_ = arr
		arr = Hx_Field_loreline_arrays_createArray()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_3
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = arr
				Hx_Field_loreline_arrays_arrayPush(_hx_tmp_4, this.VTable.Hx_Field_restoreValue(Hx_Field_loreline_arrays_arrayGet(value, i)))
			}
		}

		return arr
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_reflect_field(value, "type"), ((any)("$beatRef"))))) {
		var savedRef any
		_ = savedRef
		savedRef = value
		var beat *Hx_Obj_loreline_nbeatdecl
		_ = beat
		beat = this.VTable.Hx_Field_restoreBeat(Hx_Field_go_haxe_hxdynamic_getField(savedRef, "beat"))
		if beat == nil {
			return nil
		}

		var chain HxArray[*Hx_Obj_loreline_runtimescope]
		_ = chain
		chain = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(savedRef, "stack"), ((any)(nil))))) {
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field__restoreSavedInsertions, nil))) {
				_hx_tmp_5 = this.Hx_Field__restoreSavedInsertions
			} else {
				_hx_tmp_5 = any(map[string]any{})
			}

			var savedInsertions any
			_ = savedInsertions
			savedInsertions = _hx_tmp_5
			var _hx_tmp_6 Hx_Typedef_map
			_ = _hx_tmp_6
			if this.Hx_Field__restoreInsertionsCache != nil {
				_hx_tmp_6 = this.Hx_Field__restoreInsertionsCache
			} else {
				_hx_tmp_6 = (func() *Hx_Obj_haxe_imap {
					_hx_icast_242 := Hx_Obj_haxe_ds_intmap_CreateInstance()
					if _hx_icast_242 == nil {
						return nil
					}
					Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_242, "Hx_Obj_haxe_imap")
					return &Hx_Obj_haxe_imap{VTable: _hx_icast_242.VTable.(Hx_Obj_VTable_haxe_imap)}
				})()
			}

			var cache Hx_Typedef_map
			_ = cache
			cache = _hx_tmp_6
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[any]
				_ = _g1
				_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(savedRef, "stack"))
				for {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _g
					if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
						break
					}

					var savedScope any
					_ = savedScope
					savedScope = _g1.Get(_g)
					_g++
					var scope *Hx_Obj_loreline_runtimescope
					_ = scope
					scope = this.VTable.Hx_Field_restoreScopeItem(savedScope, savedInsertions, cache)
					if scope == nil {
						chain = HxMakeArray[*Hx_Obj_loreline_runtimescope]()
						break
					}

					_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](chain, scope))
				}
			}
		}

		return ((any)(Hx_Obj_loreline_runtimebeatref_CreateInstance(beat, chain)))
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_reflect_field(value, "type"), ((any)("$characterRef"))))) {
		var savedCharacterRef any
		_ = savedCharacterRef
		savedCharacterRef = value
		var character *Hx_Obj_loreline_runtimecharacter
		_ = character
		character = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
			var _hx_closure_src_624 func(_hx_param_key any) any
			_ = _hx_closure_src_624
			_hx_closure_src_624 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
				var hx_dyncast_243 any
				_ = hx_dyncast_243
				hx_dyncast_243 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_624(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
				var _hx_tmp_8 *Hx_Obj_loreline_runtimecharacter
				_ = _hx_tmp_8
				if hx_dyncast_243 != nil {
					_hx_tmp_8 = hx_dyncast_243.(interface {
						Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
					}).Hx_As_Hx_Obj_loreline_runtimecharacter()
				} else {
					_hx_tmp_8 = nil
				}

				return _hx_tmp_8
			}
		}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(savedCharacterRef, "name")))
		if character == nil {
			return nil
		}

		if character.Hx_Field_characterRef == nil {
			character.Hx_Field_characterRef = Hx_Obj_loreline_runtimecharacterref_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(savedCharacterRef, "name")), character)
		}

		return ((any)(character.Hx_Field_characterRef))
	}

	return this.VTable.Hx_Field_restoreFields(nil, value)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_bindFunctions(functions Hx_Typedef_loreline_functionsmap) {
	this.Hx_Field_builtins = Hx_Obj_loreline_functions_CreateInstance(this)
	this.Hx_Field_builtins.VTable.Hx_Field_bindAll(this.Hx_Field_topLevelFunctions)
	if functions != nil {
		var _hx_reserved_map *Hx_Obj_haxe_imap
		_ = _hx_reserved_map
		_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
			_hx_icast_244 := functions
			if _hx_icast_244 == nil {
				return nil
			}
			Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_244, "Hx_Obj_haxe_imap")
			return &Hx_Obj_haxe_imap{VTable: _hx_icast_244.VTable.(Hx_Obj_VTable_haxe_imap)}
		})()
		var _g_map *Hx_Obj_haxe_imap
		_ = _g_map
		_g_map = _hx_reserved_map
		var _g_keys any
		_ = _g_keys
		_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
			var key string
			_ = key
			key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
			var _g_value any
			_ = _g_value
			_g_value = _g_map.VTable.Hx_Field_get(key)
			var _g_key string
			_ = _g_key
			_g_key = key
			var key1 string
			_ = key1
			key1 = _g_key
			var _hx_reserved_func any
			_ = _hx_reserved_func
			_hx_reserved_func = _g_value
			this.VTable.Hx_Field_registerTopLevelFunction(key1, _hx_reserved_func)
		}
	}

	this.VTable.Hx_Field_refreshFunctionHelpers()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = this.Hx_Field_context.Hx_Field_scriptFunctions
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var _hx_reserved_func any
			_ = _hx_reserved_func
			_hx_reserved_func = _g1.Get(_g)
			_g++
			this.VTable.Hx_Field_bindScriptFunction(_hx_reserved_func)
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_refreshFunctionHelpers() {
	var _hx_reserved_map *Hx_Obj_haxe_imap
	_ = _hx_reserved_map
	_hx_reserved_map = (func() *Hx_Obj_haxe_imap {
		_hx_icast_245 := this.Hx_Field_topLevelFunctions.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap()
		if _hx_icast_245 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_245, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_245.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var _g_map *Hx_Obj_haxe_imap
	_ = _g_map
	_g_map = _hx_reserved_map
	var _g_keys any
	_ = _g_keys
	_g_keys = _hx_reserved_map.VTable.Hx_Field_keys()
	for {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "hasNext")
		if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
			break
		}

		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_g_keys, "next")
		var key string
		_ = key
		key = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
		var _g_value any
		_ = _g_value
		_g_value = _g_map.VTable.Hx_Field_get(key)
		var _g_key string
		_ = _g_key
		_g_key = key
		var key1 string
		_ = key1
		key1 = _g_key
		var _hx_reserved_func any
		_ = _hx_reserved_func
		_hx_reserved_func = _g_value
		if Hx_Field_stringtools_startsWith(key1, "string_") {
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = key1
			this.Hx_Field_stringHelpers.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 7, HxNullable[int32]{})), _hx_reserved_func)
		} else {
			if Hx_Field_stringtools_startsWith(key1, "array_") {
				var _hx_tmp_3 string
				_ = _hx_tmp_3
				_hx_tmp_3 = key1
				this.Hx_Field_arrayHelpers.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, 6, HxNullable[int32]{})), _hx_reserved_func)
			} else {
				if Hx_Field_stringtools_startsWith(key1, "map_") {
					var _hx_tmp_4 string
					_ = _hx_tmp_4
					_hx_tmp_4 = key1
					this.Hx_Field_mapHelpers.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, 4, HxNullable[int32]{})), _hx_reserved_func)
				} else {
					if Hx_Field_stringtools_startsWith(key1, "beat_") {
						var _hx_tmp_5 string
						_ = _hx_tmp_5
						_hx_tmp_5 = key1
						this.Hx_Field_beatHelpers.VTable.(interface {
							Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
						}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, 5, HxNullable[int32]{})), _hx_reserved_func)
					}
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_bindScriptFunction(_hx_reserved_func any) {
	var interp *Hx_Obj_loreline_lorscript_interp
	_ = interp
	interp = Hx_Obj_loreline_lorscript_interp_CreateInstance(this)
	var value any
	_ = value
	value = interp.VTable.Hx_Field_execute(Hx_Field_go_haxe_hxdynamic_getField(_hx_reserved_func, "ast"))
	this.Hx_Field_topLevelFunctions.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(Hx_Field_go_haxe_hxdynamic_getField(_hx_reserved_func, "name"), value)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_registerTopLevelFunction(key string, _hx_reserved_func any) {
	var userFunc any
	_ = userFunc
	userFunc = _hx_reserved_func
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = key
	this.Hx_Field_topLevelFunctions.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set(_hx_tmp_0, Hx_Field_reflect_makeVarArgs(func(args HxArray[any]) any {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g1
				if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(args)) {
					break
				}

				var arg any
				_ = arg
				arg = args.Get(_g1)
				_g1++
				var _hx_tmp_2 HxArray[any]
				_ = _hx_tmp_2
				_hx_tmp_2 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_2, _gthis.VTable.Hx_Field_hostValueOf(arg)))
			}
		}

		var hostArgs HxArray[any]
		_ = hostArgs
		hostArgs = _g
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = userFunc
		return Hx_Field_reflect_callMethod(nil, _hx_tmp_3, HxMakeArray[any](((any)(_gthis)), ((any)(hostArgs))).Dyn())
	}))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeStringLiteralProcessors(processors HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]) {
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var _hx_tmp_0 HxArray[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_context.Hx_Field_stringLiteralProcessors
	_ = (Hx_Array_Push[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral](_hx_tmp_0, func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
		return _gthis.VTable.Hx_Field_pluralPipeProcess(str)
	}))
	if processors != nil {
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < processors.Len()) {
				break
			}

			var p func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral
			_ = p
			p = processors.Get(_g)
			_g++
			_ = (Hx_Array_Push[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral](this.Hx_Field_context.Hx_Field_stringLiteralProcessors, p))
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeTopLevelFunction(_hx_reserved_func *Hx_Obj_loreline_nfunctiondecl) {
	if (_hx_reserved_func.Hx_Field_name).Valid != false {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if !_hx_reserved_func.Hx_Field_external {
			_hx_tmp_0 = true
		} else {
			_hx_tmp_0 = !func() func(_hx_param_key string) bool {
				var _hx_closure_src_625 func(_hx_param_key any) bool
				_ = _hx_closure_src_625
				_hx_closure_src_625 = this.Hx_Field_topLevelFunctions.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_625(((any)(_hx_param_key)))
				}
			}()((_hx_reserved_func.Hx_Field_name).Value)
		}

		if _hx_tmp_0 {
			var codeToLorscript *Hx_Obj_loreline_codetolorscript
			_ = codeToLorscript
			codeToLorscript = Hx_Obj_loreline_codetolorscript_CreateInstance()
			{
				var hx_try_state int
				_ = hx_try_state

				func() {
					defer func() {
						if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
							if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
								_ = e
								{
									var _hx_tmp_6 string
									_ = _hx_tmp_6
									if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(e, nil))) {
										_hx_tmp_6 = "null"
									} else {
										_hx_tmp_6 = Hx_Field_std_string(e)
									}

									panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Failed to parse function code: ", (_hx_tmp_6)), _hx_reserved_func.Hx_Field_pos)))))
								}
							} else {
								panic(_hx_ex)
							}
						}
					}()
					{
						var _hx_tmp_1 string
						_ = _hx_tmp_1
						_hx_tmp_1 = _hx_reserved_func.Hx_Field_code
						var _hx_tmp_2 string
						_ = _hx_tmp_2
						if _hx_reserved_func.Hx_Field_external {
							_hx_tmp_2 = " {}"
						} else {
							_hx_tmp_2 = ""
						}

						var expr string
						_ = expr
						expr = codeToLorscript.VTable.Hx_Field_process(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_1, (_hx_tmp_2)))
						var parser *Hx_Obj_loreline_lorscript_parser
						_ = parser
						parser = Hx_Obj_loreline_lorscript_parser_CreateInstance()
						parser.Hx_Field_allowJSON = true
						parser.Hx_Field_allowTypes = true
						var hx_nullable_f210 HxNullable[string]
						_ = hx_nullable_f210
						hx_nullable_f210 = _hx_reserved_func.Hx_Field_name
						var _hx_tmp_3 any
						_ = _hx_tmp_3
						if (hx_nullable_f210).Valid == true {
							_hx_tmp_3 = (hx_nullable_f210).Value
						} else {
							_hx_tmp_3 = nil
						}

						var _hx_tmp_4 string
						_ = _hx_tmp_4
						_hx_tmp_4 = expr
						var _hx_tmp_5 any
						_ = _hx_tmp_5
						_hx_tmp_5 = nil
						var scriptFunction any
						_ = scriptFunction
						scriptFunction = any(map[string]any{"name": _hx_tmp_3, "ast": parser.VTable.Hx_Field_parseString(_hx_tmp_4, _hx_tmp_5, nil)})
						_ = (Hx_Array_Push[any](HxMakeArrayView[any](this.Hx_Field_context.Hx_Field_scriptFunctions), scriptFunction))
						this.VTable.Hx_Field_bindScriptFunction(scriptFunction)
					}
				}()
				if hx_try_state == 1 {
					return
				}
			}
		}
	} else {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Top level function must have a name", _hx_reserved_func.Hx_Field_pos)))))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_wrapNext(cb func()) *Hx_Obj_loreline_evalnext {
	var wrapped *Hx_Obj_loreline_evalnext
	_ = wrapped
	wrapped = Hx_Obj_loreline_evalnext_CreateInstance()
	wrapped.Hx_Field_sync = true
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	wrapped.Hx_Field_cb = func() {
		if _gthis.Hx_Field_disposed {
			return
		}

		if wrapped.Hx_Field_sync {
			if _gthis.Hx_Field_syncCallbacks == nil {
				_gthis.Hx_Field_syncCallbacks = HxMakeArray[func()]()
			}

			_ = (Hx_Array_Push[func()](_gthis.Hx_Field_syncCallbacks, wrapped.Hx_Field_cb))
		} else {
			cb()
			_gthis.VTable.Hx_Field_flush()
			wrapped.Hx_Field_cb = nil
			if _gthis.Hx_Field_finishTrigger == wrapped {
				_gthis.VTable.Hx_Field_finish()
			}
		}
	}

	return wrapped
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_flush() {
	if this.Hx_Field_flushing {
		return
	}

	this.Hx_Field_flushing = true
	{
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							this.Hx_Field_flushing = false
							panic(Hx_Field_haxe_exception_thrown(e))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				if this.Hx_Field_syncCallbacks != nil {
					for {
						if !(this.Hx_Field_syncCallbacks.Len() > 0) {
							break
						}

						var cb func()
						_ = cb
						cb = (Hx_Array_Shift[func()](this.Hx_Field_syncCallbacks).Value)
						var prevSyncCallbacks HxArray[func()]
						_ = prevSyncCallbacks
						prevSyncCallbacks = this.Hx_Field_syncCallbacks
						this.Hx_Field_syncCallbacks = nil
						cb()
						if this.Hx_Field_syncCallbacks != nil {
							var i int32
							_ = i
							i = (this.Hx_Field_syncCallbacks.Len() - 1)
							for i >= 0 {
								(Hx_Array_Unshift[func()](prevSyncCallbacks, this.Hx_Field_syncCallbacks.Get(i)))
								i--
							}
						}

						this.Hx_Field_syncCallbacks = prevSyncCallbacks
					}
				}
			}
		}()
		if hx_try_state == 1 {
			return
		}
	}

	this.Hx_Field_flushing = false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_pop() bool {
	if this.Hx_Field_stack.Len() > 0 {
		_ = (Hx_Array_Pop[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack).Value)
		return true
	}

	return false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_push(scope *Hx_Obj_loreline_runtimescope) {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_nextScopeId
	this.Hx_Field_nextScopeId = (this.Hx_Field_nextScopeId + 1)
	scope.Hx_Field_id = _hx_tmp_0
	_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack, scope))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeTopLevelState(state *Hx_Obj_loreline_nstatedecl) {
	if state.Hx_Field_temporary {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Top level temporary states are not allowed", state.Hx_Field_pos)))))
	}

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
			var evaluated any
			_ = evaluated
			evaluated = this.VTable.Hx_Field_evaluateExpression(field.Hx_Field_value)
			Hx_Field_loreline_objects_setField(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, field.Hx_Field_name, evaluated)
			if this.VTable.Hx_Field_isOriginalScriptExpression(field.Hx_Field_value) {
				var _hx_tmp_1 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_1
				_hx_tmp_1 = this
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_originalFields
				var _hx_tmp_3 string
				_ = _hx_tmp_3
				_hx_tmp_3 = field.Hx_Field_name
				Hx_Field_loreline_objects_setField(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, this.VTable.Hx_Field_snapshotOriginalValue(evaluated))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeTopLevelBeat(beat *Hx_Obj_loreline_nbeatdecl) {
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_626 func(_hx_param_key any) bool
		_ = _hx_closure_src_626
		_hx_closure_src_626 = this.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_626(((any)(_hx_param_key)))
		}
	}()(beat.Hx_Field_name) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Duplicate top level beat: ", beat.Hx_Field_name), beat.Hx_Field_pos)))))
	}

	func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nbeatdecl) {
		var _hx_closure_src_627 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_627
		_hx_closure_src_627 = this.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_nbeatdecl) {
			_hx_closure_src_627(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(beat.Hx_Field_name, beat)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeTopLevelCharacter(character *Hx_Obj_loreline_ncharacterdecl) {
	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_628 func(_hx_param_key any) bool
		_ = _hx_closure_src_628
		_hx_closure_src_628 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_628(((any)(_hx_param_key)))
		}
	}()(character.Hx_Field_name) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Duplicate top level character: ", character.Hx_Field_name), character.Hx_Field_pos)))))
	}

	var characterState *Hx_Obj_loreline_runtimecharacter
	_ = characterState
	characterState = Hx_Obj_loreline_runtimecharacter_CreateInstance(this, (func() *Hx_Obj_loreline_astnode {
		_hx_up := character
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), nil, nil)
	func() func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_runtimecharacter) {
		var _hx_closure_src_629 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_629
		_hx_closure_src_629 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value *Hx_Obj_loreline_runtimecharacter) {
			_hx_closure_src_629(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(character.Hx_Field_name, characterState)
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = character.Hx_Field_fields
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
			var evaluated any
			_ = evaluated
			evaluated = this.VTable.Hx_Field_evaluateExpression(field.Hx_Field_value)
			Hx_Field_loreline_objects_setField(this, characterState.Hx_Field_fields, field.Hx_Field_name, evaluated)
			if this.VTable.Hx_Field_isOriginalScriptExpression(field.Hx_Field_value) {
				var _hx_tmp_1 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_1
				_hx_tmp_1 = this
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = characterState.Hx_Field_originalFields
				var _hx_tmp_3 string
				_ = _hx_tmp_3
				_hx_tmp_3 = field.Hx_Field_name
				Hx_Field_loreline_objects_setField(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, this.VTable.Hx_Field_snapshotOriginalValue(evaluated))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_initializeState(state *Hx_Obj_loreline_nstatedecl, scope *Hx_Obj_loreline_runtimescope) {
	var runtimeState *Hx_Obj_loreline_runtimestate
	_ = runtimeState
	if state.Hx_Field_temporary {
		if scope.Hx_Field_state == nil {
			scope.Hx_Field_state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
				_hx_up := state
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})(), nil, nil)
		}

		runtimeState = scope.Hx_Field_state
	} else {
		var key Hx_Typedef_haxe__int64___int64
		_ = key
		key = scope.Hx_Field_node.Hx_Field_id
		var _hx_tmp_1 int64
		_ = _hx_tmp_1
		_hx_tmp_1 = ((int64)(key))
		runtimeState = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var _hx_closure_src_630 func(high int32, low int32) any
			_ = _hx_closure_src_630
			_hx_closure_src_630 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
			return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
				var hx_dyncast_246 any
				_ = hx_dyncast_246
				hx_dyncast_246 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_630(high, low), "Hx_Obj_loreline_runtimestate")
				var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
				_ = _hx_tmp_0
				if hx_dyncast_246 != nil {
					_hx_tmp_0 = hx_dyncast_246.(interface {
						Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
					}).Hx_As_Hx_Obj_loreline_runtimestate()
				} else {
					_hx_tmp_0 = nil
				}

				return _hx_tmp_0
			}
		}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
		if runtimeState == nil {
			runtimeState = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
				_hx_up := state
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})(), nil, nil)
			{
				var tmp_key_1 Hx_Typedef_haxe__int64___int64
				_ = tmp_key_1
				tmp_key_1 = scope.Hx_Field_node.Hx_Field_id
				var _hx_tmp_2 int64
				_ = _hx_tmp_2
				_hx_tmp_2 = ((int64)(tmp_key_1))
				func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					var _hx_closure_src_631 func(high int32, low int32, value any)
					_ = _hx_closure_src_631
					_hx_closure_src_631 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
					return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
						_hx_closure_src_631(high, low, ((any)(value)))
					}
				}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), runtimeState)
			}
		}
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = state.Hx_Field_fields
		for {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			if !(_hx_tmp_3 < _g1.Len()) {
				break
			}

			var field *Hx_Obj_loreline_nobjectfield
			_ = field
			field = _g1.Get(_g)
			_g++
			if !Hx_Field_loreline_objects_fieldExists(this, runtimeState.Hx_Field_fields, field.Hx_Field_name) {
				var evaluated any
				_ = evaluated
				evaluated = this.VTable.Hx_Field_evaluateExpression(field.Hx_Field_value)
				Hx_Field_loreline_objects_setField(this, runtimeState.Hx_Field_fields, field.Hx_Field_name, evaluated)
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if !state.Hx_Field_temporary {
					_hx_tmp_4 = this.VTable.Hx_Field_isOriginalScriptExpression(field.Hx_Field_value)
				} else {
					_hx_tmp_4 = false
				}

				if _hx_tmp_4 {
					var _hx_tmp_5 *Hx_Obj_loreline_interpreter
					_ = _hx_tmp_5
					_hx_tmp_5 = this
					var _hx_tmp_6 any
					_ = _hx_tmp_6
					_hx_tmp_6 = runtimeState.Hx_Field_originalFields
					var _hx_tmp_7 string
					_ = _hx_tmp_7
					_hx_tmp_7 = field.Hx_Field_name
					Hx_Field_loreline_objects_setField(_hx_tmp_5, _hx_tmp_6, _hx_tmp_7, this.VTable.Hx_Field_snapshotOriginalValue(evaluated))
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateBeatArgs(beat *Hx_Obj_loreline_nbeatdecl, args HxArray[*Hx_Obj_loreline_nexpr], pos *Hx_Obj_loreline_position) HxArray[any] {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if args == nil {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (args.Len() == 0)
	}

	if _hx_tmp_0 {
		return nil
	}

	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if beat.Hx_Field_params != nil {
		_hx_tmp_1 = beat.Hx_Field_params.Len()
	} else {
		_hx_tmp_1 = 0
	}

	var numParams int32
	_ = numParams
	numParams = _hx_tmp_1
	if args.Len() > numParams {
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Beat ", beat.Hx_Field_name), " expects at most ")
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_toString(numParams)), " argument(s), got ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_toString(args.Len())), pos)))))
	}

	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g1
			if !(_hx_tmp_4 < args.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = args.Get(_g1)
			_g1++
			var _hx_tmp_5 HxArray[any]
			_ = _hx_tmp_5
			_hx_tmp_5 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_5, this.VTable.Hx_Field_evaluateExpression(arg)))
		}
	}

	return _g
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_seedBeatParams(scope *Hx_Obj_loreline_runtimescope, beat *Hx_Obj_loreline_nbeatdecl, values HxArray[any]) {
	var params HxArray[*Hx_Obj_loreline_nbeatparam]
	_ = params
	params = beat.Hx_Field_params
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if params == nil {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = (params.Len() == 0)
	}

	if _hx_tmp_0 {
		return
	}

	if scope.Hx_Field_state == nil {
		scope.Hx_Field_state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), nil, nil)
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = params.Len()
		for _g < _g1 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			var param *Hx_Obj_loreline_nbeatparam
			_ = param
			param = params.Get(i)
			var value any
			_ = value
			var _hx_tmp_2 bool
			_ = _hx_tmp_2
			if values != nil {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = i
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_3
				_hx_tmp_2 = (_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(values))
			} else {
				_hx_tmp_2 = false
			}

			if _hx_tmp_2 {
				value = values.Get(i)
			} else {
				if param.Hx_Field_defaultValue != nil {
					value = this.VTable.Hx_Field_evaluateExpression(param.Hx_Field_defaultValue)
				}
			}

			Hx_Field_loreline_objects_setField(this, scope.Hx_Field_state.Hx_Field_fields, param.Hx_Field_name, value)
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_finish() {
	this.Hx_Field_finishTrigger = nil
	if this.VTable.Hx_Field_isRoot() {
		this.Hx_Field_finished = true
	} else {
		_ = (Hx_Array_Remove[*Hx_Obj_loreline_interpreter](this.Hx_Field_context.Hx_Field_children, this))
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(this.Hx_Field_handleFinish)), ((any)(nil))))) {
		this.Hx_Field_handleFinish(this)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_transitionToBeat(beat *Hx_Obj_loreline_nbeatdecl, argValues__isOptional any, capturedChain__isOptional any) {
	argValues := HxDefault[HxArray[any]]()
	_ = argValues

	if argValues__isOptional != nil {
		argValues = HxConvert[HxArray[any]](argValues__isOptional)
	} else {
		argValues = nil
	}
	capturedChain := HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	_ = capturedChain

	if capturedChain__isOptional != nil {
		capturedChain = HxConvert[HxArray[*Hx_Obj_loreline_runtimescope]](capturedChain__isOptional)
	} else {
		capturedChain = nil
	}
	{
		for {
			if !(this.VTable.Hx_Field_pop()) {
				break
			}
		}

		this.Hx_Field_nextScopeId = 1
		this.Hx_Field_pendingChoiceOptions = nil
		this.Hx_Field_pendingDialogue = nil
		(Hx_Array_Resize[string](this.Hx_Field__choiceEvalTexts, 0))
		(Hx_Array_Resize[bool](this.Hx_Field__choiceEvalEnabled, 0))
		var done *Hx_Obj_loreline_evalnext
		_ = done
		done = this.VTable.Hx_Field_wrapNext(this.VTable.Hx_Field_finish)
		this.Hx_Field_finishTrigger = done
		this.VTable.Hx_Field_evalBeatRun(beat, done.Hx_Field_cb, argValues, capturedChain)
		done.Hx_Field_sync = false
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalNode(node *Hx_Obj_loreline_astnode, next func()) {
	var _g *Hx_Obj_go_haxe_hxclass
	_ = _g
	_g = func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
		var _hx_closure_src_632 func(o any) *Hx_Obj_go_haxe_hxclass
		_ = _hx_closure_src_632
		_hx_closure_src_632 = Hx_Field_type_getClass
		return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
			return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_632(((any)(o)))))
		}
	}()(node)
	if _g == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unsupported node type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_635 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_635
			_hx_closure_src_635 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_635(((any)(o)))))
			}
		}()(node))))), node.Hx_Field_pos)))))
	} else {
		switch _g {
		case Hx_Obj_loreline_nalternative_RTTI:
			{
				this.VTable.Hx_Field_evalAlternative(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
				}).Hx_As_Hx_Obj_loreline_nalternative(), next)
			}

		case Hx_Obj_loreline_nassign_RTTI:
			{
				this.VTable.Hx_Field_evalAssignment(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign
				}).Hx_As_Hx_Obj_loreline_nassign(), next)
			}

		case Hx_Obj_loreline_nbeatcall_RTTI:
			{
				this.VTable.Hx_Field_evalBeatCall(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall
				}).Hx_As_Hx_Obj_loreline_nbeatcall(), next)
			}

		case Hx_Obj_loreline_nbeatdecl_RTTI:
			{
				this.VTable.Hx_Field_evalBeatDecl(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl(), next)
			}

		case Hx_Obj_loreline_ncall_RTTI:
			{
				this.VTable.Hx_Field_evalCall(node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall(), next)
			}

		case Hx_Obj_loreline_nchoiceoption_RTTI:
			{
				this.VTable.Hx_Field_evalChoiceOption(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption
				}).Hx_As_Hx_Obj_loreline_nchoiceoption(), next)
			}

		case Hx_Obj_loreline_nchoicestatement_RTTI:
			{
				this.VTable.Hx_Field_evalChoice(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement
				}).Hx_As_Hx_Obj_loreline_nchoicestatement(), next)
			}

		case Hx_Obj_loreline_ndialoguestatement_RTTI:
			{
				this.VTable.Hx_Field_evalDialogue(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement
				}).Hx_As_Hx_Obj_loreline_ndialoguestatement(), next)
			}

		case Hx_Obj_loreline_nifstatement_RTTI:
			{
				this.VTable.Hx_Field_evalIf(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement
				}).Hx_As_Hx_Obj_loreline_nifstatement(), next)
			}

		case Hx_Obj_loreline_nstatedecl_RTTI:
			{
				this.VTable.Hx_Field_evalStateDecl(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl
				}).Hx_As_Hx_Obj_loreline_nstatedecl(), next)
			}

		case Hx_Obj_loreline_ntextstatement_RTTI:
			{
				this.VTable.Hx_Field_evalText(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement
				}).Hx_As_Hx_Obj_loreline_ntextstatement(), next)
			}

		case Hx_Obj_loreline_ntransition_RTTI:
			{
				this.VTable.Hx_Field_evalTransition(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition
				}).Hx_As_Hx_Obj_loreline_ntransition())
			}

		case Hx_Obj_loreline_nwhenstatement_RTTI:
			{
				this.VTable.Hx_Field_evalWhen(node.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nwhenstatement() *Hx_Obj_loreline_nwhenstatement
				}).Hx_As_Hx_Obj_loreline_nwhenstatement(), next)
			}

		default:
			{
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unsupported node type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
					var _hx_closure_src_636 func(o any) *Hx_Obj_go_haxe_hxclass
					_ = _hx_closure_src_636
					_hx_closure_src_636 = Hx_Field_type_getClass
					return func(o *Hx_Obj_loreline_astnode) *Hx_Obj_go_haxe_hxclass {
						return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_636(((any)(o)))))
					}
				}()(node))))), node.Hx_Field_pos)))))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalBeatDecl(beat *Hx_Obj_loreline_nbeatdecl, next func()) {
	if this.VTable.Hx_Field_get_currentScope().Hx_Field_beats == nil {
		this.VTable.Hx_Field_get_currentScope().Hx_Field_beats = HxMakeArray[*Hx_Obj_loreline_nbeatdecl]()
	} else {
		if this.VTable.Hx_Field_get_currentScope().VTable.Hx_Field_beatByName(beat.Hx_Field_name) != nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Duplicate beat with name: ", beat.Hx_Field_name), beat.Hx_Field_pos)))))
		}
	}

	_ = (Hx_Array_Push[*Hx_Obj_loreline_nbeatdecl](this.VTable.Hx_Field_get_currentScope().Hx_Field_beats, beat))
	next()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalNodeBody(beat *Hx_Obj_loreline_nbeatdecl, node *Hx_Obj_loreline_astnode, body HxArray[*Hx_Obj_loreline_astnode], insertion__isOptional any, argValues__isOptional any, capturedChain__isOptional any, shuffle__isOptional any, next func()) {
	insertion := HxDefault[*Hx_Obj_loreline_runtimeinsertion]()
	_ = insertion

	if insertion__isOptional != nil {
		insertion = HxConvert[*Hx_Obj_loreline_runtimeinsertion](insertion__isOptional)
	} else {
		insertion = nil
	}
	argValues := HxDefault[HxArray[any]]()
	_ = argValues

	if argValues__isOptional != nil {
		argValues = HxConvert[HxArray[any]](argValues__isOptional)
	} else {
		argValues = nil
	}
	capturedChain := HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	_ = capturedChain

	if capturedChain__isOptional != nil {
		capturedChain = HxConvert[HxArray[*Hx_Obj_loreline_runtimescope]](capturedChain__isOptional)
	} else {
		capturedChain = nil
	}
	shuffle := HxDefault[HxArray[int32]]()
	_ = shuffle

	if shuffle__isOptional != nil {
		shuffle = HxConvert[HxArray[int32]](shuffle__isOptional)
	} else {
		shuffle = nil
	}
	{
		this.VTable.Hx_Field_push(Hx_Obj_loreline_runtimescope_CreateInstance(nil, beat, node, nil, nil, nil, insertion, nil, shuffle))
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_2
		_hx_tmp_2 = node
		if _hx_tmp_2 == (func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})() {
			_hx_tmp_1 = (beat != nil)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = (beat.Hx_Field_params != nil)
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			this.VTable.Hx_Field_seedBeatParams(this.VTable.Hx_Field_get_currentScope(), beat, argValues)
		}

		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		var _hx_tmp_4 bool
		_ = _hx_tmp_4
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		var _hx_tmp_6 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_6
		_hx_tmp_6 = node
		if _hx_tmp_6 == (func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})() {
			_hx_tmp_5 = (beat != nil)
		} else {
			_hx_tmp_5 = false
		}

		if _hx_tmp_5 {
			_hx_tmp_4 = (capturedChain != nil)
		} else {
			_hx_tmp_4 = false
		}

		if _hx_tmp_4 {
			_hx_tmp_3 = (capturedChain.Len() > 0)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			this.VTable.Hx_Field_get_currentScope().Hx_Field_captured = capturedChain
		}

		var index int32
		_ = index
		index = 0
		var moveNext func()
		_ = moveNext
		var currentInsertion *Hx_Obj_loreline_runtimeinsertion
		_ = currentInsertion
		currentInsertion = this.VTable.Hx_Field_get_currentInsertion()
		var _gthis *Hx_Obj_loreline_interpreter
		_ = _gthis
		_gthis = this
		moveNext = func() {
			var _hx_tmp_7 bool
			_ = _hx_tmp_7
			if currentInsertion != nil {
				_hx_tmp_7 = currentInsertion.VTable.Hx_Field_isCollected()
			} else {
				_hx_tmp_7 = false
			}

			if _hx_tmp_7 {
				_gthis.VTable.Hx_Field_pop()
				next()
			} else {
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = index
				if _hx_tmp_8 < body.Len() {
					var childNode *Hx_Obj_loreline_astnode
					_ = childNode
					childNode = body.Get(index)
					_gthis.VTable.Hx_Field_get_currentScope().Hx_Field_head = childNode
					index = (index + 1)
					var done *Hx_Obj_loreline_evalnext
					_ = done
					done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
					_gthis.VTable.Hx_Field_evalNode(childNode, done.Hx_Field_cb)
					done.Hx_Field_sync = false
				} else {
					_gthis.VTable.Hx_Field_pop()
					next()
				}
			}
		}

		moveNext()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalBeatRun(beat *Hx_Obj_loreline_nbeatdecl, next func(), argValues__isOptional any, capturedChain__isOptional any) {
	argValues := HxDefault[HxArray[any]]()
	_ = argValues

	if argValues__isOptional != nil {
		argValues = HxConvert[HxArray[any]](argValues__isOptional)
	} else {
		argValues = nil
	}
	capturedChain := HxDefault[HxArray[*Hx_Obj_loreline_runtimescope]]()
	_ = capturedChain

	if capturedChain__isOptional != nil {
		capturedChain = HxConvert[HxArray[*Hx_Obj_loreline_runtimescope]](capturedChain__isOptional)
	} else {
		capturedChain = nil
	}
	{
		this.VTable.Hx_Field_incrementBeatVisitCount(beat)
		var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_0
		_hx_tmp_0 = beat
		var _hx_tmp_1 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
		var _hx_tmp_2 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_2
		_hx_tmp_2 = beat.Hx_Field_body
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = nil
		var _hx_tmp_4 HxArray[any]
		_ = _hx_tmp_4
		_hx_tmp_4 = argValues
		var _hx_tmp_5 HxArray[*Hx_Obj_loreline_runtimescope]
		_ = _hx_tmp_5
		_hx_tmp_5 = capturedChain
		this.VTable.Hx_Field_evalNodeBody(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, nil, next)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalStateDecl(state *Hx_Obj_loreline_nstatedecl, next func()) {
	var _hx_tmp_0 *Hx_Obj_loreline_nstatedecl
	_ = _hx_tmp_0
	_hx_tmp_0 = state
	this.VTable.Hx_Field_initializeState(_hx_tmp_0, this.VTable.Hx_Field_get_currentScope())
	next()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalText(text *Hx_Obj_loreline_ntextstatement, next func()) {
	if this.VTable.Hx_Field_presentRestoredDialogue((func() *Hx_Obj_loreline_astnode {
		_hx_up := text
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), next) {
		return
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if text.Hx_Field_condition != nil {
		_hx_tmp_0 = !this.VTable.Hx_Field_evaluateCondition(text.Hx_Field_condition)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		next()
		return
	}

	var str *Hx_Obj_loreline_nstringliteral
	_ = str
	str = this.VTable.Hx_Field_getTranslatedString((func() *Hx_Obj_loreline_astnode {
		_hx_up := text
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), text.Hx_Field_content)
	var content any
	_ = content
	content = this.VTable.Hx_Field_evaluateString(str)
	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := text
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_2 HxNullable[string]
	_ = _hx_tmp_2
	_hx_tmp_2 = HxNullable[string]{Value: HxStringNull}

	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(content, "text"))
	this.VTable.Hx_Field_presentDialogue(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, HxAnyToTypedArray[*Hx_Obj_loreline_texttag](Hx_Field_go_haxe_hxdynamic_getField(content, "tags")), next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalDialogue(dialogue *Hx_Obj_loreline_ndialoguestatement, next func()) {
	if this.VTable.Hx_Field_presentRestoredDialogue((func() *Hx_Obj_loreline_astnode {
		_hx_up := dialogue
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), next) {
		return
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if dialogue.Hx_Field_condition != nil {
		_hx_tmp_0 = !this.VTable.Hx_Field_evaluateCondition(dialogue.Hx_Field_condition)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		next()
		return
	}

	var str *Hx_Obj_loreline_nstringliteral
	_ = str
	str = this.VTable.Hx_Field_getTranslatedString((func() *Hx_Obj_loreline_astnode {
		_hx_up := dialogue
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), dialogue.Hx_Field_content)
	var content any
	_ = content
	content = this.VTable.Hx_Field_evaluateString(str)
	var speaker string
	_ = speaker
	speaker = dialogue.Hx_Field_character
	var speakerValue any
	_ = speakerValue
	speakerValue = this.VTable.Hx_Field_getStateFieldRaw(speaker)
	var speakerRef *Hx_Obj_loreline_runtimecharacterref
	_ = speakerRef
	speakerRef = Hx_Field_loreline_runtimecharacterref_characterOf(speakerValue)
	if speakerRef != nil {
		speaker = speakerRef.Hx_Field_name
	}

	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := dialogue
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_2 HxNullable[string]
	_ = _hx_tmp_2
	_hx_tmp_2 = HxNullableString(speaker)
	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(content, "text"))
	this.VTable.Hx_Field_presentDialogue(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, HxAnyToTypedArray[*Hx_Obj_loreline_texttag](Hx_Field_go_haxe_hxdynamic_getField(content, "tags")), next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_presentDialogue(node *Hx_Obj_loreline_astnode, character HxNullable[string], text string, tags HxArray[*Hx_Obj_loreline_texttag], next func()) {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(node))
	var hx_nullable_f211 HxNullable[string]
	_ = hx_nullable_f211
	hx_nullable_f211 = character
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	if (hx_nullable_f211).Valid == true {
		_hx_tmp_1 = (hx_nullable_f211).Value
	} else {
		_hx_tmp_1 = nil
	}

	var pending any
	_ = pending
	pending = any(map[string]any{"node": _hx_tmp_0, "character": _hx_tmp_1, "text": ((any)(text)), "tags": ((any)(tags))})
	this.Hx_Field_pendingDialogue = pending
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var _hx_tmp_2 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_2
	_hx_tmp_2 = this
	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = (character).Value
	var _hx_tmp_4 string
	_ = _hx_tmp_4
	_hx_tmp_4 = text
	var _hx_tmp_5 HxArray[*Hx_Obj_loreline_texttag]
	_ = _hx_tmp_5
	_hx_tmp_5 = tags
	this.Hx_Field_handleDialogue(_hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, func() {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_gthis.Hx_Field_pendingDialogue, pending))) {
			_gthis.Hx_Field_pendingDialogue = nil
		}

		next()
	})
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_presentRestoredDialogue(node *Hx_Obj_loreline_astnode, next func()) bool {
	var restored any
	_ = restored
	restored = this.Hx_Field_pendingDialogue
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(restored, nil))) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(restored, "node"), ((any)(node))))))
	}

	if _hx_tmp_0 {
		return false
	}

	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = node
	var hx_nullable_t66 any
	_ = hx_nullable_t66
	hx_nullable_t66 = Hx_Field_go_haxe_hxdynamic_getField(restored, "character")
	var _hx_tmp_2 HxNullable[string]
	_ = _hx_tmp_2
	if hx_nullable_t66 == nil {
		_hx_tmp_2 = HxNullable[string]{Value: HxStringNull}
	} else {
		_hx_tmp_2 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t66), Valid: true}
	}

	var _hx_tmp_3 string
	_ = _hx_tmp_3
	_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(restored, "text"))
	this.VTable.Hx_Field_presentDialogue(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, HxAnyToTypedArray[*Hx_Obj_loreline_texttag](Hx_Field_go_haxe_hxdynamic_getField(restored, "tags")), next)
	return true
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalChoice(choice *Hx_Obj_loreline_nchoicestatement, next func()) {
	var restoredOptions HxArray[*Hx_Obj_loreline_choiceoption]
	_ = restoredOptions
	restoredOptions = this.Hx_Field_pendingChoiceOptions
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if restoredOptions != nil {
		_hx_tmp_0 = this.VTable.Hx_Field_restoredOptionsMatch(choice, restoredOptions)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		this.Hx_Field_pendingChoiceOptions = nil
		(Hx_Array_Resize[string](this.Hx_Field__choiceEvalTexts, 0))
		(Hx_Array_Resize[bool](this.Hx_Field__choiceEvalEnabled, 0))
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < restoredOptions.Len()) {
					break
				}

				var opt *Hx_Obj_loreline_choiceoption
				_ = opt
				opt = restoredOptions.Get(_g)
				_g++
				_ = (Hx_Array_Push[string](this.Hx_Field__choiceEvalTexts, opt.Hx_Field_text))
				_ = (Hx_Array_Push[bool](this.Hx_Field__choiceEvalEnabled, opt.Hx_Field_enabled))
			}
		}

		this.VTable.Hx_Field_presentChoice(choice, restoredOptions, next)
		return
	}

	this.Hx_Field_pendingChoiceOptions = nil
	var options HxArray[*Hx_Obj_loreline_choiceoption]
	_ = options
	options = HxMakeArray[*Hx_Obj_loreline_choiceoption]()
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var optionsDone *Hx_Obj_loreline_evalnext
	_ = optionsDone
	optionsDone = this.VTable.Hx_Field_wrapNext(func() {
		_gthis.VTable.Hx_Field_presentChoice(choice, options, next)
	})
	var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_2
	_hx_tmp_2 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
	var _hx_tmp_3 *Hx_Obj_loreline_nchoicestatement
	_ = _hx_tmp_3
	_hx_tmp_3 = choice
	var _hx_tmp_4 HxArray[*Hx_Obj_loreline_choiceoption]
	_ = _hx_tmp_4
	_hx_tmp_4 = options
	var _hx_tmp_5 func()
	_ = _hx_tmp_5
	_hx_tmp_5 = optionsDone.Hx_Field_cb
	var _hx_tmp_6 any
	_ = _hx_tmp_6
	_hx_tmp_6 = nil
	this.VTable.Hx_Field_evalChoiceOptionsAndInsertions(_hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, (this.VTable.Hx_Field_get_currentInsertion() == nil))
	optionsDone.Hx_Field_sync = false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoredOptionsMatch(choice *Hx_Obj_loreline_nchoicestatement, options HxArray[*Hx_Obj_loreline_choiceoption]) bool {
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < options.Len()) {
				break
			}

			var option *Hx_Obj_loreline_choiceoption
			_ = option
			option = options.Get(_g)
			_g++
			if option.Hx_Field_node == nil {
				return false
			}

			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if option.Hx_Field_insertion == nil {
				var _hx_tmp_2 HxArray[*Hx_Obj_loreline_nchoiceoption]
				_ = _hx_tmp_2
				_hx_tmp_2 = choice.Hx_Field_options
				var _hx_tmp_3 *Hx_Obj_loreline_nchoiceoption
				_ = _hx_tmp_3
				_hx_tmp_3 = option.Hx_Field_node
				var _hx_tmp_4 HxArray[*Hx_Obj_loreline_nchoiceoption]
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_2
				var _hx_tmp_5 *Hx_Obj_loreline_nchoiceoption
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_3
				_hx_tmp_1 = ((Hx_Array_IndexOf[*Hx_Obj_loreline_nchoiceoption](_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{})) == -1)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				return false
			}
		}
	}

	return true
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_presentChoice(choice *Hx_Obj_loreline_nchoicestatement, options HxArray[*Hx_Obj_loreline_choiceoption], next func()) {
	var currentInsertion *Hx_Obj_loreline_runtimeinsertion
	_ = currentInsertion
	currentInsertion = this.VTable.Hx_Field_get_currentInsertion()
	if ((currentInsertion != nil) && (currentInsertion.Hx_Field_kind == ((int32)(0)))) && (currentInsertion.Hx_Field_options == nil) {
		currentInsertion.Hx_Field_stack = (Hx_Array_Concat[*Hx_Obj_loreline_runtimescope](HxMakeArray[*Hx_Obj_loreline_runtimescope](), this.Hx_Field_stack))
		currentInsertion.Hx_Field_options = options
		next()
		return
	}

	this.Hx_Field_pendingChoiceOptions = options
	var index int32
	_ = index
	index = -1
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var clearChoiceEval func()
	_ = clearChoiceEval
	clearChoiceEval = func() {
		(Hx_Array_Resize[string](_gthis.Hx_Field__choiceEvalTexts, 0))
		(Hx_Array_Resize[bool](_gthis.Hx_Field__choiceEvalEnabled, 0))
	}

	var choiceCallback *Hx_Obj_loreline_evalnext
	_ = choiceCallback
	choiceCallback = this.VTable.Hx_Field_wrapNext(func() {
		_gthis.Hx_Field_pendingChoiceOptions = nil
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if index >= 0 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = index
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _hx_tmp_1
			_hx_tmp_0 = (_hx_tmp_2 < options.Len())
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			var option *Hx_Obj_loreline_choiceoption
			_ = option
			option = options.Get(index)
			if option.Hx_Field_node.Hx_Field_once {
				_gthis.VTable.Hx_Field_setChoiceOptionChosen(option.Hx_Field_node)
			}

			var wrappedNext func()
			_ = wrappedNext
			wrappedNext = func() {
				clearChoiceEval()
				next()
			}

			if option.Hx_Field_insertion != nil {
				var scopeLevel int32
				_ = scopeLevel
				scopeLevel = _gthis.Hx_Field_stack.Len()
				for {
					if !(_gthis.Hx_Field_stack.Len() > 0) {
						break
					}

					_ = (Hx_Array_Pop[*Hx_Obj_loreline_runtimescope](_gthis.Hx_Field_stack).Value)
				}

				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 HxArray[*Hx_Obj_loreline_runtimescope]
					_ = _g1
					_g1 = option.Hx_Field_insertion.Hx_Field_stack
					for {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = _g
						if !(_hx_tmp_3 < _g1.Len()) {
							break
						}

						var scope *Hx_Obj_loreline_runtimescope
						_ = scope
						scope = _g1.Get(_g)
						_g++
						if scope.Hx_Field_insertion != nil {
							scope.Hx_Field_insertion = nil
						}

						_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](_gthis.Hx_Field_stack, scope))
					}
				}

				var _hx_tmp_4 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _hx_tmp_4
				_hx_tmp_4 = _gthis.Hx_Field_stack
				var lastScope *Hx_Obj_loreline_runtimescope
				_ = lastScope
				lastScope = _hx_tmp_4.Get((_gthis.Hx_Field_stack.Len() - 1))
				_gthis.VTable.Hx_Field_push(Hx_Obj_loreline_runtimescope_CreateInstance(nil, lastScope.Hx_Field_beat, _gthis.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getParentNode((func() *Hx_Obj_loreline_node {
					_hx_up := option.Hx_Field_node
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_node
					}
					return nil
				})()).VTable.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode(), nil, nil, (func() *Hx_Obj_loreline_astnode {
					_hx_up := option.Hx_Field_node
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_astnode
					}
					return nil
				})(), nil, nil, nil))
				_gthis.VTable.Hx_Field_resumeFromLevel(scopeLevel, wrappedNext)
			} else {
				_gthis.VTable.Hx_Field_evalChoiceOption(option.Hx_Field_node, wrappedNext)
			}
		} else {
			clearChoiceEval()
			next()
		}
	})
	var _hx_tmp_5 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_5
	_hx_tmp_5 = this
	var _hx_tmp_6 HxArray[*Hx_Obj_loreline_choiceoption]
	_ = _hx_tmp_6
	_hx_tmp_6 = options
	this.Hx_Field_handleChoice(_hx_tmp_5, _hx_tmp_6, func(index_ int32) {
		index = index_
		choiceCallback.Hx_Field_cb()
	})
	choiceCallback.Hx_Field_sync = false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalChoiceOptionsAndInsertions(beat *Hx_Obj_loreline_nbeatdecl, choice *Hx_Obj_loreline_nchoicestatement, result HxArray[*Hx_Obj_loreline_choiceoption], next func(), startIndex__isOptional any, isChoiceRoot__isOptional any) {
	startIndex := HxDefault[HxNullable[int32]]()
	_ = startIndex

	if startIndex__isOptional != nil {
		startIndex = HxConvert[HxNullable[int32]](startIndex__isOptional)
	} else {

	}
	isChoiceRoot := HxDefault[bool]()
	_ = isChoiceRoot

	if isChoiceRoot__isOptional != nil {
		isChoiceRoot = HxConvert[bool](isChoiceRoot__isOptional)
	} else {
		isChoiceRoot = true
	}
	{
		if (startIndex).Valid == false {
			var hx_nullable_t67 HxNullable[int32]
			_ = hx_nullable_t67
			hx_nullable_t67 = HxNullable[int32]{}

			startIndex = HxNullable[int32]{Value: (hx_nullable_t67).Value, Valid: (hx_nullable_t67).Valid}
		}

		var options HxArray[*Hx_Obj_loreline_nchoiceoption]
		_ = options
		options = choice.Hx_Field_options
		if isChoiceRoot {
			(Hx_Array_Resize[string](this.Hx_Field__choiceEvalTexts, 0))
			(Hx_Array_Resize[bool](this.Hx_Field__choiceEvalEnabled, 0))
			{
				var _g int32
				_ = _g
				_g = 0
				for {
					var _hx_tmp_0 int32
					_ = _hx_tmp_0
					_hx_tmp_0 = _g
					if !(_hx_tmp_0 < result.Len()) {
						break
					}

					var opt *Hx_Obj_loreline_choiceoption
					_ = opt
					opt = result.Get(_g)
					_g++
					_ = (Hx_Array_Push[string](this.Hx_Field__choiceEvalTexts, opt.Hx_Field_text))
					_ = (Hx_Array_Push[bool](this.Hx_Field__choiceEvalEnabled, opt.Hx_Field_enabled))
				}
			}
		}

		var _hx_tmp_1 HxNullable[int32]
		_ = _hx_tmp_1
		if (startIndex).Valid != false {
			_hx_tmp_1 = startIndex
		} else {
			_hx_tmp_1 = HxNullable[int32]{Value: 0, Valid: true}
		}

		var index HxNullable[int32]
		_ = index
		index = _hx_tmp_1
		var moveNext func()
		_ = moveNext
		var insertion *Hx_Obj_loreline_runtimeinsertion
		_ = insertion
		var _gthis *Hx_Obj_loreline_interpreter
		_ = _gthis
		_gthis = this
		moveNext = func() {
			if (insertion != nil) && (insertion.Hx_Field_options != nil) {
				{
					var _g int32
					_ = _g
					_g = 0
					var _g1 int32
					_ = _g1
					_g1 = insertion.Hx_Field_options.Len()
					for _g < _g1 {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = _g
						_g = (_g + 1)
						var i int32
						_ = i
						i = _hx_tmp_2
						var opt *Hx_Obj_loreline_choiceoption
						_ = opt
						opt = insertion.Hx_Field_options.Get(i)
						_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](result, opt))
					}
				}

				insertion = nil
			}

			var hx_cmp_21 int32
			_ = hx_cmp_21
			hx_cmp_21 = options.Len()
			if (index).Valid && true && ((index).Value < hx_cmp_21) {
				var _hx_tmp_3 HxArray[*Hx_Obj_loreline_nchoiceoption]
				_ = _hx_tmp_3
				_hx_tmp_3 = options
				var option *Hx_Obj_loreline_nchoiceoption
				_ = option
				option = _hx_tmp_3.Get(((int32)((index).Value)))
				(index).Value = ((index).Value + 1)
				var _hx_tmp_4 bool
				_ = _hx_tmp_4
				if option.Hx_Field_once {
					_hx_tmp_4 = _gthis.VTable.Hx_Field_isChoiceOptionChosen(option)
				} else {
					_hx_tmp_4 = false
				}

				var onceDisabled bool
				_ = onceDisabled
				onceDisabled = _hx_tmp_4
				var _hx_tmp_5 bool
				_ = _hx_tmp_5
				if !onceDisabled {
					var _hx_tmp_6 bool
					_ = _hx_tmp_6
					if option.Hx_Field_condition == nil {
						_hx_tmp_6 = true
					} else {
						_hx_tmp_6 = _gthis.VTable.Hx_Field_evaluateCondition(option.Hx_Field_condition)
					}

					_hx_tmp_5 = (_hx_tmp_6)
				} else {
					_hx_tmp_5 = false
				}

				var enabled bool
				_ = enabled
				enabled = _hx_tmp_5
				if option.Hx_Field_text != nil {
					var done *Hx_Obj_loreline_evalnext
					_ = done
					done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
					var str *Hx_Obj_loreline_nstringliteral
					_ = str
					str = _gthis.VTable.Hx_Field_getTranslatedString((func() *Hx_Obj_loreline_astnode {
						_hx_up := option
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_astnode
						}
						return nil
					})(), option.Hx_Field_text)
					var content any
					_ = content
					content = _gthis.VTable.Hx_Field_evaluateString(str)
					var owner *Hx_Obj_loreline_runtimeinsertion
					_ = owner
					owner = _gthis.VTable.Hx_Field_get_currentInsertion()
					var _hx_tmp_7 HxArray[*Hx_Obj_loreline_choiceoption]
					_ = _hx_tmp_7
					_hx_tmp_7 = result
					var _hx_tmp_8 *Hx_Obj_loreline_runtimeinsertion
					_ = _hx_tmp_8
					if (owner != nil) && (owner.Hx_Field_kind == ((int32)(0))) {
						_hx_tmp_8 = owner
					} else {
						_hx_tmp_8 = nil
					}

					_ = (Hx_Array_Push[*Hx_Obj_loreline_choiceoption](_hx_tmp_7, Hx_Obj_loreline_choiceoption_CreateInstance(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(content, "text")), HxAnyToTypedArray[*Hx_Obj_loreline_texttag](Hx_Field_go_haxe_hxdynamic_getField(content, "tags")), enabled, option, _hx_tmp_8)))
					var _hx_tmp_9 HxArray[string]
					_ = _hx_tmp_9
					_hx_tmp_9 = _gthis.Hx_Field__choiceEvalTexts
					_ = (Hx_Array_Push[string](_hx_tmp_9, Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(content, "text"))))
					_ = (Hx_Array_Push[bool](_gthis.Hx_Field__choiceEvalEnabled, enabled))
					done.Hx_Field_cb()
					done.Hx_Field_sync = false
				} else {
					if option.Hx_Field_insertion != nil {
						var done *Hx_Obj_loreline_evalnext
						_ = done
						done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
						if !enabled {
							done.Hx_Field_cb()
						} else {
							var fh *Hx_Obj_loreline_interpretercontext
							_ = fh
							fh = _gthis.Hx_Field_context
							var postfix int32
							_ = postfix
							postfix = fh.Hx_Field_nextInsertionId
							_gthis.Hx_Field_context.Hx_Field_nextInsertionId = (fh.Hx_Field_nextInsertionId + 1)
							insertion = Hx_Obj_loreline_runtimeinsertion_CreateInstance(postfix, option.Hx_Field_insertion)
							insertion.Hx_Field_parentPartialOptions = (Hx_Array_Concat[*Hx_Obj_loreline_choiceoption](HxMakeArray[*Hx_Obj_loreline_choiceoption](), result))
							insertion.Hx_Field_parentNextOptionIndex = (index).Value
							_gthis.VTable.Hx_Field_evalInsertion(insertion, done.Hx_Field_cb)
						}

						done.Hx_Field_sync = false
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid choice option", option.Hx_Field_pos)))))
					}
				}
			} else {
				next()
			}
		}

		moveNext()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalInsertion(insertion *Hx_Obj_loreline_runtimeinsertion, next func()) {
	var origin *Hx_Obj_loreline_ninsertion
	_ = origin
	origin = insertion.Hx_Field_origin
	var resolvedRef *Hx_Obj_loreline_runtimebeatref
	_ = resolvedRef
	if origin.Hx_Field_targetExpr != nil {
		var value any
		_ = value
		value = this.VTable.Hx_Field_evaluateExpression(origin.Hx_Field_targetExpr)
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = value
		resolvedRef = this.VTable.Hx_Field_resolveDynamicBeatTarget(_hx_tmp_0, (func() *Hx_Obj_loreline_astnode {
			_hx_up := origin
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
	} else {
		var beatName string
		_ = beatName
		beatName = origin.Hx_Field_target
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = beatName
		resolvedRef = this.VTable.Hx_Field_resolveBeatRefFromNode(_hx_tmp_1, (func() *Hx_Obj_loreline_node {
			_hx_up := origin
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		if resolvedRef == nil {
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = beatName
			panic(Hx_Field_haxe_exception_thrown(((any)(this.VTable.Hx_Field_beatNotFoundError(_hx_tmp_2, (func() *Hx_Obj_loreline_astnode {
				_hx_up := origin
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())))))
		}
	}

	var resolvedBeat *Hx_Obj_loreline_nbeatdecl
	_ = resolvedBeat
	resolvedBeat = resolvedRef.Hx_Field_beat
	var argValues HxArray[any]
	_ = argValues
	argValues = this.VTable.Hx_Field_evaluateBeatArgs(resolvedBeat, origin.Hx_Field_args, origin.Hx_Field_pos)
	var _hx_tmp_3 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_3
	_hx_tmp_3 = resolvedBeat
	var _hx_tmp_4 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_4
	_hx_tmp_4 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := resolvedBeat
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_5 HxArray[*Hx_Obj_loreline_astnode]
	_ = _hx_tmp_5
	_hx_tmp_5 = resolvedBeat.Hx_Field_body
	var _hx_tmp_6 *Hx_Obj_loreline_runtimeinsertion
	_ = _hx_tmp_6
	_hx_tmp_6 = insertion
	var _hx_tmp_7 HxArray[any]
	_ = _hx_tmp_7
	_hx_tmp_7 = argValues
	var _hx_tmp_8 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_8
	if resolvedRef.Hx_Field_stack.Len() > 0 {
		_hx_tmp_8 = resolvedRef.Hx_Field_stack
	} else {
		_hx_tmp_8 = nil
	}

	this.VTable.Hx_Field_evalNodeBody(_hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, _hx_tmp_7, _hx_tmp_8, nil, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalBeatCall(call *Hx_Obj_loreline_nbeatcall, next func()) {
	var value any
	_ = value
	value = this.VTable.Hx_Field_evaluateExpression(call.Hx_Field_targetExpr)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = value
	var resolvedRef *Hx_Obj_loreline_runtimebeatref
	_ = resolvedRef
	resolvedRef = this.VTable.Hx_Field_resolveDynamicBeatTarget(_hx_tmp_0, (func() *Hx_Obj_loreline_astnode {
		_hx_up := call
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})())
	var argValues HxArray[any]
	_ = argValues
	argValues = this.VTable.Hx_Field_evaluateBeatArgs(resolvedRef.Hx_Field_beat, call.Hx_Field_args, call.Hx_Field_pos)
	var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_1
	_hx_tmp_1 = resolvedRef.Hx_Field_beat
	var _hx_tmp_2 func()
	_ = _hx_tmp_2
	_hx_tmp_2 = next
	var _hx_tmp_3 HxArray[any]
	_ = _hx_tmp_3
	_hx_tmp_3 = argValues
	var _hx_tmp_4 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_4
	if resolvedRef.Hx_Field_stack.Len() > 0 {
		_hx_tmp_4 = resolvedRef.Hx_Field_stack
	} else {
		_hx_tmp_4 = nil
	}

	this.VTable.Hx_Field_evalBeatRun(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalChoiceOption(option *Hx_Obj_loreline_nchoiceoption, next func()) {
	var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_0
	_hx_tmp_0 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := option
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_2 HxArray[*Hx_Obj_loreline_astnode]
	_ = _hx_tmp_2
	_hx_tmp_2 = option.Hx_Field_body
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = nil
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = nil
	var _hx_tmp_5 any
	_ = _hx_tmp_5
	_hx_tmp_5 = nil
	this.VTable.Hx_Field_evalNodeBody(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, nil, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalIf(ifStmt *Hx_Obj_loreline_nifstatement, next func()) {
	var isTrue bool
	_ = isTrue
	isTrue = this.VTable.Hx_Field_evaluateCondition(ifStmt.Hx_Field_condition)
	var _hx_tmp_0 *Hx_Obj_loreline_nblock
	_ = _hx_tmp_0
	if isTrue {
		_hx_tmp_0 = ifStmt.Hx_Field_thenBranch
	} else {
		_hx_tmp_0 = ifStmt.Hx_Field_elseBranch
	}

	var branch *Hx_Obj_loreline_nblock
	_ = branch
	branch = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if branch != nil {
		_hx_tmp_1 = (branch.Hx_Field_body.Len() > 0)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_2
		_hx_tmp_2 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
		var _hx_tmp_3 *Hx_Obj_loreline_astnode
		_ = _hx_tmp_3
		_hx_tmp_3 = (func() *Hx_Obj_loreline_astnode {
			_hx_up := branch
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})()
		var _hx_tmp_4 HxArray[*Hx_Obj_loreline_astnode]
		_ = _hx_tmp_4
		_hx_tmp_4 = branch.Hx_Field_body
		var _hx_tmp_5 any
		_ = _hx_tmp_5
		_hx_tmp_5 = nil
		var _hx_tmp_6 any
		_ = _hx_tmp_6
		_hx_tmp_6 = nil
		var _hx_tmp_7 any
		_ = _hx_tmp_7
		_hx_tmp_7 = nil
		this.VTable.Hx_Field_evalNodeBody(_hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, _hx_tmp_7, nil, next)
	} else {
		next()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalAlternative(alt *Hx_Obj_loreline_nalternative, next func()) {
	if alt.Hx_Field_items.Len() == 0 {
		next()
		return
	}

	switch alt.Hx_Field_mode {
	case 0:
		{
			var visitCount int32
			_ = visitCount
			visitCount = this.VTable.Hx_Field_getAlternativeVisitCount(alt)
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = visitCount
			if _hx_tmp_1 < alt.Hx_Field_items.Len() {
				_hx_tmp_0 = visitCount
			} else {
				_hx_tmp_0 = (alt.Hx_Field_items.Len() - 1)
			}

			var idx int32
			_ = idx
			idx = _hx_tmp_0
			this.VTable.Hx_Field_setAlternativeVisitCount(alt, (visitCount + 1))
			var item *Hx_Obj_loreline_nblock
			_ = item
			item = alt.Hx_Field_items.Get(idx)
			var _hx_tmp_2 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_2
			_hx_tmp_2 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
			var _hx_tmp_3 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_3
			_hx_tmp_3 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := item
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_4 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_4
			_hx_tmp_4 = item.Hx_Field_body
			var _hx_tmp_5 any
			_ = _hx_tmp_5
			_hx_tmp_5 = nil
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = nil
			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = nil
			this.VTable.Hx_Field_evalNodeBody(_hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, _hx_tmp_7, nil, next)
		}

	case 1:
		{
			var visitCount int32
			_ = visitCount
			visitCount = this.VTable.Hx_Field_getAlternativeVisitCount(alt)
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = visitCount
			var idx int32
			_ = idx
			idx = (_hx_tmp_8 % alt.Hx_Field_items.Len())
			this.VTable.Hx_Field_setAlternativeVisitCount(alt, (visitCount + 1))
			var item *Hx_Obj_loreline_nblock
			_ = item
			item = alt.Hx_Field_items.Get(idx)
			var _hx_tmp_9 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_9
			_hx_tmp_9 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
			var _hx_tmp_10 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_10
			_hx_tmp_10 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := item
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_11 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_11
			_hx_tmp_11 = item.Hx_Field_body
			var _hx_tmp_12 any
			_ = _hx_tmp_12
			_hx_tmp_12 = nil
			var _hx_tmp_13 any
			_ = _hx_tmp_13
			_hx_tmp_13 = nil
			var _hx_tmp_14 any
			_ = _hx_tmp_14
			_hx_tmp_14 = nil
			this.VTable.Hx_Field_evalNodeBody(_hx_tmp_9, _hx_tmp_10, _hx_tmp_11, _hx_tmp_12, _hx_tmp_13, _hx_tmp_14, nil, next)
		}

	case 2:
		{
			var visitCount int32
			_ = visitCount
			visitCount = this.VTable.Hx_Field_getAlternativeVisitCount(alt)
			var _hx_tmp_15 int32
			_ = _hx_tmp_15
			_hx_tmp_15 = visitCount
			if _hx_tmp_15 >= alt.Hx_Field_items.Len() {
				next()
				return
			}

			this.VTable.Hx_Field_setAlternativeVisitCount(alt, (visitCount + 1))
			var item *Hx_Obj_loreline_nblock
			_ = item
			item = alt.Hx_Field_items.Get(visitCount)
			var _hx_tmp_16 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_16
			_hx_tmp_16 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
			var _hx_tmp_17 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_17
			_hx_tmp_17 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := item
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_18 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_18
			_hx_tmp_18 = item.Hx_Field_body
			var _hx_tmp_19 any
			_ = _hx_tmp_19
			_hx_tmp_19 = nil
			var _hx_tmp_20 any
			_ = _hx_tmp_20
			_hx_tmp_20 = nil
			var _hx_tmp_21 any
			_ = _hx_tmp_21
			_hx_tmp_21 = nil
			this.VTable.Hx_Field_evalNodeBody(_hx_tmp_16, _hx_tmp_17, _hx_tmp_18, _hx_tmp_19, _hx_tmp_20, _hx_tmp_21, nil, next)
		}

	case 3:
		{
			var idx int32
			_ = idx
			idx = this.Hx_Field_builtins.VTable.Hx_Field_random(0, (alt.Hx_Field_items.Len() - 1))
			var item *Hx_Obj_loreline_nblock
			_ = item
			item = alt.Hx_Field_items.Get(idx)
			var _hx_tmp_22 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_22
			_hx_tmp_22 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
			var _hx_tmp_23 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_23
			_hx_tmp_23 = (func() *Hx_Obj_loreline_astnode {
				_hx_up := item
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})()
			var _hx_tmp_24 HxArray[*Hx_Obj_loreline_astnode]
			_ = _hx_tmp_24
			_hx_tmp_24 = item.Hx_Field_body
			var _hx_tmp_25 any
			_ = _hx_tmp_25
			_hx_tmp_25 = nil
			var _hx_tmp_26 any
			_ = _hx_tmp_26
			_hx_tmp_26 = nil
			var _hx_tmp_27 any
			_ = _hx_tmp_27
			_hx_tmp_27 = nil
			this.VTable.Hx_Field_evalNodeBody(_hx_tmp_22, _hx_tmp_23, _hx_tmp_24, _hx_tmp_25, _hx_tmp_26, _hx_tmp_27, nil, next)
		}

	case 4:
		{
			var _g HxArray[int32]
			_ = _g
			_g = HxMakeArray[int32]()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 int32
				_ = _g2
				_g2 = alt.Hx_Field_items.Len()
				for _g1 < _g2 {
					var _hx_tmp_28 int32
					_ = _hx_tmp_28
					_hx_tmp_28 = _g1
					_g1 = (_g1 + 1)
					var i int32
					_ = i
					i = _hx_tmp_28
					_ = (Hx_Array_Push[int32](_g, i))
				}
			}

			var indices HxArray[int32]
			_ = indices
			indices = _g
			var i int32
			_ = i
			i = (indices.Len() - 1)
			for i > 0 {
				var j int32
				_ = j
				j = this.Hx_Field_builtins.VTable.Hx_Field_random(0, i)
				var tmp int32
				_ = tmp
				tmp = indices.Get(i)
				indices.Set(i, indices.Get(j))
				indices.Set(j, tmp)
				i--
			}

			this.VTable.Hx_Field_evalShuffledItems(alt, indices, 0, next)
		}

	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalShuffledItems(alt *Hx_Obj_loreline_nalternative, indices HxArray[int32], idx int32, next func()) {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = idx
	if _hx_tmp_0 >= indices.Len() {
		next()
		return
	}

	var item *Hx_Obj_loreline_nblock
	_ = item
	item = alt.Hx_Field_items.Get(indices.Get(idx))
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_1
	_hx_tmp_1 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
	var _hx_tmp_2 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_2
	_hx_tmp_2 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := item
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_3 HxArray[*Hx_Obj_loreline_astnode]
	_ = _hx_tmp_3
	_hx_tmp_3 = item.Hx_Field_body
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = nil
	var _hx_tmp_5 any
	_ = _hx_tmp_5
	_hx_tmp_5 = nil
	var _hx_tmp_6 any
	_ = _hx_tmp_6
	_hx_tmp_6 = nil
	var _hx_tmp_7 HxArray[int32]
	_ = _hx_tmp_7
	_hx_tmp_7 = indices
	this.VTable.Hx_Field_evalNodeBody(_hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, _hx_tmp_7, func() {
		_gthis.VTable.Hx_Field_evalShuffledItems(alt, indices, (idx + 1), next)
	})
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_restoreShuffleOrder(saved HxArray[int32], item *Hx_Obj_loreline_nblock) HxArray[int32] {
	var alt *Hx_Obj_loreline_nalternative
	_ = alt
	alt = func() func(value *Hx_Obj_loreline_node, c *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nalternative {
		var _hx_closure_src_637 func(value any, c *Hx_Obj_go_haxe_hxclass) any
		_ = _hx_closure_src_637
		_hx_closure_src_637 = Hx_Field_std_downcast
		return func(value *Hx_Obj_loreline_node, c *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nalternative {
			var hx_dyncast_247 any
			_ = hx_dyncast_247
			hx_dyncast_247 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_637(((any)(value)), c), "Hx_Obj_loreline_nalternative")
			var _hx_tmp_0 *Hx_Obj_loreline_nalternative
			_ = _hx_tmp_0
			if hx_dyncast_247 != nil {
				_hx_tmp_0 = hx_dyncast_247.(interface {
					Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative
				}).Hx_As_Hx_Obj_loreline_nalternative()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getParentNode((func() *Hx_Obj_loreline_node {
		_hx_up := item
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nalternative_RTTI)))
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (alt == nil) || (alt.Hx_Field_mode != ((int32)(4))) {
		_hx_tmp_1 = true
	} else {
		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = saved.Len()
		var _hx_tmp_3 int32
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_1 = (_hx_tmp_3 != alt.Hx_Field_items.Len())
	}

	if _hx_tmp_1 {
		return nil
	}

	var order HxArray[int32]
	_ = order
	order = HxMakeArray[int32]()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = saved.Len()
		for _g < _g1 {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_4
			var index int32
			_ = index
			index = saved.Get(i)
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			var _hx_tmp_6 bool
			_ = _hx_tmp_6
			if index < 0 {
				_hx_tmp_6 = true
			} else {
				var _hx_tmp_7 int32
				_ = _hx_tmp_7
				_hx_tmp_7 = index
				var _hx_tmp_8 int32
				_ = _hx_tmp_8
				_hx_tmp_8 = _hx_tmp_7
				_hx_tmp_6 = (_hx_tmp_8 >= saved.Len())
			}

			if _hx_tmp_6 {
				_hx_tmp_5 = true
			} else {
				var _hx_tmp_9 HxArray[int32]
				_ = _hx_tmp_9
				_hx_tmp_9 = order
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = index
				var _hx_tmp_11 HxArray[int32]
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_9
				var _hx_tmp_12 int32
				_ = _hx_tmp_12
				_hx_tmp_12 = _hx_tmp_10
				_hx_tmp_5 = ((Hx_Array_IndexOf[int32](_hx_tmp_11, _hx_tmp_12, HxNullable[int32]{})) != -1)
			}

			if _hx_tmp_5 {
				return nil
			}

			_ = (Hx_Array_Push[int32](order, index))
		}
	}

	return order
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getBeatVisitCount(beat *Hx_Obj_loreline_nbeatdecl) int32 {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = beat.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_638 func(high int32, low int32) any
		_ = _hx_closure_src_638
		_hx_closure_src_638 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_248 any
			_ = hx_dyncast_248
			hx_dyncast_248 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_638(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_248 != nil {
				_hx_tmp_0 = hx_dyncast_248.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		return 0
	}

	var count any
	_ = count
	count = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, "_visitCount")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(count, nil))) {
		return 0
	}

	return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(count)))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_incrementBeatVisitCount(beat *Hx_Obj_loreline_nbeatdecl) {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = beat.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_639 func(high int32, low int32) any
		_ = _hx_closure_src_639
		_hx_closure_src_639 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_249 any
			_ = hx_dyncast_249
			hx_dyncast_249 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_639(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_249 != nil {
				_hx_tmp_0 = hx_dyncast_249.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
			_hx_up := beat
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), nil, nil)
		{
			var tmp_key_1 Hx_Typedef_haxe__int64___int64
			_ = tmp_key_1
			tmp_key_1 = beat.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(tmp_key_1))
			func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
				var _hx_closure_src_640 func(high int32, low int32, value any)
				_ = _hx_closure_src_640
				_hx_closure_src_640 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					_hx_closure_src_640(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), state)
		}
	}

	var tmp any
	_ = tmp
	tmp = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, "_visitCount")
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(tmp, nil))) {
		_hx_tmp_3 = tmp
	} else {
		_hx_tmp_3 = ((int32)(0))
	}

	var count int32
	_ = count
	count = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(_hx_tmp_3)))
	Hx_Field_loreline_objects_setField(this, state.Hx_Field_fields, "_visitCount", (count + 1))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalWhen(when *Hx_Obj_loreline_nwhenstatement, next func()) {
	var owner *Hx_Obj_loreline_runtimeinsertion
	_ = owner
	owner = this.VTable.Hx_Field_collectingWhenInsertion((this.Hx_Field_stack.Len() - 1))
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if owner != nil {
		_hx_tmp_0 = owner.Hx_Field_whenFirst
	} else {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if (when.Hx_Field_strategy).Valid == true {
			_hx_tmp_1 = ((when.Hx_Field_strategy).Value == "first")
		} else {
			_hx_tmp_1 = false
		}

		_hx_tmp_0 = (_hx_tmp_1)
	}

	var stopAtFirst bool
	_ = stopAtFirst
	stopAtFirst = _hx_tmp_0
	var candidates HxArray[any]
	_ = candidates
	candidates = HxMakeArray[any]().Dyn()
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	var done *Hx_Obj_loreline_evalnext
	_ = done
	done = this.VTable.Hx_Field_wrapNext(func() {
		_gthis.VTable.Hx_Field_finishWhenCollection(when, candidates, owner, next)
	})
	this.VTable.Hx_Field_collectWhenRules(when, candidates, 0, stopAtFirst, owner, done.Hx_Field_cb)
	done.Hx_Field_sync = false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_collectingWhenInsertion(scopeLevel int32) *Hx_Obj_loreline_runtimeinsertion {
	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	insertion = this.VTable.Hx_Field_insertionAtOrBelow(scopeLevel)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if (insertion != nil) && (insertion.Hx_Field_kind == ((int32)(1))) {
		_hx_tmp_0 = !insertion.VTable.Hx_Field_isCollected()
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		return insertion
	} else {
		return nil
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_finishWhenCollection(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any], owner *Hx_Obj_loreline_runtimeinsertion, next func()) {
	if owner != nil {
		owner.Hx_Field_stack = (Hx_Array_Concat[*Hx_Obj_loreline_runtimescope](HxMakeArray[*Hx_Obj_loreline_runtimescope](), this.Hx_Field_stack))
		owner.Hx_Field_rules = candidates
		next()
		return
	}

	var chosen any
	_ = chosen
	chosen = this.VTable.Hx_Field_chooseWhenCandidate(when, candidates)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(chosen, nil))) {
		next()
	} else {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(chosen, "insertion"), ((any)(nil))))) {
			var hx_dyncast_250 any
			_ = hx_dyncast_250
			hx_dyncast_250 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(chosen, "rule"), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_0 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_0
			if hx_dyncast_250 != nil {
				_hx_tmp_0 = hx_dyncast_250.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_0 = nil
			}

			this.VTable.Hx_Field_playWhenRule(_hx_tmp_0, next)
		} else {
			this.VTable.Hx_Field_playInsertedWhenRule(chosen, next)
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_collectWhenRules(when *Hx_Obj_loreline_nwhenstatement, result HxArray[any], startIndex int32, stopAtFirst bool, owner *Hx_Obj_loreline_runtimeinsertion, next func()) {
	var index int32
	_ = index
	index = startIndex
	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	var moveNext func()
	_ = moveNext
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	moveNext = func() {
		if (insertion != nil) && (insertion.Hx_Field_rules != nil) {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[any]
			_ = _g1
			_g1 = insertion.Hx_Field_rules
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
					break
				}

				var candidate any
				_ = candidate
				candidate = _g1.Get(_g)
				_g++
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](result), candidate))
			}
		}

		insertion = nil
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		if stopAtFirst {
			_hx_tmp_1 = Hx_Field_loreline_interpreter_hasEligibleCandidate(result)
		} else {
			_hx_tmp_1 = false
		}

		if _hx_tmp_1 {
			next()
			return
		}

		var _hx_tmp_2 int32
		_ = _hx_tmp_2
		_hx_tmp_2 = index
		if _hx_tmp_2 >= when.Hx_Field_rules.Len() {
			next()
			return
		}

		var rule *Hx_Obj_loreline_nwhenrule
		_ = rule
		rule = when.Hx_Field_rules.Get(index)
		index = (index + 1)
		var done *Hx_Obj_loreline_evalnext
		_ = done
		done = _gthis.VTable.Hx_Field_wrapNext(moveNext)
		if rule.Hx_Field_insertion != nil {
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if rule.Hx_Field_insertionCondition == nil {
				_hx_tmp_3 = true
			} else {
				_hx_tmp_3 = _gthis.VTable.Hx_Field_evaluateCondition(rule.Hx_Field_insertionCondition)
			}

			if _hx_tmp_3 {
				var fh *Hx_Obj_loreline_interpretercontext
				_ = fh
				fh = _gthis.Hx_Field_context
				var postfix int32
				_ = postfix
				postfix = fh.Hx_Field_nextInsertionId
				_gthis.Hx_Field_context.Hx_Field_nextInsertionId = (fh.Hx_Field_nextInsertionId + 1)
				insertion = Hx_Obj_loreline_runtimeinsertion_CreateInstance(postfix, rule.Hx_Field_insertion)
				insertion.Hx_Field_kind = ((int32)(1))
				insertion.Hx_Field_whenFirst = stopAtFirst
				insertion.Hx_Field_parentPartialRules = (Hx_Array_Concat[any](HxMakeArrayView[any](HxMakeArray[any]().Dyn()), result))
				insertion.Hx_Field_parentNextRuleIndex = index
				_gthis.VTable.Hx_Field_evalInsertion(insertion, done.Hx_Field_cb)
			} else {
				done.Hx_Field_cb()
			}
		} else {
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if rule.Hx_Field_once {
				_hx_tmp_4 = _gthis.VTable.Hx_Field_isWhenRuleConsumed(rule)
			} else {
				_hx_tmp_4 = false
			}

			var consumed bool
			_ = consumed
			consumed = _hx_tmp_4
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if !consumed {
				var _hx_tmp_6 bool
				_ = _hx_tmp_6
				if rule.Hx_Field_condition == nil {
					_hx_tmp_6 = true
				} else {
					_hx_tmp_6 = _gthis.VTable.Hx_Field_evaluateCondition(rule.Hx_Field_condition)
				}

				_hx_tmp_5 = (_hx_tmp_6)
			} else {
				_hx_tmp_5 = false
			}

			var eligible bool
			_ = eligible
			eligible = _hx_tmp_5
			var _hx_tmp_7 HxArray[any]
			_ = _hx_tmp_7
			_hx_tmp_7 = HxMakeArrayView[any](result)
			_ = (Hx_Array_Push[any](_hx_tmp_7, any(map[string]any{"rule": ((any)(rule)), "eligible": ((any)(eligible)), "insertion": ((any)(owner))})))
			done.Hx_Field_cb()
		}

		done.Hx_Field_sync = false
	}

	moveNext()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_chooseWhenCandidate(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any]) any {
	var strategy HxNullable[string]
	_ = strategy
	strategy = when.Hx_Field_strategy
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if (strategy).Valid != false {
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if (strategy).Valid == false {
			_hx_tmp_2 = true
		} else {
			_hx_tmp_2 = ((strategy).Value != "first")
		}

		_hx_tmp_1 = (_hx_tmp_2)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if (strategy).Valid == false {
			_hx_tmp_3 = true
		} else {
			_hx_tmp_3 = ((strategy).Value != "pick")
		}

		_hx_tmp_0 = (_hx_tmp_3)
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		var hx_nullable_f212 HxNullable[string]
		_ = hx_nullable_f212
		hx_nullable_f212 = strategy
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		if (hx_nullable_f212).Valid == true {
			_hx_tmp_4 = (hx_nullable_f212).Value
		} else {
			_hx_tmp_4 = nil
		}

		var customStrategy any
		_ = customStrategy
		customStrategy = this.Hx_Field_topLevelFunctions.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(_hx_tmp_4)
		var _hx_tmp_5 bool
		_ = _hx_tmp_5
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(customStrategy, nil))) {
			_hx_tmp_5 = true
		} else {
			_hx_tmp_5 = !Hx_Field_reflect_isFunction(customStrategy)
		}

		if _hx_tmp_5 {
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = when.Hx_Field_strategyPos
			var _hx_tmp_6 *Hx_Obj_loreline_position
			_ = _hx_tmp_6
			if pos != nil {
				_hx_tmp_6 = pos
			} else {
				_hx_tmp_6 = when.Hx_Field_pos
			}

			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unknown when strategy: ", (strategy).Value), _hx_tmp_6)))))
		}

		return this.VTable.Hx_Field_customStrategyCandidate(when, candidates, customStrategy)
	}

	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _g1
			if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(candidates)) {
				break
			}

			var candidate any
			_ = candidate
			candidate = candidates.Get(_g1)
			_g1++
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(candidate, "eligible")))) {
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](_g), candidate))
			}
		}
	}

	var eligible HxArray[any]
	_ = eligible
	eligible = _g
	if Hx_Field_go_haxe_hxdynamic_getArrayLength(eligible) == 0 {
		return nil
	}

	if (strategy).Valid == false {
		return this.VTable.Hx_Field_mostSalientCandidate(eligible)
	} else {
		switch (strategy).Value {
		case "first":
			{
				return eligible.Get(0)
			}

		case "pick":
			{
				var _hx_tmp_8 HxArray[any]
				_ = _hx_tmp_8
				_hx_tmp_8 = eligible
				return _hx_tmp_8.Get(this.Hx_Field_builtins.VTable.Hx_Field_random(0, (Hx_Field_go_haxe_hxdynamic_getArrayLength(eligible) - 1)))
			}

		default:
			{
				return this.VTable.Hx_Field_mostSalientCandidate(eligible)
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_mostSalientCandidate(eligible HxArray[any]) any {
	var best any
	_ = best
	var bestScore int32
	_ = bestScore
	bestScore = 0
	var bestLastPlayed int32
	_ = bestLastPlayed
	bestLastPlayed = 0
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(eligible)) {
				break
			}

			var candidate any
			_ = candidate
			candidate = eligible.Get(_g)
			_g++
			var hx_dyncast_251 any
			_ = hx_dyncast_251
			hx_dyncast_251 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_1 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_1
			if hx_dyncast_251 != nil {
				_hx_tmp_1 = hx_dyncast_251.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_1 = nil
			}

			var score int32
			_ = score
			score = Hx_Field_loreline_astutils_whenRuleScore(_hx_tmp_1)
			var hx_dyncast_252 any
			_ = hx_dyncast_252
			hx_dyncast_252 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_2 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_2
			if hx_dyncast_252 != nil {
				_hx_tmp_2 = hx_dyncast_252.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_2 = nil
			}

			var lastPlayed int32
			_ = lastPlayed
			lastPlayed = this.VTable.Hx_Field_getWhenRuleLastPlayed(_hx_tmp_2)
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(best, nil))) {
				_hx_tmp_4 = true
			} else {
				_hx_tmp_4 = (score > bestScore)
			}

			if _hx_tmp_4 {
				_hx_tmp_3 = true
			} else {
				_hx_tmp_3 = ((score == bestScore) && (lastPlayed < bestLastPlayed))
			}

			if _hx_tmp_3 {
				best = candidate
				bestScore = score
				bestLastPlayed = lastPlayed
			}
		}
	}

	return best
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_customStrategyCandidate(when *Hx_Obj_loreline_nwhenstatement, candidates HxArray[any], strategy any) any {
	var records any
	_ = records
	records = Hx_Field_loreline_arrays_createArray()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = Hx_Field_go_haxe_hxdynamic_getArrayLength(candidates)
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var candidate any
			_ = candidate
			candidate = candidates.Get(i)
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = nil
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = nil
			var record any
			_ = record
			record = Hx_Field_loreline_objects_createFields(_hx_tmp_1, _hx_tmp_2, nil)
			Hx_Field_loreline_objects_setField(this, record, "index", i)
			var _hx_tmp_3 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_3
			_hx_tmp_3 = this
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = record
			Hx_Field_loreline_objects_setField(_hx_tmp_3, _hx_tmp_4, "eligible", Hx_Field_go_haxe_hxdynamic_getField(candidate, "eligible"))
			var _hx_tmp_5 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_5
			_hx_tmp_5 = this
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = record
			var hx_dyncast_253 any
			_ = hx_dyncast_253
			hx_dyncast_253 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_7 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_7
			if hx_dyncast_253 != nil {
				_hx_tmp_7 = hx_dyncast_253.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_7 = nil
			}

			Hx_Field_loreline_objects_setField(_hx_tmp_5, _hx_tmp_6, "criteria", Hx_Field_loreline_astutils_whenRuleScore(_hx_tmp_7))
			var _hx_tmp_8 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_8
			_hx_tmp_8 = this
			var _hx_tmp_9 any
			_ = _hx_tmp_9
			_hx_tmp_9 = record
			var hx_dyncast_254 any
			_ = hx_dyncast_254
			hx_dyncast_254 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_astnode")
			var _hx_tmp_10 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_10
			if hx_dyncast_254 != nil {
				_hx_tmp_10 = hx_dyncast_254.(interface {
					Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
				}).Hx_As_Hx_Obj_loreline_astnode()
			} else {
				_hx_tmp_10 = nil
			}

			Hx_Field_loreline_objects_setField(_hx_tmp_8, _hx_tmp_9, "played", this.VTable.Hx_Field_getNodeStateInt(_hx_tmp_10, "_played"))
			var _hx_tmp_11 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_11
			_hx_tmp_11 = this
			var _hx_tmp_12 any
			_ = _hx_tmp_12
			_hx_tmp_12 = record
			var hx_dyncast_255 any
			_ = hx_dyncast_255
			hx_dyncast_255 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_nwhenrule")
			var _hx_tmp_13 *Hx_Obj_loreline_nwhenrule
			_ = _hx_tmp_13
			if hx_dyncast_255 != nil {
				_hx_tmp_13 = hx_dyncast_255.(interface {
					Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule
				}).Hx_As_Hx_Obj_loreline_nwhenrule()
			} else {
				_hx_tmp_13 = nil
			}

			Hx_Field_loreline_objects_setField(_hx_tmp_11, _hx_tmp_12, "lastPlayed", this.VTable.Hx_Field_getWhenRuleLastPlayed(_hx_tmp_13))
			var _hx_tmp_14 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_14
			_hx_tmp_14 = this
			var _hx_tmp_15 any
			_ = _hx_tmp_15
			_hx_tmp_15 = record
			Hx_Field_loreline_objects_setField(_hx_tmp_14, _hx_tmp_15, "ephemeral", Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "once"))
			Hx_Field_loreline_arrays_arrayPush(records, record)
		}
	}

	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = when.Hx_Field_strategyPos
	var _hx_tmp_16 *Hx_Obj_loreline_position
	_ = _hx_tmp_16
	if pos != nil {
		_hx_tmp_16 = pos
	} else {
		_hx_tmp_16 = when.Hx_Field_pos
	}

	var pos1 *Hx_Obj_loreline_position
	_ = pos1
	pos1 = _hx_tmp_16
	var _hx_tmp_17 any
	_ = _hx_tmp_17
	_hx_tmp_17 = strategy
	var result any
	_ = result
	result = this.VTable.Hx_Field_callFunctionValue(_hx_tmp_17, HxMakeArray[any](records).Dyn(), pos1, nil)
	var _hx_tmp_18 HxNullable[int32]
	_ = _hx_tmp_18
	if Hx_Field_std_isOfType(result, Hx_Abstract_int_RTTI) {
		_hx_tmp_18 = HxNullable[int32]{Value: ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(result))), Valid: true}
	} else {
		var _hx_tmp_19 HxNullable[int32]
		_ = _hx_tmp_19
		var _hx_tmp_20 bool
		_ = _hx_tmp_20
		if Hx_Field_std_isOfType(result, Hx_Abstract_float_RTTI) {
			var _hx_tmp_21 float64
			_ = _hx_tmp_21
			_hx_tmp_21 = Hx_Field_math_ffloor(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(result))))
			var _hx_tmp_22 float64
			_ = _hx_tmp_22
			_hx_tmp_22 = _hx_tmp_21
			_hx_tmp_20 = (_hx_tmp_22 == ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(result))))
		} else {
			_hx_tmp_20 = false
		}

		if _hx_tmp_20 {
			_hx_tmp_19 = HxNullable[int32]{Value: Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(result)))), Valid: true}
		} else {
			_hx_tmp_19 = HxNullable[int32]{}
		}

		_hx_tmp_18 = _hx_tmp_19
	}

	var index HxNullable[int32]
	_ = index
	index = _hx_tmp_18
	var _hx_tmp_23 bool
	_ = _hx_tmp_23
	if (index).Valid == true {
		_hx_tmp_23 = ((index).Value == -1)
	} else {
		_hx_tmp_23 = false
	}

	if _hx_tmp_23 {
		return nil
	}

	var _hx_tmp_24 bool
	_ = _hx_tmp_24
	var _hx_tmp_25 bool
	_ = _hx_tmp_25
	if (index).Valid == false {
		_hx_tmp_25 = true
	} else {
		_hx_tmp_25 = ((index).Valid && true && ((index).Value < 0))
	}

	if _hx_tmp_25 {
		_hx_tmp_24 = true
	} else {
		var hx_cmp_22 int32
		_ = hx_cmp_22
		hx_cmp_22 = Hx_Field_go_haxe_hxdynamic_getArrayLength(candidates)
		_hx_tmp_24 = ((index).Valid && true && ((index).Value >= hx_cmp_22) || (!(index).Valid && !true))
	}

	if _hx_tmp_24 {
		var _hx_tmp_26 string
		_ = _hx_tmp_26
		_hx_tmp_26 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Strategy ", (when.Hx_Field_strategy).Value), " returned ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_26, this.VTable.Hx_Field_valueToString(result)), ", expected -1 or the index of an eligible rule"), pos1)))))
	}

	var _hx_tmp_27 HxArray[any]
	_ = _hx_tmp_27
	_hx_tmp_27 = candidates
	if !((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_getField(_hx_tmp_27.Get(((int32)((index).Value))), "eligible")))) {
		var _hx_tmp_28 string
		_ = _hx_tmp_28
		_hx_tmp_28 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Strategy ", (when.Hx_Field_strategy).Value), " returned ")
		var hx_nullable_f213 HxNullable[int32]
		_ = hx_nullable_f213
		hx_nullable_f213 = index
		var _hx_tmp_29 string
		_ = _hx_tmp_29
		if (hx_nullable_f213).Valid == true {
			_hx_tmp_29 = Hx_Field_go_haxe_hxdynamic_toString((hx_nullable_f213).Value)
		} else {
			_hx_tmp_29 = "null"
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_28, _hx_tmp_29), ", which is not an eligible rule"), pos1)))))
	}

	var _hx_tmp_30 HxArray[any]
	_ = _hx_tmp_30
	_hx_tmp_30 = candidates
	return _hx_tmp_30.Get(((int32)((index).Value)))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_playWhenRule(rule *Hx_Obj_loreline_nwhenrule, next func()) {
	this.VTable.Hx_Field_recordWhenRulePlayed(rule)
	var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_0
	_hx_tmp_0 = this.VTable.Hx_Field_get_currentScope().Hx_Field_beat
	var _hx_tmp_1 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	var _hx_tmp_2 HxArray[*Hx_Obj_loreline_astnode]
	_ = _hx_tmp_2
	_hx_tmp_2 = rule.Hx_Field_body
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = nil
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = nil
	var _hx_tmp_5 any
	_ = _hx_tmp_5
	_hx_tmp_5 = nil
	this.VTable.Hx_Field_evalNodeBody(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, nil, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_recordWhenRulePlayed(rule *Hx_Obj_loreline_nwhenrule) {
	var tick int32
	_ = tick
	tick = (this.VTable.Hx_Field_getNodeStateInt((func() *Hx_Obj_loreline_astnode {
		_hx_up := this.Hx_Field_context.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), "_whenTick") + 1)
	this.VTable.Hx_Field_setNodeStateField((func() *Hx_Obj_loreline_astnode {
		_hx_up := this.Hx_Field_context.Hx_Field_script
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), "_whenTick", tick)
	var _hx_tmp_0 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_0
	_hx_tmp_0 = (func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})()
	this.VTable.Hx_Field_setNodeStateField(_hx_tmp_0, "_played", (this.VTable.Hx_Field_getNodeStateInt((func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), "_played") + 1))
	this.VTable.Hx_Field_setNodeStateField((func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), "_lastPlayed", tick)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_playInsertedWhenRule(candidate any, next func()) {
	var hx_dyncast_258 any
	_ = hx_dyncast_258
	hx_dyncast_258 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "insertion"), "Hx_Obj_loreline_runtimeinsertion")
	var _hx_tmp_0 *Hx_Obj_loreline_runtimeinsertion
	_ = _hx_tmp_0
	if hx_dyncast_258 != nil {
		_hx_tmp_0 = hx_dyncast_258.(interface {
			Hx_As_Hx_Obj_loreline_runtimeinsertion() *Hx_Obj_loreline_runtimeinsertion
		}).Hx_As_Hx_Obj_loreline_runtimeinsertion()
	} else {
		_hx_tmp_0 = nil
	}

	var insertion *Hx_Obj_loreline_runtimeinsertion
	_ = insertion
	insertion = _hx_tmp_0
	var scopeLevel int32
	_ = scopeLevel
	scopeLevel = this.Hx_Field_stack.Len()
	for {
		if !(this.Hx_Field_stack.Len() > 0) {
			break
		}

		_ = (Hx_Array_Pop[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack).Value)
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_runtimescope]
		_ = _g1
		_g1 = insertion.Hx_Field_stack
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var scope *Hx_Obj_loreline_runtimescope
			_ = scope
			scope = _g1.Get(_g)
			_g++
			if scope.Hx_Field_insertion != nil {
				scope.Hx_Field_insertion = nil
			}

			_ = (Hx_Array_Push[*Hx_Obj_loreline_runtimescope](this.Hx_Field_stack, scope))
		}
	}

	var _hx_tmp_2 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_stack
	var lastScope *Hx_Obj_loreline_runtimescope
	_ = lastScope
	lastScope = _hx_tmp_2.Get((this.Hx_Field_stack.Len() - 1))
	var hx_dyncast_257 any
	_ = hx_dyncast_257
	hx_dyncast_257 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_node")
	var _hx_tmp_3 *Hx_Obj_loreline_node
	_ = _hx_tmp_3
	if hx_dyncast_257 != nil {
		_hx_tmp_3 = hx_dyncast_257.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
	} else {
		_hx_tmp_3 = nil
	}

	var hx_dyncast_256 any
	_ = hx_dyncast_256
	hx_dyncast_256 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(candidate, "rule"), "Hx_Obj_loreline_astnode")
	var _hx_tmp_4 *Hx_Obj_loreline_astnode
	_ = _hx_tmp_4
	if hx_dyncast_256 != nil {
		_hx_tmp_4 = hx_dyncast_256.(interface {
			Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
		}).Hx_As_Hx_Obj_loreline_astnode()
	} else {
		_hx_tmp_4 = nil
	}

	this.VTable.Hx_Field_push(Hx_Obj_loreline_runtimescope_CreateInstance(nil, lastScope.Hx_Field_beat, this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getParentNode(_hx_tmp_3).VTable.(interface {
		Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
	}).Hx_As_Hx_Obj_loreline_astnode(), nil, nil, _hx_tmp_4, nil, nil, nil))
	this.VTable.Hx_Field_resumeFromLevel(scopeLevel, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isWhenRuleConsumed(rule *Hx_Obj_loreline_nwhenrule) bool {
	return (this.VTable.Hx_Field_getNodeStateInt((func() *Hx_Obj_loreline_astnode {
		_hx_up := rule
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_astnode
		}
		return nil
	})(), "_played") > 0)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getWhenRuleLastPlayed(rule *Hx_Obj_loreline_nwhenrule) int32 {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = rule.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_641 func(high int32, low int32) any
		_ = _hx_closure_src_641
		_hx_closure_src_641 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_259 any
			_ = hx_dyncast_259
			hx_dyncast_259 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_641(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_259 != nil {
				_hx_tmp_0 = hx_dyncast_259.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		return -1
	}

	var value any
	_ = value
	value = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, "_lastPlayed")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return -1
	}

	return Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(value))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getNodeStateInt(node *Hx_Obj_loreline_astnode, field string) int32 {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = node.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_642 func(high int32, low int32) any
		_ = _hx_closure_src_642
		_hx_closure_src_642 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_260 any
			_ = hx_dyncast_260
			hx_dyncast_260 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_642(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_260 != nil {
				_hx_tmp_0 = hx_dyncast_260.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		return 0
	}

	var value any
	_ = value
	value = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, field)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return 0
	}

	return Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(value))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setNodeStateField(node *Hx_Obj_loreline_astnode, field string, value any) {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = node.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_643 func(high int32, low int32) any
		_ = _hx_closure_src_643
		_hx_closure_src_643 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_261 any
			_ = hx_dyncast_261
			hx_dyncast_261 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_643(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_261 != nil {
				_hx_tmp_0 = hx_dyncast_261.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
			_hx_up := node
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), nil, nil)
		{
			var tmp_key_1 Hx_Typedef_haxe__int64___int64
			_ = tmp_key_1
			tmp_key_1 = node.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(tmp_key_1))
			func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
				var _hx_closure_src_644 func(high int32, low int32, value any)
				_ = _hx_closure_src_644
				_hx_closure_src_644 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					_hx_closure_src_644(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), state)
		}
	}

	Hx_Field_loreline_objects_setField(this, state.Hx_Field_fields, field, value)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getAlternativeVisitCount(alt *Hx_Obj_loreline_nalternative) int32 {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = alt.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_645 func(high int32, low int32) any
		_ = _hx_closure_src_645
		_hx_closure_src_645 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_262 any
			_ = hx_dyncast_262
			hx_dyncast_262 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_645(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_262 != nil {
				_hx_tmp_0 = hx_dyncast_262.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		return 0
	}

	var count any
	_ = count
	count = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, "_visitCount")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(count, nil))) {
		return 0
	}

	return ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(count)))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setAlternativeVisitCount(alt *Hx_Obj_loreline_nalternative, count int32) {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = alt.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_646 func(high int32, low int32) any
		_ = _hx_closure_src_646
		_hx_closure_src_646 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_263 any
			_ = hx_dyncast_263
			hx_dyncast_263 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_646(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_263 != nil {
				_hx_tmp_0 = hx_dyncast_263.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
			_hx_up := alt
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), nil, nil)
		{
			var tmp_key_1 Hx_Typedef_haxe__int64___int64
			_ = tmp_key_1
			tmp_key_1 = alt.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(tmp_key_1))
			func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
				var _hx_closure_src_647 func(high int32, low int32, value any)
				_ = _hx_closure_src_647
				_hx_closure_src_647 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					_hx_closure_src_647(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), state)
		}
	}

	Hx_Field_loreline_objects_setField(this, state.Hx_Field_fields, "_visitCount", count)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isChoiceOptionChosen(option *Hx_Obj_loreline_nchoiceoption) bool {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = option.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_648 func(high int32, low int32) any
		_ = _hx_closure_src_648
		_hx_closure_src_648 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_264 any
			_ = hx_dyncast_264
			hx_dyncast_264 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_648(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_264 != nil {
				_hx_tmp_0 = hx_dyncast_264.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		return false
	}

	var chosen any
	_ = chosen
	chosen = Hx_Field_loreline_objects_getField(this, state.Hx_Field_fields, "_chosen")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(chosen, nil))) {
		return false
	}

	return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(chosen)))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_setChoiceOptionChosen(option *Hx_Obj_loreline_nchoiceoption) {
	var key Hx_Typedef_haxe__int64___int64
	_ = key
	key = option.Hx_Field_id
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(key))
	var state *Hx_Obj_loreline_runtimestate
	_ = state
	state = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
		var _hx_closure_src_649 func(high int32, low int32) any
		_ = _hx_closure_src_649
		_hx_closure_src_649 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
		return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
			var hx_dyncast_265 any
			_ = hx_dyncast_265
			hx_dyncast_265 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_649(high, low), "Hx_Obj_loreline_runtimestate")
			var _hx_tmp_0 *Hx_Obj_loreline_runtimestate
			_ = _hx_tmp_0
			if hx_dyncast_265 != nil {
				_hx_tmp_0 = hx_dyncast_265.(interface {
					Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
				}).Hx_As_Hx_Obj_loreline_runtimestate()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()(((int32)((_hx_tmp_1 >> int64(((int32)(32)))))), ((int32)(key)))
	if state == nil {
		state = Hx_Obj_loreline_runtimestate_CreateInstance(this, (func() *Hx_Obj_loreline_node {
			_hx_up := option
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})(), nil, nil)
		{
			var tmp_key_1 Hx_Typedef_haxe__int64___int64
			_ = tmp_key_1
			tmp_key_1 = option.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(tmp_key_1))
			func() func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
				var _hx_closure_src_650 func(high int32, low int32, value any)
				_ = _hx_closure_src_650
				_hx_closure_src_650 = this.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__set
				return func(high int32, low int32, value *Hx_Obj_loreline_runtimestate) {
					_hx_closure_src_650(high, low, ((any)(value)))
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(tmp_key_1)), state)
		}
	}

	Hx_Field_loreline_objects_setField(this, state.Hx_Field_fields, "_chosen", true)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalAssignment(assign *Hx_Obj_loreline_nassign, next func()) {
	var target Hx_Enum_loreline_runtimeaccess
	_ = target
	target = this.VTable.Hx_Field_resolveAssignmentTarget(assign.Hx_Field_target)
	var value any
	_ = value
	value = this.VTable.Hx_Field_evaluateExpression(assign.Hx_Field_value)
	var currentValue any
	_ = currentValue
	switch assign.Hx_Field_op.Hx_Field_enumIndex() {
	case 22:
		{
			var _hx_tmp_0 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Enum_loreline_tokentype_OpPlus{}

			currentValue = this.VTable.Hx_Field_performOperation(_hx_tmp_0, this.VTable.Hx_Field_readAccess(target), value, assign.Hx_Field_pos)
		}

	case 23:
		{
			var _hx_tmp_1 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Enum_loreline_tokentype_OpMinus{}

			currentValue = this.VTable.Hx_Field_performOperation(_hx_tmp_1, this.VTable.Hx_Field_readAccess(target), value, assign.Hx_Field_pos)
		}

	case 24:
		{
			var _hx_tmp_2 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_2
			_hx_tmp_2 = Hx_Enum_loreline_tokentype_OpMultiply{}

			currentValue = this.VTable.Hx_Field_performOperation(_hx_tmp_2, this.VTable.Hx_Field_readAccess(target), value, assign.Hx_Field_pos)
		}

	case 25:
		{
			var _hx_tmp_3 Hx_Enum_loreline_tokentype
			_ = _hx_tmp_3
			_hx_tmp_3 = Hx_Enum_loreline_tokentype_OpDivide{}

			currentValue = this.VTable.Hx_Field_performOperation(_hx_tmp_3, this.VTable.Hx_Field_readAccess(target), value, assign.Hx_Field_pos)
		}

	case 21, 26:
		{
			currentValue = value
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid assignment operator", assign.Hx_Field_pos)))))
		}
	}

	this.VTable.Hx_Field_writeAccess(target, currentValue)
	next()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_snapshotOriginalValue(value any) any {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		_hx_tmp_3 = true
	} else {
		_hx_tmp_3 = Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI)
	}

	if _hx_tmp_3 {
		_hx_tmp_2 = true
	} else {
		_hx_tmp_2 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI)
	}

	if _hx_tmp_0 {
		return value
	}

	var _hx_tmp_4 bool
	_ = _hx_tmp_4
	if Hx_Field_loreline_runtimecharacterref_characterOf(value) != nil {
		_hx_tmp_4 = true
	} else {
		_hx_tmp_4 = (Hx_Field_loreline_runtimebeatref_beatOf(value) != nil)
	}

	if _hx_tmp_4 {
		return value
	}

	if Hx_Field_loreline_arrays_isArray(value) {
		var _hx_reserved_copy any
		_ = _hx_reserved_copy
		_hx_reserved_copy = Hx_Field_loreline_arrays_createArray()
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = Hx_Field_loreline_arrays_arrayLength(value)
			for _g < _g1 {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_5
				var _hx_tmp_6 any
				_ = _hx_tmp_6
				_hx_tmp_6 = _hx_reserved_copy
				Hx_Field_loreline_arrays_arrayPush(_hx_tmp_6, this.VTable.Hx_Field_snapshotOriginalValue(Hx_Field_loreline_arrays_arrayGet(value, i)))
			}
		}

		return _hx_reserved_copy
	}

	if Hx_Field_loreline_objects_isFields(value) {
		var _hx_tmp_7 any
		_ = _hx_tmp_7
		_hx_tmp_7 = nil
		var _hx_tmp_8 any
		_ = _hx_tmp_8
		_hx_tmp_8 = nil
		var _hx_reserved_copy any
		_ = _hx_reserved_copy
		_hx_reserved_copy = Hx_Field_loreline_objects_createFields(_hx_tmp_7, _hx_tmp_8, nil)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[string]
			_ = _g1
			_g1 = Hx_Field_loreline_objects_getFields(this, value)
			for {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g
				if !(_hx_tmp_9 < _g1.Len()) {
					break
				}

				var key string
				_ = key
				key = _g1.Get(_g)
				_g++
				var _hx_tmp_10 *Hx_Obj_loreline_interpreter
				_ = _hx_tmp_10
				_hx_tmp_10 = this
				var _hx_tmp_11 any
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_reserved_copy
				var _hx_tmp_12 string
				_ = _hx_tmp_12
				_hx_tmp_12 = key
				Hx_Field_loreline_objects_setField(_hx_tmp_10, _hx_tmp_11, _hx_tmp_12, this.VTable.Hx_Field_snapshotOriginalValue(Hx_Field_loreline_objects_getField(this, value, key)))
			}
		}

		return _hx_reserved_copy
	}

	return value
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isOriginalScriptExpression(expr *Hx_Obj_loreline_nexpr) bool {
	switch func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
		var _hx_closure_src_651 func(o any) *Hx_Obj_go_haxe_hxclass
		_ = _hx_closure_src_651
		_hx_closure_src_651 = Hx_Field_type_getClass
		return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
			return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_651(((any)(o)))))
		}
	}()(expr) {
	case Hx_Obj_loreline_nbinary_RTTI:
		{
			var bin *Hx_Obj_loreline_nbinary
			_ = bin
			bin = expr.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
			}).Hx_As_Hx_Obj_loreline_nbinary()
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if this.VTable.Hx_Field_isOriginalScriptExpression(bin.Hx_Field_left) {
				_hx_tmp_0 = this.VTable.Hx_Field_isOriginalScriptExpression(bin.Hx_Field_right)
			} else {
				_hx_tmp_0 = false
			}

			return _hx_tmp_0
		}

	case Hx_Obj_loreline_naccess_RTTI, Hx_Obj_loreline_narrayaccess_RTTI, Hx_Obj_loreline_ncall_RTTI:
		{
			return false
		}

	case Hx_Obj_loreline_nliteral_RTTI:
		{
			var lit *Hx_Obj_loreline_nliteral
			_ = lit
			lit = expr.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
			}).Hx_As_Hx_Obj_loreline_nliteral()
			switch lit.Hx_Field_literalType.Hx_Field_enumIndex() {
			case 0, 1, 2:
				{
					return true
				}

			case 3:
				{
					var arr HxArray[any]
					_ = arr
					arr = HxAnyToArray(lit.Hx_Field_value)
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_1 int32
							_ = _hx_tmp_1
							_hx_tmp_1 = _g
							if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(arr)) {
								break
							}

							var elem any
							_ = elem
							elem = arr.Get(_g)
							_g++
							var _hx_tmp_2 bool
							_ = _hx_tmp_2
							if Hx_Field_std_isOfType(elem, Hx_Obj_loreline_nexpr_RTTI) {
								var hx_dyncast_266 any
								_ = hx_dyncast_266
								hx_dyncast_266 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_nexpr")
								var _hx_tmp_3 *Hx_Obj_loreline_nexpr
								_ = _hx_tmp_3
								if hx_dyncast_266 != nil {
									_hx_tmp_3 = hx_dyncast_266.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
								} else {
									_hx_tmp_3 = nil
								}

								_hx_tmp_2 = !this.VTable.Hx_Field_isOriginalScriptExpression(_hx_tmp_3)
							} else {
								_hx_tmp_2 = false
							}

							if _hx_tmp_2 {
								return false
							}
						}
					}

					return true
				}

			case 4:
				{
					var fields HxArray[*Hx_Obj_loreline_nobjectfield]
					_ = fields
					fields = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](lit.Hx_Field_value)
					{
						var _g int32
						_ = _g
						_g = 0
						for {
							var _hx_tmp_4 int32
							_ = _hx_tmp_4
							_hx_tmp_4 = _g
							if !(_hx_tmp_4 < fields.Len()) {
								break
							}

							var field *Hx_Obj_loreline_nobjectfield
							_ = field
							field = fields.Get(_g)
							_g++
							if !this.VTable.Hx_Field_isOriginalScriptExpression(field.Hx_Field_value) {
								return false
							}
						}
					}

					return true
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown("unreachable"))
				}
			}
		}

	case Hx_Obj_loreline_nstringliteral_RTTI:
		{
			var str *Hx_Obj_loreline_nstringliteral
			_ = str
			str = expr.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
			}).Hx_As_Hx_Obj_loreline_nstringliteral()
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_nstringpart]
				_ = _g1
				_g1 = str.Hx_Field_parts
				for {
					var _hx_tmp_5 int32
					_ = _hx_tmp_5
					_hx_tmp_5 = _g
					if !(_hx_tmp_5 < _g1.Len()) {
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
								continue
							}

						case 1:
							{
								return false
							}

						case 2:
							{
								var content *Hx_Obj_loreline_nstringliteral
								_ = content
								content = _g2.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_expr
								if !this.VTable.Hx_Field_isOriginalScriptExpression((func() *Hx_Obj_loreline_nexpr {
									_hx_up := content
									if _hx_up != nil {
										return &_hx_up.Hx_Obj_loreline_nexpr
									}
									return nil
								})()) {
									return false
								}
							}

						}
					}
				}
			}

			return true
		}

	case Hx_Obj_loreline_nunary_RTTI:
		{
			var un *Hx_Obj_loreline_nunary
			_ = un
			un = expr.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
			}).Hx_As_Hx_Obj_loreline_nunary()
			return this.VTable.Hx_Field_isOriginalScriptExpression(un.Hx_Field_operand)
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}

	return false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isBeatCall(node *Hx_Obj_loreline_astnode, scopeLevel__isOptional any) bool {
	scopeLevel := HxDefault[int32]()
	_ = scopeLevel

	if scopeLevel__isOptional != nil {
		scopeLevel = HxConvert[int32](scopeLevel__isOptional)
	} else {
		scopeLevel = -1
	}
	{
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_ncall_RTTI) {
			return (this.VTable.Hx_Field_resolveBeatRefFromCall(node.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall()) != nil)
		}

		return false
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_isDynamicCallForm(call *Hx_Obj_loreline_ncall) bool {
	if Hx_Field_std_isOfType(call.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
		var access *Hx_Obj_loreline_naccess
		_ = access
		access = call.Hx_Field_target.VTable.(interface {
			Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
		}).Hx_As_Hx_Obj_loreline_naccess()
		if (access.Hx_Field_target == nil) && (access.Hx_Field_name == "call") {
			return !this.VTable.Hx_Field_resolvesExistingBinding("call")
		}
	}

	return false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveBeatByName(name string, scopeLevel__isOptional any) *Hx_Obj_loreline_nbeatdecl {
	scopeLevel := HxDefault[int32]()
	_ = scopeLevel

	if scopeLevel__isOptional != nil {
		scopeLevel = HxConvert[int32](scopeLevel__isOptional)
	} else {
		scopeLevel = -1
	}
	{
		var result *Hx_Obj_loreline_nbeatdecl
		_ = result
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = scopeLevel
		this.VTable.Hx_Field_eachResolutionScope(_hx_tmp_0, func(scope *Hx_Obj_loreline_runtimescope) bool {
			var b *Hx_Obj_loreline_nbeatdecl
			_ = b
			b = scope.VTable.Hx_Field_beatByName(name)
			if b != nil {
				result = b
				return true
			}

			return false
		})
		if result != nil {
			return result
		}

		return func() func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
			var _hx_closure_src_652 func(_hx_param_key any) any
			_ = _hx_closure_src_652
			_hx_closure_src_652 = this.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
			return func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
				var hx_dyncast_267 any
				_ = hx_dyncast_267
				hx_dyncast_267 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_652(((any)(_hx_param_key))), "Hx_Obj_loreline_nbeatdecl")
				var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_1
				if hx_dyncast_267 != nil {
					_hx_tmp_1 = hx_dyncast_267.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
				} else {
					_hx_tmp_1 = nil
				}

				return _hx_tmp_1
			}
		}()(name)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveBeatRefByName(name string) *Hx_Obj_loreline_runtimebeatref {
	var result *Hx_Obj_loreline_runtimebeatref
	_ = result
	var walk func(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32)
	_ = walk
	walk = func(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32) {
		var i int32
		_ = i
		i = startIndex
		for i >= 0 {
			var scope *Hx_Obj_loreline_runtimescope
			_ = scope
			scope = scopes.Get(i)
			var b *Hx_Obj_loreline_nbeatdecl
			_ = b
			b = scope.VTable.Hx_Field_beatByName(name)
			if b != nil {
				var _hx_tmp_0 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _hx_tmp_0
				_hx_tmp_0 = scopes
				result = Hx_Obj_loreline_runtimebeatref_CreateInstance(b, (Hx_Array_Slice[*Hx_Obj_loreline_runtimescope](_hx_tmp_0, 0, HxNullable[int32]{Value: (i + 1), Valid: true})))
				return
			}

			if scope.Hx_Field_captured != nil {
				var _hx_tmp_1 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _hx_tmp_1
				_hx_tmp_1 = scope.Hx_Field_captured
				walk(_hx_tmp_1, (scope.Hx_Field_captured.Len() - 1))
				return
			}

			i--
		}
	}

	var _hx_tmp_2 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_stack
	walk(_hx_tmp_2, (this.Hx_Field_stack.Len() - 1))
	if result != nil {
		return result
	}

	var topLevel *Hx_Obj_loreline_nbeatdecl
	_ = topLevel
	topLevel = func() func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_653 func(_hx_param_key any) any
		_ = _hx_closure_src_653
		_hx_closure_src_653 = this.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_268 any
			_ = hx_dyncast_268
			hx_dyncast_268 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_653(((any)(_hx_param_key))), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_3 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_3
			if hx_dyncast_268 != nil {
				_hx_tmp_3 = hx_dyncast_268.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_3 = nil
			}

			return _hx_tmp_3
		}
	}()(name)
	if topLevel != nil {
		return Hx_Obj_loreline_runtimebeatref_CreateInstance(topLevel, HxMakeArray[*Hx_Obj_loreline_runtimescope]())
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveBeatRefFromNode(name string, node *Hx_Obj_loreline_node) *Hx_Obj_loreline_runtimebeatref {
	if node == nil {
		return this.VTable.Hx_Field_resolveBeatRefByName(name)
	}

	var lexical *Hx_Obj_loreline_nbeatdecl
	_ = lexical
	lexical = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_findBeatByNameFromNode(name, node)
	if lexical == nil {
		return nil
	}

	return Hx_Obj_loreline_runtimebeatref_CreateInstance(lexical, this.VTable.Hx_Field_captureChainForBeat(lexical))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_captureChainForBeat(beat *Hx_Obj_loreline_nbeatdecl) HxArray[*Hx_Obj_loreline_runtimescope] {
	var parentBeat *Hx_Obj_loreline_nbeatdecl
	_ = parentBeat
	parentBeat = func() func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
		var _hx_closure_src_654 func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) any
		_ = _hx_closure_src_654
		_hx_closure_src_654 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getFirstParentOfType
		return func(node *Hx_Obj_loreline_node, _hx_reserved_type *Hx_Obj_go_haxe_hxclass) *Hx_Obj_loreline_nbeatdecl {
			var hx_dyncast_269 any
			_ = hx_dyncast_269
			hx_dyncast_269 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_654(node, _hx_reserved_type), "Hx_Obj_loreline_nbeatdecl")
			var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
			_ = _hx_tmp_0
			if hx_dyncast_269 != nil {
				_hx_tmp_0 = hx_dyncast_269.(interface {
					Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
				}).Hx_As_Hx_Obj_loreline_nbeatdecl()
			} else {
				_hx_tmp_0 = nil
			}

			return _hx_tmp_0
		}
	}()((func() *Hx_Obj_loreline_node {
		_hx_up := beat
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})(), ((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)))
	if parentBeat == nil {
		return HxMakeArray[*Hx_Obj_loreline_runtimescope]()
	}

	var result HxArray[*Hx_Obj_loreline_runtimescope]
	_ = result
	var walk func(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32)
	_ = walk
	walk = func(scopes HxArray[*Hx_Obj_loreline_runtimescope], startIndex int32) {
		var i int32
		_ = i
		i = startIndex
		var parentMatch int32
		_ = parentMatch
		parentMatch = -1
		for i >= 0 {
			var scope *Hx_Obj_loreline_runtimescope
			_ = scope
			scope = scopes.Get(i)
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if scope.Hx_Field_beats != nil {
				var _hx_tmp_2 HxArray[*Hx_Obj_loreline_nbeatdecl]
				_ = _hx_tmp_2
				_hx_tmp_2 = scope.Hx_Field_beats
				var _hx_tmp_3 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_3
				_hx_tmp_3 = beat
				var _hx_tmp_4 HxArray[*Hx_Obj_loreline_nbeatdecl]
				_ = _hx_tmp_4
				_hx_tmp_4 = _hx_tmp_2
				var _hx_tmp_5 *Hx_Obj_loreline_nbeatdecl
				_ = _hx_tmp_5
				_hx_tmp_5 = _hx_tmp_3
				_hx_tmp_1 = ((Hx_Array_IndexOf[*Hx_Obj_loreline_nbeatdecl](_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{})) != -1)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				var _hx_tmp_6 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _hx_tmp_6
				_hx_tmp_6 = scopes
				result = (Hx_Array_Slice[*Hx_Obj_loreline_runtimescope](_hx_tmp_6, 0, HxNullable[int32]{Value: (i + 1), Valid: true}))
				return
			}

			if (parentMatch == -1) && (scope.Hx_Field_beat == parentBeat) {
				parentMatch = i
			}

			if scope.Hx_Field_captured != nil {
				var _hx_tmp_7 HxArray[*Hx_Obj_loreline_runtimescope]
				_ = _hx_tmp_7
				_hx_tmp_7 = scope.Hx_Field_captured
				walk(_hx_tmp_7, (scope.Hx_Field_captured.Len() - 1))
				if (result == nil) && (parentMatch != -1) {
					var _hx_tmp_8 HxArray[*Hx_Obj_loreline_runtimescope]
					_ = _hx_tmp_8
					_hx_tmp_8 = scopes
					result = (Hx_Array_Slice[*Hx_Obj_loreline_runtimescope](_hx_tmp_8, 0, HxNullable[int32]{Value: (parentMatch + 1), Valid: true}))
				}

				return
			}

			i--
		}

		if (result == nil) && (parentMatch != -1) {
			var _hx_tmp_9 HxArray[*Hx_Obj_loreline_runtimescope]
			_ = _hx_tmp_9
			_hx_tmp_9 = scopes
			result = (Hx_Array_Slice[*Hx_Obj_loreline_runtimescope](_hx_tmp_9, 0, HxNullable[int32]{Value: (parentMatch + 1), Valid: true}))
		}
	}

	var _hx_tmp_10 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_10
	_hx_tmp_10 = this.Hx_Field_stack
	walk(_hx_tmp_10, (this.Hx_Field_stack.Len() - 1))
	if result != nil {
		return result
	} else {
		return HxMakeArray[*Hx_Obj_loreline_runtimescope]()
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveDynamicBeatTarget(value any, node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_runtimebeatref {
	if Hx_Field_std_isOfType(value, Hx_Obj_loreline_runtimebeatref_RTTI) {
		var hx_dyncast_270 any
		_ = hx_dyncast_270
		hx_dyncast_270 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_runtimebeatref")
		var _hx_tmp_0 *Hx_Obj_loreline_runtimebeatref
		_ = _hx_tmp_0
		if hx_dyncast_270 != nil {
			_hx_tmp_0 = hx_dyncast_270.(interface {
				Hx_As_Hx_Obj_loreline_runtimebeatref() *Hx_Obj_loreline_runtimebeatref
			}).Hx_As_Hx_Obj_loreline_runtimebeatref()
		} else {
			_hx_tmp_0 = nil
		}

		return _hx_tmp_0
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_loreline_nbeatdecl_RTTI) {
		var hx_dyncast_271 any
		_ = hx_dyncast_271
		hx_dyncast_271 = Hx_Field_go_haxe_hxdynamic_castClass(value, "Hx_Obj_loreline_nbeatdecl")
		var _hx_tmp_1 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_1
		if hx_dyncast_271 != nil {
			_hx_tmp_1 = hx_dyncast_271.(interface {
				Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
			}).Hx_As_Hx_Obj_loreline_nbeatdecl()
		} else {
			_hx_tmp_1 = nil
		}

		return Hx_Obj_loreline_runtimebeatref_CreateInstance(_hx_tmp_1, HxMakeArray[*Hx_Obj_loreline_runtimescope]())
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		var name string
		_ = name
		name = Hx_Field_go_haxe_hxdynamic_toString(value)
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = name
		var resolved *Hx_Obj_loreline_runtimebeatref
		_ = resolved
		resolved = this.VTable.Hx_Field_resolveBeatRefFromNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
			_hx_up := node
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		if resolved == nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(this.VTable.Hx_Field_beatNotFoundError(name, node)))))
		}

		return resolved
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid dynamic beat target: ", this.VTable.Hx_Field_valueToString(value)), node.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_beatNotFoundError(name string, node *Hx_Obj_loreline_astnode) *Hx_Obj_loreline_runtimeerror {
	var _hx_tmp_0 *Hx_Obj_loreline_position
	_ = _hx_tmp_0
	if node != nil {
		_hx_tmp_0 = node.Hx_Field_pos
	} else {
		_hx_tmp_0 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
	}

	var pos *Hx_Obj_loreline_position
	_ = pos
	pos = _hx_tmp_0
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nbeatdecl]
		_ = _g1
		_g1 = func() func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nbeatdecl] {
			var _hx_closure_src_655 func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[any]
			_ = _hx_closure_src_655
			_hx_closure_src_655 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodesOfType
			return func(nodeType *Hx_Obj_go_haxe_hxclass, includeImported__isOptional any) HxArray[*Hx_Obj_loreline_nbeatdecl] {
				includeImported := HxDefault[bool]()
				_ = includeImported

				if includeImported__isOptional != nil {
					includeImported = HxConvert[bool](includeImported__isOptional)
				} else {
				}
				{
					var _hx_tmp_1 *Hx_Obj_go_haxe_hxclass
					_ = _hx_tmp_1
					_hx_tmp_1 = nodeType
					var _hx_tmp_2 any
					_ = _hx_tmp_2
					if includeImported__isOptional == nil {
						_hx_tmp_2 = nil
					} else {
						_hx_tmp_2 = includeImported
					}

					return HxMakeArrayView[*Hx_Obj_loreline_nbeatdecl](_hx_closure_src_655(_hx_tmp_1, _hx_tmp_2))
				}
			}
		}()(((*Hx_Obj_go_haxe_hxclass)(Hx_Obj_loreline_nbeatdecl_RTTI)), true)
		for {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			if !(_hx_tmp_3 < _g1.Len()) {
				break
			}

			var decl *Hx_Obj_loreline_nbeatdecl
			_ = decl
			decl = _g1.Get(_g)
			_g++
			if decl.Hx_Field_name == name {
				return Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Beat ", name), " is not visible from this scope. Store a reference (ref = "), name), ") while it is in scope to use it from elsewhere."), pos)
			}
		}
	}

	return Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Beat ", name), " not found"), pos)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveBeatRefFromCall(call *Hx_Obj_loreline_ncall) *Hx_Obj_loreline_runtimebeatref {
	if Hx_Field_std_isOfType(call.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
		var access *Hx_Obj_loreline_naccess
		_ = access
		access = call.Hx_Field_target.VTable.(interface {
			Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
		}).Hx_As_Hx_Obj_loreline_naccess()
		if access.Hx_Field_target == nil {
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = access.Hx_Field_name
			return this.VTable.Hx_Field_resolveBeatRefFromNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
				_hx_up := access
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalCall(call *Hx_Obj_loreline_ncall, next func()) {
	var resolvedRef *Hx_Obj_loreline_runtimebeatref
	_ = resolvedRef
	resolvedRef = this.VTable.Hx_Field_resolveBeatRefFromCall(call)
	if resolvedRef != nil {
		var argValues HxArray[any]
		_ = argValues
		argValues = this.VTable.Hx_Field_evaluateBeatArgs(resolvedRef.Hx_Field_beat, call.Hx_Field_args, call.Hx_Field_pos)
		var _hx_tmp_0 *Hx_Obj_loreline_nbeatdecl
		_ = _hx_tmp_0
		_hx_tmp_0 = resolvedRef.Hx_Field_beat
		var _hx_tmp_1 func()
		_ = _hx_tmp_1
		_hx_tmp_1 = next
		var _hx_tmp_2 HxArray[any]
		_ = _hx_tmp_2
		_hx_tmp_2 = argValues
		var _hx_tmp_3 HxArray[*Hx_Obj_loreline_runtimescope]
		_ = _hx_tmp_3
		if resolvedRef.Hx_Field_stack.Len() > 0 {
			_hx_tmp_3 = resolvedRef.Hx_Field_stack
		} else {
			_hx_tmp_3 = nil
		}

		this.VTable.Hx_Field_evalBeatRun(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3)
		return
	}

	this.VTable.Hx_Field_evaluateFunctionCall(call, next)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evalTransition(transition *Hx_Obj_loreline_ntransition) {
	var resolvedRef *Hx_Obj_loreline_runtimebeatref
	_ = resolvedRef
	if transition.Hx_Field_targetExpr != nil {
		var value any
		_ = value
		value = this.VTable.Hx_Field_evaluateExpression(transition.Hx_Field_targetExpr)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
			_hx_tmp_0 = (Hx_Field_go_haxe_hxdynamic_toString(value) == ".")
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if transition.Hx_Field_args != nil {
				_hx_tmp_1 = (transition.Hx_Field_args.Len() > 0)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot pass arguments when ending the script", transition.Hx_Field_pos)))))
			}

			this.VTable.Hx_Field_finish()
			return
		}

		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = value
		resolvedRef = this.VTable.Hx_Field_resolveDynamicBeatTarget(_hx_tmp_2, (func() *Hx_Obj_loreline_astnode {
			_hx_up := transition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_astnode
			}
			return nil
		})())
	} else {
		var beatName string
		_ = beatName
		beatName = transition.Hx_Field_target
		if beatName == "." {
			this.VTable.Hx_Field_finish()
			return
		}

		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = beatName
		resolvedRef = this.VTable.Hx_Field_resolveBeatRefFromNode(_hx_tmp_3, (func() *Hx_Obj_loreline_node {
			_hx_up := transition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		if resolvedRef == nil {
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = beatName
			panic(Hx_Field_haxe_exception_thrown(((any)(this.VTable.Hx_Field_beatNotFoundError(_hx_tmp_4, (func() *Hx_Obj_loreline_astnode {
				_hx_up := transition
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_astnode
				}
				return nil
			})())))))
		}
	}

	var argValues HxArray[any]
	_ = argValues
	argValues = this.VTable.Hx_Field_evaluateBeatArgs(resolvedRef.Hx_Field_beat, transition.Hx_Field_args, transition.Hx_Field_pos)
	var _hx_tmp_5 *Hx_Obj_loreline_nbeatdecl
	_ = _hx_tmp_5
	_hx_tmp_5 = resolvedRef.Hx_Field_beat
	var _hx_tmp_6 HxArray[any]
	_ = _hx_tmp_6
	_hx_tmp_6 = argValues
	var _hx_tmp_7 HxArray[*Hx_Obj_loreline_runtimescope]
	_ = _hx_tmp_7
	if resolvedRef.Hx_Field_stack.Len() > 0 {
		_hx_tmp_7 = resolvedRef.Hx_Field_stack
	} else {
		_hx_tmp_7 = nil
	}

	this.VTable.Hx_Field_transitionToBeat(_hx_tmp_5, _hx_tmp_6, _hx_tmp_7)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getTranslatedString(node *Hx_Obj_loreline_astnode, str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
	if this.Hx_Field_context.Hx_Field_translations != nil {
		var id HxNullable[string]
		_ = id
		id = this.VTable.Hx_Field_findHashCommentId(node, str)
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if (id).Valid != false {
			_hx_tmp_0 = (this.Hx_Field_context.Hx_Field_script.Hx_Field_filePath != HxStringNull)
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[string]
			_ = _g1
			_g1 = this.Hx_Field_context.Hx_Field_lens.VTable.Hx_Field_getNodeAncestorFilePaths((func() *Hx_Obj_loreline_node {
				_hx_up := node
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < _g1.Len()) {
					break
				}

				var relPath string
				_ = relPath
				relPath = _g1.Get(_g)
				_g++
				var _hx_tmp_3 string
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_stringConcat(relPath, "#")
				var scoped *Hx_Obj_loreline_nstringliteral
				_ = scoped
				scoped = func() func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
					var _hx_closure_src_656 func(_hx_param_key any) any
					_ = _hx_closure_src_656
					_hx_closure_src_656 = this.Hx_Field_context.Hx_Field_translations.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
						var hx_dyncast_272 any
						_ = hx_dyncast_272
						hx_dyncast_272 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_656(((any)(_hx_param_key))), "Hx_Obj_loreline_nstringliteral")
						var _hx_tmp_2 *Hx_Obj_loreline_nstringliteral
						_ = _hx_tmp_2
						if hx_dyncast_272 != nil {
							_hx_tmp_2 = hx_dyncast_272.(interface {
								Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
							}).Hx_As_Hx_Obj_loreline_nstringliteral()
						} else {
							_hx_tmp_2 = nil
						}

						return _hx_tmp_2
					}
				}()(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_3, (id).Value))
				if scoped != nil {
					return scoped
				}
			}
		}
	}

	return str
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_findHashCommentId(node *Hx_Obj_loreline_astnode, str *Hx_Obj_loreline_nstringliteral) HxNullable[string] {
	return Hx_Field_loreline_astutils_findHashComment(node, str)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateString(str *Hx_Obj_loreline_nstringliteral) any {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_context.Hx_Field_stringLiteralProcessors.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			str = this.Hx_Field_context.Hx_Field_stringLiteralProcessors.Get(i)(str)
		}
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var tags HxArray[*Hx_Obj_loreline_texttag]
	_ = tags
	tags = HxMakeArray[*Hx_Obj_loreline_texttag]()
	var offset int32
	_ = offset
	offset = 0
	var numParts int32
	_ = numParts
	numParts = str.Hx_Field_parts.Len()
	var keepWhitespace bool
	_ = keepWhitespace
	keepWhitespace = (str.Hx_Field_quotes != ((int32)(0)))
	var keepIndents bool
	_ = keepIndents
	keepIndents = (str.Hx_Field_quotes != ((int32)(0)))
	var keepComments bool
	_ = keepComments
	keepComments = (str.Hx_Field_quotes != ((int32)(0)))
	var trailingTextPartIndex int32
	_ = trailingTextPartIndex
	trailingTextPartIndex = -1
	if !keepWhitespace {
		var n int32
		_ = n
		n = (numParts - 1)
	hx_label_0:
		for n >= 0 {
			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = str.Hx_Field_parts.Get(n)
			switch part.Hx_Field_partType.Hx_Field_enumIndex() {
			case 0:
				{
					trailingTextPartIndex = n
					break hx_label_0
				}

			case 1:
				{
					break hx_label_0
				}

			case 2:
				{
					n--
				}

			}
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = numParts
		for _g2 < _g3 {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_1
			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = str.Hx_Field_parts.Get(i)
			{
				var _g Hx_Enum_loreline_stringparttype
				_ = _g
				_g = part.Hx_Field_partType
				switch _g.Hx_Field_enumIndex() {
				case 0:
					{
						var text string
						_ = text
						text = _g.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						{
							if !keepWhitespace {
								var l int32
								_ = l
								l = int32(utf8.RuneCountInString(text))
								var r int32
								_ = r
								r = 0
								for {
									var _hx_tmp_2 bool
									_ = _hx_tmp_2
									if r < l {
										_hx_tmp_2 = Hx_Field_stringtools_isSpace(text, r)
									} else {
										_hx_tmp_2 = false
									}

									if !(_hx_tmp_2) {
										break
									}

									r++
								}

								if r > 0 {
									var _hx_tmp_3 string
									_ = _hx_tmp_3
									_hx_tmp_3 = text
									var _hx_tmp_4 int32
									_ = _hx_tmp_4
									_hx_tmp_4 = r
									text = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_3, _hx_tmp_4, HxNullable[int32]{Value: (l - r), Valid: true}))
								}
							}

							if i == trailingTextPartIndex {
								var l int32
								_ = l
								l = int32(utf8.RuneCountInString(text))
								var r int32
								_ = r
								r = 0
								for {
									var _hx_tmp_5 bool
									_ = _hx_tmp_5
									if r < l {
										_hx_tmp_5 = Hx_Field_stringtools_isSpace(text, ((l - r) - 1))
									} else {
										_hx_tmp_5 = false
									}

									if !(_hx_tmp_5) {
										break
									}

									r++
								}

								if r > 0 {
									var _hx_tmp_6 string
									_ = _hx_tmp_6
									_hx_tmp_6 = text
									text = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_6, 0, HxNullable[int32]{Value: (l - r), Valid: true}))
								}
							}

							if !keepComments {
								text = this.VTable.Hx_Field_stripStringComments(text)
							}

							if !keepIndents {
								text = this.VTable.Hx_Field_stripStringIndent(text)
							}

							var chars HxArray[int32]
							_ = chars
							chars = Hx_Field_loreline_utf8_toCodes(text)
							var _hx_reserved_len int32
							_ = _hx_reserved_len
							_hx_reserved_len = HxMakeArrayView[any](chars).Len()
							if _hx_reserved_len > 0 {
								keepWhitespace = true
							}

							var prevIsDollar bool
							_ = prevIsDollar
							prevIsDollar = false
							var prevIsHash bool
							_ = prevIsHash
							prevIsHash = false
							var escaped bool
							_ = escaped
							escaped = false
							{
								var _g1 int32
								_ = _g1
								_g1 = 0
								var tmp__g2_1 int32
								_ = tmp__g2_1
								tmp__g2_1 = _hx_reserved_len
								for _g1 < tmp__g2_1 {
									var _hx_tmp_7 int32
									_ = _hx_tmp_7
									_hx_tmp_7 = _g1
									_g1 = (_g1 + 1)
									var tmp_i_1 int32
									_ = tmp_i_1
									tmp_i_1 = _hx_tmp_7
									var _hx_tmp_8 int32
									_ = _hx_tmp_8
									var _hx_tmp_9 bool
									_ = _hx_tmp_9
									if tmp_i_1 < 0 {
										_hx_tmp_9 = true
									} else {
										var _hx_tmp_10 int32
										_ = _hx_tmp_10
										_hx_tmp_10 = tmp_i_1
										var _hx_tmp_11 int32
										_ = _hx_tmp_11
										_hx_tmp_11 = _hx_tmp_10
										_hx_tmp_9 = (_hx_tmp_11 >= HxMakeArrayView[any](chars).Len())
									}

									if _hx_tmp_9 {
										_hx_tmp_8 = -1
									} else {
										_hx_tmp_8 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(tmp_i_1))))
									}

									var c int32
									_ = c
									c = _hx_tmp_8
									if escaped {
										if c == 110 {
											buf.VTable.Hx_Field_addChar(10)
										} else {
											if c == 114 {
												buf.VTable.Hx_Field_addChar(13)
											} else {
												if c == 116 {
													buf.VTable.Hx_Field_addChar(9)
												} else {
													buf.VTable.Hx_Field_addChar(c)
												}
											}
										}

										escaped = false
										prevIsDollar = false
										prevIsHash = false
									} else {
										if c == 92 {
											escaped = true
											prevIsDollar = false
											prevIsHash = false
										} else {
											if c == 36 {
												if prevIsDollar {
													buf.VTable.Hx_Field_addChar(c)
													prevIsDollar = false
												} else {
													prevIsDollar = true
												}

												prevIsHash = false
											} else {
												if c == 35 {
													if prevIsHash {
														buf.VTable.Hx_Field_addChar(c)
														prevIsHash = false
													} else {
														prevIsHash = true
													}

													prevIsDollar = false
												} else {
													if prevIsHash {
														buf.VTable.Hx_Field_addChar(35)
														prevIsHash = false
													}

													buf.VTable.Hx_Field_addChar(c)
												}
											}
										}
									}
								}
							}

							if prevIsHash {
								buf.VTable.Hx_Field_addChar(35)
							}

							offset = (offset + _hx_reserved_len)
						}
					}

				case 1:
					{
						var expr *Hx_Obj_loreline_nexpr
						_ = expr
						expr = _g.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
						{
							keepWhitespace = true
							var value any
							_ = value
							value = this.VTable.Hx_Field_evaluateExpression(expr)
							var text string
							_ = text
							text = this.VTable.Hx_Field_valueToString(value)
							var _hx_tmp_12 int32
							_ = _hx_tmp_12
							_hx_tmp_12 = offset
							offset = (_hx_tmp_12 + int32(utf8.RuneCountInString(text)))
							func() func(x string) {
								var _hx_closure_src_657 func(x any)
								_ = _hx_closure_src_657
								_hx_closure_src_657 = buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_657(((any)(x)))
								}
							}()(text)
						}
					}

				case 2:
					{
						var closing bool
						_ = closing
						closing = _g.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_closing
						var expr *Hx_Obj_loreline_nstringliteral
						_ = expr
						expr = _g.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_expr
						{
							var tagValue string
							_ = tagValue
							tagValue = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(this.VTable.Hx_Field_evaluateString(expr), "text"))
							var _hx_tmp_13 HxArray[*Hx_Obj_loreline_texttag]
							_ = _hx_tmp_13
							_hx_tmp_13 = tags
							_ = (Hx_Array_Push[*Hx_Obj_loreline_texttag](_hx_tmp_13, Hx_Obj_loreline_texttag_CreateInstance(closing, tagValue, offset)))
						}
					}

				}
			}
		}
	}

	return any(map[string]any{"text": ((any)(buf.VTable.Hx_Field_toString())), "tags": ((any)(tags))})
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_pluralPipeProcess(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral {
	var lastExprPartIndex int32
	_ = lastExprPartIndex
	lastExprPartIndex = -1
	var changed bool
	_ = changed
	changed = false
	var newParts HxArray[*Hx_Obj_loreline_nstringpart]
	_ = newParts
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
			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = str.Hx_Field_parts.Get(i)
			{
				var tmp__g1_1 Hx_Enum_loreline_stringparttype
				_ = tmp__g1_1
				tmp__g1_1 = part.Hx_Field_partType
				switch tmp__g1_1.Hx_Field_enumIndex() {
				case 0:
					{
						var text string
						_ = text
						text = tmp__g1_1.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
						var _hx_tmp_1 bool
						_ = _hx_tmp_1
						if lastExprPartIndex >= 0 {
							_hx_tmp_1 = this.VTable.Hx_Field_rawTextHasUnescapedPipe(text)
						} else {
							_hx_tmp_1 = false
						}

						if _hx_tmp_1 {
							var countExpr *Hx_Obj_loreline_nexpr
							_ = countExpr
							var tmp__g_1 Hx_Enum_loreline_stringparttype
							_ = tmp__g_1
							tmp__g_1 = str.Hx_Field_parts.Get(lastExprPartIndex).Hx_Field_partType
							if tmp__g_1.Hx_Field_enumIndex() == 1 {
								var expr *Hx_Obj_loreline_nexpr
								_ = expr
								expr = tmp__g_1.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
								countExpr = expr
							} else {
								countExpr = nil
							}

							if countExpr != nil {
								if !changed {
									var tmp__g_2 HxArray[*Hx_Obj_loreline_nstringpart]
									_ = tmp__g_2
									tmp__g_2 = HxMakeArray[*Hx_Obj_loreline_nstringpart]()
									{
										var tmp__g1_2 int32
										_ = tmp__g1_2
										tmp__g1_2 = 0
										var _g2 int32
										_ = _g2
										_g2 = i
										for tmp__g1_2 < _g2 {
											var _hx_tmp_2 int32
											_ = _hx_tmp_2
											_hx_tmp_2 = tmp__g1_2
											tmp__g1_2 = (tmp__g1_2 + 1)
											var j int32
											_ = j
											j = _hx_tmp_2
											_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](tmp__g_2, str.Hx_Field_parts.Get(j)))
										}
									}

									newParts = tmp__g_2
									changed = true
								}

								this.VTable.Hx_Field_transformPipeRaw(text, countExpr, part.Hx_Field_pos, newParts)
								continue
							}
						}
					}

				case 1:
					{
						lastExprPartIndex = i
					}

				case 2:
					{
					}

				}
			}

			if changed {
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](newParts, part))
			}
		}
	}

	if !changed {
		return str
	}

	return Hx_Obj_loreline_nstringliteral_CreateInstance(str.Hx_Field_id, str.Hx_Field_pos, str.Hx_Field_quotes, newParts, nil, nil)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_rawTextHasUnescapedPipe(text string) bool {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(text))
	var i int32
	_ = i
	i = 0
	for i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
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
		if c1 == 92 {
			i = (i + 2)
			continue
		}

		if c1 == 124 {
			return true
		}

		i++
	}

	return false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_transformPipeRaw(text string, countExpr *Hx_Obj_loreline_nexpr, pos *Hx_Obj_loreline_position, parts HxArray[*Hx_Obj_loreline_nstringpart]) {
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = int32(utf8.RuneCountInString(text))
	var i int32
	_ = i
	i = 0
	var segStart int32
	_ = segStart
	segStart = 0
	for i < _hx_reserved_len {
		var c HxNullable[int32]
		_ = c
		c = (Hx_Field_go_haxe_hxstring_charCodeAt(text, i))
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
		if (c1 == 92) && ((i + 1) < _hx_reserved_len) {
			i = (i + 2)
			continue
		}

		if c1 == 40 {
			var pipePos int32
			_ = pipePos
			pipePos = -1
			var closePos int32
			_ = closePos
			closePos = -1
			var j int32
			_ = j
			j = (i + 1)
			for j < _hx_reserved_len {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, j))
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				if (tmp_c_1).Valid == false {
					_hx_tmp_1 = -1
				} else {
					_hx_tmp_1 = (tmp_c_1).Value
				}

				var cj int32
				_ = cj
				cj = _hx_tmp_1
				if (cj == 92) && ((j + 1) < _hx_reserved_len) {
					j = (j + 2)
					continue
				}

				if (cj == 124) && (pipePos == -1) {
					pipePos = j
				} else {
					if cj == 41 {
						closePos = j
						break
					}
				}

				j++
			}

			if (((pipePos != -1) && (closePos != -1)) && (pipePos > (i + 1))) && (closePos > (pipePos + 1)) {
				var pos1 int32
				_ = pos1
				pos1 = (i + 1)
				var tmp__hx_reserved_len_1 HxNullable[int32]
				_ = tmp__hx_reserved_len_1
				tmp__hx_reserved_len_1 = HxNullable[int32]{Value: ((pipePos - i) - 1), Valid: true}

				var _hx_tmp_2 string
				_ = _hx_tmp_2
				if (tmp__hx_reserved_len_1).Valid != false {
					_hx_tmp_2 = (Hx_Field_go_haxe_hxstring_substr(text, pos1, tmp__hx_reserved_len_1))
				} else {
					var _hx_tmp_3 string
					_ = _hx_tmp_3
					_hx_tmp_3 = text
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = pos1
					var _hx_tmp_5 string
					_ = _hx_tmp_5
					_hx_tmp_5 = _hx_tmp_3
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = _hx_tmp_4
					_hx_tmp_2 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_5, _hx_tmp_6, HxNullable[int32]{}))
				}

				var singular string
				_ = singular
				singular = _hx_tmp_2
				var pos2 int32
				_ = pos2
				pos2 = (pipePos + 1)
				var len1 HxNullable[int32]
				_ = len1
				len1 = HxNullable[int32]{Value: ((closePos - pipePos) - 1), Valid: true}

				var _hx_tmp_7 string
				_ = _hx_tmp_7
				if (len1).Valid != false {
					_hx_tmp_7 = (Hx_Field_go_haxe_hxstring_substr(text, pos2, len1))
				} else {
					var _hx_tmp_8 string
					_ = _hx_tmp_8
					_hx_tmp_8 = text
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					_hx_tmp_9 = pos2
					var _hx_tmp_10 string
					_ = _hx_tmp_10
					_hx_tmp_10 = _hx_tmp_8
					var _hx_tmp_11 int32
					_ = _hx_tmp_11
					_hx_tmp_11 = _hx_tmp_9
					_hx_tmp_7 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_10, _hx_tmp_11, HxNullable[int32]{}))
				}

				var pluralForm string
				_ = pluralForm
				pluralForm = _hx_tmp_7
				if i > segStart {
					var tmp__hx_reserved_len_2 HxNullable[int32]
					_ = tmp__hx_reserved_len_2
					tmp__hx_reserved_len_2 = HxNullable[int32]{Value: (i - segStart), Valid: true}

					var _hx_tmp_12 HxArray[*Hx_Obj_loreline_nstringpart]
					_ = _hx_tmp_12
					_hx_tmp_12 = parts
					var _hx_tmp_13 string
					_ = _hx_tmp_13
					if (tmp__hx_reserved_len_2).Valid != false {
						_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(text, segStart, tmp__hx_reserved_len_2))
					} else {
						var _hx_tmp_14 string
						_ = _hx_tmp_14
						_hx_tmp_14 = text
						var _hx_tmp_15 int32
						_ = _hx_tmp_15
						_hx_tmp_15 = segStart
						var _hx_tmp_16 string
						_ = _hx_tmp_16
						_hx_tmp_16 = _hx_tmp_14
						var _hx_tmp_17 int32
						_ = _hx_tmp_17
						_hx_tmp_17 = _hx_tmp_15
						_hx_tmp_13 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_16, _hx_tmp_17, HxNullable[int32]{}))
					}

					_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_12, Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_13}, nil, nil)))
				}

				var _hx_tmp_18 HxArray[*Hx_Obj_loreline_nstringpart]
				_ = _hx_tmp_18
				_hx_tmp_18 = parts
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_18, Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
					_hx_up := this.VTable.Hx_Field_makePluralCall(countExpr, singular, pluralForm, pos)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_nexpr
					}
					return nil
				})()}, nil, nil)))
				i = (closePos + 1)
				segStart = i
				continue
			}
		}

		if c1 == 124 {
			var wordStart int32
			_ = wordStart
			wordStart = i
			for wordStart > segStart {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, (wordStart - 1)))
				var _hx_tmp_19 int32
				_ = _hx_tmp_19
				if (tmp_c_1).Valid == false {
					_hx_tmp_19 = -1
				} else {
					_hx_tmp_19 = (tmp_c_1).Value
				}

				var wc int32
				_ = wc
				wc = _hx_tmp_19
				if (((((wc == 32) || (wc == 9)) || (wc == 10)) || (wc == 124)) || (wc == 40)) || (wc == 41) {
					break
				}

				wordStart--
			}

			var wordEnd int32
			_ = wordEnd
			wordEnd = (i + 1)
			for wordEnd < _hx_reserved_len {
				var tmp_c_1 HxNullable[int32]
				_ = tmp_c_1
				tmp_c_1 = (Hx_Field_go_haxe_hxstring_charCodeAt(text, wordEnd))
				var _hx_tmp_20 int32
				_ = _hx_tmp_20
				if (tmp_c_1).Valid == false {
					_hx_tmp_20 = -1
				} else {
					_hx_tmp_20 = (tmp_c_1).Value
				}

				var wc int32
				_ = wc
				wc = _hx_tmp_20
				if (((((wc == 32) || (wc == 9)) || (wc == 10)) || (wc == 124)) || (wc == 40)) || (wc == 41) {
					break
				}

				wordEnd++
			}

			var len1 HxNullable[int32]
			_ = len1
			len1 = HxNullable[int32]{Value: (i - wordStart), Valid: true}

			var _hx_tmp_21 string
			_ = _hx_tmp_21
			if (len1).Valid != false {
				_hx_tmp_21 = (Hx_Field_go_haxe_hxstring_substr(text, wordStart, len1))
			} else {
				var _hx_tmp_22 string
				_ = _hx_tmp_22
				_hx_tmp_22 = text
				var _hx_tmp_23 int32
				_ = _hx_tmp_23
				_hx_tmp_23 = wordStart
				var _hx_tmp_24 string
				_ = _hx_tmp_24
				_hx_tmp_24 = _hx_tmp_22
				var _hx_tmp_25 int32
				_ = _hx_tmp_25
				_hx_tmp_25 = _hx_tmp_23
				_hx_tmp_21 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_24, _hx_tmp_25, HxNullable[int32]{}))
			}

			var word1 string
			_ = word1
			word1 = _hx_tmp_21
			var pos1 int32
			_ = pos1
			pos1 = (i + 1)
			var len2 HxNullable[int32]
			_ = len2
			len2 = HxNullable[int32]{Value: ((wordEnd - i) - 1), Valid: true}

			var _hx_tmp_26 string
			_ = _hx_tmp_26
			if (len2).Valid != false {
				_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(text, pos1, len2))
			} else {
				var _hx_tmp_27 string
				_ = _hx_tmp_27
				_hx_tmp_27 = text
				var _hx_tmp_28 int32
				_ = _hx_tmp_28
				_hx_tmp_28 = pos1
				var _hx_tmp_29 string
				_ = _hx_tmp_29
				_hx_tmp_29 = _hx_tmp_27
				var _hx_tmp_30 int32
				_ = _hx_tmp_30
				_hx_tmp_30 = _hx_tmp_28
				_hx_tmp_26 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_29, _hx_tmp_30, HxNullable[int32]{}))
			}

			var word2 string
			_ = word2
			word2 = _hx_tmp_26
			var _hx_tmp_31 bool
			_ = _hx_tmp_31
			if int32(utf8.RuneCountInString(word1)) > 0 {
				_hx_tmp_31 = (int32(utf8.RuneCountInString(word2)) > 0)
			} else {
				_hx_tmp_31 = false
			}

			if _hx_tmp_31 {
				if wordStart > segStart {
					var tmp__hx_reserved_len_1 HxNullable[int32]
					_ = tmp__hx_reserved_len_1
					tmp__hx_reserved_len_1 = HxNullable[int32]{Value: (wordStart - segStart), Valid: true}

					var _hx_tmp_32 HxArray[*Hx_Obj_loreline_nstringpart]
					_ = _hx_tmp_32
					_hx_tmp_32 = parts
					var _hx_tmp_33 string
					_ = _hx_tmp_33
					if (tmp__hx_reserved_len_1).Valid != false {
						_hx_tmp_33 = (Hx_Field_go_haxe_hxstring_substr(text, segStart, tmp__hx_reserved_len_1))
					} else {
						var _hx_tmp_34 string
						_ = _hx_tmp_34
						_hx_tmp_34 = text
						var _hx_tmp_35 int32
						_ = _hx_tmp_35
						_hx_tmp_35 = segStart
						var _hx_tmp_36 string
						_ = _hx_tmp_36
						_hx_tmp_36 = _hx_tmp_34
						var _hx_tmp_37 int32
						_ = _hx_tmp_37
						_hx_tmp_37 = _hx_tmp_35
						_hx_tmp_33 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_36, _hx_tmp_37, HxNullable[int32]{}))
					}

					_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_32, Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_33}, nil, nil)))
				}

				var _hx_tmp_38 HxArray[*Hx_Obj_loreline_nstringpart]
				_ = _hx_tmp_38
				_hx_tmp_38 = parts
				_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_38, Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Expr{(func() *Hx_Obj_loreline_nexpr {
					_hx_up := this.VTable.Hx_Field_makePluralCall(countExpr, word1, word2, pos)
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_nexpr
					}
					return nil
				})()}, nil, nil)))
				i = wordEnd
				segStart = i
				continue
			}
		}

		i++
	}

	if segStart < _hx_reserved_len {
		var len1 HxNullable[int32]
		_ = len1
		len1 = HxNullable[int32]{Value: (_hx_reserved_len - segStart), Valid: true}

		var _hx_tmp_39 HxArray[*Hx_Obj_loreline_nstringpart]
		_ = _hx_tmp_39
		_hx_tmp_39 = parts
		var _hx_tmp_40 string
		_ = _hx_tmp_40
		if (len1).Valid != false {
			_hx_tmp_40 = (Hx_Field_go_haxe_hxstring_substr(text, segStart, len1))
		} else {
			var _hx_tmp_41 string
			_ = _hx_tmp_41
			_hx_tmp_41 = text
			var _hx_tmp_42 int32
			_ = _hx_tmp_42
			_hx_tmp_42 = segStart
			var _hx_tmp_43 string
			_ = _hx_tmp_43
			_hx_tmp_43 = _hx_tmp_41
			var _hx_tmp_44 int32
			_ = _hx_tmp_44
			_hx_tmp_44 = _hx_tmp_42
			_hx_tmp_40 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_43, _hx_tmp_44, HxNullable[int32]{}))
		}

		_ = (Hx_Array_Push[*Hx_Obj_loreline_nstringpart](_hx_tmp_39, Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Raw{_hx_tmp_40}, nil, nil)))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_makePluralCall(countExpr *Hx_Obj_loreline_nexpr, singular string, pluralForm string, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_ncall {
	var target *Hx_Obj_loreline_naccess
	_ = target
	target = Hx_Obj_loreline_naccess_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, nil, "plural", nil, nil)
	var singularLiteral *Hx_Obj_loreline_nstringliteral
	_ = singularLiteral
	singularLiteral = Hx_Obj_loreline_nstringliteral_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, ((int32)(1)), HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Raw{singular}, nil, nil)), nil, nil)
	var pluralLiteral *Hx_Obj_loreline_nstringliteral
	_ = pluralLiteral
	pluralLiteral = Hx_Obj_loreline_nstringliteral_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, ((int32)(1)), HxMakeArray[*Hx_Obj_loreline_nstringpart](Hx_Obj_loreline_nstringpart_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, Hx_Enum_loreline_stringparttype_Raw{pluralForm}, nil, nil)), nil, nil)
	var _hx_tmp_0 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_0
	_hx_tmp_0 = countExpr
	var _hx_tmp_1 *Hx_Obj_loreline_nexpr
	_ = _hx_tmp_1
	_hx_tmp_1 = (func() *Hx_Obj_loreline_nexpr {
		_hx_up := singularLiteral
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})()
	return Hx_Obj_loreline_ncall_CreateInstance(Hx_Field_loreline__node_nodeid_impl__UNDEFINED, pos, (func() *Hx_Obj_loreline_nexpr {
		_hx_up := target
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})(), HxMakeArray[*Hx_Obj_loreline_nexpr](_hx_tmp_0, _hx_tmp_1, (func() *Hx_Obj_loreline_nexpr {
		_hx_up := pluralLiteral
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_nexpr
		}
		return nil
	})()), nil, nil)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_stripStringIndent(content string) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = content
	if (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "\n", HxNullable[int32]{})) == -1 {
		return content
	}

	var lines HxArray[string]
	_ = lines
	lines = (Hx_Field_go_haxe_hxstring_split(content, "\n"))
	var minIndent int32
	_ = minIndent
	minIndent = -1
	{
		var _g int32
		_ = _g
		_g = 1
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
			var line string
			_ = line
			line = lines.Get(i)
			if int32(utf8.RuneCountInString(line)) == 0 {
				continue
			}

			var indent int32
			_ = indent
			indent = 0
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = indent
				if !(_hx_tmp_2 < int32(utf8.RuneCountInString(line))) {
					break
				}

				var c HxNullable[int32]
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(line, indent))
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
				if (c1 != 32) && (c1 != 9) {
					break
				}

				indent++
			}

			if (minIndent == -1) || (indent < minIndent) {
				minIndent = indent
			}
		}
	}

	if minIndent > 0 {
		{
			var _g int32
			_ = _g
			_g = 1
			var _g1 int32
			_ = _g1
			_g1 = lines.Len()
			for _g < _g1 {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_4
				if int32(utf8.RuneCountInString(lines.Get(i))) > 0 {
					var str string
					_ = str
					str = lines.Get(i)
					var endIndex HxNullable[int32]
					_ = endIndex
					endIndex = HxNullable[int32]{}

					var _hx_tmp_5 string
					_ = _hx_tmp_5
					if (endIndex).Valid != false {
						_hx_tmp_5 = (Hx_Field_go_haxe_hxstring_substring(str, minIndent, endIndex))
					} else {
						var _hx_tmp_6 string
						_ = _hx_tmp_6
						_hx_tmp_6 = str
						var _hx_tmp_7 int32
						_ = _hx_tmp_7
						_hx_tmp_7 = minIndent
						var _hx_tmp_8 string
						_ = _hx_tmp_8
						_hx_tmp_8 = _hx_tmp_6
						var _hx_tmp_9 int32
						_ = _hx_tmp_9
						_hx_tmp_9 = _hx_tmp_7
						_hx_tmp_5 = (Hx_Field_go_haxe_hxstring_substring(_hx_tmp_8, _hx_tmp_9, HxNullable[int32]{}))
					}

					lines.Set(i, _hx_tmp_5)
				}
			}
		}

		content = (Hx_Array_Join[string](lines, "\n"))
	}

	return content
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_stripStringComments(content string) string {
	var chars HxArray[int32]
	_ = chars
	chars = Hx_Field_loreline_utf8_toCodes(content)
	var result *Hx_Obj_stringbuf
	_ = result
	result = Hx_Obj_stringbuf_CreateInstance()
	var _hx_reserved_len int32
	_ = _hx_reserved_len
	_hx_reserved_len = HxMakeArrayView[any](chars).Len()
	var i int32
	_ = i
	i = 0
	for i < _hx_reserved_len {
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
			_hx_tmp_1 = (_hx_tmp_3 >= HxMakeArrayView[any](chars).Len())
		}

		if _hx_tmp_1 {
			_hx_tmp_0 = -1
		} else {
			_hx_tmp_0 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(i))))
		}

		var c int32
		_ = c
		c = _hx_tmp_0
		var tmp bool
		_ = tmp
		if (c == 47) && ((i + 1) < _hx_reserved_len) {
			var pos int32
			_ = pos
			pos = (i + 1)
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
				_hx_tmp_5 = (_hx_tmp_7 >= HxMakeArrayView[any](chars).Len())
			}

			if _hx_tmp_5 {
				_hx_tmp_4 = -1
			} else {
				_hx_tmp_4 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
			}

			tmp = ((_hx_tmp_4) == 47)
		} else {
			tmp = false
		}

		if tmp {
			for {
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				if i < _hx_reserved_len {
					var _hx_tmp_9 int32
					_ = _hx_tmp_9
					var _hx_tmp_10 bool
					_ = _hx_tmp_10
					if i < 0 {
						_hx_tmp_10 = true
					} else {
						var _hx_tmp_11 int32
						_ = _hx_tmp_11
						_hx_tmp_11 = i
						var _hx_tmp_12 int32
						_ = _hx_tmp_12
						_hx_tmp_12 = _hx_tmp_11
						_hx_tmp_10 = (_hx_tmp_12 >= HxMakeArrayView[any](chars).Len())
					}

					if _hx_tmp_10 {
						_hx_tmp_9 = -1
					} else {
						_hx_tmp_9 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(i))))
					}

					_hx_tmp_8 = ((_hx_tmp_9) != 10)
				} else {
					_hx_tmp_8 = false
				}

				if !(_hx_tmp_8) {
					break
				}

				i++
			}

			continue
		}

		var tmp1 bool
		_ = tmp1
		if (c == 47) && ((i + 1) < _hx_reserved_len) {
			var pos int32
			_ = pos
			pos = (i + 1)
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
				_hx_tmp_14 = (_hx_tmp_16 >= HxMakeArrayView[any](chars).Len())
			}

			if _hx_tmp_14 {
				_hx_tmp_13 = -1
			} else {
				_hx_tmp_13 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
			}

			tmp1 = ((_hx_tmp_13) == 42)
		} else {
			tmp1 = false
		}

		if tmp1 {
			var hadSpaceBefore bool
			_ = hadSpaceBefore
			if i > 0 {
				var pos int32
				_ = pos
				pos = (i - 1)
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
					_hx_tmp_18 = (_hx_tmp_20 >= HxMakeArrayView[any](chars).Len())
				}

				if _hx_tmp_18 {
					_hx_tmp_17 = -1
				} else {
					_hx_tmp_17 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
				}

				hadSpaceBefore = ((_hx_tmp_17) == 32)
			} else {
				hadSpaceBefore = false
			}

			i = (i + 2)
			for {
				var tmp_tmp_1 bool
				_ = tmp_tmp_1
				if (i + 1) < _hx_reserved_len {
					var tmp_tmp1_1 bool
					_ = tmp_tmp1_1
					var _hx_tmp_21 int32
					_ = _hx_tmp_21
					var _hx_tmp_22 bool
					_ = _hx_tmp_22
					if i < 0 {
						_hx_tmp_22 = true
					} else {
						var _hx_tmp_23 int32
						_ = _hx_tmp_23
						_hx_tmp_23 = i
						var _hx_tmp_24 int32
						_ = _hx_tmp_24
						_hx_tmp_24 = _hx_tmp_23
						_hx_tmp_22 = (_hx_tmp_24 >= HxMakeArrayView[any](chars).Len())
					}

					if _hx_tmp_22 {
						_hx_tmp_21 = -1
					} else {
						_hx_tmp_21 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(i))))
					}

					if (_hx_tmp_21) == 42 {
						var pos int32
						_ = pos
						pos = (i + 1)
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
							_hx_tmp_26 = (_hx_tmp_28 >= HxMakeArrayView[any](chars).Len())
						}

						if _hx_tmp_26 {
							_hx_tmp_25 = -1
						} else {
							_hx_tmp_25 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(pos))))
						}

						tmp_tmp1_1 = ((_hx_tmp_25) == 47)
					} else {
						tmp_tmp1_1 = false
					}

					tmp_tmp_1 = !tmp_tmp1_1
				} else {
					tmp_tmp_1 = false
				}

				if !tmp_tmp_1 {
					break
				}

				i++
			}

			i = (i + 2)
			var _hx_tmp_29 bool
			_ = _hx_tmp_29
			if i < _hx_reserved_len {
				var _hx_tmp_30 int32
				_ = _hx_tmp_30
				var _hx_tmp_31 bool
				_ = _hx_tmp_31
				if i < 0 {
					_hx_tmp_31 = true
				} else {
					var _hx_tmp_32 int32
					_ = _hx_tmp_32
					_hx_tmp_32 = i
					var _hx_tmp_33 int32
					_ = _hx_tmp_33
					_hx_tmp_33 = _hx_tmp_32
					_hx_tmp_31 = (_hx_tmp_33 >= HxMakeArrayView[any](chars).Len())
				}

				if _hx_tmp_31 {
					_hx_tmp_30 = -1
				} else {
					_hx_tmp_30 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](chars).Get(i))))
				}

				_hx_tmp_29 = ((_hx_tmp_30) == 32)
			} else {
				_hx_tmp_29 = false
			}

			var hasSpaceAfter bool
			_ = hasSpaceAfter
			hasSpaceAfter = _hx_tmp_29
			if hadSpaceBefore && hasSpaceAfter {
				i++
			}

			if (!hadSpaceBefore && !hasSpaceAfter) && (i < _hx_reserved_len) {
				result.VTable.Hx_Field_addChar(32)
			}

			continue
		}

		result.VTable.Hx_Field_addChar(c)
		i++
	}

	return result.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateCondition(expr *Hx_Obj_loreline_nexpr) bool {
	var value any
	_ = value
	value = this.VTable.Hx_Field_evaluateExpression(expr)
	if Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI) {
		return (((bool)(Hx_Field_go_haxe_hxdynamic_toBool(value))) == true)
	} else {
		if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
			return (int32(utf8.RuneCountInString(Hx_Field_go_haxe_hxdynamic_toString(value))) > 0)
		} else {
			if Hx_Field_loreline_arrays_isArray(value) {
				return (Hx_Field_loreline_arrays_arrayLength(value) > 0)
			} else {
				if Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI) {
					return (((int32)(Hx_Field_go_haxe_hxdynamic_toInt(value))) != 0)
				} else {
					if Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI) {
						return (((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(value))) != ((float64)(0)))
					} else {
						return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(value, nil))))
					}
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_callFunctionValue(fn any, args HxArray[any], pos *Hx_Obj_loreline_position, next func()) any {
	{
		var hx_try_return any
		_ = hx_try_return
		var hx_try_state int
		_ = hx_try_state

		func() {
			defer func() {
				if _hx_ex := Hx_Field_haxe__exception_exception_fields__checkException(recover()); _hx_ex != nil {
					if e, _hx_ok := Hx_Field_haxe_exception_catchValue(_hx_ex, "any").(any); _hx_ok {
						_ = e
						{
							if Hx_Field_std_isOfType(e, Hx_Obj_loreline_lorscript_error_RTTI) {
								var hx_dyncast_274 any
								_ = hx_dyncast_274
								hx_dyncast_274 = Hx_Field_go_haxe_hxdynamic_castClass(e, "Hx_Obj_loreline_lorscript_error")
								var _hx_tmp_3 *Hx_Obj_loreline_lorscript_error
								_ = _hx_tmp_3
								if hx_dyncast_274 != nil {
									_hx_tmp_3 = hx_dyncast_274.(interface {
										Hx_As_Hx_Obj_loreline_lorscript_error() *Hx_Obj_loreline_lorscript_error
									}).Hx_As_Hx_Obj_loreline_lorscript_error()
								} else {
									_hx_tmp_3 = nil
								}

								var lorscriptErr *Hx_Obj_loreline_lorscript_error
								_ = lorscriptErr
								lorscriptErr = _hx_tmp_3
								var _hx_tmp_5 string
								_ = _hx_tmp_5
								_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Error when evaluating function (", Hx_Field_go_haxe_hxdynamic_toString(lorscriptErr.Hx_Field_pmin)), "-")
								var _hx_tmp_4 string
								_ = _hx_tmp_4
								_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_5, Hx_Field_go_haxe_hxdynamic_toString(lorscriptErr.Hx_Field_pmax)), "): ")
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_4, Hx_Field_std_string(lorscriptErr.Hx_Field_e)), pos)))))
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Error when calling function: ", Hx_Field_std_string(e)), pos)))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var result any
				_ = result
				result = Hx_Field_reflect_callMethod(nil, fn, args)
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(result, nil))) {
					_hx_tmp_0 = !Hx_Field_std_isOfType(result, Hx_Obj_loreline_async_RTTI)
				} else {
					_hx_tmp_0 = false
				}

				if _hx_tmp_0 {
					result = this.VTable.Hx_Field_hostValueToRuntime(result)
				}

				var _hx_tmp_1 bool
				_ = _hx_tmp_1
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(result, nil))) {
					_hx_tmp_1 = Hx_Field_std_isOfType(result, Hx_Obj_loreline_async_RTTI)
				} else {
					_hx_tmp_1 = false
				}

				if _hx_tmp_1 {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(next)), ((any)(nil))))) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot call async function in expression", pos)))))
					} else {
						var hx_dyncast_273 any
						_ = hx_dyncast_273
						hx_dyncast_273 = Hx_Field_go_haxe_hxdynamic_castClass(result, "Hx_Obj_loreline_async")
						var _hx_tmp_2 *Hx_Obj_loreline_async
						_ = _hx_tmp_2
						if hx_dyncast_273 != nil {
							_hx_tmp_2 = hx_dyncast_273.(interface{ Hx_As_Hx_Obj_loreline_async() *Hx_Obj_loreline_async }).Hx_As_Hx_Obj_loreline_async()
						} else {
							_hx_tmp_2 = nil
						}

						var asyncResult *Hx_Obj_loreline_async
						_ = asyncResult
						asyncResult = _hx_tmp_2
						asyncResult.Hx_Field_func(next)
					}
				} else {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(next)), ((any)(nil))))) {
						next()
					}
				}

				hx_try_state = 1
				hx_try_return = result
				return
			}
		}()

		return hx_try_return
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolvesExistingBinding(name string) bool {
	var found bool
	_ = found
	found = false
	var _gthis *Hx_Obj_loreline_interpreter
	_ = _gthis
	_gthis = this
	this.VTable.Hx_Field_eachResolutionScope(-1, func(scope *Hx_Obj_loreline_runtimescope) bool {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		if scope.Hx_Field_state != nil {
			_hx_tmp_0 = Hx_Field_loreline_objects_fieldExists(_gthis, scope.Hx_Field_state.Hx_Field_fields, name)
		} else {
			_hx_tmp_0 = false
		}

		if _hx_tmp_0 {
			found = true
			return true
		}

		if scope.Hx_Field_node != nil {
			var key Hx_Typedef_haxe__int64___int64
			_ = key
			key = scope.Hx_Field_node.Hx_Field_id
			var _hx_tmp_2 int64
			_ = _hx_tmp_2
			_hx_tmp_2 = ((int64)(key))
			var stateInNode *Hx_Obj_loreline_runtimestate
			_ = stateInNode
			stateInNode = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
				var _hx_closure_src_658 func(high int32, low int32) any
				_ = _hx_closure_src_658
				_hx_closure_src_658 = _gthis.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
				return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
					var hx_dyncast_275 any
					_ = hx_dyncast_275
					hx_dyncast_275 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_658(high, low), "Hx_Obj_loreline_runtimestate")
					var _hx_tmp_1 *Hx_Obj_loreline_runtimestate
					_ = _hx_tmp_1
					if hx_dyncast_275 != nil {
						_hx_tmp_1 = hx_dyncast_275.(interface {
							Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
						}).Hx_As_Hx_Obj_loreline_runtimestate()
					} else {
						_hx_tmp_1 = nil
					}

					return _hx_tmp_1
				}
			}()(((int32)((_hx_tmp_2 >> int64(((int32)(32)))))), ((int32)(key)))
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if stateInNode != nil {
				_hx_tmp_3 = Hx_Field_loreline_objects_fieldExists(_gthis, stateInNode.Hx_Field_fields, name)
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				found = true
				return true
			}
		}

		return false
	})
	if found {
		return true
	}

	if Hx_Field_loreline_objects_fieldExists(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name) {
		return true
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_659 func(_hx_param_key any) bool
		_ = _hx_closure_src_659
		_hx_closure_src_659 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_659(((any)(_hx_param_key)))
		}
	}()(name) {
		return true
	}

	if func() func(_hx_param_key string) bool {
		var _hx_closure_src_660 func(_hx_param_key any) bool
		_ = _hx_closure_src_660
		_hx_closure_src_660 = this.Hx_Field_topLevelFunctions.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_660(((any)(_hx_param_key)))
		}
	}()(name) {
		return true
	}

	var _hx_tmp_4 string
	_ = _hx_tmp_4
	_hx_tmp_4 = name
	if this.VTable.Hx_Field_resolveBeatByName(_hx_tmp_4, nil) != nil {
		return true
	}

	return false
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateDynamicCall(call *Hx_Obj_loreline_ncall, next func()) any {
	if call.Hx_Field_args.Len() == 0 {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("call() requires a target argument", call.Hx_Field_pos)))))
	}

	var targetValue any
	_ = targetValue
	targetValue = this.VTable.Hx_Field_evaluateExpression(call.Hx_Field_args.Get(0))
	if Hx_Field_loreline_runtimebeatref_beatOf(targetValue) != nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("call() can only call functions. Use beat(...) to call a beat.", call.Hx_Field_pos)))))
	}

	var fn any
	_ = fn
	if Hx_Field_reflect_isFunction(targetValue) {
		fn = targetValue
	} else {
		if Hx_Field_std_isOfType(targetValue, Hx_Obj_string_RTTI) {
			var name string
			_ = name
			name = Hx_Field_go_haxe_hxdynamic_toString(targetValue)
			var _hx_tmp_0 string
			_ = _hx_tmp_0
			_hx_tmp_0 = name
			if this.VTable.Hx_Field_resolveBeatByName(_hx_tmp_0, nil) != nil {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("call() can only call functions. Use beat(...) to call a beat.", call.Hx_Field_pos)))))
			}

			if func() func(_hx_param_key string) bool {
				var _hx_closure_src_661 func(_hx_param_key any) bool
				_ = _hx_closure_src_661
				_hx_closure_src_661 = this.Hx_Field_topLevelFunctions.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_661(((any)(_hx_param_key)))
				}
			}()(name) {
				fn = this.Hx_Field_topLevelFunctions.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
			}
		}
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(fn, nil))) {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 1
			var _g2 int32
			_ = _g2
			_g2 = call.Hx_Field_args.Len()
			for _g1 < _g2 {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g1
				_g1 = (_g1 + 1)
				var i int32
				_ = i
				i = _hx_tmp_1
				var _hx_tmp_2 HxArray[any]
				_ = _hx_tmp_2
				_hx_tmp_2 = HxMakeArrayView[any](_g)
				_ = (Hx_Array_Push[any](_hx_tmp_2, this.VTable.Hx_Field_evaluateExpression(call.Hx_Field_args.Get(i))))
			}
		}

		var values HxArray[any]
		_ = values
		values = _g
		return this.VTable.Hx_Field_callFunctionValue(fn, values, call.Hx_Field_pos, next)
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("call() target is not a function: ", this.VTable.Hx_Field_valueToString(targetValue)), call.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateFunctionCall(call *Hx_Obj_loreline_ncall, next func()) any {
	if Hx_Field_std_isOfType(call.Hx_Field_target, Hx_Obj_loreline_naccess_RTTI) {
		var access *Hx_Obj_loreline_naccess
		_ = access
		access = call.Hx_Field_target.VTable.(interface {
			Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
		}).Hx_As_Hx_Obj_loreline_naccess()
		if access.Hx_Field_target == nil {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			if access.Hx_Field_name == "call" {
				_hx_tmp_0 = !this.VTable.Hx_Field_resolvesExistingBinding("call")
			} else {
				_hx_tmp_0 = false
			}

			if _hx_tmp_0 {
				return this.VTable.Hx_Field_evaluateDynamicCall(call, next)
			}

			var target any
			_ = target
			target = this.VTable.Hx_Field_evaluateExpression(call.Hx_Field_target)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(target, nil))) {
				if Hx_Field_reflect_isFunction(target) {
					var _g HxArray[any]
					_ = _g
					_g = HxMakeArray[any]().Dyn()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[*Hx_Obj_loreline_nexpr]
						_ = _g2
						_g2 = call.Hx_Field_args
						for {
							var _hx_tmp_1 int32
							_ = _hx_tmp_1
							_hx_tmp_1 = _g1
							if !(_hx_tmp_1 < _g2.Len()) {
								break
							}

							var arg *Hx_Obj_loreline_nexpr
							_ = arg
							arg = _g2.Get(_g1)
							_g1++
							var _hx_tmp_2 HxArray[any]
							_ = _hx_tmp_2
							_hx_tmp_2 = HxMakeArrayView[any](_g)
							_ = (Hx_Array_Push[any](_hx_tmp_2, this.VTable.Hx_Field_evaluateExpression(arg)))
						}
					}

					var args HxArray[any]
					_ = args
					args = _g
					return this.VTable.Hx_Field_callFunctionValue(target, args, call.Hx_Field_pos, next)
				}
			}
		} else {
			var obj any
			_ = obj
			obj = Hx_Field_loreline_runtimecharacterref_fieldsOf(this.VTable.Hx_Field_evaluateExpression(access.Hx_Field_target))
			var helper any
			_ = helper
			var objBeat *Hx_Obj_loreline_nbeatdecl
			_ = objBeat
			objBeat = Hx_Field_loreline_runtimebeatref_beatOf(obj)
			if Hx_Field_std_isOfType(obj, Hx_Obj_string_RTTI) {
				helper = Hx_Field_loreline_objects_getStringHelper(this, access.Hx_Field_name)
			} else {
				if Hx_Field_loreline_arrays_isArray(obj) {
					helper = Hx_Field_loreline_objects_getArrayHelper(this, access.Hx_Field_name)
				} else {
					if objBeat != nil {
						helper = Hx_Field_loreline_objects_getBeatHelper(this, access.Hx_Field_name)
					} else {
						if Hx_Field_loreline_objects_isFields(obj) {
							helper = Hx_Field_loreline_objects_getMapHelper(this, access.Hx_Field_name)
						}
					}
				}
			}

			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(helper, nil))) {
				_hx_tmp_3 = Hx_Field_reflect_isFunction(helper)
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				var _g HxArray[any]
				_ = _g
				_g = HxMakeArray[any]().Dyn()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 HxArray[*Hx_Obj_loreline_nexpr]
					_ = _g2
					_g2 = call.Hx_Field_args
					for {
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g1
						if !(_hx_tmp_4 < _g2.Len()) {
							break
						}

						var arg *Hx_Obj_loreline_nexpr
						_ = arg
						arg = _g2.Get(_g1)
						_g1++
						var _hx_tmp_5 HxArray[any]
						_ = _hx_tmp_5
						_hx_tmp_5 = HxMakeArrayView[any](_g)
						_ = (Hx_Array_Push[any](_hx_tmp_5, this.VTable.Hx_Field_evaluateExpression(arg)))
					}
				}

				var args HxArray[any]
				_ = args
				args = _g
				var _hx_tmp_6 HxArray[any]
				_ = _hx_tmp_6
				_hx_tmp_6 = HxMakeArrayView[any](args)
				var _hx_tmp_7 any
				_ = _hx_tmp_7
				if objBeat != nil {
					_hx_tmp_7 = ((any)(objBeat))
				} else {
					_hx_tmp_7 = obj
				}

				(Hx_Array_Insert[any](_hx_tmp_6, 0, _hx_tmp_7))
				var result any
				_ = result
				result = Hx_Field_reflect_callMethod(nil, helper, args)
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(result, nil))) {
					_hx_tmp_8 = Hx_Field_std_isOfType(result, Hx_Obj_loreline_async_RTTI)
				} else {
					_hx_tmp_8 = false
				}

				if _hx_tmp_8 {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(next)), ((any)(nil))))) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot call async function in expression", call.Hx_Field_pos)))))
					} else {
						var hx_dyncast_276 any
						_ = hx_dyncast_276
						hx_dyncast_276 = Hx_Field_go_haxe_hxdynamic_castClass(result, "Hx_Obj_loreline_async")
						var _hx_tmp_9 *Hx_Obj_loreline_async
						_ = _hx_tmp_9
						if hx_dyncast_276 != nil {
							_hx_tmp_9 = hx_dyncast_276.(interface{ Hx_As_Hx_Obj_loreline_async() *Hx_Obj_loreline_async }).Hx_As_Hx_Obj_loreline_async()
						} else {
							_hx_tmp_9 = nil
						}

						var asyncResult *Hx_Obj_loreline_async
						_ = asyncResult
						asyncResult = _hx_tmp_9
						asyncResult.Hx_Field_func(next)
					}
				} else {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(next)), ((any)(nil))))) {
						next()
					}
				}

				return result
			} else {
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(helper, nil))) {
					var _hx_tmp_11 bool
					_ = _hx_tmp_11
					var _hx_tmp_12 bool
					_ = _hx_tmp_12
					if Hx_Field_std_isOfType(obj, Hx_Obj_string_RTTI) {
						_hx_tmp_12 = true
					} else {
						_hx_tmp_12 = Hx_Field_loreline_arrays_isArray(obj)
					}

					if _hx_tmp_12 {
						_hx_tmp_11 = true
					} else {
						_hx_tmp_11 = Hx_Field_loreline_objects_isFields(obj)
					}

					_hx_tmp_10 = !(_hx_tmp_11)
				} else {
					_hx_tmp_10 = false
				}

				if _hx_tmp_10 {
					var method any
					_ = method
					method = Hx_Field_reflect_getProperty(obj, access.Hx_Field_name)
					if Hx_Field_reflect_isFunction(method) {
						var _g HxArray[any]
						_ = _g
						_g = HxMakeArray[any]().Dyn()
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							var _g2 HxArray[*Hx_Obj_loreline_nexpr]
							_ = _g2
							_g2 = call.Hx_Field_args
							for {
								var _hx_tmp_13 int32
								_ = _hx_tmp_13
								_hx_tmp_13 = _g1
								if !(_hx_tmp_13 < _g2.Len()) {
									break
								}

								var arg *Hx_Obj_loreline_nexpr
								_ = arg
								arg = _g2.Get(_g1)
								_g1++
								var _hx_tmp_14 HxArray[any]
								_ = _hx_tmp_14
								_hx_tmp_14 = HxMakeArrayView[any](_g)
								_ = (Hx_Array_Push[any](_hx_tmp_14, this.VTable.Hx_Field_evaluateExpression(arg)))
							}
						}

						var args HxArray[any]
						_ = args
						args = _g
						var result any
						_ = result
						result = Hx_Field_reflect_callMethod(obj, method, args)
						var _hx_tmp_15 bool
						_ = _hx_tmp_15
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(result, nil))) {
							_hx_tmp_15 = Hx_Field_std_isOfType(result, Hx_Obj_loreline_async_RTTI)
						} else {
							_hx_tmp_15 = false
						}

						if _hx_tmp_15 {
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(next)), ((any)(nil))))) {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Cannot call async function in expression", call.Hx_Field_pos)))))
							} else {
								var hx_dyncast_277 any
								_ = hx_dyncast_277
								hx_dyncast_277 = Hx_Field_go_haxe_hxdynamic_castClass(result, "Hx_Obj_loreline_async")
								var _hx_tmp_16 *Hx_Obj_loreline_async
								_ = _hx_tmp_16
								if hx_dyncast_277 != nil {
									_hx_tmp_16 = hx_dyncast_277.(interface{ Hx_As_Hx_Obj_loreline_async() *Hx_Obj_loreline_async }).Hx_As_Hx_Obj_loreline_async()
								} else {
									_hx_tmp_16 = nil
								}

								var asyncResult *Hx_Obj_loreline_async
								_ = asyncResult
								asyncResult = _hx_tmp_16
								asyncResult.Hx_Field_func(next)
							}
						} else {
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(((any)(next)), ((any)(nil))))) {
								next()
							}
						}

						return result
					}
				}
			}
		}
	}

	var tmp HxNullable[string]
	_ = tmp
	tmp = HxNullableString(this.VTable.Hx_Field_printLoreline((func() *Hx_Obj_loreline_node {
		_hx_up := call.Hx_Field_target
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_loreline_node
		}
		return nil
	})()))
	var _hx_tmp_17 string
	_ = _hx_tmp_17
	if (tmp).Valid != false {
		_hx_tmp_17 = (tmp).Value
	} else {
		_hx_tmp_17 = Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_663 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_663
			_hx_closure_src_663 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_663(((any)(o)))))
			}
		}()(call.Hx_Field_target))))
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid call target: ", (_hx_tmp_17)), call.Hx_Field_pos)))))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateArrayLiteral(expr HxArray[any]) any {
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
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(expr)) {
				break
			}

			var elem any
			_ = elem
			elem = expr.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			var hx_dyncast_278 any
			_ = hx_dyncast_278
			hx_dyncast_278 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_nexpr")
			var _hx_tmp_2 *Hx_Obj_loreline_nexpr
			_ = _hx_tmp_2
			if hx_dyncast_278 != nil {
				_hx_tmp_2 = hx_dyncast_278.(interface{ Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr }).Hx_As_Hx_Obj_loreline_nexpr()
			} else {
				_hx_tmp_2 = nil
			}

			_ = (Hx_Array_Push[any](_hx_tmp_1, this.VTable.Hx_Field_evaluateExpression(_hx_tmp_2)))
		}
	}

	return ((any)(_g))
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateObjectLiteral(expr HxArray[*Hx_Obj_loreline_nobjectfield]) any {
	var obj *Hx_Obj_haxe_imap
	_ = obj
	obj = (func() *Hx_Obj_haxe_imap {
		_hx_icast_279 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_279 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_279, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_279.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < expr.Len()) {
				break
			}

			var field *Hx_Obj_loreline_nobjectfield
			_ = field
			field = expr.Get(_g)
			_g++
			var _hx_tmp_1 string
			_ = _hx_tmp_1
			_hx_tmp_1 = field.Hx_Field_name
			obj.VTable.Hx_Field_set(_hx_tmp_1, this.VTable.Hx_Field_evaluateExpression(field.Hx_Field_value))
		}
	}

	return any(obj.VTable)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_evaluateExpression(expr *Hx_Obj_loreline_nexpr) any {
	var _g *Hx_Obj_go_haxe_hxclass
	_ = _g
	_g = func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
		var _hx_closure_src_664 func(o any) *Hx_Obj_go_haxe_hxclass
		_ = _hx_closure_src_664
		_hx_closure_src_664 = Hx_Field_type_getClass
		return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
			return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_664(((any)(o)))))
		}
	}()(expr)
	if _g == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unsupported expression type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
			var _hx_closure_src_667 func(o any) *Hx_Obj_go_haxe_hxclass
			_ = _hx_closure_src_667
			_hx_closure_src_667 = Hx_Field_type_getClass
			return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
				return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_667(((any)(o)))))
			}
		}()(expr))))), expr.Hx_Field_pos)))))
	} else {
		switch _g {
		case Hx_Obj_loreline_naccess_RTTI:
			{
				var access *Hx_Obj_loreline_naccess
				_ = access
				access = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
				}).Hx_As_Hx_Obj_loreline_naccess()
				var resolved Hx_Enum_loreline_runtimeaccess
				_ = resolved
				resolved = this.VTable.Hx_Field_resolveAccess(access, access.Hx_Field_target, access.Hx_Field_name)
				return this.VTable.Hx_Field_readAccess(resolved)
			}

		case Hx_Obj_loreline_narrayaccess_RTTI:
			{
				var arrAccess *Hx_Obj_loreline_narrayaccess
				_ = arrAccess
				arrAccess = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
				}).Hx_As_Hx_Obj_loreline_narrayaccess()
				var target any
				_ = target
				target = this.VTable.Hx_Field_evaluateExpression(arrAccess.Hx_Field_target)
				var index any
				_ = index
				index = this.VTable.Hx_Field_evaluateExpression(arrAccess.Hx_Field_index)
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				if Hx_Field_loreline_arrays_isArray(target) {
					var _hx_tmp_1 bool
					_ = _hx_tmp_1
					if Hx_Field_std_isOfType(index, Hx_Abstract_int_RTTI) {
						_hx_tmp_1 = true
					} else {
						_hx_tmp_1 = Hx_Field_std_isOfType(index, Hx_Abstract_float_RTTI)
					}

					_hx_tmp_0 = (_hx_tmp_1)
				} else {
					_hx_tmp_0 = false
				}

				if _hx_tmp_0 {
					var i int32
					_ = i
					i = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(index)))))
					var _hx_tmp_2 bool
					_ = _hx_tmp_2
					if i < 0 {
						_hx_tmp_2 = true
					} else {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = i
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _hx_tmp_3
						_hx_tmp_2 = (_hx_tmp_4 >= Hx_Field_loreline_arrays_arrayLength(target))
					}

					if _hx_tmp_2 {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Array index out of bounds: ", Hx_Field_go_haxe_hxdynamic_toString(i)), arrAccess.Hx_Field_pos)))))
					} else {
						return Hx_Field_loreline_arrays_arrayGet(target, i)
					}
				} else {
					if Hx_Field_std_isOfType(index, Hx_Obj_string_RTTI) {
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(target, nil))) {
							var _hx_tmp_5 string
							_ = _hx_tmp_5
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(index, nil))) {
								_hx_tmp_5 = "null"
							} else {
								_hx_tmp_5 = Hx_Field_std_string(index)
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access field '", (_hx_tmp_5)), "' of null"), arrAccess.Hx_Field_pos)))))
						}

						var _hx_tmp_6 *Hx_Obj_loreline_interpreter
						_ = _hx_tmp_6
						_hx_tmp_6 = this
						var _hx_tmp_7 any
						_ = _hx_tmp_7
						_hx_tmp_7 = target
						return Hx_Field_loreline_objects_getField(_hx_tmp_6, _hx_tmp_7, Hx_Field_go_haxe_hxdynamic_toString(index))
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid array access", arrAccess.Hx_Field_pos)))))
					}
				}
			}

		case Hx_Obj_loreline_nbinary_RTTI:
			{
				var bin *Hx_Obj_loreline_nbinary
				_ = bin
				bin = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary
				}).Hx_As_Hx_Obj_loreline_nbinary()
				var left any
				_ = left
				left = this.VTable.Hx_Field_evaluateExpression(bin.Hx_Field_left)
				var right any
				_ = right
				right = this.VTable.Hx_Field_evaluateExpression(bin.Hx_Field_right)
				return this.VTable.Hx_Field_performOperation(bin.Hx_Field_op, left, right, bin.Hx_Field_pos)
			}

		case Hx_Obj_loreline_ncall_RTTI:
			{
				return this.VTable.Hx_Field_evaluateFunctionCall(expr.VTable.(interface{ Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall }).Hx_As_Hx_Obj_loreline_ncall(), nil)
			}

		case Hx_Obj_loreline_nliteral_RTTI:
			{
				var lit *Hx_Obj_loreline_nliteral
				_ = lit
				lit = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral
				}).Hx_As_Hx_Obj_loreline_nliteral()
				switch lit.Hx_Field_literalType.Hx_Field_enumIndex() {
				case 0, 1, 2:
					{
						return lit.Hx_Field_value
					}

				case 3:
					{
						return this.VTable.Hx_Field_evaluateArrayLiteral(HxAnyToArray(lit.Hx_Field_value))
					}

				case 4:
					{
						return this.VTable.Hx_Field_evaluateObjectLiteral(HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](lit.Hx_Field_value))
					}

				default:
					{
						panic(Hx_Field_haxe_exception_thrown("unreachable"))
					}
				}
			}

		case Hx_Obj_loreline_nstringliteral_RTTI:
			{
				var str *Hx_Obj_loreline_nstringliteral
				_ = str
				str = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
				var _hx_tmp_8 bool
				_ = _hx_tmp_8
				if str.Hx_Field_parts.Len() == 1 {
					_hx_tmp_8 = (str.Hx_Field_quotes == ((int32)(0)))
				} else {
					_hx_tmp_8 = false
				}

				if _hx_tmp_8 {
					var tmp__g_1 Hx_Enum_loreline_stringparttype
					_ = tmp__g_1
					tmp__g_1 = str.Hx_Field_parts.Get(0).Hx_Field_partType
					if tmp__g_1.Hx_Field_enumIndex() == 1 {
						var expr *Hx_Obj_loreline_nexpr
						_ = expr
						expr = tmp__g_1.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
						return this.VTable.Hx_Field_evaluateExpression(expr)
					} else {
						return Hx_Field_go_haxe_hxdynamic_getField(this.VTable.Hx_Field_evaluateString(str), "text")
					}
				} else {
					return Hx_Field_go_haxe_hxdynamic_getField(this.VTable.Hx_Field_evaluateString(str), "text")
				}
			}

		case Hx_Obj_loreline_nternary_RTTI:
			{
				var ternary *Hx_Obj_loreline_nternary
				_ = ternary
				ternary = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nternary() *Hx_Obj_loreline_nternary
				}).Hx_As_Hx_Obj_loreline_nternary()
				var condValue bool
				_ = condValue
				condValue = this.VTable.Hx_Field_evaluateCondition(ternary.Hx_Field_condition)
				if condValue {
					return this.VTable.Hx_Field_evaluateExpression(ternary.Hx_Field_trueExpr)
				} else {
					return this.VTable.Hx_Field_evaluateExpression(ternary.Hx_Field_falseExpr)
				}
			}

		case Hx_Obj_loreline_nunary_RTTI:
			{
				var un *Hx_Obj_loreline_nunary
				_ = un
				un = expr.VTable.(interface {
					Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary
				}).Hx_As_Hx_Obj_loreline_nunary()
				var operand any
				_ = operand
				operand = this.VTable.Hx_Field_evaluateExpression(un.Hx_Field_operand)
				switch un.Hx_Field_op.Hx_Field_enumIndex() {
				case 28:
					{
						if Hx_Field_std_isOfType(operand, Hx_Abstract_int_RTTI) {
							var v int32
							_ = v
							v = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(operand))))
							return ((any)(-v))
						} else {
							if Hx_Field_std_isOfType(operand, Hx_Abstract_float_RTTI) {
								var v float64
								_ = v
								v = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(operand))))
								return ((any)(-v))
							} else {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid unary operation", un.Hx_Field_pos)))))
							}
						}
					}

				case 40:
					{
						if Hx_Field_std_isOfType(operand, Hx_Abstract_bool_RTTI) {
							var v bool
							_ = v
							v = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_toBool(operand))))
							return ((any)(!v))
						} else {
							if Hx_Field_std_isOfType(operand, Hx_Obj_string_RTTI) {
								var v string
								_ = v
								v = Hx_Field_go_haxe_hxdynamic_toString(operand)
								var _hx_tmp_9 bool
								_ = _hx_tmp_9
								if v == HxStringNull {
									_hx_tmp_9 = true
								} else {
									_hx_tmp_9 = (int32(utf8.RuneCountInString(v)) == 0)
								}

								return ((any)(_hx_tmp_9))
							} else {
								if Hx_Field_loreline_arrays_isArray(operand) {
									return ((any)((Hx_Field_loreline_arrays_arrayLength(operand) == 0)))
								} else {
									if Hx_Field_std_isOfType(operand, Hx_Abstract_int_RTTI) {
										var v int32
										_ = v
										v = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_toInt(operand))))
										return ((any)((v == 0)))
									} else {
										if Hx_Field_std_isOfType(operand, Hx_Abstract_float_RTTI) {
											var v float64
											_ = v
											v = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(operand))))
											return ((any)((v == ((float64)(0)))))
										} else {
											return ((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(operand, nil))))))
										}
									}
								}
							}
						}
					}

				default:
					{
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid unary operation", un.Hx_Field_pos)))))
					}
				}
			}

		default:
			{
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Unsupported expression type: ", Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
					var _hx_closure_src_668 func(o any) *Hx_Obj_go_haxe_hxclass
					_ = _hx_closure_src_668
					_hx_closure_src_668 = Hx_Field_type_getClass
					return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
						return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_668(((any)(o)))))
					}
				}()(expr))))), expr.Hx_Field_pos)))))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_readAccess(access Hx_Enum_loreline_runtimeaccess) any {
	switch access.Hx_Field_enumIndex() {
	case 0:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_pos
			var obj any
			_ = obj
			obj = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_obj
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_name
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(obj, nil))) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot read field '", name), "' of null"), pos)))))
			}

			return Hx_Field_loreline_objects_getField(this, obj, name)
		}

	case 1:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_pos
			var array any
			_ = array
			array = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_array
			var index int32
			_ = index
			index = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_index
			return Hx_Field_loreline_arrays_arrayGet(array, index)
		}

	case 2:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_CharacterAccess).Hx_Field_pos
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_CharacterAccess).Hx_Field_name
			if func() func(_hx_param_key string) bool {
				var _hx_closure_src_669 func(_hx_param_key any) bool
				_ = _hx_closure_src_669
				_hx_closure_src_669 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_669(((any)(_hx_param_key)))
				}
			}()(name) {
				var character *Hx_Obj_loreline_runtimecharacter
				_ = character
				character = func() func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
					var _hx_closure_src_670 func(_hx_param_key any) any
					_ = _hx_closure_src_670
					_hx_closure_src_670 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
						Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
					}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_runtimecharacter {
						var hx_dyncast_280 any
						_ = hx_dyncast_280
						hx_dyncast_280 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_670(((any)(_hx_param_key))), "Hx_Obj_loreline_runtimecharacter")
						var _hx_tmp_0 *Hx_Obj_loreline_runtimecharacter
						_ = _hx_tmp_0
						if hx_dyncast_280 != nil {
							_hx_tmp_0 = hx_dyncast_280.(interface {
								Hx_As_Hx_Obj_loreline_runtimecharacter() *Hx_Obj_loreline_runtimecharacter
							}).Hx_As_Hx_Obj_loreline_runtimecharacter()
						} else {
							_hx_tmp_0 = nil
						}

						return _hx_tmp_0
					}
				}()(name)
				if character.Hx_Field_characterRef == nil {
					character.Hx_Field_characterRef = Hx_Obj_loreline_runtimecharacterref_CreateInstance(name, character)
				}

				return ((any)(character.Hx_Field_characterRef))
			} else {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Character not found: ", name), pos)))))
			}
		}

	case 3:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_FunctionAccess).Hx_Field_pos
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_FunctionAccess).Hx_Field_name
			if func() func(_hx_param_key string) bool {
				var _hx_closure_src_671 func(_hx_param_key any) bool
				_ = _hx_closure_src_671
				_hx_closure_src_671 = this.Hx_Field_topLevelFunctions.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
				return func(_hx_param_key string) bool {
					return _hx_closure_src_671(((any)(_hx_param_key)))
				}
			}()(name) {
				return this.Hx_Field_topLevelFunctions.VTable.(interface {
					Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
				}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get(name)
			} else {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Function not found: ", name), pos)))))
			}
		}

	case 4:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_BeatAccess).Hx_Field_pos
			var ref *Hx_Obj_loreline_runtimebeatref
			_ = ref
			ref = access.(Hx_Enum_loreline_runtimeaccess_BeatAccess).Hx_Field_ref
			return ((any)(ref))
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_writeAccess(access Hx_Enum_loreline_runtimeaccess, value any) {
	switch access.Hx_Field_enumIndex() {
	case 0:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_pos
			var obj any
			_ = obj
			obj = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_obj
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_FieldAccess).Hx_Field_name
			{
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(obj, nil))) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot set field '", name), "' of null"), pos)))))
				}

				Hx_Field_loreline_objects_setField(this, obj, name, value)
			}
		}

	case 1:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_pos
			var array any
			_ = array
			array = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_array
			var index int32
			_ = index
			index = access.(Hx_Enum_loreline_runtimeaccess_ArrayAccess).Hx_Field_index
			Hx_Field_loreline_arrays_arraySet(array, index, value)
		}

	case 2:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_CharacterAccess).Hx_Field_pos
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_CharacterAccess).Hx_Field_name
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot overwrite character: ", name), pos)))))
		}

	case 3:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_FunctionAccess).Hx_Field_pos
			var name string
			_ = name
			name = access.(Hx_Enum_loreline_runtimeaccess_FunctionAccess).Hx_Field_name
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot overwrite function: ", name), pos)))))
		}

	case 4:
		{
			var pos *Hx_Obj_loreline_position
			_ = pos
			pos = access.(Hx_Enum_loreline_runtimeaccess_BeatAccess).Hx_Field_pos
			var ref *Hx_Obj_loreline_runtimebeatref
			_ = ref
			ref = access.(Hx_Enum_loreline_runtimeaccess_BeatAccess).Hx_Field_ref
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot overwrite beat: ", ref.Hx_Field_beat.Hx_Field_name), pos)))))
		}

	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveAssignmentTarget(target *Hx_Obj_loreline_nexpr) Hx_Enum_loreline_runtimeaccess {
	var _g *Hx_Obj_go_haxe_hxclass
	_ = _g
	_g = func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
		var _hx_closure_src_672 func(o any) *Hx_Obj_go_haxe_hxclass
		_ = _hx_closure_src_672
		_hx_closure_src_672 = Hx_Field_type_getClass
		return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
			return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_672(((any)(o)))))
		}
	}()(target)
	if _g == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid assignment target", target.Hx_Field_pos)))))
	} else {
		switch _g {
		case Hx_Obj_loreline_naccess_RTTI:
			{
				var access *Hx_Obj_loreline_naccess
				_ = access
				access = target.VTable.(interface {
					Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
				}).Hx_As_Hx_Obj_loreline_naccess()
				return this.VTable.Hx_Field_resolveAccess(access, access.Hx_Field_target, access.Hx_Field_name)
			}

		case Hx_Obj_loreline_narrayaccess_RTTI:
			{
				var arrAccess *Hx_Obj_loreline_narrayaccess
				_ = arrAccess
				arrAccess = target.VTable.(interface {
					Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess
				}).Hx_As_Hx_Obj_loreline_narrayaccess()
				var target1 any
				_ = target1
				target1 = this.VTable.Hx_Field_evaluateExpression(arrAccess.Hx_Field_target)
				var index any
				_ = index
				index = this.VTable.Hx_Field_evaluateExpression(arrAccess.Hx_Field_index)
				var _hx_tmp_0 bool
				_ = _hx_tmp_0
				if Hx_Field_loreline_arrays_isArray(target1) {
					var _hx_tmp_1 bool
					_ = _hx_tmp_1
					if Hx_Field_std_isOfType(index, Hx_Abstract_int_RTTI) {
						_hx_tmp_1 = true
					} else {
						_hx_tmp_1 = Hx_Field_std_isOfType(index, Hx_Abstract_float_RTTI)
					}

					_hx_tmp_0 = (_hx_tmp_1)
				} else {
					_hx_tmp_0 = false
				}

				if _hx_tmp_0 {
					var i int32
					_ = i
					i = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(index)))))
					return Hx_Enum_loreline_runtimeaccess_ArrayAccess{arrAccess.Hx_Field_pos, target1, i}
				} else {
					if Hx_Field_std_isOfType(index, Hx_Obj_string_RTTI) {
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(target1, nil))) {
							var _hx_tmp_2 string
							_ = _hx_tmp_2
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(index, nil))) {
								_hx_tmp_2 = "null"
							} else {
								_hx_tmp_2 = Hx_Field_std_string(index)
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access field '", (_hx_tmp_2)), "' of null"), arrAccess.Hx_Field_pos)))))
						}

						var _hx_tmp_3 *Hx_Obj_loreline_position
						_ = _hx_tmp_3
						_hx_tmp_3 = arrAccess.Hx_Field_pos
						var _hx_tmp_4 any
						_ = _hx_tmp_4
						_hx_tmp_4 = target1
						return Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_3, _hx_tmp_4, Hx_Field_go_haxe_hxdynamic_toString(index)}
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid array access target", arrAccess.Hx_Field_pos)))))
					}
				}
			}

		default:
			{
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Invalid assignment target", target.Hx_Field_pos)))))
			}
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_resolveAccess(access *Hx_Obj_loreline_naccess, target__isOptional any, name string) Hx_Enum_loreline_runtimeaccess {
	target := HxDefault[*Hx_Obj_loreline_nexpr]()
	_ = target

	if target__isOptional != nil {
		target = HxConvert[*Hx_Obj_loreline_nexpr](target__isOptional)
	} else {
		target = nil
	}
	{
		if target != nil {
			var evaluated any
			_ = evaluated
			evaluated = this.VTable.Hx_Field_evaluateExpression(target)
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(evaluated, nil))) {
				var _hx_tmp_0 HxNullable[string]
				_ = _hx_tmp_0
				if func() func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
					var _hx_closure_src_673 func(o any) *Hx_Obj_go_haxe_hxclass
					_ = _hx_closure_src_673
					_hx_closure_src_673 = Hx_Field_type_getClass
					return func(o *Hx_Obj_loreline_nexpr) *Hx_Obj_go_haxe_hxclass {
						return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_673(((any)(o)))))
					}
				}()(target) == Hx_Obj_loreline_naccess_RTTI {
					_hx_tmp_0 = HxNullableString(target.VTable.(interface {
						Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess
					}).Hx_As_Hx_Obj_loreline_naccess().Hx_Field_name)
				} else {
					_hx_tmp_0 = HxNullable[string]{Value: HxStringNull}
				}

				var targetName HxNullable[string]
				_ = targetName
				targetName = _hx_tmp_0
				var _hx_tmp_1 string
				_ = _hx_tmp_1
				if (targetName).Valid != false {
					var _hx_tmp_2 string
					_ = _hx_tmp_2
					_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access field '", name), "' of undefined variable '")
					var _hx_tmp_3 string
					_ = _hx_tmp_3
					_hx_tmp_3 = _hx_tmp_2
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_3, (targetName).Value), "'")
				} else {
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access field '", name), "' of null value")
				}

				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(_hx_tmp_1, target.Hx_Field_pos)))))
			}

			var _hx_tmp_4 *Hx_Obj_loreline_position
			_ = _hx_tmp_4
			_hx_tmp_4 = target.Hx_Field_pos
			return Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_4, Hx_Field_loreline_runtimecharacterref_fieldsOf(evaluated), name}
		}

		var scopeAccess Hx_Enum_loreline_runtimeaccess
		_ = scopeAccess
		var _gthis *Hx_Obj_loreline_interpreter
		_ = _gthis
		_gthis = this
		this.VTable.Hx_Field_eachResolutionScope(-1, func(scope *Hx_Obj_loreline_runtimescope) bool {
			if scope.Hx_Field_state != nil {
				if Hx_Field_loreline_objects_fieldExists(_gthis, scope.Hx_Field_state.Hx_Field_fields, name) {
					var _hx_tmp_5 *Hx_Obj_loreline_position
					_ = _hx_tmp_5
					if access != nil {
						_hx_tmp_5 = access.Hx_Field_pos
					} else {
						_hx_tmp_5 = nil
					}

					var tmp *Hx_Obj_loreline_position
					_ = tmp
					tmp = _hx_tmp_5
					var pos *Hx_Obj_loreline_position
					_ = pos
					if tmp != nil {
						pos = tmp
					} else {
						var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
						_ = tmp_tmp_1
						tmp_tmp_1 = _gthis.VTable.Hx_Field_get_currentScope()
						var _hx_tmp_6 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_6
						if tmp_tmp_1 != nil {
							_hx_tmp_6 = tmp_tmp_1.Hx_Field_node
						} else {
							_hx_tmp_6 = nil
						}

						var tmp1 *Hx_Obj_loreline_astnode
						_ = tmp1
						tmp1 = _hx_tmp_6
						var _hx_tmp_7 *Hx_Obj_loreline_position
						_ = _hx_tmp_7
						if tmp1 != nil {
							_hx_tmp_7 = tmp1.Hx_Field_pos
						} else {
							_hx_tmp_7 = nil
						}

						pos = _hx_tmp_7
					}

					var _hx_tmp_8 *Hx_Obj_loreline_position
					_ = _hx_tmp_8
					if pos != nil {
						_hx_tmp_8 = pos
					} else {
						_hx_tmp_8 = _gthis.Hx_Field_context.Hx_Field_script.Hx_Field_pos
					}

					scopeAccess = Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_8, scope.Hx_Field_state.Hx_Field_fields, name}

					return true
				}
			}

			if scope.Hx_Field_node != nil {
				var key Hx_Typedef_haxe__int64___int64
				_ = key
				key = scope.Hx_Field_node.Hx_Field_id
				var _hx_tmp_10 int64
				_ = _hx_tmp_10
				_hx_tmp_10 = ((int64)(key))
				var stateInNode *Hx_Obj_loreline_runtimestate
				_ = stateInNode
				stateInNode = func() func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
					var _hx_closure_src_674 func(high int32, low int32) any
					_ = _hx_closure_src_674
					_hx_closure_src_674 = _gthis.Hx_Field_context.Hx_Field_nodeStates.Hx_Field_map.VTable.Hx_Field__get
					return func(high int32, low int32) *Hx_Obj_loreline_runtimestate {
						var hx_dyncast_281 any
						_ = hx_dyncast_281
						hx_dyncast_281 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_674(high, low), "Hx_Obj_loreline_runtimestate")
						var _hx_tmp_9 *Hx_Obj_loreline_runtimestate
						_ = _hx_tmp_9
						if hx_dyncast_281 != nil {
							_hx_tmp_9 = hx_dyncast_281.(interface {
								Hx_As_Hx_Obj_loreline_runtimestate() *Hx_Obj_loreline_runtimestate
							}).Hx_As_Hx_Obj_loreline_runtimestate()
						} else {
							_hx_tmp_9 = nil
						}

						return _hx_tmp_9
					}
				}()(((int32)((_hx_tmp_10 >> int64(((int32)(32)))))), ((int32)(key)))
				if stateInNode != nil {
					if Hx_Field_loreline_objects_fieldExists(_gthis, stateInNode.Hx_Field_fields, name) {
						var _hx_tmp_11 *Hx_Obj_loreline_position
						_ = _hx_tmp_11
						if access != nil {
							_hx_tmp_11 = access.Hx_Field_pos
						} else {
							_hx_tmp_11 = nil
						}

						var tmp *Hx_Obj_loreline_position
						_ = tmp
						tmp = _hx_tmp_11
						var pos *Hx_Obj_loreline_position
						_ = pos
						if tmp != nil {
							pos = tmp
						} else {
							var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
							_ = tmp_tmp_1
							tmp_tmp_1 = _gthis.VTable.Hx_Field_get_currentScope()
							var _hx_tmp_12 *Hx_Obj_loreline_astnode
							_ = _hx_tmp_12
							if tmp_tmp_1 != nil {
								_hx_tmp_12 = tmp_tmp_1.Hx_Field_node
							} else {
								_hx_tmp_12 = nil
							}

							var tmp1 *Hx_Obj_loreline_astnode
							_ = tmp1
							tmp1 = _hx_tmp_12
							var _hx_tmp_13 *Hx_Obj_loreline_position
							_ = _hx_tmp_13
							if tmp1 != nil {
								_hx_tmp_13 = tmp1.Hx_Field_pos
							} else {
								_hx_tmp_13 = nil
							}

							pos = _hx_tmp_13
						}

						var _hx_tmp_14 *Hx_Obj_loreline_position
						_ = _hx_tmp_14
						if pos != nil {
							_hx_tmp_14 = pos
						} else {
							_hx_tmp_14 = _gthis.Hx_Field_context.Hx_Field_script.Hx_Field_pos
						}

						scopeAccess = Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_14, stateInNode.Hx_Field_fields, name}

						return true
					}
				}
			}

			return false
		})
		if scopeAccess != nil {
			return scopeAccess
		}

		if Hx_Field_loreline_objects_fieldExists(this, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name) {
			var _hx_tmp_15 *Hx_Obj_loreline_position
			_ = _hx_tmp_15
			if access != nil {
				_hx_tmp_15 = access.Hx_Field_pos
			} else {
				_hx_tmp_15 = nil
			}

			var tmp *Hx_Obj_loreline_position
			_ = tmp
			tmp = _hx_tmp_15
			var pos *Hx_Obj_loreline_position
			_ = pos
			if tmp != nil {
				pos = tmp
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_16 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_16
				if tmp_tmp_1 != nil {
					_hx_tmp_16 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_16 = nil
				}

				var tmp1 *Hx_Obj_loreline_astnode
				_ = tmp1
				tmp1 = _hx_tmp_16
				var _hx_tmp_17 *Hx_Obj_loreline_position
				_ = _hx_tmp_17
				if tmp1 != nil {
					_hx_tmp_17 = tmp1.Hx_Field_pos
				} else {
					_hx_tmp_17 = nil
				}

				pos = _hx_tmp_17
			}

			var _hx_tmp_18 *Hx_Obj_loreline_position
			_ = _hx_tmp_18
			if pos != nil {
				_hx_tmp_18 = pos
			} else {
				_hx_tmp_18 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			return Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_18, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name}
		}

		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_675 func(_hx_param_key any) bool
			_ = _hx_closure_src_675
			_hx_closure_src_675 = this.Hx_Field_context.Hx_Field_topLevelCharacters.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_675(((any)(_hx_param_key)))
			}
		}()(name) {
			var _hx_tmp_19 *Hx_Obj_loreline_position
			_ = _hx_tmp_19
			if access != nil {
				_hx_tmp_19 = access.Hx_Field_pos
			} else {
				_hx_tmp_19 = nil
			}

			var tmp *Hx_Obj_loreline_position
			_ = tmp
			tmp = _hx_tmp_19
			var pos *Hx_Obj_loreline_position
			_ = pos
			if tmp != nil {
				pos = tmp
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_20 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_20
				if tmp_tmp_1 != nil {
					_hx_tmp_20 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_20 = nil
				}

				var tmp1 *Hx_Obj_loreline_astnode
				_ = tmp1
				tmp1 = _hx_tmp_20
				var _hx_tmp_21 *Hx_Obj_loreline_position
				_ = _hx_tmp_21
				if tmp1 != nil {
					_hx_tmp_21 = tmp1.Hx_Field_pos
				} else {
					_hx_tmp_21 = nil
				}

				pos = _hx_tmp_21
			}

			var _hx_tmp_22 *Hx_Obj_loreline_position
			_ = _hx_tmp_22
			if pos != nil {
				_hx_tmp_22 = pos
			} else {
				_hx_tmp_22 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			return Hx_Enum_loreline_runtimeaccess_CharacterAccess{_hx_tmp_22, name}
		}

		if func() func(_hx_param_key string) bool {
			var _hx_closure_src_676 func(_hx_param_key any) bool
			_ = _hx_closure_src_676
			_hx_closure_src_676 = this.Hx_Field_topLevelFunctions.VTable.(interface {
				Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
			}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
			return func(_hx_param_key string) bool {
				return _hx_closure_src_676(((any)(_hx_param_key)))
			}
		}()(name) {
			var _hx_tmp_23 *Hx_Obj_loreline_position
			_ = _hx_tmp_23
			if access != nil {
				_hx_tmp_23 = access.Hx_Field_pos
			} else {
				_hx_tmp_23 = nil
			}

			var tmp *Hx_Obj_loreline_position
			_ = tmp
			tmp = _hx_tmp_23
			var pos *Hx_Obj_loreline_position
			_ = pos
			if tmp != nil {
				pos = tmp
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_24 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_24
				if tmp_tmp_1 != nil {
					_hx_tmp_24 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_24 = nil
				}

				var tmp1 *Hx_Obj_loreline_astnode
				_ = tmp1
				tmp1 = _hx_tmp_24
				var _hx_tmp_25 *Hx_Obj_loreline_position
				_ = _hx_tmp_25
				if tmp1 != nil {
					_hx_tmp_25 = tmp1.Hx_Field_pos
				} else {
					_hx_tmp_25 = nil
				}

				pos = _hx_tmp_25
			}

			var _hx_tmp_26 *Hx_Obj_loreline_position
			_ = _hx_tmp_26
			if pos != nil {
				_hx_tmp_26 = pos
			} else {
				_hx_tmp_26 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			return Hx_Enum_loreline_runtimeaccess_FunctionAccess{_hx_tmp_26, name}
		}

		var _hx_tmp_27 string
		_ = _hx_tmp_27
		_hx_tmp_27 = name
		var beatRef *Hx_Obj_loreline_runtimebeatref
		_ = beatRef
		beatRef = this.VTable.Hx_Field_resolveBeatRefFromNode(_hx_tmp_27, (func() *Hx_Obj_loreline_node {
			_hx_up := access
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		if beatRef != nil {
			var _hx_tmp_28 *Hx_Obj_loreline_position
			_ = _hx_tmp_28
			if access != nil {
				_hx_tmp_28 = access.Hx_Field_pos
			} else {
				_hx_tmp_28 = nil
			}

			var tmp *Hx_Obj_loreline_position
			_ = tmp
			tmp = _hx_tmp_28
			var pos *Hx_Obj_loreline_position
			_ = pos
			if tmp != nil {
				pos = tmp
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_29 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_29
				if tmp_tmp_1 != nil {
					_hx_tmp_29 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_29 = nil
				}

				var tmp1 *Hx_Obj_loreline_astnode
				_ = tmp1
				tmp1 = _hx_tmp_29
				var _hx_tmp_30 *Hx_Obj_loreline_position
				_ = _hx_tmp_30
				if tmp1 != nil {
					_hx_tmp_30 = tmp1.Hx_Field_pos
				} else {
					_hx_tmp_30 = nil
				}

				pos = _hx_tmp_30
			}

			var _hx_tmp_31 *Hx_Obj_loreline_position
			_ = _hx_tmp_31
			if pos != nil {
				_hx_tmp_31 = pos
			} else {
				_hx_tmp_31 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			return Hx_Enum_loreline_runtimeaccess_BeatAccess{_hx_tmp_31, beatRef}
		}

		if !this.Hx_Field_context.Hx_Field_strictAccess {
			var _hx_tmp_32 *Hx_Obj_loreline_position
			_ = _hx_tmp_32
			if access != nil {
				_hx_tmp_32 = access.Hx_Field_pos
			} else {
				_hx_tmp_32 = nil
			}

			var tmp *Hx_Obj_loreline_position
			_ = tmp
			tmp = _hx_tmp_32
			var pos *Hx_Obj_loreline_position
			_ = pos
			if tmp != nil {
				pos = tmp
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_33 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_33
				if tmp_tmp_1 != nil {
					_hx_tmp_33 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_33 = nil
				}

				var tmp1 *Hx_Obj_loreline_astnode
				_ = tmp1
				tmp1 = _hx_tmp_33
				var _hx_tmp_34 *Hx_Obj_loreline_position
				_ = _hx_tmp_34
				if tmp1 != nil {
					_hx_tmp_34 = tmp1.Hx_Field_pos
				} else {
					_hx_tmp_34 = nil
				}

				pos = _hx_tmp_34
			}

			var _hx_tmp_35 *Hx_Obj_loreline_position
			_ = _hx_tmp_35
			if pos != nil {
				_hx_tmp_35 = pos
			} else {
				_hx_tmp_35 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			return Hx_Enum_loreline_runtimeaccess_FieldAccess{_hx_tmp_35, this.Hx_Field_context.Hx_Field_topLevelState.Hx_Field_fields, name}
		}

		var tmp string
		_ = tmp
		tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Undefined variable: ", name)
		var _hx_tmp_36 *Hx_Obj_loreline_position
		_ = _hx_tmp_36
		if access != nil {
			_hx_tmp_36 = access.Hx_Field_pos
		} else {
			_hx_tmp_36 = nil
		}

		var tmp1 *Hx_Obj_loreline_position
		_ = tmp1
		tmp1 = _hx_tmp_36
		var pos *Hx_Obj_loreline_position
		_ = pos
		if tmp1 != nil {
			pos = tmp1
		} else {
			var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
			_ = tmp_tmp_1
			tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
			var _hx_tmp_37 *Hx_Obj_loreline_astnode
			_ = _hx_tmp_37
			if tmp_tmp_1 != nil {
				_hx_tmp_37 = tmp_tmp_1.Hx_Field_node
			} else {
				_hx_tmp_37 = nil
			}

			var tmp_tmp1_1 *Hx_Obj_loreline_astnode
			_ = tmp_tmp1_1
			tmp_tmp1_1 = _hx_tmp_37
			var _hx_tmp_38 *Hx_Obj_loreline_position
			_ = _hx_tmp_38
			if tmp_tmp1_1 != nil {
				_hx_tmp_38 = tmp_tmp1_1.Hx_Field_pos
			} else {
				_hx_tmp_38 = nil
			}

			pos = _hx_tmp_38
		}

		var _hx_tmp_39 *Hx_Obj_loreline_position
		_ = _hx_tmp_39
		if pos != nil {
			_hx_tmp_39 = pos
		} else {
			_hx_tmp_39 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_39)))))
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_getTypeName(t Hx_Enum_valuetype) string {
	switch t.Hx_Field_enumIndex() {
	case 0:
		{
			return "Null"
		}

	case 1:
		{
			return "Int"
		}

	case 2:
		{
			return "Int64"
		}

	case 3:
		{
			return "Float"
		}

	case 4:
		{
			return "Bool"
		}

	case 5:
		{
			return "Object"
		}

	case 6:
		{
			return "Function"
		}

	case 7:
		{
			var c *Hx_Obj_go_haxe_hxclass
			_ = c
			c = t.(Hx_Enum_valuetype_TClass).Hx_Field_c
			return Hx_Field_type_getClassName(c)
		}

	case 8:
		{
			var e *Hx_Obj_go_haxe_hxenum
			_ = e
			e = t.(Hx_Enum_valuetype_TEnum).Hx_Field_e
			return Hx_Field_type_getEnumName(e)
		}

	case 9:
		{
			return "Unknown"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_performOperation(op Hx_Enum_loreline_tokentype, left any, right any, pos *Hx_Obj_loreline_position) any {
	var leftType Hx_Enum_valuetype
	_ = leftType
	leftType = Hx_Field_type_typeof(left)
	var rightType Hx_Enum_valuetype
	_ = rightType
	rightType = Hx_Field_type_typeof(right)
	switch op.Hx_Field_enumIndex() {
	case 27:
		{
			switch leftType.Hx_Field_enumIndex() {
			case 1:
				{
					switch rightType.Hx_Field_enumIndex() {
					case 1, 3:
						{
							var _hx_tmp_0 float64
							_ = _hx_tmp_0
							_hx_tmp_0 = Hx_Field_std_parseFloat(Hx_Field_std_string(left))
							return ((any)((_hx_tmp_0 + Hx_Field_std_parseFloat(Hx_Field_std_string(right)))))
						}

					case 7:
						{
							if rightType.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
								var _hx_tmp_1 string
								_ = _hx_tmp_1
								_hx_tmp_1 = Hx_Field_std_string(left)
								return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_1, Hx_Field_std_string(right))))
							} else {
								var _hx_tmp_2 string
								_ = _hx_tmp_2
								_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
								var tmp string
								_ = tmp
								tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_2, this.VTable.Hx_Field_getTypeName(rightType))
								var tmp1 *Hx_Obj_loreline_position
								_ = tmp1
								tmp1 = pos
								var pos1 *Hx_Obj_loreline_position
								_ = pos1
								if tmp1 != nil {
									pos1 = tmp1
								} else {
									var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
									_ = tmp_tmp_1
									tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
									var _hx_tmp_3 *Hx_Obj_loreline_astnode
									_ = _hx_tmp_3
									if tmp_tmp_1 != nil {
										_hx_tmp_3 = tmp_tmp_1.Hx_Field_node
									} else {
										_hx_tmp_3 = nil
									}

									var tmp_tmp1_1 *Hx_Obj_loreline_astnode
									_ = tmp_tmp1_1
									tmp_tmp1_1 = _hx_tmp_3
									var _hx_tmp_4 *Hx_Obj_loreline_position
									_ = _hx_tmp_4
									if tmp_tmp1_1 != nil {
										_hx_tmp_4 = tmp_tmp1_1.Hx_Field_pos
									} else {
										_hx_tmp_4 = nil
									}

									pos1 = _hx_tmp_4
								}

								var _hx_tmp_5 *Hx_Obj_loreline_position
								_ = _hx_tmp_5
								if pos1 != nil {
									_hx_tmp_5 = pos1
								} else {
									_hx_tmp_5 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
								}

								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_5)))))
							}
						}

					default:
						{
							var _hx_tmp_6 string
							_ = _hx_tmp_6
							_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_6, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_7 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_7
								if tmp_tmp_1 != nil {
									_hx_tmp_7 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_7 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_7
								var _hx_tmp_8 *Hx_Obj_loreline_position
								_ = _hx_tmp_8
								if tmp_tmp1_1 != nil {
									_hx_tmp_8 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_8 = nil
								}

								pos1 = _hx_tmp_8
							}

							var _hx_tmp_9 *Hx_Obj_loreline_position
							_ = _hx_tmp_9
							if pos1 != nil {
								_hx_tmp_9 = pos1
							} else {
								_hx_tmp_9 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_9)))))
						}
					}
				}

			case 3:
				{
					switch rightType.Hx_Field_enumIndex() {
					case 1, 3:
						{
							var _hx_tmp_10 float64
							_ = _hx_tmp_10
							_hx_tmp_10 = Hx_Field_std_parseFloat(Hx_Field_std_string(left))
							return ((any)((_hx_tmp_10 + Hx_Field_std_parseFloat(Hx_Field_std_string(right)))))
						}

					case 7:
						{
							if rightType.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
								var _hx_tmp_11 string
								_ = _hx_tmp_11
								_hx_tmp_11 = Hx_Field_std_string(left)
								return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_11, Hx_Field_std_string(right))))
							} else {
								var _hx_tmp_12 string
								_ = _hx_tmp_12
								_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
								var tmp string
								_ = tmp
								tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_12, this.VTable.Hx_Field_getTypeName(rightType))
								var tmp1 *Hx_Obj_loreline_position
								_ = tmp1
								tmp1 = pos
								var pos1 *Hx_Obj_loreline_position
								_ = pos1
								if tmp1 != nil {
									pos1 = tmp1
								} else {
									var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
									_ = tmp_tmp_1
									tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
									var _hx_tmp_13 *Hx_Obj_loreline_astnode
									_ = _hx_tmp_13
									if tmp_tmp_1 != nil {
										_hx_tmp_13 = tmp_tmp_1.Hx_Field_node
									} else {
										_hx_tmp_13 = nil
									}

									var tmp_tmp1_1 *Hx_Obj_loreline_astnode
									_ = tmp_tmp1_1
									tmp_tmp1_1 = _hx_tmp_13
									var _hx_tmp_14 *Hx_Obj_loreline_position
									_ = _hx_tmp_14
									if tmp_tmp1_1 != nil {
										_hx_tmp_14 = tmp_tmp1_1.Hx_Field_pos
									} else {
										_hx_tmp_14 = nil
									}

									pos1 = _hx_tmp_14
								}

								var _hx_tmp_15 *Hx_Obj_loreline_position
								_ = _hx_tmp_15
								if pos1 != nil {
									_hx_tmp_15 = pos1
								} else {
									_hx_tmp_15 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
								}

								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_15)))))
							}
						}

					default:
						{
							var _hx_tmp_16 string
							_ = _hx_tmp_16
							_hx_tmp_16 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_16, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_17 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_17
								if tmp_tmp_1 != nil {
									_hx_tmp_17 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_17 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_17
								var _hx_tmp_18 *Hx_Obj_loreline_position
								_ = _hx_tmp_18
								if tmp_tmp1_1 != nil {
									_hx_tmp_18 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_18 = nil
								}

								pos1 = _hx_tmp_18
							}

							var _hx_tmp_19 *Hx_Obj_loreline_position
							_ = _hx_tmp_19
							if pos1 != nil {
								_hx_tmp_19 = pos1
							} else {
								_hx_tmp_19 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_19)))))
						}
					}
				}

			case 7:
				{
					if leftType.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
						var _hx_tmp_20 string
						_ = _hx_tmp_20
						_hx_tmp_20 = Hx_Field_std_string(left)
						return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_20, Hx_Field_std_string(right))))
					} else {
						if rightType.Hx_Field_enumIndex() == 7 {
							if rightType.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
								var _hx_tmp_21 string
								_ = _hx_tmp_21
								_hx_tmp_21 = Hx_Field_std_string(left)
								return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_21, Hx_Field_std_string(right))))
							} else {
								var _hx_tmp_22 string
								_ = _hx_tmp_22
								_hx_tmp_22 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
								var tmp string
								_ = tmp
								tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_22, this.VTable.Hx_Field_getTypeName(rightType))
								var tmp1 *Hx_Obj_loreline_position
								_ = tmp1
								tmp1 = pos
								var pos1 *Hx_Obj_loreline_position
								_ = pos1
								if tmp1 != nil {
									pos1 = tmp1
								} else {
									var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
									_ = tmp_tmp_1
									tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
									var _hx_tmp_23 *Hx_Obj_loreline_astnode
									_ = _hx_tmp_23
									if tmp_tmp_1 != nil {
										_hx_tmp_23 = tmp_tmp_1.Hx_Field_node
									} else {
										_hx_tmp_23 = nil
									}

									var tmp_tmp1_1 *Hx_Obj_loreline_astnode
									_ = tmp_tmp1_1
									tmp_tmp1_1 = _hx_tmp_23
									var _hx_tmp_24 *Hx_Obj_loreline_position
									_ = _hx_tmp_24
									if tmp_tmp1_1 != nil {
										_hx_tmp_24 = tmp_tmp1_1.Hx_Field_pos
									} else {
										_hx_tmp_24 = nil
									}

									pos1 = _hx_tmp_24
								}

								var _hx_tmp_25 *Hx_Obj_loreline_position
								_ = _hx_tmp_25
								if pos1 != nil {
									_hx_tmp_25 = pos1
								} else {
									_hx_tmp_25 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
								}

								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_25)))))
							}
						} else {
							var _hx_tmp_26 string
							_ = _hx_tmp_26
							_hx_tmp_26 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_26, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_27 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_27
								if tmp_tmp_1 != nil {
									_hx_tmp_27 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_27 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_27
								var _hx_tmp_28 *Hx_Obj_loreline_position
								_ = _hx_tmp_28
								if tmp_tmp1_1 != nil {
									_hx_tmp_28 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_28 = nil
								}

								pos1 = _hx_tmp_28
							}

							var _hx_tmp_29 *Hx_Obj_loreline_position
							_ = _hx_tmp_29
							if pos1 != nil {
								_hx_tmp_29 = pos1
							} else {
								_hx_tmp_29 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_29)))))
						}
					}
				}

			default:
				{
					if rightType.Hx_Field_enumIndex() == 7 {
						if rightType.(Hx_Enum_valuetype_TClass).Hx_Field_c == Hx_Obj_string_RTTI {
							var _hx_tmp_30 string
							_ = _hx_tmp_30
							_hx_tmp_30 = Hx_Field_std_string(left)
							return ((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_30, Hx_Field_std_string(right))))
						} else {
							var _hx_tmp_31 string
							_ = _hx_tmp_31
							_hx_tmp_31 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_31, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_32 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_32
								if tmp_tmp_1 != nil {
									_hx_tmp_32 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_32 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_32
								var _hx_tmp_33 *Hx_Obj_loreline_position
								_ = _hx_tmp_33
								if tmp_tmp1_1 != nil {
									_hx_tmp_33 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_33 = nil
								}

								pos1 = _hx_tmp_33
							}

							var _hx_tmp_34 *Hx_Obj_loreline_position
							_ = _hx_tmp_34
							if pos1 != nil {
								_hx_tmp_34 = pos1
							} else {
								_hx_tmp_34 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_34)))))
						}
					} else {
						var _hx_tmp_35 string
						_ = _hx_tmp_35
						_hx_tmp_35 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot add ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
						var tmp string
						_ = tmp
						tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_35, this.VTable.Hx_Field_getTypeName(rightType))
						var tmp1 *Hx_Obj_loreline_position
						_ = tmp1
						tmp1 = pos
						var pos1 *Hx_Obj_loreline_position
						_ = pos1
						if tmp1 != nil {
							pos1 = tmp1
						} else {
							var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
							_ = tmp_tmp_1
							tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
							var _hx_tmp_36 *Hx_Obj_loreline_astnode
							_ = _hx_tmp_36
							if tmp_tmp_1 != nil {
								_hx_tmp_36 = tmp_tmp_1.Hx_Field_node
							} else {
								_hx_tmp_36 = nil
							}

							var tmp_tmp1_1 *Hx_Obj_loreline_astnode
							_ = tmp_tmp1_1
							tmp_tmp1_1 = _hx_tmp_36
							var _hx_tmp_37 *Hx_Obj_loreline_position
							_ = _hx_tmp_37
							if tmp_tmp1_1 != nil {
								_hx_tmp_37 = tmp_tmp1_1.Hx_Field_pos
							} else {
								_hx_tmp_37 = nil
							}

							pos1 = _hx_tmp_37
						}

						var _hx_tmp_38 *Hx_Obj_loreline_position
						_ = _hx_tmp_38
						if pos1 != nil {
							_hx_tmp_38 = pos1
						} else {
							_hx_tmp_38 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
						}

						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_38)))))
					}
				}
			}
		}

	case 28, 29, 30, 31:
		{
			switch leftType.Hx_Field_enumIndex() {
			case 1, 3:
				{
					switch rightType.Hx_Field_enumIndex() {
					case 1, 3:
						{
							var leftNum float64
							_ = leftNum
							leftNum = Hx_Field_std_parseFloat(Hx_Field_std_string(left))
							var rightNum float64
							_ = rightNum
							rightNum = Hx_Field_std_parseFloat(Hx_Field_std_string(right))
							switch op.Hx_Field_enumIndex() {
							case 28:
								{
									return ((any)((leftNum - rightNum)))
								}

							case 29:
								{
									return ((any)((leftNum * rightNum)))
								}

							case 30:
								{
									if rightNum == ((float64)(0)) {
										var tmp *Hx_Obj_loreline_position
										_ = tmp
										tmp = pos
										var pos1 *Hx_Obj_loreline_position
										_ = pos1
										if tmp != nil {
											pos1 = tmp
										} else {
											var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
											_ = tmp_tmp_1
											tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
											var _hx_tmp_39 *Hx_Obj_loreline_astnode
											_ = _hx_tmp_39
											if tmp_tmp_1 != nil {
												_hx_tmp_39 = tmp_tmp_1.Hx_Field_node
											} else {
												_hx_tmp_39 = nil
											}

											var tmp1 *Hx_Obj_loreline_astnode
											_ = tmp1
											tmp1 = _hx_tmp_39
											var _hx_tmp_40 *Hx_Obj_loreline_position
											_ = _hx_tmp_40
											if tmp1 != nil {
												_hx_tmp_40 = tmp1.Hx_Field_pos
											} else {
												_hx_tmp_40 = nil
											}

											pos1 = _hx_tmp_40
										}

										var _hx_tmp_41 *Hx_Obj_loreline_position
										_ = _hx_tmp_41
										if pos1 != nil {
											_hx_tmp_41 = pos1
										} else {
											_hx_tmp_41 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
										}

										panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Division by zero", _hx_tmp_41)))))
									}

									return ((any)(((float64)((leftNum / rightNum)))))
								}

							case 31:
								{
									if rightNum == ((float64)(0)) {
										var tmp *Hx_Obj_loreline_position
										_ = tmp
										tmp = pos
										var pos1 *Hx_Obj_loreline_position
										_ = pos1
										if tmp != nil {
											pos1 = tmp
										} else {
											var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
											_ = tmp_tmp_1
											tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
											var _hx_tmp_42 *Hx_Obj_loreline_astnode
											_ = _hx_tmp_42
											if tmp_tmp_1 != nil {
												_hx_tmp_42 = tmp_tmp_1.Hx_Field_node
											} else {
												_hx_tmp_42 = nil
											}

											var tmp1 *Hx_Obj_loreline_astnode
											_ = tmp1
											tmp1 = _hx_tmp_42
											var _hx_tmp_43 *Hx_Obj_loreline_position
											_ = _hx_tmp_43
											if tmp1 != nil {
												_hx_tmp_43 = tmp1.Hx_Field_pos
											} else {
												_hx_tmp_43 = nil
											}

											pos1 = _hx_tmp_43
										}

										var _hx_tmp_44 *Hx_Obj_loreline_position
										_ = _hx_tmp_44
										if pos1 != nil {
											_hx_tmp_44 = pos1
										} else {
											_hx_tmp_44 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
										}

										panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance("Modulo by zero", _hx_tmp_44)))))
									}

									return math.Mod(leftNum, rightNum)
								}

							default:
								{
									panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
								}
							}
						}

					default:
						{
							var opName string
							_ = opName
							switch op.Hx_Field_enumIndex() {
							case 28:
								{
									opName = "subtract"
								}

							case 29:
								{
									opName = "multiply"
								}

							case 30:
								{
									opName = "divide"
								}

							case 31:
								{
									opName = "modulo"
								}

							default:
								{
									opName = "perform operation on"
								}
							}

							var _hx_tmp_46 string
							_ = _hx_tmp_46
							_hx_tmp_46 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot ", opName), " ")
							var _hx_tmp_45 string
							_ = _hx_tmp_45
							_hx_tmp_45 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_46, this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_45, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_47 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_47
								if tmp_tmp_1 != nil {
									_hx_tmp_47 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_47 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_47
								var _hx_tmp_48 *Hx_Obj_loreline_position
								_ = _hx_tmp_48
								if tmp_tmp1_1 != nil {
									_hx_tmp_48 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_48 = nil
								}

								pos1 = _hx_tmp_48
							}

							var _hx_tmp_49 *Hx_Obj_loreline_position
							_ = _hx_tmp_49
							if pos1 != nil {
								_hx_tmp_49 = pos1
							} else {
								_hx_tmp_49 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_49)))))
						}
					}
				}

			default:
				{
					var opName string
					_ = opName
					switch op.Hx_Field_enumIndex() {
					case 28:
						{
							opName = "subtract"
						}

					case 29:
						{
							opName = "multiply"
						}

					case 30:
						{
							opName = "divide"
						}

					case 31:
						{
							opName = "modulo"
						}

					default:
						{
							opName = "perform operation on"
						}
					}

					var _hx_tmp_51 string
					_ = _hx_tmp_51
					_hx_tmp_51 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot ", opName), " ")
					var _hx_tmp_50 string
					_ = _hx_tmp_50
					_hx_tmp_50 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_51, this.VTable.Hx_Field_getTypeName(leftType)), " and ")
					var tmp string
					_ = tmp
					tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_50, this.VTable.Hx_Field_getTypeName(rightType))
					var tmp1 *Hx_Obj_loreline_position
					_ = tmp1
					tmp1 = pos
					var pos1 *Hx_Obj_loreline_position
					_ = pos1
					if tmp1 != nil {
						pos1 = tmp1
					} else {
						var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
						_ = tmp_tmp_1
						tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
						var _hx_tmp_52 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_52
						if tmp_tmp_1 != nil {
							_hx_tmp_52 = tmp_tmp_1.Hx_Field_node
						} else {
							_hx_tmp_52 = nil
						}

						var tmp_tmp1_1 *Hx_Obj_loreline_astnode
						_ = tmp_tmp1_1
						tmp_tmp1_1 = _hx_tmp_52
						var _hx_tmp_53 *Hx_Obj_loreline_position
						_ = _hx_tmp_53
						if tmp_tmp1_1 != nil {
							_hx_tmp_53 = tmp_tmp1_1.Hx_Field_pos
						} else {
							_hx_tmp_53 = nil
						}

						pos1 = _hx_tmp_53
					}

					var _hx_tmp_54 *Hx_Obj_loreline_position
					_ = _hx_tmp_54
					if pos1 != nil {
						_hx_tmp_54 = pos1
					} else {
						_hx_tmp_54 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
					}

					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_54)))))
				}
			}
		}

	case 32:
		{
			var leftBeat *Hx_Obj_loreline_nbeatdecl
			_ = leftBeat
			leftBeat = Hx_Field_loreline_runtimebeatref_beatOf(left)
			var rightBeat *Hx_Obj_loreline_nbeatdecl
			_ = rightBeat
			rightBeat = Hx_Field_loreline_runtimebeatref_beatOf(right)
			var leftCharacter *Hx_Obj_loreline_runtimecharacterref
			_ = leftCharacter
			leftCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(left)
			var rightCharacter *Hx_Obj_loreline_runtimecharacterref
			_ = rightCharacter
			rightCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(right)
			var result bool
			_ = result
			var _hx_tmp_55 bool
			_ = _hx_tmp_55
			if leftBeat != nil {
				_hx_tmp_55 = Hx_Field_std_isOfType(right, Hx_Obj_string_RTTI)
			} else {
				_hx_tmp_55 = false
			}

			if _hx_tmp_55 {
				var _hx_tmp_56 string
				_ = _hx_tmp_56
				_hx_tmp_56 = leftBeat.Hx_Field_name
				result = (_hx_tmp_56 == Hx_Field_go_haxe_hxdynamic_toString(right))
			} else {
				var _hx_tmp_57 bool
				_ = _hx_tmp_57
				if Hx_Field_std_isOfType(left, Hx_Obj_string_RTTI) {
					_hx_tmp_57 = (rightBeat != nil)
				} else {
					_hx_tmp_57 = false
				}

				if _hx_tmp_57 {
					result = (Hx_Field_go_haxe_hxdynamic_toString(left) == rightBeat.Hx_Field_name)
				} else {
					if (leftBeat != nil) && (rightBeat != nil) {
						result = (leftBeat.Hx_Field_name == rightBeat.Hx_Field_name)
					} else {
						var _hx_tmp_58 bool
						_ = _hx_tmp_58
						if leftCharacter != nil {
							_hx_tmp_58 = Hx_Field_std_isOfType(right, Hx_Obj_string_RTTI)
						} else {
							_hx_tmp_58 = false
						}

						if _hx_tmp_58 {
							var _hx_tmp_59 string
							_ = _hx_tmp_59
							_hx_tmp_59 = leftCharacter.Hx_Field_name
							result = (_hx_tmp_59 == Hx_Field_go_haxe_hxdynamic_toString(right))
						} else {
							var _hx_tmp_60 bool
							_ = _hx_tmp_60
							if Hx_Field_std_isOfType(left, Hx_Obj_string_RTTI) {
								_hx_tmp_60 = (rightCharacter != nil)
							} else {
								_hx_tmp_60 = false
							}

							if _hx_tmp_60 {
								result = (Hx_Field_go_haxe_hxdynamic_toString(left) == rightCharacter.Hx_Field_name)
							} else {
								if (leftCharacter != nil) || (rightCharacter != nil) {
									var _hx_tmp_61 any
									_ = _hx_tmp_61
									if leftCharacter != nil {
										_hx_tmp_61 = leftCharacter.VTable.Hx_Field_get_fields()
									} else {
										_hx_tmp_61 = left
									}

									var l any
									_ = l
									l = _hx_tmp_61
									var _hx_tmp_62 any
									_ = _hx_tmp_62
									if rightCharacter != nil {
										_hx_tmp_62 = rightCharacter.VTable.Hx_Field_get_fields()
									} else {
										_hx_tmp_62 = right
									}

									var r any
									_ = r
									r = _hx_tmp_62
									result = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, r))))
								} else {
									result = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(left, right))))
								}
							}
						}
					}
				}
			}

			switch op.Hx_Field_enumIndex() {
			case 32:
				{
					return ((any)(result))
				}

			case 33:
				{
					return ((any)(!result))
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
				}
			}
		}

	case 33:
		{
			var leftBeat *Hx_Obj_loreline_nbeatdecl
			_ = leftBeat
			leftBeat = Hx_Field_loreline_runtimebeatref_beatOf(left)
			var rightBeat *Hx_Obj_loreline_nbeatdecl
			_ = rightBeat
			rightBeat = Hx_Field_loreline_runtimebeatref_beatOf(right)
			var leftCharacter *Hx_Obj_loreline_runtimecharacterref
			_ = leftCharacter
			leftCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(left)
			var rightCharacter *Hx_Obj_loreline_runtimecharacterref
			_ = rightCharacter
			rightCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(right)
			var result bool
			_ = result
			var _hx_tmp_63 bool
			_ = _hx_tmp_63
			if leftBeat != nil {
				_hx_tmp_63 = Hx_Field_std_isOfType(right, Hx_Obj_string_RTTI)
			} else {
				_hx_tmp_63 = false
			}

			if _hx_tmp_63 {
				var _hx_tmp_64 string
				_ = _hx_tmp_64
				_hx_tmp_64 = leftBeat.Hx_Field_name
				result = (_hx_tmp_64 == Hx_Field_go_haxe_hxdynamic_toString(right))
			} else {
				var _hx_tmp_65 bool
				_ = _hx_tmp_65
				if Hx_Field_std_isOfType(left, Hx_Obj_string_RTTI) {
					_hx_tmp_65 = (rightBeat != nil)
				} else {
					_hx_tmp_65 = false
				}

				if _hx_tmp_65 {
					result = (Hx_Field_go_haxe_hxdynamic_toString(left) == rightBeat.Hx_Field_name)
				} else {
					if (leftBeat != nil) && (rightBeat != nil) {
						result = (leftBeat.Hx_Field_name == rightBeat.Hx_Field_name)
					} else {
						var _hx_tmp_66 bool
						_ = _hx_tmp_66
						if leftCharacter != nil {
							_hx_tmp_66 = Hx_Field_std_isOfType(right, Hx_Obj_string_RTTI)
						} else {
							_hx_tmp_66 = false
						}

						if _hx_tmp_66 {
							var _hx_tmp_67 string
							_ = _hx_tmp_67
							_hx_tmp_67 = leftCharacter.Hx_Field_name
							result = (_hx_tmp_67 == Hx_Field_go_haxe_hxdynamic_toString(right))
						} else {
							var _hx_tmp_68 bool
							_ = _hx_tmp_68
							if Hx_Field_std_isOfType(left, Hx_Obj_string_RTTI) {
								_hx_tmp_68 = (rightCharacter != nil)
							} else {
								_hx_tmp_68 = false
							}

							if _hx_tmp_68 {
								result = (Hx_Field_go_haxe_hxdynamic_toString(left) == rightCharacter.Hx_Field_name)
							} else {
								if (leftCharacter != nil) || (rightCharacter != nil) {
									var _hx_tmp_69 any
									_ = _hx_tmp_69
									if leftCharacter != nil {
										_hx_tmp_69 = leftCharacter.VTable.Hx_Field_get_fields()
									} else {
										_hx_tmp_69 = left
									}

									var l any
									_ = l
									l = _hx_tmp_69
									var _hx_tmp_70 any
									_ = _hx_tmp_70
									if rightCharacter != nil {
										_hx_tmp_70 = rightCharacter.VTable.Hx_Field_get_fields()
									} else {
										_hx_tmp_70 = right
									}

									var r any
									_ = r
									r = _hx_tmp_70
									result = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(l, r))))
								} else {
									result = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(left, right))))
								}
							}
						}
					}
				}
			}

			switch op.Hx_Field_enumIndex() {
			case 32:
				{
					return ((any)(result))
				}

			case 33:
				{
					return ((any)(!result))
				}

			default:
				{
					panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
				}
			}
		}

	case 34, 35, 36, 37:
		{
			switch leftType.Hx_Field_enumIndex() {
			case 1, 3:
				{
					switch rightType.Hx_Field_enumIndex() {
					case 1, 3:
						{
							var leftNum float64
							_ = leftNum
							leftNum = Hx_Field_std_parseFloat(Hx_Field_std_string(left))
							var rightNum float64
							_ = rightNum
							rightNum = Hx_Field_std_parseFloat(Hx_Field_std_string(right))
							switch op.Hx_Field_enumIndex() {
							case 34:
								{
									return ((any)((leftNum > rightNum)))
								}

							case 35:
								{
									return ((any)((leftNum < rightNum)))
								}

							case 36:
								{
									return ((any)((leftNum >= rightNum)))
								}

							case 37:
								{
									return ((any)((leftNum <= rightNum)))
								}

							default:
								{
									panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
								}
							}
						}

					default:
						{
							var _hx_tmp_71 string
							_ = _hx_tmp_71
							_hx_tmp_71 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot compare ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
							var tmp string
							_ = tmp
							tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_71, this.VTable.Hx_Field_getTypeName(rightType))
							var tmp1 *Hx_Obj_loreline_position
							_ = tmp1
							tmp1 = pos
							var pos1 *Hx_Obj_loreline_position
							_ = pos1
							if tmp1 != nil {
								pos1 = tmp1
							} else {
								var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
								_ = tmp_tmp_1
								tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
								var _hx_tmp_72 *Hx_Obj_loreline_astnode
								_ = _hx_tmp_72
								if tmp_tmp_1 != nil {
									_hx_tmp_72 = tmp_tmp_1.Hx_Field_node
								} else {
									_hx_tmp_72 = nil
								}

								var tmp_tmp1_1 *Hx_Obj_loreline_astnode
								_ = tmp_tmp1_1
								tmp_tmp1_1 = _hx_tmp_72
								var _hx_tmp_73 *Hx_Obj_loreline_position
								_ = _hx_tmp_73
								if tmp_tmp1_1 != nil {
									_hx_tmp_73 = tmp_tmp1_1.Hx_Field_pos
								} else {
									_hx_tmp_73 = nil
								}

								pos1 = _hx_tmp_73
							}

							var _hx_tmp_74 *Hx_Obj_loreline_position
							_ = _hx_tmp_74
							if pos1 != nil {
								_hx_tmp_74 = pos1
							} else {
								_hx_tmp_74 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_74)))))
						}
					}
				}

			default:
				{
					var _hx_tmp_75 string
					_ = _hx_tmp_75
					_hx_tmp_75 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot compare ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
					var tmp string
					_ = tmp
					tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_75, this.VTable.Hx_Field_getTypeName(rightType))
					var tmp1 *Hx_Obj_loreline_position
					_ = tmp1
					tmp1 = pos
					var pos1 *Hx_Obj_loreline_position
					_ = pos1
					if tmp1 != nil {
						pos1 = tmp1
					} else {
						var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
						_ = tmp_tmp_1
						tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
						var _hx_tmp_76 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_76
						if tmp_tmp_1 != nil {
							_hx_tmp_76 = tmp_tmp_1.Hx_Field_node
						} else {
							_hx_tmp_76 = nil
						}

						var tmp_tmp1_1 *Hx_Obj_loreline_astnode
						_ = tmp_tmp1_1
						tmp_tmp1_1 = _hx_tmp_76
						var _hx_tmp_77 *Hx_Obj_loreline_position
						_ = _hx_tmp_77
						if tmp_tmp1_1 != nil {
							_hx_tmp_77 = tmp_tmp1_1.Hx_Field_pos
						} else {
							_hx_tmp_77 = nil
						}

						pos1 = _hx_tmp_77
					}

					var _hx_tmp_78 *Hx_Obj_loreline_position
					_ = _hx_tmp_78
					if pos1 != nil {
						_hx_tmp_78 = pos1
					} else {
						_hx_tmp_78 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
					}

					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_78)))))
				}
			}
		}

	case 38:
		{
			if leftType.Hx_Field_enumIndex() == 4 {
				if rightType.Hx_Field_enumIndex() == 4 {
					switch op.Hx_Field_enumIndex() {
					case 38:
						{
							var _hx_tmp_79 bool
							_ = _hx_tmp_79
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(left)) {
								_hx_tmp_79 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(right)))
							} else {
								_hx_tmp_79 = false
							}

							return ((any)(_hx_tmp_79))
						}

					case 39:
						{
							var _hx_tmp_80 bool
							_ = _hx_tmp_80
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(left)) {
								_hx_tmp_80 = true
							} else {
								_hx_tmp_80 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(right)))
							}

							return ((any)(_hx_tmp_80))
						}

					default:
						{
							panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
						}
					}
				} else {
					var _hx_tmp_81 string
					_ = _hx_tmp_81
					_hx_tmp_81 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot perform logical operation on ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
					var tmp string
					_ = tmp
					tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_81, this.VTable.Hx_Field_getTypeName(rightType))
					var tmp1 *Hx_Obj_loreline_position
					_ = tmp1
					tmp1 = pos
					var pos1 *Hx_Obj_loreline_position
					_ = pos1
					if tmp1 != nil {
						pos1 = tmp1
					} else {
						var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
						_ = tmp_tmp_1
						tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
						var _hx_tmp_82 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_82
						if tmp_tmp_1 != nil {
							_hx_tmp_82 = tmp_tmp_1.Hx_Field_node
						} else {
							_hx_tmp_82 = nil
						}

						var tmp_tmp1_1 *Hx_Obj_loreline_astnode
						_ = tmp_tmp1_1
						tmp_tmp1_1 = _hx_tmp_82
						var _hx_tmp_83 *Hx_Obj_loreline_position
						_ = _hx_tmp_83
						if tmp_tmp1_1 != nil {
							_hx_tmp_83 = tmp_tmp1_1.Hx_Field_pos
						} else {
							_hx_tmp_83 = nil
						}

						pos1 = _hx_tmp_83
					}

					var _hx_tmp_84 *Hx_Obj_loreline_position
					_ = _hx_tmp_84
					if pos1 != nil {
						_hx_tmp_84 = pos1
					} else {
						_hx_tmp_84 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
					}

					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_84)))))
				}
			} else {
				var _hx_tmp_85 string
				_ = _hx_tmp_85
				_hx_tmp_85 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot perform logical operation on ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
				var tmp string
				_ = tmp
				tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_85, this.VTable.Hx_Field_getTypeName(rightType))
				var tmp1 *Hx_Obj_loreline_position
				_ = tmp1
				tmp1 = pos
				var pos1 *Hx_Obj_loreline_position
				_ = pos1
				if tmp1 != nil {
					pos1 = tmp1
				} else {
					var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
					_ = tmp_tmp_1
					tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
					var _hx_tmp_86 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_86
					if tmp_tmp_1 != nil {
						_hx_tmp_86 = tmp_tmp_1.Hx_Field_node
					} else {
						_hx_tmp_86 = nil
					}

					var tmp_tmp1_1 *Hx_Obj_loreline_astnode
					_ = tmp_tmp1_1
					tmp_tmp1_1 = _hx_tmp_86
					var _hx_tmp_87 *Hx_Obj_loreline_position
					_ = _hx_tmp_87
					if tmp_tmp1_1 != nil {
						_hx_tmp_87 = tmp_tmp1_1.Hx_Field_pos
					} else {
						_hx_tmp_87 = nil
					}

					pos1 = _hx_tmp_87
				}

				var _hx_tmp_88 *Hx_Obj_loreline_position
				_ = _hx_tmp_88
				if pos1 != nil {
					_hx_tmp_88 = pos1
				} else {
					_hx_tmp_88 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
				}

				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_88)))))
			}
		}

	case 39:
		{
			if leftType.Hx_Field_enumIndex() == 4 {
				if rightType.Hx_Field_enumIndex() == 4 {
					switch op.Hx_Field_enumIndex() {
					case 38:
						{
							var _hx_tmp_89 bool
							_ = _hx_tmp_89
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(left)) {
								_hx_tmp_89 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(right)))
							} else {
								_hx_tmp_89 = false
							}

							return ((any)(_hx_tmp_89))
						}

					case 39:
						{
							var _hx_tmp_90 bool
							_ = _hx_tmp_90
							if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(left)) {
								_hx_tmp_90 = true
							} else {
								_hx_tmp_90 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(right)))
							}

							return ((any)(_hx_tmp_90))
						}

					default:
						{
							panic(Hx_Field_haxe_exception_thrown(((any)("Unreachable"))))
						}
					}
				} else {
					var _hx_tmp_91 string
					_ = _hx_tmp_91
					_hx_tmp_91 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot perform logical operation on ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
					var tmp string
					_ = tmp
					tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_91, this.VTable.Hx_Field_getTypeName(rightType))
					var tmp1 *Hx_Obj_loreline_position
					_ = tmp1
					tmp1 = pos
					var pos1 *Hx_Obj_loreline_position
					_ = pos1
					if tmp1 != nil {
						pos1 = tmp1
					} else {
						var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
						_ = tmp_tmp_1
						tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
						var _hx_tmp_92 *Hx_Obj_loreline_astnode
						_ = _hx_tmp_92
						if tmp_tmp_1 != nil {
							_hx_tmp_92 = tmp_tmp_1.Hx_Field_node
						} else {
							_hx_tmp_92 = nil
						}

						var tmp_tmp1_1 *Hx_Obj_loreline_astnode
						_ = tmp_tmp1_1
						tmp_tmp1_1 = _hx_tmp_92
						var _hx_tmp_93 *Hx_Obj_loreline_position
						_ = _hx_tmp_93
						if tmp_tmp1_1 != nil {
							_hx_tmp_93 = tmp_tmp1_1.Hx_Field_pos
						} else {
							_hx_tmp_93 = nil
						}

						pos1 = _hx_tmp_93
					}

					var _hx_tmp_94 *Hx_Obj_loreline_position
					_ = _hx_tmp_94
					if pos1 != nil {
						_hx_tmp_94 = pos1
					} else {
						_hx_tmp_94 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
					}

					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_94)))))
				}
			} else {
				var _hx_tmp_95 string
				_ = _hx_tmp_95
				_hx_tmp_95 = Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot perform logical operation on ", this.VTable.Hx_Field_getTypeName(leftType)), " and ")
				var tmp string
				_ = tmp
				tmp = Hx_Field_go_haxe_hxdynamic_stringConcat(_hx_tmp_95, this.VTable.Hx_Field_getTypeName(rightType))
				var tmp1 *Hx_Obj_loreline_position
				_ = tmp1
				tmp1 = pos
				var pos1 *Hx_Obj_loreline_position
				_ = pos1
				if tmp1 != nil {
					pos1 = tmp1
				} else {
					var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
					_ = tmp_tmp_1
					tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
					var _hx_tmp_96 *Hx_Obj_loreline_astnode
					_ = _hx_tmp_96
					if tmp_tmp_1 != nil {
						_hx_tmp_96 = tmp_tmp_1.Hx_Field_node
					} else {
						_hx_tmp_96 = nil
					}

					var tmp_tmp1_1 *Hx_Obj_loreline_astnode
					_ = tmp_tmp1_1
					tmp_tmp1_1 = _hx_tmp_96
					var _hx_tmp_97 *Hx_Obj_loreline_position
					_ = _hx_tmp_97
					if tmp_tmp1_1 != nil {
						_hx_tmp_97 = tmp_tmp1_1.Hx_Field_pos
					} else {
						_hx_tmp_97 = nil
					}

					pos1 = _hx_tmp_97
				}

				var _hx_tmp_98 *Hx_Obj_loreline_position
				_ = _hx_tmp_98
				if pos1 != nil {
					_hx_tmp_98 = pos1
				} else {
					_hx_tmp_98 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
				}

				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_98)))))
			}
		}

	default:
		{
			var tmp string
			_ = tmp
			tmp = Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid operation: ", Hx_Field_std_string(op))
			var tmp1 *Hx_Obj_loreline_position
			_ = tmp1
			tmp1 = pos
			var pos1 *Hx_Obj_loreline_position
			_ = pos1
			if tmp1 != nil {
				pos1 = tmp1
			} else {
				var tmp_tmp_1 *Hx_Obj_loreline_runtimescope
				_ = tmp_tmp_1
				tmp_tmp_1 = this.VTable.Hx_Field_get_currentScope()
				var _hx_tmp_99 *Hx_Obj_loreline_astnode
				_ = _hx_tmp_99
				if tmp_tmp_1 != nil {
					_hx_tmp_99 = tmp_tmp_1.Hx_Field_node
				} else {
					_hx_tmp_99 = nil
				}

				var tmp_tmp1_1 *Hx_Obj_loreline_astnode
				_ = tmp_tmp1_1
				tmp_tmp1_1 = _hx_tmp_99
				var _hx_tmp_100 *Hx_Obj_loreline_position
				_ = _hx_tmp_100
				if tmp_tmp1_1 != nil {
					_hx_tmp_100 = tmp_tmp1_1.Hx_Field_pos
				} else {
					_hx_tmp_100 = nil
				}

				pos1 = _hx_tmp_100
			}

			var _hx_tmp_101 *Hx_Obj_loreline_position
			_ = _hx_tmp_101
			if pos1 != nil {
				_hx_tmp_101 = pos1
			} else {
				_hx_tmp_101 = this.Hx_Field_context.Hx_Field_script.Hx_Field_pos
			}

			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_runtimeerror_CreateInstance(tmp, _hx_tmp_101)))))
		}
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_valueToString(value any) string {
	return this.VTable.Hx_Field_valueToStringImpl(value, nil)
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_valueToStringImpl(value any, seen HxArray[any]) string {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return "null"
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		return Hx_Field_go_haxe_hxdynamic_toString(value)
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI) {
		_hx_tmp_1 = true
	} else {
		_hx_tmp_1 = Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI)
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI)
	}

	if _hx_tmp_0 {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
			return "null"
		} else {
			return Hx_Field_std_string(value)
		}
	}

	if seen == nil {
		seen = HxMakeArray[any]().Dyn()
	}

	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(seen)) {
				break
			}

			var s any
			_ = s
			s = seen.Get(_g)
			_g++
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(s, value))) {
				return "..."
			}
		}
	}

	_ = (Hx_Array_Push[any](HxMakeArrayView[any](seen), value))
	if Hx_Field_loreline_arrays_isArray(value) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(value)
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		func() func(x string) {
			var _hx_closure_src_677 func(x any)
			_ = _hx_closure_src_677
			_hx_closure_src_677 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_677(((any)(x)))
			}
		}()("[")
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_3 int32
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_3
				if i > 0 {
					func() func(x string) {
						var _hx_closure_src_678 func(x any)
						_ = _hx_closure_src_678
						_hx_closure_src_678 = buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_678(((any)(x)))
						}
					}()(", ")
				}

				func() func(x string) {
					var _hx_closure_src_679 func(x any)
					_ = _hx_closure_src_679
					_hx_closure_src_679 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_679(((any)(x)))
					}
				}()(this.VTable.Hx_Field_valueToStringImpl(Hx_Field_loreline_arrays_arrayGet(value, i), seen))
			}
		}

		func() func(x string) {
			var _hx_closure_src_680 func(x any)
			_ = _hx_closure_src_680
			_hx_closure_src_680 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_680(((any)(x)))
			}
		}()("]")
		_ = (Hx_Array_Pop[any](HxMakeArrayView[any](seen)).Value)
		return buf.VTable.Hx_Field_toString()
	}

	var asBeat *Hx_Obj_loreline_nbeatdecl
	_ = asBeat
	asBeat = Hx_Field_loreline_runtimebeatref_beatOf(value)
	if asBeat != nil {
		_ = (Hx_Array_Pop[any](HxMakeArrayView[any](seen)).Value)
		return asBeat.Hx_Field_name
	}

	var asCharacter *Hx_Obj_loreline_runtimecharacterref
	_ = asCharacter
	asCharacter = Hx_Field_loreline_runtimecharacterref_characterOf(value)
	if asCharacter != nil {
		var _hx_tmp_4 *Hx_Obj_loreline_interpreter
		_ = _hx_tmp_4
		_hx_tmp_4 = this
		var nameValue any
		_ = nameValue
		nameValue = Hx_Field_loreline_objects_getField(_hx_tmp_4, asCharacter.VTable.Hx_Field_get_fields(), "name")
		_ = (Hx_Array_Pop[any](HxMakeArrayView[any](seen)).Value)
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(nameValue, nil))) {
			return this.VTable.Hx_Field_valueToStringImpl(nameValue, nil)
		} else {
			return asCharacter.Hx_Field_name
		}
	}

	if Hx_Field_loreline_objects_isFields(value) {
		var keys HxArray[string]
		_ = keys
		keys = Hx_Field_loreline_objects_getFields(this, value)
		var _hx_tmp_5 HxArray[string]
		_ = _hx_tmp_5
		_hx_tmp_5 = keys
		(Hx_Array_Sort[string](_hx_tmp_5, func(a string, b string) int32 {
			if a < b {
				return -1
			} else {
				if a > b {
					return 1
				} else {
					return 0
				}
			}
		}))
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		func() func(x string) {
			var _hx_closure_src_681 func(x any)
			_ = _hx_closure_src_681
			_hx_closure_src_681 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_681(((any)(x)))
			}
		}()("{")
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = keys.Len()
			for _g < _g1 {
				var _hx_tmp_6 int32
				_ = _hx_tmp_6
				_hx_tmp_6 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_6
				if i > 0 {
					func() func(x string) {
						var _hx_closure_src_682 func(x any)
						_ = _hx_closure_src_682
						_hx_closure_src_682 = buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_682(((any)(x)))
						}
					}()(", ")
				}

				func() func(x string) {
					var _hx_closure_src_683 func(x any)
					_ = _hx_closure_src_683
					_hx_closure_src_683 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_683(((any)(x)))
					}
				}()(keys.Get(i))
				func() func(x string) {
					var _hx_closure_src_684 func(x any)
					_ = _hx_closure_src_684
					_hx_closure_src_684 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_684(((any)(x)))
					}
				}()(": ")
				func() func(x string) {
					var _hx_closure_src_685 func(x any)
					_ = _hx_closure_src_685
					_hx_closure_src_685 = buf.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_685(((any)(x)))
					}
				}()(this.VTable.Hx_Field_valueToStringImpl(Hx_Field_loreline_objects_getField(this, value, keys.Get(i)), seen))
			}
		}

		func() func(x string) {
			var _hx_closure_src_686 func(x any)
			_ = _hx_closure_src_686
			_hx_closure_src_686 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_686(((any)(x)))
			}
		}()("}")
		_ = (Hx_Array_Pop[any](HxMakeArrayView[any](seen)).Value)
		return buf.VTable.Hx_Field_toString()
	}

	_ = (Hx_Array_Pop[any](HxMakeArrayView[any](seen)).Value)
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return "null"
	} else {
		return Hx_Field_std_string(value)
	}
}

func (this *Hx_Obj_loreline_interpreter) Hx_Field_printLoreline(node *Hx_Obj_loreline_node) string {
	{
		var hx_try_return string
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
							hx_try_return = HxStringNull
							return
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var printer *Hx_Obj_loreline_printer
				_ = printer
				printer = Hx_Obj_loreline_printer_CreateInstance(nil, nil)
				printer.Hx_Field_enableComments = false
				hx_try_state = 1
				hx_try_return = Hx_Field_stringtools_trim(printer.VTable.Hx_Field_print(node))
				return
			}
		}()

		return hx_try_return
	}
}

var Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH int32

func Hx_Init_Hx_Field_loreline_interpreter_HOST_VALUE_MAX_DEPTH() int32 {
	return 64
}

func Hx_Field_loreline_interpreter_hasEligibleCandidate(candidates HxArray[any]) bool {
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(candidates)) {
				break
			}

			var candidate any
			_ = candidate
			candidate = candidates.Get(_g)
			_g++
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_getField(candidate, "eligible")))) {
				return true
			}
		}
	}

	return false
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_choiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_interpreter]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatdecl]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatparam]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nchoiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstatedecl]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_runtimescope]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_texttag]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[bool]()
	hxRegisterArrayConstructor[func()]()
	hxRegisterArrayConstructor[func(str *Hx_Obj_loreline_nstringliteral) *Hx_Obj_loreline_nstringliteral]()
	hxRegisterArrayConstructor[int32]()
	hxRegisterArrayConstructor[string]()
}
