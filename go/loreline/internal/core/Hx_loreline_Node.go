package core

var Hx_Abstract_loreline_nodeidstep_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NodeIdStep",
	HxMakeArray[string]("SECTION", "BRANCH", "BLOCK", "NODE"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_nodeidstep_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_nodeidstep_impl_() {
	Hx_Field_loreline__node_nodeidstep_impl__SECTION = Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__SECTION()
	Hx_Field_loreline__node_nodeidstep_impl__BRANCH = Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__BRANCH()
	Hx_Field_loreline__node_nodeidstep_impl__BLOCK = Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__BLOCK()
	Hx_Field_loreline__node_nodeidstep_impl__NODE = Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__NODE()
}

type Hx_Obj_VTable_loreline__node_nodeidstep_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_nodeidstep_impl_ struct {
	VTable Hx_Obj_VTable_loreline__node_nodeidstep_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__node_nodeidstep_impl__CreateEmptyInstance() *Hx_Obj_loreline__node_nodeidstep_impl_ {
	obj := &Hx_Obj_loreline__node_nodeidstep_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__node_nodeidstep_impl__CreateInstance() *Hx_Obj_loreline__node_nodeidstep_impl_ {
	obj := Hx_Obj_loreline__node_nodeidstep_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__node_nodeidstep_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.NodeIdStep_Impl_",
	HxMakeArray[string]("SECTION", "BRANCH", "BLOCK", "NODE"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_nodeidstep_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__node_nodeidstep_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_nodeidstep_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "SECTION":
			return &Hx_Field_loreline__node_nodeidstep_impl__SECTION
		case "BRANCH":
			return &Hx_Field_loreline__node_nodeidstep_impl__BRANCH
		case "BLOCK":
			return &Hx_Field_loreline__node_nodeidstep_impl__BLOCK
		case "NODE":
			return &Hx_Field_loreline__node_nodeidstep_impl__NODE
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__node_nodeidstep_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_nodeidstep_impl__RTTI
}

func (this *Hx_Obj_loreline__node_nodeidstep_impl_) Hx_As_Hx_Obj_loreline__node_nodeidstep_impl_() *Hx_Obj_loreline__node_nodeidstep_impl_ {
	return this
}

var Hx_Field_loreline__node_nodeidstep_impl__SECTION int32

func Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__SECTION() int32 {
	return ((int32)(1))
}

var Hx_Field_loreline__node_nodeidstep_impl__BRANCH int32

func Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__BRANCH() int32 {
	return ((int32)(2))
}

var Hx_Field_loreline__node_nodeidstep_impl__BLOCK int32

func Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__BLOCK() int32 {
	return ((int32)(3))
}

var Hx_Field_loreline__node_nodeidstep_impl__NODE int32

func Hx_Init_Hx_Field_loreline__node_nodeidstep_impl__NODE() int32 {
	return ((int32)(4))
}

var Hx_Abstract_loreline_nodeid_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NodeId",
	HxMakeArray[string]("OFFSET", "MAX", "UNDEFINED", "_hx_new", "fromInt64", "fromString", "fromTwoInt32", "packInt32", "int32GetLow", "int32SetLow", "int32GetHigh", "int32SetHigh", "section", "get_section", "set_section", "branch", "get_branch", "set_branch", "block", "get_block", "set_block", "node", "get_node", "set_node", "nextSection", "nextBranch", "nextBlock", "nextNode", "toInt64", "toString", "equals"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_nodeid_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_nodeid_impl_() {
	Hx_Field_loreline__node_nodeid_impl__OFFSET = Hx_Init_Hx_Field_loreline__node_nodeid_impl__OFFSET()
	Hx_Field_loreline__node_nodeid_impl__MAX = Hx_Init_Hx_Field_loreline__node_nodeid_impl__MAX()
	Hx_Field_loreline__node_nodeid_impl__UNDEFINED = Hx_Init_Hx_Field_loreline__node_nodeid_impl__UNDEFINED()
}

type Hx_Obj_VTable_loreline__node_nodeid_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_nodeid_impl_ struct {
	VTable Hx_Obj_VTable_loreline__node_nodeid_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__node_nodeid_impl__CreateEmptyInstance() *Hx_Obj_loreline__node_nodeid_impl_ {
	obj := &Hx_Obj_loreline__node_nodeid_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__node_nodeid_impl__CreateInstance() *Hx_Obj_loreline__node_nodeid_impl_ {
	obj := Hx_Obj_loreline__node_nodeid_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__node_nodeid_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.NodeId_Impl_",
	HxMakeArray[string]("OFFSET", "MAX", "UNDEFINED", "_hx_new", "fromInt64", "fromString", "fromTwoInt32", "packInt32", "int32GetLow", "int32SetLow", "int32GetHigh", "int32SetHigh", "get_section", "set_section", "get_branch", "set_branch", "get_block", "set_block", "get_node", "set_node", "nextSection", "nextBranch", "nextBlock", "nextNode", "toInt64", "toString", "equals"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_nodeid_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__node_nodeid_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_nodeid_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "OFFSET":
			return &Hx_Field_loreline__node_nodeid_impl__OFFSET
		case "MAX":
			return &Hx_Field_loreline__node_nodeid_impl__MAX
		case "UNDEFINED":
			return &Hx_Field_loreline__node_nodeid_impl__UNDEFINED
		case "_hx_new":
			return Hx_Field_loreline__node_nodeid_impl___hx_new
		case "fromInt64":
			return Hx_Field_loreline__node_nodeid_impl__fromInt64
		case "fromString":
			return Hx_Field_loreline__node_nodeid_impl__fromString
		case "fromTwoInt32":
			return Hx_Field_loreline__node_nodeid_impl__fromTwoInt32
		case "packInt32":
			return Hx_Field_loreline__node_nodeid_impl__packInt32
		case "int32GetLow":
			return Hx_Field_loreline__node_nodeid_impl__int32GetLow
		case "int32SetLow":
			return Hx_Field_loreline__node_nodeid_impl__int32SetLow
		case "int32GetHigh":
			return Hx_Field_loreline__node_nodeid_impl__int32GetHigh
		case "int32SetHigh":
			return Hx_Field_loreline__node_nodeid_impl__int32SetHigh
		case "get_section":
			return Hx_Field_loreline__node_nodeid_impl__get_section
		case "set_section":
			return Hx_Field_loreline__node_nodeid_impl__set_section
		case "get_branch":
			return Hx_Field_loreline__node_nodeid_impl__get_branch
		case "set_branch":
			return Hx_Field_loreline__node_nodeid_impl__set_branch
		case "get_block":
			return Hx_Field_loreline__node_nodeid_impl__get_block
		case "set_block":
			return Hx_Field_loreline__node_nodeid_impl__set_block
		case "get_node":
			return Hx_Field_loreline__node_nodeid_impl__get_node
		case "set_node":
			return Hx_Field_loreline__node_nodeid_impl__set_node
		case "nextSection":
			return Hx_Field_loreline__node_nodeid_impl__nextSection
		case "nextBranch":
			return Hx_Field_loreline__node_nodeid_impl__nextBranch
		case "nextBlock":
			return Hx_Field_loreline__node_nodeid_impl__nextBlock
		case "nextNode":
			return Hx_Field_loreline__node_nodeid_impl__nextNode
		case "toInt64":
			return Hx_Field_loreline__node_nodeid_impl__toInt64
		case "toString":
			return Hx_Field_loreline__node_nodeid_impl__toString
		case "equals":
			return Hx_Field_loreline__node_nodeid_impl__equals
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__node_nodeid_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_nodeid_impl__RTTI
}

func (this *Hx_Obj_loreline__node_nodeid_impl_) Hx_As_Hx_Obj_loreline__node_nodeid_impl_() *Hx_Obj_loreline__node_nodeid_impl_ {
	return this
}

var Hx_Field_loreline__node_nodeid_impl__OFFSET int32

func Hx_Init_Hx_Field_loreline__node_nodeid_impl__OFFSET() int32 {
	return 32768
}

var Hx_Field_loreline__node_nodeid_impl__MAX int32

func Hx_Init_Hx_Field_loreline__node_nodeid_impl__MAX() int32 {
	return 65535
}

var Hx_Field_loreline__node_nodeid_impl__UNDEFINED Hx_Typedef_haxe__int64___int64

func Hx_Init_Hx_Field_loreline__node_nodeid_impl__UNDEFINED() Hx_Typedef_haxe__int64___int64 {
	return Hx_Field_loreline__node_nodeid_impl___hx_new(0, 0, 0, 0)
}

func Hx_Field_loreline__node_nodeid_impl___hx_new(section int32, branch int32, block int32, node int32) Hx_Typedef_haxe__int64___int64 {
	if (section < 0) || (section > 65535) {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Section value (", Hx_Field_go_haxe_hxdynamic_toString(section)), ") should be between 0 and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_0, Hx_Field_go_haxe_hxdynamic_toString(65535))))))
	}

	if (branch < 0) || (branch > 65535) {
		var _hx_tmp_1 string
		_ = _hx_tmp_1
		_hx_tmp_1 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Branch value (", Hx_Field_go_haxe_hxdynamic_toString(branch)), ") should be between 0 and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_1, Hx_Field_go_haxe_hxdynamic_toString(65535))))))
	}

	if (block < 0) || (block > 65535) {
		var _hx_tmp_2 string
		_ = _hx_tmp_2
		_hx_tmp_2 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Block value (", Hx_Field_go_haxe_hxdynamic_toString(block)), ") should be between 0 and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_2, Hx_Field_go_haxe_hxdynamic_toString(65535))))))
	}

	if (node < 0) || (node > 65535) {
		var _hx_tmp_3 string
		_ = _hx_tmp_3
		_hx_tmp_3 = Hx_Field_go_haxe_hxstring_concat(Hx_Field_go_haxe_hxstring_concat("Node value (", Hx_Field_go_haxe_hxdynamic_toString(node)), ") should be between 0 and ")
		panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat(_hx_tmp_3, Hx_Field_go_haxe_hxdynamic_toString(65535))))))
	}

	var high int32
	_ = high
	high = (section - 32768)
	var low int32
	_ = low
	low = (branch - 32768)
	if high < 0 {
		high = ((high + 65536) & 65535)
	}

	if low < 0 {
		low = ((low + 65536) & 65535)
	}

	var high1 int32
	_ = high1
	high1 = (block - 32768)
	var low1 int32
	_ = low1
	low1 = (node - 32768)
	if high1 < 0 {
		high1 = ((high1 + 65536) & 65535)
	}

	if low1 < 0 {
		low1 = ((low1 + 65536) & 65535)
	}

	var h int64
	_ = h
	h = int64(((int32)(((high << 16) | (low & 65535)))))
	var l int64
	_ = l
	l = int64(((int32)(((high1 << 16) | (low1 & 65535)))))
	var _hx_tmp_4 int64
	_ = _hx_tmp_4
	_hx_tmp_4 = h
	return ((Hx_Typedef_haxe__int64___int64)(((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_4 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))))
}

func Hx_Field_loreline__node_nodeid_impl__fromInt64(value Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return value
}

func Hx_Field_loreline__node_nodeid_impl__fromString(str string) Hx_Typedef_haxe__int64___int64 {
	var parts HxArray[string]
	_ = parts
	parts = (Hx_Field_go_haxe_hxstring_split(str, "."))
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = (Hx_Field_std_parseInt(parts.Get(0))).Value
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = (Hx_Field_std_parseInt(parts.Get(1))).Value
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = (Hx_Field_std_parseInt(parts.Get(2))).Value
	return Hx_Field_loreline__node_nodeid_impl___hx_new(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, (Hx_Field_std_parseInt(parts.Get(3))).Value)
}

func Hx_Field_loreline__node_nodeid_impl__fromTwoInt32(high int32, low int32) Hx_Typedef_haxe__int64___int64 {
	var h int64
	_ = h
	h = int64(high)
	var l int64
	_ = l
	l = int64(low)
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = h
	return Hx_Field_loreline__node_nodeid_impl__fromInt64(((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_0 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK)))))
}

func Hx_Field_loreline__node_nodeid_impl__packInt32(high int32, low int32) int32 {
	if high < 0 {
		high = ((high + 65536) & 65535)
	}

	if low < 0 {
		low = ((low + 65536) & 65535)
	}

	return ((high << 16) | (low & 65535))
}

func Hx_Field_loreline__node_nodeid_impl__int32GetLow(value int32) int32 {
	var low int32
	_ = low
	low = (value & 65535)
	if low >= 32768 {
		return (low - 65536)
	} else {
		return low
	}
}

func Hx_Field_loreline__node_nodeid_impl__int32SetLow(target int32, value int32) int32 {
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	return ((target & -65536) | value)
}

func Hx_Field_loreline__node_nodeid_impl__int32GetHigh(value int32) int32 {
	var high int32
	_ = high
	high = ((value >> 16) & 65535)
	if high >= 32768 {
		return (high - 65536)
	} else {
		return high
	}
}

func Hx_Field_loreline__node_nodeid_impl__int32SetHigh(target int32, value int32) int32 {
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	return ((target & 65535) | (value << 16))
}

func Hx_Field_loreline__node_nodeid_impl__get_section(this Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)((_hx_tmp_2 >> int64(((int32)(32))))))
	var high int32
	_ = high
	high = ((_hx_tmp_1 >> 16) & 65535)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	if high >= 32768 {
		_hx_tmp_3 = (high - 65536)
	} else {
		_hx_tmp_3 = high
	}

	return ((_hx_tmp_3) + 32768)
}

func Hx_Field_loreline__node_nodeid_impl__set_section(this Hx_Typedef_haxe__int64___int64, section int32) int32 {
	var value int32
	_ = value
	value = (section - 32768)
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var h int64
	_ = h
	h = int64(((int32)(((((int32)(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))))) & 65535) | (value << 16)))))
	var l int64
	_ = l
	l = int64(((int32)(this)))
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = h
	this = ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_1 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
	return section
}

func Hx_Field_loreline__node_nodeid_impl__get_branch(this Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var low int32
	_ = low
	low = (((int32)((_hx_tmp_0 >> int64(((int32)(32)))))) & 65535)
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if low >= 32768 {
		_hx_tmp_1 = (low - 65536)
	} else {
		_hx_tmp_1 = low
	}

	return ((_hx_tmp_1) + 32768)
}

func Hx_Field_loreline__node_nodeid_impl__set_branch(this Hx_Typedef_haxe__int64___int64, branch int32) int32 {
	var value int32
	_ = value
	value = (branch - 32768)
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var h int64
	_ = h
	h = int64(((int32)(((((int32)(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))))) & -65536) | value))))
	var l int64
	_ = l
	l = int64(((int32)(this)))
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = h
	this = ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_1 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
	return branch
}

func Hx_Field_loreline__node_nodeid_impl__get_block(this Hx_Typedef_haxe__int64___int64) int32 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(this))
	var high int32
	_ = high
	high = ((_hx_tmp_0 >> 16) & 65535)
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if high >= 32768 {
		_hx_tmp_1 = (high - 65536)
	} else {
		_hx_tmp_1 = high
	}

	return ((_hx_tmp_1) + 32768)
}

func Hx_Field_loreline__node_nodeid_impl__set_block(this Hx_Typedef_haxe__int64___int64, block int32) int32 {
	var value int32
	_ = value
	value = (block - 32768)
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var h int64
	_ = h
	h = int64(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))))
	var l int64
	_ = l
	l = int64(((int32)(((((int32)(this)) & 65535) | (value << 16)))))
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = h
	this = ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_1 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
	return block
}

func Hx_Field_loreline__node_nodeid_impl__get_node(this Hx_Typedef_haxe__int64___int64) int32 {
	var low int32
	_ = low
	low = (((int32)(this)) & 65535)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if low >= 32768 {
		_hx_tmp_0 = (low - 65536)
	} else {
		_hx_tmp_0 = low
	}

	return ((_hx_tmp_0) + 32768)
}

func Hx_Field_loreline__node_nodeid_impl__set_node(this Hx_Typedef_haxe__int64___int64, node int32) int32 {
	var value int32
	_ = value
	value = (node - 32768)
	if value < 0 {
		value = ((value + 65536) & 65535)
	}

	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(this))
	var h int64
	_ = h
	h = int64(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))))
	var l int64
	_ = l
	l = int64(((int32)(((((int32)(this)) & -65536) | value))))
	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = h
	this = ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_1 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
	return node
}

func Hx_Field_loreline__node_nodeid_impl__nextSection(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = _hx_tmp_0
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int32)((_hx_tmp_2 >> int64(((int32)(32))))))
	var high int32
	_ = high
	high = ((_hx_tmp_1 >> 16) & 65535)
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	if high >= 32768 {
		_hx_tmp_3 = (high - 65536)
	} else {
		_hx_tmp_3 = high
	}

	var section int32
	_ = section
	section = ((_hx_tmp_3) + 32768)
	if section == 65535 {
		panic(Hx_Field_haxe_exception_thrown(((any)("Node id section overflow"))))
	}

	return Hx_Field_loreline__node_nodeid_impl___hx_new((section + 1), 0, 0, 0)
}

func Hx_Field_loreline__node_nodeid_impl__nextBranch(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var low int32
	_ = low
	low = (((int32)((_hx_tmp_0 >> int64(((int32)(32)))))) & 65535)
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if low >= 32768 {
		_hx_tmp_1 = (low - 65536)
	} else {
		_hx_tmp_1 = low
	}

	var branch int32
	_ = branch
	branch = ((_hx_tmp_1) + 32768)
	if branch == 65535 {
		return Hx_Field_loreline__node_nodeid_impl__nextSection(this)
	}

	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var _hx_tmp_4 int64
	_ = _hx_tmp_4
	_hx_tmp_4 = _hx_tmp_2
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = ((int32)((_hx_tmp_4 >> int64(((int32)(32))))))
	var high int32
	_ = high
	high = ((_hx_tmp_3 >> 16) & 65535)
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	if high >= 32768 {
		_hx_tmp_5 = (high - 65536)
	} else {
		_hx_tmp_5 = high
	}

	return Hx_Field_loreline__node_nodeid_impl___hx_new(((_hx_tmp_5) + 32768), (branch + 1), 0, 0)
}

func Hx_Field_loreline__node_nodeid_impl__nextBlock(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int32)(((Hx_Typedef_haxe__int64___int64)(this))))
	var high int32
	_ = high
	high = ((_hx_tmp_0 >> 16) & 65535)
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	if high >= 32768 {
		_hx_tmp_1 = (high - 65536)
	} else {
		_hx_tmp_1 = high
	}

	var block int32
	_ = block
	block = ((_hx_tmp_1) + 32768)
	if block == 65535 {
		return Hx_Field_loreline__node_nodeid_impl__nextBranch(this)
	}

	var _hx_tmp_2 int64
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var _hx_tmp_4 int64
	_ = _hx_tmp_4
	_hx_tmp_4 = _hx_tmp_2
	var _hx_tmp_3 int32
	_ = _hx_tmp_3
	_hx_tmp_3 = ((int32)((_hx_tmp_4 >> int64(((int32)(32))))))
	var high1 int32
	_ = high1
	high1 = ((_hx_tmp_3 >> 16) & 65535)
	var _hx_tmp_5 int64
	_ = _hx_tmp_5
	_hx_tmp_5 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var low int32
	_ = low
	low = (((int32)((_hx_tmp_5 >> int64(((int32)(32)))))) & 65535)
	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	if high1 >= 32768 {
		_hx_tmp_6 = (high1 - 65536)
	} else {
		_hx_tmp_6 = high1
	}

	var _hx_tmp_7 int32
	_ = _hx_tmp_7
	if low >= 32768 {
		_hx_tmp_7 = (low - 65536)
	} else {
		_hx_tmp_7 = low
	}

	return Hx_Field_loreline__node_nodeid_impl___hx_new(((_hx_tmp_6) + 32768), ((_hx_tmp_7) + 32768), (block + 1), 0)
}

