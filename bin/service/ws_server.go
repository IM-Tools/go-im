/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go_im/pkg/ws"
	"net/http"
)

type WsServe struct {}

//定义一个ws服务
func (*WsServe)WsConn(c *gin.Context) {
	// change the reqest to websocket model
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	ws.Manager.Register <- client

	go client.Read()
	go client.Write()
}
