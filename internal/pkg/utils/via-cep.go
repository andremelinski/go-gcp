package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type ViaCepDTO struct {
	Api string 
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	IBGE string `json:"ibge"`
	Gia string `json:"gia"`
	DDD string `json:"ddd"`
	Siafi string `json:"siafi"`
}

type CepInfo struct{}

func NewCepInfo() *CepInfo{
	return &CepInfo{}
}


func (c *CepInfo)GetCEPInfo(cep string) (*ViaCepDTO, error){
	ctx := context.Background()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	bytes, err := CallExternalApi(ctx, 3000, "GET", url)
	if err != nil {
		return nil, err
	}
	
	data := &ViaCepDTO{}
	json.Unmarshal(bytes, data)

	if data.Bairro == "" {
		return nil, errors.New(string(bytes))
	}

	return data, nil
}