func Hx_Field_loreline__node_nodeid_impl__nextNode(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	var low int32
	_ = low
	low = (((int32)(((Hx_Typedef_haxe__int64___int64)(this)))) & 65535)
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if low >= 32768 {
		_hx_tmp_0 = (low - 65536)
	} else {
		_hx_tmp_0 = low
	}

	var node int32
	_ = node
	node = ((_hx_tmp_0) + 32768)
	if node == 65535 {
		return Hx_Field_loreline__node_nodeid_impl__nextBlock(this)
	}

	var _hx_tmp_1 int64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var _hx_tmp_3 int64
	_ = _hx_tmp_3
	_hx_tmp_3 = _hx_tmp_1
	var _hx_tmp_2 int32
	_ = _hx_tmp_2
	_hx_tmp_2 = ((int32)((_hx_tmp_3 >> int64(((int32)(32))))))
	var high int32
	_ = high
	high = ((_hx_tmp_2 >> 16) & 65535)
	var _hx_tmp_4 int64
	_ = _hx_tmp_4
	_hx_tmp_4 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var low1 int32
	_ = low1
	low1 = (((int32)((_hx_tmp_4 >> int64(((int32)(32)))))) & 65535)
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	_hx_tmp_5 = ((int32)(((Hx_Typedef_haxe__int64___int64)(this))))
	var high1 int32
	_ = high1
	high1 = ((_hx_tmp_5 >> 16) & 65535)
	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	if high >= 32768 {
		_hx_tmp_6 = (high - 65536)
	} else {
		_hx_tmp_6 = high
	}

	var _hx_tmp_7 int32
	_ = _hx_tmp_7
	if low1 >= 32768 {
		_hx_tmp_7 = (low1 - 65536)
	} else {
		_hx_tmp_7 = low1
	}

	var _hx_tmp_8 int32
	_ = _hx_tmp_8
	if high1 >= 32768 {
		_hx_tmp_8 = (high1 - 65536)
	} else {
		_hx_tmp_8 = high1
	}

	return Hx_Field_loreline__node_nodeid_impl___hx_new(((_hx_tmp_6) + 32768), ((_hx_tmp_7) + 32768), ((_hx_tmp_8) + 32768), (node + 1))
}

func Hx_Field_loreline__node_nodeid_impl__toInt64(this Hx_Typedef_haxe__int64___int64) Hx_Typedef_haxe__int64___int64 {
	return this
}

func Hx_Field_loreline__node_nodeid_impl__toString(this Hx_Typedef_haxe__int64___int64) string {
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
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
	_hx_tmp_3 = ((int64)(((Hx_Typedef_haxe__int64___int64)(this))))
	var low int32
	_ = low
	low = (((int32)((_hx_tmp_3 >> int64(((int32)(32)))))) & 65535)
	var _hx_tmp_4 int32
	_ = _hx_tmp_4
	_hx_tmp_4 = ((int32)(((Hx_Typedef_haxe__int64___int64)(this))))
	var high1 int32
	_ = high1
	high1 = ((_hx_tmp_4 >> 16) & 65535)
	var low1 int32
	_ = low1
	low1 = (((int32)(((Hx_Typedef_haxe__int64___int64)(this)))) & 65535)
	var _hx_tmp_5 int32
	_ = _hx_tmp_5
	if high >= 32768 {
		_hx_tmp_5 = (high - 65536)
	} else {
		_hx_tmp_5 = high
	}

	var _hx_tmp_6 int32
	_ = _hx_tmp_6
	if low >= 32768 {
		_hx_tmp_6 = (low - 65536)
	} else {
		_hx_tmp_6 = low
	}

	var _hx_tmp_7 int32
	_ = _hx_tmp_7
	if high1 >= 32768 {
		_hx_tmp_7 = (high1 - 65536)
	} else {
		_hx_tmp_7 = high1
	}

	var _hx_tmp_8 int32
	_ = _hx_tmp_8
	if low1 >= 32768 {
		_hx_tmp_8 = (low1 - 65536)
	} else {
		_hx_tmp_8 = low1
	}

	return (Hx_Field_go_haxe_hxstring_concat((Hx_Field_go_haxe_hxstring_concat((Hx_Field_go_haxe_hxstring_concat((""+Hx_Field_go_haxe_hxdynamic_toString(((_hx_tmp_5)+32768))), ".")+Hx_Field_go_haxe_hxdynamic_toString(((_hx_tmp_6)+32768))), ".")+Hx_Field_go_haxe_hxdynamic_toString(((_hx_tmp_7)+32768))), ".") + Hx_Field_go_haxe_hxdynamic_toString(((_hx_tmp_8) + 32768)))
}

func Hx_Field_loreline__node_nodeid_impl__equals(a Hx_Typedef_haxe__int64___int64, b Hx_Typedef_haxe__int64___int64) bool {
	return (((Hx_Typedef_haxe__int64___int64)(a)) == ((Hx_Typedef_haxe__int64___int64)(b)))
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nodeidmap() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nodeidmap() {
}

type Hx_Obj_VTable_loreline_nodeidmap interface {
	Hx_Field_keys() *Hx_Obj_loreline__node_nodeidmapkeyiterator
	Hx_Field_keyValueIterator() *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator
	Hx_Field_get(key Hx_Typedef_haxe__int64___int64) any
	Hx_Field_set(key Hx_Typedef_haxe__int64___int64, value any)
	Hx_Field_remove(key Hx_Typedef_haxe__int64___int64)
	Hx_Field_exists(key Hx_Typedef_haxe__int64___int64) bool
	Hx_Field_clear()
	Hx_Field_iterator() *Hx_Obj_loreline__node_nodeidmapiterator
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nodeidmap struct {
	VTable       Hx_Obj_VTable_loreline_nodeidmap `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map *Hx_Obj_loreline_int64map
}

func Hx_Obj_loreline_nodeidmap_CreateEmptyInstance() *Hx_Obj_loreline_nodeidmap {
	obj := &Hx_Obj_loreline_nodeidmap{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	return obj
}

func Hx_Obj_loreline_nodeidmap_CreateInstance() *Hx_Obj_loreline_nodeidmap {
	obj := Hx_Obj_loreline_nodeidmap_CreateEmptyInstance()
	obj.Hx_New()
	return obj
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_New() {
	{
		this.Hx_Field_map = Hx_Obj_loreline_int64map_CreateInstance()
	}

}

var Hx_Obj_loreline_nodeidmap_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NodeIdMap",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "get", "set", "remove", "exists", "clear", "iterator", "keys", "keyValueIterator"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nodeidmap_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_nodeidmap_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nodeidmap_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nodeidmap_RTTI
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_As_Hx_Obj_loreline_nodeidmap() *Hx_Obj_loreline_nodeidmap {
	return this
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_get(key Hx_Typedef_haxe__int64___int64) any {
	var key1 Hx_Typedef_haxe__int64___int64
	_ = key1
	key1 = key
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(key1))
	return this.Hx_Field_map.VTable.Hx_Field__get(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key1)))
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_set(key Hx_Typedef_haxe__int64___int64, value any) {
	var key1 Hx_Typedef_haxe__int64___int64
	_ = key1
	key1 = key
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(key1))
	this.Hx_Field_map.VTable.Hx_Field__set(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key1)), value)
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_remove(key Hx_Typedef_haxe__int64___int64) {
	var key1 Hx_Typedef_haxe__int64___int64
	_ = key1
	key1 = key
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(key1))
	this.Hx_Field_map.VTable.Hx_Field__remove(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key1)))
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_exists(key Hx_Typedef_haxe__int64___int64) bool {
	var key1 Hx_Typedef_haxe__int64___int64
	_ = key1
	key1 = key
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((int64)(key1))
	return this.Hx_Field_map.VTable.Hx_Field__exists(((int32)((_hx_tmp_0 >> int64(((int32)(32)))))), ((int32)(key1)))
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_clear() {
	this.Hx_Field_map.VTable.Hx_Field_clear()
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_iterator() *Hx_Obj_loreline__node_nodeidmapiterator {
	return Hx_Obj_loreline__node_nodeidmapiterator_CreateInstance(this.Hx_Field_map)
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_keys() *Hx_Obj_loreline__node_nodeidmapkeyiterator {
	return Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateInstance(this.Hx_Field_map)
}

func (this *Hx_Obj_loreline_nodeidmap) Hx_Field_keyValueIterator() *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator {
	return Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateInstance(this.Hx_Field_map)
}

type Hx_Typedef_loreline_nodeidmapkeyval = any

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_nodeidmapiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_nodeidmapiterator() {
}

type Hx_Obj_VTable_loreline__node_nodeidmapiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_nodeidmapiterator struct {
	VTable         Hx_Obj_VTable_loreline__node_nodeidmapiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__node_nodeidmapiterator_CreateEmptyInstance() *Hx_Obj_loreline__node_nodeidmapiterator {
	obj := &Hx_Obj_loreline__node_nodeidmapiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__node_nodeidmapiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__node_nodeidmapiterator {
	obj := Hx_Obj_loreline__node_nodeidmapiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__node_nodeidmapiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.NodeIdMapIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_nodeidmapiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__node_nodeidmapiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_nodeidmapiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_nodeidmapiterator_RTTI
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_As_Hx_Obj_loreline__node_nodeidmapiterator() *Hx_Obj_loreline__node_nodeidmapiterator {
	return this
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__node_nodeidmapiterator) Hx_Field_next() any {
	var v any
	_ = v
	v = this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index)
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return v
}

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_nodeidmapkeyiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_nodeidmapkeyiterator() {
}

type Hx_Obj_VTable_loreline__node_nodeidmapkeyiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() Hx_Typedef_haxe__int64___int64
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_nodeidmapkeyiterator struct {
	VTable         Hx_Obj_VTable_loreline__node_nodeidmapkeyiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateEmptyInstance() *Hx_Obj_loreline__node_nodeidmapkeyiterator {
	obj := &Hx_Obj_loreline__node_nodeidmapkeyiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__node_nodeidmapkeyiterator {
	obj := Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__node_nodeidmapkeyiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.NodeIdMapKeyIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__node_nodeidmapkeyiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_nodeidmapkeyiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_nodeidmapkeyiterator_RTTI
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_As_Hx_Obj_loreline__node_nodeidmapkeyiterator() *Hx_Obj_loreline__node_nodeidmapkeyiterator {
	return this
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyiterator) Hx_Field_next() Hx_Typedef_haxe__int64___int64 {
	var k1 int32
	_ = k1
	k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys1).Get(this.Hx_Field_index))))
	var k2 int32
	_ = k2
	k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys2).Get(this.Hx_Field_index))))
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return Hx_Field_loreline__node_nodeid_impl__fromTwoInt32(k1, k2)
}

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_nodeidmapkeyvalueiterator() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_nodeidmapkeyvalueiterator() {
}

type Hx_Obj_VTable_loreline__node_nodeidmapkeyvalueiterator interface {
	Hx_Field_skipNulls()
	Hx_Field_hasNext() bool
	Hx_Field_next() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_nodeidmapkeyvalueiterator struct {
	VTable         Hx_Obj_VTable_loreline__node_nodeidmapkeyvalueiterator `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_map   *Hx_Obj_loreline_int64map
	Hx_Field_index int32
}

func Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateEmptyInstance() *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator {
	obj := &Hx_Obj_loreline__node_nodeidmapkeyvalueiterator{}
	obj.VTable = obj
	obj.Hx_Field_map = HxDefault[*Hx_Obj_loreline_int64map]()
	obj.Hx_Field_index = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateInstance(_hx_reserved_map *Hx_Obj_loreline_int64map) *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator {
	obj := Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateEmptyInstance()
	obj.Hx_New(_hx_reserved_map)
	return obj
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_New(_hx_reserved_map *Hx_Obj_loreline_int64map) {
	{
		this.Hx_Field_map = _hx_reserved_map
		this.Hx_Field_index = 0
		for {
			var _hx_tmp_0 bool
			_ = _hx_tmp_0
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = this.Hx_Field_index
			if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
				_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
			} else {
				_hx_tmp_0 = false
			}

			if !(_hx_tmp_0) {
				break
			}

			this.Hx_Field_index++
		}
	}

}

var Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.NodeIdMapKeyValueIterator",
	HxMakeArray[string](),
	HxMakeArray[string]("map", "index", "skipNulls", "hasNext", "next"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateInstance(HxConvert[*Hx_Obj_loreline_int64map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_nodeidmapkeyvalueiterator_RTTI
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_As_Hx_Obj_loreline__node_nodeidmapkeyvalueiterator() *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator {
	return this
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_Field_skipNulls() {
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_Field_hasNext() bool {
	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_index
	return (_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values))
}

func (this *Hx_Obj_loreline__node_nodeidmapkeyvalueiterator) Hx_Field_next() any {
	var k1 int32
	_ = k1
	k1 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys1).Get(this.Hx_Field_index))))
	var k2 int32
	_ = k2
	k2 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(HxMakeArrayView[any](this.Hx_Field_map.Hx_Field__keys2).Get(this.Hx_Field_index))))
	var v any
	_ = v
	v = this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index)
	this.Hx_Field_index++
	for {
		var _hx_tmp_0 bool
		_ = _hx_tmp_0
		var _hx_tmp_1 int32
		_ = _hx_tmp_1
		_hx_tmp_1 = this.Hx_Field_index
		if _hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(this.Hx_Field_map.Hx_Field__values) {
			_hx_tmp_0 = ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(this.Hx_Field_map.Hx_Field__values.Get(this.Hx_Field_index), nil))))
		} else {
			_hx_tmp_0 = false
		}

		if !(_hx_tmp_0) {
			break
		}

		this.Hx_Field_index++
	}

	return any(map[string]any{"key": ((any)(Hx_Field_loreline__node_nodeid_impl__fromTwoInt32(k1, k2))), "value": v})
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_node() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_node() {
}

type Hx_Obj_VTable_loreline_node interface {
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_type() string
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_node struct {
	VTable       Hx_Obj_VTable_loreline_node `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_id  Hx_Typedef_haxe__int64___int64
	Hx_Field_pos *Hx_Obj_loreline_position
}

func Hx_Obj_loreline_node_CreateEmptyInstance() *Hx_Obj_loreline_node {
	obj := &Hx_Obj_loreline_node{}
	obj.VTable = obj
	obj.Hx_Field_id = HxDefault[Hx_Typedef_haxe__int64___int64]()
	obj.Hx_Field_pos = HxDefault[*Hx_Obj_loreline_position]()
	return obj
}

func Hx_Obj_loreline_node_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position) *Hx_Obj_loreline_node {
	obj := Hx_Obj_loreline_node_CreateEmptyInstance()
	obj.Hx_New(id, pos)
	return obj
}

func (this *Hx_Obj_loreline_node) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position) {
	{
		this.Hx_Field_id = Hx_Field_loreline__node_nodeid_impl__UNDEFINED
	}
	{
		this.Hx_Field_id = id
		this.Hx_Field_pos = pos
	}

}

var Hx_Obj_loreline_node_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Node",
	HxMakeArray[string]("fromJson"),
	HxMakeArray[string]("id", "pos", "type", "toJson", "each"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_node_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline_node_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_node_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "fromJson":
			return Hx_Field_loreline_node_fromJson
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_node) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_node_RTTI
}

func (this *Hx_Obj_loreline_node) Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node {
	return this
}

func (this *Hx_Obj_loreline_node) Hx_Field_type() string {
	return "Node"
}

func (this *Hx_Obj_loreline_node) Hx_Field_toJson() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(Hx_Field_loreline__node_nodeid_impl__toString(this.Hx_Field_id)))
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = ((any)(this.VTable.Hx_Field_type()))
	return any(map[string]any{"id": _hx_tmp_0, "type": _hx_tmp_1, "pos": this.Hx_Field_pos.VTable.Hx_Field_toJson()})
}

func (this *Hx_Obj_loreline_node) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
}

func Hx_Field_loreline_node_fromJson(json any) *Hx_Obj_loreline_node {
	return Hx_Field_loreline_jsontoast_nodeFromJson(json)
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_comment() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_comment() {
}

type Hx_Obj_VTable_loreline_comment interface {
	Hx_Field_type() string
	Hx_Field_toJson() any
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_comment struct {
	Hx_Obj_loreline_node
	VTable             Hx_Obj_VTable_loreline_comment `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_content   string
	Hx_Field_multiline bool
	Hx_Field_isHash    bool
}

func Hx_Obj_loreline_comment_CreateEmptyInstance() *Hx_Obj_loreline_comment {
	obj := &Hx_Obj_loreline_comment{}
	obj.Hx_Obj_loreline_node = *Hx_Obj_loreline_node_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_content = HxDefault[string]()
	obj.Hx_Field_multiline = HxDefault[bool]()
	obj.Hx_Field_isHash = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_comment_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, content string, multiline bool, isHash__isOptional any) *Hx_Obj_loreline_comment {
	obj := Hx_Obj_loreline_comment_CreateEmptyInstance()
	obj.Hx_New(id, pos, content, multiline, isHash__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_comment) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, content string, multiline bool, isHash__isOptional any) {
	{
		isHash := HxDefault[bool]()
		_ = isHash

		if isHash__isOptional != nil {
			isHash = HxConvert[bool](isHash__isOptional)
		} else {
			isHash = false
		}
		{
			this.Hx_Obj_loreline_node.Hx_New(id, pos)
			this.Hx_Field_content = content
			this.Hx_Field_multiline = multiline
			this.Hx_Field_isHash = isHash
		}

	}
}

var Hx_Obj_loreline_comment_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Comment",
	HxMakeArray[string](),
	HxMakeArray[string]("content", "multiline", "isHash", "type", "toJson", "id", "pos", "each"),
	Hx_Obj_loreline_node_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_comment_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[bool](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_comment_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_comment_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_comment) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_comment_RTTI
}

func (this *Hx_Obj_loreline_comment) Hx_As_Hx_Obj_loreline_comment() *Hx_Obj_loreline_comment {
	return this
}

func (this *Hx_Obj_loreline_comment) Hx_Field_type() string {
	return "Comment"
}

