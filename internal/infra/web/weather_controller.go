package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vgmrs/goexpert-weather-api/internal/usecase"
	"github.com/vgmrs/goexpert-weather-api/pkg/validator"
)

type GetWeatherController struct {
	GetWeatherUsecase usecase.GetWeatherUseCase
}

func NewGetWeatherController(GetWeatherUsecase *usecase.GetWeatherUseCase) *GetWeatherController {
	return &GetWeatherController{GetWeatherUsecase: *GetWeatherUsecase}
}

func (gw *GetWeatherController) getWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	if !validator.IsValidCEP(cep) {
		http.Error(w, errors.New("invalid zipcode").Error(), http.StatusUnprocessableEntity)
		return
	}

	dto := usecase.GetWeatherInputDTO{CEP: cep}
	output, err := gw.GetWeatherUsecase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
