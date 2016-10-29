package main

import (
	"log"
	"os"
	"strings"

	"github.com/doctor-fate/safeboard/utils"

	"bufio"
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"io"
)

func find(line string, c byte, n int) (sum int) {
	for i, k := 0, 0; k < n && i != -1; k++ {
		if i = strings.IndexByte(line, c); i == -1 {
			return -1
		}
		line = line[i+1:]
		sum += i + 1
	}

	return sum
}

type Key struct {
	ip     string
	uagent string
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

		line = strings.TrimSpace(line)
		line = strings.Replace(line, "\\\"", "\"\"", -1)
		i := find(line, '"', 4)
		if line[i] == ',' {
			continue
		}
		fmt.Fprint(out, line[:i])
		fmt.Fprint(out, ",")
		fmt.Fprintln(out, line[i:])
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

	m := make(map[Key]int)
	for _, v := range records {
		key := Key{ip: v[1], uagent: v[3]}
		m[key]++
	}

	var mk Key
	for k, v := range m {
		if v > m[mk] {
			mk = k
		}
	}

	key := fmt.Sprintf("%x", md5.Sum([]byte(mk.uagent)))
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "decrypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
