package core

func Hx_ClassInit_Hx_Obj_VTable_loreline__timer_pendingtimer() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__timer_pendingtimer() {
}

type Hx_Obj_VTable_loreline__timer_pendingtimer interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__timer_pendingtimer struct {
	VTable             Hx_Obj_VTable_loreline__timer_pendingtimer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_remaining float64
	Hx_Field_done      func()
}

func Hx_Obj_loreline__timer_pendingtimer_CreateEmptyInstance() *Hx_Obj_loreline__timer_pendingtimer {
	obj := &Hx_Obj_loreline__timer_pendingtimer{}
	obj.VTable = obj
	obj.Hx_Field_remaining = HxDefault[float64]()
	obj.Hx_Field_done = HxDefault[func()]()
	return obj
}

func Hx_Obj_loreline__timer_pendingtimer_CreateInstance(remaining float64, done func()) *Hx_Obj_loreline__timer_pendingtimer {
	obj := Hx_Obj_loreline__timer_pendingtimer_CreateEmptyInstance()
	obj.Hx_New(remaining, done)
	return obj
}

func (this *Hx_Obj_loreline__timer_pendingtimer) Hx_New(remaining float64, done func()) {
	{
		this.Hx_Field_remaining = remaining
		this.Hx_Field_done = done
	}

}

var Hx_Obj_loreline__timer_pendingtimer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Timer.PendingTimer",
	HxMakeArray[string](),
	HxMakeArray[string]("remaining", "done"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__timer_pendingtimer_CreateInstance(HxConvert[float64](params.Get(0)), HxConvert[func()](params.Get(1)))
	},
	func() any {
		return Hx_Obj_loreline__timer_pendingtimer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__timer_pendingtimer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline__timer_pendingtimer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__timer_pendingtimer_RTTI
}

func (this *Hx_Obj_loreline__timer_pendingtimer) Hx_As_Hx_Obj_loreline__timer_pendingtimer() *Hx_Obj_loreline__timer_pendingtimer {
	return this
}

func Hx_ClassInit_Hx_Obj_VTable_loreline_timer() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_timer() {
	Hx_Field_loreline_timer_timers = Hx_Init_Hx_Field_loreline_timer_timers()
	Hx_Field_loreline_timer_deferredMode = Hx_Init_Hx_Field_loreline_timer_deferredMode()
}

type Hx_Obj_VTable_loreline_timer interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_timer struct {
	VTable Hx_Obj_VTable_loreline_timer `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline_timer_CreateEmptyInstance() *Hx_Obj_loreline_timer {
	obj := &Hx_Obj_loreline_timer{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline_timer_CreateInstance() *Hx_Obj_loreline_timer {
	obj := Hx_Obj_loreline_timer_CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline_timer_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Timer",
	HxMakeArray[string]("timers", "deferredMode", "enableDeferredMode", "register", "update"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_timer_CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline_timer_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_timer_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "timers":
			return &Hx_Field_loreline_timer_timers
		case "deferredMode":
			return &Hx_Field_loreline_timer_deferredMode
		case "enableDeferredMode":
			return Hx_Field_loreline_timer_enableDeferredMode
		case "register":
			return Hx_Field_loreline_timer_register
		case "update":
			return Hx_Field_loreline_timer_update
		}
		return nil
	}
}

func (this *Hx_Obj_loreline_timer) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_timer_RTTI
}

func (this *Hx_Obj_loreline_timer) Hx_As_Hx_Obj_loreline_timer() *Hx_Obj_loreline_timer {
	return this
}

var Hx_Field_loreline_timer_timers HxArray[*Hx_Obj_loreline__timer_pendingtimer]

func Hx_Init_Hx_Field_loreline_timer_timers() HxArray[*Hx_Obj_loreline__timer_pendingtimer] {
	return HxMakeArray[*Hx_Obj_loreline__timer_pendingtimer]()
}

var Hx_Field_loreline_timer_deferredMode bool

func Hx_Init_Hx_Field_loreline_timer_deferredMode() bool {
	return false
}

func Hx_Field_loreline_timer_enableDeferredMode() {
	Hx_Field_loreline_timer_deferredMode = true
}

func Hx_Field_loreline_timer_register(seconds float64, done func()) {
	var _hx_tmp_0 HxArray[*Hx_Obj_loreline__timer_pendingtimer]
	_ = _hx_tmp_0
	_hx_tmp_0 = Hx_Field_loreline_timer_timers
	_ = (Hx_Array_Push[*Hx_Obj_loreline__timer_pendingtimer](_hx_tmp_0, Hx_Obj_loreline__timer_pendingtimer_CreateInstance(seconds, done)))
}

func Hx_Field_loreline_timer_update(delta float64) {
	Hx_Field_loreline_timer_deferredMode = true
	if Hx_Field_loreline_timer_timers.Len() == 0 {
		return
	}

	var i int32
	_ = i
	i = Hx_Field_loreline_timer_timers.Len()
	for {
		i = (i - 1)
		var _hx_tmp_0 int32
		_ = _hx_tmp_0
		_hx_tmp_0 = i
		if !(_hx_tmp_0 >= 0) {
			break
		}

		var t *Hx_Obj_loreline__timer_pendingtimer
		_ = t
		t = Hx_Field_loreline_timer_timers.Get(i)
		t.Hx_Field_remaining = (t.Hx_Field_remaining - delta)
		if t.Hx_Field_remaining <= float64(0.0) {
			_ = (Hx_Array_Splice[*Hx_Obj_loreline__timer_pendingtimer](Hx_Field_loreline_timer_timers, i, 1))
			t.Hx_Field_done()
		}
	}
}

func init() {
	hxRegisterArrayConstructor[*Hx_Obj_loreline__timer_pendingtimer]()
}
