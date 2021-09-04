/**
  @author:panliang
  @data:2021/7/2
  @note
**/
package service

import (
	"encoding/json"
	"fmt"
	messageModel "go_im/im/http/models/msg"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"strconv"
	"time"
)


//byte -> map
func EnMessage(message []byte) (msg *Message) {
	err := json.Unmarshal([]byte(string(message)),&msg)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return
}
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

//get chat group user id
func GetGroupUid(group_id int) ([]GroupId,error) {
	var groups [] GroupId
	err := model.DB.Table("im_group_users").Where("group_id=?",group_id).Find(&groups).Error;if err != nil {
		return groups,err
	}
	return groups,nil
}



