package functions

import (
	"errors"
	"station/src/data"
)

// FindPaths attempts to find up to numTrains unique paths from start to end using BFS
func FindPaths(start, end string, stations map[string]*data.Station, numTrains int) ([][]string, error) {
	var paths [][]string

	// Preserve original connections
	originalConnections := make(map[string][]*data.Station)
	for name, station := range stations {
		originalConnections[name] = append([]*data.Station{}, station.Connections...)
	}

	// Attempt to find paths in a loop up to numTrains times
	for i := 0; i < numTrains; i++ {
		// Use BFS to find the shortest path
		predecessor, err := Bfs(start, end, stations)
		if err != nil {
			break // Exit the loop if no path is found
		}

		// Reconstruct path from end to start using predecessors
		path := []string{}
		for at := end; at != ""; at = predecessor[at] {
			path = append([]string{at}, path...)
		}
		paths = append(paths, path)

		// Mark the path as used by removing connections
		for j := 0; j < len(path)-1; j++ {
			current := path[j]
			next := path[j+1]
			for k, conn := range stations[current].Connections {
				if conn.Name == next {
					stations[current].Connections = append(stations[current].Connections[:k], stations[current].Connections[k+1:]...)
					break
				}
			}
			// Also remove the reverse connection
			for k, conn := range stations[next].Connections {
				if conn.Name == current {
					stations[next].Connections = append(stations[next].Connections[:k], stations[next].Connections[k+1:]...)
					break
				}
			}
		}
	}

	// Restore the original connections after finding paths
	for name, station := range stations {
		station.Connections = append([]*data.Station{}, originalConnections[name]...)
	}

	// Return the found paths or an error if no paths were found
	if len(paths) == 0 {
		return nil, errors.New("no paths found")
	}
	return paths, nil
}
