package functions

import "station/src/data"

// createOccupations generates occupation information for a given path and train
// Parameters:
//
//	path: A slice of strings representing the sequence of stations in the train's path
//	trainID: An integer identifier for the train
//
// Returns:
//
//	A slice of data.OccupationInfo structs representing the occupation of each station at each time step
func createOccupations(path []string, trainID int) []data.OccupationInfo {
	// Initialize a slice to store occupation information
	// The length of this slice will be equal to the number of stations in the path
	occupations := make([]data.OccupationInfo, len(path))

	// Iterate through each station in the path
	for i, station := range path {
		// Create an OccupationInfo struct for each station
		occupations[i] = data.OccupationInfo{
			Station: station, // The name of the current station
			Time:    i,       // The time step (assumed to be the index in the path)
			TrainID: trainID, // The ID of the train occupying this station
		}
	}

	// Return the completed slice of occupation information
	return occupations
}
