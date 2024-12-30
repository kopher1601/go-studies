package main

import (
	"fmt"
	"go-echo/db"
	"go-echo/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalln(fmt.Errorf("migrate user error: %v", err))
	}
}