func (this *Hx_Obj_loreline_comment) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_node.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "content", this.Hx_Field_content)
	Hx_Field_go_haxe_hxdynamic_setField(json, "multiline", this.Hx_Field_multiline)
	if this.Hx_Field_isHash {
		Hx_Field_go_haxe_hxdynamic_setField(json, "isHash", true)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_astnode() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_astnode() {
}

type Hx_Obj_VTable_loreline_astnode interface {
	Hx_Field_type() string
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_astnode struct {
	Hx_Obj_loreline_node
	VTable                    Hx_Obj_VTable_loreline_astnode `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_leadingComments  HxArray[*Hx_Obj_loreline_comment]
	Hx_Field_trailingComments HxArray[*Hx_Obj_loreline_comment]
}

func Hx_Obj_loreline_astnode_CreateEmptyInstance() *Hx_Obj_loreline_astnode {
	obj := &Hx_Obj_loreline_astnode{}
	obj.Hx_Obj_loreline_node = *Hx_Obj_loreline_node_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_leadingComments = HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
	obj.Hx_Field_trailingComments = HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
	return obj
}

func Hx_Obj_loreline_astnode_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_astnode {
	obj := Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.Hx_New(id, pos, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_astnode) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_node.Hx_New(id, pos)
			this.Hx_Field_leadingComments = leadingComments
			this.Hx_Field_trailingComments = trailingComments
		}

	}
}

var Hx_Obj_loreline_astnode_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.AstNode",
	HxMakeArray[string](),
	HxMakeArray[string]("leadingComments", "trailingComments", "type", "eachComment", "each", "toJson", "id", "pos"),
	Hx_Obj_loreline_node_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_astnode_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_astnode_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_astnode_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_astnode) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_astnode_RTTI
}

func (this *Hx_Obj_loreline_astnode) Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode {
	return this
}

func (this *Hx_Obj_loreline_astnode) Hx_Field_type() string {
	return "AstNode"
}

func (this *Hx_Obj_loreline_astnode) Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment)) {
	if this.Hx_Field_leadingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_leadingComments
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			handleComment(this, comment)
		}
	}

	if this.Hx_Field_trailingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_trailingComments
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			handleComment(this, comment)
		}
	}

	var _gthis *Hx_Obj_loreline_astnode
	_ = _gthis
	_gthis = this
	this.VTable.Hx_Field_each(func(node *Hx_Obj_loreline_node, _hx_blank1 *Hx_Obj_loreline_node) {
		if Hx_Field_std_isOfType(node, Hx_Obj_loreline_astnode_RTTI) {
			var astNode *Hx_Obj_loreline_astnode
			_ = astNode
			astNode = node.VTable.(interface {
				Hx_As_Hx_Obj_loreline_astnode() *Hx_Obj_loreline_astnode
			}).Hx_As_Hx_Obj_loreline_astnode()
			if astNode.Hx_Field_leadingComments != nil {
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = astNode.Hx_Field_leadingComments
				for {
					var _hx_tmp_2 int32
					_ = _hx_tmp_2
					_hx_tmp_2 = _g
					if !(_hx_tmp_2 < _g1.Len()) {
						break
					}

					var comment *Hx_Obj_loreline_comment
					_ = comment
					comment = _g1.Get(_g)
					_g++
					handleComment(_gthis, comment)
				}
			}

			if astNode.Hx_Field_trailingComments != nil {
				var _g int32
				_ = _g
				_g = 0
				var _g1 HxArray[*Hx_Obj_loreline_comment]
				_ = _g1
				_g1 = astNode.Hx_Field_trailingComments
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = _g
					if !(_hx_tmp_3 < _g1.Len()) {
						break
					}

					var comment *Hx_Obj_loreline_comment
					_ = comment
					comment = _g1.Get(_g)
					_g++
					handleComment(_gthis, comment)
				}
			}
		}
	})
}

func (this *Hx_Obj_loreline_astnode) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_node.Hx_Field_each(handleNode)
	if this.Hx_Field_leadingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_leadingComments
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := comment
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
			comment.VTable.Hx_Field_each(handleNode)
		}
	}

	if this.Hx_Field_trailingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_trailingComments
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			var _hx_tmp_3 *Hx_Obj_loreline_node
			_ = _hx_tmp_3
			_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
				_hx_up := comment
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_3, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			comment.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_astnode) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_node.Hx_Field_toJson()
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	var _hx_tmp_1 bool
	_ = _hx_tmp_1
	if this.Hx_Field_leadingComments != nil {
		_hx_tmp_1 = (this.Hx_Field_leadingComments.Len() > 0)
	} else {
		_hx_tmp_1 = false
	}

	if _hx_tmp_1 {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_2 bool
		_ = _hx_tmp_2
		if this.Hx_Field_trailingComments != nil {
			_hx_tmp_2 = (this.Hx_Field_trailingComments.Len() > 0)
		} else {
			_hx_tmp_2 = false
		}

		_hx_tmp_0 = _hx_tmp_2
	}

	if _hx_tmp_0 {
		var comments any
		_ = comments
		comments = any(map[string]any{})
		var _hx_tmp_3 bool
		_ = _hx_tmp_3
		if this.Hx_Field_leadingComments != nil {
			_hx_tmp_3 = (this.Hx_Field_leadingComments.Len() > 0)
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
				var _g2 HxArray[*Hx_Obj_loreline_comment]
				_ = _g2
				_g2 = this.Hx_Field_leadingComments
				for {
					var _hx_tmp_4 int32
					_ = _hx_tmp_4
					_hx_tmp_4 = _g1
					if !(_hx_tmp_4 < _g2.Len()) {
						break
					}

					var c *Hx_Obj_loreline_comment
					_ = c
					c = _g2.Get(_g1)
					_g1++
					var _hx_tmp_5 HxArray[any]
					_ = _hx_tmp_5
					_hx_tmp_5 = _g
					_ = (Hx_Array_Push[any](_hx_tmp_5, c.VTable.Hx_Field_toJson()))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(comments, "leading", _g)
		}

		var _hx_tmp_6 bool
		_ = _hx_tmp_6
		if this.Hx_Field_trailingComments != nil {
			_hx_tmp_6 = (this.Hx_Field_trailingComments.Len() > 0)
		} else {
			_hx_tmp_6 = false
		}

		if _hx_tmp_6 {
			var _g HxArray[any]
			_ = _g
			_g = HxMakeArray[any]().Dyn()
			{
				var _g1 int32
				_ = _g1
				_g1 = 0
				var _g2 HxArray[*Hx_Obj_loreline_comment]
				_ = _g2
				_g2 = this.Hx_Field_trailingComments
				for {
					var _hx_tmp_7 int32
					_ = _hx_tmp_7
					_hx_tmp_7 = _g1
					if !(_hx_tmp_7 < _g2.Len()) {
						break
					}

					var c *Hx_Obj_loreline_comment
					_ = c
					c = _g2.Get(_g1)
					_g1++
					var _hx_tmp_8 HxArray[any]
					_ = _hx_tmp_8
					_hx_tmp_8 = _g
					_ = (Hx_Array_Push[any](_hx_tmp_8, c.VTable.Hx_Field_toJson()))
				}
			}

			Hx_Field_go_haxe_hxdynamic_setField(comments, "trailing", _g)
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "comments", comments)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nexpr() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nexpr() {
}

type Hx_Obj_VTable_loreline_nexpr interface {
	Hx_Field_toJson() any
	Hx_Field_type() string
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nexpr struct {
	Hx_Obj_loreline_astnode
	VTable          Hx_Obj_VTable_loreline_nexpr `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_parens int32
}

func Hx_Obj_loreline_nexpr_CreateEmptyInstance() *Hx_Obj_loreline_nexpr {
	obj := &Hx_Obj_loreline_nexpr{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_parens = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nexpr_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nexpr {
	obj := Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.Hx_New(id, pos, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nexpr) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Field_parens = 0
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
		}

	}
}

var Hx_Obj_loreline_nexpr_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NExpr",
	HxMakeArray[string](),
	HxMakeArray[string]("parens", "toJson", "type", "leadingComments", "trailingComments", "eachComment", "each", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nexpr_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[any](params.Get(3)))
	},
	func() any {
		return Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nexpr_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nexpr) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nexpr_RTTI
}

func (this *Hx_Obj_loreline_nexpr) Hx_As_Hx_Obj_loreline_nexpr() *Hx_Obj_loreline_nexpr {
	return this
}

func (this *Hx_Obj_loreline_nexpr) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if this.Hx_Field_parens > 0 {
		Hx_Field_go_haxe_hxdynamic_setField(json, "parens", this.Hx_Field_parens)
	}

	return json
}

