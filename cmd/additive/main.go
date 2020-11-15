package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"os"
)

var support = []uint8{1, 2, 1, 2, 2, 1, 2, 1}

func main() {
	sourcePath := flag.String("s", "", "Source BMP file")
	targetPath := flag.String("t", "", "Target BMP file")
	message := flag.String("m", "", "Secret message")
	flag.Parse()
	source, err := read(*sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	if *message != "" {
		err = encode(source, *message, *targetPath)
	} else {
		err = decode(source, *targetPath)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func read(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bitmap, err := bmp.Decode(file)
	if err != nil {
		_ = file.Close()
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	return bitmap, nil
}

func encode(source image.Image, message, targetPath string) error {
	bounds := source.Bounds()
	target := image.NewRGBA(bounds)
	bytes, i := []byte(message), 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			colour, j := paint(source.At(x, y), bytes, i)
			target.Set(x, y, colour)
			i = j
		}
	}
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	if err = bmp.Encode(file, target); err != nil {
		_ = file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func paint(colour color.Color, bytes []byte, i int) (color.Color, int) {
	if i >= len(bytes) {
		return colour, i
	}
	j := i % len(support)
	lw := support[j] * bytes[i]
	if lw < support[j] || lw < bytes[i] {
		return colour, i
	}
	r, g, b, a := colour.RGBA()
	tr, tg, tb := uint8(r>>8), uint8(g>>8), uint8(b>>8)
	if s := tr + lw; s >= tr && s >= lw {
		tr = s
	} else if s := tg + lw; s >= tg && s >= lw {
		tg = s
	} else if s := tb + lw; s >= tb && s >= lw {
		tb = s
	} else {
		return colour, i
	}
	return color.RGBA{R: tr, G: tg, B: tb, A: uint8(a >> 8)}, i + 1
}

func decode(source image.Image, targetPath string) error {
	target, err := read(targetPath)
	if err != nil {
		return err
	}
	bounds := source.Bounds()
	if bounds != target.Bounds() {
		return fmt.Errorf(
			"main: additive got inconcistent bounds, %v != %v",
			bounds,
			target.Bounds(),
		)
	}
	message, i := make([]byte, 0), 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if w, ok := extract(source.At(x, y), target.At(x, y), i); ok {
				message = append(message, w)
				i++
			}
		}
	}
	log.Infof("main: additive decoded a message '%s'", message)
	return nil
}

func extract(source, target color.Color, i int) (byte, bool) {
	sr, sg, sb, _ := source.RGBA()
	tr, tg, tb, _ := target.RGBA()
	l := support[i%len(support)]
	if sr != tr {
		return (byte(tr>>8) - byte(sr>>8)) / l, true
	}
	if sg != tg {
		return (byte(tg>>8) - byte(sg>>8)) / l, true
	}
	if sb != tb {
		return (byte(tb>>8) - byte(sb>>8)) / l, true
	}
	return 0, false
}
