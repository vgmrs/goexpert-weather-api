package usecase

import (
	"errors"

	"github.com/vgmrs/goexpert-weather-api/internal/entity"
	"github.com/vgmrs/goexpert-weather-api/internal/infra/client"
)

type GetWeatherInputDTO struct {
	CEP string `json:"cep"`
}

type GetWeatherOutputDTO struct {
	Celsius    string `json:"temp_C"`
	Fahrenheit string `json:"temp_F"`
	Kelvin     string `json:"temp_K"`
}

type GetWeatherUseCase struct {
	CEPClient     client.CEPClientInterface
	WeatherClient client.WeatherClientInterface
}

func NewGetWeatherUseCase(
	CEPClient client.CEPClientInterface,
	WeatherClient client.WeatherClientInterface,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		CEPClient:     CEPClient,
		WeatherClient: WeatherClient,
	}
}

func (gw *GetWeatherUseCase) Execute(input GetWeatherInputDTO) (GetWeatherOutputDTO, error) {
	address, err := gw.CEPClient.GetAddress(input.CEP)
	if err != nil {
		return GetWeatherOutputDTO{}, errors.New("can not find zipcode")
	}

	climate, err := gw.WeatherClient.GetWeather(address.City)
	if err != nil {
		return GetWeatherOutputDTO{}, errors.New("can not find weather for zipcode")
	}

	weather := entity.NewWeather(climate.Celsius)

	return GetWeatherOutputDTO{
		Celsius:    weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin:     weather.Kelvin,
	}, nil
}
