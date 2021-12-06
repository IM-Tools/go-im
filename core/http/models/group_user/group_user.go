/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package group_user

import (
	"strconv"
	"time"

	"im_app/core/http/models/user"
	"im_app/pkg/model"
)

type ImGroupUsers struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	GroupId   int64  `json:"group_id"`
	Remark    string `json:"remark"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
}

func (ImGroupUsers) TableName() string {
	return "im_group_users"
}

func CreatedAll(user_ids map[string]string, group_id int64, u_id int64) (err error) {
	var group_users = make([]*ImGroupUsers, len(user_ids)+1)
	var userId = make([]int, len(user_ids)+1)
	userId = append(userId, int(u_id))
	for _, value := range user_ids {
		valueNum, _ := strconv.Atoi(value)
		userId = append(userId, valueNum)
	}
	var users []user.Users

	err = model.DB.Where("id in (?)", userId).Find(&users).Error
	if err != nil {
		return err
	}
	var i = 0
	for _, value := range users {
		group_users[i] = &ImGroupUsers{
			UserId:    value.ID,
			GroupId:   group_id,
			CreatedAt: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			Remark:    value.Email,
			Avatar:    value.Avatar,
			Name:      value.Name,
		}
		i++
	}

	err = model.DB.Model(&ImGroupUsers{}).Create(&group_users).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGroupUser(group_id string, user_id string) bool {
	var count int64
	model.DB.Table("im_group_users").
		Where("group_id=? and user_id=?", group_id, user_id).
		Count(&count)
	if count == 0 {
		return false
	}
	return true
}
