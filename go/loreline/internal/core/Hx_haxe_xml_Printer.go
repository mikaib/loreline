package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_haxe_xml_printer() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_xml_printer() {
}

type Hx_Obj_VTable_haxe_xml_printer interface {
	Hx_Field_writeNode(value *Hx_Obj_xml, tabs string)
	Hx_Field_write(input string)
	Hx_Field_newline()
	Hx_Field_hasChildren(value *Hx_Obj_xml) bool
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_xml_printer struct {
	VTable          Hx_Obj_VTable_haxe_xml_printer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_output *Hx_Obj_stringbuf
	Hx_Field_pretty bool
}

func Hx_Obj_haxe_xml_printer_CreateEmptyInstance() *Hx_Obj_haxe_xml_printer {
	obj := &Hx_Obj_haxe_xml_printer{}
	obj.VTable = obj
	obj.Hx_Field_output = HxDefault[*Hx_Obj_stringbuf]()
	obj.Hx_Field_pretty = HxDefault[bool]()
	return obj
}

func Hx_Obj_haxe_xml_printer_CreateInstance(pretty bool) *Hx_Obj_haxe_xml_printer {
	obj := Hx_Obj_haxe_xml_printer_CreateEmptyInstance()
	obj.Hx_New(pretty)
	return obj
}

func (this *Hx_Obj_haxe_xml_printer) Hx_New(pretty bool) {
	{
		this.Hx_Field_output = Hx_Obj_stringbuf_CreateInstance()
		this.Hx_Field_pretty = pretty
	}

}

var Hx_Obj_haxe_xml_printer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.xml.Printer",
	HxMakeArray[string]("print"),
	HxMakeArray[string]("output", "pretty", "writeNode", "write", "newline", "hasChildren"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_xml_printer_CreateInstance(HxConvert[bool](params.Get(0)))
	},
	func() any {
		return Hx_Obj_haxe_xml_printer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_xml_printer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "print":
			return Hx_Field_haxe_xml_printer_print
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_xml_printer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_xml_printer_RTTI
}

func (this *Hx_Obj_haxe_xml_printer) Hx_As_Hx_Obj_haxe_xml_printer() *Hx_Obj_haxe_xml_printer {
	return this
}

func (this *Hx_Obj_haxe_xml_printer) Hx_Field_writeNode(value *Hx_Obj_xml, tabs string) {
	switch value.Hx_Field_nodeType {
	case 0:
		{
			func() func(x string) {
				var _hx_closure_src_333 func(x any)
				_ = _hx_closure_src_333
				_hx_closure_src_333 = this.Hx_Field_output.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_333(((any)(x)))
				}
			}()(Hx_Field_go_haxe_hxstring_concat(tabs, "<"))
			{
				if value.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
				}

				func() func(x string) {
					var _hx_closure_src_334 func(x any)
					_ = _hx_closure_src_334
					_hx_closure_src_334 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_334(((any)(x)))
					}
				}()(value.Hx_Field_nodeName)
			}

			{
				var attribute any
				_ = attribute
				attribute = value.VTable.Hx_Field_attributes()
				for {
					var _hx_tmp_0 any
					_ = _hx_tmp_0
					_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_getField(attribute, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_0, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = Hx_Field_go_haxe_hxdynamic_getField(attribute, "next")
					var attribute1 string
					_ = attribute1
					attribute1 = Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_1, HxMakeArray[any]().Dyn()))
					func() func(x string) {
						var _hx_closure_src_335 func(x any)
						_ = _hx_closure_src_335
						_hx_closure_src_335 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_335(((any)(x)))
						}
					}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat(" ", attribute1), "=\""))
					{
						var input string
						_ = input
						input = Hx_Field_stringtools_htmlEscape((value.VTable.Hx_Field_get(attribute1)).Value, true)
						func() func(x string) {
							var _hx_closure_src_336 func(x any)
							_ = _hx_closure_src_336
							_hx_closure_src_336 = this.Hx_Field_output.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_336(((any)(x)))
							}
						}()(input)
					}

					func() func(x string) {
						var _hx_closure_src_337 func(x any)
						_ = _hx_closure_src_337
						_hx_closure_src_337 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_337(((any)(x)))
						}
					}()("\"")
				}
			}

			if this.VTable.Hx_Field_hasChildren(value) {
				func() func(x string) {
					var _hx_closure_src_338 func(x any)
					_ = _hx_closure_src_338
					_hx_closure_src_338 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_338(((any)(x)))
					}
				}()(">")
				if this.Hx_Field_pretty {
					func() func(x string) {
						var _hx_closure_src_339 func(x any)
						_ = _hx_closure_src_339
						_hx_closure_src_339 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_339(((any)(x)))
						}
					}()("\n")
				}

				{
					if (value.Hx_Field_nodeType != Hx_Field_xml_Document) && (value.Hx_Field_nodeType != Hx_Field_xml_Element) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
					}

					var _g_current int32
					_ = _g_current
					_g_current = 0
					var _g_array HxArray[*Hx_Obj_xml]
					_ = _g_array
					_g_array = value.Hx_Field_children
					for {
						var _hx_tmp_2 int32
						_ = _hx_tmp_2
						_hx_tmp_2 = _g_current
						if !(_hx_tmp_2 < _g_array.Len()) {
							break
						}

						var _hx_tmp_3 HxArray[*Hx_Obj_xml]
						_ = _hx_tmp_3
						_hx_tmp_3 = _g_array
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						_hx_tmp_4 = _g_current
						_g_current = (_g_current + 1)
						var child *Hx_Obj_xml
						_ = child
						child = _hx_tmp_3.Get(_hx_tmp_4)
						var _hx_tmp_5 *Hx_Obj_xml
						_ = _hx_tmp_5
						_hx_tmp_5 = child
						var _hx_tmp_6 string
						_ = _hx_tmp_6
						if this.Hx_Field_pretty {
							_hx_tmp_6 = Hx_Field_go_haxe_hxstring_concat(tabs, "\t")
						} else {
							_hx_tmp_6 = tabs
						}

						this.VTable.Hx_Field_writeNode(_hx_tmp_5, _hx_tmp_6)
					}
				}

				func() func(x string) {
					var _hx_closure_src_340 func(x any)
					_ = _hx_closure_src_340
					_hx_closure_src_340 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_340(((any)(x)))
					}
				}()(Hx_Field_go_haxe_hxstring_concat(tabs, "</"))
				{
					if value.Hx_Field_nodeType != Hx_Field_xml_Element {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
					}

					func() func(x string) {
						var _hx_closure_src_341 func(x any)
						_ = _hx_closure_src_341
						_hx_closure_src_341 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_341(((any)(x)))
						}
					}()(value.Hx_Field_nodeName)
				}

				func() func(x string) {
					var _hx_closure_src_342 func(x any)
					_ = _hx_closure_src_342
					_hx_closure_src_342 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_342(((any)(x)))
					}
				}()(">")
				if this.Hx_Field_pretty {
					func() func(x string) {
						var _hx_closure_src_343 func(x any)
						_ = _hx_closure_src_343
						_hx_closure_src_343 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_343(((any)(x)))
						}
					}()("\n")
				}
			} else {
				func() func(x string) {
					var _hx_closure_src_344 func(x any)
					_ = _hx_closure_src_344
					_hx_closure_src_344 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_344(((any)(x)))
					}
				}()("/>")
				if this.Hx_Field_pretty {
					func() func(x string) {
						var _hx_closure_src_345 func(x any)
						_ = _hx_closure_src_345
						_hx_closure_src_345 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_345(((any)(x)))
						}
					}()("\n")
				}
			}
		}

	case 1:
		{
			if (value.Hx_Field_nodeType == Hx_Field_xml_Document) || (value.Hx_Field_nodeType == Hx_Field_xml_Element) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
			}

			var nodeValue string
			_ = nodeValue
			nodeValue = value.Hx_Field_nodeValue
			if int32(utf8.RuneCountInString(nodeValue)) != 0 {
				{
					var _hx_tmp_7 string
					_ = _hx_tmp_7
					_hx_tmp_7 = tabs
					var _hx_tmp_8 string
					_ = _hx_tmp_8
					_hx_tmp_8 = nodeValue
					var input string
					_ = input
					input = Hx_Field_go_haxe_hxstring_concat(_hx_tmp_7, Hx_Field_stringtools_htmlEscape(_hx_tmp_8, nil))
					func() func(x string) {
						var _hx_closure_src_346 func(x any)
						_ = _hx_closure_src_346
						_hx_closure_src_346 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_346(((any)(x)))
						}
					}()(input)
				}

				if this.Hx_Field_pretty {
					func() func(x string) {
						var _hx_closure_src_347 func(x any)
						_ = _hx_closure_src_347
						_hx_closure_src_347 = this.Hx_Field_output.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_347(((any)(x)))
						}
					}()("\n")
				}
			}
		}

	case 2:
		{
			func() func(x string) {
				var _hx_closure_src_348 func(x any)
				_ = _hx_closure_src_348
				_hx_closure_src_348 = this.Hx_Field_output.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_348(((any)(x)))
				}
			}()(Hx_Field_go_haxe_hxstring_concat(tabs, "<![CDATA["))
			{
				if (value.Hx_Field_nodeType == Hx_Field_xml_Document) || (value.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
				}

				func() func(x string) {
					var _hx_closure_src_349 func(x any)
					_ = _hx_closure_src_349
					_hx_closure_src_349 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_349(((any)(x)))
					}
				}()(value.Hx_Field_nodeValue)
			}

			func() func(x string) {
				var _hx_closure_src_350 func(x any)
				_ = _hx_closure_src_350
				_hx_closure_src_350 = this.Hx_Field_output.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_350(((any)(x)))
				}
			}()("]]>")
			if this.Hx_Field_pretty {
				func() func(x string) {
					var _hx_closure_src_351 func(x any)
					_ = _hx_closure_src_351
					_hx_closure_src_351 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_351(((any)(x)))
					}
				}()("\n")
			}
		}

	case 3:
		{
			if (value.Hx_Field_nodeType == Hx_Field_xml_Document) || (value.Hx_Field_nodeType == Hx_Field_xml_Element) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
			}

			var commentContent string
			_ = commentContent
			commentContent = value.Hx_Field_nodeValue
			commentContent = Hx_Obj_ereg_CreateInstance("[\n\r\t]+", "g").VTable.Hx_Field_replace(commentContent, "")
			commentContent = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("<!--", commentContent), "-->")
			func() func(x string) {
				var _hx_closure_src_352 func(x any)
				_ = _hx_closure_src_352
				_hx_closure_src_352 = this.Hx_Field_output.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_352(((any)(x)))
				}
			}()(tabs)
			{
				var input string
				_ = input
				input = Hx_Field_stringtools_trim(commentContent)
				func() func(x string) {
					var _hx_closure_src_353 func(x any)
					_ = _hx_closure_src_353
					_hx_closure_src_353 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_353(((any)(x)))
					}
				}()(input)
			}

			if this.Hx_Field_pretty {
				func() func(x string) {
					var _hx_closure_src_354 func(x any)
					_ = _hx_closure_src_354
					_hx_closure_src_354 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_354(((any)(x)))
					}
				}()("\n")
			}
		}

	case 4:
		{
			{
				if (value.Hx_Field_nodeType == Hx_Field_xml_Document) || (value.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
				}

				func() func(x string) {
					var _hx_closure_src_355 func(x any)
					_ = _hx_closure_src_355
					_hx_closure_src_355 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_355(((any)(x)))
					}
				}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("<!DOCTYPE ", value.Hx_Field_nodeValue), ">"))
			}

			if this.Hx_Field_pretty {
				func() func(x string) {
					var _hx_closure_src_356 func(x any)
					_ = _hx_closure_src_356
					_hx_closure_src_356 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_356(((any)(x)))
					}
				}()("\n")
			}
		}

	case 5:
		{
			{
				if (value.Hx_Field_nodeType == Hx_Field_xml_Document) || (value.Hx_Field_nodeType == Hx_Field_xml_Element) {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
				}

				func() func(x string) {
					var _hx_closure_src_357 func(x any)
					_ = _hx_closure_src_357
					_hx_closure_src_357 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_357(((any)(x)))
					}
				}()(Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("<?", value.Hx_Field_nodeValue), "?>"))
			}

			if this.Hx_Field_pretty {
				func() func(x string) {
					var _hx_closure_src_358 func(x any)
					_ = _hx_closure_src_358
					_hx_closure_src_358 = this.Hx_Field_output.VTable.Hx_Field_add
					return func(x string) {
						_hx_closure_src_358(((any)(x)))
					}
				}()("\n")
			}
		}

	case 6:
		{
			if (value.Hx_Field_nodeType != Hx_Field_xml_Document) && (value.Hx_Field_nodeType != Hx_Field_xml_Element) {
				panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
			}

			var _g_current int32
			_ = _g_current
			_g_current = 0
			var _g_array HxArray[*Hx_Obj_xml]
			_ = _g_array
			_g_array = value.Hx_Field_children
			for {
				var _hx_tmp_9 int32
				_ = _hx_tmp_9
				_hx_tmp_9 = _g_current
				if !(_hx_tmp_9 < _g_array.Len()) {
					break
				}

				var _hx_tmp_10 HxArray[*Hx_Obj_xml]
				_ = _hx_tmp_10
				_hx_tmp_10 = _g_array
				var _hx_tmp_11 int32
				_ = _hx_tmp_11
				_hx_tmp_11 = _g_current
				_g_current = (_g_current + 1)
				var child *Hx_Obj_xml
				_ = child
				child = _hx_tmp_10.Get(_hx_tmp_11)
				this.VTable.Hx_Field_writeNode(child, tabs)
			}
		}

	}
}

