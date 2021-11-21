/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"im_app/im/cache"
	"im_app/pkg/jwt"
	"im_app/im/ws"
	ws2 "im_app/pkg/ws"
)


type IMService struct{}

func (*IMService) Connect(c *gin.Context) {
	conn, err := ws2.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	claims := c.MustGet("claims").(*jwt.CustomClaims)
	client := &ws.ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte)}


	var cache cache.ServiceNode

	cache.SetUserServiceNode(claims.ID)

	ws.ImManager.Register <- client

	// 开始投递任务
	//pool.AntsPool.Submit(func() {
	//
	//})
	go client.ImRead()

	go client.ImWrite()

	//pool.AntsPool.Submit(func() {
	//
	//})
}

