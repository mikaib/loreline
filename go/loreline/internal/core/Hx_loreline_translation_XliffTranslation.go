package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_translation_xlifftranslation() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_translation_xlifftranslation() {
}

type Hx_Obj_VTable_loreline_translation_xlifftranslation interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_translation_xlifftranslation struct {
	VTable Hx_Obj_VTable_loreline_translation_xlifftranslation `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_translation_xlifftranslation_CreateEmptyInstance() *Hx_Obj_loreline_translation_xlifftranslation {
	obj := &Hx_Obj_loreline_translation_xlifftranslation{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_translation_xlifftranslation_CreateInstance() *Hx_Obj_loreline_translation_xlifftranslation {
	obj := Hx_Obj_loreline_translation_xlifftranslation_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_translation_xlifftranslation_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.translation.XliffTranslation",
	HxMakeArray[string]("toLoreline", "localeMatches", "primarySubtag", "firstChildElement", "innerText", "fromScript", "escapeForXmlText", "escapeForXmlAttr", "writeEntry", "escapeLineForComment"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_translation_xlifftranslation_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_translation_xlifftranslation_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_translation_xlifftranslation_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "toLoreline":
			return Hx_Field_loreline_translation_xlifftranslation_toLoreline
		case "localeMatches":
			return Hx_Field_loreline_translation_xlifftranslation_localeMatches
		case "primarySubtag":
			return Hx_Field_loreline_translation_xlifftranslation_primarySubtag
		case "firstChildElement":
			return Hx_Field_loreline_translation_xlifftranslation_firstChildElement
		case "innerText":
			return Hx_Field_loreline_translation_xlifftranslation_innerText
		case "fromScript":
			return Hx_Field_loreline_translation_xlifftranslation_fromScript
		case "escapeForXmlText":
			return Hx_Field_loreline_translation_xlifftranslation_escapeForXmlText
		case "escapeForXmlAttr":
			return Hx_Field_loreline_translation_xlifftranslation_escapeForXmlAttr
		case "writeEntry":
			return Hx_Field_loreline_translation_xlifftranslation_writeEntry
		case "escapeLineForComment":
			return Hx_Field_loreline_translation_xlifftranslation_escapeLineForComment
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_translation_xlifftranslation) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_translation_xlifftranslation_RTTI
}

func (this *Hx_Obj_loreline_translation_xlifftranslation) Hx_As_Hx_Obj_loreline_translation_xlifftranslation() *Hx_Obj_loreline_translation_xlifftranslation {
	return this
}

func Hx_Field_loreline_translation_xlifftranslation_toLoreline(content string, locale string) string {
	var root *Hx_Obj_xml
	_ = root
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
							panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance(Hx_Field_go_haxe_hxstring_concat("Invalid XLIFF: ", Hx_Field_std_string(e)), nil)))))
						}
					} else {
						panic(_hx_ex)
					}
				}
			}()
			{
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = content
				root = Hx_Field_haxe_xml_parser_parse(_hx_tmp_0, nil)
			}
		}()
		if hx_try_state == 1 {
			return hx_try_return
		}
	}

	var xliff *Hx_Obj_xml
	_ = xliff
	{
		if (root.Hx_Field_nodeType != Hx_Field_xml_Document) && (root.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(root.Hx_Field_nodeType))))))
		}

		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[*Hx_Obj_xml]
		_ = _g_array
		_g_array = root.Hx_Field_children
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g_current
			if !(_hx_tmp_1 < _g_array.Len()) {
				break
			}

			var _hx_tmp_2 HxArray[*Hx_Obj_xml]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g_array
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g_current
			_g_current = (_g_current + 1)
			var child *Hx_Obj_xml
			_ = child
			child = _hx_tmp_2.Get(_hx_tmp_3)
			if child.Hx_Field_nodeType == ((int32)(0)) {
				if child.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(child.Hx_Field_nodeType))))))
				}

				if child.Hx_Field_nodeName == "xliff" {
					xliff = child
					break
				}
			}
		}
	}

	if xliff == nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Obj_loreline_error_CreateInstance("Invalid XLIFF: missing <xliff> root element", nil)))))
	}

	var _hx_tmp_4 HxNullable[string]
	_ = _hx_tmp_4
	if xliff.VTable.Hx_Field_exists("trgLang") {
		_hx_tmp_4 = xliff.VTable.Hx_Field_get("trgLang")
	} else {
		_hx_tmp_4 = HxNullable[string]{Value: HxStringNull}
	}

	var xliffTrgLang HxNullable[string]
	_ = xliffTrgLang
	xliffTrgLang = _hx_tmp_4
	var is2x bool
	_ = is2x
	is2x = ((xliffTrgLang).Valid != false)
	var _hx_tmp_5 bool
	_ = _hx_tmp_5
	if is2x {
		_hx_tmp_5 = !Hx_Field_loreline_translation_xlifftranslation_localeMatches((xliffTrgLang).Value, locale)
	} else {
		_hx_tmp_5 = false
	}

	if _hx_tmp_5 {
		return ""
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	var first bool
	_ = first
	first = true
	{
		var _gfile any
		_ = _gfile
		_gfile = xliff.VTable.Hx_Field_elementsNamed("file")
		for {
			var _hx_tmp_6 any
			_ = _hx_tmp_6
			_hx_tmp_6 = Hx_Field_go_haxe_hxdynamic_getField(_gfile, "hasNext")
			if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_6, HxMakeArray[any]().Dyn()))) {
				break
			}

			var _hx_tmp_7 any
			_ = _hx_tmp_7
			_hx_tmp_7 = Hx_Field_go_haxe_hxdynamic_getField(_gfile, "next")
			var hx_dyncast_363 any
			_ = hx_dyncast_363
			hx_dyncast_363 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_7, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
			var _hx_tmp_8 *Hx_Obj_xml
			_ = _hx_tmp_8
			if hx_dyncast_363 != nil {
				_hx_tmp_8 = hx_dyncast_363.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
			} else {
				_hx_tmp_8 = nil
			}

			var file *Hx_Obj_xml
			_ = file
			file = _hx_tmp_8
			if !is2x {
				var _hx_tmp_9 HxNullable[string]
				_ = _hx_tmp_9
				if file.VTable.Hx_Field_exists("target-language") {
					_hx_tmp_9 = file.VTable.Hx_Field_get("target-language")
				} else {
					_hx_tmp_9 = HxNullable[string]{Value: HxStringNull}
				}

				var fileTrgLang HxNullable[string]
				_ = fileTrgLang
				fileTrgLang = _hx_tmp_9
				var _hx_tmp_10 bool
				_ = _hx_tmp_10
				if (fileTrgLang).Valid != false {
					_hx_tmp_10 = !Hx_Field_loreline_translation_xlifftranslation_localeMatches((fileTrgLang).Value, locale)
				} else {
					_hx_tmp_10 = false
				}

				if _hx_tmp_10 {
					continue
				}
			}

			if is2x {
				var _gunit any
				_ = _gunit
				_gunit = file.VTable.Hx_Field_elementsNamed("unit")
				for {
					var _hx_tmp_11 any
					_ = _hx_tmp_11
					_hx_tmp_11 = Hx_Field_go_haxe_hxdynamic_getField(_gunit, "hasNext")
					if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_11, HxMakeArray[any]().Dyn()))) {
						break
					}

					var _hx_tmp_12 any
					_ = _hx_tmp_12
					_hx_tmp_12 = Hx_Field_go_haxe_hxdynamic_getField(_gunit, "next")
					var hx_dyncast_364 any
					_ = hx_dyncast_364
					hx_dyncast_364 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_12, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
					var _hx_tmp_13 *Hx_Obj_xml
					_ = _hx_tmp_13
					if hx_dyncast_364 != nil {
						_hx_tmp_13 = hx_dyncast_364.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
					} else {
						_hx_tmp_13 = nil
					}

					var unit *Hx_Obj_xml
					_ = unit
					unit = _hx_tmp_13
					if !unit.VTable.Hx_Field_exists("id") {
						continue
					}

					var id HxNullable[string]
					_ = id
					id = unit.VTable.Hx_Field_get("id")
					var _g HxArray[*Hx_Obj_xml]
					_ = _g
					_g = HxMakeArray[*Hx_Obj_xml]()
					{
						var _gs any
						_ = _gs
						_gs = unit.VTable.Hx_Field_elementsNamed("segment")
						for {
							var _hx_tmp_14 any
							_ = _hx_tmp_14
							_hx_tmp_14 = Hx_Field_go_haxe_hxdynamic_getField(_gs, "hasNext")
							if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_14, HxMakeArray[any]().Dyn()))) {
								break
							}

							var _hx_tmp_15 any
							_ = _hx_tmp_15
							_hx_tmp_15 = Hx_Field_go_haxe_hxdynamic_getField(_gs, "next")
							var hx_dyncast_365 any
							_ = hx_dyncast_365
							hx_dyncast_365 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_15, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
							var _hx_tmp_16 *Hx_Obj_xml
							_ = _hx_tmp_16
							if hx_dyncast_365 != nil {
								_hx_tmp_16 = hx_dyncast_365.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
							} else {
								_hx_tmp_16 = nil
							}

							var s *Hx_Obj_xml
							_ = s
							s = _hx_tmp_16
							_ = (Hx_Array_Push[*Hx_Obj_xml](_g, s))
						}
					}

					var segs HxArray[*Hx_Obj_xml]
					_ = segs
					segs = _g
					if segs.Len() == 0 {
						continue
					}

					var sourceParts HxArray[string]
					_ = sourceParts
					sourceParts = HxMakeArray[string]()
					var targetParts HxArray[string]
					_ = targetParts
					targetParts = HxMakeArray[string]()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						for {
							var _hx_tmp_17 int32
							_ = _hx_tmp_17
							_hx_tmp_17 = _g1
							if !(_hx_tmp_17 < segs.Len()) {
								break
							}

							var seg *Hx_Obj_xml
							_ = seg
							seg = segs.Get(_g1)
							_g1++
							var s string
							_ = s
							s = Hx_Field_loreline_translation_xlifftranslation_innerText(Hx_Field_loreline_translation_xlifftranslation_firstChildElement(seg, "source"))
							var t string
							_ = t
							t = Hx_Field_loreline_translation_xlifftranslation_innerText(Hx_Field_loreline_translation_xlifftranslation_firstChildElement(seg, "target"))
							if s != HxStringNull {
								_ = (Hx_Array_Push[string](sourceParts, s))
							}

							if t != HxStringNull {
								_ = (Hx_Array_Push[string](targetParts, t))
							}
						}
					}

					var src string
					_ = src
					src = (Hx_Array_Join[string](sourceParts, ""))
					var tgt string
					_ = tgt
					tgt = (Hx_Array_Join[string](targetParts, ""))
					if tgt == "" {
						continue
					}

					if !first {
						func() func(x string) {
							var _hx_closure_src_807 func(x any)
							_ = _hx_closure_src_807
							_hx_closure_src_807 = buf.VTable.Hx_Field_add
							return func(x string) {
								_hx_closure_src_807(((any)(x)))
							}
						}()("\n")
					}

					first = false
					var _hx_tmp_18 *Hx_Obj_stringbuf
					_ = _hx_tmp_18
					_hx_tmp_18 = buf
					Hx_Field_loreline_translation_xlifftranslation_writeEntry(_hx_tmp_18, (id).Value, src, tgt)
				}
			} else {
				var body *Hx_Obj_xml
				_ = body
				body = Hx_Field_loreline_translation_xlifftranslation_firstChildElement(file, "body")
				if body == nil {
					continue
				}

				{
					var _gunit any
					_ = _gunit
					_gunit = body.VTable.Hx_Field_elementsNamed("trans-unit")
					for {
						var _hx_tmp_19 any
						_ = _hx_tmp_19
						_hx_tmp_19 = Hx_Field_go_haxe_hxdynamic_getField(_gunit, "hasNext")
						if !Hx_Field_go_haxe_hxdynamic_toBool((Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_19, HxMakeArray[any]().Dyn()))) {
							break
						}

						var _hx_tmp_20 any
						_ = _hx_tmp_20
						_hx_tmp_20 = Hx_Field_go_haxe_hxdynamic_getField(_gunit, "next")
						var hx_dyncast_366 any
						_ = hx_dyncast_366
						hx_dyncast_366 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_call(_hx_tmp_20, HxMakeArray[any]().Dyn()), "Hx_Obj_xml")
						var _hx_tmp_21 *Hx_Obj_xml
						_ = _hx_tmp_21
						if hx_dyncast_366 != nil {
							_hx_tmp_21 = hx_dyncast_366.(interface{ Hx_As_Hx_Obj_xml() *Hx_Obj_xml }).Hx_As_Hx_Obj_xml()
						} else {
							_hx_tmp_21 = nil
						}

						var unit *Hx_Obj_xml
						_ = unit
						unit = _hx_tmp_21
						if !unit.VTable.Hx_Field_exists("id") {
							continue
						}

						var id HxNullable[string]
						_ = id
						id = unit.VTable.Hx_Field_get("id")
						var src string
						_ = src
						src = Hx_Field_loreline_translation_xlifftranslation_innerText(Hx_Field_loreline_translation_xlifftranslation_firstChildElement(unit, "source"))
						var tgt string
						_ = tgt
						tgt = Hx_Field_loreline_translation_xlifftranslation_innerText(Hx_Field_loreline_translation_xlifftranslation_firstChildElement(unit, "target"))
						if (tgt == HxStringNull) || (tgt == "") {
							continue
						}

						if !first {
							func() func(x string) {
								var _hx_closure_src_808 func(x any)
								_ = _hx_closure_src_808
								_hx_closure_src_808 = buf.VTable.Hx_Field_add
								return func(x string) {
									_hx_closure_src_808(((any)(x)))
								}
							}()("\n")
						}

						first = false
						var _hx_tmp_22 *Hx_Obj_stringbuf
						_ = _hx_tmp_22
						_hx_tmp_22 = buf
						var _hx_tmp_23 string
						_ = _hx_tmp_23
						_hx_tmp_23 = (id).Value
						var _hx_tmp_24 string
						_ = _hx_tmp_24
						if src != HxStringNull {
							_hx_tmp_24 = src
						} else {
							_hx_tmp_24 = ""
						}

						Hx_Field_loreline_translation_xlifftranslation_writeEntry(_hx_tmp_22, _hx_tmp_23, _hx_tmp_24, tgt)
					}
				}
			}
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_xlifftranslation_localeMatches(fileLocale string, requested string) bool {
	if fileLocale == requested {
		return true
	}

	var fileLow string
	_ = fileLow
	fileLow = Hx_Field_go_haxe_hxstring_toLowerCase(fileLocale)
	var reqLow string
	_ = reqLow
	reqLow = Hx_Field_go_haxe_hxstring_toLowerCase(requested)
	if fileLow == reqLow {
		return true
	}

	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = fileLow
	var dash int32
	_ = dash
	dash = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "-", HxNullable[int32]{}))
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	if dash < 0 {
		_hx_tmp_1 = fileLow
	} else {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		if (HxNullable[int32]{Value: dash, Valid: true}).Valid != false {
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = fileLow
			var _hx_tmp_4 string
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_tmp_3
			_hx_tmp_2 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_4, 0, HxNullable[int32]{Value: dash, Valid: true}))
		} else {
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = fileLow
			var _hx_tmp_6 string
			_ = _hx_tmp_6
			_hx_tmp_6 = _hx_tmp_5
			_hx_tmp_2 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_6, 0, HxNullable[int32]{}))
		}

		_hx_tmp_1 = _hx_tmp_2
	}

	var filePrimary string
	_ = filePrimary
	filePrimary = _hx_tmp_1
	var _hx_tmp_7 string
	_ = _hx_tmp_7
	_hx_tmp_7 = reqLow
	var dash1 int32
	_ = dash1
	dash1 = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_7, "-", HxNullable[int32]{}))
	var _hx_tmp_8 string
	_ = _hx_tmp_8
	if dash1 < 0 {
		_hx_tmp_8 = reqLow
	} else {
		var _hx_tmp_9 string
		_ = _hx_tmp_9
		if (HxNullable[int32]{Value: dash1, Valid: true}).Valid != false {
			var _hx_tmp_10 string
			_ = _hx_tmp_10
			_hx_tmp_10 = reqLow
			var _hx_tmp_11 string
			_ = _hx_tmp_11
			_hx_tmp_11 = _hx_tmp_10
			_hx_tmp_9 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_11, 0, HxNullable[int32]{Value: dash1, Valid: true}))
		} else {
			var _hx_tmp_12 string
			_ = _hx_tmp_12
			_hx_tmp_12 = reqLow
			var _hx_tmp_13 string
			_ = _hx_tmp_13
			_hx_tmp_13 = _hx_tmp_12
			_hx_tmp_9 = (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_13, 0, HxNullable[int32]{}))
		}

		_hx_tmp_8 = _hx_tmp_9
	}

	var reqPrimary string
	_ = reqPrimary
	reqPrimary = _hx_tmp_8
	return (filePrimary == reqPrimary)
}

func Hx_Field_loreline_translation_xlifftranslation_primarySubtag(s string) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = s
	var dash int32
	_ = dash
	dash = (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, "-", HxNullable[int32]{}))
	if dash < 0 {
		return s
	}

	if (HxNullable[int32]{Value: dash, Valid: true}).Valid != false {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = s
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_1, 0, HxNullable[int32]{Value: dash, Valid: true}))
	} else {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = s
		return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_2, 0, HxNullable[int32]{}))
	}
}

func Hx_Field_loreline_translation_xlifftranslation_firstChildElement(parent *Hx_Obj_xml, name string) *Hx_Obj_xml {
	{
		if (parent.Hx_Field_nodeType != Hx_Field_xml_Document) && (parent.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(parent.Hx_Field_nodeType))))))
		}

		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[*Hx_Obj_xml]
		_ = _g_array
		_g_array = parent.Hx_Field_children
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
			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_1.Get(_hx_tmp_2)
			if c.Hx_Field_nodeType == ((int32)(0)) {
				if c.Hx_Field_nodeType != Hx_Field_xml_Element {
					panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element but found ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
				}

				if c.Hx_Field_nodeName == name {
					return c
				}
			}
		}
	}

	return nil
}

func Hx_Field_loreline_translation_xlifftranslation_innerText(node *Hx_Obj_xml) string {
	if node == nil {
		return HxStringNull
	}

	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	{
		if (node.Hx_Field_nodeType != Hx_Field_xml_Document) && (node.Hx_Field_nodeType != Hx_Field_xml_Element) {
			panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, expected Element or Document but found ", Hx_Field__xml_xmltype_impl__toString(node.Hx_Field_nodeType))))))
		}

		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[*Hx_Obj_xml]
		_ = _g_array
		_g_array = node.Hx_Field_children
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
			var c *Hx_Obj_xml
			_ = c
			c = _hx_tmp_1.Get(_hx_tmp_2)
			switch c.Hx_Field_nodeType {
			case 0:
				{
					func() func(x string) {
						var _hx_closure_src_809 func(x any)
						_ = _hx_closure_src_809
						_hx_closure_src_809 = buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_809(((any)(x)))
						}
					}()(Hx_Field_loreline_translation_xlifftranslation_innerText(c))
				}

			case 1:
				{
					if (c.Hx_Field_nodeType == Hx_Field_xml_Document) || (c.Hx_Field_nodeType == Hx_Field_xml_Element) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					func() func(x string) {
						var _hx_closure_src_810 func(x any)
						_ = _hx_closure_src_810
						_hx_closure_src_810 = buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_810(((any)(x)))
						}
					}()(c.Hx_Field_nodeValue)
				}

			case 2:
				{
					if (c.Hx_Field_nodeType == Hx_Field_xml_Document) || (c.Hx_Field_nodeType == Hx_Field_xml_Element) {
						panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Bad node type, unexpected ", Hx_Field__xml_xmltype_impl__toString(c.Hx_Field_nodeType))))))
					}

					func() func(x string) {
						var _hx_closure_src_811 func(x any)
						_ = _hx_closure_src_811
						_hx_closure_src_811 = buf.VTable.Hx_Field_add
						return func(x string) {
							_hx_closure_src_811(((any)(x)))
						}
					}()(c.Hx_Field_nodeValue)
				}

			default:
				{
				}
			}
		}
	}

	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_xlifftranslation_fromScript(script *Hx_Obj_loreline_astnode, existing Hx_Typedef_map, locale string) string {
	var entries HxArray[any]
	_ = entries
	entries = Hx_Field_loreline_astutils_extractTranslatableEntries(script)
	var printer *Hx_Obj_loreline_printer
	_ = printer
	printer = Hx_Obj_loreline_printer_CreateInstance(nil, nil)
	var buf *Hx_Obj_stringbuf
	_ = buf
	buf = Hx_Obj_stringbuf_CreateInstance()
	func() func(x string) {
		var _hx_closure_src_812 func(x any)
		_ = _hx_closure_src_812
		_hx_closure_src_812 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_812(((any)(x)))
		}
	}()("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	func() func(x string) {
		var _hx_closure_src_813 func(x any)
		_ = _hx_closure_src_813
		_hx_closure_src_813 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_813(((any)(x)))
		}
	}()("<xliff version=\"1.2\" xmlns=\"urn:oasis:names:tc:xliff:document:1.2\">\n")
	func() func(x string) {
		var _hx_closure_src_814 func(x any)
		_ = _hx_closure_src_814
		_hx_closure_src_814 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_814(((any)(x)))
		}
	}()("  <file source-language=\"en\" target-language=\"")
	func() func(x string) {
		var _hx_closure_src_815 func(x any)
		_ = _hx_closure_src_815
		_hx_closure_src_815 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_815(((any)(x)))
		}
	}()(Hx_Field_loreline_translation_xlifftranslation_escapeForXmlAttr(locale))
	func() func(x string) {
		var _hx_closure_src_816 func(x any)
		_ = _hx_closure_src_816
		_hx_closure_src_816 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_816(((any)(x)))
		}
	}()("\" datatype=\"plaintext\">\n")
	func() func(x string) {
		var _hx_closure_src_817 func(x any)
		_ = _hx_closure_src_817
		_hx_closure_src_817 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_817(((any)(x)))
		}
	}()("    <body>\n")
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
			var hx_dyncast_367 any
			_ = hx_dyncast_367
			hx_dyncast_367 = Hx_Field_go_haxe_hxdynamic_castClass(Hx_Field_go_haxe_hxdynamic_getField(entry, "str"), "Hx_Obj_loreline_nstringliteral")
			var _hx_tmp_1 *Hx_Obj_loreline_nstringliteral
			_ = _hx_tmp_1
			if hx_dyncast_367 != nil {
				_hx_tmp_1 = hx_dyncast_367.(interface {
					Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
				}).Hx_As_Hx_Obj_loreline_nstringliteral()
			} else {
				_hx_tmp_1 = nil
			}

			var source string
			_ = source
			source = printer.VTable.Hx_Field_printStringLiteralAsText(_hx_tmp_1)
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			var _hx_tmp_3 bool
			_ = _hx_tmp_3
			if existing != nil {
				_hx_tmp_3 = func() func(_hx_param_key string) bool {
					var _hx_closure_src_818 func(_hx_param_key any) bool
					_ = _hx_closure_src_818
					_hx_closure_src_818 = existing.VTable.Hx_Field_exists
					return func(_hx_param_key string) bool {
						return _hx_closure_src_818(((any)(_hx_param_key)))
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id")))
			} else {
				_hx_tmp_3 = false
			}

			if _hx_tmp_3 {
				_hx_tmp_2 = printer.VTable.Hx_Field_printStringLiteralAsText(func() func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
					var _hx_closure_src_819 func(_hx_param_key any) any
					_ = _hx_closure_src_819
					_hx_closure_src_819 = existing.VTable.Hx_Field_get
					return func(_hx_param_key string) *Hx_Obj_loreline_nstringliteral {
						var hx_dyncast_368 any
						_ = hx_dyncast_368
						hx_dyncast_368 = Hx_Field_go_haxe_hxdynamic_castClass(_hx_closure_src_819(((any)(_hx_param_key))), "Hx_Obj_loreline_nstringliteral")
						var _hx_tmp_4 *Hx_Obj_loreline_nstringliteral
						_ = _hx_tmp_4
						if hx_dyncast_368 != nil {
							_hx_tmp_4 = hx_dyncast_368.(interface {
								Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral
							}).Hx_As_Hx_Obj_loreline_nstringliteral()
						} else {
							_hx_tmp_4 = nil
						}

						return _hx_tmp_4
					}
				}()(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id"))))
			} else {
				_hx_tmp_2 = source
			}

			var translation string
			_ = translation
			translation = _hx_tmp_2
			func() func(x string) {
				var _hx_closure_src_820 func(x any)
				_ = _hx_closure_src_820
				_hx_closure_src_820 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_820(((any)(x)))
				}
			}()("      <trans-unit id=\"")
			func() func(x string) {
				var _hx_closure_src_821 func(x any)
				_ = _hx_closure_src_821
				_hx_closure_src_821 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_821(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_xlifftranslation_escapeForXmlAttr(Hx_Field_go_haxe_hxdynamic_toString(Hx_Field_go_haxe_hxdynamic_getField(entry, "id"))))
			func() func(x string) {
				var _hx_closure_src_822 func(x any)
				_ = _hx_closure_src_822
				_hx_closure_src_822 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_822(((any)(x)))
				}
			}()("\">\n")
			func() func(x string) {
				var _hx_closure_src_823 func(x any)
				_ = _hx_closure_src_823
				_hx_closure_src_823 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_823(((any)(x)))
				}
			}()("        <source>")
			func() func(x string) {
				var _hx_closure_src_824 func(x any)
				_ = _hx_closure_src_824
				_hx_closure_src_824 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_824(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_xlifftranslation_escapeForXmlText(source))
			func() func(x string) {
				var _hx_closure_src_825 func(x any)
				_ = _hx_closure_src_825
				_hx_closure_src_825 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_825(((any)(x)))
				}
			}()("</source>\n")
			func() func(x string) {
				var _hx_closure_src_826 func(x any)
				_ = _hx_closure_src_826
				_hx_closure_src_826 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_826(((any)(x)))
				}
			}()("        <target>")
			func() func(x string) {
				var _hx_closure_src_827 func(x any)
				_ = _hx_closure_src_827
				_hx_closure_src_827 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_827(((any)(x)))
				}
			}()(Hx_Field_loreline_translation_xlifftranslation_escapeForXmlText(translation))
			func() func(x string) {
				var _hx_closure_src_828 func(x any)
				_ = _hx_closure_src_828
				_hx_closure_src_828 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_828(((any)(x)))
				}
			}()("</target>\n")
			func() func(x string) {
				var _hx_closure_src_829 func(x any)
				_ = _hx_closure_src_829
				_hx_closure_src_829 = buf.VTable.Hx_Field_add
				return func(x string) {
					_hx_closure_src_829(((any)(x)))
				}
			}()("      </trans-unit>\n")
		}
	}

	func() func(x string) {
		var _hx_closure_src_830 func(x any)
		_ = _hx_closure_src_830
		_hx_closure_src_830 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_830(((any)(x)))
		}
	}()("    </body>\n")
	func() func(x string) {
		var _hx_closure_src_831 func(x any)
		_ = _hx_closure_src_831
		_hx_closure_src_831 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_831(((any)(x)))
		}
	}()("  </file>\n")
	func() func(x string) {
		var _hx_closure_src_832 func(x any)
		_ = _hx_closure_src_832
		_hx_closure_src_832 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_832(((any)(x)))
		}
	}()("</xliff>\n")
	return buf.VTable.Hx_Field_toString()
}

func Hx_Field_loreline_translation_xlifftranslation_escapeForXmlText(s string) string {
	var out string
	_ = out
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "&")), "&amp;"))
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(out, "<")), "&lt;"))
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(out, ">")), "&gt;"))
	return out
}

func Hx_Field_loreline_translation_xlifftranslation_escapeForXmlAttr(s string) string {
	var out string
	_ = out
	out = Hx_Field_loreline_translation_xlifftranslation_escapeForXmlText(s)
	out = (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(out, "\"")), "&quot;"))
	return out
}

func Hx_Field_loreline_translation_xlifftranslation_writeEntry(buf Hx_Typedef_loreline_utf8buf, id string, src string, tgt string) {
	func() func(x string) {
		var _hx_closure_src_833 func(x any)
		_ = _hx_closure_src_833
		_hx_closure_src_833 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_833(((any)(x)))
		}
	}()("#")
	func() func(x string) {
		var _hx_closure_src_834 func(x any)
		_ = _hx_closure_src_834
		_hx_closure_src_834 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_834(((any)(x)))
		}
	}()(id)
	if (src != HxStringNull) && (src != "") {
		func() func(x string) {
			var _hx_closure_src_835 func(x any)
			_ = _hx_closure_src_835
			_hx_closure_src_835 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_835(((any)(x)))
			}
		}()(" // ")
		func() func(x string) {
			var _hx_closure_src_836 func(x any)
			_ = _hx_closure_src_836
			_hx_closure_src_836 = buf.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_836(((any)(x)))
			}
		}()((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(src, "\n")), " ")), "\r")), "")))
	}

	func() func(x string) {
		var _hx_closure_src_837 func(x any)
		_ = _hx_closure_src_837
		_hx_closure_src_837 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_837(((any)(x)))
		}
	}()("\n")
	func() func(x string) {
		var _hx_closure_src_838 func(x any)
		_ = _hx_closure_src_838
		_hx_closure_src_838 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_838(((any)(x)))
		}
	}()(tgt)
	func() func(x string) {
		var _hx_closure_src_839 func(x any)
		_ = _hx_closure_src_839
		_hx_closure_src_839 = buf.VTable.Hx_Field_add
		return func(x string) {
			_hx_closure_src_839(((any)(x)))
		}
	}()("\n")
}

func Hx_Field_loreline_translation_xlifftranslation_escapeLineForComment(s string) string {
	return (Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split((Hx_Array_Join[string]((Hx_Field_go_haxe_hxstring_split(s, "\n")), " ")), "\r")), ""))
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_xml]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
