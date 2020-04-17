package weather

import (
	"context"
	v1 "crl/weather/grpc/weatherpb"
	"crl/weather/src/interfaces/weather"
	"crl/weather/src/models"
	"crl/weather/src/services/weather/cases"
)

type service struct {
	weatherAPI weather.Weather
}

func NewWeatherService(weatherAPI weather.Weather) Weather {
	return &service{
		weatherAPI: weatherAPI,
	}
}

func (s *service) GetWeather(ctx context.Context, req *v1.GetWeatherRequest) (*v1.GetWeatherResponse, error) {
	res, err := cases.GetWeather(context.TODO(), &models.GetWeatherRequest{
		City: req.City,
	}, s.weatherAPI)

	if err != nil {
		return nil, err
	}

	return &v1.GetWeatherResponse{
		Temperature: res.Temperature,
	}, nil
}