func (this *Hx_Obj_loreline_nexpr) Hx_Field_type() string {
	return "Expr"
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nstatedecl() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nstatedecl() {
}

type Hx_Obj_VTable_loreline_nstatedecl interface {
	Hx_Field_type() string
	Hx_Field_get(name string) *Hx_Obj_loreline_nexpr
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nstatedecl struct {
	Hx_Obj_loreline_astnode
	VTable             Hx_Obj_VTable_loreline_nstatedecl `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_temporary bool
	Hx_Field_fields    HxArray[*Hx_Obj_loreline_nobjectfield]
	Hx_Field_style     int32
}

func Hx_Obj_loreline_nstatedecl_CreateEmptyInstance() *Hx_Obj_loreline_nstatedecl {
	obj := &Hx_Obj_loreline_nstatedecl{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_temporary = HxDefault[bool]()
	obj.Hx_Field_fields = HxDefault[HxArray[*Hx_Obj_loreline_nobjectfield]]()
	obj.Hx_Field_style = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nstatedecl_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, temporary bool, fields HxArray[*Hx_Obj_loreline_nobjectfield], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nstatedecl {
	obj := Hx_Obj_loreline_nstatedecl_CreateEmptyInstance()
	obj.Hx_New(id, pos, temporary, fields, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, temporary bool, fields HxArray[*Hx_Obj_loreline_nobjectfield], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_temporary = temporary
			this.Hx_Field_fields = fields
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nstatedecl_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NStateDecl",
	HxMakeArray[string](),
	HxMakeArray[string]("temporary", "fields", "style", "type", "get", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nstatedecl_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[bool](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_nobjectfield]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nstatedecl_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nstatedecl_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nstatedecl_RTTI
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_As_Hx_Obj_loreline_nstatedecl() *Hx_Obj_loreline_nstatedecl {
	return this
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_Field_type() string {
	return "State"
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_Field_get(name string) *Hx_Obj_loreline_nexpr {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = this.Hx_Field_fields
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
			if field.Hx_Field_name == name {
				return field.Hx_Field_value
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_nstatedecl) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_fields != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_fields.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_nobjectfield
			_ = child
			child = this.Hx_Field_fields.Get(i)
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

func (this *Hx_Obj_loreline_nstatedecl) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "temporary", this.Hx_Field_temporary)
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g2
		_g2 = this.Hx_Field_fields
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var prop *Hx_Obj_loreline_nobjectfield
			_ = prop
			prop = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_1, prop.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "fields", _g)
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nobjectfield() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nobjectfield() {
}

type Hx_Obj_VTable_loreline_nobjectfield interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nobjectfield struct {
	Hx_Obj_loreline_astnode
	VTable         Hx_Obj_VTable_loreline_nobjectfield `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name  string
	Hx_Field_value *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nobjectfield_CreateEmptyInstance() *Hx_Obj_loreline_nobjectfield {
	obj := &Hx_Obj_loreline_nobjectfield{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_value = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nobjectfield_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, value *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nobjectfield {
	obj := Hx_Obj_loreline_nobjectfield_CreateEmptyInstance()
	obj.Hx_New(id, pos, name, value, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, value *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_name = name
			this.Hx_Field_value = value
		}

	}
}

var Hx_Obj_loreline_nobjectfield_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NObjectField",
	HxMakeArray[string](),
	HxMakeArray[string]("name", "value", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nobjectfield_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nobjectfield_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nobjectfield_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nobjectfield_RTTI
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_As_Hx_Obj_loreline_nobjectfield() *Hx_Obj_loreline_nobjectfield {
	return this
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_Field_type() string {
	return "Field"
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_value != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_value
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_value.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nobjectfield) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "name", this.Hx_Field_name)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "value", this.Hx_Field_value.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ncharacterdecl() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ncharacterdecl() {
}

type Hx_Obj_VTable_loreline_ncharacterdecl interface {
	Hx_Field_type() string
	Hx_Field_get(name string) *Hx_Obj_loreline_nexpr
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ncharacterdecl struct {
	Hx_Obj_loreline_astnode
	VTable           Hx_Obj_VTable_loreline_ncharacterdecl `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name    string
	Hx_Field_namePos *Hx_Obj_loreline_position
	Hx_Field_fields  HxArray[*Hx_Obj_loreline_nobjectfield]
	Hx_Field_style   int32
}

func Hx_Obj_loreline_ncharacterdecl_CreateEmptyInstance() *Hx_Obj_loreline_ncharacterdecl {
	obj := &Hx_Obj_loreline_ncharacterdecl{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_namePos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_fields = HxDefault[HxArray[*Hx_Obj_loreline_nobjectfield]]()
	obj.Hx_Field_style = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_ncharacterdecl_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, namePos *Hx_Obj_loreline_position, fields HxArray[*Hx_Obj_loreline_nobjectfield], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ncharacterdecl {
	obj := Hx_Obj_loreline_ncharacterdecl_CreateEmptyInstance()
	obj.Hx_New(id, pos, name, namePos, fields, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, namePos *Hx_Obj_loreline_position, fields HxArray[*Hx_Obj_loreline_nobjectfield], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_name = name
			this.Hx_Field_namePos = namePos
			this.Hx_Field_fields = fields
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_ncharacterdecl_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NCharacterDecl",
	HxMakeArray[string](),
	HxMakeArray[string]("name", "namePos", "fields", "style", "type", "get", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ncharacterdecl_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[*Hx_Obj_loreline_position](params.Get(3)), HxConvert[HxArray[*Hx_Obj_loreline_nobjectfield]](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)))
	},
	func() any {
		return Hx_Obj_loreline_ncharacterdecl_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ncharacterdecl_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ncharacterdecl_RTTI
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_As_Hx_Obj_loreline_ncharacterdecl() *Hx_Obj_loreline_ncharacterdecl {
	return this
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_Field_type() string {
	return "Character"
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_Field_get(name string) *Hx_Obj_loreline_nexpr {
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g1
		_g1 = this.Hx_Field_fields
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
			if field.Hx_Field_name == name {
				return field.Hx_Field_value
			}
		}
	}

	return nil
}

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_fields != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_fields.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_nobjectfield
			_ = child
			child = this.Hx_Field_fields.Get(i)
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

func (this *Hx_Obj_loreline_ncharacterdecl) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "name", this.Hx_Field_name)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "namePos", this.Hx_Field_namePos.VTable.Hx_Field_toJson())
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nobjectfield]
		_ = _g2
		_g2 = this.Hx_Field_fields
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < _g2.Len()) {
				break
			}

			var prop *Hx_Obj_loreline_nobjectfield
			_ = prop
			prop = _g2.Get(_g1)
			_g1++
			var _hx_tmp_2 HxArray[any]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_2, prop.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "fields", _g)
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nbeatparam() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nbeatparam() {
}

type Hx_Obj_VTable_loreline_nbeatparam interface {
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nbeatparam struct {
	VTable                Hx_Obj_VTable_loreline_nbeatparam `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name         string
	Hx_Field_namePos      *Hx_Obj_loreline_position
	Hx_Field_defaultValue *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nbeatparam_CreateEmptyInstance() *Hx_Obj_loreline_nbeatparam {
	obj := &Hx_Obj_loreline_nbeatparam{}
	obj.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_namePos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_defaultValue = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nbeatparam_CreateInstance(name string, namePos *Hx_Obj_loreline_position, defaultValue__isOptional any) *Hx_Obj_loreline_nbeatparam {
	obj := Hx_Obj_loreline_nbeatparam_CreateEmptyInstance()
	obj.Hx_New(name, namePos, defaultValue__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nbeatparam) Hx_New(name string, namePos *Hx_Obj_loreline_position, defaultValue__isOptional any) {
	{
		defaultValue := HxDefault[*Hx_Obj_loreline_nexpr]()
		_ = defaultValue

		if defaultValue__isOptional != nil {
			defaultValue = HxConvert[*Hx_Obj_loreline_nexpr](defaultValue__isOptional)
		} else {
			defaultValue = nil
		}
		{
			this.Hx_Field_name = name
			this.Hx_Field_namePos = namePos
			this.Hx_Field_defaultValue = defaultValue
		}

	}
}

var Hx_Obj_loreline_nbeatparam_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NBeatParam",
	HxMakeArray[string](),
	HxMakeArray[string]("name", "namePos", "defaultValue", "toJson"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nbeatparam_CreateInstance(HxConvert[string](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[any](params.Get(2)))
	},
	func() any {
		return Hx_Obj_loreline_nbeatparam_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nbeatparam_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nbeatparam) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nbeatparam_RTTI
}

func (this *Hx_Obj_loreline_nbeatparam) Hx_As_Hx_Obj_loreline_nbeatparam() *Hx_Obj_loreline_nbeatparam {
	return this
}

func (this *Hx_Obj_loreline_nbeatparam) Hx_Field_toJson() any {
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = ((any)(this.Hx_Field_name))
	var json any
	_ = json
	json = any(map[string]any{"name": _hx_tmp_0, "namePos": this.Hx_Field_namePos.VTable.Hx_Field_toJson()})
	if this.Hx_Field_defaultValue != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "defaultValue", this.Hx_Field_defaultValue.VTable.Hx_Field_toJson())
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nbeatdecl() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nbeatdecl() {
}

type Hx_Obj_VTable_loreline_nbeatdecl interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nbeatdecl struct {
	Hx_Obj_loreline_astnode
	VTable          Hx_Obj_VTable_loreline_nbeatdecl `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name   string
	Hx_Field_body   HxArray[*Hx_Obj_loreline_astnode]
	Hx_Field_style  int32
	Hx_Field_params HxArray[*Hx_Obj_loreline_nbeatparam]
}

func Hx_Obj_loreline_nbeatdecl_CreateEmptyInstance() *Hx_Obj_loreline_nbeatdecl {
	obj := &Hx_Obj_loreline_nbeatdecl{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_name = HxDefault[string]()
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	obj.Hx_Field_style = HxDefault[int32]()
	obj.Hx_Field_params = HxDefault[HxArray[*Hx_Obj_loreline_nbeatparam]]()
	return obj
}

func Hx_Obj_loreline_nbeatdecl_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nbeatdecl {
	obj := Hx_Obj_loreline_nbeatdecl_CreateEmptyInstance()
	obj.Hx_New(id, pos, name, body, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nbeatdecl) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name string, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Field_params = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_name = name
			this.Hx_Field_body = body
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nbeatdecl_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NBeatDecl",
	HxMakeArray[string](),
	HxMakeArray[string]("name", "body", "style", "params", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nbeatdecl_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nbeatdecl_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nbeatdecl_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nbeatdecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nbeatdecl_RTTI
}

func (this *Hx_Obj_loreline_nbeatdecl) Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl {
	return this
}

func (this *Hx_Obj_loreline_nbeatdecl) Hx_Field_type() string {
	return "Beat"
}

func (this *Hx_Obj_loreline_nbeatdecl) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_params != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nbeatparam]
		_ = _g1
		_g1 = this.Hx_Field_params
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			if !(_hx_tmp_0 < _g1.Len()) {
				break
			}

			var param *Hx_Obj_loreline_nbeatparam
			_ = param
			param = _g1.Get(_g)
			_g++
			if param.Hx_Field_defaultValue != nil {
				var _hx_tmp_1 *Hx_Obj_loreline_node
				_ = _hx_tmp_1
				_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
					_hx_up := param.Hx_Field_defaultValue
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
				param.Hx_Field_defaultValue.VTable.Hx_Field_each(handleNode)
			}
		}
	}

	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_body.Len()
		for _g < _g1 {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_2
			var child *Hx_Obj_loreline_astnode
			_ = child
			child = this.Hx_Field_body.Get(i)
			var _hx_tmp_3 *Hx_Obj_loreline_node
			_ = _hx_tmp_3
			_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
				_hx_up := child
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_3, (func() *Hx_Obj_loreline_node {
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

func (this *Hx_Obj_loreline_nbeatdecl) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "name", this.Hx_Field_name)
	if this.Hx_Field_params != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_nbeatparam]
			_ = _g2
			_g2 = this.Hx_Field_params
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g1
				if !(_hx_tmp_0 < _g2.Len()) {
					break
				}

				var param *Hx_Obj_loreline_nbeatparam
				_ = param
				param = _g2.Get(_g1)
				_g1++
				var _hx_tmp_1 HxArray[any]
				_ = _hx_tmp_1
				_hx_tmp_1 = _g
				_ = (Hx_Array_Push[any](_hx_tmp_1, param.VTable.Hx_Field_toJson()))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "params", _g)
	}

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
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g1
			if !(_hx_tmp_2 < _g2.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g2.Get(_g1)
			_g1++
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_3, node.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "body", _g)
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_4, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

type Hx_Enum_loreline_stringparttype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_stringparttype()
}

type Hx_Enum_loreline_stringparttype_Tag struct {
	Hx_Field_closing bool
	Hx_Field_expr    *Hx_Obj_loreline_nstringliteral
}

func (this Hx_Enum_loreline_stringparttype_Tag) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_stringparttype_Tag) M_Hx_Enum_loreline_stringparttype()      {}
func (this Hx_Enum_loreline_stringparttype_Tag) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_loreline_stringparttype_Tag) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_stringparttype_RTTI
}
func (this Hx_Enum_loreline_stringparttype_Tag) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_closing), any(this.Hx_Field_expr))
}
func (this Hx_Enum_loreline_stringparttype_Tag) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_closing)
	case 1:
		return any(this.Hx_Field_expr)
	default:
		return nil
	}
}

type Hx_Enum_loreline_stringparttype_Raw struct {
	Hx_Field_text string
}

func (this Hx_Enum_loreline_stringparttype_Raw) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_stringparttype_Raw) M_Hx_Enum_loreline_stringparttype()      {}
func (this Hx_Enum_loreline_stringparttype_Raw) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_loreline_stringparttype_Raw) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_stringparttype_RTTI
}
func (this Hx_Enum_loreline_stringparttype_Raw) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_text))
}
func (this Hx_Enum_loreline_stringparttype_Raw) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_text)
	default:
		return nil
	}
}

type Hx_Enum_loreline_stringparttype_Expr struct {
	Hx_Field_expr *Hx_Obj_loreline_nexpr
}

func (this Hx_Enum_loreline_stringparttype_Expr) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_stringparttype_Expr) M_Hx_Enum_loreline_stringparttype()      {}
func (this Hx_Enum_loreline_stringparttype_Expr) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_loreline_stringparttype_Expr) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_stringparttype_RTTI
}
func (this Hx_Enum_loreline_stringparttype_Expr) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_expr))
}
func (this Hx_Enum_loreline_stringparttype_Expr) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_expr)
	default:
		return nil
	}
}

var Hx_Enum_loreline_stringparttype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.StringPartType",
	HxMakeArray[string]("Raw", "Expr", "Tag"),
	HxMakeArray[int32](1, 1, 2),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_stringparttype_Raw{((string)(Hx_Field_go_haxe_hxdynamic_toString(params.Get_Dyn(0))))}
		case 1:
			return Hx_Enum_loreline_stringparttype_Expr{params.Get_Dyn(0).(*Hx_Obj_loreline_nexpr)}
		case 2:
			return Hx_Enum_loreline_stringparttype_Tag{((bool)(Hx_Field_go_haxe_hxdynamic_toBool(params.Get_Dyn(0)))), params.Get_Dyn(1).(*Hx_Obj_loreline_nstringliteral)}
		default:
			return nil
		}
	},
	nil,
)

var Hx_Abstract_loreline_blockstyle_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.BlockStyle",
	HxMakeArray[string]("Plain", "Braces", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_blockstyle_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_blockstyle_impl_() {
	Hx_Field_loreline__node_blockstyle_impl__Plain = Hx_Init_Hx_Field_loreline__node_blockstyle_impl__Plain()
	Hx_Field_loreline__node_blockstyle_impl__Braces = Hx_Init_Hx_Field_loreline__node_blockstyle_impl__Braces()
}

type Hx_Obj_VTable_loreline__node_blockstyle_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_blockstyle_impl_ struct {
	VTable Hx_Obj_VTable_loreline__node_blockstyle_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__node_blockstyle_impl__CreateEmptyInstance() *Hx_Obj_loreline__node_blockstyle_impl_ {
	obj := &Hx_Obj_loreline__node_blockstyle_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__node_blockstyle_impl__CreateInstance() *Hx_Obj_loreline__node_blockstyle_impl_ {
	obj := Hx_Obj_loreline__node_blockstyle_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__node_blockstyle_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.BlockStyle_Impl_",
	HxMakeArray[string]("Plain", "Braces", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_blockstyle_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__node_blockstyle_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_blockstyle_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Plain":
			return &Hx_Field_loreline__node_blockstyle_impl__Plain
		case "Braces":
			return &Hx_Field_loreline__node_blockstyle_impl__Braces
		case "toString":
			return Hx_Field_loreline__node_blockstyle_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__node_blockstyle_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_blockstyle_impl__RTTI
}

func (this *Hx_Obj_loreline__node_blockstyle_impl_) Hx_As_Hx_Obj_loreline__node_blockstyle_impl_() *Hx_Obj_loreline__node_blockstyle_impl_ {
	return this
}

var Hx_Field_loreline__node_blockstyle_impl__Plain int32

func Hx_Init_Hx_Field_loreline__node_blockstyle_impl__Plain() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__node_blockstyle_impl__Braces int32

func Hx_Init_Hx_Field_loreline__node_blockstyle_impl__Braces() int32 {
	return ((int32)(1))
}

func Hx_Field_loreline__node_blockstyle_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "Plain"
		}

	case 1:
		{
			return "Braces"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

var Hx_Abstract_loreline_conditionstyle_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.ConditionStyle",
	HxMakeArray[string]("Plain", "Parens", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_conditionstyle_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_conditionstyle_impl_() {
	Hx_Field_loreline__node_conditionstyle_impl__Plain = Hx_Init_Hx_Field_loreline__node_conditionstyle_impl__Plain()
	Hx_Field_loreline__node_conditionstyle_impl__Parens = Hx_Init_Hx_Field_loreline__node_conditionstyle_impl__Parens()
}

type Hx_Obj_VTable_loreline__node_conditionstyle_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_conditionstyle_impl_ struct {
	VTable Hx_Obj_VTable_loreline__node_conditionstyle_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__node_conditionstyle_impl__CreateEmptyInstance() *Hx_Obj_loreline__node_conditionstyle_impl_ {
	obj := &Hx_Obj_loreline__node_conditionstyle_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__node_conditionstyle_impl__CreateInstance() *Hx_Obj_loreline__node_conditionstyle_impl_ {
	obj := Hx_Obj_loreline__node_conditionstyle_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__node_conditionstyle_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.ConditionStyle_Impl_",
	HxMakeArray[string]("Plain", "Parens", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_conditionstyle_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__node_conditionstyle_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_conditionstyle_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Plain":
			return &Hx_Field_loreline__node_conditionstyle_impl__Plain
		case "Parens":
			return &Hx_Field_loreline__node_conditionstyle_impl__Parens
		case "toString":
			return Hx_Field_loreline__node_conditionstyle_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__node_conditionstyle_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_conditionstyle_impl__RTTI
}

func (this *Hx_Obj_loreline__node_conditionstyle_impl_) Hx_As_Hx_Obj_loreline__node_conditionstyle_impl_() *Hx_Obj_loreline__node_conditionstyle_impl_ {
	return this
}

var Hx_Field_loreline__node_conditionstyle_impl__Plain int32

func Hx_Init_Hx_Field_loreline__node_conditionstyle_impl__Plain() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__node_conditionstyle_impl__Parens int32

func Hx_Init_Hx_Field_loreline__node_conditionstyle_impl__Parens() int32 {
	return ((int32)(2))
}

func Hx_Field_loreline__node_conditionstyle_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "Plain"
		}

	case 2:
		{
			return "Parens"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

var Hx_Abstract_loreline_alternativemode_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.AlternativeMode",
	HxMakeArray[string]("Sequence", "Cycle", "Once", "Pick", "Shuffle", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__node_alternativemode_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__node_alternativemode_impl_() {
	Hx_Field_loreline__node_alternativemode_impl__Sequence = Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Sequence()
	Hx_Field_loreline__node_alternativemode_impl__Cycle = Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Cycle()
	Hx_Field_loreline__node_alternativemode_impl__Once = Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Once()
	Hx_Field_loreline__node_alternativemode_impl__Pick = Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Pick()
	Hx_Field_loreline__node_alternativemode_impl__Shuffle = Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Shuffle()
}

type Hx_Obj_VTable_loreline__node_alternativemode_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__node_alternativemode_impl_ struct {
	VTable Hx_Obj_VTable_loreline__node_alternativemode_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__node_alternativemode_impl__CreateEmptyInstance() *Hx_Obj_loreline__node_alternativemode_impl_ {
	obj := &Hx_Obj_loreline__node_alternativemode_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__node_alternativemode_impl__CreateInstance() *Hx_Obj_loreline__node_alternativemode_impl_ {
	obj := Hx_Obj_loreline__node_alternativemode_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__node_alternativemode_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Node.AlternativeMode_Impl_",
	HxMakeArray[string]("Sequence", "Cycle", "Once", "Pick", "Shuffle", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__node_alternativemode_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__node_alternativemode_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__node_alternativemode_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Sequence":
			return &Hx_Field_loreline__node_alternativemode_impl__Sequence
		case "Cycle":
			return &Hx_Field_loreline__node_alternativemode_impl__Cycle
		case "Once":
			return &Hx_Field_loreline__node_alternativemode_impl__Once
		case "Pick":
			return &Hx_Field_loreline__node_alternativemode_impl__Pick
		case "Shuffle":
			return &Hx_Field_loreline__node_alternativemode_impl__Shuffle
		case "toString":
			return Hx_Field_loreline__node_alternativemode_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__node_alternativemode_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__node_alternativemode_impl__RTTI
}

func (this *Hx_Obj_loreline__node_alternativemode_impl_) Hx_As_Hx_Obj_loreline__node_alternativemode_impl_() *Hx_Obj_loreline__node_alternativemode_impl_ {
	return this
}

var Hx_Field_loreline__node_alternativemode_impl__Sequence int32

func Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Sequence() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__node_alternativemode_impl__Cycle int32

func Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Cycle() int32 {
	return ((int32)(1))
}

var Hx_Field_loreline__node_alternativemode_impl__Once int32

func Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Once() int32 {
	return ((int32)(2))
}

var Hx_Field_loreline__node_alternativemode_impl__Pick int32

func Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Pick() int32 {
	return ((int32)(3))
}

var Hx_Field_loreline__node_alternativemode_impl__Shuffle int32

func Hx_Init_Hx_Field_loreline__node_alternativemode_impl__Shuffle() int32 {
	return ((int32)(4))
}

func Hx_Field_loreline__node_alternativemode_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "Sequence"
		}

	case 1:
		{
			return "Cycle"
		}

	case 2:
		{
			return "Once"
		}

	case 3:
		{
			return "Pick"
		}

	case 4:
		{
			return "Shuffle"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nstringpart() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nstringpart() {
}

type Hx_Obj_VTable_loreline_nstringpart interface {
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nstringpart struct {
	Hx_Obj_loreline_nexpr
	VTable            Hx_Obj_VTable_loreline_nstringpart `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_partType Hx_Enum_loreline_stringparttype
}

func Hx_Obj_loreline_nstringpart_CreateEmptyInstance() *Hx_Obj_loreline_nstringpart {
	obj := &Hx_Obj_loreline_nstringpart{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_partType = HxDefault[Hx_Enum_loreline_stringparttype]()
	return obj
}

func Hx_Obj_loreline_nstringpart_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, partType Hx_Enum_loreline_stringparttype, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nstringpart {
	obj := Hx_Obj_loreline_nstringpart_CreateEmptyInstance()
	obj.Hx_New(id, pos, partType, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nstringpart) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, partType Hx_Enum_loreline_stringparttype, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_partType = partType
		}

	}
}

var Hx_Obj_loreline_nstringpart_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NStringPart",
	HxMakeArray[string](),
	HxMakeArray[string]("partType", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nstringpart_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[Hx_Enum_loreline_stringparttype](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_nstringpart_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nstringpart_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nstringpart) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nstringpart_RTTI
}

func (this *Hx_Obj_loreline_nstringpart) Hx_As_Hx_Obj_loreline_nstringpart() *Hx_Obj_loreline_nstringpart {
	return this
}

func (this *Hx_Obj_loreline_nstringpart) Hx_Field_type() string {
	return "Part"
}

func (this *Hx_Obj_loreline_nstringpart) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	{
		var _g Hx_Enum_loreline_stringparttype
		_ = _g
		_g = this.Hx_Field_partType
		switch _g.Hx_Field_enumIndex() {
		case 0:
			{
			}

		case 1:
			{
				var expr *Hx_Obj_loreline_nexpr
				_ = expr
				expr = _g.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
				{
					var _hx_tmp_0 *Hx_Obj_loreline_node
					_ = _hx_tmp_0
					_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
						_hx_up := expr
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
						_hx_up := this
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})())
					expr.VTable.Hx_Field_each(handleNode)
				}
			}

		case 2:
			{
				var expr *Hx_Obj_loreline_nexpr
				_ = expr
				expr = (func() *Hx_Obj_loreline_nexpr {
					_hx_up := _g.(Hx_Enum_loreline_stringparttype_Tag).Hx_Field_expr
					if _hx_up != nil {
						return &_hx_up.Hx_Obj_loreline_nexpr
					}
					return nil
				})()
				{
					var _hx_tmp_1 *Hx_Obj_loreline_node
					_ = _hx_tmp_1
					_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
						_hx_up := expr
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
					expr.VTable.Hx_Field_each(handleNode)
				}
			}

		}
	}
}

func (this *Hx_Obj_loreline_nstringpart) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	{
		var _g Hx_Enum_loreline_stringparttype
		_ = _g
		_g = this.Hx_Field_partType
		switch _g.Hx_Field_enumIndex() {
		case 0:
			{
				var text string
				_ = text
				text = _g.(Hx_Enum_loreline_stringparttype_Raw).Hx_Field_text
				{
					Hx_Field_go_haxe_hxdynamic_setField(json, "part", "Raw")
					Hx_Field_go_haxe_hxdynamic_setField(json, "text", text)
				}
			}

		case 1:
			{
				var expr *Hx_Obj_loreline_nexpr
				_ = expr
				expr = _g.(Hx_Enum_loreline_stringparttype_Expr).Hx_Field_expr
				{
					Hx_Field_go_haxe_hxdynamic_setField(json, "part", "Expr")
					var _hx_tmp_0 any
					_ = _hx_tmp_0
					_hx_tmp_0 = json
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "expression", expr.VTable.Hx_Field_toJson())
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
					Hx_Field_go_haxe_hxdynamic_setField(json, "part", "Tag")
					Hx_Field_go_haxe_hxdynamic_setField(json, "closing", closing)
					var _hx_tmp_1 any
					_ = _hx_tmp_1
					_hx_tmp_1 = json
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "content", expr.VTable.Hx_Field_toJson())
				}
			}

		}
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nstringliteral() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nstringliteral() {
}

type Hx_Obj_VTable_loreline_nstringliteral interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nstringliteral struct {
	Hx_Obj_loreline_nexpr
	VTable          Hx_Obj_VTable_loreline_nstringliteral `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_parts  HxArray[*Hx_Obj_loreline_nstringpart]
	Hx_Field_quotes int32
}

func Hx_Obj_loreline_nstringliteral_CreateEmptyInstance() *Hx_Obj_loreline_nstringliteral {
	obj := &Hx_Obj_loreline_nstringliteral{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_parts = HxDefault[HxArray[*Hx_Obj_loreline_nstringpart]]()
	obj.Hx_Field_quotes = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nstringliteral_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, quotes int32, parts HxArray[*Hx_Obj_loreline_nstringpart], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nstringliteral {
	obj := Hx_Obj_loreline_nstringliteral_CreateEmptyInstance()
	obj.Hx_New(id, pos, quotes, parts, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, quotes int32, parts HxArray[*Hx_Obj_loreline_nstringpart], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_parts = parts
			this.Hx_Field_quotes = quotes
		}

	}
}

var Hx_Obj_loreline_nstringliteral_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NStringLiteral",
	HxMakeArray[string](),
	HxMakeArray[string]("parts", "quotes", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nstringliteral_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[int32](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_nstringpart]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nstringliteral_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nstringliteral_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nstringliteral_RTTI
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_As_Hx_Obj_loreline_nstringliteral() *Hx_Obj_loreline_nstringliteral {
	return this
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_Field_type() string {
	return "String"
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_parts != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_parts.Len()
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
			part = this.Hx_Field_parts.Get(i)
			var _hx_tmp_1 *Hx_Obj_loreline_node
			_ = _hx_tmp_1
			_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
				_hx_up := part
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
			part.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_nstringliteral) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nstringpart]
		_ = _g2
		_g2 = this.Hx_Field_parts
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var part *Hx_Obj_loreline_nstringpart
			_ = part
			part = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = HxMakeArrayView[any](_g)
			_ = (Hx_Array_Push[any](_hx_tmp_1, part.VTable.Hx_Field_toJson()))
		}
	}

	var parts HxArray[any]
	_ = parts
	parts = _g
	Hx_Field_go_haxe_hxdynamic_setField(json, "parts", parts)
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "quotes", Hx_Field_loreline__quotes_quotes_impl__toString(this.Hx_Field_quotes))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ntextstatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ntextstatement() {
}

type Hx_Obj_VTable_loreline_ntextstatement interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ntextstatement struct {
	Hx_Obj_loreline_astnode
	VTable                  Hx_Obj_VTable_loreline_ntextstatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_content        *Hx_Obj_loreline_nstringliteral
	Hx_Field_condition      *Hx_Obj_loreline_nexpr
	Hx_Field_conditionStyle int32
	Hx_Field_conditionPos   *Hx_Obj_loreline_position
}

func Hx_Obj_loreline_ntextstatement_CreateEmptyInstance() *Hx_Obj_loreline_ntextstatement {
	obj := &Hx_Obj_loreline_ntextstatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_content = HxDefault[*Hx_Obj_loreline_nstringliteral]()
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_conditionStyle = HxDefault[int32]()
	obj.Hx_Field_conditionPos = HxDefault[*Hx_Obj_loreline_position]()
	return obj
}

func Hx_Obj_loreline_ntextstatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, content *Hx_Obj_loreline_nstringliteral, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ntextstatement {
	obj := Hx_Obj_loreline_ntextstatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, content, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, content *Hx_Obj_loreline_nstringliteral, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_content = content
			this.Hx_Field_conditionStyle = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_ntextstatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NTextStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("content", "condition", "conditionStyle", "conditionPos", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ntextstatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nstringliteral](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_ntextstatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ntextstatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ntextstatement_RTTI
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_As_Hx_Obj_loreline_ntextstatement() *Hx_Obj_loreline_ntextstatement {
	return this
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_Field_type() string {
	return "Text"
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_content != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_content
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_content.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_condition != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
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
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_ntextstatement) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "content", this.Hx_Field_content.VTable.Hx_Field_toJson())
	if this.Hx_Field_condition != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "conditionStyle", Hx_Field_loreline__node_conditionstyle_impl__toString(this.Hx_Field_conditionStyle))
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "conditionPos", this.Hx_Field_conditionPos.VTable.Hx_Field_toJson())
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ndialoguestatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ndialoguestatement() {
}

type Hx_Obj_VTable_loreline_ndialoguestatement interface {
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ndialoguestatement struct {
	Hx_Obj_loreline_astnode
	VTable                  Hx_Obj_VTable_loreline_ndialoguestatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_character      string
	Hx_Field_characterPos   *Hx_Obj_loreline_position
	Hx_Field_content        *Hx_Obj_loreline_nstringliteral
	Hx_Field_condition      *Hx_Obj_loreline_nexpr
	Hx_Field_conditionStyle int32
	Hx_Field_conditionPos   *Hx_Obj_loreline_position
}

func Hx_Obj_loreline_ndialoguestatement_CreateEmptyInstance() *Hx_Obj_loreline_ndialoguestatement {
	obj := &Hx_Obj_loreline_ndialoguestatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_character = HxDefault[string]()
	obj.Hx_Field_characterPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_content = HxDefault[*Hx_Obj_loreline_nstringliteral]()
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_conditionStyle = HxDefault[int32]()
	obj.Hx_Field_conditionPos = HxDefault[*Hx_Obj_loreline_position]()
	return obj
}

func Hx_Obj_loreline_ndialoguestatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, character string, characterPos *Hx_Obj_loreline_position, content *Hx_Obj_loreline_nstringliteral, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ndialoguestatement {
	obj := Hx_Obj_loreline_ndialoguestatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, character, characterPos, content, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, character string, characterPos *Hx_Obj_loreline_position, content *Hx_Obj_loreline_nstringliteral, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_character = character
			this.Hx_Field_characterPos = characterPos
			this.Hx_Field_content = content
			this.Hx_Field_conditionStyle = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_ndialoguestatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NDialogueStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("character", "characterPos", "content", "condition", "conditionStyle", "conditionPos", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ndialoguestatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[*Hx_Obj_loreline_position](params.Get(3)), HxConvert[*Hx_Obj_loreline_nstringliteral](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)))
	},
	func() any {
		return Hx_Obj_loreline_ndialoguestatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ndialoguestatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ndialoguestatement_RTTI
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_As_Hx_Obj_loreline_ndialoguestatement() *Hx_Obj_loreline_ndialoguestatement {
	return this
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_Field_type() string {
	return "Dialogue"
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_content != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_content
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_content.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_condition != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
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
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_ndialoguestatement) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	Hx_Field_go_haxe_hxdynamic_setField(json, "character", this.Hx_Field_character)
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "characterPos", this.Hx_Field_characterPos.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "content", this.Hx_Field_content.VTable.Hx_Field_toJson())
	if this.Hx_Field_condition != nil {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "conditionStyle", Hx_Field_loreline__node_conditionstyle_impl__toString(this.Hx_Field_conditionStyle))
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		_hx_tmp_4 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_4, "conditionPos", this.Hx_Field_conditionPos.VTable.Hx_Field_toJson())
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nchoicestatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nchoicestatement() {
}

type Hx_Obj_VTable_loreline_nchoicestatement interface {
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nchoicestatement struct {
	Hx_Obj_loreline_astnode
	VTable           Hx_Obj_VTable_loreline_nchoicestatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_options HxArray[*Hx_Obj_loreline_nchoiceoption]
	Hx_Field_style   int32
}

func Hx_Obj_loreline_nchoicestatement_CreateEmptyInstance() *Hx_Obj_loreline_nchoicestatement {
	obj := &Hx_Obj_loreline_nchoicestatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_options = HxDefault[HxArray[*Hx_Obj_loreline_nchoiceoption]]()
	obj.Hx_Field_style = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nchoicestatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, options HxArray[*Hx_Obj_loreline_nchoiceoption], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nchoicestatement {
	obj := Hx_Obj_loreline_nchoicestatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, options, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nchoicestatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, options HxArray[*Hx_Obj_loreline_nchoiceoption], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_options = options
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nchoicestatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NChoiceStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("options", "style", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nchoicestatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[HxArray[*Hx_Obj_loreline_nchoiceoption]](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_nchoicestatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nchoicestatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nchoicestatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nchoicestatement_RTTI
}

func (this *Hx_Obj_loreline_nchoicestatement) Hx_As_Hx_Obj_loreline_nchoicestatement() *Hx_Obj_loreline_nchoicestatement {
	return this
}

func (this *Hx_Obj_loreline_nchoicestatement) Hx_Field_type() string {
	return "Choice"
}

func (this *Hx_Obj_loreline_nchoicestatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_options != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_options.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_nchoiceoption
			_ = child
			child = this.Hx_Field_options.Get(i)
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

func (this *Hx_Obj_loreline_nchoicestatement) Hx_Field_toJson() any {
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
		var _g2 HxArray[*Hx_Obj_loreline_nchoiceoption]
		_ = _g2
		_g2 = this.Hx_Field_options
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g1
			if !(_hx_tmp_0 < _g2.Len()) {
				break
			}

			var option *Hx_Obj_loreline_nchoiceoption
			_ = option
			option = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_1, option.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "options", _g)
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nchoiceoption() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nchoiceoption() {
}

type Hx_Obj_VTable_loreline_nchoiceoption interface {
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nchoiceoption struct {
	Hx_Obj_loreline_astnode
	VTable                  Hx_Obj_VTable_loreline_nchoiceoption `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_text           *Hx_Obj_loreline_nstringliteral
	Hx_Field_insertion      *Hx_Obj_loreline_ninsertion
	Hx_Field_condition      *Hx_Obj_loreline_nexpr
	Hx_Field_conditionStyle int32
	Hx_Field_conditionPos   *Hx_Obj_loreline_position
	Hx_Field_body           HxArray[*Hx_Obj_loreline_astnode]
	Hx_Field_style          int32
	Hx_Field_once           bool
}

func Hx_Obj_loreline_nchoiceoption_CreateEmptyInstance() *Hx_Obj_loreline_nchoiceoption {
	obj := &Hx_Obj_loreline_nchoiceoption{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_text = HxDefault[*Hx_Obj_loreline_nstringliteral]()
	obj.Hx_Field_insertion = HxDefault[*Hx_Obj_loreline_ninsertion]()
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_conditionStyle = HxDefault[int32]()
	obj.Hx_Field_conditionPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	obj.Hx_Field_style = HxDefault[int32]()
	obj.Hx_Field_once = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_nchoiceoption_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, text *Hx_Obj_loreline_nstringliteral, insertion *Hx_Obj_loreline_ninsertion, condition *Hx_Obj_loreline_nexpr, conditionStyle int32, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nchoiceoption {
	obj := Hx_Obj_loreline_nchoiceoption_CreateEmptyInstance()
	obj.Hx_New(id, pos, text, insertion, condition, conditionStyle, body, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, text *Hx_Obj_loreline_nstringliteral, insertion *Hx_Obj_loreline_ninsertion, condition *Hx_Obj_loreline_nexpr, conditionStyle int32, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_text = text
			this.Hx_Field_insertion = insertion
			this.Hx_Field_condition = condition
			this.Hx_Field_conditionStyle = conditionStyle
			this.Hx_Field_body = body
			this.Hx_Field_style = ((int32)(0))
			this.Hx_Field_once = false
		}

	}
}

var Hx_Obj_loreline_nchoiceoption_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NChoiceOption",
	HxMakeArray[string](),
	HxMakeArray[string]("text", "insertion", "condition", "conditionStyle", "conditionPos", "body", "style", "once", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nchoiceoption_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nstringliteral](params.Get(2)), HxConvert[*Hx_Obj_loreline_ninsertion](params.Get(3)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(4)), HxConvert[int32](params.Get(5)), HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(6)), HxConvert[any](params.Get(7)), HxConvert[any](params.Get(8)))
	},
	func() any {
		return Hx_Obj_loreline_nchoiceoption_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nchoiceoption_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nchoiceoption_RTTI
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_As_Hx_Obj_loreline_nchoiceoption() *Hx_Obj_loreline_nchoiceoption {
	return this
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_Field_type() string {
	return "Option"
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_body.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_astnode
			_ = child
			child = this.Hx_Field_body.Get(i)
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

	if this.Hx_Field_text != nil {
		var _hx_tmp_2 *Hx_Obj_loreline_node
		_ = _hx_tmp_2
		_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_text
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_text.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_insertion != nil {
		var _hx_tmp_3 *Hx_Obj_loreline_node
		_ = _hx_tmp_3
		_hx_tmp_3 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_insertion
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_3, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_insertion.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_condition != nil {
		var _hx_tmp_4 *Hx_Obj_loreline_node
		_ = _hx_tmp_4
		_hx_tmp_4 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nchoiceoption) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if this.Hx_Field_text != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "text", this.Hx_Field_text.VTable.Hx_Field_toJson())
	}

	if this.Hx_Field_insertion != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "insertion", this.Hx_Field_insertion.VTable.Hx_Field_toJson())
	}

	if this.Hx_Field_condition != nil {
		var _hx_tmp_2 any
		_ = _hx_tmp_2
		_hx_tmp_2 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
		var _hx_tmp_3 any
		_ = _hx_tmp_3
		_hx_tmp_3 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "conditionStyle", Hx_Field_loreline__node_conditionstyle_impl__toString(this.Hx_Field_conditionStyle))
		var _hx_tmp_4 any
		_ = _hx_tmp_4
		_hx_tmp_4 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_4, "conditionPos", this.Hx_Field_conditionPos.VTable.Hx_Field_toJson())
	}

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
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g1
			if !(_hx_tmp_5 < _g2.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g2.Get(_g1)
			_g1++
			var _hx_tmp_6 HxArray[any]
			_ = _hx_tmp_6
			_hx_tmp_6 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_6, node.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "body", _g)
	var _hx_tmp_7 any
	_ = _hx_tmp_7
	_hx_tmp_7 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_7, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	if this.Hx_Field_once {
		Hx_Field_go_haxe_hxdynamic_setField(json, "once", true)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nwhenstatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nwhenstatement() {
}

type Hx_Obj_VTable_loreline_nwhenstatement interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nwhenstatement struct {
	Hx_Obj_loreline_astnode
	VTable               Hx_Obj_VTable_loreline_nwhenstatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_strategy    HxNullable[string]
	Hx_Field_strategyPos *Hx_Obj_loreline_position
	Hx_Field_rules       HxArray[*Hx_Obj_loreline_nwhenrule]
	Hx_Field_style       int32
}

func Hx_Obj_loreline_nwhenstatement_CreateEmptyInstance() *Hx_Obj_loreline_nwhenstatement {
	obj := &Hx_Obj_loreline_nwhenstatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_strategy = HxDefault[HxNullable[string]]()
	obj.Hx_Field_strategyPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_rules = HxDefault[HxArray[*Hx_Obj_loreline_nwhenrule]]()
	obj.Hx_Field_style = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nwhenstatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, strategy HxNullable[string], rules HxArray[*Hx_Obj_loreline_nwhenrule], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nwhenstatement {
	obj := Hx_Obj_loreline_nwhenstatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, strategy, rules, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nwhenstatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, strategy HxNullable[string], rules HxArray[*Hx_Obj_loreline_nwhenrule], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_strategy = strategy
			this.Hx_Field_rules = rules
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nwhenstatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NWhenStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("strategy", "strategyPos", "rules", "style", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nwhenstatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[HxNullable[string]](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_nwhenrule]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nwhenstatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nwhenstatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nwhenstatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nwhenstatement_RTTI
}

func (this *Hx_Obj_loreline_nwhenstatement) Hx_As_Hx_Obj_loreline_nwhenstatement() *Hx_Obj_loreline_nwhenstatement {
	return this
}

func (this *Hx_Obj_loreline_nwhenstatement) Hx_Field_type() string {
	return "When"
}

func (this *Hx_Obj_loreline_nwhenstatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_rules != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_rules.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_nwhenrule
			_ = child
			child = this.Hx_Field_rules.Get(i)
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

func (this *Hx_Obj_loreline_nwhenstatement) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if (this.Hx_Field_strategy).Valid != false {
		Hx_Field_go_haxe_hxdynamic_setField(json, "strategy", this.Hx_Field_strategy)
		if this.Hx_Field_strategyPos != nil {
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = json
			Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "strategyPos", this.Hx_Field_strategyPos.VTable.Hx_Field_toJson())
		}
	}

	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nwhenrule]
		_ = _g2
		_g2 = this.Hx_Field_rules
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < _g2.Len()) {
				break
			}

			var rule *Hx_Obj_loreline_nwhenrule
			_ = rule
			rule = _g2.Get(_g1)
			_g1++
			var _hx_tmp_2 HxArray[any]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_2, rule.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "rules", _g)
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nwhenrule() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nwhenrule() {
}

