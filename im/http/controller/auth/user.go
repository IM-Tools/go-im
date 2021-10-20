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


type (
	UsersController struct {}
	UsersList struct {
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
)
// @BasePath /api

// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 获取用户列表
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param name query string false "账号"
// @Produce json
// @Success 200
// @Router /UsersList [get]
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
// @Summary 历史消息读取[废弃]
// @Description 历史消息读取
// @Tags 历史消息读取
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param voice formData file true "图片上传"
// @Produce json
// @Success 200
// @Router /ReadMessage [get]
func (*UsersController) ReadMessage(c *gin.Context) {
	user := userModel.AuthUser
	channel_a, channel_b := helpler.ProduceChannelName(strconv.Itoa(int(user.ID)), c.Query("to_id"))
	messageModel.ReadMsg(channel_a,channel_b)
	response.SuccessResponse(gin.H{}, 200).ToJson(c)
}

func (*UsersController) FriendList(C *gin.Context)  {

}

