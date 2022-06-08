/**
  @author:panliang
  @data:2022/5/26
  @note
**/
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startCors(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{
		"tus-resumable",
		"upload-length",
		"upload-metadata",
		"cache-control",
		"x-requested-with",
		"*",
	}
	router.Use(cors.New(config))

}
