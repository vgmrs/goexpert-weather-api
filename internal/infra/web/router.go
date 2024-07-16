package web

import (
	"net/http"

	"github.com/vgmrs/goexpert-weather-api/internal/infra/client"
	"github.com/vgmrs/goexpert-weather-api/internal/usecase"
)

func Router() http.Handler {
	cepClient := client.NewCEPClient()
	weatherClient := client.NewWeatherClient()
	getWeatherUseCase := usecase.NewGetWeatherUseCase(cepClient, weatherClient)
	controller := NewGetWeatherController(getWeatherUseCase)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /weather/{cep}", controller.getWeather)

	return mux
}
