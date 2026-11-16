package core

import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_position() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_position() {
}

type Hx_Obj_VTable_loreline_position interface {
	Hx_Field_toString() string
	Hx_Field_toJson() any
	Hx_Field_withOffset(content string, additionalOffset int32, newLength__isOptional any, contentStart__isOptional any) *Hx_Obj_loreline_position
	Hx_Field_extendedTo(endPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_position
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_position struct {
	VTable          Hx_Obj_VTable_loreline_position `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_line   int32
	Hx_Field_column int32
	Hx_Field_offset int32
	Hx_Field_length int32
}

func Hx_Obj_loreline_position_CreateEmptyInstance() *Hx_Obj_loreline_position {
	obj := &Hx_Obj_loreline_position{}
	obj.VTable = obj
	obj.Hx_Field_line = HxDefault[int32]()
	obj.Hx_Field_column = HxDefault[int32]()
	obj.Hx_Field_offset = HxDefault[int32]()
	obj.Hx_Field_length = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_position_CreateInstance(line int32, column int32, offset int32, length__isOptional any) *Hx_Obj_loreline_position {
	obj := Hx_Obj_loreline_position_CreateEmptyInstance()
	obj.Hx_New(line, column, offset, length__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_position) Hx_New(line int32, column int32, offset int32, length__isOptional any) {
	{
		length := HxDefault[int32]()
		_ = length

		if length__isOptional != nil {
			length = HxConvert[int32](length__isOptional)
		} else {
			length = 0
		}
		{
			this.Hx_Field_line = line
			this.Hx_Field_column = column
			this.Hx_Field_offset = offset
			this.Hx_Field_length = length
		}

	}
}

var Hx_Obj_loreline_position_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Position",
	HxMakeArray[string]("fromContentAndIndex"),
	HxMakeArray[string]("line", "column", "offset", "length", "toString", "toJson", "withOffset", "extendedTo"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_position_CreateInstance(HxConvert[int32](params.Get(0)), HxConvert[int32](params.Get(1)), HxConvert[int32](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_position_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_position_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "fromContentAndIndex":
			return Hx_Field_loreline_position_fromContentAndIndex
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_position) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_position_RTTI
}

func (this *Hx_Obj_loreline_position) Hx_As_Hx_Obj_loreline_position() *Hx_Obj_loreline_position {
	return this
}

func (this *Hx_Obj_loreline_position) Hx_Field_toString() string {
	var _hx_tmp_2 string
	_ = _hx_tmp_2
	_hx_tmp_2 = Hx_Field_go_haxe_hxstring_concat(("(" + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_line)), ":")
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_go_haxe_hxstring_concat((_hx_tmp_2 + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_column)), ":")
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat((_hx_tmp_1 + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_offset)), ":")
	return Hx_Field_go_haxe_hxstring_concat((_hx_tmp_0 + Hx_Field_go_haxe_hxdynamic_toString(this.Hx_Field_length)), ")")
}

func (this *Hx_Obj_loreline_position) Hx_Field_toJson() any {
	var json any
	_ = json
	json = any(map[string]any{"line": ((any)(this.Hx_Field_line)), "column": ((any)(this.Hx_Field_column)), "offset": ((any)(this.Hx_Field_offset))})
	if this.Hx_Field_length != 0 {
		Hx_Field_reflect_setField(json, "length", this.Hx_Field_length)
	}

	return json
}

func (this *Hx_Obj_loreline_position) Hx_Field_withOffset(content string, additionalOffset int32, newLength__isOptional any, contentStart__isOptional any) *Hx_Obj_loreline_position {
	newLength := HxDefault[int32]()
	_ = newLength

	if newLength__isOptional != nil {
		newLength = HxConvert[int32](newLength__isOptional)
	} else {
		newLength = 0
	}
	contentStart := HxDefault[int32]()
	_ = contentStart

	if contentStart__isOptional != nil {
		contentStart = HxConvert[int32](contentStart__isOptional)
	} else {
		contentStart = 0
	}
	{
		if additionalOffset == 0 {
			return Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_offset, newLength)
		}

		var currentLine int32
		_ = currentLine
		currentLine = this.Hx_Field_line
		var currentColumn int32
		_ = currentColumn
		currentColumn = this.Hx_Field_column
		var currentOffset int32
		_ = currentOffset
		currentOffset = this.Hx_Field_offset
		if additionalOffset > 0 {
			var chars int32
			_ = chars
			chars = 0
			for chars < additionalOffset {
				var tmp bool
				_ = tmp
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = currentOffset
				var _hx_tmp_1 int32
				_ = _hx_tmp_1
				_hx_tmp_1 = contentStart
				if _hx_tmp_0 < (_hx_tmp_1 + int32(utf8.RuneCountInString(content))) {
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(content, (currentOffset - contentStart)))
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					if (c).Valid == false {
						_hx_tmp_2 = -1
					} else {
						_hx_tmp_2 = (c).Value
					}

					tmp = ((_hx_tmp_2) == 10)
				} else {
					tmp = false
				}

				if tmp {
					currentLine++
					currentColumn = 1
				} else {
					currentColumn++
				}

				chars++
				currentOffset++
			}
		} else {
			var chars int32
			_ = chars
			chars = 0
			for chars > additionalOffset {
				currentOffset--
				var tmp bool
				_ = tmp
				if currentOffset >= 0 {
					var c HxNullable[int32]
					_ = c
					c = (Hx_Field_go_haxe_hxstring_charCodeAt(content, (currentOffset - contentStart)))
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					if (c).Valid == false {
						_hx_tmp_3 = -1
					} else {
						_hx_tmp_3 = (c).Value
					}

					tmp = ((_hx_tmp_3) == 10)
				} else {
					tmp = false
				}

				if tmp {
					currentLine--
					var col int32
					_ = col
					col = 1
					var scanPos int32
					_ = scanPos
					scanPos = (currentOffset - 1)
					for scanPos >= 0 {
						var c HxNullable[int32]
						_ = c
						c = (Hx_Field_go_haxe_hxstring_charCodeAt(content, (scanPos - contentStart)))
						var _hx_tmp_4 int32
						_ = _hx_tmp_4
						if (c).Valid == false {
							_hx_tmp_4 = -1
						} else {
							_hx_tmp_4 = (c).Value
						}

						var c1 int32
						_ = c1
						c1 = _hx_tmp_4
						if c1 == 10 {
							break
						}

						col++
						scanPos--
					}

					currentColumn = col
				} else {
					currentColumn--
				}

				chars--
			}
		}

		if currentOffset < 0 {
			currentOffset = 0
			currentLine = 1
			currentColumn = 1
		}

		return Hx_Obj_loreline_position_CreateInstance(currentLine, currentColumn, currentOffset, newLength)
	}
}

func (this *Hx_Obj_loreline_position) Hx_Field_extendedTo(endPos *Hx_Obj_loreline_position) *Hx_Obj_loreline_position {
	return Hx_Obj_loreline_position_CreateInstance(this.Hx_Field_line, this.Hx_Field_column, this.Hx_Field_offset, ((endPos.Hx_Field_offset + endPos.Hx_Field_length) - this.Hx_Field_offset))
}

func Hx_Field_loreline_position_fromContentAndIndex(content string, offset int32, length__isOptional any) *Hx_Obj_loreline_position {
	length := HxDefault[int32]()
	_ = length

	if length__isOptional != nil {
		length = HxConvert[int32](length__isOptional)
	} else {
		length = 0
	}
	{
		var line int32
		_ = line
		line = 1
		var column int32
		_ = column
		column = 1
		var currentOffset int32
		_ = currentOffset
		currentOffset = 0
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = int32(utf8.RuneCountInString(content))
		if offset > _hx_reserved_len {
			offset = _hx_reserved_len
		}

		if offset < 0 {
			offset = 0
		}

		for currentOffset < offset {
			var c HxNullable[int32]
			_ = c
			c = (Hx_Field_go_haxe_hxstring_charCodeAt(content, currentOffset))
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			if (c).Valid == false {
				_hx_tmp_0 = -1
			} else {
				_hx_tmp_0 = (c).Value
			}

			if (_hx_tmp_0) == 10 {
				line++
				column = 1
			} else {
				column++
			}

			currentOffset++
		}

		return Hx_Obj_loreline_position_CreateInstance(line, column, offset, length)
	}
}
