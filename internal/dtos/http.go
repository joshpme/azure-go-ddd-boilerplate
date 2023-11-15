package dtos

import (
	domain "ddd-boilerplate/internal/domain/instruments"
	"encoding/json"
)

type IResponseDto interface {
	Serialise() ([]byte, error)
}

type LatLongResponseDto struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type InstrumentResponseDto struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Status   string             `json:"status"`
	Location string             `json:"location"`
	LatLong  LatLongResponseDto `json:"latLong"`
}
type InstrumentListResponseDto struct {
	Instruments []InstrumentResponseDto `json:"instruments"`
}

type LocationNamesDto []string

type LocationNamesResponseDto struct {
	LocationNames LocationNamesDto `json:"locationNames"`
}

type LatLongRequestDto struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type CreateInstrumentRequestDto struct {
	ID               string            `json:"id"`
	Type             string            `json:"type"`
	Status           string            `json:"status"`
	Name             string            `json:"name"`
	Location         string            `json:"location"`
	LatLong          LatLongRequestDto `json:"latLong"`
	Container        string            `json:"container"`
	CollectionPeriod int               `json:"period"`
}

type CreateInstrumentResponseDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (dto CreateInstrumentResponseDto) Serialise() ([]byte, error) {
	return json.Marshal(dto)
}

func CreateInstrumentCommandFromRequestDto(dto CreateInstrumentRequestDto) domain.CreateInstrumentCommand {
	cmd := domain.CreateInstrumentCommand{
		ID:               domain.InstrumentId(dto.ID),
		Name:             dto.Name,
		Type:             domain.InstrumentType(dto.Type),
		Status:           domain.CommissionStatus(dto.Status),
		CollectionPeriod: dto.CollectionPeriod,
		Container:        domain.Container(dto.Container),
		Location: domain.Location{
			Name: dto.Location,
			Lat:  dto.LatLong.Lat,
			Long: dto.LatLong.Long,
		},
	}
	return cmd
}
