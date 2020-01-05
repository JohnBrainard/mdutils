package config

import "unicode"

type Row struct {
	Colors string
}

func (r *Row) ColorValues() []string {
	colors := make([]string, 0)

	for _, color := range r.Colors {
		if !unicode.IsSpace(color) {
			colors = append(colors, string([]rune{color}))
		}
	}

	return colors
}
