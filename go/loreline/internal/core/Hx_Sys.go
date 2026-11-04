package core

import "os"
import "time"
import "strings"
import "runtime"
import "os/exec"
import "golang.org/x/term"

func Hx_ClassInit_Hx_Obj_VTable_sys() {
}

func Hx_StaticInit_Hx_Obj_VTable_sys() {
}

type Hx_Obj_VTable_sys interface {
	Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass
}

type Hx_Obj_sys struct {
	VTable Hx_Obj_VTable_sys `json:"-" xml:"-" yaml:"-" gorm:"-:all"`
}

func Hx_Obj_sys_CreateEmptyInstance() *Hx_Obj_sys {
	obj := &Hx_Obj_sys{}
	obj.VTable = obj
	return obj
}

func Hx_Obj_sys_CreateInstance() *Hx_Obj_sys {
	obj := Hx_Obj_sys_CreateEmptyInstance()
	return obj
}

var Hx_Obj_sys_RTTI = Hx_Obj_go_haxe_hxclass_CreateInstance(
	"Sys",
	HxMakeArray[string]("print", "println", "args", "getEnv", "putEnv", "sleep", "setTimeLocale", "getCwd", "setCwd", "systemName", "command", "exit", "time", "cpuTime", "executablePath", "programPath", "environment", "stdin", "stdout", "stderr", "getChar"),
	HxMakeArray[string](),
	nil,
	HxMakeArray[*Hx_Obj_go_haxe_hxclass](),
	func(params HxArray[any]) any {
		return Hx_Obj_sys_CreateInstance()
	},
	func() any {
		return Hx_Obj_sys_CreateEmptyInstance()
	},
	nil,
	nil,
)

func init() {
	Hx_Obj_sys_RTTI.Hx_Field_getStaticFieldPtr = func(name string) any {
		switch name {
		case "print":
			return Hx_Field_sys_print
		case "println":
			return Hx_Field_sys_println
		case "args":
			return Hx_Field_sys_args
		case "getEnv":
			return Hx_Field_sys_getEnv
		case "putEnv":
			return Hx_Field_sys_putEnv
		case "sleep":
			return Hx_Field_sys_sleep
		case "setTimeLocale":
			return Hx_Field_sys_setTimeLocale
		case "getCwd":
			return Hx_Field_sys_getCwd
		case "setCwd":
			return Hx_Field_sys_setCwd
		case "systemName":
			return Hx_Field_sys_systemName
		case "command":
			return Hx_Field_sys_command
		case "exit":
			return Hx_Field_sys_exit
		case "time":
			return Hx_Field_sys_time
		case "cpuTime":
			return Hx_Field_sys_cpuTime
		case "executablePath":
			return Hx_Field_sys_executablePath
		case "programPath":
			return Hx_Field_sys_programPath
		case "environment":
			return Hx_Field_sys_environment
		case "stdin":
			return Hx_Field_sys_stdin
		case "stdout":
			return Hx_Field_sys_stdout
		case "stderr":
			return Hx_Field_sys_stderr
		case "getChar":
			return Hx_Field_sys_getChar
		}
		return nil
	}
}

func (this *Hx_Obj_sys) Hx_Field__RTTI() *Hx_Obj_go_haxe_hxclass {
	return Hx_Obj_sys_RTTI
}

func (this *Hx_Obj_sys) Hx_As_Hx_Obj_sys() *Hx_Obj_sys {
	return this
}

func Hx_Field_sys_print(v any) {
	{
		var hx_result_25 struct {
			Error  error
			Result int
		}
		_ = hx_result_25
		hx_result_25.Result, hx_result_25.Error = (*os.Stdout).WriteString(Hx_Field_std_string(v))
		_ = hx_result_25
	}
}

func Hx_Field_sys_println(v any) {
	{
		var hx_result_26 struct {
			Error  error
			Result int
		}
		_ = hx_result_26
		hx_result_26.Result, hx_result_26.Error = (*os.Stdout).WriteString(Hx_Field_go_haxe_hxdynamic_stringConcat(Hx_Field_std_string(v), "\n"))
		_ = hx_result_26
	}
}

