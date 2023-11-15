package controller

import (
	"ddd-boilerplate/internal/database"
	"ddd-boilerplate/internal/domain/instruments"
	"ddd-boilerplate/internal/dtos"
	"net/http"
)

func ExecuteCreateCommand(conn *Connections, requestDto dtos.CreateInstrumentRequestDto) dtos.CreateInstrumentResponseDto {
	client, err := conn.DbManager.GetClient()
	if err != nil {
		return dtos.CreateInstrumentResponseDto{
			Success: false,
			Message: "Error getting db client: " + err.Error(),
		}
	}
	handler := &database.DatabaseHandler{Client: client}
	cmd := dtos.CreateInstrumentCommandFromRequestDto(requestDto)
	err = instruments.InstrumentCommandHandler(handler, cmd)
	if err != nil {
		return dtos.CreateInstrumentResponseDto{
			Success: false,
			Message: err.Error(),
		}
	}
	return dtos.CreateInstrumentResponseDto{
		Success: true,
		Message: "Successfully created instrument",
	}
}

func (conn *Connections) Create(w http.ResponseWriter, r *http.Request) {
	requestDto := dtos.CreateInstrumentRequestDto{}
	err := JsonRequest(r, &requestDto)
	if err != nil {
		BadRequest(w, err)
		return
	}
	result := ExecuteCreateCommand(conn, requestDto)
	JsonResponse(result, w)
}
