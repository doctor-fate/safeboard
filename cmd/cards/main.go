package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"

	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/doctor-fate/safeboard/utils"
)

type User struct {
	username string
	password string
	card     string
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].card < u[j].card
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func luhn(card string) bool {
	n := len(card)
	sum := 0
	for i, v := range card {
		x := int(v - '0')
		if i&1 == n&1 {
			if x *= 2; x > 9 {
				x -= 9
			}
		}
		sum += x
	}

	return (sum % 10) == 0
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

	users := make(Users, 0, len(records))
	for i, v := range records {
		if i != 0 && luhn(v[5]) {
			users = append(users, User{
				username: v[0],
				password: v[3],
				card:     v[5],
			})
		}
	}

	sort.Sort(users)
	u := users[41]
	key := u.username + ";" + u.password
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "decrypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
