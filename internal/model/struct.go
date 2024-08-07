package model

// Station represents a railway station in the network.
type Station struct {
	Name        string     // The unique name of the station
	X, Y        int        // The X and Y coordinates of the station on a 2D grid
	Connections []*Station // Slice of pointers to other Station objects that this station is directly connected to
}

// OccupationInfo keeps track of which train occupies a station at each time step
type OccupationInfo struct {
	Station string // Name of the station that is occupied
	Time    int    // The time step at which the occupation occurs
	TrainID int    // The unique identifier of the train occupying the station
}

// PathWithOccupation represents a complete path for a train, including occupation information
type PathWithOccupation struct {
	Path        []string         // Ordered slice of station names representing the train's route
	Occupations []OccupationInfo // Slice of OccupationInfo structs, providing detailed occupation data for each step in the path
}
