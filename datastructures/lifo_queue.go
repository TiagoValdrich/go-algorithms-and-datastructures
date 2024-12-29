package datastructures

import "sync"

// Thread safe queue implementation of IQueue[T]
type LifoQueue[T any] struct {
	values []T
	size   int
	mutex  sync.Mutex
}

func NewLifoQueue[T any]() IQueue[T] {
	return &LifoQueue[T]{
		values: make([]T, 0),
		size:   0,
	}
}

func (queue *LifoQueue[T]) Enqueue(value T) {
	queue.mutex.Lock()
	queue.values = append(queue.values, value)
	queue.size++
	queue.mutex.Unlock()
}

func (queue *LifoQueue[T]) Dequeue() T {
	if queue.Size() == 0 {
		var val T
		return val
	}

	queue.mutex.Lock()
	lastPosition := queue.size - 1
	value := queue.values[lastPosition]
	queue.values = queue.values[:lastPosition]
	queue.size--
	queue.mutex.Unlock()

	return value
}

func (queue *LifoQueue[T]) Size() int {
	queue.mutex.Lock()
	currentSize := queue.size
	queue.mutex.Unlock()

	return currentSize
}
