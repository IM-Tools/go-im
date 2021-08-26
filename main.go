/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package main

import (
	"github.com/gin-gonic/gin"
	"go_im/config"
	"go_im/im"
	"go_im/im/service"
	conf "go_im/pkg/config"
	log2 "go_im/pkg/log"
	"go_im/pkg/pool"
	"go_im/pkg/wordsfilter"
	"go_im/router"
)

func init() {
	config.Initialize()
	//加载敏感词库
	wordsfilter.SetTexts()
}

func main() {
	app := gin.Default()

	//初始化连接池
	im.SetupPool()

	//启动协程执行ws程序
	pool.AntsPool.Submit(func() {
		service.ImManager.ImStart()
	})

	//注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)
	//全局异常处理
	app.Use(log2.Recover)
	_ = app.Run(":" + conf.GetString("app.port"))
}
