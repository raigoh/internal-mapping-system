package main

import (
	"flag"
	"fmt"
	"os"
	"station/internal/core"
	"station/internal/io"
	"station/internal/pathfinding"
	"station/internal/utils"
	"station/internal/visualization"
	"strconv"
)

func main() {
	var visualize bool
	var help bool
	flag.BoolVar(&visualize, "v", false, "Enable visualization")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Usage = func() {}

	flag.Parse()

	if help || (len(os.Args) > 1 && os.Args[1] == "--help") {
		utils.PrintUsage()
		return
	}

	if flag.NArg() != 4 {
		fmt.Fprintf(os.Stderr, "%s%s%s\n", utils.Red, utils.ErrIncorrectArgCount, utils.Reset)
		fmt.Fprintf(os.Stderr, "Use: 'go run main.go -h' for usage information\n")
		return
	}

	args := flag.Args()
	networkMapFile := args[0]
	startStationName := args[1]
	endStationName := args[2]

	numTrains, err := strconv.Atoi(args[3])
	if err != nil || numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "%s%s%s\n", utils.Red, utils.ErrInvalidTrainCount, utils.Reset)
		return
	}

	networks, err := io.ReadMap(networkMapFile)
	if err != nil {
		printError(err)
		return
	}

	_, selectedNetwork, err := core.FindAppropriateMap(networks, startStationName, endStationName)
	if err != nil {
		printError(err)
		return
	}

	if startStationName == endStationName {
		fmt.Fprintf(os.Stderr, "%s%s%s\n", utils.Red, utils.ErrSameStartEndStation, utils.Reset)
		return
	}

	paths, occupations, err := pathfinding.FindPaths(startStationName, endStationName, selectedNetwork, numTrains)
	if err != nil {
		printError(err)
		return
	}

	_ = occupations

	if visualize {
		err = visualization.CreateVisualization(selectedNetwork, paths)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%sError creating visualization: %v%s\n", utils.Red, err, utils.Reset)
		}
	}

	pathfinding.SimTrain(paths)
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s%s%s\n", utils.Red, err.Error(), utils.Reset)
	os.Exit(1)
}
