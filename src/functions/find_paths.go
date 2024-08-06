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
	// Find all possible paths from start to end
	allPaths := findAllPaths(start, end, stations)
	if len(allPaths) == 0 {
		return nil, nil, fmt.Errorf("no paths found")
	}

	// Sort paths by length (shortest first) to prioritize shorter paths
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Select paths that don't overlap based on strict criteria
	selectedPaths := selectStrictNoOverlapPaths(allPaths, numTrains, start, end)

	// Create occupation info for selected paths
	paths := make([][]string, len(selectedPaths))
	occupations := make([][]OccupationInfo, len(selectedPaths))
	for i, path := range selectedPaths {
		paths[i] = path
		occupations[i] = createOccupations(path, i)
	}

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

// selectStrictNoOverlapPaths selects paths that don't overlap based on strict criteria
func selectStrictNoOverlapPaths(allPaths [][]string, numTrains int, start, end string) [][]string {
	selectedPaths := make([][]string, 0, numTrains)
	occupiedStations := make(map[string]map[int]bool)

	for i := 0; i < numTrains && i < len(allPaths); i++ {
		var selectedPath []string
		// Find the first path that doesn't overlap with previously selected paths
		for _, path := range allPaths {
			if !strictPathOverlaps(path, occupiedStations, start, end) {
				selectedPath = path
				break
			}
		}

		if selectedPath == nil {
			break // No more non-overlapping paths available
		}

		selectedPaths = append(selectedPaths, selectedPath)
		updateOccupiedStations(selectedPath, occupiedStations, start, end)
	}

	return selectedPaths
}

// strictPathOverlaps checks if a path overlaps with previously occupied stations
func strictPathOverlaps(path []string, occupiedStations map[string]map[int]bool, start, end string) bool {
	for time, station := range path {
		if station != start && station != end {
			if _, exists := occupiedStations[station]; exists {
				// Check if the station is occupied at the current time or adjacent time steps
				for t := max(0, time-1); t <= time+1; t++ {
					if occupiedStations[station][t] {
						return true
					}
				}
			}
		}
	}
	return false
}

// updateOccupiedStations marks stations as occupied for a given path
func updateOccupiedStations(path []string, occupiedStations map[string]map[int]bool, start, end string) {
	for time, station := range path {
		if station != start && station != end {
			if occupiedStations[station] == nil {
				occupiedStations[station] = make(map[int]bool)
			}
			occupiedStations[station][time] = true
		}
	}
}

// createOccupations generates occupation information for a given path
func createOccupations(path []string, trainID int) []OccupationInfo {
	occupations := make([]OccupationInfo, len(path))
	for i, station := range path {
		occupations[i] = OccupationInfo{Station: station, Time: i, TrainID: trainID}
	}
	return occupations
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // FindPaths attempts to find up to numTrains unique paths from start to end using BFS
// func FindPaths(start, end string, stations map[string]*data.Station, numTrains int) ([][]string, error) {
// 	var paths [][]string

// 	// Preserve original connections
// 	originalConnections := make(map[string][]*data.Station)
// 	for name, station := range stations {
// 		originalConnections[name] = append([]*data.Station{}, station.Connections...)
// 	}

// 	// Attempt to find paths in a loop up to numTrains times
// 	for i := 0; i < numTrains; i++ {
// 		// Use BFS to find the shortest path
// 		predecessor, err := Bfs(start, end, stations)
// 		if err != nil {
// 			break // Exit the loop if no path is found
// 		}

// 		// Reconstruct path from end to start using predecessors
// 		path := []string{}
// 		for at := end; at != ""; at = predecessor[at] {
// 			path = append([]string{at}, path...)
// 		}
// 		paths = append(paths, path)

// 		// Mark the path as used by removing connections
// 		for j := 0; j < len(path)-1; j++ {
// 			current := path[j]
// 			next := path[j+1]
// 			for k, conn := range stations[current].Connections {
// 				if conn.Name == next {
// 					stations[current].Connections = append(stations[current].Connections[:k], stations[current].Connections[k+1:]...)
// 					break
// 				}
// 			}
// 			// Also remove the reverse connection
// 			for k, conn := range stations[next].Connections {
// 				if conn.Name == current {
// 					stations[next].Connections = append(stations[next].Connections[:k], stations[next].Connections[k+1:]...)
// 					break
// 				}
// 			}
// 		}
// 	}

// 	// Restore the original connections after finding paths
// 	for name, station := range stations {
// 		station.Connections = append([]*data.Station{}, originalConnections[name]...)
// 	}

// 	// Return the found paths or an error if no paths were found
// 	if len(paths) == 0 {
// 		return nil, errors.New("no paths found")
// 	}

// 	return paths, nil
// }
