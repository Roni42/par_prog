package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func run(arr []int, name string, s sorter) int {
	stTime := time.Now()
	sres := s.SortImm(arr)
	runTime := int(time.Since(stTime).Nanoseconds() / 1000000)
	println(name, ": ", runTime, "| is sorted: ", checkSorted(sres))
	return runTime

}

func test(MAXN int, ARRLEN int) {
	runtime.GOMAXPROCS(4)
	var t1, t2 int = 0, 0

	for i := 0; i < MAXN; i += 1 {
		arr := genArray(int(ARRLEN))

		e1 := run(arr, strconv.Itoa(i+1)+"_seq", &seqSorter{})
		t1 += e1
		e2 := run(arr, strconv.Itoa(i+1)+"_par", &parSorter{})
		t2 += e2

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
