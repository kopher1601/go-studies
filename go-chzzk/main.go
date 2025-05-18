package main

import "go-chzzk/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}
