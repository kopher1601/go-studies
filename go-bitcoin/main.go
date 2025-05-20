package main

import (
	"fmt"
	"go-bitcoin/bitflyer"
	"go-bitcoin/config"
	"go-bitcoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
