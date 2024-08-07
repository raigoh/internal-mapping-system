package pathfinding

import "station/internal/model"

// findAllPaths uses depth-first search to find all possible paths from start to end
// Parameters:
//
//	start: The name of the starting station
//	end: The name of the destination station
//	stations: A map of all stations in the network, keyed by station name
//
// Returns:
//
//	A slice of slices, where each inner slice represents a valid path from start to end
func findAllPaths(start, end string, stations map[string]*model.Station) [][]string {
	// Initialize a slice to store all found paths
	var allPaths [][]string

	// Create a map to keep track of visited stations during the search
	visited := make(map[string]bool)

	// Define the depth-first search function
	// This is a closure that can access allPaths and visited
	var dfs func(current string, path []string)
	dfs = func(current string, path []string) {
		// Base case: if we've reached the end station
		if current == end {
			// Create a copy of the current path to avoid modifying it in future recursions
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			// Add the completed path to allPaths
			allPaths = append(allPaths, pathCopy)
			return
		}

		// Mark the current station as visited
		visited[current] = true

		// Explore all connections of the current station
		for _, neighbor := range stations[current].Connections {
			// If the neighboring station hasn't been visited, continue the search
			if !visited[neighbor.Name] {
				// Recursively call dfs for the neighbor, adding it to the path
				dfs(neighbor.Name, append(path, neighbor.Name))
			}
		}

		// Backtrack: mark the current station as unvisited
		// This allows the station to be visited again in different paths
		visited[current] = false
	}

	// Start the depth-first search from the start station
	// The initial path contains only the start station
	dfs(start, []string{start})

	// Return all found paths
	return allPaths
}
