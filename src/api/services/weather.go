package services

import (
	"net/url"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/adapters"
	"github.com/antoniocarlosmjr/weather-api-go/src/api/errors"
	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/dto"
)

func SearchWeather(zipCode string) (*dto.WeatherResponse, error) {
	resViaCep, err := adapters.SearchZipCode(zipCode)
	if resViaCep.Cep == "" {
		return nil, errors.NotFoundZipCode
	}

	if err != nil {
		return nil, errors.UnableToRetrieveZipCode
	}

	city := url.QueryEscape(resViaCep.Locale)
	resWeather, err := adapters.GetWeather(city)
	if err != nil {
		return nil, errors.UnableToRetrieveWeather
	}

	return &dto.WeatherResponse{
		TempC: resWeather.Current.TempC,
		TempK: getTemperatureKelvin(resWeather.Current.TempC),
		TempF: getTemperatureFahrenheit(resWeather.Current.TempC),
	}, nil
}

func getTemperatureFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func getTemperatureKelvin(celsius float64) float64 {
	return celsius + 273
}
