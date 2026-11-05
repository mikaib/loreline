package core

func Hx_ClassInit_Hx_Obj_VTable_haxe_rtti_xmlparser() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_rtti_xmlparser() {
}

type Hx_Obj_VTable_haxe_rtti_xmlparser interface {
	Hx_Field_xtypedef(x *Hx_Obj_xml) any
	Hx_Field_sort(l__isOptional any)
	Hx_Field_mergeRights(f1 any, f2 any) bool
	Hx_Field_mkRights(r string) Hx_Enum_haxe_rtti_rights
	Hx_Field_xtypeparams(x *Hx_Obj_xml) HxArray[Hx_Enum_haxe_rtti_ctype]
	Hx_Field_sortFields(a HxArray[any])
	Hx_Field_process(x *Hx_Obj_xml, platform string)
	Hx_Field_mergeDoc(f1 any, f2 any) bool
	Hx_Field_xmeta(x *Hx_Obj_xml) HxArray[any]
	Hx_Field_xpath(x *Hx_Obj_xml) any
	Hx_Field_mergeClasses(c any, c2 any) bool
	Hx_Field_xerror(c *Hx_Obj_xml) any
	Hx_Field_defplat() HxArray[string]
	Hx_Field_mergeAbstracts(a any, a2 any) bool
	Hx_Field_mkPath(p string) Hx_Typedef_haxe_rtti_path
	Hx_Field_mkTypeParams(p string) Hx_Typedef_haxe_rtti_typeparams
	Hx_Field_mergeEnums(e any, e2 any) bool
	Hx_Field_mergeTypedefs(t any, t2 any) bool
	Hx_Field_processElement(x *Hx_Obj_xml) Hx_Enum_haxe_rtti_typetree
	Hx_Field_xoverloads(x *Hx_Obj_xml) HxArray[any]
	Hx_Field_xenum(x *Hx_Obj_xml) any
	Hx_Field_mergeFields(f any, f2 any) bool
	Hx_Field_merge(t Hx_Enum_haxe_rtti_typetree)
	Hx_Field_xroot(x *Hx_Obj_xml)
	Hx_Field_xclass(x *Hx_Obj_xml) any
	Hx_Field_xenumfield(x *Hx_Obj_xml) any
	Hx_Field_xtype(x *Hx_Obj_xml) Hx_Enum_haxe_rtti_ctype
	Hx_Field_newField(c any, f any)
	Hx_Field_xclassfield(x *Hx_Obj_xml, defPublic__isOptional any) any
	Hx_Field_xabstract(x *Hx_Obj_xml) any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_rtti_xmlparser struct {
	VTable                Hx_Obj_VTable_haxe_rtti_xmlparser `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_root         Hx_Typedef_haxe_rtti_typeroot
	Hx_Field_curplatform  string
	Hx_Field_newField_Dyn func(this *Hx_Obj_haxe_rtti_xmlparser, c any, f any)
}

func Hx_Obj_haxe_rtti_xmlparser_CreateEmptyInstance() *Hx_Obj_haxe_rtti_xmlparser {
	obj := &Hx_Obj_haxe_rtti_xmlparser{}
	obj.VTable = obj
	obj.Hx_Field_newField_Dyn = func(this *Hx_Obj_haxe_rtti_xmlparser, c any, f any) {
	}
	obj.Hx_Field_root = HxDefault[Hx_Typedef_haxe_rtti_typeroot]()
	obj.Hx_Field_curplatform = HxDefault[string]()
	return obj
}

func Hx_Obj_haxe_rtti_xmlparser_CreateInstance() *Hx_Obj_haxe_rtti_xmlparser {
	obj := Hx_Obj_haxe_rtti_xmlparser_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_New() {
	{
		this.Hx_Field_root = ((Hx_Typedef_haxe_rtti_typeroot)(HxMakeArray[Hx_Enum_haxe_rtti_typetree]()))
	}

}

var Hx_Obj_haxe_rtti_xmlparser_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.rtti.XmlParser",
	HxMakeArray[string](),
	HxMakeArray[string]("root", "curplatform", "sort", "sortFields", "process", "mergeRights", "mergeDoc", "mergeFields", "newField", "mergeClasses", "mergeEnums", "mergeTypedefs", "mergeAbstracts", "merge", "mkPath", "mkTypeParams", "mkRights", "xerror", "xroot", "processElement", "xmeta", "xoverloads", "xpath", "xclass", "xclassfield", "xenum", "xenumfield", "xabstract", "xtypedef", "xtype", "xtypeparams", "defplat"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_rtti_xmlparser_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_rtti_xmlparser_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_rtti_xmlparser_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_rtti_xmlparser_RTTI
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_As_Hx_Obj_haxe_rtti_xmlparser() *Hx_Obj_haxe_rtti_xmlparser {
	return this
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_sort(l__isOptional any) {
	l := HxDefault[Hx_Typedef_haxe_rtti_typeroot]()
	_ = l

	if l__isOptional != nil {
		l = HxConvert[Hx_Typedef_haxe_rtti_typeroot](l__isOptional)
	} else {
		l = nil
	}
	{
		if l == nil {
			l = this.Hx_Field_root
		}

		var _hx_tmp_0 Hx_Typedef_haxe_rtti_typeroot
		_ = _hx_tmp_0
		_hx_tmp_0 = l
		(Hx_Array_Sort[Hx_Enum_haxe_rtti_typetree](_hx_tmp_0, func(e1 Hx_Enum_haxe_rtti_typetree, e2 Hx_Enum_haxe_rtti_typetree) int32 {
			var n1 string
			_ = n1
			if e1.Hx_Field_enumIndex() == 0 {
				var p string
				_ = p
				p = e1.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_name
				n1 = Hx_Field_go_haxe_hxstring_concat(" ", p)
			} else {
				n1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_haxe_rtti_typeapi_typeInfos(e1), "path"))
			}

			var n2 string
			_ = n2
			if e2.Hx_Field_enumIndex() == 0 {
				var p string
				_ = p
				p = e2.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_name
				n2 = Hx_Field_go_haxe_hxstring_concat(" ", p)
			} else {
				n2 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_haxe_rtti_typeapi_typeInfos(e2), "path"))
			}

			if n1 > n2 {
				return 1
			}

			return -1
		}))
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				if !(_hx_tmp_1 < l.Len()) {
					break
				}

				var x Hx_Enum_haxe_rtti_typetree
				_ = x
				x = l.Get(_g)
				_g++
				switch x.Hx_Field_enumIndex() {
				case 0:
					{
						var l HxArray[Hx_Enum_haxe_rtti_typetree]
						_ = l
						l = x.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_subs
						this.VTable.Hx_Field_sort(((any)(l)))
					}

				case 1:
					{
						var c any
						_ = c
						c = x.(Hx_Enum_haxe_rtti_typetree_TClassdecl).Hx_Field_c
						{
							this.VTable.Hx_Field_sortFields(HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "fields")))
							this.VTable.Hx_Field_sortFields(HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "statics")))
						}
					}

				case 2:
					{
					}

				case 3:
					{
					}

				case 4:
					{
					}

				}
			}
		}
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_sortFields(a HxArray[any]) {
	var _hx_tmp_0 HxArray[any]
	_ = _hx_tmp_0
	_hx_tmp_0 = HxMakeArrayView[any](a)
	(Hx_Array_Sort[any](_hx_tmp_0, func(f1 any, f2 any) int32 {
		var v1 bool
		_ = v1
		v1 = Hx_Field_haxe_rtti_typeapi_isVar(func() Hx_Enum_haxe_rtti_ctype {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "type"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_haxe_rtti_ctype
				return _hx_z
			}
			return _hx_d.(Hx_Enum_haxe_rtti_ctype)
		}())
		var v2 bool
		_ = v2
		v2 = Hx_Field_haxe_rtti_typeapi_isVar(func() Hx_Enum_haxe_rtti_ctype {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "type"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_haxe_rtti_ctype
				return _hx_z
			}
			return _hx_d.(Hx_Enum_haxe_rtti_ctype)
		}())
		if v1 && !v2 {
			return -1
		}

		if v2 && !v1 {
			return 1
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f1, "name"), ((any)("new"))))) {
			return -1
		}

		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f2, "name"), ((any)("new"))))) {
			return 1
		}

		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(f1, "name")
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_gt(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(f2, "name")))) {
			return 1
		}

		return -1
	}))
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_process(x *Hx_Obj_xml, platform string) {
	this.Hx_Field_curplatform = platform
	if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
	}

	this.VTable.Hx_Field_xroot(x)
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeRights(f1 any, f2 any) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	var _hx_tmp_2 bool
	_ = _hx_tmp_2
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f1, "get"), ((any)(Hx_Enum_haxe_rtti_rights_RInline{}))))) {
		_hx_tmp_2 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f1, "set"), ((any)(Hx_Enum_haxe_rtti_rights_RNo{}))))))
	} else {
		_hx_tmp_2 = false
	}

	if _hx_tmp_2 {
		_hx_tmp_1 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f2, "get"), ((any)(Hx_Enum_haxe_rtti_rights_RNormal{}))))))
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(f2, "set"), ((any)(Hx_Enum_haxe_rtti_rights_RMethod{}))))))
	} else {
		_hx_tmp_0 = false
	}

	if _hx_tmp_0 {
		Hx_Field_go_haxe_hxdynamic_setField(f1, "get", Hx_Enum_haxe_rtti_rights_RNormal{})
		Hx_Field_go_haxe_hxdynamic_setField(f1, "set", Hx_Enum_haxe_rtti_rights_RMethod{})
		return true
	}

	var _hx_tmp_3 bool
	_ = _hx_tmp_3
	var _hx_tmp_4 Hx_Enum_haxe_rtti_rights
	_ = _hx_tmp_4
	_hx_tmp_4 = func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "get"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()
	if func() func(a Hx_Enum_haxe_rtti_rights, b Hx_Enum_haxe_rtti_rights) bool {
		var _hx_closure_src_385 func(a any, b any) bool
		_ = _hx_closure_src_385
		_hx_closure_src_385 = Hx_Field_type_enumEq
		return func(a Hx_Enum_haxe_rtti_rights, b Hx_Enum_haxe_rtti_rights) bool {
			return _hx_closure_src_385(((any)(a)), ((any)(b)))
		}
	}()(_hx_tmp_4, func() Hx_Enum_haxe_rtti_rights {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "get"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_rights
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_rights)
	}()) {
		var _hx_tmp_5 Hx_Enum_haxe_rtti_rights
		_ = _hx_tmp_5
		_hx_tmp_5 = func() Hx_Enum_haxe_rtti_rights {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f1, "set"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_haxe_rtti_rights
				return _hx_z
			}
			return _hx_d.(Hx_Enum_haxe_rtti_rights)
		}()
		var _hx_tmp_6 Hx_Enum_haxe_rtti_rights
		_ = _hx_tmp_6
		_hx_tmp_6 = _hx_tmp_5
		_hx_tmp_3 = func() func(a Hx_Enum_haxe_rtti_rights, b Hx_Enum_haxe_rtti_rights) bool {
			var _hx_closure_src_386 func(a any, b any) bool
			_ = _hx_closure_src_386
			_hx_closure_src_386 = Hx_Field_type_enumEq
			return func(a Hx_Enum_haxe_rtti_rights, b Hx_Enum_haxe_rtti_rights) bool {
				return _hx_closure_src_386(((any)(a)), ((any)(b)))
			}
		}()(_hx_tmp_6, func() Hx_Enum_haxe_rtti_rights {
			_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(f2, "set"))
			if _hx_d == nil {
				var _hx_z Hx_Enum_haxe_rtti_rights
				return _hx_z
			}
			return _hx_d.(Hx_Enum_haxe_rtti_rights)
		}())
	} else {
		_hx_tmp_3 = false
	}

	return _hx_tmp_3
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeDoc(f1 any, f2 any) bool {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(f1, "doc"), nil)))))), ((any)(false))))) {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = f1
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "doc", Hx_Field_go_haxe_hxdynamic_getField(f2, "doc"))
	} else {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(f2, "doc"), nil)))))), ((any)(false))))) {
			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = f2
			Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "doc", Hx_Field_go_haxe_hxdynamic_getField(f1, "doc"))
		}
	}

	return true
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeFields(f any, f2 any) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if Hx_Field_haxe_rtti_typeapi_fieldEq(f, f2) {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 bool
		_ = _hx_tmp_1
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(f, "name")
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_4, Hx_Field_go_haxe_hxdynamic_getField(f2, "name")))) {
			var _hx_tmp_5 bool
			_ = _hx_tmp_5
			if this.VTable.Hx_Field_mergeRights(f, f2) {
				_hx_tmp_5 = true
			} else {
				_hx_tmp_5 = this.VTable.Hx_Field_mergeRights(f2, f)
			}

			_hx_tmp_3 = (_hx_tmp_5)
		} else {
			_hx_tmp_3 = false
		}

		if _hx_tmp_3 {
			_hx_tmp_2 = this.VTable.Hx_Field_mergeDoc(f, f2)
		} else {
			_hx_tmp_2 = false
		}

		if _hx_tmp_2 {
			_hx_tmp_1 = Hx_Field_haxe_rtti_typeapi_fieldEq(f, f2)
		} else {
			_hx_tmp_1 = false
		}

		_hx_tmp_0 = _hx_tmp_1
	}

	return _hx_tmp_0
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_newField(c any, f any) {
	this.Hx_Field_newField_Dyn(this, c, f)
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeClasses(c any, c2 any) bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(c, "isInterface")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(c2, "isInterface")))) {
		return false
	}

	if this.Hx_Field_curplatform != HxStringNull {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "platforms"), "push")
		Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
	}

	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(c, "isExtern")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_getField(c2, "isExtern")))) {
		Hx_Field_go_haxe_hxdynamic_setField(c, "isExtern", false)
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c2, "fields"))
		for {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var f2 any
			_ = f2
			f2 = _g1.Get(_g)
			_g++
			var found any
			_ = found
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				var _g3 HxArray[any]
				_ = _g3
				_g3 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"))
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g2
					if !(_hx_tmp_4 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g3)) {
						break
					}

					var f any
					_ = f
					f = _g3.Get(_g2)
					_g2++
					if this.VTable.Hx_Field_mergeFields(f, f2) {
						found = f
						break
					}
				}
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(found, nil))) {
				this.VTable.Hx_Field_newField(c, f2)
				var _hx_tmp_5 any
				_ = _hx_tmp_5
				_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "fields"), "push")
				Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any](f2).Dyn())
			} else {
				if this.Hx_Field_curplatform != HxStringNull {
					var _hx_tmp_6 any
					_ = _hx_tmp_6
					_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(found, "platforms"), "push")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
				}
			}
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 HxArray[any]
		_ = _g3
		_g3 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c2, "statics"))
		for {
			var _hx_tmp_7 int32
			_ = _hx_tmp_7
			_hx_tmp_7 = _g2
			if !(_hx_tmp_7 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g3)) {
				break
			}

			var f2 any
			_ = f2
			f2 = _g3.Get(_g2)
			_g2++
			var found any
			_ = found
			{
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[any]
				_ = _g1
				_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(c, "statics"))
				for {
					var _hx_tmp_8 int32
					_ = _hx_tmp_8
					_hx_tmp_8 = _g
					if !(_hx_tmp_8 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
						break
					}

					var f any
					_ = f
					f = _g1.Get(_g)
					_g++
					if this.VTable.Hx_Field_mergeFields(f, f2) {
						found = f
						break
					}
				}
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(found, nil))) {
				this.VTable.Hx_Field_newField(c, f2)
				var _hx_tmp_9 any
				_ = _hx_tmp_9
				_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(c, "statics"), "push")
				Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any](f2).Dyn())
			} else {
				if this.Hx_Field_curplatform != HxStringNull {
					var _hx_tmp_10 any
					_ = _hx_tmp_10
					_hx_tmp_10 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(found, "platforms"), "push")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_10, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
				}
			}
		}
	}

	return true
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeEnums(e any, e2 any) bool {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(e, "isExtern")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_getField(e2, "isExtern")))) {
		return false
	}

	if this.Hx_Field_curplatform != HxStringNull {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(e, "platforms"), "push")
		Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[any]
		_ = _g1
		_g1 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(e2, "constructors"))
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
				break
			}

			var c2 any
			_ = c2
			c2 = _g1.Get(_g)
			_g++
			var found any
			_ = found
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				var _g3 HxArray[any]
				_ = _g3
				_g3 = HxAnyToArray(Hx_Field_go_haxe_hxdynamic_getField(e, "constructors"))
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g2
					if !(_hx_tmp_3 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g3)) {
						break
					}

					var c any
					_ = c
					c = _g3.Get(_g2)
					_g2++
					if Hx_Field_haxe_rtti_typeapi_constructorEq(c, c2) {
						found = c
						break
					}
				}
			}

			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(found, nil))) {
				var _hx_tmp_4 any
				_ = _hx_tmp_4
				_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(e, "constructors"), "push")
				Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any](c2).Dyn())
			} else {
				if this.Hx_Field_curplatform != HxStringNull {
					var _hx_tmp_5 any
					_ = _hx_tmp_5
					_hx_tmp_5 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(found, "platforms"), "push")
					Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_5, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
				}
			}
		}
	}

	return true
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeTypedefs(t any, t2 any) bool {
	if this.Hx_Field_curplatform == HxStringNull {
		return false
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(t, "platforms"), "push")
	Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_curplatform
	func() func(_hx_param_key string, _hx_param_value Hx_Enum_haxe_rtti_ctype) {
		var hx_dyncast_51 any
		_ = hx_dyncast_51
		hx_dyncast_51 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(t, "types"), "Hx_Obj_haxe_imap")
		var _hx_tmp_1 *Hx_Obj_haxe_imap
		_ = _hx_tmp_1
		if hx_dyncast_51 != nil {
			_hx_tmp_1 = hx_dyncast_51.(*Hx_Obj_haxe_imap)
		} else {
			_hx_tmp_1 = nil
		}

		var _hx_closure_src_387 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_387
		_hx_closure_src_387 = _hx_tmp_1.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value Hx_Enum_haxe_rtti_ctype) {
			_hx_closure_src_387(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()(_hx_tmp_2, func() Hx_Enum_haxe_rtti_ctype {
		_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(t2, "type"))
		if _hx_d == nil {
			var _hx_z Hx_Enum_haxe_rtti_ctype
			return _hx_z
		}
		return _hx_d.(Hx_Enum_haxe_rtti_ctype)
	}())
	return true
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mergeAbstracts(a any, a2 any) bool {
	if this.Hx_Field_curplatform == HxStringNull {
		return false
	}

	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a, "to"), "length")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a2, "to"), "length")))) {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a, "from"), "length")
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = _hx_tmp_2
		_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a2, "from"), "length")))))
	}

	if _hx_tmp_0 {
		return false
	}

	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a, "to"), "length"))))
		for _g < _g1 {
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_4
			var _hx_tmp_5 Hx_Enum_haxe_rtti_ctype
			_ = _hx_tmp_5
			_hx_tmp_5 = func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getArrayIndex(Hx_Field_go_haxe_hxdynamic_getField(a, "to"), i), "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()
			if !Hx_Field_haxe_rtti_typeapi_typeEq(_hx_tmp_5, func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getArrayIndex(Hx_Field_go_haxe_hxdynamic_getField(a2, "to"), i), "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()) {
				return false
			}
		}
	}

	{
		var _g2 int32
		_ = _g2
		_g2 = 0
		var _g3 int32
		_ = _g3
		_g3 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a, "from"), "length"))))
		for _g2 < _g3 {
			var _hx_tmp_6 int32
			_ = _hx_tmp_6
			_hx_tmp_6 = _g2
			_g2 = (_g2 + 1)
			var i int32
			_ = i
			i = _hx_tmp_6
			var _hx_tmp_7 Hx_Enum_haxe_rtti_ctype
			_ = _hx_tmp_7
			_hx_tmp_7 = func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getArrayIndex(Hx_Field_go_haxe_hxdynamic_getField(a, "from"), i), "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()
			if !Hx_Field_haxe_rtti_typeapi_typeEq(_hx_tmp_7, func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getArrayIndex(Hx_Field_go_haxe_hxdynamic_getField(a2, "from"), i), "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()) {
				return false
			}
		}
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(a2, "impl"), nil))) {
		var _hx_tmp_8 any
		_ = _hx_tmp_8
		_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(a, "impl")
		this.VTable.Hx_Field_mergeClasses(_hx_tmp_8, Hx_Field_go_haxe_hxdynamic_getField(a2, "impl"))
	}

	var _hx_tmp_9 any
	_ = _hx_tmp_9
	_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(a, "platforms"), "push")
	Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any](this.Hx_Field_curplatform).Dyn())
	return true
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_merge(t Hx_Enum_haxe_rtti_typetree) {
	var inf any
	_ = inf
	inf = Hx_Field_haxe_rtti_typeapi_typeInfos(t)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(inf, "path"), "split")
	var pack HxArray[string]
	_ = pack
	pack = HxAnyToTypedArray[string](Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any](((any)("."))).Dyn()))
	var cur Hx_Typedef_haxe_rtti_typeroot
	_ = cur
	cur = this.Hx_Field_root
	var curpack HxArray[string]
	_ = curpack
	curpack = HxMakeArray[string]()
	_ = (Hx_Array_Pop[string](pack))
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < pack.Len()) {
				break
			}

			var p string
			_ = p
			p = pack.Get(_g)
			_g++
			var found bool
			_ = found
			found = false
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g1
					if !(_hx_tmp_2 < cur.Len()) {
						break
					}

					var pk Hx_Enum_haxe_rtti_typetree
					_ = pk
					pk = cur.Get(_g1)
					_g1++
					if pk.Hx_Field_enumIndex() == 0 {
						var pname string
						_ = pname
						pname = pk.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_name
						var subs HxArray[Hx_Enum_haxe_rtti_typetree]
						_ = subs
						subs = pk.(Hx_Enum_haxe_rtti_typetree_TPackage).Hx_Field_subs
						if pname == p {
							found = true
							cur = ((Hx_Typedef_haxe_rtti_typeroot)(subs))
							break
						}
					}
				}
			}

			_ = (Hx_Array_Push[string](curpack, p))
			if !found {
				var pk HxArray[Hx_Enum_haxe_rtti_typetree]
				_ = pk
				pk = HxMakeArray[Hx_Enum_haxe_rtti_typetree]()
				var _hx_tmp_3 Hx_Typedef_haxe_rtti_typeroot
				_ = _hx_tmp_3
				_hx_tmp_3 = cur
				var _hx_tmp_4 string
				_ = _hx_tmp_4
				_hx_tmp_4 = p
				_ = (Hx_Array_Push[Hx_Enum_haxe_rtti_typetree](_hx_tmp_3, Hx_Enum_haxe_rtti_typetree_TPackage{_hx_tmp_4, (Hx_Array_Join[string](curpack, ".")), pk}))
				cur = ((Hx_Typedef_haxe_rtti_typeroot)(pk))
			}
		}
	}

	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		for {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g1
			if !(_hx_tmp_5 < cur.Len()) {
				break
			}

			var ct Hx_Enum_haxe_rtti_typetree
			_ = ct
			ct = cur.Get(_g1)
			_g1++
			if ct.Hx_Field_enumIndex() == 0 {
				continue
			}

			var tinf any
			_ = tinf
			tinf = Hx_Field_haxe_rtti_typeapi_typeInfos(ct)
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "path")
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_6, Hx_Field_go_haxe_hxdynamic_getField(inf, "path")))) {
				var sameType bool
				_ = sameType
				sameType = true
				var _hx_tmp_7 bool
				_ = _hx_tmp_7
				_hx_tmp_7 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(tinf, "doc"), nil)))))), ((any)(false))))))
				if _hx_tmp_7 != ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(inf, "doc"), nil)))))), ((any)(false)))))) {
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(((any)(((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(Hx_Field_go_haxe_hxdynamic_getField(inf, "doc"), nil)))))), ((any)(false))))) {
						var _hx_tmp_8 any
						_ = _hx_tmp_8
						_hx_tmp_8 = inf
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_8, "doc", Hx_Field_go_haxe_hxdynamic_getField(tinf, "doc"))
					} else {
						var _hx_tmp_9 any
						_ = _hx_tmp_9
						_hx_tmp_9 = tinf
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_9, "doc", Hx_Field_go_haxe_hxdynamic_getField(inf, "doc"))
					}
				}

				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_go_haxe_hxdynamic_getField(tinf, "path"), ((any)("haxe._Int64.NativeInt64"))))) {
					continue
				}

				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				var _hx_tmp_11 bool
				_ = _hx_tmp_11
				var _hx_tmp_12 any
				_ = _hx_tmp_12
				_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "module")
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_12, Hx_Field_go_haxe_hxdynamic_getField(inf, "module")))) {
					var _hx_tmp_13 any
					_ = _hx_tmp_13
					_hx_tmp_13 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "doc")
					var _hx_tmp_14 any
					_ = _hx_tmp_14
					_hx_tmp_14 = _hx_tmp_13
					_hx_tmp_11 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_14, Hx_Field_go_haxe_hxdynamic_getField(inf, "doc")))))
				} else {
					_hx_tmp_11 = false
				}

				if _hx_tmp_11 {
					var _hx_tmp_15 any
					_ = _hx_tmp_15
					_hx_tmp_15 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "isPrivate")
					var _hx_tmp_16 any
					_ = _hx_tmp_16
					_hx_tmp_16 = _hx_tmp_15
					_hx_tmp_10 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(_hx_tmp_16, Hx_Field_go_haxe_hxdynamic_getField(inf, "isPrivate")))))
				} else {
					_hx_tmp_10 = false
				}

				if _hx_tmp_10 {
					switch ct.Hx_Field_enumIndex() {
					case 0:
						{
							sameType = false
						}

					case 1:
						{
							var c any
							_ = c
							c = ct.(Hx_Enum_haxe_rtti_typetree_TClassdecl).Hx_Field_c
							if t.Hx_Field_enumIndex() == 1 {
								var c2 any
								_ = c2
								c2 = t.(Hx_Enum_haxe_rtti_typetree_TClassdecl).Hx_Field_c
								if this.VTable.Hx_Field_mergeClasses(c, c2) {
									return
								}
							} else {
								sameType = false
							}
						}

					case 2:
						{
							var e any
							_ = e
							e = ct.(Hx_Enum_haxe_rtti_typetree_TEnumdecl).Hx_Field_e
							if t.Hx_Field_enumIndex() == 2 {
								var e2 any
								_ = e2
								e2 = t.(Hx_Enum_haxe_rtti_typetree_TEnumdecl).Hx_Field_e
								if this.VTable.Hx_Field_mergeEnums(e, e2) {
									return
								}
							} else {
								sameType = false
							}
						}

					case 3:
						{
							var td any
							_ = td
							td = ct.(Hx_Enum_haxe_rtti_typetree_TTypedecl).Hx_Field_t
							if t.Hx_Field_enumIndex() == 3 {
								var td2 any
								_ = td2
								td2 = t.(Hx_Enum_haxe_rtti_typetree_TTypedecl).Hx_Field_t
								if this.VTable.Hx_Field_mergeTypedefs(td, td2) {
									return
								}
							}
						}

					case 4:
						{
							var a any
							_ = a
							a = ct.(Hx_Enum_haxe_rtti_typetree_TAbstractdecl).Hx_Field_a
							if t.Hx_Field_enumIndex() == 4 {
								var a2 any
								_ = a2
								a2 = t.(Hx_Enum_haxe_rtti_typetree_TAbstractdecl).Hx_Field_a
								if this.VTable.Hx_Field_mergeAbstracts(a, a2) {
									return
								}
							} else {
								sameType = false
							}
						}

					}
				}

				var _hx_tmp_17 string
				_ = _hx_tmp_17
				var _hx_tmp_18 any
				_ = _hx_tmp_18
				_hx_tmp_18 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "module")
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_18, Hx_Field_go_haxe_hxdynamic_getField(inf, "module")))) {
					var _hx_tmp_19 any
					_ = _hx_tmp_19
					_hx_tmp_19 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(((any)("module ")), Hx_Field_go_haxe_hxdynamic_getField(inf, "module")), ((any)(" should be ")))
					var _hx_tmp_20 any
					_ = _hx_tmp_20
					_hx_tmp_20 = _hx_tmp_19
					_hx_tmp_17 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_20, Hx_Field_go_haxe_hxdynamic_getField(tinf, "module")))
				} else {
					var _hx_tmp_21 string
					_ = _hx_tmp_21
					var _hx_tmp_22 any
					_ = _hx_tmp_22
					_hx_tmp_22 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "doc")
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_22, Hx_Field_go_haxe_hxdynamic_getField(inf, "doc")))) {
						_hx_tmp_21 = "documentation is different"
					} else {
						var _hx_tmp_23 string
						_ = _hx_tmp_23
						var _hx_tmp_24 any
						_ = _hx_tmp_24
						_hx_tmp_24 = Hx_Field_go_haxe_hxdynamic_getField(tinf, "isPrivate")
						if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(_hx_tmp_24, Hx_Field_go_haxe_hxdynamic_getField(inf, "isPrivate")))) {
							_hx_tmp_23 = "private flag is different"
						} else {
							var _hx_tmp_25 string
							_ = _hx_tmp_25
							if !sameType {
								_hx_tmp_25 = "type kind is different"
							} else {
								_hx_tmp_25 = "could not merge definition"
							}

							_hx_tmp_23 = _hx_tmp_25
						}

						_hx_tmp_21 = _hx_tmp_23
					}

					_hx_tmp_17 = _hx_tmp_21
				}

				var msg string
				_ = msg
				msg = _hx_tmp_17
				var _hx_tmp_26 any
				_ = _hx_tmp_26
				_hx_tmp_26 = Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(((any)("Incompatibilities between ")), Hx_Field_go_haxe_hxdynamic_getField(tinf, "path")), ((any)(" in ")))
				var _hx_tmp_27 any
				_ = _hx_tmp_27
				_hx_tmp_27 = Hx_Field_go_haxe_hxdynamic_getField(Hx_Field_go_haxe_hxdynamic_getField(tinf, "platforms"), "join")
				panic(Hx_Field_haxe_exception_thrown(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(Hx_Field_go_haxe_hxdynamic_add(_hx_tmp_26, Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_27, HxMakeArray[any](((any)(","))).Dyn())), ((any)(" and "))), ((any)(this.Hx_Field_curplatform))), ((any)(" ("))), ((any)(msg))), ((any)(")")))))
			}
		}
	}

	_ = (Hx_Array_Push[Hx_Enum_haxe_rtti_typetree](cur, t))
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mkPath(p string) Hx_Typedef_haxe_rtti_path {
	return ((Hx_Typedef_haxe_rtti_path)(p))
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mkTypeParams(p string) Hx_Typedef_haxe_rtti_typeparams {
	var pl HxArray[string]
	_ = pl
	pl = (Hx_Field_go_haxe_hxstring_split(p, ":"))
	if pl.Get(0) == "" {
		return ((Hx_Typedef_haxe_rtti_typeparams)(HxMakeArray[string]()))
	}

	return ((Hx_Typedef_haxe_rtti_typeparams)(pl))
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_mkRights(r string) Hx_Enum_haxe_rtti_rights {
	switch r {
	case "dynamic":
		{
			return Hx_Enum_haxe_rtti_rights_RDynamic{}
		}

	case "inline":
		{
			return Hx_Enum_haxe_rtti_rights_RInline{}
		}

	case "method":
		{
			return Hx_Enum_haxe_rtti_rights_RMethod{}
		}

	case "null":
		{
			return Hx_Enum_haxe_rtti_rights_RNo{}
		}

	default:
		{
			return Hx_Enum_haxe_rtti_rights_RCall{r}
		}
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xerror(c *Hx_Obj_xml) any {
	var tmp string
	_ = tmp
	if c.Hx_Field_nodeType == Hx_Field_xml_Document {
		tmp = "Document"
	} else {
		if c.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
		}

		tmp = c.Hx_Field_nodeName
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid ", tmp)))))
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xroot(x *Hx_Obj_xml) {
	var _gc any
	_ = _gc
	_gc = x.VTable.Hx_Field_elements()
	for {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
		if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
			break
		}

		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
		var hx_dyncast_52 any
		_ = hx_dyncast_52
		hx_dyncast_52 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
		var _hx_tmp_2 *Hx_Obj_xml
		_ = _hx_tmp_2
		if hx_dyncast_52 != nil {
			_hx_tmp_2 = hx_dyncast_52.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
		} else {
			_hx_tmp_2 = nil
		}

		var c *Hx_Obj_xml
		_ = c
		c = _hx_tmp_2
		this.VTable.Hx_Field_merge(this.VTable.Hx_Field_processElement(c))
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_processElement(x *Hx_Obj_xml) Hx_Enum_haxe_rtti_typetree {
	if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
	}

	var c *Hx_Obj_xml
	_ = c
	c = x
	var _g string
	_ = _g
	if c.Hx_Field_nodeType == Hx_Field_xml_Document {
		_g = "Document"
	} else {
		if c.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
		}

		_g = c.Hx_Field_nodeName
	}

	switch _g {
	case "abstract":
		{
			return Hx_Enum_haxe_rtti_typetree_TAbstractdecl{this.VTable.Hx_Field_xabstract(c)}
		}

	case "class":
		{
			return Hx_Enum_haxe_rtti_typetree_TClassdecl{this.VTable.Hx_Field_xclass(c)}
		}

	case "enum":
		{
			return Hx_Enum_haxe_rtti_typetree_TEnumdecl{this.VTable.Hx_Field_xenum(c)}
		}

	case "typedef":
		{
			return Hx_Enum_haxe_rtti_typetree_TTypedecl{this.VTable.Hx_Field_xtypedef(c)}
		}

	default:
		{
			return func() Hx_Enum_haxe_rtti_typetree {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_xerror(c))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_typetree
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_typetree)
			}()
		}
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xmeta(x *Hx_Obj_xml) HxArray[any] {
	var ml HxArray[any]
	_ = ml
	ml = HxMakeArray[any]().Dyn()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_xml]
		_ = _g1
		_g1 = Hx_Field_haxe_xml__access_nodelistaccess_impl__resolve(x, "m")
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var m *Hx_Obj_xml
			_ = m
			m = _g1.Get(_g)
			_g++
			var pl HxArray[string]
			_ = pl
			pl = HxMakeArray[string]()
			{
				var _g2 int32
				_ = _g2
				_g2 = 0
				var _g3 HxArray[*Hx_Obj_xml]
				_ = _g3
				_g3 = Hx_Field_haxe_xml__access_nodelistaccess_impl__resolve(m, "e")
				for {
					var _hx_tmp_1 int32
					_ = _hx_tmp_1
					_hx_tmp_1 = _g2
					if !(_hx_tmp_1 < _g3.Len()) {
						break
					}

					var p *Hx_Obj_xml
					_ = p
					p = _g3.Get(_g2)
					_g2++
					var _hx_tmp_2 HxArray[string]
					_ = _hx_tmp_2
					_hx_tmp_2 = pl
					_ = (Hx_Array_Push[string](_hx_tmp_2, Hx_Field_haxe_xml__access_access_impl__get_innerHTML(p)))
				}
			}

			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = HxMakeArrayView[any](ml)
			_ = (Hx_Array_Push[any](_hx_tmp_3, any(map[string]any{"name": ((any)(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(m, "n"))), "params": ((any)(pl))})))
		}
	}

	return ml
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xoverloads(x *Hx_Obj_xml) HxArray[any] {
	var l HxArray[any]
	_ = l
	l = HxMakeArray[any]().Dyn()
	{
		var _gm any
		_ = _gm
		_gm = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gm, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gm, "next")
			var hx_dyncast_53 any
			_ = hx_dyncast_53
			hx_dyncast_53 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_53 != nil {
				_hx_tmp_2 = hx_dyncast_53.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var m *Hx_Obj_xml
			_ = m
			m = _hx_tmp_2
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = HxMakeArrayView[any](l)
			var _hx_tmp_4 *Hx_Obj_xml
			_ = _hx_tmp_4
			_hx_tmp_4 = m
			_ = (Hx_Array_Push[any](_hx_tmp_3, this.VTable.Hx_Field_xclassfield(_hx_tmp_4, nil)))
		}
	}

	return l
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xpath(x *Hx_Obj_xml) any {
	var path Hx_Typedef_haxe_rtti_path
	_ = path
	path = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))
	var params HxArray[Hx_Enum_haxe_rtti_ctype]
	_ = params
	params = HxMakeArray[Hx_Enum_haxe_rtti_ctype]()
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_54 any
			_ = hx_dyncast_54
			hx_dyncast_54 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_54 != nil {
				_hx_tmp_2 = hx_dyncast_54.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			var _hx_tmp_3 HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = _hx_tmp_3
			_hx_tmp_3 = params
			_ = (Hx_Array_Push[Hx_Enum_haxe_rtti_ctype](_hx_tmp_3, this.VTable.Hx_Field_xtype(c)))
		}
	}

	return any(map[string]any{"path": ((any)(path)), "params": ((any)(params))})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xclass(x *Hx_Obj_xml) any {
	var csuper any
	_ = csuper
	var doc HxNullable[string]
	_ = doc
	doc = HxNullable[string]{Value: HxStringNull}

	var tdynamic Hx_Enum_haxe_rtti_ctype
	_ = tdynamic
	var interfaces HxArray[any]
	_ = interfaces
	interfaces = HxMakeArray[any]().Dyn()
	var fields HxArray[any]
	_ = fields
	fields = HxMakeArray[any]().Dyn()
	var statics HxArray[any]
	_ = statics
	statics = HxMakeArray[any]().Dyn()
	var meta HxArray[any]
	_ = meta
	meta = HxMakeArray[any]().Dyn()
	var isInterface bool
	_ = isInterface
	isInterface = x.VTable.Hx_Field_exists("interface")
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_55 any
			_ = hx_dyncast_55
			hx_dyncast_55 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_55 != nil {
				_hx_tmp_2 = hx_dyncast_55.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			{
				var _g string
				_ = _g
				if c.Hx_Field_nodeType == Hx_Field_xml_Document {
					_g = "Document"
				} else {
					if c.Hx_Field_nodeType != Hx_Field_xml_Element {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					_g = c.Hx_Field_nodeName
				}

				switch _g {
				case "extends":
					{
						if isInterface {
							var _hx_tmp_3 HxArray[any]
							_ = _hx_tmp_3
							_hx_tmp_3 = HxMakeArrayView[any](interfaces)
							_ = (Hx_Array_Push[any](_hx_tmp_3, this.VTable.Hx_Field_xpath(c)))
						} else {
							csuper = this.VTable.Hx_Field_xpath(c)
						}
					}

				case "haxe_doc":
					{
						doc = HxNullableString(Hx_Field_haxe_xml__access_access_impl__get_innerData(c))
					}

				case "haxe_dynamic":
					{
						var x *Hx_Obj_xml
						_ = x
						x = c.VTable.Hx_Field_firstElement()
						if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
						}

						tdynamic = this.VTable.Hx_Field_xtype(x)
					}

				case "implements":
					{
						var _hx_tmp_4 HxArray[any]
						_ = _hx_tmp_4
						_hx_tmp_4 = HxMakeArrayView[any](interfaces)
						_ = (Hx_Array_Push[any](_hx_tmp_4, this.VTable.Hx_Field_xpath(c)))
					}

				case "meta":
					{
						meta = this.VTable.Hx_Field_xmeta(c)
					}

				default:
					{
						if c.VTable.Hx_Field_exists("static") {
							var _hx_tmp_5 HxArray[any]
							_ = _hx_tmp_5
							_hx_tmp_5 = HxMakeArrayView[any](statics)
							var _hx_tmp_6 *Hx_Obj_xml
							_ = _hx_tmp_6
							_hx_tmp_6 = c
							_ = (Hx_Array_Push[any](_hx_tmp_5, this.VTable.Hx_Field_xclassfield(_hx_tmp_6, nil)))
						} else {
							var _hx_tmp_7 HxArray[any]
							_ = _hx_tmp_7
							_hx_tmp_7 = HxMakeArrayView[any](fields)
							var _hx_tmp_8 *Hx_Obj_xml
							_ = _hx_tmp_8
							_hx_tmp_8 = c
							_ = (Hx_Array_Push[any](_hx_tmp_7, this.VTable.Hx_Field_xclassfield(_hx_tmp_8, nil)))
						}
					}
				}
			}
		}
	}

	var _hx_tmp_9 HxNullable[string]
	_ = _hx_tmp_9
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "file") {
		_hx_tmp_9 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "file"))
	} else {
		_hx_tmp_9 = HxNullable[string]{Value: HxStringNull}
	}

	var hx_nullable_f185 HxNullable[string]
	_ = hx_nullable_f185
	hx_nullable_f185 = _hx_tmp_9
	var _hx_tmp_10 any
	_ = _hx_tmp_10
	if (hx_nullable_f185).Valid == true {
		_hx_tmp_10 = (hx_nullable_f185).Value
	} else {
		_hx_tmp_10 = nil
	}

	var _hx_tmp_11 any
	_ = _hx_tmp_11
	_hx_tmp_11 = ((any)(this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))))
	var _hx_tmp_12 Hx_Typedef_haxe_rtti_path
	_ = _hx_tmp_12
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "module") {
		_hx_tmp_12 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "module"))
	} else {
		_hx_tmp_12 = ((Hx_Typedef_haxe_rtti_path)(HxStringNull))
	}

	var hx_nullable_f186 HxNullable[string]
	_ = hx_nullable_f186
	hx_nullable_f186 = doc
	var _hx_tmp_13 any
	_ = _hx_tmp_13
	if (hx_nullable_f186).Valid == true {
		_hx_tmp_13 = (hx_nullable_f186).Value
	} else {
		_hx_tmp_13 = nil
	}

	var _hx_tmp_14 any
	_ = _hx_tmp_14
	_hx_tmp_14 = ((any)(x.VTable.Hx_Field_exists("private")))
	var _hx_tmp_15 any
	_ = _hx_tmp_15
	_hx_tmp_15 = ((any)(x.VTable.Hx_Field_exists("extern")))
	var _hx_tmp_16 any
	_ = _hx_tmp_16
	_hx_tmp_16 = ((any)(x.VTable.Hx_Field_exists("final")))
	var _hx_tmp_17 any
	_ = _hx_tmp_17
	_hx_tmp_17 = ((any)(isInterface))
	var _hx_tmp_18 any
	_ = _hx_tmp_18
	_hx_tmp_18 = ((any)(this.VTable.Hx_Field_mkTypeParams(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "params"))))
	var _hx_tmp_19 any
	_ = _hx_tmp_19
	_hx_tmp_19 = csuper
	var _hx_tmp_20 any
	_ = _hx_tmp_20
	_hx_tmp_20 = ((any)(interfaces))
	var _hx_tmp_21 any
	_ = _hx_tmp_21
	_hx_tmp_21 = ((any)(fields))
	var _hx_tmp_22 any
	_ = _hx_tmp_22
	_hx_tmp_22 = ((any)(statics))
	var _hx_tmp_23 any
	_ = _hx_tmp_23
	_hx_tmp_23 = ((any)(tdynamic))
	return any(map[string]any{"file": _hx_tmp_10, "path": _hx_tmp_11, "module": ((any)(_hx_tmp_12)), "doc": _hx_tmp_13, "isPrivate": _hx_tmp_14, "isExtern": _hx_tmp_15, "isFinal": _hx_tmp_16, "isInterface": _hx_tmp_17, "params": _hx_tmp_18, "superClass": _hx_tmp_19, "interfaces": _hx_tmp_20, "fields": _hx_tmp_21, "statics": _hx_tmp_22, "tdynamic": _hx_tmp_23, "platforms": ((any)(this.VTable.Hx_Field_defplat())), "meta": ((any)(meta))})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xclassfield(x *Hx_Obj_xml, defPublic__isOptional any) any {
	defPublic := HxDefault[HxNullable[bool]]()
	_ = defPublic

	if defPublic__isOptional != nil {
		defPublic = HxConvert[HxNullable[bool]](defPublic__isOptional)
	} else {

	}
	{
		if (defPublic).Valid == false {
			defPublic = HxNullable[bool]{Value: false, Valid: true}
		}

		var e any
		_ = e
		e = x.VTable.Hx_Field_elements()
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(e, "next")
		var hx_dyncast_56 any
		_ = hx_dyncast_56
		hx_dyncast_56 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
		var _hx_tmp_1 *Hx_Obj_xml
		_ = _hx_tmp_1
		if hx_dyncast_56 != nil {
			_hx_tmp_1 = hx_dyncast_56.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
		} else {
			_hx_tmp_1 = nil
		}

		var t Hx_Enum_haxe_rtti_ctype
		_ = t
		t = this.VTable.Hx_Field_xtype(_hx_tmp_1)
		var doc HxNullable[string]
		_ = doc
		doc = HxNullable[string]{Value: HxStringNull}

		var meta HxArray[any]
		_ = meta
		meta = HxMakeArray[any]().Dyn()
		var overloads HxArray[any]
		_ = overloads
		{
			var _gc any
			_ = _gc
			_gc = e
			for {
				var _hx_tmp_2 any
				_ = _hx_tmp_2
				_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
				if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()))) {
					break
				}

				var _hx_tmp_3 any
				_ = _hx_tmp_3
				_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
				var hx_dyncast_57 any
				_ = hx_dyncast_57
				hx_dyncast_57 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
				var _hx_tmp_4 *Hx_Obj_xml
				_ = _hx_tmp_4
				if hx_dyncast_57 != nil {
					_hx_tmp_4 = hx_dyncast_57.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
				} else {
					_hx_tmp_4 = nil
				}

				var c *Hx_Obj_xml
				_ = c
				c = _hx_tmp_4
				{
					var _g string
					_ = _g
					if c.Hx_Field_nodeType == Hx_Field_xml_Document {
						_g = "Document"
					} else {
						if c.Hx_Field_nodeType != Hx_Field_xml_Element {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
						}

						_g = c.Hx_Field_nodeName
					}

					switch _g {
					case "haxe_doc":
						{
							doc = HxNullableString(Hx_Field_haxe_xml__access_access_impl__get_innerData(c))
						}

					case "meta":
						{
							meta = this.VTable.Hx_Field_xmeta(c)
						}

					case "overloads":
						{
							overloads = this.VTable.Hx_Field_xoverloads(c)
						}

					default:
						{
							this.VTable.Hx_Field_xerror(c)
						}
					}
				}
			}
		}

		var tmp string
		_ = tmp
		if x.Hx_Field_nodeType == Hx_Field_xml_Document {
			tmp = "Document"
		} else {
			if x.Hx_Field_nodeType != Hx_Field_xml_Element {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
			}

			tmp = x.Hx_Field_nodeName
		}

		var _hx_tmp_5 any
		_ = _hx_tmp_5
		_hx_tmp_5 = ((any)(tmp))
		var _hx_tmp_6 any
		_ = _hx_tmp_6
		_hx_tmp_6 = ((any)(t))
		var _hx_tmp_8 bool
		_ = _hx_tmp_8
		if x.VTable.Hx_Field_exists("public") {
			_hx_tmp_8 = true
		} else {
			_hx_tmp_8 = (defPublic).Value
		}

		var _hx_tmp_7 any
		_ = _hx_tmp_7
		_hx_tmp_7 = ((any)(_hx_tmp_8))
		var _hx_tmp_9 any
		_ = _hx_tmp_9
		_hx_tmp_9 = ((any)(x.VTable.Hx_Field_exists("final")))
		var _hx_tmp_10 any
		_ = _hx_tmp_10
		_hx_tmp_10 = ((any)(x.VTable.Hx_Field_exists("override")))
		var _hx_tmp_11 HxNullable[int32]
		_ = _hx_tmp_11
		if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "line") {
			_hx_tmp_11 = Hx_Field_std_parseInt(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "line"))
		} else {
			_hx_tmp_11 = HxNullable[int32]{}
		}

		var hx_nullable_f187 HxNullable[int32]
		_ = hx_nullable_f187
		hx_nullable_f187 = _hx_tmp_11
		var _hx_tmp_12 any
		_ = _hx_tmp_12
		if (hx_nullable_f187).Valid == true {
			_hx_tmp_12 = (hx_nullable_f187).Value
		} else {
			_hx_tmp_12 = nil
		}

		var hx_nullable_f188 HxNullable[string]
		_ = hx_nullable_f188
		hx_nullable_f188 = doc
		var _hx_tmp_13 any
		_ = _hx_tmp_13
		if (hx_nullable_f188).Valid == true {
			_hx_tmp_13 = (hx_nullable_f188).Value
		} else {
			_hx_tmp_13 = nil
		}

		var _hx_tmp_14 Hx_Enum_haxe_rtti_rights
		_ = _hx_tmp_14
		if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "get") {
			_hx_tmp_14 = this.VTable.Hx_Field_mkRights(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "get"))
		} else {
			_hx_tmp_14 = Hx_Enum_haxe_rtti_rights_RNormal{}
		}

		var _hx_tmp_15 Hx_Enum_haxe_rtti_rights
		_ = _hx_tmp_15
		if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "set") {
			_hx_tmp_15 = this.VTable.Hx_Field_mkRights(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "set"))
		} else {
			_hx_tmp_15 = Hx_Enum_haxe_rtti_rights_RNormal{}
		}

		var _hx_tmp_16 Hx_Typedef_haxe_rtti_typeparams
		_ = _hx_tmp_16
		if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "params") {
			_hx_tmp_16 = this.VTable.Hx_Field_mkTypeParams(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "params"))
		} else {
			_hx_tmp_16 = ((Hx_Typedef_haxe_rtti_typeparams)(HxMakeArray[string]()))
		}

		var _hx_tmp_17 any
		_ = _hx_tmp_17
		_hx_tmp_17 = ((any)(this.VTable.Hx_Field_defplat()))
		var _hx_tmp_18 any
		_ = _hx_tmp_18
		_hx_tmp_18 = ((any)(meta))
		var _hx_tmp_19 any
		_ = _hx_tmp_19
		_hx_tmp_19 = ((any)(overloads))
		var _hx_tmp_20 HxNullable[string]
		_ = _hx_tmp_20
		if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "expr") {
			_hx_tmp_20 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "expr"))
		} else {
			_hx_tmp_20 = HxNullable[string]{Value: HxStringNull}
		}

		var hx_nullable_f189 HxNullable[string]
		_ = hx_nullable_f189
		hx_nullable_f189 = _hx_tmp_20
		var _hx_tmp_21 any
		_ = _hx_tmp_21
		if (hx_nullable_f189).Valid == true {
			_hx_tmp_21 = (hx_nullable_f189).Value
		} else {
			_hx_tmp_21 = nil
		}

		return any(map[string]any{"name": _hx_tmp_5, "type": _hx_tmp_6, "isPublic": _hx_tmp_7, "isFinal": _hx_tmp_9, "isOverride": _hx_tmp_10, "line": _hx_tmp_12, "doc": _hx_tmp_13, "get": ((any)(_hx_tmp_14)), "set": ((any)(_hx_tmp_15)), "params": ((any)(_hx_tmp_16)), "platforms": _hx_tmp_17, "meta": _hx_tmp_18, "overloads": _hx_tmp_19, "expr": _hx_tmp_21})
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xenum(x *Hx_Obj_xml) any {
	var cl HxArray[any]
	_ = cl
	cl = HxMakeArray[any]().Dyn()
	var doc HxNullable[string]
	_ = doc
	doc = HxNullable[string]{Value: HxStringNull}

	var meta HxArray[any]
	_ = meta
	meta = HxMakeArray[any]().Dyn()
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_58 any
			_ = hx_dyncast_58
			hx_dyncast_58 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_58 != nil {
				_hx_tmp_2 = hx_dyncast_58.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			var tmp string
			_ = tmp
			if c.Hx_Field_nodeType == Hx_Field_xml_Document {
				tmp = "Document"
			} else {
				if c.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
				}

				tmp = c.Hx_Field_nodeName
			}

			if tmp == "haxe_doc" {
				doc = HxNullableString(Hx_Field_haxe_xml__access_access_impl__get_innerData(c))
			} else {
				var tmp_tmp_1 string
				_ = tmp_tmp_1
				if c.Hx_Field_nodeType == Hx_Field_xml_Document {
					tmp_tmp_1 = "Document"
				} else {
					if c.Hx_Field_nodeType != Hx_Field_xml_Element {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					tmp_tmp_1 = c.Hx_Field_nodeName
				}

				if tmp_tmp_1 == "meta" {
					meta = this.VTable.Hx_Field_xmeta(c)
				} else {
					var _hx_tmp_3 HxArray[any]
					_ = _hx_tmp_3
					_hx_tmp_3 = HxMakeArrayView[any](cl)
					_ = (Hx_Array_Push[any](_hx_tmp_3, this.VTable.Hx_Field_xenumfield(c)))
				}
			}
		}
	}

	var _hx_tmp_4 HxNullable[string]
	_ = _hx_tmp_4
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "file") {
		_hx_tmp_4 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "file"))
	} else {
		_hx_tmp_4 = HxNullable[string]{Value: HxStringNull}
	}

	var hx_nullable_f190 HxNullable[string]
	_ = hx_nullable_f190
	hx_nullable_f190 = _hx_tmp_4
	var _hx_tmp_5 any
	_ = _hx_tmp_5
	if (hx_nullable_f190).Valid == true {
		_hx_tmp_5 = (hx_nullable_f190).Value
	} else {
		_hx_tmp_5 = nil
	}

	var _hx_tmp_6 any
	_ = _hx_tmp_6
	_hx_tmp_6 = ((any)(this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))))
	var _hx_tmp_7 Hx_Typedef_haxe_rtti_path
	_ = _hx_tmp_7
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "module") {
		_hx_tmp_7 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "module"))
	} else {
		_hx_tmp_7 = ((Hx_Typedef_haxe_rtti_path)(HxStringNull))
	}

	var hx_nullable_f191 HxNullable[string]
	_ = hx_nullable_f191
	hx_nullable_f191 = doc
	var _hx_tmp_8 any
	_ = _hx_tmp_8
	if (hx_nullable_f191).Valid == true {
		_hx_tmp_8 = (hx_nullable_f191).Value
	} else {
		_hx_tmp_8 = nil
	}

	var _hx_tmp_9 any
	_ = _hx_tmp_9
	_hx_tmp_9 = ((any)(x.VTable.Hx_Field_exists("private")))
	var _hx_tmp_10 any
	_ = _hx_tmp_10
	_hx_tmp_10 = ((any)(x.VTable.Hx_Field_exists("extern")))
	var _hx_tmp_11 any
	_ = _hx_tmp_11
	_hx_tmp_11 = ((any)(this.VTable.Hx_Field_mkTypeParams(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "params"))))
	var _hx_tmp_12 any
	_ = _hx_tmp_12
	_hx_tmp_12 = ((any)(cl))
	return any(map[string]any{"file": _hx_tmp_5, "path": _hx_tmp_6, "module": ((any)(_hx_tmp_7)), "doc": _hx_tmp_8, "isPrivate": _hx_tmp_9, "isExtern": _hx_tmp_10, "params": _hx_tmp_11, "constructors": _hx_tmp_12, "platforms": ((any)(this.VTable.Hx_Field_defplat())), "meta": ((any)(meta))})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xenumfield(x *Hx_Obj_xml) any {
	var args HxArray[any]
	_ = args
	var docElements any
	_ = docElements
	docElements = x.VTable.Hx_Field_elementsNamed("haxe_doc")
	var _hx_tmp_0 *Hx_Obj_xml
	_ = _hx_tmp_0
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(docElements, "hasNext")
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn())))) {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxdynamic_getField(docElements, "next")
		var hx_dyncast_61 any
		_ = hx_dyncast_61
		hx_dyncast_61 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_2, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
		var _hx_tmp_3 *Hx_Obj_xml
		_ = _hx_tmp_3
		if hx_dyncast_61 != nil {
			_hx_tmp_3 = hx_dyncast_61.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
		} else {
			_hx_tmp_3 = nil
		}

		_hx_tmp_0 = _hx_tmp_3
	} else {
		_hx_tmp_0 = nil
	}

	var xdoc *Hx_Obj_xml
	_ = xdoc
	xdoc = _hx_tmp_0
	var _hx_tmp_4 HxArray[any]
	_ = _hx_tmp_4
	if Hx_Field_haxe_xml__access_hasnodeaccess_impl__resolve(x, "meta") {
		_hx_tmp_4 = this.VTable.Hx_Field_xmeta(Hx_Field_haxe_xml__access_nodeaccess_impl__resolve(x, "meta"))
	} else {
		_hx_tmp_4 = HxMakeArray[any]().Dyn()
	}

	var meta HxArray[any]
	_ = meta
	meta = _hx_tmp_4
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "a") {
		var names HxArray[string]
		_ = names
		names = (Hx_Field_go_haxe_hxstring_split(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "a"), ":"))
		var elts any
		_ = elts
		elts = x.VTable.Hx_Field_elements()
		args = HxMakeArray[any]().Dyn()
		{
			var _g int32
			_ = _g
			_g = 0
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g
				if !(_hx_tmp_5 < names.Len()) {
					break
				}

				var c string
				_ = c
				c = names.Get(_g)
				_g++
				var opt bool
				_ = opt
				opt = false
				if (Hx_Field_go_haxe_hxstring_charAt(c, 0)) == "?" {
					opt = true
					var _hx_tmp_6 string
					_ = _hx_tmp_6
					_hx_tmp_6 = c
					c = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_6, 1, HxNullable[int32]{}))
				}

				var _hx_tmp_7 HxArray[any]
				_ = _hx_tmp_7
				_hx_tmp_7 = HxMakeArrayView[any](args)
				var _hx_tmp_8 any
				_ = _hx_tmp_8
				_hx_tmp_8 = ((any)(c))
				var _hx_tmp_9 any
				_ = _hx_tmp_9
				_hx_tmp_9 = ((any)(opt))
				var _hx_tmp_10 any
				_ = _hx_tmp_10
				_hx_tmp_10 = Hx_Field_go_haxe_hxdynamic_getField(elts, "next")
				var hx_dyncast_60 any
				_ = hx_dyncast_60
				hx_dyncast_60 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_10, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
				var _hx_tmp_11 *Hx_Obj_xml
				_ = _hx_tmp_11
				if hx_dyncast_60 != nil {
					_hx_tmp_11 = hx_dyncast_60.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
				} else {
					_hx_tmp_11 = nil
				}

				_ = (Hx_Array_Push[any](_hx_tmp_7, any(map[string]any{"name": _hx_tmp_8, "opt": _hx_tmp_9, "t": ((any)(this.VTable.Hx_Field_xtype(_hx_tmp_11)))})))
			}
		}
	}

	var tmp string
	_ = tmp
	if x.Hx_Field_nodeType == Hx_Field_xml_Document {
		tmp = "Document"
	} else {
		if x.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
		}

		tmp = x.Hx_Field_nodeName
	}

	var tmp1 string
	_ = tmp1
	if xdoc == nil {
		tmp1 = HxStringNull
	} else {
		if (xdoc.Hx_Field_nodeType != Hx_Field_xml_Document) && (xdoc.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(xdoc.Hx_Field_nodeType))))))
		}

		tmp1 = Hx_Field_haxe_xml__access_access_impl__get_innerData(xdoc)
	}

	var _hx_tmp_12 any
	_ = _hx_tmp_12
	_hx_tmp_12 = ((any)(tmp))
	var _hx_tmp_13 any
	_ = _hx_tmp_13
	_hx_tmp_13 = ((any)(args))
	var _hx_tmp_14 any
	_ = _hx_tmp_14
	_hx_tmp_14 = ((any)(tmp1))
	var _hx_tmp_15 any
	_ = _hx_tmp_15
	_hx_tmp_15 = ((any)(meta))
	return any(map[string]any{"name": _hx_tmp_12, "args": _hx_tmp_13, "doc": _hx_tmp_14, "meta": _hx_tmp_15, "platforms": ((any)(this.VTable.Hx_Field_defplat()))})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xabstract(x *Hx_Obj_xml) any {
	var doc HxNullable[string]
	_ = doc
	doc = HxNullable[string]{Value: HxStringNull}

	var impl any
	_ = impl
	var athis Hx_Enum_haxe_rtti_ctype
	_ = athis
	var meta HxArray[any]
	_ = meta
	meta = HxMakeArray[any]().Dyn()
	var to HxArray[any]
	_ = to
	to = HxMakeArray[any]().Dyn()
	var from HxArray[any]
	_ = from
	from = HxMakeArray[any]().Dyn()
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_62 any
			_ = hx_dyncast_62
			hx_dyncast_62 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_62 != nil {
				_hx_tmp_2 = hx_dyncast_62.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			{
				var _g string
				_ = _g
				if c.Hx_Field_nodeType == Hx_Field_xml_Document {
					_g = "Document"
				} else {
					if c.Hx_Field_nodeType != Hx_Field_xml_Element {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					_g = c.Hx_Field_nodeName
				}

				switch _g {
				case "from":
					{
						var _gt any
						_ = _gt
						_gt = c.VTable.Hx_Field_elements()
						for {
							var _hx_tmp_3 any
							_ = _hx_tmp_3
							_hx_tmp_3 = Hx_Field_go_haxe_hxdynamic_getField(_gt, "hasNext")
							if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_3, HxMakeArray[any]().Dyn()))) {
								break
							}

							var _hx_tmp_4 any
							_ = _hx_tmp_4
							_hx_tmp_4 = Hx_Field_go_haxe_hxdynamic_getField(_gt, "next")
							var hx_dyncast_63 any
							_ = hx_dyncast_63
							hx_dyncast_63 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_4, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
							var _hx_tmp_5 *Hx_Obj_xml
							_ = _hx_tmp_5
							if hx_dyncast_63 != nil {
								_hx_tmp_5 = hx_dyncast_63.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
							} else {
								_hx_tmp_5 = nil
							}

							var t *Hx_Obj_xml
							_ = t
							t = _hx_tmp_5
							var x *Hx_Obj_xml
							_ = x
							x = t.VTable.Hx_Field_firstElement()
							if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
							}

							var _hx_tmp_6 HxArray[any]
							_ = _hx_tmp_6
							_hx_tmp_6 = HxMakeArrayView[any](from)
							var _hx_tmp_7 any
							_ = _hx_tmp_7
							_hx_tmp_7 = ((any)(this.VTable.Hx_Field_xtype(x)))
							var _hx_tmp_8 HxNullable[string]
							_ = _hx_tmp_8
							if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(t, "field") {
								_hx_tmp_8 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(t, "field"))
							} else {
								_hx_tmp_8 = HxNullable[string]{Value: HxStringNull}
							}

							var hx_nullable_f192 HxNullable[string]
							_ = hx_nullable_f192
							hx_nullable_f192 = _hx_tmp_8
							var _hx_tmp_9 any
							_ = _hx_tmp_9
							if (hx_nullable_f192).Valid == true {
								_hx_tmp_9 = (hx_nullable_f192).Value
							} else {
								_hx_tmp_9 = nil
							}

							_ = (Hx_Array_Push[any](_hx_tmp_6, any(map[string]any{"t": _hx_tmp_7, "field": _hx_tmp_9})))
						}
					}

				case "haxe_doc":
					{
						doc = HxNullableString(Hx_Field_haxe_xml__access_access_impl__get_innerData(c))
					}

				case "impl":
					{
						impl = this.VTable.Hx_Field_xclass(Hx_Field_haxe_xml__access_nodeaccess_impl__resolve(c, "class"))
					}

				case "meta":
					{
						meta = this.VTable.Hx_Field_xmeta(c)
					}

				case "this":
					{
						var x *Hx_Obj_xml
						_ = x
						x = c.VTable.Hx_Field_firstElement()
						if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
						}

						athis = this.VTable.Hx_Field_xtype(x)
					}

				case "to":
					{
						var _gt any
						_ = _gt
						_gt = c.VTable.Hx_Field_elements()
						for {
							var _hx_tmp_10 any
							_ = _hx_tmp_10
							_hx_tmp_10 = Hx_Field_go_haxe_hxdynamic_getField(_gt, "hasNext")
							if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_10, HxMakeArray[any]().Dyn()))) {
								break
							}

							var _hx_tmp_11 any
							_ = _hx_tmp_11
							_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_getField(_gt, "next")
							var hx_dyncast_64 any
							_ = hx_dyncast_64
							hx_dyncast_64 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_11, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
							var _hx_tmp_12 *Hx_Obj_xml
							_ = _hx_tmp_12
							if hx_dyncast_64 != nil {
								_hx_tmp_12 = hx_dyncast_64.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
							} else {
								_hx_tmp_12 = nil
							}

							var t *Hx_Obj_xml
							_ = t
							t = _hx_tmp_12
							var x *Hx_Obj_xml
							_ = x
							x = t.VTable.Hx_Field_firstElement()
							if (x.Hx_Field_nodeType != Hx_Field_xml_Document) && (x.Hx_Field_nodeType != Hx_Field_xml_Element) {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
							}

							var _hx_tmp_13 HxArray[any]
							_ = _hx_tmp_13
							_hx_tmp_13 = HxMakeArrayView[any](to)
							var _hx_tmp_14 any
							_ = _hx_tmp_14
							_hx_tmp_14 = ((any)(this.VTable.Hx_Field_xtype(x)))
							var _hx_tmp_15 HxNullable[string]
							_ = _hx_tmp_15
							if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(t, "field") {
								_hx_tmp_15 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(t, "field"))
							} else {
								_hx_tmp_15 = HxNullable[string]{Value: HxStringNull}
							}

							var hx_nullable_f193 HxNullable[string]
							_ = hx_nullable_f193
							hx_nullable_f193 = _hx_tmp_15
							var _hx_tmp_16 any
							_ = _hx_tmp_16
							if (hx_nullable_f193).Valid == true {
								_hx_tmp_16 = (hx_nullable_f193).Value
							} else {
								_hx_tmp_16 = nil
							}

							_ = (Hx_Array_Push[any](_hx_tmp_13, any(map[string]any{"t": _hx_tmp_14, "field": _hx_tmp_16})))
						}
					}

				default:
					{
						this.VTable.Hx_Field_xerror(c)
					}
				}
			}
		}
	}

	var _hx_tmp_17 HxNullable[string]
	_ = _hx_tmp_17
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "file") {
		_hx_tmp_17 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "file"))
	} else {
		_hx_tmp_17 = HxNullable[string]{Value: HxStringNull}
	}

	var hx_nullable_f194 HxNullable[string]
	_ = hx_nullable_f194
	hx_nullable_f194 = _hx_tmp_17
	var _hx_tmp_18 any
	_ = _hx_tmp_18
	if (hx_nullable_f194).Valid == true {
		_hx_tmp_18 = (hx_nullable_f194).Value
	} else {
		_hx_tmp_18 = nil
	}

	var _hx_tmp_19 any
	_ = _hx_tmp_19
	_hx_tmp_19 = ((any)(this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))))
	var _hx_tmp_20 Hx_Typedef_haxe_rtti_path
	_ = _hx_tmp_20
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "module") {
		_hx_tmp_20 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "module"))
	} else {
		_hx_tmp_20 = ((Hx_Typedef_haxe_rtti_path)(HxStringNull))
	}

	var hx_nullable_f195 HxNullable[string]
	_ = hx_nullable_f195
	hx_nullable_f195 = doc
	var _hx_tmp_21 any
	_ = _hx_tmp_21
	if (hx_nullable_f195).Valid == true {
		_hx_tmp_21 = (hx_nullable_f195).Value
	} else {
		_hx_tmp_21 = nil
	}

	var _hx_tmp_22 any
	_ = _hx_tmp_22
	_hx_tmp_22 = ((any)(x.VTable.Hx_Field_exists("private")))
	var _hx_tmp_23 any
	_ = _hx_tmp_23
	_hx_tmp_23 = ((any)(this.VTable.Hx_Field_mkTypeParams(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "params"))))
	return any(map[string]any{"file": _hx_tmp_18, "path": _hx_tmp_19, "module": ((any)(_hx_tmp_20)), "doc": _hx_tmp_21, "isPrivate": _hx_tmp_22, "params": _hx_tmp_23, "platforms": ((any)(this.VTable.Hx_Field_defplat())), "meta": ((any)(meta)), "athis": ((any)(athis)), "to": ((any)(to)), "from": ((any)(from)), "impl": impl})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xtypedef(x *Hx_Obj_xml) any {
	var doc HxNullable[string]
	_ = doc
	doc = HxNullable[string]{Value: HxStringNull}

	var t Hx_Enum_haxe_rtti_ctype
	_ = t
	var meta HxArray[any]
	_ = meta
	meta = HxMakeArray[any]().Dyn()
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_65 any
			_ = hx_dyncast_65
			hx_dyncast_65 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_65 != nil {
				_hx_tmp_2 = hx_dyncast_65.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			var tmp string
			_ = tmp
			if c.Hx_Field_nodeType == Hx_Field_xml_Document {
				tmp = "Document"
			} else {
				if c.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
				}

				tmp = c.Hx_Field_nodeName
			}

			if tmp == "haxe_doc" {
				doc = HxNullableString(Hx_Field_haxe_xml__access_access_impl__get_innerData(c))
			} else {
				var tmp_tmp_1 string
				_ = tmp_tmp_1
				if c.Hx_Field_nodeType == Hx_Field_xml_Document {
					tmp_tmp_1 = "Document"
				} else {
					if c.Hx_Field_nodeType != Hx_Field_xml_Element {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					tmp_tmp_1 = c.Hx_Field_nodeName
				}

				if tmp_tmp_1 == "meta" {
					meta = this.VTable.Hx_Field_xmeta(c)
				} else {
					t = this.VTable.Hx_Field_xtype(c)
				}
			}
		}
	}

	var types *Hx_Obj_haxe_ds_stringmap
	_ = types
	types = Hx_Obj_haxe_ds_stringmap_CreateInstance()
	if this.Hx_Field_curplatform != HxStringNull {
		func() func(_hx_param_key string, _hx_param_value Hx_Enum_haxe_rtti_ctype) {
			var _hx_closure_src_388 func(_hx_param_key any, _hx_param_value any)
			_ = _hx_closure_src_388
			_hx_closure_src_388 = types.VTable.Hx_Field_set
			return func(_hx_param_key string, _hx_param_value Hx_Enum_haxe_rtti_ctype) {
				_hx_closure_src_388(((any)(_hx_param_key)), ((any)(_hx_param_value)))
			}
		}()(this.Hx_Field_curplatform, t)
	}

	var _hx_tmp_3 HxNullable[string]
	_ = _hx_tmp_3
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "file") {
		_hx_tmp_3 = HxNullableString(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "file"))
	} else {
		_hx_tmp_3 = HxNullable[string]{Value: HxStringNull}
	}

	var hx_nullable_f196 HxNullable[string]
	_ = hx_nullable_f196
	hx_nullable_f196 = _hx_tmp_3
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	if (hx_nullable_f196).Valid == true {
		_hx_tmp_4 = (hx_nullable_f196).Value
	} else {
		_hx_tmp_4 = nil
	}

	var _hx_tmp_5 any
	_ = _hx_tmp_5
	_hx_tmp_5 = ((any)(this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))))
	var _hx_tmp_6 Hx_Typedef_haxe_rtti_path
	_ = _hx_tmp_6
	if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "module") {
		_hx_tmp_6 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "module"))
	} else {
		_hx_tmp_6 = ((Hx_Typedef_haxe_rtti_path)(HxStringNull))
	}

	var hx_nullable_f197 HxNullable[string]
	_ = hx_nullable_f197
	hx_nullable_f197 = doc
	var _hx_tmp_7 any
	_ = _hx_tmp_7
	if (hx_nullable_f197).Valid == true {
		_hx_tmp_7 = (hx_nullable_f197).Value
	} else {
		_hx_tmp_7 = nil
	}

	var _hx_tmp_8 any
	_ = _hx_tmp_8
	_hx_tmp_8 = ((any)(x.VTable.Hx_Field_exists("private")))
	var _hx_tmp_9 any
	_ = _hx_tmp_9
	_hx_tmp_9 = ((any)(this.VTable.Hx_Field_mkTypeParams(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "params"))))
	var _hx_tmp_10 any
	_ = _hx_tmp_10
	_hx_tmp_10 = ((any)(t))
	var _hx_tmp_11 any
	_ = _hx_tmp_11
	_hx_tmp_11 = any((func() *Hx_Obj_haxe_imap {
		_hx_icast_66 := types
		if _hx_icast_66 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_66, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_66.VTable.(Hx_Obj_VTable_haxe_imap)}
	})().VTable)
	return any(map[string]any{"file": _hx_tmp_4, "path": _hx_tmp_5, "module": ((any)(_hx_tmp_6)), "doc": _hx_tmp_7, "isPrivate": _hx_tmp_8, "params": _hx_tmp_9, "type": _hx_tmp_10, "types": _hx_tmp_11, "platforms": ((any)(this.VTable.Hx_Field_defplat())), "meta": ((any)(meta))})
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xtype(x *Hx_Obj_xml) Hx_Enum_haxe_rtti_ctype {
	var _g string
	_ = _g
	if x.Hx_Field_nodeType == Hx_Field_xml_Document {
		_g = "Document"
	} else {
		if x.Hx_Field_nodeType != Hx_Field_xml_Element {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(x.Hx_Field_nodeType))))))
		}

		_g = x.Hx_Field_nodeName
	}

	switch _g {
	case "a":
		{
			var fields HxArray[any]
			_ = fields
			fields = HxMakeArray[any]().Dyn()
			{
				var _gf any
				_ = _gf
				_gf = x.VTable.Hx_Field_elements()
				for {
					var _hx_tmp_0 any
					_ = _hx_tmp_0
					_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gf, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gf, "next")
					var hx_dyncast_67 any
					_ = hx_dyncast_67
					hx_dyncast_67 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
					var _hx_tmp_2 *Hx_Obj_xml
					_ = _hx_tmp_2
					if hx_dyncast_67 != nil {
						_hx_tmp_2 = hx_dyncast_67.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
					} else {
						_hx_tmp_2 = nil
					}

					var f *Hx_Obj_xml
					_ = f
					f = _hx_tmp_2
					var _hx_tmp_3 *Hx_Obj_xml
					_ = _hx_tmp_3
					_hx_tmp_3 = f
					var f1 any
					_ = f1
					f1 = this.VTable.Hx_Field_xclassfield(_hx_tmp_3, HxNullable[bool]{Value: true, Valid: true})
					var _hx_tmp_4 any
					_ = _hx_tmp_4
					_hx_tmp_4 = f1
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_4, "platforms", HxMakeArray[string]())
					_ = (Hx_Array_Push[any](HxMakeArrayView[any](fields), f1))
				}
			}

			return Hx_Enum_haxe_rtti_ctype_CAnonymous{fields}
		}

	case "c":
		{
			var _hx_tmp_5 Hx_Typedef_haxe_rtti_path
			_ = _hx_tmp_5
			_hx_tmp_5 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))
			return Hx_Enum_haxe_rtti_ctype_CClass{_hx_tmp_5, this.VTable.Hx_Field_xtypeparams(x)}
		}

	case "d":
		{
			var t Hx_Enum_haxe_rtti_ctype
			_ = t
			var tx *Hx_Obj_xml
			_ = tx
			tx = x.VTable.Hx_Field_firstElement()
			if tx != nil {
				if (tx.Hx_Field_nodeType != Hx_Field_xml_Document) && (tx.Hx_Field_nodeType != Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid nodeType ", Hx_Field__xml_xmltype_impl__toString(tx.Hx_Field_nodeType))))))
				}

				t = this.VTable.Hx_Field_xtype(tx)
			}

			return Hx_Enum_haxe_rtti_ctype_CDynamic{t}
		}

	case "e":
		{
			var _hx_tmp_6 Hx_Typedef_haxe_rtti_path
			_ = _hx_tmp_6
			_hx_tmp_6 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))
			return Hx_Enum_haxe_rtti_ctype_CEnum{_hx_tmp_6, this.VTable.Hx_Field_xtypeparams(x)}
		}

	case "f":
		{
			var args HxArray[any]
			_ = args
			args = HxMakeArray[any]().Dyn()
			var aname HxArray[string]
			_ = aname
			aname = (Hx_Field_go_haxe_hxstring_split(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "a"), ":"))
			var eargs_current int32
			_ = eargs_current
			eargs_current = 0
			var eargs_array HxArray[string]
			_ = eargs_array
			eargs_array = aname
			var _hx_tmp_7 *Hx_Obj_haxe_iterators_arrayiterator
			_ = _hx_tmp_7
			if Hx_Field_haxe_xml__access_hasattribaccess_impl__resolve(x, "v") {
				_hx_tmp_7 = Hx_Obj_haxe_iterators_arrayiterator_CreateInstance(HxMakeArrayView[any]((Hx_Field_go_haxe_hxstring_split(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "v"), ":"))))
			} else {
				_hx_tmp_7 = nil
			}

			var evalues *Hx_Obj_haxe_iterators_arrayiterator
			_ = evalues
			evalues = _hx_tmp_7
			{
				var _ge any
				_ = _ge
				_ge = x.VTable.Hx_Field_elements()
				for {
					var _hx_tmp_8 any
					_ = _hx_tmp_8
					_hx_tmp_8 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_8, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_9 any
					_ = _hx_tmp_9
					_hx_tmp_9 = Hx_Field_go_haxe_hxdynamic_getField(_ge, "next")
					var hx_dyncast_68 any
					_ = hx_dyncast_68
					hx_dyncast_68 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_9, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
					var _hx_tmp_10 *Hx_Obj_xml
					_ = _hx_tmp_10
					if hx_dyncast_68 != nil {
						_hx_tmp_10 = hx_dyncast_68.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
					} else {
						_hx_tmp_10 = nil
					}

					var e *Hx_Obj_xml
					_ = e
					e = _hx_tmp_10
					var opt bool
					_ = opt
					opt = false
					var _hx_tmp_11 HxNullable[string]
					_ = _hx_tmp_11
					var _hx_tmp_12 int32
					_ = _hx_tmp_12
					_hx_tmp_12 = eargs_current
					if _hx_tmp_12 < eargs_array.Len() {
						var hx_nullable_a46 HxArray[string]
						_ = hx_nullable_a46
						hx_nullable_a46 = eargs_array
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = eargs_current
						eargs_current = (eargs_current + 1)
						var hx_nullable_i47 int32
						_ = hx_nullable_i47
						hx_nullable_i47 = _hx_tmp_13
						var _hx_tmp_14 HxNullable[string]
						_ = _hx_tmp_14
						if hx_nullable_i47 >= 0 && hx_nullable_i47 < hx_nullable_a46.Len() {
							_hx_tmp_14 = HxNullable[string]{Value: hx_nullable_a46.Get(hx_nullable_i47), Valid: true}
						} else {
							_hx_tmp_14 = HxNullable[string]{Value: HxStringNull}
						}

						_hx_tmp_11 = _hx_tmp_14
					} else {
						_hx_tmp_11 = HxNullable[string]{Value: HxStringNull}
					}

					var a HxNullable[string]
					_ = a
					a = _hx_tmp_11
					if (a).Valid == false {
						a = HxNullableString("")
					}

					if (Hx_Field_go_haxe_hxstring_charAt(a.Value, 0)) == "?" {
						opt = true
						var _hx_tmp_15 string
						_ = _hx_tmp_15
						_hx_tmp_15 = a.Value
						a = HxNullableString((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_15, 1, HxNullable[int32]{})))
					}

					var _hx_tmp_16 HxNullable[string]
					_ = _hx_tmp_16
					var _hx_tmp_17 bool
					_ = _hx_tmp_17
					if evalues == nil {
						_hx_tmp_17 = true
					} else {
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = evalues.Hx_Field_current
						var _hx_tmp_19 int32
						_ = _hx_tmp_19
						_hx_tmp_19 = _hx_tmp_18
						_hx_tmp_17 = (_hx_tmp_19 >= HxMakeArrayView[string](evalues.Hx_Field_array).Len())
					}

					if _hx_tmp_17 {
						_hx_tmp_16 = HxNullable[string]{Value: HxStringNull}
					} else {
						var hx_nullable_a48 HxArray[string]
						_ = hx_nullable_a48
						hx_nullable_a48 = HxMakeArrayView[string](evalues.Hx_Field_array)
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						_hx_tmp_20 = evalues.Hx_Field_current
						evalues.Hx_Field_current = (evalues.Hx_Field_current + 1)
						var hx_nullable_i49 int32
						_ = hx_nullable_i49
						hx_nullable_i49 = _hx_tmp_20
						var _hx_tmp_21 HxNullable[string]
						_ = _hx_tmp_21
						if hx_nullable_i49 >= 0 && hx_nullable_i49 < hx_nullable_a48.Len() {
							_hx_tmp_21 = HxNullable[string]{Value: hx_nullable_a48.Get(hx_nullable_i49), Valid: true}
						} else {
							_hx_tmp_21 = HxNullable[string]{Value: HxStringNull}
						}

						_hx_tmp_16 = _hx_tmp_21
					}

					var v HxNullable[string]
					_ = v
					v = _hx_tmp_16
					var _hx_tmp_22 HxArray[any]
					_ = _hx_tmp_22
					_hx_tmp_22 = HxMakeArrayView[any](args)
					var hx_nullable_f198 HxNullable[string]
					_ = hx_nullable_f198
					hx_nullable_f198 = a
					var _hx_tmp_23 any
					_ = _hx_tmp_23
					if (hx_nullable_f198).Valid == true {
						_hx_tmp_23 = (hx_nullable_f198).Value
					} else {
						_hx_tmp_23 = nil
					}

					var _hx_tmp_24 any
					_ = _hx_tmp_24
					_hx_tmp_24 = ((any)(opt))
					var _hx_tmp_25 any
					_ = _hx_tmp_25
					_hx_tmp_25 = ((any)(this.VTable.Hx_Field_xtype(e)))
					var _hx_tmp_26 HxNullable[string]
					_ = _hx_tmp_26
					var _hx_tmp_27 bool
					_ = _hx_tmp_27
					if (v).Valid == true {
						_hx_tmp_27 = ((v).Value == "")
					} else {
						_hx_tmp_27 = false
					}

					if _hx_tmp_27 {
						_hx_tmp_26 = HxNullable[string]{Value: HxStringNull}
					} else {
						_hx_tmp_26 = v
					}

					var hx_nullable_f199 HxNullable[string]
					_ = hx_nullable_f199
					hx_nullable_f199 = _hx_tmp_26
					var _hx_tmp_28 any
					_ = _hx_tmp_28
					if (hx_nullable_f199).Valid == true {
						_hx_tmp_28 = (hx_nullable_f199).Value
					} else {
						_hx_tmp_28 = nil
					}

					_ = (Hx_Array_Push[any](_hx_tmp_22, any(map[string]any{"name": _hx_tmp_23, "opt": _hx_tmp_24, "t": _hx_tmp_25, "value": _hx_tmp_28})))
				}
			}

			var _hx_tmp_29 HxArray[any]
			_ = _hx_tmp_29
			_hx_tmp_29 = args
			var ret any
			_ = ret
			ret = _hx_tmp_29.Get((Hx_Field_go_haxe_hxdynamic_getArrayLength(args) - 1))
			_ = (Hx_Array_Remove[any](HxMakeArrayView[any](args), ret))
			var _hx_tmp_30 HxArray[any]
			_ = _hx_tmp_30
			_hx_tmp_30 = args
			return Hx_Enum_haxe_rtti_ctype_CFunction{_hx_tmp_30, func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(Hx_Field_go_haxe_hxdynamic_getField(ret, "t"))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()}
		}

	case "t":
		{
			var _hx_tmp_31 Hx_Typedef_haxe_rtti_path
			_ = _hx_tmp_31
			_hx_tmp_31 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))
			return Hx_Enum_haxe_rtti_ctype_CTypedef{_hx_tmp_31, this.VTable.Hx_Field_xtypeparams(x)}
		}

	case "unknown":
		{
			return Hx_Enum_haxe_rtti_ctype_CUnknown{}
		}

	case "x":
		{
			var _hx_tmp_32 Hx_Typedef_haxe_rtti_path
			_ = _hx_tmp_32
			_hx_tmp_32 = this.VTable.Hx_Field_mkPath(Hx_Field_haxe_xml__access_attribaccess_impl__resolve(x, "path"))
			return Hx_Enum_haxe_rtti_ctype_CAbstract{_hx_tmp_32, this.VTable.Hx_Field_xtypeparams(x)}
		}

	default:
		{
			return func() Hx_Enum_haxe_rtti_ctype {
				_hx_d := Hx_Field_go_haxe_hxdynamic_ensureInterface(this.VTable.Hx_Field_xerror(x))
				if _hx_d == nil {
					var _hx_z Hx_Enum_haxe_rtti_ctype
					return _hx_z
				}
				return _hx_d.(Hx_Enum_haxe_rtti_ctype)
			}()
		}
	}
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_xtypeparams(x *Hx_Obj_xml) HxArray[Hx_Enum_haxe_rtti_ctype] {
	var p HxArray[Hx_Enum_haxe_rtti_ctype]
	_ = p
	p = HxMakeArray[Hx_Enum_haxe_rtti_ctype]()
	{
		var _gc any
		_ = _gc
		_gc = x.VTable.Hx_Field_elements()
		for {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_1 any
			_ = _hx_tmp_1
			_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(_gc, "next")
			var hx_dyncast_69 any
			_ = hx_dyncast_69
			hx_dyncast_69 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_2 *Hx_Obj_xml
			_ = _hx_tmp_2
			if hx_dyncast_69 != nil {
				_hx_tmp_2 = hx_dyncast_69.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_2 = nil
			}

			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_2
			var _hx_tmp_3 HxArray[Hx_Enum_haxe_rtti_ctype]
			_ = _hx_tmp_3
			_hx_tmp_3 = p
			_ = (Hx_Array_Push[Hx_Enum_haxe_rtti_ctype](_hx_tmp_3, this.VTable.Hx_Field_xtype(c)))
		}
	}

	return p
}

func (this *Hx_Obj_haxe_rtti_xmlparser) Hx_Field_defplat() HxArray[string] {
	var l HxArray[string]
	_ = l
	l = HxMakeArray[string]()
	if this.Hx_Field_curplatform != HxStringNull {
		_ = (Hx_Array_Push[string](l, this.Hx_Field_curplatform))
	}

	return l
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_xml]()
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_ctype]()
	hxRegisterArrayConstructor[Hx_Enum_haxe_rtti_typetree]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
