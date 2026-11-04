package core

import "os"
import "runtime/pprof"
import "io"

func Hx_ClassInit_Hx_Obj_VTable_go_haxe_hxprofile() {
}

func Hx_StaticInit_Hx_Obj_VTable_go_haxe_hxprofile() {
	Hx_Field_go_haxe_hxprofile_cpuFile = Hx_Init_Hx_Field_go_haxe_hxprofile_cpuFile()
	Hx_Field_go_haxe_hxprofile_cpuPath = Hx_Init_Hx_Field_go_haxe_hxprofile_cpuPath()
	Hx_Field_go_haxe_hxprofile_memPath = Hx_Init_Hx_Field_go_haxe_hxprofile_memPath()
}

type Hx_Obj_VTable_go_haxe_hxprofile interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_go_haxe_hxprofile struct {
	VTable Hx_Obj_VTable_go_haxe_hxprofile `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_go_haxe_hxprofile_CreateEmptyInstance() *Hx_Obj_go_haxe_hxprofile {
	obj := &Hx_Obj_go_haxe_hxprofile{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_go_haxe_hxprofile_CreateInstance() *Hx_Obj_go_haxe_hxprofile {
	obj := Hx_Obj_go_haxe_hxprofile_CreateEmptyInstance()
	return obj
}

var Hx_Obj_go_haxe_hxprofile_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"go.haxe.HxProfile",
	HxMakeArray[string]("cpuFile", "cpuPath", "memPath", "start", "stop", "writeHeap"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_go_haxe_hxprofile_CreateInstance()
	},
	func() any {
		return Hx_Obj_go_haxe_hxprofile_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_go_haxe_hxprofile_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "cpuFile":
			return &Hx_Field_go_haxe_hxprofile_cpuFile
		case "cpuPath":
			return &Hx_Field_go_haxe_hxprofile_cpuPath
		case "memPath":
			return &Hx_Field_go_haxe_hxprofile_memPath
		case "start":
			return Hx_Field_go_haxe_hxprofile_start
		case "stop":
			return Hx_Field_go_haxe_hxprofile_stop
		case "writeHeap":
			return Hx_Field_go_haxe_hxprofile_writeHeap
		}
		return nil
	}
}

func (this *Hx_Obj_go_haxe_hxprofile) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_go_haxe_hxprofile_RTTI
}

func (this *Hx_Obj_go_haxe_hxprofile) Hx_As_Hx_Obj_go_haxe_hxprofile() *Hx_Obj_go_haxe_hxprofile {
	return this
}

var Hx_Field_go_haxe_hxprofile_cpuFile *os.File

func Hx_Init_Hx_Field_go_haxe_hxprofile_cpuFile() *os.File {
	return ((*os.File)(nil))
}

var Hx_Field_go_haxe_hxprofile_cpuPath string

func Hx_Init_Hx_Field_go_haxe_hxprofile_cpuPath() string {
	return "cpu.pprof"
}

var Hx_Field_go_haxe_hxprofile_memPath string

func Hx_Init_Hx_Field_go_haxe_hxprofile_memPath() string {
	return "mem.pprof"
}

func Hx_Field_go_haxe_hxprofile_start() {
}

func Hx_Field_go_haxe_hxprofile_stop() {
}

func Hx_Field_go_haxe_hxprofile_writeHeap(name string) {
	var hx_result_14 struct {
		Error  error
		Result *os.File
	}
	_ = hx_result_14
	hx_result_14.Result, hx_result_14.Error = os.Create(Hx_Field_go_haxe_hxdynamic_stringConcat(name, ".pprof"))
	var this1 struct {
		Error  error
		Result *os.File
	}
	_ = this1
	this1 = hx_result_14
	var res struct {
		Error  error
		Result *os.File
	}
	_ = res
	res = ((struct {
		Error  error
		Result *os.File
	})(this1))
	if res.Error == nil {
		pprof.WriteHeapProfile(((io.Writer)(res.Result)))
		(*res.Result).Close()
	}
}
