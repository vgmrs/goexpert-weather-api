package entity

import (
	"fmt"

	"github.com/vgmrs/goexpert-weather-api/pkg/temperature"
)

type Weather struct {
	Celsius    string
	Fahrenheit string
	Kelvin     string
}

func NewWeather(c float64) *Weather {
	f := temperature.CelsiusToFahrenheit(c)
	k := temperature.CelsiusToKelvin(c)

	celsius := fmt.Sprintf("%.2f", c)
	fahrenheit := fmt.Sprintf("%.2f", f)
	kelvin := fmt.Sprintf("%.2f", k)

	return &Weather{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     kelvin,
	}
}
