package main

import "go-chat-server/network"

func main() {
	n := network.NewServer()

	err := n.StartServer()
	if err != nil {
		panic(err)
	}
}
