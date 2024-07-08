package composite

import (
	"github.com/andremelinski/go-gcp/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/go-gcp/internal/pkg/utils"
	"github.com/andremelinski/go-gcp/internal/pkg/web"
	"github.com/andremelinski/go-gcp/internal/usecases"
)

func TemperatureLocationComposite(apiKey string) *handlers.TemperatureLocationHandler {

	httpHandler := web.NewWebResponseHandler()
	
	weatherApi := utils.NewWeatherInfo(apiKey)
	viaCep := utils.NewCepInfo()

	cepUsecase := usecases.NewLocationUseCase(viaCep)
	tempUseCase := usecases.NewClimateUseCase(weatherApi)
	
	return handlers.NewTemperatureLocationHandler(cepUsecase, tempUseCase, httpHandler)
}
