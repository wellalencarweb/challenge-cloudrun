package main

import (
	"github.com/wellalencarweb/challenge-cloudrun/config"
	"github.com/wellalencarweb/challenge-cloudrun/internal/pkg/dependencies"
)

func main() {
	configs, configsErr := config.LoadConfig(".")
	if configsErr != nil {
		panic(configsErr)
	}

	deps := dependencies.Build(configs)
	deps.WebServer.Start()
}
