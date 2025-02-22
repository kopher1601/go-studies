package controllers

import (
	"go-fintech/app/models"
	"go-fintech/bitflyer"
	"go-fintech/config"
	"log"
)

func StreamIngestionData() {
	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	for ticker := range tickerChannl {
		log.Printf("action=StreamIngestionData, %v", ticker)
		for _, duration := range config.Config.Durations {
			isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
			if isCreated == true && duration == config.Config.TradeDuration {
				// TODO
			}
		}
	}
}
