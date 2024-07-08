package usecases

type ILocationInfo interface {
	GetLocationInfo(cep string) (*LocationOutputDTO, error)
}

type IWeatherInfo interface{
	GetClimateUseCaseByName(name string) (*ClimateInfoDTO, error)
}