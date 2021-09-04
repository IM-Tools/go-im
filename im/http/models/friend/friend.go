/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package friend

import (
	 "go_im/im/http/models/user"
	"go_im/pkg/model"
	"time"
)

type ImFriends struct {
	ID uint64 `json:"id"`
	Mid uint64 `json:"mid"`
	Fid uint64 `json:"fid"`
	Status int `json:"status"`
	CreatedAt string `json:"created_at"`
	Note string `json:"note"`
	Users user.Users `json:"users" gorm:"foreignKey:Fid;references:ID"`
}

func (ImFriends ) TableName() string  {
	return "im_friends"
}

func GetFriendList(user_id uint64)  ([]ImFriends,error) {
	var friends []ImFriends
	err := model.DB.Preload("Users").Where("id=?",user_id).Find(&friends).Error;
	if err!=nil{
		return friends,err
	}
	return friends,nil
}

func AddFriends(mid uint64, fid uint64) error  {
	result := model.DB.Create(ImFriends{Mid: mid,
		Fid: fid,
		Status: 0,
		CreatedAt:time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			}).Error

	if result != nil {
		return result
	}
	return nil
}
