/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	Auth "go_im/bin/http/controller/auth"
	"go_im/bin/http/middleware"
	"go_im/bin/service"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func RegisterApiRoutes(router *gin.Engine) {
	//允许跨域
	weibo := new(Auth.WeiBoController)
	auth := new(Auth.AuthController)
	users := new(Auth.UsersController)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}//允许请求的方法
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"}//允许的Header
	router.Use(cors.New(config))
	//router.Use(middleware.CrosHandler())
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world!!!",
		})
	})
	router.GET("/api/WeiBoCallBack", weibo.WeiBoCallBack)
	router.GET("/api/giteeCallBack", auth.GiteeCallBack)
	api := router.Group("/api").Use(middleware.Auth())
	{
		api.POST("/me", auth.Me)
		api.POST("/refresh", auth.Refresh)
		api.GET("/UsersList", users.GetUsersList)
		//将该连接升级为ws
	}
	wsServe := new(service.WsServe)
	ws := router.Group("/serve").Use(middleware.WsAuth())
	{
		ws.GET("/ws-con", wsServe.WsConn)
		//将该连接升级为ws
	}
}
