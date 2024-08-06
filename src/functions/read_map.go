package functions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"station/src/data"
	"strconv"
	"strings"
)

// ReadMap reads and parses the network map from the file
func ReadMap(filepath string) (map[string]*data.Station, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	allStations := make(map[string]*data.Station)
	currentNetwork := ""
	scanner := bufio.NewScanner(file)

	inStationsSection := false
	inConnectionsSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(strings.Split(scanner.Text(), "#")[0])
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "---") {
			currentNetwork = strings.Trim(line, "- ")
			inStationsSection = false
			inConnectionsSection = false
			continue
		}

		switch line {
		case "stations:":
			inStationsSection = true
			inConnectionsSection = false
		case "connections:":
			inStationsSection = false
			inConnectionsSection = true
		default:
			if inStationsSection {
				if err := parseStation(line, allStations, currentNetwork); err != nil {
					return nil, err
				}
			} else if inConnectionsSection {
				if err := parseConnection(line, allStations, currentNetwork); err != nil {
					return nil, err
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allStations, nil
}

func parseStation(line string, stations map[string]*data.Station, network string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid station format in network %s: %s", network, line)
	}

	name := strings.TrimSpace(parts[0])
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
		return fmt.Errorf("invalid station name in network %s: %s", network, name)
	}

	x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || x < 0 {
		return fmt.Errorf("invalid x coordinate for station %s in network %s", name, network)
	}

	y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil || y < 0 {
		return fmt.Errorf("invalid y coordinate for station %s in network %s", name, network)
	}

	if _, exists := stations[name]; exists {
		return fmt.Errorf("duplicate station name in network %s: %s", network, name)
	}

	stations[name] = &data.Station{Name: name, X: x, Y: y, Connections: []*data.Station{}}
	return nil
}

func parseConnection(line string, stations map[string]*data.Station, network string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid connection format in network %s: %s", network, line)
	}

	station1 := strings.TrimSpace(parts[0])
	station2 := strings.TrimSpace(parts[1])

	if station1 == station2 {
		return fmt.Errorf("self loop connection for station in network %s: %s", network, station1)
	}

	s1, exists1 := stations[station1]
	s2, exists2 := stations[station2]

	if !exists1 || !exists2 {
		return fmt.Errorf("station does not exist in network %s", network)
	}

	for _, conn := range s1.Connections {
		if conn.Name == station2 {
			return fmt.Errorf("duplicate connection between %s and %s in network %s", station1, station2, network)
		}
	}

	s1.Connections = append(s1.Connections, s2)
	s2.Connections = append(s2.Connections, s1)
	return nil
}
