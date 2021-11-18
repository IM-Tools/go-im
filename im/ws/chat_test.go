/**
  @author:panliang
  @data:2021/11/13
  @note
**/
package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"im_app/pkg/wordsfilter"
	"log"
)


func init()  {

}

func (manager *ImClientManager) ImStart() {
	for  {
		select {
		case conn := <-ImManager.Register:
			manager.ImSend([]byte(string("用户上线了")), conn)
		case conn := <-ImManager.Unregister:
			manager.ImSend([]byte(string("用户离线了")), conn)
		case message := <-ImManager.Broadcast:
			var conn_id = "test"
			if data, ok := manager.ImClientMap[conn_id]; ok {
				manager.ImSend(message, data)
			}
		}
	}
}

func (manager *ImClientManager) ImSend(message []byte, ignore *ImClient) {
	data, ok := manager.ImClientMap[ignore.ID]
	if ok {
		data.Send <- message
	}
}

func (c *ImClient) ImRead() {
	//关闭客户端注册 关闭socket连接
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
		if string(message) == "HeartBeat" {
			c.Socket.WriteMessage(websocket.TextMessage, []byte(`{"code":0,"data":"heartbeat ok"}`))
			continue
		}
		msg := new(Msg)
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Fatal(err)
		}

		if wordsfilter.MsgFilter(msg.Msg) {
			c.Socket.WriteMessage(websocket.TextMessage, []byte(`{"code":401,"data":"禁止发送敏感词！"}`))
			continue
		} else {
			if msg.ChannelType == 1 {
				data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
					msg.Msg, msg.FromId, msg.ToId, msg.MsgType, msg.ChannelType)
				c.Socket.WriteMessage(websocket.TextMessage, []byte(data))
			}

		}

		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		ImManager.Broadcast <- jsonMessage
	}
}


func (c *ImClient) ImWrite() {
	defer Close(c)

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

func Close(c *ImClient) error  {
	return c.Socket.Close()
}
