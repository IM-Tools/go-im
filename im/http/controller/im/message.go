/**
  @author:panliang
  @data:2021/8/9
  @note
**/
package im

import (
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"im_app/im/http/models/group_user"
	"im_app/im/http/models/msg"
	userModel "im_app/im/http/models/user"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/response"
)

type (
	MessageController struct{}
	ImMessage         struct {
		ID          int64                  `json:"id"`
		Msg         string                  `json:"msg"`
		CreatedAt   string                  `json:"created_at"`
		FromId      int64                  `json:"user_id"`
		ToId        int64                  `json:"send_id"`
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
	channel_type := c.DefaultQuery("channel_type", "1")
	user := userModel.AuthUser
	from_id := user.ID
	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}

	var Users []userModel.Users
	model.DB.Where("id=?", to_id).First(&Users)

	var MsgList []ImMessage
	// 生成频道标识符号 用户查询用户信息
	toid,_ :=strconv.Atoi(to_id)
	channel_a, channel_b := helpler.ProduceChannelName(int64(from_id), int64(toid))

	query := model.DB.
		Table("im_messages").
		Where("(channel = ?  or channel= ?) and channel_type=?    order by created_at desc", channel_a, channel_b, channel_type)

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

func SortGroupByAge(list []msg.ImMessage) {
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
// @Param to_id query string true "群聊id"
// @Produce json
// @Success 200
// @Router /GetGroupMessageList [get]
func (*MessageController) GetGroupMessageList(c *gin.Context) {
	to_id := c.Query("to_id")
	channel_type := c.DefaultQuery("channel_type", "2")
	user := userModel.AuthUser

	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}
	var MsgList []msg.ImMessage
	// 生成频道标识符号 用户查询用户信息
	channel_a := helpler.ProduceChannelGroupName(to_id)

	list := model.DB.
		Preload("Users").
		Where("channel =? and channel_type=?    order by created_at desc", channel_a, channel_type).
		Limit(40).
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
