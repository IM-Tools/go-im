package router

import (
	"github.com/gin-gonic/gin"
	"go_im/bin/http/middleware"
	"go_im/bin/service"
)

func RegisterIMRouters(router *gin.Engine) {
	IMservice := new(service.IMservice)

	ws := router.Group("/serve").Use(middleware.WsAuth())
	{
		ws.GET("/ws-con", IMservice.WsConn)
	}
}
