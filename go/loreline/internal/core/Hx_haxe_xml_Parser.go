package core

import "unicode/utf8"

var Hx_Abstract_haxe_xml__parser_s_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Parser.S",
	HxMakeArray[string]("IGNORE_SPACES", "BEGIN", "BEGIN_NODE", "TAG_NAME", "BODY", "ATTRIB_NAME", "EQUALS", "ATTVAL_BEGIN", "ATTRIB_VAL", "CHILDS", "CLOSE", "WAIT_END", "WAIT_END_RET", "PCDATA", "HEADER", "COMMENT", "DOCTYPE", "CDATA", "ESCAPE"),
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

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml__parser_s_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml__parser_s_impl_() {
	Hx_Field_haxe_xml__parser_s_impl__IGNORE_SPACES = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__IGNORE_SPACES()
	Hx_Field_haxe_xml__parser_s_impl__BEGIN = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BEGIN()
	Hx_Field_haxe_xml__parser_s_impl__BEGIN_NODE = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BEGIN_NODE()
	Hx_Field_haxe_xml__parser_s_impl__TAG_NAME = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__TAG_NAME()
	Hx_Field_haxe_xml__parser_s_impl__BODY = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BODY()
	Hx_Field_haxe_xml__parser_s_impl__ATTRIB_NAME = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTRIB_NAME()
	Hx_Field_haxe_xml__parser_s_impl__EQUALS = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__EQUALS()
	Hx_Field_haxe_xml__parser_s_impl__ATTVAL_BEGIN = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTVAL_BEGIN()
	Hx_Field_haxe_xml__parser_s_impl__ATTRIB_VAL = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTRIB_VAL()
	Hx_Field_haxe_xml__parser_s_impl__CHILDS = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CHILDS()
	Hx_Field_haxe_xml__parser_s_impl__CLOSE = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CLOSE()
	Hx_Field_haxe_xml__parser_s_impl__WAIT_END = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__WAIT_END()
	Hx_Field_haxe_xml__parser_s_impl__WAIT_END_RET = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__WAIT_END_RET()
	Hx_Field_haxe_xml__parser_s_impl__PCDATA = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__PCDATA()
	Hx_Field_haxe_xml__parser_s_impl__HEADER = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__HEADER()
	Hx_Field_haxe_xml__parser_s_impl__COMMENT = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__COMMENT()
	Hx_Field_haxe_xml__parser_s_impl__DOCTYPE = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__DOCTYPE()
	Hx_Field_haxe_xml__parser_s_impl__CDATA = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CDATA()
	Hx_Field_haxe_xml__parser_s_impl__ESCAPE = Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ESCAPE()
}

