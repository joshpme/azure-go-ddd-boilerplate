package dtos

import (
	domain "ddd-boilerplate/internal/domain/instruments"
	"ddd-boilerplate/internal/domain/location"
	"github.com/google/uuid"
	"strconv"
)

type EventDto struct {
	EventType string `json:"eventType"`
}

type CreateInstrumentEventDto struct {
	ID               string `json:"id"`
	PartitionKey     string `json:"partitionKey"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	Status           string `json:"status"`
	Location         string `json:"location"`
	LatLong          string `json:"latLong"`
	CollectionPeriod int    `json:"collectionPeriod"`
	Container        string `json:"container"`
	EventType        string `json:"eventType"`
}

func CreateInstrumentEventFromDatabaseDto(dto CreateInstrumentEventDto) (domain.CreateInstrumentEvent, error) {
	loc, err := location.CreateLocation(dto.LatLong, dto.Location)
	if err != nil {
		return domain.CreateInstrumentEvent{}, err
	}
	event := domain.CreateInstrumentEvent{
		ID:               domain.InstrumentId(dto.PartitionKey),
		Name:             dto.Name,
		Type:             domain.InstrumentType(dto.Type),
		Status:           domain.CommissionStatus(dto.Status),
		CollectionPeriod: dto.CollectionPeriod,
		Container:        domain.Container(dto.Container),
		Location:         loc,
	}
	return event, nil
}

func CreateInstrumentEventToDatabaseDto(s domain.CreateInstrumentEvent) CreateInstrumentEventDto {
	latLong := strconv.FormatFloat(s.Location.Lat, 'f', 6, 64) + "," + strconv.FormatFloat(s.Location.Long, 'f', 6, 64)
	return CreateInstrumentEventDto{
		ID:               uuid.NewString(),
		PartitionKey:     string(s.ID),
		Name:             s.Name,
		Type:             string(s.Type),
		Status:           string(s.Status),
		Location:         s.Location.Name,
		LatLong:          latLong,
		CollectionPeriod: s.CollectionPeriod,
		Container:        string(s.Container),
		EventType:        string(s.EventName()),
	}
}
