package main

import (
	"fmt"
)

func main() {
	arr := utils.genArray(100)
	// fmt.Printf("array before qsort: %v\n", arr)
	sort.seqQsort(arr)
	// fmt.Printf("array after qsort: %v", arr)
	fmt.Printf("is sorted: %v", utils.checkSorted(arr))
}
