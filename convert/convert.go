package convert

import (
	"image"
	"strings"
)

const reset = "\x1b[0m"

var shadeBlocks = []string{"  ", "░░", "▒▒", "▓▓", "██"}

const shadeRange = 0xffff/5 + 1

func ConvertImageToShadeBlocks(img image.Image, pal ANSIPalette) string {
	var builder strings.Builder
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		prev := -1
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			index := pal.ColorIndex(img.At(x, y))
			if index != prev {
				writeCode(&builder, pal.ForegroundCode(index))
				prev = index
			}
			_, _, _, a := img.At(x, y).RGBA()
			block := shadeBlocks[a/shadeRange]
			builder.WriteString(block)
		}
		builder.WriteString(reset)
		builder.WriteRune('\n')
	}
	return builder.String()
}

const upperHalfBlock = "▀"

func ConvertImageToHalfBlocks(img image.Image, pal ANSIPalette) string {
	var builder strings.Builder
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y += 2 {
		topPrev := -1
		bottomPrev := -1
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			index := pal.ColorIndex(img.At(x, y))
			if index != topPrev {
				writeCode(&builder, pal.ForegroundCode(index))
				topPrev = index
			}
			// For images with odd number of lines we leave the lower half
			// in default background color ("transparent") for the last line.
			if y < img.Bounds().Max.Y {
				index = pal.ColorIndex(img.At(x, y+1))
				if index != bottomPrev {
					writeCode(&builder, pal.BackgroundCode(index))
					bottomPrev = index
				}
			}
			builder.WriteString(upperHalfBlock)
		}
		builder.WriteString(reset)
		builder.WriteRune('\n')
	}
	return builder.String()
}

func writeCode(b *strings.Builder, code string) {
	b.WriteString("\x1b[")
	b.WriteString(code)
	b.WriteRune('m')
}
