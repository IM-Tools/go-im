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
	Websocket *websocket.Conn
	err       error
)

func App(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	Websocket, err = upgrader.Upgrade(w, r, nil)
	return Websocket, err
}
