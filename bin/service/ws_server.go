/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var cache_user_id =  "im_cache_user_id";

type WsServe struct {}

//定义一个ws服务
func (*WsServe)WsConn(c *gin.Context) {

	var (
		err error
		data []byte
	)


	// 将http请求升级为websocket协议
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)//
		return
	}


	//uuids := uuid.NewV4().String()
	// websocket connect
	//client := &ws.Client{ID: uuids, Socket: conn, Send: make(chan []byte)}

	//需要判断双方都在线
	for {
		if _,data,err  =  conn.ReadMessage();err!= nil {
			goto ERR
		}
	   if err =	conn.WriteMessage(websocket.TextMessage,data);err!=nil{
		goto ERR
	   }
	}

	//ws.Manager.Register <- client
	//
	//go client.Read()
	//go client.Write()

	ERR:
		conn.Close()
}

