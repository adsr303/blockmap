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
	b, err := os.ReadFile(imgfile)
	if err != nil {
		log.Fatalf("reading %s: %v", imgfile, err)
	}
	r := bytes.NewReader(b)
	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalf("reading %s: %v", imgfile, err)
	}
	fmt.Print(convert.ConvertImageToTerminal(img))
}
