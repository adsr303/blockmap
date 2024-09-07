package convert

import (
	"fmt"
	"image"
	"math"
)

func calculateFit(bounds image.Rectangle, columns, lines int) (int, int, error) {
	dx, dy := bounds.Dx(), bounds.Dy()
	if columns <= 0 || lines <= 0 {
		return 0, 0, fmt.Errorf("invalid terminal size: %dx%x", columns, lines)
	}
	if dx <= 0 || dy <= 0 {
		return 0, 0, fmt.Errorf("invalid image bounts: %v", bounds)
	}
	if columns >= dx && lines >= dy { // TODO Rethink condition - terminfo automargin etc.
		return dx, dy, nil
	}
	ratiox := float64(dx) / float64(columns)
	ratioy := float64(dy) / float64(lines)
	ratio := math.Max(ratiox, ratioy)
	return int(float64(dx) / ratio), int(float64(dy) / ratio), nil
}
