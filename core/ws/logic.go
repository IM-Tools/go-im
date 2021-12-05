/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	messageModel "im_app/core/http/models/msg"
	"im_app/core/http/models/user"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/mq"
	"im_app/pkg/zaplog"
	"log"
	"net"
	"strconv"
	"time"
)

// 这个文件里面代码很乱找时间梳理一下

// group message insert db
func PutGroupData(msg *Msg, is_read int, channel_type int) {
	channel_a := helpler.ProduceChannelGroupName(strconv.Itoa(msg.ToId))
	fid := int64(msg.FromId)
	tid := int64(msg.ToId)
	user := messageModel.ImMessage{FromId: fid,
		ToId:      tid,
		Msg:       msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		Channel:   channel_a, IsRead: is_read, MsgType: msg.MsgType, ChannelType: channel_type}
	model.DB.Create(&user)
	return
}

//私人消息入队
func MqPersonalPublish(msg []byte, to_id int) {
	ch, err := mq.RabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	err = ch.Publish(
		"",                              // exchange
		"personal_"+strconv.Itoa(to_id), // routing key
		false,                           // mandatory
		false,                           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   string(1),
			Type:        "AgentJob",
			Body:        msg,
		})
	if err != nil {
		log.Fatalf("发送错误")
	}
}

//组消息入库
func MqGroupPublish(msg []byte, to_id int) {
	ch, err := mq.RabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.Publish(
		"",                           // exchange
		"group_"+strconv.Itoa(to_id), // routing key
		false,                        // mandatory
		false,                        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   string(1),
			Type:        "AgentJob",
			Body:        msg,
		})
	if err != nil {
		log.Fatalf("发送错误")
	}
	return
}

// 私人消息同步消费
func MqPersonalConsumption(conn *ImClient, user_id int64) {
	ch, err := mq.RabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"personal_"+strconv.Itoa(int(user_id)), // name
		false,                                  // durable
		false,                                  // delete when unused
		false,                                  // exclusive
		false,                                  // no-wait
		nil,                                    // arguments
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		conn.Send <- msg.Body
	}
	return
}

func MqGroupConsumption(conn *ImClient, user_id int64) {
	ch, err := mq.RabbitMq.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"group_"+strconv.Itoa(int(user_id)), // name
		false,                               // durable
		false,                               // delete when unused
		false,                               // exclusive
		false,                               // no-wait
		nil,                                 // arguments
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		conn.Send <- msg.Body
	}
}

// The private chat insert db
func PutData(msg *Msg, is_read int, channel_type int) {

	channel_a, _ := helpler.ProduceChannelName(int64(msg.FromId), int64(msg.ToId))
	fid := int64(msg.FromId)
	tid := int64(msg.ToId)
	message := messageModel.ImMessage{FromId: fid,
		ToId:      tid,
		Msg:       msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		Channel:   channel_a, IsRead: is_read, MsgType: msg.MsgType, ChannelType: channel_type}
	result := model.DB.Create(&message)
	if result.Error != nil {
		zaplog.Error("数据拆入异常")
	}

	return
}

func PushUserOnlineNotification(conn *ImClient, id int64) {
	var msgList []ImMessage
	list := model.DB.Where("to_id=? and is_read=?", id, 0).Find(&msgList)
	if list.Error != nil {
		zaplog.Error("异常", list.Error)
	}
	for key, _ := range msgList {
		data, _ := json.Marshal(&Msg{Code: SendOk, Msg: msgList[key].Msg,
			FromId: msgList[key].FromId, ToId: msgList[key].ToId,
			Status: 0, MsgType: msgList[key].MsgType, ChannelType: msgList[key].ChannelType})
		conn.Send <- data
	}
}

func PushUserOfflineNotification(manager *ImClientManager, conn *ImClient) {
	if _, ok := manager.ImClientMap[conn.ID]; ok {

		user.SetUserStatus(conn.ID, 0)
		// conn.Socket.Close()
	}
	// 推送离线消息
	jsonMessage, _ := json.Marshal(&OnlineMsg{Code: connOut, Msg: "用户离线了", ID: conn.ID, ChannelType: 3})
	for _, wsConn := range manager.ImClientMap {
		wsConn.Socket.WriteMessage(websocket.TextMessage, jsonMessage)
	}

}

// 用户被挤下线
func CrowdedOffline(user_id int64) {
	manager := new(ImClientManager)
	if conn, ok := manager.ImClientMap[user_id]; ok {
		jsonMessage, _ := json.Marshal(&ImOnlineMsg{Code: CrowdedOk, Msg: "账号已在别处登录", ID: conn.ID, ChannelType: 3})

		fmt.Println(jsonMessage)
		conn.Send <- jsonMessage

		conn.Socket.Close()
		close(conn.Send)
		delete(manager.ImClientMap, conn.ID)
	}
}

// byte -> struct
func EnMessage(message []byte) (msg *Message) {
	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return
}

func DeMessage(message *Msg) []byte {
	byte_msg, err := json.Marshal(message)
	if err != nil {
		log.Fatal("异常", err)
	}
	return byte_msg
}

// get chat group user id
func GetGroupUid(group_id int) ([]GroupId, error) {
	var groups []GroupId
	err := model.DB.Table("im_group_users").Where("group_id=?", group_id).Find(&groups).Error
	if err != nil {
		return groups, err
	}
	return groups, nil
}

func IsIpPort(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		// 匹配上
		return true
	}
}
