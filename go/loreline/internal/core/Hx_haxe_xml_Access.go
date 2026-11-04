package core

var Hx_Abstract_haxe_xml__access_nodeaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.NodeAccess",
	HxMakeArray[string]("resolve"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_nodeaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_nodeaccess_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_nodeaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_nodeaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_nodeaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_nodeaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_nodeaccess_impl_ {
	obj := &Hx_Obj_haxe_xml__access_nodeaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_nodeaccess_impl__CreateInstance() *Hx_Obj_haxe_xml__access_nodeaccess_impl_ {
	obj := Hx_Obj_haxe_xml__access_nodeaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_nodeaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.NodeAccess_Impl_",
	HxMakeArray[string]("resolve"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_nodeaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_nodeaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_nodeaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "resolve":
			return Hx_Field_haxe_xml__access_nodeaccess_impl__resolve
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_nodeaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_nodeaccess_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_nodeaccess_impl_) Hx_As_Hx_Obj_haxe_xml__access_nodeaccess_impl_() *Hx_Obj_haxe_xml__access_nodeaccess_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_nodeaccess_impl__resolve(this *Hx_Obj_xml, name string) *Hx_Obj_xml {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(this.VTable.Hx_Field_elementsNamed(name), "next")
	var hx_dyncast_49 any
	_ = hx_dyncast_49
	hx_dyncast_49 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
	var _hx_tmp_1 *Hx_Obj_xml
	_ = _hx_tmp_1
	if hx_dyncast_49 != nil {
		_hx_tmp_1 = hx_dyncast_49.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
	} else {
		_hx_tmp_1 = nil
	}

	var x *Hx_Obj_xml
	_ = x
	x = _hx_tmp_1
	if x == nil {
		var xname string
		_ = xname
		if this.Hx_Field_nodeType == Hx_Field_xml_Document {
			xname = "Document"
		} else {
			if this.Hx_Field_nodeType != Hx_Field_xml_Element {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
			}

			xname = this.Hx_Field_nodeName
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(xname, " is missing element "), name)))))
	}

	if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
	}

	return x
}

var Hx_Abstract_haxe_xml__access_attribaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.AttribAccess",
	HxMakeArray[string]("resolve", "_hx_set"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_attribaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_attribaccess_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_attribaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_attribaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_attribaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_attribaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_attribaccess_impl_ {
	obj := &Hx_Obj_haxe_xml__access_attribaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_attribaccess_impl__CreateInstance() *Hx_Obj_haxe_xml__access_attribaccess_impl_ {
	obj := Hx_Obj_haxe_xml__access_attribaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_attribaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.AttribAccess_Impl_",
	HxMakeArray[string]("resolve", "_hx_set"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_attribaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_attribaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_attribaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "resolve":
			return Hx_Field_haxe_xml__access_attribaccess_impl__resolve
		case "_hx_set":
			return Hx_Field_haxe_xml__access_attribaccess_impl___hx_set
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_attribaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_attribaccess_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_attribaccess_impl_) Hx_As_Hx_Obj_haxe_xml__access_attribaccess_impl_() *Hx_Obj_haxe_xml__access_attribaccess_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_attribaccess_impl__resolve(this *Hx_Obj_xml, name string) string {
	if this.Hx_Field_nodeType == Hx_Field_xml_Document {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access document attribute ", name)))))
	}

	var v HxNullable[string]
	_ = v
	v = this.VTable.Hx_Field_get(name)
	if (v).Valid == false {
		if this.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_go_haxe_hxdynamic_stringConcat(this.Hx_Field_nodeName, " is missing attribute "), name)))))
	}

	return (v).Value
}

func Hx_Field_haxe_xml__access_attribaccess_impl___hx_set(this *Hx_Obj_xml, name string, value string) string {
	if this.Hx_Field_nodeType == Hx_Field_xml_Document {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access document attribute ", name)))))
	}

	this.VTable.Hx_Field_set(name, value)
	return value
}

