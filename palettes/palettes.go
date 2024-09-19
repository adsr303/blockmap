package palettes

import (
	"fmt"
	"image/color"
	"strconv"
)

type ANSIPalette interface {
	ColorIndex(c color.Color) int
	ForegroundCode(index int) string
	BackgroundCode(index int) string
}

var (
	ANSI    = ansi{}
	ANSI256 = ansi256{}
	ANSIRGB = ansirgb{}
)

type ansi struct{}

func (a ansi) ColorIndex(c color.Color) int {
	return getOpaqueColorIndex(palette8, c)
}

func (a ansi) ForegroundCode(index int) string {
	return strconv.Itoa(index + 30)
}

func (a ansi) BackgroundCode(index int) string {
	return strconv.Itoa(index + 40)
}

type ansi256 struct{}

func (a ansi256) ColorIndex(c color.Color) int {
	return getOpaqueColorIndex(palette256, c)
}

func (a ansi256) ForegroundCode(index int) string {
	return fmt.Sprintf("38;5;%d", index)
}

func (a ansi256) BackgroundCode(index int) string {
	return fmt.Sprintf("48;5;%d", index)
}

type ansirgb struct{}

func (a ansirgb) ColorIndex(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int((r & 0xff00 << 8) | (g & 0xff00) | (b & 0xff00 >> 8))
}

func (a ansirgb) ForegroundCode(index int) string {
	r, g, b := splitRGB(index)
	return fmt.Sprintf("38;2;%d;%d;%d", r, g, b)
}

func (a ansirgb) BackgroundCode(index int) string {
	r, g, b := splitRGB(index)
	return fmt.Sprintf("48;2;%d;%d;%d", r, g, b)
}

func splitRGB(index int) (int, int, int) {
	r := index & 0xff0000 >> 16
	g := index & 0xff00 >> 8
	b := index & 0xff
	return r, g, b
}

func getOpaqueColorIndex(p color.Palette, c color.Color) int {
	r, g, b, _ := c.RGBA()
	opaque := color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 0xff}
	return p.Index(opaque)
}

// See https://en.wikipedia.org/wiki/ANSI_escape_code#3-bit_and_4-bit
var palette8 color.Palette

// See https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
var palette256 color.Palette

func init() {
	// Standard & high-intensity colors
	palette256 = append(palette256, color.Black)
	for _, num := range []uint32{
		0x800000,
		0x008000,
		0x808000,
		0x000080,
		0x800080,
		0x008080,
		0xc0c0c0,
		0x808080,
		0xff0000,
		0x00ff00,
		0xffff00,
		0x0000ff,
		0xff00ff,
		0x00ffff,
	} {
		palette256 = append(palette256, makeNRGBA(num))
	}
	palette256 = append(palette256, color.White)

	// 216 colors
	edge := []uint8{0x00, 0x5f, 0x87, 0xaf, 0xd7, 0xff}
	for r := range 6 {
		for g := range 6 {
			for b := range 6 {
				palette256 = append(palette256, color.NRGBA{edge[r], edge[g], edge[b], 0xff})
			}
		}
	}

	// Grayscale colors
	for i := range 24 {
		palette256 = append(palette256, color.Gray{uint8(0x08 + i*0x0a)})
	}

	palette8 = palette256[:8]
}

func makeNRGBA(num uint32) color.Color {
	r := (num >> 16) & 0xff
	g := (num >> 8) & 0xff
	b := num & 0xff
	return color.NRGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
