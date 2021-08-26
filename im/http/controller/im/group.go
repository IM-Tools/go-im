/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package im

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go_im/im/http/models/group"
	"go_im/im/http/models/group_user"
	userModel "go_im/im/http/models/user"
	"go_im/im/http/validates"
	"go_im/pkg/helpler"
	log2 "go_im/pkg/log"
	"go_im/pkg/model"
	"go_im/pkg/response"
	"net/http"
	"reflect"
)

type GroupController struct {}
type Groups struct {
	GroupId string `json:"group_id"`
}

func (*GroupController) List(c *gin.Context){

	user :=userModel.AuthUser
	var groupId []Groups
	err :=model.DB.Table("im_group_users").
		Where("user_id=?",user.ID).
		Group("group_id").
		Find(&groupId).Error; if err !=nil{
		fmt.Println(err)
	}
	v := reflect.ValueOf(groupId)
	group_slice := make([]string, v.Len())
	for key,value := range groupId {
		group_slice[key] = value.GroupId
	}
	list,err :=group.GetGroupUserList(group_slice)
	if err != nil {
		log2.Warning(err.Error())
		response.FailResponse(http.StatusInternalServerError,"服务器错误")
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}
//创建一个新群聊
func (*GroupController) Create(c *gin.Context){
	user :=userModel.AuthUser

	_groups := validates.CreateGroupParams{
		GroupName: c.PostForm("group_name"),
		UserId:c.PostFormMap("user_id") ,
	}
	fmt.Println(_groups)
	rules := govalidator.MapData{
		"group_name": []string{"required","between:2,20"},
		//"user_id": []string{"required"},
	}
	opts := govalidator.Options{
		Data:          &_groups,
		Rules:         rules,
		TagIdentifier: "valid",
	}
	errs := govalidator.New(opts).ValidateStruct()

	if len(errs) >0 {
		data, _ := json.MarshalIndent(errs, "", "  ")
		var  result =  helpler.JsonToMap(data)
		response.ErrorResponse(http.StatusInternalServerError,"参数不合格",result).ToJson(c)
		return
	}
	if len(_groups.UserId) > 50 {
		response.ErrorResponse(http.StatusInternalServerError,"默认只能邀请50人入群").ToJson(c)
	}

	id,err :=group.Created(user.ID,_groups.GroupName);if err != nil {
		fmt.Println("异常")
		response.ErrorResponse(http.StatusInternalServerError,"创建异常").ToJson(c)
		return
	}
	err = group_user.CreatedAll(_groups.UserId,id,user.ID)
	if err != nil {
		response.ErrorResponse(http.StatusInternalServerError,"创建异常").ToJson(c)
		return
	}
	response.SuccessResponse().ToJson(c)
	return
}

func (*GroupController) RemoveGroup(c *gin.Context){
	group_id := c.PostForm("group_id")
	if len(group_id) == 0 {
		response.ErrorResponse(http.StatusInternalServerError,"参数不合格").ToJson(c)
		return
	}
	model.DB.Where("id=?",group_id).Delete(&group.ImGroups{});
	model.DB.Where("group_id=?",group_id).Delete(&group_user.ImGroupUsers{});
	response.SuccessResponse().ToJson(c)
	return
}

func (*GroupController) DeleteUser(c *gin.Context){

}


