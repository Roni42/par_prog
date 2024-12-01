package main

func _seqQsort(source []int, from, to int) {
	if from < to {
		p1, p2 := partition(source, from, to)
		if p1-from > 1 {
			_seqQsort(source, from, p1)
		}
		if to-p2 > 1 {
			_seqQsort(source, p2, to)
		}
	}
}

func seqQsort(arr []int) []int {
	res := make([]int, len(arr), len(arr))
	copy(res, arr)
	_seqQsort(res, 0, len(arr))
	return res
}
