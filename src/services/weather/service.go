package weather

import (
	"context"
	v1 "crl/weather/grpc/weatherpb"
)

type service struct {
}

func NewWeather() Weather {
	return &service{}
}

func (s *service) GetWeather(ctx context.Context, req *v1.GetWeatherRequest) (*v1.GetWeatherResponse, error) {
	return &v1.GetWeatherResponse{
		temperature: "10",
	}, nil
}
