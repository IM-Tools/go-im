/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import (
	"github.com/gin-gonic/gin"
	userModel "go_im/bin/http/models/user"
	"go_im/pkg/model"
	"go_im/pkg/response"
)

type UsersController struct {
	
}

type UsersList struct {
	ID uint64 `json:"id"`
	Email string  `json:"email"`
	Avatar string `json:"avatar"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

var UsersModel userModel.Users

//获取用户列表
func (*UsersController)GetUsersList(c *gin.Context)  {
	name := c.Query("name")
	var Users []UsersList

	query := model.DB.Model(UsersModel)

	if len(name) >0 {
		query = query.Where("name like ?","%"+name+"%")
	}
	query = query.Select("id","name","avatar","status","created_at").Find(&Users)

	response.SuccessResponse(map[string]interface{}{
		"list":Users,
	},200).ToJson(c)
}
