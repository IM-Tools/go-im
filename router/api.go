/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package router

import (
	"github.com/gin-gonic/gin"
	Auth "go_im/bin/http/controller/auth"
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
	router.GET("/api/WeiBoCallBack",weibo.WeiBoCallBack)

	//将该连接升级为ws
	ws := new(service.WsServe)
	router.GET("/ws-con",ws.WsConn)



}