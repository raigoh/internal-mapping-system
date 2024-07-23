package main

import (
	"fmt"
	"os"
	"station/src/functions"
	"strconv"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error: Incorrect number of arguments")
		return
	}

	networkMap := os.Args[1]
	startStationName := os.Args[2]
	endStationName := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil || numTrains <= 0 {
		fmt.Fprintln(os.Stderr, "Error: number_of_trains must be a valid positive integer")
		return
	}

	stations, err := functions.ReadMap(networkMap)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading network map:", err)
		return
	}

	if _, exists := stations[startStationName]; !exists {
		fmt.Fprintf(os.Stderr, "Error: Start station %s does not exist\n", startStationName)
		return
	}

	if _, exists := stations[endStationName]; !exists {
		fmt.Fprintf(os.Stderr, "Error: End station %s does not exist\n", endStationName)
		return
	}

	if startStationName == endStationName {
		fmt.Fprintln(os.Stderr, "Error: Start and end station are the same")
		return
	}

	paths, err := functions.FindPaths(startStationName, endStationName, stations, numTrains)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error finding paths:", err)
		return
	}
	functions.SimulateTrains2(paths, numTrains)
}
