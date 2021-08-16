/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package group

import (
	"go_im/im/http/models/user"
	"go_im/pkg/model"
)
type ImGroups struct {
	ID uint64 `json:"id"`
	UserId uint64 `json:"user_id" gorm:"index"`
	GroupName string `json:"group_name"`
	Info string `json:"info"`
	CreatedAt string `json:"created_at"`
	GroupAvatar string `json:"group_avatar"`
	Users []user.Users `json:"users" gorm:"foreignKey:UserId;references:ID"`
}
func (ImGroups) TableName() string {
	return "im_groups"
}
func GetGroupUserList(id uint64) ([]ImGroups,error) {
	var group []ImGroups
	 err := model.DB.Joins("Users").Where("user_id=?",id).Find(&group).Error;
	 if err!=nil{
		 return group,err
	 }
	return group,nil
}
