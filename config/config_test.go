package config

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const configTOML = `
[Colors]
A = "FF0000" # Alpha, Space Keys
B = "00FF00" # Punctuation, Symbol Keyslion
C = "0000FF" # Modifier Keys
D = "00FFFF" # Navigation Keys
E = "FFFF00" # Number Keys

[[Layers]]
KeyLEDCount = 87
EdgeLEDCount = 32

KeyLEDs = '''
	CCC B CCCC DDD
	C AAAAAAA BBB C D
	C AAAAAAAAA BB C
	C AAAAAAAAAA BBB DDD
	B EEEEEEEEEE BB C DDD
	C CCCC CCCC CCCC DDD
'''

EdgeLEDs = '''
	ABCABCABCABCA BAB ABCABCABCABCA ABA
'''
`

func TestRead(t *testing.T) {
	t.Run("Read returns a Config instance", func(t *testing.T) {
		reader := strings.NewReader(configTOML)
		config, err := Read(reader)
		if err != nil {
			t.Fatalf("unexpected error: %v, cause: %v", err, errors.Unwrap(err))
		}

		expected := Config{
			Colors: map[string]HexColor{
				"A": "FF0000",
				"B": "00FF00",
				"C": "0000FF",
				"D": "00FFFF",
				"E": "FFFF00",
			},
			Layers: []Layer{{
				KeyLEDCount:  87,
				EdgeLEDCount: 32,

				KeyLEDs:  LEDsFromString("CCCBCCCCDDDCAAAAAAABBBCDCAAAAAAAAABBCCAAAAAAAAAABBBDDDBEEEEEEEEEEBBCDDDCCCCCCCCCCCCCDDD"),
				EdgeLEDs: LEDsFromString("ABCABCABCABCABABABCABCABCABCAABA"),
			}},
		}

		if diff := cmp.Diff(expected, config); diff != "" {
			t.Errorf("result not expected: %v", diff)
		}
	})
}
