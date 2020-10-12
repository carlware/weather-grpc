package cmd

import (
	"context"
	"crl/weather/src/interfaces/weather/openweather"
	"crl/weather/src/models"
	"crl/weather/src/services/weather/cases"
	"fmt"

	"github.com/spf13/cobra"
)

var isCelsius bool
var city string

var getWeatherCmd = &cobra.Command{
	Use:     "get-temp",
	Aliases: []string{"t"},
	Short:   "get temperature",
	Run: func(cmd *cobra.Command, args []string) {
		weatherAPI := openweather.NewOpenWeatherClient()

		req, _ := cases.GetWeather(context.TODO(), &models.GetWeatherRequest{
			City: city,
		}, weatherAPI)
		fmt.Println(city, " - ", req.Temperature, "K")
	},
}

func init() {
	getWeatherCmd.Flags().StringVarP(&city, "city", "c", "Corregidora", "City")
	//getWeatherCmd.Flags().BoolVarP(&isCelsius, "celsius", "f", false, "Celsis")
}
