package pathfinding

import (
	"fmt"
	"strings"
)

// SimTrain simulates the movement of trains along their paths and prints the simulation results
// Parameters:
//
//	paths: A slice of paths, where each path is a slice of station names representing a train's route
func SimTrain(paths [][]string) {
	// Find the length of the longest path
	maxLen := 0
	for _, path := range paths {
		if len(path) > maxLen {
			maxLen = len(path) - 1 // Subtract 1 because we start from step 1, not 0
		}
	}

	// Simulate each step of the train movements
	for step := 1; step < maxLen+1; step++ {
		changes := false        // Flag to track if any train moved in this step
		movements := []string{} // Slice to store each train's movement

		// Check each train's position at this step
		for trainID, path := range paths {
			// Conditions for recording a train's position:
			// 1. The current step is within the train's path length
			// 2. It's either the first step (step == 0) or the train has moved to a new station
			if step < len(path) && (step == 0 || path[step] != path[step-1]) {
				movements = append(movements, fmt.Sprintf("T%d-%s", trainID+1, path[step]))
				changes = true // Mark that at least one train moved
			}
		}

		// Only print the line if there were any changes in this step
		if changes {
			fmt.Println(strings.Join(movements, " "))
		}
	}
}
