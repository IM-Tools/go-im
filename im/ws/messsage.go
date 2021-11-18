/**
  @author:panliang
  @data:2021/11/18
  @note
**/
package ws

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

// 消息投递下发
func LaunchMessage(msg_byte []byte, manager *ImClientManager) {
	// 单聊
	message := EnMessage(msg_byte)
	message.Mes.Code = 200
	msg := jsonMessage(message.Mes)

	if message.Mes.ChannelType == 1 {
		id := strconv.Itoa(message.Mes.ToId)
		if conn, ok := manager.ImClientMap[id]; ok {
			PutData(message.Mes, 0, 1)
			conn.Send <- msg
		}
		return
	}
	// 群聊
	groups, _ := GetGroupUid(message.Mes.ToId)
	PutGroupData(message.Mes, 1, message.Mes.ChannelType)

	for _, value := range groups {
		if data, ok := manager.ImClientMap[value.UserId]; ok {
			MqGroupPublish(msg, message.Mes.ToId)
			data.Send <- msg
		}
	}
	return
}

// 上线消息通知
func LaunchOnlineMsg(id string, manager *ImClientManager) {
	message, _ := json.Marshal(&ImOnlineMsg{Code: connOk, Msg: "用户上线啦", ID: id, ChannelType: 3})
	for _, conn := range manager.ImClientMap {
		conn.Socket.WriteMessage(websocket.TextMessage, message)
	}
	return
}

func jsonMessage(message *Msg) []byte {
	byte_msg, err := json.Marshal(message)
	if err != nil {
		log.Fatal("异常", err)
	}
	return byte_msg
}
