/**
 @author:panliang
 @data:2021/11/14
 @note
**/
package ws

import "github.com/gorilla/websocket"

type ConnContainerHandler interface {
	Run()
	Read()
	Write(c *Client)
}

// 容器
type ConnContainer struct {
	id   int64
	Conn *Devices
}

// 设备管理
type Devices struct {
	Client map[int64]*Client
	Size   int64
	Close  bool
}

// 客户端连接实例
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}
