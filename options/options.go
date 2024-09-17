package options

import (
	"errors"
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/adsr303/blockmap/convert"
	"github.com/adsr303/blockmap/terminal"
)

type Options struct {
	UseShadeBlocks bool
	// `[auto[-LINES]|COLUMNSxLINES]`
	Fit string
	// `[ansi|bash|echo|sh]`
	Format string
	// `[auto|ansi|ansi256|ansirgb]`
	Colors string
}

var (
	ErrInvalidFitFormat = errors.New("fit format")
	ErrInvalidColors    = errors.New("colors argument")
)

func (o Options) GetFitRect(term terminal.Terminfo) (image.Rectangle, error) {
	dx, dy, err := parseFit(o.Fit, term)
	if err != nil {
		return image.Rectangle{}, err
	}
	if o.UseShadeBlocks {
		dx /= 2
	} else {
		dy *= 2
	}
	return image.Rect(0, 0, dx, dy), nil
}

var fitAuto = regexp.MustCompile(`^auto-(\d+)$`)
var fitSize = regexp.MustCompile(`^(\d+)x(\d+)$`)

const bigSize = 1 << 12

// parseFit finds the desired maximum dimensions of the output that the user
// wants to generate based on command-line argument and terminal size.
func parseFit(fit string, term terminal.Terminfo) (int, int, error) {
	switch fit {
	case "", "none":
		return bigSize, bigSize, nil
	case "auto":
		return term.Columns, term.Lines, nil
	}
	m := fitAuto.FindStringSubmatch(fit)
	if len(m) == 2 {
		minusLines, err := strconv.Atoi(m[1])
		if err != nil {
			return 0, 0, unexpectedError(err)
		}
		return term.Columns, term.Lines - minusLines, nil
	}
	m = fitSize.FindStringSubmatch(fit)
	if len(m) == 3 {
		columns, err := strconv.Atoi(m[1])
		if err != nil {
			return 0, 0, unexpectedError(err)
		}
		lines, err := strconv.Atoi(m[2])
		if err != nil {
			return 0, 0, unexpectedError(err)
		}
		return columns, lines, nil
	}
	return 0, 0, ErrInvalidFitFormat
}

func unexpectedError(e error) error {
	return fmt.Errorf("unexpected error: %w", e)
}

func (o Options) GetPalette(term terminal.Terminfo) (convert.ANSIPalette, error) {
	switch o.Colors {
	case "ansi":
		return convert.ANSI, nil
	case "ansi256":
		return convert.ANSI256, nil
	case "ansirgb":
		return convert.ANSIRGB, nil
	case "auto":
		switch term.Colors {
		case terminal.Colors3bit:
			return convert.ANSI, nil
		case terminal.Colors8bit:
			return convert.ANSI256, nil
		case terminal.Colors24bit:
			return convert.ANSIRGB, nil
		default:
			return convert.ANSI, nil
		}
	default:
		return nil, ErrInvalidColors
	}
}
