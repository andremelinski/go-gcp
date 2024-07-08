package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
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
	// TODO vai pra env
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	
	if err != nil {
		return nil, err
	}
	data := &ViaCepDTO{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}

	if data.Bairro == "" {
		return nil, errors.New(string(res))
	}

	return data, nil
}