package main

import (
	"log"
	"strconv"

	"strings"

	"io/ioutil"
	"regexp"

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

type Operators struct {
	data [1024]string
	top  int
}

func (w *Operators) Push(x string) {
	w.data[w.top] = x
	w.top++
}

func (w *Operators) Peek() string {
	return w.data[w.top-1]
}

func (w *Operators) Pop() string {
	w.top--
	return w.data[w.top]
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
	var (
		q Operands
		w Operators
	)

	buf, err := ioutil.ReadFile("task.txt")
	if err != nil {
		log.Fatalln(err)
	}

	r, err := regexp.Compile("[A-Za-z0-9]+|[!)(#>=]")
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range r.FindAll(buf, -1) {
		switch s := string(v); s {
		case "#", "!", ">", "=":
			for w.top != 0 && w.Peek() != "(" {
				op(w.Pop(), &q)
			}
			w.Push(s)
		case "(":
			w.Push(s)
		case ")":
			for o := w.Pop(); o != "("; o = w.Pop() {
				op(o, &q)
			}
		default:
			var n uint64
			if len(s) > 2 && s[:2] == "0t" {
				n, err = strconv.ParseUint(s[2:], 13, 64)
			} else {
				n, err = strconv.ParseUint(s, 7, 64)
			}
			if err != nil {
				log.Fatalln(err)
			}
			q.Push(n)
		}
	}

	for w.top != 0 {
		op(w.Pop(), &q)
	}

	key := "0t" + strings.ToUpper(strconv.FormatUint(q.Pop(), 13))
	err = utils.WriteKeyAndDecrypt([]byte(key), "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
