package utils

type IClimateInfoAPI interface{
	GetClimateInfo(place string) (*WeatherApiDTO, error)
}

type ICepInfoAPI interface{
	GetCEPInfo(cep string) (*ViaCepDTO, error)
}