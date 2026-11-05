package core

import "time"

func Hx_ClassInit_Hx_Obj_VTable_date() {
}

func Hx_StaticInit_Hx_Obj_VTable_date() {
}

type Hx_Obj_VTable_date interface {
	Hx_Field_getUTCFullYear() int32
	Hx_Field_getUTCMonth() int32
	Hx_Field_getUTCMinutes() int32
	Hx_Field_getUTCDay() int32
	Hx_Field_toString() string
	Hx_Field_getUTCSeconds() int32
	Hx_Field_getFullYear() int32
	Hx_Field_getDate() int32
	Hx_Field_getHours() int32
	Hx_Field_getMinutes() int32
	Hx_Field_getSeconds() int32
	Hx_Field_getDay() int32
	Hx_Field_getUTCDate() int32
	Hx_Field_getTime() float64
	Hx_Field_getTimezoneOffset() int32
	Hx_Field_getUTCHours() int32
	Hx_Field_getMonth() int32
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_date struct {
	VTable     Hx_Obj_VTable_date `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_t time.Time
}

func Hx_Obj_date_CreateEmptyInstance() *Hx_Obj_date {
	obj := &Hx_Obj_date{}
	obj.VTable = obj
	obj.Hx_Field_t = HxDefault[time.Time]()
	return obj
}

func Hx_Obj_date_CreateInstance(year int32, month int32, day int32, hour int32, min int32, sec int32) *Hx_Obj_date {
	obj := Hx_Obj_date_CreateEmptyInstance()
	obj.Hx_New(year, month, day, hour, min, sec)
	return obj
}

func (this *Hx_Obj_date) Hx_New(year int32, month int32, day int32, hour int32, min int32, sec int32) {
	{
		var _hx_tmp_0 int
		_ = _hx_tmp_0
		_hx_tmp_0 = int(year)
		var _hx_tmp_1 time.Month
		_ = _hx_tmp_1
		_hx_tmp_1 = ((time.Month)((month + 1)))
		var _hx_tmp_2 int
		_ = _hx_tmp_2
		_hx_tmp_2 = int(day)
		var _hx_tmp_3 int
		_ = _hx_tmp_3
		_hx_tmp_3 = int(hour)
		var _hx_tmp_4 int
		_ = _hx_tmp_4
		_hx_tmp_4 = int(min)
		var _hx_tmp_5 int
		_ = _hx_tmp_5
		_hx_tmp_5 = int(sec)
		var _hx_tmp_6 int
		_ = _hx_tmp_6
		_hx_tmp_6 = int(((int32)(0)))
		this.Hx_Field_t = time.Date(_hx_tmp_0, _hx_tmp_1, _hx_tmp_2, _hx_tmp_3, _hx_tmp_4, _hx_tmp_5, _hx_tmp_6, time.Local)
	}

}

var Hx_Obj_date_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Date",
	HxMakeArray[string]("now", "createEmpty", "fromInt", "fromTime", "fromString"),
	HxMakeArray[string]("t", "getTime", "getFullYear", "getMonth", "getDate", "getHours", "getMinutes", "getSeconds", "getDay", "getUTCFullYear", "getUTCMonth", "getUTCDate", "getUTCHours", "getUTCMinutes", "getUTCSeconds", "getUTCDay", "getTimezoneOffset", "toString"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_date_CreateInstance(HxConvert[int32](params.Get(0)), HxConvert[int32](params.Get(1)), HxConvert[int32](params.Get(2)), HxConvert[int32](params.Get(3)), HxConvert[int32](params.Get(4)), HxConvert[int32](params.Get(5)))
	},
	func() any {
		return Hx_Obj_date_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_date_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "now":
			return Hx_Field_date_now
		case "createEmpty":
			return Hx_Field_date_createEmpty
		case "fromInt":
			return Hx_Field_date_fromInt
		case "fromTime":
			return Hx_Field_date_fromTime
		case "fromString":
			return Hx_Field_date_fromString
		}
		return nil
	}
}

func (this *Hx_Obj_date) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_date_RTTI
}

func (this *Hx_Obj_date) Hx_As_Hx_Obj_date() *Hx_Obj_date {
	return this
}

func (this *Hx_Obj_date) Hx_Field_getTime() float64 {
	var t1 time.Time
	_ = t1
	t1 = this.Hx_Field_t
	var _hx_tmp_0 int
	_ = _hx_tmp_0
	_hx_tmp_0 = t1.Year()
	if _hx_tmp_0 == int(((int32)(0))) {
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = int(((int32)(1970)))
		var _hx_tmp_2 int
		_ = _hx_tmp_2
		_hx_tmp_2 = int(((int32)(0)))
		t1 = t1.AddDate(_hx_tmp_1, _hx_tmp_2, int(((int32)(0))))
	}

	return ((float64)(t1.UnixMilli()))
}

func (this *Hx_Obj_date) Hx_Field_getFullYear() int32 {
	return ((int32)(this.Hx_Field_t.Year()))
}

func (this *Hx_Obj_date) Hx_Field_getMonth() int32 {
	return (((int32)(this.Hx_Field_t.Month())) - 1)
}

func (this *Hx_Obj_date) Hx_Field_getDate() int32 {
	return ((int32)(this.Hx_Field_t.Day()))
}

func (this *Hx_Obj_date) Hx_Field_getHours() int32 {
	return ((int32)(this.Hx_Field_t.Hour()))
}

func (this *Hx_Obj_date) Hx_Field_getMinutes() int32 {
	return ((int32)(this.Hx_Field_t.Minute()))
}

func (this *Hx_Obj_date) Hx_Field_getSeconds() int32 {
	return ((int32)(this.Hx_Field_t.Second()))
}

func (this *Hx_Obj_date) Hx_Field_getDay() int32 {
	return ((int32)(this.Hx_Field_t.Weekday()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCFullYear() int32 {
	var yr int
	_ = yr
	yr = this.Hx_Field_t.UTC().Year()
	var _hx_tmp_0 int
	_ = _hx_tmp_0
	_hx_tmp_0 = yr
	if _hx_tmp_0 == int(((int32)(0))) {
		return 1970
	}

	return ((int32)(yr))
}

func (this *Hx_Obj_date) Hx_Field_getUTCMonth() int32 {
	return (((int32)(this.Hx_Field_t.UTC().Month())) - 1)
}

func (this *Hx_Obj_date) Hx_Field_getUTCDate() int32 {
	return ((int32)(this.Hx_Field_t.UTC().Day()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCHours() int32 {
	return ((int32)(this.Hx_Field_t.UTC().Hour()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCMinutes() int32 {
	return ((int32)(this.Hx_Field_t.UTC().Minute()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCSeconds() int32 {
	return ((int32)(this.Hx_Field_t.UTC().Second()))
}

func (this *Hx_Obj_date) Hx_Field_getUTCDay() int32 {
	var _hx_tmp_0 int
	_ = _hx_tmp_0
	_hx_tmp_0 = this.Hx_Field_t.UTC().Year()
	if _hx_tmp_0 == int(((int32)(0))) {
		var _hx_tmp_1 int
		_ = _hx_tmp_1
		_hx_tmp_1 = int(((int32)(1970)))
		var _hx_tmp_2 int
		_ = _hx_tmp_2
		_hx_tmp_2 = int(((int32)(0)))
		return ((int32)(this.Hx_Field_t.AddDate(_hx_tmp_1, _hx_tmp_2, int(((int32)(0)))).Weekday()))
	}

	return ((int32)(this.Hx_Field_t.UTC().Weekday()))
}

func (this *Hx_Obj_date) Hx_Field_getTimezoneOffset() int32 {
	var hx_tuple_8 struct {
		Name   string
		Offset int
	}
	_ = hx_tuple_8
	hx_tuple_8.Name, hx_tuple_8.Offset = this.Hx_Field_t.Zone()
	return ((int32)(hx_tuple_8.Offset))
}

func (this *Hx_Obj_date) Hx_Field_toString() string {
	return this.Hx_Field_t.Format("2006-01-02 15:04:05")
}

func Hx_Field_date_now() *Hx_Obj_date {
	var d *Hx_Obj_date
	_ = d
	d = Hx_Field_date_createEmpty()
	d.Hx_Field_t = time.Now()
	return d
}

func Hx_Field_date_createEmpty() *Hx_Obj_date {
	return Hx_Obj_date_CreateInstance(0, 0, 0, 0, 0, 0)
}

func Hx_Field_date_fromInt(t int32) *Hx_Obj_date {
	var d *Hx_Obj_date
	_ = d
	d = Hx_Field_date_createEmpty()
	var _hx_tmp_0 int64
	_ = _hx_tmp_0
	_hx_tmp_0 = int64(t)
	d.Hx_Field_t = time.Unix(_hx_tmp_0, int64(((int32)(0))))
	return d
}

func Hx_Field_date_fromTime(t float64) *Hx_Obj_date {
	var d *Hx_Obj_date
	_ = d
	d = Hx_Field_date_createEmpty()
	d.Hx_Field_t = time.UnixMilli(((int64)(Hx_Field_haxe_int64helper_fromFloat(t))))
	return d
}

func Hx_Field_date_fromString(s string) *Hx_Obj_date {
	var d *Hx_Obj_date
	_ = d
	d = Hx_Field_date_createEmpty()
	var hx_result_9 struct {
		Error  error
		Result time.Time
	}
	_ = hx_result_9
	hx_result_9.Result, hx_result_9.Error = time.Parse(time.DateTime, s)
	var res struct {
		Error  error
		Result time.Time
	}
	_ = res
	res = hx_result_9
	var tmp bool
	_ = tmp
	if res.Error != nil {
		tmp = false
	} else {
		tmp = true
	}

	if tmp {
		var tmp_tmp_1 time.Time
		_ = tmp_tmp_1
		if res.Error != nil {
			var e error
			_ = e
			e = res.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r time.Time
			_ = r
			r = res.Result
			tmp_tmp_1 = r
		}

		d.Hx_Field_t = tmp_tmp_1
		return d
	}

	var hx_result_10 struct {
		Error  error
		Result time.Time
	}
	_ = hx_result_10
	hx_result_10.Result, hx_result_10.Error = time.Parse(time.DateOnly, s)
	res = hx_result_10
	var tmp1 bool
	_ = tmp1
	if res.Error != nil {
		tmp1 = false
	} else {
		tmp1 = true
	}

	if tmp1 {
		var tmp_tmp_1 time.Time
		_ = tmp_tmp_1
		if res.Error != nil {
			var e error
			_ = e
			e = res.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r time.Time
			_ = r
			r = res.Result
			tmp_tmp_1 = r
		}

		d.Hx_Field_t = tmp_tmp_1
		return d
	}

	var hx_result_11 struct {
		Error  error
		Result time.Time
	}
	_ = hx_result_11
	hx_result_11.Result, hx_result_11.Error = time.Parse(time.TimeOnly, s)
	res = hx_result_11
	var tmp2 bool
	_ = tmp2
	if res.Error != nil {
		tmp2 = false
	} else {
		tmp2 = true
	}

	if tmp2 {
		var tmp_tmp_1 time.Time
		_ = tmp_tmp_1
		if res.Error != nil {
			var e error
			_ = e
			e = res.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r time.Time
			_ = r
			r = res.Result
			tmp_tmp_1 = r
		}

		d.Hx_Field_t = tmp_tmp_1
		return d
	}

	panic(Hx_Field_haxe_exception_thrown(((any)(Hx_Field_go_haxe_hxstring_concat("Invalid date format : ", s)))))
}
