package router

import (
	"github.com/gin-gonic/gin"
	"go_im/im/http/middleware"
	"go_im/im/service"
)
func RegisterIMRouters(router *gin.Engine) {
	IMService := new(service.IMService)
	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", IMService.Connect)
	}
}
