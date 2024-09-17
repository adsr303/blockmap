package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/adsr303/blockmap/convert"
	"github.com/adsr303/blockmap/options"
	"github.com/adsr303/blockmap/terminal"
	_ "golang.org/x/image/bmp"
)

func main() {
	var opts options.Options
	flag.BoolVar(&opts.UseShadeBlocks, "shade", false, "use double-size shade block characters for alpha")
	flag.StringVar(&opts.Fit, "fit", "none", "fit image within specified size; one of:\nnone, auto, auto-LINES, COLUMNSxLINES")
	flag.StringVar(&opts.Colors, "colors", "auto", "ANSI color palette to use for rendering; one of:\nauto, ansi, ansi256, ansirgb")
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
	imgfile := flag.Arg(0)
	img, err := readImage(imgfile)
	if err != nil {
		log.Fatal(err)
	}
	term := terminal.GetTerminfo()
	rect, err := opts.GetFitRect(term)
	if err != nil {
		log.Fatal(err)
	}
	if rect.Dx() < img.Bounds().Dx() || rect.Dy() < img.Bounds().Dy() {
		img, err = convert.ResizeImage(img, rect)
		if err != nil {
			log.Fatal(err)
		}
	}
	palette, err := opts.GetPalette(term)
	if err != nil {
		log.Fatal(err)
	}
	if opts.UseShadeBlocks {
		fmt.Print(convert.ConvertImageToShadeBlocks(img, palette))
	} else {
		fmt.Print(convert.ConvertImageToHalfBlocks(img, palette))
	}
}

func readImage(imgfile string) (image.Image, error) {
	b, err := os.ReadFile(imgfile)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", imgfile, err)
	}
	r := bytes.NewReader(b)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", imgfile, err)
	}
	return img, nil
}
