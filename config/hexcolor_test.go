package config

import "testing"

func TestHexColor_RGB(t *testing.T) {
	tests := []struct {
		name  string
		c     HexColor
		wantR uint8
		wantG uint8
		wantB uint8
	}{
		{
			name:  "white",
			c:     "FFFFFF",
			wantR: 255,
			wantG: 255,
			wantB: 255,
		},
		{
			name:  "red",
			c:     "FF0000",
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "green",
			c:     "FF00",
			wantR: 0,
			wantG: 255,
			wantB: 0,
		},
		{
			name:  "blue",
			c:     "FF",
			wantR: 0,
			wantG: 0,
			wantB: 255,
		},
		{
			name:  "black",
			c:     "0",
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "random",
			c:     "3b5187",
			wantR: 59,
			wantG: 81,
			wantB: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := tt.c.RGB()
			if gotR != tt.wantR {
				t.Errorf("RGB() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("RGB() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("RGB() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
