package main

import (
	"flag"
	"fmt"
	"go-chat-server/config"
	"go-chat-server/network"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()
	c := config.NewConfig(*pathFlag)
	fmt.Println(c)

	n := network.NewServer()

	err := n.StartServer()
	if err != nil {
		panic(err)
	}
}
