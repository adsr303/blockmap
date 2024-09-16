package convert

import (
	"image/color"
	"testing"
)

func TestAnsi256_ColorIndex(t *testing.T) {
	var a ansi256
	index := a.ColorIndex(color.Black)
	if index != 0 {
		t.Errorf("expected index 0, got %d", index)
	}
}

func TestAnsi24bit_ColorIndex(t *testing.T) {
	var a ansi24bit
	index := a.ColorIndex(color.White)
	if index != 0xffffff {
		t.Errorf("expected #FFFFFF, got #%06X", index)
	}
}

func TestGetOpaqueColorIndex(t *testing.T) {
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
