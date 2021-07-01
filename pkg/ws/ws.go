/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package ws

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
)


//客户端管理
type ClientManager struct {
	Clients map[*Client]bool //存放ws长链接
	Broadcast  chan []byte //收集消息分发到client
	Register   chan *Client //新创建的长链接
	Unregister chan *Client //注销的长链接
}

//ws client
type Client struct {
	ID     string //客户端id
	Socket *websocket.Conn //长链接
	Send   chan []byte //需要发送的消息
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type Msg struct {
	Code int `json:"code,omitempty"`
	FromId int `json:"from_i,omitempty"`
	Msg string `json:"msg,omitempty"`
	ToId int `json:"to_id,omitempty"`
	Status int `json:"status,omitempty"`
}

const (
	connOut = -0
	okStatus = 0 //连接成功
	SendOk = 200 //发送成功
)

// 创建客户端管理器
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

//启动websocket
func (manager *ClientManager) Start() {
	for {
		select {
		//如果有新的链接通过channel把链接传递给conn
		case conn := <-manager.Register:
			//将客户端的链接设置为true
			manager.Clients[conn] = true
			//把返回链接成功的消息json格式化
			jsonMessage, _ := json.Marshal(&Msg{Code:okStatus ,Msg: "/A new socket has connected."})
			//调用客户端方法发送消息
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Msg{Code:connOut,Msg: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
			//广播
		case message := <-manager.Broadcast:
			//msg :=new(Msg)
			//
			//msgs := gjson.Get(string(message),"msg")
			//msg.Msg = msgs.Str
			//fromid := gjson.Get(string(message),"from_id")
			//toid := gjson.Get(string(message),"to_id")
			//msg.FromId = fromid.Index
			//msg.ToId = fromid.Index
			//fid_channel := SendUuid(fromid.Num)
			//toid_channel := SendUuid(toid.Num)
			//////遍历已经连接的客户端，把消息发送给他们
			for conn := range manager.Clients {
			//	if conn.ID == fid_channel {
			//		msg.Status = 1
			//		data,_ :=Encode(msg)
			//		fmt.Println("字节数组",data)
			//
			//
			//		select {
			//		case
			//		conn.Send <- data:
			//		default:
			//			close(conn.Send)
			//			delete(manager.Clients, conn)
			//
			//		}
			//
			//
			//	}
			//	if conn.ID == toid_channel {
			//		msg.Status = 0
			//		data,_ :=	Encode(msg)
			//		fmt.Println("字节数组",data)
			//		select {
			//		case
			//			conn.Send <- data:
			//		default:
			//			close(conn.Send)
			//			delete(manager.Clients, conn)
			//
			//		}
			//	}else {

					select {

					case conn.Send <-message:

					default:
						close(conn.Send)
						delete(manager.Clients, conn)
					}

			}
		}
	}
}

// 发送消息到客户端
func (manager *ClientManager) Send(message []byte, ignore *Client) {

	for conn := range manager.Clients {

		if conn != ignore {
			conn.Send <- message
		}
	}
}

//消息投递
func (c *Client) Read() {

	//关闭客户端注册 关闭socket连接
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		fmt.Println(message)

		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}

		fmt.Println("投递客户端消息",message)


		//将数据投递到对应的客户端
		Manager.Broadcast <- message

	}
}

//从客户端消费消息
func (c *Client) Write() {
	//关闭socket连接
	defer func() {
		c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

//将字节数组转化为结构体
func EnMessage(message *[]byte) (msg *Msg) {

	err := json.Unmarshal(*message,&msg)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	//msg = new(Msg)
	//
	//if err :=	json.Unmarshal([]byte(string(*message)),&msg);err!= nil {
	//	fmt.Println(err)
	//}
	return
}

//数据
func SendUuid(id float64) (uuids string) {
	  uuids = "channel_"+ strconv.Itoa(int(id))
	  return
}


func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