func Hx_Field_sys_args() HxArray[string] {
	var self []string
	_ = self
	self = os.Args
	var args HxArray[string]
	_ = args
	args = HxMakeArray(self...)
	_ = (Hx_Array_Shift[string](args))
	return args
}

func Hx_Field_sys_getEnv(s string) string {
	var hx_tuple_10 struct {
		P0 string
		P1 bool
	}
	_ = hx_tuple_10
	hx_tuple_10.P0, hx_tuple_10.P1 = os.LookupEnv(s)
	var tup struct {
		P0 string
		P1 bool
	}
	_ = tup
	tup = ((struct {
		P0 string
		P1 bool
	})(hx_tuple_10))
	if !tup.P1 {
		return HxStringNull
	}

	return tup.P0
}

func Hx_Field_sys_putEnv(s string, v HxNullable[string]) {
	if (v).Valid == false {
		var this1 error
		_ = this1
		this1 = os.Unsetenv(s)
		if this1 != nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(this1.Error()))))
		}
	} else {
		var _hx_tmp_0 string
		_ = _hx_tmp_0
		_hx_tmp_0 = s
		var this1 error
		_ = this1
		this1 = os.Setenv(_hx_tmp_0, (v).Value)
		if this1 != nil {
			panic(Hx_Field_haxe_exception_thrown(((any)(this1.Error()))))
		}
	}
}

func Hx_Field_sys_sleep(seconds float64) {
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = (seconds * ((float64)(1000)))
	time.Sleep(((time.Duration)((_hx_tmp_0 * ((float64)(time.Millisecond))))))
}

func Hx_Field_sys_setTimeLocale(loc string) bool {
	var hx_result_27 struct {
		Error  error
		Result *time.Location
	}
	_ = hx_result_27
	hx_result_27.Result, hx_result_27.Error = time.LoadLocation(loc)
	var locale struct {
		Error  error
		Result *time.Location
	}
	_ = locale
	locale = hx_result_27
	var tmp bool
	_ = tmp
	if locale.Error != nil {
		tmp = false
	} else {
		tmp = true
	}

	if !tmp {
		return false
	}

	var tmp1 *time.Location
	_ = tmp1
	if locale.Error != nil {
		var e error
		_ = e
		e = locale.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r *time.Location
		_ = r
		r = locale.Result
		tmp1 = r
	}

	time.Local = tmp1
	return true
}

func Hx_Field_sys_getCwd() string {
	var hx_result_28 struct {
		Error  error
		Result string
	}
	_ = hx_result_28
	hx_result_28.Result, hx_result_28.Error = os.Getwd()
	var this1 struct {
		Error  error
		Result string
	}
	_ = this1
	this1 = hx_result_28
	var ret string
	_ = ret
	if this1.Error != nil {
		var e error
		_ = e
		e = this1.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r string
		_ = r
		r = this1.Result
		ret = r
	}

	var sep rune
	_ = sep
	sep = os.PathSeparator
	var trailingSlash string
	_ = trailingSlash
	trailingSlash = string(sep)
	if strings.HasSuffix(ret, trailingSlash) {
		return ret
	}

	return Hx_Field_go_haxe_hxdynamic_stringConcat(ret, trailingSlash)
}

func Hx_Field_sys_setCwd(s string) {
	var this1 error
	_ = this1
	this1 = os.Chdir(s)
	if this1 != nil {
		panic(Hx_Field_haxe_exception_thrown(((any)(this1.Error()))))
	}
}

func Hx_Field_sys_systemName() string {
	switch runtime.GOOS {
	case "android", "linux":
		{
			return "Linux"
		}

	case "darwin":
		{
			return "Mac"
		}

	case "freebsd", "netbsd", "openbsd":
		{
			return "BSD"
		}

	case "windows":
		{
			return "Windows"
		}

	default:
		{
			return runtime.GOOS
		}
	}
}

