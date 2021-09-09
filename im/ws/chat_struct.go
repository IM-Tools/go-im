/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package ws

import "github.com/gorilla/websocket"

type ImClientManager struct {
	ImClientMap map[string]*ImClient //存放在线用户连接
	Broadcast  chan []byte           //收集消息分发到客户端
	Register   chan *ImClient        //新注册的长连接
	Unregister chan *ImClient        //已注销的长连接
}

type ImClient struct {
	ID string
	Socket *websocket.Conn
	Send chan []byte
}

//客户度啊管理器
var ImManager = ImClientManager{
	ImClientMap:make(map[string]*ImClient),
	Broadcast:  make(chan []byte),
	Register:   make(chan *ImClient),
	Unregister: make(chan *ImClient),
}

type ImOnlineMsg struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	ID   string `json:"id,omitempty"`
	ChannelType   int   `json:"channel_type"` //1 私聊 2 群聊 3 广播
}

//消息结构体
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

//发送的消息
type Msg struct {
	Code    int    `json:"code,omitempty"`
	FromId  int    `json:"from_id,omitempty"`
	Msg     string `json:"msg,omitempty"`
	ToId    int    `json:"to_id,omitempty"`
	Status  int    `json:"status,omitempty"`
	MsgType int    `json:"msg_type,omitempty"`
	ChannelType   int   `json:"channel_type"`
}
type ImMessage struct {
	ID        uint64 `json:"id"`
	Msg       string `json:"msg"`
	CreatedAt string `json:"created_at"`
	FromId    int    `json:"user_id"`
	ToId      int    `json:"send_id"`
	Channel   string `json:"channel"`
	IsRead    int    `json:"is_read"`
	MsgType   int    `json:"msg_type"`
	ChannelType   int   `json:"channel_type"`
}

//离线和上线消息
type OnlineMsg struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	ID   string `json:"id,omitempty"`
	ChannelType   int   `json:"channel_type"`
}

//定义的一些状态码

const (
	connOut = 5000 //断开链接
	connOk  = 1000 //连接成功
	SendOk  = 200  //发送成功
)

//存储房间号
type GroupId struct {
	UserId string `json:"user_id"`
}

//根据房间号
type GroupMap struct {
	GroupIds map[int]*GroupId
}


