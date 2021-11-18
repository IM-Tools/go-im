/**
 @author:panliang
 @data:2021/11/14
 @note
**/
package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"

	"im_app/pkg/wordsfilter"
	"im_app/pkg/zaplog"
)

type ImClient struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

type ImOnlineMsg struct {
	Code        int    `json:"code,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ID          string `json:"id,omitempty"`
	ChannelType int    `json:"channel_type"` // 1 私聊 2 群聊 3 广播
}

// 消息结构体
type Message struct {
	Sender    string `json:"sender,omitempty"`
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
	ID          uint64 `json:"id"`
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
	ID          string `json:"id,omitempty"`
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
	UserId string `json:"user_id"`
}

// 根据房间号
type GroupMap struct {
	GroupIds map[int]*GroupId
}

// 消息投递
func (c *ImClient) ImRead() {
	// 关闭客户端注册 关闭socket连接
	defer func() {
		ImManager.Unregister <- c
		c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			ImManager.Unregister <- c
			c.Socket.Close()
			// 使用 break 打断for循环写入
			break
		}
		if string(message) == "HeartBeat" {
			c.Socket.WriteMessage(Text, []byte(`{"code":0,"data":"heartbeat ok"}`))
			continue
		}
		msg := new(Msg)

		if len(message) < 0 {
			continue
		}

		err = json.Unmarshal(message, &msg)
		if err != nil {
			zaplog.ZapLogger.Named("常")
			continue
		}

		if wordsfilter.MsgFilter(msg.Msg) {
			c.Socket.WriteMessage(Text, []byte(`{"code":401,"data":"禁止发送敏感词！"}`))
			continue
		}

		if msg.ChannelType == 1 {
			data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
				msg.Msg, msg.FromId, msg.ToId, msg.MsgType, msg.ChannelType)
			c.Socket.WriteMessage(Text, []byte(data))
		}

		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Mes: msg})
		ImManager.Broadcast <- jsonMessage

	}
}

// 从客户端消费消息
func (c *ImClient) ImWrite() {
	// 关闭socket连接
	defer c.Socket.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 关闭
				c.Socket.WriteMessage(Clone, []byte{})
				return
			}
			c.Socket.WriteMessage(Text, message)
		}
	}
}
