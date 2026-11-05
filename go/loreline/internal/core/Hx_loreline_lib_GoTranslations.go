package core

func Hx_ClassInit_Hx_Obj_VTable_loreline_lib_gotranslations() {
}

func Hx_StaticInit_Hx_Obj_VTable_loreline_lib_gotranslations() {
}

type Hx_Obj_VTable_loreline_lib_gotranslations interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_loreline_lib_gotranslations struct {
	VTable                Hx_Obj_VTable_loreline_lib_gotranslations `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
	Hx_Field_translations Hx_Typedef_map
}

func Hx_Obj_loreline_lib_gotranslations_CreateEmptyInstance() *Hx_Obj_loreline_lib_gotranslations {
	obj := &Hx_Obj_loreline_lib_gotranslations{}
	obj.VTable = obj
	obj.Hx_Field_translations = HxDefault[Hx_Typedef_map]()
	return obj
}

func Hx_Obj_loreline_lib_gotranslations_CreateInstance(translations Hx_Typedef_map) *Hx_Obj_loreline_lib_gotranslations {
	obj := Hx_Obj_loreline_lib_gotranslations_CreateEmptyInstance()
	obj.Hx_New(translations)
	return obj
}

func (this *Hx_Obj_loreline_lib_gotranslations) Hx_New(translations Hx_Typedef_map) {
	{
		this.Hx_Field_translations = translations
	}

}

var Hx_Obj_loreline_lib_gotranslations_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"loreline.lib.GoTranslations",
	HxMakeArray[string](),
	HxMakeArray[string]("translations"),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_loreline_lib_gotranslations_CreateInstance(HxConvert[Hx_Typedef_map](params.Get(0)))
	},
	func() any {
		return Hx_Obj_loreline_lib_gotranslations_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_loreline_lib_gotranslations_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		return nil
	}
}

func (this *Hx_Obj_loreline_lib_gotranslations) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_loreline_lib_gotranslations_RTTI
}

func (this *Hx_Obj_loreline_lib_gotranslations) Hx_As_Hx_Obj_loreline_lib_gotranslations() *Hx_Obj_loreline_lib_gotranslations {
	return this
}
