package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_gonode() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_gonode() {
}

type Hx_Obj_VTable_loreline_lib_gonode interface {
	Hx_Field_length() int32
	Hx_Field_id() string
	Hx_Field_toJson(pretty bool) string
	Hx_Field_type() string
	Hx_Field_line() int32
	Hx_Field_column() int32
	Hx_Field_offset() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_gonode struct {
	VTable        Hx_Obj_VTable_loreline_lib_gonode `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_node *Hx_Obj_loreline_node
}

func Hx_Obj_loreline_lib_gonode_CreateEmptyInstance() *Hx_Obj_loreline_lib_gonode {
	obj := &Hx_Obj_loreline_lib_gonode{}
	obj.VTable = obj
	obj.Hx_Field_node = HxDefault[*Hx_Obj_loreline_node]()
	return obj
}

func Hx_Obj_loreline_lib_gonode_CreateInstance(node *Hx_Obj_loreline_node) *Hx_Obj_loreline_lib_gonode {
	obj := Hx_Obj_loreline_lib_gonode_CreateEmptyInstance()
	obj.Hx_New(node)
	return obj
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_New(node *Hx_Obj_loreline_node) {
	{
		this.Hx_Field_node = node
	}

}

var Hx_Obj_loreline_lib_gonode_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoNode",
	HxMakeArray[string](),
	HxMakeArray[string]("node", "type", "line", "column", "offset", "length", "id", "toJson"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_gonode_CreateInstance(HxConvert[*Hx_Obj_loreline_node](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lib_gonode_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_gonode_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_gonode_RTTI
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_As_Hx_Obj_loreline_lib_gonode() *Hx_Obj_loreline_lib_gonode {
	return this
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_type() string {
	return this.Hx_Field_node.VTable.Hx_Field_type()
}

func (this *Hx_Obj_loreline_lib_gonode) Type() string {
	return this.VTable.Hx_Field_type()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_line() int32 {
	if this.Hx_Field_node.Hx_Field_pos != nil {
		return this.Hx_Field_node.Hx_Field_pos.Hx_Field_line
	} else {
		return 0
	}
}

func (this *Hx_Obj_loreline_lib_gonode) Line() int32 {
	return this.VTable.Hx_Field_line()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_column() int32 {
	if this.Hx_Field_node.Hx_Field_pos != nil {
		return this.Hx_Field_node.Hx_Field_pos.Hx_Field_column
	} else {
		return 0
	}
}

func (this *Hx_Obj_loreline_lib_gonode) Column() int32 {
	return this.VTable.Hx_Field_column()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_offset() int32 {
	if this.Hx_Field_node.Hx_Field_pos != nil {
		return this.Hx_Field_node.Hx_Field_pos.Hx_Field_offset
	} else {
		return 0
	}
}

func (this *Hx_Obj_loreline_lib_gonode) Offset() int32 {
	return this.VTable.Hx_Field_offset()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_length() int32 {
	if this.Hx_Field_node.Hx_Field_pos != nil {
		return this.Hx_Field_node.Hx_Field_pos.Hx_Field_length
	} else {
		return 0
	}
}

func (this *Hx_Obj_loreline_lib_gonode) Length() int32 {
	return this.VTable.Hx_Field_length()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_id() string {
	return Hx_Field_loreline__node_nodeid_impl__toString(this.Hx_Field_node.Hx_Field_id)
}

func (this *Hx_Obj_loreline_lib_gonode) Id() string {
	return this.VTable.Hx_Field_id()
}

func (this *Hx_Obj_loreline_lib_gonode) Hx_Field_toJson(pretty bool) string {
	return Hx_Field_loreline_json_stringify(this.Hx_Field_node.VTable.Hx_Field_toJson(), pretty)
}

func (this *Hx_Obj_loreline_lib_gonode) ToJson(pretty bool) string {
	return this.VTable.Hx_Field_toJson(pretty)
}
