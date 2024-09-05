package convert

import "testing"

func TestPalette256(t *testing.T) {
	if len(palette256) != 256 {
		t.Errorf("expected 256 colors, got %d", len(palette256))
	}
}
