package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/gabrielPossa/fc-go-gcr-weatherCEP/internal/cep"
	"github.com/gabrielPossa/fc-go-gcr-weatherCEP/internal/weather"
	"github.com/gabrielPossa/fc-go-gcr-weatherCEP/pkg/utils"
)

var digitCheck = regexp.MustCompile("^\\d{8}$")

type response struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeatherByCEP(w http.ResponseWriter, r *http.Request) {
	cepString := chi.URLParam(r, "CEP")

	cepString = strings.Replace(cepString, "-", "", -1)

	if !digitCheck.MatchString(cepString) {
		log.Println("CEP invalido,CEP deve ser composto por 8 números. Formatos aceitos: 12345678 ou 12345-678")
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	CEP, err := cep.FetchCEPData(r.Context(), cepString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if CEP.Erro == "true" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weatherQ := fmt.Sprintf("%s,%s", CEP.Localidade, CEP.Estado)

	weatherData, err := weather.GetWeatherData(r.Context(), weatherQ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := response{
		TempC: weatherData.Current.TempC,
		TempF: utils.CelciusToFahrenheit(weatherData.Current.TempC),
		TempK: utils.CelciusToKelvin(weatherData.Current.TempC),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
