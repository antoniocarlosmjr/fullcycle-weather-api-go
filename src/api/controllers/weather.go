package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	apiErrors "github.com/antoniocarlosmjr/weather-api-go/src/api/errors"
	"github.com/antoniocarlosmjr/weather-api-go/src/api/services"
)

func GetWeatherHandle(w http.ResponseWriter, r *http.Request) {
	zipCode := r.URL.Query().Get("cep")
	if zipCode == "" {
		http.Error(w, apiErrors.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		return
	}

	matched, err := regexp.MatchString(`^\d{8}$`, zipCode)
	if err != nil || !matched {
		http.Error(w, apiErrors.InvalidZipCode.Error(), http.StatusUnprocessableEntity)
		return
	}

	response, err := services.SearchWeather(zipCode)
	if err != nil && errors.Is(err, apiErrors.NotFoundZipCode) {
		http.Error(w, apiErrors.NotFoundZipCode.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
