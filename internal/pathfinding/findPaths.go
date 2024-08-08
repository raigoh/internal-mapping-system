package pathfinding

import (
	"fmt"
	"sort"
	"station/internal/core"
	"station/internal/model"
	"station/internal/utils"
)

// FindPaths attempts to find all possible paths and select the best ones for multiple trains
// It returns the selected paths, their occupation information, and any error encountered
// FindPaths attempts to find all possible paths and select the best ones for multiple trains
// It returns the selected paths, their occupation information, and any error encountered
func FindPaths(start, end string, stations map[string]*model.Station, numTrains int) ([][]string, [][]model.OccupationInfo, error) {
	// Check if start and end stations exist
	startExists := false
	endExists := false

	if _, exists := stations[start]; exists {
		startExists = true
	}
	if _, exists := stations[end]; exists {
		endExists = true
	}

	if !startExists && !endExists {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrStartStationNotExist, utils.Reset)
	}

	if !startExists {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrStartStationNotExist, utils.Reset)
	}

	if !endExists {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrEndStationNotExist, utils.Reset)
	}

	// Check if start and end stations are the same
	if start == end {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrSameStartEndStation, utils.Reset)
	}

	// Check if the number of trains is valid
	if numTrains <= 0 {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrInvalidTrainCount, utils.Reset)
	}

	// Find all possible paths between the start and end stations
	allPaths := findAllPaths(start, end, stations)

	// If no paths are found, return an error
	if len(allPaths) == 0 {
		return nil, nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoPath, utils.Reset)
	}

	// Sort the paths by length (shortest first)
	// This helps in selecting optimal paths, as shorter paths are generally preferred
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Select the optimal paths based on the number of trains
	// This function likely implements some logic to choose diverse and efficient paths
	selectedPaths := selectOptimalPaths(allPaths, numTrains, start, end)

	// Initialize slices to store the final paths and their occupation information
	paths := make([][]string, len(selectedPaths))
	occupations := make([][]model.OccupationInfo, len(selectedPaths))

	// For each selected path, create the corresponding occupation information
	for i, path := range selectedPaths {
		paths[i] = path
		// Create occupation information for each path, using the path index as the train ID
		occupations[i] = core.CreateOccupations(path, i)
	}

	// Return the selected paths, their occupation information, and nil error
	return paths, occupations, nil
}
