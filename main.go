package main

import (
	"net/http"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/controllers"
)

func main() {
	http.HandleFunc("/api/v1/weather", controllers.GetWeatherHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
