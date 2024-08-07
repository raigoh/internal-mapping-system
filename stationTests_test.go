package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// Test cases with format: {map file, start station, end station, number of trains, expected number of turns}
// Each test case represents a different scenario to test the pathfinding algorithm
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

// Colors for output - ANSI escape codes for colored console output
// These constants are used to color the output in the terminal for better readability
const (
	ColorRed   = "\033[31m" // Used for error messages
	ColorGreen = "\033[32m" // Used for success messages
	ColorReset = "\033[0m"  // Resets the color back to default
)

// runCommand executes the main program with given parameters and returns its output
// Parameters:
//
//	mapFile: path to the network map file
//	start: name of the start station
//	end: name of the end station
//	numTrains: number of trains to route
//
// Returns:
//
//	output of the command (string) and any error encountered
func runCommand(mapFile, start, end string, numTrains int) (string, error) {
	// Prepare the command to run the main program
	cmd := exec.Command("go", "run", ".", mapFile, start, end, fmt.Sprint(numTrains))
	var out bytes.Buffer
	cmd.Stdout = &out // Capture standard output
	cmd.Stderr = &out // Capture standard error
	err := cmd.Run()  // Run the command
	return out.String(), err
}

// TestPathfinder is the main test function that runs all test cases
func TestPathfinder(t *testing.T) {
	for _, tc := range testCases {
		// Create a subtest for each test case
		t.Run(fmt.Sprintf("%s_%s", tc.startStation, tc.endStation), func(t *testing.T) {
			// Run the command with the test case parameters
			output, err := runCommand(tc.mapFile, tc.startStation, tc.endStation, tc.numberOfTrains)

			// Check if the output contains the error message about too many stations
			if strings.Contains(output, "Error: Map contains more than 10000 stations") {
				t.Errorf("%sError: Map contains more than 10000 stations%s", ColorRed, ColorReset)
				return
			}

			// Check if there was an error running the command
			if err != nil {
				t.Errorf("%sFailed to run command: %v%s", ColorRed, err, ColorReset)
				return
			}

			// Count the number of turns (lines) in the output
			lines := strings.Split(strings.TrimSpace(output), "\n")
			actualTurns := len(lines)

			// Compare actual turns with expected turns
			if actualTurns != tc.expectedTurns {
				t.Errorf("%sWanted minimum %d turns, got %d turns.\nOutput:\n%s%s",
					ColorRed, tc.expectedTurns, actualTurns, output, ColorReset)
			} else {
				// If test passes, print success message with green color
				fmt.Printf("%sTest passed for %s to %s with %d trains.%s\nOutput:\n%s%s",
					ColorGreen, tc.startStation, tc.endStation, tc.numberOfTrains,
					ColorGreen, output, ColorReset)
			}
		})
	}
}
