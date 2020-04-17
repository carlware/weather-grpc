package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "service",
	Short: "Microservice to manage weather api",
}

func init() {
	mainCmd.AddCommand(getWeatherCmd)
	// mainCmd.AddCommand(serverCmd)
}

// Execute main command
func Execute() error {
	return mainCmd.Execute()
}
