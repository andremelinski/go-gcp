package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andremelinski/go-gcp/internal/pkg/web"
	"github.com/andremelinski/go-gcp/internal/usecases"
)

type LocationTemperatureHandler struct{
	CepUseCase usecases.ILocationInfo
	TempUseCase usecases.IWeatherInfo
	HttpResponse web.IWebResponseHandler
}

func NewLocationTemperatureHandler (cepUseCase usecases.ILocationInfo,tempUseCase usecases.IWeatherInfo, httpResponse web.IWebResponseHandler) *LocationTemperatureHandler{
	return &LocationTemperatureHandler{
		cepUseCase,
		tempUseCase,
		httpResponse,
	}
}

func(lc *LocationTemperatureHandler) CityTemperature(w http.ResponseWriter, r *http.Request){
	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	// if err := validateInput(zipStr); err != nil {
	// 	lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	
	info, err := lc.CepUseCase.GetLocationInfo(zipStr)
	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
	}
	climateInfo, err := lc.TempUseCase.GetClimateUseCaseByName(info.Localidade)

	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
	}

	lc.HttpResponse.Respond(w, http.StatusOK, climateInfo)
}