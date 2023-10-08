package algorithms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/algorithms"
)

func TestBFS(t *testing.T) {
	t.Run("BFS should return the shortest path between two nodes", func(t *testing.T) {
		graph := map[string][]string{
			"ben":     {"alf", "rim"},
			"rim":     {"john"},
			"alf":     {"edward"},
			"john":    {"edward"},
			"edward":  {"jim", "bjorn"},
			"jim":     {"joe"},
			"joe":     {"catarina"},
			"bjorn":   {"romenia", "irenia"},
			"romenia": {"catarina"},
			"irenia":  {"catarina"},
		}
		destination := "catarina"
		startFrom := "ben"

		path := algorithms.BFS(algorithms.BfsParams{
			Vertexes:    graph,
			StartFrom:   startFrom,
			Destination: destination,
		})

		expectedPath := []string{"ben", "alf", "rim", "edward", "john", "jim", "bjorn", "joe", "romenia", "irenia", "catarina"}

		assert.Equal(t, expectedPath, path)
	})
}
