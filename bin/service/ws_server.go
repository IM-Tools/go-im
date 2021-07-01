/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go_im/pkg/jwt"
	"go_im/pkg/ws"
	"net/http"
)

type WsServe struct {}

//所有逻辑 初始化连接 就订阅所有好友全部的频道
//定义一个ws服务
func (*WsServe)WsConn(c *gin.Context) {
	//开启携程启动默认程序
	// 将http请求升级为websocket协议
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)//
		return
	}
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	uuids := "channel_"+claims.ID
	client := &ws.Client{ID: uuids, Socket: conn, Send: make(chan []byte)}
	//注册一个新链接
	ws.Manager.Register <- client
	//启动协程读消息
	go client.Read()
    //启动协程写消息
	go client.Write()

}