var Hx_Abstract_haxe_xml__access_hasattribaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.HasAttribAccess",
	HxMakeArray[string]("resolve"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_hasattribaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_hasattribaccess_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_hasattribaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_hasattribaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_hasattribaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_hasattribaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_hasattribaccess_impl_ {
	obj := &Hx_Obj_haxe_xml__access_hasattribaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_hasattribaccess_impl__CreateInstance() *Hx_Obj_haxe_xml__access_hasattribaccess_impl_ {
	obj := Hx_Obj_haxe_xml__access_hasattribaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_hasattribaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.HasAttribAccess_Impl_",
	HxMakeArray[string]("resolve"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_hasattribaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_hasattribaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_hasattribaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "resolve":
			return Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_hasattribaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_hasattribaccess_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_hasattribaccess_impl_) Hx_As_Hx_Obj_haxe_xml__access_hasattribaccess_impl_() *Hx_Obj_haxe_xml__access_hasattribaccess_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(this *Hx_Obj_xml, name string) bool {
	if this.Hx_Field_nodeType == Hx_Field_xml_Document {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Cannot access document attribute ", name)))))
	}

	return this.VTable.Hx_Field_exists(name)
}

var Hx_Abstract_haxe_xml__access_hasnodeaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.HasNodeAccess",
	HxMakeArray[string]("resolve"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_hasnodeaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_hasnodeaccess_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_hasnodeaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_hasnodeaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_hasnodeaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_hasnodeaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_hasnodeaccess_impl_ {
	obj := &Hx_Obj_haxe_xml__access_hasnodeaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_hasnodeaccess_impl__CreateInstance() *Hx_Obj_haxe_xml__access_hasnodeaccess_impl_ {
	obj := Hx_Obj_haxe_xml__access_hasnodeaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_hasnodeaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.HasNodeAccess_Impl_",
	HxMakeArray[string]("resolve"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_hasnodeaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_hasnodeaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_hasnodeaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "resolve":
			return Hx_Field_haxe_xml__access_hasnodeaccess_impl__resolve
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_hasnodeaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_hasnodeaccess_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_hasnodeaccess_impl_) Hx_As_Hx_Obj_haxe_xml__access_hasnodeaccess_impl_() *Hx_Obj_haxe_xml__access_hasnodeaccess_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_hasnodeaccess_impl__resolve(this *Hx_Obj_xml, name string) bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(this.VTable.Hx_Field_elementsNamed(name), "hasNext")
	return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))))
}

var Hx_Abstract_haxe_xml__access_nodelistaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.NodeListAccess",
	HxMakeArray[string]("resolve"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_nodelistaccess_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_nodelistaccess_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_nodelistaccess_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_nodelistaccess_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_nodelistaccess_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_nodelistaccess_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_nodelistaccess_impl_ {
	obj := &Hx_Obj_haxe_xml__access_nodelistaccess_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_nodelistaccess_impl__CreateInstance() *Hx_Obj_haxe_xml__access_nodelistaccess_impl_ {
	obj := Hx_Obj_haxe_xml__access_nodelistaccess_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_nodelistaccess_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.NodeListAccess_Impl_",
	HxMakeArray[string]("resolve"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_nodelistaccess_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_nodelistaccess_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_nodelistaccess_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "resolve":
			return Hx_Field_haxe_xml__access_nodelistaccess_impl__resolve
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_nodelistaccess_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_nodelistaccess_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_nodelistaccess_impl_) Hx_As_Hx_Obj_haxe_xml__access_nodelistaccess_impl_() *Hx_Obj_haxe_xml__access_nodelistaccess_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_nodelistaccess_impl__resolve(this *Hx_Obj_xml, name string) HxArray[*Hx_Obj_xml] {
	var l HxArray[*Hx_Obj_xml]
	_ = l
	l = HxMakeArray[*Hx_Obj_xml]()
	{
		var x any
		_ = x
		x = this.VTable.Hx_Field_elementsNamed(name)
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(x, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(x, "next")
			var hx_dyncast_50 any
			_ = hx_dyncast_50
			hx_dyncast_50 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_50 != nil {
				_hx_tmp_2 = hx_dyncast_50.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var x1 *Hx_Obj_xml
			_ = x1
			x1 = _hx_tmp_2
			if (x1.Hx_Field_nodeType != Hx_Field_xml_Document) && (x1.Hx_Field_nodeType != Hx_Field_xml_Element) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x1.Hx_Field_nodeType))))))
			}

			_ = (Hx_Array_Push[*Hx_Obj_xml](l, x1))
		}
	}

	return l
}

