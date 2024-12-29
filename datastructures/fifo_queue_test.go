package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

func TestQueue(t *testing.T) {
	queue := datastructures.NewFifoQueue[int]()

	t.Run("Queue should store values in FIFO order", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			queue.Enqueue(i)
		}

		assert.Equal(t, 5, queue.Size())
	})

	t.Run("Queue should return values in FIFO order", func(t *testing.T) {
		size := queue.Size()

		for i := 1; i <= size; i++ {
			value := queue.Dequeue()
			assert.Equal(t, i, value)
		}
	})

	t.Run("Queue should return it's size", func(t *testing.T) {
		assert.Equal(t, queue.Size(), 0)

		for i := 1; i <= 3; i++ {
			queue.Enqueue(i)
		}

		assert.Equal(t, 3, queue.Size())
	})
}
