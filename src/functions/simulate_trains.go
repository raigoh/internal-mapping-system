package functions

import (
	"fmt"
)

// trainsPaths represents the state and route of a train
type trainsPaths struct {
	TrainID   int
	Route     []string
	RouteStep int
	Finished  bool
	Skip      bool
}

// SimulateTrains2 simulates train movements along predefined paths
func SimulateTrains2(paths [][]string, numTrains int) {
	var trains []trainsPaths
	var tempRoute []string
	routeIndex := 0

	shortestPathLen := int(^uint(0) >> 1) // Max int value
	longestPathLen := 0
	doubleTheShort := false

	// Determine the shortest and longest paths
	for _, route := range paths {
		if len(route) < shortestPathLen {
			shortestPathLen = len(route)
		}
		if len(route) > longestPathLen {
			longestPathLen = len(route)
		}
	}

	if shortestPathLen*2 <= longestPathLen {
		doubleTheShort = true
	}

	// Initialize trains with routes
	for i := 0; i < numTrains; i++ {
		if routeIndex >= len(paths) {
			routeIndex = 0
		}

		skipTurn := false
		if i == numTrains-1 && doubleTheShort {
			shortLen := 0
			skipTurn = true
			for _, route := range paths {
				if len(route) < shortLen || shortLen == 0 {
					shortLen = len(route)
					tempRoute = route
				}
			}
		} else {
			tempRoute = paths[routeIndex]
		}

		trains = append(trains, trainsPaths{TrainID: i, Route: tempRoute, RouteStep: 0, Finished: false, Skip: skipTurn})
		routeIndex++
	}

	// Prepare paths for printTestOutput
	simulatedPaths := make([][]string, numTrains)
	for i := range simulatedPaths {
		simulatedPaths[i] = make([]string, 0)
	}

	// Simulate train movements
	allFinished := false
	for step := 0; !allFinished; step++ {
		checkLimit := (step+1)*len(paths) - 1
		if checkLimit >= numTrains {
			checkLimit = numTrains - 1
		}

		for t := 0; t <= checkLimit; t++ {
			train := &trains[t]
			if !train.Finished {
				if !train.Skip {
					simulatedPaths[t] = append(simulatedPaths[t], train.Route[train.RouteStep])
					train.RouteStep++
				} else {
					train.Skip = false
					simulatedPaths[t] = append(simulatedPaths[t], train.Route[train.RouteStep])
				}
				if train.RouteStep >= len(train.Route)-1 {
					train.Finished = true
				}
			} else {
				// Add the last station again for finished trains
				simulatedPaths[t] = append(simulatedPaths[t], train.Route[len(train.Route)-1])
			}
		}

		// Check if all trains have finished
		allFinished = true
		for _, train := range trains {
			if !train.Finished {
				allFinished = false
				break
			}
		}
	}

	// Print the simulated paths using the new function
	printTestOutput(simulatedPaths)
}

// printTestOutput prints the test output in a more readable format
func printTestOutput(paths [][]string) {
	maxLen := 0
	for _, path := range paths {
		if len(path) > maxLen {
			maxLen = len(path)
		}
	}

	for step := 0; step < maxLen; step++ {
		changes := false
		line := fmt.Sprintf("Step %d: ", step)
		for trainID, path := range paths {
			if step < len(path) && (step == 0 || path[step] != path[step-1]) {
				line += fmt.Sprintf("T%d-%s ", trainID+1, path[step])
				changes = true
			}
		}
		if changes {
			fmt.Println(line)
		}
	}
}
