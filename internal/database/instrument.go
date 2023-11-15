package database

import (
	instrumentDomain "ddd-boilerplate/internal/domain/instruments"
	"ddd-boilerplate/internal/domain/location"
	"ddd-boilerplate/internal/dtos"
	"encoding/json"
)

func (manager *DatabaseHandler) GetInstruments() ([]instrumentDomain.InstrumentAggregate, error) {
	viewContainer, err := manager.ViewContainer()
	if err != nil {
		return nil, err
	}

	rows, err := getRows(instrumentDomain.InstrumentAggregate{}.PartitionKey(), "select * from c", viewContainer)
	if err != nil {
		return nil, err
	}

	instruments := make([]instrumentDomain.InstrumentAggregate, 0)
	for _, row := range rows {
		var dto dtos.InstrumentAggregateDto
		err = json.Unmarshal(row, &dto)
		if err != nil {
			return nil, err
		}
		instrument, err := dtos.InstrumentAggregateFromDatabaseDto(dto)
		if err != nil {
			return nil, err
		}
		instruments = append(instruments, instrument)
	}

	return instruments, nil
}

func (manager *DatabaseHandler) GetLocationNames() (location.LocationNames, error) {
	viewContainer, err := manager.ViewContainer()
	if err != nil {
		return nil, err
	}

	rows, err := getRows(instrumentDomain.InstrumentAggregate{}.PartitionKey(), "select distinct c.location from c", viewContainer)
	if err != nil {
		return nil, err
	}

	locationNames := make([]string, 0)
	for _, row := range rows {
		var dto dtos.LocationNameDto
		err = json.Unmarshal(row, &dto)
		if err != nil {
			return nil, err
		}
		locationNames = append(locationNames, dto.Location)
	}

	return locationNames, nil
}
