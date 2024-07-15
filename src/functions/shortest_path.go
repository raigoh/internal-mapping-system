package functions

import (
	"errors"
	"station/src/data"
)

// Breadth-first search to find shortest path
func Bfs(start, end string, stations map[string]*data.Station) (map[string]string, error) {
	queue := []string{start}
	visited := make(map[string]bool)
	predecessor := make(map[string]string)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return predecessor, nil
		}

		for _, neighbor := range stations[current].Connections {
			if !visited[neighbor.Name] {
				queue = append(queue, neighbor.Name)
				visited[neighbor.Name] = true
				predecessor[neighbor.Name] = current
			}
		}
	}

	return nil, errors.New("no path found")
}
