package functions

import (
	"fmt"
	"sort"
	"station/src/data"
)

// OccupationInfo keeps track of which train occupies a station at each time step
type OccupationInfo struct {
	Station string // Name of the station
	Time    int    // Time step
	TrainID int    // ID of the train occupying the station
}

// PathWithOccupation represents a path and its corresponding occupation information
type PathWithOccupation struct {
	Path        []string         // Sequence of stations in the path
	Occupations []OccupationInfo // Occupation information for each step in the path
}

// FindPaths attempts to find all possible paths and select the best ones for multiple trains
// It returns the selected paths, their occupation information, and any error encountered
func FindPaths(start, end string, stations map[string]*data.Station, numTrains int) ([][]string, [][]OccupationInfo, error) {
	allPaths := findAllPaths(start, end, stations)
	if len(allPaths) == 0 {
		return nil, nil, fmt.Errorf("no paths found")
	}

	// Sort paths by length (shortest first)
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	selectedPaths := selectOptimalPaths(allPaths, numTrains, start, end)

	paths := make([][]string, len(selectedPaths))
	occupations := make([][]OccupationInfo, len(selectedPaths))
	for i, path := range selectedPaths {
		paths[i] = path
		occupations[i] = createOccupations(path, i)
	}
	// fmt.Println("PATH", paths)
	return paths, occupations, nil
}

// findAllPaths uses depth-first search to find all possible paths from start to end
func findAllPaths(start, end string, stations map[string]*data.Station) [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)

	// Define the depth-first search function
	var dfs func(current string, path []string)
	dfs = func(current string, path []string) {
		if current == end {
			// If we've reached the end, add the current path to allPaths
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			allPaths = append(allPaths, pathCopy)
			return
		}

		visited[current] = true
		// Explore all connections of the current station
		for _, neighbor := range stations[current].Connections {
			if !visited[neighbor.Name] {
				dfs(neighbor.Name, append(path, neighbor.Name))
			}
		}
		visited[current] = false // Backtrack
	}

	// Start the DFS from the start station
	dfs(start, []string{start})
	return allPaths
}

func selectOptimalPaths(allPaths [][]string, numTrains int, start, end string) [][]string {
	selectedPaths := make([][]string, 0, numTrains)
	occupiedStations := make(map[string]map[int]bool)

	// Helper function to check if a path conflicts with existing paths
	pathConflicts := func(path []string, startTime int) bool {
		for t, station := range path {
			if station != start && station != end {
				if occupied, exists := occupiedStations[station]; exists {
					if occupied[startTime+t] {
						return true
					}
				}
			}
		}
		return false
	}

	// Helper function to add a path to the selected paths
	addPath := func(path []string, startTime int) {
		delayedPath := make([]string, startTime+len(path))
		for i := 0; i < startTime; i++ {
			delayedPath[i] = start
		}
		copy(delayedPath[startTime:], path)
		selectedPaths = append(selectedPaths, delayedPath)
		for t, station := range path {
			if station != start && station != end {
				if occupiedStations[station] == nil {
					occupiedStations[station] = make(map[int]bool)
				}
				occupiedStations[station][startTime+t] = true
			}
		}
	}

	maxPathLength := len(allPaths[len(allPaths)-1])
	timeStep := 0

	for len(selectedPaths) < numTrains {
		for _, path := range allPaths {
			if len(selectedPaths) >= numTrains {
				break
			}
			if !pathConflicts(path, timeStep) {
				addPath(path, timeStep)
			}
		}
		timeStep++
		if timeStep > maxPathLength*numTrains {
			break // Safety check to prevent infinite loop
		}
	}

	return selectedPaths
}

func createOccupations(path []string, trainID int) []OccupationInfo {
	occupations := make([]OccupationInfo, len(path))
	for i, station := range path {
		occupations[i] = OccupationInfo{Station: station, Time: i, TrainID: trainID}
	}
	return occupations
}

// func applyDelay(path []string, delay int) []string {
// 	delayedPath := make([]string, delay+len(path))
// 	copy(delayedPath[delay:], path)
// 	for i := 0; i < delay; i++ {
// 		delayedPath[i] = path[0] // Stay at the start station during the delay
// 	}
// 	return delayedPath
// }

// // Helper function to check if a slice of slices contains a specific slice
// func contains(paths [][]string, path []string) bool {
// 	for _, p := range paths {
// 		if reflect.DeepEqual(p, path) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // strictPathOverlaps remains the same as in the original code
// func strictPathOverlaps(path []string, occupiedStations map[string]map[int]bool, start, end string) bool {
// 	for time, station := range path {
// 		if station != start && station != end {
// 			if _, exists := occupiedStations[station]; exists {
// 				// Check if the station is occupied at the current time or adjacent time steps
// 				for t := max(0, time-1); t <= time+1; t++ {
// 					if occupiedStations[station][t] {
// 						return true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// // updateOccupiedStations remains the same as in the original code
// func updateOccupiedStations(path []string, occupiedStations map[string]map[int]bool, start, end string) {
// 	for time, station := range path {
// 		if station != start && station != end {
// 			if occupiedStations[station] == nil {
// 				occupiedStations[station] = make(map[int]bool)
// 			}
// 			occupiedStations[station][time] = true
// 		}
// 	}
// }

// // createOccupations generates occupation information for a given path
// func createOccupations(path []string, trainID int) []OccupationInfo {
// 	occupations := make([]OccupationInfo, len(path))
// 	for i, station := range path {
// 		occupations[i] = OccupationInfo{Station: station, Time: i, TrainID: trainID}
// 	}
// 	return occupations
// }

// // max returns the maximum of two integers
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
