package main

import (
	"fmt"
	"os"
	"station/src/colors"
	"station/src/data"
	"station/src/functions"
	"strconv"
)

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "%sError: Incorrect number of arguments%s\n", colors.Red, colors.Reset)
		return
	}

	// Extract command-line arguments
	networkMap := os.Args[1]       // Path to the network map file
	startStationName := os.Args[2] // Name of the start station
	endStationName := os.Args[3]   // Name of the end station

	// Parse the number of trains, ensuring it's a positive integer
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil || numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "%sError: number_of_trains must be a valid positive integer%s\n", colors.Red, colors.Reset)
		return
	}

	// Read and parse the network map
	networks, err := functions.ReadMap(networkMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError: 1. Reading network map:\n       2. %v%s\n",
			functions.ColorRed,
			err,
			functions.ColorReset)
		return
	}

	// Find the network that contains both the start and end stations
	var selectedNetwork map[string]*data.Station
	for _, network := range networks {
		if _, startExists := network[startStationName]; startExists {
			if _, endExists := network[endStationName]; endExists {
				selectedNetwork = network
				break
			}
		}
	}

	if selectedNetwork == nil {
		fmt.Fprintf(os.Stderr, "%sError: No network contains both %s and %s%s\n", functions.ColorRed, startStationName, endStationName, functions.ColorReset)
		return
	}

	// Verify that the start and end stations are different
	if startStationName == endStationName {
		fmt.Fprintf(os.Stderr, "%sError: Start and end station are the same%s\n", functions.ColorRed, functions.ColorReset)
		return

	}

	// Find paths for the trains
	paths, occupations, err := functions.FindPaths(startStationName, endStationName, selectedNetwork, numTrains)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError finding paths: %v%s\n", functions.ColorRed, err, functions.ColorReset)
		return
	}

	// Note: occupations data is available here but not used in this version
	_ = occupations

	// Simulate and display the train movements
	functions.SimTrain(paths)

}
