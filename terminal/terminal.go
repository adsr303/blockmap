package terminal

import (
	"os"
	"strconv"
	"strings"
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
	columns := getNumEnv("COLUMNS")
	lines := getNumEnv("LINES")
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

func getNumEnv(name string) int {
	s := os.Getenv(name)
	if s == "" {
		return 0
	}
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}
