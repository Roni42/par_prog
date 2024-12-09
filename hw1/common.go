package main

type sorter interface {
	partition(arr []int, from, to int) (int, int)
	Sort(arr []int)
	SortImm(arr []int) []int
}

type qsortBase struct{}

func (*qsortBase) partition(arr []int, from, to int) (int, int) {
	index := randFromTo(from, to)

	pivot := arr[index]
	lc, ec := 0, 0
	i := from
	for i < to {
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
