package utils

type IClimateInfoAPI interface{
	GetWeatherInfo(place string) (*WeatherApiDTO, error)
}

type ICepInfoAPI interface{
	GetCEPInfo(cep string) (*ViaCepDTO, error)
}