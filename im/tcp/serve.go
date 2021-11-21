/**
  @author:panliang
  @data:2021/9/26
  @note
**/
package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"im_app/im"
	"im_app/im/service"
	"im_app/pkg/config"
	NewJwt "im_app/pkg/jwt"
	"im_app/pkg/pool"
)

type TcpClient struct {
	ID       int64 // 用户id
	UserName string // 用户名称
	Ch       client // 客户端消息通道
}

// 客户端集合
type client chan<- string

// Client manager
type TcpClientManager struct {
	ClientMap map[int64]*TcpClient
	ch        chan string
}

var Manager = TcpClientManager{
	ClientMap: make(map[int64]*TcpClient),
}

var (
	entering = make(chan TcpClient) // 上线消息
	leaving  = make(chan TcpClient) // 离线消息
	messages = make(chan string)    // 所有接受的客户消息
)

func init() {
	// 加载池
	im.SetupPool()
}

func StartTcpServe() {
	listener, err := net.Listen("tcp", ":"+config.GetString("app.tcp_port"))
	if err != nil {
		log.Fatal(err)
	}
	pool.AntsPool.Submit(func() {
		broadcaster()
	})
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		// 使用协程池管理使用协程
		go handleConn(conn)
	}
}

// 广播器
func broadcaster() {
	clients := Manager.ClientMap
	for {
		select {
		case msg := <-messages:
			// 可以根据投递的消息判断是私聊消息、还是广播消息、还是群聊消息 离线消息可以写入消息中间件
			for name, cli := range clients {
				select {
				case cli.Ch <- msg:
				default:
					fmt.Fprintf(os.Stderr, "发送消息失败: %s: %s\n", name, msg)
				}
			}
		case cliSt := <-entering:
			// 用户上线了
			var users []string
			for _, user := range clients {
				users = append(users, user.UserName)
			}
			if len(users) > 1 {
				cliSt.Ch <- fmt.Sprintf("房间里面的其他用户: %s", strings.Join(users, "; "))
			} else {
				cliSt.Ch <- "你是第一个加入房间的人."
			}
			clients[cliSt.ID].Ch = cliSt.Ch
		case cliSt := <-leaving:
			delete(clients, cliSt.ID)
			close(cliSt.Ch)
		}
	}
}

var (
	err    error
	claims *NewJwt.CustomClaims
)

// 处理连接
func handleConn(conn net.Conn) {
	// 执行登录逻辑操作 读取用户输入账号和密码
	token := getToken(conn)

	jwt := NewJwt.NewJWT()
	claims, err = jwt.ParseToken(token)
	fmt.Println(claims)

	if err != nil {
		data := fmt.Sprintf(`{"code":401,"msg":"%s","errMsg":"%s"}`, "用户身份验证失败", err.Error())
		conn.Write([]byte(data))
		conn.Close()
		return
	}

	fmt.Fprintf(conn, "name %q is existed\r\ntry other name: ", "登录成功")

	// username:= clientRegisterUser(conn)
	// password:= clientRegisterPwd(conn)
	// 执行登录操作
	tcpDao := new(service.TcpDao)
	users, err := tcpDao.GetUser(claims.ID)

	if err != nil {
		conn.Close()
	}
	// 断开清理用户连接
	defer func() {
		conn.Close()
	}()
	// 用户信息存储
	Manager.ClientMap[users.ID] = &TcpClient{ID: users.ID, UserName: users.Name}
	// 客户端消息写入
	ch := make(chan string)
	go clientWriter(conn, ch)
	// 用户信息结构体
	chl := TcpClient{ID: users.ID, UserName: users.Name, Ch: ch}

	ch <- "你是-" + users.Name
	messages <- users.Name + "登录成功"
	entering <- chl
	inputFunc := func(sig chan<- struct{}) {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			sig <- struct{}{}
			if len(strings.TrimSpace(input.Text())) == 0 {
				continue
			}
			messages <- users.Name + ": " + input.Text()
		}
		// 注意，忽略input.Err()中可能的错误
	}
	inputWithTimeout(conn, 300*time.Second, inputFunc)

	leaving <- chl
	messages <- users.Name + " 已经下线了"
}

func getToken(conn net.Conn) (who string) {
	inputFunc := func(sig chan<- struct{}) {
		input := bufio.NewScanner(conn)
		ch := make(chan bool)
		data := fmt.Sprintf(`{"code":1,"msg":"%s"}`, "请输入token")
		conn.Write([]byte(data))
		// fmt.Fprint(conn, "请输入token:") // 注意，忽略网络层面的错误
		for input.Scan() {
			if len(strings.TrimSpace(input.Text())) == 0 { // 禁止发送纯空白字符
				continue
			}
			who = input.Text()
			if <-ch {
				break
			}
			fmt.Fprintf(conn, "name %q is existed\r\ntry other name: ", who)
		}
		// 注意，忽略input.Err()中可能的错误
	}
	inputWithTimeout(conn, 15*time.Second, inputFunc)
	return who
}

// 获取用户账号
func clientRegisterUser(conn net.Conn) (who string) {
	inputFunc := func(sig chan<- struct{}) {
		input := bufio.NewScanner(conn)
		ch := make(chan bool)
		fmt.Fprint(conn, "请输入登录的账号: ") // 注意，忽略网络层面的错误
		for input.Scan() {
			if len(strings.TrimSpace(input.Text())) == 0 { // 禁止发送纯空白字符
				continue
			}
			who = input.Text()
			if <-ch {
				break
			}
			fmt.Fprintf(conn, "name %q is existed\r\ntry other name: ", who)
		}
		// 注意，忽略input.Err()中可能的错误
	}
	inputWithTimeout(conn, 15*time.Second, inputFunc)
	return who
}

// 获取用户密码
func clientRegisterPwd(conn net.Conn) (who string) {
	inputFunc := func(sig chan<- struct{}) {
		input := bufio.NewScanner(conn)
		ch := make(chan bool)
		fmt.Fprint(conn, "请输入登录的密码: ") // 注意，忽略网络层面的错误
		for input.Scan() {
			if len(strings.TrimSpace(input.Text())) == 0 { // 禁止发送纯空白字符
				continue
			}
			who = input.Text()
			if <-ch {
				break
			}
			fmt.Fprintf(conn, "name %q is existed\r\ntry other name: ", who)
		}
		// 注意，忽略input.Err()中可能的错误
	}
	inputWithTimeout(conn, 15*time.Second, inputFunc)
	return who
}

// 超时退出功能
func inputWithTimeout(conn net.Conn, timeout time.Duration, input func(sig chan<- struct{})) {
	done := make(chan struct{}, 2)
	inputSignal := make(chan struct{})
	go func() {
		timer := time.NewTimer(timeout)
		for {
			select {
			case <-inputSignal:
				timer.Reset(timeout)
			case <-timer.C:
				// 超时，断开连接
				done <- struct{}{}
				return
			}
		}
	}()
	pool.AntsPool.Submit(func() {
		input(inputSignal)
		done <- struct{}{}
	})
	<-done
}

// 客户端消息写入
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Println(msg)
		fmt.Fprintln(conn, msg+"\r") // 注意，忽略网络层面的错误
	}
}
