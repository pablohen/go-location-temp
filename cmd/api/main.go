package main

import (
	"log"

	"go-location-temp/internal/infrastructure/http"
	"go-location-temp/internal/infrastructure/weather"
	"go-location-temp/internal/infrastructure/zipcode"
	"go-location-temp/internal/interface/handler"
	"go-location-temp/internal/usecase"
)

func main() {
	zipCodeRepo := zipcode.NewViaCEPRepository()
	weatherRepo := weather.NewWeatherAPIRepository()

	temperatureUseCase := usecase.NewTemperatureByZipCodeUseCase(zipCodeRepo, weatherRepo)

	temperatureHandler := handler.NewTemperatureHandler(temperatureUseCase)

	server := http.NewServer(temperatureHandler)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
