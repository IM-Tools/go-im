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
	"go.uber.org/zap"

	"im_app/im/http/models/group"
	"im_app/im/http/models/group_user"
	userModel "im_app/im/http/models/user"
	"im_app/im/http/validates"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/response"
	log2 "im_app/pkg/zaplog"
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
		log2.ZapLogger.Error("获取群聊列表异常", zap.Error(err))
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
		log2.ZapLogger.Error("创建群聊异常", zap.Error(err))
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

func (*GroupController) DeleteUser(c *gin.Context) {

}
