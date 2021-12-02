/**
  @author:panliang
  @data:2021/9/8
  @note
**/
package core

import (
	"github.com/gin-gonic/gin"
	"im_app/core/ws"
	conf "im_app/pkg/config"
	"im_app/pkg/zaplog"
	"im_app/router"
)
var app_cluster_model = conf.GetBool("core.app_cluster_model")

func StartHttp() {
	app := gin.Default()
	// 初始化各种池
	SetupPool()
	// 启动ws服务
	go ws.ImManager.Start()
	// 启动rpc服务
	if app_cluster_model == true {
		go ws.StartRpc()
	}

	// 注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)
	// 全局异常处理
	app.Use(zaplog.Recover)

	_ = app.Run(":" + conf.GetString("core.port"))

}
