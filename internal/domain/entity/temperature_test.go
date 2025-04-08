package entity_test

import (
	"testing"

	"go-location-temp/internal/domain/entity"
)

func TestNewTemperature(t *testing.T) {
	testCases := []struct {
		name           string
		celsius        float64
		wantFahrenheit float64
		wantKelvin     float64
	}{
		{
			name:           "0 Celsius",
			celsius:        0,
			wantFahrenheit: 32,
			wantKelvin:     273,
		},
		{
			name:           "100 Celsius",
			celsius:        100,
			wantFahrenheit: 212,
			wantKelvin:     373,
		},
		{
			name:           "Negative temperature",
			celsius:        -10,
			wantFahrenheit: 14,
			wantKelvin:     263,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			temp := entity.NewTemperature(tc.celsius)

			if temp.Celsius != tc.celsius {
				t.Errorf("Expected Celsius to be %.2f, got %.2f", tc.celsius, temp.Celsius)
			}

			if temp.Fahrenheit != tc.wantFahrenheit {
				t.Errorf("Expected Fahrenheit to be %.2f, got %.2f", tc.wantFahrenheit, temp.Fahrenheit)
			}

			if temp.Kelvin != tc.wantKelvin {
				t.Errorf("Expected Kelvin to be %.2f, got %.2f", tc.wantKelvin, temp.Kelvin)
			}
		})
	}
}
