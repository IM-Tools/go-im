package router

import (
	"github.com/gin-gonic/gin"
	"go_im/im/http/controller/im"
	"go_im/im/http/middleware"
)
func RegisterIMRouters(router *gin.Engine) {
	IMService := new(im.IMService)
	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", IMService.Connect)
	}
}
