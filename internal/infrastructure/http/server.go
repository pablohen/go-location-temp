package http

import (
	"log"
	"net/http"
	"os"

	"go-location-temp/internal/interface/handler"
)

type Server struct {
	temperatureHandler *handler.TemperatureHandler
}

func NewServer(temperatureHandler *handler.TemperatureHandler) *Server {
	return &Server{
		temperatureHandler: temperatureHandler,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/temperature/", s.temperatureHandler.GetTemperatureByZipCode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	return http.ListenAndServe(":"+port, nil)
}
