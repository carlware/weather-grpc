package cmd

import (
	"crl/weather/cli/dispatchers/grpc"
	"fmt"
	"net"

	log "github.com/inconshreveable/log15"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
)

var port string
var host string

var serverCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Starts the weather gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		var group run.Group

		ln, err := net.Listen("tcp", host+":"+port)
		if err != nil {
			log.Error("Unable to bind port", err)
		}

		// Attach the dispatcher
		server, err := grpc.New()
		if err != nil {
			log.Error("Unable to start GRPC server", err)
			panic(fmt.Sprintf("Cannot start service %s", err))
		}

		// Add to goroutine group
		group.Add(
			func() error {
				log.Info("GRPC server listening ", "address:", ln.Addr().String())
				return server.Serve(ln)
			},
			func(e error) {
				log.Info("Shutting GRPC server down")
				server.GracefulStop()
			},
		)
		if err := group.Run(); err != nil {
			log.Error("Unable to start GRPC server", err)
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(&host, "host", "H", "localhost", "hostname")
	serverCmd.Flags().StringVarP(&port, "port", "P", "9003", "port number")
}
