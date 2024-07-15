package usecase

import (
	"github.com/vgmrs/goexpert-weather-api/internal/entity"
	"github.com/vgmrs/goexpert-weather-api/internal/infra/client"
)

type GetWeatherInputDTO struct {
	CEP string `json:"cep"`
}

type GetWeatherOutputDTO struct {
	Celsius    string `json:"celsius"`
	Kelvin     string `json:"kelvin"`
	Fahrenheit string `json:"fahrenheit"`
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
		return GetWeatherOutputDTO{}, err
	}

	climate, err := gw.WeatherClient.GetWeather(address.City)
	if err != nil {
		return GetWeatherOutputDTO{}, err
	}

	weather := entity.NewWeather(climate.Celsius)

	return GetWeatherOutputDTO{
		Celsius:    weather.Celsius,
		Kelvin:     weather.Kelvin,
		Fahrenheit: weather.Fahrenheit,
	}, nil
}
