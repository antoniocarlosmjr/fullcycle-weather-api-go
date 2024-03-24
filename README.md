# weather-api-go
This API is developed in Golang as a form of assessment for full cycle postgraduate courses.

## How to run

1) Clone the repository
2) Run the command `docker-compose up` in the root folder of the project
3) Access the URL `http://localhost:8080/api/v1/weather?cep=cep_number` in your browser or Postman, replacing `cep_number` with the cep of the city you want to know the weather.
4) The API will return a JSON with the weather information of the city you requested.
5) To stop the API, run the command `docker-compose down` in the root folder of the project.

## How to run tests

1) Run the command `go test ./...` in the root folder of the project
2) The tests will run and show the results in the terminal

## How to run the API remotely

1) Access url `` create in Cloud Run.