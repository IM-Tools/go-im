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
	"net/url"

	"google.golang.org/grpc"

	"im_app/im/ws"
)

type ImRPCHandler interface {
	SendRpcMsg(message []byte, node string)
}

type ImRpcClient struct {
}

// 使用rpc开始投递消息
func (prc *ImRpcClient) SendRpcMsg(message []byte, node string) {
	data := ws.EnMessage(message)
	msg := new(ws.Msg)
	err := json.Unmarshal([]byte(data.Content), &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgs, _ := url.QueryUnescape(msg.Msg)
	msg.Msg = msgs

	conn, err := grpc.Dial(node, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ImRpcServiceClient := NewImRpcServiceClient(conn)

	resp, err := ImRpcServiceClient.
		SendMessage(context.Background(),
			&MessageRequest{Code: 200,
				FromId:      int32(msg.FromId),
				Msg:         msg.Msg,
				ToId:        int32(msg.ToId),
				Status:      1,
				MsgType:     int32(msg.MsgType),
				ChannelType: int32(msg.ChannelType)})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("调用gRPC方法成功，ProdStock=", resp)
	return
}
