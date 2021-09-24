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
	messageModel "go_im/im/http/models/msg"
	"go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"go_im/pkg/mq"
	"log"
	"strconv"
	"time"
)

//group message insert db
func PutGroupData(msg *Msg,is_read int,channel_type int) {
	channel_a := helpler.ProduceChannelGroupName(strconv.Itoa(msg.ToId))
	fid := uint64(msg.FromId)
	tid := uint64(msg.ToId)
	user := messageModel.ImMessage{FromId:fid,
		ToId: tid,
		Msg: msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05"),
		Channel:channel_a,IsRead: is_read,MsgType: msg.MsgType,ChannelType: channel_type}
	model.DB.Create(&user)
	return
}
func MqPersonalPublish(msg []byte,to_id int)  {
	ch,err := mq.RabbitMq.Channel()
	if err!= nil {
		log.Fatal(err)
	}
	defer ch.Close()
	err = ch.Publish(
		"",  // exchange
		"personal_"+strconv.Itoa(to_id), // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   string(1),
			Type:        "AgentJob",
			Body: msg  ,
		})
	if err != nil {
		log.Fatalf("发送错误")
	}
}

func MqGroupPublish(msg []byte,to_id int)  {
	ch,err := mq.RabbitMq.Channel()
	if err!= nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.Publish(
		"",  // exchange
		"group_"+strconv.Itoa(to_id), // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   string(1),
			Type:        "AgentJob",
			Body: msg  ,
		})
	if err != nil {
		log.Fatalf("发送错误")
	}
	return
}

func MqPersonalConsumption(conn *ImClient,user_id int64)  {
	ch,err :=mq.RabbitMq.Channel()
	if err!=nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"personal_"+strconv.Itoa(int(user_id)), // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

func MqGroupConsumption(conn *ImClient,user_id int64)  {
	ch,err :=mq.RabbitMq.Channel()
	if err!=nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"group_"+strconv.Itoa(int(user_id)), // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

//The private chat insert db
func PutData(msg *Msg,is_read int,channel_type int) {
	channel_a,_ := helpler.ProduceChannelName( strconv.Itoa(msg.FromId), strconv.Itoa(msg.ToId))
	fid := uint64(msg.FromId)
	tid := uint64(msg.ToId)
	user := messageModel.ImMessage{FromId:fid,
		ToId: tid,
		Msg: msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05"),
		Channel:channel_a,IsRead: is_read,MsgType: msg.MsgType,ChannelType: channel_type}
	model.DB.Create(&user)
	return
}


func PushUserOnlineNotification(conn *ImClient,id int64)  {
	var msgList []ImMessage
	list := model.DB.Where("to_id=? and is_read=?", id, 0).Find(&msgList)
	if list.Error != nil {
		fmt.Println(list.Error)
	}
	for key, _ := range msgList {
		data, _ := json.Marshal(&Msg{Code: SendOk, Msg: msgList[key].Msg,
			FromId: msgList[key].FromId, ToId: msgList[key].ToId,
			Status: 0, MsgType: msgList[key].MsgType,ChannelType: msgList[key].ChannelType})
		conn.Send <- data
	}
}

func PushUserOfflineNotification(manager *ImClientManager,conn *ImClient)  {
	if _,ok := manager.ImClientMap[conn.ID];ok {
		id, _ := strconv.ParseInt(conn.ID, 10, 64)
		user.SetUserStatus(uint64(id), 0)
		conn.Socket.Close()
		close(conn.Send)
		delete(manager.ImClientMap, conn.ID)
	}
	//推送离线消息
	jsonMessage, _ := json.Marshal(&OnlineMsg{Code: connOut, Msg: "用户离线了" + conn.ID, ID: conn.ID,ChannelType: 3})
	for _,wsConn := range manager.ImClientMap {
		wsConn.Socket.WriteMessage(websocket.TextMessage,jsonMessage)
	}

}