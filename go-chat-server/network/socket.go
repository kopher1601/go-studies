package network

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-chat-server/types"
	"net/http"
)

// Http -> WebSocket 으로 업그레이드
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  types.SocketBufferSize,
	WriteBufferSize: types.MessageBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Room struct {
	Forward chan *message // 수신되는 메시지를 보관하는 값
	// 들어오는 메시지를 다른 클라이언트들에게 전송을 한다.

	Join  chan *Client // 소켓이 연결되는 경우에 작동
	Leave chan *Client // 소켓이 끊어지는 경우에 대해서 작동

	Clients map[*Client]bool // 전체 방에 있는 클라이언트 정보를 저장
}

type message struct {
	Name    string
	Message string
	Time    int64
}

type Client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
}
