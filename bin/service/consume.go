/**
  @author:panliang
  @data:2021/6/30
  @note
**/
package service

import (
	"encoding/json"
	"fmt"
	messageModel "go_im/bin/http/models/msg"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"strconv"
	"time"
)

type Msg struct {
	FromId int `json:"from_id"`
	Msg string `json:"msg"`
	ToId int `json:"to_id"`
}


//将消息数据写入数据库
func PutData(data []byte) {
	msg := new(Msg)

	if err :=	json.Unmarshal([]byte(string(data)),&msg);err!= nil {
		fmt.Println(err)
	}

	channel_a,_ := helpler.ProduceChannelName( strconv.Itoa(msg.FromId), strconv.Itoa(msg.ToId))

	fid := uint64(msg.FromId)
	tid := uint64(msg.ToId)

	user := messageModel.ImMessage{FromId:fid,
		ToId: tid,
		Msg: msg.Msg,
		CreatedAt: time.Unix(time.Now().Unix(), 0,).Format("2006-01-02 15:04:05"),
		Channel: channel_a}

	model.DB.Create(&user)

	return
}


