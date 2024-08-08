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
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer file.Close()

	allNetworks := make(map[string]map[string]*model.Station)
	currentNetwork := "default" // Start with a default network
	currentStations := make(map[string]*model.Station)
	allNetworks[currentNetwork] = currentStations

	scanner := bufio.NewScanner(file)

	inStationsSection := false
	inConnectionsSection := false
	hasStationsSection := false
	hasConnectionsSection := false
	totalStations := 0

	for scanner.Scan() {
		line := strings.TrimSpace(strings.Split(scanner.Text(), "#")[0])
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "---") {
			if currentNetwork != "default" {
				if !hasStationsSection {
					return nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoStationsSection, utils.Reset)
				}
				if !hasConnectionsSection {
					return nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoConnectionsSection, utils.Reset)
				}
			}
			currentNetwork = strings.Trim(line, "- ")
			currentStations = make(map[string]*model.Station)
			allNetworks[currentNetwork] = currentStations
			inStationsSection = false
			inConnectionsSection = false
			hasStationsSection = false
			hasConnectionsSection = false
			continue
		}

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
			if inStationsSection {
				if err := parseStation(line, currentStations, currentNetwork); err != nil {
					return nil, err
				}
				totalStations++
				if totalStations > 10000 {
					return nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrTooManyStations, utils.Reset)
				}
			} else if inConnectionsSection {
				if err := parseConnection(line, currentStations, currentNetwork); err != nil {
					return nil, err
				}
			} else {
				return nil, fmt.Errorf("%sError: Found data outside of stations or connections section%s", utils.Red, utils.Reset)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	if currentNetwork != "default" {
		if !hasStationsSection {
			return nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoStationsSection, utils.Reset)
		}
		if !hasConnectionsSection {
			return nil, fmt.Errorf("%s%s%s", utils.Red, utils.ErrNoConnectionsSection, utils.Reset)
		}
	}

	if len(allNetworks) == 1 && currentNetwork == "default" {
		return nil, fmt.Errorf("%s%s%s", utils.ErrNoPath, utils.Red, utils.Reset)
	}

	return allNetworks, nil
}

// parseStation parses a single station line and adds the station to the stations map
func parseStation(line string, stations map[string]*model.Station, network string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return fmt.Errorf("%sInvalid station format in network %s: %s%s", utils.Red, network, line, utils.Reset)
	}

	name := strings.TrimSpace(parts[0])
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrInvalidStationName(name, 0), utils.Reset)
	}

	x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || x < 0 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrInvalidCoordinate(x, 0, name), utils.Reset)
	}

	y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil || y < 0 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrInvalidCoordinate(0, y, name), utils.Reset)
	}

	if _, exists := stations[name]; exists {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrDuplicateStationNames, utils.Reset)
	}

	for _, station := range stations {
		if station.X == x && station.Y == y {
			return fmt.Errorf("%s%s%s", utils.Red, utils.ErrSameCoordinates, utils.Reset)
		}
	}

	stations[name] = &model.Station{Name: name, X: x, Y: y, Connections: []*model.Station{}}
	return nil
}

func parseConnection(line string, stations map[string]*model.Station, network string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("%sInvalid connection format in network %s: %s%s", utils.Red, network, line, utils.Reset)
	}

	station1 := strings.TrimSpace(parts[0])
	station2 := strings.TrimSpace(parts[1])

	if station1 == station2 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrSameStartEndStation, utils.Reset)
	}

	s1, exists1 := stations[station1]
	s2, exists2 := stations[station2]

	if !exists1 && !exists2 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrEndStationNotExist, utils.Reset)
	}

	if !exists1 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrStartStationNotExist, utils.Reset)
	}
	if !exists2 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrEndStationNotExist, utils.Reset)
	}

	for _, conn := range s1.Connections {
		if conn.Name == station2 {
			return fmt.Errorf("%s%s%s", utils.Red, utils.ErrDuplicateConnection(station1, station2), utils.Reset)
		}
	}
	for _, conn := range s2.Connections {
		if conn.Name == station1 {
			return fmt.Errorf("%s%s%s", utils.Red, utils.ErrDuplicateConnection(station2, station1), utils.Reset)
		}
	}

	s1.Connections = append(s1.Connections, s2)
	s2.Connections = append(s2.Connections, s1)
	return nil
}
