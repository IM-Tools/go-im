package router

import (
	"github.com/gin-gonic/gin"
	"go_im/bin/http/middleware"
	"go_im/bin/service"
)

func RegisterIMRouters(router *gin.Engine) {
	IMservice := new(service.IMservice)

	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", IMservice.Connect)
	}
}
