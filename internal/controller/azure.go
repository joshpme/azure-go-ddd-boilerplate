package controller

import (
	"ddd-boilerplate/internal/database"
	"encoding/json"
	"net/http"
)

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

func JsonResponse(data interface{}, w http.ResponseWriter) {
	responseJson, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	JsonResponse(err.Error(), w)
}

func JsonRequest(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

type Connections struct {
	DbManager database.IStorageHandler
}
