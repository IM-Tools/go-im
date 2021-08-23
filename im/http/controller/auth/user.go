/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import "C"
import (
	"github.com/gin-gonic/gin"
	messageModel "go_im/im/http/models/msg"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"strconv"
)

type UsersController struct {
}

type UsersList struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Msg       string `json:"msg"`
	Status       int `json:"status"`
	IsRead     int `json:"is_read"`
	SendTime     string `json:"send_time"`
	SendMsg     string `json:"send_msg"`
	MsgTotal     int `json:"msg_total"`
}

func (*UsersController) GetUsersList(c *gin.Context) {
	name := c.Query("name")

	user := userModel.AuthUser
	var Users []UsersList
	//将自己信息排除掉
	query := model.DB.Model(userModel.Users{}).Where("id <> ?", user.ID)
	if len(name) > 0 {
		query = query.Where("name like ?", "%"+name+"%")
	}
	query = query.Select("id", "name", "avatar", "status", "created_at").Find(&Users)
	response.SuccessResponse(map[string]interface{}{
		"list": Users,
	}, 200).ToJson(c)
}

func (*UsersController) ReadMessage(c *gin.Context) {
	user := userModel.AuthUser
	channel_a, channel_b := helpler.ProduceChannelName(strconv.Itoa(int(user.ID)), c.Query("to_id"))
	messageModel.ReadMsg(channel_a,channel_b)
	response.SuccessResponse(gin.H{}, 200).ToJson(c)
}

