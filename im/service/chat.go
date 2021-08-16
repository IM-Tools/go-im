/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go_im/im/http/models/user"
	"go_im/pkg/model"
	"strconv"
	"sync"
)

var mutexKey sync.Mutex

func (manager *ImClientManager) ImStart() {
	for  {
		select {
		case conn := <-ImManager.Register:
			//新增锁 防止并发写
			mutexKey.Lock()
			manager.ImClientMap[conn.ID] = &ImClient{ID: conn.ID,Socket: conn.Socket,Send:conn.Send}
			mutexKey.Unlock()
			jsonMessage, _ := json.Marshal(&ImOnlineMsg{Code: connOk, Msg: "用户上线啦", ID: conn.ID})
			id, _ := strconv.ParseInt(conn.ID, 10, 64)
			user.SetUserStatus(uint64(id), 1)
			manager.ImSend(jsonMessage, conn)

			//用户上线通知
			go func() {
				var msgList []ImMessage
				list := model.DB.Where("to_id=? and is_read=?", id, 0).Find(&msgList)
				if list.Error != nil {
					fmt.Println(list.Error)
				}
				for key, _ := range msgList {
					data, _ := json.Marshal(&Msg{Code: SendOk, Msg: msgList[key].Msg,
						FromId: msgList[key].FromId, ToId: msgList[key].ToId,
						Status: 0, MsgType: msgList[key].MsgType})
					conn.Send <- data
				}
			}()
		case conn := <-ImManager.Unregister:
			if _,ok :=manager.ImClientMap[conn.ID];ok {
				id, _ := strconv.ParseInt(conn.ID, 10, 64)
				user.SetUserStatus(uint64(id), 0)
				//关闭连接释放资源
				conn.Socket.Close()
				close(conn.Send)
				delete(manager.ImClientMap, conn.ID)
				jsonMessage, _ := json.Marshal(&OnlineMsg{Code: connOut, Msg: "用户离线了" + conn.ID, ID: conn.ID})
				manager.ImSend(jsonMessage, conn)
			}
		case message := <-ImManager.Broadcast:
			data := EnMessage(message)
			msg := new(Msg)
			err := json.Unmarshal([]byte(data.Content), &msg)
			if err != nil {
				fmt.Println(err)
			}
			jsonMessage_from, _ := json.Marshal(&Msg{Code: SendOk, Msg: msg.Msg,
				FromId: msg.FromId,
				ToId:   msg.ToId, Status: 0, MsgType: msg.MsgType})
			conn_id := strconv.Itoa(msg.ToId)

			if data,ok :=manager.ImClientMap[conn_id];ok {
				go PutData(msg, 1)
				data.Send <- jsonMessage_from
			} else {
				go PutData(msg, 0)
			}
		}
	}

}

func (manager *ImClientManager) ImSend(message []byte, ignore *ImClient) {
	data,ok := manager.ImClientMap[ignore.ID]
	fmt.Println(ignore.ID)
	if ok {
		data.Send <- message
	}
}


//消息投递
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
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		ImManager.Broadcast <- jsonMessage
	}
}

//从客户端消费消息
func (c *ImClient) ImWrite() {
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
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

