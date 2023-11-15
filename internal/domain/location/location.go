package location

import (
	"ddd-boilerplate/internal/domain/instruments"
	"errors"
	"strconv"
	"strings"
)

func CreateLocation(latLong string, location string) (instruments.Location, error) {
	parts := strings.Split(latLong, ",")
	if len(parts) != 2 {
		return instruments.Location{}, errors.New("invalid latLong (no comma)")
	}
	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return instruments.Location{}, errors.New("invalid latLong (lat could not be converted to a f64)")
	}
	long, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return instruments.Location{}, errors.New("invalid latLong (long could not be converted to a f64)")
	}
	return instruments.Location{
		Name: location,
		Lat:  lat,
		Long: long,
	}, nil
}
