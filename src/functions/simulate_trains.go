package functions

import (
	"fmt"
	"strings"
)

// SimulateTrains simulates train movements and outputs the movements to the terminal
func SimulateTrains(paths [][]string, numTrains int) {
	// Initialize train locations at the starting position
	trainLocations := make(map[string]int)
	for i := 0; i < numTrains; i++ {
		trainLocations[fmt.Sprintf("T%d", i+1)] = 0
	}

	// Simulation loop, runs until all trains have finished their routes
	for step := 0; ; step++ {
		move := []string{}  // Stores movements made in the current step
		finishedTrains := 0 // Tracks how many trains have finished their routes

		// Iterate over each train
		for i := 0; i < numTrains; i++ {
			trainName := fmt.Sprintf("T%d", i+1)

			// If there are no more paths for this train, mark it as finished
			if i >= len(paths) {
				finishedTrains++
				continue
			}
			currentPos := trainLocations[trainName]

			// If the train has reached the end of its path, mark it as finished
			if currentPos >= len(paths[i])-1 {
				finishedTrains++
				continue
			}

			// Determine the next station on the current train's path
			nextStation := paths[i][currentPos+1]
			move = append(move, fmt.Sprintf("%s-%s", trainName, nextStation))
			trainLocations[trainName]++
		}

		// If all trains have finished their routes, exit the simulation loop
		if finishedTrains == numTrains {
			break
		}

		// Output the movements made in the current step
		fmt.Println(strings.Join(move, " "))
	}
}
