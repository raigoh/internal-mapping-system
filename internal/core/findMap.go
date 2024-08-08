package core

import (
	"fmt"
	"station/internal/model"
	"station/internal/utils"
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
func FindAppropriateMap(networks map[string]map[string]*model.Station, start, end string) (string, map[string]*model.Station, error) {
	startExists := false
	endExists := false

	// Iterate through all available networks to check if the start and end stations exist
	for _, network := range networks {
		if _, exists := network[start]; exists {
			startExists = true
		}
		if _, exists := network[end]; exists {
			endExists = true
		}
	}

	// If start station does not exist in any network
	if !startExists {
		return "", nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrStartStationNotExist, utils.Reset)
	}

	// If end station does not exist in any network
	if !endExists {
		return "", nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrEndStationNotExist, utils.Reset)
	}

	// Iterate through all available networks to find one containing both the start and end stations
	for name, network := range networks {
		if _, startExists := network[start]; startExists {
			if _, endExists := network[end]; endExists {
				return name, network, nil
			}
		}
	}

	// If no suitable network is found
	return "", nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoPath, utils.Reset)
}
