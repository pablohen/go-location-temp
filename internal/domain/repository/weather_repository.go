package repository

import "go-location-temp/internal/domain/entity"

type WeatherRepository interface {
	GetTemperatureByCity(city, state, country string) (*entity.Temperature, error)
}
