package main

import (
	"runtime"
	"sync"
)

type parSorter struct {
	qsortBase
	s seqSorter
}

func (p *parSorter) Sort(arr []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	p.sort(arr, 0, len(arr), &wg, runtime.GOMAXPROCS(0))
	wg.Wait()
}

func (p *parSorter) SortImm(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	p.Sort(res)
	return res
}

func (s *parSorter) sort(arr []int, from, to int, wg *sync.WaitGroup, maxDepth int) {
	defer wg.Done()
	// println("from to", from, to)

	if maxDepth <= 0 {
		s.s.sort(arr, from, to)
		return
	}

	if from < to {
		// println("in")

		p1, p2 := s.partition(arr, from, to)
		// printArr(arr, from, p1, false)
		// printArr(arr, p1, to, false)
		var wg1 sync.WaitGroup
		if p1-from > 1 {
			wg1.Add(1)
			go s.sort(arr, from, p1, &wg1, maxDepth-1)
		}
		if to-p2 > 1 {
			wg1.Add(1)
			go s.sort(arr, p2, to, &wg1, maxDepth-1)
		}
		wg1.Wait()
	}
}