func (this *Hx_Obj_haxe_xml_printer) Hx_Field_write(input string) {
	func() func(x string) {
		var _hx_closure_src_359 func(x any)
		_ = _hx_closure_src_359
		_hx_closure_src_359 = this.Hx_Field_output.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_359(((any)(x)))
		}
	}()(input)
}

func (this *Hx_Obj_haxe_xml_printer) Hx_Field_newline() {
	if this.Hx_Field_pretty {
		func() func(x string) {
			var _hx_closure_src_360 func(x any)
			_ = _hx_closure_src_360
			_hx_closure_src_360 = this.Hx_Field_output.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_360(((any)(x)))
			}
		}()("\n")
	}
}

func (this *Hx_Obj_haxe_xml_printer) Hx_Field_hasChildren(value *Hx_Obj_xml) bool {
	{
		if (value.Hx_Field_nodeType != Hx_Field_xml_Document) && (value.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(value.Hx_Field_nodeType))))))
		}

		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[*Hx_Obj_xml]
		_ = _g_array
		_g_array = value.Hx_Field_children
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
			var child *Hx_Obj_xml
			_ = child
			child = _hx_tmp_1.Get(_hx_tmp_2)
			switch child.Hx_Field_nodeType {
			case 0, 1:
				{
					return true
				}

			case 2, 3:
				{
					if (child.Hx_Field_nodeType == Hx_Field_xml_Document) || (child.Hx_Field_nodeType == Hx_Field_xml_Element) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(child.Hx_Field_nodeType))))))
					}

					var s string
					_ = s
					s = child.Hx_Field_nodeValue
					var l int32
					_ = l
					l = int32(utf8.RuneCountInString(s))
					var r int32
					_ = r
					r = 0
					for {
						var _hx_tmp_3 bool
						_ = _hx_tmp_3
						if r < l {
							_hx_tmp_3 = Hx_Field_stringtools_isSpace(s, r)
						} else {
							_hx_tmp_3 = false
						}

						if !(_hx_tmp_3) {
							break
						}

						r++
					}

					var _hx_tmp_4 string
					_ = _hx_tmp_4
					if r > 0 {
						var _hx_tmp_5 string
						_ = _hx_tmp_5
						_hx_tmp_5 = s
						var _hx_tmp_6 int32
						_ = _hx_tmp_6
						_hx_tmp_6 = r
						var _hx_tmp_7 string
						_ = _hx_tmp_7
						_hx_tmp_7 = _hx_tmp_5
						var _hx_tmp_8 int32
						_ = _hx_tmp_8
						_hx_tmp_8 = _hx_tmp_6
						_hx_tmp_4 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_7, _hx_tmp_8, HxNullable[int32]{Value: (l - r), Valid: true}))
					} else {
						_hx_tmp_4 = s
					}

					if int32(utf8.RuneCountInString((_hx_tmp_4))) != 0 {
						return true
					}
				}

			default:
				{
				}
			}
		}
	}

	return false
}

func Hx_Field_haxe_xml_printer_print(xml *Hx_Obj_xml, pretty__isOptional any) string {
	pretty := HxDefault[bool]()
	_ = pretty

	if pretty__isOptional != nil {
		pretty = HxConvert[bool](pretty__isOptional)
	} else {
		pretty = false
	}
	{
		var printer *Hx_Obj_haxe_xml_printer
		_ = printer
		printer = Hx_Obj_haxe_xml_printer_CreateInstance(pretty)
		printer.VTable.Hx_Field_writeNode(xml, "")
		return printer.Hx_Field_output.VTable.Hx_Field_toString()
	}
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_xml]()
}
