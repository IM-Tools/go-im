/**
  @author:panliang
  @data:2021/9/3
  @note
**/
package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartTcpServe()  {
	listener,err := net.Listen("tcp",":8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for  {
		conn,err :=listener.Accept()
		if err !=nil {
			log.Fatal(err)
			continue
		}
		go HandleConn(conn)

	}
}

func HandleConn(conn net.Conn)  {
	ch := make(chan string)
	go clientWriter(conn,ch)
	who := conn.RemoteAddr().String()
	ch <-"测试"+who //单发给自己
	messages<-who+"自己" //广播自己
	entering<-ch //将自己加入广播频道
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages<-who+":"+input.Text()
	}
	leaving<-ch
	messages<-who+"离开了"
	conn.Close()
}

func clientWriter(conn net.Conn,ch <-chan string)  {
	for msg:=range ch {
		fmt.Fprintf(conn,"%s\n",msg)
	}
}

type client chan<-string //消息通道
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接受的客户消息
)

func broadcaster()  {
	clients :=make(map[client]bool)
	for  {
		select {
		case msg:= <-messages:
			for cli :=range clients {
				cli<-msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
