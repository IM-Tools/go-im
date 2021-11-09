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
	"go_im/pkg/pool"
	"go_im/pkg/zaplog"
	"go_im/router"
	"google.golang.org/grpc"
	"log"
	"net"
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
	startRpc()
	_ = app.Run(":" + conf.GetString("app.port"))

}

func startRpc()  {
	rpcServer := grpc.NewServer()

	grpc2.RegisterImRpcServiceServer(rpcServer, new(grpc2.ImRpcServer));

	listener, err := net.Listen("tcp", ":"+conf.GetString("app.grpc_port"))
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = rpcServer.Serve(listener)

}


