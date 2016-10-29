package main

import (
	"log"
	"os"

	"container/heap"

	"github.com/doctor-fate/safeboard/utils"
	"golang.org/x/image/bmp"
)

type Vertex struct {
	dist   uint32
	value  uint32
	index  int
	number int
	parent *Vertex
}

type PQ []*Vertex

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	v := x.(*Vertex)
	v.index = n
	*pq = append(*pq, v)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	v := old[n-1]
	v.index = -1
	*pq = old[0 : n-1]
	return v
}

func (pq *PQ) update(v *Vertex, dist uint32) {
	v.dist = dist
	heap.Fix(pq, v.index)
}

func adjacent(v *Vertex, maze []Vertex, n int) (adj []*Vertex) {
	p := v.number
	i, j := p/n, p%n

	if i > 0 {
		adj = append(adj, &maze[p-n])
	}
	if i < n-1 {
		adj = append(adj, &maze[p+n])
	}
	if j > 0 {
		adj = append(adj, &maze[p-1])
	}
	if j < n-1 {
		adj = append(adj, &maze[p+1])
	}

	return
}

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

	n := img.Bounds().Max.Y
	maze := make([]Vertex, n*n)
	for k, i := 0, 0; i < n; i++ {
		for j := 0; j < n; j, k = j+1, k+1 {
			col := img.At(j, i)
			r, g, b, _ := col.RGBA()
			maze[k] = Vertex{
				dist:   0xFFFF,
				value:  (r >> 8) + (g >> 8) + (b >> 8),
				index:  k,
				number: k,
			}
		}
	}

	pq := make(PQ, len(maze))
	for i := range maze {
		pq[i] = &maze[i]
	}

	maze[0].dist = maze[0].value
	heap.Init(&pq)
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Vertex)
		for _, u := range adjacent(v, maze, n) {
			if u.index != -1 && v.dist+u.value < u.dist {
				pq.update(u, v.dist+u.value)
				u.parent = v
			}
		}
	}

	var path []*Vertex
	for v := &maze[n*n-1]; v != nil; v = v.parent {
		path = append(path, v)
	}

	key := []byte{
		byte(maze[n*n-1].dist),
	}

	n = len(path)
	for i := 1; i <= n && i < 32; i++ {
		v := path[n-i]
		key = append(key, byte(v.value))
	}

	err = utils.WriteKeyAndDecrypt(key, "key.txt", "encrypted", "dectypted.dat")
	if err != nil {
		log.Fatalln(err)
	}
}
