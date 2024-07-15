package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vgmrs/goexpert-weather-api/internal/infra/client"
)

func TestGetWeatherUseCase_Execute_Success(t *testing.T) {
	cepClient := client.NewCEPClient()
	weatherClient := client.NewWeatherClient()

	useCase := NewGetWeatherUseCase(cepClient, weatherClient)

	input := GetWeatherInputDTO{CEP: "01001000"}
	output, err := useCase.Execute(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, output.Celsius)
	assert.NotEmpty(t, output.Kelvin)
	assert.NotEmpty(t, output.Fahrenheit)
}

func TestGetWeatherUseCase_Execute_InvalidCEPError(t *testing.T) {
	cepClient := client.NewCEPClient()
	weatherClient := client.NewWeatherClient()

	useCase := NewGetWeatherUseCase(cepClient, weatherClient)

	input := GetWeatherInputDTO{CEP: "invalid-cep"}
	output, err := useCase.Execute(input)

	assert.Error(t, err)
	assert.Empty(t, output)
}
