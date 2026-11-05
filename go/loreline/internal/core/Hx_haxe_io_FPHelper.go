package core

import "math"

func Hx_ClassInit_Hx_Obj_VTable_haxe_io_fphelper() {
}

func Hx_StaticInit_Hx_Obj_VTable_haxe_io_fphelper() {
}

type Hx_Obj_VTable_haxe_io_fphelper interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_haxe_io_fphelper struct {
	VTable Hx_Obj_VTable_haxe_io_fphelper `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_haxe_io_fphelper_CreateEmptyInstance() *Hx_Obj_haxe_io_fphelper {
	obj := &Hx_Obj_haxe_io_fphelper{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_haxe_io_fphelper_CreateInstance() *Hx_Obj_haxe_io_fphelper {
	obj := Hx_Obj_haxe_io_fphelper_CreateEmptyInstance()
	return obj
}

var Hx_Obj_haxe_io_fphelper_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"haxe.io.FPHelper",
	HxMakeArray[string]("i32ToFloat", "floatToI32", "i64ToDouble", "doubleToI64"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_haxe_io_fphelper_CreateInstance()
	},
	func() any {
		return Hx_Obj_haxe_io_fphelper_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_haxe_io_fphelper_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "i32ToFloat":
			return Hx_Field_haxe_io_fphelper_i32ToFloat
		case "floatToI32":
			return Hx_Field_haxe_io_fphelper_floatToI32
		case "i64ToDouble":
			return Hx_Field_haxe_io_fphelper_i64ToDouble
		case "doubleToI64":
			return Hx_Field_haxe_io_fphelper_doubleToI64
		}
		return nil
	}
}

func (this *Hx_Obj_haxe_io_fphelper) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_haxe_io_fphelper_RTTI
}

func (this *Hx_Obj_haxe_io_fphelper) Hx_As_Hx_Obj_haxe_io_fphelper() *Hx_Obj_haxe_io_fphelper {
	return this
}

func Hx_Field_haxe_io_fphelper_i32ToFloat(i int32) float32 {
	return Hx_Field__single_single_impl__fromFloat32(math.Float32frombits(((uint32)(i))))
}

func Hx_Field_haxe_io_fphelper_floatToI32(f float32) int32 {
	return ((int32)(math.Float32bits(Hx_Field__single_single_impl__toFloat32(f))))
}

func Hx_Field_haxe_io_fphelper_i64ToDouble(low int32, high int32) float64 {
	var bits uint64
	_ = bits
	bits = (uint64(uint32(high)) << 32) | uint64(uint32(low))
	return math.Float64frombits(bits)
}

func Hx_Field_haxe_io_fphelper_doubleToI64(v float64) Hx_Typedef_haxe__int64___int64 {
	var bits uint64
	_ = bits
	bits = math.Float64bits(v)
	var low int32
	_ = low
	low = int32(uint32(bits))
	var high int32
	_ = high
	high = int32(uint32(bits >> 32))
	var h int64
	_ = h
	h = int64(high)
	var l int64
	_ = l
	l = int64(low)
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = h
	return ((Hx_Typedef_haxe__int64___int64)(((_hx_tmp_0 << int64(((int32)(32)))) | (l & Hx_Field_haxe__int64_int64_impl__MASK))))
}
