package main

func SeqBFS(graph Graph, start int, result []int) []int {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true
	for len(queue) > 0 {
		level := queue
		queue = nil

		for _, node := range level {
			for _, neigh := range graph[node] {
				if !visited[neigh] {
					result[neigh] = result[node] + 1
					visited[neigh] = true
					queue = append(queue, neigh)
				}
			}
		}
	}
	return result
}
