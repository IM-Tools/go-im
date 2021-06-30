/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	messageModel "go_im/bin/http/models/msg"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"net/http"
	"strconv"
	"time"
)

var cache_user_id =  "im_cache_user_id";

type WsServe struct {}

//所有逻辑 初始化连接 就订阅所有好友全部的频道

type Msg struct {
	FromId int `json:"from_id"`
	Msg string `json:"msg"`
	ToId int `json:"to_id"`
}


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
		// 判断用户是否在线「在频道内
		// 在投递并且入库
		// 不在将消息放入mysql和redis 用户上线即通知消费

	   if err =	conn.WriteMessage(websocket.TextMessage,data);err!=nil{
		goto ERR
	   }
	   //携程消费
	   go putData(data)

	}

	//ws.Manager.Register <- client
	//
	//go client.Read()
	//go client.Write()

	ERR:
		conn.Close()
}
//put
func putData(data []byte) {
	msg := new(Msg)

	if err :=	json.Unmarshal([]byte(string(data)),&msg);err!= nil {
		fmt.Println(err)
	}

	channel_a,_ := helpler.ProduceChannelName( strconv.Itoa(msg.FromId), strconv.Itoa(msg.ToId))

	fid := uint64(msg.FromId)
	tid := uint64(msg.ToId)

	user := messageModel.ImMessage{FromId:fid,
		ToId: tid,
		Msg: msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05"),
		Channel: channel_a}

	model.DB.Create(&user)

	return
}

