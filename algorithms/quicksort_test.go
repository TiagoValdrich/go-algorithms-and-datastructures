package algorithms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/algorithms"
)

func TestQuick(t *testing.T) {
	t.Run("Quick should return a sorted list", func(t *testing.T) {
		nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

		sortedList := algorithms.Quicksort(nums)

		expectedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		assert.Equal(t, expectedList, sortedList)
	})
}
