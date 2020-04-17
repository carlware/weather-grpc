package main

import (
	"crl/weather/cli/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("Unable to complete command execution", "err", err)
	}
}
