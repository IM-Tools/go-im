/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import "go_im/pkg/model"

type Users struct {
	ID int64
	Email string  `valid:"email"`
	Password string  `valid:"password"`
	Avatar,Name string
	OauthType int
	OauthId string
	CreatedAt string
	PasswordComfirm string ` gorm:"-" valid:"password_comfirm"`
}


func (a Users) GetAvatar() string {
	if a.Avatar =="" {
		return "https://learnku.com/users/27407"
	}
	return a.Avatar
}


func GetUsers(OauthId string) (Users,error) {

	var user Users

	if err := model.DB.Where("oauth_id =?",OauthId).First(&user).Error; err != nil {
		return user,err
	}

	return user,nil
}


