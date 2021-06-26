/**
  @author:panliang
  @data:2021/6/25
  @note
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	NewJwt "go_im/pkg/jwt"
	"net/http"
)

func WsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusForbidden, map[string]interface{}{
				"code":403,
				"msg":"token不能为空",
			})
			c.Abort()
			return
		} else {
			//开始鉴权
			jwt := NewJwt.NewJWT()
			claims,err := jwt.ParseToken(token)
			if err != nil {

				if err == NewJwt.TokenExpired {
					c.JSON(http.StatusOK, gin.H{
						"status": 500,
						"msg":err.Error(),
					})
				} else {
					c.JSON(http.StatusForbidden, map[string]interface{}{
						"code":500,
						"msg":err.Error(),
					})
				}
				c.Abort()
				return
			}
			c.Set("claims", claims)
		}
	}
}

