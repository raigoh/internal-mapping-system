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

	// Simulate train movements
	allFinished := false
	for step := 1; !allFinished; step++ {
		checkLimit := step*len(paths) - 1
		if checkLimit >= numTrains {
			checkLimit = numTrains - 1
		}

		for t := 0; t <= checkLimit; t++ {
			train := &trains[t]
			if !train.Finished {
				if !train.Skip {
					train.RouteStep++
					fmt.Print("T", train.TrainID+1, "-", train.Route[train.RouteStep], " ")
				} else {
					train.Skip = false
				}
				if train.RouteStep >= len(train.Route)-1 {
					train.Finished = true
				}
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

		fmt.Println("")
	}
}
