/**
  @author:panliang
  @data:2021/8/9
  @note
**/
package im

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
)

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

type MessageController struct {}

func (*MessageController) InformationHistory(c *gin.Context) {
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