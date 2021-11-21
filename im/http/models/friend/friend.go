/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package friend

import (
	"time"

	"im_app/im/http/models/user"
	"im_app/pkg/model"
)

type ImFriends struct {
	ID        int64     `json:"id"`
	MId       int64     `json:"m_id"`
	FId       int64     `json:"f_id"`
	Status    int        `json:"status"`
	CreatedAt string     `json:"created_at"`
	Note      string     `json:"note"`
	Users     user.Users `json:"users" gorm:"foreignKey:FId;references:ID"`
}

func (ImFriends) TableName() string {
	return "im_friends"
}

func GetFriendList(user_id int64) ([]ImFriends, error) {
	var friends []ImFriends
	err := model.DB.Preload("Users").Where("m_id=?", user_id).Find(&friends).Error
	if err != nil {
		return friends, err
	}
	return friends, nil
}

func AddFriends(mid int64, fid int64) error {
	result := model.DB.Create(ImFriends{MId: mid,
		FId:       fid,
		Status:    0,
		CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	}).Error

	if result != nil {
		return result
	}
	return nil
}
