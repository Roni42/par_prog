package main

import (
	"log"
	"time"
)

func generateCubicGraph(side int) Graph {
	graph := make(Graph)

	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			for z := 0; z < side; z++ {
				node := x*side*side + y*side + z
				neighbors := []int{}

				if x > 0 {
					neighbors = append(neighbors, (x-1)*side*side+y*side+z)
				}
				if x < side-1 {
					neighbors = append(neighbors, (x+1)*side*side+y*side+z)
				}
				if y > 0 {
					neighbors = append(neighbors, x*side*side+(y-1)*side+z)
				}
				if y < side-1 {
					neighbors = append(neighbors, x*side*side+(y+1)*side+z)
				}
				if z > 0 {
					neighbors = append(neighbors, x*side*side+y*side+(z-1))
				}
				if z < side-1 {
					neighbors = append(neighbors, x*side*side+y*side+(z+1))
				}

				graph[node] = neighbors
			}
		}
	}

	return graph
}

type Graph map[int][]int

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}

func test(gr Graph, times int64) {
	var s1, s2 int64
	for i := range times {
		log.Println(">> time ", i, " out of ", times)
		n := time.Now()
		r1 := SeqBFS(gr, 0, make([]int, len(gr)))
		s1 += time.Since(n).Nanoseconds()

		n = time.Now()
		r2 := ParBFS(gr, 0, make([]int, len(gr)))
		s2 += time.Since(n).Nanoseconds()

		if !eq(r1, r2) {
			panic("Different r1, r2!")
		}
		log.Printf("### %d Result avg time: seq %.3f, par %.3f, speed up: %f",
			i, float64(s1)/float64(i+1),
			float64(s2)/float64(i+1), float64(s1)/float64(s2))

	}
	log.Printf("Result avg time: \nseq: %.3f,\npar: %.3f,\nspeed up: %f",
		float64(s1)/float64(times),
		float64(s2)/float64(times),
		float64(s1)/float64(s2))

}

func main() {
	graph := Graph{
		0: {1, 3},
		1: {2, 3},
		2: {4, 5},
		3: {},
		4: {},
		5: {6},
		6: {},
	}
	test(graph, 1)

	gr := generateCubicGraph(100)
	log.Println("size of graph: ", len(gr))
	test(gr, 5)
}
