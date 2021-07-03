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
	"go_im/pkg/ws"
	"go_im/router"
)

func init() {
	config.Initialize()
}

func main() {
	app := gin.Default()
	//加载连接池
	bin.SetupDB()
	//启动协程执行开始程序
	go ws.Manager.Start()

	//注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)

	_ = app.Run(":" + conf.GetString("app.port"))
}
