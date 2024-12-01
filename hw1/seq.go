package main

type seqSorter struct {
	qsortBase
}

func (s *seqSorter) sort(arr []int, from, to int) {
	if from < to {
		p1, p2 := s.partition(arr, from, to)
		if p1-from > 1 {
			s.sort(arr, from, p1)
		}
		if to-p2 > 1 {
			s.sort(arr, p2, to)
		}
	}
}

func (s *seqSorter) Sort(arr []int) {
	s.sort(arr, 0, len(arr))
}

func (s *seqSorter) SortImm(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	s.Sort(res)
	return res
}
