/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package im

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"im_app/core/http/models/group"
	"im_app/core/http/models/group_user"
	userModel "im_app/core/http/models/user"
	"im_app/core/http/validates"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/response"
	"im_app/pkg/zaplog"
)

type (
	GroupController struct{}
	Groups          struct {
		GroupId string `json:"group_id"`
	}
)

// @BasePath /api

// @Summary 获取群聊列表
// @Description 获取群聊列表
// @Tags 获取群聊列表
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Produce json
// @Success 200
// @Router /GetGroupList [get]
func (*GroupController) List(c *gin.Context) {

	user := userModel.AuthUser
	var groupId []Groups
	err := model.DB.Table("im_group_users").
		Where("user_id=?", user.ID).
		Group("group_id").
		Find(&groupId).Error
	if err != nil {
		fmt.Println(err)
	}
	v := reflect.ValueOf(groupId)
	group_slice := make([]string, v.Len())
	for key, value := range groupId {
		group_slice[key] = value.GroupId
	}
	fmt.Println(group_slice)
	list, err := group.GetGroupUserList(group_slice)

	if err != nil {
		zaplog.Error("----获取群聊列表异常", err)
		response.FailResponse(http.StatusInternalServerError, "服务器错误")
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}

// @BasePath /api

// @Summary 创建群聊
// @Description 创建群聊
// @Tags 创建群聊
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_name formData string true "群聊名称"
// @Param user_id formData array true "群聊用户"
// @Produce json
// @Success 200
// @Router /CreateGroup [post]
func (*GroupController) Create(c *gin.Context) {
	user := userModel.AuthUser

	_groups := validates.CreateGroupParams{
		GroupName: c.PostForm("group_name"),
		UserId:    c.PostFormMap("user_id"),
	}
	fmt.Println(_groups)
	rules := govalidator.MapData{
		"group_name": []string{"required", "between:2,20"},
		// "user_id": []string{"required"},
	}
	opts := govalidator.Options{
		Data:          &_groups,
		Rules:         rules,
		TagIdentifier: "valid",
	}
	errs := govalidator.New(opts).ValidateStruct()

	if len(errs) > 0 {

		data, _ := json.MarshalIndent(errs, "", "  ")
		var result = helpler.JsonToMap(data)
		response.ErrorResponse(http.StatusInternalServerError, "参数不合格", result).ToJson(c)
		return
	}
	if len(_groups.UserId) > 50 {
		response.ErrorResponse(http.StatusInternalServerError, "默认只能邀请50人入群").ToJson(c)
	}

	id, err := group.Created(user.ID, _groups.GroupName)
	if err != nil {
		response.ErrorResponse(http.StatusInternalServerError, "创建异常").ToJson(c)
		return
	}
	err = group_user.CreatedAll(_groups.UserId, id, user.ID)
	if err != nil {
		zaplog.Error("----创建群聊异常", err)
		response.ErrorResponse(http.StatusInternalServerError, "创建异常").ToJson(c)
		return
	}
	response.SuccessResponse().ToJson(c)
	return
}

// @BasePath /api

// @Summary 删除群聊
// @Description 删除群聊
// @Tags 删除群聊
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_id formData string true "群聊id"
// @Produce json
// @Success 200
// @Router /RemoveGroup [post]
func (*GroupController) RemoveGroup(c *gin.Context) {
	group_id := c.PostForm("group_id")
	if len(group_id) == 0 {
		response.ErrorResponse(http.StatusInternalServerError, "参数不合格").ToJson(c)
		return
	}
	model.DB.Where("id=?", group_id).Delete(&group.ImGroups{})
	model.DB.Where("group_id=?", group_id).Delete(&group_user.ImGroupUsers{})
	response.SuccessResponse().ToJson(c)
	return
}

// @BasePath /api

// @Summary 移除群聊用户
// @Description 移除群聊用户
// @Tags 移除群聊用户
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_id formData string true "群聊id"
// @Param user_id formData string true "用户id"
// @Produce json
// @Success 200
// @Router /RemovedUserFromGroup [post]
func (*GroupController) RemovedUserFromGroup(c *gin.Context) {

	_group := validates.RemoveUserFormGroupFrom{
		GroupId: c.PostForm("group_id"),
		UserId:  c.PostForm("user_id"),
	}
	errs := validates.ValidateRemoveGroupForm(_group)

	if len(errs) > 0 {
		response.FailResponse(401, "error", errs)
	}
	g_id, _ := group.GetGroupUserId(_group.GroupId)

	if userModel.AuthUser.ID != g_id {
		response.FailResponse(401, "没有权限删除群成员！").ToJson(c)
		return
	}
	model.DB.Table("im_group_users").
		Where("user_id", _group.UserId).
		Delete(&group_user.ImGroupUsers{})

	response.SuccessResponse().ToJson(c)
	return
}
