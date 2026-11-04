package core

import "math"
import "math/rand"

func Hx_ClassInit_Hx_Obj_VTable_math() {
}

func Hx_StaticInit_Hx_Obj_VTable_math() {
	Hx_Field_math_PI = Hx_Init_Hx_Field_math_PI()
	Hx_Field_math_NEGATIVE_INFINITY = Hx_Init_Hx_Field_math_NEGATIVE_INFINITY()
	Hx_Field_math_POSITIVE_INFINITY = Hx_Init_Hx_Field_math_POSITIVE_INFINITY()
	Hx_Field_math_NaN = Hx_Init_Hx_Field_math_NaN()
}

type Hx_Obj_VTable_math interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_math struct {
	VTable Hx_Obj_VTable_math `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_math_CreateEmptyInstance() *Hx_Obj_math {
	obj := &Hx_Obj_math{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_math_CreateInstance() *Hx_Obj_math {
	obj := Hx_Obj_math_CreateEmptyInstance()
	return obj
}

var Hx_Obj_math_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Math",
	HxMakeArray[string]("PI", "NEGATIVE_INFINITY", "POSITIVE_INFINITY", "NaN", "abs", "min", "max", "sin", "cos", "tan", "asin", "acos", "atan", "atan2", "exp", "log", "pow", "sqrt", "isNaN", "ffloor", "fceil", "fround", "round", "floor", "ceil", "random", "isFinite"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_math_CreateInstance()
	},
	func() any {
		return Hx_Obj_math_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_math_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "PI":
			return &Hx_Field_math_PI
		case "NEGATIVE_INFINITY":
			return &Hx_Field_math_NEGATIVE_INFINITY
		case "POSITIVE_INFINITY":
			return &Hx_Field_math_POSITIVE_INFINITY
		case "NaN":
			return &Hx_Field_math_NaN
		case "abs":
			return Hx_Field_math_abs
		case "min":
			return Hx_Field_math_min
		case "max":
			return Hx_Field_math_max
		case "sin":
			return Hx_Field_math_sin
		case "cos":
			return Hx_Field_math_cos
		case "tan":
			return Hx_Field_math_tan
		case "asin":
			return Hx_Field_math_asin
		case "acos":
			return Hx_Field_math_acos
		case "atan":
			return Hx_Field_math_atan
		case "atan2":
			return Hx_Field_math_atan2
		case "exp":
			return Hx_Field_math_exp
		case "log":
			return Hx_Field_math_log
		case "pow":
			return Hx_Field_math_pow
		case "sqrt":
			return Hx_Field_math_sqrt
		case "isNaN":
			return Hx_Field_math_isNaN
		case "ffloor":
			return Hx_Field_math_ffloor
		case "fceil":
			return Hx_Field_math_fceil
		case "fround":
			return Hx_Field_math_fround
		case "round":
			return Hx_Field_math_round
		case "floor":
			return Hx_Field_math_floor
		case "ceil":
			return Hx_Field_math_ceil
		case "random":
			return Hx_Field_math_random
		case "isFinite":
			return Hx_Field_math_isFinite
		}
		return nil
	}
}

func (this *Hx_Obj_math) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_math_RTTI
}

func (this *Hx_Obj_math) Hx_As_Hx_Obj_math() *Hx_Obj_math {
	return this
}

var Hx_Field_math_PI float64

func Hx_Init_Hx_Field_math_PI() float64 {
	return math.Pi
}

var Hx_Field_math_NEGATIVE_INFINITY float64

func Hx_Init_Hx_Field_math_NEGATIVE_INFINITY() float64 {
	return math.Inf(int(((int32)(-1))))
}

var Hx_Field_math_POSITIVE_INFINITY float64

func Hx_Init_Hx_Field_math_POSITIVE_INFINITY() float64 {
	return math.Inf(int(((int32)(1))))
}

var Hx_Field_math_NaN float64

func Hx_Init_Hx_Field_math_NaN() float64 {
	return math.NaN()
}

func Hx_Field_math_abs(v float64) float64 {
	return math.Abs(v)
}

func Hx_Field_math_min(a float64, b float64) float64 {
	if Hx_Field_math_isNaN(a) {
		return a
	}

	if Hx_Field_math_isNaN(b) {
		return b
	}

	return math.Min(a, b)
}

func Hx_Field_math_max(a float64, b float64) float64 {
	if Hx_Field_math_isNaN(a) {
		return a
	}

	if Hx_Field_math_isNaN(b) {
		return b
	}

	return math.Max(a, b)
}

func Hx_Field_math_sin(v float64) float64 {
	return math.Sin(v)
}

func Hx_Field_math_cos(v float64) float64 {
	return math.Cos(v)
}

func Hx_Field_math_tan(v float64) float64 {
	return math.Tan(v)
}

func Hx_Field_math_asin(v float64) float64 {
	return math.Asin(v)
}

func Hx_Field_math_acos(v float64) float64 {
	return math.Acos(v)
}

func Hx_Field_math_atan(v float64) float64 {
	return math.Atan(v)
}

func Hx_Field_math_atan2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}

func Hx_Field_math_exp(v float64) float64 {
	return math.Exp(v)
}

func Hx_Field_math_log(v float64) float64 {
	return math.Log(v)
}

func Hx_Field_math_pow(v float64, exp float64) float64 {
	return math.Pow(v, exp)
}

func Hx_Field_math_sqrt(v float64) float64 {
	return math.Sqrt(v)
}

func Hx_Field_math_isNaN(f float64) bool {
	return math.IsNaN(f)
}

func Hx_Field_math_ffloor(v float64) float64 {
	return math.Floor(v)
}

func Hx_Field_math_fceil(v float64) float64 {
	return math.Ceil(v)
}

func Hx_Field_math_fround(v float64) float64 {
	if v < float64(0.0) {
		v = (v + float64(0.000001))
	}

	return math.Round(v)
}

func Hx_Field_math_round(v float64) int32 {
	return ((int32)(int32(math.Floor((v + float64(0.5))))))
}

func Hx_Field_math_floor(v float64) int32 {
	return ((int32)(int32(Hx_Field_math_ffloor(v))))
}

func Hx_Field_math_ceil(v float64) int32 {
	return ((int32)(int32(Hx_Field_math_fceil(v))))
}

func Hx_Field_math_random() float64 {
	return rand.Float64()
}

func Hx_Field_math_isFinite(f float64) bool {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if !math.IsNaN(f) {
		var _hx_tmp_1 float64
		_ = _hx_tmp_1
		_hx_tmp_1 = f
		var _hx_tmp_2 float64
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = !math.IsInf(_hx_tmp_2, int(((int32)(0))))
	} else {
		_hx_tmp_0 = false
	}

	return _hx_tmp_0
}
