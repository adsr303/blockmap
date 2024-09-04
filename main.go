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
	fmt.Print(convert.ConvertImageToTerminal(img))
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
