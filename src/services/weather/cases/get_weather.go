package cases

import (
	"context"
	"crl/weather/src/models"
)

func GetWeather(ctx context.Context, req *models.GetWeatherRequest) (*models.GetWeatherResponse, error) {
	return &models.GetWeatherResponse{
		Temperature: "10",
	}, nil
}
