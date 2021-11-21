/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package ws

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"im_app/im/ws/rpc"
	conf "im_app/pkg/config"
	"im_app/pkg/zaplog"
	"log"
	"net"
)

var RpcServer = grpc.NewServer()

type ImRpcServerHandler interface {
	StartRpc()
}

type ImRpcServer struct {
}

// 启动rpc服务
func StartRpc() {

	rpc.RegisterImRpcServiceServer(RpcServer, new(ImRpcServer))

	listener, err := net.Listen("tcp", ":"+conf.GetString("app.grpc_port"))
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = RpcServer.Serve(listener)

}

// rpc消息投递
func (ps *ImRpcServer) SendMessage(ctx context.Context, request *rpc.MessageRequest) (*rpc.MessageResponse, error) {
	jsonMessage_from, _ := json.Marshal(&RpcMsg{Code: int(request.Code), Msg: request.Msg,
		FromId: int(request.FromId),
		ToId:   int(request.ToId), Status: 1, MsgType: int(request.MsgType), ChannelType: int(request.ChannelType)})

	zaplog.Info(jsonMessage_from)
	if data, ok := ImManager.ImClientMap[int64(request.ToId)]; ok {
		data.Send <- jsonMessage_from
	} else {
		MqPersonalPublish(jsonMessage_from, int(request.ToId))
	}
	return &rpc.MessageResponse{Code: 200}, nil
}
