package functions

import "station/src/data"

// Find all paths using BFS
func FindPaths(start, end string, stations map[string]*data.Station, numTrains int) ([][]string, error) {
	var paths [][]string

	// Preserve original connections
	originalConnections := make(map[string][]*data.Station)
	for name, station := range stations {
		originalConnections[name] = append([]*data.Station{}, station.Connections...)
	}

	for i := 0; i < numTrains; i++ {
		predecessor, err := Bfs(start, end, stations)
		if err != nil {
			return nil, err
		}

		// Reconstruct path
		path := []string{}
		for at := end; at != ""; at = predecessor[at] {
			path = append([]string{at}, path...)
		}
		paths = append(paths, path)

		// Mark the path as used
		for j := 0; j < len(path)-1; j++ {
			current := path[j]
			next := path[j+1]
			for k, conn := range stations[current].Connections {
				if conn.Name == next {
					stations[current].Connections = append(stations[current].Connections[:k], stations[current].Connections[k+1:]...)
					break
				}
			}
			for k, conn := range stations[next].Connections {
				if conn.Name == current {
					stations[next].Connections = append(stations[next].Connections[:k], stations[next].Connections[k+1:]...)
					break
				}
			}
		}

		// Restore the original connections for the next iteration
		for name, station := range stations {
			station.Connections = append([]*data.Station{}, originalConnections[name]...)
		}
	}
	return paths, nil
}
