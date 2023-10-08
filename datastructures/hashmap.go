package datastructures

import "sync"

type IHashMap[Key comparable, Value any] interface {
	Keys() []Key
	Values() []Value
	Get(key Key) Value
	Set(key Key, val Value)
	Delete(key Key)
	Iter() map[Key]Value
	Size() int
	Clear()
}

// Thread safe hashmap implementation of IHashMap[Key, Value]
type HashMap[Key comparable, Value any] struct {
	values map[Key]Value
	size   int
	mutex  sync.Mutex
}

func NewHashMap[Key comparable, Value any]() IHashMap[Key, Value] {
	return &HashMap[Key, Value]{
		values: make(map[Key]Value),
		size:   0,
	}
}

func (hashMap *HashMap[Key, Value]) Keys() []Key {
	var keys []Key

	hashMap.mutex.Lock()

	for key := range hashMap.values {
		keys = append(keys, key)
	}

	hashMap.mutex.Unlock()

	return keys
}

func (hashMap *HashMap[Key, Value]) Values() []Value {
	var values []Value

	hashMap.mutex.Lock()

	for _, val := range hashMap.values {
		values = append(values, val)
	}

	hashMap.mutex.Unlock()

	return values
}

func (hashMap *HashMap[Key, Value]) Get(key Key) Value {
	var value Value

	hashMap.mutex.Lock()

	if val, exists := hashMap.values[key]; exists {
		value = val
	}

	hashMap.mutex.Unlock()

	return value
}

func (hashMap *HashMap[Key, Value]) Set(key Key, val Value) {
	hashMap.mutex.Lock()
	hashMap.values[key] = val
	hashMap.size++
	hashMap.mutex.Unlock()
}

func (hashMap *HashMap[Key, Value]) Delete(key Key) {
	hashMap.mutex.Lock()
	delete(hashMap.values, key)
	hashMap.size--
	hashMap.mutex.Unlock()
}

func (hashMap *HashMap[Key, Value]) Clear() {
	hashMap.mutex.Lock()
	hashMap.values = make(map[Key]Value)
	hashMap.size = 0
	hashMap.mutex.Unlock()
}

func (hashMap *HashMap[Key, Value]) Iter() map[Key]Value {
	hashMap.mutex.Lock()
	values := hashMap.values
	hashMap.mutex.Unlock()

	return values
}

func (hashMap *HashMap[Key, Value]) Size() int {
	hashMap.mutex.Lock()
	size := hashMap.size
	hashMap.mutex.Unlock()

	return size
}
