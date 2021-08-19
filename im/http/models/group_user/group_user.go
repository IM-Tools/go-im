/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package group_user

import (
	"go_im/im/http/models/user"
	"go_im/pkg/model"
	"strconv"
	"time"
)

type ImGroupUsers struct {
	ID uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	GroupId uint64 `json:"group_id"`
	Remark string `json:"remark"`
	Avatar string `json:"avatar"`
}



func (ImGroupUsers) TableName() string {
	return "im_group_users"
}

func CreatedAll(user_ids map[string]string,group_id uint64 ) (err error)  {
	//默认只能邀请50个人组群
	var group_users [50]*ImGroupUsers

	var userId []uint64

	for key,value := range user_ids {
		keyNum, _ := strconv.Atoi(key)
		valueNum, _ := strconv.Atoi(value)
		userId[keyNum] = uint64(valueNum)

	}
	var users []user.Users

	err = model.DB.Where("id=(?)",userId).Find(&users).Error;if err !=nil {
		return err
	}
	var i =0
	for _,value := range users {
		group_users[i] = &ImGroupUsers{
			UserId: value.ID,
			GroupId:group_id,
			CreatedAt:time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			Remark:"暂无😄",
			Avatar:value.Avatar,
		}
		i++
	}

	err = model.DB.Model(&ImGroupUsers{}).Create(group_users).Error; if err != nil {
		return err
	}
	return nil
}
