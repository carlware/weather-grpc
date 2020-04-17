package openweather

import (
	"crl/weather/src/interfaces/weather"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type service struct {
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Response struct {
	Main Main `json:"main"`
}

func NewOpenWeatherClient() weather.Weather {
	return &service{}
}

func (s *service) GetTempByCity(city string) (string, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=53f278b7bac5d2745ac0617a7ce21da3", city)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return strconv.FormatFloat(responseObject.Main.Temp, 'f', 6, 64), nil

}
