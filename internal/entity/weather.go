package entity

import (
	"fmt"

	"github.com/vgmrs/goexpert-weather-api/pkg/temperature"
)

type Weather struct {
	Celsius    string
	Kelvin     string
	Fahrenheit string
}

func NewWeather(c float64) *Weather {
	k := temperature.CelsiusToKelvin(c)
	f := temperature.CelsiusToFahrenheit(c)

	celsius := fmt.Sprintf("%.2f", c)
	kelvin := fmt.Sprintf("%.2f", k)
	fahrenheit := fmt.Sprintf("%.2f", f)

	return &Weather{
		Celsius:    celsius,
		Kelvin:     kelvin,
		Fahrenheit: fahrenheit,
	}
}
