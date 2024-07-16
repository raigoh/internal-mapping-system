package functions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"station/src/data"
	"strconv"
	"strings"
)

// ReadMap reads and parses the network map from the file
func ReadMap(filepath string) (map[string]*data.Station, error) {
	// Open the network map file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize stations map to store parsed station data
	stations := make(map[string]*data.Station)

	// Initialize scanner to read from file
	scanner := bufio.NewScanner(file)

	// Flags to track which section (stations or connections) the scanner is currently in
	inStationsSection := false
	inConnectionsSection := false

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// Remove comments and trim leading/trailing whitespace
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx] // Remove comments if present
		}
		line = strings.TrimSpace(line)
		// Skip empty lines
		if line == "" {
			continue
		}

		// Determine current section based on line content
		if line == "stations:" {
			inStationsSection = true
			inConnectionsSection = false
			continue
		} else if line == "connections:" {
			inStationsSection = false
			inConnectionsSection = true
			continue
		}

		// Process lines in the stations section
		if inStationsSection {
			// Split the line by commas to extract station details
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, errors.New("invalid station format")
			}
			// Extract and validate station name
			name := strings.TrimSpace(parts[0])
			if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
				return nil, fmt.Errorf("invalid station name: %s", name)
			}
			// Extract and validate x coordinate
			x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil || x < 0 {
				return nil, fmt.Errorf("invalid x coordinate for station %s", name)
			}
			// Extract and validate y coordinate
			y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err != nil || y < 0 {
				return nil, fmt.Errorf("invalide y coordinate for station %s", name)
			}
			// Check for duplicate station names
			if _, exists := stations[name]; exists {
				return nil, fmt.Errorf("duplicate station name: %s", name)
			}
			// Create a new Station object and add it to the stations map
			stations[name] = &data.Station{Name: name, X: x, Y: y, Connections: []*data.Station{}}

			// Process lines in the connections section
		} else if inConnectionsSection {
			// Split the line by dash to extract connection details
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid connection format: %s", line)
			}
			// Extract and trim station names
			station1 := strings.TrimSpace(parts[0])
			station2 := strings.TrimSpace(parts[1])

			// Validate that the connection doesn't loop back to the same station
			if station1 == station2 {
				return nil, fmt.Errorf("self loop connection for station: %s", station1)
			}
			// Check if both stations in the connection exist in the stations map
			if _, exists := stations[station1]; !exists {
				return nil, fmt.Errorf("station %s does not exist", station1)
			}
			if _, exists := stations[station2]; !exists {
				return nil, fmt.Errorf("station %s does not exist", station2)
			}

			// Check for duplicate connections between stations
			for _, conn := range stations[station1].Connections {
				if conn.Name == station2 {
					return nil, fmt.Errorf("duplicate connection between %s and %s", station1, station2)
				}
			}
			// Add bidirectional connections between station1 and station2
			stations[station1].Connections = append(stations[station1].Connections, stations[station2])
			stations[station2].Connections = append(stations[station2].Connections, stations[station1])
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// Return the parsed stations map
	return stations, nil
}
