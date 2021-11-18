/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	"im_app/im/ws"
	conf "im_app/pkg/config"
)

var RpcServer = grpc.NewServer()

type ImRpcServerHandler interface {
	StartRpc()
}

type ImRpcServer struct {
}

// 启动rpc服务
func StartRpc() {

	RegisterImRpcServiceServer(RpcServer, new(ImRpcServer))

	listener, err := net.Listen("tcp", ":"+conf.GetString("app.grpc_port"))
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	_ = RpcServer.Serve(listener)

}

// rpc消息投递
func (ps *ImRpcServer) SendMessage(ctx context.Context, request *MessageRequest) (*MessageResponse, error) {
	jsonMessage_from, _ := json.Marshal(&ws.Msg{Code: int(request.Code), Msg: request.Msg,
		FromId: int(request.FromId),
		ToId:   int(request.ToId), Status: 1, MsgType: int(request.MsgType), ChannelType: int(request.ChannelType)})
	var manager ws.ImClientManager
	fmt.Println(jsonMessage_from)
	to_id := strconv.Itoa(int(request.ToId))

	if data, ok := manager.ImClientMap[to_id]; ok {
		data.Send <- jsonMessage_from
	} else {
		ws.MqPersonalPublish(jsonMessage_from, int(request.ToId))
	}
	return &MessageResponse{Code: 200}, nil
}
