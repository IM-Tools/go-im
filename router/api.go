/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	Auth "go_im/bin/http/controller/auth"
	"go_im/bin/http/middleware"
)

func RegisterApiRoutes(router *gin.Engine) {
	//允许跨域
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                                                 //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}                                                                      //允许请求的方法
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"} //允许的Header
	router.Use(cors.New(config))

	weibo := new(Auth.WeiBoController)
	auth := new(Auth.AuthController)
	users := new(Auth.UsersController)

	apiRouter := router.Group("/api")
	apiRouter.Group("")
	{
		apiRouter.GET("/WeiBoCallBack", weibo.WeiBoCallBack)
		apiRouter.GET("/giteeCallBack", auth.GiteeCallBack)
		apiRouter.POST("/login", auth.Login)

		apiRouter.Use(middleware.Auth())
		{
			apiRouter.POST("/me", auth.Me)
			apiRouter.POST("/refresh", auth.Refresh)
			apiRouter.GET("/UsersList", users.GetUsersList)
			apiRouter.GET("/InformationHistory", users.InformationHistory)
		}
	}
}