type Hx_Obj_VTable_loreline_nwhenrule interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nwhenrule struct {
	Hx_Obj_loreline_astnode
	VTable                           Hx_Obj_VTable_loreline_nwhenrule `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_condition               *Hx_Obj_loreline_nexpr
	Hx_Field_insertion               *Hx_Obj_loreline_ninsertion
	Hx_Field_insertionCondition      *Hx_Obj_loreline_nexpr
	Hx_Field_insertionConditionStyle int32
	Hx_Field_body                    HxArray[*Hx_Obj_loreline_astnode]
	Hx_Field_style                   int32
	Hx_Field_once                    bool
}

func Hx_Obj_loreline_nwhenrule_CreateEmptyInstance() *Hx_Obj_loreline_nwhenrule {
	obj := &Hx_Obj_loreline_nwhenrule{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_insertion = HxDefault[*Hx_Obj_loreline_ninsertion]()
	obj.Hx_Field_insertionCondition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_insertionConditionStyle = HxDefault[int32]()
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	obj.Hx_Field_style = HxDefault[int32]()
	obj.Hx_Field_once = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_nwhenrule_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nwhenrule {
	obj := Hx_Obj_loreline_nwhenrule_CreateEmptyInstance()
	obj.Hx_New(id, pos, condition, body, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nwhenrule) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_condition = condition
			this.Hx_Field_insertion = nil
			this.Hx_Field_insertionCondition = nil
			this.Hx_Field_insertionConditionStyle = ((int32)(0))
			this.Hx_Field_body = body
			this.Hx_Field_style = ((int32)(0))
			this.Hx_Field_once = false
		}

	}
}

var Hx_Obj_loreline_nwhenrule_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NWhenRule",
	HxMakeArray[string](),
	HxMakeArray[string]("condition", "insertion", "insertionCondition", "insertionConditionStyle", "body", "style", "once", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nwhenrule_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nwhenrule_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nwhenrule_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nwhenrule) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nwhenrule_RTTI
}

func (this *Hx_Obj_loreline_nwhenrule) Hx_As_Hx_Obj_loreline_nwhenrule() *Hx_Obj_loreline_nwhenrule {
	return this
}

func (this *Hx_Obj_loreline_nwhenrule) Hx_Field_type() string {
	return "WhenRule"
}

func (this *Hx_Obj_loreline_nwhenrule) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_condition != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_insertion != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_insertion
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
		this.Hx_Field_insertion.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_insertionCondition != nil {
		var _hx_tmp_2 *Hx_Obj_loreline_node
		_ = _hx_tmp_2
		_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_insertionCondition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_insertionCondition.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_body.Len()
		for _g < _g1 {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_3
			var child *Hx_Obj_loreline_astnode
			_ = child
			child = this.Hx_Field_body.Get(i)
			var _hx_tmp_4 *Hx_Obj_loreline_node
			_ = _hx_tmp_4
			_hx_tmp_4 = (func() *Hx_Obj_loreline_node {
				_hx_up := child
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
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

func (this *Hx_Obj_loreline_nwhenrule) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if this.Hx_Field_condition != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
	}

	if this.Hx_Field_insertion != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "insertion", this.Hx_Field_insertion.VTable.Hx_Field_toJson())
		if this.Hx_Field_insertionCondition != nil {
			var _hx_tmp_2 any
			_ = _hx_tmp_2
			_hx_tmp_2 = json
			Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "insertionCondition", this.Hx_Field_insertionCondition.VTable.Hx_Field_toJson())
			var _hx_tmp_3 any
			_ = _hx_tmp_3
			_hx_tmp_3 = json
			Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "insertionConditionStyle", Hx_Field_loreline__node_conditionstyle_impl__toString(this.Hx_Field_insertionConditionStyle))
		}
	}

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
			var _hx_tmp_4 int32
			_ = _hx_tmp_4
			_hx_tmp_4 = _g1
			if !(_hx_tmp_4 < _g2.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g2.Get(_g1)
			_g1++
			var _hx_tmp_5 HxArray[any]
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_5, node.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "body", _g)
	var _hx_tmp_6 any
	_ = _hx_tmp_6
	_hx_tmp_6 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_6, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	if this.Hx_Field_once {
		Hx_Field_go_haxe_hxdynamic_setField(json, "once", true)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nblock() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nblock() {
}

type Hx_Obj_VTable_loreline_nblock interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nblock struct {
	Hx_Obj_loreline_astnode
	VTable         Hx_Obj_VTable_loreline_nblock `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_body  HxArray[*Hx_Obj_loreline_astnode]
	Hx_Field_style int32
}

func Hx_Obj_loreline_nblock_CreateEmptyInstance() *Hx_Obj_loreline_nblock {
	obj := &Hx_Obj_loreline_nblock{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_body = HxDefault[HxArray[*Hx_Obj_loreline_astnode]]()
	obj.Hx_Field_style = HxDefault[int32]()
	return obj
}

func Hx_Obj_loreline_nblock_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nblock {
	obj := Hx_Obj_loreline_nblock_CreateEmptyInstance()
	obj.Hx_New(id, pos, body, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nblock) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, body HxArray[*Hx_Obj_loreline_astnode], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_body = body
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nblock_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NBlock",
	HxMakeArray[string](),
	HxMakeArray[string]("body", "style", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nblock_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[HxArray[*Hx_Obj_loreline_astnode]](params.Get(2)), HxConvert[any](params.Get(3)), HxConvert[any](params.Get(4)))
	},
	func() any {
		return Hx_Obj_loreline_nblock_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nblock_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nblock) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nblock_RTTI
}

func (this *Hx_Obj_loreline_nblock) Hx_As_Hx_Obj_loreline_nblock() *Hx_Obj_loreline_nblock {
	return this
}

func (this *Hx_Obj_loreline_nblock) Hx_Field_type() string {
	return "Block"
}

func (this *Hx_Obj_loreline_nblock) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_body != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_body.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_astnode
			_ = child
			child = this.Hx_Field_body.Get(i)
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

func (this *Hx_Obj_loreline_nblock) Hx_Field_toJson() any {
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

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g2.Get(_g1)
			_g1++
			var _hx_tmp_1 HxArray[any]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_1, node.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "body", _g)
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nifstatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nifstatement() {
}

type Hx_Obj_VTable_loreline_nifstatement interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nifstatement struct {
	Hx_Obj_loreline_astnode
	VTable                        Hx_Obj_VTable_loreline_nifstatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_condition            *Hx_Obj_loreline_nexpr
	Hx_Field_conditionStyle       int32
	Hx_Field_thenBranch           *Hx_Obj_loreline_nblock
	Hx_Field_elseBranch           *Hx_Obj_loreline_nblock
	Hx_Field_elseLeadingComments  HxArray[*Hx_Obj_loreline_comment]
	Hx_Field_elseTrailingComments HxArray[*Hx_Obj_loreline_comment]
}

func Hx_Obj_loreline_nifstatement_CreateEmptyInstance() *Hx_Obj_loreline_nifstatement {
	obj := &Hx_Obj_loreline_nifstatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_conditionStyle = HxDefault[int32]()
	obj.Hx_Field_thenBranch = HxDefault[*Hx_Obj_loreline_nblock]()
	obj.Hx_Field_elseBranch = HxDefault[*Hx_Obj_loreline_nblock]()
	obj.Hx_Field_elseLeadingComments = HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
	obj.Hx_Field_elseTrailingComments = HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
	return obj
}

