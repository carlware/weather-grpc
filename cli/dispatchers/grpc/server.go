package grpc

import (
	"crl/weather/grpc/weatherpb"
	"crl/weather/src/interfaces/weather/openweather"
	"crl/weather/src/services/weather"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func grpcServer() (*grpc.Server, error) {
	sopts := []grpc.ServerOption{}

	sopts = append(sopts, grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)

	server := grpc.NewServer(sopts...)

	// Health
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	// initialize OpenWeather
	weatherAPI := openweather.NewOpenWeatherClient()

	// services
	weatherService := weather.NewWeatherService(weatherAPI)
	weatherpb.RegisterWeatherAPIServer(server, weatherService)

	// Reflection
	reflection.Register(server)

	// Return no error
	return server, nil
}
