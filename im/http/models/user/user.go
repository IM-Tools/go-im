/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import (
	"go_im/pkg/model"
)

type Users struct {
	ID              uint64 `json:"id"`
	Email           string `valid:"email" json:"email"`
	Password        string `valid:"password"`
	Avatar          string `json:"avatar"`
	Name            string `json:"name"`
	OauthType       int
	OauthId         string
	CreatedAt       string `json:"created_at"`
	PasswordComfirm string ` gorm:"-" valid:"password_comfirm"`
	Bio string `json:"bio"`
}

// 当前登录用户
var AuthUser *Users

func (a Users) GetAvatar() string {
	if a.Avatar == "" {
		return "https://learnku.com/users/27407"
	}
	return a.Avatar
}

//设置用户上下线状态
func SetUserStatus(id uint64 ,status int )  {

	model.DB.Model(&Users{}).Where("id=?",id).Update("status",status)
}