func Hx_Obj_loreline_nifstatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, conditionStyle int32, thenBranch *Hx_Obj_loreline_nblock, elseBranch *Hx_Obj_loreline_nblock, leadingComments__isOptional any, trailingComments__isOptional any, elseLeadingComments__isOptional any, elseTrailingComments__isOptional any) *Hx_Obj_loreline_nifstatement {
	obj := Hx_Obj_loreline_nifstatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, condition, conditionStyle, thenBranch, elseBranch, leadingComments__isOptional, trailingComments__isOptional, elseLeadingComments__isOptional, elseTrailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nifstatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, conditionStyle int32, thenBranch *Hx_Obj_loreline_nblock, elseBranch *Hx_Obj_loreline_nblock, leadingComments__isOptional any, trailingComments__isOptional any, elseLeadingComments__isOptional any, elseTrailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		elseLeadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = elseLeadingComments

		if elseLeadingComments__isOptional != nil {
			elseLeadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](elseLeadingComments__isOptional)
		} else {
			elseLeadingComments = nil
		}
		elseTrailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = elseTrailingComments

		if elseTrailingComments__isOptional != nil {
			elseTrailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](elseTrailingComments__isOptional)
		} else {
			elseTrailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_condition = condition
			this.Hx_Field_conditionStyle = conditionStyle
			this.Hx_Field_thenBranch = thenBranch
			this.Hx_Field_elseBranch = elseBranch
			this.Hx_Field_elseLeadingComments = elseLeadingComments
			this.Hx_Field_elseTrailingComments = elseTrailingComments
		}

	}
}

var Hx_Obj_loreline_nifstatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NIfStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("condition", "conditionStyle", "thenBranch", "elseBranch", "elseLeadingComments", "elseTrailingComments", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nifstatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[int32](params.Get(3)), HxConvert[*Hx_Obj_loreline_nblock](params.Get(4)), HxConvert[*Hx_Obj_loreline_nblock](params.Get(5)), HxConvert[any](params.Get(6)), HxConvert[any](params.Get(7)), HxConvert[any](params.Get(8)), HxConvert[any](params.Get(9)))
	},
	func() any {
		return Hx_Obj_loreline_nifstatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nifstatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nifstatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nifstatement_RTTI
}

func (this *Hx_Obj_loreline_nifstatement) Hx_As_Hx_Obj_loreline_nifstatement() *Hx_Obj_loreline_nifstatement {
	return this
}

func (this *Hx_Obj_loreline_nifstatement) Hx_Field_type() string {
	return "If"
}

func (this *Hx_Obj_loreline_nifstatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_condition != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_thenBranch != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_thenBranch
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
		this.Hx_Field_thenBranch.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_elseBranch != nil {
		var _hx_tmp_2 *Hx_Obj_loreline_node
		_ = _hx_tmp_2
		_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_elseBranch
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_elseBranch.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_elseLeadingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_elseLeadingComments
		for {
			var _hx_tmp_3 int32
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			if !(_hx_tmp_3 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			var _hx_tmp_4 *Hx_Obj_loreline_node
			_ = _hx_tmp_4
			_hx_tmp_4 = (func() *Hx_Obj_loreline_node {
				_hx_up := comment
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			comment.VTable.Hx_Field_each(handleNode)
		}
	}

	if this.Hx_Field_elseTrailingComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_comment]
		_ = _g1
		_g1 = this.Hx_Field_elseTrailingComments
		for {
			var _hx_tmp_5 int32
			_ = _hx_tmp_5
			_hx_tmp_5 = _g
			if !(_hx_tmp_5 < _g1.Len()) {
				break
			}

			var comment *Hx_Obj_loreline_comment
			_ = comment
			comment = _g1.Get(_g)
			_g++
			var _hx_tmp_6 *Hx_Obj_loreline_node
			_ = _hx_tmp_6
			_hx_tmp_6 = (func() *Hx_Obj_loreline_node {
				_hx_up := comment
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_6, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			comment.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_nifstatement) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "conditionStyle", Hx_Field_loreline__node_conditionstyle_impl__toString(this.Hx_Field_conditionStyle))
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_astnode]
		_ = _g2
		_g2 = this.Hx_Field_thenBranch.Hx_Field_body
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g1
			if !(_hx_tmp_2 < _g2.Len()) {
				break
			}

			var node *Hx_Obj_loreline_astnode
			_ = node
			node = _g2.Get(_g1)
			_g1++
			var _hx_tmp_3 HxArray[any]
			_ = _hx_tmp_3
			_hx_tmp_3 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_3, node.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "thenBranch", _g)
	var _hx_tmp_4 any
	_ = _hx_tmp_4
	_hx_tmp_4 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_4, "thenStyle", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_thenBranch.Hx_Field_style))
	if this.Hx_Field_elseBranch != nil {
		var tmp__g_1 HxArray[any]
		_ = tmp__g_1
		tmp__g_1 = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_astnode]
			_ = _g2
			_g2 = this.Hx_Field_elseBranch.Hx_Field_body
			for {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = _g1
				if !(_hx_tmp_5 < _g2.Len()) {
					break
				}

				var node *Hx_Obj_loreline_astnode
				_ = node
				node = _g2.Get(_g1)
				_g1++
				var _hx_tmp_6 HxArray[any]
				_ = _hx_tmp_6
				_hx_tmp_6 = tmp__g_1
				_ = (Hx_Array_Push[any](_hx_tmp_6, node.VTable.Hx_Field_toJson()))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "elseBranch", tmp__g_1)
		var _hx_tmp_7 any
		_ = _hx_tmp_7
		_hx_tmp_7 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_7, "elseStyle", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_elseBranch.Hx_Field_style))
		var _hx_tmp_8 bool
		_ = _hx_tmp_8
		var _hx_tmp_9 bool
		_ = _hx_tmp_9
		if this.Hx_Field_elseLeadingComments != nil {
			_hx_tmp_9 = (this.Hx_Field_elseLeadingComments.Len() > 0)
		} else {
			_hx_tmp_9 = false
		}

		if _hx_tmp_9 {
			_hx_tmp_8 = true
		} else {
			var _hx_tmp_10 bool
			_ = _hx_tmp_10
			if this.Hx_Field_elseTrailingComments != nil {
				_hx_tmp_10 = (this.Hx_Field_elseTrailingComments.Len() > 0)
			} else {
				_hx_tmp_10 = false
			}

			_hx_tmp_8 = _hx_tmp_10
		}

		if _hx_tmp_8 {
			var comments any
			_ = comments
			comments = Hx_Field_go_haxe_hxdynamic_getField(json, "comments")
			var _hx_tmp_11 any
			_ = _hx_tmp_11
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(comments, nil))) {
				_hx_tmp_11 = comments
			} else {
				_hx_tmp_11 = any(map[string]any{})
			}

			var comments1 any
			_ = comments1
			comments1 = _hx_tmp_11
			var _hx_tmp_12 bool
			_ = _hx_tmp_12
			if this.Hx_Field_elseLeadingComments != nil {
				_hx_tmp_12 = (this.Hx_Field_elseLeadingComments.Len() > 0)
			} else {
				_hx_tmp_12 = false
			}

			if _hx_tmp_12 {
				var tmp__g_2 HxArray[any]
				_ = tmp__g_2
				tmp__g_2 = HxMakeArray[any]().Dyn()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 HxArray[*Hx_Obj_loreline_comment]
					_ = _g2
					_g2 = this.Hx_Field_elseLeadingComments
					for {
						var _hx_tmp_13 int32
						_ = _hx_tmp_13
						_hx_tmp_13 = _g1
						if !(_hx_tmp_13 < _g2.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = _g2.Get(_g1)
						_g1++
						var _hx_tmp_14 HxArray[any]
						_ = _hx_tmp_14
						_hx_tmp_14 = tmp__g_2
						_ = (Hx_Array_Push[any](_hx_tmp_14, c.VTable.Hx_Field_toJson()))
					}
				}

				Hx_Field_go_haxe_hxdynamic_setField(comments1, "elseLeading", tmp__g_2)
			}

			var _hx_tmp_15 bool
			_ = _hx_tmp_15
			if this.Hx_Field_elseTrailingComments != nil {
				_hx_tmp_15 = (this.Hx_Field_elseTrailingComments.Len() > 0)
			} else {
				_hx_tmp_15 = false
			}

			if _hx_tmp_15 {
				var tmp__g_2 HxArray[any]
				_ = tmp__g_2
				tmp__g_2 = HxMakeArray[any]().Dyn()
				{
					var _g1 int32
					_ = _g1
					_g1 = 0
					var _g2 HxArray[*Hx_Obj_loreline_comment]
					_ = _g2
					_g2 = this.Hx_Field_elseTrailingComments
					for {
						var _hx_tmp_16 int32
						_ = _hx_tmp_16
						_hx_tmp_16 = _g1
						if !(_hx_tmp_16 < _g2.Len()) {
							break
						}

						var c *Hx_Obj_loreline_comment
						_ = c
						c = _g2.Get(_g1)
						_g1++
						var _hx_tmp_17 HxArray[any]
						_ = _hx_tmp_17
						_hx_tmp_17 = tmp__g_2
						_ = (Hx_Array_Push[any](_hx_tmp_17, c.VTable.Hx_Field_toJson()))
					}
				}

				Hx_Field_go_haxe_hxdynamic_setField(comments1, "elseTrailing", tmp__g_2)
			}

			Hx_Field_go_haxe_hxdynamic_setField(json, "comments", comments1)
		}
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nalternative() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nalternative() {
}

type Hx_Obj_VTable_loreline_nalternative interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nalternative struct {
	Hx_Obj_loreline_astnode
	VTable                     Hx_Obj_VTable_loreline_nalternative `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_mode              int32
	Hx_Field_items             HxArray[*Hx_Obj_loreline_nblock]
	Hx_Field_style             int32
	Hx_Field_separatorComments HxArray[HxArray[*Hx_Obj_loreline_comment]]
}

func Hx_Obj_loreline_nalternative_CreateEmptyInstance() *Hx_Obj_loreline_nalternative {
	obj := &Hx_Obj_loreline_nalternative{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_mode = HxDefault[int32]()
	obj.Hx_Field_items = HxDefault[HxArray[*Hx_Obj_loreline_nblock]]()
	obj.Hx_Field_style = HxDefault[int32]()
	obj.Hx_Field_separatorComments = HxDefault[HxArray[HxArray[*Hx_Obj_loreline_comment]]]()
	return obj
}

func Hx_Obj_loreline_nalternative_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, mode int32, items HxArray[*Hx_Obj_loreline_nblock], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nalternative {
	obj := Hx_Obj_loreline_nalternative_CreateEmptyInstance()
	obj.Hx_New(id, pos, mode, items, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nalternative) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, mode int32, items HxArray[*Hx_Obj_loreline_nblock], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_mode = mode
			this.Hx_Field_items = items
			this.Hx_Field_style = ((int32)(0))
		}

	}
}

var Hx_Obj_loreline_nalternative_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NAlternative",
	HxMakeArray[string](),
	HxMakeArray[string]("mode", "items", "style", "separatorComments", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nalternative_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[int32](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_nblock]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nalternative_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nalternative_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nalternative) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nalternative_RTTI
}

func (this *Hx_Obj_loreline_nalternative) Hx_As_Hx_Obj_loreline_nalternative() *Hx_Obj_loreline_nalternative {
	return this
}

func (this *Hx_Obj_loreline_nalternative) Hx_Field_type() string {
	return "Alternative"
}

func (this *Hx_Obj_loreline_nalternative) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_items != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = this.Hx_Field_items.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			var child *Hx_Obj_loreline_nblock
			_ = child
			child = this.Hx_Field_items.Get(i)
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

	if this.Hx_Field_separatorComments != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[HxArray[*Hx_Obj_loreline_comment]]
		_ = _g1
		_g1 = this.Hx_Field_separatorComments
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < _g1.Len()) {
				break
			}

			var comments HxArray[*Hx_Obj_loreline_comment]
			_ = comments
			comments = _g1.Get(_g)
			_g++
			if comments != nil {
				var tmp__g_1 int32
				_ = tmp__g_1
				tmp__g_1 = 0
				for {
					var _hx_tmp_3 int32
					_ = _hx_tmp_3
					_hx_tmp_3 = tmp__g_1
					if !(_hx_tmp_3 < comments.Len()) {
						break
					}

					var comment *Hx_Obj_loreline_comment
					_ = comment
					comment = comments.Get(tmp__g_1)
					tmp__g_1++
					var _hx_tmp_4 *Hx_Obj_loreline_node
					_ = _hx_tmp_4
					_hx_tmp_4 = (func() *Hx_Obj_loreline_node {
						_hx_up := comment
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})()
					handleNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
						_hx_up := this
						if _hx_up != nil {
							return &_hx_up.Hx_Obj_loreline_node
						}
						return nil
					})())
					comment.VTable.Hx_Field_each(handleNode)
				}
			}
		}
	}
}

func (this *Hx_Obj_loreline_nalternative) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "mode", Hx_Field_loreline__node_alternativemode_impl__toString(this.Hx_Field_mode))
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nblock]
		_ = _g2
		_g2 = this.Hx_Field_items
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g1
			if !(_hx_tmp_1 < _g2.Len()) {
				break
			}

			var item *Hx_Obj_loreline_nblock
			_ = item
			item = _g2.Get(_g1)
			_g1++
			var _hx_tmp_2 HxArray[any]
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_2, item.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "items", _g)
	var _hx_tmp_3 any
	_ = _hx_tmp_3
	_hx_tmp_3 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_3, "style", Hx_Field_loreline__node_blockstyle_impl__toString(this.Hx_Field_style))
	if this.Hx_Field_separatorComments != nil {
		var tmp__g_1 HxArray[any]
		_ = tmp__g_1
		tmp__g_1 = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[HxArray[*Hx_Obj_loreline_comment]]
			_ = _g2
			_g2 = this.Hx_Field_separatorComments
			for {
				var _hx_tmp_4 int32
				_ = _hx_tmp_4
				_hx_tmp_4 = _g1
				if !(_hx_tmp_4 < _g2.Len()) {
					break
				}

				var comments HxArray[*Hx_Obj_loreline_comment]
				_ = comments
				comments = _g2.Get(_g1)
				_g1++
				var tmp any
				_ = tmp
				if comments != nil {
					var tmp__g_2 HxArray[any]
					_ = tmp__g_2
					tmp__g_2 = HxMakeArray[any]().Dyn()
					{
						var tmp__g1_1 int32
						_ = tmp__g1_1
						tmp__g1_1 = 0
						for {
							var _hx_tmp_5 int32
							_ = _hx_tmp_5
							_hx_tmp_5 = tmp__g1_1
							if !(_hx_tmp_5 < comments.Len()) {
								break
							}

							var c *Hx_Obj_loreline_comment
							_ = c
							c = comments.Get(tmp__g1_1)
							tmp__g1_1++
							var _hx_tmp_6 HxArray[any]
							_ = _hx_tmp_6
							_hx_tmp_6 = tmp__g_2
							_ = (Hx_Array_Push[any](_hx_tmp_6, c.VTable.Hx_Field_toJson()))
						}
					}

					tmp = ((any)(tmp__g_2))
				} else {
					tmp = nil
				}

				_ = (Hx_Array_Push[any](tmp__g_1, tmp))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "separatorComments", tmp__g_1)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ncall() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ncall() {
}

type Hx_Obj_VTable_loreline_ncall interface {
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ncall struct {
	Hx_Obj_loreline_nexpr
	VTable          Hx_Obj_VTable_loreline_ncall `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target *Hx_Obj_loreline_nexpr
	Hx_Field_args   HxArray[*Hx_Obj_loreline_nexpr]
}

func Hx_Obj_loreline_ncall_CreateEmptyInstance() *Hx_Obj_loreline_ncall {
	obj := &Hx_Obj_loreline_ncall{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_args = HxDefault[HxArray[*Hx_Obj_loreline_nexpr]]()
	return obj
}

func Hx_Obj_loreline_ncall_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, args HxArray[*Hx_Obj_loreline_nexpr], leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ncall {
	obj := Hx_Obj_loreline_ncall_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, args, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ncall) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, args HxArray[*Hx_Obj_loreline_nexpr], leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_args = args
		}

	}
}

var Hx_Obj_loreline_ncall_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NCall",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "args", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ncall_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[HxArray[*Hx_Obj_loreline_nexpr]](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_ncall_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ncall_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ncall) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ncall_RTTI
}

func (this *Hx_Obj_loreline_ncall) Hx_As_Hx_Obj_loreline_ncall() *Hx_Obj_loreline_ncall {
	return this
}

func (this *Hx_Obj_loreline_ncall) Hx_Field_type() string {
	return "Call"
}

func (this *Hx_Obj_loreline_ncall) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_target != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_target
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_target.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_args != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g1
		_g1 = this.Hx_Field_args
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = _g1.Get(_g)
			_g++
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			arg.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_ncall) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "target", this.Hx_Field_target.VTable.Hx_Field_toJson())
	var _g HxArray[any]
	_ = _g
	_g = HxMakeArray[any]().Dyn()
	{
		var _g1 int32
		_ = _g1
		_g1 = 0
		var _g2 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g2
		_g2 = this.Hx_Field_args
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
			_hx_tmp_2 = _g
			_ = (Hx_Array_Push[any](_hx_tmp_2, arg.VTable.Hx_Field_toJson()))
		}
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "args", _g)
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ntransition() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ntransition() {
}

type Hx_Obj_VTable_loreline_ntransition interface {
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ntransition struct {
	Hx_Obj_loreline_astnode
	VTable              Hx_Obj_VTable_loreline_ntransition `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target     string
	Hx_Field_targetPos  *Hx_Obj_loreline_position
	Hx_Field_targetExpr *Hx_Obj_loreline_nexpr
	Hx_Field_args       HxArray[*Hx_Obj_loreline_nexpr]
}

func Hx_Obj_loreline_ntransition_CreateEmptyInstance() *Hx_Obj_loreline_ntransition {
	obj := &Hx_Obj_loreline_ntransition{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[string]()
	obj.Hx_Field_targetPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_targetExpr = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_args = HxDefault[HxArray[*Hx_Obj_loreline_nexpr]]()
	return obj
}

func Hx_Obj_loreline_ntransition_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target string, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ntransition {
	obj := Hx_Obj_loreline_ntransition_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, targetPos, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ntransition) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target string, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Field_targetExpr = nil
		}
		{
			this.Hx_Field_args = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_targetPos = targetPos
		}

	}
}

var Hx_Obj_loreline_ntransition_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NTransition",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "targetPos", "targetExpr", "args", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ntransition_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[*Hx_Obj_loreline_position](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_ntransition_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ntransition_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ntransition) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ntransition_RTTI
}

func (this *Hx_Obj_loreline_ntransition) Hx_As_Hx_Obj_loreline_ntransition() *Hx_Obj_loreline_ntransition {
	return this
}

func (this *Hx_Obj_loreline_ntransition) Hx_Field_type() string {
	return "Transition"
}

func (this *Hx_Obj_loreline_ntransition) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_targetExpr != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_targetExpr
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_targetExpr.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_args != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g1
		_g1 = this.Hx_Field_args
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = _g1.Get(_g)
			_g++
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			arg.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_ntransition) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if this.Hx_Field_targetExpr != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "targetExpr", this.Hx_Field_targetExpr.VTable.Hx_Field_toJson())
	} else {
		Hx_Field_go_haxe_hxdynamic_setField(json, "target", this.Hx_Field_target)
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "targetPos", this.Hx_Field_targetPos.VTable.Hx_Field_toJson())
	if this.Hx_Field_args != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _g2
			_g2 = this.Hx_Field_args
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < _g2.Len()) {
					break
				}

				var arg *Hx_Obj_loreline_nexpr
				_ = arg
				arg = _g2.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[any]
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_ = (Hx_Array_Push[any](_hx_tmp_3, arg.VTable.Hx_Field_toJson()))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "args", _g)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_ninsertion() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_ninsertion() {
}

type Hx_Obj_VTable_loreline_ninsertion interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_ninsertion struct {
	Hx_Obj_loreline_astnode
	VTable              Hx_Obj_VTable_loreline_ninsertion `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target     string
	Hx_Field_targetPos  *Hx_Obj_loreline_position
	Hx_Field_targetExpr *Hx_Obj_loreline_nexpr
	Hx_Field_args       HxArray[*Hx_Obj_loreline_nexpr]
}

func Hx_Obj_loreline_ninsertion_CreateEmptyInstance() *Hx_Obj_loreline_ninsertion {
	obj := &Hx_Obj_loreline_ninsertion{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[string]()
	obj.Hx_Field_targetPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_targetExpr = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_args = HxDefault[HxArray[*Hx_Obj_loreline_nexpr]]()
	return obj
}

func Hx_Obj_loreline_ninsertion_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target string, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_ninsertion {
	obj := Hx_Obj_loreline_ninsertion_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, targetPos, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_ninsertion) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target string, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Field_targetExpr = nil
		}
		{
			this.Hx_Field_args = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_targetPos = targetPos
		}

	}
}

