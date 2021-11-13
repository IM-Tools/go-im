/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package friend_record

import (
	"go_im/pkg/model"
	"os/user"
	"strconv"
	"time"
)

type ImFriendRecord struct {
	ID          uint64    `json:"id"`
	UserId      uint64    `json:"user_id"`
	Fid         uint64    `json:"f_id"`
	Status      int64     `json:"status"`
	CreatedAt   string    `json:"created_at"`
	Information string    `json:"information"`
	Users       user.User `json:"users" gorm:"foreignKey:UserId;references:ID"`
}

func (*ImFriendRecord) TableName() string {
	return "im_friend_record"
}

func GetFriendRecordList(user_id uint64) ([]ImFriendRecord, error) {
	var friends []ImFriendRecord
	err := model.DB.Preload("Users").Where("Fid=?", user_id).Find(&friends).Error
	if err != nil {
		return friends, err
	}
	return friends, nil
}

func AddRecords(user_id uint64, f_id string, information string) error {
	friend_id, _ := strconv.Atoi(f_id)

	result := model.DB.Create(ImFriendRecord{UserId: user_id,
		Fid:         uint64(friend_id),
		Status:      0,
		CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		Information: information,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil

}
