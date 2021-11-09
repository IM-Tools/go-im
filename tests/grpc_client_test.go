/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package tests

import (
	"context"
	"fmt"
	"go_im/config"
	grpc2 "go_im/im/grpc"
	"google.golang.org/grpc"
	"log"
	"testing"
)
//
func init()  {
	config.Initialize()
}

func TestGrpcClient(t *testing.T) {

	//conn,err := grpc.Dial(":"+ conf.GetString("app.grpc_port") ,grpc.WithInsecure())
	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ImRpcServiceClient := grpc2.NewImRpcServiceClient(conn)

	resp ,err := ImRpcServiceClient.SendMessage(context.Background(),&grpc2.MessageRequest{Name:2})

	if err != nil {
		t.Error("调用gRPC方法错误:",err)
		return
	}
	fmt.Println("调用gRPC方法成功:")
	fmt.Println("调用gRPC方法成功，ProdStock=",resp)
}

