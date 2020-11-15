package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"os"
)

var support = []uint32{5, 18, 23, 4, 7, 1, 9, 2, 2, 9, 1, 7, 4, 23, 18, 5}

func main() {
	sourcePath := flag.String("s", "", "Source BMP file")
	targetPath := flag.String("t", "", "Target BMP file")
	message := flag.String("m", "", "Secret message")
	flag.Parse()
	file, err := os.Open(*sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	source, err := bmp.Decode(file)
	if err != nil {
		_ = file.Close()
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	bounds := source.Bounds()
	target := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			target.Set(x, y, source.At(x, y))
		}
	}
	width := bounds.Max.X - bounds.Min.X
	for i, w := range *message {
		x, y := (i<<1)%width, (i<<1)/width
		r, g, b, a := source.At(x, y).RGBA()
		lw := support[i%len(support)] * uint32(w)
		target.SetRGBA(
			x,
			y,
			color.RGBA{
				R: uint8((r + lw) >> 8),
				G: uint8((g + lw) >> 8),
				B: uint8((b + lw) >> 8),
				A: uint8(a >> 8),
			},
		)
	}
	file, err = os.Create(*targetPath)
	if err != nil {
		log.Fatal(err)
	}
	if err = bmp.Encode(file, target); err != nil {
		_ = file.Close()
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
