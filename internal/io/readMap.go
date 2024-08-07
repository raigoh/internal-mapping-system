package io

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"station/internal/model"
	"station/internal/utils"
	"strconv"
	"strings"
)

// ReadMap reads and parses the network map from the specified file
// It returns a map of network names to maps of station names to Station structs, and any error encountered
func ReadMap(filepath string) (map[string]map[string]*model.Station, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer file.Close()

	// Initialize a map to store all networks
	allNetworks := make(map[string]map[string]*model.Station)
	currentNetwork := ""
	var currentStations map[string]*model.Station
	scanner := bufio.NewScanner(file)

	// Flags to track parsing state and section presence
	inStationsSection := false
	inConnectionsSection := false
	hasStationsSection := false
	hasConnectionsSection := false
	totalStations := 0 // Counter for total stations

	// Read the file line by line
	for scanner.Scan() {
		// Remove comments and leading/trailing whitespace
		line := strings.TrimSpace(strings.Split(scanner.Text(), "#")[0])
		if line == "" {
			continue
		}

		// Check for network delimiter
		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "---") {
			// Validate previous network if it exists
			if currentNetwork != "" {
				if !hasStationsSection {
					return nil, fmt.Errorf("%sError: Network '%s' does not contain a 'stations:' section%s", utils.Red, currentNetwork, utils.Reset)
				}
				if !hasConnectionsSection {
					return nil, fmt.Errorf("%sError: Network '%s' does not contain a 'connections:' section%s", utils.Red, currentNetwork, utils.Reset)
				}
			}
			// Start a new network
			currentNetwork = strings.Trim(line, "- ")
			currentStations = make(map[string]*model.Station)
			allNetworks[currentNetwork] = currentStations
			// Reset flags for the new network
			inStationsSection = false
			inConnectionsSection = false
			hasStationsSection = false
			hasConnectionsSection = false
			continue
		}

		// Determine which section we're in and set flags accordingly
		switch line {
		case "stations:":
			inStationsSection = true
			inConnectionsSection = false
			hasStationsSection = true
		case "connections:":
			inStationsSection = false
			inConnectionsSection = true
			hasConnectionsSection = true
		default:
			// Ensure we're inside a declared network before parsing data
			if currentNetwork == "" {
				return nil, fmt.Errorf("%sError: Found data before network declaration%s", utils.Red, utils.Reset)
			}
			// Parse station or connection based on current section
			if inStationsSection {
				if err := parseStation(line, currentStations, currentNetwork); err != nil {
					return nil, fmt.Errorf("%s%v%s", utils.Red, err, utils.Reset)
				}
				totalStations++ // Increment the total station count
				// Check if total stations exceed 10000
				if totalStations > 10000 {
					fmt.Fprintf(os.Stderr, "%sError: Map contains more than 10000 stations%s\n", utils.Red, utils.Reset)
					os.Exit(1)
				}
			} else if inConnectionsSection {
				if err := parseConnection(line, currentStations, currentNetwork); err != nil {
					return nil, fmt.Errorf("%s%v%s", utils.Red, err, utils.Reset)
				}
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("%sError: %v%s", utils.Red, err, utils.Reset)
	}

	// Validate the last network
	if currentNetwork != "" {
		if !hasStationsSection {
			return nil, fmt.Errorf("%sError: Network '%s' does not contain a 'stations:' section%s", utils.Red, currentNetwork, utils.Reset)
		}
		if !hasConnectionsSection {
			return nil, fmt.Errorf("%sError: Network '%s' does not contain a 'connections:' section%s", utils.Red, currentNetwork, utils.Reset)
		}
	}

	// Ensure at least one valid network was found
	if len(allNetworks) == 0 {
		return nil, fmt.Errorf("%sError: No valid networks found in the map file%s", utils.Red, utils.Reset)
	}
	return allNetworks, nil
}

// parseStation parses a single station line and adds the station to the stations map
func parseStation(line string, stations map[string]*model.Station, network string) error {
	// Split the line into parts: name, x, y
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return fmt.Errorf("%sInvalid station format in network %s: %s%s", utils.Red, network, line, utils.Reset)
	}

	// Parse and validate station name
	name := strings.TrimSpace(parts[0])
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
		return fmt.Errorf("%sInvalid station name in network %s: %s%s", utils.Red, network, name, utils.Reset)
	}

	// Parse and validate x coordinate
	x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || x < 0 {
		return fmt.Errorf("%sInvalid x coordinate for station %s in network %s%s", utils.Red, name, network, utils.Reset)
	}

	// Parse and validate y coordinate
	y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil || y < 0 {
		return fmt.Errorf("%sInvalid y coordinate for station %s in network %s%s", utils.Red, name, network, utils.Reset)
	}

	// Check for duplicate station names
	if _, exists := stations[name]; exists {
		return fmt.Errorf("%sDuplicate station name in network %s: %s%s", utils.Red, network, name, utils.Reset)
	}

	// Check for duplicate coordinates
	for _, station := range stations {
		if station.X == x && station.Y == y {
			return fmt.Errorf("%sError: Two stations exist at the same coordinates (%d, %d) in network %s%s", utils.Red, x, y, network, utils.Reset)
		}
	}

	// Add the new station to the stations map
	stations[name] = &model.Station{Name: name, X: x, Y: y, Connections: []*model.Station{}}
	return nil
}

// parseConnection parses a single connection line and updates the stations' connections
func parseConnection(line string, stations map[string]*model.Station, network string) error {
	// Split the line into two station names
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("%sInvalid connection format in network %s: %s%s", utils.Red, network, line, utils.Reset)
	}

	station1 := strings.TrimSpace(parts[0])
	station2 := strings.TrimSpace(parts[1])

	// Check for self-loop connections
	if station1 == station2 {
		return fmt.Errorf("%sSelf loop connection for station in network %s: %s%s", utils.Red, network, station1, utils.Reset)
	}

	// Retrieve station objects and check if they exist
	s1, exists1 := stations[station1]
	s2, exists2 := stations[station2]

	if !exists1 || !exists2 {
		return fmt.Errorf("%sStation does not exist in network %s%s", utils.Red, network, utils.Reset)
	}

	// Check for duplicate connections (in both directions)
	for _, conn := range s1.Connections {
		if conn.Name == station2 {
			return fmt.Errorf("%sError: Duplicate connection between %s and %s in network %s%s", utils.Red, station1, station2, network, utils.Reset)
		}
	}
	for _, conn := range s2.Connections {
		if conn.Name == station1 {
			return fmt.Errorf("%sError: Duplicate connection between %s and %s in network %s%s", utils.Red, station2, station1, network, utils.Reset)
		}
	}

	// Add bidirectional connection between stations
	s1.Connections = append(s1.Connections, s2)
	s2.Connections = append(s2.Connections, s1)
	return nil
}
