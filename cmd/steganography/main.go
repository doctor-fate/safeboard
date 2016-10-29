package main

import (
	"log"
	"os"

	"image"
	"image/color"

	"github.com/doctor-fate/safeboard/utils"
	"golang.org/x/image/bmp"
)

func openImage(filename string) (image.Image, error) {
	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	img, err := bmp.Decode(in)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func main() {
	imga, err := openImage("a.bmp")
	if err != nil {
		log.Fatalln(err)
	}
	imgb, err := openImage("b.bmp")
	if err != nil {
		log.Fatalln(err)
	}

	out := image.NewRGBA(imga.Bounds())
	x := imga.Bounds().Max.X
	y := imga.Bounds().Max.Y
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ra, ga, ba, _ := imga.At(j, i).RGBA()
			rb, gb, bb, _ := imgb.At(j, i).RGBA()
			out.Set(j, i, color.RGBA{
				R: byte((ra%2 + rb%2) * 127),
				G: byte((ga%2 + gb%2) * 127),
				B: byte((ba%2 + bb%2) * 127),
				A: 255,
			})
		}
	}

	f, err := os.Create("out.bmp")
	if err != nil {
		log.Fatalln(err)
	}
	if err := bmp.Encode(f, out); err != nil {
		log.Fatalln(err)
	}

	key := "We give you the power of lolcatz"
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
