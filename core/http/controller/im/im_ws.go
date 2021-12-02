/**
  @author:panliang
  @data:2021/8/13
  @note
**/
package im

import (
	"github.com/gin-gonic/gin"
	"im_app/pkg/config"
	"net/http"
	"im_app/core/cache"
	"im_app/pkg/jwt"
	"im_app/core/ws"
	ws2 "im_app/pkg/ws"
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
	client := &ws.ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte) }

	if app_cluster_model {
		var cache cache.ServiceNode
		cache.SetUserServiceNode(claims.ID)
	}

	ws.ImManager.Register <- client

	go client.ImRead()

	go client.ImWrite()
}

