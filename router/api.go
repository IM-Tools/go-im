/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
	Auth "go_im/bin/http/controller/auth"
	"go_im/bin/http/middleware"
	"go_im/bin/service"
)

var router *gin.Engine

func RegisterApiRoutes(router *gin.Engine)  {
	//允许跨域
	router.Use(middleware.CrosHandler())

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world!!!",
		})
	})
	weibo := new(Auth.WeiBoController)
	auth := new(Auth.AuthController)
	users := new(Auth.UsersController)
	router.GET("/api/WeiBoCallBack",weibo.WeiBoCallBack)
	router.GET("/api/giteeCallBack",auth.GiteeCallBack)

	api := router.Group("/api").Use(middleware.Auth())
	{
		api.POST("/me",auth.Me)
		api.POST("/refresh",auth.Refresh)
		api.GET("/usersList",users.GetUsersList)
		//将该连接升级为ws

	}

	wsServe := new(service.WsServe)
	ws := router.Group("/serve").Use(middleware.WsAuth())
	{
		ws.GET("/ws-con",wsServe.WsConn)
		//将该连接升级为ws
	}



}