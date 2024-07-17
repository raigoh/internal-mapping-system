package functions

import (
	"fmt"
	"strings"
	"time"
)

// New struct for trains
type trainsPaths struct {
	TrainID   int
	Route     []string
	RouteStep int
	Finnished bool
	Skip      bool
}

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

// SimulateTrains simulates train movements and outputs the movements to the terminal
func SimulateTrains2(paths [][]string, numTrains int) {

	//Store all trains in a slice called trains
	var trains []trainsPaths
	var tempRoute []string
	temp := 0

	shortestPathLen := 99999999
	longestPathLen := 0
	doubleTheShort := false

	//Find the shortest and longest path in the paths string
	for _, r := range paths {

		if shortestPathLen > len(r) {
			shortestPathLen = len(r)
		}
		if longestPathLen < len(r) {
			longestPathLen = len(r)
		}
	}

	//fmt.Println("Shortest path: ", shortestPath, " with the length of ", shortestPathLen)
	//fmt.Println("Longest path: ", longestPath, " with the length of ", longestPathLen)

	if shortestPathLen*2 <= longestPathLen {
		//fmt.Println("Its better to have two shortest paths in a row than longer and shorter with the last trains")
		doubleTheShort = true
	}
	//Make every train a part of a slice as a trainsPath structure. That way, every train can have their own route dedicated to them
	for x := range numTrains {
		if temp >= len(paths) {
			temp = 0
		}

		skipTurn := false

		if x == numTrains-1 && doubleTheShort {
			//Check the shortest path
			shortLen := 0
			skipTurn = true
			for _, str := range paths {
				if len(str) <= shortLen || shortLen == 0 {
					shortLen = len(str)
					tempRoute = str
					//fmt.Println("Shortest: ", tempRoute)
					time.Sleep(1 * time.Second)
				}
			}

		} else {
			tempRoute = paths[temp]
			//fmt.Println(tempRoute)
		}

		trains = append(trains, trainsPaths{TrainID: x, Route: tempRoute, RouteStep: 0, Finnished: false, Skip: skipTurn})

		//Temp chooses the route for every train
		temp++
	}

	allFinnished := false

	for step := 1; !allFinnished; step++ {

		check := step*len(paths) - 1

		if check >= numTrains {
			check = numTrains - 1
		}

		// Go throu the steps.. t = step
		for t := 0; t <= check; t++ {

			// If train has finnished, skip it
			if !trains[t].Finnished {

				// increase route step
				if !trains[t].Skip {
					trains[t].RouteStep++
					fmt.Print("T", t+1, "-", trains[t].Route[trains[t].RouteStep], " ")
				} else {
					trains[t].Skip = false
				}
				// Check if the train reached the last station
				if trains[t].RouteStep >= len(trains[t].Route)-1 {
					trains[t].Finnished = true
				}
			}

			// Check if all trains have finnished and break from the for loop
			allFinnished = true
			for _, t := range trains {
				if !t.Finnished {
					allFinnished = false
				}
			}
			if allFinnished {
				break
			}

		}
		// This just marks the end of a turn, by starting a new turn
		fmt.Println("")
	}
}
