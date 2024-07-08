package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/andremelinski/go-gcp/internal/pkg/web"
	"github.com/andremelinski/go-gcp/internal/usecases"
)

type TemperatureLocationHandler struct{
	CepUseCase usecases.ILocationInfo
	TempUseCase usecases.IWeatherInfo
	HttpResponse web.IWebResponseHandler
}

func NewTemperatureLocationHandler (cepUseCase usecases.ILocationInfo,tempUseCase usecases.IWeatherInfo, httpResponse web.IWebResponseHandler) *TemperatureLocationHandler{
	return &TemperatureLocationHandler{
		cepUseCase,
		tempUseCase,
		httpResponse,
	}
}

func(lc *TemperatureLocationHandler) CityTemperature(w http.ResponseWriter, r *http.Request){
	qs := r.URL.Query()
	zipStr := qs.Get("zipcode")

	if err := validateInput(zipStr); err != nil {
		lc.HttpResponse.RespondWithError(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	info, err := lc.CepUseCase.GetLocationInfo(zipStr)
	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
		return 
	}

	climateInfo, err := lc.TempUseCase.GetClimateUseCaseByName(info.Localidade)

	if err != nil {
		fmt.Println(err)
		lc.HttpResponse.RespondWithError(w, http.StatusBadRequest, errors.New("can not find zipcode"))
		return 
	}

	lc.HttpResponse.Respond(w, http.StatusOK, climateInfo)
}

func validateInput(zipcode string) error {
	if zipcode == "" {
		return errors.New("invalid zipcode")
	}

	matched, err := regexp.MatchString(`\b\d{5}[\-]?\d{3}\b`, zipcode)
	if !matched || err != nil {
		return errors.New("invalid zipcode")
	}

	return nil
}
