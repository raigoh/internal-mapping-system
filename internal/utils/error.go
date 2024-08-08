package utils

import (
	"fmt"
)

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
	ErrInvalidStationNames   = "Error: Invalid station name in network"
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
	ErrTooManyStations      = "Error: Map contains more than 10000 stations"
)

// Enhanced error messages
func ErrDuplicateConnection(station1, station2 string) string {
	return fmt.Sprintf("Error: Duplicate connection between %s and %s", station1, station2)
}

func ErrInvalidCoordinate(isX bool, coord int, stationName string) string {
	coordType := "x"
	if !isX {
		coordType = "y"
	}
	return fmt.Sprintf("Error: Invalid %s coordinate for station %s", coordType, stationName)
}

func ErrNoConnectionsSections(network string) error {
	return fmt.Errorf("Error: Network '%s' does not contain a 'connections:' section", network)
}

// func ErrInvalidStationFormat(network, line string) error {
// 	return fmt.Errorf("Error: Invalid station format in network %s: %s", network, line)
// }

func ErrInvalidCoordinat(isX bool, coord int, stationName string) error {
	coordType := "x"
	if !isX {
		coordType = "y"
	}
	return fmt.Errorf("Error: Invalid %s coordinate for station %s", coordType, stationName)
}

func ErrInvalidConnectionFormat(network, line string) error {
	return fmt.Errorf("Error: Invalid connection format in network %s: %s", network, line)
}

func ErrSameStationConnection(station, network string) error {
	return fmt.Errorf("Error: Start and end station '%s' are the same in network '%s'", station, network)
}

func ErrStationNotExist(station, network string) error {
	return fmt.Errorf("Error: Station '%s' does not exist in network '%s'", station, network)
}

func ErrNoStationsSections(network string) error {
	return fmt.Errorf("Error: Network '%s' does not contain a 'stations:' section", network)
}

func ErrNoNetwork() error {
	return fmt.Errorf("Error: The map does not contain any networks")
}

// func ErrDataOutsideSection(network string) error {
// 	return fmt.Errorf("Error: Found data outside of stations or connections section in network '%s'", network)
// }
