/**
  @author:panliang
  @data:2021/6/22
  @note
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	NewJwt "go_im/pkg/jwt"
	"go_im/pkg/response"
)

//路由中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			response.FailResponse(403,"token不能为空").ToJson(c)
		} else {
			//开始鉴权
			jwt := NewJwt.NewJWT()
			claims,err := jwt.ParseToken(token)
			if err != nil {
				if err == NewJwt.TokenExpired {
					response.FailResponse(500,err.Error()).ToJson(c)
				} else {
					response.FailResponse(500,err.Error()).ToJson(c)
				}
				c.Abort()
				return
			}
			c.Set("claims", claims)
		}
	}
}

