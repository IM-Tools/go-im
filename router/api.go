/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	Auth "im_app/core/http/controller/auth"
	"im_app/core/http/controller/im"
	"im_app/core/http/middleware"
	"im_app/docs"
)

func RegisterApiRoutes(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"}
	router.Use(cors.New(config))

	weibo := new(Auth.WeiBoController)
	auth := new(Auth.AuthController)
	users := new(Auth.UsersController)
	sm := new(im.SmApiController)
	uploads := new(im.UploadController)
	group := new(im.GroupController)
	message := new(im.MessageController)
	friends := new(im.FriendController)
	maps := new(im.MapController)
	session := new(im.SessionController)

	docs.SwaggerInfo.BasePath = "/api"

	apiRouter := router.Group("/api")
	apiRouter.Group("")
	{
		apiRouter.POST("/login", auth.Login) // account  login
		apiRouter.GET("/seedRegisteredEmail", auth.SeedRegisteredEmail)
		apiRouter.POST("/registered", auth.Registered) //registered user account

		apiRouter.GET("/WeiBoCallBack", weibo.WeiBoCallBack) // weibo auth
		apiRouter.GET("/getApiToken", sm.GetApiToken)        // get sm token
		apiRouter.GET("/WxCallback", auth.WxCallback)        // get user list
		apiRouter.Use(middleware.Auth(), middleware.GinLogger(), middleware.GinRecovery(true))
		{
			apiRouter.POST("/me", auth.Me)                     // get user info
			apiRouter.POST("/UpdatePwd", auth.UpdatePwd)       // update Pwd
			apiRouter.PUT("/user", auth.Update)                // get user info
			apiRouter.GET("/UsersList", users.GetUsersList)    // get user list
			apiRouter.POST("/bindingEmail", auth.BindingEmail) //binding email

			apiRouter.GET("/InformationHistory", message.InformationHistory)   //get message list
			apiRouter.GET("/GetGroupMessageList", message.GetGroupMessageList) //get message list
			apiRouter.GET("/GetMessageList", message.GetList)                  //get message list

			apiRouter.POST("/UploadImg", sm.UploadImg)                  //upload img
			apiRouter.POST("/UploadVoiceFile", uploads.UploadVoiceFile) //upload voice file

			apiRouter.GET("/ReadMessage", users.ReadMessage)  //read message
			apiRouter.GET("/GetLongitude", maps.GetLongitude) //read message

			apiRouter.GET("/GetGroupList", group.List)                          //get group list
			apiRouter.GET("/GetGroupDetails", group.Show)                       //get group list
			apiRouter.POST("/CreateGroup", group.Create)                        //add group
			apiRouter.POST("/RemoveGroup", group.RemoveGroup)                   //add group
			apiRouter.POST("/RemovedUserFromGroup", group.RemovedUserFromGroup) //remove user  group
			apiRouter.POST("/JoinGroup", group.JoinGroup)                       //remove user  group

			apiRouter.GET("/FriendList", friends.GetList) // get user list
			apiRouter.GET("/GetFriendForRecord", friends.GetFriendForRecord)
			apiRouter.POST("/SendFriendRequest", friends.SendFriendRequest)
			apiRouter.POST("/ByFriendRequest", friends.ByFriendRequest)
			apiRouter.POST("/RemoveFriend", friends.RemoveFriend)
			apiRouter.POST("/FriendPlacedTop", friends.FriendPlacedTop)
			apiRouter.POST("/UpdateFriendNote", friends.UpdateFriendNote)

			apiRouter.POST("/AddSession", session.Create)
			apiRouter.GET("/GetSessionList", session.GetSessionList)
			apiRouter.POST("/DelSession", session.DelSession)
			apiRouter.POST("/SetSessionTop", session.SetSessionTop)

		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/GetLongitude", maps.GetLongitude) //read message
}
