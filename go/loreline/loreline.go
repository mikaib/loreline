package loreline

import (
	"fmt"
	"os"

	"github.com/jeremyfa/loreline/go/loreline/internal/core"
)

var library *core.Hx_Obj_loreline_lib_golibrary

func init() {
	core.Hx_Boot()
	library = core.Hx_Obj_loreline_lib_golibrary_CreateInstance()
}

type TextTag struct {
	Value   string
	Offset  int
	Closing bool
}

type ChoiceOption struct {
	Text    string
	Tags    []TextTag
	Enabled bool
}

type DialogueHandler func(it *Interpreter, character string, text string, tags []TextTag, advance func())

type ChoiceHandler func(it *Interpreter, options []ChoiceOption, choose func(index int))

type FinishHandler func(it *Interpreter)

type Handlers struct {
	Dialogue DialogueHandler
	Choice   ChoiceHandler
	Finish   FinishHandler
}

type Script struct {
	script *core.Hx_Obj_loreline_lib_goscript
}

type Interpreter struct {
	interp *core.Hx_Obj_loreline_lib_gointerpreter
}

func Parse(source string) (*Script, error) {
	return ParseWithImports(source, "", nil)
}

func ParseFile(path string) (*Script, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ParseWithImports(string(source), path, func(importPath string) (string, bool) {
		data, err := os.ReadFile(importPath)
		return string(data), err == nil
	})
}

func ParseWithImports(source string, path string, readFile func(path string) (string, bool)) (script *Script, err error) {
	defer recoverError(&err)

	s, err := library.Parse(source, path, func(importPath string) string {
		if readFile == nil {
			return core.HxStringNull
		}
		data, ok := readFile(importPath)
		if !ok {
			return core.HxStringNull
		}
		return data
	})
	if err != nil {
		return nil, err
	}

	return &Script{script: s}, nil
}

func (s *Script) Play(handlers Handlers, beat string) (it *Interpreter, err error) {
	defer recoverError(&err)

	it = &Interpreter{}
	dialogue, choice, finish := it.bridges(handlers)
	interp, err := s.script.Play(dialogue, choice, finish, beat)
	if err != nil {
		return nil, err
	}

	it.interp = interp
	return it, nil
}

func (s *Script) Resume(handlers Handlers, save string, beat string) (it *Interpreter, err error) {
	defer recoverError(&err)

	it = &Interpreter{}
	dialogue, choice, finish := it.bridges(handlers)
	interp, err := s.script.Resume(dialogue, choice, finish, save, beat)
	if err != nil {
		return nil, err
	}

	it.interp = interp
	return it, nil
}

func (s *Script) Format() string {
	return s.script.Format()
}

func (it *Interpreter) Start(beat string) (err error) {
	defer recoverError(&err)
	return it.interp.Start(beat)
}

func (it *Interpreter) Save() (save string, err error) {
	defer recoverError(&err)
	return it.interp.Save()
}

func (it *Interpreter) Restore(save string) (err error) {
	defer recoverError(&err)
	return it.interp.Restore(save)
}

func (it *Interpreter) StateField(name string) any {
	return fromHaxe(it.interp.StateField(name))
}

func (it *Interpreter) SetStateField(name string, value any) {
	it.interp.SetStateField(name, toHaxe(value))
}

func (it *Interpreter) CharacterField(character string, name string) any {
	return fromHaxe(it.interp.CharacterField(character, name))
}

func (it *Interpreter) SetCharacterField(character string, name string, value any) {
	it.interp.SetCharacterField(character, name, toHaxe(value))
}

func (it *Interpreter) bridges(handlers Handlers) (
	core.Hx_Typedef_loreline_lib_godialoguehandler,
	core.Hx_Typedef_loreline_lib_gochoicehandler,
	core.Hx_Typedef_loreline_lib_gofinishhandler,
) {
	dialogue := func(interp *core.Hx_Obj_loreline_lib_gointerpreter, character string, text string, tags core.HxArray[*core.Hx_Obj_loreline_texttag], advance func()) {
		it.interp = interp
		if handlers.Dialogue == nil {
			advance()
			return
		}
		handlers.Dialogue(it, character, text, convertTags(tags), advance)
	}

	choice := func(interp *core.Hx_Obj_loreline_lib_gointerpreter, options core.HxArray[*core.Hx_Obj_loreline_choiceoption], choose func(int32)) {
		it.interp = interp
		converted := make([]ChoiceOption, options.Len())
		for i := range converted {
			option := options.Get(int32(i))
			converted[i] = ChoiceOption{
				Text:    option.Hx_Field_text,
				Tags:    convertTags(option.Hx_Field_tags),
				Enabled: option.Hx_Field_enabled,
			}
		}
		if handlers.Choice == nil {
			for i, option := range converted {
				if option.Enabled {
					choose(int32(i))
					return
				}
			}
			return
		}
		handlers.Choice(it, converted, func(index int) {
			choose(int32(index))
		})
	}

	finish := func(interp *core.Hx_Obj_loreline_lib_gointerpreter) {
		it.interp = interp
		if handlers.Finish != nil {
			handlers.Finish(it)
		}
	}

	return dialogue, choice, finish
}

func convertTags(tags core.HxArray[*core.Hx_Obj_loreline_texttag]) []TextTag {
	if tags == nil {
		return nil
	}

	converted := make([]TextTag, tags.Len())
	for i := range converted {
		tag := tags.Get(int32(i))
		converted[i] = TextTag{
			Value:   tag.Hx_Field_value,
			Offset:  int(tag.Hx_Field_offset),
			Closing: tag.Hx_Field_closing,
		}
	}
	return converted
}

func recoverError(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("loreline: %v", r)
	}
}

func fromHaxe(value any) any {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		if v == core.HxStringNull {
			return nil
		}
		return v
	case int32:
		return int(v)
	case map[string]any:
		converted := make(map[string]any, len(v))
		for key, item := range v {
			converted[key] = fromHaxe(item)
		}
		return converted
	case interface{ Dyn() core.HxArray[any] }:
		items := v.Dyn()
		converted := make([]any, items.Len())
		for i := range converted {
			converted[i] = fromHaxe(items.Get(int32(i)))
		}
		return converted
	default:
		return value
	}
}

func toHaxe(value any) any {
	switch v := value.(type) {
	case nil:
		return nil
	case int:
		return int32(v)
	case int64:
		return int32(v)
	case float32:
		return float64(v)
	case []any:
		items := make([]any, len(v))
		for i, item := range v {
			items[i] = toHaxe(item)
		}
		return core.HxMakeArray[any](items...)
	case map[string]any:
		converted := make(map[string]any, len(v))
		for key, item := range v {
			converted[key] = toHaxe(item)
		}
		return converted
	default:
		return value
	}
}
