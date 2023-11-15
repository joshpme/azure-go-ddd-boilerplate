package dtos

import (
	domain "ddd-boilerplate/internal/domain/instruments"
)

type LatLongDto struct {
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

type InstrumentAggregateDto struct {
	ID               string     `json:"id"`
	PartitionKey     string     `json:"partitionKey"`
	Name             string     `json:"name"`
	Type             string     `json:"type"`
	Status           string     `json:"status"`
	Location         string     `json:"location"`
	LatLong          LatLongDto `json:"latLong"`
	CollectionPeriod int        `json:"collectionPeriod"`
	Container        string     `json:"container"`
}

type LocationNameDto struct {
	Location string `json:"location"`
}

func InstrumentAggregateToDatabaseDto(ia domain.InstrumentAggregate) InstrumentAggregateDto {

	return InstrumentAggregateDto{
		ID:           string(ia.ID),
		Name:         ia.Name,
		PartitionKey: ia.PartitionKey(),
		Type:         string(ia.Type),
		Status:       string(ia.Status),
		Location:     ia.Location.Name,
		LatLong: LatLongDto{
			Lat:  ia.Location.Lat,
			Long: ia.Location.Long,
		},
		Container: string(ia.Container),
	}
}

func InstrumentAggregateFromDatabaseDto(dto InstrumentAggregateDto) (domain.InstrumentAggregate, error) {
	ia := domain.InstrumentAggregate{
		ID:        domain.InstrumentId(dto.ID),
		Name:      dto.Name,
		Type:      domain.InstrumentType(dto.Type),
		Status:    domain.CommissionStatus(dto.Status),
		Container: domain.Container(dto.Container),
		Location: domain.Location{
			Name: dto.Location,
			Lat:  dto.LatLong.Lat,
			Long: dto.LatLong.Long,
		},
	}
	return ia, nil
}
