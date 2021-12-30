/**
  @author:panliang
  @data:2021/8/9
  @note
**/
package im

import (
	"im_app/core/http/models/group_message"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"im_app/core/http/models/group_user"
	userModel "im_app/core/http/models/user"
	"im_app/pkg/model"
	"im_app/pkg/response"
)

type (
	MessageController struct{}
	ImMessage         struct {
		ID          int64                   `json:"id"`
		Msg         string                  `json:"msg"`
		CreatedAt   string                  `json:"created_at"`
		FromId      int64                   `json:"user_id"`
		ToId        int64                   `json:"to_id"`
		Channel     string                  `json:"channel"`
		Status      int                     `json:"status"`
		IsRead      int                     `json:"is_read"`
		MsgType     int                     `json:"msg_type"`
		ChannelType int                     `json:"channel_type"`
		Users       userModel.Users         `json:"users,omitempty" gorm:"foreignKey:FromId;references:ID"`
		Group       group_user.ImGroupUsers `json:"group,omitempty" gorm:"foreignKey:FromId;references:ID"`
	}
)

var total int64

// @BasePath /api

// @Summary 获取用户历史消息
// @Description 获取用户历史消息
// @Tags 获取用户历史消息
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param to_id query string true "用户id"
// @Param pageSize query string false "分页条数"
// @Param page query string false "第几页"
// @Produce json
// @Success 200
// @Router /InformationHistory [get]
func (*MessageController) InformationHistory(c *gin.Context) {
	to_id := c.Query("to_id")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "40"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	user := userModel.AuthUser

	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}

	var Users []userModel.Users
	model.DB.Where("id=?", to_id).First(&Users)

	var MsgList []ImMessage

	query := model.DB.
		Table("im_messages").
		Where("from_id = ? and to_id =? and  channel_type=1    order by created_at desc", user.ID, to_id)

	query.Count(&total)

	list := query.
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Select("id,msg,created_at,from_id,to_id,channel,msg_type").
		Find(&MsgList)

	if list.Error != nil {
		return
	}
	for key, value := range MsgList {
		MsgList[key].CreatedAt = carbon.Parse(value.CreatedAt).SetLocale("zh-CN").DiffForHumans()
		if value.FromId == user.ID {
			MsgList[key].Status = 0
		} else {
			MsgList[key].Status = 1
		}
	}
	SortByAge(MsgList)
	response.SuccessResponse(gin.H{
		"list": MsgList,
		"user": Users[0],
		"mate": gin.H{
			"pageSize": pageSize,
			"page":     page,
			"total":    total,
		}}, 200).ToJson(c)
}

func SortByAge(list []ImMessage) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
}

func SortGroupByAge(list []group_message.ImGroupMessages) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
}

// @BasePath /api

// @Summary 获取群聊历史消息
// @Description 获取群聊历史消息
// @Tags 获取群聊历史消息
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_id query string true "群聊id"
// @Produce json
// @Success 200
// @Router /GetGroupMessageList [get]
func (*MessageController) GetGroupMessageList(c *gin.Context) {

	group_id := c.Query("group_id")

	user := userModel.AuthUser
	if len(group_id) < 0 {
		response.FailResponse(401, "group_id不能为空").ToJson(c)
		return
	}
	var MsgList []group_message.ImGroupMessages

	list := model.DB.
		Preload("Users").
		Where("group_id =? order by created_at desc", group_id).
		Limit(40).
		Find(&MsgList)

	if list.Error != nil {
		response.FailResponse(500, "数据查询错误").ToJson(c)
		return
	}
	for key, value := range MsgList {

		if value.FromId == user.ID {
			MsgList[key].Status = 0
		} else {
			MsgList[key].Status = 1
		}
	}
	SortGroupByAge(MsgList)
	response.SuccessResponse(MsgList, 200).ToJson(c)
}

func (*MessageController) GetList(c *gin.Context) {
	user := userModel.AuthUser
	var MsgList []ImMessage
	model.DB.Table("im_messages").
		Where("from_id=?", user.ID).
		Order("created_at").
		Find(&MsgList)

	response.SuccessResponse(MsgList, 200).ToJson(c)
	return
}