var Hx_Obj_loreline_ninsertion_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NInsertion",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "targetPos", "targetExpr", "args", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_ninsertion_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[string](params.Get(2)), HxConvert[*Hx_Obj_loreline_position](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_ninsertion_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_ninsertion_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_ninsertion) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_ninsertion_RTTI
}

func (this *Hx_Obj_loreline_ninsertion) Hx_As_Hx_Obj_loreline_ninsertion() *Hx_Obj_loreline_ninsertion {
	return this
}

func (this *Hx_Obj_loreline_ninsertion) Hx_Field_type() string {
	return "Insertion"
}

func (this *Hx_Obj_loreline_ninsertion) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_targetExpr != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_targetExpr
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_targetExpr.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_args != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g1
		_g1 = this.Hx_Field_args
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = _g1.Get(_g)
			_g++
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			arg.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_ninsertion) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	if this.Hx_Field_targetExpr != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "targetExpr", this.Hx_Field_targetExpr.VTable.Hx_Field_toJson())
	} else {
		Hx_Field_go_haxe_hxdynamic_setField(json, "target", this.Hx_Field_target)
	}

	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "targetPos", this.Hx_Field_targetPos.VTable.Hx_Field_toJson())
	if this.Hx_Field_args != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _g2
			_g2 = this.Hx_Field_args
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < _g2.Len()) {
					break
				}

				var arg *Hx_Obj_loreline_nexpr
				_ = arg
				arg = _g2.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[any]
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_ = (Hx_Array_Push[any](_hx_tmp_3, arg.VTable.Hx_Field_toJson()))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "args", _g)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nbeatcall() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nbeatcall() {
}

type Hx_Obj_VTable_loreline_nbeatcall interface {
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nbeatcall struct {
	Hx_Obj_loreline_astnode
	VTable              Hx_Obj_VTable_loreline_nbeatcall `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_targetExpr *Hx_Obj_loreline_nexpr
	Hx_Field_targetPos  *Hx_Obj_loreline_position
	Hx_Field_args       HxArray[*Hx_Obj_loreline_nexpr]
}

func Hx_Obj_loreline_nbeatcall_CreateEmptyInstance() *Hx_Obj_loreline_nbeatcall {
	obj := &Hx_Obj_loreline_nbeatcall{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_targetExpr = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_targetPos = HxDefault[*Hx_Obj_loreline_position]()
	obj.Hx_Field_args = HxDefault[HxArray[*Hx_Obj_loreline_nexpr]]()
	return obj
}

func Hx_Obj_loreline_nbeatcall_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, targetExpr *Hx_Obj_loreline_nexpr, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nbeatcall {
	obj := Hx_Obj_loreline_nbeatcall_CreateEmptyInstance()
	obj.Hx_New(id, pos, targetExpr, targetPos, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, targetExpr *Hx_Obj_loreline_nexpr, targetPos *Hx_Obj_loreline_position, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Field_args = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_targetExpr = targetExpr
			this.Hx_Field_targetPos = targetPos
		}

	}
}

var Hx_Obj_loreline_nbeatcall_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NBeatCall",
	HxMakeArray[string](),
	HxMakeArray[string]("targetExpr", "targetPos", "args", "type", "each", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nbeatcall_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[*Hx_Obj_loreline_position](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nbeatcall_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nbeatcall_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nbeatcall_RTTI
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_As_Hx_Obj_loreline_nbeatcall() *Hx_Obj_loreline_nbeatcall {
	return this
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_Field_type() string {
	return "BeatCall"
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_targetExpr != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_targetExpr
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_targetExpr.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_args != nil {
		var _g int32
		_ = _g
		_g = 0
		var _g1 HxArray[*Hx_Obj_loreline_nexpr]
		_ = _g1
		_g1 = this.Hx_Field_args
		for {
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = _g
			if !(_hx_tmp_1 < _g1.Len()) {
				break
			}

			var arg *Hx_Obj_loreline_nexpr
			_ = arg
			arg = _g1.Get(_g)
			_g++
			var _hx_tmp_2 *Hx_Obj_loreline_node
			_ = _hx_tmp_2
			_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
				_hx_up := arg
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})()
			handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
				_hx_up := this
				if _hx_up != nil {
					return &_hx_up.Hx_Obj_loreline_node
				}
				return nil
			})())
			arg.VTable.Hx_Field_each(handleNode)
		}
	}
}

func (this *Hx_Obj_loreline_nbeatcall) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "targetExpr", this.Hx_Field_targetExpr.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "targetPos", this.Hx_Field_targetPos.VTable.Hx_Field_toJson())
	if this.Hx_Field_args != nil {
		var _g HxArray[any]
		_ = _g
		_g = HxMakeArray[any]().Dyn()
		{
			var _g1 int32
			_ = _g1
			_g1 = 0
			var _g2 HxArray[*Hx_Obj_loreline_nexpr]
			_ = _g2
			_g2 = this.Hx_Field_args
			for {
				var _hx_tmp_2 int32
				_ = _hx_tmp_2
				_hx_tmp_2 = _g1
				if !(_hx_tmp_2 < _g2.Len()) {
					break
				}

				var arg *Hx_Obj_loreline_nexpr
				_ = arg
				arg = _g2.Get(_g1)
				_g1++
				var _hx_tmp_3 HxArray[any]
				_ = _hx_tmp_3
				_hx_tmp_3 = _g
				_ = (Hx_Array_Push[any](_hx_tmp_3, arg.VTable.Hx_Field_toJson()))
			}
		}

		Hx_Field_go_haxe_hxdynamic_setField(json, "args", _g)
	}

	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nfunctiondecl() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nfunctiondecl() {
}

type Hx_Obj_VTable_loreline_nfunctiondecl interface {
	Hx_Field_type() string
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nfunctiondecl struct {
	Hx_Obj_loreline_nexpr
	VTable            Hx_Obj_VTable_loreline_nfunctiondecl `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_name     HxNullable[string]
	Hx_Field_args     HxArray[string]
	Hx_Field_code     string
	Hx_Field_external bool
}

func Hx_Obj_loreline_nfunctiondecl_CreateEmptyInstance() *Hx_Obj_loreline_nfunctiondecl {
	obj := &Hx_Obj_loreline_nfunctiondecl{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_name = HxDefault[HxNullable[string]]()
	obj.Hx_Field_args = HxDefault[HxArray[string]]()
	obj.Hx_Field_code = HxDefault[string]()
	obj.Hx_Field_external = HxDefault[bool]()
	return obj
}

func Hx_Obj_loreline_nfunctiondecl_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name HxNullable[string], args HxArray[string], code string, external bool, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nfunctiondecl {
	obj := Hx_Obj_loreline_nfunctiondecl_CreateEmptyInstance()
	obj.Hx_New(id, pos, name, args, code, external, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nfunctiondecl) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, name HxNullable[string], args HxArray[string], code string, external bool, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_name = name
			this.Hx_Field_args = args
			this.Hx_Field_code = code
			this.Hx_Field_external = external
		}

	}
}

var Hx_Obj_loreline_nfunctiondecl_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NFunctionDecl",
	HxMakeArray[string](),
	HxMakeArray[string]("name", "args", "code", "external", "type", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "each", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nfunctiondecl_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[HxNullable[string]](params.Get(2)), HxConvert[HxArray[string]](params.Get(3)), HxConvert[string](params.Get(4)), HxConvert[bool](params.Get(5)), HxConvert[any](params.Get(6)), HxConvert[any](params.Get(7)))
	},
	func() any {
		return Hx_Obj_loreline_nfunctiondecl_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nfunctiondecl_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nfunctiondecl) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nfunctiondecl_RTTI
}

func (this *Hx_Obj_loreline_nfunctiondecl) Hx_As_Hx_Obj_loreline_nfunctiondecl() *Hx_Obj_loreline_nfunctiondecl {
	return this
}

func (this *Hx_Obj_loreline_nfunctiondecl) Hx_Field_type() string {
	return "Function"
}

func (this *Hx_Obj_loreline_nfunctiondecl) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	if (this.Hx_Field_name).Valid != false {
		Hx_Field_go_haxe_hxdynamic_setField(json, "name", this.Hx_Field_name)
	}

	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "args", (Hx_Array_Concat[string](HxMakeArray[string](), this.Hx_Field_args)))
	Hx_Field_go_haxe_hxdynamic_setField(json, "code", this.Hx_Field_code)
	Hx_Field_go_haxe_hxdynamic_setField(json, "external", this.Hx_Field_external)
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nliteral() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nliteral() {
}

type Hx_Obj_VTable_loreline_nliteral interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nliteral struct {
	Hx_Obj_loreline_nexpr
	VTable               Hx_Obj_VTable_loreline_nliteral `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_value       any
	Hx_Field_literalType Hx_Enum_loreline_literaltype
}

func Hx_Obj_loreline_nliteral_CreateEmptyInstance() *Hx_Obj_loreline_nliteral {
	obj := &Hx_Obj_loreline_nliteral{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_value = HxDefault[any]()
	obj.Hx_Field_literalType = HxDefault[Hx_Enum_loreline_literaltype]()
	return obj
}

func Hx_Obj_loreline_nliteral_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, value any, literalType Hx_Enum_loreline_literaltype, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nliteral {
	obj := Hx_Obj_loreline_nliteral_CreateEmptyInstance()
	obj.Hx_New(id, pos, value, literalType, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nliteral) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, value any, literalType Hx_Enum_loreline_literaltype, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_value = value
			this.Hx_Field_literalType = literalType
		}

	}
}

var Hx_Obj_loreline_nliteral_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NLiteral",
	HxMakeArray[string](),
	HxMakeArray[string]("value", "literalType", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nliteral_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[any](params.Get(2)), HxConvert[Hx_Enum_loreline_literaltype](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nliteral_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nliteral_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nliteral) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nliteral_RTTI
}

func (this *Hx_Obj_loreline_nliteral) Hx_As_Hx_Obj_loreline_nliteral() *Hx_Obj_loreline_nliteral {
	return this
}

func (this *Hx_Obj_loreline_nliteral) Hx_Field_type() string {
	return "Literal"
}

func (this *Hx_Obj_loreline_nliteral) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	{
		var _g Hx_Enum_loreline_literaltype
		_ = _g
		_g = this.Hx_Field_literalType
		switch _g.Hx_Field_enumIndex() {
		case 0, 1, 2:
			{
			}

		case 3:
			{
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_value, nil))) {
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var _g1 HxArray[any]
					_ = _g1
					_g1 = HxAnyToArray(this.Hx_Field_value)
					for {
						var _hx_tmp_0 int32
						_ = _hx_tmp_0
						_hx_tmp_0 = tmp__g_1
						if !(_hx_tmp_0 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g1)) {
							break
						}

						var elem any
						_ = elem
						elem = _g1.Get(tmp__g_1)
						tmp__g_1++
						if Hx_Field_std_isOfType(elem, Hx_Obj_loreline_node_RTTI) {
							var hx_dyncast_29 any
							_ = hx_dyncast_29
							hx_dyncast_29 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_node")
							var _hx_tmp_1 *Hx_Obj_loreline_node
							_ = _hx_tmp_1
							if hx_dyncast_29 != nil {
								_hx_tmp_1 = hx_dyncast_29.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
							} else {
								_hx_tmp_1 = nil
							}

							var node *Hx_Obj_loreline_node
							_ = node
							node = _hx_tmp_1
							var _hx_tmp_2 *Hx_Obj_loreline_node
							_ = _hx_tmp_2
							_hx_tmp_2 = node
							handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
								_hx_up := this
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							node.VTable.Hx_Field_each(handleNode)
						}
					}
				}
			}

		case 4:
			{
				var style int32
				_ = style
				style = _g.(Hx_Enum_loreline_literaltype_Object).Hx_Field_style
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_value, nil))) {
					var tmp__g_1 int32
					_ = tmp__g_1
					tmp__g_1 = 0
					var _g1 HxArray[*Hx_Obj_loreline_nobjectfield]
					_ = _g1
					_g1 = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](this.Hx_Field_value)
					for {
						var _hx_tmp_3 int32
						_ = _hx_tmp_3
						_hx_tmp_3 = tmp__g_1
						if !(_hx_tmp_3 < _g1.Len()) {
							break
						}

						var field *Hx_Obj_loreline_nobjectfield
						_ = field
						field = _g1.Get(tmp__g_1)
						tmp__g_1++
						if Hx_Field_std_isOfType(field.Hx_Field_value, Hx_Obj_loreline_node_RTTI) {
							var node *Hx_Obj_loreline_node
							_ = node
							node = (func() *Hx_Obj_loreline_node {
								_hx_up := field.Hx_Field_value
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})()
							var _hx_tmp_4 *Hx_Obj_loreline_node
							_ = _hx_tmp_4
							_hx_tmp_4 = node
							handleNode(_hx_tmp_4, (func() *Hx_Obj_loreline_node {
								_hx_up := this
								if _hx_up != nil {
									return &_hx_up.Hx_Obj_loreline_node
								}
								return nil
							})())
							node.VTable.Hx_Field_each(handleNode)
						}
					}
				}
			}

		}
	}
}

func (this *Hx_Obj_loreline_nliteral) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "literal", Hx_Field_type_enumConstructor(((Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue)(this.Hx_Field_literalType))))
	{
		var _g Hx_Enum_loreline_literaltype
		_ = _g
		_g = this.Hx_Field_literalType
		switch _g.Hx_Field_enumIndex() {
		case 0, 1, 2:
			{
				Hx_Field_go_haxe_hxdynamic_setField(json, "value", this.Hx_Field_value)
			}

		case 3:
			{
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_value, nil))) {
					var tmp__g_1 HxArray[any]
					_ = tmp__g_1
					tmp__g_1 = HxMakeArray[any]().Dyn()
					{
						var _g1 int32
						_ = _g1
						_g1 = 0
						var _g2 HxArray[any]
						_ = _g2
						_g2 = HxAnyToArray(this.Hx_Field_value)
						for {
							var _hx_tmp_1 int32
							_ = _hx_tmp_1
							_hx_tmp_1 = _g1
							if !(_hx_tmp_1 < Hx_Field_go_haxe_hxdynamic_getArrayLength(_g2)) {
								break
							}

							var elem any
							_ = elem
							elem = _g2.Get(_g1)
							_g1++
							if Hx_Field_std_isOfType(elem, Hx_Obj_loreline_node_RTTI) {
								var _hx_tmp_2 HxArray[any]
								_ = _hx_tmp_2
								_hx_tmp_2 = tmp__g_1
								var hx_dyncast_30 any
								_ = hx_dyncast_30
								hx_dyncast_30 = Hx_Field_go_haxe_hxdynamic_castClass(elem, "Hx_Obj_loreline_node")
								var _hx_tmp_3 *Hx_Obj_loreline_node
								_ = _hx_tmp_3
								if hx_dyncast_30 != nil {
									_hx_tmp_3 = hx_dyncast_30.(interface{ Hx_As_Hx_Obj_loreline_node() *Hx_Obj_loreline_node }).Hx_As_Hx_Obj_loreline_node()
								} else {
									_hx_tmp_3 = nil
								}

								_ = (Hx_Array_Push[any](_hx_tmp_2, _hx_tmp_3.VTable.Hx_Field_toJson()))
							} else {
								_ = (Hx_Array_Push[any](tmp__g_1, elem))
							}
						}
					}

					Hx_Field_go_haxe_hxdynamic_setField(json, "value", tmp__g_1)
				}
			}

		case 4:
			{
				var style int32
				_ = style
				style = _g.(Hx_Enum_loreline_literaltype_Object).Hx_Field_style
				{
					if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(this.Hx_Field_value, nil))) {
						var tmp__g_1 HxArray[any]
						_ = tmp__g_1
						tmp__g_1 = HxMakeArray[any]().Dyn()
						{
							var _g1 int32
							_ = _g1
							_g1 = 0
							var _g2 HxArray[*Hx_Obj_loreline_nobjectfield]
							_ = _g2
							_g2 = HxAnyToTypedArray[*Hx_Obj_loreline_nobjectfield](this.Hx_Field_value)
							for {
								var _hx_tmp_4 int32
								_ = _hx_tmp_4
								_hx_tmp_4 = _g1
								if !(_hx_tmp_4 < _g2.Len()) {
									break
								}

								var field *Hx_Obj_loreline_nobjectfield
								_ = field
								field = _g2.Get(_g1)
								_g1++
								var _hx_tmp_5 HxArray[any]
								_ = _hx_tmp_5
								_hx_tmp_5 = tmp__g_1
								_ = (Hx_Array_Push[any](_hx_tmp_5, field.VTable.Hx_Field_toJson()))
							}
						}

						Hx_Field_go_haxe_hxdynamic_setField(json, "value", tmp__g_1)
					} else {
						var _hx_tmp_6 any
						_ = _hx_tmp_6
						_hx_tmp_6 = json
						Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_6, "value", HxMakeArray[any]().Dyn())
					}

					var _hx_tmp_7 any
					_ = _hx_tmp_7
					_hx_tmp_7 = json
					Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_7, "style", Hx_Field_loreline__node_blockstyle_impl__toString(style))
				}
			}

		}
	}

	return json
}

type Hx_Enum_loreline_literaltype interface {
	Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue
	M_Hx_Enum_loreline_literaltype()
}

type Hx_Enum_loreline_literaltype_Object struct {
	Hx_Field_style int32
}

func (this Hx_Enum_loreline_literaltype_Object) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_literaltype_Object) M_Hx_Enum_loreline_literaltype()         {}
func (this Hx_Enum_loreline_literaltype_Object) Hx_Field_enumIndex() int32               { return 4 }
func (this Hx_Enum_loreline_literaltype_Object) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_literaltype_RTTI
}
func (this Hx_Enum_loreline_literaltype_Object) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any](any(this.Hx_Field_style))
}
func (this Hx_Enum_loreline_literaltype_Object) Hx_Field_enumParameter(index int32) any {
	switch index {
	case 0:
		return any(this.Hx_Field_style)
	default:
		return nil
	}
}

type Hx_Enum_loreline_literaltype_Number struct{}

func (this Hx_Enum_loreline_literaltype_Number) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_literaltype_Number) M_Hx_Enum_loreline_literaltype()         {}
func (this Hx_Enum_loreline_literaltype_Number) Hx_Field_enumIndex() int32               { return 0 }
func (this Hx_Enum_loreline_literaltype_Number) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_literaltype_RTTI
}
func (this Hx_Enum_loreline_literaltype_Number) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_literaltype_Number) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_literaltype_Null struct{}

func (this Hx_Enum_loreline_literaltype_Null) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_literaltype_Null) M_Hx_Enum_loreline_literaltype()         {}
func (this Hx_Enum_loreline_literaltype_Null) Hx_Field_enumIndex() int32               { return 2 }
func (this Hx_Enum_loreline_literaltype_Null) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_literaltype_RTTI
}
func (this Hx_Enum_loreline_literaltype_Null) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_literaltype_Null) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_literaltype_Boolean struct{}

func (this Hx_Enum_loreline_literaltype_Boolean) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_literaltype_Boolean) M_Hx_Enum_loreline_literaltype()         {}
func (this Hx_Enum_loreline_literaltype_Boolean) Hx_Field_enumIndex() int32               { return 1 }
func (this Hx_Enum_loreline_literaltype_Boolean) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_literaltype_RTTI
}
func (this Hx_Enum_loreline_literaltype_Boolean) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_literaltype_Boolean) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

type Hx_Enum_loreline_literaltype_Array struct{}

func (this Hx_Enum_loreline_literaltype_Array) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass { return nil }
func (this Hx_Enum_loreline_literaltype_Array) M_Hx_Enum_loreline_literaltype()         {}
func (this Hx_Enum_loreline_literaltype_Array) Hx_Field_enumIndex() int32               { return 3 }
func (this Hx_Enum_loreline_literaltype_Array) Hx_Field_enumType() *Hx_Obj_go_haxe_hxenum {
	return Hx_Enum_loreline_literaltype_RTTI
}
func (this Hx_Enum_loreline_literaltype_Array) Hx_Field_enumParams() HxArray[any] {
	return HxMakeArray[any]()
}
func (this Hx_Enum_loreline_literaltype_Array) Hx_Field_enumParameter(index int32) any {
	switch index {
	default:
		return nil
	}
}

var Hx_Enum_loreline_literaltype_RTTI = Hx_Obj_go_haxe_hxenum_CreateInstance(
	"loreline.LiteralType",
	HxMakeArray[string]("Number", "Boolean", "Null", "Array", "Object"),
	HxMakeArray[int32](0, 0, 0, 0, 1),
	func(index int32, params HxArray[any]) Hx_Obj_VTable_go_haxe__hxenumvalue__hxenumvalue {
		switch index {
		case 0:
			return Hx_Enum_loreline_literaltype_Number{}
		case 1:
			return Hx_Enum_loreline_literaltype_Boolean{}
		case 2:
			return Hx_Enum_loreline_literaltype_Null{}
		case 3:
			return Hx_Enum_loreline_literaltype_Array{}
		case 4:
			return Hx_Enum_loreline_literaltype_Object{params.Get_Dyn(0).(int32)}
		default:
			return nil
		}
	},
	nil,
)

func Hx_ClassInit_Hx_Obj_VTable_loreline_naccess() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_naccess() {
}

type Hx_Obj_VTable_loreline_naccess interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_naccess struct {
	Hx_Obj_loreline_nexpr
	VTable          Hx_Obj_VTable_loreline_naccess `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target *Hx_Obj_loreline_nexpr
	Hx_Field_name   string
}

