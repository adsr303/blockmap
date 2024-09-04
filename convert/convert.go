package convert

import (
	"image"
	"strings"

	"github.com/fatih/color"
)

func ConvertImageToTerminal(img image.Image) string {
	color.NoColor = false
	var builder strings.Builder
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		var prev *color.Color
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			c := getColorIndex(r, g, b)
			co := color.Set(colors[c])
			if !co.Equals(prev) {
				co.SetWriter(&builder)
			}
			block := blocks[a/alphaRange]
			builder.WriteString(block)
			prev = co
		}
		prev.UnsetWriter(&builder)
		builder.WriteRune('\n')
	}
	return builder.String()
}

var blocks = []string{"  ", "░░", "▒▒", "▓▓", "██"}
var colors = []color.Attribute{
	color.FgBlack,
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,
	color.FgWhite,
	color.FgHiBlack,
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
	color.FgHiWhite,
}

const alphaRange = 0xffff/5 + 1
const halfBright = 0x7fff

func getColorIndex(r, g, b uint32) int {
	var result int
	if r > halfBright {
		result |= 0x1
	}
	if g > halfBright {
		result |= 0x2
	}
	if b > halfBright {
		result |= 0x4
	}
	if (r+g+b)/3 > 0xbfff {
		result += 8
	}
	return result
}
