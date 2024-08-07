package pathfinding

// selectOptimalPaths selects the best paths for multiple trains while avoiding conflicts
// Parameters:
//
//	allPaths: A slice of all possible paths, each path being a slice of station names
//	numTrains: The number of trains to schedule
//	start, end: The names of the start and end stations
//
// Returns:
//
//	A slice of selected paths, where each path is a slice of station names
func selectOptimalPaths(allPaths [][]string, numTrains int, start, end string) [][]string {
	// Initialize slice to store selected paths and map to track occupied stations
	selectedPaths := make([][]string, 0, numTrains)
	occupiedStations := make(map[string]map[int]bool)

	// Helper function to check if a path conflicts with existing paths
	pathConflicts := func(path []string, startTime int) bool {
		for t, station := range path {
			if station != start && station != end {
				if occupied, exists := occupiedStations[station]; exists {
					if occupied[startTime+t] {
						return true // Conflict found
					}
				}
			}
		}
		return false // No conflicts
	}

	// Helper function to add a path to the selected paths
	addPath := func(path []string, startTime int) {
		// Create a new path with delay at the start if necessary
		delayedPath := make([]string, startTime+len(path))
		for i := 0; i < startTime; i++ {
			delayedPath[i] = start // Train waits at start station
		}
		copy(delayedPath[startTime:], path)

		// Add the path to selected paths
		selectedPaths = append(selectedPaths, delayedPath)

		// Mark stations as occupied for this path
		for t, station := range path {
			if station != start && station != end {
				if occupiedStations[station] == nil {
					occupiedStations[station] = make(map[int]bool)
				}
				occupiedStations[station][startTime+t] = true
			}
		}
	}

	// Find the maximum path length for the safety check
	maxPathLength := len(allPaths[len(allPaths)-1])
	timeStep := 0

	// Main loop to select paths
	for len(selectedPaths) < numTrains {
		for _, path := range allPaths {
			if len(selectedPaths) >= numTrains {
				break // Exit if we've selected enough paths
			}

			// Check for conflicts and add paths
			if !pathConflicts(path, timeStep) {
				if len(selectedPaths) != numTrains-1 {
					// Add path if it's not the last train
					addPath(path, timeStep)
				} else {
					// Special handling for the last train
					if len(allPaths) == 2 && len(allPaths[0])+1 < len(allPaths[1]) {
						// Choose shorter path with a delay if it's significantly shorter
						addPath(allPaths[0], timeStep+1)
					} else {
						addPath(path, timeStep)
					}
				}
			}
		}

		timeStep++
		// Safety check to prevent infinite loop
		if timeStep > maxPathLength*numTrains {
			break
		}
	}

	return selectedPaths
}
