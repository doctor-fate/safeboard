package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/doctor-fate/safeboard/utils"
)

func main() {
	in, err := os.Open("task.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sum := uint64(0)
	for i := 1; ; i++ {
		var s string
		if _, err := fmt.Fscan(in, &s); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		n, err := strconv.ParseUint(s, 7, 64)
		if err != nil {
			log.Fatalf("Parsing error: %s on line %d", err.Error(), err)
		}
		sum += n
	}

	key := strconv.FormatUint(sum, 7)
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
