/**
  @author:panliang
  @data:2021/9/3
  @note
**/
package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)




func StartTcpClient()  {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		log.Fatal(err)
	}
	//login(conn)
	done := make(chan struct{})
	go func() {
		fmt.Println(os.Stdout)
		io.Copy(os.Stdout, conn) // 注意：忽略错误
		log.Println("done")
		done <- struct{}{} // 通知主 goroutine 的信号
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // 等待后台 goroutine 完成
}

//func login(conn net.Conn)  {
//	username :=getClientLoginName(conn)
//	password :=getClientLoginPwd(conn)
//	fmt.Println(username,password)
//
//
//
//
//}


func mustCopy(dst io.Writer,src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// 获取登录信息

func getClientLoginName(conn net.Conn) (username string)  {
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn,"请输入你需要登录账号：")
	for input.Scan() {
		if len(strings.TrimSpace(input.Text())) == 0 {

			continue
		}
		username = input.Text()
		break;
	}
	return username
}

func getClientLoginPwd(conn net.Conn) (password string)  {
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn,"请输入你需要登录密码：")
	for input.Scan() {
		if len(strings.TrimSpace(input.Text())) == 0 {

			continue
		}
		password = input.Text()
		break;
	}
	return password
}