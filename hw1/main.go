package main

import (
	"fmt"
	"runtime"
	"time"
)

var cycles = 5

func main() {
	runtime.GOMAXPROCS(4)

	t1, t2 := 0, 0

	for range cycles {
		arr := genArray(1e1)

		// s1 := time.Now()
		// res1 := seqQsort(arr)
		// t1 += int(time.Since(s1).Nanoseconds() / 1000000)
		// if checkSorted(arr) || !checkSorted(res1) {
		// 	panic("broken sort")
		// }

		s2 := time.Now()
		parQsort(arr)
		t2 += int(time.Since(s2).Nanoseconds() / 1000000)
		if !checkSorted(arr) {
			panic("broken sort")
		}

	}
	fmt.Println("result: ", t1/cycles, " vs ", t2/cycles)
}
