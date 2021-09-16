/**
  @author:panliang
  @data:2021/9/15
  @note
**/
package mq

import (
	"github.com/streadway/amqp"
	"go_im/pkg/config"
	"log"
)

var RabbitMq *amqp.Connection
var err error

//加载mq
func ConnectMQ() *amqp.Connection  {
	RabbitMq,err = amqp.Dial("amqp://"+config.GetString("rabbitmq.user")+":"+
		config.GetString("rabbitmq.password")+ "@"+
		config.GetString("rabbitmq.host")+":"+
		config.GetString("rabbitmq.port")+"/")
	if err != nil {
		log.Fatal("rabbitmq连接失败")
	}
	//defer RabbitMq.Close()

	return RabbitMq
}