var Hx_Abstract_haxe_xml_access_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml.Access",
	HxMakeArray[string]("x", "get_x", "name", "get_name", "innerData", "innerHTML", "node", "get_node", "nodes", "get_nodes", "att", "get_att", "has", "get_has", "hasNode", "get_hasNode", "elements", "get_elements", "_hx_new", "get_innerData", "get_innerHTML"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__access_access_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__access_access_impl_() {
}

type Hx_Obj_VTable_haxe_xml__access_access_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__access_access_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__access_access_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__access_access_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__access_access_impl_ {
	obj := &Hx_Obj_haxe_xml__access_access_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__access_access_impl__CreateInstance() *Hx_Obj_haxe_xml__access_access_impl_ {
	obj := Hx_Obj_haxe_xml__access_access_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__access_access_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Access.Access_Impl_",
	HxMakeArray[string]("get_x", "get_name", "get_node", "get_nodes", "get_att", "get_has", "get_hasNode", "get_elements", "_hx_new", "get_innerData", "get_innerHTML"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__access_access_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__access_access_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__access_access_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "get_x":
			return Hx_Field_haxe_xml__access_access_impl__get_x
		case "get_name":
			return Hx_Field_haxe_xml__access_access_impl__get_name
		case "get_node":
			return Hx_Field_haxe_xml__access_access_impl__get_node
		case "get_nodes":
			return Hx_Field_haxe_xml__access_access_impl__get_nodes
		case "get_att":
			return Hx_Field_haxe_xml__access_access_impl__get_att
		case "get_has":
			return Hx_Field_haxe_xml__access_access_impl__get_has
		case "get_hasNode":
			return Hx_Field_haxe_xml__access_access_impl__get_hasNode
		case "get_elements":
			return Hx_Field_haxe_xml__access_access_impl__get_elements
		case "_hx_new":
			return Hx_Field_haxe_xml__access_access_impl___hx_new
		case "get_innerData":
			return Hx_Field_haxe_xml__access_access_impl__get_innerData
		case "get_innerHTML":
			return Hx_Field_haxe_xml__access_access_impl__get_innerHTML
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__access_access_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__access_access_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__access_access_impl_) Hx_As_Hx_Obj_haxe_xml__access_access_impl_() *Hx_Obj_haxe_xml__access_access_impl_ {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_x(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_name(this *Hx_Obj_xml) string {
	if this.Hx_Field_nodeType == Hx_Field_xml_Document {
		return "Document"
	} else {
		if this.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
		}

		return this.Hx_Field_nodeName
	}
}

func Hx_Field_haxe_xml__access_access_impl__get_node(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_nodes(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_att(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_has(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_hasNode(this *Hx_Obj_xml) *Hx_Obj_xml {
	return this
}

func Hx_Field_haxe_xml__access_access_impl__get_elements(this *Hx_Obj_xml) any {
	return this.VTable.Hx_Field_elements()
}

func Hx_Field_haxe_xml__access_access_impl___hx_new(x *Hx_Obj_xml) *Hx_Obj_xml {
	if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
	}

	return x
}

func Hx_Field_haxe_xml__access_access_impl__get_innerData(this *Hx_Obj_xml) string {
	if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
	}

	var it_current int32
	_ = it_current
	it_current = 0
	var it_array HxArray[*Hx_Obj_xml]
	_ = it_array
	it_array = this.Hx_Field_children
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = it_current
	if _hx_tmp_0 >= it_array.Len() {
		var tmp string
		_ = tmp
		if this.Hx_Field_nodeType == Hx_Field_xml_Document {
			tmp = "Document"
		} else {
			if this.Hx_Field_nodeType != Hx_Field_xml_Element {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
			}

			tmp = this.Hx_Field_nodeName
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(tmp, " does not have data")))))
	}

	var _hx_tmp_1 HxArray[*Hx_Obj_xml]
	_ = _hx_tmp_1
	_hx_tmp_1 = it_array
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = it_current
	it_current = (it_current + 1)
	var v *Hx_Obj_xml
	_ = v
	v = _hx_tmp_1.Get(_hx_tmp_2)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = it_current
	if _hx_tmp_3 < it_array.Len() {
		var _hx_tmp_4 HxArray[*Hx_Obj_xml]
		_ = _hx_tmp_4
		_hx_tmp_4 = it_array
		var _hx_tmp_5 int32
		_ = _hx_tmp_5
		_hx_tmp_5 = it_current
		it_current = (it_current + 1)
		var n *Hx_Obj_xml
		_ = n
		n = _hx_tmp_4.Get(_hx_tmp_5)
		var tmp bool
		_ = tmp
		if (v.Hx_Field_nodeType == Hx_Field_xml_PCData) && (n.Hx_Field_nodeType == Hx_Field_xml_CData) {
			if (v.Hx_Field_nodeType == Hx_Field_xml_Document) || (v.Hx_Field_nodeType == Hx_Field_xml_Element) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(v.Hx_Field_nodeType))))))
			}

			tmp = (Hx_Field_stringtools_trim(v.Hx_Field_nodeValue) == "")
		} else {
			tmp = false
		}

		if tmp {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = it_current
			if _hx_tmp_6 >= it_array.Len() {
				if (n.Hx_Field_nodeType == Hx_Field_xml_Document) || (n.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(n.Hx_Field_nodeType))))))
				}

				return n.Hx_Field_nodeValue
			}

			var _hx_tmp_7 HxArray[*Hx_Obj_xml]
			_ = _hx_tmp_7
			_hx_tmp_7 = it_array
			var _hx_tmp_8 int32
			_ = _hx_tmp_8
			_hx_tmp_8 = it_current
			it_current = (it_current + 1)
			var n2 *Hx_Obj_xml
			_ = n2
			n2 = _hx_tmp_7.Get(_hx_tmp_8)
			var tmp_tmp_1 bool
			_ = tmp_tmp_1
			if n2.Hx_Field_nodeType == Hx_Field_xml_PCData {
				if (n2.Hx_Field_nodeType == Hx_Field_xml_Document) || (n2.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(n2.Hx_Field_nodeType))))))
				}

				tmp_tmp_1 = (Hx_Field_stringtools_trim(n2.Hx_Field_nodeValue) == "")
			} else {
				tmp_tmp_1 = false
			}

			var _hx_tmp_9 bool
			_ = _hx_tmp_9
			if tmp_tmp_1 {
				var _hx_tmp_10 int32
				_ = _hx_tmp_10
				_hx_tmp_10 = it_current
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _hx_tmp_10
				_hx_tmp_9 = (_hx_tmp_11 >= it_array.Len())
			} else {
				_hx_tmp_9 = false
			}

			if _hx_tmp_9 {
				if (n.Hx_Field_nodeType == Hx_Field_xml_Document) || (n.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(n.Hx_Field_nodeType))))))
				}

				return n.Hx_Field_nodeValue
			}
		}

		var tmp1 string
		_ = tmp1
		if this.Hx_Field_nodeType == Hx_Field_xml_Document {
			tmp1 = "Document"
		} else {
			if this.Hx_Field_nodeType != Hx_Field_xml_Element {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
			}

			tmp1 = this.Hx_Field_nodeName
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(tmp1, " does not only have data")))))
	}

	if (v.Hx_Field_nodeType != Hx_Field_xml_PCData) && (v.Hx_Field_nodeType != Hx_Field_xml_CData) {
		var tmp string
		_ = tmp
		if this.Hx_Field_nodeType == Hx_Field_xml_Document {
			tmp = "Document"
		} else {
			if this.Hx_Field_nodeType != Hx_Field_xml_Element {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
			}

			tmp = this.Hx_Field_nodeName
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat(tmp, " does not have data")))))
	}

	if (v.Hx_Field_nodeType == Hx_Field_xml_Document) || (v.Hx_Field_nodeType == Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(v.Hx_Field_nodeType))))))
	}

	return v.Hx_Field_nodeValue
}

