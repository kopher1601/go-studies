package main

import (
	"flag"
	"fmt"
	"go-chzzk/config"
	"go-chzzk/network"
	"go-chzzk/repository"
	"go-chzzk/service"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", "1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	if rep, err := repository.NewRepository(c); err != nil {
		panic(err)
	} else {
		service := service.NewService(rep)
		network.NewServer(service, rep)
		fmt.Println(rep)
	}

	fmt.Println(c)
	//n := network.NewServer()
	//n.StartServer()
}
