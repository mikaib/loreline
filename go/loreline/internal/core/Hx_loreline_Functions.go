package core

import "math"
import "unicode/utf8"

func Hx_ClassInit_Hx_Obj_VTable_loreline_functions() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_functions() {
}

type Hx_Obj_VTable_loreline_functions interface {
	Hx_Field_beat_visits(name__isOptional any) int32
	Hx_Field_array_has(array any, value any) bool
	Hx_Field_string_index(text string, needle string) int32
	Hx_Field_array_index(array any, value any) int32
	Hx_Field_ceil(n float64) int32
	Hx_Field_string_ends(text string, suffix string) bool
	Hx_Field_array_join(array any, sep string) string
	Hx_Field_random(min int32, max int32) int32
	Hx_Field_seed_random(seed__isOptional any) any
	Hx_Field_string_starts(text string, prefix string) bool
	Hx_Field_abs(n float64) float64
	Hx_Field_random_float(min float64, max float64) float64
	Hx_Field_string_repeat(text string, count int32) string
	Hx_Field_rng() float64
	Hx_Field_array_remove(array any, value any) bool
	Hx_Field_max(a float64, b float64) float64
	Hx_Field_bindAll(target Hx_Typedef_map)
	Hx_Field_string_(value any) any
	Hx_Field_string_contains(text string, needle string) bool
	Hx_Field_clamp(v float64, lo float64, hi float64) float64
	Hx_Field_string_upper(text string) string
	Hx_Field_map_copy(_hx_reserved_map any) any
	Hx_Field_choices() HxArray[any]
	Hx_Field_round(n float64) int32
	Hx_Field_bool(value any) bool
	Hx_Field_string_length(text string) int32
	Hx_Field_string_replace(text string, from string, to string) string
	Hx_Field_string_trim(text string) string
	Hx_Field_plural(count any, singular string, plural_form string) string
	Hx_Field_chance(n int32) bool
	Hx_Field_array_add(array any, value any) any
	Hx_Field_array_copy(array any) any
	Hx_Field_current_beat() any
	Hx_Field_string_lower(text string) string
	Hx_Field_array_length(array any) int32
	Hx_Field_array_pop(array any) any
	Hx_Field_array_shift(array any) any
	Hx_Field_array_pick(array any) any
	Hx_Field_map_has(_hx_reserved_map any, key string) bool
	Hx_Field_map_remove(_hx_reserved_map any, key string) bool
	Hx_Field_has_beat(name any) bool
	Hx_Field_floor(n float64) int32
	Hx_Field_pow(base float64, exp float64) float64
	Hx_Field_wait(seconds float64) *Hx_Obj_loreline_async
	Hx_Field_float_(value any) any
	Hx_Field_array_prepend(array any, value any) any
	Hx_Field_choices_disabled() HxArray[any]
	Hx_Field_string_split(text string, sep string) HxArray[string]
	Hx_Field_map_get(_hx_reserved_map any, key string) any
	Hx_Field_choices_all() HxArray[any]
	Hx_Field_min(a float64, b float64) float64
	Hx_Field_string_sub(text string, start any, _hx_reserved_len any) string
	Hx_Field_array_reverse(array any) any
	Hx_Field_map_keys(_hx_reserved_map any) HxArray[string]
	Hx_Field_map_set(_hx_reserved_map any, key string, value any) any
	Hx_Field_array_sort(array any) any
	Hx_Field_array_shuffle(array any) any
	Hx_Field_map_length(_hx_reserved_map any) int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_functions struct {
	VTable               Hx_Obj_VTable_loreline_functions `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_interpreter *Hx_Obj_loreline_interpreter
}

func Hx_Obj_loreline_functions_CreateEmptyInstance() *Hx_Obj_loreline_functions {
	obj := &Hx_Obj_loreline_functions{}
	obj.VTable = obj
	obj.Hx_Field_interpreter = HxDefault[*Hx_Obj_loreline_interpreter]()
	return obj
}

func Hx_Obj_loreline_functions_CreateInstance(interpreter *Hx_Obj_loreline_interpreter) *Hx_Obj_loreline_functions {
	obj := Hx_Obj_loreline_functions_CreateEmptyInstance()
	obj.Hx_New(interpreter)
	return obj
}

func (this *Hx_Obj_loreline_functions) Hx_New(interpreter *Hx_Obj_loreline_interpreter) {
	{
		this.Hx_Field_interpreter = interpreter
	}

}

var Hx_Obj_loreline_functions_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Functions",
	HxMakeArray[string](),
	HxMakeArray[string]("interpreter", "bindAll", "rng", "floor", "ceil", "round", "abs", "min", "max", "clamp", "pow", "random", "chance", "seed_random", "random_float", "wait", "float_", "string_", "bool", "string_length", "string_upper", "string_lower", "string_contains", "string_replace", "string_split", "string_trim", "string_index", "string_sub", "string_starts", "string_ends", "string_repeat", "plural", "array_length", "array_add", "array_pop", "array_prepend", "array_shift", "array_remove", "array_index", "array_has", "array_sort", "array_reverse", "array_join", "array_pick", "array_shuffle", "array_copy", "map_length", "map_keys", "map_has", "map_get", "map_set", "map_remove", "map_copy", "current_beat", "has_beat", "beat_visits", "choices", "choices_disabled", "choices_all"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_functions_CreateInstance(HxConvert[*Hx_Obj_loreline_interpreter](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_functions_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_functions_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_functions) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_functions_RTTI
}

func (this *Hx_Obj_loreline_functions) Hx_As_Hx_Obj_loreline_functions() *Hx_Obj_loreline_functions {
	return this
}

func (this *Hx_Obj_loreline_functions) Hx_Field_bindAll(target Hx_Typedef_map) {
	target.VTable.Hx_Field_set("floor", ((any)(this.VTable.Hx_Field_floor)))
	target.VTable.Hx_Field_set("ceil", ((any)(this.VTable.Hx_Field_ceil)))
	target.VTable.Hx_Field_set("round", ((any)(this.VTable.Hx_Field_round)))
	target.VTable.Hx_Field_set("abs", ((any)(this.VTable.Hx_Field_abs)))
	target.VTable.Hx_Field_set("min", ((any)(this.VTable.Hx_Field_min)))
	target.VTable.Hx_Field_set("max", ((any)(this.VTable.Hx_Field_max)))
	target.VTable.Hx_Field_set("clamp", ((any)(this.VTable.Hx_Field_clamp)))
	target.VTable.Hx_Field_set("pow", ((any)(this.VTable.Hx_Field_pow)))
	target.VTable.Hx_Field_set("random", ((any)(this.VTable.Hx_Field_random)))
	target.VTable.Hx_Field_set("chance", ((any)(this.VTable.Hx_Field_chance)))
	target.VTable.Hx_Field_set("seed_random", ((any)(this.VTable.Hx_Field_seed_random)))
	target.VTable.Hx_Field_set("random_float", ((any)(this.VTable.Hx_Field_random_float)))
	target.VTable.Hx_Field_set("wait", ((any)(this.VTable.Hx_Field_wait)))
	target.VTable.Hx_Field_set("float", ((any)(this.VTable.Hx_Field_float_)))
	target.VTable.Hx_Field_set("string", ((any)(this.VTable.Hx_Field_string_)))
	target.VTable.Hx_Field_set("bool", ((any)(this.VTable.Hx_Field_bool)))
	target.VTable.Hx_Field_set("string_length", ((any)(this.VTable.Hx_Field_string_length)))
	target.VTable.Hx_Field_set("string_upper", ((any)(this.VTable.Hx_Field_string_upper)))
	target.VTable.Hx_Field_set("string_lower", ((any)(this.VTable.Hx_Field_string_lower)))
	target.VTable.Hx_Field_set("string_contains", ((any)(this.VTable.Hx_Field_string_contains)))
	target.VTable.Hx_Field_set("string_replace", ((any)(this.VTable.Hx_Field_string_replace)))
	target.VTable.Hx_Field_set("string_split", ((any)(this.VTable.Hx_Field_string_split)))
	target.VTable.Hx_Field_set("string_trim", ((any)(this.VTable.Hx_Field_string_trim)))
	target.VTable.Hx_Field_set("string_index", ((any)(this.VTable.Hx_Field_string_index)))
	target.VTable.Hx_Field_set("string_sub", ((any)(this.VTable.Hx_Field_string_sub)))
	target.VTable.Hx_Field_set("string_starts", ((any)(this.VTable.Hx_Field_string_starts)))
	target.VTable.Hx_Field_set("string_ends", ((any)(this.VTable.Hx_Field_string_ends)))
	target.VTable.Hx_Field_set("string_repeat", ((any)(this.VTable.Hx_Field_string_repeat)))
	target.VTable.Hx_Field_set("plural", ((any)(this.VTable.Hx_Field_plural)))
	target.VTable.Hx_Field_set("array_length", ((any)(this.VTable.Hx_Field_array_length)))
	target.VTable.Hx_Field_set("array_add", ((any)(this.VTable.Hx_Field_array_add)))
	target.VTable.Hx_Field_set("array_pop", ((any)(this.VTable.Hx_Field_array_pop)))
	target.VTable.Hx_Field_set("array_prepend", ((any)(this.VTable.Hx_Field_array_prepend)))
	target.VTable.Hx_Field_set("array_shift", ((any)(this.VTable.Hx_Field_array_shift)))
	target.VTable.Hx_Field_set("array_remove", ((any)(this.VTable.Hx_Field_array_remove)))
	target.VTable.Hx_Field_set("array_index", ((any)(this.VTable.Hx_Field_array_index)))
	target.VTable.Hx_Field_set("array_has", ((any)(this.VTable.Hx_Field_array_has)))
	target.VTable.Hx_Field_set("array_sort", ((any)(this.VTable.Hx_Field_array_sort)))
	target.VTable.Hx_Field_set("array_reverse", ((any)(this.VTable.Hx_Field_array_reverse)))
	target.VTable.Hx_Field_set("array_join", ((any)(this.VTable.Hx_Field_array_join)))
	target.VTable.Hx_Field_set("array_pick", ((any)(this.VTable.Hx_Field_array_pick)))
	target.VTable.Hx_Field_set("array_shuffle", ((any)(this.VTable.Hx_Field_array_shuffle)))
	target.VTable.Hx_Field_set("array_copy", ((any)(this.VTable.Hx_Field_array_copy)))
	target.VTable.Hx_Field_set("map_length", ((any)(this.VTable.Hx_Field_map_length)))
	target.VTable.Hx_Field_set("map_keys", ((any)(this.VTable.Hx_Field_map_keys)))
	target.VTable.Hx_Field_set("map_has", ((any)(this.VTable.Hx_Field_map_has)))
	target.VTable.Hx_Field_set("map_get", ((any)(this.VTable.Hx_Field_map_get)))
	target.VTable.Hx_Field_set("map_set", ((any)(this.VTable.Hx_Field_map_set)))
	target.VTable.Hx_Field_set("map_remove", ((any)(this.VTable.Hx_Field_map_remove)))
	target.VTable.Hx_Field_set("map_copy", ((any)(this.VTable.Hx_Field_map_copy)))
	target.VTable.Hx_Field_set("current_beat", ((any)(this.VTable.Hx_Field_current_beat)))
	target.VTable.Hx_Field_set("has_beat", ((any)(this.VTable.Hx_Field_has_beat)))
	target.VTable.Hx_Field_set("beat_visits", ((any)(this.VTable.Hx_Field_beat_visits)))
	target.VTable.Hx_Field_set("choices", ((any)(this.VTable.Hx_Field_choices)))
	target.VTable.Hx_Field_set("choices_disabled", ((any)(this.VTable.Hx_Field_choices_disabled)))
	target.VTable.Hx_Field_set("choices_all", ((any)(this.VTable.Hx_Field_choices_all)))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_rng() float64 {
	var context *Hx_Obj_loreline_interpretercontext
	_ = context
	context = this.Hx_Field_interpreter.Hx_Field_context
	if context.Hx_Field_random == nil {
		context.Hx_Field_random = Hx_Obj_loreline_random_CreateInstance(nil)
	}

	var _this *Hx_Obj_loreline_random
	_ = _this
	_this = context.Hx_Field_random
	_this.Hx_Field_seed = math.Mod((_this.Hx_Field_seed * ((float64)(16807))), ((float64)(2147483647)))
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = _this.Hx_Field_seed
	return (((float64)(((_hx_tmp_0) / ((float64)(2147483647))))) + float64(0.000000000233))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_floor(n float64) int32 {
	return Hx_Field_math_floor(n)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_ceil(n float64) int32 {
	return Hx_Field_math_ceil(n)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_round(n float64) int32 {
	return Hx_Field_math_round(n)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_abs(n float64) float64 {
	return Hx_Field_math_abs(n)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_min(a float64, b float64) float64 {
	return Hx_Field_math_min(a, b)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_max(a float64, b float64) float64 {
	return Hx_Field_math_max(a, b)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_clamp(v float64, lo float64, hi float64) float64 {
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = lo
	return Hx_Field_math_max(_hx_tmp_0, Hx_Field_math_min(hi, v))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_pow(base float64, exp float64) float64 {
	return Hx_Field_math_pow(base, exp)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_random(min int32, max int32) int32 {
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((float64)(min))
	return Hx_Field_math_floor((_hx_tmp_0 + (this.VTable.Hx_Field_rng() * ((float64)(((max + 1) - min))))))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_chance(n int32) bool {
	return (Hx_Field_math_floor((this.VTable.Hx_Field_rng() * ((float64)(n)))) == 0)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_seed_random(seed__isOptional any) any {
	seed := HxDefault[HxNullable[float64]]()
	_ = seed

	if seed__isOptional != nil {
		seed = HxConvert[HxNullable[float64]](seed__isOptional)
	} else {

	}
	{
		if (seed).Valid == false {
			var hx_nullable_t57 HxNullable[float64]
			_ = hx_nullable_t57
			hx_nullable_t57 = HxNullable[float64]{}

			seed = HxNullable[float64]{Value: (hx_nullable_t57).Value, Valid: (hx_nullable_t57).Valid}
		}

		this.Hx_Field_interpreter.VTable.Hx_Field_seedRandom(seed)
		return nil
	}
}

func (this *Hx_Obj_loreline_functions) Hx_Field_random_float(min float64, max float64) float64 {
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = min
	return (_hx_tmp_0 + (this.VTable.Hx_Field_rng() * (max - min)))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_wait(seconds float64) *Hx_Obj_loreline_async {
	return Hx_Obj_loreline_async_CreateInstance(func(done func()) {
		if Hx_Field_loreline_timer_deferredMode {
			Hx_Field_loreline_timer_register(seconds, done)
		} else {
			Hx_Field_sys_sleep(seconds)
			done()
		}
	})
}

func (this *Hx_Obj_loreline_functions) Hx_Field_float_(value any) any {
	if Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI) {
		return ((any)(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(value)))))
	}

	if Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI) {
		return ((any)((((float64)(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(value))))) * float64(1.0))))
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		return ((any)(Hx_Field_std_parseFloat(Hx_Field_go_haxe_hxdynamic_toString(value))))
	}

	if Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI) {
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(value)) {
			return ((float64)(float64(1.0)))
		} else {
			return ((float64)(float64(0.0)))
		}
	}

	return ((float64)(float64(0.0)))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_(value any) any {
	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return ((any)("null"))
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		return ((any)(Hx_Field_go_haxe_hxdynamic_toString(value)))
	}

	if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(value, nil))) {
		return ((any)("null"))
	} else {
		return ((any)(Hx_Field_std_string(value)))
	}
}

func (this *Hx_Obj_loreline_functions) Hx_Field_bool(value any) bool {
	if Hx_Field_std_isOfType(value, Hx_Abstract_bool_RTTI) {
		return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(value)))
	}

	if Hx_Field_std_isOfType(value, Hx_Obj_string_RTTI) {
		return (int32(utf8.RuneCountInString(Hx_Field_go_haxe_hxdynamic_toString(value))) > 0)
	}

	if Hx_Field_loreline_arrays_isArray(value) {
		return (Hx_Field_loreline_arrays_arrayLength(value) > 0)
	}

	if Hx_Field_std_isOfType(value, Hx_Abstract_int_RTTI) {
		return (((int32)(Hx_Field_go_haxe_hxdynamic_toInt(value))) != 0)
	}

	if Hx_Field_std_isOfType(value, Hx_Abstract_float_RTTI) {
		return (((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(value))) != ((float64)(0)))
	}

	return ((bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_nequals(value, nil))))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_length(text string) int32 {
	return int32(utf8.RuneCountInString(text))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_upper(text string) string {
	return Hx_Field_go_haxe_hxstring_toUpperCase(text)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_lower(text string) string {
	return Hx_Field_go_haxe_hxstring_toLowerCase(text)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_contains(text string, needle string) bool {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = needle
	return ((Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{})) != -1)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_replace(text string, from string, to string) string {
	return Hx_Field_stringtools_replace(text, from, to)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_split(text string, sep string) HxArray[string] {
	return (Hx_Field_go_haxe_hxstring_split(text, sep))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_trim(text string) string {
	return Hx_Field_stringtools_trim(text)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_index(text string, needle string) int32 {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	var _hx_tmp_1 string
	_ = _hx_tmp_1
	_hx_tmp_1 = needle
	return (Hx_Field_go_haxe_hxstring_indexOf(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{}))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_sub(text string, start any, _hx_reserved_len any) string {
	var _hx_tmp_0 string
	_ = _hx_tmp_0
	_hx_tmp_0 = text
	var _hx_tmp_1 int32
	_ = _hx_tmp_1
	_hx_tmp_1 = Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(start)))))
	return (Hx_Field_go_haxe_hxstring_substr(_hx_tmp_0, _hx_tmp_1, HxNullable[int32]{Value: Hx_Field_std_int(((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(Hx_Field_go_haxe_hxdynamic_toFloat(_hx_reserved_len))))), Valid: true}))
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_starts(text string, prefix string) bool {
	return Hx_Field_stringtools_startsWith(text, prefix)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_ends(text string, suffix string) bool {
	return Hx_Field_stringtools_endsWith(text, suffix)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_string_repeat(text string, count int32) string {
	var result *Hx_Obj_stringbuf
	_ = result
	result = Hx_Obj_stringbuf_CreateInstance()
	var i int32
	_ = i
	i = 0
	for i < count {
		func() func(x string) {
			var _hx_closure_src_534 func(x any)
			_ = _hx_closure_src_534
			_hx_closure_src_534 = result.VTable.Hx_Field_add
			return func(x string) {
				_hx_closure_src_534(((any)(x)))
			}
		}()(text)
		i++
	}

	return result.VTable.Hx_Field_toString()
}

func (this *Hx_Obj_loreline_functions) Hx_Field_plural(count any, singular string, plural_form string) string {
	var n float64
	_ = n
	n = ((float64)(0))
	if Hx_Field_std_isOfType(count, Hx_Abstract_int_RTTI) {
		n = (((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(((int32)(Hx_Field_go_haxe_hxdynamic_toInt(count)))))) * float64(1.0))
	} else {
		if Hx_Field_std_isOfType(count, Hx_Abstract_float_RTTI) {
			n = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(count)))
		}
	}

	if n == ((float64)(1)) {
		return singular
	} else {
		return plural_form
	}
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_length(array any) int32 {
	if Hx_Field_loreline_arrays_isArray(array) {
		return Hx_Field_loreline_arrays_arrayLength(array)
	}

	return 0
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_add(array any, value any) any {
	Hx_Field_loreline_arrays_arrayPush(array, value)
	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_pop(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		return Hx_Field_loreline_arrays_arrayPop(array)
	}

	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_prepend(array any, value any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		Hx_Field_loreline_arrays_arrayInsert(array, 0, value)
	}

	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_shift(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		return Hx_Field_loreline_arrays_arrayShift(array)
	}

	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_remove(array any, value any) bool {
	if Hx_Field_loreline_arrays_isArray(array) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_loreline_arrays_arrayGet(array, i), value))) {
					Hx_Field_loreline_arrays_arrayRemoveAt(array, i)
					return true
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_index(array any, value any) int32 {
	if Hx_Field_loreline_arrays_isArray(array) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_loreline_arrays_arrayGet(array, i), value))) {
					return i
				}
			}
		}
	}

	return -1
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_has(array any, value any) bool {
	if Hx_Field_loreline_arrays_isArray(array) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
		{
			var _g int32
			_ = _g
			_g = 0
			var _g1 int32
			_ = _g1
			_g1 = _hx_reserved_len
			for _g < _g1 {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				_g = (_g + 1)
				var i int32
				_ = i
				i = _hx_tmp_0
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(Hx_Field_loreline_arrays_arrayGet(array, i), value))) {
					return true
				}
			}
		}
	}

	return false
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_sort(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		var _hx_tmp_0 any
		_ = _hx_tmp_0
		_hx_tmp_0 = array
		Hx_Field_loreline_arrays_arraySort(_hx_tmp_0, func(a any, b any) int32 {
			var _hx_tmp_1 bool
			_ = _hx_tmp_1
			if Hx_Field_std_isOfType(a, Hx_Abstract_float_RTTI) {
				_hx_tmp_1 = Hx_Field_std_isOfType(b, Hx_Abstract_float_RTTI)
			} else {
				_hx_tmp_1 = false
			}

			if _hx_tmp_1 {
				var _hx_tmp_2 float64
				_ = _hx_tmp_2
				_hx_tmp_2 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(a)))
				if _hx_tmp_2 < ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(b))) {
					return -1
				} else {
					var _hx_tmp_3 float64
					_ = _hx_tmp_3
					_hx_tmp_3 = ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(a)))
					if _hx_tmp_3 > ((float64)(Hx_Field_go_haxe_hxdynamic_toFloat(b))) {
						return 1
					} else {
						return 0
					}
				}
			}

			var _hx_tmp_4 bool
			_ = _hx_tmp_4
			if Hx_Field_std_isOfType(a, Hx_Abstract_int_RTTI) {
				_hx_tmp_4 = Hx_Field_std_isOfType(b, Hx_Abstract_int_RTTI)
			} else {
				_hx_tmp_4 = false
			}

			if _hx_tmp_4 {
				var _hx_tmp_5 int32
				_ = _hx_tmp_5
				_hx_tmp_5 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(a)))
				if _hx_tmp_5 < ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(b))) {
					return -1
				} else {
					var _hx_tmp_6 int32
					_ = _hx_tmp_6
					_hx_tmp_6 = ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(a)))
					if _hx_tmp_6 > ((int32)(Hx_Field_go_haxe_hxdynamic_toInt(b))) {
						return 1
					} else {
						return 0
					}
				}
			}

			var _hx_tmp_7 string
			_ = _hx_tmp_7
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
				_hx_tmp_7 = "null"
			} else {
				_hx_tmp_7 = Hx_Field_std_string(a)
			}

			var _hx_tmp_8 string
			_ = _hx_tmp_8
			if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(b, nil))) {
				_hx_tmp_8 = "null"
			} else {
				_hx_tmp_8 = Hx_Field_std_string(b)
			}

			if (_hx_tmp_7) < (_hx_tmp_8) {
				return -1
			} else {
				var _hx_tmp_9 string
				_ = _hx_tmp_9
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(a, nil))) {
					_hx_tmp_9 = "null"
				} else {
					_hx_tmp_9 = Hx_Field_std_string(a)
				}

				var _hx_tmp_10 string
				_ = _hx_tmp_10
				if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(b, nil))) {
					_hx_tmp_10 = "null"
				} else {
					_hx_tmp_10 = Hx_Field_std_string(b)
				}

				if (_hx_tmp_9) > (_hx_tmp_10) {
					return 1
				} else {
					return 0
				}
			}
		})
	}

	return array
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_reverse(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		Hx_Field_loreline_arrays_arrayReverse(array)
	}

	return array
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_join(array any, sep string) string {
	if Hx_Field_loreline_arrays_isArray(array) {
		return Hx_Field_loreline_arrays_arrayJoin(array, sep)
	}

	return ""
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_pick(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		var _hx_reserved_len int32
		_ = _hx_reserved_len
		_hx_reserved_len = Hx_Field_loreline_arrays_arrayLength(array)
		if _hx_reserved_len == 0 {
			return nil
		}

		var idx int32
		_ = idx
		idx = Hx_Field_math_floor((this.VTable.Hx_Field_rng() * ((float64)(_hx_reserved_len))))
		return Hx_Field_loreline_arrays_arrayGet(array, idx)
	}

	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_shuffle(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		var i int32
		_ = i
		i = (Hx_Field_loreline_arrays_arrayLength(array) - 1)
		for i > 0 {
			var j int32
			_ = j
			j = Hx_Field_math_floor((this.VTable.Hx_Field_rng() * ((float64)((i + 1)))))
			var tmp any
			_ = tmp
			tmp = Hx_Field_loreline_arrays_arrayGet(array, i)
			var _hx_tmp_0 any
			_ = _hx_tmp_0
			_hx_tmp_0 = array
			var _hx_tmp_1 int32
			_ = _hx_tmp_1
			_hx_tmp_1 = i
			Hx_Field_loreline_arrays_arraySet(_hx_tmp_0, _hx_tmp_1, Hx_Field_loreline_arrays_arrayGet(array, j))
			Hx_Field_loreline_arrays_arraySet(array, j, tmp)
			i--
		}
	}

	return array
}

func (this *Hx_Obj_loreline_functions) Hx_Field_array_copy(array any) any {
	if Hx_Field_loreline_arrays_isArray(array) {
		return Hx_Field_loreline_arrays_arrayCopy(array)
	}

	return array
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_length(_hx_reserved_map any) int32 {
	return Hx_Field_loreline_objects_getFields(this.Hx_Field_interpreter, _hx_reserved_map).Len()
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_keys(_hx_reserved_map any) HxArray[string] {
	var fields HxArray[string]
	_ = fields
	fields = Hx_Field_loreline_objects_getFields(this.Hx_Field_interpreter, _hx_reserved_map)
	var _hx_tmp_0 HxArray[string]
	_ = _hx_tmp_0
	_hx_tmp_0 = fields
	(Hx_Array_Sort[string](_hx_tmp_0, func(a string, b string) int32 {
		if a < b {
			return -1
		} else {
			if a > b {
				return 1
			} else {
				return 0
			}
		}
	}))
	return fields
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_has(_hx_reserved_map any, key string) bool {
	return Hx_Field_loreline_objects_fieldExists(this.Hx_Field_interpreter, _hx_reserved_map, key)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_get(_hx_reserved_map any, key string) any {
	return Hx_Field_loreline_objects_getField(this.Hx_Field_interpreter, _hx_reserved_map, key)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_set(_hx_reserved_map any, key string, value any) any {
	Hx_Field_loreline_objects_setField(this.Hx_Field_interpreter, _hx_reserved_map, key, value)
	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_remove(_hx_reserved_map any, key string) bool {
	if Hx_Field_loreline_objects_fieldExists(this.Hx_Field_interpreter, _hx_reserved_map, key) {
		return Hx_Field_loreline_objects_removeField(this.Hx_Field_interpreter, _hx_reserved_map, key)
	}

	return false
}

func (this *Hx_Obj_loreline_functions) Hx_Field_map_copy(_hx_reserved_map any) any {
	var keys HxArray[string]
	_ = keys
	keys = Hx_Field_loreline_objects_getFields(this.Hx_Field_interpreter, _hx_reserved_map)
	var _hx_tmp_0 *Hx_Obj_loreline_interpreter
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_interpreter
	var _hx_tmp_1 any
	_ = _hx_tmp_1
	_hx_tmp_1 = nil
	var _hx_reserved_copy any
	_ = _hx_reserved_copy
	_hx_reserved_copy = Hx_Field_loreline_objects_createFields(_hx_tmp_0, _hx_tmp_1, nil)
	{
		var _g int32
		_ = _g
		_g = 0
		for {
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g
			if !(_hx_tmp_2 < keys.Len()) {
				break
			}

			var key string
			_ = key
			key = keys.Get(_g)
			_g++
			var _hx_tmp_3 *Hx_Obj_loreline_interpreter
			_ = _hx_tmp_3
			_hx_tmp_3 = this.Hx_Field_interpreter
			var _hx_tmp_4 any
			_ = _hx_tmp_4
			_hx_tmp_4 = _hx_reserved_copy
			var _hx_tmp_5 string
			_ = _hx_tmp_5
			_hx_tmp_5 = key
			Hx_Field_loreline_objects_setField(_hx_tmp_3, _hx_tmp_4, _hx_tmp_5, Hx_Field_loreline_objects_getField(this.Hx_Field_interpreter, _hx_reserved_map, key))
		}
	}

	return _hx_reserved_copy
}

func (this *Hx_Obj_loreline_functions) Hx_Field_current_beat() any {
	var i int32
	_ = i
	i = (this.Hx_Field_interpreter.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		var scope *Hx_Obj_loreline_runtimescope
		_ = scope
		scope = this.Hx_Field_interpreter.Hx_Field_stack.Get(i)
		if scope.Hx_Field_beat != nil {
			return ((any)(scope.Hx_Field_beat))
		}

		i--
	}

	return nil
}

func (this *Hx_Obj_loreline_functions) Hx_Field_has_beat(name any) bool {
	if Hx_Field_loreline_runtimebeatref_beatOf(name) != nil {
		return true
	}

	var nameStr string
	_ = nameStr
	nameStr = Hx_Field_go_haxe_hxdynamic_toString(name)
	var i int32
	_ = i
	i = (this.Hx_Field_interpreter.Hx_Field_stack.Len() - 1)
	for i >= 0 {
		var scope *Hx_Obj_loreline_runtimescope
		_ = scope
		scope = this.Hx_Field_interpreter.Hx_Field_stack.Get(i)
		if (scope.Hx_Field_beat != nil) && (scope.Hx_Field_beat.Hx_Field_body != nil) {
			var _g int32
			_ = _g
			_g = 0
			var _g1 HxArray[*Hx_Obj_loreline_astnode]
			_ = _g1
			_g1 = scope.Hx_Field_beat.Hx_Field_body
			for {
				var _hx_tmp_0 int32
				_ = _hx_tmp_0
				_hx_tmp_0 = _g
				if !(_hx_tmp_0 < _g1.Len()) {
					break
				}

				var node *Hx_Obj_loreline_astnode
				_ = node
				node = _g1.Get(_g)
				_g++
				if Hx_Field_std_isOfType(node, Hx_Obj_loreline_nbeatdecl_RTTI) {
					var beatDecl *Hx_Obj_loreline_nbeatdecl
					_ = beatDecl
					beatDecl = node.VTable.(interface {
						Hx_As_Hx_Obj_loreline_nbeatdecl() *Hx_Obj_loreline_nbeatdecl
					}).Hx_As_Hx_Obj_loreline_nbeatdecl()
					if beatDecl.Hx_Field_name == nameStr {
						return true
					}
				}
			}
		}

		i--
	}

	return func() func(_hx_param_key string) bool {
		var _hx_closure_src_535 func(_hx_param_key any) bool
		_ = _hx_closure_src_535
		_hx_closure_src_535 = this.Hx_Field_interpreter.Hx_Field_context.Hx_Field_topLevelBeats.VTable.(interface {
			Hx_As_Hx_Obj_haxe_ds_stringmap() *Hx_Obj_haxe_ds_stringmap
		}).Hx_As_Hx_Obj_haxe_ds_stringmap().VTable.Hx_Field_exists
		return func(_hx_param_key string) bool {
			return _hx_closure_src_535(((any)(_hx_param_key)))
		}
	}()(nameStr)
}

func (this *Hx_Obj_loreline_functions) Hx_Field_beat_visits(name__isOptional any) int32 {
	name := HxDefault[any]()
	_ = name

	if name__isOptional != nil {
		name = HxConvert[any](name__isOptional)
	} else {
		name = nil
	}
	{
		if (bool)(Hx_Field_go_haxe_hxdynamic_toBool(Hx_Field_go_haxe_hxdynamic_equals(name, nil))) {
			var i int32
			_ = i
			i = (this.Hx_Field_interpreter.Hx_Field_stack.Len() - 1)
			for i >= 0 {
				var scope *Hx_Obj_loreline_runtimescope
				_ = scope
				scope = this.Hx_Field_interpreter.Hx_Field_stack.Get(i)
				if scope.Hx_Field_beat != nil {
					return this.Hx_Field_interpreter.VTable.Hx_Field_getBeatVisitCount(scope.Hx_Field_beat)
				}

				i--
			}

			return 0
		} else {
			if Hx_Field_loreline_runtimebeatref_beatOf(name) != nil {
				return this.Hx_Field_interpreter.VTable.Hx_Field_getBeatVisitCount(Hx_Field_loreline_runtimebeatref_beatOf(name))
			} else {
				var _hx_tmp_0 string
				_ = _hx_tmp_0
				_hx_tmp_0 = Hx_Field_go_haxe_hxdynamic_toString(name)
				var beat *Hx_Obj_loreline_nbeatdecl
				_ = beat
				beat = this.Hx_Field_interpreter.VTable.Hx_Field_resolveBeatByName(_hx_tmp_0, nil)
				if beat == nil {
					return 0
				}

				return this.Hx_Field_interpreter.VTable.Hx_Field_getBeatVisitCount(beat)
			}
		}
	}
}

func (this *Hx_Obj_loreline_functions) Hx_Field_choices() HxArray[any] {
	var texts HxArray[string]
	_ = texts
	texts = this.Hx_Field_interpreter.Hx_Field__choiceEvalTexts
	var enabled HxArray[bool]
	_ = enabled
	enabled = this.Hx_Field_interpreter.Hx_Field__choiceEvalEnabled
	var result HxArray[any]
	_ = result
	result = HxMakeArray[any]().Dyn()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = texts.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			if enabled.Get(i) {
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](result), texts.Get(i)))
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_functions) Hx_Field_choices_disabled() HxArray[any] {
	var texts HxArray[string]
	_ = texts
	texts = this.Hx_Field_interpreter.Hx_Field__choiceEvalTexts
	var enabled HxArray[bool]
	_ = enabled
	enabled = this.Hx_Field_interpreter.Hx_Field__choiceEvalEnabled
	var result HxArray[any]
	_ = result
	result = HxMakeArray[any]().Dyn()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = texts.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			if !enabled.Get(i) {
				_ = (Hx_Array_Push[any](HxMakeArrayView[any](result), texts.Get(i)))
			}
		}
	}

	return result
}

func (this *Hx_Obj_loreline_functions) Hx_Field_choices_all() HxArray[any] {
	var texts HxArray[string]
	_ = texts
	texts = this.Hx_Field_interpreter.Hx_Field__choiceEvalTexts
	var result HxArray[any]
	_ = result
	result = HxMakeArray[any]().Dyn()
	{
		var _g int32
		_ = _g
		_g = 0
		var _g1 int32
		_ = _g1
		_g1 = texts.Len()
		for _g < _g1 {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g
			_g = (_g + 1)
			var i int32
			_ = i
			i = _hx_tmp_0
			_ = (Hx_Array_Push[any](HxMakeArrayView[any](result), texts.Get(i)))
		}
	}

	return result
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline_astnode]()
	hxRegisterArrayConstructor[any]()
	hxRegisterArrayConstructor[bool]()
	hxRegisterArrayConstructor[string]()
}
