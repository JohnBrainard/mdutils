package ctrlutils

import (
	"strings"

	"github.com/JohnBrainard/droputils/config"
)

type LedInstruction struct {
	ID0, ID1, ID2, ID3 uint32
	Red, Green, Blue   uint8
	MatchLayers        []int
}

func NewLedInstruction(ids []uint32, color config.HexColor, matchLayers []int) LedInstruction {
	if len(ids) > 4 {
		panic("max of 4 ids values supported")
	}

	var instruction LedInstruction
	instruction.MatchLayers = matchLayers

	for i, id := range ids {
		switch i {
		case 0:
			instruction.ID0 = id

		case 1:
			instruction.ID1 = id

		case 2:
			instruction.ID2 = id

		case 3:
			instruction.ID3 = id
		}
	}

	instruction.Red, instruction.Green, instruction.Blue = color.RGB()

	return instruction
}

func (i LedInstruction) IDs() []uint32 {
	ids := make([]uint32, 0)

	switch {
	case i.ID3 != 0:
		ids = append([]uint32{i.ID3}, ids...)
		fallthrough

	case i.ID2 != 0:
		ids = append([]uint32{i.ID2}, ids...)
		fallthrough

	case i.ID1 != 0:
		ids = append([]uint32{i.ID1}, ids...)
		fallthrough

	case i.ID0 != 0:
		ids = append([]uint32{i.ID0}, ids...)
	}

	return ids
}

func (i LedInstruction) Flags() string {
	flags := []string{
		"LED_FLAG_MATCH_ID",
		"LED_FLAG_USE_RGB",
	}

	if len(i.MatchLayers) != 0 {
		flags = append(flags, "LED_FLAG_MATCH_LAYER")
	}

	return strings.Join(flags, " | ")
}
