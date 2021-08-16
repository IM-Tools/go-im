/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
	messageModel "go_im/im/http/models/msg"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"sort"
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
type ImMsgList struct {
	ID        uint64 `json:"id"`
	Msg       string `json:"msg"`
	CreatedAt string `json:"created_at"`
	FromId    uint64 `json:"from_id"`
	ToId      uint64 `json:"to_id"`
	Channel   string `json:"channel"`
	Status    int    `json:"status"`
	MsgType   int `json:"msg_type"`

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
	//生成频道标识符号 用户查询用户信息
	channel_a, channel_b := helpler.ProduceChannelName(from_id, to_id)
	fmt.Println(channel_b,channel_a)
	list := model.DB.
		Model(messageModel.ImMessage{}).
		Where("channel = ?  or channel= ?  order by created_at desc", channel_a, channel_b).
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
	SortByAge(MsgList)
	response.SuccessResponse( MsgList, 200).ToJson(c)
}

func SortByAge(list []ImMsgList)  {
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
}

func (*UsersController) ReadMessage(c *gin.Context) {
	user := userModel.AuthUser
	channel_a, channel_b := helpler.ProduceChannelName(strconv.Itoa(int(user.ID)), c.Query("to_id"))
	messageModel.ReadMsg(channel_a,channel_b)
	response.SuccessResponse(gin.H{}, 200).ToJson(c)
}

