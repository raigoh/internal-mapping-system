package functions

import (
	"fmt"
	"station/src/data"
)

// FindAppropriateMap selects the most appropriate map based on the start and end stations
func FindAppropriateMap(networks map[string]map[string]*data.Station, start, end string) (string, map[string]*data.Station, error) {
	for name, network := range networks {
		if _, startExists := network[start]; startExists {
			if _, endExists := network[end]; endExists {
				return name, network, nil
			}
		}
	}
	return "", nil, fmt.Errorf("no map contains both the start and end stations")
}
