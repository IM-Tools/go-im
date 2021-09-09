/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Auth "go_im/im/http/controller/auth"
	"go_im/im/http/controller/im"
	"go_im/im/http/middleware"
)

func RegisterApiRoutes(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"}
	router.Use(cors.New(config))
	weibo := new(Auth.WeiBoController)
	auth  := new(Auth.AuthController)
	users := new(Auth.UsersController)
	sm    := new(im.SmApiController)
	uploads := new(im.UploadController)
	group := new(im.GroupController)
	message := new(im.MessageController)
	friends := new(im.FriendController)

	apiRouter := router.Group("/api")
	apiRouter.Group("")
	{
		apiRouter.POST("/login", auth.Login)                 // account  login
		apiRouter.GET("/WeiBoCallBack", weibo.WeiBoCallBack) // weibo auth
		apiRouter.GET("/getApiToken",  sm.GetApiToken)       // get sm token
		apiRouter.Use(middleware.Auth())
		{

			apiRouter.POST("/me", auth.Me)                  // get user info
			apiRouter.GET("/UsersList", users.GetUsersList) // get user list

			apiRouter.GET("/InformationHistory", message.InformationHistory) //get message list
			apiRouter.GET("/GetGroupMessageList", message.GetGroupMessageList) //get message list
			apiRouter.POST("/UploadImg", sm.UploadImg)                  //upload img
			apiRouter.POST("/UploadVoiceFile", uploads.UploadVoiceFile) //upload voice file
			apiRouter.GET("/ReadMessage", users.ReadMessage)            //read message

			apiRouter.GET("/GetGroupList", group.List)                  //get group list
			apiRouter.POST("/CreateGroup", group.Create)                //add group
			apiRouter.POST("/RemoveGroup", group.RemoveGroup)                //add group

			apiRouter.GET("/GetFriendsList", friends.GetList)
			apiRouter.GET("/GetFriendForRecord", friends.GetFriendForRecord)
			apiRouter.POST("/SendFriendRequest", friends.SendFriendRequest)
			apiRouter.POST("/ByFriendRequest", friends.ByFriendRequest)
		}
	}
}
