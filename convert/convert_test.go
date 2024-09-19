package convert

import (
	"image"
	"image/color"
	"strings"
	"testing"

	"github.com/adsr303/blockmap/palettes"
)

func TestConvertImageToHalfBlocks(t *testing.T) {
	cases := []struct {
		img   image.Image
		lines int
	}{
		{makeRect(3, 2), 1},
		{makeRect(3, 3), 2},
		{makeRect(3, 4), 2},
		{makeRect(3, 5), 3},
	}
	for _, c := range cases {
		s := ConvertImageToHalfBlocks(c.img, palettes.ANSI256)
		if !strings.HasSuffix(s, "\n") {
			t.Error("expected newline at end")
		}
		lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")
		if len(lines) != c.lines {
			t.Errorf("expected %d lines, got %d", c.lines, len(lines))
		}
	}
}

func makeRect(width, height int) image.Image {
	img := image.NewGray(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			img.Set(x, y, color.Black)
		}
	}
	return img
}
