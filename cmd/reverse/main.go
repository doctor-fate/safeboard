package main

import (
	"io/ioutil"
	"log"

	"github.com/doctor-fate/safeboard/decryptor"
)

func main() {
	dec, err := decryptor.DecryptFiles("encrypted", "key.txt")
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("decrypted.dat", dec, 0644); err != nil {
		log.Fatalln(err)
	}
}
