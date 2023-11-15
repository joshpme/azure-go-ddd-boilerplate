package controller

import (
	"ddd-boilerplate/internal/database"
	domain "ddd-boilerplate/internal/domain/instruments"
	"ddd-boilerplate/internal/domain/location"
	"ddd-boilerplate/internal/dtos"
	"net/http"
)

func (conn *Connections) List(w http.ResponseWriter, r *http.Request) {
	client, err := conn.DbManager.GetClient()
	if err != nil {
		JsonResponse(err.Error(), w)
		return
	}

	handler := &database.DatabaseHandler{Client: client}
	domainInstruments, err := domain.GetInstrumentList(handler)
	instruments := make([]dtos.InstrumentResponseDto, 0)
	for _, instrument := range domainInstruments {
		instruments = append(instruments, dtos.InstrumentResponseDto{
			ID:       string(instrument.ID),
			Name:     instrument.Name,
			Type:     string(instrument.Type),
			Status:   string(instrument.Status),
			Location: instrument.Location.Name,
			LatLong: dtos.LatLongResponseDto{
				Lat:  instrument.Location.Lat,
				Long: instrument.Location.Long,
			},
		})
	}
	instrumentListResponseDto := dtos.InstrumentListResponseDto{Instruments: instruments}

	if err != nil {
		JsonResponse(err.Error(), w)
		return
	}

	JsonResponse(instrumentListResponseDto, w)
}

func (conn *Connections) Locations(w http.ResponseWriter, r *http.Request) {
	client, err := conn.DbManager.GetClient()
	if err != nil {
		JsonResponse(err.Error(), w)
		return
	}
	handler := &database.DatabaseHandler{Client: client}
	domainLocationNames, err := location.GetLocationNames(handler)
	locationNames := make([]string, 0)
	for _, locationName := range domainLocationNames.LocationNames {
		locationNames = append(locationNames, locationName)
	}
	locationNamesResponseDto := dtos.LocationNamesResponseDto{LocationNames: locationNames}

	if err != nil {
		JsonResponse(err.Error(), w)
		return
	}

	JsonResponse(locationNamesResponseDto, w)
}
