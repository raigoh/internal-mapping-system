package io

import (
	"fmt"
	"station/internal/model"
	"station/internal/utils"
	"strings"
)

func parseConnection(line string, stations map[string]*model.Station, network string, startStation string, endStation string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return utils.ErrInvalidConnectionFormat(network, line)
	}

	station1 := strings.TrimSpace(parts[0])
	station2 := strings.TrimSpace(parts[1])

	if station1 == station2 {
		return fmt.Errorf(utils.ErrSameStartEndStation)
	}

	s1, exists1 := stations[station1]
	s2, exists2 := stations[station2]

	_, exists3 := stations[startStation]
	_, exists4 := stations[endStation]

	if !exists3 {
		return fmt.Errorf(utils.ErrStartStationNotExist)
	}
	if !exists4 {
		return fmt.Errorf(utils.ErrEndStationNotExist)
	}

	if !exists1 {

		return fmt.Errorf(utils.ErrStationDoesNotExistInConnections)

	}
	if !exists2 {
		return fmt.Errorf(utils.ErrStationDoesNotExistInConnections)

	}

	// Check for duplicate connections
	for _, conn := range s1.Connections {
		if conn.Name == station2 {
			return fmt.Errorf("%s%s", utils.ErrDuplicateConnection(station1, station2), utils.Reset)
		}
	}

	s1.Connections = append(s1.Connections, s2)
	s2.Connections = append(s2.Connections, s1)
	return nil
}
