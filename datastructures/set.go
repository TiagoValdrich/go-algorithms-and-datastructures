package datastructures

type ISet[T comparable] interface {
	Add(value T)
	Remove(value T)
	Has(value T) bool
	Values() []T
	Size() int
}

type Set[T comparable] struct {
	values map[T]*struct{}
}

func NewSet[T comparable]() ISet[T] {
	return &Set[T]{
		values: make(map[T]*struct{}),
	}
}

func (set *Set[T]) Add(value T) {
	set.values[value] = nil
}

func (set *Set[T]) Remove(value T) {
	delete(set.values, value)
}

func (set *Set[T]) Has(value T) bool {
	_, exists := set.values[value]

	return exists
}

func (set *Set[T]) Values() []T {
	var values []T

	for val := range set.values {
		values = append(values, val)
	}

	return values
}

func (set *Set[T]) Size() int {
	return len(set.values)
}
