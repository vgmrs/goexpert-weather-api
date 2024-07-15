package client

type CEPOutputDTO struct {
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Unit         string `json:"unit"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type WeatherOutputDTO struct {
	Celsius   float64 `json:"celsius"`
	Condition string  `json:"condition"`
}

type CEPClientInterface interface {
	GetAddress(cep string) (*CEPOutputDTO, error)
}

type WeatherClientInterface interface {
	GetWeather(city string) (*WeatherOutputDTO, error)
}
