package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/doctor-fate/safeboard/utils"
)

type Operands struct {
	data [1024]uint64
	top  int
}

func (q *Operands) Push(x uint64) {
	q.data[q.top] = x
	q.top++
}

func (q *Operands) Pop() uint64 {
	q.top--
	return q.data[q.top]
}

func op(s string, q *Operands) {
	b, a := q.Pop(), q.Pop()
	switch s {
	case "#":
		q.Push(^(a & b))
	case "!":
		q.Push(^(a | b))
	case ">":
		q.Push(^a | b)
	case "=":
		q.Push((^a | b) & (a | ^b))
	}
}

func main() {
	in, err := os.Open("task.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var q Operands
	for i := 1; ; i++ {
		var s string
		if _, err := fmt.Fscan(in, &s); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		switch s {
		case "#", "!", ">", "=":
			op(s, &q)
		default:
			n, err := strconv.ParseUint(s, 7, 64)
			if err != nil {
				log.Fatalf("Parsing error: %s on line %d", err.Error(), err)
			}
			q.Push(n)
		}
	}

	key := strconv.FormatUint(q.Pop(), 7)
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
