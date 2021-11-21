/**
  @author:panliang
  @data:2021/9/8
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"im_app/im/ws"
	conf "im_app/pkg/config"
	"im_app/pkg/zaplog"
	"im_app/router"
)

func StartHttp() {
	app := gin.Default()
	// 初始化各种池
	SetupPool()
	// 启动ws服务
	go ws.ImManager.Start()
	// 启动rpc服务
	go ws.StartRpc()
	// 注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)
	// 全局异常处理
	app.Use(zaplog.Recover)

	_ = app.Run(":" + conf.GetString("app.port"))

}
