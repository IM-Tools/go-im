/**
  @author:panliang
  @data:2021/11/16
  @note
**/
package client

import (
	"github.com/gorilla/websocket"
	message2 "im_app/im/ws_test/message"
	"sync"
)

type ClientContainerHandler interface {
	NewClient(uid int64,conn *websocket.Conn) *Client
	SetDevice(uid int64,device_type int,conn *websocket.Conn)
	Start()
}

type ClientContainer struct {
	mu *sync.Mutex
	Devices map[int64] *Devices
}


type Devices struct {
	Client map[int]*Client  //客户端
	Type int `json:"type"`
}

type Client struct {
	ID     int64
	Socket *websocket.Conn
	Message chan *message2.Messages
}

func (c *Client) Read() {

}

func Write()  {

}

// 创建客户端
func (c *Client)NewClient(uid int64,conn *websocket.Conn) *Client  {
	client := new(Client)
	client.ID = uid
	client.Socket = conn
	client.Message = make(chan *message2.Messages,40)
	return client
}
//设置当前设备客户端



func (container *ClientContainer)SetDevice(uid int64,device_type int,client *Client)  {
	if data, ok := container.Devices[uid]; ok {
		if _, ok := data.Client[device_type]; ok {
			return
		}
		data.Client[device_type] = client
	} else {
		container.Devices[uid].Client[device_type] = client
	}
	return
}
func Start(c *Client)  {

}

