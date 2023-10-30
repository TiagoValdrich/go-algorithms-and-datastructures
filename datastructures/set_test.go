package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

func TestSet(t *testing.T) {
	s := datastructures.NewSet[string]()

	t.Run("should be able to add items to the set", func(t *testing.T) {
		s.Add("a")
		s.Add("b")
		s.Add("c")

		assert.Equal(t, 3, s.Size())
	})

	t.Run("should be able to remove items from the set", func(t *testing.T) {
		s.Remove("a")

		currentValues := s.Values()

		assert.ElementsMatch(t, []string{"b", "c"}, currentValues)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("should be able to check if an item exists in the set", func(t *testing.T) {
		setHasB := s.Has("b")

		assert.Equal(t, true, setHasB)
	})

	t.Run("should be able to get all values from the set", func(t *testing.T) {
		values := s.Values()

		assert.ElementsMatch(t, []string{"b", "c"}, values)
	})
}
