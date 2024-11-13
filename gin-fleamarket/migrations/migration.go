package main

import (
	"fmt"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database : %s", err.Error()))
	}
}
