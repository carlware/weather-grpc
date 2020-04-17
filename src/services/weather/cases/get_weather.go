package cases

import (
	"context"
	"crl/weather/src/models"

	"crl/weather/src/interfaces/weather"
)

func GetWeather(ctx context.Context, req *models.GetWeatherRequest, weatherApi weather.Weather) (*models.GetWeatherResponse, error) {
	resp, err := weatherApi.GetTempByCity(req.City)
	if err != nil {
		return nil, err
	}
	return &models.GetWeatherResponse{
		Temperature: resp,
	}, nil
}
