package socket_pkg

import (
	"github.com/gorilla/websocket"
	// "github.com/ishan/Chat_App/pkg/websocket"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	Pool     *Pool
	Username string
}
