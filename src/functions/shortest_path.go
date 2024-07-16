package functions

import (
	"errors"
	"station/src/data"
)

// Breadth-first search to find shortest path
func Bfs(start, end string, stations map[string]*data.Station) (map[string]string, error) {
	// Initialize queue with the starting station
	queue := []string{start}
	visited := make(map[string]bool)       // Keeps track of visited stations
	predecessor := make(map[string]string) // Maps each station to its predecessor in the shortest path
	visited[start] = true                  // Mark the starting station as visited

	// Loop until all stations reachable from start are processed
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue the current station

		// If we reach the destination station, return the predecessors which represent the shortest path
		if current == end {
			return predecessor, nil
		}

		// Visit all neighboring stations of the current station
		for _, neighbor := range stations[current].Connections {
			if !visited[neighbor.Name] {
				queue = append(queue, neighbor.Name) // Enqueue the neighboring station if not visited
				visited[neighbor.Name] = true        // Mark the neighboring station as visited
				predecessor[neighbor.Name] = current // Record the predecessor of the neighboring station
			}
		}
	}

	// If no path is found to the destination, return an error
	return nil, errors.New("no path found")
}
