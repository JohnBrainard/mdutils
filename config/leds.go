package config

import (
	"unicode"
)

type LEDs []string

func LEDsFromString(str string) LEDs {
	runes := []rune(str)
	count := len(runes)
	result := make(LEDs, 0, count)

	for _, r := range runes {
		if !unicode.IsSpace(r) {
			result = append(result, string(r))
		}
	}

	return result
}

func (l *LEDs) UnmarshalText(data []byte) error {
	*l = LEDsFromString(string(data))

	return nil
}
