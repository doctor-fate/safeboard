package main

import (
	"log"
	"os"

	"github.com/doctor-fate/safeboard/utils"
	"golang.org/x/image/bmp"
)

func main() {
	in, err := os.Open("task.bmp")
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	img, err := bmp.Decode(in)
	if err != nil {
		log.Fatalln(err)
	}

	key := make([]byte, 0, 32)
outer:
	for i := 0; i < img.Bounds().Max.Y; i++ {
		for j := 0; j < img.Bounds().Max.X; j++ {
			col := img.At(j, i)
			r, g, b, _ := col.RGBA()
			if r == g && r == b {
				key = append(key, byte(r>>8))
			}

			if len(key) >= 32 {
				break outer
			}
		}
	}

	err = utils.WriteKeyAndDecrypt(key, "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
