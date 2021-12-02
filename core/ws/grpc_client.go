/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"im_app/core/ws/rpc"
	"im_app/pkg/zaplog"
	"log"
)

//
//type ImRPCHandler interface {
//	SendRpcMsg(message []byte, node string)
//}
//
//type ImRpcClient struct {
//}

// 使用rpc开始投递消息

// 发送的消息
type RpcMsg struct {
	Code        int    `json:"code,omitempty"`
	FromId      int    `json:"from_id,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ToId        int    `json:"to_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	MsgType     int    `json:"msg_type,omitempty"`
	ChannelType int    `json:"channel_type"`
}


func  SendRpcMsg(message []byte, node string) {

	var msg RpcMsg
	err := json.Unmarshal(message, &msg)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}

	conn, err := grpc.Dial(node, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ImRpcServiceClient := rpc.NewImRpcServiceClient(conn)

	resp, err := ImRpcServiceClient.
		SendMessage(context.Background(),
			&rpc.MessageRequest{Code: 200,
				FromId:      int32(msg.FromId),
				Msg:         msg.Msg,
				ToId:        int32(msg.ToId),
				Status:      1,
				MsgType:     int32(msg.MsgType),
				ChannelType: int32(msg.ChannelType)})
	if err != nil {
		zaplog.Info("异常",err)
		return
	}
	zaplog.Info("调用gRPC方法成功",resp)
	return
}
