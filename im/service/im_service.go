package service

import (
	"github.com/gin-gonic/gin"
	"go_im/pkg/jwt"
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
	//client := &ws.Client{ID: claims.ID, Socket: conn, Send: make(chan []byte)}
	client := &ImClient{ID: claims.ID, Socket: conn, Send: make(chan []byte)}
	ImManager.Register <- client
	go client.ImRead()
	go client.ImWrite()
}


