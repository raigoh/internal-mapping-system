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
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stations := make(map[string]*data.Station)
	scanner := bufio.NewScanner(file)
	inStationsSection := false
	inConnectionsSection := false
	for scanner.Scan() {
		line := scanner.Text()
		// Remove comments and trim spaces
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "stations:" {
			inStationsSection = true
			inConnectionsSection = false
			continue
		} else if line == "connections:" {
			inStationsSection = false
			inConnectionsSection = true
			continue
		}

		if inStationsSection {
			parts := strings.Split(line, ",")
			if len(parts) != 3 {
				return nil, errors.New("invalid station format")
			}
			name := strings.TrimSpace(parts[0])
			if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
				return nil, fmt.Errorf("invalid station name: %s", name)
			}
			x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil || x < 0 {
				return nil, fmt.Errorf("invalid x coordinate for station %s", name)
			}
			y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
			if err != nil || y < 0 {
				return nil, fmt.Errorf("invalide y coordinate for station %s", name)
			}
			if _, exists := stations[name]; exists {
				return nil, fmt.Errorf("duplicate station name: %s", name)
			}
			stations[name] = &data.Station{Name: name, X: x, Y: y, Connections: []*data.Station{}}
		} else if inConnectionsSection {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid connection format: %s", line)
			}
			station1 := strings.TrimSpace(parts[0])
			station2 := strings.TrimSpace(parts[1])
			if station1 == station2 {
				return nil, fmt.Errorf("self loop connection for station: %s", station1)
			}
			if _, exists := stations[station1]; !exists {
				return nil, fmt.Errorf("station %s does not exist", station1)
			}
			if _, exists := stations[station2]; !exists {
				return nil, fmt.Errorf("station %s does not exist", station2)
			}

			for _, conn := range stations[station1].Connections {
				if conn.Name == station2 {
					return nil, fmt.Errorf("duplicate connection between %s and %s", station1, station2)
				}
			}
			stations[station1].Connections = append(stations[station1].Connections, stations[station2])
			stations[station2].Connections = append(stations[station2].Connections, stations[station1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if !inStationsSection {
		return nil, errors.New("error: Map does not contain a 'stations:' section")
	}
	if !inConnectionsSection {
		return nil, errors.New("error: Map does not contain a 'connections:' section")
	}
	return stations, nil
}
