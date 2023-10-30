package algorithms

import "github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"

// Greedy algorithm to find the "ideal" states for a radio station to cover all states
func Greedy(states []string, radioStations map[string]datastructures.ISet[string]) []string {
	availableStates := datastructures.NewSet[string]()
	finalStations := make([]string, 0)

	for _, state := range states {
		availableStates.Add(state)
	}

	for availableStates.Size() > 0 {
		bestStation := ""
		currentCoveredStates := datastructures.NewSet[string]()

		for station, states := range radioStations {
			coveredStates := datastructures.NewSet[string]()

			for _, state := range states.Values() {
				if availableStates.Has(state) {
					coveredStates.Add(state)
				}
			}

			if coveredStates.Size() > currentCoveredStates.Size() {
				bestStation = station
				currentCoveredStates = coveredStates
			}
		}

		for _, state := range currentCoveredStates.Values() {
			availableStates.Remove(state)
		}

		finalStations = append(finalStations, bestStation)
	}

	return finalStations
}
