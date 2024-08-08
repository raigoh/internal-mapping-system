package tests

import (
	"os/exec"
	"path/filepath"
	"station/internal/utils"
	"strconv"
	"strings"
	"testing"
)

// findMainGo searches for main.go in the parent directory
func findMainGo() (string, error) {
	dir, err := filepath.Abs("../..")
	if err != nil {
		return "", err
	}
	mainPath := filepath.Join(dir, "main.go")
	return mainPath, nil
}

// TestErrorCases tests various error scenarios
func TestErrorCases(t *testing.T) {
	mainPath, err := findMainGo()
	if err != nil {
		t.Fatalf("Failed to find main.go: %v", err)
	}

	// Get the absolute path to the tests directory
	testsDir, err := filepath.Abs(".")
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	errorTestCases := []struct {
		mapFile        string
		startStation   string
		endStation     string
		numberOfTrains int
		expectedError  string
	}{
		{"10no-start-station_london.txt", "waterloo", "st_pancras", 2, "Error: Start station does not exist"},
		{"11no-end-station_london.txt", "waterloo", "st_pancras", 2, "Error: End station does not exist"},
		{"12same-start-end_london.txt", "waterloo", "waterloo", 2, "Error: Start and end station are the same"},
		{"13no-path_london.txt", "waterloo", "st_pancras", 2, "Error: no paths found"},
		{"14duplicate-routes_london.txt", "waterloo", "st_pancras", 2, "Error: Duplicate connection between"},
		{"16no-valid-coord_london.txt", "waterloo", "st_pancras", 2, "Error: Invalid x coordinate for station"},
		{"17same-coords_london.txt", "waterloo", "st_pancras", 2, "Error: Two stations exist at the same coordinates"},
		{"18station-not-exist_london.txt", "waterloo", "st_pancras", 2, "Error: Connetion to non existing station"},
		{"19duplicate-names_london.txt", "waterloo", "st_pancras", 2, "Error: Duplicate station names"},
		{"21no-stations_london.txt", "waterloo", "st_pancras", 2, "Error: Network 'London Network Map' does not contain a 'stations:' section"},
		{"22no-connections_london.txt", "waterloo", "st_pancras", 2, "Error: Network 'London Network Map' does not contain a 'connections:' section"},
		{"23over-tenK.txt", "station1", "station10001", 2, "Error: Map contains more than 10000 stations"},
		{"invalidname_london.txt", "waterloo", "st_pancras", 2, "Error: Invalid station name in network"},
		{"network.map", "waterloo", "st_pancras", -2, "Error: Number of trains is not a valid positive integer"},
	}

	for _, tc := range errorTestCases {
		testName := tc.mapFile
		t.Run(testName, func(t *testing.T) {
			var mapPath string
			if tc.mapFile == "network.map" {
				mapPath = filepath.Join(testsDir, tc.mapFile)
			} else {
				mapPath = filepath.Join(testsDir, tc.mapFile)
			}

			cmd := exec.Command("go", "run", mainPath, mapPath, tc.startStation, tc.endStation, strconv.Itoa(tc.numberOfTrains))
			output, err := cmd.CombinedOutput()

			if err == nil {
				t.Errorf("%sFAILED: %s - Expected an error, but got none, output is: %s%s", utils.Red, testName, output, utils.Reset)
			} else if !strings.Contains(string(output), tc.expectedError) {
				t.Errorf("%sFAILED: %s - Expected error containing '%s', but got: %s%s", utils.Red, testName, tc.expectedError, string(output), utils.Reset)
			} else {
				t.Logf("%sPASSED: %s - Got expected error: %s%s", utils.Green, testName, strings.TrimSpace(string(output)), utils.Reset)
			}
		})
	}
}
