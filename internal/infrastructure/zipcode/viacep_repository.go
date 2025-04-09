package zipcode

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"go-location-temp/internal/domain/repository"
)

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"` // City
	UF          string `json:"uf"`         // State
	IBGE        string `json:"ibge"`
	Gia         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        string `json:"erro"`
}

type ViaCEPRepository struct {
	BaseURL string
	Client  *http.Client
}

func NewViaCEPRepository() *ViaCEPRepository {
	return &ViaCEPRepository{
		BaseURL: "https://viacep.com.br/ws",
		Client:  &http.Client{},
	}
}

func (r *ViaCEPRepository) GetLocationByZipCode(zipCode string) (*repository.Location, error) {
	if !r.isValidZipCodeFormat(zipCode) {
		return nil, errors.New("invalid zipcode")
	}

	zipCode = regexp.MustCompile(`\D`).ReplaceAllString(zipCode, "")

	url := fmt.Sprintf("%s/%s/json", r.BaseURL, zipCode)
	resp, err := r.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var viaCEPResp ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return nil, err
	}

	if viaCEPResp.Erro == "true" {
		return nil, errors.New("can not find zipcode")
	}

	return &repository.Location{
		City:    viaCEPResp.Localidade,
		State:   viaCEPResp.UF,
		Country: "Brazil",
		ZipCode: zipCode,
	}, nil
}

func (r *ViaCEPRepository) isValidZipCodeFormat(zipCode string) bool {
	digitsOnly := regexp.MustCompile(`\D`).ReplaceAllString(zipCode, "")
	return len(digitsOnly) == 8
}
