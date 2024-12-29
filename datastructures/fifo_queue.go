package datastructures

import "sync"

type IQueue[T any] interface {
	Enqueue(value T)
	Dequeue() T
	Size() int
}

// Thread safe queue implementation of IFifoQueue[T]
type FifoQueue[T any] struct {
	values []T
	size   int
	mutex  sync.Mutex
}

func NewFifoQueue[T any]() IQueue[T] {
	return &FifoQueue[T]{
		values: make([]T, 0),
		size:   0,
	}
}

func (queue *FifoQueue[T]) Enqueue(value T) {
	queue.mutex.Lock()
	queue.values = append(queue.values, value)
	queue.size++
	queue.mutex.Unlock()
}

func (queue *FifoQueue[T]) Dequeue() T {
	if queue.Size() == 0 {
		var val T
		return val
	}

	queue.mutex.Lock()
	value := queue.values[0]
	queue.values = queue.values[1:]
	queue.size--
	queue.mutex.Unlock()

	return value
}

func (queue *FifoQueue[T]) Size() int {
	queue.mutex.Lock()
	currentSize := queue.size
	queue.mutex.Unlock()

	return currentSize
}
