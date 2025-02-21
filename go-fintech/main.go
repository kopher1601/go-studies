package main

import (
	"fmt"
	"go-fintech/bitflyer"
	"go-fintech/config"
	"go-fintech/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
