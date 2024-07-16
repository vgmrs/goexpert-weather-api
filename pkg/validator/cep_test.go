package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		name   string
		cep    string
		expect bool
	}{
		{"CEP valid with dash", "12345-678", true},
		{"CEP valid without dash", "12345678", true},
		{"CEP invalid with letters", "1234A-678", false},
		{"CEP invalid with special characters", "1234@-678", false},
		{"CEP invalid length less than 8 characters", "1234-678", false},
		{"CEP empty", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidCEP(test.cep)
			assert.Equal(t, test.expect, result, "IsValidCEP(%s)", test.cep)
		})
	}
}
