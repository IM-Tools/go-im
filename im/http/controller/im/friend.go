/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_im/im/http/models/friend"
	"go_im/im/http/models/friend_record"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/model"
	"go_im/pkg/response"
	log2 "go_im/pkg/zaplog"
	"net/http"
	"reflect"
)
type FriendController struct{}

// @Summary 获取好友列表
// @Description 获取好友列表
// @Tags 获取好友列表
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Produce json
// @Success 200
// @Router /FriendList [get]
func (*FriendController) GetList(c *gin.Context)  {
	user := userModel.AuthUser
	var friendId []friend.ImFriends
	err:= model.DB.Select("f_id").Where("m_id=?",user.ID).Find(&friendId).Error
	if err !=nil{
		fmt.Println(err)
	}

	v := reflect.ValueOf(friendId)
	group_slice := make([]uint64, v.Len())
	for key,value := range friendId {
		group_slice[key] = value.FId
	}
	list,err :=userModel.GetFriendListV2(group_slice)
	if err != nil {
		log2.ZapLogger.Error("获取好友列表异常",zap.Error(err))
		response.FailResponse(http.StatusInternalServerError,"服务器错误")
		return
	}
	response.SuccessResponse(map[string]interface{}{
		"list": list,
	}, 200).ToJson(c)
	return
}
// @Summary 获取好友申请记录
// @Description 获取好友申请记录
// @Tags 获取好友申请记录
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Produce json
// @Success 200
// @Router /GetFriendForRecord [get]
func (*FriendController ) GetFriendForRecord(c *gin.Context)  {
	list,err := friend_record.GetFriendRecordList(userModel.AuthUser.ID)
	if err !=nil {
		response.FailResponse(500, "获取好友申请记录异常").ToJson(c)
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}

// @BasePath /api

// @Summary 发送好友请求
// @Description 发送好友请求接口
// @Tags 发送好友请求接口
// @Accept multipart/form-data
// @Produce json
// @Param information formData string true "请求描述"
// @Param f_id formData string true "用户id"
// @Param client_type formData string false "客户端类型 0.网页端登录 1.设备端登录"
// @Success 200
// @Router /SendFriendRequest [post]
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
// @BasePath /api

// @Summary 同意好友请求
// @Description 同意好友请求接口
// @Tags 同意好友请求接口
// @Accept multipart/form-data
// @Produce json
// @Param id formData string true "请求记录id"
// @Success 200
// @Router /ByFriendRequest [post]
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