type Hx_Obj_VTable_haxe_xml__parser_s_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml__parser_s_impl_ struct {
	VTable Hx_Obj_VTable_haxe_xml__parser_s_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml__parser_s_impl__CreateEmptyInstance() *Hx_Obj_haxe_xml__parser_s_impl_ {
	obj := &Hx_Obj_haxe_xml__parser_s_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml__parser_s_impl__CreateInstance() *Hx_Obj_haxe_xml__parser_s_impl_ {
	obj := Hx_Obj_haxe_xml__parser_s_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml__parser_s_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml._Parser.S_Impl_",
	HxMakeArray[string]("IGNORE_SPACES", "BEGIN", "BEGIN_NODE", "TAG_NAME", "BODY", "ATTRIB_NAME", "EQUALS", "ATTVAL_BEGIN", "ATTRIB_VAL", "CHILDS", "CLOSE", "WAIT_END", "WAIT_END_RET", "PCDATA", "HEADER", "COMMENT", "DOCTYPE", "CDATA", "ESCAPE"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml__parser_s_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml__parser_s_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml__parser_s_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "IGNORE_SPACES":
			return &Hx_Field_haxe_xml__parser_s_impl__IGNORE_SPACES
		case "BEGIN":
			return &Hx_Field_haxe_xml__parser_s_impl__BEGIN
		case "BEGIN_NODE":
			return &Hx_Field_haxe_xml__parser_s_impl__BEGIN_NODE
		case "TAG_NAME":
			return &Hx_Field_haxe_xml__parser_s_impl__TAG_NAME
		case "BODY":
			return &Hx_Field_haxe_xml__parser_s_impl__BODY
		case "ATTRIB_NAME":
			return &Hx_Field_haxe_xml__parser_s_impl__ATTRIB_NAME
		case "EQUALS":
			return &Hx_Field_haxe_xml__parser_s_impl__EQUALS
		case "ATTVAL_BEGIN":
			return &Hx_Field_haxe_xml__parser_s_impl__ATTVAL_BEGIN
		case "ATTRIB_VAL":
			return &Hx_Field_haxe_xml__parser_s_impl__ATTRIB_VAL
		case "CHILDS":
			return &Hx_Field_haxe_xml__parser_s_impl__CHILDS
		case "CLOSE":
			return &Hx_Field_haxe_xml__parser_s_impl__CLOSE
		case "WAIT_END":
			return &Hx_Field_haxe_xml__parser_s_impl__WAIT_END
		case "WAIT_END_RET":
			return &Hx_Field_haxe_xml__parser_s_impl__WAIT_END_RET
		case "PCDATA":
			return &Hx_Field_haxe_xml__parser_s_impl__PCDATA
		case "HEADER":
			return &Hx_Field_haxe_xml__parser_s_impl__HEADER
		case "COMMENT":
			return &Hx_Field_haxe_xml__parser_s_impl__COMMENT
		case "DOCTYPE":
			return &Hx_Field_haxe_xml__parser_s_impl__DOCTYPE
		case "CDATA":
			return &Hx_Field_haxe_xml__parser_s_impl__CDATA
		case "ESCAPE":
			return &Hx_Field_haxe_xml__parser_s_impl__ESCAPE
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml__parser_s_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml__parser_s_impl__RTTI
}

func (this *Hx_Obj_haxe_xml__parser_s_impl_) Hx_As_Hx_Obj_haxe_xml__parser_s_impl_() *Hx_Obj_haxe_xml__parser_s_impl_ {
	return this
}

var Hx_Field_haxe_xml__parser_s_impl__IGNORE_SPACES int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__IGNORE_SPACES() int32 {
	return ((int32)(0))
}

var Hx_Field_haxe_xml__parser_s_impl__BEGIN int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BEGIN() int32 {
	return ((int32)(1))
}

var Hx_Field_haxe_xml__parser_s_impl__BEGIN_NODE int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BEGIN_NODE() int32 {
	return ((int32)(2))
}

var Hx_Field_haxe_xml__parser_s_impl__TAG_NAME int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__TAG_NAME() int32 {
	return ((int32)(3))
}

var Hx_Field_haxe_xml__parser_s_impl__BODY int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__BODY() int32 {
	return ((int32)(4))
}

var Hx_Field_haxe_xml__parser_s_impl__ATTRIB_NAME int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTRIB_NAME() int32 {
	return ((int32)(5))
}

var Hx_Field_haxe_xml__parser_s_impl__EQUALS int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__EQUALS() int32 {
	return ((int32)(6))
}

var Hx_Field_haxe_xml__parser_s_impl__ATTVAL_BEGIN int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTVAL_BEGIN() int32 {
	return ((int32)(7))
}

var Hx_Field_haxe_xml__parser_s_impl__ATTRIB_VAL int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ATTRIB_VAL() int32 {
	return ((int32)(8))
}

var Hx_Field_haxe_xml__parser_s_impl__CHILDS int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CHILDS() int32 {
	return ((int32)(9))
}

var Hx_Field_haxe_xml__parser_s_impl__CLOSE int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CLOSE() int32 {
	return ((int32)(10))
}

var Hx_Field_haxe_xml__parser_s_impl__WAIT_END int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__WAIT_END() int32 {
	return ((int32)(11))
}

var Hx_Field_haxe_xml__parser_s_impl__WAIT_END_RET int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__WAIT_END_RET() int32 {
	return ((int32)(12))
}

var Hx_Field_haxe_xml__parser_s_impl__PCDATA int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__PCDATA() int32 {
	return ((int32)(13))
}

var Hx_Field_haxe_xml__parser_s_impl__HEADER int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__HEADER() int32 {
	return ((int32)(14))
}

var Hx_Field_haxe_xml__parser_s_impl__COMMENT int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__COMMENT() int32 {
	return ((int32)(15))
}

var Hx_Field_haxe_xml__parser_s_impl__DOCTYPE int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__DOCTYPE() int32 {
	return ((int32)(16))
}

var Hx_Field_haxe_xml__parser_s_impl__CDATA int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__CDATA() int32 {
	return ((int32)(17))
}

var Hx_Field_haxe_xml__parser_s_impl__ESCAPE int32

func Hx_Init_Hx_Field_haxe_xml__parser_s_impl__ESCAPE() int32 {
	return ((int32)(18))
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml_xmlparserexception() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml_xmlparserexception() {
}

type Hx_Obj_VTable_haxe_xml_xmlparserexception interface {
	Hx_Field_toString() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml_xmlparserexception struct {
	VTable                  Hx_Obj_VTable_haxe_xml_xmlparserexception `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_message        string
	Hx_Field_lineNumber     int32
	Hx_Field_positionAtLine int32
	Hx_Field_position       int32
	Hx_Field_xml            string
}

func Hx_Obj_haxe_xml_xmlparserexception_CreateEmptyInstance() *Hx_Obj_haxe_xml_xmlparserexception {
	obj := &Hx_Obj_haxe_xml_xmlparserexception{}
	obj.VTable = obj
	obj.Hx_Field_message = HxDefault[string]()
	obj.Hx_Field_lineNumber = HxDefault[int32]()
	obj.Hx_Field_positionAtLine = HxDefault[int32]()
	obj.Hx_Field_position = HxDefault[int32]()
	obj.Hx_Field_xml = HxDefault[string]()
	return obj
}

func Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(message string, xml string, position int32) *Hx_Obj_haxe_xml_xmlparserexception {
	obj := Hx_Obj_haxe_xml_xmlparserexception_CreateEmptyInstance()
	obj.Hx_New(message, xml, position)
	return obj
}

func (this *Hx_Obj_haxe_xml_xmlparserexception) Hx_New(message string, xml string, position int32) {
	{
		this.Hx_Field_xml = xml
		this.Hx_Field_message = message
		this.Hx_Field_position = position
		this.Hx_Field_lineNumber = 1
		this.Hx_Field_positionAtLine = 0
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = position
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				var c int32
				_ = c
				c = (Hx_Field_go_haxe_hxstring_charCodeAt(xml, i)).Value
				if c == 10 {
					this.Hx_Field_lineNumber++
					this.Hx_Field_positionAtLine = 0
				} else {
					if c != 13 {
						this.Hx_Field_positionAtLine++
					}
				}
			}
		}
	}

}

var Hx_Obj_haxe_xml_xmlparserexception_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml.XmlParserException",
	HxMakeArray[string](),
	HxMakeArray[string]("message", "lineNumber", "positionAtLine", "position", "xml", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[string](params.Get(1)), HxConvert[int32](params.Get(2)))
	},
	func() any {
		return Hx_Obj_haxe_xml_xmlparserexception_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml_xmlparserexception_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_haxe_xml_xmlparserexception) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml_xmlparserexception_RTTI
}

func (this *Hx_Obj_haxe_xml_xmlparserexception) Hx_As_Hx_Obj_haxe_xml_xmlparserexception() *Hx_Obj_haxe_xml_xmlparserexception {
	return this
}

func (this *Hx_Obj_haxe_xml_xmlparserexception) Hx_Field_toString() string {
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(Hx_Field_type_getClassName(((*Hx_Obj_go_haxe_hxclass)(func() func(o *Hx_Obj_haxe_xml_xmlparserexception) *Hx_Obj_go_haxe_hxclass {
		var _hx_closure_src_362 func(o any) *Hx_Obj_go_haxe_hxclass
		_ = _hx_closure_src_362
		_hx_closure_src_362 = Hx_Field_type_getClass
		return func(o *Hx_Obj_haxe_xml_xmlparserexception) *Hx_Obj_go_haxe_hxclass {
			return ((*Hx_Obj_go_haxe_hxclass)(_hx_closure_src_362(((any)(o)))))
		}
	}()(this)))), ": "), this.Hx_Field_message), " at line ")
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat((_hx_tmp_1 + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_lineNumber)), " char ")
	return (_hx_tmp_0 + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_positionAtLine))
}

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml_parser() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml_parser() {
	Hx_Field_haxe_xml_parser_escapes = Hx_Init_Hx_Field_haxe_xml_parser_escapes()
}

type Hx_Obj_VTable_haxe_xml_parser interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml_parser struct {
	VTable Hx_Obj_VTable_haxe_xml_parser `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_xml_parser_CreateEmptyInstance() *Hx_Obj_haxe_xml_parser {
	obj := &Hx_Obj_haxe_xml_parser{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_xml_parser_CreateInstance() *Hx_Obj_haxe_xml_parser {
	obj := Hx_Obj_haxe_xml_parser_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_xml_parser_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml.Parser",
	HxMakeArray[string]("escapes", "parse", "doParse", "isValidChar"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml_parser_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_xml_parser_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml_parser_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "escapes":
			return &Hx_Field_haxe_xml_parser_escapes
		case "parse":
			return Hx_Field_haxe_xml_parser_parse
		case "doParse":
			return Hx_Field_haxe_xml_parser_doParse
		case "isValidChar":
			return Hx_Field_haxe_xml_parser_isValidChar
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml_parser) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml_parser_RTTI
}

func (this *Hx_Obj_haxe_xml_parser) Hx_As_Hx_Obj_haxe_xml_parser() *Hx_Obj_haxe_xml_parser {
	return this
}

var Hx_Field_haxe_xml_parser_escapes *Hx_Obj_haxe_ds_stringmap

func Hx_Init_Hx_Field_haxe_xml_parser_escapes() *Hx_Obj_haxe_ds_stringmap {
	var h *Hx_Obj_haxe_ds_stringmap
	_ = h
	h = Hx_Obj_haxe_ds_stringmap_CreateInstance()
	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_363 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_363
		_hx_closure_src_363 = h.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_363(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("lt", "<")
	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_364 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_364
		_hx_closure_src_364 = h.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_364(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("gt", ">")
	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_365 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_365
		_hx_closure_src_365 = h.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_365(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("amp", "&")
	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_366 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_366
		_hx_closure_src_366 = h.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_366(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("quot", "\"")
	func() func(_hx_param_key string, _hx_param_value string) {
		var _hx_closure_src_367 func(_hx_param_key any, _hx_param_value any)
		_ = _hx_closure_src_367
		_hx_closure_src_367 = h.VTable.Hx_Field_set
		return func(_hx_param_key string, _hx_param_value string) {
			_hx_closure_src_367(((any)(_hx_param_key)), ((any)(_hx_param_value)))
		}
	}()("apos", "'")
	return h
}

func Hx_Field_haxe_xml_parser_parse(str string, strict__isOptional any) *Hx_Obj_xml {
	strict := HxDefault[bool]()
	_ = strict

	if strict__isOptional != nil {
		strict = HxConvert[bool](strict__isOptional)
	} else {
		strict = false
	}
	{
		var doc *Hx_Obj_xml
		_ = doc
		doc = Hx_Field_xml_createDocument()
		Hx_Field_haxe_xml_parser_doParse(str, strict, 0, doc)
		return doc
	}
}

func Hx_Field_haxe_xml_parser_doParse(str string, strict bool, p__isOptional any, parent__isOptional any) int32 {
	p := HxDefault[int32]()
	_ = p

	if p__isOptional != nil {
		p = HxConvert[int32](p__isOptional)
	} else {
		p = 0
	}
	parent := HxDefault[*Hx_Obj_xml]()
	_ = parent

	if parent__isOptional != nil {
		parent = HxConvert[*Hx_Obj_xml](parent__isOptional)
	} else {
		parent = nil
	}
	{
		var xml *Hx_Obj_xml
		_ = xml
		var state int32
		_ = state
		state = ((int32)(1))
		var next int32
		_ = next
		next = ((int32)(1))
		var aname HxNullable[string]
		_ = aname
		aname = HxNullable[string]{Value: HxStringNull}

		var start int32
		_ = start
		start = 0
		var nsubs int32
		_ = nsubs
		nsubs = 0
		var nbrackets int32
		_ = nbrackets
		nbrackets = 0
		var buf *Hx_Obj_stringbuf
		_ = buf
		buf = Hx_Obj_stringbuf_CreateInstance()
		var escapeNext int32
		_ = escapeNext
		escapeNext = ((int32)(1))
		var attrValQuote int32
		_ = attrValQuote
		attrValQuote = -1
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = p
			if !(_hx_tmp_0 < int32(utf8.RuneCountInString(str))) {
				break
			}

			var c int32
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(str, p)).Value
			switch state {
			case 0:
				{
					switch c {
					case 9, 10, 13, 32:
						{
						}

					default:
						{
							state = next
							continue
						}
					}
				}

			case 1:
				{
					if c == 60 {
						state = ((int32)(0))
						next = ((int32)(2))
					} else {
						start = p
						state = ((int32)(13))
						continue
					}
				}

			case 2:
				{
					switch c {
					case 33:
						{
							if (Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 91 {
								p = (p + 2)
								var _hx_tmp_1 string
								_ = _hx_tmp_1
								_hx_tmp_1 = str
								var _hx_tmp_2 int32
								_ = _hx_tmp_2
								_hx_tmp_2 = p
								if Hx_Field_go_haxe_hxstring_toUpperCase((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, _hx_tmp_2, HxNullable[int32]{Value: 6, Valid: true}))) != "CDATA[" {
									panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected <![CDATA[", str, p)))))
								}

								p = (p + 5)
								state = ((int32)(17))
								start = (p + 1)
							} else {
								var _hx_tmp_3 bool
								_ = _hx_tmp_3
								if (Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 68 {
									_hx_tmp_3 = true
								} else {
									_hx_tmp_3 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 100)
								}

								if _hx_tmp_3 {
									var _hx_tmp_4 string
									_ = _hx_tmp_4
									_hx_tmp_4 = str
									var _hx_tmp_5 int32
									_ = _hx_tmp_5
									_hx_tmp_5 = (p + 2)
									if Hx_Field_go_haxe_hxstring_toUpperCase((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, _hx_tmp_5, HxNullable[int32]{Value: 6, Valid: true}))) != "OCTYPE" {
										panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected <!DOCTYPE", str, p)))))
									}

									p = (p + 8)
									state = ((int32)(16))
									start = (p + 1)
								} else {
									var _hx_tmp_6 bool
									_ = _hx_tmp_6
									if (Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value != 45 {
										_hx_tmp_6 = true
									} else {
										_hx_tmp_6 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 2))).Value != 45)
									}

									if _hx_tmp_6 {
										panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected <!--", str, p)))))
									} else {
										p = (p + 2)
										state = ((int32)(15))
										start = (p + 1)
									}
								}
							}
						}

					case 47:
						{
							if parent == nil {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected node name", str, p)))))
							}

							start = (p + 1)
							state = ((int32)(0))
							next = ((int32)(10))
						}

					case 63:
						{
							state = ((int32)(14))
							start = p
						}

					default:
						{
							state = ((int32)(3))
							start = p
							continue
						}
					}
				}

			case 3:
				{
					if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
						if p == start {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected node name", str, p)))))
						}

						var _hx_tmp_7 string
						_ = _hx_tmp_7
						_hx_tmp_7 = str
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						_hx_tmp_8 = start
						xml = Hx_Field_xml_createElement((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, _hx_tmp_8, HxNullable[int32]{Value: (p - start), Valid: true})))
						{
							parent.VTable.Hx_Field_addChild(xml)
							nsubs++
						}

						state = ((int32)(0))
						next = ((int32)(4))
						continue
					}
				}

			case 4:
				{
					switch c {
					case 47:
						{
							state = ((int32)(11))
						}

					case 62:
						{
							state = ((int32)(9))
						}

					default:
						{
							state = ((int32)(5))
							start = p
							continue
						}
					}
				}

			case 5:
				{
					if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
						if start == p {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected attribute name", str, p)))))
						}

						var _hx_tmp_9 string
						_ = _hx_tmp_9
						_hx_tmp_9 = str
						var _hx_tmp_10 int32
						_ = _hx_tmp_10
						_hx_tmp_10 = start
						var tmp string
						_ = tmp
						tmp = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_9, _hx_tmp_10, HxNullable[int32]{Value: (p - start), Valid: true}))
						aname = HxNullableString(tmp)
						if xml.VTable.Hx_Field_exists((aname).Value) {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Duplicate attribute [", (aname).Value), "]"), str, p)))))
						}

						state = ((int32)(0))
						next = ((int32)(6))
						continue
					}
				}

			case 6:
				{
					if c == 61 {
						state = ((int32)(0))
						next = ((int32)(7))
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected =", str, p)))))
					}
				}

			case 7:
				{
					switch c {
					case 34, 39:
						{
							buf = Hx_Obj_stringbuf_CreateInstance()
							state = ((int32)(8))
							start = (p + 1)
							attrValQuote = c
						}

					default:
						{
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected \"", str, p)))))
						}
					}
				}

			case 8:
				{
					switch c {
					case 38:
						{
							var _hx_tmp_11 string
							_ = _hx_tmp_11
							_hx_tmp_11 = str
							var _hx_tmp_12 int32
							_ = _hx_tmp_12
							_hx_tmp_12 = start
							buf.VTable.Hx_Field_addSub(_hx_tmp_11, _hx_tmp_12, HxNullable[int32]{Value: (p - start), Valid: true})
							state = ((int32)(18))
							escapeNext = ((int32)(8))
							start = (p + 1)
						}

					case 60, 62:
						{
							if strict {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Invalid unescaped ", Hx_Field_go_haxe_hxstring_fromCharCode(c)), " in attribute value"), str, p)))))
							} else {
								if c == attrValQuote {
									var _hx_tmp_13 string
									_ = _hx_tmp_13
									_hx_tmp_13 = str
									var _hx_tmp_14 int32
									_ = _hx_tmp_14
									_hx_tmp_14 = start
									buf.VTable.Hx_Field_addSub(_hx_tmp_13, _hx_tmp_14, HxNullable[int32]{Value: (p - start), Valid: true})
									var val string
									_ = val
									val = buf.VTable.Hx_Field_toString()
									buf = Hx_Obj_stringbuf_CreateInstance()
									xml.VTable.Hx_Field_set((aname).Value, val)
									state = ((int32)(0))
									next = ((int32)(4))
								}
							}
						}

					default:
						{
							if c == attrValQuote {
								var _hx_tmp_15 string
								_ = _hx_tmp_15
								_hx_tmp_15 = str
								var _hx_tmp_16 int32
								_ = _hx_tmp_16
								_hx_tmp_16 = start
								buf.VTable.Hx_Field_addSub(_hx_tmp_15, _hx_tmp_16, HxNullable[int32]{Value: (p - start), Valid: true})
								var val string
								_ = val
								val = buf.VTable.Hx_Field_toString()
								buf = Hx_Obj_stringbuf_CreateInstance()
								xml.VTable.Hx_Field_set((aname).Value, val)
								state = ((int32)(0))
								next = ((int32)(4))
							}
						}
					}
				}

			case 9:
				{
					p = Hx_Field_haxe_xml_parser_doParse(str, strict, p, xml)
					start = p
					state = ((int32)(1))
				}

			case 10:
				{
					if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) {
						if start == p {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected node name", str, p)))))
						}

						var _hx_tmp_17 string
						_ = _hx_tmp_17
						_hx_tmp_17 = str
						var _hx_tmp_18 int32
						_ = _hx_tmp_18
						_hx_tmp_18 = start
						var v string
						_ = v
						v = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_17, _hx_tmp_18, HxNullable[int32]{Value: (p - start), Valid: true}))
						if (parent == nil) || (parent.Hx_Field_nodeType != ((int32)(0))) {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unexpected </", v), ">, tag is not open"), str, p)))))
						}

						if parent.Hx_Field_nodeType != Hx_Field_xml_Element {
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(parent.Hx_Field_nodeType))))))
						}

						if v != parent.Hx_Field_nodeName {
							if parent.Hx_Field_nodeType != Hx_Field_xml_Element {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(parent.Hx_Field_nodeType))))))
							}

							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Expected </", parent.Hx_Field_nodeName), ">"), str, p)))))
						}

						state = ((int32)(0))
						next = ((int32)(12))
						continue
					}
				}

			case 11:
				{
					if c == 62 {
						state = ((int32)(1))
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected >", str, p)))))
					}
				}

			case 12:
				{
					if c == 62 {
						if nsubs == 0 {
							parent.VTable.Hx_Field_addChild(Hx_Field_xml_createPCData(""))
						}

						return p
					} else {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Expected >", str, p)))))
					}
				}

			case 13:
				{
					if c == 60 {
						var _hx_tmp_19 string
						_ = _hx_tmp_19
						_hx_tmp_19 = str
						var _hx_tmp_20 int32
						_ = _hx_tmp_20
						_hx_tmp_20 = start
						buf.VTable.Hx_Field_addSub(_hx_tmp_19, _hx_tmp_20, HxNullable[int32]{Value: (p - start), Valid: true})
						var child *Hx_Obj_xml
						_ = child
						child = Hx_Field_xml_createPCData(buf.VTable.Hx_Field_toString())
						buf = Hx_Obj_stringbuf_CreateInstance()
						{
							parent.VTable.Hx_Field_addChild(child)
							nsubs++
						}

						state = ((int32)(0))
						next = ((int32)(2))
					} else {
						if c == 38 {
							var _hx_tmp_21 string
							_ = _hx_tmp_21
							_hx_tmp_21 = str
							var _hx_tmp_22 int32
							_ = _hx_tmp_22
							_hx_tmp_22 = start
							buf.VTable.Hx_Field_addSub(_hx_tmp_21, _hx_tmp_22, HxNullable[int32]{Value: (p - start), Valid: true})
							state = ((int32)(18))
							escapeNext = ((int32)(13))
							start = (p + 1)
						}
					}
				}

			case 14:
				{
					var _hx_tmp_23 bool
					_ = _hx_tmp_23
					if c == 63 {
						_hx_tmp_23 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 62)
					} else {
						_hx_tmp_23 = false
					}

					if _hx_tmp_23 {
						p++
						var _hx_tmp_24 string
						_ = _hx_tmp_24
						_hx_tmp_24 = str
						var _hx_tmp_25 int32
						_ = _hx_tmp_25
						_hx_tmp_25 = (start + 1)
						var str1 string
						_ = str1
						str1 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_24, _hx_tmp_25, HxNullable[int32]{Value: ((p - start) - 2), Valid: true}))
						{
							parent.VTable.Hx_Field_addChild(Hx_Field_xml_createProcessingInstruction(str1))
							nsubs++
						}

						state = ((int32)(1))
					}
				}

			case 15:
				{
					var _hx_tmp_26 bool
					_ = _hx_tmp_26
					var _hx_tmp_27 bool
					_ = _hx_tmp_27
					if c == 45 {
						_hx_tmp_27 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 45)
					} else {
						_hx_tmp_27 = false
					}

					if _hx_tmp_27 {
						_hx_tmp_26 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 2))).Value == 62)
					} else {
						_hx_tmp_26 = false
					}

					if _hx_tmp_26 {
						{
							var _hx_tmp_28 string
							_ = _hx_tmp_28
							_hx_tmp_28 = str
							var _hx_tmp_29 int32
							_ = _hx_tmp_29
							_hx_tmp_29 = start
							parent.VTable.Hx_Field_addChild(Hx_Field_xml_createComment((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_28, _hx_tmp_29, HxNullable[int32]{Value: (p - start), Valid: true}))))
							nsubs++
						}

						p = (p + 2)
						state = ((int32)(1))
					}
				}

			case 16:
				{
					if c == 91 {
						nbrackets++
					} else {
						if c == 93 {
							nbrackets--
						} else {
							if (c == 62) && (nbrackets == 0) {
								{
									var _hx_tmp_30 string
									_ = _hx_tmp_30
									_hx_tmp_30 = str
									var _hx_tmp_31 int32
									_ = _hx_tmp_31
									_hx_tmp_31 = start
									parent.VTable.Hx_Field_addChild(Hx_Field_xml_createDocType((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_30, _hx_tmp_31, HxNullable[int32]{Value: (p - start), Valid: true}))))
									nsubs++
								}

								state = ((int32)(1))
							}
						}
					}
				}

			case 17:
				{
					var _hx_tmp_32 bool
					_ = _hx_tmp_32
					var _hx_tmp_33 bool
					_ = _hx_tmp_33
					if c == 93 {
						_hx_tmp_33 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 1))).Value == 93)
					} else {
						_hx_tmp_33 = false
					}

					if _hx_tmp_33 {
						_hx_tmp_32 = ((Hx_Field_go_haxe_hxstring_charCodeAt(str, (p + 2))).Value == 62)
					} else {
						_hx_tmp_32 = false
					}

					if _hx_tmp_32 {
						var _hx_tmp_34 string
						_ = _hx_tmp_34
						_hx_tmp_34 = str
						var _hx_tmp_35 int32
						_ = _hx_tmp_35
						_hx_tmp_35 = start
						var child *Hx_Obj_xml
						_ = child
						child = Hx_Field_xml_createCData((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_34, _hx_tmp_35, HxNullable[int32]{Value: (p - start), Valid: true})))
						{
							parent.VTable.Hx_Field_addChild(child)
							nsubs++
						}

						p = (p + 2)
						state = ((int32)(1))
					}
				}

			case 18:
				{
					if c == 59 {
						var _hx_tmp_36 string
						_ = _hx_tmp_36
						_hx_tmp_36 = str
						var _hx_tmp_37 int32
						_ = _hx_tmp_37
						_hx_tmp_37 = start
						var s string
						_ = s
						s = Hx_Field_go_haxe_hxstring_toLowerCase((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_36, _hx_tmp_37, HxNullable[int32]{Value: (p - start), Valid: true})))
						if (Hx_Field_go_haxe_hxstring_charCodeAt(s, 0)).Value == 35 {
							var _hx_tmp_38 HxNullable[int32]
							_ = _hx_tmp_38
							if (Hx_Field_go_haxe_hxstring_charCodeAt(s, 1)).Value == 120 {
								var _hx_tmp_39 string
								_ = _hx_tmp_39
								_hx_tmp_39 = s
								var _hx_tmp_40 string
								_ = _hx_tmp_40
								_hx_tmp_40 = _hx_tmp_39
								_hx_tmp_38 = Hx_Field_std_parseInt(Hx_Field_go_haxe_hxstring_concat("0", (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_40, 1, HxNullable[int32]{Value: (int32(utf8.RuneCountInString(s)) - 1), Valid: true}))))
							} else {
								var _hx_tmp_41 string
								_ = _hx_tmp_41
								_hx_tmp_41 = s
								var _hx_tmp_42 string
								_ = _hx_tmp_42
								_hx_tmp_42 = _hx_tmp_41
								_hx_tmp_38 = Hx_Field_std_parseInt((Hx_Field_go_haxe_hxstring_substr(_hx_tmp_42, 1, HxNullable[int32]{Value: (int32(utf8.RuneCountInString(s)) - 1), Valid: true})))
							}

							var tmp_c_1 HxNullable[int32]
							_ = tmp_c_1
							tmp_c_1 = _hx_tmp_38
							buf.VTable.Hx_Field_addChar((tmp_c_1).Value)
						} else {
							if !func() func(_hx_param_key string) bool {
								var _hx_closure_src_368 func(_hx_param_key any) bool
								_ = _hx_closure_src_368
								_hx_closure_src_368 = Hx_Field_haxe_xml_parser_escapes.VTable.Hx_Field_exists
								return func(_hx_param_key string) bool {
									return _hx_closure_src_368(((any)(_hx_param_key)))
								}
							}()(s) {
								if strict {
									panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat("Undefined entity: ", s), str, p)))))
								}

								func() func(x string) {
									var _hx_closure_src_369 func(x any)
									_ = _hx_closure_src_369
									_hx_closure_src_369 = buf.VTable.Hx_Field_add
									return func(x string) {
										_hx_closure_src_369(((any)(x)))
									}
								}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("&", s), ";"))
							} else {
								func() func(x HxNullable[string]) {
									var _hx_closure_src_370 func(x any)
									_ = _hx_closure_src_370
									_hx_closure_src_370 = buf.VTable.Hx_Field_add
									return func(x HxNullable[string]) {
										var hx_nullable_f184 HxNullable[string]
										_ = hx_nullable_f184
										hx_nullable_f184 = x
										var _hx_tmp_43 any
										_ = _hx_tmp_43
										if (hx_nullable_f184).Valid == true {
											_hx_tmp_43 = (hx_nullable_f184).Value
										} else {
											_hx_tmp_43 = nil
										}

										_hx_closure_src_370(_hx_tmp_43)
									}
								}()(func() func(_hx_param_key string) HxNullable[string] {
									var _hx_closure_src_371 func(_hx_param_key any) any
									_ = _hx_closure_src_371
									_hx_closure_src_371 = Hx_Field_haxe_xml_parser_escapes.VTable.Hx_Field_get
									return func(_hx_param_key string) HxNullable[string] {
										var hx_nullable_t44 any
										_ = hx_nullable_t44
										hx_nullable_t44 = _hx_closure_src_371(((any)(_hx_param_key)))
										var _hx_tmp_44 HxNullable[string]
										_ = _hx_tmp_44
										if hx_nullable_t44 == nil {
											_hx_tmp_44 = HxNullable[string]{Value: HxStringNull}
										} else {
											_hx_tmp_44 = HxNullable[string]{Value: Hx_Field_go_haxe_hxdynamic_toString(hx_nullable_t44), Valid: true}
										}

										return _hx_tmp_44
									}
								}()(s))
							}
						}

						start = (p + 1)
						state = escapeNext
					} else {
						if !((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45)) && (c != 35) {
							if strict {
								panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat("Invalid character in entity: ", Hx_Field_go_haxe_hxstring_fromCharCode(c)), str, p)))))
							}

							buf.VTable.Hx_Field_addChar(38)
							var _hx_tmp_45 string
							_ = _hx_tmp_45
							_hx_tmp_45 = str
							var _hx_tmp_46 int32
							_ = _hx_tmp_46
							_hx_tmp_46 = start
							buf.VTable.Hx_Field_addSub(_hx_tmp_45, _hx_tmp_46, HxNullable[int32]{Value: (p - start), Valid: true})
							p--
							start = (p + 1)
							state = escapeNext
						}
					}
				}

			}

			p++
		}

		if state == ((int32)(1)) {
			start = p
			state = ((int32)(13))
		}

		if state == ((int32)(13)) {
			if parent.Hx_Field_nodeType == ((int32)(0)) {
				if parent.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(parent.Hx_Field_nodeType))))))
				}

				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Unclosed node <", parent.Hx_Field_nodeName), ">"), str, p)))))
			}

			if p != start {
				var _hx_tmp_47 string
				_ = _hx_tmp_47
				_hx_tmp_47 = str
				var _hx_tmp_48 int32
				_ = _hx_tmp_48
				_hx_tmp_48 = start
				buf.VTable.Hx_Field_addSub(_hx_tmp_47, _hx_tmp_48, HxNullable[int32]{Value: (p - start), Valid: true})
			}

			var str1 string
			_ = str1
			str1 = buf.VTable.Hx_Field_toString()
			var _hx_tmp_49 bool
			_ = _hx_tmp_49
			if int32(utf8.RuneCountInString(str1)) > 0 {
				_hx_tmp_49 = true
			} else {
				_hx_tmp_49 = (nsubs == 0)
			}

			if _hx_tmp_49 {
				parent.VTable.Hx_Field_addChild(Hx_Field_xml_createPCData(buf.VTable.Hx_Field_toString()))
				nsubs++
			}

			return p
		}

		if (!strict && (state == ((int32)(18)))) && (escapeNext == ((int32)(13))) {
			buf.VTable.Hx_Field_addChar(38)
			var _hx_tmp_50 string
			_ = _hx_tmp_50
			_hx_tmp_50 = str
			var _hx_tmp_51 int32
			_ = _hx_tmp_51
			_hx_tmp_51 = start
			buf.VTable.Hx_Field_addSub(_hx_tmp_50, _hx_tmp_51, HxNullable[int32]{Value: (p - start), Valid: true})
			{
				parent.VTable.Hx_Field_addChild(Hx_Field_xml_createPCData(buf.VTable.Hx_Field_toString()))
				nsubs++
			}

			return p
		}

		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_haxe_xml_xmlparserexception_CreateInstance("Unexpected end", str, p)))))
	}
}

func Hx_Field_haxe_xml_parser_isValidChar(c int32) bool {
	return ((((((((c >= 97) && (c <= 122)) || ((c >= 65) && (c <= 90))) || ((c >= 48) && (c <= 57))) || (c == 58)) || (c == 46)) || (c == 95)) || (c == 45))
}
