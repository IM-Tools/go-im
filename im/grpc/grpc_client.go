/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package grpc

import (
	"context"
	"fmt"
	conf "go_im/pkg/config"
	"google.golang.org/grpc"
	"log"
)

// 投递消息
func SendRpcMsg() {
	conn, err := grpc.Dial(":"+conf.GetString("app.grpc_port"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ImRpcServiceClient := NewImRpcServiceClient(conn)
	data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
		"ces", 2, 3, 1, 1)
	resp, err := ImRpcServiceClient.SendMessage(context.Background(), &MessageRequest{Msg: data})

	if err != nil {
		return
	}
	fmt.Println("调用gRPC方法成功，ProdStock=", resp)
	return
}
