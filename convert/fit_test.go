package convert

import (
	"image"
	"testing"
)

func TestCalculateFit(t *testing.T) {
	cases := []struct {
		name            string
		r               image.Rectangle
		columns, lines  int
		wantCol, wantLn int
	}{
		{"exact height", image.Rect(0, 0, 25, 25), 80, 25, 25, 25},
		{"icon 32", image.Rect(0, 0, 32, 32), 80, 25, 25, 25},
		{"icon 16", image.Rect(0, 0, 16, 16), 80, 25, 16, 16},
		{"vertical terminal", image.Rect(0, 0, 320, 240), 80, 125, 80, 60},
		{"tall image", image.Rect(0, 0, 180, 320), 80, 25, 14, 25},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dx, dy, err := calculateFit(c.r, c.columns, c.lines)
			if dx != c.wantCol {
				t.Errorf("expected width=%d, got %d", c.wantCol, dx)
			}
			if dy != c.wantLn {
				t.Errorf("expected height=%d, got %d", c.wantLn, dy)
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
