package usecases

import (
	"github.com/andremelinski/go-gcp/internal/pkg/utils"
)

type ClimateInfoDTO struct{
	Celsius float64 `json:"temp_C"` 
	Fahrenheit float64 `json:"temp_F"`
	Kelvin float64 `json:"temp_K"`
}
// retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin. 
type ClimateUseCase struct {
	WeatheInfo utils.IClimateInfoAPI
}

func NewClimateUseCase(climateApi utils.IClimateInfoAPI)*ClimateUseCase{
	return &ClimateUseCase{
		climateApi,
	}
}

func (l *ClimateUseCase)GetClimateUseCaseByName(name string) (*ClimateInfoDTO, error){
	weatherInfo, err := l.WeatheInfo.GetClimateInfo(name)

	if err != nil {
		return nil, err
	}

	return &ClimateInfoDTO{
		Celsius: weatherInfo.Current.TempC,
		Fahrenheit: weatherInfo.Current.TempF,
		Kelvin: weatherInfo.Current.TempC + 273,
	}, nil
}