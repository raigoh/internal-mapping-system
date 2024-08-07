package main

import (
	"fmt"
	"os"
	"path/filepath"
	"station/internal/core"
	"station/internal/io"
	"station/internal/pathfinding"
	"station/internal/utils"
	"strconv"
)

func main() {
	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		utils.PrintUsage()
		return
	}

	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "%sError: Incorrect number of arguments%s\n", utils.Red, utils.Reset)
		utils.PrintUsage()
		return
	}

	// Extract command-line arguments
	networkMapArg := os.Args[1] // Path to the network map file as provided in arguments
	startStationName := os.Args[2]
	endStationName := os.Args[3]

	// Parse the number of trains, ensuring it's a positive integer
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil || numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "%sError: number_of_trains must be a valid positive integer%s\n", utils.Red, utils.Reset)
		return
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError: Unable to get current working directory: %v%s\n", utils.Red, err, utils.Reset)
		return
	}

	// Construct the full path to the network map file
	networkMap := filepath.Join(cwd, networkMapArg)

	// If the file doesn't exist, try looking in the stations folder at the project root
	if _, err := os.Stat(networkMap); os.IsNotExist(err) {
		projectRoot := filepath.Dir(filepath.Dir(cwd)) // Assuming cmd is one level deep in the project structure
		networkMap = filepath.Join(projectRoot, "stations", networkMapArg)
	}

	// Read and parse the network map
	networks, err := io.ReadMap(networkMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError: 1. Reading network map:\n       2. %v%s\n",
			utils.Red,
			err,
			utils.Reset)
		return
	}

	// Find the appropriate map containing both start and end stations
	_, selectedNetwork, err := core.FindAppropriateMap(networks, startStationName, endStationName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s%v%s\n", utils.Red, err, utils.Reset)
		return
	}

	// Verify that the start and end stations are different
	if startStationName == endStationName {
		fmt.Fprintf(os.Stderr, "%sError: Start and end station are the same%s\n", utils.Red, utils.Reset)
		return
	}

	// Find paths for the trains
	paths, occupations, err := pathfinding.FindPaths(startStationName, endStationName, selectedNetwork, numTrains)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError finding paths: %v%s\n", utils.Red, err, utils.Reset)
		return
	}

	// Note: occupations data is available here but not used in this version
	_ = occupations

	// Simulate and display the train movements
	pathfinding.SimTrain(paths)
}
