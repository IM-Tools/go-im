/**
  @author:panliang
  @data:2021/11/18
  @note
**/
package ws

import (
	"encoding/json"
	"fmt"
	"im_app/pkg/wordsfilter"
	"im_app/pkg/zaplog"
	"strconv"

	"github.com/gorilla/websocket"
)

// 消息投递下发
func (manager *ImClientManager)LaunchMessage(msg_byte []byte) {
	// 消息传输协议可以优化 可以使用自定义二进制协议
	// json传输协议 格式转换比较消耗性能
	// 当然这个方法也可以优惠避免多次转换
	message := EnMessage(msg_byte)
	msg     := DeMessage(message.Mes)

	if message.Mes.ChannelType == 1 {
		id := strconv.Itoa(message.Mes.ToId)
		if conn, ok := manager.ImClientMap[id]; ok {
			zaplog.Info("数据入库:",message.Mes)
			PutData(message.Mes, 1, 1)
			conn.Send <- msg
		} else {
			PutData(message.Mes, 0, 1)
			MqPersonalConsumption(conn,int64(message.Mes.ToId))
		}
		return
	}
	// 群聊 获取群聊用户信息可以做数据缓存 --
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
func (manager *ImClientManager)LaunchOnlineMsg(id string) {
	message, _ := json.Marshal(&ImOnlineMsg{Code: connOk, Msg: "用户上线啦", ID: id, ChannelType: 3})
	for _, conn := range manager.ImClientMap {
		conn.Socket.WriteMessage(websocket.TextMessage, message)
	}
	return
}


// 消息处理方法
func (c *ImClient)PullMessageHandler(message []byte)  {

	if len(message) < 0 {
		return
	}
	if string(message) == "HeartBeat" {
		LaunchTicklingAckMsg([]byte(`{"code":0,"data":"heartbeat ok"}`),c)
		return
	}

	msg := new(Msg)
	err := json.Unmarshal(message, &msg)
	if err != nil {
		zaplog.Error("消息解析异常-----",err)
		return
	}

	if wordsfilter.MsgFilter(msg.Msg) {
		LaunchTicklingAckMsg([]byte(`{"code":401,"data":"禁止发送敏感词！"}`),c)
		return
	}

	if msg.ChannelType == 1 {
		data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
			msg.Msg, msg.FromId, msg.ToId, msg.MsgType, msg.ChannelType)
		LaunchTicklingAckMsg([]byte(data),c)
	}
	messageByte, _ := json.Marshal(&Message{Sender: c.ID, Mes: msg})

	ImManager.Broadcast <- messageByte
	return
}


//
func LaunchTicklingAckMsg(msg []byte,conn *ImClient)  {
	conn.Socket.WriteMessage(websocket.TextMessage,msg)
	return
}
