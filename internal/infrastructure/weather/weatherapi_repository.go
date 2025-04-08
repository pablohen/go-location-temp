package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go-location-temp/internal/domain/entity"
)

type WeatherAPIResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type WeatherAPIRepository struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewWeatherAPIRepository() *WeatherAPIRepository {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		apiKey = "default_key"
	}

	return &WeatherAPIRepository{
		BaseURL: "https://api.weatherapi.com/v1",
		APIKey:  apiKey,
		Client:  &http.Client{},
	}
}

func (r *WeatherAPIRepository) GetTemperatureByCity(city, state, country string) (*entity.Temperature, error) {
	locationQuery := fmt.Sprintf("%s,%s,%s", city, state, country)

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", r.BaseURL, r.APIKey, locationQuery)
	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var weatherResp WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, err
	}

	return entity.NewTemperature(weatherResp.Current.TempC), nil
}
