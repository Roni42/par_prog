package main

import (
	"sync"
)

func ParBFS(graph Graph, start int, result []int) []int {
	var (
		visited = make([]bool, len(graph))
		queue   = []int{start}
		workers = 4
		next    []int
		mtx     sync.Mutex
		wg      sync.WaitGroup
	)

	visited[start] = true

	for len(queue) > 0 {
		chunk := (len(queue) + workers - 1) / workers
		next = []int{}

		wg.Add(workers)
		for w := 0; w < workers; w++ {
			go func(worker int) {
				defer wg.Done()
				localNext := []int{}

				st := worker * chunk
				end := min((worker+1)*chunk, len(queue))

				if st >= end {
					return
				}

				for _, node := range queue[st:end] {
					for _, neighbor := range graph[node] {
						if !visited[neighbor] {
							mtx.Lock()
							if !visited[neighbor] {
								visited[neighbor] = true
								result[neighbor] = result[node] + 1
								localNext = append(localNext, neighbor)
							}
							mtx.Unlock()
						}
					}
				}

				mtx.Lock()
				next = append(next, localNext...)
				mtx.Unlock()
			}(w)
		}

		wg.Wait()
		queue = next
	}

	return result
}
