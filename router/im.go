package router

import (
	"github.com/gin-gonic/gin"
	"im_app/core/http/controller/im"
	"im_app/core/http/middleware"
)

func RegisterIMRouters(router *gin.Engine) {
	IMService := new(im.IMService)

	// im聊天
	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", IMService.Connect)

	}

	// 客服
	router.GET("/im/customer_service", IMService.CustomerService)
}
