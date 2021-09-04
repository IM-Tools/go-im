/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package im


import (
	"github.com/gin-gonic/gin"
	"go_im/im/service"
	"go_im/pkg/jwt"
	"go_im/pkg/pool"
	"go_im/pkg/ws"
	"net/http"
)

type IMService struct{}

func (*IMService) Connect(c *gin.Context) {
	conn, err := ws.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	client := &service.ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte)}
	service.ImManager.Register <- client

	//开始投递任务
	pool.AntsPool.Submit(func() {
		client.ImRead()
	})

	pool.AntsPool.Submit(func() {
		client.ImWrite()
	})
}