func Hx_Field_sys_command(cmd string, args__isOptional any) int32 {
	args := HxDefault[HxArray[string]]()
	_ = args

	if args__isOptional != nil {
		args = HxConvert[HxArray[string]](args__isOptional)
	} else {
		args = nil
	}
	{
		var _hx_tmp_0 *exec.Cmd
		_ = _hx_tmp_0
		if args == nil {
			var _hx_tmp_1 *exec.Cmd
			_ = _hx_tmp_1
			if runtime.GOOS == "windows" {
				_hx_tmp_1 = exec.Command("cmd", "/C", cmd)
			} else {
				_hx_tmp_1 = exec.Command("/bin/sh", "-c", cmd)
			}

			_hx_tmp_0 = _hx_tmp_1
		} else {
			var _hx_tmp_2 string
			_ = _hx_tmp_2
			_hx_tmp_2 = cmd
			var _hx_tmp_3 string
			_ = _hx_tmp_3
			_hx_tmp_3 = _hx_tmp_2
			_hx_tmp_0 = exec.Command(_hx_tmp_3, args.Underlying()...)
		}

		var proc *exec.Cmd
		_ = proc
		proc = _hx_tmp_0
		var hx_result_29 struct {
			Error  error
			Result []byte
		}
		_ = hx_result_29
		hx_result_29.Result, hx_result_29.Error = (*proc).Output()
		var res struct {
			Error  error
			Result []byte
		}
		_ = res
		res = hx_result_29
		var t struct {
			Error  error
			Result []byte
		}
		_ = t
		t = ((struct {
			Error  error
			Result []byte
		})(((struct {
			Error  error
			Result []byte
		})(res))))
		var output []byte
		_ = output
		output = t.Result
		var err error
		_ = err
		err = t.Error
		var exitCode int
		_ = exitCode
		exitCode = int(((int32)(0)))
		var exitError *exec.ExitError
		_ = exitError
		var ok bool
		_ = ok
		ok = false
		if err != nil {
			exitError, ok = err.(*exec.ExitError)
			if ok {
				exitCode = exitError.ProcessState.ExitCode()
			} else {
				exitCode = int(((int32)(-1)))
			}
		} else {
			Hx_Field_sys_println(string(output))
		}

		return ((int32)(exitCode))
	}
}

func Hx_Field_sys_exit(code int32) {
	Hx_Field_go_haxe_hxprofile_stop()
	os.Exit(int(code))
}

func Hx_Field_sys_time() float64 {
	var tn time.Time
	_ = tn
	tn = time.Now()
	var _hx_tmp_0 float64
	_ = _hx_tmp_0
	_hx_tmp_0 = ((float64)(tn.UnixNano()))
	var sec float64
	_ = sec
	sec = ((float64)((_hx_tmp_0 / ((float64)(time.Second)))))
	var hx_tuple_11 struct {
		Name   string
		Offset int
	}
	_ = hx_tuple_11
	hx_tuple_11.Name, hx_tuple_11.Offset = tn.Local().Zone()
	var off int
	_ = off
	off = hx_tuple_11.Offset
	return (sec + ((float64)(off)))
}

func Hx_Field_sys_cpuTime() float64 {
	return Hx_Field_sys_time()
}

func Hx_Field_sys_executablePath() string {
	return Hx_Field_sys_programPath()
}

func Hx_Field_sys_programPath() string {
	return os.Args[0]
}

