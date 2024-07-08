package main

import (
	"github.com/andremelinski/go-gcp/configs"
	"github.com/andremelinski/go-gcp/internal/composite"
	web_infra "github.com/andremelinski/go-gcp/internal/infra/web"
)

func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	tempHandler := composite.TemperatureLocationComposite(configs.WEATHER_API_KEY)

	webRouter := web_infra.NewWebRouter(tempHandler)
	webServer := web_infra.NewWebServer(
		configs.WebServerPort,
		webRouter.BuildHandlers(),
	)

	webServer.Start()
}