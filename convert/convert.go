package convert

import (
	"image"
	"strings"
)

func ConvertImageToTerminal(img image.Image) string {
	var builder strings.Builder
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			v := r/3 + g/3 + b/3
			builder.WriteRune(gradient[v/alphaRange])
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

var gradient = []rune{' ', '░', '▒', '▓', '█'}

const alphaRange = 0xffff / 5
