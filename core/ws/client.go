/**
 @author:panliang
 @data:2021/11/14
 @note
**/
package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ImClient struct {
	ID     int64           //客户端id
	Socket *websocket.Conn //
	Send   chan []byte
	Mux    sync.RWMutex
}

type ImOnlineMsg struct {
	Code        int    `json:"code,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ID          int64  `json:"id,omitempty"`
	ChannelType int    `json:"channel_type"` // 1 私聊 2 群聊 3 广播
}

// 消息结构体
type Message struct {
	Sender    int64  `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Mes       *Msg
}

// 发送的消息
type Msg struct {
	Code        int    `json:"code,omitempty"`
	FromId      int    `json:"from_id,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ToId        int    `json:"to_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	MsgType     int    `json:"msg_type,omitempty"`
	ChannelType int    `json:"channel_type"`
}
type ImMessage struct {
	ID          int64  `json:"id"`
	Msg         string `json:"msg"`
	CreatedAt   string `json:"created_at"`
	FromId      int    `json:"user_id"`
	ToId        int    `json:"send_id"`
	Channel     string `json:"channel"`
	IsRead      int    `json:"is_read"`
	MsgType     int    `json:"msg_type"`
	ChannelType int    `json:"channel_type"`
}

// 离线和上线消息
type OnlineMsg struct {
	Code        int    `json:"code,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ID          int64  `json:"id,omitempty"`
	ChannelType int    `json:"channel_type"`
}

// 定义的一些状态码
const (
	heartBeat = 0    // 心跳
	connOut   = 5000 // 离线
	connOk    = 1000 // 上线
	SendOk    = 200  // 消息投递成功
	CrowdedOk = 4001 // 已在别处登录
)

var (
	Text  = websocket.TextMessage  // 文本消息指令
	Clone = websocket.CloseMessage // 关闭指令
)

// 存储房间号
type GroupId struct {
	UserId int64 `json:"user_id"`
}

// 根据房间号
type GroupMap struct {
	GroupIds map[int]*GroupId
}

// 消息投递
func (c *ImClient) ImRead() {
	defer func() {
		ImManager.Unregister <- c
		c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			ImManager.Unregister <- c
			c.Socket.Close()
			break
		}
		c.PullMessageHandler(message)

	}
}

// 从客户端消费消息
func (c *ImClient) ImWrite() {

	defer c.Socket.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(Clone, []byte{})
				return
			}
			c.Socket.WriteMessage(Text, message)
		}
	}
}
