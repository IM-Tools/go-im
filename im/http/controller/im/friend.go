/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_im/im/http/models/friend"
	"go_im/im/http/models/friend_record"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/model"
	"go_im/pkg/response"
)
type FriendController struct{}


func (*FriendController) GetList(c *gin.Context)  {
	list,err := friend.GetFriendList(userModel.AuthUser.ID)
	if err !=nil {
		response.FailResponse(500, "获取好友列表异常").ToJson(c)
		return
	}
	fmt.Println(userModel.AuthUser.ID)
	response.SuccessResponse(list).ToJson(c)
	return
}

func (*FriendController ) GetFriendForRecord(c *gin.Context)  {
	list,err := friend_record.GetFriendRecordList(userModel.AuthUser.ID)
	if err !=nil {
		response.FailResponse(500, "获取好友申请记录异常").ToJson(c)
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}


func (*FriendController) SendFriendRequest(c *gin.Context)  {
	information :=c.PostForm("information")
	f_id :=c.PostForm("f_id")

	err := friend_record.AddRecords(userModel.AuthUser.ID,f_id,information)
	if err != nil {
		response.FailResponse(500, "添加失败").ToJson(c)
		return
	}
	response.SuccessResponse().ToJson(c)
	return

}

func (*FriendController) ByFriendRequest(c *gin.Context)  {
	friend_records := friend_record.ImFriendRecord{}
	id :=c.PostForm("id")
	result := model.DB.Table("im_friend_records").
		Where("id=?",id).
		First(&friend_records)
	if result.Error != nil {
		response.FailResponse(500, "查询数据异常").ToJson(c)
		return
	}
	friend.AddFriends(friend_records.UserId,friend_records.Fid)
	friend.AddFriends(friend_records.Fid,friend_records.UserId)
	//投递一条消息
	response.SuccessResponse().ToJson(c)
	return
}

