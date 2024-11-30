package sort

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

func partition(arr []int, from, to int) (int, int) {
	index := utils.randFromTo(from, to)

	pivot := arr[index]
	lc, ec := 0, 0
	i := from
	for i < to {
		// fmt.Println("i", i)
		if arr[i] < pivot {
			arr[from+lc], arr[i] = arr[i], arr[from+lc] // i <--> first equal|greater element
			if ec != 0 {
				arr[from+lc+ec], arr[i] = arr[i], arr[from+lc+ec]
			}
			lc++
		} else if arr[i] == pivot {
			arr[from+lc+ec], arr[i] = arr[i], arr[from+lc+ec]
			ec++
		}
		i++
	}
	return from + lc, from + lc + ec
}

func seqQsort(arr []int) {
	_seqQsort(arr, 0, len(arr))
}
