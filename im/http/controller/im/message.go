/**
  @author:panliang
  @data:2021/8/9
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
	"go_im/im/http/models/msg"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"sort"
	"strconv"
)


type (
	MessageController struct {}
	ImMessage struct {
		ID        uint64 `json:"id"`
		Msg       string `json:"msg"`
		CreatedAt string  `json:"created_at"`
		FromId uint64 `json:"form_id"`
		ToId uint64 `json:"send_id"`
		Channel string `json:"channel"`
		Status int `json:"status"`
		IsRead     int `json:"is_read"`
		MsgType int `json:"msg_type"`
		ChannelType int  `json:"channel_type"`
		//Users userModel.Users `json:"users,omitempty" gorm:"foreignKey:FromId;references:ID"`
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
	pageSize,_ := strconv.Atoi(c.DefaultQuery("pageSize","40"))
	page,_ := strconv.Atoi(c.DefaultQuery("page","1"))
	channel_type := c.DefaultQuery("channel_type","1")
	user := userModel.AuthUser
	from_id := cast.ToString(user.ID)
	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}

	var Users []userModel.Users
	model.DB.Where("id=?",to_id).First(&Users)


	var MsgList []ImMessage
	//生成频道标识符号 用户查询用户信息
	channel_a, channel_b := helpler.ProduceChannelName(from_id, to_id)

	query := model.DB.
		Table("im_messages").

		Where("(channel = ?  or channel= ?) and channel_type=?    order by created_at desc", channel_a, channel_b,channel_type)

	query.Count(&total)

	list := query.
		Offset((page-1)*pageSize).
		Limit(pageSize).
		Select("id,msg,created_at,from_id,to_id,channel,msg_type").
		Find(&MsgList)

	if list.Error != nil {
		return
	}
	from_ids, _ := cast.ToUint64E(user.ID)
	for key, value := range MsgList {
		MsgList[key].CreatedAt = carbon.Parse(value.CreatedAt).SetLocale("zh-CN").DiffForHumans()
		if value.FromId == from_ids {
			MsgList[key].Status = 0
		} else {
			MsgList[key].Status = 1
		}
	}
	SortByAge(MsgList)
	response.SuccessResponse(gin.H{
		"list":MsgList,
		"user":Users[0],
		"mate":gin.H{
			"pageSize":pageSize,
			"page":page,
			"total":total,
		}}, 200).ToJson(c)
}

func SortByAge(list []ImMessage)  {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
}

func SortGroupByAge(list []msg.ImMessage)  {
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
func (*MessageController)GetGroupMessageList(c *gin.Context)  {
	to_id := c.Query("to_id")
	channel_type := c.DefaultQuery("channel_type","2")
	user := userModel.AuthUser

	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}
	var MsgList []msg.ImMessage
	//生成频道标识符号 用户查询用户信息
	channel_a := helpler.ProduceChannelGroupName(to_id)

	list := model.DB.
		Preload("Users").
		Where("channel =? and channel_type=?    order by created_at desc", channel_a,channel_type).
		Limit(40).
		Select("id,msg,created_at,from_id,to_id,channel,msg_type").
		Find(&MsgList)

	if list.Error != nil {
		return
	}
	from_ids, _ := cast.ToUint64E(user.ID)
	for key, value := range MsgList {

		MsgList[key].CreatedAt = carbon.Parse(value.CreatedAt).SetLocale("zh-CN").DiffForHumans()

		if value.FromId == from_ids {
			MsgList[key].Status = 0
		} else {
			MsgList[key].Status = 1
		}
	}
	SortGroupByAge(MsgList)
	response.SuccessResponse( MsgList, 200).ToJson(c)
}

