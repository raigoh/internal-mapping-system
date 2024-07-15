package functions

import (
	"fmt"
	"strings"
)

// SimulateTrains simulates train movements and outputs the movements to the terminal
func SimulateTrains(paths [][]string, numTrains int) {
	trainLocations := make(map[string]int)

	for i := 0; i < numTrains; i++ {
		trainLocations[fmt.Sprintf("T%d", i+1)] = 0
	}

	for step := 0; ; step++ {
		move := []string{}
		finishedTrains := 0
		for i := 0; i < numTrains; i++ {
			trainName := fmt.Sprintf("T%d", i+1)
			currentPos := trainLocations[trainName]
			if currentPos >= len(paths[i])-1 {
				finishedTrains++
				continue
			}
			nextStation := paths[i][currentPos+1]
			move = append(move, fmt.Sprintf("%s-%s", trainName, nextStation))
			trainLocations[trainName]++
		}
		if finishedTrains == numTrains {
			break
		}
		fmt.Println(strings.Join(move, " "))
	}
}
