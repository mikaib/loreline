package core

import "math"

func Hx_ClassInit_Hx_Obj_VTable_loreline_random() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_random() {
}

type Hx_Obj_VTable_loreline_random interface {
	Hx_Field_between(min int32, max int32) int32
	Hx_Field_reset(initialSeed__isOptional any)
	Hx_Field_next() float64
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_random struct {
	VTable               Hx_Obj_VTable_loreline_random `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_seed        float64
	Hx_Field_initialSeed float64
}

func Hx_Obj_loreline_random_CreateEmptyInstance() *Hx_Obj_loreline_random {
	obj := &Hx_Obj_loreline_random{}
	obj.VTable = obj
	obj.Hx_Field_seed = HxDefault[float64]()
	obj.Hx_Field_initialSeed = HxDefault[float64]()
	return obj
}

func Hx_Obj_loreline_random_CreateInstance(seed__isOptional any) *Hx_Obj_loreline_random {
	obj := Hx_Obj_loreline_random_CreateEmptyInstance()
	obj.Hx_New(seed__isOptional)
	return obj
}

func (this *Hx_Obj_loreline_random) Hx_New(seed__isOptional any) {
	{
		seed := HxDefault[HxNullable[float64]]()
		_ = seed

		if seed__isOptional != nil {
			seed = HxConvert[HxNullable[float64]](seed__isOptional)
		} else {

		}
		{
			if (seed).Valid == false {
				var hx_nullable_t53 HxNullable[float64]
				_ = hx_nullable_t53
				hx_nullable_t53 = HxNullable[float64]{}

				seed = HxNullable[float64]{Value: (hx_nullable_t53).Value, Valid: (hx_nullable_t53).Valid}
			}

			if (seed).Valid == false {
				var now float64
				_ = now
				now = Hx_Field_sys_time()
				seed = HxNullable[float64]{Value: (now * ((float64)(1000000))), Valid: true}

				var _hx_tmp_0 float64
				_ = _hx_tmp_0
				_hx_tmp_0 = (seed).Value
				(seed).Value = (_hx_tmp_0 + (Hx_Field_math_random() * ((float64)(100000))))
				var _hx_tmp_1 float64
				_ = _hx_tmp_1
				_hx_tmp_1 = (seed).Value
				(seed).Value = (_hx_tmp_1 + (Hx_Field_sys_cpuTime() * ((float64)(10000))))
			}

			this.Hx_Field_seed = Hx_Field_loreline_random_normalizeSeed((seed).Value)
			this.Hx_Field_initialSeed = this.Hx_Field_seed
		}

	}
}

var Hx_Obj_loreline_random_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Random",
	HxMakeArray[string]("normalizeSeed"),
	HxMakeArray[string]("seed", "initialSeed", "next", "between", "reset"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_random_CreateInstance(HxConvert[any](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_random_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_random_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "normalizeSeed":
			return Hx_Field_loreline_random_normalizeSeed
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_random) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_random_RTTI
}

func (this *Hx_Obj_loreline_random) Hx_As_Hx_Obj_loreline_random() *Hx_Obj_loreline_random {
	return this
}

func (this *Hx_Obj_loreline_random) Hx_Field_next() float64 {
	this.Hx_Field_seed = math.Mod((this.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_seed
	return (((float64)(((_hx_tmp_0) / ((float64)(2147483647))))) + float64(0.000000000233))
}

func (this *Hx_Obj_loreline_random) Hx_Field_between(min int32, max int32) int32 {
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((float64)(min))
	var _hx_tmp_1 float64
	_ = _hx_tmp_1
	_hx_tmp_1 = ((float64)((max - min)))
	this.Hx_Field_seed = math.Mod((this.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
	var _hx_tmp_2 float64
	_ = _hx_tmp_2
	_hx_tmp_2 = this.Hx_Field_seed
	return Hx_Field_math_floor((_hx_tmp_0 + (_hx_tmp_1 * (((float64)(((_hx_tmp_2) / ((float64)(2147483647))))) + float64(0.000000000233)))))
}

func (this *Hx_Obj_loreline_random) Hx_Field_reset(initialSeed__isOptional any) {
	initialSeed := HxDefault[HxNullable[float64]]()
	_ = initialSeed

	if initialSeed__isOptional != nil {
		initialSeed = HxConvert[HxNullable[float64]](initialSeed__isOptional)
	} else {

	}
	{
		if (initialSeed).Valid == false {
			var hx_nullable_t52 HxNullable[float64]
			_ = hx_nullable_t52
			hx_nullable_t52 = HxNullable[float64]{}

			initialSeed = HxNullable[float64]{Value: (hx_nullable_t52).Value, Valid: (hx_nullable_t52).Valid}
		}

		if (initialSeed).Valid != false {
			this.Hx_Field_initialSeed = (initialSeed).Value
		}

		this.Hx_Field_seed = this.Hx_Field_initialSeed
	}
}

func Hx_Field_loreline_random_normalizeSeed(seed float64) float64 {
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if Hx_Field_math_isNaN(seed) {
		_hx_tmp_0 = true
	} else {
		_hx_tmp_0 = !Hx_Field_math_isFinite(seed)
	}

	if _hx_tmp_0 {
		return float64(1.0)
	}

	var result float64
	_ = result
	result = Hx_Field_math_ffloor(math.Mod(Hx_Field_math_abs(seed), float64(2147483647.0)))
	if result == ((float64)(0)) {
		return float64(1.0)
	} else {
		return result
	}
}
