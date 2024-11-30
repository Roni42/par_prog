package utils

import (
	"fmt"
	"math/rand/v2"
)

func genArray(n int) []int {
	a := make([]int, n)
	for i := range n {
		a[i] = rand.N[int](n * 2)
	}
	return a
}

func printArr(source []int, from, to int, needPrev bool) {
	if needPrev {
		fmt.Printf("\n= array from %d to %d:\n", from, to)
	}
	for i := from; i < to; i += 1 {
		fmt.Printf("%d ", source[i])
	}
	fmt.Println("")
}

func randFromTo(from, to int) int {
	return from + rand.N[int](to-from)
}

func checkSorted(arr []int) bool {
	c := arr[0]
	for _, a := range arr {
		if c > a {
			return false
		}
		c = a
	}
	return true
}
