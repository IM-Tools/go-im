/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package group_user

import (
	"strconv"
	"time"

	"im_app/im/http/models/user"
	"im_app/pkg/model"
)

type ImGroupUsers struct {
	ID        uint64 `json:"id"`
	UserId    uint64 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	GroupId   uint64 `json:"group_id"`
	Remark    string `json:"remark"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
}

func (ImGroupUsers) TableName() string {
	return "im_group_users"
}

func CreatedAll(user_ids map[string]string, group_id uint64, u_id uint64) (err error) {
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
