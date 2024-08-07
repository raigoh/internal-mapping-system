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

var errorTestCases = []struct {
	mapFile        string
	startStation   string
	endStation     string
	numberOfTrains int
}{
	{"/tests/errors/14duplicate-routes_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/22no-connections_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/19duplicate-names_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/10no-start-station_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/11no-end-station_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/21no-stations_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/16no-valid-coord_london.txt", "waterloo", "st_pancras", 2},
	{"network.map", "waterloo", "st_pancras", -2},
	{"tests/errors/invalidname_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/17same-coords_london.txt", "waterloo", "st_pancras", 2},
	{"network.map", "waterloo", "waterloo", 2},
	{"tests/errors/13no-path_london.txt", "waterloo", "st_pancras", 2},
	{"tests/errors/23over-tenK.txt", "station1", "station10001", 2},
}

// runCommand executes the main program with given parameters and returns its output
func runCommand(mapFile, start, end string, numTrains int) (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Construct the path to the main.go file (in the cmd folder)
	mainPath := filepath.Join(filepath.Dir(cwd), "cmd", "main.go")

	// Construct the full path to the map file (in the stations folder)
	//mapPath := filepath.Join(filepath.Dir(cwd), mapFile)

	// Prepare the command to run the main program
	cmd := exec.Command("go", "run", mainPath, mapFile, start, end, fmt.Sprint(numTrains))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	return out.String(), err
}

// TestPathfinder is the main test function that runs all test cases
func TestPathfinder(t *testing.T) {
	for _, tc := range testCases {
		// Create a subtest for each test case
		t.Run(fmt.Sprintf("%s_%s", tc.startStation, tc.endStation), func(t *testing.T) {
			// Run the command and get its output
			output, err := runCommand(tc.mapFile, tc.startStation, tc.endStation, tc.numberOfTrains)

			// Check if the map is too large (contains more than 10000 stations)
			if strings.Contains(output, "Error: Map contains more than 10000 stations") {
				t.Logf("%sMap contains more than 10000 stations, skipping test%s", utils.Red, utils.Reset)
				return
			}

			// Check if the command execution failed
			if err != nil {
				t.Errorf("%sFailed to run command: %v\nOutput: %s%s", utils.Red, err, output, utils.Reset)
				return
			}

			// Count the number of turns (lines) in the output
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

func TestErrors(t *testing.T) {
	for _, tc := range errorTestCases {
		output, _ := runCommand(tc.mapFile, tc.startStation, tc.endStation, tc.numberOfTrains)
		if strings.HasPrefix(output, "[31mError:") {
			t.Logf("%sGot error as expected: %s%s", utils.Green, output, utils.Reset)
		} else {
			t.Errorf("%sCould not get an error, Output: %s%s", utils.Red, output, utils.Reset)
		}
	}
}
