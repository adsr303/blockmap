package terminal

import (
	"os"
	"strings"

	"golang.org/x/term"
)

type Colors int

const (
	ColorsUnknown Colors = 0
	Colors3bit    Colors = 8
	Colors4bit    Colors = 16
	Colors8bit    Colors = 256
)

type Terminfo struct {
	Columns, Lines int
	Colors
}

func GetTerminfo() Terminfo {
	columns, lines, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		// TODO For now assuming dumb terminals, might return error later
		columns = 80
		lines = 24
	}
	var colors Colors
	term := os.Getenv("TERM")
	switch {
	case term == "":
		colors = Colors3bit
	case strings.HasSuffix(term, "-256color"):
		colors = Colors8bit
	}
	colorterm := os.Getenv("COLORTERM")
	switch colorterm {
	case "24bit", "millions":
		colors = Colors8bit
	}
	return Terminfo{Columns: columns, Lines: lines, Colors: colors}
}
