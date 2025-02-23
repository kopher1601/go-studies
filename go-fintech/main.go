package main

import (
	"fmt"
	"go-fintech/app/models"
	"go-fintech/config"
	"go-fintech/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}
