package ctrlutils

import (
	"fmt"
	"io"
	"text/template"

	"github.com/JohnBrainard/droputils/config"
)

const ledInstructionTemplate = `
led_instruction_t led_instructions[] = {
	{{range $instructionIndex, $instruction := .Instructions}}
	{
		.flags = {{$instruction.Flags}},
		{{- if $instruction.MatchLayers }}
		.layer = {{index $instruction.MatchLayers 0}},
		{{end}}

		{{- range $idIndex, $id := $instruction.IDs }}

		// {{FormatIDBinary $id}}
		.id{{$idIndex}} = {{FormatID $id}},
		{{- end }}
		.r = {{$instruction.Red}},
		.g = {{$instruction.Green}},
		.b = {{$instruction.Blue}},
	},
	{{end}}

	{.end = 1}
};`

type KeymapRenderer struct {
	keyMap *KeyMap
	cfg    *config.Config
}

func NewKeymapRenderer(keyMap *KeyMap, cfg *config.Config) KeymapRenderer {
	return KeymapRenderer{
		keyMap: keyMap,
		cfg:    cfg,
	}
}

func (r KeymapRenderer) Render(writer io.Writer) error {
	funcMap := template.FuncMap{
		"FormatID": func(id uint32) string {
			return fmt.Sprintf("0x%X", id)
		},
		"FormatIDBinary": func(id uint32) string {
			return fmt.Sprintf("0b%b", id)
		},
	}

	var tplText = ledInstructionTemplate
	if len(r.cfg.Template.Template) != 0 {
		tplText = r.cfg.Template.Template
	}

	tpl, err := template.New("keymap-template").
		Funcs(funcMap).
		Parse(tplText)

	if err != nil {
		return err
	}
	return tpl.Execute(writer, r.keyMap)
}
