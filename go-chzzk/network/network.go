package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Network struct {
	engine *gin.Engine
}

func NewServer() *Network {
	n := &Network{engine: gin.New()}

	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.Default())

	r := NewRoom()
	go r.RunInit()

	n.engine.GET("/room", r.SocketServe)

	return n
}

func (n *Network) StartServer() error {
	log.Println("Starting server...")
	return n.engine.Run(":8080")
}
