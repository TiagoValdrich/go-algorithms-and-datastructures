package algorithms_test

import (
	"testing"

	"github.com/tiagovaldrich/go-algorithms-and-datastructures/algorithms"
)

func TestBinarySearch(t *testing.T) {
	t.Run("BinarySearch should return the index of the searched value", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		index := algorithms.BinarySearch(array, 4)

		if index != 3 {
			t.Errorf("BinarySearch(array, 4) returned %d, expected 3", index)
		}
	})

	t.Run("BinarySearch should return -1 if the searched value is not present in the array", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		index := algorithms.BinarySearch(array, 11)

		if index != -1 {
			t.Errorf("BinarySearch(array, 11) returned %d, expected -1", index)
		}
	})
}
