package core

import "reflect"

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxstringnull() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxstringnull() {
}

type Hx_Obj_VTable_go_haxe_hxstringnull interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxstringnull struct {
	VTable Hx_Obj_VTable_go_haxe_hxstringnull `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go_haxe_hxstringnull_CreateEmptyInstance() *Hx_Obj_go_haxe_hxstringnull {
	obj := &Hx_Obj_go_haxe_hxstringnull{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go_haxe_hxstringnull_CreateInstance() *Hx_Obj_go_haxe_hxstringnull {
	obj := Hx_Obj_go_haxe_hxstringnull_CreateEmptyInstance()
	return obj
}

var Hx_Obj_go_haxe_hxstringnull_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxStringNull",
	HxMakeArray[string]("isNull", "dynIsStringNull"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxstringnull_CreateInstance()
	},
	func() any {
		return Hx_Obj_go_haxe_hxstringnull_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxstringnull_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "isNull":
			return Hx_Field_go_haxe_hxstringnull_isNull
		case "dynIsStringNull":
			return Hx_Field_go_haxe_hxstringnull_dynIsStringNull
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxstringnull) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxstringnull_RTTI
}

func (this *Hx_Obj_go_haxe_hxstringnull) Hx_As_Hx_Obj_go_haxe_hxstringnull() *Hx_Obj_go_haxe_hxstringnull {
	return this
}

func Hx_Field_go_haxe_hxstringnull_isNull(s string) bool {
	return isStringNull(s)
}

func Hx_Field_go_haxe_hxstringnull_dynIsStringNull(d any) bool {
	if d == nil {
		return false
	}

	var v reflect.Value
	_ = v
	v = Hx_Field_go_haxe_hxdynamic_ensureConcreteValue(d)
	var _hx_tmp_0 bool
	_ = _hx_tmp_0
	if !v.IsValid() {
		_hx_tmp_0 = true
	} else {
		var _hx_tmp_1 reflect.Kind
		_ = _hx_tmp_1
		_hx_tmp_1 = v.Kind()
		var _hx_tmp_2 reflect.Kind
		_ = _hx_tmp_2
		_hx_tmp_2 = _hx_tmp_1
		_hx_tmp_0 = (_hx_tmp_2 != reflect.String)
	}

	if _hx_tmp_0 {
		return false
	}

	return isStringNull(v.String())
}
