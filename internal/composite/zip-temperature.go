package composite

import (
	"github.com/andremelinski/go-gcp/internal/infra/web/webserver/handlers"
	"github.com/andremelinski/go-gcp/internal/pkg/utils"
	"github.com/andremelinski/go-gcp/internal/pkg/web"
	"github.com/andremelinski/go-gcp/internal/usecases"
)

func TemperatureLocationComposite(apiKey string) *handlers.LocalTemperatureHandler {

	httpHandler := web.NewWebResponseHandler()
	handlerExternalApi := utils.NewHandlerExternalApi()
	
	weatherApi := utils.NewWeatherInfo(apiKey, handlerExternalApi)
	viaCep := utils.NewCepInfo(handlerExternalApi)

	cepUsecase := usecases.NewLocationUseCase(viaCep)
	tempUseCase := usecases.NewClimateUseCase(weatherApi)
	
	return handlers.NewLocalTemperatureHandler(cepUsecase, tempUseCase, httpHandler)
}
