package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"station/internal/utils"
	"strings"
	"testing"
)

// Test cases with format: {map file, start station, end station, number of trains, expected number of turns}
var testCases = []struct {
	mapFile        string
	startStation   string
	endStation     string
	numberOfTrains int
	expectedTurns  int
}{
	{"network.map", "waterloo", "st_pancras", 4, 3},
	{"network.map", "beginning", "terminus", 20, 11},
	{"network.map", "beethoven", "part", 9, 6},
	{"network.map", "small", "large", 9, 8},
	{"network.map", "two", "four", 4, 6},
	{"network.map", "jungle", "desert", 10, 8},
	{"network.map", "bond_square", "space_port", 4, 6},
}

func runCommand(mapFile, start, end string, numTrains int) (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Construct the path to the main.go file (in the cmd folder)
	mainPath := filepath.Join(filepath.Dir(cwd), "cmd", "main.go")

	// Construct the full path to the map file (in the stations folder)
	mapPath := filepath.Join(filepath.Dir(cwd), mapFile)

	// Prepare the command to run the main program
	cmd := exec.Command("go", "run", mainPath, mapPath, start, end, fmt.Sprint(numTrains))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	return out.String(), err
}

func TestPathfinder(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.startStation, tc.endStation), func(t *testing.T) {
			output, err := runCommand(tc.mapFile, tc.startStation, tc.endStation, tc.numberOfTrains)

			if strings.Contains(output, "Error: Map contains more than 10000 stations") {
				t.Logf("%sMap contains more than 10000 stations, skipping test%s", utils.Red, utils.Reset)
				return
			}

			if err != nil {
				t.Errorf("%sFailed to run command: %v\nOutput: %s%s", utils.Red, err, output, utils.Reset)
				return
			}

			lines := strings.Split(strings.TrimSpace(output), "\n")
			actualTurns := len(lines)

			// Compare actual turns with expected turns
			if actualTurns != tc.expectedTurns {
				t.Errorf("%sWanted minimum %d turns, got %d turns.\nOutput:\n%s%s",
					utils.Red, tc.expectedTurns, actualTurns, output, utils.Reset)
			} else {
				// If test passes, print success message with green color
				fmt.Printf("%sTest passed for %s to %s with %d trains.%s\nOutput:\n%s%s",
					utils.Green, tc.startStation, tc.endStation, tc.numberOfTrains,
					utils.Green, output, utils.Reset)
			}
		})
	}
}
