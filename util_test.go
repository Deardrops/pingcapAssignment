package main

import "testing"

func TestIsLetter(t *testing.T) {
	areaTests := []struct {
		in  uint8
		out bool
	}{
		{'3', false},
		{'z', true},
		{'+', false},
	}

	for _, tt := range areaTests {
		got := isLetter(tt.in)
		if got != tt.out {
			t.Errorf("with %c, got %v, want %v\n", tt.in, got, tt.out)
		}
	}
}
