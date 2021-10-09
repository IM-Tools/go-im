/**
  @author:panliang
  @data:2021/9/8
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"go_im/im/ws"
	conf "go_im/pkg/config"
	"go_im/pkg/pool"
	"go_im/pkg/zaplog"
	"go_im/router"
)

func StartHttp()  {
	app := gin.Default()
	//初始化连接池
	SetupPool()
	//启动协程执行ws程序
	pool.AntsPool.Submit(func() {
		ws.ImManager.ImStart()
	})

	//注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)
	//全局异常处理
	app.Use(zaplog.Recover)

	_ = app.Run(":" + conf.GetString("app.port"))
}


