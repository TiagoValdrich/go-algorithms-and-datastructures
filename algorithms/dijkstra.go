package algorithms

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

type WeightedGraph map[string]map[string]int

func getNextLowestCostNode(costs map[string]int, visited map[string]bool) string {
	lowestCost := math.MaxInt
	lowestCostNode := ""

	for node, nodeCost := range costs {
		isVisited := visited[node]

		if nodeCost < lowestCost && !isVisited {
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func getPath(parents map[string]string, start string, end string) string {
	path := []string{}
	parent := parents[end]

	if parent == "" {
		return ""
	}

	path = append(path, end)

	for parent != "" {
		path = append(path, parent)
		parent = parents[parent]
	}

	slices.Reverse(path)

	return strings.Join(path, ";")
}

func Dijkstra(weightedGraph WeightedGraph, start string, end string) string {
	costs := make(map[string]int)
	parents := make(map[string]string)
	visited := make(map[string]bool)
	queue := datastructures.NewQueue[string]()

	// Initialize costs and parents for the rest of the nodes
	for node := range weightedGraph {
		if node != start {
			costs[node] = math.MaxInt
			parents[node] = ""
		}
	}

	// Initialize costs and parents
	for node, cost := range weightedGraph[start] {
		costs[node] = cost
		parents[node] = start
	}

	lowestCostNode := getNextLowestCostNode(costs, visited)
	queue.Enqueue(lowestCostNode)

	for queue.Size() > 0 {
		node := queue.Dequeue()

		if node == "" {
			continue
		}

		fmt.Println(node)

		nodeCost := costs[node]
		nodeNeighbors := weightedGraph[node]

		for neighborNode, neighborNodeCost := range nodeNeighbors {
			newCost := neighborNodeCost + nodeCost
			currentCost := costs[neighborNode]

			if newCost < currentCost {
				costs[neighborNode] = newCost
				parents[neighborNode] = node
			}
		}

		visited[node] = true
		lowestCostNode = getNextLowestCostNode(costs, visited)
		queue.Enqueue(lowestCostNode)
	}

	return getPath(parents, start, end)
}
