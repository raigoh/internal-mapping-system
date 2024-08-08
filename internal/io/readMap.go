package io

import (
	"bufio"
	"fmt"
	"os"
	"station/internal/model"
	"station/internal/utils"
	"strings"
)

const maxStations = 10000

// ReadMap reads and parses the network map from the specified file.
// It returns a map of network names to maps of station names to Station structs, and any error encountered.
func ReadMap(filepath string, startStation string, endStation string) (map[string]map[string]*model.Station, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	allNetworks := make(map[string]map[string]*model.Station)
	var currentNetwork string
	var currentStations map[string]*model.Station

	inStationsSection := false
	inConnectionsSection := false
	hasStationsSection := false
	hasConnectionsSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(strings.Split(scanner.Text(), "#")[0])
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "---") {
			// Validate the previous network before starting a new one
			if currentNetwork != "" {
				if err := validateNetwork(currentNetwork, hasStationsSection, hasConnectionsSection); err != nil {
					return nil, err
				}
			}

			// Start a new network
			currentNetwork = strings.Trim(line, "- ")
			currentStations = make(map[string]*model.Station)
			allNetworks[currentNetwork] = currentStations

			// Reset section flags
			inStationsSection = false
			inConnectionsSection = false
			hasStationsSection = false
			hasConnectionsSection = false
			continue
		}

		// Only process station data if we're inside a network section
		if currentNetwork == "" {
			return nil, fmt.Errorf("Error: Data found outside of a network section")
		}

		switch line {
		case "stations:":
			inStationsSection = true
			hasStationsSection = true
			inConnectionsSection = false
		case "connections:":
			inConnectionsSection = true
			hasConnectionsSection = true
			inStationsSection = false
		default:
			if inStationsSection {
				if len(currentStations) >= maxStations {
					return nil, fmt.Errorf("%s", utils.ErrTooManyStations)
				}

				if err := parseStation(line, currentStations, currentNetwork); err != nil {
					return nil, err
				}
			} else if inConnectionsSection {
				if err := parseConnection(line, currentStations, currentNetwork, startStation, endStation); err != nil {
					return nil, err
				}
			} else {
				return nil, utils.ErrNoStationsSections(currentNetwork)
			}
		}
	}

	// Validate the last network
	if err := validateNetwork(currentNetwork, hasStationsSection, hasConnectionsSection); err != nil {
		return nil, err
	}

	if len(allNetworks) == 0 {
		return nil, utils.ErrNoNetwork()
	}

	return allNetworks, nil
}

// validateNetwork checks that a network has both a stations section and a connections section.
func validateNetwork(network string, hasStations, hasConnections bool) error {
	if network == "" {
		return utils.ErrNoNetwork()
	}
	if !hasStations {
		return utils.ErrNoStationsSections(network)
	}
	if !hasConnections {
		return utils.ErrNoConnectionsSections(network)
	}
	return nil
}
