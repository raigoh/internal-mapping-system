package io

import (
	"fmt"
	"regexp"
	"station/internal/model"
	"station/internal/utils"
	"strconv"
	"strings"
)

// parseStation parses a single station line and adds the station to the stations map
func parseStation(line string, stations map[string]*model.Station, network string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return utils.ErrNoConnectionsSections(network)
	}

	name := strings.TrimSpace(parts[0])
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(name) {
		return fmt.Errorf(utils.ErrInvalidStationNames)
	}

	x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || x < 0 {
		return fmt.Errorf("%s%s%s", utils.Red, utils.ErrInvalidCoordinate(true, x, name), utils.Reset)
	}

	y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil || y < 0 {
		return fmt.Errorf("%s%s%s", utils.ErrInvalidCoordinate(false, y, name), utils.Reset)
	}

	if _, exists := stations[name]; exists {
		return fmt.Errorf(utils.ErrDuplicateStationNames)
	}

	for _, station := range stations {
		if station.X == x && station.Y == y {
			return fmt.Errorf(utils.ErrSameCoordinates)
		}
	}

	stations[name] = &model.Station{Name: name, X: x, Y: y, Connections: []*model.Station{}}
	return nil
}
