package usecases

import (
	"github.com/andremelinski/go-gcp/internal/pkg/utils"
)

type LocationInputDTO struct {
	Cep    string  `json:"cep"`
}

type LocationOutputDTO struct{
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	DDD string `json:"ddd"`
}

type LocationUseCase struct {
	cepInfo utils.ICepInfoAPI
}

func NewLocationUseCase(cepInfo utils.ICepInfoAPI)*LocationUseCase{
	return &LocationUseCase{
		cepInfo,
	}
}

func (l *LocationUseCase)GetLocationInfo(cep string) (*LocationOutputDTO, error){
	cepnfo, err := l.cepInfo.GetCEPInfo(cep)

	if err != nil {
		return nil, err
	}

	return &LocationOutputDTO{
		Cep: cepnfo.Cep, 
		Logradouro: cepnfo.Logradouro, 
		Complemento: cepnfo.Complemento,
		Bairro: cepnfo.Bairro,
		Localidade: cepnfo.Localidade,
		UF: cepnfo.UF,
		DDD: cepnfo.DDD, 
	}, nil
}