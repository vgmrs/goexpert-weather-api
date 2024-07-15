package temperature

func CelsiusToFahrenheit(c float64) float64 {
	return (c * 1.8) + 32
}

func CelsiusToKelvin(c float64) float64 {
	return c + 273
}
