/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	messageModel "go_im/bin/http/models/msg"
	userModel "go_im/bin/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"go_im/pkg/response"
)

type UsersController struct {
}

type UsersList struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type ImMsgList struct {
	ID        uint64 `json:"id"`
	Msg       string `json:"msg"`
	CreatedAt string `json:"created_at"`
	FromId    uint64 `json:"from_id"`
	ToId      uint64 `json:"to_id"`
	Channel   string `json:"channel"`
	Status    int    `json:"status"`
}

// 获取用户列表
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

// 获取与单个用户消息列表
func (*UsersController) InformationHistory(c *gin.Context) {
	to_id := c.Query("to_id")
	user := userModel.AuthUser
	from_id := cast.ToString(user.ID)

	if len(to_id) < 0 {
		response.FailResponse(500, "用户id不能为空").ToJson(c)
	}
	var MsgList []ImMsgList
	//生成频道
	channel_a, channel_b := helpler.ProduceChannelName(from_id, to_id)
	list := model.DB.
		Model(messageModel.ImMessage{}).
		Where("channel = ?  or channel= ?", channel_a, channel_b).
		Limit(20).
		Select("id,msg,created_at,from_id,to_id,channel").
		Find(&MsgList)

	if list.Error != nil {
		return
	}
	from_ids, _ := cast.ToUint64E(user.ID)

	for key, value := range MsgList {
		if value.FromId == from_ids {
			MsgList[key].Status = 0
		} else {
			MsgList[key].Status = 1
		}
	}
	response.SuccessResponse(MsgList, 200).ToJson(c)
}
