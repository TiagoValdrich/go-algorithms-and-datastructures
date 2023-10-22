package algorithms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/algorithms"
)

func TestDijkstra(t *testing.T) {
	weightedGraph := algorithms.WeightedGraph{
		"start": map[string]int{
			"a": 5,
			"b": 2,
		},
		"a": map[string]int{
			"c": 4,
			"d": 2,
		},
		"b": map[string]int{
			"a": 8,
			"d": 7,
		},
		"c": map[string]int{
			"d":   6,
			"end": 3,
		},
		"d": map[string]int{
			"end": 1,
		},
		"end": map[string]int{},
	}

	result := algorithms.Dijkstra(weightedGraph, "start", "end")

	assert.Equal(t, "start;a;d;end", result)
}
