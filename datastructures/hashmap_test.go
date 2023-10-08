package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

func TestHashMap(t *testing.T) {
	hashmap := datastructures.NewHashMap[string, int]()

	t.Run("HashMap should store values", func(t *testing.T) {
		hashmap.Set("Tiago", 23)
		hashmap.Set("Alfredo", 21)
		hashmap.Set("Gabriela", 20)
		hashmap.Set("Nathalia", 24)
	})

	t.Run("HashMap should allow get stored values", func(t *testing.T) {
		tiagoValue := hashmap.Get("Tiago")
		alfredoValue := hashmap.Get("Alfredo")
		gabrielaValue := hashmap.Get("Gabriela")
		nathaliaValue := hashmap.Get("Nathalia")

		assert.Equal(t, tiagoValue, 23)
		assert.Equal(t, alfredoValue, 21)
		assert.Equal(t, gabrielaValue, 20)
		assert.Equal(t, nathaliaValue, 24)
	})

	t.Run("HashMap should return the it's size", func(t *testing.T) {
		size := hashmap.Size()

		assert.Equal(t, size, 4)
	})

	t.Run("HashMap should return a list containing stored keys", func(t *testing.T) {
		keys := hashmap.Keys()

		assert.Contains(t, keys, "Tiago")
		assert.Contains(t, keys, "Alfredo")
		assert.Contains(t, keys, "Gabriela")
		assert.Contains(t, keys, "Nathalia")
	})

	t.Run("HashMap should return a list containing stored values", func(t *testing.T) {
		values := hashmap.Values()

		assert.Contains(t, values, 23)
		assert.Contains(t, values, 21)
		assert.Contains(t, values, 20)
		assert.Contains(t, values, 24)
	})

	t.Run("HashMap should allow iterate over it's values", func(t *testing.T) {
		var keysFromIteration []string
		var valuesFromIteration []int

		for key, value := range hashmap.Iter() {
			keysFromIteration = append(keysFromIteration, key)
			valuesFromIteration = append(valuesFromIteration, value)
		}

		expectedKeys := []string{"Tiago", "Alfredo", "Gabriela", "Nathalia"}
		expectedValues := []int{23, 21, 20, 24}

		assert.ElementsMatch(t, keysFromIteration, expectedKeys)
		assert.ElementsMatch(t, valuesFromIteration, expectedValues)
	})

	t.Run("HashMap should delete stored values", func(t *testing.T) {
		hashmap.Delete("Tiago")
		hashmap.Delete("Alfredo")

		tiagoValue := hashmap.Get("Tiago")
		alfredoValue := hashmap.Get("Alfredo")

		assert.Equal(t, tiagoValue, 0)
		assert.Equal(t, alfredoValue, 0)
	})

	t.Run("HashMap should clear all stored values", func(t *testing.T) {
		gabrielaValue := hashmap.Get("Gabriela")
		nathaliaValue := hashmap.Get("Nathalia")

		assert.Equal(t, gabrielaValue, 20)
		assert.Equal(t, nathaliaValue, 24)
		assert.Equal(t, hashmap.Size(), 2)

		hashmap.Clear()

		gabrielaValue = hashmap.Get("Gabriela")
		nathaliaValue = hashmap.Get("Nathalia")

		assert.Equal(t, gabrielaValue, 0)
		assert.Equal(t, nathaliaValue, 0)
	})
}