func Hx_Field_sys_environment() Hx_Typedef_map {
	var em Hx_Typedef_map
	_ = em
	em = (func() *Hx_Obj_haxe_imap {
		_hx_icast_90 := Hx_Obj_haxe_ds_stringmap_CreateInstance()
		if _hx_icast_90 == nil {
			return nil
		}
		Hx_Field_go_haxe_hxdynamic_castClass(_hx_icast_90, "Hx_Obj_haxe_imap")
		return &Hx_Obj_haxe_imap{VTable: _hx_icast_90.VTable.(Hx_Obj_VTable_haxe_imap)}
	})()
	var ev []string
	_ = ev
	ev = os.Environ()
	{
		var self []string
		_ = self
		self = ev
		var _g_current int32
		_ = _g_current
		_g_current = 0
		var _g_array HxArray[string]
		_ = _g_array
		_g_array = HxMakeArray(self...)
		for {
			var _hx_tmp_0 int32
			_ = _hx_tmp_0
			_hx_tmp_0 = _g_current
			if !(_hx_tmp_0 < _g_array.Len()) {
				break
			}

			var _hx_tmp_1 HxArray[string]
			_ = _hx_tmp_1
			_hx_tmp_1 = _g_array
			var _hx_tmp_2 int32
			_ = _hx_tmp_2
			_hx_tmp_2 = _g_current
			_g_current = (_g_current + 1)
			var entry string
			_ = entry
			entry = _hx_tmp_1.Get(_hx_tmp_2)
			var parts HxArray[string]
			_ = parts
			parts = (Hx_Field_go_haxe_hxstring_split(entry, "="))
			var name string
			_ = name
			name = (Hx_Array_Shift[string](parts)).Value
			var value string
			_ = value
			value = (Hx_Array_Join[string](parts, "="))
			func() func(_hx_param_key string, _hx_param_value string) {
				var _hx_closure_src_437 func(_hx_param_key any, _hx_param_value any)
				_ = _hx_closure_src_437
				_hx_closure_src_437 = em.VTable.Hx_Field_set
				return func(_hx_param_key string, _hx_param_value string) {
					_hx_closure_src_437(((any)(_hx_param_key)), ((any)(_hx_param_value)))
				}
			}()(name, value)
		}
	}

	return em
}

func Hx_Field_sys_stdin() *Hx_Obj_haxe_io_input {
	return (func() *Hx_Obj_haxe_io_input {
		_hx_up := Hx_Obj_sys_io_fileinput_CreateInstance(os.Stdin)
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_haxe_io_input
		}
		return nil
	})()
}

func Hx_Field_sys_stdout() *Hx_Obj_haxe_io_output {
	return (func() *Hx_Obj_haxe_io_output {
		_hx_up := Hx_Obj_sys_io_fileoutput_CreateInstance(os.Stdout)
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_haxe_io_output
		}
		return nil
	})()
}

func Hx_Field_sys_stderr() *Hx_Obj_haxe_io_output {
	return (func() *Hx_Obj_haxe_io_output {
		_hx_up := Hx_Obj_sys_io_fileoutput_CreateInstance(os.Stderr)
		if _hx_up != nil {
			return &_hx_up.Hx_Obj_haxe_io_output
		}
		return nil
	})()
}

func Hx_Field_sys_getChar(echo bool) int32 {
	var hx_result_30 struct {
		Error  error
		Result *term.State
	}
	_ = hx_result_30
	hx_result_30.Result, hx_result_30.Error = term.MakeRaw(int((*os.Stdin).Fd()))
	var this1 struct {
		Error  error
		Result *term.State
	}
	_ = this1
	this1 = hx_result_30
	var oldState *term.State
	_ = oldState
	if this1.Error != nil {
		var e error
		_ = e
		e = this1.Error
		panic(Hx_Field_haxe_exception_thrown(((any)(e))))
	} else {
		var r *term.State
		_ = r
		r = this1.Result
		oldState = r
	}

	defer func() {
		term.Restore(int((*os.Stdin).Fd()), oldState)
	}()
	var length HxNullable[int32]
	_ = length
	length = HxNullable[int32]{Value: 1, Valid: true}

	var _hx_tmp_0 int32
	_ = _hx_tmp_0
	if (length).Valid != false {
		_hx_tmp_0 = (length).Value
	} else {
		_hx_tmp_0 = 0
	}

	var b []byte
	_ = b
	b = make([]byte, _hx_tmp_0)
	{
		var hx_result_31 struct {
			Error  error
			Result int
		}
		_ = hx_result_31
		hx_result_31.Result, hx_result_31.Error = (*os.Stdin).Read(b)
		var this2 struct {
			Error  error
			Result int
		}
		_ = this2
		this2 = hx_result_31
		if this2.Error != nil {
			var e error
			_ = e
			e = this2.Error
			panic(Hx_Field_haxe_exception_thrown(((any)(e))))
		} else {
			var r int
			_ = r
			r = this2.Result
		}
	}

	if echo {
		Hx_Field_sys_println(Hx_Field_go_haxe_hxstring_fromCharCode(((int32)(b[0]))))
	}

	return ((int32)(b[0]))
}

func init() {
	hxRegisterArrayConstructor[string]()
}
