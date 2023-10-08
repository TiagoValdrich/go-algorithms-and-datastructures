package algorithms

import (
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

type BfsParams struct {
	Vertexes    map[string][]string
	StartFrom   string
	Destination string
}

// Breadth First Search implementation
// For learning purposes, it just accepts a map of string slices
func BFS(params BfsParams) []string {
	path := []string{}
	visited := make(map[string]bool)
	queue := datastructures.NewQueue[string]()

	queue.Enqueue(params.StartFrom)

	for queue.Size() > 0 {
		currentVertex := queue.Dequeue()
		visited[currentVertex] = true
		path = append(path, currentVertex)

		if currentVertex == params.Destination {
			return path
		}

		neighbors := params.Vertexes[currentVertex]

		for _, vertex := range neighbors {
			if _, alreadyVisited := visited[vertex]; !alreadyVisited {
				queue.Enqueue(vertex)
			}
		}
	}

	return []string{}
}
