/**
  @author:panliang
  @data:2021/12/20
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im_app/core/http/models/group"
	"im_app/core/http/models/session"
	user2 "im_app/core/http/models/user"
	"im_app/core/http/validates"
	"im_app/pkg/model"
	"im_app/pkg/response"
	"strconv"
	"time"
)

type SessionController struct {
}

// @BasePath /api

// @Summary 获取会话列表
// @Description 获取会话列表
// @Tags 获取会话列表
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Produce json
// @Success 200
// @Router /GetSessionList [post]
func (*SessionController) GetSessionList(c *gin.Context) {
	var list []session.ImSessions
	err := model.DB.Table("im_sessions").Where("m_id=?", user2.AuthUser.ID).
		Order("top_status desc").
		Order("top_time desc").
		Find(&list).Error
	if err != nil {
		response.ErrorResponse(500, "error").ToJson(c)
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}

// @BasePath /api

// @Summary 添加会话信息
// @Description 添加会话信息
// @Tags 添加会话信息
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param f_id formData string true "好友id或者群聊id"
// @Param channel_type formData string true "会话类型"
// @Produce json
// @Success 200
// @Router /AddSession [post]
func (*SessionController) Create(c *gin.Context) {

	user_id := c.PostForm("f_id")
	channel_type := c.PostForm("channel_type")
	_user := validates.AddSessionFrom{
		UserId:      user_id,
		ChannelType: channel_type,
	}
	errs := validates.ValidateAddSession(_user)
	fmt.Println(_user)
	if len(errs) > 0 {
		response.FailResponse(401, "error", errs).ToJson(c)
		return
	}

	f_id, _ := strconv.Atoi(user_id)
	c_type, _ := strconv.Atoi(channel_type)

	if c_type == 1 {
		var user user2.Users

		fmt.Println(2)
		err := model.DB.Table("im_users").Where("id=?", user_id).First(&user).Error
		if err != nil {
			response.FailResponse(500, "用户数据不存在")
			return
		}
		fmt.Println(user)
		sessionData := session.ImSessions{
			Name:        user.Name,
			MId:         user2.AuthUser.ID,
			FId:         int64(f_id),
			CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			TopTime:     time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			TopStatus:   0,
			ChannelType: c_type,
			Avatar:      user.Avatar,
			Status:      0,
		}
		model.DB.Model(&session.ImSessions{}).Create(&sessionData)
		response.SuccessResponse(&sessionData).ToJson(c)
		return

	} else {

		var groups group.ImGroups
		err := model.DB.Table("im_groups").Where("id=?", user_id).First(&groups).Error
		if err != nil {
			response.FailResponse(500, "群聊数据不存在")
			return
		}
		sessionData := session.ImSessions{
			Name:        groups.GroupName,
			MId:         user2.AuthUser.ID,
			FId:         int64(f_id),
			CreatedAt:   time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			TopTime:     time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
			TopStatus:   0,
			ChannelType: c_type,
			Avatar:      groups.GroupAvatar,
			Status:      0,
		}
		model.DB.Model(&session.ImSessions{}).Create(&sessionData)

		response.SuccessResponse(&sessionData).ToJson(c)
		return

	}
}

// @BasePath /api

// @Summary 会话置顶
// @Description 会话置顶
// @Tags 会话置顶
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param f_id formData string true "好友id或者群聊id"
// @Param top_status formData string true "0 正常 1置顶"
// @Produce json
// @Success 200
// @Router /DelSession [post]
func (*SessionController) SetSessionTop(c *gin.Context) {

	user_id := c.PostForm("f_id")
	top_status := c.PostForm("top_status")
	_status, _ := strconv.Atoi(top_status)
	model.DB.Model(&session.ImSessions{}).Where("m_id=? and f_id=? ", user2.AuthUser.ID, user_id).
		Updates(map[string]interface{}{
			"top_status": _status, "TopTime": time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		})
	response.SuccessResponse(c)
	return
}

// @BasePath /api

// @Summary 删除会话信息
// @Description 删除会话信息
// @Tags 删除会话信息
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param f_id formData string true "好友id或者群聊id"
// @Produce json
// @Success 200
// @Router /DelSession [post]
func (*SessionController) DelSession(c *gin.Context) {
	user_id := c.PostForm("f_id")
	if len(user_id) < 1 {
		response.FailResponse(401, "f_id不能为空")
		return
	}
	err := model.DB.Table("im_sessions").Where("m_id=? and f_id=?",
		user2.AuthUser.ID, user_id).Delete(&session.ImSessions{}).Error
	if err != nil {
		response.FailResponse(500, "删除失败").ToJson(c)
		return
	}
	response.SuccessResponse().ToJson(c)
	return
}
