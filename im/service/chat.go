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
	"go_im/pkg/pool"
	"go_im/pkg/wordsfilter"
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
			jsonMessage, _ := json.Marshal(&ImOnlineMsg{Code: connOk, Msg: "用户上线啦", ID: conn.ID,ChannelType:3})
			id, _ := strconv.ParseInt(conn.ID, 10, 64)
			user.SetUserStatus(uint64(id), 1)
			manager.ImSend(jsonMessage, conn)

			//用户上线通知
			pool.AntsPool.Submit(func() {
				PushUserOnlineNotification(conn,id)
			})

		case conn := <-ImManager.Unregister:
			PushUserOfflineNotification(manager,conn)
		case message := <-ImManager.Broadcast:
			data := EnMessage(message)
			msg := new(Msg)
			err := json.Unmarshal([]byte(data.Content), &msg)
			if err != nil {
				fmt.Println(err)
			}
			jsonMessage_from, _ := json.Marshal(&Msg{Code: SendOk, Msg: msg.Msg,
				FromId: msg.FromId,
				ToId:   msg.ToId, Status:1, MsgType: msg.MsgType,ChannelType: msg.ChannelType})


			if msg.ChannelType == 1 {
				conn_id := strconv.Itoa(msg.ToId)
				if data,ok :=manager.ImClientMap[conn_id];ok {
					pool.AntsPool.Submit(func() {
						PutData(msg, 1,msg.ChannelType)
					})
					data.Send <- jsonMessage_from
				} else {
					pool.AntsPool.Submit(func() {
						PutData(msg, 0,msg.ChannelType)
					})
				}
			} else {

				//群聊消息消费
				groups,_ := GetGroupUid(msg.ToId)

				for _,value :=range groups {
					if data,ok := manager.ImClientMap[value.UserId];ok {
						pool.AntsPool.Submit(func() {
							PutGroupData(msg, 1,msg.ChannelType)
						})
						data.Send <- jsonMessage_from
					}
				}
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
		msg := new(Msg)
		err = json.Unmarshal(message, &msg)
		if err !=nil {
			fmt.Println(err)
		}

		if wordsfilter.MsgFilter(msg.Msg) {
			c.Socket.WriteMessage(websocket.TextMessage, []byte(`{"code":401,"data":"禁止发送敏感词！"}`))
			continue
		} else {
			if msg.ChannelType == 1 {
				data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
					msg.Msg, msg.FromId,msg.ToId,msg.MsgType,msg.ChannelType)

				c.Socket.WriteMessage(websocket.TextMessage, []byte(data))
			}

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

