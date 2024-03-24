package adapters

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/entities"
)

func SearchZipCode(zipCode string) (*entities.ZipCode, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + zipCode + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c entities.ZipCode
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
