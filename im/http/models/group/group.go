/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package group

import (
	"go_im/pkg/model"
)

type ImGroups struct {
	ID uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	GroupName string `json:"group_name"`
	Info string `json:"info"`
	CreatedAt string `json:"created_at"`
	GroupAvatar string `json:"group_avatar"`
}

func GetGroupList(id uint64) ([]ImGroups,error) {
	var group []ImGroups
	if err := model.DB.Where("user_id=?",id).Find(&group).Error;err!=nil{
		return group,err
	}
	return group,nil
}
