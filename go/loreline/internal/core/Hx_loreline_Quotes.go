package core

var Hx_Abstract_loreline_quotes_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.Quotes",
	HxMakeArray[string]("Unquoted", "DoubleQuotes", "toString"),
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

func Hx_ClassInit_Hx_Obj_VTable_loreline__quotes_quotes_impl_() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline__quotes_quotes_impl_() {
	Hx_Field_loreline__quotes_quotes_impl__Unquoted = Hx_Init_Hx_Field_loreline__quotes_quotes_impl__Unquoted()
	Hx_Field_loreline__quotes_quotes_impl__DoubleQuotes = Hx_Init_Hx_Field_loreline__quotes_quotes_impl__DoubleQuotes()
}

type Hx_Obj_VTable_loreline__quotes_quotes_impl_ interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline__quotes_quotes_impl_ struct {
	VTable Hx_Obj_VTable_loreline__quotes_quotes_impl_ `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_loreline__quotes_quotes_impl__CreateEmptyInstance() *Hx_Obj_loreline__quotes_quotes_impl_ {
	obj := &Hx_Obj_loreline__quotes_quotes_impl_{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_loreline__quotes_quotes_impl__CreateInstance() *Hx_Obj_loreline__quotes_quotes_impl_ {
	obj := Hx_Obj_loreline__quotes_quotes_impl__CreateEmptyInstance()
	return obj
}

var Hx_Obj_loreline__quotes_quotes_impl__RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline._Quotes.Quotes_Impl_",
	HxMakeArray[string]("Unquoted", "DoubleQuotes", "toString"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline__quotes_quotes_impl__CreateInstance()
	},
	func() any {
		return Hx_Obj_loreline__quotes_quotes_impl__CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline__quotes_quotes_impl__RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "Unquoted":
			return &Hx_Field_loreline__quotes_quotes_impl__Unquoted
		case "DoubleQuotes":
			return &Hx_Field_loreline__quotes_quotes_impl__DoubleQuotes
		case "toString":
			return Hx_Field_loreline__quotes_quotes_impl__toString
		}
		return nil
	}
}

func (this *Hx_Obj_loreline__quotes_quotes_impl_) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline__quotes_quotes_impl__RTTI
}

func (this *Hx_Obj_loreline__quotes_quotes_impl_) Hx_As_Hx_Obj_loreline__quotes_quotes_impl_() *Hx_Obj_loreline__quotes_quotes_impl_ {
	return this
}

var Hx_Field_loreline__quotes_quotes_impl__Unquoted int32

func Hx_Init_Hx_Field_loreline__quotes_quotes_impl__Unquoted() int32 {
	return ((int32)(0))
}

var Hx_Field_loreline__quotes_quotes_impl__DoubleQuotes int32

func Hx_Init_Hx_Field_loreline__quotes_quotes_impl__DoubleQuotes() int32 {
	return ((int32)(1))
}

func Hx_Field_loreline__quotes_quotes_impl__toString(this int32) string {
	switch this {
	case 0:
		{
			return "Unquoted"
		}

	case 1:
		{
			return "DoubleQuotes"
		}

	default:
		{
			panic(Hx_Field_haxe_exception_thrown("unreachable"))
		}
	}
}
