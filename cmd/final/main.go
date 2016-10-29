package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func openImage(filename string) (image.Image, error) {
	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	img, err := png.Decode(in)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	img, err := openImage(path)
	if err != nil {
		return err
	}

	out := image.NewRGBA(img.Bounds())
	x := img.Bounds().Max.X
	y := img.Bounds().Max.Y
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ra, ga, ba, _ := img.At(j, i).RGBA()
			out.Set(j, i, color.RGBA{
				R: byte((ra%2 + ra%2) * 127),
				G: byte((ga%2 + ga%2) * 127),
				B: byte((ba%2 + ba%2) * 127),
				A: 255,
			})
		}
	}

	f, err := os.Create("pictures_out/" + info.Name())
	if err != nil {
		return err
	}
	if err := png.Encode(f, out); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := os.Mkdir("pictures_out", 0775); err != nil {
		log.Fatalln(err)
	}

	if err := filepath.Walk("pictures/", walk); err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("http://safeboard.kaspersky.ru/%x", md5.Sum([]byte("Hackathon 2016 Ultimate Prize")))
	fmt.Println(url)
}
