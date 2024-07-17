package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// Helper function to run the main program with specified arguments and return the output
func runMain(t *testing.T, args []string) string {
	cmd := exec.Command("go", append([]string{"run", "."}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed with error: %v\nOutput: %s", err, out.String())
	}
	return out.String()
}

func TestStations(t *testing.T) {
	tests := []struct {
		mapFile       string
		start         string
		end           string
		numTrains     int
		expectedTurns int
	}{
		{"network.map", "waterloo", "st_pancras", 4, 3},
		{"network.map", "beginning", "terminus", 20, 11},
		{"network.map", "beethoven", "part", 4, 6},
		{"network.map", "small", "large", 4, 8},
		{"network.map", "two", "four", 4, 6},
		{"network.map", "jungle", "desert", 4, 8},
		{"network.map", "bond_square", "space_port", 4, 6},
		{"network.map", "london", "network", 2, 2},
		{"network.map", "london", "network", 1, 1},
		{"network.map", "london", "network", 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.start+"-"+tt.end, func(t *testing.T) {
			args := []string{tt.mapFile, tt.start, tt.end, strconv.Itoa(tt.numTrains)}
			output := runMain(t, args)
			lines := strings.Split(strings.TrimSpace(output), "\n")
			numTurns := len(lines)
			if numTurns != tt.expectedTurns {
				t.Errorf("Test %s-%s failed - wanted minimum %d, got %d turns", tt.start, tt.end, tt.expectedTurns, numTurns)
			}
		})
	}
}
