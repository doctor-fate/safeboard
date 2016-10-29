package main

import (
	"log"
	"os"

	"encoding/csv"

	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/doctor-fate/safeboard/utils"
)

type Record struct {
	timestamp string
	username  string
	from      string
	to        string
	amount    string
}

func main() {
	in, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create("data_fix.csv")
	if err != nil {
		log.Fatalln(err)
	}

	br := bufio.NewReader(in)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		line = strings.Replace(line, ",\"", "\",", -1)
		line = strings.TrimSpace(line)
		fmt.Fprintln(out, line)
	}

	in.Close()
	out.Close()

	in, err = os.Open("data_fix.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	cr := csv.NewReader(in)
	cr.TrimLeadingSpace = true
	records, err := cr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string][]Record)
	for _, v := range records {
		from := v[3]
		m[from] = append(m[from], Record{
			v[0], v[1], v[2], v[3], v[4],
		})
	}

	var mk string
	for k, v := range m {
		if len(v) > len(m[mk]) {
			mk = k
		}
	}

	key := m[mk][0].username + ";" + mk
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "decrypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
