/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"im_app/core/cache"
	"im_app/core/ws"
	"im_app/pkg/config"
	"im_app/pkg/jwt"
	ws2 "im_app/pkg/ws"
	"net/http"
	"time"
)

var app_cluster_model = config.GetBool("core.app_cluster_model")

type IMService struct{}

func (*IMService) Connect(c *gin.Context) {
	conn, err := ws2.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	claims := c.MustGet("claims").(*jwt.CustomClaims)

	//&ws.Devices{Socket: conn}
	client := &ws.ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte)}

	if app_cluster_model {
		var cache cache.ServiceNode
		cache.SetUserServiceNode(claims.ID)
	}

	ws.ImManager.Register <- client

	go client.ImRead()

	go client.ImWrite()
}

func (*IMService) CustomerService(c *gin.Context) {
	conn, err := ws2.App(c.Writer, c.Request)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	mid := c.Query("mid")

	if len(mid) == 0 {
		conn.WriteMessage(1, []byte(string("参数有误--连接中断")))
		conn.Close()
		return
	} else {
		timeInt := time.Now().Unix()
		//&ws.Devices{Socket: conn}
		client := &ws.ImClient{ID: timeInt, Socket: conn, Send: make(chan []byte)}

		ws.ImManager.Register <- client

		go client.ImRead()

		go client.ImWrite()
	}

}
