package main

import (
	"go-bitcoin/config"
	"go-bitcoin/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
