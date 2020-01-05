package config

import (
	"errors"
	"unicode"
)

type Layer struct {
	MatchLayers []int
	Keys        string
	Edges       string

	KeyLEDs     LEDs
	KeyLEDCount int

	EdgeLEDs     LEDs
	EdgeLEDCount int
}

func (l *Layer) EdgeColors() []string {
	var colors []string

	runes := []rune(l.Edges)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			colors = append(colors, string([]rune{r}))
		}
	}

	return colors
}

func (l *Layer) Validate() {
	if l.KeyLEDCount < len(l.KeyLEDs) {
		panic(errors.New("KeyLEDCount must be greater than or equal to the number of KeyLEDs"))
	}

	if l.EdgeLEDCount < len(l.EdgeLEDs) {
		panic(errors.New("EdgeLEDCount must be greater than or equal to the number of EdgeLEDs"))
	}
}
