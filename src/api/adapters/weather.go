package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/errors"
	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/entities"
)

func GetWeather(locale string) (*entities.Weather, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?q=%s&key=3841b81037a5427eb51191826241702", locale)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.UnableToRetrieveWeather
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c entities.Weather
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
