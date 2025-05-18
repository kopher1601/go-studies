package network

import (
	"go-chzzk/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Room struct {
	Forward chan *message // 수신되는 메시지를 보관하는 값
	// 들어오는 메시지를 다른 클라이언트들에게 전송을 한다.

	Join    chan *client     // Socket 이 연결되는 경우에 작동
	Leave   chan *client     // Socket 이 끊어지는 경우에 작동
	Clients map[*client]bool // 현재 방에 있는 Client 정보를 저장
}

type message struct {
	Name    string
	Message string
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Clients: make(map[*client]bool),
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		panic(err)
	}

	client := &client{
		Socket: socket,
		Send:   make(chan *message, types.MessageBufferSize),
		Room:   r,
		Name:   userCookie.Value,
	}

	r.Join <- client

	defer func() {
		r.Leave <- client
	}()
}
