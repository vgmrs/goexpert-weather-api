package temperature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius    float64
		fahrenheit float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.6},
	}

	for _, test := range tests {
		fahrenheit := CelsiusToFahrenheit(test.celsius)
		assert.Equal(t, test.fahrenheit, fahrenheit)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		celsius float64
		kelvin  float64
	}{
		{0, 273},
		{100, 373},
		{-273, 0},
		{25, 298},
	}

	for _, test := range tests {
		kelvin := CelsiusToKelvin(test.celsius)
		assert.Equal(t, test.kelvin, kelvin)
	}
}
