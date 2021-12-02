/**
  @author:panliang
  @data:2021/11/14
  @note
**/
package tests

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"im_app/config"
	"im_app/pkg/jwt"
	"log"
	"math/big"
	"sync"
	"testing"
	"time"
)

func init() {
	config.Initialize()
}

var addr = flag.String("addr", "127.0.0.1:9502", "http service address")

type TestHandler interface {

	Send(id int,msg []byte)
}

var ClientManager =ConnMap{
	client: make(map[int]*Client),
}

type ConnMap struct {
	client map[int]*Client
}
type Client struct {
	Token     string
	Conn *websocket.Conn
	Queue   chan []byte
	Mu sync.Mutex
}


var  numbers = 50000

func TestApp(t *testing.T) {

	for i := 1; i < numbers; i++ {
		ClientManager.start(i)
	}
	fmt.Println("5w连接存储成功:")

	for j := 1; j < numbers; j++ {

		if conn,ok := ClientManager.client[j];ok{
			time.Sleep(time.Microsecond*2)
			wg.Add(1)
			 conn.Send(j)

		}
	}

	wg.Wait()


}

func (c *ConnMap)start(i int)  {
	name := fmt.Sprintf("测试%d",i)
	token := jwt.GenerateToken(int64(i), name, "test", "2540463097@qq.com", 1)

	u := "ws://127.0.0.1:9502/core/connect?token="+token
	fmt.Println(name)
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err,u)
	}

	mutexKey.Lock()
	c.client[i]= &Client{Conn: conn,Token: token,Queue: make(chan []byte,40)}
	mutexKey.Unlock()
	go c.client[i].write() //执行

}
var mutexKey sync.Mutex

func (c *Client)write()  {
	// 关闭socket连接
	defer c.Conn.Close()
	for {
		select {
		case message, ok := <-c.Queue:
			if !ok {
				// 关闭
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Mu.Lock()
			c.Conn.WriteMessage(websocket.TextMessage, message)
			c.Mu.Unlock()
		}
	}
}

//消息推送
func (c *Client)Send(i int)  {


	t_id :=random()
	data := fmt.Sprintf(`{"msg":"%s","from_id":%v,"to_id":%v,"status":0,"msg_type":%v,"channel_type":%v}`,
		"test",i,t_id,  1,1)
	//消息投递
	c.Queue <- []byte(data)
	c.Queue <- []byte(data)
	c.Queue <- []byte(data)
	c.Queue <- []byte(data)
	c.Queue <- []byte(data)
}

func random() int  {
	max := big.NewInt(5000)
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal("rand:", err)
	}
	return i.BitLen()
}

