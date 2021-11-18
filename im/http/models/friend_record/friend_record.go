/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package friend_record

import (
	"strconv"
	"time"

	"im_app/pkg/model"
)

type ImFriendRecords struct {
	ID          uint64 `json:"id"`
	UserId      uint64 `json:"user_id"`
	FId         uint64 `json:"f_id"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"created_at"`
	Information string `json:"information"`
	// Users       user.User `json:"users" gorm:"foreignKey:UserId;references:ID"`
}

func (*ImFriendRecords) TableName() string {
	return "im_friend_records"
}

func GetFriendRecordList(user_id uint64) ([]ImFriendRecords, error) {
	var friends []ImFriendRecords
	err := model.DB.Preload("Users").Where("f_id=?", user_id).Find(&friends).Error
	if err != nil {
		return friends, err
	}
	return friends, nil
}

func AddRecords(user_id uint64, f_id string, information string) error {
	friend_id, _ := strconv.Atoi(f_id)

	result := model.DB.Create(&ImFriendRecords{UserId: user_id,
		FId:         uint64(friend_id),
		Status:      0,
		CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		Information: information,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil

}
