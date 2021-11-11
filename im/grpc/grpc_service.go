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
	"go_im/im/ws"
	"strconv"
)

type ImRpcServer struct {
	
}
// rpc消息投递
func (ps *ImRpcServer)SendMessage(ctx context.Context,request *MessageRequest)(*MessageResponse,error)  {
	jsonMessage_from, _ := json.Marshal(&ws.Msg{Code: int(request.Code), Msg:request.Msg ,
		FromId: int(request.FromId),
		ToId:   int(request.ToId), Status:1, MsgType: int(request.MsgType),ChannelType:int( request.ChannelType)})
	var  manager ws.ImClientManager
	fmt.Println(jsonMessage_from)
	to_id := strconv.Itoa(int(request.ToId))

	if data,ok :=manager.ImClientMap[to_id];ok {
		data.Send <- jsonMessage_from
	} else {
		ws.MqPersonalPublish(jsonMessage_from,int(request.ToId))
	}
	return &MessageResponse{Code: 200},nil
}
