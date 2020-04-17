package grpc

import (
	"google.golang.org/grpc"
)

type application struct {
	server *grpc.Server
}

var app *application

func New() (*grpc.Server, error) {
	var err error

	app = &application{}

	app.server, err = grpcServer()
	return app.server, err
}
