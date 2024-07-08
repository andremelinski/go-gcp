package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type WeatherApiDTO struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		WindchillC float64 `json:"windchill_c"`
		WindchillF float64 `json:"windchill_f"`
		HeatindexC float64 `json:"heatindex_c"`
		HeatindexF float64 `json:"heatindex_f"`
		DewpointC  float64 `json:"dewpoint_c"`
		DewpointF  float64 `json:"dewpoint_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
		AirQuality struct {
			Co           float64 `json:"co"`
			No2          float64 `json:"no2"`
			O3           float64 `json:"o3"`
			So2          float64 `json:"so2"`
			Pm25         float64 `json:"pm2_5"`
			Pm10         float64 `json:"pm10"`
			UsEpaIndex   int     `json:"us-epa-index"`
			GbDefraIndex int     `json:"gb-defra-index"`
		} `json:"air_quality"`
	} `json:"current"`
}

type WeatherInfo struct{
	apiKey string
}

func NewWeatherInfo(apiKey string) *WeatherInfo{
	return &WeatherInfo{
		apiKey,
	}
}

func (c *WeatherInfo)GetWeatherInfo(place string) (*WeatherApiDTO, error){
	ctx := context.Background()

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=curitiba&aqi=yes", c.apiKey)

	bytes, err := CallExternalApi(ctx, 1000, "GET", url)
	
	if err != nil {
		return nil, err
	}

	dto := &WeatherApiDTO{}
	json.Unmarshal(bytes, dto)
	if dto.Location.Name == "" {
		return nil, errors.New("no matching location found")
	}

	return dto, nil
}