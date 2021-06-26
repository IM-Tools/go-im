/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)



type ClientManager struct {
	Clients map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

//ws client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
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
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
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

//消息读取
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()

		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		fmt.Println(jsonMessage)
		Manager.Broadcast <- jsonMessage
	}
}

//关闭连接
func (c *Client) Write() {
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
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}


