package tests

import (
	"fmt"
	"os/exec"
	"station/internal/utils"
	"strings"
	"testing"
)

// TestValidCases tests various valid scenarios
func TestValidCases(t *testing.T) {
	mainPath, err := findMainGo()
	if err != nil {
		t.Fatalf("Failed to find main.go: %v", err)
	}

	validTestCases := []struct {
		mapFile        string
		startStation   string
		endStation     string
		numberOfTrains int
		expectedTurns  int
	}{
		{"../network.map", "waterloo", "st_pancras", 4, 3},
		{"../network.map", "beginning", "terminus", 20, 11},
		{"../network.map", "beethoven", "part", 9, 6},
		{"../network.map", "small", "large", 9, 8},
		{"../network.map", "two", "four", 4, 6},
		{"../network.map", "jungle", "desert", 10, 8},
		{"../network.map", "bond_square", "space_port", 4, 6},
	}

	for _, tc := range validTestCases {
		t.Run(tc.mapFile, func(t *testing.T) {
			cmd := exec.Command("go", "run", mainPath, tc.mapFile, tc.startStation, tc.endStation, fmt.Sprintf("%d", tc.numberOfTrains))
			outputBytes, err := cmd.CombinedOutput()
			output := string(outputBytes) // Convert []byte to string
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
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
