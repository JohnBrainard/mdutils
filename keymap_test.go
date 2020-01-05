package ctrlutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/JohnBrainard/droputils/config"
)

func TestKeyMap_CreateInstructions(t *testing.T) {
	cfg := config.Config{
		Colors: map[string]config.HexColor{
			"A": "FF0000",
			"B": "00FF00",
		},
		Layers: []config.Layer{{
			KeyLEDCount:  87,
			EdgeLEDCount: 32,

			KeyLEDs:  config.LEDsFromString("AAA B AAAA BBB	A BBBBBBB AAA B A A BBBBBBBBB AA BA BBBBBBBBBB AAA BBB A BBBBBBBBBB AA B AAA B BBBB BBBB BBBB BBB"),
			EdgeLEDs: config.LEDsFromString("ABABABABABABA BAB BABABABABABAB ABA"),
		}},
	}
	keyMap := KeyMap{cfg: cfg}

	t.Run("creates led instructions", func(t *testing.T) {
		result := keyMap.CreateInstructions(cfg.Layers[0])
		var expected = []LedInstruction{
			{
				ID0: 0x01b808f7,
				ID1: 0x47002c,
				ID2: 0xaa800076,
				ID3: 0x55552a,
				Red: 0xff,
			},
			{
				ID0:   0xfe47f708,
				ID1:   0xffb8ffd3,
				ID2:   0x557fff89,
				ID3:   0x2aaad5,
				Green: 0xff,
			},
		}

		if diff := cmp.Diff(expected, result); diff != "" {
			t.Errorf("result not expected: %v", diff)
		}
	})
}

func TestKeyMap_RowKeyMaps(t *testing.T) {
	var keyMap KeyMap

	t.Run("creates a valid key mask", func(t *testing.T) {
		row := config.Row{
			Colors: "CCC B CCCC DDD",
		}

		var expected uint32 = 0b11101111000
		var result = keyMap.RowKeyMaps(row, "C")

		if expected != result {
			t.Errorf("expected %b, got %b instead", expected, result)
		}
	})
}

func TestKeyMap_CreateIDS(t *testing.T) {
	tests := []struct {
		name     string
		arg      string
		expected []uint32
	}{
		{
			name: "All 1s",
			arg:  "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111",
			expected: []uint32{
				0xFFFFFFFF,
				0xFFFFFFFF,
				0x007FFFFF,
			},
		},
		{
			name: "Real values",
			arg:  "111011110001000000011101100000000011010000000000111000100000000001101110000000000000000",
			expected: []uint32{
				0x01b808f7, 0x47002c, 0x76,
			},
		},
	}

	// var binaryString = `111011110001000000011101100000000011010000000000111000100000000001101110000000000000000`
	keyMap := KeyMap{}
	//
	for _, tt := range tests {
		t.Run("converts a binary string into 32 bit uints", func(t *testing.T) {
			result := keyMap.CreateIDS(tt.arg)
			if diff := cmp.Diff(tt.expected, result); diff != "" {
				t.Errorf("result not expected: %v", diff)
			}
		})
	}

}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "even length string",
			args: "ABCD",
			want: "DCBA",
		},
		{
			name: "odd length string",
			args: "ABCDE",
			want: "EDCBA",
		},
		{
			name: "zero length string",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
