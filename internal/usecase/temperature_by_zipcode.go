package usecase

import (
	"go-location-temp/internal/domain/entity"
	"go-location-temp/internal/domain/repository"
)

type TemperatureByZipCodeUseCase struct {
	zipCodeRepo repository.ZipCodeRepository
	weatherRepo repository.WeatherRepository
}

func NewTemperatureByZipCodeUseCase(
	zipCodeRepo repository.ZipCodeRepository,
	weatherRepo repository.WeatherRepository,
) *TemperatureByZipCodeUseCase {
	return &TemperatureByZipCodeUseCase{
		zipCodeRepo: zipCodeRepo,
		weatherRepo: weatherRepo,
	}
}

func (uc *TemperatureByZipCodeUseCase) Execute(zipCode string) (*entity.Temperature, error) {
	location, err := uc.zipCodeRepo.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	temperature, err := uc.weatherRepo.GetTemperatureByCity(location.City, location.State, location.Country)
	if err != nil {
		return nil, err
	}

	return temperature, nil
}
