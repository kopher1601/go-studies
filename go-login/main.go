package main

import (
	"fmt"
	"go-login/db"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("MySQL接続OK")
}
