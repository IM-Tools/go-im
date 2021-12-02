/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package friend_record

import (
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"im_app/pkg/model"
	"strconv"
	"time"
)

type ImFriendRecords struct {
	ID          int64 `json:"id"`
	UserId      int64 `json:"user_id"`
	FId         int64 `json:"f_id"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"created_at"`
	Information string `json:"information"`
	User ImUsers `json:"users" gorm:"foreignKey:UserId;references:ID"`
}



type ImUsers struct {
	Name string `json:"name"`
	Avatar          string `json:"avatar"`
	ID      int64 `json:"id"`
}


func (*ImFriendRecords) TableName() string {
	return "im_friend_records"
}

func (f *ImFriendRecords) AfterFind(tx *gorm.DB) (err error) {
	f.CreatedAt = carbon.Parse(f.CreatedAt).SetLocale("zh-CN").DiffForHumans()
	return
}


func GetFriendRecordList(user_id int64) ([]ImFriendRecords, error) {
	var friends []ImFriendRecords
	err := model.DB.Preload("User").
		Where("f_id=? and status=0", user_id).
		Order("created_at desc").
		Find(&friends).Error
	if err != nil {
		return friends, err
	}
	return friends, nil
}

func AddRecords(user_id int64, f_id string, information string) error {
	friend_id, _ := strconv.Atoi(f_id)

	result := model.DB.Create(&ImFriendRecords{UserId: user_id,
		FId:         int64(friend_id),
		Status:      0,
		CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		Information: information,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil

}
