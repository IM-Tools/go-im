/**
  @author:panliang
  @data:2021/6/30
  @note
**/
package msg

import (
	"fmt"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/model"
)

type ImMessage struct {
	ID        uint64 `json:"id"`
	Msg       string `json:"msg"`
	CreatedAt string  `json:"created_at"`
	FromId uint64 `json:"user_id"`
	ToId uint64 `json:"send_id"`
	Channel string `json:"channel"`
	Status int `json:"status"`
	IsRead     int `json:"is_read"`
	MsgType int `json:"msg_type"`
	ChannelType int  `json:"channel_type"`
	//Users userModel.Users `json:"users" gorm:"foreignKey:ID;references:FromId"`
	Users userModel.Users `json:"users,omitempty" gorm:"foreignKey:FromId;references:ID"`
}

func (ImMessage) TableName() string {
	return "im_messages"
}

// 获取离线消息列表
func GetOfflineMessage(id uint64) (msg *[]ImMessage)  {
	list := model.DB.Where("id=?",id).Find(&msg)
	if list.Error != nil {
		fmt.Println(list.Error)
	}
	return msg
}
func ReadMsg(channel_a string,channel_b string)  {
	model.DB.Model(&ImMessage{}).Where("channel = ?  or channel= ?", channel_a, channel_b).Update("is_read",1)
}






