package functions

import (
	"fmt"
	"strings"
)

// New struct for trains
type trainsPaths struct {
	TrainID   int
	Route     []string
	RouteStep int
	Finnished bool
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
	temp := 0

	//Make every train a part of a slice as a trainsPath structure. That way, every train can have their own route dedicated to them
	for x := range numTrains {
		if temp >= len(paths) {
			temp = 0
		}
		trains = append(trains, trainsPaths{TrainID: x, Route: paths[temp], RouteStep: 0, Finnished: false})

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
				trains[t].RouteStep++
				fmt.Print("T", t+1, "-", trains[t].Route[trains[t].RouteStep], " ")

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
		// This just marks the end of a turn, by starting a new line
		fmt.Println("")
	}
}
