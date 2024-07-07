package usecases

type CepInfo struct {
	viacep u
}

type LocationInfoDTO struct{
	Cep string `json:"cep"` 
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	DDD string `json:"ddd"`
}

func GetTemperatureByCep(cep string) (*LocationInfoDTO, error){
	cepnfo := viacep.GetCEPInfo(cep)

	return &LocationInfoDTO{

	}
}