func Hx_Field_haxe_xml__access_access_impl__get_innerHTML(this *Hx_Obj_xml) string {
	var s *Hx_Obj_stringbuf
	_ = s
	s = Hx_Obj_stringbuf_CreateInstance()
	{
		if (this.Hx_Field_nodeType != Hx_Field_xml_Document) && (this.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxdynamic_stringConcat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(this.Hx_Field_nodeType))))))
		}

		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[*Hx_Obj_xml]
		_ = _g_array
		_g_array = this.Hx_Field_children
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g_current
			if !(_hx_tmp_0 < _g_array.Len()) {
				break
			}

			var _hx_tmp_1 HxArray[*Hx_Obj_xml]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g_array
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g_current
			_g_current = (_g_current + 1)
			var x *Hx_Obj_xml
			_ = x
			x = _hx_tmp_1.Get(_hx_tmp_2)
			var _hx_tmp_3 *Hx_Obj_xml
			_ = _hx_tmp_3
			_hx_tmp_3 = x
			func() func(x string) {
				var _hx_closure_src_384 func(x any)
				_ = _hx_closure_src_384
				_hx_closure_src_384 = s.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_384(((any)(x)))
				}
			}()(Hx_Field_haxe_xml_printer_print(_hx_tmp_3, nil))
		}
	}

	return s.VTable.Hx_Field_toString()
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_xml]()
}
