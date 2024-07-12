package temperature

func CelsiusToFahrenheit(c float32) float32 {
	return (c*1.8) + 32
}

func CelsiusToKelvin(c float32) float32 {
	return c + 273
}
