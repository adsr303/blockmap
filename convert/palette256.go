package convert

import "image/color"

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
}

func makeNRGBA(num uint32) color.Color {
	r := (num >> 16) & 0xff
	g := (num >> 8) & 0xff
	b := num & 0xff
	return color.NRGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
