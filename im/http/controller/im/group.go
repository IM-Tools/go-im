/**
  @author:panliang
  @data:2021/7/13
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"go_im/im/http/models/group"
	userModel "go_im/im/http/models/user"
	log2 "go_im/pkg/log"
	"go_im/pkg/response"
)

type GroupController struct {}

func (*GroupController) List(c *gin.Context){
	user :=userModel.AuthUser
	list,err :=group.GetGroupUserList(user.ID)
	if err != nil {
		log2.Warning(err.Error())
		response.FailResponse(500,"服务器错误")
		return
	}
	response.SuccessResponse(list).ToJson(c)
	return
}
func (*GroupController) Create(){

}
func (*GroupController) RemoveUser(){

}
func (*GroupController) Delete(){

}


