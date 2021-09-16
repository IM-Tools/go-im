/**
  @author:panliang
  @data:2021/9/3
  @note
**/
package tcp

import (
	"bufio"
	"fmt"
	userModel "go_im/im/http/models/user"
	"go_im/pkg/helpler"
	"go_im/pkg/model"
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
	login(conn)
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

func login(conn net.Conn)  {
	username :=getClientLoginMsg("username",conn)
	password :=getClientLoginMsg("password",conn)

	var users userModel.Users
	model.DB.Model(&userModel.Users{}).Where("name = ?",username).Find(&users)
	if users.ID == 0 {
		log.Fatal("用户不存在")
	}
	if !helpler.ComparePasswords(users.Password, password) {
		log.Fatal("账号或者密码错误")
	}


}


func mustCopy(dst io.Writer,src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// 获取登录信息

func getClientLoginMsg(filed string,conn net.Conn) (who string)  {
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn,"input "+filed+":")
	for input.Scan() {
		if len(strings.TrimSpace(input.Text())) == 0 {
			continue
		}
		who = input.Text()
		break;
	}

	return who
}