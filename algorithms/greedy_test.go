package algorithms_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/algorithms"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

func TestGreedy(t *testing.T) {
	t.Run("greedy algorithm should return the best radio station list that will cover all the states", func(t *testing.T) {
		states := []string{"rs", "sc", "pr", "sp", "rj", "es", "mt", "bh", "mg"}
		radioStations := map[string]datastructures.ISet[string]{}

		radioOne := datastructures.NewSet[string]()
		radioOneStates := []string{"rs", "sc", "pr"}

		for _, state := range radioOneStates {
			radioOne.Add(state)
		}

		radioStations["one"] = radioOne

		radioTwo := datastructures.NewSet[string]()
		radioTwoStates := []string{"sp", "rj", "es"}

		for _, state := range radioTwoStates {
			radioTwo.Add(state)
		}

		radioStations["two"] = radioTwo

		radioThree := datastructures.NewSet[string]()
		radioThreeStates := []string{"sp", "mt", "mg"}

		for _, state := range radioThreeStates {
			radioThree.Add(state)
		}

		radioStations["three"] = radioThree

		radioFour := datastructures.NewSet[string]()
		radioFourStates := []string{"bh", "mt"}

		for _, state := range radioFourStates {
			radioFour.Add(state)
		}

		radioStations["four"] = radioFour

		radioFive := datastructures.NewSet[string]()
		radioFiveStates := []string{"bh", "mg"}

		for _, state := range radioFiveStates {
			radioFive.Add(state)
		}

		radioStations["five"] = radioFive

		result := algorithms.Greedy(states, radioStations)

		t.Log("Greedy result", strings.Join(result, ";"))

		assert.ElementsMatch(t, result, []string{"one", "two", "three", "four"})
	})
}
