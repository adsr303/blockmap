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
	_ "golang.org/x/image/bmp"
)

func main() {
	var useShadeBlocks bool
	flag.BoolVar(&useShadeBlocks, "shade", false, "use double-size shade block characters for alpha")
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
	if useShadeBlocks {
		fmt.Print(convert.ConvertImageToShadeBlocks(img))
	} else {
		fmt.Print(convert.ConvertImageToHalfBlocks(img))
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
