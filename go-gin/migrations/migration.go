package main

import (
	"fmt"
	"go-gin/infra"
	"go-gin/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate items: %s", err))
	}
}
