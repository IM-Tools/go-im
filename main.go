/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package main

import (
	"github.com/gin-gonic/gin"
	"go_im/bin"
	"go_im/config"
	conf "go_im/pkg/config"
	"go_im/router"
)

func init()  {
	config.Initialize()
}



func main()  {



	app := gin.Default()
	//加载连接池
	bin.SetupDB()
	//注册路由
	router.RegisterApiRoutes(app)

	app.Run(":"+conf.GetString("app.port"))
}
