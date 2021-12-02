/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package tests

import (
	"context"
	"google.golang.org/grpc"
	"im_app/config"
	grpc2 "im_app/core/ws/rpc"
	"im_app/pkg/zaplog"
	"log"
	"testing"
)
//
func init()  {
	config.Initialize()
}

func TestGrpcClient(t *testing.T) {
	//conn,err := grpc.Dial(":"+ conf.GetString("core.grpc_port") ,grpc.WithInsecure())
	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ImRpcServiceClient := grpc2.NewImRpcServiceClient(conn)

	resp ,err := ImRpcServiceClient.
		SendMessage(context.Background(),
			&grpc2.MessageRequest{Code: 200,FromId: 31,Msg: "Grpc",ToId: 30,Status: 1,MsgType: 1,ChannelType: 1})

	if err != nil {
		t.Error("调用gRPC方法错误:",err)
		return
	}
	zaplog.Info("服务调用成功---",resp)
	return
}


