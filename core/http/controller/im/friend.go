/**
  @author:panliang
  @data:2021/9/4
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im_app/core/http/models/friend"
	"im_app/core/http/models/friend_record"
	userModel "im_app/core/http/models/user"
	"im_app/pkg/model"
	"im_app/pkg/response"
	"im_app/pkg/zaplog"
	"net/http"
	"reflect"
	"strconv"
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
func (*FriendController) GetList(c *gin.Context) {
	user := userModel.AuthUser
	var friendId []friend.ImFriends
	err := model.DB.Select("f_id").Where("m_id=?", user.ID).Find(&friendId).Error
	if err != nil {
		fmt.Println(err)
	}

	v := reflect.ValueOf(friendId)
	group_slice := make([]int64, v.Len())
	for key, value := range friendId {
		group_slice[key] = value.FId
	}
	list, err := userModel.GetFriendListV2(group_slice)
	if err != nil {
		zaplog.Error("----获取好友列表异常", err)
		response.FailResponse(http.StatusInternalServerError, "服务器错误")
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
func (*FriendController) GetFriendForRecord(c *gin.Context) {

	list, err := friend_record.GetFriendRecordList(userModel.AuthUser.ID)
	if err != nil {
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
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param information formData string true "请求描述"
// @Param f_id formData string true "用户id"
// @Param client_type formData string false "客户端类型 0.网页端登录 1.设备端登录 2.app端"
// @Success 200
// @Router /SendFriendRequest [post]
func (*FriendController) SendFriendRequest(c *gin.Context) {

	information := c.PostForm("information")
	f_id := c.PostForm("f_id")
	fId, _ := strconv.Atoi(f_id)

	if int64(fId) == userModel.AuthUser.ID {
		response.FailResponse(401, "请勿添加自己为好友").ToJson(c)
		return
	}
	var friend friend.ImFriends

	model.DB.Table("im_friends").
		Where("status=1 and f_id=? and m_id=?", f_id, userModel.AuthUser.ID).
		Find(&friend)

	if friend.ID == 0 {
		err := friend_record.AddRecords(userModel.AuthUser.ID, f_id, information)
		if err != nil {
			response.FailResponse(500, "添加失败").ToJson(c)
			return
		}
		response.SuccessResponse().ToJson(c)
		return
	} else {
		response.FailResponse(401, "已经是好友关系了，请勿重复添加")
		return
	}

}

type ImFriendRecords struct {
	ID     int64 `json:"id"`
	UserId int64 `json:"user_id"`
	FId    int64 `json:"f_id"`
	Status int   `json:"status"`
}

// @BasePath /api

// @Summary 同意好友请求接口
// @Description 同意好友请求接口
// @Tags 同意好友请求接口
// @Accept multipart/form-data
// @Produce json
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param information formData string true "请求描述"
// @Param id formData string true "请求记录id"
// @Param status formData bool true  "1.同意 0 拒绝"
// @Success 200
// @Router /ByFriendRequest [post]
func (*FriendController) ByFriendRequest(c *gin.Context) {

	id := c.PostForm("id")
	sta := c.PostForm("status")
	status, _ := strconv.Atoi(sta)

	var friends ImFriendRecords
	err := model.DB.Where("id=?", id).
		First(&friends).Error
	if err != nil {
		response.FailResponse(500, "添加失败").ToJson(c)
		return
	}

	if status == 0 {
		friends.Status = 2
		model.DB.Save(&friends)
		// 投递一条消息
		response.FailResponse(500, "已经拒绝了~").ToJson(c)
	} else {
		friend.AddFriends(friends.UserId, friends.FId)
		friend.AddFriends(friends.FId, friends.UserId)
		friends.Status = 1
		model.DB.Save(&friends)

		// 投递一条消息
		response.SuccessResponse().ToJson(c)
		return

	}

}