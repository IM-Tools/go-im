/**
  @author:panliang
  @data:2021/9/8
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	grpc2 "go_im/im/grpc"
	"go_im/im/ws"
	conf "go_im/pkg/config"
	"go_im/pkg/zaplog"
	"go_im/router"
)

func StartHttp()  {
	app := gin.Default()
	// 初始化各种池
	SetupPool()
	// 启动ws服务
	go  ws.ImManager.ImStart()
	// 启动rpc服务
	go grpc2.StartRpc()
	//注册路由
	router.RegisterApiRoutes(app)
	router.RegisterIMRouters(app)
	//全局异常处理
	app.Use(zaplog.Recover)

	_ = app.Run(":" + conf.GetString("app.port"))

}

