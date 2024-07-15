package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CEPClient struct{}

func NewCEPClient() *CEPClient {
	return &CEPClient{}
}

type CEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c CEPClient) GetAddress(cep string) (*CEPOutputDTO, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data CEPResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &CEPOutputDTO{
		Street:       data.Logradouro,
		Complement:   data.Complemento,
		Unit:         data.Unidade,
		Neighborhood: data.Bairro,
		City:         data.Localidade,
		State:        data.UF,
	}, nil
}
