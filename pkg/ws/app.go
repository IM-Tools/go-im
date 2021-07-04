package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	Webscoket *websocket.Conn
	err       error
)

func Hanlder(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	Webscoket, err = upgrader.Upgrade(w, r, nil)
	return Webscoket, err
}
