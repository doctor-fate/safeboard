package main

import (
	"log"
	"os"

	"github.com/doctor-fate/safeboard/utils"

	"encoding/csv"
)

func main() {
	in, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	r := csv.NewReader(in)
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string]int)
	for _, v := range records {
		key := v[2]
		m[key]++
	}

	var mk string
	for k, v := range m {
		if v > m[mk] {
			mk = k
		}
	}

	err = utils.WriteKeyAndDecrypt([]byte(mk), "key.txt", "encrypted", "decrypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
