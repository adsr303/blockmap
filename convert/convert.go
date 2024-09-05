package convert

import (
	"fmt"
	"image"
	"strings"
)

func ConvertImageToTerminal(img image.Image) string {
	var builder strings.Builder
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		prev := -1
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			index := getOpaqueColorIndex(palette256, img.At(x, y))
			if index != prev {
				builder.WriteString(fmt.Sprintf("\x1b[38;5;%dm", index))
				prev = index
			}
			_, _, _, a := img.At(x, y).RGBA()
			block := alphaBlocks[a/alphaRange]
			builder.WriteString(block)
		}
		builder.WriteString(reset)
		builder.WriteRune('\n')
	}
	return builder.String()
}

const reset = "\x1b[0m"

var alphaBlocks = []string{"  ", "░░", "▒▒", "▓▓", "██"}

const alphaRange = 0xffff/5 + 1
