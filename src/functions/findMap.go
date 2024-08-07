package functions

import (
	"fmt"
	"station/src/data"
)

// FindAppropriateMap selects the most appropriate map based on the start and end stations
// Parameters:
//
//	networks: A map of network names to their corresponding station maps
//	start: The name of the starting station
//	end: The name of the destination station
//
// Returns:
//
//	string: The name of the appropriate network
//	map[string]*data.Station: The selected network's station map
//	error: An error if no appropriate map is found
func FindAppropriateMap(networks map[string]map[string]*data.Station, start, end string) (string, map[string]*data.Station, error) {
	// Iterate through all available networks
	for name, network := range networks {
		// Check if the start station exists in the current network
		if _, startExists := network[start]; startExists {
			// If the start station exists, check if the end station also exists
			if _, endExists := network[end]; endExists {
				// If both start and end stations exist in this network, return it
				return name, network, nil
			}
		}
	}

	// If no suitable network is found, return an error with red color
	return "", nil, fmt.Errorf("%sError: No map contains both the start and end stations%s", ColorRed, ColorReset)
}
