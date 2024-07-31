package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const APIKey string = ""

type WeatherClient struct{}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{}
}

type WeatherResponse struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
}

func (w *WeatherClient) GetWeather(city string) (*WeatherOutputDTO, error) {
	escapedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", APIKey, escapedCity)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data WeatherResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &WeatherOutputDTO{
		Celsius:   data.Current.TempC,
		Condition: data.Current.Condition.Text,
	}, nil
}
