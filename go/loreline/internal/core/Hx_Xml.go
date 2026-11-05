package core

var Hx_Abstract_xmltype_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"XmlType",
	HxMakeArray[string]("Element", "PCData", "CData", "Comment", "DocType", "ProcessingInstruction", "Document", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable__xml_xmltype_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable__xml_xmltype_impl_() {
	Hx_Field__xml_xmltype_impl__Element = Hx_Init_Hx_Field__xml_xmltype_impl__Element()
	Hx_Field__xml_xmltype_impl__PCData = Hx_Init_Hx_Field__xml_xmltype_impl__PCData()
	Hx_Field__xml_xmltype_impl__CData = Hx_Init_Hx_Field__xml_xmltype_impl__CData()
	Hx_Field__xml_xmltype_impl__Comment = Hx_Init_Hx_Field__xml_xmltype_impl__Comment()
	Hx_Field__xml_xmltype_impl__DocType = Hx_Init_Hx_Field__xml_xmltype_impl__DocType()
	Hx_Field__xml_xmltype_impl__ProcessingInstruction = Hx_Init_Hx_Field__xml_xmltype_impl__ProcessingInstruction()
	Hx_Field__xml_xmltype_impl__Document = Hx_Init_Hx_Field__xml_xmltype_impl__Document()
}

type Hx_Obj_VTable__xml_xmltype_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj__xml_xmltype_impl_ struct {
	VTable Hx_Obj_VTable__xml_xmltype_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj__xml_xmltype_impl__CreateEmptyInstance() *Hx_Obj__xml_xmltype_impl_ {
	obj := &Hx_Obj__xml_xmltype_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj__xml_xmltype_impl__CreateInstance() *Hx_Obj__xml_xmltype_impl_ {
	obj := Hx_Obj__xml_xmltype_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj__xml_xmltype_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"_Xml.XmlType_Impl_",
	HxMakeArray[string]("Element", "PCData", "CData", "Comment", "DocType", "ProcessingInstruction", "Document", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj__xml_xmltype_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj__xml_xmltype_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj__xml_xmltype_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Element":
			return &Hx_Field__xml_xmltype_impl__Element
		case "PCData":
			return &Hx_Field__xml_xmltype_impl__PCData
		case "CData":
			return &Hx_Field__xml_xmltype_impl__CData
		case "Comment":
			return &Hx_Field__xml_xmltype_impl__Comment
		case "DocType":
			return &Hx_Field__xml_xmltype_impl__DocType
		case "ProcessingInstruction":
			return &Hx_Field__xml_xmltype_impl__ProcessingInstruction
		case "Document":
			return &Hx_Field__xml_xmltype_impl__Document
		case "toString":
			return Hx_Field__xml_xmltype_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj__xml_xmltype_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj__xml_xmltype_impl__RTTI
}

func (this *Hx_Obj__xml_xmltype_impl_) Hx_As_Hx_Obj__xml_xmltype_impl_() *Hx_Obj__xml_xmltype_impl_ {
	return this
}

var Hx_Field__xml_xmltype_impl__Element int32

func Hx_Init_Hx_Field__xml_xmltype_impl__Element() int32 {
	return ((int32)(0))
}

var Hx_Field__xml_xmltype_impl__PCData int32

func Hx_Init_Hx_Field__xml_xmltype_impl__PCData() int32 {
	return ((int32)(1))
}

var Hx_Field__xml_xmltype_impl__CData int32

func Hx_Init_Hx_Field__xml_xmltype_impl__CData() int32 {
	return ((int32)(2))
}

var Hx_Field__xml_xmltype_impl__Comment int32

func Hx_Init_Hx_Field__xml_xmltype_impl__Comment() int32 {
	return ((int32)(3))
}

var Hx_Field__xml_xmltype_impl__DocType int32

func Hx_Init_Hx_Field__xml_xmltype_impl__DocType() int32 {
	return ((int32)(4))
}

var Hx_Field__xml_xmltype_impl__ProcessingInstruction int32

func Hx_Init_Hx_Field__xml_xmltype_impl__ProcessingInstruction() int32 {
	return ((int32)(5))
}

var Hx_Field__xml_xmltype_impl__Document int32

func Hx_Init_Hx_Field__xml_xmltype_impl__Document() int32 {
	return ((int32)(6))
}

func Hx_Field__xml_xmltype_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "Element"
		}

	case 1:
		{
			return "PCData"
		}

	case 2:
		{
			return "CData"
		}

	case 3:
		{
			return "Comment"
		}

	case 4:
		{
			return "DocType"
		}

	case 5:
		{
			return "ProcessingInstruction"
		}

	case 6:
		{
			return "Document"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func Hx_ClassInit_Hx_Obj_VTable_xml() {
}

func Hx_StaticInit_Hx_Obj_VTable_xml() {
	Hx_Field_xml_Element = Hx_Init_Hx_Field_xml_Element()
	Hx_Field_xml_PCData = Hx_Init_Hx_Field_xml_PCData()
	Hx_Field_xml_CData = Hx_Init_Hx_Field_xml_CData()
	Hx_Field_xml_Comment = Hx_Init_Hx_Field_xml_Comment()
	Hx_Field_xml_DocType = Hx_Init_Hx_Field_xml_DocType()
	Hx_Field_xml_ProcessingInstruction = Hx_Init_Hx_Field_xml_ProcessingInstruction()
	Hx_Field_xml_Document = Hx_Init_Hx_Field_xml_Document()
}

type Hx_Obj_VTable_xml interface {
	Hx_Field_set_nodeValue(v string) string
	Hx_Field_attributes() any
	Hx_Field_iterator() any
	Hx_Field_elements() any
	Hx_Field_toString() string
	Hx_Field_get_nodeValue() string
	Hx_Field_remove(att string)
	Hx_Field_firstElement() *Hx_Obj_xml
	Hx_Field_ensureElementType()
	Hx_Field_get_nodeName() string
	Hx_Field_elementsNamed(name string) any
	Hx_Field_firstChild() *Hx_Obj_xml
	Hx_Field_addChild(x *Hx_Obj_xml)
	Hx_Field_set_nodeName(v string) string
	Hx_Field_get(att string) HxNullable[string]
	Hx_Field_set(att string, value string)
	Hx_Field_exists(att string) bool
	Hx_Field_removeChild(x *Hx_Obj_xml) bool
	Hx_Field_insertChild(x *Hx_Obj_xml, pos int32)
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_xml struct {
	VTable                Hx_Obj_VTable_xml `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_nodeType     int32
	Hx_Field_nodeName     string
	Hx_Field_nodeValue    string
	Hx_Field_parent       *Hx_Obj_xml
	Hx_Field_children     HxArray[*Hx_Obj_xml]
	Hx_Field_attributeMap Hx_Typedef_map
}

func Hx_Obj_xml_CreateEmptyInstance() *Hx_Obj_xml {
	obj := &Hx_Obj_xml{}
	obj.VTable = obj
	obj.Hx_Field_nodeType = HxDefault[int32]()
	obj.Hx_Field_nodeName = HxDefault[string]()
	obj.Hx_Field_nodeValue = HxDefault[string]()
	obj.Hx_Field_parent = HxDefault[*Hx_Obj_xml]()
	obj.Hx_Field_children = HxDefault[HxArray[*Hx_Obj_xml]]()
	obj.Hx_Field_attributeMap = HxDefault[Hx_Typedef_map]()
	return obj
}

func Hx_Obj_xml_CreateInstance(nodeType int32) *Hx_Obj_xml {
	obj := Hx_Obj_xml_CreateEmptyInstance()
	obj.Hx_New(nodeType)
	return obj
}

func (this *Hx_Obj_xml) Hx_New(nodeType int32) {
	{
		this.Hx_Field_nodeType = nodeType
		switch nodeType {
		case Hx_Field_xml_Document:
			{
				this.Hx_Field_children = HxMakeArray[*Hx_Obj_xml]()
			}

		case Hx_Field_xml_Element:
			{
				this.Hx_Field_children = HxMakeArray[*Hx_Obj_xml]()
				this.Hx_Field_attributeMap = (func() *Hx_Obj_haxe_imap {
					_hx_icast_48 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
					if _hx_icast_48 == nil {
						return nil
					}
					Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_48, "Hx_Obj_haxe_imap")
					return &Hx_Obj_haxe_imap{VTable: _hx_icast_48.VTable.(Hx_Obj_VTable_haxe_imap)}
				})()
			}

		default:
			{
			}
		}
	}

}

var Hx_Obj_xml_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Xml",
	HxMakeArray[string]("Element", "PCData", "CData", "Comment", "DocType", "ProcessingInstruction", "Document", "parse", "createElement", "createPCData", "createCData", "createComment", "createDocType", "createProcessingInstruction", "createDocument"),
	HxMakeArray[string]("nodeType", "nodeName", "nodeValue", "parent", "children", "attributeMap", "get_nodeName", "set_nodeName", "get_nodeValue", "set_nodeValue", "get", "set", "remove", "exists", "attributes", "iterator", "elements", "elementsNamed", "firstChild", "firstElement", "addChild", "removeChild", "insertChild", "toString", "ensureElementType"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_xml_CreateInstance(HxConvert[int32](params.Get(0)))
	},
	func() any {
		return Hx_Obj_xml_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_xml_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Element":
			return &Hx_Field_xml_Element
		case "PCData":
			return &Hx_Field_xml_PCData
		case "CData":
			return &Hx_Field_xml_CData
		case "Comment":
			return &Hx_Field_xml_Comment
		case "DocType":
			return &Hx_Field_xml_DocType
		case "ProcessingInstruction":
			return &Hx_Field_xml_ProcessingInstruction
		case "Document":
			return &Hx_Field_xml_Document
		case "parse":
			return Hx_Field_xml_parse
		case "createElement":
			return Hx_Field_xml_createElement
		case "createPCData":
			return Hx_Field_xml_createPCData
		case "createCData":
			return Hx_Field_xml_createCData
		case "createComment":
			return Hx_Field_xml_createComment
		case "createDocType":
			return Hx_Field_xml_createDocType
		case "createProcessingInstruction":
			return Hx_Field_xml_createProcessingInstruction
		case "createDocument":
			return Hx_Field_xml_createDocument
		}
		return nil
	}
}

func (this *Hx_Obj_xml) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_xml_RTTI
}

func (this *Hx_Obj_xml) Hx_As_Hx_Obj_xml() *Hx_Obj_xml {
	return this
}

func (this *Hx_Obj_xml) Hx_Field_get_nodeName() string {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return this.Hx_Field_nodeName
}

func (this *Hx_Obj_xml) Hx_Field_set_nodeName(v string) string {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	this.Hx_Field_nodeName = v
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_nodeName
	return _hx_tmp_0
}

func (this *Hx_Obj_xml) Hx_Field_get_nodeValue() string {
	if (this.Hx_Field_nodeType == Hx_Field_xml_Document) || (this.Hx_Field_nodeType == Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return this.Hx_Field_nodeValue
}

func (this *Hx_Obj_xml) Hx_Field_set_nodeValue(v string) string {
	if (this.Hx_Field_nodeType == Hx_Field_xml_Document) || (this.Hx_Field_nodeType == Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	this.Hx_Field_nodeValue = v
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_nodeValue
	return _hx_tmp_0
}

func (this *Hx_Obj_xml) Hx_Field_get(att string) HxNullable[string] {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return func() func(_hx_param_key string) HxNullable[string] {
		var _hx_closure_src_372 func(_hx_param_key any) any
		_ = _hx_closure_src_372
		_hx_closure_src_372 = this.Hx_Field_attributeMap.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_get
		return func(_hx_param_key string) HxNullable[string] {
			var hx_nullable_t45 any
			_ = hx_nullable_t45
			hx_nullable_t45 = _hx_closure_src_372(((any)(_hx_param_key)))
			var _hx_tmp_0 HxNullable[string]
			_ = _hx_tmp_0
			if hx_nullable_t45 == nil {
				_hx_tmp_0 = HxNullable[string]{Value: HxStringNull}
			} else {
				_hx_tmp_0 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t45), Valid: true}
			}

			return _hx_tmp_0
		}
	}()(att)
}

func (this *Hx_Obj_xml) Hx_Field_set(att string, value string) {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_373 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_373
		_hx_closure_src_373 = this.Hx_Field_attributeMap.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_373(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(att, value)
}

func (this *Hx_Obj_xml) Hx_Field_remove(att string) {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	func() func(_hx_param_key string) bool {
		var _hx_closure_src_374 func(_hx_param_key any) bool
		_ = _hx_closure_src_374
		_hx_closure_src_374 = this.Hx_Field_attributeMap.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_remove
		return func(_hx_param_key string) bool {
			return _hx_closure_src_374(((any)(_hx_param_key)))
		}
	}()(att)
}

func (this *Hx_Obj_xml) Hx_Field_exists(att string) bool {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return func() func(_hx_param_key string) bool {
		var _hx_closure_src_375 func(_hx_param_key any) bool
		_ = _hx_closure_src_375
		_hx_closure_src_375 = this.Hx_Field_attributeMap.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_375(((any)(_hx_param_key)))
		}
	}()(att)
}

func (this *Hx_Obj_xml) Hx_Field_attributes() any {
	if this.Hx_Field_nodeType != Hx_Field_xml_Element {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return this.Hx_Field_attributeMap.VTable.(interface {
		Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
	}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_keys()
}

func (this *Hx_Obj_xml) Hx_Field_iterator() any {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any](this.Hx_Field_children))))
}

func (this *Hx_Obj_xml) Hx_Field_elements() any {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	var _g HxArray[*Hx_Obj_xml]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_xml]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_xml]
		_ = _g2
		_g2 = this.Hx_Field_children
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var child *Hx_Obj_xml
			_ = child
			child = _g2.Get(_g1)
			_g1++
			if child.Hx_Field_nodeType == Hx_Field_xml_Element {
				_ = (Hx_Array_Push[*Hx_Obj_xml](_g, child))
			}
		}
	}

	var ret HxArray[*Hx_Obj_xml]
	_ = ret
	ret = _g
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any](ret))))
}

func (this *Hx_Obj_xml) Hx_Field_elementsNamed(name string) any {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	var _g HxArray[*Hx_Obj_xml]
	_ = _g
	_g = HxMakeArray[*Hx_Obj_xml]()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_xml]
		_ = _g2
		_g2 = this.Hx_Field_children
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var child *Hx_Obj_xml
			_ = child
			child = _g2.Get(_g1)
			_g1++
			var tmp bool
			_ = tmp
			if child.Hx_Field_nodeType == Hx_Field_xml_Element {
				if child.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(child.Hx_Field_nodeType))))))
				}

				tmp = (child.Hx_Field_nodeName == name)
			} else {
				tmp = false
			}

			if tmp {
				_ = (Hx_Array_Push[*Hx_Obj_xml](_g, child))
			}
		}
	}

	var ret HxArray[*Hx_Obj_xml]
	_ = ret
	ret = _g
	return ((any)(Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any](ret))))
}

func (this *Hx_Obj_xml) Hx_Field_firstChild() *Hx_Obj_xml {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	return this.Hx_Field_children.Get(0)
}

func (this *Hx_Obj_xml) Hx_Field_firstElement() *Hx_Obj_xml {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_xml]
		_ = _g1
		_g1 = this.Hx_Field_children
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var child *Hx_Obj_xml
			_ = child
			child = _g1.Get(_g)
			_g++
			if child.Hx_Field_nodeType == Hx_Field_xml_Element {
				return child
			}
		}
	}

	return nil
}

func (this *Hx_Obj_xml) Hx_Field_addChild(x *Hx_Obj_xml) {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	if x.Hx_Field_parent != nil {
		x.Hx_Field_parent.VTable.Hx_Field_removeChild(x)
	}

	_ = (Hx_Array_Push[*Hx_Obj_xml](this.Hx_Field_children, x))
	x.Hx_Field_parent = this
}

func (this *Hx_Obj_xml) Hx_Field_removeChild(x *Hx_Obj_xml) bool {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	if Hx_Array_Remove[*Hx_Obj_xml](this.Hx_Field_children, x) {
		x.Hx_Field_parent = nil
		return true
	}

	return false
}

func (this *Hx_Obj_xml) Hx_Field_insertChild(x *Hx_Obj_xml, pos int32) {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	if x.Hx_Field_parent != nil {
		_ = (Hx_Array_Remove[*Hx_Obj_xml](x.Hx_Field_parent.Hx_Field_children, x))
	}

	(Hx_Array_Insert[*Hx_Obj_xml](this.Hx_Field_children, pos, x))
	x.Hx_Field_parent = this
}

func (this *Hx_Obj_xml) Hx_Field_toString() string {
	var _hx_tmp_0 *Hx_Obj_xml
	_ = _hx_tmp_0
	_hx_tmp_0 = this
	return Hx_Field_haxe_xml_printer_print(_hx_tmp_0, nil)
}

func (this *Hx_Obj_xml) Hx_Field_ensureElementType() {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}
}

var Hx_Field_xml_Element int32

func Hx_Init_Hx_Field_xml_Element() int32 {
	return ((int32)(0))
}

var Hx_Field_xml_PCData int32

func Hx_Init_Hx_Field_xml_PCData() int32 {
	return ((int32)(1))
}

var Hx_Field_xml_CData int32

func Hx_Init_Hx_Field_xml_CData() int32 {
	return ((int32)(2))
}

var Hx_Field_xml_Comment int32

func Hx_Init_Hx_Field_xml_Comment() int32 {
	return ((int32)(3))
}

var Hx_Field_xml_DocType int32

func Hx_Init_Hx_Field_xml_DocType() int32 {
	return ((int32)(4))
}

var Hx_Field_xml_ProcessingInstruction int32

func Hx_Init_Hx_Field_xml_ProcessingInstruction() int32 {
	return ((int32)(5))
}

var Hx_Field_xml_Document int32

func Hx_Init_Hx_Field_xml_Document() int32 {
	return ((int32)(6))
}

func Hx_Field_xml_parse(str string) *Hx_Obj_xml {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = str
	return Hx_Field_haxe_xml_parser_parse(_hx_tmp_0, nil)
}

func Hx_Field_xml_createElement(name string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_Element)
	{
		if xml.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeName = name
	}

	return xml
}

func Hx_Field_xml_createPCData(data string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_PCData)
	{
		if (xml.Hx_Field_nodeType == Hx_Field_xml_Document) || (xml.Hx_Field_nodeType == Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeValue = data
	}

	return xml
}

func Hx_Field_xml_createCData(data string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_CData)
	{
		if (xml.Hx_Field_nodeType == Hx_Field_xml_Document) || (xml.Hx_Field_nodeType == Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeValue = data
	}

	return xml
}

func Hx_Field_xml_createComment(data string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_Comment)
	{
		if (xml.Hx_Field_nodeType == Hx_Field_xml_Document) || (xml.Hx_Field_nodeType == Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeValue = data
	}

	return xml
}

func Hx_Field_xml_createDocType(data string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_DocType)
	{
		if (xml.Hx_Field_nodeType == Hx_Field_xml_Document) || (xml.Hx_Field_nodeType == Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeValue = data
	}

	return xml
}

func Hx_Field_xml_createProcessingInstruction(data string) *Hx_Obj_xml {
	var xml *Hx_Obj_xml
	_ = xml
	xml = Hx_Obj_xml_CreateInstance(Hx_Field_xml_ProcessingInstruction)
	{
		if (xml.Hx_Field_nodeType == Hx_Field_xml_Document) || (xml.Hx_Field_nodeType == Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(xml.Hx_Field_nodeType))))))
		}

		xml.Hx_Field_nodeValue = data
	}

	return xml
}

func Hx_Field_xml_createDocument() *Hx_Obj_xml {
	return Hx_Obj_xml_CreateInstance(Hx_Field_xml_Document)
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_xml]()
}
