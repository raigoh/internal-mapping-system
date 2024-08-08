package utils

import "fmt"

const (
	// Command Line Errors
	ErrIncorrectArgCount = "Error: Incorrect number of command line arguments"
	ErrTooFewArgs        = "Error: Too few command line arguments"
	ErrTooManyArgs       = "Error: Too many command line arguments"

	// Station Errors
	ErrStartStationNotExist  = "Error: Start station does not exist"
	ErrEndStationNotExist    = "Error: End station does not exist"
	ErrSameStartEndStation   = "Error: Start and end station are the same"
	ErrDuplicateStationNames = "Error: Duplicate station names"
	ErrInvalidStationNames   = "Error: Invalid station names"
	ErrSameCoordinates       = "Error: Two stations exist at the same coordinates"

	// Connection Errors
	ErrNoPath                = "Error: no paths found"
	ErrDuplicateConnections  = "Error: Duplicate connections, including those which are described in reverse"
	ErrNonexistentConnection = "Error: Connection with a station which does not exist"

	// Input Validation Errors
	ErrInvalidTrainCount  = "Error: Number of trains is not a valid positive integer"
	ErrInvalidCoordinates = "Error: Coordinates which are not valid positive integers"

	// Map Structure Errors
	ErrNoStationsSection    = "Error: The map does not contain a \"stations:\" section"
	ErrNoConnectionsSection = "Error: The map does not contain a \"connections:\" section"
	ErrTooManyStations      = "Error: A map contains more than 10000 stations"
)

// Enhanced error messages
func ErrDuplicateConnection(station1, station2 string) string {
	return fmt.Sprintf("Error: Duplicate connection between %s and %s", station1, station2)
}

func ErrInvalidStationName(name string, line int) string {
	return fmt.Sprintf("Error: Invalid station name \"%s\" on line %d", name, line)
}

func ErrInvalidCoordinate(x, y int, stationName string) string {
	return fmt.Sprintf("Error: Coordinate [%d,%d] is not a valid positive integer for station %s", x, y, stationName)
}
