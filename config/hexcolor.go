package config

import (
	"fmt"
	"strconv"
)

type HexColor string

func (c HexColor) RGB() (r uint8, g uint8, b uint8) {
	color, err := strconv.ParseUint(string(c), 16, 32)
	if err != nil {
		panic(fmt.Errorf("'%v' is not a valid RGB color string: %v", c, err))
	}

	b = uint8(color % 256)
	g = uint8((color / 256) % 256)
	r = uint8((color / 65_536) % 256)

	return
}
