package main

import (
	"bytes"
	"fmt"
	"os/exec"
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

// Colors for output
const (
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorReset = "\033[0m"
)

func runCommand(mapFile, start, end string, numTrains int) (string, error) {
	cmd := exec.Command("go", "run", ".", mapFile, start, end, fmt.Sprint(numTrains))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestPathfinder(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.startStation, tc.endStation), func(t *testing.T) {
			output, err := runCommand(tc.mapFile, tc.startStation, tc.endStation, tc.numberOfTrains)
			if err != nil {
				t.Errorf("%sFailed to run command: %v%s", ColorRed, err, ColorReset)
				return
			}

			lines := strings.Split(strings.TrimSpace(output), "\n")
			actualTurns := len(lines)

			if actualTurns != tc.expectedTurns {
				t.Errorf("%sWanted minimum %d turns, got %d turns.\nOutput:\n%s%s", ColorRed, tc.expectedTurns, actualTurns, output, ColorReset)
			} else {
				fmt.Printf("%sTest passed for %s to %s with %d trains.%s\nOutput:\n%s%s", ColorGreen, tc.startStation, tc.endStation, tc.numberOfTrains, ColorGreen, output, ColorReset)
			}
		})
	}
}
