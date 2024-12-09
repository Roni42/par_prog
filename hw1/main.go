package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(MAXN int, ARRLEN int) {
	runtime.GOMAXPROCS(4)
	var t1, t2 int = 0, 0

	for i := 0; i < MAXN; i += 1 {
		arr := genArray(int(ARRLEN))

		var s seqSorter
		s1 := time.Now()
		sres := s.SortImm(arr)
		e1 := int(time.Since(s1).Nanoseconds() / 1000000)
		t1 += e1
		println("1: ", e1, checkSorted(sres))

		var p parSorter
		s2 := time.Now()
		pres := p.SortImm(arr)
		e2 := int(time.Since(s2).Nanoseconds() / 1000000)
		t2 += e2
		println("2: ", e2, checkSorted(pres))
		fmt.Printf("Speedup: %.2fx\n", float64(e1)/float64(e2))

		println("===")
	}
	seqTime, parTime := t1/MAXN, t2/MAXN

	fmt.Printf("seq average time: %d, par average time: %d\n", seqTime, parTime)

	var speedup float64 = float64(seqTime) / float64(parTime)
	fmt.Printf("Speedup: %.2fx\n", speedup)
}

func main() {
	test(5, 1e8)
}
