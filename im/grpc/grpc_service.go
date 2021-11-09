/**
  @author:panliang
  @data:2021/11/9
  @note
**/
package grpc

import "context"

type ImRpcServer struct {
	
}

func (ps *ImRpcServer)SendMessage(ctx context.Context,request *MessageRequest)(*MessageResponse,error)  {
	return &MessageResponse{Greeting: request.Name},nil
}