/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package group

import (
	"im_app/core/http/models/group_user"
	"im_app/pkg/model"
	"time"
)

type ImGroups struct {
	ID          int64                     `json:"id"`
	UserId      int64                     `json:"user_id" gorm:"index"`
	GroupName   string                    `json:"group_name"`
	Info        string                    `json:"info"`
	CreatedAt   string                    `json:"created_at"`
	GroupAvatar string                    `json:"group_avatar"`
	Users       []group_user.ImGroupUsers `json:"users" gorm:"foreignKey:GroupId;references:ID"`
}

func (ImGroups) TableName() string {
	return "im_groups"
}
func GetGroupUserList(group_id []string) ([]ImGroups, error) {
	var group []ImGroups
	err := model.DB.Preload("Users").Where("id in (?)", group_id).Find(&group).Error
	if err != nil {
		return group, err
	}
	return group, nil
}

func Created(user_id int64, group_name string) (id int64, err error) {
	group := ImGroups{
		UserId:      user_id,
		GroupName:   group_name,
		Info:        "暂无",
		CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		GroupAvatar: "https://api.prodless.com/avatar.png",
	}
	result := model.DB.Create(&group)

	if result.Error != nil {
		return
	}
	return group.ID, nil
}

// 获取群创建者id

func GetGroupUserId(id string) (g_id int64, errs error) {
	var group ImGroups
	model.DB.Table("im_groups").Where("id=?", id).First(&group)

	if errs != nil {
		return group.ID, errs
	}
	return group.UserId, nil
}
