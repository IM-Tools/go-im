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
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world!!!",
		})
	})
	weibo := new(Auth.WeiBo)
	auth := new(Auth.AuthController)
	router.GET("/api/WeiBoCallBack",weibo.WeiBoCallBack)


	api := router.Group("/api").Use(middleware.Auth())
	{
		api.POST("/me",auth.Me)
		//将该连接升级为ws
		ws := new(service.WsServe)
		api.GET("/ws-con",ws.WsConn)
	}
}