package ctrlutils

import (
	"strconv"
	"strings"

	"github.com/JohnBrainard/droputils/config"
)

type KeyMap struct {
	cfg config.Config
}

func New(cfg config.Config) KeyMap {
	return KeyMap{cfg: cfg}
}

func (m *KeyMap) Instructions() []LedInstruction {
	instructions := make([]LedInstruction, 0)

	for _, layer := range m.cfg.Layers {
		instructions = append(instructions, m.CreateInstructions(layer)...)
	}

	// for _, row := range m.cfg.Keys.Rows {
	// 	instructions = append(instructions, m.CreateInstructions(row)...)
	// }

	return instructions
}

func (m *KeyMap) CreateInstructions(layer config.Layer) []LedInstruction {
	instructions := make([]LedInstruction, 0)

	for colorKey, colorValue := range m.cfg.Colors {
		leds := m.CreateBinaryString(colorKey, layer)
		if leds != "" {
			ids := m.CreateIDS(leds)

			instruction := NewLedInstruction(ids, colorValue)
			instructions = append(instructions, instruction)
		}
	}

	return instructions
}

func (m *KeyMap) RowKeyMaps(row config.Row, color string) uint32 {
	var keyMask uint32

	for _, rowColor := range row.ColorValues() {
		keyMask = keyMask << 1
		if color == rowColor {
			keyMask |= 1
		}
	}

	return keyMask
}

func (m *KeyMap) CreateBinaryString(colorId string, layer config.Layer) string {
	var str strings.Builder
	var colorUsed = false

	for _, c := range layer.KeyLEDs {
		if c == colorId {
			str.WriteByte('1')
			colorUsed = true
		} else {
			str.WriteByte('0')
		}
	}

	for i := len(layer.KeyLEDs); i < layer.KeyLEDCount; i++ {
		str.WriteByte('0')
	}

	for _, c := range layer.EdgeLEDs {
		if c == colorId {
			str.WriteByte('1')
			colorUsed = true
		} else {
			str.WriteByte('0')
		}
	}

	for i := len(layer.EdgeLEDs); i < layer.EdgeLEDCount; i++ {
		str.WriteByte('0')
	}

	if colorUsed {
		return str.String()
	}
	return ""
}

func (m *KeyMap) CreateIDS(binary string) []uint32 {
	result := make([]uint32, 0)

	var i = 0
	for {
		low := i
		high := minInt(i+32, len(binary))

		uintVal, err := strconv.ParseUint(Reverse(binary[low:high]), 2, 32)
		if err != nil {
			panic(err)
		}
		result = append(result, uint32(uintVal))

		i += 32

		if i > len(binary) {
			break
		}
	}

	return result
}

func Reverse(str string) string {
	runes := []rune(str)
	l := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[l-i] = runes[l-i], runes[i]
	}
	return string(runes)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ESC - HOME
// PgUp - L. Shift
// Z - R. Arrow +
