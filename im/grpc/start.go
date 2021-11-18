/**
  @author:panliang
  @data:2021/11/10
  @note
**/
package grpc

import (
	conf "im_app/pkg/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

//启动rpc服务
func StartRpc() {
	rpcServer := grpc.NewServer()

	RegisterImRpcServiceServer(rpcServer, new(ImRpcServer))

	listener, err := net.Listen("tcp", ":"+conf.GetString("app.grpc_port"))
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = rpcServer.Serve(listener)

}
