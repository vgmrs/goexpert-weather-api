package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressByCEPRealAPI(t *testing.T) {
	cep := "01001000"
	expected := &CEPOutputDTO{
		Street:       "Praça da Sé",
		Complement:   "lado ímpar",
		Unit:         "",
		Neighborhood: "Sé",
		City:         "São Paulo",
		State:        "SP",
	}

	cepClient := NewCEPClient()
	address, err := cepClient.GetAddress(cep)
	assert.NoError(t, err)

	assert.Equal(t, expected.Street, address.Street)
	assert.Equal(t, expected.Complement, address.Complement)
	assert.Equal(t, expected.Unit, address.Unit)
	assert.Equal(t, expected.Neighborhood, address.Neighborhood)
	assert.Equal(t, expected.City, address.City)
	assert.Equal(t, expected.State, address.State)
}
