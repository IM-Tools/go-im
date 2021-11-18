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
	"testing"
)

func init() {
	config.Initialize()
}

var addr = flag.String("addr", "127.0.0.1:9502", "http service address")

type ConnMap struct {
	client map[int]*Client
}
type Client struct {
	Token string
	Conn  *websocket.Conn
}



func TestApp(t *testing.T) {
	//ConnMap := new(ConnMap)
	for i := 2; i < 500; i++ {
		name := fmt.Sprintf("测试%d", i)
		token := jwt.GenerateToken(uint64(i), name, "test", "2540463097@qq.com", 1)
		client := new(Client)
		client.Token = token
		u := "ws://127.0.0.1:9502/im/connect?token="+token
		c, _, err := websocket.DefaultDialer.Dial(u, nil)
		if err != nil {
			log.Fatal("dial:", err,u)
		}
		client.Conn = c
		data := fmt.Sprintf(`{"msg":"%s","from_id":%v,"to_id":%v,"status":0,"msg_type":%v}`,
			"test", i,30,  1)
		client.Conn.WriteMessage(websocket.TextMessage,[]byte(data))
		//jsonMessage, _ := json.Marshal(&ws.Message{Sender: strconv.Itoa(i), Content: string(data)})
		//ws.ImManager.Broadcast <- jsonMessage
		//ConnMap.client[i] = client
	}

	//for _,cli :=range ConnMap.client {
	//	data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
	//		"test", random(), random(), 1, 1)
	//	go cli.Conn.WriteMessage(websocket.TextMessage,[]byte(data))
	//}

}

func random() int  {
	max := big.NewInt(100)
	i, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal("rand:", err)
	}
	return i.BitLen()
}

