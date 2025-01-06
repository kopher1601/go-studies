package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-chzzk/repository"
	"go-chzzk/service"
	"log"
)

type Server struct {
	engine *gin.Engine

	service    *service.Service
	repository *repository.Repository

	port string
	ip   string
}

func NewServer(service *service.Service, repository *repository.Repository, port string) *Server {
	n := &Server{
		engine:     gin.New(),
		service:    service,
		repository: repository,
		port:       port,
	}

	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.Default())

	r := NewRoom()
	go r.Run()

	n.engine.GET("/room", r.ServeHTTP)

	return n
}

func (n *Server) StartServer() error {
	log.Println("Starting server...")
	return n.engine.Run(":8080")
}
