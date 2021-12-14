/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import "C"
import (
	"github.com/gin-gonic/gin"
	messageModel "im_app/core/http/models/msg"
	userModel "im_app/core/http/models/user"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/response"
	"strconv"
	"time"
)

type (
	UsersController struct{}
	UsersList       struct {
		ID            int64     `json:"id"`
		Email         string    `json:"email"`
		Avatar        string    `json:"avatar"`
		Name          string    `json:"name"`
		Msg           string    `json:"msg"`
		Status        int       `json:"status"`
		IsRead        int       `json:"is_read"`
		SendTime      string    `json:"send_time"`
		SendMsg       string    `json:"send_msg"`
		MsgTotal      int       `json:"msg_total"`
		ClientType    int       `json:"client_type"`
		Bio           int       `json:"bio"`
		Sex           int       `json:"sex"`
		LastLoginTime time.Time `gorm:"type:time" json:"last_login_time"`
	}
	NotFriendList struct {
		Name   string `json:"name"`
		ID     int64  `json:"id"`
		Avatar string `json:"avatar"`
		Status string `json:"status"`
	}
	Result struct {
		FId int64 `json:"f_id"`
	}
)

// @BasePath /api

// @Summary 根据昵称查询非好友用户列表
// @Description 根据昵称查询非好友用户列表
// @Tags 根据昵称查询非好友用户列表
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
	if len(name) < 1 {
		response.FailResponse(500, "请输入需要添加好友的昵称").ToJson(c)
		return
	}

	user := userModel.AuthUser
	subQuery := model.DB.Select("f_id").
		Group("f_id").
		Where("m_id=?", user.ID).
		Table("im_friends")

	userList, err := userModel.GetNotFriendList(subQuery, userModel.AuthUser.ID, name)
	if err != nil {
		response.FailResponse(500, "接口查询异常").ToJson(c)
		return
	}
	response.SuccessResponse(map[string]interface{}{
		"list": userList,
	}, 200).ToJson(c)
	return
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
	to_id := c.Query("to_id")
	toid, _ := strconv.Atoi(to_id)
	channel_a, channel_b := helpler.ProduceChannelName(int64(user.ID), int64(toid))
	messageModel.ReadMsg(channel_a, channel_b)
	response.SuccessResponse(gin.H{}, 200).ToJson(c)
}
