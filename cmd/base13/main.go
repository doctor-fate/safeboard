package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"strings"

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

		var n uint64
		if len(s) > 2 && s[:2] == "0t" {
			n, err = strconv.ParseUint(s[2:], 13, 64)
		} else {
			n, err = strconv.ParseUint(s, 7, 64)
		}
		if err != nil {
			log.Fatalf("Parsing error: %s on line %d", err.Error(), err)
		}
		sum += n
	}

	key := "0t" + strings.ToUpper(strconv.FormatUint(sum, 13))
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
