package options

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/adsr303/blockmap/terminal"
)

type Options struct {
	// `[auto[-LINES]|COLUMNSxLINES]`
	Fit string
	// `[ansi|bash|echo|sh]`
	Format string
	// `[auto|8|16|256]`
	Colors string
}

var ErrInvalidFitFormat = errors.New("fit format")

var fitAuto = regexp.MustCompile(`^auto(-\d+)?$`)
var fitSize = regexp.MustCompile(`^(\d+)x(\d+)$`)

// parseFit finds the desired maximum dimensions of the output that the user
// wants to generate based on command-line argument and terminal size.
func parseFit(fit string, term terminal.Terminfo) (int, int, error) {
	if fit == "" {
		return term.Columns, term.Lines, nil
	}
	m := fitAuto.FindStringSubmatch(fit)
	switch len(m) {
	case 1:
		return term.Columns, term.Lines, nil
	case 2:
		if m[1] == "" {
			return term.Columns, term.Lines, nil
		}
		minusLines, err := strconv.Atoi(m[1])
		if err != nil {
			return 0, 0, unexpectedError(err)
		}
		return term.Columns, term.Lines + minusLines, nil
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
