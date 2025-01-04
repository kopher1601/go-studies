package main

import (
	"flag"
	"fmt"
	"go-chzzk/config"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", "1010", "port set")

func main() {
	flag.Parse()

	c := config.NewConfig(*pathFlag)

	fmt.Println(c)
	//n := network.NewServer()
	//n.StartServer()
}
