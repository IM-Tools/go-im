/**
  @author:panliang
  @data:2021/12/7
  @note
**/
package services

type SendMessageHandler interface {
	MessageDelivery() //投递一条消息到推送服务
}

type MessageService struct {
}

func (*MessageService) MessageDelivery() {
	//组装消息
}
