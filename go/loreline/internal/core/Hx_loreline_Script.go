package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_script() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_script() {
}

type Hx_Obj_VTable_loreline_script interface {
	Hx_Field_type() string
	Hx_Field_eachExcludingImported(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_iterator() *Hx_Obj_loreline__script_scriptbodyiterator
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_script struct {
	Hx_Obj_loreline_astnode
	VTable              Hx_Obj_VTable_loreline_script `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_indentSize int32
	Hx_Field_filePath   string
	Hx_Field_body       HxArray[*Hx_Obj_loreline_astnode]
}

func Hx_Obj_loreline_script_CreateEmptyInstance() *Hx_Obj_loreline_script {
	obj := &Hx_Obj_loreline_script{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_indentSize = HxDefault[int32]()
	obj.Hx_Field_filePath = HxDefault[string]()
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	return obj
}

func Hx_Obj_loreline_script_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, body HxArray[*Hx_Obj_loreline_astnode]) *Hx_Obj_loreline_script {
	obj := Hx_Obj_loreline_script_CreateEmptyInstance()
	obj.Hx_New(id, pos, body)
	return obj
}

func (this *Hx_Obj_loreline_script) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, body HxArray[*Hx_Obj_loreline_astnode]) {
	{
		this.Hx_Field_indentSize = 2
	}
	{
		this.Hx_Field_filePath = HxStringNull
	}
	{
		var _hx_tmp_0 Hx_Typedef_haxe__int64___int64
		_ = _hx_tmp_0
		_hx_tmp_0 = id
		var _hx_tmp_1 *Hx_Obj_loreline_position
		_ = _hx_tmp_1
		_hx_tmp_1 = pos
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = nil
		this.Hx_Obj_loreline_astnode.Hx_New(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, nil)
		this.Hx_Field_body = body
	}

}

var Hx_Obj_loreline_script_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Script",
	HxMakeArray[string]("fromJson"),
	HxMakeArray[string]("indentSize", "filePath", "body", "toJson", "type", "eachExcludingImported", "each", "iterator", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_script_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(2)))
	},
	func() any {
		return Hx_Obj_loreline_script_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_script_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "fromJson":
			return Hx_Field_loreline_script_fromJson
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_script) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_script_RTTI
}

func (this *Hx_Obj_loreline_script) Hx_As_Hx_Obj_loreline_script() *Hx_Obj_loreline_script {
	return this
}

func (this *Hx_Obj_loreline_script) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g2
		_g2 = this.Hx_Field_body
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var decl *Hx_Obj_loreline_astnode
			_ = decl
			decl = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_1, decl.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "body", _g)
	Hx_Field_go_haxe_hxdynamic_setField(json, "indentSize", this.Hx_Field_indentSize)
	if this.Hx_Field_filePath != HxStringNull {
		Hx_Field_go_haxe_hxdynamic_setField(json, "filePath", this.Hx_Field_filePath)
	}

	return json
}

func (this *Hx_Obj_loreline_script) Hx_Field_type() string {
	return "Script"
}

func (this *Hx_Obj_loreline_script) Hx_Field_eachExcludingImported(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g1
		_g1 = this.Hx_Field_body
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var child *Hx_Obj_loreline_astnode
			_ = child
			child = _g1.Get(_g)
			_g++
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := child
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_1, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			if !Hx_Field_std_isOfType(child, Hx_Obj_loreline_nimportstatement_RTTI) {
				child.VTable.Hx_Field_each(handleNode)
			}
		}
	}
}

func (this *Hx_Obj_loreline_script) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g1
		_g1 = this.Hx_Field_body
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var child *Hx_Obj_loreline_astnode
			_ = child
			child = _g1.Get(_g)
			_g++
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := child
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_1, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			child.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_script) Hx_Field_iterator() *Hx_Obj_loreline__script_scriptbodyiterator {
	return Hx_Obj_loreline__script_scriptbodyiterator_CreateInstance(this.Hx_Field_body)
}

func Hx_Field_loreline_script_fromJson(json any) *Hx_Obj_loreline_script {
	return Hx_Field_loreline_jsontoast_scriptFromJson(json)
}

func Hx_ClassInit_Hx_Obj_VTable_loreline__script_scriptbodyiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__script_scriptbodyiterator() {
}

type Hx_Obj_VTable_loreline__script_scriptbodyiterator interface {
	Hx_Field_fillBody(body HxArray[*Hx_Obj_loreline_astnode])
	Hx_Field_hasNext() bool
	Hx_Field_next() *Hx_Obj_loreline_astnode
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__script_scriptbodyiterator struct {
	VTable            Hx_Obj_VTable_loreline__script_scriptbodyiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_body     HxArray[*Hx_Obj_loreline_astnode]
	Hx_Field_index    int32
	Hx_Field_flatBody HxArray[*Hx_Obj_loreline_astnode]
}

func Hx_Obj_loreline__script_scriptbodyiterator_CreateEmptyInstance() *Hx_Obj_loreline__script_scriptbodyiterator {
	obj := &Hx_Obj_loreline__script_scriptbodyiterator{}
	obj.VTable = obj
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	obj.Hx_Field_index = HxDefault[int32]()
	obj.Hx_Field_flatBody = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	return obj
}

func Hx_Obj_loreline__script_scriptbodyiterator_CreateInstance(body HxArray[*Hx_Obj_loreline_astnode]) *Hx_Obj_loreline__script_scriptbodyiterator {
	obj := Hx_Obj_loreline__script_scriptbodyiterator_CreateEmptyInstance()
	obj.Hx_New(body)
	return obj
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_New(body HxArray[*Hx_Obj_loreline_astnode]) {
	{
		this.Hx_Field_body = body
		this.Hx_Field_index = 0
		this.Hx_Field_flatBody = HxMakeArray[*Hx_Obj_loreline_astnode]()
		this.VTable.Hx_Field_fillBody(body)
	}

}

var Hx_Obj_loreline__script_scriptbodyiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Script.ScriptBodyIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("body", "index", "flatBody", "fillBody", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__script_scriptbodyiterator_CreateInstance(HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__script_scriptbodyiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__script_scriptbodyiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__script_scriptbodyiterator_RTTI
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_As_Hx_Obj_loreline__script_scriptbodyiterator() *Hx_Obj_loreline__script_scriptbodyiterator {
	return this
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_Field_fillBody(body HxArray[*Hx_Obj_loreline_astnode]) {
	var _g int32
	_ = _g
	_g = 0
	var _g1 int32
	_ = _g1
	_g1 = body.Len()
	for _g < _g1 {
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = _g
		_g = (_g + 1)
		var i int32
		_ = i
		i = _hx_tmp_0
		var node *Hx_Obj_loreline_astnode
		_ = node
		node = body.Get(i)
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nimportstatement_RTTI) {
			var importNode *Hx_Obj_loreline_nimportstatement
			_ = importNode
			importNode = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement
			}).Hx_As_Hx_Obj_loreline_nimportstatement()
			if importNode.Hx_Field_script != nil {
				this.VTable.Hx_Field_fillBody(importNode.Hx_Field_script.Hx_Field_body)
			}
		}

		_ = (Hx_Array_Push[*Hx_Obj_loreline_astnode](this.Hx_Field_flatBody, body.Get(i)))
	}
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < this.Hx_Field_flatBody.Len())
}

func (this *Hx_Obj_loreline__script_scriptbodyiterator) Hx_Field_next() *Hx_Obj_loreline_astnode {
	var v *Hx_Obj_loreline_astnode
	_ = v
	v = this.Hx_Field_flatBody.Get(this.Hx_Field_index)
	this.Hx_Field_index++
	return v
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[any]()
}
