package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	utils_dto "github.com/andremelinski/go-gcp/internal/pkg/utils/dto"
	utils_interface "github.com/andremelinski/go-gcp/internal/pkg/utils/interface"
)


type WeatherInfo struct{
	apiKey string
	handlerExternalApi utils_interface.IHandlerExternalApi
}

func NewWeatherInfo(apiKey string, handlerExternalApi utils_interface.IHandlerExternalApi) *WeatherInfo{
	return &WeatherInfo{
		apiKey,
		handlerExternalApi,
	}
}

func (c *WeatherInfo)GetWeatherInfo(place string) (*utils_dto.WeatherApiDTO, error){
	ctx := context.Background()

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes", c.apiKey, place)

	bytes, err := c.handlerExternalApi.CallExternalApi(ctx, 1000, "GET", url)
	
	if err != nil {
		return nil, err
	}

	dto := &utils_dto.WeatherApiDTO{}
	json.Unmarshal(bytes, dto)
	if dto.Location.Name == "" {
		return nil, errors.New("no matching location found")
	}
	return dto, nil
}