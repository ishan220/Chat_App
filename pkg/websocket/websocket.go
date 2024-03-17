package socket_pkg

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ishan/Chat_App/model"
	// "github.com/ishan/Chat_App/pkg/httpserver"
)

type Message struct {
	Type string     `json:"type"`
	User string     `json:"user"`
	Chat model.Chat `json:"chat"`
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	CheckOrigin: func(r *http.Request) bool { //currently no checking and just allow any connection
		return true
	},
}
