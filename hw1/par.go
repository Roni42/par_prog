package main

func _parQsort(_source []int, ch chan int) {

	// source := make([]int, len(_source), len(_source))
	// copy(source, _source)

	l := len(_source)
	if len(_source) >= 1 {
		p1, p2 := partition(_source, 0, l)
		less := make([]int, p1)
		copy(less, _source[:p1])
		great := make([]int, l-p2)
		copy(great, _source[p2:])

		ch1 := make(chan int, p1)
		ch2 := make(chan int, l-p2)

		if p1 > 1 {
			go _parQsort(less, ch1)
			println("panic ", p1, l)
			for i := range ch1 {
				ch <- i
			}
		} else if p1 == 1 {
			ch <- _source[0]
		}

		for range p2 - p1 {
			ch <- _source[p1]
		}

		if l-p2 > 1 {
			go _parQsort(great, ch2)
			for i := range ch2 {
				ch <- i
			}
		} else if l-p2 == 1 {
			ch <- _source[p1]
		}
	}

	close(ch)
}

func parQsort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	ch := make(chan int)
	_parQsort(res, ch)
	return res
}
