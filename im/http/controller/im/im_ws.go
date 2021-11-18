/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package im

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"im_app/im/cache"
	ws2 "im_app/im/ws"
	client2 "im_app/im/ws_test/client"
	"im_app/pkg/jwt"
	"im_app/pkg/pool"
	"im_app/pkg/ws"
	"net/http"
	"strconv"
)

type IMService struct{}



func (*IMService) Connect(c *gin.Context) {
	conn, err := ws.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	client := &ws2.ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte)}

	ID ,err := strconv.Atoi(claims.ID)
	if err != nil{
		fmt.Println(err)
	}
	cache.SetUserServiceNode(int64(ID))

	ws2.ImManager.Register <- client

	//开始投递任务
	pool.AntsPool.Submit(func() {
		client.ImRead()
	})

	pool.AntsPool.Submit(func() {
		client.ImWrite()
	})
}


func (*IMService) Connect2(c *gin.Context) {
	conn, err := ws.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	id,_ := strconv.Atoi(claims.ID)

	client := new(client2.Client)

	clients := client.NewClient(int64(id),conn)

	Container :=new(client2.ClientContainer)
	Container.SetDevice(int64(id),claims.ClientType,clients)

}

