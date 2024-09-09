package options

import (
	"math"
	"testing"

	"github.com/adsr303/blockmap/terminal"
)

func TestParseFit(t *testing.T) {
	cases := []struct {
		value          string
		columns, lines int
		wantError      bool
	}{
		{"", math.MaxInt, math.MaxInt, false},
		{"none", math.MaxInt, math.MaxInt, false},
		{"auto", 80, 24, false},
		{"auto-2", 80, 22, false},
		{"32x32", 32, 32, false},
		{"32", 0, 0, true},
		{"32-2", 0, 0, true},
		{"garbage", 0, 0, true},
	}
	term := terminal.Terminfo{Columns: 80, Lines: 24, Colors: terminal.Colors8bit}
	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			columns, lines, err := parseFit(c.value, term)
			if c.wantError {
				switch err {
				case ErrInvalidFitFormat:
					return
				case nil:
					t.Errorf("expected %T, got nil and %dx%d", ErrInvalidFitFormat, columns, lines)
				default:
					t.Errorf("expected %T, got %v", ErrInvalidFitFormat, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if columns != c.columns || lines != c.lines {
					t.Errorf("expected %dx%d, got %d, %d", c.columns, c.lines, columns, lines)
				}
			}
		})
	}
}