func Hx_Obj_loreline_naccess_CreateEmptyInstance() *Hx_Obj_loreline_naccess {
	obj := &Hx_Obj_loreline_naccess{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_name = HxDefault[string]()
	return obj
}

func Hx_Obj_loreline_naccess_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, name string, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_naccess {
	obj := Hx_Obj_loreline_naccess_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, name, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_naccess) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, name string, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_name = name
		}

	}
}

var Hx_Obj_loreline_naccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NAccess",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "name", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_naccess_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[string](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_naccess_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_naccess_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_naccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_naccess_RTTI
}

func (this *Hx_Obj_loreline_naccess) Hx_As_Hx_Obj_loreline_naccess() *Hx_Obj_loreline_naccess {
	return this
}

func (this *Hx_Obj_loreline_naccess) Hx_Field_type() string {
	return "Access"
}

func (this *Hx_Obj_loreline_naccess) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_target != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_target
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_target.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_naccess) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	if this.Hx_Field_target != nil {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "target", this.Hx_Field_target.VTable.Hx_Field_toJson())
	}

	Hx_Field_go_haxe_hxdynamic_setField(json, "name", this.Hx_Field_name)
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nassign() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nassign() {
}

type Hx_Obj_VTable_loreline_nassign interface {
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nassign struct {
	Hx_Obj_loreline_nexpr
	VTable          Hx_Obj_VTable_loreline_nassign `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target *Hx_Obj_loreline_nexpr
	Hx_Field_op     Hx_Enum_loreline_tokentype
	Hx_Field_value  *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nassign_CreateEmptyInstance() *Hx_Obj_loreline_nassign {
	obj := &Hx_Obj_loreline_nassign{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_op = HxDefault[Hx_Enum_loreline_tokentype]()
	obj.Hx_Field_value = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nassign_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, op Hx_Enum_loreline_tokentype, value *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nassign {
	obj := Hx_Obj_loreline_nassign_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, op, value, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nassign) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, op Hx_Enum_loreline_tokentype, value *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_op = op
			this.Hx_Field_value = value
		}

	}
}

var Hx_Obj_loreline_nassign_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NAssign",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "op", "value", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nassign_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[Hx_Enum_loreline_tokentype](params.Get(3)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)))
	},
	func() any {
		return Hx_Obj_loreline_nassign_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nassign_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nassign) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nassign_RTTI
}

func (this *Hx_Obj_loreline_nassign) Hx_As_Hx_Obj_loreline_nassign() *Hx_Obj_loreline_nassign {
	return this
}

func (this *Hx_Obj_loreline_nassign) Hx_Field_type() string {
	return "Assign"
}

func (this *Hx_Obj_loreline_nassign) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_target != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_target
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_target.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_value != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_value
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
		this.Hx_Field_value.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nassign) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "target", this.Hx_Field_target.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "op", Hx_Field_std_string(this.Hx_Field_op))
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "value", this.Hx_Field_value.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_narrayaccess() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_narrayaccess() {
}

type Hx_Obj_VTable_loreline_narrayaccess interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_narrayaccess struct {
	Hx_Obj_loreline_nexpr
	VTable          Hx_Obj_VTable_loreline_narrayaccess `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_target *Hx_Obj_loreline_nexpr
	Hx_Field_index  *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_narrayaccess_CreateEmptyInstance() *Hx_Obj_loreline_narrayaccess {
	obj := &Hx_Obj_loreline_narrayaccess{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_target = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_index = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_narrayaccess_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, index *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_narrayaccess {
	obj := Hx_Obj_loreline_narrayaccess_CreateEmptyInstance()
	obj.Hx_New(id, pos, target, index, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, target *Hx_Obj_loreline_nexpr, index *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_target = target
			this.Hx_Field_index = index
		}

	}
}

var Hx_Obj_loreline_narrayaccess_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NArrayAccess",
	HxMakeArray[string](),
	HxMakeArray[string]("target", "index", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_narrayaccess_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_narrayaccess_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_narrayaccess_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_narrayaccess_RTTI
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_As_Hx_Obj_loreline_narrayaccess() *Hx_Obj_loreline_narrayaccess {
	return this
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_Field_type() string {
	return "ArrayAccess"
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_target != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_target
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_target.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_index != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_index
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
		this.Hx_Field_index.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_narrayaccess) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "target", this.Hx_Field_target.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "index", this.Hx_Field_index.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nbinary() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nbinary() {
}

type Hx_Obj_VTable_loreline_nbinary interface {
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nbinary struct {
	Hx_Obj_loreline_nexpr
	VTable         Hx_Obj_VTable_loreline_nbinary `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_left  *Hx_Obj_loreline_nexpr
	Hx_Field_op    Hx_Enum_loreline_tokentype
	Hx_Field_right *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nbinary_CreateEmptyInstance() *Hx_Obj_loreline_nbinary {
	obj := &Hx_Obj_loreline_nbinary{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_left = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_op = HxDefault[Hx_Enum_loreline_tokentype]()
	obj.Hx_Field_right = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nbinary_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, left *Hx_Obj_loreline_nexpr, op Hx_Enum_loreline_tokentype, right *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nbinary {
	obj := Hx_Obj_loreline_nbinary_CreateEmptyInstance()
	obj.Hx_New(id, pos, left, op, right, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nbinary) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, left *Hx_Obj_loreline_nexpr, op Hx_Enum_loreline_tokentype, right *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_left = left
			this.Hx_Field_op = op
			this.Hx_Field_right = right
		}

	}
}

var Hx_Obj_loreline_nbinary_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NBinary",
	HxMakeArray[string](),
	HxMakeArray[string]("left", "op", "right", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nbinary_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[Hx_Enum_loreline_tokentype](params.Get(3)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)))
	},
	func() any {
		return Hx_Obj_loreline_nbinary_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nbinary_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nbinary) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nbinary_RTTI
}

func (this *Hx_Obj_loreline_nbinary) Hx_As_Hx_Obj_loreline_nbinary() *Hx_Obj_loreline_nbinary {
	return this
}

func (this *Hx_Obj_loreline_nbinary) Hx_Field_type() string {
	return "Binary"
}

func (this *Hx_Obj_loreline_nbinary) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_left != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_left
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_left.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_right != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_right
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
		this.Hx_Field_right.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nbinary) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "left", this.Hx_Field_left.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "op", Hx_Field_std_string(this.Hx_Field_op))
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "right", this.Hx_Field_right.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nunary() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nunary() {
}

type Hx_Obj_VTable_loreline_nunary interface {
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nunary struct {
	Hx_Obj_loreline_nexpr
	VTable           Hx_Obj_VTable_loreline_nunary `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_op      Hx_Enum_loreline_tokentype
	Hx_Field_operand *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nunary_CreateEmptyInstance() *Hx_Obj_loreline_nunary {
	obj := &Hx_Obj_loreline_nunary{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_op = HxDefault[Hx_Enum_loreline_tokentype]()
	obj.Hx_Field_operand = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nunary_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, op Hx_Enum_loreline_tokentype, operand *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nunary {
	obj := Hx_Obj_loreline_nunary_CreateEmptyInstance()
	obj.Hx_New(id, pos, op, operand, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nunary) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, op Hx_Enum_loreline_tokentype, operand *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_op = op
			this.Hx_Field_operand = operand
		}

	}
}

var Hx_Obj_loreline_nunary_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NUnary",
	HxMakeArray[string](),
	HxMakeArray[string]("op", "operand", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nunary_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[Hx_Enum_loreline_tokentype](params.Get(2)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nunary_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nunary_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nunary) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nunary_RTTI
}

func (this *Hx_Obj_loreline_nunary) Hx_As_Hx_Obj_loreline_nunary() *Hx_Obj_loreline_nunary {
	return this
}

func (this *Hx_Obj_loreline_nunary) Hx_Field_type() string {
	return "Unary"
}

func (this *Hx_Obj_loreline_nunary) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_operand != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_operand
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_operand.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nunary) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "op", Hx_Field_std_string(this.Hx_Field_op))
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "operand", this.Hx_Field_operand.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nternary() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nternary() {
}

type Hx_Obj_VTable_loreline_nternary interface {
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field_type() string
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nternary struct {
	Hx_Obj_loreline_nexpr
	VTable             Hx_Obj_VTable_loreline_nternary `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_condition *Hx_Obj_loreline_nexpr
	Hx_Field_trueExpr  *Hx_Obj_loreline_nexpr
	Hx_Field_falseExpr *Hx_Obj_loreline_nexpr
}

func Hx_Obj_loreline_nternary_CreateEmptyInstance() *Hx_Obj_loreline_nternary {
	obj := &Hx_Obj_loreline_nternary{}
	obj.Hx_Obj_loreline_nexpr = *Hx_Obj_loreline_nexpr_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_nexpr.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_condition = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_trueExpr = HxDefault[*Hx_Obj_loreline_nexpr]()
	obj.Hx_Field_falseExpr = HxDefault[*Hx_Obj_loreline_nexpr]()
	return obj
}

func Hx_Obj_loreline_nternary_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, trueExpr *Hx_Obj_loreline_nexpr, falseExpr *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nternary {
	obj := Hx_Obj_loreline_nternary_CreateEmptyInstance()
	obj.Hx_New(id, pos, condition, trueExpr, falseExpr, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nternary) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, condition *Hx_Obj_loreline_nexpr, trueExpr *Hx_Obj_loreline_nexpr, falseExpr *Hx_Obj_loreline_nexpr, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_nexpr.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_condition = condition
			this.Hx_Field_trueExpr = trueExpr
			this.Hx_Field_falseExpr = falseExpr
		}

	}
}

var Hx_Obj_loreline_nternary_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NTernary",
	HxMakeArray[string](),
	HxMakeArray[string]("condition", "trueExpr", "falseExpr", "type", "each", "toJson", "parens", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_nexpr_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nternary_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(2)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(3)), HxConvert[*Hx_Obj_loreline_nexpr](params.Get(4)), HxConvert[any](params.Get(5)), HxConvert[any](params.Get(6)))
	},
	func() any {
		return Hx_Obj_loreline_nternary_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nternary_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nternary) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nternary_RTTI
}

func (this *Hx_Obj_loreline_nternary) Hx_As_Hx_Obj_loreline_nternary() *Hx_Obj_loreline_nternary {
	return this
}

func (this *Hx_Obj_loreline_nternary) Hx_Field_type() string {
	return "Ternary"
}

func (this *Hx_Obj_loreline_nternary) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_nexpr.Hx_Field_each(handleNode)
	if this.Hx_Field_condition != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_condition
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_condition.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_trueExpr != nil {
		var _hx_tmp_1 *Hx_Obj_loreline_node
		_ = _hx_tmp_1
		_hx_tmp_1 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_trueExpr
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
		this.Hx_Field_trueExpr.VTable.Hx_Field_each(handleNode)
	}

	if this.Hx_Field_falseExpr != nil {
		var _hx_tmp_2 *Hx_Obj_loreline_node
		_ = _hx_tmp_2
		_hx_tmp_2 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_falseExpr
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_2, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_falseExpr.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nternary) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_nexpr.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "condition", this.Hx_Field_condition.VTable.Hx_Field_toJson())
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "trueExpr", this.Hx_Field_trueExpr.VTable.Hx_Field_toJson())
	var _hx_tmp_2 any
	_ = _hx_tmp_2
	_hx_tmp_2 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_2, "falseExpr", this.Hx_Field_falseExpr.VTable.Hx_Field_toJson())
	return json
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_nimportstatement() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_nimportstatement() {
}

type Hx_Obj_VTable_loreline_nimportstatement interface {
	Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node))
	Hx_Field_type() string
	Hx_Field_toJson() any
	Hx_Field_eachComment(handleComment func(node *Hx_Obj_loreline_astnode, comment *Hx_Obj_loreline_comment))
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_nimportstatement struct {
	Hx_Obj_loreline_astnode
	VTable          Hx_Obj_VTable_loreline_nimportstatement `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_path   *Hx_Obj_loreline_nstringliteral
	Hx_Field_script *Hx_Obj_loreline_script
}

func Hx_Obj_loreline_nimportstatement_CreateEmptyInstance() *Hx_Obj_loreline_nimportstatement {
	obj := &Hx_Obj_loreline_nimportstatement{}
	obj.Hx_Obj_loreline_astnode = *Hx_Obj_loreline_astnode_CreateEmptyInstance()
	obj.VTable = obj
	obj.Hx_Obj_loreline_astnode.VTable = obj
	obj.Hx_Obj_loreline_node.VTable = obj
	obj.Hx_Field_path = HxDefault[*Hx_Obj_loreline_nstringliteral]()
	obj.Hx_Field_script = HxDefault[*Hx_Obj_loreline_script]()
	return obj
}

func Hx_Obj_loreline_nimportstatement_CreateInstance(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, path *Hx_Obj_loreline_nstringliteral, script *Hx_Obj_loreline_script, leadingComments__isOptional any, trailingComments__isOptional any) *Hx_Obj_loreline_nimportstatement {
	obj := Hx_Obj_loreline_nimportstatement_CreateEmptyInstance()
	obj.Hx_New(id, pos, path, script, leadingComments__isOptional, trailingComments__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_New(id Hx_Typedef_haxe__int64___int64, pos *Hx_Obj_loreline_position, path *Hx_Obj_loreline_nstringliteral, script *Hx_Obj_loreline_script, leadingComments__isOptional any, trailingComments__isOptional any) {
	{
		leadingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = leadingComments

		if leadingComments__isOptional != nil {
			leadingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](leadingComments__isOptional)
		} else {
			leadingComments = nil
		}
		trailingComments := HxDefault[HxArray[*Hx_Obj_loreline_comment]]()
		_ = trailingComments

		if trailingComments__isOptional != nil {
			trailingComments = HxConvert[HxArray[*Hx_Obj_loreline_comment]](trailingComments__isOptional)
		} else {
			trailingComments = nil
		}
		{
			this.Hx_Obj_loreline_astnode.Hx_New(id, pos, leadingComments, trailingComments)
			this.Hx_Field_path = path
			this.Hx_Field_script = script
		}

	}
}

var Hx_Obj_loreline_nimportstatement_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.NImportStatement",
	HxMakeArray[string](),
	HxMakeArray[string]("path", "script", "each", "type", "toJson", "leadingComments", "trailingComments", "eachComment", "id", "pos"),
	Hx_Obj_loreline_astnode_RTTI,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_nimportstatement_CreateInstance(HxConvert[Hx_Typedef_haxe__int64___int64](params.Get(0)), HxConvert[*Hx_Obj_loreline_position](params.Get(1)), HxConvert[*Hx_Obj_loreline_nstringliteral](params.Get(2)), HxConvert[*Hx_Obj_loreline_script](params.Get(3)), HxConvert[any](params.Get(4)), HxConvert[any](params.Get(5)))
	},
	func() any {
		return Hx_Obj_loreline_nimportstatement_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_nimportstatement_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_nimportstatement_RTTI
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_As_Hx_Obj_loreline_nimportstatement() *Hx_Obj_loreline_nimportstatement {
	return this
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_Field_each(handleNode func(node *Hx_Obj_loreline_node, parent *Hx_Obj_loreline_node)) {
	this.Hx_Obj_loreline_astnode.Hx_Field_each(handleNode)
	if this.Hx_Field_script != nil {
		var _hx_tmp_0 *Hx_Obj_loreline_node
		_ = _hx_tmp_0
		_hx_tmp_0 = (func() *Hx_Obj_loreline_node {
			_hx_up := this.Hx_Field_script
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})()
		handleNode(_hx_tmp_0, (func() *Hx_Obj_loreline_node {
			_hx_up := this
			if _hx_up != nil {
				return &_hx_up.Hx_Obj_loreline_node
			}
			return nil
		})())
		this.Hx_Field_script.VTable.Hx_Field_each(handleNode)
	}
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_Field_type() string {
	return "Import"
}

func (this *Hx_Obj_loreline_nimportstatement) Hx_Field_toJson() any {
	var json any
	_ = json
	json = this.Hx_Obj_loreline_astnode.Hx_Field_toJson()
	var _hx_tmp_0 any
	_ = _hx_tmp_0
	_hx_tmp_0 = json
	Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_0, "path", this.Hx_Field_path.VTable.Hx_Field_toJson())
	if this.Hx_Field_script != nil {
		var _hx_tmp_1 any
		_ = _hx_tmp_1
		_hx_tmp_1 = json
		Hx_Field_go_haxe_hxdynamic_setField(_hx_tmp_1, "script", this.Hx_Field_script.VTable.Hx_Field_toJson())
	}

	return json
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_comment]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nbeatparam]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nblock]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nchoiceoption]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nexpr]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nobjectfield]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nstringpart]()
	hxRegisterArrayConstructor[*Hx_Obj_loreline_nwhenrule]()
	hxRegisterArrayConstructor[HxArray[*Hx_Obj_loreline_comment]]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[string]()
}
