package network

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-chzzk/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	engine *gin.Engine

	service *service.Service

	port string
	ip   string
}

func NewServer(service *service.Service, port string) *Server {
	s := &Server{engine: gin.New(), service: service, port: port}

	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	registerServer(s)

	return s
}

func (s *Server) StartServer() error {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	log.Printf("Start Tx Server")
	return s.engine.Run(s.port)
}
