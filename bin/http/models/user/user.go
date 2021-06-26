/**
  @author:panliang
  @data:2021/6/21
  @note
**/
package user

import (
	"fmt"
	"go_im/pkg/model"
)

type Users struct {
	ID uint64 `json:"id"`
	Email string  `valid:"email" json:"email"`
	Password string  `valid:"password"`
	Avatar string `json:"avatar"`
	Name  string `json:"name"` 
	OauthType int
	OauthId string
	CreatedAt string `json:"created_at"`
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

//获取用户列表
func GetUserList() (Users,error){
   var user Users
	err := model.DB.Select("id","name","avatar","status","created_at").Find(&user)
	if err!= nil {
		fmt.Println(err)
	}
	return user,nil
}


