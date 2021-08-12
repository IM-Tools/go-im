package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_im/bin/http/middleware"
	"go_im/bin/service"
)

func RegisterIMRouters(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                                                 //允许所有域名
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}                                                                      //允许请求的方法
	config.AllowHeaders = []string{"tus-resumable", "upload-length", "upload-metadata", "cache-control", "x-requested-with", "*"} //允许的Header
	router.Use(cors.New(config))

	IMservice := new(service.IMservice)

	ws := router.Group("/im").Use(middleware.Auth())
	{
		ws.GET("/connect", IMservice.Connect)
	}
}
