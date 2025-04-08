package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-location-temp/internal/usecase"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type TemperatureHandler struct {
	useCase *usecase.TemperatureByZipCodeUseCase
}

func NewTemperatureHandler(useCase *usecase.TemperatureByZipCodeUseCase) *TemperatureHandler {
	return &TemperatureHandler{
		useCase: useCase,
	}
}

func (h *TemperatureHandler) GetTemperatureByZipCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		h.respondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}
	zipCode := pathParts[len(pathParts)-1]

	temperature, err := h.useCase.Execute(zipCode)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			h.respondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
			return
		}
		if err.Error() == "can not find zipcode" {
			h.respondWithError(w, http.StatusNotFound, "can not find zipcode")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "failed to get temperature")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}

func (h *TemperatureHandler) respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
