package convert

import (
	"image/color"
	"testing"
)

func TestGetOpaqueIndex(t *testing.T) {
	index := getOpaqueColorIndex(palette256, color.White)
	if index != 15 {
		t.Errorf("expected index 15, got %d", index)
	}
}

func TestPalette256(t *testing.T) {
	if len(palette256) != 256 {
		t.Errorf("expected 256 colors, got %d", len(palette256))
	}
}
